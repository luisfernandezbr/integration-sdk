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

	// IntegrationMutationResponseTable is the default table name
	IntegrationMutationResponseTable datamodel.ModelNameType = "agent_integrationmutationresponse"

	// IntegrationMutationResponseModelName is the model name
	IntegrationMutationResponseModelName datamodel.ModelNameType = "agent.IntegrationMutationResponse"
)

const (
	// IntegrationMutationResponseModelAgentReceivedRequestDateColumn is the column json value agent_received_request_date
	IntegrationMutationResponseModelAgentReceivedRequestDateColumn = "agent_received_request_date"
	// IntegrationMutationResponseModelAgentReceivedRequestDateEpochColumn is the column json value epoch
	IntegrationMutationResponseModelAgentReceivedRequestDateEpochColumn = "epoch"
	// IntegrationMutationResponseModelAgentReceivedRequestDateOffsetColumn is the column json value offset
	IntegrationMutationResponseModelAgentReceivedRequestDateOffsetColumn = "offset"
	// IntegrationMutationResponseModelAgentReceivedRequestDateRfc3339Column is the column json value rfc3339
	IntegrationMutationResponseModelAgentReceivedRequestDateRfc3339Column = "rfc3339"
	// IntegrationMutationResponseModelAgentRequestSentDateColumn is the column json value agent_request_sent_date
	IntegrationMutationResponseModelAgentRequestSentDateColumn = "agent_request_sent_date"
	// IntegrationMutationResponseModelAgentRequestSentDateEpochColumn is the column json value epoch
	IntegrationMutationResponseModelAgentRequestSentDateEpochColumn = "epoch"
	// IntegrationMutationResponseModelAgentRequestSentDateOffsetColumn is the column json value offset
	IntegrationMutationResponseModelAgentRequestSentDateOffsetColumn = "offset"
	// IntegrationMutationResponseModelAgentRequestSentDateRfc3339Column is the column json value rfc3339
	IntegrationMutationResponseModelAgentRequestSentDateRfc3339Column = "rfc3339"
	// IntegrationMutationResponseModelAgentResponseSentDateColumn is the column json value agent_response_sent_date
	IntegrationMutationResponseModelAgentResponseSentDateColumn = "agent_response_sent_date"
	// IntegrationMutationResponseModelAgentResponseSentDateEpochColumn is the column json value epoch
	IntegrationMutationResponseModelAgentResponseSentDateEpochColumn = "epoch"
	// IntegrationMutationResponseModelAgentResponseSentDateOffsetColumn is the column json value offset
	IntegrationMutationResponseModelAgentResponseSentDateOffsetColumn = "offset"
	// IntegrationMutationResponseModelAgentResponseSentDateRfc3339Column is the column json value rfc3339
	IntegrationMutationResponseModelAgentResponseSentDateRfc3339Column = "rfc3339"
	// IntegrationMutationResponseModelArchitectureColumn is the column json value architecture
	IntegrationMutationResponseModelArchitectureColumn = "architecture"
	// IntegrationMutationResponseModelCustomerIDColumn is the column json value customer_id
	IntegrationMutationResponseModelCustomerIDColumn = "customer_id"
	// IntegrationMutationResponseModelDataColumn is the column json value data
	IntegrationMutationResponseModelDataColumn = "data"
	// IntegrationMutationResponseModelDistroColumn is the column json value distro
	IntegrationMutationResponseModelDistroColumn = "distro"
	// IntegrationMutationResponseModelErrorColumn is the column json value error
	IntegrationMutationResponseModelErrorColumn = "error"
	// IntegrationMutationResponseModelErrorCodeColumn is the column json value error_code
	IntegrationMutationResponseModelErrorCodeColumn = "error_code"
	// IntegrationMutationResponseModelEventDateColumn is the column json value event_date
	IntegrationMutationResponseModelEventDateColumn = "event_date"
	// IntegrationMutationResponseModelEventDateEpochColumn is the column json value epoch
	IntegrationMutationResponseModelEventDateEpochColumn = "epoch"
	// IntegrationMutationResponseModelEventDateOffsetColumn is the column json value offset
	IntegrationMutationResponseModelEventDateOffsetColumn = "offset"
	// IntegrationMutationResponseModelEventDateRfc3339Column is the column json value rfc3339
	IntegrationMutationResponseModelEventDateRfc3339Column = "rfc3339"
	// IntegrationMutationResponseModelFreeSpaceColumn is the column json value free_space
	IntegrationMutationResponseModelFreeSpaceColumn = "free_space"
	// IntegrationMutationResponseModelGoVersionColumn is the column json value go_version
	IntegrationMutationResponseModelGoVersionColumn = "go_version"
	// IntegrationMutationResponseModelHostnameColumn is the column json value hostname
	IntegrationMutationResponseModelHostnameColumn = "hostname"
	// IntegrationMutationResponseModelIDColumn is the column json value id
	IntegrationMutationResponseModelIDColumn = "id"
	// IntegrationMutationResponseModelJobIDColumn is the column json value job_id
	IntegrationMutationResponseModelJobIDColumn = "job_id"
	// IntegrationMutationResponseModelLastExportDateColumn is the column json value last_export_date
	IntegrationMutationResponseModelLastExportDateColumn = "last_export_date"
	// IntegrationMutationResponseModelLastExportDateEpochColumn is the column json value epoch
	IntegrationMutationResponseModelLastExportDateEpochColumn = "epoch"
	// IntegrationMutationResponseModelLastExportDateOffsetColumn is the column json value offset
	IntegrationMutationResponseModelLastExportDateOffsetColumn = "offset"
	// IntegrationMutationResponseModelLastExportDateRfc3339Column is the column json value rfc3339
	IntegrationMutationResponseModelLastExportDateRfc3339Column = "rfc3339"
	// IntegrationMutationResponseModelMemoryColumn is the column json value memory
	IntegrationMutationResponseModelMemoryColumn = "memory"
	// IntegrationMutationResponseModelMessageColumn is the column json value message
	IntegrationMutationResponseModelMessageColumn = "message"
	// IntegrationMutationResponseModelNumCPUColumn is the column json value num_cpu
	IntegrationMutationResponseModelNumCPUColumn = "num_cpu"
	// IntegrationMutationResponseModelOSColumn is the column json value os
	IntegrationMutationResponseModelOSColumn = "os"
	// IntegrationMutationResponseModelRefIDColumn is the column json value ref_id
	IntegrationMutationResponseModelRefIDColumn = "ref_id"
	// IntegrationMutationResponseModelRefTypeColumn is the column json value ref_type
	IntegrationMutationResponseModelRefTypeColumn = "ref_type"
	// IntegrationMutationResponseModelRequestIDColumn is the column json value request_id
	IntegrationMutationResponseModelRequestIDColumn = "request_id"
	// IntegrationMutationResponseModelSuccessColumn is the column json value success
	IntegrationMutationResponseModelSuccessColumn = "success"
	// IntegrationMutationResponseModelSystemIDColumn is the column json value system_id
	IntegrationMutationResponseModelSystemIDColumn = "system_id"
	// IntegrationMutationResponseModelTypeColumn is the column json value type
	IntegrationMutationResponseModelTypeColumn = "type"
	// IntegrationMutationResponseModelUpdatedObjectsColumn is the column json value updated_objects
	IntegrationMutationResponseModelUpdatedObjectsColumn = "updated_objects"
	// IntegrationMutationResponseModelUptimeColumn is the column json value uptime
	IntegrationMutationResponseModelUptimeColumn = "uptime"
	// IntegrationMutationResponseModelUUIDColumn is the column json value uuid
	IntegrationMutationResponseModelUUIDColumn = "uuid"
	// IntegrationMutationResponseModelVersionColumn is the column json value version
	IntegrationMutationResponseModelVersionColumn = "version"
	// IntegrationMutationResponseModelWebappRequestDateColumn is the column json value webapp_request_date
	IntegrationMutationResponseModelWebappRequestDateColumn = "webapp_request_date"
	// IntegrationMutationResponseModelWebappRequestDateEpochColumn is the column json value epoch
	IntegrationMutationResponseModelWebappRequestDateEpochColumn = "epoch"
	// IntegrationMutationResponseModelWebappRequestDateOffsetColumn is the column json value offset
	IntegrationMutationResponseModelWebappRequestDateOffsetColumn = "offset"
	// IntegrationMutationResponseModelWebappRequestDateRfc3339Column is the column json value rfc3339
	IntegrationMutationResponseModelWebappRequestDateRfc3339Column = "rfc3339"
	// IntegrationMutationResponseModelWebappResponseColumn is the column json value webapp_response
	IntegrationMutationResponseModelWebappResponseColumn = "webapp_response"
)

