// DO NOT EDIT -- generated code

// Package work - the system which contains project work
package work

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
	"strings"
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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// IssueTopic is the default topic name
	IssueTopic datamodel.TopicNameType = "work_Issue_topic"

	// IssueStream is the default stream name
	IssueStream datamodel.TopicNameType = "work_Issue_stream"

	// IssueTable is the default table name
	IssueTable datamodel.TopicNameType = "work_Issue"

	// IssueModelName is the model name
	IssueModelName datamodel.ModelNameType = "work.Issue"
)

const (
	// IssueAssigneeRefIDColumn is the assignee_ref_id column name
	IssueAssigneeRefIDColumn = "assignee_ref_id"
	// IssueCreatedColumn is the created column name
	IssueCreatedColumn = "created"
	// IssueCreatorRefIDColumn is the creator_ref_id column name
	IssueCreatorRefIDColumn = "creator_ref_id"
	// IssueCustomFieldsColumn is the customFields column name
	IssueCustomFieldsColumn = "customFields"
	// IssueCustomerIDColumn is the customer_id column name
	IssueCustomerIDColumn = "customer_id"
	// IssueDueDateColumn is the due_date column name
	IssueDueDateColumn = "due_date"
	// IssueIDColumn is the id column name
	IssueIDColumn = "id"
	// IssueIdentifierColumn is the identifier column name
	IssueIdentifierColumn = "identifier"
	// IssueParentIDColumn is the parent_id column name
	IssueParentIDColumn = "parent_id"
	// IssuePlannedEndAtColumn is the planned_end_ts column name
	IssuePlannedEndAtColumn = "planned_end_ts"
	// IssuePlannedStartAtColumn is the planned_start_ts column name
	IssuePlannedStartAtColumn = "planned_start_ts"
	// IssuePriorityColumn is the priority column name
	IssuePriorityColumn = "priority"
	// IssueProjectIDColumn is the project_id column name
	IssueProjectIDColumn = "project_id"
	// IssueRefIDColumn is the ref_id column name
	IssueRefIDColumn = "ref_id"
	// IssueRefTypeColumn is the ref_type column name
	IssueRefTypeColumn = "ref_type"
	// IssueReporterRefIDColumn is the reporter_ref_id column name
	IssueReporterRefIDColumn = "reporter_ref_id"
	// IssueResolutionColumn is the resolution column name
	IssueResolutionColumn = "resolution"
	// IssueStatusColumn is the status column name
	IssueStatusColumn = "status"
	// IssueTagsColumn is the tags column name
	IssueTagsColumn = "tags"
	// IssueTitleColumn is the title column name
	IssueTitleColumn = "title"
	// IssueTypeColumn is the type column name
	IssueTypeColumn = "type"
	// IssueUpdatedColumn is the updated column name
	IssueUpdatedColumn = "updated"
	// IssueURLColumn is the url column name
	IssueURLColumn = "url"
)

// IssueCustomFields represents the object structure for customFields
type IssueCustomFields struct {
	// ID the id of the custom field
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// Name the name of the custom field
	Name string `json:"name" bson:"name" yaml:"name" faker:"-"`
	// Value the value of the custom field
	Value string `json:"value" bson:"value" yaml:"value" faker:"-"`
}

func (o *IssueCustomFields) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// ID the id of the custom field
		"id": o.ID,
		// Name the name of the custom field
		"name": o.Name,
		// Value the value of the custom field
		"value": o.Value,
	}
}

