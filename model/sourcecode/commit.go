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
	"github.com/pinpt/go-common/event"
	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
)

const (
	// CommitTopic is the default topic name
	CommitTopic datamodel.TopicNameType = "sourcecode_Commit_topic"

	// CommitStream is the default stream name
	CommitStream datamodel.TopicNameType = "sourcecode_Commit_stream"

	// CommitTable is the default table name
	CommitTable datamodel.TopicNameType = "sourcecode_Commit"

	// CommitModelName is the model name
	CommitModelName datamodel.ModelNameType = "sourcecode.Commit"
)

const (
	// CommitIDColumn is the id column name
	CommitIDColumn = "id"
	// CommitRefIDColumn is the ref_id column name
	CommitRefIDColumn = "ref_id"
	// CommitRefTypeColumn is the ref_type column name
	CommitRefTypeColumn = "ref_type"
	// CommitCustomerIDColumn is the customer_id column name
	CommitCustomerIDColumn = "customer_id"
	// CommitRepoIDColumn is the repo_id column name
	CommitRepoIDColumn = "repo_id"
	// CommitShaColumn is the sha column name
	CommitShaColumn = "sha"
	// CommitMessageColumn is the message column name
	CommitMessageColumn = "message"
	// CommitURLColumn is the url column name
	CommitURLColumn = "url"
	// CommitCreatedAtColumn is the created_ts column name
	CommitCreatedAtColumn = "created_ts"
	// CommitCreatedColumn is the created column name
	CommitCreatedColumn = "created"
	// CommitBranchColumn is the branch column name
	CommitBranchColumn = "branch"
	// CommitAdditionsColumn is the additions column name
	CommitAdditionsColumn = "additions"
	// CommitDeletionsColumn is the deletions column name
	CommitDeletionsColumn = "deletions"
	// CommitFilesChangedColumn is the files_changed column name
	CommitFilesChangedColumn = "files_changed"
	// CommitAuthorRefIDColumn is the author_ref_id column name
	CommitAuthorRefIDColumn = "author_ref_id"
	// CommitCommitterRefIDColumn is the committer_ref_id column name
	CommitCommitterRefIDColumn = "committer_ref_id"
	// CommitOrdinalColumn is the ordinal column name
	CommitOrdinalColumn = "ordinal"
	// CommitLocColumn is the loc column name
	CommitLocColumn = "loc"
	// CommitSlocColumn is the sloc column name
	CommitSlocColumn = "sloc"
	// CommitCommentsColumn is the comments column name
	CommitCommentsColumn = "comments"
	// CommitBlanksColumn is the blanks column name
	CommitBlanksColumn = "blanks"
	// CommitSizeColumn is the size column name
	CommitSizeColumn = "size"
	// CommitComplexityColumn is the complexity column name
	CommitComplexityColumn = "complexity"
	// CommitGpgSignedColumn is the gpg_signed column name
	CommitGpgSignedColumn = "gpg_signed"
	// CommitExcludedColumn is the excluded column name
	CommitExcludedColumn = "excluded"
)

