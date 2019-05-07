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

	cluster "github.com/bsm/sarama-cluster"
	kafka "github.com/dangkaka/go-kafka-avro"
	"github.com/linkedin/goavro"
	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
	"github.com/pinpt/integration-sdk/util"
)

// IssueDefaultTopic is the default topic name
const IssueDefaultTopic = "work.Issue.topic"

// IssueDefaultStream is the default stream name
const IssueDefaultStream = "work.Issue.topicstream"

// IssueDefaultTable is the default table name
const IssueDefaultTable = "work.Issue"

// Issue the issue is a specific work item for a project
type Issue struct {
	// built in types

	ID         string `json:"id" yaml:"id"`
	RefID      string `json:"ref_id" yaml:"ref_id"`
	RefType    string `json:"ref_type" yaml:"ref_type"`
	CustomerID string `json:"customer_id" yaml:"customer_id"`
	Hashcode   string `json:"hashcode" yaml:"hashcode"`

	// custom types

	Title             string `json:"title" yaml:"title"`
	Identifier        string `json:"identifier" yaml:"identifier"`
	URL               string `json:"url" yaml:"url"`
	CreatedAt         int64  `json:"created_ts" yaml:"created_ts"`
	RefCreatorUserID  string `json:"ref_creator_user_id" yaml:"ref_creator_user_id"`
	RefReporterUserID string `json:"ref_reporter_user_id" yaml:"ref_reporter_user_id"`
	RefAssigneeUserID string `json:"ref_assignee_user_id" yaml:"ref_assignee_user_id"`
	Key               string `json:"key" yaml:"key"`
	IssueType         string `json:"issue_type" yaml:"issue_type"`
	ProjectID         string `json:"project_id" yaml:"project_id"`
	Labels            string `json:"labels" yaml:"labels"`
	ParentID          string `json:"parent_id" yaml:"parent_id"`
	Status            string `json:"status" yaml:"status"`
	Resolution        string `json:"resolution" yaml:"resolution"`
	Duedate           string `json:"duedate" yaml:"duedate"`
	IssuetypeID       string `json:"issuetype_id" yaml:"issuetype_id"`
	Combo             string `json:"combo" yaml:"combo"`
	Comp              string `json:"comp" yaml:"comp"`
	Active            string `json:"active" yaml:"active"`
	Username          string `json:"username" yaml:"username"`
	Displayname       string `json:"displayname" yaml:"displayname"`
	Email             string `json:"email" yaml:"email"`
	AvatarURL         string `json:"avatar_url" yaml:"avatar_url"`
	PriorityID        string `json:"priority_id" yaml:"priority_id"`
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
		"id":                   o.GetID(),
		"ref_id":               o.GetRefID(),
		"ref_type":             o.RefType,
		"customer_id":          o.CustomerID,
		"hashcode":             o.Hash(),
		"title":                o.Title,
		"identifier":           o.Identifier,
		"url":                  o.URL,
		"created_ts":           o.CreatedAt,
		"ref_creator_user_id":  o.RefCreatorUserID,
		"ref_reporter_user_id": o.RefReporterUserID,
		"ref_assignee_user_id": o.RefAssigneeUserID,
		"key":                  o.Key,
		"issue_type":           o.IssueType,
		"project_id":           o.ProjectID,
		"labels":               o.Labels,
		"parent_id":            o.ParentID,
		"status":               o.Status,
		"resolution":           o.Resolution,
		"duedate":              o.Duedate,
		"issuetype_id":         o.IssuetypeID,
		"combo":                o.Combo,
		"comp":                 o.Comp,
		"active":               o.Active,
		"username":             o.Username,
		"displayname":          o.Displayname,
		"email":                o.Email,
		"avatar_url":           o.AvatarURL,
		"priority_id":          o.PriorityID,
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
		o.Title = fmt.Sprintf("%v", val)
	}
	if val, ok := kv["identifier"].(string); ok {
		o.Identifier = val
	} else {
		val := kv["identifier"]
		o.Identifier = fmt.Sprintf("%v", val)
	}
	if val, ok := kv["url"].(string); ok {
		o.URL = val
	} else {
		val := kv["url"]
		o.URL = fmt.Sprintf("%v", val)
	}
	if val, ok := kv["created_ts"].(int64); ok {
		o.CreatedAt = val
	} else {
		val := kv["created_ts"]
		o.CreatedAt = number.ToInt64Any(val)
	}
	if val, ok := kv["ref_creator_user_id"].(string); ok {
		o.RefCreatorUserID = val
	} else {
		val := kv["ref_creator_user_id"]
		o.RefCreatorUserID = fmt.Sprintf("%v", val)
	}
	if val, ok := kv["ref_reporter_user_id"].(string); ok {
		o.RefReporterUserID = val
	} else {
		val := kv["ref_reporter_user_id"]
		o.RefReporterUserID = fmt.Sprintf("%v", val)
	}
	if val, ok := kv["ref_assignee_user_id"].(string); ok {
		o.RefAssigneeUserID = val
	} else {
		val := kv["ref_assignee_user_id"]
		o.RefAssigneeUserID = fmt.Sprintf("%v", val)
	}
	if val, ok := kv["key"].(string); ok {
		o.Key = val
	} else {
		val := kv["key"]
		o.Key = fmt.Sprintf("%v", val)
	}
	if val, ok := kv["issue_type"].(string); ok {
		o.IssueType = val
	} else {
		val := kv["issue_type"]
		o.IssueType = fmt.Sprintf("%v", val)
	}
	if val, ok := kv["project_id"].(string); ok {
		o.ProjectID = val
	} else {
		val := kv["project_id"]
		o.ProjectID = fmt.Sprintf("%v", val)
	}
	if val, ok := kv["labels"].(string); ok {
		o.Labels = val
	} else {
		val := kv["labels"]
		o.Labels = fmt.Sprintf("%v", val)
	}
	if val, ok := kv["parent_id"].(string); ok {
		o.ParentID = val
	} else {
		val := kv["parent_id"]
		o.ParentID = fmt.Sprintf("%v", val)
	}
	if val, ok := kv["status"].(string); ok {
		o.Status = val
	} else {
		val := kv["status"]
		o.Status = fmt.Sprintf("%v", val)
	}
	if val, ok := kv["resolution"].(string); ok {
		o.Resolution = val
	} else {
		val := kv["resolution"]
		o.Resolution = fmt.Sprintf("%v", val)
	}
	if val, ok := kv["duedate"].(string); ok {
		o.Duedate = val
	} else {
		val := kv["duedate"]
		o.Duedate = fmt.Sprintf("%v", val)
	}
	if val, ok := kv["issuetype_id"].(string); ok {
		o.IssuetypeID = val
	} else {
		val := kv["issuetype_id"]
		o.IssuetypeID = fmt.Sprintf("%v", val)
	}
	if val, ok := kv["combo"].(string); ok {
		o.Combo = val
	} else {
		val := kv["combo"]
		o.Combo = fmt.Sprintf("%v", val)
	}
	if val, ok := kv["comp"].(string); ok {
		o.Comp = val
	} else {
		val := kv["comp"]
		o.Comp = fmt.Sprintf("%v", val)
	}
	if val, ok := kv["active"].(string); ok {
		o.Active = val
	} else {
		val := kv["active"]
		o.Active = fmt.Sprintf("%v", val)
	}
	if val, ok := kv["username"].(string); ok {
		o.Username = val
	} else {
		val := kv["username"]
		o.Username = fmt.Sprintf("%v", val)
	}
	if val, ok := kv["displayname"].(string); ok {
		o.Displayname = val
	} else {
		val := kv["displayname"]
		o.Displayname = fmt.Sprintf("%v", val)
	}
	if val, ok := kv["email"].(string); ok {
		o.Email = val
	} else {
		val := kv["email"]
		o.Email = fmt.Sprintf("%v", val)
	}
	if val, ok := kv["avatar_url"].(string); ok {
		o.AvatarURL = val
	} else {
		val := kv["avatar_url"]
		o.AvatarURL = fmt.Sprintf("%v", val)
	}
	if val, ok := kv["priority_id"].(string); ok {
		o.PriorityID = val
	} else {
		val := kv["priority_id"]
		o.PriorityID = fmt.Sprintf("%v", val)
	}
	// make sure that these have values if empty
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *Issue) Hash() string {
	if o.Hashcode == "" {
		args := make([]interface{}, 0)
		args = append(args, o.GetID())
		args = append(args, o.GetRefID())
		args = append(args, o.RefType)
		args = append(args, o.Title)
		args = append(args, o.Identifier)
		args = append(args, o.URL)
		args = append(args, o.CreatedAt)
		args = append(args, o.RefCreatorUserID)
		args = append(args, o.RefReporterUserID)
		args = append(args, o.RefAssigneeUserID)
		args = append(args, o.Key)
		args = append(args, o.IssueType)
		args = append(args, o.ProjectID)
		args = append(args, o.Labels)
		args = append(args, o.ParentID)
		args = append(args, o.Status)
		args = append(args, o.Resolution)
		args = append(args, o.Duedate)
		args = append(args, o.IssuetypeID)
		args = append(args, o.Combo)
		args = append(args, o.Comp)
		args = append(args, o.Active)
		args = append(args, o.Username)
		args = append(args, o.Displayname)
		args = append(args, o.Email)
		args = append(args, o.AvatarURL)
		args = append(args, o.PriorityID)
		o.Hashcode = hash.Values(args...)
	}
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
				"name": "url",
				"type": "string",
			},
			map[string]interface{}{
				"name": "created_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "ref_creator_user_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "ref_reporter_user_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "ref_assignee_user_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "key",
				"type": "string",
			},
			map[string]interface{}{
				"name": "issue_type",
				"type": "string",
			},
			map[string]interface{}{
				"name": "project_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "labels",
				"type": "string",
			},
			map[string]interface{}{
				"name": "parent_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "status",
				"type": "string",
			},
			map[string]interface{}{
				"name": "resolution",
				"type": "string",
			},
			map[string]interface{}{
				"name": "duedate",
				"type": "string",
			},
			map[string]interface{}{
				"name": "issuetype_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "combo",
				"type": "string",
			},
			map[string]interface{}{
				"name": "comp",
				"type": "string",
			},
			map[string]interface{}{
				"name": "active",
				"type": "string",
			},
			map[string]interface{}{
				"name": "username",
				"type": "string",
			},
			map[string]interface{}{
				"name": "displayname",
				"type": "string",
			},
			map[string]interface{}{
				"name": "email",
				"type": "string",
			},
			map[string]interface{}{
				"name": "avatar_url",
				"type": "string",
			},
			map[string]interface{}{
				"name": "priority_id",
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
	builder.WriteString(", TIMESTAMP='created_ts'")
	builder.WriteString(");")
	return builder.String()
}

