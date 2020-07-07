// DO NOT EDIT -- generated code

// Package agent -
package agent

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/bxcodec/faker"
	"github.com/pinpt/go-common/v10/datamodel"
	"github.com/pinpt/go-common/v10/datetime"
	"github.com/pinpt/go-common/v10/hash"
	pjson "github.com/pinpt/go-common/v10/json"
	"github.com/pinpt/go-common/v10/slice"
	pstrings "github.com/pinpt/go-common/v10/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (

	// MutateDataTable is the default table name
	MutateDataTable datamodel.ModelNameType = "agent_mutatedata"

	// MutateDataModelName is the model name
	MutateDataModelName datamodel.ModelNameType = "agent.MutateData"
)

const (
	// MutateDataModelCustomerIDColumn is the column json value customer_id
	MutateDataModelCustomerIDColumn = "customer_id"
	// MutateDataModelIDColumn is the column json value id
	MutateDataModelIDColumn = "id"
	// MutateDataModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	MutateDataModelIntegrationInstanceIDColumn = "integration_instance_id"
	// MutateDataModelModelColumn is the column json value model
	MutateDataModelModelColumn = "model"
	// MutateDataModelPullColumn is the column json value pull
	MutateDataModelPullColumn = "pull"
	// MutateDataModelPushColumn is the column json value push
	MutateDataModelPushColumn = "push"
	// MutateDataModelRefIDColumn is the column json value ref_id
	MutateDataModelRefIDColumn = "ref_id"
	// MutateDataModelRefTypeColumn is the column json value ref_type
	MutateDataModelRefTypeColumn = "ref_type"
	// MutateDataModelSetColumn is the column json value set
	MutateDataModelSetColumn = "set"
	// MutateDataModelUnsetColumn is the column json value unset
	MutateDataModelUnsetColumn = "unset"
)

// MutateData request to change data on a per field basis vs the entire object
type MutateData struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the ID of the object to mutate. this object must exist or will not succeed
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// Model the name of the model
	Model string `json:"model" codec:"model" bson:"model" yaml:"model" faker:"-"`
	// Pull the fields to remove in the case an array like value where the value is the field name
	Pull []string `json:"pull" codec:"pull" bson:"pull" yaml:"pull" faker:"-"`
	// Push the fields to append in the case an array like value where the value is a serialized JSON value
	Push map[string]string `json:"push" codec:"push" bson:"push" yaml:"push" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// Set the fields to set where the value is a serialized JSON value
	Set map[string]string `json:"set" codec:"set" bson:"set" yaml:"set" faker:"-"`
	// Unset the field names to unset
	Unset []string `json:"unset" codec:"unset" bson:"unset" yaml:"unset" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*MutateData)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*MutateData)(nil)

func toMutateDataObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *MutateData:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of MutateData
func (o *MutateData) String() string {
	return fmt.Sprintf("agent.MutateData<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *MutateData) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *MutateData) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *MutateData) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *MutateData) GetModelName() datamodel.ModelNameType {
	return MutateDataModelName
}

// NewMutateDataID provides a template for generating an ID field for MutateData
func NewMutateDataID(customerID string) string {
	return hash.Values(customerID, datetime.EpochNow(), randomString(64))
}

func (o *MutateData) setDefaults(frommap bool) {
	if o.Pull == nil {
		o.Pull = make([]string, 0)
	}
	if o.Unset == nil {
		o.Unset = make([]string, 0)
	}

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
func (o *MutateData) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *MutateData) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *MutateData) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *MutateData) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *MutateData) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *MutateData) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *MutateData) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *MutateData) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *MutateData) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = MutateDataModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *MutateData) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *MutateData) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of MutateData
func (o *MutateData) Clone() datamodel.Model {
	c := new(MutateData)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *MutateData) Anon() datamodel.Model {
	c := new(MutateData)
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
func (o *MutateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *MutateData) UnmarshalJSON(data []byte) error {
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
func (o *MutateData) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *MutateData) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two MutateData objects are equal
func (o *MutateData) IsEqual(other *MutateData) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *MutateData) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"customer_id":             toMutateDataObject(o.CustomerID, false),
		"id":                      toMutateDataObject(o.ID, false),
		"integration_instance_id": toMutateDataObject(o.IntegrationInstanceID, true),
		"model":                   toMutateDataObject(o.Model, false),
		"pull":                    toMutateDataObject(o.Pull, false),
		"push":                    toMutateDataObject(o.Push, false),
		"ref_id":                  toMutateDataObject(o.RefID, false),
		"ref_type":                toMutateDataObject(o.RefType, false),
		"set":                     toMutateDataObject(o.Set, false),
		"unset":                   toMutateDataObject(o.Unset, false),
		"hashcode":                toMutateDataObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *MutateData) FromMap(kv map[string]interface{}) {

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
	if val, ok := kv["model"].(string); ok {
		o.Model = val
	} else {
		if val, ok := kv["model"]; ok {
			if val == nil {
				o.Model = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Model = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["pull"]; ok {
		if val != nil {
			na := make([]string, 0)
			if a, ok := val.([]string); ok {
				na = append(na, a...)
			} else {
				if a, ok := val.([]interface{}); ok {
					for _, ae := range a {
						if av, ok := ae.(string); ok {
							na = append(na, av)
						} else {
							if badMap, ok := ae.(map[interface{}]interface{}); ok {
								ae = slice.ConvertToStringToInterface(badMap)
							}
							b, _ := json.Marshal(ae)
							var av string
							if err := json.Unmarshal(b, &av); err != nil {
								panic("unsupported type for pull field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else if s, ok := val.(string); ok {
					for _, sv := range strings.Split(s, ",") {
						na = append(na, strings.TrimSpace(sv))
					}
				} else if a, ok := val.(primitive.A); ok {
					for _, ae := range a {
						if av, ok := ae.(string); ok {
							na = append(na, av)
						} else {
							b, _ := json.Marshal(ae)
							var av string
							if err := json.Unmarshal(b, &av); err != nil {
								panic("unsupported type for pull field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for pull field")
				}
			}
			o.Pull = na
		}
	}
	if o.Pull == nil {
		o.Pull = make([]string, 0)
	}
	if val, ok := kv["push"]; ok {
		if val != nil {
			kv := make(map[string]string)
			if m, ok := val.(map[string]string); ok {
				kv = m
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					for k, v := range m {
						if mv, ok := v.(string); ok {

							kv[k] = mv

						} else {
							kv[k] = pjson.Stringify(v)
						}
					}
				} else {
					panic("unsupported type for push field entry: " + reflect.TypeOf(val).String())
				}
			}
			o.Push = kv
		}
	}
	if o.Push == nil {
		// here
		o.Push = make(map[string]string)
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
	if val, ok := kv["set"]; ok {
		if val != nil {
			kv := make(map[string]string)
			if m, ok := val.(map[string]string); ok {
				kv = m
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					for k, v := range m {
						if mv, ok := v.(string); ok {

							kv[k] = mv

						} else {
							kv[k] = pjson.Stringify(v)
						}
					}
				} else {
					panic("unsupported type for set field entry: " + reflect.TypeOf(val).String())
				}
			}
			o.Set = kv
		}
	}
	if o.Set == nil {
		// here
		o.Set = make(map[string]string)
	}
	if val, ok := kv["unset"]; ok {
		if val != nil {
			na := make([]string, 0)
			if a, ok := val.([]string); ok {
				na = append(na, a...)
			} else {
				if a, ok := val.([]interface{}); ok {
					for _, ae := range a {
						if av, ok := ae.(string); ok {
							na = append(na, av)
						} else {
							if badMap, ok := ae.(map[interface{}]interface{}); ok {
								ae = slice.ConvertToStringToInterface(badMap)
							}
							b, _ := json.Marshal(ae)
							var av string
							if err := json.Unmarshal(b, &av); err != nil {
								panic("unsupported type for unset field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else if s, ok := val.(string); ok {
					for _, sv := range strings.Split(s, ",") {
						na = append(na, strings.TrimSpace(sv))
					}
				} else if a, ok := val.(primitive.A); ok {
					for _, ae := range a {
						if av, ok := ae.(string); ok {
							na = append(na, av)
						} else {
							b, _ := json.Marshal(ae)
							var av string
							if err := json.Unmarshal(b, &av); err != nil {
								panic("unsupported type for unset field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for unset field")
				}
			}
			o.Unset = na
		}
	}
	if o.Unset == nil {
		o.Unset = make([]string, 0)
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *MutateData) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.Model)
	args = append(args, o.Pull)
	args = append(args, o.Push)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.Set)
	args = append(args, o.Unset)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// MutateDataPartial is a partial struct for upsert mutations for MutateData
type MutateDataPartial struct {
	// ID the ID of the object to mutate. this object must exist or will not succeed
	ID *string `json:"id,omitempty"`
	// Model the name of the model
	Model *string `json:"model,omitempty"`
	// Pull the fields to remove in the case an array like value where the value is the field name
	Pull []string `json:"pull,omitempty"`
	// Push the fields to append in the case an array like value where the value is a serialized JSON value
	Push map[string]string `json:"push,omitempty"`
	// Set the fields to set where the value is a serialized JSON value
	Set map[string]string `json:"set,omitempty"`
	// Unset the field names to unset
	Unset []string `json:"unset,omitempty"`
}

var _ datamodel.PartialModel = (*MutateDataPartial)(nil)

// GetModelName returns the name of the model
func (o *MutateDataPartial) GetModelName() datamodel.ModelNameType {
	return MutateDataModelName
}

// ToMap returns the object as a map
func (o *MutateDataPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"id":    toMutateDataObject(o.ID, true),
		"model": toMutateDataObject(o.Model, true),
		"pull":  toMutateDataObject(o.Pull, true),
		"push":  toMutateDataObject(o.Push, true),
		"set":   toMutateDataObject(o.Set, true),
		"unset": toMutateDataObject(o.Unset, true),
	}
	for k, v := range kv {
		if v == nil || reflect.ValueOf(v).IsZero() {
			delete(kv, k)
		} else {

			if k == "pull" {
				if arr, ok := v.([]string); ok {
					if len(arr) == 0 {
						delete(kv, k)
					}
				}
			}

			if k == "unset" {
				if arr, ok := v.([]string); ok {
					if len(arr) == 0 {
						delete(kv, k)
					}
				}
			}
		}
	}
	return kv
}

// Stringify returns the object in JSON format as a string
func (o *MutateDataPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *MutateDataPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *MutateDataPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *MutateDataPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *MutateDataPartial) FromMap(kv map[string]interface{}) {
	if val, ok := kv["id"].(*string); ok {
		o.ID = val
	} else if val, ok := kv["id"].(string); ok {
		o.ID = &val
	} else {
		if val, ok := kv["id"]; ok {
			if val == nil {
				o.ID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.ID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["model"].(*string); ok {
		o.Model = val
	} else if val, ok := kv["model"].(string); ok {
		o.Model = &val
	} else {
		if val, ok := kv["model"]; ok {
			if val == nil {
				o.Model = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Model = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["pull"]; ok {
		if val != nil {
			na := make([]string, 0)
			if a, ok := val.([]string); ok {
				na = append(na, a...)
			} else {
				if a, ok := val.([]interface{}); ok {
					for _, ae := range a {
						if av, ok := ae.(string); ok {
							na = append(na, av)
						} else {
							if badMap, ok := ae.(map[interface{}]interface{}); ok {
								ae = slice.ConvertToStringToInterface(badMap)
							}
							b, _ := json.Marshal(ae)
							var av string
							if err := json.Unmarshal(b, &av); err != nil {
								panic("unsupported type for pull field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else if s, ok := val.(string); ok {
					for _, sv := range strings.Split(s, ",") {
						na = append(na, strings.TrimSpace(sv))
					}
				} else if a, ok := val.(primitive.A); ok {
					for _, ae := range a {
						if av, ok := ae.(string); ok {
							na = append(na, av)
						} else {
							b, _ := json.Marshal(ae)
							var av string
							if err := json.Unmarshal(b, &av); err != nil {
								panic("unsupported type for pull field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for pull field")
				}
			}
			o.Pull = na
		}
	}
	if o.Pull == nil {
		o.Pull = make([]string, 0)
	}
	if val, ok := kv["push"]; ok {
		if val != nil {
			kv := make(map[string]string)
			if m, ok := val.(map[string]string); ok {
				kv = m
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					for k, v := range m {
						if mv, ok := v.(*string); ok {

							kv[k] = *mv

						} else {
							panic("unsupported type for push key:" + k + " field entry types: " + reflect.TypeOf(v).String())
						}
					}
				} else {
					panic("unsupported type for push field entry: " + reflect.TypeOf(val).String())
				}
			}
			o.Push = kv
		}
	}
	if o.Push == nil {
		// here
		o.Push = make(map[string]string)
	}
	if val, ok := kv["set"]; ok {
		if val != nil {
			kv := make(map[string]string)
			if m, ok := val.(map[string]string); ok {
				kv = m
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					for k, v := range m {
						if mv, ok := v.(*string); ok {

							kv[k] = *mv

						} else {
							panic("unsupported type for set key:" + k + " field entry types: " + reflect.TypeOf(v).String())
						}
					}
				} else {
					panic("unsupported type for set field entry: " + reflect.TypeOf(val).String())
				}
			}
			o.Set = kv
		}
	}
	if o.Set == nil {
		// here
		o.Set = make(map[string]string)
	}
	if val, ok := kv["unset"]; ok {
		if val != nil {
			na := make([]string, 0)
			if a, ok := val.([]string); ok {
				na = append(na, a...)
			} else {
				if a, ok := val.([]interface{}); ok {
					for _, ae := range a {
						if av, ok := ae.(string); ok {
							na = append(na, av)
						} else {
							if badMap, ok := ae.(map[interface{}]interface{}); ok {
								ae = slice.ConvertToStringToInterface(badMap)
							}
							b, _ := json.Marshal(ae)
							var av string
							if err := json.Unmarshal(b, &av); err != nil {
								panic("unsupported type for unset field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else if s, ok := val.(string); ok {
					for _, sv := range strings.Split(s, ",") {
						na = append(na, strings.TrimSpace(sv))
					}
				} else if a, ok := val.(primitive.A); ok {
					for _, ae := range a {
						if av, ok := ae.(string); ok {
							na = append(na, av)
						} else {
							b, _ := json.Marshal(ae)
							var av string
							if err := json.Unmarshal(b, &av); err != nil {
								panic("unsupported type for unset field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for unset field")
				}
			}
			o.Unset = na
		}
	}
	if o.Unset == nil {
		o.Unset = make([]string, 0)
	}
	o.setDefaults(false)
}
