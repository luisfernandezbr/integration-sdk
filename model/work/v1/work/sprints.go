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
	"github.com/pinpt/go-common/number"
	"github.com/pinpt/integration-sdk/util"
)

// SprintsDefaultTopic is the default topic name
const SprintsDefaultTopic = "work_Sprints_topic"

// SprintsDefaultStream is the default stream name
const SprintsDefaultStream = "work_Sprints_stream"

// SprintsDefaultTable is the default table name
const SprintsDefaultTable = "work_Sprints"

// Sprints sprint details
type Sprints struct {
	// built in types

	ID         string `json:"id" yaml:"id"`
	RefID      string `json:"ref_id" yaml:"ref_id"`
	RefType    string `json:"ref_type" yaml:"ref_type"`
	CustomerID string `json:"customer_id" yaml:"customer_id"`
	Hashcode   string `json:"hashcode" yaml:"hashcode"`

	// custom types

	Name        string `json:"name" yaml:"name"`
	Identifier  string `json:"identifier" yaml:"identifier"`
	Status      string `json:"status" yaml:"status"`
	StartedAt   int64  `json:"started_ts" yaml:"started_ts"`
	EndedAt     int64  `json:"ended_ts" yaml:"ended_ts"`
	CompletedAt int64  `json:"completed_ts" yaml:"completed_ts"`
}

// String returns a string representation of Sprints
func (o *Sprints) String() string {
	return fmt.Sprintf("work.v1.Sprints<%s>", o.ID)
}

func (o *Sprints) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Sprints) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Sprints", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetRefID returns the RefID for the object
func (o *Sprints) GetRefID() string {
	return o.RefID
}

// MarshalJSON returns the bytes for marshaling to json
func (o *Sprints) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Sprints) UnmarshalJSON(data []byte) error {
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
func (o *Sprints) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Sprints objects are equal
func (o *Sprints) IsEqual(other *Sprints) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Sprints) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":           o.GetID(),
		"ref_id":       o.GetRefID(),
		"ref_type":     o.RefType,
		"customer_id":  o.CustomerID,
		"hashcode":     o.Hash(),
		"name":         o.Name,
		"identifier":   o.Identifier,
		"status":       o.Status,
		"started_ts":   o.StartedAt,
		"ended_ts":     o.EndedAt,
		"completed_ts": o.CompletedAt,
	}
}

// FromMap attempts to load data into object from a map
func (o *Sprints) FromMap(kv map[string]interface{}) {
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
	if val, ok := kv["status"].(string); ok {
		o.Status = val
	} else {
		val := kv["status"]
		if val == nil {
			o.Status = ""
		} else {
			o.Status = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["started_ts"].(int64); ok {
		o.StartedAt = val
	} else {
		val := kv["started_ts"]
		if val == nil {
			o.StartedAt = number.ToInt64Any(nil)
		} else {
			o.StartedAt = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["ended_ts"].(int64); ok {
		o.EndedAt = val
	} else {
		val := kv["ended_ts"]
		if val == nil {
			o.EndedAt = number.ToInt64Any(nil)
		} else {
			o.EndedAt = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["completed_ts"].(int64); ok {
		o.CompletedAt = val
	} else {
		val := kv["completed_ts"]
		if val == nil {
			o.CompletedAt = number.ToInt64Any(nil)
		} else {
			o.CompletedAt = number.ToInt64Any(val)
		}
	}
	// make sure that these have values if empty
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *Sprints) Hash() string {
	if o.Hashcode == "" {
		args := make([]interface{}, 0)
		args = append(args, o.GetID())
		args = append(args, o.GetRefID())
		args = append(args, o.RefType)
		args = append(args, o.Name)
		args = append(args, o.Identifier)
		args = append(args, o.Status)
		args = append(args, o.StartedAt)
		args = append(args, o.EndedAt)
		args = append(args, o.CompletedAt)
		o.Hashcode = hash.Values(args...)
	}
	return o.Hashcode
}

// CreateSprintsAvroSchemaSpec creates the avro schema specification for Sprints
func CreateSprintsAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "work.v1",
		"name":         "Sprints",
		"connect.name": "work.v1.Sprints",
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
				"name": "identifier",
				"type": "string",
			},
			map[string]interface{}{
				"name": "status",
				"type": "string",
			},
			map[string]interface{}{
				"name": "started_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "ended_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "completed_ts",
				"type": "long",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// CreateSprintsAvroSchema creates the avro schema for Sprints
func CreateSprintsAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(CreateSprintsAvroSchemaSpec())
}

// CreateSprintsKQLStreamSQL creates KQL Stream SQL for Sprints
func CreateSprintsKQLStreamSQL() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("CREATE STREAM %s ", SprintsDefaultStream))
	builder.WriteString(fmt.Sprintf("WITH (KAFKA_TOPIC='%s', VALUE_FORMAT='AVRO', KEY='id'", SprintsDefaultTopic))
	builder.WriteString(", TIMESTAMP='started_ts'")
	builder.WriteString(");")
	return builder.String()
}

// CreateSprintsKQLTableSQL creates KQL Table SQL for Sprints
func CreateSprintsKQLTableSQL() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("CREATE TABLE %s ", SprintsDefaultTable))
	builder.WriteString(fmt.Sprintf("WITH (KAFKA_TOPIC='%s', VALUE_FORMAT='AVRO', KEY='id'", SprintsDefaultTopic))
	builder.WriteString(", TIMESTAMP='started_ts'")
	builder.WriteString(");")
	return builder.String()
}

// TransformSprintsFunc is a function for transforming Sprints during processing
type TransformSprintsFunc func(input *Sprints) (*Sprints, error)

// CreateSprintsPipe creates a pipe for processing Sprints items
func CreateSprintsPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformSprintsFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := CreateSprintsInputStream(input, errors)
	var stream chan Sprints
	if len(transforms) > 0 {
		stream = make(chan Sprints, 1000)
	} else {
		stream = inch
	}
	outdone := CreateSprintsOutputStream(output, stream, errors)
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

// CreateSprintsInputStreamDir creates a channel for reading Sprints as JSON newlines from a directory of files
func CreateSprintsInputStreamDir(dir string, errors chan<- error, transforms ...TransformSprintsFunc) (chan Sprints, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/work/v1/sprints\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan Sprints)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for sprints")
		ch := make(chan Sprints)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return CreateSprintsInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan Sprints)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// CreateSprintsInputStreamFile creates an channel for reading Sprints as JSON newlines from filename