// Issue the issue is a specific work item for a project
type Issue struct {
	// AssigneeRefID user id of the assignee
	AssigneeRefID string `json:"assignee_ref_id" bson:"assignee_ref_id" yaml:"assignee_ref_id" faker:"-"`
	// Created the date that the issue was created
	Created string `json:"created" bson:"created" yaml:"created" faker:"-"`
	// CreatorRefID user id of the creator
	CreatorRefID string `json:"creator_ref_id" bson:"creator_ref_id" yaml:"creator_ref_id" faker:"-"`
	// CustomFields list of custom fields and their values
	CustomFields []IssueCustomFields `json:"customFields" bson:"customFields" yaml:"customFields" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// DueDate due date of the issue
	DueDate string `json:"due_date" bson:"due_date" yaml:"due_date" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// Identifier the common identifier for the issue
	Identifier string `json:"identifier" bson:"identifier" yaml:"identifier" faker:"issue_id"`
	// ParentID parent issue id, if any
	ParentID string `json:"parent_id" bson:"parent_id" yaml:"parent_id" faker:"-"`
	// PlannedEndAt the timestamp in UTC that the issue was planned to end
	PlannedEndAt int64 `json:"planned_end_ts" bson:"planned_end_ts" yaml:"planned_end_ts" faker:"-"`
	// PlannedStartAt the timestamp in UTC that the issue was planned to start
	PlannedStartAt int64 `json:"planned_start_ts" bson:"planned_start_ts" yaml:"planned_start_ts" faker:"-"`
	// Priority priority of the issue
	Priority string `json:"priority" bson:"priority" yaml:"priority" faker:"-"`
	// ProjectID unique project id
	ProjectID string `json:"project_id" bson:"project_id" yaml:"project_id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// ReporterRefID user id of the reporter
	ReporterRefID string `json:"reporter_ref_id" bson:"reporter_ref_id" yaml:"reporter_ref_id" faker:"-"`
	// Resolution resolution of the issue
	Resolution string `json:"resolution" bson:"resolution" yaml:"resolution" faker:"-"`
	// Status status of the issue
	Status string `json:"status" bson:"status" yaml:"status" faker:"-"`
	// Tags tags on the issue
	Tags []string `json:"tags" bson:"tags" yaml:"tags" faker:"-"`
	// Title the issue title
	Title string `json:"title" bson:"title" yaml:"title" faker:"issue_title"`
	// Type type of issue
	Type string `json:"type" bson:"type" yaml:"type" faker:"-"`
	// Updated the date that the issue was updated
	Updated string `json:"updated" bson:"updated" yaml:"updated" faker:"-"`
	// URL the url to the issue page
	URL string `json:"url" bson:"url" yaml:"url" faker:"url"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Issue)(nil)

func toIssueObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toIssueObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if o == nil {
		return toIssueObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toIssueObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toIssueObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toIssueObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toIssueObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toIssueObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toIssueObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toIssueObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toIssueObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toIssueObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toIssueObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case map[string]interface{}:
		return o
	case *map[string]interface{}:
		return v
	case map[string]string:
		return v
	case *map[string]string:
		return *v
	case *Issue:
		return v.ToMap()
	case Issue:
		return v.ToMap()
	case []string, []int64, []float64, []bool:
		return o
	case *[]string:
		return (*(o.(*[]string)))
	case *[]int64:
		return (*(o.(*[]int64)))
	case *[]float64:
		return (*(o.(*[]float64)))
	case *[]bool:
		return (*(o.(*[]bool)))
	case []interface{}:
		a := o.([]interface{})
		arr := make([]interface{}, 0)
		for _, av := range a {
			arr = append(arr, toIssueObject(av, isavro, false, ""))
		}
		return arr

	case IssueCustomFields:
		vv := o.(IssueCustomFields)
		return vv.ToMap()
	case *IssueCustomFields:
		return (*o.(*IssueCustomFields)).ToMap()
	case []IssueCustomFields:
		arr := make([]interface{}, 0)
		for _, i := range o.([]IssueCustomFields) {
			arr = append(arr, i.ToMap())
		}
		return arr
	case *[]IssueCustomFields:
		arr := make([]interface{}, 0)
		vv := o.(*[]IssueCustomFields)
		for _, i := range *vv {
			arr = append(arr, i.ToMap())
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of Issue
func (o *Issue) String() string {
	return fmt.Sprintf("work.Issue<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Issue) GetTopicName() datamodel.TopicNameType {
	return IssueTopic
}

// GetModelName returns the name of the model
func (o *Issue) GetModelName() datamodel.ModelNameType {
	return IssueModelName
}

func (o *Issue) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Issue) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Issue", o.CustomerID, o.RefType, o.GetRefID())
	}
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
	var dt interface{} = o.Updated
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
	retention, err := time.ParseDuration("168h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 168h0m0s. " + err.Error())
	}

	ttl, err := time.ParseDuration("0s")
	if err != nil {
		ttl = 0
	}
	return &datamodel.ModelTopicConfig{
		Key:               "project_id",
		Timestamp:         "updated",
		NumPartitions:     8,
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetStateKey returns a key for use in state store
func (o *Issue) GetStateKey() string {
	key := "project_id"
	return fmt.Sprintf("%s_%s", key, o.GetID())
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
	return nil
}

var cachedCodecIssue *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *Issue) GetAvroCodec() *goavro.Codec {
	if cachedCodecIssue == nil {
		c, err := GetIssueAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecIssue = c
	}
	return cachedCodecIssue
}

