// DO NOT EDIT -- generated code

// Package activity - the pipeline models
package activity

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"sync"
	"time"

	"github.com/bxcodec/faker"
	"github.com/linkedin/goavro"
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/eventing"
	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
)

const (
	// ActivityTopic is the default topic name
	ActivityTopic datamodel.TopicNameType = "pipeline_activity_Activity_topic"

	// ActivityStream is the default stream name
	ActivityStream datamodel.TopicNameType = "pipeline_activity_Activity_stream"

	// ActivityTable is the default table name
	ActivityTable datamodel.TopicNameType = "pipeline_activity_Activity"

	// ActivityModelName is the model name
	ActivityModelName datamodel.ModelNameType = "pipeline.activity.Activity"
)

const (
	// ActivityActivityDateColumn is the activity_date column name
	ActivityActivityDateColumn = "activity_date"
	// ActivityActivityDateColumnEpochColumn is the epoch column property of the ActivityDate name
	ActivityActivityDateColumnEpochColumn = "activity_date->epoch"
	// ActivityActivityDateColumnOffsetColumn is the offset column property of the ActivityDate name
	ActivityActivityDateColumnOffsetColumn = "activity_date->offset"
	// ActivityActivityDateColumnRfc3339Column is the rfc3339 column property of the ActivityDate name
	ActivityActivityDateColumnRfc3339Column = "activity_date->rfc3339"
	// ActivityCustomerIDColumn is the customer_id column name
	ActivityCustomerIDColumn = "customer_id"
	// ActivityIDColumn is the id column name
	ActivityIDColumn = "id"
	// ActivityRefIDColumn is the ref_id column name
	ActivityRefIDColumn = "ref_id"
	// ActivityRefTypeColumn is the ref_type column name
	ActivityRefTypeColumn = "ref_type"
	// ActivityTypeColumn is the type column name
	ActivityTypeColumn = "type"
	// ActivityUserColumn is the user column name
	ActivityUserColumn = "user"
	// ActivityUserColumnIDColumn is the id column property of the User name
	ActivityUserColumnIDColumn = "user->id"
	// ActivityUserColumnTeamIDColumn is the team_id column property of the User name
	ActivityUserColumnTeamIDColumn = "user->team_id"
)

// ActivityActivityDate represents the object structure for activity_date
type ActivityActivityDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toActivityActivityDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toActivityActivityDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	// nested => true prefix => ActivityActivityDate name => ActivityActivityDate
	switch v := o.(type) {
	case *ActivityActivityDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *ActivityActivityDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toActivityActivityDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toActivityActivityDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toActivityActivityDateObject(o.Rfc3339, isavro, false, "string"),
	}
}