// IntegrationMutationResponseAgentReceivedRequestDate represents the object structure for agent_received_request_date
type IntegrationMutationResponseAgentReceivedRequestDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIntegrationMutationResponseAgentReceivedRequestDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationMutationResponseAgentReceivedRequestDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *IntegrationMutationResponseAgentReceivedRequestDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIntegrationMutationResponseAgentReceivedRequestDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIntegrationMutationResponseAgentReceivedRequestDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIntegrationMutationResponseAgentReceivedRequestDateObject(o.Rfc3339, false),
	}
}

func (o *IntegrationMutationResponseAgentReceivedRequestDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationMutationResponseAgentReceivedRequestDate) FromMap(kv map[string]interface{}) {

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

// IntegrationMutationResponseAgentRequestSentDate represents the object structure for agent_request_sent_date
type IntegrationMutationResponseAgentRequestSentDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIntegrationMutationResponseAgentRequestSentDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationMutationResponseAgentRequestSentDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *IntegrationMutationResponseAgentRequestSentDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIntegrationMutationResponseAgentRequestSentDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIntegrationMutationResponseAgentRequestSentDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIntegrationMutationResponseAgentRequestSentDateObject(o.Rfc3339, false),
	}
}

func (o *IntegrationMutationResponseAgentRequestSentDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationMutationResponseAgentRequestSentDate) FromMap(kv map[string]interface{}) {

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

// IntegrationMutationResponseAgentResponseSentDate represents the object structure for agent_response_sent_date
type IntegrationMutationResponseAgentResponseSentDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIntegrationMutationResponseAgentResponseSentDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationMutationResponseAgentResponseSentDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *IntegrationMutationResponseAgentResponseSentDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIntegrationMutationResponseAgentResponseSentDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIntegrationMutationResponseAgentResponseSentDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIntegrationMutationResponseAgentResponseSentDateObject(o.Rfc3339, false),
	}
}

func (o *IntegrationMutationResponseAgentResponseSentDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationMutationResponseAgentResponseSentDate) FromMap(kv map[string]interface{}) {

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

// IntegrationMutationResponseErrorCode is the enumeration type for error_code
type IntegrationMutationResponseErrorCode int32

// UnmarshalBSONValue for unmarshaling value
func (v *IntegrationMutationResponseErrorCode) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = IntegrationMutationResponseErrorCode(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "UNKNOWN":
			*v = IntegrationMutationResponseErrorCode(0)
		case "NOT_FOUND":
			*v = IntegrationMutationResponseErrorCode(1)
		case "NOT_AUTHORIZED":
			*v = IntegrationMutationResponseErrorCode(2)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v IntegrationMutationResponseErrorCode) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "UNKNOWN":
		v = 0
	case "NOT_FOUND":
		v = 1
	case "NOT_AUTHORIZED":
		v = 2
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v IntegrationMutationResponseErrorCode) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("UNKNOWN")
	case 1:
		return json.Marshal("NOT_FOUND")
	case 2:
		return json.Marshal("NOT_AUTHORIZED")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for ErrorCode
