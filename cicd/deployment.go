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
	// DeploymentTopic is the default topic name
	DeploymentTopic datamodel.TopicNameType = "cicd_Deployment_topic"

	// DeploymentStream is the default stream name
	DeploymentStream datamodel.TopicNameType = "cicd_Deployment_stream"

	// DeploymentTable is the default table name
	DeploymentTable datamodel.TopicNameType = "cicd_Deployment"

	// DeploymentModelName is the model name
	DeploymentModelName datamodel.ModelNameType = "cicd.Deployment"
)

const (
	// DeploymentAutomatedColumn is the automated column name
	DeploymentAutomatedColumn = "automated"
	// DeploymentBuildRefIDColumn is the build_ref_id column name
	DeploymentBuildRefIDColumn = "build_ref_id"
	// DeploymentCommitShaColumn is the commit_sha column name
	DeploymentCommitShaColumn = "commit_sha"
	// DeploymentCustomerIDColumn is the customer_id column name
	DeploymentCustomerIDColumn = "customer_id"
	// DeploymentEndDateColumn is the end_date column name
	DeploymentEndDateColumn = "end_date"
	// DeploymentEndDateColumnEpochColumn is the epoch column property of the EndDate name
	DeploymentEndDateColumnEpochColumn = "end_date->epoch"
	// DeploymentEndDateColumnOffsetColumn is the offset column property of the EndDate name
	DeploymentEndDateColumnOffsetColumn = "end_date->offset"
	// DeploymentEndDateColumnRfc3339Column is the rfc3339 column property of the EndDate name
	DeploymentEndDateColumnRfc3339Column = "end_date->rfc3339"
	// DeploymentEnvironmentColumn is the environment column name
	DeploymentEnvironmentColumn = "environment"
	// DeploymentIDColumn is the id column name
	DeploymentIDColumn = "id"
	// DeploymentRefIDColumn is the ref_id column name
	DeploymentRefIDColumn = "ref_id"
	// DeploymentRefTypeColumn is the ref_type column name
	DeploymentRefTypeColumn = "ref_type"
	// DeploymentRepoNameColumn is the repo_name column name
	DeploymentRepoNameColumn = "repo_name"
	// DeploymentStartDateColumn is the start_date column name
	DeploymentStartDateColumn = "start_date"
	// DeploymentStartDateColumnEpochColumn is the epoch column property of the StartDate name
	DeploymentStartDateColumnEpochColumn = "start_date->epoch"
	// DeploymentStartDateColumnOffsetColumn is the offset column property of the StartDate name
	DeploymentStartDateColumnOffsetColumn = "start_date->offset"
	// DeploymentStartDateColumnRfc3339Column is the rfc3339 column property of the StartDate name
	DeploymentStartDateColumnRfc3339Column = "start_date->rfc3339"
	// DeploymentStatusColumn is the status column name
	DeploymentStatusColumn = "status"
	// DeploymentUpdatedAtColumn is the updated_ts column name
	DeploymentUpdatedAtColumn = "updated_ts"
	// DeploymentURLColumn is the url column name
	DeploymentURLColumn = "url"
)

// DeploymentEndDate represents the object structure for end_date
type DeploymentEndDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toDeploymentEndDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toDeploymentEndDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *DeploymentEndDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *DeploymentEndDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toDeploymentEndDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toDeploymentEndDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toDeploymentEndDateObject(o.Rfc3339, isavro, false, "string"),
	}
}

func (o *DeploymentEndDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *DeploymentEndDate) FromMap(kv map[string]interface{}) {

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

// DeploymentEnvironment is the enumeration type for environment
type DeploymentEnvironment int32

// String returns the string value for Environment
func (v DeploymentEnvironment) String() string {
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
	DeploymentEnvironmentProduction DeploymentEnvironment = 0
	// EnvironmentDevelopment is the enumeration value for development
	DeploymentEnvironmentDevelopment DeploymentEnvironment = 1
	// EnvironmentBeta is the enumeration value for beta
	DeploymentEnvironmentBeta DeploymentEnvironment = 2
	// EnvironmentStaging is the enumeration value for staging
	DeploymentEnvironmentStaging DeploymentEnvironment = 3
	// EnvironmentTest is the enumeration value for test
	DeploymentEnvironmentTest DeploymentEnvironment = 4
	// EnvironmentOther is the enumeration value for other
	DeploymentEnvironmentOther DeploymentEnvironment = 5
)

// DeploymentStartDate represents the object structure for start_date
type DeploymentStartDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toDeploymentStartDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toDeploymentStartDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *DeploymentStartDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *DeploymentStartDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toDeploymentStartDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toDeploymentStartDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toDeploymentStartDateObject(o.Rfc3339, isavro, false, "string"),
	}
}

