// DO NOT EDIT -- generated code

// Package customer - customer specific data
package customer

import (
	"bufio"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"time"

	"github.com/bxcodec/faker"
	"github.com/linkedin/goavro"
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/eventing"
	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
	pstrings "github.com/pinpt/go-common/strings"
)

const (
	// UserTopic is the default topic name
	UserTopic datamodel.TopicNameType = "customer_User_topic"

	// UserStream is the default stream name
	UserStream datamodel.TopicNameType = "customer_User_stream"

	// UserTable is the default table name
	UserTable datamodel.TopicNameType = "customer_User"

	// UserModelName is the model name
	UserModelName datamodel.ModelNameType = "customer.User"
)

const (
	// UserIDColumn is the id column name
	UserIDColumn = "id"
	// UserRefIDColumn is the ref_id column name
	UserRefIDColumn = "ref_id"
	// UserRefTypeColumn is the ref_type column name
	UserRefTypeColumn = "ref_type"
	// UserCustomerIDColumn is the customer_id column name
	UserCustomerIDColumn = "customer_id"
	// UserCreatedAtColumn is the created_ts column name
	UserCreatedAtColumn = "created_ts"
	// UserUpdatedAtColumn is the updated_ts column name
	UserUpdatedAtColumn = "updated_ts"
	// UserNameColumn is the name column name
	UserNameColumn = "name"
	// UserEmailColumn is the email column name
	UserEmailColumn = "email"
	// UserTitleColumn is the title column name
	UserTitleColumn = "title"
	// UserLocationColumn is the location column name
	UserLocationColumn = "location"
	// UserAvatarURLColumn is the avatar_url column name
	UserAvatarURLColumn = "avatar_url"
	// UserManagerIDColumn is the manager_id column name
	UserManagerIDColumn = "manager_id"
	// UserOwnerColumn is the owner column name
	UserOwnerColumn = "owner"
	// UserActiveColumn is the active column name
	UserActiveColumn = "active"
	// UserTrackableColumn is the trackable column name
	UserTrackableColumn = "trackable"
	// UserDeletedAtColumn is the deleted_ts column name
	UserDeletedAtColumn = "deleted_ts"
	// UserHiredAtColumn is the hired_ts column name
	UserHiredAtColumn = "hired_ts"
	// UserTerminatedAtColumn is the terminated_ts column name
	UserTerminatedAtColumn = "terminated_ts"
	// UserCostCenterIDColumn is the cost_center_id column name
	UserCostCenterIDColumn = "cost_center_id"
	// UserTeamIDColumn is the team_id column name
	UserTeamIDColumn = "team_id"
)

// User the customer's user record
type User struct {
	// built in types

	ID         string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	RefID      string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	RefType    string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	Hashcode   string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`

	CreatedAt int64 `json:"created_ts" bson:"created_ts" yaml:"created_ts" faker:"-"`
	UpdatedAt int64 `json:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`

	// custom types

	// Name name of the user
	Name string `json:"name" bson:"name" yaml:"name" faker:"person"`
	// Email the email of the user
	Email string `json:"email" bson:"email" yaml:"email" faker:"email"`
	// Title the title of the user
	Title *string `json:"title" bson:"title" yaml:"title" faker:"jobtitle"`
	// Location the location of the user
	Location *string `json:"location" bson:"location" yaml:"location" faker:"location"`
	// AvatarURL the user avatar url
	AvatarURL *string `json:"avatar_url" bson:"avatar_url" yaml:"avatar_url" faker:"avatar"`
	// ManagerID the manager user id
	ManagerID *string `json:"manager_id" bson:"manager_id" yaml:"manager_id" faker:"-"`
	// Owner if true, the user is an owner of the account
	Owner bool `json:"owner" bson:"owner" yaml:"owner" faker:"-"`
	// Active if true, the user is active and able to login
	Active bool `json:"active" bson:"active" yaml:"active" faker:"-"`
	// Trackable if true, the user is trackable in the pinpoint system
	Trackable bool `json:"trackable" bson:"trackable" yaml:"trackable" faker:"-"`
	// DeletedAt when the user record was deleted in epoch timestamp
	DeletedAt *int64 `json:"deleted_ts" bson:"deleted_ts" yaml:"deleted_ts" faker:"-"`
	// HiredAt when the user was hired in epoch timestamp
	HiredAt *int64 `json:"hired_ts" bson:"hired_ts" yaml:"hired_ts" faker:"-"`
	// TerminatedAt when the user was terminated in epoch timestamp
	TerminatedAt *int64 `json:"terminated_ts" bson:"terminated_ts" yaml:"terminated_ts" faker:"-"`
	// CostCenterID the id of the cost center
	CostCenterID *string `json:"cost_center_id" bson:"cost_center_id" yaml:"cost_center_id" faker:"-"`
	// TeamID the team id that the user is part of
	TeamID *string `json:"team_id" bson:"team_id" yaml:"team_id" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*User)(nil)

func toUserObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toUserObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if o == nil {
		return toUserObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toUserObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toUserObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toUserObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toUserObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toUserObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toUserObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toUserObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toUserObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toUserObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toUserObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case map[string]interface{}:
		return o
	case *map[string]interface{}:
		return v
	case map[string]string:
		return v
	case *map[string]string:
		return *v
	case *User:
		return v.ToMap()
	case User:
		return v.ToMap()
	case []string, []int64, []float64, []bool:
		return o
	case *[]string:
		return (*(o.(*[]string)))
	case *[]int64:
		return (*(o.(*[]int64)))
	case *[]float64:
		return (*(o.(*[]float64)))
	case *[]bool:
		return (*(o.(*[]bool)))
	case []interface{}:
		a := o.([]interface{})
		arr := make([]interface{}, 0)
		for _, av := range a {
			arr = append(arr, toUserObject(av, isavro, false, ""))
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of User
func (o *User) String() string {
	return fmt.Sprintf("customer.User<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *User) GetTopicName() datamodel.TopicNameType {
	return UserTopic
}

// GetModelName returns the name of the model
func (o *User) GetModelName() datamodel.ModelNameType {
	return UserModelName
}

func (o *User) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *User) GetID() string {
	if o.ID == "" {
		// set the id from the spec provided in the model
		o.ID = hash.Values(o.CustomerID, o.Email)
	}
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
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *User) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *User) IsMaterialized() bool {
	return true
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *User) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	idletime, err := time.ParseDuration("1s")
	if err != nil {
		panic(err)
	}
	return &datamodel.ModelMaterializeConfig{
		KeyName:   "id",
		TableName: "customer_user",
		BatchSize: 5000,
		IdleTime:  idletime,
	}
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
	duration, err := time.ParseDuration("168h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 168h0m0s. " + err.Error())
	}
	return &datamodel.ModelTopicConfig{
		Key:               "id",
		Timestamp:         "",
		NumPartitions:     8,
		ReplicationFactor: 3,
		Retention:         duration,
		MaxSize:           5242880,
	}
}

// GetCustomerID will return the customer_id
func (o *User) GetCustomerID() string {
	return o.CustomerID
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
	// make sure that these have values if empty
	o.setDefaults()
	o.FromMap(kv)
	return nil
}

var cachedCodecUser *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *User) GetAvroCodec() *goavro.Codec {
	if cachedCodecUser == nil {
		c, err := GetUserAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecUser = c
	}
	return cachedCodecUser
}

// ToAvroBinary returns the data as Avro binary data
func (o *User) ToAvroBinary() ([]byte, *goavro.Codec, error) {
	kv := o.ToMap(true)
	jbuf, _ := json.Marshal(kv)
	codec := o.GetAvroCodec()
	native, _, err := codec.NativeFromTextual(jbuf)
	if err != nil {
		return nil, nil, err
	}
	// Convert native Go form to binary Avro data
	buf, err := codec.BinaryFromNative(nil, native)
	return buf, codec, err
}

