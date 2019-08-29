// DO NOT EDIT -- generated code

// Package cicd - the system which contains CI/CD
package cicd

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
)

const (
	// BuildTopic is the default topic name
	BuildTopic datamodel.TopicNameType = "cicd_Build_topic"

	// BuildStream is the default stream name
	BuildStream datamodel.TopicNameType = "cicd_Build_stream"

	// BuildTable is the default table name
	BuildTable datamodel.TopicNameType = "cicd_build"

	// BuildModelName is the model name
	BuildModelName datamodel.ModelNameType = "cicd.Build"
)

const (
	// BuildAutomatedColumn is the automated column name
	BuildAutomatedColumn = "automated"
	// BuildCommitShaColumn is the commit_sha column name
	BuildCommitShaColumn = "commit_sha"
	// BuildCustomerIDColumn is the customer_id column name
	BuildCustomerIDColumn = "customer_id"
	// BuildEndDateColumn is the end_date column name
	BuildEndDateColumn = "end_date"
	// BuildEndDateColumnEpochColumn is the epoch column property of the EndDate name
	BuildEndDateColumnEpochColumn = "end_date->epoch"
	// BuildEndDateColumnOffsetColumn is the offset column property of the EndDate name
	BuildEndDateColumnOffsetColumn = "end_date->offset"
	// BuildEndDateColumnRfc3339Column is the rfc3339 column property of the EndDate name
	BuildEndDateColumnRfc3339Column = "end_date->rfc3339"
	// BuildEnvironmentColumn is the environment column name
	BuildEnvironmentColumn = "environment"
	// BuildIDColumn is the id column name
	BuildIDColumn = "id"
	// BuildRefIDColumn is the ref_id column name
	BuildRefIDColumn = "ref_id"
	// BuildRefTypeColumn is the ref_type column name
	BuildRefTypeColumn = "ref_type"
	// BuildRepoNameColumn is the repo_name column name
	BuildRepoNameColumn = "repo_name"
	// BuildStartDateColumn is the start_date column name
	BuildStartDateColumn = "start_date"
	// BuildStartDateColumnEpochColumn is the epoch column property of the StartDate name
	BuildStartDateColumnEpochColumn = "start_date->epoch"
	// BuildStartDateColumnOffsetColumn is the offset column property of the StartDate name
	BuildStartDateColumnOffsetColumn = "start_date->offset"
	// BuildStartDateColumnRfc3339Column is the rfc3339 column property of the StartDate name
	BuildStartDateColumnRfc3339Column = "start_date->rfc3339"
	// BuildStatusColumn is the status column name
	BuildStatusColumn = "status"
	// BuildUpdatedAtColumn is the updated_ts column name
	BuildUpdatedAtColumn = "updated_ts"
	// BuildURLColumn is the url column name
	BuildURLColumn = "url"
)

// BuildEndDate represents the object structure for end_date
type BuildEndDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toBuildEndDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toBuildEndDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *BuildEndDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *BuildEndDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toBuildEndDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toBuildEndDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toBuildEndDateObject(o.Rfc3339, isavro, false, "string"),
	}
}

func (o *BuildEndDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *BuildEndDate) FromMap(kv map[string]interface{}) {

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

// BuildEnvironment is the enumeration type for environment
type BuildEnvironment int32

// String returns the string value for Environment
func (v BuildEnvironment) String() string {
	switch int32(v) {
	case 0:
		return "PRODUCTION"
	case 1:
		return "DEVELOPMENT"
	case 2:
		return "BETA"
	case 3:
		return "STAGING"
	case 4:
		return "TEST"
	case 5:
		return "OTHER"
	}
	return "unset"
}

const (
	// EnvironmentProduction is the enumeration value for production
	BuildEnvironmentProduction BuildEnvironment = 0
	// EnvironmentDevelopment is the enumeration value for development
	BuildEnvironmentDevelopment BuildEnvironment = 1
	// EnvironmentBeta is the enumeration value for beta
	BuildEnvironmentBeta BuildEnvironment = 2
	// EnvironmentStaging is the enumeration value for staging
	BuildEnvironmentStaging BuildEnvironment = 3
	// EnvironmentTest is the enumeration value for test
	BuildEnvironmentTest BuildEnvironment = 4
	// EnvironmentOther is the enumeration value for other
	BuildEnvironmentOther BuildEnvironment = 5
)

// BuildStartDate represents the object structure for start_date
type BuildStartDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toBuildStartDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toBuildStartDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *BuildStartDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *BuildStartDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toBuildStartDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toBuildStartDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toBuildStartDateObject(o.Rfc3339, isavro, false, "string"),
	}
}

