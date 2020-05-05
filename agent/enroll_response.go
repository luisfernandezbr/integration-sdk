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

	// EnrollResponseTable is the default table name
	EnrollResponseTable datamodel.ModelNameType = "agent_enrollresponse"

	// EnrollResponseModelName is the model name
	EnrollResponseModelName datamodel.ModelNameType = "agent.EnrollResponse"
)

const (
	// EnrollResponseModelApikeyColumn is the column json value apikey
	EnrollResponseModelApikeyColumn = "apikey"
	// EnrollResponseModelArchitectureColumn is the column json value architecture
	EnrollResponseModelArchitectureColumn = "architecture"
	// EnrollResponseModelCustomerIDColumn is the column json value customer_id
	EnrollResponseModelCustomerIDColumn = "customer_id"
	// EnrollResponseModelDataColumn is the column json value data
	EnrollResponseModelDataColumn = "data"
	// EnrollResponseModelDistroColumn is the column json value distro
	EnrollResponseModelDistroColumn = "distro"
	// EnrollResponseModelErrorColumn is the column json value error
	EnrollResponseModelErrorColumn = "error"
	// EnrollResponseModelEventDateColumn is the column json value event_date
	EnrollResponseModelEventDateColumn = "event_date"
	// EnrollResponseModelEventDateEpochColumn is the column json value epoch
	EnrollResponseModelEventDateEpochColumn = "epoch"
	// EnrollResponseModelEventDateOffsetColumn is the column json value offset
	EnrollResponseModelEventDateOffsetColumn = "offset"
	// EnrollResponseModelEventDateRfc3339Column is the column json value rfc3339
	EnrollResponseModelEventDateRfc3339Column = "rfc3339"
	// EnrollResponseModelFreeSpaceColumn is the column json value free_space
	EnrollResponseModelFreeSpaceColumn = "free_space"
	// EnrollResponseModelGoVersionColumn is the column json value go_version
	EnrollResponseModelGoVersionColumn = "go_version"
	// EnrollResponseModelHostnameColumn is the column json value hostname
	EnrollResponseModelHostnameColumn = "hostname"
	// EnrollResponseModelIDColumn is the column json value id
	EnrollResponseModelIDColumn = "id"
	// EnrollResponseModelLastExportDateColumn is the column json value last_export_date
	EnrollResponseModelLastExportDateColumn = "last_export_date"
	// EnrollResponseModelLastExportDateEpochColumn is the column json value epoch
	EnrollResponseModelLastExportDateEpochColumn = "epoch"
	// EnrollResponseModelLastExportDateOffsetColumn is the column json value offset
	EnrollResponseModelLastExportDateOffsetColumn = "offset"
	// EnrollResponseModelLastExportDateRfc3339Column is the column json value rfc3339
	EnrollResponseModelLastExportDateRfc3339Column = "rfc3339"
	// EnrollResponseModelMemoryColumn is the column json value memory
	EnrollResponseModelMemoryColumn = "memory"
	// EnrollResponseModelMessageColumn is the column json value message
	EnrollResponseModelMessageColumn = "message"
	// EnrollResponseModelNumCPUColumn is the column json value num_cpu
	EnrollResponseModelNumCPUColumn = "num_cpu"
	// EnrollResponseModelOSColumn is the column json value os
	EnrollResponseModelOSColumn = "os"
	// EnrollResponseModelRefIDColumn is the column json value ref_id
	EnrollResponseModelRefIDColumn = "ref_id"
	// EnrollResponseModelRefTypeColumn is the column json value ref_type
	EnrollResponseModelRefTypeColumn = "ref_type"
	// EnrollResponseModelRequestIDColumn is the column json value request_id
	EnrollResponseModelRequestIDColumn = "request_id"
	// EnrollResponseModelSuccessColumn is the column json value success
	EnrollResponseModelSuccessColumn = "success"
	// EnrollResponseModelSystemIDColumn is the column json value system_id
	EnrollResponseModelSystemIDColumn = "system_id"
	// EnrollResponseModelTypeColumn is the column json value type
	EnrollResponseModelTypeColumn = "type"
	// EnrollResponseModelUptimeColumn is the column json value uptime
	EnrollResponseModelUptimeColumn = "uptime"
	// EnrollResponseModelUUIDColumn is the column json value uuid
	EnrollResponseModelUUIDColumn = "uuid"
	// EnrollResponseModelVersionColumn is the column json value version
	EnrollResponseModelVersionColumn = "version"
)