// Stringify returns the object in JSON format as a string
func (o *User) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two User objects are equal
func (o *User) IsEqual(other *User) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *User) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	return map[string]interface{}{
		"id":             o.GetID(),
		"ref_id":         o.GetRefID(),
		"ref_type":       o.RefType,
		"customer_id":    o.CustomerID,
		"hashcode":       o.Hash(),
		"created_ts":     o.CreatedAt,
		"updated_ts":     o.UpdatedAt,
		"name":           toUserObject(o.Name, isavro, false, "string"),
		"email":          toUserObject(o.Email, isavro, false, "string"),
		"title":          toUserObject(o.Title, isavro, true, "string"),
		"location":       toUserObject(o.Location, isavro, true, "string"),
		"avatar_url":     toUserObject(o.AvatarURL, isavro, true, "string"),
		"manager_id":     toUserObject(o.ManagerID, isavro, true, "string"),
		"owner":          toUserObject(o.Owner, isavro, false, "boolean"),
		"active":         toUserObject(o.Active, isavro, false, "boolean"),
		"trackable":      toUserObject(o.Trackable, isavro, false, "boolean"),
		"deleted_ts":     toUserObject(o.DeletedAt, isavro, true, "long"),
		"hired_ts":       toUserObject(o.HiredAt, isavro, true, "long"),
		"terminated_ts":  toUserObject(o.TerminatedAt, isavro, true, "long"),
		"cost_center_id": toUserObject(o.CostCenterID, isavro, true, "string"),
		"team_id":        toUserObject(o.TeamID, isavro, true, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *User) FromMap(kv map[string]interface{}) {
	// make sure that these have values if empty
	o.setDefaults()
	if val, ok := kv["id"].(string); ok {
		o.ID = val
	} else if val, ok := kv["_id"].(string); ok {
		o.ID = val
	}
	if val, ok := kv["ref_id"].(string); ok {
		o.RefID = val
	}
	if val, ok := kv["ref_type"].(string); ok {
		o.RefType = val
	}
	if val, ok := kv["customer_id"].(string); ok {
		o.CustomerID = val
	}
	if val, ok := kv["created_ts"].(int64); ok {
		o.CreatedAt = val
	} else if val, ok := kv["created_ts"].(time.Time); ok {
		o.CreatedAt = datetime.TimeToEpoch(val)
	}
	if val, ok := kv["updated_ts"].(int64); ok {
		o.UpdatedAt = val
	} else if val, ok := kv["updated_ts"].(time.Time); ok {
		o.UpdatedAt = datetime.TimeToEpoch(val)
	}
	if val, ok := kv["name"].(string); ok {
		o.Name = val
	} else {
		val := kv["name"]
		if val == nil {
			o.Name = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Name = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["email"].(string); ok {
		o.Email = val
	} else {
		val := kv["email"]
		if val == nil {
			o.Email = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Email = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["title"].(*string); ok {
		o.Title = val
	} else if val, ok := kv["title"].(string); ok {
		o.Title = &val
	} else {
		val := kv["title"]
		if val == nil {
			o.Title = pstrings.Pointer("")
		} else {
			// if coming in as avro union, convert it back
			if kv, ok := val.(map[string]interface{}); ok {
				val = kv["string"]
			}
			o.Title = pstrings.Pointer(fmt.Sprintf("%v", val))
		}
	}
	if val, ok := kv["location"].(*string); ok {
		o.Location = val
	} else if val, ok := kv["location"].(string); ok {
		o.Location = &val
	} else {
		val := kv["location"]
		if val == nil {
			o.Location = pstrings.Pointer("")
		} else {
			// if coming in as avro union, convert it back
			if kv, ok := val.(map[string]interface{}); ok {
				val = kv["string"]
			}
			o.Location = pstrings.Pointer(fmt.Sprintf("%v", val))
		}
	}
	if val, ok := kv["avatar_url"].(*string); ok {
		o.AvatarURL = val
	} else if val, ok := kv["avatar_url"].(string); ok {
		o.AvatarURL = &val
	} else {
		val := kv["avatar_url"]
		if val == nil {
			o.AvatarURL = pstrings.Pointer("")
		} else {
			// if coming in as avro union, convert it back
			if kv, ok := val.(map[string]interface{}); ok {
				val = kv["string"]
			}
			o.AvatarURL = pstrings.Pointer(fmt.Sprintf("%v", val))
		}
	}
	if val, ok := kv["manager_id"].(*string); ok {
		o.ManagerID = val
	} else if val, ok := kv["manager_id"].(string); ok {
		o.ManagerID = &val
	} else {
		val := kv["manager_id"]
		if val == nil {
			o.ManagerID = pstrings.Pointer("")
		} else {
			// if coming in as avro union, convert it back
			if kv, ok := val.(map[string]interface{}); ok {
				val = kv["string"]
			}
			o.ManagerID = pstrings.Pointer(fmt.Sprintf("%v", val))
		}
	}
	if val, ok := kv["owner"].(bool); ok {
		o.Owner = val
	} else {
		val := kv["owner"]
		if val == nil {
			o.Owner = number.ToBoolAny(nil)
		} else {
			o.Owner = number.ToBoolAny(val)
		}
	}
	if val, ok := kv["active"].(bool); ok {
		o.Active = val
	} else {
		val := kv["active"]
		if val == nil {
			o.Active = number.ToBoolAny(nil)
		} else {
			o.Active = number.ToBoolAny(val)
		}
	}
	if val, ok := kv["trackable"].(bool); ok {
		o.Trackable = val
	} else {
		val := kv["trackable"]
		if val == nil {
			o.Trackable = number.ToBoolAny(nil)
		} else {
			o.Trackable = number.ToBoolAny(val)
		}
	}
	if val, ok := kv["deleted_ts"].(*int64); ok {
		o.DeletedAt = val
	} else if val, ok := kv["deleted_ts"].(int64); ok {
		o.DeletedAt = &val
	} else {
		val := kv["deleted_ts"]
		if val == nil {
			o.DeletedAt = number.Int64Pointer(number.ToInt64Any(nil))
		} else {
			// if coming in as avro union, convert it back
			if kv, ok := val.(map[string]interface{}); ok {
				val = kv["int64"]
			}
			o.DeletedAt = number.Int64Pointer(number.ToInt64Any(val))
		}
	}
	if val, ok := kv["hired_ts"].(*int64); ok {
		o.HiredAt = val
	} else if val, ok := kv["hired_ts"].(int64); ok {
		o.HiredAt = &val
	} else {
		val := kv["hired_ts"]
		if val == nil {
			o.HiredAt = number.Int64Pointer(number.ToInt64Any(nil))
		} else {
			// if coming in as avro union, convert it back
			if kv, ok := val.(map[string]interface{}); ok {
				val = kv["int64"]
			}
			o.HiredAt = number.Int64Pointer(number.ToInt64Any(val))
		}
	}
	if val, ok := kv["terminated_ts"].(*int64); ok {
		o.TerminatedAt = val
	} else if val, ok := kv["terminated_ts"].(int64); ok {
		o.TerminatedAt = &val
	} else {
		val := kv["terminated_ts"]
		if val == nil {
			o.TerminatedAt = number.Int64Pointer(number.ToInt64Any(nil))
		} else {
			// if coming in as avro union, convert it back
			if kv, ok := val.(map[string]interface{}); ok {
				val = kv["int64"]
			}
			o.TerminatedAt = number.Int64Pointer(number.ToInt64Any(val))
		}
	}
	if val, ok := kv["cost_center_id"].(*string); ok {
		o.CostCenterID = val
	} else if val, ok := kv["cost_center_id"].(string); ok {
		o.CostCenterID = &val
	} else {
		val := kv["cost_center_id"]
		if val == nil {
			o.CostCenterID = pstrings.Pointer("")
		} else {
			// if coming in as avro union, convert it back
			if kv, ok := val.(map[string]interface{}); ok {
				val = kv["string"]
			}
			o.CostCenterID = pstrings.Pointer(fmt.Sprintf("%v", val))
		}
	}
	if val, ok := kv["team_id"].(*string); ok {
		o.TeamID = val
	} else if val, ok := kv["team_id"].(string); ok {
		o.TeamID = &val
	} else {
		val := kv["team_id"]
		if val == nil {
			o.TeamID = pstrings.Pointer("")
		} else {
			// if coming in as avro union, convert it back
			if kv, ok := val.(map[string]interface{}); ok {
				val = kv["string"]
			}
			o.TeamID = pstrings.Pointer(fmt.Sprintf("%v", val))
		}
	}
}

// Hash will return a hashcode for the object
func (o *User) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.GetID())
	args = append(args, o.GetRefID())
	args = append(args, o.RefType)
	args = append(args, o.CustomerID)
	args = append(args, o.Name)
	args = append(args, o.Email)
	args = append(args, o.Title)
	args = append(args, o.Location)
	args = append(args, o.AvatarURL)
	args = append(args, o.ManagerID)
	args = append(args, o.Owner)
	args = append(args, o.Active)
	args = append(args, o.Trackable)
	args = append(args, o.DeletedAt)
	args = append(args, o.HiredAt)
	args = append(args, o.TerminatedAt)
	args = append(args, o.CostCenterID)
	args = append(args, o.TeamID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// CreateUser creates a new User in the database
func CreateUser(ctx context.Context, db datamodel.Storage, o *User) error {
	o.setDefaults()
	return db.Create(ctx, o)
}

// DeleteUser deletes a User in the database
func DeleteUser(ctx context.Context, db datamodel.Storage, o *User) error {
	o.setDefaults()
	return db.Delete(ctx, o)
}

// UpdateUser updates a User in the database
func UpdateUser(ctx context.Context, db datamodel.Storage, o *User) error {
	o.setDefaults()
	return db.Update(ctx, o)
}

// FindUser returns a User from the database
func FindUser(ctx context.Context, db datamodel.Storage, id string) (*User, error) {
	kv, err := db.FindOne(ctx, UserModelName, id)
	if err != nil {
		return nil, err
	}
	if kv == nil {
		return nil, nil
	}
	return kv.(*User), nil
}

// FindUsers returns all User from the database matching keys
func FindUsers(ctx context.Context, db datamodel.Storage, kv map[string]interface{}) ([]*User, error) {
	res, err := db.Find(ctx, UserModelName, kv)
	if err != nil {
		return nil, err
	}
	if res != nil {
		arr := make([]*User, 0)
		for _, m := range res {
			arr = append(arr, m.(*User))
		}
		return arr, nil
	}
	return nil, nil
}

// GetUserAvroSchemaSpec creates the avro schema specification for User
func GetUserAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "customer",
		"name":         "User",
		"connect.name": "customer.User",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "ref_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "ref_type",
				"type": "string",
			},
			map[string]interface{}{
				"name": "customer_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "hashcode",
				"type": "string",
			},
			map[string]interface{}{
				"name": "created_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "updated_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "name",
				"type": "string",
			},
			map[string]interface{}{
				"name": "email",
				"type": "string",
			},
			map[string]interface{}{
				"name":    "title",
				"type":    []interface{}{"null", "string"},
				"default": nil,
			},
			map[string]interface{}{
				"name":    "location",
				"type":    []interface{}{"null", "string"},
				"default": nil,
			},
			map[string]interface{}{
				"name":    "avatar_url",
				"type":    []interface{}{"null", "string"},
				"default": nil,
			},
			map[string]interface{}{
				"name":    "manager_id",
				"type":    []interface{}{"null", "string"},
				"default": nil,
			},
			map[string]interface{}{
				"name": "owner",
				"type": "boolean",
			},
			map[string]interface{}{
				"name": "active",
				"type": "boolean",
			},
			map[string]interface{}{
				"name": "trackable",
				"type": "boolean",
			},
			map[string]interface{}{
				"name":    "deleted_ts",
				"type":    []interface{}{"null", "long"},
				"default": nil,
			},
			map[string]interface{}{
				"name":    "hired_ts",
				"type":    []interface{}{"null", "long"},
				"default": nil,
			},
			map[string]interface{}{
				"name":    "terminated_ts",
				"type":    []interface{}{"null", "long"},
				"default": nil,
			},
			map[string]interface{}{
				"name":    "cost_center_id",
				"type":    []interface{}{"null", "string"},
				"default": nil,
			},
			map[string]interface{}{
				"name":    "team_id",
				"type":    []interface{}{"null", "string"},
				"default": nil,
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetUserAvroSchema creates the avro schema for User
func GetUserAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetUserAvroSchemaSpec())
}

// TransformUserFunc is a function for transforming User during processing
type TransformUserFunc func(input *User) (*User, error)

// NewUserPipe creates a pipe for processing User items
func NewUserPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformUserFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewUserInputStream(input, errors)
	var stream chan User
	if len(transforms) > 0 {
		stream = make(chan User, 1000)
	} else {
		stream = inch
	}
	outdone := NewUserOutputStream(output, stream, errors)
	go func() {
		if len(transforms) > 0 {
			var stop bool
			for item := range inch {
				input := &item
				for _, transform := range transforms {
					out, err := transform(input)
					if err != nil {
						stop = true
						errors <- err
						break
					}
					if out == nil {
						input = nil
						break
					} else {
						input = out
					}
				}
				if stop {
					break
				}
				if input != nil {
					stream <- *input
				}
			}
			close(stream)
		}
		<-indone
		<-outdone
		done <- true
	}()
	return done
}

// NewUserInputStreamDir creates a channel for reading User as JSON newlines from a directory of files
func NewUserInputStreamDir(dir string, errors chan<- error, transforms ...TransformUserFunc) (chan User, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/customer/user\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan User)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for user")
		ch := make(chan User)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewUserInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan User)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewUserInputStreamFile creates an channel for reading User as JSON newlines from filename
func NewUserInputStreamFile(filename string, errors chan<- error, transforms ...TransformUserFunc) (chan User, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan User)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	var f io.ReadCloser = of
	if filepath.Ext(filename) == ".gz" {
		gz, err := gzip.NewReader(f)
		if err != nil {
			of.Close()
			errors <- err
			ch := make(chan User)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewUserInputStream(f, errors, transforms...)
}

// NewUserInputStream creates an channel for reading User as JSON newlines from stream
func NewUserInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformUserFunc) (chan User, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan User, 1000)
	go func() {
		defer func() { stream.Close(); close(ch); done <- true }()
		r := bufio.NewReader(stream)
		for {
			buf, err := r.ReadBytes('\n')
			if err != nil {
				if err == io.EOF {
					return
				}
				errors <- err
				return
			}
			var item User
			if err := json.Unmarshal(buf, &item); err != nil {
				errors <- err
				return
			}
			in := &item
			var skip bool
			for _, transform := range transforms {
				in, err = transform(in)
				if err != nil {
					errors <- err
					return
				}
				if in == nil {
					skip = true
					break
				}
			}
			if !skip {
				ch <- *in
			}
		}
	}()
	return ch, done
}