// Commit the commit is a specific change in a repo
type Commit struct {
	// built in types

	ID         string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	RefID      string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	RefType    string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	Hashcode   string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`

	// custom types

	// RepoID the unique id for the repo
	RepoID string `json:"repo_id" bson:"repo_id" yaml:"repo_id" faker:"-"`
	// Sha the unique sha for the commit
	Sha string `json:"sha" bson:"sha" yaml:"sha" faker:"sha"`
	// Message the commit message
	Message string `json:"message" bson:"message" yaml:"message" faker:"commit_message"`
	// URL the url to the commit detail
	URL string `json:"url" bson:"url" yaml:"url" faker:"url"`
	// CreatedAt the timestamp in UTC that the commit was created
	CreatedAt int64 `json:"created_ts" bson:"created_ts" yaml:"created_ts" faker:"-"`
	// Created date in rfc3339 format
	Created string `json:"created" bson:"created" yaml:"created" faker:"-"`
	// Branch the branch that the commit was made to
	Branch string `json:"branch" bson:"branch" yaml:"branch" faker:"-"`
	// Additions the number of additions for the commit
	Additions int64 `json:"additions" bson:"additions" yaml:"additions" faker:"-"`
	// Deletions the number of deletions for the commit
	Deletions int64 `json:"deletions" bson:"deletions" yaml:"deletions" faker:"-"`
	// FilesChanged the number of files changed for the commit
	FilesChanged int64 `json:"files_changed" bson:"files_changed" yaml:"files_changed" faker:"-"`
	// AuthorRefID the author ref_id in the source system
	AuthorRefID string `json:"author_ref_id" bson:"author_ref_id" yaml:"author_ref_id" faker:"-"`
	// CommitterRefID the committer ref_id in the source system
	CommitterRefID string `json:"committer_ref_id" bson:"committer_ref_id" yaml:"committer_ref_id" faker:"-"`
	// Ordinal the order of the commit in the commit stream
	Ordinal int64 `json:"ordinal" bson:"ordinal" yaml:"ordinal" faker:"-"`
	// Loc the number of lines in the commit
	Loc int64 `json:"loc" bson:"loc" yaml:"loc" faker:"-"`
	// Sloc the number of source lines in the commit
	Sloc int64 `json:"sloc" bson:"sloc" yaml:"sloc" faker:"-"`
	// Comments the number of comment lines in the commit
	Comments int64 `json:"comments" bson:"comments" yaml:"comments" faker:"-"`
	// Blanks the number of blank lines in the commit
	Blanks int64 `json:"blanks" bson:"blanks" yaml:"blanks" faker:"-"`
	// Size the size of all files in the commit
	Size int64 `json:"size" bson:"size" yaml:"size" faker:"-"`
	// Complexity the complexity value for the change
	Complexity int64 `json:"complexity" bson:"complexity" yaml:"complexity" faker:"-"`
	// GpgSigned if the commit was signed
	GpgSigned bool `json:"gpg_signed" bson:"gpg_signed" yaml:"gpg_signed" faker:"-"`
	// Excluded if the commit was excluded
	Excluded bool `json:"excluded" bson:"excluded" yaml:"excluded" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Commit)(nil)

func toCommitObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toCommitObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if o == nil {
		return toCommitObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toCommitObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toCommitObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toCommitObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toCommitObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toCommitObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toCommitObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toCommitObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toCommitObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toCommitObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toCommitObjectNil(isavro, isoptional)
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
	case *Commit:
		return v.ToMap()
	case Commit:
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
			arr = append(arr, toCommitObject(av, isavro, false, ""))
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of Commit
func (o *Commit) String() string {
	return fmt.Sprintf("sourcecode.Commit<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Commit) GetTopicName() datamodel.TopicNameType {
	return CommitTopic
}

// GetModelName returns the name of the model
func (o *Commit) GetModelName() datamodel.ModelNameType {
	return CommitModelName
}

func (o *Commit) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Commit) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Commit", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Commit) GetTopicKey() string {
	var i interface{} = o.RepoID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Commit) GetTimestamp() time.Time {
	var dt interface{} = o.CreatedAt
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
	panic("not sure how to handle the date time format for Commit")
}

// GetRefID returns the RefID for the object
func (o *Commit) GetRefID() string {
	o.RefID = o.Sha
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Commit) IsMaterialized() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Commit) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Commit) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *Commit) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = CommitModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Commit) GetTopicConfig() *datamodel.ModelTopicConfig {
	duration, err := time.ParseDuration("168h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 168h0m0s. " + err.Error())
	}
	return &datamodel.ModelTopicConfig{
		Key:               "repo_id",
		Timestamp:         "created_ts",
		NumPartitions:     8,
		ReplicationFactor: 3,
		Retention:         duration,
		MaxSize:           5242880,
	}
}

