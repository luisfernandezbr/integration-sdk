// DO NOT EDIT -- generated code

// Package agent -
package agent

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/bxcodec/faker"
	"github.com/pinpt/go-common/v10/datamodel"
	"github.com/pinpt/go-common/v10/datetime"
	"github.com/pinpt/go-common/v10/hash"
	pjson "github.com/pinpt/go-common/v10/json"
	"github.com/pinpt/go-common/v10/number"
	pstrings "github.com/pinpt/go-common/v10/strings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

const (

	// WebappIntegrationMutationRequestTable is the default table name
	WebappIntegrationMutationRequestTable datamodel.ModelNameType = "agent_webappintegrationmutationrequest"

	// WebappIntegrationMutationRequestModelName is the model name
	WebappIntegrationMutationRequestModelName datamodel.ModelNameType = "agent.WebappIntegrationMutationRequest"
)

const (
	// WebappIntegrationMutationRequestModelActionColumn is the column json value action
	WebappIntegrationMutationRequestModelActionColumn = "action"
	// WebappIntegrationMutationRequestModelCustomerIDColumn is the column json value customer_id
	WebappIntegrationMutationRequestModelCustomerIDColumn = "customer_id"
	// WebappIntegrationMutationRequestModelDataColumn is the column json value data
	WebappIntegrationMutationRequestModelDataColumn = "data"
	// WebappIntegrationMutationRequestModelIDColumn is the column json value id
	WebappIntegrationMutationRequestModelIDColumn = "id"
	// WebappIntegrationMutationRequestModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	WebappIntegrationMutationRequestModelIntegrationInstanceIDColumn = "integration_instance_id"
	// WebappIntegrationMutationRequestModelIntegrationNameColumn is the column json value integration_name
	WebappIntegrationMutationRequestModelIntegrationNameColumn = "integration_name"
	// WebappIntegrationMutationRequestModelJobIDColumn is the column json value job_id
	WebappIntegrationMutationRequestModelJobIDColumn = "job_id"
	// WebappIntegrationMutationRequestModelRefIDColumn is the column json value ref_id
	WebappIntegrationMutationRequestModelRefIDColumn = "ref_id"
	// WebappIntegrationMutationRequestModelRefTypeColumn is the column json value ref_type
	WebappIntegrationMutationRequestModelRefTypeColumn = "ref_type"
	// WebappIntegrationMutationRequestModelRequestDateColumn is the column json value request_date
	WebappIntegrationMutationRequestModelRequestDateColumn = "request_date"
	// WebappIntegrationMutationRequestModelRequestDateEpochColumn is the column json value epoch
	WebappIntegrationMutationRequestModelRequestDateEpochColumn = "epoch"
	// WebappIntegrationMutationRequestModelRequestDateOffsetColumn is the column json value offset
	WebappIntegrationMutationRequestModelRequestDateOffsetColumn = "offset"
	// WebappIntegrationMutationRequestModelRequestDateRfc3339Column is the column json value rfc3339
	WebappIntegrationMutationRequestModelRequestDateRfc3339Column = "rfc3339"
	// WebappIntegrationMutationRequestModelSystemTypeColumn is the column json value system_type
	WebappIntegrationMutationRequestModelSystemTypeColumn = "system_type"
)

// WebappIntegrationMutationRequestAction is the enumeration type for action
type WebappIntegrationMutationRequestAction int32

// toWebappIntegrationMutationRequestActionPointer is the enumeration pointer type for action
func toWebappIntegrationMutationRequestActionPointer(v int32) *WebappIntegrationMutationRequestAction {
	nv := WebappIntegrationMutationRequestAction(v)
	return &nv
}

// toWebappIntegrationMutationRequestActionEnum is the enumeration pointer wrapper for action
func toWebappIntegrationMutationRequestActionEnum(v *WebappIntegrationMutationRequestAction) string {
	if v == nil {
		return toWebappIntegrationMutationRequestActionPointer(0).String()
	}
	return v.String()
}

