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
	pstrings "github.com/pinpt/go-common/v10/strings"
)

const (

	// Oauth1UserIdentityRequestTable is the default table name
	Oauth1UserIdentityRequestTable datamodel.ModelNameType = "agent_oauth1useridentityrequest"

	// Oauth1UserIdentityRequestModelName is the model name
	Oauth1UserIdentityRequestModelName datamodel.ModelNameType = "agent.Oauth1UserIdentityRequest"
)

const (
	// Oauth1UserIdentityRequestModelConsumerKeyColumn is the column json value consumer_key
	Oauth1UserIdentityRequestModelConsumerKeyColumn = "consumer_key"
	// Oauth1UserIdentityRequestModelConsumerSecretColumn is the column json value consumer_secret
	Oauth1UserIdentityRequestModelConsumerSecretColumn = "consumer_secret"
	// Oauth1UserIdentityRequestModelCustomerIDColumn is the column json value customer_id
	Oauth1UserIdentityRequestModelCustomerIDColumn = "customer_id"
	// Oauth1UserIdentityRequestModelEnrollmentIDColumn is the column json value enrollment_id
	Oauth1UserIdentityRequestModelEnrollmentIDColumn = "enrollment_id"
	// Oauth1UserIdentityRequestModelIDColumn is the column json value id
	Oauth1UserIdentityRequestModelIDColumn = "id"
	// Oauth1UserIdentityRequestModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	Oauth1UserIdentityRequestModelIntegrationInstanceIDColumn = "integration_instance_id"
	// Oauth1UserIdentityRequestModelPrivateKeyColumn is the column json value private_key
	Oauth1UserIdentityRequestModelPrivateKeyColumn = "private_key"
	// Oauth1UserIdentityRequestModelRefIDColumn is the column json value ref_id
	Oauth1UserIdentityRequestModelRefIDColumn = "ref_id"
	// Oauth1UserIdentityRequestModelRefTypeColumn is the column json value ref_type
	Oauth1UserIdentityRequestModelRefTypeColumn = "ref_type"
	// Oauth1UserIdentityRequestModelURLColumn is the column json value url
	Oauth1UserIdentityRequestModelURLColumn = "url"
)

// Oauth1UserIdentityRequest An OAuth1 request sent to the agent to get user identity details
type Oauth1UserIdentityRequest struct {
	// ConsumerKey The Consumer Key
	ConsumerKey string `json:"consumer_key" codec:"consumer_key" bson:"consumer_key" yaml:"consumer_key" faker:"-"`
	// ConsumerSecret The Consumer Secret
	ConsumerSecret string `json:"consumer_secret" codec:"consumer_secret" bson:"consumer_secret" yaml:"consumer_secret" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// EnrollmentID The id of the on premise agent enrollment to route this to. Leave null for cloud agent.
	EnrollmentID *string `json:"enrollment_id,omitempty" codec:"enrollment_id,omitempty" bson:"enrollment_id" yaml:"enrollment_id,omitempty" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// PrivateKey A private key for signing the oauth request.
	PrivateKey string `json:"private_key" codec:"private_key" bson:"private_key" yaml:"private_key" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// URL The URL to the endpoint for fetching the data
	URL string `json:"url" codec:"url" bson:"url" yaml:"url" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Oauth1UserIdentityRequest)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*Oauth1UserIdentityRequest)(nil)

func toOauth1UserIdentityRequestObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *Oauth1UserIdentityRequest:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of Oauth1UserIdentityRequest
func (o *Oauth1UserIdentityRequest) String() string {
	return fmt.Sprintf("agent.Oauth1UserIdentityRequest<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Oauth1UserIdentityRequest) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *Oauth1UserIdentityRequest) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *Oauth1UserIdentityRequest) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *Oauth1UserIdentityRequest) GetModelName() datamodel.ModelNameType {
	return Oauth1UserIdentityRequestModelName
}

// NewOauth1UserIdentityRequestID provides a template for generating an ID field for Oauth1UserIdentityRequest
func NewOauth1UserIdentityRequestID(customerID string, refType string, refID string) string {
	return hash.Values("Oauth1UserIdentityRequest", customerID, refType, refID)
}

func (o *Oauth1UserIdentityRequest) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Oauth1UserIdentityRequest", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Oauth1UserIdentityRequest) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Oauth1UserIdentityRequest) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Oauth1UserIdentityRequest) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *Oauth1UserIdentityRequest) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Oauth1UserIdentityRequest) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *Oauth1UserIdentityRequest) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Oauth1UserIdentityRequest) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Oauth1UserIdentityRequest) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *Oauth1UserIdentityRequest) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = Oauth1UserIdentityRequestModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Oauth1UserIdentityRequest) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *Oauth1UserIdentityRequest) GetCustomerID() string {
	return o.CustomerID
}

