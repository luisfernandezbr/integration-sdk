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

// ChangelogDefaultTopic is the default topic name
const ChangelogDefaultTopic = "work_Changelog_topic"

// ChangelogDefaultStream is the default stream name
const ChangelogDefaultStream = "work_Changelog_stream"

// ChangelogDefaultTable is the default table name
const ChangelogDefaultTable = "work_Changelog"

// Changelog change log
type Changelog struct {
	// built in types

	ID         string `json:"id" yaml:"id"`
	RefID      string `json:"ref_id" yaml:"ref_id"`
	RefType    string `json:"ref_type" yaml:"ref_type"`
	CustomerID string `json:"customer_id" yaml:"customer_id"`
	Hashcode   string `json:"hashcode" yaml:"hashcode"`

	// custom types

	IssueID    string `json:"issue_id" yaml:"issue_id"`
	CreatedAt  int64  `json:"created_ts" yaml:"created_ts"`
	Ordinal    int64  `json:"ordinal" yaml:"ordinal"`
	UserID     string `json:"user_id" yaml:"user_id"`
	UserActive bool   `json:"user_active" yaml:"user_active"`
	Field      string `json:"field" yaml:"field"`
	FieldType  string `json:"field_type" yaml:"field_type"`
	From       string `json:"from" yaml:"from"`
	FromString string `json:"from_string" yaml:"from_string"`
	To         string `json:"to" yaml:"to"`
	ToString   string `json:"to_string" yaml:"to_string"`
}

// String returns a string representation of Changelog
func (o *Changelog) String() string {
	return fmt.Sprintf("work.v1.Changelog<%s>", o.ID)
}

func (o *Changelog) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Changelog) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Changelog", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetRefID returns the RefID for the object
func (o *Changelog) GetRefID() string {
	return o.RefID
}

// MarshalJSON returns the bytes for marshaling to json
func (o *Changelog) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Changelog) UnmarshalJSON(data []byte) error {
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
func (o *Changelog) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Changelog objects are equal
func (o *Changelog) IsEqual(other *Changelog) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Changelog) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":          o.GetID(),
		"ref_id":      o.GetRefID(),
		"ref_type":    o.RefType,
		"customer_id": o.CustomerID,
		"hashcode":    o.Hash(),
		"issue_id":    o.IssueID,
		"created_ts":  o.CreatedAt,
		"ordinal":     o.Ordinal,
		"user_id":     o.UserID,
		"user_active": o.UserActive,
		"field":       o.Field,
		"field_type":  o.FieldType,
		"from":        o.From,
		"from_string": o.FromString,
		"to":          o.To,
		"to_string":   o.ToString,
	}
}

