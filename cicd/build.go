// DO NOT EDIT -- generated code

// Package cicd - the system which contains CI/CD
package cicd

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
	pstrings "github.com/pinpt/go-common/strings"
)

const (
	// BuildTopic is the default topic name
	BuildTopic datamodel.TopicNameType = "cicd_Build_topic"

	// BuildModelName is the model name
	BuildModelName datamodel.ModelNameType = "cicd.Build"
)

const (
	// BuildAutomatedColumn is the automated column name
	BuildAutomatedColumn = "Automated"
	// BuildCommitShaColumn is the commit_sha column name
	BuildCommitShaColumn = "CommitSha"
	// BuildCustomerIDColumn is the customer_id column name
	BuildCustomerIDColumn = "CustomerID"
	// BuildEndDateColumn is the end_date column name
	BuildEndDateColumn = "EndDate"
	// BuildEndDateColumnEpochColumn is the epoch column property of the EndDate name
	BuildEndDateColumnEpochColumn = "EndDate.Epoch"
	// BuildEndDateColumnOffsetColumn is the offset column property of the EndDate name
	BuildEndDateColumnOffsetColumn = "EndDate.Offset"
	// BuildEndDateColumnRfc3339Column is the rfc3339 column property of the EndDate name
	BuildEndDateColumnRfc3339Column = "EndDate.Rfc3339"
	// BuildEnvironmentColumn is the environment column name
	BuildEnvironmentColumn = "Environment"
	// BuildIDColumn is the id column name
	BuildIDColumn = "ID"
	// BuildRefIDColumn is the ref_id column name
	BuildRefIDColumn = "RefID"
	// BuildRefTypeColumn is the ref_type column name
	BuildRefTypeColumn = "RefType"
	// BuildRepoNameColumn is the repo_name column name
	BuildRepoNameColumn = "RepoName"
	// BuildStartDateColumn is the start_date column name
	BuildStartDateColumn = "StartDate"
	// BuildStartDateColumnEpochColumn is the epoch column property of the StartDate name
	BuildStartDateColumnEpochColumn = "StartDate.Epoch"
	// BuildStartDateColumnOffsetColumn is the offset column property of the StartDate name
	BuildStartDateColumnOffsetColumn = "StartDate.Offset"
	// BuildStartDateColumnRfc3339Column is the rfc3339 column property of the StartDate name
	BuildStartDateColumnRfc3339Column = "StartDate.Rfc3339"
	// BuildStatusColumn is the status column name
	BuildStatusColumn = "Status"
	// BuildUpdatedAtColumn is the updated_ts column name
	BuildUpdatedAtColumn = "UpdatedAt"
	// BuildURLColumn is the url column name
	BuildURLColumn = "URL"
)

// BuildEndDate represents the object structure for end_date
type BuildEndDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toBuildEndDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *BuildEndDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *BuildEndDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toBuildEndDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toBuildEndDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toBuildEndDateObject(o.Rfc3339, false),
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

// UnmarshalJSON unmarshals the enum value
func (v BuildEnvironment) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "PRODUCTION":
		v = 0
	case "DEVELOPMENT":
		v = 1
	case "BETA":
		v = 2
	case "STAGING":
		v = 3
	case "TEST":
		v = 4
	case "OTHER":
		v = 5
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v BuildEnvironment) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("PRODUCTION")
	case 1:
		return json.Marshal("DEVELOPMENT")
	case 2:
		return json.Marshal("BETA")
	case 3:
		return json.Marshal("STAGING")
	case 4:
		return json.Marshal("TEST")
	case 5:
		return json.Marshal("OTHER")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

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
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toBuildStartDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *BuildStartDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *BuildStartDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toBuildStartDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toBuildStartDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toBuildStartDateObject(o.Rfc3339, false),
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

// UnmarshalJSON unmarshals the enum value
func (v BuildStatus) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "PASS":
		v = 0
	case "FAIL":
		v = 1
	case "CANCEL":
		v = 2
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v BuildStatus) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("PASS")
	case 1:
		return json.Marshal("FAIL")
	case 2:
		return json.Marshal("CANCEL")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

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
	Automated bool `json:"automated" codec:"automated" bson:"automated" yaml:"automated" faker:"-"`
	// CommitSha the commit sha for the commit that triggered the build
	CommitSha string `json:"commit_sha" codec:"commit_sha" bson:"commit_sha" yaml:"commit_sha" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// EndDate the date when the build finished
	EndDate BuildEndDate `json:"end_date" codec:"end_date" bson:"end_date" yaml:"end_date" faker:"-"`
	// Environment the environment for the build
	Environment BuildEnvironment `json:"environment" codec:"environment" bson:"environment" yaml:"environment" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RepoName the name of the repo
	RepoName string `json:"repo_name" codec:"repo_name" bson:"repo_name" yaml:"repo_name" faker:"-"`
	// StartDate the date when the build started
	StartDate BuildStartDate `json:"start_date" codec:"start_date" bson:"start_date" yaml:"start_date" faker:"-"`
	// Status the status of the build
	Status BuildStatus `json:"status" codec:"status" bson:"status" yaml:"status" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// URL the url to the build status page
	URL *string `json:"url,omitempty" codec:"url,omitempty" bson:"url" yaml:"url,omitempty" faker:"url"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Build)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*Build)(nil)

func toBuildObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *Build:
		return v.ToMap()

	case BuildEndDate:
		return v.ToMap()

	case BuildEnvironment:
		return v.String()

	case BuildStartDate:
		return v.ToMap()

	case BuildStatus:
		return v.String()

	default:
		return o
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

// GetStreamName returns the name of the stream
func (o *Build) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *Build) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *Build) GetModelName() datamodel.ModelNameType {
	return BuildModelName
}

// NewBuildID provides a template for generating an ID field for Build
func NewBuildID(customerID string, refType string, refID string) string {
	return hash.Values("Build", customerID, refType, refID)
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

// IsMutable returns true if the model is mutable
func (o *Build) IsMutable() bool {
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
		CleanupPolicy:     datamodel.CleanupPolicy("compact"),
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
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
func (o *Build) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"automated":   toBuildObject(o.Automated, false),
		"commit_sha":  toBuildObject(o.CommitSha, false),
		"customer_id": toBuildObject(o.CustomerID, false),
		"end_date":    toBuildObject(o.EndDate, false),

		"environment": o.Environment.String(),
		"id":          toBuildObject(o.ID, false),
		"ref_id":      toBuildObject(o.RefID, false),
		"ref_type":    toBuildObject(o.RefType, false),
		"repo_name":   toBuildObject(o.RepoName, false),
		"start_date":  toBuildObject(o.StartDate, false),

		"status":     o.Status.String(),
		"updated_ts": toBuildObject(o.UpdatedAt, false),
		"url":        toBuildObject(o.URL, true),
		"hashcode":   toBuildObject(o.Hashcode, false),
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
				// if coming in as map, convert it back
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
