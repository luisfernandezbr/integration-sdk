// DO NOT EDIT -- generated code

// Package agent - the agent communicates with pinpoint cloud to send integration data
package agent

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"sync"
	"time"

	"github.com/bxcodec/faker"
	"github.com/linkedin/goavro"
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/eventing"
	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
	pstrings "github.com/pinpt/go-common/strings"
)

const (
	// UninstallTopic is the default topic name
	UninstallTopic datamodel.TopicNameType = "agent_Uninstall_topic"

	// UninstallStream is the default stream name
	UninstallStream datamodel.TopicNameType = "agent_Uninstall_stream"

	// UninstallTable is the default table name
	UninstallTable datamodel.TopicNameType = "agent_Uninstall"

	// UninstallModelName is the model name
	UninstallModelName datamodel.ModelNameType = "agent.Uninstall"
)

const (
	// UninstallArchitectureColumn is the architecture column name
	UninstallArchitectureColumn = "architecture"
	// UninstallCustomerIDColumn is the customer_id column name
	UninstallCustomerIDColumn = "customer_id"
	// UninstallDataColumn is the data column name
	UninstallDataColumn = "data"
	// UninstallDistroColumn is the distro column name
	UninstallDistroColumn = "distro"
	// UninstallErrorColumn is the error column name
	UninstallErrorColumn = "error"
	// UninstallEventDateColumn is the event_date column name
	UninstallEventDateColumn = "event_date"
	// UninstallEventDateColumnEpochColumn is the epoch column property of the EventDate name
	UninstallEventDateColumnEpochColumn = "event_date->epoch"
	// UninstallEventDateColumnOffsetColumn is the offset column property of the EventDate name
	UninstallEventDateColumnOffsetColumn = "event_date->offset"
	// UninstallEventDateColumnRfc3339Column is the rfc3339 column property of the EventDate name
	UninstallEventDateColumnRfc3339Column = "event_date->rfc3339"
	// UninstallFreeSpaceColumn is the free_space column name
	UninstallFreeSpaceColumn = "free_space"
	// UninstallGoVersionColumn is the go_version column name
	UninstallGoVersionColumn = "go_version"
	// UninstallHostnameColumn is the hostname column name
	UninstallHostnameColumn = "hostname"
	// UninstallIDColumn is the id column name
	UninstallIDColumn = "id"
	// UninstallLastExportDateColumn is the last_export_date column name
	UninstallLastExportDateColumn = "last_export_date"
	// UninstallLastExportDateColumnEpochColumn is the epoch column property of the LastExportDate name
	UninstallLastExportDateColumnEpochColumn = "last_export_date->epoch"
	// UninstallLastExportDateColumnOffsetColumn is the offset column property of the LastExportDate name
	UninstallLastExportDateColumnOffsetColumn = "last_export_date->offset"
	// UninstallLastExportDateColumnRfc3339Column is the rfc3339 column property of the LastExportDate name
	UninstallLastExportDateColumnRfc3339Column = "last_export_date->rfc3339"
	// UninstallMemoryColumn is the memory column name
	UninstallMemoryColumn = "memory"
	// UninstallMessageColumn is the message column name
	UninstallMessageColumn = "message"
	// UninstallNumCPUColumn is the num_cpu column name
	UninstallNumCPUColumn = "num_cpu"
	// UninstallOSColumn is the os column name
	UninstallOSColumn = "os"
	// UninstallRefIDColumn is the ref_id column name
	UninstallRefIDColumn = "ref_id"
	// UninstallRefTypeColumn is the ref_type column name
	UninstallRefTypeColumn = "ref_type"
	// UninstallRequestIDColumn is the request_id column name
	UninstallRequestIDColumn = "request_id"
	// UninstallSuccessColumn is the success column name
	UninstallSuccessColumn = "success"
	// UninstallTypeColumn is the type column name
	UninstallTypeColumn = "type"
	// UninstallUpdatedAtColumn is the updated_ts column name
	UninstallUpdatedAtColumn = "updated_ts"
	// UninstallUptimeColumn is the uptime column name
	UninstallUptimeColumn = "uptime"
	// UninstallUUIDColumn is the uuid column name
	UninstallUUIDColumn = "uuid"
	// UninstallVersionColumn is the version column name
	UninstallVersionColumn = "version"
)