// CreateIssueKQLTableSQL creates KQL Table SQL for Issue
func CreateIssueKQLTableSQL() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("CREATE TABLE %s ", IssueDefaultTable))
	builder.WriteString(fmt.Sprintf("WITH (KAFKA_TOPIC='%s', VALUE_FORMAT='AVRO', KEY='id'", IssueDefaultTopic))
	builder.WriteString(", TIMESTAMP='created_ts'")
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
func CreateIssueProducer(producer *util.KafkaProducer, ch chan Issue, errors chan<- error) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		schemaspec := CreateIssueAvroSchemaSpec()
		ctx := context.Background()
		for item := range ch {
			if err := producer.Send(ctx, schemaspec, []byte(item.ID), []byte(item.Stringify())); err != nil {
				errors <- fmt.Errorf("error sending %s. %v", item, err)
			}
		}
	}()
	return done
}

// CreateIssueConsumer will stream data from the default topic into the provided channel
func CreateIssueConsumer(kafkaServers []string, schemaRegistryServers []string, topic string, consumerGroupID string, ch chan Issue, errors chan<- error) (<-chan bool, chan<- bool) {
	return CreateIssueConsumerForTopic(kafkaServers, schemaRegistryServers, IssueDefaultTopic, consumerGroupID, ch, errors)
}

// CreateIssueConsumerForTopic will stream data from the topic into the provided channel
func CreateIssueConsumerForTopic(kafkaServers []string, schemaRegistryServers []string, topic string, consumerGroupID string, ch chan Issue, errors chan<- error) (<-chan bool, chan<- bool) {
	done := make(chan bool, 1)
	closed := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		consumerCallbacks := kafka.ConsumerCallbacks{
			OnDataReceived: func(msg kafka.Message) {
				var object Issue
				if err := json.Unmarshal([]byte(msg.Value), &object); err != nil {
					errors <- fmt.Errorf("error unmarshaling json data into Issue: %s", err)
					return
				}
				ch <- object
			},
			OnError: func(err error) {
				errors <- err
			},
			OnNotification: func(notification *cluster.Notification) {
			},
		}
		consumer, err := kafka.NewAvroConsumer(kafkaServers, schemaRegistryServers, topic, consumerGroupID, consumerCallbacks)
		if err != nil {
			errors <- err
			return
		}
		go consumer.Consume()
		select {
		case <-closed:
			consumer.Close()
		}
	}()
	return done, closed
}
