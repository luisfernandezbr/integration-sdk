// DO NOT EDIT -- generated code

// Package sourcecode - the system which contains source code
package sourcecode

import (
	"encoding/json"
	"fmt"
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

	// PullRequestReviewTable is the default table name
	PullRequestReviewTable datamodel.ModelNameType = "sourcecode_pullrequestreview"

	// PullRequestReviewModelName is the model name
	PullRequestReviewModelName datamodel.ModelNameType = "sourcecode.PullRequestReview"
)

const (
	// PullRequestReviewModelCreatedDateColumn is the column json value created_date
	PullRequestReviewModelCreatedDateColumn = "created_date"
	// PullRequestReviewModelCreatedDateEpochColumn is the column json value epoch
	PullRequestReviewModelCreatedDateEpochColumn = "epoch"
	// PullRequestReviewModelCreatedDateOffsetColumn is the column json value offset
	PullRequestReviewModelCreatedDateOffsetColumn = "offset"
	// PullRequestReviewModelCreatedDateRfc3339Column is the column json value rfc3339
	PullRequestReviewModelCreatedDateRfc3339Column = "rfc3339"
	// PullRequestReviewModelCustomerIDColumn is the column json value customer_id
	PullRequestReviewModelCustomerIDColumn = "customer_id"
	// PullRequestReviewModelIDColumn is the column json value id
	PullRequestReviewModelIDColumn = "id"
	// PullRequestReviewModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	PullRequestReviewModelIntegrationInstanceIDColumn = "integration_instance_id"
	// PullRequestReviewModelPullRequestIDColumn is the column json value pull_request_id
	PullRequestReviewModelPullRequestIDColumn = "pull_request_id"
	// PullRequestReviewModelRefIDColumn is the column json value ref_id
	PullRequestReviewModelRefIDColumn = "ref_id"
	// PullRequestReviewModelRefTypeColumn is the column json value ref_type
	PullRequestReviewModelRefTypeColumn = "ref_type"
	// PullRequestReviewModelRepoIDColumn is the column json value repo_id
	PullRequestReviewModelRepoIDColumn = "repo_id"
	// PullRequestReviewModelStateColumn is the column json value state
	PullRequestReviewModelStateColumn = "state"
	// PullRequestReviewModelURLColumn is the column json value url
	PullRequestReviewModelURLColumn = "url"
	// PullRequestReviewModelUserRefIDColumn is the column json value user_ref_id
	PullRequestReviewModelUserRefIDColumn = "user_ref_id"
)

// PullRequestReviewCreatedDate represents the object structure for created_date
type PullRequestReviewCreatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toPullRequestReviewCreatedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *PullRequestReviewCreatedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *PullRequestReviewCreatedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toPullRequestReviewCreatedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toPullRequestReviewCreatedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toPullRequestReviewCreatedDateObject(o.Rfc3339, false),
	}
}

func (o *PullRequestReviewCreatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *PullRequestReviewCreatedDate) FromMap(kv map[string]interface{}) {

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

// PullRequestReviewState is the enumeration type for state
type PullRequestReviewState int32

// UnmarshalBSONValue for unmarshaling value
func (v *PullRequestReviewState) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = PullRequestReviewState(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "APPROVED":
			*v = PullRequestReviewState(0)
		case "COMMENTED":
			*v = PullRequestReviewState(1)
		case "CHANGES_REQUESTED":
			*v = PullRequestReviewState(2)
		case "PENDING":
			*v = PullRequestReviewState(3)
		case "DISMISSED":
			*v = PullRequestReviewState(4)
		case "REQUESTED":
			*v = PullRequestReviewState(5)
		case "REQUEST_REMOVED":
			*v = PullRequestReviewState(6)
		case "ASSIGNED":
			*v = PullRequestReviewState(7)
		case "UNASSIGNED":
			*v = PullRequestReviewState(8)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v PullRequestReviewState) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "APPROVED":
		v = 0
	case "COMMENTED":
		v = 1
	case "CHANGES_REQUESTED":
		v = 2
	case "PENDING":
		v = 3
	case "DISMISSED":
		v = 4
	case "REQUESTED":
		v = 5
	case "REQUEST_REMOVED":
		v = 6
	case "ASSIGNED":
		v = 7
	case "UNASSIGNED":
		v = 8
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v PullRequestReviewState) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("APPROVED")
	case 1:
		return json.Marshal("COMMENTED")
	case 2:
		return json.Marshal("CHANGES_REQUESTED")
	case 3:
		return json.Marshal("PENDING")
	case 4:
		return json.Marshal("DISMISSED")
	case 5:
		return json.Marshal("REQUESTED")
	case 6:
		return json.Marshal("REQUEST_REMOVED")
	case 7:
		return json.Marshal("ASSIGNED")
	case 8:
		return json.Marshal("UNASSIGNED")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for State
