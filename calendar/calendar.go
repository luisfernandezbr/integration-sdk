// DO NOT EDIT -- generated code

// Package calendar - public calendar data models
package calendar

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
	// CalendarTopic is the default topic name
	CalendarTopic datamodel.TopicNameType = "calendar_Calendar"

	// CalendarTable is the default table name
	CalendarTable datamodel.ModelNameType = "calendar_calendar"

	// CalendarModelName is the model name
	CalendarModelName datamodel.ModelNameType = "calendar.Calendar"
)

const (
	// CalendarModelActiveColumn is the column json value active
	CalendarModelActiveColumn = "active"
	// CalendarModelCustomerIDColumn is the column json value customer_id
	CalendarModelCustomerIDColumn = "customer_id"
	// CalendarModelDescriptionColumn is the column json value description
	CalendarModelDescriptionColumn = "description"
	// CalendarModelEnabledColumn is the column json value enabled
	CalendarModelEnabledColumn = "enabled"
	// CalendarModelIDColumn is the column json value id
	CalendarModelIDColumn = "id"
	// CalendarModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	CalendarModelIntegrationInstanceIDColumn = "integration_instance_id"
	// CalendarModelNameColumn is the column json value name
	CalendarModelNameColumn = "name"
	// CalendarModelRefIDColumn is the column json value ref_id
	CalendarModelRefIDColumn = "ref_id"
	// CalendarModelRefTypeColumn is the column json value ref_type
	CalendarModelRefTypeColumn = "ref_type"
	// CalendarModelUpdatedAtColumn is the column json value updated_ts
	CalendarModelUpdatedAtColumn = "updated_ts"
	// CalendarModelUserRefIDColumn is the column json value user_ref_id
	CalendarModelUserRefIDColumn = "user_ref_id"
)

