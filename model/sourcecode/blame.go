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
	pstrings "github.com/pinpt/go-common/strings"
)

const (
	// BlameTopic is the default topic name
	BlameTopic datamodel.TopicNameType = "sourcecode_Blame_topic"

	// BlameStream is the default stream name
	BlameStream datamodel.TopicNameType = "sourcecode_Blame_stream"

	// BlameTable is the default table name
	BlameTable datamodel.TopicNameType = "sourcecode_Blame"

	// BlameModelName is the model name
	BlameModelName datamodel.ModelNameType = "sourcecode.Blame"
)

const (
	// BlameIDColumn is the id column name
	BlameIDColumn = "id"
	// BlameRefIDColumn is the ref_id column name
	BlameRefIDColumn = "ref_id"
	// BlameRefTypeColumn is the ref_type column name
	BlameRefTypeColumn = "ref_type"
	// BlameCustomerIDColumn is the customer_id column name
	BlameCustomerIDColumn = "customer_id"
	// BlameDateAtColumn is the date_ts column name
	BlameDateAtColumn = "date_ts"
	// BlameDateColumn is the date column name
	BlameDateColumn = "date"
	// BlameRepoIDColumn is the repo_id column name
	BlameRepoIDColumn = "repo_id"
	// BlameFilenameColumn is the filename column name
	BlameFilenameColumn = "filename"
	// BlameLanguageColumn is the language column name
	BlameLanguageColumn = "language"
	// BlameSizeColumn is the size column name
	BlameSizeColumn = "size"
	// BlameLocColumn is the loc column name
	BlameLocColumn = "loc"
	// BlameSlocColumn is the sloc column name
	BlameSlocColumn = "sloc"
	// BlameBlanksColumn is the blanks column name
	BlameBlanksColumn = "blanks"
	// BlameCommentsColumn is the comments column name
	BlameCommentsColumn = "comments"
	// BlameComplexityColumn is the complexity column name
	BlameComplexityColumn = "complexity"
	// BlameShaColumn is the sha column name
	BlameShaColumn = "sha"
	// BlameCommitIDColumn is the commit_id column name
	BlameCommitIDColumn = "commit_id"
	// BlameExcludedColumn is the excluded column name
	BlameExcludedColumn = "excluded"
	// BlameExcludedReasonColumn is the excluded_reason column name
	BlameExcludedReasonColumn = "excluded_reason"
	// BlameLicenseColumn is the license column name
	BlameLicenseColumn = "license"
	// BlameStatusColumn is the status column name
	BlameStatusColumn = "status"
	// BlameLinesColumn is the lines column name
	BlameLinesColumn = "lines"
)

// Status is the enumeration type for status
type Status int32

// String returns the string value for Status
func (v Status) String() string {
	switch int32(v) {
	case 0:
		return "added"
	case 1:
		return "modified"
	case 2:
		return "removed"
	}
	return "unset"
}

const (
	// StatusAdded is the enumeration value for added
	StatusAdded Status = 0
	// StatusModified is the enumeration value for modified
	StatusModified Status = 1
	// StatusRemoved is the enumeration value for removed
	StatusRemoved Status = 2
)

// BlameLines represents the object structure for lines
type BlameLines struct {
	// Sha the sha when this line was last changed
	Sha string `json:"sha" bson:"sha" yaml:"sha" faker:"-"`
	// AuthorRefID the author ref_id of this line when last changed
	AuthorRefID string `json:"author_ref_id" bson:"author_ref_id" yaml:"author_ref_id" faker:"-"`
	// Date the change date in RFC3339 format of this line when last changed
	Date string `json:"date" bson:"date" yaml:"date" faker:"-"`
	// Comment if the line is a comment
	Comment bool `json:"comment" bson:"comment" yaml:"comment" faker:"-"`
	// Code if the line is sourcecode
	Code bool `json:"code" bson:"code" yaml:"code" faker:"-"`
	// Blank if the line is a blank line
	Blank bool `json:"blank" bson:"blank" yaml:"blank" faker:"-"`
}

