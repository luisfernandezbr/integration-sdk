// DO NOT EDIT -- generated code

// Package agent -
package agent

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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

const (

	// IntegrationMutationRequestTable is the default table name
	IntegrationMutationRequestTable datamodel.ModelNameType = "agent_integrationmutationrequest"

	// IntegrationMutationRequestModelName is the model name
	IntegrationMutationRequestModelName datamodel.ModelNameType = "agent.IntegrationMutationRequest"
)

// IntegrationMutationRequestAction is the enumeration type for action
type IntegrationMutationRequestAction int32

// UnmarshalBSONValue for unmarshaling value
func (v *IntegrationMutationRequestAction) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = IntegrationMutationRequestAction(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "ISSUE_ADD_COMMENT":
			*v = IntegrationMutationRequestAction(0)
		case "ISSUE_SET_TITLE":
			*v = IntegrationMutationRequestAction(1)
		case "ISSUE_SET_STATUS":
			*v = IntegrationMutationRequestAction(2)
		case "ISSUE_SET_PRIORITY":
			*v = IntegrationMutationRequestAction(3)
		case "ISSUE_SET_ASSIGNEE":
			*v = IntegrationMutationRequestAction(4)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v IntegrationMutationRequestAction) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "ISSUE_ADD_COMMENT":
		v = 0
	case "ISSUE_SET_TITLE":
		v = 1
	case "ISSUE_SET_STATUS":
		v = 2
	case "ISSUE_SET_PRIORITY":
		v = 3
	case "ISSUE_SET_ASSIGNEE":
		v = 4
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v IntegrationMutationRequestAction) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("ISSUE_ADD_COMMENT")
	case 1:
		return json.Marshal("ISSUE_SET_TITLE")
	case 2:
		return json.Marshal("ISSUE_SET_STATUS")
	case 3:
		return json.Marshal("ISSUE_SET_PRIORITY")
	case 4:
		return json.Marshal("ISSUE_SET_ASSIGNEE")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for Action
func (v IntegrationMutationRequestAction) String() string {
	switch int32(v) {
	case 0:
		return "ISSUE_ADD_COMMENT"
	case 1:
		return "ISSUE_SET_TITLE"
	case 2:
		return "ISSUE_SET_STATUS"
	case 3:
		return "ISSUE_SET_PRIORITY"
	case 4:
		return "ISSUE_SET_ASSIGNEE"
	}
	return "unset"
}

const (
	// ActionIssueAddComment is the enumeration value for issue_add_comment
	IntegrationMutationRequestActionIssueAddComment IntegrationMutationRequestAction = 0
	// ActionIssueSetTitle is the enumeration value for issue_set_title
	IntegrationMutationRequestActionIssueSetTitle IntegrationMutationRequestAction = 1
	// ActionIssueSetStatus is the enumeration value for issue_set_status
	IntegrationMutationRequestActionIssueSetStatus IntegrationMutationRequestAction = 2
	// ActionIssueSetPriority is the enumeration value for issue_set_priority
	IntegrationMutationRequestActionIssueSetPriority IntegrationMutationRequestAction = 3
	// ActionIssueSetAssignee is the enumeration value for issue_set_assignee
	IntegrationMutationRequestActionIssueSetAssignee IntegrationMutationRequestAction = 4
)

// IntegrationMutationRequestAuthorization represents the object structure for authorization
type IntegrationMutationRequestAuthorization struct {
	// RefreshToken Refresh token
	RefreshToken string `json:"refresh_token" codec:"refresh_token" bson:"refresh_token" yaml:"refresh_token" faker:"-"`
	// URL URL of instance if relevant
	URL *string `json:"url,omitempty" codec:"url,omitempty" bson:"url" yaml:"url,omitempty" faker:"-"`
}

func toIntegrationMutationRequestAuthorizationObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationMutationRequestAuthorization:
		return v.ToMap()

	default:
		return o
	}
}

