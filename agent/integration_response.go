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
)

const (
	// IntegrationResponseTopic is the default topic name
	IntegrationResponseTopic datamodel.TopicNameType = "agent_IntegrationResponse_topic"

	// IntegrationResponseModelName is the model name
	IntegrationResponseModelName datamodel.ModelNameType = "agent.IntegrationResponse"
)

const (
	// IntegrationResponseArchitectureColumn is the architecture column name
	IntegrationResponseArchitectureColumn = "Architecture"
	// IntegrationResponseAuthorizationColumn is the authorization column name
	IntegrationResponseAuthorizationColumn = "Authorization"
	// IntegrationResponseCustomerIDColumn is the customer_id column name
	IntegrationResponseCustomerIDColumn = "CustomerID"
	// IntegrationResponseDataColumn is the data column name
	IntegrationResponseDataColumn = "Data"
	// IntegrationResponseDistroColumn is the distro column name
	IntegrationResponseDistroColumn = "Distro"
	// IntegrationResponseErrorColumn is the error column name
	IntegrationResponseErrorColumn = "Error"
	// IntegrationResponseEventDateColumn is the event_date column name
	IntegrationResponseEventDateColumn = "EventDate"
	// IntegrationResponseEventDateColumnEpochColumn is the epoch column property of the EventDate name
	IntegrationResponseEventDateColumnEpochColumn = "EventDate.Epoch"
	// IntegrationResponseEventDateColumnOffsetColumn is the offset column property of the EventDate name
	IntegrationResponseEventDateColumnOffsetColumn = "EventDate.Offset"
	// IntegrationResponseEventDateColumnRfc3339Column is the rfc3339 column property of the EventDate name
	IntegrationResponseEventDateColumnRfc3339Column = "EventDate.Rfc3339"
	// IntegrationResponseFreeSpaceColumn is the free_space column name
	IntegrationResponseFreeSpaceColumn = "FreeSpace"
	// IntegrationResponseGoVersionColumn is the go_version column name
	IntegrationResponseGoVersionColumn = "GoVersion"
	// IntegrationResponseHostnameColumn is the hostname column name
	IntegrationResponseHostnameColumn = "Hostname"
	// IntegrationResponseIDColumn is the id column name
	IntegrationResponseIDColumn = "ID"
	// IntegrationResponseLastExportDateColumn is the last_export_date column name
	IntegrationResponseLastExportDateColumn = "LastExportDate"
	// IntegrationResponseLastExportDateColumnEpochColumn is the epoch column property of the LastExportDate name
	IntegrationResponseLastExportDateColumnEpochColumn = "LastExportDate.Epoch"
	// IntegrationResponseLastExportDateColumnOffsetColumn is the offset column property of the LastExportDate name
	IntegrationResponseLastExportDateColumnOffsetColumn = "LastExportDate.Offset"
	// IntegrationResponseLastExportDateColumnRfc3339Column is the rfc3339 column property of the LastExportDate name
	IntegrationResponseLastExportDateColumnRfc3339Column = "LastExportDate.Rfc3339"
	// IntegrationResponseMemoryColumn is the memory column name
	IntegrationResponseMemoryColumn = "Memory"
	// IntegrationResponseMessageColumn is the message column name
	IntegrationResponseMessageColumn = "Message"
	// IntegrationResponseNumCPUColumn is the num_cpu column name
	IntegrationResponseNumCPUColumn = "NumCPU"
	// IntegrationResponseOSColumn is the os column name
	IntegrationResponseOSColumn = "OS"
	// IntegrationResponseRefIDColumn is the ref_id column name
	IntegrationResponseRefIDColumn = "RefID"
	// IntegrationResponseRefTypeColumn is the ref_type column name
	IntegrationResponseRefTypeColumn = "RefType"
	// IntegrationResponseRequestIDColumn is the request_id column name
	IntegrationResponseRequestIDColumn = "RequestID"
	// IntegrationResponseSuccessColumn is the success column name
	IntegrationResponseSuccessColumn = "Success"
	// IntegrationResponseSystemIDColumn is the system_id column name
	IntegrationResponseSystemIDColumn = "SystemID"
	// IntegrationResponseTypeColumn is the type column name
	IntegrationResponseTypeColumn = "Type"
	// IntegrationResponseUpdatedAtColumn is the updated_ts column name
	IntegrationResponseUpdatedAtColumn = "UpdatedAt"
	// IntegrationResponseUptimeColumn is the uptime column name
	IntegrationResponseUptimeColumn = "Uptime"
	// IntegrationResponseUUIDColumn is the uuid column name
	IntegrationResponseUUIDColumn = "UUID"
	// IntegrationResponseVersionColumn is the version column name
	IntegrationResponseVersionColumn = "Version"
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
				o.Epoch = number.ToInt64Any(nil)
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
				o.Offset = number.ToInt64Any(nil)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				o.Epoch = number.ToInt64Any(nil)
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
				o.Offset = number.ToInt64Any(nil)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// IntegrationResponseType is the enumeration type for type
type IntegrationResponseType int32

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
		return "UNINSTALL"
	case 10:
		return "UPGRADE"
	case 11:
		return "START"
	case 12:
		return "STOP"
	}
	return "unset"
}

