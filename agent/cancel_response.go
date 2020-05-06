// DO NOT EDIT -- generated code

// Package agent - the agent communicates with pinpoint cloud to send integration data
package agent

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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

const (

	// CancelResponseTable is the default table name
	CancelResponseTable datamodel.ModelNameType = "agent_cancelresponse"

	// CancelResponseModelName is the model name
	CancelResponseModelName datamodel.ModelNameType = "agent.CancelResponse"
)

const (
	// CancelResponseModelArchitectureColumn is the column json value architecture
	CancelResponseModelArchitectureColumn = "architecture"
	// CancelResponseModelCancelDateColumn is the column json value cancel_date
	CancelResponseModelCancelDateColumn = "cancel_date"
	// CancelResponseModelCancelDateEpochColumn is the column json value epoch
	CancelResponseModelCancelDateEpochColumn = "epoch"
	// CancelResponseModelCancelDateOffsetColumn is the column json value offset
	CancelResponseModelCancelDateOffsetColumn = "offset"
	// CancelResponseModelCancelDateRfc3339Column is the column json value rfc3339
	CancelResponseModelCancelDateRfc3339Column = "rfc3339"
	// CancelResponseModelCustomerIDColumn is the column json value customer_id
	CancelResponseModelCustomerIDColumn = "customer_id"
	// CancelResponseModelDataColumn is the column json value data
	CancelResponseModelDataColumn = "data"
	// CancelResponseModelDistroColumn is the column json value distro
	CancelResponseModelDistroColumn = "distro"
	// CancelResponseModelErrorColumn is the column json value error
	CancelResponseModelErrorColumn = "error"
	// CancelResponseModelEventDateColumn is the column json value event_date
	CancelResponseModelEventDateColumn = "event_date"
	// CancelResponseModelEventDateEpochColumn is the column json value epoch
	CancelResponseModelEventDateEpochColumn = "epoch"
	// CancelResponseModelEventDateOffsetColumn is the column json value offset
	CancelResponseModelEventDateOffsetColumn = "offset"
	// CancelResponseModelEventDateRfc3339Column is the column json value rfc3339
	CancelResponseModelEventDateRfc3339Column = "rfc3339"
	// CancelResponseModelFreeSpaceColumn is the column json value free_space
	CancelResponseModelFreeSpaceColumn = "free_space"
	// CancelResponseModelGoVersionColumn is the column json value go_version
	CancelResponseModelGoVersionColumn = "go_version"
	// CancelResponseModelHostnameColumn is the column json value hostname
	CancelResponseModelHostnameColumn = "hostname"
	// CancelResponseModelIDColumn is the column json value id
	CancelResponseModelIDColumn = "id"
	// CancelResponseModelLastExportDateColumn is the column json value last_export_date
	CancelResponseModelLastExportDateColumn = "last_export_date"
	// CancelResponseModelLastExportDateEpochColumn is the column json value epoch
	CancelResponseModelLastExportDateEpochColumn = "epoch"
	// CancelResponseModelLastExportDateOffsetColumn is the column json value offset
	CancelResponseModelLastExportDateOffsetColumn = "offset"
	// CancelResponseModelLastExportDateRfc3339Column is the column json value rfc3339
	CancelResponseModelLastExportDateRfc3339Column = "rfc3339"
	// CancelResponseModelMemoryColumn is the column json value memory
	CancelResponseModelMemoryColumn = "memory"
	// CancelResponseModelMessageColumn is the column json value message
	CancelResponseModelMessageColumn = "message"
	// CancelResponseModelNumCPUColumn is the column json value num_cpu
	CancelResponseModelNumCPUColumn = "num_cpu"
	// CancelResponseModelOSColumn is the column json value os
	CancelResponseModelOSColumn = "os"
	// CancelResponseModelRefIDColumn is the column json value ref_id
	CancelResponseModelRefIDColumn = "ref_id"
	// CancelResponseModelRefTypeColumn is the column json value ref_type
	CancelResponseModelRefTypeColumn = "ref_type"
	// CancelResponseModelRequestIDColumn is the column json value request_id
	CancelResponseModelRequestIDColumn = "request_id"
	// CancelResponseModelSuccessColumn is the column json value success
	CancelResponseModelSuccessColumn = "success"
	// CancelResponseModelSystemIDColumn is the column json value system_id
	CancelResponseModelSystemIDColumn = "system_id"
	// CancelResponseModelTypeColumn is the column json value type
	CancelResponseModelTypeColumn = "type"
	// CancelResponseModelUptimeColumn is the column json value uptime
	CancelResponseModelUptimeColumn = "uptime"
	// CancelResponseModelUUIDColumn is the column json value uuid
	CancelResponseModelUUIDColumn = "uuid"
	// CancelResponseModelVersionColumn is the column json value version
	CancelResponseModelVersionColumn = "version"
)