func (v IntegrationMutationResponseErrorCode) String() string {
	switch int32(v) {
	case 0:
		return "UNKNOWN"
	case 1:
		return "NOT_FOUND"
	case 2:
		return "NOT_AUTHORIZED"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *IntegrationMutationResponseErrorCode) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case IntegrationMutationResponseErrorCode:
		*v = val
	case int32:
		*v = IntegrationMutationResponseErrorCode(int32(val))
	case int:
		*v = IntegrationMutationResponseErrorCode(int32(val))
	case string:
		switch val {
		case "UNKNOWN":
			*v = IntegrationMutationResponseErrorCode(0)
		case "NOT_FOUND":
			*v = IntegrationMutationResponseErrorCode(1)
		case "NOT_AUTHORIZED":
			*v = IntegrationMutationResponseErrorCode(2)
		}
	}
	return nil
}

const (
	// IntegrationMutationResponseErrorCodeUnknown is the enumeration value for unknown
	IntegrationMutationResponseErrorCodeUnknown IntegrationMutationResponseErrorCode = 0
	// IntegrationMutationResponseErrorCodeNotFound is the enumeration value for not_found
	IntegrationMutationResponseErrorCodeNotFound IntegrationMutationResponseErrorCode = 1
	// IntegrationMutationResponseErrorCodeNotAuthorized is the enumeration value for not_authorized
	IntegrationMutationResponseErrorCodeNotAuthorized IntegrationMutationResponseErrorCode = 2
)

// IntegrationMutationResponseEventDate represents the object structure for event_date
type IntegrationMutationResponseEventDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIntegrationMutationResponseEventDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationMutationResponseEventDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *IntegrationMutationResponseEventDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIntegrationMutationResponseEventDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIntegrationMutationResponseEventDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIntegrationMutationResponseEventDateObject(o.Rfc3339, false),
	}
}

func (o *IntegrationMutationResponseEventDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationMutationResponseEventDate) FromMap(kv map[string]interface{}) {

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

// IntegrationMutationResponseLastExportDate represents the object structure for last_export_date
type IntegrationMutationResponseLastExportDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIntegrationMutationResponseLastExportDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationMutationResponseLastExportDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *IntegrationMutationResponseLastExportDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIntegrationMutationResponseLastExportDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIntegrationMutationResponseLastExportDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIntegrationMutationResponseLastExportDateObject(o.Rfc3339, false),
	}
}

func (o *IntegrationMutationResponseLastExportDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationMutationResponseLastExportDate) FromMap(kv map[string]interface{}) {

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

// IntegrationMutationResponseType is the enumeration type for type
type IntegrationMutationResponseType int32

// UnmarshalBSONValue for unmarshaling value
func (v *IntegrationMutationResponseType) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = IntegrationMutationResponseType(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "ENROLL":
			*v = IntegrationMutationResponseType(0)
		case "PING":
			*v = IntegrationMutationResponseType(1)
		case "CRASH":
			*v = IntegrationMutationResponseType(2)
		case "LOG":
			*v = IntegrationMutationResponseType(3)
		case "INTEGRATION":
			*v = IntegrationMutationResponseType(4)
		case "EXPORT":
			*v = IntegrationMutationResponseType(5)
		case "PROJECT":
			*v = IntegrationMutationResponseType(6)
		case "REPO":
			*v = IntegrationMutationResponseType(7)
		case "USER":
			*v = IntegrationMutationResponseType(8)
		case "CALENDAR":
			*v = IntegrationMutationResponseType(9)
		case "UNINSTALL":
			*v = IntegrationMutationResponseType(10)
		case "UPGRADE":
			*v = IntegrationMutationResponseType(11)
		case "START":
			*v = IntegrationMutationResponseType(12)
		case "STOP":
			*v = IntegrationMutationResponseType(13)
		case "PAUSE":
			*v = IntegrationMutationResponseType(14)
		case "RESUME":
			*v = IntegrationMutationResponseType(15)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v IntegrationMutationResponseType) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "ENROLL":
		v = 0
	case "PING":
		v = 1
	case "CRASH":
		v = 2
	case "LOG":
		v = 3
	case "INTEGRATION":
		v = 4
	case "EXPORT":
		v = 5
	case "PROJECT":
		v = 6
	case "REPO":
		v = 7
	case "USER":
		v = 8
	case "CALENDAR":
		v = 9
	case "UNINSTALL":
		v = 10
	case "UPGRADE":
		v = 11
	case "START":
		v = 12
	case "STOP":
		v = 13
	case "PAUSE":
		v = 14
	case "RESUME":
		v = 15
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v IntegrationMutationResponseType) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("ENROLL")
	case 1:
		return json.Marshal("PING")
	case 2:
		return json.Marshal("CRASH")
	case 3:
		return json.Marshal("LOG")
	case 4:
		return json.Marshal("INTEGRATION")
	case 5:
		return json.Marshal("EXPORT")
	case 6:
		return json.Marshal("PROJECT")
	case 7:
		return json.Marshal("REPO")
	case 8:
		return json.Marshal("USER")
	case 9:
		return json.Marshal("CALENDAR")
	case 10:
		return json.Marshal("UNINSTALL")
	case 11:
		return json.Marshal("UPGRADE")
	case 12:
		return json.Marshal("START")
	case 13:
		return json.Marshal("STOP")
	case 14:
		return json.Marshal("PAUSE")
	case 15:
		return json.Marshal("RESUME")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for Type
