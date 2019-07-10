// DO NOT EDIT -- generated code

// Package sourcecode - the pipeline models
package sourcecode

import (
	"bufio"
	"bytes"
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
	"sync"
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
	pstrings "github.com/pinpt/go-common/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// CommitTopic is the default topic name
	CommitTopic datamodel.TopicNameType = "pipeline_sourcecode_Commit_topic"

	// CommitStream is the default stream name
	CommitStream datamodel.TopicNameType = "pipeline_sourcecode_Commit_stream"

	// CommitTable is the default table name
	CommitTable datamodel.TopicNameType = "pipeline_sourcecode_Commit"

	// CommitModelName is the model name
	CommitModelName datamodel.ModelNameType = "pipeline.sourcecode.Commit"
)

const (
	// CommitAdditionsColumn is the additions column name
	CommitAdditionsColumn = "additions"
	// CommitAuthorColumn is the author column name
	CommitAuthorColumn = "author"
	// CommitAuthorRefIDColumn is the author_ref_id column name
	CommitAuthorRefIDColumn = "author_ref_id"
	// CommitBlanksColumn is the blanks column name
	CommitBlanksColumn = "blanks"
	// CommitBranchColumn is the branch column name
	CommitBranchColumn = "branch"
	// CommitCommentsColumn is the comments column name
	CommitCommentsColumn = "comments"
	// CommitCommitterColumn is the committer column name
	CommitCommitterColumn = "committer"
	// CommitCommitterRefIDColumn is the committer_ref_id column name
	CommitCommitterRefIDColumn = "committer_ref_id"
	// CommitComplexityColumn is the complexity column name
	CommitComplexityColumn = "complexity"
	// CommitCreatedColumn is the created column name
	CommitCreatedColumn = "created"
	// CommitCreatedAtColumn is the created_ts column name
	CommitCreatedAtColumn = "created_ts"
	// CommitCustomerIDColumn is the customer_id column name
	CommitCustomerIDColumn = "customer_id"
	// CommitDeletionsColumn is the deletions column name
	CommitDeletionsColumn = "deletions"
	// CommitExcludedColumn is the excluded column name
	CommitExcludedColumn = "excluded"
	// CommitFilesColumn is the files column name
	CommitFilesColumn = "files"
	// CommitFilesChangedColumn is the files_changed column name
	CommitFilesChangedColumn = "files_changed"
	// CommitGpgSignedColumn is the gpg_signed column name
	CommitGpgSignedColumn = "gpg_signed"
	// CommitIDColumn is the id column name
	CommitIDColumn = "id"
	// CommitIssueIDColumn is the issue_id column name
	CommitIssueIDColumn = "issue_id"
	// CommitLinkedColumn is the linked column name
	CommitLinkedColumn = "linked"
	// CommitLocColumn is the loc column name
	CommitLocColumn = "loc"
	// CommitMessageColumn is the message column name
	CommitMessageColumn = "message"
	// CommitOrdinalColumn is the ordinal column name
	CommitOrdinalColumn = "ordinal"
	// CommitProjectIDColumn is the project_id column name
	CommitProjectIDColumn = "project_id"
	// CommitRefIDColumn is the ref_id column name
	CommitRefIDColumn = "ref_id"
	// CommitRefTypeColumn is the ref_type column name
	CommitRefTypeColumn = "ref_type"
	// CommitRepoIDColumn is the repo_id column name
	CommitRepoIDColumn = "repo_id"
	// CommitShaColumn is the sha column name
	CommitShaColumn = "sha"
	// CommitSizeColumn is the size column name
	CommitSizeColumn = "size"
	// CommitSlocColumn is the sloc column name
	CommitSlocColumn = "sloc"
	// CommitURLColumn is the url column name
	CommitURLColumn = "url"
)

// CommitAuthor represents the object structure for author
type CommitAuthor struct {
	// ID the corporate user id
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// TeamID the corporate team id
	TeamID string `json:"team_id" bson:"team_id" yaml:"team_id" faker:"-"`
}

func (o *CommitAuthor) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// ID the corporate user id
		"id": o.ID,
		// TeamID the corporate team id
		"team_id": o.TeamID,
	}
}

// CommitCommitter represents the object structure for committer
type CommitCommitter struct {
	// ID the corporate user id
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// TeamID the corporate team id
	TeamID string `json:"team_id" bson:"team_id" yaml:"team_id" faker:"-"`
}

func (o *CommitCommitter) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// ID the corporate user id
		"id": o.ID,
		// TeamID the corporate team id
		"team_id": o.TeamID,
	}
}