func (o *BuildStartDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *BuildStartDate) FromMap(kv map[string]interface{}) {

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

// BuildStatus is the enumeration type for status
type BuildStatus int32

// String returns the string value for Status
func (v BuildStatus) String() string {
	switch int32(v) {
	case 0:
		return "PASS"
	case 1:
		return "FAIL"
	case 2:
		return "CANCEL"
	}
	return "unset"
}

const (
	// StatusPass is the enumeration value for pass
	BuildStatusPass BuildStatus = 0
	// StatusFail is the enumeration value for fail
	BuildStatusFail BuildStatus = 1
	// StatusCancel is the enumeration value for cancel
	BuildStatusCancel BuildStatus = 2
)

// Build an individual build
type Build struct {
	// Automated if the build was automated or manual
	Automated bool `json:"automated" bson:"automated" yaml:"automated" faker:"-"`
	// CommitSha the commit sha for the commit that triggered the build
	CommitSha string `json:"commit_sha" bson:"commit_sha" yaml:"commit_sha" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// EndDate the date when the build finished
	EndDate BuildEndDate `json:"end_date" bson:"end_date" yaml:"end_date" faker:"-"`
	// Environment the environment for the build
	Environment BuildEnvironment `json:"environment" bson:"environment" yaml:"environment" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RepoName the name of the repo
	RepoName string `json:"repo_name" bson:"repo_name" yaml:"repo_name" faker:"-"`
	// StartDate the date when the build started
	StartDate BuildStartDate `json:"start_date" bson:"start_date" yaml:"start_date" faker:"-"`
	// Status the status of the build
	Status BuildStatus `json:"status" bson:"status" yaml:"status" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// URL the url to the build status page
	URL *string `json:"url" bson:"url" yaml:"url" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Build)(nil)

func toBuildObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toBuildObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *Build:
		return v.ToMap(isavro)

	case BuildEndDate:
		return v.ToMap(isavro)

	case BuildEnvironment:
		return v.String()

	case BuildStartDate:
		return v.ToMap(isavro)

	case BuildStatus:
		return v.String()

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

// String returns a string representation of Build
func (o *Build) String() string {
	return fmt.Sprintf("cicd.Build<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Build) GetTopicName() datamodel.TopicNameType {
	return BuildTopic
}

// GetModelName returns the name of the model
func (o *Build) GetModelName() datamodel.ModelNameType {
	return BuildModelName
}

func (o *Build) setDefaults(frommap bool) {
	if o.URL == nil {
		o.URL = &emptyString
	}

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Build", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Build) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Build) GetTopicKey() string {
	var i interface{} = o.CustomerID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Build) GetTimestamp() time.Time {
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
	panic("not sure how to handle the date time format for Build")
}

// GetRefID returns the RefID for the object
func (o *Build) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Build) IsMaterialized() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Build) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Build) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *Build) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = BuildModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Build) GetTopicConfig() *datamodel.ModelTopicConfig {
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
		Key:               "customer_id",
		Timestamp:         "updated_ts",
		NumPartitions:     8,
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetStateKey returns a key for use in state store
func (o *Build) GetStateKey() string {
	key := "customer_id"
	return fmt.Sprintf("%s_%s", key, o.GetID())
}

// GetCustomerID will return the customer_id
func (o *Build) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of Build
func (o *Build) Clone() datamodel.Model {
	c := new(Build)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Build) Anon() datamodel.Model {
	c := new(Build)
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
func (o *Build) MarshalBinary() ([]byte, error) {
	return o.MarshalJSON()
}

func (o *Build) UnmarshalBinary(data []byte) error {
	return o.UnmarshalJSON(data)
}

// MarshalJSON returns the bytes for marshaling to json
func (o *Build) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Build) UnmarshalJSON(data []byte) error {
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

var cachedCodecBuild *goavro.Codec
var cachedCodecBuildLock sync.Mutex

// GetAvroCodec returns the avro codec for this model
func (o *Build) GetAvroCodec() *goavro.Codec {
	cachedCodecBuildLock.Lock()
	if cachedCodecBuild == nil {
		c, err := GetBuildAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecBuild = c
	}
	cachedCodecBuildLock.Unlock()
	return cachedCodecBuild
}

// ToAvroBinary returns the data as Avro binary data
func (o *Build) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *Build) FromAvroBinary(value []byte) error {
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
func (o *Build) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Build objects are equal
func (o *Build) IsEqual(other *Build) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Build) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	o.setDefaults(false)
	return map[string]interface{}{
		"automated":   toBuildObject(o.Automated, isavro, false, "boolean"),
		"commit_sha":  toBuildObject(o.CommitSha, isavro, false, "string"),
		"customer_id": toBuildObject(o.CustomerID, isavro, false, "string"),
		"end_date":    toBuildObject(o.EndDate, isavro, false, "end_date"),
		"environment": toBuildObject(o.Environment, isavro, false, "environment"),
		"id":          toBuildObject(o.ID, isavro, false, "string"),
		"ref_id":      toBuildObject(o.RefID, isavro, false, "string"),
		"ref_type":    toBuildObject(o.RefType, isavro, false, "string"),
		"repo_name":   toBuildObject(o.RepoName, isavro, false, "string"),
		"start_date":  toBuildObject(o.StartDate, isavro, false, "start_date"),
		"status":      toBuildObject(o.Status, isavro, false, "status"),
		"updated_ts":  toBuildObject(o.UpdatedAt, isavro, false, "long"),
		"url":         toBuildObject(o.URL, isavro, true, "string"),
		"hashcode":    toBuildObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *Build) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["automated"].(bool); ok {
		o.Automated = val
	} else {
		if val, ok := kv["automated"]; ok {
			if val == nil {
				o.Automated = number.ToBoolAny(nil)
			} else {
				o.Automated = number.ToBoolAny(val)
			}
		}
	}

	if val, ok := kv["commit_sha"].(string); ok {
		o.CommitSha = val
	} else {
		if val, ok := kv["commit_sha"]; ok {
			if val == nil {
				o.CommitSha = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.CommitSha = fmt.Sprintf("%v", val)
			}
		}
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

	if val, ok := kv["end_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.EndDate.FromMap(kv)
		} else if sv, ok := val.(BuildEndDate); ok {
			// struct
			o.EndDate = sv
		} else if sp, ok := val.(*BuildEndDate); ok {
			// struct pointer
			o.EndDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.EndDate.Epoch = dt.Epoch
			o.EndDate.Rfc3339 = dt.Rfc3339
			o.EndDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.EndDate.Epoch = dt.Epoch
			o.EndDate.Rfc3339 = dt.Rfc3339
			o.EndDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.EndDate.Epoch = dt.Epoch
				o.EndDate.Rfc3339 = dt.Rfc3339
				o.EndDate.Offset = dt.Offset
			}
		}
	} else {
		o.EndDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["environment"].(BuildEnvironment); ok {
		o.Environment = val
	} else {
		if em, ok := kv["environment"].(map[string]interface{}); ok {
			ev := em["cicd.environment"].(string)
			switch ev {
			case "production", "PRODUCTION":
				o.Environment = 0
			case "development", "DEVELOPMENT":
				o.Environment = 1
			case "beta", "BETA":
				o.Environment = 2
			case "staging", "STAGING":
				o.Environment = 3
			case "test", "TEST":
				o.Environment = 4
			case "other", "OTHER":
				o.Environment = 5
			}
		}
		if em, ok := kv["environment"].(string); ok {
			switch em {
			case "production", "PRODUCTION":
				o.Environment = 0
			case "development", "DEVELOPMENT":
				o.Environment = 1
			case "beta", "BETA":
				o.Environment = 2
			case "staging", "STAGING":
				o.Environment = 3
			case "test", "TEST":
				o.Environment = 4
			case "other", "OTHER":
				o.Environment = 5
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

	if val, ok := kv["repo_name"].(string); ok {
		o.RepoName = val
	} else {
		if val, ok := kv["repo_name"]; ok {
			if val == nil {
				o.RepoName = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.RepoName = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["start_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.StartDate.FromMap(kv)
		} else if sv, ok := val.(BuildStartDate); ok {
			// struct
			o.StartDate = sv
		} else if sp, ok := val.(*BuildStartDate); ok {
			// struct pointer
			o.StartDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.StartDate.Epoch = dt.Epoch
			o.StartDate.Rfc3339 = dt.Rfc3339
			o.StartDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.StartDate.Epoch = dt.Epoch
			o.StartDate.Rfc3339 = dt.Rfc3339
			o.StartDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.StartDate.Epoch = dt.Epoch
				o.StartDate.Rfc3339 = dt.Rfc3339
				o.StartDate.Offset = dt.Offset
			}
		}
	} else {
		o.StartDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["status"].(BuildStatus); ok {
		o.Status = val
	} else {
		if em, ok := kv["status"].(map[string]interface{}); ok {
			ev := em["cicd.status"].(string)
			switch ev {
			case "pass", "PASS":
				o.Status = 0
			case "fail", "FAIL":
				o.Status = 1
			case "cancel", "CANCEL":
				o.Status = 2
			}
		}
		if em, ok := kv["status"].(string); ok {
			switch em {
			case "pass", "PASS":
				o.Status = 0
			case "fail", "FAIL":
				o.Status = 1
			case "cancel", "CANCEL":
				o.Status = 2
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

	if val, ok := kv["url"].(*string); ok {
		o.URL = val
	} else if val, ok := kv["url"].(string); ok {
		o.URL = &val
	} else {
		if val, ok := kv["url"]; ok {
			if val == nil {
				o.URL = pstrings.Pointer("")
			} else {
				// if coming in as avro union, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.URL = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *Build) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Automated)
	args = append(args, o.CommitSha)
	args = append(args, o.CustomerID)
	args = append(args, o.EndDate)
	args = append(args, o.Environment)
	args = append(args, o.ID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RepoName)
	args = append(args, o.StartDate)
	args = append(args, o.Status)
	args = append(args, o.UpdatedAt)
	args = append(args, o.URL)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetBuildAvroSchemaSpec creates the avro schema specification for Build
func GetBuildAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "cicd",
		"name":      "Build",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "hashcode",
				"type": "string",
			},
			map[string]interface{}{
				"name": "automated",
				"type": "boolean",
			},
			map[string]interface{}{
				"name": "commit_sha",
				"type": "string",
			},
			map[string]interface{}{
				"name": "customer_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "end_date",
				"type": map[string]interface{}{"doc": "the date when the build finished", "fields": []interface{}{map[string]interface{}{"doc": "the date in epoch format", "name": "epoch", "type": "long"}, map[string]interface{}{"doc": "the timezone offset from GMT", "name": "offset", "type": "long"}, map[string]interface{}{"doc": "the date in RFC3339 format", "name": "rfc3339", "type": "string"}}, "name": "end_date", "type": "record"},
			},
			map[string]interface{}{
				"name": "environment",
				"type": map[string]interface{}{
					"type":    "enum",
					"name":    "environment",
					"symbols": []interface{}{"PRODUCTION", "DEVELOPMENT", "BETA", "STAGING", "TEST", "OTHER"},
				},
			},
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
				"name": "repo_name",
				"type": "string",
			},
			map[string]interface{}{
				"name": "start_date",
				"type": map[string]interface{}{"doc": "the date when the build started", "fields": []interface{}{map[string]interface{}{"doc": "the date in epoch format", "name": "epoch", "type": "long"}, map[string]interface{}{"doc": "the timezone offset from GMT", "name": "offset", "type": "long"}, map[string]interface{}{"doc": "the date in RFC3339 format", "name": "rfc3339", "type": "string"}}, "name": "start_date", "type": "record"},
			},
			map[string]interface{}{
				"name": "status",
				"type": map[string]interface{}{
					"type":    "enum",
					"name":    "status",
					"symbols": []interface{}{"PASS", "FAIL", "CANCEL"},
				},
			},
			map[string]interface{}{
				"name": "updated_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name":    "url",
				"type":    []interface{}{"null", "string"},
				"default": nil,
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *Build) GetEventAPIConfig() datamodel.EventAPIConfig {
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

// GetBuildAvroSchema creates the avro schema for Build
func GetBuildAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetBuildAvroSchemaSpec())
}

// TransformBuildFunc is a function for transforming Build during processing
type TransformBuildFunc func(input *Build) (*Build, error)

// NewBuildPipe creates a pipe for processing Build items
func NewBuildPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformBuildFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewBuildInputStream(input, errors)
	var stream chan Build
	if len(transforms) > 0 {
		stream = make(chan Build, 1000)
	} else {
		stream = inch
	}
	outdone := NewBuildOutputStream(output, stream, errors)
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

// NewBuildInputStreamDir creates a channel for reading Build as JSON newlines from a directory of files
func NewBuildInputStreamDir(dir string, errors chan<- error, transforms ...TransformBuildFunc) (chan Build, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/cicd/build\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan Build)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for build")
		ch := make(chan Build)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewBuildInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan Build)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewBuildInputStreamFile creates an channel for reading Build as JSON newlines from filename
func NewBuildInputStreamFile(filename string, errors chan<- error, transforms ...TransformBuildFunc) (chan Build, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan Build)
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
			ch := make(chan Build)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewBuildInputStream(f, errors, transforms...)
}

// NewBuildInputStream creates an channel for reading Build as JSON newlines from stream
func NewBuildInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformBuildFunc) (chan Build, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan Build, 1000)
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
			var item Build
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

// NewBuildOutputStreamDir will output json newlines from channel and save in dir
func NewBuildOutputStreamDir(dir string, ch chan Build, errors chan<- error, transforms ...TransformBuildFunc) <-chan bool {
	fp := filepath.Join(dir, "/cicd/build\\.json(\\.gz)?$")
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
	return NewBuildOutputStream(gz, ch, errors, transforms...)
}

// NewBuildOutputStream will output json newlines from channel to the stream
func NewBuildOutputStream(stream io.WriteCloser, ch chan Build, errors chan<- error, transforms ...TransformBuildFunc) <-chan bool {
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

// BuildSendEvent is an event detail for sending data
type BuildSendEvent struct {
	Build   *Build
	headers map[string]string
	time    time.Time
	key     string
}

var _ datamodel.ModelSendEvent = (*BuildSendEvent)(nil)

// Key is the key to use for the message
func (e *BuildSendEvent) Key() string {
	if e.key == "" {
		return e.Build.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *BuildSendEvent) Object() datamodel.Model {
	return e.Build
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *BuildSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *BuildSendEvent) Timestamp() time.Time {
	return e.time
}

// BuildSendEventOpts is a function handler for setting opts
type BuildSendEventOpts func(o *BuildSendEvent)

// WithBuildSendEventKey sets the key value to a value different than the object ID
func WithBuildSendEventKey(key string) BuildSendEventOpts {
	return func(o *BuildSendEvent) {
		o.key = key
	}
}

// WithBuildSendEventTimestamp sets the timestamp value
func WithBuildSendEventTimestamp(tv time.Time) BuildSendEventOpts {
	return func(o *BuildSendEvent) {
		o.time = tv
	}
}

// WithBuildSendEventHeader sets the timestamp value
func WithBuildSendEventHeader(key, value string) BuildSendEventOpts {
	return func(o *BuildSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewBuildSendEvent returns a new BuildSendEvent instance
func NewBuildSendEvent(o *Build, opts ...BuildSendEventOpts) *BuildSendEvent {
	res := &BuildSendEvent{
		Build: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewBuildProducer will stream data from the channel
func NewBuildProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
	done := make(chan bool, 1)
	emptyTime := time.Unix(0, 0)
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
				if object, ok := item.Object().(*Build); ok {
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
					msg := eventing.Message{
						Encoding:  eventing.AvroEncoding,
						Key:       item.Key(),
						Value:     binary,
						Codec:     codec,
						Headers:   headers,
						Timestamp: tv,
						Partition: -1, // select any partition based on partitioner strategy in kafka
						Topic:     object.GetTopicName().String(),
					}
					if err := producer.Send(ctx, msg); err != nil {
						errors <- fmt.Errorf("error sending %s. %v", object.String(), err)
					}
				} else {
					errors <- fmt.Errorf("invalid event received. expected an object of type cicd.Build but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewBuildConsumer will stream data from the topic into the provided channel
func NewBuildConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object Build
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into cicd.Build: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avro data into cicd.Build: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for cicd.Build")
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

			ch <- &BuildReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object Build
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &BuildReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// BuildReceiveEvent is an event detail for receiving data
type BuildReceiveEvent struct {
	Build   *Build
	message eventing.Message
	eof     bool
}

var _ datamodel.ModelReceiveEvent = (*BuildReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *BuildReceiveEvent) Object() datamodel.Model {
	return e.Build
}

// Message returns the underlying message data for the event
func (e *BuildReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *BuildReceiveEvent) EOF() bool {
	return e.eof
}

// BuildProducer implements the datamodel.ModelEventProducer
type BuildProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*BuildProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *BuildProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *BuildProducer) Close() error {
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
func (o *Build) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *Build) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &BuildProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewBuildProducer(newctx, producer, ch, errors, empty),
	}
}

// NewBuildProducerChannel returns a channel which can be used for producing Model events
func NewBuildProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewBuildProducerChannelSize(producer, 0, errors)
}

// NewBuildProducerChannelSize returns a channel which can be used for producing Model events
func NewBuildProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &BuildProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewBuildProducer(newctx, producer, ch, errors, empty),
	}
}

// BuildConsumer implements the datamodel.ModelEventConsumer
type BuildConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*BuildConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *BuildConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *BuildConsumer) Close() error {
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
func (o *Build) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &BuildConsumer{
		ch:       ch,
		callback: NewBuildConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewBuildConsumerChannel returns a consumer channel which can be used to consume Model events
func NewBuildConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &BuildConsumer{
		ch:       ch,
		callback: NewBuildConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
