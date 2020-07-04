// DO NOT EDIT -- generated code

// Package work - the system which contains project work
package work

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
	"github.com/pinpt/go-common/v10/number"
	pstrings "github.com/pinpt/go-common/v10/strings"
)

const (

	// IssuePriorityTable is the default table name
	IssuePriorityTable datamodel.ModelNameType = "work_issuepriority"

	// IssuePriorityModelName is the model name
	IssuePriorityModelName datamodel.ModelNameType = "work.IssuePriority"
)

const (
	// IssuePriorityModelColorColumn is the column json value color
	IssuePriorityModelColorColumn = "color"
	// IssuePriorityModelCustomerIDColumn is the column json value customer_id
	IssuePriorityModelCustomerIDColumn = "customer_id"
	// IssuePriorityModelDescriptionColumn is the column json value description
	IssuePriorityModelDescriptionColumn = "description"
	// IssuePriorityModelIconURLColumn is the column json value icon_url
	IssuePriorityModelIconURLColumn = "icon_url"
	// IssuePriorityModelIDColumn is the column json value id
	IssuePriorityModelIDColumn = "id"
	// IssuePriorityModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	IssuePriorityModelIntegrationInstanceIDColumn = "integration_instance_id"
	// IssuePriorityModelNameColumn is the column json value name
	IssuePriorityModelNameColumn = "name"
	// IssuePriorityModelOrderColumn is the column json value order
	IssuePriorityModelOrderColumn = "order"
	// IssuePriorityModelRefIDColumn is the column json value ref_id
	IssuePriorityModelRefIDColumn = "ref_id"
	// IssuePriorityModelRefTypeColumn is the column json value ref_type
	IssuePriorityModelRefTypeColumn = "ref_type"
)

