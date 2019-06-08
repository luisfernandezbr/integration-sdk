// DO NOT EDIT -- generated code

// Package codequality - the system which contains code quality
package codequality

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

// ProjectTopic is the default topic name
const ProjectTopic datamodel.TopicNameType = "codequality_Project_topic"

// ProjectStream is the default stream name
const ProjectStream datamodel.TopicNameType = "codequality_Project_stream"

// ProjectTable is the default table name
const ProjectTable datamodel.TopicNameType = "codequality_Project"

// ProjectModelName is the model name
const ProjectModelName datamodel.ModelNameType = "codequality.Project"

// Project project information
type Project struct {
	// built in types

	ID         string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	RefID      string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	RefType    string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	Hashcode   string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
	// custom types

	// Identifier the common identifier of the project
	Identifier string `json:"identifier" bson:"identifier" yaml:"identifier" faker:"-"`
	// Name the name of the project
	Name string `json:"name" bson:"name" yaml:"name" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Project)(nil)

func toProjectObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toProjectObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if o == nil {
		return toProjectObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toProjectObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toProjectObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toProjectObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toProjectObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toProjectObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toProjectObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toProjectObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toProjectObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toProjectObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toProjectObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case map[string]interface{}:
		return o
	case *map[string]interface{}:
		return v
	case *Project:
		return v.ToMap()
	case Project:
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
			arr = append(arr, toProjectObject(av, isavro, false, ""))
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of Project
func (o *Project) String() string {
	return fmt.Sprintf("codequality.Project<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Project) GetTopicName() datamodel.TopicNameType {
	return ProjectTopic
}

// GetModelName returns the name of the model
func (o *Project) GetModelName() datamodel.ModelNameType {
	return ProjectModelName
}

func (o *Project) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Project) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Project", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Project) GetTopicKey() string {
	var i interface{} = o.ID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetRefID returns the RefID for the object
func (o *Project) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Project) IsMaterialized() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Project) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Project) IsEvented() bool {
	return true
}

// GetTopicConfig returns the topic config object
func (o *Project) GetTopicConfig() *datamodel.ModelTopicConfig {
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

// Clone returns an exact copy of Project
func (o *Project) Clone() datamodel.Model {
	c := new(Project)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Project) Anon() datamodel.Model {
	c := new(Project)
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
func (o *Project) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Project) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	// make sure that these have values if empty
	o.setDefaults()
	return nil
}

var cachedCodecProject *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *Project) GetAvroCodec() *goavro.Codec {
	if cachedCodecProject == nil {
		c, err := GetProjectAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecProject = c
	}
	return cachedCodecProject
}

// ToAvroBinary returns the data as Avro binary data
func (o *Project) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *Project) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Project objects are equal
func (o *Project) IsEqual(other *Project) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Project) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	return map[string]interface{}{
		"id":          o.GetID(),
		"ref_id":      o.GetRefID(),
		"ref_type":    o.RefType,
		"customer_id": o.CustomerID,
		"hashcode":    o.Hash(),
		"identifier":  toProjectObject(o.Identifier, isavro, false, "string"),
		"name":        toProjectObject(o.Name, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *Project) FromMap(kv map[string]interface{}) {
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
	if val, ok := kv["identifier"].(string); ok {
		o.Identifier = val
	} else {
		val := kv["identifier"]
		if val == nil {
			o.Identifier = ""
		} else {
			o.Identifier = fmt.Sprintf("%v", val)
		}
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
	// make sure that these have values if empty
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *Project) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.GetID())
	args = append(args, o.GetRefID())
	args = append(args, o.RefType)
	args = append(args, o.CustomerID)
	args = append(args, o.Identifier)
	args = append(args, o.Name)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetProjectAvroSchemaSpec creates the avro schema specification for Project
func GetProjectAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "codequality",
		"name":         "Project",
		"connect.name": "codequality.Project",
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
				"name": "identifier",
				"type": "string",
			},
			map[string]interface{}{
				"name": "name",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetProjectAvroSchema creates the avro schema for Project
func GetProjectAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetProjectAvroSchemaSpec())
}

// TransformProjectFunc is a function for transforming Project during processing
type TransformProjectFunc func(input *Project) (*Project, error)

// NewProjectPipe creates a pipe for processing Project items
func NewProjectPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformProjectFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewProjectInputStream(input, errors)
	var stream chan Project
	if len(transforms) > 0 {
		stream = make(chan Project, 1000)
	} else {
		stream = inch
	}
	outdone := NewProjectOutputStream(output, stream, errors)
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

// NewProjectInputStreamDir creates a channel for reading Project as JSON newlines from a directory of files
func NewProjectInputStreamDir(dir string, errors chan<- error, transforms ...TransformProjectFunc) (chan Project, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/codequality/project\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan Project)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for project")
		ch := make(chan Project)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewProjectInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan Project)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewProjectInputStreamFile creates an channel for reading Project as JSON newlines from filename
func NewProjectInputStreamFile(filename string, errors chan<- error, transforms ...TransformProjectFunc) (chan Project, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan Project)
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
			ch := make(chan Project)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewProjectInputStream(f, errors, transforms...)
}

// NewProjectInputStream creates an channel for reading Project as JSON newlines from stream
func NewProjectInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformProjectFunc) (chan Project, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan Project, 1000)
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
			var item Project
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

// NewProjectOutputStreamDir will output json newlines from channel and save in dir
func NewProjectOutputStreamDir(dir string, ch chan Project, errors chan<- error, transforms ...TransformProjectFunc) <-chan bool {
	fp := filepath.Join(dir, "/codequality/project\\.json(\\.gz)?$")
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
	return NewProjectOutputStream(gz, ch, errors, transforms...)
}

// NewProjectOutputStream will output json newlines from channel to the stream
func NewProjectOutputStream(stream io.WriteCloser, ch chan Project, errors chan<- error, transforms ...TransformProjectFunc) <-chan bool {
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

// ProjectSendEvent is an event detail for sending data
type ProjectSendEvent struct {
	Project *Project
	headers map[string]string
	time    time.Time
	key     string
}

var _ datamodel.ModelSendEvent = (*ProjectSendEvent)(nil)

// Key is the key to use for the message
func (e *ProjectSendEvent) Key() string {
	if e.key == "" {
		return e.Project.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *ProjectSendEvent) Object() datamodel.Model {
	return e.Project
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *ProjectSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *ProjectSendEvent) Timestamp() time.Time {
	return e.time
}

// ProjectSendEventOpts is a function handler for setting opts
type ProjectSendEventOpts func(o *ProjectSendEvent)

// WithProjectSendEventKey sets the key value to a value different than the object ID
func WithProjectSendEventKey(key string) ProjectSendEventOpts {
	return func(o *ProjectSendEvent) {
		o.key = key
	}
}

// WithProjectSendEventTimestamp sets the timestamp value
func WithProjectSendEventTimestamp(tv time.Time) ProjectSendEventOpts {
	return func(o *ProjectSendEvent) {
		o.time = tv
	}
}

// WithProjectSendEventHeader sets the timestamp value
func WithProjectSendEventHeader(key, value string) ProjectSendEventOpts {
	return func(o *ProjectSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewProjectSendEvent returns a new ProjectSendEvent instance
func NewProjectSendEvent(o *Project, opts ...ProjectSendEventOpts) *ProjectSendEvent {
	res := &ProjectSendEvent{
		Project: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewProjectProducer will stream data from the channel
func NewProjectProducer(producer event.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		ctx := context.Background()
		for item := range ch {
			if object, ok := item.Object().(*Project); ok {
				binary, codec, err := object.ToAvroBinary()
				if err != nil {
					errors <- fmt.Errorf("error encoding %s to avro binary data. %v", object.String(), err)
					return
				}
				headers := map[string]string{
					"customer_id": object.CustomerID,
					"model":       ProjectModelName.String(),
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
				errors <- fmt.Errorf("invalid event received. expected an object of type codequality.Project but received on of type %v", reflect.TypeOf(item.Object()))
			}
		}
	}()
	return done
}

// NewProjectConsumer will stream data from the topic into the provided channel
func NewProjectConsumer(consumer event.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) {
	consumer.Consume(event.ConsumerCallback{
		OnDataReceived: func(msg event.Message) error {
			var object Project
			if err := json.Unmarshal(msg.Value, &object); err != nil {
				return fmt.Errorf("error unmarshaling json data into codequality.Project: %s", err)
			}
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &ProjectReceiveEvent{&object, msg}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
	})
}

// ProjectReceiveEvent is an event detail for receiving data
type ProjectReceiveEvent struct {
	Project *Project
	message event.Message
}

var _ datamodel.ModelReceiveEvent = (*ProjectReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *ProjectReceiveEvent) Object() datamodel.Model {
	return e.Project
}

// Message returns the underlying message data for the event
func (e *ProjectReceiveEvent) Message() event.Message {
	return e.message
}

// ProjectProducer implements the datamodel.ModelEventProducer
type ProjectProducer struct {
	ch   chan datamodel.ModelSendEvent
	done <-chan bool
}

var _ datamodel.ModelEventProducer = (*ProjectProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *ProjectProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *ProjectProducer) Close() error {
	close(p.ch)
	<-p.done
	return nil
}

// NewProducerChannel returns a channel which can be used for producing Model events
func (o *Project) NewProducerChannel(producer event.Producer, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent)
	return &ProjectProducer{
		ch:   ch,
		done: NewProjectProducer(producer, ch, errors),
	}
}

// NewProjectProducerChannel returns a channel which can be used for producing Model events
func NewProjectProducerChannel(producer event.Producer, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent)
	return &ProjectProducer{
		ch:   ch,
		done: NewProjectProducer(producer, ch, errors),
	}
}

// ProjectConsumer implements the datamodel.ModelEventConsumer
type ProjectConsumer struct {
	ch chan datamodel.ModelReceiveEvent
}

var _ datamodel.ModelEventConsumer = (*ProjectConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *ProjectConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *ProjectConsumer) Close() error {
	close(c.ch)
	return nil
}

// NewConsumerChannel returns a consumer channel which can be used to consume Model events
func (o *Project) NewConsumerChannel(consumer event.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	NewProjectConsumer(consumer, ch, errors)
	return &ProjectConsumer{
		ch: ch,
	}
}

// NewProjectConsumerChannel returns a consumer channel which can be used to consume Model events
func NewProjectConsumerChannel(consumer event.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	NewProjectConsumer(consumer, ch, errors)
	return &ProjectConsumer{
		ch: ch,
	}
}
