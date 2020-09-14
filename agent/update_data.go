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
	"github.com/pinpt/go-common/v10/number"
	"github.com/pinpt/go-common/v10/slice"
	pstrings "github.com/pinpt/go-common/v10/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// UpdateDataTopic is the default topic name
	UpdateDataTopic datamodel.TopicNameType = "agent_UpdateData"

	// UpdateDataTable is the default table name
	UpdateDataTable datamodel.ModelNameType = "agent_updatedata"

	// UpdateDataModelName is the model name
	UpdateDataModelName datamodel.ModelNameType = "agent.UpdateData"
)

const (
	// UpdateDataModelCustomerIDColumn is the column json value customer_id
	UpdateDataModelCustomerIDColumn = "customer_id"
	// UpdateDataModelIDColumn is the column json value id
	UpdateDataModelIDColumn = "id"
	// UpdateDataModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	UpdateDataModelIntegrationInstanceIDColumn = "integration_instance_id"
	// UpdateDataModelModelColumn is the column json value model
	UpdateDataModelModelColumn = "model"
	// UpdateDataModelPullColumn is the column json value pull
	UpdateDataModelPullColumn = "pull"
	// UpdateDataModelPushColumn is the column json value push
	UpdateDataModelPushColumn = "push"
	// UpdateDataModelRefIDColumn is the column json value ref_id
	UpdateDataModelRefIDColumn = "ref_id"
	// UpdateDataModelRefTypeColumn is the column json value ref_type
	UpdateDataModelRefTypeColumn = "ref_type"
	// UpdateDataModelSetColumn is the column json value set
	UpdateDataModelSetColumn = "set"
	// UpdateDataModelUnsetColumn is the column json value unset
	UpdateDataModelUnsetColumn = "unset"
	// UpdateDataModelUpdatedAtColumn is the column json value updated_ts
	UpdateDataModelUpdatedAtColumn = "updated_ts"
)

// UpdateData request to update data on a per field basis vs the entire object
type UpdateData struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the ID of the object to mutate. this object must exist or will not succeed
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// Model the name of the model
	Model string `json:"model" codec:"model" bson:"model" yaml:"model" faker:"-"`
	// Pull the fields to remove in the case an array like value where the value is the field name
	Pull map[string]string `json:"pull" codec:"pull" bson:"pull" yaml:"pull" faker:"-"`
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
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*UpdateData)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*UpdateData)(nil)

func toUpdateDataObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *UpdateData:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of UpdateData
func (o *UpdateData) String() string {
	return fmt.Sprintf("agent.UpdateData<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *UpdateData) GetTopicName() datamodel.TopicNameType {
	return UpdateDataTopic
}

// GetStreamName returns the name of the stream
func (o *UpdateData) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *UpdateData) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *UpdateData) GetModelName() datamodel.ModelNameType {
	return UpdateDataModelName
}

// NewUpdateDataID provides a template for generating an ID field for UpdateData
func NewUpdateDataID(customerID string, refType string, refID string) string {
	return hash.Values("UpdateData", customerID, refType, refID)
}

func (o *UpdateData) setDefaults(frommap bool) {
	if o.Unset == nil {
		o.Unset = make([]string, 0)
	}

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("UpdateData", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *UpdateData) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *UpdateData) GetTopicKey() string {
	var i interface{} = o.ID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *UpdateData) GetTimestamp() time.Time {
	var dt interface{} = o.UpdatedAt
	switch v := dt.(type) {
	case int64:
		return datetime.DateFromEpoch(v).UTC()
	case string:
		tv, err := datetime.ISODateToTime(v)
		if err != nil {
			panic(err)
		}
		return tv.UTC()
	case time.Time:
		return v.UTC()
	}
	panic("not sure how to handle the date time format for UpdateData")
}

