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
)

const (

	// CrashTable is the default table name
	CrashTable datamodel.ModelNameType = "agent_crash"

	// CrashModelName is the model name
	CrashModelName datamodel.ModelNameType = "agent.Crash"
)

const (
	// CrashModelArchitectureColumn is the column json value architecture
	CrashModelArchitectureColumn = "architecture"
	// CrashModelComponentColumn is the column json value component
	CrashModelComponentColumn = "component"
	// CrashModelCrashDateColumn is the column json value crash_date
	CrashModelCrashDateColumn = "crash_date"
	// CrashModelCrashDateEpochColumn is the column json value epoch
	CrashModelCrashDateEpochColumn = "epoch"
	// CrashModelCrashDateOffsetColumn is the column json value offset
	CrashModelCrashDateOffsetColumn = "offset"
	// CrashModelCrashDateRfc3339Column is the column json value rfc3339
	CrashModelCrashDateRfc3339Column = "rfc3339"
	// CrashModelCustomerIDColumn is the column json value customer_id
	CrashModelCustomerIDColumn = "customer_id"
	// CrashModelDataColumn is the column json value data
	CrashModelDataColumn = "data"
	// CrashModelDistroColumn is the column json value distro
	CrashModelDistroColumn = "distro"
	// CrashModelErrorColumn is the column json value error
	CrashModelErrorColumn = "error"
	// CrashModelEventDateColumn is the column json value event_date
	CrashModelEventDateColumn = "event_date"
	// CrashModelEventDateEpochColumn is the column json value epoch
	CrashModelEventDateEpochColumn = "epoch"
	// CrashModelEventDateOffsetColumn is the column json value offset
	CrashModelEventDateOffsetColumn = "offset"
	// CrashModelEventDateRfc3339Column is the column json value rfc3339
	CrashModelEventDateRfc3339Column = "rfc3339"
	// CrashModelFreeSpaceColumn is the column json value free_space
	CrashModelFreeSpaceColumn = "free_space"
	// CrashModelGoVersionColumn is the column json value go_version
	CrashModelGoVersionColumn = "go_version"
	// CrashModelHostnameColumn is the column json value hostname
	CrashModelHostnameColumn = "hostname"
	// CrashModelIDColumn is the column json value id
	CrashModelIDColumn = "id"
	// CrashModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	CrashModelIntegrationInstanceIDColumn = "integration_instance_id"
	// CrashModelLastExportDateColumn is the column json value last_export_date
	CrashModelLastExportDateColumn = "last_export_date"
	// CrashModelLastExportDateEpochColumn is the column json value epoch
	CrashModelLastExportDateEpochColumn = "epoch"
	// CrashModelLastExportDateOffsetColumn is the column json value offset
	CrashModelLastExportDateOffsetColumn = "offset"
	// CrashModelLastExportDateRfc3339Column is the column json value rfc3339
	CrashModelLastExportDateRfc3339Column = "rfc3339"
	// CrashModelMemoryColumn is the column json value memory
	CrashModelMemoryColumn = "memory"
	// CrashModelMessageColumn is the column json value message
	CrashModelMessageColumn = "message"
	// CrashModelNumCPUColumn is the column json value num_cpu
	CrashModelNumCPUColumn = "num_cpu"
	// CrashModelOSColumn is the column json value os
	CrashModelOSColumn = "os"
	// CrashModelRefIDColumn is the column json value ref_id
	CrashModelRefIDColumn = "ref_id"
	// CrashModelRefTypeColumn is the column json value ref_type
	CrashModelRefTypeColumn = "ref_type"
	// CrashModelRequestIDColumn is the column json value request_id
	CrashModelRequestIDColumn = "request_id"
	// CrashModelSuccessColumn is the column json value success
	CrashModelSuccessColumn = "success"
	// CrashModelSystemIDColumn is the column json value system_id
	CrashModelSystemIDColumn = "system_id"
	// CrashModelTypeColumn is the column json value type
	CrashModelTypeColumn = "type"
	// CrashModelUptimeColumn is the column json value uptime
	CrashModelUptimeColumn = "uptime"
	// CrashModelUUIDColumn is the column json value uuid
	CrashModelUUIDColumn = "uuid"
	// CrashModelVersionColumn is the column json value version
	CrashModelVersionColumn = "version"
)

