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
	// DeploymentTopic is the default topic name
	DeploymentTopic datamodel.TopicNameType = "cicd_Deployment_topic"

	// DeploymentModelName is the model name
	DeploymentModelName datamodel.ModelNameType = "cicd.Deployment"
)

const (
	// DeploymentAutomatedColumn is the automated column name
	DeploymentAutomatedColumn = "Automated"
	// DeploymentBuildRefIDColumn is the build_ref_id column name
	DeploymentBuildRefIDColumn = "BuildRefID"
	// DeploymentCommitShaColumn is the commit_sha column name
	DeploymentCommitShaColumn = "CommitSha"
	// DeploymentCustomerIDColumn is the customer_id column name
	DeploymentCustomerIDColumn = "CustomerID"
	// DeploymentEndDateColumn is the end_date column name
	DeploymentEndDateColumn = "EndDate"
	// DeploymentEndDateColumnEpochColumn is the epoch column property of the EndDate name
	DeploymentEndDateColumnEpochColumn = "EndDate.Epoch"
	// DeploymentEndDateColumnOffsetColumn is the offset column property of the EndDate name
	DeploymentEndDateColumnOffsetColumn = "EndDate.Offset"
	// DeploymentEndDateColumnRfc3339Column is the rfc3339 column property of the EndDate name
	DeploymentEndDateColumnRfc3339Column = "EndDate.Rfc3339"
	// DeploymentEnvironmentColumn is the environment column name
	DeploymentEnvironmentColumn = "Environment"
	// DeploymentIDColumn is the id column name
	DeploymentIDColumn = "ID"
	// DeploymentRefIDColumn is the ref_id column name
	DeploymentRefIDColumn = "RefID"
	// DeploymentRefTypeColumn is the ref_type column name
	DeploymentRefTypeColumn = "RefType"
	// DeploymentRepoNameColumn is the repo_name column name
	DeploymentRepoNameColumn = "RepoName"
	// DeploymentStartDateColumn is the start_date column name
	DeploymentStartDateColumn = "StartDate"
	// DeploymentStartDateColumnEpochColumn is the epoch column property of the StartDate name
	DeploymentStartDateColumnEpochColumn = "StartDate.Epoch"
	// DeploymentStartDateColumnOffsetColumn is the offset column property of the StartDate name
	DeploymentStartDateColumnOffsetColumn = "StartDate.Offset"
	// DeploymentStartDateColumnRfc3339Column is the rfc3339 column property of the StartDate name
	DeploymentStartDateColumnRfc3339Column = "StartDate.Rfc3339"
	// DeploymentStatusColumn is the status column name
	DeploymentStatusColumn = "Status"
	// DeploymentUpdatedAtColumn is the updated_ts column name
	DeploymentUpdatedAtColumn = "UpdatedAt"
	// DeploymentURLColumn is the url column name
	DeploymentURLColumn = "URL"
)

// DeploymentEndDate represents the object structure for end_date
type DeploymentEndDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toDeploymentEndDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *DeploymentEndDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *DeploymentEndDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toDeploymentEndDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toDeploymentEndDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toDeploymentEndDateObject(o.Rfc3339, false),
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
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toDeploymentStartDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *DeploymentStartDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *DeploymentStartDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toDeploymentStartDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toDeploymentStartDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toDeploymentStartDateObject(o.Rfc3339, false),
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
	Automated bool `json:"automated" codec:"automated" bson:"automated" yaml:"automated" faker:"-"`
	// BuildRefID the build id that is associated with the deployment (if any)
	BuildRefID string `json:"build_ref_id" codec:"build_ref_id" bson:"build_ref_id" yaml:"build_ref_id" faker:"-"`
	// CommitSha the commit sha for the commit that triggered the deployment
	CommitSha string `json:"commit_sha" codec:"commit_sha" bson:"commit_sha" yaml:"commit_sha" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// EndDate the date when the deployment finished
	EndDate DeploymentEndDate `json:"end_date" codec:"end_date" bson:"end_date" yaml:"end_date" faker:"-"`
	// Environment the environment for the deployment
	Environment DeploymentEnvironment `json:"environment" codec:"environment" bson:"environment" yaml:"environment" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RepoName the name of the repo
	RepoName string `json:"repo_name" codec:"repo_name" bson:"repo_name" yaml:"repo_name" faker:"-"`
	// StartDate the date when the deployment started
	StartDate DeploymentStartDate `json:"start_date" codec:"start_date" bson:"start_date" yaml:"start_date" faker:"-"`
	// Status the status of the deployment
	Status DeploymentStatus `json:"status" codec:"status" bson:"status" yaml:"status" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// URL the url to the deployment status page
	URL *string `json:"url,omitempty" codec:"url,omitempty" bson:"url" yaml:"url,omitempty" faker:"url"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Deployment)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*Deployment)(nil)

func toDeploymentObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *Deployment:
		return v.ToMap()

	case DeploymentEndDate:
		return v.ToMap()

	case DeploymentEnvironment:
		return v.String()

	case DeploymentStartDate:
		return v.ToMap()

	case DeploymentStatus:
		return v.String()

	default:
		return o
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

// GetStreamName returns the name of the stream
func (o *Deployment) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *Deployment) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *Deployment) GetModelName() datamodel.ModelNameType {
	return DeploymentModelName
}

// NewDeploymentID provides a template for generating an ID field for Deployment
func NewDeploymentID(customerID string, refType string, refID string) string {
	return hash.Values("Deployment", customerID, refType, refID)
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
	var i interface{} = o.CustomerID
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

// IsMutable returns true if the model is mutable
func (o *Deployment) IsMutable() bool {
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
func (o *Deployment) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"automated":    toDeploymentObject(o.Automated, false),
		"build_ref_id": toDeploymentObject(o.BuildRefID, false),
		"commit_sha":   toDeploymentObject(o.CommitSha, false),
		"customer_id":  toDeploymentObject(o.CustomerID, false),
		"end_date":     toDeploymentObject(o.EndDate, false),
		"environment":  toDeploymentObject(o.Environment, false),
		"id":           toDeploymentObject(o.ID, false),
		"ref_id":       toDeploymentObject(o.RefID, false),
		"ref_type":     toDeploymentObject(o.RefType, false),
		"repo_name":    toDeploymentObject(o.RepoName, false),
		"start_date":   toDeploymentObject(o.StartDate, false),
		"status":       toDeploymentObject(o.Status, false),
		"updated_ts":   toDeploymentObject(o.UpdatedAt, false),
		"url":          toDeploymentObject(o.URL, true),
		"hashcode":     toDeploymentObject(o.Hashcode, false),
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
