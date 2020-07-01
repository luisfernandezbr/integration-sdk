// DO NOT EDIT -- generated code

// Package agent -
package agent

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

	// WebhookResponseTable is the default table name
	WebhookResponseTable datamodel.ModelNameType = "agent_webhookresponse"

	// WebhookResponseModelName is the model name
	WebhookResponseModelName datamodel.ModelNameType = "agent.WebhookResponse"
)

const (
	// WebhookResponseModelAgentReceivedDateColumn is the column json value agent_received_date
	WebhookResponseModelAgentReceivedDateColumn = "agent_received_date"
	// WebhookResponseModelAgentReceivedDateEpochColumn is the column json value epoch
	WebhookResponseModelAgentReceivedDateEpochColumn = "epoch"
	// WebhookResponseModelAgentReceivedDateOffsetColumn is the column json value offset
	WebhookResponseModelAgentReceivedDateOffsetColumn = "offset"
	// WebhookResponseModelAgentReceivedDateRfc3339Column is the column json value rfc3339
	WebhookResponseModelAgentReceivedDateRfc3339Column = "rfc3339"
	// WebhookResponseModelAgentResponseSentDateColumn is the column json value agent_response_sent_date
	WebhookResponseModelAgentResponseSentDateColumn = "agent_response_sent_date"
	// WebhookResponseModelAgentResponseSentDateEpochColumn is the column json value epoch
	WebhookResponseModelAgentResponseSentDateEpochColumn = "epoch"
	// WebhookResponseModelAgentResponseSentDateOffsetColumn is the column json value offset
	WebhookResponseModelAgentResponseSentDateOffsetColumn = "offset"
	// WebhookResponseModelAgentResponseSentDateRfc3339Column is the column json value rfc3339
	WebhookResponseModelAgentResponseSentDateRfc3339Column = "rfc3339"
	// WebhookResponseModelArchitectureColumn is the column json value architecture
	WebhookResponseModelArchitectureColumn = "architecture"
	// WebhookResponseModelCustomerIDColumn is the column json value customer_id
	WebhookResponseModelCustomerIDColumn = "customer_id"
	// WebhookResponseModelDataColumn is the column json value data
	WebhookResponseModelDataColumn = "data"
	// WebhookResponseModelDistroColumn is the column json value distro
	WebhookResponseModelDistroColumn = "distro"
	// WebhookResponseModelErrorColumn is the column json value error
	WebhookResponseModelErrorColumn = "error"
	// WebhookResponseModelEventAPIReceivedDateColumn is the column json value event_api_received_date
	WebhookResponseModelEventAPIReceivedDateColumn = "event_api_received_date"
	// WebhookResponseModelEventAPIReceivedDateEpochColumn is the column json value epoch
	WebhookResponseModelEventAPIReceivedDateEpochColumn = "epoch"
	// WebhookResponseModelEventAPIReceivedDateOffsetColumn is the column json value offset
	WebhookResponseModelEventAPIReceivedDateOffsetColumn = "offset"
	// WebhookResponseModelEventAPIReceivedDateRfc3339Column is the column json value rfc3339
	WebhookResponseModelEventAPIReceivedDateRfc3339Column = "rfc3339"
	// WebhookResponseModelEventDateColumn is the column json value event_date
	WebhookResponseModelEventDateColumn = "event_date"
	// WebhookResponseModelEventDateEpochColumn is the column json value epoch
	WebhookResponseModelEventDateEpochColumn = "epoch"
	// WebhookResponseModelEventDateOffsetColumn is the column json value offset
	WebhookResponseModelEventDateOffsetColumn = "offset"
	// WebhookResponseModelEventDateRfc3339Column is the column json value rfc3339
	WebhookResponseModelEventDateRfc3339Column = "rfc3339"
	// WebhookResponseModelFreeSpaceColumn is the column json value free_space
	WebhookResponseModelFreeSpaceColumn = "free_space"
	// WebhookResponseModelGoVersionColumn is the column json value go_version
	WebhookResponseModelGoVersionColumn = "go_version"
	// WebhookResponseModelHostnameColumn is the column json value hostname
	WebhookResponseModelHostnameColumn = "hostname"
	// WebhookResponseModelIDColumn is the column json value id
	WebhookResponseModelIDColumn = "id"
	// WebhookResponseModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	WebhookResponseModelIntegrationInstanceIDColumn = "integration_instance_id"
	// WebhookResponseModelJobIDColumn is the column json value job_id
	WebhookResponseModelJobIDColumn = "job_id"
	// WebhookResponseModelLastExportDateColumn is the column json value last_export_date
	WebhookResponseModelLastExportDateColumn = "last_export_date"
	// WebhookResponseModelLastExportDateEpochColumn is the column json value epoch
	WebhookResponseModelLastExportDateEpochColumn = "epoch"
	// WebhookResponseModelLastExportDateOffsetColumn is the column json value offset
	WebhookResponseModelLastExportDateOffsetColumn = "offset"
	// WebhookResponseModelLastExportDateRfc3339Column is the column json value rfc3339
	WebhookResponseModelLastExportDateRfc3339Column = "rfc3339"
	// WebhookResponseModelMemoryColumn is the column json value memory
	WebhookResponseModelMemoryColumn = "memory"
	// WebhookResponseModelMessageColumn is the column json value message
	WebhookResponseModelMessageColumn = "message"
	// WebhookResponseModelNumCPUColumn is the column json value num_cpu
	WebhookResponseModelNumCPUColumn = "num_cpu"
	// WebhookResponseModelOperatorReceivedDateColumn is the column json value operator_received_date
	WebhookResponseModelOperatorReceivedDateColumn = "operator_received_date"
	// WebhookResponseModelOperatorReceivedDateEpochColumn is the column json value epoch
	WebhookResponseModelOperatorReceivedDateEpochColumn = "epoch"
	// WebhookResponseModelOperatorReceivedDateOffsetColumn is the column json value offset
	WebhookResponseModelOperatorReceivedDateOffsetColumn = "offset"
	// WebhookResponseModelOperatorReceivedDateRfc3339Column is the column json value rfc3339
	WebhookResponseModelOperatorReceivedDateRfc3339Column = "rfc3339"
	// WebhookResponseModelOSColumn is the column json value os
	WebhookResponseModelOSColumn = "os"
	// WebhookResponseModelRefIDColumn is the column json value ref_id
	WebhookResponseModelRefIDColumn = "ref_id"
	// WebhookResponseModelRefTypeColumn is the column json value ref_type
	WebhookResponseModelRefTypeColumn = "ref_type"
	// WebhookResponseModelRequestIDColumn is the column json value request_id
	WebhookResponseModelRequestIDColumn = "request_id"
	// WebhookResponseModelSuccessColumn is the column json value success
	WebhookResponseModelSuccessColumn = "success"
	// WebhookResponseModelSystemIDColumn is the column json value system_id
	WebhookResponseModelSystemIDColumn = "system_id"
	// WebhookResponseModelTypeColumn is the column json value type
	WebhookResponseModelTypeColumn = "type"
	// WebhookResponseModelUpdatedObjectsColumn is the column json value updated_objects
	WebhookResponseModelUpdatedObjectsColumn = "updated_objects"
	// WebhookResponseModelUptimeColumn is the column json value uptime
	WebhookResponseModelUptimeColumn = "uptime"
	// WebhookResponseModelUUIDColumn is the column json value uuid
	WebhookResponseModelUUIDColumn = "uuid"
	// WebhookResponseModelVersionColumn is the column json value version
	WebhookResponseModelVersionColumn = "version"
)