// FromMap attempts to load data into object from a map
func (o *Changelog) FromMap(kv map[string]interface{}) {
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
	if val, ok := kv["issue_id"].(string); ok {
		o.IssueID = val
	} else {
		val := kv["issue_id"]
		if val == nil {
			o.IssueID = ""
		} else {
			o.IssueID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["created_ts"].(int64); ok {
		o.CreatedAt = val
	} else {
		val := kv["created_ts"]
		if val == nil {
			o.CreatedAt = number.ToInt64Any(nil)
		} else {
			o.CreatedAt = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["ordinal"].(int64); ok {
		o.Ordinal = val
	} else {
		val := kv["ordinal"]
		if val == nil {
			o.Ordinal = number.ToInt64Any(nil)
		} else {
			o.Ordinal = number.ToInt64Any(val)
		}
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
	if val, ok := kv["user_active"].(bool); ok {
		o.UserActive = val
	} else {
		val := kv["user_active"]
		if val == nil {
			o.UserActive = number.ToBoolAny(nil)
		} else {
			o.UserActive = number.ToBoolAny(val)
		}
	}
	if val, ok := kv["field"].(string); ok {
		o.Field = val
	} else {
		val := kv["field"]
		if val == nil {
			o.Field = ""
		} else {
			o.Field = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["field_type"].(string); ok {
		o.FieldType = val
	} else {
		val := kv["field_type"]
		if val == nil {
			o.FieldType = ""
		} else {
			o.FieldType = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["from"].(string); ok {
		o.From = val
	} else {
		val := kv["from"]
		if val == nil {
			o.From = ""
		} else {
			o.From = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["from_string"].(string); ok {
		o.FromString = val
	} else {
		val := kv["from_string"]
		if val == nil {
			o.FromString = ""
		} else {
			o.FromString = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["to"].(string); ok {
		o.To = val
	} else {
		val := kv["to"]
		if val == nil {
			o.To = ""
		} else {
			o.To = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["to_string"].(string); ok {
		o.ToString = val
	} else {
		val := kv["to_string"]
		if val == nil {
			o.ToString = ""
		} else {
			o.ToString = fmt.Sprintf("%v", val)
		}
	}
	// make sure that these have values if empty
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *Changelog) Hash() string {
	if o.Hashcode == "" {
		args := make([]interface{}, 0)
		args = append(args, o.GetID())
		args = append(args, o.GetRefID())
		args = append(args, o.RefType)
		args = append(args, o.IssueID)
		args = append(args, o.CreatedAt)
		args = append(args, o.Ordinal)
		args = append(args, o.UserID)
		args = append(args, o.UserActive)
		args = append(args, o.Field)
		args = append(args, o.FieldType)
		args = append(args, o.From)
		args = append(args, o.FromString)
		args = append(args, o.To)
		args = append(args, o.ToString)
		o.Hashcode = hash.Values(args...)
	}
	return o.Hashcode
}

// CreateChangelogAvroSchemaSpec creates the avro schema specification for Changelog
func CreateChangelogAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "work.v1",
		"name":         "Changelog",
		"connect.name": "work.v1.Changelog",
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
				"name": "issue_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "created_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "ordinal",
				"type": "long",
			},
			map[string]interface{}{
				"name": "user_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "user_active",
				"type": "boolean",
			},
			map[string]interface{}{
				"name": "field",
				"type": "string",
			},
			map[string]interface{}{
				"name": "field_type",
				"type": "string",
			},
			map[string]interface{}{
				"name": "from",
				"type": "string",
			},
			map[string]interface{}{
				"name": "from_string",
				"type": "string",
			},
			map[string]interface{}{
				"name": "to",
				"type": "string",
			},
			map[string]interface{}{
				"name": "to_string",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// CreateChangelogAvroSchema creates the avro schema for Changelog
func CreateChangelogAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(CreateChangelogAvroSchemaSpec())
}

// CreateChangelogKQLStreamSQL creates KQL Stream SQL for Changelog
func CreateChangelogKQLStreamSQL() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("CREATE STREAM %s ", ChangelogDefaultStream))
	builder.WriteString(fmt.Sprintf("WITH (KAFKA_TOPIC='%s', VALUE_FORMAT='AVRO', KEY='id'", ChangelogDefaultTopic))
	builder.WriteString(", TIMESTAMP='created_ts'")
	builder.WriteString(");")
	return builder.String()
}

// CreateChangelogKQLTableSQL creates KQL Table SQL for Changelog
func CreateChangelogKQLTableSQL() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("CREATE TABLE %s ", ChangelogDefaultTable))
	builder.WriteString(fmt.Sprintf("WITH (KAFKA_TOPIC='%s', VALUE_FORMAT='AVRO', KEY='id'", ChangelogDefaultTopic))
	builder.WriteString(", TIMESTAMP='created_ts'")
	builder.WriteString(");")
	return builder.String()
}

// TransformChangelogFunc is a function for transforming Changelog during processing
type TransformChangelogFunc func(input *Changelog) (*Changelog, error)

// CreateChangelogPipe creates a pipe for processing Changelog items
func CreateChangelogPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformChangelogFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := CreateChangelogInputStream(input, errors)
	var stream chan Changelog
	if len(transforms) > 0 {
		stream = make(chan Changelog, 1000)
	} else {
		stream = inch
	}
	outdone := CreateChangelogOutputStream(output, stream, errors)
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

// CreateChangelogInputStreamDir creates a channel for reading Changelog as JSON newlines from a directory of files
func CreateChangelogInputStreamDir(dir string, errors chan<- error, transforms ...TransformChangelogFunc) (chan Changelog, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/work/v1/changelog\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan Changelog)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for changelog")
		ch := make(chan Changelog)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return CreateChangelogInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan Changelog)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// CreateChangelogInputStreamFile creates an channel for reading Changelog as JSON newlines from filename
func CreateChangelogInputStreamFile(filename string, errors chan<- error, transforms ...TransformChangelogFunc) (chan Changelog, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan Changelog)
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
			ch := make(chan Changelog)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return CreateChangelogInputStream(f, errors, transforms...)
}

// CreateChangelogInputStream creates an channel for reading Changelog as JSON newlines from stream
func CreateChangelogInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformChangelogFunc) (chan Changelog, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan Changelog, 1000)
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
			var item Changelog
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

// CreateChangelogOutputStreamDir will output json newlines from channel and save in dir
func CreateChangelogOutputStreamDir(dir string, ch chan Changelog, errors chan<- error, transforms ...TransformChangelogFunc) <-chan bool {
	fp := filepath.Join(dir, "/work/v1/changelog\\.json(\\.gz)?$")
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
	return CreateChangelogOutputStream(gz, ch, errors, transforms...)
}

// CreateChangelogOutputStream will output json newlines from channel to the stream
func CreateChangelogOutputStream(stream io.WriteCloser, ch chan Changelog, errors chan<- error, transforms ...TransformChangelogFunc) <-chan bool {
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

// CreateChangelogProducer will stream data from the channel
func CreateChangelogProducer(producer util.Producer, ch chan Changelog, errors chan<- error) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		schemaspec := CreateChangelogAvroSchemaSpec()
		ctx := context.Background()
		for item := range ch {
			if err := producer.Send(ctx, schemaspec, []byte(item.ID), []byte(item.Stringify())); err != nil {
				errors <- fmt.Errorf("error sending %s. %v", item.String(), err)
			}
		}
	}()
	return done
}

// CreateChangelogConsumer will stream data from the default topic into the provided channel
func CreateChangelogConsumer(factory util.ConsumerFactory, topic string, ch chan Changelog, errors chan<- error) (<-chan bool, chan<- bool) {
	return CreateChangelogConsumerForTopic(factory, ChangelogDefaultTopic, ch, errors)
}

// CreateChangelogConsumerForTopic will stream data from the topic into the provided channel
func CreateChangelogConsumerForTopic(factory util.ConsumerFactory, topic string, ch chan Changelog, errors chan<- error) (<-chan bool, chan<- bool) {
	done := make(chan bool, 1)
	closed := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		callback := util.ConsumerCallback{
			OnDataReceived: func(key []byte, value []byte) error {
				var object Changelog
				if err := json.Unmarshal(value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into Changelog: %s", err)
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
