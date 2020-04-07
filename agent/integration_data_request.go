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

	// IntegrationDataRequestTable is the default table name
	IntegrationDataRequestTable datamodel.ModelNameType = "agent_integrationdatarequest"

	// IntegrationDataRequestModelName is the model name
	IntegrationDataRequestModelName datamodel.ModelNameType = "agent.IntegrationDataRequest"
)

const (
	// IntegrationDataRequestModelActionColumn is the column json value action
	IntegrationDataRequestModelActionColumn = "action"
	// IntegrationDataRequestModelAuthorizationColumn is the column json value authorization
	IntegrationDataRequestModelAuthorizationColumn = "authorization"
	// IntegrationDataRequestModelAuthorizationRefreshTokenColumn is the column json value refresh_token
	IntegrationDataRequestModelAuthorizationRefreshTokenColumn = "refresh_token"
	// IntegrationDataRequestModelAuthorizationURLColumn is the column json value url
	IntegrationDataRequestModelAuthorizationURLColumn = "url"
	// IntegrationDataRequestModelCustomerIDColumn is the column json value customer_id
	IntegrationDataRequestModelCustomerIDColumn = "customer_id"
	// IntegrationDataRequestModelDataColumn is the column json value data
	IntegrationDataRequestModelDataColumn = "data"
	// IntegrationDataRequestModelIDColumn is the column json value id
	IntegrationDataRequestModelIDColumn = "id"
	// IntegrationDataRequestModelIntegrationNameColumn is the column json value integration_name
	IntegrationDataRequestModelIntegrationNameColumn = "integration_name"
	// IntegrationDataRequestModelJobIDColumn is the column json value job_id
	IntegrationDataRequestModelJobIDColumn = "job_id"
	// IntegrationDataRequestModelRefIDColumn is the column json value ref_id
	IntegrationDataRequestModelRefIDColumn = "ref_id"
	// IntegrationDataRequestModelRefTypeColumn is the column json value ref_type
	IntegrationDataRequestModelRefTypeColumn = "ref_type"
	// IntegrationDataRequestModelRequestDateColumn is the column json value request_date
	IntegrationDataRequestModelRequestDateColumn = "request_date"
	// IntegrationDataRequestModelRequestDateEpochColumn is the column json value epoch
	IntegrationDataRequestModelRequestDateEpochColumn = "epoch"
	// IntegrationDataRequestModelRequestDateOffsetColumn is the column json value offset
	IntegrationDataRequestModelRequestDateOffsetColumn = "offset"
	// IntegrationDataRequestModelRequestDateRfc3339Column is the column json value rfc3339
	IntegrationDataRequestModelRequestDateRfc3339Column = "rfc3339"
	// IntegrationDataRequestModelSystemTypeColumn is the column json value system_type
	IntegrationDataRequestModelSystemTypeColumn = "system_type"
	// IntegrationDataRequestModelUUIDColumn is the column json value uuid
	IntegrationDataRequestModelUUIDColumn = "uuid"
)

// IntegrationDataRequestAction is the enumeration type for action
type IntegrationDataRequestAction int32

// UnmarshalBSONValue for unmarshaling value
func (v *IntegrationDataRequestAction) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = IntegrationDataRequestAction(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "ISSUE_GET_TRANSITIONS":
			*v = IntegrationDataRequestAction(0)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v IntegrationDataRequestAction) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "ISSUE_GET_TRANSITIONS":
		v = 0
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v IntegrationDataRequestAction) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("ISSUE_GET_TRANSITIONS")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for Action
func (v IntegrationDataRequestAction) String() string {
	switch int32(v) {
	case 0:
		return "ISSUE_GET_TRANSITIONS"
	}
	return "unset"
}

const (
	// IntegrationDataRequestActionIssueGetTransitions is the enumeration value for issue_get_transitions
	IntegrationDataRequestActionIssueGetTransitions IntegrationDataRequestAction = 0
)

