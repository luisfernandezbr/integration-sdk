// DO NOT EDIT -- generated code

// Package agent - the agent communicates with pinpoint cloud to send integration data
package agent

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"sync"
	"time"

	"github.com/bxcodec/faker"
	"github.com/linkedin/goavro"
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/eventing"
	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
	pstrings "github.com/pinpt/go-common/strings"
)

const (
	// ProjectResponseTopic is the default topic name
	ProjectResponseTopic datamodel.TopicNameType = "agent_ProjectResponse_topic"

	// ProjectResponseStream is the default stream name
	ProjectResponseStream datamodel.TopicNameType = "agent_ProjectResponse_stream"

	// ProjectResponseTable is the default table name
	ProjectResponseTable datamodel.TopicNameType = "agent_ProjectResponse"

	// ProjectResponseModelName is the model name
	ProjectResponseModelName datamodel.ModelNameType = "agent.ProjectResponse"
)

const (
	// ProjectResponseArchitectureColumn is the architecture column name
	ProjectResponseArchitectureColumn = "architecture"
	// ProjectResponseCustomerIDColumn is the customer_id column name
	ProjectResponseCustomerIDColumn = "customer_id"
	// ProjectResponseDataColumn is the data column name
	ProjectResponseDataColumn = "data"
	// ProjectResponseDistroColumn is the distro column name
	ProjectResponseDistroColumn = "distro"
	// ProjectResponseErrorColumn is the error column name
	ProjectResponseErrorColumn = "error"
	// ProjectResponseEventDateColumn is the event_date column name
	ProjectResponseEventDateColumn = "event_date"
	// ProjectResponseEventDateColumnEpochColumn is the epoch column property of the EventDate name
	ProjectResponseEventDateColumnEpochColumn = "event_date->epoch"
	// ProjectResponseEventDateColumnOffsetColumn is the offset column property of the EventDate name
	ProjectResponseEventDateColumnOffsetColumn = "event_date->offset"
	// ProjectResponseEventDateColumnRfc3339Column is the rfc3339 column property of the EventDate name
	ProjectResponseEventDateColumnRfc3339Column = "event_date->rfc3339"
	// ProjectResponseFreeSpaceColumn is the free_space column name
	ProjectResponseFreeSpaceColumn = "free_space"
	// ProjectResponseGoVersionColumn is the go_version column name
	ProjectResponseGoVersionColumn = "go_version"
	// ProjectResponseHostnameColumn is the hostname column name
	ProjectResponseHostnameColumn = "hostname"
	// ProjectResponseIDColumn is the id column name
	ProjectResponseIDColumn = "id"
	// ProjectResponseIntegrationIDColumn is the integration_id column name
	ProjectResponseIntegrationIDColumn = "integration_id"
	// ProjectResponseMemoryColumn is the memory column name
	ProjectResponseMemoryColumn = "memory"
	// ProjectResponseMessageColumn is the message column name
	ProjectResponseMessageColumn = "message"
	// ProjectResponseNumCPUColumn is the num_cpu column name
	ProjectResponseNumCPUColumn = "num_cpu"
	// ProjectResponseOSColumn is the os column name
	ProjectResponseOSColumn = "os"
	// ProjectResponseProjectsColumn is the projects column name
	ProjectResponseProjectsColumn = "projects"
	// ProjectResponseProjectsColumnActiveColumn is the active column property of the Projects name
	ProjectResponseProjectsColumnActiveColumn = "projects->active"
	// ProjectResponseProjectsColumnCategoryColumn is the category column property of the Projects name
	ProjectResponseProjectsColumnCategoryColumn = "projects->category"
	// ProjectResponseProjectsColumnDescriptionColumn is the description column property of the Projects name
	ProjectResponseProjectsColumnDescriptionColumn = "projects->description"
	// ProjectResponseProjectsColumnIdentifierColumn is the identifier column property of the Projects name
	ProjectResponseProjectsColumnIdentifierColumn = "projects->identifier"
	// ProjectResponseProjectsColumnLastIssueColumn is the last_issue column property of the Projects name
	ProjectResponseProjectsColumnLastIssueColumn = "projects->last_issue"
	// ProjectResponseProjectsColumnNameColumn is the name column property of the Projects name
	ProjectResponseProjectsColumnNameColumn = "projects->name"
	// ProjectResponseProjectsColumnProjectIDColumn is the project_id column property of the Projects name
	ProjectResponseProjectsColumnProjectIDColumn = "projects->project_id"
	// ProjectResponseProjectsColumnTotalIssuesColumn is the total_issues column property of the Projects name
	ProjectResponseProjectsColumnTotalIssuesColumn = "projects->total_issues"
	// ProjectResponseProjectsColumnURLColumn is the url column property of the Projects name
	ProjectResponseProjectsColumnURLColumn = "projects->url"
	// ProjectResponseRefIDColumn is the ref_id column name
	ProjectResponseRefIDColumn = "ref_id"
	// ProjectResponseRefTypeColumn is the ref_type column name
	ProjectResponseRefTypeColumn = "ref_type"
	// ProjectResponseRequestIDColumn is the request_id column name
	ProjectResponseRequestIDColumn = "request_id"
	// ProjectResponseSuccessColumn is the success column name
	ProjectResponseSuccessColumn = "success"
	// ProjectResponseTypeColumn is the type column name
	ProjectResponseTypeColumn = "type"
	// ProjectResponseUUIDColumn is the uuid column name
	ProjectResponseUUIDColumn = "uuid"
	// ProjectResponseVersionColumn is the version column name
	ProjectResponseVersionColumn = "version"
)

// ProjectResponseEventDate represents the object structure for event_date
type ProjectResponseEventDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toProjectResponseEventDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toProjectResponseEventDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	// nested => true prefix => ProjectResponseEventDate name => ProjectResponseEventDate
	switch v := o.(type) {
	case *ProjectResponseEventDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *ProjectResponseEventDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toProjectResponseEventDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toProjectResponseEventDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toProjectResponseEventDateObject(o.Rfc3339, isavro, false, "string"),
	}
}