// NewUserOutputStreamDir will output json newlines from channel and save in dir
func NewUserOutputStreamDir(dir string, ch chan User, errors chan<- error, transforms ...TransformUserFunc) <-chan bool {
	fp := filepath.Join(dir, "/customer/user\\.json(\\.gz)?$")
	os.MkdirAll(filepath.Dir(fp), 0777)
	of, err := os.Create(fp)
	if err != nil {
		errors <- err
		done := make(chan bool, 1)
		done <- true
		return done
	}
	gz, err := gzip.NewWriterLevel(of, gzip.BestCompression)
	if err != nil {
		errors <- err
		done := make(chan bool, 1)
		done <- true
		return done
	}
	return NewUserOutputStream(gz, ch, errors, transforms...)
}

// NewUserOutputStream will output json newlines from channel to the stream
func NewUserOutputStream(stream io.WriteCloser, ch chan User, errors chan<- error, transforms ...TransformUserFunc) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() {
			if gz, ok := stream.(*gzip.Writer); ok {
				gz.Flush()
				gz.Close()
			}
			stream.Close()
			done <- true
		}()
		for item := range ch {
			in := &item
			var skip bool
			var err error
			for _, transform := range transforms {
				in, err = transform(in)
				if err != nil {
					errors <- err
					return
				}
				if in == nil {
					skip = true
					break
				}
			}
			if !skip {
				buf, err := json.Marshal(in)
				if err != nil {
					errors <- err
					return
				}
				stream.Write(buf)
				stream.Write([]byte{'\n'})
			}
		}
	}()
	return done
}

