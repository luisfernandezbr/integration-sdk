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

	// WebappIntegrationDataRequestTable is the default table name
	WebappIntegrationDataRequestTable datamodel.ModelNameType = "agent_webappintegrationdatarequest"

	// WebappIntegrationDataRequestModelName is the model name
	WebappIntegrationDataRequestModelName datamodel.ModelNameType = "agent.WebappIntegrationDataRequest"
)

const (
	// WebappIntegrationDataRequestModelActionColumn is the column json value action
	WebappIntegrationDataRequestModelActionColumn = "action"
	// WebappIntegrationDataRequestModelCustomerIDColumn is the column json value customer_id
	WebappIntegrationDataRequestModelCustomerIDColumn = "customer_id"
	// WebappIntegrationDataRequestModelDataColumn is the column json value data
	WebappIntegrationDataRequestModelDataColumn = "data"
	// WebappIntegrationDataRequestModelIDColumn is the column json value id
	WebappIntegrationDataRequestModelIDColumn = "id"
	// WebappIntegrationDataRequestModelIntegrationNameColumn is the column json value integration_name
	WebappIntegrationDataRequestModelIntegrationNameColumn = "integration_name"
	// WebappIntegrationDataRequestModelJobIDColumn is the column json value job_id
	WebappIntegrationDataRequestModelJobIDColumn = "job_id"
	// WebappIntegrationDataRequestModelRefIDColumn is the column json value ref_id
	WebappIntegrationDataRequestModelRefIDColumn = "ref_id"
	// WebappIntegrationDataRequestModelRefTypeColumn is the column json value ref_type
	WebappIntegrationDataRequestModelRefTypeColumn = "ref_type"
	// WebappIntegrationDataRequestModelSystemTypeColumn is the column json value system_type
	WebappIntegrationDataRequestModelSystemTypeColumn = "system_type"
)

// WebappIntegrationDataRequestAction is the enumeration type for action
type WebappIntegrationDataRequestAction int32

// UnmarshalBSONValue for unmarshaling value
func (v *WebappIntegrationDataRequestAction) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = WebappIntegrationDataRequestAction(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "ISSUE_GET_TRANSITIONS":
			*v = WebappIntegrationDataRequestAction(0)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v WebappIntegrationDataRequestAction) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "ISSUE_GET_TRANSITIONS":
		v = 0
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v WebappIntegrationDataRequestAction) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("ISSUE_GET_TRANSITIONS")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for Action
func (v WebappIntegrationDataRequestAction) String() string {
	switch int32(v) {
	case 0:
		return "ISSUE_GET_TRANSITIONS"
	}
	return "unset"
}

const (
	// WebappIntegrationDataRequestActionIssueGetTransitions is the enumeration value for issue_get_transitions
	WebappIntegrationDataRequestActionIssueGetTransitions WebappIntegrationDataRequestAction = 0
)

// WebappIntegrationDataRequestSystemType is the enumeration type for system_type
type WebappIntegrationDataRequestSystemType int32

// UnmarshalBSONValue for unmarshaling value
func (v *WebappIntegrationDataRequestSystemType) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = WebappIntegrationDataRequestSystemType(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "WORK":
			*v = WebappIntegrationDataRequestSystemType(0)
		case "SOURCECODE":
			*v = WebappIntegrationDataRequestSystemType(1)
		case "CODEQUALITY":
			*v = WebappIntegrationDataRequestSystemType(2)
		case "USER":
			*v = WebappIntegrationDataRequestSystemType(3)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v WebappIntegrationDataRequestSystemType) UnmarshalJSON(buf []byte) error {
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
func (v WebappIntegrationDataRequestSystemType) MarshalJSON() ([]byte, error) {
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
func (v WebappIntegrationDataRequestSystemType) String() string {
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
	// WebappIntegrationDataRequestSystemTypeWork is the enumeration value for work
	WebappIntegrationDataRequestSystemTypeWork WebappIntegrationDataRequestSystemType = 0
	// WebappIntegrationDataRequestSystemTypeSourcecode is the enumeration value for sourcecode
	WebappIntegrationDataRequestSystemTypeSourcecode WebappIntegrationDataRequestSystemType = 1
	// WebappIntegrationDataRequestSystemTypeCodequality is the enumeration value for codequality
	WebappIntegrationDataRequestSystemTypeCodequality WebappIntegrationDataRequestSystemType = 2
	// WebappIntegrationDataRequestSystemTypeUser is the enumeration value for user
	WebappIntegrationDataRequestSystemTypeUser WebappIntegrationDataRequestSystemType = 3
)

// WebappIntegrationDataRequest request to get data from integration for immediate display in the webapp (from webapp to agent operator)
type WebappIntegrationDataRequest struct {
	// Action Request action
	Action WebappIntegrationDataRequestAction `json:"action" codec:"action" bson:"action" yaml:"action" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Data Data request parameters that vary based on action. Stored as JSON string.
	Data string `json:"data" codec:"data" bson:"data" yaml:"data" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationName Name of the integration binary
	IntegrationName string `json:"integration_name" codec:"integration_name" bson:"integration_name" yaml:"integration_name" faker:"-"`
	// JobID The job ID
	JobID string `json:"job_id" codec:"job_id" bson:"job_id" yaml:"job_id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// SystemType The system type of the integration
	SystemType WebappIntegrationDataRequestSystemType `json:"system_type" codec:"system_type" bson:"system_type" yaml:"system_type" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*WebappIntegrationDataRequest)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*WebappIntegrationDataRequest)(nil)

func toWebappIntegrationDataRequestObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *WebappIntegrationDataRequest:
		return v.ToMap()

	case WebappIntegrationDataRequestAction:
		return v.String()

	case WebappIntegrationDataRequestSystemType:
		return v.String()
	default:
		return o
	}
}