func (o *ActivityActivityDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ActivityActivityDate) FromMap(kv map[string]interface{}) {

	if val, ok := kv["epoch"].(int64); ok {
		o.Epoch = val
	} else {
		if val, ok := kv["epoch"]; ok {
			if val == nil {
				o.Epoch = number.ToInt64Any(nil)
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
				o.Offset = number.ToInt64Any(nil)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// ActivityType is the enumeration type for type
type ActivityType int32

// String returns the string value for Type
func (v ActivityType) String() string {
	switch int32(v) {
	case 0:
		return "COMMIT"
	case 1:
		return "ISSUE_CREATED"
	case 2:
		return "ISSUE_CLOSED"
	case 3:
		return "ISSUE_STATUS_CHANGE"
	case 4:
		return "ISSUE_COMMENT"
	case 5:
		return "BRANCH_CREATED"
	case 6:
		return "PULL_REQUEST_CREATED"
	case 7:
		return "PULL_REQUEST_MERGED"
	case 8:
		return "PULL_REQUEST_REVIEWED"
	case 9:
		return "BRANCH_MERGED"
	case 10:
		return "UNLINKED_COMMIT"
	}
	return "unset"
}

const (
	// TypeCommit is the enumeration value for commit
	ActivityTypeCommit ActivityType = 0
	// TypeIssueCreated is the enumeration value for issue_created
	ActivityTypeIssueCreated ActivityType = 1
	// TypeIssueClosed is the enumeration value for issue_closed
	ActivityTypeIssueClosed ActivityType = 2
	// TypeIssueStatusChange is the enumeration value for issue_status_change
	ActivityTypeIssueStatusChange ActivityType = 3
	// TypeIssueComment is the enumeration value for issue_comment
	ActivityTypeIssueComment ActivityType = 4
	// TypeBranchCreated is the enumeration value for branch_created
	ActivityTypeBranchCreated ActivityType = 5
	// TypePullRequestCreated is the enumeration value for pull_request_created
	ActivityTypePullRequestCreated ActivityType = 6
	// TypePullRequestMerged is the enumeration value for pull_request_merged
	ActivityTypePullRequestMerged ActivityType = 7
	// TypePullRequestReviewed is the enumeration value for pull_request_reviewed
	ActivityTypePullRequestReviewed ActivityType = 8
	// TypeBranchMerged is the enumeration value for branch_merged
	ActivityTypeBranchMerged ActivityType = 9
	// TypeUnlinkedCommit is the enumeration value for unlinked_commit
	ActivityTypeUnlinkedCommit ActivityType = 10
)

// ActivityUser represents the object structure for user
type ActivityUser struct {
	// ID the corporate user id
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// TeamID the corporate team id
	TeamID string `json:"team_id" bson:"team_id" yaml:"team_id" faker:"-"`
}

func toActivityUserObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toActivityUserObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	// nested => true prefix => ActivityUser name => ActivityUser
	switch v := o.(type) {
	case *ActivityUser:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *ActivityUser) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// ID the corporate user id
		"id": toActivityUserObject(o.ID, isavro, false, "string"),
		// TeamID the corporate team id
		"team_id": toActivityUserObject(o.TeamID, isavro, false, "string"),
	}
}

func (o *ActivityUser) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ActivityUser) FromMap(kv map[string]interface{}) {

	if val, ok := kv["id"].(string); ok {
		o.ID = val
	} else {
		if val, ok := kv["id"]; ok {
			if val == nil {
				o.ID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.ID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["team_id"].(string); ok {
		o.TeamID = val
	} else {
		if val, ok := kv["team_id"]; ok {
			if val == nil {
				o.TeamID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.TeamID = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// Activity activity table is a log of activity accross all integrations
type Activity struct {
	// ActivityDate date object
	ActivityDate ActivityActivityDate `json:"activity_date" bson:"activity_date" yaml:"activity_date" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// Type the type of activity
	Type ActivityType `json:"type" bson:"type" yaml:"type" faker:"-"`
	// User the user related to the activity
	User ActivityUser `json:"user" bson:"user" yaml:"user" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Activity)(nil)

func toActivityObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toActivityObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	// nested => false prefix => Activity name => Activity
	switch v := o.(type) {
	case *Activity:
		return v.ToMap(isavro)

	case ActivityActivityDate:
		return v.ToMap(isavro)

		// is nested enum Type
	case ActivityType:
		return v.String()

	case ActivityUser:
		return v.ToMap(isavro)
	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

// String returns a string representation of Activity
func (o *Activity) String() string {
	return fmt.Sprintf("pipeline.activity.Activity<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Activity) GetTopicName() datamodel.TopicNameType {
	return ActivityTopic
}

// GetModelName returns the name of the model
func (o *Activity) GetModelName() datamodel.ModelNameType {
	return ActivityModelName
}

func (o *Activity) setDefaults(frommap bool) {

	o.GetID()

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Activity) GetID() string {
	if o.ID == "" {
		o.ID = hash.Values(o.CustomerID, o.RefID, o.User.ID, o.Type, o.ActivityDate.Epoch)
	}
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Activity) GetTopicKey() string {
	var i interface{} = o.ID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Activity) GetTimestamp() time.Time {
	var dt interface{} = o.ActivityDate
	switch v := dt.(type) {
	case int64:
		return datetime.DateFromEpoch(v).UTC()
	case string:
		tv, err := datetime.ISODateToTime(v)
		if err != nil {
			panic(err)
		}
		return tv.UTC()
	case time.Time:
		return v.UTC()
	case ActivityActivityDate:
		return datetime.DateFromEpoch(v.Epoch)
	}
	panic("not sure how to handle the date time format for Activity")
}

// GetRefID returns the RefID for the object
func (o *Activity) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Activity) IsMaterialized() bool {
	return true
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Activity) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	idletime, err := time.ParseDuration("1s")
	if err != nil {
		panic(err)
	}
	return &datamodel.ModelMaterializeConfig{
		KeyName:   "id",
		TableName: "pipeline_activity_activity",
		BatchSize: 5000,
		IdleTime:  idletime,
	}
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Activity) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *Activity) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = ActivityModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Activity) GetTopicConfig() *datamodel.ModelTopicConfig {
	retention, err := time.ParseDuration("168h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 168h0m0s. " + err.Error())
	}

	ttl, err := time.ParseDuration("0s")
	if err != nil {
		ttl = 0
	}
	return &datamodel.ModelTopicConfig{
		Key:               "id",
		Timestamp:         "activity_date",
		NumPartitions:     8,
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetStateKey returns a key for use in state store
func (o *Activity) GetStateKey() string {
	key := "id"
	return fmt.Sprintf("%s_%s", key, o.GetID())
}

// GetCustomerID will return the customer_id
func (o *Activity) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of Activity
func (o *Activity) Clone() datamodel.Model {
	c := new(Activity)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Activity) Anon() datamodel.Model {
	c := new(Activity)
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
func (o *Activity) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Activity) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

var cachedCodecActivity *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *Activity) GetAvroCodec() *goavro.Codec {
	if cachedCodecActivity == nil {
		c, err := GetActivityAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecActivity = c
	}
	return cachedCodecActivity
}

// ToAvroBinary returns the data as Avro binary data
func (o *Activity) ToAvroBinary() ([]byte, *goavro.Codec, error) {
	kv := o.ToMap(true)
	jbuf, _ := json.Marshal(kv)
	codec := o.GetAvroCodec()
	native, _, err := codec.NativeFromTextual(jbuf)
	if err != nil {
		return nil, nil, err
	}
	// Convert native Go form to binary Avro data
	buf, err := codec.BinaryFromNative(nil, native)
	return buf, codec, err
}

// FromAvroBinary will convert from Avro binary data into data in this object
func (o *Activity) FromAvroBinary(value []byte) error {
	var nullHeader = []byte{byte(0)}
	// if this still has the schema encoded in the header, move past it to the avro payload
	if bytes.HasPrefix(value, nullHeader) {
		value = value[5:]
	}
	kv, _, err := o.GetAvroCodec().NativeFromBinary(value)
	if err != nil {
		return err
	}
	o.FromMap(kv.(map[string]interface{}))
	return nil
}

// Stringify returns the object in JSON format as a string
func (o *Activity) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Activity objects are equal
func (o *Activity) IsEqual(other *Activity) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Activity) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	o.setDefaults(true)
	return map[string]interface{}{
		"activity_date": toActivityObject(o.ActivityDate, isavro, false, "activity_date"),
		"customer_id":   toActivityObject(o.CustomerID, isavro, false, "string"),
		"id":            toActivityObject(o.ID, isavro, false, "string"),
		"ref_id":        toActivityObject(o.RefID, isavro, false, "string"),
		"ref_type":      toActivityObject(o.RefType, isavro, false, "string"),
		"type":          toActivityObject(o.Type, isavro, false, "type"),
		"user":          toActivityObject(o.User, isavro, false, "user"),
		"hashcode":      toActivityObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *Activity) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["activity_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.ActivityDate.FromMap(kv)
		} else if sv, ok := val.(ActivityActivityDate); ok {
			// struct
			o.ActivityDate = sv
		} else if sp, ok := val.(*ActivityActivityDate); ok {
			// struct pointer
			o.ActivityDate = *sp
		}
	} else {
		o.ActivityDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["customer_id"].(string); ok {
		o.CustomerID = val
	} else {
		if val, ok := kv["customer_id"]; ok {
			if val == nil {
				o.CustomerID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.RefType = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["type"].(ActivityType); ok {
		o.Type = val
	} else {
		if em, ok := kv["type"].(map[string]interface{}); ok {
			ev := em["pipeline.activity.type"].(string)
			switch ev {
			case "commit", "COMMIT":
				o.Type = 0
			case "issue_created", "ISSUE_CREATED":
				o.Type = 1
			case "issue_closed", "ISSUE_CLOSED":
				o.Type = 2
			case "issue_status_change", "ISSUE_STATUS_CHANGE":
				o.Type = 3
			case "issue_comment", "ISSUE_COMMENT":
				o.Type = 4
			case "branch_created", "BRANCH_CREATED":
				o.Type = 5
			case "pull_request_created", "PULL_REQUEST_CREATED":
				o.Type = 6
			case "pull_request_merged", "PULL_REQUEST_MERGED":
				o.Type = 7
			case "pull_request_reviewed", "PULL_REQUEST_REVIEWED":
				o.Type = 8
			case "branch_merged", "BRANCH_MERGED":
				o.Type = 9
			case "unlinked_commit", "UNLINKED_COMMIT":
				o.Type = 10
			}
		}
		if em, ok := kv["type"].(string); ok {
			switch em {
			case "commit", "COMMIT":
				o.Type = 0
			case "issue_created", "ISSUE_CREATED":
				o.Type = 1
			case "issue_closed", "ISSUE_CLOSED":
				o.Type = 2
			case "issue_status_change", "ISSUE_STATUS_CHANGE":
				o.Type = 3
			case "issue_comment", "ISSUE_COMMENT":
				o.Type = 4
			case "branch_created", "BRANCH_CREATED":
				o.Type = 5
			case "pull_request_created", "PULL_REQUEST_CREATED":
				o.Type = 6
			case "pull_request_merged", "PULL_REQUEST_MERGED":
				o.Type = 7
			case "pull_request_reviewed", "PULL_REQUEST_REVIEWED":
				o.Type = 8
			case "branch_merged", "BRANCH_MERGED":
				o.Type = 9
			case "unlinked_commit", "UNLINKED_COMMIT":
				o.Type = 10
			}
		}
	}

	if val, ok := kv["user"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.User.FromMap(kv)
		} else if sv, ok := val.(ActivityUser); ok {
			// struct
			o.User = sv
		} else if sp, ok := val.(*ActivityUser); ok {
			// struct pointer
			o.User = *sp
		}
	} else {
		o.User.FromMap(map[string]interface{}{})
	}

	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *Activity) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.ActivityDate)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.Type)
	args = append(args, o.User)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetActivityAvroSchemaSpec creates the avro schema specification for Activity
func GetActivityAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "pipeline.activity",
		"name":      "Activity",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "hashcode",
				"type": "string",
			},
			map[string]interface{}{
				"name": "activity_date",
				"type": map[string]interface{}{"type": "record", "name": "activity_date", "fields": []interface{}{map[string]interface{}{"type": "long", "name": "epoch", "doc": "the date in epoch format"}, map[string]interface{}{"type": "long", "name": "offset", "doc": "the timezone offset from GMT"}, map[string]interface{}{"type": "string", "name": "rfc3339", "doc": "the date in RFC3339 format"}}, "doc": "date object"},
			},
			map[string]interface{}{
				"name": "customer_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "ref_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "ref_type",
				"type": "string",
			},
			map[string]interface{}{
				"name": "type",
				"type": map[string]interface{}{
					"type":    "enum",
					"name":    "type",
					"symbols": []interface{}{"COMMIT", "ISSUE_CREATED", "ISSUE_CLOSED", "ISSUE_STATUS_CHANGE", "ISSUE_COMMENT", "BRANCH_CREATED", "PULL_REQUEST_CREATED", "PULL_REQUEST_MERGED", "PULL_REQUEST_REVIEWED", "BRANCH_MERGED", "UNLINKED_COMMIT"},
				},
			},
			map[string]interface{}{
				"name": "user",
				"type": map[string]interface{}{"type": "record", "name": "user", "fields": []interface{}{map[string]interface{}{"type": "string", "name": "id", "doc": "the corporate user id"}, map[string]interface{}{"type": "string", "name": "team_id", "doc": "the corporate team id"}}, "doc": "the user related to the activity"},
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *Activity) GetEventAPIConfig() datamodel.EventAPIConfig {
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

// GetActivityAvroSchema creates the avro schema for Activity
func GetActivityAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetActivityAvroSchemaSpec())
}

// TransformActivityFunc is a function for transforming Activity during processing
type TransformActivityFunc func(input *Activity) (*Activity, error)

// NewActivityPipe creates a pipe for processing Activity items
func NewActivityPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformActivityFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewActivityInputStream(input, errors)
	var stream chan Activity
	if len(transforms) > 0 {
		stream = make(chan Activity, 1000)
	} else {
		stream = inch
	}
	outdone := NewActivityOutputStream(output, stream, errors)
	go func() {
		if len(transforms) > 0 {
			var stop bool
			for item := range inch {
				input := &item
				for _, transform := range transforms {
					out, err := transform(input)
					if err != nil {
						stop = true
						errors <- err
						break
					}
					if out == nil {
						input = nil
						break
					} else {
						input = out
					}
				}
				if stop {
					break
				}
				if input != nil {
					stream <- *input
				}
			}
			close(stream)
		}
		<-indone
		<-outdone
		done <- true
	}()
	return done
}

// NewActivityInputStreamDir creates a channel for reading Activity as JSON newlines from a directory of files
func NewActivityInputStreamDir(dir string, errors chan<- error, transforms ...TransformActivityFunc) (chan Activity, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/pipeline.activity/activity\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan Activity)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for activity")
		ch := make(chan Activity)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewActivityInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan Activity)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewActivityInputStreamFile creates an channel for reading Activity as JSON newlines from filename
