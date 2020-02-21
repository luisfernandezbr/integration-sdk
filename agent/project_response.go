// DO NOT EDIT -- generated code

// Package agent - the agent communicates with pinpoint cloud to send integration data
package agent

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/bxcodec/faker"
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
	pstrings "github.com/pinpt/go-common/strings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (

	// ProjectResponseTable is the default table name
	ProjectResponseTable datamodel.ModelNameType = "agent_projectresponse"

	// ProjectResponseModelName is the model name
	ProjectResponseModelName datamodel.ModelNameType = "agent.ProjectResponse"
)

// ProjectResponseEventDate represents the object structure for event_date
type ProjectResponseEventDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toProjectResponseEventDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ProjectResponseEventDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *ProjectResponseEventDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toProjectResponseEventDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toProjectResponseEventDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toProjectResponseEventDateObject(o.Rfc3339, false),
	}
}

func (o *ProjectResponseEventDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ProjectResponseEventDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// ProjectResponseLastExportDate represents the object structure for last_export_date
type ProjectResponseLastExportDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toProjectResponseLastExportDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ProjectResponseLastExportDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *ProjectResponseLastExportDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toProjectResponseLastExportDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toProjectResponseLastExportDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toProjectResponseLastExportDateObject(o.Rfc3339, false),
	}
}

func (o *ProjectResponseLastExportDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ProjectResponseLastExportDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// ProjectResponseProjectsError is the enumeration type for error
type ProjectResponseProjectsError int32

// UnmarshalBSONValue for unmarshaling value
func (v *ProjectResponseProjectsError) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = ProjectResponseProjectsError(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "NONE":
			*v = ProjectResponseProjectsError(0)
		case "PERMISSIONS":
			*v = ProjectResponseProjectsError(1)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v ProjectResponseProjectsError) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "NONE":
		v = 0
	case "PERMISSIONS":
		v = 1
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v ProjectResponseProjectsError) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("NONE")
	case 1:
		return json.Marshal("PERMISSIONS")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for ProjectsError
func (v ProjectResponseProjectsError) String() string {
	switch int32(v) {
	case 0:
		return "NONE"
	case 1:
		return "PERMISSIONS"
	}
	return "unset"
}

const (
	// ProjectsErrorNONE is the enumeration value for NONE
	ProjectResponseProjectsErrorNONE ProjectResponseProjectsError = 0
	// ProjectsErrorPERMISSIONS is the enumeration value for PERMISSIONS
	ProjectResponseProjectsErrorPERMISSIONS ProjectResponseProjectsError = 1
)

// ProjectResponseProjectsLastIssueCreatedDate represents the object structure for created_date
type ProjectResponseProjectsLastIssueCreatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toProjectResponseProjectsLastIssueCreatedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ProjectResponseProjectsLastIssueCreatedDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *ProjectResponseProjectsLastIssueCreatedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toProjectResponseProjectsLastIssueCreatedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toProjectResponseProjectsLastIssueCreatedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toProjectResponseProjectsLastIssueCreatedDateObject(o.Rfc3339, false),
	}
}

