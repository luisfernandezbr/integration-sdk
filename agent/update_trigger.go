// DO NOT EDIT -- generated code

// Package agent - the agent communicates with pinpoint cloud to send integration data
package agent

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/bxcodec/faker"
	"github.com/pinpt/go-common/v10/datamodel"
	"github.com/pinpt/go-common/v10/hash"
	pjson "github.com/pinpt/go-common/v10/json"
	pstrings "github.com/pinpt/go-common/v10/strings"
)

const (

	// UpdateTriggerTable is the default table name
	UpdateTriggerTable datamodel.ModelNameType = "agent_updatetrigger"

	// UpdateTriggerModelName is the model name
	UpdateTriggerModelName datamodel.ModelNameType = "agent.UpdateTrigger"
)

const (
	// UpdateTriggerModelCustomerIDColumn is the column json value customer_id
	UpdateTriggerModelCustomerIDColumn = "customer_id"
	// UpdateTriggerModelIDColumn is the column json value id
	UpdateTriggerModelIDColumn = "id"
	// UpdateTriggerModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	UpdateTriggerModelIntegrationInstanceIDColumn = "integration_instance_id"
	// UpdateTriggerModelRefIDColumn is the column json value ref_id
	UpdateTriggerModelRefIDColumn = "ref_id"
	// UpdateTriggerModelRefTypeColumn is the column json value ref_type
	UpdateTriggerModelRefTypeColumn = "ref_type"
	// UpdateTriggerModelUUIDColumn is the column json value uuid
	UpdateTriggerModelUUIDColumn = "uuid"
	// UpdateTriggerModelVersionColumn is the column json value version
	UpdateTriggerModelVersionColumn = "version"
)

// UpdateTrigger used to trigger an agent.UpdateRequest
type UpdateTrigger struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// UUID UUID of the agent to update
	UUID *string `json:"uuid,omitempty" codec:"uuid,omitempty" bson:"uuid" yaml:"uuid,omitempty" faker:"-"`
	// Version
	Version *string `json:"version,omitempty" codec:"version,omitempty" bson:"version" yaml:"version,omitempty" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*UpdateTrigger)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*UpdateTrigger)(nil)

func toUpdateTriggerObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *UpdateTrigger:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of UpdateTrigger
func (o *UpdateTrigger) String() string {
	return fmt.Sprintf("agent.UpdateTrigger<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *UpdateTrigger) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *UpdateTrigger) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *UpdateTrigger) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *UpdateTrigger) GetModelName() datamodel.ModelNameType {
	return UpdateTriggerModelName
}

// NewUpdateTriggerID provides a template for generating an ID field for UpdateTrigger
func NewUpdateTriggerID(customerID string, refType string, refID string) string {
	return hash.Values("UpdateTrigger", customerID, refType, refID)
}

func (o *UpdateTrigger) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("UpdateTrigger", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *UpdateTrigger) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *UpdateTrigger) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *UpdateTrigger) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *UpdateTrigger) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *UpdateTrigger) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *UpdateTrigger) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *UpdateTrigger) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *UpdateTrigger) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *UpdateTrigger) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = UpdateTriggerModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *UpdateTrigger) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *UpdateTrigger) GetCustomerID() string {

	return o.CustomerID

}

// SetCustomerID will return the customer_id
func (o *UpdateTrigger) SetCustomerID(id string) {

	o.CustomerID = id

}

// GetRefType will return the ref_type
func (o *UpdateTrigger) GetRefType() string {
	return o.RefType
}

// SetRefType will return the ref_type
func (o *UpdateTrigger) SetRefType(t string) {
	o.RefType = t
}

// Clone returns an exact copy of UpdateTrigger
func (o *UpdateTrigger) Clone() datamodel.Model {
	c := new(UpdateTrigger)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *UpdateTrigger) Anon() datamodel.Model {
	c := new(UpdateTrigger)
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
func (o *UpdateTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *UpdateTrigger) UnmarshalJSON(data []byte) error {
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
func (o *UpdateTrigger) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *UpdateTrigger) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two UpdateTrigger objects are equal
func (o *UpdateTrigger) IsEqual(other *UpdateTrigger) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *UpdateTrigger) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"customer_id":             toUpdateTriggerObject(o.CustomerID, false),
		"id":                      toUpdateTriggerObject(o.ID, false),
		"integration_instance_id": toUpdateTriggerObject(o.IntegrationInstanceID, true),
		"ref_id":                  toUpdateTriggerObject(o.RefID, false),
		"ref_type":                toUpdateTriggerObject(o.RefType, false),
		"uuid":                    toUpdateTriggerObject(o.UUID, true),
		"version":                 toUpdateTriggerObject(o.Version, true),
		"hashcode":                toUpdateTriggerObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *UpdateTrigger) FromMap(kv map[string]interface{}) {

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
	if val, ok := kv["uuid"].(*string); ok {
		o.UUID = val
	} else if val, ok := kv["uuid"].(string); ok {
		o.UUID = &val
	} else {
		if val, ok := kv["uuid"]; ok {
			if val == nil {
				o.UUID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.UUID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["version"].(*string); ok {
		o.Version = val
	} else if val, ok := kv["version"].(string); ok {
		o.Version = &val
	} else {
		if val, ok := kv["version"]; ok {
			if val == nil {
				o.Version = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Version = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *UpdateTrigger) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.UUID)
	args = append(args, o.Version)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// SetIntegrationInstanceID will set the integration instance ID
func (o *UpdateTrigger) SetIntegrationInstanceID(id string) {
	if id == "" {
		o.IntegrationInstanceID = nil
	} else {
		o.IntegrationInstanceID = &id
	}
}

// GetIntegrationInstanceID will return the integration instance ID
func (o *UpdateTrigger) GetIntegrationInstanceID() *string {
	return o.IntegrationInstanceID
}

// UpdateTriggerPartial is a partial struct for upsert mutations for UpdateTrigger
type UpdateTriggerPartial struct {
	// UUID UUID of the agent to update
	UUID *string `json:"uuid,omitempty"`
	// Version
	Version *string `json:"version,omitempty"`
}

var _ datamodel.PartialModel = (*UpdateTriggerPartial)(nil)

// GetModelName returns the name of the model
func (o *UpdateTriggerPartial) GetModelName() datamodel.ModelNameType {
	return UpdateTriggerModelName
}

// ToMap returns the object as a map
func (o *UpdateTriggerPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"uuid":    toUpdateTriggerObject(o.UUID, true),
		"version": toUpdateTriggerObject(o.Version, true),
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
func (o *UpdateTriggerPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *UpdateTriggerPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *UpdateTriggerPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *UpdateTriggerPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *UpdateTriggerPartial) FromMap(kv map[string]interface{}) {
	if val, ok := kv["uuid"].(*string); ok {
		o.UUID = val
	} else if val, ok := kv["uuid"].(string); ok {
		o.UUID = &val
	} else {
		if val, ok := kv["uuid"]; ok {
			if val == nil {
				o.UUID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.UUID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["version"].(*string); ok {
		o.Version = val
	} else if val, ok := kv["version"].(string); ok {
		o.Version = &val
	} else {
		if val, ok := kv["version"]; ok {
			if val == nil {
				o.Version = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Version = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	o.setDefaults(false)
}
