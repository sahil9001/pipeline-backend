import grpc from "k6/net/grpc";
import http from "k6/http";

import { check, group } from "k6";

import * as pipeline from "./grpc-pipeline-public.js";
import * as pipelineWithJwt from "./grpc-pipeline-public-with-jwt.js";
import * as pipelinePrivate from "./grpc-pipeline-private.js";
import * as trigger from "./grpc-trigger.js";
import * as triggerAsync from "./grpc-trigger-async.js";

const client = new grpc.Client();

client.load(["proto/vdp/pipeline/v1alpha"], "pipeline_public_service.proto");
client.load(["proto/vdp/connector/v1alpha"], "connector_public_service.proto");

import * as constant from "./const.js";

export let options = {
  setupTimeout: "300s",
  insecureSkipTLSVerify: true,
  thresholds: {
    checks: ["rate == 1.0"],
  },
};

export function setup() {
  client.connect(constant.connectorGRPCPublicHost, {
    plaintext: true,
    timeout: "10s",
  });

  var loginResp = http.request("POST", `${constant.mgmtPublicHost}/v1alpha/auth/login`, JSON.stringify({
    "username": constant.defaultUsername,
    "password": constant.defaultPassword,
  }))

  check(loginResp, {
    [`POST ${constant.mgmtPublicHost}/v1alpha/auth/login response status is 200`]: (
      r
    ) => r.status === 200,
  });

  var metadata = {
    "metadata": {
      "Authorization": `Bearer ${loginResp.json().access_token}`
    },
    "timeout": "600s",
  }

  group(
    "Connector Backend API: Create a CSV destination connector 1",
    function () {
      check(
        client.invoke(
          "vdp.connector.v1alpha.ConnectorPublicService/CreateUserConnectorResource",
          {
            parent: `${constant.namespace}`,
            connector_resource: {
              id: constant.dstCSVConnID1,
              connector_definition_name:
                "connector-definitions/airbyte-destination-csv",
              configuration: {
                destination_path: "/local/pipeline-backend-test-1",
              },
            },
          },
          metadata
        ),
        {
          "vdp.connector.v1alpha.ConnectorPublicService/CreateUserConnectorResource CSV response StatusOK":
            (r) => r.status === grpc.StatusOK,
        }
      );
      client.invoke(
        "vdp.connector.v1alpha.ConnectorPublicService/ConnectUserConnectorResource",
        {
          name: `${constant.namespace}/connector-resources/${constant.dstCSVConnID1}`,
        },
        metadata
      );
    }
  );
  group(
    "Connector Backend API: Create a CSV destination connector 2",
    function () {
      check(
        client.invoke(
          "vdp.connector.v1alpha.ConnectorPublicService/CreateUserConnectorResource",
          {
            parent: `${constant.namespace}`,
            connectorResource: {
              id: constant.dstCSVConnID2,
              connector_definition_name:
                "connector-definitions/airbyte-destination-csv",
              configuration: {
                destination_path: "/local/pipeline-backend-test-2",
              },
            },
          },
          metadata
        ),
        {
          "vdp.connector.v1alpha.ConnectorPublicService/CreateUserConnectorResource CSV response StatusOK":
            (r) => r.status === grpc.StatusOK,
        }
      );
      client.invoke(
        "vdp.connector.v1alpha.ConnectorPublicService/ConnectUserConnectorResource",
        {
          name: `${constant.namespace}/connector-resources/${constant.dstCSVConnID2}`,
        },
        metadata
      );
    }
  );

  client.close();
  return metadata
}

export default function (metadata) {
  /*
   * Pipelines API - API CALLS
   */

  // Health check
  {
    group("Pipelines API: Health check", () => {
      client.connect(constant.pipelineGRPCPublicHost, {
        plaintext: true,
      });
      check(
        client.invoke(
          "vdp.pipeline.v1alpha.PipelinePublicService/Liveness",
          {}
        ),
        {
          "GET /health/pipeline response status is StatusOK": (r) =>
            r.status === grpc.StatusOK,
        }
      );
      client.close();
    });
  }

  if (!constant.apiGatewayMode) {
    pipelinePrivate.CheckList(metadata)
    pipelinePrivate.CheckLookUp(metadata)

  } else {
    pipelineWithJwt.CheckCreate(metadata)
    pipelineWithJwt.CheckList(metadata)
    pipelineWithJwt.CheckGet(metadata)
    pipelineWithJwt.CheckUpdate(metadata)
    pipelineWithJwt.CheckRename(metadata)
    pipelineWithJwt.CheckLookUp(metadata)
    pipeline.CheckCreate(metadata)
    pipeline.CheckList(metadata)
    pipeline.CheckGet(metadata)
    pipeline.CheckUpdate(metadata)
    pipeline.CheckRename(metadata)
    pipeline.CheckLookUp(metadata)

    trigger.CheckTrigger(metadata);
    triggerAsync.CheckTrigger(metadata);
  }
}

export function teardown(metadata) {
  group("Pipeline API: Delete all pipelines created by this test", () => {
    client.connect(constant.pipelineGRPCPublicHost, {
      plaintext: true,
    });

    for (const pipeline of client.invoke(
      "vdp.pipeline.v1alpha.PipelinePublicService/ListUserPipelines",
      {
        parent: `${constant.namespace}`,
        pageSize: 1000,
      },
      metadata
    ).message.pipelines) {
      check(
        client.invoke(
          `vdp.pipeline.v1alpha.PipelinePublicService/DeleteUserPipeline`,
          {
            name: `${constant.namespace}/pipelines/${pipeline.id}`,
          },
          metadata
        ),
        {
          [`vdp.pipeline.v1alpha.PipelinePublicService/DeleteUserPipeline response StatusOK`]:
            (r) => r.status === grpc.StatusOK,
        }
      );
    }

    client.close();
  });

  client.connect(constant.connectorGRPCPublicHost, {
    plaintext: true,
  });

  group(
    "Connector Backend API: Delete the csv destination connector 1",
    function () {
      check(
        client.invoke(
          `vdp.connector.v1alpha.ConnectorPublicService/DeleteUserConnectorResource`,
          {
            name: `${constant.namespace}/connector-resources/${constant.dstCSVConnID1}`,
          },
          metadata
        ),
        {
          [`vdp.connector.v1alpha.ConnectorPublicService/DeleteUserConnectorResource response StatusOK`]:
            (r) => r.status === grpc.StatusOK,
        }
      );
    }
  );
  group(
    "Connector Backend API: Delete the csv destination connector 2",
    function () {
      check(
        client.invoke(
          `vdp.connector.v1alpha.ConnectorPublicService/DeleteUserConnectorResource`,
          {
            name: `${constant.namespace}/connector-resources/${constant.dstCSVConnID2}`,
          },
          metadata
        ),
        {
          [`vdp.connector.v1alpha.ConnectorPublicService/DeleteUserConnectorResource response StatusOK`]:
            (r) => r.status === grpc.StatusOK,
        }
      );
    }
  );

  client.close();
}