func NewActivityInputStreamFile(filename string, errors chan<- error, transforms ...TransformActivityFunc) (chan Activity, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan Activity)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	var f io.ReadCloser = of
	if filepath.Ext(filename) == ".gz" {
		gz, err := gzip.NewReader(f)
		if err != nil {
			of.Close()
			errors <- err
			ch := make(chan Activity)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewActivityInputStream(f, errors, transforms...)
}

// NewActivityInputStream creates an channel for reading Activity as JSON newlines from stream
func NewActivityInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformActivityFunc) (chan Activity, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan Activity, 1000)
	go func() {
		defer func() { stream.Close(); close(ch); done <- true }()
		r := bufio.NewReader(stream)
		for {
			buf, err := r.ReadBytes('\n')
			if err != nil {
				if err == io.EOF {
					return
				}
				errors <- err
				return
			}
			var item Activity
			if err := json.Unmarshal(buf, &item); err != nil {
				errors <- err
				return
			}
			in := &item
			var skip bool
			for _, transform := range transforms {
				in, err = transform(in)
				if err != nil {
					errors <- err
					return
				}
				if in == nil {
					skip = true
					break
				}
			}
			if !skip {
				ch <- *in
			}
		}
	}()
	return ch, done
}

// NewActivityOutputStreamDir will output json newlines from channel and save in dir
func NewActivityOutputStreamDir(dir string, ch chan Activity, errors chan<- error, transforms ...TransformActivityFunc) <-chan bool {
	fp := filepath.Join(dir, "/pipeline.activity/activity\\.json(\\.gz)?$")
	os.MkdirAll(filepath.Dir(fp), 0777)
	of, err := os.Create(fp)
	if err != nil {
		errors <- err
		done := make(chan bool, 1)
		done <- true
		return done
	}
	gz, err := gzip.NewWriterLevel(of, gzip.BestCompression)
	if err != nil {
		errors <- err
		done := make(chan bool, 1)
		done <- true
		return done
	}
	return NewActivityOutputStream(gz, ch, errors, transforms...)
}

