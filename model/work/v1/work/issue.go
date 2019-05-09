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

// IssueDefaultTopic is the default topic name
const IssueDefaultTopic = "work_Issue_topic"

// IssueDefaultStream is the default stream name
const IssueDefaultStream = "work_Issue_stream"

// IssueDefaultTable is the default table name
const IssueDefaultTable = "work_Issue"

// Issue the issue is a specific work item for a project
type Issue struct {
	// built in types

	ID         string `json:"id" yaml:"id"`
	RefID      string `json:"ref_id" yaml:"ref_id"`
	RefType    string `json:"ref_type" yaml:"ref_type"`
	CustomerID string `json:"customer_id" yaml:"customer_id"`
	Hashcode   string `json:"hashcode" yaml:"hashcode"`

	// custom types

	Title          string   `json:"title" yaml:"title"`
	Identifier     string   `json:"identifier" yaml:"identifier"`
	ProjectID      string   `json:"project_id" yaml:"project_id"`
	URL            string   `json:"url" yaml:"url"`
	CreatedAt      int64    `json:"created_ts" yaml:"created_ts"`
	UpdatedAt      int64    `json:"updated_ts" yaml:"updated_ts"`
	PlannedStartAt int64    `json:"planned_start_ts" yaml:"planned_start_ts"`
	PlannedEndAt   int64    `json:"planned_end_ts" yaml:"planned_end_ts"`
	DueDateAt      int64    `json:"due_date_ts" yaml:"due_date_ts"`
	Priority       string   `json:"priority" yaml:"priority"`
	Type           string   `json:"type" yaml:"type"`
	Status         string   `json:"status" yaml:"status"`
	CreatorRefID   string   `json:"creator_ref_id" yaml:"creator_ref_id"`
	ReporterRefID  string   `json:"reporter_ref_id" yaml:"reporter_ref_id"`
	AssigneeRefID  string   `json:"assignee_ref_id" yaml:"assignee_ref_id"`
	AuthorRefID    string   `json:"author_ref_id" yaml:"author_ref_id"`
	Tags           []string `json:"tags" yaml:"tags"`
	ParentID       string   `json:"parent_id" yaml:"parent_id"`
	Resolution     string   `json:"resolution" yaml:"resolution"`
}

// String returns a string representation of Issue
func (o *Issue) String() string {
	return fmt.Sprintf("work.v1.Issue<%s>", o.ID)
}

func (o *Issue) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Issue) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Issue", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetRefID returns the RefID for the object
func (o *Issue) GetRefID() string {
	return o.RefID
}

// MarshalJSON returns the bytes for marshaling to json
func (o *Issue) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Issue) UnmarshalJSON(data []byte) error {
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
func (o *Issue) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Issue objects are equal
func (o *Issue) IsEqual(other *Issue) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Issue) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":               o.GetID(),
		"ref_id":           o.GetRefID(),
		"ref_type":         o.RefType,
		"customer_id":      o.CustomerID,
		"hashcode":         o.Hash(),
		"title":            o.Title,
		"identifier":       o.Identifier,
		"project_id":       o.ProjectID,
		"url":              o.URL,
		"created_ts":       o.CreatedAt,
		"updated_ts":       o.UpdatedAt,
		"planned_start_ts": o.PlannedStartAt,
		"planned_end_ts":   o.PlannedEndAt,
		"due_date_ts":      o.DueDateAt,
		"priority":         o.Priority,
		"type":             o.Type,
		"status":           o.Status,
		"creator_ref_id":   o.CreatorRefID,
		"reporter_ref_id":  o.ReporterRefID,
		"assignee_ref_id":  o.AssigneeRefID,
		"author_ref_id":    o.AuthorRefID,
		"tags":             o.Tags,
		"parent_id":        o.ParentID,
		"resolution":       o.Resolution,
	}
}

