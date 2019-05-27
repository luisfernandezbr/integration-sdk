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
	"github.com/pinpt/go-common/event"
	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
)

// CommitFileTopic is the default topic name
const CommitFileTopic datamodel.TopicNameType = "sourcecode_CommitFile_topic"

// CommitFileStream is the default stream name
const CommitFileStream datamodel.TopicNameType = "sourcecode_CommitFile_stream"

// CommitFileTable is the default table name
const CommitFileTable datamodel.TopicNameType = "sourcecode_CommitFile"

// CommitFileModelName is the model name
const CommitFileModelName datamodel.ModelNameType = "sourcecode.CommitFile"

// CommitFile the file change for a specific commit
type CommitFile struct {
	// built in types

	ID         string `json:"commit_file_id" yaml:"commit_file_id" faker:"-"`
	RefID      string `json:"ref_id" yaml:"ref_id" faker:"-"`
	RefType    string `json:"ref_type" yaml:"ref_type" faker:"-"`
	CustomerID string `json:"customer_id" yaml:"customer_id" faker:"-"`
	Hashcode   string `json:"hashcode" yaml:"hashcode" faker:"-"`
	// custom types

	// CreatedAt the timestamp in UTC that the commit was created
	CreatedAt int64 `json:"created_ts" yaml:"created_ts" faker:"-"`
	// CommitID the unique id for the commit
	CommitID string `json:"commit_id" yaml:"commit_id" faker:"-"`
	// RepoID the unique id for the repo
	RepoID string `json:"repo_id" yaml:"repo_id" faker:"-"`
	// Filename the filename
	Filename string `json:"filename" yaml:"filename" faker:"-"`
	// Additions the number of additions for the commit file
	Additions int64 `json:"additions" yaml:"additions" faker:"-"`
	// Deletions the number of deletions for the commit file
	Deletions int64 `json:"deletions" yaml:"deletions" faker:"-"`
	// Status the status of the change
	Status string `json:"status" yaml:"status" faker:"-"`
	// Binary indicates if the file was detected to be a binary file
	Binary bool `json:"binary" yaml:"binary" faker:"-"`
	// Language the language that was detected for the file
	Language string `json:"language" yaml:"language" faker:"-"`
	// Excluded if the file was excluded from processing
	Excluded bool `json:"excluded" yaml:"excluded" faker:"-"`
	// ExcludedReason if the file was excluded, the reason
	ExcludedReason string `json:"excluded_reason" yaml:"excluded_reason" faker:"-"`
	// Ordinal the order value for the file in the change set
	Ordinal int64 `json:"ordinal" yaml:"ordinal" faker:"-"`
	// Loc the number of lines in the file
	Loc int64 `json:"loc" yaml:"loc" faker:"-"`
	// Sloc the number of source lines in the file
	Sloc int64 `json:"sloc" yaml:"sloc" faker:"-"`
	// Blanks the number of blank lines in the file
	Blanks int64 `json:"blanks" yaml:"blanks" faker:"-"`
	// Comments the number of comment lines in the file
	Comments int64 `json:"comments" yaml:"comments" faker:"-"`
	// Complexity the complexity value for the file change
	Complexity int64 `json:"complexity" yaml:"complexity" faker:"-"`
	// License the license which was detected for the file
	License string `json:"license" yaml:"license" faker:"-"`
	// LicenseConfidence the license confidence from the detection engine
	LicenseConfidence float64 `json:"license_confidence" yaml:"license_confidence" faker:"-"`
	// Renamed if the file was renamed
	Renamed bool `json:"renamed" yaml:"renamed" faker:"-"`
	// RenamedFrom the original file name
	RenamedFrom string `json:"renamed_from" yaml:"renamed_from" faker:"-"`
	// RenamedTo the final file name
	RenamedTo string `json:"renamed_to" yaml:"renamed_to" faker:"-"`
	// Size the size of the file
	Size int64 `json:"size" yaml:"size" faker:"-"`
	// SkippedReason the reason why the file was skipped
	SkippedReason string `json:"skipped_reason" yaml:"skipped_reason" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*CommitFile)(nil)

func toCommitFileObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toCommitFileObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if o == nil {
		return toCommitFileObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toCommitFileObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toCommitFileObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toCommitFileObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toCommitFileObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toCommitFileObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toCommitFileObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toCommitFileObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toCommitFileObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toCommitFileObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toCommitFileObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case map[string]interface{}:
		return o
	case *map[string]interface{}:
		return v
	case *CommitFile:
		return v.ToMap()
	case CommitFile:
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
			arr = append(arr, toCommitFileObject(av, isavro, false, ""))
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of CommitFile
func (o *CommitFile) String() string {
	return fmt.Sprintf("sourcecode.CommitFile<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *CommitFile) GetTopicName() datamodel.TopicNameType {
	return CommitFileTopic
}

// GetModelName returns the name of the model
func (o *CommitFile) GetModelName() datamodel.ModelNameType {
	return CommitFileModelName
}

func (o *CommitFile) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *CommitFile) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("CommitFile", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetRefID returns the RefID for the object
func (o *CommitFile) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *CommitFile) IsMaterialized() bool {
	return true
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *CommitFile) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return &datamodel.ModelMaterializeConfig{
		KeyName:   "commit_file_id",
		TableName: "sourcecode_commitfile",
	}
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *CommitFile) IsEvented() bool {
	return true
}

// GetTopicConfig returns the topic config object
func (o *CommitFile) GetTopicConfig() *datamodel.ModelTopicConfig {
	duration, err := time.ParseDuration("168h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 168h0m0s. " + err.Error())
	}
	return &datamodel.ModelTopicConfig{
		NumPartitions:     4,
		ReplicationFactor: 1,
		Retention:         duration,
		MaxSize:           5242880,
	}
}

// Clone returns an exact copy of CommitFile
func (o *CommitFile) Clone() *CommitFile {
	c := new(CommitFile)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *CommitFile) Anon() *CommitFile {
	c := new(CommitFile)
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
func (o *CommitFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *CommitFile) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	// make sure that these have values if empty
	o.setDefaults()
	return nil
}

var cachedCodecCommitFile *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *CommitFile) GetAvroCodec() *goavro.Codec {
	if cachedCodecCommitFile == nil {
		c, err := GetCommitFileAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecCommitFile = c
	}
	return cachedCodecCommitFile
}

// ToAvroBinary returns the data as Avro binary data
func (o *CommitFile) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *CommitFile) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two CommitFile objects are equal
func (o *CommitFile) IsEqual(other *CommitFile) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *CommitFile) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	return map[string]interface{}{
		"commit_file_id":     o.GetID(),
		"ref_id":             o.GetRefID(),
		"ref_type":           o.RefType,
		"customer_id":        o.CustomerID,
		"hashcode":           o.Hash(),
		"created_ts":         toCommitFileObject(o.CreatedAt, isavro, false, "long"),
		"commit_id":          toCommitFileObject(o.CommitID, isavro, false, "string"),
		"repo_id":            toCommitFileObject(o.RepoID, isavro, false, "string"),
		"filename":           toCommitFileObject(o.Filename, isavro, false, "string"),
		"additions":          toCommitFileObject(o.Additions, isavro, false, "long"),
		"deletions":          toCommitFileObject(o.Deletions, isavro, false, "long"),
		"status":             toCommitFileObject(o.Status, isavro, false, "string"),
		"binary":             toCommitFileObject(o.Binary, isavro, false, "boolean"),
		"language":           toCommitFileObject(o.Language, isavro, false, "string"),
		"excluded":           toCommitFileObject(o.Excluded, isavro, false, "boolean"),
		"excluded_reason":    toCommitFileObject(o.ExcludedReason, isavro, false, "string"),
		"ordinal":            toCommitFileObject(o.Ordinal, isavro, false, "long"),
		"loc":                toCommitFileObject(o.Loc, isavro, false, "long"),
		"sloc":               toCommitFileObject(o.Sloc, isavro, false, "long"),
		"blanks":             toCommitFileObject(o.Blanks, isavro, false, "long"),
		"comments":           toCommitFileObject(o.Comments, isavro, false, "long"),
		"complexity":         toCommitFileObject(o.Complexity, isavro, false, "long"),
		"license":            toCommitFileObject(o.License, isavro, false, "string"),
		"license_confidence": toCommitFileObject(o.LicenseConfidence, isavro, false, "float"),
		"renamed":            toCommitFileObject(o.Renamed, isavro, false, "boolean"),
		"renamed_from":       toCommitFileObject(o.RenamedFrom, isavro, false, "string"),
		"renamed_to":         toCommitFileObject(o.RenamedTo, isavro, false, "string"),
		"size":               toCommitFileObject(o.Size, isavro, false, "long"),
		"skipped_reason":     toCommitFileObject(o.SkippedReason, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *CommitFile) FromMap(kv map[string]interface{}) {
	if val, ok := kv["commit_file_id"].(string); ok {
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
	if val, ok := kv["commit_id"].(string); ok {
		o.CommitID = val
	} else {
		val := kv["commit_id"]
		if val == nil {
			o.CommitID = ""
		} else {
			o.CommitID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["repo_id"].(string); ok {
		o.RepoID = val
	} else {
		val := kv["repo_id"]
		if val == nil {
			o.RepoID = ""
		} else {
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
			o.Filename = fmt.Sprintf("%v", val)
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
	if val, ok := kv["binary"].(bool); ok {
		o.Binary = val
	} else {
		val := kv["binary"]
		if val == nil {
			o.Binary = number.ToBoolAny(nil)
		} else {
			o.Binary = number.ToBoolAny(val)
		}
	}
	if val, ok := kv["language"].(string); ok {
		o.Language = val
	} else {
		val := kv["language"]
		if val == nil {
			o.Language = ""
		} else {
			o.Language = fmt.Sprintf("%v", val)
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
			o.ExcludedReason = fmt.Sprintf("%v", val)
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
	if val, ok := kv["license"].(string); ok {
		o.License = val
	} else {
		val := kv["license"]
		if val == nil {
			o.License = ""
		} else {
			o.License = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["license_confidence"].(float64); ok {
		o.LicenseConfidence = val
	} else {
		val := kv["license_confidence"]
		if val == nil {
			o.LicenseConfidence = number.ToFloat64Any(nil)
		} else {
			o.LicenseConfidence = number.ToFloat64Any(val)
		}
	}
	if val, ok := kv["renamed"].(bool); ok {
		o.Renamed = val
	} else {
		val := kv["renamed"]
		if val == nil {
			o.Renamed = number.ToBoolAny(nil)
		} else {
			o.Renamed = number.ToBoolAny(val)
		}
	}
	if val, ok := kv["renamed_from"].(string); ok {
		o.RenamedFrom = val
	} else {
		val := kv["renamed_from"]
		if val == nil {
			o.RenamedFrom = ""
		} else {
			o.RenamedFrom = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["renamed_to"].(string); ok {
		o.RenamedTo = val
	} else {
		val := kv["renamed_to"]
		if val == nil {
			o.RenamedTo = ""
		} else {
			o.RenamedTo = fmt.Sprintf("%v", val)
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
	if val, ok := kv["skipped_reason"].(string); ok {
		o.SkippedReason = val
	} else {
		val := kv["skipped_reason"]
		if val == nil {
			o.SkippedReason = ""
		} else {
			o.SkippedReason = fmt.Sprintf("%v", val)
		}
	}
	// make sure that these have values if empty
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *CommitFile) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.GetID())
	args = append(args, o.GetRefID())
	args = append(args, o.RefType)
	args = append(args, o.CustomerID)
	args = append(args, o.CreatedAt)
	args = append(args, o.CommitID)
	args = append(args, o.RepoID)
	args = append(args, o.Filename)
	args = append(args, o.Additions)
	args = append(args, o.Deletions)
	args = append(args, o.Status)
	args = append(args, o.Binary)
	args = append(args, o.Language)
	args = append(args, o.Excluded)
	args = append(args, o.ExcludedReason)
	args = append(args, o.Ordinal)
	args = append(args, o.Loc)
	args = append(args, o.Sloc)
	args = append(args, o.Blanks)
	args = append(args, o.Comments)
	args = append(args, o.Complexity)
	args = append(args, o.License)
	args = append(args, o.LicenseConfidence)
	args = append(args, o.Renamed)
	args = append(args, o.RenamedFrom)
	args = append(args, o.RenamedTo)
	args = append(args, o.Size)
	args = append(args, o.SkippedReason)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetCommitFileAvroSchemaSpec creates the avro schema specification for CommitFile
func GetCommitFileAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "sourcecode",
		"name":         "CommitFile",
		"connect.name": "sourcecode.CommitFile",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "commit_file_id",
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
				"name": "created_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "commit_id",
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
				"name": "additions",
				"type": "long",
			},
			map[string]interface{}{
				"name": "deletions",
				"type": "long",
			},
			map[string]interface{}{
				"name": "status",
				"type": "string",
			},
			map[string]interface{}{
				"name": "binary",
				"type": "boolean",
			},
			map[string]interface{}{
				"name": "language",
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
				"name": "license",
				"type": "string",
			},
			map[string]interface{}{
				"name": "license_confidence",
				"type": "float",
			},
			map[string]interface{}{
				"name": "renamed",
				"type": "boolean",
			},
			map[string]interface{}{
				"name": "renamed_from",
				"type": "string",
			},
			map[string]interface{}{
				"name": "renamed_to",
				"type": "string",
			},
			map[string]interface{}{
				"name": "size",
				"type": "long",
			},
			map[string]interface{}{
				"name": "skipped_reason",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetCommitFileAvroSchema creates the avro schema for CommitFile
func GetCommitFileAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetCommitFileAvroSchemaSpec())
}

// TransformCommitFileFunc is a function for transforming CommitFile during processing
type TransformCommitFileFunc func(input *CommitFile) (*CommitFile, error)

// NewCommitFilePipe creates a pipe for processing CommitFile items
func NewCommitFilePipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformCommitFileFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewCommitFileInputStream(input, errors)
	var stream chan CommitFile
	if len(transforms) > 0 {
		stream = make(chan CommitFile, 1000)
	} else {
		stream = inch
	}
	outdone := NewCommitFileOutputStream(output, stream, errors)
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

// NewCommitFileInputStreamDir creates a channel for reading CommitFile as JSON newlines from a directory of files
func NewCommitFileInputStreamDir(dir string, errors chan<- error, transforms ...TransformCommitFileFunc) (chan CommitFile, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/sourcecode/commit_file\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan CommitFile)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for commit_file")
		ch := make(chan CommitFile)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewCommitFileInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan CommitFile)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewCommitFileInputStreamFile creates an channel for reading CommitFile as JSON newlines from filename
func NewCommitFileInputStreamFile(filename string, errors chan<- error, transforms ...TransformCommitFileFunc) (chan CommitFile, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan CommitFile)
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
			ch := make(chan CommitFile)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewCommitFileInputStream(f, errors, transforms...)
}

// NewCommitFileInputStream creates an channel for reading CommitFile as JSON newlines from stream
func NewCommitFileInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformCommitFileFunc) (chan CommitFile, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan CommitFile, 1000)
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
			var item CommitFile
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

// NewCommitFileOutputStreamDir will output json newlines from channel and save in dir
func NewCommitFileOutputStreamDir(dir string, ch chan CommitFile, errors chan<- error, transforms ...TransformCommitFileFunc) <-chan bool {
	fp := filepath.Join(dir, "/sourcecode/commit_file\\.json(\\.gz)?$")
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
	return NewCommitFileOutputStream(gz, ch, errors, transforms...)
}

// NewCommitFileOutputStream will output json newlines from channel to the stream
func NewCommitFileOutputStream(stream io.WriteCloser, ch chan CommitFile, errors chan<- error, transforms ...TransformCommitFileFunc) <-chan bool {
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

// CommitFileSendEvent is an event detail for sending data
type CommitFileSendEvent struct {
	CommitFile *CommitFile
	headers    map[string]string
	time       time.Time
	key        string
}

var _ datamodel.ModelSendEvent = (*CommitFileSendEvent)(nil)

// Key is the key to use for the message
func (e *CommitFileSendEvent) Key() string {
	if e.key == "" {
		return e.CommitFile.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *CommitFileSendEvent) Object() datamodel.Model {
	return e.CommitFile
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *CommitFileSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *CommitFileSendEvent) Timestamp() time.Time {
	return e.time
}

// CommitFileSendEventOpts is a function handler for setting opts
type CommitFileSendEventOpts func(o *CommitFileSendEvent)

// WithCommitFileSendEventKey sets the key value to a value different than the object ID
func WithCommitFileSendEventKey(key string) CommitFileSendEventOpts {
	return func(o *CommitFileSendEvent) {
		o.key = key
	}
}

// WithCommitFileSendEventTimestamp sets the timestamp value
func WithCommitFileSendEventTimestamp(tv time.Time) CommitFileSendEventOpts {
	return func(o *CommitFileSendEvent) {
		o.time = tv
	}
}

// WithCommitFileSendEventHeader sets the timestamp value
func WithCommitFileSendEventHeader(key, value string) CommitFileSendEventOpts {
	return func(o *CommitFileSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewCommitFileSendEvent returns a new CommitFileSendEvent instance
func NewCommitFileSendEvent(o *CommitFile, opts ...CommitFileSendEventOpts) *CommitFileSendEvent {
	res := &CommitFileSendEvent{
		CommitFile: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewCommitFileProducer will stream data from the channel
func NewCommitFileProducer(producer event.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		ctx := context.Background()
		for item := range ch {
			if object, ok := item.Object().(*CommitFile); ok {
				binary, codec, err := object.ToAvroBinary()
				if err != nil {
					errors <- fmt.Errorf("error encoding %s to avro binary data. %v", object.String(), err)
					return
				}
				headers := map[string]string{
					"customer_id": object.CustomerID,
				}
				for k, v := range item.Headers() {
					headers[k] = v
				}
				msg := event.Message{
					Key:     item.Key(),
					Value:   binary,
					Codec:   codec,
					Headers: headers,
				}
				if err := producer.Send(ctx, msg); err != nil {
					errors <- fmt.Errorf("error sending %s. %v", object.String(), err)
				}
			} else {
				errors <- fmt.Errorf("invalid event received. expected an object of type sourcecode.CommitFile but received on of type %v", reflect.TypeOf(item.Object()))
			}
		}
	}()
	return done
}

// NewCommitFileConsumer will stream data from the topic into the provided channel
func NewCommitFileConsumer(consumer event.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) {
	consumer.Consume(event.ConsumerCallback{
		OnDataReceived: func(msg event.Message) error {
			var object CommitFile
			if err := json.Unmarshal(msg.Value, &object); err != nil {
				return fmt.Errorf("error unmarshaling json data into sourcecode.CommitFile: %s", err)
			}
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &CommitFileReceiveEvent{&object, msg}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
	})
}

// CommitFileReceiveEvent is an event detail for receiving data
type CommitFileReceiveEvent struct {
	CommitFile *CommitFile
	message    event.Message
}

var _ datamodel.ModelReceiveEvent = (*CommitFileReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *CommitFileReceiveEvent) Object() datamodel.Model {
	return e.CommitFile
}

// Message returns the underlying message data for the event
func (e *CommitFileReceiveEvent) Message() event.Message {
	return e.message
}

// CommitFileProducer implements the datamodel.ModelEventProducer
type CommitFileProducer struct {
	ch   chan datamodel.ModelSendEvent
	done <-chan bool
}

var _ datamodel.ModelEventProducer = (*CommitFileProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *CommitFileProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *CommitFileProducer) Close() error {
	close(p.ch)
	<-p.done
	return nil
}

// NewProducerChannel returns a channel which can be used for producing Model events
func (o *CommitFile) NewProducerChannel(producer event.Producer, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent)
	return &CommitFileProducer{
		ch:   ch,
		done: NewCommitFileProducer(producer, ch, errors),
	}
}

// NewCommitFileProducerChannel returns a channel which can be used for producing Model events
func NewCommitFileProducerChannel(producer event.Producer, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent)
	return &CommitFileProducer{
		ch:   ch,
		done: NewCommitFileProducer(producer, ch, errors),
	}
}

// CommitFileConsumer implements the datamodel.ModelEventConsumer
type CommitFileConsumer struct {
	ch chan datamodel.ModelReceiveEvent
}

var _ datamodel.ModelEventConsumer = (*CommitFileConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *CommitFileConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *CommitFileConsumer) Close() error {
	close(c.ch)
	return nil
}

// NewConsumerChannel returns a consumer channel which can be used to consume Model events
func (o *CommitFile) NewConsumerChannel(consumer event.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	NewCommitFileConsumer(consumer, ch, errors)
	return &CommitFileConsumer{
		ch: ch,
	}
}

// NewCommitFileConsumerChannel returns a consumer channel which can be used to consume Model events
func NewCommitFileConsumerChannel(consumer event.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	NewCommitFileConsumer(consumer, ch, errors)
	return &CommitFileConsumer{
		ch: ch,
	}
}
