// DO NOT EDIT -- generated code

// Package agent -
package agent

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

	// WebhookRequestTable is the default table name
	WebhookRequestTable datamodel.ModelNameType = "agent_webhookrequest"

	// WebhookRequestModelName is the model name
	WebhookRequestModelName datamodel.ModelNameType = "agent.WebhookRequest"
)

const (
	// WebhookRequestModelCustomerIDColumn is the column json value customer_id
	WebhookRequestModelCustomerIDColumn = "customer_id"
	// WebhookRequestModelDataColumn is the column json value data
	WebhookRequestModelDataColumn = "data"
	// WebhookRequestModelEncryptedAuthorizationColumn is the column json value encrypted_authorization
	WebhookRequestModelEncryptedAuthorizationColumn = "encrypted_authorization"
	// WebhookRequestModelEventAPIReceivedDateColumn is the column json value event_api_received_date
	WebhookRequestModelEventAPIReceivedDateColumn = "event_api_received_date"
	// WebhookRequestModelEventAPIReceivedDateEpochColumn is the column json value epoch
	WebhookRequestModelEventAPIReceivedDateEpochColumn = "epoch"
	// WebhookRequestModelEventAPIReceivedDateOffsetColumn is the column json value offset
	WebhookRequestModelEventAPIReceivedDateOffsetColumn = "offset"
	// WebhookRequestModelEventAPIReceivedDateRfc3339Column is the column json value rfc3339
	WebhookRequestModelEventAPIReceivedDateRfc3339Column = "rfc3339"
	// WebhookRequestModelHeadersColumn is the column json value headers
	WebhookRequestModelHeadersColumn = "headers"
	// WebhookRequestModelIDColumn is the column json value id
	WebhookRequestModelIDColumn = "id"
	// WebhookRequestModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	WebhookRequestModelIntegrationInstanceIDColumn = "integration_instance_id"
	// WebhookRequestModelIntegrationNameColumn is the column json value integration_name
	WebhookRequestModelIntegrationNameColumn = "integration_name"
	// WebhookRequestModelJobIDColumn is the column json value job_id
	WebhookRequestModelJobIDColumn = "job_id"
	// WebhookRequestModelOperatorReceivedDateColumn is the column json value operator_received_date
	WebhookRequestModelOperatorReceivedDateColumn = "operator_received_date"
	// WebhookRequestModelOperatorReceivedDateEpochColumn is the column json value epoch
	WebhookRequestModelOperatorReceivedDateEpochColumn = "epoch"
	// WebhookRequestModelOperatorReceivedDateOffsetColumn is the column json value offset
	WebhookRequestModelOperatorReceivedDateOffsetColumn = "offset"
	// WebhookRequestModelOperatorReceivedDateRfc3339Column is the column json value rfc3339
	WebhookRequestModelOperatorReceivedDateRfc3339Column = "rfc3339"
	// WebhookRequestModelRefIDColumn is the column json value ref_id
	WebhookRequestModelRefIDColumn = "ref_id"
	// WebhookRequestModelRefTypeColumn is the column json value ref_type
	WebhookRequestModelRefTypeColumn = "ref_type"
	// WebhookRequestModelRequestDateColumn is the column json value request_date
	WebhookRequestModelRequestDateColumn = "request_date"
	// WebhookRequestModelRequestDateEpochColumn is the column json value epoch
	WebhookRequestModelRequestDateEpochColumn = "epoch"
	// WebhookRequestModelRequestDateOffsetColumn is the column json value offset
	WebhookRequestModelRequestDateOffsetColumn = "offset"
	// WebhookRequestModelRequestDateRfc3339Column is the column json value rfc3339
	WebhookRequestModelRequestDateRfc3339Column = "rfc3339"
	// WebhookRequestModelSystemTypeColumn is the column json value system_type
	WebhookRequestModelSystemTypeColumn = "system_type"
	// WebhookRequestModelUUIDColumn is the column json value uuid
	WebhookRequestModelUUIDColumn = "uuid"
)

