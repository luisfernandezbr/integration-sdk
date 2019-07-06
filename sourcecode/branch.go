// DO NOT EDIT -- generated code

// Package sourcecode - the system which contains source code
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
)

const (
	// BranchTopic is the default topic name
	BranchTopic datamodel.TopicNameType = "sourcecode_Branch_topic"

	// BranchStream is the default stream name
	BranchStream datamodel.TopicNameType = "sourcecode_Branch_stream"

	// BranchTable is the default table name
	BranchTable datamodel.TopicNameType = "sourcecode_Branch"

	// BranchModelName is the model name
	BranchModelName datamodel.ModelNameType = "sourcecode.Branch"
)

const (
	// BranchAheadDefaultCountColumn is the ahead_default_count column name
	BranchAheadDefaultCountColumn = "ahead_default_count"
	// BranchBehindDefaultCountColumn is the behind_default_count column name
	BranchBehindDefaultCountColumn = "behind_default_count"
	// BranchBranchedFromCommitsColumn is the branched_from_commits column name
	BranchBranchedFromCommitsColumn = "branched_from_commits"
	// BranchCommitsColumn is the commits column name
	BranchCommitsColumn = "commits"
	// BranchCustomerIDColumn is the customer_id column name
	BranchCustomerIDColumn = "customer_id"
	// BranchDefaultColumn is the default column name
	BranchDefaultColumn = "default"
	// BranchIDColumn is the id column name
	BranchIDColumn = "id"
	// BranchMergeCommitColumn is the merge_commit column name
	BranchMergeCommitColumn = "merge_commit"
	// BranchMergedColumn is the merged column name
	BranchMergedColumn = "merged"
	// BranchNameColumn is the name column name
	BranchNameColumn = "name"
	// BranchRefIDColumn is the ref_id column name
	BranchRefIDColumn = "ref_id"
	// BranchRefTypeColumn is the ref_type column name
	BranchRefTypeColumn = "ref_type"
	// BranchRepoIDColumn is the repo_id column name
	BranchRepoIDColumn = "repo_id"
)

// Branch git branches
type Branch struct {
	// AheadDefaultCount ahead default count
	AheadDefaultCount int64 `json:"ahead_default_count" bson:"ahead_default_count" yaml:"ahead_default_count" faker:"-"`
	// BehindDefaultCount behind default count
	BehindDefaultCount int64 `json:"behind_default_count" bson:"behind_default_count" yaml:"behind_default_count" faker:"-"`
	// BranchedFromCommits branched from commits
	BranchedFromCommits []string `json:"branched_from_commits" bson:"branched_from_commits" yaml:"branched_from_commits" faker:"-"`
	// Commits list of commits on this branch
	Commits []string `json:"commits" bson:"commits" yaml:"commits" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Default wether is the default branch or not
	Default bool `json:"default" bson:"default" yaml:"default" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// MergeCommit commit from the merge
	MergeCommit string `json:"merge_commit" bson:"merge_commit" yaml:"merge_commit" faker:"-"`
	// Merged wether it has been merged
	Merged bool `json:"merged" bson:"merged" yaml:"merged" faker:"-"`
	// Name name of the branch
	Name string `json:"name" bson:"name" yaml:"name" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RepoID the unique id for the repo
	RepoID string `json:"repo_id" bson:"repo_id" yaml:"repo_id" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Branch)(nil)

func toBranchObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toBranchObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if o == nil {
		return toBranchObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toBranchObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toBranchObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toBranchObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toBranchObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toBranchObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toBranchObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toBranchObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toBranchObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toBranchObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toBranchObjectNil(isavro, isoptional)
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
	case *Branch:
		return v.ToMap()
	case Branch:
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
			arr = append(arr, toBranchObject(av, isavro, false, ""))
		}
		return arr

	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of Branch