func (o *IntegrationMutationRequestAuthorization) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// RefreshToken Refresh token
		"refresh_token": toIntegrationMutationRequestAuthorizationObject(o.RefreshToken, false),
		// URL URL of instance if relevant
		"url": toIntegrationMutationRequestAuthorizationObject(o.URL, true),
	}
}

func (o *IntegrationMutationRequestAuthorization) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationMutationRequestAuthorization) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["refresh_token"].(string); ok {
		o.RefreshToken = val
	} else {
		if val, ok := kv["refresh_token"]; ok {
			if val == nil {
				o.RefreshToken = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.RefreshToken = fmt.Sprintf("%v", val)
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

// IntegrationMutationRequestRequestDate represents the object structure for request_date
type IntegrationMutationRequestRequestDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIntegrationMutationRequestRequestDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationMutationRequestRequestDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *IntegrationMutationRequestRequestDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIntegrationMutationRequestRequestDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIntegrationMutationRequestRequestDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIntegrationMutationRequestRequestDateObject(o.Rfc3339, false),
	}
}

func (o *IntegrationMutationRequestRequestDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationMutationRequestRequestDate) FromMap(kv map[string]interface{}) {

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

// IntegrationMutationRequestSystemType is the enumeration type for system_type
type IntegrationMutationRequestSystemType int32

// UnmarshalBSONValue for unmarshaling value
func (v *IntegrationMutationRequestSystemType) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = IntegrationMutationRequestSystemType(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "WORK":
			*v = IntegrationMutationRequestSystemType(0)
		case "SOURCECODE":
			*v = IntegrationMutationRequestSystemType(1)
		case "CODEQUALITY":
			*v = IntegrationMutationRequestSystemType(2)
		case "USER":
			*v = IntegrationMutationRequestSystemType(3)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v IntegrationMutationRequestSystemType) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "WORK":
		v = 0
	case "SOURCECODE":
		v = 1
	case "CODEQUALITY":
		v = 2
	case "USER":
		v = 3
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v IntegrationMutationRequestSystemType) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("WORK")
	case 1:
		return json.Marshal("SOURCECODE")
	case 2:
		return json.Marshal("CODEQUALITY")
	case 3:
		return json.Marshal("USER")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for SystemType
func (v IntegrationMutationRequestSystemType) String() string {
	switch int32(v) {
	case 0:
		return "WORK"
	case 1:
		return "SOURCECODE"
	case 2:
		return "CODEQUALITY"
	case 3:
		return "USER"
	}
	return "unset"
}

const (
	// SystemTypeWork is the enumeration value for work
	IntegrationMutationRequestSystemTypeWork IntegrationMutationRequestSystemType = 0
	// SystemTypeSourcecode is the enumeration value for sourcecode
	IntegrationMutationRequestSystemTypeSourcecode IntegrationMutationRequestSystemType = 1
	// SystemTypeCodequality is the enumeration value for codequality
	IntegrationMutationRequestSystemTypeCodequality IntegrationMutationRequestSystemType = 2
	// SystemTypeUser is the enumeration value for user
	IntegrationMutationRequestSystemTypeUser IntegrationMutationRequestSystemType = 3
)

// IntegrationMutationRequest request to change data in integration going from agent service to agent
type IntegrationMutationRequest struct {
	// Action Action to perform
	Action IntegrationMutationRequestAction `json:"action" codec:"action" bson:"action" yaml:"action" faker:"-"`
	// Authorization Authorization information
	Authorization IntegrationMutationRequestAuthorization `json:"authorization" codec:"authorization" bson:"authorization" yaml:"authorization" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Data Mutation parameters that vary based on action as JSON string.
	Data string `json:"data" codec:"data" bson:"data" yaml:"data" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationName Name of the integration binary
	IntegrationName string `json:"integration_name" codec:"integration_name" bson:"integration_name" yaml:"integration_name" faker:"-"`
	// JobID The job ID
	JobID string `json:"job_id" codec:"job_id" bson:"job_id" yaml:"job_id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RequestDate the date when the request was made
	RequestDate IntegrationMutationRequestRequestDate `json:"request_date" codec:"request_date" bson:"request_date" yaml:"request_date" faker:"-"`
	// SystemType The system type of the integration
	SystemType IntegrationMutationRequestSystemType `json:"system_type" codec:"system_type" bson:"system_type" yaml:"system_type" faker:"-"`
	// UUID the agent unique identifier
	UUID string `json:"uuid" codec:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*IntegrationMutationRequest)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*IntegrationMutationRequest)(nil)

func toIntegrationMutationRequestObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationMutationRequest:
		return v.ToMap()

	case IntegrationMutationRequestAction:
		return v.String()

	case IntegrationMutationRequestAuthorization:
		return v.ToMap()

	case IntegrationMutationRequestRequestDate:
		return v.ToMap()

	case IntegrationMutationRequestSystemType:
		return v.String()

	default:
		return o
	}
}

// String returns a string representation of IntegrationMutationRequest
func (o *IntegrationMutationRequest) String() string {
	return fmt.Sprintf("agent.IntegrationMutationRequest<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *IntegrationMutationRequest) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *IntegrationMutationRequest) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *IntegrationMutationRequest) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *IntegrationMutationRequest) GetModelName() datamodel.ModelNameType {
	return IntegrationMutationRequestModelName
}

// NewIntegrationMutationRequestID provides a template for generating an ID field for IntegrationMutationRequest
func NewIntegrationMutationRequestID(customerID string, refType string, refID string) string {
	return hash.Values("IntegrationMutationRequest", customerID, refType, refID)
}

func (o *IntegrationMutationRequest) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("IntegrationMutationRequest", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *IntegrationMutationRequest) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *IntegrationMutationRequest) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *IntegrationMutationRequest) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *IntegrationMutationRequest) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *IntegrationMutationRequest) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *IntegrationMutationRequest) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *IntegrationMutationRequest) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *IntegrationMutationRequest) IsEvented() bool {
	return false
}

