// DO NOT EDIT -- generated code

// Package work - the system which contains project work
package work

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/bxcodec/faker"
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
	pnumber "github.com/pinpt/go-common/number"
	"github.com/pinpt/go-common/slice"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// IssueTopic is the default topic name
	IssueTopic datamodel.TopicNameType = "work_Issue_topic"

	// IssueTable is the default table name
	IssueTable datamodel.ModelNameType = "work_issue"

	// IssueModelName is the model name
	IssueModelName datamodel.ModelNameType = "work.Issue"
)

// IssueChangeLogCreatedDate represents the object structure for created_date
type IssueChangeLogCreatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIssueChangeLogCreatedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IssueChangeLogCreatedDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *IssueChangeLogCreatedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIssueChangeLogCreatedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIssueChangeLogCreatedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIssueChangeLogCreatedDateObject(o.Rfc3339, false),
	}
}

func (o *IssueChangeLogCreatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IssueChangeLogCreatedDate) FromMap(kv map[string]interface{}) {

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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// IssueChangeLogField is the enumeration type for field
type IssueChangeLogField int32

// UnmarshalBSONValue for unmarshaling value
func (v *IssueChangeLogField) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = IssueChangeLogField(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "ASSIGNEE_REF_ID":
			*v = IssueChangeLogField(0)
		case "DUE_DATE":
			*v = IssueChangeLogField(1)
		case "EPIC_ID":
			*v = IssueChangeLogField(2)
		case "IDENTIFIER":
			*v = IssueChangeLogField(3)
		case "PARENT_ID":
			*v = IssueChangeLogField(4)
		case "PRIORITY":
			*v = IssueChangeLogField(5)
		case "PROJECT_ID":
			*v = IssueChangeLogField(6)
		case "REPORTER_REF_ID":
			*v = IssueChangeLogField(7)
		case "RESOLUTION":
			*v = IssueChangeLogField(8)
		case "SPRINT_IDS":
			*v = IssueChangeLogField(9)
		case "STATUS":
			*v = IssueChangeLogField(10)
		case "TAGS":
			*v = IssueChangeLogField(11)
		case "TITLE":
			*v = IssueChangeLogField(12)
		case "TYPE":
			*v = IssueChangeLogField(13)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v IssueChangeLogField) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "ASSIGNEE_REF_ID":
		v = 0
	case "DUE_DATE":
		v = 1
	case "EPIC_ID":
		v = 2
	case "IDENTIFIER":
		v = 3
	case "PARENT_ID":
		v = 4
	case "PRIORITY":
		v = 5
	case "PROJECT_ID":
		v = 6
	case "REPORTER_REF_ID":
		v = 7
	case "RESOLUTION":
		v = 8
	case "SPRINT_IDS":
		v = 9
	case "STATUS":
		v = 10
	case "TAGS":
		v = 11
	case "TITLE":
		v = 12
	case "TYPE":
		v = 13
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v IssueChangeLogField) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("ASSIGNEE_REF_ID")
	case 1:
		return json.Marshal("DUE_DATE")
	case 2:
		return json.Marshal("EPIC_ID")
	case 3:
		return json.Marshal("IDENTIFIER")
	case 4:
		return json.Marshal("PARENT_ID")
	case 5:
		return json.Marshal("PRIORITY")
	case 6:
		return json.Marshal("PROJECT_ID")
	case 7:
		return json.Marshal("REPORTER_REF_ID")
	case 8:
		return json.Marshal("RESOLUTION")
	case 9:
		return json.Marshal("SPRINT_IDS")
	case 10:
		return json.Marshal("STATUS")
	case 11:
		return json.Marshal("TAGS")
	case 12:
		return json.Marshal("TITLE")
	case 13:
		return json.Marshal("TYPE")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for ChangeLogField
func (v IssueChangeLogField) String() string {
	switch int32(v) {
	case 0:
		return "ASSIGNEE_REF_ID"
	case 1:
		return "DUE_DATE"
	case 2:
		return "EPIC_ID"
	case 3:
		return "IDENTIFIER"
	case 4:
		return "PARENT_ID"
	case 5:
		return "PRIORITY"
	case 6:
		return "PROJECT_ID"
	case 7:
		return "REPORTER_REF_ID"
	case 8:
		return "RESOLUTION"
	case 9:
		return "SPRINT_IDS"
	case 10:
		return "STATUS"
	case 11:
		return "TAGS"
	case 12:
		return "TITLE"
	case 13:
		return "TYPE"
	}
	return "unset"
}

