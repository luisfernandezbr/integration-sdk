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

	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/integration-sdk/util"
)

// CustomFieldDefaultTopic is the default topic name
const CustomFieldDefaultTopic = "work_CustomField_topic"

// CustomFieldDefaultStream is the default stream name
const CustomFieldDefaultStream = "work_CustomField_stream"

// CustomFieldDefaultTable is the default table name
const CustomFieldDefaultTable = "work_CustomField"

// CustomField user defined fields
type CustomField struct {
	// built in types

	ID         string `json:"custom_field_id" yaml:"custom_field_id"`
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

func toCustomFieldObject(o interface{}, isavro bool) interface{} {
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
	case *CustomField:
		return v.ToMap()
	case CustomField:
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
			arr = append(arr, toCustomFieldObject(av, isavro))
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of CustomField
func (o *CustomField) String() string {
	return fmt.Sprintf("work.CustomField<%s>", o.ID)
}

func (o *CustomField) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *CustomField) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("CustomField", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetRefID returns the RefID for the object
func (o *CustomField) GetRefID() string {
	return o.RefID
}

// MarshalJSON returns the bytes for marshaling to json
func (o *CustomField) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *CustomField) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	// make sure that these have values if empty
	o.setDefaults()
	return nil
}

var cachedCodecCustomField *goavro.Codec

// ToAvroBinary returns the data as Avro binary data
func (o *CustomField) ToAvroBinary() ([]byte, *goavro.Codec, error) {
	if cachedCodecCustomField == nil {
		c, err := CreateCustomFieldAvroSchema()
		if err != nil {
			return nil, nil, err
		}
		cachedCodecCustomField = c
	}
	kv := o.ToMap(true)
	jbuf, _ := json.Marshal(kv)
	native, _, err := cachedCodecCustomField.NativeFromTextual(jbuf)
	if err != nil {
		return nil, nil, err
	}
	// Convert native Go form to binary Avro data
	buf, err := cachedCodecCustomField.BinaryFromNative(nil, native)
	return buf, cachedCodecCustomField, err
}

// Stringify returns the object in JSON format as a string
func (o *CustomField) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two CustomField objects are equal
func (o *CustomField) IsEqual(other *CustomField) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *CustomField) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	return map[string]interface{}{
		"custom_field_id": o.GetID(),
		"ref_id":          o.GetRefID(),
		"ref_type":        o.RefType,
		"customer_id":     o.CustomerID,
		"hashcode":        o.Hash(),
		"name":            toCustomFieldObject(o.Name, isavro),
		"key":             toCustomFieldObject(o.Key, isavro),
	}
}

// FromMap attempts to load data into object from a map
func (o *CustomField) FromMap(kv map[string]interface{}) {
	if val, ok := kv["custom_field_id"].(string); ok {
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
func (o *CustomField) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.GetID())
	args = append(args, o.GetRefID())
	args = append(args, o.RefType)
	args = append(args, o.CustomerID)
	args = append(args, o.Name)
	args = append(args, o.Key)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// CreateCustomFieldAvroSchemaSpec creates the avro schema specification for CustomField
func CreateCustomFieldAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "work",
		"name":         "CustomField",
		"connect.name": "work.CustomField",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "custom_field_id",
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

// CreateCustomFieldAvroSchema creates the avro schema for CustomField
func CreateCustomFieldAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(CreateCustomFieldAvroSchemaSpec())
}

// TransformCustomFieldFunc is a function for transforming CustomField during processing
type TransformCustomFieldFunc func(input *CustomField) (*CustomField, error)

// CreateCustomFieldPipe creates a pipe for processing CustomField items
func CreateCustomFieldPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformCustomFieldFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := CreateCustomFieldInputStream(input, errors)
	var stream chan CustomField
	if len(transforms) > 0 {
		stream = make(chan CustomField, 1000)
	} else {
		stream = inch
	}
	outdone := CreateCustomFieldOutputStream(output, stream, errors)
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

// CreateCustomFieldInputStreamDir creates a channel for reading CustomField as JSON newlines from a directory of files
func CreateCustomFieldInputStreamDir(dir string, errors chan<- error, transforms ...TransformCustomFieldFunc) (chan CustomField, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/work/custom_field\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan CustomField)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for custom_field")
		ch := make(chan CustomField)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return CreateCustomFieldInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan CustomField)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// CreateCustomFieldInputStreamFile creates an channel for reading CustomField as JSON newlines from filename
func CreateCustomFieldInputStreamFile(filename string, errors chan<- error, transforms ...TransformCustomFieldFunc) (chan CustomField, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan CustomField)
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
			ch := make(chan CustomField)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return CreateCustomFieldInputStream(f, errors, transforms...)
}

// CreateCustomFieldInputStream creates an channel for reading CustomField as JSON newlines from stream
func CreateCustomFieldInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformCustomFieldFunc) (chan CustomField, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan CustomField, 1000)
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
			var item CustomField
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

// CreateCustomFieldOutputStreamDir will output json newlines from channel and save in dir
func CreateCustomFieldOutputStreamDir(dir string, ch chan CustomField, errors chan<- error, transforms ...TransformCustomFieldFunc) <-chan bool {
	fp := filepath.Join(dir, "/work/custom_field\\.json(\\.gz)?$")
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
	return CreateCustomFieldOutputStream(gz, ch, errors, transforms...)
}

// CreateCustomFieldOutputStream will output json newlines from channel to the stream
func CreateCustomFieldOutputStream(stream io.WriteCloser, ch chan CustomField, errors chan<- error, transforms ...TransformCustomFieldFunc) <-chan bool {
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

// CreateCustomFieldProducer will stream data from the channel
func CreateCustomFieldProducer(producer util.Producer, ch chan CustomField, errors chan<- error) <-chan bool {
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

// CreateCustomFieldConsumer will stream data from the default topic into the provided channel
func CreateCustomFieldConsumer(factory util.ConsumerFactory, topic string, ch chan CustomField, errors chan<- error) (<-chan bool, chan<- bool) {
	return CreateCustomFieldConsumerForTopic(factory, CustomFieldDefaultTopic, ch, errors)
}

// CreateCustomFieldConsumerForTopic will stream data from the topic into the provided channel
func CreateCustomFieldConsumerForTopic(factory util.ConsumerFactory, topic string, ch chan CustomField, errors chan<- error) (<-chan bool, chan<- bool) {
	done := make(chan bool, 1)
	closed := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		callback := util.ConsumerCallback{
			OnDataReceived: func(key []byte, value []byte) error {
				var object CustomField
				if err := json.Unmarshal(value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into CustomField: %s", err)
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
