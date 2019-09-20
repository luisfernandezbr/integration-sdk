// DO NOT EDIT -- generated code

// Package sourcecode - the system which contains source code
package sourcecode

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/bxcodec/faker"
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
)

const (
	// PullRequestCommentTopic is the default topic name
	PullRequestCommentTopic datamodel.TopicNameType = "sourcecode_PullRequestComment_topic"

	// PullRequestCommentStream is the default stream name
	PullRequestCommentStream datamodel.TopicNameType = "sourcecode_PullRequestComment_stream"

	// PullRequestCommentTable is the default table name
	PullRequestCommentTable datamodel.TopicNameType = "sourcecode_pullrequestcomment"

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
	// PullRequestCommentUpdatedAtColumn is the updated_ts column name
	PullRequestCommentUpdatedAtColumn = "updated_ts"
	// PullRequestCommentURLColumn is the url column name
	PullRequestCommentURLColumn = "url"
	// PullRequestCommentUserRefIDColumn is the user_ref_id column name
	PullRequestCommentUserRefIDColumn = "user_ref_id"
)

// PullRequestCommentCreatedDate represents the object structure for created_date
type PullRequestCommentCreatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toPullRequestCommentCreatedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *PullRequestCommentCreatedDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *PullRequestCommentCreatedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toPullRequestCommentCreatedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toPullRequestCommentCreatedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toPullRequestCommentCreatedDateObject(o.Rfc3339, false),
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
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toPullRequestCommentUpdatedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *PullRequestCommentUpdatedDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *PullRequestCommentUpdatedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toPullRequestCommentUpdatedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toPullRequestCommentUpdatedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toPullRequestCommentUpdatedDateObject(o.Rfc3339, false),
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
	Body string `json:"body" codec:"body" bson:"body" yaml:"body" faker:"commit_message"`
	// CreatedDate the timestamp in UTC that the comment was created
	CreatedDate PullRequestCommentCreatedDate `json:"created_date" codec:"created_date" bson:"created_date" yaml:"created_date" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// PullRequestID the pull request this comment is associated with
	PullRequestID string `json:"pull_request_id" codec:"pull_request_id" bson:"pull_request_id" yaml:"pull_request_id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RepoID the unique id for the repo
	RepoID string `json:"repo_id" codec:"repo_id" bson:"repo_id" yaml:"repo_id" faker:"-"`
	// UpdatedDate the timestamp in UTC that the comment was closed
	UpdatedDate PullRequestCommentUpdatedDate `json:"updated_date" codec:"updated_date" bson:"updated_date" yaml:"updated_date" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// URL the URL to the source system for this comment
	URL string `json:"url" codec:"url" bson:"url" yaml:"url" faker:"url"`
	// UserRefID the user ref_id in the source system
	UserRefID string `json:"user_ref_id" codec:"user_ref_id" bson:"user_ref_id" yaml:"user_ref_id" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*PullRequestComment)(nil)

func toPullRequestCommentObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *PullRequestComment:
		return v.ToMap()

	case PullRequestCommentCreatedDate:
		return v.ToMap()

	case PullRequestCommentUpdatedDate:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of PullRequestComment
func (o *PullRequestComment) String() string {
	return fmt.Sprintf("sourcecode.PullRequestComment<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *PullRequestComment) GetTopicName() datamodel.TopicNameType {
	return PullRequestCommentTopic
}

// GetModelName returns the name of the model
func (o *PullRequestComment) GetModelName() datamodel.ModelNameType {
	return PullRequestCommentModelName
}

// GetStreamName returns the name of the stream
func (o *PullRequestComment) GetStreamName() string {
	return PullRequestCommentStream.String()
}

// GetTableName returns the name of the table
func (o *PullRequestComment) GetTableName() string {
	return PullRequestCommentTable.String()
}

// NewPullRequestCommentID provides a template for generating an ID field for PullRequestComment
func NewPullRequestCommentID(customerID string, refID string, refType string, RepoID string) string {
	return hash.Values(customerID, refID, refType, RepoID)
}

func (o *PullRequestComment) setDefaults(frommap bool) {

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
func (o *PullRequestComment) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *PullRequestComment) GetTopicKey() string {
	var i interface{} = o.RepoID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *PullRequestComment) GetTimestamp() time.Time {
	var dt interface{} = o.CreatedDate
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
	case PullRequestCommentCreatedDate:
		return datetime.DateFromEpoch(v.Epoch)
	}
	panic("not sure how to handle the date time format for PullRequestComment")
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
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *PullRequestComment) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = PullRequestCommentModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *PullRequestComment) GetTopicConfig() *datamodel.ModelTopicConfig {
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
		Timestamp:         "created_date",
		NumPartitions:     8,
		CleanupPolicy:     datamodel.CleanupPolicy("compact"),
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
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
func (o *PullRequestComment) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"body":            toPullRequestCommentObject(o.Body, false),
		"created_date":    toPullRequestCommentObject(o.CreatedDate, false),
		"customer_id":     toPullRequestCommentObject(o.CustomerID, false),
		"id":              toPullRequestCommentObject(o.ID, false),
		"pull_request_id": toPullRequestCommentObject(o.PullRequestID, false),
		"ref_id":          toPullRequestCommentObject(o.RefID, false),
		"ref_type":        toPullRequestCommentObject(o.RefType, false),
		"repo_id":         toPullRequestCommentObject(o.RepoID, false),
		"updated_date":    toPullRequestCommentObject(o.UpdatedDate, false),
		"updated_ts":      toPullRequestCommentObject(o.UpdatedAt, false),
		"url":             toPullRequestCommentObject(o.URL, false),
		"user_ref_id":     toPullRequestCommentObject(o.UserRefID, false),
		"hashcode":        toPullRequestCommentObject(o.Hashcode, false),
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
	args = append(args, o.UpdatedAt)
	args = append(args, o.URL)
	args = append(args, o.UserRefID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
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
