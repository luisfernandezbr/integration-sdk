// DO NOT EDIT - this code is generated

package datamodel

import (
	"github.com/pinpt/go-common/datamodel"

	dm_admin "github.com/pinpt/go-datamodel/admin"
	dm_agent "github.com/pinpt/go-datamodel/agent"
	dm_auth "github.com/pinpt/go-datamodel/auth"
	dm_codequality "github.com/pinpt/go-datamodel/codequality"
	dm_customer "github.com/pinpt/go-datamodel/customer"
	dm_ops_db "github.com/pinpt/go-datamodel/ops/db"
	dm_pipeline_activity "github.com/pinpt/go-datamodel/pipeline/activity"
	dm_pipeline_customer "github.com/pinpt/go-datamodel/pipeline/customer"
	dm_pipeline_integration "github.com/pinpt/go-datamodel/pipeline/integration"
	dm_pipeline_signal "github.com/pinpt/go-datamodel/pipeline/signal"
	dm_pipeline_sourcecode "github.com/pinpt/go-datamodel/pipeline/sourcecode"
	dm_pipeline_work "github.com/pinpt/go-datamodel/pipeline/work"
	dm_sourcecode "github.com/pinpt/go-datamodel/sourcecode"
	dm_web "github.com/pinpt/go-datamodel/web"
	dm_work "github.com/pinpt/go-datamodel/work"
)

// New returns a new instanceof from a ModelNameType
func New(name datamodel.ModelNameType) datamodel.Model {
	switch name {
	case "admin.Agent":
		o := new(dm_admin.Agent)
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
	case "agent.ProjectRequest":
		o := new(dm_agent.ProjectRequest)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.ProjectResponse":
		o := new(dm_agent.ProjectResponse)
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
	case "agent.UserRequest":
		o := new(dm_agent.UserRequest)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.UserResponse":
		o := new(dm_agent.UserResponse)
		o.FromMap(map[string]interface{}{})
		return o
	case "auth.ACLGrant":
		o := new(dm_auth.ACLGrant)
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
	case "ops.db.Change":
		o := new(dm_ops_db.Change)
		o.FromMap(map[string]interface{}{})
		return o
	case "pipeline.activity.Activity":
		o := new(dm_pipeline_activity.Activity)
		o.FromMap(map[string]interface{}{})
		return o
	case "pipeline.customer.User":
		o := new(dm_pipeline_customer.User)
		o.FromMap(map[string]interface{}{})
		return o
	case "pipeline.integration.User":
		o := new(dm_pipeline_integration.User)
		o.FromMap(map[string]interface{}{})
		return o
	case "pipeline.signal.Signal":
		o := new(dm_pipeline_signal.Signal)
		o.FromMap(map[string]interface{}{})
		return o
	case "pipeline.sourcecode.Commit":
		o := new(dm_pipeline_sourcecode.Commit)
		o.FromMap(map[string]interface{}{})
		return o
	case "pipeline.work.CustomField":
		o := new(dm_pipeline_work.CustomField)
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
	case "web.Hook":
		o := new(dm_web.Hook)
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
	case "admin_Agent_topic":
		o := new(dm_admin.Agent)
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
	case "agent_ProjectRequest_topic":
		o := new(dm_agent.ProjectRequest)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent_ProjectResponse_topic":
		o := new(dm_agent.ProjectResponse)
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
	case "agent_UserRequest_topic":
		o := new(dm_agent.UserRequest)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent_UserResponse_topic":
		o := new(dm_agent.UserResponse)
		o.FromMap(map[string]interface{}{})
		return o
	case "auth_ACLGrant_topic":
		o := new(dm_auth.ACLGrant)
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
	case "ops_db_Change_topic":
		o := new(dm_ops_db.Change)
		o.FromMap(map[string]interface{}{})
		return o
	case "pipeline_activity_Activity_topic":
		o := new(dm_pipeline_activity.Activity)
		o.FromMap(map[string]interface{}{})
		return o
	case "pipeline_customer_User_topic":
		o := new(dm_pipeline_customer.User)
		o.FromMap(map[string]interface{}{})
		return o
	case "pipeline_integration_User_topic":
		o := new(dm_pipeline_integration.User)
		o.FromMap(map[string]interface{}{})
		return o
	case "pipeline_signal_Signal_topic":
		o := new(dm_pipeline_signal.Signal)
		o.FromMap(map[string]interface{}{})
		return o
	case "pipeline_sourcecode_Commit_topic":
		o := new(dm_pipeline_sourcecode.Commit)
		o.FromMap(map[string]interface{}{})
		return o
	case "pipeline_work_CustomField_topic":
		o := new(dm_pipeline_work.CustomField)
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
	case "web_Hook_topic":
		o := new(dm_web.Hook)
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
		datamodel.TopicNameType("admin_Agent_topic"),
		datamodel.TopicNameType("auth_ACLGrant_topic"),
		datamodel.TopicNameType("customer_CostCenter_topic"),
		datamodel.TopicNameType("customer_Team_topic"),
		datamodel.TopicNameType("customer_User_topic"),
		datamodel.TopicNameType("pipeline_activity_Activity_topic"),
		datamodel.TopicNameType("pipeline_integration_User_topic"),
		datamodel.TopicNameType("pipeline_signal_Signal_topic"),
		datamodel.TopicNameType("pipeline_sourcecode_Commit_topic"),
		datamodel.TopicNameType("pipeline_work_CustomField_topic"),
		datamodel.TopicNameType("sourcecode_Branch_topic"),
		datamodel.TopicNameType("sourcecode_PullRequestComment_topic"),
		datamodel.TopicNameType("sourcecode_PullRequestReview_topic"),
		datamodel.TopicNameType("sourcecode_PullRequest_topic"),
		datamodel.TopicNameType("sourcecode_Repo_topic"),
		datamodel.TopicNameType("sourcecode_User_topic"),
		datamodel.TopicNameType("work_Project_topic"),
		datamodel.TopicNameType("work_User_topic"),
	}
}

