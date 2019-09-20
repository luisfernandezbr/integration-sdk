// DO NOT EDIT -- generated code

// Package agent - the agent communicates with pinpoint cloud to send integration data
package agent

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/bxcodec/faker"
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
	"github.com/pinpt/go-common/slice"
	pstrings "github.com/pinpt/go-common/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// UserResponseTopic is the default topic name
	UserResponseTopic datamodel.TopicNameType = "agent_UserResponse_topic"

	// UserResponseStream is the default stream name
	UserResponseStream datamodel.TopicNameType = "agent_UserResponse_stream"

	// UserResponseTable is the default table name
	UserResponseTable datamodel.TopicNameType = "agent_userresponse"

	// UserResponseModelName is the model name
	UserResponseModelName datamodel.ModelNameType = "agent.UserResponse"
)

const (
	// UserResponseArchitectureColumn is the architecture column name
	UserResponseArchitectureColumn = "architecture"
	// UserResponseCustomerIDColumn is the customer_id column name
	UserResponseCustomerIDColumn = "customer_id"
	// UserResponseDataColumn is the data column name
	UserResponseDataColumn = "data"
	// UserResponseDistroColumn is the distro column name
	UserResponseDistroColumn = "distro"
	// UserResponseErrorColumn is the error column name
	UserResponseErrorColumn = "error"
	// UserResponseEventDateColumn is the event_date column name
	UserResponseEventDateColumn = "event_date"
	// UserResponseEventDateColumnEpochColumn is the epoch column property of the EventDate name
	UserResponseEventDateColumnEpochColumn = "event_date->epoch"
	// UserResponseEventDateColumnOffsetColumn is the offset column property of the EventDate name
	UserResponseEventDateColumnOffsetColumn = "event_date->offset"
	// UserResponseEventDateColumnRfc3339Column is the rfc3339 column property of the EventDate name
	UserResponseEventDateColumnRfc3339Column = "event_date->rfc3339"
	// UserResponseFreeSpaceColumn is the free_space column name
	UserResponseFreeSpaceColumn = "free_space"
	// UserResponseGoVersionColumn is the go_version column name
	UserResponseGoVersionColumn = "go_version"
	// UserResponseHostnameColumn is the hostname column name
	UserResponseHostnameColumn = "hostname"
	// UserResponseIDColumn is the id column name
	UserResponseIDColumn = "id"
	// UserResponseIntegrationIDColumn is the integration_id column name
	UserResponseIntegrationIDColumn = "integration_id"
	// UserResponseLastExportDateColumn is the last_export_date column name
	UserResponseLastExportDateColumn = "last_export_date"
	// UserResponseLastExportDateColumnEpochColumn is the epoch column property of the LastExportDate name
	UserResponseLastExportDateColumnEpochColumn = "last_export_date->epoch"
	// UserResponseLastExportDateColumnOffsetColumn is the offset column property of the LastExportDate name
	UserResponseLastExportDateColumnOffsetColumn = "last_export_date->offset"
	// UserResponseLastExportDateColumnRfc3339Column is the rfc3339 column property of the LastExportDate name
	UserResponseLastExportDateColumnRfc3339Column = "last_export_date->rfc3339"
	// UserResponseMemoryColumn is the memory column name
	UserResponseMemoryColumn = "memory"
	// UserResponseMessageColumn is the message column name
	UserResponseMessageColumn = "message"
	// UserResponseNumCPUColumn is the num_cpu column name
	UserResponseNumCPUColumn = "num_cpu"
	// UserResponseOSColumn is the os column name
	UserResponseOSColumn = "os"
	// UserResponseRefIDColumn is the ref_id column name
	UserResponseRefIDColumn = "ref_id"
	// UserResponseRefTypeColumn is the ref_type column name
	UserResponseRefTypeColumn = "ref_type"
	// UserResponseRequestIDColumn is the request_id column name
	UserResponseRequestIDColumn = "request_id"
	// UserResponseSuccessColumn is the success column name
	UserResponseSuccessColumn = "success"
	// UserResponseSystemIDColumn is the system_id column name
	UserResponseSystemIDColumn = "system_id"
	// UserResponseTypeColumn is the type column name
	UserResponseTypeColumn = "type"
	// UserResponseUpdatedAtColumn is the updated_ts column name
	UserResponseUpdatedAtColumn = "updated_ts"
	// UserResponseUptimeColumn is the uptime column name
	UserResponseUptimeColumn = "uptime"
	// UserResponseUsersColumn is the users column name
	UserResponseUsersColumn = "users"
	// UserResponseUsersColumnActiveColumn is the active column property of the Users name
	UserResponseUsersColumnActiveColumn = "users->active"
	// UserResponseUsersColumnAssociatedRefIDColumn is the associated_ref_id column property of the Users name
	UserResponseUsersColumnAssociatedRefIDColumn = "users->associated_ref_id"
	// UserResponseUsersColumnAvatarURLColumn is the avatar_url column property of the Users name
	UserResponseUsersColumnAvatarURLColumn = "users->avatar_url"
	// UserResponseUsersColumnCustomerIDColumn is the customer_id column property of the Users name
	UserResponseUsersColumnCustomerIDColumn = "users->customer_id"
	// UserResponseUsersColumnEmailsColumn is the emails column property of the Users name
	UserResponseUsersColumnEmailsColumn = "users->emails"
	// UserResponseUsersColumnGroupsColumn is the groups column property of the Users name
	UserResponseUsersColumnGroupsColumn = "users->groups"
	// UserResponseUsersColumnIDColumn is the id column property of the Users name
	UserResponseUsersColumnIDColumn = "users->id"
	// UserResponseUsersColumnNameColumn is the name column property of the Users name
	UserResponseUsersColumnNameColumn = "users->name"
	// UserResponseUsersColumnRefIDColumn is the ref_id column property of the Users name
	UserResponseUsersColumnRefIDColumn = "users->ref_id"
	// UserResponseUsersColumnRefTypeColumn is the ref_type column property of the Users name
	UserResponseUsersColumnRefTypeColumn = "users->ref_type"
	// UserResponseUsersColumnUsernameColumn is the username column property of the Users name
	UserResponseUsersColumnUsernameColumn = "users->username"
	// UserResponseUUIDColumn is the uuid column name
	UserResponseUUIDColumn = "uuid"
	// UserResponseVersionColumn is the version column name
	UserResponseVersionColumn = "version"
)

