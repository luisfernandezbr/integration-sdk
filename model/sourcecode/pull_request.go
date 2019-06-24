// DO NOT EDIT -- generated code

// Package sourcecode - the system which contains source code
package sourcecode

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
	"time"

	"github.com/bxcodec/faker"
	"github.com/linkedin/goavro"
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/eventing"
	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
)

const (
	// PullRequestTopic is the default topic name
	PullRequestTopic datamodel.TopicNameType = "sourcecode_PullRequest_topic"

	// PullRequestStream is the default stream name
	PullRequestStream datamodel.TopicNameType = "sourcecode_PullRequest_stream"

	// PullRequestTable is the default table name
	PullRequestTable datamodel.TopicNameType = "sourcecode_PullRequest"

	// PullRequestModelName is the model name
	PullRequestModelName datamodel.ModelNameType = "sourcecode.PullRequest"
)

const (
	// PullRequestIDColumn is the id column name
	PullRequestIDColumn = "id"
	// PullRequestRefIDColumn is the ref_id column name
	PullRequestRefIDColumn = "ref_id"
	// PullRequestRefTypeColumn is the ref_type column name
	PullRequestRefTypeColumn = "ref_type"
	// PullRequestCustomerIDColumn is the customer_id column name
	PullRequestCustomerIDColumn = "customer_id"
	// PullRequestRepoIDColumn is the repo_id column name
	PullRequestRepoIDColumn = "repo_id"
	// PullRequestTitleColumn is the title column name
	PullRequestTitleColumn = "title"
	// PullRequestDescriptionColumn is the description column name
	PullRequestDescriptionColumn = "description"
	// PullRequestURLColumn is the url column name
	PullRequestURLColumn = "url"
	// PullRequestCreatedAtColumn is the created_ts column name
	PullRequestCreatedAtColumn = "created_ts"
	// PullRequestMergedAtColumn is the merged_ts column name
	PullRequestMergedAtColumn = "merged_ts"
	// PullRequestClosedAtColumn is the closed_ts column name
	PullRequestClosedAtColumn = "closed_ts"
	// PullRequestClosedByRefIDColumn is the closed_by_ref_id column name
	PullRequestClosedByRefIDColumn = "closed_by_ref_id"
	// PullRequestClosedByLoginColumn is the closed_by_login column name
	PullRequestClosedByLoginColumn = "closed_by_login"
	// PullRequestUpdatedAtColumn is the updated_ts column name
	PullRequestUpdatedAtColumn = "updated_ts"
	// PullRequestStatusColumn is the status column name
	PullRequestStatusColumn = "status"
	// PullRequestUserRefIDColumn is the user_ref_id column name
	PullRequestUserRefIDColumn = "user_ref_id"
)

