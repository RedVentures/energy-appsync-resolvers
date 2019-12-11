package resolvers

import (
	"context"
	"encoding/json"
	"reflect"
)

type resolver struct {
	function interface{}
}

func (r *resolver) hasContext() bool {
	return reflect.TypeOf(r.function).NumIn() == 2
}

func (r *resolver) hasPayload() bool {
	return reflect.TypeOf(r.function).NumIn() > 0
}

func (r *resolver) call(ctx context.Context, p json.RawMessage) (interface{}, error) {
	args := make([]reflect.Value, 0, 2)
	hasContext := r.hasContext()

	if hasContext {
		args = append(args, reflect.ValueOf(ctx))
	}

	if r.hasPayload() {
		var index int
		if hasContext {
			index = 1
		}

		pld := payload{p}
		val, err := pld.parse(reflect.TypeOf(r.function).In(index))
		if err != nil {
			return nil, err
		}

		args = append(args, val)
	}

	returnValues := reflect.ValueOf(r.function).Call(args)
	var returnData interface{}
	var returnError error

	if len(returnValues) == 2 {
		returnData = returnValues[0].Interface()
	}

	if err := returnValues[len(returnValues)-1].Interface(); err != nil {
		returnError = returnValues[len(returnValues)-1].Interface().(error)
	}

	return returnData, returnError
}
