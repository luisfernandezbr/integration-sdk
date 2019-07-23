// DO NOT EDIT -- generated code

// Package admin - admin specific data models
package admin

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
	// AgentTopic is the default topic name
	AgentTopic datamodel.TopicNameType = "admin_Agent_topic"

	// AgentStream is the default stream name
	AgentStream datamodel.TopicNameType = "admin_Agent_stream"

	// AgentTable is the default table name
	AgentTable datamodel.TopicNameType = "admin_Agent"

	// AgentModelName is the model name
	AgentModelName datamodel.ModelNameType = "admin.Agent"
)

const (
	// AgentCompletedDateColumn is the completed_date column name
	AgentCompletedDateColumn = "completed_date"
	// AgentCompletedDateColumnEpochColumn is the epoch column property of the CompletedDate name
	AgentCompletedDateColumnEpochColumn = "completed_date->epoch"
	// AgentCompletedDateColumnOffsetColumn is the offset column property of the CompletedDate name
	AgentCompletedDateColumnOffsetColumn = "completed_date->offset"
	// AgentCompletedDateColumnRfc3339Column is the rfc3339 column property of the CompletedDate name
	AgentCompletedDateColumnRfc3339Column = "completed_date->rfc3339"
	// AgentCreatedAtColumn is the created_ts column name
	AgentCreatedAtColumn = "created_ts"
	// AgentCustomerIDColumn is the customer_id column name
	AgentCustomerIDColumn = "customer_id"
	// AgentErrorColumn is the error column name
	AgentErrorColumn = "error"
	// AgentIDColumn is the id column name
	AgentIDColumn = "id"
	// AgentRefIDColumn is the ref_id column name
	AgentRefIDColumn = "ref_id"
	// AgentRefTypeColumn is the ref_type column name
	AgentRefTypeColumn = "ref_type"
	// AgentRerunHistoricalColumn is the rerun_historical column name
	AgentRerunHistoricalColumn = "rerun_historical"
	// AgentUpdatedAtColumn is the updated_ts column name
	AgentUpdatedAtColumn = "updated_ts"
	// AgentUUIDColumn is the uuid column name
	AgentUUIDColumn = "uuid"
)

// AgentCompletedDate represents the object structure for completed_date
type AgentCompletedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func (o *AgentCompletedDate) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": o.Epoch,
		// Offset the timezone offset from GMT
		"offset": o.Offset,
		// Rfc3339 the date in RFC3339 format
		"rfc3339": o.Rfc3339,
	}
}

// Agent Agent metadata for an enrolled agent
type Agent struct {
	// CompletedDate Last time the agent completed setup
	CompletedDate AgentCompletedDate `json:"completed_date" bson:"completed_date" yaml:"completed_date" faker:"-"`
	// CreatedAt the date the record was created in Epoch time
	CreatedAt int64 `json:"created_ts" bson:"created_ts" yaml:"created_ts" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Error Message if the agent is currently erroring out
	Error *string `json:"error" bson:"error" yaml:"error" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RerunHistorical Last re-run of historical data
	RerunHistorical int64 `json:"rerun_historical" bson:"rerun_historical" yaml:"rerun_historical" faker:"-"`
	// UpdatedAt the date the record was updated in Epoch time
	UpdatedAt int64 `json:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// UUID The uuid
	UUID string `json:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Agent)(nil)

func toAgentObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toAgentObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {

	if res := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); res != nil {
		return res
	}
	switch v := o.(type) {
	case *Agent:
		return v.ToMap()
	case Agent:
		return v.ToMap()

	case AgentCompletedDate:
		vv := o.(AgentCompletedDate)
		return vv.ToMap()
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of Agent
func (o *Agent) String() string {
	return fmt.Sprintf("admin.Agent<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Agent) GetTopicName() datamodel.TopicNameType {
	return AgentTopic
}

// GetModelName returns the name of the model
func (o *Agent) GetModelName() datamodel.ModelNameType {
	return AgentModelName
}

func (o *Agent) setDefaults() {
	if o.Error == nil {
		o.Error = &emptyString
	}

	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Agent) GetID() string {
	if o.ID == "" {
		// set the id from the spec provided in the model
		o.ID = hash.Values(o.CustomerID, o.RefID)
	}
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Agent) GetTopicKey() string {
	var i interface{} = o.ID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Agent) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *Agent) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Agent) IsMaterialized() bool {
	return true
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Agent) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	idletime, err := time.ParseDuration("1s")
	if err != nil {
		panic(err)
	}
	return &datamodel.ModelMaterializeConfig{
		KeyName:   "id",
		TableName: "admin_agent",
		BatchSize: 5000,
		IdleTime:  idletime,
	}
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Agent) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *Agent) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = AgentModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Agent) GetTopicConfig() *datamodel.ModelTopicConfig {
	retention, err := time.ParseDuration("168h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 168h0m0s. " + err.Error())
	}

	ttl, err := time.ParseDuration("0s")
	if err != nil {
		ttl = 0
	}
	return &datamodel.ModelTopicConfig{
		Key:               "id",
		Timestamp:         "",
		NumPartitions:     8,
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetStateKey returns a key for use in state store
func (o *Agent) GetStateKey() string {
	key := "id"
	return fmt.Sprintf("%s_%s", key, o.GetID())
}

// GetCustomerID will return the customer_id
func (o *Agent) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of Agent
func (o *Agent) Clone() datamodel.Model {
	c := new(Agent)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Agent) Anon() datamodel.Model {
	c := new(Agent)
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
func (o *Agent) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Agent) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

var cachedCodecAgent *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *Agent) GetAvroCodec() *goavro.Codec {
	if cachedCodecAgent == nil {
		c, err := GetAgentAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecAgent = c
	}
	return cachedCodecAgent
}