// GetRefID returns the RefID for the object
func (o *UpdateData) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *UpdateData) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *UpdateData) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *UpdateData) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *UpdateData) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *UpdateData) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = UpdateDataModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *UpdateData) GetTopicConfig() *datamodel.ModelTopicConfig {
	retention, err := time.ParseDuration("87360h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 87360h0m0s. " + err.Error())
	}

	ttl, err := time.ParseDuration("0s")
	if err != nil {
		ttl = 0
	}
	if ttl == 0 && retention != 0 {
		ttl = retention // they should be the same if not set
	}
	return &datamodel.ModelTopicConfig{
		Key:               "id",
		Timestamp:         "updated_ts",
		NumPartitions:     8,
		CleanupPolicy:     datamodel.CleanupPolicy("compact"),
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetCustomerID will return the customer_id
func (o *UpdateData) GetCustomerID() string {

	return o.CustomerID

}

// SetCustomerID will return the customer_id
func (o *UpdateData) SetCustomerID(id string) {

	o.CustomerID = id

}

// GetRefType will return the ref_type
func (o *UpdateData) GetRefType() string {

	return o.RefType

}

// SetRefType will return the ref_type
func (o *UpdateData) SetRefType(t string) {

	o.RefType = id

}

// Clone returns an exact copy of UpdateData
func (o *UpdateData) Clone() datamodel.Model {
	c := new(UpdateData)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *UpdateData) Anon() datamodel.Model {
	c := new(UpdateData)
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
func (o *UpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *UpdateData) UnmarshalJSON(data []byte) error {
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
func (o *UpdateData) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *UpdateData) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two UpdateData objects are equal
func (o *UpdateData) IsEqual(other *UpdateData) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *UpdateData) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"customer_id":             toUpdateDataObject(o.CustomerID, false),
		"id":                      toUpdateDataObject(o.ID, false),
		"integration_instance_id": toUpdateDataObject(o.IntegrationInstanceID, true),
		"model":                   toUpdateDataObject(o.Model, false),
		"pull":                    toUpdateDataObject(o.Pull, false),
		"push":                    toUpdateDataObject(o.Push, false),
		"ref_id":                  toUpdateDataObject(o.RefID, false),
		"ref_type":                toUpdateDataObject(o.RefType, false),
		"set":                     toUpdateDataObject(o.Set, false),
		"unset":                   toUpdateDataObject(o.Unset, false),
		"updated_ts":              toUpdateDataObject(o.UpdatedAt, false),
		"hashcode":                toUpdateDataObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *UpdateData) FromMap(kv map[string]interface{}) {

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
					panic("unsupported type for pull field entry: " + reflect.TypeOf(val).String())
				}
			}
			o.Pull = kv
		}
	}
	if o.Pull == nil {
		// here
		o.Pull = make(map[string]string)
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
	if val, ok := kv["updated_ts"].(int64); ok {
		o.UpdatedAt = val
	} else {
		if val, ok := kv["updated_ts"]; ok {
			if val == nil {
				o.UpdatedAt = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.UpdatedAt = number.ToInt64Any(val)
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *UpdateData) Hash() string {
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
	args = append(args, o.UpdatedAt)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// SetIntegrationInstanceID will set the integration instance ID
func (o *UpdateData) SetIntegrationInstanceID(id string) {
	if id == "" {
		o.IntegrationInstanceID = nil
	} else {
		o.IntegrationInstanceID = &id
	}
}

// GetIntegrationInstanceID will return the integration instance ID
func (o *UpdateData) GetIntegrationInstanceID() *string {
	return o.IntegrationInstanceID
}

// UpdateDataPartial is a partial struct for upsert mutations for UpdateData
type UpdateDataPartial struct {
	// ID the ID of the object to mutate. this object must exist or will not succeed
	ID *string `json:"id,omitempty"`
	// Model the name of the model
	Model *string `json:"model,omitempty"`
	// Pull the fields to remove in the case an array like value where the value is the field name
	Pull map[string]string `json:"pull,omitempty"`
	// Push the fields to append in the case an array like value where the value is a serialized JSON value
	Push map[string]string `json:"push,omitempty"`
	// Set the fields to set where the value is a serialized JSON value
	Set map[string]string `json:"set,omitempty"`
	// Unset the field names to unset
	Unset []string `json:"unset,omitempty"`
}

var _ datamodel.PartialModel = (*UpdateDataPartial)(nil)

// GetModelName returns the name of the model
func (o *UpdateDataPartial) GetModelName() datamodel.ModelNameType {
	return UpdateDataModelName
}

// ToMap returns the object as a map
func (o *UpdateDataPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"id":    toUpdateDataObject(o.ID, true),
		"model": toUpdateDataObject(o.Model, true),
		"pull":  toUpdateDataObject(o.Pull, true),
		"push":  toUpdateDataObject(o.Push, true),
		"set":   toUpdateDataObject(o.Set, true),
		"unset": toUpdateDataObject(o.Unset, true),
	}
	for k, v := range kv {
		if v == nil || reflect.ValueOf(v).IsZero() {
			delete(kv, k)
		} else {

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
func (o *UpdateDataPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *UpdateDataPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *UpdateDataPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *UpdateDataPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *UpdateDataPartial) FromMap(kv map[string]interface{}) {
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
			kv := make(map[string]string)
			if m, ok := val.(map[string]string); ok {
				kv = m
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					for k, v := range m {
						if mv, ok := v.(*string); ok {

							kv[k] = *mv

						} else {
							panic("unsupported type for pull key:" + k + " field entry types: " + reflect.TypeOf(v).String())
						}
					}
				} else {
					panic("unsupported type for pull field entry: " + reflect.TypeOf(val).String())
				}
			}
			o.Pull = kv
		}
	}
	if o.Pull == nil {
		// here
		o.Pull = make(map[string]string)
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
