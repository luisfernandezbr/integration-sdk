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
	// PullReuqestCommitTopic is the default topic name
	PullReuqestCommitTopic datamodel.TopicNameType = "sourcecode_PullReuqestCommit_topic"

	// PullReuqestCommitStream is the default stream name
	PullReuqestCommitStream datamodel.TopicNameType = "sourcecode_PullReuqestCommit_stream"

	// PullReuqestCommitTable is the default table name
	PullReuqestCommitTable datamodel.TopicNameType = "sourcecode_pullreuqestcommit"

	// PullReuqestCommitModelName is the model name
	PullReuqestCommitModelName datamodel.ModelNameType = "sourcecode.PullReuqestCommit"
)

const (
	// PullReuqestCommitAdditionsColumn is the additions column name
	PullReuqestCommitAdditionsColumn = "additions"
	// PullReuqestCommitAuthorRefIDColumn is the author_ref_id column name
	PullReuqestCommitAuthorRefIDColumn = "author_ref_id"
	// PullReuqestCommitBranchIDColumn is the branch_id column name
	PullReuqestCommitBranchIDColumn = "branch_id"
	// PullReuqestCommitCommitterRefIDColumn is the committer_ref_id column name
	PullReuqestCommitCommitterRefIDColumn = "committer_ref_id"
	// PullReuqestCommitCreatedDateColumn is the created_date column name
	PullReuqestCommitCreatedDateColumn = "created_date"
	// PullReuqestCommitCreatedDateColumnEpochColumn is the epoch column property of the CreatedDate name
	PullReuqestCommitCreatedDateColumnEpochColumn = "created_date->epoch"
	// PullReuqestCommitCreatedDateColumnOffsetColumn is the offset column property of the CreatedDate name
	PullReuqestCommitCreatedDateColumnOffsetColumn = "created_date->offset"
	// PullReuqestCommitCreatedDateColumnRfc3339Column is the rfc3339 column property of the CreatedDate name
	PullReuqestCommitCreatedDateColumnRfc3339Column = "created_date->rfc3339"
	// PullReuqestCommitCustomerIDColumn is the customer_id column name
	PullReuqestCommitCustomerIDColumn = "customer_id"
	// PullReuqestCommitDeletionsColumn is the deletions column name
	PullReuqestCommitDeletionsColumn = "deletions"
	// PullReuqestCommitIDColumn is the id column name
	PullReuqestCommitIDColumn = "id"
	// PullReuqestCommitMessageColumn is the message column name
	PullReuqestCommitMessageColumn = "message"
	// PullReuqestCommitRefIDColumn is the ref_id column name
	PullReuqestCommitRefIDColumn = "ref_id"
	// PullReuqestCommitRefTypeColumn is the ref_type column name
	PullReuqestCommitRefTypeColumn = "ref_type"
	// PullReuqestCommitRepoIDColumn is the repo_id column name
	PullReuqestCommitRepoIDColumn = "repo_id"
	// PullReuqestCommitShaColumn is the sha column name
	PullReuqestCommitShaColumn = "sha"
	// PullReuqestCommitUpdatedAtColumn is the updated_ts column name
	PullReuqestCommitUpdatedAtColumn = "updated_ts"
	// PullReuqestCommitURLColumn is the url column name
	PullReuqestCommitURLColumn = "url"
)

// PullReuqestCommitCreatedDate represents the object structure for created_date
type PullReuqestCommitCreatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toPullReuqestCommitCreatedDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toPullReuqestCommitCreatedDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *PullReuqestCommitCreatedDate:
		return v.ToMap(isavro)

	default:
		return o
	}
}

func (o *PullReuqestCommitCreatedDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toPullReuqestCommitCreatedDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toPullReuqestCommitCreatedDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toPullReuqestCommitCreatedDateObject(o.Rfc3339, isavro, false, "string"),
	}
}

