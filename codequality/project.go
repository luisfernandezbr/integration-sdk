// DO NOT EDIT -- generated code

// Package codequality - the system which contains code quality
package codequality

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
	// ProjectTopic is the default topic name
	ProjectTopic datamodel.TopicNameType = "codequality_Project_topic"

	// ProjectModelName is the model name
	ProjectModelName datamodel.ModelNameType = "codequality.Project"
)

const (
	// ProjectCustomerIDColumn is the customer_id column name
	ProjectCustomerIDColumn = "CustomerID"
	// ProjectIDColumn is the id column name
	ProjectIDColumn = "ID"
	// ProjectIdentifierColumn is the identifier column name
	ProjectIdentifierColumn = "Identifier"
	// ProjectNameColumn is the name column name
	ProjectNameColumn = "Name"
	// ProjectRefIDColumn is the ref_id column name
	ProjectRefIDColumn = "RefID"
	// ProjectRefTypeColumn is the ref_type column name
	ProjectRefTypeColumn = "RefType"
	// ProjectUpdatedAtColumn is the updated_ts column name
	ProjectUpdatedAtColumn = "UpdatedAt"
)

// Project project information
type Project struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// Identifier the common identifier of the project
	Identifier string `json:"identifier" codec:"identifier" bson:"identifier" yaml:"identifier" faker:"-"`
	// Name the name of the project
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
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
	return ProjectTopic
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
	var i interface{} = o.CustomerID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Project) GetTimestamp() time.Time {
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
	panic("not sure how to handle the date time format for Project")
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
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *Project) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = ProjectModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Project) GetTopicConfig() *datamodel.ModelTopicConfig {
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
func (o *Project) GetCustomerID() string {

	return o.CustomerID

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

// IsEqual returns true if the two Project objects are equal
func (o *Project) IsEqual(other *Project) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Project) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"customer_id": toProjectObject(o.CustomerID, false),
		"id":          toProjectObject(o.ID, false),
		"identifier":  toProjectObject(o.Identifier, false),
		"name":        toProjectObject(o.Name, false),
		"ref_id":      toProjectObject(o.RefID, false),
		"ref_type":    toProjectObject(o.RefType, false),
		"updated_ts":  toProjectObject(o.UpdatedAt, false),
		"hashcode":    toProjectObject(o.Hashcode, false),
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

	if val, ok := kv["identifier"].(string); ok {
		o.Identifier = val
	} else {
		if val, ok := kv["identifier"]; ok {
			if val == nil {
				o.Identifier = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Identifier = fmt.Sprintf("%v", val)
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
func (o *Project) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.Identifier)
	args = append(args, o.Name)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.UpdatedAt)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *Project) GetEventAPIConfig() datamodel.EventAPIConfig {
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