// FromMap attempts to load data into object from a map
func (o *Issue) FromMap(kv map[string]interface{}) {
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
	if val, ok := kv["title"].(string); ok {
		o.Title = val
	} else {
		val := kv["title"]
		if val == nil {
			o.Title = ""
		} else {
			o.Title = fmt.Sprintf("%v", val)
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
	if val, ok := kv["project_id"].(string); ok {
		o.ProjectID = val
	} else {
		val := kv["project_id"]
		if val == nil {
			o.ProjectID = ""
		} else {
			o.ProjectID = fmt.Sprintf("%v", val)
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
	if val, ok := kv["updated_ts"].(int64); ok {
		o.UpdatedAt = val
	} else {
		val := kv["updated_ts"]
		if val == nil {
			o.UpdatedAt = number.ToInt64Any(nil)
		} else {
			o.UpdatedAt = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["planned_start_ts"].(int64); ok {
		o.PlannedStartAt = val
	} else {
		val := kv["planned_start_ts"]
		if val == nil {
			o.PlannedStartAt = number.ToInt64Any(nil)
		} else {
			o.PlannedStartAt = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["planned_end_ts"].(int64); ok {
		o.PlannedEndAt = val
	} else {
		val := kv["planned_end_ts"]
		if val == nil {
			o.PlannedEndAt = number.ToInt64Any(nil)
		} else {
			o.PlannedEndAt = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["due_date_ts"].(int64); ok {
		o.DueDateAt = val
	} else {
		val := kv["due_date_ts"]
		if val == nil {
			o.DueDateAt = number.ToInt64Any(nil)
		} else {
			o.DueDateAt = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["priority"].(string); ok {
		o.Priority = val
	} else {
		val := kv["priority"]
		if val == nil {
			o.Priority = ""
		} else {
			o.Priority = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["type"].(string); ok {
		o.Type = val
	} else {
		val := kv["type"]
		if val == nil {
			o.Type = ""
		} else {
			o.Type = fmt.Sprintf("%v", val)
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
	if val, ok := kv["creator_ref_id"].(string); ok {
		o.CreatorRefID = val
	} else {
		val := kv["creator_ref_id"]
		if val == nil {
			o.CreatorRefID = ""
		} else {
			o.CreatorRefID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["reporter_ref_id"].(string); ok {
		o.ReporterRefID = val
	} else {
		val := kv["reporter_ref_id"]
		if val == nil {
			o.ReporterRefID = ""
		} else {
			o.ReporterRefID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["assignee_ref_id"].(string); ok {
		o.AssigneeRefID = val
	} else {
		val := kv["assignee_ref_id"]
		if val == nil {
			o.AssigneeRefID = ""
		} else {
			o.AssigneeRefID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["author_ref_id"].(string); ok {
		o.AuthorRefID = val
	} else {
		val := kv["author_ref_id"]
		if val == nil {
			o.AuthorRefID = ""
		} else {
			o.AuthorRefID = fmt.Sprintf("%v", val)
		}
	}
	val := kv["tags"]
	if val == nil {
		o.Tags = []string{}
	} else {
		o.Tags = append(o.Tags, fmt.Sprintf("%v", val))
	}
	if val, ok := kv["parent_id"].(string); ok {
		o.ParentID = val
	} else {
		val := kv["parent_id"]
		if val == nil {
			o.ParentID = ""
		} else {
			o.ParentID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["resolution"].(string); ok {
		o.Resolution = val
	} else {
		val := kv["resolution"]
		if val == nil {
			o.Resolution = ""
		} else {
			o.Resolution = fmt.Sprintf("%v", val)
		}
	}
	// make sure that these have values if empty
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *Issue) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.GetID())
	args = append(args, o.GetRefID())
	args = append(args, o.RefType)
	args = append(args, o.Title)
	args = append(args, o.Identifier)
	args = append(args, o.ProjectID)
	args = append(args, o.URL)
	args = append(args, o.CreatedAt)
	args = append(args, o.UpdatedAt)
	args = append(args, o.PlannedStartAt)
	args = append(args, o.PlannedEndAt)
	args = append(args, o.DueDateAt)
	args = append(args, o.Priority)
	args = append(args, o.Type)
	args = append(args, o.Status)
	args = append(args, o.CreatorRefID)
	args = append(args, o.ReporterRefID)
	args = append(args, o.AssigneeRefID)
	args = append(args, o.AuthorRefID)
	args = append(args, o.Tags)
	args = append(args, o.ParentID)
	args = append(args, o.Resolution)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// CreateIssueAvroSchemaSpec creates the avro schema specification for Issue
func CreateIssueAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "work.v1",
		"name":         "Issue",
		"connect.name": "work.v1.Issue",
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
				"name": "title",
				"type": "string",
			},
			map[string]interface{}{
				"name": "identifier",
				"type": "string",
			},
			map[string]interface{}{
				"name": "project_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "url",
				"type": "string",
			},
			map[string]interface{}{
				"name": "created_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "updated_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "planned_start_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "planned_end_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "due_date_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "priority",
				"type": "string",
			},
			map[string]interface{}{
				"name": "type",
				"type": "string",
			},
			map[string]interface{}{
				"name": "status",
				"type": "string",
			},
			map[string]interface{}{
				"name": "creator_ref_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "reporter_ref_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "assignee_ref_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "author_ref_id",
				"type": "string",
			},
			map[string]interface{}{
				"name":  "tags",
				"type":  "list",
				"items": "string",
			},
			map[string]interface{}{
				"name": "parent_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "resolution",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// CreateIssueAvroSchema creates the avro schema for Issue
func CreateIssueAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(CreateIssueAvroSchemaSpec())
}

// CreateIssueKQLStreamSQL creates KQL Stream SQL for Issue
func CreateIssueKQLStreamSQL() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("CREATE STREAM %s ", IssueDefaultStream))
	builder.WriteString(fmt.Sprintf("WITH (KAFKA_TOPIC='%s', VALUE_FORMAT='AVRO', KEY='id'", IssueDefaultTopic))
	builder.WriteString(", TIMESTAMP='updated_ts'")
	builder.WriteString(");")
	return builder.String()
}

// CreateIssueKQLTableSQL creates KQL Table SQL for Issue
func CreateIssueKQLTableSQL() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("CREATE TABLE %s ", IssueDefaultTable))
	builder.WriteString(fmt.Sprintf("WITH (KAFKA_TOPIC='%s', VALUE_FORMAT='AVRO', KEY='id'", IssueDefaultTopic))
	builder.WriteString(", TIMESTAMP='updated_ts'")
	builder.WriteString(");")
	return builder.String()
}

// TransformIssueFunc is a function for transforming Issue during processing
type TransformIssueFunc func(input *Issue) (*Issue, error)

// CreateIssuePipe creates a pipe for processing Issue items
func CreateIssuePipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformIssueFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := CreateIssueInputStream(input, errors)
	var stream chan Issue
	if len(transforms) > 0 {
		stream = make(chan Issue, 1000)
	} else {
		stream = inch
	}
	outdone := CreateIssueOutputStream(output, stream, errors)
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

// CreateIssueInputStreamDir creates a channel for reading Issue as JSON newlines from a directory of files
func CreateIssueInputStreamDir(dir string, errors chan<- error, transforms ...TransformIssueFunc) (chan Issue, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/work/v1/issue\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan Issue)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for issue")
		ch := make(chan Issue)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return CreateIssueInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan Issue)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// CreateIssueInputStreamFile creates an channel for reading Issue as JSON newlines from filename
func CreateIssueInputStreamFile(filename string, errors chan<- error, transforms ...TransformIssueFunc) (chan Issue, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan Issue)
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
			ch := make(chan Issue)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return CreateIssueInputStream(f, errors, transforms...)
}

// CreateIssueInputStream creates an channel for reading Issue as JSON newlines from stream
func CreateIssueInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformIssueFunc) (chan Issue, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan Issue, 1000)
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
			var item Issue
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

// CreateIssueOutputStreamDir will output json newlines from channel and save in dir
func CreateIssueOutputStreamDir(dir string, ch chan Issue, errors chan<- error, transforms ...TransformIssueFunc) <-chan bool {
	fp := filepath.Join(dir, "/work/v1/issue\\.json(\\.gz)?$")
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
	return CreateIssueOutputStream(gz, ch, errors, transforms...)
}

// CreateIssueOutputStream will output json newlines from channel to the stream
func CreateIssueOutputStream(stream io.WriteCloser, ch chan Issue, errors chan<- error, transforms ...TransformIssueFunc) <-chan bool {
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

// CreateIssueProducer will stream data from the channel
func CreateIssueProducer(producer util.Producer, ch chan Issue, errors chan<- error) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		schemaspec := CreateIssueAvroSchemaSpec()
		ctx := context.Background()
		for item := range ch {
			if err := producer.Send(ctx, schemaspec, []byte(item.ID), []byte(item.Stringify())); err != nil {
				errors <- fmt.Errorf("error sending %s. %v", item.String(), err)
			}
		}
	}()
	return done
}

// CreateIssueConsumer will stream data from the default topic into the provided channel
func CreateIssueConsumer(factory util.ConsumerFactory, topic string, ch chan Issue, errors chan<- error) (<-chan bool, chan<- bool) {
	return CreateIssueConsumerForTopic(factory, IssueDefaultTopic, ch, errors)
}

// CreateIssueConsumerForTopic will stream data from the topic into the provided channel
func CreateIssueConsumerForTopic(factory util.ConsumerFactory, topic string, ch chan Issue, errors chan<- error) (<-chan bool, chan<- bool) {
	done := make(chan bool, 1)
	closed := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		callback := util.ConsumerCallback{
			OnDataReceived: func(key []byte, value []byte) error {
				var object Issue
				if err := json.Unmarshal(value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into Issue: %s", err)
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