// String returns a string representation of WebappIntegrationDataRequest
func (o *WebappIntegrationDataRequest) String() string {
	return fmt.Sprintf("agent.WebappIntegrationDataRequest<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *WebappIntegrationDataRequest) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *WebappIntegrationDataRequest) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *WebappIntegrationDataRequest) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *WebappIntegrationDataRequest) GetModelName() datamodel.ModelNameType {
	return WebappIntegrationDataRequestModelName
}

// NewWebappIntegrationDataRequestID provides a template for generating an ID field for WebappIntegrationDataRequest
func NewWebappIntegrationDataRequestID(customerID string, refType string, refID string) string {
	return hash.Values("WebappIntegrationDataRequest", customerID, refType, refID)
}

func (o *WebappIntegrationDataRequest) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("WebappIntegrationDataRequest", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *WebappIntegrationDataRequest) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *WebappIntegrationDataRequest) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *WebappIntegrationDataRequest) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *WebappIntegrationDataRequest) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *WebappIntegrationDataRequest) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *WebappIntegrationDataRequest) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *WebappIntegrationDataRequest) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *WebappIntegrationDataRequest) IsEvented() bool {
	return false
}

// GetTopicConfig returns the topic config object
func (o *WebappIntegrationDataRequest) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *WebappIntegrationDataRequest) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of WebappIntegrationDataRequest
func (o *WebappIntegrationDataRequest) Clone() datamodel.Model {
	c := new(WebappIntegrationDataRequest)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *WebappIntegrationDataRequest) Anon() datamodel.Model {
	c := new(WebappIntegrationDataRequest)
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
func (o *WebappIntegrationDataRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *WebappIntegrationDataRequest) UnmarshalJSON(data []byte) error {
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
func (o *WebappIntegrationDataRequest) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *WebappIntegrationDataRequest) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two WebappIntegrationDataRequest objects are equal
func (o *WebappIntegrationDataRequest) IsEqual(other *WebappIntegrationDataRequest) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *WebappIntegrationDataRequest) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{

		"action":           o.Action.String(),
		"customer_id":      toWebappIntegrationDataRequestObject(o.CustomerID, false),
		"data":             toWebappIntegrationDataRequestObject(o.Data, false),
		"id":               toWebappIntegrationDataRequestObject(o.ID, false),
		"integration_name": toWebappIntegrationDataRequestObject(o.IntegrationName, false),
		"job_id":           toWebappIntegrationDataRequestObject(o.JobID, false),
		"ref_id":           toWebappIntegrationDataRequestObject(o.RefID, false),
		"ref_type":         toWebappIntegrationDataRequestObject(o.RefType, false),

		"system_type": o.SystemType.String(),
		"hashcode":    toWebappIntegrationDataRequestObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *WebappIntegrationDataRequest) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["action"].(WebappIntegrationDataRequestAction); ok {
		o.Action = val
	} else {
		if em, ok := kv["action"].(map[string]interface{}); ok {
			ev := em["agent.action"].(string)
			switch ev {
			case "issue_get_transitions", "ISSUE_GET_TRANSITIONS":
				o.Action = 0
			}
		}
		if em, ok := kv["action"].(string); ok {
			switch em {
			case "issue_get_transitions", "ISSUE_GET_TRANSITIONS":
				o.Action = 0
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

	if val, ok := kv["system_type"].(WebappIntegrationDataRequestSystemType); ok {
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
func (o *WebappIntegrationDataRequest) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Action)
	args = append(args, o.CustomerID)
	args = append(args, o.Data)
	args = append(args, o.ID)
	args = append(args, o.IntegrationName)
	args = append(args, o.JobID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.SystemType)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *WebappIntegrationDataRequest) GetEventAPIConfig() datamodel.EventAPIConfig {
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