const (
	// ChangeLogFieldAssigneeRefID is the enumeration value for assignee_ref_id
	IssueChangeLogFieldAssigneeRefID IssueChangeLogField = 0
	// ChangeLogFieldDueDate is the enumeration value for due_date
	IssueChangeLogFieldDueDate IssueChangeLogField = 1
	// ChangeLogFieldEpicID is the enumeration value for epic_id
	IssueChangeLogFieldEpicID IssueChangeLogField = 2
	// ChangeLogFieldIdentifier is the enumeration value for identifier
	IssueChangeLogFieldIdentifier IssueChangeLogField = 3
	// ChangeLogFieldParentID is the enumeration value for parent_id
	IssueChangeLogFieldParentID IssueChangeLogField = 4
	// ChangeLogFieldPriority is the enumeration value for priority
	IssueChangeLogFieldPriority IssueChangeLogField = 5
	// ChangeLogFieldProjectID is the enumeration value for project_id
	IssueChangeLogFieldProjectID IssueChangeLogField = 6
	// ChangeLogFieldReporterRefID is the enumeration value for reporter_ref_id
	IssueChangeLogFieldReporterRefID IssueChangeLogField = 7
	// ChangeLogFieldResolution is the enumeration value for resolution
	IssueChangeLogFieldResolution IssueChangeLogField = 8
	// ChangeLogFieldSprintIds is the enumeration value for sprint_ids
	IssueChangeLogFieldSprintIds IssueChangeLogField = 9
	// ChangeLogFieldStatus is the enumeration value for status
	IssueChangeLogFieldStatus IssueChangeLogField = 10
	// ChangeLogFieldTags is the enumeration value for tags
	IssueChangeLogFieldTags IssueChangeLogField = 11
	// ChangeLogFieldTitle is the enumeration value for title
	IssueChangeLogFieldTitle IssueChangeLogField = 12
	// ChangeLogFieldType is the enumeration value for type
	IssueChangeLogFieldType IssueChangeLogField = 13
)

// IssueChangeLog represents the object structure for change_log
type IssueChangeLog struct {
	// CreatedDate the date when this change was created
	CreatedDate IssueChangeLogCreatedDate `json:"created_date" codec:"created_date" bson:"created_date" yaml:"created_date" faker:"-"`
	// Field the field that was changed
	Field IssueChangeLogField `json:"field" codec:"field" bson:"field" yaml:"field" faker:"-"`
	// From id of the change from
	From string `json:"from" codec:"from" bson:"from" yaml:"from" faker:"-"`
	// FromString human readable representation of the change from, used mostly for debug
	FromString string `json:"from_string" codec:"from_string" bson:"from_string" yaml:"from_string" faker:"-"`
	// Ordinal so we can order correctly in queries since dates could be equal
	Ordinal int64 `json:"ordinal" codec:"ordinal" bson:"ordinal" yaml:"ordinal" faker:"-"`
	// RefID id of the changelog in the source system, for debug purposes
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// To id of the change to
	To string `json:"to" codec:"to" bson:"to" yaml:"to" faker:"-"`
	// ToString human readable representation name of the change to, used mostly for debug
	ToString string `json:"to_string" codec:"to_string" bson:"to_string" yaml:"to_string" faker:"-"`
	// UserID id of the user of this change
	UserID string `json:"user_id" codec:"user_id" bson:"user_id" yaml:"user_id" faker:"-"`
}

func toIssueChangeLogObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IssueChangeLog:
		return v.ToMap()

	case IssueChangeLogCreatedDate:
		return v.ToMap()

	case IssueChangeLogField:
		return v.String()

	default:
		return o
	}
}

func (o *IssueChangeLog) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// CreatedDate the date when this change was created
		"created_date": toIssueChangeLogObject(o.CreatedDate, false),
		// Field the field that was changed
		"field": toIssueChangeLogObject(o.Field, false),
		// From id of the change from
		"from": toIssueChangeLogObject(o.From, false),
		// FromString human readable representation of the change from, used mostly for debug
		"from_string": toIssueChangeLogObject(o.FromString, false),
		// Ordinal so we can order correctly in queries since dates could be equal
		"ordinal": toIssueChangeLogObject(o.Ordinal, false),
		// RefID id of the changelog in the source system, for debug purposes
		"ref_id": toIssueChangeLogObject(o.RefID, false),
		// To id of the change to
		"to": toIssueChangeLogObject(o.To, false),
		// ToString human readable representation name of the change to, used mostly for debug
		"to_string": toIssueChangeLogObject(o.ToString, false),
		// UserID id of the user of this change
		"user_id": toIssueChangeLogObject(o.UserID, false),
	}
}

