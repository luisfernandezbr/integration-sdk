// DO NOT EDIT -- generated code

// Package calendar - public calendar data models
package calendar

import (
	"encoding/json"
	"fmt"
	"reflect"
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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (

	// EventTable is the default table name
	EventTable datamodel.ModelNameType = "calendar_event"

	// EventModelName is the model name
	EventModelName datamodel.ModelNameType = "calendar.Event"
)

const (
	// EventModelBusyColumn is the column json value busy
	EventModelBusyColumn = "busy"
	// EventModelCalendarIDColumn is the column json value calendar_id
	EventModelCalendarIDColumn = "calendar_id"
	// EventModelCustomerIDColumn is the column json value customer_id
	EventModelCustomerIDColumn = "customer_id"
	// EventModelDescriptionColumn is the column json value description
	EventModelDescriptionColumn = "description"
	// EventModelEndDateColumn is the column json value end_date
	EventModelEndDateColumn = "end_date"
	// EventModelEndDateEpochColumn is the column json value epoch
	EventModelEndDateEpochColumn = "epoch"
	// EventModelEndDateOffsetColumn is the column json value offset
	EventModelEndDateOffsetColumn = "offset"
	// EventModelEndDateRfc3339Column is the column json value rfc3339
	EventModelEndDateRfc3339Column = "rfc3339"
	// EventModelIDColumn is the column json value id
	EventModelIDColumn = "id"
	// EventModelLocationColumn is the column json value location
	EventModelLocationColumn = "location"
	// EventModelLocationDetailsColumn is the column json value details
	EventModelLocationDetailsColumn = "details"
	// EventModelLocationNameColumn is the column json value name
	EventModelLocationNameColumn = "name"
	// EventModelLocationURLColumn is the column json value url
	EventModelLocationURLColumn = "url"
	// EventModelNameColumn is the column json value name
	EventModelNameColumn = "name"
	// EventModelOwnerRefIDColumn is the column json value owner_ref_id
	EventModelOwnerRefIDColumn = "owner_ref_id"
	// EventModelParticipantsColumn is the column json value participants
	EventModelParticipantsColumn = "participants"
	// EventModelParticipantsStatusColumn is the column json value status
	EventModelParticipantsStatusColumn = "status"
	// EventModelParticipantsUserRefIDColumn is the column json value user_ref_id
	EventModelParticipantsUserRefIDColumn = "user_ref_id"
	// EventModelRefIDColumn is the column json value ref_id
	EventModelRefIDColumn = "ref_id"
	// EventModelRefTypeColumn is the column json value ref_type
	EventModelRefTypeColumn = "ref_type"
	// EventModelStartDateColumn is the column json value start_date
	EventModelStartDateColumn = "start_date"
	// EventModelStartDateEpochColumn is the column json value epoch
	EventModelStartDateEpochColumn = "epoch"
	// EventModelStartDateOffsetColumn is the column json value offset
	EventModelStartDateOffsetColumn = "offset"
	// EventModelStartDateRfc3339Column is the column json value rfc3339
	EventModelStartDateRfc3339Column = "rfc3339"
	// EventModelStatusColumn is the column json value status
	EventModelStatusColumn = "status"
)

// EventEndDate represents the object structure for end_date
type EventEndDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toEventEndDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *EventEndDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *EventEndDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toEventEndDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toEventEndDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toEventEndDateObject(o.Rfc3339, false),
	}
}

func (o *EventEndDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *EventEndDate) FromMap(kv map[string]interface{}) {

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

// EventLocation represents the object structure for location
type EventLocation struct {
	// Details raw data from the location property the calendar
	Details string `json:"details" codec:"details" bson:"details" yaml:"details" faker:"-"`
	// Name the name of the place
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
	// URL link to the meeting if any
	URL string `json:"url" codec:"url" bson:"url" yaml:"url" faker:"-"`
}

func toEventLocationObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *EventLocation:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *EventLocation) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Details raw data from the location property the calendar
		"details": toEventLocationObject(o.Details, false),
		// Name the name of the place
		"name": toEventLocationObject(o.Name, false),
		// URL link to the meeting if any
		"url": toEventLocationObject(o.URL, false),
	}
}

