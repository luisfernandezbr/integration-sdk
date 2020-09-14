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
	"github.com/pinpt/go-common/v10/datetime"
	"github.com/pinpt/go-common/v10/hash"
	pjson "github.com/pinpt/go-common/v10/json"
	"github.com/pinpt/go-common/v10/number"
	pstrings "github.com/pinpt/go-common/v10/strings"
)

const (

	// UpdateRequestTable is the default table name
	UpdateRequestTable datamodel.ModelNameType = "agent_updaterequest"

	// UpdateRequestModelName is the model name
	UpdateRequestModelName datamodel.ModelNameType = "agent.UpdateRequest"
)

const (
	// UpdateRequestModelCustomerIDColumn is the column json value customer_id
	UpdateRequestModelCustomerIDColumn = "customer_id"
	// UpdateRequestModelIDColumn is the column json value id
	UpdateRequestModelIDColumn = "id"
	// UpdateRequestModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	UpdateRequestModelIntegrationInstanceIDColumn = "integration_instance_id"
	// UpdateRequestModelRefIDColumn is the column json value ref_id
	UpdateRequestModelRefIDColumn = "ref_id"
	// UpdateRequestModelRefTypeColumn is the column json value ref_type
	UpdateRequestModelRefTypeColumn = "ref_type"
	// UpdateRequestModelRequestDateColumn is the column json value request_date
	UpdateRequestModelRequestDateColumn = "request_date"
	// UpdateRequestModelRequestDateEpochColumn is the column json value epoch
	UpdateRequestModelRequestDateEpochColumn = "epoch"
	// UpdateRequestModelRequestDateOffsetColumn is the column json value offset
	UpdateRequestModelRequestDateOffsetColumn = "offset"
	// UpdateRequestModelRequestDateRfc3339Column is the column json value rfc3339
	UpdateRequestModelRequestDateRfc3339Column = "rfc3339"
	// UpdateRequestModelToVersionColumn is the column json value to_version
	UpdateRequestModelToVersionColumn = "to_version"
	// UpdateRequestModelUUIDColumn is the column json value uuid
	UpdateRequestModelUUIDColumn = "uuid"
)

// UpdateRequestRequestDate represents the object structure for request_date
type UpdateRequestRequestDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toUpdateRequestRequestDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *UpdateRequestRequestDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *UpdateRequestRequestDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toUpdateRequestRequestDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toUpdateRequestRequestDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toUpdateRequestRequestDateObject(o.Rfc3339, false),
	}
}

