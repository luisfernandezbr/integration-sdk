// DO NOT EDIT -- generated code

// Package agent - the agent communicates with pinpoint cloud to send integration data
package agent

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/bxcodec/faker"
	"github.com/pinpt/go-common/v10/datamodel"
	"github.com/pinpt/go-common/v10/datetime"
	"github.com/pinpt/go-common/v10/graphql"
	"github.com/pinpt/go-common/v10/hash"
	pjson "github.com/pinpt/go-common/v10/json"
	"github.com/pinpt/go-common/v10/number"
	pstrings "github.com/pinpt/go-common/v10/strings"
)

const (

	// IntegrationInstanceStatTable is the default table name
	IntegrationInstanceStatTable datamodel.ModelNameType = "agent_integrationinstancestat"

	// IntegrationInstanceStatModelName is the model name
	IntegrationInstanceStatModelName datamodel.ModelNameType = "agent.IntegrationInstanceStat"
)

const (
	// IntegrationInstanceStatModelCreatedAtColumn is the column json value created_ts
	IntegrationInstanceStatModelCreatedAtColumn = "created_ts"
	// IntegrationInstanceStatModelCustomerIDColumn is the column json value customer_id
	IntegrationInstanceStatModelCustomerIDColumn = "customer_id"
	// IntegrationInstanceStatModelIDColumn is the column json value id
	IntegrationInstanceStatModelIDColumn = "id"
	// IntegrationInstanceStatModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	IntegrationInstanceStatModelIntegrationInstanceIDColumn = "integration_instance_id"
	// IntegrationInstanceStatModelLastExportCompletedDateColumn is the column json value last_export_completed_date
	IntegrationInstanceStatModelLastExportCompletedDateColumn = "last_export_completed_date"
	// IntegrationInstanceStatModelLastExportCompletedDateEpochColumn is the column json value epoch
	IntegrationInstanceStatModelLastExportCompletedDateEpochColumn = "epoch"
	// IntegrationInstanceStatModelLastExportCompletedDateOffsetColumn is the column json value offset
	IntegrationInstanceStatModelLastExportCompletedDateOffsetColumn = "offset"
	// IntegrationInstanceStatModelLastExportCompletedDateRfc3339Column is the column json value rfc3339
	IntegrationInstanceStatModelLastExportCompletedDateRfc3339Column = "rfc3339"
	// IntegrationInstanceStatModelLastExportRequestedDateColumn is the column json value last_export_requested_date
	IntegrationInstanceStatModelLastExportRequestedDateColumn = "last_export_requested_date"
	// IntegrationInstanceStatModelLastExportRequestedDateEpochColumn is the column json value epoch
	IntegrationInstanceStatModelLastExportRequestedDateEpochColumn = "epoch"
	// IntegrationInstanceStatModelLastExportRequestedDateOffsetColumn is the column json value offset
	IntegrationInstanceStatModelLastExportRequestedDateOffsetColumn = "offset"
	// IntegrationInstanceStatModelLastExportRequestedDateRfc3339Column is the column json value rfc3339
	IntegrationInstanceStatModelLastExportRequestedDateRfc3339Column = "rfc3339"
	// IntegrationInstanceStatModelLastHistoricalCompletedDateColumn is the column json value last_historical_completed_date
	IntegrationInstanceStatModelLastHistoricalCompletedDateColumn = "last_historical_completed_date"
	// IntegrationInstanceStatModelLastHistoricalCompletedDateEpochColumn is the column json value epoch
	IntegrationInstanceStatModelLastHistoricalCompletedDateEpochColumn = "epoch"
	// IntegrationInstanceStatModelLastHistoricalCompletedDateOffsetColumn is the column json value offset
	IntegrationInstanceStatModelLastHistoricalCompletedDateOffsetColumn = "offset"
	// IntegrationInstanceStatModelLastHistoricalCompletedDateRfc3339Column is the column json value rfc3339
	IntegrationInstanceStatModelLastHistoricalCompletedDateRfc3339Column = "rfc3339"
	// IntegrationInstanceStatModelLastHistoricalRequestedDateColumn is the column json value last_historical_requested_date
	IntegrationInstanceStatModelLastHistoricalRequestedDateColumn = "last_historical_requested_date"
	// IntegrationInstanceStatModelLastHistoricalRequestedDateEpochColumn is the column json value epoch
	IntegrationInstanceStatModelLastHistoricalRequestedDateEpochColumn = "epoch"
	// IntegrationInstanceStatModelLastHistoricalRequestedDateOffsetColumn is the column json value offset
	IntegrationInstanceStatModelLastHistoricalRequestedDateOffsetColumn = "offset"
	// IntegrationInstanceStatModelLastHistoricalRequestedDateRfc3339Column is the column json value rfc3339
	IntegrationInstanceStatModelLastHistoricalRequestedDateRfc3339Column = "rfc3339"
	// IntegrationInstanceStatModelLastProcessingDateColumn is the column json value last_processing_date
	IntegrationInstanceStatModelLastProcessingDateColumn = "last_processing_date"
	// IntegrationInstanceStatModelLastProcessingDateEpochColumn is the column json value epoch
	IntegrationInstanceStatModelLastProcessingDateEpochColumn = "epoch"
	// IntegrationInstanceStatModelLastProcessingDateOffsetColumn is the column json value offset
	IntegrationInstanceStatModelLastProcessingDateOffsetColumn = "offset"
	// IntegrationInstanceStatModelLastProcessingDateRfc3339Column is the column json value rfc3339
	IntegrationInstanceStatModelLastProcessingDateRfc3339Column = "rfc3339"
	// IntegrationInstanceStatModelRefIDColumn is the column json value ref_id
	IntegrationInstanceStatModelRefIDColumn = "ref_id"
	// IntegrationInstanceStatModelRefTypeColumn is the column json value ref_type
	IntegrationInstanceStatModelRefTypeColumn = "ref_type"
	// IntegrationInstanceStatModelUpdatedAtColumn is the column json value updated_ts
	IntegrationInstanceStatModelUpdatedAtColumn = "updated_ts"
)

