// DO NOT EDIT -- generated code

// Package sourcecode - the system which contains source code
package sourcecode

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"reflect"
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
	pstrings "github.com/pinpt/go-common/strings"
)

const (
	// PullRequestCommitTopic is the default topic name
	PullRequestCommitTopic datamodel.TopicNameType = "sourcecode_PullRequestCommit_topic"

	// PullRequestCommitStream is the default stream name
	PullRequestCommitStream datamodel.TopicNameType = "sourcecode_PullRequestCommit_stream"

	// PullRequestCommitTable is the default table name
	PullRequestCommitTable datamodel.TopicNameType = "sourcecode_pullrequestcommit"

	// PullRequestCommitModelName is the model name
	PullRequestCommitModelName datamodel.ModelNameType = "sourcecode.PullRequestCommit"
)

const (
	// PullRequestCommitAdditionsColumn is the additions column name
	PullRequestCommitAdditionsColumn = "additions"
	// PullRequestCommitAuthorRefIDColumn is the author_ref_id column name
	PullRequestCommitAuthorRefIDColumn = "author_ref_id"
	// PullRequestCommitBranchIDColumn is the branch_id column name
	PullRequestCommitBranchIDColumn = "branch_id"
	// PullRequestCommitCommitterRefIDColumn is the committer_ref_id column name
	PullRequestCommitCommitterRefIDColumn = "committer_ref_id"
	// PullRequestCommitCreatedDateColumn is the created_date column name
	PullRequestCommitCreatedDateColumn = "created_date"
	// PullRequestCommitCreatedDateColumnEpochColumn is the epoch column property of the CreatedDate name
	PullRequestCommitCreatedDateColumnEpochColumn = "created_date->epoch"
	// PullRequestCommitCreatedDateColumnOffsetColumn is the offset column property of the CreatedDate name
	PullRequestCommitCreatedDateColumnOffsetColumn = "created_date->offset"
	// PullRequestCommitCreatedDateColumnRfc3339Column is the rfc3339 column property of the CreatedDate name
	PullRequestCommitCreatedDateColumnRfc3339Column = "created_date->rfc3339"
	// PullRequestCommitCustomerIDColumn is the customer_id column name
	PullRequestCommitCustomerIDColumn = "customer_id"
	// PullRequestCommitDeletionsColumn is the deletions column name
	PullRequestCommitDeletionsColumn = "deletions"
	// PullRequestCommitIDColumn is the id column name
	PullRequestCommitIDColumn = "id"
	// PullRequestCommitMessageColumn is the message column name
	PullRequestCommitMessageColumn = "message"
	// PullRequestCommitPullRequestIDColumn is the pull_request_id column name
	PullRequestCommitPullRequestIDColumn = "pull_request_id"
	// PullRequestCommitRefIDColumn is the ref_id column name
	PullRequestCommitRefIDColumn = "ref_id"
	// PullRequestCommitRefTypeColumn is the ref_type column name
	PullRequestCommitRefTypeColumn = "ref_type"
	// PullRequestCommitRepoIDColumn is the repo_id column name
	PullRequestCommitRepoIDColumn = "repo_id"
	// PullRequestCommitShaColumn is the sha column name
	PullRequestCommitShaColumn = "sha"
	// PullRequestCommitUpdatedAtColumn is the updated_ts column name
	PullRequestCommitUpdatedAtColumn = "updated_ts"
	// PullRequestCommitURLColumn is the url column name
	PullRequestCommitURLColumn = "url"
)

// PullRequestCommitCreatedDate represents the object structure for created_date
type PullRequestCommitCreatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toPullRequestCommitCreatedDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toPullRequestCommitCreatedDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *PullRequestCommitCreatedDate:
		return v.ToMap(isavro)

	default:
		return o
	}
}

func (o *PullRequestCommitCreatedDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toPullRequestCommitCreatedDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toPullRequestCommitCreatedDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toPullRequestCommitCreatedDateObject(o.Rfc3339, isavro, false, "string"),
	}
}

