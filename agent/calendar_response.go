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

	// CalendarResponseTable is the default table name
	CalendarResponseTable datamodel.ModelNameType = "agent_calendarresponse"

	// CalendarResponseModelName is the model name
	CalendarResponseModelName datamodel.ModelNameType = "agent.CalendarResponse"
)

const (
	// CalendarResponseModelArchitectureColumn is the column json value architecture
	CalendarResponseModelArchitectureColumn = "architecture"
	// CalendarResponseModelCalendarsColumn is the column json value calendars
	CalendarResponseModelCalendarsColumn = "calendars"
	// CalendarResponseModelCalendarsActiveColumn is the column json value active
	CalendarResponseModelCalendarsActiveColumn = "active"
	// CalendarResponseModelCalendarsDescriptionColumn is the column json value description
	CalendarResponseModelCalendarsDescriptionColumn = "description"
	// CalendarResponseModelCalendarsNameColumn is the column json value name
	CalendarResponseModelCalendarsNameColumn = "name"
	// CalendarResponseModelCalendarsRefIDColumn is the column json value ref_id
	CalendarResponseModelCalendarsRefIDColumn = "ref_id"
	// CalendarResponseModelCalendarsRefTypeColumn is the column json value ref_type
	CalendarResponseModelCalendarsRefTypeColumn = "ref_type"
	// CalendarResponseModelCustomerIDColumn is the column json value customer_id
	CalendarResponseModelCustomerIDColumn = "customer_id"
	// CalendarResponseModelDataColumn is the column json value data
	CalendarResponseModelDataColumn = "data"
	// CalendarResponseModelDistroColumn is the column json value distro
	CalendarResponseModelDistroColumn = "distro"
	// CalendarResponseModelErrorColumn is the column json value error
	CalendarResponseModelErrorColumn = "error"
	// CalendarResponseModelEventDateColumn is the column json value event_date
	CalendarResponseModelEventDateColumn = "event_date"
	// CalendarResponseModelEventDateEpochColumn is the column json value epoch
	CalendarResponseModelEventDateEpochColumn = "epoch"
	// CalendarResponseModelEventDateOffsetColumn is the column json value offset
	CalendarResponseModelEventDateOffsetColumn = "offset"
	// CalendarResponseModelEventDateRfc3339Column is the column json value rfc3339
	CalendarResponseModelEventDateRfc3339Column = "rfc3339"
	// CalendarResponseModelFreeSpaceColumn is the column json value free_space
	CalendarResponseModelFreeSpaceColumn = "free_space"
	// CalendarResponseModelGoVersionColumn is the column json value go_version
	CalendarResponseModelGoVersionColumn = "go_version"
	// CalendarResponseModelHostnameColumn is the column json value hostname
	CalendarResponseModelHostnameColumn = "hostname"
	// CalendarResponseModelIDColumn is the column json value id
	CalendarResponseModelIDColumn = "id"
	// CalendarResponseModelIntegrationIDColumn is the column json value integration_id
	CalendarResponseModelIntegrationIDColumn = "integration_id"
	// CalendarResponseModelLastExportDateColumn is the column json value last_export_date
	CalendarResponseModelLastExportDateColumn = "last_export_date"
	// CalendarResponseModelLastExportDateEpochColumn is the column json value epoch
	CalendarResponseModelLastExportDateEpochColumn = "epoch"
	// CalendarResponseModelLastExportDateOffsetColumn is the column json value offset
	CalendarResponseModelLastExportDateOffsetColumn = "offset"
	// CalendarResponseModelLastExportDateRfc3339Column is the column json value rfc3339
	CalendarResponseModelLastExportDateRfc3339Column = "rfc3339"
	// CalendarResponseModelMemoryColumn is the column json value memory
	CalendarResponseModelMemoryColumn = "memory"
	// CalendarResponseModelMessageColumn is the column json value message
	CalendarResponseModelMessageColumn = "message"
	// CalendarResponseModelNumCPUColumn is the column json value num_cpu
	CalendarResponseModelNumCPUColumn = "num_cpu"
	// CalendarResponseModelOSColumn is the column json value os
	CalendarResponseModelOSColumn = "os"
	// CalendarResponseModelRefIDColumn is the column json value ref_id
	CalendarResponseModelRefIDColumn = "ref_id"
	// CalendarResponseModelRefTypeColumn is the column json value ref_type
	CalendarResponseModelRefTypeColumn = "ref_type"
	// CalendarResponseModelRequestIDColumn is the column json value request_id
	CalendarResponseModelRequestIDColumn = "request_id"
	// CalendarResponseModelSuccessColumn is the column json value success
	CalendarResponseModelSuccessColumn = "success"
	// CalendarResponseModelSystemIDColumn is the column json value system_id
	CalendarResponseModelSystemIDColumn = "system_id"
	// CalendarResponseModelTypeColumn is the column json value type
	CalendarResponseModelTypeColumn = "type"
	// CalendarResponseModelUptimeColumn is the column json value uptime
	CalendarResponseModelUptimeColumn = "uptime"
	// CalendarResponseModelUUIDColumn is the column json value uuid
	CalendarResponseModelUUIDColumn = "uuid"
	// CalendarResponseModelVersionColumn is the column json value version
	CalendarResponseModelVersionColumn = "version"
)

