// DO NOT EDIT - this code is generated

package datamodel

import (
	"strings"

	"github.com/linkedin/goavro"
	"github.com/pinpt/go-datamodel/codequality"
	"github.com/pinpt/go-datamodel/customer"
	"github.com/pinpt/go-datamodel/sourcecode"
	"github.com/pinpt/go-datamodel/work"
)

// TopicNameType is a type for the name of a topic
type TopicNameType string

// ModelNameType is a type for the model name
type ModelNameType string

// Model is a generic model interface that all our models implement
type Model interface {
	// GetID returns the ID for the instance
	GetID() string
	// ToAvroBinary converts the instance to binary avro
	ToAvroBinary() ([]byte, *goavro.Codec, error)
	// Stringfy converts the instance to JSON string
	Stringify() string
	// ToMap converts the instance to a map
	ToMap(avro ...bool) map[string]interface{}
	// FromMap sets the properties of the instance from the map
	FromMap(kv map[string]interface{})
}

// New returns a new instanceof from a ModelNameType
func New(name ModelNameType) Model {
	switch name {
	case "codequality.Project":
		return new(codequality.Project)
	case "codequality.Metric":
		return new(codequality.Metric)
	case "customer.CostCenter":
		return new(customer.CostCenter)
	case "customer.Team":
		return new(customer.Team)
	case "customer.User":
		return new(customer.User)
	case "sourcecode.PullRequest":
		return new(sourcecode.PullRequest)
	case "sourcecode.Repo":
		return new(sourcecode.Repo)
	case "sourcecode.User":
		return new(sourcecode.User)
	case "sourcecode.Branch":
		return new(sourcecode.Branch)
	case "sourcecode.Changelog":
		return new(sourcecode.Changelog)
	case "sourcecode.Commit":
		return new(sourcecode.Commit)
	case "sourcecode.CommitFile":
		return new(sourcecode.CommitFile)
	case "work.CustomField":
		return new(work.CustomField)
	case "work.Issue":
		return new(work.Issue)
	case "work.Project":
		return new(work.Project)
	case "work.Sprint":
		return new(work.Sprint)
	case "work.User":
		return new(work.User)
	case "work.Changelog":
		return new(work.Changelog)
	}
	panic("invalid type specific: " + name)
}

// NewFromTopic returns a new instanceof from a TopicNameType
func NewFromTopic(name TopicNameType) interface{} {
	switch name {
	case "codequality_Project_topic":
		return new(codequality.Project)
	case "codequality_Metric_topic":
		return new(codequality.Metric)
	case "customer_CostCenter_topic":
		return new(customer.CostCenter)
	case "customer_Team_topic":
		return new(customer.Team)
	case "customer_User_topic":
		return new(customer.User)
	case "sourcecode_Repo_topic":
		return new(sourcecode.Repo)
	case "sourcecode_User_topic":
		return new(sourcecode.User)
	case "sourcecode_Branch_topic":
		return new(sourcecode.Branch)
	case "sourcecode_Changelog_topic":
		return new(sourcecode.Changelog)
	case "sourcecode_Commit_topic":
		return new(sourcecode.Commit)
	case "sourcecode_CommitFile_topic":
		return new(sourcecode.CommitFile)
	case "sourcecode_PullRequest_topic":
		return new(sourcecode.PullRequest)
	case "work_User_topic":
		return new(work.User)
	case "work_Changelog_topic":
		return new(work.Changelog)
	case "work_CustomField_topic":
		return new(work.CustomField)
	case "work_Issue_topic":
		return new(work.Issue)
	case "work_Project_topic":
		return new(work.Project)
	case "work_Sprint_topic":
		return new(work.Sprint)
	}
	panic("invalid type specific: " + name)
}
