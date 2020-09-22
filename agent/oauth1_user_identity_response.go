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

	// Oauth1UserIdentityResponseTable is the default table name
	Oauth1UserIdentityResponseTable datamodel.ModelNameType = "agent_oauth1useridentityresponse"

	// Oauth1UserIdentityResponseModelName is the model name
	Oauth1UserIdentityResponseModelName datamodel.ModelNameType = "agent.Oauth1UserIdentityResponse"
)

const (
	// Oauth1UserIdentityResponseModelCustomerIDColumn is the column json value customer_id
	Oauth1UserIdentityResponseModelCustomerIDColumn = "customer_id"
	// Oauth1UserIdentityResponseModelIDColumn is the column json value id
	Oauth1UserIdentityResponseModelIDColumn = "id"
	// Oauth1UserIdentityResponseModelIdentityColumn is the column json value identity
	Oauth1UserIdentityResponseModelIdentityColumn = "identity"
	// Oauth1UserIdentityResponseModelIdentityAvatarURLColumn is the column json value avatar_url
	Oauth1UserIdentityResponseModelIdentityAvatarURLColumn = "avatar_url"
	// Oauth1UserIdentityResponseModelIdentityEmailColumn is the column json value email
	Oauth1UserIdentityResponseModelIdentityEmailColumn = "email"
	// Oauth1UserIdentityResponseModelIdentityNameColumn is the column json value name
	Oauth1UserIdentityResponseModelIdentityNameColumn = "name"
	// Oauth1UserIdentityResponseModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	Oauth1UserIdentityResponseModelIntegrationInstanceIDColumn = "integration_instance_id"
	// Oauth1UserIdentityResponseModelMessageColumn is the column json value message
	Oauth1UserIdentityResponseModelMessageColumn = "message"
	// Oauth1UserIdentityResponseModelRefIDColumn is the column json value ref_id
	Oauth1UserIdentityResponseModelRefIDColumn = "ref_id"
	// Oauth1UserIdentityResponseModelRefTypeColumn is the column json value ref_type
	Oauth1UserIdentityResponseModelRefTypeColumn = "ref_type"
	// Oauth1UserIdentityResponseModelSuccessColumn is the column json value success
	Oauth1UserIdentityResponseModelSuccessColumn = "success"
)

// Oauth1UserIdentityResponseIdentity represents the object structure for identity
type Oauth1UserIdentityResponseIdentity struct {
	// AvatarURL The avatar if one exists
	AvatarURL *string `json:"avatar_url,omitempty" codec:"avatar_url,omitempty" bson:"avatar_url" yaml:"avatar_url,omitempty" faker:"-"`
	// Email The verified email if one exists
	Email *string `json:"email,omitempty" codec:"email,omitempty" bson:"email" yaml:"email,omitempty" faker:"-"`
	// Name The display name of the user
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
}

func toOauth1UserIdentityResponseIdentityObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *Oauth1UserIdentityResponseIdentity:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *Oauth1UserIdentityResponseIdentity) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// AvatarURL The avatar if one exists
		"avatar_url": toOauth1UserIdentityResponseIdentityObject(o.AvatarURL, true),
		// Email The verified email if one exists
		"email": toOauth1UserIdentityResponseIdentityObject(o.Email, true),
		// Name The display name of the user
		"name": toOauth1UserIdentityResponseIdentityObject(o.Name, false),
	}
}

