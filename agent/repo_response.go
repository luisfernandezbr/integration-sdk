// DO NOT EDIT -- generated code

// Package agent - the agent communicates with pinpoint cloud to send integration data
package agent

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/bxcodec/faker"
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
	pstrings "github.com/pinpt/go-common/strings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (

	// RepoResponseTable is the default table name
	RepoResponseTable datamodel.ModelNameType = "agent_reporesponse"

	// RepoResponseModelName is the model name
	RepoResponseModelName datamodel.ModelNameType = "agent.RepoResponse"
)

const (
	// RepoResponseModelArchitectureColumn is the column json value architecture
	RepoResponseModelArchitectureColumn = "architecture"
	// RepoResponseModelCustomerIDColumn is the column json value customer_id
	RepoResponseModelCustomerIDColumn = "customer_id"
	// RepoResponseModelDataColumn is the column json value data
	RepoResponseModelDataColumn = "data"
	// RepoResponseModelDistroColumn is the column json value distro
	RepoResponseModelDistroColumn = "distro"
	// RepoResponseModelErrorColumn is the column json value error
	RepoResponseModelErrorColumn = "error"
	// RepoResponseModelEventDateColumn is the column json value event_date
	RepoResponseModelEventDateColumn = "event_date"
	// RepoResponseModelEventDateEpochColumn is the column json value epoch
	RepoResponseModelEventDateEpochColumn = "epoch"
	// RepoResponseModelEventDateOffsetColumn is the column json value offset
	RepoResponseModelEventDateOffsetColumn = "offset"
	// RepoResponseModelEventDateRfc3339Column is the column json value rfc3339
	RepoResponseModelEventDateRfc3339Column = "rfc3339"
	// RepoResponseModelFreeSpaceColumn is the column json value free_space
	RepoResponseModelFreeSpaceColumn = "free_space"
	// RepoResponseModelGoVersionColumn is the column json value go_version
	RepoResponseModelGoVersionColumn = "go_version"
	// RepoResponseModelHostnameColumn is the column json value hostname
	RepoResponseModelHostnameColumn = "hostname"
	// RepoResponseModelIDColumn is the column json value id
	RepoResponseModelIDColumn = "id"
	// RepoResponseModelIntegrationIDColumn is the column json value integration_id
	RepoResponseModelIntegrationIDColumn = "integration_id"
	// RepoResponseModelLastExportDateColumn is the column json value last_export_date
	RepoResponseModelLastExportDateColumn = "last_export_date"
	// RepoResponseModelLastExportDateEpochColumn is the column json value epoch
	RepoResponseModelLastExportDateEpochColumn = "epoch"
	// RepoResponseModelLastExportDateOffsetColumn is the column json value offset
	RepoResponseModelLastExportDateOffsetColumn = "offset"
	// RepoResponseModelLastExportDateRfc3339Column is the column json value rfc3339
	RepoResponseModelLastExportDateRfc3339Column = "rfc3339"
	// RepoResponseModelMemoryColumn is the column json value memory
	RepoResponseModelMemoryColumn = "memory"
	// RepoResponseModelMessageColumn is the column json value message
	RepoResponseModelMessageColumn = "message"
	// RepoResponseModelNumCPUColumn is the column json value num_cpu
	RepoResponseModelNumCPUColumn = "num_cpu"
	// RepoResponseModelOSColumn is the column json value os
	RepoResponseModelOSColumn = "os"
	// RepoResponseModelRefIDColumn is the column json value ref_id
	RepoResponseModelRefIDColumn = "ref_id"
	// RepoResponseModelRefTypeColumn is the column json value ref_type
	RepoResponseModelRefTypeColumn = "ref_type"
	// RepoResponseModelReposColumn is the column json value repos
	RepoResponseModelReposColumn = "repos"
	// RepoResponseModelReposActiveColumn is the column json value active
	RepoResponseModelReposActiveColumn = "active"
	// RepoResponseModelReposCreatedDateColumn is the column json value created_date
	RepoResponseModelReposCreatedDateColumn = "created_date"
	// RepoResponseModelReposCreatedDateEpochColumn is the column json value epoch
	RepoResponseModelReposCreatedDateEpochColumn = "epoch"
	// RepoResponseModelReposCreatedDateOffsetColumn is the column json value offset
	RepoResponseModelReposCreatedDateOffsetColumn = "offset"
	// RepoResponseModelReposCreatedDateRfc3339Column is the column json value rfc3339
	RepoResponseModelReposCreatedDateRfc3339Column = "rfc3339"
	// RepoResponseModelReposDescriptionColumn is the column json value description
	RepoResponseModelReposDescriptionColumn = "description"
	// RepoResponseModelReposErrorColumn is the column json value error
	RepoResponseModelReposErrorColumn = "error"
	// RepoResponseModelReposLanguageColumn is the column json value language
	RepoResponseModelReposLanguageColumn = "language"
	// RepoResponseModelReposNameColumn is the column json value name
	RepoResponseModelReposNameColumn = "name"
	// RepoResponseModelReposRefIDColumn is the column json value ref_id
	RepoResponseModelReposRefIDColumn = "ref_id"
	// RepoResponseModelReposRefTypeColumn is the column json value ref_type
	RepoResponseModelReposRefTypeColumn = "ref_type"
	// RepoResponseModelReposWebhookPermissionColumn is the column json value webhook_permission
	RepoResponseModelReposWebhookPermissionColumn = "webhook_permission"
	// RepoResponseModelRequestIDColumn is the column json value request_id
	RepoResponseModelRequestIDColumn = "request_id"
	// RepoResponseModelSuccessColumn is the column json value success
	RepoResponseModelSuccessColumn = "success"
	// RepoResponseModelSystemIDColumn is the column json value system_id
	RepoResponseModelSystemIDColumn = "system_id"
	// RepoResponseModelTypeColumn is the column json value type
	RepoResponseModelTypeColumn = "type"
	// RepoResponseModelUptimeColumn is the column json value uptime
	RepoResponseModelUptimeColumn = "uptime"
	// RepoResponseModelUUIDColumn is the column json value uuid
	RepoResponseModelUUIDColumn = "uuid"
	// RepoResponseModelVersionColumn is the column json value version
	RepoResponseModelVersionColumn = "version"
)