// SetCustomerID will return the customer_id
func (o *Oauth1UserIdentityRequest) SetCustomerID(id string) {
	o.CustomerID = id
}

// GetRefType will return the ref_type
func (o *Oauth1UserIdentityRequest) GetRefType() string {
	return o.RefType
}

// SetRefType will return the ref_type
func (o *Oauth1UserIdentityRequest) SetRefType(t string) {
	o.RefType = t
}

// Clone returns an exact copy of Oauth1UserIdentityRequest
func (o *Oauth1UserIdentityRequest) Clone() datamodel.Model {
	c := new(Oauth1UserIdentityRequest)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Oauth1UserIdentityRequest) Anon() datamodel.Model {
	c := new(Oauth1UserIdentityRequest)
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
func (o *Oauth1UserIdentityRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Oauth1UserIdentityRequest) UnmarshalJSON(data []byte) error {
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
func (o *Oauth1UserIdentityRequest) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *Oauth1UserIdentityRequest) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two Oauth1UserIdentityRequest objects are equal
func (o *Oauth1UserIdentityRequest) IsEqual(other *Oauth1UserIdentityRequest) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Oauth1UserIdentityRequest) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"consumer_key":            toOauth1UserIdentityRequestObject(o.ConsumerKey, false),
		"consumer_secret":         toOauth1UserIdentityRequestObject(o.ConsumerSecret, false),
		"customer_id":             toOauth1UserIdentityRequestObject(o.CustomerID, false),
		"enrollment_id":           toOauth1UserIdentityRequestObject(o.EnrollmentID, true),
		"id":                      toOauth1UserIdentityRequestObject(o.ID, false),
		"integration_instance_id": toOauth1UserIdentityRequestObject(o.IntegrationInstanceID, true),
		"private_key":             toOauth1UserIdentityRequestObject(o.PrivateKey, false),
		"ref_id":                  toOauth1UserIdentityRequestObject(o.RefID, false),
		"ref_type":                toOauth1UserIdentityRequestObject(o.RefType, false),
		"url":                     toOauth1UserIdentityRequestObject(o.URL, false),
		"hashcode":                toOauth1UserIdentityRequestObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *Oauth1UserIdentityRequest) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["consumer_key"].(string); ok {
		o.ConsumerKey = val
	} else {
		if val, ok := kv["consumer_key"]; ok {
			if val == nil {
				o.ConsumerKey = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.ConsumerKey = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["consumer_secret"].(string); ok {
		o.ConsumerSecret = val
	} else {
		if val, ok := kv["consumer_secret"]; ok {
			if val == nil {
				o.ConsumerSecret = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.ConsumerSecret = fmt.Sprintf("%v", val)
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
	if val, ok := kv["enrollment_id"].(*string); ok {
		o.EnrollmentID = val
	} else if val, ok := kv["enrollment_id"].(string); ok {
		o.EnrollmentID = &val
	} else {
		if val, ok := kv["enrollment_id"]; ok {
			if val == nil {
				o.EnrollmentID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.EnrollmentID = pstrings.Pointer(fmt.Sprintf("%v", val))
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
	if val, ok := kv["private_key"].(string); ok {
		o.PrivateKey = val
	} else {
		if val, ok := kv["private_key"]; ok {
			if val == nil {
				o.PrivateKey = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.PrivateKey = fmt.Sprintf("%v", val)
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
	if val, ok := kv["url"].(string); ok {
		o.URL = val
	} else {
		if val, ok := kv["url"]; ok {
			if val == nil {
				o.URL = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.URL = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *Oauth1UserIdentityRequest) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.ConsumerKey)
	args = append(args, o.ConsumerSecret)
	args = append(args, o.CustomerID)
	args = append(args, o.EnrollmentID)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.PrivateKey)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.URL)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// SetIntegrationInstanceID will set the integration instance ID
func (o *Oauth1UserIdentityRequest) SetIntegrationInstanceID(id string) {
	if id == "" {
		o.IntegrationInstanceID = nil
	} else {
		o.IntegrationInstanceID = &id
	}
}

// GetIntegrationInstanceID will return the integration instance ID
func (o *Oauth1UserIdentityRequest) GetIntegrationInstanceID() *string {
	return o.IntegrationInstanceID
}

// Oauth1UserIdentityRequestPartial is a partial struct for upsert mutations for Oauth1UserIdentityRequest
type Oauth1UserIdentityRequestPartial struct {
	// ConsumerKey The Consumer Key
	ConsumerKey *string `json:"consumer_key,omitempty"`
	// ConsumerSecret The Consumer Secret
	ConsumerSecret *string `json:"consumer_secret,omitempty"`
	// EnrollmentID The id of the on premise agent enrollment to route this to. Leave null for cloud agent.
	EnrollmentID *string `json:"enrollment_id,omitempty"`
	// PrivateKey A private key for signing the oauth request.
	PrivateKey *string `json:"private_key,omitempty"`
	// URL The URL to the endpoint for fetching the data
	URL *string `json:"url,omitempty"`
}

var _ datamodel.PartialModel = (*Oauth1UserIdentityRequestPartial)(nil)

// GetModelName returns the name of the model
func (o *Oauth1UserIdentityRequestPartial) GetModelName() datamodel.ModelNameType {
	return Oauth1UserIdentityRequestModelName
}

// ToMap returns the object as a map
func (o *Oauth1UserIdentityRequestPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"consumer_key":    toOauth1UserIdentityRequestObject(o.ConsumerKey, true),
		"consumer_secret": toOauth1UserIdentityRequestObject(o.ConsumerSecret, true),
		"enrollment_id":   toOauth1UserIdentityRequestObject(o.EnrollmentID, true),
		"private_key":     toOauth1UserIdentityRequestObject(o.PrivateKey, true),
		"url":             toOauth1UserIdentityRequestObject(o.URL, true),
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
func (o *Oauth1UserIdentityRequestPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *Oauth1UserIdentityRequestPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Oauth1UserIdentityRequestPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *Oauth1UserIdentityRequestPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *Oauth1UserIdentityRequestPartial) FromMap(kv map[string]interface{}) {
	if val, ok := kv["consumer_key"].(*string); ok {
		o.ConsumerKey = val
	} else if val, ok := kv["consumer_key"].(string); ok {
		o.ConsumerKey = &val
	} else {
		if val, ok := kv["consumer_key"]; ok {
			if val == nil {
				o.ConsumerKey = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.ConsumerKey = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["consumer_secret"].(*string); ok {
		o.ConsumerSecret = val
	} else if val, ok := kv["consumer_secret"].(string); ok {
		o.ConsumerSecret = &val
	} else {
		if val, ok := kv["consumer_secret"]; ok {
			if val == nil {
				o.ConsumerSecret = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.ConsumerSecret = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["enrollment_id"].(*string); ok {
		o.EnrollmentID = val
	} else if val, ok := kv["enrollment_id"].(string); ok {
		o.EnrollmentID = &val
	} else {
		if val, ok := kv["enrollment_id"]; ok {
			if val == nil {
				o.EnrollmentID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.EnrollmentID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["private_key"].(*string); ok {
		o.PrivateKey = val
	} else if val, ok := kv["private_key"].(string); ok {
		o.PrivateKey = &val
	} else {
		if val, ok := kv["private_key"]; ok {
			if val == nil {
				o.PrivateKey = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.PrivateKey = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["url"].(*string); ok {
		o.URL = val
	} else if val, ok := kv["url"].(string); ok {
		o.URL = &val
	} else {
		if val, ok := kv["url"]; ok {
			if val == nil {
				o.URL = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.URL = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	o.setDefaults(false)
}