func (v IntegrationMutationResponseType) String() string {
	switch int32(v) {
	case 0:
		return "ENROLL"
	case 1:
		return "PING"
	case 2:
		return "CRASH"
	case 3:
		return "LOG"
	case 4:
		return "INTEGRATION"
	case 5:
		return "EXPORT"
	case 6:
		return "PROJECT"
	case 7:
		return "REPO"
	case 8:
		return "USER"
	case 9:
		return "CALENDAR"
	case 10:
		return "UNINSTALL"
	case 11:
		return "UPGRADE"
	case 12:
		return "START"
	case 13:
		return "STOP"
	case 14:
		return "PAUSE"
	case 15:
		return "RESUME"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *IntegrationMutationResponseType) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case IntegrationMutationResponseType:
		*v = val
	case int32:
		*v = IntegrationMutationResponseType(int32(val))
	case int:
		*v = IntegrationMutationResponseType(int32(val))
	case string:
		switch val {
		case "ENROLL":
			*v = IntegrationMutationResponseType(0)
		case "PING":
			*v = IntegrationMutationResponseType(1)
		case "CRASH":
			*v = IntegrationMutationResponseType(2)
		case "LOG":
			*v = IntegrationMutationResponseType(3)
		case "INTEGRATION":
			*v = IntegrationMutationResponseType(4)
		case "EXPORT":
			*v = IntegrationMutationResponseType(5)
		case "PROJECT":
			*v = IntegrationMutationResponseType(6)
		case "REPO":
			*v = IntegrationMutationResponseType(7)
		case "USER":
			*v = IntegrationMutationResponseType(8)
		case "CALENDAR":
			*v = IntegrationMutationResponseType(9)
		case "UNINSTALL":
			*v = IntegrationMutationResponseType(10)
		case "UPGRADE":
			*v = IntegrationMutationResponseType(11)
		case "START":
			*v = IntegrationMutationResponseType(12)
		case "STOP":
			*v = IntegrationMutationResponseType(13)
		case "PAUSE":
			*v = IntegrationMutationResponseType(14)
		case "RESUME":
			*v = IntegrationMutationResponseType(15)
		}
	}
	return nil
}

const (
	// IntegrationMutationResponseTypeEnroll is the enumeration value for enroll
	IntegrationMutationResponseTypeEnroll IntegrationMutationResponseType = 0
	// IntegrationMutationResponseTypePing is the enumeration value for ping
	IntegrationMutationResponseTypePing IntegrationMutationResponseType = 1
	// IntegrationMutationResponseTypeCrash is the enumeration value for crash
	IntegrationMutationResponseTypeCrash IntegrationMutationResponseType = 2
	// IntegrationMutationResponseTypeLog is the enumeration value for log
	IntegrationMutationResponseTypeLog IntegrationMutationResponseType = 3
	// IntegrationMutationResponseTypeIntegration is the enumeration value for integration
	IntegrationMutationResponseTypeIntegration IntegrationMutationResponseType = 4
	// IntegrationMutationResponseTypeExport is the enumeration value for export
	IntegrationMutationResponseTypeExport IntegrationMutationResponseType = 5
	// IntegrationMutationResponseTypeProject is the enumeration value for project
	IntegrationMutationResponseTypeProject IntegrationMutationResponseType = 6
	// IntegrationMutationResponseTypeRepo is the enumeration value for repo
	IntegrationMutationResponseTypeRepo IntegrationMutationResponseType = 7
	// IntegrationMutationResponseTypeUser is the enumeration value for user
	IntegrationMutationResponseTypeUser IntegrationMutationResponseType = 8
	// IntegrationMutationResponseTypeCalendar is the enumeration value for calendar
	IntegrationMutationResponseTypeCalendar IntegrationMutationResponseType = 9
	// IntegrationMutationResponseTypeUninstall is the enumeration value for uninstall
	IntegrationMutationResponseTypeUninstall IntegrationMutationResponseType = 10
	// IntegrationMutationResponseTypeUpgrade is the enumeration value for upgrade
	IntegrationMutationResponseTypeUpgrade IntegrationMutationResponseType = 11
	// IntegrationMutationResponseTypeStart is the enumeration value for start
	IntegrationMutationResponseTypeStart IntegrationMutationResponseType = 12
	// IntegrationMutationResponseTypeStop is the enumeration value for stop
	IntegrationMutationResponseTypeStop IntegrationMutationResponseType = 13
	// IntegrationMutationResponseTypePause is the enumeration value for pause
	IntegrationMutationResponseTypePause IntegrationMutationResponseType = 14
	// IntegrationMutationResponseTypeResume is the enumeration value for resume
	IntegrationMutationResponseTypeResume IntegrationMutationResponseType = 15
)

// IntegrationMutationResponseWebappRequestDate represents the object structure for webapp_request_date
type IntegrationMutationResponseWebappRequestDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIntegrationMutationResponseWebappRequestDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationMutationResponseWebappRequestDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *IntegrationMutationResponseWebappRequestDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIntegrationMutationResponseWebappRequestDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIntegrationMutationResponseWebappRequestDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIntegrationMutationResponseWebappRequestDateObject(o.Rfc3339, false),
	}
}