func (o *BlameLines) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// Sha the sha when this line was last changed
		"sha": o.Sha,
		// AuthorRefID the author ref_id of this line when last changed
		"author_ref_id": o.AuthorRefID,
		// Date the change date in RFC3339 format of this line when last changed
		"date": o.Date,
		// Comment if the line is a comment
		"comment": o.Comment,
		// Code if the line is sourcecode
		"code": o.Code,
		// Blank if the line is a blank line
		"blank": o.Blank,
	}
}

// Blame the blame detail for each commit in a repo
type Blame struct {
	// built in types

	ID         string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	RefID      string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	RefType    string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	Hashcode   string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`

	// custom types

	// DateAt the date of the change
	DateAt int64 `json:"date_ts" bson:"date_ts" yaml:"date_ts" faker:"-"`
	// Date the date of the change in RFC3339 format
	Date string `json:"date" bson:"date" yaml:"date" faker:"-"`
	// RepoID the unique id for the repo
	RepoID string `json:"repo_id" bson:"repo_id" yaml:"repo_id" faker:"-"`
	// Filename the filename
	Filename string `json:"filename" bson:"filename" yaml:"filename" faker:"-"`
	// Language the detected language
	Language string `json:"language" bson:"language" yaml:"language" faker:"-"`
	// Size the size of the file
	Size int64 `json:"size" bson:"size" yaml:"size" faker:"-"`
	// Loc the count of lines in the file
	Loc int64 `json:"loc" bson:"loc" yaml:"loc" faker:"-"`
	// Sloc the count of source lines in the file based on language rules
	Sloc int64 `json:"sloc" bson:"sloc" yaml:"sloc" faker:"-"`
	// Blanks the count of blank lines in the file
	Blanks int64 `json:"blanks" bson:"blanks" yaml:"blanks" faker:"-"`
	// Comments the count of comment lines in the file based on language rules
	Comments int64 `json:"comments" bson:"comments" yaml:"comments" faker:"-"`
	// Complexity the cyclomatic complexity for the change
	Complexity int64 `json:"complexity" bson:"complexity" yaml:"complexity" faker:"-"`
	// Sha the commit SHA
	Sha string `json:"sha" bson:"sha" yaml:"sha" faker:"-"`
	// CommitID the commit ID
	CommitID string `json:"commit_id" bson:"commit_id" yaml:"commit_id" faker:"-"`
	// Excluded if the result was excluded
	Excluded bool `json:"excluded" bson:"excluded" yaml:"excluded" faker:"-"`
	// ExcludedReason why the result was excluded
	ExcludedReason string `json:"excluded_reason" bson:"excluded_reason" yaml:"excluded_reason" faker:"-"`
	// License if a license was detected in the file, what was the license SPDX
	License *string `json:"license" bson:"license" yaml:"license" faker:"-"`
	// Status the status of the change
	Status Status `json:"status" bson:"status" yaml:"status" faker:"-"`
	// Lines the individual line attributions
	Lines []BlameLines `json:"lines" bson:"lines" yaml:"lines" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Blame)(nil)

func toBlameObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toBlameObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if o == nil {
		return toBlameObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toBlameObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toBlameObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toBlameObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toBlameObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toBlameObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toBlameObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toBlameObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toBlameObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toBlameObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toBlameObjectNil(isavro, isoptional)
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
	case *Blame:
		return v.ToMap()
	case Blame:
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
			arr = append(arr, toBlameObject(av, isavro, false, ""))
		}
		return arr
	case Status:
		if !isavro {
			return (o.(Status)).String()
		}
		return map[string]string{
			"sourcecode.status": (o.(Status)).String(),
		}
	case *Status:
		if !isavro {
			return (o.(*Status)).String()
		}
		return map[string]string{
			"sourcecode.status": (o.(*Status)).String(),
		}
	case BlameLines:
		vv := o.(BlameLines)
		return vv.ToMap()
	case *BlameLines:
		return (*o.(*BlameLines)).ToMap()
	case []BlameLines:
		arr := make([]interface{}, 0)
		for _, i := range o.([]BlameLines) {
			arr = append(arr, i.ToMap())
		}
		return arr
	case *[]BlameLines:
		arr := make([]interface{}, 0)
		vv := o.(*[]BlameLines)
		for _, i := range *vv {
			arr = append(arr, i.ToMap())
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of Blame
func (o *Blame) String() string {
	return fmt.Sprintf("sourcecode.Blame<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Blame) GetTopicName() datamodel.TopicNameType {
	return BlameTopic
}

// GetModelName returns the name of the model
func (o *Blame) GetModelName() datamodel.ModelNameType {
	return BlameModelName
}

func (o *Blame) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Blame) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Blame", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Blame) GetTopicKey() string {
	var i interface{} = o.RepoID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Blame) GetTimestamp() time.Time {
	var dt interface{} = o.DateAt
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
	panic("not sure how to handle the date time format for Blame")
}