func (o *DeploymentStartDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *DeploymentStartDate) FromMap(kv map[string]interface{}) {

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

// DeploymentStatus is the enumeration type for status
type DeploymentStatus int32

// String returns the string value for Status
func (v DeploymentStatus) String() string {
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
	DeploymentStatusPass DeploymentStatus = 0
	// StatusFail is the enumeration value for fail
	DeploymentStatusFail DeploymentStatus = 1
	// StatusCancel is the enumeration value for cancel
	DeploymentStatusCancel DeploymentStatus = 2
)

// Deployment an individual deployment
type Deployment struct {
	// Automated if the deployment was automated or manual
	Automated bool `json:"automated" bson:"automated" yaml:"automated" faker:"-"`
	// BuildRefID the build id that is associated with the deployment (if any)
	BuildRefID string `json:"build_ref_id" bson:"build_ref_id" yaml:"build_ref_id" faker:"-"`
	// CommitSha the commit sha for the commit that triggered the deployment
	CommitSha string `json:"commit_sha" bson:"commit_sha" yaml:"commit_sha" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// EndDate the date when the deployment finished
	EndDate DeploymentEndDate `json:"end_date" bson:"end_date" yaml:"end_date" faker:"-"`
	// Environment the environment for the deployment
	Environment DeploymentEnvironment `json:"environment" bson:"environment" yaml:"environment" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RepoName the name of the repo
	RepoName string `json:"repo_name" bson:"repo_name" yaml:"repo_name" faker:"-"`
	// StartDate the date when the deployment started
	StartDate DeploymentStartDate `json:"start_date" bson:"start_date" yaml:"start_date" faker:"-"`
	// Status the status of the deployment
	Status DeploymentStatus `json:"status" bson:"status" yaml:"status" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// URL the url to the deployment status page
	URL *string `json:"url" bson:"url" yaml:"url" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Deployment)(nil)

func toDeploymentObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toDeploymentObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *Deployment:
		return v.ToMap(isavro)

	case DeploymentEndDate:
		return v.ToMap(isavro)

	case DeploymentEnvironment:
		return v.String()

	case DeploymentStartDate:
		return v.ToMap(isavro)

	case DeploymentStatus:
		return v.String()

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

// String returns a string representation of Deployment
func (o *Deployment) String() string {
	return fmt.Sprintf("cicd.Deployment<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Deployment) GetTopicName() datamodel.TopicNameType {
	return DeploymentTopic
}

// GetModelName returns the name of the model
func (o *Deployment) GetModelName() datamodel.ModelNameType {
	return DeploymentModelName
}

func (o *Deployment) setDefaults(frommap bool) {
	if o.URL == nil {
		o.URL = &emptyString
	}

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Deployment", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Deployment) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Deployment) GetTopicKey() string {
	var i interface{} = o.ID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Deployment) GetTimestamp() time.Time {
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
	panic("not sure how to handle the date time format for Deployment")
}

// GetRefID returns the RefID for the object
func (o *Deployment) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Deployment) IsMaterialized() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Deployment) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Deployment) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *Deployment) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = DeploymentModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Deployment) GetTopicConfig() *datamodel.ModelTopicConfig {
	retention, err := time.ParseDuration("168h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 168h0m0s. " + err.Error())
	}

	ttl, err := time.ParseDuration("0s")
	if err != nil {
		ttl = 0
	}
	if ttl == 0 && retention != 0 {
		ttl = retention // they should be the same if not set
	}
	return &datamodel.ModelTopicConfig{
		Key:               "id",
		Timestamp:         "updated_ts",
		NumPartitions:     8,
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetStateKey returns a key for use in state store
func (o *Deployment) GetStateKey() string {
	key := "id"
	return fmt.Sprintf("%s_%s", key, o.GetID())
}

// GetCustomerID will return the customer_id
func (o *Deployment) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of Deployment
func (o *Deployment) Clone() datamodel.Model {
	c := new(Deployment)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Deployment) Anon() datamodel.Model {
	c := new(Deployment)
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
func (o *Deployment) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Deployment) UnmarshalJSON(data []byte) error {
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

var cachedCodecDeployment *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *Deployment) GetAvroCodec() *goavro.Codec {
	if cachedCodecDeployment == nil {
		c, err := GetDeploymentAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecDeployment = c
	}
	return cachedCodecDeployment
}

// ToAvroBinary returns the data as Avro binary data
func (o *Deployment) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *Deployment) FromAvroBinary(value []byte) error {
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
func (o *Deployment) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Deployment objects are equal
func (o *Deployment) IsEqual(other *Deployment) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Deployment) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	o.setDefaults(false)
	return map[string]interface{}{
		"automated":    toDeploymentObject(o.Automated, isavro, false, "boolean"),
		"build_ref_id": toDeploymentObject(o.BuildRefID, isavro, false, "string"),
		"commit_sha":   toDeploymentObject(o.CommitSha, isavro, false, "string"),
		"customer_id":  toDeploymentObject(o.CustomerID, isavro, false, "string"),
		"end_date":     toDeploymentObject(o.EndDate, isavro, false, "end_date"),
		"environment":  toDeploymentObject(o.Environment, isavro, false, "environment"),
		"id":           toDeploymentObject(o.ID, isavro, false, "string"),
		"ref_id":       toDeploymentObject(o.RefID, isavro, false, "string"),
		"ref_type":     toDeploymentObject(o.RefType, isavro, false, "string"),
		"repo_name":    toDeploymentObject(o.RepoName, isavro, false, "string"),
		"start_date":   toDeploymentObject(o.StartDate, isavro, false, "start_date"),
		"status":       toDeploymentObject(o.Status, isavro, false, "status"),
		"updated_ts":   toDeploymentObject(o.UpdatedAt, isavro, false, "long"),
		"url":          toDeploymentObject(o.URL, isavro, true, "string"),
		"hashcode":     toDeploymentObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *Deployment) FromMap(kv map[string]interface{}) {

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

	if val, ok := kv["build_ref_id"].(string); ok {
		o.BuildRefID = val
	} else {
		if val, ok := kv["build_ref_id"]; ok {
			if val == nil {
				o.BuildRefID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.BuildRefID = fmt.Sprintf("%v", val)
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
		} else if sv, ok := val.(DeploymentEndDate); ok {
			// struct
			o.EndDate = sv
		} else if sp, ok := val.(*DeploymentEndDate); ok {
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

	if val, ok := kv["environment"].(DeploymentEnvironment); ok {
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
		} else if sv, ok := val.(DeploymentStartDate); ok {
			// struct
			o.StartDate = sv
		} else if sp, ok := val.(*DeploymentStartDate); ok {
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

	if val, ok := kv["status"].(DeploymentStatus); ok {
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
func (o *Deployment) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Automated)
	args = append(args, o.BuildRefID)
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

// GetDeploymentAvroSchemaSpec creates the avro schema specification for Deployment
func GetDeploymentAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "cicd",
		"name":      "Deployment",
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
				"name": "build_ref_id",
				"type": "string",
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
				"type": map[string]interface{}{"doc": "the date when the deployment finished", "fields": []interface{}{map[string]interface{}{"doc": "the date in epoch format", "name": "epoch", "type": "long"}, map[string]interface{}{"doc": "the timezone offset from GMT", "name": "offset", "type": "long"}, map[string]interface{}{"doc": "the date in RFC3339 format", "name": "rfc3339", "type": "string"}}, "name": "end_date", "type": "record"},
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
				"type": map[string]interface{}{"doc": "the date when the deployment started", "fields": []interface{}{map[string]interface{}{"doc": "the date in epoch format", "name": "epoch", "type": "long"}, map[string]interface{}{"doc": "the timezone offset from GMT", "name": "offset", "type": "long"}, map[string]interface{}{"doc": "the date in RFC3339 format", "name": "rfc3339", "type": "string"}}, "name": "start_date", "type": "record"},
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
func (o *Deployment) GetEventAPIConfig() datamodel.EventAPIConfig {
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

// GetDeploymentAvroSchema creates the avro schema for Deployment
func GetDeploymentAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetDeploymentAvroSchemaSpec())
}

// TransformDeploymentFunc is a function for transforming Deployment during processing
type TransformDeploymentFunc func(input *Deployment) (*Deployment, error)

// NewDeploymentPipe creates a pipe for processing Deployment items
func NewDeploymentPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformDeploymentFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewDeploymentInputStream(input, errors)
	var stream chan Deployment
	if len(transforms) > 0 {
		stream = make(chan Deployment, 1000)
	} else {
		stream = inch
	}
	outdone := NewDeploymentOutputStream(output, stream, errors)
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

// NewDeploymentInputStreamDir creates a channel for reading Deployment as JSON newlines from a directory of files
func NewDeploymentInputStreamDir(dir string, errors chan<- error, transforms ...TransformDeploymentFunc) (chan Deployment, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/cicd/deployment\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan Deployment)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for deployment")
		ch := make(chan Deployment)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewDeploymentInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan Deployment)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewDeploymentInputStreamFile creates an channel for reading Deployment as JSON newlines from filename
func NewDeploymentInputStreamFile(filename string, errors chan<- error, transforms ...TransformDeploymentFunc) (chan Deployment, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan Deployment)
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
			ch := make(chan Deployment)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewDeploymentInputStream(f, errors, transforms...)
}

// NewDeploymentInputStream creates an channel for reading Deployment as JSON newlines from stream
func NewDeploymentInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformDeploymentFunc) (chan Deployment, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan Deployment, 1000)
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
			var item Deployment
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

// NewDeploymentOutputStreamDir will output json newlines from channel and save in dir
func NewDeploymentOutputStreamDir(dir string, ch chan Deployment, errors chan<- error, transforms ...TransformDeploymentFunc) <-chan bool {
	fp := filepath.Join(dir, "/cicd/deployment\\.json(\\.gz)?$")
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
	return NewDeploymentOutputStream(gz, ch, errors, transforms...)
}

// NewDeploymentOutputStream will output json newlines from channel to the stream
func NewDeploymentOutputStream(stream io.WriteCloser, ch chan Deployment, errors chan<- error, transforms ...TransformDeploymentFunc) <-chan bool {
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

// DeploymentSendEvent is an event detail for sending data
type DeploymentSendEvent struct {
	Deployment *Deployment
	headers    map[string]string
	time       time.Time
	key        string
}

var _ datamodel.ModelSendEvent = (*DeploymentSendEvent)(nil)

// Key is the key to use for the message
func (e *DeploymentSendEvent) Key() string {
	if e.key == "" {
		return e.Deployment.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *DeploymentSendEvent) Object() datamodel.Model {
	return e.Deployment
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *DeploymentSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *DeploymentSendEvent) Timestamp() time.Time {
	return e.time
}

// DeploymentSendEventOpts is a function handler for setting opts
type DeploymentSendEventOpts func(o *DeploymentSendEvent)

// WithDeploymentSendEventKey sets the key value to a value different than the object ID
func WithDeploymentSendEventKey(key string) DeploymentSendEventOpts {
	return func(o *DeploymentSendEvent) {
		o.key = key
	}
}

// WithDeploymentSendEventTimestamp sets the timestamp value
func WithDeploymentSendEventTimestamp(tv time.Time) DeploymentSendEventOpts {
	return func(o *DeploymentSendEvent) {
		o.time = tv
	}
}

// WithDeploymentSendEventHeader sets the timestamp value
func WithDeploymentSendEventHeader(key, value string) DeploymentSendEventOpts {
	return func(o *DeploymentSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewDeploymentSendEvent returns a new DeploymentSendEvent instance
func NewDeploymentSendEvent(o *Deployment, opts ...DeploymentSendEventOpts) *DeploymentSendEvent {
	res := &DeploymentSendEvent{
		Deployment: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewDeploymentProducer will stream data from the channel
func NewDeploymentProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
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
				if object, ok := item.Object().(*Deployment); ok {
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
						Topic:     object.GetTopicName().String(),
					}
					if err := producer.Send(ctx, msg); err != nil {
						errors <- fmt.Errorf("error sending %s. %v", object.String(), err)
					}
				} else {
					errors <- fmt.Errorf("invalid event received. expected an object of type cicd.Deployment but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewDeploymentConsumer will stream data from the topic into the provided channel
func NewDeploymentConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object Deployment
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into cicd.Deployment: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avro data into cicd.Deployment: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for cicd.Deployment")
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

			ch <- &DeploymentReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object Deployment
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &DeploymentReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// DeploymentReceiveEvent is an event detail for receiving data
type DeploymentReceiveEvent struct {
	Deployment *Deployment
	message    eventing.Message
	eof        bool
}

var _ datamodel.ModelReceiveEvent = (*DeploymentReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *DeploymentReceiveEvent) Object() datamodel.Model {
	return e.Deployment
}

// Message returns the underlying message data for the event
func (e *DeploymentReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *DeploymentReceiveEvent) EOF() bool {
	return e.eof
}

// DeploymentProducer implements the datamodel.ModelEventProducer
type DeploymentProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*DeploymentProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *DeploymentProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *DeploymentProducer) Close() error {
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
func (o *Deployment) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *Deployment) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &DeploymentProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewDeploymentProducer(newctx, producer, ch, errors, empty),
	}
}

// NewDeploymentProducerChannel returns a channel which can be used for producing Model events
func NewDeploymentProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewDeploymentProducerChannelSize(producer, 0, errors)
}

// NewDeploymentProducerChannelSize returns a channel which can be used for producing Model events
func NewDeploymentProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &DeploymentProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewDeploymentProducer(newctx, producer, ch, errors, empty),
	}
}

// DeploymentConsumer implements the datamodel.ModelEventConsumer
type DeploymentConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*DeploymentConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *DeploymentConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *DeploymentConsumer) Close() error {
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
func (o *Deployment) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &DeploymentConsumer{
		ch:       ch,
		callback: NewDeploymentConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewDeploymentConsumerChannel returns a consumer channel which can be used to consume Model events
func NewDeploymentConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &DeploymentConsumer{
		ch:       ch,
		callback: NewDeploymentConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
