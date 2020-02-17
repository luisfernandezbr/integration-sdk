// DO NOT EDIT - this code is generated

package datamodel

import (
	"github.com/pinpt/go-common/datamodel"

	dm_agent "github.com/pinpt/integration-sdk/agent"
	dm_cicd "github.com/pinpt/integration-sdk/cicd"
	dm_codequality "github.com/pinpt/integration-sdk/codequality"
	dm_customer "github.com/pinpt/integration-sdk/customer"
	dm_sourcecode "github.com/pinpt/integration-sdk/sourcecode"
	dm_work "github.com/pinpt/integration-sdk/work"
)

// New returns a new instanceof from a ModelNameType
func New(name datamodel.ModelNameType) datamodel.Model {
	switch name {
	case "agent.CancelRequest":
		o := new(dm_agent.CancelRequest)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.CancelRequestTrigger":
		o := new(dm_agent.CancelRequestTrigger)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.CancelResponse":
		o := new(dm_agent.CancelResponse)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.CodequalityRequest":
		o := new(dm_agent.CodequalityRequest)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.CodequalityResponse":
		o := new(dm_agent.CodequalityResponse)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.CodequalityTrigger":
		o := new(dm_agent.CodequalityTrigger)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.Crash":
		o := new(dm_agent.Crash)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.Enabled":
		o := new(dm_agent.Enabled)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.EnrollRequest":
		o := new(dm_agent.EnrollRequest)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.EnrollResponse":
		o := new(dm_agent.EnrollResponse)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.Event":
		o := new(dm_agent.Event)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.ExportRequest":
		o := new(dm_agent.ExportRequest)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.ExportResponse":
		o := new(dm_agent.ExportResponse)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.ExportTrigger":
		o := new(dm_agent.ExportTrigger)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.IntegrationMutationRequest":
		o := new(dm_agent.IntegrationMutationRequest)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.IntegrationMutationResponse":
		o := new(dm_agent.IntegrationMutationResponse)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.IntegrationRequest":
		o := new(dm_agent.IntegrationRequest)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.IntegrationResponse":
		o := new(dm_agent.IntegrationResponse)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.Log":
		o := new(dm_agent.Log)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.Pause":
		o := new(dm_agent.Pause)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.Ping":
		o := new(dm_agent.Ping)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.ProjectRequest":
		o := new(dm_agent.ProjectRequest)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.ProjectResponse":
		o := new(dm_agent.ProjectResponse)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.ProjectTrigger":
		o := new(dm_agent.ProjectTrigger)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.RepoRequest":
		o := new(dm_agent.RepoRequest)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.RepoResponse":
		o := new(dm_agent.RepoResponse)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.RepoTrigger":
		o := new(dm_agent.RepoTrigger)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.Resume":
		o := new(dm_agent.Resume)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.Start":
		o := new(dm_agent.Start)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.Stop":
		o := new(dm_agent.Stop)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.Uninstall":
		o := new(dm_agent.Uninstall)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.UninstallRequest":
		o := new(dm_agent.UninstallRequest)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.UninstallResponse":
		o := new(dm_agent.UninstallResponse)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.UpdateRequest":
		o := new(dm_agent.UpdateRequest)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.UpdateResponse":
		o := new(dm_agent.UpdateResponse)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.UpdateTrigger":
		o := new(dm_agent.UpdateTrigger)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.WebappIntegrationMutationRequest":
		o := new(dm_agent.WebappIntegrationMutationRequest)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.WebappIntegrationMutationResponse":
		o := new(dm_agent.WebappIntegrationMutationResponse)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.WorkStatusRequest":
		o := new(dm_agent.WorkStatusRequest)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.WorkStatusResponse":
		o := new(dm_agent.WorkStatusResponse)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.WorkStatusTrigger":
		o := new(dm_agent.WorkStatusTrigger)
		o.FromMap(map[string]interface{}{})
		return o
	case "cicd.Build":
		o := new(dm_cicd.Build)
		o.FromMap(map[string]interface{}{})
		return o
	case "cicd.Deployment":
		o := new(dm_cicd.Deployment)
		o.FromMap(map[string]interface{}{})
		return o
	case "codequality.Metric":
		o := new(dm_codequality.Metric)
		o.FromMap(map[string]interface{}{})
		return o
	case "codequality.Project":
		o := new(dm_codequality.Project)
		o.FromMap(map[string]interface{}{})
		return o
	case "customer.CostCenter":
		o := new(dm_customer.CostCenter)
		o.FromMap(map[string]interface{}{})
		return o
	case "customer.Team":
		o := new(dm_customer.Team)
		o.FromMap(map[string]interface{}{})
		return o
	case "sourcecode.Blame":
		o := new(dm_sourcecode.Blame)
		o.FromMap(map[string]interface{}{})
		return o
	case "sourcecode.Branch":
		o := new(dm_sourcecode.Branch)
		o.FromMap(map[string]interface{}{})
		return o
	case "sourcecode.Commit":
		o := new(dm_sourcecode.Commit)
		o.FromMap(map[string]interface{}{})
		return o
	case "sourcecode.PullRequest":
		o := new(dm_sourcecode.PullRequest)
		o.FromMap(map[string]interface{}{})
		return o
	case "sourcecode.PullRequestBranch":
		o := new(dm_sourcecode.PullRequestBranch)
		o.FromMap(map[string]interface{}{})
		return o
	case "sourcecode.PullRequestComment":
		o := new(dm_sourcecode.PullRequestComment)
		o.FromMap(map[string]interface{}{})
		return o
	case "sourcecode.PullRequestCommit":
		o := new(dm_sourcecode.PullRequestCommit)
		o.FromMap(map[string]interface{}{})
		return o
	case "sourcecode.PullRequestReview":
		o := new(dm_sourcecode.PullRequestReview)
		o.FromMap(map[string]interface{}{})
		return o
	case "sourcecode.Repo":
		o := new(dm_sourcecode.Repo)
		o.FromMap(map[string]interface{}{})
		return o
	case "sourcecode.User":
		o := new(dm_sourcecode.User)
		o.FromMap(map[string]interface{}{})
		return o
	case "work.Issue":
		o := new(dm_work.Issue)
		o.FromMap(map[string]interface{}{})
		return o
	case "work.Project":
		o := new(dm_work.Project)
		o.FromMap(map[string]interface{}{})
		return o
	case "work.Sprint":
		o := new(dm_work.Sprint)
		o.FromMap(map[string]interface{}{})
		return o
	case "work.User":
		o := new(dm_work.User)
		o.FromMap(map[string]interface{}{})
		return o
	}
	return nil
}

