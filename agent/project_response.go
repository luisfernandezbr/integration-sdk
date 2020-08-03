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
	"github.com/pinpt/go-common/v10/datetime"
	"github.com/pinpt/go-common/v10/hash"
	pjson "github.com/pinpt/go-common/v10/json"
	"github.com/pinpt/go-common/v10/number"
	pstrings "github.com/pinpt/go-common/v10/strings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (

	// ProjectResponseTable is the default table name
	ProjectResponseTable datamodel.ModelNameType = "agent_projectresponse"

	// ProjectResponseModelName is the model name
	ProjectResponseModelName datamodel.ModelNameType = "agent.ProjectResponse"
)

const (
	// ProjectResponseModelArchitectureColumn is the column json value architecture
	ProjectResponseModelArchitectureColumn = "architecture"
	// ProjectResponseModelCustomerIDColumn is the column json value customer_id
	ProjectResponseModelCustomerIDColumn = "customer_id"
	// ProjectResponseModelDataColumn is the column json value data
	ProjectResponseModelDataColumn = "data"
	// ProjectResponseModelDistroColumn is the column json value distro
	ProjectResponseModelDistroColumn = "distro"
	// ProjectResponseModelErrorColumn is the column json value error
	ProjectResponseModelErrorColumn = "error"
	// ProjectResponseModelEventDateColumn is the column json value event_date
	ProjectResponseModelEventDateColumn = "event_date"
	// ProjectResponseModelEventDateEpochColumn is the column json value epoch
	ProjectResponseModelEventDateEpochColumn = "epoch"
	// ProjectResponseModelEventDateOffsetColumn is the column json value offset
	ProjectResponseModelEventDateOffsetColumn = "offset"
	// ProjectResponseModelEventDateRfc3339Column is the column json value rfc3339
	ProjectResponseModelEventDateRfc3339Column = "rfc3339"
	// ProjectResponseModelFreeSpaceColumn is the column json value free_space
	ProjectResponseModelFreeSpaceColumn = "free_space"
	// ProjectResponseModelGoVersionColumn is the column json value go_version
	ProjectResponseModelGoVersionColumn = "go_version"
	// ProjectResponseModelHostnameColumn is the column json value hostname
	ProjectResponseModelHostnameColumn = "hostname"
	// ProjectResponseModelIDColumn is the column json value id
	ProjectResponseModelIDColumn = "id"
	// ProjectResponseModelIntegrationIDColumn is the column json value integration_id
	ProjectResponseModelIntegrationIDColumn = "integration_id"
	// ProjectResponseModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	ProjectResponseModelIntegrationInstanceIDColumn = "integration_instance_id"
	// ProjectResponseModelLastExportDateColumn is the column json value last_export_date
	ProjectResponseModelLastExportDateColumn = "last_export_date"
	// ProjectResponseModelLastExportDateEpochColumn is the column json value epoch
	ProjectResponseModelLastExportDateEpochColumn = "epoch"
	// ProjectResponseModelLastExportDateOffsetColumn is the column json value offset
	ProjectResponseModelLastExportDateOffsetColumn = "offset"
	// ProjectResponseModelLastExportDateRfc3339Column is the column json value rfc3339
	ProjectResponseModelLastExportDateRfc3339Column = "rfc3339"
	// ProjectResponseModelMemoryColumn is the column json value memory
	ProjectResponseModelMemoryColumn = "memory"
	// ProjectResponseModelMessageColumn is the column json value message
	ProjectResponseModelMessageColumn = "message"
	// ProjectResponseModelNumCPUColumn is the column json value num_cpu
	ProjectResponseModelNumCPUColumn = "num_cpu"
	// ProjectResponseModelOSColumn is the column json value os
	ProjectResponseModelOSColumn = "os"
	// ProjectResponseModelProjectsColumn is the column json value projects
	ProjectResponseModelProjectsColumn = "projects"
	// ProjectResponseModelProjectsActiveColumn is the column json value active
	ProjectResponseModelProjectsActiveColumn = "active"
	// ProjectResponseModelProjectsCategoryColumn is the column json value category
	ProjectResponseModelProjectsCategoryColumn = "category"
	// ProjectResponseModelProjectsDescriptionColumn is the column json value description
	ProjectResponseModelProjectsDescriptionColumn = "description"
	// ProjectResponseModelProjectsErrorColumn is the column json value error
	ProjectResponseModelProjectsErrorColumn = "error"
	// ProjectResponseModelProjectsIdentifierColumn is the column json value identifier
	ProjectResponseModelProjectsIdentifierColumn = "identifier"
	// ProjectResponseModelProjectsNameColumn is the column json value name
	ProjectResponseModelProjectsNameColumn = "name"
	// ProjectResponseModelProjectsRefIDColumn is the column json value ref_id
	ProjectResponseModelProjectsRefIDColumn = "ref_id"
	// ProjectResponseModelProjectsRefTypeColumn is the column json value ref_type
	ProjectResponseModelProjectsRefTypeColumn = "ref_type"
	// ProjectResponseModelProjectsURLColumn is the column json value url
	ProjectResponseModelProjectsURLColumn = "url"
	// ProjectResponseModelProjectsWebhookPermissionColumn is the column json value webhook_permission
	ProjectResponseModelProjectsWebhookPermissionColumn = "webhook_permission"
	// ProjectResponseModelRefIDColumn is the column json value ref_id
	ProjectResponseModelRefIDColumn = "ref_id"
	// ProjectResponseModelRefTypeColumn is the column json value ref_type
	ProjectResponseModelRefTypeColumn = "ref_type"
	// ProjectResponseModelRequestIDColumn is the column json value request_id
	ProjectResponseModelRequestIDColumn = "request_id"
	// ProjectResponseModelSuccessColumn is the column json value success
	ProjectResponseModelSuccessColumn = "success"
	// ProjectResponseModelSystemIDColumn is the column json value system_id
	ProjectResponseModelSystemIDColumn = "system_id"
	// ProjectResponseModelTypeColumn is the column json value type
	ProjectResponseModelTypeColumn = "type"
	// ProjectResponseModelUptimeColumn is the column json value uptime
	ProjectResponseModelUptimeColumn = "uptime"
	// ProjectResponseModelUUIDColumn is the column json value uuid
	ProjectResponseModelUUIDColumn = "uuid"
	// ProjectResponseModelVersionColumn is the column json value version
	ProjectResponseModelVersionColumn = "version"
)