// PullRequest the pull request for a given repo
type PullRequest struct {
	// built in types

	ID         string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	RefID      string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	RefType    string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	Hashcode   string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`

	// custom types

	// RepoID the unique id for the repo
	RepoID string `json:"repo_id" bson:"repo_id" yaml:"repo_id" faker:"-"`
	// Title the title of the pull request
	Title string `json:"title" bson:"title" yaml:"title" faker:"commit_message"`
	// Description the description of the pull request
	Description string `json:"description" bson:"description" yaml:"description" faker:"-"`
	// URL the url to the pull request home page
	URL string `json:"url" bson:"url" yaml:"url" faker:"url"`
	// CreatedAt the timestamp in UTC that the pull request was created
	CreatedAt int64 `json:"created_ts" bson:"created_ts" yaml:"created_ts" faker:"-"`
	// MergedAt the timestamp in UTC that the pull request was merged
	MergedAt int64 `json:"merged_ts" bson:"merged_ts" yaml:"merged_ts" faker:"-"`
	// ClosedAt the timestamp in UTC that the pull request was closed
	ClosedAt int64 `json:"closed_ts" bson:"closed_ts" yaml:"closed_ts" faker:"-"`
	// ClosedByRefID the id of user who closed the pull request
	ClosedByRefID string `json:"closed_by_ref_id" bson:"closed_by_ref_id" yaml:"closed_by_ref_id" faker:"-"`
	// ClosedByLogin the username/login of user who closed the pull request
	ClosedByLogin string `json:"closed_by_login" bson:"closed_by_login" yaml:"closed_by_login" faker:"-"`
	// UpdatedAt the timestamp in UTC that the pull request was closed
	UpdatedAt int64 `json:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// Status the status of the pull request
	Status string `json:"status" bson:"status" yaml:"status" faker:"-"`
	// UserRefID the user ref_id in the source system
	UserRefID string `json:"user_ref_id" bson:"user_ref_id" yaml:"user_ref_id" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*PullRequest)(nil)

func toPullRequestObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toPullRequestObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if o == nil {
		return toPullRequestObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toPullRequestObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toPullRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toPullRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toPullRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toPullRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toPullRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toPullRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toPullRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toPullRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toPullRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case map[string]interface{}:
		return o
	case *map[string]interface{}:
		return v
	case map[string]string:
		return v
	case *map[string]string:
		return *v
	case *PullRequest:
		return v.ToMap()
	case PullRequest:
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
			arr = append(arr, toPullRequestObject(av, isavro, false, ""))
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of PullRequest
func (o *PullRequest) String() string {
	return fmt.Sprintf("sourcecode.PullRequest<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *PullRequest) GetTopicName() datamodel.TopicNameType {
	return PullRequestTopic
}

// GetModelName returns the name of the model
func (o *PullRequest) GetModelName() datamodel.ModelNameType {
	return PullRequestModelName
}

func (o *PullRequest) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *PullRequest) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("PullRequest", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *PullRequest) GetTopicKey() string {
	var i interface{} = o.RepoID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *PullRequest) GetTimestamp() time.Time {
	var dt interface{} = o.UpdatedAt
	switch v := dt.(type) {
	case int64:
		return datetime.DateFromEpoch(v).UTC()
	case string:
		tv, err := datetime.ISODateToTime(v)
		if err != nil {
			panic(err)
		}
		return tv.UTC()
	case time.Time:
		return v.UTC()
	}
	panic("not sure how to handle the date time format for PullRequest")
}

// GetRefID returns the RefID for the object
func (o *PullRequest) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *PullRequest) IsMaterialized() bool {
	return true
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *PullRequest) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	idletime, err := time.ParseDuration("1s")
	if err != nil {
		panic(err)
	}
	return &datamodel.ModelMaterializeConfig{
		KeyName:   "id",
		TableName: "sourcecode_pullrequest",
		BatchSize: 5000,
		IdleTime:  idletime,
	}
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *PullRequest) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *PullRequest) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = PullRequestModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *PullRequest) GetTopicConfig() *datamodel.ModelTopicConfig {
	duration, err := time.ParseDuration("168h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 168h0m0s. " + err.Error())
	}
	return &datamodel.ModelTopicConfig{
		Key:               "repo_id",
		Timestamp:         "updated_ts",
		NumPartitions:     8,
		ReplicationFactor: 3,
		Retention:         duration,
		MaxSize:           5242880,
	}
}

// GetCustomerID will return the customer_id
func (o *PullRequest) GetCustomerID() string {
	return o.CustomerID
}

// Clone returns an exact copy of PullRequest
func (o *PullRequest) Clone() datamodel.Model {
	c := new(PullRequest)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *PullRequest) Anon() datamodel.Model {
	c := new(PullRequest)
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
func (o *PullRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *PullRequest) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	// make sure that these have values if empty
	o.setDefaults()
	o.FromMap(kv)
	return nil
}

var cachedCodecPullRequest *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *PullRequest) GetAvroCodec() *goavro.Codec {
	if cachedCodecPullRequest == nil {
		c, err := GetPullRequestAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecPullRequest = c
	}
	return cachedCodecPullRequest
}

// ToAvroBinary returns the data as Avro binary data
func (o *PullRequest) ToAvroBinary() ([]byte, *goavro.Codec, error) {
	kv := o.ToMap(true)
	jbuf, _ := json.Marshal(kv)
	codec := o.GetAvroCodec()
	native, _, err := codec.NativeFromTextual(jbuf)
	if err != nil {
		return nil, nil, err
	}
	// Convert native Go form to binary Avro data
	buf, err := codec.BinaryFromNative(nil, native)
	return buf, codec, err
}

// Stringify returns the object in JSON format as a string
func (o *PullRequest) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two PullRequest objects are equal
func (o *PullRequest) IsEqual(other *PullRequest) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *PullRequest) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	return map[string]interface{}{
		"id":               o.GetID(),
		"ref_id":           o.GetRefID(),
		"ref_type":         o.RefType,
		"customer_id":      o.CustomerID,
		"hashcode":         o.Hash(),
		"repo_id":          toPullRequestObject(o.RepoID, isavro, false, "string"),
		"title":            toPullRequestObject(o.Title, isavro, false, "string"),
		"description":      toPullRequestObject(o.Description, isavro, false, "string"),
		"url":              toPullRequestObject(o.URL, isavro, false, "string"),
		"created_ts":       toPullRequestObject(o.CreatedAt, isavro, false, "long"),
		"merged_ts":        toPullRequestObject(o.MergedAt, isavro, false, "long"),
		"closed_ts":        toPullRequestObject(o.ClosedAt, isavro, false, "long"),
		"closed_by_ref_id": toPullRequestObject(o.ClosedByRefID, isavro, false, "string"),
		"closed_by_login":  toPullRequestObject(o.ClosedByLogin, isavro, false, "string"),
		"updated_ts":       toPullRequestObject(o.UpdatedAt, isavro, false, "long"),
		"status":           toPullRequestObject(o.Status, isavro, false, "string"),
		"user_ref_id":      toPullRequestObject(o.UserRefID, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *PullRequest) FromMap(kv map[string]interface{}) {
	// make sure that these have values if empty
	o.setDefaults()
	if val, ok := kv["id"].(string); ok {
		o.ID = val
	} else if val, ok := kv["_id"].(string); ok {
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
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.RepoID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["title"].(string); ok {
		o.Title = val
	} else {
		val := kv["title"]
		if val == nil {
			o.Title = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Title = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["description"].(string); ok {
		o.Description = val
	} else {
		val := kv["description"]
		if val == nil {
			o.Description = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Description = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["url"].(string); ok {
		o.URL = val
	} else {
		val := kv["url"]
		if val == nil {
			o.URL = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
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
	if val, ok := kv["merged_ts"].(int64); ok {
		o.MergedAt = val
	} else {
		val := kv["merged_ts"]
		if val == nil {
			o.MergedAt = number.ToInt64Any(nil)
		} else {
			o.MergedAt = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["closed_ts"].(int64); ok {
		o.ClosedAt = val
	} else {
		val := kv["closed_ts"]
		if val == nil {
			o.ClosedAt = number.ToInt64Any(nil)
		} else {
			o.ClosedAt = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["closed_by_ref_id"].(string); ok {
		o.ClosedByRefID = val
	} else {
		val := kv["closed_by_ref_id"]
		if val == nil {
			o.ClosedByRefID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.ClosedByRefID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["closed_by_login"].(string); ok {
		o.ClosedByLogin = val
	} else {
		val := kv["closed_by_login"]
		if val == nil {
			o.ClosedByLogin = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.ClosedByLogin = fmt.Sprintf("%v", val)
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
	if val, ok := kv["status"].(string); ok {
		o.Status = val
	} else {
		val := kv["status"]
		if val == nil {
			o.Status = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Status = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["user_ref_id"].(string); ok {
		o.UserRefID = val
	} else {
		val := kv["user_ref_id"]
		if val == nil {
			o.UserRefID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.UserRefID = fmt.Sprintf("%v", val)
		}
	}
}

// Hash will return a hashcode for the object
func (o *PullRequest) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.GetID())
	args = append(args, o.GetRefID())
	args = append(args, o.RefType)
	args = append(args, o.CustomerID)
	args = append(args, o.RepoID)
	args = append(args, o.Title)
	args = append(args, o.Description)
	args = append(args, o.URL)
	args = append(args, o.CreatedAt)
	args = append(args, o.MergedAt)
	args = append(args, o.ClosedAt)
	args = append(args, o.ClosedByRefID)
	args = append(args, o.ClosedByLogin)
	args = append(args, o.UpdatedAt)
	args = append(args, o.Status)
	args = append(args, o.UserRefID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetPullRequestAvroSchemaSpec creates the avro schema specification for PullRequest
func GetPullRequestAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "sourcecode",
		"name":      "PullRequest",
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
				"name": "repo_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "title",
				"type": "string",
			},
			map[string]interface{}{
				"name": "description",
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
				"name": "merged_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "closed_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "closed_by_ref_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "closed_by_login",
				"type": "string",
			},
			map[string]interface{}{
				"name": "updated_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "status",
				"type": "string",
			},
			map[string]interface{}{
				"name": "user_ref_id",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetPullRequestAvroSchema creates the avro schema for PullRequest
func GetPullRequestAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetPullRequestAvroSchemaSpec())
}

// TransformPullRequestFunc is a function for transforming PullRequest during processing
type TransformPullRequestFunc func(input *PullRequest) (*PullRequest, error)

// NewPullRequestPipe creates a pipe for processing PullRequest items
func NewPullRequestPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformPullRequestFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewPullRequestInputStream(input, errors)
	var stream chan PullRequest
	if len(transforms) > 0 {
		stream = make(chan PullRequest, 1000)
	} else {
		stream = inch
	}
	outdone := NewPullRequestOutputStream(output, stream, errors)
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

// NewPullRequestInputStreamDir creates a channel for reading PullRequest as JSON newlines from a directory of files
func NewPullRequestInputStreamDir(dir string, errors chan<- error, transforms ...TransformPullRequestFunc) (chan PullRequest, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/sourcecode/pull_request\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan PullRequest)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for pull_request")
		ch := make(chan PullRequest)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewPullRequestInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan PullRequest)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewPullRequestInputStreamFile creates an channel for reading PullRequest as JSON newlines from filename
func NewPullRequestInputStreamFile(filename string, errors chan<- error, transforms ...TransformPullRequestFunc) (chan PullRequest, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan PullRequest)
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
			ch := make(chan PullRequest)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewPullRequestInputStream(f, errors, transforms...)
}

// NewPullRequestInputStream creates an channel for reading PullRequest as JSON newlines from stream
func NewPullRequestInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformPullRequestFunc) (chan PullRequest, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan PullRequest, 1000)
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
			var item PullRequest
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

// NewPullRequestOutputStreamDir will output json newlines from channel and save in dir
func NewPullRequestOutputStreamDir(dir string, ch chan PullRequest, errors chan<- error, transforms ...TransformPullRequestFunc) <-chan bool {
	fp := filepath.Join(dir, "/sourcecode/pull_request\\.json(\\.gz)?$")
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
	return NewPullRequestOutputStream(gz, ch, errors, transforms...)
}

// NewPullRequestOutputStream will output json newlines from channel to the stream
func NewPullRequestOutputStream(stream io.WriteCloser, ch chan PullRequest, errors chan<- error, transforms ...TransformPullRequestFunc) <-chan bool {
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

// PullRequestSendEvent is an event detail for sending data
type PullRequestSendEvent struct {
	PullRequest *PullRequest
	headers     map[string]string
	time        time.Time
	key         string
}

var _ datamodel.ModelSendEvent = (*PullRequestSendEvent)(nil)

// Key is the key to use for the message
func (e *PullRequestSendEvent) Key() string {
	if e.key == "" {
		return e.PullRequest.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *PullRequestSendEvent) Object() datamodel.Model {
	return e.PullRequest
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *PullRequestSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *PullRequestSendEvent) Timestamp() time.Time {
	return e.time
}

// PullRequestSendEventOpts is a function handler for setting opts
type PullRequestSendEventOpts func(o *PullRequestSendEvent)

// WithPullRequestSendEventKey sets the key value to a value different than the object ID
func WithPullRequestSendEventKey(key string) PullRequestSendEventOpts {
	return func(o *PullRequestSendEvent) {
		o.key = key
	}
}

// WithPullRequestSendEventTimestamp sets the timestamp value
func WithPullRequestSendEventTimestamp(tv time.Time) PullRequestSendEventOpts {
	return func(o *PullRequestSendEvent) {
		o.time = tv
	}
}

// WithPullRequestSendEventHeader sets the timestamp value
func WithPullRequestSendEventHeader(key, value string) PullRequestSendEventOpts {
	return func(o *PullRequestSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewPullRequestSendEvent returns a new PullRequestSendEvent instance
func NewPullRequestSendEvent(o *PullRequest, opts ...PullRequestSendEventOpts) *PullRequestSendEvent {
	res := &PullRequestSendEvent{
		PullRequest: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewPullRequestProducer will stream data from the channel
func NewPullRequestProducer(producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		ctx := context.Background()
		for item := range ch {
			if object, ok := item.Object().(*PullRequest); ok {
				binary, codec, err := object.ToAvroBinary()
				if err != nil {
					errors <- fmt.Errorf("error encoding %s to avro binary data. %v", object.String(), err)
					return
				}
				headers := map[string]string{}
				object.SetEventHeaders(headers)
				for k, v := range item.Headers() {
					headers[k] = v
				}
				tv := item.Timestamp()
				if tv.IsZero() {
					tv = object.GetTimestamp() // if not provided in the message, use the objects value
				}
				if tv.IsZero() {
					tv = time.Now() // if its still zero, use the ingest time
				}
				msg := eventing.Message{
					Encoding:  eventing.AvroEncoding,
					Key:       item.Key(),
					Value:     binary,
					Codec:     codec,
					Headers:   headers,
					Timestamp: tv,
					Topic:     object.GetTopicName().String(),
				}
				if err := producer.Send(ctx, msg); err != nil {
					errors <- fmt.Errorf("error sending %s. %v", object.String(), err)
				}
			} else {
				errors <- fmt.Errorf("invalid event received. expected an object of type sourcecode.PullRequest but received on of type %v", reflect.TypeOf(item.Object()))
			}
		}
	}()
	return done
}

// NewPullRequestConsumer will stream data from the topic into the provided channel
func NewPullRequestConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) {
	consumer.Consume(&eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object PullRequest
			if err := json.Unmarshal(msg.Value, &object); err != nil {
				return fmt.Errorf("error unmarshaling json data into sourcecode.PullRequest: %s", err)
			}
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &PullRequestReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object PullRequest
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &PullRequestReceiveEvent{nil, msg, true}
		},
	})
}

// PullRequestReceiveEvent is an event detail for receiving data
type PullRequestReceiveEvent struct {
	PullRequest *PullRequest
	message     eventing.Message
	eof         bool
}

var _ datamodel.ModelReceiveEvent = (*PullRequestReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *PullRequestReceiveEvent) Object() datamodel.Model {
	return e.PullRequest
}

// Message returns the underlying message data for the event
func (e *PullRequestReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *PullRequestReceiveEvent) EOF() bool {
	return e.eof
}

// PullRequestProducer implements the datamodel.ModelEventProducer
type PullRequestProducer struct {
	ch   chan datamodel.ModelSendEvent
	done <-chan bool
}

var _ datamodel.ModelEventProducer = (*PullRequestProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *PullRequestProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *PullRequestProducer) Close() error {
	close(p.ch)
	<-p.done
	return nil
}

// NewProducerChannel returns a channel which can be used for producing Model events
func (o *PullRequest) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent)
	return &PullRequestProducer{
		ch:   ch,
		done: NewPullRequestProducer(producer, ch, errors),
	}
}

// NewPullRequestProducerChannel returns a channel which can be used for producing Model events
func NewPullRequestProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent)
	return &PullRequestProducer{
		ch:   ch,
		done: NewPullRequestProducer(producer, ch, errors),
	}
}

// PullRequestConsumer implements the datamodel.ModelEventConsumer
type PullRequestConsumer struct {
	ch chan datamodel.ModelReceiveEvent
}

var _ datamodel.ModelEventConsumer = (*PullRequestConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *PullRequestConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *PullRequestConsumer) Close() error {
	close(c.ch)
	return nil
}

// NewConsumerChannel returns a consumer channel which can be used to consume Model events
func (o *PullRequest) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	NewPullRequestConsumer(consumer, ch, errors)
	return &PullRequestConsumer{
		ch: ch,
	}
}

// NewPullRequestConsumerChannel returns a consumer channel which can be used to consume Model events
func NewPullRequestConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	NewPullRequestConsumer(consumer, ch, errors)
	return &PullRequestConsumer{
		ch: ch,
	}
}