// IntegrationInstanceStatLastExportCompletedDate represents the object structure for last_export_completed_date
type IntegrationInstanceStatLastExportCompletedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIntegrationInstanceStatLastExportCompletedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationInstanceStatLastExportCompletedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *IntegrationInstanceStatLastExportCompletedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIntegrationInstanceStatLastExportCompletedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIntegrationInstanceStatLastExportCompletedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIntegrationInstanceStatLastExportCompletedDateObject(o.Rfc3339, false),
	}
}

func (o *IntegrationInstanceStatLastExportCompletedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationInstanceStatLastExportCompletedDate) FromMap(kv map[string]interface{}) {

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

// IntegrationInstanceStatLastExportRequestedDate represents the object structure for last_export_requested_date
type IntegrationInstanceStatLastExportRequestedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIntegrationInstanceStatLastExportRequestedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationInstanceStatLastExportRequestedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *IntegrationInstanceStatLastExportRequestedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIntegrationInstanceStatLastExportRequestedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIntegrationInstanceStatLastExportRequestedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIntegrationInstanceStatLastExportRequestedDateObject(o.Rfc3339, false),
	}
}

func (o *IntegrationInstanceStatLastExportRequestedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationInstanceStatLastExportRequestedDate) FromMap(kv map[string]interface{}) {

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

// IntegrationInstanceStatLastHistoricalCompletedDate represents the object structure for last_historical_completed_date
type IntegrationInstanceStatLastHistoricalCompletedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIntegrationInstanceStatLastHistoricalCompletedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationInstanceStatLastHistoricalCompletedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *IntegrationInstanceStatLastHistoricalCompletedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIntegrationInstanceStatLastHistoricalCompletedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIntegrationInstanceStatLastHistoricalCompletedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIntegrationInstanceStatLastHistoricalCompletedDateObject(o.Rfc3339, false),
	}
}

func (o *IntegrationInstanceStatLastHistoricalCompletedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationInstanceStatLastHistoricalCompletedDate) FromMap(kv map[string]interface{}) {

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

// IntegrationInstanceStatLastHistoricalRequestedDate represents the object structure for last_historical_requested_date
type IntegrationInstanceStatLastHistoricalRequestedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIntegrationInstanceStatLastHistoricalRequestedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationInstanceStatLastHistoricalRequestedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *IntegrationInstanceStatLastHistoricalRequestedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIntegrationInstanceStatLastHistoricalRequestedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIntegrationInstanceStatLastHistoricalRequestedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIntegrationInstanceStatLastHistoricalRequestedDateObject(o.Rfc3339, false),
	}
}

func (o *IntegrationInstanceStatLastHistoricalRequestedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationInstanceStatLastHistoricalRequestedDate) FromMap(kv map[string]interface{}) {

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

// IntegrationInstanceStatLastProcessingDate represents the object structure for last_processing_date
type IntegrationInstanceStatLastProcessingDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIntegrationInstanceStatLastProcessingDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationInstanceStatLastProcessingDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *IntegrationInstanceStatLastProcessingDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIntegrationInstanceStatLastProcessingDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIntegrationInstanceStatLastProcessingDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIntegrationInstanceStatLastProcessingDateObject(o.Rfc3339, false),
	}
}

func (o *IntegrationInstanceStatLastProcessingDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationInstanceStatLastProcessingDate) FromMap(kv map[string]interface{}) {

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

// IntegrationInstanceStat Data about an integration instance that is useful but too common to be sent in a dbchange.
type IntegrationInstanceStat struct {
	// CreatedAt the date the record was created in Epoch time
	CreatedAt int64 `json:"created_ts" codec:"created_ts" bson:"created_ts" yaml:"created_ts" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the id of the integration instance that these stats are about.
	IntegrationInstanceID string `json:"integration_instance_id" codec:"integration_instance_id" bson:"integration_instance_id" yaml:"integration_instance_id" faker:"-"`
	// LastExportCompletedDate when the export response was received
	LastExportCompletedDate IntegrationInstanceStatLastExportCompletedDate `json:"last_export_completed_date" codec:"last_export_completed_date" bson:"last_export_completed_date" yaml:"last_export_completed_date" faker:"-"`
	// LastExportRequestedDate when the export request was made
	LastExportRequestedDate IntegrationInstanceStatLastExportRequestedDate `json:"last_export_requested_date" codec:"last_export_requested_date" bson:"last_export_requested_date" yaml:"last_export_requested_date" faker:"-"`
	// LastHistoricalCompletedDate when the last hitorical export completed
	LastHistoricalCompletedDate IntegrationInstanceStatLastHistoricalCompletedDate `json:"last_historical_completed_date" codec:"last_historical_completed_date" bson:"last_historical_completed_date" yaml:"last_historical_completed_date" faker:"-"`
	// LastHistoricalRequestedDate when the last historical export was requested
	LastHistoricalRequestedDate IntegrationInstanceStatLastHistoricalRequestedDate `json:"last_historical_requested_date" codec:"last_historical_requested_date" bson:"last_historical_requested_date" yaml:"last_historical_requested_date" faker:"-"`
	// LastProcessingDate when the data was last processed
	LastProcessingDate IntegrationInstanceStatLastProcessingDate `json:"last_processing_date" codec:"last_processing_date" bson:"last_processing_date" yaml:"last_processing_date" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// UpdatedAt the date the record was updated in Epoch time
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*IntegrationInstanceStat)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*IntegrationInstanceStat)(nil)

func toIntegrationInstanceStatObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationInstanceStat:
		return v.ToMap()

	case IntegrationInstanceStatLastExportCompletedDate:
		return v.ToMap()

	case IntegrationInstanceStatLastExportRequestedDate:
		return v.ToMap()

	case IntegrationInstanceStatLastHistoricalCompletedDate:
		return v.ToMap()

	case IntegrationInstanceStatLastHistoricalRequestedDate:
		return v.ToMap()

	case IntegrationInstanceStatLastProcessingDate:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of IntegrationInstanceStat
func (o *IntegrationInstanceStat) String() string {
	return fmt.Sprintf("agent.IntegrationInstanceStat<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *IntegrationInstanceStat) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *IntegrationInstanceStat) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *IntegrationInstanceStat) GetTableName() string {
	return IntegrationInstanceStatTable.String()
}

// GetModelName returns the name of the model
func (o *IntegrationInstanceStat) GetModelName() datamodel.ModelNameType {
	return IntegrationInstanceStatModelName
}

// NewIntegrationInstanceStatID provides a template for generating an ID field for IntegrationInstanceStat
func NewIntegrationInstanceStatID(IntegrationInstanceID string) string {
	return hash.Values(IntegrationInstanceID)
}

func (o *IntegrationInstanceStat) setDefaults(frommap bool) {

	if o.ID == "" {
		// set the id from the spec provided in the model
		o.ID = hash.Values(o.IntegrationInstanceID)
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *IntegrationInstanceStat) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *IntegrationInstanceStat) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *IntegrationInstanceStat) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *IntegrationInstanceStat) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *IntegrationInstanceStat) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *IntegrationInstanceStat) IsMutable() bool {
	return true
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *IntegrationInstanceStat) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *IntegrationInstanceStat) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *IntegrationInstanceStat) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = IntegrationInstanceStatModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *IntegrationInstanceStat) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *IntegrationInstanceStat) GetCustomerID() string {
	return o.CustomerID
}

// SetCustomerID will return the customer_id
func (o *IntegrationInstanceStat) SetCustomerID(id string) {
	o.CustomerID = id
}

// GetRefType will return the ref_type
func (o *IntegrationInstanceStat) GetRefType() string {
	return o.RefType
}

// SetRefType will return the ref_type
func (o *IntegrationInstanceStat) SetRefType(t string) {
	o.RefType = t
}

