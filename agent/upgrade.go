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
	// UpgradeTopic is the default topic name
	UpgradeTopic datamodel.TopicNameType = "agent_Upgrade_topic"

	// UpgradeModelName is the model name
	UpgradeModelName datamodel.ModelNameType = "agent.Upgrade"
)

const (
	// UpgradeArchitectureColumn is the architecture column name
	UpgradeArchitectureColumn = "Architecture"
	// UpgradeCustomerIDColumn is the customer_id column name
	UpgradeCustomerIDColumn = "CustomerID"
	// UpgradeDataColumn is the data column name
	UpgradeDataColumn = "Data"
	// UpgradeDistroColumn is the distro column name
	UpgradeDistroColumn = "Distro"
	// UpgradeErrorColumn is the error column name
	UpgradeErrorColumn = "Error"
	// UpgradeEventDateColumn is the event_date column name
	UpgradeEventDateColumn = "EventDate"
	// UpgradeEventDateColumnEpochColumn is the epoch column property of the EventDate name
	UpgradeEventDateColumnEpochColumn = "EventDate.Epoch"
	// UpgradeEventDateColumnOffsetColumn is the offset column property of the EventDate name
	UpgradeEventDateColumnOffsetColumn = "EventDate.Offset"
	// UpgradeEventDateColumnRfc3339Column is the rfc3339 column property of the EventDate name
	UpgradeEventDateColumnRfc3339Column = "EventDate.Rfc3339"
	// UpgradeFreeSpaceColumn is the free_space column name
	UpgradeFreeSpaceColumn = "FreeSpace"
	// UpgradeFromVersionColumn is the from_version column name
	UpgradeFromVersionColumn = "FromVersion"
	// UpgradeGoVersionColumn is the go_version column name
	UpgradeGoVersionColumn = "GoVersion"
	// UpgradeHostnameColumn is the hostname column name
	UpgradeHostnameColumn = "Hostname"
	// UpgradeIDColumn is the id column name
	UpgradeIDColumn = "ID"
	// UpgradeLastExportDateColumn is the last_export_date column name
	UpgradeLastExportDateColumn = "LastExportDate"
	// UpgradeLastExportDateColumnEpochColumn is the epoch column property of the LastExportDate name
	UpgradeLastExportDateColumnEpochColumn = "LastExportDate.Epoch"
	// UpgradeLastExportDateColumnOffsetColumn is the offset column property of the LastExportDate name
	UpgradeLastExportDateColumnOffsetColumn = "LastExportDate.Offset"
	// UpgradeLastExportDateColumnRfc3339Column is the rfc3339 column property of the LastExportDate name
	UpgradeLastExportDateColumnRfc3339Column = "LastExportDate.Rfc3339"
	// UpgradeMemoryColumn is the memory column name
	UpgradeMemoryColumn = "Memory"
	// UpgradeMessageColumn is the message column name
	UpgradeMessageColumn = "Message"
	// UpgradeNumCPUColumn is the num_cpu column name
	UpgradeNumCPUColumn = "NumCPU"
	// UpgradeOSColumn is the os column name
	UpgradeOSColumn = "OS"
	// UpgradeRefIDColumn is the ref_id column name
	UpgradeRefIDColumn = "RefID"
	// UpgradeRefTypeColumn is the ref_type column name
	UpgradeRefTypeColumn = "RefType"
	// UpgradeRequestIDColumn is the request_id column name
	UpgradeRequestIDColumn = "RequestID"
	// UpgradeSuccessColumn is the success column name
	UpgradeSuccessColumn = "Success"
	// UpgradeSystemIDColumn is the system_id column name
	UpgradeSystemIDColumn = "SystemID"
	// UpgradeToVersionColumn is the to_version column name
	UpgradeToVersionColumn = "ToVersion"
	// UpgradeTypeColumn is the type column name
	UpgradeTypeColumn = "Type"
	// UpgradeUpdatedAtColumn is the updated_ts column name
	UpgradeUpdatedAtColumn = "UpdatedAt"
	// UpgradeUptimeColumn is the uptime column name
	UpgradeUptimeColumn = "Uptime"
	// UpgradeUUIDColumn is the uuid column name
	UpgradeUUIDColumn = "UUID"
	// UpgradeVersionColumn is the version column name
	UpgradeVersionColumn = "Version"
)

// UpgradeEventDate represents the object structure for event_date
type UpgradeEventDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toUpgradeEventDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *UpgradeEventDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *UpgradeEventDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toUpgradeEventDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toUpgradeEventDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toUpgradeEventDateObject(o.Rfc3339, false),
	}
}

func (o *UpgradeEventDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *UpgradeEventDate) FromMap(kv map[string]interface{}) {

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

// UpgradeLastExportDate represents the object structure for last_export_date
type UpgradeLastExportDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toUpgradeLastExportDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *UpgradeLastExportDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *UpgradeLastExportDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toUpgradeLastExportDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toUpgradeLastExportDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toUpgradeLastExportDateObject(o.Rfc3339, false),
	}
}

