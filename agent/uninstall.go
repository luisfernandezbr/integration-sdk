// DO NOT EDIT -- generated code

// Package agent - the agent communicates with pinpoint cloud to send integration data
package agent

import (
	"encoding/json"
	"fmt"
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

	// UninstallTable is the default table name
	UninstallTable datamodel.ModelNameType = "agent_uninstall"

	// UninstallModelName is the model name
	UninstallModelName datamodel.ModelNameType = "agent.Uninstall"
)

const (
	// UninstallModelArchitectureColumn is the column json value architecture
	UninstallModelArchitectureColumn = "architecture"
	// UninstallModelCustomerIDColumn is the column json value customer_id
	UninstallModelCustomerIDColumn = "customer_id"
	// UninstallModelDataColumn is the column json value data
	UninstallModelDataColumn = "data"
	// UninstallModelDistroColumn is the column json value distro
	UninstallModelDistroColumn = "distro"
	// UninstallModelErrorColumn is the column json value error
	UninstallModelErrorColumn = "error"
	// UninstallModelEventDateColumn is the column json value event_date
	UninstallModelEventDateColumn = "event_date"
	// UninstallModelEventDateEpochColumn is the column json value epoch
	UninstallModelEventDateEpochColumn = "epoch"
	// UninstallModelEventDateOffsetColumn is the column json value offset
	UninstallModelEventDateOffsetColumn = "offset"
	// UninstallModelEventDateRfc3339Column is the column json value rfc3339
	UninstallModelEventDateRfc3339Column = "rfc3339"
	// UninstallModelFreeSpaceColumn is the column json value free_space
	UninstallModelFreeSpaceColumn = "free_space"
	// UninstallModelGoVersionColumn is the column json value go_version
	UninstallModelGoVersionColumn = "go_version"
	// UninstallModelHostnameColumn is the column json value hostname
	UninstallModelHostnameColumn = "hostname"
	// UninstallModelIDColumn is the column json value id
	UninstallModelIDColumn = "id"
	// UninstallModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	UninstallModelIntegrationInstanceIDColumn = "integration_instance_id"
	// UninstallModelLastExportDateColumn is the column json value last_export_date
	UninstallModelLastExportDateColumn = "last_export_date"
	// UninstallModelLastExportDateEpochColumn is the column json value epoch
	UninstallModelLastExportDateEpochColumn = "epoch"
	// UninstallModelLastExportDateOffsetColumn is the column json value offset
	UninstallModelLastExportDateOffsetColumn = "offset"
	// UninstallModelLastExportDateRfc3339Column is the column json value rfc3339
	UninstallModelLastExportDateRfc3339Column = "rfc3339"
	// UninstallModelMemoryColumn is the column json value memory
	UninstallModelMemoryColumn = "memory"
	// UninstallModelMessageColumn is the column json value message
	UninstallModelMessageColumn = "message"
	// UninstallModelNumCPUColumn is the column json value num_cpu
	UninstallModelNumCPUColumn = "num_cpu"
	// UninstallModelOSColumn is the column json value os
	UninstallModelOSColumn = "os"
	// UninstallModelRefIDColumn is the column json value ref_id
	UninstallModelRefIDColumn = "ref_id"
	// UninstallModelRefTypeColumn is the column json value ref_type
	UninstallModelRefTypeColumn = "ref_type"
	// UninstallModelRequestIDColumn is the column json value request_id
	UninstallModelRequestIDColumn = "request_id"
	// UninstallModelSuccessColumn is the column json value success
	UninstallModelSuccessColumn = "success"
	// UninstallModelSystemIDColumn is the column json value system_id
	UninstallModelSystemIDColumn = "system_id"
	// UninstallModelTypeColumn is the column json value type
	UninstallModelTypeColumn = "type"
	// UninstallModelUptimeColumn is the column json value uptime
	UninstallModelUptimeColumn = "uptime"
	// UninstallModelUUIDColumn is the column json value uuid
	UninstallModelUUIDColumn = "uuid"
	// UninstallModelVersionColumn is the column json value version
	UninstallModelVersionColumn = "version"
)