// UninstallEventDate represents the object structure for event_date
type UninstallEventDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toUninstallEventDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toUninstallEventDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *UninstallEventDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *UninstallEventDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toUninstallEventDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toUninstallEventDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toUninstallEventDateObject(o.Rfc3339, isavro, false, "string"),
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

// UninstallLastExportDate represents the object structure for last_export_date
type UninstallLastExportDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toUninstallLastExportDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toUninstallLastExportDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *UninstallLastExportDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *UninstallLastExportDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toUninstallLastExportDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toUninstallLastExportDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toUninstallLastExportDateObject(o.Rfc3339, isavro, false, "string"),
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

// UninstallType is the enumeration type for type
type UninstallType int32

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
	UninstallTypeEnroll UninstallType = 0
	// TypePing is the enumeration value for ping
	UninstallTypePing UninstallType = 1
	// TypeCrash is the enumeration value for crash
	UninstallTypeCrash UninstallType = 2
	// TypeLog is the enumeration value for log
	UninstallTypeLog UninstallType = 3
	// TypeIntegration is the enumeration value for integration
	UninstallTypeIntegration UninstallType = 4
	// TypeExport is the enumeration value for export
	UninstallTypeExport UninstallType = 5
	// TypeProject is the enumeration value for project
	UninstallTypeProject UninstallType = 6
	// TypeRepo is the enumeration value for repo
	UninstallTypeRepo UninstallType = 7
	// TypeUser is the enumeration value for user
	UninstallTypeUser UninstallType = 8
	// TypeUninstall is the enumeration value for uninstall
	UninstallTypeUninstall UninstallType = 9
	// TypeUpgrade is the enumeration value for upgrade
	UninstallTypeUpgrade UninstallType = 10
	// TypeStart is the enumeration value for start
	UninstallTypeStart UninstallType = 11
	// TypeStop is the enumeration value for stop
	UninstallTypeStop UninstallType = 12
)

// Uninstall an agent event to indicate the agent was uninstalled
type Uninstall struct {
	// Architecture the architecture of the agent machine
	Architecture string `json:"architecture" bson:"architecture" yaml:"architecture" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Data extra data that is specific about this event
	Data *string `json:"data" bson:"data" yaml:"data" faker:"-"`
	// Distro the agent os distribution
	Distro string `json:"distro" bson:"distro" yaml:"distro" faker:"-"`
	// Error an error message related to this event
	Error *string `json:"error" bson:"error" yaml:"error" faker:"-"`
	// EventDate the date of the event
	EventDate UninstallEventDate `json:"event_date" bson:"event_date" yaml:"event_date" faker:"-"`
	// FreeSpace the amount of free space in bytes for the agent machine
	FreeSpace int64 `json:"free_space" bson:"free_space" yaml:"free_space" faker:"-"`
	// GoVersion the go version that the agent build was built with
	GoVersion string `json:"go_version" bson:"go_version" yaml:"go_version" faker:"-"`
	// Hostname the agent hostname
	Hostname string `json:"hostname" bson:"hostname" yaml:"hostname" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// LastExportDate the last export date
	LastExportDate UninstallLastExportDate `json:"last_export_date" bson:"last_export_date" yaml:"last_export_date" faker:"-"`
	// Memory the amount of memory in bytes for the agent machine
	Memory int64 `json:"memory" bson:"memory" yaml:"memory" faker:"-"`
	// Message a message related to this event
	Message string `json:"message" bson:"message" yaml:"message" faker:"-"`
	// NumCPU the number of CPU the agent is running
	NumCPU int64 `json:"num_cpu" bson:"num_cpu" yaml:"num_cpu" faker:"-"`
	// OS the agent operating system
	OS string `json:"os" bson:"os" yaml:"os" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RequestID the request id that this response is correlated to
	RequestID string `json:"request_id" bson:"request_id" yaml:"request_id" faker:"-"`
	// Success if the response was successful
	Success bool `json:"success" bson:"success" yaml:"success" faker:"-"`
	// Type the type of event
	Type UninstallType `json:"type" bson:"type" yaml:"type" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// Uptime the uptime in milliseconds since the agent started
	Uptime int64 `json:"uptime" bson:"uptime" yaml:"uptime" faker:"-"`
	// UUID the agent unique identifier
	UUID string `json:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
	// Version the agent version
	Version string `json:"version" bson:"version" yaml:"version" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Uninstall)(nil)

func toUninstallObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toUninstallObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *Uninstall:
		return v.ToMap(isavro)

	case UninstallEventDate:
		return v.ToMap(isavro)

	case UninstallLastExportDate:
		return v.ToMap(isavro)

	case UninstallType:
		return v.String()

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

// String returns a string representation of Uninstall
func (o *Uninstall) String() string {
	return fmt.Sprintf("agent.Uninstall<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Uninstall) GetTopicName() datamodel.TopicNameType {
	return UninstallTopic
}

// GetModelName returns the name of the model
func (o *Uninstall) GetModelName() datamodel.ModelNameType {
	return UninstallModelName
}

func (o *Uninstall) setDefaults(frommap bool) {
	if o.Data == nil {
		o.Data = &emptyString
	}
	if o.Error == nil {
		o.Error = &emptyString
	}

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
	var i interface{} = o.UUID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Uninstall) GetTimestamp() time.Time {
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
	panic("not sure how to handle the date time format for Uninstall")
}

// GetRefID returns the RefID for the object
func (o *Uninstall) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Uninstall) IsMaterialized() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Uninstall) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Uninstall) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *Uninstall) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = UninstallModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Uninstall) GetTopicConfig() *datamodel.ModelTopicConfig {
	retention, err := time.ParseDuration("168h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 168h0m0s. " + err.Error())
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
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetStateKey returns a key for use in state store
func (o *Uninstall) GetStateKey() string {
	key := "uuid"
	return fmt.Sprintf("%s_%s", key, o.GetID())
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

var cachedCodecUninstall *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *Uninstall) GetAvroCodec() *goavro.Codec {
	if cachedCodecUninstall == nil {
		c, err := GetUninstallAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecUninstall = c
	}
	return cachedCodecUninstall
}

// ToAvroBinary returns the data as Avro binary data
func (o *Uninstall) ToAvroBinary() ([]byte, *goavro.Codec, error) {
	kv := o.ToMap(true)
	jbuf, _ := json.Marshal(kv)
	codec := o.GetAvroCodec()
	native, _, err := codec.NativeFromTextual(jbuf)
	if err != nil {
		return nil, nil, err
	}
	// Convert native Go form to binary Avro data
	buf, err := codec.BinaryFromNative(nil, native)
	return buf, codec, err
}

// FromAvroBinary will convert from Avro binary data into data in this object
func (o *Uninstall) FromAvroBinary(value []byte) error {
	var nullHeader = []byte{byte(0)}
	// if this still has the schema encoded in the header, move past it to the avro payload
	if bytes.HasPrefix(value, nullHeader) {
		value = value[5:]
	}
	kv, _, err := o.GetAvroCodec().NativeFromBinary(value)
	if err != nil {
		return err
	}
	o.FromMap(kv.(map[string]interface{}))
	return nil
}

// Stringify returns the object in JSON format as a string
func (o *Uninstall) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Uninstall objects are equal
func (o *Uninstall) IsEqual(other *Uninstall) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Uninstall) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	o.setDefaults(false)
	return map[string]interface{}{
		"architecture":     toUninstallObject(o.Architecture, isavro, false, "string"),
		"customer_id":      toUninstallObject(o.CustomerID, isavro, false, "string"),
		"data":             toUninstallObject(o.Data, isavro, true, "string"),
		"distro":           toUninstallObject(o.Distro, isavro, false, "string"),
		"error":            toUninstallObject(o.Error, isavro, true, "string"),
		"event_date":       toUninstallObject(o.EventDate, isavro, false, "event_date"),
		"free_space":       toUninstallObject(o.FreeSpace, isavro, false, "long"),
		"go_version":       toUninstallObject(o.GoVersion, isavro, false, "string"),
		"hostname":         toUninstallObject(o.Hostname, isavro, false, "string"),
		"id":               toUninstallObject(o.ID, isavro, false, "string"),
		"last_export_date": toUninstallObject(o.LastExportDate, isavro, false, "last_export_date"),
		"memory":           toUninstallObject(o.Memory, isavro, false, "long"),
		"message":          toUninstallObject(o.Message, isavro, false, "string"),
		"num_cpu":          toUninstallObject(o.NumCPU, isavro, false, "long"),
		"os":               toUninstallObject(o.OS, isavro, false, "string"),
		"ref_id":           toUninstallObject(o.RefID, isavro, false, "string"),
		"ref_type":         toUninstallObject(o.RefType, isavro, false, "string"),
		"request_id":       toUninstallObject(o.RequestID, isavro, false, "string"),
		"success":          toUninstallObject(o.Success, isavro, false, "boolean"),
		"type":             toUninstallObject(o.Type, isavro, false, "type"),
		"updated_ts":       toUninstallObject(o.UpdatedAt, isavro, false, "long"),
		"uptime":           toUninstallObject(o.Uptime, isavro, false, "long"),
		"uuid":             toUninstallObject(o.UUID, isavro, false, "string"),
		"version":          toUninstallObject(o.Version, isavro, false, "string"),
		"hashcode":         toUninstallObject(o.Hashcode, isavro, false, "string"),
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
				// if coming in as avro union, convert it back
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
				// if coming in as avro union, convert it back
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
		} else if sv, ok := val.(UninstallLastExportDate); ok {
			// struct
			o.LastExportDate = sv
		} else if sp, ok := val.(*UninstallLastExportDate); ok {
			// struct pointer
			o.LastExportDate = *sp
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
	args = append(args, o.LastExportDate)
	args = append(args, o.Memory)
	args = append(args, o.Message)
	args = append(args, o.NumCPU)
	args = append(args, o.OS)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RequestID)
	args = append(args, o.Success)
	args = append(args, o.Type)
	args = append(args, o.UpdatedAt)
	args = append(args, o.Uptime)
	args = append(args, o.UUID)
	args = append(args, o.Version)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetUninstallAvroSchemaSpec creates the avro schema specification for Uninstall
func GetUninstallAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "agent",
		"name":      "Uninstall",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "hashcode",
				"type": "string",
			},
			map[string]interface{}{
				"name": "architecture",
				"type": "string",
			},
			map[string]interface{}{
				"name": "customer_id",
				"type": "string",
			},
			map[string]interface{}{
				"name":    "data",
				"type":    []interface{}{"null", "string"},
				"default": nil,
			},
			map[string]interface{}{
				"name": "distro",
				"type": "string",
			},
			map[string]interface{}{
				"name":    "error",
				"type":    []interface{}{"null", "string"},
				"default": nil,
			},
			map[string]interface{}{
				"name": "event_date",
				"type": map[string]interface{}{"doc": "the date of the event", "fields": []interface{}{map[string]interface{}{"doc": "the date in epoch format", "name": "epoch", "type": "long"}, map[string]interface{}{"doc": "the timezone offset from GMT", "name": "offset", "type": "long"}, map[string]interface{}{"doc": "the date in RFC3339 format", "name": "rfc3339", "type": "string"}}, "name": "event_date", "type": "record"},
			},
			map[string]interface{}{
				"name": "free_space",
				"type": "long",
			},
			map[string]interface{}{
				"name": "go_version",
				"type": "string",
			},
			map[string]interface{}{
				"name": "hostname",
				"type": "string",
			},
			map[string]interface{}{
				"name": "id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "last_export_date",
				"type": map[string]interface{}{"doc": "the last export date", "fields": []interface{}{map[string]interface{}{"doc": "the date in epoch format", "name": "epoch", "type": "long"}, map[string]interface{}{"doc": "the timezone offset from GMT", "name": "offset", "type": "long"}, map[string]interface{}{"doc": "the date in RFC3339 format", "name": "rfc3339", "type": "string"}}, "name": "last_export_date", "type": "record"},
			},
			map[string]interface{}{
				"name": "memory",
				"type": "long",
			},
			map[string]interface{}{
				"name": "message",
				"type": "string",
			},
			map[string]interface{}{
				"name": "num_cpu",
				"type": "long",
			},
			map[string]interface{}{
				"name": "os",
				"type": "string",
			},
			map[string]interface{}{
				"name": "ref_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "ref_type",
				"type": "string",
			},
			map[string]interface{}{
				"name": "request_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "success",
				"type": "boolean",
			},
			map[string]interface{}{
				"name": "type",
				"type": map[string]interface{}{
					"type":    "enum",
					"name":    "type",
					"symbols": []interface{}{"ENROLL", "PING", "CRASH", "LOG", "INTEGRATION", "EXPORT", "PROJECT", "REPO", "USER", "UNINSTALL", "UPGRADE", "START", "STOP"},
				},
			},
			map[string]interface{}{
				"name": "updated_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "uptime",
				"type": "long",
			},
			map[string]interface{}{
				"name": "uuid",
				"type": "string",
			},
			map[string]interface{}{
				"name": "version",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *Uninstall) GetEventAPIConfig() datamodel.EventAPIConfig {
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

// GetUninstallAvroSchema creates the avro schema for Uninstall
func GetUninstallAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetUninstallAvroSchemaSpec())
}

