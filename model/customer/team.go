// DO NOT EDIT -- generated code

// Package customer - customer specific data
package customer

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
	pstrings "github.com/pinpt/go-common/strings"
)

// TeamTopic is the default topic name
const TeamTopic datamodel.TopicNameType = "customer_Team_topic"

// TeamStream is the default stream name
const TeamStream datamodel.TopicNameType = "customer_Team_stream"

// TeamTable is the default table name
const TeamTable datamodel.TopicNameType = "customer_Team"

// TeamModelName is the model name
const TeamModelName datamodel.ModelNameType = "customer.Team"

// Team a team is a grouping of one or more users
type Team struct {
	// built in types

	ID         string `json:"team_id" bson:"team_id" yaml:"team_id" faker:"-"`
	MongoID    string `json:"_id" bson:"_id" yaml:"_id" faker:"-"` // generated and used internally, do not set
	RefID      string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	RefType    string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	Hashcode   string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
	// custom types

	// Name the name of the team
	Name string `json:"name" bson:"name" yaml:"name" faker:"team"`
	// ParentID the parent id of the team
	ParentID *string `json:"parent_id" bson:"parent_id" yaml:"parent_id" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Team)(nil)

func toTeamObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toTeamObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if o == nil {
		return toTeamObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toTeamObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toTeamObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toTeamObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toTeamObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toTeamObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toTeamObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toTeamObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toTeamObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toTeamObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toTeamObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case map[string]interface{}:
		return o
	case *map[string]interface{}:
		return v
	case *Team:
		return v.ToMap()
	case Team:
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
			arr = append(arr, toTeamObject(av, isavro, false, ""))
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of Team
func (o *Team) String() string {
	return fmt.Sprintf("customer.Team<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Team) GetTopicName() datamodel.TopicNameType {
	return TeamTopic
}

// GetModelName returns the name of the model
func (o *Team) GetModelName() datamodel.ModelNameType {
	return TeamModelName
}

func (o *Team) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Team) GetID() string {
	if o.ID == "" {
		// set the id from the spec provided in the model
		o.ID = hash.Values(o.CustomerID, o.Name)
	}
	if o.MongoID == "" {
		o.MongoID = o.ID
	}
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Team) GetTopicKey() string {
	var i interface{} = o.ID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetRefID returns the RefID for the object
func (o *Team) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Team) IsMaterialized() bool {
	return true
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Team) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	idletime, err := time.ParseDuration("1s")
	if err != nil {
		panic(err)
	}
	return &datamodel.ModelMaterializeConfig{
		KeyName:   "team_id",
		TableName: "customer_team",
		BatchSize: 5000,
		IdleTime:  idletime,
	}
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Team) IsEvented() bool {
	return true
}

// GetTopicConfig returns the topic config object
func (o *Team) GetTopicConfig() *datamodel.ModelTopicConfig {
	duration, err := time.ParseDuration("168h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 168h0m0s. " + err.Error())
	}
	return &datamodel.ModelTopicConfig{
		Key:               "team_id",
		Timestamp:         "",
		NumPartitions:     4,
		ReplicationFactor: 1,
		Retention:         duration,
		MaxSize:           5242880,
	}
}

// Clone returns an exact copy of Team
func (o *Team) Clone() datamodel.Model {
	c := new(Team)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Team) Anon() datamodel.Model {
	c := new(Team)
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
func (o *Team) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Team) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	// make sure that these have values if empty
	o.setDefaults()
	return nil
}

var cachedCodecTeam *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *Team) GetAvroCodec() *goavro.Codec {
	if cachedCodecTeam == nil {
		c, err := GetTeamAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecTeam = c
	}
	return cachedCodecTeam
}

// ToAvroBinary returns the data as Avro binary data
func (o *Team) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *Team) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Team objects are equal
func (o *Team) IsEqual(other *Team) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Team) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	return map[string]interface{}{
		"team_id":     o.GetID(),
		"ref_id":      o.GetRefID(),
		"ref_type":    o.RefType,
		"customer_id": o.CustomerID,
		"hashcode":    o.Hash(),
		"name":        toTeamObject(o.Name, isavro, false, "string"),
		"parent_id":   toTeamObject(o.ParentID, isavro, true, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *Team) FromMap(kv map[string]interface{}) {
	if val, ok := kv["team_id"].(string); ok {
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
	if val, ok := kv["parent_id"].(*string); ok {
		o.ParentID = val
	} else if val, ok := kv["parent_id"].(string); ok {
		o.ParentID = &val
	} else {
		val := kv["parent_id"]
		if val == nil {
			o.ParentID = pstrings.Pointer("")
		} else {
			o.ParentID = pstrings.Pointer(fmt.Sprintf("%v", val))
		}
	}
	// make sure that these have values if empty
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *Team) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.GetID())
	args = append(args, o.GetRefID())
	args = append(args, o.RefType)
	args = append(args, o.CustomerID)
	args = append(args, o.Name)
	args = append(args, o.ParentID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// CreateTeam creates a new Team in the database
func CreateTeam(ctx context.Context, db datamodel.Storage, o *Team) error {
	o.setDefaults()
	return db.Create(ctx, o)
}

// DeleteTeam deletes a Team in the database
func DeleteTeam(ctx context.Context, db datamodel.Storage, o *Team) error {
	o.setDefaults()
	return db.Delete(ctx, o)
}

// UpdateTeam updates a Team in the database
func UpdateTeam(ctx context.Context, db datamodel.Storage, o *Team) error {
	o.setDefaults()
	return db.Update(ctx, o)
}

// FindTeam returns a Team from the database
func FindTeam(ctx context.Context, db datamodel.Storage, id string) (*Team, error) {
	kv, err := db.FindOne(ctx, TeamModelName, id)
	if err != nil {
		return nil, err
	}
	if kv == nil {
		return nil, nil
	}
	return kv.(*Team), nil
}

// FindTeams returns all Team from the database matching keys
func FindTeams(ctx context.Context, db datamodel.Storage, kv map[string]interface{}) ([]*Team, error) {
	res, err := db.Find(ctx, TeamModelName, kv)
	if err != nil {
		return nil, err
	}
	if res != nil {
		arr := make([]*Team, 0)
		for _, m := range res {
			arr = append(arr, m.(*Team))
		}
		return arr, nil
	}
	return nil, nil
}

// GetTeamAvroSchemaSpec creates the avro schema specification for Team
func GetTeamAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "customer",
		"name":         "Team",
		"connect.name": "customer.Team",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "team_id",
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
				"name":    "parent_id",
				"type":    []interface{}{"null", "string"},
				"default": nil,
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetTeamAvroSchema creates the avro schema for Team
func GetTeamAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetTeamAvroSchemaSpec())
}

// TransformTeamFunc is a function for transforming Team during processing
type TransformTeamFunc func(input *Team) (*Team, error)

// NewTeamPipe creates a pipe for processing Team items
func NewTeamPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformTeamFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewTeamInputStream(input, errors)
	var stream chan Team
	if len(transforms) > 0 {
		stream = make(chan Team, 1000)
	} else {
		stream = inch
	}
	outdone := NewTeamOutputStream(output, stream, errors)
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

// NewTeamInputStreamDir creates a channel for reading Team as JSON newlines from a directory of files
func NewTeamInputStreamDir(dir string, errors chan<- error, transforms ...TransformTeamFunc) (chan Team, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/customer/team\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan Team)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for team")
		ch := make(chan Team)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewTeamInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan Team)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewTeamInputStreamFile creates an channel for reading Team as JSON newlines from filename
func NewTeamInputStreamFile(filename string, errors chan<- error, transforms ...TransformTeamFunc) (chan Team, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan Team)
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
			ch := make(chan Team)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewTeamInputStream(f, errors, transforms...)
}

// NewTeamInputStream creates an channel for reading Team as JSON newlines from stream
func NewTeamInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformTeamFunc) (chan Team, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan Team, 1000)
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
			var item Team
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

// NewTeamOutputStreamDir will output json newlines from channel and save in dir
func NewTeamOutputStreamDir(dir string, ch chan Team, errors chan<- error, transforms ...TransformTeamFunc) <-chan bool {
	fp := filepath.Join(dir, "/customer/team\\.json(\\.gz)?$")
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
	return NewTeamOutputStream(gz, ch, errors, transforms...)
}

// NewTeamOutputStream will output json newlines from channel to the stream
func NewTeamOutputStream(stream io.WriteCloser, ch chan Team, errors chan<- error, transforms ...TransformTeamFunc) <-chan bool {
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

// TeamSendEvent is an event detail for sending data
type TeamSendEvent struct {
	Team    *Team
	headers map[string]string
	time    time.Time
	key     string
}

var _ datamodel.ModelSendEvent = (*TeamSendEvent)(nil)

// Key is the key to use for the message
func (e *TeamSendEvent) Key() string {
	if e.key == "" {
		return e.Team.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *TeamSendEvent) Object() datamodel.Model {
	return e.Team
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *TeamSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *TeamSendEvent) Timestamp() time.Time {
	return e.time
}

// TeamSendEventOpts is a function handler for setting opts
type TeamSendEventOpts func(o *TeamSendEvent)

// WithTeamSendEventKey sets the key value to a value different than the object ID
func WithTeamSendEventKey(key string) TeamSendEventOpts {
	return func(o *TeamSendEvent) {
		o.key = key
	}
}

// WithTeamSendEventTimestamp sets the timestamp value
func WithTeamSendEventTimestamp(tv time.Time) TeamSendEventOpts {
	return func(o *TeamSendEvent) {
		o.time = tv
	}
}

// WithTeamSendEventHeader sets the timestamp value
func WithTeamSendEventHeader(key, value string) TeamSendEventOpts {
	return func(o *TeamSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewTeamSendEvent returns a new TeamSendEvent instance
func NewTeamSendEvent(o *Team, opts ...TeamSendEventOpts) *TeamSendEvent {
	res := &TeamSendEvent{
		Team: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewTeamProducer will stream data from the channel
func NewTeamProducer(producer event.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		ctx := context.Background()
		for item := range ch {
			if object, ok := item.Object().(*Team); ok {
				binary, codec, err := object.ToAvroBinary()
				if err != nil {
					errors <- fmt.Errorf("error encoding %s to avro binary data. %v", object.String(), err)
					return
				}
				headers := map[string]string{
					"customer_id": object.CustomerID,
					"model":       TeamModelName.String(),
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
				errors <- fmt.Errorf("invalid event received. expected an object of type customer.Team but received on of type %v", reflect.TypeOf(item.Object()))
			}
		}
	}()
	return done
}

// NewTeamConsumer will stream data from the topic into the provided channel
func NewTeamConsumer(consumer event.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) {
	consumer.Consume(event.ConsumerCallback{
		OnDataReceived: func(msg event.Message) error {
			var object Team
			if err := json.Unmarshal(msg.Value, &object); err != nil {
				return fmt.Errorf("error unmarshaling json data into customer.Team: %s", err)
			}
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &TeamReceiveEvent{&object, msg}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
	})
}

// TeamReceiveEvent is an event detail for receiving data
type TeamReceiveEvent struct {
	Team    *Team
	message event.Message
}

var _ datamodel.ModelReceiveEvent = (*TeamReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *TeamReceiveEvent) Object() datamodel.Model {
	return e.Team
}

// Message returns the underlying message data for the event
func (e *TeamReceiveEvent) Message() event.Message {
	return e.message
}

// TeamProducer implements the datamodel.ModelEventProducer
type TeamProducer struct {
	ch   chan datamodel.ModelSendEvent
	done <-chan bool
}

var _ datamodel.ModelEventProducer = (*TeamProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *TeamProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *TeamProducer) Close() error {
	close(p.ch)
	<-p.done
	return nil
}

// NewProducerChannel returns a channel which can be used for producing Model events
func (o *Team) NewProducerChannel(producer event.Producer, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent)
	return &TeamProducer{
		ch:   ch,
		done: NewTeamProducer(producer, ch, errors),
	}
}

// NewTeamProducerChannel returns a channel which can be used for producing Model events
func NewTeamProducerChannel(producer event.Producer, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent)
	return &TeamProducer{
		ch:   ch,
		done: NewTeamProducer(producer, ch, errors),
	}
}

// TeamConsumer implements the datamodel.ModelEventConsumer
type TeamConsumer struct {
	ch chan datamodel.ModelReceiveEvent
}

var _ datamodel.ModelEventConsumer = (*TeamConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *TeamConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *TeamConsumer) Close() error {
	close(c.ch)
	return nil
}

// NewConsumerChannel returns a consumer channel which can be used to consume Model events
func (o *Team) NewConsumerChannel(consumer event.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	NewTeamConsumer(consumer, ch, errors)
	return &TeamConsumer{
		ch: ch,
	}
}

// NewTeamConsumerChannel returns a consumer channel which can be used to consume Model events
func NewTeamConsumerChannel(consumer event.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	NewTeamConsumer(consumer, ch, errors)
	return &TeamConsumer{
		ch: ch,
	}
}
