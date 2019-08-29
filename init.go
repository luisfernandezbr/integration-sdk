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
	case "agent.Upgrade":
		o := new(dm_agent.Upgrade)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.UserRequest":
		o := new(dm_agent.UserRequest)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.UserResponse":
		o := new(dm_agent.UserResponse)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.UserTrigger":
		o := new(dm_agent.UserTrigger)
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
	case "customer.User":
		o := new(dm_customer.User)
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
	case "sourcecode.PullRequestComment":
		o := new(dm_sourcecode.PullRequestComment)
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
	case "work.Changelog":
		o := new(dm_work.Changelog)
		o.FromMap(map[string]interface{}{})
		return o
	case "work.CustomField":
		o := new(dm_work.CustomField)
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
	case "agent_CodequalityRequest_topic":
		o := new(dm_agent.CodequalityRequest)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent_CodequalityResponse_topic":
		o := new(dm_agent.CodequalityResponse)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent_CodequalityTrigger_topic":
		o := new(dm_agent.CodequalityTrigger)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent_Crash_topic":
		o := new(dm_agent.Crash)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent_Enabled_topic":
		o := new(dm_agent.Enabled)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent_EnrollRequest_topic":
		o := new(dm_agent.EnrollRequest)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent_EnrollResponse_topic":
		o := new(dm_agent.EnrollResponse)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent_Event_topic":
		o := new(dm_agent.Event)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent_ExportRequest_topic":
		o := new(dm_agent.ExportRequest)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent_ExportResponse_topic":
		o := new(dm_agent.ExportResponse)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent_ExportTrigger_topic":
		o := new(dm_agent.ExportTrigger)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent_IntegrationRequest_topic":
		o := new(dm_agent.IntegrationRequest)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent_IntegrationResponse_topic":
		o := new(dm_agent.IntegrationResponse)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent_Log_topic":
		o := new(dm_agent.Log)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent_Ping_topic":
		o := new(dm_agent.Ping)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent_ProjectRequest_topic":
		o := new(dm_agent.ProjectRequest)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent_ProjectResponse_topic":
		o := new(dm_agent.ProjectResponse)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent_ProjectTrigger_topic":
		o := new(dm_agent.ProjectTrigger)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent_RepoRequest_topic":
		o := new(dm_agent.RepoRequest)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent_RepoResponse_topic":
		o := new(dm_agent.RepoResponse)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent_RepoTrigger_topic":
		o := new(dm_agent.RepoTrigger)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent_Start_topic":
		o := new(dm_agent.Start)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent_Stop_topic":
		o := new(dm_agent.Stop)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent_Uninstall_topic":
		o := new(dm_agent.Uninstall)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent_Upgrade_topic":
		o := new(dm_agent.Upgrade)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent_UserRequest_topic":
		o := new(dm_agent.UserRequest)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent_UserResponse_topic":
		o := new(dm_agent.UserResponse)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent_UserTrigger_topic":
		o := new(dm_agent.UserTrigger)
		o.FromMap(map[string]interface{}{})
		return o
	case "cicd_Build_topic":
		o := new(dm_cicd.Build)
		o.FromMap(map[string]interface{}{})
		return o
	case "cicd_Deployment_topic":
		o := new(dm_cicd.Deployment)
		o.FromMap(map[string]interface{}{})
		return o
	case "codequality_Metric_topic":
		o := new(dm_codequality.Metric)
		o.FromMap(map[string]interface{}{})
		return o
	case "codequality_Project_topic":
		o := new(dm_codequality.Project)
		o.FromMap(map[string]interface{}{})
		return o
	case "customer_CostCenter_topic":
		o := new(dm_customer.CostCenter)
		o.FromMap(map[string]interface{}{})
		return o
	case "customer_Team_topic":
		o := new(dm_customer.Team)
		o.FromMap(map[string]interface{}{})
		return o
	case "customer_User_topic":
		o := new(dm_customer.User)
		o.FromMap(map[string]interface{}{})
		return o
	case "sourcecode_Blame_topic":
		o := new(dm_sourcecode.Blame)
		o.FromMap(map[string]interface{}{})
		return o
	case "sourcecode_Branch_topic":
		o := new(dm_sourcecode.Branch)
		o.FromMap(map[string]interface{}{})
		return o
	case "sourcecode_Commit_topic":
		o := new(dm_sourcecode.Commit)
		o.FromMap(map[string]interface{}{})
		return o
	case "sourcecode_PullRequest_topic":
		o := new(dm_sourcecode.PullRequest)
		o.FromMap(map[string]interface{}{})
		return o
	case "sourcecode_PullRequestComment_topic":
		o := new(dm_sourcecode.PullRequestComment)
		o.FromMap(map[string]interface{}{})
		return o
	case "sourcecode_PullRequestReview_topic":
		o := new(dm_sourcecode.PullRequestReview)
		o.FromMap(map[string]interface{}{})
		return o
	case "sourcecode_Repo_topic":
		o := new(dm_sourcecode.Repo)
		o.FromMap(map[string]interface{}{})
		return o
	case "sourcecode_User_topic":
		o := new(dm_sourcecode.User)
		o.FromMap(map[string]interface{}{})
		return o
	case "work_Changelog_topic":
		o := new(dm_work.Changelog)
		o.FromMap(map[string]interface{}{})
		return o
	case "work_CustomField_topic":
		o := new(dm_work.CustomField)
		o.FromMap(map[string]interface{}{})
		return o
	case "work_Issue_topic":
		o := new(dm_work.Issue)
		o.FromMap(map[string]interface{}{})
		return o
	case "work_Project_topic":
		o := new(dm_work.Project)
		o.FromMap(map[string]interface{}{})
		return o
	case "work_Sprint_topic":
		o := new(dm_work.Sprint)
		o.FromMap(map[string]interface{}{})
		return o
	case "work_User_topic":
		o := new(dm_work.User)
		o.FromMap(map[string]interface{}{})
		return o
	}
	return nil
}