func (o *IssueChangeLog) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IssueChangeLog) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["created_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.CreatedDate.FromMap(kv)
		} else if sv, ok := val.(IssueChangeLogCreatedDate); ok {
			// struct
			o.CreatedDate = sv
		} else if sp, ok := val.(*IssueChangeLogCreatedDate); ok {
			// struct pointer
			o.CreatedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.CreatedDate.Epoch = dt.Epoch
			o.CreatedDate.Rfc3339 = dt.Rfc3339
			o.CreatedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.CreatedDate.Epoch = dt.Epoch
			o.CreatedDate.Rfc3339 = dt.Rfc3339
			o.CreatedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.CreatedDate.Epoch = dt.Epoch
				o.CreatedDate.Rfc3339 = dt.Rfc3339
				o.CreatedDate.Offset = dt.Offset
			}
		}
	} else {
		o.CreatedDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["field"].(IssueChangeLogField); ok {
		o.Field = val
	} else {
		if em, ok := kv["field"].(map[string]interface{}); ok {
			ev := em["work.field"].(string)
			switch ev {
			case "assignee_ref_id", "ASSIGNEE_REF_ID":
				o.Field = 0
			case "due_date", "DUE_DATE":
				o.Field = 1
			case "epic_id", "EPIC_ID":
				o.Field = 2
			case "identifier", "IDENTIFIER":
				o.Field = 3
			case "parent_id", "PARENT_ID":
				o.Field = 4
			case "priority", "PRIORITY":
				o.Field = 5
			case "project_id", "PROJECT_ID":
				o.Field = 6
			case "reporter_ref_id", "REPORTER_REF_ID":
				o.Field = 7
			case "resolution", "RESOLUTION":
				o.Field = 8
			case "sprint_ids", "SPRINT_IDS":
				o.Field = 9
			case "status", "STATUS":
				o.Field = 10
			case "tags", "TAGS":
				o.Field = 11
			case "title", "TITLE":
				o.Field = 12
			case "type", "TYPE":
				o.Field = 13
			}
		}
		if em, ok := kv["field"].(string); ok {
			switch em {
			case "assignee_ref_id", "ASSIGNEE_REF_ID":
				o.Field = 0
			case "due_date", "DUE_DATE":
				o.Field = 1
			case "epic_id", "EPIC_ID":
				o.Field = 2
			case "identifier", "IDENTIFIER":
				o.Field = 3
			case "parent_id", "PARENT_ID":
				o.Field = 4
			case "priority", "PRIORITY":
				o.Field = 5
			case "project_id", "PROJECT_ID":
				o.Field = 6
			case "reporter_ref_id", "REPORTER_REF_ID":
				o.Field = 7
			case "resolution", "RESOLUTION":
				o.Field = 8
			case "sprint_ids", "SPRINT_IDS":
				o.Field = 9
			case "status", "STATUS":
				o.Field = 10
			case "tags", "TAGS":
				o.Field = 11
			case "title", "TITLE":
				o.Field = 12
			case "type", "TYPE":
				o.Field = 13
			}
		}
	}

	if val, ok := kv["from"].(string); ok {
		o.From = val
	} else {
		if val, ok := kv["from"]; ok {
			if val == nil {
				o.From = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.From = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["from_string"].(string); ok {
		o.FromString = val
	} else {
		if val, ok := kv["from_string"]; ok {
			if val == nil {
				o.FromString = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.FromString = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["ordinal"].(int64); ok {
		o.Ordinal = val
	} else {
		if val, ok := kv["ordinal"]; ok {
			if val == nil {
				o.Ordinal = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Ordinal = number.ToInt64Any(val)
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

	if val, ok := kv["to"].(string); ok {
		o.To = val
	} else {
		if val, ok := kv["to"]; ok {
			if val == nil {
				o.To = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.To = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["to_string"].(string); ok {
		o.ToString = val
	} else {
		if val, ok := kv["to_string"]; ok {
			if val == nil {
				o.ToString = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.ToString = fmt.Sprintf("%v", val)
			}
		}
	}

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
	o.setDefaults(false)
}

// IssueCreatedDate represents the object structure for created_date
type IssueCreatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIssueCreatedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IssueCreatedDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *IssueCreatedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIssueCreatedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIssueCreatedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIssueCreatedDateObject(o.Rfc3339, false),
	}
}

func (o *IssueCreatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IssueCreatedDate) FromMap(kv map[string]interface{}) {

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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// IssueDueDate represents the object structure for due_date
type IssueDueDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIssueDueDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IssueDueDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *IssueDueDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIssueDueDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIssueDueDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIssueDueDateObject(o.Rfc3339, false),
	}
}

func (o *IssueDueDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IssueDueDate) FromMap(kv map[string]interface{}) {

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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// IssuePlannedEndDate represents the object structure for planned_end_date
type IssuePlannedEndDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIssuePlannedEndDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IssuePlannedEndDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *IssuePlannedEndDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIssuePlannedEndDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIssuePlannedEndDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIssuePlannedEndDateObject(o.Rfc3339, false),
	}
}

func (o *IssuePlannedEndDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IssuePlannedEndDate) FromMap(kv map[string]interface{}) {

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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// IssuePlannedStartDate represents the object structure for planned_start_date
type IssuePlannedStartDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIssuePlannedStartDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IssuePlannedStartDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *IssuePlannedStartDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIssuePlannedStartDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIssuePlannedStartDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIssuePlannedStartDateObject(o.Rfc3339, false),
	}
}

func (o *IssuePlannedStartDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IssuePlannedStartDate) FromMap(kv map[string]interface{}) {

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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// IssueUpdatedDate represents the object structure for updated_date
type IssueUpdatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIssueUpdatedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IssueUpdatedDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *IssueUpdatedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIssueUpdatedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIssueUpdatedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIssueUpdatedDateObject(o.Rfc3339, false),
	}
}

func (o *IssueUpdatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IssueUpdatedDate) FromMap(kv map[string]interface{}) {

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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// Issue the issue is a specific work item for a project
type Issue struct {
	// AssigneeRefID user id of the assignee
	AssigneeRefID string `json:"assignee_ref_id" codec:"assignee_ref_id" bson:"assignee_ref_id" yaml:"assignee_ref_id" faker:"-"`
	// ChangeLog the change history of this issue
	ChangeLog []IssueChangeLog `json:"change_log" codec:"change_log" bson:"change_log" yaml:"change_log" faker:"-"`
	// CreatedDate the date that the issue was created
	CreatedDate IssueCreatedDate `json:"created_date" codec:"created_date" bson:"created_date" yaml:"created_date" faker:"-"`
	// CreatorRefID user id of the creator
	CreatorRefID string `json:"creator_ref_id" codec:"creator_ref_id" bson:"creator_ref_id" yaml:"creator_ref_id" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Description the issue description
	Description string `json:"description" codec:"description" bson:"description" yaml:"description" faker:"-"`
	// DueDate due date of the issue
	DueDate IssueDueDate `json:"due_date" codec:"due_date" bson:"due_date" yaml:"due_date" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// Identifier the common identifier for the issue
	Identifier string `json:"identifier" codec:"identifier" bson:"identifier" yaml:"identifier" faker:"issue_id"`
	// ParentID parent issue id, if any
	ParentID string `json:"parent_id" codec:"parent_id" bson:"parent_id" yaml:"parent_id" faker:"-"`
	// PlannedEndDate the date that the issue was planned to end
	PlannedEndDate IssuePlannedEndDate `json:"planned_end_date" codec:"planned_end_date" bson:"planned_end_date" yaml:"planned_end_date" faker:"-"`
	// PlannedStartDate the date that the issue was planned to start
	PlannedStartDate IssuePlannedStartDate `json:"planned_start_date" codec:"planned_start_date" bson:"planned_start_date" yaml:"planned_start_date" faker:"-"`
	// Priority priority of the issue
	Priority string `json:"priority" codec:"priority" bson:"priority" yaml:"priority" faker:"-"`
	// ProjectID unique project id
	ProjectID string `json:"project_id" codec:"project_id" bson:"project_id" yaml:"project_id" faker:"-"`
	// PullRequestIds pull request ids linked to this issue, if available
	PullRequestIds []string `json:"pull_request_ids" codec:"pull_request_ids" bson:"pull_request_ids" yaml:"pull_request_ids" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// ReporterRefID user id of the reporter
	ReporterRefID string `json:"reporter_ref_id" codec:"reporter_ref_id" bson:"reporter_ref_id" yaml:"reporter_ref_id" faker:"-"`
	// Resolution resolution of the issue
	Resolution string `json:"resolution" codec:"resolution" bson:"resolution" yaml:"resolution" faker:"-"`
	// Status status of the issue
	Status string `json:"status" codec:"status" bson:"status" yaml:"status" faker:"-"`
	// StoryPoints the story points estimation for the issue
	StoryPoints *float64 `json:"story_points,omitempty" codec:"story_points,omitempty" bson:"story_points" yaml:"story_points,omitempty" faker:"-"`
	// Tags tags on the issue
	Tags []string `json:"tags" codec:"tags" bson:"tags" yaml:"tags" faker:"-"`
	// Title the issue title
	Title string `json:"title" codec:"title" bson:"title" yaml:"title" faker:"issue_title"`
	// Type type of issue
	Type string `json:"type" codec:"type" bson:"type" yaml:"type" faker:"-"`
	// UpdatedDate the date that the issue was updated
	UpdatedDate IssueUpdatedDate `json:"updated_date" codec:"updated_date" bson:"updated_date" yaml:"updated_date" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// URL the url to the sprint page
	URL string `json:"url" codec:"url" bson:"url" yaml:"url" faker:"url"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Issue)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*Issue)(nil)

func toIssueObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *Issue:
		return v.ToMap()

	case []IssueChangeLog:
		arr := make([]interface{}, 0)
		for _, i := range v {
			arr = append(arr, i.ToMap())
		}
		return arr

	case IssueCreatedDate:
		return v.ToMap()

	case IssueDueDate:
		return v.ToMap()

	case IssuePlannedEndDate:
		return v.ToMap()

	case IssuePlannedStartDate:
		return v.ToMap()

	case IssueUpdatedDate:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of Issue
func (o *Issue) String() string {
	return fmt.Sprintf("work.Issue<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Issue) GetTopicName() datamodel.TopicNameType {
	return IssueTopic
}

// GetStreamName returns the name of the stream
func (o *Issue) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *Issue) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *Issue) GetModelName() datamodel.ModelNameType {
	return IssueModelName
}

// NewIssueID provides a template for generating an ID field for Issue
func NewIssueID(customerID string, refType string, refID string) string {
	return hash.Values("Issue", customerID, refType, refID)
}

func (o *Issue) setDefaults(frommap bool) {
	if o.ChangeLog == nil {
		o.ChangeLog = make([]IssueChangeLog, 0)
	}
	if o.PullRequestIds == nil {
		o.PullRequestIds = make([]string, 0)
	}
	if o.StoryPoints == nil {
		o.StoryPoints = pnumber.Float64Pointer(0)
	}
	if o.Tags == nil {
		o.Tags = make([]string, 0)
	}

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Issue", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Issue) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Issue) GetTopicKey() string {
	var i interface{} = o.ProjectID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Issue) GetTimestamp() time.Time {
	var dt interface{} = o.CreatedDate
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
	case IssueCreatedDate:
		return datetime.DateFromEpoch(v.Epoch)
	}
	panic("not sure how to handle the date time format for Issue")
}

// GetRefID returns the RefID for the object
func (o *Issue) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Issue) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *Issue) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Issue) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Issue) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *Issue) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = IssueModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Issue) GetTopicConfig() *datamodel.ModelTopicConfig {
	retention, err := time.ParseDuration("87360h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 87360h0m0s. " + err.Error())
	}

	ttl, err := time.ParseDuration("0s")
	if err != nil {
		ttl = 0
	}
	if ttl == 0 && retention != 0 {
		ttl = retention // they should be the same if not set
	}
	return &datamodel.ModelTopicConfig{
		Key:               "project_id",
		Timestamp:         "created_date",
		NumPartitions:     8,
		CleanupPolicy:     datamodel.CleanupPolicy("compact"),
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetCustomerID will return the customer_id
func (o *Issue) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of Issue
func (o *Issue) Clone() datamodel.Model {
	c := new(Issue)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Issue) Anon() datamodel.Model {
	c := new(Issue)
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
func (o *Issue) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Issue) UnmarshalJSON(data []byte) error {
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
func (o *Issue) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Issue objects are equal
func (o *Issue) IsEqual(other *Issue) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Issue) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"assignee_ref_id":    toIssueObject(o.AssigneeRefID, false),
		"change_log":         toIssueObject(o.ChangeLog, false),
		"created_date":       toIssueObject(o.CreatedDate, false),
		"creator_ref_id":     toIssueObject(o.CreatorRefID, false),
		"customer_id":        toIssueObject(o.CustomerID, false),
		"description":        toIssueObject(o.Description, false),
		"due_date":           toIssueObject(o.DueDate, false),
		"id":                 toIssueObject(o.ID, false),
		"identifier":         toIssueObject(o.Identifier, false),
		"parent_id":          toIssueObject(o.ParentID, false),
		"planned_end_date":   toIssueObject(o.PlannedEndDate, false),
		"planned_start_date": toIssueObject(o.PlannedStartDate, false),
		"priority":           toIssueObject(o.Priority, false),
		"project_id":         toIssueObject(o.ProjectID, false),
		"pull_request_ids":   toIssueObject(o.PullRequestIds, false),
		"ref_id":             toIssueObject(o.RefID, false),
		"ref_type":           toIssueObject(o.RefType, false),
		"reporter_ref_id":    toIssueObject(o.ReporterRefID, false),
		"resolution":         toIssueObject(o.Resolution, false),
		"status":             toIssueObject(o.Status, false),
		"story_points":       toIssueObject(o.StoryPoints, true),
		"tags":               toIssueObject(o.Tags, false),
		"title":              toIssueObject(o.Title, false),
		"type":               toIssueObject(o.Type, false),
		"updated_date":       toIssueObject(o.UpdatedDate, false),
		"updated_ts":         toIssueObject(o.UpdatedAt, false),
		"url":                toIssueObject(o.URL, false),
		"hashcode":           toIssueObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *Issue) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["assignee_ref_id"].(string); ok {
		o.AssigneeRefID = val
	} else {
		if val, ok := kv["assignee_ref_id"]; ok {
			if val == nil {
				o.AssigneeRefID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.AssigneeRefID = fmt.Sprintf("%v", val)
			}
		}
	}

	if o == nil {

		o.ChangeLog = make([]IssueChangeLog, 0)

	}
	if val, ok := kv["change_log"]; ok {
		if sv, ok := val.([]IssueChangeLog); ok {
			o.ChangeLog = sv
		} else if sp, ok := val.([]*IssueChangeLog); ok {
			o.ChangeLog = o.ChangeLog[:0]
			for _, e := range sp {
				o.ChangeLog = append(o.ChangeLog, *e)
			}
		} else if a, ok := val.(primitive.A); ok {
			for _, ae := range a {
				if av, ok := ae.(IssueChangeLog); ok {
					o.ChangeLog = append(o.ChangeLog, av)
				} else if av, ok := ae.(primitive.M); ok {
					var fm IssueChangeLog
					fm.FromMap(av)
					o.ChangeLog = append(o.ChangeLog, fm)
				} else {
					b, _ := json.Marshal(ae)
					var av IssueChangeLog
					if err := json.Unmarshal(b, &av); err != nil {
						panic("unsupported type for change_log field entry: " + reflect.TypeOf(ae).String())
					}
					o.ChangeLog = append(o.ChangeLog, av)
				}
			}
		} else if arr, ok := val.([]interface{}); ok {
			for _, item := range arr {
				if r, ok := item.(IssueChangeLog); ok {
					o.ChangeLog = append(o.ChangeLog, r)
				} else if r, ok := item.(map[string]interface{}); ok {
					var fm IssueChangeLog
					fm.FromMap(r)
					o.ChangeLog = append(o.ChangeLog, fm)
				} else if r, ok := item.(primitive.M); ok {
					fm := IssueChangeLog{}
					fm.FromMap(r)
					o.ChangeLog = append(o.ChangeLog, fm)
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
							var fm IssueChangeLog
							fm.FromMap(m[0].Interface().(map[string]interface{}))
							o.ChangeLog = append(o.ChangeLog, fm)
						}
					}
				}
			}
		}
	}

	if val, ok := kv["created_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.CreatedDate.FromMap(kv)
		} else if sv, ok := val.(IssueCreatedDate); ok {
			// struct
			o.CreatedDate = sv
		} else if sp, ok := val.(*IssueCreatedDate); ok {
			// struct pointer
			o.CreatedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.CreatedDate.Epoch = dt.Epoch
			o.CreatedDate.Rfc3339 = dt.Rfc3339
			o.CreatedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.CreatedDate.Epoch = dt.Epoch
			o.CreatedDate.Rfc3339 = dt.Rfc3339
			o.CreatedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.CreatedDate.Epoch = dt.Epoch
				o.CreatedDate.Rfc3339 = dt.Rfc3339
				o.CreatedDate.Offset = dt.Offset
			}
		}
	} else {
		o.CreatedDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["creator_ref_id"].(string); ok {
		o.CreatorRefID = val
	} else {
		if val, ok := kv["creator_ref_id"]; ok {
			if val == nil {
				o.CreatorRefID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.CreatorRefID = fmt.Sprintf("%v", val)
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

	if val, ok := kv["description"].(string); ok {
		o.Description = val
	} else {
		if val, ok := kv["description"]; ok {
			if val == nil {
				o.Description = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Description = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["due_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.DueDate.FromMap(kv)
		} else if sv, ok := val.(IssueDueDate); ok {
			// struct
			o.DueDate = sv
		} else if sp, ok := val.(*IssueDueDate); ok {
			// struct pointer
			o.DueDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.DueDate.Epoch = dt.Epoch
			o.DueDate.Rfc3339 = dt.Rfc3339
			o.DueDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.DueDate.Epoch = dt.Epoch
			o.DueDate.Rfc3339 = dt.Rfc3339
			o.DueDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.DueDate.Epoch = dt.Epoch
				o.DueDate.Rfc3339 = dt.Rfc3339
				o.DueDate.Offset = dt.Offset
			}
		}
	} else {
		o.DueDate.FromMap(map[string]interface{}{})
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

	if val, ok := kv["parent_id"].(string); ok {
		o.ParentID = val
	} else {
		if val, ok := kv["parent_id"]; ok {
			if val == nil {
				o.ParentID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.ParentID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["planned_end_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.PlannedEndDate.FromMap(kv)
		} else if sv, ok := val.(IssuePlannedEndDate); ok {
			// struct
			o.PlannedEndDate = sv
		} else if sp, ok := val.(*IssuePlannedEndDate); ok {
			// struct pointer
			o.PlannedEndDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.PlannedEndDate.Epoch = dt.Epoch
			o.PlannedEndDate.Rfc3339 = dt.Rfc3339
			o.PlannedEndDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.PlannedEndDate.Epoch = dt.Epoch
			o.PlannedEndDate.Rfc3339 = dt.Rfc3339
			o.PlannedEndDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.PlannedEndDate.Epoch = dt.Epoch
				o.PlannedEndDate.Rfc3339 = dt.Rfc3339
				o.PlannedEndDate.Offset = dt.Offset
			}
		}
	} else {
		o.PlannedEndDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["planned_start_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.PlannedStartDate.FromMap(kv)
		} else if sv, ok := val.(IssuePlannedStartDate); ok {
			// struct
			o.PlannedStartDate = sv
		} else if sp, ok := val.(*IssuePlannedStartDate); ok {
			// struct pointer
			o.PlannedStartDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.PlannedStartDate.Epoch = dt.Epoch
			o.PlannedStartDate.Rfc3339 = dt.Rfc3339
			o.PlannedStartDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.PlannedStartDate.Epoch = dt.Epoch
			o.PlannedStartDate.Rfc3339 = dt.Rfc3339
			o.PlannedStartDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.PlannedStartDate.Epoch = dt.Epoch
				o.PlannedStartDate.Rfc3339 = dt.Rfc3339
				o.PlannedStartDate.Offset = dt.Offset
			}
		}
	} else {
		o.PlannedStartDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["priority"].(string); ok {
		o.Priority = val
	} else {
		if val, ok := kv["priority"]; ok {
			if val == nil {
				o.Priority = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Priority = fmt.Sprintf("%v", val)
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

	if val, ok := kv["pull_request_ids"]; ok {
		if val != nil {
			na := make([]string, 0)
			if a, ok := val.([]string); ok {
				na = append(na, a...)
			} else {
				if a, ok := val.([]interface{}); ok {
					for _, ae := range a {
						if av, ok := ae.(string); ok {
							na = append(na, av)
						} else {
							if badMap, ok := ae.(map[interface{}]interface{}); ok {
								ae = slice.ConvertToStringToInterface(badMap)
							}
							b, _ := json.Marshal(ae)
							var av string
							if err := json.Unmarshal(b, &av); err != nil {
								panic("unsupported type for pull_request_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else if s, ok := val.(string); ok {
					for _, sv := range strings.Split(s, ",") {
						na = append(na, strings.TrimSpace(sv))
					}
				} else if a, ok := val.(primitive.A); ok {
					for _, ae := range a {
						if av, ok := ae.(string); ok {
							na = append(na, av)
						} else {
							b, _ := json.Marshal(ae)
							var av string
							if err := json.Unmarshal(b, &av); err != nil {
								panic("unsupported type for pull_request_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for pull_request_ids field")
				}
			}
			o.PullRequestIds = na
		}
	}
	if o.PullRequestIds == nil {
		o.PullRequestIds = make([]string, 0)
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

	if val, ok := kv["reporter_ref_id"].(string); ok {
		o.ReporterRefID = val
	} else {
		if val, ok := kv["reporter_ref_id"]; ok {
			if val == nil {
				o.ReporterRefID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.ReporterRefID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["resolution"].(string); ok {
		o.Resolution = val
	} else {
		if val, ok := kv["resolution"]; ok {
			if val == nil {
				o.Resolution = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Resolution = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["status"].(string); ok {
		o.Status = val
	} else {
		if val, ok := kv["status"]; ok {
			if val == nil {
				o.Status = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Status = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["story_points"].(*float64); ok {
		o.StoryPoints = val
	} else if val, ok := kv["story_points"].(float64); ok {
		o.StoryPoints = &val
	} else {
		if val, ok := kv["story_points"]; ok {
			if val == nil {
				o.StoryPoints = number.Float64Pointer(number.ToFloat64Any(nil))
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["float64"]
				}
				o.StoryPoints = number.Float64Pointer(number.ToFloat64Any(val))
			}
		}
	}

	if val, ok := kv["tags"]; ok {
		if val != nil {
			na := make([]string, 0)
			if a, ok := val.([]string); ok {
				na = append(na, a...)
			} else {
				if a, ok := val.([]interface{}); ok {
					for _, ae := range a {
						if av, ok := ae.(string); ok {
							na = append(na, av)
						} else {
							if badMap, ok := ae.(map[interface{}]interface{}); ok {
								ae = slice.ConvertToStringToInterface(badMap)
							}
							b, _ := json.Marshal(ae)
							var av string
							if err := json.Unmarshal(b, &av); err != nil {
								panic("unsupported type for tags field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else if s, ok := val.(string); ok {
					for _, sv := range strings.Split(s, ",") {
						na = append(na, strings.TrimSpace(sv))
					}
				} else if a, ok := val.(primitive.A); ok {
					for _, ae := range a {
						if av, ok := ae.(string); ok {
							na = append(na, av)
						} else {
							b, _ := json.Marshal(ae)
							var av string
							if err := json.Unmarshal(b, &av); err != nil {
								panic("unsupported type for tags field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for tags field")
				}
			}
			o.Tags = na
		}
	}
	if o.Tags == nil {
		o.Tags = make([]string, 0)
	}

	if val, ok := kv["title"].(string); ok {
		o.Title = val
	} else {
		if val, ok := kv["title"]; ok {
			if val == nil {
				o.Title = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Title = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["type"].(string); ok {
		o.Type = val
	} else {
		if val, ok := kv["type"]; ok {
			if val == nil {
				o.Type = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Type = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["updated_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.UpdatedDate.FromMap(kv)
		} else if sv, ok := val.(IssueUpdatedDate); ok {
			// struct
			o.UpdatedDate = sv
		} else if sp, ok := val.(*IssueUpdatedDate); ok {
			// struct pointer
			o.UpdatedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.UpdatedDate.Epoch = dt.Epoch
			o.UpdatedDate.Rfc3339 = dt.Rfc3339
			o.UpdatedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.UpdatedDate.Epoch = dt.Epoch
			o.UpdatedDate.Rfc3339 = dt.Rfc3339
			o.UpdatedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.UpdatedDate.Epoch = dt.Epoch
				o.UpdatedDate.Rfc3339 = dt.Rfc3339
				o.UpdatedDate.Offset = dt.Offset
			}
		}
	} else {
		o.UpdatedDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["updated_ts"].(int64); ok {
		o.UpdatedAt = val
	} else {
		if val, ok := kv["updated_ts"]; ok {
			if val == nil {
				o.UpdatedAt = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.UpdatedAt = number.ToInt64Any(val)
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

// Hash will return a hashcode for the object
func (o *Issue) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.AssigneeRefID)
	args = append(args, o.ChangeLog)
	args = append(args, o.CreatedDate)
	args = append(args, o.CreatorRefID)
	args = append(args, o.CustomerID)
	args = append(args, o.Description)
	args = append(args, o.DueDate)
	args = append(args, o.ID)
	args = append(args, o.Identifier)
	args = append(args, o.ParentID)
	args = append(args, o.PlannedEndDate)
	args = append(args, o.PlannedStartDate)
	args = append(args, o.Priority)
	args = append(args, o.ProjectID)
	args = append(args, o.PullRequestIds)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.ReporterRefID)
	args = append(args, o.Resolution)
	args = append(args, o.Status)
	args = append(args, o.StoryPoints)
	args = append(args, o.Tags)
	args = append(args, o.Title)
	args = append(args, o.Type)
	args = append(args, o.UpdatedDate)
	args = append(args, o.UpdatedAt)
	args = append(args, o.URL)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *Issue) GetEventAPIConfig() datamodel.EventAPIConfig {
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
