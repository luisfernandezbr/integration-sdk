package gen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	assert := assert.New(t)
	var schema Schema
	schema.Name = "work"
	schema.Description = "the test model"
	schema.Version = 1
	schema.Models = map[string]Model{
		"project": Model{
			Columns: []Column{
				Column{
					Name: "foo_bar",
					Type: StringColumnType,
				},
				Column{
					Name: "b",
					Type: BooleanColumnType,
				},
				Column{
					Name:     "c",
					Type:     IntColumnType,
					Optional: true,
				},
			},
		},
	}
	var out strings.Builder
	assert.NoError(Generate(schema, "project", &out))
}