// NewFromTopic returns a new instanceof from a TopicNameType
func NewFromTopic(name datamodel.TopicNameType) datamodel.Model {
	switch name {
	case "sourcecode_repo_topic":
		o := new(dm_sourcecode.Repo)
		o.FromMap(map[string]interface{}{})
		return o
	case "sourcecode_user_topic":
		o := new(dm_sourcecode.User)
		o.FromMap(map[string]interface{}{})
		return o
	case "work_project_topic":
		o := new(dm_work.Project)
		o.FromMap(map[string]interface{}{})
		return o
	case "work_sprint_topic":
		o := new(dm_work.Sprint)
		o.FromMap(map[string]interface{}{})
		return o
	case "work_user_topic":
		o := new(dm_work.User)
		o.FromMap(map[string]interface{}{})
		return o
	}
	return nil
}

// GetMaterializedTopics returns an array of topics to be materialized
func GetMaterializedTopics() []datamodel.TopicNameType {
	return []datamodel.TopicNameType{}
}

// GetTopics returns an array of topics that are configured to be evented
func GetTopics() []datamodel.TopicNameType {
	return []datamodel.TopicNameType{
		datamodel.TopicNameType("sourcecode_repo_topic"),
		datamodel.TopicNameType("sourcecode_user_topic"),
		datamodel.TopicNameType("work_project_topic"),
		datamodel.TopicNameType("work_sprint_topic"),
		datamodel.TopicNameType("work_user_topic"),
	}
}

