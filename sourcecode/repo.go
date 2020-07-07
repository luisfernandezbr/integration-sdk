// DO NOT EDIT -- generated code

// Package sourcecode - the system which contains source code
package sourcecode

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
	// RepoTopic is the default topic name
	RepoTopic datamodel.TopicNameType = "sourcecode_Repo"

	// RepoTable is the default table name
	RepoTable datamodel.ModelNameType = "sourcecode_repo"

	// RepoModelName is the model name
	RepoModelName datamodel.ModelNameType = "sourcecode.Repo"
)

const (
	// RepoModelActiveColumn is the column json value active
	RepoModelActiveColumn = "active"
	// RepoModelAffiliationColumn is the column json value affiliation
	RepoModelAffiliationColumn = "affiliation"
	// RepoModelCustomerIDColumn is the column json value customer_id
	RepoModelCustomerIDColumn = "customer_id"
	// RepoModelDefaultBranchColumn is the column json value default_branch
	RepoModelDefaultBranchColumn = "default_branch"
	// RepoModelDescriptionColumn is the column json value description
	RepoModelDescriptionColumn = "description"
	// RepoModelIDColumn is the column json value id
	RepoModelIDColumn = "id"
	// RepoModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	RepoModelIntegrationInstanceIDColumn = "integration_instance_id"
	// RepoModelLanguageColumn is the column json value language
	RepoModelLanguageColumn = "language"
	// RepoModelNameColumn is the column json value name
	RepoModelNameColumn = "name"
	// RepoModelRefIDColumn is the column json value ref_id
	RepoModelRefIDColumn = "ref_id"
	// RepoModelRefTypeColumn is the column json value ref_type
	RepoModelRefTypeColumn = "ref_type"
	// RepoModelUpdatedAtColumn is the column json value updated_ts
	RepoModelUpdatedAtColumn = "updated_ts"
	// RepoModelURLColumn is the column json value url
	RepoModelURLColumn = "url"
	// RepoModelVisibilityColumn is the column json value visibility
	RepoModelVisibilityColumn = "visibility"
)

// RepoAffiliation is the enumeration type for affiliation
type RepoAffiliation int32

// toRepoAffiliationPointer is the enumeration pointer type for affiliation
func toRepoAffiliationPointer(v int32) *RepoAffiliation {
	nv := RepoAffiliation(v)
	return &nv
}

// toRepoAffiliationEnum is the enumeration pointer wrapper for affiliation
func toRepoAffiliationEnum(v *RepoAffiliation) string {
	if v == nil {
		return toRepoAffiliationPointer(0).String()
	}
	return v.String()
}

// UnmarshalBSONValue for unmarshaling value
func (v *RepoAffiliation) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = RepoAffiliation(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "ORGANIZATION":
			*v = RepoAffiliation(0)
		case "USER":
			*v = RepoAffiliation(1)
		case "THIRDPARTY":
			*v = RepoAffiliation(2)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v RepoAffiliation) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "ORGANIZATION":
		v = 0
	case "USER":
		v = 1
	case "THIRDPARTY":
		v = 2
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v RepoAffiliation) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("ORGANIZATION")
	case 1:
		return json.Marshal("USER")
	case 2:
		return json.Marshal("THIRDPARTY")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for Affiliation
func (v RepoAffiliation) String() string {
	switch int32(v) {
	case 0:
		return "ORGANIZATION"
	case 1:
		return "USER"
	case 2:
		return "THIRDPARTY"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *RepoAffiliation) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case RepoAffiliation:
		*v = val
	case int32:
		*v = RepoAffiliation(int32(val))
	case int:
		*v = RepoAffiliation(int32(val))
	case string:
		switch val {
		case "ORGANIZATION":
			*v = RepoAffiliation(0)
		case "USER":
			*v = RepoAffiliation(1)
		case "THIRDPARTY":
			*v = RepoAffiliation(2)
		}
	}
	return nil
}

const (
	// RepoAffiliationOrganization is the enumeration value for organization
	RepoAffiliationOrganization RepoAffiliation = 0
	// RepoAffiliationUser is the enumeration value for user
	RepoAffiliationUser RepoAffiliation = 1
	// RepoAffiliationThirdparty is the enumeration value for thirdparty
	RepoAffiliationThirdparty RepoAffiliation = 2
)

