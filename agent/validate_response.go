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

	// ValidateResponseTable is the default table name
	ValidateResponseTable datamodel.ModelNameType = "agent_validateresponse"

	// ValidateResponseModelName is the model name
	ValidateResponseModelName datamodel.ModelNameType = "agent.ValidateResponse"
)

const (
	// ValidateResponseModelCustomerIDColumn is the column json value customer_id
	ValidateResponseModelCustomerIDColumn = "customer_id"
	// ValidateResponseModelErrorColumn is the column json value error
	ValidateResponseModelErrorColumn = "error"
	// ValidateResponseModelIDColumn is the column json value id
	ValidateResponseModelIDColumn = "id"
	// ValidateResponseModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	ValidateResponseModelIntegrationInstanceIDColumn = "integration_instance_id"
	// ValidateResponseModelRefIDColumn is the column json value ref_id
	ValidateResponseModelRefIDColumn = "ref_id"
	// ValidateResponseModelRefTypeColumn is the column json value ref_type
	ValidateResponseModelRefTypeColumn = "ref_type"
	// ValidateResponseModelResultColumn is the column json value result
	ValidateResponseModelResultColumn = "result"
	// ValidateResponseModelSessionIDColumn is the column json value session_id
	ValidateResponseModelSessionIDColumn = "session_id"
	// ValidateResponseModelSuccessColumn is the column json value success
	ValidateResponseModelSuccessColumn = "success"
)

// ValidateResponse A validation response sent from the agent with the result of the validation
type ValidateResponse struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Error if the validation was not successful, the description of the error
	Error *string `json:"error,omitempty" codec:"error,omitempty" bson:"error" yaml:"error,omitempty" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// Result The result of the validation in JSON
	Result *string `json:"result,omitempty" codec:"result,omitempty" bson:"result" yaml:"result,omitempty" faker:"-"`
	// SessionID A unique session ID from the request
	SessionID string `json:"session_id" codec:"session_id" bson:"session_id" yaml:"session_id" faker:"-"`
	// Success if the validation was successful or not
	Success bool `json:"success" codec:"success" bson:"success" yaml:"success" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*ValidateResponse)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*ValidateResponse)(nil)

func toValidateResponseObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ValidateResponse:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of ValidateResponse
func (o *ValidateResponse) String() string {
	return fmt.Sprintf("agent.ValidateResponse<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *ValidateResponse) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *ValidateResponse) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *ValidateResponse) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *ValidateResponse) GetModelName() datamodel.ModelNameType {
	return ValidateResponseModelName
}

// NewValidateResponseID provides a template for generating an ID field for ValidateResponse
func NewValidateResponseID(customerID string, refType string, refID string) string {
	return hash.Values("ValidateResponse", customerID, refType, refID)
}

func (o *ValidateResponse) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("ValidateResponse", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *ValidateResponse) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *ValidateResponse) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *ValidateResponse) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *ValidateResponse) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *ValidateResponse) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *ValidateResponse) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *ValidateResponse) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *ValidateResponse) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *ValidateResponse) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = ValidateResponseModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *ValidateResponse) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *ValidateResponse) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of ValidateResponse
func (o *ValidateResponse) Clone() datamodel.Model {
	c := new(ValidateResponse)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *ValidateResponse) Anon() datamodel.Model {
	c := new(ValidateResponse)
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
func (o *ValidateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *ValidateResponse) UnmarshalJSON(data []byte) error {
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
func (o *ValidateResponse) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *ValidateResponse) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two ValidateResponse objects are equal
func (o *ValidateResponse) IsEqual(other *ValidateResponse) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *ValidateResponse) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"customer_id":             toValidateResponseObject(o.CustomerID, false),
		"error":                   toValidateResponseObject(o.Error, true),
		"id":                      toValidateResponseObject(o.ID, false),
		"integration_instance_id": toValidateResponseObject(o.IntegrationInstanceID, true),
		"ref_id":                  toValidateResponseObject(o.RefID, false),
		"ref_type":                toValidateResponseObject(o.RefType, false),
		"result":                  toValidateResponseObject(o.Result, true),
		"session_id":              toValidateResponseObject(o.SessionID, false),
		"success":                 toValidateResponseObject(o.Success, false),
		"hashcode":                toValidateResponseObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *ValidateResponse) FromMap(kv map[string]interface{}) {

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
	if val, ok := kv["error"].(*string); ok {
		o.Error = val
	} else if val, ok := kv["error"].(string); ok {
		o.Error = &val
	} else {
		if val, ok := kv["error"]; ok {
			if val == nil {
				o.Error = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Error = pstrings.Pointer(fmt.Sprintf("%v", val))
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
	if val, ok := kv["result"].(*string); ok {
		o.Result = val
	} else if val, ok := kv["result"].(string); ok {
		o.Result = &val
	} else {
		if val, ok := kv["result"]; ok {
			if val == nil {
				o.Result = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Result = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["session_id"].(string); ok {
		o.SessionID = val
	} else {
		if val, ok := kv["session_id"]; ok {
			if val == nil {
				o.SessionID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.SessionID = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["success"].(bool); ok {
		o.Success = val
	} else {
		if val, ok := kv["success"]; ok {
			if val == nil {
				o.Success = false
			} else {
				o.Success = number.ToBoolAny(val)
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *ValidateResponse) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CustomerID)
	args = append(args, o.Error)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.Result)
	args = append(args, o.SessionID)
	args = append(args, o.Success)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// ValidateResponsePartial is a partial struct for upsert mutations for ValidateResponse
type ValidateResponsePartial struct {
	// Error if the validation was not successful, the description of the error
	Error *string `json:"error,omitempty"`
	// Result The result of the validation in JSON
	Result *string `json:"result,omitempty"`
	// SessionID A unique session ID from the request
	SessionID *string `json:"session_id,omitempty"`
	// Success if the validation was successful or not
	Success *bool `json:"success,omitempty"`
}

var _ datamodel.PartialModel = (*ValidateResponsePartial)(nil)

// GetModelName returns the name of the model
func (o *ValidateResponsePartial) GetModelName() datamodel.ModelNameType {
	return ValidateResponseModelName
}

// ToMap returns the object as a map
func (o *ValidateResponsePartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"error":      toValidateResponseObject(o.Error, true),
		"result":     toValidateResponseObject(o.Result, true),
		"session_id": toValidateResponseObject(o.SessionID, true),
		"success":    toValidateResponseObject(o.Success, true),
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
func (o *ValidateResponsePartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *ValidateResponsePartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *ValidateResponsePartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *ValidateResponsePartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *ValidateResponsePartial) FromMap(kv map[string]interface{}) {
	if val, ok := kv["error"].(*string); ok {
		o.Error = val
	} else if val, ok := kv["error"].(string); ok {
		o.Error = &val
	} else {
		if val, ok := kv["error"]; ok {
			if val == nil {
				o.Error = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Error = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["result"].(*string); ok {
		o.Result = val
	} else if val, ok := kv["result"].(string); ok {
		o.Result = &val
	} else {
		if val, ok := kv["result"]; ok {
			if val == nil {
				o.Result = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Result = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["session_id"].(*string); ok {
		o.SessionID = val
	} else if val, ok := kv["session_id"].(string); ok {
		o.SessionID = &val
	} else {
		if val, ok := kv["session_id"]; ok {
			if val == nil {
				o.SessionID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.SessionID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["success"].(*bool); ok {
		o.Success = val
	} else if val, ok := kv["success"].(bool); ok {
		o.Success = &val
	} else {
		if val, ok := kv["success"]; ok {
			if val == nil {
				o.Success = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Success = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}
	o.setDefaults(false)
}