func (o *UpdateRequestRequestDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *UpdateRequestRequestDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["epoch"].(int64); ok {
		o.Epoch = val
	} else {
		if val, ok := kv["epoch"]; ok {
			if val == nil {
				o.Epoch = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Epoch = number.ToInt64Any(val)
			}
		}
	}
	if val, ok := kv["offset"].(int64); ok {
		o.Offset = val
	} else {
		if val, ok := kv["offset"]; ok {
			if val == nil {
				o.Offset = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Offset = number.ToInt64Any(val)
			}
		}
	}
	if val, ok := kv["rfc3339"].(string); ok {
		o.Rfc3339 = val
	} else {
		if val, ok := kv["rfc3339"]; ok {
			if val == nil {
				o.Rfc3339 = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// UpdateRequest request agent update
type UpdateRequest struct {
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
	// RequestDate the date when the request was made
	RequestDate UpdateRequestRequestDate `json:"request_date" codec:"request_date" bson:"request_date" yaml:"request_date" faker:"-"`
	// ToVersion the version we're installing
	ToVersion string `json:"to_version" codec:"to_version" bson:"to_version" yaml:"to_version" faker:"-"`
	// UUID the agent unique identifier
	UUID string `json:"uuid" codec:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*UpdateRequest)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*UpdateRequest)(nil)

func toUpdateRequestObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *UpdateRequest:
		return v.ToMap()

	case UpdateRequestRequestDate:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of UpdateRequest
func (o *UpdateRequest) String() string {
	return fmt.Sprintf("agent.UpdateRequest<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *UpdateRequest) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *UpdateRequest) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *UpdateRequest) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *UpdateRequest) GetModelName() datamodel.ModelNameType {
	return UpdateRequestModelName
}

// NewUpdateRequestID provides a template for generating an ID field for UpdateRequest
func NewUpdateRequestID(customerID string, refType string, refID string) string {
	return hash.Values("UpdateRequest", customerID, refType, refID)
}

func (o *UpdateRequest) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("UpdateRequest", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *UpdateRequest) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *UpdateRequest) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *UpdateRequest) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *UpdateRequest) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *UpdateRequest) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *UpdateRequest) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *UpdateRequest) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *UpdateRequest) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *UpdateRequest) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = UpdateRequestModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *UpdateRequest) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *UpdateRequest) GetCustomerID() string {

	return o.CustomerID

}

// SetCustomerID will return the customer_id
func (o *UpdateRequest) SetCustomerID(id string) {

	o.CustomerID = id

}

// GetRefType will return the ref_type
func (o *UpdateRequest) GetRefType() string {

	return o.RefType

}

// SetRefType will return the ref_type
func (o *UpdateRequest) SetRefType(t string) {

	o.RefType = t

}

// Clone returns an exact copy of UpdateRequest
func (o *UpdateRequest) Clone() datamodel.Model {
	c := new(UpdateRequest)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *UpdateRequest) Anon() datamodel.Model {
	c := new(UpdateRequest)
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
func (o *UpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *UpdateRequest) UnmarshalJSON(data []byte) error {
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
func (o *UpdateRequest) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *UpdateRequest) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two UpdateRequest objects are equal
func (o *UpdateRequest) IsEqual(other *UpdateRequest) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *UpdateRequest) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"customer_id":             toUpdateRequestObject(o.CustomerID, false),
		"id":                      toUpdateRequestObject(o.ID, false),
		"integration_instance_id": toUpdateRequestObject(o.IntegrationInstanceID, true),
		"ref_id":                  toUpdateRequestObject(o.RefID, false),
		"ref_type":                toUpdateRequestObject(o.RefType, false),
		"request_date":            toUpdateRequestObject(o.RequestDate, false),
		"to_version":              toUpdateRequestObject(o.ToVersion, false),
		"uuid":                    toUpdateRequestObject(o.UUID, false),
		"hashcode":                toUpdateRequestObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *UpdateRequest) FromMap(kv map[string]interface{}) {

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

	if val, ok := kv["request_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.RequestDate.FromMap(kv)
		} else if sv, ok := val.(UpdateRequestRequestDate); ok {
			// struct
			o.RequestDate = sv
		} else if sp, ok := val.(*UpdateRequestRequestDate); ok {
			// struct pointer
			o.RequestDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.RequestDate.Epoch = dt.Epoch
			o.RequestDate.Rfc3339 = dt.Rfc3339
			o.RequestDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.RequestDate.Epoch = dt.Epoch
			o.RequestDate.Rfc3339 = dt.Rfc3339
			o.RequestDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.RequestDate.Epoch = dt.Epoch
				o.RequestDate.Rfc3339 = dt.Rfc3339
				o.RequestDate.Offset = dt.Offset
			}
		}
	} else {
		o.RequestDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["to_version"].(string); ok {
		o.ToVersion = val
	} else {
		if val, ok := kv["to_version"]; ok {
			if val == nil {
				o.ToVersion = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.ToVersion = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["uuid"].(string); ok {
		o.UUID = val
	} else {
		if val, ok := kv["uuid"]; ok {
			if val == nil {
				o.UUID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.UUID = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *UpdateRequest) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RequestDate)
	args = append(args, o.ToVersion)
	args = append(args, o.UUID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// SetIntegrationInstanceID will set the integration instance ID
func (o *UpdateRequest) SetIntegrationInstanceID(id string) {
	if id == "" {
		o.IntegrationInstanceID = nil
	} else {
		o.IntegrationInstanceID = &id
	}
}

// GetIntegrationInstanceID will return the integration instance ID
func (o *UpdateRequest) GetIntegrationInstanceID() *string {
	return o.IntegrationInstanceID
}

// UpdateRequestPartial is a partial struct for upsert mutations for UpdateRequest
type UpdateRequestPartial struct {
	// RequestDate the date when the request was made
	RequestDate *UpdateRequestRequestDate `json:"request_date,omitempty"`
	// ToVersion the version we're installing
	ToVersion *string `json:"to_version,omitempty"`
	// UUID the agent unique identifier
	UUID *string `json:"uuid,omitempty"`
}

var _ datamodel.PartialModel = (*UpdateRequestPartial)(nil)

// GetModelName returns the name of the model
func (o *UpdateRequestPartial) GetModelName() datamodel.ModelNameType {
	return UpdateRequestModelName
}

// ToMap returns the object as a map
func (o *UpdateRequestPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"request_date": toUpdateRequestObject(o.RequestDate, true),
		"to_version":   toUpdateRequestObject(o.ToVersion, true),
		"uuid":         toUpdateRequestObject(o.UUID, true),
	}
	for k, v := range kv {
		if v == nil || reflect.ValueOf(v).IsZero() {
			delete(kv, k)
		} else {
			if k == "request_date" {
				if dt, ok := v.(*UpdateRequestRequestDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}
		}
	}
	return kv
}

// Stringify returns the object in JSON format as a string
func (o *UpdateRequestPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *UpdateRequestPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *UpdateRequestPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *UpdateRequestPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *UpdateRequestPartial) FromMap(kv map[string]interface{}) {

	if o.RequestDate == nil {
		o.RequestDate = &UpdateRequestRequestDate{}
	}

	if val, ok := kv["request_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.RequestDate.FromMap(kv)
		} else if sv, ok := val.(UpdateRequestRequestDate); ok {
			// struct
			o.RequestDate = &sv
		} else if sp, ok := val.(*UpdateRequestRequestDate); ok {
			// struct pointer
			o.RequestDate = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.RequestDate.Epoch = dt.Epoch
			o.RequestDate.Rfc3339 = dt.Rfc3339
			o.RequestDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.RequestDate.Epoch = dt.Epoch
			o.RequestDate.Rfc3339 = dt.Rfc3339
			o.RequestDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.RequestDate.Epoch = dt.Epoch
				o.RequestDate.Rfc3339 = dt.Rfc3339
				o.RequestDate.Offset = dt.Offset
			}
		}
	} else {
		o.RequestDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["to_version"].(*string); ok {
		o.ToVersion = val
	} else if val, ok := kv["to_version"].(string); ok {
		o.ToVersion = &val
	} else {
		if val, ok := kv["to_version"]; ok {
			if val == nil {
				o.ToVersion = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.ToVersion = pstrings.Pointer(fmt.Sprintf("%v", val))
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