// CrashCrashDate represents the object structure for crash_date
type CrashCrashDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toCrashCrashDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CrashCrashDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *CrashCrashDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toCrashCrashDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toCrashCrashDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toCrashCrashDateObject(o.Rfc3339, false),
	}
}

func (o *CrashCrashDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *CrashCrashDate) FromMap(kv map[string]interface{}) {

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

// CrashEventDate represents the object structure for event_date
type CrashEventDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toCrashEventDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CrashEventDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *CrashEventDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toCrashEventDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toCrashEventDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toCrashEventDateObject(o.Rfc3339, false),
	}
}

func (o *CrashEventDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *CrashEventDate) FromMap(kv map[string]interface{}) {

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

// CrashLastExportDate represents the object structure for last_export_date
type CrashLastExportDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toCrashLastExportDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CrashLastExportDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *CrashLastExportDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toCrashLastExportDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toCrashLastExportDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toCrashLastExportDateObject(o.Rfc3339, false),
	}
}

func (o *CrashLastExportDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *CrashLastExportDate) FromMap(kv map[string]interface{}) {

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

// CrashType is the enumeration type for type
type CrashType int32

// toCrashTypePointer is the enumeration pointer type for type
func toCrashTypePointer(v int32) *CrashType {
	nv := CrashType(v)
	return &nv
}

// toCrashTypeEnum is the enumeration pointer wrapper for type
func toCrashTypeEnum(v *CrashType) string {
	if v == nil {
		return toCrashTypePointer(0).String()
	}
	return v.String()
}

// UnmarshalBSONValue for unmarshaling value
func (v *CrashType) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = CrashType(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "ENROLL":
			*v = CrashType(0)
		case "PING":
			*v = CrashType(1)
		case "CRASH":
			*v = CrashType(2)
		case "LOG":
			*v = CrashType(3)
		case "INTEGRATION":
			*v = CrashType(4)
		case "EXPORT":
			*v = CrashType(5)
		case "PROJECT":
			*v = CrashType(6)
		case "REPO":
			*v = CrashType(7)
		case "USER":
			*v = CrashType(8)
		case "CALENDAR":
			*v = CrashType(9)
		case "UNINSTALL":
			*v = CrashType(10)
		case "UPGRADE":
			*v = CrashType(11)
		case "START":
			*v = CrashType(12)
		case "STOP":
			*v = CrashType(13)
		case "PAUSE":
			*v = CrashType(14)
		case "RESUME":
			*v = CrashType(15)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v *CrashType) UnmarshalJSON(buf []byte) error {
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
func (v CrashType) MarshalJSON() ([]byte, error) {
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
func (v CrashType) String() string {
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
func (v *CrashType) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case CrashType:
		*v = val
	case int32:
		*v = CrashType(int32(val))
	case int:
		*v = CrashType(int32(val))
	case string:
		switch val {
		case "ENROLL":
			*v = CrashType(0)
		case "PING":
			*v = CrashType(1)
		case "CRASH":
			*v = CrashType(2)
		case "LOG":
			*v = CrashType(3)
		case "INTEGRATION":
			*v = CrashType(4)
		case "EXPORT":
			*v = CrashType(5)
		case "PROJECT":
			*v = CrashType(6)
		case "REPO":
			*v = CrashType(7)
		case "USER":
			*v = CrashType(8)
		case "CALENDAR":
			*v = CrashType(9)
		case "UNINSTALL":
			*v = CrashType(10)
		case "UPGRADE":
			*v = CrashType(11)
		case "START":
			*v = CrashType(12)
		case "STOP":
			*v = CrashType(13)
		case "PAUSE":
			*v = CrashType(14)
		case "RESUME":
			*v = CrashType(15)
		}
	}
	return nil
}

const (
	// CrashTypeEnroll is the enumeration value for enroll
	CrashTypeEnroll CrashType = 0
	// CrashTypePing is the enumeration value for ping
	CrashTypePing CrashType = 1
	// CrashTypeCrash is the enumeration value for crash
	CrashTypeCrash CrashType = 2
	// CrashTypeLog is the enumeration value for log
	CrashTypeLog CrashType = 3
	// CrashTypeIntegration is the enumeration value for integration
	CrashTypeIntegration CrashType = 4
	// CrashTypeExport is the enumeration value for export
	CrashTypeExport CrashType = 5
	// CrashTypeProject is the enumeration value for project
	CrashTypeProject CrashType = 6
	// CrashTypeRepo is the enumeration value for repo
	CrashTypeRepo CrashType = 7
	// CrashTypeUser is the enumeration value for user
	CrashTypeUser CrashType = 8
	// CrashTypeCalendar is the enumeration value for calendar
	CrashTypeCalendar CrashType = 9
	// CrashTypeUninstall is the enumeration value for uninstall
	CrashTypeUninstall CrashType = 10
	// CrashTypeUpgrade is the enumeration value for upgrade
	CrashTypeUpgrade CrashType = 11
	// CrashTypeStart is the enumeration value for start
	CrashTypeStart CrashType = 12
	// CrashTypeStop is the enumeration value for stop
	CrashTypeStop CrashType = 13
	// CrashTypePause is the enumeration value for pause
	CrashTypePause CrashType = 14
	// CrashTypeResume is the enumeration value for resume
	CrashTypeResume CrashType = 15
)

// Crash an agent event to indicate that the agent crashed
type Crash struct {
	// Architecture the architecture of the agent machine
	Architecture string `json:"architecture" codec:"architecture" bson:"architecture" yaml:"architecture" faker:"-"`
	// Component Component that caused panic. Could be one of integrations, export subcommands or service itself.
	Component string `json:"component" codec:"component" bson:"component" yaml:"component" faker:"-"`
	// CrashDate The date when crash happened. We may not be able to send the crash immediately due to network errors.
	CrashDate CrashCrashDate `json:"crash_date" codec:"crash_date" bson:"crash_date" yaml:"crash_date" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Data extra data that is specific about this event
	Data *string `json:"data,omitempty" codec:"data,omitempty" bson:"data" yaml:"data,omitempty" faker:"-"`
	// Distro the agent os distribution
	Distro string `json:"distro" codec:"distro" bson:"distro" yaml:"distro" faker:"-"`
	// Error an error message related to this event
	Error *string `json:"error,omitempty" codec:"error,omitempty" bson:"error" yaml:"error,omitempty" faker:"-"`
	// EventDate the date of the event
	EventDate CrashEventDate `json:"event_date" codec:"event_date" bson:"event_date" yaml:"event_date" faker:"-"`
	// FreeSpace the amount of free space in bytes for the agent machine
	FreeSpace int64 `json:"free_space" codec:"free_space" bson:"free_space" yaml:"free_space" faker:"-"`
	// GoVersion the go version that the agent build was built with
	GoVersion string `json:"go_version" codec:"go_version" bson:"go_version" yaml:"go_version" faker:"-"`
	// Hostname the agent hostname
	Hostname string `json:"hostname" codec:"hostname" bson:"hostname" yaml:"hostname" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// LastExportDate the last export date
	LastExportDate CrashLastExportDate `json:"last_export_date" codec:"last_export_date" bson:"last_export_date" yaml:"last_export_date" faker:"-"`
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
	Type CrashType `json:"type" codec:"type" bson:"type" yaml:"type" faker:"-"`
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
var _ datamodel.Model = (*Crash)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*Crash)(nil)

func toCrashObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *Crash:
		return v.ToMap()

	case CrashCrashDate:
		return v.ToMap()

	case CrashEventDate:
		return v.ToMap()

	case CrashLastExportDate:
		return v.ToMap()

	case CrashType:
		return v.String()

	default:
		return o
	}
}