func (o *Branch) String() string {
	return fmt.Sprintf("sourcecode.Branch<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Branch) GetTopicName() datamodel.TopicNameType {
	return BranchTopic
}

// GetModelName returns the name of the model
func (o *Branch) GetModelName() datamodel.ModelNameType {
	return BranchModelName
}

func (o *Branch) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Branch) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Branch", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Branch) GetTopicKey() string {
	var i interface{} = o.RepoID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Branch) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *Branch) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Branch) IsMaterialized() bool {
	return true
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Branch) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	idletime, err := time.ParseDuration("1s")
	if err != nil {
		panic(err)
	}
	return &datamodel.ModelMaterializeConfig{
		KeyName:   "id",
		TableName: "sourcecode_branch",
		BatchSize: 5000,
		IdleTime:  idletime,
	}
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Branch) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *Branch) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = BranchModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Branch) GetTopicConfig() *datamodel.ModelTopicConfig {
	duration, err := time.ParseDuration("168h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 168h0m0s. " + err.Error())
	}
	return &datamodel.ModelTopicConfig{
		Key:               "repo_id",
		Timestamp:         "",
		NumPartitions:     8,
		ReplicationFactor: 3,
		Retention:         duration,
		MaxSize:           5242880,
	}
}

// GetStateKey returns a key for use in state store
func (o *Branch) GetStateKey() string {
	key := "repo_id"
	return fmt.Sprintf("%s_%s", key, o.GetID())
}

// GetCustomerID will return the customer_id
func (o *Branch) GetCustomerID() string {
	return o.CustomerID
}

// Clone returns an exact copy of Branch
func (o *Branch) Clone() datamodel.Model {
	c := new(Branch)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Branch) Anon() datamodel.Model {
	c := new(Branch)
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
func (o *Branch) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Branch) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

var cachedCodecBranch *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *Branch) GetAvroCodec() *goavro.Codec {
	if cachedCodecBranch == nil {
		c, err := GetBranchAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecBranch = c
	}
	return cachedCodecBranch
}

