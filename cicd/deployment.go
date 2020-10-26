// DO NOT EDIT -- generated code

// Package cicd - the system which contains CI/CD
package cicd

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/bxcodec/faker"
	"github.com/pinpt/go-common/v10/datamodel"
	"github.com/pinpt/go-common/v10/datetime"
	"github.com/pinpt/go-common/v10/hash"
	pjson "github.com/pinpt/go-common/v10/json"
	"github.com/pinpt/go-common/v10/number"
	"github.com/pinpt/go-common/v10/slice"
	pstrings "github.com/pinpt/go-common/v10/strings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (

	// DeploymentTable is the default table name
	DeploymentTable datamodel.ModelNameType = "cicd_deployment"

	// DeploymentModelName is the model name
	DeploymentModelName datamodel.ModelNameType = "cicd.Deployment"
)

const (
	// DeploymentModelAutomatedColumn is the column json value automated
	DeploymentModelAutomatedColumn = "automated"
	// DeploymentModelBuildRefIdsColumn is the column json value build_ref_ids
	DeploymentModelBuildRefIdsColumn = "build_ref_ids"
	// DeploymentModelCustomerIDColumn is the column json value customer_id
	DeploymentModelCustomerIDColumn = "customer_id"
	// DeploymentModelEndDateColumn is the column json value end_date
	DeploymentModelEndDateColumn = "end_date"
	// DeploymentModelEndDateEpochColumn is the column json value epoch
	DeploymentModelEndDateEpochColumn = "epoch"
	// DeploymentModelEndDateOffsetColumn is the column json value offset
	DeploymentModelEndDateOffsetColumn = "offset"
	// DeploymentModelEndDateRfc3339Column is the column json value rfc3339
	DeploymentModelEndDateRfc3339Column = "rfc3339"
	// DeploymentModelEnvironmentColumn is the column json value environment
	DeploymentModelEnvironmentColumn = "environment"
	// DeploymentModelIDColumn is the column json value id
	DeploymentModelIDColumn = "id"
	// DeploymentModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	DeploymentModelIntegrationInstanceIDColumn = "integration_instance_id"
	// DeploymentModelMessageColumn is the column json value message
	DeploymentModelMessageColumn = "message"
	// DeploymentModelRefIDColumn is the column json value ref_id
	DeploymentModelRefIDColumn = "ref_id"
	// DeploymentModelRefTypeColumn is the column json value ref_type
	DeploymentModelRefTypeColumn = "ref_type"
	// DeploymentModelShasColumn is the column json value shas
	DeploymentModelShasColumn = "shas"
	// DeploymentModelStartDateColumn is the column json value start_date
	DeploymentModelStartDateColumn = "start_date"
	// DeploymentModelStartDateEpochColumn is the column json value epoch
	DeploymentModelStartDateEpochColumn = "epoch"
	// DeploymentModelStartDateOffsetColumn is the column json value offset
	DeploymentModelStartDateOffsetColumn = "offset"
	// DeploymentModelStartDateRfc3339Column is the column json value rfc3339
	DeploymentModelStartDateRfc3339Column = "rfc3339"
	// DeploymentModelStatusColumn is the column json value status
	DeploymentModelStatusColumn = "status"
	// DeploymentModelURLColumn is the column json value url
	DeploymentModelURLColumn = "url"
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

// ToMap returns the object as a map
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

// DeploymentEnvironment is the enumeration type for environment
type DeploymentEnvironment int32

// toDeploymentEnvironmentPointer is the enumeration pointer type for environment
func toDeploymentEnvironmentPointer(v int32) *DeploymentEnvironment {
	nv := DeploymentEnvironment(v)
	return &nv
}

// toDeploymentEnvironmentEnum is the enumeration pointer wrapper for environment
func toDeploymentEnvironmentEnum(v *DeploymentEnvironment) string {
	if v == nil {
		return toDeploymentEnvironmentPointer(0).String()
	}
	return v.String()
}

// UnmarshalBSONValue for unmarshaling value
func (v *DeploymentEnvironment) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = DeploymentEnvironment(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "PRODUCTION":
			*v = DeploymentEnvironment(0)
		case "DEVELOPMENT":
			*v = DeploymentEnvironment(1)
		case "BETA":
			*v = DeploymentEnvironment(2)
		case "STAGING":
			*v = DeploymentEnvironment(3)
		case "TEST":
			*v = DeploymentEnvironment(4)
		case "OTHER":
			*v = DeploymentEnvironment(5)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v *DeploymentEnvironment) UnmarshalJSON(buf []byte) error {
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
func (v DeploymentEnvironment) MarshalJSON() ([]byte, error) {
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

// FromInterface for decoding from an interface
func (v *DeploymentEnvironment) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case DeploymentEnvironment:
		*v = val
	case int32:
		*v = DeploymentEnvironment(int32(val))
	case int:
		*v = DeploymentEnvironment(int32(val))
	case string:
		switch val {
		case "PRODUCTION":
			*v = DeploymentEnvironment(0)
		case "DEVELOPMENT":
			*v = DeploymentEnvironment(1)
		case "BETA":
			*v = DeploymentEnvironment(2)
		case "STAGING":
			*v = DeploymentEnvironment(3)
		case "TEST":
			*v = DeploymentEnvironment(4)
		case "OTHER":
			*v = DeploymentEnvironment(5)
		}
	}
	return nil
}

const (
	// DeploymentEnvironmentProduction is the enumeration value for production
	DeploymentEnvironmentProduction DeploymentEnvironment = 0
	// DeploymentEnvironmentDevelopment is the enumeration value for development
	DeploymentEnvironmentDevelopment DeploymentEnvironment = 1
	// DeploymentEnvironmentBeta is the enumeration value for beta
	DeploymentEnvironmentBeta DeploymentEnvironment = 2
	// DeploymentEnvironmentStaging is the enumeration value for staging
	DeploymentEnvironmentStaging DeploymentEnvironment = 3
	// DeploymentEnvironmentTest is the enumeration value for test
	DeploymentEnvironmentTest DeploymentEnvironment = 4
	// DeploymentEnvironmentOther is the enumeration value for other
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

// ToMap returns the object as a map
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

// DeploymentStatus is the enumeration type for status
type DeploymentStatus int32

// toDeploymentStatusPointer is the enumeration pointer type for status
func toDeploymentStatusPointer(v int32) *DeploymentStatus {
	nv := DeploymentStatus(v)
	return &nv
}

// toDeploymentStatusEnum is the enumeration pointer wrapper for status
func toDeploymentStatusEnum(v *DeploymentStatus) string {
	if v == nil {
		return toDeploymentStatusPointer(0).String()
	}
	return v.String()
}

// UnmarshalBSONValue for unmarshaling value
func (v *DeploymentStatus) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = DeploymentStatus(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "RUNNING":
			*v = DeploymentStatus(0)
		case "PASS":
			*v = DeploymentStatus(1)
		case "FAIL":
			*v = DeploymentStatus(2)
		case "CANCEL":
			*v = DeploymentStatus(3)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v *DeploymentStatus) UnmarshalJSON(buf []byte) error {
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
func (v DeploymentStatus) MarshalJSON() ([]byte, error) {
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
func (v DeploymentStatus) String() string {
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
func (v *DeploymentStatus) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case DeploymentStatus:
		*v = val
	case int32:
		*v = DeploymentStatus(int32(val))
	case int:
		*v = DeploymentStatus(int32(val))
	case string:
		switch val {
		case "RUNNING":
			*v = DeploymentStatus(0)
		case "PASS":
			*v = DeploymentStatus(1)
		case "FAIL":
			*v = DeploymentStatus(2)
		case "CANCEL":
			*v = DeploymentStatus(3)
		}
	}
	return nil
}

const (
	// DeploymentStatusRunning is the enumeration value for running
	DeploymentStatusRunning DeploymentStatus = 0
	// DeploymentStatusPass is the enumeration value for pass
	DeploymentStatusPass DeploymentStatus = 1
	// DeploymentStatusFail is the enumeration value for fail
	DeploymentStatusFail DeploymentStatus = 2
	// DeploymentStatusCancel is the enumeration value for cancel
	DeploymentStatusCancel DeploymentStatus = 3
)

// Deployment an individual deployment in a continous deployment system
type Deployment struct {
	// Automated if the deployment was automated or manual
	Automated bool `json:"automated" codec:"automated" bson:"automated" yaml:"automated" faker:"-"`
	// BuildRefIds the build ids that is associated with the deployment (if any)
	BuildRefIds []string `json:"build_ref_ids" codec:"build_ref_ids" bson:"build_ref_ids" yaml:"build_ref_ids" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// EndDate the date when the deployment finished
	EndDate DeploymentEndDate `json:"end_date" codec:"end_date" bson:"end_date" yaml:"end_date" faker:"-"`
	// Environment the environment for the deployment
	Environment DeploymentEnvironment `json:"environment" codec:"environment" bson:"environment" yaml:"environment" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// Message an optional message about the build
	Message *string `json:"message,omitempty" codec:"message,omitempty" bson:"message" yaml:"message,omitempty" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// Shas the commit shas for the commit that triggered the deployment
	Shas []string `json:"shas" codec:"shas" bson:"shas" yaml:"shas" faker:"-"`
	// StartDate the date when the deployment started
	StartDate DeploymentStartDate `json:"start_date" codec:"start_date" bson:"start_date" yaml:"start_date" faker:"-"`
	// Status the status of the deployment
	Status DeploymentStatus `json:"status" codec:"status" bson:"status" yaml:"status" faker:"-"`
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
	return ""
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
	if o.BuildRefIds == nil {
		o.BuildRefIds = make([]string, 0)
	}
	if o.Shas == nil {
		o.Shas = make([]string, 0)
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
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Deployment) GetTimestamp() time.Time {
	return time.Now().UTC()
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
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *Deployment) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = DeploymentModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Deployment) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *Deployment) GetCustomerID() string {
	return o.CustomerID
}

// SetCustomerID will return the customer_id
func (o *Deployment) SetCustomerID(id string) {
	o.CustomerID = id
}

// GetRefType will return the ref_type
func (o *Deployment) GetRefType() string {
	return o.RefType
}

// SetRefType will return the ref_type
func (o *Deployment) SetRefType(t string) {
	o.RefType = t
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

// StringifyPretty returns the object in JSON format as a string prettified
func (o *Deployment) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two Deployment objects are equal
func (o *Deployment) IsEqual(other *Deployment) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Deployment) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"automated":     toDeploymentObject(o.Automated, false),
		"build_ref_ids": toDeploymentObject(o.BuildRefIds, false),
		"customer_id":   toDeploymentObject(o.CustomerID, false),
		"end_date":      toDeploymentObject(o.EndDate, false),

		"environment":             o.Environment.String(),
		"id":                      toDeploymentObject(o.ID, false),
		"integration_instance_id": toDeploymentObject(o.IntegrationInstanceID, true),
		"message":                 toDeploymentObject(o.Message, true),
		"ref_id":                  toDeploymentObject(o.RefID, false),
		"ref_type":                toDeploymentObject(o.RefType, false),
		"shas":                    toDeploymentObject(o.Shas, false),
		"start_date":              toDeploymentObject(o.StartDate, false),

		"status":   o.Status.String(),
		"url":      toDeploymentObject(o.URL, true),
		"hashcode": toDeploymentObject(o.Hashcode, false),
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
				o.Automated = false
			} else {
				o.Automated = number.ToBoolAny(val)
			}
		}
	}
	if val, ok := kv["build_ref_ids"]; ok {
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
								panic("unsupported type for build_ref_ids field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for build_ref_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for build_ref_ids field")
				}
			}
			o.BuildRefIds = na
		}
	}
	if o.BuildRefIds == nil {
		o.BuildRefIds = make([]string, 0)
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
	if val, ok := kv["shas"]; ok {
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
								panic("unsupported type for shas field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for shas field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for shas field")
				}
			}
			o.Shas = na
		}
	}
	if o.Shas == nil {
		o.Shas = make([]string, 0)
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

	if val, ok := kv["status"].(DeploymentStatus); ok {
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
func (o *Deployment) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Automated)
	args = append(args, o.BuildRefIds)
	args = append(args, o.CustomerID)
	args = append(args, o.EndDate)
	args = append(args, o.Environment)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.Message)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.Shas)
	args = append(args, o.StartDate)
	args = append(args, o.Status)
	args = append(args, o.URL)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// SetIntegrationInstanceID will set the integration instance ID
func (o *Deployment) SetIntegrationInstanceID(id string) {
	if id == "" {
		o.IntegrationInstanceID = nil
	} else {
		o.IntegrationInstanceID = &id
	}
}

// GetIntegrationInstanceID will return the integration instance ID
func (o *Deployment) GetIntegrationInstanceID() *string {
	return o.IntegrationInstanceID
}

// DeploymentPartial is a partial struct for upsert mutations for Deployment
type DeploymentPartial struct {
	// Automated if the deployment was automated or manual
	Automated *bool `json:"automated,omitempty"`
	// BuildRefIds the build ids that is associated with the deployment (if any)
	BuildRefIds []string `json:"build_ref_ids,omitempty"`
	// EndDate the date when the deployment finished
	EndDate *DeploymentEndDate `json:"end_date,omitempty"`
	// Environment the environment for the deployment
	Environment *DeploymentEnvironment `json:"environment,omitempty"`
	// Message an optional message about the build
	Message *string `json:"message,omitempty"`
	// Shas the commit shas for the commit that triggered the deployment
	Shas []string `json:"shas,omitempty"`
	// StartDate the date when the deployment started
	StartDate *DeploymentStartDate `json:"start_date,omitempty"`
	// Status the status of the deployment
	Status *DeploymentStatus `json:"status,omitempty"`
	// URL the url to the deployment status page
	URL *string `json:"url,omitempty"`
}

var _ datamodel.PartialModel = (*DeploymentPartial)(nil)

// GetModelName returns the name of the model
func (o *DeploymentPartial) GetModelName() datamodel.ModelNameType {
	return DeploymentModelName
}

// ToMap returns the object as a map
func (o *DeploymentPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"automated":     toDeploymentObject(o.Automated, true),
		"build_ref_ids": toDeploymentObject(o.BuildRefIds, true),
		"end_date":      toDeploymentObject(o.EndDate, true),

		"environment": toDeploymentEnvironmentEnum(o.Environment),
		"message":     toDeploymentObject(o.Message, true),
		"shas":        toDeploymentObject(o.Shas, true),
		"start_date":  toDeploymentObject(o.StartDate, true),

		"status": toDeploymentStatusEnum(o.Status),
		"url":    toDeploymentObject(o.URL, true),
	}
	for k, v := range kv {
		if v == nil || reflect.ValueOf(v).IsZero() {
			delete(kv, k)
		} else {

			if k == "build_ref_ids" {
				if arr, ok := v.([]string); ok {
					if len(arr) == 0 {
						delete(kv, k)
					}
				}
			}
			if k == "end_date" {
				if dt, ok := v.(*DeploymentEndDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}

			if k == "shas" {
				if arr, ok := v.([]string); ok {
					if len(arr) == 0 {
						delete(kv, k)
					}
				}
			}
			if k == "start_date" {
				if dt, ok := v.(*DeploymentStartDate); ok {
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
func (o *DeploymentPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *DeploymentPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *DeploymentPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *DeploymentPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *DeploymentPartial) FromMap(kv map[string]interface{}) {
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
	if val, ok := kv["build_ref_ids"]; ok {
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
								panic("unsupported type for build_ref_ids field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for build_ref_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for build_ref_ids field")
				}
			}
			o.BuildRefIds = na
		}
	}
	if o.BuildRefIds == nil {
		o.BuildRefIds = make([]string, 0)
	}

	if o.EndDate == nil {
		o.EndDate = &DeploymentEndDate{}
	}

	if val, ok := kv["end_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.EndDate.FromMap(kv)
		} else if sv, ok := val.(DeploymentEndDate); ok {
			// struct
			o.EndDate = &sv
		} else if sp, ok := val.(*DeploymentEndDate); ok {
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

	if val, ok := kv["environment"].(*DeploymentEnvironment); ok {
		o.Environment = val
	} else if val, ok := kv["environment"].(DeploymentEnvironment); ok {
		o.Environment = &val
	} else {
		if val, ok := kv["environment"]; ok {
			if val == nil {
				o.Environment = toDeploymentEnvironmentPointer(0)
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["DeploymentEnvironment"]
				}
				// this is an enum pointer
				if em, ok := val.(string); ok {
					switch em {
					case "production", "PRODUCTION":
						o.Environment = toDeploymentEnvironmentPointer(0)
					case "development", "DEVELOPMENT":
						o.Environment = toDeploymentEnvironmentPointer(1)
					case "beta", "BETA":
						o.Environment = toDeploymentEnvironmentPointer(2)
					case "staging", "STAGING":
						o.Environment = toDeploymentEnvironmentPointer(3)
					case "test", "TEST":
						o.Environment = toDeploymentEnvironmentPointer(4)
					case "other", "OTHER":
						o.Environment = toDeploymentEnvironmentPointer(5)
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
	if val, ok := kv["shas"]; ok {
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
								panic("unsupported type for shas field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for shas field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for shas field")
				}
			}
			o.Shas = na
		}
	}
	if o.Shas == nil {
		o.Shas = make([]string, 0)
	}

	if o.StartDate == nil {
		o.StartDate = &DeploymentStartDate{}
	}

	if val, ok := kv["start_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.StartDate.FromMap(kv)
		} else if sv, ok := val.(DeploymentStartDate); ok {
			// struct
			o.StartDate = &sv
		} else if sp, ok := val.(*DeploymentStartDate); ok {
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

	if val, ok := kv["status"].(*DeploymentStatus); ok {
		o.Status = val
	} else if val, ok := kv["status"].(DeploymentStatus); ok {
		o.Status = &val
	} else {
		if val, ok := kv["status"]; ok {
			if val == nil {
				o.Status = toDeploymentStatusPointer(0)
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["DeploymentStatus"]
				}
				// this is an enum pointer
				if em, ok := val.(string); ok {
					switch em {
					case "running", "RUNNING":
						o.Status = toDeploymentStatusPointer(0)
					case "pass", "PASS":
						o.Status = toDeploymentStatusPointer(1)
					case "fail", "FAIL":
						o.Status = toDeploymentStatusPointer(2)
					case "cancel", "CANCEL":
						o.Status = toDeploymentStatusPointer(3)
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
