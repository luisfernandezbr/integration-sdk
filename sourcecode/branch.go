// DO NOT EDIT -- generated code

// Package sourcecode - the system which contains source code
package sourcecode

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/bxcodec/faker"
	"github.com/linkedin/goavro"
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/eventing"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
	"github.com/pinpt/go-common/slice"
	pstrings "github.com/pinpt/go-common/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// BranchTopic is the default topic name
	BranchTopic datamodel.TopicNameType = "sourcecode_Branch_topic"

	// BranchStream is the default stream name
	BranchStream datamodel.TopicNameType = "sourcecode_Branch_stream"

	// BranchTable is the default table name
	BranchTable datamodel.TopicNameType = "sourcecode_branch"

	// BranchModelName is the model name
	BranchModelName datamodel.ModelNameType = "sourcecode.Branch"
)

const (
	// BranchAheadDefaultCountColumn is the ahead_default_count column name
	BranchAheadDefaultCountColumn = "ahead_default_count"
	// BranchBehindDefaultCountColumn is the behind_default_count column name
	BranchBehindDefaultCountColumn = "behind_default_count"
	// BranchBranchedFromCommitIdsColumn is the branched_from_commit_ids column name
	BranchBranchedFromCommitIdsColumn = "branched_from_commit_ids"
	// BranchBranchedFromCommitShasColumn is the branched_from_commit_shas column name
	BranchBranchedFromCommitShasColumn = "branched_from_commit_shas"
	// BranchCommitIdsColumn is the commit_ids column name
	BranchCommitIdsColumn = "commit_ids"
	// BranchCommitShasColumn is the commit_shas column name
	BranchCommitShasColumn = "commit_shas"
	// BranchCustomerIDColumn is the customer_id column name
	BranchCustomerIDColumn = "customer_id"
	// BranchDefaultColumn is the default column name
	BranchDefaultColumn = "default"
	// BranchFirstBranchedFromCommitIDColumn is the first_branched_from_commit_id column name
	BranchFirstBranchedFromCommitIDColumn = "first_branched_from_commit_id"
	// BranchFirstBranchedFromCommitShaColumn is the first_branched_from_commit_sha column name
	BranchFirstBranchedFromCommitShaColumn = "first_branched_from_commit_sha"
	// BranchFirstCommitIDColumn is the first_commit_id column name
	BranchFirstCommitIDColumn = "first_commit_id"
	// BranchFirstCommitShaColumn is the first_commit_sha column name
	BranchFirstCommitShaColumn = "first_commit_sha"
	// BranchIDColumn is the id column name
	BranchIDColumn = "id"
	// BranchMergeCommitIDColumn is the merge_commit_id column name
	BranchMergeCommitIDColumn = "merge_commit_id"
	// BranchMergeCommitShaColumn is the merge_commit_sha column name
	BranchMergeCommitShaColumn = "merge_commit_sha"
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
	// BranchUpdatedAtColumn is the updated_ts column name
	BranchUpdatedAtColumn = "updated_ts"
	// BranchURLColumn is the url column name
	BranchURLColumn = "url"
)