// ToAvroBinary returns the data as Avro binary data
func (o *Branch) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *Branch) FromAvroBinary(value []byte) error {
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
func (o *Branch) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Branch objects are equal
func (o *Branch) IsEqual(other *Branch) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Branch) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
		if o.BranchedFromCommits == nil {
			o.BranchedFromCommits = make([]string, 0)
		}
		if o.Commits == nil {
			o.Commits = make([]string, 0)
		}
	}
	o.setDefaults()
	return map[string]interface{}{
		"ahead_default_count":   toBranchObject(o.AheadDefaultCount, isavro, false, "long"),
		"behind_default_count":  toBranchObject(o.BehindDefaultCount, isavro, false, "long"),
		"branched_from_commits": toBranchObject(o.BranchedFromCommits, isavro, false, "branched_from_commits"),
		"commits":               toBranchObject(o.Commits, isavro, false, "commits"),
		"customer_id":           toBranchObject(o.CustomerID, isavro, false, "string"),
		"default":               toBranchObject(o.Default, isavro, false, "boolean"),
		"id":                    toBranchObject(o.ID, isavro, false, "string"),
		"merge_commit":          toBranchObject(o.MergeCommit, isavro, false, "string"),
		"merged":                toBranchObject(o.Merged, isavro, false, "boolean"),
		"name":                  toBranchObject(o.Name, isavro, false, "string"),
		"ref_id":                toBranchObject(o.RefID, isavro, false, "string"),
		"ref_type":              toBranchObject(o.RefType, isavro, false, "string"),
		"repo_id":               toBranchObject(o.RepoID, isavro, false, "string"),
		"hashcode":              toBranchObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *Branch) FromMap(kv map[string]interface{}) {
	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["ahead_default_count"].(int64); ok {
		o.AheadDefaultCount = val
	} else {
		val := kv["ahead_default_count"]
		if val == nil {
			o.AheadDefaultCount = number.ToInt64Any(nil)
		} else {
			if tv, ok := val.(time.Time); ok {
				val = datetime.TimeToEpoch(tv)
			}
			o.AheadDefaultCount = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["behind_default_count"].(int64); ok {
		o.BehindDefaultCount = val
	} else {
		val := kv["behind_default_count"]
		if val == nil {
			o.BehindDefaultCount = number.ToInt64Any(nil)
		} else {
			if tv, ok := val.(time.Time); ok {
				val = datetime.TimeToEpoch(tv)
			}
			o.BehindDefaultCount = number.ToInt64Any(val)
		}
	}
	if val := kv["branched_from_commits"]; val != nil {
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
							panic("unsupported type for branched_from_commits field entry: " + reflect.TypeOf(ae).String())
						}
						na = append(na, av)
					}
				}
			} else if s, ok := val.(string); ok {
				for _, sv := range strings.Split(s, ",") {
					na = append(na, strings.TrimSpace(sv))
				}
			} else {
				fmt.Println(reflect.TypeOf(val).String())
				panic("unsupported type for branched_from_commits field")
			}
		}
		o.BranchedFromCommits = na
	} else {
		o.BranchedFromCommits = []string{}
	}
	if o.BranchedFromCommits == nil {
		o.BranchedFromCommits = make([]string, 0)
	}
	if val := kv["commits"]; val != nil {
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
							panic("unsupported type for commits field entry: " + reflect.TypeOf(ae).String())
						}
						na = append(na, av)
					}
				}
			} else if s, ok := val.(string); ok {
				for _, sv := range strings.Split(s, ",") {
					na = append(na, strings.TrimSpace(sv))
				}
			} else {
				fmt.Println(reflect.TypeOf(val).String())
				panic("unsupported type for commits field")
			}
		}
		o.Commits = na
	} else {
		o.Commits = []string{}
	}
	if o.Commits == nil {
		o.Commits = make([]string, 0)
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
	if val, ok := kv["default"].(bool); ok {
		o.Default = val
	} else {
		val := kv["default"]
		if val == nil {
			o.Default = number.ToBoolAny(nil)
		} else {
			o.Default = number.ToBoolAny(val)
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
	if val, ok := kv["merge_commit"].(string); ok {
		o.MergeCommit = val
	} else {
		val := kv["merge_commit"]
		if val == nil {
			o.MergeCommit = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.MergeCommit = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["merged"].(bool); ok {
		o.Merged = val
	} else {
		val := kv["merged"]
		if val == nil {
			o.Merged = number.ToBoolAny(nil)
		} else {
			o.Merged = number.ToBoolAny(val)
		}
	}
	if val, ok := kv["name"].(string); ok {
		o.Name = val
	} else {
		val := kv["name"]
		if val == nil {
			o.Name = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Name = fmt.Sprintf("%v", val)
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
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *Branch) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.AheadDefaultCount)
	args = append(args, o.BehindDefaultCount)
	args = append(args, o.BranchedFromCommits)
	args = append(args, o.Commits)
	args = append(args, o.CustomerID)
	args = append(args, o.Default)
	args = append(args, o.ID)
	args = append(args, o.MergeCommit)
	args = append(args, o.Merged)
	args = append(args, o.Name)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RepoID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetBranchAvroSchemaSpec creates the avro schema specification for Branch
func GetBranchAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "sourcecode",
		"name":      "Branch",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "hashcode",
				"type": "string",
			},
			map[string]interface{}{
				"name": "ahead_default_count",
				"type": "long",
			},
			map[string]interface{}{
				"name": "behind_default_count",
				"type": "long",
			},
			map[string]interface{}{
				"name": "branched_from_commits",
				"type": map[string]interface{}{"type": "array", "name": "branched_from_commits", "items": "string"},
			},
			map[string]interface{}{
				"name": "commits",
				"type": map[string]interface{}{"type": "array", "name": "commits", "items": "string"},
			},
			map[string]interface{}{
				"name": "customer_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "default",
				"type": "boolean",
			},
			map[string]interface{}{
				"name": "id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "merge_commit",
				"type": "string",
			},
			map[string]interface{}{
				"name": "merged",
				"type": "boolean",
			},
			map[string]interface{}{
				"name": "name",
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
				"name": "repo_id",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetBranchAvroSchema creates the avro schema for Branch
func GetBranchAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetBranchAvroSchemaSpec())
}

// TransformBranchFunc is a function for transforming Branch during processing
type TransformBranchFunc func(input *Branch) (*Branch, error)

// NewBranchPipe creates a pipe for processing Branch items
func NewBranchPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformBranchFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewBranchInputStream(input, errors)
	var stream chan Branch
	if len(transforms) > 0 {
		stream = make(chan Branch, 1000)
	} else {
		stream = inch
	}
	outdone := NewBranchOutputStream(output, stream, errors)
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

// NewBranchInputStreamDir creates a channel for reading Branch as JSON newlines from a directory of files
func NewBranchInputStreamDir(dir string, errors chan<- error, transforms ...TransformBranchFunc) (chan Branch, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/sourcecode/branch\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan Branch)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for branch")
		ch := make(chan Branch)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewBranchInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan Branch)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewBranchInputStreamFile creates an channel for reading Branch as JSON newlines from filename
func NewBranchInputStreamFile(filename string, errors chan<- error, transforms ...TransformBranchFunc) (chan Branch, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan Branch)
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
			ch := make(chan Branch)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewBranchInputStream(f, errors, transforms...)
}

// NewBranchInputStream creates an channel for reading Branch as JSON newlines from stream
func NewBranchInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformBranchFunc) (chan Branch, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan Branch, 1000)
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
			var item Branch
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