// ToAvroBinary returns the data as Avro binary data
func (o *Agent) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *Agent) FromAvroBinary(value []byte) error {
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
func (o *Agent) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Agent objects are equal
func (o *Agent) IsEqual(other *Agent) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Agent) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	o.setDefaults()
	return map[string]interface{}{
		"completed_date":   toAgentObject(o.CompletedDate, isavro, false, "completed_date"),
		"created_ts":       toAgentObject(o.CreatedAt, isavro, false, "long"),
		"customer_id":      toAgentObject(o.CustomerID, isavro, false, "string"),
		"error":            toAgentObject(o.Error, isavro, true, "string"),
		"id":               toAgentObject(o.ID, isavro, false, "string"),
		"ref_id":           toAgentObject(o.RefID, isavro, false, "string"),
		"ref_type":         toAgentObject(o.RefType, isavro, false, "string"),
		"rerun_historical": toAgentObject(o.RerunHistorical, isavro, false, "long"),
		"updated_ts":       toAgentObject(o.UpdatedAt, isavro, false, "long"),
		"uuid":             toAgentObject(o.UUID, isavro, false, "string"),
		"hashcode":         toAgentObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *Agent) FromMap(kv map[string]interface{}) {
	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["completed_date"].(AgentCompletedDate); ok {
		o.CompletedDate = val
	} else {
		val := kv["completed_date"]
		if val == nil {
			o.CompletedDate = AgentCompletedDate{}
		} else {
			o.CompletedDate = AgentCompletedDate{}
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
			json.Unmarshal(b, &o.CompletedDate)

		}
	}
	if val, ok := kv["created_ts"].(int64); ok {
		o.CreatedAt = val
	} else {
		val := kv["created_ts"]
		if val == nil {
			o.CreatedAt = number.ToInt64Any(nil)
		} else {
			if tv, ok := val.(time.Time); ok {
				val = datetime.TimeToEpoch(tv)
			}
			o.CreatedAt = number.ToInt64Any(val)
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
	if val, ok := kv["rerun_historical"].(int64); ok {
		o.RerunHistorical = val
	} else {
		val := kv["rerun_historical"]
		if val == nil {
			o.RerunHistorical = number.ToInt64Any(nil)
		} else {
			if tv, ok := val.(time.Time); ok {
				val = datetime.TimeToEpoch(tv)
			}
			o.RerunHistorical = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["updated_ts"].(int64); ok {
		o.UpdatedAt = val
	} else {
		val := kv["updated_ts"]
		if val == nil {
			o.UpdatedAt = number.ToInt64Any(nil)
		} else {
			if tv, ok := val.(time.Time); ok {
				val = datetime.TimeToEpoch(tv)
			}
			o.UpdatedAt = number.ToInt64Any(val)
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
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *Agent) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CompletedDate)
	args = append(args, o.CreatedAt)
	args = append(args, o.CustomerID)
	args = append(args, o.Error)
	args = append(args, o.ID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RerunHistorical)
	args = append(args, o.UpdatedAt)
	args = append(args, o.UUID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// CreateAgent creates a new Agent in the database
func CreateAgent(ctx context.Context, db datamodel.Storage, o *Agent) error {
	o.setDefaults()
	return db.Create(ctx, o)
}

// DeleteAgent deletes a Agent in the database
func DeleteAgent(ctx context.Context, db datamodel.Storage, o *Agent) error {
	o.setDefaults()
	return db.Delete(ctx, o)
}

// UpdateAgent updates a Agent in the database
func UpdateAgent(ctx context.Context, db datamodel.Storage, o *Agent) error {
	o.setDefaults()
	return db.Update(ctx, o)
}

// FindAgent returns a Agent from the database
func FindAgent(ctx context.Context, db datamodel.Storage, id string) (*Agent, error) {
	kv, err := db.FindOne(ctx, AgentModelName, id)
	if err != nil {
		return nil, err
	}
	if kv == nil {
		return nil, nil
	}
	return kv.(*Agent), nil
}

// FindAgents returns all Agent from the database matching keys
func FindAgents(ctx context.Context, db datamodel.Storage, kv map[string]interface{}) ([]*Agent, error) {
	res, err := db.Find(ctx, AgentModelName, kv)
	if err != nil {
		return nil, err
	}
	if res != nil {
		arr := make([]*Agent, 0)
		for _, m := range res {
			arr = append(arr, m.(*Agent))
		}
		return arr, nil
	}
	return nil, nil
}

// GetAgentAvroSchemaSpec creates the avro schema specification for Agent
func GetAgentAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "admin",
		"name":      "Agent",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "hashcode",
				"type": "string",
			},
			map[string]interface{}{
				"name": "completed_date",
				"type": map[string]interface{}{"type": "record", "name": "completed_date", "fields": []interface{}{map[string]interface{}{"doc": "the date in epoch format", "type": "long", "name": "epoch"}, map[string]interface{}{"type": "long", "name": "offset", "doc": "the timezone offset from GMT"}, map[string]interface{}{"name": "rfc3339", "doc": "the date in RFC3339 format", "type": "string"}}, "doc": "Last time the agent completed setup"},
			},
			map[string]interface{}{
				"name": "created_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "customer_id",
				"type": "string",
			},
			map[string]interface{}{
				"name":    "error",
				"type":    []interface{}{"null", "string"},
				"default": nil,
			},
			map[string]interface{}{
				"name": "id",
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
				"name": "rerun_historical",
				"type": "long",
			},
			map[string]interface{}{
				"name": "updated_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "uuid",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *Agent) GetEventAPIConfig() datamodel.EventAPIConfig {
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

// GetAgentAvroSchema creates the avro schema for Agent
func GetAgentAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetAgentAvroSchemaSpec())
}

// TransformAgentFunc is a function for transforming Agent during processing
type TransformAgentFunc func(input *Agent) (*Agent, error)

// NewAgentPipe creates a pipe for processing Agent items
func NewAgentPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformAgentFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewAgentInputStream(input, errors)
	var stream chan Agent
	if len(transforms) > 0 {
		stream = make(chan Agent, 1000)
	} else {
		stream = inch
	}
	outdone := NewAgentOutputStream(output, stream, errors)
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

// NewAgentInputStreamDir creates a channel for reading Agent as JSON newlines from a directory of files
func NewAgentInputStreamDir(dir string, errors chan<- error, transforms ...TransformAgentFunc) (chan Agent, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/admin/agent\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan Agent)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for agent")
		ch := make(chan Agent)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewAgentInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan Agent)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewAgentInputStreamFile creates an channel for reading Agent as JSON newlines from filename
func NewAgentInputStreamFile(filename string, errors chan<- error, transforms ...TransformAgentFunc) (chan Agent, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan Agent)
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
			ch := make(chan Agent)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewAgentInputStream(f, errors, transforms...)
}

// NewAgentInputStream creates an channel for reading Agent as JSON newlines from stream
func NewAgentInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformAgentFunc) (chan Agent, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan Agent, 1000)
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
			var item Agent
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

// NewAgentOutputStreamDir will output json newlines from channel and save in dir
func NewAgentOutputStreamDir(dir string, ch chan Agent, errors chan<- error, transforms ...TransformAgentFunc) <-chan bool {
	fp := filepath.Join(dir, "/admin/agent\\.json(\\.gz)?$")
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
	return NewAgentOutputStream(gz, ch, errors, transforms...)
}

// NewAgentOutputStream will output json newlines from channel to the stream
func NewAgentOutputStream(stream io.WriteCloser, ch chan Agent, errors chan<- error, transforms ...TransformAgentFunc) <-chan bool {
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

// AgentSendEvent is an event detail for sending data
type AgentSendEvent struct {
	Agent   *Agent
	headers map[string]string
	time    time.Time
	key     string
}

var _ datamodel.ModelSendEvent = (*AgentSendEvent)(nil)

// Key is the key to use for the message
func (e *AgentSendEvent) Key() string {
	if e.key == "" {
		return e.Agent.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *AgentSendEvent) Object() datamodel.Model {
	return e.Agent
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *AgentSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *AgentSendEvent) Timestamp() time.Time {
	return e.time
}

// AgentSendEventOpts is a function handler for setting opts
type AgentSendEventOpts func(o *AgentSendEvent)

// WithAgentSendEventKey sets the key value to a value different than the object ID
func WithAgentSendEventKey(key string) AgentSendEventOpts {
	return func(o *AgentSendEvent) {
		o.key = key
	}
}

// WithAgentSendEventTimestamp sets the timestamp value
func WithAgentSendEventTimestamp(tv time.Time) AgentSendEventOpts {
	return func(o *AgentSendEvent) {
		o.time = tv
	}
}

// WithAgentSendEventHeader sets the timestamp value
func WithAgentSendEventHeader(key, value string) AgentSendEventOpts {
	return func(o *AgentSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewAgentSendEvent returns a new AgentSendEvent instance
func NewAgentSendEvent(o *Agent, opts ...AgentSendEventOpts) *AgentSendEvent {
	res := &AgentSendEvent{
		Agent: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewAgentProducer will stream data from the channel
func NewAgentProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
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
				if object, ok := item.Object().(*Agent); ok {
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
					errors <- fmt.Errorf("invalid event received. expected an object of type admin.Agent but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewAgentConsumer will stream data from the topic into the provided channel
func NewAgentConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object Agent
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into admin.Agent: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avro data into admin.Agent: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for admin.Agent")
			}

			// ignore messages that have exceeded the TTL
			cfg := object.GetTopicConfig()
			if cfg != nil && cfg.TTL != 0 && msg.Timestamp.UTC().Add(cfg.TTL).Sub(time.Now().UTC()) < 0 {
				return nil
			}
			msg.Codec = object.GetAvroCodec() // match the codec

			ch <- &AgentReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object Agent
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &AgentReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// AgentReceiveEvent is an event detail for receiving data
type AgentReceiveEvent struct {
	Agent   *Agent
	message eventing.Message
	eof     bool
}

var _ datamodel.ModelReceiveEvent = (*AgentReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *AgentReceiveEvent) Object() datamodel.Model {
	return e.Agent
}

// Message returns the underlying message data for the event
func (e *AgentReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *AgentReceiveEvent) EOF() bool {
	return e.eof
}

// AgentProducer implements the datamodel.ModelEventProducer
type AgentProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*AgentProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *AgentProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *AgentProducer) Close() error {
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
func (o *Agent) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *Agent) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &AgentProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewAgentProducer(newctx, producer, ch, errors, empty),
	}
}

// NewAgentProducerChannel returns a channel which can be used for producing Model events
func NewAgentProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewAgentProducerChannelSize(producer, 0, errors)
}

// NewAgentProducerChannelSize returns a channel which can be used for producing Model events
func NewAgentProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &AgentProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewAgentProducer(newctx, producer, ch, errors, empty),
	}
}

// AgentConsumer implements the datamodel.ModelEventConsumer
type AgentConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*AgentConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *AgentConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *AgentConsumer) Close() error {
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
func (o *Agent) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &AgentConsumer{
		ch:       ch,
		callback: NewAgentConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewAgentConsumerChannel returns a consumer channel which can be used to consume Model events
func NewAgentConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &AgentConsumer{
		ch:       ch,
		callback: NewAgentConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