// IntegrationDataRequestAuthorization represents the object structure for authorization
type IntegrationDataRequestAuthorization struct {
	// RefreshToken Refresh token
	RefreshToken string `json:"refresh_token" codec:"refresh_token" bson:"refresh_token" yaml:"refresh_token" faker:"-"`
	// URL URL of instance if relevant
	URL *string `json:"url,omitempty" codec:"url,omitempty" bson:"url" yaml:"url,omitempty" faker:"-"`
}

func toIntegrationDataRequestAuthorizationObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationDataRequestAuthorization:
		return v.ToMap()

	default:
		return o
	}
}

func (o *IntegrationDataRequestAuthorization) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// RefreshToken Refresh token
		"refresh_token": toIntegrationDataRequestAuthorizationObject(o.RefreshToken, false),
		// URL URL of instance if relevant
		"url": toIntegrationDataRequestAuthorizationObject(o.URL, true),
	}
}

func (o *IntegrationDataRequestAuthorization) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationDataRequestAuthorization) FromMap(kv map[string]interface{}) {

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

// IntegrationDataRequestRequestDate represents the object structure for request_date
type IntegrationDataRequestRequestDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIntegrationDataRequestRequestDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationDataRequestRequestDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *IntegrationDataRequestRequestDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIntegrationDataRequestRequestDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIntegrationDataRequestRequestDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIntegrationDataRequestRequestDateObject(o.Rfc3339, false),
	}
}

func (o *IntegrationDataRequestRequestDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationDataRequestRequestDate) FromMap(kv map[string]interface{}) {

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

// IntegrationDataRequestSystemType is the enumeration type for system_type
type IntegrationDataRequestSystemType int32

// UnmarshalBSONValue for unmarshaling value
func (v *IntegrationDataRequestSystemType) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = IntegrationDataRequestSystemType(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "WORK":
			*v = IntegrationDataRequestSystemType(0)
		case "SOURCECODE":
			*v = IntegrationDataRequestSystemType(1)
		case "CODEQUALITY":
			*v = IntegrationDataRequestSystemType(2)
		case "USER":
			*v = IntegrationDataRequestSystemType(3)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v IntegrationDataRequestSystemType) UnmarshalJSON(buf []byte) error {
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
func (v IntegrationDataRequestSystemType) MarshalJSON() ([]byte, error) {
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
func (v IntegrationDataRequestSystemType) String() string {
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
	// IntegrationDataRequestSystemTypeWork is the enumeration value for work
	IntegrationDataRequestSystemTypeWork IntegrationDataRequestSystemType = 0
	// IntegrationDataRequestSystemTypeSourcecode is the enumeration value for sourcecode
	IntegrationDataRequestSystemTypeSourcecode IntegrationDataRequestSystemType = 1
	// IntegrationDataRequestSystemTypeCodequality is the enumeration value for codequality
	IntegrationDataRequestSystemTypeCodequality IntegrationDataRequestSystemType = 2
	// IntegrationDataRequestSystemTypeUser is the enumeration value for user
	IntegrationDataRequestSystemTypeUser IntegrationDataRequestSystemType = 3
)

// IntegrationDataRequest request to get data from integration for immediate display in the webapp (from agent operator to agent)
type IntegrationDataRequest struct {
	// Action Request action
	Action IntegrationDataRequestAction `json:"action" codec:"action" bson:"action" yaml:"action" faker:"-"`
	// Authorization Authorization information
	Authorization IntegrationDataRequestAuthorization `json:"authorization" codec:"authorization" bson:"authorization" yaml:"authorization" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Data Data request parameters that vary based on action. Stored as JSON string.
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
	RequestDate IntegrationDataRequestRequestDate `json:"request_date" codec:"request_date" bson:"request_date" yaml:"request_date" faker:"-"`
	// SystemType The system type of the integration
	SystemType IntegrationDataRequestSystemType `json:"system_type" codec:"system_type" bson:"system_type" yaml:"system_type" faker:"-"`
	// UUID the agent unique identifier
	UUID string `json:"uuid" codec:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*IntegrationDataRequest)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*IntegrationDataRequest)(nil)

func toIntegrationDataRequestObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationDataRequest:
		return v.ToMap()

	case IntegrationDataRequestAction:
		return v.String()

	case IntegrationDataRequestAuthorization:
		return v.ToMap()

	case IntegrationDataRequestRequestDate:
		return v.ToMap()

	case IntegrationDataRequestSystemType:
		return v.String()

	default:
		return o
	}
}