// UnmarshalBSONValue for unmarshaling value
func (v *WebappIntegrationMutationRequestAction) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = WebappIntegrationMutationRequestAction(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "ISSUE_ADD_COMMENT":
			*v = WebappIntegrationMutationRequestAction(0)
		case "ISSUE_SET_TITLE":
			*v = WebappIntegrationMutationRequestAction(1)
		case "ISSUE_SET_STATUS":
			*v = WebappIntegrationMutationRequestAction(2)
		case "ISSUE_SET_PRIORITY":
			*v = WebappIntegrationMutationRequestAction(3)
		case "ISSUE_SET_ASSIGNEE":
			*v = WebappIntegrationMutationRequestAction(4)
		case "ISSUE_GET_TRANSITIONS":
			*v = WebappIntegrationMutationRequestAction(5)
		case "PR_SET_TITLE":
			*v = WebappIntegrationMutationRequestAction(6)
		case "PR_SET_DESCRIPTION":
			*v = WebappIntegrationMutationRequestAction(7)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v *WebappIntegrationMutationRequestAction) UnmarshalJSON(buf []byte) error {
	var val string
	if err := json.Unmarshal(buf, &val); err != nil {
		return err
	}
	switch val {
	case "ISSUE_ADD_COMMENT":
		*v = 0
	case "ISSUE_SET_TITLE":
		*v = 1
	case "ISSUE_SET_STATUS":
		*v = 2
	case "ISSUE_SET_PRIORITY":
		*v = 3
	case "ISSUE_SET_ASSIGNEE":
		*v = 4
	case "ISSUE_GET_TRANSITIONS":
		*v = 5
	case "PR_SET_TITLE":
		*v = 6
	case "PR_SET_DESCRIPTION":
		*v = 7
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v WebappIntegrationMutationRequestAction) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("ISSUE_ADD_COMMENT")
	case 1:
		return json.Marshal("ISSUE_SET_TITLE")
	case 2:
		return json.Marshal("ISSUE_SET_STATUS")
	case 3:
		return json.Marshal("ISSUE_SET_PRIORITY")
	case 4:
		return json.Marshal("ISSUE_SET_ASSIGNEE")
	case 5:
		return json.Marshal("ISSUE_GET_TRANSITIONS")
	case 6:
		return json.Marshal("PR_SET_TITLE")
	case 7:
		return json.Marshal("PR_SET_DESCRIPTION")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for Action
func (v WebappIntegrationMutationRequestAction) String() string {
	switch int32(v) {
	case 0:
		return "ISSUE_ADD_COMMENT"
	case 1:
		return "ISSUE_SET_TITLE"
	case 2:
		return "ISSUE_SET_STATUS"
	case 3:
		return "ISSUE_SET_PRIORITY"
	case 4:
		return "ISSUE_SET_ASSIGNEE"
	case 5:
		return "ISSUE_GET_TRANSITIONS"
	case 6:
		return "PR_SET_TITLE"
	case 7:
		return "PR_SET_DESCRIPTION"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *WebappIntegrationMutationRequestAction) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case WebappIntegrationMutationRequestAction:
		*v = val
	case int32:
		*v = WebappIntegrationMutationRequestAction(int32(val))
	case int:
		*v = WebappIntegrationMutationRequestAction(int32(val))
	case string:
		switch val {
		case "ISSUE_ADD_COMMENT":
			*v = WebappIntegrationMutationRequestAction(0)
		case "ISSUE_SET_TITLE":
			*v = WebappIntegrationMutationRequestAction(1)
		case "ISSUE_SET_STATUS":
			*v = WebappIntegrationMutationRequestAction(2)
		case "ISSUE_SET_PRIORITY":
			*v = WebappIntegrationMutationRequestAction(3)
		case "ISSUE_SET_ASSIGNEE":
			*v = WebappIntegrationMutationRequestAction(4)
		case "ISSUE_GET_TRANSITIONS":
			*v = WebappIntegrationMutationRequestAction(5)
		case "PR_SET_TITLE":
			*v = WebappIntegrationMutationRequestAction(6)
		case "PR_SET_DESCRIPTION":
			*v = WebappIntegrationMutationRequestAction(7)
		}
	}
	return nil
}