// TransformUninstallFunc is a function for transforming Uninstall during processing
type TransformUninstallFunc func(input *Uninstall) (*Uninstall, error)

// NewUninstallPipe creates a pipe for processing Uninstall items
func NewUninstallPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformUninstallFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewUninstallInputStream(input, errors)
	var stream chan Uninstall
	if len(transforms) > 0 {
		stream = make(chan Uninstall, 1000)
	} else {
		stream = inch
	}
	outdone := NewUninstallOutputStream(output, stream, errors)
	go func() {
		if len(transforms) > 0 {
			var stop bool
			for item := range inch {
				input := &item
				for _, transform := range transforms {
					out, err := transform(input)
					if err != nil {
						stop = true
						errors <- err
						break
					}
					if out == nil {
						input = nil
						break
					} else {
						input = out
					}
				}
				if stop {
					break
				}
				if input != nil {
					stream <- *input
				}
			}
			close(stream)
		}
		<-indone
		<-outdone
		done <- true
	}()
	return done
}

// NewUninstallInputStreamDir creates a channel for reading Uninstall as JSON newlines from a directory of files
func NewUninstallInputStreamDir(dir string, errors chan<- error, transforms ...TransformUninstallFunc) (chan Uninstall, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/agent/uninstall\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan Uninstall)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for uninstall")
		ch := make(chan Uninstall)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewUninstallInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan Uninstall)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewUninstallInputStreamFile creates an channel for reading Uninstall as JSON newlines from filename
func NewUninstallInputStreamFile(filename string, errors chan<- error, transforms ...TransformUninstallFunc) (chan Uninstall, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan Uninstall)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	var f io.ReadCloser = of
	if filepath.Ext(filename) == ".gz" {
		gz, err := gzip.NewReader(f)
		if err != nil {
			of.Close()
			errors <- err
			ch := make(chan Uninstall)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewUninstallInputStream(f, errors, transforms...)
}

// NewUninstallInputStream creates an channel for reading Uninstall as JSON newlines from stream
func NewUninstallInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformUninstallFunc) (chan Uninstall, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan Uninstall, 1000)
	go func() {
		defer func() { stream.Close(); close(ch); done <- true }()
		r := bufio.NewReader(stream)
		for {
			buf, err := r.ReadBytes('\n')
			if err != nil {
				if err == io.EOF {
					return
				}
				errors <- err
				return
			}
			var item Uninstall
			if err := json.Unmarshal(buf, &item); err != nil {
				errors <- err
				return
			}
			in := &item
			var skip bool
			for _, transform := range transforms {
				in, err = transform(in)
				if err != nil {
					errors <- err
					return
				}
				if in == nil {
					skip = true
					break
				}
			}
			if !skip {
				ch <- *in
			}
		}
	}()
	return ch, done
}