// EnrollResponseEventDate represents the object structure for event_date
type EnrollResponseEventDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toEnrollResponseEventDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *EnrollResponseEventDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *EnrollResponseEventDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toEnrollResponseEventDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toEnrollResponseEventDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toEnrollResponseEventDateObject(o.Rfc3339, false),
	}
}

func (o *EnrollResponseEventDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *EnrollResponseEventDate) FromMap(kv map[string]interface{}) {

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

// EnrollResponseLastExportDate represents the object structure for last_export_date
type EnrollResponseLastExportDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toEnrollResponseLastExportDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *EnrollResponseLastExportDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *EnrollResponseLastExportDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toEnrollResponseLastExportDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toEnrollResponseLastExportDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toEnrollResponseLastExportDateObject(o.Rfc3339, false),
	}
}

func (o *EnrollResponseLastExportDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *EnrollResponseLastExportDate) FromMap(kv map[string]interface{}) {

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

// EnrollResponseType is the enumeration type for type
type EnrollResponseType int32

// UnmarshalBSONValue for unmarshaling value
func (v *EnrollResponseType) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = EnrollResponseType(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "ENROLL":
			*v = EnrollResponseType(0)
		case "PING":
			*v = EnrollResponseType(1)
		case "CRASH":
			*v = EnrollResponseType(2)
		case "LOG":
			*v = EnrollResponseType(3)
		case "INTEGRATION":
			*v = EnrollResponseType(4)
		case "EXPORT":
			*v = EnrollResponseType(5)
		case "PROJECT":
			*v = EnrollResponseType(6)
		case "REPO":
			*v = EnrollResponseType(7)
		case "USER":
			*v = EnrollResponseType(8)
		case "CALENDAR":
			*v = EnrollResponseType(9)
		case "UNINSTALL":
			*v = EnrollResponseType(10)
		case "UPGRADE":
			*v = EnrollResponseType(11)
		case "START":
			*v = EnrollResponseType(12)
		case "STOP":
			*v = EnrollResponseType(13)
		case "PAUSE":
			*v = EnrollResponseType(14)
		case "RESUME":
			*v = EnrollResponseType(15)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v EnrollResponseType) UnmarshalJSON(buf []byte) error {
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
func (v EnrollResponseType) MarshalJSON() ([]byte, error) {
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
func (v EnrollResponseType) String() string {
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
func (v *EnrollResponseType) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case int32:
		*v = EnrollResponseType(int32(val))
	case int:
		*v = EnrollResponseType(int32(val))
	case string:
		switch val {
		case "ENROLL":
			*v = EnrollResponseType(0)
		case "PING":
			*v = EnrollResponseType(1)
		case "CRASH":
			*v = EnrollResponseType(2)
		case "LOG":
			*v = EnrollResponseType(3)
		case "INTEGRATION":
			*v = EnrollResponseType(4)
		case "EXPORT":
			*v = EnrollResponseType(5)
		case "PROJECT":
			*v = EnrollResponseType(6)
		case "REPO":
			*v = EnrollResponseType(7)
		case "USER":
			*v = EnrollResponseType(8)
		case "CALENDAR":
			*v = EnrollResponseType(9)
		case "UNINSTALL":
			*v = EnrollResponseType(10)
		case "UPGRADE":
			*v = EnrollResponseType(11)
		case "START":
			*v = EnrollResponseType(12)
		case "STOP":
			*v = EnrollResponseType(13)
		case "PAUSE":
			*v = EnrollResponseType(14)
		case "RESUME":
			*v = EnrollResponseType(15)
		}
	}
	return nil
}

const (
	// EnrollResponseTypeEnroll is the enumeration value for enroll
	EnrollResponseTypeEnroll EnrollResponseType = 0
	// EnrollResponseTypePing is the enumeration value for ping
	EnrollResponseTypePing EnrollResponseType = 1
	// EnrollResponseTypeCrash is the enumeration value for crash
	EnrollResponseTypeCrash EnrollResponseType = 2
	// EnrollResponseTypeLog is the enumeration value for log
	EnrollResponseTypeLog EnrollResponseType = 3
	// EnrollResponseTypeIntegration is the enumeration value for integration
	EnrollResponseTypeIntegration EnrollResponseType = 4
	// EnrollResponseTypeExport is the enumeration value for export
	EnrollResponseTypeExport EnrollResponseType = 5
	// EnrollResponseTypeProject is the enumeration value for project
	EnrollResponseTypeProject EnrollResponseType = 6
	// EnrollResponseTypeRepo is the enumeration value for repo
	EnrollResponseTypeRepo EnrollResponseType = 7
	// EnrollResponseTypeUser is the enumeration value for user
	EnrollResponseTypeUser EnrollResponseType = 8
	// EnrollResponseTypeCalendar is the enumeration value for calendar
	EnrollResponseTypeCalendar EnrollResponseType = 9
	// EnrollResponseTypeUninstall is the enumeration value for uninstall
	EnrollResponseTypeUninstall EnrollResponseType = 10
	// EnrollResponseTypeUpgrade is the enumeration value for upgrade
	EnrollResponseTypeUpgrade EnrollResponseType = 11
	// EnrollResponseTypeStart is the enumeration value for start
	EnrollResponseTypeStart EnrollResponseType = 12
	// EnrollResponseTypeStop is the enumeration value for stop
	EnrollResponseTypeStop EnrollResponseType = 13
	// EnrollResponseTypePause is the enumeration value for pause
	EnrollResponseTypePause EnrollResponseType = 14
	// EnrollResponseTypeResume is the enumeration value for resume
	EnrollResponseTypeResume EnrollResponseType = 15
)

// EnrollResponse an agent response to an the enroll action
type EnrollResponse struct {
	// Apikey The API key that the agent should use when communicating with the cloud
	Apikey string `json:"apikey" codec:"apikey" bson:"apikey" yaml:"apikey" faker:"-"`
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
	EventDate EnrollResponseEventDate `json:"event_date" codec:"event_date" bson:"event_date" yaml:"event_date" faker:"-"`
	// FreeSpace the amount of free space in bytes for the agent machine
	FreeSpace int64 `json:"free_space" codec:"free_space" bson:"free_space" yaml:"free_space" faker:"-"`
	// GoVersion the go version that the agent build was built with
	GoVersion string `json:"go_version" codec:"go_version" bson:"go_version" yaml:"go_version" faker:"-"`
	// Hostname the agent hostname
	Hostname string `json:"hostname" codec:"hostname" bson:"hostname" yaml:"hostname" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// LastExportDate the last export date
	LastExportDate EnrollResponseLastExportDate `json:"last_export_date" codec:"last_export_date" bson:"last_export_date" yaml:"last_export_date" faker:"-"`
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
	Type EnrollResponseType `json:"type" codec:"type" bson:"type" yaml:"type" faker:"-"`
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
var _ datamodel.Model = (*EnrollResponse)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*EnrollResponse)(nil)

func toEnrollResponseObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *EnrollResponse:
		return v.ToMap()

	case EnrollResponseEventDate:
		return v.ToMap()

	case EnrollResponseLastExportDate:
		return v.ToMap()

	case EnrollResponseType:
		return v.String()

	default:
		return o
	}
}

// String returns a string representation of EnrollResponse
func (o *EnrollResponse) String() string {
	return fmt.Sprintf("agent.EnrollResponse<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *EnrollResponse) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *EnrollResponse) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *EnrollResponse) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *EnrollResponse) GetModelName() datamodel.ModelNameType {
	return EnrollResponseModelName
}

// NewEnrollResponseID provides a template for generating an ID field for EnrollResponse
func NewEnrollResponseID(customerID string, refType string, refID string) string {
	return hash.Values("EnrollResponse", customerID, refType, refID)
}

func (o *EnrollResponse) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("EnrollResponse", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *EnrollResponse) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *EnrollResponse) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *EnrollResponse) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *EnrollResponse) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *EnrollResponse) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *EnrollResponse) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *EnrollResponse) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *EnrollResponse) IsEvented() bool {
	return false
}

// GetTopicConfig returns the topic config object
func (o *EnrollResponse) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *EnrollResponse) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of EnrollResponse
func (o *EnrollResponse) Clone() datamodel.Model {
	c := new(EnrollResponse)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *EnrollResponse) Anon() datamodel.Model {
	c := new(EnrollResponse)
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
func (o *EnrollResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *EnrollResponse) UnmarshalJSON(data []byte) error {
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
func (o *EnrollResponse) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *EnrollResponse) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two EnrollResponse objects are equal
func (o *EnrollResponse) IsEqual(other *EnrollResponse) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *EnrollResponse) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"apikey":           toEnrollResponseObject(o.Apikey, false),
		"architecture":     toEnrollResponseObject(o.Architecture, false),
		"customer_id":      toEnrollResponseObject(o.CustomerID, false),
		"data":             toEnrollResponseObject(o.Data, true),
		"distro":           toEnrollResponseObject(o.Distro, false),
		"error":            toEnrollResponseObject(o.Error, true),
		"event_date":       toEnrollResponseObject(o.EventDate, false),
		"free_space":       toEnrollResponseObject(o.FreeSpace, false),
		"go_version":       toEnrollResponseObject(o.GoVersion, false),
		"hostname":         toEnrollResponseObject(o.Hostname, false),
		"id":               toEnrollResponseObject(o.ID, false),
		"last_export_date": toEnrollResponseObject(o.LastExportDate, false),
		"memory":           toEnrollResponseObject(o.Memory, false),
		"message":          toEnrollResponseObject(o.Message, false),
		"num_cpu":          toEnrollResponseObject(o.NumCPU, false),
		"os":               toEnrollResponseObject(o.OS, false),
		"ref_id":           toEnrollResponseObject(o.RefID, false),
		"ref_type":         toEnrollResponseObject(o.RefType, false),
		"request_id":       toEnrollResponseObject(o.RequestID, false),
		"success":          toEnrollResponseObject(o.Success, false),
		"system_id":        toEnrollResponseObject(o.SystemID, false),

		"type":     o.Type.String(),
		"uptime":   toEnrollResponseObject(o.Uptime, false),
		"uuid":     toEnrollResponseObject(o.UUID, false),
		"version":  toEnrollResponseObject(o.Version, false),
		"hashcode": toEnrollResponseObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *EnrollResponse) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["apikey"].(string); ok {
		o.Apikey = val
	} else {
		if val, ok := kv["apikey"]; ok {
			if val == nil {
				o.Apikey = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Apikey = fmt.Sprintf("%v", val)
			}
		}
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
		} else if sv, ok := val.(EnrollResponseEventDate); ok {
			// struct
			o.EventDate = sv
		} else if sp, ok := val.(*EnrollResponseEventDate); ok {
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
		} else if sv, ok := val.(EnrollResponseLastExportDate); ok {
			// struct
			o.LastExportDate = sv
		} else if sp, ok := val.(*EnrollResponseLastExportDate); ok {
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

	if val, ok := kv["type"].(EnrollResponseType); ok {
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
func (o *EnrollResponse) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Apikey)
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
func (o *EnrollResponse) GetEventAPIConfig() datamodel.EventAPIConfig {
	return datamodel.EventAPIConfig{
		Publish: datamodel.EventAPIPublish{
			Public: false,
		},
		Subscribe: datamodel.EventAPISubscribe{
			Public: true,
			Key:    "uuid",
		},
	}
}
