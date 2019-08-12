// DO NOT EDIT -- generated code

// Package customer - customer specific data
package customer

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"sync"
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
	"github.com/pinpt/go-common/slice"
	pstrings "github.com/pinpt/go-common/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	// UserActiveColumn is the active column name
	UserActiveColumn = "active"
	// UserAvatarURLColumn is the avatar_url column name
	UserAvatarURLColumn = "avatar_url"
	// UserCostCenterIDColumn is the cost_center_id column name
	UserCostCenterIDColumn = "cost_center_id"
	// UserCreatedAtColumn is the created_ts column name
	UserCreatedAtColumn = "created_ts"
	// UserCustomerIDColumn is the customer_id column name
	UserCustomerIDColumn = "customer_id"
	// UserDeletedDateColumn is the deleted_date column name
	UserDeletedDateColumn = "deleted_date"
	// UserDeletedDateColumnEpochColumn is the epoch column property of the DeletedDate name
	UserDeletedDateColumnEpochColumn = "deleted_date->epoch"
	// UserDeletedDateColumnOffsetColumn is the offset column property of the DeletedDate name
	UserDeletedDateColumnOffsetColumn = "deleted_date->offset"
	// UserDeletedDateColumnRfc3339Column is the rfc3339 column property of the DeletedDate name
	UserDeletedDateColumnRfc3339Column = "deleted_date->rfc3339"
	// UserEmailColumn is the email column name
	UserEmailColumn = "email"
	// UserHiredDateColumn is the hired_date column name
	UserHiredDateColumn = "hired_date"
	// UserHiredDateColumnEpochColumn is the epoch column property of the HiredDate name
	UserHiredDateColumnEpochColumn = "hired_date->epoch"
	// UserHiredDateColumnOffsetColumn is the offset column property of the HiredDate name
	UserHiredDateColumnOffsetColumn = "hired_date->offset"
	// UserHiredDateColumnRfc3339Column is the rfc3339 column property of the HiredDate name
	UserHiredDateColumnRfc3339Column = "hired_date->rfc3339"
	// UserIDColumn is the id column name
	UserIDColumn = "id"
	// UserLocationColumn is the location column name
	UserLocationColumn = "location"
	// UserManagerIDColumn is the manager_id column name
	UserManagerIDColumn = "manager_id"
	// UserNameColumn is the name column name
	UserNameColumn = "name"
	// UserOwnerColumn is the owner column name
	UserOwnerColumn = "owner"
	// UserPrimaryTeamIDColumn is the primary_team_id column name
	UserPrimaryTeamIDColumn = "primary_team_id"
	// UserRefIDColumn is the ref_id column name
	UserRefIDColumn = "ref_id"
	// UserRefTypeColumn is the ref_type column name
	UserRefTypeColumn = "ref_type"
	// UserRoleIdsColumn is the role_ids column name
	UserRoleIdsColumn = "role_ids"
	// UserTeamIdsColumn is the team_ids column name
	UserTeamIdsColumn = "team_ids"
	// UserTerminatedDateColumn is the terminated_date column name
	UserTerminatedDateColumn = "terminated_date"
	// UserTerminatedDateColumnEpochColumn is the epoch column property of the TerminatedDate name
	UserTerminatedDateColumnEpochColumn = "terminated_date->epoch"
	// UserTerminatedDateColumnOffsetColumn is the offset column property of the TerminatedDate name
	UserTerminatedDateColumnOffsetColumn = "terminated_date->offset"
	// UserTerminatedDateColumnRfc3339Column is the rfc3339 column property of the TerminatedDate name
	UserTerminatedDateColumnRfc3339Column = "terminated_date->rfc3339"
	// UserTitleColumn is the title column name
	UserTitleColumn = "title"
	// UserTrackableColumn is the trackable column name
	UserTrackableColumn = "trackable"
	// UserUpdatedAtColumn is the updated_ts column name
	UserUpdatedAtColumn = "updated_ts"
)

