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
	"strings"

	"github.com/bxcodec/faker"
	"github.com/linkedin/goavro"
	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	number "github.com/pinpt/go-common/number"
	"github.com/pinpt/go-datamodel/datamodel"
)

// IssueTopic is the default topic name
const IssueTopic datamodel.TopicNameType = "work_Issue_topic"

// IssueStream is the default stream name
const IssueStream datamodel.TopicNameType = "work_Issue_stream"

// IssueTable is the default table name
const IssueTable datamodel.TopicNameType = "work_Issue"

// IssueModelName is the model name
const IssueModelName datamodel.ModelNameType = "work.Issue"

// Issue the issue is a specific work item for a project
type Issue struct {
	// built in types

	ID         string `json:"issue_id" yaml:"issue_id" faker:"-"`
	RefID      string `json:"ref_id" yaml:"ref_id" faker:"-"`
	RefType    string `json:"ref_type" yaml:"ref_type" faker:"-"`
	CustomerID string `json:"customer_id" yaml:"customer_id" faker:"-"`
	Hashcode   string `json:"hashcode" yaml:"hashcode" faker:"-"`
	// custom types

	// Title the issue title
	Title string `json:"title" yaml:"title" faker:"issue_title"`
	// Identifier the common identifier for the issue
	Identifier string `json:"identifier" yaml:"identifier" faker:"issue_id"`
	// ProjectID unique project id
	ProjectID string `json:"project_id" yaml:"project_id" faker:"-"`
	// URL the url to the issue page
	URL string `json:"url" yaml:"url" faker:"url"`
	// CreatedAt the timestamp in UTC that the issue was created
	CreatedAt int64 `json:"created_ts" yaml:"created_ts" faker:"-"`
	// UpdatedAt the timestamp in UTC that the issue was updated
	UpdatedAt int64 `json:"updated_ts" yaml:"updated_ts" faker:"-"`
	// PlannedStartAt the timestamp in UTC that the issue was planned to start
	PlannedStartAt int64 `json:"planned_start_ts" yaml:"planned_start_ts" faker:"-"`
	// PlannedEndAt the timestamp in UTC that the issue was planned to end
	PlannedEndAt int64 `json:"planned_end_ts" yaml:"planned_end_ts" faker:"-"`
	// DueDateAt due date of the issue
	DueDateAt int64 `json:"due_date_ts" yaml:"due_date_ts" faker:"-"`
	// Priority priority of the issue
	Priority string `json:"priority" yaml:"priority" faker:"-"`
	// Type type of issue
	Type string `json:"type" yaml:"type" faker:"-"`
	// Status status of the issue
	Status string `json:"status" yaml:"status" faker:"-"`
	// CreatorRefID user id of the creator
	CreatorRefID string `json:"creator_ref_id" yaml:"creator_ref_id" faker:"-"`
	// ReporterRefID user id of the reporter
	ReporterRefID string `json:"reporter_ref_id" yaml:"reporter_ref_id" faker:"-"`
	// AssigneeRefID user id of the assignee
	AssigneeRefID string `json:"assignee_ref_id" yaml:"assignee_ref_id" faker:"-"`
	// AuthorRefID user id of the author
	AuthorRefID string `json:"author_ref_id" yaml:"author_ref_id" faker:"-"`
	// Tags tags on the issue
	Tags []string `json:"tags" yaml:"tags" faker:"-"`
	// ParentID parent issue id, if any
	ParentID string `json:"parent_id" yaml:"parent_id" faker:"-"`
	// Resolution resolution of the issue
	Resolution string `json:"resolution" yaml:"resolution" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Issue)(nil)

func toIssueObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toIssueObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if o == nil {
		return toIssueObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toIssueObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toIssueObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toIssueObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toIssueObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toIssueObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toIssueObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toIssueObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toIssueObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toIssueObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toIssueObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case map[string]interface{}:
		return o
	case *map[string]interface{}:
		return v
	case *Issue:
		return v.ToMap()
	case Issue:
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
			arr = append(arr, toIssueObject(av, isavro, false, ""))
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of Issue
func (o *Issue) String() string {
	return fmt.Sprintf("work.Issue<%s>", o.ID)
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

// Clone returns an exact copy of Issue
func (o *Issue) Clone() *Issue {
	c := new(Issue)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Issue) Anon() *Issue {
	c := new(Issue)
	if err := faker.FakeData(c); err != nil {
		panic("couldn't create anon version of object: " + err.Error())
	}
	kv := c.ToMap()
	for k, v := range o.ToMap() {
		if _, ok := kv[k]; !ok {
			kv[k] = v
		}
	}
	c.FromMap(kv)
	return c
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

var cachedCodecIssue *goavro.Codec

// ToAvroBinary returns the data as Avro binary data
func (o *Issue) ToAvroBinary() ([]byte, *goavro.Codec, error) {
	if cachedCodecIssue == nil {
		c, err := CreateIssueAvroSchema()
		if err != nil {
			return nil, nil, err
		}
		cachedCodecIssue = c
	}
	kv := o.ToMap(true)
	jbuf, _ := json.Marshal(kv)
	native, _, err := cachedCodecIssue.NativeFromTextual(jbuf)
	if err != nil {
		return nil, nil, err
	}
	// Convert native Go form to binary Avro data
	buf, err := cachedCodecIssue.BinaryFromNative(nil, native)
	return buf, cachedCodecIssue, err
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
func (o *Issue) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
		if o.Tags == nil {
			o.Tags = make([]string, 0)
		}
	}
	return map[string]interface{}{
		"issue_id":         o.GetID(),
		"ref_id":           o.GetRefID(),
		"ref_type":         o.RefType,
		"customer_id":      o.CustomerID,
		"hashcode":         o.Hash(),
		"title":            toIssueObject(o.Title, isavro, false, "string"),
		"identifier":       toIssueObject(o.Identifier, isavro, false, "string"),
		"project_id":       toIssueObject(o.ProjectID, isavro, false, "string"),
		"url":              toIssueObject(o.URL, isavro, false, "string"),
		"created_ts":       toIssueObject(o.CreatedAt, isavro, false, "long"),
		"updated_ts":       toIssueObject(o.UpdatedAt, isavro, false, "long"),
		"planned_start_ts": toIssueObject(o.PlannedStartAt, isavro, false, "long"),
		"planned_end_ts":   toIssueObject(o.PlannedEndAt, isavro, false, "long"),
		"due_date_ts":      toIssueObject(o.DueDateAt, isavro, false, "long"),
		"priority":         toIssueObject(o.Priority, isavro, false, "string"),
		"type":             toIssueObject(o.Type, isavro, false, "string"),
		"status":           toIssueObject(o.Status, isavro, false, "string"),
		"creator_ref_id":   toIssueObject(o.CreatorRefID, isavro, false, "string"),
		"reporter_ref_id":  toIssueObject(o.ReporterRefID, isavro, false, "string"),
		"assignee_ref_id":  toIssueObject(o.AssigneeRefID, isavro, false, "string"),
		"author_ref_id":    toIssueObject(o.AuthorRefID, isavro, false, "string"),
		"tags":             toIssueObject(o.Tags, isavro, false, "string"),
		"parent_id":        toIssueObject(o.ParentID, isavro, false, "string"),
		"resolution":       toIssueObject(o.Resolution, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *Issue) FromMap(kv map[string]interface{}) {
	if val, ok := kv["issue_id"].(string); ok {
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
	if val := kv["tags"]; val != nil {
		na := make([]string, 0)
		if a, ok := val.([]string); ok {
			na = append(na, a...)
		} else {
			if a, ok := val.([]interface{}); ok {
				for _, ae := range a {
					if av, ok := ae.(string); ok {
						na = append(na, av)
					} else {
						panic("unsupported type for tags field entry: " + reflect.TypeOf(ae).String())
					}
				}
			} else if s, ok := val.(string); ok {
				for _, sv := range strings.Split(s, ",") {
					na = append(na, strings.TrimSpace(sv))
				}
			} else {
				fmt.Println(reflect.TypeOf(val).String())
				panic("unsupported type for tags field")
			}
		}
		o.Tags = na
	} else {
		o.Tags = []string{}
	}
	if o.Tags == nil {
		o.Tags = make([]string, 0)
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
	args = append(args, o.CustomerID)
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
		"namespace":    "work",
		"name":         "Issue",
		"connect.name": "work.Issue",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "issue_id",
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
				"name": "tags",
				"type": map[string]interface{}{
					"type":  "array",
					"items": "string",
				},
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
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/work/issue\\.json(\\.gz)?$"))
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
	fp := filepath.Join(dir, "/work/issue\\.json(\\.gz)?$")
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
func CreateIssueProducer(producer datamodel.Producer, ch chan Issue, errors chan<- error) <-chan bool {
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

// CreateIssueConsumer will stream data from the default topic into the provided channel
func CreateIssueConsumer(factory datamodel.ConsumerFactory, topic datamodel.TopicNameType, ch chan Issue, errors chan<- error) (<-chan bool, chan<- bool) {
	return CreateIssueConsumerForTopic(factory, IssueTopic, ch, errors)
}

// CreateIssueConsumerForTopic will stream data from the topic into the provided channel
func CreateIssueConsumerForTopic(factory datamodel.ConsumerFactory, topic datamodel.TopicNameType, ch chan Issue, errors chan<- error) (<-chan bool, chan<- bool) {
	done := make(chan bool, 1)
	closed := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		callback := datamodel.ConsumerCallback{
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
		consumer, err := factory.CreateConsumer(string(topic), callback)
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
