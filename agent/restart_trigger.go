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
)

const (

	// RestartTriggerTable is the default table name
	RestartTriggerTable datamodel.ModelNameType = "agent_restarttrigger"

	// RestartTriggerModelName is the model name
	RestartTriggerModelName datamodel.ModelNameType = "agent.RestartTrigger"
)

const (
	// RestartTriggerModelCustomerIDColumn is the column json value customer_id
	RestartTriggerModelCustomerIDColumn = "customer_id"
	// RestartTriggerModelIDColumn is the column json value id
	RestartTriggerModelIDColumn = "id"
	// RestartTriggerModelRefIDColumn is the column json value ref_id
	RestartTriggerModelRefIDColumn = "ref_id"
	// RestartTriggerModelRefTypeColumn is the column json value ref_type
	RestartTriggerModelRefTypeColumn = "ref_type"
	// RestartTriggerModelUUIDColumn is the column json value uuid
	RestartTriggerModelUUIDColumn = "uuid"
)

// RestartTrigger used to make the agent service restart
type RestartTrigger struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// UUID This UUID of the agent to trigger
	UUID string `json:"uuid" codec:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*RestartTrigger)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*RestartTrigger)(nil)

func toRestartTriggerObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *RestartTrigger:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of RestartTrigger
func (o *RestartTrigger) String() string {
	return fmt.Sprintf("agent.RestartTrigger<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *RestartTrigger) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *RestartTrigger) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *RestartTrigger) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *RestartTrigger) GetModelName() datamodel.ModelNameType {
	return RestartTriggerModelName
}

// NewRestartTriggerID provides a template for generating an ID field for RestartTrigger
func NewRestartTriggerID(customerID string, refType string, refID string) string {
	return hash.Values("RestartTrigger", customerID, refType, refID)
}

func (o *RestartTrigger) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("RestartTrigger", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *RestartTrigger) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *RestartTrigger) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *RestartTrigger) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *RestartTrigger) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *RestartTrigger) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *RestartTrigger) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *RestartTrigger) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *RestartTrigger) IsEvented() bool {
	return false
}

// GetTopicConfig returns the topic config object
func (o *RestartTrigger) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *RestartTrigger) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of RestartTrigger
func (o *RestartTrigger) Clone() datamodel.Model {
	c := new(RestartTrigger)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *RestartTrigger) Anon() datamodel.Model {
	c := new(RestartTrigger)
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
func (o *RestartTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *RestartTrigger) UnmarshalJSON(data []byte) error {
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
func (o *RestartTrigger) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *RestartTrigger) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two RestartTrigger objects are equal
func (o *RestartTrigger) IsEqual(other *RestartTrigger) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *RestartTrigger) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"customer_id": toRestartTriggerObject(o.CustomerID, false),
		"id":          toRestartTriggerObject(o.ID, false),
		"ref_id":      toRestartTriggerObject(o.RefID, false),
		"ref_type":    toRestartTriggerObject(o.RefType, false),
		"uuid":        toRestartTriggerObject(o.UUID, false),
		"hashcode":    toRestartTriggerObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *RestartTrigger) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
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
func (o *RestartTrigger) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.UUID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}
