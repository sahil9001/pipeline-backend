import http from "k6/http";
import { check, group } from "k6";
import { randomString } from "https://jslib.k6.io/k6-utils/1.1.0/index.js";

import {
  pipelinePublicHost,
  defaultUsername,
  dbIDPrefix
} from "./const.js";

import { deepEqual } from "./helper.js";

const defaultPageSize = 10;

export function CheckIntegrations() {
  group("Integration API: Get integration", () => {
    // Inexistent component
    check(http.request("GET", `${pipelinePublicHost}/v1beta/integrations/restapio`, null, null), {
      "GET /v1beta/integrations/restapio response status is 404": (r) => r.status === 404,
      "GET /v1beta/integrations/restapio response contains end-user message": (r) => r.json().message === "Integration does not exist.",
    });

    // Component without setup
    check(http.request("GET", `${pipelinePublicHost}/v1beta/integrations/document`, null, null), {
      "GET /v1beta/integrations/document response status is 404": (r) => r.status === 404,
      "GET /v1beta/integrations/document response contains end-user message": (r) => r.json().message === "Integration does not exist.",
    });

    var id = "slack";
    var cdef = http.request("GET", `${pipelinePublicHost}/v1beta/connector-definitions/${id}`, null, null).
      json().connectorDefinition;

    var integration = {
      uid: cdef.uid,
      id: cdef.id,
      title: cdef.title,
      description: cdef.description,
      vendor: cdef.vendor,
      icon: cdef.icon,
      setupSchema: null,
      schemas: [], // Deprecated
      view: "VIEW_BASIC"
    };

    var oAuthConfig = {
      authUrl: "https://slack.com/oauth/v2/authorize",
      accessUrl: "https://slack.com/api/oauth.v2.access",
      scopes: []
    };

    // Basic view
    check(http.request("GET", `${pipelinePublicHost}/v1beta/integrations/${id}`, null, null), {
      [`GET /v1beta/integrations/${id} response status is 200`]: (r) => r.status === 200,
      [`GET /v1beta/integrations/${id} response contains expected integration`]: (r) => deepEqual(r.json().integration, integration),
    });

    // Full view
    check(http.request("GET", `${pipelinePublicHost}/v1beta/integrations/${id}?view=VIEW_FULL`, null, null), {
      [`GET /v1beta/integrations/${id}?view=VIEW_FULL response status is 200`]: (r) => r.status === 200,
      [`GET /v1beta/integrations/${id}?view=VIEW_FULL response contains schema`]: (r) => r.json().integration.setupSchema.required[0] === "token",
      [`GET /v1beta/integrations/${id}?view=VIEW_FULL response contains OAuth config`]: (r) => deepEqual(r.json().integration.oAuthConfig, oAuthConfig),
      // Deprecated
      [`DEPRECATED GET /v1beta/integrations/${id}?view=VIEW_FULL response contains schema`]: (r) => r.json().integration.schemas[0].method === "METHOD_DICTIONARY",
    });
  });

  group("Integration API: List integrations", () => {
    // Default pagination.
    var firstPage = http.request("GET", `${pipelinePublicHost}/v1beta/integrations`, null, null);
    check(firstPage, {
      "GET /v1beta/integrations response status is 200": (r) => r.status === 200,
      "GET /v1beta/integrations response totalSize > 0": (r) => r.json().totalSize > 0,
      "GET /v1beta/integrations has default page size": (r) => r.json().integrations.length === defaultPageSize,
    });

    // Non-default pagination, non-first page
    var tokenPageTwo = firstPage.json().nextPageToken;
    check(http.request("GET", `${pipelinePublicHost}/v1beta/integrations?pageSize=2&pageToken=${tokenPageTwo}`, null, null), {
      [`GET /v1beta/integrations?pageSize=2&pageToken=${tokenPageTwo} response status is 200`]: (r) => r.status === 200,
      [`GET /v1beta/integrations?pageSize=2&pageToken=${tokenPageTwo} has page size 2"`]: (r) => r.json().integrations.length === 2,
      [`GET /v1beta/integrations?pageSize=2&pageToken=${tokenPageTwo} has different elements than page 1"`]: (r) =>
        r.json().integrations[0].id != firstPage.json().integrations[0].id,
    });

    // Filter fuzzy title
    check(http.request("GET", `${pipelinePublicHost}/v1beta/integrations?filter=qIntegration="que"`, null, null), {
      [`GET /v1beta/integrations?filter=qIntegration="que" response status is 200`]: (r) => r.status === 200,
      [`GET /v1beta/integrations?filter=qIntegration="que" response totalSize > 0`]: (r) => r.json().totalSize === 1,
      [`GET /v1beta/integrations?filter=qIntegration="que" returns BigQuery integration`]: (r) => r.json().integrations[0].title === "BigQuery",
    });

    // Filter fuzzy vendor
    check(http.request("GET", `${pipelinePublicHost}/v1beta/integrations?filter=qIntegration="labs"`, null, null), {
      [`GET /v1beta/integrations?filter=qIntegration="labs" response status is 200`]: (r) => r.status === 200,
      [`GET /v1beta/integrations?filter=qIntegration="labs" response totalSize > 0`]: (r) => r.json().totalSize === 1,
      [`GET /v1beta/integrations?filter=qIntegration="labs" returns Redis integration`]: (r) => r.json().integrations[0].title === "Redis",
      [`GET /v1beta/integrations?filter=qIntegration="labs" intgration vendor matches Redis Labs`]: (r) => r.json().integrations[0].vendor === "Redis Labs",
    });
  });
}

