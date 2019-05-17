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

	"github.com/linkedin/goavro"
	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/integration-sdk/util"
)

// UserMappingDefaultTopic is the default topic name
const UserMappingDefaultTopic = "customer_UserMapping_topic"

// UserMappingDefaultStream is the default stream name
const UserMappingDefaultStream = "customer_UserMapping_stream"

// UserMappingDefaultTable is the default table name
const UserMappingDefaultTable = "customer_UserMapping"

// UserMapping blah blah blah
type UserMapping struct {
	// built in types

	ID         string `json:"user_mapping_id" yaml:"user_mapping_id"`
	RefID      string `json:"ref_id" yaml:"ref_id"`
	RefType    string `json:"ref_type" yaml:"ref_type"`
	CustomerID string `json:"customer_id" yaml:"customer_id"`
	Hashcode   string `json:"hashcode" yaml:"hashcode"`

	// custom types

	// UserID _
	UserID string `json:"user_id" yaml:"user_id"`
}

func toUserMappingObject(o interface{}, isavro bool) interface{} {
	if o == nil {
		return nil
	}
	switch v := o.(type) {
	case nil:
		return nil
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		return v
	case *string, *int, *int8, *int16, *int32, *int64, *float32, *float64, *bool:
		return v
	case map[string]interface{}:
		return o
	case *map[string]interface{}:
		return v
	case *UserMapping:
		return v.ToMap()
	case UserMapping:
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
			arr = append(arr, toUserMappingObject(av, isavro))
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of UserMapping
func (o *UserMapping) String() string {
	return fmt.Sprintf("customer.UserMapping<%s>", o.ID)
}

func (o *UserMapping) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *UserMapping) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("UserMapping", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetRefID returns the RefID for the object
func (o *UserMapping) GetRefID() string {
	return o.RefID
}

// MarshalJSON returns the bytes for marshaling to json
func (o *UserMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *UserMapping) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	// make sure that these have values if empty
	o.setDefaults()
	return nil
}

var cachedCodecUserMapping *goavro.Codec

// ToAvroBinary returns the data as Avro binary data
func (o *UserMapping) ToAvroBinary() ([]byte, *goavro.Codec, error) {
	if cachedCodecUserMapping == nil {
		c, err := CreateUserMappingAvroSchema()
		if err != nil {
			return nil, nil, err
		}
		cachedCodecUserMapping = c
	}
	kv := o.ToMap(true)
	jbuf, _ := json.Marshal(kv)
	native, _, err := cachedCodecUserMapping.NativeFromTextual(jbuf)
	if err != nil {
		return nil, nil, err
	}
	// Convert native Go form to binary Avro data
	buf, err := cachedCodecUserMapping.BinaryFromNative(nil, native)
	return buf, cachedCodecUserMapping, err
}

// Stringify returns the object in JSON format as a string
func (o *UserMapping) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two UserMapping objects are equal
func (o *UserMapping) IsEqual(other *UserMapping) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *UserMapping) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	return map[string]interface{}{
		"user_mapping_id": o.GetID(),
		"ref_id":          o.GetRefID(),
		"ref_type":        o.RefType,
		"customer_id":     o.CustomerID,
		"hashcode":        o.Hash(),
		"user_id":         toUserMappingObject(o.UserID, isavro),
	}
}

// FromMap attempts to load data into object from a map
func (o *UserMapping) FromMap(kv map[string]interface{}) {
	if val, ok := kv["user_mapping_id"].(string); ok {
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
	if val, ok := kv["user_id"].(string); ok {
		o.UserID = val
	} else {
		val := kv["user_id"]
		if val == nil {
			o.UserID = ""
		} else {
			o.UserID = fmt.Sprintf("%v", val)
		}
	}
	// make sure that these have values if empty
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *UserMapping) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.GetID())
	args = append(args, o.GetRefID())
	args = append(args, o.RefType)
	args = append(args, o.UserID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// CreateUserMappingAvroSchemaSpec creates the avro schema specification for UserMapping
func CreateUserMappingAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "customer",
		"name":         "UserMapping",
		"connect.name": "customer.UserMapping",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "user_mapping_id",
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
				"name": "user_id",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// CreateUserMappingAvroSchema creates the avro schema for UserMapping
func CreateUserMappingAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(CreateUserMappingAvroSchemaSpec())
}