func (o *ProjectResponseProjectsLastIssueCreatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ProjectResponseProjectsLastIssueCreatedDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
	UserID string `json:"user_id" codec:"user_id" bson:"user_id" yaml:"user_id" faker:"-"`
	// Name the user name
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"name"`
	// AvatarURL the avatar url
	AvatarURL string `json:"avatar_url" codec:"avatar_url" bson:"avatar_url" yaml:"avatar_url" faker:"avatar"`
}

func toProjectResponseProjectsLastIssueLastUserObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ProjectResponseProjectsLastIssueLastUser:
		return v.ToMap()

	default:
		return o
	}
}

func (o *ProjectResponseProjectsLastIssueLastUser) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// UserID the work project user id
		"user_id": toProjectResponseProjectsLastIssueLastUserObject(o.UserID, false),
		// Name the user name
		"name": toProjectResponseProjectsLastIssueLastUserObject(o.Name, false),
		// AvatarURL the avatar url
		"avatar_url": toProjectResponseProjectsLastIssueLastUserObject(o.AvatarURL, false),
	}
}

func (o *ProjectResponseProjectsLastIssueLastUser) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ProjectResponseProjectsLastIssueLastUser) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["user_id"].(string); ok {
		o.UserID = val
	} else {
		if val, ok := kv["user_id"]; ok {
			if val == nil {
				o.UserID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
	IssueID string `json:"issue_id" codec:"issue_id" bson:"issue_id" yaml:"issue_id" faker:"-"`
	// Identifier the issue key from the source
	Identifier string `json:"identifier" codec:"identifier" bson:"identifier" yaml:"identifier" faker:"-"`
	// CreatedDate the date of the change
	CreatedDate ProjectResponseProjectsLastIssueCreatedDate `json:"created_date" codec:"created_date" bson:"created_date" yaml:"created_date" faker:"-"`
	// LastUser the last user
	LastUser ProjectResponseProjectsLastIssueLastUser `json:"last_user" codec:"last_user" bson:"last_user" yaml:"last_user" faker:"-"`
}

func toProjectResponseProjectsLastIssueObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ProjectResponseProjectsLastIssue:
		return v.ToMap()

	case ProjectResponseProjectsLastIssueCreatedDate:
		return v.ToMap()

	case ProjectResponseProjectsLastIssueLastUser:
		return v.ToMap()
	default:
		return o
	}
}

func (o *ProjectResponseProjectsLastIssue) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// IssueID issue id
		"issue_id": toProjectResponseProjectsLastIssueObject(o.IssueID, false),
		// Identifier the issue key from the source
		"identifier": toProjectResponseProjectsLastIssueObject(o.Identifier, false),
		// CreatedDate the date of the change
		"created_date": toProjectResponseProjectsLastIssueObject(o.CreatedDate, false),
		// LastUser the last user
		"last_user": toProjectResponseProjectsLastIssueObject(o.LastUser, false),
	}
}

func (o *ProjectResponseProjectsLastIssue) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ProjectResponseProjectsLastIssue) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["issue_id"].(string); ok {
		o.IssueID = val
	} else {
		if val, ok := kv["issue_id"]; ok {
			if val == nil {
				o.IssueID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
	Active bool `json:"active" codec:"active" bson:"active" yaml:"active" faker:"-"`
	// Category the project category
	Category *string `json:"category,omitempty" codec:"category,omitempty" bson:"category" yaml:"category,omitempty" faker:"-"`
	// Description the description of the project
	Description *string `json:"description,omitempty" codec:"description,omitempty" bson:"description" yaml:"description,omitempty" faker:"-"`
	// Error reason why the project is being set to inactive
	Error ProjectResponseProjectsError `json:"error" codec:"error" bson:"error" yaml:"error" faker:"-"`
	// Identifier the common identifier for the project
	Identifier string `json:"identifier" codec:"identifier" bson:"identifier" yaml:"identifier" faker:"-"`
	// LastIssue last issue for this project
	LastIssue ProjectResponseProjectsLastIssue `json:"last_issue" codec:"last_issue" bson:"last_issue" yaml:"last_issue" faker:"-"`
	// Name the name of the project
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
	// RefID the id of the project
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the record type
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// TotalIssues the total issues count for the project
	TotalIssues int64 `json:"total_issues" codec:"total_issues" bson:"total_issues" yaml:"total_issues" faker:"-"`
	// URL the url to the project home page
	URL string `json:"url" codec:"url" bson:"url" yaml:"url" faker:"-"`
}

func toProjectResponseProjectsObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ProjectResponseProjects:
		return v.ToMap()

	case ProjectResponseProjectsError:
		return v.String()

	case ProjectResponseProjectsLastIssue:
		return v.ToMap()

	default:
		return o
	}
}

func (o *ProjectResponseProjects) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Active the status of the project
		"active": toProjectResponseProjectsObject(o.Active, false),
		// Category the project category
		"category": toProjectResponseProjectsObject(o.Category, true),
		// Description the description of the project
		"description": toProjectResponseProjectsObject(o.Description, true),
		// Error reason why the project is being set to inactive
		"error": toProjectResponseProjectsObject(o.Error, false),
		// Identifier the common identifier for the project
		"identifier": toProjectResponseProjectsObject(o.Identifier, false),
		// LastIssue last issue for this project
		"last_issue": toProjectResponseProjectsObject(o.LastIssue, false),
		// Name the name of the project
		"name": toProjectResponseProjectsObject(o.Name, false),
		// RefID the id of the project
		"ref_id": toProjectResponseProjectsObject(o.RefID, false),
		// RefType the record type
		"ref_type": toProjectResponseProjectsObject(o.RefType, false),
		// TotalIssues the total issues count for the project
		"total_issues": toProjectResponseProjectsObject(o.TotalIssues, false),
		// URL the url to the project home page
		"url": toProjectResponseProjectsObject(o.URL, false),
	}
}

func (o *ProjectResponseProjects) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ProjectResponseProjects) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

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
				// if coming in as map, convert it back
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
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Description = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if val, ok := kv["error"].(ProjectResponseProjectsError); ok {
		o.Error = val
	} else {
		if em, ok := kv["error"].(map[string]interface{}); ok {
			ev := em["agent.error"].(string)
			switch ev {
			case "none", "NONE":
				o.Error = 0
			case "permissions", "PERMISSIONS":
				o.Error = 1
			}
		}
		if em, ok := kv["error"].(string); ok {
			switch em {
			case "none", "NONE":
				o.Error = 0
			case "permissions", "PERMISSIONS":
				o.Error = 1
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Name = fmt.Sprintf("%v", val)
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.RefType = fmt.Sprintf("%v", val)
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.URL = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// ProjectResponseType is the enumeration type for type
type ProjectResponseType int32

// UnmarshalBSONValue for unmarshaling value
func (v *ProjectResponseType) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = ProjectResponseType(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "ENROLL":
			*v = ProjectResponseType(0)
		case "PING":
			*v = ProjectResponseType(1)
		case "CRASH":
			*v = ProjectResponseType(2)
		case "LOG":
			*v = ProjectResponseType(3)
		case "INTEGRATION":
			*v = ProjectResponseType(4)
		case "EXPORT":
			*v = ProjectResponseType(5)
		case "PROJECT":
			*v = ProjectResponseType(6)
		case "REPO":
			*v = ProjectResponseType(7)
		case "USER":
			*v = ProjectResponseType(8)
		case "UNINSTALL":
			*v = ProjectResponseType(9)
		case "UPGRADE":
			*v = ProjectResponseType(10)
		case "START":
			*v = ProjectResponseType(11)
		case "STOP":
			*v = ProjectResponseType(12)
		case "PAUSE":
			*v = ProjectResponseType(13)
		case "RESUME":
			*v = ProjectResponseType(14)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v ProjectResponseType) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "ENROLL":
		v = 0
	case "PING":
		v = 1
	case "CRASH":
		v = 2
	case "LOG":
		v = 3
	case "INTEGRATION":
		v = 4
	case "EXPORT":
		v = 5
	case "PROJECT":
		v = 6
	case "REPO":
		v = 7
	case "USER":
		v = 8
	case "UNINSTALL":
		v = 9
	case "UPGRADE":
		v = 10
	case "START":
		v = 11
	case "STOP":
		v = 12
	case "PAUSE":
		v = 13
	case "RESUME":
		v = 14
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v ProjectResponseType) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("ENROLL")
	case 1:
		return json.Marshal("PING")
	case 2:
		return json.Marshal("CRASH")
	case 3:
		return json.Marshal("LOG")
	case 4:
		return json.Marshal("INTEGRATION")
	case 5:
		return json.Marshal("EXPORT")
	case 6:
		return json.Marshal("PROJECT")
	case 7:
		return json.Marshal("REPO")
	case 8:
		return json.Marshal("USER")
	case 9:
		return json.Marshal("UNINSTALL")
	case 10:
		return json.Marshal("UPGRADE")
	case 11:
		return json.Marshal("START")
	case 12:
		return json.Marshal("STOP")
	case 13:
		return json.Marshal("PAUSE")
	case 14:
		return json.Marshal("RESUME")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

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
		return "LOG"
	case 4:
		return "INTEGRATION"
	case 5:
		return "EXPORT"
	case 6:
		return "PROJECT"
	case 7:
		return "REPO"
	case 8:
		return "USER"
	case 9:
		return "UNINSTALL"
	case 10:
		return "UPGRADE"
	case 11:
		return "START"
	case 12:
		return "STOP"
	case 13:
		return "PAUSE"
	case 14:
		return "RESUME"
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
	// TypeLog is the enumeration value for log
	ProjectResponseTypeLog ProjectResponseType = 3
	// TypeIntegration is the enumeration value for integration
	ProjectResponseTypeIntegration ProjectResponseType = 4
	// TypeExport is the enumeration value for export
	ProjectResponseTypeExport ProjectResponseType = 5
	// TypeProject is the enumeration value for project
	ProjectResponseTypeProject ProjectResponseType = 6
	// TypeRepo is the enumeration value for repo
	ProjectResponseTypeRepo ProjectResponseType = 7
	// TypeUser is the enumeration value for user
	ProjectResponseTypeUser ProjectResponseType = 8
	// TypeUninstall is the enumeration value for uninstall
	ProjectResponseTypeUninstall ProjectResponseType = 9
	// TypeUpgrade is the enumeration value for upgrade
	ProjectResponseTypeUpgrade ProjectResponseType = 10
	// TypeStart is the enumeration value for start
	ProjectResponseTypeStart ProjectResponseType = 11
	// TypeStop is the enumeration value for stop
	ProjectResponseTypeStop ProjectResponseType = 12
	// TypePause is the enumeration value for pause
	ProjectResponseTypePause ProjectResponseType = 13
	// TypeResume is the enumeration value for resume
	ProjectResponseTypeResume ProjectResponseType = 14
)

// ProjectResponse an agent response to an action request adding project(s)
type ProjectResponse struct {
	// Architecture the architecture of the agent machine
	Architecture string `json:"architecture" codec:"architecture" bson:"architecture" yaml:"architecture" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Data extra data that is specific about this event
	Data *string `json:"data,omitempty" codec:"data,omitempty" bson:"data" yaml:"data,omitempty" faker:"-"`
	// Distro the agent os distribution
	Distro string `json:"distro" codec:"distro" bson:"distro" yaml:"distro" faker:"-"`
	// Error an error message related to this event
	Error *string `json:"error,omitempty" codec:"error,omitempty" bson:"error" yaml:"error,omitempty" faker:"-"`
	// EventDate the date of the event
	EventDate ProjectResponseEventDate `json:"event_date" codec:"event_date" bson:"event_date" yaml:"event_date" faker:"-"`
	// FreeSpace the amount of free space in bytes for the agent machine
	FreeSpace int64 `json:"free_space" codec:"free_space" bson:"free_space" yaml:"free_space" faker:"-"`
	// GoVersion the go version that the agent build was built with
	GoVersion string `json:"go_version" codec:"go_version" bson:"go_version" yaml:"go_version" faker:"-"`
	// Hostname the agent hostname
	Hostname string `json:"hostname" codec:"hostname" bson:"hostname" yaml:"hostname" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationID the integration id
	IntegrationID string `json:"integration_id" codec:"integration_id" bson:"integration_id" yaml:"integration_id" faker:"-"`
	// LastExportDate the last export date
	LastExportDate ProjectResponseLastExportDate `json:"last_export_date" codec:"last_export_date" bson:"last_export_date" yaml:"last_export_date" faker:"-"`
	// Memory the amount of memory in bytes for the agent machine
	Memory int64 `json:"memory" codec:"memory" bson:"memory" yaml:"memory" faker:"-"`
	// Message a message related to this event
	Message string `json:"message" codec:"message" bson:"message" yaml:"message" faker:"-"`
	// NumCPU the number of CPU the agent is running
	NumCPU int64 `json:"num_cpu" codec:"num_cpu" bson:"num_cpu" yaml:"num_cpu" faker:"-"`
	// OS the agent operating system
	OS string `json:"os" codec:"os" bson:"os" yaml:"os" faker:"-"`
	// Projects the projects exported
	Projects []ProjectResponseProjects `json:"projects" codec:"projects" bson:"projects" yaml:"projects" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RequestID the request id that this response is correlated to
	RequestID string `json:"request_id" codec:"request_id" bson:"request_id" yaml:"request_id" faker:"-"`
	// Success if the response was successful
	Success bool `json:"success" codec:"success" bson:"success" yaml:"success" faker:"-"`
	// SystemID system unique device ID
	SystemID string `json:"system_id" codec:"system_id" bson:"system_id" yaml:"system_id" faker:"-"`
	// Type the type of event
	Type ProjectResponseType `json:"type" codec:"type" bson:"type" yaml:"type" faker:"-"`
	// Uptime the uptime in milliseconds since the agent started
	Uptime int64 `json:"uptime" codec:"uptime" bson:"uptime" yaml:"uptime" faker:"-"`
	// UUID the agent unique identifier
	UUID string `json:"uuid" codec:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
	// Version the agent version
	Version string `json:"version" codec:"version" bson:"version" yaml:"version" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*ProjectResponse)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*ProjectResponse)(nil)

func toProjectResponseObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ProjectResponse:
		return v.ToMap()

	case ProjectResponseEventDate:
		return v.ToMap()

	case ProjectResponseLastExportDate:
		return v.ToMap()

	case []ProjectResponseProjects:
		arr := make([]interface{}, 0)
		for _, i := range v {
			arr = append(arr, i.ToMap())
		}
		return arr

	case ProjectResponseType:
		return v.String()

	default:
		return o
	}
}

// String returns a string representation of ProjectResponse
func (o *ProjectResponse) String() string {
	return fmt.Sprintf("agent.ProjectResponse<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *ProjectResponse) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *ProjectResponse) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *ProjectResponse) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *ProjectResponse) GetModelName() datamodel.ModelNameType {
	return ProjectResponseModelName
}

// NewProjectResponseID provides a template for generating an ID field for ProjectResponse
func NewProjectResponseID(customerID string, refType string, refID string) string {
	return hash.Values("ProjectResponse", customerID, refType, refID)
}

func (o *ProjectResponse) setDefaults(frommap bool) {
	if o.Data == nil {
		o.Data = pstrings.Pointer("")
	}
	if o.Error == nil {
		o.Error = pstrings.Pointer("")
	}
	if o.Projects == nil {
		o.Projects = make([]ProjectResponseProjects, 0)
	}

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("ProjectResponse", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *ProjectResponse) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *ProjectResponse) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *ProjectResponse) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *ProjectResponse) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *ProjectResponse) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *ProjectResponse) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *ProjectResponse) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *ProjectResponse) IsEvented() bool {
	return false
}