// GetRefID returns the RefID for the object
func (o *Blame) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Blame) IsMaterialized() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Blame) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Blame) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *Blame) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = BlameModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Blame) GetTopicConfig() *datamodel.ModelTopicConfig {
	duration, err := time.ParseDuration("168h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 168h0m0s. " + err.Error())
	}
	return &datamodel.ModelTopicConfig{
		Key:               "repo_id",
		Timestamp:         "date_ts",
		NumPartitions:     8,
		ReplicationFactor: 3,
		Retention:         duration,
		MaxSize:           5242880,
	}
}

// GetCustomerID will return the customer_id
func (o *Blame) GetCustomerID() string {
	return o.CustomerID
}

// Clone returns an exact copy of Blame
func (o *Blame) Clone() datamodel.Model {
	c := new(Blame)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Blame) Anon() datamodel.Model {
	c := new(Blame)
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
func (o *Blame) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Blame) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	// make sure that these have values if empty
	o.setDefaults()
	o.FromMap(kv)
	return nil
}

var cachedCodecBlame *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *Blame) GetAvroCodec() *goavro.Codec {
	if cachedCodecBlame == nil {
		c, err := GetBlameAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecBlame = c
	}
	return cachedCodecBlame
}

// ToAvroBinary returns the data as Avro binary data
func (o *Blame) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *Blame) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Blame objects are equal
func (o *Blame) IsEqual(other *Blame) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Blame) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
		if o.Lines == nil {
			o.Lines = make([]BlameLines, 0)
		}
	}
	return map[string]interface{}{
		"id":              o.GetID(),
		"ref_id":          o.GetRefID(),
		"ref_type":        o.RefType,
		"customer_id":     o.CustomerID,
		"hashcode":        o.Hash(),
		"date_ts":         toBlameObject(o.DateAt, isavro, false, "long"),
		"date":            toBlameObject(o.Date, isavro, false, "string"),
		"repo_id":         toBlameObject(o.RepoID, isavro, false, "string"),
		"filename":        toBlameObject(o.Filename, isavro, false, "string"),
		"language":        toBlameObject(o.Language, isavro, false, "string"),
		"size":            toBlameObject(o.Size, isavro, false, "long"),
		"loc":             toBlameObject(o.Loc, isavro, false, "long"),
		"sloc":            toBlameObject(o.Sloc, isavro, false, "long"),
		"blanks":          toBlameObject(o.Blanks, isavro, false, "long"),
		"comments":        toBlameObject(o.Comments, isavro, false, "long"),
		"complexity":      toBlameObject(o.Complexity, isavro, false, "long"),
		"sha":             toBlameObject(o.Sha, isavro, false, "string"),
		"commit_id":       toBlameObject(o.CommitID, isavro, false, "string"),
		"excluded":        toBlameObject(o.Excluded, isavro, false, "boolean"),
		"excluded_reason": toBlameObject(o.ExcludedReason, isavro, false, "string"),
		"license":         toBlameObject(o.License, isavro, true, "string"),
		"status":          toBlameObject(o.Status, isavro, false, "status"),
		"lines":           toBlameObject(o.Lines, isavro, false, "lines"),
	}
}

