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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (

	// IntegrationInstanceTable is the default table name
	IntegrationInstanceTable datamodel.ModelNameType = "agent_integrationinstance"

	// IntegrationInstanceModelName is the model name
	IntegrationInstanceModelName datamodel.ModelNameType = "agent.IntegrationInstance"
)

const (
	// IntegrationInstanceModelActiveColumn is the column json value active
	IntegrationInstanceModelActiveColumn = "active"
	// IntegrationInstanceModelConfigColumn is the column json value config
	IntegrationInstanceModelConfigColumn = "config"
	// IntegrationInstanceModelCreatedByProfileIDColumn is the column json value created_by_profile_id
	IntegrationInstanceModelCreatedByProfileIDColumn = "created_by_profile_id"
	// IntegrationInstanceModelCreatedByUserIDColumn is the column json value created_by_user_id
	IntegrationInstanceModelCreatedByUserIDColumn = "created_by_user_id"
	// IntegrationInstanceModelCreatedDateColumn is the column json value created_date
	IntegrationInstanceModelCreatedDateColumn = "created_date"
	// IntegrationInstanceModelCreatedDateEpochColumn is the column json value epoch
	IntegrationInstanceModelCreatedDateEpochColumn = "epoch"
	// IntegrationInstanceModelCreatedDateOffsetColumn is the column json value offset
	IntegrationInstanceModelCreatedDateOffsetColumn = "offset"
	// IntegrationInstanceModelCreatedDateRfc3339Column is the column json value rfc3339
	IntegrationInstanceModelCreatedDateRfc3339Column = "rfc3339"
	// IntegrationInstanceModelCreatedAtColumn is the column json value created_ts
	IntegrationInstanceModelCreatedAtColumn = "created_ts"
	// IntegrationInstanceModelCustomerIDColumn is the column json value customer_id
	IntegrationInstanceModelCustomerIDColumn = "customer_id"
	// IntegrationInstanceModelEntityErrorsColumn is the column json value entity_errors
	IntegrationInstanceModelEntityErrorsColumn = "entity_errors"
	// IntegrationInstanceModelEntityErrorsErrorColumn is the column json value error
	IntegrationInstanceModelEntityErrorsErrorColumn = "error"
	// IntegrationInstanceModelEntityErrorsIDColumn is the column json value id
	IntegrationInstanceModelEntityErrorsIDColumn = "id"
	// IntegrationInstanceModelEntityErrorsRefIDColumn is the column json value ref_id
	IntegrationInstanceModelEntityErrorsRefIDColumn = "ref_id"
	// IntegrationInstanceModelErrorMessageColumn is the column json value error_message
	IntegrationInstanceModelErrorMessageColumn = "error_message"
	// IntegrationInstanceModelErroredColumn is the column json value errored
	IntegrationInstanceModelErroredColumn = "errored"
	// IntegrationInstanceModelExportAcknowledgedColumn is the column json value export_acknowledged
	IntegrationInstanceModelExportAcknowledgedColumn = "export_acknowledged"
	// IntegrationInstanceModelIDColumn is the column json value id
	IntegrationInstanceModelIDColumn = "id"
	// IntegrationInstanceModelIntegrationIDColumn is the column json value integration_id
	IntegrationInstanceModelIntegrationIDColumn = "integration_id"
	// IntegrationInstanceModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	IntegrationInstanceModelIntegrationInstanceIDColumn = "integration_instance_id"
	// IntegrationInstanceModelIntervalColumn is the column json value interval
	IntegrationInstanceModelIntervalColumn = "interval"
	// IntegrationInstanceModelLastExportCompletedDateColumn is the column json value last_export_completed_date
	IntegrationInstanceModelLastExportCompletedDateColumn = "last_export_completed_date"
	// IntegrationInstanceModelLastExportCompletedDateEpochColumn is the column json value epoch
	IntegrationInstanceModelLastExportCompletedDateEpochColumn = "epoch"
	// IntegrationInstanceModelLastExportCompletedDateOffsetColumn is the column json value offset
	IntegrationInstanceModelLastExportCompletedDateOffsetColumn = "offset"
	// IntegrationInstanceModelLastExportCompletedDateRfc3339Column is the column json value rfc3339
	IntegrationInstanceModelLastExportCompletedDateRfc3339Column = "rfc3339"
	// IntegrationInstanceModelLastExportRequestedDateColumn is the column json value last_export_requested_date
	IntegrationInstanceModelLastExportRequestedDateColumn = "last_export_requested_date"
	// IntegrationInstanceModelLastExportRequestedDateEpochColumn is the column json value epoch
	IntegrationInstanceModelLastExportRequestedDateEpochColumn = "epoch"
	// IntegrationInstanceModelLastExportRequestedDateOffsetColumn is the column json value offset
	IntegrationInstanceModelLastExportRequestedDateOffsetColumn = "offset"
	// IntegrationInstanceModelLastExportRequestedDateRfc3339Column is the column json value rfc3339
	IntegrationInstanceModelLastExportRequestedDateRfc3339Column = "rfc3339"
	// IntegrationInstanceModelLastProcessingDateColumn is the column json value last_processing_date
	IntegrationInstanceModelLastProcessingDateColumn = "last_processing_date"
	// IntegrationInstanceModelLastProcessingDateEpochColumn is the column json value epoch
	IntegrationInstanceModelLastProcessingDateEpochColumn = "epoch"
	// IntegrationInstanceModelLastProcessingDateOffsetColumn is the column json value offset
	IntegrationInstanceModelLastProcessingDateOffsetColumn = "offset"
	// IntegrationInstanceModelLastProcessingDateRfc3339Column is the column json value rfc3339
	IntegrationInstanceModelLastProcessingDateRfc3339Column = "rfc3339"
	// IntegrationInstanceModelLocationColumn is the column json value location
	IntegrationInstanceModelLocationColumn = "location"
	// IntegrationInstanceModelNameColumn is the column json value name
	IntegrationInstanceModelNameColumn = "name"
	// IntegrationInstanceModelPausedColumn is the column json value paused
	IntegrationInstanceModelPausedColumn = "paused"
	// IntegrationInstanceModelProcessedColumn is the column json value processed
	IntegrationInstanceModelProcessedColumn = "processed"
	// IntegrationInstanceModelRefIDColumn is the column json value ref_id
	IntegrationInstanceModelRefIDColumn = "ref_id"
	// IntegrationInstanceModelRefTypeColumn is the column json value ref_type
	IntegrationInstanceModelRefTypeColumn = "ref_type"
	// IntegrationInstanceModelStateColumn is the column json value state
	IntegrationInstanceModelStateColumn = "state"
	// IntegrationInstanceModelThrottledColumn is the column json value throttled
	IntegrationInstanceModelThrottledColumn = "throttled"
	// IntegrationInstanceModelThrottledUntilColumn is the column json value throttled_until
	IntegrationInstanceModelThrottledUntilColumn = "throttled_until"
	// IntegrationInstanceModelThrottledUntilEpochColumn is the column json value epoch
	IntegrationInstanceModelThrottledUntilEpochColumn = "epoch"
	// IntegrationInstanceModelThrottledUntilOffsetColumn is the column json value offset
	IntegrationInstanceModelThrottledUntilOffsetColumn = "offset"
	// IntegrationInstanceModelThrottledUntilRfc3339Column is the column json value rfc3339
	IntegrationInstanceModelThrottledUntilRfc3339Column = "rfc3339"
	// IntegrationInstanceModelUpdatedAtColumn is the column json value updated_ts
	IntegrationInstanceModelUpdatedAtColumn = "updated_ts"
	// IntegrationInstanceModelWebhooksColumn is the column json value webhooks
	IntegrationInstanceModelWebhooksColumn = "webhooks"
	// IntegrationInstanceModelWebhooksEnabledColumn is the column json value enabled
	IntegrationInstanceModelWebhooksEnabledColumn = "enabled"
	// IntegrationInstanceModelWebhooksErrorMessageColumn is the column json value error_message
	IntegrationInstanceModelWebhooksErrorMessageColumn = "error_message"
	// IntegrationInstanceModelWebhooksErroredColumn is the column json value errored
	IntegrationInstanceModelWebhooksErroredColumn = "errored"
	// IntegrationInstanceModelWebhooksRefIDColumn is the column json value ref_id
	IntegrationInstanceModelWebhooksRefIDColumn = "ref_id"
	// IntegrationInstanceModelWebhooksURLColumn is the column json value url
	IntegrationInstanceModelWebhooksURLColumn = "url"
)

// IntegrationInstanceCreatedDate represents the object structure for created_date
type IntegrationInstanceCreatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIntegrationInstanceCreatedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationInstanceCreatedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *IntegrationInstanceCreatedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIntegrationInstanceCreatedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIntegrationInstanceCreatedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIntegrationInstanceCreatedDateObject(o.Rfc3339, false),
	}
}

