// DO NOT EDIT -- generated code

// Package sourcecode - the system which contains source code
package sourcecode

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"sync"
	"time"

	"github.com/bxcodec/faker"
	"github.com/linkedin/goavro"
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
)

const (

	// PullRequestCommentModelName is the model name
	PullRequestCommentModelName datamodel.ModelNameType = "sourcecode.PullRequestComment"
)

const (
	// PullRequestCommentBodyColumn is the body column name
	PullRequestCommentBodyColumn = "body"
	// PullRequestCommentCreatedDateColumn is the created_date column name
	PullRequestCommentCreatedDateColumn = "created_date"
	// PullRequestCommentCreatedDateColumnEpochColumn is the epoch column property of the CreatedDate name
	PullRequestCommentCreatedDateColumnEpochColumn = "created_date->epoch"
	// PullRequestCommentCreatedDateColumnOffsetColumn is the offset column property of the CreatedDate name
	PullRequestCommentCreatedDateColumnOffsetColumn = "created_date->offset"
	// PullRequestCommentCreatedDateColumnRfc3339Column is the rfc3339 column property of the CreatedDate name
	PullRequestCommentCreatedDateColumnRfc3339Column = "created_date->rfc3339"
	// PullRequestCommentCustomerIDColumn is the customer_id column name
	PullRequestCommentCustomerIDColumn = "customer_id"
	// PullRequestCommentIDColumn is the id column name
	PullRequestCommentIDColumn = "id"
	// PullRequestCommentPullRequestIDColumn is the pull_request_id column name
	PullRequestCommentPullRequestIDColumn = "pull_request_id"
	// PullRequestCommentRefIDColumn is the ref_id column name
	PullRequestCommentRefIDColumn = "ref_id"
	// PullRequestCommentRefTypeColumn is the ref_type column name
	PullRequestCommentRefTypeColumn = "ref_type"
	// PullRequestCommentRepoIDColumn is the repo_id column name
	PullRequestCommentRepoIDColumn = "repo_id"
	// PullRequestCommentUpdatedDateColumn is the updated_date column name
	PullRequestCommentUpdatedDateColumn = "updated_date"
	// PullRequestCommentUpdatedDateColumnEpochColumn is the epoch column property of the UpdatedDate name
	PullRequestCommentUpdatedDateColumnEpochColumn = "updated_date->epoch"
	// PullRequestCommentUpdatedDateColumnOffsetColumn is the offset column property of the UpdatedDate name
	PullRequestCommentUpdatedDateColumnOffsetColumn = "updated_date->offset"
	// PullRequestCommentUpdatedDateColumnRfc3339Column is the rfc3339 column property of the UpdatedDate name
	PullRequestCommentUpdatedDateColumnRfc3339Column = "updated_date->rfc3339"
	// PullRequestCommentUserRefIDColumn is the user_ref_id column name
	PullRequestCommentUserRefIDColumn = "user_ref_id"
)

// PullRequestCommentCreatedDate represents the object structure for created_date
type PullRequestCommentCreatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toPullRequestCommentCreatedDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toPullRequestCommentCreatedDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *PullRequestCommentCreatedDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *PullRequestCommentCreatedDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toPullRequestCommentCreatedDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toPullRequestCommentCreatedDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toPullRequestCommentCreatedDateObject(o.Rfc3339, isavro, false, "string"),
	}
}

func (o *PullRequestCommentCreatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *PullRequestCommentCreatedDate) FromMap(kv map[string]interface{}) {

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

// PullRequestCommentUpdatedDate represents the object structure for updated_date
type PullRequestCommentUpdatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toPullRequestCommentUpdatedDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toPullRequestCommentUpdatedDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *PullRequestCommentUpdatedDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *PullRequestCommentUpdatedDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toPullRequestCommentUpdatedDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toPullRequestCommentUpdatedDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toPullRequestCommentUpdatedDateObject(o.Rfc3339, isavro, false, "string"),
	}
}

