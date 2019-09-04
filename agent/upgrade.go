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
	// UpgradeTopic is the default topic name
	UpgradeTopic datamodel.TopicNameType = "agent_Upgrade_topic"

	// UpgradeStream is the default stream name
	UpgradeStream datamodel.TopicNameType = "agent_Upgrade_stream"

	// UpgradeTable is the default table name
	UpgradeTable datamodel.TopicNameType = "agent_upgrade"

	// UpgradeModelName is the model name
	UpgradeModelName datamodel.ModelNameType = "agent.Upgrade"
)

const (
	// UpgradeArchitectureColumn is the architecture column name
	UpgradeArchitectureColumn = "architecture"
	// UpgradeCustomerIDColumn is the customer_id column name
	UpgradeCustomerIDColumn = "customer_id"
	// UpgradeDataColumn is the data column name
	UpgradeDataColumn = "data"
	// UpgradeDistroColumn is the distro column name
	UpgradeDistroColumn = "distro"
	// UpgradeErrorColumn is the error column name
	UpgradeErrorColumn = "error"
	// UpgradeEventDateColumn is the event_date column name
	UpgradeEventDateColumn = "event_date"
	// UpgradeEventDateColumnEpochColumn is the epoch column property of the EventDate name
	UpgradeEventDateColumnEpochColumn = "event_date->epoch"
	// UpgradeEventDateColumnOffsetColumn is the offset column property of the EventDate name
	UpgradeEventDateColumnOffsetColumn = "event_date->offset"
	// UpgradeEventDateColumnRfc3339Column is the rfc3339 column property of the EventDate name
	UpgradeEventDateColumnRfc3339Column = "event_date->rfc3339"
	// UpgradeFreeSpaceColumn is the free_space column name
	UpgradeFreeSpaceColumn = "free_space"
	// UpgradeFromVersionColumn is the from_version column name
	UpgradeFromVersionColumn = "from_version"
	// UpgradeGoVersionColumn is the go_version column name
	UpgradeGoVersionColumn = "go_version"
	// UpgradeHostnameColumn is the hostname column name
	UpgradeHostnameColumn = "hostname"
	// UpgradeIDColumn is the id column name
	UpgradeIDColumn = "id"
	// UpgradeLastExportDateColumn is the last_export_date column name
	UpgradeLastExportDateColumn = "last_export_date"
	// UpgradeLastExportDateColumnEpochColumn is the epoch column property of the LastExportDate name
	UpgradeLastExportDateColumnEpochColumn = "last_export_date->epoch"
	// UpgradeLastExportDateColumnOffsetColumn is the offset column property of the LastExportDate name
	UpgradeLastExportDateColumnOffsetColumn = "last_export_date->offset"
	// UpgradeLastExportDateColumnRfc3339Column is the rfc3339 column property of the LastExportDate name
	UpgradeLastExportDateColumnRfc3339Column = "last_export_date->rfc3339"
	// UpgradeMemoryColumn is the memory column name
	UpgradeMemoryColumn = "memory"
	// UpgradeMessageColumn is the message column name
	UpgradeMessageColumn = "message"
	// UpgradeNumCPUColumn is the num_cpu column name
	UpgradeNumCPUColumn = "num_cpu"
	// UpgradeOSColumn is the os column name
	UpgradeOSColumn = "os"
	// UpgradeRefIDColumn is the ref_id column name
	UpgradeRefIDColumn = "ref_id"
	// UpgradeRefTypeColumn is the ref_type column name
	UpgradeRefTypeColumn = "ref_type"
	// UpgradeRequestIDColumn is the request_id column name
	UpgradeRequestIDColumn = "request_id"
	// UpgradeSuccessColumn is the success column name
	UpgradeSuccessColumn = "success"
	// UpgradeSystemIDColumn is the system_id column name
	UpgradeSystemIDColumn = "system_id"
	// UpgradeToVersionColumn is the to_version column name
	UpgradeToVersionColumn = "to_version"
	// UpgradeTypeColumn is the type column name
	UpgradeTypeColumn = "type"
	// UpgradeUpdatedAtColumn is the updated_ts column name
	UpgradeUpdatedAtColumn = "updated_ts"
	// UpgradeUptimeColumn is the uptime column name
	UpgradeUptimeColumn = "uptime"
	// UpgradeUUIDColumn is the uuid column name
	UpgradeUUIDColumn = "uuid"
	// UpgradeVersionColumn is the version column name
	UpgradeVersionColumn = "version"
)