// FromMap attempts to load data into object from a map
func (o *Blame) FromMap(kv map[string]interface{}) {
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
	if val, ok := kv["date_ts"].(int64); ok {
		o.DateAt = val
	} else {
		val := kv["date_ts"]
		if val == nil {
			o.DateAt = number.ToInt64Any(nil)
		} else {
			o.DateAt = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["date"].(string); ok {
		o.Date = val
	} else {
		val := kv["date"]
		if val == nil {
			o.Date = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Date = fmt.Sprintf("%v", val)
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
	if val, ok := kv["filename"].(string); ok {
		o.Filename = val
	} else {
		val := kv["filename"]
		if val == nil {
			o.Filename = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Filename = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["language"].(string); ok {
		o.Language = val
	} else {
		val := kv["language"]
		if val == nil {
			o.Language = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Language = fmt.Sprintf("%v", val)
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
	if val, ok := kv["commit_id"].(string); ok {
		o.CommitID = val
	} else {
		val := kv["commit_id"]
		if val == nil {
			o.CommitID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.CommitID = fmt.Sprintf("%v", val)
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
	if val, ok := kv["excluded_reason"].(string); ok {
		o.ExcludedReason = val
	} else {
		val := kv["excluded_reason"]
		if val == nil {
			o.ExcludedReason = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.ExcludedReason = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["license"].(*string); ok {
		o.License = val
	} else if val, ok := kv["license"].(string); ok {
		o.License = &val
	} else {
		val := kv["license"]
		if val == nil {
			o.License = pstrings.Pointer("")
		} else {
			// if coming in as avro union, convert it back
			if kv, ok := val.(map[string]interface{}); ok {
				val = kv["string"]
			}
			o.License = pstrings.Pointer(fmt.Sprintf("%v", val))
		}
	}
	if val, ok := kv["status"].(Status); ok {
		o.Status = val
	} else {
		if em, ok := kv["status"].(map[string]interface{}); ok {
			ev := em["sourcecode.status"].(string)
			switch ev {
			case "added":
				o.Status = 0
			case "modified":
				o.Status = 1
			case "removed":
				o.Status = 2
			}
		}
		if em, ok := kv["status"].(string); ok {
			switch em {
			case "added":
				o.Status = 0
			case "modified":
				o.Status = 1
			case "removed":
				o.Status = 2
			}
		}
	}
	if val := kv["lines"]; val != nil {
		na := make([]BlameLines, 0)
		if a, ok := val.([]BlameLines); ok {
			na = append(na, a...)
		} else {
			if a, ok := val.([]interface{}); ok {
				for _, ae := range a {
					if av, ok := ae.(BlameLines); ok {
						na = append(na, av)
					} else {
						b, _ := json.Marshal(ae)
						var av BlameLines
						if err := json.Unmarshal(b, &av); err != nil {
							panic("unsupported type for lines field entry: " + reflect.TypeOf(ae).String())
						}
						na = append(na, av)
					}
				}
			} else {
				fmt.Println(reflect.TypeOf(val).String())
				panic("unsupported type for lines field")
			}
		}
		o.Lines = na
	} else {
		o.Lines = []BlameLines{}
	}
	if o.Lines == nil {
		o.Lines = make([]BlameLines, 0)
	}
}

// Hash will return a hashcode for the object
func (o *Blame) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.GetID())
	args = append(args, o.GetRefID())
	args = append(args, o.RefType)
	args = append(args, o.CustomerID)
	args = append(args, o.DateAt)
	args = append(args, o.Date)
	args = append(args, o.RepoID)
	args = append(args, o.Filename)
	args = append(args, o.Language)
	args = append(args, o.Size)
	args = append(args, o.Loc)
	args = append(args, o.Sloc)
	args = append(args, o.Blanks)
	args = append(args, o.Comments)
	args = append(args, o.Complexity)
	args = append(args, o.Sha)
	args = append(args, o.CommitID)
	args = append(args, o.Excluded)
	args = append(args, o.ExcludedReason)
	args = append(args, o.License)
	args = append(args, o.Status)
	args = append(args, o.Lines)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetBlameAvroSchemaSpec creates the avro schema specification for Blame
func GetBlameAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "sourcecode",
		"name":         "Blame",
		"connect.name": "sourcecode.Blame",
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
				"name": "date_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "date",
				"type": "string",
			},
			map[string]interface{}{
				"name": "repo_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "filename",
				"type": "string",
			},
			map[string]interface{}{
				"name": "language",
				"type": "string",
			},
			map[string]interface{}{
				"name": "size",
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
				"name": "blanks",
				"type": "long",
			},
			map[string]interface{}{
				"name": "comments",
				"type": "long",
			},
			map[string]interface{}{
				"name": "complexity",
				"type": "long",
			},
			map[string]interface{}{
				"name": "sha",
				"type": "string",
			},
			map[string]interface{}{
				"name": "commit_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "excluded",
				"type": "boolean",
			},
			map[string]interface{}{
				"name": "excluded_reason",
				"type": "string",
			},
			map[string]interface{}{
				"name":    "license",
				"type":    []interface{}{"null", "string"},
				"default": nil,
			},
			map[string]interface{}{
				"name": "status",
				"type": []interface{}{
					map[string]interface{}{
						"type":    "enum",
						"name":    "status",
						"symbols": []interface{}{"added", "modified", "removed"},
					},
				},
			},
			map[string]interface{}{
				"name": "lines",
				"type": map[string]interface{}{"type": "array", "name": "lines", "items": map[string]interface{}{"type": "record", "name": "lines", "fields": []interface{}{map[string]interface{}{"type": "string", "name": "sha", "doc": "the sha when this line was last changed"}, map[string]interface{}{"name": "author_ref_id", "doc": "the author ref_id of this line when last changed", "type": "string"}, map[string]interface{}{"type": "string", "name": "date", "doc": "the change date in RFC3339 format of this line when last changed"}, map[string]interface{}{"type": "boolean", "name": "comment", "doc": "if the line is a comment"}, map[string]interface{}{"type": "boolean", "name": "code", "doc": "if the line is sourcecode"}, map[string]interface{}{"doc": "if the line is a blank line", "type": "boolean", "name": "blank"}}, "doc": "the individual line attributions"}},
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetBlameAvroSchema creates the avro schema for Blame
func GetBlameAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetBlameAvroSchemaSpec())
}

// TransformBlameFunc is a function for transforming Blame during processing
type TransformBlameFunc func(input *Blame) (*Blame, error)

// NewBlamePipe creates a pipe for processing Blame items
func NewBlamePipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformBlameFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewBlameInputStream(input, errors)
	var stream chan Blame
	if len(transforms) > 0 {
		stream = make(chan Blame, 1000)
	} else {
		stream = inch
	}
	outdone := NewBlameOutputStream(output, stream, errors)
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

// NewBlameInputStreamDir creates a channel for reading Blame as JSON newlines from a directory of files
func NewBlameInputStreamDir(dir string, errors chan<- error, transforms ...TransformBlameFunc) (chan Blame, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/sourcecode/blame\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan Blame)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for blame")
		ch := make(chan Blame)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewBlameInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan Blame)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewBlameInputStreamFile creates an channel for reading Blame as JSON newlines from filename
func NewBlameInputStreamFile(filename string, errors chan<- error, transforms ...TransformBlameFunc) (chan Blame, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan Blame)
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
			ch := make(chan Blame)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewBlameInputStream(f, errors, transforms...)
}

// NewBlameInputStream creates an channel for reading Blame as JSON newlines from stream
func NewBlameInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformBlameFunc) (chan Blame, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan Blame, 1000)
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
			var item Blame
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

// NewBlameOutputStreamDir will output json newlines from channel and save in dir
func NewBlameOutputStreamDir(dir string, ch chan Blame, errors chan<- error, transforms ...TransformBlameFunc) <-chan bool {
	fp := filepath.Join(dir, "/sourcecode/blame\\.json(\\.gz)?$")
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
	return NewBlameOutputStream(gz, ch, errors, transforms...)
}

// NewBlameOutputStream will output json newlines from channel to the stream
func NewBlameOutputStream(stream io.WriteCloser, ch chan Blame, errors chan<- error, transforms ...TransformBlameFunc) <-chan bool {
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

// BlameSendEvent is an event detail for sending data
type BlameSendEvent struct {
	Blame   *Blame
	headers map[string]string
	time    time.Time
	key     string
}

var _ datamodel.ModelSendEvent = (*BlameSendEvent)(nil)

// Key is the key to use for the message
func (e *BlameSendEvent) Key() string {
	if e.key == "" {
		return e.Blame.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *BlameSendEvent) Object() datamodel.Model {
	return e.Blame
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *BlameSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *BlameSendEvent) Timestamp() time.Time {
	return e.time
}

// BlameSendEventOpts is a function handler for setting opts
type BlameSendEventOpts func(o *BlameSendEvent)

// WithBlameSendEventKey sets the key value to a value different than the object ID
func WithBlameSendEventKey(key string) BlameSendEventOpts {
	return func(o *BlameSendEvent) {
		o.key = key
	}
}

// WithBlameSendEventTimestamp sets the timestamp value
func WithBlameSendEventTimestamp(tv time.Time) BlameSendEventOpts {
	return func(o *BlameSendEvent) {
		o.time = tv
	}
}

// WithBlameSendEventHeader sets the timestamp value
func WithBlameSendEventHeader(key, value string) BlameSendEventOpts {
	return func(o *BlameSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewBlameSendEvent returns a new BlameSendEvent instance
func NewBlameSendEvent(o *Blame, opts ...BlameSendEventOpts) *BlameSendEvent {
	res := &BlameSendEvent{
		Blame: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewBlameProducer will stream data from the channel
func NewBlameProducer(producer event.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		ctx := context.Background()
		for item := range ch {
			if object, ok := item.Object().(*Blame); ok {
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
				errors <- fmt.Errorf("invalid event received. expected an object of type sourcecode.Blame but received on of type %v", reflect.TypeOf(item.Object()))
			}
		}
	}()
	return done
}

// NewBlameConsumer will stream data from the topic into the provided channel
func NewBlameConsumer(consumer event.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) {
	consumer.Consume(event.ConsumerCallback{
		OnDataReceived: func(msg event.Message) error {
			var object Blame
			if err := json.Unmarshal(msg.Value, &object); err != nil {
				return fmt.Errorf("error unmarshaling json data into sourcecode.Blame: %s", err)
			}
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &BlameReceiveEvent{&object, msg}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
	})
}

// BlameReceiveEvent is an event detail for receiving data
type BlameReceiveEvent struct {
	Blame   *Blame
	message event.Message
}

var _ datamodel.ModelReceiveEvent = (*BlameReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *BlameReceiveEvent) Object() datamodel.Model {
	return e.Blame
}

// Message returns the underlying message data for the event
func (e *BlameReceiveEvent) Message() event.Message {
	return e.message
}

// BlameProducer implements the datamodel.ModelEventProducer
type BlameProducer struct {
	ch   chan datamodel.ModelSendEvent
	done <-chan bool
}

var _ datamodel.ModelEventProducer = (*BlameProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *BlameProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *BlameProducer) Close() error {
	close(p.ch)
	<-p.done
	return nil
}

// NewProducerChannel returns a channel which can be used for producing Model events
func (o *Blame) NewProducerChannel(producer event.Producer, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent)
	return &BlameProducer{
		ch:   ch,
		done: NewBlameProducer(producer, ch, errors),
	}
}

// NewBlameProducerChannel returns a channel which can be used for producing Model events
func NewBlameProducerChannel(producer event.Producer, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent)
	return &BlameProducer{
		ch:   ch,
		done: NewBlameProducer(producer, ch, errors),
	}
}

// BlameConsumer implements the datamodel.ModelEventConsumer
type BlameConsumer struct {
	ch chan datamodel.ModelReceiveEvent
}

var _ datamodel.ModelEventConsumer = (*BlameConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *BlameConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *BlameConsumer) Close() error {
	close(c.ch)
	return nil
}

// NewConsumerChannel returns a consumer channel which can be used to consume Model events
func (o *Blame) NewConsumerChannel(consumer event.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	NewBlameConsumer(consumer, ch, errors)
	return &BlameConsumer{
		ch: ch,
	}
}

// NewBlameConsumerChannel returns a consumer channel which can be used to consume Model events
func NewBlameConsumerChannel(consumer event.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	NewBlameConsumer(consumer, ch, errors)
	return &BlameConsumer{
		ch: ch,
	}
}