func CreateSprintsInputStreamFile(filename string, errors chan<- error, transforms ...TransformSprintsFunc) (chan Sprints, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan Sprints)
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
			ch := make(chan Sprints)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return CreateSprintsInputStream(f, errors, transforms...)
}

// CreateSprintsInputStream creates an channel for reading Sprints as JSON newlines from stream
func CreateSprintsInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformSprintsFunc) (chan Sprints, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan Sprints, 1000)
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
			var item Sprints
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

// CreateSprintsOutputStreamDir will output json newlines from channel and save in dir
func CreateSprintsOutputStreamDir(dir string, ch chan Sprints, errors chan<- error, transforms ...TransformSprintsFunc) <-chan bool {
	fp := filepath.Join(dir, "/work/v1/sprints\\.json(\\.gz)?$")
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
	return CreateSprintsOutputStream(gz, ch, errors, transforms...)
}

// CreateSprintsOutputStream will output json newlines from channel to the stream
func CreateSprintsOutputStream(stream io.WriteCloser, ch chan Sprints, errors chan<- error, transforms ...TransformSprintsFunc) <-chan bool {
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

// CreateSprintsProducer will stream data from the channel
func CreateSprintsProducer(producer util.Producer, ch chan Sprints, errors chan<- error) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		schemaspec := CreateSprintsAvroSchemaSpec()
		ctx := context.Background()
		for item := range ch {
			if err := producer.Send(ctx, schemaspec, []byte(item.ID), []byte(item.Stringify())); err != nil {
				errors <- fmt.Errorf("error sending %s. %v", item.String(), err)
			}
		}
	}()
	return done
}

// CreateSprintsConsumer will stream data from the default topic into the provided channel
func CreateSprintsConsumer(factory util.ConsumerFactory, topic string, ch chan Sprints, errors chan<- error) (<-chan bool, chan<- bool) {
	return CreateSprintsConsumerForTopic(factory, SprintsDefaultTopic, ch, errors)
}

// CreateSprintsConsumerForTopic will stream data from the topic into the provided channel
func CreateSprintsConsumerForTopic(factory util.ConsumerFactory, topic string, ch chan Sprints, errors chan<- error) (<-chan bool, chan<- bool) {
	done := make(chan bool, 1)
	closed := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		callback := util.ConsumerCallback{
			OnDataReceived: func(key []byte, value []byte) error {
				var object Sprints
				if err := json.Unmarshal(value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into Sprints: %s", err)
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