// RepoResponseEventDate represents the object structure for event_date
type RepoResponseEventDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toRepoResponseEventDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *RepoResponseEventDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *RepoResponseEventDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toRepoResponseEventDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toRepoResponseEventDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toRepoResponseEventDateObject(o.Rfc3339, false),
	}
}

func (o *RepoResponseEventDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *RepoResponseEventDate) FromMap(kv map[string]interface{}) {

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

// RepoResponseLastExportDate represents the object structure for last_export_date
type RepoResponseLastExportDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toRepoResponseLastExportDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *RepoResponseLastExportDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *RepoResponseLastExportDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toRepoResponseLastExportDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toRepoResponseLastExportDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toRepoResponseLastExportDateObject(o.Rfc3339, false),
	}
}

func (o *RepoResponseLastExportDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *RepoResponseLastExportDate) FromMap(kv map[string]interface{}) {

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

// RepoResponseReposCreatedDate represents the object structure for created_date
type RepoResponseReposCreatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toRepoResponseReposCreatedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *RepoResponseReposCreatedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *RepoResponseReposCreatedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toRepoResponseReposCreatedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toRepoResponseReposCreatedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toRepoResponseReposCreatedDateObject(o.Rfc3339, false),
	}
}

func (o *RepoResponseReposCreatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *RepoResponseReposCreatedDate) FromMap(kv map[string]interface{}) {

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

// RepoResponseReposError is the enumeration type for error
type RepoResponseReposError int32

// UnmarshalBSONValue for unmarshaling value
func (v *RepoResponseReposError) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = RepoResponseReposError(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "NONE":
			*v = RepoResponseReposError(0)
		case "PERMISSIONS":
			*v = RepoResponseReposError(1)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v RepoResponseReposError) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "NONE":
		v = 0
	case "PERMISSIONS":
		v = 1
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v RepoResponseReposError) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("NONE")
	case 1:
		return json.Marshal("PERMISSIONS")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for ReposError
