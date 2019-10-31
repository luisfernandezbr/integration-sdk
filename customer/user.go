// DO NOT EDIT -- generated code

// Package customer - customer specific data
package customer

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
	pstrings "github.com/pinpt/go-common/strings"
)

const (
	// UserTopic is the default topic name
	UserTopic datamodel.TopicNameType = "customer_User_topic"

	// UserModelName is the model name
	UserModelName datamodel.ModelNameType = "customer.User"
)

const (
	// UserAvatarURLColumn is the avatar_url column name
	UserAvatarURLColumn = "AvatarURL"
	// UserCostCenterIDColumn is the cost_center_id column name
	UserCostCenterIDColumn = "CostCenterID"
	// UserCustomerIDColumn is the customer_id column name
	UserCustomerIDColumn = "CustomerID"
	// UserEmailColumn is the email column name
	UserEmailColumn = "Email"
	// UserHiredDateColumn is the hired_date column name
	UserHiredDateColumn = "HiredDate"
	// UserHiredDateColumnEpochColumn is the epoch column property of the HiredDate name
	UserHiredDateColumnEpochColumn = "HiredDate.Epoch"
	// UserHiredDateColumnOffsetColumn is the offset column property of the HiredDate name
	UserHiredDateColumnOffsetColumn = "HiredDate.Offset"
	// UserHiredDateColumnRfc3339Column is the rfc3339 column property of the HiredDate name
	UserHiredDateColumnRfc3339Column = "HiredDate.Rfc3339"
	// UserIDColumn is the id column name
	UserIDColumn = "ID"
	// UserLocationColumn is the location column name
	UserLocationColumn = "Location"
	// UserManagerIDColumn is the manager_id column name
	UserManagerIDColumn = "ManagerID"
	// UserNameColumn is the name column name
	UserNameColumn = "Name"
	// UserPrimaryTeamIDColumn is the primary_team_id column name
	UserPrimaryTeamIDColumn = "PrimaryTeamID"
	// UserRefIDColumn is the ref_id column name
	UserRefIDColumn = "RefID"
	// UserRefTypeColumn is the ref_type column name
	UserRefTypeColumn = "RefType"
	// UserTerminatedDateColumn is the terminated_date column name
	UserTerminatedDateColumn = "TerminatedDate"
	// UserTerminatedDateColumnEpochColumn is the epoch column property of the TerminatedDate name
	UserTerminatedDateColumnEpochColumn = "TerminatedDate.Epoch"
	// UserTerminatedDateColumnOffsetColumn is the offset column property of the TerminatedDate name
	UserTerminatedDateColumnOffsetColumn = "TerminatedDate.Offset"
	// UserTerminatedDateColumnRfc3339Column is the rfc3339 column property of the TerminatedDate name
	UserTerminatedDateColumnRfc3339Column = "TerminatedDate.Rfc3339"
	// UserTitleColumn is the title column name
	UserTitleColumn = "Title"
	// UserUpdatedAtColumn is the updated_ts column name
	UserUpdatedAtColumn = "UpdatedAt"
)

// UserHiredDate represents the object structure for hired_date
type UserHiredDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toUserHiredDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *UserHiredDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *UserHiredDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toUserHiredDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toUserHiredDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toUserHiredDateObject(o.Rfc3339, false),
	}
}

func (o *UserHiredDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *UserHiredDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

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
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toUserTerminatedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *UserTerminatedDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *UserTerminatedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toUserTerminatedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toUserTerminatedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toUserTerminatedDateObject(o.Rfc3339, false),
	}
}