// UserResponseEventDate represents the object structure for event_date
type UserResponseEventDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toUserResponseEventDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *UserResponseEventDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *UserResponseEventDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toUserResponseEventDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toUserResponseEventDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toUserResponseEventDateObject(o.Rfc3339, false),
	}
}

func (o *UserResponseEventDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *UserResponseEventDate) FromMap(kv map[string]interface{}) {

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

// UserResponseLastExportDate represents the object structure for last_export_date
type UserResponseLastExportDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toUserResponseLastExportDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *UserResponseLastExportDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *UserResponseLastExportDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toUserResponseLastExportDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toUserResponseLastExportDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toUserResponseLastExportDateObject(o.Rfc3339, false),
	}
}

func (o *UserResponseLastExportDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *UserResponseLastExportDate) FromMap(kv map[string]interface{}) {

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

// UserResponseType is the enumeration type for type
type UserResponseType int32

// String returns the string value for Type
func (v UserResponseType) String() string {
	switch int32(v) {
	case 0:
		return "ENROLL"
	case 1:
		return "PING"
	case 2:
		return "CRASH"
	case 3:
		return "LOG"
	case 4:
		return "INTEGRATION"
	case 5:
		return "EXPORT"
	case 6:
		return "PROJECT"
	case 7:
		return "REPO"
	case 8:
		return "USER"
	case 9:
		return "UNINSTALL"
	case 10:
		return "UPGRADE"
	case 11:
		return "START"
	case 12:
		return "STOP"
	}
	return "unset"
}

const (
	// TypeEnroll is the enumeration value for enroll
	UserResponseTypeEnroll UserResponseType = 0
	// TypePing is the enumeration value for ping
	UserResponseTypePing UserResponseType = 1
	// TypeCrash is the enumeration value for crash
	UserResponseTypeCrash UserResponseType = 2
	// TypeLog is the enumeration value for log
	UserResponseTypeLog UserResponseType = 3
	// TypeIntegration is the enumeration value for integration
	UserResponseTypeIntegration UserResponseType = 4
	// TypeExport is the enumeration value for export
	UserResponseTypeExport UserResponseType = 5
	// TypeProject is the enumeration value for project
	UserResponseTypeProject UserResponseType = 6
	// TypeRepo is the enumeration value for repo
	UserResponseTypeRepo UserResponseType = 7
	// TypeUser is the enumeration value for user
	UserResponseTypeUser UserResponseType = 8
	// TypeUninstall is the enumeration value for uninstall
	UserResponseTypeUninstall UserResponseType = 9
	// TypeUpgrade is the enumeration value for upgrade
	UserResponseTypeUpgrade UserResponseType = 10
	// TypeStart is the enumeration value for start
	UserResponseTypeStart UserResponseType = 11
	// TypeStop is the enumeration value for stop
	UserResponseTypeStop UserResponseType = 12
)

// UserResponseUsersGroups represents the object structure for groups
type UserResponseUsersGroups struct {
	// GroupID Group id
	GroupID string `json:"group_id" codec:"group_id" bson:"group_id" yaml:"group_id" faker:"-"`
	// Name Group name
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
}

func toUserResponseUsersGroupsObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *UserResponseUsersGroups:
		return v.ToMap()

	default:
		return o
	}
}

