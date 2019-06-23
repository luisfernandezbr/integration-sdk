// DO NOT EDIT -- generated code

// Package admin - admin specific data models
package admin

import (
	"bufio"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
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
	// AgentIDColumn is the id column name
	AgentIDColumn = "id"
	// AgentRefIDColumn is the ref_id column name
	AgentRefIDColumn = "ref_id"
	// AgentRefTypeColumn is the ref_type column name
	AgentRefTypeColumn = "ref_type"
	// AgentCustomerIDColumn is the customer_id column name
	AgentCustomerIDColumn = "customer_id"
	// AgentCreatedAtColumn is the created_ts column name
	AgentCreatedAtColumn = "created_ts"
	// AgentUpdatedAtColumn is the updated_ts column name
	AgentUpdatedAtColumn = "updated_ts"
	// AgentApikeyColumn is the apikey column name
	AgentApikeyColumn = "apikey"
	// AgentRerunHistoricalColumn is the rerun_historical column name
	AgentRerunHistoricalColumn = "rerun_historical"
	// AgentUUIDColumn is the uuid column name
	AgentUUIDColumn = "uuid"
	// AgentCompletedAtColumn is the completed_ts column name
	AgentCompletedAtColumn = "completed_ts"
)

// Agent Agent metadata
type Agent struct {
	// built in types

	ID         string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	RefID      string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	RefType    string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	Hashcode   string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`

	CreatedAt int64 `json:"created_ts" bson:"created_ts" yaml:"created_ts" faker:"-"`
	UpdatedAt int64 `json:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`

	// custom types

	// Apikey The apikey for this agent to use
	Apikey string `json:"apikey" bson:"apikey" yaml:"apikey" faker:"-"`
	// RerunHistorical Last re-run of historical data
	RerunHistorical int64 `json:"rerun_historical" bson:"rerun_historical" yaml:"rerun_historical" faker:"-"`
	// UUID The uuid
	UUID string `json:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
	// CompletedAt Last time the agent completed
	CompletedAt int64 `json:"completed_ts" bson:"completed_ts" yaml:"completed_ts" faker:"-"`
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
	if o == nil {
		return toAgentObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toAgentObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toAgentObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toAgentObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toAgentObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toAgentObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toAgentObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toAgentObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toAgentObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toAgentObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toAgentObjectNil(isavro, isoptional)
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
	case *Agent:
		return v.ToMap()
	case Agent:
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
			arr = append(arr, toAgentObject(av, isavro, false, ""))
		}
		return arr
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
	duration, err := time.ParseDuration("168h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 168h0m0s. " + err.Error())
	}
	return &datamodel.ModelTopicConfig{
		Key:               "id",
		Timestamp:         "",
		NumPartitions:     8,
		ReplicationFactor: 3,
		Retention:         duration,
		MaxSize:           5242880,
	}
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
	// make sure that these have values if empty
	o.setDefaults()
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

// Stringify returns the object in JSON format as a string
func (o *Agent) Stringify() string {
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
	return map[string]interface{}{
		"id":               o.GetID(),
		"ref_id":           o.GetRefID(),
		"ref_type":         o.RefType,
		"customer_id":      o.CustomerID,
		"hashcode":         o.Hash(),
		"created_ts":       o.CreatedAt,
		"updated_ts":       o.UpdatedAt,
		"apikey":           toAgentObject(o.Apikey, isavro, false, "string"),
		"rerun_historical": toAgentObject(o.RerunHistorical, isavro, false, "long"),
		"uuid":             toAgentObject(o.UUID, isavro, false, "string"),
		"completed_ts":     toAgentObject(o.CompletedAt, isavro, false, "long"),
	}
}