func (o *UserTerminatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *UserTerminatedDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

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
	// AvatarURL the user avatar url
	AvatarURL *string `json:"avatar_url,omitempty" codec:"avatar_url,omitempty" bson:"avatar_url" yaml:"avatar_url,omitempty" faker:"avatar"`
	// CostCenterID the id of the cost center
	CostCenterID *string `json:"cost_center_id,omitempty" codec:"cost_center_id,omitempty" bson:"cost_center_id" yaml:"cost_center_id,omitempty" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Email the email of the user
	Email string `json:"email" codec:"email" bson:"email" yaml:"email" faker:"email"`
	// HiredDate when the user was hired in epoch timestamp
	HiredDate UserHiredDate `json:"hired_date" codec:"hired_date" bson:"hired_date" yaml:"hired_date" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// Location the location of the user
	Location *string `json:"location,omitempty" codec:"location,omitempty" bson:"location" yaml:"location,omitempty" faker:"location"`
	// ManagerID the manager user id
	ManagerID *string `json:"manager_id,omitempty" codec:"manager_id,omitempty" bson:"manager_id" yaml:"manager_id,omitempty" faker:"-"`
	// Name name of the user
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"person"`
	// PrimaryTeamID the team id of the user's primary team
	PrimaryTeamID *string `json:"primary_team_id,omitempty" codec:"primary_team_id,omitempty" bson:"primary_team_id" yaml:"primary_team_id,omitempty" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// TerminatedDate when the user was terminated in epoch timestamp
	TerminatedDate UserTerminatedDate `json:"terminated_date" codec:"terminated_date" bson:"terminated_date" yaml:"terminated_date" faker:"-"`
	// Title the title of the user
	Title *string `json:"title,omitempty" codec:"title,omitempty" bson:"title" yaml:"title,omitempty" faker:"jobtitle"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
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

	case UserHiredDate:
		return v.ToMap()

	case UserTerminatedDate:
		return v.ToMap()

	default:
		return o
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
func NewUserID(customerID string, Email string) string {
	return hash.Values(customerID, Email)
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
	if o.Title == nil {
		o.Title = &emptyString
	}

	if o.ID == "" {
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
	var i interface{} = o.CustomerID
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
func (o *User) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"avatar_url":      toUserObject(o.AvatarURL, true),
		"cost_center_id":  toUserObject(o.CostCenterID, true),
		"customer_id":     toUserObject(o.CustomerID, false),
		"email":           toUserObject(o.Email, false),
		"hired_date":      toUserObject(o.HiredDate, false),
		"id":              toUserObject(o.ID, false),
		"location":        toUserObject(o.Location, true),
		"manager_id":      toUserObject(o.ManagerID, true),
		"name":            toUserObject(o.Name, false),
		"primary_team_id": toUserObject(o.PrimaryTeamID, true),
		"ref_id":          toUserObject(o.RefID, false),
		"ref_type":        toUserObject(o.RefType, false),
		"terminated_date": toUserObject(o.TerminatedDate, false),
		"title":           toUserObject(o.Title, true),
		"updated_ts":      toUserObject(o.UpdatedAt, false),
		"hashcode":        toUserObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *User) FromMap(kv map[string]interface{}) {

	o.ID = ""

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

	if val, ok := kv["cost_center_id"].(*string); ok {
		o.CostCenterID = val
	} else if val, ok := kv["cost_center_id"].(string); ok {
		o.CostCenterID = &val
	} else {
		if val, ok := kv["cost_center_id"]; ok {
			if val == nil {
				o.CostCenterID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.CostCenterID = pstrings.Pointer(fmt.Sprintf("%v", val))
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
				// if coming in as map, convert it back
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
				// if coming in as map, convert it back
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

	if val, ok := kv["primary_team_id"].(*string); ok {
		o.PrimaryTeamID = val
	} else if val, ok := kv["primary_team_id"].(string); ok {
		o.PrimaryTeamID = &val
	} else {
		if val, ok := kv["primary_team_id"]; ok {
			if val == nil {
				o.PrimaryTeamID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
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
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Title = pstrings.Pointer(fmt.Sprintf("%v", val))
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
	args = append(args, o.AvatarURL)
	args = append(args, o.CostCenterID)
	args = append(args, o.CustomerID)
	args = append(args, o.Email)
	args = append(args, o.HiredDate)
	args = append(args, o.ID)
	args = append(args, o.Location)
	args = append(args, o.ManagerID)
	args = append(args, o.Name)
	args = append(args, o.PrimaryTeamID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.TerminatedDate)
	args = append(args, o.Title)
	args = append(args, o.UpdatedAt)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
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