func (o *Oauth1UserIdentityResponseIdentity) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *Oauth1UserIdentityResponseIdentity) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["avatar_url"].(*string); ok {
		o.AvatarURL = val
	} else if val, ok := kv["avatar_url"].(string); ok {
		o.AvatarURL = &val
	} else {
		if val, ok := kv["avatar_url"]; ok {
			if val == nil {
				o.AvatarURL = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.AvatarURL = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["email"].(*string); ok {
		o.Email = val
	} else if val, ok := kv["email"].(string); ok {
		o.Email = &val
	} else {
		if val, ok := kv["email"]; ok {
			if val == nil {
				o.Email = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Email = pstrings.Pointer(fmt.Sprintf("%v", val))
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
	o.setDefaults(false)
}

// Oauth1UserIdentityResponse An OAuth1 response from the agent with user identity details
type Oauth1UserIdentityResponse struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// Identity the identity if successful
	Identity *Oauth1UserIdentityResponseIdentity `json:"identity,omitempty" codec:"identity,omitempty" bson:"identity" yaml:"identity,omitempty" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// Message if the request was not successful the error message
	Message *string `json:"message,omitempty" codec:"message,omitempty" bson:"message" yaml:"message,omitempty" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// Success if the request was successful or not
	Success bool `json:"success" codec:"success" bson:"success" yaml:"success" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Oauth1UserIdentityResponse)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*Oauth1UserIdentityResponse)(nil)

func toOauth1UserIdentityResponseObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *Oauth1UserIdentityResponse:
		return v.ToMap()

	case *Oauth1UserIdentityResponseIdentity:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of Oauth1UserIdentityResponse
func (o *Oauth1UserIdentityResponse) String() string {
	return fmt.Sprintf("agent.Oauth1UserIdentityResponse<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Oauth1UserIdentityResponse) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *Oauth1UserIdentityResponse) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *Oauth1UserIdentityResponse) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *Oauth1UserIdentityResponse) GetModelName() datamodel.ModelNameType {
	return Oauth1UserIdentityResponseModelName
}

// NewOauth1UserIdentityResponseID provides a template for generating an ID field for Oauth1UserIdentityResponse
func NewOauth1UserIdentityResponseID(customerID string, refType string, refID string) string {
	return hash.Values("Oauth1UserIdentityResponse", customerID, refType, refID)
}

func (o *Oauth1UserIdentityResponse) setDefaults(frommap bool) {
	if o.Identity == nil {
		o.Identity = &Oauth1UserIdentityResponseIdentity{}
	}

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Oauth1UserIdentityResponse", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Oauth1UserIdentityResponse) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Oauth1UserIdentityResponse) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Oauth1UserIdentityResponse) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *Oauth1UserIdentityResponse) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Oauth1UserIdentityResponse) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *Oauth1UserIdentityResponse) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Oauth1UserIdentityResponse) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Oauth1UserIdentityResponse) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *Oauth1UserIdentityResponse) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = Oauth1UserIdentityResponseModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Oauth1UserIdentityResponse) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *Oauth1UserIdentityResponse) GetCustomerID() string {
	return o.CustomerID
}

// SetCustomerID will return the customer_id
func (o *Oauth1UserIdentityResponse) SetCustomerID(id string) {
	o.CustomerID = id
}

// GetRefType will return the ref_type
func (o *Oauth1UserIdentityResponse) GetRefType() string {
	return o.RefType
}

// SetRefType will return the ref_type
func (o *Oauth1UserIdentityResponse) SetRefType(t string) {
	o.RefType = t
}

// Clone returns an exact copy of Oauth1UserIdentityResponse
func (o *Oauth1UserIdentityResponse) Clone() datamodel.Model {
	c := new(Oauth1UserIdentityResponse)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Oauth1UserIdentityResponse) Anon() datamodel.Model {
	c := new(Oauth1UserIdentityResponse)
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
func (o *Oauth1UserIdentityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Oauth1UserIdentityResponse) UnmarshalJSON(data []byte) error {
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
func (o *Oauth1UserIdentityResponse) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *Oauth1UserIdentityResponse) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two Oauth1UserIdentityResponse objects are equal
func (o *Oauth1UserIdentityResponse) IsEqual(other *Oauth1UserIdentityResponse) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Oauth1UserIdentityResponse) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"customer_id":             toOauth1UserIdentityResponseObject(o.CustomerID, false),
		"id":                      toOauth1UserIdentityResponseObject(o.ID, false),
		"identity":                toOauth1UserIdentityResponseObject(o.Identity, true),
		"integration_instance_id": toOauth1UserIdentityResponseObject(o.IntegrationInstanceID, true),
		"message":                 toOauth1UserIdentityResponseObject(o.Message, true),
		"ref_id":                  toOauth1UserIdentityResponseObject(o.RefID, false),
		"ref_type":                toOauth1UserIdentityResponseObject(o.RefType, false),
		"success":                 toOauth1UserIdentityResponseObject(o.Success, false),
		"hashcode":                toOauth1UserIdentityResponseObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *Oauth1UserIdentityResponse) FromMap(kv map[string]interface{}) {

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

	if o.Identity == nil {
		o.Identity = &Oauth1UserIdentityResponseIdentity{}
	}

	if val, ok := kv["identity"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.Identity.FromMap(kv)
		} else if sv, ok := val.(Oauth1UserIdentityResponseIdentity); ok {
			// struct
			o.Identity = &sv
		} else if sp, ok := val.(*Oauth1UserIdentityResponseIdentity); ok {
			// struct pointer
			o.Identity = sp
		}
	} else {
		o.Identity.FromMap(map[string]interface{}{})
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
	if val, ok := kv["message"].(*string); ok {
		o.Message = val
	} else if val, ok := kv["message"].(string); ok {
		o.Message = &val
	} else {
		if val, ok := kv["message"]; ok {
			if val == nil {
				o.Message = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Message = pstrings.Pointer(fmt.Sprintf("%v", val))
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
func (o *Oauth1UserIdentityResponse) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.Identity)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.Message)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.Success)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// SetIntegrationInstanceID will set the integration instance ID
func (o *Oauth1UserIdentityResponse) SetIntegrationInstanceID(id string) {
	if id == "" {
		o.IntegrationInstanceID = nil
	} else {
		o.IntegrationInstanceID = &id
	}
}

// GetIntegrationInstanceID will return the integration instance ID
func (o *Oauth1UserIdentityResponse) GetIntegrationInstanceID() *string {
	return o.IntegrationInstanceID
}

// Oauth1UserIdentityResponsePartial is a partial struct for upsert mutations for Oauth1UserIdentityResponse
type Oauth1UserIdentityResponsePartial struct {
	// Identity the identity if successful
	Identity *Oauth1UserIdentityResponseIdentity `json:"identity,omitempty"`
	// Message if the request was not successful the error message
	Message *string `json:"message,omitempty"`
	// Success if the request was successful or not
	Success *bool `json:"success,omitempty"`
}

var _ datamodel.PartialModel = (*Oauth1UserIdentityResponsePartial)(nil)

// GetModelName returns the name of the model
func (o *Oauth1UserIdentityResponsePartial) GetModelName() datamodel.ModelNameType {
	return Oauth1UserIdentityResponseModelName
}

// ToMap returns the object as a map
func (o *Oauth1UserIdentityResponsePartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"identity": toOauth1UserIdentityResponseObject(o.Identity, true),
		"message":  toOauth1UserIdentityResponseObject(o.Message, true),
		"success":  toOauth1UserIdentityResponseObject(o.Success, true),
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
func (o *Oauth1UserIdentityResponsePartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *Oauth1UserIdentityResponsePartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Oauth1UserIdentityResponsePartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *Oauth1UserIdentityResponsePartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *Oauth1UserIdentityResponsePartial) FromMap(kv map[string]interface{}) {

	if o.Identity == nil {
		o.Identity = &Oauth1UserIdentityResponseIdentity{}
	}

	if val, ok := kv["identity"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.Identity.FromMap(kv)
		} else if sv, ok := val.(Oauth1UserIdentityResponseIdentity); ok {
			// struct
			o.Identity = &sv
		} else if sp, ok := val.(*Oauth1UserIdentityResponseIdentity); ok {
			// struct pointer
			o.Identity = sp
		}
	} else {
		o.Identity.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["message"].(*string); ok {
		o.Message = val
	} else if val, ok := kv["message"].(string); ok {
		o.Message = &val
	} else {
		if val, ok := kv["message"]; ok {
			if val == nil {
				o.Message = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Message = pstrings.Pointer(fmt.Sprintf("%v", val))
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