// Clone returns an exact copy of IntegrationInstanceStat
func (o *IntegrationInstanceStat) Clone() datamodel.Model {
	c := new(IntegrationInstanceStat)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *IntegrationInstanceStat) Anon() datamodel.Model {
	c := new(IntegrationInstanceStat)
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
func (o *IntegrationInstanceStat) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *IntegrationInstanceStat) UnmarshalJSON(data []byte) error {
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
func (o *IntegrationInstanceStat) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *IntegrationInstanceStat) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two IntegrationInstanceStat objects are equal
func (o *IntegrationInstanceStat) IsEqual(other *IntegrationInstanceStat) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *IntegrationInstanceStat) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"created_ts":                     toIntegrationInstanceStatObject(o.CreatedAt, false),
		"customer_id":                    toIntegrationInstanceStatObject(o.CustomerID, false),
		"id":                             toIntegrationInstanceStatObject(o.ID, false),
		"integration_instance_id":        toIntegrationInstanceStatObject(o.IntegrationInstanceID, false),
		"last_export_completed_date":     toIntegrationInstanceStatObject(o.LastExportCompletedDate, false),
		"last_export_requested_date":     toIntegrationInstanceStatObject(o.LastExportRequestedDate, false),
		"last_historical_completed_date": toIntegrationInstanceStatObject(o.LastHistoricalCompletedDate, false),
		"last_historical_requested_date": toIntegrationInstanceStatObject(o.LastHistoricalRequestedDate, false),
		"last_processing_date":           toIntegrationInstanceStatObject(o.LastProcessingDate, false),
		"ref_id":                         toIntegrationInstanceStatObject(o.RefID, false),
		"ref_type":                       toIntegrationInstanceStatObject(o.RefType, false),
		"updated_ts":                     toIntegrationInstanceStatObject(o.UpdatedAt, false),
		"hashcode":                       toIntegrationInstanceStatObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationInstanceStat) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["created_ts"].(int64); ok {
		o.CreatedAt = val
	} else {
		if val, ok := kv["created_ts"]; ok {
			if val == nil {
				o.CreatedAt = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.CreatedAt = number.ToInt64Any(val)
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
	if val, ok := kv["integration_instance_id"].(string); ok {
		o.IntegrationInstanceID = val
	} else {
		if val, ok := kv["integration_instance_id"]; ok {
			if val == nil {
				o.IntegrationInstanceID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.IntegrationInstanceID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["last_export_completed_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastExportCompletedDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationInstanceStatLastExportCompletedDate); ok {
			// struct
			o.LastExportCompletedDate = sv
		} else if sp, ok := val.(*IntegrationInstanceStatLastExportCompletedDate); ok {
			// struct pointer
			o.LastExportCompletedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.LastExportCompletedDate.Epoch = dt.Epoch
			o.LastExportCompletedDate.Rfc3339 = dt.Rfc3339
			o.LastExportCompletedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.LastExportCompletedDate.Epoch = dt.Epoch
			o.LastExportCompletedDate.Rfc3339 = dt.Rfc3339
			o.LastExportCompletedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.LastExportCompletedDate.Epoch = dt.Epoch
				o.LastExportCompletedDate.Rfc3339 = dt.Rfc3339
				o.LastExportCompletedDate.Offset = dt.Offset
			}
		}
	} else {
		o.LastExportCompletedDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["last_export_requested_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastExportRequestedDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationInstanceStatLastExportRequestedDate); ok {
			// struct
			o.LastExportRequestedDate = sv
		} else if sp, ok := val.(*IntegrationInstanceStatLastExportRequestedDate); ok {
			// struct pointer
			o.LastExportRequestedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.LastExportRequestedDate.Epoch = dt.Epoch
			o.LastExportRequestedDate.Rfc3339 = dt.Rfc3339
			o.LastExportRequestedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.LastExportRequestedDate.Epoch = dt.Epoch
			o.LastExportRequestedDate.Rfc3339 = dt.Rfc3339
			o.LastExportRequestedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.LastExportRequestedDate.Epoch = dt.Epoch
				o.LastExportRequestedDate.Rfc3339 = dt.Rfc3339
				o.LastExportRequestedDate.Offset = dt.Offset
			}
		}
	} else {
		o.LastExportRequestedDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["last_historical_completed_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastHistoricalCompletedDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationInstanceStatLastHistoricalCompletedDate); ok {
			// struct
			o.LastHistoricalCompletedDate = sv
		} else if sp, ok := val.(*IntegrationInstanceStatLastHistoricalCompletedDate); ok {
			// struct pointer
			o.LastHistoricalCompletedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.LastHistoricalCompletedDate.Epoch = dt.Epoch
			o.LastHistoricalCompletedDate.Rfc3339 = dt.Rfc3339
			o.LastHistoricalCompletedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.LastHistoricalCompletedDate.Epoch = dt.Epoch
			o.LastHistoricalCompletedDate.Rfc3339 = dt.Rfc3339
			o.LastHistoricalCompletedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.LastHistoricalCompletedDate.Epoch = dt.Epoch
				o.LastHistoricalCompletedDate.Rfc3339 = dt.Rfc3339
				o.LastHistoricalCompletedDate.Offset = dt.Offset
			}
		}
	} else {
		o.LastHistoricalCompletedDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["last_historical_requested_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastHistoricalRequestedDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationInstanceStatLastHistoricalRequestedDate); ok {
			// struct
			o.LastHistoricalRequestedDate = sv
		} else if sp, ok := val.(*IntegrationInstanceStatLastHistoricalRequestedDate); ok {
			// struct pointer
			o.LastHistoricalRequestedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.LastHistoricalRequestedDate.Epoch = dt.Epoch
			o.LastHistoricalRequestedDate.Rfc3339 = dt.Rfc3339
			o.LastHistoricalRequestedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.LastHistoricalRequestedDate.Epoch = dt.Epoch
			o.LastHistoricalRequestedDate.Rfc3339 = dt.Rfc3339
			o.LastHistoricalRequestedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.LastHistoricalRequestedDate.Epoch = dt.Epoch
				o.LastHistoricalRequestedDate.Rfc3339 = dt.Rfc3339
				o.LastHistoricalRequestedDate.Offset = dt.Offset
			}
		}
	} else {
		o.LastHistoricalRequestedDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["last_processing_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastProcessingDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationInstanceStatLastProcessingDate); ok {
			// struct
			o.LastProcessingDate = sv
		} else if sp, ok := val.(*IntegrationInstanceStatLastProcessingDate); ok {
			// struct pointer
			o.LastProcessingDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.LastProcessingDate.Epoch = dt.Epoch
			o.LastProcessingDate.Rfc3339 = dt.Rfc3339
			o.LastProcessingDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.LastProcessingDate.Epoch = dt.Epoch
			o.LastProcessingDate.Rfc3339 = dt.Rfc3339
			o.LastProcessingDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.LastProcessingDate.Epoch = dt.Epoch
				o.LastProcessingDate.Rfc3339 = dt.Rfc3339
				o.LastProcessingDate.Offset = dt.Offset
			}
		}
	} else {
		o.LastProcessingDate.FromMap(map[string]interface{}{})
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
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *IntegrationInstanceStat) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CreatedAt)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.LastExportCompletedDate)
	args = append(args, o.LastExportRequestedDate)
	args = append(args, o.LastHistoricalCompletedDate)
	args = append(args, o.LastHistoricalRequestedDate)
	args = append(args, o.LastProcessingDate)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.UpdatedAt)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// SetIntegrationInstanceID will set the integration instance ID
func (o *IntegrationInstanceStat) SetIntegrationInstanceID(id string) {
	o.IntegrationInstanceID = id
}

// GetIntegrationInstanceID will return the integration instance ID
func (o *IntegrationInstanceStat) GetIntegrationInstanceID() *string {
	return &o.IntegrationInstanceID
}

// GetHydrationQuery returns a query for all fields, and one level deep of relations.
// This query requires "id" to be in the query variables.
func (o *IntegrationInstanceStat) GetHydrationQuery() string {
	return `query GoIntegrationInstanceStatQuery($id: ID!) {
	agent {
		IntegrationInstanceStat(_id: $id) {
			created_ts
			customer_id
			_id
			integration_instance_id
			last_export_completed_date {
			epoch
			offset
			rfc3339
			}
			last_export_requested_date {
			epoch
			offset
			rfc3339
			}
			last_historical_completed_date {
			epoch
			offset
			rfc3339
			}
			last_historical_requested_date {
			epoch
			offset
			rfc3339
			}
			last_processing_date {
			epoch
			offset
			rfc3339
			}
			ref_id
			ref_type
			updated_ts
		}
	}
}
`
}

