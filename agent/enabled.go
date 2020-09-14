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

	// EnabledTable is the default table name
	EnabledTable datamodel.ModelNameType = "agent_enabled"

	// EnabledModelName is the model name
	EnabledModelName datamodel.ModelNameType = "agent.Enabled"
)

const (
	// EnabledModelArchitectureColumn is the column json value architecture
	EnabledModelArchitectureColumn = "architecture"
	// EnabledModelCustomerIDColumn is the column json value customer_id
	EnabledModelCustomerIDColumn = "customer_id"
	// EnabledModelDataColumn is the column json value data
	EnabledModelDataColumn = "data"
	// EnabledModelDistroColumn is the column json value distro
	EnabledModelDistroColumn = "distro"
	// EnabledModelErrorColumn is the column json value error
	EnabledModelErrorColumn = "error"
	// EnabledModelEventDateColumn is the column json value event_date
	EnabledModelEventDateColumn = "event_date"
	// EnabledModelEventDateEpochColumn is the column json value epoch
	EnabledModelEventDateEpochColumn = "epoch"
	// EnabledModelEventDateOffsetColumn is the column json value offset
	EnabledModelEventDateOffsetColumn = "offset"
	// EnabledModelEventDateRfc3339Column is the column json value rfc3339
	EnabledModelEventDateRfc3339Column = "rfc3339"
	// EnabledModelFreeSpaceColumn is the column json value free_space
	EnabledModelFreeSpaceColumn = "free_space"
	// EnabledModelGoVersionColumn is the column json value go_version
	EnabledModelGoVersionColumn = "go_version"
	// EnabledModelHostnameColumn is the column json value hostname
	EnabledModelHostnameColumn = "hostname"
	// EnabledModelIDColumn is the column json value id
	EnabledModelIDColumn = "id"
	// EnabledModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	EnabledModelIntegrationInstanceIDColumn = "integration_instance_id"
	// EnabledModelLastExportDateColumn is the column json value last_export_date
	EnabledModelLastExportDateColumn = "last_export_date"
	// EnabledModelLastExportDateEpochColumn is the column json value epoch
	EnabledModelLastExportDateEpochColumn = "epoch"
	// EnabledModelLastExportDateOffsetColumn is the column json value offset
	EnabledModelLastExportDateOffsetColumn = "offset"
	// EnabledModelLastExportDateRfc3339Column is the column json value rfc3339
	EnabledModelLastExportDateRfc3339Column = "rfc3339"
	// EnabledModelMemoryColumn is the column json value memory
	EnabledModelMemoryColumn = "memory"
	// EnabledModelMessageColumn is the column json value message
	EnabledModelMessageColumn = "message"
	// EnabledModelNumCPUColumn is the column json value num_cpu
	EnabledModelNumCPUColumn = "num_cpu"
	// EnabledModelOSColumn is the column json value os
	EnabledModelOSColumn = "os"
	// EnabledModelRefIDColumn is the column json value ref_id
	EnabledModelRefIDColumn = "ref_id"
	// EnabledModelRefTypeColumn is the column json value ref_type
	EnabledModelRefTypeColumn = "ref_type"
	// EnabledModelRequestIDColumn is the column json value request_id
	EnabledModelRequestIDColumn = "request_id"
	// EnabledModelSuccessColumn is the column json value success
	EnabledModelSuccessColumn = "success"
	// EnabledModelSystemIDColumn is the column json value system_id
	EnabledModelSystemIDColumn = "system_id"
	// EnabledModelTypeColumn is the column json value type
	EnabledModelTypeColumn = "type"
	// EnabledModelUptimeColumn is the column json value uptime
	EnabledModelUptimeColumn = "uptime"
	// EnabledModelUUIDColumn is the column json value uuid
	EnabledModelUUIDColumn = "uuid"
	// EnabledModelVersionColumn is the column json value version
	EnabledModelVersionColumn = "version"
)