func (o *IntegrationInstanceCreatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationInstanceCreatedDate) FromMap(kv map[string]interface{}) {

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

// IntegrationInstanceEntityErrors represents the object structure for entity_errors
type IntegrationInstanceEntityErrors struct {
	// Error error message if set integration failed when trying to export
	Error string `json:"error" codec:"error" bson:"error" yaml:"error" faker:"-"`
	// ID entity id
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// RefID entity ref_id
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
}

func toIntegrationInstanceEntityErrorsObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationInstanceEntityErrors:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *IntegrationInstanceEntityErrors) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Error error message if set integration failed when trying to export
		"error": toIntegrationInstanceEntityErrorsObject(o.Error, false),
		// ID entity id
		"id": toIntegrationInstanceEntityErrorsObject(o.ID, false),
		// RefID entity ref_id
		"ref_id": toIntegrationInstanceEntityErrorsObject(o.RefID, false),
	}
}

func (o *IntegrationInstanceEntityErrors) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationInstanceEntityErrors) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["error"].(string); ok {
		o.Error = val
	} else {
		if val, ok := kv["error"]; ok {
			if val == nil {
				o.Error = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Error = fmt.Sprintf("%v", val)
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
	o.setDefaults(false)
}

// IntegrationInstanceLastExportCompletedDate represents the object structure for last_export_completed_date
type IntegrationInstanceLastExportCompletedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIntegrationInstanceLastExportCompletedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationInstanceLastExportCompletedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *IntegrationInstanceLastExportCompletedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIntegrationInstanceLastExportCompletedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIntegrationInstanceLastExportCompletedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIntegrationInstanceLastExportCompletedDateObject(o.Rfc3339, false),
	}
}

func (o *IntegrationInstanceLastExportCompletedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationInstanceLastExportCompletedDate) FromMap(kv map[string]interface{}) {

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

// IntegrationInstanceLastExportRequestedDate represents the object structure for last_export_requested_date
type IntegrationInstanceLastExportRequestedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIntegrationInstanceLastExportRequestedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationInstanceLastExportRequestedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *IntegrationInstanceLastExportRequestedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIntegrationInstanceLastExportRequestedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIntegrationInstanceLastExportRequestedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIntegrationInstanceLastExportRequestedDateObject(o.Rfc3339, false),
	}
}

func (o *IntegrationInstanceLastExportRequestedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationInstanceLastExportRequestedDate) FromMap(kv map[string]interface{}) {

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

// IntegrationInstanceLastProcessingDate represents the object structure for last_processing_date
type IntegrationInstanceLastProcessingDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIntegrationInstanceLastProcessingDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationInstanceLastProcessingDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *IntegrationInstanceLastProcessingDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIntegrationInstanceLastProcessingDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIntegrationInstanceLastProcessingDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIntegrationInstanceLastProcessingDateObject(o.Rfc3339, false),
	}
}

func (o *IntegrationInstanceLastProcessingDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationInstanceLastProcessingDate) FromMap(kv map[string]interface{}) {

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

// IntegrationInstanceLocation is the enumeration type for location
type IntegrationInstanceLocation int32

// toIntegrationInstanceLocationPointer is the enumeration pointer type for location
func toIntegrationInstanceLocationPointer(v int32) *IntegrationInstanceLocation {
	nv := IntegrationInstanceLocation(v)
	return &nv
}

// toIntegrationInstanceLocationEnum is the enumeration pointer wrapper for location
func toIntegrationInstanceLocationEnum(v *IntegrationInstanceLocation) string {
	if v == nil {
		return toIntegrationInstanceLocationPointer(0).String()
	}
	return v.String()
}

// UnmarshalBSONValue for unmarshaling value
func (v *IntegrationInstanceLocation) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = IntegrationInstanceLocation(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "PRIVATE":
			*v = IntegrationInstanceLocation(0)
		case "CLOUD":
			*v = IntegrationInstanceLocation(1)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v IntegrationInstanceLocation) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "PRIVATE":
		v = 0
	case "CLOUD":
		v = 1
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v IntegrationInstanceLocation) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("PRIVATE")
	case 1:
		return json.Marshal("CLOUD")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for Location
func (v IntegrationInstanceLocation) String() string {
	switch int32(v) {
	case 0:
		return "PRIVATE"
	case 1:
		return "CLOUD"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *IntegrationInstanceLocation) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case IntegrationInstanceLocation:
		*v = val
	case int32:
		*v = IntegrationInstanceLocation(int32(val))
	case int:
		*v = IntegrationInstanceLocation(int32(val))
	case string:
		switch val {
		case "PRIVATE":
			*v = IntegrationInstanceLocation(0)
		case "CLOUD":
			*v = IntegrationInstanceLocation(1)
		}
	}
	return nil
}

const (
	// IntegrationInstanceLocationPrivate is the enumeration value for private
	IntegrationInstanceLocationPrivate IntegrationInstanceLocation = 0
	// IntegrationInstanceLocationCloud is the enumeration value for cloud
	IntegrationInstanceLocationCloud IntegrationInstanceLocation = 1
)

// IntegrationInstanceState is the enumeration type for state
type IntegrationInstanceState int32

// toIntegrationInstanceStatePointer is the enumeration pointer type for state
func toIntegrationInstanceStatePointer(v int32) *IntegrationInstanceState {
	nv := IntegrationInstanceState(v)
	return &nv
}

// toIntegrationInstanceStateEnum is the enumeration pointer wrapper for state
func toIntegrationInstanceStateEnum(v *IntegrationInstanceState) string {
	if v == nil {
		return toIntegrationInstanceStatePointer(0).String()
	}
	return v.String()
}

// UnmarshalBSONValue for unmarshaling value
func (v *IntegrationInstanceState) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = IntegrationInstanceState(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "IDLE":
			*v = IntegrationInstanceState(0)
		case "EXPORTING":
			*v = IntegrationInstanceState(1)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v IntegrationInstanceState) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "IDLE":
		v = 0
	case "EXPORTING":
		v = 1
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v IntegrationInstanceState) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("IDLE")
	case 1:
		return json.Marshal("EXPORTING")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for State
func (v IntegrationInstanceState) String() string {
	switch int32(v) {
	case 0:
		return "IDLE"
	case 1:
		return "EXPORTING"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *IntegrationInstanceState) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case IntegrationInstanceState:
		*v = val
	case int32:
		*v = IntegrationInstanceState(int32(val))
	case int:
		*v = IntegrationInstanceState(int32(val))
	case string:
		switch val {
		case "IDLE":
			*v = IntegrationInstanceState(0)
		case "EXPORTING":
			*v = IntegrationInstanceState(1)
		}
	}
	return nil
}

const (
	// IntegrationInstanceStateIdle is the enumeration value for idle
	IntegrationInstanceStateIdle IntegrationInstanceState = 0
	// IntegrationInstanceStateExporting is the enumeration value for exporting
	IntegrationInstanceStateExporting IntegrationInstanceState = 1
)

// IntegrationInstanceThrottledUntil represents the object structure for throttled_until
type IntegrationInstanceThrottledUntil struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIntegrationInstanceThrottledUntilObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationInstanceThrottledUntil:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *IntegrationInstanceThrottledUntil) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIntegrationInstanceThrottledUntilObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIntegrationInstanceThrottledUntilObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIntegrationInstanceThrottledUntilObject(o.Rfc3339, false),
	}
}

func (o *IntegrationInstanceThrottledUntil) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationInstanceThrottledUntil) FromMap(kv map[string]interface{}) {

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

// IntegrationInstanceWebhooks represents the object structure for webhooks
type IntegrationInstanceWebhooks struct {
	// Enabled if webhooks are enabled for this instance
	Enabled bool `json:"enabled" codec:"enabled" bson:"enabled" yaml:"enabled" faker:"-"`
	// ErrorMessage the error message
	ErrorMessage *string `json:"error_message,omitempty" codec:"error_message,omitempty" bson:"error_message" yaml:"error_message,omitempty" faker:"-"`
	// Errored if the webhook has an error
	Errored bool `json:"errored" codec:"errored" bson:"errored" yaml:"errored" faker:"-"`
	// RefID an optional reference id related to the webhook and what its related to
	RefID *string `json:"ref_id,omitempty" codec:"ref_id,omitempty" bson:"ref_id" yaml:"ref_id,omitempty" faker:"-"`
	// URL the url the webhook for the webhook
	URL *string `json:"url,omitempty" codec:"url,omitempty" bson:"url" yaml:"url,omitempty" faker:"-"`
}

func toIntegrationInstanceWebhooksObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationInstanceWebhooks:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *IntegrationInstanceWebhooks) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Enabled if webhooks are enabled for this instance
		"enabled": toIntegrationInstanceWebhooksObject(o.Enabled, false),
		// ErrorMessage the error message
		"error_message": toIntegrationInstanceWebhooksObject(o.ErrorMessage, true),
		// Errored if the webhook has an error
		"errored": toIntegrationInstanceWebhooksObject(o.Errored, false),
		// RefID an optional reference id related to the webhook and what its related to
		"ref_id": toIntegrationInstanceWebhooksObject(o.RefID, true),
		// URL the url the webhook for the webhook
		"url": toIntegrationInstanceWebhooksObject(o.URL, true),
	}
}

