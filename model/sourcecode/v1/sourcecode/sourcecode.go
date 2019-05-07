// DO NOT EDIT -- generated code

// Package sourcecode - the system which contains source code
package sourcecode

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
	builder.WriteString(CreateCommitKQLStreamSQL())
	builder.WriteString("\n")
	builder.WriteString(CreateCommitKQLTableSQL())
	builder.WriteString("\n")
	builder.WriteString(CreateCommitFileKQLStreamSQL())
	builder.WriteString("\n")
	builder.WriteString(CreateCommitFileKQLTableSQL())
	builder.WriteString("\n")
	builder.WriteString(CreatePullRequestKQLStreamSQL())
	builder.WriteString("\n")
	builder.WriteString(CreatePullRequestKQLTableSQL())
	builder.WriteString("\n")
	builder.WriteString(CreateRepoKQLStreamSQL())
	builder.WriteString("\n")
	builder.WriteString(CreateRepoKQLTableSQL())
	builder.WriteString("\n")
	return ioutil.WriteFile(filepath.Join(dir, "sourcecode.v1.ksql"), []byte(builder.String()), 0644)
}

// GenerateAvroSchemaSpec will generate the Avro schema to directory for each model
func GenerateAvroSchemaSpec(dir string) error {
	if err := ioutil.WriteFile(filepath.Join(dir, "sourcecode.v1.Changelog.avro"), []byte(CreateChangelogAvroSchemaSpec()), 0644); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(dir, "sourcecode.v1.Commit.avro"), []byte(CreateCommitAvroSchemaSpec()), 0644); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(dir, "sourcecode.v1.CommitFile.avro"), []byte(CreateCommitFileAvroSchemaSpec()), 0644); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(dir, "sourcecode.v1.PullRequest.avro"), []byte(CreatePullRequestAvroSchemaSpec()), 0644); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(dir, "sourcecode.v1.Repo.avro"), []byte(CreateRepoAvroSchemaSpec()), 0644); err != nil {
		return err
	}
	return nil
}

func init() {
}
