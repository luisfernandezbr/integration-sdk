// DO NOT EDIT -- generated code

// Package sourcecode - the system which contains source code
package sourcecode

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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

const (
	// UserTopic is the default topic name
	UserTopic datamodel.TopicNameType = "sourcecode_User"

	// UserTable is the default table name
	UserTable datamodel.ModelNameType = "sourcecode_user"

	// UserModelName is the model name
	UserModelName datamodel.ModelNameType = "sourcecode.User"
)

const (
	// UserModelAssociatedRefIDColumn is the column json value associated_ref_id
	UserModelAssociatedRefIDColumn = "associated_ref_id"
	// UserModelAvatarURLColumn is the column json value avatar_url
	UserModelAvatarURLColumn = "avatar_url"
	// UserModelCustomerIDColumn is the column json value customer_id
	UserModelCustomerIDColumn = "customer_id"
	// UserModelEmailColumn is the column json value email
	UserModelEmailColumn = "email"
	// UserModelIDColumn is the column json value id
	UserModelIDColumn = "id"
	// UserModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	UserModelIntegrationInstanceIDColumn = "integration_instance_id"
	// UserModelMemberColumn is the column json value member
	UserModelMemberColumn = "member"
	// UserModelNameColumn is the column json value name
	UserModelNameColumn = "name"
	// UserModelRefIDColumn is the column json value ref_id
	UserModelRefIDColumn = "ref_id"
	// UserModelRefTypeColumn is the column json value ref_type
	UserModelRefTypeColumn = "ref_type"
	// UserModelTypeColumn is the column json value type
	UserModelTypeColumn = "type"
	// UserModelUpdatedAtColumn is the column json value updated_ts
	UserModelUpdatedAtColumn = "updated_ts"
	// UserModelURLColumn is the column json value url
	UserModelURLColumn = "url"
	// UserModelUsernameColumn is the column json value username
	UserModelUsernameColumn = "username"
)

// UserType is the enumeration type for type
type UserType int32

// toUserTypePointer is the enumeration pointer type for type
func toUserTypePointer(v int32) *UserType {
	nv := UserType(v)
	return &nv
}

// toUserTypeEnum is the enumeration pointer wrapper for type
func toUserTypeEnum(v *UserType) string {
	if v == nil {
		return toUserTypePointer(0).String()
	}
	return v.String()
}

// UnmarshalBSONValue for unmarshaling value
func (v *UserType) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = UserType(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "HUMAN":
			*v = UserType(0)
		case "BOT":
			*v = UserType(1)
		case "DELETED_SPECIAL_USER":
			*v = UserType(2)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v *UserType) UnmarshalJSON(buf []byte) error {
	var val string
	if err := json.Unmarshal(buf, &val); err != nil {
		return err
	}
	switch val {
	case "HUMAN":
		*v = 0
	case "BOT":
		*v = 1
	case "DELETED_SPECIAL_USER":
		*v = 2
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v UserType) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("HUMAN")
	case 1:
		return json.Marshal("BOT")
	case 2:
		return json.Marshal("DELETED_SPECIAL_USER")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for Type
func (v UserType) String() string {
	switch int32(v) {
	case 0:
		return "HUMAN"
	case 1:
		return "BOT"
	case 2:
		return "DELETED_SPECIAL_USER"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *UserType) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case UserType:
		*v = val
	case int32:
		*v = UserType(int32(val))
	case int:
		*v = UserType(int32(val))
	case string:
		switch val {
		case "HUMAN":
			*v = UserType(0)
		case "BOT":
			*v = UserType(1)
		case "DELETED_SPECIAL_USER":
			*v = UserType(2)
		}
	}
	return nil
}

const (
	// UserTypeHuman is the enumeration value for human
	UserTypeHuman UserType = 0
	// UserTypeBot is the enumeration value for bot
	UserTypeBot UserType = 1
	// UserTypeDeletedSpecialUser is the enumeration value for deleted_special_user
	UserTypeDeletedSpecialUser UserType = 2
)