func (o *EventLocation) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *EventLocation) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["details"].(string); ok {
		o.Details = val
	} else {
		if val, ok := kv["details"]; ok {
			if val == nil {
				o.Details = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Details = fmt.Sprintf("%v", val)
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
	o.setDefaults(false)
}

// EventParticipantsStatus is the enumeration type for status
type EventParticipantsStatus int32

// UnmarshalBSONValue for unmarshaling value
func (v *EventParticipantsStatus) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = EventParticipantsStatus(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "UNKNOWN":
			*v = EventParticipantsStatus(0)
		case "MAYBE":
			*v = EventParticipantsStatus(1)
		case "GOING":
			*v = EventParticipantsStatus(2)
		case "NOT_GOING":
			*v = EventParticipantsStatus(3)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v EventParticipantsStatus) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "UNKNOWN":
		v = 0
	case "MAYBE":
		v = 1
	case "GOING":
		v = 2
	case "NOT_GOING":
		v = 3
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v EventParticipantsStatus) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("UNKNOWN")
	case 1:
		return json.Marshal("MAYBE")
	case 2:
		return json.Marshal("GOING")
	case 3:
		return json.Marshal("NOT_GOING")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for ParticipantsStatus
func (v EventParticipantsStatus) String() string {
	switch int32(v) {
	case 0:
		return "UNKNOWN"
	case 1:
		return "MAYBE"
	case 2:
		return "GOING"
	case 3:
		return "NOT_GOING"
	}
	return "unset"
}

const (
	// EventParticipantsStatusUnknown is the enumeration value for unknown
	EventParticipantsStatusUnknown EventParticipantsStatus = 0
	// EventParticipantsStatusMaybe is the enumeration value for maybe
	EventParticipantsStatusMaybe EventParticipantsStatus = 1
	// EventParticipantsStatusGoing is the enumeration value for going
	EventParticipantsStatusGoing EventParticipantsStatus = 2
	// EventParticipantsStatusNotGoing is the enumeration value for not_going
	EventParticipantsStatusNotGoing EventParticipantsStatus = 3
)

// EventParticipants represents the object structure for participants
type EventParticipants struct {
	// Status the status of the participant
	Status EventParticipantsStatus `json:"status" codec:"status" bson:"status" yaml:"status" faker:"-"`
	// UserRefID the user ref_id of the participant
	UserRefID string `json:"user_ref_id" codec:"user_ref_id" bson:"user_ref_id" yaml:"user_ref_id" faker:"-"`
}

func toEventParticipantsObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *EventParticipants:
		return v.ToMap()

	case EventParticipantsStatus:
		return v.String()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *EventParticipants) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Status the status of the participant
		"status": toEventParticipantsObject(o.Status, false),
		// UserRefID the user ref_id of the participant
		"user_ref_id": toEventParticipantsObject(o.UserRefID, false),
	}
}