// CalendarResponseCalendars represents the object structure for calendars
type CalendarResponseCalendars struct {
	// Active the status of the calendar
	Active bool `json:"active" codec:"active" bson:"active" yaml:"active" faker:"-"`
	// Description the description of the calendar
	Description string `json:"description" codec:"description" bson:"description" yaml:"description" faker:"-"`
	// Name the name of the calendar
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
	// RefID the calendar ID
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the record type
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
}

func toCalendarResponseCalendarsObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CalendarResponseCalendars:
		return v.ToMap()

	default:
		return o
	}
}

func (o *CalendarResponseCalendars) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Active the status of the calendar
		"active": toCalendarResponseCalendarsObject(o.Active, false),
		// Description the description of the calendar
		"description": toCalendarResponseCalendarsObject(o.Description, false),
		// Name the name of the calendar
		"name": toCalendarResponseCalendarsObject(o.Name, false),
		// RefID the calendar ID
		"ref_id": toCalendarResponseCalendarsObject(o.RefID, false),
		// RefType the record type
		"ref_type": toCalendarResponseCalendarsObject(o.RefType, false),
	}
}

func (o *CalendarResponseCalendars) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *CalendarResponseCalendars) FromMap(kv map[string]interface{}) {

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
	o.setDefaults(false)
}

// CalendarResponseEventDate represents the object structure for event_date
type CalendarResponseEventDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toCalendarResponseEventDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CalendarResponseEventDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *CalendarResponseEventDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toCalendarResponseEventDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toCalendarResponseEventDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toCalendarResponseEventDateObject(o.Rfc3339, false),
	}
}

func (o *CalendarResponseEventDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *CalendarResponseEventDate) FromMap(kv map[string]interface{}) {

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

// CalendarResponseLastExportDate represents the object structure for last_export_date
type CalendarResponseLastExportDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toCalendarResponseLastExportDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CalendarResponseLastExportDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *CalendarResponseLastExportDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toCalendarResponseLastExportDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toCalendarResponseLastExportDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toCalendarResponseLastExportDateObject(o.Rfc3339, false),
	}
}