func (o *UserResponseUsersGroups) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// GroupID Group id
		"group_id": toUserResponseUsersGroupsObject(o.GroupID, false),
		// Name Group name
		"name": toUserResponseUsersGroupsObject(o.Name, false),
	}
}

func (o *UserResponseUsersGroups) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *UserResponseUsersGroups) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["group_id"].(string); ok {
		o.GroupID = val
	} else {
		if val, ok := kv["group_id"]; ok {
			if val == nil {
				o.GroupID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.GroupID = fmt.Sprintf("%v", val)
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
	o.setDefaults(false)
}

// UserResponseUsers represents the object structure for users
type UserResponseUsers struct {
	// Active if user is active
	Active bool `json:"active" codec:"active" bson:"active" yaml:"active" faker:"-"`
	// AssociatedRefID the ref id associated for this user in another system
	AssociatedRefID *string `json:"associated_ref_id,omitempty" codec:"associated_ref_id,omitempty" bson:"associated_ref_id" yaml:"associated_ref_id,omitempty" faker:"-"`
	// AvatarURL the url to users avatar
	AvatarURL *string `json:"avatar_url,omitempty" codec:"avatar_url,omitempty" bson:"avatar_url" yaml:"avatar_url,omitempty" faker:"avatar"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Emails the email for the user
	Emails []string `json:"emails" codec:"emails" bson:"emails" yaml:"emails" faker:"-"`
	// Groups Group names
	Groups []UserResponseUsersGroups `json:"groups" codec:"groups" bson:"groups" yaml:"groups" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// Name the name of the user
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"person"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// Username the username of the user
	Username string `json:"username" codec:"username" bson:"username" yaml:"username" faker:"username"`
}

func toUserResponseUsersObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *UserResponseUsers:
		return v.ToMap()

	case []UserResponseUsersGroups:
		arr := make([]interface{}, 0)
		for _, i := range v {
			arr = append(arr, i.ToMap())
		}
		return arr

	default:
		return o
	}
}