// NewActivityOutputStream will output json newlines from channel to the stream
func NewActivityOutputStream(stream io.WriteCloser, ch chan Activity, errors chan<- error, transforms ...TransformActivityFunc) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() {
			if gz, ok := stream.(*gzip.Writer); ok {
				gz.Flush()
				gz.Close()
			}
			stream.Close()
			done <- true
		}()
		for item := range ch {
			in := &item
			var skip bool
			var err error
			for _, transform := range transforms {
				in, err = transform(in)
				if err != nil {
					errors <- err
					return
				}
				if in == nil {
					skip = true
					break
				}
			}
			if !skip {
				buf, err := json.Marshal(in)
				if err != nil {
					errors <- err
					return
				}
				stream.Write(buf)
				stream.Write([]byte{'\n'})
			}
		}
	}()
	return done
}

// ActivitySendEvent is an event detail for sending data
type ActivitySendEvent struct {
	Activity *Activity
	headers  map[string]string
	time     time.Time
	key      string
}

var _ datamodel.ModelSendEvent = (*ActivitySendEvent)(nil)

// Key is the key to use for the message
func (e *ActivitySendEvent) Key() string {
	if e.key == "" {
		return e.Activity.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *ActivitySendEvent) Object() datamodel.Model {
	return e.Activity
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *ActivitySendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *ActivitySendEvent) Timestamp() time.Time {
	return e.time
}

// ActivitySendEventOpts is a function handler for setting opts
type ActivitySendEventOpts func(o *ActivitySendEvent)

// WithActivitySendEventKey sets the key value to a value different than the object ID
func WithActivitySendEventKey(key string) ActivitySendEventOpts {
	return func(o *ActivitySendEvent) {
		o.key = key
	}
}

// WithActivitySendEventTimestamp sets the timestamp value
func WithActivitySendEventTimestamp(tv time.Time) ActivitySendEventOpts {
	return func(o *ActivitySendEvent) {
		o.time = tv
	}
}

// WithActivitySendEventHeader sets the timestamp value
func WithActivitySendEventHeader(key, value string) ActivitySendEventOpts {
	return func(o *ActivitySendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewActivitySendEvent returns a new ActivitySendEvent instance
func NewActivitySendEvent(o *Activity, opts ...ActivitySendEventOpts) *ActivitySendEvent {
	res := &ActivitySendEvent{
		Activity: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewActivityProducer will stream data from the channel
func NewActivityProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		for {
			select {
			case <-ctx.Done():
				return
			case item := <-ch:
				if item == nil {
					empty <- true
					return
				}
				if object, ok := item.Object().(*Activity); ok {
					binary, codec, err := object.ToAvroBinary()
					if err != nil {
						errors <- fmt.Errorf("error encoding %s to avro binary data. %v", object.String(), err)
						return
					}
					headers := map[string]string{}
					object.SetEventHeaders(headers)
					for k, v := range item.Headers() {
						headers[k] = v
					}
					tv := item.Timestamp()
					if tv.IsZero() {
						tv = object.GetTimestamp() // if not provided in the message, use the objects value
					}
					if tv.IsZero() {
						tv = time.Now() // if its still zero, use the ingest time
					}
					msg := eventing.Message{
						Encoding:  eventing.AvroEncoding,
						Key:       item.Key(),
						Value:     binary,
						Codec:     codec,
						Headers:   headers,
						Timestamp: tv,
						Topic:     object.GetTopicName().String(),
					}
					if err := producer.Send(ctx, msg); err != nil {
						errors <- fmt.Errorf("error sending %s. %v", object.String(), err)
					}
				} else {
					errors <- fmt.Errorf("invalid event received. expected an object of type pipeline.activity.Activity but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewActivityConsumer will stream data from the topic into the provided channel
func NewActivityConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object Activity
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into pipeline.activity.Activity: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avro data into pipeline.activity.Activity: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for pipeline.activity.Activity")
			}

			// ignore messages that have exceeded the TTL
			cfg := object.GetTopicConfig()
			if cfg != nil && cfg.TTL != 0 && msg.Timestamp.UTC().Add(cfg.TTL).Sub(time.Now().UTC()) < 0 {
				return nil
			}
			msg.Codec = object.GetAvroCodec() // match the codec

			ch <- &ActivityReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object Activity
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &ActivityReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// ActivityReceiveEvent is an event detail for receiving data
type ActivityReceiveEvent struct {
	Activity *Activity
	message  eventing.Message
	eof      bool
}

var _ datamodel.ModelReceiveEvent = (*ActivityReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *ActivityReceiveEvent) Object() datamodel.Model {
	return e.Activity
}

// Message returns the underlying message data for the event
func (e *ActivityReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *ActivityReceiveEvent) EOF() bool {
	return e.eof
}

// ActivityProducer implements the datamodel.ModelEventProducer
type ActivityProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*ActivityProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *ActivityProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *ActivityProducer) Close() error {
	p.mu.Lock()
	closed := p.closed
	p.closed = true
	p.mu.Unlock()
	if !closed {
		close(p.ch)
		<-p.empty
		p.cancel()
		<-p.done
	}
	return nil
}

// NewProducerChannel returns a channel which can be used for producing Model events
func (o *Activity) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *Activity) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &ActivityProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewActivityProducer(newctx, producer, ch, errors, empty),
	}
}

// NewActivityProducerChannel returns a channel which can be used for producing Model events
func NewActivityProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewActivityProducerChannelSize(producer, 0, errors)
}

// NewActivityProducerChannelSize returns a channel which can be used for producing Model events
func NewActivityProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &ActivityProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewActivityProducer(newctx, producer, ch, errors, empty),
	}
}

// ActivityConsumer implements the datamodel.ModelEventConsumer
type ActivityConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*ActivityConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *ActivityConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *ActivityConsumer) Close() error {
	c.mu.Lock()
	closed := c.closed
	c.closed = true
	c.mu.Unlock()
	var err error
	if !closed {
		c.callback.Close()
		err = c.consumer.Close()
	}
	return err
}

// NewConsumerChannel returns a consumer channel which can be used to consume Model events
func (o *Activity) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &ActivityConsumer{
		ch:       ch,
		callback: NewActivityConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewActivityConsumerChannel returns a consumer channel which can be used to consume Model events
func NewActivityConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &ActivityConsumer{
		ch:       ch,
		callback: NewActivityConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