// NewUninstallOutputStreamDir will output json newlines from channel and save in dir
func NewUninstallOutputStreamDir(dir string, ch chan Uninstall, errors chan<- error, transforms ...TransformUninstallFunc) <-chan bool {
	fp := filepath.Join(dir, "/agent/uninstall\\.json(\\.gz)?$")
	os.MkdirAll(filepath.Dir(fp), 0777)
	of, err := os.Create(fp)
	if err != nil {
		errors <- err
		done := make(chan bool, 1)
		done <- true
		return done
	}
	gz, err := gzip.NewWriterLevel(of, gzip.BestCompression)
	if err != nil {
		errors <- err
		done := make(chan bool, 1)
		done <- true
		return done
	}
	return NewUninstallOutputStream(gz, ch, errors, transforms...)
}

// NewUninstallOutputStream will output json newlines from channel to the stream
func NewUninstallOutputStream(stream io.WriteCloser, ch chan Uninstall, errors chan<- error, transforms ...TransformUninstallFunc) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() {
			if gz, ok := stream.(*gzip.Writer); ok {
				gz.Flush()
				gz.Close()
			}
			stream.Close()
			done <- true
		}()
		for item := range ch {
			in := &item
			var skip bool
			var err error
			for _, transform := range transforms {
				in, err = transform(in)
				if err != nil {
					errors <- err
					return
				}
				if in == nil {
					skip = true
					break
				}
			}
			if !skip {
				buf, err := json.Marshal(in)
				if err != nil {
					errors <- err
					return
				}
				stream.Write(buf)
				stream.Write([]byte{'\n'})
			}
		}
	}()
	return done
}

// UninstallSendEvent is an event detail for sending data
type UninstallSendEvent struct {
	Uninstall *Uninstall
	headers   map[string]string
	time      time.Time
	key       string
}

var _ datamodel.ModelSendEvent = (*UninstallSendEvent)(nil)

