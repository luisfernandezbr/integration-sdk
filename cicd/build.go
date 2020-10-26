// DO NOT EDIT -- generated code

// Package cicd - the system which contains CI/CD
package cicd

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/bxcodec/faker"
	"github.com/pinpt/go-common/v10/datamodel"
	"github.com/pinpt/go-common/v10/datetime"
	"github.com/pinpt/go-common/v10/hash"
	pjson "github.com/pinpt/go-common/v10/json"
	"github.com/pinpt/go-common/v10/number"
	pstrings "github.com/pinpt/go-common/v10/strings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

const (

	// BuildTable is the default table name
	BuildTable datamodel.ModelNameType = "cicd_build"

	// BuildModelName is the model name
	BuildModelName datamodel.ModelNameType = "cicd.Build"
)

const (
	// BuildModelAutomatedColumn is the column json value automated
	BuildModelAutomatedColumn = "automated"
	// BuildModelCustomerIDColumn is the column json value customer_id
	BuildModelCustomerIDColumn = "customer_id"
	// BuildModelEndDateColumn is the column json value end_date
	BuildModelEndDateColumn = "end_date"
	// BuildModelEndDateEpochColumn is the column json value epoch
	BuildModelEndDateEpochColumn = "epoch"
	// BuildModelEndDateOffsetColumn is the column json value offset
	BuildModelEndDateOffsetColumn = "offset"
	// BuildModelEndDateRfc3339Column is the column json value rfc3339
	BuildModelEndDateRfc3339Column = "rfc3339"
	// BuildModelEnvironmentColumn is the column json value environment
	BuildModelEnvironmentColumn = "environment"
	// BuildModelIDColumn is the column json value id
	BuildModelIDColumn = "id"
	// BuildModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	BuildModelIntegrationInstanceIDColumn = "integration_instance_id"
	// BuildModelMessageColumn is the column json value message
	BuildModelMessageColumn = "message"
	// BuildModelPullrequestURLColumn is the column json value pullrequest_url
	BuildModelPullrequestURLColumn = "pullrequest_url"
	// BuildModelRefIDColumn is the column json value ref_id
	BuildModelRefIDColumn = "ref_id"
	// BuildModelRefTypeColumn is the column json value ref_type
	BuildModelRefTypeColumn = "ref_type"
	// BuildModelShaColumn is the column json value sha
	BuildModelShaColumn = "sha"
	// BuildModelStartDateColumn is the column json value start_date
	BuildModelStartDateColumn = "start_date"
	// BuildModelStartDateEpochColumn is the column json value epoch
	BuildModelStartDateEpochColumn = "epoch"
	// BuildModelStartDateOffsetColumn is the column json value offset
	BuildModelStartDateOffsetColumn = "offset"
	// BuildModelStartDateRfc3339Column is the column json value rfc3339
	BuildModelStartDateRfc3339Column = "rfc3339"
	// BuildModelStatusColumn is the column json value status
	BuildModelStatusColumn = "status"
	// BuildModelURLColumn is the column json value url
	BuildModelURLColumn = "url"
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

// ToMap returns the object as a map
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
				o.Epoch = 0
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
				o.Offset = 0
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// BuildEnvironment is the enumeration type for environment
type BuildEnvironment int32

// toBuildEnvironmentPointer is the enumeration pointer type for environment
func toBuildEnvironmentPointer(v int32) *BuildEnvironment {
	nv := BuildEnvironment(v)
	return &nv
}

// toBuildEnvironmentEnum is the enumeration pointer wrapper for environment
func toBuildEnvironmentEnum(v *BuildEnvironment) string {
	if v == nil {
		return toBuildEnvironmentPointer(0).String()
	}
	return v.String()
}

// UnmarshalBSONValue for unmarshaling value
func (v *BuildEnvironment) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = BuildEnvironment(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "PRODUCTION":
			*v = BuildEnvironment(0)
		case "DEVELOPMENT":
			*v = BuildEnvironment(1)
		case "BETA":
			*v = BuildEnvironment(2)
		case "STAGING":
			*v = BuildEnvironment(3)
		case "TEST":
			*v = BuildEnvironment(4)
		case "OTHER":
			*v = BuildEnvironment(5)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v *BuildEnvironment) UnmarshalJSON(buf []byte) error {
	var val string
	if err := json.Unmarshal(buf, &val); err != nil {
		return err
	}
	switch val {
	case "PRODUCTION":
		*v = 0
	case "DEVELOPMENT":
		*v = 1
	case "BETA":
		*v = 2
	case "STAGING":
		*v = 3
	case "TEST":
		*v = 4
	case "OTHER":
		*v = 5
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

// FromInterface for decoding from an interface
func (v *BuildEnvironment) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case BuildEnvironment:
		*v = val
	case int32:
		*v = BuildEnvironment(int32(val))
	case int:
		*v = BuildEnvironment(int32(val))
	case string:
		switch val {
		case "PRODUCTION":
			*v = BuildEnvironment(0)
		case "DEVELOPMENT":
			*v = BuildEnvironment(1)
		case "BETA":
			*v = BuildEnvironment(2)
		case "STAGING":
			*v = BuildEnvironment(3)
		case "TEST":
			*v = BuildEnvironment(4)
		case "OTHER":
			*v = BuildEnvironment(5)
		}
	}
	return nil
}

const (
	// BuildEnvironmentProduction is the enumeration value for production
	BuildEnvironmentProduction BuildEnvironment = 0
	// BuildEnvironmentDevelopment is the enumeration value for development
	BuildEnvironmentDevelopment BuildEnvironment = 1
	// BuildEnvironmentBeta is the enumeration value for beta
	BuildEnvironmentBeta BuildEnvironment = 2
	// BuildEnvironmentStaging is the enumeration value for staging
	BuildEnvironmentStaging BuildEnvironment = 3
	// BuildEnvironmentTest is the enumeration value for test
	BuildEnvironmentTest BuildEnvironment = 4
	// BuildEnvironmentOther is the enumeration value for other
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

// ToMap returns the object as a map
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
				o.Epoch = 0
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
				o.Offset = 0
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// BuildStatus is the enumeration type for status
type BuildStatus int32

// toBuildStatusPointer is the enumeration pointer type for status
func toBuildStatusPointer(v int32) *BuildStatus {
	nv := BuildStatus(v)
	return &nv
}

// toBuildStatusEnum is the enumeration pointer wrapper for status
func toBuildStatusEnum(v *BuildStatus) string {
	if v == nil {
		return toBuildStatusPointer(0).String()
	}
	return v.String()
}

// UnmarshalBSONValue for unmarshaling value
func (v *BuildStatus) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = BuildStatus(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "RUNNING":
			*v = BuildStatus(0)
		case "PASS":
			*v = BuildStatus(1)
		case "FAIL":
			*v = BuildStatus(2)
		case "CANCEL":
			*v = BuildStatus(3)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v *BuildStatus) UnmarshalJSON(buf []byte) error {
	var val string
	if err := json.Unmarshal(buf, &val); err != nil {
		return err
	}
	switch val {
	case "RUNNING":
		*v = 0
	case "PASS":
		*v = 1
	case "FAIL":
		*v = 2
	case "CANCEL":
		*v = 3
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v BuildStatus) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("RUNNING")
	case 1:
		return json.Marshal("PASS")
	case 2:
		return json.Marshal("FAIL")
	case 3:
		return json.Marshal("CANCEL")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for Status
func (v BuildStatus) String() string {
	switch int32(v) {
	case 0:
		return "RUNNING"
	case 1:
		return "PASS"
	case 2:
		return "FAIL"
	case 3:
		return "CANCEL"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *BuildStatus) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case BuildStatus:
		*v = val
	case int32:
		*v = BuildStatus(int32(val))
	case int:
		*v = BuildStatus(int32(val))
	case string:
		switch val {
		case "RUNNING":
			*v = BuildStatus(0)
		case "PASS":
			*v = BuildStatus(1)
		case "FAIL":
			*v = BuildStatus(2)
		case "CANCEL":
			*v = BuildStatus(3)
		}
	}
	return nil
}

const (
	// BuildStatusRunning is the enumeration value for running
	BuildStatusRunning BuildStatus = 0
	// BuildStatusPass is the enumeration value for pass
	BuildStatusPass BuildStatus = 1
	// BuildStatusFail is the enumeration value for fail
	BuildStatusFail BuildStatus = 2
	// BuildStatusCancel is the enumeration value for cancel
	BuildStatusCancel BuildStatus = 3
)

// Build an individual build in a continous integration system
type Build struct {
	// Automated if the build was automated or manual
	Automated bool `json:"automated" codec:"automated" bson:"automated" yaml:"automated" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// EndDate the date when the build finished
	EndDate BuildEndDate `json:"end_date" codec:"end_date" bson:"end_date" yaml:"end_date" faker:"-"`
	// Environment the environment for the build
	Environment BuildEnvironment `json:"environment" codec:"environment" bson:"environment" yaml:"environment" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// Message an optional message about the build
	Message *string `json:"message,omitempty" codec:"message,omitempty" bson:"message" yaml:"message,omitempty" faker:"-"`
	// PullrequestURL the pull request url that triggered this build
	PullrequestURL *string `json:"pullrequest_url,omitempty" codec:"pullrequest_url,omitempty" bson:"pullrequest_url" yaml:"pullrequest_url,omitempty" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// Sha the commit sha for the commit that triggered the build
	Sha string `json:"sha" codec:"sha" bson:"sha" yaml:"sha" faker:"-"`
	// StartDate the date when the build started
	StartDate BuildStartDate `json:"start_date" codec:"start_date" bson:"start_date" yaml:"start_date" faker:"-"`
	// Status the status of the build
	Status BuildStatus `json:"status" codec:"status" bson:"status" yaml:"status" faker:"-"`
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
	return ""
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
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Build) GetTimestamp() time.Time {
	return time.Now().UTC()
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
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *Build) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = BuildModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Build) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *Build) GetCustomerID() string {
	return o.CustomerID
}

// SetCustomerID will return the customer_id
func (o *Build) SetCustomerID(id string) {
	o.CustomerID = id
}

// GetRefType will return the ref_type
func (o *Build) GetRefType() string {
	return o.RefType
}

// SetRefType will return the ref_type
func (o *Build) SetRefType(t string) {
	o.RefType = t
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

// StringifyPretty returns the object in JSON format as a string prettified
func (o *Build) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
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
		"customer_id": toBuildObject(o.CustomerID, false),
		"end_date":    toBuildObject(o.EndDate, false),

		"environment":             o.Environment.String(),
		"id":                      toBuildObject(o.ID, false),
		"integration_instance_id": toBuildObject(o.IntegrationInstanceID, true),
		"message":                 toBuildObject(o.Message, true),
		"pullrequest_url":         toBuildObject(o.PullrequestURL, true),
		"ref_id":                  toBuildObject(o.RefID, false),
		"ref_type":                toBuildObject(o.RefType, false),
		"sha":                     toBuildObject(o.Sha, false),
		"start_date":              toBuildObject(o.StartDate, false),

		"status":   o.Status.String(),
		"url":      toBuildObject(o.URL, true),
		"hashcode": toBuildObject(o.Hashcode, false),
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
				o.Automated = false
			} else {
				o.Automated = number.ToBoolAny(val)
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
			dt := datetime.NewDateWithTime(tv)
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.ID = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["integration_instance_id"].(*string); ok {
		o.IntegrationInstanceID = val
	} else if val, ok := kv["integration_instance_id"].(string); ok {
		o.IntegrationInstanceID = &val
	} else {
		if val, ok := kv["integration_instance_id"]; ok {
			if val == nil {
				o.IntegrationInstanceID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.IntegrationInstanceID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["message"].(*string); ok {
		o.Message = val
	} else if val, ok := kv["message"].(string); ok {
		o.Message = &val
	} else {
		if val, ok := kv["message"]; ok {
			if val == nil {
				o.Message = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Message = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["pullrequest_url"].(*string); ok {
		o.PullrequestURL = val
	} else if val, ok := kv["pullrequest_url"].(string); ok {
		o.PullrequestURL = &val
	} else {
		if val, ok := kv["pullrequest_url"]; ok {
			if val == nil {
				o.PullrequestURL = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.PullrequestURL = pstrings.Pointer(fmt.Sprintf("%v", val))
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.RefType = fmt.Sprintf("%v", val)
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Sha = fmt.Sprintf("%v", val)
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
			dt := datetime.NewDateWithTime(tv)
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
			case "running", "RUNNING":
				o.Status = 0
			case "pass", "PASS":
				o.Status = 1
			case "fail", "FAIL":
				o.Status = 2
			case "cancel", "CANCEL":
				o.Status = 3
			}
		}
		if em, ok := kv["status"].(string); ok {
			switch em {
			case "running", "RUNNING":
				o.Status = 0
			case "pass", "PASS":
				o.Status = 1
			case "fail", "FAIL":
				o.Status = 2
			case "cancel", "CANCEL":
				o.Status = 3
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
	args = append(args, o.CustomerID)
	args = append(args, o.EndDate)
	args = append(args, o.Environment)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.Message)
	args = append(args, o.PullrequestURL)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.Sha)
	args = append(args, o.StartDate)
	args = append(args, o.Status)
	args = append(args, o.URL)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// SetIntegrationInstanceID will set the integration instance ID
func (o *Build) SetIntegrationInstanceID(id string) {
	if id == "" {
		o.IntegrationInstanceID = nil
	} else {
		o.IntegrationInstanceID = &id
	}
}

// GetIntegrationInstanceID will return the integration instance ID
func (o *Build) GetIntegrationInstanceID() *string {
	return o.IntegrationInstanceID
}

// BuildPartial is a partial struct for upsert mutations for Build
type BuildPartial struct {
	// Automated if the build was automated or manual
	Automated *bool `json:"automated,omitempty"`
	// EndDate the date when the build finished
	EndDate *BuildEndDate `json:"end_date,omitempty"`
	// Environment the environment for the build
	Environment *BuildEnvironment `json:"environment,omitempty"`
	// Message an optional message about the build
	Message *string `json:"message,omitempty"`
	// PullrequestURL the pull request url that triggered this build
	PullrequestURL *string `json:"pullrequest_url,omitempty"`
	// Sha the commit sha for the commit that triggered the build
	Sha *string `json:"sha,omitempty"`
	// StartDate the date when the build started
	StartDate *BuildStartDate `json:"start_date,omitempty"`
	// Status the status of the build
	Status *BuildStatus `json:"status,omitempty"`
	// URL the url to the build status page
	URL *string `json:"url,omitempty"`
}

var _ datamodel.PartialModel = (*BuildPartial)(nil)

// GetModelName returns the name of the model
func (o *BuildPartial) GetModelName() datamodel.ModelNameType {
	return BuildModelName
}

// ToMap returns the object as a map
func (o *BuildPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"automated": toBuildObject(o.Automated, true),
		"end_date":  toBuildObject(o.EndDate, true),

		"environment":     toBuildEnvironmentEnum(o.Environment),
		"message":         toBuildObject(o.Message, true),
		"pullrequest_url": toBuildObject(o.PullrequestURL, true),
		"sha":             toBuildObject(o.Sha, true),
		"start_date":      toBuildObject(o.StartDate, true),

		"status": toBuildStatusEnum(o.Status),
		"url":    toBuildObject(o.URL, true),
	}
	for k, v := range kv {
		if v == nil || reflect.ValueOf(v).IsZero() {
			delete(kv, k)
		} else {
			if k == "end_date" {
				if dt, ok := v.(*BuildEndDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}
			if k == "start_date" {
				if dt, ok := v.(*BuildStartDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}
		}
	}
	return kv
}

// Stringify returns the object in JSON format as a string
func (o *BuildPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *BuildPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *BuildPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *BuildPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *BuildPartial) FromMap(kv map[string]interface{}) {
	if val, ok := kv["automated"].(*bool); ok {
		o.Automated = val
	} else if val, ok := kv["automated"].(bool); ok {
		o.Automated = &val
	} else {
		if val, ok := kv["automated"]; ok {
			if val == nil {
				o.Automated = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Automated = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}

	if o.EndDate == nil {
		o.EndDate = &BuildEndDate{}
	}

	if val, ok := kv["end_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.EndDate.FromMap(kv)
		} else if sv, ok := val.(BuildEndDate); ok {
			// struct
			o.EndDate = &sv
		} else if sp, ok := val.(*BuildEndDate); ok {
			// struct pointer
			o.EndDate = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.EndDate.Epoch = dt.Epoch
			o.EndDate.Rfc3339 = dt.Rfc3339
			o.EndDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
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

	if val, ok := kv["environment"].(*BuildEnvironment); ok {
		o.Environment = val
	} else if val, ok := kv["environment"].(BuildEnvironment); ok {
		o.Environment = &val
	} else {
		if val, ok := kv["environment"]; ok {
			if val == nil {
				o.Environment = toBuildEnvironmentPointer(0)
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["BuildEnvironment"]
				}
				// this is an enum pointer
				if em, ok := val.(string); ok {
					switch em {
					case "production", "PRODUCTION":
						o.Environment = toBuildEnvironmentPointer(0)
					case "development", "DEVELOPMENT":
						o.Environment = toBuildEnvironmentPointer(1)
					case "beta", "BETA":
						o.Environment = toBuildEnvironmentPointer(2)
					case "staging", "STAGING":
						o.Environment = toBuildEnvironmentPointer(3)
					case "test", "TEST":
						o.Environment = toBuildEnvironmentPointer(4)
					case "other", "OTHER":
						o.Environment = toBuildEnvironmentPointer(5)
					}
				}
			}
		}
	}
	if val, ok := kv["message"].(*string); ok {
		o.Message = val
	} else if val, ok := kv["message"].(string); ok {
		o.Message = &val
	} else {
		if val, ok := kv["message"]; ok {
			if val == nil {
				o.Message = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Message = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["pullrequest_url"].(*string); ok {
		o.PullrequestURL = val
	} else if val, ok := kv["pullrequest_url"].(string); ok {
		o.PullrequestURL = &val
	} else {
		if val, ok := kv["pullrequest_url"]; ok {
			if val == nil {
				o.PullrequestURL = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.PullrequestURL = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["sha"].(*string); ok {
		o.Sha = val
	} else if val, ok := kv["sha"].(string); ok {
		o.Sha = &val
	} else {
		if val, ok := kv["sha"]; ok {
			if val == nil {
				o.Sha = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Sha = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if o.StartDate == nil {
		o.StartDate = &BuildStartDate{}
	}

	if val, ok := kv["start_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.StartDate.FromMap(kv)
		} else if sv, ok := val.(BuildStartDate); ok {
			// struct
			o.StartDate = &sv
		} else if sp, ok := val.(*BuildStartDate); ok {
			// struct pointer
			o.StartDate = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.StartDate.Epoch = dt.Epoch
			o.StartDate.Rfc3339 = dt.Rfc3339
			o.StartDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
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

	if val, ok := kv["status"].(*BuildStatus); ok {
		o.Status = val
	} else if val, ok := kv["status"].(BuildStatus); ok {
		o.Status = &val
	} else {
		if val, ok := kv["status"]; ok {
			if val == nil {
				o.Status = toBuildStatusPointer(0)
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["BuildStatus"]
				}
				// this is an enum pointer
				if em, ok := val.(string); ok {
					switch em {
					case "running", "RUNNING":
						o.Status = toBuildStatusPointer(0)
					case "pass", "PASS":
						o.Status = toBuildStatusPointer(1)
					case "fail", "FAIL":
						o.Status = toBuildStatusPointer(2)
					case "cancel", "CANCEL":
						o.Status = toBuildStatusPointer(3)
					}
				}
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
