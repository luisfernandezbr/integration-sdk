// DO NOT EDIT -- generated code

// Package codequality - the system which contains code quality
package codequality

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

	// MetricTable is the default table name
	MetricTable datamodel.ModelNameType = "codequality_metric"

	// MetricModelName is the model name
	MetricModelName datamodel.ModelNameType = "codequality.Metric"
)

const (
	// MetricModelBranchColumn is the column json value branch
	MetricModelBranchColumn = "branch"
	// MetricModelCommitIDColumn is the column json value commit_id
	MetricModelCommitIDColumn = "commit_id"
	// MetricModelCommitShaColumn is the column json value commit_sha
	MetricModelCommitShaColumn = "commit_sha"
	// MetricModelCreatedDateColumn is the column json value created_date
	MetricModelCreatedDateColumn = "created_date"
	// MetricModelCreatedDateEpochColumn is the column json value epoch
	MetricModelCreatedDateEpochColumn = "epoch"
	// MetricModelCreatedDateOffsetColumn is the column json value offset
	MetricModelCreatedDateOffsetColumn = "offset"
	// MetricModelCreatedDateRfc3339Column is the column json value rfc3339
	MetricModelCreatedDateRfc3339Column = "rfc3339"
	// MetricModelCustomerIDColumn is the column json value customer_id
	MetricModelCustomerIDColumn = "customer_id"
	// MetricModelIDColumn is the column json value id
	MetricModelIDColumn = "id"
	// MetricModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	MetricModelIntegrationInstanceIDColumn = "integration_instance_id"
	// MetricModelNameColumn is the column json value name
	MetricModelNameColumn = "name"
	// MetricModelProjectIDColumn is the column json value project_id
	MetricModelProjectIDColumn = "project_id"
	// MetricModelPullRequestIDColumn is the column json value pull_request_id
	MetricModelPullRequestIDColumn = "pull_request_id"
	// MetricModelPullRequestRefIDColumn is the column json value pull_request_ref_id
	MetricModelPullRequestRefIDColumn = "pull_request_ref_id"
	// MetricModelRefIDColumn is the column json value ref_id
	MetricModelRefIDColumn = "ref_id"
	// MetricModelRefTypeColumn is the column json value ref_type
	MetricModelRefTypeColumn = "ref_type"
	// MetricModelRepoIDColumn is the column json value repo_id
	MetricModelRepoIDColumn = "repo_id"
	// MetricModelStatusColumn is the column json value status
	MetricModelStatusColumn = "status"
	// MetricModelValueColumn is the column json value value
	MetricModelValueColumn = "value"
)

// MetricCreatedDate represents the object structure for created_date
type MetricCreatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toMetricCreatedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *MetricCreatedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *MetricCreatedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toMetricCreatedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toMetricCreatedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toMetricCreatedDateObject(o.Rfc3339, false),
	}
}