func (o *PullReuqestCommitCreatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *PullReuqestCommitCreatedDate) FromMap(kv map[string]interface{}) {

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

// PullReuqestCommit the pull request commit is a specific change in a repo that was extracted from a pull request
type PullReuqestCommit struct {
	// Additions the number of additions for the commit
	Additions int64 `json:"additions" bson:"additions" yaml:"additions" faker:"-"`
	// AuthorRefID the author ref_id in the source system
	AuthorRefID string `json:"author_ref_id" bson:"author_ref_id" yaml:"author_ref_id" faker:"-"`
	// BranchID the branch that the commit was a part of
	BranchID string `json:"branch_id" bson:"branch_id" yaml:"branch_id" faker:"-"`
	// CommitterRefID the committer ref_id in the source system
	CommitterRefID string `json:"committer_ref_id" bson:"committer_ref_id" yaml:"committer_ref_id" faker:"-"`
	// CreatedDate date when the commit was created
	CreatedDate PullReuqestCommitCreatedDate `json:"created_date" bson:"created_date" yaml:"created_date" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Deletions the number of deletions for the commit
	Deletions int64 `json:"deletions" bson:"deletions" yaml:"deletions" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// Message the commit message
	Message string `json:"message" bson:"message" yaml:"message" faker:"commit_message"`
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
var _ datamodel.Model = (*PullReuqestCommit)(nil)

func toPullReuqestCommitObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toPullReuqestCommitObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *PullReuqestCommit:
		return v.ToMap(isavro)

	case PullReuqestCommitCreatedDate:
		return v.ToMap(isavro)

	default:
		return o
	}
}

