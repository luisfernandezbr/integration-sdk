// DO NOT EDIT -- generated code

// Package work - the system which contains project work
package work

import "context"

type Db interface {
	Create(ctx context.Context, kv map[string]interface{}) error
	Update(ctx context.Context, id string, kv map[string]interface{}) error
	Delete(ctx context.Context, id string) error
	FindOne(ctx context.Context, id string) (map[string]interface{}, error)
	Find(ctx context.Context, kv map[string]interface{}) ([]map[string]interface{}, error)
}

func init() {
}