func getIntegrationInstanceStatQueryFields() string {
	var sb strings.Builder

	// scalar
	sb.WriteString("\t\t\tcreated_ts\n")
	// scalar
	sb.WriteString("\t\t\tcustomer_id\n")
	// id
	sb.WriteString("\t\t\t_id\n")
	// scalar
	sb.WriteString("\t\t\tintegration_instance_id\n")
	// object with fields
	sb.WriteString("\t\t\tlast_export_completed_date {\n")

	// scalar
	sb.WriteString("\t\t\tepoch\n")
	// scalar
	sb.WriteString("\t\t\toffset\n")
	// scalar
	sb.WriteString("\t\t\trfc3339\n")
	sb.WriteString("\t\t\t}\n")
	// object with fields
	sb.WriteString("\t\t\tlast_export_requested_date {\n")

	// scalar
	sb.WriteString("\t\t\tepoch\n")
	// scalar
	sb.WriteString("\t\t\toffset\n")
	// scalar
	sb.WriteString("\t\t\trfc3339\n")
	sb.WriteString("\t\t\t}\n")
	// object with fields
	sb.WriteString("\t\t\tlast_historical_completed_date {\n")

	// scalar
	sb.WriteString("\t\t\tepoch\n")
	// scalar
	sb.WriteString("\t\t\toffset\n")
	// scalar
	sb.WriteString("\t\t\trfc3339\n")
	sb.WriteString("\t\t\t}\n")
	// object with fields
	sb.WriteString("\t\t\tlast_historical_requested_date {\n")

	// scalar
	sb.WriteString("\t\t\tepoch\n")
	// scalar
	sb.WriteString("\t\t\toffset\n")
	// scalar
	sb.WriteString("\t\t\trfc3339\n")
	sb.WriteString("\t\t\t}\n")
	// object with fields
	sb.WriteString("\t\t\tlast_processing_date {\n")

	// scalar
	sb.WriteString("\t\t\tepoch\n")
	// scalar
	sb.WriteString("\t\t\toffset\n")
	// scalar
	sb.WriteString("\t\t\trfc3339\n")
	sb.WriteString("\t\t\t}\n")
	// scalar
	sb.WriteString("\t\t\tref_id\n")
	// scalar
	sb.WriteString("\t\t\tref_type\n")
	// scalar
	sb.WriteString("\t\t\tupdated_ts\n")
	return sb.String()
}

// IntegrationInstanceStatPageInfo is a grapqhl PageInfo
type IntegrationInstanceStatPageInfo struct {
	StartCursor     string `json:"startCursor,omitempty"`
	EndCursor       string `json:"endCursor,omitempty"`
	HasNextPage     bool   `json:"hasNextPage,omitempty"`
	HasPreviousPage bool   `json:"hasPreviousPage,omitempty"`
}

// IntegrationInstanceStatConnection is a grapqhl connection
type IntegrationInstanceStatConnection struct {
	Edges      []*IntegrationInstanceStatEdge   `json:"edges,omitempty"`
	PageInfo   IntegrationInstanceStatPageInfo  `json:"pageInfo,omitempty"`
	CacheInfo  IntegrationInstanceStatCacheInfo `json:"cacheInfo,omitempty"`
	TotalCount *int64                           `json:"totalCount,omitempty"`
}

// IntegrationInstanceStatEdge is a grapqhl edge
type IntegrationInstanceStatEdge struct {
	Node *IntegrationInstanceStat `json:"node,omitempty"`
}

// QueryManyIntegrationInstanceStatNode is a grapqhl query many node
type QueryManyIntegrationInstanceStatNode struct {
	Object *IntegrationInstanceStatConnection `json:"IntegrationInstanceStats,omitempty"`
}

// QueryManyIntegrationInstanceStatData is a grapqhl query many data node
type QueryManyIntegrationInstanceStatData struct {
	Data *QueryManyIntegrationInstanceStatNode `json:"agent,omitempty"`
}

// QueryOneIntegrationInstanceStatNode is a grapqhl query one node
type QueryOneIntegrationInstanceStatNode struct {
	Object *IntegrationInstanceStat `json:"IntegrationInstanceStat,omitempty"`
}

// QueryOneIntegrationInstanceStatData is a grapqhl query one data node
type QueryOneIntegrationInstanceStatData struct {
	Data *QueryOneIntegrationInstanceStatNode `json:"agent,omitempty"`
}

// IntegrationInstanceStatQuery is query struct
type IntegrationInstanceStatQuery struct {
	Filters []string      `json:"filters,omitempty"`
	Params  []interface{} `json:"params,omitempty"`
}

// IntegrationInstanceStatOrder is the order direction
type IntegrationInstanceStatOrder *string

var (
	// IntegrationInstanceStatOrderDesc is descending
	IntegrationInstanceStatOrderDesc IntegrationInstanceStatOrder = pstrings.Pointer("DESC")
	// IntegrationInstanceStatOrderAsc is ascending
	IntegrationInstanceStatOrderAsc IntegrationInstanceStatOrder = pstrings.Pointer("ASC")
)

// IntegrationInstanceStatQueryInput is query input struct
type IntegrationInstanceStatQueryInput struct {
	First   *int64                        `json:"first,omitempty"`
	Last    *int64                        `json:"last,omitempty"`
	Before  *string                       `json:"before,omitempty"`
	After   *string                       `json:"after,omitempty"`
	Query   *IntegrationInstanceStatQuery `json:"query,omitempty"`
	OrderBy *string                       `json:"orderBy,omitempty"`
	Order   IntegrationInstanceStatOrder  `json:"order,omitempty"`
	NoCache boolean                       `json:"nocache,omitempty"`
}

// NewIntegrationInstanceStatQuery is a convenience for building a *IntegrationInstanceStatQuery
func NewIntegrationInstanceStatQuery(params ...interface{}) *IntegrationInstanceStatQueryInput {
	if len(params)%2 != 0 {
		panic("incorrect number of arguments passed")
	}
	q := &IntegrationInstanceStatQuery{
		Filters: make([]string, 0),
		Params:  make([]interface{}, 0),
	}
	for i := 0; i < len(params); i += 2 {
		q.Filters = append(q.Filters, params[i].(string))
		q.Params = append(q.Params, params[i+1])
	}
	return &IntegrationInstanceStatQueryInput{
		Query: q,
	}
}

