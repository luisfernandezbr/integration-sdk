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
	// EventTopic is the default topic name
	EventTopic datamodel.TopicNameType = "agent_Event_topic"

	// EventStream is the default stream name
	EventStream datamodel.TopicNameType = "agent_Event_stream"

	// EventTable is the default table name
	EventTable datamodel.TopicNameType = "agent_Event"

	// EventModelName is the model name
	EventModelName datamodel.ModelNameType = "agent.Event"
)

const (
	// EventArchitectureColumn is the architecture column name
	EventArchitectureColumn = "architecture"
	// EventCustomerIDColumn is the customer_id column name
	EventCustomerIDColumn = "customer_id"
	// EventDataColumn is the data column name
	EventDataColumn = "data"
	// EventDateColumn is the date column name
	EventDateColumn = "date"
	// EventDistroColumn is the distro column name
	EventDistroColumn = "distro"
	// EventErrorColumn is the error column name
	EventErrorColumn = "error"
	// EventFreeSpaceColumn is the free_space column name
	EventFreeSpaceColumn = "free_space"
	// EventGoVersionColumn is the go_version column name
	EventGoVersionColumn = "go_version"
	// EventHostnameColumn is the hostname column name
	EventHostnameColumn = "hostname"
	// EventIDColumn is the id column name
	EventIDColumn = "id"
	// EventMemoryColumn is the memory column name
	EventMemoryColumn = "memory"
	// EventMessageColumn is the message column name
	EventMessageColumn = "message"
	// EventNumCPUColumn is the num_cpu column name
	EventNumCPUColumn = "num_cpu"
	// EventOSColumn is the os column name
	EventOSColumn = "os"
	// EventRefIDColumn is the ref_id column name
	EventRefIDColumn = "ref_id"
	// EventRefTypeColumn is the ref_type column name
	EventRefTypeColumn = "ref_type"
	// EventTypeColumn is the type column name
	EventTypeColumn = "type"
	// EventUUIDColumn is the uuid column name
	EventUUIDColumn = "uuid"
	// EventVersionColumn is the version column name
	EventVersionColumn = "version"
)

// EventDate represents the object structure for date
type EventDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func (o *EventDate) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": o.Epoch,
		// Offset the timezone offset from GMT
		"offset": o.Offset,
		// Rfc3339 the date in RFC3339 format
		"rfc3339": o.Rfc3339,
	}
}

// Type is the enumeration type for type
type EventType int32

// String returns the string value for Type
func (v EventType) String() string {
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
	}
	return "unset"
}

const (
	// TypeEnroll is the enumeration value for enroll
	EventTypeEnroll EventType = 0
	// TypePing is the enumeration value for ping
	EventTypePing EventType = 1
	// TypeCrash is the enumeration value for crash
	EventTypeCrash EventType = 2
	// TypeIntegration is the enumeration value for integration
	EventTypeIntegration EventType = 3
	// TypeExport is the enumeration value for export
	EventTypeExport EventType = 4
	// TypeProject is the enumeration value for project
	EventTypeProject EventType = 5
)