// WebhookResponseAgentReceivedDate represents the object structure for agent_received_date
type WebhookResponseAgentReceivedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toWebhookResponseAgentReceivedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *WebhookResponseAgentReceivedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *WebhookResponseAgentReceivedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toWebhookResponseAgentReceivedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toWebhookResponseAgentReceivedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toWebhookResponseAgentReceivedDateObject(o.Rfc3339, false),
	}
}

func (o *WebhookResponseAgentReceivedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *WebhookResponseAgentReceivedDate) FromMap(kv map[string]interface{}) {

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

// WebhookResponseAgentResponseSentDate represents the object structure for agent_response_sent_date
type WebhookResponseAgentResponseSentDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toWebhookResponseAgentResponseSentDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *WebhookResponseAgentResponseSentDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *WebhookResponseAgentResponseSentDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toWebhookResponseAgentResponseSentDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toWebhookResponseAgentResponseSentDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toWebhookResponseAgentResponseSentDateObject(o.Rfc3339, false),
	}
}

func (o *WebhookResponseAgentResponseSentDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *WebhookResponseAgentResponseSentDate) FromMap(kv map[string]interface{}) {

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

// WebhookResponseEventAPIReceivedDate represents the object structure for event_api_received_date
type WebhookResponseEventAPIReceivedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toWebhookResponseEventAPIReceivedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *WebhookResponseEventAPIReceivedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *WebhookResponseEventAPIReceivedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toWebhookResponseEventAPIReceivedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toWebhookResponseEventAPIReceivedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toWebhookResponseEventAPIReceivedDateObject(o.Rfc3339, false),
	}
}

func (o *WebhookResponseEventAPIReceivedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *WebhookResponseEventAPIReceivedDate) FromMap(kv map[string]interface{}) {

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

// WebhookResponseEventDate represents the object structure for event_date
type WebhookResponseEventDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toWebhookResponseEventDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *WebhookResponseEventDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *WebhookResponseEventDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toWebhookResponseEventDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toWebhookResponseEventDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toWebhookResponseEventDateObject(o.Rfc3339, false),
	}
}

func (o *WebhookResponseEventDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *WebhookResponseEventDate) FromMap(kv map[string]interface{}) {

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

// WebhookResponseLastExportDate represents the object structure for last_export_date
type WebhookResponseLastExportDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toWebhookResponseLastExportDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *WebhookResponseLastExportDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *WebhookResponseLastExportDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toWebhookResponseLastExportDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toWebhookResponseLastExportDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toWebhookResponseLastExportDateObject(o.Rfc3339, false),
	}
}

func (o *WebhookResponseLastExportDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *WebhookResponseLastExportDate) FromMap(kv map[string]interface{}) {

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

// WebhookResponseOperatorReceivedDate represents the object structure for operator_received_date
type WebhookResponseOperatorReceivedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toWebhookResponseOperatorReceivedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *WebhookResponseOperatorReceivedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *WebhookResponseOperatorReceivedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toWebhookResponseOperatorReceivedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toWebhookResponseOperatorReceivedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toWebhookResponseOperatorReceivedDateObject(o.Rfc3339, false),
	}
}

func (o *WebhookResponseOperatorReceivedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *WebhookResponseOperatorReceivedDate) FromMap(kv map[string]interface{}) {

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

// WebhookResponseType is the enumeration type for type
type WebhookResponseType int32

// UnmarshalBSONValue for unmarshaling value
func (v *WebhookResponseType) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = WebhookResponseType(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "ENROLL":
			*v = WebhookResponseType(0)
		case "PING":
			*v = WebhookResponseType(1)
		case "CRASH":
			*v = WebhookResponseType(2)
		case "LOG":
			*v = WebhookResponseType(3)
		case "INTEGRATION":
			*v = WebhookResponseType(4)
		case "EXPORT":
			*v = WebhookResponseType(5)
		case "PROJECT":
			*v = WebhookResponseType(6)
		case "REPO":
			*v = WebhookResponseType(7)
		case "USER":
			*v = WebhookResponseType(8)
		case "CALENDAR":
			*v = WebhookResponseType(9)
		case "UNINSTALL":
			*v = WebhookResponseType(10)
		case "UPGRADE":
			*v = WebhookResponseType(11)
		case "START":
			*v = WebhookResponseType(12)
		case "STOP":
			*v = WebhookResponseType(13)
		case "PAUSE":
			*v = WebhookResponseType(14)
		case "RESUME":
			*v = WebhookResponseType(15)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v WebhookResponseType) UnmarshalJSON(buf []byte) error {
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
func (v WebhookResponseType) MarshalJSON() ([]byte, error) {
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
func (v WebhookResponseType) String() string {
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
func (v *WebhookResponseType) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case WebhookResponseType:
		*v = val
	case int32:
		*v = WebhookResponseType(int32(val))
	case int:
		*v = WebhookResponseType(int32(val))
	case string:
		switch val {
		case "ENROLL":
			*v = WebhookResponseType(0)
		case "PING":
			*v = WebhookResponseType(1)
		case "CRASH":
			*v = WebhookResponseType(2)
		case "LOG":
			*v = WebhookResponseType(3)
		case "INTEGRATION":
			*v = WebhookResponseType(4)
		case "EXPORT":
			*v = WebhookResponseType(5)
		case "PROJECT":
			*v = WebhookResponseType(6)
		case "REPO":
			*v = WebhookResponseType(7)
		case "USER":
			*v = WebhookResponseType(8)
		case "CALENDAR":
			*v = WebhookResponseType(9)
		case "UNINSTALL":
			*v = WebhookResponseType(10)
		case "UPGRADE":
			*v = WebhookResponseType(11)
		case "START":
			*v = WebhookResponseType(12)
		case "STOP":
			*v = WebhookResponseType(13)
		case "PAUSE":
			*v = WebhookResponseType(14)
		case "RESUME":
			*v = WebhookResponseType(15)
		}
	}
	return nil
}

const (
	// WebhookResponseTypeEnroll is the enumeration value for enroll
	WebhookResponseTypeEnroll WebhookResponseType = 0
	// WebhookResponseTypePing is the enumeration value for ping
	WebhookResponseTypePing WebhookResponseType = 1
	// WebhookResponseTypeCrash is the enumeration value for crash
	WebhookResponseTypeCrash WebhookResponseType = 2
	// WebhookResponseTypeLog is the enumeration value for log
	WebhookResponseTypeLog WebhookResponseType = 3
	// WebhookResponseTypeIntegration is the enumeration value for integration
	WebhookResponseTypeIntegration WebhookResponseType = 4
	// WebhookResponseTypeExport is the enumeration value for export
	WebhookResponseTypeExport WebhookResponseType = 5
	// WebhookResponseTypeProject is the enumeration value for project
	WebhookResponseTypeProject WebhookResponseType = 6
	// WebhookResponseTypeRepo is the enumeration value for repo
	WebhookResponseTypeRepo WebhookResponseType = 7
	// WebhookResponseTypeUser is the enumeration value for user
	WebhookResponseTypeUser WebhookResponseType = 8
	// WebhookResponseTypeCalendar is the enumeration value for calendar
	WebhookResponseTypeCalendar WebhookResponseType = 9
	// WebhookResponseTypeUninstall is the enumeration value for uninstall
	WebhookResponseTypeUninstall WebhookResponseType = 10
	// WebhookResponseTypeUpgrade is the enumeration value for upgrade
	WebhookResponseTypeUpgrade WebhookResponseType = 11
	// WebhookResponseTypeStart is the enumeration value for start
	WebhookResponseTypeStart WebhookResponseType = 12
	// WebhookResponseTypeStop is the enumeration value for stop
	WebhookResponseTypeStop WebhookResponseType = 13
	// WebhookResponseTypePause is the enumeration value for pause
	WebhookResponseTypePause WebhookResponseType = 14
	// WebhookResponseTypeResume is the enumeration value for resume
	WebhookResponseTypeResume WebhookResponseType = 15
)

// WebhookResponse response for data mutation going back from agent (used by pipeline and webapp)
type WebhookResponse struct {
	// AgentReceivedDate set by agent when event is received
	AgentReceivedDate WebhookResponseAgentReceivedDate `json:"agent_received_date" codec:"agent_received_date" bson:"agent_received_date" yaml:"agent_received_date" faker:"-"`
	// AgentResponseSentDate set by agent after completing processing
	AgentResponseSentDate WebhookResponseAgentResponseSentDate `json:"agent_response_sent_date" codec:"agent_response_sent_date" bson:"agent_response_sent_date" yaml:"agent_response_sent_date" faker:"-"`
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
	// EventAPIReceivedDate set by event api when it receives webhook data
	EventAPIReceivedDate WebhookResponseEventAPIReceivedDate `json:"event_api_received_date" codec:"event_api_received_date" bson:"event_api_received_date" yaml:"event_api_received_date" faker:"-"`
	// EventDate the date of the event
	EventDate WebhookResponseEventDate `json:"event_date" codec:"event_date" bson:"event_date" yaml:"event_date" faker:"-"`
	// FreeSpace the amount of free space in bytes for the agent machine
	FreeSpace int64 `json:"free_space" codec:"free_space" bson:"free_space" yaml:"free_space" faker:"-"`
	// GoVersion the go version that the agent build was built with
	GoVersion string `json:"go_version" codec:"go_version" bson:"go_version" yaml:"go_version" faker:"-"`
	// Hostname the agent hostname
	Hostname string `json:"hostname" codec:"hostname" bson:"hostname" yaml:"hostname" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// JobID The job ID
	JobID string `json:"job_id" codec:"job_id" bson:"job_id" yaml:"job_id" faker:"-"`
	// LastExportDate the last export date
	LastExportDate WebhookResponseLastExportDate `json:"last_export_date" codec:"last_export_date" bson:"last_export_date" yaml:"last_export_date" faker:"-"`
	// Memory the amount of memory in bytes for the agent machine
	Memory int64 `json:"memory" codec:"memory" bson:"memory" yaml:"memory" faker:"-"`
	// Message a message related to this event
	Message string `json:"message" codec:"message" bson:"message" yaml:"message" faker:"-"`
	// NumCPU the number of CPU the agent is running
	NumCPU int64 `json:"num_cpu" codec:"num_cpu" bson:"num_cpu" yaml:"num_cpu" faker:"-"`
	// OperatorReceivedDate set by operator when it receives webhook data
	OperatorReceivedDate WebhookResponseOperatorReceivedDate `json:"operator_received_date" codec:"operator_received_date" bson:"operator_received_date" yaml:"operator_received_date" faker:"-"`
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
	Type WebhookResponseType `json:"type" codec:"type" bson:"type" yaml:"type" faker:"-"`
	// UpdatedObjects Updated objects as JSON map[type][]object for use in pipeline
	UpdatedObjects string `json:"updated_objects" codec:"updated_objects" bson:"updated_objects" yaml:"updated_objects" faker:"-"`
	// Uptime the uptime in milliseconds since the agent started
	Uptime int64 `json:"uptime" codec:"uptime" bson:"uptime" yaml:"uptime" faker:"-"`
	// UUID the agent unique identifier
	UUID string `json:"uuid" codec:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
	// Version the agent version
	Version string `json:"version" codec:"version" bson:"version" yaml:"version" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*WebhookResponse)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*WebhookResponse)(nil)

func toWebhookResponseObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *WebhookResponse:
		return v.ToMap()

	case WebhookResponseAgentReceivedDate:
		return v.ToMap()

	case WebhookResponseAgentResponseSentDate:
		return v.ToMap()

	case WebhookResponseEventAPIReceivedDate:
		return v.ToMap()

	case WebhookResponseEventDate:
		return v.ToMap()

	case WebhookResponseLastExportDate:
		return v.ToMap()

	case WebhookResponseOperatorReceivedDate:
		return v.ToMap()

	case WebhookResponseType:
		return v.String()

	default:
		return o
	}
}

// String returns a string representation of WebhookResponse
func (o *WebhookResponse) String() string {
	return fmt.Sprintf("agent.WebhookResponse<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *WebhookResponse) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *WebhookResponse) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *WebhookResponse) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *WebhookResponse) GetModelName() datamodel.ModelNameType {
	return WebhookResponseModelName
}

// NewWebhookResponseID provides a template for generating an ID field for WebhookResponse
func NewWebhookResponseID(customerID string, refType string, refID string) string {
	return hash.Values("WebhookResponse", customerID, refType, refID)
}

func (o *WebhookResponse) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("WebhookResponse", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *WebhookResponse) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *WebhookResponse) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *WebhookResponse) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *WebhookResponse) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *WebhookResponse) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *WebhookResponse) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *WebhookResponse) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *WebhookResponse) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *WebhookResponse) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = WebhookResponseModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *WebhookResponse) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *WebhookResponse) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of WebhookResponse
func (o *WebhookResponse) Clone() datamodel.Model {
	c := new(WebhookResponse)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *WebhookResponse) Anon() datamodel.Model {
	c := new(WebhookResponse)
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
func (o *WebhookResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *WebhookResponse) UnmarshalJSON(data []byte) error {
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
func (o *WebhookResponse) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *WebhookResponse) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two WebhookResponse objects are equal
func (o *WebhookResponse) IsEqual(other *WebhookResponse) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *WebhookResponse) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"agent_received_date":      toWebhookResponseObject(o.AgentReceivedDate, false),
		"agent_response_sent_date": toWebhookResponseObject(o.AgentResponseSentDate, false),
		"architecture":             toWebhookResponseObject(o.Architecture, false),
		"customer_id":              toWebhookResponseObject(o.CustomerID, false),
		"data":                     toWebhookResponseObject(o.Data, true),
		"distro":                   toWebhookResponseObject(o.Distro, false),
		"error":                    toWebhookResponseObject(o.Error, true),
		"event_api_received_date":  toWebhookResponseObject(o.EventAPIReceivedDate, false),
		"event_date":               toWebhookResponseObject(o.EventDate, false),
		"free_space":               toWebhookResponseObject(o.FreeSpace, false),
		"go_version":               toWebhookResponseObject(o.GoVersion, false),
		"hostname":                 toWebhookResponseObject(o.Hostname, false),
		"id":                       toWebhookResponseObject(o.ID, false),
		"integration_instance_id":  toWebhookResponseObject(o.IntegrationInstanceID, true),
		"job_id":                   toWebhookResponseObject(o.JobID, false),
		"last_export_date":         toWebhookResponseObject(o.LastExportDate, false),
		"memory":                   toWebhookResponseObject(o.Memory, false),
		"message":                  toWebhookResponseObject(o.Message, false),
		"num_cpu":                  toWebhookResponseObject(o.NumCPU, false),
		"operator_received_date":   toWebhookResponseObject(o.OperatorReceivedDate, false),
		"os":                       toWebhookResponseObject(o.OS, false),
		"ref_id":                   toWebhookResponseObject(o.RefID, false),
		"ref_type":                 toWebhookResponseObject(o.RefType, false),
		"request_id":               toWebhookResponseObject(o.RequestID, false),
		"success":                  toWebhookResponseObject(o.Success, false),
		"system_id":                toWebhookResponseObject(o.SystemID, false),

		"type":            o.Type.String(),
		"updated_objects": toWebhookResponseObject(o.UpdatedObjects, false),
		"uptime":          toWebhookResponseObject(o.Uptime, false),
		"uuid":            toWebhookResponseObject(o.UUID, false),
		"version":         toWebhookResponseObject(o.Version, false),
		"hashcode":        toWebhookResponseObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *WebhookResponse) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["agent_received_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.AgentReceivedDate.FromMap(kv)
		} else if sv, ok := val.(WebhookResponseAgentReceivedDate); ok {
			// struct
			o.AgentReceivedDate = sv
		} else if sp, ok := val.(*WebhookResponseAgentReceivedDate); ok {
			// struct pointer
			o.AgentReceivedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.AgentReceivedDate.Epoch = dt.Epoch
			o.AgentReceivedDate.Rfc3339 = dt.Rfc3339
			o.AgentReceivedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.AgentReceivedDate.Epoch = dt.Epoch
			o.AgentReceivedDate.Rfc3339 = dt.Rfc3339
			o.AgentReceivedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.AgentReceivedDate.Epoch = dt.Epoch
				o.AgentReceivedDate.Rfc3339 = dt.Rfc3339
				o.AgentReceivedDate.Offset = dt.Offset
			}
		}
	} else {
		o.AgentReceivedDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["agent_response_sent_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.AgentResponseSentDate.FromMap(kv)
		} else if sv, ok := val.(WebhookResponseAgentResponseSentDate); ok {
			// struct
			o.AgentResponseSentDate = sv
		} else if sp, ok := val.(*WebhookResponseAgentResponseSentDate); ok {
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

	if val, ok := kv["event_api_received_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.EventAPIReceivedDate.FromMap(kv)
		} else if sv, ok := val.(WebhookResponseEventAPIReceivedDate); ok {
			// struct
			o.EventAPIReceivedDate = sv
		} else if sp, ok := val.(*WebhookResponseEventAPIReceivedDate); ok {
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

	if val, ok := kv["event_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.EventDate.FromMap(kv)
		} else if sv, ok := val.(WebhookResponseEventDate); ok {
			// struct
			o.EventDate = sv
		} else if sp, ok := val.(*WebhookResponseEventDate); ok {
			// struct pointer
			o.EventDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.EventDate.Epoch = dt.Epoch
			o.EventDate.Rfc3339 = dt.Rfc3339
			o.EventDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.EventDate.Epoch = dt.Epoch
			o.EventDate.Rfc3339 = dt.Rfc3339
			o.EventDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.EventDate.Epoch = dt.Epoch
				o.EventDate.Rfc3339 = dt.Rfc3339
				o.EventDate.Offset = dt.Offset
			}
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
		} else if sv, ok := val.(WebhookResponseLastExportDate); ok {
			// struct
			o.LastExportDate = sv
		} else if sp, ok := val.(*WebhookResponseLastExportDate); ok {
			// struct pointer
			o.LastExportDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.LastExportDate.Epoch = dt.Epoch
			o.LastExportDate.Rfc3339 = dt.Rfc3339
			o.LastExportDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.LastExportDate.Epoch = dt.Epoch
			o.LastExportDate.Rfc3339 = dt.Rfc3339
			o.LastExportDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.LastExportDate.Epoch = dt.Epoch
				o.LastExportDate.Rfc3339 = dt.Rfc3339
				o.LastExportDate.Offset = dt.Offset
			}
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

	if val, ok := kv["operator_received_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.OperatorReceivedDate.FromMap(kv)
		} else if sv, ok := val.(WebhookResponseOperatorReceivedDate); ok {
			// struct
			o.OperatorReceivedDate = sv
		} else if sp, ok := val.(*WebhookResponseOperatorReceivedDate); ok {
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
	if val, ok := kv["type"].(WebhookResponseType); ok {
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
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *WebhookResponse) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.AgentReceivedDate)
	args = append(args, o.AgentResponseSentDate)
	args = append(args, o.Architecture)
	args = append(args, o.CustomerID)
	args = append(args, o.Data)
	args = append(args, o.Distro)
	args = append(args, o.Error)
	args = append(args, o.EventAPIReceivedDate)
	args = append(args, o.EventDate)
	args = append(args, o.FreeSpace)
	args = append(args, o.GoVersion)
	args = append(args, o.Hostname)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.JobID)
	args = append(args, o.LastExportDate)
	args = append(args, o.Memory)
	args = append(args, o.Message)
	args = append(args, o.NumCPU)
	args = append(args, o.OperatorReceivedDate)
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
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}