// CommitFiles represents the object structure for files
type CommitFiles struct {
	// Additions the number of additions for the commit file
	Additions int64 `json:"additions" bson:"additions" yaml:"additions" faker:"-"`
	// Binary indicates if the file was detected to be a binary file
	Binary bool `json:"binary" bson:"binary" yaml:"binary" faker:"-"`
	// Blanks the number of blank lines in the file
	Blanks int64 `json:"blanks" bson:"blanks" yaml:"blanks" faker:"-"`
	// Comments the number of comment lines in the file
	Comments int64 `json:"comments" bson:"comments" yaml:"comments" faker:"-"`
	// CommitID the unique id for the commit
	CommitID string `json:"commit_id" bson:"commit_id" yaml:"commit_id" faker:"-"`
	// Complexity the complexity value for the file change
	Complexity int64 `json:"complexity" bson:"complexity" yaml:"complexity" faker:"-"`
	// CreatedAt the timestamp in UTC that the commit was created
	CreatedAt int64 `json:"created_ts" bson:"created_ts" yaml:"created_ts" faker:"-"`
	// Deletions the number of deletions for the commit file
	Deletions int64 `json:"deletions" bson:"deletions" yaml:"deletions" faker:"-"`
	// Excluded if the file was excluded from processing
	Excluded bool `json:"excluded" bson:"excluded" yaml:"excluded" faker:"-"`
	// ExcludedReason if the file was excluded, the reason
	ExcludedReason string `json:"excluded_reason" bson:"excluded_reason" yaml:"excluded_reason" faker:"-"`
	// Filename the filename
	Filename string `json:"filename" bson:"filename" yaml:"filename" faker:"-"`
	// Language the language that was detected for the file
	Language string `json:"language" bson:"language" yaml:"language" faker:"-"`
	// License the license which was detected for the file
	License string `json:"license" bson:"license" yaml:"license" faker:"-"`
	// LicenseConfidence the license confidence from the detection engine
	LicenseConfidence float64 `json:"license_confidence" bson:"license_confidence" yaml:"license_confidence" faker:"-"`
	// Loc the number of lines in the file
	Loc int64 `json:"loc" bson:"loc" yaml:"loc" faker:"-"`
	// Ordinal the order value for the file in the change set
	Ordinal int64 `json:"ordinal" bson:"ordinal" yaml:"ordinal" faker:"-"`
	// Renamed if the file was renamed
	Renamed bool `json:"renamed" bson:"renamed" yaml:"renamed" faker:"-"`
	// RenamedFrom the original file name
	RenamedFrom string `json:"renamed_from" bson:"renamed_from" yaml:"renamed_from" faker:"-"`
	// RenamedTo the final file name
	RenamedTo string `json:"renamed_to" bson:"renamed_to" yaml:"renamed_to" faker:"-"`
	// RepoID the unique id for the repo
	RepoID string `json:"repo_id" bson:"repo_id" yaml:"repo_id" faker:"-"`
	// Size the size of the file
	Size int64 `json:"size" bson:"size" yaml:"size" faker:"-"`
	// Sloc the number of source lines in the file
	Sloc int64 `json:"sloc" bson:"sloc" yaml:"sloc" faker:"-"`
	// Status the status of the change
	Status string `json:"status" bson:"status" yaml:"status" faker:"-"`
}

func (o *CommitFiles) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// Additions the number of additions for the commit file
		"additions": o.Additions,
		// Binary indicates if the file was detected to be a binary file
		"binary": o.Binary,
		// Blanks the number of blank lines in the file
		"blanks": o.Blanks,
		// Comments the number of comment lines in the file
		"comments": o.Comments,
		// CommitID the unique id for the commit
		"commit_id": o.CommitID,
		// Complexity the complexity value for the file change
		"complexity": o.Complexity,
		// CreatedAt the timestamp in UTC that the commit was created
		"created_ts": o.CreatedAt,
		// Deletions the number of deletions for the commit file
		"deletions": o.Deletions,
		// Excluded if the file was excluded from processing
		"excluded": o.Excluded,
		// ExcludedReason if the file was excluded, the reason
		"excluded_reason": o.ExcludedReason,
		// Filename the filename
		"filename": o.Filename,
		// Language the language that was detected for the file
		"language": o.Language,
		// License the license which was detected for the file
		"license": o.License,
		// LicenseConfidence the license confidence from the detection engine
		"license_confidence": o.LicenseConfidence,
		// Loc the number of lines in the file
		"loc": o.Loc,
		// Ordinal the order value for the file in the change set
		"ordinal": o.Ordinal,
		// Renamed if the file was renamed
		"renamed": o.Renamed,
		// RenamedFrom the original file name
		"renamed_from": o.RenamedFrom,
		// RenamedTo the final file name
		"renamed_to": o.RenamedTo,
		// RepoID the unique id for the repo
		"repo_id": o.RepoID,
		// Size the size of the file
		"size": o.Size,
		// Sloc the number of source lines in the file
		"sloc": o.Sloc,
		// Status the status of the change
		"status": o.Status,
	}
}