// String returns a string representation of IntegrationDataRequest
func (o *IntegrationDataRequest) String() string {
	return fmt.Sprintf("agent.IntegrationDataRequest<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *IntegrationDataRequest) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *IntegrationDataRequest) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *IntegrationDataRequest) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *IntegrationDataRequest) GetModelName() datamodel.ModelNameType {
	return IntegrationDataRequestModelName
}

// NewIntegrationDataRequestID provides a template for generating an ID field for IntegrationDataRequest
func NewIntegrationDataRequestID(customerID string, refType string, refID string) string {
	return hash.Values("IntegrationDataRequest", customerID, refType, refID)
}

func (o *IntegrationDataRequest) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("IntegrationDataRequest", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *IntegrationDataRequest) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *IntegrationDataRequest) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *IntegrationDataRequest) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *IntegrationDataRequest) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *IntegrationDataRequest) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *IntegrationDataRequest) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *IntegrationDataRequest) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *IntegrationDataRequest) IsEvented() bool {
	return false
}

// GetTopicConfig returns the topic config object
func (o *IntegrationDataRequest) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *IntegrationDataRequest) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of IntegrationDataRequest
func (o *IntegrationDataRequest) Clone() datamodel.Model {
	c := new(IntegrationDataRequest)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *IntegrationDataRequest) Anon() datamodel.Model {
	c := new(IntegrationDataRequest)
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
func (o *IntegrationDataRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *IntegrationDataRequest) UnmarshalJSON(data []byte) error {
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
func (o *IntegrationDataRequest) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *IntegrationDataRequest) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two IntegrationDataRequest objects are equal
func (o *IntegrationDataRequest) IsEqual(other *IntegrationDataRequest) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *IntegrationDataRequest) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{

		"action":           o.Action.String(),
		"authorization":    toIntegrationDataRequestObject(o.Authorization, false),
		"customer_id":      toIntegrationDataRequestObject(o.CustomerID, false),
		"data":             toIntegrationDataRequestObject(o.Data, false),
		"id":               toIntegrationDataRequestObject(o.ID, false),
		"integration_name": toIntegrationDataRequestObject(o.IntegrationName, false),
		"job_id":           toIntegrationDataRequestObject(o.JobID, false),
		"ref_id":           toIntegrationDataRequestObject(o.RefID, false),
		"ref_type":         toIntegrationDataRequestObject(o.RefType, false),
		"request_date":     toIntegrationDataRequestObject(o.RequestDate, false),

		"system_type": o.SystemType.String(),
		"uuid":        toIntegrationDataRequestObject(o.UUID, false),
		"hashcode":    toIntegrationDataRequestObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationDataRequest) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["action"].(IntegrationDataRequestAction); ok {
		o.Action = val
	} else {
		if em, ok := kv["action"].(map[string]interface{}); ok {
			ev := em["agent.action"].(string)
			switch ev {
			case "issue_get_transitions", "ISSUE_GET_TRANSITIONS":
				o.Action = 0
			}
		}
		if em, ok := kv["action"].(string); ok {
			switch em {
			case "issue_get_transitions", "ISSUE_GET_TRANSITIONS":
				o.Action = 0
			}
		}
	}

	if val, ok := kv["authorization"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.Authorization.FromMap(kv)
		} else if sv, ok := val.(IntegrationDataRequestAuthorization); ok {
			// struct
			o.Authorization = sv
		} else if sp, ok := val.(*IntegrationDataRequestAuthorization); ok {
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
		} else if sv, ok := val.(IntegrationDataRequestRequestDate); ok {
			// struct
			o.RequestDate = sv
		} else if sp, ok := val.(*IntegrationDataRequestRequestDate); ok {
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

	if val, ok := kv["system_type"].(IntegrationDataRequestSystemType); ok {
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
func (o *IntegrationDataRequest) Hash() string {
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
func (o *IntegrationDataRequest) GetEventAPIConfig() datamodel.EventAPIConfig {
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