// UpgradeEventDate represents the object structure for event_date
type UpgradeEventDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toUpgradeEventDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toUpgradeEventDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *UpgradeEventDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *UpgradeEventDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toUpgradeEventDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toUpgradeEventDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toUpgradeEventDateObject(o.Rfc3339, isavro, false, "string"),
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
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toUpgradeLastExportDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toUpgradeLastExportDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *UpgradeLastExportDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *UpgradeLastExportDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toUpgradeLastExportDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toUpgradeLastExportDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toUpgradeLastExportDateObject(o.Rfc3339, isavro, false, "string"),
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
	EventDate UpgradeEventDate `json:"event_date" bson:"event_date" yaml:"event_date" faker:"-"`
	// FreeSpace the amount of free space in bytes for the agent machine
	FreeSpace int64 `json:"free_space" bson:"free_space" yaml:"free_space" faker:"-"`
	// FromVersion the version we're upgrading
	FromVersion string `json:"from_version" bson:"from_version" yaml:"from_version" faker:"-"`
	// GoVersion the go version that the agent build was built with
	GoVersion string `json:"go_version" bson:"go_version" yaml:"go_version" faker:"-"`
	// Hostname the agent hostname
	Hostname string `json:"hostname" bson:"hostname" yaml:"hostname" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// LastExportDate the last export date
	LastExportDate UpgradeLastExportDate `json:"last_export_date" bson:"last_export_date" yaml:"last_export_date" faker:"-"`
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
	// SystemID system unique device ID
	SystemID string `json:"system_id" bson:"system_id" yaml:"system_id" faker:"-"`
	// ToVersion the version we're installing
	ToVersion string `json:"to_version" bson:"to_version" yaml:"to_version" faker:"-"`
	// Type the type of event
	Type UpgradeType `json:"type" bson:"type" yaml:"type" faker:"-"`
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
var _ datamodel.Model = (*Upgrade)(nil)

func toUpgradeObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toUpgradeObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *Upgrade:
		return v.ToMap(isavro)

	case UpgradeEventDate:
		return v.ToMap(isavro)

	case UpgradeLastExportDate:
		return v.ToMap(isavro)

	case UpgradeType:
		return v.String()

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
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

// GetStateKey returns a key for use in state store
func (o *Upgrade) GetStateKey() string {
	key := "uuid"
	return fmt.Sprintf("%s_%s", key, o.GetID())
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

// MarshalBinary returns the bytes for marshaling to binary
func (o *Upgrade) MarshalBinary() ([]byte, error) {
	return o.MarshalJSON()
}

func (o *Upgrade) UnmarshalBinary(data []byte) error {
	return o.UnmarshalJSON(data)
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

var cachedCodecUpgrade *goavro.Codec
var cachedCodecUpgradeLock sync.Mutex

// GetAvroCodec returns the avro codec for this model
func (o *Upgrade) GetAvroCodec() *goavro.Codec {
	cachedCodecUpgradeLock.Lock()
	if cachedCodecUpgrade == nil {
		c, err := GetUpgradeAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecUpgrade = c
	}
	cachedCodecUpgradeLock.Unlock()
	return cachedCodecUpgrade
}

// ToAvroBinary returns the data as Avro binary data
func (o *Upgrade) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *Upgrade) FromAvroBinary(value []byte) error {
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
func (o *Upgrade) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Upgrade objects are equal
func (o *Upgrade) IsEqual(other *Upgrade) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Upgrade) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	o.setDefaults(false)
	return map[string]interface{}{
		"architecture":     toUpgradeObject(o.Architecture, isavro, false, "string"),
		"customer_id":      toUpgradeObject(o.CustomerID, isavro, false, "string"),
		"data":             toUpgradeObject(o.Data, isavro, true, "string"),
		"distro":           toUpgradeObject(o.Distro, isavro, false, "string"),
		"error":            toUpgradeObject(o.Error, isavro, true, "string"),
		"event_date":       toUpgradeObject(o.EventDate, isavro, false, "event_date"),
		"free_space":       toUpgradeObject(o.FreeSpace, isavro, false, "long"),
		"from_version":     toUpgradeObject(o.FromVersion, isavro, false, "string"),
		"go_version":       toUpgradeObject(o.GoVersion, isavro, false, "string"),
		"hostname":         toUpgradeObject(o.Hostname, isavro, false, "string"),
		"id":               toUpgradeObject(o.ID, isavro, false, "string"),
		"last_export_date": toUpgradeObject(o.LastExportDate, isavro, false, "last_export_date"),
		"memory":           toUpgradeObject(o.Memory, isavro, false, "long"),
		"message":          toUpgradeObject(o.Message, isavro, false, "string"),
		"num_cpu":          toUpgradeObject(o.NumCPU, isavro, false, "long"),
		"os":               toUpgradeObject(o.OS, isavro, false, "string"),
		"ref_id":           toUpgradeObject(o.RefID, isavro, false, "string"),
		"ref_type":         toUpgradeObject(o.RefType, isavro, false, "string"),
		"request_id":       toUpgradeObject(o.RequestID, isavro, false, "string"),
		"success":          toUpgradeObject(o.Success, isavro, false, "boolean"),
		"system_id":        toUpgradeObject(o.SystemID, isavro, false, "string"),
		"to_version":       toUpgradeObject(o.ToVersion, isavro, false, "string"),
		"type":             toUpgradeObject(o.Type, isavro, false, "type"),
		"updated_ts":       toUpgradeObject(o.UpdatedAt, isavro, false, "long"),
		"uptime":           toUpgradeObject(o.Uptime, isavro, false, "long"),
		"uuid":             toUpgradeObject(o.UUID, isavro, false, "string"),
		"version":          toUpgradeObject(o.Version, isavro, false, "string"),
		"hashcode":         toUpgradeObject(o.Hashcode, isavro, false, "string"),
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
		} else if sv, ok := val.(UpgradeEventDate); ok {
			// struct
			o.EventDate = sv
		} else if sp, ok := val.(*UpgradeEventDate); ok {
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

// GetUpgradeAvroSchemaSpec creates the avro schema specification for Upgrade
func GetUpgradeAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "agent",
		"name":      "Upgrade",
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
				"name": "from_version",
				"type": "string",
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
				"name": "system_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "to_version",
				"type": "string",
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

// GetUpgradeAvroSchema creates the avro schema for Upgrade
func GetUpgradeAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetUpgradeAvroSchemaSpec())
}

// TransformUpgradeFunc is a function for transforming Upgrade during processing
type TransformUpgradeFunc func(input *Upgrade) (*Upgrade, error)

// NewUpgradePipe creates a pipe for processing Upgrade items
func NewUpgradePipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformUpgradeFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewUpgradeInputStream(input, errors)
	var stream chan Upgrade
	if len(transforms) > 0 {
		stream = make(chan Upgrade, 1000)
	} else {
		stream = inch
	}
	outdone := NewUpgradeOutputStream(output, stream, errors)
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

// NewUpgradeInputStreamDir creates a channel for reading Upgrade as JSON newlines from a directory of files
func NewUpgradeInputStreamDir(dir string, errors chan<- error, transforms ...TransformUpgradeFunc) (chan Upgrade, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/agent/upgrade\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan Upgrade)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for upgrade")
		ch := make(chan Upgrade)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewUpgradeInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan Upgrade)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewUpgradeInputStreamFile creates an channel for reading Upgrade as JSON newlines from filename
func NewUpgradeInputStreamFile(filename string, errors chan<- error, transforms ...TransformUpgradeFunc) (chan Upgrade, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan Upgrade)
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
			ch := make(chan Upgrade)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewUpgradeInputStream(f, errors, transforms...)
}

// NewUpgradeInputStream creates an channel for reading Upgrade as JSON newlines from stream
func NewUpgradeInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformUpgradeFunc) (chan Upgrade, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan Upgrade, 1000)
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
			var item Upgrade
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

// NewUpgradeOutputStreamDir will output json newlines from channel and save in dir
func NewUpgradeOutputStreamDir(dir string, ch chan Upgrade, errors chan<- error, transforms ...TransformUpgradeFunc) <-chan bool {
	fp := filepath.Join(dir, "/agent/upgrade\\.json(\\.gz)?$")
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
	return NewUpgradeOutputStream(gz, ch, errors, transforms...)
}

// NewUpgradeOutputStream will output json newlines from channel to the stream
func NewUpgradeOutputStream(stream io.WriteCloser, ch chan Upgrade, errors chan<- error, transforms ...TransformUpgradeFunc) <-chan bool {
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

// UpgradeSendEvent is an event detail for sending data
type UpgradeSendEvent struct {
	Upgrade *Upgrade
	headers map[string]string
	time    time.Time
	key     string
}

var _ datamodel.ModelSendEvent = (*UpgradeSendEvent)(nil)

// Key is the key to use for the message
func (e *UpgradeSendEvent) Key() string {
	if e.key == "" {
		return e.Upgrade.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *UpgradeSendEvent) Object() datamodel.Model {
	return e.Upgrade
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *UpgradeSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *UpgradeSendEvent) Timestamp() time.Time {
	return e.time
}

// UpgradeSendEventOpts is a function handler for setting opts
type UpgradeSendEventOpts func(o *UpgradeSendEvent)

// WithUpgradeSendEventKey sets the key value to a value different than the object ID
func WithUpgradeSendEventKey(key string) UpgradeSendEventOpts {
	return func(o *UpgradeSendEvent) {
		o.key = key
	}
}

// WithUpgradeSendEventTimestamp sets the timestamp value
func WithUpgradeSendEventTimestamp(tv time.Time) UpgradeSendEventOpts {
	return func(o *UpgradeSendEvent) {
		o.time = tv
	}
}

// WithUpgradeSendEventHeader sets the timestamp value
func WithUpgradeSendEventHeader(key, value string) UpgradeSendEventOpts {
	return func(o *UpgradeSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewUpgradeSendEvent returns a new UpgradeSendEvent instance
func NewUpgradeSendEvent(o *Upgrade, opts ...UpgradeSendEventOpts) *UpgradeSendEvent {
	res := &UpgradeSendEvent{
		Upgrade: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewUpgradeProducer will stream data from the channel
func NewUpgradeProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
	done := make(chan bool, 1)
	emptyTime := time.Unix(0, 0)
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
				if object, ok := item.Object().(*Upgrade); ok {
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
					if tv.IsZero() || tv.Equal(emptyTime) {
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
						Partition: -1, // select any partition based on partitioner strategy in kafka
						Topic:     object.GetTopicName().String(),
					}
					if err := producer.Send(ctx, msg); err != nil {
						errors <- fmt.Errorf("error sending %s. %v", object.String(), err)
					}
				} else {
					errors <- fmt.Errorf("invalid event received. expected an object of type agent.Upgrade but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewUpgradeConsumer will stream data from the topic into the provided channel
func NewUpgradeConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object Upgrade
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into agent.Upgrade: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avro data into agent.Upgrade: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for agent.Upgrade")
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

			ch <- &UpgradeReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object Upgrade
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &UpgradeReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// UpgradeReceiveEvent is an event detail for receiving data
type UpgradeReceiveEvent struct {
	Upgrade *Upgrade
	message eventing.Message
	eof     bool
}

var _ datamodel.ModelReceiveEvent = (*UpgradeReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *UpgradeReceiveEvent) Object() datamodel.Model {
	return e.Upgrade
}

// Message returns the underlying message data for the event
func (e *UpgradeReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *UpgradeReceiveEvent) EOF() bool {
	return e.eof
}

// UpgradeProducer implements the datamodel.ModelEventProducer
type UpgradeProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*UpgradeProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *UpgradeProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *UpgradeProducer) Close() error {
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
func (o *Upgrade) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *Upgrade) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &UpgradeProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewUpgradeProducer(newctx, producer, ch, errors, empty),
	}
}

// NewUpgradeProducerChannel returns a channel which can be used for producing Model events
func NewUpgradeProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewUpgradeProducerChannelSize(producer, 0, errors)
}

// NewUpgradeProducerChannelSize returns a channel which can be used for producing Model events
func NewUpgradeProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &UpgradeProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewUpgradeProducer(newctx, producer, ch, errors, empty),
	}
}

// UpgradeConsumer implements the datamodel.ModelEventConsumer
type UpgradeConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*UpgradeConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *UpgradeConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *UpgradeConsumer) Close() error {
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
func (o *Upgrade) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &UpgradeConsumer{
		ch:       ch,
		callback: NewUpgradeConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewUpgradeConsumerChannel returns a consumer channel which can be used to consume Model events
func NewUpgradeConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &UpgradeConsumer{
		ch:       ch,
		callback: NewUpgradeConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