func (o *IntegrationMutationResponseWebappRequestDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationMutationResponseWebappRequestDate) FromMap(kv map[string]interface{}) {

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

// IntegrationMutationResponse response for data mutation going back from agent (used by pipeline and webapp)
type IntegrationMutationResponse struct {
	// AgentReceivedRequestDate set by agent when event is received
	AgentReceivedRequestDate IntegrationMutationResponseAgentReceivedRequestDate `json:"agent_received_request_date" codec:"agent_received_request_date" bson:"agent_received_request_date" yaml:"agent_received_request_date" faker:"-"`
	// AgentRequestSentDate set by agent operator
	AgentRequestSentDate IntegrationMutationResponseAgentRequestSentDate `json:"agent_request_sent_date" codec:"agent_request_sent_date" bson:"agent_request_sent_date" yaml:"agent_request_sent_date" faker:"-"`
	// AgentResponseSentDate set by agent after completing processing
	AgentResponseSentDate IntegrationMutationResponseAgentResponseSentDate `json:"agent_response_sent_date" codec:"agent_response_sent_date" bson:"agent_response_sent_date" yaml:"agent_response_sent_date" faker:"-"`
	// Architecture the architecture of the agent machine
	Architecture string `json:"architecture" codec:"architecture" bson:"architecture" yaml:"architecture" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Data extra data that is specific about this event
	Data *string `json:"data,omitempty" codec:"data,omitempty" bson:"data" yaml:"data,omitempty" faker:"-"`
	// Distro the agent os distribution
	Distro string `json:"distro" codec:"distro" bson:"distro" yaml:"distro" faker:"-"`
	// Error an error message related to this event
	Error *string `json:"error,omitempty" codec:"error,omitempty" bson:"error" yaml:"error,omitempty" faker:"-"`
	// ErrorCode Error code
	ErrorCode IntegrationMutationResponseErrorCode `json:"error_code" codec:"error_code" bson:"error_code" yaml:"error_code" faker:"-"`
	// EventDate the date of the event
	EventDate IntegrationMutationResponseEventDate `json:"event_date" codec:"event_date" bson:"event_date" yaml:"event_date" faker:"-"`
	// FreeSpace the amount of free space in bytes for the agent machine
	FreeSpace int64 `json:"free_space" codec:"free_space" bson:"free_space" yaml:"free_space" faker:"-"`
	// GoVersion the go version that the agent build was built with
	GoVersion string `json:"go_version" codec:"go_version" bson:"go_version" yaml:"go_version" faker:"-"`
	// Hostname the agent hostname
	Hostname string `json:"hostname" codec:"hostname" bson:"hostname" yaml:"hostname" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// JobID The job ID
	JobID string `json:"job_id" codec:"job_id" bson:"job_id" yaml:"job_id" faker:"-"`
	// LastExportDate the last export date
	LastExportDate IntegrationMutationResponseLastExportDate `json:"last_export_date" codec:"last_export_date" bson:"last_export_date" yaml:"last_export_date" faker:"-"`
	// Memory the amount of memory in bytes for the agent machine
	Memory int64 `json:"memory" codec:"memory" bson:"memory" yaml:"memory" faker:"-"`
	// Message a message related to this event
	Message string `json:"message" codec:"message" bson:"message" yaml:"message" faker:"-"`
	// NumCPU the number of CPU the agent is running
	NumCPU int64 `json:"num_cpu" codec:"num_cpu" bson:"num_cpu" yaml:"num_cpu" faker:"-"`
	// OS the agent operating system
	OS string `json:"os" codec:"os" bson:"os" yaml:"os" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RequestID the request id that this response is correlated to
	RequestID string `json:"request_id" codec:"request_id" bson:"request_id" yaml:"request_id" faker:"-"`
	// Success if the response was successful
	Success bool `json:"success" codec:"success" bson:"success" yaml:"success" faker:"-"`
	// SystemID system unique device ID
	SystemID string `json:"system_id" codec:"system_id" bson:"system_id" yaml:"system_id" faker:"-"`
	// Type the type of event
	Type IntegrationMutationResponseType `json:"type" codec:"type" bson:"type" yaml:"type" faker:"-"`
	// UpdatedObjects Updated objects as JSON map[type][]object for use in pipeline
	UpdatedObjects string `json:"updated_objects" codec:"updated_objects" bson:"updated_objects" yaml:"updated_objects" faker:"-"`
	// Uptime the uptime in milliseconds since the agent started
	Uptime int64 `json:"uptime" codec:"uptime" bson:"uptime" yaml:"uptime" faker:"-"`
	// UUID the agent unique identifier
	UUID string `json:"uuid" codec:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
	// Version the agent version
	Version string `json:"version" codec:"version" bson:"version" yaml:"version" faker:"-"`
	// WebappRequestDate set by webapp using browser time
	WebappRequestDate IntegrationMutationResponseWebappRequestDate `json:"webapp_request_date" codec:"webapp_request_date" bson:"webapp_request_date" yaml:"webapp_request_date" faker:"-"`
	// WebappResponse Arbitrary response to be used in webapp
	WebappResponse string `json:"webapp_response" codec:"webapp_response" bson:"webapp_response" yaml:"webapp_response" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*IntegrationMutationResponse)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*IntegrationMutationResponse)(nil)

func toIntegrationMutationResponseObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationMutationResponse:
		return v.ToMap()

	case IntegrationMutationResponseAgentReceivedRequestDate:
		return v.ToMap()

	case IntegrationMutationResponseAgentRequestSentDate:
		return v.ToMap()

	case IntegrationMutationResponseAgentResponseSentDate:
		return v.ToMap()

	case IntegrationMutationResponseErrorCode:
		return v.String()

	case IntegrationMutationResponseEventDate:
		return v.ToMap()

	case IntegrationMutationResponseLastExportDate:
		return v.ToMap()

	case IntegrationMutationResponseType:
		return v.String()

	case IntegrationMutationResponseWebappRequestDate:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of IntegrationMutationResponse
func (o *IntegrationMutationResponse) String() string {
	return fmt.Sprintf("agent.IntegrationMutationResponse<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *IntegrationMutationResponse) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *IntegrationMutationResponse) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *IntegrationMutationResponse) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *IntegrationMutationResponse) GetModelName() datamodel.ModelNameType {
	return IntegrationMutationResponseModelName
}

// NewIntegrationMutationResponseID provides a template for generating an ID field for IntegrationMutationResponse
func NewIntegrationMutationResponseID(customerID string, refType string, refID string) string {
	return hash.Values("IntegrationMutationResponse", customerID, refType, refID)
}

func (o *IntegrationMutationResponse) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("IntegrationMutationResponse", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *IntegrationMutationResponse) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *IntegrationMutationResponse) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *IntegrationMutationResponse) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *IntegrationMutationResponse) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *IntegrationMutationResponse) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *IntegrationMutationResponse) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *IntegrationMutationResponse) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *IntegrationMutationResponse) IsEvented() bool {
	return false
}

// GetTopicConfig returns the topic config object
func (o *IntegrationMutationResponse) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *IntegrationMutationResponse) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of IntegrationMutationResponse
func (o *IntegrationMutationResponse) Clone() datamodel.Model {
	c := new(IntegrationMutationResponse)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *IntegrationMutationResponse) Anon() datamodel.Model {
	c := new(IntegrationMutationResponse)
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
func (o *IntegrationMutationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *IntegrationMutationResponse) UnmarshalJSON(data []byte) error {
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
func (o *IntegrationMutationResponse) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *IntegrationMutationResponse) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two IntegrationMutationResponse objects are equal
func (o *IntegrationMutationResponse) IsEqual(other *IntegrationMutationResponse) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *IntegrationMutationResponse) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"agent_received_request_date": toIntegrationMutationResponseObject(o.AgentReceivedRequestDate, false),
		"agent_request_sent_date":     toIntegrationMutationResponseObject(o.AgentRequestSentDate, false),
		"agent_response_sent_date":    toIntegrationMutationResponseObject(o.AgentResponseSentDate, false),
		"architecture":                toIntegrationMutationResponseObject(o.Architecture, false),
		"customer_id":                 toIntegrationMutationResponseObject(o.CustomerID, false),
		"data":                        toIntegrationMutationResponseObject(o.Data, true),
		"distro":                      toIntegrationMutationResponseObject(o.Distro, false),
		"error":                       toIntegrationMutationResponseObject(o.Error, true),

		"error_code":       o.ErrorCode.String(),
		"event_date":       toIntegrationMutationResponseObject(o.EventDate, false),
		"free_space":       toIntegrationMutationResponseObject(o.FreeSpace, false),
		"go_version":       toIntegrationMutationResponseObject(o.GoVersion, false),
		"hostname":         toIntegrationMutationResponseObject(o.Hostname, false),
		"id":               toIntegrationMutationResponseObject(o.ID, false),
		"job_id":           toIntegrationMutationResponseObject(o.JobID, false),
		"last_export_date": toIntegrationMutationResponseObject(o.LastExportDate, false),
		"memory":           toIntegrationMutationResponseObject(o.Memory, false),
		"message":          toIntegrationMutationResponseObject(o.Message, false),
		"num_cpu":          toIntegrationMutationResponseObject(o.NumCPU, false),
		"os":               toIntegrationMutationResponseObject(o.OS, false),
		"ref_id":           toIntegrationMutationResponseObject(o.RefID, false),
		"ref_type":         toIntegrationMutationResponseObject(o.RefType, false),
		"request_id":       toIntegrationMutationResponseObject(o.RequestID, false),
		"success":          toIntegrationMutationResponseObject(o.Success, false),
		"system_id":        toIntegrationMutationResponseObject(o.SystemID, false),

		"type":                o.Type.String(),
		"updated_objects":     toIntegrationMutationResponseObject(o.UpdatedObjects, false),
		"uptime":              toIntegrationMutationResponseObject(o.Uptime, false),
		"uuid":                toIntegrationMutationResponseObject(o.UUID, false),
		"version":             toIntegrationMutationResponseObject(o.Version, false),
		"webapp_request_date": toIntegrationMutationResponseObject(o.WebappRequestDate, false),
		"webapp_response":     toIntegrationMutationResponseObject(o.WebappResponse, false),
		"hashcode":            toIntegrationMutationResponseObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationMutationResponse) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["agent_received_request_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.AgentReceivedRequestDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationMutationResponseAgentReceivedRequestDate); ok {
			// struct
			o.AgentReceivedRequestDate = sv
		} else if sp, ok := val.(*IntegrationMutationResponseAgentReceivedRequestDate); ok {
			// struct pointer
			o.AgentReceivedRequestDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.AgentReceivedRequestDate.Epoch = dt.Epoch
			o.AgentReceivedRequestDate.Rfc3339 = dt.Rfc3339
			o.AgentReceivedRequestDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.AgentReceivedRequestDate.Epoch = dt.Epoch
			o.AgentReceivedRequestDate.Rfc3339 = dt.Rfc3339
			o.AgentReceivedRequestDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.AgentReceivedRequestDate.Epoch = dt.Epoch
				o.AgentReceivedRequestDate.Rfc3339 = dt.Rfc3339
				o.AgentReceivedRequestDate.Offset = dt.Offset
			}
		}
	} else {
		o.AgentReceivedRequestDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["agent_request_sent_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.AgentRequestSentDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationMutationResponseAgentRequestSentDate); ok {
			// struct
			o.AgentRequestSentDate = sv
		} else if sp, ok := val.(*IntegrationMutationResponseAgentRequestSentDate); ok {
			// struct pointer
			o.AgentRequestSentDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.AgentRequestSentDate.Epoch = dt.Epoch
			o.AgentRequestSentDate.Rfc3339 = dt.Rfc3339
			o.AgentRequestSentDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.AgentRequestSentDate.Epoch = dt.Epoch
			o.AgentRequestSentDate.Rfc3339 = dt.Rfc3339
			o.AgentRequestSentDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.AgentRequestSentDate.Epoch = dt.Epoch
				o.AgentRequestSentDate.Rfc3339 = dt.Rfc3339
				o.AgentRequestSentDate.Offset = dt.Offset
			}
		}
	} else {
		o.AgentRequestSentDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["agent_response_sent_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.AgentResponseSentDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationMutationResponseAgentResponseSentDate); ok {
			// struct
			o.AgentResponseSentDate = sv
		} else if sp, ok := val.(*IntegrationMutationResponseAgentResponseSentDate); ok {
			// struct pointer
			o.AgentResponseSentDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.AgentResponseSentDate.Epoch = dt.Epoch
			o.AgentResponseSentDate.Rfc3339 = dt.Rfc3339
			o.AgentResponseSentDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.AgentResponseSentDate.Epoch = dt.Epoch
			o.AgentResponseSentDate.Rfc3339 = dt.Rfc3339
			o.AgentResponseSentDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.AgentResponseSentDate.Epoch = dt.Epoch
				o.AgentResponseSentDate.Rfc3339 = dt.Rfc3339
				o.AgentResponseSentDate.Offset = dt.Offset
			}
		}
	} else {
		o.AgentResponseSentDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["architecture"].(string); ok {
		o.Architecture = val
	} else {
		if val, ok := kv["architecture"]; ok {
			if val == nil {
				o.Architecture = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Architecture = fmt.Sprintf("%v", val)
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

	if val, ok := kv["data"].(*string); ok {
		o.Data = val
	} else if val, ok := kv["data"].(string); ok {
		o.Data = &val
	} else {
		if val, ok := kv["data"]; ok {
			if val == nil {
				o.Data = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Data = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if val, ok := kv["distro"].(string); ok {
		o.Distro = val
	} else {
		if val, ok := kv["distro"]; ok {
			if val == nil {
				o.Distro = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Distro = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["error"].(*string); ok {
		o.Error = val
	} else if val, ok := kv["error"].(string); ok {
		o.Error = &val
	} else {
		if val, ok := kv["error"]; ok {
			if val == nil {
				o.Error = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Error = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if val, ok := kv["error_code"].(IntegrationMutationResponseErrorCode); ok {
		o.ErrorCode = val
	} else {
		if em, ok := kv["error_code"].(map[string]interface{}); ok {

			ev := em["agent.error_code"].(string)
			switch ev {
			case "unknown", "UNKNOWN":
				o.ErrorCode = 0
			case "not_found", "NOT_FOUND":
				o.ErrorCode = 1
			case "not_authorized", "NOT_AUTHORIZED":
				o.ErrorCode = 2
			}
		}
		if em, ok := kv["error_code"].(string); ok {
			switch em {
			case "unknown", "UNKNOWN":
				o.ErrorCode = 0
			case "not_found", "NOT_FOUND":
				o.ErrorCode = 1
			case "not_authorized", "NOT_AUTHORIZED":
				o.ErrorCode = 2
			}
		}
	}

	if val, ok := kv["event_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.EventDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationMutationResponseEventDate); ok {
			// struct
			o.EventDate = sv
		} else if sp, ok := val.(*IntegrationMutationResponseEventDate); ok {
			// struct pointer
			o.EventDate = *sp
		}
	} else {
		o.EventDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["free_space"].(int64); ok {
		o.FreeSpace = val
	} else {
		if val, ok := kv["free_space"]; ok {
			if val == nil {
				o.FreeSpace = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.FreeSpace = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["go_version"].(string); ok {
		o.GoVersion = val
	} else {
		if val, ok := kv["go_version"]; ok {
			if val == nil {
				o.GoVersion = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.GoVersion = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["hostname"].(string); ok {
		o.Hostname = val
	} else {
		if val, ok := kv["hostname"]; ok {
			if val == nil {
				o.Hostname = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Hostname = fmt.Sprintf("%v", val)
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

	if val, ok := kv["last_export_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastExportDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationMutationResponseLastExportDate); ok {
			// struct
			o.LastExportDate = sv
		} else if sp, ok := val.(*IntegrationMutationResponseLastExportDate); ok {
			// struct pointer
			o.LastExportDate = *sp
		}
	} else {
		o.LastExportDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["memory"].(int64); ok {
		o.Memory = val
	} else {
		if val, ok := kv["memory"]; ok {
			if val == nil {
				o.Memory = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Memory = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["message"].(string); ok {
		o.Message = val
	} else {
		if val, ok := kv["message"]; ok {
			if val == nil {
				o.Message = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Message = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["num_cpu"].(int64); ok {
		o.NumCPU = val
	} else {
		if val, ok := kv["num_cpu"]; ok {
			if val == nil {
				o.NumCPU = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.NumCPU = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["os"].(string); ok {
		o.OS = val
	} else {
		if val, ok := kv["os"]; ok {
			if val == nil {
				o.OS = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.OS = fmt.Sprintf("%v", val)
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

	if val, ok := kv["request_id"].(string); ok {
		o.RequestID = val
	} else {
		if val, ok := kv["request_id"]; ok {
			if val == nil {
				o.RequestID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.RequestID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["success"].(bool); ok {
		o.Success = val
	} else {
		if val, ok := kv["success"]; ok {
			if val == nil {
				o.Success = false
			} else {
				o.Success = number.ToBoolAny(val)
			}
		}
	}

	if val, ok := kv["system_id"].(string); ok {
		o.SystemID = val
	} else {
		if val, ok := kv["system_id"]; ok {
			if val == nil {
				o.SystemID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.SystemID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["type"].(IntegrationMutationResponseType); ok {
		o.Type = val
	} else {
		if em, ok := kv["type"].(map[string]interface{}); ok {

			ev := em["agent.type"].(string)
			switch ev {
			case "enroll", "ENROLL":
				o.Type = 0
			case "ping", "PING":
				o.Type = 1
			case "crash", "CRASH":
				o.Type = 2
			case "log", "LOG":
				o.Type = 3
			case "integration", "INTEGRATION":
				o.Type = 4
			case "export", "EXPORT":
				o.Type = 5
			case "project", "PROJECT":
				o.Type = 6
			case "repo", "REPO":
				o.Type = 7
			case "user", "USER":
				o.Type = 8
			case "calendar", "CALENDAR":
				o.Type = 9
			case "uninstall", "UNINSTALL":
				o.Type = 10
			case "upgrade", "UPGRADE":
				o.Type = 11
			case "start", "START":
				o.Type = 12
			case "stop", "STOP":
				o.Type = 13
			case "pause", "PAUSE":
				o.Type = 14
			case "resume", "RESUME":
				o.Type = 15
			}
		}
		if em, ok := kv["type"].(string); ok {
			switch em {
			case "enroll", "ENROLL":
				o.Type = 0
			case "ping", "PING":
				o.Type = 1
			case "crash", "CRASH":
				o.Type = 2
			case "log", "LOG":
				o.Type = 3
			case "integration", "INTEGRATION":
				o.Type = 4
			case "export", "EXPORT":
				o.Type = 5
			case "project", "PROJECT":
				o.Type = 6
			case "repo", "REPO":
				o.Type = 7
			case "user", "USER":
				o.Type = 8
			case "calendar", "CALENDAR":
				o.Type = 9
			case "uninstall", "UNINSTALL":
				o.Type = 10
			case "upgrade", "UPGRADE":
				o.Type = 11
			case "start", "START":
				o.Type = 12
			case "stop", "STOP":
				o.Type = 13
			case "pause", "PAUSE":
				o.Type = 14
			case "resume", "RESUME":
				o.Type = 15
			}
		}
	}

	if val, ok := kv["updated_objects"].(string); ok {
		o.UpdatedObjects = val
	} else {
		if val, ok := kv["updated_objects"]; ok {
			if val == nil {
				o.UpdatedObjects = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.UpdatedObjects = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["uptime"].(int64); ok {
		o.Uptime = val
	} else {
		if val, ok := kv["uptime"]; ok {
			if val == nil {
				o.Uptime = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Uptime = number.ToInt64Any(val)
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

	if val, ok := kv["version"].(string); ok {
		o.Version = val
	} else {
		if val, ok := kv["version"]; ok {
			if val == nil {
				o.Version = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Version = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["webapp_request_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.WebappRequestDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationMutationResponseWebappRequestDate); ok {
			// struct
			o.WebappRequestDate = sv
		} else if sp, ok := val.(*IntegrationMutationResponseWebappRequestDate); ok {
			// struct pointer
			o.WebappRequestDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.WebappRequestDate.Epoch = dt.Epoch
			o.WebappRequestDate.Rfc3339 = dt.Rfc3339
			o.WebappRequestDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.WebappRequestDate.Epoch = dt.Epoch
			o.WebappRequestDate.Rfc3339 = dt.Rfc3339
			o.WebappRequestDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.WebappRequestDate.Epoch = dt.Epoch
				o.WebappRequestDate.Rfc3339 = dt.Rfc3339
				o.WebappRequestDate.Offset = dt.Offset
			}
		}
	} else {
		o.WebappRequestDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["webapp_response"].(string); ok {
		o.WebappResponse = val
	} else {
		if val, ok := kv["webapp_response"]; ok {
			if val == nil {
				o.WebappResponse = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.WebappResponse = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *IntegrationMutationResponse) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.AgentReceivedRequestDate)
	args = append(args, o.AgentRequestSentDate)
	args = append(args, o.AgentResponseSentDate)
	args = append(args, o.Architecture)
	args = append(args, o.CustomerID)
	args = append(args, o.Data)
	args = append(args, o.Distro)
	args = append(args, o.Error)
	args = append(args, o.ErrorCode)
	args = append(args, o.EventDate)
	args = append(args, o.FreeSpace)
	args = append(args, o.GoVersion)
	args = append(args, o.Hostname)
	args = append(args, o.ID)
	args = append(args, o.JobID)
	args = append(args, o.LastExportDate)
	args = append(args, o.Memory)
	args = append(args, o.Message)
	args = append(args, o.NumCPU)
	args = append(args, o.OS)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RequestID)
	args = append(args, o.Success)
	args = append(args, o.SystemID)
	args = append(args, o.Type)
	args = append(args, o.UpdatedObjects)
	args = append(args, o.Uptime)
	args = append(args, o.UUID)
	args = append(args, o.Version)
	args = append(args, o.WebappRequestDate)
	args = append(args, o.WebappResponse)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *IntegrationMutationResponse) GetEventAPIConfig() datamodel.EventAPIConfig {
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
