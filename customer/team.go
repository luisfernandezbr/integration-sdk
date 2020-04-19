// DO NOT EDIT -- generated code

// Package customer - customer specific data
package customer

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/bxcodec/faker"
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/graphql"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
	"github.com/pinpt/go-common/slice"
	pstrings "github.com/pinpt/go-common/strings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (

	// TeamTable is the default table name
	TeamTable datamodel.ModelNameType = "customer_team"

	// TeamModelName is the model name
	TeamModelName datamodel.ModelNameType = "customer.Team"
)

const (
	// TeamModelActiveColumn is the column json value active
	TeamModelActiveColumn = "active"
	// TeamModelChildrenIdsColumn is the column json value children_ids
	TeamModelChildrenIdsColumn = "children_ids"
	// TeamModelCreatedAtColumn is the column json value created_ts
	TeamModelCreatedAtColumn = "created_ts"
	// TeamModelCustomerIDColumn is the column json value customer_id
	TeamModelCustomerIDColumn = "customer_id"
	// TeamModelDeletedColumn is the column json value deleted
	TeamModelDeletedColumn = "deleted"
	// TeamModelDeletedDateColumn is the column json value deleted_date
	TeamModelDeletedDateColumn = "deleted_date"
	// TeamModelDeletedDateEpochColumn is the column json value epoch
	TeamModelDeletedDateEpochColumn = "epoch"
	// TeamModelDeletedDateOffsetColumn is the column json value offset
	TeamModelDeletedDateOffsetColumn = "offset"
	// TeamModelDeletedDateRfc3339Column is the column json value rfc3339
	TeamModelDeletedDateRfc3339Column = "rfc3339"
	// TeamModelDescriptionColumn is the column json value description
	TeamModelDescriptionColumn = "description"
	// TeamModelIDColumn is the column json value id
	TeamModelIDColumn = "id"
	// TeamModelLeafColumn is the column json value leaf
	TeamModelLeafColumn = "leaf"
	// TeamModelNameColumn is the column json value name
	TeamModelNameColumn = "name"
	// TeamModelParentIdsColumn is the column json value parent_ids
	TeamModelParentIdsColumn = "parent_ids"
	// TeamModelProjectIdsColumn is the column json value project_ids
	TeamModelProjectIdsColumn = "project_ids"
	// TeamModelRefIDColumn is the column json value ref_id
	TeamModelRefIDColumn = "ref_id"
	// TeamModelRefTypeColumn is the column json value ref_type
	TeamModelRefTypeColumn = "ref_type"
	// TeamModelRepoIdsColumn is the column json value repo_ids
	TeamModelRepoIdsColumn = "repo_ids"
	// TeamModelUsersTypeColumn is the column json value type
	TeamModelUsersTypeColumn = "type"
	// TeamModelUpdatedAtColumn is the column json value updated_ts
	TeamModelUpdatedAtColumn = "updated_ts"
	// TeamModelUsersColumn is the column json value users
	TeamModelUsersColumn = "users"
	// TeamModelUsersAvatarURLColumn is the column json value avatar_url
	TeamModelUsersAvatarURLColumn = "avatar_url"
	// TeamModelUsersIDColumn is the column json value id
	TeamModelUsersIDColumn = "id"
	// TeamModelUsersNameColumn is the column json value name
	TeamModelUsersNameColumn = "name"
	// TeamModelUsersNicknameColumn is the column json value nickname
	TeamModelUsersNicknameColumn = "nickname"
	// TeamModelUsersProfileIDColumn is the column json value profile_id
	TeamModelUsersProfileIDColumn = "profile_id"
	// TeamModelUsersRefIDColumn is the column json value ref_id
	TeamModelUsersRefIDColumn = "ref_id"
	// TeamModelUsersTeamIDColumn is the column json value team_id
	TeamModelUsersTeamIDColumn = "team_id"
	// TeamModelUsersURLColumn is the column json value url
	TeamModelUsersURLColumn = "url"
)