// WebhookRequestEventAPIReceivedDate represents the object structure for event_api_received_date
type WebhookRequestEventAPIReceivedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toWebhookRequestEventAPIReceivedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *WebhookRequestEventAPIReceivedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *WebhookRequestEventAPIReceivedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toWebhookRequestEventAPIReceivedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toWebhookRequestEventAPIReceivedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toWebhookRequestEventAPIReceivedDateObject(o.Rfc3339, false),
	}
}

func (o *WebhookRequestEventAPIReceivedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *WebhookRequestEventAPIReceivedDate) FromMap(kv map[string]interface{}) {

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

// WebhookRequestOperatorReceivedDate represents the object structure for operator_received_date
type WebhookRequestOperatorReceivedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toWebhookRequestOperatorReceivedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *WebhookRequestOperatorReceivedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *WebhookRequestOperatorReceivedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toWebhookRequestOperatorReceivedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toWebhookRequestOperatorReceivedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toWebhookRequestOperatorReceivedDateObject(o.Rfc3339, false),
	}
}

func (o *WebhookRequestOperatorReceivedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *WebhookRequestOperatorReceivedDate) FromMap(kv map[string]interface{}) {

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

// WebhookRequestRequestDate represents the object structure for request_date
type WebhookRequestRequestDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toWebhookRequestRequestDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *WebhookRequestRequestDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *WebhookRequestRequestDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toWebhookRequestRequestDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toWebhookRequestRequestDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toWebhookRequestRequestDateObject(o.Rfc3339, false),
	}
}