func (v PullRequestReviewState) String() string {
	switch int32(v) {
	case 0:
		return "APPROVED"
	case 1:
		return "COMMENTED"
	case 2:
		return "CHANGES_REQUESTED"
	case 3:
		return "PENDING"
	case 4:
		return "DISMISSED"
	case 5:
		return "REQUESTED"
	case 6:
		return "REQUEST_REMOVED"
	case 7:
		return "ASSIGNED"
	case 8:
		return "UNASSIGNED"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *PullRequestReviewState) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case PullRequestReviewState:
		*v = val
	case int32:
		*v = PullRequestReviewState(int32(val))
	case int:
		*v = PullRequestReviewState(int32(val))
	case string:
		switch val {
		case "APPROVED":
			*v = PullRequestReviewState(0)
		case "COMMENTED":
			*v = PullRequestReviewState(1)
		case "CHANGES_REQUESTED":
			*v = PullRequestReviewState(2)
		case "PENDING":
			*v = PullRequestReviewState(3)
		case "DISMISSED":
			*v = PullRequestReviewState(4)
		case "REQUESTED":
			*v = PullRequestReviewState(5)
		case "REQUEST_REMOVED":
			*v = PullRequestReviewState(6)
		case "ASSIGNED":
			*v = PullRequestReviewState(7)
		case "UNASSIGNED":
			*v = PullRequestReviewState(8)
		}
	}
	return nil
}

const (
	// PullRequestReviewStateApproved is the enumeration value for approved
	PullRequestReviewStateApproved PullRequestReviewState = 0
	// PullRequestReviewStateCommented is the enumeration value for commented
	PullRequestReviewStateCommented PullRequestReviewState = 1
	// PullRequestReviewStateChangesRequested is the enumeration value for changes_requested
	PullRequestReviewStateChangesRequested PullRequestReviewState = 2
	// PullRequestReviewStatePending is the enumeration value for pending
	PullRequestReviewStatePending PullRequestReviewState = 3
	// PullRequestReviewStateDismissed is the enumeration value for dismissed
	PullRequestReviewStateDismissed PullRequestReviewState = 4
	// PullRequestReviewStateRequested is the enumeration value for requested
	PullRequestReviewStateRequested PullRequestReviewState = 5
	// PullRequestReviewStateRequestRemoved is the enumeration value for request_removed
	PullRequestReviewStateRequestRemoved PullRequestReviewState = 6
	// PullRequestReviewStateAssigned is the enumeration value for assigned
	PullRequestReviewStateAssigned PullRequestReviewState = 7
	// PullRequestReviewStateUnassigned is the enumeration value for unassigned
	PullRequestReviewStateUnassigned PullRequestReviewState = 8
)