// RepoVisibility is the enumeration type for visibility
type RepoVisibility int32

// toRepoVisibilityPointer is the enumeration pointer type for visibility
func toRepoVisibilityPointer(v int32) *RepoVisibility {
	nv := RepoVisibility(v)
	return &nv
}

// toRepoVisibilityEnum is the enumeration pointer wrapper for visibility
func toRepoVisibilityEnum(v *RepoVisibility) string {
	if v == nil {
		return toRepoVisibilityPointer(0).String()
	}
	return v.String()
}

// UnmarshalBSONValue for unmarshaling value
func (v *RepoVisibility) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = RepoVisibility(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "PRIVATE":
			*v = RepoVisibility(0)
		case "PUBLIC":
			*v = RepoVisibility(1)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v RepoVisibility) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "PRIVATE":
		v = 0
	case "PUBLIC":
		v = 1
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v RepoVisibility) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("PRIVATE")
	case 1:
		return json.Marshal("PUBLIC")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for Visibility
func (v RepoVisibility) String() string {
	switch int32(v) {
	case 0:
		return "PRIVATE"
	case 1:
		return "PUBLIC"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *RepoVisibility) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case RepoVisibility:
		*v = val
	case int32:
		*v = RepoVisibility(int32(val))
	case int:
		*v = RepoVisibility(int32(val))
	case string:
		switch val {
		case "PRIVATE":
			*v = RepoVisibility(0)
		case "PUBLIC":
			*v = RepoVisibility(1)
		}
	}
	return nil
}

const (
	// RepoVisibilityPrivate is the enumeration value for private
	RepoVisibilityPrivate RepoVisibility = 0
	// RepoVisibilityPublic is the enumeration value for public
	RepoVisibilityPublic RepoVisibility = 1
)

