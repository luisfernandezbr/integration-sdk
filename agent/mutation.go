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
	pstrings "github.com/pinpt/go-common/v10/strings"
)

const (

	// MutationTable is the default table name
	MutationTable datamodel.ModelNameType = "agent_mutation"

	// MutationModelName is the model name
	MutationModelName datamodel.ModelNameType = "agent.Mutation"
)

const (
	// MutationModelCustomerIDColumn is the column json value customer_id
	MutationModelCustomerIDColumn = "customer_id"
	// MutationModelIDColumn is the column json value id
	MutationModelIDColumn = "id"
	// MutationModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	MutationModelIntegrationInstanceIDColumn = "integration_instance_id"
	// MutationModelPayloadColumn is the column json value payload
	MutationModelPayloadColumn = "payload"
	// MutationModelRefIDColumn is the column json value ref_id
	MutationModelRefIDColumn = "ref_id"
	// MutationModelRefTypeColumn is the column json value ref_type
	MutationModelRefTypeColumn = "ref_type"
)

// Mutation request to change data in integration going from agent service to agent
type Mutation struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// Payload the mutation body.
	Payload string `json:"payload" codec:"payload" bson:"payload" yaml:"payload" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Mutation)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*Mutation)(nil)

func toMutationObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *Mutation:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of Mutation
func (o *Mutation) String() string {
	return fmt.Sprintf("agent.Mutation<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Mutation) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *Mutation) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *Mutation) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *Mutation) GetModelName() datamodel.ModelNameType {
	return MutationModelName
}

// NewMutationID provides a template for generating an ID field for Mutation
func NewMutationID(customerID string) string {
	return hash.Values(customerID, datetime.EpochNow(), randomString(64))
}

func (o *Mutation) setDefaults(frommap bool) {

	if o.ID == "" {
		o.ID = hash.Values(o.CustomerID, datetime.EpochNow(), randomString(64))
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Mutation) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Mutation) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Mutation) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *Mutation) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Mutation) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *Mutation) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Mutation) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Mutation) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *Mutation) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = MutationModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Mutation) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *Mutation) GetCustomerID() string {
	return o.CustomerID
}

// SetCustomerID will return the customer_id
func (o *Mutation) SetCustomerID(id string) {
	o.CustomerID = id
}

// GetRefType will return the ref_type
func (o *Mutation) GetRefType() string {
	return o.RefType
}

// SetRefType will return the ref_type
func (o *Mutation) SetRefType(t string) {
	o.RefType = t
}

// Clone returns an exact copy of Mutation
func (o *Mutation) Clone() datamodel.Model {
	c := new(Mutation)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Mutation) Anon() datamodel.Model {
	c := new(Mutation)
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
func (o *Mutation) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Mutation) UnmarshalJSON(data []byte) error {
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
func (o *Mutation) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *Mutation) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two Mutation objects are equal
func (o *Mutation) IsEqual(other *Mutation) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Mutation) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"customer_id":             toMutationObject(o.CustomerID, false),
		"id":                      toMutationObject(o.ID, false),
		"integration_instance_id": toMutationObject(o.IntegrationInstanceID, true),
		"payload":                 toMutationObject(o.Payload, false),
		"ref_id":                  toMutationObject(o.RefID, false),
		"ref_type":                toMutationObject(o.RefType, false),
		"hashcode":                toMutationObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *Mutation) FromMap(kv map[string]interface{}) {

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
	if val, ok := kv["payload"].(string); ok {
		o.Payload = val
	} else {
		if val, ok := kv["payload"]; ok {
			if val == nil {
				o.Payload = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Payload = fmt.Sprintf("%v", val)
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
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *Mutation) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.Payload)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// SetIntegrationInstanceID will set the integration instance ID
func (o *Mutation) SetIntegrationInstanceID(id string) {
	if id == "" {
		o.IntegrationInstanceID = nil
	} else {
		o.IntegrationInstanceID = &id
	}
}

// GetIntegrationInstanceID will return the integration instance ID
func (o *Mutation) GetIntegrationInstanceID() *string {
	return o.IntegrationInstanceID
}

// MutationPartial is a partial struct for upsert mutations for Mutation
type MutationPartial struct {
	// Payload the mutation body.
	Payload *string `json:"payload,omitempty"`
}

var _ datamodel.PartialModel = (*MutationPartial)(nil)

// GetModelName returns the name of the model
func (o *MutationPartial) GetModelName() datamodel.ModelNameType {
	return MutationModelName
}

// ToMap returns the object as a map
func (o *MutationPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"payload": toMutationObject(o.Payload, true),
	}
	for k, v := range kv {
		if v == nil || reflect.ValueOf(v).IsZero() {
			delete(kv, k)
		} else {
		}
	}
	return kv
}

// Stringify returns the object in JSON format as a string
func (o *MutationPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *MutationPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *MutationPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *MutationPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *MutationPartial) FromMap(kv map[string]interface{}) {
	if val, ok := kv["payload"].(*string); ok {
		o.Payload = val
	} else if val, ok := kv["payload"].(string); ok {
		o.Payload = &val
	} else {
		if val, ok := kv["payload"]; ok {
			if val == nil {
				o.Payload = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Payload = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	o.setDefaults(false)
}