func (o *IntegrationInstanceWebhooks) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationInstanceWebhooks) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["enabled"].(bool); ok {
		o.Enabled = val
	} else {
		if val, ok := kv["enabled"]; ok {
			if val == nil {
				o.Enabled = false
			} else {
				o.Enabled = number.ToBoolAny(val)
			}
		}
	}
	if val, ok := kv["error_message"].(*string); ok {
		o.ErrorMessage = val
	} else if val, ok := kv["error_message"].(string); ok {
		o.ErrorMessage = &val
	} else {
		if val, ok := kv["error_message"]; ok {
			if val == nil {
				o.ErrorMessage = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.ErrorMessage = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["errored"].(bool); ok {
		o.Errored = val
	} else {
		if val, ok := kv["errored"]; ok {
			if val == nil {
				o.Errored = false
			} else {
				o.Errored = number.ToBoolAny(val)
			}
		}
	}
	if val, ok := kv["ref_id"].(*string); ok {
		o.RefID = val
	} else if val, ok := kv["ref_id"].(string); ok {
		o.RefID = &val
	} else {
		if val, ok := kv["ref_id"]; ok {
			if val == nil {
				o.RefID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.RefID = pstrings.Pointer(fmt.Sprintf("%v", val))
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

// IntegrationInstance The integration instance for a customer
type IntegrationInstance struct {
	// Active If true, the integration is still active
	Active bool `json:"active" codec:"active" bson:"active" yaml:"active" faker:"-"`
	// Config the integration configuration controlled by the integration itself
	Config *string `json:"config,omitempty" codec:"config,omitempty" bson:"config" yaml:"config,omitempty" faker:"-"`
	// CreatedByProfileID The id of the profile for the user that created the integration
	CreatedByProfileID *string `json:"created_by_profile_id,omitempty" codec:"created_by_profile_id,omitempty" bson:"created_by_profile_id" yaml:"created_by_profile_id,omitempty" faker:"-"`
	// CreatedByUserID The id of the user that created the integration
	CreatedByUserID *string `json:"created_by_user_id,omitempty" codec:"created_by_user_id,omitempty" bson:"created_by_user_id" yaml:"created_by_user_id,omitempty" faker:"-"`
	// CreatedDate when the integration was created
	CreatedDate IntegrationInstanceCreatedDate `json:"created_date" codec:"created_date" bson:"created_date" yaml:"created_date" faker:"-"`
	// CreatedAt the date the record was created in Epoch time
	CreatedAt int64 `json:"created_ts" codec:"created_ts" bson:"created_ts" yaml:"created_ts" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// EntityErrors export status and error per entity in the integration
	EntityErrors []IntegrationInstanceEntityErrors `json:"entity_errors" codec:"entity_errors" bson:"entity_errors" yaml:"entity_errors" faker:"-"`
	// ErrorMessage The error message from an export run
	ErrorMessage *string `json:"error_message,omitempty" codec:"error_message,omitempty" bson:"error_message" yaml:"error_message,omitempty" faker:"-"`
	// Errored If authorization failed by the agent or any other error
	Errored *bool `json:"errored,omitempty" codec:"errored,omitempty" bson:"errored" yaml:"errored,omitempty" faker:"-"`
	// ExportAcknowledged Set to true an export has been recieved by the agent.
	ExportAcknowledged *bool `json:"export_acknowledged,omitempty" codec:"export_acknowledged,omitempty" bson:"export_acknowledged" yaml:"export_acknowledged,omitempty" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationID The unique id for the integration
	IntegrationID string `json:"integration_id" codec:"integration_id" bson:"integration_id" yaml:"integration_id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// Interval the interval in milliseconds for how often an export job is scheduled
	Interval int64 `json:"interval" codec:"interval" bson:"interval" yaml:"interval" faker:"-"`
	// LastExportCompletedDate when the export response was received
	LastExportCompletedDate IntegrationInstanceLastExportCompletedDate `json:"last_export_completed_date" codec:"last_export_completed_date" bson:"last_export_completed_date" yaml:"last_export_completed_date" faker:"-"`
	// LastExportRequestedDate when the export request was made
	LastExportRequestedDate IntegrationInstanceLastExportRequestedDate `json:"last_export_requested_date" codec:"last_export_requested_date" bson:"last_export_requested_date" yaml:"last_export_requested_date" faker:"-"`
	// LastProcessingDate when the data was last processed
	LastProcessingDate IntegrationInstanceLastProcessingDate `json:"last_processing_date" codec:"last_processing_date" bson:"last_processing_date" yaml:"last_processing_date" faker:"-"`
	// Location The location of this integration (on-premise / private or cloud)
	Location IntegrationInstanceLocation `json:"location" codec:"location" bson:"location" yaml:"location" faker:"-"`
	// Name The user friendly name of the integration
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
	// Paused true if the agent is paused and should not start new scheduled jobs
	Paused bool `json:"paused" codec:"paused" bson:"paused" yaml:"paused" faker:"-"`
	// Processed If the integration has been processed at least once
	Processed *bool `json:"processed,omitempty" codec:"processed,omitempty" bson:"processed" yaml:"processed,omitempty" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// State the current state of the integration
	State IntegrationInstanceState `json:"state" codec:"state" bson:"state" yaml:"state" faker:"-"`
	// Throttled Set to true when integration is throttled.
	Throttled *bool `json:"throttled,omitempty" codec:"throttled,omitempty" bson:"throttled" yaml:"throttled,omitempty" faker:"-"`
	// ThrottledUntil After throttling integration, we set this field for estimated resume date.
	ThrottledUntil *IntegrationInstanceThrottledUntil `json:"throttled_until,omitempty" codec:"throttled_until,omitempty" bson:"throttled_until" yaml:"throttled_until,omitempty" faker:"-"`
	// UpdatedAt the date the record was updated in Epoch time
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// Webhooks for any webhooks installed at the integration instance
	Webhooks []IntegrationInstanceWebhooks `json:"webhooks" codec:"webhooks" bson:"webhooks" yaml:"webhooks" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*IntegrationInstance)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*IntegrationInstance)(nil)

func toIntegrationInstanceObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationInstance:
		return v.ToMap()

	case IntegrationInstanceCreatedDate:
		return v.ToMap()

	case []IntegrationInstanceEntityErrors:
		arr := make([]interface{}, 0)
		for _, i := range v {
			arr = append(arr, i.ToMap())
		}
		return arr

	case IntegrationInstanceLastExportCompletedDate:
		return v.ToMap()

	case IntegrationInstanceLastExportRequestedDate:
		return v.ToMap()

	case IntegrationInstanceLastProcessingDate:
		return v.ToMap()

	case IntegrationInstanceLocation:
		return v.String()

	case IntegrationInstanceState:
		return v.String()

	case *IntegrationInstanceThrottledUntil:
		return v.ToMap()

	case []IntegrationInstanceWebhooks:
		arr := make([]interface{}, 0)
		for _, i := range v {
			arr = append(arr, i.ToMap())
		}
		return arr

	default:
		return o
	}
}

// String returns a string representation of IntegrationInstance
func (o *IntegrationInstance) String() string {
	return fmt.Sprintf("agent.IntegrationInstance<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *IntegrationInstance) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *IntegrationInstance) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *IntegrationInstance) GetTableName() string {
	return IntegrationInstanceTable.String()
}

// GetModelName returns the name of the model
func (o *IntegrationInstance) GetModelName() datamodel.ModelNameType {
	return IntegrationInstanceModelName
}

// NewIntegrationInstanceID provides a template for generating an ID field for IntegrationInstance
func NewIntegrationInstanceID(customerID string) string {
	return hash.Values(customerID, datetime.EpochNow(), randomString(64))
}

func (o *IntegrationInstance) setDefaults(frommap bool) {
	if o.EntityErrors == nil {
		o.EntityErrors = make([]IntegrationInstanceEntityErrors, 0)
	}
	if o.ThrottledUntil == nil {
		o.ThrottledUntil = &IntegrationInstanceThrottledUntil{}
	}
	if o.Webhooks == nil {
		o.Webhooks = make([]IntegrationInstanceWebhooks, 0)
	}

	if o.ID == "" {
		// set the id from the spec provided in the model
		o.ID = hash.Values(o.CustomerID, datetime.EpochNow(), randomString(64))
	}

	if o.Errored == nil {
		var v bool
		o.Errored = &v
	}

	if o.ExportAcknowledged == nil {
		var v bool
		o.ExportAcknowledged = &v
	}

	{

	}

	if o.Processed == nil {
		var v bool
		o.Processed = &v
	}

	if o.Throttled == nil {
		var v bool
		o.Throttled = &v
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *IntegrationInstance) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *IntegrationInstance) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *IntegrationInstance) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *IntegrationInstance) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *IntegrationInstance) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *IntegrationInstance) IsMutable() bool {
	return true
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *IntegrationInstance) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *IntegrationInstance) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *IntegrationInstance) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = IntegrationInstanceModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *IntegrationInstance) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *IntegrationInstance) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of IntegrationInstance
func (o *IntegrationInstance) Clone() datamodel.Model {
	c := new(IntegrationInstance)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *IntegrationInstance) Anon() datamodel.Model {
	c := new(IntegrationInstance)
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
func (o *IntegrationInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *IntegrationInstance) UnmarshalJSON(data []byte) error {
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
func (o *IntegrationInstance) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *IntegrationInstance) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two IntegrationInstance objects are equal
func (o *IntegrationInstance) IsEqual(other *IntegrationInstance) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *IntegrationInstance) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"active":                     toIntegrationInstanceObject(o.Active, false),
		"config":                     toIntegrationInstanceObject(o.Config, true),
		"created_by_profile_id":      toIntegrationInstanceObject(o.CreatedByProfileID, true),
		"created_by_user_id":         toIntegrationInstanceObject(o.CreatedByUserID, true),
		"created_date":               toIntegrationInstanceObject(o.CreatedDate, false),
		"created_ts":                 toIntegrationInstanceObject(o.CreatedAt, false),
		"customer_id":                toIntegrationInstanceObject(o.CustomerID, false),
		"entity_errors":              toIntegrationInstanceObject(o.EntityErrors, false),
		"error_message":              toIntegrationInstanceObject(o.ErrorMessage, true),
		"errored":                    toIntegrationInstanceObject(o.Errored, true),
		"export_acknowledged":        toIntegrationInstanceObject(o.ExportAcknowledged, true),
		"id":                         toIntegrationInstanceObject(o.ID, false),
		"integration_id":             toIntegrationInstanceObject(o.IntegrationID, false),
		"integration_instance_id":    toIntegrationInstanceObject(o.IntegrationInstanceID, true),
		"interval":                   toIntegrationInstanceObject(o.Interval, false),
		"last_export_completed_date": toIntegrationInstanceObject(o.LastExportCompletedDate, false),
		"last_export_requested_date": toIntegrationInstanceObject(o.LastExportRequestedDate, false),
		"last_processing_date":       toIntegrationInstanceObject(o.LastProcessingDate, false),

		"location":  o.Location.String(),
		"name":      toIntegrationInstanceObject(o.Name, false),
		"paused":    toIntegrationInstanceObject(o.Paused, false),
		"processed": toIntegrationInstanceObject(o.Processed, true),
		"ref_id":    toIntegrationInstanceObject(o.RefID, false),
		"ref_type":  toIntegrationInstanceObject(o.RefType, false),

		"state":           o.State.String(),
		"throttled":       toIntegrationInstanceObject(o.Throttled, true),
		"throttled_until": toIntegrationInstanceObject(o.ThrottledUntil, true),
		"updated_ts":      toIntegrationInstanceObject(o.UpdatedAt, false),
		"webhooks":        toIntegrationInstanceObject(o.Webhooks, false),
		"hashcode":        toIntegrationInstanceObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationInstance) FromMap(kv map[string]interface{}) {

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
	if val, ok := kv["config"].(*string); ok {
		o.Config = val
	} else if val, ok := kv["config"].(string); ok {
		o.Config = &val
	} else {
		if val, ok := kv["config"]; ok {
			if val == nil {
				o.Config = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Config = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["created_by_profile_id"].(*string); ok {
		o.CreatedByProfileID = val
	} else if val, ok := kv["created_by_profile_id"].(string); ok {
		o.CreatedByProfileID = &val
	} else {
		if val, ok := kv["created_by_profile_id"]; ok {
			if val == nil {
				o.CreatedByProfileID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.CreatedByProfileID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["created_by_user_id"].(*string); ok {
		o.CreatedByUserID = val
	} else if val, ok := kv["created_by_user_id"].(string); ok {
		o.CreatedByUserID = &val
	} else {
		if val, ok := kv["created_by_user_id"]; ok {
			if val == nil {
				o.CreatedByUserID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.CreatedByUserID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if val, ok := kv["created_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.CreatedDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationInstanceCreatedDate); ok {
			// struct
			o.CreatedDate = sv
		} else if sp, ok := val.(*IntegrationInstanceCreatedDate); ok {
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

	if o == nil {

		o.EntityErrors = make([]IntegrationInstanceEntityErrors, 0)

	}
	if val, ok := kv["entity_errors"]; ok {
		if sv, ok := val.([]IntegrationInstanceEntityErrors); ok {
			o.EntityErrors = sv
		} else if sp, ok := val.([]*IntegrationInstanceEntityErrors); ok {
			o.EntityErrors = o.EntityErrors[:0]
			for _, e := range sp {
				o.EntityErrors = append(o.EntityErrors, *e)
			}
		} else if a, ok := val.(primitive.A); ok {
			for _, ae := range a {
				if av, ok := ae.(IntegrationInstanceEntityErrors); ok {
					o.EntityErrors = append(o.EntityErrors, av)
				} else if av, ok := ae.(primitive.M); ok {
					var fm IntegrationInstanceEntityErrors
					fm.FromMap(av)
					o.EntityErrors = append(o.EntityErrors, fm)
				} else {
					b, _ := json.Marshal(ae)
					bkv := make(map[string]interface{})
					json.Unmarshal(b, &bkv)
					var av IntegrationInstanceEntityErrors
					av.FromMap(bkv)
					o.EntityErrors = append(o.EntityErrors, av)
				}
			}
		} else if arr, ok := val.([]interface{}); ok {
			for _, item := range arr {
				if r, ok := item.(IntegrationInstanceEntityErrors); ok {
					o.EntityErrors = append(o.EntityErrors, r)
				} else if r, ok := item.(map[string]interface{}); ok {
					var fm IntegrationInstanceEntityErrors
					fm.FromMap(r)
					o.EntityErrors = append(o.EntityErrors, fm)
				} else if r, ok := item.(primitive.M); ok {
					fm := IntegrationInstanceEntityErrors{}
					fm.FromMap(r)
					o.EntityErrors = append(o.EntityErrors, fm)
				}
			}
		} else {
			arr := reflect.ValueOf(val)
			if arr.Kind() == reflect.Slice {
				for i := 0; i < arr.Len(); i++ {
					item := arr.Index(i)
					if item.CanAddr() {
						v := item.Addr().MethodByName("ToMap")
						if !v.IsNil() {
							m := v.Call([]reflect.Value{})
							var fm IntegrationInstanceEntityErrors
							fm.FromMap(m[0].Interface().(map[string]interface{}))
							o.EntityErrors = append(o.EntityErrors, fm)
						}
					}
				}
			}
		}
	}

	if val, ok := kv["error_message"].(*string); ok {
		o.ErrorMessage = val
	} else if val, ok := kv["error_message"].(string); ok {
		o.ErrorMessage = &val
	} else {
		if val, ok := kv["error_message"]; ok {
			if val == nil {
				o.ErrorMessage = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.ErrorMessage = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["errored"].(*bool); ok {
		o.Errored = val
	} else if val, ok := kv["errored"].(bool); ok {
		o.Errored = &val
	} else {
		if val, ok := kv["errored"]; ok {
			if val == nil {
				o.Errored = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Errored = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}
	if val, ok := kv["export_acknowledged"].(*bool); ok {
		o.ExportAcknowledged = val
	} else if val, ok := kv["export_acknowledged"].(bool); ok {
		o.ExportAcknowledged = &val
	} else {
		if val, ok := kv["export_acknowledged"]; ok {
			if val == nil {
				o.ExportAcknowledged = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.ExportAcknowledged = number.BoolPointer(number.ToBoolAny(val))
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
	if val, ok := kv["integration_id"].(string); ok {
		o.IntegrationID = val
	} else {
		if val, ok := kv["integration_id"]; ok {
			if val == nil {
				o.IntegrationID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.IntegrationID = fmt.Sprintf("%v", val)
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
	if val, ok := kv["interval"].(int64); ok {
		o.Interval = val
	} else {
		if val, ok := kv["interval"]; ok {
			if val == nil {
				o.Interval = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Interval = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["last_export_completed_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastExportCompletedDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationInstanceLastExportCompletedDate); ok {
			// struct
			o.LastExportCompletedDate = sv
		} else if sp, ok := val.(*IntegrationInstanceLastExportCompletedDate); ok {
			// struct pointer
			o.LastExportCompletedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.LastExportCompletedDate.Epoch = dt.Epoch
			o.LastExportCompletedDate.Rfc3339 = dt.Rfc3339
			o.LastExportCompletedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
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
		} else if sv, ok := val.(IntegrationInstanceLastExportRequestedDate); ok {
			// struct
			o.LastExportRequestedDate = sv
		} else if sp, ok := val.(*IntegrationInstanceLastExportRequestedDate); ok {
			// struct pointer
			o.LastExportRequestedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.LastExportRequestedDate.Epoch = dt.Epoch
			o.LastExportRequestedDate.Rfc3339 = dt.Rfc3339
			o.LastExportRequestedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
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

	if val, ok := kv["last_processing_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastProcessingDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationInstanceLastProcessingDate); ok {
			// struct
			o.LastProcessingDate = sv
		} else if sp, ok := val.(*IntegrationInstanceLastProcessingDate); ok {
			// struct pointer
			o.LastProcessingDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.LastProcessingDate.Epoch = dt.Epoch
			o.LastProcessingDate.Rfc3339 = dt.Rfc3339
			o.LastProcessingDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
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

	if val, ok := kv["location"].(IntegrationInstanceLocation); ok {
		o.Location = val
	} else {
		if em, ok := kv["location"].(map[string]interface{}); ok {

			ev := em["agent.location"].(string)
			switch ev {
			case "private", "PRIVATE":
				o.Location = 0
			case "cloud", "CLOUD":
				o.Location = 1
			}
		}
		if em, ok := kv["location"].(string); ok {
			switch em {
			case "private", "PRIVATE":
				o.Location = 0
			case "cloud", "CLOUD":
				o.Location = 1
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
	if val, ok := kv["paused"].(bool); ok {
		o.Paused = val
	} else {
		if val, ok := kv["paused"]; ok {
			if val == nil {
				o.Paused = false
			} else {
				o.Paused = number.ToBoolAny(val)
			}
		}
	}
	if val, ok := kv["processed"].(*bool); ok {
		o.Processed = val
	} else if val, ok := kv["processed"].(bool); ok {
		o.Processed = &val
	} else {
		if val, ok := kv["processed"]; ok {
			if val == nil {
				o.Processed = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Processed = number.BoolPointer(number.ToBoolAny(val))
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
	if val, ok := kv["state"].(IntegrationInstanceState); ok {
		o.State = val
	} else {
		if em, ok := kv["state"].(map[string]interface{}); ok {

			ev := em["agent.state"].(string)
			switch ev {
			case "idle", "IDLE":
				o.State = 0
			case "exporting", "EXPORTING":
				o.State = 1
			}
		}
		if em, ok := kv["state"].(string); ok {
			switch em {
			case "idle", "IDLE":
				o.State = 0
			case "exporting", "EXPORTING":
				o.State = 1
			}
		}
	}
	if val, ok := kv["throttled"].(*bool); ok {
		o.Throttled = val
	} else if val, ok := kv["throttled"].(bool); ok {
		o.Throttled = &val
	} else {
		if val, ok := kv["throttled"]; ok {
			if val == nil {
				o.Throttled = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Throttled = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}

	if o.ThrottledUntil == nil {
		o.ThrottledUntil = &IntegrationInstanceThrottledUntil{}
	}

	if val, ok := kv["throttled_until"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.ThrottledUntil.FromMap(kv)
		} else if sv, ok := val.(IntegrationInstanceThrottledUntil); ok {
			// struct
			o.ThrottledUntil = &sv
		} else if sp, ok := val.(*IntegrationInstanceThrottledUntil); ok {
			// struct pointer
			o.ThrottledUntil = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.ThrottledUntil.Epoch = dt.Epoch
			o.ThrottledUntil.Rfc3339 = dt.Rfc3339
			o.ThrottledUntil.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.ThrottledUntil.Epoch = dt.Epoch
			o.ThrottledUntil.Rfc3339 = dt.Rfc3339
			o.ThrottledUntil.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.ThrottledUntil.Epoch = dt.Epoch
				o.ThrottledUntil.Rfc3339 = dt.Rfc3339
				o.ThrottledUntil.Offset = dt.Offset
			}
		}
	} else {
		o.ThrottledUntil.FromMap(map[string]interface{}{})
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

	if o == nil {

		o.Webhooks = make([]IntegrationInstanceWebhooks, 0)

	}
	if val, ok := kv["webhooks"]; ok {
		if sv, ok := val.([]IntegrationInstanceWebhooks); ok {
			o.Webhooks = sv
		} else if sp, ok := val.([]*IntegrationInstanceWebhooks); ok {
			o.Webhooks = o.Webhooks[:0]
			for _, e := range sp {
				o.Webhooks = append(o.Webhooks, *e)
			}
		} else if a, ok := val.(primitive.A); ok {
			for _, ae := range a {
				if av, ok := ae.(IntegrationInstanceWebhooks); ok {
					o.Webhooks = append(o.Webhooks, av)
				} else if av, ok := ae.(primitive.M); ok {
					var fm IntegrationInstanceWebhooks
					fm.FromMap(av)
					o.Webhooks = append(o.Webhooks, fm)
				} else {
					b, _ := json.Marshal(ae)
					bkv := make(map[string]interface{})
					json.Unmarshal(b, &bkv)
					var av IntegrationInstanceWebhooks
					av.FromMap(bkv)
					o.Webhooks = append(o.Webhooks, av)
				}
			}
		} else if arr, ok := val.([]interface{}); ok {
			for _, item := range arr {
				if r, ok := item.(IntegrationInstanceWebhooks); ok {
					o.Webhooks = append(o.Webhooks, r)
				} else if r, ok := item.(map[string]interface{}); ok {
					var fm IntegrationInstanceWebhooks
					fm.FromMap(r)
					o.Webhooks = append(o.Webhooks, fm)
				} else if r, ok := item.(primitive.M); ok {
					fm := IntegrationInstanceWebhooks{}
					fm.FromMap(r)
					o.Webhooks = append(o.Webhooks, fm)
				}
			}
		} else {
			arr := reflect.ValueOf(val)
			if arr.Kind() == reflect.Slice {
				for i := 0; i < arr.Len(); i++ {
					item := arr.Index(i)
					if item.CanAddr() {
						v := item.Addr().MethodByName("ToMap")
						if !v.IsNil() {
							m := v.Call([]reflect.Value{})
							var fm IntegrationInstanceWebhooks
							fm.FromMap(m[0].Interface().(map[string]interface{}))
							o.Webhooks = append(o.Webhooks, fm)
						}
					}
				}
			}
		}
	}

	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *IntegrationInstance) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Active)
	args = append(args, o.Config)
	args = append(args, o.CreatedByProfileID)
	args = append(args, o.CreatedByUserID)
	args = append(args, o.CreatedDate)
	args = append(args, o.CreatedAt)
	args = append(args, o.CustomerID)
	args = append(args, o.EntityErrors)
	args = append(args, o.ErrorMessage)
	args = append(args, o.Errored)
	args = append(args, o.ExportAcknowledged)
	args = append(args, o.ID)
	args = append(args, o.IntegrationID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.Interval)
	args = append(args, o.LastExportCompletedDate)
	args = append(args, o.LastExportRequestedDate)
	args = append(args, o.LastProcessingDate)
	args = append(args, o.Location)
	args = append(args, o.Name)
	args = append(args, o.Paused)
	args = append(args, o.Processed)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.State)
	args = append(args, o.Throttled)
	args = append(args, o.ThrottledUntil)
	args = append(args, o.UpdatedAt)
	args = append(args, o.Webhooks)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

func getIntegrationInstanceQueryFields() string {
	var sb strings.Builder

	// scalar
	sb.WriteString("\t\t\tactive\n")
	// scalar
	sb.WriteString("\t\t\tconfig\n")
	// scalar
	sb.WriteString("\t\t\tcreated_by_profile_id\n")
	// scalar
	sb.WriteString("\t\t\tcreated_by_user_id\n")
	// object with fields
	sb.WriteString("\t\t\tcreated_date {\n")

	// scalar
	sb.WriteString("\t\t\tepoch\n")
	// scalar
	sb.WriteString("\t\t\toffset\n")
	// scalar
	sb.WriteString("\t\t\trfc3339\n")
	sb.WriteString("\t\t\t}\n")
	// scalar
	sb.WriteString("\t\t\tcreated_ts\n")
	// scalar
	sb.WriteString("\t\t\tcustomer_id\n")
	// object with fields
	sb.WriteString("\t\t\tentity_errors {\n")

	// scalar
	sb.WriteString("\t\t\terror\n")
	// scalar
	sb.WriteString("\t\t\tid\n")
	// scalar
	sb.WriteString("\t\t\tref_id\n")
	sb.WriteString("\t\t\t}\n")
	// scalar
	sb.WriteString("\t\t\terror_message\n")
	// scalar
	sb.WriteString("\t\t\terrored\n")
	// scalar
	sb.WriteString("\t\t\texport_acknowledged\n")
	// id
	sb.WriteString("\t\t\t_id\n")
	// scalar
	sb.WriteString("\t\t\tintegration_id\n")
	// scalar
	sb.WriteString("\t\t\tintegration_instance_id\n")
	// scalar
	sb.WriteString("\t\t\tinterval\n")
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
	sb.WriteString("\t\t\tlast_processing_date {\n")

	// scalar
	sb.WriteString("\t\t\tepoch\n")
	// scalar
	sb.WriteString("\t\t\toffset\n")
	// scalar
	sb.WriteString("\t\t\trfc3339\n")
	sb.WriteString("\t\t\t}\n")
	// scalar
	sb.WriteString("\t\t\tlocation\n")
	// scalar
	sb.WriteString("\t\t\tname\n")
	// scalar
	sb.WriteString("\t\t\tpaused\n")
	// scalar
	sb.WriteString("\t\t\tprocessed\n")
	// scalar
	sb.WriteString("\t\t\tref_id\n")
	// scalar
	sb.WriteString("\t\t\tref_type\n")
	// scalar
	sb.WriteString("\t\t\tstate\n")
	// scalar
	sb.WriteString("\t\t\tthrottled\n")
	// object with fields
	sb.WriteString("\t\t\tthrottled_until {\n")

	// scalar
	sb.WriteString("\t\t\tepoch\n")
	// scalar
	sb.WriteString("\t\t\toffset\n")
	// scalar
	sb.WriteString("\t\t\trfc3339\n")
	sb.WriteString("\t\t\t}\n")
	// scalar
	sb.WriteString("\t\t\tupdated_ts\n")
	// object with fields
	sb.WriteString("\t\t\twebhooks {\n")

	// scalar
	sb.WriteString("\t\t\tenabled\n")
	// scalar
	sb.WriteString("\t\t\terror_message\n")
	// scalar
	sb.WriteString("\t\t\terrored\n")
	// scalar
	sb.WriteString("\t\t\tref_id\n")
	// scalar
	sb.WriteString("\t\t\turl\n")
	sb.WriteString("\t\t\t}\n")
	return sb.String()
}

// IntegrationInstancePageInfo is a grapqhl PageInfo
type IntegrationInstancePageInfo struct {
	StartCursor     string `json:"startCursor,omitempty"`
	EndCursor       string `json:"endCursor,omitempty"`
	HasNextPage     bool   `json:"hasNextPage,omitempty"`
	HasPreviousPage bool   `json:"hasPreviousPage,omitempty"`
}

// IntegrationInstanceConnection is a grapqhl connection
type IntegrationInstanceConnection struct {
	Edges      []*IntegrationInstanceEdge  `json:"edges,omitempty"`
	PageInfo   IntegrationInstancePageInfo `json:"pageInfo,omitempty"`
	TotalCount *int64                      `json:"totalCount,omitempty"`
}

// IntegrationInstanceEdge is a grapqhl edge
type IntegrationInstanceEdge struct {
	Node *IntegrationInstance `json:"node,omitempty"`
}

// QueryManyIntegrationInstanceNode is a grapqhl query many node
type QueryManyIntegrationInstanceNode struct {
	Object *IntegrationInstanceConnection `json:"IntegrationInstances,omitempty"`
}

// QueryManyIntegrationInstanceData is a grapqhl query many data node
type QueryManyIntegrationInstanceData struct {
	Data *QueryManyIntegrationInstanceNode `json:"agent,omitempty"`
}

// QueryOneIntegrationInstanceNode is a grapqhl query one node
type QueryOneIntegrationInstanceNode struct {
	Object *IntegrationInstance `json:"IntegrationInstance,omitempty"`
}

// QueryOneIntegrationInstanceData is a grapqhl query one data node
type QueryOneIntegrationInstanceData struct {
	Data *QueryOneIntegrationInstanceNode `json:"agent,omitempty"`
}

// IntegrationInstanceQuery is query struct
type IntegrationInstanceQuery struct {
	Filters []string      `json:"filters,omitempty"`
	Params  []interface{} `json:"params,omitempty"`
}

// IntegrationInstanceOrder is the order direction
type IntegrationInstanceOrder *string

var (
	// IntegrationInstanceOrderDesc is descending
	IntegrationInstanceOrderDesc IntegrationInstanceOrder = pstrings.Pointer("DESC")
	// IntegrationInstanceOrderAsc is ascending
	IntegrationInstanceOrderAsc IntegrationInstanceOrder = pstrings.Pointer("ASC")
)

// IntegrationInstanceQueryInput is query input struct
type IntegrationInstanceQueryInput struct {
	First   *int64                    `json:"first,omitempty"`
	Last    *int64                    `json:"last,omitempty"`
	Before  *string                   `json:"before,omitempty"`
	After   *string                   `json:"after,omitempty"`
	Query   *IntegrationInstanceQuery `json:"query,omitempty"`
	OrderBy *string                   `json:"orderBy,omitempty"`
	Order   IntegrationInstanceOrder  `json:"order,omitempty"`
}

// NewIntegrationInstanceQuery is a convenience for building a *IntegrationInstanceQuery
func NewIntegrationInstanceQuery(params ...interface{}) *IntegrationInstanceQueryInput {
	if len(params)%2 != 0 {
		panic("incorrect number of arguments passed")
	}
	q := &IntegrationInstanceQuery{
		Filters: make([]string, 0),
		Params:  make([]interface{}, 0),
	}
	for i := 0; i < len(params); i += 2 {
		q.Filters = append(q.Filters, params[i].(string))
		q.Params = append(q.Params, params[i+1])
	}
	return &IntegrationInstanceQueryInput{
		Query: q,
	}
}

// FindIntegrationInstance will query an IntegrationInstance by id
func FindIntegrationInstance(client graphql.Client, id string) (*IntegrationInstance, error) {
	variables := make(graphql.Variables)
	variables["id"] = id
	var sb strings.Builder
	sb.WriteString("query IntegrationInstanceQuery($id: ID) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tIntegrationInstance(_id: $id) {\n")
	sb.WriteString(getIntegrationInstanceQueryFields())
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res QueryOneIntegrationInstanceData
	if err := client.Query(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	if res.Data != nil {
		return res.Data.Object, nil
	}
	return nil, nil
}

// FindIntegrationInstances will query for any IntegrationInstances matching the query
func FindIntegrationInstances(client graphql.Client, input *IntegrationInstanceQueryInput) (*IntegrationInstanceConnection, error) {
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
	sb.WriteString("query IntegrationInstanceQueryMany($first: Int, $last: Int, $before: Cursor, $after: Cursor, $query: QueryInput, $order: SortOrderEnum, $orderBy: AgentIntegrationInstanceColumnEnum) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tIntegrationInstances(first: $first last: $last before: $before after: $after query: $query orderBy: $orderBy order: $order) {\n")
	sb.WriteString("\t\t\tpageInfo {\n")
	sb.WriteString("\t\t\t\tstartCursor\n")
	sb.WriteString("\t\t\t\tendCursor\n")
	sb.WriteString("\t\t\t\thasNextPage\n")
	sb.WriteString("\t\t\t\thasPreviousPage\n")
	sb.WriteString("\t\t\t}\n")
	sb.WriteString("\t\t\ttotalCount\n")
	sb.WriteString("\t\t\tedges {\n")
	sb.WriteString("\t\t\t\tnode {\n")
	sb.WriteString(getIntegrationInstanceQueryFields())
	sb.WriteString("\t\t\t\t}\n")
	sb.WriteString("\t\t\t}\n")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res QueryManyIntegrationInstanceData
	if err := client.Query(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	return res.Data.Object, nil
}

// FindIntegrationInstancesPaginatedCallback is a callback function for handling each page
type FindIntegrationInstancesPaginatedCallback func(conn *IntegrationInstanceConnection) (bool, error)

// FindIntegrationInstancesPaginated will query for any IntegrationInstances matching the query and return each page callback
func FindIntegrationInstancesPaginated(client graphql.Client, query *IntegrationInstanceQuery, pageSize int64, callback FindIntegrationInstancesPaginatedCallback) error {
	input := &IntegrationInstanceQueryInput{
		First: &pageSize,
		Query: query,
	}
	for {
		res, err := FindIntegrationInstances(client, input)
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

// UpdateIntegrationInstanceNode is a grapqhl update node
type UpdateIntegrationInstanceNode struct {
	Object *IntegrationInstance `json:"updateIntegrationInstance,omitempty"`
}

// UpdateIntegrationInstanceData is a grapqhl update data node
type UpdateIntegrationInstanceData struct {
	Data *UpdateIntegrationInstanceNode `json:"agent,omitempty"`
}

// ExecIntegrationInstanceUpdateMutation returns a graphql update mutation result for IntegrationInstance
func ExecIntegrationInstanceUpdateMutation(client graphql.Client, id string, input graphql.Variables, upsert bool) (*IntegrationInstance, error) {
	variables := make(graphql.Variables)
	variables["id"] = id
	variables["upsert"] = upsert
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation IntegrationInstanceUpdateMutation($id: String, $input: UpdateAgentIntegrationInstanceInput, $upsert: Boolean) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tupdateIntegrationInstance(_id: $id, input: $input, upsert: $upsert) {\n")
	sb.WriteString(getIntegrationInstanceQueryFields())
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateIntegrationInstanceData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	return res.Data.Object, nil
}

// ExecIntegrationInstanceSilentUpdateMutation returns a graphql update mutation result for IntegrationInstance
func ExecIntegrationInstanceSilentUpdateMutation(client graphql.Client, id string, input graphql.Variables, upsert bool) error {
	variables := make(graphql.Variables)
	variables["id"] = id
	variables["upsert"] = upsert
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation IntegrationInstanceUpdateMutation($id: String, $input: UpdateAgentIntegrationInstanceInput, $upsert: Boolean) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tupdateIntegrationInstance(_id: $id, input: $input, upsert: $upsert) {\n")
	sb.WriteString("\t\t\t_id")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateIntegrationInstanceData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return err
	}
	return nil
}

func CreateIntegrationInstance(client graphql.Client, model IntegrationInstance) error {
	variables := make(graphql.Variables)
	input := model.ToMap()
	delete(input, "hashcode")
	delete(input, "customer_id")
	delete(input, "created_ts")
	delete(input, "id")
	input["_id"] = model.GetID()
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation CreateIntegrationInstance($input: CreateAgentIntegrationInstanceInput!) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tcreateIntegrationInstance(input: $input) {\n")
	sb.WriteString("\t\t\t_id")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateIntegrationInstanceData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return err
	}
	return nil
}

// IntegrationInstancePartial is a partial struct for upsert mutations for IntegrationInstance
type IntegrationInstancePartial struct {
	// Active If true, the integration is still active
	Active *bool `json:"active,omitempty"`
	// Config the integration configuration controlled by the integration itself
	Config *string `json:"config,omitempty"`
	// CreatedByProfileID The id of the profile for the user that created the integration
	CreatedByProfileID *string `json:"created_by_profile_id,omitempty"`
	// CreatedByUserID The id of the user that created the integration
	CreatedByUserID *string `json:"created_by_user_id,omitempty"`
	// CreatedDate when the integration was created
	CreatedDate *IntegrationInstanceCreatedDate `json:"created_date,omitempty"`
	// EntityErrors export status and error per entity in the integration
	EntityErrors []IntegrationInstanceEntityErrors `json:"entity_errors,omitempty"`
	// ErrorMessage The error message from an export run
	ErrorMessage *string `json:"error_message,omitempty"`
	// Errored If authorization failed by the agent or any other error
	Errored *bool `json:"errored,omitempty"`
	// ExportAcknowledged Set to true an export has been recieved by the agent.
	ExportAcknowledged *bool `json:"export_acknowledged,omitempty"`
	// IntegrationID The unique id for the integration
	IntegrationID *string `json:"integration_id,omitempty"`
	// Interval the interval in milliseconds for how often an export job is scheduled
	Interval *int64 `json:"interval,omitempty"`
	// LastExportCompletedDate when the export response was received
	LastExportCompletedDate *IntegrationInstanceLastExportCompletedDate `json:"last_export_completed_date,omitempty"`
	// LastExportRequestedDate when the export request was made
	LastExportRequestedDate *IntegrationInstanceLastExportRequestedDate `json:"last_export_requested_date,omitempty"`
	// LastProcessingDate when the data was last processed
	LastProcessingDate *IntegrationInstanceLastProcessingDate `json:"last_processing_date,omitempty"`
	// Location The location of this integration (on-premise / private or cloud)
	Location *IntegrationInstanceLocation `json:"location,omitempty"`
	// Name The user friendly name of the integration
	Name *string `json:"name,omitempty"`
	// Paused true if the agent is paused and should not start new scheduled jobs
	Paused *bool `json:"paused,omitempty"`
	// Processed If the integration has been processed at least once
	Processed *bool `json:"processed,omitempty"`
	// State the current state of the integration
	State *IntegrationInstanceState `json:"state,omitempty"`
	// Throttled Set to true when integration is throttled.
	Throttled *bool `json:"throttled,omitempty"`
	// ThrottledUntil After throttling integration, we set this field for estimated resume date.
	ThrottledUntil *IntegrationInstanceThrottledUntil `json:"throttled_until,omitempty"`
	// Webhooks for any webhooks installed at the integration instance
	Webhooks []IntegrationInstanceWebhooks `json:"webhooks,omitempty"`
}

var _ datamodel.PartialModel = (*IntegrationInstancePartial)(nil)

// GetModelName returns the name of the model
func (o *IntegrationInstancePartial) GetModelName() datamodel.ModelNameType {
	return IntegrationInstanceModelName
}

// ToMap returns the object as a map
func (o *IntegrationInstancePartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"active":                     toIntegrationInstanceObject(o.Active, true),
		"config":                     toIntegrationInstanceObject(o.Config, true),
		"created_by_profile_id":      toIntegrationInstanceObject(o.CreatedByProfileID, true),
		"created_by_user_id":         toIntegrationInstanceObject(o.CreatedByUserID, true),
		"created_date":               toIntegrationInstanceObject(o.CreatedDate, true),
		"entity_errors":              toIntegrationInstanceObject(o.EntityErrors, true),
		"error_message":              toIntegrationInstanceObject(o.ErrorMessage, true),
		"errored":                    toIntegrationInstanceObject(o.Errored, true),
		"export_acknowledged":        toIntegrationInstanceObject(o.ExportAcknowledged, true),
		"integration_id":             toIntegrationInstanceObject(o.IntegrationID, true),
		"interval":                   toIntegrationInstanceObject(o.Interval, true),
		"last_export_completed_date": toIntegrationInstanceObject(o.LastExportCompletedDate, true),
		"last_export_requested_date": toIntegrationInstanceObject(o.LastExportRequestedDate, true),
		"last_processing_date":       toIntegrationInstanceObject(o.LastProcessingDate, true),

		"location":  toIntegrationInstanceLocationEnum(o.Location),
		"name":      toIntegrationInstanceObject(o.Name, true),
		"paused":    toIntegrationInstanceObject(o.Paused, true),
		"processed": toIntegrationInstanceObject(o.Processed, true),

		"state":           toIntegrationInstanceStateEnum(o.State),
		"throttled":       toIntegrationInstanceObject(o.Throttled, true),
		"throttled_until": toIntegrationInstanceObject(o.ThrottledUntil, true),
		"webhooks":        toIntegrationInstanceObject(o.Webhooks, true),
	}
	for k, v := range kv {
		if v == nil || reflect.ValueOf(v).IsZero() {
			delete(kv, k)
		} else {
			if k == "created_date" {
				if dt, ok := v.(*IntegrationInstanceCreatedDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}

			if k == "entity_errors" {
				if arr, ok := v.([]IntegrationInstanceEntityErrors); ok {
					if len(arr) == 0 {
						delete(kv, k)
					}
				}
			}
			if k == "last_export_completed_date" {
				if dt, ok := v.(*IntegrationInstanceLastExportCompletedDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}
			if k == "last_export_requested_date" {
				if dt, ok := v.(*IntegrationInstanceLastExportRequestedDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}
			if k == "last_processing_date" {
				if dt, ok := v.(*IntegrationInstanceLastProcessingDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}
			if k == "throttled_until" {
				if dt, ok := v.(*IntegrationInstanceThrottledUntil); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}

			if k == "webhooks" {
				if arr, ok := v.([]IntegrationInstanceWebhooks); ok {
					if len(arr) == 0 {
						delete(kv, k)
					}
				}
			}
		}
	}
	return kv
}

// Stringify returns the object in JSON format as a string
func (o *IntegrationInstancePartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *IntegrationInstancePartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *IntegrationInstancePartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *IntegrationInstancePartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *IntegrationInstancePartial) FromMap(kv map[string]interface{}) {
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
	if val, ok := kv["config"].(*string); ok {
		o.Config = val
	} else if val, ok := kv["config"].(string); ok {
		o.Config = &val
	} else {
		if val, ok := kv["config"]; ok {
			if val == nil {
				o.Config = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Config = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["created_by_profile_id"].(*string); ok {
		o.CreatedByProfileID = val
	} else if val, ok := kv["created_by_profile_id"].(string); ok {
		o.CreatedByProfileID = &val
	} else {
		if val, ok := kv["created_by_profile_id"]; ok {
			if val == nil {
				o.CreatedByProfileID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.CreatedByProfileID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["created_by_user_id"].(*string); ok {
		o.CreatedByUserID = val
	} else if val, ok := kv["created_by_user_id"].(string); ok {
		o.CreatedByUserID = &val
	} else {
		if val, ok := kv["created_by_user_id"]; ok {
			if val == nil {
				o.CreatedByUserID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.CreatedByUserID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if o.CreatedDate == nil {
		o.CreatedDate = &IntegrationInstanceCreatedDate{}
	}

	if val, ok := kv["created_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.CreatedDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationInstanceCreatedDate); ok {
			// struct
			o.CreatedDate = &sv
		} else if sp, ok := val.(*IntegrationInstanceCreatedDate); ok {
			// struct pointer
			o.CreatedDate = sp
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

	if o == nil {

		o.EntityErrors = make([]IntegrationInstanceEntityErrors, 0)

	}
	if val, ok := kv["entity_errors"]; ok {
		if sv, ok := val.([]IntegrationInstanceEntityErrors); ok {
			o.EntityErrors = sv
		} else if sp, ok := val.([]*IntegrationInstanceEntityErrors); ok {
			o.EntityErrors = o.EntityErrors[:0]
			for _, e := range sp {
				o.EntityErrors = append(o.EntityErrors, *e)
			}
		} else if a, ok := val.(primitive.A); ok {
			for _, ae := range a {
				if av, ok := ae.(IntegrationInstanceEntityErrors); ok {
					o.EntityErrors = append(o.EntityErrors, av)
				} else if av, ok := ae.(primitive.M); ok {
					var fm IntegrationInstanceEntityErrors
					fm.FromMap(av)
					o.EntityErrors = append(o.EntityErrors, fm)
				} else {
					b, _ := json.Marshal(ae)
					bkv := make(map[string]interface{})
					json.Unmarshal(b, &bkv)
					var av IntegrationInstanceEntityErrors
					av.FromMap(bkv)
					o.EntityErrors = append(o.EntityErrors, av)
				}
			}
		} else if arr, ok := val.([]interface{}); ok {
			for _, item := range arr {
				if r, ok := item.(IntegrationInstanceEntityErrors); ok {
					o.EntityErrors = append(o.EntityErrors, r)
				} else if r, ok := item.(map[string]interface{}); ok {
					var fm IntegrationInstanceEntityErrors
					fm.FromMap(r)
					o.EntityErrors = append(o.EntityErrors, fm)
				} else if r, ok := item.(primitive.M); ok {
					fm := IntegrationInstanceEntityErrors{}
					fm.FromMap(r)
					o.EntityErrors = append(o.EntityErrors, fm)
				}
			}
		} else {
			arr := reflect.ValueOf(val)
			if arr.Kind() == reflect.Slice {
				for i := 0; i < arr.Len(); i++ {
					item := arr.Index(i)
					if item.CanAddr() {
						v := item.Addr().MethodByName("ToMap")
						if !v.IsNil() {
							m := v.Call([]reflect.Value{})
							var fm IntegrationInstanceEntityErrors
							fm.FromMap(m[0].Interface().(map[string]interface{}))
							o.EntityErrors = append(o.EntityErrors, fm)
						}
					}
				}
			}
		}
	}

	if val, ok := kv["error_message"].(*string); ok {
		o.ErrorMessage = val
	} else if val, ok := kv["error_message"].(string); ok {
		o.ErrorMessage = &val
	} else {
		if val, ok := kv["error_message"]; ok {
			if val == nil {
				o.ErrorMessage = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.ErrorMessage = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["errored"].(*bool); ok {
		o.Errored = val
	} else if val, ok := kv["errored"].(bool); ok {
		o.Errored = &val
	} else {
		if val, ok := kv["errored"]; ok {
			if val == nil {
				o.Errored = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Errored = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}
	if val, ok := kv["export_acknowledged"].(*bool); ok {
		o.ExportAcknowledged = val
	} else if val, ok := kv["export_acknowledged"].(bool); ok {
		o.ExportAcknowledged = &val
	} else {
		if val, ok := kv["export_acknowledged"]; ok {
			if val == nil {
				o.ExportAcknowledged = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.ExportAcknowledged = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}
	if val, ok := kv["integration_id"].(*string); ok {
		o.IntegrationID = val
	} else if val, ok := kv["integration_id"].(string); ok {
		o.IntegrationID = &val
	} else {
		if val, ok := kv["integration_id"]; ok {
			if val == nil {
				o.IntegrationID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.IntegrationID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["interval"].(*int64); ok {
		o.Interval = val
	} else if val, ok := kv["interval"].(int64); ok {
		o.Interval = &val
	} else {
		if val, ok := kv["interval"]; ok {
			if val == nil {
				o.Interval = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["long"]
				}
				o.Interval = number.Int64Pointer(number.ToInt64Any(val))
			}
		}
	}

	if o.LastExportCompletedDate == nil {
		o.LastExportCompletedDate = &IntegrationInstanceLastExportCompletedDate{}
	}

	if val, ok := kv["last_export_completed_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastExportCompletedDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationInstanceLastExportCompletedDate); ok {
			// struct
			o.LastExportCompletedDate = &sv
		} else if sp, ok := val.(*IntegrationInstanceLastExportCompletedDate); ok {
			// struct pointer
			o.LastExportCompletedDate = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.LastExportCompletedDate.Epoch = dt.Epoch
			o.LastExportCompletedDate.Rfc3339 = dt.Rfc3339
			o.LastExportCompletedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
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
		o.LastExportRequestedDate = &IntegrationInstanceLastExportRequestedDate{}
	}

	if val, ok := kv["last_export_requested_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastExportRequestedDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationInstanceLastExportRequestedDate); ok {
			// struct
			o.LastExportRequestedDate = &sv
		} else if sp, ok := val.(*IntegrationInstanceLastExportRequestedDate); ok {
			// struct pointer
			o.LastExportRequestedDate = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.LastExportRequestedDate.Epoch = dt.Epoch
			o.LastExportRequestedDate.Rfc3339 = dt.Rfc3339
			o.LastExportRequestedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
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

	if o.LastProcessingDate == nil {
		o.LastProcessingDate = &IntegrationInstanceLastProcessingDate{}
	}

	if val, ok := kv["last_processing_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastProcessingDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationInstanceLastProcessingDate); ok {
			// struct
			o.LastProcessingDate = &sv
		} else if sp, ok := val.(*IntegrationInstanceLastProcessingDate); ok {
			// struct pointer
			o.LastProcessingDate = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.LastProcessingDate.Epoch = dt.Epoch
			o.LastProcessingDate.Rfc3339 = dt.Rfc3339
			o.LastProcessingDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
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

	if val, ok := kv["location"].(*IntegrationInstanceLocation); ok {
		o.Location = val
	} else if val, ok := kv["location"].(IntegrationInstanceLocation); ok {
		o.Location = &val
	} else {
		if val, ok := kv["location"]; ok {
			if val == nil {
				o.Location = toIntegrationInstanceLocationPointer(0)
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["IntegrationInstanceLocation"]
				}
				// this is an enum pointer
				if em, ok := val.(string); ok {
					switch em {
					case "private", "PRIVATE":
						o.Location = toIntegrationInstanceLocationPointer(0)
					case "cloud", "CLOUD":
						o.Location = toIntegrationInstanceLocationPointer(1)
					}
				}
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
	if val, ok := kv["paused"].(*bool); ok {
		o.Paused = val
	} else if val, ok := kv["paused"].(bool); ok {
		o.Paused = &val
	} else {
		if val, ok := kv["paused"]; ok {
			if val == nil {
				o.Paused = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Paused = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}
	if val, ok := kv["processed"].(*bool); ok {
		o.Processed = val
	} else if val, ok := kv["processed"].(bool); ok {
		o.Processed = &val
	} else {
		if val, ok := kv["processed"]; ok {
			if val == nil {
				o.Processed = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Processed = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}
	if val, ok := kv["state"].(*IntegrationInstanceState); ok {
		o.State = val
	} else if val, ok := kv["state"].(IntegrationInstanceState); ok {
		o.State = &val
	} else {
		if val, ok := kv["state"]; ok {
			if val == nil {
				o.State = toIntegrationInstanceStatePointer(0)
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["IntegrationInstanceState"]
				}
				// this is an enum pointer
				if em, ok := val.(string); ok {
					switch em {
					case "idle", "IDLE":
						o.State = toIntegrationInstanceStatePointer(0)
					case "exporting", "EXPORTING":
						o.State = toIntegrationInstanceStatePointer(1)
					}
				}
			}
		}
	}
	if val, ok := kv["throttled"].(*bool); ok {
		o.Throttled = val
	} else if val, ok := kv["throttled"].(bool); ok {
		o.Throttled = &val
	} else {
		if val, ok := kv["throttled"]; ok {
			if val == nil {
				o.Throttled = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Throttled = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}

	if o.ThrottledUntil == nil {
		o.ThrottledUntil = &IntegrationInstanceThrottledUntil{}
	}

	if val, ok := kv["throttled_until"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.ThrottledUntil.FromMap(kv)
		} else if sv, ok := val.(IntegrationInstanceThrottledUntil); ok {
			// struct
			o.ThrottledUntil = &sv
		} else if sp, ok := val.(*IntegrationInstanceThrottledUntil); ok {
			// struct pointer
			o.ThrottledUntil = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.ThrottledUntil.Epoch = dt.Epoch
			o.ThrottledUntil.Rfc3339 = dt.Rfc3339
			o.ThrottledUntil.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.ThrottledUntil.Epoch = dt.Epoch
			o.ThrottledUntil.Rfc3339 = dt.Rfc3339
			o.ThrottledUntil.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.ThrottledUntil.Epoch = dt.Epoch
				o.ThrottledUntil.Rfc3339 = dt.Rfc3339
				o.ThrottledUntil.Offset = dt.Offset
			}
		}
	} else {
		o.ThrottledUntil.FromMap(map[string]interface{}{})
	}

	if o == nil {

		o.Webhooks = make([]IntegrationInstanceWebhooks, 0)

	}
	if val, ok := kv["webhooks"]; ok {
		if sv, ok := val.([]IntegrationInstanceWebhooks); ok {
			o.Webhooks = sv
		} else if sp, ok := val.([]*IntegrationInstanceWebhooks); ok {
			o.Webhooks = o.Webhooks[:0]
			for _, e := range sp {
				o.Webhooks = append(o.Webhooks, *e)
			}
		} else if a, ok := val.(primitive.A); ok {
			for _, ae := range a {
				if av, ok := ae.(IntegrationInstanceWebhooks); ok {
					o.Webhooks = append(o.Webhooks, av)
				} else if av, ok := ae.(primitive.M); ok {
					var fm IntegrationInstanceWebhooks
					fm.FromMap(av)
					o.Webhooks = append(o.Webhooks, fm)
				} else {
					b, _ := json.Marshal(ae)
					bkv := make(map[string]interface{})
					json.Unmarshal(b, &bkv)
					var av IntegrationInstanceWebhooks
					av.FromMap(bkv)
					o.Webhooks = append(o.Webhooks, av)
				}
			}
		} else if arr, ok := val.([]interface{}); ok {
			for _, item := range arr {
				if r, ok := item.(IntegrationInstanceWebhooks); ok {
					o.Webhooks = append(o.Webhooks, r)
				} else if r, ok := item.(map[string]interface{}); ok {
					var fm IntegrationInstanceWebhooks
					fm.FromMap(r)
					o.Webhooks = append(o.Webhooks, fm)
				} else if r, ok := item.(primitive.M); ok {
					fm := IntegrationInstanceWebhooks{}
					fm.FromMap(r)
					o.Webhooks = append(o.Webhooks, fm)
				}
			}
		} else {
			arr := reflect.ValueOf(val)
			if arr.Kind() == reflect.Slice {
				for i := 0; i < arr.Len(); i++ {
					item := arr.Index(i)
					if item.CanAddr() {
						v := item.Addr().MethodByName("ToMap")
						if !v.IsNil() {
							m := v.Call([]reflect.Value{})
							var fm IntegrationInstanceWebhooks
							fm.FromMap(m[0].Interface().(map[string]interface{}))
							o.Webhooks = append(o.Webhooks, fm)
						}
					}
				}
			}
		}
	}

	o.setDefaults(false)
}