// ToAvroBinary returns the data as Avro binary data
func (o *Issue) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *Issue) FromAvroBinary(value []byte) error {
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
func (o *Issue) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Issue objects are equal
func (o *Issue) IsEqual(other *Issue) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Issue) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
		if o.CustomFields == nil {
			o.CustomFields = make([]IssueCustomFields, 0)
		}
		if o.Tags == nil {
			o.Tags = make([]string, 0)
		}
	}
	o.setDefaults()
	return map[string]interface{}{
		"assignee_ref_id":  toIssueObject(o.AssigneeRefID, isavro, false, "string"),
		"created":          toIssueObject(o.Created, isavro, false, "string"),
		"creator_ref_id":   toIssueObject(o.CreatorRefID, isavro, false, "string"),
		"customFields":     toIssueObject(o.CustomFields, isavro, false, "customFields"),
		"customer_id":      toIssueObject(o.CustomerID, isavro, false, "string"),
		"due_date":         toIssueObject(o.DueDate, isavro, false, "string"),
		"id":               toIssueObject(o.ID, isavro, false, "string"),
		"identifier":       toIssueObject(o.Identifier, isavro, false, "string"),
		"parent_id":        toIssueObject(o.ParentID, isavro, false, "string"),
		"planned_end_ts":   toIssueObject(o.PlannedEndAt, isavro, false, "long"),
		"planned_start_ts": toIssueObject(o.PlannedStartAt, isavro, false, "long"),
		"priority":         toIssueObject(o.Priority, isavro, false, "string"),
		"project_id":       toIssueObject(o.ProjectID, isavro, false, "string"),
		"ref_id":           toIssueObject(o.RefID, isavro, false, "string"),
		"ref_type":         toIssueObject(o.RefType, isavro, false, "string"),
		"reporter_ref_id":  toIssueObject(o.ReporterRefID, isavro, false, "string"),
		"resolution":       toIssueObject(o.Resolution, isavro, false, "string"),
		"status":           toIssueObject(o.Status, isavro, false, "string"),
		"tags":             toIssueObject(o.Tags, isavro, false, "tags"),
		"title":            toIssueObject(o.Title, isavro, false, "string"),
		"type":             toIssueObject(o.Type, isavro, false, "string"),
		"updated":          toIssueObject(o.Updated, isavro, false, "string"),
		"url":              toIssueObject(o.URL, isavro, false, "string"),
		"hashcode":         toIssueObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *Issue) FromMap(kv map[string]interface{}) {
	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["assignee_ref_id"].(string); ok {
		o.AssigneeRefID = val
	} else {
		val := kv["assignee_ref_id"]
		if val == nil {
			o.AssigneeRefID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.AssigneeRefID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["created"].(string); ok {
		o.Created = val
	} else {
		val := kv["created"]
		if val == nil {
			o.Created = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Created = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["creator_ref_id"].(string); ok {
		o.CreatorRefID = val
	} else {
		val := kv["creator_ref_id"]
		if val == nil {
			o.CreatorRefID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.CreatorRefID = fmt.Sprintf("%v", val)
		}
	}
	if val := kv["customFields"]; val != nil {
		na := make([]IssueCustomFields, 0)
		if a, ok := val.([]IssueCustomFields); ok {
			na = append(na, a...)
		} else {
			if a, ok := val.([]interface{}); ok {
				for _, ae := range a {
					if av, ok := ae.(IssueCustomFields); ok {
						na = append(na, av)
					} else {
						b, _ := json.Marshal(ae)
						var av IssueCustomFields
						if err := json.Unmarshal(b, &av); err != nil {
							panic("unsupported type for customFields field entry: " + reflect.TypeOf(ae).String())
						}
						na = append(na, av)
					}
				}
			} else if a, ok := val.(primitive.A); ok {
				for _, ae := range a {
					if av, ok := ae.(IssueCustomFields); ok {
						na = append(na, av)
					} else {
						b, _ := json.Marshal(ae)
						var av IssueCustomFields
						if err := json.Unmarshal(b, &av); err != nil {
							panic("unsupported type for customFields field entry: " + reflect.TypeOf(ae).String())
						}
						na = append(na, av)
					}
				}
			} else {
				fmt.Println(reflect.TypeOf(val).String())
				panic("unsupported type for customFields field")
			}
		}
		o.CustomFields = na
	} else {
		o.CustomFields = []IssueCustomFields{}
	}
	if o.CustomFields == nil {
		o.CustomFields = make([]IssueCustomFields, 0)
	}
	if val, ok := kv["customer_id"].(string); ok {
		o.CustomerID = val
	} else {
		val := kv["customer_id"]
		if val == nil {
			o.CustomerID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.CustomerID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["due_date"].(string); ok {
		o.DueDate = val
	} else {
		val := kv["due_date"]
		if val == nil {
			o.DueDate = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.DueDate = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["id"].(string); ok {
		o.ID = val
	} else {
		val := kv["id"]
		if val == nil {
			o.ID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.ID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["identifier"].(string); ok {
		o.Identifier = val
	} else {
		val := kv["identifier"]
		if val == nil {
			o.Identifier = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Identifier = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["parent_id"].(string); ok {
		o.ParentID = val
	} else {
		val := kv["parent_id"]
		if val == nil {
			o.ParentID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.ParentID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["planned_end_ts"].(int64); ok {
		o.PlannedEndAt = val
	} else {
		val := kv["planned_end_ts"]
		if val == nil {
			o.PlannedEndAt = number.ToInt64Any(nil)
		} else {
			if tv, ok := val.(time.Time); ok {
				val = datetime.TimeToEpoch(tv)
			}
			o.PlannedEndAt = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["planned_start_ts"].(int64); ok {
		o.PlannedStartAt = val
	} else {
		val := kv["planned_start_ts"]
		if val == nil {
			o.PlannedStartAt = number.ToInt64Any(nil)
		} else {
			if tv, ok := val.(time.Time); ok {
				val = datetime.TimeToEpoch(tv)
			}
			o.PlannedStartAt = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["priority"].(string); ok {
		o.Priority = val
	} else {
		val := kv["priority"]
		if val == nil {
			o.Priority = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Priority = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["project_id"].(string); ok {
		o.ProjectID = val
	} else {
		val := kv["project_id"]
		if val == nil {
			o.ProjectID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.ProjectID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["ref_id"].(string); ok {
		o.RefID = val
	} else {
		val := kv["ref_id"]
		if val == nil {
			o.RefID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.RefID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["ref_type"].(string); ok {
		o.RefType = val
	} else {
		val := kv["ref_type"]
		if val == nil {
			o.RefType = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.RefType = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["reporter_ref_id"].(string); ok {
		o.ReporterRefID = val
	} else {
		val := kv["reporter_ref_id"]
		if val == nil {
			o.ReporterRefID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.ReporterRefID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["resolution"].(string); ok {
		o.Resolution = val
	} else {
		val := kv["resolution"]
		if val == nil {
			o.Resolution = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Resolution = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["status"].(string); ok {
		o.Status = val
	} else {
		val := kv["status"]
		if val == nil {
			o.Status = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Status = fmt.Sprintf("%v", val)
		}
	}
	if val := kv["tags"]; val != nil {
		na := make([]string, 0)
		if a, ok := val.([]string); ok {
			na = append(na, a...)
		} else {
			if a, ok := val.([]interface{}); ok {
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
	} else {
		o.Tags = []string{}
	}
	if o.Tags == nil {
		o.Tags = make([]string, 0)
	}
	if val, ok := kv["title"].(string); ok {
		o.Title = val
	} else {
		val := kv["title"]
		if val == nil {
			o.Title = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Title = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["type"].(string); ok {
		o.Type = val
	} else {
		val := kv["type"]
		if val == nil {
			o.Type = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Type = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["updated"].(string); ok {
		o.Updated = val
	} else {
		val := kv["updated"]
		if val == nil {
			o.Updated = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Updated = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["url"].(string); ok {
		o.URL = val
	} else {
		val := kv["url"]
		if val == nil {
			o.URL = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.URL = fmt.Sprintf("%v", val)
		}
	}
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *Issue) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.AssigneeRefID)
	args = append(args, o.Created)
	args = append(args, o.CreatorRefID)
	args = append(args, o.CustomFields)
	args = append(args, o.CustomerID)
	args = append(args, o.DueDate)
	args = append(args, o.ID)
	args = append(args, o.Identifier)
	args = append(args, o.ParentID)
	args = append(args, o.PlannedEndAt)
	args = append(args, o.PlannedStartAt)
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
	args = append(args, o.Updated)
	args = append(args, o.URL)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetIssueAvroSchemaSpec creates the avro schema specification for Issue
func GetIssueAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "work",
		"name":      "Issue",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "hashcode",
				"type": "string",
			},
			map[string]interface{}{
				"name": "assignee_ref_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "created",
				"type": "string",
			},
			map[string]interface{}{
				"name": "creator_ref_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "customFields",
				"type": map[string]interface{}{"name": "customFields", "items": map[string]interface{}{"type": "record", "name": "customFields", "fields": []interface{}{map[string]interface{}{"type": "string", "name": "id", "doc": "the id of the custom field"}, map[string]interface{}{"type": "string", "name": "name", "doc": "the name of the custom field"}, map[string]interface{}{"type": "string", "name": "value", "doc": "the value of the custom field"}}, "doc": "list of custom fields and their values"}, "type": "array"},
			},
			map[string]interface{}{
				"name": "customer_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "due_date",
				"type": "string",
			},
			map[string]interface{}{
				"name": "id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "identifier",
				"type": "string",
			},
			map[string]interface{}{
				"name": "parent_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "planned_end_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "planned_start_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "priority",
				"type": "string",
			},
			map[string]interface{}{
				"name": "project_id",
				"type": "string",
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
				"name": "reporter_ref_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "resolution",
				"type": "string",
			},
			map[string]interface{}{
				"name": "status",
				"type": "string",
			},
			map[string]interface{}{
				"name": "tags",
				"type": map[string]interface{}{"type": "array", "name": "tags", "items": "string"},
			},
			map[string]interface{}{
				"name": "title",
				"type": "string",
			},
			map[string]interface{}{
				"name": "type",
				"type": "string",
			},
			map[string]interface{}{
				"name": "updated",
				"type": "string",
			},
			map[string]interface{}{
				"name": "url",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetIssueAvroSchema creates the avro schema for Issue
func GetIssueAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetIssueAvroSchemaSpec())
}

// TransformIssueFunc is a function for transforming Issue during processing
type TransformIssueFunc func(input *Issue) (*Issue, error)

// NewIssuePipe creates a pipe for processing Issue items
func NewIssuePipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformIssueFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewIssueInputStream(input, errors)
	var stream chan Issue
	if len(transforms) > 0 {
		stream = make(chan Issue, 1000)
	} else {
		stream = inch
	}
	outdone := NewIssueOutputStream(output, stream, errors)
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

// NewIssueInputStreamDir creates a channel for reading Issue as JSON newlines from a directory of files
func NewIssueInputStreamDir(dir string, errors chan<- error, transforms ...TransformIssueFunc) (chan Issue, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/work/issue\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan Issue)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for issue")
		ch := make(chan Issue)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewIssueInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan Issue)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewIssueInputStreamFile creates an channel for reading Issue as JSON newlines from filename
func NewIssueInputStreamFile(filename string, errors chan<- error, transforms ...TransformIssueFunc) (chan Issue, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan Issue)
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
			ch := make(chan Issue)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewIssueInputStream(f, errors, transforms...)
}

// NewIssueInputStream creates an channel for reading Issue as JSON newlines from stream
func NewIssueInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformIssueFunc) (chan Issue, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan Issue, 1000)
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
			var item Issue
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

// NewIssueOutputStreamDir will output json newlines from channel and save in dir
func NewIssueOutputStreamDir(dir string, ch chan Issue, errors chan<- error, transforms ...TransformIssueFunc) <-chan bool {
	fp := filepath.Join(dir, "/work/issue\\.json(\\.gz)?$")
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
	return NewIssueOutputStream(gz, ch, errors, transforms...)
}

// NewIssueOutputStream will output json newlines from channel to the stream
func NewIssueOutputStream(stream io.WriteCloser, ch chan Issue, errors chan<- error, transforms ...TransformIssueFunc) <-chan bool {
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

// IssueSendEvent is an event detail for sending data
type IssueSendEvent struct {
	Issue   *Issue
	headers map[string]string
	time    time.Time
	key     string
}

var _ datamodel.ModelSendEvent = (*IssueSendEvent)(nil)

// Key is the key to use for the message
func (e *IssueSendEvent) Key() string {
	if e.key == "" {
		return e.Issue.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *IssueSendEvent) Object() datamodel.Model {
	return e.Issue
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *IssueSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *IssueSendEvent) Timestamp() time.Time {
	return e.time
}

// IssueSendEventOpts is a function handler for setting opts
type IssueSendEventOpts func(o *IssueSendEvent)

// WithIssueSendEventKey sets the key value to a value different than the object ID
func WithIssueSendEventKey(key string) IssueSendEventOpts {
	return func(o *IssueSendEvent) {
		o.key = key
	}
}

// WithIssueSendEventTimestamp sets the timestamp value
func WithIssueSendEventTimestamp(tv time.Time) IssueSendEventOpts {
	return func(o *IssueSendEvent) {
		o.time = tv
	}
}

// WithIssueSendEventHeader sets the timestamp value
func WithIssueSendEventHeader(key, value string) IssueSendEventOpts {
	return func(o *IssueSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewIssueSendEvent returns a new IssueSendEvent instance
func NewIssueSendEvent(o *Issue, opts ...IssueSendEventOpts) *IssueSendEvent {
	res := &IssueSendEvent{
		Issue: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewIssueProducer will stream data from the channel
func NewIssueProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
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
				if object, ok := item.Object().(*Issue); ok {
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
					errors <- fmt.Errorf("invalid event received. expected an object of type work.Issue but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewIssueConsumer will stream data from the topic into the provided channel
func NewIssueConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object Issue
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into work.Issue: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avro data into work.Issue: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for work.Issue")
			}

			// ignore messages that have exceeded the TTL
			cfg := object.GetTopicConfig()
			if cfg != nil && cfg.TTL != 0 && msg.Timestamp.UTC().Add(cfg.TTL).Sub(time.Now().UTC()) < 0 {
				return nil
			}
			msg.Codec = object.GetAvroCodec() // match the codec

			ch <- &IssueReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object Issue
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &IssueReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// IssueReceiveEvent is an event detail for receiving data
type IssueReceiveEvent struct {
	Issue   *Issue
	message eventing.Message
	eof     bool
}

var _ datamodel.ModelReceiveEvent = (*IssueReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *IssueReceiveEvent) Object() datamodel.Model {
	return e.Issue
}

// Message returns the underlying message data for the event
func (e *IssueReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *IssueReceiveEvent) EOF() bool {
	return e.eof
}

// IssueProducer implements the datamodel.ModelEventProducer
type IssueProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*IssueProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *IssueProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *IssueProducer) Close() error {
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
func (o *Issue) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *Issue) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &IssueProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewIssueProducer(newctx, producer, ch, errors, empty),
	}
}

// NewIssueProducerChannel returns a channel which can be used for producing Model events
func NewIssueProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewIssueProducerChannelSize(producer, 0, errors)
}

// NewIssueProducerChannelSize returns a channel which can be used for producing Model events
func NewIssueProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &IssueProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewIssueProducer(newctx, producer, ch, errors, empty),
	}
}

// IssueConsumer implements the datamodel.ModelEventConsumer
type IssueConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*IssueConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *IssueConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *IssueConsumer) Close() error {
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
func (o *Issue) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &IssueConsumer{
		ch:       ch,
		callback: NewIssueConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewIssueConsumerChannel returns a consumer channel which can be used to consume Model events
func NewIssueConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &IssueConsumer{
		ch:       ch,
		callback: NewIssueConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
