// DO NOT EDIT -- generated code

// Package codequality - the system which contains code quality
package codequality

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

	// ProjectTable is the default table name
	ProjectTable datamodel.ModelNameType = "codequality_project"

	// ProjectModelName is the model name
	ProjectModelName datamodel.ModelNameType = "codequality.Project"
)

const (
	// ProjectModelCustomerIDColumn is the column json value customer_id
	ProjectModelCustomerIDColumn = "customer_id"
	// ProjectModelIDColumn is the column json value id
	ProjectModelIDColumn = "id"
	// ProjectModelIdentifierColumn is the column json value identifier
	ProjectModelIdentifierColumn = "identifier"
	// ProjectModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	ProjectModelIntegrationInstanceIDColumn = "integration_instance_id"
	// ProjectModelNameColumn is the column json value name
	ProjectModelNameColumn = "name"
	// ProjectModelRefIDColumn is the column json value ref_id
	ProjectModelRefIDColumn = "ref_id"
	// ProjectModelRefTypeColumn is the column json value ref_type
	ProjectModelRefTypeColumn = "ref_type"
)

// Project project information
type Project struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// Identifier the common identifier of the project
	Identifier string `json:"identifier" codec:"identifier" bson:"identifier" yaml:"identifier" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// Name the name of the project
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Project)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*Project)(nil)

func toProjectObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *Project:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of Project
func (o *Project) String() string {
	return fmt.Sprintf("codequality.Project<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Project) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *Project) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *Project) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *Project) GetModelName() datamodel.ModelNameType {
	return ProjectModelName
}

// NewProjectID provides a template for generating an ID field for Project
func NewProjectID(customerID string, refType string, refID string) string {
	return hash.Values("Project", customerID, refType, refID)
}

func (o *Project) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Project", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Project) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Project) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Project) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *Project) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Project) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *Project) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Project) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Project) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *Project) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = ProjectModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Project) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *Project) GetCustomerID() string {

	return o.CustomerID

}

// SetCustomerID will return the customer_id
func (o *Project) SetCustomerID(id string) {

	o.CustomerID = id

}

// GetRefType will return the ref_type
func (o *Project) GetRefType() string {
	return o.RefType
}

// SetRefType will return the ref_type
func (o *Project) SetRefType(t string) {
	o.RefType = t
}

// Clone returns an exact copy of Project
func (o *Project) Clone() datamodel.Model {
	c := new(Project)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Project) Anon() datamodel.Model {
	c := new(Project)
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
func (o *Project) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Project) UnmarshalJSON(data []byte) error {
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
func (o *Project) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *Project) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two Project objects are equal
func (o *Project) IsEqual(other *Project) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Project) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"customer_id":             toProjectObject(o.CustomerID, false),
		"id":                      toProjectObject(o.ID, false),
		"identifier":              toProjectObject(o.Identifier, false),
		"integration_instance_id": toProjectObject(o.IntegrationInstanceID, true),
		"name":                    toProjectObject(o.Name, false),
		"ref_id":                  toProjectObject(o.RefID, false),
		"ref_type":                toProjectObject(o.RefType, false),
		"hashcode":                toProjectObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *Project) FromMap(kv map[string]interface{}) {

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
	if val, ok := kv["identifier"].(string); ok {
		o.Identifier = val
	} else {
		if val, ok := kv["identifier"]; ok {
			if val == nil {
				o.Identifier = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Identifier = fmt.Sprintf("%v", val)
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
	if val, ok := kv["name"].(string); ok {
		o.Name = val
	} else {
		if val, ok := kv["name"]; ok {
			if val == nil {
				o.Name = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Name = fmt.Sprintf("%v", val)
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
func (o *Project) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.Identifier)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.Name)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// SetIntegrationInstanceID will set the integration instance ID
func (o *Project) SetIntegrationInstanceID(id string) {
	if id == "" {
		o.IntegrationInstanceID = nil
	} else {
		o.IntegrationInstanceID = &id
	}
}

// GetIntegrationInstanceID will return the integration instance ID
func (o *Project) GetIntegrationInstanceID() *string {
	return o.IntegrationInstanceID
}

// ProjectPartial is a partial struct for upsert mutations for Project
type ProjectPartial struct {
	// Identifier the common identifier of the project
	Identifier *string `json:"identifier,omitempty"`
	// Name the name of the project
	Name *string `json:"name,omitempty"`
}

var _ datamodel.PartialModel = (*ProjectPartial)(nil)

// GetModelName returns the name of the model
func (o *ProjectPartial) GetModelName() datamodel.ModelNameType {
	return ProjectModelName
}

// ToMap returns the object as a map
func (o *ProjectPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"identifier": toProjectObject(o.Identifier, true),
		"name":       toProjectObject(o.Name, true),
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
func (o *ProjectPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *ProjectPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *ProjectPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *ProjectPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *ProjectPartial) FromMap(kv map[string]interface{}) {
	if val, ok := kv["identifier"].(*string); ok {
		o.Identifier = val
	} else if val, ok := kv["identifier"].(string); ok {
		o.Identifier = &val
	} else {
		if val, ok := kv["identifier"]; ok {
			if val == nil {
				o.Identifier = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Identifier = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["name"].(*string); ok {
		o.Name = val
	} else if val, ok := kv["name"].(string); ok {
		o.Name = &val
	} else {
		if val, ok := kv["name"]; ok {
			if val == nil {
				o.Name = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Name = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	o.setDefaults(false)
}
