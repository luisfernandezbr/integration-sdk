// DO NOT EDIT -- generated code

// Package sourcecode - the system which contains source code
package sourcecode

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
	"github.com/pinpt/go-common/event"
	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
)

// RepoTopic is the default topic name
const RepoTopic datamodel.TopicNameType = "sourcecode_Repo_topic"

// RepoStream is the default stream name
const RepoStream datamodel.TopicNameType = "sourcecode_Repo_stream"

// RepoTable is the default table name
const RepoTable datamodel.TopicNameType = "sourcecode_Repo"

// RepoModelName is the model name
const RepoModelName datamodel.ModelNameType = "sourcecode.Repo"

// Repo the repo holds source code
type Repo struct {
	// built in types

	ID         string `json:"repo_id" bson:"repo_id" yaml:"repo_id" faker:"-"`
	MongoID    string `json:"_id" bson:"_id" yaml:"_id" faker:"-"` // generated and used internally, do not set
	RefID      string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	RefType    string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	Hashcode   string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
	// custom types

	// Name the name of the repo
	Name string `json:"name" bson:"name" yaml:"name" faker:"repo"`
	// URL the url to the repo home page
	URL string `json:"url" bson:"url" yaml:"url" faker:"url"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Repo)(nil)

func toRepoObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toRepoObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if o == nil {
		return toRepoObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toRepoObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toRepoObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toRepoObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toRepoObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toRepoObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toRepoObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toRepoObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toRepoObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toRepoObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toRepoObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case map[string]interface{}:
		return o
	case *map[string]interface{}:
		return v
	case *Repo:
		return v.ToMap()
	case Repo:
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
			arr = append(arr, toRepoObject(av, isavro, false, ""))
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of Repo
func (o *Repo) String() string {
	return fmt.Sprintf("sourcecode.Repo<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Repo) GetTopicName() datamodel.TopicNameType {
	return RepoTopic
}

// GetModelName returns the name of the model
func (o *Repo) GetModelName() datamodel.ModelNameType {
	return RepoModelName
}

func (o *Repo) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Repo) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Repo", o.CustomerID, o.RefType, o.GetRefID())
	}
	if o.MongoID == "" {
		o.MongoID = o.ID
	}
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Repo) GetTopicKey() string {
	var i interface{} = o.ID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetRefID returns the RefID for the object
func (o *Repo) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Repo) IsMaterialized() bool {
	return true
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Repo) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	idletime, err := time.ParseDuration("1s")
	if err != nil {
		panic(err)
	}
	return &datamodel.ModelMaterializeConfig{
		KeyName:   "repo_id",
		TableName: "sourcecode_repo",
		BatchSize: 5000,
		IdleTime:  idletime,
	}
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Repo) IsEvented() bool {
	return true
}

// GetTopicConfig returns the topic config object
func (o *Repo) GetTopicConfig() *datamodel.ModelTopicConfig {
	duration, err := time.ParseDuration("168h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 168h0m0s. " + err.Error())
	}
	return &datamodel.ModelTopicConfig{
		Key:               "repo_id",
		Timestamp:         "",
		NumPartitions:     4,
		ReplicationFactor: 1,
		Retention:         duration,
		MaxSize:           5242880,
	}
}

// Clone returns an exact copy of Repo
func (o *Repo) Clone() datamodel.Model {
	c := new(Repo)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Repo) Anon() datamodel.Model {
	c := new(Repo)
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
func (o *Repo) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Repo) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	// make sure that these have values if empty
	o.setDefaults()
	return nil
}

var cachedCodecRepo *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *Repo) GetAvroCodec() *goavro.Codec {
	if cachedCodecRepo == nil {
		c, err := GetRepoAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecRepo = c
	}
	return cachedCodecRepo
}

// ToAvroBinary returns the data as Avro binary data
func (o *Repo) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *Repo) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Repo objects are equal
func (o *Repo) IsEqual(other *Repo) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Repo) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	return map[string]interface{}{
		"repo_id":     o.GetID(),
		"ref_id":      o.GetRefID(),
		"ref_type":    o.RefType,
		"customer_id": o.CustomerID,
		"hashcode":    o.Hash(),
		"name":        toRepoObject(o.Name, isavro, false, "string"),
		"url":         toRepoObject(o.URL, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *Repo) FromMap(kv map[string]interface{}) {
	if val, ok := kv["repo_id"].(string); ok {
		o.ID = val
		o.MongoID = val
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
	if val, ok := kv["name"].(string); ok {
		o.Name = val
	} else {
		val := kv["name"]
		if val == nil {
			o.Name = ""
		} else {
			o.Name = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["url"].(string); ok {
		o.URL = val
	} else {
		val := kv["url"]
		if val == nil {
			o.URL = ""
		} else {
			o.URL = fmt.Sprintf("%v", val)
		}
	}
	// make sure that these have values if empty
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *Repo) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.GetID())
	args = append(args, o.GetRefID())
	args = append(args, o.RefType)
	args = append(args, o.CustomerID)
	args = append(args, o.Name)
	args = append(args, o.URL)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetRepoAvroSchemaSpec creates the avro schema specification for Repo
func GetRepoAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "sourcecode",
		"name":         "Repo",
		"connect.name": "sourcecode.Repo",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "repo_id",
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
				"name": "name",
				"type": "string",
			},
			map[string]interface{}{
				"name": "url",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetRepoAvroSchema creates the avro schema for Repo
func GetRepoAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetRepoAvroSchemaSpec())
}

// TransformRepoFunc is a function for transforming Repo during processing
type TransformRepoFunc func(input *Repo) (*Repo, error)

// NewRepoPipe creates a pipe for processing Repo items
func NewRepoPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformRepoFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewRepoInputStream(input, errors)
	var stream chan Repo
	if len(transforms) > 0 {
		stream = make(chan Repo, 1000)
	} else {
		stream = inch
	}
	outdone := NewRepoOutputStream(output, stream, errors)
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

// NewRepoInputStreamDir creates a channel for reading Repo as JSON newlines from a directory of files
func NewRepoInputStreamDir(dir string, errors chan<- error, transforms ...TransformRepoFunc) (chan Repo, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/sourcecode/repo\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan Repo)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for repo")
		ch := make(chan Repo)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewRepoInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan Repo)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewRepoInputStreamFile creates an channel for reading Repo as JSON newlines from filename
func NewRepoInputStreamFile(filename string, errors chan<- error, transforms ...TransformRepoFunc) (chan Repo, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan Repo)
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
			ch := make(chan Repo)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewRepoInputStream(f, errors, transforms...)
}

// NewRepoInputStream creates an channel for reading Repo as JSON newlines from stream
func NewRepoInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformRepoFunc) (chan Repo, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan Repo, 1000)
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
			var item Repo
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

// NewRepoOutputStreamDir will output json newlines from channel and save in dir
func NewRepoOutputStreamDir(dir string, ch chan Repo, errors chan<- error, transforms ...TransformRepoFunc) <-chan bool {
	fp := filepath.Join(dir, "/sourcecode/repo\\.json(\\.gz)?$")
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
	return NewRepoOutputStream(gz, ch, errors, transforms...)
}

// NewRepoOutputStream will output json newlines from channel to the stream
func NewRepoOutputStream(stream io.WriteCloser, ch chan Repo, errors chan<- error, transforms ...TransformRepoFunc) <-chan bool {
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

// RepoSendEvent is an event detail for sending data
type RepoSendEvent struct {
	Repo    *Repo
	headers map[string]string
	time    time.Time
	key     string
}

var _ datamodel.ModelSendEvent = (*RepoSendEvent)(nil)

// Key is the key to use for the message
func (e *RepoSendEvent) Key() string {
	if e.key == "" {
		return e.Repo.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *RepoSendEvent) Object() datamodel.Model {
	return e.Repo
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *RepoSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *RepoSendEvent) Timestamp() time.Time {
	return e.time
}

// RepoSendEventOpts is a function handler for setting opts
type RepoSendEventOpts func(o *RepoSendEvent)

// WithRepoSendEventKey sets the key value to a value different than the object ID
func WithRepoSendEventKey(key string) RepoSendEventOpts {
	return func(o *RepoSendEvent) {
		o.key = key
	}
}

// WithRepoSendEventTimestamp sets the timestamp value
func WithRepoSendEventTimestamp(tv time.Time) RepoSendEventOpts {
	return func(o *RepoSendEvent) {
		o.time = tv
	}
}

// WithRepoSendEventHeader sets the timestamp value
func WithRepoSendEventHeader(key, value string) RepoSendEventOpts {
	return func(o *RepoSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewRepoSendEvent returns a new RepoSendEvent instance
func NewRepoSendEvent(o *Repo, opts ...RepoSendEventOpts) *RepoSendEvent {
	res := &RepoSendEvent{
		Repo: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewRepoProducer will stream data from the channel
func NewRepoProducer(producer event.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		ctx := context.Background()
		for item := range ch {
			if object, ok := item.Object().(*Repo); ok {
				binary, codec, err := object.ToAvroBinary()
				if err != nil {
					errors <- fmt.Errorf("error encoding %s to avro binary data. %v", object.String(), err)
					return
				}
				headers := map[string]string{
					"customer_id": object.CustomerID,
					"model":       RepoModelName.String(),
				}
				for k, v := range item.Headers() {
					headers[k] = v
				}
				msg := event.Message{
					Key:     item.Key(),
					Value:   binary,
					Codec:   codec,
					Headers: headers,
				}
				if err := producer.Send(ctx, msg); err != nil {
					errors <- fmt.Errorf("error sending %s. %v", object.String(), err)
				}
			} else {
				errors <- fmt.Errorf("invalid event received. expected an object of type sourcecode.Repo but received on of type %v", reflect.TypeOf(item.Object()))
			}
		}
	}()
	return done
}

// NewRepoConsumer will stream data from the topic into the provided channel
func NewRepoConsumer(consumer event.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) {
	consumer.Consume(event.ConsumerCallback{
		OnDataReceived: func(msg event.Message) error {
			var object Repo
			if err := json.Unmarshal(msg.Value, &object); err != nil {
				return fmt.Errorf("error unmarshaling json data into sourcecode.Repo: %s", err)
			}
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &RepoReceiveEvent{&object, msg}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
	})
}

// RepoReceiveEvent is an event detail for receiving data
type RepoReceiveEvent struct {
	Repo    *Repo
	message event.Message
}

var _ datamodel.ModelReceiveEvent = (*RepoReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *RepoReceiveEvent) Object() datamodel.Model {
	return e.Repo
}

// Message returns the underlying message data for the event
func (e *RepoReceiveEvent) Message() event.Message {
	return e.message
}

// RepoProducer implements the datamodel.ModelEventProducer
type RepoProducer struct {
	ch   chan datamodel.ModelSendEvent
	done <-chan bool
}

var _ datamodel.ModelEventProducer = (*RepoProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *RepoProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *RepoProducer) Close() error {
	close(p.ch)
	<-p.done
	return nil
}

// NewProducerChannel returns a channel which can be used for producing Model events
func (o *Repo) NewProducerChannel(producer event.Producer, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent)
	return &RepoProducer{
		ch:   ch,
		done: NewRepoProducer(producer, ch, errors),
	}
}

// NewRepoProducerChannel returns a channel which can be used for producing Model events
func NewRepoProducerChannel(producer event.Producer, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent)
	return &RepoProducer{
		ch:   ch,
		done: NewRepoProducer(producer, ch, errors),
	}
}

// RepoConsumer implements the datamodel.ModelEventConsumer
type RepoConsumer struct {
	ch chan datamodel.ModelReceiveEvent
}

var _ datamodel.ModelEventConsumer = (*RepoConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *RepoConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *RepoConsumer) Close() error {
	close(c.ch)
	return nil
}

// NewConsumerChannel returns a consumer channel which can be used to consume Model events
func (o *Repo) NewConsumerChannel(consumer event.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	NewRepoConsumer(consumer, ch, errors)
	return &RepoConsumer{
		ch: ch,
	}
}

// NewRepoConsumerChannel returns a consumer channel which can be used to consume Model events
func NewRepoConsumerChannel(consumer event.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	NewRepoConsumer(consumer, ch, errors)
	return &RepoConsumer{
		ch: ch,
	}
}