func (o *CalendarResponseLastExportDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *CalendarResponseLastExportDate) FromMap(kv map[string]interface{}) {

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

// CalendarResponseType is the enumeration type for type
type CalendarResponseType int32

// UnmarshalBSONValue for unmarshaling value
func (v *CalendarResponseType) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = CalendarResponseType(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "ENROLL":
			*v = CalendarResponseType(0)
		case "PING":
			*v = CalendarResponseType(1)
		case "CRASH":
			*v = CalendarResponseType(2)
		case "LOG":
			*v = CalendarResponseType(3)
		case "INTEGRATION":
			*v = CalendarResponseType(4)
		case "EXPORT":
			*v = CalendarResponseType(5)
		case "PROJECT":
			*v = CalendarResponseType(6)
		case "REPO":
			*v = CalendarResponseType(7)
		case "USER":
			*v = CalendarResponseType(8)
		case "CALENDAR":
			*v = CalendarResponseType(9)
		case "UNINSTALL":
			*v = CalendarResponseType(10)
		case "UPGRADE":
			*v = CalendarResponseType(11)
		case "START":
			*v = CalendarResponseType(12)
		case "STOP":
			*v = CalendarResponseType(13)
		case "PAUSE":
			*v = CalendarResponseType(14)
		case "RESUME":
			*v = CalendarResponseType(15)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v CalendarResponseType) UnmarshalJSON(buf []byte) error {
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
func (v CalendarResponseType) MarshalJSON() ([]byte, error) {
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
func (v CalendarResponseType) String() string {
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

const (
	// CalendarResponseTypeEnroll is the enumeration value for enroll
	CalendarResponseTypeEnroll CalendarResponseType = 0
	// CalendarResponseTypePing is the enumeration value for ping
	CalendarResponseTypePing CalendarResponseType = 1
	// CalendarResponseTypeCrash is the enumeration value for crash
	CalendarResponseTypeCrash CalendarResponseType = 2
	// CalendarResponseTypeLog is the enumeration value for log
	CalendarResponseTypeLog CalendarResponseType = 3
	// CalendarResponseTypeIntegration is the enumeration value for integration
	CalendarResponseTypeIntegration CalendarResponseType = 4
	// CalendarResponseTypeExport is the enumeration value for export
	CalendarResponseTypeExport CalendarResponseType = 5
	// CalendarResponseTypeProject is the enumeration value for project
	CalendarResponseTypeProject CalendarResponseType = 6
	// CalendarResponseTypeRepo is the enumeration value for repo
	CalendarResponseTypeRepo CalendarResponseType = 7
	// CalendarResponseTypeUser is the enumeration value for user
	CalendarResponseTypeUser CalendarResponseType = 8
	// CalendarResponseTypeCalendar is the enumeration value for calendar
	CalendarResponseTypeCalendar CalendarResponseType = 9
	// CalendarResponseTypeUninstall is the enumeration value for uninstall
	CalendarResponseTypeUninstall CalendarResponseType = 10
	// CalendarResponseTypeUpgrade is the enumeration value for upgrade
	CalendarResponseTypeUpgrade CalendarResponseType = 11
	// CalendarResponseTypeStart is the enumeration value for start
	CalendarResponseTypeStart CalendarResponseType = 12
	// CalendarResponseTypeStop is the enumeration value for stop
	CalendarResponseTypeStop CalendarResponseType = 13
	// CalendarResponseTypePause is the enumeration value for pause
	CalendarResponseTypePause CalendarResponseType = 14
	// CalendarResponseTypeResume is the enumeration value for resume
	CalendarResponseTypeResume CalendarResponseType = 15
)

// CalendarResponse an agent response to an action request adding calendar(s)
type CalendarResponse struct {
	// Architecture the architecture of the agent machine
	Architecture string `json:"architecture" codec:"architecture" bson:"architecture" yaml:"architecture" faker:"-"`
	// Calendars the calendars exported
	Calendars []CalendarResponseCalendars `json:"calendars" codec:"calendars" bson:"calendars" yaml:"calendars" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Data extra data that is specific about this event
	Data *string `json:"data,omitempty" codec:"data,omitempty" bson:"data" yaml:"data,omitempty" faker:"-"`
	// Distro the agent os distribution
	Distro string `json:"distro" codec:"distro" bson:"distro" yaml:"distro" faker:"-"`
	// Error an error message related to this event
	Error *string `json:"error,omitempty" codec:"error,omitempty" bson:"error" yaml:"error,omitempty" faker:"-"`
	// EventDate the date of the event
	EventDate CalendarResponseEventDate `json:"event_date" codec:"event_date" bson:"event_date" yaml:"event_date" faker:"-"`
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
	LastExportDate CalendarResponseLastExportDate `json:"last_export_date" codec:"last_export_date" bson:"last_export_date" yaml:"last_export_date" faker:"-"`
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
	Type CalendarResponseType `json:"type" codec:"type" bson:"type" yaml:"type" faker:"-"`
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
var _ datamodel.Model = (*CalendarResponse)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*CalendarResponse)(nil)

func toCalendarResponseObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CalendarResponse:
		return v.ToMap()

	case []CalendarResponseCalendars:
		arr := make([]interface{}, 0)
		for _, i := range v {
			arr = append(arr, i.ToMap())
		}
		return arr

	case CalendarResponseEventDate:
		return v.ToMap()

	case CalendarResponseLastExportDate:
		return v.ToMap()

	case CalendarResponseType:
		return v.String()

	default:
		return o
	}
}

// String returns a string representation of CalendarResponse
func (o *CalendarResponse) String() string {
	return fmt.Sprintf("agent.CalendarResponse<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *CalendarResponse) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *CalendarResponse) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *CalendarResponse) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *CalendarResponse) GetModelName() datamodel.ModelNameType {
	return CalendarResponseModelName
}

// NewCalendarResponseID provides a template for generating an ID field for CalendarResponse
func NewCalendarResponseID(customerID string, refType string, refID string) string {
	return hash.Values("CalendarResponse", customerID, refType, refID)
}

func (o *CalendarResponse) setDefaults(frommap bool) {
	if o.Calendars == nil {
		o.Calendars = make([]CalendarResponseCalendars, 0)
	}

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("CalendarResponse", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *CalendarResponse) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *CalendarResponse) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *CalendarResponse) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *CalendarResponse) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *CalendarResponse) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *CalendarResponse) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *CalendarResponse) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *CalendarResponse) IsEvented() bool {
	return false
}

// GetTopicConfig returns the topic config object
func (o *CalendarResponse) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *CalendarResponse) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of CalendarResponse
func (o *CalendarResponse) Clone() datamodel.Model {
	c := new(CalendarResponse)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *CalendarResponse) Anon() datamodel.Model {
	c := new(CalendarResponse)
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
func (o *CalendarResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *CalendarResponse) UnmarshalJSON(data []byte) error {
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
func (o *CalendarResponse) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *CalendarResponse) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two CalendarResponse objects are equal
func (o *CalendarResponse) IsEqual(other *CalendarResponse) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *CalendarResponse) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"architecture":     toCalendarResponseObject(o.Architecture, false),
		"calendars":        toCalendarResponseObject(o.Calendars, false),
		"customer_id":      toCalendarResponseObject(o.CustomerID, false),
		"data":             toCalendarResponseObject(o.Data, true),
		"distro":           toCalendarResponseObject(o.Distro, false),
		"error":            toCalendarResponseObject(o.Error, true),
		"event_date":       toCalendarResponseObject(o.EventDate, false),
		"free_space":       toCalendarResponseObject(o.FreeSpace, false),
		"go_version":       toCalendarResponseObject(o.GoVersion, false),
		"hostname":         toCalendarResponseObject(o.Hostname, false),
		"id":               toCalendarResponseObject(o.ID, false),
		"integration_id":   toCalendarResponseObject(o.IntegrationID, false),
		"last_export_date": toCalendarResponseObject(o.LastExportDate, false),
		"memory":           toCalendarResponseObject(o.Memory, false),
		"message":          toCalendarResponseObject(o.Message, false),
		"num_cpu":          toCalendarResponseObject(o.NumCPU, false),
		"os":               toCalendarResponseObject(o.OS, false),
		"ref_id":           toCalendarResponseObject(o.RefID, false),
		"ref_type":         toCalendarResponseObject(o.RefType, false),
		"request_id":       toCalendarResponseObject(o.RequestID, false),
		"success":          toCalendarResponseObject(o.Success, false),
		"system_id":        toCalendarResponseObject(o.SystemID, false),

		"type":     o.Type.String(),
		"uptime":   toCalendarResponseObject(o.Uptime, false),
		"uuid":     toCalendarResponseObject(o.UUID, false),
		"version":  toCalendarResponseObject(o.Version, false),
		"hashcode": toCalendarResponseObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *CalendarResponse) FromMap(kv map[string]interface{}) {

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

	if o == nil {

		o.Calendars = make([]CalendarResponseCalendars, 0)

	}
	if val, ok := kv["calendars"]; ok {
		if sv, ok := val.([]CalendarResponseCalendars); ok {
			o.Calendars = sv
		} else if sp, ok := val.([]*CalendarResponseCalendars); ok {
			o.Calendars = o.Calendars[:0]
			for _, e := range sp {
				o.Calendars = append(o.Calendars, *e)
			}
		} else if a, ok := val.(primitive.A); ok {
			for _, ae := range a {
				if av, ok := ae.(CalendarResponseCalendars); ok {
					o.Calendars = append(o.Calendars, av)
				} else if av, ok := ae.(primitive.M); ok {
					var fm CalendarResponseCalendars
					fm.FromMap(av)
					o.Calendars = append(o.Calendars, fm)
				} else {
					b, _ := json.Marshal(ae)
					bkv := make(map[string]interface{})
					json.Unmarshal(b, &bkv)
					var av CalendarResponseCalendars
					av.FromMap(bkv)
					o.Calendars = append(o.Calendars, av)
				}
			}
		} else if arr, ok := val.([]interface{}); ok {
			for _, item := range arr {
				if r, ok := item.(CalendarResponseCalendars); ok {
					o.Calendars = append(o.Calendars, r)
				} else if r, ok := item.(map[string]interface{}); ok {
					var fm CalendarResponseCalendars
					fm.FromMap(r)
					o.Calendars = append(o.Calendars, fm)
				} else if r, ok := item.(primitive.M); ok {
					fm := CalendarResponseCalendars{}
					fm.FromMap(r)
					o.Calendars = append(o.Calendars, fm)
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
							var fm CalendarResponseCalendars
							fm.FromMap(m[0].Interface().(map[string]interface{}))
							o.Calendars = append(o.Calendars, fm)
						}
					}
				}
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
		} else if sv, ok := val.(CalendarResponseEventDate); ok {
			// struct
			o.EventDate = sv
		} else if sp, ok := val.(*CalendarResponseEventDate); ok {
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
		} else if sv, ok := val.(CalendarResponseLastExportDate); ok {
			// struct
			o.LastExportDate = sv
		} else if sp, ok := val.(*CalendarResponseLastExportDate); ok {
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

	if val, ok := kv["type"].(CalendarResponseType); ok {
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
func (o *CalendarResponse) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Architecture)
	args = append(args, o.Calendars)
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
	args = append(args, o.Uptime)
	args = append(args, o.UUID)
	args = append(args, o.Version)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *CalendarResponse) GetEventAPIConfig() datamodel.EventAPIConfig {
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