// FromMap attempts to load data into object from a map
func (o *Agent) FromMap(kv map[string]interface{}) {
	// make sure that these have values if empty
	o.setDefaults()
	if val, ok := kv["id"].(string); ok {
		o.ID = val
	} else if val, ok := kv["_id"].(string); ok {
		o.ID = val
	}
	if val, ok := kv["ref_id"].(string); ok {
		o.RefID = val
	}
	if val, ok := kv["ref_type"].(string); ok {
		o.RefType = val
	}
	if val, ok := kv["customer_id"].(string); ok {
		o.CustomerID = val
	}
	if val, ok := kv["created_ts"].(int64); ok {
		o.CreatedAt = val
	} else if val, ok := kv["created_ts"].(time.Time); ok {
		o.CreatedAt = datetime.TimeToEpoch(val)
	}
	if val, ok := kv["updated_ts"].(int64); ok {
		o.UpdatedAt = val
	} else if val, ok := kv["updated_ts"].(time.Time); ok {
		o.UpdatedAt = datetime.TimeToEpoch(val)
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
	if val, ok := kv["rerun_historical"].(int64); ok {
		o.RerunHistorical = val
	} else {
		val := kv["rerun_historical"]
		if val == nil {
			o.RerunHistorical = number.ToInt64Any(nil)
		} else {
			o.RerunHistorical = number.ToInt64Any(val)
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
	if val, ok := kv["completed_ts"].(int64); ok {
		o.CompletedAt = val
	} else {
		val := kv["completed_ts"]
		if val == nil {
			o.CompletedAt = number.ToInt64Any(nil)
		} else {
			o.CompletedAt = number.ToInt64Any(val)
		}
	}
}

// Hash will return a hashcode for the object
func (o *Agent) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.GetID())
	args = append(args, o.GetRefID())
	args = append(args, o.RefType)
	args = append(args, o.CustomerID)
	args = append(args, o.Apikey)
	args = append(args, o.RerunHistorical)
	args = append(args, o.UUID)
	args = append(args, o.CompletedAt)
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
		"type":         "record",
		"namespace":    "admin",
		"name":         "Agent",
		"connect.name": "admin.Agent",
		"fields": []map[string]interface{}{
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
				"name": "customer_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "hashcode",
				"type": "string",
			},
			map[string]interface{}{
				"name": "created_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "updated_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "apikey",
				"type": "string",
			},
			map[string]interface{}{
				"name": "rerun_historical",
				"type": "long",
			},
			map[string]interface{}{
				"name": "uuid",
				"type": "string",
			},
			map[string]interface{}{
				"name": "completed_ts",
				"type": "long",
			},
		},
	}
	return pjson.Stringify(spec, true)
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
func NewAgentProducer(producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		ctx := context.Background()
		for item := range ch {
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
					Key:       item.Key(),
					Value:     binary,
					Codec:     codec,
					Headers:   headers,
					Timestamp: tv,
				}
				if err := producer.Send(ctx, msg); err != nil {
					errors <- fmt.Errorf("error sending %s. %v", object.String(), err)
				}
			} else {
				errors <- fmt.Errorf("invalid event received. expected an object of type admin.Agent but received on of type %v", reflect.TypeOf(item.Object()))
			}
		}
	}()
	return done
}

// NewAgentConsumer will stream data from the topic into the provided channel
func NewAgentConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) {
	consumer.Consume(&eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object Agent
			if err := json.Unmarshal(msg.Value, &object); err != nil {
				return fmt.Errorf("error unmarshaling json data into admin.Agent: %s", err)
			}
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &AgentReceiveEvent{&object, msg}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
	})
}

// AgentReceiveEvent is an event detail for receiving data
type AgentReceiveEvent struct {
	Agent   *Agent
	message eventing.Message
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

// AgentProducer implements the datamodel.ModelEventProducer
type AgentProducer struct {
	ch   chan datamodel.ModelSendEvent
	done <-chan bool
}

var _ datamodel.ModelEventProducer = (*AgentProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *AgentProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *AgentProducer) Close() error {
	close(p.ch)
	<-p.done
	return nil
}

// NewProducerChannel returns a channel which can be used for producing Model events
func (o *Agent) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent)
	return &AgentProducer{
		ch:   ch,
		done: NewAgentProducer(producer, ch, errors),
	}
}

// NewAgentProducerChannel returns a channel which can be used for producing Model events
func NewAgentProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent)
	return &AgentProducer{
		ch:   ch,
		done: NewAgentProducer(producer, ch, errors),
	}
}

// AgentConsumer implements the datamodel.ModelEventConsumer
type AgentConsumer struct {
	ch chan datamodel.ModelReceiveEvent
}

var _ datamodel.ModelEventConsumer = (*AgentConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *AgentConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *AgentConsumer) Close() error {
	close(c.ch)
	return nil
}

// NewConsumerChannel returns a consumer channel which can be used to consume Model events
func (o *Agent) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	NewAgentConsumer(consumer, ch, errors)
	return &AgentConsumer{
		ch: ch,
	}
}

// NewAgentConsumerChannel returns a consumer channel which can be used to consume Model events
func NewAgentConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	NewAgentConsumer(consumer, ch, errors)
	return &AgentConsumer{
		ch: ch,
	}
}
