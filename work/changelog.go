// DO NOT EDIT -- generated code

// Package work - the system which contains project work
package work

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/bxcodec/faker"
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
)

const (
	// ChangelogTopic is the default topic name
	ChangelogTopic datamodel.TopicNameType = "work_Changelog_topic"

	// ChangelogModelName is the model name
	ChangelogModelName datamodel.ModelNameType = "work.Changelog"
)

const (
	// ChangelogCreatedDateColumn is the created_date column name
	ChangelogCreatedDateColumn = "CreatedDate"
	// ChangelogCreatedDateColumnEpochColumn is the epoch column property of the CreatedDate name
	ChangelogCreatedDateColumnEpochColumn = "CreatedDate.Epoch"
	// ChangelogCreatedDateColumnOffsetColumn is the offset column property of the CreatedDate name
	ChangelogCreatedDateColumnOffsetColumn = "CreatedDate.Offset"
	// ChangelogCreatedDateColumnRfc3339Column is the rfc3339 column property of the CreatedDate name
	ChangelogCreatedDateColumnRfc3339Column = "CreatedDate.Rfc3339"
	// ChangelogCustomerIDColumn is the customer_id column name
	ChangelogCustomerIDColumn = "CustomerID"
	// ChangelogFieldColumn is the field column name
	ChangelogFieldColumn = "Field"
	// ChangelogFieldTypeColumn is the field_type column name
	ChangelogFieldTypeColumn = "FieldType"
	// ChangelogFromColumn is the from column name
	ChangelogFromColumn = "From"
	// ChangelogFromStringColumn is the from_string column name
	ChangelogFromStringColumn = "FromString"
	// ChangelogIDColumn is the id column name
	ChangelogIDColumn = "ID"
	// ChangelogIssueIDColumn is the issue_id column name
	ChangelogIssueIDColumn = "IssueID"
	// ChangelogOrdinalColumn is the ordinal column name
	ChangelogOrdinalColumn = "Ordinal"
	// ChangelogProjectIDColumn is the project_id column name
	ChangelogProjectIDColumn = "ProjectID"
	// ChangelogRefIDColumn is the ref_id column name
	ChangelogRefIDColumn = "RefID"
	// ChangelogRefTypeColumn is the ref_type column name
	ChangelogRefTypeColumn = "RefType"
	// ChangelogToColumn is the to column name
	ChangelogToColumn = "To"
	// ChangelogToStringColumn is the to_string column name
	ChangelogToStringColumn = "ToString"
	// ChangelogUpdatedAtColumn is the updated_ts column name
	ChangelogUpdatedAtColumn = "UpdatedAt"
	// ChangelogUserIDColumn is the user_id column name
	ChangelogUserIDColumn = "UserID"
)

// ChangelogCreatedDate represents the object structure for created_date
type ChangelogCreatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toChangelogCreatedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ChangelogCreatedDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *ChangelogCreatedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toChangelogCreatedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toChangelogCreatedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toChangelogCreatedDateObject(o.Rfc3339, false),
	}
}