func (o *ProjectResponseEventDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ProjectResponseEventDate) FromMap(kv map[string]interface{}) {

	if val, ok := kv["epoch"].(int64); ok {
		o.Epoch = val
	} else {
		if val, ok := kv["epoch"]; ok {
			if val == nil {
				o.Epoch = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Epoch = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["offset"].(int64); ok {
		o.Offset = val
	} else {
		if val, ok := kv["offset"]; ok {
			if val == nil {
				o.Offset = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Offset = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["rfc3339"].(string); ok {
		o.Rfc3339 = val
	} else {
		if val, ok := kv["rfc3339"]; ok {
			if val == nil {
				o.Rfc3339 = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// ProjectResponseProjectsLastIssueCreatedDate represents the object structure for created_date
type ProjectResponseProjectsLastIssueCreatedDate struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toProjectResponseProjectsLastIssueCreatedDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toProjectResponseProjectsLastIssueCreatedDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	// nested => true prefix => ProjectResponseProjectsLastIssueCreatedDate name => ProjectResponseProjectsLastIssueCreatedDate
	switch v := o.(type) {
	case *ProjectResponseProjectsLastIssueCreatedDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *ProjectResponseProjectsLastIssueCreatedDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// CustomerID the customer id for the model instance
		"customer_id": toProjectResponseProjectsLastIssueCreatedDateObject(o.CustomerID, isavro, false, "string"),
		// Epoch the date in epoch format
		"epoch": toProjectResponseProjectsLastIssueCreatedDateObject(o.Epoch, isavro, false, "long"),
		// ID the primary key for the model instance
		"id": toProjectResponseProjectsLastIssueCreatedDateObject(o.ID, isavro, false, "string"),
		// Offset the timezone offset from GMT
		"offset": toProjectResponseProjectsLastIssueCreatedDateObject(o.Offset, isavro, false, "long"),
		// RefID the source system id for the model instance
		"ref_id": toProjectResponseProjectsLastIssueCreatedDateObject(o.RefID, isavro, false, "string"),
		// RefType the source system identifier for the model instance
		"ref_type": toProjectResponseProjectsLastIssueCreatedDateObject(o.RefType, isavro, false, "string"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toProjectResponseProjectsLastIssueCreatedDateObject(o.Rfc3339, isavro, false, "string"),
	}
}

func (o *ProjectResponseProjectsLastIssueCreatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ProjectResponseProjectsLastIssueCreatedDate) FromMap(kv map[string]interface{}) {

	if val, ok := kv["customer_id"].(string); ok {
		o.CustomerID = val
	} else {
		if val, ok := kv["customer_id"]; ok {
			if val == nil {
				o.CustomerID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.CustomerID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["epoch"].(int64); ok {
		o.Epoch = val
	} else {
		if val, ok := kv["epoch"]; ok {
			if val == nil {
				o.Epoch = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Epoch = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["id"].(string); ok {
		o.ID = val
	} else {
		if val, ok := kv["id"]; ok {
			if val == nil {
				o.ID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.ID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["offset"].(int64); ok {
		o.Offset = val
	} else {
		if val, ok := kv["offset"]; ok {
			if val == nil {
				o.Offset = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Offset = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["ref_id"].(string); ok {
		o.RefID = val
	} else {
		if val, ok := kv["ref_id"]; ok {
			if val == nil {
				o.RefID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.RefID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["ref_type"].(string); ok {
		o.RefType = val
	} else {
		if val, ok := kv["ref_type"]; ok {
			if val == nil {
				o.RefType = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.RefType = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["rfc3339"].(string); ok {
		o.Rfc3339 = val
	} else {
		if val, ok := kv["rfc3339"]; ok {
			if val == nil {
				o.Rfc3339 = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// ProjectResponseProjectsLastIssueLastUser represents the object structure for last_user
type ProjectResponseProjectsLastIssueLastUser struct {
	// UserID the work project user id
	UserID string `json:"user_id" bson:"user_id" yaml:"user_id" faker:"-"`
	// Name the user name
	Name string `json:"name" bson:"name" yaml:"name" faker:"name"`
	// AvatarURL the avatar url
	AvatarURL string `json:"avatar_url" bson:"avatar_url" yaml:"avatar_url" faker:"avatar"`
}

func toProjectResponseProjectsLastIssueLastUserObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toProjectResponseProjectsLastIssueLastUserObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	// nested => true prefix => ProjectResponseProjectsLastIssueLastUser name => ProjectResponseProjectsLastIssueLastUser
	switch v := o.(type) {
	case *ProjectResponseProjectsLastIssueLastUser:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *ProjectResponseProjectsLastIssueLastUser) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// UserID the work project user id
		"user_id": toProjectResponseProjectsLastIssueLastUserObject(o.UserID, isavro, false, "string"),
		// Name the user name
		"name": toProjectResponseProjectsLastIssueLastUserObject(o.Name, isavro, false, "string"),
		// AvatarURL the avatar url
		"avatar_url": toProjectResponseProjectsLastIssueLastUserObject(o.AvatarURL, isavro, false, "string"),
	}
}

func (o *ProjectResponseProjectsLastIssueLastUser) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ProjectResponseProjectsLastIssueLastUser) FromMap(kv map[string]interface{}) {

	if val, ok := kv["user_id"].(string); ok {
		o.UserID = val
	} else {
		if val, ok := kv["user_id"]; ok {
			if val == nil {
				o.UserID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.UserID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["name"].(string); ok {
		o.Name = val
	} else {
		if val, ok := kv["name"]; ok {
			if val == nil {
				o.Name = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Name = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["avatar_url"].(string); ok {
		o.AvatarURL = val
	} else {
		if val, ok := kv["avatar_url"]; ok {
			if val == nil {
				o.AvatarURL = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.AvatarURL = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// ProjectResponseProjectsLastIssue represents the object structure for last_issue
type ProjectResponseProjectsLastIssue struct {
	// IssueID issue id
	IssueID string `json:"issue_id" bson:"issue_id" yaml:"issue_id" faker:"-"`
	// Identifier the issue key from the source
	Identifier string `json:"identifier" bson:"identifier" yaml:"identifier" faker:"-"`
	// CreatedDate the date of the change
	CreatedDate ProjectResponseProjectsLastIssueCreatedDate `json:"created_date" bson:"created_date" yaml:"created_date" faker:"-"`
	// LastUser the last user
	LastUser ProjectResponseProjectsLastIssueLastUser `json:"last_user" bson:"last_user" yaml:"last_user" faker:"-"`
}

func toProjectResponseProjectsLastIssueObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toProjectResponseProjectsLastIssueObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	// nested => true prefix => ProjectResponseProjectsLastIssue name => ProjectResponseProjectsLastIssue
	switch v := o.(type) {
	case *ProjectResponseProjectsLastIssue:
		return v.ToMap(isavro)

	case ProjectResponseProjectsLastIssueCreatedDate:
		return v.ToMap(isavro)

	case ProjectResponseProjectsLastIssueLastUser:
		return v.ToMap(isavro)
	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *ProjectResponseProjectsLastIssue) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// IssueID issue id
		"issue_id": toProjectResponseProjectsLastIssueObject(o.IssueID, isavro, false, "string"),
		// Identifier the issue key from the source
		"identifier": toProjectResponseProjectsLastIssueObject(o.Identifier, isavro, false, "string"),
		// CreatedDate the date of the change
		"created_date": toProjectResponseProjectsLastIssueObject(o.CreatedDate, isavro, false, "created_date"),
		// LastUser the last user
		"last_user": toProjectResponseProjectsLastIssueObject(o.LastUser, isavro, false, "last_user"),
	}
}

func (o *ProjectResponseProjectsLastIssue) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ProjectResponseProjectsLastIssue) FromMap(kv map[string]interface{}) {

	if val, ok := kv["issue_id"].(string); ok {
		o.IssueID = val
	} else {
		if val, ok := kv["issue_id"]; ok {
			if val == nil {
				o.IssueID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.IssueID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["identifier"].(string); ok {
		o.Identifier = val
	} else {
		if val, ok := kv["identifier"]; ok {
			if val == nil {
				o.Identifier = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Identifier = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["created_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.CreatedDate.FromMap(kv)
		} else if sv, ok := val.(ProjectResponseProjectsLastIssueCreatedDate); ok {
			// struct
			o.CreatedDate = sv
		} else if sp, ok := val.(*ProjectResponseProjectsLastIssueCreatedDate); ok {
			// struct pointer
			o.CreatedDate = *sp
		}
	} else {
		o.CreatedDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["last_user"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastUser.FromMap(kv)
		} else if sv, ok := val.(ProjectResponseProjectsLastIssueLastUser); ok {
			// struct
			o.LastUser = sv
		} else if sp, ok := val.(*ProjectResponseProjectsLastIssueLastUser); ok {
			// struct pointer
			o.LastUser = *sp
		}
	} else {
		o.LastUser.FromMap(map[string]interface{}{})
	}

	o.setDefaults(false)
}

// ProjectResponseProjects represents the object structure for projects
type ProjectResponseProjects struct {
	// Active the status of the project
	Active bool `json:"active" bson:"active" yaml:"active" faker:"-"`
	// Category the project category
	Category *string `json:"category" bson:"category" yaml:"category" faker:"-"`
	// Description the description of the project
	Description *string `json:"description" bson:"description" yaml:"description" faker:"-"`
	// Identifier the common identifier for the project
	Identifier string `json:"identifier" bson:"identifier" yaml:"identifier" faker:"-"`
	// LastIssue last issue for this project
	LastIssue ProjectResponseProjectsLastIssue `json:"last_issue" bson:"last_issue" yaml:"last_issue" faker:"-"`
	// Name the name of the project
	Name string `json:"name" bson:"name" yaml:"name" faker:"-"`
	// ProjectID the id of the project
	ProjectID string `json:"project_id" bson:"project_id" yaml:"project_id" faker:"-"`
	// TotalIssues the total issues count for the project
	TotalIssues int64 `json:"total_issues" bson:"total_issues" yaml:"total_issues" faker:"-"`
	// URL the url to the project home page
	URL string `json:"url" bson:"url" yaml:"url" faker:"-"`
}

func toProjectResponseProjectsObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toProjectResponseProjectsObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	// nested => true prefix => ProjectResponseProjects name => ProjectResponseProjects
	switch v := o.(type) {
	case *ProjectResponseProjects:
		return v.ToMap(isavro)

	case ProjectResponseProjectsLastIssue:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *ProjectResponseProjects) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Active the status of the project
		"active": toProjectResponseProjectsObject(o.Active, isavro, false, "boolean"),
		// Category the project category
		"category": toProjectResponseProjectsObject(o.Category, isavro, true, "string"),
		// Description the description of the project
		"description": toProjectResponseProjectsObject(o.Description, isavro, true, "string"),
		// Identifier the common identifier for the project
		"identifier": toProjectResponseProjectsObject(o.Identifier, isavro, false, "string"),
		// LastIssue last issue for this project
		"last_issue": toProjectResponseProjectsObject(o.LastIssue, isavro, false, "last_issue"),
		// Name the name of the project
		"name": toProjectResponseProjectsObject(o.Name, isavro, false, "string"),
		// ProjectID the id of the project
		"project_id": toProjectResponseProjectsObject(o.ProjectID, isavro, false, "string"),
		// TotalIssues the total issues count for the project
		"total_issues": toProjectResponseProjectsObject(o.TotalIssues, isavro, false, "long"),
		// URL the url to the project home page
		"url": toProjectResponseProjectsObject(o.URL, isavro, false, "string"),
	}
}

func (o *ProjectResponseProjects) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ProjectResponseProjects) FromMap(kv map[string]interface{}) {

	if val, ok := kv["active"].(bool); ok {
		o.Active = val
	} else {
		if val, ok := kv["active"]; ok {
			if val == nil {
				o.Active = number.ToBoolAny(nil)
			} else {
				o.Active = number.ToBoolAny(val)
			}
		}
	}

	if val, ok := kv["category"].(*string); ok {
		o.Category = val
	} else if val, ok := kv["category"].(string); ok {
		o.Category = &val
	} else {
		if val, ok := kv["category"]; ok {
			if val == nil {
				o.Category = pstrings.Pointer("")
			} else {
				// if coming in as avro union, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Category = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if val, ok := kv["description"].(*string); ok {
		o.Description = val
	} else if val, ok := kv["description"].(string); ok {
		o.Description = &val
	} else {
		if val, ok := kv["description"]; ok {
			if val == nil {
				o.Description = pstrings.Pointer("")
			} else {
				// if coming in as avro union, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Description = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if val, ok := kv["identifier"].(string); ok {
		o.Identifier = val
	} else {
		if val, ok := kv["identifier"]; ok {
			if val == nil {
				o.Identifier = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Identifier = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["last_issue"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastIssue.FromMap(kv)
		} else if sv, ok := val.(ProjectResponseProjectsLastIssue); ok {
			// struct
			o.LastIssue = sv
		} else if sp, ok := val.(*ProjectResponseProjectsLastIssue); ok {
			// struct pointer
			o.LastIssue = *sp
		}
	} else {
		o.LastIssue.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["name"].(string); ok {
		o.Name = val
	} else {
		if val, ok := kv["name"]; ok {
			if val == nil {
				o.Name = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Name = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["project_id"].(string); ok {
		o.ProjectID = val
	} else {
		if val, ok := kv["project_id"]; ok {
			if val == nil {
				o.ProjectID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.ProjectID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["total_issues"].(int64); ok {
		o.TotalIssues = val
	} else {
		if val, ok := kv["total_issues"]; ok {
			if val == nil {
				o.TotalIssues = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.TotalIssues = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["url"].(string); ok {
		o.URL = val
	} else {
		if val, ok := kv["url"]; ok {
			if val == nil {
				o.URL = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.URL = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// ProjectResponseType is the enumeration type for type
type ProjectResponseType int32

// String returns the string value for Type
func (v ProjectResponseType) String() string {
	switch int32(v) {
	case 0:
		return "ENROLL"
	case 1:
		return "PING"
	case 2:
		return "CRASH"
	case 3:
		return "INTEGRATION"
	case 4:
		return "EXPORT"
	case 5:
		return "PROJECT"
	case 6:
		return "REPO"
	case 7:
		return "USER"
	}
	return "unset"
}

const (
	// TypeEnroll is the enumeration value for enroll
	ProjectResponseTypeEnroll ProjectResponseType = 0
	// TypePing is the enumeration value for ping
	ProjectResponseTypePing ProjectResponseType = 1
	// TypeCrash is the enumeration value for crash
	ProjectResponseTypeCrash ProjectResponseType = 2
	// TypeIntegration is the enumeration value for integration
	ProjectResponseTypeIntegration ProjectResponseType = 3
	// TypeExport is the enumeration value for export
	ProjectResponseTypeExport ProjectResponseType = 4
	// TypeProject is the enumeration value for project
	ProjectResponseTypeProject ProjectResponseType = 5
	// TypeRepo is the enumeration value for repo
	ProjectResponseTypeRepo ProjectResponseType = 6
	// TypeUser is the enumeration value for user
	ProjectResponseTypeUser ProjectResponseType = 7
)

// ProjectResponse an agent response to an action request adding project(s)
type ProjectResponse struct {
	// Architecture the architecture of the agent machine
	Architecture string `json:"architecture" bson:"architecture" yaml:"architecture" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Data extra data that is specific about this event
	Data *string `json:"data" bson:"data" yaml:"data" faker:"-"`
	// Distro the agent os distribution
	Distro string `json:"distro" bson:"distro" yaml:"distro" faker:"-"`
	// Error an error message related to this event
	Error *string `json:"error" bson:"error" yaml:"error" faker:"-"`
	// EventDate the date of the event
	EventDate ProjectResponseEventDate `json:"event_date" bson:"event_date" yaml:"event_date" faker:"-"`
	// FreeSpace the amount of free space in bytes for the agent machine
	FreeSpace int64 `json:"free_space" bson:"free_space" yaml:"free_space" faker:"-"`
	// GoVersion the go version that the agent build was built with
	GoVersion string `json:"go_version" bson:"go_version" yaml:"go_version" faker:"-"`
	// Hostname the agent hostname
	Hostname string `json:"hostname" bson:"hostname" yaml:"hostname" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationID the integration id
	IntegrationID string `json:"integration_id" bson:"integration_id" yaml:"integration_id" faker:"-"`
	// Memory the amount of memory in bytes for the agent machine
	Memory int64 `json:"memory" bson:"memory" yaml:"memory" faker:"-"`
	// Message a message related to this event
	Message string `json:"message" bson:"message" yaml:"message" faker:"-"`
	// NumCPU the number of CPU the agent is running
	NumCPU int64 `json:"num_cpu" bson:"num_cpu" yaml:"num_cpu" faker:"-"`
	// OS the agent operating system
	OS string `json:"os" bson:"os" yaml:"os" faker:"-"`
	// Projects the projects exported
	Projects []ProjectResponseProjects `json:"projects" bson:"projects" yaml:"projects" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RequestID the request id that this response is correlated to
	RequestID string `json:"request_id" bson:"request_id" yaml:"request_id" faker:"-"`
	// Success if the response was successful
	Success bool `json:"success" bson:"success" yaml:"success" faker:"-"`
	// Type the type of event
	Type ProjectResponseType `json:"type" bson:"type" yaml:"type" faker:"-"`
	// UUID the agent unique identifier
	UUID string `json:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
	// Version the agent version
	Version string `json:"version" bson:"version" yaml:"version" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*ProjectResponse)(nil)

func toProjectResponseObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toProjectResponseObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	// nested => false prefix => ProjectResponse name => ProjectResponse
	switch v := o.(type) {
	case *ProjectResponse:
		return v.ToMap(isavro)

	case ProjectResponseEventDate:
		return v.ToMap(isavro)

	case []ProjectResponseProjects:
		return v

		// is nested enum Type
	case ProjectResponseType:
		return v.String()

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

// String returns a string representation of ProjectResponse
func (o *ProjectResponse) String() string {
	return fmt.Sprintf("agent.ProjectResponse<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *ProjectResponse) GetTopicName() datamodel.TopicNameType {
	return ProjectResponseTopic
}

// GetModelName returns the name of the model
func (o *ProjectResponse) GetModelName() datamodel.ModelNameType {
	return ProjectResponseModelName
}

func (o *ProjectResponse) setDefaults(frommap bool) {
	if o.Data == nil {
		o.Data = &emptyString
	}
	if o.Error == nil {
		o.Error = &emptyString
	}
	if o.Projects == nil {
		o.Projects = make([]ProjectResponseProjects, 0)
	}

	o.GetID()

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *ProjectResponse) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("ProjectResponse", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *ProjectResponse) GetTopicKey() string {
	var i interface{} = o.UUID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *ProjectResponse) GetTimestamp() time.Time {
	var dt interface{} = o.EventDate
	switch v := dt.(type) {
	case int64:
		return datetime.DateFromEpoch(v).UTC()
	case string:
		tv, err := datetime.ISODateToTime(v)
		if err != nil {
			panic(err)
		}
		return tv.UTC()
	case time.Time:
		return v.UTC()
	case ProjectResponseEventDate:
		return datetime.DateFromEpoch(v.Epoch)
	}
	panic("not sure how to handle the date time format for ProjectResponse")
}

// GetRefID returns the RefID for the object
func (o *ProjectResponse) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *ProjectResponse) IsMaterialized() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *ProjectResponse) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *ProjectResponse) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *ProjectResponse) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = ProjectResponseModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *ProjectResponse) GetTopicConfig() *datamodel.ModelTopicConfig {
	retention, err := time.ParseDuration("168h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 168h0m0s. " + err.Error())
	}

	ttl, err := time.ParseDuration("0s")
	if err != nil {
		ttl = 0
	}
	return &datamodel.ModelTopicConfig{
		Key:               "uuid",
		Timestamp:         "event_date",
		NumPartitions:     8,
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetStateKey returns a key for use in state store
func (o *ProjectResponse) GetStateKey() string {
	key := "uuid"
	return fmt.Sprintf("%s_%s", key, o.GetID())
}

// GetCustomerID will return the customer_id
func (o *ProjectResponse) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of ProjectResponse
func (o *ProjectResponse) Clone() datamodel.Model {
	c := new(ProjectResponse)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *ProjectResponse) Anon() datamodel.Model {
	c := new(ProjectResponse)
	if err := faker.FakeData(c); err != nil {
		panic("couldn't create anon version of object: " + err.Error())
	}
	kv := c.ToMap()
	for k, v := range o.ToMap() {
		if _, ok := kv[k]; !ok {
			kv[k] = v
		}
	}
	c.FromMap(kv)
	return c
}

// MarshalJSON returns the bytes for marshaling to json
func (o *ProjectResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *ProjectResponse) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

var cachedCodecProjectResponse *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *ProjectResponse) GetAvroCodec() *goavro.Codec {
	if cachedCodecProjectResponse == nil {
		c, err := GetProjectResponseAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecProjectResponse = c
	}
	return cachedCodecProjectResponse
}

// ToAvroBinary returns the data as Avro binary data
func (o *ProjectResponse) ToAvroBinary() ([]byte, *goavro.Codec, error) {
	kv := o.ToMap(true)
	jbuf, _ := json.Marshal(kv)
	codec := o.GetAvroCodec()
	native, _, err := codec.NativeFromTextual(jbuf)
	if err != nil {
		return nil, nil, err
	}
	// Convert native Go form to binary Avro data
	buf, err := codec.BinaryFromNative(nil, native)
	return buf, codec, err
}

// FromAvroBinary will convert from Avro binary data into data in this object
func (o *ProjectResponse) FromAvroBinary(value []byte) error {
	var nullHeader = []byte{byte(0)}
	// if this still has the schema encoded in the header, move past it to the avro payload
	if bytes.HasPrefix(value, nullHeader) {
		value = value[5:]
	}
	kv, _, err := o.GetAvroCodec().NativeFromBinary(value)
	if err != nil {
		return err
	}
	o.FromMap(kv.(map[string]interface{}))
	return nil
}

// Stringify returns the object in JSON format as a string
func (o *ProjectResponse) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two ProjectResponse objects are equal
func (o *ProjectResponse) IsEqual(other *ProjectResponse) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *ProjectResponse) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
		if o.Projects == nil {
			o.Projects = make([]ProjectResponseProjects, 0)
		}
	}
	o.setDefaults(true)
	return map[string]interface{}{
		"architecture":   toProjectResponseObject(o.Architecture, isavro, false, "string"),
		"customer_id":    toProjectResponseObject(o.CustomerID, isavro, false, "string"),
		"data":           toProjectResponseObject(o.Data, isavro, true, "string"),
		"distro":         toProjectResponseObject(o.Distro, isavro, false, "string"),
		"error":          toProjectResponseObject(o.Error, isavro, true, "string"),
		"event_date":     toProjectResponseObject(o.EventDate, isavro, false, "event_date"),
		"free_space":     toProjectResponseObject(o.FreeSpace, isavro, false, "long"),
		"go_version":     toProjectResponseObject(o.GoVersion, isavro, false, "string"),
		"hostname":       toProjectResponseObject(o.Hostname, isavro, false, "string"),
		"id":             toProjectResponseObject(o.ID, isavro, false, "string"),
		"integration_id": toProjectResponseObject(o.IntegrationID, isavro, false, "string"),
		"memory":         toProjectResponseObject(o.Memory, isavro, false, "long"),
		"message":        toProjectResponseObject(o.Message, isavro, false, "string"),
		"num_cpu":        toProjectResponseObject(o.NumCPU, isavro, false, "long"),
		"os":             toProjectResponseObject(o.OS, isavro, false, "string"),
		"projects":       toProjectResponseObject(o.Projects, isavro, false, "projects"),
		"ref_id":         toProjectResponseObject(o.RefID, isavro, false, "string"),
		"ref_type":       toProjectResponseObject(o.RefType, isavro, false, "string"),
		"request_id":     toProjectResponseObject(o.RequestID, isavro, false, "string"),
		"success":        toProjectResponseObject(o.Success, isavro, false, "boolean"),
		"type":           toProjectResponseObject(o.Type, isavro, false, "type"),
		"uuid":           toProjectResponseObject(o.UUID, isavro, false, "string"),
		"version":        toProjectResponseObject(o.Version, isavro, false, "string"),
		"hashcode":       toProjectResponseObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *ProjectResponse) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["architecture"].(string); ok {
		o.Architecture = val
	} else {
		if val, ok := kv["architecture"]; ok {
			if val == nil {
				o.Architecture = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Architecture = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["customer_id"].(string); ok {
		o.CustomerID = val
	} else {
		if val, ok := kv["customer_id"]; ok {
			if val == nil {
				o.CustomerID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.CustomerID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["data"].(*string); ok {
		o.Data = val
	} else if val, ok := kv["data"].(string); ok {
		o.Data = &val
	} else {
		if val, ok := kv["data"]; ok {
			if val == nil {
				o.Data = pstrings.Pointer("")
			} else {
				// if coming in as avro union, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Data = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if val, ok := kv["distro"].(string); ok {
		o.Distro = val
	} else {
		if val, ok := kv["distro"]; ok {
			if val == nil {
				o.Distro = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Distro = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["error"].(*string); ok {
		o.Error = val
	} else if val, ok := kv["error"].(string); ok {
		o.Error = &val
	} else {
		if val, ok := kv["error"]; ok {
			if val == nil {
				o.Error = pstrings.Pointer("")
			} else {
				// if coming in as avro union, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Error = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if val, ok := kv["event_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.EventDate.FromMap(kv)
		} else if sv, ok := val.(ProjectResponseEventDate); ok {
			// struct
			o.EventDate = sv
		} else if sp, ok := val.(*ProjectResponseEventDate); ok {
			// struct pointer
			o.EventDate = *sp
		}
	} else {
		o.EventDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["free_space"].(int64); ok {
		o.FreeSpace = val
	} else {
		if val, ok := kv["free_space"]; ok {
			if val == nil {
				o.FreeSpace = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.FreeSpace = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["go_version"].(string); ok {
		o.GoVersion = val
	} else {
		if val, ok := kv["go_version"]; ok {
			if val == nil {
				o.GoVersion = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.GoVersion = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["hostname"].(string); ok {
		o.Hostname = val
	} else {
		if val, ok := kv["hostname"]; ok {
			if val == nil {
				o.Hostname = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Hostname = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["id"].(string); ok {
		o.ID = val
	} else {
		if val, ok := kv["id"]; ok {
			if val == nil {
				o.ID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.ID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["integration_id"].(string); ok {
		o.IntegrationID = val
	} else {
		if val, ok := kv["integration_id"]; ok {
			if val == nil {
				o.IntegrationID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.IntegrationID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["memory"].(int64); ok {
		o.Memory = val
	} else {
		if val, ok := kv["memory"]; ok {
			if val == nil {
				o.Memory = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Memory = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["message"].(string); ok {
		o.Message = val
	} else {
		if val, ok := kv["message"]; ok {
			if val == nil {
				o.Message = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Message = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["num_cpu"].(int64); ok {
		o.NumCPU = val
	} else {
		if val, ok := kv["num_cpu"]; ok {
			if val == nil {
				o.NumCPU = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.NumCPU = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["os"].(string); ok {
		o.OS = val
	} else {
		if val, ok := kv["os"]; ok {
			if val == nil {
				o.OS = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.OS = fmt.Sprintf("%v", val)
			}
		}
	}

	if o == nil {

		o.Projects = make([]ProjectResponseProjects, 0)

	}
	if val, ok := kv["projects"]; ok {
		if sv, ok := val.([]ProjectResponseProjects); ok {
			o.Projects = sv
		} else if sp, ok := val.([]*ProjectResponseProjects); ok {
			o.Projects = o.Projects[:0]
			for _, e := range sp {
				o.Projects = append(o.Projects, *e)
			}
		}
	}

	if val, ok := kv["ref_id"].(string); ok {
		o.RefID = val
	} else {
		if val, ok := kv["ref_id"]; ok {
			if val == nil {
				o.RefID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.RefID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["ref_type"].(string); ok {
		o.RefType = val
	} else {
		if val, ok := kv["ref_type"]; ok {
			if val == nil {
				o.RefType = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.RefType = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["request_id"].(string); ok {
		o.RequestID = val
	} else {
		if val, ok := kv["request_id"]; ok {
			if val == nil {
				o.RequestID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.RequestID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["success"].(bool); ok {
		o.Success = val
	} else {
		if val, ok := kv["success"]; ok {
			if val == nil {
				o.Success = number.ToBoolAny(nil)
			} else {
				o.Success = number.ToBoolAny(val)
			}
		}
	}

	if val, ok := kv["type"].(ProjectResponseType); ok {
		o.Type = val
	} else {
		if em, ok := kv["type"].(map[string]interface{}); ok {
			ev := em["agent.type"].(string)
			switch ev {
			case "enroll", "ENROLL":
				o.Type = 0
			case "ping", "PING":
				o.Type = 1
			case "crash", "CRASH":
				o.Type = 2
			case "integration", "INTEGRATION":
				o.Type = 3
			case "export", "EXPORT":
				o.Type = 4
			case "project", "PROJECT":
				o.Type = 5
			case "repo", "REPO":
				o.Type = 6
			case "user", "USER":
				o.Type = 7
			}
		}
		if em, ok := kv["type"].(string); ok {
			switch em {
			case "enroll", "ENROLL":
				o.Type = 0
			case "ping", "PING":
				o.Type = 1
			case "crash", "CRASH":
				o.Type = 2
			case "integration", "INTEGRATION":
				o.Type = 3
			case "export", "EXPORT":
				o.Type = 4
			case "project", "PROJECT":
				o.Type = 5
			case "repo", "REPO":
				o.Type = 6
			case "user", "USER":
				o.Type = 7
			}
		}
	}

	if val, ok := kv["uuid"].(string); ok {
		o.UUID = val
	} else {
		if val, ok := kv["uuid"]; ok {
			if val == nil {
				o.UUID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.UUID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["version"].(string); ok {
		o.Version = val
	} else {
		if val, ok := kv["version"]; ok {
			if val == nil {
				o.Version = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Version = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *ProjectResponse) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Architecture)
	args = append(args, o.CustomerID)
	args = append(args, o.Data)
	args = append(args, o.Distro)
	args = append(args, o.Error)
	args = append(args, o.EventDate)
	args = append(args, o.FreeSpace)
	args = append(args, o.GoVersion)
	args = append(args, o.Hostname)
	args = append(args, o.ID)
	args = append(args, o.IntegrationID)
	args = append(args, o.Memory)
	args = append(args, o.Message)
	args = append(args, o.NumCPU)
	args = append(args, o.OS)
	args = append(args, o.Projects)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RequestID)
	args = append(args, o.Success)
	args = append(args, o.Type)
	args = append(args, o.UUID)
	args = append(args, o.Version)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetProjectResponseAvroSchemaSpec creates the avro schema specification for ProjectResponse
func GetProjectResponseAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "agent",
		"name":      "ProjectResponse",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "hashcode",
				"type": "string",
			},
			map[string]interface{}{
				"name": "architecture",
				"type": "string",
			},
			map[string]interface{}{
				"name": "customer_id",
				"type": "string",
			},
			map[string]interface{}{
				"name":    "data",
				"type":    []interface{}{"null", "string"},
				"default": nil,
			},
			map[string]interface{}{
				"name": "distro",
				"type": "string",
			},
			map[string]interface{}{
				"name":    "error",
				"type":    []interface{}{"null", "string"},
				"default": nil,
			},
			map[string]interface{}{
				"name": "event_date",
				"type": map[string]interface{}{"name": "event_date", "fields": []interface{}{map[string]interface{}{"type": "long", "name": "epoch", "doc": "the date in epoch format"}, map[string]interface{}{"type": "long", "name": "offset", "doc": "the timezone offset from GMT"}, map[string]interface{}{"type": "string", "name": "rfc3339", "doc": "the date in RFC3339 format"}}, "doc": "the date of the event", "type": "record"},
			},
			map[string]interface{}{
				"name": "free_space",
				"type": "long",
			},
			map[string]interface{}{
				"name": "go_version",
				"type": "string",
			},
			map[string]interface{}{
				"name": "hostname",
				"type": "string",
			},
			map[string]interface{}{
				"name": "id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "integration_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "memory",
				"type": "long",
			},
			map[string]interface{}{
				"name": "message",
				"type": "string",
			},
			map[string]interface{}{
				"name": "num_cpu",
				"type": "long",
			},
			map[string]interface{}{
				"name": "os",
				"type": "string",
			},
			map[string]interface{}{
				"name": "projects",
				"type": map[string]interface{}{"type": "array", "name": "projects", "items": map[string]interface{}{"type": "record", "name": "projects", "fields": []interface{}{map[string]interface{}{"type": "boolean", "name": "active", "doc": "the status of the project"}, map[string]interface{}{"default": nil, "type": []interface{}{"null", "string"}, "name": "category", "doc": "the project category"}, map[string]interface{}{"default": nil, "type": []interface{}{"null", "string"}, "name": "description", "doc": "the description of the project"}, map[string]interface{}{"type": "string", "name": "identifier", "doc": "the common identifier for the project"}, map[string]interface{}{"name": "last_issue", "doc": "last issue for this project", "type": map[string]interface{}{"type": "record", "name": "projects.last_issue", "fields": []interface{}{map[string]interface{}{"type": "string", "name": "issue_id", "doc": "issue id"}, map[string]interface{}{"type": "string", "name": "identifier", "doc": "the issue key from the source"}, map[string]interface{}{"type": map[string]interface{}{"type": "record", "name": "last_issue.created_date", "fields": []interface{}{map[string]interface{}{"name": "customer_id", "doc": "the customer id for the model instance", "type": "string"}, map[string]interface{}{"name": "epoch", "doc": "the date in epoch format", "type": "long"}, map[string]interface{}{"type": "string", "name": "id", "doc": "the primary key for the model instance"}, map[string]interface{}{"name": "offset", "doc": "the timezone offset from GMT", "type": "long"}, map[string]interface{}{"type": "string", "name": "ref_id", "doc": "the source system id for the model instance"}, map[string]interface{}{"doc": "the source system identifier for the model instance", "type": "string", "name": "ref_type"}, map[string]interface{}{"name": "rfc3339", "doc": "the date in RFC3339 format", "type": "string"}}, "doc": "the date of the change"}, "name": "created_date", "doc": "the date of the change"}, map[string]interface{}{"type": map[string]interface{}{"type": "record", "name": "last_issue.last_user", "fields": []interface{}{map[string]interface{}{"type": "string", "name": "user_id", "doc": "the work project user id"}, map[string]interface{}{"doc": "the user name", "type": "string", "name": "name"}, map[string]interface{}{"type": "string", "name": "avatar_url", "doc": "the avatar url"}}, "doc": "the last user"}, "name": "last_user", "doc": "the last user"}}, "doc": "last issue for this project"}}, map[string]interface{}{"name": "name", "doc": "the name of the project", "type": "string"}, map[string]interface{}{"type": "string", "name": "project_id", "doc": "the id of the project"}, map[string]interface{}{"type": "long", "name": "total_issues", "doc": "the total issues count for the project"}, map[string]interface{}{"type": "string", "name": "url", "doc": "the url to the project home page"}}, "doc": "the projects exported"}},
			},
			map[string]interface{}{
				"name": "ref_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "ref_type",
				"type": "string",
			},
			map[string]interface{}{
				"name": "request_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "success",
				"type": "boolean",
			},
			map[string]interface{}{
				"name": "type",
				"type": map[string]interface{}{
					"type":    "enum",
					"name":    "type",
					"symbols": []interface{}{"ENROLL", "PING", "CRASH", "INTEGRATION", "EXPORT", "PROJECT", "REPO", "USER"},
				},
			},
			map[string]interface{}{
				"name": "uuid",
				"type": "string",
			},
			map[string]interface{}{
				"name": "version",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *ProjectResponse) GetEventAPIConfig() datamodel.EventAPIConfig {
	return datamodel.EventAPIConfig{
		Publish: datamodel.EventAPIPublish{
			Public: false,
		},
		Subscribe: datamodel.EventAPISubscribe{
			Public: false,
			Key:    "",
		},
	}
}

// GetProjectResponseAvroSchema creates the avro schema for ProjectResponse
func GetProjectResponseAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetProjectResponseAvroSchemaSpec())
}

// TransformProjectResponseFunc is a function for transforming ProjectResponse during processing
type TransformProjectResponseFunc func(input *ProjectResponse) (*ProjectResponse, error)

// NewProjectResponsePipe creates a pipe for processing ProjectResponse items
func NewProjectResponsePipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformProjectResponseFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewProjectResponseInputStream(input, errors)
	var stream chan ProjectResponse
	if len(transforms) > 0 {
		stream = make(chan ProjectResponse, 1000)
	} else {
		stream = inch
	}
	outdone := NewProjectResponseOutputStream(output, stream, errors)
	go func() {
		if len(transforms) > 0 {
			var stop bool
			for item := range inch {
				input := &item
				for _, transform := range transforms {
					out, err := transform(input)
					if err != nil {
						stop = true
						errors <- err
						break
					}
					if out == nil {
						input = nil
						break
					} else {
						input = out
					}
				}
				if stop {
					break
				}
				if input != nil {
					stream <- *input
				}
			}
			close(stream)
		}
		<-indone
		<-outdone
		done <- true
	}()
	return done
}

// NewProjectResponseInputStreamDir creates a channel for reading ProjectResponse as JSON newlines from a directory of files
func NewProjectResponseInputStreamDir(dir string, errors chan<- error, transforms ...TransformProjectResponseFunc) (chan ProjectResponse, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/agent/project_response\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan ProjectResponse)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for project_response")
		ch := make(chan ProjectResponse)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewProjectResponseInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan ProjectResponse)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewProjectResponseInputStreamFile creates an channel for reading ProjectResponse as JSON newlines from filename
func NewProjectResponseInputStreamFile(filename string, errors chan<- error, transforms ...TransformProjectResponseFunc) (chan ProjectResponse, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan ProjectResponse)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	var f io.ReadCloser = of
	if filepath.Ext(filename) == ".gz" {
		gz, err := gzip.NewReader(f)
		if err != nil {
			of.Close()
			errors <- err
			ch := make(chan ProjectResponse)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewProjectResponseInputStream(f, errors, transforms...)
}

// NewProjectResponseInputStream creates an channel for reading ProjectResponse as JSON newlines from stream
func NewProjectResponseInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformProjectResponseFunc) (chan ProjectResponse, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan ProjectResponse, 1000)
	go func() {
		defer func() { stream.Close(); close(ch); done <- true }()
		r := bufio.NewReader(stream)
		for {
			buf, err := r.ReadBytes('\n')
			if err != nil {
				if err == io.EOF {
					return
				}
				errors <- err
				return
			}
			var item ProjectResponse
			if err := json.Unmarshal(buf, &item); err != nil {
				errors <- err
				return
			}
			in := &item
			var skip bool
			for _, transform := range transforms {
				in, err = transform(in)
				if err != nil {
					errors <- err
					return
				}
				if in == nil {
					skip = true
					break
				}
			}
			if !skip {
				ch <- *in
			}
		}
	}()
	return ch, done
}

// NewProjectResponseOutputStreamDir will output json newlines from channel and save in dir
func NewProjectResponseOutputStreamDir(dir string, ch chan ProjectResponse, errors chan<- error, transforms ...TransformProjectResponseFunc) <-chan bool {
	fp := filepath.Join(dir, "/agent/project_response\\.json(\\.gz)?$")
	os.MkdirAll(filepath.Dir(fp), 0777)
	of, err := os.Create(fp)
	if err != nil {
		errors <- err
		done := make(chan bool, 1)
		done <- true
		return done
	}
	gz, err := gzip.NewWriterLevel(of, gzip.BestCompression)
	if err != nil {
		errors <- err
		done := make(chan bool, 1)
		done <- true
		return done
	}
	return NewProjectResponseOutputStream(gz, ch, errors, transforms...)
}

// NewProjectResponseOutputStream will output json newlines from channel to the stream
func NewProjectResponseOutputStream(stream io.WriteCloser, ch chan ProjectResponse, errors chan<- error, transforms ...TransformProjectResponseFunc) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() {
			if gz, ok := stream.(*gzip.Writer); ok {
				gz.Flush()
				gz.Close()
			}
			stream.Close()
			done <- true
		}()
		for item := range ch {
			in := &item
			var skip bool
			var err error
			for _, transform := range transforms {
				in, err = transform(in)
				if err != nil {
					errors <- err
					return
				}
				if in == nil {
					skip = true
					break
				}
			}
			if !skip {
				buf, err := json.Marshal(in)
				if err != nil {
					errors <- err
					return
				}
				stream.Write(buf)
				stream.Write([]byte{'\n'})
			}
		}
	}()
	return done
}

// ProjectResponseSendEvent is an event detail for sending data
type ProjectResponseSendEvent struct {
	ProjectResponse *ProjectResponse
	headers         map[string]string
	time            time.Time
	key             string
}

var _ datamodel.ModelSendEvent = (*ProjectResponseSendEvent)(nil)

// Key is the key to use for the message
func (e *ProjectResponseSendEvent) Key() string {
	if e.key == "" {
		return e.ProjectResponse.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *ProjectResponseSendEvent) Object() datamodel.Model {
	return e.ProjectResponse
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *ProjectResponseSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *ProjectResponseSendEvent) Timestamp() time.Time {
	return e.time
}

// ProjectResponseSendEventOpts is a function handler for setting opts
type ProjectResponseSendEventOpts func(o *ProjectResponseSendEvent)

// WithProjectResponseSendEventKey sets the key value to a value different than the object ID
func WithProjectResponseSendEventKey(key string) ProjectResponseSendEventOpts {
	return func(o *ProjectResponseSendEvent) {
		o.key = key
	}
}

// WithProjectResponseSendEventTimestamp sets the timestamp value
func WithProjectResponseSendEventTimestamp(tv time.Time) ProjectResponseSendEventOpts {
	return func(o *ProjectResponseSendEvent) {
		o.time = tv
	}
}

// WithProjectResponseSendEventHeader sets the timestamp value
func WithProjectResponseSendEventHeader(key, value string) ProjectResponseSendEventOpts {
	return func(o *ProjectResponseSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewProjectResponseSendEvent returns a new ProjectResponseSendEvent instance
func NewProjectResponseSendEvent(o *ProjectResponse, opts ...ProjectResponseSendEventOpts) *ProjectResponseSendEvent {
	res := &ProjectResponseSendEvent{
		ProjectResponse: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewProjectResponseProducer will stream data from the channel
func NewProjectResponseProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		for {
			select {
			case <-ctx.Done():
				return
			case item := <-ch:
				if item == nil {
					empty <- true
					return
				}
				if object, ok := item.Object().(*ProjectResponse); ok {
					binary, codec, err := object.ToAvroBinary()
					if err != nil {
						errors <- fmt.Errorf("error encoding %s to avro binary data. %v", object.String(), err)
						return
					}
					headers := map[string]string{}
					object.SetEventHeaders(headers)
					for k, v := range item.Headers() {
						headers[k] = v
					}
					tv := item.Timestamp()
					if tv.IsZero() {
						tv = object.GetTimestamp() // if not provided in the message, use the objects value
					}
					if tv.IsZero() {
						tv = time.Now() // if its still zero, use the ingest time
					}
					msg := eventing.Message{
						Encoding:  eventing.AvroEncoding,
						Key:       item.Key(),
						Value:     binary,
						Codec:     codec,
						Headers:   headers,
						Timestamp: tv,
						Topic:     object.GetTopicName().String(),
					}
					if err := producer.Send(ctx, msg); err != nil {
						errors <- fmt.Errorf("error sending %s. %v", object.String(), err)
					}
				} else {
					errors <- fmt.Errorf("invalid event received. expected an object of type agent.ProjectResponse but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewProjectResponseConsumer will stream data from the topic into the provided channel
func NewProjectResponseConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object ProjectResponse
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into agent.ProjectResponse: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avro data into agent.ProjectResponse: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for agent.ProjectResponse")
			}

			// ignore messages that have exceeded the TTL
			cfg := object.GetTopicConfig()
			if cfg != nil && cfg.TTL != 0 && msg.Timestamp.UTC().Add(cfg.TTL).Sub(time.Now().UTC()) < 0 {
				return nil
			}
			msg.Codec = object.GetAvroCodec() // match the codec

			ch <- &ProjectResponseReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object ProjectResponse
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &ProjectResponseReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// ProjectResponseReceiveEvent is an event detail for receiving data
type ProjectResponseReceiveEvent struct {
	ProjectResponse *ProjectResponse
	message         eventing.Message
	eof             bool
}

var _ datamodel.ModelReceiveEvent = (*ProjectResponseReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *ProjectResponseReceiveEvent) Object() datamodel.Model {
	return e.ProjectResponse
}

// Message returns the underlying message data for the event
func (e *ProjectResponseReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *ProjectResponseReceiveEvent) EOF() bool {
	return e.eof
}

// ProjectResponseProducer implements the datamodel.ModelEventProducer
type ProjectResponseProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*ProjectResponseProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *ProjectResponseProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *ProjectResponseProducer) Close() error {
	p.mu.Lock()
	closed := p.closed
	p.closed = true
	p.mu.Unlock()
	if !closed {
		close(p.ch)
		<-p.empty
		p.cancel()
		<-p.done
	}
	return nil
}

// NewProducerChannel returns a channel which can be used for producing Model events
func (o *ProjectResponse) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *ProjectResponse) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &ProjectResponseProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewProjectResponseProducer(newctx, producer, ch, errors, empty),
	}
}

// NewProjectResponseProducerChannel returns a channel which can be used for producing Model events
func NewProjectResponseProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewProjectResponseProducerChannelSize(producer, 0, errors)
}

// NewProjectResponseProducerChannelSize returns a channel which can be used for producing Model events
func NewProjectResponseProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &ProjectResponseProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewProjectResponseProducer(newctx, producer, ch, errors, empty),
	}
}

// ProjectResponseConsumer implements the datamodel.ModelEventConsumer
type ProjectResponseConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*ProjectResponseConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *ProjectResponseConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *ProjectResponseConsumer) Close() error {
	c.mu.Lock()
	closed := c.closed
	c.closed = true
	c.mu.Unlock()
	var err error
	if !closed {
		c.callback.Close()
		err = c.consumer.Close()
	}
	return err
}

// NewConsumerChannel returns a consumer channel which can be used to consume Model events
func (o *ProjectResponse) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &ProjectResponseConsumer{
		ch:       ch,
		callback: NewProjectResponseConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewProjectResponseConsumerChannel returns a consumer channel which can be used to consume Model events
func NewProjectResponseConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &ProjectResponseConsumer{
		ch:       ch,
		callback: NewProjectResponseConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
