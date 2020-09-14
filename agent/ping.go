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

	// PingTable is the default table name
	PingTable datamodel.ModelNameType = "agent_ping"

	// PingModelName is the model name
	PingModelName datamodel.ModelNameType = "agent.Ping"
)

const (
	// PingModelArchitectureColumn is the column json value architecture
	PingModelArchitectureColumn = "architecture"
	// PingModelCustomerIDColumn is the column json value customer_id
	PingModelCustomerIDColumn = "customer_id"
	// PingModelDataColumn is the column json value data
	PingModelDataColumn = "data"
	// PingModelDistroColumn is the column json value distro
	PingModelDistroColumn = "distro"
	// PingModelErrorColumn is the column json value error
	PingModelErrorColumn = "error"
	// PingModelEventDateColumn is the column json value event_date
	PingModelEventDateColumn = "event_date"
	// PingModelEventDateEpochColumn is the column json value epoch
	PingModelEventDateEpochColumn = "epoch"
	// PingModelEventDateOffsetColumn is the column json value offset
	PingModelEventDateOffsetColumn = "offset"
	// PingModelEventDateRfc3339Column is the column json value rfc3339
	PingModelEventDateRfc3339Column = "rfc3339"
	// PingModelExportingColumn is the column json value exporting
	PingModelExportingColumn = "exporting"
	// PingModelFreeSpaceColumn is the column json value free_space
	PingModelFreeSpaceColumn = "free_space"
	// PingModelGoVersionColumn is the column json value go_version
	PingModelGoVersionColumn = "go_version"
	// PingModelHostnameColumn is the column json value hostname
	PingModelHostnameColumn = "hostname"
	// PingModelIDColumn is the column json value id
	PingModelIDColumn = "id"
	// PingModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	PingModelIntegrationInstanceIDColumn = "integration_instance_id"
	// PingModelLastExportDateColumn is the column json value last_export_date
	PingModelLastExportDateColumn = "last_export_date"
	// PingModelLastExportDateEpochColumn is the column json value epoch
	PingModelLastExportDateEpochColumn = "epoch"
	// PingModelLastExportDateOffsetColumn is the column json value offset
	PingModelLastExportDateOffsetColumn = "offset"
	// PingModelLastExportDateRfc3339Column is the column json value rfc3339
	PingModelLastExportDateRfc3339Column = "rfc3339"
	// PingModelMemoryColumn is the column json value memory
	PingModelMemoryColumn = "memory"
	// PingModelMessageColumn is the column json value message
	PingModelMessageColumn = "message"
	// PingModelNumCPUColumn is the column json value num_cpu
	PingModelNumCPUColumn = "num_cpu"
	// PingModelOnboardingColumn is the column json value onboarding
	PingModelOnboardingColumn = "onboarding"
	// PingModelOSColumn is the column json value os
	PingModelOSColumn = "os"
	// PingModelRefIDColumn is the column json value ref_id
	PingModelRefIDColumn = "ref_id"
	// PingModelRefTypeColumn is the column json value ref_type
	PingModelRefTypeColumn = "ref_type"
	// PingModelRequestIDColumn is the column json value request_id
	PingModelRequestIDColumn = "request_id"
	// PingModelStateColumn is the column json value state
	PingModelStateColumn = "state"
	// PingModelSuccessColumn is the column json value success
	PingModelSuccessColumn = "success"
	// PingModelSystemIDColumn is the column json value system_id
	PingModelSystemIDColumn = "system_id"
	// PingModelTypeColumn is the column json value type
	PingModelTypeColumn = "type"
	// PingModelUptimeColumn is the column json value uptime
	PingModelUptimeColumn = "uptime"
	// PingModelUUIDColumn is the column json value uuid
	PingModelUUIDColumn = "uuid"
	// PingModelVersionColumn is the column json value version
	PingModelVersionColumn = "version"
)

// PingEventDate represents the object structure for event_date
type PingEventDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toPingEventDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *PingEventDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *PingEventDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toPingEventDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toPingEventDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toPingEventDateObject(o.Rfc3339, false),
	}
}

func (o *PingEventDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *PingEventDate) FromMap(kv map[string]interface{}) {

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

// PingLastExportDate represents the object structure for last_export_date
type PingLastExportDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toPingLastExportDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *PingLastExportDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *PingLastExportDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toPingLastExportDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toPingLastExportDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toPingLastExportDateObject(o.Rfc3339, false),
	}
}