// UninstallEventDate represents the object structure for event_date
type UninstallEventDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toUninstallEventDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *UninstallEventDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *UninstallEventDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toUninstallEventDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toUninstallEventDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toUninstallEventDateObject(o.Rfc3339, false),
	}
}

func (o *UninstallEventDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *UninstallEventDate) FromMap(kv map[string]interface{}) {

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

// UninstallLastExportDate represents the object structure for last_export_date
type UninstallLastExportDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toUninstallLastExportDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *UninstallLastExportDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *UninstallLastExportDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toUninstallLastExportDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toUninstallLastExportDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toUninstallLastExportDateObject(o.Rfc3339, false),
	}
}

func (o *UninstallLastExportDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *UninstallLastExportDate) FromMap(kv map[string]interface{}) {

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

// UninstallType is the enumeration type for type
type UninstallType int32

// UnmarshalBSONValue for unmarshaling value
func (v *UninstallType) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = UninstallType(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "ENROLL":
			*v = UninstallType(0)
		case "PING":
			*v = UninstallType(1)
		case "CRASH":
			*v = UninstallType(2)
		case "LOG":
			*v = UninstallType(3)
		case "INTEGRATION":
			*v = UninstallType(4)
		case "EXPORT":
			*v = UninstallType(5)
		case "PROJECT":
			*v = UninstallType(6)
		case "REPO":
			*v = UninstallType(7)
		case "USER":
			*v = UninstallType(8)
		case "CALENDAR":
			*v = UninstallType(9)
		case "UNINSTALL":
			*v = UninstallType(10)
		case "UPGRADE":
			*v = UninstallType(11)
		case "START":
			*v = UninstallType(12)
		case "STOP":
			*v = UninstallType(13)
		case "PAUSE":
			*v = UninstallType(14)
		case "RESUME":
			*v = UninstallType(15)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v UninstallType) UnmarshalJSON(buf []byte) error {
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
func (v UninstallType) MarshalJSON() ([]byte, error) {
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
func (v UninstallType) String() string {
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
func (v *UninstallType) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case UninstallType:
		*v = val
	case int32:
		*v = UninstallType(int32(val))
	case int:
		*v = UninstallType(int32(val))
	case string:
		switch val {
		case "ENROLL":
			*v = UninstallType(0)
		case "PING":
			*v = UninstallType(1)
		case "CRASH":
			*v = UninstallType(2)
		case "LOG":
			*v = UninstallType(3)
		case "INTEGRATION":
			*v = UninstallType(4)
		case "EXPORT":
			*v = UninstallType(5)
		case "PROJECT":
			*v = UninstallType(6)
		case "REPO":
			*v = UninstallType(7)
		case "USER":
			*v = UninstallType(8)
		case "CALENDAR":
			*v = UninstallType(9)
		case "UNINSTALL":
			*v = UninstallType(10)
		case "UPGRADE":
			*v = UninstallType(11)
		case "START":
			*v = UninstallType(12)
		case "STOP":
			*v = UninstallType(13)
		case "PAUSE":
			*v = UninstallType(14)
		case "RESUME":
			*v = UninstallType(15)
		}
	}
	return nil
}

const (
	// UninstallTypeEnroll is the enumeration value for enroll
	UninstallTypeEnroll UninstallType = 0
	// UninstallTypePing is the enumeration value for ping
	UninstallTypePing UninstallType = 1
	// UninstallTypeCrash is the enumeration value for crash
	UninstallTypeCrash UninstallType = 2
	// UninstallTypeLog is the enumeration value for log
	UninstallTypeLog UninstallType = 3
	// UninstallTypeIntegration is the enumeration value for integration
	UninstallTypeIntegration UninstallType = 4
	// UninstallTypeExport is the enumeration value for export
	UninstallTypeExport UninstallType = 5
	// UninstallTypeProject is the enumeration value for project
	UninstallTypeProject UninstallType = 6
	// UninstallTypeRepo is the enumeration value for repo
	UninstallTypeRepo UninstallType = 7
	// UninstallTypeUser is the enumeration value for user
	UninstallTypeUser UninstallType = 8
	// UninstallTypeCalendar is the enumeration value for calendar
	UninstallTypeCalendar UninstallType = 9
	// UninstallTypeUninstall is the enumeration value for uninstall
	UninstallTypeUninstall UninstallType = 10
	// UninstallTypeUpgrade is the enumeration value for upgrade
	UninstallTypeUpgrade UninstallType = 11
	// UninstallTypeStart is the enumeration value for start
	UninstallTypeStart UninstallType = 12
	// UninstallTypeStop is the enumeration value for stop
	UninstallTypeStop UninstallType = 13
	// UninstallTypePause is the enumeration value for pause
	UninstallTypePause UninstallType = 14
	// UninstallTypeResume is the enumeration value for resume
	UninstallTypeResume UninstallType = 15
)

// Uninstall an agent event to indicate the agent was uninstalled
type Uninstall struct {
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
	EventDate UninstallEventDate `json:"event_date" codec:"event_date" bson:"event_date" yaml:"event_date" faker:"-"`
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
	LastExportDate UninstallLastExportDate `json:"last_export_date" codec:"last_export_date" bson:"last_export_date" yaml:"last_export_date" faker:"-"`
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
	Type UninstallType `json:"type" codec:"type" bson:"type" yaml:"type" faker:"-"`
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
var _ datamodel.Model = (*Uninstall)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*Uninstall)(nil)

func toUninstallObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *Uninstall:
		return v.ToMap()

	case UninstallEventDate:
		return v.ToMap()

	case UninstallLastExportDate:
		return v.ToMap()

	case UninstallType:
		return v.String()

	default:
		return o
	}
}

