// DO NOT EDIT -- generated code

// Package work - the system which contains project work
package work

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

// GenerateKQL will generate SQL statements for all models into file in directory
func GenerateKQL(dir string) error {
	var builder strings.Builder
	builder.WriteString(CreateChangelogKQLStreamSQL())
	builder.WriteString("\n")
	builder.WriteString(CreateChangelogKQLTableSQL())
	builder.WriteString("\n")
	builder.WriteString(CreateFieldTypeKQLStreamSQL())
	builder.WriteString("\n")
	builder.WriteString(CreateFieldTypeKQLTableSQL())
	builder.WriteString("\n")
	builder.WriteString(CreateIssueKQLStreamSQL())
	builder.WriteString("\n")
	builder.WriteString(CreateIssueKQLTableSQL())
	builder.WriteString("\n")
	builder.WriteString(CreateProjectKQLStreamSQL())
	builder.WriteString("\n")
	builder.WriteString(CreateProjectKQLTableSQL())
	builder.WriteString("\n")
	builder.WriteString(CreateUserKQLStreamSQL())
	builder.WriteString("\n")
	builder.WriteString(CreateUserKQLTableSQL())
	builder.WriteString("\n")
	return ioutil.WriteFile(filepath.Join(dir, "work.v1.ksql"), []byte(builder.String()), 0644)
}

// GenerateAvroSchemaSpec will generate the Avro schema to directory for each model
func GenerateAvroSchemaSpec(dir string) error {
	if err := ioutil.WriteFile(filepath.Join(dir, "work.v1.Changelog.avro"), []byte(CreateChangelogAvroSchemaSpec()), 0644); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(dir, "work.v1.FieldType.avro"), []byte(CreateFieldTypeAvroSchemaSpec()), 0644); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(dir, "work.v1.Issue.avro"), []byte(CreateIssueAvroSchemaSpec()), 0644); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(dir, "work.v1.Project.avro"), []byte(CreateProjectAvroSchemaSpec()), 0644); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(dir, "work.v1.User.avro"), []byte(CreateUserAvroSchemaSpec()), 0644); err != nil {
		return err
	}
	return nil
}

func init() {
}
