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

// ProjectDefaultTopic is the default topic name
const ProjectDefaultTopic = "work_Project_topic"

// ProjectDefaultStream is the default stream name
const ProjectDefaultStream = "work_Project_stream"

// ProjectDefaultTable is the default table name
const ProjectDefaultTable = "work_Project"

// Project the project holds work
type Project struct {
	// built in types

	ID         string `json:"project_id" yaml:"project_id"`
	RefID      string `json:"ref_id" yaml:"ref_id"`
	RefType    string `json:"ref_type" yaml:"ref_type"`
	CustomerID string `json:"customer_id" yaml:"customer_id"`
	Hashcode   string `json:"hashcode" yaml:"hashcode"`

	// custom types

	// Name the name of the project
	Name string `json:"name" yaml:"name"`
	// Identifier the common identifier for the project
	Identifier string `json:"identifier" yaml:"identifier"`
	// URL the url to the project home page
	URL string `json:"url" yaml:"url"`
}

func toProjectObject(o interface{}, isavro bool) interface{} {
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
	case *Project:
		val := o.(*Project)
		return val.ToMap()
	case Project:
		val := o.(Project)
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
			arr = append(arr, toProjectObject(av, isavro))
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of Project
func (o *Project) String() string {
	return fmt.Sprintf("work.Project<%s>", o.ID)
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

// GetRefID returns the RefID for the object
func (o *Project) GetRefID() string {
	return o.RefID
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

// ToAvroBinary returns the data as Avro binary data
func (o *Project) ToAvroBinary() ([]byte, *goavro.Codec, error) {
	if cachedCodecProject == nil {
		c, err := CreateProjectAvroSchema()
		if err != nil {
			return nil, nil, err
		}
		cachedCodecProject = c
	}
	kv := o.ToMap(true)
	jbuf, _ := json.Marshal(kv)
	native, _, err := cachedCodecProject.NativeFromTextual(jbuf)
	if err != nil {
		return nil, nil, err
	}
	// Convert native Go form to binary Avro data
	buf, err := cachedCodecProject.BinaryFromNative(nil, native)
	return buf, cachedCodecProject, err
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
	return map[string]interface{}{
		"project_id":  o.GetID(),
		"ref_id":      o.GetRefID(),
		"ref_type":    o.RefType,
		"customer_id": o.CustomerID,
		"hashcode":    o.Hash(),
		"name":        toProjectObject(o.Name, isavro),
		"identifier":  toProjectObject(o.Identifier, isavro),
		"url":         toProjectObject(o.URL, isavro),
	}
}

// FromMap attempts to load data into object from a map
func (o *Project) FromMap(kv map[string]interface{}) {
	if val, ok := kv["project_id"].(string); ok {
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
func (o *Project) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.GetID())
	args = append(args, o.GetRefID())
	args = append(args, o.RefType)
	args = append(args, o.Name)
	args = append(args, o.Identifier)
	args = append(args, o.URL)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// CreateProjectAvroSchemaSpec creates the avro schema specification for Project
func CreateProjectAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "work",
		"name":         "Project",
		"connect.name": "work.Project",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "project_id",
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
				"name": "identifier",
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

// CreateProjectAvroSchema creates the avro schema for Project
func CreateProjectAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(CreateProjectAvroSchemaSpec())
}

// TransformProjectFunc is a function for transforming Project during processing
type TransformProjectFunc func(input *Project) (*Project, error)

// CreateProjectPipe creates a pipe for processing Project items
func CreateProjectPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformProjectFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := CreateProjectInputStream(input, errors)
	var stream chan Project
	if len(transforms) > 0 {
		stream = make(chan Project, 1000)
	} else {
		stream = inch
	}
	outdone := CreateProjectOutputStream(output, stream, errors)
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

// CreateProjectInputStreamDir creates a channel for reading Project as JSON newlines from a directory of files
func CreateProjectInputStreamDir(dir string, errors chan<- error, transforms ...TransformProjectFunc) (chan Project, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/work/project\\.json(\\.gz)?$"))
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
		return CreateProjectInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan Project)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// CreateProjectInputStreamFile creates an channel for reading Project as JSON newlines from filename
func CreateProjectInputStreamFile(filename string, errors chan<- error, transforms ...TransformProjectFunc) (chan Project, <-chan bool) {
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
	return CreateProjectInputStream(f, errors, transforms...)
}

// CreateProjectInputStream creates an channel for reading Project as JSON newlines from stream
func CreateProjectInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformProjectFunc) (chan Project, <-chan bool) {
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

// CreateProjectOutputStreamDir will output json newlines from channel and save in dir
func CreateProjectOutputStreamDir(dir string, ch chan Project, errors chan<- error, transforms ...TransformProjectFunc) <-chan bool {
	fp := filepath.Join(dir, "/work/project\\.json(\\.gz)?$")
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
	return CreateProjectOutputStream(gz, ch, errors, transforms...)
}

// CreateProjectOutputStream will output json newlines from channel to the stream
func CreateProjectOutputStream(stream io.WriteCloser, ch chan Project, errors chan<- error, transforms ...TransformProjectFunc) <-chan bool {
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

// CreateProjectProducer will stream data from the channel
func CreateProjectProducer(producer util.Producer, ch chan Project, errors chan<- error) <-chan bool {
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

// CreateProjectConsumer will stream data from the default topic into the provided channel
func CreateProjectConsumer(factory util.ConsumerFactory, topic string, ch chan Project, errors chan<- error) (<-chan bool, chan<- bool) {
	return CreateProjectConsumerForTopic(factory, ProjectDefaultTopic, ch, errors)
}

// CreateProjectConsumerForTopic will stream data from the topic into the provided channel
func CreateProjectConsumerForTopic(factory util.ConsumerFactory, topic string, ch chan Project, errors chan<- error) (<-chan bool, chan<- bool) {
	done := make(chan bool, 1)
	closed := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		callback := util.ConsumerCallback{
			OnDataReceived: func(key []byte, value []byte) error {
				var object Project
				if err := json.Unmarshal(value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into Project: %s", err)
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
