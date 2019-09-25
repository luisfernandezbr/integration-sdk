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
	"github.com/pinpt/go-common/slice"
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

const (
	// IssueAssigneeRefIDColumn is the assignee_ref_id column name
	IssueAssigneeRefIDColumn = "AssigneeRefID"
	// IssueCreatedDateColumn is the created_date column name
	IssueCreatedDateColumn = "CreatedDate"
	// IssueCreatedDateColumnEpochColumn is the epoch column property of the CreatedDate name
	IssueCreatedDateColumnEpochColumn = "CreatedDate.Epoch"
	// IssueCreatedDateColumnOffsetColumn is the offset column property of the CreatedDate name
	IssueCreatedDateColumnOffsetColumn = "CreatedDate.Offset"
	// IssueCreatedDateColumnRfc3339Column is the rfc3339 column property of the CreatedDate name
	IssueCreatedDateColumnRfc3339Column = "CreatedDate.Rfc3339"
	// IssueCreatorRefIDColumn is the creator_ref_id column name
	IssueCreatorRefIDColumn = "CreatorRefID"
	// IssueCustomFieldsColumn is the customFields column name
	IssueCustomFieldsColumn = "CustomFields"
	// IssueCustomFieldsColumnIDColumn is the id column property of the CustomFields name
	IssueCustomFieldsColumnIDColumn = "CustomFields.ID"
	// IssueCustomFieldsColumnNameColumn is the name column property of the CustomFields name
	IssueCustomFieldsColumnNameColumn = "CustomFields.Name"
	// IssueCustomFieldsColumnValueColumn is the value column property of the CustomFields name
	IssueCustomFieldsColumnValueColumn = "CustomFields.Value"
	// IssueCustomerIDColumn is the customer_id column name
	IssueCustomerIDColumn = "CustomerID"
	// IssueDueDateColumn is the due_date column name
	IssueDueDateColumn = "DueDate"
	// IssueDueDateColumnEpochColumn is the epoch column property of the DueDate name
	IssueDueDateColumnEpochColumn = "DueDate.Epoch"
	// IssueDueDateColumnOffsetColumn is the offset column property of the DueDate name
	IssueDueDateColumnOffsetColumn = "DueDate.Offset"
	// IssueDueDateColumnRfc3339Column is the rfc3339 column property of the DueDate name
	IssueDueDateColumnRfc3339Column = "DueDate.Rfc3339"
	// IssueIDColumn is the id column name
	IssueIDColumn = "ID"
	// IssueIdentifierColumn is the identifier column name
	IssueIdentifierColumn = "Identifier"
	// IssueParentIDColumn is the parent_id column name
	IssueParentIDColumn = "ParentID"
	// IssuePriorityColumn is the priority column name
	IssuePriorityColumn = "Priority"
	// IssueProjectIDColumn is the project_id column name
	IssueProjectIDColumn = "ProjectID"
	// IssueRefIDColumn is the ref_id column name
	IssueRefIDColumn = "RefID"
	// IssueRefTypeColumn is the ref_type column name
	IssueRefTypeColumn = "RefType"
	// IssueReporterRefIDColumn is the reporter_ref_id column name
	IssueReporterRefIDColumn = "ReporterRefID"
	// IssueResolutionColumn is the resolution column name
	IssueResolutionColumn = "Resolution"
	// IssueStatusColumn is the status column name
	IssueStatusColumn = "Status"
	// IssueTagsColumn is the tags column name
	IssueTagsColumn = "Tags"
	// IssueTitleColumn is the title column name
	IssueTitleColumn = "Title"
	// IssueTypeColumn is the type column name
	IssueTypeColumn = "Type"
	// IssueUpdatedDateColumn is the updated_date column name
	IssueUpdatedDateColumn = "UpdatedDate"
	// IssueUpdatedDateColumnEpochColumn is the epoch column property of the UpdatedDate name
	IssueUpdatedDateColumnEpochColumn = "UpdatedDate.Epoch"
	// IssueUpdatedDateColumnOffsetColumn is the offset column property of the UpdatedDate name
	IssueUpdatedDateColumnOffsetColumn = "UpdatedDate.Offset"
	// IssueUpdatedDateColumnRfc3339Column is the rfc3339 column property of the UpdatedDate name
	IssueUpdatedDateColumnRfc3339Column = "UpdatedDate.Rfc3339"
	// IssueUpdatedAtColumn is the updated_ts column name
	IssueUpdatedAtColumn = "UpdatedAt"
	// IssueURLColumn is the url column name
	IssueURLColumn = "URL"
)

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

// IssueCustomFields represents the object structure for customFields
type IssueCustomFields struct {
	// ID the id of the custom field
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// Name the name of the custom field
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
	// Value the value of the custom field
	Value string `json:"value" codec:"value" bson:"value" yaml:"value" faker:"-"`
}

func toIssueCustomFieldsObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IssueCustomFields:
		return v.ToMap()

	default:
		return o
	}
}

func (o *IssueCustomFields) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// ID the id of the custom field
		"id": toIssueCustomFieldsObject(o.ID, false),
		// Name the name of the custom field
		"name": toIssueCustomFieldsObject(o.Name, false),
		// Value the value of the custom field
		"value": toIssueCustomFieldsObject(o.Value, false),
	}
}