// GetTopicConfig returns the topic config object
func (o *IntegrationMutationRequest) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *IntegrationMutationRequest) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of IntegrationMutationRequest
func (o *IntegrationMutationRequest) Clone() datamodel.Model {
	c := new(IntegrationMutationRequest)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *IntegrationMutationRequest) Anon() datamodel.Model {
	c := new(IntegrationMutationRequest)
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
func (o *IntegrationMutationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *IntegrationMutationRequest) UnmarshalJSON(data []byte) error {
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
func (o *IntegrationMutationRequest) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two IntegrationMutationRequest objects are equal
func (o *IntegrationMutationRequest) IsEqual(other *IntegrationMutationRequest) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *IntegrationMutationRequest) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{

		"action":           o.Action.String(),
		"authorization":    toIntegrationMutationRequestObject(o.Authorization, false),
		"customer_id":      toIntegrationMutationRequestObject(o.CustomerID, false),
		"data":             toIntegrationMutationRequestObject(o.Data, false),
		"id":               toIntegrationMutationRequestObject(o.ID, false),
		"integration_name": toIntegrationMutationRequestObject(o.IntegrationName, false),
		"job_id":           toIntegrationMutationRequestObject(o.JobID, false),
		"ref_id":           toIntegrationMutationRequestObject(o.RefID, false),
		"ref_type":         toIntegrationMutationRequestObject(o.RefType, false),
		"request_date":     toIntegrationMutationRequestObject(o.RequestDate, false),

		"system_type": o.SystemType.String(),
		"uuid":        toIntegrationMutationRequestObject(o.UUID, false),
		"hashcode":    toIntegrationMutationRequestObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationMutationRequest) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["action"].(IntegrationMutationRequestAction); ok {
		o.Action = val
	} else {
		if em, ok := kv["action"].(map[string]interface{}); ok {
			ev := em["agent.action"].(string)
			switch ev {
			case "issue_add_comment", "ISSUE_ADD_COMMENT":
				o.Action = 0
			case "issue_set_title", "ISSUE_SET_TITLE":
				o.Action = 1
			case "issue_set_status", "ISSUE_SET_STATUS":
				o.Action = 2
			case "issue_set_priority", "ISSUE_SET_PRIORITY":
				o.Action = 3
			case "issue_set_assignee", "ISSUE_SET_ASSIGNEE":
				o.Action = 4
			}
		}
		if em, ok := kv["action"].(string); ok {
			switch em {
			case "issue_add_comment", "ISSUE_ADD_COMMENT":
				o.Action = 0
			case "issue_set_title", "ISSUE_SET_TITLE":
				o.Action = 1
			case "issue_set_status", "ISSUE_SET_STATUS":
				o.Action = 2
			case "issue_set_priority", "ISSUE_SET_PRIORITY":
				o.Action = 3
			case "issue_set_assignee", "ISSUE_SET_ASSIGNEE":
				o.Action = 4
			}
		}
	}

	if val, ok := kv["authorization"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.Authorization.FromMap(kv)
		} else if sv, ok := val.(IntegrationMutationRequestAuthorization); ok {
			// struct
			o.Authorization = sv
		} else if sp, ok := val.(*IntegrationMutationRequestAuthorization); ok {
			// struct pointer
			o.Authorization = *sp
		}
	} else {
		o.Authorization.FromMap(map[string]interface{}{})
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

	if val, ok := kv["data"].(string); ok {
		o.Data = val
	} else {
		if val, ok := kv["data"]; ok {
			if val == nil {
				o.Data = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Data = fmt.Sprintf("%v", val)
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

	if val, ok := kv["integration_name"].(string); ok {
		o.IntegrationName = val
	} else {
		if val, ok := kv["integration_name"]; ok {
			if val == nil {
				o.IntegrationName = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.IntegrationName = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["job_id"].(string); ok {
		o.JobID = val
	} else {
		if val, ok := kv["job_id"]; ok {
			if val == nil {
				o.JobID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.JobID = fmt.Sprintf("%v", val)
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

	if val, ok := kv["request_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.RequestDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationMutationRequestRequestDate); ok {
			// struct
			o.RequestDate = sv
		} else if sp, ok := val.(*IntegrationMutationRequestRequestDate); ok {
			// struct pointer
			o.RequestDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.RequestDate.Epoch = dt.Epoch
			o.RequestDate.Rfc3339 = dt.Rfc3339
			o.RequestDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.RequestDate.Epoch = dt.Epoch
			o.RequestDate.Rfc3339 = dt.Rfc3339
			o.RequestDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.RequestDate.Epoch = dt.Epoch
				o.RequestDate.Rfc3339 = dt.Rfc3339
				o.RequestDate.Offset = dt.Offset
			}
		}
	} else {
		o.RequestDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["system_type"].(IntegrationMutationRequestSystemType); ok {
		o.SystemType = val
	} else {
		if em, ok := kv["system_type"].(map[string]interface{}); ok {
			ev := em["agent.system_type"].(string)
			switch ev {
			case "work", "WORK":
				o.SystemType = 0
			case "sourcecode", "SOURCECODE":
				o.SystemType = 1
			case "codequality", "CODEQUALITY":
				o.SystemType = 2
			case "user", "USER":
				o.SystemType = 3
			}
		}
		if em, ok := kv["system_type"].(string); ok {
			switch em {
			case "work", "WORK":
				o.SystemType = 0
			case "sourcecode", "SOURCECODE":
				o.SystemType = 1
			case "codequality", "CODEQUALITY":
				o.SystemType = 2
			case "user", "USER":
				o.SystemType = 3
			}
		}
	}

	if val, ok := kv["uuid"].(string); ok {
		o.UUID = val
	} else {
		if val, ok := kv["uuid"]; ok {
			if val == nil {
				o.UUID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.UUID = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *IntegrationMutationRequest) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Action)
	args = append(args, o.Authorization)
	args = append(args, o.CustomerID)
	args = append(args, o.Data)
	args = append(args, o.ID)
	args = append(args, o.IntegrationName)
	args = append(args, o.JobID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RequestDate)
	args = append(args, o.SystemType)
	args = append(args, o.UUID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *IntegrationMutationRequest) GetEventAPIConfig() datamodel.EventAPIConfig {
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
