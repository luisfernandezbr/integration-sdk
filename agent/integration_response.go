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

	// IntegrationResponseTable is the default table name
	IntegrationResponseTable datamodel.ModelNameType = "agent_integrationresponse"

	// IntegrationResponseModelName is the model name
	IntegrationResponseModelName datamodel.ModelNameType = "agent.IntegrationResponse"
)

const (
	// IntegrationResponseModelArchitectureColumn is the column json value architecture
	IntegrationResponseModelArchitectureColumn = "architecture"
	// IntegrationResponseModelAuthorizationColumn is the column json value authorization
	IntegrationResponseModelAuthorizationColumn = "authorization"
	// IntegrationResponseModelCustomerIDColumn is the column json value customer_id
	IntegrationResponseModelCustomerIDColumn = "customer_id"
	// IntegrationResponseModelDataColumn is the column json value data
	IntegrationResponseModelDataColumn = "data"
	// IntegrationResponseModelDistroColumn is the column json value distro
	IntegrationResponseModelDistroColumn = "distro"
	// IntegrationResponseModelErrorColumn is the column json value error
	IntegrationResponseModelErrorColumn = "error"
	// IntegrationResponseModelEventDateColumn is the column json value event_date
	IntegrationResponseModelEventDateColumn = "event_date"
	// IntegrationResponseModelEventDateEpochColumn is the column json value epoch
	IntegrationResponseModelEventDateEpochColumn = "epoch"
	// IntegrationResponseModelEventDateOffsetColumn is the column json value offset
	IntegrationResponseModelEventDateOffsetColumn = "offset"
	// IntegrationResponseModelEventDateRfc3339Column is the column json value rfc3339
	IntegrationResponseModelEventDateRfc3339Column = "rfc3339"
	// IntegrationResponseModelFreeSpaceColumn is the column json value free_space
	IntegrationResponseModelFreeSpaceColumn = "free_space"
	// IntegrationResponseModelGoVersionColumn is the column json value go_version
	IntegrationResponseModelGoVersionColumn = "go_version"
	// IntegrationResponseModelHostnameColumn is the column json value hostname
	IntegrationResponseModelHostnameColumn = "hostname"
	// IntegrationResponseModelIDColumn is the column json value id
	IntegrationResponseModelIDColumn = "id"
	// IntegrationResponseModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	IntegrationResponseModelIntegrationInstanceIDColumn = "integration_instance_id"
	// IntegrationResponseModelLastExportDateColumn is the column json value last_export_date
	IntegrationResponseModelLastExportDateColumn = "last_export_date"
	// IntegrationResponseModelLastExportDateEpochColumn is the column json value epoch
	IntegrationResponseModelLastExportDateEpochColumn = "epoch"
	// IntegrationResponseModelLastExportDateOffsetColumn is the column json value offset
	IntegrationResponseModelLastExportDateOffsetColumn = "offset"
	// IntegrationResponseModelLastExportDateRfc3339Column is the column json value rfc3339
	IntegrationResponseModelLastExportDateRfc3339Column = "rfc3339"
	// IntegrationResponseModelMemoryColumn is the column json value memory
	IntegrationResponseModelMemoryColumn = "memory"
	// IntegrationResponseModelMessageColumn is the column json value message
	IntegrationResponseModelMessageColumn = "message"
	// IntegrationResponseModelNumCPUColumn is the column json value num_cpu
	IntegrationResponseModelNumCPUColumn = "num_cpu"
	// IntegrationResponseModelOSColumn is the column json value os
	IntegrationResponseModelOSColumn = "os"
	// IntegrationResponseModelRefIDColumn is the column json value ref_id
	IntegrationResponseModelRefIDColumn = "ref_id"
	// IntegrationResponseModelRefTypeColumn is the column json value ref_type
	IntegrationResponseModelRefTypeColumn = "ref_type"
	// IntegrationResponseModelRequestIDColumn is the column json value request_id
	IntegrationResponseModelRequestIDColumn = "request_id"
	// IntegrationResponseModelServerVersionColumn is the column json value server_version
	IntegrationResponseModelServerVersionColumn = "server_version"
	// IntegrationResponseModelSuccessColumn is the column json value success
	IntegrationResponseModelSuccessColumn = "success"
	// IntegrationResponseModelSystemIDColumn is the column json value system_id
	IntegrationResponseModelSystemIDColumn = "system_id"
	// IntegrationResponseModelTypeColumn is the column json value type
	IntegrationResponseModelTypeColumn = "type"
	// IntegrationResponseModelUptimeColumn is the column json value uptime
	IntegrationResponseModelUptimeColumn = "uptime"
	// IntegrationResponseModelUUIDColumn is the column json value uuid
	IntegrationResponseModelUUIDColumn = "uuid"
	// IntegrationResponseModelVersionColumn is the column json value version
	IntegrationResponseModelVersionColumn = "version"
)

