// DO NOT EDIT -- generated code

// Package security - the system which contains security information
package security

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

	// AdvisoryTable is the default table name
	AdvisoryTable datamodel.ModelNameType = "security_advisory"

	// AdvisoryModelName is the model name
	AdvisoryModelName datamodel.ModelNameType = "security.Advisory"
)

const (
	// AdvisoryModelCreatedDateColumn is the column json value created_date
	AdvisoryModelCreatedDateColumn = "created_date"
	// AdvisoryModelCreatedDateEpochColumn is the column json value epoch
	AdvisoryModelCreatedDateEpochColumn = "epoch"
	// AdvisoryModelCreatedDateOffsetColumn is the column json value offset
	AdvisoryModelCreatedDateOffsetColumn = "offset"
	// AdvisoryModelCreatedDateRfc3339Column is the column json value rfc3339
	AdvisoryModelCreatedDateRfc3339Column = "rfc3339"
	// AdvisoryModelCustomerIDColumn is the column json value customer_id
	AdvisoryModelCustomerIDColumn = "customer_id"
	// AdvisoryModelDescriptionColumn is the column json value description
	AdvisoryModelDescriptionColumn = "description"
	// AdvisoryModelIDColumn is the column json value id
	AdvisoryModelIDColumn = "id"
	// AdvisoryModelIdentifiersColumn is the column json value identifiers
	AdvisoryModelIdentifiersColumn = "identifiers"
	// AdvisoryModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	AdvisoryModelIntegrationInstanceIDColumn = "integration_instance_id"
	// AdvisoryModelRefIDColumn is the column json value ref_id
	AdvisoryModelRefIDColumn = "ref_id"
	// AdvisoryModelRefTypeColumn is the column json value ref_type
	AdvisoryModelRefTypeColumn = "ref_type"
	// AdvisoryModelRepoNameColumn is the column json value repo_name
	AdvisoryModelRepoNameColumn = "repo_name"
	// AdvisoryModelResolvedColumn is the column json value resolved
	AdvisoryModelResolvedColumn = "resolved"
	// AdvisoryModelResolvedDateColumn is the column json value resolved_date
	AdvisoryModelResolvedDateColumn = "resolved_date"
	// AdvisoryModelResolvedDateEpochColumn is the column json value epoch
	AdvisoryModelResolvedDateEpochColumn = "epoch"
	// AdvisoryModelResolvedDateOffsetColumn is the column json value offset
	AdvisoryModelResolvedDateOffsetColumn = "offset"
	// AdvisoryModelResolvedDateRfc3339Column is the column json value rfc3339
	AdvisoryModelResolvedDateRfc3339Column = "rfc3339"
	// AdvisoryModelSeverityColumn is the column json value severity
	AdvisoryModelSeverityColumn = "severity"
	// AdvisoryModelTitleColumn is the column json value title
	AdvisoryModelTitleColumn = "title"
	// AdvisoryModelURLColumn is the column json value url
	AdvisoryModelURLColumn = "url"
)

// AdvisoryCreatedDate represents the object structure for created_date
type AdvisoryCreatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toAdvisoryCreatedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *AdvisoryCreatedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *AdvisoryCreatedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toAdvisoryCreatedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toAdvisoryCreatedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toAdvisoryCreatedDateObject(o.Rfc3339, false),
	}
}

func (o *AdvisoryCreatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *AdvisoryCreatedDate) FromMap(kv map[string]interface{}) {

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

// AdvisoryResolvedDate represents the object structure for resolved_date
type AdvisoryResolvedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toAdvisoryResolvedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *AdvisoryResolvedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *AdvisoryResolvedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toAdvisoryResolvedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toAdvisoryResolvedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toAdvisoryResolvedDateObject(o.Rfc3339, false),
	}
}

