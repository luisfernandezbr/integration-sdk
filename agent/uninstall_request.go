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

	// UninstallRequestTable is the default table name
	UninstallRequestTable datamodel.ModelNameType = "agent_uninstallrequest"

	// UninstallRequestModelName is the model name
	UninstallRequestModelName datamodel.ModelNameType = "agent.UninstallRequest"
)

const (
	// UninstallRequestModelArchitectureColumn is the column json value architecture
	UninstallRequestModelArchitectureColumn = "architecture"
	// UninstallRequestModelCustomerIDColumn is the column json value customer_id
	UninstallRequestModelCustomerIDColumn = "customer_id"
	// UninstallRequestModelDataColumn is the column json value data
	UninstallRequestModelDataColumn = "data"
	// UninstallRequestModelDistroColumn is the column json value distro
	UninstallRequestModelDistroColumn = "distro"
	// UninstallRequestModelErrorColumn is the column json value error
	UninstallRequestModelErrorColumn = "error"
	// UninstallRequestModelEventDateColumn is the column json value event_date
	UninstallRequestModelEventDateColumn = "event_date"
	// UninstallRequestModelEventDateEpochColumn is the column json value epoch
	UninstallRequestModelEventDateEpochColumn = "epoch"
	// UninstallRequestModelEventDateOffsetColumn is the column json value offset
	UninstallRequestModelEventDateOffsetColumn = "offset"
	// UninstallRequestModelEventDateRfc3339Column is the column json value rfc3339
	UninstallRequestModelEventDateRfc3339Column = "rfc3339"
	// UninstallRequestModelFreeSpaceColumn is the column json value free_space
	UninstallRequestModelFreeSpaceColumn = "free_space"
	// UninstallRequestModelGoVersionColumn is the column json value go_version
	UninstallRequestModelGoVersionColumn = "go_version"
	// UninstallRequestModelHostnameColumn is the column json value hostname
	UninstallRequestModelHostnameColumn = "hostname"
	// UninstallRequestModelIDColumn is the column json value id
	UninstallRequestModelIDColumn = "id"
	// UninstallRequestModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	UninstallRequestModelIntegrationInstanceIDColumn = "integration_instance_id"
	// UninstallRequestModelLastExportDateColumn is the column json value last_export_date
	UninstallRequestModelLastExportDateColumn = "last_export_date"
	// UninstallRequestModelLastExportDateEpochColumn is the column json value epoch
	UninstallRequestModelLastExportDateEpochColumn = "epoch"
	// UninstallRequestModelLastExportDateOffsetColumn is the column json value offset
	UninstallRequestModelLastExportDateOffsetColumn = "offset"
	// UninstallRequestModelLastExportDateRfc3339Column is the column json value rfc3339
	UninstallRequestModelLastExportDateRfc3339Column = "rfc3339"
	// UninstallRequestModelMemoryColumn is the column json value memory
	UninstallRequestModelMemoryColumn = "memory"
	// UninstallRequestModelMessageColumn is the column json value message
	UninstallRequestModelMessageColumn = "message"
	// UninstallRequestModelNumCPUColumn is the column json value num_cpu
	UninstallRequestModelNumCPUColumn = "num_cpu"
	// UninstallRequestModelOSColumn is the column json value os
	UninstallRequestModelOSColumn = "os"
	// UninstallRequestModelRefIDColumn is the column json value ref_id
	UninstallRequestModelRefIDColumn = "ref_id"
	// UninstallRequestModelRefTypeColumn is the column json value ref_type
	UninstallRequestModelRefTypeColumn = "ref_type"
	// UninstallRequestModelRequestIDColumn is the column json value request_id
	UninstallRequestModelRequestIDColumn = "request_id"
	// UninstallRequestModelSuccessColumn is the column json value success
	UninstallRequestModelSuccessColumn = "success"
	// UninstallRequestModelSystemIDColumn is the column json value system_id
	UninstallRequestModelSystemIDColumn = "system_id"
	// UninstallRequestModelTypeColumn is the column json value type
	UninstallRequestModelTypeColumn = "type"
	// UninstallRequestModelUptimeColumn is the column json value uptime
	UninstallRequestModelUptimeColumn = "uptime"
	// UninstallRequestModelUUIDColumn is the column json value uuid
	UninstallRequestModelUUIDColumn = "uuid"
	// UninstallRequestModelVersionColumn is the column json value version
	UninstallRequestModelVersionColumn = "version"
)