func (o *UserResponseUsers) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Active if user is active
		"active": toUserResponseUsersObject(o.Active, false),
		// AssociatedRefID the ref id associated for this user in another system
		"associated_ref_id": toUserResponseUsersObject(o.AssociatedRefID, true),
		// AvatarURL the url to users avatar
		"avatar_url": toUserResponseUsersObject(o.AvatarURL, true),
		// CustomerID the customer id for the model instance
		"customer_id": toUserResponseUsersObject(o.CustomerID, false),
		// Emails the email for the user
		"emails": toUserResponseUsersObject(o.Emails, false),
		// Groups Group names
		"groups": toUserResponseUsersObject(o.Groups, false),
		// ID the primary key for the model instance
		"id": toUserResponseUsersObject(o.ID, false),
		// Name the name of the user
		"name": toUserResponseUsersObject(o.Name, false),
		// RefID the source system id for the model instance
		"ref_id": toUserResponseUsersObject(o.RefID, false),
		// RefType the source system identifier for the model instance
		"ref_type": toUserResponseUsersObject(o.RefType, false),
		// Username the username of the user
		"username": toUserResponseUsersObject(o.Username, false),
	}
}

func (o *UserResponseUsers) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *UserResponseUsers) FromMap(kv map[string]interface{}) {

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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.CustomerID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["emails"]; ok {
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
								panic("unsupported type for emails field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for emails field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for emails field")
				}
			}
			o.Emails = na
		}
	}
	if o.Emails == nil {
		o.Emails = make([]string, 0)
	}

	if o == nil {

		o.Groups = make([]UserResponseUsersGroups, 0)

	}
	if val, ok := kv["groups"]; ok {
		if sv, ok := val.([]UserResponseUsersGroups); ok {
			o.Groups = sv
		} else if sp, ok := val.([]*UserResponseUsersGroups); ok {
			o.Groups = o.Groups[:0]
			for _, e := range sp {
				o.Groups = append(o.Groups, *e)
			}
		} else if a, ok := val.(primitive.A); ok {
			for _, ae := range a {
				if av, ok := ae.(UserResponseUsersGroups); ok {
					o.Groups = append(o.Groups, av)
				} else if av, ok := ae.(primitive.M); ok {
					var fm UserResponseUsersGroups
					fm.FromMap(av)
					o.Groups = append(o.Groups, fm)
				} else {
					b, _ := json.Marshal(ae)
					var av UserResponseUsersGroups
					if err := json.Unmarshal(b, &av); err != nil {
						panic("unsupported type for groups field entry: " + reflect.TypeOf(ae).String())
					}
					o.Groups = append(o.Groups, av)
				}
			}
		} else if arr, ok := val.([]interface{}); ok {
			for _, item := range arr {
				if r, ok := item.(UserResponseUsersGroups); ok {
					o.Groups = append(o.Groups, r)
				} else if r, ok := item.(map[string]interface{}); ok {
					var fm UserResponseUsersGroups
					fm.FromMap(r)
					o.Groups = append(o.Groups, fm)
				} else if r, ok := item.(primitive.M); ok {
					fm := UserResponseUsersGroups{}
					fm.FromMap(r)
					o.Groups = append(o.Groups, fm)
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
							var fm UserResponseUsersGroups
							fm.FromMap(m[0].Interface().(map[string]interface{}))
							o.Groups = append(o.Groups, fm)
						}
					}
				}
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.ID = fmt.Sprintf("%v", val)
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

	if val, ok := kv["username"].(string); ok {
		o.Username = val
	} else {
		if val, ok := kv["username"]; ok {
			if val == nil {
				o.Username = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Username = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// UserResponse an agent response to an action request adding user(s)
type UserResponse struct {
	// Architecture the architecture of the agent machine
	Architecture string `json:"architecture" codec:"architecture" bson:"architecture" yaml:"architecture" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Data extra data that is specific about this event
	Data *string `json:"data,omitempty" codec:"data,omitempty" bson:"data" yaml:"data,omitempty" faker:"-"`
	// Distro the agent os distribution
	Distro string `json:"distro" codec:"distro" bson:"distro" yaml:"distro" faker:"-"`
	// Error an error message related to this event
	Error *string `json:"error,omitempty" codec:"error,omitempty" bson:"error" yaml:"error,omitempty" faker:"-"`
	// EventDate the date of the event
	EventDate UserResponseEventDate `json:"event_date" codec:"event_date" bson:"event_date" yaml:"event_date" faker:"-"`
	// FreeSpace the amount of free space in bytes for the agent machine
	FreeSpace int64 `json:"free_space" codec:"free_space" bson:"free_space" yaml:"free_space" faker:"-"`
	// GoVersion the go version that the agent build was built with
	GoVersion string `json:"go_version" codec:"go_version" bson:"go_version" yaml:"go_version" faker:"-"`
	// Hostname the agent hostname
	Hostname string `json:"hostname" codec:"hostname" bson:"hostname" yaml:"hostname" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationID the integration id
	IntegrationID string `json:"integration_id" codec:"integration_id" bson:"integration_id" yaml:"integration_id" faker:"-"`
	// LastExportDate the last export date
	LastExportDate UserResponseLastExportDate `json:"last_export_date" codec:"last_export_date" bson:"last_export_date" yaml:"last_export_date" faker:"-"`
	// Memory the amount of memory in bytes for the agent machine
	Memory int64 `json:"memory" codec:"memory" bson:"memory" yaml:"memory" faker:"-"`
	// Message a message related to this event
	Message string `json:"message" codec:"message" bson:"message" yaml:"message" faker:"-"`
	// NumCPU the number of CPU the agent is running
	NumCPU int64 `json:"num_cpu" codec:"num_cpu" bson:"num_cpu" yaml:"num_cpu" faker:"-"`
	// OS the agent operating system
	OS string `json:"os" codec:"os" bson:"os" yaml:"os" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RequestID the request id that this response is correlated to
	RequestID string `json:"request_id" codec:"request_id" bson:"request_id" yaml:"request_id" faker:"-"`
	// Success if the response was successful
	Success bool `json:"success" codec:"success" bson:"success" yaml:"success" faker:"-"`
	// SystemID system unique device ID
	SystemID string `json:"system_id" codec:"system_id" bson:"system_id" yaml:"system_id" faker:"-"`
	// Type the type of event
	Type UserResponseType `json:"type" codec:"type" bson:"type" yaml:"type" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// Uptime the uptime in milliseconds since the agent started
	Uptime int64 `json:"uptime" codec:"uptime" bson:"uptime" yaml:"uptime" faker:"-"`
	// Users the exported users
	Users []UserResponseUsers `json:"users" codec:"users" bson:"users" yaml:"users" faker:"-"`
	// UUID the agent unique identifier
	UUID string `json:"uuid" codec:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
	// Version the agent version
	Version string `json:"version" codec:"version" bson:"version" yaml:"version" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*UserResponse)(nil)

func toUserResponseObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *UserResponse:
		return v.ToMap()

	case UserResponseEventDate:
		return v.ToMap()

	case UserResponseLastExportDate:
		return v.ToMap()

	case UserResponseType:
		return v.String()

	case []UserResponseUsers:
		arr := make([]interface{}, 0)
		for _, i := range v {
			arr = append(arr, i.ToMap())
		}
		return arr

	default:
		return o
	}
}

// String returns a string representation of UserResponse
func (o *UserResponse) String() string {
	return fmt.Sprintf("agent.UserResponse<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *UserResponse) GetTopicName() datamodel.TopicNameType {
	return UserResponseTopic
}

// GetModelName returns the name of the model
func (o *UserResponse) GetModelName() datamodel.ModelNameType {
	return UserResponseModelName
}

// GetStreamName returns the name of the stream
func (o *UserResponse) GetStreamName() string {
	return UserResponseStream.String()
}

// GetTableName returns the name of the table
func (o *UserResponse) GetTableName() string {
	return UserResponseTable.String()
}

// NewUserResponseID provides a template for generating an ID field for UserResponse
func NewUserResponseID(customerID string, refType string, refID string) string {
	return hash.Values("UserResponse", customerID, refType, refID)
}

func (o *UserResponse) setDefaults(frommap bool) {
	if o.Data == nil {
		o.Data = &emptyString
	}
	if o.Error == nil {
		o.Error = &emptyString
	}
	if o.Users == nil {
		o.Users = make([]UserResponseUsers, 0)
	}

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("UserResponse", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *UserResponse) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *UserResponse) GetTopicKey() string {
	var i interface{} = o.UUID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *UserResponse) GetTimestamp() time.Time {
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
	panic("not sure how to handle the date time format for UserResponse")
}

// GetRefID returns the RefID for the object
func (o *UserResponse) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *UserResponse) IsMaterialized() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *UserResponse) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *UserResponse) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *UserResponse) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = UserResponseModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *UserResponse) GetTopicConfig() *datamodel.ModelTopicConfig {
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
		Key:               "uuid",
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
func (o *UserResponse) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of UserResponse
func (o *UserResponse) Clone() datamodel.Model {
	c := new(UserResponse)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *UserResponse) Anon() datamodel.Model {
	c := new(UserResponse)
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
func (o *UserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *UserResponse) UnmarshalJSON(data []byte) error {
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
func (o *UserResponse) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two UserResponse objects are equal
func (o *UserResponse) IsEqual(other *UserResponse) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *UserResponse) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"architecture":     toUserResponseObject(o.Architecture, false),
		"customer_id":      toUserResponseObject(o.CustomerID, false),
		"data":             toUserResponseObject(o.Data, true),
		"distro":           toUserResponseObject(o.Distro, false),
		"error":            toUserResponseObject(o.Error, true),
		"event_date":       toUserResponseObject(o.EventDate, false),
		"free_space":       toUserResponseObject(o.FreeSpace, false),
		"go_version":       toUserResponseObject(o.GoVersion, false),
		"hostname":         toUserResponseObject(o.Hostname, false),
		"id":               toUserResponseObject(o.ID, false),
		"integration_id":   toUserResponseObject(o.IntegrationID, false),
		"last_export_date": toUserResponseObject(o.LastExportDate, false),
		"memory":           toUserResponseObject(o.Memory, false),
		"message":          toUserResponseObject(o.Message, false),
		"num_cpu":          toUserResponseObject(o.NumCPU, false),
		"os":               toUserResponseObject(o.OS, false),
		"ref_id":           toUserResponseObject(o.RefID, false),
		"ref_type":         toUserResponseObject(o.RefType, false),
		"request_id":       toUserResponseObject(o.RequestID, false),
		"success":          toUserResponseObject(o.Success, false),
		"system_id":        toUserResponseObject(o.SystemID, false),
		"type":             toUserResponseObject(o.Type, false),
		"updated_ts":       toUserResponseObject(o.UpdatedAt, false),
		"uptime":           toUserResponseObject(o.Uptime, false),
		"users":            toUserResponseObject(o.Users, false),
		"uuid":             toUserResponseObject(o.UUID, false),
		"version":          toUserResponseObject(o.Version, false),
		"hashcode":         toUserResponseObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *UserResponse) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["architecture"].(string); ok {
		o.Architecture = val
	} else {
		if val, ok := kv["architecture"]; ok {
			if val == nil {
				o.Architecture = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Architecture = fmt.Sprintf("%v", val)
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

	if val, ok := kv["data"].(*string); ok {
		o.Data = val
	} else if val, ok := kv["data"].(string); ok {
		o.Data = &val
	} else {
		if val, ok := kv["data"]; ok {
			if val == nil {
				o.Data = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Data = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if val, ok := kv["distro"].(string); ok {
		o.Distro = val
	} else {
		if val, ok := kv["distro"]; ok {
			if val == nil {
				o.Distro = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Distro = fmt.Sprintf("%v", val)
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

	if val, ok := kv["event_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.EventDate.FromMap(kv)
		} else if sv, ok := val.(UserResponseEventDate); ok {
			// struct
			o.EventDate = sv
		} else if sp, ok := val.(*UserResponseEventDate); ok {
			// struct pointer
			o.EventDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.EventDate.Epoch = dt.Epoch
			o.EventDate.Rfc3339 = dt.Rfc3339
			o.EventDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.EventDate.Epoch = dt.Epoch
			o.EventDate.Rfc3339 = dt.Rfc3339
			o.EventDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.EventDate.Epoch = dt.Epoch
				o.EventDate.Rfc3339 = dt.Rfc3339
				o.EventDate.Offset = dt.Offset
			}
		}
	} else {
		o.EventDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["free_space"].(int64); ok {
		o.FreeSpace = val
	} else {
		if val, ok := kv["free_space"]; ok {
			if val == nil {
				o.FreeSpace = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.FreeSpace = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["go_version"].(string); ok {
		o.GoVersion = val
	} else {
		if val, ok := kv["go_version"]; ok {
			if val == nil {
				o.GoVersion = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.GoVersion = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["hostname"].(string); ok {
		o.Hostname = val
	} else {
		if val, ok := kv["hostname"]; ok {
			if val == nil {
				o.Hostname = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Hostname = fmt.Sprintf("%v", val)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.ID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["integration_id"].(string); ok {
		o.IntegrationID = val
	} else {
		if val, ok := kv["integration_id"]; ok {
			if val == nil {
				o.IntegrationID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.IntegrationID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["last_export_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastExportDate.FromMap(kv)
		} else if sv, ok := val.(UserResponseLastExportDate); ok {
			// struct
			o.LastExportDate = sv
		} else if sp, ok := val.(*UserResponseLastExportDate); ok {
			// struct pointer
			o.LastExportDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.LastExportDate.Epoch = dt.Epoch
			o.LastExportDate.Rfc3339 = dt.Rfc3339
			o.LastExportDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.LastExportDate.Epoch = dt.Epoch
			o.LastExportDate.Rfc3339 = dt.Rfc3339
			o.LastExportDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.LastExportDate.Epoch = dt.Epoch
				o.LastExportDate.Rfc3339 = dt.Rfc3339
				o.LastExportDate.Offset = dt.Offset
			}
		}
	} else {
		o.LastExportDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["memory"].(int64); ok {
		o.Memory = val
	} else {
		if val, ok := kv["memory"]; ok {
			if val == nil {
				o.Memory = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Memory = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["message"].(string); ok {
		o.Message = val
	} else {
		if val, ok := kv["message"]; ok {
			if val == nil {
				o.Message = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Message = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["num_cpu"].(int64); ok {
		o.NumCPU = val
	} else {
		if val, ok := kv["num_cpu"]; ok {
			if val == nil {
				o.NumCPU = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.NumCPU = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["os"].(string); ok {
		o.OS = val
	} else {
		if val, ok := kv["os"]; ok {
			if val == nil {
				o.OS = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.OS = fmt.Sprintf("%v", val)
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

	if val, ok := kv["request_id"].(string); ok {
		o.RequestID = val
	} else {
		if val, ok := kv["request_id"]; ok {
			if val == nil {
				o.RequestID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.RequestID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["success"].(bool); ok {
		o.Success = val
	} else {
		if val, ok := kv["success"]; ok {
			if val == nil {
				o.Success = number.ToBoolAny(nil)
			} else {
				o.Success = number.ToBoolAny(val)
			}
		}
	}

	if val, ok := kv["system_id"].(string); ok {
		o.SystemID = val
	} else {
		if val, ok := kv["system_id"]; ok {
			if val == nil {
				o.SystemID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.SystemID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["type"].(UserResponseType); ok {
		o.Type = val
	} else {
		if em, ok := kv["type"].(map[string]interface{}); ok {
			ev := em["agent.type"].(string)
			switch ev {
			case "enroll", "ENROLL":
				o.Type = 0
			case "ping", "PING":
				o.Type = 1
			case "crash", "CRASH":
				o.Type = 2
			case "log", "LOG":
				o.Type = 3
			case "integration", "INTEGRATION":
				o.Type = 4
			case "export", "EXPORT":
				o.Type = 5
			case "project", "PROJECT":
				o.Type = 6
			case "repo", "REPO":
				o.Type = 7
			case "user", "USER":
				o.Type = 8
			case "uninstall", "UNINSTALL":
				o.Type = 9
			case "upgrade", "UPGRADE":
				o.Type = 10
			case "start", "START":
				o.Type = 11
			case "stop", "STOP":
				o.Type = 12
			}
		}
		if em, ok := kv["type"].(string); ok {
			switch em {
			case "enroll", "ENROLL":
				o.Type = 0
			case "ping", "PING":
				o.Type = 1
			case "crash", "CRASH":
				o.Type = 2
			case "log", "LOG":
				o.Type = 3
			case "integration", "INTEGRATION":
				o.Type = 4
			case "export", "EXPORT":
				o.Type = 5
			case "project", "PROJECT":
				o.Type = 6
			case "repo", "REPO":
				o.Type = 7
			case "user", "USER":
				o.Type = 8
			case "uninstall", "UNINSTALL":
				o.Type = 9
			case "upgrade", "UPGRADE":
				o.Type = 10
			case "start", "START":
				o.Type = 11
			case "stop", "STOP":
				o.Type = 12
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

	if val, ok := kv["uptime"].(int64); ok {
		o.Uptime = val
	} else {
		if val, ok := kv["uptime"]; ok {
			if val == nil {
				o.Uptime = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Uptime = number.ToInt64Any(val)
			}
		}
	}

	if o == nil {

		o.Users = make([]UserResponseUsers, 0)

	}
	if val, ok := kv["users"]; ok {
		if sv, ok := val.([]UserResponseUsers); ok {
			o.Users = sv
		} else if sp, ok := val.([]*UserResponseUsers); ok {
			o.Users = o.Users[:0]
			for _, e := range sp {
				o.Users = append(o.Users, *e)
			}
		} else if a, ok := val.(primitive.A); ok {
			for _, ae := range a {
				if av, ok := ae.(UserResponseUsers); ok {
					o.Users = append(o.Users, av)
				} else if av, ok := ae.(primitive.M); ok {
					var fm UserResponseUsers
					fm.FromMap(av)
					o.Users = append(o.Users, fm)
				} else {
					b, _ := json.Marshal(ae)
					var av UserResponseUsers
					if err := json.Unmarshal(b, &av); err != nil {
						panic("unsupported type for users field entry: " + reflect.TypeOf(ae).String())
					}
					o.Users = append(o.Users, av)
				}
			}
		} else if arr, ok := val.([]interface{}); ok {
			for _, item := range arr {
				if r, ok := item.(UserResponseUsers); ok {
					o.Users = append(o.Users, r)
				} else if r, ok := item.(map[string]interface{}); ok {
					var fm UserResponseUsers
					fm.FromMap(r)
					o.Users = append(o.Users, fm)
				} else if r, ok := item.(primitive.M); ok {
					fm := UserResponseUsers{}
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
							var fm UserResponseUsers
							fm.FromMap(m[0].Interface().(map[string]interface{}))
							o.Users = append(o.Users, fm)
						}
					}
				}
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.UUID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["version"].(string); ok {
		o.Version = val
	} else {
		if val, ok := kv["version"]; ok {
			if val == nil {
				o.Version = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Version = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *UserResponse) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Architecture)
	args = append(args, o.CustomerID)
	args = append(args, o.Data)
	args = append(args, o.Distro)
	args = append(args, o.Error)
	args = append(args, o.EventDate)
	args = append(args, o.FreeSpace)
	args = append(args, o.GoVersion)
	args = append(args, o.Hostname)
	args = append(args, o.ID)
	args = append(args, o.IntegrationID)
	args = append(args, o.LastExportDate)
	args = append(args, o.Memory)
	args = append(args, o.Message)
	args = append(args, o.NumCPU)
	args = append(args, o.OS)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RequestID)
	args = append(args, o.Success)
	args = append(args, o.SystemID)
	args = append(args, o.Type)
	args = append(args, o.UpdatedAt)
	args = append(args, o.Uptime)
	args = append(args, o.Users)
	args = append(args, o.UUID)
	args = append(args, o.Version)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *UserResponse) GetEventAPIConfig() datamodel.EventAPIConfig {
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
