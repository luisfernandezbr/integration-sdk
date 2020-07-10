// DO NOT EDIT - this code is generated

package datamodel

import (
	"github.com/pinpt/go-common/v10/datamodel"

	dm_agent "github.com/pinpt/integration-sdk/agent"
	dm_calendar "github.com/pinpt/integration-sdk/calendar"
	dm_cicd "github.com/pinpt/integration-sdk/cicd"
	dm_codequality "github.com/pinpt/integration-sdk/codequality"
	dm_customer "github.com/pinpt/integration-sdk/customer"
	dm_sourcecode "github.com/pinpt/integration-sdk/sourcecode"
	dm_web "github.com/pinpt/integration-sdk/web"
	dm_work "github.com/pinpt/integration-sdk/work"
	dm_work_agile "github.com/pinpt/integration-sdk/work/agile"
)

// New returns a new instanceof from a ModelNameType
func New(name datamodel.ModelNameType) datamodel.Model {
	switch name {
	case "agent.CalendarRequest":
		o := new(dm_agent.CalendarRequest)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.CalendarResponse":
		o := new(dm_agent.CalendarResponse)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.CalendarTrigger":
		o := new(dm_agent.CalendarTrigger)
		o.FromMap(map[string]interface{}{})
		return o
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
	case "agent.Enrollment":
		o := new(dm_agent.Enrollment)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.Event":
		o := new(dm_agent.Event)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.Export":
		o := new(dm_agent.Export)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.ExportData":
		o := new(dm_agent.ExportData)
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
	case "agent.IntegrationInstance":
		o := new(dm_agent.IntegrationInstance)
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
	case "agent.UpdateData":
		o := new(dm_agent.UpdateData)
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
	case "agent.WebhookRequest":
		o := new(dm_agent.WebhookRequest)
		o.FromMap(map[string]interface{}{})
		return o
	case "agent.WebhookResponse":
		o := new(dm_agent.WebhookResponse)
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
	case "calendar.Calendar":
		o := new(dm_calendar.Calendar)
		o.FromMap(map[string]interface{}{})
		return o
	case "calendar.Event":
		o := new(dm_calendar.Event)
		o.FromMap(map[string]interface{}{})
		return o
	case "calendar.User":
		o := new(dm_calendar.User)
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
	case "sourcecode.RepoError":
		o := new(dm_sourcecode.RepoError)
		o.FromMap(map[string]interface{}{})
		return o
	case "sourcecode.RepoWebhook":
		o := new(dm_sourcecode.RepoWebhook)
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
	case "work.Config":
		o := new(dm_work.Config)
		o.FromMap(map[string]interface{}{})
		return o
	case "work.Issue":
		o := new(dm_work.Issue)
		o.FromMap(map[string]interface{}{})
		return o
	case "work.IssueComment":
		o := new(dm_work.IssueComment)
		o.FromMap(map[string]interface{}{})
		return o
	case "work.IssuePriority":
		o := new(dm_work.IssuePriority)
		o.FromMap(map[string]interface{}{})
		return o
	case "work.IssueStatus":
		o := new(dm_work.IssueStatus)
		o.FromMap(map[string]interface{}{})
		return o
	case "work.IssueType":
		o := new(dm_work.IssueType)
		o.FromMap(map[string]interface{}{})
		return o
	case "work.KanbanBoard":
		o := new(dm_work.KanbanBoard)
		o.FromMap(map[string]interface{}{})
		return o
	case "work.Project":
		o := new(dm_work.Project)
		o.FromMap(map[string]interface{}{})
		return o
	case "work.ProjectCapability":
		o := new(dm_work.ProjectCapability)
		o.FromMap(map[string]interface{}{})
		return o
	case "work.ProjectError":
		o := new(dm_work.ProjectError)
		o.FromMap(map[string]interface{}{})
		return o
	case "work.ProjectWebhook":
		o := new(dm_work.ProjectWebhook)
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
	case "work.agile.Board":
		o := new(dm_work_agile.Board)
		o.FromMap(map[string]interface{}{})
		return o
	case "work.agile.Kanban":
		o := new(dm_work_agile.Kanban)
		o.FromMap(map[string]interface{}{})
		return o
	case "work.agile.Sprint":
		o := new(dm_work_agile.Sprint)
		o.FromMap(map[string]interface{}{})
		return o
	}
	return nil
}