// Event event data sent from the agent
type Event struct {
	// Architecture the architecture of the agent machine
	Architecture string `json:"architecture" bson:"architecture" yaml:"architecture" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Data extra data that is specific about this event
	Data *string `json:"data" bson:"data" yaml:"data" faker:"-"`
	// Date the date of the event
	Date EventDate `json:"date" bson:"date" yaml:"date" faker:"-"`
	// Distro the agent os distribution
	Distro string `json:"distro" bson:"distro" yaml:"distro" faker:"-"`
	// Error an error message related to this event
	Error *string `json:"error" bson:"error" yaml:"error" faker:"-"`
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
	// Type the type of event
	Type EventType `json:"type" bson:"type" yaml:"type" faker:"-"`
	// UUID the agent unique identifier
	UUID string `json:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
	// Version the agent version
	Version string `json:"version" bson:"version" yaml:"version" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Event)(nil)

func toEventObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toEventObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if o == nil {
		return toEventObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toEventObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toEventObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toEventObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toEventObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toEventObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toEventObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toEventObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toEventObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toEventObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toEventObjectNil(isavro, isoptional)
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
	case *Event:
		return v.ToMap()
	case Event:
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
			arr = append(arr, toEventObject(av, isavro, false, ""))
		}
		return arr

	case EventDate:
		vv := o.(EventDate)
		return vv.ToMap()
	case *EventDate:
		return (*o.(*EventDate)).ToMap()
	case []EventDate:
		arr := make([]interface{}, 0)
		for _, i := range o.([]EventDate) {
			arr = append(arr, i.ToMap())
		}
		return arr
	case *[]EventDate:
		arr := make([]interface{}, 0)
		vv := o.(*[]EventDate)
		for _, i := range *vv {
			arr = append(arr, i.ToMap())
		}
		return arr
	case EventType:
		if !isavro {
			return (o.(EventType)).String()
		}
		return map[string]string{
			"agent.type": (o.(EventType)).String(),
		}
	case *EventType:
		if !isavro {
			return (o.(*EventType)).String()
		}
		return map[string]string{
			"agent.type": (o.(*EventType)).String(),
		}
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of Event
func (o *Event) String() string {
	return fmt.Sprintf("agent.Event<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Event) GetTopicName() datamodel.TopicNameType {
	return EventTopic
}

// GetModelName returns the name of the model
func (o *Event) GetModelName() datamodel.ModelNameType {
	return EventModelName
}

func (o *Event) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Event) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Event", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Event) GetTopicKey() string {
	var i interface{} = o.UUID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Event) GetTimestamp() time.Time {
	var dt interface{} = o.Date
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
	case EventDate:
		return datetime.DateFromEpoch(v.Epoch)
	}
	panic("not sure how to handle the date time format for Event")
}

// GetRefID returns the RefID for the object
func (o *Event) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Event) IsMaterialized() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Event) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Event) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *Event) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = EventModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Event) GetTopicConfig() *datamodel.ModelTopicConfig {
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
		Timestamp:         "date",
		NumPartitions:     8,
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetStateKey returns a key for use in state store
func (o *Event) GetStateKey() string {
	key := "uuid"
	return fmt.Sprintf("%s_%s", key, o.GetID())
}

// GetCustomerID will return the customer_id
func (o *Event) GetCustomerID() string {
	return o.CustomerID
}

// Clone returns an exact copy of Event
func (o *Event) Clone() datamodel.Model {
	c := new(Event)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Event) Anon() datamodel.Model {
	c := new(Event)
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
func (o *Event) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Event) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

var cachedCodecEvent *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *Event) GetAvroCodec() *goavro.Codec {
	if cachedCodecEvent == nil {
		c, err := GetEventAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecEvent = c
	}
	return cachedCodecEvent
}

// ToAvroBinary returns the data as Avro binary data
func (o *Event) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *Event) FromAvroBinary(value []byte) error {
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
func (o *Event) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Event objects are equal
func (o *Event) IsEqual(other *Event) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Event) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	o.setDefaults()
	return map[string]interface{}{
		"architecture": toEventObject(o.Architecture, isavro, false, "string"),
		"customer_id":  toEventObject(o.CustomerID, isavro, false, "string"),
		"data":         toEventObject(o.Data, isavro, true, "string"),
		"date":         toEventObject(o.Date, isavro, false, "date"),
		"distro":       toEventObject(o.Distro, isavro, false, "string"),
		"error":        toEventObject(o.Error, isavro, true, "string"),
		"free_space":   toEventObject(o.FreeSpace, isavro, false, "long"),
		"go_version":   toEventObject(o.GoVersion, isavro, false, "string"),
		"hostname":     toEventObject(o.Hostname, isavro, false, "string"),
		"id":           toEventObject(o.ID, isavro, false, "string"),
		"memory":       toEventObject(o.Memory, isavro, false, "long"),
		"message":      toEventObject(o.Message, isavro, false, "string"),
		"num_cpu":      toEventObject(o.NumCPU, isavro, false, "long"),
		"os":           toEventObject(o.OS, isavro, false, "string"),
		"ref_id":       toEventObject(o.RefID, isavro, false, "string"),
		"ref_type":     toEventObject(o.RefType, isavro, false, "string"),
		"type":         toEventObject(o.Type, isavro, false, "type"),
		"uuid":         toEventObject(o.UUID, isavro, false, "string"),
		"version":      toEventObject(o.Version, isavro, false, "string"),
		"hashcode":     toEventObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *Event) FromMap(kv map[string]interface{}) {
	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
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
	if val, ok := kv["date"].(EventDate); ok {
		o.Date = val
	} else {
		val := kv["date"]
		if val == nil {
			o.Date = EventDate{}
		} else {
			o.Date = EventDate{}
			b, _ := json.Marshal(val)
			json.Unmarshal(b, &o.Date)

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
	if val, ok := kv["type"].(EventType); ok {
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
func (o *Event) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Architecture)
	args = append(args, o.CustomerID)
	args = append(args, o.Data)
	args = append(args, o.Date)
	args = append(args, o.Distro)
	args = append(args, o.Error)
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
	args = append(args, o.Type)
	args = append(args, o.UUID)
	args = append(args, o.Version)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetEventAvroSchemaSpec creates the avro schema specification for Event
func GetEventAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "agent",
		"name":      "Event",
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
				"name": "date",
				"type": map[string]interface{}{"type": "record", "name": "date", "fields": []interface{}{map[string]interface{}{"type": "long", "name": "epoch", "doc": "the date in epoch format"}, map[string]interface{}{"type": "long", "name": "offset", "doc": "the timezone offset from GMT"}, map[string]interface{}{"type": "string", "name": "rfc3339", "doc": "the date in RFC3339 format"}}, "doc": "the date of the event"},
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
				"name": "type",
				"type": []interface{}{
					map[string]interface{}{
						"type":    "enum",
						"name":    "type",
						"symbols": []interface{}{"enroll", "ping", "crash", "integration", "export", "project"},
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

// GetEventAvroSchema creates the avro schema for Event
func GetEventAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetEventAvroSchemaSpec())
}

// TransformEventFunc is a function for transforming Event during processing
type TransformEventFunc func(input *Event) (*Event, error)

// NewEventPipe creates a pipe for processing Event items
func NewEventPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformEventFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewEventInputStream(input, errors)
	var stream chan Event
	if len(transforms) > 0 {
		stream = make(chan Event, 1000)
	} else {
		stream = inch
	}
	outdone := NewEventOutputStream(output, stream, errors)
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

// NewEventInputStreamDir creates a channel for reading Event as JSON newlines from a directory of files
func NewEventInputStreamDir(dir string, errors chan<- error, transforms ...TransformEventFunc) (chan Event, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/agent/event\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan Event)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for event")
		ch := make(chan Event)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewEventInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan Event)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewEventInputStreamFile creates an channel for reading Event as JSON newlines from filename
func NewEventInputStreamFile(filename string, errors chan<- error, transforms ...TransformEventFunc) (chan Event, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan Event)
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
			ch := make(chan Event)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewEventInputStream(f, errors, transforms...)
}

// NewEventInputStream creates an channel for reading Event as JSON newlines from stream
func NewEventInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformEventFunc) (chan Event, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan Event, 1000)
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
			var item Event
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

// NewEventOutputStreamDir will output json newlines from channel and save in dir
func NewEventOutputStreamDir(dir string, ch chan Event, errors chan<- error, transforms ...TransformEventFunc) <-chan bool {
	fp := filepath.Join(dir, "/agent/event\\.json(\\.gz)?$")
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
	return NewEventOutputStream(gz, ch, errors, transforms...)
}

// NewEventOutputStream will output json newlines from channel to the stream
func NewEventOutputStream(stream io.WriteCloser, ch chan Event, errors chan<- error, transforms ...TransformEventFunc) <-chan bool {
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

// EventSendEvent is an event detail for sending data
type EventSendEvent struct {
	Event   *Event
	headers map[string]string
	time    time.Time
	key     string
}

var _ datamodel.ModelSendEvent = (*EventSendEvent)(nil)

// Key is the key to use for the message
func (e *EventSendEvent) Key() string {
	if e.key == "" {
		return e.Event.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *EventSendEvent) Object() datamodel.Model {
	return e.Event
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *EventSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *EventSendEvent) Timestamp() time.Time {
	return e.time
}

// EventSendEventOpts is a function handler for setting opts
type EventSendEventOpts func(o *EventSendEvent)

// WithEventSendEventKey sets the key value to a value different than the object ID
func WithEventSendEventKey(key string) EventSendEventOpts {
	return func(o *EventSendEvent) {
		o.key = key
	}
}

// WithEventSendEventTimestamp sets the timestamp value
func WithEventSendEventTimestamp(tv time.Time) EventSendEventOpts {
	return func(o *EventSendEvent) {
		o.time = tv
	}
}

// WithEventSendEventHeader sets the timestamp value
func WithEventSendEventHeader(key, value string) EventSendEventOpts {
	return func(o *EventSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewEventSendEvent returns a new EventSendEvent instance
func NewEventSendEvent(o *Event, opts ...EventSendEventOpts) *EventSendEvent {
	res := &EventSendEvent{
		Event: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewEventProducer will stream data from the channel
func NewEventProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
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
				if object, ok := item.Object().(*Event); ok {
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
					errors <- fmt.Errorf("invalid event received. expected an object of type agent.Event but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewEventConsumer will stream data from the topic into the provided channel
func NewEventConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object Event
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into agent.Event: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avro data into agent.Event: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for agent.Event")
			}
			// ignore messages that have exceeded the TTL
			cfg := object.GetTopicConfig()
			if cfg != nil && cfg.TTL != 0 && msg.Timestamp.Add(cfg.TTL).Sub(time.Now()) < 0 {
				return nil
			}
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &EventReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object Event
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &EventReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// EventReceiveEvent is an event detail for receiving data
type EventReceiveEvent struct {
	Event   *Event
	message eventing.Message
	eof     bool
}

var _ datamodel.ModelReceiveEvent = (*EventReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *EventReceiveEvent) Object() datamodel.Model {
	return e.Event
}

// Message returns the underlying message data for the event
func (e *EventReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *EventReceiveEvent) EOF() bool {
	return e.eof
}

// EventProducer implements the datamodel.ModelEventProducer
type EventProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*EventProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *EventProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *EventProducer) Close() error {
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
func (o *Event) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *Event) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &EventProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewEventProducer(newctx, producer, ch, errors, empty),
	}
}

// NewEventProducerChannel returns a channel which can be used for producing Model events
func NewEventProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewEventProducerChannelSize(producer, 0, errors)
}

// NewEventProducerChannelSize returns a channel which can be used for producing Model events
func NewEventProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &EventProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewEventProducer(newctx, producer, ch, errors, empty),
	}
}

// EventConsumer implements the datamodel.ModelEventConsumer
type EventConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*EventConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *EventConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *EventConsumer) Close() error {
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
func (o *Event) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &EventConsumer{
		ch:       ch,
		callback: NewEventConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewEventConsumerChannel returns a consumer channel which can be used to consume Model events
func NewEventConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &EventConsumer{
		ch:       ch,
		callback: NewEventConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