// Key is the key to use for the message
func (e *UninstallSendEvent) Key() string {
	if e.key == "" {
		return e.Uninstall.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *UninstallSendEvent) Object() datamodel.Model {
	return e.Uninstall
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *UninstallSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *UninstallSendEvent) Timestamp() time.Time {
	return e.time
}

// UninstallSendEventOpts is a function handler for setting opts
type UninstallSendEventOpts func(o *UninstallSendEvent)

// WithUninstallSendEventKey sets the key value to a value different than the object ID
func WithUninstallSendEventKey(key string) UninstallSendEventOpts {
	return func(o *UninstallSendEvent) {
		o.key = key
	}
}

// WithUninstallSendEventTimestamp sets the timestamp value
func WithUninstallSendEventTimestamp(tv time.Time) UninstallSendEventOpts {
	return func(o *UninstallSendEvent) {
		o.time = tv
	}
}

// WithUninstallSendEventHeader sets the timestamp value
func WithUninstallSendEventHeader(key, value string) UninstallSendEventOpts {
	return func(o *UninstallSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewUninstallSendEvent returns a new UninstallSendEvent instance
func NewUninstallSendEvent(o *Uninstall, opts ...UninstallSendEventOpts) *UninstallSendEvent {
	res := &UninstallSendEvent{
		Uninstall: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewUninstallProducer will stream data from the channel
func NewUninstallProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		for {
			select {
			case <-ctx.Done():
				return
			case item := <-ch:
				if item == nil {
					empty <- true
					return
				}
				if object, ok := item.Object().(*Uninstall); ok {
					binary, codec, err := object.ToAvroBinary()
					if err != nil {
						errors <- fmt.Errorf("error encoding %s to avro binary data. %v", object.String(), err)
						return
					}
					headers := map[string]string{}
					object.SetEventHeaders(headers)
					for k, v := range item.Headers() {
						headers[k] = v
					}
					tv := item.Timestamp()
					if tv.IsZero() {
						tv = object.GetTimestamp() // if not provided in the message, use the objects value
					}
					if tv.IsZero() {
						tv = time.Now() // if its still zero, use the ingest time
					}
					// add generated message headers
					headers["message-id"] = pstrings.NewUUIDV4()
					headers["message-ts"] = fmt.Sprintf("%v", datetime.EpochNow())
					msg := eventing.Message{
						Encoding:  eventing.AvroEncoding,
						Key:       item.Key(),
						Value:     binary,
						Codec:     codec,
						Headers:   headers,
						Timestamp: tv,
						Topic:     object.GetTopicName().String(),
					}
					if err := producer.Send(ctx, msg); err != nil {
						errors <- fmt.Errorf("error sending %s. %v", object.String(), err)
					}
				} else {
					errors <- fmt.Errorf("invalid event received. expected an object of type agent.Uninstall but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewUninstallConsumer will stream data from the topic into the provided channel
func NewUninstallConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object Uninstall
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into agent.Uninstall: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avro data into agent.Uninstall: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for agent.Uninstall")
			}

			// ignore messages that have exceeded the TTL
			cfg := object.GetTopicConfig()
			if cfg != nil && cfg.TTL != 0 && msg.Timestamp.UTC().Add(cfg.TTL).Sub(time.Now().UTC()) < 0 {
				// if disable auto and we're skipping, we need to commit the message
				if !msg.IsAutoCommit() {
					msg.Commit()
				}
				return nil
			}
			msg.Codec = object.GetAvroCodec() // match the codec

			ch <- &UninstallReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object Uninstall
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &UninstallReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// UninstallReceiveEvent is an event detail for receiving data
type UninstallReceiveEvent struct {
	Uninstall *Uninstall
	message   eventing.Message
	eof       bool
}

var _ datamodel.ModelReceiveEvent = (*UninstallReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *UninstallReceiveEvent) Object() datamodel.Model {
	return e.Uninstall
}

// Message returns the underlying message data for the event
func (e *UninstallReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *UninstallReceiveEvent) EOF() bool {
	return e.eof
}

// UninstallProducer implements the datamodel.ModelEventProducer
type UninstallProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*UninstallProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *UninstallProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *UninstallProducer) Close() error {
	p.mu.Lock()
	closed := p.closed
	p.closed = true
	p.mu.Unlock()
	if !closed {
		close(p.ch)
		<-p.empty
		p.cancel()
		<-p.done
	}
	return nil
}

// NewProducerChannel returns a channel which can be used for producing Model events
func (o *Uninstall) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *Uninstall) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &UninstallProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewUninstallProducer(newctx, producer, ch, errors, empty),
	}
}

// NewUninstallProducerChannel returns a channel which can be used for producing Model events
func NewUninstallProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewUninstallProducerChannelSize(producer, 0, errors)
}

// NewUninstallProducerChannelSize returns a channel which can be used for producing Model events
func NewUninstallProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &UninstallProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewUninstallProducer(newctx, producer, ch, errors, empty),
	}
}

// UninstallConsumer implements the datamodel.ModelEventConsumer
type UninstallConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*UninstallConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *UninstallConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *UninstallConsumer) Close() error {
	c.mu.Lock()
	closed := c.closed
	c.closed = true
	c.mu.Unlock()
	var err error
	if !closed {
		c.callback.Close()
		err = c.consumer.Close()
	}
	return err
}

// NewConsumerChannel returns a consumer channel which can be used to consume Model events
func (o *Uninstall) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &UninstallConsumer{
		ch:       ch,
		callback: NewUninstallConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewUninstallConsumerChannel returns a consumer channel which can be used to consume Model events
func NewUninstallConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &UninstallConsumer{
		ch:       ch,
		callback: NewUninstallConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
