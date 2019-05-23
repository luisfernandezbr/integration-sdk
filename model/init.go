// DO NOT EDIT - this code is generated

package datamodel

import (
	"strings"

	"github.com/pinpt/go-datamodel/codequality"
	"github.com/pinpt/go-datamodel/customer"
	"github.com/pinpt/go-datamodel/sourcecode"
	"github.com/pinpt/go-datamodel/work"
)


func New(name string) interface{} {
	switch name {
		case "codequality.Metric":
			return new(codequality.Metric)
		case "codequality.Project":
			return new(codequality.Project)
		case "customer.Team":
			return new(customer.Team)
		case "customer.User":
			return new(customer.User)
		case "customer.CostCenter":
			return new(customer.CostCenter)
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
		case "work.User":
			return new(work.User)
		case "work.Changelog":
			return new(work.Changelog)
		case "work.CustomField":
			return new(work.CustomField)
		case "work.Issue":
			return new(work.Issue)
		case "work.Project":
			return new(work.Project)
		case "work.Sprint":
			return new(work.Sprint)
	}
	panic("invalid type specific: " + name)
}