// IssuePriority the priority type of an issue
type IssuePriority struct {
	// Color an optional color in hex for display
	Color *string `json:"color,omitempty" codec:"color,omitempty" bson:"color" yaml:"color,omitempty" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Description the description of the priority
	Description *string `json:"description,omitempty" codec:"description,omitempty" bson:"description" yaml:"description,omitempty" faker:"-"`
	// IconURL an optional url for the icon for the priority
	IconURL *string `json:"icon_url,omitempty" codec:"icon_url,omitempty" bson:"icon_url" yaml:"icon_url,omitempty" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// Name the name of the priority
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
	// Order a rank order
	Order int64 `json:"order" codec:"order" bson:"order" yaml:"order" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*IssuePriority)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*IssuePriority)(nil)

func toIssuePriorityObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IssuePriority:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of IssuePriority
func (o *IssuePriority) String() string {
	return fmt.Sprintf("work.IssuePriority<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *IssuePriority) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *IssuePriority) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *IssuePriority) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *IssuePriority) GetModelName() datamodel.ModelNameType {
	return IssuePriorityModelName
}

// NewIssuePriorityID provides a template for generating an ID field for IssuePriority
func NewIssuePriorityID(customerID string, refType string, refID string) string {
	return hash.Values(customerID, refType, refID)
}

func (o *IssuePriority) setDefaults(frommap bool) {

	if o.ID == "" {
		o.ID = hash.Values(o.CustomerID, o.RefType, o.RefID)
	}

	{

	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *IssuePriority) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *IssuePriority) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *IssuePriority) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *IssuePriority) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *IssuePriority) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *IssuePriority) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *IssuePriority) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *IssuePriority) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *IssuePriority) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = IssuePriorityModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *IssuePriority) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *IssuePriority) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of IssuePriority
func (o *IssuePriority) Clone() datamodel.Model {
	c := new(IssuePriority)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *IssuePriority) Anon() datamodel.Model {
	c := new(IssuePriority)
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
func (o *IssuePriority) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *IssuePriority) UnmarshalJSON(data []byte) error {
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
func (o *IssuePriority) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *IssuePriority) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two IssuePriority objects are equal
func (o *IssuePriority) IsEqual(other *IssuePriority) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *IssuePriority) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"color":                   toIssuePriorityObject(o.Color, true),
		"customer_id":             toIssuePriorityObject(o.CustomerID, false),
		"description":             toIssuePriorityObject(o.Description, true),
		"icon_url":                toIssuePriorityObject(o.IconURL, true),
		"id":                      toIssuePriorityObject(o.ID, false),
		"integration_instance_id": toIssuePriorityObject(o.IntegrationInstanceID, true),
		"name":                    toIssuePriorityObject(o.Name, false),
		"order":                   toIssuePriorityObject(o.Order, false),
		"ref_id":                  toIssuePriorityObject(o.RefID, false),
		"ref_type":                toIssuePriorityObject(o.RefType, false),
		"hashcode":                toIssuePriorityObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *IssuePriority) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["color"].(*string); ok {
		o.Color = val
	} else if val, ok := kv["color"].(string); ok {
		o.Color = &val
	} else {
		if val, ok := kv["color"]; ok {
			if val == nil {
				o.Color = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Color = pstrings.Pointer(fmt.Sprintf("%v", val))
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
	if val, ok := kv["description"].(*string); ok {
		o.Description = val
	} else if val, ok := kv["description"].(string); ok {
		o.Description = &val
	} else {
		if val, ok := kv["description"]; ok {
			if val == nil {
				o.Description = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Description = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["icon_url"].(*string); ok {
		o.IconURL = val
	} else if val, ok := kv["icon_url"].(string); ok {
		o.IconURL = &val
	} else {
		if val, ok := kv["icon_url"]; ok {
			if val == nil {
				o.IconURL = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.IconURL = pstrings.Pointer(fmt.Sprintf("%v", val))
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
	if val, ok := kv["order"].(int64); ok {
		o.Order = val
	} else {
		if val, ok := kv["order"]; ok {
			if val == nil {
				o.Order = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Order = number.ToInt64Any(val)
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
func (o *IssuePriority) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Color)
	args = append(args, o.CustomerID)
	args = append(args, o.Description)
	args = append(args, o.IconURL)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.Name)
	args = append(args, o.Order)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// IssuePriorityPartial is a partial struct for upsert mutations for IssuePriority
type IssuePriorityPartial struct {
	// Color an optional color in hex for display
	Color *string `json:"color,omitempty"`
	// Description the description of the priority
	Description *string `json:"description,omitempty"`
	// IconURL an optional url for the icon for the priority
	IconURL *string `json:"icon_url,omitempty"`
	// Name the name of the priority
	Name *string `json:"name,omitempty"`
	// Order a rank order
	Order *int64 `json:"order,omitempty"`
}

var _ datamodel.PartialModel = (*IssuePriorityPartial)(nil)

// GetModelName returns the name of the model
func (o *IssuePriorityPartial) GetModelName() datamodel.ModelNameType {
	return IssuePriorityModelName
}

// ToMap returns the object as a map
func (o *IssuePriorityPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"color":       toIssuePriorityObject(o.Color, true),
		"description": toIssuePriorityObject(o.Description, true),
		"icon_url":    toIssuePriorityObject(o.IconURL, true),
		"name":        toIssuePriorityObject(o.Name, true),
		"order":       toIssuePriorityObject(o.Order, true),
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
func (o *IssuePriorityPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *IssuePriorityPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *IssuePriorityPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *IssuePriorityPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *IssuePriorityPartial) FromMap(kv map[string]interface{}) {
	if val, ok := kv["color"].(*string); ok {
		o.Color = val
	} else if val, ok := kv["color"].(string); ok {
		o.Color = &val
	} else {
		if val, ok := kv["color"]; ok {
			if val == nil {
				o.Color = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Color = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["description"].(*string); ok {
		o.Description = val
	} else if val, ok := kv["description"].(string); ok {
		o.Description = &val
	} else {
		if val, ok := kv["description"]; ok {
			if val == nil {
				o.Description = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Description = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["icon_url"].(*string); ok {
		o.IconURL = val
	} else if val, ok := kv["icon_url"].(string); ok {
		o.IconURL = &val
	} else {
		if val, ok := kv["icon_url"]; ok {
			if val == nil {
				o.IconURL = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.IconURL = pstrings.Pointer(fmt.Sprintf("%v", val))
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
	if val, ok := kv["order"].(*int64); ok {
		o.Order = val
	} else if val, ok := kv["order"].(int64); ok {
		o.Order = &val
	} else {
		if val, ok := kv["order"]; ok {
			if val == nil {
				o.Order = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["long"]
				}
				o.Order = number.Int64Pointer(number.ToInt64Any(val))
			}
		}
	}
	o.setDefaults(false)
}