func (o *MetricCreatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *MetricCreatedDate) FromMap(kv map[string]interface{}) {

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

// MetricStatus is the enumeration type for status
type MetricStatus int32

// toMetricStatusPointer is the enumeration pointer type for status
func toMetricStatusPointer(v int32) *MetricStatus {
	nv := MetricStatus(v)
	return &nv
}

// toMetricStatusEnum is the enumeration pointer wrapper for status
func toMetricStatusEnum(v *MetricStatus) string {
	if v == nil {
		return toMetricStatusPointer(0).String()
	}
	return v.String()
}

// UnmarshalBSONValue for unmarshaling value
func (v *MetricStatus) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = MetricStatus(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "UNSUPPORTED":
			*v = MetricStatus(0)
		case "NOT_ANALYSED":
			*v = MetricStatus(1)
		case "ANALYSED":
			*v = MetricStatus(2)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v *MetricStatus) UnmarshalJSON(buf []byte) error {
	var val string
	if err := json.Unmarshal(buf, &val); err != nil {
		return err
	}
	switch val {
	case "UNSUPPORTED":
		*v = 0
	case "NOT_ANALYSED":
		*v = 1
	case "ANALYSED":
		*v = 2
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v MetricStatus) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("UNSUPPORTED")
	case 1:
		return json.Marshal("NOT_ANALYSED")
	case 2:
		return json.Marshal("ANALYSED")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for Status
func (v MetricStatus) String() string {
	switch int32(v) {
	case 0:
		return "UNSUPPORTED"
	case 1:
		return "NOT_ANALYSED"
	case 2:
		return "ANALYSED"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *MetricStatus) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case MetricStatus:
		*v = val
	case int32:
		*v = MetricStatus(int32(val))
	case int:
		*v = MetricStatus(int32(val))
	case string:
		switch val {
		case "UNSUPPORTED":
			*v = MetricStatus(0)
		case "NOT_ANALYSED":
			*v = MetricStatus(1)
		case "ANALYSED":
			*v = MetricStatus(2)
		}
	}
	return nil
}

const (
	// MetricStatusUnsupported is the enumeration value for unsupported
	MetricStatusUnsupported MetricStatus = 0
	// MetricStatusNotAnalysed is the enumeration value for not_analysed
	MetricStatusNotAnalysed MetricStatus = 1
	// MetricStatusAnalysed is the enumeration value for analysed
	MetricStatusAnalysed MetricStatus = 2
)

// Metric individual metric details
type Metric struct {
	// Branch branch name
	Branch *string `json:"branch,omitempty" codec:"branch,omitempty" bson:"branch" yaml:"branch,omitempty" faker:"-"`
	// CommitID the commit id
	CommitID *string `json:"commit_id,omitempty" codec:"commit_id,omitempty" bson:"commit_id" yaml:"commit_id,omitempty" faker:"-"`
	// CommitSha commit sha
	CommitSha *string `json:"commit_sha,omitempty" codec:"commit_sha,omitempty" bson:"commit_sha" yaml:"commit_sha,omitempty" faker:"-"`
	// CreatedDate the date when the metric was created
	CreatedDate MetricCreatedDate `json:"created_date" codec:"created_date" bson:"created_date" yaml:"created_date" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// Name the metric name
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
	// ProjectID the project id
	ProjectID string `json:"project_id" codec:"project_id" bson:"project_id" yaml:"project_id" faker:"-"`
	// PullRequestID the pull request this commit is associated with
	PullRequestID *string `json:"pull_request_id,omitempty" codec:"pull_request_id,omitempty" bson:"pull_request_id" yaml:"pull_request_id,omitempty" faker:"-"`
	// PullRequestRefID the original pull request id this commit is associated with
	PullRequestRefID *string `json:"pull_request_ref_id,omitempty" codec:"pull_request_ref_id,omitempty" bson:"pull_request_ref_id" yaml:"pull_request_ref_id,omitempty" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RepoID repo id
	RepoID *string `json:"repo_id,omitempty" codec:"repo_id,omitempty" bson:"repo_id" yaml:"repo_id,omitempty" faker:"-"`
	// Status status of the analysis for this commit
	Status MetricStatus `json:"status" codec:"status" bson:"status" yaml:"status" faker:"-"`
	// Value the value of the metric
	Value string `json:"value" codec:"value" bson:"value" yaml:"value" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Metric)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*Metric)(nil)

func toMetricObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *Metric:
		return v.ToMap()

	case MetricCreatedDate:
		return v.ToMap()

	case MetricStatus:
		return v.String()

	default:
		return o
	}
}

// String returns a string representation of Metric
func (o *Metric) String() string {
	return fmt.Sprintf("codequality.Metric<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Metric) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *Metric) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *Metric) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *Metric) GetModelName() datamodel.ModelNameType {
	return MetricModelName
}

// NewMetricID provides a template for generating an ID field for Metric
func NewMetricID(customerID string, refType string, refID string) string {
	return hash.Values("Metric", customerID, refType, refID)
}

func (o *Metric) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Metric", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Metric) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Metric) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Metric) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *Metric) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Metric) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *Metric) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Metric) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Metric) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *Metric) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = MetricModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Metric) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *Metric) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of Metric
func (o *Metric) Clone() datamodel.Model {
	c := new(Metric)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Metric) Anon() datamodel.Model {
	c := new(Metric)
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
func (o *Metric) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Metric) UnmarshalJSON(data []byte) error {
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
func (o *Metric) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *Metric) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two Metric objects are equal
func (o *Metric) IsEqual(other *Metric) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Metric) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"branch":                  toMetricObject(o.Branch, true),
		"commit_id":               toMetricObject(o.CommitID, true),
		"commit_sha":              toMetricObject(o.CommitSha, true),
		"created_date":            toMetricObject(o.CreatedDate, false),
		"customer_id":             toMetricObject(o.CustomerID, false),
		"id":                      toMetricObject(o.ID, false),
		"integration_instance_id": toMetricObject(o.IntegrationInstanceID, true),
		"name":                    toMetricObject(o.Name, false),
		"project_id":              toMetricObject(o.ProjectID, false),
		"pull_request_id":         toMetricObject(o.PullRequestID, true),
		"pull_request_ref_id":     toMetricObject(o.PullRequestRefID, true),
		"ref_id":                  toMetricObject(o.RefID, false),
		"ref_type":                toMetricObject(o.RefType, false),
		"repo_id":                 toMetricObject(o.RepoID, true),

		"status":   o.Status.String(),
		"value":    toMetricObject(o.Value, false),
		"hashcode": toMetricObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *Metric) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["branch"].(*string); ok {
		o.Branch = val
	} else if val, ok := kv["branch"].(string); ok {
		o.Branch = &val
	} else {
		if val, ok := kv["branch"]; ok {
			if val == nil {
				o.Branch = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Branch = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["commit_id"].(*string); ok {
		o.CommitID = val
	} else if val, ok := kv["commit_id"].(string); ok {
		o.CommitID = &val
	} else {
		if val, ok := kv["commit_id"]; ok {
			if val == nil {
				o.CommitID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.CommitID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["commit_sha"].(*string); ok {
		o.CommitSha = val
	} else if val, ok := kv["commit_sha"].(string); ok {
		o.CommitSha = &val
	} else {
		if val, ok := kv["commit_sha"]; ok {
			if val == nil {
				o.CommitSha = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.CommitSha = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if val, ok := kv["created_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.CreatedDate.FromMap(kv)
		} else if sv, ok := val.(MetricCreatedDate); ok {
			// struct
			o.CreatedDate = sv
		} else if sp, ok := val.(*MetricCreatedDate); ok {
			// struct pointer
			o.CreatedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.CreatedDate.Epoch = dt.Epoch
			o.CreatedDate.Rfc3339 = dt.Rfc3339
			o.CreatedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
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
	if val, ok := kv["name"].(string); ok {
		o.Name = val
	} else {
		if val, ok := kv["name"]; ok {
			if val == nil {
				o.Name = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Name = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["project_id"].(string); ok {
		o.ProjectID = val
	} else {
		if val, ok := kv["project_id"]; ok {
			if val == nil {
				o.ProjectID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.ProjectID = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["pull_request_id"].(*string); ok {
		o.PullRequestID = val
	} else if val, ok := kv["pull_request_id"].(string); ok {
		o.PullRequestID = &val
	} else {
		if val, ok := kv["pull_request_id"]; ok {
			if val == nil {
				o.PullRequestID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.PullRequestID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["pull_request_ref_id"].(*string); ok {
		o.PullRequestRefID = val
	} else if val, ok := kv["pull_request_ref_id"].(string); ok {
		o.PullRequestRefID = &val
	} else {
		if val, ok := kv["pull_request_ref_id"]; ok {
			if val == nil {
				o.PullRequestRefID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.PullRequestRefID = pstrings.Pointer(fmt.Sprintf("%v", val))
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
	if val, ok := kv["repo_id"].(*string); ok {
		o.RepoID = val
	} else if val, ok := kv["repo_id"].(string); ok {
		o.RepoID = &val
	} else {
		if val, ok := kv["repo_id"]; ok {
			if val == nil {
				o.RepoID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.RepoID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["status"].(MetricStatus); ok {
		o.Status = val
	} else {
		if em, ok := kv["status"].(map[string]interface{}); ok {

			ev := em["codequality.status"].(string)
			switch ev {
			case "unsupported", "UNSUPPORTED":
				o.Status = 0
			case "not_analysed", "NOT_ANALYSED":
				o.Status = 1
			case "analysed", "ANALYSED":
				o.Status = 2
			}
		}
		if em, ok := kv["status"].(string); ok {
			switch em {
			case "unsupported", "UNSUPPORTED":
				o.Status = 0
			case "not_analysed", "NOT_ANALYSED":
				o.Status = 1
			case "analysed", "ANALYSED":
				o.Status = 2
			}
		}
	}
	if val, ok := kv["value"].(string); ok {
		o.Value = val
	} else {
		if val, ok := kv["value"]; ok {
			if val == nil {
				o.Value = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Value = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *Metric) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Branch)
	args = append(args, o.CommitID)
	args = append(args, o.CommitSha)
	args = append(args, o.CreatedDate)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.Name)
	args = append(args, o.ProjectID)
	args = append(args, o.PullRequestID)
	args = append(args, o.PullRequestRefID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RepoID)
	args = append(args, o.Status)
	args = append(args, o.Value)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// SetIntegrationInstanceID will set the integration instance ID
func (o *Metric) SetIntegrationInstanceID(id string) {
	if id == "" {
		o.IntegrationInstanceID = nil
	} else {
		o.IntegrationInstanceID = &id
	}
}

// GetIntegrationInstanceID will return the integration instance ID
func (o *Metric) GetIntegrationInstanceID() *string {
	return o.IntegrationInstanceID
}

// MetricPartial is a partial struct for upsert mutations for Metric
type MetricPartial struct {
	// Branch branch name
	Branch *string `json:"branch,omitempty"`
	// CommitID the commit id
	CommitID *string `json:"commit_id,omitempty"`
	// CommitSha commit sha
	CommitSha *string `json:"commit_sha,omitempty"`
	// CreatedDate the date when the metric was created
	CreatedDate *MetricCreatedDate `json:"created_date,omitempty"`
	// Name the metric name
	Name *string `json:"name,omitempty"`
	// ProjectID the project id
	ProjectID *string `json:"project_id,omitempty"`
	// PullRequestID the pull request this commit is associated with
	PullRequestID *string `json:"pull_request_id,omitempty"`
	// PullRequestRefID the original pull request id this commit is associated with
	PullRequestRefID *string `json:"pull_request_ref_id,omitempty"`
	// RepoID repo id
	RepoID *string `json:"repo_id,omitempty"`
	// Status status of the analysis for this commit
	Status *MetricStatus `json:"status,omitempty"`
	// Value the value of the metric
	Value *string `json:"value,omitempty"`
}

var _ datamodel.PartialModel = (*MetricPartial)(nil)

// GetModelName returns the name of the model
func (o *MetricPartial) GetModelName() datamodel.ModelNameType {
	return MetricModelName
}

// ToMap returns the object as a map
func (o *MetricPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"branch":              toMetricObject(o.Branch, true),
		"commit_id":           toMetricObject(o.CommitID, true),
		"commit_sha":          toMetricObject(o.CommitSha, true),
		"created_date":        toMetricObject(o.CreatedDate, true),
		"name":                toMetricObject(o.Name, true),
		"project_id":          toMetricObject(o.ProjectID, true),
		"pull_request_id":     toMetricObject(o.PullRequestID, true),
		"pull_request_ref_id": toMetricObject(o.PullRequestRefID, true),
		"repo_id":             toMetricObject(o.RepoID, true),

		"status": toMetricStatusEnum(o.Status),
		"value":  toMetricObject(o.Value, true),
	}
	for k, v := range kv {
		if v == nil || reflect.ValueOf(v).IsZero() {
			delete(kv, k)
		} else {
			if k == "created_date" {
				if dt, ok := v.(*MetricCreatedDate); ok {
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
func (o *MetricPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *MetricPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *MetricPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *MetricPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *MetricPartial) FromMap(kv map[string]interface{}) {
	if val, ok := kv["branch"].(*string); ok {
		o.Branch = val
	} else if val, ok := kv["branch"].(string); ok {
		o.Branch = &val
	} else {
		if val, ok := kv["branch"]; ok {
			if val == nil {
				o.Branch = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Branch = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["commit_id"].(*string); ok {
		o.CommitID = val
	} else if val, ok := kv["commit_id"].(string); ok {
		o.CommitID = &val
	} else {
		if val, ok := kv["commit_id"]; ok {
			if val == nil {
				o.CommitID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.CommitID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["commit_sha"].(*string); ok {
		o.CommitSha = val
	} else if val, ok := kv["commit_sha"].(string); ok {
		o.CommitSha = &val
	} else {
		if val, ok := kv["commit_sha"]; ok {
			if val == nil {
				o.CommitSha = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.CommitSha = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if o.CreatedDate == nil {
		o.CreatedDate = &MetricCreatedDate{}
	}

	if val, ok := kv["created_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.CreatedDate.FromMap(kv)
		} else if sv, ok := val.(MetricCreatedDate); ok {
			// struct
			o.CreatedDate = &sv
		} else if sp, ok := val.(*MetricCreatedDate); ok {
			// struct pointer
			o.CreatedDate = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.CreatedDate.Epoch = dt.Epoch
			o.CreatedDate.Rfc3339 = dt.Rfc3339
			o.CreatedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
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

	if val, ok := kv["name"].(*string); ok {
		o.Name = val
	} else if val, ok := kv["name"].(string); ok {
		o.Name = &val
	} else {
		if val, ok := kv["name"]; ok {
			if val == nil {
				o.Name = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Name = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["project_id"].(*string); ok {
		o.ProjectID = val
	} else if val, ok := kv["project_id"].(string); ok {
		o.ProjectID = &val
	} else {
		if val, ok := kv["project_id"]; ok {
			if val == nil {
				o.ProjectID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.ProjectID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["pull_request_id"].(*string); ok {
		o.PullRequestID = val
	} else if val, ok := kv["pull_request_id"].(string); ok {
		o.PullRequestID = &val
	} else {
		if val, ok := kv["pull_request_id"]; ok {
			if val == nil {
				o.PullRequestID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.PullRequestID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["pull_request_ref_id"].(*string); ok {
		o.PullRequestRefID = val
	} else if val, ok := kv["pull_request_ref_id"].(string); ok {
		o.PullRequestRefID = &val
	} else {
		if val, ok := kv["pull_request_ref_id"]; ok {
			if val == nil {
				o.PullRequestRefID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.PullRequestRefID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["repo_id"].(*string); ok {
		o.RepoID = val
	} else if val, ok := kv["repo_id"].(string); ok {
		o.RepoID = &val
	} else {
		if val, ok := kv["repo_id"]; ok {
			if val == nil {
				o.RepoID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.RepoID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["status"].(*MetricStatus); ok {
		o.Status = val
	} else if val, ok := kv["status"].(MetricStatus); ok {
		o.Status = &val
	} else {
		if val, ok := kv["status"]; ok {
			if val == nil {
				o.Status = toMetricStatusPointer(0)
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["MetricStatus"]
				}
				// this is an enum pointer
				if em, ok := val.(string); ok {
					switch em {
					case "unsupported", "UNSUPPORTED":
						o.Status = toMetricStatusPointer(0)
					case "not_analysed", "NOT_ANALYSED":
						o.Status = toMetricStatusPointer(1)
					case "analysed", "ANALYSED":
						o.Status = toMetricStatusPointer(2)
					}
				}
			}
		}
	}
	if val, ok := kv["value"].(*string); ok {
		o.Value = val
	} else if val, ok := kv["value"].(string); ok {
		o.Value = &val
	} else {
		if val, ok := kv["value"]; ok {
			if val == nil {
				o.Value = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Value = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	o.setDefaults(false)
}