// EnabledEventDate represents the object structure for event_date
type EnabledEventDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toEnabledEventDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *EnabledEventDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *EnabledEventDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toEnabledEventDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toEnabledEventDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toEnabledEventDateObject(o.Rfc3339, false),
	}
}

func (o *EnabledEventDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *EnabledEventDate) FromMap(kv map[string]interface{}) {

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

// EnabledLastExportDate represents the object structure for last_export_date
type EnabledLastExportDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toEnabledLastExportDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *EnabledLastExportDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *EnabledLastExportDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toEnabledLastExportDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toEnabledLastExportDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toEnabledLastExportDateObject(o.Rfc3339, false),
	}
}

func (o *EnabledLastExportDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *EnabledLastExportDate) FromMap(kv map[string]interface{}) {

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

// EnabledType is the enumeration type for type
type EnabledType int32

// toEnabledTypePointer is the enumeration pointer type for type
func toEnabledTypePointer(v int32) *EnabledType {
	nv := EnabledType(v)
	return &nv
}

// toEnabledTypeEnum is the enumeration pointer wrapper for type
func toEnabledTypeEnum(v *EnabledType) string {
	if v == nil {
		return toEnabledTypePointer(0).String()
	}
	return v.String()
}

// UnmarshalBSONValue for unmarshaling value
func (v *EnabledType) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = EnabledType(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "ENROLL":
			*v = EnabledType(0)
		case "PING":
			*v = EnabledType(1)
		case "CRASH":
			*v = EnabledType(2)
		case "LOG":
			*v = EnabledType(3)
		case "INTEGRATION":
			*v = EnabledType(4)
		case "EXPORT":
			*v = EnabledType(5)
		case "PROJECT":
			*v = EnabledType(6)
		case "REPO":
			*v = EnabledType(7)
		case "USER":
			*v = EnabledType(8)
		case "CALENDAR":
			*v = EnabledType(9)
		case "UNINSTALL":
			*v = EnabledType(10)
		case "UPGRADE":
			*v = EnabledType(11)
		case "START":
			*v = EnabledType(12)
		case "STOP":
			*v = EnabledType(13)
		case "PAUSE":
			*v = EnabledType(14)
		case "RESUME":
			*v = EnabledType(15)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v *EnabledType) UnmarshalJSON(buf []byte) error {
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
func (v EnabledType) MarshalJSON() ([]byte, error) {
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
func (v EnabledType) String() string {
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
func (v *EnabledType) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case EnabledType:
		*v = val
	case int32:
		*v = EnabledType(int32(val))
	case int:
		*v = EnabledType(int32(val))
	case string:
		switch val {
		case "ENROLL":
			*v = EnabledType(0)
		case "PING":
			*v = EnabledType(1)
		case "CRASH":
			*v = EnabledType(2)
		case "LOG":
			*v = EnabledType(3)
		case "INTEGRATION":
			*v = EnabledType(4)
		case "EXPORT":
			*v = EnabledType(5)
		case "PROJECT":
			*v = EnabledType(6)
		case "REPO":
			*v = EnabledType(7)
		case "USER":
			*v = EnabledType(8)
		case "CALENDAR":
			*v = EnabledType(9)
		case "UNINSTALL":
			*v = EnabledType(10)
		case "UPGRADE":
			*v = EnabledType(11)
		case "START":
			*v = EnabledType(12)
		case "STOP":
			*v = EnabledType(13)
		case "PAUSE":
			*v = EnabledType(14)
		case "RESUME":
			*v = EnabledType(15)
		}
	}
	return nil
}

const (
	// EnabledTypeEnroll is the enumeration value for enroll
	EnabledTypeEnroll EnabledType = 0
	// EnabledTypePing is the enumeration value for ping
	EnabledTypePing EnabledType = 1
	// EnabledTypeCrash is the enumeration value for crash
	EnabledTypeCrash EnabledType = 2
	// EnabledTypeLog is the enumeration value for log
	EnabledTypeLog EnabledType = 3
	// EnabledTypeIntegration is the enumeration value for integration
	EnabledTypeIntegration EnabledType = 4
	// EnabledTypeExport is the enumeration value for export
	EnabledTypeExport EnabledType = 5
	// EnabledTypeProject is the enumeration value for project
	EnabledTypeProject EnabledType = 6
	// EnabledTypeRepo is the enumeration value for repo
	EnabledTypeRepo EnabledType = 7
	// EnabledTypeUser is the enumeration value for user
	EnabledTypeUser EnabledType = 8
	// EnabledTypeCalendar is the enumeration value for calendar
	EnabledTypeCalendar EnabledType = 9
	// EnabledTypeUninstall is the enumeration value for uninstall
	EnabledTypeUninstall EnabledType = 10
	// EnabledTypeUpgrade is the enumeration value for upgrade
	EnabledTypeUpgrade EnabledType = 11
	// EnabledTypeStart is the enumeration value for start
	EnabledTypeStart EnabledType = 12
	// EnabledTypeStop is the enumeration value for stop
	EnabledTypeStop EnabledType = 13
	// EnabledTypePause is the enumeration value for pause
	EnabledTypePause EnabledType = 14
	// EnabledTypeResume is the enumeration value for resume
	EnabledTypeResume EnabledType = 15
)

// Enabled an agent event to indicate that it's enabled and ready for actions
type Enabled struct {
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
	EventDate EnabledEventDate `json:"event_date" codec:"event_date" bson:"event_date" yaml:"event_date" faker:"-"`
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
	LastExportDate EnabledLastExportDate `json:"last_export_date" codec:"last_export_date" bson:"last_export_date" yaml:"last_export_date" faker:"-"`
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
	Type EnabledType `json:"type" codec:"type" bson:"type" yaml:"type" faker:"-"`
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
var _ datamodel.Model = (*Enabled)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*Enabled)(nil)

func toEnabledObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *Enabled:
		return v.ToMap()

	case EnabledEventDate:
		return v.ToMap()

	case EnabledLastExportDate:
		return v.ToMap()

	case EnabledType:
		return v.String()

	default:
		return o
	}
}