// String returns a string representation of Crash
func (o *Crash) String() string {
	return fmt.Sprintf("agent.Crash<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Crash) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *Crash) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *Crash) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *Crash) GetModelName() datamodel.ModelNameType {
	return CrashModelName
}

// NewCrashID provides a template for generating an ID field for Crash
func NewCrashID(customerID string, refType string, refID string) string {
	return hash.Values("Crash", customerID, refType, refID)
}

func (o *Crash) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Crash", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Crash) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Crash) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Crash) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *Crash) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Crash) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *Crash) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Crash) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Crash) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *Crash) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = CrashModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Crash) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *Crash) GetCustomerID() string {
	return o.CustomerID
}

// SetCustomerID will return the customer_id
func (o *Crash) SetCustomerID(id string) {
	o.CustomerID = id
}

// GetRefType will return the ref_type
func (o *Crash) GetRefType() string {
	return o.RefType
}

// SetRefType will return the ref_type
func (o *Crash) SetRefType(t string) {
	o.RefType = t
}

// Clone returns an exact copy of Crash
func (o *Crash) Clone() datamodel.Model {
	c := new(Crash)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Crash) Anon() datamodel.Model {
	c := new(Crash)
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
func (o *Crash) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Crash) UnmarshalJSON(data []byte) error {
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
func (o *Crash) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *Crash) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two Crash objects are equal
func (o *Crash) IsEqual(other *Crash) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Crash) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"architecture":            toCrashObject(o.Architecture, false),
		"component":               toCrashObject(o.Component, false),
		"crash_date":              toCrashObject(o.CrashDate, false),
		"customer_id":             toCrashObject(o.CustomerID, false),
		"data":                    toCrashObject(o.Data, true),
		"distro":                  toCrashObject(o.Distro, false),
		"error":                   toCrashObject(o.Error, true),
		"event_date":              toCrashObject(o.EventDate, false),
		"free_space":              toCrashObject(o.FreeSpace, false),
		"go_version":              toCrashObject(o.GoVersion, false),
		"hostname":                toCrashObject(o.Hostname, false),
		"id":                      toCrashObject(o.ID, false),
		"integration_instance_id": toCrashObject(o.IntegrationInstanceID, true),
		"last_export_date":        toCrashObject(o.LastExportDate, false),
		"memory":                  toCrashObject(o.Memory, false),
		"message":                 toCrashObject(o.Message, false),
		"num_cpu":                 toCrashObject(o.NumCPU, false),
		"os":                      toCrashObject(o.OS, false),
		"ref_id":                  toCrashObject(o.RefID, false),
		"ref_type":                toCrashObject(o.RefType, false),
		"request_id":              toCrashObject(o.RequestID, false),
		"success":                 toCrashObject(o.Success, false),
		"system_id":               toCrashObject(o.SystemID, false),

		"type":     o.Type.String(),
		"uptime":   toCrashObject(o.Uptime, false),
		"uuid":     toCrashObject(o.UUID, false),
		"version":  toCrashObject(o.Version, false),
		"hashcode": toCrashObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *Crash) FromMap(kv map[string]interface{}) {

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
	if val, ok := kv["component"].(string); ok {
		o.Component = val
	} else {
		if val, ok := kv["component"]; ok {
			if val == nil {
				o.Component = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Component = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["crash_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.CrashDate.FromMap(kv)
		} else if sv, ok := val.(CrashCrashDate); ok {
			// struct
			o.CrashDate = sv
		} else if sp, ok := val.(*CrashCrashDate); ok {
			// struct pointer
			o.CrashDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.CrashDate.Epoch = dt.Epoch
			o.CrashDate.Rfc3339 = dt.Rfc3339
			o.CrashDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.CrashDate.Epoch = dt.Epoch
			o.CrashDate.Rfc3339 = dt.Rfc3339
			o.CrashDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.CrashDate.Epoch = dt.Epoch
				o.CrashDate.Rfc3339 = dt.Rfc3339
				o.CrashDate.Offset = dt.Offset
			}
		}
	} else {
		o.CrashDate.FromMap(map[string]interface{}{})
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
		} else if sv, ok := val.(CrashEventDate); ok {
			// struct
			o.EventDate = sv
		} else if sp, ok := val.(*CrashEventDate); ok {
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
		} else if sv, ok := val.(CrashLastExportDate); ok {
			// struct
			o.LastExportDate = sv
		} else if sp, ok := val.(*CrashLastExportDate); ok {
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
	if val, ok := kv["type"].(CrashType); ok {
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
func (o *Crash) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Architecture)
	args = append(args, o.Component)
	args = append(args, o.CrashDate)
	args = append(args, o.CustomerID)
	args = append(args, o.Data)
	args = append(args, o.Distro)
	args = append(args, o.Error)
	args = append(args, o.EventDate)
	args = append(args, o.FreeSpace)
	args = append(args, o.GoVersion)
	args = append(args, o.Hostname)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
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

// SetIntegrationInstanceID will set the integration instance ID
func (o *Crash) SetIntegrationInstanceID(id string) {
	if id == "" {
		o.IntegrationInstanceID = nil
	} else {
		o.IntegrationInstanceID = &id
	}
}

// GetIntegrationInstanceID will return the integration instance ID
func (o *Crash) GetIntegrationInstanceID() *string {
	return o.IntegrationInstanceID
}

// CrashPartial is a partial struct for upsert mutations for Crash
type CrashPartial struct {
	// Architecture the architecture of the agent machine
	Architecture *string `json:"architecture,omitempty"`
	// Component Component that caused panic. Could be one of integrations, export subcommands or service itself.
	Component *string `json:"component,omitempty"`
	// CrashDate The date when crash happened. We may not be able to send the crash immediately due to network errors.
	CrashDate *CrashCrashDate `json:"crash_date,omitempty"`
	// Data extra data that is specific about this event
	Data *string `json:"data,omitempty"`
	// Distro the agent os distribution
	Distro *string `json:"distro,omitempty"`
	// Error an error message related to this event
	Error *string `json:"error,omitempty"`
	// EventDate the date of the event
	EventDate *CrashEventDate `json:"event_date,omitempty"`
	// FreeSpace the amount of free space in bytes for the agent machine
	FreeSpace *int64 `json:"free_space,omitempty"`
	// GoVersion the go version that the agent build was built with
	GoVersion *string `json:"go_version,omitempty"`
	// Hostname the agent hostname
	Hostname *string `json:"hostname,omitempty"`
	// LastExportDate the last export date
	LastExportDate *CrashLastExportDate `json:"last_export_date,omitempty"`
	// Memory the amount of memory in bytes for the agent machine
	Memory *int64 `json:"memory,omitempty"`
	// Message a message related to this event
	Message *string `json:"message,omitempty"`
	// NumCPU the number of CPU the agent is running
	NumCPU *int64 `json:"num_cpu,omitempty"`
	// OS the agent operating system
	OS *string `json:"os,omitempty"`
	// RequestID the request id that this response is correlated to
	RequestID *string `json:"request_id,omitempty"`
	// Success if the response was successful
	Success *bool `json:"success,omitempty"`
	// SystemID system unique device ID
	SystemID *string `json:"system_id,omitempty"`
	// Type the type of event
	Type *CrashType `json:"type,omitempty"`
	// Uptime the uptime in milliseconds since the agent started
	Uptime *int64 `json:"uptime,omitempty"`
	// UUID the agent unique identifier
	UUID *string `json:"uuid,omitempty"`
	// Version the agent version
	Version *string `json:"version,omitempty"`
}

var _ datamodel.PartialModel = (*CrashPartial)(nil)

// GetModelName returns the name of the model
func (o *CrashPartial) GetModelName() datamodel.ModelNameType {
	return CrashModelName
}

// ToMap returns the object as a map
func (o *CrashPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"architecture":     toCrashObject(o.Architecture, true),
		"component":        toCrashObject(o.Component, true),
		"crash_date":       toCrashObject(o.CrashDate, true),
		"data":             toCrashObject(o.Data, true),
		"distro":           toCrashObject(o.Distro, true),
		"error":            toCrashObject(o.Error, true),
		"event_date":       toCrashObject(o.EventDate, true),
		"free_space":       toCrashObject(o.FreeSpace, true),
		"go_version":       toCrashObject(o.GoVersion, true),
		"hostname":         toCrashObject(o.Hostname, true),
		"last_export_date": toCrashObject(o.LastExportDate, true),
		"memory":           toCrashObject(o.Memory, true),
		"message":          toCrashObject(o.Message, true),
		"num_cpu":          toCrashObject(o.NumCPU, true),
		"os":               toCrashObject(o.OS, true),
		"request_id":       toCrashObject(o.RequestID, true),
		"success":          toCrashObject(o.Success, true),
		"system_id":        toCrashObject(o.SystemID, true),

		"type":    toCrashTypeEnum(o.Type),
		"uptime":  toCrashObject(o.Uptime, true),
		"uuid":    toCrashObject(o.UUID, true),
		"version": toCrashObject(o.Version, true),
	}
	for k, v := range kv {
		if v == nil || reflect.ValueOf(v).IsZero() {
			delete(kv, k)
		} else {
			if k == "crash_date" {
				if dt, ok := v.(*CrashCrashDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}
			if k == "event_date" {
				if dt, ok := v.(*CrashEventDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}
			if k == "last_export_date" {
				if dt, ok := v.(*CrashLastExportDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}
		}
	}
	return kv
}

// Stringify returns the object in JSON format as a string
func (o *CrashPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *CrashPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *CrashPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *CrashPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *CrashPartial) FromMap(kv map[string]interface{}) {
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
	if val, ok := kv["component"].(*string); ok {
		o.Component = val
	} else if val, ok := kv["component"].(string); ok {
		o.Component = &val
	} else {
		if val, ok := kv["component"]; ok {
			if val == nil {
				o.Component = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Component = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if o.CrashDate == nil {
		o.CrashDate = &CrashCrashDate{}
	}

	if val, ok := kv["crash_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.CrashDate.FromMap(kv)
		} else if sv, ok := val.(CrashCrashDate); ok {
			// struct
			o.CrashDate = &sv
		} else if sp, ok := val.(*CrashCrashDate); ok {
			// struct pointer
			o.CrashDate = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.CrashDate.Epoch = dt.Epoch
			o.CrashDate.Rfc3339 = dt.Rfc3339
			o.CrashDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.CrashDate.Epoch = dt.Epoch
			o.CrashDate.Rfc3339 = dt.Rfc3339
			o.CrashDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.CrashDate.Epoch = dt.Epoch
				o.CrashDate.Rfc3339 = dt.Rfc3339
				o.CrashDate.Offset = dt.Offset
			}
		}
	} else {
		o.CrashDate.FromMap(map[string]interface{}{})
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
		o.EventDate = &CrashEventDate{}
	}

	if val, ok := kv["event_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.EventDate.FromMap(kv)
		} else if sv, ok := val.(CrashEventDate); ok {
			// struct
			o.EventDate = &sv
		} else if sp, ok := val.(*CrashEventDate); ok {
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

	if o.LastExportDate == nil {
		o.LastExportDate = &CrashLastExportDate{}
	}

	if val, ok := kv["last_export_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastExportDate.FromMap(kv)
		} else if sv, ok := val.(CrashLastExportDate); ok {
			// struct
			o.LastExportDate = &sv
		} else if sp, ok := val.(*CrashLastExportDate); ok {
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
	if val, ok := kv["type"].(*CrashType); ok {
		o.Type = val
	} else if val, ok := kv["type"].(CrashType); ok {
		o.Type = &val
	} else {
		if val, ok := kv["type"]; ok {
			if val == nil {
				o.Type = toCrashTypePointer(0)
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["CrashType"]
				}
				// this is an enum pointer
				if em, ok := val.(string); ok {
					switch em {
					case "enroll", "ENROLL":
						o.Type = toCrashTypePointer(0)
					case "ping", "PING":
						o.Type = toCrashTypePointer(1)
					case "crash", "CRASH":
						o.Type = toCrashTypePointer(2)
					case "log", "LOG":
						o.Type = toCrashTypePointer(3)
					case "integration", "INTEGRATION":
						o.Type = toCrashTypePointer(4)
					case "export", "EXPORT":
						o.Type = toCrashTypePointer(5)
					case "project", "PROJECT":
						o.Type = toCrashTypePointer(6)
					case "repo", "REPO":
						o.Type = toCrashTypePointer(7)
					case "user", "USER":
						o.Type = toCrashTypePointer(8)
					case "calendar", "CALENDAR":
						o.Type = toCrashTypePointer(9)
					case "uninstall", "UNINSTALL":
						o.Type = toCrashTypePointer(10)
					case "upgrade", "UPGRADE":
						o.Type = toCrashTypePointer(11)
					case "start", "START":
						o.Type = toCrashTypePointer(12)
					case "stop", "STOP":
						o.Type = toCrashTypePointer(13)
					case "pause", "PAUSE":
						o.Type = toCrashTypePointer(14)
					case "resume", "RESUME":
						o.Type = toCrashTypePointer(15)
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