func (o *ChangelogCreatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ChangelogCreatedDate) FromMap(kv map[string]interface{}) {

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

// ChangelogField is the enumeration type for field
type ChangelogField int32

// UnmarshalBSON unmarshals the enum value
func (v ChangelogField) UnmarshalBSON(buf []byte) error {
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

// String returns the string value for Field
func (v ChangelogField) String() string {
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
	// FieldAssigneeRefID is the enumeration value for assignee_ref_id
	ChangelogFieldAssigneeRefID ChangelogField = 0
	// FieldDueDate is the enumeration value for due_date
	ChangelogFieldDueDate ChangelogField = 1
	// FieldEpicID is the enumeration value for epic_id
	ChangelogFieldEpicID ChangelogField = 2
	// FieldIdentifier is the enumeration value for identifier
	ChangelogFieldIdentifier ChangelogField = 3
	// FieldParentID is the enumeration value for parent_id
	ChangelogFieldParentID ChangelogField = 4
	// FieldPriority is the enumeration value for priority
	ChangelogFieldPriority ChangelogField = 5
	// FieldProjectID is the enumeration value for project_id
	ChangelogFieldProjectID ChangelogField = 6
	// FieldReporterRefID is the enumeration value for reporter_ref_id
	ChangelogFieldReporterRefID ChangelogField = 7
	// FieldResolution is the enumeration value for resolution
	ChangelogFieldResolution ChangelogField = 8
	// FieldSprintIds is the enumeration value for sprint_ids
	ChangelogFieldSprintIds ChangelogField = 9
	// FieldStatus is the enumeration value for status
	ChangelogFieldStatus ChangelogField = 10
	// FieldTags is the enumeration value for tags
	ChangelogFieldTags ChangelogField = 11
	// FieldTitle is the enumeration value for title
	ChangelogFieldTitle ChangelogField = 12
	// FieldType is the enumeration value for type
	ChangelogFieldType ChangelogField = 13
)

// Changelog an individual change to an issue
type Changelog struct {
	// CreatedDate the date when this change was created
	CreatedDate ChangelogCreatedDate `json:"created_date" codec:"created_date" bson:"created_date" yaml:"created_date" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Field the field that was changed
	Field ChangelogField `json:"field" codec:"field" bson:"field" yaml:"field" faker:"-"`
	// FieldType type of the field that was changed
	FieldType string `json:"field_type" codec:"field_type" bson:"field_type" yaml:"field_type" faker:"-"`
	// From id of the change from
	From string `json:"from" codec:"from" bson:"from" yaml:"from" faker:"-"`
	// FromString human readable representation of the change from, used mostly for debug
	FromString string `json:"from_string" codec:"from_string" bson:"from_string" yaml:"from_string" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IssueID id of the issue
	IssueID string `json:"issue_id" codec:"issue_id" bson:"issue_id" yaml:"issue_id" faker:"-"`
	// Ordinal so we can order correctly in queries since dates could be equal
	Ordinal int64 `json:"ordinal" codec:"ordinal" bson:"ordinal" yaml:"ordinal" faker:"-"`
	// ProjectID project id of the issue
	ProjectID string `json:"project_id" codec:"project_id" bson:"project_id" yaml:"project_id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// To id of the change to
	To string `json:"to" codec:"to" bson:"to" yaml:"to" faker:"-"`
	// ToString human readable representation name of the change to, used mostly for debug
	ToString string `json:"to_string" codec:"to_string" bson:"to_string" yaml:"to_string" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// UserID id of the user of this change
	UserID string `json:"user_id" codec:"user_id" bson:"user_id" yaml:"user_id" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Changelog)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*Changelog)(nil)

func toChangelogObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *Changelog:
		return v.ToMap()

	case ChangelogCreatedDate:
		return v.ToMap()

	case ChangelogField:
		return v.String()

	default:
		return o
	}
}

// String returns a string representation of Changelog
func (o *Changelog) String() string {
	return fmt.Sprintf("work.Changelog<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Changelog) GetTopicName() datamodel.TopicNameType {
	return ChangelogTopic
}

// GetStreamName returns the name of the stream
func (o *Changelog) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *Changelog) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *Changelog) GetModelName() datamodel.ModelNameType {
	return ChangelogModelName
}

// NewChangelogID provides a template for generating an ID field for Changelog
func NewChangelogID(customerID string, refType string, refID string) string {
	return hash.Values("Changelog", customerID, refType, refID)
}

func (o *Changelog) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Changelog", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Changelog) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Changelog) GetTopicKey() string {
	var i interface{} = o.ProjectID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Changelog) GetTimestamp() time.Time {
	var dt interface{} = o.UpdatedAt
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
	}
	panic("not sure how to handle the date time format for Changelog")
}

// GetRefID returns the RefID for the object
func (o *Changelog) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Changelog) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *Changelog) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Changelog) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Changelog) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *Changelog) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = ChangelogModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Changelog) GetTopicConfig() *datamodel.ModelTopicConfig {
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
		Timestamp:         "updated_ts",
		NumPartitions:     8,
		CleanupPolicy:     datamodel.CleanupPolicy("compact"),
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetCustomerID will return the customer_id
func (o *Changelog) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of Changelog
func (o *Changelog) Clone() datamodel.Model {
	c := new(Changelog)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Changelog) Anon() datamodel.Model {
	c := new(Changelog)
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
func (o *Changelog) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Changelog) UnmarshalJSON(data []byte) error {
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
func (o *Changelog) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Changelog objects are equal
func (o *Changelog) IsEqual(other *Changelog) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Changelog) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"created_date": toChangelogObject(o.CreatedDate, false),
		"customer_id":  toChangelogObject(o.CustomerID, false),
		"field":        toChangelogObject(o.Field, false),
		"field_type":   toChangelogObject(o.FieldType, false),
		"from":         toChangelogObject(o.From, false),
		"from_string":  toChangelogObject(o.FromString, false),
		"id":           toChangelogObject(o.ID, false),
		"issue_id":     toChangelogObject(o.IssueID, false),
		"ordinal":      toChangelogObject(o.Ordinal, false),
		"project_id":   toChangelogObject(o.ProjectID, false),
		"ref_id":       toChangelogObject(o.RefID, false),
		"ref_type":     toChangelogObject(o.RefType, false),
		"to":           toChangelogObject(o.To, false),
		"to_string":    toChangelogObject(o.ToString, false),
		"updated_ts":   toChangelogObject(o.UpdatedAt, false),
		"user_id":      toChangelogObject(o.UserID, false),
		"hashcode":     toChangelogObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *Changelog) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["created_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.CreatedDate.FromMap(kv)
		} else if sv, ok := val.(ChangelogCreatedDate); ok {
			// struct
			o.CreatedDate = sv
		} else if sp, ok := val.(*ChangelogCreatedDate); ok {
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

	if val, ok := kv["field"].(ChangelogField); ok {
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

	if val, ok := kv["field_type"].(string); ok {
		o.FieldType = val
	} else {
		if val, ok := kv["field_type"]; ok {
			if val == nil {
				o.FieldType = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.FieldType = fmt.Sprintf("%v", val)
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

// Hash will return a hashcode for the object
func (o *Changelog) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CreatedDate)
	args = append(args, o.CustomerID)
	args = append(args, o.Field)
	args = append(args, o.FieldType)
	args = append(args, o.From)
	args = append(args, o.FromString)
	args = append(args, o.ID)
	args = append(args, o.IssueID)
	args = append(args, o.Ordinal)
	args = append(args, o.ProjectID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.To)
	args = append(args, o.ToString)
	args = append(args, o.UpdatedAt)
	args = append(args, o.UserID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *Changelog) GetEventAPIConfig() datamodel.EventAPIConfig {
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