const (
	// WebappIntegrationMutationRequestActionIssueAddComment is the enumeration value for issue_add_comment
	WebappIntegrationMutationRequestActionIssueAddComment WebappIntegrationMutationRequestAction = 0
	// WebappIntegrationMutationRequestActionIssueSetTitle is the enumeration value for issue_set_title
	WebappIntegrationMutationRequestActionIssueSetTitle WebappIntegrationMutationRequestAction = 1
	// WebappIntegrationMutationRequestActionIssueSetStatus is the enumeration value for issue_set_status
	WebappIntegrationMutationRequestActionIssueSetStatus WebappIntegrationMutationRequestAction = 2
	// WebappIntegrationMutationRequestActionIssueSetPriority is the enumeration value for issue_set_priority
	WebappIntegrationMutationRequestActionIssueSetPriority WebappIntegrationMutationRequestAction = 3
	// WebappIntegrationMutationRequestActionIssueSetAssignee is the enumeration value for issue_set_assignee
	WebappIntegrationMutationRequestActionIssueSetAssignee WebappIntegrationMutationRequestAction = 4
	// WebappIntegrationMutationRequestActionIssueGetTransitions is the enumeration value for issue_get_transitions
	WebappIntegrationMutationRequestActionIssueGetTransitions WebappIntegrationMutationRequestAction = 5
	// WebappIntegrationMutationRequestActionPrSetTitle is the enumeration value for pr_set_title
	WebappIntegrationMutationRequestActionPrSetTitle WebappIntegrationMutationRequestAction = 6
	// WebappIntegrationMutationRequestActionPrSetDescription is the enumeration value for pr_set_description
	WebappIntegrationMutationRequestActionPrSetDescription WebappIntegrationMutationRequestAction = 7
)

// WebappIntegrationMutationRequestRequestDate represents the object structure for request_date
type WebappIntegrationMutationRequestRequestDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toWebappIntegrationMutationRequestRequestDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *WebappIntegrationMutationRequestRequestDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *WebappIntegrationMutationRequestRequestDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toWebappIntegrationMutationRequestRequestDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toWebappIntegrationMutationRequestRequestDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toWebappIntegrationMutationRequestRequestDateObject(o.Rfc3339, false),
	}
}

func (o *WebappIntegrationMutationRequestRequestDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *WebappIntegrationMutationRequestRequestDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["epoch"].(int64); ok {
		o.Epoch = val
	} else {
		if val, ok := kv["epoch"]; ok {
			if val == nil {
				o.Epoch = 0
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
				o.Offset = 0
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

// WebappIntegrationMutationRequestSystemType is the enumeration type for system_type
type WebappIntegrationMutationRequestSystemType int32

// toWebappIntegrationMutationRequestSystemTypePointer is the enumeration pointer type for system_type
func toWebappIntegrationMutationRequestSystemTypePointer(v int32) *WebappIntegrationMutationRequestSystemType {
	nv := WebappIntegrationMutationRequestSystemType(v)
	return &nv
}

// toWebappIntegrationMutationRequestSystemTypeEnum is the enumeration pointer wrapper for system_type
func toWebappIntegrationMutationRequestSystemTypeEnum(v *WebappIntegrationMutationRequestSystemType) string {
	if v == nil {
		return toWebappIntegrationMutationRequestSystemTypePointer(0).String()
	}
	return v.String()
}

// UnmarshalBSONValue for unmarshaling value
func (v *WebappIntegrationMutationRequestSystemType) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = WebappIntegrationMutationRequestSystemType(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "WORK":
			*v = WebappIntegrationMutationRequestSystemType(0)
		case "SOURCECODE":
			*v = WebappIntegrationMutationRequestSystemType(1)
		case "CODEQUALITY":
			*v = WebappIntegrationMutationRequestSystemType(2)
		case "USER":
			*v = WebappIntegrationMutationRequestSystemType(3)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v *WebappIntegrationMutationRequestSystemType) UnmarshalJSON(buf []byte) error {
	var val string
	if err := json.Unmarshal(buf, &val); err != nil {
		return err
	}
	switch val {
	case "WORK":
		*v = 0
	case "SOURCECODE":
		*v = 1
	case "CODEQUALITY":
		*v = 2
	case "USER":
		*v = 3
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v WebappIntegrationMutationRequestSystemType) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("WORK")
	case 1:
		return json.Marshal("SOURCECODE")
	case 2:
		return json.Marshal("CODEQUALITY")
	case 3:
		return json.Marshal("USER")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for SystemType
func (v WebappIntegrationMutationRequestSystemType) String() string {
	switch int32(v) {
	case 0:
		return "WORK"
	case 1:
		return "SOURCECODE"
	case 2:
		return "CODEQUALITY"
	case 3:
		return "USER"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *WebappIntegrationMutationRequestSystemType) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case WebappIntegrationMutationRequestSystemType:
		*v = val
	case int32:
		*v = WebappIntegrationMutationRequestSystemType(int32(val))
	case int:
		*v = WebappIntegrationMutationRequestSystemType(int32(val))
	case string:
		switch val {
		case "WORK":
			*v = WebappIntegrationMutationRequestSystemType(0)
		case "SOURCECODE":
			*v = WebappIntegrationMutationRequestSystemType(1)
		case "CODEQUALITY":
			*v = WebappIntegrationMutationRequestSystemType(2)
		case "USER":
			*v = WebappIntegrationMutationRequestSystemType(3)
		}
	}
	return nil
}

const (
	// WebappIntegrationMutationRequestSystemTypeWork is the enumeration value for work
	WebappIntegrationMutationRequestSystemTypeWork WebappIntegrationMutationRequestSystemType = 0
	// WebappIntegrationMutationRequestSystemTypeSourcecode is the enumeration value for sourcecode
	WebappIntegrationMutationRequestSystemTypeSourcecode WebappIntegrationMutationRequestSystemType = 1
	// WebappIntegrationMutationRequestSystemTypeCodequality is the enumeration value for codequality
	WebappIntegrationMutationRequestSystemTypeCodequality WebappIntegrationMutationRequestSystemType = 2
	// WebappIntegrationMutationRequestSystemTypeUser is the enumeration value for user
	WebappIntegrationMutationRequestSystemTypeUser WebappIntegrationMutationRequestSystemType = 3
)

// WebappIntegrationMutationRequest request to change data in integration going from webapp to agent service
type WebappIntegrationMutationRequest struct {
	// Action Action to perform
	Action WebappIntegrationMutationRequestAction `json:"action" codec:"action" bson:"action" yaml:"action" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Data Mutation parameters that vary based on action as JSON string.
	Data string `json:"data" codec:"data" bson:"data" yaml:"data" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// IntegrationName Name of the integration binary
	IntegrationName string `json:"integration_name" codec:"integration_name" bson:"integration_name" yaml:"integration_name" faker:"-"`
	// JobID The job ID
	JobID string `json:"job_id" codec:"job_id" bson:"job_id" yaml:"job_id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RequestDate set by webapp using browser time
	RequestDate WebappIntegrationMutationRequestRequestDate `json:"request_date" codec:"request_date" bson:"request_date" yaml:"request_date" faker:"-"`
	// SystemType The system type of the integration
	SystemType WebappIntegrationMutationRequestSystemType `json:"system_type" codec:"system_type" bson:"system_type" yaml:"system_type" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*WebappIntegrationMutationRequest)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*WebappIntegrationMutationRequest)(nil)

func toWebappIntegrationMutationRequestObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *WebappIntegrationMutationRequest:
		return v.ToMap()

	case WebappIntegrationMutationRequestAction:
		return v.String()

	case WebappIntegrationMutationRequestRequestDate:
		return v.ToMap()

	case WebappIntegrationMutationRequestSystemType:
		return v.String()
	default:
		return o
	}
}

// String returns a string representation of WebappIntegrationMutationRequest
func (o *WebappIntegrationMutationRequest) String() string {
	return fmt.Sprintf("agent.WebappIntegrationMutationRequest<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *WebappIntegrationMutationRequest) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *WebappIntegrationMutationRequest) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *WebappIntegrationMutationRequest) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *WebappIntegrationMutationRequest) GetModelName() datamodel.ModelNameType {
	return WebappIntegrationMutationRequestModelName
}

// NewWebappIntegrationMutationRequestID provides a template for generating an ID field for WebappIntegrationMutationRequest
func NewWebappIntegrationMutationRequestID(customerID string, refType string, refID string) string {
	return hash.Values("WebappIntegrationMutationRequest", customerID, refType, refID)
}

func (o *WebappIntegrationMutationRequest) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("WebappIntegrationMutationRequest", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *WebappIntegrationMutationRequest) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *WebappIntegrationMutationRequest) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *WebappIntegrationMutationRequest) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *WebappIntegrationMutationRequest) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *WebappIntegrationMutationRequest) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *WebappIntegrationMutationRequest) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *WebappIntegrationMutationRequest) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *WebappIntegrationMutationRequest) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *WebappIntegrationMutationRequest) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = WebappIntegrationMutationRequestModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *WebappIntegrationMutationRequest) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *WebappIntegrationMutationRequest) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of WebappIntegrationMutationRequest
func (o *WebappIntegrationMutationRequest) Clone() datamodel.Model {
	c := new(WebappIntegrationMutationRequest)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *WebappIntegrationMutationRequest) Anon() datamodel.Model {
	c := new(WebappIntegrationMutationRequest)
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
func (o *WebappIntegrationMutationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *WebappIntegrationMutationRequest) UnmarshalJSON(data []byte) error {
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
func (o *WebappIntegrationMutationRequest) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *WebappIntegrationMutationRequest) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two WebappIntegrationMutationRequest objects are equal
func (o *WebappIntegrationMutationRequest) IsEqual(other *WebappIntegrationMutationRequest) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *WebappIntegrationMutationRequest) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{

		"action":                  o.Action.String(),
		"customer_id":             toWebappIntegrationMutationRequestObject(o.CustomerID, false),
		"data":                    toWebappIntegrationMutationRequestObject(o.Data, false),
		"id":                      toWebappIntegrationMutationRequestObject(o.ID, false),
		"integration_instance_id": toWebappIntegrationMutationRequestObject(o.IntegrationInstanceID, true),
		"integration_name":        toWebappIntegrationMutationRequestObject(o.IntegrationName, false),
		"job_id":                  toWebappIntegrationMutationRequestObject(o.JobID, false),
		"ref_id":                  toWebappIntegrationMutationRequestObject(o.RefID, false),
		"ref_type":                toWebappIntegrationMutationRequestObject(o.RefType, false),
		"request_date":            toWebappIntegrationMutationRequestObject(o.RequestDate, false),

		"system_type": o.SystemType.String(),
		"hashcode":    toWebappIntegrationMutationRequestObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *WebappIntegrationMutationRequest) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["action"].(WebappIntegrationMutationRequestAction); ok {
		o.Action = val
	} else {
		if em, ok := kv["action"].(map[string]interface{}); ok {

			ev := em["agent.action"].(string)
			switch ev {
			case "issue_add_comment", "ISSUE_ADD_COMMENT":
				o.Action = 0
			case "issue_set_title", "ISSUE_SET_TITLE":
				o.Action = 1
			case "issue_set_status", "ISSUE_SET_STATUS":
				o.Action = 2
			case "issue_set_priority", "ISSUE_SET_PRIORITY":
				o.Action = 3
			case "issue_set_assignee", "ISSUE_SET_ASSIGNEE":
				o.Action = 4
			case "issue_get_transitions", "ISSUE_GET_TRANSITIONS":
				o.Action = 5
			case "pr_set_title", "PR_SET_TITLE":
				o.Action = 6
			case "pr_set_description", "PR_SET_DESCRIPTION":
				o.Action = 7
			}
		}
		if em, ok := kv["action"].(string); ok {
			switch em {
			case "issue_add_comment", "ISSUE_ADD_COMMENT":
				o.Action = 0
			case "issue_set_title", "ISSUE_SET_TITLE":
				o.Action = 1
			case "issue_set_status", "ISSUE_SET_STATUS":
				o.Action = 2
			case "issue_set_priority", "ISSUE_SET_PRIORITY":
				o.Action = 3
			case "issue_set_assignee", "ISSUE_SET_ASSIGNEE":
				o.Action = 4
			case "issue_get_transitions", "ISSUE_GET_TRANSITIONS":
				o.Action = 5
			case "pr_set_title", "PR_SET_TITLE":
				o.Action = 6
			case "pr_set_description", "PR_SET_DESCRIPTION":
				o.Action = 7
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
	if val, ok := kv["data"].(string); ok {
		o.Data = val
	} else {
		if val, ok := kv["data"]; ok {
			if val == nil {
				o.Data = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Data = fmt.Sprintf("%v", val)
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
	if val, ok := kv["integration_instance_id"].(*string); ok {
		o.IntegrationInstanceID = val
	} else if val, ok := kv["integration_instance_id"].(string); ok {
		o.IntegrationInstanceID = &val
	} else {
		if val, ok := kv["integration_instance_id"]; ok {
			if val == nil {
				o.IntegrationInstanceID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.IntegrationInstanceID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["integration_name"].(string); ok {
		o.IntegrationName = val
	} else {
		if val, ok := kv["integration_name"]; ok {
			if val == nil {
				o.IntegrationName = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.IntegrationName = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["job_id"].(string); ok {
		o.JobID = val
	} else {
		if val, ok := kv["job_id"]; ok {
			if val == nil {
				o.JobID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.JobID = fmt.Sprintf("%v", val)
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

	if val, ok := kv["request_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.RequestDate.FromMap(kv)
		} else if sv, ok := val.(WebappIntegrationMutationRequestRequestDate); ok {
			// struct
			o.RequestDate = sv
		} else if sp, ok := val.(*WebappIntegrationMutationRequestRequestDate); ok {
			// struct pointer
			o.RequestDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.RequestDate.Epoch = dt.Epoch
			o.RequestDate.Rfc3339 = dt.Rfc3339
			o.RequestDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.RequestDate.Epoch = dt.Epoch
			o.RequestDate.Rfc3339 = dt.Rfc3339
			o.RequestDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.RequestDate.Epoch = dt.Epoch
				o.RequestDate.Rfc3339 = dt.Rfc3339
				o.RequestDate.Offset = dt.Offset
			}
		}
	} else {
		o.RequestDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["system_type"].(WebappIntegrationMutationRequestSystemType); ok {
		o.SystemType = val
	} else {
		if em, ok := kv["system_type"].(map[string]interface{}); ok {

			ev := em["agent.system_type"].(string)
			switch ev {
			case "work", "WORK":
				o.SystemType = 0
			case "sourcecode", "SOURCECODE":
				o.SystemType = 1
			case "codequality", "CODEQUALITY":
				o.SystemType = 2
			case "user", "USER":
				o.SystemType = 3
			}
		}
		if em, ok := kv["system_type"].(string); ok {
			switch em {
			case "work", "WORK":
				o.SystemType = 0
			case "sourcecode", "SOURCECODE":
				o.SystemType = 1
			case "codequality", "CODEQUALITY":
				o.SystemType = 2
			case "user", "USER":
				o.SystemType = 3
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *WebappIntegrationMutationRequest) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Action)
	args = append(args, o.CustomerID)
	args = append(args, o.Data)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.IntegrationName)
	args = append(args, o.JobID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RequestDate)
	args = append(args, o.SystemType)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// SetIntegrationInstanceID will set the integration instance ID
func (o *WebappIntegrationMutationRequest) SetIntegrationInstanceID(id string) {
	if id == "" {
		o.IntegrationInstanceID = nil
	} else {
		o.IntegrationInstanceID = &id
	}
}

// GetIntegrationInstanceID will return the integration instance ID
func (o *WebappIntegrationMutationRequest) GetIntegrationInstanceID() *string {
	return o.IntegrationInstanceID
}

// WebappIntegrationMutationRequestPartial is a partial struct for upsert mutations for WebappIntegrationMutationRequest
type WebappIntegrationMutationRequestPartial struct {
	// Action Action to perform
	Action *WebappIntegrationMutationRequestAction `json:"action,omitempty"`
	// Data Mutation parameters that vary based on action as JSON string.
	Data *string `json:"data,omitempty"`
	// IntegrationName Name of the integration binary
	IntegrationName *string `json:"integration_name,omitempty"`
	// JobID The job ID
	JobID *string `json:"job_id,omitempty"`
	// RequestDate set by webapp using browser time
	RequestDate *WebappIntegrationMutationRequestRequestDate `json:"request_date,omitempty"`
	// SystemType The system type of the integration
	SystemType *WebappIntegrationMutationRequestSystemType `json:"system_type,omitempty"`
}

var _ datamodel.PartialModel = (*WebappIntegrationMutationRequestPartial)(nil)

// GetModelName returns the name of the model
func (o *WebappIntegrationMutationRequestPartial) GetModelName() datamodel.ModelNameType {
	return WebappIntegrationMutationRequestModelName
}

// ToMap returns the object as a map
func (o *WebappIntegrationMutationRequestPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{

		"action":           toWebappIntegrationMutationRequestActionEnum(o.Action),
		"data":             toWebappIntegrationMutationRequestObject(o.Data, true),
		"integration_name": toWebappIntegrationMutationRequestObject(o.IntegrationName, true),
		"job_id":           toWebappIntegrationMutationRequestObject(o.JobID, true),
		"request_date":     toWebappIntegrationMutationRequestObject(o.RequestDate, true),

		"system_type": toWebappIntegrationMutationRequestSystemTypeEnum(o.SystemType),
	}
	for k, v := range kv {
		if v == nil || reflect.ValueOf(v).IsZero() {
			delete(kv, k)
		} else {
			if k == "request_date" {
				if dt, ok := v.(*WebappIntegrationMutationRequestRequestDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}
		}
	}
	return kv
}

// Stringify returns the object in JSON format as a string
func (o *WebappIntegrationMutationRequestPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *WebappIntegrationMutationRequestPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *WebappIntegrationMutationRequestPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *WebappIntegrationMutationRequestPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *WebappIntegrationMutationRequestPartial) FromMap(kv map[string]interface{}) {
	if val, ok := kv["action"].(*WebappIntegrationMutationRequestAction); ok {
		o.Action = val
	} else if val, ok := kv["action"].(WebappIntegrationMutationRequestAction); ok {
		o.Action = &val
	} else {
		if val, ok := kv["action"]; ok {
			if val == nil {
				o.Action = toWebappIntegrationMutationRequestActionPointer(0)
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["WebappIntegrationMutationRequestAction"]
				}
				// this is an enum pointer
				if em, ok := val.(string); ok {
					switch em {
					case "issue_add_comment", "ISSUE_ADD_COMMENT":
						o.Action = toWebappIntegrationMutationRequestActionPointer(0)
					case "issue_set_title", "ISSUE_SET_TITLE":
						o.Action = toWebappIntegrationMutationRequestActionPointer(1)
					case "issue_set_status", "ISSUE_SET_STATUS":
						o.Action = toWebappIntegrationMutationRequestActionPointer(2)
					case "issue_set_priority", "ISSUE_SET_PRIORITY":
						o.Action = toWebappIntegrationMutationRequestActionPointer(3)
					case "issue_set_assignee", "ISSUE_SET_ASSIGNEE":
						o.Action = toWebappIntegrationMutationRequestActionPointer(4)
					case "issue_get_transitions", "ISSUE_GET_TRANSITIONS":
						o.Action = toWebappIntegrationMutationRequestActionPointer(5)
					case "pr_set_title", "PR_SET_TITLE":
						o.Action = toWebappIntegrationMutationRequestActionPointer(6)
					case "pr_set_description", "PR_SET_DESCRIPTION":
						o.Action = toWebappIntegrationMutationRequestActionPointer(7)
					}
				}
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
	if val, ok := kv["integration_name"].(*string); ok {
		o.IntegrationName = val
	} else if val, ok := kv["integration_name"].(string); ok {
		o.IntegrationName = &val
	} else {
		if val, ok := kv["integration_name"]; ok {
			if val == nil {
				o.IntegrationName = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.IntegrationName = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["job_id"].(*string); ok {
		o.JobID = val
	} else if val, ok := kv["job_id"].(string); ok {
		o.JobID = &val
	} else {
		if val, ok := kv["job_id"]; ok {
			if val == nil {
				o.JobID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.JobID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if o.RequestDate == nil {
		o.RequestDate = &WebappIntegrationMutationRequestRequestDate{}
	}

	if val, ok := kv["request_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.RequestDate.FromMap(kv)
		} else if sv, ok := val.(WebappIntegrationMutationRequestRequestDate); ok {
			// struct
			o.RequestDate = &sv
		} else if sp, ok := val.(*WebappIntegrationMutationRequestRequestDate); ok {
			// struct pointer
			o.RequestDate = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.RequestDate.Epoch = dt.Epoch
			o.RequestDate.Rfc3339 = dt.Rfc3339
			o.RequestDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.RequestDate.Epoch = dt.Epoch
			o.RequestDate.Rfc3339 = dt.Rfc3339
			o.RequestDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.RequestDate.Epoch = dt.Epoch
				o.RequestDate.Rfc3339 = dt.Rfc3339
				o.RequestDate.Offset = dt.Offset
			}
		}
	} else {
		o.RequestDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["system_type"].(*WebappIntegrationMutationRequestSystemType); ok {
		o.SystemType = val
	} else if val, ok := kv["system_type"].(WebappIntegrationMutationRequestSystemType); ok {
		o.SystemType = &val
	} else {
		if val, ok := kv["system_type"]; ok {
			if val == nil {
				o.SystemType = toWebappIntegrationMutationRequestSystemTypePointer(0)
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["WebappIntegrationMutationRequestSystemType"]
				}
				// this is an enum pointer
				if em, ok := val.(string); ok {
					switch em {
					case "work", "WORK":
						o.SystemType = toWebappIntegrationMutationRequestSystemTypePointer(0)
					case "sourcecode", "SOURCECODE":
						o.SystemType = toWebappIntegrationMutationRequestSystemTypePointer(1)
					case "codequality", "CODEQUALITY":
						o.SystemType = toWebappIntegrationMutationRequestSystemTypePointer(2)
					case "user", "USER":
						o.SystemType = toWebappIntegrationMutationRequestSystemTypePointer(3)
					}
				}
			}
		}
	}
	o.setDefaults(false)
}