// Commit the enriched sourcecode.Commit
type Commit struct {
	// Additions the number of additions for the commit
	Additions int64 `json:"additions" bson:"additions" yaml:"additions" faker:"-"`
	// Author the author that authored the commit
	Author CommitAuthor `json:"author" bson:"author" yaml:"author" faker:"-"`
	// AuthorRefID the author ref_id in the source system
	AuthorRefID string `json:"author_ref_id" bson:"author_ref_id" yaml:"author_ref_id" faker:"-"`
	// Blanks the number of blank lines in the commit
	Blanks int64 `json:"blanks" bson:"blanks" yaml:"blanks" faker:"-"`
	// Branch the branch that the commit was made to
	Branch string `json:"branch" bson:"branch" yaml:"branch" faker:"-"`
	// Comments the number of comment lines in the commit
	Comments int64 `json:"comments" bson:"comments" yaml:"comments" faker:"-"`
	// Committer the committer that commmitted the commit
	Committer CommitCommitter `json:"committer" bson:"committer" yaml:"committer" faker:"-"`
	// CommitterRefID the committer ref_id in the source system
	CommitterRefID string `json:"committer_ref_id" bson:"committer_ref_id" yaml:"committer_ref_id" faker:"-"`
	// Complexity the complexity value for the change
	Complexity int64 `json:"complexity" bson:"complexity" yaml:"complexity" faker:"-"`
	// Created date in rfc3339 format
	Created string `json:"created" bson:"created" yaml:"created" faker:"-"`
	// CreatedAt the timestamp in UTC that the commit was created
	CreatedAt int64 `json:"created_ts" bson:"created_ts" yaml:"created_ts" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Deletions the number of deletions for the commit
	Deletions int64 `json:"deletions" bson:"deletions" yaml:"deletions" faker:"-"`
	// Excluded if the commit was excluded
	Excluded bool `json:"excluded" bson:"excluded" yaml:"excluded" faker:"-"`
	// Files the files touched by this commit
	Files []CommitFiles `json:"files" bson:"files" yaml:"files" faker:"-"`
	// FilesChanged the number of files changed for the commit
	FilesChanged int64 `json:"files_changed" bson:"files_changed" yaml:"files_changed" faker:"-"`
	// GpgSigned if the commit was signed
	GpgSigned bool `json:"gpg_signed" bson:"gpg_signed" yaml:"gpg_signed" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// IssueID the issues this commit is linked to
	IssueID []string `json:"issue_id" bson:"issue_id" yaml:"issue_id" faker:"-"`
	// Linked if the commit is linked to an issue
	Linked bool `json:"linked" bson:"linked" yaml:"linked" faker:"-"`
	// Loc the number of lines in the commit
	Loc int64 `json:"loc" bson:"loc" yaml:"loc" faker:"-"`
	// Message the commit message
	Message string `json:"message" bson:"message" yaml:"message" faker:"commit_message"`
	// Ordinal the order of the commit in the commit stream
	Ordinal int64 `json:"ordinal" bson:"ordinal" yaml:"ordinal" faker:"-"`
	// ProjectID the project of the issue that this commit is linked to
	ProjectID *string `json:"project_id" bson:"project_id" yaml:"project_id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RepoID the unique id for the repo
	RepoID string `json:"repo_id" bson:"repo_id" yaml:"repo_id" faker:"-"`
	// Sha the unique sha for the commit
	Sha string `json:"sha" bson:"sha" yaml:"sha" faker:"sha"`
	// Size the size of all files in the commit
	Size int64 `json:"size" bson:"size" yaml:"size" faker:"-"`
	// Sloc the number of source lines in the commit
	Sloc int64 `json:"sloc" bson:"sloc" yaml:"sloc" faker:"-"`
	// URL the url to the commit detail
	URL string `json:"url" bson:"url" yaml:"url" faker:"url"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
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

	case CommitAuthor:
		vv := o.(CommitAuthor)
		return vv.ToMap()
	case *CommitAuthor:
		return (*o.(*CommitAuthor)).ToMap()
	case []CommitAuthor:
		arr := make([]interface{}, 0)
		for _, i := range o.([]CommitAuthor) {
			arr = append(arr, i.ToMap())
		}
		return arr
	case *[]CommitAuthor:
		arr := make([]interface{}, 0)
		vv := o.(*[]CommitAuthor)
		for _, i := range *vv {
			arr = append(arr, i.ToMap())
		}
		return arr
	case CommitCommitter:
		vv := o.(CommitCommitter)
		return vv.ToMap()
	case *CommitCommitter:
		return (*o.(*CommitCommitter)).ToMap()
	case []CommitCommitter:
		arr := make([]interface{}, 0)
		for _, i := range o.([]CommitCommitter) {
			arr = append(arr, i.ToMap())
		}
		return arr
	case *[]CommitCommitter:
		arr := make([]interface{}, 0)
		vv := o.(*[]CommitCommitter)
		for _, i := range *vv {
			arr = append(arr, i.ToMap())
		}
		return arr
	case CommitFiles:
		vv := o.(CommitFiles)
		return vv.ToMap()
	case *CommitFiles:
		return (*o.(*CommitFiles)).ToMap()
	case []CommitFiles:
		arr := make([]interface{}, 0)
		for _, i := range o.([]CommitFiles) {
			arr = append(arr, i.ToMap())
		}
		return arr
	case *[]CommitFiles:
		arr := make([]interface{}, 0)
		vv := o.(*[]CommitFiles)
		for _, i := range *vv {
			arr = append(arr, i.ToMap())
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of Commit
func (o *Commit) String() string {
	return fmt.Sprintf("pipeline.sourcecode.Commit<%s>", o.ID)
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
	return true
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Commit) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	idletime, err := time.ParseDuration("1s")
	if err != nil {
		panic(err)
	}
	return &datamodel.ModelMaterializeConfig{
		KeyName:   "id",
		TableName: "pipeline_sourcecode_commit",
		BatchSize: 5000,
		IdleTime:  idletime,
	}
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
	retention, err := time.ParseDuration("168h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 168h0m0s. " + err.Error())
	}

	ttl, err := time.ParseDuration("0s")
	if err != nil {
		ttl = 0
	}
	return &datamodel.ModelTopicConfig{
		Key:               "repo_id",
		Timestamp:         "created_ts",
		NumPartitions:     8,
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetStateKey returns a key for use in state store
func (o *Commit) GetStateKey() string {
	key := "repo_id"
	return fmt.Sprintf("%s_%s", key, o.GetID())
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

// FromAvroBinary will convert from Avro binary data into data in this object
func (o *Commit) FromAvroBinary(value []byte) error {
	var nullHeader = []byte{byte(0)}
	// if this still has the schema encoded in the header, move past it to the avro payload
	if bytes.HasPrefix(value, nullHeader) {
		value = value[5:]
	}
	kv, _, err := o.GetAvroCodec().NativeFromBinary(value)
	if err != nil {
		return err
	}
	o.FromMap(kv.(map[string]interface{}))
	return nil
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
		if o.Files == nil {
			o.Files = make([]CommitFiles, 0)
		}
		if o.IssueID == nil {
			o.IssueID = make([]string, 0)
		}
	}
	o.setDefaults()
	return map[string]interface{}{
		"additions":        toCommitObject(o.Additions, isavro, false, "long"),
		"author":           toCommitObject(o.Author, isavro, false, "author"),
		"author_ref_id":    toCommitObject(o.AuthorRefID, isavro, false, "string"),
		"blanks":           toCommitObject(o.Blanks, isavro, false, "long"),
		"branch":           toCommitObject(o.Branch, isavro, false, "string"),
		"comments":         toCommitObject(o.Comments, isavro, false, "long"),
		"committer":        toCommitObject(o.Committer, isavro, false, "committer"),
		"committer_ref_id": toCommitObject(o.CommitterRefID, isavro, false, "string"),
		"complexity":       toCommitObject(o.Complexity, isavro, false, "long"),
		"created":          toCommitObject(o.Created, isavro, false, "string"),
		"created_ts":       toCommitObject(o.CreatedAt, isavro, false, "long"),
		"customer_id":      toCommitObject(o.CustomerID, isavro, false, "string"),
		"deletions":        toCommitObject(o.Deletions, isavro, false, "long"),
		"excluded":         toCommitObject(o.Excluded, isavro, false, "boolean"),
		"files":            toCommitObject(o.Files, isavro, false, "files"),
		"files_changed":    toCommitObject(o.FilesChanged, isavro, false, "long"),
		"gpg_signed":       toCommitObject(o.GpgSigned, isavro, false, "boolean"),
		"id":               toCommitObject(o.ID, isavro, false, "string"),
		"issue_id":         toCommitObject(o.IssueID, isavro, false, "issue_id"),
		"linked":           toCommitObject(o.Linked, isavro, false, "boolean"),
		"loc":              toCommitObject(o.Loc, isavro, false, "long"),
		"message":          toCommitObject(o.Message, isavro, false, "string"),
		"ordinal":          toCommitObject(o.Ordinal, isavro, false, "long"),
		"project_id":       toCommitObject(o.ProjectID, isavro, true, "string"),
		"ref_id":           toCommitObject(o.RefID, isavro, false, "string"),
		"ref_type":         toCommitObject(o.RefType, isavro, false, "string"),
		"repo_id":          toCommitObject(o.RepoID, isavro, false, "string"),
		"sha":              toCommitObject(o.Sha, isavro, false, "string"),
		"size":             toCommitObject(o.Size, isavro, false, "long"),
		"sloc":             toCommitObject(o.Sloc, isavro, false, "long"),
		"url":              toCommitObject(o.URL, isavro, false, "string"),
		"hashcode":         toCommitObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *Commit) FromMap(kv map[string]interface{}) {
	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["additions"].(int64); ok {
		o.Additions = val
	} else {
		val := kv["additions"]
		if val == nil {
			o.Additions = number.ToInt64Any(nil)
		} else {
			if tv, ok := val.(time.Time); ok {
				val = datetime.TimeToEpoch(tv)
			}
			o.Additions = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["author"].(CommitAuthor); ok {
		o.Author = val
	} else {
		val := kv["author"]
		if val == nil {
			o.Author = CommitAuthor{}
		} else {
			o.Author = CommitAuthor{}
			b, _ := json.Marshal(val)
			json.Unmarshal(b, &o.Author)

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
	if val, ok := kv["blanks"].(int64); ok {
		o.Blanks = val
	} else {
		val := kv["blanks"]
		if val == nil {
			o.Blanks = number.ToInt64Any(nil)
		} else {
			if tv, ok := val.(time.Time); ok {
				val = datetime.TimeToEpoch(tv)
			}
			o.Blanks = number.ToInt64Any(val)
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
	if val, ok := kv["comments"].(int64); ok {
		o.Comments = val
	} else {
		val := kv["comments"]
		if val == nil {
			o.Comments = number.ToInt64Any(nil)
		} else {
			if tv, ok := val.(time.Time); ok {
				val = datetime.TimeToEpoch(tv)
			}
			o.Comments = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["committer"].(CommitCommitter); ok {
		o.Committer = val
	} else {
		val := kv["committer"]
		if val == nil {
			o.Committer = CommitCommitter{}
		} else {
			o.Committer = CommitCommitter{}
			b, _ := json.Marshal(val)
			json.Unmarshal(b, &o.Committer)

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
	if val, ok := kv["complexity"].(int64); ok {
		o.Complexity = val
	} else {
		val := kv["complexity"]
		if val == nil {
			o.Complexity = number.ToInt64Any(nil)
		} else {
			if tv, ok := val.(time.Time); ok {
				val = datetime.TimeToEpoch(tv)
			}
			o.Complexity = number.ToInt64Any(val)
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
	if val, ok := kv["created_ts"].(int64); ok {
		o.CreatedAt = val
	} else {
		val := kv["created_ts"]
		if val == nil {
			o.CreatedAt = number.ToInt64Any(nil)
		} else {
			if tv, ok := val.(time.Time); ok {
				val = datetime.TimeToEpoch(tv)
			}
			o.CreatedAt = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["customer_id"].(string); ok {
		o.CustomerID = val
	} else {
		val := kv["customer_id"]
		if val == nil {
			o.CustomerID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.CustomerID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["deletions"].(int64); ok {
		o.Deletions = val
	} else {
		val := kv["deletions"]
		if val == nil {
			o.Deletions = number.ToInt64Any(nil)
		} else {
			if tv, ok := val.(time.Time); ok {
				val = datetime.TimeToEpoch(tv)
			}
			o.Deletions = number.ToInt64Any(val)
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
	if val := kv["files"]; val != nil {
		na := make([]CommitFiles, 0)
		if a, ok := val.([]CommitFiles); ok {
			na = append(na, a...)
		} else {
			if a, ok := val.([]interface{}); ok {
				for _, ae := range a {
					if av, ok := ae.(CommitFiles); ok {
						na = append(na, av)
					} else {
						b, _ := json.Marshal(ae)
						var av CommitFiles
						if err := json.Unmarshal(b, &av); err != nil {
							panic("unsupported type for files field entry: " + reflect.TypeOf(ae).String())
						}
						na = append(na, av)
					}
				}
			} else if a, ok := val.(primitive.A); ok {
				for _, ae := range a {
					if av, ok := ae.(CommitFiles); ok {
						na = append(na, av)
					} else {
						b, _ := json.Marshal(ae)
						var av CommitFiles
						if err := json.Unmarshal(b, &av); err != nil {
							panic("unsupported type for files field entry: " + reflect.TypeOf(ae).String())
						}
						na = append(na, av)
					}
				}
			} else {
				fmt.Println(reflect.TypeOf(val).String())
				panic("unsupported type for files field")
			}
		}
		o.Files = na
	} else {
		o.Files = []CommitFiles{}
	}
	if o.Files == nil {
		o.Files = make([]CommitFiles, 0)
	}
	if val, ok := kv["files_changed"].(int64); ok {
		o.FilesChanged = val
	} else {
		val := kv["files_changed"]
		if val == nil {
			o.FilesChanged = number.ToInt64Any(nil)
		} else {
			if tv, ok := val.(time.Time); ok {
				val = datetime.TimeToEpoch(tv)
			}
			o.FilesChanged = number.ToInt64Any(val)
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
	if val, ok := kv["id"].(string); ok {
		o.ID = val
	} else {
		val := kv["id"]
		if val == nil {
			o.ID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.ID = fmt.Sprintf("%v", val)
		}
	}
	if val := kv["issue_id"]; val != nil {
		na := make([]string, 0)
		if a, ok := val.([]string); ok {
			na = append(na, a...)
		} else {
			if a, ok := val.([]interface{}); ok {
				for _, ae := range a {
					if av, ok := ae.(string); ok {
						na = append(na, av)
					} else {
						b, _ := json.Marshal(ae)
						var av string
						if err := json.Unmarshal(b, &av); err != nil {
							panic("unsupported type for issue_id field entry: " + reflect.TypeOf(ae).String())
						}
						na = append(na, av)
					}
				}
			} else if s, ok := val.(string); ok {
				for _, sv := range strings.Split(s, ",") {
					na = append(na, strings.TrimSpace(sv))
				}
			} else if a, ok := val.(primitive.A); ok {
				for _, ae := range a {
					if av, ok := ae.(string); ok {
						na = append(na, av)
					} else {
						b, _ := json.Marshal(ae)
						var av string
						if err := json.Unmarshal(b, &av); err != nil {
							panic("unsupported type for issue_id field entry: " + reflect.TypeOf(ae).String())
						}
						na = append(na, av)
					}
				}
			} else {
				fmt.Println(reflect.TypeOf(val).String())
				panic("unsupported type for issue_id field")
			}
		}
		o.IssueID = na
	} else {
		o.IssueID = []string{}
	}
	if o.IssueID == nil {
		o.IssueID = make([]string, 0)
	}
	if val, ok := kv["linked"].(bool); ok {
		o.Linked = val
	} else {
		val := kv["linked"]
		if val == nil {
			o.Linked = number.ToBoolAny(nil)
		} else {
			o.Linked = number.ToBoolAny(val)
		}
	}
	if val, ok := kv["loc"].(int64); ok {
		o.Loc = val
	} else {
		val := kv["loc"]
		if val == nil {
			o.Loc = number.ToInt64Any(nil)
		} else {
			if tv, ok := val.(time.Time); ok {
				val = datetime.TimeToEpoch(tv)
			}
			o.Loc = number.ToInt64Any(val)
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
	if val, ok := kv["ordinal"].(int64); ok {
		o.Ordinal = val
	} else {
		val := kv["ordinal"]
		if val == nil {
			o.Ordinal = number.ToInt64Any(nil)
		} else {
			if tv, ok := val.(time.Time); ok {
				val = datetime.TimeToEpoch(tv)
			}
			o.Ordinal = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["project_id"].(*string); ok {
		o.ProjectID = val
	} else if val, ok := kv["project_id"].(string); ok {
		o.ProjectID = &val
	} else {
		val := kv["project_id"]
		if val == nil {
			o.ProjectID = pstrings.Pointer("")
		} else {
			// if coming in as avro union, convert it back
			if kv, ok := val.(map[string]interface{}); ok {
				val = kv["string"]
			}
			o.ProjectID = pstrings.Pointer(fmt.Sprintf("%v", val))
		}
	}
	if val, ok := kv["ref_id"].(string); ok {
		o.RefID = val
	} else {
		val := kv["ref_id"]
		if val == nil {
			o.RefID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.RefID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["ref_type"].(string); ok {
		o.RefType = val
	} else {
		val := kv["ref_type"]
		if val == nil {
			o.RefType = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.RefType = fmt.Sprintf("%v", val)
		}
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
	if val, ok := kv["size"].(int64); ok {
		o.Size = val
	} else {
		val := kv["size"]
		if val == nil {
			o.Size = number.ToInt64Any(nil)
		} else {
			if tv, ok := val.(time.Time); ok {
				val = datetime.TimeToEpoch(tv)
			}
			o.Size = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["sloc"].(int64); ok {
		o.Sloc = val
	} else {
		val := kv["sloc"]
		if val == nil {
			o.Sloc = number.ToInt64Any(nil)
		} else {
			if tv, ok := val.(time.Time); ok {
				val = datetime.TimeToEpoch(tv)
			}
			o.Sloc = number.ToInt64Any(val)
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
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *Commit) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Additions)
	args = append(args, o.Author)
	args = append(args, o.AuthorRefID)
	args = append(args, o.Blanks)
	args = append(args, o.Branch)
	args = append(args, o.Comments)
	args = append(args, o.Committer)
	args = append(args, o.CommitterRefID)
	args = append(args, o.Complexity)
	args = append(args, o.Created)
	args = append(args, o.CreatedAt)
	args = append(args, o.CustomerID)
	args = append(args, o.Deletions)
	args = append(args, o.Excluded)
	args = append(args, o.Files)
	args = append(args, o.FilesChanged)
	args = append(args, o.GpgSigned)
	args = append(args, o.ID)
	args = append(args, o.IssueID)
	args = append(args, o.Linked)
	args = append(args, o.Loc)
	args = append(args, o.Message)
	args = append(args, o.Ordinal)
	args = append(args, o.ProjectID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RepoID)
	args = append(args, o.Sha)
	args = append(args, o.Size)
	args = append(args, o.Sloc)
	args = append(args, o.URL)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetCommitAvroSchemaSpec creates the avro schema specification for Commit
func GetCommitAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "pipeline.sourcecode",
		"name":      "Commit",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "hashcode",
				"type": "string",
			},
			map[string]interface{}{
				"name": "additions",
				"type": "long",
			},
			map[string]interface{}{
				"name": "author",
				"type": map[string]interface{}{"type": "record", "name": "author", "fields": []interface{}{map[string]interface{}{"type": "string", "name": "id", "doc": "the corporate user id"}, map[string]interface{}{"doc": "the corporate team id", "type": "string", "name": "team_id"}}, "doc": "the author that authored the commit"},
			},
			map[string]interface{}{
				"name": "author_ref_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "blanks",
				"type": "long",
			},
			map[string]interface{}{
				"name": "branch",
				"type": "string",
			},
			map[string]interface{}{
				"name": "comments",
				"type": "long",
			},
			map[string]interface{}{
				"name": "committer",
				"type": map[string]interface{}{"fields": []interface{}{map[string]interface{}{"type": "string", "name": "id", "doc": "the corporate user id"}, map[string]interface{}{"type": "string", "name": "team_id", "doc": "the corporate team id"}}, "doc": "the committer that commmitted the commit", "type": "record", "name": "committer"},
			},
			map[string]interface{}{
				"name": "committer_ref_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "complexity",
				"type": "long",
			},
			map[string]interface{}{
				"name": "created",
				"type": "string",
			},
			map[string]interface{}{
				"name": "created_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "customer_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "deletions",
				"type": "long",
			},
			map[string]interface{}{
				"name": "excluded",
				"type": "boolean",
			},
			map[string]interface{}{
				"name": "files",
				"type": map[string]interface{}{"items": map[string]interface{}{"type": "record", "name": "files", "fields": []interface{}{map[string]interface{}{"name": "additions", "doc": "the number of additions for the commit file", "type": "long"}, map[string]interface{}{"type": "boolean", "name": "binary", "doc": "indicates if the file was detected to be a binary file"}, map[string]interface{}{"type": "long", "name": "blanks", "doc": "the number of blank lines in the file"}, map[string]interface{}{"type": "long", "name": "comments", "doc": "the number of comment lines in the file"}, map[string]interface{}{"doc": "the unique id for the commit", "type": "string", "name": "commit_id"}, map[string]interface{}{"name": "complexity", "doc": "the complexity value for the file change", "type": "long"}, map[string]interface{}{"type": "long", "name": "created_ts", "doc": "the timestamp in UTC that the commit was created"}, map[string]interface{}{"type": "long", "name": "deletions", "doc": "the number of deletions for the commit file"}, map[string]interface{}{"type": "boolean", "name": "excluded", "doc": "if the file was excluded from processing"}, map[string]interface{}{"type": "string", "name": "excluded_reason", "doc": "if the file was excluded, the reason"}, map[string]interface{}{"name": "filename", "doc": "the filename", "type": "string"}, map[string]interface{}{"doc": "the language that was detected for the file", "type": "string", "name": "language"}, map[string]interface{}{"name": "license", "doc": "the license which was detected for the file", "type": "string"}, map[string]interface{}{"doc": "the license confidence from the detection engine", "type": "float", "name": "license_confidence"}, map[string]interface{}{"type": "long", "name": "loc", "doc": "the number of lines in the file"}, map[string]interface{}{"type": "long", "name": "ordinal", "doc": "the order value for the file in the change set"}, map[string]interface{}{"name": "renamed", "doc": "if the file was renamed", "type": "boolean"}, map[string]interface{}{"type": "string", "name": "renamed_from", "doc": "the original file name"}, map[string]interface{}{"type": "string", "name": "renamed_to", "doc": "the final file name"}, map[string]interface{}{"name": "repo_id", "doc": "the unique id for the repo", "type": "string"}, map[string]interface{}{"type": "long", "name": "size", "doc": "the size of the file"}, map[string]interface{}{"type": "long", "name": "sloc", "doc": "the number of source lines in the file"}, map[string]interface{}{"type": "string", "name": "status", "doc": "the status of the change"}}, "doc": "the files touched by this commit"}, "type": "array", "name": "files"},
			},
			map[string]interface{}{
				"name": "files_changed",
				"type": "long",
			},
			map[string]interface{}{
				"name": "gpg_signed",
				"type": "boolean",
			},
			map[string]interface{}{
				"name": "id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "issue_id",
				"type": map[string]interface{}{"type": "array", "name": "issue_id", "items": "string"},
			},
			map[string]interface{}{
				"name": "linked",
				"type": "boolean",
			},
			map[string]interface{}{
				"name": "loc",
				"type": "long",
			},
			map[string]interface{}{
				"name": "message",
				"type": "string",
			},
			map[string]interface{}{
				"name": "ordinal",
				"type": "long",
			},
			map[string]interface{}{
				"name":    "project_id",
				"type":    []interface{}{"null", "string"},
				"default": nil,
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
				"name": "repo_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "sha",
				"type": "string",
			},
			map[string]interface{}{
				"name": "size",
				"type": "long",
			},
			map[string]interface{}{
				"name": "sloc",
				"type": "long",
			},
			map[string]interface{}{
				"name": "url",
				"type": "string",
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
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/pipeline.sourcecode/commit\\.json(\\.gz)?$"))
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
	fp := filepath.Join(dir, "/pipeline.sourcecode/commit\\.json(\\.gz)?$")
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
func NewCommitProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		for {
			select {
			case <-ctx.Done():
				return
			case item := <-ch:
				if item == nil {
					empty <- true
					return
				}
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
					errors <- fmt.Errorf("invalid event received. expected an object of type pipeline.sourcecode.Commit but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewCommitConsumer will stream data from the topic into the provided channel
func NewCommitConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object Commit
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into pipeline.sourcecode.Commit: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avro data into pipeline.sourcecode.Commit: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for pipeline.sourcecode.Commit")
			}
			// ignore messages that have exceeded the TTL
			cfg := object.GetTopicConfig()
			if cfg != nil && cfg.TTL != 0 && msg.Timestamp.Add(cfg.TTL).Sub(time.Now()) < 0 {
				return nil
			}
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &CommitReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object Commit
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &CommitReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// CommitReceiveEvent is an event detail for receiving data
type CommitReceiveEvent struct {
	Commit  *Commit
	message eventing.Message
	eof     bool
}

var _ datamodel.ModelReceiveEvent = (*CommitReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *CommitReceiveEvent) Object() datamodel.Model {
	return e.Commit
}

// Message returns the underlying message data for the event
func (e *CommitReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *CommitReceiveEvent) EOF() bool {
	return e.eof
}

// CommitProducer implements the datamodel.ModelEventProducer
type CommitProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*CommitProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *CommitProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *CommitProducer) Close() error {
	p.mu.Lock()
	closed := p.closed
	p.closed = true
	p.mu.Unlock()
	if !closed {
		close(p.ch)
		<-p.empty
		p.cancel()
		<-p.done
	}
	return nil
}

// NewProducerChannel returns a channel which can be used for producing Model events
func (o *Commit) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *Commit) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &CommitProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewCommitProducer(newctx, producer, ch, errors, empty),
	}
}

// NewCommitProducerChannel returns a channel which can be used for producing Model events
func NewCommitProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewCommitProducerChannelSize(producer, 0, errors)
}

// NewCommitProducerChannelSize returns a channel which can be used for producing Model events
func NewCommitProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &CommitProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewCommitProducer(newctx, producer, ch, errors, empty),
	}
}

// CommitConsumer implements the datamodel.ModelEventConsumer
type CommitConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*CommitConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *CommitConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *CommitConsumer) Close() error {
	c.mu.Lock()
	closed := c.closed
	c.closed = true
	c.mu.Unlock()
	var err error
	if !closed {
		c.callback.Close()
		err = c.consumer.Close()
	}
	return err
}

// NewConsumerChannel returns a consumer channel which can be used to consume Model events
func (o *Commit) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &CommitConsumer{
		ch:       ch,
		callback: NewCommitConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewCommitConsumerChannel returns a consumer channel which can be used to consume Model events
func NewCommitConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &CommitConsumer{
		ch:       ch,
		callback: NewCommitConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
