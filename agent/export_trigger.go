// DO NOT EDIT -- generated code

// Package agent - the agent communicates with pinpoint cloud to send integration data
package agent

import (
	"bufio"
	"bytes"
	"compress/gzip"
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
	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
)

const (

	// ExportTriggerModelName is the model name
	ExportTriggerModelName datamodel.ModelNameType = "agent.ExportTrigger"
)

const (
	// ExportTriggerCustomerIDColumn is the customer_id column name
	ExportTriggerCustomerIDColumn = "customer_id"
	// ExportTriggerIDColumn is the id column name
	ExportTriggerIDColumn = "id"
	// ExportTriggerRefIDColumn is the ref_id column name
	ExportTriggerRefIDColumn = "ref_id"
	// ExportTriggerRefTypeColumn is the ref_type column name
	ExportTriggerRefTypeColumn = "ref_type"
)

// ExportTrigger used to trigger an agent.ExportRequest
type ExportTrigger struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*ExportTrigger)(nil)

func toExportTriggerObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toExportTriggerObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *ExportTrigger:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

// String returns a string representation of ExportTrigger
func (o *ExportTrigger) String() string {
	return fmt.Sprintf("agent.ExportTrigger<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *ExportTrigger) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetModelName returns the name of the model
func (o *ExportTrigger) GetModelName() datamodel.ModelNameType {
	return ExportTriggerModelName
}

func (o *ExportTrigger) setDefaults(frommap bool) {

	o.GetID()

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *ExportTrigger) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("ExportTrigger", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *ExportTrigger) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *ExportTrigger) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *ExportTrigger) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *ExportTrigger) IsMaterialized() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *ExportTrigger) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *ExportTrigger) IsEvented() bool {
	return false
}

// GetTopicConfig returns the topic config object
func (o *ExportTrigger) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetStateKey returns a key for use in state store
func (o *ExportTrigger) GetStateKey() string {
	key := ""
	return fmt.Sprintf("%s_%s", key, o.GetID())
}

// GetCustomerID will return the customer_id
func (o *ExportTrigger) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of ExportTrigger
func (o *ExportTrigger) Clone() datamodel.Model {
	c := new(ExportTrigger)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *ExportTrigger) Anon() datamodel.Model {
	c := new(ExportTrigger)
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
func (o *ExportTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *ExportTrigger) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

var cachedCodecExportTrigger *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *ExportTrigger) GetAvroCodec() *goavro.Codec {
	if cachedCodecExportTrigger == nil {
		c, err := GetExportTriggerAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecExportTrigger = c
	}
	return cachedCodecExportTrigger
}

// ToAvroBinary returns the data as Avro binary data
func (o *ExportTrigger) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *ExportTrigger) FromAvroBinary(value []byte) error {
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
func (o *ExportTrigger) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two ExportTrigger objects are equal
func (o *ExportTrigger) IsEqual(other *ExportTrigger) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *ExportTrigger) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	o.setDefaults(true)
	return map[string]interface{}{
		"customer_id": toExportTriggerObject(o.CustomerID, isavro, false, "string"),
		"id":          toExportTriggerObject(o.ID, isavro, false, "string"),
		"ref_id":      toExportTriggerObject(o.RefID, isavro, false, "string"),
		"ref_type":    toExportTriggerObject(o.RefType, isavro, false, "string"),
		"hashcode":    toExportTriggerObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *ExportTrigger) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
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
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *ExportTrigger) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetExportTriggerAvroSchemaSpec creates the avro schema specification for ExportTrigger
func GetExportTriggerAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "agent",
		"name":      "ExportTrigger",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "hashcode",
				"type": "string",
			},
			map[string]interface{}{
				"name": "customer_id",
				"type": "string",
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
		},
	}
	return pjson.Stringify(spec, true)
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *ExportTrigger) GetEventAPIConfig() datamodel.EventAPIConfig {
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

// GetExportTriggerAvroSchema creates the avro schema for ExportTrigger
func GetExportTriggerAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetExportTriggerAvroSchemaSpec())
}

// TransformExportTriggerFunc is a function for transforming ExportTrigger during processing
type TransformExportTriggerFunc func(input *ExportTrigger) (*ExportTrigger, error)

// NewExportTriggerPipe creates a pipe for processing ExportTrigger items
func NewExportTriggerPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformExportTriggerFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewExportTriggerInputStream(input, errors)
	var stream chan ExportTrigger
	if len(transforms) > 0 {
		stream = make(chan ExportTrigger, 1000)
	} else {
		stream = inch
	}
	outdone := NewExportTriggerOutputStream(output, stream, errors)
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

// NewExportTriggerInputStreamDir creates a channel for reading ExportTrigger as JSON newlines from a directory of files
func NewExportTriggerInputStreamDir(dir string, errors chan<- error, transforms ...TransformExportTriggerFunc) (chan ExportTrigger, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/agent/export_trigger\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan ExportTrigger)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for export_trigger")
		ch := make(chan ExportTrigger)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewExportTriggerInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan ExportTrigger)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewExportTriggerInputStreamFile creates an channel for reading ExportTrigger as JSON newlines from filename
func NewExportTriggerInputStreamFile(filename string, errors chan<- error, transforms ...TransformExportTriggerFunc) (chan ExportTrigger, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan ExportTrigger)
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
			ch := make(chan ExportTrigger)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewExportTriggerInputStream(f, errors, transforms...)
}

// NewExportTriggerInputStream creates an channel for reading ExportTrigger as JSON newlines from stream
func NewExportTriggerInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformExportTriggerFunc) (chan ExportTrigger, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan ExportTrigger, 1000)
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
			var item ExportTrigger
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

// NewExportTriggerOutputStreamDir will output json newlines from channel and save in dir
func NewExportTriggerOutputStreamDir(dir string, ch chan ExportTrigger, errors chan<- error, transforms ...TransformExportTriggerFunc) <-chan bool {
	fp := filepath.Join(dir, "/agent/export_trigger\\.json(\\.gz)?$")
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
	return NewExportTriggerOutputStream(gz, ch, errors, transforms...)
}

// NewExportTriggerOutputStream will output json newlines from channel to the stream
func NewExportTriggerOutputStream(stream io.WriteCloser, ch chan ExportTrigger, errors chan<- error, transforms ...TransformExportTriggerFunc) <-chan bool {
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