// FindIntegrationInstanceStat will query an IntegrationInstanceStat by id
func FindIntegrationInstanceStat(client graphql.Client, id string) (*IntegrationInstanceStat, error) {
	variables := make(graphql.Variables)
	variables["id"] = id
	var sb strings.Builder
	sb.WriteString("query GoIntegrationInstanceStatQuery($id: ID!) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tIntegrationInstanceStat(_id: $id) {\n")
	sb.WriteString(getIntegrationInstanceStatQueryFields())
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res QueryOneIntegrationInstanceStatData
	if err := client.Query(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	if res.Data != nil {
		return res.Data.Object, nil
	}
	return nil, nil
}

// FindIntegrationInstanceStatWithoutCache will query an IntegrationInstanceStat by id avoiding the cache
func FindIntegrationInstanceStatWithoutCache(client graphql.Client, id string) (*IntegrationInstanceStat, error) {
	variables := make(graphql.Variables)
	variables["id"] = id
	variables["nocache"] = true
	var sb strings.Builder
	sb.WriteString("query GoIntegrationInstanceStatQuery($id: ID!, $nocache: Boolean) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tIntegrationInstanceStat(_id: $id, nocache: $nocache) {\n")
	sb.WriteString(getIntegrationInstanceStatQueryFields())
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res QueryOneIntegrationInstanceStatData
	if err := client.Query(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	if res.Data != nil {
		return res.Data.Object, nil
	}
	return nil, nil
}

// FindIntegrationInstanceStats will query for any IntegrationInstanceStats matching the query
func FindIntegrationInstanceStats(client graphql.Client, input *IntegrationInstanceStatQueryInput) (*IntegrationInstanceStatConnection, error) {
	variables := make(graphql.Variables)
	if input != nil {
		variables["first"] = input.First
		variables["last"] = input.Last
		variables["before"] = input.Before
		variables["after"] = input.After
		variables["query"] = input.Query
		if input.OrderBy != nil {
			variables["orderBy"] = *input.OrderBy
		}
		if input.Order != nil {
			variables["order"] = *input.Order
		}
	}
	var sb strings.Builder
	sb.WriteString("query GoIntegrationInstanceStatQueryMany($first: Int, $last: Int, $before: Cursor, $after: Cursor, $query: QueryInput, $order: SortOrderEnum, $orderBy: AgentIntegrationInstanceStatColumnEnum, $nocache: Boolean) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tIntegrationInstanceStats(first: $first last: $last before: $before after: $after query: $query orderBy: $orderBy order: $order nocache: $nocache) {\n")
	sb.WriteString("\t\t\tpageInfo {\n")
	sb.WriteString("\t\t\t\tstartCursor\n")
	sb.WriteString("\t\t\t\tendCursor\n")
	sb.WriteString("\t\t\t\thasNextPage\n")
	sb.WriteString("\t\t\t\thasPreviousPage\n")
	sb.WriteString("\t\t\t}\n")
	sb.WriteString("\t\t\tcacheInfo {\n")
	sb.WriteString("\t\t\t\tcached\n")
	sb.WriteString("\t\t\t\tid\n")
	sb.WriteString("\t\t\t\tetag\n")
	sb.WriteString("\t\t\t}\n")
	sb.WriteString("\t\t\ttotalCount\n")
	sb.WriteString("\t\t\tedges {\n")
	sb.WriteString("\t\t\t\tnode {\n")
	sb.WriteString(getIntegrationInstanceStatQueryFields())
	sb.WriteString("\t\t\t\t}\n")
	sb.WriteString("\t\t\t}\n")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res QueryManyIntegrationInstanceStatData
	if err := client.Query(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	return res.Data.Object, nil
}

// FindIntegrationInstanceStatsPaginatedCallback is a callback function for handling each page
type FindIntegrationInstanceStatsPaginatedCallback func(conn *IntegrationInstanceStatConnection) (bool, error)

// FindIntegrationInstanceStatsPaginated will query for any IntegrationInstanceStats matching the query and return each page callback
func FindIntegrationInstanceStatsPaginated(client graphql.Client, query *IntegrationInstanceStatQuery, pageSize int64, callback FindIntegrationInstanceStatsPaginatedCallback) error {
	input := &IntegrationInstanceStatQueryInput{
		First: &pageSize,
		Query: query,
	}
	for {
		res, err := FindIntegrationInstanceStats(client, input)
		if err != nil {
			return err
		}
		if res == nil {
			break
		}
		ok, err := callback(res)
		if err != nil {
			return err
		}
		if !ok || !res.PageInfo.HasNextPage {
			break
		}
		input.After = &res.PageInfo.EndCursor
	}
	return nil
}

// UpdateIntegrationInstanceStatNode is a grapqhl update node
type UpdateIntegrationInstanceStatNode struct {
	Object *IntegrationInstanceStat `json:"updateIntegrationInstanceStat,omitempty"`
}

// UpdateIntegrationInstanceStatData is a grapqhl update data node
type UpdateIntegrationInstanceStatData struct {
	Data *UpdateIntegrationInstanceStatNode `json:"agent,omitempty"`
}

// ExecIntegrationInstanceStatUpdateMutation returns a graphql update mutation result for IntegrationInstanceStat
func ExecIntegrationInstanceStatUpdateMutation(client graphql.Client, id string, input graphql.Variables, upsert bool) (*IntegrationInstanceStat, error) {
	if !upsert && id == "" {
		return nil, fmt.Errorf("id is required with upsert false")
	}
	variables := make(graphql.Variables)
	variables["id"] = id
	variables["upsert"] = upsert
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation GoIntegrationInstanceStatUpdateMutation($id: String, $input: UpdateAgentIntegrationInstanceStatInput, $upsert: Boolean) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tupdateIntegrationInstanceStat(_id: $id, input: $input, upsert: $upsert) {\n")
	sb.WriteString(getIntegrationInstanceStatQueryFields())
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateIntegrationInstanceStatData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	return res.Data.Object, nil
}

// ExecIntegrationInstanceStatSilentUpdateMutation returns a graphql update mutation result for IntegrationInstanceStat
func ExecIntegrationInstanceStatSilentUpdateMutation(client graphql.Client, id string, input graphql.Variables, upsert bool) error {
	if !upsert && id == "" {
		return fmt.Errorf("id is required with upsert false")
	}
	variables := make(graphql.Variables)
	variables["id"] = id
	variables["upsert"] = upsert
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation GoIntegrationInstanceStatUpdateMutation($id: String, $input: UpdateAgentIntegrationInstanceStatInput, $upsert: Boolean) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tupdateIntegrationInstanceStat(_id: $id, input: $input, upsert: $upsert) {\n")
	sb.WriteString("\t\t\t_id")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateIntegrationInstanceStatData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return err
	}
	return nil
}

// ExecIntegrationInstanceStatDeleteMutation executes a graphql delete mutation for IntegrationInstanceStat
func ExecIntegrationInstanceStatDeleteMutation(client graphql.Client, id string) error {
	variables := make(graphql.Variables)
	variables["id"] = id
	var sb strings.Builder
	sb.WriteString("mutation GoIntegrationInstanceStatDeleteMutation($id: String!) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tdeleteIntegrationInstanceStat(_id: $id)\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res interface{}
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return err
	}
	return nil
}

func CreateIntegrationInstanceStat(client graphql.Client, model IntegrationInstanceStat) error {
	variables := make(graphql.Variables)
	input := model.ToMap()
	delete(input, "hashcode")
	delete(input, "customer_id")
	delete(input, "created_ts")
	delete(input, "id")
	input["_id"] = model.GetID()
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation GoCreateIntegrationInstanceStat($input: CreateAgentIntegrationInstanceStatInput!) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tcreateIntegrationInstanceStat(input: $input) {\n")
	sb.WriteString("\t\t\t_id")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateIntegrationInstanceStatData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return err
	}
	return nil
}

// IntegrationInstanceStatPartial is a partial struct for upsert mutations for IntegrationInstanceStat
type IntegrationInstanceStatPartial struct {
	// LastExportCompletedDate when the export response was received
	LastExportCompletedDate *IntegrationInstanceStatLastExportCompletedDate `json:"last_export_completed_date,omitempty"`
	// LastExportRequestedDate when the export request was made
	LastExportRequestedDate *IntegrationInstanceStatLastExportRequestedDate `json:"last_export_requested_date,omitempty"`
	// LastHistoricalCompletedDate when the last hitorical export completed
	LastHistoricalCompletedDate *IntegrationInstanceStatLastHistoricalCompletedDate `json:"last_historical_completed_date,omitempty"`
	// LastHistoricalRequestedDate when the last historical export was requested
	LastHistoricalRequestedDate *IntegrationInstanceStatLastHistoricalRequestedDate `json:"last_historical_requested_date,omitempty"`
	// LastProcessingDate when the data was last processed
	LastProcessingDate *IntegrationInstanceStatLastProcessingDate `json:"last_processing_date,omitempty"`
}

var _ datamodel.PartialModel = (*IntegrationInstanceStatPartial)(nil)

// GetModelName returns the name of the model
func (o *IntegrationInstanceStatPartial) GetModelName() datamodel.ModelNameType {
	return IntegrationInstanceStatModelName
}

// ToMap returns the object as a map
func (o *IntegrationInstanceStatPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"last_export_completed_date":     toIntegrationInstanceStatObject(o.LastExportCompletedDate, true),
		"last_export_requested_date":     toIntegrationInstanceStatObject(o.LastExportRequestedDate, true),
		"last_historical_completed_date": toIntegrationInstanceStatObject(o.LastHistoricalCompletedDate, true),
		"last_historical_requested_date": toIntegrationInstanceStatObject(o.LastHistoricalRequestedDate, true),
		"last_processing_date":           toIntegrationInstanceStatObject(o.LastProcessingDate, true),
	}
	for k, v := range kv {
		if v == nil || reflect.ValueOf(v).IsZero() {
			delete(kv, k)
		} else {
			if k == "last_export_completed_date" {
				if dt, ok := v.(*IntegrationInstanceStatLastExportCompletedDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}
			if k == "last_export_requested_date" {
				if dt, ok := v.(*IntegrationInstanceStatLastExportRequestedDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}
			if k == "last_historical_completed_date" {
				if dt, ok := v.(*IntegrationInstanceStatLastHistoricalCompletedDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}
			if k == "last_historical_requested_date" {
				if dt, ok := v.(*IntegrationInstanceStatLastHistoricalRequestedDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}
			if k == "last_processing_date" {
				if dt, ok := v.(*IntegrationInstanceStatLastProcessingDate); ok {
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
func (o *IntegrationInstanceStatPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *IntegrationInstanceStatPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *IntegrationInstanceStatPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *IntegrationInstanceStatPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *IntegrationInstanceStatPartial) FromMap(kv map[string]interface{}) {

	if o.LastExportCompletedDate == nil {
		o.LastExportCompletedDate = &IntegrationInstanceStatLastExportCompletedDate{}
	}

	if val, ok := kv["last_export_completed_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastExportCompletedDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationInstanceStatLastExportCompletedDate); ok {
			// struct
			o.LastExportCompletedDate = &sv
		} else if sp, ok := val.(*IntegrationInstanceStatLastExportCompletedDate); ok {
			// struct pointer
			o.LastExportCompletedDate = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.LastExportCompletedDate.Epoch = dt.Epoch
			o.LastExportCompletedDate.Rfc3339 = dt.Rfc3339
			o.LastExportCompletedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.LastExportCompletedDate.Epoch = dt.Epoch
			o.LastExportCompletedDate.Rfc3339 = dt.Rfc3339
			o.LastExportCompletedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.LastExportCompletedDate.Epoch = dt.Epoch
				o.LastExportCompletedDate.Rfc3339 = dt.Rfc3339
				o.LastExportCompletedDate.Offset = dt.Offset
			}
		}
	} else {
		o.LastExportCompletedDate.FromMap(map[string]interface{}{})
	}

	if o.LastExportRequestedDate == nil {
		o.LastExportRequestedDate = &IntegrationInstanceStatLastExportRequestedDate{}
	}

	if val, ok := kv["last_export_requested_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastExportRequestedDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationInstanceStatLastExportRequestedDate); ok {
			// struct
			o.LastExportRequestedDate = &sv
		} else if sp, ok := val.(*IntegrationInstanceStatLastExportRequestedDate); ok {
			// struct pointer
			o.LastExportRequestedDate = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.LastExportRequestedDate.Epoch = dt.Epoch
			o.LastExportRequestedDate.Rfc3339 = dt.Rfc3339
			o.LastExportRequestedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.LastExportRequestedDate.Epoch = dt.Epoch
			o.LastExportRequestedDate.Rfc3339 = dt.Rfc3339
			o.LastExportRequestedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.LastExportRequestedDate.Epoch = dt.Epoch
				o.LastExportRequestedDate.Rfc3339 = dt.Rfc3339
				o.LastExportRequestedDate.Offset = dt.Offset
			}
		}
	} else {
		o.LastExportRequestedDate.FromMap(map[string]interface{}{})
	}

	if o.LastHistoricalCompletedDate == nil {
		o.LastHistoricalCompletedDate = &IntegrationInstanceStatLastHistoricalCompletedDate{}
	}

	if val, ok := kv["last_historical_completed_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastHistoricalCompletedDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationInstanceStatLastHistoricalCompletedDate); ok {
			// struct
			o.LastHistoricalCompletedDate = &sv
		} else if sp, ok := val.(*IntegrationInstanceStatLastHistoricalCompletedDate); ok {
			// struct pointer
			o.LastHistoricalCompletedDate = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.LastHistoricalCompletedDate.Epoch = dt.Epoch
			o.LastHistoricalCompletedDate.Rfc3339 = dt.Rfc3339
			o.LastHistoricalCompletedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.LastHistoricalCompletedDate.Epoch = dt.Epoch
			o.LastHistoricalCompletedDate.Rfc3339 = dt.Rfc3339
			o.LastHistoricalCompletedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.LastHistoricalCompletedDate.Epoch = dt.Epoch
				o.LastHistoricalCompletedDate.Rfc3339 = dt.Rfc3339
				o.LastHistoricalCompletedDate.Offset = dt.Offset
			}
		}
	} else {
		o.LastHistoricalCompletedDate.FromMap(map[string]interface{}{})
	}

	if o.LastHistoricalRequestedDate == nil {
		o.LastHistoricalRequestedDate = &IntegrationInstanceStatLastHistoricalRequestedDate{}
	}

	if val, ok := kv["last_historical_requested_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastHistoricalRequestedDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationInstanceStatLastHistoricalRequestedDate); ok {
			// struct
			o.LastHistoricalRequestedDate = &sv
		} else if sp, ok := val.(*IntegrationInstanceStatLastHistoricalRequestedDate); ok {
			// struct pointer
			o.LastHistoricalRequestedDate = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.LastHistoricalRequestedDate.Epoch = dt.Epoch
			o.LastHistoricalRequestedDate.Rfc3339 = dt.Rfc3339
			o.LastHistoricalRequestedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.LastHistoricalRequestedDate.Epoch = dt.Epoch
			o.LastHistoricalRequestedDate.Rfc3339 = dt.Rfc3339
			o.LastHistoricalRequestedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.LastHistoricalRequestedDate.Epoch = dt.Epoch
				o.LastHistoricalRequestedDate.Rfc3339 = dt.Rfc3339
				o.LastHistoricalRequestedDate.Offset = dt.Offset
			}
		}
	} else {
		o.LastHistoricalRequestedDate.FromMap(map[string]interface{}{})
	}

	if o.LastProcessingDate == nil {
		o.LastProcessingDate = &IntegrationInstanceStatLastProcessingDate{}
	}

	if val, ok := kv["last_processing_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastProcessingDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationInstanceStatLastProcessingDate); ok {
			// struct
			o.LastProcessingDate = &sv
		} else if sp, ok := val.(*IntegrationInstanceStatLastProcessingDate); ok {
			// struct pointer
			o.LastProcessingDate = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.LastProcessingDate.Epoch = dt.Epoch
			o.LastProcessingDate.Rfc3339 = dt.Rfc3339
			o.LastProcessingDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.LastProcessingDate.Epoch = dt.Epoch
			o.LastProcessingDate.Rfc3339 = dt.Rfc3339
			o.LastProcessingDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.LastProcessingDate.Epoch = dt.Epoch
				o.LastProcessingDate.Rfc3339 = dt.Rfc3339
				o.LastProcessingDate.Offset = dt.Offset
			}
		}
	} else {
		o.LastProcessingDate.FromMap(map[string]interface{}{})
	}

	o.setDefaults(false)
}
