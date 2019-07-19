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
	// EnrollResponseTopic is the default topic name
	EnrollResponseTopic datamodel.TopicNameType = "agent_EnrollResponse_topic"

	// EnrollResponseStream is the default stream name
	EnrollResponseStream datamodel.TopicNameType = "agent_EnrollResponse_stream"

	// EnrollResponseTable is the default table name
	EnrollResponseTable datamodel.TopicNameType = "agent_EnrollResponse"

	// EnrollResponseModelName is the model name
	EnrollResponseModelName datamodel.ModelNameType = "agent.EnrollResponse"
)

const (
	// EnrollResponseApikeyColumn is the apikey column name
	EnrollResponseApikeyColumn = "apikey"
	// EnrollResponseArchitectureColumn is the architecture column name
	EnrollResponseArchitectureColumn = "architecture"
	// EnrollResponseCustomerIDColumn is the customer_id column name
	EnrollResponseCustomerIDColumn = "customer_id"
	// EnrollResponseDataColumn is the data column name
	EnrollResponseDataColumn = "data"
	// EnrollResponseDistroColumn is the distro column name
	EnrollResponseDistroColumn = "distro"
	// EnrollResponseErrorColumn is the error column name
	EnrollResponseErrorColumn = "error"
	// EnrollResponseEventDateColumn is the event_date column name
	EnrollResponseEventDateColumn = "event_date"
	// EnrollResponseEventDateColumnEpochColumn is the epoch column property of the EventDate name
	EnrollResponseEventDateColumnEpochColumn = "event_date->epoch"
	// EnrollResponseEventDateColumnOffsetColumn is the offset column property of the EventDate name
	EnrollResponseEventDateColumnOffsetColumn = "event_date->offset"
	// EnrollResponseEventDateColumnRfc3339Column is the rfc3339 column property of the EventDate name
	EnrollResponseEventDateColumnRfc3339Column = "event_date->rfc3339"
	// EnrollResponseFreeSpaceColumn is the free_space column name
	EnrollResponseFreeSpaceColumn = "free_space"
	// EnrollResponseGoVersionColumn is the go_version column name
	EnrollResponseGoVersionColumn = "go_version"
	// EnrollResponseHostnameColumn is the hostname column name
	EnrollResponseHostnameColumn = "hostname"
	// EnrollResponseIDColumn is the id column name
	EnrollResponseIDColumn = "id"
	// EnrollResponseMemoryColumn is the memory column name
	EnrollResponseMemoryColumn = "memory"
	// EnrollResponseMessageColumn is the message column name
	EnrollResponseMessageColumn = "message"
	// EnrollResponseNumCPUColumn is the num_cpu column name
	EnrollResponseNumCPUColumn = "num_cpu"
	// EnrollResponseOSColumn is the os column name
	EnrollResponseOSColumn = "os"
	// EnrollResponseRefIDColumn is the ref_id column name
	EnrollResponseRefIDColumn = "ref_id"
	// EnrollResponseRefTypeColumn is the ref_type column name
	EnrollResponseRefTypeColumn = "ref_type"
	// EnrollResponseRequestIDColumn is the request_id column name
	EnrollResponseRequestIDColumn = "request_id"
	// EnrollResponseSuccessColumn is the success column name
	EnrollResponseSuccessColumn = "success"
	// EnrollResponseTypeColumn is the type column name
	EnrollResponseTypeColumn = "type"
	// EnrollResponseUUIDColumn is the uuid column name
	EnrollResponseUUIDColumn = "uuid"
	// EnrollResponseVersionColumn is the version column name
	EnrollResponseVersionColumn = "version"
)

// EnrollResponseEventDate represents the object structure for event_date
type EnrollResponseEventDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func (o *EnrollResponseEventDate) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": o.Epoch,
		// Offset the timezone offset from GMT
		"offset": o.Offset,
		// Rfc3339 the date in RFC3339 format
		"rfc3339": o.Rfc3339,
	}
}

// EnrollResponseType is the enumeration type for type
type EnrollResponseType int32

// String returns the string value for Type
func (v EnrollResponseType) String() string {
	switch int32(v) {
	case 0:
		return "enroll"
	case 1:
		return "ping"
	case 2:
		return "crash"
	case 3:
		return "integration"
	case 4:
		return "export"
	case 5:
		return "project"
	case 6:
		return "user"
	case 7:
		return "repo"
	}
	return "unset"
}

