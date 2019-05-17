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

// RepoMappingDefaultTopic is the default topic name
const RepoMappingDefaultTopic = "customer_RepoMapping_topic"

// RepoMappingDefaultStream is the default stream name
const RepoMappingDefaultStream = "customer_RepoMapping_stream"

// RepoMappingDefaultTable is the default table name
const RepoMappingDefaultTable = "customer_RepoMapping"

// RepoMapping blah blah blah
type RepoMapping struct {
	// built in types

	ID         string `json:"repo_mapping_id" yaml:"repo_mapping_id"`
	RefID      string `json:"ref_id" yaml:"ref_id"`
	RefType    string `json:"ref_type" yaml:"ref_type"`
	CustomerID string `json:"customer_id" yaml:"customer_id"`
	Hashcode   string `json:"hashcode" yaml:"hashcode"`

	// custom types

	// RepoID blah blah
	RepoID string `json:"repo_id" yaml:"repo_id"`
}

func toRepoMappingObject(o interface{}, isavro bool) interface{} {
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
	case *RepoMapping:
		return v.ToMap()
	case RepoMapping:
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
			arr = append(arr, toRepoMappingObject(av, isavro))
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of RepoMapping
func (o *RepoMapping) String() string {
	return fmt.Sprintf("customer.RepoMapping<%s>", o.ID)
}

func (o *RepoMapping) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *RepoMapping) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("RepoMapping", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetRefID returns the RefID for the object
func (o *RepoMapping) GetRefID() string {
	return o.RefID
}

// MarshalJSON returns the bytes for marshaling to json
func (o *RepoMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *RepoMapping) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	// make sure that these have values if empty
	o.setDefaults()
	return nil
}

var cachedCodecRepoMapping *goavro.Codec

// ToAvroBinary returns the data as Avro binary data
func (o *RepoMapping) ToAvroBinary() ([]byte, *goavro.Codec, error) {
	if cachedCodecRepoMapping == nil {
		c, err := CreateRepoMappingAvroSchema()
		if err != nil {
			return nil, nil, err
		}
		cachedCodecRepoMapping = c
	}
	kv := o.ToMap(true)
	jbuf, _ := json.Marshal(kv)
	native, _, err := cachedCodecRepoMapping.NativeFromTextual(jbuf)
	if err != nil {
		return nil, nil, err
	}
	// Convert native Go form to binary Avro data
	buf, err := cachedCodecRepoMapping.BinaryFromNative(nil, native)
	return buf, cachedCodecRepoMapping, err
}

// Stringify returns the object in JSON format as a string
func (o *RepoMapping) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two RepoMapping objects are equal
func (o *RepoMapping) IsEqual(other *RepoMapping) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *RepoMapping) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	return map[string]interface{}{
		"repo_mapping_id": o.GetID(),
		"ref_id":          o.GetRefID(),
		"ref_type":        o.RefType,
		"customer_id":     o.CustomerID,
		"hashcode":        o.Hash(),
		"repo_id":         toRepoMappingObject(o.RepoID, isavro),
	}
}

// FromMap attempts to load data into object from a map
func (o *RepoMapping) FromMap(kv map[string]interface{}) {
	if val, ok := kv["repo_mapping_id"].(string); ok {
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
	if val, ok := kv["repo_id"].(string); ok {
		o.RepoID = val
	} else {
		val := kv["repo_id"]
		if val == nil {
			o.RepoID = ""
		} else {
			o.RepoID = fmt.Sprintf("%v", val)
		}
	}
	// make sure that these have values if empty
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *RepoMapping) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.GetID())
	args = append(args, o.GetRefID())
	args = append(args, o.RefType)
	args = append(args, o.RepoID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// CreateRepoMappingAvroSchemaSpec creates the avro schema specification for RepoMapping
func CreateRepoMappingAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "customer",
		"name":         "RepoMapping",
		"connect.name": "customer.RepoMapping",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "repo_mapping_id",
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
				"name": "repo_id",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// CreateRepoMappingAvroSchema creates the avro schema for RepoMapping
func CreateRepoMappingAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(CreateRepoMappingAvroSchemaSpec())
}

// TransformRepoMappingFunc is a function for transforming RepoMapping during processing
type TransformRepoMappingFunc func(input *RepoMapping) (*RepoMapping, error)

// CreateRepoMappingPipe creates a pipe for processing RepoMapping items
func CreateRepoMappingPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformRepoMappingFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := CreateRepoMappingInputStream(input, errors)
	var stream chan RepoMapping
	if len(transforms) > 0 {
		stream = make(chan RepoMapping, 1000)
	} else {
		stream = inch
	}
	outdone := CreateRepoMappingOutputStream(output, stream, errors)
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

// CreateRepoMappingInputStreamDir creates a channel for reading RepoMapping as JSON newlines from a directory of files
func CreateRepoMappingInputStreamDir(dir string, errors chan<- error, transforms ...TransformRepoMappingFunc) (chan RepoMapping, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/customer/repo_mapping\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan RepoMapping)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for repo_mapping")
		ch := make(chan RepoMapping)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return CreateRepoMappingInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan RepoMapping)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// CreateRepoMappingInputStreamFile creates an channel for reading RepoMapping as JSON newlines from filename
func CreateRepoMappingInputStreamFile(filename string, errors chan<- error, transforms ...TransformRepoMappingFunc) (chan RepoMapping, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan RepoMapping)
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
			ch := make(chan RepoMapping)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return CreateRepoMappingInputStream(f, errors, transforms...)
}

// CreateRepoMappingInputStream creates an channel for reading RepoMapping as JSON newlines from stream
func CreateRepoMappingInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformRepoMappingFunc) (chan RepoMapping, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan RepoMapping, 1000)
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
			var item RepoMapping
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

// CreateRepoMappingOutputStreamDir will output json newlines from channel and save in dir
func CreateRepoMappingOutputStreamDir(dir string, ch chan RepoMapping, errors chan<- error, transforms ...TransformRepoMappingFunc) <-chan bool {
	fp := filepath.Join(dir, "/customer/repo_mapping\\.json(\\.gz)?$")
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
	return CreateRepoMappingOutputStream(gz, ch, errors, transforms...)
}

// CreateRepoMappingOutputStream will output json newlines from channel to the stream
func CreateRepoMappingOutputStream(stream io.WriteCloser, ch chan RepoMapping, errors chan<- error, transforms ...TransformRepoMappingFunc) <-chan bool {
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

// CreateRepoMappingProducer will stream data from the channel
func CreateRepoMappingProducer(producer util.Producer, ch chan RepoMapping, errors chan<- error) <-chan bool {
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

// CreateRepoMappingConsumer will stream data from the default topic into the provided channel
func CreateRepoMappingConsumer(factory util.ConsumerFactory, topic string, ch chan RepoMapping, errors chan<- error) (<-chan bool, chan<- bool) {
	return CreateRepoMappingConsumerForTopic(factory, RepoMappingDefaultTopic, ch, errors)
}

// CreateRepoMappingConsumerForTopic will stream data from the topic into the provided channel
func CreateRepoMappingConsumerForTopic(factory util.ConsumerFactory, topic string, ch chan RepoMapping, errors chan<- error) (<-chan bool, chan<- bool) {
	done := make(chan bool, 1)
	closed := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		callback := util.ConsumerCallback{
			OnDataReceived: func(key []byte, value []byte) error {
				var object RepoMapping
				if err := json.Unmarshal(value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into RepoMapping: %s", err)
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
