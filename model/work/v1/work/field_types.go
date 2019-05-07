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
	"regexp"
	"strings"

	"github.com/linkedin/goavro"
	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/integration-sdk/util"
)

// FieldTypesDefaultTopic is the default topic name
const FieldTypesDefaultTopic = "work_FieldTypes_topic"

// FieldTypesDefaultStream is the default stream name
const FieldTypesDefaultStream = "work_FieldTypes_stream"

// FieldTypesDefaultTable is the default table name
const FieldTypesDefaultTable = "work_FieldTypes"

// FieldTypes user defined fields
type FieldTypes struct {
	// built in types

	ID         string `json:"id" yaml:"id"`
	RefID      string `json:"ref_id" yaml:"ref_id"`
	RefType    string `json:"ref_type" yaml:"ref_type"`
	CustomerID string `json:"customer_id" yaml:"customer_id"`
	Hashcode   string `json:"hashcode" yaml:"hashcode"`

	// custom types

	Name string `json:"name" yaml:"name"`
	Key  string `json:"key" yaml:"key"`
}

// String returns a string representation of FieldTypes
func (o *FieldTypes) String() string {
	return fmt.Sprintf("work.v1.FieldTypes<%s>", o.ID)
}

func (o *FieldTypes) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *FieldTypes) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("FieldTypes", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetRefID returns the RefID for the object
func (o *FieldTypes) GetRefID() string {
	return o.RefID
}

// MarshalJSON returns the bytes for marshaling to json
func (o *FieldTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *FieldTypes) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	// make sure that these have values if empty
	o.setDefaults()
	return nil
}

// Stringify returns the object in JSON format as a string
func (o *FieldTypes) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two FieldTypes objects are equal
func (o *FieldTypes) IsEqual(other *FieldTypes) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *FieldTypes) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":          o.GetID(),
		"ref_id":      o.GetRefID(),
		"ref_type":    o.RefType,
		"customer_id": o.CustomerID,
		"hashcode":    o.Hash(),
		"name":        o.Name,
		"key":         o.Key,
	}
}

// FromMap attempts to load data into object from a map
func (o *FieldTypes) FromMap(kv map[string]interface{}) {
	if val, ok := kv["id"].(string); ok {
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
func (o *FieldTypes) Hash() string {
	if o.Hashcode == "" {
		args := make([]interface{}, 0)
		args = append(args, o.GetID())
		args = append(args, o.GetRefID())
		args = append(args, o.RefType)
		args = append(args, o.Name)
		args = append(args, o.Key)
		o.Hashcode = hash.Values(args...)
	}
	return o.Hashcode
}

// CreateFieldTypesAvroSchemaSpec creates the avro schema specification for FieldTypes
func CreateFieldTypesAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "work.v1",
		"name":         "FieldTypes",
		"connect.name": "work.v1.FieldTypes",
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

// CreateFieldTypesAvroSchema creates the avro schema for FieldTypes
func CreateFieldTypesAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(CreateFieldTypesAvroSchemaSpec())
}

// CreateFieldTypesKQLStreamSQL creates KQL Stream SQL for FieldTypes
func CreateFieldTypesKQLStreamSQL() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("CREATE STREAM %s ", FieldTypesDefaultStream))
	builder.WriteString(fmt.Sprintf("WITH (KAFKA_TOPIC='%s', VALUE_FORMAT='AVRO', KEY='id'", FieldTypesDefaultTopic))
	builder.WriteString(");")
	return builder.String()
}

// CreateFieldTypesKQLTableSQL creates KQL Table SQL for FieldTypes
func CreateFieldTypesKQLTableSQL() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("CREATE TABLE %s ", FieldTypesDefaultTable))
	builder.WriteString(fmt.Sprintf("WITH (KAFKA_TOPIC='%s', VALUE_FORMAT='AVRO', KEY='id'", FieldTypesDefaultTopic))
	builder.WriteString(");")
	return builder.String()
}

// TransformFieldTypesFunc is a function for transforming FieldTypes during processing
type TransformFieldTypesFunc func(input *FieldTypes) (*FieldTypes, error)

// CreateFieldTypesPipe creates a pipe for processing FieldTypes items
func CreateFieldTypesPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformFieldTypesFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := CreateFieldTypesInputStream(input, errors)
	var stream chan FieldTypes
	if len(transforms) > 0 {
		stream = make(chan FieldTypes, 1000)
	} else {
		stream = inch
	}
	outdone := CreateFieldTypesOutputStream(output, stream, errors)
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

// CreateFieldTypesInputStreamDir creates a channel for reading FieldTypes as JSON newlines from a directory of files
func CreateFieldTypesInputStreamDir(dir string, errors chan<- error, transforms ...TransformFieldTypesFunc) (chan FieldTypes, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/work/v1/field_types\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan FieldTypes)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for field_types")
		ch := make(chan FieldTypes)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return CreateFieldTypesInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan FieldTypes)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// CreateFieldTypesInputStreamFile creates an channel for reading FieldTypes as JSON newlines from filename
func CreateFieldTypesInputStreamFile(filename string, errors chan<- error, transforms ...TransformFieldTypesFunc) (chan FieldTypes, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan FieldTypes)
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
			ch := make(chan FieldTypes)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return CreateFieldTypesInputStream(f, errors, transforms...)
}

// CreateFieldTypesInputStream creates an channel for reading FieldTypes as JSON newlines from stream
func CreateFieldTypesInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformFieldTypesFunc) (chan FieldTypes, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan FieldTypes, 1000)
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
			var item FieldTypes
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

// CreateFieldTypesOutputStreamDir will output json newlines from channel and save in dir
func CreateFieldTypesOutputStreamDir(dir string, ch chan FieldTypes, errors chan<- error, transforms ...TransformFieldTypesFunc) <-chan bool {
	fp := filepath.Join(dir, "/work/v1/field_types\\.json(\\.gz)?$")
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
	return CreateFieldTypesOutputStream(gz, ch, errors, transforms...)
}

// CreateFieldTypesOutputStream will output json newlines from channel to the stream
func CreateFieldTypesOutputStream(stream io.WriteCloser, ch chan FieldTypes, errors chan<- error, transforms ...TransformFieldTypesFunc) <-chan bool {
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

// CreateFieldTypesProducer will stream data from the channel
func CreateFieldTypesProducer(producer util.Producer, ch chan FieldTypes, errors chan<- error) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		schemaspec := CreateFieldTypesAvroSchemaSpec()
		ctx := context.Background()
		for item := range ch {
			if err := producer.Send(ctx, schemaspec, []byte(item.ID), []byte(item.Stringify())); err != nil {
				errors <- fmt.Errorf("error sending %s. %v", item.String(), err)
			}
		}
	}()
	return done
}

// CreateFieldTypesConsumer will stream data from the default topic into the provided channel
func CreateFieldTypesConsumer(factory util.ConsumerFactory, topic string, ch chan FieldTypes, errors chan<- error) (<-chan bool, chan<- bool) {
	return CreateFieldTypesConsumerForTopic(factory, FieldTypesDefaultTopic, ch, errors)
}

// CreateFieldTypesConsumerForTopic will stream data from the topic into the provided channel
func CreateFieldTypesConsumerForTopic(factory util.ConsumerFactory, topic string, ch chan FieldTypes, errors chan<- error) (<-chan bool, chan<- bool) {
	done := make(chan bool, 1)
	closed := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		callback := util.ConsumerCallback{
			OnDataReceived: func(key []byte, value []byte) error {
				var object FieldTypes
				if err := json.Unmarshal(value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into FieldTypes: %s", err)
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