const (
	// TypeEnroll is the enumeration value for enroll
	EnrollResponseTypeEnroll EnrollResponseType = 0
	// TypePing is the enumeration value for ping
	EnrollResponseTypePing EnrollResponseType = 1
	// TypeCrash is the enumeration value for crash
	EnrollResponseTypeCrash EnrollResponseType = 2
	// TypeIntegration is the enumeration value for integration
	EnrollResponseTypeIntegration EnrollResponseType = 3
	// TypeExport is the enumeration value for export
	EnrollResponseTypeExport EnrollResponseType = 4
	// TypeProject is the enumeration value for project
	EnrollResponseTypeProject EnrollResponseType = 5
	// TypeUser is the enumeration value for user
	EnrollResponseTypeUser EnrollResponseType = 6
	// TypeRepo is the enumeration value for repo
	EnrollResponseTypeRepo EnrollResponseType = 7
)

// EnrollResponse an agent response to an the enroll action
type EnrollResponse struct {
	// Apikey The API key that the agent should use when communicating with the cloud
	Apikey string `json:"apikey" bson:"apikey" yaml:"apikey" faker:"-"`
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
	EventDate EnrollResponseEventDate `json:"event_date" bson:"event_date" yaml:"event_date" faker:"-"`
	// FreeSpace the amount of free space in bytes for the agent machine
	FreeSpace int64 `json:"free_space" bson:"free_space" yaml:"free_space" faker:"-"`
	// GoVersion the go version that the agent build was built with
	GoVersion string `json:"go_version" bson:"go_version" yaml:"go_version" faker:"-"`
	// Hostname the agent hostname
	Hostname string `json:"hostname" bson:"hostname" yaml:"hostname" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
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
	Type EnrollResponseType `json:"type" bson:"type" yaml:"type" faker:"-"`
	// UUID the agent unique identifier
	UUID string `json:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
	// Version the agent version
	Version string `json:"version" bson:"version" yaml:"version" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*EnrollResponse)(nil)

func toEnrollResponseObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toEnrollResponseObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if o == nil {
		return toEnrollResponseObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toEnrollResponseObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toEnrollResponseObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toEnrollResponseObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toEnrollResponseObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toEnrollResponseObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toEnrollResponseObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toEnrollResponseObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toEnrollResponseObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toEnrollResponseObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toEnrollResponseObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case map[string]interface{}:
		return o
	case *map[string]interface{}:
		return v
	case map[string]string:
		return v
	case *map[string]string:
		return *v
	case *EnrollResponse:
		return v.ToMap()
	case EnrollResponse:
		return v.ToMap()
	case []string, []int64, []float64, []bool:
		return o
	case *[]string:
		return (*(o.(*[]string)))
	case *[]int64:
		return (*(o.(*[]int64)))
	case *[]float64:
		return (*(o.(*[]float64)))
	case *[]bool:
		return (*(o.(*[]bool)))
	case []interface{}:
		a := o.([]interface{})
		arr := make([]interface{}, 0)
		for _, av := range a {
			arr = append(arr, toEnrollResponseObject(av, isavro, false, ""))
		}
		return arr

	case EnrollResponseEventDate:
		vv := o.(EnrollResponseEventDate)
		return vv.ToMap()
	case *EnrollResponseEventDate:
		return map[string]interface{}{
			"agent.event_date": (*o.(*EnrollResponseEventDate)).ToMap(),
		}
	case []EnrollResponseEventDate:
		arr := make([]interface{}, 0)
		for _, i := range o.([]EnrollResponseEventDate) {
			arr = append(arr, i.ToMap())
		}
		return arr
	case *[]EnrollResponseEventDate:
		arr := make([]interface{}, 0)
		vv := o.(*[]EnrollResponseEventDate)
		for _, i := range *vv {
			arr = append(arr, i.ToMap())
		}
		return arr
	case EnrollResponseType:
		if !isavro {
			return (o.(EnrollResponseType)).String()
		}
		return map[string]string{
			"agent.type": (o.(EnrollResponseType)).String(),
		}
	case *EnrollResponseType:
		if !isavro {
			return (o.(*EnrollResponseType)).String()
		}
		return map[string]string{
			"agent.type": (o.(*EnrollResponseType)).String(),
		}
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of EnrollResponse
func (o *EnrollResponse) String() string {
	return fmt.Sprintf("agent.EnrollResponse<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *EnrollResponse) GetTopicName() datamodel.TopicNameType {
	return EnrollResponseTopic
}

// GetModelName returns the name of the model
func (o *EnrollResponse) GetModelName() datamodel.ModelNameType {
	return EnrollResponseModelName
}

func (o *EnrollResponse) setDefaults() {
	if o.Data == nil {
		o.Data = &emptyString
	}
	if o.Error == nil {
		o.Error = &emptyString
	}

	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *EnrollResponse) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("EnrollResponse", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *EnrollResponse) GetTopicKey() string {
	var i interface{} = o.UUID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *EnrollResponse) GetTimestamp() time.Time {
	var dt interface{} = o.EventDate
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
	case EnrollResponseEventDate:
		return datetime.DateFromEpoch(v.Epoch)
	}
	panic("not sure how to handle the date time format for EnrollResponse")
}

// GetRefID returns the RefID for the object
func (o *EnrollResponse) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *EnrollResponse) IsMaterialized() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *EnrollResponse) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *EnrollResponse) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *EnrollResponse) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = EnrollResponseModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *EnrollResponse) GetTopicConfig() *datamodel.ModelTopicConfig {
	retention, err := time.ParseDuration("168h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 168h0m0s. " + err.Error())
	}

	ttl, err := time.ParseDuration("0s")
	if err != nil {
		ttl = 0
	}
	return &datamodel.ModelTopicConfig{
		Key:               "uuid",
		Timestamp:         "event_date",
		NumPartitions:     8,
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetStateKey returns a key for use in state store
func (o *EnrollResponse) GetStateKey() string {
	key := "uuid"
	return fmt.Sprintf("%s_%s", key, o.GetID())
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
	return nil
}

var cachedCodecEnrollResponse *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *EnrollResponse) GetAvroCodec() *goavro.Codec {
	if cachedCodecEnrollResponse == nil {
		c, err := GetEnrollResponseAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecEnrollResponse = c
	}
	return cachedCodecEnrollResponse
}

// ToAvroBinary returns the data as Avro binary data
func (o *EnrollResponse) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *EnrollResponse) FromAvroBinary(value []byte) error {
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
func (o *EnrollResponse) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two EnrollResponse objects are equal
func (o *EnrollResponse) IsEqual(other *EnrollResponse) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *EnrollResponse) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	o.setDefaults()
	return map[string]interface{}{
		"apikey":       toEnrollResponseObject(o.Apikey, isavro, false, "string"),
		"architecture": toEnrollResponseObject(o.Architecture, isavro, false, "string"),
		"customer_id":  toEnrollResponseObject(o.CustomerID, isavro, false, "string"),
		"data":         toEnrollResponseObject(o.Data, isavro, true, "string"),
		"distro":       toEnrollResponseObject(o.Distro, isavro, false, "string"),
		"error":        toEnrollResponseObject(o.Error, isavro, true, "string"),
		"event_date":   toEnrollResponseObject(o.EventDate, isavro, false, "event_date"),
		"free_space":   toEnrollResponseObject(o.FreeSpace, isavro, false, "long"),
		"go_version":   toEnrollResponseObject(o.GoVersion, isavro, false, "string"),
		"hostname":     toEnrollResponseObject(o.Hostname, isavro, false, "string"),
		"id":           toEnrollResponseObject(o.ID, isavro, false, "string"),
		"memory":       toEnrollResponseObject(o.Memory, isavro, false, "long"),
		"message":      toEnrollResponseObject(o.Message, isavro, false, "string"),
		"num_cpu":      toEnrollResponseObject(o.NumCPU, isavro, false, "long"),
		"os":           toEnrollResponseObject(o.OS, isavro, false, "string"),
		"ref_id":       toEnrollResponseObject(o.RefID, isavro, false, "string"),
		"ref_type":     toEnrollResponseObject(o.RefType, isavro, false, "string"),
		"request_id":   toEnrollResponseObject(o.RequestID, isavro, false, "string"),
		"success":      toEnrollResponseObject(o.Success, isavro, false, "boolean"),
		"type":         toEnrollResponseObject(o.Type, isavro, false, "type"),
		"uuid":         toEnrollResponseObject(o.UUID, isavro, false, "string"),
		"version":      toEnrollResponseObject(o.Version, isavro, false, "string"),
		"hashcode":     toEnrollResponseObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *EnrollResponse) FromMap(kv map[string]interface{}) {
	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["apikey"].(string); ok {
		o.Apikey = val
	} else {
		val := kv["apikey"]
		if val == nil {
			o.Apikey = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Apikey = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["architecture"].(string); ok {
		o.Architecture = val
	} else {
		val := kv["architecture"]
		if val == nil {
			o.Architecture = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Architecture = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["customer_id"].(string); ok {
		o.CustomerID = val
	} else {
		val := kv["customer_id"]
		if val == nil {
			o.CustomerID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.CustomerID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["data"].(*string); ok {
		o.Data = val
	} else if val, ok := kv["data"].(string); ok {
		o.Data = &val
	} else {
		val := kv["data"]
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
	if val, ok := kv["distro"].(string); ok {
		o.Distro = val
	} else {
		val := kv["distro"]
		if val == nil {
			o.Distro = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Distro = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["error"].(*string); ok {
		o.Error = val
	} else if val, ok := kv["error"].(string); ok {
		o.Error = &val
	} else {
		val := kv["error"]
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
	if val, ok := kv["event_date"].(EnrollResponseEventDate); ok {
		o.EventDate = val
	} else {
		val := kv["event_date"]
		if val == nil {
			o.EventDate = EnrollResponseEventDate{}
		} else {
			o.EventDate = EnrollResponseEventDate{}
			if m, ok := val.(map[interface{}]interface{}); ok {
				si := make(map[string]interface{})
				for k, v := range m {
					if key, ok := k.(string); ok {
						si[key] = v
					}
				}
				val = si
			}
			b, _ := json.Marshal(val)
			json.Unmarshal(b, &o.EventDate)

		}
	}
	if val, ok := kv["free_space"].(int64); ok {
		o.FreeSpace = val
	} else {
		val := kv["free_space"]
		if val == nil {
			o.FreeSpace = number.ToInt64Any(nil)
		} else {
			if tv, ok := val.(time.Time); ok {
				val = datetime.TimeToEpoch(tv)
			}
			o.FreeSpace = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["go_version"].(string); ok {
		o.GoVersion = val
	} else {
		val := kv["go_version"]
		if val == nil {
			o.GoVersion = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.GoVersion = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["hostname"].(string); ok {
		o.Hostname = val
	} else {
		val := kv["hostname"]
		if val == nil {
			o.Hostname = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Hostname = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["id"].(string); ok {
		o.ID = val
	} else {
		val := kv["id"]
		if val == nil {
			o.ID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.ID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["memory"].(int64); ok {
		o.Memory = val
	} else {
		val := kv["memory"]
		if val == nil {
			o.Memory = number.ToInt64Any(nil)
		} else {
			if tv, ok := val.(time.Time); ok {
				val = datetime.TimeToEpoch(tv)
			}
			o.Memory = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["message"].(string); ok {
		o.Message = val
	} else {
		val := kv["message"]
		if val == nil {
			o.Message = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Message = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["num_cpu"].(int64); ok {
		o.NumCPU = val
	} else {
		val := kv["num_cpu"]
		if val == nil {
			o.NumCPU = number.ToInt64Any(nil)
		} else {
			if tv, ok := val.(time.Time); ok {
				val = datetime.TimeToEpoch(tv)
			}
			o.NumCPU = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["os"].(string); ok {
		o.OS = val
	} else {
		val := kv["os"]
		if val == nil {
			o.OS = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.OS = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["ref_id"].(string); ok {
		o.RefID = val
	} else {
		val := kv["ref_id"]
		if val == nil {
			o.RefID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.RefID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["ref_type"].(string); ok {
		o.RefType = val
	} else {
		val := kv["ref_type"]
		if val == nil {
			o.RefType = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.RefType = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["request_id"].(string); ok {
		o.RequestID = val
	} else {
		val := kv["request_id"]
		if val == nil {
			o.RequestID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.RequestID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["success"].(bool); ok {
		o.Success = val
	} else {
		val := kv["success"]
		if val == nil {
			o.Success = number.ToBoolAny(nil)
		} else {
			o.Success = number.ToBoolAny(val)
		}
	}
	if val, ok := kv["type"].(EnrollResponseType); ok {
		o.Type = val
	} else {
		if em, ok := kv["type"].(map[string]interface{}); ok {
			ev := em["agent.type"].(string)
			switch ev {
			case "enroll":
				o.Type = 0
			case "ping":
				o.Type = 1
			case "crash":
				o.Type = 2
			case "integration":
				o.Type = 3
			case "export":
				o.Type = 4
			case "project":
				o.Type = 5
			case "user":
				o.Type = 6
			case "repo":
				o.Type = 7
			}
		}
		if em, ok := kv["type"].(string); ok {
			switch em {
			case "enroll":
				o.Type = 0
			case "ping":
				o.Type = 1
			case "crash":
				o.Type = 2
			case "integration":
				o.Type = 3
			case "export":
				o.Type = 4
			case "project":
				o.Type = 5
			case "user":
				o.Type = 6
			case "repo":
				o.Type = 7
			}
		}
	}
	if val, ok := kv["uuid"].(string); ok {
		o.UUID = val
	} else {
		val := kv["uuid"]
		if val == nil {
			o.UUID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.UUID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["version"].(string); ok {
		o.Version = val
	} else {
		val := kv["version"]
		if val == nil {
			o.Version = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Version = fmt.Sprintf("%v", val)
		}
	}
	o.setDefaults()
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
	args = append(args, o.Memory)
	args = append(args, o.Message)
	args = append(args, o.NumCPU)
	args = append(args, o.OS)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RequestID)
	args = append(args, o.Success)
	args = append(args, o.Type)
	args = append(args, o.UUID)
	args = append(args, o.Version)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetEnrollResponseAvroSchemaSpec creates the avro schema specification for EnrollResponse
func GetEnrollResponseAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "agent",
		"name":      "EnrollResponse",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "hashcode",
				"type": "string",
			},
			map[string]interface{}{
				"name": "apikey",
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
				"type": map[string]interface{}{"fields": []interface{}{map[string]interface{}{"type": "long", "name": "epoch", "doc": "the date in epoch format"}, map[string]interface{}{"doc": "the timezone offset from GMT", "type": "long", "name": "offset"}, map[string]interface{}{"type": "string", "name": "rfc3339", "doc": "the date in RFC3339 format"}}, "doc": "the date of the event", "type": "record", "name": "event_date"},
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
				"type": []interface{}{
					map[string]interface{}{
						"type":    "enum",
						"name":    "type",
						"symbols": []interface{}{"enroll", "ping", "crash", "integration", "export", "project", "user", "repo"},
					},
				},
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

// GetEnrollResponseAvroSchema creates the avro schema for EnrollResponse
func GetEnrollResponseAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetEnrollResponseAvroSchemaSpec())
}

// TransformEnrollResponseFunc is a function for transforming EnrollResponse during processing
type TransformEnrollResponseFunc func(input *EnrollResponse) (*EnrollResponse, error)

// NewEnrollResponsePipe creates a pipe for processing EnrollResponse items
func NewEnrollResponsePipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformEnrollResponseFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewEnrollResponseInputStream(input, errors)
	var stream chan EnrollResponse
	if len(transforms) > 0 {
		stream = make(chan EnrollResponse, 1000)
	} else {
		stream = inch
	}
	outdone := NewEnrollResponseOutputStream(output, stream, errors)
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

// NewEnrollResponseInputStreamDir creates a channel for reading EnrollResponse as JSON newlines from a directory of files
func NewEnrollResponseInputStreamDir(dir string, errors chan<- error, transforms ...TransformEnrollResponseFunc) (chan EnrollResponse, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/agent/enroll_response\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan EnrollResponse)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for enroll_response")
		ch := make(chan EnrollResponse)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewEnrollResponseInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan EnrollResponse)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewEnrollResponseInputStreamFile creates an channel for reading EnrollResponse as JSON newlines from filename
func NewEnrollResponseInputStreamFile(filename string, errors chan<- error, transforms ...TransformEnrollResponseFunc) (chan EnrollResponse, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan EnrollResponse)
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
			ch := make(chan EnrollResponse)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewEnrollResponseInputStream(f, errors, transforms...)
}

// NewEnrollResponseInputStream creates an channel for reading EnrollResponse as JSON newlines from stream
func NewEnrollResponseInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformEnrollResponseFunc) (chan EnrollResponse, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan EnrollResponse, 1000)
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
			var item EnrollResponse
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

// NewEnrollResponseOutputStreamDir will output json newlines from channel and save in dir
func NewEnrollResponseOutputStreamDir(dir string, ch chan EnrollResponse, errors chan<- error, transforms ...TransformEnrollResponseFunc) <-chan bool {
	fp := filepath.Join(dir, "/agent/enroll_response\\.json(\\.gz)?$")
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
	return NewEnrollResponseOutputStream(gz, ch, errors, transforms...)
}

// NewEnrollResponseOutputStream will output json newlines from channel to the stream
func NewEnrollResponseOutputStream(stream io.WriteCloser, ch chan EnrollResponse, errors chan<- error, transforms ...TransformEnrollResponseFunc) <-chan bool {
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

// EnrollResponseSendEvent is an event detail for sending data
type EnrollResponseSendEvent struct {
	EnrollResponse *EnrollResponse
	headers        map[string]string
	time           time.Time
	key            string
}

var _ datamodel.ModelSendEvent = (*EnrollResponseSendEvent)(nil)

// Key is the key to use for the message
func (e *EnrollResponseSendEvent) Key() string {
	if e.key == "" {
		return e.EnrollResponse.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *EnrollResponseSendEvent) Object() datamodel.Model {
	return e.EnrollResponse
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *EnrollResponseSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *EnrollResponseSendEvent) Timestamp() time.Time {
	return e.time
}

// EnrollResponseSendEventOpts is a function handler for setting opts
type EnrollResponseSendEventOpts func(o *EnrollResponseSendEvent)

// WithEnrollResponseSendEventKey sets the key value to a value different than the object ID
func WithEnrollResponseSendEventKey(key string) EnrollResponseSendEventOpts {
	return func(o *EnrollResponseSendEvent) {
		o.key = key
	}
}

// WithEnrollResponseSendEventTimestamp sets the timestamp value
func WithEnrollResponseSendEventTimestamp(tv time.Time) EnrollResponseSendEventOpts {
	return func(o *EnrollResponseSendEvent) {
		o.time = tv
	}
}

// WithEnrollResponseSendEventHeader sets the timestamp value
func WithEnrollResponseSendEventHeader(key, value string) EnrollResponseSendEventOpts {
	return func(o *EnrollResponseSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewEnrollResponseSendEvent returns a new EnrollResponseSendEvent instance
func NewEnrollResponseSendEvent(o *EnrollResponse, opts ...EnrollResponseSendEventOpts) *EnrollResponseSendEvent {
	res := &EnrollResponseSendEvent{
		EnrollResponse: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewEnrollResponseProducer will stream data from the channel
func NewEnrollResponseProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
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
				if object, ok := item.Object().(*EnrollResponse); ok {
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
					errors <- fmt.Errorf("invalid event received. expected an object of type agent.EnrollResponse but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewEnrollResponseConsumer will stream data from the topic into the provided channel
func NewEnrollResponseConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object EnrollResponse
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into agent.EnrollResponse: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avro data into agent.EnrollResponse: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for agent.EnrollResponse")
			}

			// ignore messages that have exceeded the TTL
			cfg := object.GetTopicConfig()
			if cfg != nil && cfg.TTL != 0 && msg.Timestamp.UTC().Add(cfg.TTL).Sub(time.Now().UTC()) < 0 {
				return nil
			}
			msg.Codec = object.GetAvroCodec() // match the codec

			ch <- &EnrollResponseReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object EnrollResponse
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &EnrollResponseReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// EnrollResponseReceiveEvent is an event detail for receiving data
type EnrollResponseReceiveEvent struct {
	EnrollResponse *EnrollResponse
	message        eventing.Message
	eof            bool
}

var _ datamodel.ModelReceiveEvent = (*EnrollResponseReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *EnrollResponseReceiveEvent) Object() datamodel.Model {
	return e.EnrollResponse
}

// Message returns the underlying message data for the event
func (e *EnrollResponseReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *EnrollResponseReceiveEvent) EOF() bool {
	return e.eof
}

// EnrollResponseProducer implements the datamodel.ModelEventProducer
type EnrollResponseProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*EnrollResponseProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *EnrollResponseProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *EnrollResponseProducer) Close() error {
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
func (o *EnrollResponse) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *EnrollResponse) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &EnrollResponseProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewEnrollResponseProducer(newctx, producer, ch, errors, empty),
	}
}

// NewEnrollResponseProducerChannel returns a channel which can be used for producing Model events
func NewEnrollResponseProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewEnrollResponseProducerChannelSize(producer, 0, errors)
}

// NewEnrollResponseProducerChannelSize returns a channel which can be used for producing Model events
func NewEnrollResponseProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &EnrollResponseProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewEnrollResponseProducer(newctx, producer, ch, errors, empty),
	}
}

// EnrollResponseConsumer implements the datamodel.ModelEventConsumer
type EnrollResponseConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*EnrollResponseConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *EnrollResponseConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *EnrollResponseConsumer) Close() error {
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
func (o *EnrollResponse) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &EnrollResponseConsumer{
		ch:       ch,
		callback: NewEnrollResponseConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewEnrollResponseConsumerChannel returns a consumer channel which can be used to consume Model events
func NewEnrollResponseConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &EnrollResponseConsumer{
		ch:       ch,
		callback: NewEnrollResponseConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