// IntegrationResponseEventDate represents the object structure for event_date
type IntegrationResponseEventDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIntegrationResponseEventDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationResponseEventDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *IntegrationResponseEventDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIntegrationResponseEventDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIntegrationResponseEventDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIntegrationResponseEventDateObject(o.Rfc3339, false),
	}
}

func (o *IntegrationResponseEventDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationResponseEventDate) FromMap(kv map[string]interface{}) {

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

// IntegrationResponseLastExportDate represents the object structure for last_export_date
type IntegrationResponseLastExportDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIntegrationResponseLastExportDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationResponseLastExportDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *IntegrationResponseLastExportDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIntegrationResponseLastExportDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIntegrationResponseLastExportDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIntegrationResponseLastExportDateObject(o.Rfc3339, false),
	}
}

func (o *IntegrationResponseLastExportDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationResponseLastExportDate) FromMap(kv map[string]interface{}) {

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

// IntegrationResponseType is the enumeration type for type
type IntegrationResponseType int32

// UnmarshalBSONValue for unmarshaling value
func (v *IntegrationResponseType) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = IntegrationResponseType(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "ENROLL":
			*v = IntegrationResponseType(0)
		case "PING":
			*v = IntegrationResponseType(1)
		case "CRASH":
			*v = IntegrationResponseType(2)
		case "LOG":
			*v = IntegrationResponseType(3)
		case "INTEGRATION":
			*v = IntegrationResponseType(4)
		case "EXPORT":
			*v = IntegrationResponseType(5)
		case "PROJECT":
			*v = IntegrationResponseType(6)
		case "REPO":
			*v = IntegrationResponseType(7)
		case "USER":
			*v = IntegrationResponseType(8)
		case "CALENDAR":
			*v = IntegrationResponseType(9)
		case "UNINSTALL":
			*v = IntegrationResponseType(10)
		case "UPGRADE":
			*v = IntegrationResponseType(11)
		case "START":
			*v = IntegrationResponseType(12)
		case "STOP":
			*v = IntegrationResponseType(13)
		case "PAUSE":
			*v = IntegrationResponseType(14)
		case "RESUME":
			*v = IntegrationResponseType(15)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v IntegrationResponseType) UnmarshalJSON(buf []byte) error {
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
func (v IntegrationResponseType) MarshalJSON() ([]byte, error) {
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
func (v IntegrationResponseType) String() string {
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
func (v *IntegrationResponseType) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case IntegrationResponseType:
		*v = val
	case int32:
		*v = IntegrationResponseType(int32(val))
	case int:
		*v = IntegrationResponseType(int32(val))
	case string:
		switch val {
		case "ENROLL":
			*v = IntegrationResponseType(0)
		case "PING":
			*v = IntegrationResponseType(1)
		case "CRASH":
			*v = IntegrationResponseType(2)
		case "LOG":
			*v = IntegrationResponseType(3)
		case "INTEGRATION":
			*v = IntegrationResponseType(4)
		case "EXPORT":
			*v = IntegrationResponseType(5)
		case "PROJECT":
			*v = IntegrationResponseType(6)
		case "REPO":
			*v = IntegrationResponseType(7)
		case "USER":
			*v = IntegrationResponseType(8)
		case "CALENDAR":
			*v = IntegrationResponseType(9)
		case "UNINSTALL":
			*v = IntegrationResponseType(10)
		case "UPGRADE":
			*v = IntegrationResponseType(11)
		case "START":
			*v = IntegrationResponseType(12)
		case "STOP":
			*v = IntegrationResponseType(13)
		case "PAUSE":
			*v = IntegrationResponseType(14)
		case "RESUME":
			*v = IntegrationResponseType(15)
		}
	}
	return nil
}

const (
	// IntegrationResponseTypeEnroll is the enumeration value for enroll
	IntegrationResponseTypeEnroll IntegrationResponseType = 0
	// IntegrationResponseTypePing is the enumeration value for ping
	IntegrationResponseTypePing IntegrationResponseType = 1
	// IntegrationResponseTypeCrash is the enumeration value for crash
	IntegrationResponseTypeCrash IntegrationResponseType = 2
	// IntegrationResponseTypeLog is the enumeration value for log
	IntegrationResponseTypeLog IntegrationResponseType = 3
	// IntegrationResponseTypeIntegration is the enumeration value for integration
	IntegrationResponseTypeIntegration IntegrationResponseType = 4
	// IntegrationResponseTypeExport is the enumeration value for export
	IntegrationResponseTypeExport IntegrationResponseType = 5
	// IntegrationResponseTypeProject is the enumeration value for project
	IntegrationResponseTypeProject IntegrationResponseType = 6
	// IntegrationResponseTypeRepo is the enumeration value for repo
	IntegrationResponseTypeRepo IntegrationResponseType = 7
	// IntegrationResponseTypeUser is the enumeration value for user
	IntegrationResponseTypeUser IntegrationResponseType = 8
	// IntegrationResponseTypeCalendar is the enumeration value for calendar
	IntegrationResponseTypeCalendar IntegrationResponseType = 9
	// IntegrationResponseTypeUninstall is the enumeration value for uninstall
	IntegrationResponseTypeUninstall IntegrationResponseType = 10
	// IntegrationResponseTypeUpgrade is the enumeration value for upgrade
	IntegrationResponseTypeUpgrade IntegrationResponseType = 11
	// IntegrationResponseTypeStart is the enumeration value for start
	IntegrationResponseTypeStart IntegrationResponseType = 12
	// IntegrationResponseTypeStop is the enumeration value for stop
	IntegrationResponseTypeStop IntegrationResponseType = 13
	// IntegrationResponseTypePause is the enumeration value for pause
	IntegrationResponseTypePause IntegrationResponseType = 14
	// IntegrationResponseTypeResume is the enumeration value for resume
	IntegrationResponseTypeResume IntegrationResponseType = 15
)

// IntegrationResponse an agent response to an action request adding an integration
type IntegrationResponse struct {
	// Architecture the architecture of the agent machine
	Architecture string `json:"architecture" codec:"architecture" bson:"architecture" yaml:"architecture" faker:"-"`
	// Authorization the encrypted authorization data for this integration
	Authorization string `json:"authorization" codec:"authorization" bson:"authorization" yaml:"authorization" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Data extra data that is specific about this event
	Data *string `json:"data,omitempty" codec:"data,omitempty" bson:"data" yaml:"data,omitempty" faker:"-"`
	// Distro the agent os distribution
	Distro string `json:"distro" codec:"distro" bson:"distro" yaml:"distro" faker:"-"`
	// Error an error message related to this event
	Error *string `json:"error,omitempty" codec:"error,omitempty" bson:"error" yaml:"error,omitempty" faker:"-"`
	// EventDate the date of the event
	EventDate IntegrationResponseEventDate `json:"event_date" codec:"event_date" bson:"event_date" yaml:"event_date" faker:"-"`
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
	LastExportDate IntegrationResponseLastExportDate `json:"last_export_date" codec:"last_export_date" bson:"last_export_date" yaml:"last_export_date" faker:"-"`
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
	// ServerVersion the server version for this integration
	ServerVersion *string `json:"server_version,omitempty" codec:"server_version,omitempty" bson:"server_version" yaml:"server_version,omitempty" faker:"-"`
	// Success if the response was successful
	Success bool `json:"success" codec:"success" bson:"success" yaml:"success" faker:"-"`
	// SystemID system unique device ID
	SystemID string `json:"system_id" codec:"system_id" bson:"system_id" yaml:"system_id" faker:"-"`
	// Type the type of event
	Type IntegrationResponseType `json:"type" codec:"type" bson:"type" yaml:"type" faker:"-"`
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
var _ datamodel.Model = (*IntegrationResponse)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*IntegrationResponse)(nil)

func toIntegrationResponseObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationResponse:
		return v.ToMap()

	case IntegrationResponseEventDate:
		return v.ToMap()

	case IntegrationResponseLastExportDate:
		return v.ToMap()

	case IntegrationResponseType:
		return v.String()

	default:
		return o
	}
}

// String returns a string representation of IntegrationResponse
func (o *IntegrationResponse) String() string {
	return fmt.Sprintf("agent.IntegrationResponse<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *IntegrationResponse) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *IntegrationResponse) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *IntegrationResponse) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *IntegrationResponse) GetModelName() datamodel.ModelNameType {
	return IntegrationResponseModelName
}

// NewIntegrationResponseID provides a template for generating an ID field for IntegrationResponse
func NewIntegrationResponseID(customerID string, refType string, refID string) string {
	return hash.Values("IntegrationResponse", customerID, refType, refID)
}

func (o *IntegrationResponse) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("IntegrationResponse", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *IntegrationResponse) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *IntegrationResponse) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *IntegrationResponse) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *IntegrationResponse) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *IntegrationResponse) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *IntegrationResponse) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *IntegrationResponse) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *IntegrationResponse) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *IntegrationResponse) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = IntegrationResponseModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *IntegrationResponse) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *IntegrationResponse) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of IntegrationResponse
func (o *IntegrationResponse) Clone() datamodel.Model {
	c := new(IntegrationResponse)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *IntegrationResponse) Anon() datamodel.Model {
	c := new(IntegrationResponse)
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
func (o *IntegrationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *IntegrationResponse) UnmarshalJSON(data []byte) error {
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
func (o *IntegrationResponse) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *IntegrationResponse) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two IntegrationResponse objects are equal
func (o *IntegrationResponse) IsEqual(other *IntegrationResponse) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *IntegrationResponse) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"architecture":            toIntegrationResponseObject(o.Architecture, false),
		"authorization":           toIntegrationResponseObject(o.Authorization, false),
		"customer_id":             toIntegrationResponseObject(o.CustomerID, false),
		"data":                    toIntegrationResponseObject(o.Data, true),
		"distro":                  toIntegrationResponseObject(o.Distro, false),
		"error":                   toIntegrationResponseObject(o.Error, true),
		"event_date":              toIntegrationResponseObject(o.EventDate, false),
		"free_space":              toIntegrationResponseObject(o.FreeSpace, false),
		"go_version":              toIntegrationResponseObject(o.GoVersion, false),
		"hostname":                toIntegrationResponseObject(o.Hostname, false),
		"id":                      toIntegrationResponseObject(o.ID, false),
		"integration_instance_id": toIntegrationResponseObject(o.IntegrationInstanceID, true),
		"last_export_date":        toIntegrationResponseObject(o.LastExportDate, false),
		"memory":                  toIntegrationResponseObject(o.Memory, false),
		"message":                 toIntegrationResponseObject(o.Message, false),
		"num_cpu":                 toIntegrationResponseObject(o.NumCPU, false),
		"os":                      toIntegrationResponseObject(o.OS, false),
		"ref_id":                  toIntegrationResponseObject(o.RefID, false),
		"ref_type":                toIntegrationResponseObject(o.RefType, false),
		"request_id":              toIntegrationResponseObject(o.RequestID, false),
		"server_version":          toIntegrationResponseObject(o.ServerVersion, true),
		"success":                 toIntegrationResponseObject(o.Success, false),
		"system_id":               toIntegrationResponseObject(o.SystemID, false),

		"type":     o.Type.String(),
		"uptime":   toIntegrationResponseObject(o.Uptime, false),
		"uuid":     toIntegrationResponseObject(o.UUID, false),
		"version":  toIntegrationResponseObject(o.Version, false),
		"hashcode": toIntegrationResponseObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationResponse) FromMap(kv map[string]interface{}) {

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
	if val, ok := kv["authorization"].(string); ok {
		o.Authorization = val
	} else {
		if val, ok := kv["authorization"]; ok {
			if val == nil {
				o.Authorization = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Authorization = fmt.Sprintf("%v", val)
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
		} else if sv, ok := val.(IntegrationResponseEventDate); ok {
			// struct
			o.EventDate = sv
		} else if sp, ok := val.(*IntegrationResponseEventDate); ok {
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
		} else if sv, ok := val.(IntegrationResponseLastExportDate); ok {
			// struct
			o.LastExportDate = sv
		} else if sp, ok := val.(*IntegrationResponseLastExportDate); ok {
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
	if val, ok := kv["server_version"].(*string); ok {
		o.ServerVersion = val
	} else if val, ok := kv["server_version"].(string); ok {
		o.ServerVersion = &val
	} else {
		if val, ok := kv["server_version"]; ok {
			if val == nil {
				o.ServerVersion = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.ServerVersion = pstrings.Pointer(fmt.Sprintf("%v", val))
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
	if val, ok := kv["type"].(IntegrationResponseType); ok {
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
func (o *IntegrationResponse) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Architecture)
	args = append(args, o.Authorization)
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
	args = append(args, o.ServerVersion)
	args = append(args, o.Success)
	args = append(args, o.SystemID)
	args = append(args, o.Type)
	args = append(args, o.Uptime)
	args = append(args, o.UUID)
	args = append(args, o.Version)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}