// GetModelNames returns an array of model names
func GetModelNames() []datamodel.ModelNameType {
	return []datamodel.ModelNameType{
		datamodel.ModelNameType("agent.CancelRequest"),
		datamodel.ModelNameType("agent.CancelRequestTrigger"),
		datamodel.ModelNameType("agent.CancelResponse"),
		datamodel.ModelNameType("agent.CodequalityRequest"),
		datamodel.ModelNameType("agent.CodequalityResponse"),
		datamodel.ModelNameType("agent.CodequalityTrigger"),
		datamodel.ModelNameType("agent.Crash"),
		datamodel.ModelNameType("agent.Enabled"),
		datamodel.ModelNameType("agent.EnrollRequest"),
		datamodel.ModelNameType("agent.EnrollResponse"),
		datamodel.ModelNameType("agent.Event"),
		datamodel.ModelNameType("agent.ExportRequest"),
		datamodel.ModelNameType("agent.ExportResponse"),
		datamodel.ModelNameType("agent.ExportTrigger"),
		datamodel.ModelNameType("agent.IntegrationMutationRequest"),
		datamodel.ModelNameType("agent.IntegrationMutationResponse"),
		datamodel.ModelNameType("agent.IntegrationRequest"),
		datamodel.ModelNameType("agent.IntegrationResponse"),
		datamodel.ModelNameType("agent.Log"),
		datamodel.ModelNameType("agent.Pause"),
		datamodel.ModelNameType("agent.Ping"),
		datamodel.ModelNameType("agent.ProjectRequest"),
		datamodel.ModelNameType("agent.ProjectResponse"),
		datamodel.ModelNameType("agent.ProjectTrigger"),
		datamodel.ModelNameType("agent.RepoRequest"),
		datamodel.ModelNameType("agent.RepoResponse"),
		datamodel.ModelNameType("agent.RepoTrigger"),
		datamodel.ModelNameType("agent.Resume"),
		datamodel.ModelNameType("agent.Start"),
		datamodel.ModelNameType("agent.Stop"),
		datamodel.ModelNameType("agent.Uninstall"),
		datamodel.ModelNameType("agent.UninstallRequest"),
		datamodel.ModelNameType("agent.UninstallResponse"),
		datamodel.ModelNameType("agent.UpdateRequest"),
		datamodel.ModelNameType("agent.UpdateResponse"),
		datamodel.ModelNameType("agent.UpdateTrigger"),
		datamodel.ModelNameType("agent.WebappIntegrationMutationRequest"),
		datamodel.ModelNameType("agent.WebappIntegrationMutationResponse"),
		datamodel.ModelNameType("agent.WorkStatusRequest"),
		datamodel.ModelNameType("agent.WorkStatusResponse"),
		datamodel.ModelNameType("agent.WorkStatusTrigger"),
		datamodel.ModelNameType("cicd.Build"),
		datamodel.ModelNameType("cicd.Deployment"),
		datamodel.ModelNameType("codequality.Metric"),
		datamodel.ModelNameType("codequality.Project"),
		datamodel.ModelNameType("customer.CostCenter"),
		datamodel.ModelNameType("customer.Team"),
		datamodel.ModelNameType("sourcecode.Blame"),
		datamodel.ModelNameType("sourcecode.Branch"),
		datamodel.ModelNameType("sourcecode.Commit"),
		datamodel.ModelNameType("sourcecode.PullRequest"),
		datamodel.ModelNameType("sourcecode.PullRequestBranch"),
		datamodel.ModelNameType("sourcecode.PullRequestComment"),
		datamodel.ModelNameType("sourcecode.PullRequestCommit"),
		datamodel.ModelNameType("sourcecode.PullRequestReview"),
		datamodel.ModelNameType("sourcecode.Repo"),
		datamodel.ModelNameType("sourcecode.User"),
		datamodel.ModelNameType("work.Issue"),
		datamodel.ModelNameType("work.Project"),
		datamodel.ModelNameType("work.Sprint"),
		datamodel.ModelNameType("work.User"),
	}
}