func (o *IssueCustomFields) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IssueCustomFields) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
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

	if val, ok := kv["value"].(string); ok {
		o.Value = val
	} else {
		if val, ok := kv["value"]; ok {
			if val == nil {
				o.Value = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Value = fmt.Sprintf("%v", val)
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
	// CreatedDate the date that the issue was created
	CreatedDate IssueCreatedDate `json:"created_date" codec:"created_date" bson:"created_date" yaml:"created_date" faker:"-"`
	// CreatorRefID user id of the creator
	CreatorRefID string `json:"creator_ref_id" codec:"creator_ref_id" bson:"creator_ref_id" yaml:"creator_ref_id" faker:"-"`
	// CustomFields list of custom fields and their values
	CustomFields []IssueCustomFields `json:"customFields" codec:"customFields" bson:"customFields" yaml:"customFields" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// DueDate due date of the issue
	DueDate IssueDueDate `json:"due_date" codec:"due_date" bson:"due_date" yaml:"due_date" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// Identifier the common identifier for the issue
	Identifier string `json:"identifier" codec:"identifier" bson:"identifier" yaml:"identifier" faker:"issue_id"`
	// ParentID parent issue id, if any
	ParentID string `json:"parent_id" codec:"parent_id" bson:"parent_id" yaml:"parent_id" faker:"-"`
	// Priority priority of the issue
	Priority string `json:"priority" codec:"priority" bson:"priority" yaml:"priority" faker:"-"`
	// ProjectID unique project id
	ProjectID string `json:"project_id" codec:"project_id" bson:"project_id" yaml:"project_id" faker:"-"`
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
	// URL the url to the issue page
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

	case IssueCreatedDate:
		return v.ToMap()

	case []IssueCustomFields:
		arr := make([]interface{}, 0)
		for _, i := range v {
			arr = append(arr, i.ToMap())
		}
		return arr

	case IssueDueDate:
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
	return IssueTable.String()
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
	if o.CustomFields == nil {
		o.CustomFields = make([]IssueCustomFields, 0)
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
		"assignee_ref_id": toIssueObject(o.AssigneeRefID, false),
		"created_date":    toIssueObject(o.CreatedDate, false),
		"creator_ref_id":  toIssueObject(o.CreatorRefID, false),
		"customFields":    toIssueObject(o.CustomFields, false),
		"customer_id":     toIssueObject(o.CustomerID, false),
		"due_date":        toIssueObject(o.DueDate, false),
		"id":              toIssueObject(o.ID, false),
		"identifier":      toIssueObject(o.Identifier, false),
		"parent_id":       toIssueObject(o.ParentID, false),
		"priority":        toIssueObject(o.Priority, false),
		"project_id":      toIssueObject(o.ProjectID, false),
		"ref_id":          toIssueObject(o.RefID, false),
		"ref_type":        toIssueObject(o.RefType, false),
		"reporter_ref_id": toIssueObject(o.ReporterRefID, false),
		"resolution":      toIssueObject(o.Resolution, false),
		"status":          toIssueObject(o.Status, false),
		"tags":            toIssueObject(o.Tags, false),
		"title":           toIssueObject(o.Title, false),
		"type":            toIssueObject(o.Type, false),
		"updated_date":    toIssueObject(o.UpdatedDate, false),
		"updated_ts":      toIssueObject(o.UpdatedAt, false),
		"url":             toIssueObject(o.URL, false),
		"hashcode":        toIssueObject(o.Hashcode, false),
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

	if o == nil {

		o.CustomFields = make([]IssueCustomFields, 0)

	}
	if val, ok := kv["customFields"]; ok {
		if sv, ok := val.([]IssueCustomFields); ok {
			o.CustomFields = sv
		} else if sp, ok := val.([]*IssueCustomFields); ok {
			o.CustomFields = o.CustomFields[:0]
			for _, e := range sp {
				o.CustomFields = append(o.CustomFields, *e)
			}
		} else if a, ok := val.(primitive.A); ok {
			for _, ae := range a {
				if av, ok := ae.(IssueCustomFields); ok {
					o.CustomFields = append(o.CustomFields, av)
				} else if av, ok := ae.(primitive.M); ok {
					var fm IssueCustomFields
					fm.FromMap(av)
					o.CustomFields = append(o.CustomFields, fm)
				} else {
					b, _ := json.Marshal(ae)
					var av IssueCustomFields
					if err := json.Unmarshal(b, &av); err != nil {
						panic("unsupported type for customFields field entry: " + reflect.TypeOf(ae).String())
					}
					o.CustomFields = append(o.CustomFields, av)
				}
			}
		} else if arr, ok := val.([]interface{}); ok {
			for _, item := range arr {
				if r, ok := item.(IssueCustomFields); ok {
					o.CustomFields = append(o.CustomFields, r)
				} else if r, ok := item.(map[string]interface{}); ok {
					var fm IssueCustomFields
					fm.FromMap(r)
					o.CustomFields = append(o.CustomFields, fm)
				} else if r, ok := item.(primitive.M); ok {
					fm := IssueCustomFields{}
					fm.FromMap(r)
					o.CustomFields = append(o.CustomFields, fm)
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
							var fm IssueCustomFields
							fm.FromMap(m[0].Interface().(map[string]interface{}))
							o.CustomFields = append(o.CustomFields, fm)
						}
					}
				}
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
	args = append(args, o.CreatedDate)
	args = append(args, o.CreatorRefID)
	args = append(args, o.CustomFields)
	args = append(args, o.CustomerID)
	args = append(args, o.DueDate)
	args = append(args, o.ID)
	args = append(args, o.Identifier)
	args = append(args, o.ParentID)
	args = append(args, o.Priority)
	args = append(args, o.ProjectID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.ReporterRefID)
	args = append(args, o.Resolution)
	args = append(args, o.Status)
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
