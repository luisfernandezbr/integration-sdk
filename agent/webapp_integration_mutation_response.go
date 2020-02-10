// DO NOT EDIT -- generated code

// Package agent -
package agent

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/bxcodec/faker"
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	pstrings "github.com/pinpt/go-common/strings"
)

const (

	// WebappIntegrationMutationResponseTable is the default table name
	WebappIntegrationMutationResponseTable datamodel.ModelNameType = "agent_webappintegrationmutationresponse"

	// WebappIntegrationMutationResponseModelName is the model name
	WebappIntegrationMutationResponseModelName datamodel.ModelNameType = "agent.WebappIntegrationMutationResponse"
)

// WebappIntegrationMutationResponse
type WebappIntegrationMutationResponse struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Error
	Error string `json:"error" codec:"error" bson:"error" yaml:"error" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*WebappIntegrationMutationResponse)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*WebappIntegrationMutationResponse)(nil)

func toWebappIntegrationMutationResponseObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *WebappIntegrationMutationResponse:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of WebappIntegrationMutationResponse
func (o *WebappIntegrationMutationResponse) String() string {
	return fmt.Sprintf("agent.WebappIntegrationMutationResponse<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *WebappIntegrationMutationResponse) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *WebappIntegrationMutationResponse) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *WebappIntegrationMutationResponse) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *WebappIntegrationMutationResponse) GetModelName() datamodel.ModelNameType {
	return WebappIntegrationMutationResponseModelName
}

// NewWebappIntegrationMutationResponseID provides a template for generating an ID field for WebappIntegrationMutationResponse
func NewWebappIntegrationMutationResponseID(customerID string, refType string, refID string) string {
	return hash.Values("WebappIntegrationMutationResponse", customerID, refType, refID)
}

func (o *WebappIntegrationMutationResponse) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("WebappIntegrationMutationResponse", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *WebappIntegrationMutationResponse) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *WebappIntegrationMutationResponse) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *WebappIntegrationMutationResponse) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *WebappIntegrationMutationResponse) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *WebappIntegrationMutationResponse) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *WebappIntegrationMutationResponse) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *WebappIntegrationMutationResponse) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *WebappIntegrationMutationResponse) IsEvented() bool {
	return false
}

// GetTopicConfig returns the topic config object
func (o *WebappIntegrationMutationResponse) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *WebappIntegrationMutationResponse) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of WebappIntegrationMutationResponse
func (o *WebappIntegrationMutationResponse) Clone() datamodel.Model {
	c := new(WebappIntegrationMutationResponse)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *WebappIntegrationMutationResponse) Anon() datamodel.Model {
	c := new(WebappIntegrationMutationResponse)
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
func (o *WebappIntegrationMutationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *WebappIntegrationMutationResponse) UnmarshalJSON(data []byte) error {
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
func (o *WebappIntegrationMutationResponse) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two WebappIntegrationMutationResponse objects are equal
func (o *WebappIntegrationMutationResponse) IsEqual(other *WebappIntegrationMutationResponse) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *WebappIntegrationMutationResponse) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"customer_id": toWebappIntegrationMutationResponseObject(o.CustomerID, false),
		"error":       toWebappIntegrationMutationResponseObject(o.Error, false),
		"id":          toWebappIntegrationMutationResponseObject(o.ID, false),
		"ref_id":      toWebappIntegrationMutationResponseObject(o.RefID, false),
		"ref_type":    toWebappIntegrationMutationResponseObject(o.RefType, false),
		"hashcode":    toWebappIntegrationMutationResponseObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *WebappIntegrationMutationResponse) FromMap(kv map[string]interface{}) {

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

	if val, ok := kv["error"].(string); ok {
		o.Error = val
	} else {
		if val, ok := kv["error"]; ok {
			if val == nil {
				o.Error = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Error = fmt.Sprintf("%v", val)
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
func (o *WebappIntegrationMutationResponse) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CustomerID)
	args = append(args, o.Error)
	args = append(args, o.ID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *WebappIntegrationMutationResponse) GetEventAPIConfig() datamodel.EventAPIConfig {
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
