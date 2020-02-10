// DO NOT EDIT -- generated code

// Package agent -
package agent

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/bxcodec/faker"
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	pstrings "github.com/pinpt/go-common/strings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

const (

	// WebappIntegrationMutationRequestTable is the default table name
	WebappIntegrationMutationRequestTable datamodel.ModelNameType = "agent_webappintegrationmutationrequest"

	// WebappIntegrationMutationRequestModelName is the model name
	WebappIntegrationMutationRequestModelName datamodel.ModelNameType = "agent.WebappIntegrationMutationRequest"
)

// WebappIntegrationMutationRequestAction is the enumeration type for action
type WebappIntegrationMutationRequestAction int32

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
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v WebappIntegrationMutationRequestAction) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "ISSUE_ADD_COMMENT":
		v = 0
	case "ISSUE_SET_TITLE":
		v = 1
	case "ISSUE_SET_STATUS":
		v = 2
	case "ISSUE_SET_PRIORITY":
		v = 3
	case "ISSUE_SET_ASSIGNEE":
		v = 4
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
	}
	return "unset"
}

const (
	// ActionIssueAddComment is the enumeration value for issue_add_comment
	WebappIntegrationMutationRequestActionIssueAddComment WebappIntegrationMutationRequestAction = 0
	// ActionIssueSetTitle is the enumeration value for issue_set_title
	WebappIntegrationMutationRequestActionIssueSetTitle WebappIntegrationMutationRequestAction = 1
	// ActionIssueSetStatus is the enumeration value for issue_set_status
	WebappIntegrationMutationRequestActionIssueSetStatus WebappIntegrationMutationRequestAction = 2
	// ActionIssueSetPriority is the enumeration value for issue_set_priority
	WebappIntegrationMutationRequestActionIssueSetPriority WebappIntegrationMutationRequestAction = 3
	// ActionIssueSetAssignee is the enumeration value for issue_set_assignee
	WebappIntegrationMutationRequestActionIssueSetAssignee WebappIntegrationMutationRequestAction = 4
)

// WebappIntegrationMutationRequestSystemType is the enumeration type for system_type
type WebappIntegrationMutationRequestSystemType int32

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
func (v WebappIntegrationMutationRequestSystemType) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "WORK":
		v = 0
	case "SOURCECODE":
		v = 1
	case "CODEQUALITY":
		v = 2
	case "USER":
		v = 3
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

const (
	// SystemTypeWork is the enumeration value for work
	WebappIntegrationMutationRequestSystemTypeWork WebappIntegrationMutationRequestSystemType = 0
	// SystemTypeSourcecode is the enumeration value for sourcecode
	WebappIntegrationMutationRequestSystemTypeSourcecode WebappIntegrationMutationRequestSystemType = 1
	// SystemTypeCodequality is the enumeration value for codequality
	WebappIntegrationMutationRequestSystemTypeCodequality WebappIntegrationMutationRequestSystemType = 2
	// SystemTypeUser is the enumeration value for user
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
	// IntegrationName Name of the integration binary
	IntegrationName string `json:"integration_name" codec:"integration_name" bson:"integration_name" yaml:"integration_name" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
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

// IsEqual returns true if the two WebappIntegrationMutationRequest objects are equal
func (o *WebappIntegrationMutationRequest) IsEqual(other *WebappIntegrationMutationRequest) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *WebappIntegrationMutationRequest) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{

		"action":           o.Action.String(),
		"customer_id":      toWebappIntegrationMutationRequestObject(o.CustomerID, false),
		"data":             toWebappIntegrationMutationRequestObject(o.Data, false),
		"id":               toWebappIntegrationMutationRequestObject(o.ID, false),
		"integration_name": toWebappIntegrationMutationRequestObject(o.IntegrationName, false),
		"ref_id":           toWebappIntegrationMutationRequestObject(o.RefID, false),
		"ref_type":         toWebappIntegrationMutationRequestObject(o.RefType, false),

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
	args = append(args, o.IntegrationName)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.SystemType)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *WebappIntegrationMutationRequest) GetEventAPIConfig() datamodel.EventAPIConfig {
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