// User the source code user
type User struct {
	// AssociatedRefID the ref id associated for this user in another system
	AssociatedRefID *string `json:"associated_ref_id,omitempty" codec:"associated_ref_id,omitempty" bson:"associated_ref_id" yaml:"associated_ref_id,omitempty" faker:"-"`
	// AvatarURL the url to users avatar
	AvatarURL *string `json:"avatar_url,omitempty" codec:"avatar_url,omitempty" bson:"avatar_url" yaml:"avatar_url,omitempty" faker:"avatar"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Email the email for the user
	Email *string `json:"email,omitempty" codec:"email,omitempty" bson:"email" yaml:"email,omitempty" faker:"email"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// Member if the user is a member of organization
	//
	// Deprecated: no longer used
	Member bool `json:"member" codec:"member" bson:"member" yaml:"member" faker:"-"`
	// Name the name of the user
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"person"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// Type type of the user
	Type UserType `json:"type" codec:"type" bson:"type" yaml:"type" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// URL the url to the user's page
	URL *string `json:"url,omitempty" codec:"url,omitempty" bson:"url" yaml:"url,omitempty" faker:"-"`
	// Username username of the user
	Username *string `json:"username,omitempty" codec:"username,omitempty" bson:"username" yaml:"username,omitempty" faker:"username"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*User)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*User)(nil)

func toUserObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *User:
		return v.ToMap()

	case UserType:
		return v.String()

	default:
		return o
	}
}

// String returns a string representation of User
func (o *User) String() string {
	return fmt.Sprintf("sourcecode.User<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *User) GetTopicName() datamodel.TopicNameType {
	return UserTopic
}

// GetStreamName returns the name of the stream
func (o *User) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *User) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *User) GetModelName() datamodel.ModelNameType {
	return UserModelName
}

// NewUserID provides a template for generating an ID field for User
func NewUserID(customerID string, refType string, refID string) string {
	return hash.Values("User", customerID, refType, refID)
}

func (o *User) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("User", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *User) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *User) GetTopicKey() string {
	var i interface{} = o.ID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *User) GetTimestamp() time.Time {
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
	panic("not sure how to handle the date time format for User")
}