// String returns a string representation of Enabled
func (o *Enabled) String() string {
	return fmt.Sprintf("agent.Enabled<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Enabled) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *Enabled) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *Enabled) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *Enabled) GetModelName() datamodel.ModelNameType {
	return EnabledModelName
}

// NewEnabledID provides a template for generating an ID field for Enabled
func NewEnabledID(customerID string, refType string, refID string) string {
	return hash.Values("Enabled", customerID, refType, refID)
}

func (o *Enabled) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Enabled", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Enabled) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Enabled) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Enabled) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *Enabled) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Enabled) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *Enabled) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Enabled) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Enabled) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *Enabled) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = EnabledModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Enabled) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *Enabled) GetCustomerID() string {

	return o.CustomerID

}

// SetCustomerID will return the customer_id
func (o *Enabled) SetCustomerID(id string) {

	o.CustomerID = id

}

// GetRefType will return the ref_type
func (o *Enabled) GetRefType() string {
	return o.RefType
}

// SetRefType will return the ref_type
func (o *Enabled) SetRefType(t string) {
	o.RefType = t
}

// Clone returns an exact copy of Enabled
func (o *Enabled) Clone() datamodel.Model {
	c := new(Enabled)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Enabled) Anon() datamodel.Model {
	c := new(Enabled)
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
func (o *Enabled) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Enabled) UnmarshalJSON(data []byte) error {
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
func (o *Enabled) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *Enabled) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two Enabled objects are equal
func (o *Enabled) IsEqual(other *Enabled) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Enabled) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"architecture":            toEnabledObject(o.Architecture, false),
		"customer_id":             toEnabledObject(o.CustomerID, false),
		"data":                    toEnabledObject(o.Data, true),
		"distro":                  toEnabledObject(o.Distro, false),
		"error":                   toEnabledObject(o.Error, true),
		"event_date":              toEnabledObject(o.EventDate, false),
		"free_space":              toEnabledObject(o.FreeSpace, false),
		"go_version":              toEnabledObject(o.GoVersion, false),
		"hostname":                toEnabledObject(o.Hostname, false),
		"id":                      toEnabledObject(o.ID, false),
		"integration_instance_id": toEnabledObject(o.IntegrationInstanceID, true),
		"last_export_date":        toEnabledObject(o.LastExportDate, false),
		"memory":                  toEnabledObject(o.Memory, false),
		"message":                 toEnabledObject(o.Message, false),
		"num_cpu":                 toEnabledObject(o.NumCPU, false),
		"os":                      toEnabledObject(o.OS, false),
		"ref_id":                  toEnabledObject(o.RefID, false),
		"ref_type":                toEnabledObject(o.RefType, false),
		"request_id":              toEnabledObject(o.RequestID, false),
		"success":                 toEnabledObject(o.Success, false),
		"system_id":               toEnabledObject(o.SystemID, false),

		"type":     o.Type.String(),
		"uptime":   toEnabledObject(o.Uptime, false),
		"uuid":     toEnabledObject(o.UUID, false),
		"version":  toEnabledObject(o.Version, false),
		"hashcode": toEnabledObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *Enabled) FromMap(kv map[string]interface{}) {

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
		} else if sv, ok := val.(EnabledEventDate); ok {
			// struct
			o.EventDate = sv
		} else if sp, ok := val.(*EnabledEventDate); ok {
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
		} else if sv, ok := val.(EnabledLastExportDate); ok {
			// struct
			o.LastExportDate = sv
		} else if sp, ok := val.(*EnabledLastExportDate); ok {
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
	if val, ok := kv["type"].(EnabledType); ok {
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
func (o *Enabled) Hash() string {
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
func (o *Enabled) SetIntegrationInstanceID(id string) {
	if id == "" {
		o.IntegrationInstanceID = nil
	} else {
		o.IntegrationInstanceID = &id
	}
}

// GetIntegrationInstanceID will return the integration instance ID
func (o *Enabled) GetIntegrationInstanceID() *string {
	return o.IntegrationInstanceID
}

// EnabledPartial is a partial struct for upsert mutations for Enabled
type EnabledPartial struct {
	// Architecture the architecture of the agent machine
	Architecture *string `json:"architecture,omitempty"`
	// Data extra data that is specific about this event
	Data *string `json:"data,omitempty"`
	// Distro the agent os distribution
	Distro *string `json:"distro,omitempty"`
	// Error an error message related to this event
	Error *string `json:"error,omitempty"`
	// EventDate the date of the event
	EventDate *EnabledEventDate `json:"event_date,omitempty"`
	// FreeSpace the amount of free space in bytes for the agent machine
	FreeSpace *int64 `json:"free_space,omitempty"`
	// GoVersion the go version that the agent build was built with
	GoVersion *string `json:"go_version,omitempty"`
	// Hostname the agent hostname
	Hostname *string `json:"hostname,omitempty"`
	// LastExportDate the last export date
	LastExportDate *EnabledLastExportDate `json:"last_export_date,omitempty"`
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
	Type *EnabledType `json:"type,omitempty"`
	// Uptime the uptime in milliseconds since the agent started
	Uptime *int64 `json:"uptime,omitempty"`
	// UUID the agent unique identifier
	UUID *string `json:"uuid,omitempty"`
	// Version the agent version
	Version *string `json:"version,omitempty"`
}

var _ datamodel.PartialModel = (*EnabledPartial)(nil)

// GetModelName returns the name of the model
func (o *EnabledPartial) GetModelName() datamodel.ModelNameType {
	return EnabledModelName
}

// ToMap returns the object as a map
func (o *EnabledPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"architecture":     toEnabledObject(o.Architecture, true),
		"data":             toEnabledObject(o.Data, true),
		"distro":           toEnabledObject(o.Distro, true),
		"error":            toEnabledObject(o.Error, true),
		"event_date":       toEnabledObject(o.EventDate, true),
		"free_space":       toEnabledObject(o.FreeSpace, true),
		"go_version":       toEnabledObject(o.GoVersion, true),
		"hostname":         toEnabledObject(o.Hostname, true),
		"last_export_date": toEnabledObject(o.LastExportDate, true),
		"memory":           toEnabledObject(o.Memory, true),
		"message":          toEnabledObject(o.Message, true),
		"num_cpu":          toEnabledObject(o.NumCPU, true),
		"os":               toEnabledObject(o.OS, true),
		"request_id":       toEnabledObject(o.RequestID, true),
		"success":          toEnabledObject(o.Success, true),
		"system_id":        toEnabledObject(o.SystemID, true),

		"type":    toEnabledTypeEnum(o.Type),
		"uptime":  toEnabledObject(o.Uptime, true),
		"uuid":    toEnabledObject(o.UUID, true),
		"version": toEnabledObject(o.Version, true),
	}
	for k, v := range kv {
		if v == nil || reflect.ValueOf(v).IsZero() {
			delete(kv, k)
		} else {
			if k == "event_date" {
				if dt, ok := v.(*EnabledEventDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}
			if k == "last_export_date" {
				if dt, ok := v.(*EnabledLastExportDate); ok {
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
func (o *EnabledPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *EnabledPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *EnabledPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *EnabledPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *EnabledPartial) FromMap(kv map[string]interface{}) {
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
		o.EventDate = &EnabledEventDate{}
	}

	if val, ok := kv["event_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.EventDate.FromMap(kv)
		} else if sv, ok := val.(EnabledEventDate); ok {
			// struct
			o.EventDate = &sv
		} else if sp, ok := val.(*EnabledEventDate); ok {
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
		o.LastExportDate = &EnabledLastExportDate{}
	}

	if val, ok := kv["last_export_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastExportDate.FromMap(kv)
		} else if sv, ok := val.(EnabledLastExportDate); ok {
			// struct
			o.LastExportDate = &sv
		} else if sp, ok := val.(*EnabledLastExportDate); ok {
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
	if val, ok := kv["type"].(*EnabledType); ok {
		o.Type = val
	} else if val, ok := kv["type"].(EnabledType); ok {
		o.Type = &val
	} else {
		if val, ok := kv["type"]; ok {
			if val == nil {
				o.Type = toEnabledTypePointer(0)
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["EnabledType"]
				}
				// this is an enum pointer
				if em, ok := val.(string); ok {
					switch em {
					case "enroll", "ENROLL":
						o.Type = toEnabledTypePointer(0)
					case "ping", "PING":
						o.Type = toEnabledTypePointer(1)
					case "crash", "CRASH":
						o.Type = toEnabledTypePointer(2)
					case "log", "LOG":
						o.Type = toEnabledTypePointer(3)
					case "integration", "INTEGRATION":
						o.Type = toEnabledTypePointer(4)
					case "export", "EXPORT":
						o.Type = toEnabledTypePointer(5)
					case "project", "PROJECT":
						o.Type = toEnabledTypePointer(6)
					case "repo", "REPO":
						o.Type = toEnabledTypePointer(7)
					case "user", "USER":
						o.Type = toEnabledTypePointer(8)
					case "calendar", "CALENDAR":
						o.Type = toEnabledTypePointer(9)
					case "uninstall", "UNINSTALL":
						o.Type = toEnabledTypePointer(10)
					case "upgrade", "UPGRADE":
						o.Type = toEnabledTypePointer(11)
					case "start", "START":
						o.Type = toEnabledTypePointer(12)
					case "stop", "STOP":
						o.Type = toEnabledTypePointer(13)
					case "pause", "PAUSE":
						o.Type = toEnabledTypePointer(14)
					case "resume", "RESUME":
						o.Type = toEnabledTypePointer(15)
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