// UninstallRequestEventDate represents the object structure for event_date
type UninstallRequestEventDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toUninstallRequestEventDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *UninstallRequestEventDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *UninstallRequestEventDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toUninstallRequestEventDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toUninstallRequestEventDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toUninstallRequestEventDateObject(o.Rfc3339, false),
	}
}

func (o *UninstallRequestEventDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *UninstallRequestEventDate) FromMap(kv map[string]interface{}) {

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

// UninstallRequestLastExportDate represents the object structure for last_export_date
type UninstallRequestLastExportDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toUninstallRequestLastExportDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *UninstallRequestLastExportDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *UninstallRequestLastExportDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toUninstallRequestLastExportDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toUninstallRequestLastExportDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toUninstallRequestLastExportDateObject(o.Rfc3339, false),
	}
}

func (o *UninstallRequestLastExportDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *UninstallRequestLastExportDate) FromMap(kv map[string]interface{}) {

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

// UninstallRequestType is the enumeration type for type
type UninstallRequestType int32

// toUninstallRequestTypePointer is the enumeration pointer type for type
func toUninstallRequestTypePointer(v int32) *UninstallRequestType {
	nv := UninstallRequestType(v)
	return &nv
}

// toUninstallRequestTypeEnum is the enumeration pointer wrapper for type
func toUninstallRequestTypeEnum(v *UninstallRequestType) string {
	if v == nil {
		return toUninstallRequestTypePointer(0).String()
	}
	return v.String()
}

// UnmarshalBSONValue for unmarshaling value
func (v *UninstallRequestType) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = UninstallRequestType(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "ENROLL":
			*v = UninstallRequestType(0)
		case "PING":
			*v = UninstallRequestType(1)
		case "CRASH":
			*v = UninstallRequestType(2)
		case "LOG":
			*v = UninstallRequestType(3)
		case "INTEGRATION":
			*v = UninstallRequestType(4)
		case "EXPORT":
			*v = UninstallRequestType(5)
		case "PROJECT":
			*v = UninstallRequestType(6)
		case "REPO":
			*v = UninstallRequestType(7)
		case "USER":
			*v = UninstallRequestType(8)
		case "CALENDAR":
			*v = UninstallRequestType(9)
		case "UNINSTALL":
			*v = UninstallRequestType(10)
		case "UPGRADE":
			*v = UninstallRequestType(11)
		case "START":
			*v = UninstallRequestType(12)
		case "STOP":
			*v = UninstallRequestType(13)
		case "PAUSE":
			*v = UninstallRequestType(14)
		case "RESUME":
			*v = UninstallRequestType(15)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v *UninstallRequestType) UnmarshalJSON(buf []byte) error {
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
func (v UninstallRequestType) MarshalJSON() ([]byte, error) {
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
func (v UninstallRequestType) String() string {
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
func (v *UninstallRequestType) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case UninstallRequestType:
		*v = val
	case int32:
		*v = UninstallRequestType(int32(val))
	case int:
		*v = UninstallRequestType(int32(val))
	case string:
		switch val {
		case "ENROLL":
			*v = UninstallRequestType(0)
		case "PING":
			*v = UninstallRequestType(1)
		case "CRASH":
			*v = UninstallRequestType(2)
		case "LOG":
			*v = UninstallRequestType(3)
		case "INTEGRATION":
			*v = UninstallRequestType(4)
		case "EXPORT":
			*v = UninstallRequestType(5)
		case "PROJECT":
			*v = UninstallRequestType(6)
		case "REPO":
			*v = UninstallRequestType(7)
		case "USER":
			*v = UninstallRequestType(8)
		case "CALENDAR":
			*v = UninstallRequestType(9)
		case "UNINSTALL":
			*v = UninstallRequestType(10)
		case "UPGRADE":
			*v = UninstallRequestType(11)
		case "START":
			*v = UninstallRequestType(12)
		case "STOP":
			*v = UninstallRequestType(13)
		case "PAUSE":
			*v = UninstallRequestType(14)
		case "RESUME":
			*v = UninstallRequestType(15)
		}
	}
	return nil
}

const (
	// UninstallRequestTypeEnroll is the enumeration value for enroll
	UninstallRequestTypeEnroll UninstallRequestType = 0
	// UninstallRequestTypePing is the enumeration value for ping
	UninstallRequestTypePing UninstallRequestType = 1
	// UninstallRequestTypeCrash is the enumeration value for crash
	UninstallRequestTypeCrash UninstallRequestType = 2
	// UninstallRequestTypeLog is the enumeration value for log
	UninstallRequestTypeLog UninstallRequestType = 3
	// UninstallRequestTypeIntegration is the enumeration value for integration
	UninstallRequestTypeIntegration UninstallRequestType = 4
	// UninstallRequestTypeExport is the enumeration value for export
	UninstallRequestTypeExport UninstallRequestType = 5
	// UninstallRequestTypeProject is the enumeration value for project
	UninstallRequestTypeProject UninstallRequestType = 6
	// UninstallRequestTypeRepo is the enumeration value for repo
	UninstallRequestTypeRepo UninstallRequestType = 7
	// UninstallRequestTypeUser is the enumeration value for user
	UninstallRequestTypeUser UninstallRequestType = 8
	// UninstallRequestTypeCalendar is the enumeration value for calendar
	UninstallRequestTypeCalendar UninstallRequestType = 9
	// UninstallRequestTypeUninstall is the enumeration value for uninstall
	UninstallRequestTypeUninstall UninstallRequestType = 10
	// UninstallRequestTypeUpgrade is the enumeration value for upgrade
	UninstallRequestTypeUpgrade UninstallRequestType = 11
	// UninstallRequestTypeStart is the enumeration value for start
	UninstallRequestTypeStart UninstallRequestType = 12
	// UninstallRequestTypeStop is the enumeration value for stop
	UninstallRequestTypeStop UninstallRequestType = 13
	// UninstallRequestTypePause is the enumeration value for pause
	UninstallRequestTypePause UninstallRequestType = 14
	// UninstallRequestTypeResume is the enumeration value for resume
	UninstallRequestTypeResume UninstallRequestType = 15
)

// UninstallRequest an event to request agent to uninstall itself
type UninstallRequest struct {
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
	EventDate UninstallRequestEventDate `json:"event_date" codec:"event_date" bson:"event_date" yaml:"event_date" faker:"-"`
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
	LastExportDate UninstallRequestLastExportDate `json:"last_export_date" codec:"last_export_date" bson:"last_export_date" yaml:"last_export_date" faker:"-"`
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
	Type UninstallRequestType `json:"type" codec:"type" bson:"type" yaml:"type" faker:"-"`
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
var _ datamodel.Model = (*UninstallRequest)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*UninstallRequest)(nil)

func toUninstallRequestObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *UninstallRequest:
		return v.ToMap()

	case UninstallRequestEventDate:
		return v.ToMap()

	case UninstallRequestLastExportDate:
		return v.ToMap()

	case UninstallRequestType:
		return v.String()

	default:
		return o
	}
}

