//go:generate compogen readme ./config ./README.mdx --extraContents TASK_JQ=.compogen/extra-jq.mdx --extraContents bottom=.compogen/bottom.mdx
package json

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	_ "embed"

	"github.com/itchyny/gojq"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/instill-ai/pipeline-backend/pkg/component/base"
	"github.com/instill-ai/x/errmsg"
)

const (
	taskMarshal   = "TASK_MARSHAL"
	taskUnmarshal = "TASK_UNMARSHAL"
	taskJQ        = "TASK_JQ"
)

var (
	//go:embed config/definition.json
	definitionJSON []byte
	//go:embed config/tasks.json
	tasksJSON []byte

	once sync.Once
	comp *component
)

type component struct {
	base.Component
}

type execution struct {
	base.ComponentExecution

	execute func(*structpb.Struct) (*structpb.Struct, error)
}

// Init returns an implementation of IOperator that processes JSON objects.
func Init(bc base.Component) *component {
	once.Do(func() {
		comp = &component{Component: bc}
		err := comp.LoadDefinition(definitionJSON, nil, tasksJSON, nil)
		if err != nil {
			panic(err)
		}
	})
	return comp
}

// CreateExecution initializes a connector executor that can be used in a
// pipeline trigger.
func (c *component) CreateExecution(x base.ComponentExecution) (base.IExecution, error) {
	e := &execution{ComponentExecution: x}

	switch x.Task {
	case taskMarshal:
		e.execute = e.marshal
	case taskUnmarshal:
		e.execute = e.unmarshal
	case taskJQ:
		e.execute = e.jq
	default:
		return nil, errmsg.AddMessage(
			fmt.Errorf("not supported task: %s", x.Task),
			fmt.Sprintf("%s task is not supported.", x.Task),
		)
	}
	return e, nil
}

func (e *execution) marshal(in *structpb.Struct) (*structpb.Struct, error) {
	out := new(structpb.Struct)

	b, err := protojson.Marshal(in.Fields["json"])
	if err != nil {
		return nil, errmsg.AddMessage(err, "Couldn't convert the provided object to JSON.")
	}

	out.Fields = map[string]*structpb.Value{
		"string": structpb.NewStringValue(string(b)),
	}

	return out, nil
}

func (e *execution) unmarshal(in *structpb.Struct) (*structpb.Struct, error) {
	out := new(structpb.Struct)

	b := []byte(in.Fields["string"].GetStringValue())
	obj := new(structpb.Value)
	if err := protojson.Unmarshal(b, obj); err != nil {
		return nil, errmsg.AddMessage(err, "Couldn't parse the JSON string. Please check the syntax is correct.")
	}

	out.Fields = map[string]*structpb.Value{"json": obj}

	return out, nil
}

func (e *execution) jq(in *structpb.Struct) (*structpb.Struct, error) {
	out := new(structpb.Struct)

	input := in.Fields["json-value"].AsInterface()
	if input == nil {
		b := []byte(in.Fields["json-string"].GetStringValue())
		if err := json.Unmarshal(b, &input); err != nil {
			return nil, errmsg.AddMessage(err, "Couldn't parse the JSON input. Please check the syntax is correct.")
		}
	}

	queryStr := in.Fields["jq-filter"].GetStringValue()
	q, err := gojq.Parse(queryStr)
	if err != nil {
		// Error messages from gojq are human-friendly enough.
		msg := fmt.Sprintf("Couldn't parse the jq filter: %s. Please check the syntax is correct.", err.Error())
		return nil, errmsg.AddMessage(err, msg)
	}

	results := []any{}
	iter := q.Run(input)
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}

		if err, ok := v.(error); ok {
			msg := fmt.Sprintf("Couldn't apply the jq filter: %s.", err.Error())
			return nil, errmsg.AddMessage(err, msg)
		}

		results = append(results, v)
	}

	list, err := structpb.NewList(results)
	if err != nil {
		return nil, err
	}

	out.Fields = map[string]*structpb.Value{
		"results": structpb.NewListValue(list),
	}

	return out, nil
}

func (e *execution) Execute(ctx context.Context, jobs []*base.Job) error {
	return base.SequentialExecutor(ctx, jobs, e.execute)
}