func (o *UpgradeLastExportDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *UpgradeLastExportDate) FromMap(kv map[string]interface{}) {

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

// UpgradeType is the enumeration type for type
type UpgradeType int32

// UnmarshalJSON unmarshals the enum value
func (v UpgradeType) UnmarshalJSON(buf []byte) error {
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
	case "UNINSTALL":
		v = 9
	case "UPGRADE":
		v = 10
	case "START":
		v = 11
	case "STOP":
		v = 12
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v UpgradeType) MarshalJSON() ([]byte, error) {
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
		return json.Marshal("UNINSTALL")
	case 10:
		return json.Marshal("UPGRADE")
	case 11:
		return json.Marshal("START")
	case 12:
		return json.Marshal("STOP")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for Type
func (v UpgradeType) String() string {
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
	UpgradeTypeEnroll UpgradeType = 0
	// TypePing is the enumeration value for ping
	UpgradeTypePing UpgradeType = 1
	// TypeCrash is the enumeration value for crash
	UpgradeTypeCrash UpgradeType = 2
	// TypeLog is the enumeration value for log
	UpgradeTypeLog UpgradeType = 3
	// TypeIntegration is the enumeration value for integration
	UpgradeTypeIntegration UpgradeType = 4
	// TypeExport is the enumeration value for export
	UpgradeTypeExport UpgradeType = 5
	// TypeProject is the enumeration value for project
	UpgradeTypeProject UpgradeType = 6
	// TypeRepo is the enumeration value for repo
	UpgradeTypeRepo UpgradeType = 7
	// TypeUser is the enumeration value for user
	UpgradeTypeUser UpgradeType = 8
	// TypeUninstall is the enumeration value for uninstall
	UpgradeTypeUninstall UpgradeType = 9
	// TypeUpgrade is the enumeration value for upgrade
	UpgradeTypeUpgrade UpgradeType = 10
	// TypeStart is the enumeration value for start
	UpgradeTypeStart UpgradeType = 11
	// TypeStop is the enumeration value for stop
	UpgradeTypeStop UpgradeType = 12
)

// Upgrade an agent event to indicate the agent version was upgraded
type Upgrade struct {
	// Architecture the architecture of the agent machine
	Architecture string `json:"architecture" yaml:"architecture" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" yaml:"customer_id" faker:"-"`
	// Data extra data that is specific about this event
	Data *string `json:"data,omitempty" yaml:"data,omitempty" faker:"-"`
	// Distro the agent os distribution
	Distro string `json:"distro" yaml:"distro" faker:"-"`
	// Error an error message related to this event
	Error *string `json:"error,omitempty" yaml:"error,omitempty" faker:"-"`
	// EventDate the date of the event
	EventDate UpgradeEventDate `json:"event_date" yaml:"event_date" faker:"-"`
	// FreeSpace the amount of free space in bytes for the agent machine
	FreeSpace int64 `json:"free_space" yaml:"free_space" faker:"-"`
	// FromVersion the version we're upgrading
	FromVersion string `json:"from_version" yaml:"from_version" faker:"-"`
	// GoVersion the go version that the agent build was built with
	GoVersion string `json:"go_version" yaml:"go_version" faker:"-"`
	// Hostname the agent hostname
	Hostname string `json:"hostname" yaml:"hostname" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// LastExportDate the last export date
	LastExportDate UpgradeLastExportDate `json:"last_export_date" yaml:"last_export_date" faker:"-"`
	// Memory the amount of memory in bytes for the agent machine
	Memory int64 `json:"memory" yaml:"memory" faker:"-"`
	// Message a message related to this event
	Message string `json:"message" yaml:"message" faker:"-"`
	// NumCPU the number of CPU the agent is running
	NumCPU int64 `json:"num_cpu" yaml:"num_cpu" faker:"-"`
	// OS the agent operating system
	OS string `json:"os" yaml:"os" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" yaml:"ref_type" faker:"-"`
	// RequestID the request id that this response is correlated to
	RequestID string `json:"request_id" yaml:"request_id" faker:"-"`
	// Success if the response was successful
	Success bool `json:"success" yaml:"success" faker:"-"`
	// SystemID system unique device ID
	SystemID string `json:"system_id" yaml:"system_id" faker:"-"`
	// ToVersion the version we're installing
	ToVersion string `json:"to_version" yaml:"to_version" faker:"-"`
	// Type the type of event
	Type UpgradeType `json:"type" yaml:"type" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" yaml:"updated_ts" faker:"-"`
	// Uptime the uptime in milliseconds since the agent started
	Uptime int64 `json:"uptime" yaml:"uptime" faker:"-"`
	// UUID the agent unique identifier
	UUID string `json:"uuid" yaml:"uuid" faker:"-"`
	// Version the agent version
	Version string `json:"version" yaml:"version" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Upgrade)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*Upgrade)(nil)

func toUpgradeObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *Upgrade:
		return v.ToMap()

	case UpgradeEventDate:
		return v.ToMap()

	case UpgradeLastExportDate:
		return v.ToMap()

	case UpgradeType:
		return v.String()

	default:
		return o
	}
}

// String returns a string representation of Upgrade
func (o *Upgrade) String() string {
	return fmt.Sprintf("agent.Upgrade<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Upgrade) GetTopicName() datamodel.TopicNameType {
	return UpgradeTopic
}

// GetStreamName returns the name of the stream
func (o *Upgrade) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *Upgrade) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *Upgrade) GetModelName() datamodel.ModelNameType {
	return UpgradeModelName
}

// NewUpgradeID provides a template for generating an ID field for Upgrade
func NewUpgradeID(customerID string, refType string, refID string) string {
	return hash.Values("Upgrade", customerID, refType, refID)
}

func (o *Upgrade) setDefaults(frommap bool) {
	if o.Data == nil {
		o.Data = &emptyString
	}
	if o.Error == nil {
		o.Error = &emptyString
	}

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Upgrade", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Upgrade) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Upgrade) GetTopicKey() string {
	var i interface{} = o.UUID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Upgrade) GetTimestamp() time.Time {
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
	panic("not sure how to handle the date time format for Upgrade")
}

// GetRefID returns the RefID for the object
func (o *Upgrade) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Upgrade) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *Upgrade) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Upgrade) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Upgrade) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *Upgrade) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = UpgradeModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Upgrade) GetTopicConfig() *datamodel.ModelTopicConfig {
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
func (o *Upgrade) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of Upgrade
func (o *Upgrade) Clone() datamodel.Model {
	c := new(Upgrade)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Upgrade) Anon() datamodel.Model {
	c := new(Upgrade)
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
func (o *Upgrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Upgrade) UnmarshalJSON(data []byte) error {
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
func (o *Upgrade) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Upgrade objects are equal
func (o *Upgrade) IsEqual(other *Upgrade) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Upgrade) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"architecture":     toUpgradeObject(o.Architecture, false),
		"customer_id":      toUpgradeObject(o.CustomerID, false),
		"data":             toUpgradeObject(o.Data, true),
		"distro":           toUpgradeObject(o.Distro, false),
		"error":            toUpgradeObject(o.Error, true),
		"event_date":       toUpgradeObject(o.EventDate, false),
		"free_space":       toUpgradeObject(o.FreeSpace, false),
		"from_version":     toUpgradeObject(o.FromVersion, false),
		"go_version":       toUpgradeObject(o.GoVersion, false),
		"hostname":         toUpgradeObject(o.Hostname, false),
		"id":               toUpgradeObject(o.ID, false),
		"last_export_date": toUpgradeObject(o.LastExportDate, false),
		"memory":           toUpgradeObject(o.Memory, false),
		"message":          toUpgradeObject(o.Message, false),
		"num_cpu":          toUpgradeObject(o.NumCPU, false),
		"os":               toUpgradeObject(o.OS, false),
		"ref_id":           toUpgradeObject(o.RefID, false),
		"ref_type":         toUpgradeObject(o.RefType, false),
		"request_id":       toUpgradeObject(o.RequestID, false),
		"success":          toUpgradeObject(o.Success, false),
		"system_id":        toUpgradeObject(o.SystemID, false),
		"to_version":       toUpgradeObject(o.ToVersion, false),

		"type":       o.Type.String(),
		"updated_ts": toUpgradeObject(o.UpdatedAt, false),
		"uptime":     toUpgradeObject(o.Uptime, false),
		"uuid":       toUpgradeObject(o.UUID, false),
		"version":    toUpgradeObject(o.Version, false),
		"hashcode":   toUpgradeObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *Upgrade) FromMap(kv map[string]interface{}) {

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
		} else if sv, ok := val.(UpgradeEventDate); ok {
			// struct
			o.EventDate = sv
		} else if sp, ok := val.(*UpgradeEventDate); ok {
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

	if val, ok := kv["from_version"].(string); ok {
		o.FromVersion = val
	} else {
		if val, ok := kv["from_version"]; ok {
			if val == nil {
				o.FromVersion = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.FromVersion = fmt.Sprintf("%v", val)
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
		} else if sv, ok := val.(UpgradeLastExportDate); ok {
			// struct
			o.LastExportDate = sv
		} else if sp, ok := val.(*UpgradeLastExportDate); ok {
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

	if val, ok := kv["to_version"].(string); ok {
		o.ToVersion = val
	} else {
		if val, ok := kv["to_version"]; ok {
			if val == nil {
				o.ToVersion = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.ToVersion = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["type"].(UpgradeType); ok {
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
func (o *Upgrade) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Architecture)
	args = append(args, o.CustomerID)
	args = append(args, o.Data)
	args = append(args, o.Distro)
	args = append(args, o.Error)
	args = append(args, o.EventDate)
	args = append(args, o.FreeSpace)
	args = append(args, o.FromVersion)
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
	args = append(args, o.ToVersion)
	args = append(args, o.Type)
	args = append(args, o.UpdatedAt)
	args = append(args, o.Uptime)
	args = append(args, o.UUID)
	args = append(args, o.Version)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *Upgrade) GetEventAPIConfig() datamodel.EventAPIConfig {
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