// GetRefID returns the RefID for the object
func (o *User) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *User) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *User) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *User) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *User) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *User) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = UserModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *User) GetTopicConfig() *datamodel.ModelTopicConfig {
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
		NumPartitions:     128,
		CleanupPolicy:     datamodel.CleanupPolicy("compact"),
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetCustomerID will return the customer_id
func (o *User) GetCustomerID() string {
	return o.CustomerID
}

// SetCustomerID will return the customer_id
func (o *User) SetCustomerID(id string) {
	o.CustomerID = id
}

// GetRefType will return the ref_type
func (o *User) GetRefType() string {
	return o.RefType
}

// SetRefType will return the ref_type
func (o *User) SetRefType(t string) {
	o.RefType = t
}

// Clone returns an exact copy of User
func (o *User) Clone() datamodel.Model {
	c := new(User)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *User) Anon() datamodel.Model {
	c := new(User)
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
func (o *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *User) UnmarshalJSON(data []byte) error {
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
func (o *User) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *User) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two User objects are equal
func (o *User) IsEqual(other *User) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *User) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"associated_ref_id":       toUserObject(o.AssociatedRefID, true),
		"avatar_url":              toUserObject(o.AvatarURL, true),
		"customer_id":             toUserObject(o.CustomerID, false),
		"email":                   toUserObject(o.Email, true),
		"id":                      toUserObject(o.ID, false),
		"integration_instance_id": toUserObject(o.IntegrationInstanceID, true),
		"member":                  toUserObject(o.Member, false),
		"name":                    toUserObject(o.Name, false),
		"ref_id":                  toUserObject(o.RefID, false),
		"ref_type":                toUserObject(o.RefType, false),

		"type":       o.Type.String(),
		"updated_ts": toUserObject(o.UpdatedAt, false),
		"url":        toUserObject(o.URL, true),
		"username":   toUserObject(o.Username, true),
		"hashcode":   toUserObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *User) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["associated_ref_id"].(*string); ok {
		o.AssociatedRefID = val
	} else if val, ok := kv["associated_ref_id"].(string); ok {
		o.AssociatedRefID = &val
	} else {
		if val, ok := kv["associated_ref_id"]; ok {
			if val == nil {
				o.AssociatedRefID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.AssociatedRefID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
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
	// Deprecated
	if val, ok := kv["member"].(bool); ok {
		o.Member = val
	} else {
		if val, ok := kv["member"]; ok {
			if val == nil {
				o.Member = false
			} else {
				o.Member = number.ToBoolAny(val)
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
	if val, ok := kv["type"].(UserType); ok {
		o.Type = val
	} else {
		if em, ok := kv["type"].(map[string]interface{}); ok {

			ev := em["sourcecode.type"].(string)
			switch ev {
			case "human", "HUMAN":
				o.Type = 0
			case "bot", "BOT":
				o.Type = 1
			case "deleted_special_user", "DELETED_SPECIAL_USER":
				o.Type = 2
			}
		}
		if em, ok := kv["type"].(string); ok {
			switch em {
			case "human", "HUMAN":
				o.Type = 0
			case "bot", "BOT":
				o.Type = 1
			case "deleted_special_user", "DELETED_SPECIAL_USER":
				o.Type = 2
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
	if val, ok := kv["username"].(*string); ok {
		o.Username = val
	} else if val, ok := kv["username"].(string); ok {
		o.Username = &val
	} else {
		if val, ok := kv["username"]; ok {
			if val == nil {
				o.Username = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Username = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *User) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.AssociatedRefID)
	args = append(args, o.AvatarURL)
	args = append(args, o.CustomerID)
	args = append(args, o.Email)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.Member)
	args = append(args, o.Name)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.Type)
	args = append(args, o.UpdatedAt)
	args = append(args, o.URL)
	args = append(args, o.Username)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// SetIntegrationInstanceID will set the integration instance ID
func (o *User) SetIntegrationInstanceID(id string) {
	if id == "" {
		o.IntegrationInstanceID = nil
	} else {
		o.IntegrationInstanceID = &id
	}
}

// GetIntegrationInstanceID will return the integration instance ID
func (o *User) GetIntegrationInstanceID() *string {
	return o.IntegrationInstanceID
}

// UserPartial is a partial struct for upsert mutations for User
type UserPartial struct {
	// AssociatedRefID the ref id associated for this user in another system
	AssociatedRefID *string `json:"associated_ref_id,omitempty"`
	// AvatarURL the url to users avatar
	AvatarURL *string `json:"avatar_url,omitempty"`
	// Email the email for the user
	Email *string `json:"email,omitempty"`
	// Member if the user is a member of organization
	//
	// Deprecated: no longer used
	Member *bool `json:"member,omitempty"`
	// Name the name of the user
	Name *string `json:"name,omitempty"`
	// Type type of the user
	Type *UserType `json:"type,omitempty"`
	// URL the url to the user's page
	URL *string `json:"url,omitempty"`
	// Username username of the user
	Username *string `json:"username,omitempty"`
}

var _ datamodel.PartialModel = (*UserPartial)(nil)

// GetModelName returns the name of the model
func (o *UserPartial) GetModelName() datamodel.ModelNameType {
	return UserModelName
}

// ToMap returns the object as a map
func (o *UserPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"associated_ref_id": toUserObject(o.AssociatedRefID, true),
		"avatar_url":        toUserObject(o.AvatarURL, true),
		"email":             toUserObject(o.Email, true),
		"member":            toUserObject(o.Member, true),
		"name":              toUserObject(o.Name, true),

		"type":     toUserTypeEnum(o.Type),
		"url":      toUserObject(o.URL, true),
		"username": toUserObject(o.Username, true),
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
func (o *UserPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *UserPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *UserPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *UserPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *UserPartial) FromMap(kv map[string]interface{}) {
	if val, ok := kv["associated_ref_id"].(*string); ok {
		o.AssociatedRefID = val
	} else if val, ok := kv["associated_ref_id"].(string); ok {
		o.AssociatedRefID = &val
	} else {
		if val, ok := kv["associated_ref_id"]; ok {
			if val == nil {
				o.AssociatedRefID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.AssociatedRefID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
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
	// Deprecated
	if val, ok := kv["member"].(*bool); ok {
		o.Member = val
	} else if val, ok := kv["member"].(bool); ok {
		o.Member = &val
	} else {
		if val, ok := kv["member"]; ok {
			if val == nil {
				o.Member = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Member = number.BoolPointer(number.ToBoolAny(val))
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
	if val, ok := kv["type"].(*UserType); ok {
		o.Type = val
	} else if val, ok := kv["type"].(UserType); ok {
		o.Type = &val
	} else {
		if val, ok := kv["type"]; ok {
			if val == nil {
				o.Type = toUserTypePointer(0)
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["UserType"]
				}
				// this is an enum pointer
				if em, ok := val.(string); ok {
					switch em {
					case "human", "HUMAN":
						o.Type = toUserTypePointer(0)
					case "bot", "BOT":
						o.Type = toUserTypePointer(1)
					case "deleted_special_user", "DELETED_SPECIAL_USER":
						o.Type = toUserTypePointer(2)
					}
				}
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
	if val, ok := kv["username"].(*string); ok {
		o.Username = val
	} else if val, ok := kv["username"].(string); ok {
		o.Username = &val
	} else {
		if val, ok := kv["username"]; ok {
			if val == nil {
				o.Username = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Username = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	o.setDefaults(false)
}
