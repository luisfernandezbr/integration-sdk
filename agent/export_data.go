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
	"github.com/pinpt/go-common/v10/hash"
	pjson "github.com/pinpt/go-common/v10/json"
	pstrings "github.com/pinpt/go-common/v10/strings"
)

const (

	// ExportDataTable is the default table name
	ExportDataTable datamodel.ModelNameType = "agent_exportdata"

	// ExportDataModelName is the model name
	ExportDataModelName datamodel.ModelNameType = "agent.ExportData"
)

const (
	// ExportDataModelCustomerIDColumn is the column json value customer_id
	ExportDataModelCustomerIDColumn = "customer_id"
	// ExportDataModelIDColumn is the column json value id
	ExportDataModelIDColumn = "id"
	// ExportDataModelIntegrationIDColumn is the column json value integration_id
	ExportDataModelIntegrationIDColumn = "integration_id"
	// ExportDataModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	ExportDataModelIntegrationInstanceIDColumn = "integration_instance_id"
	// ExportDataModelJobIDColumn is the column json value job_id
	ExportDataModelJobIDColumn = "job_id"
	// ExportDataModelObjectsColumn is the column json value objects
	ExportDataModelObjectsColumn = "objects"
	// ExportDataModelRefIDColumn is the column json value ref_id
	ExportDataModelRefIDColumn = "ref_id"
	// ExportDataModelRefTypeColumn is the column json value ref_type
	ExportDataModelRefTypeColumn = "ref_type"
)

// ExportData export data from a streaming export
type ExportData struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationID The ID of the integration
	IntegrationID string `json:"integration_id" codec:"integration_id" bson:"integration_id" yaml:"integration_id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// JobID The job ID
	JobID string `json:"job_id" codec:"job_id" bson:"job_id" yaml:"job_id" faker:"-"`
	// Objects Objects as JSON map[type][]object for use in pipeline
	Objects string `json:"objects" codec:"objects" bson:"objects" yaml:"objects" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*ExportData)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*ExportData)(nil)

func toExportDataObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ExportData:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of ExportData
func (o *ExportData) String() string {
	return fmt.Sprintf("agent.ExportData<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *ExportData) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *ExportData) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *ExportData) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *ExportData) GetModelName() datamodel.ModelNameType {
	return ExportDataModelName
}

// NewExportDataID provides a template for generating an ID field for ExportData
func NewExportDataID(customerID string, refType string, refID string) string {
	return hash.Values("ExportData", customerID, refType, refID)
}

func (o *ExportData) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("ExportData", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *ExportData) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *ExportData) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *ExportData) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *ExportData) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *ExportData) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *ExportData) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *ExportData) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *ExportData) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *ExportData) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = ExportDataModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *ExportData) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *ExportData) GetCustomerID() string {
	return o.CustomerID
}

// SetCustomerID will return the customer_id
func (o *ExportData) SetCustomerID(id string) {
	o.CustomerID = id
}

// GetRefType will return the ref_type
func (o *ExportData) GetRefType() string {
	return o.RefType
}

// SetRefType will return the ref_type
func (o *ExportData) SetRefType(t string) {
	o.RefType = t
}

// Clone returns an exact copy of ExportData
func (o *ExportData) Clone() datamodel.Model {
	c := new(ExportData)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *ExportData) Anon() datamodel.Model {
	c := new(ExportData)
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
func (o *ExportData) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *ExportData) UnmarshalJSON(data []byte) error {
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
func (o *ExportData) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *ExportData) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two ExportData objects are equal
func (o *ExportData) IsEqual(other *ExportData) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *ExportData) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"customer_id":             toExportDataObject(o.CustomerID, false),
		"id":                      toExportDataObject(o.ID, false),
		"integration_id":          toExportDataObject(o.IntegrationID, false),
		"integration_instance_id": toExportDataObject(o.IntegrationInstanceID, true),
		"job_id":                  toExportDataObject(o.JobID, false),
		"objects":                 toExportDataObject(o.Objects, false),
		"ref_id":                  toExportDataObject(o.RefID, false),
		"ref_type":                toExportDataObject(o.RefType, false),
		"hashcode":                toExportDataObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *ExportData) FromMap(kv map[string]interface{}) {

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
	if val, ok := kv["integration_id"].(string); ok {
		o.IntegrationID = val
	} else {
		if val, ok := kv["integration_id"]; ok {
			if val == nil {
				o.IntegrationID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.IntegrationID = fmt.Sprintf("%v", val)
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
	if val, ok := kv["objects"].(string); ok {
		o.Objects = val
	} else {
		if val, ok := kv["objects"]; ok {
			if val == nil {
				o.Objects = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Objects = fmt.Sprintf("%v", val)
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
func (o *ExportData) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.IntegrationID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.JobID)
	args = append(args, o.Objects)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// SetIntegrationInstanceID will set the integration instance ID
func (o *ExportData) SetIntegrationInstanceID(id string) {
	if id == "" {
		o.IntegrationInstanceID = nil
	} else {
		o.IntegrationInstanceID = &id
	}
}

// GetIntegrationInstanceID will return the integration instance ID
func (o *ExportData) GetIntegrationInstanceID() *string {
	return o.IntegrationInstanceID
}

// ExportDataPartial is a partial struct for upsert mutations for ExportData
type ExportDataPartial struct {
	// IntegrationID The ID of the integration
	IntegrationID *string `json:"integration_id,omitempty"`
	// JobID The job ID
	JobID *string `json:"job_id,omitempty"`
	// Objects Objects as JSON map[type][]object for use in pipeline
	Objects *string `json:"objects,omitempty"`
}

var _ datamodel.PartialModel = (*ExportDataPartial)(nil)

// GetModelName returns the name of the model
func (o *ExportDataPartial) GetModelName() datamodel.ModelNameType {
	return ExportDataModelName
}

// ToMap returns the object as a map
func (o *ExportDataPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"integration_id": toExportDataObject(o.IntegrationID, true),
		"job_id":         toExportDataObject(o.JobID, true),
		"objects":        toExportDataObject(o.Objects, true),
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
func (o *ExportDataPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *ExportDataPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *ExportDataPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *ExportDataPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *ExportDataPartial) FromMap(kv map[string]interface{}) {
	if val, ok := kv["integration_id"].(*string); ok {
		o.IntegrationID = val
	} else if val, ok := kv["integration_id"].(string); ok {
		o.IntegrationID = &val
	} else {
		if val, ok := kv["integration_id"]; ok {
			if val == nil {
				o.IntegrationID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.IntegrationID = pstrings.Pointer(fmt.Sprintf("%v", val))
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
	if val, ok := kv["objects"].(*string); ok {
		o.Objects = val
	} else if val, ok := kv["objects"].(string); ok {
		o.Objects = &val
	} else {
		if val, ok := kv["objects"]; ok {
			if val == nil {
				o.Objects = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Objects = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	o.setDefaults(false)
}