func (o *AdvisoryResolvedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *AdvisoryResolvedDate) FromMap(kv map[string]interface{}) {

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

// AdvisorySeverity is the enumeration type for severity
type AdvisorySeverity int32

// toAdvisorySeverityPointer is the enumeration pointer type for severity
func toAdvisorySeverityPointer(v int32) *AdvisorySeverity {
	nv := AdvisorySeverity(v)
	return &nv
}

// toAdvisorySeverityEnum is the enumeration pointer wrapper for severity
func toAdvisorySeverityEnum(v *AdvisorySeverity) string {
	if v == nil {
		return toAdvisorySeverityPointer(0).String()
	}
	return v.String()
}

// UnmarshalBSONValue for unmarshaling value
func (v *AdvisorySeverity) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = AdvisorySeverity(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "CRITICAL":
			*v = AdvisorySeverity(0)
		case "HIGH":
			*v = AdvisorySeverity(1)
		case "MODERATE":
			*v = AdvisorySeverity(2)
		case "LOW":
			*v = AdvisorySeverity(3)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v *AdvisorySeverity) UnmarshalJSON(buf []byte) error {
	var val string
	if err := json.Unmarshal(buf, &val); err != nil {
		return err
	}
	switch val {
	case "CRITICAL":
		*v = 0
	case "HIGH":
		*v = 1
	case "MODERATE":
		*v = 2
	case "LOW":
		*v = 3
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v AdvisorySeverity) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("CRITICAL")
	case 1:
		return json.Marshal("HIGH")
	case 2:
		return json.Marshal("MODERATE")
	case 3:
		return json.Marshal("LOW")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for Severity
func (v AdvisorySeverity) String() string {
	switch int32(v) {
	case 0:
		return "CRITICAL"
	case 1:
		return "HIGH"
	case 2:
		return "MODERATE"
	case 3:
		return "LOW"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *AdvisorySeverity) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case AdvisorySeverity:
		*v = val
	case int32:
		*v = AdvisorySeverity(int32(val))
	case int:
		*v = AdvisorySeverity(int32(val))
	case string:
		switch val {
		case "CRITICAL":
			*v = AdvisorySeverity(0)
		case "HIGH":
			*v = AdvisorySeverity(1)
		case "MODERATE":
			*v = AdvisorySeverity(2)
		case "LOW":
			*v = AdvisorySeverity(3)
		}
	}
	return nil
}

const (
	// AdvisorySeverityCritical is the enumeration value for critical
	AdvisorySeverityCritical AdvisorySeverity = 0
	// AdvisorySeverityHigh is the enumeration value for high
	AdvisorySeverityHigh AdvisorySeverity = 1
	// AdvisorySeverityModerate is the enumeration value for moderate
	AdvisorySeverityModerate AdvisorySeverity = 2
	// AdvisorySeverityLow is the enumeration value for low
	AdvisorySeverityLow AdvisorySeverity = 3
)

// Advisory a security advisory
type Advisory struct {
	// CreatedDate the date when the advisory was created
	CreatedDate AdvisoryCreatedDate `json:"created_date" codec:"created_date" bson:"created_date" yaml:"created_date" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Description the detailed summary of the advisory
	Description string `json:"description" codec:"description" bson:"description" yaml:"description" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// Identifiers the advisory identifiers such as GHSA-v4rh-8p82-6h5w or CVE-2020-7661
	Identifiers []string `json:"identifiers" codec:"identifiers" bson:"identifiers" yaml:"identifiers" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RepoName the repo name from the source system that this advisory is associated
	RepoName string `json:"repo_name" codec:"repo_name" bson:"repo_name" yaml:"repo_name" faker:"-"`
	// Resolved if the advisory is resolved
	Resolved bool `json:"resolved" codec:"resolved" bson:"resolved" yaml:"resolved" faker:"-"`
	// ResolvedDate if the advisory is resolved, the date it was patched
	ResolvedDate AdvisoryResolvedDate `json:"resolved_date" codec:"resolved_date" bson:"resolved_date" yaml:"resolved_date" faker:"-"`
	// Severity the severity of the advisory
	Severity AdvisorySeverity `json:"severity" codec:"severity" bson:"severity" yaml:"severity" faker:"-"`
	// Title the title of the advisory
	Title string `json:"title" codec:"title" bson:"title" yaml:"title" faker:"-"`
	// URL the url to the advisory detail page
	URL string `json:"url" codec:"url" bson:"url" yaml:"url" faker:"url"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Advisory)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*Advisory)(nil)

func toAdvisoryObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *Advisory:
		return v.ToMap()

	case AdvisoryCreatedDate:
		return v.ToMap()

	case AdvisoryResolvedDate:
		return v.ToMap()

	case AdvisorySeverity:
		return v.String()

	default:
		return o
	}
}