// ProjectResponseEventDate represents the object structure for event_date
type ProjectResponseEventDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toProjectResponseEventDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ProjectResponseEventDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *ProjectResponseEventDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toProjectResponseEventDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toProjectResponseEventDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toProjectResponseEventDateObject(o.Rfc3339, false),
	}
}

func (o *ProjectResponseEventDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ProjectResponseEventDate) FromMap(kv map[string]interface{}) {

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

// ProjectResponseLastExportDate represents the object structure for last_export_date
type ProjectResponseLastExportDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toProjectResponseLastExportDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ProjectResponseLastExportDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *ProjectResponseLastExportDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toProjectResponseLastExportDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toProjectResponseLastExportDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toProjectResponseLastExportDateObject(o.Rfc3339, false),
	}
}

func (o *ProjectResponseLastExportDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ProjectResponseLastExportDate) FromMap(kv map[string]interface{}) {

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

// ProjectResponseProjectsError is the enumeration type for error
type ProjectResponseProjectsError int32

// toProjectResponseProjectsErrorPointer is the enumeration pointer type for error
func toProjectResponseProjectsErrorPointer(v int32) *ProjectResponseProjectsError {
	nv := ProjectResponseProjectsError(v)
	return &nv
}

// toProjectResponseProjectsErrorEnum is the enumeration pointer wrapper for error
func toProjectResponseProjectsErrorEnum(v *ProjectResponseProjectsError) string {
	if v == nil {
		return toProjectResponseProjectsErrorPointer(0).String()
	}
	return v.String()
}

// UnmarshalBSONValue for unmarshaling value
func (v *ProjectResponseProjectsError) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = ProjectResponseProjectsError(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "NONE":
			*v = ProjectResponseProjectsError(0)
		case "PERMISSIONS":
			*v = ProjectResponseProjectsError(1)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v *ProjectResponseProjectsError) UnmarshalJSON(buf []byte) error {
	var val string
	if err := json.Unmarshal(buf, &val); err != nil {
		return err
	}
	switch val {
	case "NONE":
		*v = 0
	case "PERMISSIONS":
		*v = 1
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v ProjectResponseProjectsError) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("NONE")
	case 1:
		return json.Marshal("PERMISSIONS")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for ProjectsError
func (v ProjectResponseProjectsError) String() string {
	switch int32(v) {
	case 0:
		return "NONE"
	case 1:
		return "PERMISSIONS"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *ProjectResponseProjectsError) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case ProjectResponseProjectsError:
		*v = val
	case int32:
		*v = ProjectResponseProjectsError(int32(val))
	case int:
		*v = ProjectResponseProjectsError(int32(val))
	case string:
		switch val {
		case "NONE":
			*v = ProjectResponseProjectsError(0)
		case "PERMISSIONS":
			*v = ProjectResponseProjectsError(1)
		}
	}
	return nil
}

const (
	// ProjectResponseProjectsErrorNone is the enumeration value for none
	ProjectResponseProjectsErrorNone ProjectResponseProjectsError = 0
	// ProjectResponseProjectsErrorPermissions is the enumeration value for permissions
	ProjectResponseProjectsErrorPermissions ProjectResponseProjectsError = 1
)

// ProjectResponseProjects represents the object structure for projects
type ProjectResponseProjects struct {
	// Active the status of the project
	Active bool `json:"active" codec:"active" bson:"active" yaml:"active" faker:"-"`
	// Category the project category
	Category *string `json:"category,omitempty" codec:"category,omitempty" bson:"category" yaml:"category,omitempty" faker:"-"`
	// Description the description of the project
	Description *string `json:"description,omitempty" codec:"description,omitempty" bson:"description" yaml:"description,omitempty" faker:"-"`
	// Error reason why the project is being set to inactive
	Error ProjectResponseProjectsError `json:"error" codec:"error" bson:"error" yaml:"error" faker:"-"`
	// Identifier the common identifier for the project
	Identifier string `json:"identifier" codec:"identifier" bson:"identifier" yaml:"identifier" faker:"-"`
	// Name the name of the project
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
	// RefID the id of the project
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the record type
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// URL the url to the project home page
	URL string `json:"url" codec:"url" bson:"url" yaml:"url" faker:"-"`
	// WebhookPermission if user has permission to install webhook for this project
	WebhookPermission bool `json:"webhook_permission" codec:"webhook_permission" bson:"webhook_permission" yaml:"webhook_permission" faker:"-"`
}

func toProjectResponseProjectsObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ProjectResponseProjects:
		return v.ToMap()

	case ProjectResponseProjectsError:
		return v.String()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *ProjectResponseProjects) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Active the status of the project
		"active": toProjectResponseProjectsObject(o.Active, false),
		// Category the project category
		"category": toProjectResponseProjectsObject(o.Category, true),
		// Description the description of the project
		"description": toProjectResponseProjectsObject(o.Description, true),
		// Error reason why the project is being set to inactive
		"error": toProjectResponseProjectsObject(o.Error, false),
		// Identifier the common identifier for the project
		"identifier": toProjectResponseProjectsObject(o.Identifier, false),
		// Name the name of the project
		"name": toProjectResponseProjectsObject(o.Name, false),
		// RefID the id of the project
		"ref_id": toProjectResponseProjectsObject(o.RefID, false),
		// RefType the record type
		"ref_type": toProjectResponseProjectsObject(o.RefType, false),
		// URL the url to the project home page
		"url": toProjectResponseProjectsObject(o.URL, false),
		// WebhookPermission if user has permission to install webhook for this project
		"webhook_permission": toProjectResponseProjectsObject(o.WebhookPermission, false),
	}
}