// GetMaterializedTopics returns an array of topics to be materialized
func GetMaterializedTopics() []datamodel.TopicNameType {
	return []datamodel.TopicNameType{
		datamodel.TopicNameType("customer_CostCenter_topic"),
		datamodel.TopicNameType("customer_Team_topic"),
		datamodel.TopicNameType("customer_User_topic"),
	}
}

// GetTopics returns an array of topics that are configured
func GetTopics() []datamodel.TopicNameType {
	return []datamodel.TopicNameType{
		datamodel.TopicNameType("agent_CodequalityRequest_topic"),
		datamodel.TopicNameType("agent_CodequalityResponse_topic"),
		datamodel.TopicNameType("agent_CodequalityTrigger_topic"),
		datamodel.TopicNameType("agent_Crash_topic"),
		datamodel.TopicNameType("agent_Enabled_topic"),
		datamodel.TopicNameType("agent_EnrollRequest_topic"),
		datamodel.TopicNameType("agent_EnrollResponse_topic"),
		datamodel.TopicNameType("agent_Event_topic"),
		datamodel.TopicNameType("agent_ExportRequest_topic"),
		datamodel.TopicNameType("agent_ExportResponse_topic"),
		datamodel.TopicNameType("agent_ExportTrigger_topic"),
		datamodel.TopicNameType("agent_IntegrationRequest_topic"),
		datamodel.TopicNameType("agent_IntegrationResponse_topic"),
		datamodel.TopicNameType("agent_Log_topic"),
		datamodel.TopicNameType("agent_Ping_topic"),
		datamodel.TopicNameType("agent_ProjectRequest_topic"),
		datamodel.TopicNameType("agent_ProjectResponse_topic"),
		datamodel.TopicNameType("agent_ProjectTrigger_topic"),
		datamodel.TopicNameType("agent_RepoRequest_topic"),
		datamodel.TopicNameType("agent_RepoResponse_topic"),
		datamodel.TopicNameType("agent_RepoTrigger_topic"),
		datamodel.TopicNameType("agent_Start_topic"),
		datamodel.TopicNameType("agent_Stop_topic"),
		datamodel.TopicNameType("agent_Uninstall_topic"),
		datamodel.TopicNameType("agent_Upgrade_topic"),
		datamodel.TopicNameType("agent_UserRequest_topic"),
		datamodel.TopicNameType("agent_UserResponse_topic"),
		datamodel.TopicNameType("agent_UserTrigger_topic"),
		datamodel.TopicNameType("cicd_Build_topic"),
		datamodel.TopicNameType("cicd_Deployment_topic"),
		datamodel.TopicNameType("codequality_Metric_topic"),
		datamodel.TopicNameType("codequality_Project_topic"),
		datamodel.TopicNameType("customer_CostCenter_topic"),
		datamodel.TopicNameType("customer_Team_topic"),
		datamodel.TopicNameType("customer_User_topic"),
		datamodel.TopicNameType("sourcecode_Blame_topic"),
		datamodel.TopicNameType("sourcecode_Branch_topic"),
		datamodel.TopicNameType("sourcecode_Commit_topic"),
		datamodel.TopicNameType("sourcecode_Repo_topic"),
		datamodel.TopicNameType("sourcecode_User_topic"),
		datamodel.TopicNameType("work_Changelog_topic"),
		datamodel.TopicNameType("work_CustomField_topic"),
		datamodel.TopicNameType("work_Issue_topic"),
		datamodel.TopicNameType("work_Project_topic"),
		datamodel.TopicNameType("work_Sprint_topic"),
		datamodel.TopicNameType("work_User_topic"),
	}
}