func (o *EventParticipants) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *EventParticipants) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["status"].(EventParticipantsStatus); ok {
		o.Status = val
	} else {
		if em, ok := kv["status"].(map[string]interface{}); ok {

			ev := em["calendar.status"].(string)
			switch ev {
			case "unknown", "UNKNOWN":
				o.Status = 0
			case "maybe", "MAYBE":
				o.Status = 1
			case "going", "GOING":
				o.Status = 2
			case "not_going", "NOT_GOING":
				o.Status = 3
			}
		}
		if em, ok := kv["status"].(string); ok {
			switch em {
			case "unknown", "UNKNOWN":
				o.Status = 0
			case "maybe", "MAYBE":
				o.Status = 1
			case "going", "GOING":
				o.Status = 2
			case "not_going", "NOT_GOING":
				o.Status = 3
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

// EventStartDate represents the object structure for start_date
type EventStartDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toEventStartDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *EventStartDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *EventStartDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toEventStartDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toEventStartDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toEventStartDateObject(o.Rfc3339, false),
	}
}

func (o *EventStartDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *EventStartDate) FromMap(kv map[string]interface{}) {

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

// EventStatus is the enumeration type for status
type EventStatus int32

// UnmarshalBSONValue for unmarshaling value
func (v *EventStatus) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = EventStatus(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "CONFIRMED":
			*v = EventStatus(0)
		case "TENTATIVE":
			*v = EventStatus(1)
		case "CANCELLED":
			*v = EventStatus(2)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v EventStatus) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "CONFIRMED":
		v = 0
	case "TENTATIVE":
		v = 1
	case "CANCELLED":
		v = 2
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v EventStatus) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("CONFIRMED")
	case 1:
		return json.Marshal("TENTATIVE")
	case 2:
		return json.Marshal("CANCELLED")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for Status
func (v EventStatus) String() string {
	switch int32(v) {
	case 0:
		return "CONFIRMED"
	case 1:
		return "TENTATIVE"
	case 2:
		return "CANCELLED"
	}
	return "unset"
}

const (
	// EventStatusConfirmed is the enumeration value for confirmed
	EventStatusConfirmed EventStatus = 0
	// EventStatusTentative is the enumeration value for tentative
	EventStatusTentative EventStatus = 1
	// EventStatusCancelled is the enumeration value for cancelled
	EventStatusCancelled EventStatus = 2
)

// Event event on a calendar
type Event struct {
	// Busy true if the user is marked as busy in this event
	Busy bool `json:"busy" codec:"busy" bson:"busy" yaml:"busy" faker:"-"`
	// CalendarID unique project id this event belongs to
	CalendarID string `json:"calendar_id" codec:"calendar_id" bson:"calendar_id" yaml:"calendar_id" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Description the description of the event
	Description string `json:"description" codec:"description" bson:"description" yaml:"description" faker:"-"`
	// EndDate the event end date
	EndDate EventEndDate `json:"end_date" codec:"end_date" bson:"end_date" yaml:"end_date" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// Location location of the event, could be a place or a link to a meeting
	Location EventLocation `json:"location" codec:"location" bson:"location" yaml:"location" faker:"-"`
	// Name the name of the event
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
	// OwnerRefID owner ref_id of the event
	OwnerRefID string `json:"owner_ref_id" codec:"owner_ref_id" bson:"owner_ref_id" yaml:"owner_ref_id" faker:"-"`
	// Participants participants of the event
	Participants []EventParticipants `json:"participants" codec:"participants" bson:"participants" yaml:"participants" faker:"-"`
	// RefID the calendar ID
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the record type
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// StartDate the event start date
	StartDate EventStartDate `json:"start_date" codec:"start_date" bson:"start_date" yaml:"start_date" faker:"-"`
	// Status the status of the user in this event
	Status EventStatus `json:"status" codec:"status" bson:"status" yaml:"status" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Event)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*Event)(nil)

func toEventObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *Event:
		return v.ToMap()

	case EventEndDate:
		return v.ToMap()

	case EventLocation:
		return v.ToMap()

	case []EventParticipants:
		arr := make([]interface{}, 0)
		for _, i := range v {
			arr = append(arr, i.ToMap())
		}
		return arr

	case EventStartDate:
		return v.ToMap()

	case EventStatus:
		return v.String()
	default:
		return o
	}
}

// String returns a string representation of Event
func (o *Event) String() string {
	return fmt.Sprintf("calendar.Event<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Event) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *Event) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *Event) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *Event) GetModelName() datamodel.ModelNameType {
	return EventModelName
}

// NewEventID provides a template for generating an ID field for Event
func NewEventID(customerID string, refType string, refID string) string {
	return hash.Values("Event", customerID, refType, refID)
}

func (o *Event) setDefaults(frommap bool) {
	if o.Participants == nil {
		o.Participants = make([]EventParticipants, 0)
	}

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Event", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Event) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Event) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Event) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *Event) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Event) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *Event) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Event) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Event) IsEvented() bool {
	return false
}

// GetTopicConfig returns the topic config object
func (o *Event) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *Event) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of Event
func (o *Event) Clone() datamodel.Model {
	c := new(Event)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Event) Anon() datamodel.Model {
	c := new(Event)
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
func (o *Event) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Event) UnmarshalJSON(data []byte) error {
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
func (o *Event) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *Event) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two Event objects are equal
func (o *Event) IsEqual(other *Event) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Event) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"busy":         toEventObject(o.Busy, false),
		"calendar_id":  toEventObject(o.CalendarID, false),
		"customer_id":  toEventObject(o.CustomerID, false),
		"description":  toEventObject(o.Description, false),
		"end_date":     toEventObject(o.EndDate, false),
		"id":           toEventObject(o.ID, false),
		"location":     toEventObject(o.Location, false),
		"name":         toEventObject(o.Name, false),
		"owner_ref_id": toEventObject(o.OwnerRefID, false),
		"participants": toEventObject(o.Participants, false),
		"ref_id":       toEventObject(o.RefID, false),
		"ref_type":     toEventObject(o.RefType, false),
		"start_date":   toEventObject(o.StartDate, false),

		"status":   o.Status.String(),
		"hashcode": toEventObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *Event) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["busy"].(bool); ok {
		o.Busy = val
	} else {
		if val, ok := kv["busy"]; ok {
			if val == nil {
				o.Busy = false
			} else {
				o.Busy = number.ToBoolAny(val)
			}
		}
	}

	if val, ok := kv["calendar_id"].(string); ok {
		o.CalendarID = val
	} else {
		if val, ok := kv["calendar_id"]; ok {
			if val == nil {
				o.CalendarID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.CalendarID = fmt.Sprintf("%v", val)
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

	if val, ok := kv["description"].(string); ok {
		o.Description = val
	} else {
		if val, ok := kv["description"]; ok {
			if val == nil {
				o.Description = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Description = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["end_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.EndDate.FromMap(kv)
		} else if sv, ok := val.(EventEndDate); ok {
			// struct
			o.EndDate = sv
		} else if sp, ok := val.(*EventEndDate); ok {
			// struct pointer
			o.EndDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.EndDate.Epoch = dt.Epoch
			o.EndDate.Rfc3339 = dt.Rfc3339
			o.EndDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.EndDate.Epoch = dt.Epoch
			o.EndDate.Rfc3339 = dt.Rfc3339
			o.EndDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.EndDate.Epoch = dt.Epoch
				o.EndDate.Rfc3339 = dt.Rfc3339
				o.EndDate.Offset = dt.Offset
			}
		}
	} else {
		o.EndDate.FromMap(map[string]interface{}{})
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

	if val, ok := kv["location"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.Location.FromMap(kv)
		} else if sv, ok := val.(EventLocation); ok {
			// struct
			o.Location = sv
		} else if sp, ok := val.(*EventLocation); ok {
			// struct pointer
			o.Location = *sp
		}
	} else {
		o.Location.FromMap(map[string]interface{}{})
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

	if val, ok := kv["owner_ref_id"].(string); ok {
		o.OwnerRefID = val
	} else {
		if val, ok := kv["owner_ref_id"]; ok {
			if val == nil {
				o.OwnerRefID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.OwnerRefID = fmt.Sprintf("%v", val)
			}
		}
	}

	if o == nil {

		o.Participants = make([]EventParticipants, 0)

	}
	if val, ok := kv["participants"]; ok {
		if sv, ok := val.([]EventParticipants); ok {
			o.Participants = sv
		} else if sp, ok := val.([]*EventParticipants); ok {
			o.Participants = o.Participants[:0]
			for _, e := range sp {
				o.Participants = append(o.Participants, *e)
			}
		} else if a, ok := val.(primitive.A); ok {
			for _, ae := range a {
				if av, ok := ae.(EventParticipants); ok {
					o.Participants = append(o.Participants, av)
				} else if av, ok := ae.(primitive.M); ok {
					var fm EventParticipants
					fm.FromMap(av)
					o.Participants = append(o.Participants, fm)
				} else {
					b, _ := json.Marshal(ae)
					bkv := make(map[string]interface{})
					json.Unmarshal(b, &bkv)
					var av EventParticipants
					av.FromMap(bkv)
					o.Participants = append(o.Participants, av)
				}
			}
		} else if arr, ok := val.([]interface{}); ok {
			for _, item := range arr {
				if r, ok := item.(EventParticipants); ok {
					o.Participants = append(o.Participants, r)
				} else if r, ok := item.(map[string]interface{}); ok {
					var fm EventParticipants
					fm.FromMap(r)
					o.Participants = append(o.Participants, fm)
				} else if r, ok := item.(primitive.M); ok {
					fm := EventParticipants{}
					fm.FromMap(r)
					o.Participants = append(o.Participants, fm)
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
							var fm EventParticipants
							fm.FromMap(m[0].Interface().(map[string]interface{}))
							o.Participants = append(o.Participants, fm)
						}
					}
				}
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

	if val, ok := kv["start_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.StartDate.FromMap(kv)
		} else if sv, ok := val.(EventStartDate); ok {
			// struct
			o.StartDate = sv
		} else if sp, ok := val.(*EventStartDate); ok {
			// struct pointer
			o.StartDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.StartDate.Epoch = dt.Epoch
			o.StartDate.Rfc3339 = dt.Rfc3339
			o.StartDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.StartDate.Epoch = dt.Epoch
			o.StartDate.Rfc3339 = dt.Rfc3339
			o.StartDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.StartDate.Epoch = dt.Epoch
				o.StartDate.Rfc3339 = dt.Rfc3339
				o.StartDate.Offset = dt.Offset
			}
		}
	} else {
		o.StartDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["status"].(EventStatus); ok {
		o.Status = val
	} else {
		if em, ok := kv["status"].(map[string]interface{}); ok {

			ev := em["calendar.status"].(string)
			switch ev {
			case "confirmed", "CONFIRMED":
				o.Status = 0
			case "tentative", "TENTATIVE":
				o.Status = 1
			case "cancelled", "CANCELLED":
				o.Status = 2
			}
		}
		if em, ok := kv["status"].(string); ok {
			switch em {
			case "confirmed", "CONFIRMED":
				o.Status = 0
			case "tentative", "TENTATIVE":
				o.Status = 1
			case "cancelled", "CANCELLED":
				o.Status = 2
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *Event) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Busy)
	args = append(args, o.CalendarID)
	args = append(args, o.CustomerID)
	args = append(args, o.Description)
	args = append(args, o.EndDate)
	args = append(args, o.ID)
	args = append(args, o.Location)
	args = append(args, o.Name)
	args = append(args, o.OwnerRefID)
	args = append(args, o.Participants)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.StartDate)
	args = append(args, o.Status)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *Event) GetEventAPIConfig() datamodel.EventAPIConfig {
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