// String returns a string representation of Advisory
func (o *Advisory) String() string {
	return fmt.Sprintf("security.Advisory<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Advisory) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *Advisory) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *Advisory) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *Advisory) GetModelName() datamodel.ModelNameType {
	return AdvisoryModelName
}

// NewAdvisoryID provides a template for generating an ID field for Advisory
func NewAdvisoryID(customerID string, refType string, refID string) string {
	return hash.Values("Advisory", customerID, refType, refID)
}

func (o *Advisory) setDefaults(frommap bool) {
	if o.Identifiers == nil {
		o.Identifiers = make([]string, 0)
	}

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Advisory", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Advisory) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Advisory) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Advisory) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *Advisory) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Advisory) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *Advisory) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Advisory) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Advisory) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *Advisory) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = AdvisoryModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Advisory) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *Advisory) GetCustomerID() string {
	return o.CustomerID
}

// SetCustomerID will return the customer_id
func (o *Advisory) SetCustomerID(id string) {
	o.CustomerID = id
}

// GetRefType will return the ref_type
func (o *Advisory) GetRefType() string {
	return o.RefType
}

// SetRefType will return the ref_type
func (o *Advisory) SetRefType(t string) {
	o.RefType = t
}

// Clone returns an exact copy of Advisory
func (o *Advisory) Clone() datamodel.Model {
	c := new(Advisory)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Advisory) Anon() datamodel.Model {
	c := new(Advisory)
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
func (o *Advisory) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Advisory) UnmarshalJSON(data []byte) error {
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
func (o *Advisory) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *Advisory) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two Advisory objects are equal
func (o *Advisory) IsEqual(other *Advisory) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Advisory) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"created_date":            toAdvisoryObject(o.CreatedDate, false),
		"customer_id":             toAdvisoryObject(o.CustomerID, false),
		"description":             toAdvisoryObject(o.Description, false),
		"id":                      toAdvisoryObject(o.ID, false),
		"identifiers":             toAdvisoryObject(o.Identifiers, false),
		"integration_instance_id": toAdvisoryObject(o.IntegrationInstanceID, true),
		"ref_id":                  toAdvisoryObject(o.RefID, false),
		"ref_type":                toAdvisoryObject(o.RefType, false),
		"repo_name":               toAdvisoryObject(o.RepoName, false),
		"resolved":                toAdvisoryObject(o.Resolved, false),
		"resolved_date":           toAdvisoryObject(o.ResolvedDate, false),

		"severity": o.Severity.String(),
		"title":    toAdvisoryObject(o.Title, false),
		"url":      toAdvisoryObject(o.URL, false),
		"hashcode": toAdvisoryObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *Advisory) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["created_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.CreatedDate.FromMap(kv)
		} else if sv, ok := val.(AdvisoryCreatedDate); ok {
			// struct
			o.CreatedDate = sv
		} else if sp, ok := val.(*AdvisoryCreatedDate); ok {
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
	if val, ok := kv["description"].(string); ok {
		o.Description = val
	} else {
		if val, ok := kv["description"]; ok {
			if val == nil {
				o.Description = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Description = fmt.Sprintf("%v", val)
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
	if val, ok := kv["identifiers"]; ok {
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
								panic("unsupported type for identifiers field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for identifiers field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for identifiers field")
				}
			}
			o.Identifiers = na
		}
	}
	if o.Identifiers == nil {
		o.Identifiers = make([]string, 0)
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
	if val, ok := kv["repo_name"].(string); ok {
		o.RepoName = val
	} else {
		if val, ok := kv["repo_name"]; ok {
			if val == nil {
				o.RepoName = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.RepoName = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["resolved"].(bool); ok {
		o.Resolved = val
	} else {
		if val, ok := kv["resolved"]; ok {
			if val == nil {
				o.Resolved = false
			} else {
				o.Resolved = number.ToBoolAny(val)
			}
		}
	}

	if val, ok := kv["resolved_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.ResolvedDate.FromMap(kv)
		} else if sv, ok := val.(AdvisoryResolvedDate); ok {
			// struct
			o.ResolvedDate = sv
		} else if sp, ok := val.(*AdvisoryResolvedDate); ok {
			// struct pointer
			o.ResolvedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.ResolvedDate.Epoch = dt.Epoch
			o.ResolvedDate.Rfc3339 = dt.Rfc3339
			o.ResolvedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.ResolvedDate.Epoch = dt.Epoch
			o.ResolvedDate.Rfc3339 = dt.Rfc3339
			o.ResolvedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.ResolvedDate.Epoch = dt.Epoch
				o.ResolvedDate.Rfc3339 = dt.Rfc3339
				o.ResolvedDate.Offset = dt.Offset
			}
		}
	} else {
		o.ResolvedDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["severity"].(AdvisorySeverity); ok {
		o.Severity = val
	} else {
		if em, ok := kv["severity"].(map[string]interface{}); ok {

			ev := em["security.severity"].(string)
			switch ev {
			case "critical", "CRITICAL":
				o.Severity = 0
			case "high", "HIGH":
				o.Severity = 1
			case "moderate", "MODERATE":
				o.Severity = 2
			case "low", "LOW":
				o.Severity = 3
			}
		}
		if em, ok := kv["severity"].(string); ok {
			switch em {
			case "critical", "CRITICAL":
				o.Severity = 0
			case "high", "HIGH":
				o.Severity = 1
			case "moderate", "MODERATE":
				o.Severity = 2
			case "low", "LOW":
				o.Severity = 3
			}
		}
	}
	if val, ok := kv["title"].(string); ok {
		o.Title = val
	} else {
		if val, ok := kv["title"]; ok {
			if val == nil {
				o.Title = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Title = fmt.Sprintf("%v", val)
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.URL = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *Advisory) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CreatedDate)
	args = append(args, o.CustomerID)
	args = append(args, o.Description)
	args = append(args, o.ID)
	args = append(args, o.Identifiers)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RepoName)
	args = append(args, o.Resolved)
	args = append(args, o.ResolvedDate)
	args = append(args, o.Severity)
	args = append(args, o.Title)
	args = append(args, o.URL)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// SetIntegrationInstanceID will set the integration instance ID
func (o *Advisory) SetIntegrationInstanceID(id string) {
	if id == "" {
		o.IntegrationInstanceID = nil
	} else {
		o.IntegrationInstanceID = &id
	}
}

// GetIntegrationInstanceID will return the integration instance ID
func (o *Advisory) GetIntegrationInstanceID() *string {
	return o.IntegrationInstanceID
}

// AdvisoryPartial is a partial struct for upsert mutations for Advisory
type AdvisoryPartial struct {
	// CreatedDate the date when the advisory was created
	CreatedDate *AdvisoryCreatedDate `json:"created_date,omitempty"`
	// Description the detailed summary of the advisory
	Description *string `json:"description,omitempty"`
	// Identifiers the advisory identifiers such as GHSA-v4rh-8p82-6h5w or CVE-2020-7661
	Identifiers []string `json:"identifiers,omitempty"`
	// RepoName the repo name from the source system that this advisory is associated
	RepoName *string `json:"repo_name,omitempty"`
	// Resolved if the advisory is resolved
	Resolved *bool `json:"resolved,omitempty"`
	// ResolvedDate if the advisory is resolved, the date it was patched
	ResolvedDate *AdvisoryResolvedDate `json:"resolved_date,omitempty"`
	// Severity the severity of the advisory
	Severity *AdvisorySeverity `json:"severity,omitempty"`
	// Title the title of the advisory
	Title *string `json:"title,omitempty"`
	// URL the url to the advisory detail page
	URL *string `json:"url,omitempty"`
}

var _ datamodel.PartialModel = (*AdvisoryPartial)(nil)

// GetModelName returns the name of the model
func (o *AdvisoryPartial) GetModelName() datamodel.ModelNameType {
	return AdvisoryModelName
}

// ToMap returns the object as a map
func (o *AdvisoryPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"created_date":  toAdvisoryObject(o.CreatedDate, true),
		"description":   toAdvisoryObject(o.Description, true),
		"identifiers":   toAdvisoryObject(o.Identifiers, true),
		"repo_name":     toAdvisoryObject(o.RepoName, true),
		"resolved":      toAdvisoryObject(o.Resolved, true),
		"resolved_date": toAdvisoryObject(o.ResolvedDate, true),

		"severity": toAdvisorySeverityEnum(o.Severity),
		"title":    toAdvisoryObject(o.Title, true),
		"url":      toAdvisoryObject(o.URL, true),
	}
	for k, v := range kv {
		if v == nil || reflect.ValueOf(v).IsZero() {
			delete(kv, k)
		} else {
			if k == "created_date" {
				if dt, ok := v.(*AdvisoryCreatedDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}

			if k == "identifiers" {
				if arr, ok := v.([]string); ok {
					if len(arr) == 0 {
						delete(kv, k)
					}
				}
			}
			if k == "resolved_date" {
				if dt, ok := v.(*AdvisoryResolvedDate); ok {
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
func (o *AdvisoryPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *AdvisoryPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *AdvisoryPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *AdvisoryPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *AdvisoryPartial) FromMap(kv map[string]interface{}) {

	if o.CreatedDate == nil {
		o.CreatedDate = &AdvisoryCreatedDate{}
	}

	if val, ok := kv["created_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.CreatedDate.FromMap(kv)
		} else if sv, ok := val.(AdvisoryCreatedDate); ok {
			// struct
			o.CreatedDate = &sv
		} else if sp, ok := val.(*AdvisoryCreatedDate); ok {
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

	if val, ok := kv["description"].(*string); ok {
		o.Description = val
	} else if val, ok := kv["description"].(string); ok {
		o.Description = &val
	} else {
		if val, ok := kv["description"]; ok {
			if val == nil {
				o.Description = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Description = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["identifiers"]; ok {
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
								panic("unsupported type for identifiers field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for identifiers field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for identifiers field")
				}
			}
			o.Identifiers = na
		}
	}
	if o.Identifiers == nil {
		o.Identifiers = make([]string, 0)
	}
	if val, ok := kv["repo_name"].(*string); ok {
		o.RepoName = val
	} else if val, ok := kv["repo_name"].(string); ok {
		o.RepoName = &val
	} else {
		if val, ok := kv["repo_name"]; ok {
			if val == nil {
				o.RepoName = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.RepoName = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["resolved"].(*bool); ok {
		o.Resolved = val
	} else if val, ok := kv["resolved"].(bool); ok {
		o.Resolved = &val
	} else {
		if val, ok := kv["resolved"]; ok {
			if val == nil {
				o.Resolved = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Resolved = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}

	if o.ResolvedDate == nil {
		o.ResolvedDate = &AdvisoryResolvedDate{}
	}

	if val, ok := kv["resolved_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.ResolvedDate.FromMap(kv)
		} else if sv, ok := val.(AdvisoryResolvedDate); ok {
			// struct
			o.ResolvedDate = &sv
		} else if sp, ok := val.(*AdvisoryResolvedDate); ok {
			// struct pointer
			o.ResolvedDate = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.ResolvedDate.Epoch = dt.Epoch
			o.ResolvedDate.Rfc3339 = dt.Rfc3339
			o.ResolvedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.ResolvedDate.Epoch = dt.Epoch
			o.ResolvedDate.Rfc3339 = dt.Rfc3339
			o.ResolvedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.ResolvedDate.Epoch = dt.Epoch
				o.ResolvedDate.Rfc3339 = dt.Rfc3339
				o.ResolvedDate.Offset = dt.Offset
			}
		}
	} else {
		o.ResolvedDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["severity"].(*AdvisorySeverity); ok {
		o.Severity = val
	} else if val, ok := kv["severity"].(AdvisorySeverity); ok {
		o.Severity = &val
	} else {
		if val, ok := kv["severity"]; ok {
			if val == nil {
				o.Severity = toAdvisorySeverityPointer(0)
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["AdvisorySeverity"]
				}
				// this is an enum pointer
				if em, ok := val.(string); ok {
					switch em {
					case "critical", "CRITICAL":
						o.Severity = toAdvisorySeverityPointer(0)
					case "high", "HIGH":
						o.Severity = toAdvisorySeverityPointer(1)
					case "moderate", "MODERATE":
						o.Severity = toAdvisorySeverityPointer(2)
					case "low", "LOW":
						o.Severity = toAdvisorySeverityPointer(3)
					}
				}
			}
		}
	}
	if val, ok := kv["title"].(*string); ok {
		o.Title = val
	} else if val, ok := kv["title"].(string); ok {
		o.Title = &val
	} else {
		if val, ok := kv["title"]; ok {
			if val == nil {
				o.Title = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Title = pstrings.Pointer(fmt.Sprintf("%v", val))
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