// String returns a string representation of UninstallRequest
func (o *UninstallRequest) String() string {
	return fmt.Sprintf("agent.UninstallRequest<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *UninstallRequest) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *UninstallRequest) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *UninstallRequest) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *UninstallRequest) GetModelName() datamodel.ModelNameType {
	return UninstallRequestModelName
}

// NewUninstallRequestID provides a template for generating an ID field for UninstallRequest
func NewUninstallRequestID(customerID string, refType string, refID string) string {
	return hash.Values("UninstallRequest", customerID, refType, refID)
}

func (o *UninstallRequest) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("UninstallRequest", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *UninstallRequest) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *UninstallRequest) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *UninstallRequest) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *UninstallRequest) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *UninstallRequest) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *UninstallRequest) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *UninstallRequest) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *UninstallRequest) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *UninstallRequest) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = UninstallRequestModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *UninstallRequest) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *UninstallRequest) GetCustomerID() string {
	return o.CustomerID
}

// SetCustomerID will return the customer_id
func (o *UninstallRequest) SetCustomerID(id string) {
	o.CustomerID = id
}

// GetRefType will return the ref_type
func (o *UninstallRequest) GetRefType() string {
	return o.RefType
}

// SetRefType will return the ref_type
func (o *UninstallRequest) SetRefType(t string) {
	o.RefType = t
}