export function CheckConnections(data) {
  var connectionID = dbIDPrefix + randomString(8);
  var collectionPath = `/v1beta/namespaces/${defaultUsername}/connections`;
  var resourcePath = `${collectionPath}/${connectionID}`;
  var integrationID = "slack";

  var setup = {"token": "one2THREE"};

  group("Integration API: Create connection", () => {
    var path = collectionPath;

    // Successful creation: dictionary
    var dictReq = http.request(
      "POST",
      pipelinePublicHost + path,
      JSON.stringify({
        id: connectionID,
        integrationId: integrationID,
        method: "METHOD_DICTIONARY",
        setup: setup,
      }),
      data.header
    );
    check(dictReq, {
      [`POST ${path} (dictionary) response status is 201`]: (r) => r.status === 201,
      [`POST ${path} (dictionary) has a UID`]: (r) => r.json().connection.uid.length > 0,
      [`POST ${path} (dictionary) has a creation time`]: (r) => new Date(r.json().connection.createTime).getTime() > new Date().setTime(0),
    });

    // Successful creation: dictionary
    var oAuthReq = http.request(
      "POST",
      pipelinePublicHost + path,
      JSON.stringify({
        id: connectionID + "-oauth",
        integrationId: integrationID,
        method: "METHOD_OAUTH",
        setup: setup,
        scopes: ["commands", "incoming-webhook"],
        oAuthAccessDetails: {
          ok: true,
          access_token: "xoxb-17653672481-19874698323-pdFZKVeTuE8sk7oOcBrzbqgy",
          token_type: "bot",
          scope: "commands,incoming-webhook",
          bot_user_id: "U0KRQLJ9H",
          app_id: "A0KRD7HC3",
          team: {
              name: "Slack Pickleball Team",
              id: "T9TK3CUKW"
          },
          enterprise: {
              name: "slack-pickleball",
              id: "E12345678"
          },
          authed_user: {
              id: "U1234",
              scope: "chat:write",
              access_token: "xoxp-1234",
              token_type: "user"
          }
        }
      }),
      data.header
    );
    check(oAuthReq, {
      [`POST ${path} (OAuth) response status is 201`]: (r) => r.status === 201,
      [`POST ${path} (OAuth) has a UID`]: (r) => r.json().connection.uid.length > 0,
      [`POST ${path} (OAuth) has a creation time`]: (r) => new Date(r.json().connection.createTime).getTime() > new Date().setTime(0),
    });

    // Check ID format
    var invalidID = dbIDPrefix + "This-Is-Invalid";
    var invalidIDReq = http.request(
      "POST",
      pipelinePublicHost + path,
      JSON.stringify({
        id: invalidID,
        integrationId: integrationID,
        method: "METHOD_DICTIONARY",
        setup: {"token": "foo"},
      }),
      data.header
    );
    check(invalidIDReq, {
      [`POST ${path} response status is 400 with ID ${invalidID}`]: (r) => r.status === 400,
    });

    var invalidSetupReq = http.request(
      "POST",
      pipelinePublicHost + path,
      JSON.stringify({
        id: dbIDPrefix + randomString(16),
        integrationId: integrationID,
        method: "METHOD_DICTIONARY",
        setup: {"token": 234},
      }),
      data.header
    );
    check(invalidIDReq, {
      [`POST ${path} response status is 400 with invalid setup`]: (r) => r.status === 400,
    });

    var invalidDictReq = http.request(
      "POST",
      pipelinePublicHost + path,
      JSON.stringify({
        id: dbIDPrefix + randomString(16),
        integrationId: integrationID,
        method: "METHOD_DICTIONARY",
        setup: {"token": "valid"},
        scopes: ["commands", "incoming-webhook"],
      }),
      data.header
    );
    check(invalidIDReq, {
      [`POST ${path} response status is 400 with invalid method payload`]: (r) => r.status === 400,
    });
  });

  group("Integration API: Get connection", () => {
    var path = resourcePath + "-oauth";

    check(http.request("GET", pipelinePublicHost + path + "aaa", null, data.header), {
      [`GET ${path + "aaa"} response status is 404`]: (r) => r.status === 404,
    });

    // Basic view
    check(http.request("GET", pipelinePublicHost + path, null, data.header), {
      [`GET ${path} response status is 200`]: (r) => r.status === 200,
      [`GET ${path} has basic view`]: (r) => r.json().connection.view === "VIEW_BASIC",
      [`GET ${path} has setup hidden`]: (r) => r.json().connection.setup === null,
      [`GET ${path} has integration ID`]: (r) => r.json().connection.integrationId === integrationID,
      [`GET ${path} has integration title`]: (r) => r.json().connection.integrationTitle === "Slack",
    });

    // Full view
    check(http.request("GET", pipelinePublicHost + path + "?view=VIEW_FULL", null, data.header), {
      [`GET ${path + "?view=VIEW_FULL"} response status is 200`]: (r) => r.status === 200,
      [`GET ${path + "?view=VIEW_FULL"} has full view`]: (r) => r.json().connection.view === "VIEW_FULL",
      [`GET ${path + "?view=VIEW_FULL"} has setup`]: (r) => r.json().connection.setup != null,
      [`GET ${path + "?view=VIEW_FULL"} has setup value`]: (r) => r.json().connection.setup.password === setup.password, // TODO: redact
      [`GET ${path + "?view=VIEW_FULL"} has scopes`]: (r) => r.json().connection.scopes.length > 0,
      [`GET ${path + "?view=VIEW_FULL"} has OAuth details`]: (r) => r.json().connection.oAuthAccessDetails.access_token.length > 0, // TODO redact
    });
  });

  group("Integration API: List connections", () => {
    var path = collectionPath;
    var nConnections = 12;
    // Connections have been created in previous tests.
    var totalConnections = nConnections + 2;
    var integrationID = "openai";

    for (var i = 0; i < nConnections; i++) {
      var req = http.request(
        "POST",
        pipelinePublicHost + path,
        JSON.stringify({
          id: dbIDPrefix + randomString(8),
          integrationId: integrationID,
          method: "METHOD_DICTIONARY",
          setup: {
            "api-key": randomString(16),
          },
        }),
        data.header
      );
      check(req, { [`POST ${path}[${i}] response status is 201`]: (r) => r.status === 201 });
    }


    // With connection ID filter
    var pathWithFilter =  path + `?filter=qConnection="${dbIDPrefix}"`;
    var firstPage = http.request("GET", pipelinePublicHost + pathWithFilter, null, data.header);
    check(firstPage, {
      [`GET ${pathWithFilter} response status is 200`]: (r) => r.status === 200,
      [`GET ${pathWithFilter} response has totalSize = ${totalConnections}`]: (r) =>
        r.json().totalSize === totalConnections,
      [`GET ${pathWithFilter} response has default page size`]: (r) =>
        r.json().connections.length === defaultPageSize,
    });

    var pathWithToken = pathWithFilter + `&pageToken=${firstPage.json().nextPageToken}`;
    check(http.request("GET", pipelinePublicHost + pathWithToken, null, data.header), {
      [`GET ${pathWithToken} response status is 200`]: (r) => r.status === 200,
      [`GET ${pathWithToken} response has totalSize = ${totalConnections}`]: (r) =>
        r.json().totalSize === totalConnections,
      [`GET ${pathWithToken} response has remaining items`]: (r) =>
        r.json().connections.length === totalConnections - defaultPageSize,
      [`GET ${pathWithToken} response has no more pages`]: (r) => r.json().nextPageToken === "",
    });

    // With integration ID filter
    var pathWithIntegration = pathWithFilter + `%20AND%20integrationId='${integrationID}'`;
    check(http.request("GET", pipelinePublicHost + pathWithIntegration, null, data.header), {
      [`GET ${pathWithIntegration} response status is 200`]: (r) => r.status === 200,
      [`GET ${pathWithIntegration} response has totalSize = ${nConnections}`]: (r) =>
        r.json().totalSize === nConnections,
      [`GET ${pathWithIntegration} response contains connections for ${integrationID} integration`]: (r) =>
        r.json().connections[0].integrationId === integrationID,
    });
  });

  group("Integration API: List pipelines by connection", () => {
    const yamlRecipe = `
version: v1beta
variable:
  recipients:
    instill-format: array:string
output:
  resp:
    title: Response
    value: \${email-0.output.result}
component:
  email-0:
    type: email
    input:
      recipients: \${variable.recipients}
      cc: null
      bcc: null
      subject: "Dummy email"
      message: "Hi I'm testing integrations"
    condition: null
    task: TASK_SEND_EMAIL
    setup: \${connection.${connectionID}}
`;

    const nPipelines = 30;
    for (var i = 0; i < nPipelines; i++) {
      var id = dbIDPrefix + randomString(8);
      if (i == 0) {
        id = dbIDPrefix + "foobar";
      }
      var reqBody = {
        id: id,
        description: randomString(10),
        rawRecipe: yamlRecipe,
      };


      check(
        http.request(
          "POST",
          `${pipelinePublicHost}/v1beta/namespaces/${defaultUsername}/pipelines`,
          JSON.stringify(reqBody),
          data.header
        ),
        {
          [`POST /v1beta/namespaces/${defaultUsername}/pipelines ${reqBody.id} response status is 201`]:
          (r) => r.status === 201,
        }
      );
    }

    var path = resourcePath + `/referenced-pipelines?pageSize=${nPipelines - 5}`;
    var firstPage = http.request("GET", pipelinePublicHost + path, null, data.header);
    check(firstPage, {
      [`GET ${path} response status is 200`]: (r) => r.status === 200,
      [`GET ${path} response has totalSize = ${nPipelines}`]: (r) => r.json().totalSize === nPipelines,
      [`GET ${path} response has page size ${nPipelines - 5}`]: (r) => r.json().pipelineIds.length === nPipelines - 5,
    });

    var pathWithToken = path + `&pageToken=${firstPage.json().nextPageToken}`;
    check(http.request("GET", pipelinePublicHost + pathWithToken, null, data.header), {
      [`GET ${pathWithToken} response status is 200`]: (r) => r.status === 200,
      [`GET ${pathWithToken} response has totalSize = ${nPipelines}`]: (r) => r.json().totalSize === nPipelines,
      [`GET ${pathWithToken} response has remaining items`]: (r) => r.json().pipelineIds.length === 5,
      [`GET ${pathWithToken} response has no more pages`]: (r) => r.json().nextPageToken === "",
    });

    var pathWithFilter = path + `&filter=q="fooba"`;
    check(http.request("GET", pipelinePublicHost + pathWithFilter, null, data.header), {
      [`GET ${pathWithToken} response status is 200`]: (r) => r.status === 200,
      [`GET ${pathWithToken} response has totalSize = 1`]: (r) => r.json().totalSize === 1,
    });
  });

  group("Integration API: Update connection", () => {
    var path = resourcePath;
    var originalConn = http.request(
      "GET",
      pipelinePublicHost + path,
      null,
      data.header
    ).json().connection;

    var newToken = "new-token";
    var newID = dbIDPrefix + "my-new-id";
    var newMethod = "METHOD_OAUTH";
    var scopes = ["foo"];

    var req = http.request(
      "PATCH",
      pipelinePublicHost + path,
      JSON.stringify({
        uid: "should-be-ignored",
        method: newMethod,
        scopes: scopes,
        id: newID,
        // Fields with an underlying structpb.Struct type (setup,
        // oAuthAccessDetails) will be updated in block.
        setup: {"token": newToken},
        oAuthAccessDetails: {
          token_type: "bot",
          enterprise: {
            name: "slack-pickleball",
            id: "E12345678"
          },
        },
      }),
      data.header
    );

    check(req, {
      [`PATCH ${path} response status 200`]: (r) => r.status === 200,
      [`PATCH ${path} contains new ID`]: (r) => r.json().connection.id === newID,
      [`PATCH ${path} contains new method`]: (r) => r.json().connection.method === newMethod,
      [`PATCH ${path} contains new setup`]: (r) => r.json().connection.setup.token === newToken,
      [`PATCH ${path} contains scopes`]: (r) => r.json().connection.scopes[0] === scopes[0],
      [`PATCH ${path} didn't modify UID`]: (r) => r.json().connection.uid === originalConn.uid,
    });

    resourcePath = `${collectionPath}/${newID}`;
    path = resourcePath;

    check(http.request("GET", pipelinePublicHost + path + "?view=VIEW_FULL", null, data.header), {
      [`GET ${path + "?view=VIEW_FULL"} response status is 200`]: (r) => r.status === 200,
      [`GET ${path + "?view=VIEW_FULL"} has new ID value`]: (r) => r.json().connection.id === newID,
      [`GET ${path + "?view=VIEW_FULL"} has new method`]: (r) => r.json().connection.method === newMethod,
      [`GET ${path + "?view=VIEW_FULL"} has scopes`]: (r) => r.json().connection.setup.token === newToken,
      [`GET ${path + "?view=VIEW_FULL"} has scopes`]: (r) =>
        r.json().connection.scopes[0] === scopes[0],
      [`GET ${path + "?view=VIEW_FULL"} has new setup value`]: (r) =>
        r.json().connection.setup.token === newToken,
    });
  });

  group("Integration API: Delete connection", () => {
    var path = resourcePath;
    check(http.request("DELETE", pipelinePublicHost + path, null, data.header), {
      [`DELETE ${path} response status 204`]: (r) => r.status === 204,
    });

    check(http.request("GET", pipelinePublicHost + path, null, data.header), {
      [`GET ${path} response status is 404`]: (r) => r.status === 404,
    });

    check(http.request("DELETE", pipelinePublicHost + path, null, data.header), {
      [`DELETE ${path} response status 404`]: (r) => r.status === 404,
    });
  });
}