// TeamDeletedDate represents the object structure for deleted_date
type TeamDeletedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toTeamDeletedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *TeamDeletedDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *TeamDeletedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toTeamDeletedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toTeamDeletedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toTeamDeletedDateObject(o.Rfc3339, false),
	}
}

func (o *TeamDeletedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *TeamDeletedDate) FromMap(kv map[string]interface{}) {

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

// TeamUsersType is the enumeration type for type
type TeamUsersType int32

// UnmarshalBSONValue for unmarshaling value
func (v *TeamUsersType) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = TeamUsersType(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "NONE":
			*v = TeamUsersType(0)
		case "TRACKABLE":
			*v = TeamUsersType(1)
		case "BOT":
			*v = TeamUsersType(2)
		case "UNTRACKABLE":
			*v = TeamUsersType(3)
		case "UNMAPPED":
			*v = TeamUsersType(4)
		case "TERMINATED":
			*v = TeamUsersType(5)
		case "DELETED":
			*v = TeamUsersType(6)
		case "INVITED":
			*v = TeamUsersType(7)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v TeamUsersType) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "NONE":
		v = 0
	case "TRACKABLE":
		v = 1
	case "BOT":
		v = 2
	case "UNTRACKABLE":
		v = 3
	case "UNMAPPED":
		v = 4
	case "TERMINATED":
		v = 5
	case "DELETED":
		v = 6
	case "INVITED":
		v = 7
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v TeamUsersType) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("NONE")
	case 1:
		return json.Marshal("TRACKABLE")
	case 2:
		return json.Marshal("BOT")
	case 3:
		return json.Marshal("UNTRACKABLE")
	case 4:
		return json.Marshal("UNMAPPED")
	case 5:
		return json.Marshal("TERMINATED")
	case 6:
		return json.Marshal("DELETED")
	case 7:
		return json.Marshal("INVITED")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for UsersType
func (v TeamUsersType) String() string {
	switch int32(v) {
	case 0:
		return "NONE"
	case 1:
		return "TRACKABLE"
	case 2:
		return "BOT"
	case 3:
		return "UNTRACKABLE"
	case 4:
		return "UNMAPPED"
	case 5:
		return "TERMINATED"
	case 6:
		return "DELETED"
	case 7:
		return "INVITED"
	}
	return "unset"
}

const (
	// TeamUsersTypeNone is the enumeration value for none
	TeamUsersTypeNone TeamUsersType = 0
	// TeamUsersTypeTrackable is the enumeration value for trackable
	TeamUsersTypeTrackable TeamUsersType = 1
	// TeamUsersTypeBot is the enumeration value for bot
	TeamUsersTypeBot TeamUsersType = 2
	// TeamUsersTypeUntrackable is the enumeration value for untrackable
	TeamUsersTypeUntrackable TeamUsersType = 3
	// TeamUsersTypeUnmapped is the enumeration value for unmapped
	TeamUsersTypeUnmapped TeamUsersType = 4
	// TeamUsersTypeTerminated is the enumeration value for terminated
	TeamUsersTypeTerminated TeamUsersType = 5
	// TeamUsersTypeDeleted is the enumeration value for deleted
	TeamUsersTypeDeleted TeamUsersType = 6
	// TeamUsersTypeInvited is the enumeration value for invited
	TeamUsersTypeInvited TeamUsersType = 7
)

