package resolvers

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type payload struct {
	Message json.RawMessage
}

func (p payload) parse(argsType reflect.Type) (reflect.Value, error) {
	args := reflect.New(argsType)

	if err := json.Unmarshal(p.Message, args.Interface()); err != nil {
		return args.Elem(), fmt.Errorf("Unable to prepare payload: %s", err.Error())
	}

	return args.Elem(), nil
}