// String returns a string representation of Uninstall
func (o *Uninstall) String() string {
	return fmt.Sprintf("agent.Uninstall<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Uninstall) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *Uninstall) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *Uninstall) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *Uninstall) GetModelName() datamodel.ModelNameType {
	return UninstallModelName
}

// NewUninstallID provides a template for generating an ID field for Uninstall
func NewUninstallID(customerID string, refType string, refID string) string {
	return hash.Values("Uninstall", customerID, refType, refID)
}

func (o *Uninstall) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Uninstall", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Uninstall) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Uninstall) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Uninstall) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *Uninstall) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Uninstall) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *Uninstall) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Uninstall) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Uninstall) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *Uninstall) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = UninstallModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Uninstall) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *Uninstall) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of Uninstall
func (o *Uninstall) Clone() datamodel.Model {
	c := new(Uninstall)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Uninstall) Anon() datamodel.Model {
	c := new(Uninstall)
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
func (o *Uninstall) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Uninstall) UnmarshalJSON(data []byte) error {
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
func (o *Uninstall) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *Uninstall) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two Uninstall objects are equal
func (o *Uninstall) IsEqual(other *Uninstall) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Uninstall) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"architecture":            toUninstallObject(o.Architecture, false),
		"customer_id":             toUninstallObject(o.CustomerID, false),
		"data":                    toUninstallObject(o.Data, true),
		"distro":                  toUninstallObject(o.Distro, false),
		"error":                   toUninstallObject(o.Error, true),
		"event_date":              toUninstallObject(o.EventDate, false),
		"free_space":              toUninstallObject(o.FreeSpace, false),
		"go_version":              toUninstallObject(o.GoVersion, false),
		"hostname":                toUninstallObject(o.Hostname, false),
		"id":                      toUninstallObject(o.ID, false),
		"integration_instance_id": toUninstallObject(o.IntegrationInstanceID, true),
		"last_export_date":        toUninstallObject(o.LastExportDate, false),
		"memory":                  toUninstallObject(o.Memory, false),
		"message":                 toUninstallObject(o.Message, false),
		"num_cpu":                 toUninstallObject(o.NumCPU, false),
		"os":                      toUninstallObject(o.OS, false),
		"ref_id":                  toUninstallObject(o.RefID, false),
		"ref_type":                toUninstallObject(o.RefType, false),
		"request_id":              toUninstallObject(o.RequestID, false),
		"success":                 toUninstallObject(o.Success, false),
		"system_id":               toUninstallObject(o.SystemID, false),

		"type":     o.Type.String(),
		"uptime":   toUninstallObject(o.Uptime, false),
		"uuid":     toUninstallObject(o.UUID, false),
		"version":  toUninstallObject(o.Version, false),
		"hashcode": toUninstallObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *Uninstall) FromMap(kv map[string]interface{}) {

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
		} else if sv, ok := val.(UninstallEventDate); ok {
			// struct
			o.EventDate = sv
		} else if sp, ok := val.(*UninstallEventDate); ok {
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
		} else if sv, ok := val.(UninstallLastExportDate); ok {
			// struct
			o.LastExportDate = sv
		} else if sp, ok := val.(*UninstallLastExportDate); ok {
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
	if val, ok := kv["type"].(UninstallType); ok {
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
func (o *Uninstall) Hash() string {
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

// UninstallPartial is a partial struct for upsert mutations for Uninstall
type UninstallPartial struct {
	// Architecture the architecture of the agent machine
	Architecture *string `json:"architecture,omitempty"`
	// Data extra data that is specific about this event
	Data *string `json:"data,omitempty"`
	// Distro the agent os distribution
	Distro *string `json:"distro,omitempty"`
	// Error an error message related to this event
	Error *string `json:"error,omitempty"`
	// EventDate the date of the event
	EventDate *UninstallEventDate `json:"event_date,omitempty"`
	// FreeSpace the amount of free space in bytes for the agent machine
	FreeSpace *int64 `json:"free_space,omitempty"`
	// GoVersion the go version that the agent build was built with
	GoVersion *string `json:"go_version,omitempty"`
	// Hostname the agent hostname
	Hostname *string `json:"hostname,omitempty"`
	// LastExportDate the last export date
	LastExportDate *UninstallLastExportDate `json:"last_export_date,omitempty"`
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
	Type *UninstallType `json:"type,omitempty"`
	// Uptime the uptime in milliseconds since the agent started
	Uptime *int64 `json:"uptime,omitempty"`
	// UUID the agent unique identifier
	UUID *string `json:"uuid,omitempty"`
	// Version the agent version
	Version *string `json:"version,omitempty"`
}

var _ datamodel.PartialModel = (*UninstallPartial)(nil)

// GetModelName returns the name of the model
func (o *UninstallPartial) GetModelName() datamodel.ModelNameType {
	return UninstallModelName
}

// ToMap returns the object as a map
func (o *UninstallPartial) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"architecture":     toUninstallObject(o.Architecture, true),
		"data":             toUninstallObject(o.Data, true),
		"distro":           toUninstallObject(o.Distro, true),
		"error":            toUninstallObject(o.Error, true),
		"event_date":       toUninstallObject(o.EventDate, true),
		"free_space":       toUninstallObject(o.FreeSpace, true),
		"go_version":       toUninstallObject(o.GoVersion, true),
		"hostname":         toUninstallObject(o.Hostname, true),
		"last_export_date": toUninstallObject(o.LastExportDate, true),
		"memory":           toUninstallObject(o.Memory, true),
		"message":          toUninstallObject(o.Message, true),
		"num_cpu":          toUninstallObject(o.NumCPU, true),
		"os":               toUninstallObject(o.OS, true),
		"request_id":       toUninstallObject(o.RequestID, true),
		"success":          toUninstallObject(o.Success, true),
		"system_id":        toUninstallObject(o.SystemID, true),
		"type":             o.Type.String(),
		"uptime":           toUninstallObject(o.Uptime, true),
		"uuid":             toUninstallObject(o.UUID, true),
		"version":          toUninstallObject(o.Version, true),
	}
}

// Stringify returns the object in JSON format as a string
func (o *UninstallPartial) Stringify() string {
	return pjson.Stringify(o)
}