// GetCustomerID will return the customer_id
func (o *Commit) GetCustomerID() string {
	return o.CustomerID
}

// Clone returns an exact copy of Commit
func (o *Commit) Clone() datamodel.Model {
	c := new(Commit)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Commit) Anon() datamodel.Model {
	c := new(Commit)
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
func (o *Commit) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Commit) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	// make sure that these have values if empty
	o.setDefaults()
	o.FromMap(kv)
	return nil
}

var cachedCodecCommit *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *Commit) GetAvroCodec() *goavro.Codec {
	if cachedCodecCommit == nil {
		c, err := GetCommitAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecCommit = c
	}
	return cachedCodecCommit
}

// ToAvroBinary returns the data as Avro binary data
func (o *Commit) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *Commit) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Commit objects are equal
func (o *Commit) IsEqual(other *Commit) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Commit) ToMap(avro ...bool) map[string]interface{} {
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
		"repo_id":          toCommitObject(o.RepoID, isavro, false, "string"),
		"sha":              toCommitObject(o.Sha, isavro, false, "string"),
		"message":          toCommitObject(o.Message, isavro, false, "string"),
		"url":              toCommitObject(o.URL, isavro, false, "string"),
		"created_ts":       toCommitObject(o.CreatedAt, isavro, false, "long"),
		"created":          toCommitObject(o.Created, isavro, false, "string"),
		"branch":           toCommitObject(o.Branch, isavro, false, "string"),
		"additions":        toCommitObject(o.Additions, isavro, false, "long"),
		"deletions":        toCommitObject(o.Deletions, isavro, false, "long"),
		"files_changed":    toCommitObject(o.FilesChanged, isavro, false, "long"),
		"author_ref_id":    toCommitObject(o.AuthorRefID, isavro, false, "string"),
		"committer_ref_id": toCommitObject(o.CommitterRefID, isavro, false, "string"),
		"ordinal":          toCommitObject(o.Ordinal, isavro, false, "long"),
		"loc":              toCommitObject(o.Loc, isavro, false, "long"),
		"sloc":             toCommitObject(o.Sloc, isavro, false, "long"),
		"comments":         toCommitObject(o.Comments, isavro, false, "long"),
		"blanks":           toCommitObject(o.Blanks, isavro, false, "long"),
		"size":             toCommitObject(o.Size, isavro, false, "long"),
		"complexity":       toCommitObject(o.Complexity, isavro, false, "long"),
		"gpg_signed":       toCommitObject(o.GpgSigned, isavro, false, "boolean"),
		"excluded":         toCommitObject(o.Excluded, isavro, false, "boolean"),
	}
}