// GetTopicConfig returns the topic config object
func (o *ProjectResponse) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
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
	if idstr, ok := kv["id"].(string); ok {
		o.ID = idstr
	}
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
func (o *ProjectResponse) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"architecture":     toProjectResponseObject(o.Architecture, false),
		"customer_id":      toProjectResponseObject(o.CustomerID, false),
		"data":             toProjectResponseObject(o.Data, true),
		"distro":           toProjectResponseObject(o.Distro, false),
		"error":            toProjectResponseObject(o.Error, true),
		"event_date":       toProjectResponseObject(o.EventDate, false),
		"free_space":       toProjectResponseObject(o.FreeSpace, false),
		"go_version":       toProjectResponseObject(o.GoVersion, false),
		"hostname":         toProjectResponseObject(o.Hostname, false),
		"id":               toProjectResponseObject(o.ID, false),
		"integration_id":   toProjectResponseObject(o.IntegrationID, false),
		"last_export_date": toProjectResponseObject(o.LastExportDate, false),
		"memory":           toProjectResponseObject(o.Memory, false),
		"message":          toProjectResponseObject(o.Message, false),
		"num_cpu":          toProjectResponseObject(o.NumCPU, false),
		"os":               toProjectResponseObject(o.OS, false),
		"projects":         toProjectResponseObject(o.Projects, false),
		"ref_id":           toProjectResponseObject(o.RefID, false),
		"ref_type":         toProjectResponseObject(o.RefType, false),
		"request_id":       toProjectResponseObject(o.RequestID, false),
		"success":          toProjectResponseObject(o.Success, false),
		"system_id":        toProjectResponseObject(o.SystemID, false),

		"type":     o.Type.String(),
		"uptime":   toProjectResponseObject(o.Uptime, false),
		"uuid":     toProjectResponseObject(o.UUID, false),
		"version":  toProjectResponseObject(o.Version, false),
		"hashcode": toProjectResponseObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *ProjectResponse) FromMap(kv map[string]interface{}) {

	o.ID = ""

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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
				// if coming in as map, convert it back
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
				// if coming in as map, convert it back
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.IntegrationID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["last_export_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastExportDate.FromMap(kv)
		} else if sv, ok := val.(ProjectResponseLastExportDate); ok {
			// struct
			o.LastExportDate = sv
		} else if sp, ok := val.(*ProjectResponseLastExportDate); ok {
			// struct pointer
			o.LastExportDate = *sp
		}
	} else {
		o.LastExportDate.FromMap(map[string]interface{}{})
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
		} else if a, ok := val.(primitive.A); ok {
			for _, ae := range a {
				if av, ok := ae.(ProjectResponseProjects); ok {
					o.Projects = append(o.Projects, av)
				} else if av, ok := ae.(primitive.M); ok {
					var fm ProjectResponseProjects
					fm.FromMap(av)
					o.Projects = append(o.Projects, fm)
				} else {
					b, _ := json.Marshal(ae)
					var av ProjectResponseProjects
					if err := json.Unmarshal(b, &av); err != nil {
						panic("unsupported type for projects field entry: " + reflect.TypeOf(ae).String())
					}
					o.Projects = append(o.Projects, av)
				}
			}
		} else if arr, ok := val.([]interface{}); ok {
			for _, item := range arr {
				if r, ok := item.(ProjectResponseProjects); ok {
					o.Projects = append(o.Projects, r)
				} else if r, ok := item.(map[string]interface{}); ok {
					var fm ProjectResponseProjects
					fm.FromMap(r)
					o.Projects = append(o.Projects, fm)
				} else if r, ok := item.(primitive.M); ok {
					fm := ProjectResponseProjects{}
					fm.FromMap(r)
					o.Projects = append(o.Projects, fm)
				}
			}
		} else {
			arr := reflect.ValueOf(val)
			if arr.Kind() == reflect.Slice {
				for i := 0; i < arr.Len(); i++ {
					item := arr.Index(i)
					if item.CanAddr() {
						v := item.Addr().MethodByName("ToMap")
						if !v.IsNil() {
							m := v.Call([]reflect.Value{})
							var fm ProjectResponseProjects
							fm.FromMap(m[0].Interface().(map[string]interface{}))
							o.Projects = append(o.Projects, fm)
						}
					}
				}
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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

	if val, ok := kv["system_id"].(string); ok {
		o.SystemID = val
	} else {
		if val, ok := kv["system_id"]; ok {
			if val == nil {
				o.SystemID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.SystemID = fmt.Sprintf("%v", val)
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
			case "log", "LOG":
				o.Type = 3
			case "integration", "INTEGRATION":
				o.Type = 4
			case "export", "EXPORT":
				o.Type = 5
			case "project", "PROJECT":
				o.Type = 6
			case "repo", "REPO":
				o.Type = 7
			case "user", "USER":
				o.Type = 8
			case "uninstall", "UNINSTALL":
				o.Type = 9
			case "upgrade", "UPGRADE":
				o.Type = 10
			case "start", "START":
				o.Type = 11
			case "stop", "STOP":
				o.Type = 12
			case "pause", "PAUSE":
				o.Type = 13
			case "resume", "RESUME":
				o.Type = 14
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
			case "log", "LOG":
				o.Type = 3
			case "integration", "INTEGRATION":
				o.Type = 4
			case "export", "EXPORT":
				o.Type = 5
			case "project", "PROJECT":
				o.Type = 6
			case "repo", "REPO":
				o.Type = 7
			case "user", "USER":
				o.Type = 8
			case "uninstall", "UNINSTALL":
				o.Type = 9
			case "upgrade", "UPGRADE":
				o.Type = 10
			case "start", "START":
				o.Type = 11
			case "stop", "STOP":
				o.Type = 12
			case "pause", "PAUSE":
				o.Type = 13
			case "resume", "RESUME":
				o.Type = 14
			}
		}
	}

	if val, ok := kv["uptime"].(int64); ok {
		o.Uptime = val
	} else {
		if val, ok := kv["uptime"]; ok {
			if val == nil {
				o.Uptime = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Uptime = number.ToInt64Any(val)
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
	args = append(args, o.LastExportDate)
	args = append(args, o.Memory)
	args = append(args, o.Message)
	args = append(args, o.NumCPU)
	args = append(args, o.OS)
	args = append(args, o.Projects)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RequestID)
	args = append(args, o.Success)
	args = append(args, o.SystemID)
	args = append(args, o.Type)
	args = append(args, o.Uptime)
	args = append(args, o.UUID)
	args = append(args, o.Version)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
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