// TeamUsers represents the object structure for users
type TeamUsers struct {
	// AvatarURL the avatar of the user in case there is no relation. this might not be set always if source system doesn't support
	AvatarURL *string `json:"avatar_url,omitempty" codec:"avatar_url,omitempty" bson:"avatar_url" yaml:"avatar_url,omitempty" faker:"-"`
	// ID the corporate user id
	ID string `json:"id" codec:"id" bson:"id" yaml:"id" faker:"-"`
	// Name the name of the user
	Name *string `json:"name,omitempty" codec:"name,omitempty" bson:"name" yaml:"name,omitempty" faker:"-"`
	// Nickname the nick name of the user (usually the first name if they have a profile)
	Nickname *string `json:"nickname,omitempty" codec:"nickname,omitempty" bson:"nickname" yaml:"nickname,omitempty" faker:"-"`
	// ProfileID the universal profile id for this user
	ProfileID *string `json:"profile_id,omitempty" codec:"profile_id,omitempty" bson:"profile_id" yaml:"profile_id,omitempty" faker:"-"`
	// RefID the ref_id of the user in the source system
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// TeamID the corporate team id
	TeamID string `json:"team_id" codec:"team_id" bson:"team_id" yaml:"team_id" faker:"-"`
	// Type the type of user
	Type TeamUsersType `json:"type" codec:"type" bson:"type" yaml:"type" faker:"-"`
	// URL a url to the user in the source system
	URL *string `json:"url,omitempty" codec:"url,omitempty" bson:"url" yaml:"url,omitempty" faker:"-"`
}

func toTeamUsersObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *TeamUsers:
		return v.ToMap()

	case TeamUsersType:
		return v.String()

	default:
		return o
	}
}

func (o *TeamUsers) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// AvatarURL the avatar of the user in case there is no relation. this might not be set always if source system doesn't support
		"avatar_url": toTeamUsersObject(o.AvatarURL, true),
		// ID the corporate user id
		"id": toTeamUsersObject(o.ID, false),
		// Name the name of the user
		"name": toTeamUsersObject(o.Name, true),
		// Nickname the nick name of the user (usually the first name if they have a profile)
		"nickname": toTeamUsersObject(o.Nickname, true),
		// ProfileID the universal profile id for this user
		"profile_id": toTeamUsersObject(o.ProfileID, true),
		// RefID the ref_id of the user in the source system
		"ref_id": toTeamUsersObject(o.RefID, false),
		// TeamID the corporate team id
		"team_id": toTeamUsersObject(o.TeamID, false),
		// Type the type of user
		"type": toTeamUsersObject(o.Type, false),
		// URL a url to the user in the source system
		"url": toTeamUsersObject(o.URL, true),
	}
}

