// DO NOT EDIT - this code is generated

package datamodel

import (
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-datamodel/admin"
	"github.com/pinpt/go-datamodel/agent"
	"github.com/pinpt/go-datamodel/auth"
	"github.com/pinpt/go-datamodel/codequality"
	"github.com/pinpt/go-datamodel/customer"
	"github.com/pinpt/go-datamodel/pipeline"
	"github.com/pinpt/go-datamodel/sourcecode"
	"github.com/pinpt/go-datamodel/web"
	"github.com/pinpt/go-datamodel/work"
)

// New returns a new instanceof from a ModelNameType
func New(name datamodel.ModelNameType) datamodel.Model {
	switch name {
	case "sourcecode.CommitFile":
		return new(sourcecode.CommitFile)
	case "sourcecode.PullRequest":
		return new(sourcecode.PullRequest)
	case "sourcecode.Repo":
		return new(sourcecode.Repo)
	case "sourcecode.User":
		return new(sourcecode.User)
	case "sourcecode.Blame":
		return new(sourcecode.Blame)
	case "sourcecode.Branch":
		return new(sourcecode.Branch)
	case "sourcecode.Commit":
		return new(sourcecode.Commit)
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
	case "work.User":
		return new(work.User)
	case "customer.Team":
		return new(customer.Team)
	case "customer.User":
		return new(customer.User)
	case "customer.CostCenter":
		return new(customer.CostCenter)
	case "codequality.Metric":
		return new(codequality.Metric)
	case "codequality.Project":
		return new(codequality.Project)
	case "auth.ACLGrant":
		return new(auth.ACLGrant)
	case "admin.Agent":
		return new(admin.Agent)
	case "agent.Event":
		return new(agent.Event)
	case "pipeline.User":
		return new(pipeline.User)
	case "web.Hook":
		return new(web.Hook)
	}
	return nil
}

// NewFromTopic returns a new instanceof from a TopicNameType
func NewFromTopic(name datamodel.TopicNameType) datamodel.Model {
	switch name {
	case "auth_ACLGrant_topic":
		return new(auth.ACLGrant)
	case "sourcecode_User_topic":
		return new(sourcecode.User)
	case "sourcecode_Blame_topic":
		return new(sourcecode.Blame)
	case "sourcecode_Branch_topic":
		return new(sourcecode.Branch)
	case "sourcecode_Commit_topic":
		return new(sourcecode.Commit)
	case "sourcecode_CommitFile_topic":
		return new(sourcecode.CommitFile)
	case "sourcecode_PullRequest_topic":
		return new(sourcecode.PullRequest)
	case "sourcecode_Repo_topic":
		return new(sourcecode.Repo)
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
	case "work_User_topic":
		return new(work.User)
	case "customer_Team_topic":
		return new(customer.Team)
	case "customer_User_topic":
		return new(customer.User)
	case "customer_CostCenter_topic":
		return new(customer.CostCenter)
	case "codequality_Metric_topic":
		return new(codequality.Metric)
	case "codequality_Project_topic":
		return new(codequality.Project)
	case "web_Hook_topic":
		return new(web.Hook)
	case "admin_Agent_topic":
		return new(admin.Agent)
	case "agent_Event_topic":
		return new(agent.Event)
	case "pipeline_User_topic":
		return new(pipeline.User)
	}
	return nil
}

// GetMaterializedTopics returns an array of topics to be materialized
func GetMaterializedTopics() []datamodel.TopicNameType {
	return []datamodel.TopicNameType{
		datamodel.TopicNameType("customer_Team_topic"),
		datamodel.TopicNameType("customer_User_topic"),
		datamodel.TopicNameType("customer_CostCenter_topic"),
		datamodel.TopicNameType("auth_ACLGrant_topic"),
		datamodel.TopicNameType("sourcecode_Branch_topic"),
		datamodel.TopicNameType("sourcecode_CommitFile_topic"),
		datamodel.TopicNameType("sourcecode_PullRequest_topic"),
		datamodel.TopicNameType("sourcecode_Repo_topic"),
		datamodel.TopicNameType("sourcecode_User_topic"),
		datamodel.TopicNameType("admin_Agent_topic"),
	}
}

// GetTopics returns an array of topics that are configured
func GetTopics() []datamodel.TopicNameType {
	return []datamodel.TopicNameType{
		datamodel.TopicNameType("customer_Team_topic"),
		datamodel.TopicNameType("customer_User_topic"),
		datamodel.TopicNameType("codequality_Metric_topic"),
		datamodel.TopicNameType("customer_CostCenter_topic"),
		datamodel.TopicNameType("codequality_Project_topic"),
		datamodel.TopicNameType("agent_Event_topic"),
		datamodel.TopicNameType("auth_ACLGrant_topic"),
		datamodel.TopicNameType("pipeline_User_topic"),
		datamodel.TopicNameType("sourcecode_Branch_topic"),
		datamodel.TopicNameType("sourcecode_Commit_topic"),
		datamodel.TopicNameType("sourcecode_CommitFile_topic"),
		datamodel.TopicNameType("sourcecode_PullRequest_topic"),
		datamodel.TopicNameType("sourcecode_Repo_topic"),
		datamodel.TopicNameType("sourcecode_User_topic"),
		datamodel.TopicNameType("sourcecode_Blame_topic"),
		datamodel.TopicNameType("web_Hook_topic"),
		datamodel.TopicNameType("work_Changelog_topic"),
		datamodel.TopicNameType("work_CustomField_topic"),
		datamodel.TopicNameType("work_Issue_topic"),
		datamodel.TopicNameType("work_Project_topic"),
		datamodel.TopicNameType("work_Sprint_topic"),
		datamodel.TopicNameType("work_User_topic"),
		datamodel.TopicNameType("admin_Agent_topic"),
	}
}

// GetModelNames returns an array of model names that are configured
func GetModelNames() []datamodel.ModelNameType {
	return []datamodel.ModelNameType{
		datamodel.ModelNameType("customer.Team"),
		datamodel.ModelNameType("customer.User"),
		datamodel.ModelNameType("codequality.Metric"),
		datamodel.ModelNameType("customer.CostCenter"),
		datamodel.ModelNameType("codequality.Project"),
		datamodel.ModelNameType("agent.Event"),
		datamodel.ModelNameType("auth.ACLGrant"),
		datamodel.ModelNameType("pipeline.User"),
		datamodel.ModelNameType("sourcecode.Branch"),
		datamodel.ModelNameType("sourcecode.Commit"),
		datamodel.ModelNameType("sourcecode.CommitFile"),
		datamodel.ModelNameType("sourcecode.PullRequest"),
		datamodel.ModelNameType("sourcecode.Repo"),
		datamodel.ModelNameType("sourcecode.User"),
		datamodel.ModelNameType("sourcecode.Blame"),
		datamodel.ModelNameType("web.Hook"),
		datamodel.ModelNameType("work.Changelog"),
		datamodel.ModelNameType("work.CustomField"),
		datamodel.ModelNameType("work.Issue"),
		datamodel.ModelNameType("work.Project"),
		datamodel.ModelNameType("work.Sprint"),
		datamodel.ModelNameType("work.User"),
		datamodel.ModelNameType("admin.Agent"),
	}
}