// String returns a string representation of PullReuqestCommit
func (o *PullReuqestCommit) String() string {
	return fmt.Sprintf("sourcecode.PullReuqestCommit<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *PullReuqestCommit) GetTopicName() datamodel.TopicNameType {
	return PullReuqestCommitTopic
}

// GetModelName returns the name of the model
func (o *PullReuqestCommit) GetModelName() datamodel.ModelNameType {
	return PullReuqestCommitModelName
}

// GetStreamName returns the name of the stream
func (o *PullReuqestCommit) GetStreamName() string {
	return PullReuqestCommitStream.String()
}

// GetTableName returns the name of the table
func (o *PullReuqestCommit) GetTableName() string {
	return PullReuqestCommitTable.String()
}

// NewPullReuqestCommitID provides a template for generating an ID field for PullReuqestCommit
func NewPullReuqestCommitID(customerID string, refID string, refType string, RepoID string) string {
	return hash.Values(customerID, refID, refType, RepoID)
}

func (o *PullReuqestCommit) setDefaults(frommap bool) {

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
func (o *PullReuqestCommit) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *PullReuqestCommit) GetTopicKey() string {
	var i interface{} = o.RepoID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *PullReuqestCommit) GetTimestamp() time.Time {
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
	panic("not sure how to handle the date time format for PullReuqestCommit")
}

// GetRefID returns the RefID for the object
func (o *PullReuqestCommit) GetRefID() string {
	o.RefID = o.Sha
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *PullReuqestCommit) IsMaterialized() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *PullReuqestCommit) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *PullReuqestCommit) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *PullReuqestCommit) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = PullReuqestCommitModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *PullReuqestCommit) GetTopicConfig() *datamodel.ModelTopicConfig {
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
func (o *PullReuqestCommit) GetStateKey() string {
	key := "repo_id"
	return fmt.Sprintf("%s_%s", key, o.GetID())
}

// GetCustomerID will return the customer_id
func (o *PullReuqestCommit) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of PullReuqestCommit
func (o *PullReuqestCommit) Clone() datamodel.Model {
	c := new(PullReuqestCommit)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *PullReuqestCommit) Anon() datamodel.Model {
	c := new(PullReuqestCommit)
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
func (o *PullReuqestCommit) MarshalBinary() ([]byte, error) {
	return o.MarshalJSON()
}

func (o *PullReuqestCommit) UnmarshalBinary(data []byte) error {
	return o.UnmarshalJSON(data)
}

// MarshalJSON returns the bytes for marshaling to json
func (o *PullReuqestCommit) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *PullReuqestCommit) UnmarshalJSON(data []byte) error {
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

var cachedCodecPullReuqestCommit *goavro.Codec
var cachedCodecPullReuqestCommitLock sync.Mutex

// GetAvroCodec returns the avro codec for this model
func (o *PullReuqestCommit) GetAvroCodec() *goavro.Codec {
	cachedCodecPullReuqestCommitLock.Lock()
	if cachedCodecPullReuqestCommit == nil {
		c, err := GetPullReuqestCommitAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecPullReuqestCommit = c
	}
	cachedCodecPullReuqestCommitLock.Unlock()
	return cachedCodecPullReuqestCommit
}

// ToAvroBinary returns the data as Avro binary data
func (o *PullReuqestCommit) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *PullReuqestCommit) FromAvroBinary(value []byte) error {
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
func (o *PullReuqestCommit) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two PullReuqestCommit objects are equal
func (o *PullReuqestCommit) IsEqual(other *PullReuqestCommit) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *PullReuqestCommit) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	o.setDefaults(false)
	return map[string]interface{}{
		"additions":        toPullReuqestCommitObject(o.Additions, isavro, false, "long"),
		"author_ref_id":    toPullReuqestCommitObject(o.AuthorRefID, isavro, false, "string"),
		"branch_id":        toPullReuqestCommitObject(o.BranchID, isavro, false, "string"),
		"committer_ref_id": toPullReuqestCommitObject(o.CommitterRefID, isavro, false, "string"),
		"created_date":     toPullReuqestCommitObject(o.CreatedDate, isavro, false, "created_date"),
		"customer_id":      toPullReuqestCommitObject(o.CustomerID, isavro, false, "string"),
		"deletions":        toPullReuqestCommitObject(o.Deletions, isavro, false, "long"),
		"id":               toPullReuqestCommitObject(o.ID, isavro, false, "string"),
		"message":          toPullReuqestCommitObject(o.Message, isavro, false, "string"),
		"ref_id":           toPullReuqestCommitObject(o.RefID, isavro, false, "string"),
		"ref_type":         toPullReuqestCommitObject(o.RefType, isavro, false, "string"),
		"repo_id":          toPullReuqestCommitObject(o.RepoID, isavro, false, "string"),
		"sha":              toPullReuqestCommitObject(o.Sha, isavro, false, "string"),
		"updated_ts":       toPullReuqestCommitObject(o.UpdatedAt, isavro, false, "long"),
		"url":              toPullReuqestCommitObject(o.URL, isavro, false, "string"),
		"hashcode":         toPullReuqestCommitObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *PullReuqestCommit) FromMap(kv map[string]interface{}) {

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
		} else if sv, ok := val.(PullReuqestCommitCreatedDate); ok {
			// struct
			o.CreatedDate = sv
		} else if sp, ok := val.(*PullReuqestCommitCreatedDate); ok {
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
func (o *PullReuqestCommit) Hash() string {
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
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RepoID)
	args = append(args, o.Sha)
	args = append(args, o.UpdatedAt)
	args = append(args, o.URL)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetPullReuqestCommitAvroSchemaSpec creates the avro schema specification for PullReuqestCommit
func GetPullReuqestCommitAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "sourcecode",
		"name":      "PullReuqestCommit",
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
func (o *PullReuqestCommit) GetEventAPIConfig() datamodel.EventAPIConfig {
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

// GetPullReuqestCommitAvroSchema creates the avro schema for PullReuqestCommit
func GetPullReuqestCommitAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetPullReuqestCommitAvroSchemaSpec())
}

// PullReuqestCommitSendEvent is an event detail for sending data
type PullReuqestCommitSendEvent struct {
	PullReuqestCommit *PullReuqestCommit
	headers           map[string]string
	time              time.Time
	key               string
}

var _ datamodel.ModelSendEvent = (*PullReuqestCommitSendEvent)(nil)

// Key is the key to use for the message
func (e *PullReuqestCommitSendEvent) Key() string {
	if e.key == "" {
		return e.PullReuqestCommit.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *PullReuqestCommitSendEvent) Object() datamodel.Model {
	return e.PullReuqestCommit
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *PullReuqestCommitSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *PullReuqestCommitSendEvent) Timestamp() time.Time {
	return e.time
}

// PullReuqestCommitSendEventOpts is a function handler for setting opts
type PullReuqestCommitSendEventOpts func(o *PullReuqestCommitSendEvent)

// WithPullReuqestCommitSendEventKey sets the key value to a value different than the object ID
func WithPullReuqestCommitSendEventKey(key string) PullReuqestCommitSendEventOpts {
	return func(o *PullReuqestCommitSendEvent) {
		o.key = key
	}
}

// WithPullReuqestCommitSendEventTimestamp sets the timestamp value
func WithPullReuqestCommitSendEventTimestamp(tv time.Time) PullReuqestCommitSendEventOpts {
	return func(o *PullReuqestCommitSendEvent) {
		o.time = tv
	}
}

// WithPullReuqestCommitSendEventHeader sets the timestamp value
func WithPullReuqestCommitSendEventHeader(key, value string) PullReuqestCommitSendEventOpts {
	return func(o *PullReuqestCommitSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewPullReuqestCommitSendEvent returns a new PullReuqestCommitSendEvent instance
func NewPullReuqestCommitSendEvent(o *PullReuqestCommit, opts ...PullReuqestCommitSendEventOpts) *PullReuqestCommitSendEvent {
	res := &PullReuqestCommitSendEvent{
		PullReuqestCommit: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewPullReuqestCommitProducer will stream data from the channel
func NewPullReuqestCommitProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
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
				if object, ok := item.Object().(*PullReuqestCommit); ok {
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
					errors <- fmt.Errorf("invalid event received. expected an object of type sourcecode.PullReuqestCommit but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewPullReuqestCommitConsumer will stream data from the topic into the provided channel
func NewPullReuqestCommitConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object PullReuqestCommit
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into sourcecode.PullReuqestCommit: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avro data into sourcecode.PullReuqestCommit: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for sourcecode.PullReuqestCommit")
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

			ch <- &PullReuqestCommitReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object PullReuqestCommit
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &PullReuqestCommitReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// PullReuqestCommitReceiveEvent is an event detail for receiving data
type PullReuqestCommitReceiveEvent struct {
	PullReuqestCommit *PullReuqestCommit
	message           eventing.Message
	eof               bool
}

var _ datamodel.ModelReceiveEvent = (*PullReuqestCommitReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *PullReuqestCommitReceiveEvent) Object() datamodel.Model {
	return e.PullReuqestCommit
}

// Message returns the underlying message data for the event
func (e *PullReuqestCommitReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *PullReuqestCommitReceiveEvent) EOF() bool {
	return e.eof
}

// PullReuqestCommitProducer implements the datamodel.ModelEventProducer
type PullReuqestCommitProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*PullReuqestCommitProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *PullReuqestCommitProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *PullReuqestCommitProducer) Close() error {
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
func (o *PullReuqestCommit) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *PullReuqestCommit) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &PullReuqestCommitProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewPullReuqestCommitProducer(newctx, producer, ch, errors, empty),
	}
}

// NewPullReuqestCommitProducerChannel returns a channel which can be used for producing Model events
func NewPullReuqestCommitProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewPullReuqestCommitProducerChannelSize(producer, 0, errors)
}

// NewPullReuqestCommitProducerChannelSize returns a channel which can be used for producing Model events
func NewPullReuqestCommitProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &PullReuqestCommitProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewPullReuqestCommitProducer(newctx, producer, ch, errors, empty),
	}
}

// PullReuqestCommitConsumer implements the datamodel.ModelEventConsumer
type PullReuqestCommitConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*PullReuqestCommitConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *PullReuqestCommitConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *PullReuqestCommitConsumer) Close() error {
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
func (o *PullReuqestCommit) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &PullReuqestCommitConsumer{
		ch:       ch,
		callback: NewPullReuqestCommitConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewPullReuqestCommitConsumerChannel returns a consumer channel which can be used to consume Model events
func NewPullReuqestCommitConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &PullReuqestCommitConsumer{
		ch:       ch,
		callback: NewPullReuqestCommitConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