const (
	// TypeEnroll is the enumeration value for enroll
	IntegrationResponseTypeEnroll IntegrationResponseType = 0
	// TypePing is the enumeration value for ping
	IntegrationResponseTypePing IntegrationResponseType = 1
	// TypeCrash is the enumeration value for crash
	IntegrationResponseTypeCrash IntegrationResponseType = 2
	// TypeLog is the enumeration value for log
	IntegrationResponseTypeLog IntegrationResponseType = 3
	// TypeIntegration is the enumeration value for integration
	IntegrationResponseTypeIntegration IntegrationResponseType = 4
	// TypeExport is the enumeration value for export
	IntegrationResponseTypeExport IntegrationResponseType = 5
	// TypeProject is the enumeration value for project
	IntegrationResponseTypeProject IntegrationResponseType = 6
	// TypeRepo is the enumeration value for repo
	IntegrationResponseTypeRepo IntegrationResponseType = 7
	// TypeUser is the enumeration value for user
	IntegrationResponseTypeUser IntegrationResponseType = 8
	// TypeUninstall is the enumeration value for uninstall
	IntegrationResponseTypeUninstall IntegrationResponseType = 9
	// TypeUpgrade is the enumeration value for upgrade
	IntegrationResponseTypeUpgrade IntegrationResponseType = 10
	// TypeStart is the enumeration value for start
	IntegrationResponseTypeStart IntegrationResponseType = 11
	// TypeStop is the enumeration value for stop
	IntegrationResponseTypeStop IntegrationResponseType = 12
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
	// Success if the response was successful
	Success bool `json:"success" codec:"success" bson:"success" yaml:"success" faker:"-"`
	// SystemID system unique device ID
	SystemID string `json:"system_id" codec:"system_id" bson:"system_id" yaml:"system_id" faker:"-"`
	// Type the type of event
	Type IntegrationResponseType `json:"type" codec:"type" bson:"type" yaml:"type" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
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
	return IntegrationResponseTopic
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
	if o.Data == nil {
		o.Data = &emptyString
	}
	if o.Error == nil {
		o.Error = &emptyString
	}

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
	var i interface{} = o.UUID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *IntegrationResponse) GetTimestamp() time.Time {
	var dt interface{} = o.UpdatedAt
	switch v := dt.(type) {
	case int64:
		return datetime.DateFromEpoch(v).UTC()
	case string:
		tv, err := datetime.ISODateToTime(v)
		if err != nil {
			panic(err)
		}
		return tv.UTC()
	case time.Time:
		return v.UTC()
	}
	panic("not sure how to handle the date time format for IntegrationResponse")
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
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *IntegrationResponse) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = IntegrationResponseModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *IntegrationResponse) GetTopicConfig() *datamodel.ModelTopicConfig {
	retention, err := time.ParseDuration("87360h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 87360h0m0s. " + err.Error())
	}

	ttl, err := time.ParseDuration("0s")
	if err != nil {
		ttl = 0
	}
	if ttl == 0 && retention != 0 {
		ttl = retention // they should be the same if not set
	}
	return &datamodel.ModelTopicConfig{
		Key:               "uuid",
		Timestamp:         "updated_ts",
		NumPartitions:     8,
		CleanupPolicy:     datamodel.CleanupPolicy("compact"),
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
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

// IsEqual returns true if the two IntegrationResponse objects are equal
func (o *IntegrationResponse) IsEqual(other *IntegrationResponse) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *IntegrationResponse) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"architecture":     toIntegrationResponseObject(o.Architecture, false),
		"authorization":    toIntegrationResponseObject(o.Authorization, false),
		"customer_id":      toIntegrationResponseObject(o.CustomerID, false),
		"data":             toIntegrationResponseObject(o.Data, true),
		"distro":           toIntegrationResponseObject(o.Distro, false),
		"error":            toIntegrationResponseObject(o.Error, true),
		"event_date":       toIntegrationResponseObject(o.EventDate, false),
		"free_space":       toIntegrationResponseObject(o.FreeSpace, false),
		"go_version":       toIntegrationResponseObject(o.GoVersion, false),
		"hostname":         toIntegrationResponseObject(o.Hostname, false),
		"id":               toIntegrationResponseObject(o.ID, false),
		"last_export_date": toIntegrationResponseObject(o.LastExportDate, false),
		"memory":           toIntegrationResponseObject(o.Memory, false),
		"message":          toIntegrationResponseObject(o.Message, false),
		"num_cpu":          toIntegrationResponseObject(o.NumCPU, false),
		"os":               toIntegrationResponseObject(o.OS, false),
		"ref_id":           toIntegrationResponseObject(o.RefID, false),
		"ref_type":         toIntegrationResponseObject(o.RefType, false),
		"request_id":       toIntegrationResponseObject(o.RequestID, false),
		"success":          toIntegrationResponseObject(o.Success, false),
		"system_id":        toIntegrationResponseObject(o.SystemID, false),
		"type":             toIntegrationResponseObject(o.Type, false),
		"updated_ts":       toIntegrationResponseObject(o.UpdatedAt, false),
		"uptime":           toIntegrationResponseObject(o.Uptime, false),
		"uuid":             toIntegrationResponseObject(o.UUID, false),
		"version":          toIntegrationResponseObject(o.Version, false),
		"hashcode":         toIntegrationResponseObject(o.Hashcode, false),
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				o.FreeSpace = number.ToInt64Any(nil)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.ID = fmt.Sprintf("%v", val)
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
				o.Memory = number.ToInt64Any(nil)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				o.NumCPU = number.ToInt64Any(nil)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				o.Success = number.ToBoolAny(nil)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
			case "uninstall", "UNINSTALL":
				o.Type = 9
			case "upgrade", "UPGRADE":
				o.Type = 10
			case "start", "START":
				o.Type = 11
			case "stop", "STOP":
				o.Type = 12
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
			case "uninstall", "UNINSTALL":
				o.Type = 9
			case "upgrade", "UPGRADE":
				o.Type = 10
			case "start", "START":
				o.Type = 11
			case "stop", "STOP":
				o.Type = 12
			}
		}
	}

	if val, ok := kv["updated_ts"].(int64); ok {
		o.UpdatedAt = val
	} else {
		if val, ok := kv["updated_ts"]; ok {
			if val == nil {
				o.UpdatedAt = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.UpdatedAt = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["uptime"].(int64); ok {
		o.Uptime = val
	} else {
		if val, ok := kv["uptime"]; ok {
			if val == nil {
				o.Uptime = number.ToInt64Any(nil)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
	args = append(args, o.UpdatedAt)
	args = append(args, o.Uptime)
	args = append(args, o.UUID)
	args = append(args, o.Version)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *IntegrationResponse) GetEventAPIConfig() datamodel.EventAPIConfig {
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
