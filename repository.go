package resolvers

import (
	"context"
	"fmt"
	"reflect"
)

// Repository stores all resolvers
type Repository map[string]resolver

// Add stores a new resolver
func (r Repository) Add(resolve string, handler interface{}) error {
	err := validators.run(reflect.TypeOf(handler))

	if err == nil {
		r[resolve] = resolver{
			function: handler,
		}
	}

	return err
}

// Handle responds to the AppSync request
func (r Repository) Handle(ctx context.Context, in invocation) (interface{}, error) {
	handler, found := r[in.Resolve]

	if found {
		return handler.call(ctx, in.payload())
	}

	return nil, fmt.Errorf("No resolver found: %s", in.Resolve)
}