// UserSendEvent is an event detail for sending data
type UserSendEvent struct {
	User    *User
	headers map[string]string
	time    time.Time
	key     string
}

var _ datamodel.ModelSendEvent = (*UserSendEvent)(nil)

// Key is the key to use for the message
func (e *UserSendEvent) Key() string {
	if e.key == "" {
		return e.User.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *UserSendEvent) Object() datamodel.Model {
	return e.User
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *UserSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *UserSendEvent) Timestamp() time.Time {
	return e.time
}

// UserSendEventOpts is a function handler for setting opts
type UserSendEventOpts func(o *UserSendEvent)

// WithUserSendEventKey sets the key value to a value different than the object ID
func WithUserSendEventKey(key string) UserSendEventOpts {
	return func(o *UserSendEvent) {
		o.key = key
	}
}

// WithUserSendEventTimestamp sets the timestamp value
func WithUserSendEventTimestamp(tv time.Time) UserSendEventOpts {
	return func(o *UserSendEvent) {
		o.time = tv
	}
}

// WithUserSendEventHeader sets the timestamp value
func WithUserSendEventHeader(key, value string) UserSendEventOpts {
	return func(o *UserSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewUserSendEvent returns a new UserSendEvent instance
func NewUserSendEvent(o *User, opts ...UserSendEventOpts) *UserSendEvent {
	res := &UserSendEvent{
		User: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewUserProducer will stream data from the channel
func NewUserProducer(producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		ctx := context.Background()
		for item := range ch {
			if object, ok := item.Object().(*User); ok {
				binary, codec, err := object.ToAvroBinary()
				if err != nil {
					errors <- fmt.Errorf("error encoding %s to avro binary data. %v", object.String(), err)
					return
				}
				headers := map[string]string{}
				object.SetEventHeaders(headers)
				for k, v := range item.Headers() {
					headers[k] = v
				}
				tv := item.Timestamp()
				if tv.IsZero() {
					tv = object.GetTimestamp() // if not provided in the message, use the objects value
				}
				if tv.IsZero() {
					tv = time.Now() // if its still zero, use the ingest time
				}
				msg := eventing.Message{
					Key:       item.Key(),
					Value:     binary,
					Codec:     codec,
					Headers:   headers,
					Timestamp: tv,
					Topic:     object.GetTopicName().String(),
				}
				if err := producer.Send(ctx, msg); err != nil {
					errors <- fmt.Errorf("error sending %s. %v", object.String(), err)
				}
			} else {
				errors <- fmt.Errorf("invalid event received. expected an object of type customer.User but received on of type %v", reflect.TypeOf(item.Object()))
			}
		}
	}()
	return done
}

// NewUserConsumer will stream data from the topic into the provided channel
func NewUserConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) {
	consumer.Consume(&eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object User
			if err := json.Unmarshal(msg.Value, &object); err != nil {
				return fmt.Errorf("error unmarshaling json data into customer.User: %s", err)
			}
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &UserReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object User
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &UserReceiveEvent{nil, msg, true}
		},
	})
}

// UserReceiveEvent is an event detail for receiving data
type UserReceiveEvent struct {
	User    *User
	message eventing.Message
	eof     bool
}

var _ datamodel.ModelReceiveEvent = (*UserReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *UserReceiveEvent) Object() datamodel.Model {
	return e.User
}

// Message returns the underlying message data for the event
func (e *UserReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *UserReceiveEvent) EOF() bool {
	return e.eof
}

// UserProducer implements the datamodel.ModelEventProducer
type UserProducer struct {
	ch   chan datamodel.ModelSendEvent
	done <-chan bool
}

var _ datamodel.ModelEventProducer = (*UserProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *UserProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *UserProducer) Close() error {
	close(p.ch)
	<-p.done
	return nil
}

// NewProducerChannel returns a channel which can be used for producing Model events
func (o *User) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent)
	return &UserProducer{
		ch:   ch,
		done: NewUserProducer(producer, ch, errors),
	}
}

// NewUserProducerChannel returns a channel which can be used for producing Model events
func NewUserProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent)
	return &UserProducer{
		ch:   ch,
		done: NewUserProducer(producer, ch, errors),
	}
}

// UserConsumer implements the datamodel.ModelEventConsumer
type UserConsumer struct {
	ch chan datamodel.ModelReceiveEvent
}

var _ datamodel.ModelEventConsumer = (*UserConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *UserConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *UserConsumer) Close() error {
	close(c.ch)
	return nil
}

// NewConsumerChannel returns a consumer channel which can be used to consume Model events
func (o *User) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	NewUserConsumer(consumer, ch, errors)
	return &UserConsumer{
		ch: ch,
	}
}

// NewUserConsumerChannel returns a consumer channel which can be used to consume Model events
func NewUserConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	NewUserConsumer(consumer, ch, errors)
	return &UserConsumer{
		ch: ch,
	}
}