func (o *PullRequestCommentUpdatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *PullRequestCommentUpdatedDate) FromMap(kv map[string]interface{}) {

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

// PullRequestComment the comment for a given pull request
type PullRequestComment struct {
	// Body the body of the comment
	Body string `json:"body" bson:"body" yaml:"body" faker:"commit_message"`
	// CreatedDate the timestamp in UTC that the comment was created
	CreatedDate PullRequestCommentCreatedDate `json:"created_date" bson:"created_date" yaml:"created_date" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// PullRequestID the pull request this comment is associated with
	PullRequestID string `json:"pull_request_id" bson:"pull_request_id" yaml:"pull_request_id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RepoID the unique id for the repo
	RepoID string `json:"repo_id" bson:"repo_id" yaml:"repo_id" faker:"-"`
	// UpdatedDate the timestamp in UTC that the comment was closed
	UpdatedDate PullRequestCommentUpdatedDate `json:"updated_date" bson:"updated_date" yaml:"updated_date" faker:"-"`
	// UserRefID the user ref_id in the source system
	UserRefID string `json:"user_ref_id" bson:"user_ref_id" yaml:"user_ref_id" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*PullRequestComment)(nil)

func toPullRequestCommentObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toPullRequestCommentObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *PullRequestComment:
		return v.ToMap(isavro)

	case PullRequestCommentCreatedDate:
		return v.ToMap(isavro)

	case PullRequestCommentUpdatedDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

// String returns a string representation of PullRequestComment
func (o *PullRequestComment) String() string {
	return fmt.Sprintf("sourcecode.PullRequestComment<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *PullRequestComment) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetModelName returns the name of the model
func (o *PullRequestComment) GetModelName() datamodel.ModelNameType {
	return PullRequestCommentModelName
}

func (o *PullRequestComment) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("PullRequestComment", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *PullRequestComment) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *PullRequestComment) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *PullRequestComment) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *PullRequestComment) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *PullRequestComment) IsMaterialized() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *PullRequestComment) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *PullRequestComment) IsEvented() bool {
	return false
}

// GetTopicConfig returns the topic config object
func (o *PullRequestComment) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetStateKey returns a key for use in state store
func (o *PullRequestComment) GetStateKey() string {
	key := ""
	return fmt.Sprintf("%s_%s", key, o.GetID())
}

// GetCustomerID will return the customer_id
func (o *PullRequestComment) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of PullRequestComment
func (o *PullRequestComment) Clone() datamodel.Model {
	c := new(PullRequestComment)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *PullRequestComment) Anon() datamodel.Model {
	c := new(PullRequestComment)
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
func (o *PullRequestComment) MarshalBinary() ([]byte, error) {
	return o.MarshalJSON()
}

func (o *PullRequestComment) UnmarshalBinary(data []byte) error {
	return o.UnmarshalJSON(data)
}

// MarshalJSON returns the bytes for marshaling to json
func (o *PullRequestComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *PullRequestComment) UnmarshalJSON(data []byte) error {
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

var cachedCodecPullRequestComment *goavro.Codec
var cachedCodecPullRequestCommentLock sync.Mutex

// GetAvroCodec returns the avro codec for this model
func (o *PullRequestComment) GetAvroCodec() *goavro.Codec {
	cachedCodecPullRequestCommentLock.Lock()
	if cachedCodecPullRequestComment == nil {
		c, err := GetPullRequestCommentAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecPullRequestComment = c
	}
	cachedCodecPullRequestCommentLock.Unlock()
	return cachedCodecPullRequestComment
}

// ToAvroBinary returns the data as Avro binary data
func (o *PullRequestComment) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *PullRequestComment) FromAvroBinary(value []byte) error {
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
func (o *PullRequestComment) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two PullRequestComment objects are equal
func (o *PullRequestComment) IsEqual(other *PullRequestComment) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *PullRequestComment) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	o.setDefaults(false)
	return map[string]interface{}{
		"body":            toPullRequestCommentObject(o.Body, isavro, false, "string"),
		"created_date":    toPullRequestCommentObject(o.CreatedDate, isavro, false, "created_date"),
		"customer_id":     toPullRequestCommentObject(o.CustomerID, isavro, false, "string"),
		"id":              toPullRequestCommentObject(o.ID, isavro, false, "string"),
		"pull_request_id": toPullRequestCommentObject(o.PullRequestID, isavro, false, "string"),
		"ref_id":          toPullRequestCommentObject(o.RefID, isavro, false, "string"),
		"ref_type":        toPullRequestCommentObject(o.RefType, isavro, false, "string"),
		"repo_id":         toPullRequestCommentObject(o.RepoID, isavro, false, "string"),
		"updated_date":    toPullRequestCommentObject(o.UpdatedDate, isavro, false, "updated_date"),
		"user_ref_id":     toPullRequestCommentObject(o.UserRefID, isavro, false, "string"),
		"hashcode":        toPullRequestCommentObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *PullRequestComment) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["body"].(string); ok {
		o.Body = val
	} else {
		if val, ok := kv["body"]; ok {
			if val == nil {
				o.Body = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Body = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["created_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.CreatedDate.FromMap(kv)
		} else if sv, ok := val.(PullRequestCommentCreatedDate); ok {
			// struct
			o.CreatedDate = sv
		} else if sp, ok := val.(*PullRequestCommentCreatedDate); ok {
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

	if val, ok := kv["updated_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.UpdatedDate.FromMap(kv)
		} else if sv, ok := val.(PullRequestCommentUpdatedDate); ok {
			// struct
			o.UpdatedDate = sv
		} else if sp, ok := val.(*PullRequestCommentUpdatedDate); ok {
			// struct pointer
			o.UpdatedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.UpdatedDate.Epoch = dt.Epoch
			o.UpdatedDate.Rfc3339 = dt.Rfc3339
			o.UpdatedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.UpdatedDate.Epoch = dt.Epoch
			o.UpdatedDate.Rfc3339 = dt.Rfc3339
			o.UpdatedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.UpdatedDate.Epoch = dt.Epoch
				o.UpdatedDate.Rfc3339 = dt.Rfc3339
				o.UpdatedDate.Offset = dt.Offset
			}
		}
	} else {
		o.UpdatedDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["user_ref_id"].(string); ok {
		o.UserRefID = val
	} else {
		if val, ok := kv["user_ref_id"]; ok {
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
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *PullRequestComment) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Body)
	args = append(args, o.CreatedDate)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.PullRequestID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RepoID)
	args = append(args, o.UpdatedDate)
	args = append(args, o.UserRefID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetPullRequestCommentAvroSchemaSpec creates the avro schema specification for PullRequestComment
func GetPullRequestCommentAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "sourcecode",
		"name":      "PullRequestComment",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "hashcode",
				"type": "string",
			},
			map[string]interface{}{
				"name": "body",
				"type": "string",
			},
			map[string]interface{}{
				"name": "created_date",
				"type": map[string]interface{}{"doc": "the timestamp in UTC that the comment was created", "fields": []interface{}{map[string]interface{}{"doc": "the date in epoch format", "name": "epoch", "type": "long"}, map[string]interface{}{"doc": "the timezone offset from GMT", "name": "offset", "type": "long"}, map[string]interface{}{"doc": "the date in RFC3339 format", "name": "rfc3339", "type": "string"}}, "name": "created_date", "type": "record"},
			},
			map[string]interface{}{
				"name": "customer_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "id",
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
				"name": "updated_date",
				"type": map[string]interface{}{"doc": "the timestamp in UTC that the comment was closed", "fields": []interface{}{map[string]interface{}{"doc": "the date in epoch format", "name": "epoch", "type": "long"}, map[string]interface{}{"doc": "the timezone offset from GMT", "name": "offset", "type": "long"}, map[string]interface{}{"doc": "the date in RFC3339 format", "name": "rfc3339", "type": "string"}}, "name": "updated_date", "type": "record"},
			},
			map[string]interface{}{
				"name": "user_ref_id",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *PullRequestComment) GetEventAPIConfig() datamodel.EventAPIConfig {
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

// GetPullRequestCommentAvroSchema creates the avro schema for PullRequestComment
func GetPullRequestCommentAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetPullRequestCommentAvroSchemaSpec())
}

// TransformPullRequestCommentFunc is a function for transforming PullRequestComment during processing
type TransformPullRequestCommentFunc func(input *PullRequestComment) (*PullRequestComment, error)

// NewPullRequestCommentPipe creates a pipe for processing PullRequestComment items
func NewPullRequestCommentPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformPullRequestCommentFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewPullRequestCommentInputStream(input, errors)
	var stream chan PullRequestComment
	if len(transforms) > 0 {
		stream = make(chan PullRequestComment, 1000)
	} else {
		stream = inch
	}
	outdone := NewPullRequestCommentOutputStream(output, stream, errors)
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

// NewPullRequestCommentInputStreamDir creates a channel for reading PullRequestComment as JSON newlines from a directory of files
func NewPullRequestCommentInputStreamDir(dir string, errors chan<- error, transforms ...TransformPullRequestCommentFunc) (chan PullRequestComment, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/sourcecode/pull_request_comment\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan PullRequestComment)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for pull_request_comment")
		ch := make(chan PullRequestComment)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewPullRequestCommentInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan PullRequestComment)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewPullRequestCommentInputStreamFile creates an channel for reading PullRequestComment as JSON newlines from filename
func NewPullRequestCommentInputStreamFile(filename string, errors chan<- error, transforms ...TransformPullRequestCommentFunc) (chan PullRequestComment, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan PullRequestComment)
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
			ch := make(chan PullRequestComment)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewPullRequestCommentInputStream(f, errors, transforms...)
}

// NewPullRequestCommentInputStream creates an channel for reading PullRequestComment as JSON newlines from stream
func NewPullRequestCommentInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformPullRequestCommentFunc) (chan PullRequestComment, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan PullRequestComment, 1000)
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
			var item PullRequestComment
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

// NewPullRequestCommentOutputStreamDir will output json newlines from channel and save in dir
func NewPullRequestCommentOutputStreamDir(dir string, ch chan PullRequestComment, errors chan<- error, transforms ...TransformPullRequestCommentFunc) <-chan bool {
	fp := filepath.Join(dir, "/sourcecode/pull_request_comment\\.json(\\.gz)?$")
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
	return NewPullRequestCommentOutputStream(gz, ch, errors, transforms...)
}

// NewPullRequestCommentOutputStream will output json newlines from channel to the stream
func NewPullRequestCommentOutputStream(stream io.WriteCloser, ch chan PullRequestComment, errors chan<- error, transforms ...TransformPullRequestCommentFunc) <-chan bool {
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