// GetModelNames returns an array of model names that are configured
func GetModelNames() []datamodel.ModelNameType {
	return []datamodel.ModelNameType{
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
		datamodel.ModelNameType("agent.IntegrationRequest"),
		datamodel.ModelNameType("agent.IntegrationResponse"),
		datamodel.ModelNameType("agent.Log"),
		datamodel.ModelNameType("agent.Ping"),
		datamodel.ModelNameType("agent.ProjectRequest"),
		datamodel.ModelNameType("agent.ProjectResponse"),
		datamodel.ModelNameType("agent.ProjectTrigger"),
		datamodel.ModelNameType("agent.RepoRequest"),
		datamodel.ModelNameType("agent.RepoResponse"),
		datamodel.ModelNameType("agent.RepoTrigger"),
		datamodel.ModelNameType("agent.Start"),
		datamodel.ModelNameType("agent.Stop"),
		datamodel.ModelNameType("agent.Uninstall"),
		datamodel.ModelNameType("agent.Upgrade"),
		datamodel.ModelNameType("agent.UserRequest"),
		datamodel.ModelNameType("agent.UserResponse"),
		datamodel.ModelNameType("agent.UserTrigger"),
		datamodel.ModelNameType("cicd.Build"),
		datamodel.ModelNameType("cicd.Deployment"),
		datamodel.ModelNameType("codequality.Metric"),
		datamodel.ModelNameType("codequality.Project"),
		datamodel.ModelNameType("customer.CostCenter"),
		datamodel.ModelNameType("customer.Team"),
		datamodel.ModelNameType("customer.User"),
		datamodel.ModelNameType("sourcecode.Blame"),
		datamodel.ModelNameType("sourcecode.Branch"),
		datamodel.ModelNameType("sourcecode.Commit"),
		datamodel.ModelNameType("sourcecode.PullRequest"),
		datamodel.ModelNameType("sourcecode.PullRequestComment"),
		datamodel.ModelNameType("sourcecode.PullRequestReview"),
		datamodel.ModelNameType("sourcecode.Repo"),
		datamodel.ModelNameType("sourcecode.User"),
		datamodel.ModelNameType("work.Changelog"),
		datamodel.ModelNameType("work.CustomField"),
		datamodel.ModelNameType("work.Issue"),
		datamodel.ModelNameType("work.Project"),
		datamodel.ModelNameType("work.Sprint"),
		datamodel.ModelNameType("work.User"),
	}
}