func (o *WebhookRequestRequestDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *WebhookRequestRequestDate) FromMap(kv map[string]interface{}) {

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

// WebhookRequestSystemType is the enumeration type for system_type
type WebhookRequestSystemType int32

// UnmarshalBSONValue for unmarshaling value
func (v *WebhookRequestSystemType) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = WebhookRequestSystemType(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "WORK":
			*v = WebhookRequestSystemType(0)
		case "SOURCECODE":
			*v = WebhookRequestSystemType(1)
		case "CODEQUALITY":
			*v = WebhookRequestSystemType(2)
		case "USER":
			*v = WebhookRequestSystemType(3)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v WebhookRequestSystemType) UnmarshalJSON(buf []byte) error {
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
func (v WebhookRequestSystemType) MarshalJSON() ([]byte, error) {
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
func (v WebhookRequestSystemType) String() string {
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

// FromInterface for decoding from an interface
func (v *WebhookRequestSystemType) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case WebhookRequestSystemType:
		*v = val
	case int32:
		*v = WebhookRequestSystemType(int32(val))
	case int:
		*v = WebhookRequestSystemType(int32(val))
	case string:
		switch val {
		case "WORK":
			*v = WebhookRequestSystemType(0)
		case "SOURCECODE":
			*v = WebhookRequestSystemType(1)
		case "CODEQUALITY":
			*v = WebhookRequestSystemType(2)
		case "USER":
			*v = WebhookRequestSystemType(3)
		}
	}
	return nil
}

const (
	// WebhookRequestSystemTypeWork is the enumeration value for work
	WebhookRequestSystemTypeWork WebhookRequestSystemType = 0
	// WebhookRequestSystemTypeSourcecode is the enumeration value for sourcecode
	WebhookRequestSystemTypeSourcecode WebhookRequestSystemType = 1
	// WebhookRequestSystemTypeCodequality is the enumeration value for codequality
	WebhookRequestSystemTypeCodequality WebhookRequestSystemType = 2
	// WebhookRequestSystemTypeUser is the enumeration value for user
	WebhookRequestSystemTypeUser WebhookRequestSystemType = 3
)

// WebhookRequest webhook data for agent to process with additional auth info
type WebhookRequest struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Data Webhook data
	Data string `json:"data" codec:"data" bson:"data" yaml:"data" faker:"-"`
	// EncryptedAuthorization Auth information encrypted by agent
	EncryptedAuthorization string `json:"encrypted_authorization" codec:"encrypted_authorization" bson:"encrypted_authorization" yaml:"encrypted_authorization" faker:"-"`
	// EventAPIReceivedDate set by event api when it receives webhook data
	EventAPIReceivedDate WebhookRequestEventAPIReceivedDate `json:"event_api_received_date" codec:"event_api_received_date" bson:"event_api_received_date" yaml:"event_api_received_date" faker:"-"`
	// Headers the headers of the incoming webhook
	Headers map[string]string `json:"headers" codec:"headers" bson:"headers" yaml:"headers" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// IntegrationName Name of the integration binary
	IntegrationName string `json:"integration_name" codec:"integration_name" bson:"integration_name" yaml:"integration_name" faker:"-"`
	// JobID ID for this webhook
	JobID string `json:"job_id" codec:"job_id" bson:"job_id" yaml:"job_id" faker:"-"`
	// OperatorReceivedDate set by operator when it receives webhook data
	OperatorReceivedDate WebhookRequestOperatorReceivedDate `json:"operator_received_date" codec:"operator_received_date" bson:"operator_received_date" yaml:"operator_received_date" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RequestDate the date when the request was made
	RequestDate WebhookRequestRequestDate `json:"request_date" codec:"request_date" bson:"request_date" yaml:"request_date" faker:"-"`
	// SystemType The system type of the integration
	SystemType WebhookRequestSystemType `json:"system_type" codec:"system_type" bson:"system_type" yaml:"system_type" faker:"-"`
	// UUID the agent unique identifier
	UUID string `json:"uuid" codec:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*WebhookRequest)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*WebhookRequest)(nil)

func toWebhookRequestObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *WebhookRequest:
		return v.ToMap()

	case WebhookRequestEventAPIReceivedDate:
		return v.ToMap()

	case WebhookRequestOperatorReceivedDate:
		return v.ToMap()

	case WebhookRequestRequestDate:
		return v.ToMap()

	case WebhookRequestSystemType:
		return v.String()

	default:
		return o
	}
}

// String returns a string representation of WebhookRequest
func (o *WebhookRequest) String() string {
	return fmt.Sprintf("agent.WebhookRequest<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *WebhookRequest) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *WebhookRequest) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *WebhookRequest) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *WebhookRequest) GetModelName() datamodel.ModelNameType {
	return WebhookRequestModelName
}

// NewWebhookRequestID provides a template for generating an ID field for WebhookRequest
func NewWebhookRequestID(customerID string, refType string, refID string) string {
	return hash.Values("WebhookRequest", customerID, refType, refID)
}

func (o *WebhookRequest) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("WebhookRequest", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *WebhookRequest) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *WebhookRequest) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *WebhookRequest) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *WebhookRequest) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *WebhookRequest) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *WebhookRequest) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *WebhookRequest) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *WebhookRequest) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *WebhookRequest) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = WebhookRequestModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *WebhookRequest) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *WebhookRequest) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of WebhookRequest
func (o *WebhookRequest) Clone() datamodel.Model {
	c := new(WebhookRequest)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *WebhookRequest) Anon() datamodel.Model {
	c := new(WebhookRequest)
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
func (o *WebhookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *WebhookRequest) UnmarshalJSON(data []byte) error {
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
func (o *WebhookRequest) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *WebhookRequest) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two WebhookRequest objects are equal
func (o *WebhookRequest) IsEqual(other *WebhookRequest) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *WebhookRequest) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"customer_id":             toWebhookRequestObject(o.CustomerID, false),
		"data":                    toWebhookRequestObject(o.Data, false),
		"encrypted_authorization": toWebhookRequestObject(o.EncryptedAuthorization, false),
		"event_api_received_date": toWebhookRequestObject(o.EventAPIReceivedDate, false),
		"headers":                 toWebhookRequestObject(o.Headers, false),
		"id":                      toWebhookRequestObject(o.ID, false),
		"integration_instance_id": toWebhookRequestObject(o.IntegrationInstanceID, true),
		"integration_name":        toWebhookRequestObject(o.IntegrationName, false),
		"job_id":                  toWebhookRequestObject(o.JobID, false),
		"operator_received_date":  toWebhookRequestObject(o.OperatorReceivedDate, false),
		"ref_id":                  toWebhookRequestObject(o.RefID, false),
		"ref_type":                toWebhookRequestObject(o.RefType, false),
		"request_date":            toWebhookRequestObject(o.RequestDate, false),

		"system_type": o.SystemType.String(),
		"uuid":        toWebhookRequestObject(o.UUID, false),
		"hashcode":    toWebhookRequestObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *WebhookRequest) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
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
	if val, ok := kv["encrypted_authorization"].(string); ok {
		o.EncryptedAuthorization = val
	} else {
		if val, ok := kv["encrypted_authorization"]; ok {
			if val == nil {
				o.EncryptedAuthorization = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.EncryptedAuthorization = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["event_api_received_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.EventAPIReceivedDate.FromMap(kv)
		} else if sv, ok := val.(WebhookRequestEventAPIReceivedDate); ok {
			// struct
			o.EventAPIReceivedDate = sv
		} else if sp, ok := val.(*WebhookRequestEventAPIReceivedDate); ok {
			// struct pointer
			o.EventAPIReceivedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.EventAPIReceivedDate.Epoch = dt.Epoch
			o.EventAPIReceivedDate.Rfc3339 = dt.Rfc3339
			o.EventAPIReceivedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.EventAPIReceivedDate.Epoch = dt.Epoch
			o.EventAPIReceivedDate.Rfc3339 = dt.Rfc3339
			o.EventAPIReceivedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.EventAPIReceivedDate.Epoch = dt.Epoch
				o.EventAPIReceivedDate.Rfc3339 = dt.Rfc3339
				o.EventAPIReceivedDate.Offset = dt.Offset
			}
		}
	} else {
		o.EventAPIReceivedDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["headers"]; ok {
		if val != nil {
			kv := make(map[string]string)
			if m, ok := val.(map[string]string); ok {
				kv = m
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					for k, v := range m {
						if mv, ok := v.(string); ok {
							kv[k] = mv
						} else {
							kv[k] = pjson.Stringify(v)
						}
					}
				} else {
					panic("unsupported type for headers field entry: " + reflect.TypeOf(val).String())
				}
			}
			o.Headers = kv
		}
	}
	if o.Headers == nil {
		// here
		o.Headers = make(map[string]string)
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

	if val, ok := kv["operator_received_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.OperatorReceivedDate.FromMap(kv)
		} else if sv, ok := val.(WebhookRequestOperatorReceivedDate); ok {
			// struct
			o.OperatorReceivedDate = sv
		} else if sp, ok := val.(*WebhookRequestOperatorReceivedDate); ok {
			// struct pointer
			o.OperatorReceivedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.OperatorReceivedDate.Epoch = dt.Epoch
			o.OperatorReceivedDate.Rfc3339 = dt.Rfc3339
			o.OperatorReceivedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.OperatorReceivedDate.Epoch = dt.Epoch
			o.OperatorReceivedDate.Rfc3339 = dt.Rfc3339
			o.OperatorReceivedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.OperatorReceivedDate.Epoch = dt.Epoch
				o.OperatorReceivedDate.Rfc3339 = dt.Rfc3339
				o.OperatorReceivedDate.Offset = dt.Offset
			}
		}
	} else {
		o.OperatorReceivedDate.FromMap(map[string]interface{}{})
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
		} else if sv, ok := val.(WebhookRequestRequestDate); ok {
			// struct
			o.RequestDate = sv
		} else if sp, ok := val.(*WebhookRequestRequestDate); ok {
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

	if val, ok := kv["system_type"].(WebhookRequestSystemType); ok {
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
func (o *WebhookRequest) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CustomerID)
	args = append(args, o.Data)
	args = append(args, o.EncryptedAuthorization)
	args = append(args, o.EventAPIReceivedDate)
	args = append(args, o.Headers)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.IntegrationName)
	args = append(args, o.JobID)
	args = append(args, o.OperatorReceivedDate)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RequestDate)
	args = append(args, o.SystemType)
	args = append(args, o.UUID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}
