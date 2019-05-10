package gen

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/serenize/snaker"

	"github.com/pinpt/go-common/fileutil"

	yaml "gopkg.in/yaml.v2"
)

// ColumnType is the column type
type ColumnType string

const (
	// NullColumnType is the type for null values
	NullColumnType ColumnType = "null"
	// BooleanColumnType is the type for boolean values
	BooleanColumnType ColumnType = "boolean"
	// IntColumnType is the type for int values
	IntColumnType ColumnType = "int"
	// LongColumnType is the type for long values
	LongColumnType ColumnType = "long"
	// FloatColumnType is the type for float values
	FloatColumnType ColumnType = "float"
	// DoubleColumnType is the type for double values
	DoubleColumnType ColumnType = "double"
	// BytesColumnType is the type for bytes values
	BytesColumnType ColumnType = "bytes"
	// StringColumnType is the type for string values
	StringColumnType ColumnType = "string"
)

// Column in the schema
type Column struct {
	Name        string
	Description string
	Type        ColumnType
	Key         bool
	Optional    bool
	Timestamp   bool
	RefID       bool

	// internal
	IsArray bool
	IsMap   bool
}

// Model in the schema
type Model struct {
	Description string
	Columns     []*Column
	Extends     []string
}

// Schema is the schema for a particular domain
type Schema struct {
	Name        string
	Version     int
	Description string
	Models      map[string]Model
}

func isArray(val string) (bool, string) {
	if strings.HasPrefix(val, "Array<") {
		return true, val[6 : len(val)-1]
	}
	return false, val
}

func isMap(val string) (bool, string) {
	if strings.HasPrefix(val, "Map<") {
		return true, val[4 : len(val)-1]
	}
	return false, val
}

func camel(val string) string {
	return snaker.SnakeToCamel(strings.ToUpper(val[0:1]) + val[1:])
}

// Name of the model
func (m *Model) Name(s *Schema, name string) string {
	return fmt.Sprintf("%s.v%v.%s", s.Name, s.Version, name)
}

// Resolve any extends schema
func (s *Schema) Resolve(dirs []string) error {
	for mn, m := range s.Models {
		if m.Extends != nil && len(m.Extends) > 0 {
			for _, name := range m.Extends {
				var found bool
				// first we search our current model
				for cmname, cm := range s.Models {
					if name == cmname || name == cm.Name(s, camel(cmname)) {
						m.Columns = append(m.Columns, cm.Columns...)
						found = true
						break
					}
				}
				if found {
					break
				}
				for _, dir := range dirs {
					files, err := fileutil.FindFiles(dir, regexp.MustCompile("\\.y[a]?ml$"))
					if err != nil {
						return err
					}
					if len(files) > 0 {
						for _, fp := range files {
							rel, _ := filepath.Abs(fp)
							of, err := os.Open(rel)
							if err != nil {
								return fmt.Errorf("error opening: %v. %v", rel, err)
							}
							defer of.Close()
							ns := &Schema{}
							if err := ns.DecodeYAML(of); err != nil {
								return fmt.Errorf("error decoding: %v. %v", rel, err)
							}
							for cmname, cm := range ns.Models {
								if name == cmname || name == cm.Name(ns, camel(cmname)) {
									m.Columns = append(m.Columns, cm.Columns...)
									s.Models[mn] = m
									found = true
									break
								}
							}
							if found {
								break
							}
						}
						if found {
							break
						}
					}
				}
			}
		}
	}
	return nil
}

// DecodeYAML will decode the schema from reader as yaml
func (s *Schema) DecodeYAML(r io.Reader) error {
	if err := yaml.NewDecoder(r).Decode(s); err != nil {
		return err
	}
	for _, m := range s.Models {
		for _, col := range m.Columns {
			ok, t := isArray(string(col.Type))
			if ok {
				col.IsArray = true
				col.Type = ColumnType(t)
			}
			ok, t = isMap(string(col.Type))
			if ok {
				col.IsMap = true
				col.Type = ColumnType(t)
			}
		}
	}
	return nil
}

// EncodeYAML will encode the schema to the writer as yaml
func (s *Schema) EncodeYAML(w io.Writer) error {
	return yaml.NewEncoder(w).Encode(s)
}

// GetSchemaFilePath returns a full file path for a given schema
func (s *Schema) GetSchemaFilePath(basedir string, name string) string {
	return filepath.Join(basedir, fmt.Sprintf("%s/v%d/%s/%s.go", s.Name, s.Version, s.Name, name))
}