// Clone returns an exact copy of UninstallRequest
func (o *UninstallRequest) Clone() datamodel.Model {
	c := new(UninstallRequest)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *UninstallRequest) Anon() datamodel.Model {
	c := new(UninstallRequest)
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
func (o *UninstallRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *UninstallRequest) UnmarshalJSON(data []byte) error {
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
func (o *UninstallRequest) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *UninstallRequest) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two UninstallRequest objects are equal
func (o *UninstallRequest) IsEqual(other *UninstallRequest) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *UninstallRequest) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"architecture":            toUninstallRequestObject(o.Architecture, false),
		"customer_id":             toUninstallRequestObject(o.CustomerID, false),
		"data":                    toUninstallRequestObject(o.Data, true),
		"distro":                  toUninstallRequestObject(o.Distro, false),
		"error":                   toUninstallRequestObject(o.Error, true),
		"event_date":              toUninstallRequestObject(o.EventDate, false),
		"free_space":              toUninstallRequestObject(o.FreeSpace, false),
		"go_version":              toUninstallRequestObject(o.GoVersion, false),
		"hostname":                toUninstallRequestObject(o.Hostname, false),
		"id":                      toUninstallRequestObject(o.ID, false),
		"integration_instance_id": toUninstallRequestObject(o.IntegrationInstanceID, true),
		"last_export_date":        toUninstallRequestObject(o.LastExportDate, false),
		"memory":                  toUninstallRequestObject(o.Memory, false),
		"message":                 toUninstallRequestObject(o.Message, false),
		"num_cpu":                 toUninstallRequestObject(o.NumCPU, false),
		"os":                      toUninstallRequestObject(o.OS, false),
		"ref_id":                  toUninstallRequestObject(o.RefID, false),
		"ref_type":                toUninstallRequestObject(o.RefType, false),
		"request_id":              toUninstallRequestObject(o.RequestID, false),
		"success":                 toUninstallRequestObject(o.Success, false),
		"system_id":               toUninstallRequestObject(o.SystemID, false),

		"type":     o.Type.String(),
		"uptime":   toUninstallRequestObject(o.Uptime, false),
		"uuid":     toUninstallRequestObject(o.UUID, false),
		"version":  toUninstallRequestObject(o.Version, false),
		"hashcode": toUninstallRequestObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *UninstallRequest) FromMap(kv map[string]interface{}) {

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
		} else if sv, ok := val.(UninstallRequestEventDate); ok {
			// struct
			o.EventDate = sv
		} else if sp, ok := val.(*UninstallRequestEventDate); ok {
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
		} else if sv, ok := val.(UninstallRequestLastExportDate); ok {
			// struct
			o.LastExportDate = sv
		} else if sp, ok := val.(*UninstallRequestLastExportDate); ok {
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
	if val, ok := kv["type"].(UninstallRequestType); ok {
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
func (o *UninstallRequest) Hash() string {
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
func (o *UninstallRequest) SetIntegrationInstanceID(id string) {
	if id == "" {
		o.IntegrationInstanceID = nil
	} else {
		o.IntegrationInstanceID = &id
	}
}

// GetIntegrationInstanceID will return the integration instance ID
func (o *UninstallRequest) GetIntegrationInstanceID() *string {
	return o.IntegrationInstanceID
}

// UninstallRequestPartial is a partial struct for upsert mutations for UninstallRequest
type UninstallRequestPartial struct {
	// Architecture the architecture of the agent machine
	Architecture *string `json:"architecture,omitempty"`
	// Data extra data that is specific about this event
	Data *string `json:"data,omitempty"`
	// Distro the agent os distribution
	Distro *string `json:"distro,omitempty"`
	// Error an error message related to this event
	Error *string `json:"error,omitempty"`
	// EventDate the date of the event
	EventDate *UninstallRequestEventDate `json:"event_date,omitempty"`
	// FreeSpace the amount of free space in bytes for the agent machine
	FreeSpace *int64 `json:"free_space,omitempty"`
	// GoVersion the go version that the agent build was built with
	GoVersion *string `json:"go_version,omitempty"`
	// Hostname the agent hostname
	Hostname *string `json:"hostname,omitempty"`
	// LastExportDate the last export date
	LastExportDate *UninstallRequestLastExportDate `json:"last_export_date,omitempty"`
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
	Type *UninstallRequestType `json:"type,omitempty"`
	// Uptime the uptime in milliseconds since the agent started
	Uptime *int64 `json:"uptime,omitempty"`
	// UUID the agent unique identifier
	UUID *string `json:"uuid,omitempty"`
	// Version the agent version
	Version *string `json:"version,omitempty"`
}

var _ datamodel.PartialModel = (*UninstallRequestPartial)(nil)

// GetModelName returns the name of the model
func (o *UninstallRequestPartial) GetModelName() datamodel.ModelNameType {
	return UninstallRequestModelName
}

// ToMap returns the object as a map
func (o *UninstallRequestPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"architecture":     toUninstallRequestObject(o.Architecture, true),
		"data":             toUninstallRequestObject(o.Data, true),
		"distro":           toUninstallRequestObject(o.Distro, true),
		"error":            toUninstallRequestObject(o.Error, true),
		"event_date":       toUninstallRequestObject(o.EventDate, true),
		"free_space":       toUninstallRequestObject(o.FreeSpace, true),
		"go_version":       toUninstallRequestObject(o.GoVersion, true),
		"hostname":         toUninstallRequestObject(o.Hostname, true),
		"last_export_date": toUninstallRequestObject(o.LastExportDate, true),
		"memory":           toUninstallRequestObject(o.Memory, true),
		"message":          toUninstallRequestObject(o.Message, true),
		"num_cpu":          toUninstallRequestObject(o.NumCPU, true),
		"os":               toUninstallRequestObject(o.OS, true),
		"request_id":       toUninstallRequestObject(o.RequestID, true),
		"success":          toUninstallRequestObject(o.Success, true),
		"system_id":        toUninstallRequestObject(o.SystemID, true),

		"type":    toUninstallRequestTypeEnum(o.Type),
		"uptime":  toUninstallRequestObject(o.Uptime, true),
		"uuid":    toUninstallRequestObject(o.UUID, true),
		"version": toUninstallRequestObject(o.Version, true),
	}
	for k, v := range kv {
		if v == nil || reflect.ValueOf(v).IsZero() {
			delete(kv, k)
		} else {
			if k == "event_date" {
				if dt, ok := v.(*UninstallRequestEventDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}
			if k == "last_export_date" {
				if dt, ok := v.(*UninstallRequestLastExportDate); ok {
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
func (o *UninstallRequestPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *UninstallRequestPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *UninstallRequestPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *UninstallRequestPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *UninstallRequestPartial) FromMap(kv map[string]interface{}) {
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
		o.EventDate = &UninstallRequestEventDate{}
	}

	if val, ok := kv["event_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.EventDate.FromMap(kv)
		} else if sv, ok := val.(UninstallRequestEventDate); ok {
			// struct
			o.EventDate = &sv
		} else if sp, ok := val.(*UninstallRequestEventDate); ok {
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
		o.LastExportDate = &UninstallRequestLastExportDate{}
	}

	if val, ok := kv["last_export_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastExportDate.FromMap(kv)
		} else if sv, ok := val.(UninstallRequestLastExportDate); ok {
			// struct
			o.LastExportDate = &sv
		} else if sp, ok := val.(*UninstallRequestLastExportDate); ok {
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
	if val, ok := kv["type"].(*UninstallRequestType); ok {
		o.Type = val
	} else if val, ok := kv["type"].(UninstallRequestType); ok {
		o.Type = &val
	} else {
		if val, ok := kv["type"]; ok {
			if val == nil {
				o.Type = toUninstallRequestTypePointer(0)
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["UninstallRequestType"]
				}
				// this is an enum pointer
				if em, ok := val.(string); ok {
					switch em {
					case "enroll", "ENROLL":
						o.Type = toUninstallRequestTypePointer(0)
					case "ping", "PING":
						o.Type = toUninstallRequestTypePointer(1)
					case "crash", "CRASH":
						o.Type = toUninstallRequestTypePointer(2)
					case "log", "LOG":
						o.Type = toUninstallRequestTypePointer(3)
					case "integration", "INTEGRATION":
						o.Type = toUninstallRequestTypePointer(4)
					case "export", "EXPORT":
						o.Type = toUninstallRequestTypePointer(5)
					case "project", "PROJECT":
						o.Type = toUninstallRequestTypePointer(6)
					case "repo", "REPO":
						o.Type = toUninstallRequestTypePointer(7)
					case "user", "USER":
						o.Type = toUninstallRequestTypePointer(8)
					case "calendar", "CALENDAR":
						o.Type = toUninstallRequestTypePointer(9)
					case "uninstall", "UNINSTALL":
						o.Type = toUninstallRequestTypePointer(10)
					case "upgrade", "UPGRADE":
						o.Type = toUninstallRequestTypePointer(11)
					case "start", "START":
						o.Type = toUninstallRequestTypePointer(12)
					case "stop", "STOP":
						o.Type = toUninstallRequestTypePointer(13)
					case "pause", "PAUSE":
						o.Type = toUninstallRequestTypePointer(14)
					case "resume", "RESUME":
						o.Type = toUninstallRequestTypePointer(15)
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
