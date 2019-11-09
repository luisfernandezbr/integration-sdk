// DO NOT EDIT -- generated code

// Package work - the system which contains project work
package work

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/bxcodec/faker"
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
)

const (
	// CustomFieldTopic is the default topic name
	CustomFieldTopic datamodel.TopicNameType = "work_CustomField_topic"

	// CustomFieldModelName is the model name
	CustomFieldModelName datamodel.ModelNameType = "work.CustomField"
)

const (
	// CustomFieldCustomerIDColumn is the customer_id column name
	CustomFieldCustomerIDColumn = "CustomerID"
	// CustomFieldIDColumn is the id column name
	CustomFieldIDColumn = "ID"
	// CustomFieldKeyColumn is the key column name
	CustomFieldKeyColumn = "Key"
	// CustomFieldNameColumn is the name column name
	CustomFieldNameColumn = "Name"
	// CustomFieldRefIDColumn is the ref_id column name
	CustomFieldRefIDColumn = "RefID"
	// CustomFieldRefTypeColumn is the ref_type column name
	CustomFieldRefTypeColumn = "RefType"
	// CustomFieldUpdatedAtColumn is the updated_ts column name
	CustomFieldUpdatedAtColumn = "UpdatedAt"
)

// CustomField user defined fields
type CustomField struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// Key key of the field
	Key string `json:"key" yaml:"key" faker:"-"`
	// Name the name of the field
	Name string `json:"name" yaml:"name" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" yaml:"ref_type" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" yaml:"updated_ts" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*CustomField)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*CustomField)(nil)

func toCustomFieldObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CustomField:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of CustomField
func (o *CustomField) String() string {
	return fmt.Sprintf("work.CustomField<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *CustomField) GetTopicName() datamodel.TopicNameType {
	return CustomFieldTopic
}

// GetStreamName returns the name of the stream
func (o *CustomField) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *CustomField) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *CustomField) GetModelName() datamodel.ModelNameType {
	return CustomFieldModelName
}

// NewCustomFieldID provides a template for generating an ID field for CustomField
func NewCustomFieldID(customerID string, refType string, refID string) string {
	return hash.Values("CustomField", customerID, refType, refID)
}

func (o *CustomField) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("CustomField", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *CustomField) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *CustomField) GetTopicKey() string {
	var i interface{} = o.CustomerID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *CustomField) GetTimestamp() time.Time {
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
	panic("not sure how to handle the date time format for CustomField")
}

// GetRefID returns the RefID for the object
func (o *CustomField) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *CustomField) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *CustomField) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *CustomField) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *CustomField) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *CustomField) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = CustomFieldModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *CustomField) GetTopicConfig() *datamodel.ModelTopicConfig {
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
		Key:               "customer_id",
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
func (o *CustomField) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of CustomField
func (o *CustomField) Clone() datamodel.Model {
	c := new(CustomField)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *CustomField) Anon() datamodel.Model {
	c := new(CustomField)
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
func (o *CustomField) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *CustomField) UnmarshalJSON(data []byte) error {
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
func (o *CustomField) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two CustomField objects are equal
func (o *CustomField) IsEqual(other *CustomField) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *CustomField) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"customer_id": toCustomFieldObject(o.CustomerID, false),
		"id":          toCustomFieldObject(o.ID, false),
		"key":         toCustomFieldObject(o.Key, false),
		"name":        toCustomFieldObject(o.Name, false),
		"ref_id":      toCustomFieldObject(o.RefID, false),
		"ref_type":    toCustomFieldObject(o.RefType, false),
		"updated_ts":  toCustomFieldObject(o.UpdatedAt, false),
		"hashcode":    toCustomFieldObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *CustomField) FromMap(kv map[string]interface{}) {

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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.ID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["key"].(string); ok {
		o.Key = val
	} else {
		if val, ok := kv["key"]; ok {
			if val == nil {
				o.Key = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Key = fmt.Sprintf("%v", val)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				o.UpdatedAt = number.ToInt64Any(nil)
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
func (o *CustomField) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.Key)
	args = append(args, o.Name)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.UpdatedAt)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *CustomField) GetEventAPIConfig() datamodel.EventAPIConfig {
	return datamodel.EventAPIConfig{
		Publish: datamodel.EventAPIPublish{
			Public: false,
		},
		Subscribe: datamodel.EventAPISubscribe{
			Public: false,
			Key:    "",
		},
	}
}