// PullRequestReview the review for a given pull request
type PullRequestReview struct {
	// CreatedDate the timestamp in UTC that the review was created
	CreatedDate PullRequestReviewCreatedDate `json:"created_date" codec:"created_date" bson:"created_date" yaml:"created_date" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// PullRequestID the pull request this review is associated with
	PullRequestID string `json:"pull_request_id" codec:"pull_request_id" bson:"pull_request_id" yaml:"pull_request_id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RepoID the unique id for the repo
	RepoID string `json:"repo_id" codec:"repo_id" bson:"repo_id" yaml:"repo_id" faker:"-"`
	// State the state of the review
	State PullRequestReviewState `json:"state" codec:"state" bson:"state" yaml:"state" faker:"-"`
	// URL the URL to the source system for this review
	URL string `json:"url" codec:"url" bson:"url" yaml:"url" faker:"url"`
	// UserRefID the user ref_id in the source system
	UserRefID string `json:"user_ref_id" codec:"user_ref_id" bson:"user_ref_id" yaml:"user_ref_id" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*PullRequestReview)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*PullRequestReview)(nil)

func toPullRequestReviewObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *PullRequestReview:
		return v.ToMap()

	case PullRequestReviewCreatedDate:
		return v.ToMap()

	case PullRequestReviewState:
		return v.String()

	default:
		return o
	}
}

// String returns a string representation of PullRequestReview
func (o *PullRequestReview) String() string {
	return fmt.Sprintf("sourcecode.PullRequestReview<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *PullRequestReview) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *PullRequestReview) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *PullRequestReview) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *PullRequestReview) GetModelName() datamodel.ModelNameType {
	return PullRequestReviewModelName
}

// NewPullRequestReviewID provides a template for generating an ID field for PullRequestReview
func NewPullRequestReviewID(customerID string, refID string, refType string, RepoID string) string {
	return hash.Values(customerID, refID, refType, RepoID)
}

func (o *PullRequestReview) setDefaults(frommap bool) {

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
func (o *PullRequestReview) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *PullRequestReview) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *PullRequestReview) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *PullRequestReview) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *PullRequestReview) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *PullRequestReview) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *PullRequestReview) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *PullRequestReview) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *PullRequestReview) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = PullRequestReviewModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *PullRequestReview) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *PullRequestReview) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of PullRequestReview
func (o *PullRequestReview) Clone() datamodel.Model {
	c := new(PullRequestReview)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *PullRequestReview) Anon() datamodel.Model {
	c := new(PullRequestReview)
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
func (o *PullRequestReview) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *PullRequestReview) UnmarshalJSON(data []byte) error {
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
func (o *PullRequestReview) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *PullRequestReview) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two PullRequestReview objects are equal
func (o *PullRequestReview) IsEqual(other *PullRequestReview) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *PullRequestReview) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"created_date":            toPullRequestReviewObject(o.CreatedDate, false),
		"customer_id":             toPullRequestReviewObject(o.CustomerID, false),
		"id":                      toPullRequestReviewObject(o.ID, false),
		"integration_instance_id": toPullRequestReviewObject(o.IntegrationInstanceID, true),
		"pull_request_id":         toPullRequestReviewObject(o.PullRequestID, false),
		"ref_id":                  toPullRequestReviewObject(o.RefID, false),
		"ref_type":                toPullRequestReviewObject(o.RefType, false),
		"repo_id":                 toPullRequestReviewObject(o.RepoID, false),

		"state":       o.State.String(),
		"url":         toPullRequestReviewObject(o.URL, false),
		"user_ref_id": toPullRequestReviewObject(o.UserRefID, false),
		"hashcode":    toPullRequestReviewObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *PullRequestReview) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["created_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.CreatedDate.FromMap(kv)
		} else if sv, ok := val.(PullRequestReviewCreatedDate); ok {
			// struct
			o.CreatedDate = sv
		} else if sp, ok := val.(*PullRequestReviewCreatedDate); ok {
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
	if val, ok := kv["pull_request_id"].(string); ok {
		o.PullRequestID = val
	} else {
		if val, ok := kv["pull_request_id"]; ok {
			if val == nil {
				o.PullRequestID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
	if val, ok := kv["repo_id"].(string); ok {
		o.RepoID = val
	} else {
		if val, ok := kv["repo_id"]; ok {
			if val == nil {
				o.RepoID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.RepoID = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["state"].(PullRequestReviewState); ok {
		o.State = val
	} else {
		if em, ok := kv["state"].(map[string]interface{}); ok {

			ev := em["sourcecode.state"].(string)
			switch ev {
			case "approved", "APPROVED":
				o.State = 0
			case "commented", "COMMENTED":
				o.State = 1
			case "changes_requested", "CHANGES_REQUESTED":
				o.State = 2
			case "pending", "PENDING":
				o.State = 3
			case "dismissed", "DISMISSED":
				o.State = 4
			case "requested", "REQUESTED":
				o.State = 5
			case "request_removed", "REQUEST_REMOVED":
				o.State = 6
			case "assigned", "ASSIGNED":
				o.State = 7
			case "unassigned", "UNASSIGNED":
				o.State = 8
			}
		}
		if em, ok := kv["state"].(string); ok {
			switch em {
			case "approved", "APPROVED":
				o.State = 0
			case "commented", "COMMENTED":
				o.State = 1
			case "changes_requested", "CHANGES_REQUESTED":
				o.State = 2
			case "pending", "PENDING":
				o.State = 3
			case "dismissed", "DISMISSED":
				o.State = 4
			case "requested", "REQUESTED":
				o.State = 5
			case "request_removed", "REQUEST_REMOVED":
				o.State = 6
			case "assigned", "ASSIGNED":
				o.State = 7
			case "unassigned", "UNASSIGNED":
				o.State = 8
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
	if val, ok := kv["user_ref_id"].(string); ok {
		o.UserRefID = val
	} else {
		if val, ok := kv["user_ref_id"]; ok {
			if val == nil {
				o.UserRefID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.UserRefID = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *PullRequestReview) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CreatedDate)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.PullRequestID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RepoID)
	args = append(args, o.State)
	args = append(args, o.URL)
	args = append(args, o.UserRefID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}