// Branch git branches
type Branch struct {
	// AheadDefaultCount the number of commits that this branch is ahead of the default branch
	AheadDefaultCount int64 `json:"ahead_default_count" bson:"ahead_default_count" yaml:"ahead_default_count" faker:"-"`
	// BehindDefaultCount the number of commits that this branch is behind from the default branch
	BehindDefaultCount int64 `json:"behind_default_count" bson:"behind_default_count" yaml:"behind_default_count" faker:"-"`
	// BranchedFromCommitIds the commit ids from which the branch was created
	BranchedFromCommitIds []string `json:"branched_from_commit_ids" bson:"branched_from_commit_ids" yaml:"branched_from_commit_ids" faker:"-"`
	// BranchedFromCommitShas the commit shas from which the branch was created
	BranchedFromCommitShas []string `json:"branched_from_commit_shas" bson:"branched_from_commit_shas" yaml:"branched_from_commit_shas" faker:"-"`
	// CommitIds list of commit ids contained on this branch
	CommitIds []string `json:"commit_ids" bson:"commit_ids" yaml:"commit_ids" faker:"-"`
	// CommitShas list of commit shas contained on this branch
	CommitShas []string `json:"commit_shas" bson:"commit_shas" yaml:"commit_shas" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Default true if the branch is the default branch
	Default bool `json:"default" bson:"default" yaml:"default" faker:"-"`
	// FirstBranchedFromCommitID the first commit id from which the branch was created
	FirstBranchedFromCommitID string `json:"first_branched_from_commit_id" bson:"first_branched_from_commit_id" yaml:"first_branched_from_commit_id" faker:"-"`
	// FirstBranchedFromCommitSha the first commit sha from which the branch was created
	FirstBranchedFromCommitSha string `json:"first_branched_from_commit_sha" bson:"first_branched_from_commit_sha" yaml:"first_branched_from_commit_sha" faker:"-"`
	// FirstCommitID the first commit id on the branch
	FirstCommitID string `json:"first_commit_id" bson:"first_commit_id" yaml:"first_commit_id" faker:"-"`
	// FirstCommitSha the first commit sha on the branch
	FirstCommitSha string `json:"first_commit_sha" bson:"first_commit_sha" yaml:"first_commit_sha" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// MergeCommitID commit id in which the branch was merged
	MergeCommitID string `json:"merge_commit_id" bson:"merge_commit_id" yaml:"merge_commit_id" faker:"-"`
	// MergeCommitSha commit sha in which the branch was merged
	MergeCommitSha string `json:"merge_commit_sha" bson:"merge_commit_sha" yaml:"merge_commit_sha" faker:"-"`
	// Merged true if the branch has been merged into the default branch
	Merged bool `json:"merged" bson:"merged" yaml:"merged" faker:"-"`
	// Name name of the branch
	Name string `json:"name" bson:"name" yaml:"name" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RepoID the unique id for the repo
	RepoID string `json:"repo_id" bson:"repo_id" yaml:"repo_id" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// URL the URL to the source system for this branch
	URL string `json:"url" bson:"url" yaml:"url" faker:"url"`
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
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *Branch:
		return v.ToMap(isavro)

	default:
		return o
	}
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

// GetStreamName returns the name of the stream
func (o *Branch) GetStreamName() string {
	return BranchStream.String()
}

// GetTableName returns the name of the table
func (o *Branch) GetTableName() string {
	return BranchTable.String()
}

// NewBranchID provides a template for generating an ID field for Branch
func NewBranchID(refType string, RepoID string, customerID string, Name string, FirstCommitID string) string {
	return hash.Values(refType, RepoID, customerID, Name, FirstCommitID)
}