// CancelResponseCancelDate represents the object structure for cancel_date
type CancelResponseCancelDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toCancelResponseCancelDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CancelResponseCancelDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *CancelResponseCancelDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toCancelResponseCancelDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toCancelResponseCancelDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toCancelResponseCancelDateObject(o.Rfc3339, false),
	}
}

func (o *CancelResponseCancelDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *CancelResponseCancelDate) FromMap(kv map[string]interface{}) {

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

// CancelResponseEventDate represents the object structure for event_date
type CancelResponseEventDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toCancelResponseEventDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CancelResponseEventDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *CancelResponseEventDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toCancelResponseEventDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toCancelResponseEventDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toCancelResponseEventDateObject(o.Rfc3339, false),
	}
}

func (o *CancelResponseEventDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *CancelResponseEventDate) FromMap(kv map[string]interface{}) {

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

// CancelResponseLastExportDate represents the object structure for last_export_date
type CancelResponseLastExportDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toCancelResponseLastExportDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CancelResponseLastExportDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *CancelResponseLastExportDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toCancelResponseLastExportDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toCancelResponseLastExportDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toCancelResponseLastExportDateObject(o.Rfc3339, false),
	}
}

func (o *CancelResponseLastExportDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *CancelResponseLastExportDate) FromMap(kv map[string]interface{}) {

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

// CancelResponseType is the enumeration type for type
type CancelResponseType int32

// UnmarshalBSONValue for unmarshaling value
func (v *CancelResponseType) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = CancelResponseType(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "ENROLL":
			*v = CancelResponseType(0)
		case "PING":
			*v = CancelResponseType(1)
		case "CRASH":
			*v = CancelResponseType(2)
		case "LOG":
			*v = CancelResponseType(3)
		case "INTEGRATION":
			*v = CancelResponseType(4)
		case "EXPORT":
			*v = CancelResponseType(5)
		case "PROJECT":
			*v = CancelResponseType(6)
		case "REPO":
			*v = CancelResponseType(7)
		case "USER":
			*v = CancelResponseType(8)
		case "CALENDAR":
			*v = CancelResponseType(9)
		case "UNINSTALL":
			*v = CancelResponseType(10)
		case "UPGRADE":
			*v = CancelResponseType(11)
		case "START":
			*v = CancelResponseType(12)
		case "STOP":
			*v = CancelResponseType(13)
		case "PAUSE":
			*v = CancelResponseType(14)
		case "RESUME":
			*v = CancelResponseType(15)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v CancelResponseType) UnmarshalJSON(buf []byte) error {
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
func (v CancelResponseType) MarshalJSON() ([]byte, error) {
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
func (v CancelResponseType) String() string {
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
func (v *CancelResponseType) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case CancelResponseType:
		*v = val
	case int32:
		*v = CancelResponseType(int32(val))
	case int:
		*v = CancelResponseType(int32(val))
	case string:
		switch val {
		case "ENROLL":
			*v = CancelResponseType(0)
		case "PING":
			*v = CancelResponseType(1)
		case "CRASH":
			*v = CancelResponseType(2)
		case "LOG":
			*v = CancelResponseType(3)
		case "INTEGRATION":
			*v = CancelResponseType(4)
		case "EXPORT":
			*v = CancelResponseType(5)
		case "PROJECT":
			*v = CancelResponseType(6)
		case "REPO":
			*v = CancelResponseType(7)
		case "USER":
			*v = CancelResponseType(8)
		case "CALENDAR":
			*v = CancelResponseType(9)
		case "UNINSTALL":
			*v = CancelResponseType(10)
		case "UPGRADE":
			*v = CancelResponseType(11)
		case "START":
			*v = CancelResponseType(12)
		case "STOP":
			*v = CancelResponseType(13)
		case "PAUSE":
			*v = CancelResponseType(14)
		case "RESUME":
			*v = CancelResponseType(15)
		}
	}
	return nil
}

const (
	// CancelResponseTypeEnroll is the enumeration value for enroll
	CancelResponseTypeEnroll CancelResponseType = 0
	// CancelResponseTypePing is the enumeration value for ping
	CancelResponseTypePing CancelResponseType = 1
	// CancelResponseTypeCrash is the enumeration value for crash
	CancelResponseTypeCrash CancelResponseType = 2
	// CancelResponseTypeLog is the enumeration value for log
	CancelResponseTypeLog CancelResponseType = 3
	// CancelResponseTypeIntegration is the enumeration value for integration
	CancelResponseTypeIntegration CancelResponseType = 4
	// CancelResponseTypeExport is the enumeration value for export
	CancelResponseTypeExport CancelResponseType = 5
	// CancelResponseTypeProject is the enumeration value for project
	CancelResponseTypeProject CancelResponseType = 6
	// CancelResponseTypeRepo is the enumeration value for repo
	CancelResponseTypeRepo CancelResponseType = 7
	// CancelResponseTypeUser is the enumeration value for user
	CancelResponseTypeUser CancelResponseType = 8
	// CancelResponseTypeCalendar is the enumeration value for calendar
	CancelResponseTypeCalendar CancelResponseType = 9
	// CancelResponseTypeUninstall is the enumeration value for uninstall
	CancelResponseTypeUninstall CancelResponseType = 10
	// CancelResponseTypeUpgrade is the enumeration value for upgrade
	CancelResponseTypeUpgrade CancelResponseType = 11
	// CancelResponseTypeStart is the enumeration value for start
	CancelResponseTypeStart CancelResponseType = 12
	// CancelResponseTypeStop is the enumeration value for stop
	CancelResponseTypeStop CancelResponseType = 13
	// CancelResponseTypePause is the enumeration value for pause
	CancelResponseTypePause CancelResponseType = 14
	// CancelResponseTypeResume is the enumeration value for resume
	CancelResponseTypeResume CancelResponseType = 15
)

// CancelResponse an agent response to an action request for cancel
type CancelResponse struct {
	// Architecture the architecture of the agent machine
	Architecture string `json:"architecture" codec:"architecture" bson:"architecture" yaml:"architecture" faker:"-"`
	// CancelDate the cancel date
	CancelDate CancelResponseCancelDate `json:"cancel_date" codec:"cancel_date" bson:"cancel_date" yaml:"cancel_date" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Data extra data that is specific about this event
	Data *string `json:"data,omitempty" codec:"data,omitempty" bson:"data" yaml:"data,omitempty" faker:"-"`
	// Distro the agent os distribution
	Distro string `json:"distro" codec:"distro" bson:"distro" yaml:"distro" faker:"-"`
	// Error an error message related to this event
	Error *string `json:"error,omitempty" codec:"error,omitempty" bson:"error" yaml:"error,omitempty" faker:"-"`
	// EventDate the date of the event
	EventDate CancelResponseEventDate `json:"event_date" codec:"event_date" bson:"event_date" yaml:"event_date" faker:"-"`
	// FreeSpace the amount of free space in bytes for the agent machine
	FreeSpace int64 `json:"free_space" codec:"free_space" bson:"free_space" yaml:"free_space" faker:"-"`
	// GoVersion the go version that the agent build was built with
	GoVersion string `json:"go_version" codec:"go_version" bson:"go_version" yaml:"go_version" faker:"-"`
	// Hostname the agent hostname
	Hostname string `json:"hostname" codec:"hostname" bson:"hostname" yaml:"hostname" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// LastExportDate the last export date
	LastExportDate CancelResponseLastExportDate `json:"last_export_date" codec:"last_export_date" bson:"last_export_date" yaml:"last_export_date" faker:"-"`
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
	Type CancelResponseType `json:"type" codec:"type" bson:"type" yaml:"type" faker:"-"`
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
var _ datamodel.Model = (*CancelResponse)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*CancelResponse)(nil)

func toCancelResponseObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CancelResponse:
		return v.ToMap()

	case CancelResponseCancelDate:
		return v.ToMap()

	case CancelResponseEventDate:
		return v.ToMap()

	case CancelResponseLastExportDate:
		return v.ToMap()

	case CancelResponseType:
		return v.String()

	default:
		return o
	}
}