func (o *ProjectResponseProjects) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ProjectResponseProjects) FromMap(kv map[string]interface{}) {

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
	if val, ok := kv["category"].(*string); ok {
		o.Category = val
	} else if val, ok := kv["category"].(string); ok {
		o.Category = &val
	} else {
		if val, ok := kv["category"]; ok {
			if val == nil {
				o.Category = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Category = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["description"].(*string); ok {
		o.Description = val
	} else if val, ok := kv["description"].(string); ok {
		o.Description = &val
	} else {
		if val, ok := kv["description"]; ok {
			if val == nil {
				o.Description = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Description = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["error"].(ProjectResponseProjectsError); ok {
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
	if val, ok := kv["identifier"].(string); ok {
		o.Identifier = val
	} else {
		if val, ok := kv["identifier"]; ok {
			if val == nil {
				o.Identifier = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Identifier = fmt.Sprintf("%v", val)
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

// ProjectResponseType is the enumeration type for type
type ProjectResponseType int32

// toProjectResponseTypePointer is the enumeration pointer type for type
func toProjectResponseTypePointer(v int32) *ProjectResponseType {
	nv := ProjectResponseType(v)
	return &nv
}

// toProjectResponseTypeEnum is the enumeration pointer wrapper for type
func toProjectResponseTypeEnum(v *ProjectResponseType) string {
	if v == nil {
		return toProjectResponseTypePointer(0).String()
	}
	return v.String()
}

// UnmarshalBSONValue for unmarshaling value
func (v *ProjectResponseType) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = ProjectResponseType(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "ENROLL":
			*v = ProjectResponseType(0)
		case "PING":
			*v = ProjectResponseType(1)
		case "CRASH":
			*v = ProjectResponseType(2)
		case "LOG":
			*v = ProjectResponseType(3)
		case "INTEGRATION":
			*v = ProjectResponseType(4)
		case "EXPORT":
			*v = ProjectResponseType(5)
		case "PROJECT":
			*v = ProjectResponseType(6)
		case "REPO":
			*v = ProjectResponseType(7)
		case "USER":
			*v = ProjectResponseType(8)
		case "CALENDAR":
			*v = ProjectResponseType(9)
		case "UNINSTALL":
			*v = ProjectResponseType(10)
		case "UPGRADE":
			*v = ProjectResponseType(11)
		case "START":
			*v = ProjectResponseType(12)
		case "STOP":
			*v = ProjectResponseType(13)
		case "PAUSE":
			*v = ProjectResponseType(14)
		case "RESUME":
			*v = ProjectResponseType(15)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v *ProjectResponseType) UnmarshalJSON(buf []byte) error {
	var val string
	if err := json.Unmarshal(buf, &val); err != nil {
		return err
	}
	switch val {
	case "ENROLL":
		*v = 0
	case "PING":
		*v = 1
	case "CRASH":
		*v = 2
	case "LOG":
		*v = 3
	case "INTEGRATION":
		*v = 4
	case "EXPORT":
		*v = 5
	case "PROJECT":
		*v = 6
	case "REPO":
		*v = 7
	case "USER":
		*v = 8
	case "CALENDAR":
		*v = 9
	case "UNINSTALL":
		*v = 10
	case "UPGRADE":
		*v = 11
	case "START":
		*v = 12
	case "STOP":
		*v = 13
	case "PAUSE":
		*v = 14
	case "RESUME":
		*v = 15
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v ProjectResponseType) MarshalJSON() ([]byte, error) {
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
func (v ProjectResponseType) String() string {
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
func (v *ProjectResponseType) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case ProjectResponseType:
		*v = val
	case int32:
		*v = ProjectResponseType(int32(val))
	case int:
		*v = ProjectResponseType(int32(val))
	case string:
		switch val {
		case "ENROLL":
			*v = ProjectResponseType(0)
		case "PING":
			*v = ProjectResponseType(1)
		case "CRASH":
			*v = ProjectResponseType(2)
		case "LOG":
			*v = ProjectResponseType(3)
		case "INTEGRATION":
			*v = ProjectResponseType(4)
		case "EXPORT":
			*v = ProjectResponseType(5)
		case "PROJECT":
			*v = ProjectResponseType(6)
		case "REPO":
			*v = ProjectResponseType(7)
		case "USER":
			*v = ProjectResponseType(8)
		case "CALENDAR":
			*v = ProjectResponseType(9)
		case "UNINSTALL":
			*v = ProjectResponseType(10)
		case "UPGRADE":
			*v = ProjectResponseType(11)
		case "START":
			*v = ProjectResponseType(12)
		case "STOP":
			*v = ProjectResponseType(13)
		case "PAUSE":
			*v = ProjectResponseType(14)
		case "RESUME":
			*v = ProjectResponseType(15)
		}
	}
	return nil
}

const (
	// ProjectResponseTypeEnroll is the enumeration value for enroll
	ProjectResponseTypeEnroll ProjectResponseType = 0
	// ProjectResponseTypePing is the enumeration value for ping
	ProjectResponseTypePing ProjectResponseType = 1
	// ProjectResponseTypeCrash is the enumeration value for crash
	ProjectResponseTypeCrash ProjectResponseType = 2
	// ProjectResponseTypeLog is the enumeration value for log
	ProjectResponseTypeLog ProjectResponseType = 3
	// ProjectResponseTypeIntegration is the enumeration value for integration
	ProjectResponseTypeIntegration ProjectResponseType = 4
	// ProjectResponseTypeExport is the enumeration value for export
	ProjectResponseTypeExport ProjectResponseType = 5
	// ProjectResponseTypeProject is the enumeration value for project
	ProjectResponseTypeProject ProjectResponseType = 6
	// ProjectResponseTypeRepo is the enumeration value for repo
	ProjectResponseTypeRepo ProjectResponseType = 7
	// ProjectResponseTypeUser is the enumeration value for user
	ProjectResponseTypeUser ProjectResponseType = 8
	// ProjectResponseTypeCalendar is the enumeration value for calendar
	ProjectResponseTypeCalendar ProjectResponseType = 9
	// ProjectResponseTypeUninstall is the enumeration value for uninstall
	ProjectResponseTypeUninstall ProjectResponseType = 10
	// ProjectResponseTypeUpgrade is the enumeration value for upgrade
	ProjectResponseTypeUpgrade ProjectResponseType = 11
	// ProjectResponseTypeStart is the enumeration value for start
	ProjectResponseTypeStart ProjectResponseType = 12
	// ProjectResponseTypeStop is the enumeration value for stop
	ProjectResponseTypeStop ProjectResponseType = 13
	// ProjectResponseTypePause is the enumeration value for pause
	ProjectResponseTypePause ProjectResponseType = 14
	// ProjectResponseTypeResume is the enumeration value for resume
	ProjectResponseTypeResume ProjectResponseType = 15
)

// ProjectResponse an agent response to an action request adding project(s)
type ProjectResponse struct {
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
	EventDate ProjectResponseEventDate `json:"event_date" codec:"event_date" bson:"event_date" yaml:"event_date" faker:"-"`
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
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// LastExportDate the last export date
	LastExportDate ProjectResponseLastExportDate `json:"last_export_date" codec:"last_export_date" bson:"last_export_date" yaml:"last_export_date" faker:"-"`
	// Memory the amount of memory in bytes for the agent machine
	Memory int64 `json:"memory" codec:"memory" bson:"memory" yaml:"memory" faker:"-"`
	// Message a message related to this event
	Message string `json:"message" codec:"message" bson:"message" yaml:"message" faker:"-"`
	// NumCPU the number of CPU the agent is running
	NumCPU int64 `json:"num_cpu" codec:"num_cpu" bson:"num_cpu" yaml:"num_cpu" faker:"-"`
	// OS the agent operating system
	OS string `json:"os" codec:"os" bson:"os" yaml:"os" faker:"-"`
	// Projects the projects exported
	Projects []ProjectResponseProjects `json:"projects" codec:"projects" bson:"projects" yaml:"projects" faker:"-"`
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
	Type ProjectResponseType `json:"type" codec:"type" bson:"type" yaml:"type" faker:"-"`
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
var _ datamodel.Model = (*ProjectResponse)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*ProjectResponse)(nil)

func toProjectResponseObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ProjectResponse:
		return v.ToMap()

	case ProjectResponseEventDate:
		return v.ToMap()

	case ProjectResponseLastExportDate:
		return v.ToMap()

	case []ProjectResponseProjects:
		arr := make([]interface{}, 0)
		for _, i := range v {
			arr = append(arr, i.ToMap())
		}
		return arr

	case ProjectResponseType:
		return v.String()

	default:
		return o
	}
}

// String returns a string representation of ProjectResponse
func (o *ProjectResponse) String() string {
	return fmt.Sprintf("agent.ProjectResponse<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *ProjectResponse) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *ProjectResponse) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *ProjectResponse) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *ProjectResponse) GetModelName() datamodel.ModelNameType {
	return ProjectResponseModelName
}

// NewProjectResponseID provides a template for generating an ID field for ProjectResponse
func NewProjectResponseID(customerID string, refType string, refID string) string {
	return hash.Values("ProjectResponse", customerID, refType, refID)
}

func (o *ProjectResponse) setDefaults(frommap bool) {
	if o.Projects == nil {
		o.Projects = make([]ProjectResponseProjects, 0)
	}

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("ProjectResponse", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *ProjectResponse) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *ProjectResponse) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *ProjectResponse) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *ProjectResponse) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *ProjectResponse) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *ProjectResponse) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *ProjectResponse) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *ProjectResponse) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *ProjectResponse) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = ProjectResponseModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *ProjectResponse) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *ProjectResponse) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of ProjectResponse
func (o *ProjectResponse) Clone() datamodel.Model {
	c := new(ProjectResponse)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *ProjectResponse) Anon() datamodel.Model {
	c := new(ProjectResponse)
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
func (o *ProjectResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *ProjectResponse) UnmarshalJSON(data []byte) error {
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
func (o *ProjectResponse) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *ProjectResponse) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two ProjectResponse objects are equal
func (o *ProjectResponse) IsEqual(other *ProjectResponse) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *ProjectResponse) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"architecture":            toProjectResponseObject(o.Architecture, false),
		"customer_id":             toProjectResponseObject(o.CustomerID, false),
		"data":                    toProjectResponseObject(o.Data, true),
		"distro":                  toProjectResponseObject(o.Distro, false),
		"error":                   toProjectResponseObject(o.Error, true),
		"event_date":              toProjectResponseObject(o.EventDate, false),
		"free_space":              toProjectResponseObject(o.FreeSpace, false),
		"go_version":              toProjectResponseObject(o.GoVersion, false),
		"hostname":                toProjectResponseObject(o.Hostname, false),
		"id":                      toProjectResponseObject(o.ID, false),
		"integration_id":          toProjectResponseObject(o.IntegrationID, false),
		"integration_instance_id": toProjectResponseObject(o.IntegrationInstanceID, true),
		"last_export_date":        toProjectResponseObject(o.LastExportDate, false),
		"memory":                  toProjectResponseObject(o.Memory, false),
		"message":                 toProjectResponseObject(o.Message, false),
		"num_cpu":                 toProjectResponseObject(o.NumCPU, false),
		"os":                      toProjectResponseObject(o.OS, false),
		"projects":                toProjectResponseObject(o.Projects, false),
		"ref_id":                  toProjectResponseObject(o.RefID, false),
		"ref_type":                toProjectResponseObject(o.RefType, false),
		"request_id":              toProjectResponseObject(o.RequestID, false),
		"success":                 toProjectResponseObject(o.Success, false),
		"system_id":               toProjectResponseObject(o.SystemID, false),

		"type":     o.Type.String(),
		"uptime":   toProjectResponseObject(o.Uptime, false),
		"uuid":     toProjectResponseObject(o.UUID, false),
		"version":  toProjectResponseObject(o.Version, false),
		"hashcode": toProjectResponseObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *ProjectResponse) FromMap(kv map[string]interface{}) {

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
		} else if sv, ok := val.(ProjectResponseEventDate); ok {
			// struct
			o.EventDate = sv
		} else if sp, ok := val.(*ProjectResponseEventDate); ok {
			// struct pointer
			o.EventDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.EventDate.Epoch = dt.Epoch
			o.EventDate.Rfc3339 = dt.Rfc3339
			o.EventDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
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

	if val, ok := kv["last_export_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastExportDate.FromMap(kv)
		} else if sv, ok := val.(ProjectResponseLastExportDate); ok {
			// struct
			o.LastExportDate = sv
		} else if sp, ok := val.(*ProjectResponseLastExportDate); ok {
			// struct pointer
			o.LastExportDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.LastExportDate.Epoch = dt.Epoch
			o.LastExportDate.Rfc3339 = dt.Rfc3339
			o.LastExportDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
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

	if o == nil {

		o.Projects = make([]ProjectResponseProjects, 0)

	}
	if val, ok := kv["projects"]; ok {
		if sv, ok := val.([]ProjectResponseProjects); ok {
			o.Projects = sv
		} else if sp, ok := val.([]*ProjectResponseProjects); ok {
			o.Projects = o.Projects[:0]
			for _, e := range sp {
				o.Projects = append(o.Projects, *e)
			}
		} else if a, ok := val.(primitive.A); ok {
			for _, ae := range a {
				if av, ok := ae.(ProjectResponseProjects); ok {
					o.Projects = append(o.Projects, av)
				} else if av, ok := ae.(primitive.M); ok {
					var fm ProjectResponseProjects
					fm.FromMap(av)
					o.Projects = append(o.Projects, fm)
				} else {
					b, _ := json.Marshal(ae)
					bkv := make(map[string]interface{})
					json.Unmarshal(b, &bkv)
					var av ProjectResponseProjects
					av.FromMap(bkv)
					o.Projects = append(o.Projects, av)
				}
			}
		} else if arr, ok := val.([]interface{}); ok {
			for _, item := range arr {
				if r, ok := item.(ProjectResponseProjects); ok {
					o.Projects = append(o.Projects, r)
				} else if r, ok := item.(map[string]interface{}); ok {
					var fm ProjectResponseProjects
					fm.FromMap(r)
					o.Projects = append(o.Projects, fm)
				} else if r, ok := item.(primitive.M); ok {
					fm := ProjectResponseProjects{}
					fm.FromMap(r)
					o.Projects = append(o.Projects, fm)
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
							var fm ProjectResponseProjects
							fm.FromMap(m[0].Interface().(map[string]interface{}))
							o.Projects = append(o.Projects, fm)
						}
					}
				}
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
	if val, ok := kv["type"].(ProjectResponseType); ok {
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
func (o *ProjectResponse) Hash() string {
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
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.LastExportDate)
	args = append(args, o.Memory)
	args = append(args, o.Message)
	args = append(args, o.NumCPU)
	args = append(args, o.OS)
	args = append(args, o.Projects)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
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

// ProjectResponsePartial is a partial struct for upsert mutations for ProjectResponse
type ProjectResponsePartial struct {
	// Architecture the architecture of the agent machine
	Architecture *string `json:"architecture,omitempty"`
	// Data extra data that is specific about this event
	Data *string `json:"data,omitempty"`
	// Distro the agent os distribution
	Distro *string `json:"distro,omitempty"`
	// Error an error message related to this event
	Error *string `json:"error,omitempty"`
	// EventDate the date of the event
	EventDate *ProjectResponseEventDate `json:"event_date,omitempty"`
	// FreeSpace the amount of free space in bytes for the agent machine
	FreeSpace *int64 `json:"free_space,omitempty"`
	// GoVersion the go version that the agent build was built with
	GoVersion *string `json:"go_version,omitempty"`
	// Hostname the agent hostname
	Hostname *string `json:"hostname,omitempty"`
	// IntegrationID the integration id
	IntegrationID *string `json:"integration_id,omitempty"`
	// LastExportDate the last export date
	LastExportDate *ProjectResponseLastExportDate `json:"last_export_date,omitempty"`
	// Memory the amount of memory in bytes for the agent machine
	Memory *int64 `json:"memory,omitempty"`
	// Message a message related to this event
	Message *string `json:"message,omitempty"`
	// NumCPU the number of CPU the agent is running
	NumCPU *int64 `json:"num_cpu,omitempty"`
	// OS the agent operating system
	OS *string `json:"os,omitempty"`
	// Projects the projects exported
	Projects []ProjectResponseProjects `json:"projects,omitempty"`
	// RequestID the request id that this response is correlated to
	RequestID *string `json:"request_id,omitempty"`
	// Success if the response was successful
	Success *bool `json:"success,omitempty"`
	// SystemID system unique device ID
	SystemID *string `json:"system_id,omitempty"`
	// Type the type of event
	Type *ProjectResponseType `json:"type,omitempty"`
	// Uptime the uptime in milliseconds since the agent started
	Uptime *int64 `json:"uptime,omitempty"`
	// UUID the agent unique identifier
	UUID *string `json:"uuid,omitempty"`
	// Version the agent version
	Version *string `json:"version,omitempty"`
}

var _ datamodel.PartialModel = (*ProjectResponsePartial)(nil)

// GetModelName returns the name of the model
func (o *ProjectResponsePartial) GetModelName() datamodel.ModelNameType {
	return ProjectResponseModelName
}

// ToMap returns the object as a map
func (o *ProjectResponsePartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"architecture":     toProjectResponseObject(o.Architecture, true),
		"data":             toProjectResponseObject(o.Data, true),
		"distro":           toProjectResponseObject(o.Distro, true),
		"error":            toProjectResponseObject(o.Error, true),
		"event_date":       toProjectResponseObject(o.EventDate, true),
		"free_space":       toProjectResponseObject(o.FreeSpace, true),
		"go_version":       toProjectResponseObject(o.GoVersion, true),
		"hostname":         toProjectResponseObject(o.Hostname, true),
		"integration_id":   toProjectResponseObject(o.IntegrationID, true),
		"last_export_date": toProjectResponseObject(o.LastExportDate, true),
		"memory":           toProjectResponseObject(o.Memory, true),
		"message":          toProjectResponseObject(o.Message, true),
		"num_cpu":          toProjectResponseObject(o.NumCPU, true),
		"os":               toProjectResponseObject(o.OS, true),
		"projects":         toProjectResponseObject(o.Projects, true),
		"request_id":       toProjectResponseObject(o.RequestID, true),
		"success":          toProjectResponseObject(o.Success, true),
		"system_id":        toProjectResponseObject(o.SystemID, true),

		"type":    toProjectResponseTypeEnum(o.Type),
		"uptime":  toProjectResponseObject(o.Uptime, true),
		"uuid":    toProjectResponseObject(o.UUID, true),
		"version": toProjectResponseObject(o.Version, true),
	}
	for k, v := range kv {
		if v == nil || reflect.ValueOf(v).IsZero() {
			delete(kv, k)
		} else {
			if k == "event_date" {
				if dt, ok := v.(*ProjectResponseEventDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}
			if k == "last_export_date" {
				if dt, ok := v.(*ProjectResponseLastExportDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}

			if k == "projects" {
				if arr, ok := v.([]ProjectResponseProjects); ok {
					if len(arr) == 0 {
						delete(kv, k)
					}
				}
			}
		}
	}
	return kv
}

// Stringify returns the object in JSON format as a string
func (o *ProjectResponsePartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *ProjectResponsePartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *ProjectResponsePartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *ProjectResponsePartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *ProjectResponsePartial) FromMap(kv map[string]interface{}) {
	if val, ok := kv["architecture"].(*string); ok {
		o.Architecture = val
	} else if val, ok := kv["architecture"].(string); ok {
		o.Architecture = &val
	} else {
		if val, ok := kv["architecture"]; ok {
			if val == nil {
				o.Architecture = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Architecture = pstrings.Pointer(fmt.Sprintf("%v", val))
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
	if val, ok := kv["distro"].(*string); ok {
		o.Distro = val
	} else if val, ok := kv["distro"].(string); ok {
		o.Distro = &val
	} else {
		if val, ok := kv["distro"]; ok {
			if val == nil {
				o.Distro = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Distro = pstrings.Pointer(fmt.Sprintf("%v", val))
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

	if o.EventDate == nil {
		o.EventDate = &ProjectResponseEventDate{}
	}

	if val, ok := kv["event_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.EventDate.FromMap(kv)
		} else if sv, ok := val.(ProjectResponseEventDate); ok {
			// struct
			o.EventDate = &sv
		} else if sp, ok := val.(*ProjectResponseEventDate); ok {
			// struct pointer
			o.EventDate = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.EventDate.Epoch = dt.Epoch
			o.EventDate.Rfc3339 = dt.Rfc3339
			o.EventDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
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

	if val, ok := kv["free_space"].(*int64); ok {
		o.FreeSpace = val
	} else if val, ok := kv["free_space"].(int64); ok {
		o.FreeSpace = &val
	} else {
		if val, ok := kv["free_space"]; ok {
			if val == nil {
				o.FreeSpace = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["long"]
				}
				o.FreeSpace = number.Int64Pointer(number.ToInt64Any(val))
			}
		}
	}
	if val, ok := kv["go_version"].(*string); ok {
		o.GoVersion = val
	} else if val, ok := kv["go_version"].(string); ok {
		o.GoVersion = &val
	} else {
		if val, ok := kv["go_version"]; ok {
			if val == nil {
				o.GoVersion = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.GoVersion = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["hostname"].(*string); ok {
		o.Hostname = val
	} else if val, ok := kv["hostname"].(string); ok {
		o.Hostname = &val
	} else {
		if val, ok := kv["hostname"]; ok {
			if val == nil {
				o.Hostname = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Hostname = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["integration_id"].(*string); ok {
		o.IntegrationID = val
	} else if val, ok := kv["integration_id"].(string); ok {
		o.IntegrationID = &val
	} else {
		if val, ok := kv["integration_id"]; ok {
			if val == nil {
				o.IntegrationID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.IntegrationID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if o.LastExportDate == nil {
		o.LastExportDate = &ProjectResponseLastExportDate{}
	}

	if val, ok := kv["last_export_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastExportDate.FromMap(kv)
		} else if sv, ok := val.(ProjectResponseLastExportDate); ok {
			// struct
			o.LastExportDate = &sv
		} else if sp, ok := val.(*ProjectResponseLastExportDate); ok {
			// struct pointer
			o.LastExportDate = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.LastExportDate.Epoch = dt.Epoch
			o.LastExportDate.Rfc3339 = dt.Rfc3339
			o.LastExportDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
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

	if val, ok := kv["memory"].(*int64); ok {
		o.Memory = val
	} else if val, ok := kv["memory"].(int64); ok {
		o.Memory = &val
	} else {
		if val, ok := kv["memory"]; ok {
			if val == nil {
				o.Memory = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["long"]
				}
				o.Memory = number.Int64Pointer(number.ToInt64Any(val))
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
	if val, ok := kv["num_cpu"].(*int64); ok {
		o.NumCPU = val
	} else if val, ok := kv["num_cpu"].(int64); ok {
		o.NumCPU = &val
	} else {
		if val, ok := kv["num_cpu"]; ok {
			if val == nil {
				o.NumCPU = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["long"]
				}
				o.NumCPU = number.Int64Pointer(number.ToInt64Any(val))
			}
		}
	}
	if val, ok := kv["os"].(*string); ok {
		o.OS = val
	} else if val, ok := kv["os"].(string); ok {
		o.OS = &val
	} else {
		if val, ok := kv["os"]; ok {
			if val == nil {
				o.OS = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.OS = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if o == nil {

		o.Projects = make([]ProjectResponseProjects, 0)

	}
	if val, ok := kv["projects"]; ok {
		if sv, ok := val.([]ProjectResponseProjects); ok {
			o.Projects = sv
		} else if sp, ok := val.([]*ProjectResponseProjects); ok {
			o.Projects = o.Projects[:0]
			for _, e := range sp {
				o.Projects = append(o.Projects, *e)
			}
		} else if a, ok := val.(primitive.A); ok {
			for _, ae := range a {
				if av, ok := ae.(ProjectResponseProjects); ok {
					o.Projects = append(o.Projects, av)
				} else if av, ok := ae.(primitive.M); ok {
					var fm ProjectResponseProjects
					fm.FromMap(av)
					o.Projects = append(o.Projects, fm)
				} else {
					b, _ := json.Marshal(ae)
					bkv := make(map[string]interface{})
					json.Unmarshal(b, &bkv)
					var av ProjectResponseProjects
					av.FromMap(bkv)
					o.Projects = append(o.Projects, av)
				}
			}
		} else if arr, ok := val.([]interface{}); ok {
			for _, item := range arr {
				if r, ok := item.(ProjectResponseProjects); ok {
					o.Projects = append(o.Projects, r)
				} else if r, ok := item.(map[string]interface{}); ok {
					var fm ProjectResponseProjects
					fm.FromMap(r)
					o.Projects = append(o.Projects, fm)
				} else if r, ok := item.(primitive.M); ok {
					fm := ProjectResponseProjects{}
					fm.FromMap(r)
					o.Projects = append(o.Projects, fm)
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
							var fm ProjectResponseProjects
							fm.FromMap(m[0].Interface().(map[string]interface{}))
							o.Projects = append(o.Projects, fm)
						}
					}
				}
			}
		}
	}

	if val, ok := kv["request_id"].(*string); ok {
		o.RequestID = val
	} else if val, ok := kv["request_id"].(string); ok {
		o.RequestID = &val
	} else {
		if val, ok := kv["request_id"]; ok {
			if val == nil {
				o.RequestID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.RequestID = pstrings.Pointer(fmt.Sprintf("%v", val))
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
	if val, ok := kv["system_id"].(*string); ok {
		o.SystemID = val
	} else if val, ok := kv["system_id"].(string); ok {
		o.SystemID = &val
	} else {
		if val, ok := kv["system_id"]; ok {
			if val == nil {
				o.SystemID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.SystemID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["type"].(*ProjectResponseType); ok {
		o.Type = val
	} else if val, ok := kv["type"].(ProjectResponseType); ok {
		o.Type = &val
	} else {
		if val, ok := kv["type"]; ok {
			if val == nil {
				o.Type = toProjectResponseTypePointer(0)
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["ProjectResponseType"]
				}
				// this is an enum pointer
				if em, ok := val.(string); ok {
					switch em {
					case "enroll", "ENROLL":
						o.Type = toProjectResponseTypePointer(0)
					case "ping", "PING":
						o.Type = toProjectResponseTypePointer(1)
					case "crash", "CRASH":
						o.Type = toProjectResponseTypePointer(2)
					case "log", "LOG":
						o.Type = toProjectResponseTypePointer(3)
					case "integration", "INTEGRATION":
						o.Type = toProjectResponseTypePointer(4)
					case "export", "EXPORT":
						o.Type = toProjectResponseTypePointer(5)
					case "project", "PROJECT":
						o.Type = toProjectResponseTypePointer(6)
					case "repo", "REPO":
						o.Type = toProjectResponseTypePointer(7)
					case "user", "USER":
						o.Type = toProjectResponseTypePointer(8)
					case "calendar", "CALENDAR":
						o.Type = toProjectResponseTypePointer(9)
					case "uninstall", "UNINSTALL":
						o.Type = toProjectResponseTypePointer(10)
					case "upgrade", "UPGRADE":
						o.Type = toProjectResponseTypePointer(11)
					case "start", "START":
						o.Type = toProjectResponseTypePointer(12)
					case "stop", "STOP":
						o.Type = toProjectResponseTypePointer(13)
					case "pause", "PAUSE":
						o.Type = toProjectResponseTypePointer(14)
					case "resume", "RESUME":
						o.Type = toProjectResponseTypePointer(15)
					}
				}
			}
		}
	}
	if val, ok := kv["uptime"].(*int64); ok {
		o.Uptime = val
	} else if val, ok := kv["uptime"].(int64); ok {
		o.Uptime = &val
	} else {
		if val, ok := kv["uptime"]; ok {
			if val == nil {
				o.Uptime = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["long"]
				}
				o.Uptime = number.Int64Pointer(number.ToInt64Any(val))
			}
		}
	}
	if val, ok := kv["uuid"].(*string); ok {
		o.UUID = val
	} else if val, ok := kv["uuid"].(string); ok {
		o.UUID = &val
	} else {
		if val, ok := kv["uuid"]; ok {
			if val == nil {
				o.UUID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.UUID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["version"].(*string); ok {
		o.Version = val
	} else if val, ok := kv["version"].(string); ok {
		o.Version = &val
	} else {
		if val, ok := kv["version"]; ok {
			if val == nil {
				o.Version = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Version = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	o.setDefaults(false)
}