func (o *Branch) setDefaults(frommap bool) {
	if o.BranchedFromCommitIds == nil {
		o.BranchedFromCommitIds = make([]string, 0)
	}
	if o.BranchedFromCommitShas == nil {
		o.BranchedFromCommitShas = make([]string, 0)
	}
	if o.CommitIds == nil {
		o.CommitIds = make([]string, 0)
	}
	if o.CommitShas == nil {
		o.CommitShas = make([]string, 0)
	}

	if o.ID == "" {
		o.ID = hash.Values(o.RefType, o.RepoID, o.CustomerID, o.Name, o.FirstCommitID)
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Branch) GetID() string {
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
	panic("not sure how to handle the date time format for Branch")
}

// GetRefID returns the RefID for the object
func (o *Branch) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Branch) IsMaterialized() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Branch) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
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
	retention, err := time.ParseDuration("87360h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 87360h0m0s. " + err.Error())
	}

	ttl, err := time.ParseDuration("0s")
	if err != nil {
		ttl = 0
	}
	if ttl == 0 && retention != 0 {
		ttl = retention // they should be the same if not set
	}
	return &datamodel.ModelTopicConfig{
		Key:               "repo_id",
		Timestamp:         "updated_ts",
		NumPartitions:     8,
		CleanupPolicy:     datamodel.CleanupPolicy("compact"),
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
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

// MarshalBinary returns the bytes for marshaling to binary
func (o *Branch) MarshalBinary() ([]byte, error) {
	return o.MarshalJSON()
}

func (o *Branch) UnmarshalBinary(data []byte) error {
	return o.UnmarshalJSON(data)
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
	if idstr, ok := kv["id"].(string); ok {
		o.ID = idstr
	}
	return nil
}

var cachedCodecBranch *goavro.Codec
var cachedCodecBranchLock sync.Mutex

// GetAvroCodec returns the avro codec for this model
func (o *Branch) GetAvroCodec() *goavro.Codec {
	cachedCodecBranchLock.Lock()
	if cachedCodecBranch == nil {
		c, err := GetBranchAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecBranch = c
	}
	cachedCodecBranchLock.Unlock()
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
	o.Hash()
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
		if o.BranchedFromCommitIds == nil {
			o.BranchedFromCommitIds = make([]string, 0)
		}
		if o.BranchedFromCommitShas == nil {
			o.BranchedFromCommitShas = make([]string, 0)
		}
		if o.CommitIds == nil {
			o.CommitIds = make([]string, 0)
		}
		if o.CommitShas == nil {
			o.CommitShas = make([]string, 0)
		}
	}
	o.setDefaults(false)
	return map[string]interface{}{
		"ahead_default_count":            toBranchObject(o.AheadDefaultCount, isavro, false, "long"),
		"behind_default_count":           toBranchObject(o.BehindDefaultCount, isavro, false, "long"),
		"branched_from_commit_ids":       toBranchObject(o.BranchedFromCommitIds, isavro, false, "branched_from_commit_ids"),
		"branched_from_commit_shas":      toBranchObject(o.BranchedFromCommitShas, isavro, false, "branched_from_commit_shas"),
		"commit_ids":                     toBranchObject(o.CommitIds, isavro, false, "commit_ids"),
		"commit_shas":                    toBranchObject(o.CommitShas, isavro, false, "commit_shas"),
		"customer_id":                    toBranchObject(o.CustomerID, isavro, false, "string"),
		"default":                        toBranchObject(o.Default, isavro, false, "boolean"),
		"first_branched_from_commit_id":  toBranchObject(o.FirstBranchedFromCommitID, isavro, false, "string"),
		"first_branched_from_commit_sha": toBranchObject(o.FirstBranchedFromCommitSha, isavro, false, "string"),
		"first_commit_id":                toBranchObject(o.FirstCommitID, isavro, false, "string"),
		"first_commit_sha":               toBranchObject(o.FirstCommitSha, isavro, false, "string"),
		"id":                             toBranchObject(o.ID, isavro, false, "string"),
		"merge_commit_id":                toBranchObject(o.MergeCommitID, isavro, false, "string"),
		"merge_commit_sha":               toBranchObject(o.MergeCommitSha, isavro, false, "string"),
		"merged":                         toBranchObject(o.Merged, isavro, false, "boolean"),
		"name":                           toBranchObject(o.Name, isavro, false, "string"),
		"ref_id":                         toBranchObject(o.RefID, isavro, false, "string"),
		"ref_type":                       toBranchObject(o.RefType, isavro, false, "string"),
		"repo_id":                        toBranchObject(o.RepoID, isavro, false, "string"),
		"updated_ts":                     toBranchObject(o.UpdatedAt, isavro, false, "long"),
		"url":                            toBranchObject(o.URL, isavro, false, "string"),
		"hashcode":                       toBranchObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *Branch) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["ahead_default_count"].(int64); ok {
		o.AheadDefaultCount = val
	} else {
		if val, ok := kv["ahead_default_count"]; ok {
			if val == nil {
				o.AheadDefaultCount = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.AheadDefaultCount = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["behind_default_count"].(int64); ok {
		o.BehindDefaultCount = val
	} else {
		if val, ok := kv["behind_default_count"]; ok {
			if val == nil {
				o.BehindDefaultCount = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.BehindDefaultCount = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["branched_from_commit_ids"]; ok {
		if val != nil {
			na := make([]string, 0)
			if a, ok := val.([]string); ok {
				na = append(na, a...)
			} else {
				if a, ok := val.([]interface{}); ok {
					for _, ae := range a {
						if av, ok := ae.(string); ok {
							na = append(na, av)
						} else {
							if badMap, ok := ae.(map[interface{}]interface{}); ok {
								ae = slice.ConvertToStringToInterface(badMap)
							}
							b, _ := json.Marshal(ae)
							var av string
							if err := json.Unmarshal(b, &av); err != nil {
								panic("unsupported type for branched_from_commit_ids field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for branched_from_commit_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for branched_from_commit_ids field")
				}
			}
			o.BranchedFromCommitIds = na
		}
	}
	if o.BranchedFromCommitIds == nil {
		o.BranchedFromCommitIds = make([]string, 0)
	}

	if val, ok := kv["branched_from_commit_shas"]; ok {
		if val != nil {
			na := make([]string, 0)
			if a, ok := val.([]string); ok {
				na = append(na, a...)
			} else {
				if a, ok := val.([]interface{}); ok {
					for _, ae := range a {
						if av, ok := ae.(string); ok {
							na = append(na, av)
						} else {
							if badMap, ok := ae.(map[interface{}]interface{}); ok {
								ae = slice.ConvertToStringToInterface(badMap)
							}
							b, _ := json.Marshal(ae)
							var av string
							if err := json.Unmarshal(b, &av); err != nil {
								panic("unsupported type for branched_from_commit_shas field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for branched_from_commit_shas field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for branched_from_commit_shas field")
				}
			}
			o.BranchedFromCommitShas = na
		}
	}
	if o.BranchedFromCommitShas == nil {
		o.BranchedFromCommitShas = make([]string, 0)
	}

	if val, ok := kv["commit_ids"]; ok {
		if val != nil {
			na := make([]string, 0)
			if a, ok := val.([]string); ok {
				na = append(na, a...)
			} else {
				if a, ok := val.([]interface{}); ok {
					for _, ae := range a {
						if av, ok := ae.(string); ok {
							na = append(na, av)
						} else {
							if badMap, ok := ae.(map[interface{}]interface{}); ok {
								ae = slice.ConvertToStringToInterface(badMap)
							}
							b, _ := json.Marshal(ae)
							var av string
							if err := json.Unmarshal(b, &av); err != nil {
								panic("unsupported type for commit_ids field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for commit_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for commit_ids field")
				}
			}
			o.CommitIds = na
		}
	}
	if o.CommitIds == nil {
		o.CommitIds = make([]string, 0)
	}

	if val, ok := kv["commit_shas"]; ok {
		if val != nil {
			na := make([]string, 0)
			if a, ok := val.([]string); ok {
				na = append(na, a...)
			} else {
				if a, ok := val.([]interface{}); ok {
					for _, ae := range a {
						if av, ok := ae.(string); ok {
							na = append(na, av)
						} else {
							if badMap, ok := ae.(map[interface{}]interface{}); ok {
								ae = slice.ConvertToStringToInterface(badMap)
							}
							b, _ := json.Marshal(ae)
							var av string
							if err := json.Unmarshal(b, &av); err != nil {
								panic("unsupported type for commit_shas field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for commit_shas field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for commit_shas field")
				}
			}
			o.CommitShas = na
		}
	}
	if o.CommitShas == nil {
		o.CommitShas = make([]string, 0)
	}

	if val, ok := kv["customer_id"].(string); ok {
		o.CustomerID = val
	} else {
		if val, ok := kv["customer_id"]; ok {
			if val == nil {
				o.CustomerID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.CustomerID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["default"].(bool); ok {
		o.Default = val
	} else {
		if val, ok := kv["default"]; ok {
			if val == nil {
				o.Default = number.ToBoolAny(nil)
			} else {
				o.Default = number.ToBoolAny(val)
			}
		}
	}

	if val, ok := kv["first_branched_from_commit_id"].(string); ok {
		o.FirstBranchedFromCommitID = val
	} else {
		if val, ok := kv["first_branched_from_commit_id"]; ok {
			if val == nil {
				o.FirstBranchedFromCommitID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.FirstBranchedFromCommitID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["first_branched_from_commit_sha"].(string); ok {
		o.FirstBranchedFromCommitSha = val
	} else {
		if val, ok := kv["first_branched_from_commit_sha"]; ok {
			if val == nil {
				o.FirstBranchedFromCommitSha = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.FirstBranchedFromCommitSha = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["first_commit_id"].(string); ok {
		o.FirstCommitID = val
	} else {
		if val, ok := kv["first_commit_id"]; ok {
			if val == nil {
				o.FirstCommitID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.FirstCommitID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["first_commit_sha"].(string); ok {
		o.FirstCommitSha = val
	} else {
		if val, ok := kv["first_commit_sha"]; ok {
			if val == nil {
				o.FirstCommitSha = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.FirstCommitSha = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["id"].(string); ok {
		o.ID = val
	} else {
		if val, ok := kv["id"]; ok {
			if val == nil {
				o.ID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.ID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["merge_commit_id"].(string); ok {
		o.MergeCommitID = val
	} else {
		if val, ok := kv["merge_commit_id"]; ok {
			if val == nil {
				o.MergeCommitID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.MergeCommitID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["merge_commit_sha"].(string); ok {
		o.MergeCommitSha = val
	} else {
		if val, ok := kv["merge_commit_sha"]; ok {
			if val == nil {
				o.MergeCommitSha = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.MergeCommitSha = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["merged"].(bool); ok {
		o.Merged = val
	} else {
		if val, ok := kv["merged"]; ok {
			if val == nil {
				o.Merged = number.ToBoolAny(nil)
			} else {
				o.Merged = number.ToBoolAny(val)
			}
		}
	}

	if val, ok := kv["name"].(string); ok {
		o.Name = val
	} else {
		if val, ok := kv["name"]; ok {
			if val == nil {
				o.Name = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Name = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["ref_id"].(string); ok {
		o.RefID = val
	} else {
		if val, ok := kv["ref_id"]; ok {
			if val == nil {
				o.RefID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.RefID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["ref_type"].(string); ok {
		o.RefType = val
	} else {
		if val, ok := kv["ref_type"]; ok {
			if val == nil {
				o.RefType = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.RefType = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["repo_id"].(string); ok {
		o.RepoID = val
	} else {
		if val, ok := kv["repo_id"]; ok {
			if val == nil {
				o.RepoID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.RepoID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["updated_ts"].(int64); ok {
		o.UpdatedAt = val
	} else {
		if val, ok := kv["updated_ts"]; ok {
			if val == nil {
				o.UpdatedAt = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.UpdatedAt = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["url"].(string); ok {
		o.URL = val
	} else {
		if val, ok := kv["url"]; ok {
			if val == nil {
				o.URL = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.URL = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *Branch) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.AheadDefaultCount)
	args = append(args, o.BehindDefaultCount)
	args = append(args, o.BranchedFromCommitIds)
	args = append(args, o.BranchedFromCommitShas)
	args = append(args, o.CommitIds)
	args = append(args, o.CommitShas)
	args = append(args, o.CustomerID)
	args = append(args, o.Default)
	args = append(args, o.FirstBranchedFromCommitID)
	args = append(args, o.FirstBranchedFromCommitSha)
	args = append(args, o.FirstCommitID)
	args = append(args, o.FirstCommitSha)
	args = append(args, o.ID)
	args = append(args, o.MergeCommitID)
	args = append(args, o.MergeCommitSha)
	args = append(args, o.Merged)
	args = append(args, o.Name)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RepoID)
	args = append(args, o.UpdatedAt)
	args = append(args, o.URL)
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
				"name": "branched_from_commit_ids",
				"type": map[string]interface{}{"items": "string", "name": "branched_from_commit_ids", "type": "array"},
			},
			map[string]interface{}{
				"name": "branched_from_commit_shas",
				"type": map[string]interface{}{"items": "string", "name": "branched_from_commit_shas", "type": "array"},
			},
			map[string]interface{}{
				"name": "commit_ids",
				"type": map[string]interface{}{"items": "string", "name": "commit_ids", "type": "array"},
			},
			map[string]interface{}{
				"name": "commit_shas",
				"type": map[string]interface{}{"items": "string", "name": "commit_shas", "type": "array"},
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
				"name": "first_branched_from_commit_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "first_branched_from_commit_sha",
				"type": "string",
			},
			map[string]interface{}{
				"name": "first_commit_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "first_commit_sha",
				"type": "string",
			},
			map[string]interface{}{
				"name": "id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "merge_commit_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "merge_commit_sha",
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
			map[string]interface{}{
				"name": "updated_ts",
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

// GetEventAPIConfig returns the EventAPIConfig
func (o *Branch) GetEventAPIConfig() datamodel.EventAPIConfig {
	return datamodel.EventAPIConfig{
		Publish: datamodel.EventAPIPublish{
			Public: false,
		},
		Subscribe: datamodel.EventAPISubscribe{
			Public: false,
			Key:    "",
		},
	}
}

// GetBranchAvroSchema creates the avro schema for Branch
func GetBranchAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetBranchAvroSchemaSpec())
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
	emptyTime := time.Unix(0, 0)
	var numPartitions int
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
					if numPartitions == 0 {
						numPartitions = object.GetTopicConfig().NumPartitions
					}
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
					if tv.IsZero() || tv.Equal(emptyTime) {
						tv = time.Now() // if its still zero, use the ingest time
					}
					// add generated message headers
					headers["message-id"] = pstrings.NewUUIDV4()
					headers["message-ts"] = fmt.Sprintf("%v", datetime.EpochNow())
					// determine the partition selection by using the partition key
					// and taking the modulo over the number of partitions for the topic
					partition := hash.Modulo(item.Key(), numPartitions)
					msg := eventing.Message{
						Encoding:  eventing.AvroEncoding,
						Key:       item.(ModelWithID).GetID(),
						Value:     binary,
						Codec:     codec,
						Headers:   headers,
						Timestamp: tv,
						Partition: int32(partition),
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
					return fmt.Errorf("error unmarshaling avro data into sourcecode.Branch: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for sourcecode.Branch")
			}

			// ignore messages that have exceeded the TTL
			cfg := object.GetTopicConfig()
			if cfg != nil && cfg.TTL != 0 && msg.Timestamp.UTC().Add(cfg.TTL).Sub(time.Now().UTC()) < 0 {
				// if disable auto and we're skipping, we need to commit the message
				if !msg.IsAutoCommit() {
					msg.Commit()
				}
				return nil
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