func (v RepoResponseReposError) String() string {
	switch int32(v) {
	case 0:
		return "NONE"
	case 1:
		return "PERMISSIONS"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *RepoResponseReposError) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case RepoResponseReposError:
		*v = val
	case int32:
		*v = RepoResponseReposError(int32(val))
	case int:
		*v = RepoResponseReposError(int32(val))
	case string:
		switch val {
		case "NONE":
			*v = RepoResponseReposError(0)
		case "PERMISSIONS":
			*v = RepoResponseReposError(1)
		}
	}
	return nil
}

const (
	// RepoResponseReposErrorNone is the enumeration value for none
	RepoResponseReposErrorNone RepoResponseReposError = 0
	// RepoResponseReposErrorPermissions is the enumeration value for permissions
	RepoResponseReposErrorPermissions RepoResponseReposError = 1
)

// RepoResponseRepos represents the object structure for repos
type RepoResponseRepos struct {
	// Active the status of the repo determined by an Admin
	Active bool `json:"active" codec:"active" bson:"active" yaml:"active" faker:"-"`
	// CreatedDate the creation date
	CreatedDate RepoResponseReposCreatedDate `json:"created_date" codec:"created_date" bson:"created_date" yaml:"created_date" faker:"-"`
	// Description the description of the repository
	Description string `json:"description" codec:"description" bson:"description" yaml:"description" faker:"sentence"`
	// Error reason why the repo is being set to inactive
	Error RepoResponseReposError `json:"error" codec:"error" bson:"error" yaml:"error" faker:"-"`
	// Language the programming language defined for the repository
	Language string `json:"language" codec:"language" bson:"language" yaml:"language" faker:"-"`
	// Name the name of the repository
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"repo"`
	// RefID the repo ID
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the record type
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// WebhookPermission if user has permission to install webhook for this repo
	WebhookPermission bool `json:"webhook_permission" codec:"webhook_permission" bson:"webhook_permission" yaml:"webhook_permission" faker:"-"`
}

func toRepoResponseReposObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *RepoResponseRepos:
		return v.ToMap()

	case RepoResponseReposCreatedDate:
		return v.ToMap()

	case RepoResponseReposError:
		return v.String()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *RepoResponseRepos) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Active the status of the repo determined by an Admin
		"active": toRepoResponseReposObject(o.Active, false),
		// CreatedDate the creation date
		"created_date": toRepoResponseReposObject(o.CreatedDate, false),
		// Description the description of the repository
		"description": toRepoResponseReposObject(o.Description, false),
		// Error reason why the repo is being set to inactive
		"error": toRepoResponseReposObject(o.Error, false),
		// Language the programming language defined for the repository
		"language": toRepoResponseReposObject(o.Language, false),
		// Name the name of the repository
		"name": toRepoResponseReposObject(o.Name, false),
		// RefID the repo ID
		"ref_id": toRepoResponseReposObject(o.RefID, false),
		// RefType the record type
		"ref_type": toRepoResponseReposObject(o.RefType, false),
		// WebhookPermission if user has permission to install webhook for this repo
		"webhook_permission": toRepoResponseReposObject(o.WebhookPermission, false),
	}
}

func (o *RepoResponseRepos) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *RepoResponseRepos) FromMap(kv map[string]interface{}) {

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

	if val, ok := kv["created_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.CreatedDate.FromMap(kv)
		} else if sv, ok := val.(RepoResponseReposCreatedDate); ok {
			// struct
			o.CreatedDate = sv
		} else if sp, ok := val.(*RepoResponseReposCreatedDate); ok {
			// struct pointer
			o.CreatedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.CreatedDate.Epoch = dt.Epoch
			o.CreatedDate.Rfc3339 = dt.Rfc3339
			o.CreatedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.CreatedDate.Epoch = dt.Epoch
			o.CreatedDate.Rfc3339 = dt.Rfc3339
			o.CreatedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.CreatedDate.Epoch = dt.Epoch
				o.CreatedDate.Rfc3339 = dt.Rfc3339
				o.CreatedDate.Offset = dt.Offset
			}
		}
	} else {
		o.CreatedDate.FromMap(map[string]interface{}{})
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
	if val, ok := kv["error"].(RepoResponseReposError); ok {
		o.Error = val
	} else {
		if em, ok := kv["error"].(map[string]interface{}); ok {

			ev := em["agent.error"].(string)
			switch ev {
			case "none", "NONE":
				o.Error = 0
			case "permissions", "PERMISSIONS":
				o.Error = 1
			}
		}
		if em, ok := kv["error"].(string); ok {
			switch em {
			case "none", "NONE":
				o.Error = 0
			case "permissions", "PERMISSIONS":
				o.Error = 1
			}
		}
	}
	if val, ok := kv["language"].(string); ok {
		o.Language = val
	} else {
		if val, ok := kv["language"]; ok {
			if val == nil {
				o.Language = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Language = fmt.Sprintf("%v", val)
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
	if val, ok := kv["webhook_permission"].(bool); ok {
		o.WebhookPermission = val
	} else {
		if val, ok := kv["webhook_permission"]; ok {
			if val == nil {
				o.WebhookPermission = false
			} else {
				o.WebhookPermission = number.ToBoolAny(val)
			}
		}
	}
	o.setDefaults(false)
}

// RepoResponseType is the enumeration type for type
type RepoResponseType int32

// UnmarshalBSONValue for unmarshaling value
func (v *RepoResponseType) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = RepoResponseType(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "ENROLL":
			*v = RepoResponseType(0)
		case "PING":
			*v = RepoResponseType(1)
		case "CRASH":
			*v = RepoResponseType(2)
		case "LOG":
			*v = RepoResponseType(3)
		case "INTEGRATION":
			*v = RepoResponseType(4)
		case "EXPORT":
			*v = RepoResponseType(5)
		case "PROJECT":
			*v = RepoResponseType(6)
		case "REPO":
			*v = RepoResponseType(7)
		case "USER":
			*v = RepoResponseType(8)
		case "CALENDAR":
			*v = RepoResponseType(9)
		case "UNINSTALL":
			*v = RepoResponseType(10)
		case "UPGRADE":
			*v = RepoResponseType(11)
		case "START":
			*v = RepoResponseType(12)
		case "STOP":
			*v = RepoResponseType(13)
		case "PAUSE":
			*v = RepoResponseType(14)
		case "RESUME":
			*v = RepoResponseType(15)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v RepoResponseType) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "ENROLL":
		v = 0
	case "PING":
		v = 1
	case "CRASH":
		v = 2
	case "LOG":
		v = 3
	case "INTEGRATION":
		v = 4
	case "EXPORT":
		v = 5
	case "PROJECT":
		v = 6
	case "REPO":
		v = 7
	case "USER":
		v = 8
	case "CALENDAR":
		v = 9
	case "UNINSTALL":
		v = 10
	case "UPGRADE":
		v = 11
	case "START":
		v = 12
	case "STOP":
		v = 13
	case "PAUSE":
		v = 14
	case "RESUME":
		v = 15
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v RepoResponseType) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("ENROLL")
	case 1:
		return json.Marshal("PING")
	case 2:
		return json.Marshal("CRASH")
	case 3:
		return json.Marshal("LOG")
	case 4:
		return json.Marshal("INTEGRATION")
	case 5:
		return json.Marshal("EXPORT")
	case 6:
		return json.Marshal("PROJECT")
	case 7:
		return json.Marshal("REPO")
	case 8:
		return json.Marshal("USER")
	case 9:
		return json.Marshal("CALENDAR")
	case 10:
		return json.Marshal("UNINSTALL")
	case 11:
		return json.Marshal("UPGRADE")
	case 12:
		return json.Marshal("START")
	case 13:
		return json.Marshal("STOP")
	case 14:
		return json.Marshal("PAUSE")
	case 15:
		return json.Marshal("RESUME")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for Type
func (v RepoResponseType) String() string {
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
		return "CALENDAR"
	case 10:
		return "UNINSTALL"
	case 11:
		return "UPGRADE"
	case 12:
		return "START"
	case 13:
		return "STOP"
	case 14:
		return "PAUSE"
	case 15:
		return "RESUME"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *RepoResponseType) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case RepoResponseType:
		*v = val
	case int32:
		*v = RepoResponseType(int32(val))
	case int:
		*v = RepoResponseType(int32(val))
	case string:
		switch val {
		case "ENROLL":
			*v = RepoResponseType(0)
		case "PING":
			*v = RepoResponseType(1)
		case "CRASH":
			*v = RepoResponseType(2)
		case "LOG":
			*v = RepoResponseType(3)
		case "INTEGRATION":
			*v = RepoResponseType(4)
		case "EXPORT":
			*v = RepoResponseType(5)
		case "PROJECT":
			*v = RepoResponseType(6)
		case "REPO":
			*v = RepoResponseType(7)
		case "USER":
			*v = RepoResponseType(8)
		case "CALENDAR":
			*v = RepoResponseType(9)
		case "UNINSTALL":
			*v = RepoResponseType(10)
		case "UPGRADE":
			*v = RepoResponseType(11)
		case "START":
			*v = RepoResponseType(12)
		case "STOP":
			*v = RepoResponseType(13)
		case "PAUSE":
			*v = RepoResponseType(14)
		case "RESUME":
			*v = RepoResponseType(15)
		}
	}
	return nil
}

const (
	// RepoResponseTypeEnroll is the enumeration value for enroll
	RepoResponseTypeEnroll RepoResponseType = 0
	// RepoResponseTypePing is the enumeration value for ping
	RepoResponseTypePing RepoResponseType = 1
	// RepoResponseTypeCrash is the enumeration value for crash
	RepoResponseTypeCrash RepoResponseType = 2
	// RepoResponseTypeLog is the enumeration value for log
	RepoResponseTypeLog RepoResponseType = 3
	// RepoResponseTypeIntegration is the enumeration value for integration
	RepoResponseTypeIntegration RepoResponseType = 4
	// RepoResponseTypeExport is the enumeration value for export
	RepoResponseTypeExport RepoResponseType = 5
	// RepoResponseTypeProject is the enumeration value for project
	RepoResponseTypeProject RepoResponseType = 6
	// RepoResponseTypeRepo is the enumeration value for repo
	RepoResponseTypeRepo RepoResponseType = 7
	// RepoResponseTypeUser is the enumeration value for user
	RepoResponseTypeUser RepoResponseType = 8
	// RepoResponseTypeCalendar is the enumeration value for calendar
	RepoResponseTypeCalendar RepoResponseType = 9
	// RepoResponseTypeUninstall is the enumeration value for uninstall
	RepoResponseTypeUninstall RepoResponseType = 10
	// RepoResponseTypeUpgrade is the enumeration value for upgrade
	RepoResponseTypeUpgrade RepoResponseType = 11
	// RepoResponseTypeStart is the enumeration value for start
	RepoResponseTypeStart RepoResponseType = 12
	// RepoResponseTypeStop is the enumeration value for stop
	RepoResponseTypeStop RepoResponseType = 13
	// RepoResponseTypePause is the enumeration value for pause
	RepoResponseTypePause RepoResponseType = 14
	// RepoResponseTypeResume is the enumeration value for resume
	RepoResponseTypeResume RepoResponseType = 15
)

// RepoResponse an agent response to an action request adding repo(s)
type RepoResponse struct {
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
	EventDate RepoResponseEventDate `json:"event_date" codec:"event_date" bson:"event_date" yaml:"event_date" faker:"-"`
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
	LastExportDate RepoResponseLastExportDate `json:"last_export_date" codec:"last_export_date" bson:"last_export_date" yaml:"last_export_date" faker:"-"`
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
	// Repos the repos exported
	Repos []RepoResponseRepos `json:"repos" codec:"repos" bson:"repos" yaml:"repos" faker:"-"`
	// RequestID the request id that this response is correlated to
	RequestID string `json:"request_id" codec:"request_id" bson:"request_id" yaml:"request_id" faker:"-"`
	// Success if the response was successful
	Success bool `json:"success" codec:"success" bson:"success" yaml:"success" faker:"-"`
	// SystemID system unique device ID
	SystemID string `json:"system_id" codec:"system_id" bson:"system_id" yaml:"system_id" faker:"-"`
	// Type the type of event
	Type RepoResponseType `json:"type" codec:"type" bson:"type" yaml:"type" faker:"-"`
	// Uptime the uptime in milliseconds since the agent started
	Uptime int64 `json:"uptime" codec:"uptime" bson:"uptime" yaml:"uptime" faker:"-"`
	// UUID the agent unique identifier
	UUID string `json:"uuid" codec:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
	// Version the agent version
	Version string `json:"version" codec:"version" bson:"version" yaml:"version" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*RepoResponse)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*RepoResponse)(nil)

func toRepoResponseObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *RepoResponse:
		return v.ToMap()

	case RepoResponseEventDate:
		return v.ToMap()

	case RepoResponseLastExportDate:
		return v.ToMap()

	case []RepoResponseRepos:
		arr := make([]interface{}, 0)
		for _, i := range v {
			arr = append(arr, i.ToMap())
		}
		return arr

	case RepoResponseType:
		return v.String()

	default:
		return o
	}
}

// String returns a string representation of RepoResponse
func (o *RepoResponse) String() string {
	return fmt.Sprintf("agent.RepoResponse<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *RepoResponse) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *RepoResponse) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *RepoResponse) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *RepoResponse) GetModelName() datamodel.ModelNameType {
	return RepoResponseModelName
}

// NewRepoResponseID provides a template for generating an ID field for RepoResponse
func NewRepoResponseID(customerID string, refType string, refID string) string {
	return hash.Values("RepoResponse", customerID, refType, refID)
}

func (o *RepoResponse) setDefaults(frommap bool) {
	if o.Repos == nil {
		o.Repos = make([]RepoResponseRepos, 0)
	}

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("RepoResponse", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *RepoResponse) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *RepoResponse) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *RepoResponse) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *RepoResponse) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *RepoResponse) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *RepoResponse) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *RepoResponse) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *RepoResponse) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *RepoResponse) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = RepoResponseModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *RepoResponse) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *RepoResponse) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of RepoResponse
func (o *RepoResponse) Clone() datamodel.Model {
	c := new(RepoResponse)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *RepoResponse) Anon() datamodel.Model {
	c := new(RepoResponse)
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
func (o *RepoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *RepoResponse) UnmarshalJSON(data []byte) error {
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
func (o *RepoResponse) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *RepoResponse) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two RepoResponse objects are equal
func (o *RepoResponse) IsEqual(other *RepoResponse) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *RepoResponse) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"architecture":     toRepoResponseObject(o.Architecture, false),
		"customer_id":      toRepoResponseObject(o.CustomerID, false),
		"data":             toRepoResponseObject(o.Data, true),
		"distro":           toRepoResponseObject(o.Distro, false),
		"error":            toRepoResponseObject(o.Error, true),
		"event_date":       toRepoResponseObject(o.EventDate, false),
		"free_space":       toRepoResponseObject(o.FreeSpace, false),
		"go_version":       toRepoResponseObject(o.GoVersion, false),
		"hostname":         toRepoResponseObject(o.Hostname, false),
		"id":               toRepoResponseObject(o.ID, false),
		"integration_id":   toRepoResponseObject(o.IntegrationID, false),
		"last_export_date": toRepoResponseObject(o.LastExportDate, false),
		"memory":           toRepoResponseObject(o.Memory, false),
		"message":          toRepoResponseObject(o.Message, false),
		"num_cpu":          toRepoResponseObject(o.NumCPU, false),
		"os":               toRepoResponseObject(o.OS, false),
		"ref_id":           toRepoResponseObject(o.RefID, false),
		"ref_type":         toRepoResponseObject(o.RefType, false),
		"repos":            toRepoResponseObject(o.Repos, false),
		"request_id":       toRepoResponseObject(o.RequestID, false),
		"success":          toRepoResponseObject(o.Success, false),
		"system_id":        toRepoResponseObject(o.SystemID, false),

		"type":     o.Type.String(),
		"uptime":   toRepoResponseObject(o.Uptime, false),
		"uuid":     toRepoResponseObject(o.UUID, false),
		"version":  toRepoResponseObject(o.Version, false),
		"hashcode": toRepoResponseObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *RepoResponse) FromMap(kv map[string]interface{}) {

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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
		} else if sv, ok := val.(RepoResponseEventDate); ok {
			// struct
			o.EventDate = sv
		} else if sp, ok := val.(*RepoResponseEventDate); ok {
			// struct pointer
			o.EventDate = *sp
		}
	} else {
		o.EventDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["free_space"].(int64); ok {
		o.FreeSpace = val
	} else {
		if val, ok := kv["free_space"]; ok {
			if val == nil {
				o.FreeSpace = 0
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
	if val, ok := kv["integration_id"].(string); ok {
		o.IntegrationID = val
	} else {
		if val, ok := kv["integration_id"]; ok {
			if val == nil {
				o.IntegrationID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.IntegrationID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["last_export_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastExportDate.FromMap(kv)
		} else if sv, ok := val.(RepoResponseLastExportDate); ok {
			// struct
			o.LastExportDate = sv
		} else if sp, ok := val.(*RepoResponseLastExportDate); ok {
			// struct pointer
			o.LastExportDate = *sp
		}
	} else {
		o.LastExportDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["memory"].(int64); ok {
		o.Memory = val
	} else {
		if val, ok := kv["memory"]; ok {
			if val == nil {
				o.Memory = 0
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
				o.NumCPU = 0
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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

	if o == nil {

		o.Repos = make([]RepoResponseRepos, 0)

	}
	if val, ok := kv["repos"]; ok {
		if sv, ok := val.([]RepoResponseRepos); ok {
			o.Repos = sv
		} else if sp, ok := val.([]*RepoResponseRepos); ok {
			o.Repos = o.Repos[:0]
			for _, e := range sp {
				o.Repos = append(o.Repos, *e)
			}
		} else if a, ok := val.(primitive.A); ok {
			for _, ae := range a {
				if av, ok := ae.(RepoResponseRepos); ok {
					o.Repos = append(o.Repos, av)
				} else if av, ok := ae.(primitive.M); ok {
					var fm RepoResponseRepos
					fm.FromMap(av)
					o.Repos = append(o.Repos, fm)
				} else {
					b, _ := json.Marshal(ae)
					bkv := make(map[string]interface{})
					json.Unmarshal(b, &bkv)
					var av RepoResponseRepos
					av.FromMap(bkv)
					o.Repos = append(o.Repos, av)
				}
			}
		} else if arr, ok := val.([]interface{}); ok {
			for _, item := range arr {
				if r, ok := item.(RepoResponseRepos); ok {
					o.Repos = append(o.Repos, r)
				} else if r, ok := item.(map[string]interface{}); ok {
					var fm RepoResponseRepos
					fm.FromMap(r)
					o.Repos = append(o.Repos, fm)
				} else if r, ok := item.(primitive.M); ok {
					fm := RepoResponseRepos{}
					fm.FromMap(r)
					o.Repos = append(o.Repos, fm)
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
							var fm RepoResponseRepos
							fm.FromMap(m[0].Interface().(map[string]interface{}))
							o.Repos = append(o.Repos, fm)
						}
					}
				}
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
				o.Success = false
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.SystemID = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["type"].(RepoResponseType); ok {
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
			case "calendar", "CALENDAR":
				o.Type = 9
			case "uninstall", "UNINSTALL":
				o.Type = 10
			case "upgrade", "UPGRADE":
				o.Type = 11
			case "start", "START":
				o.Type = 12
			case "stop", "STOP":
				o.Type = 13
			case "pause", "PAUSE":
				o.Type = 14
			case "resume", "RESUME":
				o.Type = 15
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
			case "calendar", "CALENDAR":
				o.Type = 9
			case "uninstall", "UNINSTALL":
				o.Type = 10
			case "upgrade", "UPGRADE":
				o.Type = 11
			case "start", "START":
				o.Type = 12
			case "stop", "STOP":
				o.Type = 13
			case "pause", "PAUSE":
				o.Type = 14
			case "resume", "RESUME":
				o.Type = 15
			}
		}
	}
	if val, ok := kv["uptime"].(int64); ok {
		o.Uptime = val
	} else {
		if val, ok := kv["uptime"]; ok {
			if val == nil {
				o.Uptime = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Uptime = number.ToInt64Any(val)
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Version = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *RepoResponse) Hash() string {
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
	args = append(args, o.Repos)
	args = append(args, o.RequestID)
	args = append(args, o.Success)
	args = append(args, o.SystemID)
	args = append(args, o.Type)
	args = append(args, o.Uptime)
	args = append(args, o.UUID)
	args = append(args, o.Version)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}
