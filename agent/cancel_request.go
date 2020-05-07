// DO NOT EDIT -- generated code

// Package agent - the agent communicates with pinpoint cloud to send integration data
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

	// CancelRequestTable is the default table name
	CancelRequestTable datamodel.ModelNameType = "agent_cancelrequest"

	// CancelRequestModelName is the model name
	CancelRequestModelName datamodel.ModelNameType = "agent.CancelRequest"
)

const (
	// CancelRequestModelCommandColumn is the column json value command
	CancelRequestModelCommandColumn = "command"
	// CancelRequestModelCustomerIDColumn is the column json value customer_id
	CancelRequestModelCustomerIDColumn = "customer_id"
	// CancelRequestModelIDColumn is the column json value id
	CancelRequestModelIDColumn = "id"
	// CancelRequestModelRefIDColumn is the column json value ref_id
	CancelRequestModelRefIDColumn = "ref_id"
	// CancelRequestModelRefTypeColumn is the column json value ref_type
	CancelRequestModelRefTypeColumn = "ref_type"
	// CancelRequestModelRequestDateColumn is the column json value request_date
	CancelRequestModelRequestDateColumn = "request_date"
	// CancelRequestModelRequestDateEpochColumn is the column json value epoch
	CancelRequestModelRequestDateEpochColumn = "epoch"
	// CancelRequestModelRequestDateOffsetColumn is the column json value offset
	CancelRequestModelRequestDateOffsetColumn = "offset"
	// CancelRequestModelRequestDateRfc3339Column is the column json value rfc3339
	CancelRequestModelRequestDateRfc3339Column = "rfc3339"
	// CancelRequestModelUUIDColumn is the column json value uuid
	CancelRequestModelUUIDColumn = "uuid"
)

// CancelRequestCommand is the enumeration type for command
type CancelRequestCommand int32

// UnmarshalBSONValue for unmarshaling value
func (v *CancelRequestCommand) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = CancelRequestCommand(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "EXPORT":
			*v = CancelRequestCommand(0)
		case "ONBOARD":
			*v = CancelRequestCommand(1)
		case "INTEGRATION":
			*v = CancelRequestCommand(2)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v CancelRequestCommand) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "EXPORT":
		v = 0
	case "ONBOARD":
		v = 1
	case "INTEGRATION":
		v = 2
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v CancelRequestCommand) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("EXPORT")
	case 1:
		return json.Marshal("ONBOARD")
	case 2:
		return json.Marshal("INTEGRATION")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for Command
func (v CancelRequestCommand) String() string {
	switch int32(v) {
	case 0:
		return "EXPORT"
	case 1:
		return "ONBOARD"
	case 2:
		return "INTEGRATION"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *CancelRequestCommand) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case CancelRequestCommand:
		*v = val
	case int32:
		*v = CancelRequestCommand(int32(val))
	case int:
		*v = CancelRequestCommand(int32(val))
	case string:
		switch val {
		case "EXPORT":
			*v = CancelRequestCommand(0)
		case "ONBOARD":
			*v = CancelRequestCommand(1)
		case "INTEGRATION":
			*v = CancelRequestCommand(2)
		}
	}
	return nil
}

const (
	// CancelRequestCommandExport is the enumeration value for export
	CancelRequestCommandExport CancelRequestCommand = 0
	// CancelRequestCommandOnboard is the enumeration value for onboard
	CancelRequestCommandOnboard CancelRequestCommand = 1
	// CancelRequestCommandIntegration is the enumeration value for integration
	CancelRequestCommandIntegration CancelRequestCommand = 2
)

// CancelRequestRequestDate represents the object structure for request_date
type CancelRequestRequestDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toCancelRequestRequestDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CancelRequestRequestDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *CancelRequestRequestDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toCancelRequestRequestDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toCancelRequestRequestDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toCancelRequestRequestDateObject(o.Rfc3339, false),
	}
}

func (o *CancelRequestRequestDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *CancelRequestRequestDate) FromMap(kv map[string]interface{}) {

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

// CancelRequest an agent action to request a cancel
type CancelRequest struct {
	// Command The command to cancel
	Command CancelRequestCommand `json:"command" codec:"command" bson:"command" yaml:"command" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RequestDate the date when the request was made
	RequestDate CancelRequestRequestDate `json:"request_date" codec:"request_date" bson:"request_date" yaml:"request_date" faker:"-"`
	// UUID the agent unique identifier
	UUID string `json:"uuid" codec:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*CancelRequest)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*CancelRequest)(nil)

func toCancelRequestObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CancelRequest:
		return v.ToMap()

	case CancelRequestCommand:
		return v.String()

	case CancelRequestRequestDate:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of CancelRequest
func (o *CancelRequest) String() string {
	return fmt.Sprintf("agent.CancelRequest<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *CancelRequest) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *CancelRequest) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *CancelRequest) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *CancelRequest) GetModelName() datamodel.ModelNameType {
	return CancelRequestModelName
}

// NewCancelRequestID provides a template for generating an ID field for CancelRequest
func NewCancelRequestID(customerID string, refType string, refID string) string {
	return hash.Values("CancelRequest", customerID, refType, refID)
}

func (o *CancelRequest) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("CancelRequest", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *CancelRequest) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *CancelRequest) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *CancelRequest) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *CancelRequest) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *CancelRequest) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *CancelRequest) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *CancelRequest) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *CancelRequest) IsEvented() bool {
	return false
}

// GetTopicConfig returns the topic config object
func (o *CancelRequest) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *CancelRequest) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of CancelRequest
func (o *CancelRequest) Clone() datamodel.Model {
	c := new(CancelRequest)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *CancelRequest) Anon() datamodel.Model {
	c := new(CancelRequest)
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
func (o *CancelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *CancelRequest) UnmarshalJSON(data []byte) error {
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
func (o *CancelRequest) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *CancelRequest) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two CancelRequest objects are equal
func (o *CancelRequest) IsEqual(other *CancelRequest) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *CancelRequest) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{

		"command":      o.Command.String(),
		"customer_id":  toCancelRequestObject(o.CustomerID, false),
		"id":           toCancelRequestObject(o.ID, false),
		"ref_id":       toCancelRequestObject(o.RefID, false),
		"ref_type":     toCancelRequestObject(o.RefType, false),
		"request_date": toCancelRequestObject(o.RequestDate, false),
		"uuid":         toCancelRequestObject(o.UUID, false),
		"hashcode":     toCancelRequestObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *CancelRequest) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["command"].(CancelRequestCommand); ok {
		o.Command = val
	} else {
		if em, ok := kv["command"].(map[string]interface{}); ok {

			ev := em["agent.command"].(string)
			switch ev {
			case "export", "EXPORT":
				o.Command = 0
			case "onboard", "ONBOARD":
				o.Command = 1
			case "integration", "INTEGRATION":
				o.Command = 2
			}
		}
		if em, ok := kv["command"].(string); ok {
			switch em {
			case "export", "EXPORT":
				o.Command = 0
			case "onboard", "ONBOARD":
				o.Command = 1
			case "integration", "INTEGRATION":
				o.Command = 2
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
		} else if sv, ok := val.(CancelRequestRequestDate); ok {
			// struct
			o.RequestDate = sv
		} else if sp, ok := val.(*CancelRequestRequestDate); ok {
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
func (o *CancelRequest) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Command)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RequestDate)
	args = append(args, o.UUID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *CancelRequest) GetEventAPIConfig() datamodel.EventAPIConfig {
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