// Calendar details for the given integration calendar (deprecated)
type Calendar struct {
	// Active the status of the calendar
	Active bool `json:"active" codec:"active" bson:"active" yaml:"active" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Description the description of the calendar
	Description string `json:"description" codec:"description" bson:"description" yaml:"description" faker:"-"`
	// Enabled Wether this user's calendar can be exported or not
	Enabled bool `json:"enabled" codec:"enabled" bson:"enabled" yaml:"enabled" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// Name the name of the calendar
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
	// RefID the calendar ID
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the record type
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// UserRefID the user ref_id that owns the calendar
	UserRefID string `json:"user_ref_id" codec:"user_ref_id" bson:"user_ref_id" yaml:"user_ref_id" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Calendar)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*Calendar)(nil)

func toCalendarObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *Calendar:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of Calendar
func (o *Calendar) String() string {
	return fmt.Sprintf("calendar.Calendar<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Calendar) GetTopicName() datamodel.TopicNameType {
	return CalendarTopic
}

// GetStreamName returns the name of the stream
func (o *Calendar) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *Calendar) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *Calendar) GetModelName() datamodel.ModelNameType {
	return CalendarModelName
}

// NewCalendarID provides a template for generating an ID field for Calendar
func NewCalendarID(customerID string, refType string, refID string) string {
	return hash.Values("Calendar", customerID, refType, refID)
}

func (o *Calendar) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Calendar", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Calendar) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Calendar) GetTopicKey() string {
	var i interface{} = o.ID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Calendar) GetTimestamp() time.Time {
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
	panic("not sure how to handle the date time format for Calendar")
}

// GetRefID returns the RefID for the object
func (o *Calendar) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Calendar) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *Calendar) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Calendar) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Calendar) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *Calendar) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = CalendarModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Calendar) GetTopicConfig() *datamodel.ModelTopicConfig {
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
func (o *Calendar) GetCustomerID() string {
	return o.CustomerID
}

// SetCustomerID will return the customer_id
func (o *Calendar) SetCustomerID(id string) {
	o.CustomerID = id
}

// GetRefType will return the ref_type
func (o *Calendar) GetRefType() string {
	return o.RefType
}

// SetRefType will return the ref_type
func (o *Calendar) SetRefType(t string) {
	o.RefType = t
}

// Clone returns an exact copy of Calendar
func (o *Calendar) Clone() datamodel.Model {
	c := new(Calendar)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Calendar) Anon() datamodel.Model {
	c := new(Calendar)
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
func (o *Calendar) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Calendar) UnmarshalJSON(data []byte) error {
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
func (o *Calendar) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *Calendar) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two Calendar objects are equal
func (o *Calendar) IsEqual(other *Calendar) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Calendar) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"active":                  toCalendarObject(o.Active, false),
		"customer_id":             toCalendarObject(o.CustomerID, false),
		"description":             toCalendarObject(o.Description, false),
		"enabled":                 toCalendarObject(o.Enabled, false),
		"id":                      toCalendarObject(o.ID, false),
		"integration_instance_id": toCalendarObject(o.IntegrationInstanceID, true),
		"name":                    toCalendarObject(o.Name, false),
		"ref_id":                  toCalendarObject(o.RefID, false),
		"ref_type":                toCalendarObject(o.RefType, false),
		"updated_ts":              toCalendarObject(o.UpdatedAt, false),
		"user_ref_id":             toCalendarObject(o.UserRefID, false),
		"hashcode":                toCalendarObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *Calendar) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["active"].(bool); ok {
		o.Active = val
	} else {
		if val, ok := kv["active"]; ok {
			if val == nil {
				o.Active = false
			} else {
				o.Active = number.ToBoolAny(val)
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
	if val, ok := kv["description"].(string); ok {
		o.Description = val
	} else {
		if val, ok := kv["description"]; ok {
			if val == nil {
				o.Description = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Description = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["enabled"].(bool); ok {
		o.Enabled = val
	} else {
		if val, ok := kv["enabled"]; ok {
			if val == nil {
				o.Enabled = false
			} else {
				o.Enabled = number.ToBoolAny(val)
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
	if val, ok := kv["user_ref_id"].(string); ok {
		o.UserRefID = val
	} else {
		if val, ok := kv["user_ref_id"]; ok {
			if val == nil {
				o.UserRefID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.UserRefID = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *Calendar) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Active)
	args = append(args, o.CustomerID)
	args = append(args, o.Description)
	args = append(args, o.Enabled)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.Name)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.UpdatedAt)
	args = append(args, o.UserRefID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// SetIntegrationInstanceID will set the integration instance ID
func (o *Calendar) SetIntegrationInstanceID(id string) {
	if id == "" {
		o.IntegrationInstanceID = nil
	} else {
		o.IntegrationInstanceID = &id
	}
}

// GetIntegrationInstanceID will return the integration instance ID
func (o *Calendar) GetIntegrationInstanceID() *string {
	return o.IntegrationInstanceID
}

// CalendarPartial is a partial struct for upsert mutations for Calendar
type CalendarPartial struct {
	// Active the status of the calendar
	Active *bool `json:"active,omitempty"`
	// Description the description of the calendar
	Description *string `json:"description,omitempty"`
	// Enabled Wether this user's calendar can be exported or not
	Enabled *bool `json:"enabled,omitempty"`
	// Name the name of the calendar
	Name *string `json:"name,omitempty"`
	// RefID the calendar ID
	RefID *string `json:"ref_id,omitempty"`
	// RefType the record type
	RefType *string `json:"ref_type,omitempty"`
	// UserRefID the user ref_id that owns the calendar
	UserRefID *string `json:"user_ref_id,omitempty"`
}

var _ datamodel.PartialModel = (*CalendarPartial)(nil)

// GetModelName returns the name of the model
func (o *CalendarPartial) GetModelName() datamodel.ModelNameType {
	return CalendarModelName
}

// ToMap returns the object as a map
func (o *CalendarPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"active":      toCalendarObject(o.Active, true),
		"description": toCalendarObject(o.Description, true),
		"enabled":     toCalendarObject(o.Enabled, true),
		"name":        toCalendarObject(o.Name, true),
		"ref_id":      toCalendarObject(o.RefID, true),
		"ref_type":    toCalendarObject(o.RefType, true),
		"user_ref_id": toCalendarObject(o.UserRefID, true),
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
func (o *CalendarPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *CalendarPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *CalendarPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *CalendarPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *CalendarPartial) FromMap(kv map[string]interface{}) {
	if val, ok := kv["active"].(*bool); ok {
		o.Active = val
	} else if val, ok := kv["active"].(bool); ok {
		o.Active = &val
	} else {
		if val, ok := kv["active"]; ok {
			if val == nil {
				o.Active = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Active = number.BoolPointer(number.ToBoolAny(val))
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
	if val, ok := kv["enabled"].(*bool); ok {
		o.Enabled = val
	} else if val, ok := kv["enabled"].(bool); ok {
		o.Enabled = &val
	} else {
		if val, ok := kv["enabled"]; ok {
			if val == nil {
				o.Enabled = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Enabled = number.BoolPointer(number.ToBoolAny(val))
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
	if val, ok := kv["ref_id"].(*string); ok {
		o.RefID = val
	} else if val, ok := kv["ref_id"].(string); ok {
		o.RefID = &val
	} else {
		if val, ok := kv["ref_id"]; ok {
			if val == nil {
				o.RefID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.RefID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["ref_type"].(*string); ok {
		o.RefType = val
	} else if val, ok := kv["ref_type"].(string); ok {
		o.RefType = &val
	} else {
		if val, ok := kv["ref_type"]; ok {
			if val == nil {
				o.RefType = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.RefType = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["user_ref_id"].(*string); ok {
		o.UserRefID = val
	} else if val, ok := kv["user_ref_id"].(string); ok {
		o.UserRefID = &val
	} else {
		if val, ok := kv["user_ref_id"]; ok {
			if val == nil {
				o.UserRefID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.UserRefID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	o.setDefaults(false)
}