func (o *TeamUsers) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *TeamUsers) FromMap(kv map[string]interface{}) {

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

	if val, ok := kv["nickname"].(*string); ok {
		o.Nickname = val
	} else if val, ok := kv["nickname"].(string); ok {
		o.Nickname = &val
	} else {
		if val, ok := kv["nickname"]; ok {
			if val == nil {
				o.Nickname = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Nickname = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if val, ok := kv["profile_id"].(*string); ok {
		o.ProfileID = val
	} else if val, ok := kv["profile_id"].(string); ok {
		o.ProfileID = &val
	} else {
		if val, ok := kv["profile_id"]; ok {
			if val == nil {
				o.ProfileID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.ProfileID = pstrings.Pointer(fmt.Sprintf("%v", val))
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

	if val, ok := kv["team_id"].(string); ok {
		o.TeamID = val
	} else {
		if val, ok := kv["team_id"]; ok {
			if val == nil {
				o.TeamID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.TeamID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["type"].(TeamUsersType); ok {
		o.Type = val
	} else {
		if em, ok := kv["type"].(map[string]interface{}); ok {

			ev := em["customer.type"].(string)
			switch ev {
			case "none", "NONE":
				o.Type = 0
			case "trackable", "TRACKABLE":
				o.Type = 1
			case "bot", "BOT":
				o.Type = 2
			case "untrackable", "UNTRACKABLE":
				o.Type = 3
			case "unmapped", "UNMAPPED":
				o.Type = 4
			case "terminated", "TERMINATED":
				o.Type = 5
			case "deleted", "DELETED":
				o.Type = 6
			case "invited", "INVITED":
				o.Type = 7
			}
		}
		if em, ok := kv["type"].(string); ok {
			switch em {
			case "none", "NONE":
				o.Type = 0
			case "trackable", "TRACKABLE":
				o.Type = 1
			case "bot", "BOT":
				o.Type = 2
			case "untrackable", "UNTRACKABLE":
				o.Type = 3
			case "unmapped", "UNMAPPED":
				o.Type = 4
			case "terminated", "TERMINATED":
				o.Type = 5
			case "deleted", "DELETED":
				o.Type = 6
			case "invited", "INVITED":
				o.Type = 7
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

// Team a team is a grouping of one or more users
type Team struct {
	// Active whether the team is tracked in pinpoint
	Active bool `json:"active" codec:"active" bson:"active" yaml:"active" faker:"-"`
	// ChildrenIds the children_ids for this team
	ChildrenIds []string `json:"children_ids" codec:"children_ids" bson:"children_ids" yaml:"children_ids" faker:"-"`
	// CreatedAt the date the record was created in Epoch time
	CreatedAt int64 `json:"created_ts" codec:"created_ts" bson:"created_ts" yaml:"created_ts" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Deleted delete flag for a team. true === deleted
	Deleted bool `json:"deleted" codec:"deleted" bson:"deleted" yaml:"deleted" faker:"-"`
	// DeletedDate when the user profile was soft deleted
	DeletedDate TeamDeletedDate `json:"deleted_date" codec:"deleted_date" bson:"deleted_date" yaml:"deleted_date" faker:"-"`
	// Description the description of the team
	Description string `json:"description" codec:"description" bson:"description" yaml:"description" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// Leaf True when team has no children_ids
	Leaf bool `json:"leaf" codec:"leaf" bson:"leaf" yaml:"leaf" faker:"-"`
	// Name the name of the team
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"team"`
	// ParentIds the parent_ids for this team
	ParentIds []string `json:"parent_ids" codec:"parent_ids" bson:"parent_ids" yaml:"parent_ids" faker:"-"`
	// ProjectIds project ids for the team
	ProjectIds []string `json:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RepoIds repo ids for the team
	RepoIds []string `json:"-"`
	// UpdatedAt the date the record was updated in Epoch time
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// Users users for the team
	Users []TeamUsers `json:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Team)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*Team)(nil)

func toTeamObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *Team:
		return v.ToMap()

	case TeamDeletedDate:
		return v.ToMap()

	case []TeamUsers:
		arr := make([]interface{}, 0)
		for _, i := range v {
			arr = append(arr, i.ToMap())
		}
		return arr

	default:
		return o
	}
}

// String returns a string representation of Team
func (o *Team) String() string {
	return fmt.Sprintf("customer.Team<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Team) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *Team) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *Team) GetTableName() string {
	return TeamTable.String()
}

// GetModelName returns the name of the model
func (o *Team) GetModelName() datamodel.ModelNameType {
	return TeamModelName
}

// NewTeamID provides a template for generating an ID field for Team
func NewTeamID(Name string, customerID string) string {
	return hash.Values(Name, customerID, datetime.EpochNow(), randomString(64))
}

func (o *Team) setDefaults(frommap bool) {
	if o.ChildrenIds == nil {
		o.ChildrenIds = make([]string, 0)
	}
	if o.ParentIds == nil {
		o.ParentIds = make([]string, 0)
	}
	if o.ProjectIds == nil {
		o.ProjectIds = make([]string, 0)
	}
	if o.RepoIds == nil {
		o.RepoIds = make([]string, 0)
	}
	if o.Users == nil {
		o.Users = make([]TeamUsers, 0)
	}

	if o.ID == "" {
		// set the id from the spec provided in the model
		o.ID = hash.Values(o.Name, o.CustomerID, datetime.EpochNow(), randomString(64))
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Team) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Team) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Team) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *Team) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Team) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *Team) IsMutable() bool {
	return true
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Team) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Team) IsEvented() bool {
	return false
}

// GetTopicConfig returns the topic config object
func (o *Team) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *Team) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of Team
func (o *Team) Clone() datamodel.Model {
	c := new(Team)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Team) Anon() datamodel.Model {
	c := new(Team)
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
func (o *Team) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Team) UnmarshalJSON(data []byte) error {
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
func (o *Team) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *Team) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two Team objects are equal
func (o *Team) IsEqual(other *Team) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Team) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"active":       toTeamObject(o.Active, false),
		"children_ids": toTeamObject(o.ChildrenIds, false),
		"created_ts":   toTeamObject(o.CreatedAt, false),
		"customer_id":  toTeamObject(o.CustomerID, false),
		"deleted":      toTeamObject(o.Deleted, false),
		"deleted_date": toTeamObject(o.DeletedDate, false),
		"description":  toTeamObject(o.Description, false),
		"id":           toTeamObject(o.ID, false),
		"leaf":         toTeamObject(o.Leaf, false),
		"name":         toTeamObject(o.Name, false),
		"parent_ids":   toTeamObject(o.ParentIds, false),
		"project_ids":  toTeamObject(o.ProjectIds, false),
		"ref_id":       toTeamObject(o.RefID, false),
		"ref_type":     toTeamObject(o.RefType, false),
		"repo_ids":     toTeamObject(o.RepoIds, false),
		"updated_ts":   toTeamObject(o.UpdatedAt, false),
		"users":        toTeamObject(o.Users, false),
		"hashcode":     toTeamObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *Team) FromMap(kv map[string]interface{}) {

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
				o.Active = false
			} else {
				o.Active = number.ToBoolAny(val)
			}
		}
	}

	if val, ok := kv["children_ids"]; ok {
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
								panic("unsupported type for children_ids field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for children_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for children_ids field")
				}
			}
			o.ChildrenIds = na
		}
	}
	if o.ChildrenIds == nil {
		o.ChildrenIds = make([]string, 0)
	}

	if val, ok := kv["created_ts"].(int64); ok {
		o.CreatedAt = val
	} else {
		if val, ok := kv["created_ts"]; ok {
			if val == nil {
				o.CreatedAt = 0
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

	if val, ok := kv["deleted"].(bool); ok {
		o.Deleted = val
	} else {
		if val, ok := kv["deleted"]; ok {
			if val == nil {
				o.Deleted = false
			} else {
				o.Deleted = number.ToBoolAny(val)
			}
		}
	}

	if val, ok := kv["deleted_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.DeletedDate.FromMap(kv)
		} else if sv, ok := val.(TeamDeletedDate); ok {
			// struct
			o.DeletedDate = sv
		} else if sp, ok := val.(*TeamDeletedDate); ok {
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

	if val, ok := kv["description"].(string); ok {
		o.Description = val
	} else {
		if val, ok := kv["description"]; ok {
			if val == nil {
				o.Description = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Description = fmt.Sprintf("%v", val)
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

	if val, ok := kv["leaf"].(bool); ok {
		o.Leaf = val
	} else {
		if val, ok := kv["leaf"]; ok {
			if val == nil {
				o.Leaf = false
			} else {
				o.Leaf = number.ToBoolAny(val)
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

	if val, ok := kv["parent_ids"]; ok {
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
								panic("unsupported type for parent_ids field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for parent_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for parent_ids field")
				}
			}
			o.ParentIds = na
		}
	}
	if o.ParentIds == nil {
		o.ParentIds = make([]string, 0)
	}

	if val, ok := kv["project_ids"]; ok {
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
								panic("unsupported type for project_ids field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for project_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for project_ids field")
				}
			}
			o.ProjectIds = na
		}
	}
	if o.ProjectIds == nil {
		o.ProjectIds = make([]string, 0)
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

	if val, ok := kv["repo_ids"]; ok {
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
								panic("unsupported type for repo_ids field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for repo_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for repo_ids field")
				}
			}
			o.RepoIds = na
		}
	}
	if o.RepoIds == nil {
		o.RepoIds = make([]string, 0)
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

	if o == nil {

		o.Users = make([]TeamUsers, 0)

	}
	if val, ok := kv["users"]; ok {
		if sv, ok := val.([]TeamUsers); ok {
			o.Users = sv
		} else if sp, ok := val.([]*TeamUsers); ok {
			o.Users = o.Users[:0]
			for _, e := range sp {
				o.Users = append(o.Users, *e)
			}
		} else if a, ok := val.(primitive.A); ok {
			for _, ae := range a {
				if av, ok := ae.(TeamUsers); ok {
					o.Users = append(o.Users, av)
				} else if av, ok := ae.(primitive.M); ok {
					var fm TeamUsers
					fm.FromMap(av)
					o.Users = append(o.Users, fm)
				} else {
					b, _ := json.Marshal(ae)
					bkv := make(map[string]interface{})
					json.Unmarshal(b, &bkv)
					var av TeamUsers
					av.FromMap(bkv)
					o.Users = append(o.Users, av)
				}
			}
		} else if arr, ok := val.([]interface{}); ok {
			for _, item := range arr {
				if r, ok := item.(TeamUsers); ok {
					o.Users = append(o.Users, r)
				} else if r, ok := item.(map[string]interface{}); ok {
					var fm TeamUsers
					fm.FromMap(r)
					o.Users = append(o.Users, fm)
				} else if r, ok := item.(primitive.M); ok {
					fm := TeamUsers{}
					fm.FromMap(r)
					o.Users = append(o.Users, fm)
				}
			}
		} else {
			arr := reflect.ValueOf(val)
			if arr.Kind() == reflect.Slice {
				for i := 0; i < arr.Len(); i++ {
					item := arr.Index(i)
					if item.CanAddr() {
						v := item.Addr().MethodByName("ToMap")
						if !v.IsNil() {
							m := v.Call([]reflect.Value{})
							var fm TeamUsers
							fm.FromMap(m[0].Interface().(map[string]interface{}))
							o.Users = append(o.Users, fm)
						}
					}
				}
			}
		}
	}

	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *Team) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Active)
	args = append(args, o.ChildrenIds)
	args = append(args, o.CreatedAt)
	args = append(args, o.CustomerID)
	args = append(args, o.Deleted)
	args = append(args, o.DeletedDate)
	args = append(args, o.Description)
	args = append(args, o.ID)
	args = append(args, o.Leaf)
	args = append(args, o.Name)
	args = append(args, o.ParentIds)
	args = append(args, o.ProjectIds)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RepoIds)
	args = append(args, o.UpdatedAt)
	args = append(args, o.Users)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *Team) GetEventAPIConfig() datamodel.EventAPIConfig {
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

func getTeamQueryFields() string {
	var sb strings.Builder

	// scalar
	sb.WriteString("\t\t\tactive\n")
	// scalar
	sb.WriteString("\t\t\tchildren_ids\n")
	// scalar
	sb.WriteString("\t\t\tcreated_ts\n")
	// scalar
	sb.WriteString("\t\t\tcustomer_id\n")
	// scalar
	sb.WriteString("\t\t\tdeleted\n")
	// object with fields
	sb.WriteString("\t\t\tdeleted_date {\n")

	// scalar
	sb.WriteString("\t\t\tepoch\n")
	// scalar
	sb.WriteString("\t\t\toffset\n")
	// scalar
	sb.WriteString("\t\t\trfc3339\n")
	sb.WriteString("\t\t\t}\n")
	// scalar
	sb.WriteString("\t\t\tdescription\n")
	// id
	sb.WriteString("\t\t\t_id\n")
	// scalar
	sb.WriteString("\t\t\tleaf\n")
	// scalar
	sb.WriteString("\t\t\tname\n")
	// scalar
	sb.WriteString("\t\t\tparent_ids\n")
	// scalar
	sb.WriteString("\t\t\tproject_ids\n")
	// scalar
	sb.WriteString("\t\t\tref_id\n")
	// scalar
	sb.WriteString("\t\t\tref_type\n")
	// scalar
	sb.WriteString("\t\t\trepo_ids\n")
	// scalar
	sb.WriteString("\t\t\tupdated_ts\n")
	// object with fields
	sb.WriteString("\t\t\tusers {\n")

	// scalar
	sb.WriteString("\t\t\tavatar_url\n")
	// scalar
	sb.WriteString("\t\t\tid\n")
	// scalar
	sb.WriteString("\t\t\tname\n")
	// scalar
	sb.WriteString("\t\t\tnickname\n")
	// scalar
	sb.WriteString("\t\t\tprofile_id\n")
	// scalar
	sb.WriteString("\t\t\tref_id\n")
	// scalar
	sb.WriteString("\t\t\tteam_id\n")
	// scalar
	sb.WriteString("\t\t\ttype\n")
	// scalar
	sb.WriteString("\t\t\turl\n")
	sb.WriteString("\t\t\t}\n")
	return sb.String()
}

// TeamPageInfo is a grapqhl PageInfo
type TeamPageInfo struct {
	StartCursor     string `json:"startCursor,omitempty"`
	EndCursor       string `json:"endCursor,omitempty"`
	HasNextPage     bool   `json:"hasNextPage,omitempty"`
	HasPreviousPage bool   `json:"hasPreviousPage,omitempty"`
}

// TeamConnection is a grapqhl connection
type TeamConnection struct {
	Edges      []*TeamEdge  `json:"edges,omitempty"`
	PageInfo   TeamPageInfo `json:"pageInfo,omitempty"`
	TotalCount *int64       `json:"totalCount,omitempty"`
}

// TeamEdge is a grapqhl edge
type TeamEdge struct {
	Node *Team `json:"node,omitempty"`
}

// QueryManyTeamNode is a grapqhl query many node
type QueryManyTeamNode struct {
	Object *TeamConnection `json:"Teams,omitempty"`
}

// QueryManyTeamData is a grapqhl query many data node
type QueryManyTeamData struct {
	Data *QueryManyTeamNode `json:"customer,omitempty"`
}

// QueryOneTeamNode is a grapqhl query one node
type QueryOneTeamNode struct {
	Object *Team `json:"Team,omitempty"`
}

// QueryOneTeamData is a grapqhl query one data node
type QueryOneTeamData struct {
	Data *QueryOneTeamNode `json:"customer,omitempty"`
}

// TeamQuery is query struct
type TeamQuery struct {
	Filters []string      `json:"filters,omitempty"`
	Params  []interface{} `json:"params,omitempty"`
}

// TeamQueryInput is query input struct
type TeamQueryInput struct {
	First  *int64     `json:"first,omitempty"`
	Last   *int64     `json:"last,omitempty"`
	Before *string    `json:"before,omitempty"`
	After  *string    `json:"after,omitempty"`
	Query  *TeamQuery `json:"query,omitempty"`
}

// NewTeamQuery is a convenience for building a *TeamQuery
func NewTeamQuery(params ...interface{}) *TeamQueryInput {
	if len(params)%2 != 0 {
		panic("incorrect number of arguments passed")
	}
	q := &TeamQuery{
		Filters: make([]string, 0),
		Params:  make([]interface{}, 0),
	}
	for i := 0; i < len(params); i += 2 {
		q.Filters = append(q.Filters, params[i].(string))
		q.Params = append(q.Params, params[i+1])
	}
	return &TeamQueryInput{
		Query: q,
	}
}

// FindTeam will query an Team by id
func FindTeam(client graphql.Client, id string) (*Team, error) {
	variables := make(graphql.Variables)
	variables["id"] = id
	var sb strings.Builder
	sb.WriteString("query TeamQuery($id: ID) {\n")
	sb.WriteString("\tcustomer {\n")
	sb.WriteString("\t\tTeam(_id: $id) {\n")
	sb.WriteString(getTeamQueryFields())
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res QueryOneTeamData
	if err := client.Query(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	if res.Data != nil {
		return res.Data.Object, nil
	}
	return nil, nil
}

// FindTeams will query for any Teams matching the query
func FindTeams(client graphql.Client, input *TeamQueryInput) (*TeamConnection, error) {
	variables := make(graphql.Variables)
	if input != nil {
		variables["first"] = input.First
		variables["last"] = input.Last
		variables["before"] = input.Before
		variables["after"] = input.After
		variables["query"] = input.Query
	}
	var sb strings.Builder
	sb.WriteString("query TeamQueryMany($first: Int, $last: Int, $before: Cursor, $after: Cursor, $query: QueryInput) {\n")
	sb.WriteString("\tcustomer {\n")
	sb.WriteString("\t\tTeams(first: $first last: $last before: $before after: $after query: $query) {\n")
	sb.WriteString("\t\t\tpageInfo {\n")
	sb.WriteString("\t\t\t\tstartCursor\n")
	sb.WriteString("\t\t\t\tendCursor\n")
	sb.WriteString("\t\t\t\thasNextPage\n")
	sb.WriteString("\t\t\t\thasPreviousPage\n")
	sb.WriteString("\t\t\t}\n")
	sb.WriteString("\t\t\ttotalCount\n")
	sb.WriteString("\t\t\tedges {\n")
	sb.WriteString("\t\t\t\tnode {\n")
	sb.WriteString(getTeamQueryFields())
	sb.WriteString("\t\t\t\t}\n")
	sb.WriteString("\t\t\t}\n")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res QueryManyTeamData
	if err := client.Query(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	return res.Data.Object, nil
}

// FindTeamsPaginatedCallback is a callback function for handling each page
type FindTeamsPaginatedCallback func(conn *TeamConnection) (bool, error)

// FindTeams will query for any Teams matching the query and return each page callback
func FindTeamsPaginated(client graphql.Client, query *TeamQuery, pageSize int64, callback FindTeamsPaginatedCallback) error {
	input := &TeamQueryInput{
		First: &pageSize,
		Query: query,
	}
	for {
		res, err := FindTeams(client, input)
		if err != nil {
			return err
		}
		if res == nil {
			break
		}
		ok, err := callback(res)
		if err != nil {
			return err
		}
		if !ok || !res.PageInfo.HasNextPage {
			break
		}
		input.After = &res.PageInfo.EndCursor
	}
	return nil
}

// UpdateTeamNode is a grapqhl update node
type UpdateTeamNode struct {
	Object *Team `json:"updateTeam,omitempty"`
}

// UpdateTeamData is a grapqhl update data node
type UpdateTeamData struct {
	Data *UpdateTeamNode `json:"customer,omitempty"`
}

// ExecTeamUpdateMutation returns a graphql update mutation result for Team
func ExecTeamUpdateMutation(client graphql.Client, id string, input graphql.Variables, upsert bool) (*Team, error) {
	variables := make(graphql.Variables)
	variables["id"] = id
	variables["upsert"] = upsert
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation TeamUpdateMutation($id: String, $input: UpdateCustomerTeamInput, $upsert: Boolean) {\n")
	sb.WriteString("\tcustomer {\n")
	sb.WriteString("\t\tupdateTeam(_id: $id, input: $input, upsert: $upsert) {\n")
	sb.WriteString(getTeamQueryFields())
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateTeamData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	return res.Data.Object, nil
}
