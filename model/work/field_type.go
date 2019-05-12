// DO NOT EDIT -- generated code

// Package work - the system which contains project work
package work

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

// FieldTypeDefaultTopic is the default topic name
const FieldTypeDefaultTopic = "work_FieldType_topic"

// FieldTypeDefaultStream is the default stream name
const FieldTypeDefaultStream = "work_FieldType_stream"

// FieldTypeDefaultTable is the default table name
const FieldTypeDefaultTable = "work_FieldType"

// FieldType user defined fields
type FieldType struct {
	// built in types

	ID         string `json:"field_type_id" yaml:"field_type_id"`
	RefID      string `json:"ref_id" yaml:"ref_id"`
	RefType    string `json:"ref_type" yaml:"ref_type"`
	CustomerID string `json:"customer_id" yaml:"customer_id"`
	Hashcode   string `json:"hashcode" yaml:"hashcode"`

	// custom types

	// Name the name of the field
	Name string `json:"name" yaml:"name"`
	// Key key of the field
	Key string `json:"key" yaml:"key"`
}

func toFieldTypeObject(o interface{}, isavro bool) interface{} {
	if o == nil {
		return nil
	}
	switch o.(type) {
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		return o
	case *string:
		return *(o.(*string))
	case *int:
		return *(o.(*int))
	case *int8:
		return *(o.(*int8))
	case *int16:
		return *(o.(*int16))
	case *int32:
		return *(o.(*int32))
	case *int64:
		return *(o.(*int64))
	case *float32:
		return *(o.(*float32))
	case *float64:
		return *(o.(*float64))
	case *bool:
		return *(o.(*bool))
	case map[string]interface{}:
		return o
	case *map[string]interface{}:
		return *(o.(*interface{}))
	case *FieldType:
		val := o.(*FieldType)
		return val.ToMap()
	case FieldType:
		val := o.(FieldType)
		return val.ToMap()
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
			arr = append(arr, toFieldTypeObject(av, isavro))
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of FieldType
func (o *FieldType) String() string {
	return fmt.Sprintf("work.FieldType<%s>", o.ID)
}

func (o *FieldType) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *FieldType) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("FieldType", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetRefID returns the RefID for the object
func (o *FieldType) GetRefID() string {
	return o.RefID
}

// MarshalJSON returns the bytes for marshaling to json
func (o *FieldType) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *FieldType) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	// make sure that these have values if empty
	o.setDefaults()
	return nil
}

var cachedCodecFieldType *goavro.Codec

// ToAvroBinary returns the data as Avro binary data
func (o *FieldType) ToAvroBinary() ([]byte, *goavro.Codec, error) {
	if cachedCodecFieldType == nil {
		c, err := CreateFieldTypeAvroSchema()
		if err != nil {
			return nil, nil, err
		}
		cachedCodecFieldType = c
	}
	kv := o.ToMap(true)
	jbuf, _ := json.Marshal(kv)
	native, _, err := cachedCodecFieldType.NativeFromTextual(jbuf)
	if err != nil {
		return nil, nil, err
	}
	// Convert native Go form to binary Avro data
	buf, err := cachedCodecFieldType.BinaryFromNative(nil, native)
	return buf, cachedCodecFieldType, err
}

// Stringify returns the object in JSON format as a string
func (o *FieldType) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two FieldType objects are equal
func (o *FieldType) IsEqual(other *FieldType) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *FieldType) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	return map[string]interface{}{
		"field_type_id": o.GetID(),
		"ref_id":        o.GetRefID(),
		"ref_type":      o.RefType,
		"customer_id":   o.CustomerID,
		"hashcode":      o.Hash(),
		"name":          toFieldTypeObject(o.Name, isavro),
		"key":           toFieldTypeObject(o.Key, isavro),
	}
}