// UserDeletedDate represents the object structure for deleted_date
type UserDeletedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toUserDeletedDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toUserDeletedDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *UserDeletedDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *UserDeletedDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toUserDeletedDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toUserDeletedDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toUserDeletedDateObject(o.Rfc3339, isavro, false, "string"),
	}
}

func (o *UserDeletedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *UserDeletedDate) FromMap(kv map[string]interface{}) {

	if val, ok := kv["epoch"].(int64); ok {
		o.Epoch = val
	} else {
		if val, ok := kv["epoch"]; ok {
			if val == nil {
				o.Epoch = number.ToInt64Any(nil)
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
				o.Offset = number.ToInt64Any(nil)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// UserHiredDate represents the object structure for hired_date
type UserHiredDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toUserHiredDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toUserHiredDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *UserHiredDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *UserHiredDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toUserHiredDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toUserHiredDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toUserHiredDateObject(o.Rfc3339, isavro, false, "string"),
	}
}

func (o *UserHiredDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *UserHiredDate) FromMap(kv map[string]interface{}) {

	if val, ok := kv["epoch"].(int64); ok {
		o.Epoch = val
	} else {
		if val, ok := kv["epoch"]; ok {
			if val == nil {
				o.Epoch = number.ToInt64Any(nil)
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
				o.Offset = number.ToInt64Any(nil)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// UserTerminatedDate represents the object structure for terminated_date
type UserTerminatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toUserTerminatedDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toUserTerminatedDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *UserTerminatedDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *UserTerminatedDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toUserTerminatedDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toUserTerminatedDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toUserTerminatedDateObject(o.Rfc3339, isavro, false, "string"),
	}
}

func (o *UserTerminatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *UserTerminatedDate) FromMap(kv map[string]interface{}) {

	if val, ok := kv["epoch"].(int64); ok {
		o.Epoch = val
	} else {
		if val, ok := kv["epoch"]; ok {
			if val == nil {
				o.Epoch = number.ToInt64Any(nil)
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
				o.Offset = number.ToInt64Any(nil)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// User the customer's user record
type User struct {
	// Active if true, the user is active and able to login
	Active bool `json:"active" bson:"active" yaml:"active" faker:"-"`
	// AvatarURL the user avatar url
	AvatarURL *string `json:"avatar_url" bson:"avatar_url" yaml:"avatar_url" faker:"avatar"`
	// CostCenterID the id of the cost center
	CostCenterID *string `json:"cost_center_id" bson:"cost_center_id" yaml:"cost_center_id" faker:"-"`
	// CreatedAt the date the record was created in Epoch time
	CreatedAt int64 `json:"created_ts" bson:"created_ts" yaml:"created_ts" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// DeletedDate when the user record was deleted in epoch timestamp
	DeletedDate UserDeletedDate `json:"deleted_date" bson:"deleted_date" yaml:"deleted_date" faker:"-"`
	// Email the email of the user
	Email string `json:"email" bson:"email" yaml:"email" faker:"email"`
	// HiredDate when the user was hired in epoch timestamp
	HiredDate UserHiredDate `json:"hired_date" bson:"hired_date" yaml:"hired_date" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// Location the location of the user
	Location *string `json:"location" bson:"location" yaml:"location" faker:"location"`
	// ManagerID the manager user id
	ManagerID *string `json:"manager_id" bson:"manager_id" yaml:"manager_id" faker:"-"`
	// Name name of the user
	Name string `json:"name" bson:"name" yaml:"name" faker:"person"`
	// Owner if true, the user is an owner of the account
	Owner bool `json:"owner" bson:"owner" yaml:"owner" faker:"-"`
	// PrimaryTeamID the team id of the user's primary team
	PrimaryTeamID *string `json:"primary_team_id" bson:"primary_team_id" yaml:"primary_team_id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RoleIds the auth role ids for the user
	RoleIds []string `json:"role_ids" bson:"role_ids" yaml:"role_ids" faker:"-"`
	// TeamIds the team ids that the user is part of
	TeamIds []string `json:"team_ids" bson:"team_ids" yaml:"team_ids" faker:"-"`
	// TerminatedDate when the user was terminated in epoch timestamp
	TerminatedDate UserTerminatedDate `json:"terminated_date" bson:"terminated_date" yaml:"terminated_date" faker:"-"`
	// Title the title of the user
	Title *string `json:"title" bson:"title" yaml:"title" faker:"jobtitle"`
	// Trackable if true, the user is trackable in the pinpoint system
	Trackable bool `json:"trackable" bson:"trackable" yaml:"trackable" faker:"-"`
	// UpdatedAt the date the record was updated in Epoch time
	UpdatedAt int64 `json:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
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
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *User:
		return v.ToMap(isavro)

	case UserDeletedDate:
		return v.ToMap(isavro)

	case UserHiredDate:
		return v.ToMap(isavro)

	case UserTerminatedDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
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

func (o *User) setDefaults(frommap bool) {
	if o.AvatarURL == nil {
		o.AvatarURL = &emptyString
	}
	if o.CostCenterID == nil {
		o.CostCenterID = &emptyString
	}
	if o.Location == nil {
		o.Location = &emptyString
	}
	if o.ManagerID == nil {
		o.ManagerID = &emptyString
	}
	if o.PrimaryTeamID == nil {
		o.PrimaryTeamID = &emptyString
	}
	if o.RoleIds == nil {
		o.RoleIds = make([]string, 0)
	}
	if o.TeamIds == nil {
		o.TeamIds = make([]string, 0)
	}
	if o.Title == nil {
		o.Title = &emptyString
	}

	if o.ID == "" {
		// set the id from the spec provided in the model
		o.ID = hash.Values(o.CustomerID, o.Email)
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
	retention, err := time.ParseDuration("168h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 168h0m0s. " + err.Error())
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
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetStateKey returns a key for use in state store
func (o *User) GetStateKey() string {
	key := "id"
	return fmt.Sprintf("%s_%s", key, o.GetID())
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
	o.FromMap(kv)
	if idstr, ok := kv["id"].(string); ok {
		o.ID = idstr
	}
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

// FromAvroBinary will convert from Avro binary data into data in this object
func (o *User) FromAvroBinary(value []byte) error {
	var nullHeader = []byte{byte(0)}
	// if this still has the schema encoded in the header, move past it to the avro payload
	if bytes.HasPrefix(value, nullHeader) {
		value = value[5:]
	}
	kv, _, err := o.GetAvroCodec().NativeFromBinary(value)
	if err != nil {
		return err
	}
	o.FromMap(kv.(map[string]interface{}))
	return nil
}

// Stringify returns the object in JSON format as a string
func (o *User) Stringify() string {
	o.Hash()
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
		if o.RoleIds == nil {
			o.RoleIds = make([]string, 0)
		}
		if o.TeamIds == nil {
			o.TeamIds = make([]string, 0)
		}
	}
	o.setDefaults(false)
	return map[string]interface{}{
		"active":          toUserObject(o.Active, isavro, false, "boolean"),
		"avatar_url":      toUserObject(o.AvatarURL, isavro, true, "string"),
		"cost_center_id":  toUserObject(o.CostCenterID, isavro, true, "string"),
		"created_ts":      toUserObject(o.CreatedAt, isavro, false, "long"),
		"customer_id":     toUserObject(o.CustomerID, isavro, false, "string"),
		"deleted_date":    toUserObject(o.DeletedDate, isavro, false, "deleted_date"),
		"email":           toUserObject(o.Email, isavro, false, "string"),
		"hired_date":      toUserObject(o.HiredDate, isavro, false, "hired_date"),
		"id":              toUserObject(o.ID, isavro, false, "string"),
		"location":        toUserObject(o.Location, isavro, true, "string"),
		"manager_id":      toUserObject(o.ManagerID, isavro, true, "string"),
		"name":            toUserObject(o.Name, isavro, false, "string"),
		"owner":           toUserObject(o.Owner, isavro, false, "boolean"),
		"primary_team_id": toUserObject(o.PrimaryTeamID, isavro, true, "string"),
		"ref_id":          toUserObject(o.RefID, isavro, false, "string"),
		"ref_type":        toUserObject(o.RefType, isavro, false, "string"),
		"role_ids":        toUserObject(o.RoleIds, isavro, false, "role_ids"),
		"team_ids":        toUserObject(o.TeamIds, isavro, false, "team_ids"),
		"terminated_date": toUserObject(o.TerminatedDate, isavro, false, "terminated_date"),
		"title":           toUserObject(o.Title, isavro, true, "string"),
		"trackable":       toUserObject(o.Trackable, isavro, false, "boolean"),
		"updated_ts":      toUserObject(o.UpdatedAt, isavro, false, "long"),
		"hashcode":        toUserObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *User) FromMap(kv map[string]interface{}) {

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
				o.Active = number.ToBoolAny(nil)
			} else {
				o.Active = number.ToBoolAny(val)
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
				// if coming in as avro union, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.AvatarURL = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if val, ok := kv["cost_center_id"].(*string); ok {
		o.CostCenterID = val
	} else if val, ok := kv["cost_center_id"].(string); ok {
		o.CostCenterID = &val
	} else {
		if val, ok := kv["cost_center_id"]; ok {
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
	}

	if val, ok := kv["created_ts"].(int64); ok {
		o.CreatedAt = val
	} else {
		if val, ok := kv["created_ts"]; ok {
			if val == nil {
				o.CreatedAt = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.CreatedAt = number.ToInt64Any(val)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.CustomerID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["deleted_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.DeletedDate.FromMap(kv)
		} else if sv, ok := val.(UserDeletedDate); ok {
			// struct
			o.DeletedDate = sv
		} else if sp, ok := val.(*UserDeletedDate); ok {
			// struct pointer
			o.DeletedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.DeletedDate.Epoch = dt.Epoch
			o.DeletedDate.Rfc3339 = dt.Rfc3339
			o.DeletedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.DeletedDate.Epoch = dt.Epoch
			o.DeletedDate.Rfc3339 = dt.Rfc3339
			o.DeletedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.DeletedDate.Epoch = dt.Epoch
				o.DeletedDate.Rfc3339 = dt.Rfc3339
				o.DeletedDate.Offset = dt.Offset
			}
		}
	} else {
		o.DeletedDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["email"].(string); ok {
		o.Email = val
	} else {
		if val, ok := kv["email"]; ok {
			if val == nil {
				o.Email = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Email = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["hired_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.HiredDate.FromMap(kv)
		} else if sv, ok := val.(UserHiredDate); ok {
			// struct
			o.HiredDate = sv
		} else if sp, ok := val.(*UserHiredDate); ok {
			// struct pointer
			o.HiredDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.HiredDate.Epoch = dt.Epoch
			o.HiredDate.Rfc3339 = dt.Rfc3339
			o.HiredDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.HiredDate.Epoch = dt.Epoch
			o.HiredDate.Rfc3339 = dt.Rfc3339
			o.HiredDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.HiredDate.Epoch = dt.Epoch
				o.HiredDate.Rfc3339 = dt.Rfc3339
				o.HiredDate.Offset = dt.Offset
			}
		}
	} else {
		o.HiredDate.FromMap(map[string]interface{}{})
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

	if val, ok := kv["location"].(*string); ok {
		o.Location = val
	} else if val, ok := kv["location"].(string); ok {
		o.Location = &val
	} else {
		if val, ok := kv["location"]; ok {
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
	}

	if val, ok := kv["manager_id"].(*string); ok {
		o.ManagerID = val
	} else if val, ok := kv["manager_id"].(string); ok {
		o.ManagerID = &val
	} else {
		if val, ok := kv["manager_id"]; ok {
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

	if val, ok := kv["owner"].(bool); ok {
		o.Owner = val
	} else {
		if val, ok := kv["owner"]; ok {
			if val == nil {
				o.Owner = number.ToBoolAny(nil)
			} else {
				o.Owner = number.ToBoolAny(val)
			}
		}
	}

	if val, ok := kv["primary_team_id"].(*string); ok {
		o.PrimaryTeamID = val
	} else if val, ok := kv["primary_team_id"].(string); ok {
		o.PrimaryTeamID = &val
	} else {
		if val, ok := kv["primary_team_id"]; ok {
			if val == nil {
				o.PrimaryTeamID = pstrings.Pointer("")
			} else {
				// if coming in as avro union, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.PrimaryTeamID = pstrings.Pointer(fmt.Sprintf("%v", val))
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

	if val, ok := kv["role_ids"]; ok {
		if val != nil {
			na := make([]string, 0)
			if a, ok := val.([]string); ok {
				na = append(na, a...)
			} else {
				if a, ok := val.([]interface{}); ok {
					for _, ae := range a {
						if av, ok := ae.(string); ok {
							na = append(na, av)
						} else {
							if badMap, ok := ae.(map[interface{}]interface{}); ok {
								ae = slice.ConvertToStringToInterface(badMap)
							}
							b, _ := json.Marshal(ae)
							var av string
							if err := json.Unmarshal(b, &av); err != nil {
								panic("unsupported type for role_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else if s, ok := val.(string); ok {
					for _, sv := range strings.Split(s, ",") {
						na = append(na, strings.TrimSpace(sv))
					}
				} else if a, ok := val.(primitive.A); ok {
					for _, ae := range a {
						if av, ok := ae.(string); ok {
							na = append(na, av)
						} else {
							b, _ := json.Marshal(ae)
							var av string
							if err := json.Unmarshal(b, &av); err != nil {
								panic("unsupported type for role_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for role_ids field")
				}
			}
			o.RoleIds = na
		}
	}
	if o.RoleIds == nil {
		o.RoleIds = make([]string, 0)
	}

	if val, ok := kv["team_ids"]; ok {
		if val != nil {
			na := make([]string, 0)
			if a, ok := val.([]string); ok {
				na = append(na, a...)
			} else {
				if a, ok := val.([]interface{}); ok {
					for _, ae := range a {
						if av, ok := ae.(string); ok {
							na = append(na, av)
						} else {
							if badMap, ok := ae.(map[interface{}]interface{}); ok {
								ae = slice.ConvertToStringToInterface(badMap)
							}
							b, _ := json.Marshal(ae)
							var av string
							if err := json.Unmarshal(b, &av); err != nil {
								panic("unsupported type for team_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else if s, ok := val.(string); ok {
					for _, sv := range strings.Split(s, ",") {
						na = append(na, strings.TrimSpace(sv))
					}
				} else if a, ok := val.(primitive.A); ok {
					for _, ae := range a {
						if av, ok := ae.(string); ok {
							na = append(na, av)
						} else {
							b, _ := json.Marshal(ae)
							var av string
							if err := json.Unmarshal(b, &av); err != nil {
								panic("unsupported type for team_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for team_ids field")
				}
			}
			o.TeamIds = na
		}
	}
	if o.TeamIds == nil {
		o.TeamIds = make([]string, 0)
	}

	if val, ok := kv["terminated_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.TerminatedDate.FromMap(kv)
		} else if sv, ok := val.(UserTerminatedDate); ok {
			// struct
			o.TerminatedDate = sv
		} else if sp, ok := val.(*UserTerminatedDate); ok {
			// struct pointer
			o.TerminatedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.TerminatedDate.Epoch = dt.Epoch
			o.TerminatedDate.Rfc3339 = dt.Rfc3339
			o.TerminatedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.TerminatedDate.Epoch = dt.Epoch
			o.TerminatedDate.Rfc3339 = dt.Rfc3339
			o.TerminatedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.TerminatedDate.Epoch = dt.Epoch
				o.TerminatedDate.Rfc3339 = dt.Rfc3339
				o.TerminatedDate.Offset = dt.Offset
			}
		}
	} else {
		o.TerminatedDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["title"].(*string); ok {
		o.Title = val
	} else if val, ok := kv["title"].(string); ok {
		o.Title = &val
	} else {
		if val, ok := kv["title"]; ok {
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
	}

	if val, ok := kv["trackable"].(bool); ok {
		o.Trackable = val
	} else {
		if val, ok := kv["trackable"]; ok {
			if val == nil {
				o.Trackable = number.ToBoolAny(nil)
			} else {
				o.Trackable = number.ToBoolAny(val)
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
func (o *User) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Active)
	args = append(args, o.AvatarURL)
	args = append(args, o.CostCenterID)
	args = append(args, o.CreatedAt)
	args = append(args, o.CustomerID)
	args = append(args, o.DeletedDate)
	args = append(args, o.Email)
	args = append(args, o.HiredDate)
	args = append(args, o.ID)
	args = append(args, o.Location)
	args = append(args, o.ManagerID)
	args = append(args, o.Name)
	args = append(args, o.Owner)
	args = append(args, o.PrimaryTeamID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RoleIds)
	args = append(args, o.TeamIds)
	args = append(args, o.TerminatedDate)
	args = append(args, o.Title)
	args = append(args, o.Trackable)
	args = append(args, o.UpdatedAt)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// CreateUser creates a new User in the database
func CreateUser(ctx context.Context, db datamodel.Storage, o *User) error {
	o.setDefaults(true)
	return db.Create(ctx, o)
}

// DeleteUser deletes a User in the database
func DeleteUser(ctx context.Context, db datamodel.Storage, o *User) error {
	o.setDefaults(true)
	return db.Delete(ctx, o)
}

// UpdateUser updates a User in the database
func UpdateUser(ctx context.Context, db datamodel.Storage, o *User) error {
	o.setDefaults(true)
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
		"type":      "record",
		"namespace": "customer",
		"name":      "User",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "hashcode",
				"type": "string",
			},
			map[string]interface{}{
				"name": "active",
				"type": "boolean",
			},
			map[string]interface{}{
				"name":    "avatar_url",
				"type":    []interface{}{"null", "string"},
				"default": nil,
			},
			map[string]interface{}{
				"name":    "cost_center_id",
				"type":    []interface{}{"null", "string"},
				"default": nil,
			},
			map[string]interface{}{
				"name": "created_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "customer_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "deleted_date",
				"type": map[string]interface{}{"doc": "when the user record was deleted in epoch timestamp", "fields": []interface{}{map[string]interface{}{"doc": "the date in epoch format", "name": "epoch", "type": "long"}, map[string]interface{}{"doc": "the timezone offset from GMT", "name": "offset", "type": "long"}, map[string]interface{}{"doc": "the date in RFC3339 format", "name": "rfc3339", "type": "string"}}, "name": "deleted_date", "type": "record"},
			},
			map[string]interface{}{
				"name": "email",
				"type": "string",
			},
			map[string]interface{}{
				"name": "hired_date",
				"type": map[string]interface{}{"doc": "when the user was hired in epoch timestamp", "fields": []interface{}{map[string]interface{}{"doc": "the date in epoch format", "name": "epoch", "type": "long"}, map[string]interface{}{"doc": "the timezone offset from GMT", "name": "offset", "type": "long"}, map[string]interface{}{"doc": "the date in RFC3339 format", "name": "rfc3339", "type": "string"}}, "name": "hired_date", "type": "record"},
			},
			map[string]interface{}{
				"name": "id",
				"type": "string",
			},
			map[string]interface{}{
				"name":    "location",
				"type":    []interface{}{"null", "string"},
				"default": nil,
			},
			map[string]interface{}{
				"name":    "manager_id",
				"type":    []interface{}{"null", "string"},
				"default": nil,
			},
			map[string]interface{}{
				"name": "name",
				"type": "string",
			},
			map[string]interface{}{
				"name": "owner",
				"type": "boolean",
			},
			map[string]interface{}{
				"name":    "primary_team_id",
				"type":    []interface{}{"null", "string"},
				"default": nil,
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
				"name": "role_ids",
				"type": map[string]interface{}{"items": "string", "name": "role_ids", "type": "array"},
			},
			map[string]interface{}{
				"name": "team_ids",
				"type": map[string]interface{}{"items": "string", "name": "team_ids", "type": "array"},
			},
			map[string]interface{}{
				"name": "terminated_date",
				"type": map[string]interface{}{"doc": "when the user was terminated in epoch timestamp", "fields": []interface{}{map[string]interface{}{"doc": "the date in epoch format", "name": "epoch", "type": "long"}, map[string]interface{}{"doc": "the timezone offset from GMT", "name": "offset", "type": "long"}, map[string]interface{}{"doc": "the date in RFC3339 format", "name": "rfc3339", "type": "string"}}, "name": "terminated_date", "type": "record"},
			},
			map[string]interface{}{
				"name":    "title",
				"type":    []interface{}{"null", "string"},
				"default": nil,
			},
			map[string]interface{}{
				"name": "trackable",
				"type": "boolean",
			},
			map[string]interface{}{
				"name": "updated_ts",
				"type": "long",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *User) GetEventAPIConfig() datamodel.EventAPIConfig {
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
func NewUserProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		for {
			select {
			case <-ctx.Done():
				return
			case item := <-ch:
				if item == nil {
					empty <- true
					return
				}
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
					// add generated message headers
					headers["message-id"] = pstrings.NewUUIDV4()
					headers["message-ts"] = fmt.Sprintf("%v", datetime.EpochNow())
					msg := eventing.Message{
						Encoding:  eventing.AvroEncoding,
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
		}
	}()
	return done
}

// NewUserConsumer will stream data from the topic into the provided channel
func NewUserConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object User
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into customer.User: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avro data into customer.User: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for customer.User")
			}

			// ignore messages that have exceeded the TTL
			cfg := object.GetTopicConfig()
			if cfg != nil && cfg.TTL != 0 && msg.Timestamp.UTC().Add(cfg.TTL).Sub(time.Now().UTC()) < 0 {
				// if disable auto and we're skipping, we need to commit the message
				if !msg.IsAutoCommit() {
					msg.Commit()
				}
				return nil
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
	}
	consumer.Consume(adapter)
	return adapter
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
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*UserProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *UserProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *UserProducer) Close() error {
	p.mu.Lock()
	closed := p.closed
	p.closed = true
	p.mu.Unlock()
	if !closed {
		close(p.ch)
		<-p.empty
		p.cancel()
		<-p.done
	}
	return nil
}

// NewProducerChannel returns a channel which can be used for producing Model events
func (o *User) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *User) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &UserProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewUserProducer(newctx, producer, ch, errors, empty),
	}
}

// NewUserProducerChannel returns a channel which can be used for producing Model events
func NewUserProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewUserProducerChannelSize(producer, 0, errors)
}

// NewUserProducerChannelSize returns a channel which can be used for producing Model events
func NewUserProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &UserProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewUserProducer(newctx, producer, ch, errors, empty),
	}
}

// UserConsumer implements the datamodel.ModelEventConsumer
type UserConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*UserConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *UserConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *UserConsumer) Close() error {
	c.mu.Lock()
	closed := c.closed
	c.closed = true
	c.mu.Unlock()
	var err error
	if !closed {
		c.callback.Close()
		err = c.consumer.Close()
	}
	return err
}

// NewConsumerChannel returns a consumer channel which can be used to consume Model events
func (o *User) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &UserConsumer{
		ch:       ch,
		callback: NewUserConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewUserConsumerChannel returns a consumer channel which can be used to consume Model events
func NewUserConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &UserConsumer{
		ch:       ch,
		callback: NewUserConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
