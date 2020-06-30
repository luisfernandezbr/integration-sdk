// DO NOT EDIT -- generated code

// Package web - the web model
package web

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
)

const (

	// HookTable is the default table name
	HookTable datamodel.ModelNameType = "web_hook"

	// HookModelName is the model name
	HookModelName datamodel.ModelNameType = "web.Hook"
)

const (
	// HookModelDataColumn is the column json value data
	HookModelDataColumn = "data"
	// HookModelHeadersColumn is the column json value headers
	HookModelHeadersColumn = "headers"
	// HookModelIDColumn is the column json value id
	HookModelIDColumn = "id"
	// HookModelReceivedDateColumn is the column json value received_date
	HookModelReceivedDateColumn = "received_date"
	// HookModelReceivedDateEpochColumn is the column json value epoch
	HookModelReceivedDateEpochColumn = "epoch"
	// HookModelReceivedDateOffsetColumn is the column json value offset
	HookModelReceivedDateOffsetColumn = "offset"
	// HookModelReceivedDateRfc3339Column is the column json value rfc3339
	HookModelReceivedDateRfc3339Column = "rfc3339"
	// HookModelSystemColumn is the column json value system
	HookModelSystemColumn = "system"
	// HookModelTokenColumn is the column json value token
	HookModelTokenColumn = "token"
)

// HookReceivedDate represents the object structure for received_date
type HookReceivedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toHookReceivedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *HookReceivedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *HookReceivedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toHookReceivedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toHookReceivedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toHookReceivedDateObject(o.Rfc3339, false),
	}
}

func (o *HookReceivedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *HookReceivedDate) FromMap(kv map[string]interface{}) {

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

// Hook hook is a webhook event which is received from an external source
type Hook struct {
	// Data the webhook data payload base64 encoded
	Data string `json:"data" codec:"data" bson:"data" yaml:"data" faker:"-"`
	// Headers the headers of the incoming webhook
	Headers map[string]string `json:"headers" codec:"headers" bson:"headers" yaml:"headers" faker:"-"`
	// ID the primary key for this model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// ReceivedDate the date when the hook was received
	ReceivedDate HookReceivedDate `json:"received_date" codec:"received_date" bson:"received_date" yaml:"received_date" faker:"-"`
	// System the name of the system sending the hook
	System string `json:"system" codec:"system" bson:"system" yaml:"system" faker:"-"`
	// Token the token part of the url
	Token string `json:"token" codec:"token" bson:"token" yaml:"token" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Hook)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*Hook)(nil)

func toHookObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *Hook:
		return v.ToMap()

	case HookReceivedDate:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of Hook
func (o *Hook) String() string {
	return fmt.Sprintf("web.Hook<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Hook) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *Hook) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *Hook) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *Hook) GetModelName() datamodel.ModelNameType {
	return HookModelName
}

// NewHookID provides a template for generating an ID field for Hook
func NewHookID(System string, Token string, ReceivedDateEpoch int64) string {
	return hash.Values(System, Token, ReceivedDateEpoch)
}

func (o *Hook) setDefaults(frommap bool) {

	if o.ID == "" {
		o.ID = hash.Values(o.System, o.Token, o.ReceivedDate.Epoch)
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// GetID returns the ID for the object
func (o *Hook) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Hook) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Hook) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// IsMaterialized returns true if the model is materialized
func (o *Hook) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *Hook) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Hook) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Hook) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *Hook) SetEventHeaders(kv map[string]string) {
	kv["model"] = HookModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Hook) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// Clone returns an exact copy of Hook
func (o *Hook) Clone() datamodel.Model {
	c := new(Hook)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Hook) Anon() datamodel.Model {
	c := new(Hook)
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
func (o *Hook) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Hook) UnmarshalJSON(data []byte) error {
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
func (o *Hook) Stringify() string {
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *Hook) StringifyPretty() string {
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two Hook objects are equal
func (o *Hook) IsEqual(other *Hook) bool {
	return o.GetID() == other.GetID()
}

// ToMap returns the object as a map
func (o *Hook) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"data":          toHookObject(o.Data, false),
		"headers":       toHookObject(o.Headers, false),
		"id":            toHookObject(o.ID, false),
		"received_date": toHookObject(o.ReceivedDate, false),
		"system":        toHookObject(o.System, false),
		"token":         toHookObject(o.Token, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *Hook) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
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

	if val, ok := kv["received_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.ReceivedDate.FromMap(kv)
		} else if sv, ok := val.(HookReceivedDate); ok {
			// struct
			o.ReceivedDate = sv
		} else if sp, ok := val.(*HookReceivedDate); ok {
			// struct pointer
			o.ReceivedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.ReceivedDate.Epoch = dt.Epoch
			o.ReceivedDate.Rfc3339 = dt.Rfc3339
			o.ReceivedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.ReceivedDate.Epoch = dt.Epoch
			o.ReceivedDate.Rfc3339 = dt.Rfc3339
			o.ReceivedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.ReceivedDate.Epoch = dt.Epoch
				o.ReceivedDate.Rfc3339 = dt.Rfc3339
				o.ReceivedDate.Offset = dt.Offset
			}
		}
	} else {
		o.ReceivedDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["system"].(string); ok {
		o.System = val
	} else {
		if val, ok := kv["system"]; ok {
			if val == nil {
				o.System = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.System = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["token"].(string); ok {
		o.Token = val
	} else {
		if val, ok := kv["token"]; ok {
			if val == nil {
				o.Token = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Token = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}