// FromMap attempts to load data into object from a map
func (o *FieldType) FromMap(kv map[string]interface{}) {
	if val, ok := kv["field_type_id"].(string); ok {
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
	if val, ok := kv["key"].(string); ok {
		o.Key = val
	} else {
		val := kv["key"]
		if val == nil {
			o.Key = ""
		} else {
			o.Key = fmt.Sprintf("%v", val)
		}
	}
	// make sure that these have values if empty
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *FieldType) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.GetID())
	args = append(args, o.GetRefID())
	args = append(args, o.RefType)
	args = append(args, o.Name)
	args = append(args, o.Key)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// CreateFieldTypeAvroSchemaSpec creates the avro schema specification for FieldType
func CreateFieldTypeAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "work",
		"name":         "FieldType",
		"connect.name": "work.FieldType",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "field_type_id",
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
				"name": "key",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// CreateFieldTypeAvroSchema creates the avro schema for FieldType
func CreateFieldTypeAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(CreateFieldTypeAvroSchemaSpec())
}

// TransformFieldTypeFunc is a function for transforming FieldType during processing
type TransformFieldTypeFunc func(input *FieldType) (*FieldType, error)

// CreateFieldTypePipe creates a pipe for processing FieldType items
func CreateFieldTypePipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformFieldTypeFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := CreateFieldTypeInputStream(input, errors)
	var stream chan FieldType
	if len(transforms) > 0 {
		stream = make(chan FieldType, 1000)
	} else {
		stream = inch
	}
	outdone := CreateFieldTypeOutputStream(output, stream, errors)
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

// CreateFieldTypeInputStreamDir creates a channel for reading FieldType as JSON newlines from a directory of files
func CreateFieldTypeInputStreamDir(dir string, errors chan<- error, transforms ...TransformFieldTypeFunc) (chan FieldType, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/work/field_type\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan FieldType)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for field_type")
		ch := make(chan FieldType)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return CreateFieldTypeInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan FieldType)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// CreateFieldTypeInputStreamFile creates an channel for reading FieldType as JSON newlines from filename
func CreateFieldTypeInputStreamFile(filename string, errors chan<- error, transforms ...TransformFieldTypeFunc) (chan FieldType, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan FieldType)
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
			ch := make(chan FieldType)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return CreateFieldTypeInputStream(f, errors, transforms...)
}

// CreateFieldTypeInputStream creates an channel for reading FieldType as JSON newlines from stream
func CreateFieldTypeInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformFieldTypeFunc) (chan FieldType, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan FieldType, 1000)
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
			var item FieldType
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

// CreateFieldTypeOutputStreamDir will output json newlines from channel and save in dir
func CreateFieldTypeOutputStreamDir(dir string, ch chan FieldType, errors chan<- error, transforms ...TransformFieldTypeFunc) <-chan bool {
	fp := filepath.Join(dir, "/work/field_type\\.json(\\.gz)?$")
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
	return CreateFieldTypeOutputStream(gz, ch, errors, transforms...)
}

// CreateFieldTypeOutputStream will output json newlines from channel to the stream
func CreateFieldTypeOutputStream(stream io.WriteCloser, ch chan FieldType, errors chan<- error, transforms ...TransformFieldTypeFunc) <-chan bool {
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

// CreateFieldTypeProducer will stream data from the channel
func CreateFieldTypeProducer(producer util.Producer, ch chan FieldType, errors chan<- error) <-chan bool {
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

// CreateFieldTypeConsumer will stream data from the default topic into the provided channel
func CreateFieldTypeConsumer(factory util.ConsumerFactory, topic string, ch chan FieldType, errors chan<- error) (<-chan bool, chan<- bool) {
	return CreateFieldTypeConsumerForTopic(factory, FieldTypeDefaultTopic, ch, errors)
}

// CreateFieldTypeConsumerForTopic will stream data from the topic into the provided channel
func CreateFieldTypeConsumerForTopic(factory util.ConsumerFactory, topic string, ch chan FieldType, errors chan<- error) (<-chan bool, chan<- bool) {
	done := make(chan bool, 1)
	closed := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		callback := util.ConsumerCallback{
			OnDataReceived: func(key []byte, value []byte) error {
				var object FieldType
				if err := json.Unmarshal(value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into FieldType: %s", err)
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