// NewBranchOutputStreamDir will output json newlines from channel and save in dir
func NewBranchOutputStreamDir(dir string, ch chan Branch, errors chan<- error, transforms ...TransformBranchFunc) <-chan bool {
	fp := filepath.Join(dir, "/sourcecode/branch\\.json(\\.gz)?$")
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
	return NewBranchOutputStream(gz, ch, errors, transforms...)
}

// NewBranchOutputStream will output json newlines from channel to the stream
func NewBranchOutputStream(stream io.WriteCloser, ch chan Branch, errors chan<- error, transforms ...TransformBranchFunc) <-chan bool {
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

// BranchSendEvent is an event detail for sending data
type BranchSendEvent struct {
	Branch  *Branch
	headers map[string]string
	time    time.Time
	key     string
}

var _ datamodel.ModelSendEvent = (*BranchSendEvent)(nil)

// Key is the key to use for the message
func (e *BranchSendEvent) Key() string {
	if e.key == "" {
		return e.Branch.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *BranchSendEvent) Object() datamodel.Model {
	return e.Branch
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *BranchSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *BranchSendEvent) Timestamp() time.Time {
	return e.time
}

// BranchSendEventOpts is a function handler for setting opts
type BranchSendEventOpts func(o *BranchSendEvent)

// WithBranchSendEventKey sets the key value to a value different than the object ID
func WithBranchSendEventKey(key string) BranchSendEventOpts {
	return func(o *BranchSendEvent) {
		o.key = key
	}
}

// WithBranchSendEventTimestamp sets the timestamp value
func WithBranchSendEventTimestamp(tv time.Time) BranchSendEventOpts {
	return func(o *BranchSendEvent) {
		o.time = tv
	}
}

// WithBranchSendEventHeader sets the timestamp value
func WithBranchSendEventHeader(key, value string) BranchSendEventOpts {
	return func(o *BranchSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewBranchSendEvent returns a new BranchSendEvent instance
func NewBranchSendEvent(o *Branch, opts ...BranchSendEventOpts) *BranchSendEvent {
	res := &BranchSendEvent{
		Branch: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewBranchProducer will stream data from the channel
func NewBranchProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
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
				if object, ok := item.Object().(*Branch); ok {
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
					errors <- fmt.Errorf("invalid event received. expected an object of type sourcecode.Branch but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewBranchConsumer will stream data from the topic into the provided channel
func NewBranchConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object Branch
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into sourcecode.Branch: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avri data into sourcecode.Branch: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for sourcecode.Branch")
			}
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &BranchReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object Branch
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &BranchReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// BranchReceiveEvent is an event detail for receiving data
type BranchReceiveEvent struct {
	Branch  *Branch
	message eventing.Message
	eof     bool
}

var _ datamodel.ModelReceiveEvent = (*BranchReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *BranchReceiveEvent) Object() datamodel.Model {
	return e.Branch
}

// Message returns the underlying message data for the event
func (e *BranchReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *BranchReceiveEvent) EOF() bool {
	return e.eof
}

// BranchProducer implements the datamodel.ModelEventProducer
type BranchProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*BranchProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *BranchProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *BranchProducer) Close() error {
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
func (o *Branch) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *Branch) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &BranchProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewBranchProducer(newctx, producer, ch, errors, empty),
	}
}

// NewBranchProducerChannel returns a channel which can be used for producing Model events
func NewBranchProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewBranchProducerChannelSize(producer, 0, errors)
}

// NewBranchProducerChannelSize returns a channel which can be used for producing Model events
func NewBranchProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &BranchProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewBranchProducer(newctx, producer, ch, errors, empty),
	}
}

// BranchConsumer implements the datamodel.ModelEventConsumer
type BranchConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*BranchConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *BranchConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *BranchConsumer) Close() error {
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
func (o *Branch) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &BranchConsumer{
		ch:       ch,
		callback: NewBranchConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewBranchConsumerChannel returns a consumer channel which can be used to consume Model events
func NewBranchConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &BranchConsumer{
		ch:       ch,
		callback: NewBranchConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