func (o *PingLastExportDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *PingLastExportDate) FromMap(kv map[string]interface{}) {

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

// PingState is the enumeration type for state
type PingState int32

// toPingStatePointer is the enumeration pointer type for state
func toPingStatePointer(v int32) *PingState {
	nv := PingState(v)
	return &nv
}

// toPingStateEnum is the enumeration pointer wrapper for state
func toPingStateEnum(v *PingState) string {
	if v == nil {
		return toPingStatePointer(0).String()
	}
	return v.String()
}

// UnmarshalBSONValue for unmarshaling value
func (v *PingState) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = PingState(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "IDLE":
			*v = PingState(0)
		case "STARTING":
			*v = PingState(1)
		case "STOPPING":
			*v = PingState(2)
		case "EXPORTING":
			*v = PingState(3)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v *PingState) UnmarshalJSON(buf []byte) error {
	var val string
	if err := json.Unmarshal(buf, &val); err != nil {
		return err
	}
	switch val {
	case "IDLE":
		*v = 0
	case "STARTING":
		*v = 1
	case "STOPPING":
		*v = 2
	case "EXPORTING":
		*v = 3
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v PingState) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("IDLE")
	case 1:
		return json.Marshal("STARTING")
	case 2:
		return json.Marshal("STOPPING")
	case 3:
		return json.Marshal("EXPORTING")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for State
func (v PingState) String() string {
	switch int32(v) {
	case 0:
		return "IDLE"
	case 1:
		return "STARTING"
	case 2:
		return "STOPPING"
	case 3:
		return "EXPORTING"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *PingState) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case PingState:
		*v = val
	case int32:
		*v = PingState(int32(val))
	case int:
		*v = PingState(int32(val))
	case string:
		switch val {
		case "IDLE":
			*v = PingState(0)
		case "STARTING":
			*v = PingState(1)
		case "STOPPING":
			*v = PingState(2)
		case "EXPORTING":
			*v = PingState(3)
		}
	}
	return nil
}

const (
	// PingStateIdle is the enumeration value for idle
	PingStateIdle PingState = 0
	// PingStateStarting is the enumeration value for starting
	PingStateStarting PingState = 1
	// PingStateStopping is the enumeration value for stopping
	PingStateStopping PingState = 2
	// PingStateExporting is the enumeration value for exporting
	PingStateExporting PingState = 3
)

// PingType is the enumeration type for type
type PingType int32

// toPingTypePointer is the enumeration pointer type for type
func toPingTypePointer(v int32) *PingType {
	nv := PingType(v)
	return &nv
}

// toPingTypeEnum is the enumeration pointer wrapper for type
func toPingTypeEnum(v *PingType) string {
	if v == nil {
		return toPingTypePointer(0).String()
	}
	return v.String()
}

// UnmarshalBSONValue for unmarshaling value
func (v *PingType) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = PingType(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "ENROLL":
			*v = PingType(0)
		case "PING":
			*v = PingType(1)
		case "CRASH":
			*v = PingType(2)
		case "LOG":
			*v = PingType(3)
		case "INTEGRATION":
			*v = PingType(4)
		case "EXPORT":
			*v = PingType(5)
		case "PROJECT":
			*v = PingType(6)
		case "REPO":
			*v = PingType(7)
		case "USER":
			*v = PingType(8)
		case "CALENDAR":
			*v = PingType(9)
		case "UNINSTALL":
			*v = PingType(10)
		case "UPGRADE":
			*v = PingType(11)
		case "START":
			*v = PingType(12)
		case "STOP":
			*v = PingType(13)
		case "PAUSE":
			*v = PingType(14)
		case "RESUME":
			*v = PingType(15)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v *PingType) UnmarshalJSON(buf []byte) error {
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
func (v PingType) MarshalJSON() ([]byte, error) {
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
func (v PingType) String() string {
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
func (v *PingType) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case PingType:
		*v = val
	case int32:
		*v = PingType(int32(val))
	case int:
		*v = PingType(int32(val))
	case string:
		switch val {
		case "ENROLL":
			*v = PingType(0)
		case "PING":
			*v = PingType(1)
		case "CRASH":
			*v = PingType(2)
		case "LOG":
			*v = PingType(3)
		case "INTEGRATION":
			*v = PingType(4)
		case "EXPORT":
			*v = PingType(5)
		case "PROJECT":
			*v = PingType(6)
		case "REPO":
			*v = PingType(7)
		case "USER":
			*v = PingType(8)
		case "CALENDAR":
			*v = PingType(9)
		case "UNINSTALL":
			*v = PingType(10)
		case "UPGRADE":
			*v = PingType(11)
		case "START":
			*v = PingType(12)
		case "STOP":
			*v = PingType(13)
		case "PAUSE":
			*v = PingType(14)
		case "RESUME":
			*v = PingType(15)
		}
	}
	return nil
}

const (
	// PingTypeEnroll is the enumeration value for enroll
	PingTypeEnroll PingType = 0
	// PingTypePing is the enumeration value for ping
	PingTypePing PingType = 1
	// PingTypeCrash is the enumeration value for crash
	PingTypeCrash PingType = 2
	// PingTypeLog is the enumeration value for log
	PingTypeLog PingType = 3
	// PingTypeIntegration is the enumeration value for integration
	PingTypeIntegration PingType = 4
	// PingTypeExport is the enumeration value for export
	PingTypeExport PingType = 5
	// PingTypeProject is the enumeration value for project
	PingTypeProject PingType = 6
	// PingTypeRepo is the enumeration value for repo
	PingTypeRepo PingType = 7
	// PingTypeUser is the enumeration value for user
	PingTypeUser PingType = 8
	// PingTypeCalendar is the enumeration value for calendar
	PingTypeCalendar PingType = 9
	// PingTypeUninstall is the enumeration value for uninstall
	PingTypeUninstall PingType = 10
	// PingTypeUpgrade is the enumeration value for upgrade
	PingTypeUpgrade PingType = 11
	// PingTypeStart is the enumeration value for start
	PingTypeStart PingType = 12
	// PingTypeStop is the enumeration value for stop
	PingTypeStop PingType = 13
	// PingTypePause is the enumeration value for pause
	PingTypePause PingType = 14
	// PingTypeResume is the enumeration value for resume
	PingTypeResume PingType = 15
)

// Ping an agent event for recording a ping
type Ping struct {
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
	EventDate PingEventDate `json:"event_date" codec:"event_date" bson:"event_date" yaml:"event_date" faker:"-"`
	// Exporting flag to specify if export is in progress
	Exporting bool `json:"exporting" codec:"exporting" bson:"exporting" yaml:"exporting" faker:"-"`
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
	LastExportDate PingLastExportDate `json:"last_export_date" codec:"last_export_date" bson:"last_export_date" yaml:"last_export_date" faker:"-"`
	// Memory the amount of memory in bytes for the agent machine
	Memory int64 `json:"memory" codec:"memory" bson:"memory" yaml:"memory" faker:"-"`
	// Message a message related to this event
	Message string `json:"message" codec:"message" bson:"message" yaml:"message" faker:"-"`
	// NumCPU the number of CPU the agent is running
	NumCPU int64 `json:"num_cpu" codec:"num_cpu" bson:"num_cpu" yaml:"num_cpu" faker:"-"`
	// Onboarding flag to specify if onboarding events are in progress. could be multiple at once
	Onboarding bool `json:"onboarding" codec:"onboarding" bson:"onboarding" yaml:"onboarding" faker:"-"`
	// OS the agent operating system
	OS string `json:"os" codec:"os" bson:"os" yaml:"os" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RequestID the request id that this response is correlated to
	RequestID string `json:"request_id" codec:"request_id" bson:"request_id" yaml:"request_id" faker:"-"`
	// State the state of the agent
	State PingState `json:"state" codec:"state" bson:"state" yaml:"state" faker:"-"`
	// Success if the response was successful
	Success bool `json:"success" codec:"success" bson:"success" yaml:"success" faker:"-"`
	// SystemID system unique device ID
	SystemID string `json:"system_id" codec:"system_id" bson:"system_id" yaml:"system_id" faker:"-"`
	// Type the type of event
	Type PingType `json:"type" codec:"type" bson:"type" yaml:"type" faker:"-"`
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
var _ datamodel.Model = (*Ping)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*Ping)(nil)

func toPingObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *Ping:
		return v.ToMap()

	case PingEventDate:
		return v.ToMap()

	case PingLastExportDate:
		return v.ToMap()

	case PingState:
		return v.String()

	case PingType:
		return v.String()

	default:
		return o
	}
}

// String returns a string representation of Ping
func (o *Ping) String() string {
	return fmt.Sprintf("agent.Ping<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Ping) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *Ping) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *Ping) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *Ping) GetModelName() datamodel.ModelNameType {
	return PingModelName
}

// NewPingID provides a template for generating an ID field for Ping
func NewPingID(customerID string, refType string, refID string) string {
	return hash.Values("Ping", customerID, refType, refID)
}

func (o *Ping) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Ping", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Ping) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Ping) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Ping) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *Ping) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Ping) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *Ping) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Ping) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Ping) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *Ping) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = PingModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Ping) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *Ping) GetCustomerID() string {

	return o.CustomerID

}

// SetCustomerID will return the customer_id
func (o *Ping) SetCustomerID(id string) {

	o.CustomerID = id

}

// GetRefType will return the ref_type
func (o *Ping) GetRefType() string {

	return o.RefType

}

// SetRefType will return the ref_type
func (o *Ping) SetRefType(t string) {

	o.RefType = id

}

// Clone returns an exact copy of Ping
func (o *Ping) Clone() datamodel.Model {
	c := new(Ping)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Ping) Anon() datamodel.Model {
	c := new(Ping)
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
func (o *Ping) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Ping) UnmarshalJSON(data []byte) error {
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
func (o *Ping) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *Ping) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two Ping objects are equal
func (o *Ping) IsEqual(other *Ping) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Ping) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"architecture":            toPingObject(o.Architecture, false),
		"customer_id":             toPingObject(o.CustomerID, false),
		"data":                    toPingObject(o.Data, true),
		"distro":                  toPingObject(o.Distro, false),
		"error":                   toPingObject(o.Error, true),
		"event_date":              toPingObject(o.EventDate, false),
		"exporting":               toPingObject(o.Exporting, false),
		"free_space":              toPingObject(o.FreeSpace, false),
		"go_version":              toPingObject(o.GoVersion, false),
		"hostname":                toPingObject(o.Hostname, false),
		"id":                      toPingObject(o.ID, false),
		"integration_instance_id": toPingObject(o.IntegrationInstanceID, true),
		"last_export_date":        toPingObject(o.LastExportDate, false),
		"memory":                  toPingObject(o.Memory, false),
		"message":                 toPingObject(o.Message, false),
		"num_cpu":                 toPingObject(o.NumCPU, false),
		"onboarding":              toPingObject(o.Onboarding, false),
		"os":                      toPingObject(o.OS, false),
		"ref_id":                  toPingObject(o.RefID, false),
		"ref_type":                toPingObject(o.RefType, false),
		"request_id":              toPingObject(o.RequestID, false),

		"state":     o.State.String(),
		"success":   toPingObject(o.Success, false),
		"system_id": toPingObject(o.SystemID, false),

		"type":     o.Type.String(),
		"uptime":   toPingObject(o.Uptime, false),
		"uuid":     toPingObject(o.UUID, false),
		"version":  toPingObject(o.Version, false),
		"hashcode": toPingObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *Ping) FromMap(kv map[string]interface{}) {

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
		} else if sv, ok := val.(PingEventDate); ok {
			// struct
			o.EventDate = sv
		} else if sp, ok := val.(*PingEventDate); ok {
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

	if val, ok := kv["exporting"].(bool); ok {
		o.Exporting = val
	} else {
		if val, ok := kv["exporting"]; ok {
			if val == nil {
				o.Exporting = false
			} else {
				o.Exporting = number.ToBoolAny(val)
			}
		}
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
		} else if sv, ok := val.(PingLastExportDate); ok {
			// struct
			o.LastExportDate = sv
		} else if sp, ok := val.(*PingLastExportDate); ok {
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
	if val, ok := kv["onboarding"].(bool); ok {
		o.Onboarding = val
	} else {
		if val, ok := kv["onboarding"]; ok {
			if val == nil {
				o.Onboarding = false
			} else {
				o.Onboarding = number.ToBoolAny(val)
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
	if val, ok := kv["state"].(PingState); ok {
		o.State = val
	} else {
		if em, ok := kv["state"].(map[string]interface{}); ok {

			ev := em["agent.state"].(string)
			switch ev {
			case "idle", "IDLE":
				o.State = 0
			case "starting", "STARTING":
				o.State = 1
			case "stopping", "STOPPING":
				o.State = 2
			case "exporting", "EXPORTING":
				o.State = 3
			}
		}
		if em, ok := kv["state"].(string); ok {
			switch em {
			case "idle", "IDLE":
				o.State = 0
			case "starting", "STARTING":
				o.State = 1
			case "stopping", "STOPPING":
				o.State = 2
			case "exporting", "EXPORTING":
				o.State = 3
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
	if val, ok := kv["type"].(PingType); ok {
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
func (o *Ping) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Architecture)
	args = append(args, o.CustomerID)
	args = append(args, o.Data)
	args = append(args, o.Distro)
	args = append(args, o.Error)
	args = append(args, o.EventDate)
	args = append(args, o.Exporting)
	args = append(args, o.FreeSpace)
	args = append(args, o.GoVersion)
	args = append(args, o.Hostname)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.LastExportDate)
	args = append(args, o.Memory)
	args = append(args, o.Message)
	args = append(args, o.NumCPU)
	args = append(args, o.Onboarding)
	args = append(args, o.OS)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RequestID)
	args = append(args, o.State)
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
func (o *Ping) SetIntegrationInstanceID(id string) {
	if id == "" {
		o.IntegrationInstanceID = nil
	} else {
		o.IntegrationInstanceID = &id
	}
}

// GetIntegrationInstanceID will return the integration instance ID
func (o *Ping) GetIntegrationInstanceID() *string {
	return o.IntegrationInstanceID
}

// PingPartial is a partial struct for upsert mutations for Ping
type PingPartial struct {
	// Architecture the architecture of the agent machine
	Architecture *string `json:"architecture,omitempty"`
	// Data extra data that is specific about this event
	Data *string `json:"data,omitempty"`
	// Distro the agent os distribution
	Distro *string `json:"distro,omitempty"`
	// Error an error message related to this event
	Error *string `json:"error,omitempty"`
	// EventDate the date of the event
	EventDate *PingEventDate `json:"event_date,omitempty"`
	// Exporting flag to specify if export is in progress
	Exporting *bool `json:"exporting,omitempty"`
	// FreeSpace the amount of free space in bytes for the agent machine
	FreeSpace *int64 `json:"free_space,omitempty"`
	// GoVersion the go version that the agent build was built with
	GoVersion *string `json:"go_version,omitempty"`
	// Hostname the agent hostname
	Hostname *string `json:"hostname,omitempty"`
	// LastExportDate the last export date
	LastExportDate *PingLastExportDate `json:"last_export_date,omitempty"`
	// Memory the amount of memory in bytes for the agent machine
	Memory *int64 `json:"memory,omitempty"`
	// Message a message related to this event
	Message *string `json:"message,omitempty"`
	// NumCPU the number of CPU the agent is running
	NumCPU *int64 `json:"num_cpu,omitempty"`
	// Onboarding flag to specify if onboarding events are in progress. could be multiple at once
	Onboarding *bool `json:"onboarding,omitempty"`
	// OS the agent operating system
	OS *string `json:"os,omitempty"`
	// RequestID the request id that this response is correlated to
	RequestID *string `json:"request_id,omitempty"`
	// State the state of the agent
	State *PingState `json:"state,omitempty"`
	// Success if the response was successful
	Success *bool `json:"success,omitempty"`
	// SystemID system unique device ID
	SystemID *string `json:"system_id,omitempty"`
	// Type the type of event
	Type *PingType `json:"type,omitempty"`
	// Uptime the uptime in milliseconds since the agent started
	Uptime *int64 `json:"uptime,omitempty"`
	// UUID the agent unique identifier
	UUID *string `json:"uuid,omitempty"`
	// Version the agent version
	Version *string `json:"version,omitempty"`
}

var _ datamodel.PartialModel = (*PingPartial)(nil)

// GetModelName returns the name of the model
func (o *PingPartial) GetModelName() datamodel.ModelNameType {
	return PingModelName
}

// ToMap returns the object as a map
func (o *PingPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"architecture":     toPingObject(o.Architecture, true),
		"data":             toPingObject(o.Data, true),
		"distro":           toPingObject(o.Distro, true),
		"error":            toPingObject(o.Error, true),
		"event_date":       toPingObject(o.EventDate, true),
		"exporting":        toPingObject(o.Exporting, true),
		"free_space":       toPingObject(o.FreeSpace, true),
		"go_version":       toPingObject(o.GoVersion, true),
		"hostname":         toPingObject(o.Hostname, true),
		"last_export_date": toPingObject(o.LastExportDate, true),
		"memory":           toPingObject(o.Memory, true),
		"message":          toPingObject(o.Message, true),
		"num_cpu":          toPingObject(o.NumCPU, true),
		"onboarding":       toPingObject(o.Onboarding, true),
		"os":               toPingObject(o.OS, true),
		"request_id":       toPingObject(o.RequestID, true),

		"state":     toPingStateEnum(o.State),
		"success":   toPingObject(o.Success, true),
		"system_id": toPingObject(o.SystemID, true),

		"type":    toPingTypeEnum(o.Type),
		"uptime":  toPingObject(o.Uptime, true),
		"uuid":    toPingObject(o.UUID, true),
		"version": toPingObject(o.Version, true),
	}
	for k, v := range kv {
		if v == nil || reflect.ValueOf(v).IsZero() {
			delete(kv, k)
		} else {
			if k == "event_date" {
				if dt, ok := v.(*PingEventDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}
			if k == "last_export_date" {
				if dt, ok := v.(*PingLastExportDate); ok {
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
func (o *PingPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *PingPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *PingPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *PingPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *PingPartial) FromMap(kv map[string]interface{}) {
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
		o.EventDate = &PingEventDate{}
	}

	if val, ok := kv["event_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.EventDate.FromMap(kv)
		} else if sv, ok := val.(PingEventDate); ok {
			// struct
			o.EventDate = &sv
		} else if sp, ok := val.(*PingEventDate); ok {
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

	if val, ok := kv["exporting"].(*bool); ok {
		o.Exporting = val
	} else if val, ok := kv["exporting"].(bool); ok {
		o.Exporting = &val
	} else {
		if val, ok := kv["exporting"]; ok {
			if val == nil {
				o.Exporting = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Exporting = number.BoolPointer(number.ToBoolAny(val))
			}
		}
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
		o.LastExportDate = &PingLastExportDate{}
	}

	if val, ok := kv["last_export_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastExportDate.FromMap(kv)
		} else if sv, ok := val.(PingLastExportDate); ok {
			// struct
			o.LastExportDate = &sv
		} else if sp, ok := val.(*PingLastExportDate); ok {
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
	if val, ok := kv["onboarding"].(*bool); ok {
		o.Onboarding = val
	} else if val, ok := kv["onboarding"].(bool); ok {
		o.Onboarding = &val
	} else {
		if val, ok := kv["onboarding"]; ok {
			if val == nil {
				o.Onboarding = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Onboarding = number.BoolPointer(number.ToBoolAny(val))
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
	if val, ok := kv["state"].(*PingState); ok {
		o.State = val
	} else if val, ok := kv["state"].(PingState); ok {
		o.State = &val
	} else {
		if val, ok := kv["state"]; ok {
			if val == nil {
				o.State = toPingStatePointer(0)
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["PingState"]
				}
				// this is an enum pointer
				if em, ok := val.(string); ok {
					switch em {
					case "idle", "IDLE":
						o.State = toPingStatePointer(0)
					case "starting", "STARTING":
						o.State = toPingStatePointer(1)
					case "stopping", "STOPPING":
						o.State = toPingStatePointer(2)
					case "exporting", "EXPORTING":
						o.State = toPingStatePointer(3)
					}
				}
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
	if val, ok := kv["type"].(*PingType); ok {
		o.Type = val
	} else if val, ok := kv["type"].(PingType); ok {
		o.Type = &val
	} else {
		if val, ok := kv["type"]; ok {
			if val == nil {
				o.Type = toPingTypePointer(0)
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["PingType"]
				}
				// this is an enum pointer
				if em, ok := val.(string); ok {
					switch em {
					case "enroll", "ENROLL":
						o.Type = toPingTypePointer(0)
					case "ping", "PING":
						o.Type = toPingTypePointer(1)
					case "crash", "CRASH":
						o.Type = toPingTypePointer(2)
					case "log", "LOG":
						o.Type = toPingTypePointer(3)
					case "integration", "INTEGRATION":
						o.Type = toPingTypePointer(4)
					case "export", "EXPORT":
						o.Type = toPingTypePointer(5)
					case "project", "PROJECT":
						o.Type = toPingTypePointer(6)
					case "repo", "REPO":
						o.Type = toPingTypePointer(7)
					case "user", "USER":
						o.Type = toPingTypePointer(8)
					case "calendar", "CALENDAR":
						o.Type = toPingTypePointer(9)
					case "uninstall", "UNINSTALL":
						o.Type = toPingTypePointer(10)
					case "upgrade", "UPGRADE":
						o.Type = toPingTypePointer(11)
					case "start", "START":
						o.Type = toPingTypePointer(12)
					case "stop", "STOP":
						o.Type = toPingTypePointer(13)
					case "pause", "PAUSE":
						o.Type = toPingTypePointer(14)
					case "resume", "RESUME":
						o.Type = toPingTypePointer(15)
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
