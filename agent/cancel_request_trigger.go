// DO NOT EDIT -- generated code

// Package agent - the agent communicates with pinpoint cloud to send integration data
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

	// CancelRequestTriggerTable is the default table name
	CancelRequestTriggerTable datamodel.ModelNameType = "agent_cancelrequesttrigger"

	// CancelRequestTriggerModelName is the model name
	CancelRequestTriggerModelName datamodel.ModelNameType = "agent.CancelRequestTrigger"
)

const (
	// CancelRequestTriggerModelCommandColumn is the column json value command
	CancelRequestTriggerModelCommandColumn = "command"
	// CancelRequestTriggerModelCustomerIDColumn is the column json value customer_id
	CancelRequestTriggerModelCustomerIDColumn = "customer_id"
	// CancelRequestTriggerModelIDColumn is the column json value id
	CancelRequestTriggerModelIDColumn = "id"
	// CancelRequestTriggerModelRefIDColumn is the column json value ref_id
	CancelRequestTriggerModelRefIDColumn = "ref_id"
	// CancelRequestTriggerModelRefTypeColumn is the column json value ref_type
	CancelRequestTriggerModelRefTypeColumn = "ref_type"
	// CancelRequestTriggerModelUUIDColumn is the column json value uuid
	CancelRequestTriggerModelUUIDColumn = "uuid"
)

// CancelRequestTriggerCommand is the enumeration type for command
type CancelRequestTriggerCommand int32

// UnmarshalBSONValue for unmarshaling value
func (v *CancelRequestTriggerCommand) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = CancelRequestTriggerCommand(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "EXPORT":
			*v = CancelRequestTriggerCommand(0)
		case "ONBOARD":
			*v = CancelRequestTriggerCommand(1)
		case "INTEGRATION":
			*v = CancelRequestTriggerCommand(2)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v CancelRequestTriggerCommand) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "EXPORT":
		v = 0
	case "ONBOARD":
		v = 1
	case "INTEGRATION":
		v = 2
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v CancelRequestTriggerCommand) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("EXPORT")
	case 1:
		return json.Marshal("ONBOARD")
	case 2:
		return json.Marshal("INTEGRATION")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for Command
func (v CancelRequestTriggerCommand) String() string {
	switch int32(v) {
	case 0:
		return "EXPORT"
	case 1:
		return "ONBOARD"
	case 2:
		return "INTEGRATION"
	}
	return "unset"
}

const (
	// CancelRequestTriggerCommandExport is the enumeration value for export
	CancelRequestTriggerCommandExport CancelRequestTriggerCommand = 0
	// CancelRequestTriggerCommandOnboard is the enumeration value for onboard
	CancelRequestTriggerCommandOnboard CancelRequestTriggerCommand = 1
	// CancelRequestTriggerCommandIntegration is the enumeration value for integration
	CancelRequestTriggerCommandIntegration CancelRequestTriggerCommand = 2
)

// CancelRequestTrigger used to trigger agent.CancelRequest
type CancelRequestTrigger struct {
	// Command The command to cancel
	Command CancelRequestTriggerCommand `json:"command" codec:"command" bson:"command" yaml:"command" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// UUID The agent UUID
	UUID string `json:"uuid" codec:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*CancelRequestTrigger)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*CancelRequestTrigger)(nil)

func toCancelRequestTriggerObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CancelRequestTrigger:
		return v.ToMap()

	case CancelRequestTriggerCommand:
		return v.String()

	default:
		return o
	}
}

// String returns a string representation of CancelRequestTrigger
func (o *CancelRequestTrigger) String() string {
	return fmt.Sprintf("agent.CancelRequestTrigger<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *CancelRequestTrigger) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *CancelRequestTrigger) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *CancelRequestTrigger) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *CancelRequestTrigger) GetModelName() datamodel.ModelNameType {
	return CancelRequestTriggerModelName
}

// NewCancelRequestTriggerID provides a template for generating an ID field for CancelRequestTrigger
func NewCancelRequestTriggerID(customerID string, refType string, refID string) string {
	return hash.Values("CancelRequestTrigger", customerID, refType, refID)
}

func (o *CancelRequestTrigger) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("CancelRequestTrigger", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *CancelRequestTrigger) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *CancelRequestTrigger) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *CancelRequestTrigger) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *CancelRequestTrigger) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *CancelRequestTrigger) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *CancelRequestTrigger) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *CancelRequestTrigger) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *CancelRequestTrigger) IsEvented() bool {
	return false
}

// GetTopicConfig returns the topic config object
func (o *CancelRequestTrigger) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *CancelRequestTrigger) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of CancelRequestTrigger
func (o *CancelRequestTrigger) Clone() datamodel.Model {
	c := new(CancelRequestTrigger)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *CancelRequestTrigger) Anon() datamodel.Model {
	c := new(CancelRequestTrigger)
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
func (o *CancelRequestTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *CancelRequestTrigger) UnmarshalJSON(data []byte) error {
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
func (o *CancelRequestTrigger) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *CancelRequestTrigger) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two CancelRequestTrigger objects are equal
func (o *CancelRequestTrigger) IsEqual(other *CancelRequestTrigger) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *CancelRequestTrigger) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{

		"command":     o.Command.String(),
		"customer_id": toCancelRequestTriggerObject(o.CustomerID, false),
		"id":          toCancelRequestTriggerObject(o.ID, false),
		"ref_id":      toCancelRequestTriggerObject(o.RefID, false),
		"ref_type":    toCancelRequestTriggerObject(o.RefType, false),
		"uuid":        toCancelRequestTriggerObject(o.UUID, false),
		"hashcode":    toCancelRequestTriggerObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *CancelRequestTrigger) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["command"].(CancelRequestTriggerCommand); ok {
		o.Command = val
	} else {
		if em, ok := kv["command"].(map[string]interface{}); ok {

			ev := em["agent.command"].(string)
			switch ev {
			case "export", "EXPORT":
				o.Command = 0
			case "onboard", "ONBOARD":
				o.Command = 1
			case "integration", "INTEGRATION":
				o.Command = 2
			}
		}
		if em, ok := kv["command"].(string); ok {
			switch em {
			case "export", "EXPORT":
				o.Command = 0
			case "onboard", "ONBOARD":
				o.Command = 1
			case "integration", "INTEGRATION":
				o.Command = 2
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
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *CancelRequestTrigger) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Command)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.UUID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *CancelRequestTrigger) GetEventAPIConfig() datamodel.EventAPIConfig {
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