// NewPartial returns a new instanceof from a string
func NewPartial(name datamodel.ModelNameType) datamodel.PartialModel {
	switch name {
	case "agent.CalendarRequest":
		return new(dm_agent.CalendarRequestPartial)
	case "agent.CalendarResponse":
		return new(dm_agent.CalendarResponsePartial)
	case "agent.CalendarTrigger":
		return new(dm_agent.CalendarTriggerPartial)
	case "agent.CancelRequest":
		return new(dm_agent.CancelRequestPartial)
	case "agent.CancelRequestTrigger":
		return new(dm_agent.CancelRequestTriggerPartial)
	case "agent.CancelResponse":
		return new(dm_agent.CancelResponsePartial)
	case "agent.CodequalityRequest":
		return new(dm_agent.CodequalityRequestPartial)
	case "agent.CodequalityResponse":
		return new(dm_agent.CodequalityResponsePartial)
	case "agent.CodequalityTrigger":
		return new(dm_agent.CodequalityTriggerPartial)
	case "agent.Crash":
		return new(dm_agent.CrashPartial)
	case "agent.Enabled":
		return new(dm_agent.EnabledPartial)
	case "agent.EnrollRequest":
		return new(dm_agent.EnrollRequestPartial)
	case "agent.EnrollResponse":
		return new(dm_agent.EnrollResponsePartial)
	case "agent.Enrollment":
		return new(dm_agent.EnrollmentPartial)
	case "agent.Event":
		return new(dm_agent.EventPartial)
	case "agent.Export":
		return new(dm_agent.ExportPartial)
	case "agent.ExportData":
		return new(dm_agent.ExportDataPartial)
	case "agent.ExportRequest":
		return new(dm_agent.ExportRequestPartial)
	case "agent.ExportResponse":
		return new(dm_agent.ExportResponsePartial)
	case "agent.ExportTrigger":
		return new(dm_agent.ExportTriggerPartial)
	case "agent.IntegrationInstance":
		return new(dm_agent.IntegrationInstancePartial)
	case "agent.IntegrationMutationRequest":
		return new(dm_agent.IntegrationMutationRequestPartial)
	case "agent.IntegrationMutationResponse":
		return new(dm_agent.IntegrationMutationResponsePartial)
	case "agent.IntegrationRequest":
		return new(dm_agent.IntegrationRequestPartial)
	case "agent.IntegrationResponse":
		return new(dm_agent.IntegrationResponsePartial)
	case "agent.Log":
		return new(dm_agent.LogPartial)
	case "agent.Pause":
		return new(dm_agent.PausePartial)
	case "agent.Ping":
		return new(dm_agent.PingPartial)
	case "agent.ProjectRequest":
		return new(dm_agent.ProjectRequestPartial)
	case "agent.ProjectResponse":
		return new(dm_agent.ProjectResponsePartial)
	case "agent.ProjectTrigger":
		return new(dm_agent.ProjectTriggerPartial)
	case "agent.RepoRequest":
		return new(dm_agent.RepoRequestPartial)
	case "agent.RepoResponse":
		return new(dm_agent.RepoResponsePartial)
	case "agent.RepoTrigger":
		return new(dm_agent.RepoTriggerPartial)
	case "agent.Resume":
		return new(dm_agent.ResumePartial)
	case "agent.Start":
		return new(dm_agent.StartPartial)
	case "agent.Stop":
		return new(dm_agent.StopPartial)
	case "agent.Uninstall":
		return new(dm_agent.UninstallPartial)
	case "agent.UninstallRequest":
		return new(dm_agent.UninstallRequestPartial)
	case "agent.UninstallResponse":
		return new(dm_agent.UninstallResponsePartial)
	case "agent.UpdateData":
		return new(dm_agent.UpdateDataPartial)
	case "agent.UpdateRequest":
		return new(dm_agent.UpdateRequestPartial)
	case "agent.UpdateResponse":
		return new(dm_agent.UpdateResponsePartial)
	case "agent.UpdateTrigger":
		return new(dm_agent.UpdateTriggerPartial)
	case "agent.WebappIntegrationMutationRequest":
		return new(dm_agent.WebappIntegrationMutationRequestPartial)
	case "agent.WebhookRequest":
		return new(dm_agent.WebhookRequestPartial)
	case "agent.WebhookResponse":
		return new(dm_agent.WebhookResponsePartial)
	case "agent.WorkStatusRequest":
		return new(dm_agent.WorkStatusRequestPartial)
	case "agent.WorkStatusResponse":
		return new(dm_agent.WorkStatusResponsePartial)
	case "agent.WorkStatusTrigger":
		return new(dm_agent.WorkStatusTriggerPartial)
	case "calendar.Calendar":
		return new(dm_calendar.CalendarPartial)
	case "calendar.Event":
		return new(dm_calendar.EventPartial)
	case "calendar.User":
		return new(dm_calendar.UserPartial)
	case "cicd.Build":
		return new(dm_cicd.BuildPartial)
	case "cicd.Deployment":
		return new(dm_cicd.DeploymentPartial)
	case "codequality.Metric":
		return new(dm_codequality.MetricPartial)
	case "codequality.Project":
		return new(dm_codequality.ProjectPartial)
	case "customer.CostCenter":
		return new(dm_customer.CostCenterPartial)
	case "customer.Team":
		return new(dm_customer.TeamPartial)
	case "sourcecode.Blame":
		return new(dm_sourcecode.BlamePartial)
	case "sourcecode.Branch":
		return new(dm_sourcecode.BranchPartial)
	case "sourcecode.Commit":
		return new(dm_sourcecode.CommitPartial)
	case "sourcecode.PullRequest":
		return new(dm_sourcecode.PullRequestPartial)
	case "sourcecode.PullRequestBranch":
		return new(dm_sourcecode.PullRequestBranchPartial)
	case "sourcecode.PullRequestComment":
		return new(dm_sourcecode.PullRequestCommentPartial)
	case "sourcecode.PullRequestCommit":
		return new(dm_sourcecode.PullRequestCommitPartial)
	case "sourcecode.PullRequestReview":
		return new(dm_sourcecode.PullRequestReviewPartial)
	case "sourcecode.Repo":
		return new(dm_sourcecode.RepoPartial)
	case "sourcecode.RepoError":
		return new(dm_sourcecode.RepoErrorPartial)
	case "sourcecode.RepoWebhook":
		return new(dm_sourcecode.RepoWebhookPartial)
	case "sourcecode.User":
		return new(dm_sourcecode.UserPartial)
	case "web.Hook":
		return new(dm_web.HookPartial)
	case "work.Config":
		return new(dm_work.ConfigPartial)
	case "work.Issue":
		return new(dm_work.IssuePartial)
	case "work.IssueComment":
		return new(dm_work.IssueCommentPartial)
	case "work.IssuePriority":
		return new(dm_work.IssuePriorityPartial)
	case "work.IssueStatus":
		return new(dm_work.IssueStatusPartial)
	case "work.IssueType":
		return new(dm_work.IssueTypePartial)
	case "work.KanbanBoard":
		return new(dm_work.KanbanBoardPartial)
	case "work.Project":
		return new(dm_work.ProjectPartial)
	case "work.ProjectCapability":
		return new(dm_work.ProjectCapabilityPartial)
	case "work.ProjectError":
		return new(dm_work.ProjectErrorPartial)
	case "work.ProjectWebhook":
		return new(dm_work.ProjectWebhookPartial)
	case "work.Sprint":
		return new(dm_work.SprintPartial)
	case "work.User":
		return new(dm_work.UserPartial)
	case "work.agile.Board":
		return new(dm_work_agile.BoardPartial)
	case "work.agile.Kanban":
		return new(dm_work_agile.KanbanPartial)
	case "work.agile.Sprint":
		return new(dm_work_agile.SprintPartial)
	}
	return nil
}