// FromMap attempts to load data into object from a map
func (o *Commit) FromMap(kv map[string]interface{}) {
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
	if val, ok := kv["sha"].(string); ok {
		o.Sha = val
	} else {
		val := kv["sha"]
		if val == nil {
			o.Sha = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Sha = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["message"].(string); ok {
		o.Message = val
	} else {
		val := kv["message"]
		if val == nil {
			o.Message = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Message = fmt.Sprintf("%v", val)
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
	if val, ok := kv["created"].(string); ok {
		o.Created = val
	} else {
		val := kv["created"]
		if val == nil {
			o.Created = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Created = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["branch"].(string); ok {
		o.Branch = val
	} else {
		val := kv["branch"]
		if val == nil {
			o.Branch = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Branch = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["additions"].(int64); ok {
		o.Additions = val
	} else {
		val := kv["additions"]
		if val == nil {
			o.Additions = number.ToInt64Any(nil)
		} else {
			o.Additions = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["deletions"].(int64); ok {
		o.Deletions = val
	} else {
		val := kv["deletions"]
		if val == nil {
			o.Deletions = number.ToInt64Any(nil)
		} else {
			o.Deletions = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["files_changed"].(int64); ok {
		o.FilesChanged = val
	} else {
		val := kv["files_changed"]
		if val == nil {
			o.FilesChanged = number.ToInt64Any(nil)
		} else {
			o.FilesChanged = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["author_ref_id"].(string); ok {
		o.AuthorRefID = val
	} else {
		val := kv["author_ref_id"]
		if val == nil {
			o.AuthorRefID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.AuthorRefID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["committer_ref_id"].(string); ok {
		o.CommitterRefID = val
	} else {
		val := kv["committer_ref_id"]
		if val == nil {
			o.CommitterRefID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.CommitterRefID = fmt.Sprintf("%v", val)
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
	if val, ok := kv["loc"].(int64); ok {
		o.Loc = val
	} else {
		val := kv["loc"]
		if val == nil {
			o.Loc = number.ToInt64Any(nil)
		} else {
			o.Loc = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["sloc"].(int64); ok {
		o.Sloc = val
	} else {
		val := kv["sloc"]
		if val == nil {
			o.Sloc = number.ToInt64Any(nil)
		} else {
			o.Sloc = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["comments"].(int64); ok {
		o.Comments = val
	} else {
		val := kv["comments"]
		if val == nil {
			o.Comments = number.ToInt64Any(nil)
		} else {
			o.Comments = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["blanks"].(int64); ok {
		o.Blanks = val
	} else {
		val := kv["blanks"]
		if val == nil {
			o.Blanks = number.ToInt64Any(nil)
		} else {
			o.Blanks = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["size"].(int64); ok {
		o.Size = val
	} else {
		val := kv["size"]
		if val == nil {
			o.Size = number.ToInt64Any(nil)
		} else {
			o.Size = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["complexity"].(int64); ok {
		o.Complexity = val
	} else {
		val := kv["complexity"]
		if val == nil {
			o.Complexity = number.ToInt64Any(nil)
		} else {
			o.Complexity = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["gpg_signed"].(bool); ok {
		o.GpgSigned = val
	} else {
		val := kv["gpg_signed"]
		if val == nil {
			o.GpgSigned = number.ToBoolAny(nil)
		} else {
			o.GpgSigned = number.ToBoolAny(val)
		}
	}
	if val, ok := kv["excluded"].(bool); ok {
		o.Excluded = val
	} else {
		val := kv["excluded"]
		if val == nil {
			o.Excluded = number.ToBoolAny(nil)
		} else {
			o.Excluded = number.ToBoolAny(val)
		}
	}
}

// Hash will return a hashcode for the object
func (o *Commit) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.GetID())
	args = append(args, o.GetRefID())
	args = append(args, o.RefType)
	args = append(args, o.CustomerID)
	args = append(args, o.RepoID)
	args = append(args, o.Sha)
	args = append(args, o.Message)
	args = append(args, o.URL)
	args = append(args, o.CreatedAt)
	args = append(args, o.Created)
	args = append(args, o.Branch)
	args = append(args, o.Additions)
	args = append(args, o.Deletions)
	args = append(args, o.FilesChanged)
	args = append(args, o.AuthorRefID)
	args = append(args, o.CommitterRefID)
	args = append(args, o.Ordinal)
	args = append(args, o.Loc)
	args = append(args, o.Sloc)
	args = append(args, o.Comments)
	args = append(args, o.Blanks)
	args = append(args, o.Size)
	args = append(args, o.Complexity)
	args = append(args, o.GpgSigned)
	args = append(args, o.Excluded)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetCommitAvroSchemaSpec creates the avro schema specification for Commit
func GetCommitAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "sourcecode",
		"name":         "Commit",
		"connect.name": "sourcecode.Commit",
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
				"name": "sha",
				"type": "string",
			},
			map[string]interface{}{
				"name": "message",
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
				"name": "created",
				"type": "string",
			},
			map[string]interface{}{
				"name": "branch",
				"type": "string",
			},
			map[string]interface{}{
				"name": "additions",
				"type": "long",
			},
			map[string]interface{}{
				"name": "deletions",
				"type": "long",
			},
			map[string]interface{}{
				"name": "files_changed",
				"type": "long",
			},
			map[string]interface{}{
				"name": "author_ref_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "committer_ref_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "ordinal",
				"type": "long",
			},
			map[string]interface{}{
				"name": "loc",
				"type": "long",
			},
			map[string]interface{}{
				"name": "sloc",
				"type": "long",
			},
			map[string]interface{}{
				"name": "comments",
				"type": "long",
			},
			map[string]interface{}{
				"name": "blanks",
				"type": "long",
			},
			map[string]interface{}{
				"name": "size",
				"type": "long",
			},
			map[string]interface{}{
				"name": "complexity",
				"type": "long",
			},
			map[string]interface{}{
				"name": "gpg_signed",
				"type": "boolean",
			},
			map[string]interface{}{
				"name": "excluded",
				"type": "boolean",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetCommitAvroSchema creates the avro schema for Commit
func GetCommitAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetCommitAvroSchemaSpec())
}

// TransformCommitFunc is a function for transforming Commit during processing
type TransformCommitFunc func(input *Commit) (*Commit, error)

// NewCommitPipe creates a pipe for processing Commit items
func NewCommitPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformCommitFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewCommitInputStream(input, errors)
	var stream chan Commit
	if len(transforms) > 0 {
		stream = make(chan Commit, 1000)
	} else {
		stream = inch
	}
	outdone := NewCommitOutputStream(output, stream, errors)
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

// NewCommitInputStreamDir creates a channel for reading Commit as JSON newlines from a directory of files
func NewCommitInputStreamDir(dir string, errors chan<- error, transforms ...TransformCommitFunc) (chan Commit, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/sourcecode/commit\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan Commit)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for commit")
		ch := make(chan Commit)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewCommitInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan Commit)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewCommitInputStreamFile creates an channel for reading Commit as JSON newlines from filename
func NewCommitInputStreamFile(filename string, errors chan<- error, transforms ...TransformCommitFunc) (chan Commit, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan Commit)
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
			ch := make(chan Commit)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewCommitInputStream(f, errors, transforms...)
}

// NewCommitInputStream creates an channel for reading Commit as JSON newlines from stream
func NewCommitInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformCommitFunc) (chan Commit, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan Commit, 1000)
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
			var item Commit
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

// NewCommitOutputStreamDir will output json newlines from channel and save in dir
func NewCommitOutputStreamDir(dir string, ch chan Commit, errors chan<- error, transforms ...TransformCommitFunc) <-chan bool {
	fp := filepath.Join(dir, "/sourcecode/commit\\.json(\\.gz)?$")
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
	return NewCommitOutputStream(gz, ch, errors, transforms...)
}

// NewCommitOutputStream will output json newlines from channel to the stream
func NewCommitOutputStream(stream io.WriteCloser, ch chan Commit, errors chan<- error, transforms ...TransformCommitFunc) <-chan bool {
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

// CommitSendEvent is an event detail for sending data
type CommitSendEvent struct {
	Commit  *Commit
	headers map[string]string
	time    time.Time
	key     string
}

var _ datamodel.ModelSendEvent = (*CommitSendEvent)(nil)

// Key is the key to use for the message
func (e *CommitSendEvent) Key() string {
	if e.key == "" {
		return e.Commit.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *CommitSendEvent) Object() datamodel.Model {
	return e.Commit
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *CommitSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *CommitSendEvent) Timestamp() time.Time {
	return e.time
}

// CommitSendEventOpts is a function handler for setting opts
type CommitSendEventOpts func(o *CommitSendEvent)

// WithCommitSendEventKey sets the key value to a value different than the object ID
func WithCommitSendEventKey(key string) CommitSendEventOpts {
	return func(o *CommitSendEvent) {
		o.key = key
	}
}

// WithCommitSendEventTimestamp sets the timestamp value
func WithCommitSendEventTimestamp(tv time.Time) CommitSendEventOpts {
	return func(o *CommitSendEvent) {
		o.time = tv
	}
}

// WithCommitSendEventHeader sets the timestamp value
func WithCommitSendEventHeader(key, value string) CommitSendEventOpts {
	return func(o *CommitSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewCommitSendEvent returns a new CommitSendEvent instance
func NewCommitSendEvent(o *Commit, opts ...CommitSendEventOpts) *CommitSendEvent {
	res := &CommitSendEvent{
		Commit: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewCommitProducer will stream data from the channel
func NewCommitProducer(producer event.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		ctx := context.Background()
		for item := range ch {
			if object, ok := item.Object().(*Commit); ok {
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
				msg := event.Message{
					Key:       item.Key(),
					Value:     binary,
					Codec:     codec,
					Headers:   headers,
					Timestamp: tv,
				}
				if err := producer.Send(ctx, msg); err != nil {
					errors <- fmt.Errorf("error sending %s. %v", object.String(), err)
				}
			} else {
				errors <- fmt.Errorf("invalid event received. expected an object of type sourcecode.Commit but received on of type %v", reflect.TypeOf(item.Object()))
			}
		}
	}()
	return done
}

// NewCommitConsumer will stream data from the topic into the provided channel
func NewCommitConsumer(consumer event.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) {
	consumer.Consume(event.ConsumerCallback{
		OnDataReceived: func(msg event.Message) error {
			var object Commit
			if err := json.Unmarshal(msg.Value, &object); err != nil {
				return fmt.Errorf("error unmarshaling json data into sourcecode.Commit: %s", err)
			}
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &CommitReceiveEvent{&object, msg}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
	})
}

// CommitReceiveEvent is an event detail for receiving data
type CommitReceiveEvent struct {
	Commit  *Commit
	message event.Message
}

var _ datamodel.ModelReceiveEvent = (*CommitReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *CommitReceiveEvent) Object() datamodel.Model {
	return e.Commit
}

// Message returns the underlying message data for the event
func (e *CommitReceiveEvent) Message() event.Message {
	return e.message
}

// CommitProducer implements the datamodel.ModelEventProducer
type CommitProducer struct {
	ch   chan datamodel.ModelSendEvent
	done <-chan bool
}

var _ datamodel.ModelEventProducer = (*CommitProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *CommitProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *CommitProducer) Close() error {
	close(p.ch)
	<-p.done
	return nil
}

// NewProducerChannel returns a channel which can be used for producing Model events
func (o *Commit) NewProducerChannel(producer event.Producer, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent)
	return &CommitProducer{
		ch:   ch,
		done: NewCommitProducer(producer, ch, errors),
	}
}

// NewCommitProducerChannel returns a channel which can be used for producing Model events
func NewCommitProducerChannel(producer event.Producer, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent)
	return &CommitProducer{
		ch:   ch,
		done: NewCommitProducer(producer, ch, errors),
	}
}

// CommitConsumer implements the datamodel.ModelEventConsumer
type CommitConsumer struct {
	ch chan datamodel.ModelReceiveEvent
}

var _ datamodel.ModelEventConsumer = (*CommitConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *CommitConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *CommitConsumer) Close() error {
	close(c.ch)
	return nil
}

// NewConsumerChannel returns a consumer channel which can be used to consume Model events
func (o *Commit) NewConsumerChannel(consumer event.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	NewCommitConsumer(consumer, ch, errors)
	return &CommitConsumer{
		ch: ch,
	}
}

// NewCommitConsumerChannel returns a consumer channel which can be used to consume Model events
func NewCommitConsumerChannel(consumer event.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	NewCommitConsumer(consumer, ch, errors)
	return &CommitConsumer{
		ch: ch,
	}
}