// String returns a string representation of CancelResponse
func (o *CancelResponse) String() string {
	return fmt.Sprintf("agent.CancelResponse<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *CancelResponse) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *CancelResponse) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *CancelResponse) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *CancelResponse) GetModelName() datamodel.ModelNameType {
	return CancelResponseModelName
}

// NewCancelResponseID provides a template for generating an ID field for CancelResponse
func NewCancelResponseID(customerID string, refType string, refID string) string {
	return hash.Values("CancelResponse", customerID, refType, refID)
}

func (o *CancelResponse) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("CancelResponse", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *CancelResponse) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *CancelResponse) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *CancelResponse) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *CancelResponse) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *CancelResponse) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *CancelResponse) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *CancelResponse) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *CancelResponse) IsEvented() bool {
	return false
}

// GetTopicConfig returns the topic config object
func (o *CancelResponse) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *CancelResponse) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of CancelResponse
func (o *CancelResponse) Clone() datamodel.Model {
	c := new(CancelResponse)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *CancelResponse) Anon() datamodel.Model {
	c := new(CancelResponse)
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
func (o *CancelResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *CancelResponse) UnmarshalJSON(data []byte) error {
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
func (o *CancelResponse) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *CancelResponse) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two CancelResponse objects are equal
func (o *CancelResponse) IsEqual(other *CancelResponse) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *CancelResponse) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"architecture":     toCancelResponseObject(o.Architecture, false),
		"cancel_date":      toCancelResponseObject(o.CancelDate, false),
		"customer_id":      toCancelResponseObject(o.CustomerID, false),
		"data":             toCancelResponseObject(o.Data, true),
		"distro":           toCancelResponseObject(o.Distro, false),
		"error":            toCancelResponseObject(o.Error, true),
		"event_date":       toCancelResponseObject(o.EventDate, false),
		"free_space":       toCancelResponseObject(o.FreeSpace, false),
		"go_version":       toCancelResponseObject(o.GoVersion, false),
		"hostname":         toCancelResponseObject(o.Hostname, false),
		"id":               toCancelResponseObject(o.ID, false),
		"last_export_date": toCancelResponseObject(o.LastExportDate, false),
		"memory":           toCancelResponseObject(o.Memory, false),
		"message":          toCancelResponseObject(o.Message, false),
		"num_cpu":          toCancelResponseObject(o.NumCPU, false),
		"os":               toCancelResponseObject(o.OS, false),
		"ref_id":           toCancelResponseObject(o.RefID, false),
		"ref_type":         toCancelResponseObject(o.RefType, false),
		"request_id":       toCancelResponseObject(o.RequestID, false),
		"success":          toCancelResponseObject(o.Success, false),
		"system_id":        toCancelResponseObject(o.SystemID, false),

		"type":     o.Type.String(),
		"uptime":   toCancelResponseObject(o.Uptime, false),
		"uuid":     toCancelResponseObject(o.UUID, false),
		"version":  toCancelResponseObject(o.Version, false),
		"hashcode": toCancelResponseObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *CancelResponse) FromMap(kv map[string]interface{}) {

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

	if val, ok := kv["cancel_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.CancelDate.FromMap(kv)
		} else if sv, ok := val.(CancelResponseCancelDate); ok {
			// struct
			o.CancelDate = sv
		} else if sp, ok := val.(*CancelResponseCancelDate); ok {
			// struct pointer
			o.CancelDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.CancelDate.Epoch = dt.Epoch
			o.CancelDate.Rfc3339 = dt.Rfc3339
			o.CancelDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.CancelDate.Epoch = dt.Epoch
			o.CancelDate.Rfc3339 = dt.Rfc3339
			o.CancelDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.CancelDate.Epoch = dt.Epoch
				o.CancelDate.Rfc3339 = dt.Rfc3339
				o.CancelDate.Offset = dt.Offset
			}
		}
	} else {
		o.CancelDate.FromMap(map[string]interface{}{})
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
		} else if sv, ok := val.(CancelResponseEventDate); ok {
			// struct
			o.EventDate = sv
		} else if sp, ok := val.(*CancelResponseEventDate); ok {
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

	if val, ok := kv["last_export_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastExportDate.FromMap(kv)
		} else if sv, ok := val.(CancelResponseLastExportDate); ok {
			// struct
			o.LastExportDate = sv
		} else if sp, ok := val.(*CancelResponseLastExportDate); ok {
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

	if val, ok := kv["type"].(CancelResponseType); ok {
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
func (o *CancelResponse) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Architecture)
	args = append(args, o.CancelDate)
	args = append(args, o.CustomerID)
	args = append(args, o.Data)
	args = append(args, o.Distro)
	args = append(args, o.Error)
	args = append(args, o.EventDate)
	args = append(args, o.FreeSpace)
	args = append(args, o.GoVersion)
	args = append(args, o.Hostname)
	args = append(args, o.ID)
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
func (o *CancelResponse) GetEventAPIConfig() datamodel.EventAPIConfig {
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