// NewFromTopic returns a new instanceof from a TopicNameType
func NewFromTopic(name datamodel.TopicNameType) datamodel.Model {
	switch name {
	case "agent_UpdateData":
		o := new(dm_agent.UpdateData)
		o.FromMap(map[string]interface{}{})
		return o
	case "calendar_Calendar":
		o := new(dm_calendar.Calendar)
		o.FromMap(map[string]interface{}{})
		return o
	case "calendar_Event":
		o := new(dm_calendar.Event)
		o.FromMap(map[string]interface{}{})
		return o
	case "calendar_User":
		o := new(dm_calendar.User)
		o.FromMap(map[string]interface{}{})
		return o
	case "sourcecode_Repo":
		o := new(dm_sourcecode.Repo)
		o.FromMap(map[string]interface{}{})
		return o
	case "sourcecode_RepoError":
		o := new(dm_sourcecode.RepoError)
		o.FromMap(map[string]interface{}{})
		return o
	case "sourcecode_RepoWebhook":
		o := new(dm_sourcecode.RepoWebhook)
		o.FromMap(map[string]interface{}{})
		return o
	case "sourcecode_User":
		o := new(dm_sourcecode.User)
		o.FromMap(map[string]interface{}{})
		return o
	case "work_Config":
		o := new(dm_work.Config)
		o.FromMap(map[string]interface{}{})
		return o
	case "work_KanbanBoard":
		o := new(dm_work.KanbanBoard)
		o.FromMap(map[string]interface{}{})
		return o
	case "work_Project":
		o := new(dm_work.Project)
		o.FromMap(map[string]interface{}{})
		return o
	case "work_ProjectCapability":
		o := new(dm_work.ProjectCapability)
		o.FromMap(map[string]interface{}{})
		return o
	case "work_ProjectError":
		o := new(dm_work.ProjectError)
		o.FromMap(map[string]interface{}{})
		return o
	case "work_ProjectWebhook":
		o := new(dm_work.ProjectWebhook)
		o.FromMap(map[string]interface{}{})
		return o
	case "work_Sprint":
		o := new(dm_work.Sprint)
		o.FromMap(map[string]interface{}{})
		return o
	case "work_User":
		o := new(dm_work.User)
		o.FromMap(map[string]interface{}{})
		return o
	case "work_agile_Board":
		o := new(dm_work_agile.Board)
		o.FromMap(map[string]interface{}{})
		return o
	case "work_agile_Kanban":
		o := new(dm_work_agile.Kanban)
		o.FromMap(map[string]interface{}{})
		return o
	case "work_agile_Sprint":
		o := new(dm_work_agile.Sprint)
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
		datamodel.TopicNameType("agent_UpdateData"),
		datamodel.TopicNameType("calendar_Calendar"),
		datamodel.TopicNameType("calendar_Event"),
		datamodel.TopicNameType("calendar_User"),
		datamodel.TopicNameType("sourcecode_Repo"),
		datamodel.TopicNameType("sourcecode_RepoError"),
		datamodel.TopicNameType("sourcecode_RepoWebhook"),
		datamodel.TopicNameType("sourcecode_User"),
		datamodel.TopicNameType("work_Config"),
		datamodel.TopicNameType("work_KanbanBoard"),
		datamodel.TopicNameType("work_Project"),
		datamodel.TopicNameType("work_ProjectCapability"),
		datamodel.TopicNameType("work_ProjectError"),
		datamodel.TopicNameType("work_ProjectWebhook"),
		datamodel.TopicNameType("work_Sprint"),
		datamodel.TopicNameType("work_User"),
		datamodel.TopicNameType("work_agile_Board"),
		datamodel.TopicNameType("work_agile_Kanban"),
		datamodel.TopicNameType("work_agile_Sprint"),
	}
}