// TransformUserMappingFunc is a function for transforming UserMapping during processing
type TransformUserMappingFunc func(input *UserMapping) (*UserMapping, error)

// CreateUserMappingPipe creates a pipe for processing UserMapping items
func CreateUserMappingPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformUserMappingFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := CreateUserMappingInputStream(input, errors)
	var stream chan UserMapping
	if len(transforms) > 0 {
		stream = make(chan UserMapping, 1000)
	} else {
		stream = inch
	}
	outdone := CreateUserMappingOutputStream(output, stream, errors)
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

// CreateUserMappingInputStreamDir creates a channel for reading UserMapping as JSON newlines from a directory of files
func CreateUserMappingInputStreamDir(dir string, errors chan<- error, transforms ...TransformUserMappingFunc) (chan UserMapping, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/customer/user_mapping\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan UserMapping)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for user_mapping")
		ch := make(chan UserMapping)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return CreateUserMappingInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan UserMapping)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// CreateUserMappingInputStreamFile creates an channel for reading UserMapping as JSON newlines from filename
func CreateUserMappingInputStreamFile(filename string, errors chan<- error, transforms ...TransformUserMappingFunc) (chan UserMapping, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan UserMapping)
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
			ch := make(chan UserMapping)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return CreateUserMappingInputStream(f, errors, transforms...)
}

// CreateUserMappingInputStream creates an channel for reading UserMapping as JSON newlines from stream
func CreateUserMappingInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformUserMappingFunc) (chan UserMapping, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan UserMapping, 1000)
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
			var item UserMapping
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

// CreateUserMappingOutputStreamDir will output json newlines from channel and save in dir
func CreateUserMappingOutputStreamDir(dir string, ch chan UserMapping, errors chan<- error, transforms ...TransformUserMappingFunc) <-chan bool {
	fp := filepath.Join(dir, "/customer/user_mapping\\.json(\\.gz)?$")
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
	return CreateUserMappingOutputStream(gz, ch, errors, transforms...)
}

// CreateUserMappingOutputStream will output json newlines from channel to the stream
func CreateUserMappingOutputStream(stream io.WriteCloser, ch chan UserMapping, errors chan<- error, transforms ...TransformUserMappingFunc) <-chan bool {
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

// CreateUserMappingProducer will stream data from the channel
func CreateUserMappingProducer(producer util.Producer, ch chan UserMapping, errors chan<- error) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		ctx := context.Background()
		for item := range ch {
			binary, codec, err := item.ToAvroBinary()
			if err != nil {
				errors <- fmt.Errorf("error encoding %s to avro binary data. %v", item.String(), err)
				return
			}
			if err := producer.Send(ctx, codec, []byte(item.ID), binary); err != nil {
				errors <- fmt.Errorf("error sending %s. %v", item.String(), err)
			}
		}
	}()
	return done
}

// CreateUserMappingConsumer will stream data from the default topic into the provided channel
func CreateUserMappingConsumer(factory util.ConsumerFactory, topic string, ch chan UserMapping, errors chan<- error) (<-chan bool, chan<- bool) {
	return CreateUserMappingConsumerForTopic(factory, UserMappingDefaultTopic, ch, errors)
}

// CreateUserMappingConsumerForTopic will stream data from the topic into the provided channel
func CreateUserMappingConsumerForTopic(factory util.ConsumerFactory, topic string, ch chan UserMapping, errors chan<- error) (<-chan bool, chan<- bool) {
	done := make(chan bool, 1)
	closed := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		callback := util.ConsumerCallback{
			OnDataReceived: func(key []byte, value []byte) error {
				var object UserMapping
				if err := json.Unmarshal(value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into UserMapping: %s", err)
				}
				ch <- object
				return nil
			},
			OnErrorReceived: func(err error) {
				errors <- err
			},
		}
		consumer, err := factory.CreateConsumer(topic, callback)
		if err != nil {
			errors <- err
			return
		}
		select {
		case <-closed:
			consumer.Close()
		}
	}()
	return done, closed
}