// GetTopics returns an array of topics that are configured
func GetTopics() []datamodel.TopicNameType {
	return []datamodel.TopicNameType{
		datamodel.TopicNameType("admin_Agent_topic"),
		datamodel.TopicNameType("agent_Enabled_topic"),
		datamodel.TopicNameType("agent_EnrollRequest_topic"),
		datamodel.TopicNameType("agent_EnrollResponse_topic"),
		datamodel.TopicNameType("agent_Event_topic"),
		datamodel.TopicNameType("agent_ExportRequest_topic"),
		datamodel.TopicNameType("agent_ExportResponse_topic"),
		datamodel.TopicNameType("agent_ExportTrigger_topic"),
		datamodel.TopicNameType("agent_IntegrationRequest_topic"),
		datamodel.TopicNameType("agent_IntegrationResponse_topic"),
		datamodel.TopicNameType("agent_ProjectRequest_topic"),
		datamodel.TopicNameType("agent_ProjectResponse_topic"),
		datamodel.TopicNameType("agent_RepoRequest_topic"),
		datamodel.TopicNameType("agent_RepoResponse_topic"),
		datamodel.TopicNameType("agent_UserRequest_topic"),
		datamodel.TopicNameType("agent_UserResponse_topic"),
		datamodel.TopicNameType("auth_ACLGrant_topic"),
		datamodel.TopicNameType("codequality_Metric_topic"),
		datamodel.TopicNameType("codequality_Project_topic"),
		datamodel.TopicNameType("customer_CostCenter_topic"),
		datamodel.TopicNameType("customer_Team_topic"),
		datamodel.TopicNameType("customer_User_topic"),
		datamodel.TopicNameType("ops_db_Change_topic"),
		datamodel.TopicNameType("pipeline_activity_Activity_topic"),
		datamodel.TopicNameType("pipeline_customer_User_topic"),
		datamodel.TopicNameType("pipeline_integration_User_topic"),
		datamodel.TopicNameType("pipeline_signal_Signal_topic"),
		datamodel.TopicNameType("pipeline_sourcecode_Commit_topic"),
		datamodel.TopicNameType("pipeline_work_CustomField_topic"),
		datamodel.TopicNameType("sourcecode_Blame_topic"),
		datamodel.TopicNameType("sourcecode_Branch_topic"),
		datamodel.TopicNameType("sourcecode_Commit_topic"),
		datamodel.TopicNameType("sourcecode_PullRequestComment_topic"),
		datamodel.TopicNameType("sourcecode_PullRequestReview_topic"),
		datamodel.TopicNameType("sourcecode_PullRequest_topic"),
		datamodel.TopicNameType("sourcecode_Repo_topic"),
		datamodel.TopicNameType("sourcecode_User_topic"),
		datamodel.TopicNameType("web_Hook_topic"),
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
		datamodel.ModelNameType("admin.Agent"),
		datamodel.ModelNameType("agent.Enabled"),
		datamodel.ModelNameType("agent.EnrollRequest"),
		datamodel.ModelNameType("agent.EnrollResponse"),
		datamodel.ModelNameType("agent.Event"),
		datamodel.ModelNameType("agent.ExportRequest"),
		datamodel.ModelNameType("agent.ExportResponse"),
		datamodel.ModelNameType("agent.ExportTrigger"),
		datamodel.ModelNameType("agent.IntegrationRequest"),
		datamodel.ModelNameType("agent.IntegrationResponse"),
		datamodel.ModelNameType("agent.ProjectRequest"),
		datamodel.ModelNameType("agent.ProjectResponse"),
		datamodel.ModelNameType("agent.RepoRequest"),
		datamodel.ModelNameType("agent.RepoResponse"),
		datamodel.ModelNameType("agent.UserRequest"),
		datamodel.ModelNameType("agent.UserResponse"),
		datamodel.ModelNameType("auth.ACLGrant"),
		datamodel.ModelNameType("codequality.Metric"),
		datamodel.ModelNameType("codequality.Project"),
		datamodel.ModelNameType("customer.CostCenter"),
		datamodel.ModelNameType("customer.Team"),
		datamodel.ModelNameType("customer.User"),
		datamodel.ModelNameType("ops.db.Change"),
		datamodel.ModelNameType("pipeline.activity.Activity"),
		datamodel.ModelNameType("pipeline.customer.User"),
		datamodel.ModelNameType("pipeline.integration.User"),
		datamodel.ModelNameType("pipeline.signal.Signal"),
		datamodel.ModelNameType("pipeline.sourcecode.Commit"),
		datamodel.ModelNameType("pipeline.work.CustomField"),
		datamodel.ModelNameType("sourcecode.Blame"),
		datamodel.ModelNameType("sourcecode.Branch"),
		datamodel.ModelNameType("sourcecode.Commit"),
		datamodel.ModelNameType("sourcecode.PullRequest"),
		datamodel.ModelNameType("sourcecode.PullRequestComment"),
		datamodel.ModelNameType("sourcecode.PullRequestReview"),
		datamodel.ModelNameType("sourcecode.Repo"),
		datamodel.ModelNameType("sourcecode.User"),
		datamodel.ModelNameType("web.Hook"),
		datamodel.ModelNameType("work.Changelog"),
		datamodel.ModelNameType("work.CustomField"),
		datamodel.ModelNameType("work.Issue"),
		datamodel.ModelNameType("work.Project"),
		datamodel.ModelNameType("work.Sprint"),
		datamodel.ModelNameType("work.User"),
	}
}
