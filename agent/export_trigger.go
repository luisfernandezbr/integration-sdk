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
	"github.com/pinpt/go-common/v10/number"
	pstrings "github.com/pinpt/go-common/v10/strings"
)

const (

	// ExportTriggerTable is the default table name
	ExportTriggerTable datamodel.ModelNameType = "agent_exporttrigger"

	// ExportTriggerModelName is the model name
	ExportTriggerModelName datamodel.ModelNameType = "agent.ExportTrigger"
)

const (
	// ExportTriggerModelCustomerIDColumn is the column json value customer_id
	ExportTriggerModelCustomerIDColumn = "customer_id"
	// ExportTriggerModelIDColumn is the column json value id
	ExportTriggerModelIDColumn = "id"
	// ExportTriggerModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	ExportTriggerModelIntegrationInstanceIDColumn = "integration_instance_id"
	// ExportTriggerModelRefIDColumn is the column json value ref_id
	ExportTriggerModelRefIDColumn = "ref_id"
	// ExportTriggerModelRefTypeColumn is the column json value ref_type
	ExportTriggerModelRefTypeColumn = "ref_type"
	// ExportTriggerModelReprocessHistoricalColumn is the column json value reprocess_historical
	ExportTriggerModelReprocessHistoricalColumn = "reprocess_historical"
	// ExportTriggerModelUUIDColumn is the column json value uuid
	ExportTriggerModelUUIDColumn = "uuid"
)

// ExportTrigger used to trigger an agent.ExportRequest
type ExportTrigger struct {
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
	// ReprocessHistorical This field is to differentiate between historical and incrementals
	ReprocessHistorical bool `json:"reprocess_historical" codec:"reprocess_historical" bson:"reprocess_historical" yaml:"reprocess_historical" faker:"-"`
	// UUID This UUID of the agent to trigger
	UUID *string `json:"uuid,omitempty" codec:"uuid,omitempty" bson:"uuid" yaml:"uuid,omitempty" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*ExportTrigger)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*ExportTrigger)(nil)

func toExportTriggerObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ExportTrigger:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of ExportTrigger
func (o *ExportTrigger) String() string {
	return fmt.Sprintf("agent.ExportTrigger<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *ExportTrigger) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *ExportTrigger) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *ExportTrigger) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *ExportTrigger) GetModelName() datamodel.ModelNameType {
	return ExportTriggerModelName
}

// NewExportTriggerID provides a template for generating an ID field for ExportTrigger
func NewExportTriggerID(customerID string, refType string, refID string) string {
	return hash.Values("ExportTrigger", customerID, refType, refID)
}

func (o *ExportTrigger) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("ExportTrigger", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *ExportTrigger) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *ExportTrigger) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *ExportTrigger) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *ExportTrigger) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *ExportTrigger) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *ExportTrigger) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *ExportTrigger) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *ExportTrigger) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *ExportTrigger) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = ExportTriggerModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *ExportTrigger) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *ExportTrigger) GetCustomerID() string {

	return o.CustomerID

}

// SetCustomerID will return the customer_id
func (o *ExportTrigger) SetCustomerID(id string) {

	o.CustomerID = id

}

// GetRefType will return the ref_type
func (o *ExportTrigger) GetRefType() string {

	return o.RefType

}

// SetRefType will return the ref_type
func (o *ExportTrigger) SetRefType(t string) {

	o.RefType = t

}

// Clone returns an exact copy of ExportTrigger
func (o *ExportTrigger) Clone() datamodel.Model {
	c := new(ExportTrigger)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *ExportTrigger) Anon() datamodel.Model {
	c := new(ExportTrigger)
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
func (o *ExportTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *ExportTrigger) UnmarshalJSON(data []byte) error {
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
func (o *ExportTrigger) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *ExportTrigger) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two ExportTrigger objects are equal
func (o *ExportTrigger) IsEqual(other *ExportTrigger) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *ExportTrigger) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"customer_id":             toExportTriggerObject(o.CustomerID, false),
		"id":                      toExportTriggerObject(o.ID, false),
		"integration_instance_id": toExportTriggerObject(o.IntegrationInstanceID, true),
		"ref_id":                  toExportTriggerObject(o.RefID, false),
		"ref_type":                toExportTriggerObject(o.RefType, false),
		"reprocess_historical":    toExportTriggerObject(o.ReprocessHistorical, false),
		"uuid":                    toExportTriggerObject(o.UUID, true),
		"hashcode":                toExportTriggerObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *ExportTrigger) FromMap(kv map[string]interface{}) {

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
	if val, ok := kv["reprocess_historical"].(bool); ok {
		o.ReprocessHistorical = val
	} else {
		if val, ok := kv["reprocess_historical"]; ok {
			if val == nil {
				o.ReprocessHistorical = false
			} else {
				o.ReprocessHistorical = number.ToBoolAny(val)
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
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *ExportTrigger) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.ReprocessHistorical)
	args = append(args, o.UUID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// SetIntegrationInstanceID will set the integration instance ID
func (o *ExportTrigger) SetIntegrationInstanceID(id string) {
	if id == "" {
		o.IntegrationInstanceID = nil
	} else {
		o.IntegrationInstanceID = &id
	}
}

// GetIntegrationInstanceID will return the integration instance ID
func (o *ExportTrigger) GetIntegrationInstanceID() *string {
	return o.IntegrationInstanceID
}

// ExportTriggerPartial is a partial struct for upsert mutations for ExportTrigger
type ExportTriggerPartial struct {
	// ReprocessHistorical This field is to differentiate between historical and incrementals
	ReprocessHistorical *bool `json:"reprocess_historical,omitempty"`
	// UUID This UUID of the agent to trigger
	UUID *string `json:"uuid,omitempty"`
}

var _ datamodel.PartialModel = (*ExportTriggerPartial)(nil)

// GetModelName returns the name of the model
func (o *ExportTriggerPartial) GetModelName() datamodel.ModelNameType {
	return ExportTriggerModelName
}

// ToMap returns the object as a map
func (o *ExportTriggerPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"reprocess_historical": toExportTriggerObject(o.ReprocessHistorical, true),
		"uuid":                 toExportTriggerObject(o.UUID, true),
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
func (o *ExportTriggerPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *ExportTriggerPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *ExportTriggerPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *ExportTriggerPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *ExportTriggerPartial) FromMap(kv map[string]interface{}) {
	if val, ok := kv["reprocess_historical"].(*bool); ok {
		o.ReprocessHistorical = val
	} else if val, ok := kv["reprocess_historical"].(bool); ok {
		o.ReprocessHistorical = &val
	} else {
		if val, ok := kv["reprocess_historical"]; ok {
			if val == nil {
				o.ReprocessHistorical = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.ReprocessHistorical = number.BoolPointer(number.ToBoolAny(val))
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
	o.setDefaults(false)
}