func (o *PullRequestCommitCreatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *PullRequestCommitCreatedDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["epoch"].(int64); ok {
		o.Epoch = val
	} else {
		if val, ok := kv["epoch"]; ok {
			if val == nil {
				o.Epoch = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Epoch = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["offset"].(int64); ok {
		o.Offset = val
	} else {
		if val, ok := kv["offset"]; ok {
			if val == nil {
				o.Offset = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Offset = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["rfc3339"].(string); ok {
		o.Rfc3339 = val
	} else {
		if val, ok := kv["rfc3339"]; ok {
			if val == nil {
				o.Rfc3339 = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// PullRequestCommit the pull request commit is a specific change in a repo that was extracted from a pull request
type PullRequestCommit struct {
	// Additions the number of additions for the commit
	Additions int64 `json:"additions" bson:"additions" yaml:"additions" faker:"-"`
	// AuthorRefID the author ref_id in the source system
	AuthorRefID string `json:"author_ref_id" bson:"author_ref_id" yaml:"author_ref_id" faker:"-"`
	// BranchID the branch that the commit was a part of
	BranchID string `json:"branch_id" bson:"branch_id" yaml:"branch_id" faker:"-"`
	// CommitterRefID the committer ref_id in the source system
	CommitterRefID string `json:"committer_ref_id" bson:"committer_ref_id" yaml:"committer_ref_id" faker:"-"`
	// CreatedDate date when the commit was created
	CreatedDate PullRequestCommitCreatedDate `json:"created_date" bson:"created_date" yaml:"created_date" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Deletions the number of deletions for the commit
	Deletions int64 `json:"deletions" bson:"deletions" yaml:"deletions" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// Message the commit message
	Message string `json:"message" bson:"message" yaml:"message" faker:"commit_message"`
	// PullRequestID the pull request this commit was taken from
	PullRequestID string `json:"pull_request_id" bson:"pull_request_id" yaml:"pull_request_id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RepoID the unique id for the repo
	RepoID string `json:"repo_id" bson:"repo_id" yaml:"repo_id" faker:"-"`
	// Sha the unique sha for the commit
	Sha string `json:"sha" bson:"sha" yaml:"sha" faker:"sha"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// URL the url to the commit detail
	URL string `json:"url" bson:"url" yaml:"url" faker:"url"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*PullRequestCommit)(nil)

func toPullRequestCommitObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toPullRequestCommitObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *PullRequestCommit:
		return v.ToMap(isavro)

	case PullRequestCommitCreatedDate:
		return v.ToMap(isavro)

	default:
		return o
	}
}

// String returns a string representation of PullRequestCommit
func (o *PullRequestCommit) String() string {
	return fmt.Sprintf("sourcecode.PullRequestCommit<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *PullRequestCommit) GetTopicName() datamodel.TopicNameType {
	return PullRequestCommitTopic
}

// GetModelName returns the name of the model
func (o *PullRequestCommit) GetModelName() datamodel.ModelNameType {
	return PullRequestCommitModelName
}

// GetStreamName returns the name of the stream
func (o *PullRequestCommit) GetStreamName() string {
	return PullRequestCommitStream.String()
}

// GetTableName returns the name of the table
func (o *PullRequestCommit) GetTableName() string {
	return PullRequestCommitTable.String()
}

// NewPullRequestCommitID provides a template for generating an ID field for PullRequestCommit
func NewPullRequestCommitID(customerID string, refID string, refType string, RepoID string) string {
	return hash.Values(customerID, refID, refType, RepoID)
}

func (o *PullRequestCommit) setDefaults(frommap bool) {

	if o.ID == "" {
		o.ID = hash.Values(o.CustomerID, o.RefID, o.RefType, o.RepoID)
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *PullRequestCommit) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *PullRequestCommit) GetTopicKey() string {
	var i interface{} = o.RepoID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *PullRequestCommit) GetTimestamp() time.Time {
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
	panic("not sure how to handle the date time format for PullRequestCommit")
}

// GetRefID returns the RefID for the object
func (o *PullRequestCommit) GetRefID() string {
	o.RefID = o.Sha
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *PullRequestCommit) IsMaterialized() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *PullRequestCommit) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *PullRequestCommit) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *PullRequestCommit) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = PullRequestCommitModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *PullRequestCommit) GetTopicConfig() *datamodel.ModelTopicConfig {
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
func (o *PullRequestCommit) GetStateKey() string {
	key := "repo_id"
	return fmt.Sprintf("%s_%s", key, o.GetID())
}

// GetCustomerID will return the customer_id
func (o *PullRequestCommit) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of PullRequestCommit
func (o *PullRequestCommit) Clone() datamodel.Model {
	c := new(PullRequestCommit)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *PullRequestCommit) Anon() datamodel.Model {
	c := new(PullRequestCommit)
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
func (o *PullRequestCommit) MarshalBinary() ([]byte, error) {
	return o.MarshalJSON()
}

func (o *PullRequestCommit) UnmarshalBinary(data []byte) error {
	return o.UnmarshalJSON(data)
}

// MarshalJSON returns the bytes for marshaling to json
func (o *PullRequestCommit) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *PullRequestCommit) UnmarshalJSON(data []byte) error {
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

var cachedCodecPullRequestCommit *goavro.Codec
var cachedCodecPullRequestCommitLock sync.Mutex

// GetAvroCodec returns the avro codec for this model
func (o *PullRequestCommit) GetAvroCodec() *goavro.Codec {
	cachedCodecPullRequestCommitLock.Lock()
	if cachedCodecPullRequestCommit == nil {
		c, err := GetPullRequestCommitAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecPullRequestCommit = c
	}
	cachedCodecPullRequestCommitLock.Unlock()
	return cachedCodecPullRequestCommit
}

// ToAvroBinary returns the data as Avro binary data
func (o *PullRequestCommit) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *PullRequestCommit) FromAvroBinary(value []byte) error {
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
func (o *PullRequestCommit) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two PullRequestCommit objects are equal
func (o *PullRequestCommit) IsEqual(other *PullRequestCommit) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *PullRequestCommit) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	o.setDefaults(false)
	return map[string]interface{}{
		"additions":        toPullRequestCommitObject(o.Additions, isavro, false, "long"),
		"author_ref_id":    toPullRequestCommitObject(o.AuthorRefID, isavro, false, "string"),
		"branch_id":        toPullRequestCommitObject(o.BranchID, isavro, false, "string"),
		"committer_ref_id": toPullRequestCommitObject(o.CommitterRefID, isavro, false, "string"),
		"created_date":     toPullRequestCommitObject(o.CreatedDate, isavro, false, "created_date"),
		"customer_id":      toPullRequestCommitObject(o.CustomerID, isavro, false, "string"),
		"deletions":        toPullRequestCommitObject(o.Deletions, isavro, false, "long"),
		"id":               toPullRequestCommitObject(o.ID, isavro, false, "string"),
		"message":          toPullRequestCommitObject(o.Message, isavro, false, "string"),
		"pull_request_id":  toPullRequestCommitObject(o.PullRequestID, isavro, false, "string"),
		"ref_id":           toPullRequestCommitObject(o.RefID, isavro, false, "string"),
		"ref_type":         toPullRequestCommitObject(o.RefType, isavro, false, "string"),
		"repo_id":          toPullRequestCommitObject(o.RepoID, isavro, false, "string"),
		"sha":              toPullRequestCommitObject(o.Sha, isavro, false, "string"),
		"updated_ts":       toPullRequestCommitObject(o.UpdatedAt, isavro, false, "long"),
		"url":              toPullRequestCommitObject(o.URL, isavro, false, "string"),
		"hashcode":         toPullRequestCommitObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *PullRequestCommit) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["additions"].(int64); ok {
		o.Additions = val
	} else {
		if val, ok := kv["additions"]; ok {
			if val == nil {
				o.Additions = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Additions = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["author_ref_id"].(string); ok {
		o.AuthorRefID = val
	} else {
		if val, ok := kv["author_ref_id"]; ok {
			if val == nil {
				o.AuthorRefID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.AuthorRefID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["branch_id"].(string); ok {
		o.BranchID = val
	} else {
		if val, ok := kv["branch_id"]; ok {
			if val == nil {
				o.BranchID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.BranchID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["committer_ref_id"].(string); ok {
		o.CommitterRefID = val
	} else {
		if val, ok := kv["committer_ref_id"]; ok {
			if val == nil {
				o.CommitterRefID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.CommitterRefID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["created_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.CreatedDate.FromMap(kv)
		} else if sv, ok := val.(PullRequestCommitCreatedDate); ok {
			// struct
			o.CreatedDate = sv
		} else if sp, ok := val.(*PullRequestCommitCreatedDate); ok {
			// struct pointer
			o.CreatedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.CreatedDate.Epoch = dt.Epoch
			o.CreatedDate.Rfc3339 = dt.Rfc3339
			o.CreatedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.CreatedDate.Epoch = dt.Epoch
			o.CreatedDate.Rfc3339 = dt.Rfc3339
			o.CreatedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.CreatedDate.Epoch = dt.Epoch
				o.CreatedDate.Rfc3339 = dt.Rfc3339
				o.CreatedDate.Offset = dt.Offset
			}
		}
	} else {
		o.CreatedDate.FromMap(map[string]interface{}{})
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

	if val, ok := kv["deletions"].(int64); ok {
		o.Deletions = val
	} else {
		if val, ok := kv["deletions"]; ok {
			if val == nil {
				o.Deletions = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Deletions = number.ToInt64Any(val)
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

	if val, ok := kv["message"].(string); ok {
		o.Message = val
	} else {
		if val, ok := kv["message"]; ok {
			if val == nil {
				o.Message = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Message = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["pull_request_id"].(string); ok {
		o.PullRequestID = val
	} else {
		if val, ok := kv["pull_request_id"]; ok {
			if val == nil {
				o.PullRequestID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.PullRequestID = fmt.Sprintf("%v", val)
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

	if val, ok := kv["sha"].(string); ok {
		o.Sha = val
	} else {
		if val, ok := kv["sha"]; ok {
			if val == nil {
				o.Sha = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Sha = fmt.Sprintf("%v", val)
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
func (o *PullRequestCommit) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Additions)
	args = append(args, o.AuthorRefID)
	args = append(args, o.BranchID)
	args = append(args, o.CommitterRefID)
	args = append(args, o.CreatedDate)
	args = append(args, o.CustomerID)
	args = append(args, o.Deletions)
	args = append(args, o.ID)
	args = append(args, o.Message)
	args = append(args, o.PullRequestID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RepoID)
	args = append(args, o.Sha)
	args = append(args, o.UpdatedAt)
	args = append(args, o.URL)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetPullRequestCommitAvroSchemaSpec creates the avro schema specification for PullRequestCommit
func GetPullRequestCommitAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "sourcecode",
		"name":      "PullRequestCommit",
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
				"name": "author_ref_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "branch_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "committer_ref_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "created_date",
				"type": map[string]interface{}{"doc": "date when the commit was created", "fields": []interface{}{map[string]interface{}{"doc": "the date in epoch format", "name": "epoch", "type": "long"}, map[string]interface{}{"doc": "the timezone offset from GMT", "name": "offset", "type": "long"}, map[string]interface{}{"doc": "the date in RFC3339 format", "name": "rfc3339", "type": "string"}}, "name": "created_date", "type": "record"},
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
				"name": "id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "message",
				"type": "string",
			},
			map[string]interface{}{
				"name": "pull_request_id",
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
				"name": "sha",
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
func (o *PullRequestCommit) GetEventAPIConfig() datamodel.EventAPIConfig {
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

// GetPullRequestCommitAvroSchema creates the avro schema for PullRequestCommit
func GetPullRequestCommitAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetPullRequestCommitAvroSchemaSpec())
}

// PullRequestCommitSendEvent is an event detail for sending data
type PullRequestCommitSendEvent struct {
	PullRequestCommit *PullRequestCommit
	headers           map[string]string
	time              time.Time
	key               string
}

var _ datamodel.ModelSendEvent = (*PullRequestCommitSendEvent)(nil)

// Key is the key to use for the message
func (e *PullRequestCommitSendEvent) Key() string {
	if e.key == "" {
		return e.PullRequestCommit.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *PullRequestCommitSendEvent) Object() datamodel.Model {
	return e.PullRequestCommit
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *PullRequestCommitSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *PullRequestCommitSendEvent) Timestamp() time.Time {
	return e.time
}

// PullRequestCommitSendEventOpts is a function handler for setting opts
type PullRequestCommitSendEventOpts func(o *PullRequestCommitSendEvent)

// WithPullRequestCommitSendEventKey sets the key value to a value different than the object ID
func WithPullRequestCommitSendEventKey(key string) PullRequestCommitSendEventOpts {
	return func(o *PullRequestCommitSendEvent) {
		o.key = key
	}
}

// WithPullRequestCommitSendEventTimestamp sets the timestamp value
func WithPullRequestCommitSendEventTimestamp(tv time.Time) PullRequestCommitSendEventOpts {
	return func(o *PullRequestCommitSendEvent) {
		o.time = tv
	}
}

// WithPullRequestCommitSendEventHeader sets the timestamp value
func WithPullRequestCommitSendEventHeader(key, value string) PullRequestCommitSendEventOpts {
	return func(o *PullRequestCommitSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewPullRequestCommitSendEvent returns a new PullRequestCommitSendEvent instance
func NewPullRequestCommitSendEvent(o *PullRequestCommit, opts ...PullRequestCommitSendEventOpts) *PullRequestCommitSendEvent {
	res := &PullRequestCommitSendEvent{
		PullRequestCommit: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewPullRequestCommitProducer will stream data from the channel
func NewPullRequestCommitProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
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
				if object, ok := item.Object().(*PullRequestCommit); ok {
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
						Key:       object.GetID(),
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
					errors <- fmt.Errorf("invalid event received. expected an object of type sourcecode.PullRequestCommit but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewPullRequestCommitConsumer will stream data from the topic into the provided channel
func NewPullRequestCommitConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object PullRequestCommit
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into sourcecode.PullRequestCommit: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avro data into sourcecode.PullRequestCommit: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for sourcecode.PullRequestCommit")
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

			ch <- &PullRequestCommitReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object PullRequestCommit
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &PullRequestCommitReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// PullRequestCommitReceiveEvent is an event detail for receiving data
type PullRequestCommitReceiveEvent struct {
	PullRequestCommit *PullRequestCommit
	message           eventing.Message
	eof               bool
}

var _ datamodel.ModelReceiveEvent = (*PullRequestCommitReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *PullRequestCommitReceiveEvent) Object() datamodel.Model {
	return e.PullRequestCommit
}

// Message returns the underlying message data for the event
func (e *PullRequestCommitReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *PullRequestCommitReceiveEvent) EOF() bool {
	return e.eof
}

// PullRequestCommitProducer implements the datamodel.ModelEventProducer
type PullRequestCommitProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*PullRequestCommitProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *PullRequestCommitProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *PullRequestCommitProducer) Close() error {
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
func (o *PullRequestCommit) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *PullRequestCommit) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &PullRequestCommitProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewPullRequestCommitProducer(newctx, producer, ch, errors, empty),
	}
}

// NewPullRequestCommitProducerChannel returns a channel which can be used for producing Model events
func NewPullRequestCommitProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewPullRequestCommitProducerChannelSize(producer, 0, errors)
}

// NewPullRequestCommitProducerChannelSize returns a channel which can be used for producing Model events
func NewPullRequestCommitProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &PullRequestCommitProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewPullRequestCommitProducer(newctx, producer, ch, errors, empty),
	}
}

// PullRequestCommitConsumer implements the datamodel.ModelEventConsumer
type PullRequestCommitConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*PullRequestCommitConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *PullRequestCommitConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *PullRequestCommitConsumer) Close() error {
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
func (o *PullRequestCommit) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &PullRequestCommitConsumer{
		ch:       ch,
		callback: NewPullRequestCommitConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewPullRequestCommitConsumerChannel returns a consumer channel which can be used to consume Model events
func NewPullRequestCommitConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &PullRequestCommitConsumer{
		ch:       ch,
		callback: NewPullRequestCommitConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