// GetModelNames returns an array of model names
func GetModelNames() []datamodel.ModelNameType {
	return []datamodel.ModelNameType{
		datamodel.ModelNameType("agent.CalendarRequest"),
		datamodel.ModelNameType("agent.CalendarResponse"),
		datamodel.ModelNameType("agent.CalendarTrigger"),
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
		datamodel.ModelNameType("agent.Enrollment"),
		datamodel.ModelNameType("agent.Event"),
		datamodel.ModelNameType("agent.Export"),
		datamodel.ModelNameType("agent.ExportData"),
		datamodel.ModelNameType("agent.ExportRequest"),
		datamodel.ModelNameType("agent.ExportResponse"),
		datamodel.ModelNameType("agent.ExportTrigger"),
		datamodel.ModelNameType("agent.IntegrationInstance"),
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
		datamodel.ModelNameType("agent.UpdateData"),
		datamodel.ModelNameType("agent.UpdateRequest"),
		datamodel.ModelNameType("agent.UpdateResponse"),
		datamodel.ModelNameType("agent.UpdateTrigger"),
		datamodel.ModelNameType("agent.WebappIntegrationMutationRequest"),
		datamodel.ModelNameType("agent.WebhookRequest"),
		datamodel.ModelNameType("agent.WebhookResponse"),
		datamodel.ModelNameType("agent.WorkStatusRequest"),
		datamodel.ModelNameType("agent.WorkStatusResponse"),
		datamodel.ModelNameType("agent.WorkStatusTrigger"),
		datamodel.ModelNameType("calendar.Calendar"),
		datamodel.ModelNameType("calendar.Event"),
		datamodel.ModelNameType("calendar.User"),
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
		datamodel.ModelNameType("sourcecode.RepoError"),
		datamodel.ModelNameType("sourcecode.RepoWebhook"),
		datamodel.ModelNameType("sourcecode.User"),
		datamodel.ModelNameType("web.Hook"),
		datamodel.ModelNameType("work.Config"),
		datamodel.ModelNameType("work.Issue"),
		datamodel.ModelNameType("work.IssueComment"),
		datamodel.ModelNameType("work.IssuePriority"),
		datamodel.ModelNameType("work.IssueStatus"),
		datamodel.ModelNameType("work.IssueType"),
		datamodel.ModelNameType("work.KanbanBoard"),
		datamodel.ModelNameType("work.Project"),
		datamodel.ModelNameType("work.ProjectCapability"),
		datamodel.ModelNameType("work.ProjectError"),
		datamodel.ModelNameType("work.ProjectWebhook"),
		datamodel.ModelNameType("work.Sprint"),
		datamodel.ModelNameType("work.User"),
		datamodel.ModelNameType("work.agile.Board"),
		datamodel.ModelNameType("work.agile.Kanban"),
		datamodel.ModelNameType("work.agile.Sprint"),
	}
}