// Repo the repo holds source code
type Repo struct {
	// Active the status of the repo
	Active bool `json:"active" codec:"active" bson:"active" yaml:"active" faker:"-"`
	// Affiliation the affiliation to the repo owner
	Affiliation RepoAffiliation `json:"affiliation" codec:"affiliation" bson:"affiliation" yaml:"affiliation" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// DefaultBranch the repo default branch
	DefaultBranch string `json:"default_branch" codec:"default_branch" bson:"default_branch" yaml:"default_branch" faker:"-"`
	// Description brief explanation of the repo
	Description string `json:"description" codec:"description" bson:"description" yaml:"description" faker:"sentence"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// Language the programming language the source code is primarily written in
	Language string `json:"language" codec:"language" bson:"language" yaml:"language" faker:"-"`
	// Name the name of the repo
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"repo"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// URL the url to the repo home page
	URL string `json:"url" codec:"url" bson:"url" yaml:"url" faker:"url"`
	// Visibility the visibility of the repo
	Visibility RepoVisibility `json:"visibility" codec:"visibility" bson:"visibility" yaml:"visibility" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Repo)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*Repo)(nil)

func toRepoObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *Repo:
		return v.ToMap()

	case RepoAffiliation:
		return v.String()

	case RepoVisibility:
		return v.String()
	default:
		return o
	}
}

// String returns a string representation of Repo
func (o *Repo) String() string {
	return fmt.Sprintf("sourcecode.Repo<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Repo) GetTopicName() datamodel.TopicNameType {
	return RepoTopic
}

// GetStreamName returns the name of the stream
func (o *Repo) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *Repo) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *Repo) GetModelName() datamodel.ModelNameType {
	return RepoModelName
}

// NewRepoID provides a template for generating an ID field for Repo
func NewRepoID(customerID string, refType string, refID string) string {
	return hash.Values("Repo", customerID, refType, refID)
}

func (o *Repo) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Repo", o.CustomerID, o.RefType, o.GetRefID())
	}

	{
		v := "master"

		o.DefaultBranch = v

	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Repo) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Repo) GetTopicKey() string {
	var i interface{} = o.ID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Repo) GetTimestamp() time.Time {
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
	panic("not sure how to handle the date time format for Repo")
}

// GetRefID returns the RefID for the object
func (o *Repo) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Repo) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *Repo) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Repo) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Repo) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *Repo) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = RepoModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Repo) GetTopicConfig() *datamodel.ModelTopicConfig {
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
		Key:               "id",
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
func (o *Repo) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of Repo
func (o *Repo) Clone() datamodel.Model {
	c := new(Repo)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Repo) Anon() datamodel.Model {
	c := new(Repo)
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
func (o *Repo) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Repo) UnmarshalJSON(data []byte) error {
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
func (o *Repo) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *Repo) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two Repo objects are equal
func (o *Repo) IsEqual(other *Repo) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Repo) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"active": toRepoObject(o.Active, false),

		"affiliation":             o.Affiliation.String(),
		"customer_id":             toRepoObject(o.CustomerID, false),
		"default_branch":          toRepoObject(o.DefaultBranch, false),
		"description":             toRepoObject(o.Description, false),
		"id":                      toRepoObject(o.ID, false),
		"integration_instance_id": toRepoObject(o.IntegrationInstanceID, true),
		"language":                toRepoObject(o.Language, false),
		"name":                    toRepoObject(o.Name, false),
		"ref_id":                  toRepoObject(o.RefID, false),
		"ref_type":                toRepoObject(o.RefType, false),
		"updated_ts":              toRepoObject(o.UpdatedAt, false),
		"url":                     toRepoObject(o.URL, false),

		"visibility": o.Visibility.String(),
		"hashcode":   toRepoObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *Repo) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["active"].(bool); ok {
		o.Active = val
	} else {
		if val, ok := kv["active"]; ok {
			if val == nil {
				o.Active = false
			} else {
				o.Active = number.ToBoolAny(val)
			}
		}
	}
	if val, ok := kv["affiliation"].(RepoAffiliation); ok {
		o.Affiliation = val
	} else {
		if em, ok := kv["affiliation"].(map[string]interface{}); ok {

			ev := em["sourcecode.affiliation"].(string)
			switch ev {
			case "organization", "ORGANIZATION":
				o.Affiliation = 0
			case "user", "USER":
				o.Affiliation = 1
			case "thirdparty", "THIRDPARTY":
				o.Affiliation = 2
			}
		}
		if em, ok := kv["affiliation"].(string); ok {
			switch em {
			case "organization", "ORGANIZATION":
				o.Affiliation = 0
			case "user", "USER":
				o.Affiliation = 1
			case "thirdparty", "THIRDPARTY":
				o.Affiliation = 2
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
	if val, ok := kv["default_branch"].(string); ok {
		o.DefaultBranch = val
	} else {
		if val, ok := kv["default_branch"]; ok {
			if val == nil {
				o.DefaultBranch = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.DefaultBranch = fmt.Sprintf("%v", val)
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
	if val, ok := kv["language"].(string); ok {
		o.Language = val
	} else {
		if val, ok := kv["language"]; ok {
			if val == nil {
				o.Language = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Language = fmt.Sprintf("%v", val)
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
	if val, ok := kv["updated_ts"].(int64); ok {
		o.UpdatedAt = val
	} else {
		if val, ok := kv["updated_ts"]; ok {
			if val == nil {
				o.UpdatedAt = 0
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
	if val, ok := kv["visibility"].(RepoVisibility); ok {
		o.Visibility = val
	} else {
		if em, ok := kv["visibility"].(map[string]interface{}); ok {

			ev := em["sourcecode.visibility"].(string)
			switch ev {
			case "private", "PRIVATE":
				o.Visibility = 0
			case "public", "PUBLIC":
				o.Visibility = 1
			}
		}
		if em, ok := kv["visibility"].(string); ok {
			switch em {
			case "private", "PRIVATE":
				o.Visibility = 0
			case "public", "PUBLIC":
				o.Visibility = 1
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *Repo) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Active)
	args = append(args, o.Affiliation)
	args = append(args, o.CustomerID)
	args = append(args, o.DefaultBranch)
	args = append(args, o.Description)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.Language)
	args = append(args, o.Name)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.UpdatedAt)
	args = append(args, o.URL)
	args = append(args, o.Visibility)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// RepoPartial is a partial struct for upsert mutations for Repo
type RepoPartial struct {
	// Active the status of the repo
	Active *bool `json:"active,omitempty"`
	// Affiliation the affiliation to the repo owner
	Affiliation *RepoAffiliation `json:"affiliation,omitempty"`
	// DefaultBranch the repo default branch
	DefaultBranch *string `json:"default_branch,omitempty"`
	// Description brief explanation of the repo
	Description *string `json:"description,omitempty"`
	// Language the programming language the source code is primarily written in
	Language *string `json:"language,omitempty"`
	// Name the name of the repo
	Name *string `json:"name,omitempty"`
	// URL the url to the repo home page
	URL *string `json:"url,omitempty"`
	// Visibility the visibility of the repo
	Visibility *RepoVisibility `json:"visibility,omitempty"`
}

var _ datamodel.PartialModel = (*RepoPartial)(nil)

// GetModelName returns the name of the model
func (o *RepoPartial) GetModelName() datamodel.ModelNameType {
	return RepoModelName
}

// ToMap returns the object as a map
func (o *RepoPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"active": toRepoObject(o.Active, true),

		"affiliation":    toRepoAffiliationEnum(o.Affiliation),
		"default_branch": toRepoObject(o.DefaultBranch, true),
		"description":    toRepoObject(o.Description, true),
		"language":       toRepoObject(o.Language, true),
		"name":           toRepoObject(o.Name, true),
		"url":            toRepoObject(o.URL, true),

		"visibility": toRepoVisibilityEnum(o.Visibility),
	}
	for k, v := range kv {
		if v == nil || reflect.ValueOf(v).IsZero() {
			delete(kv, k)
		} else {
		}
	}
	return kv
}

// Stringify returns the object in JSON format as a string
func (o *RepoPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *RepoPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *RepoPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *RepoPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *RepoPartial) FromMap(kv map[string]interface{}) {
	if val, ok := kv["active"].(*bool); ok {
		o.Active = val
	} else if val, ok := kv["active"].(bool); ok {
		o.Active = &val
	} else {
		if val, ok := kv["active"]; ok {
			if val == nil {
				o.Active = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Active = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}
	if val, ok := kv["affiliation"].(*RepoAffiliation); ok {
		o.Affiliation = val
	} else if val, ok := kv["affiliation"].(RepoAffiliation); ok {
		o.Affiliation = &val
	} else {
		if val, ok := kv["affiliation"]; ok {
			if val == nil {
				o.Affiliation = toRepoAffiliationPointer(0)
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["RepoAffiliation"]
				}
				// this is an enum pointer
				if em, ok := val.(string); ok {
					switch em {
					case "organization", "ORGANIZATION":
						o.Affiliation = toRepoAffiliationPointer(0)
					case "user", "USER":
						o.Affiliation = toRepoAffiliationPointer(1)
					case "thirdparty", "THIRDPARTY":
						o.Affiliation = toRepoAffiliationPointer(2)
					}
				}
			}
		}
	}
	if val, ok := kv["default_branch"].(*string); ok {
		o.DefaultBranch = val
	} else if val, ok := kv["default_branch"].(string); ok {
		o.DefaultBranch = &val
	} else {
		if val, ok := kv["default_branch"]; ok {
			if val == nil {
				o.DefaultBranch = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.DefaultBranch = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
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
	if val, ok := kv["language"].(*string); ok {
		o.Language = val
	} else if val, ok := kv["language"].(string); ok {
		o.Language = &val
	} else {
		if val, ok := kv["language"]; ok {
			if val == nil {
				o.Language = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Language = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
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
	if val, ok := kv["visibility"].(*RepoVisibility); ok {
		o.Visibility = val
	} else if val, ok := kv["visibility"].(RepoVisibility); ok {
		o.Visibility = &val
	} else {
		if val, ok := kv["visibility"]; ok {
			if val == nil {
				o.Visibility = toRepoVisibilityPointer(0)
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["RepoVisibility"]
				}
				// this is an enum pointer
				if em, ok := val.(string); ok {
					switch em {
					case "private", "PRIVATE":
						o.Visibility = toRepoVisibilityPointer(0)
					case "public", "PUBLIC":
						o.Visibility = toRepoVisibilityPointer(1)
					}
				}
			}
		}
	}
	o.setDefaults(false)
}
