// DO NOT EDIT -- generated code

// Package customer - customer specific data
package customer

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
	"strings"
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
	// TeamTopic is the default topic name
	TeamTopic datamodel.TopicNameType = "customer_Team_topic"

	// TeamStream is the default stream name
	TeamStream datamodel.TopicNameType = "customer_Team_stream"

	// TeamTable is the default table name
	TeamTable datamodel.TopicNameType = "customer_Team"

	// TeamModelName is the model name
	TeamModelName datamodel.ModelNameType = "customer.Team"
)

const (
	// TeamActiveColumn is the active column name
	TeamActiveColumn = "active"
	// TeamChildrenIdsColumn is the children_ids column name
	TeamChildrenIdsColumn = "children_ids"
	// TeamCreatedAtColumn is the created_ts column name
	TeamCreatedAtColumn = "created_ts"
	// TeamCustomerIDColumn is the customer_id column name
	TeamCustomerIDColumn = "customer_id"
	// TeamDescriptionColumn is the description column name
	TeamDescriptionColumn = "description"
	// TeamIDColumn is the id column name
	TeamIDColumn = "id"
	// TeamLeafColumn is the leaf column name
	TeamLeafColumn = "leaf"
	// TeamNameColumn is the name column name
	TeamNameColumn = "name"
	// TeamParentIdsColumn is the parent_ids column name
	TeamParentIdsColumn = "parent_ids"
	// TeamRefIDColumn is the ref_id column name
	TeamRefIDColumn = "ref_id"
	// TeamRefTypeColumn is the ref_type column name
	TeamRefTypeColumn = "ref_type"
	// TeamUpdatedAtColumn is the updated_ts column name
	TeamUpdatedAtColumn = "updated_ts"
)

// Team a team is a grouping of one or more users
type Team struct {
	// Active whether the team is tracked in pinpoint
	Active *bool `json:"active" bson:"active" yaml:"active" faker:"-"`
	// ChildrenIds the children_ids for this team
	ChildrenIds []string `json:"children_ids" bson:"children_ids" yaml:"children_ids" faker:"-"`
	// CreatedAt the date the record was created in Epoch time
	CreatedAt int64 `json:"created_ts" bson:"created_ts" yaml:"created_ts" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Description the description of the team
	Description string `json:"description" bson:"description" yaml:"description" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// Leaf True when team has no children_ids
	Leaf bool `json:"leaf" bson:"leaf" yaml:"leaf" faker:"-"`
	// Name the name of the team
	Name string `json:"name" bson:"name" yaml:"name" faker:"team"`
	// ParentIds the parent_ids for this team
	ParentIds []string `json:"parent_ids" bson:"parent_ids" yaml:"parent_ids" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// UpdatedAt the date the record was updated in Epoch time
	UpdatedAt int64 `json:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Team)(nil)

func toTeamObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toTeamObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if o == nil {
		return toTeamObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toTeamObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toTeamObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toTeamObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toTeamObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toTeamObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toTeamObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toTeamObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toTeamObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toTeamObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toTeamObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case map[string]interface{}:
		return o
	case *map[string]interface{}:
		return v
	case map[string]string:
		return v
	case *map[string]string:
		return *v
	case *Team:
		return v.ToMap()
	case Team:
		return v.ToMap()
	case []string, []int64, []float64, []bool:
		return o
	case *[]string:
		return (*(o.(*[]string)))
	case *[]int64:
		return (*(o.(*[]int64)))
	case *[]float64:
		return (*(o.(*[]float64)))
	case *[]bool:
		return (*(o.(*[]bool)))
	case []interface{}:
		a := o.([]interface{})
		arr := make([]interface{}, 0)
		for _, av := range a {
			arr = append(arr, toTeamObject(av, isavro, false, ""))
		}
		return arr

	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of Team
func (o *Team) String() string {
	return fmt.Sprintf("customer.Team<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Team) GetTopicName() datamodel.TopicNameType {
	return TeamTopic
}

// GetModelName returns the name of the model
func (o *Team) GetModelName() datamodel.ModelNameType {
	return TeamModelName
}

func (o *Team) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Team) GetID() string {
	if o.ID == "" {
		// set the id from the spec provided in the model
		o.ID = hash.Values(o.CustomerID, randomString(64))
	}
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Team) GetTopicKey() string {
	var i interface{} = o.ID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Team) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *Team) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Team) IsMaterialized() bool {
	return true
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Team) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	idletime, err := time.ParseDuration("1s")
	if err != nil {
		panic(err)
	}
	return &datamodel.ModelMaterializeConfig{
		KeyName:   "id",
		TableName: "customer_team",
		BatchSize: 5000,
		IdleTime:  idletime,
	}
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Team) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *Team) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = TeamModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Team) GetTopicConfig() *datamodel.ModelTopicConfig {
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
		Timestamp:         "",
		NumPartitions:     8,
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetStateKey returns a key for use in state store
func (o *Team) GetStateKey() string {
	key := "id"
	return fmt.Sprintf("%s_%s", key, o.GetID())
}

// GetCustomerID will return the customer_id
func (o *Team) GetCustomerID() string {
	return o.CustomerID
}

// Clone returns an exact copy of Team
func (o *Team) Clone() datamodel.Model {
	c := new(Team)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Team) Anon() datamodel.Model {
	c := new(Team)
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
func (o *Team) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Team) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

var cachedCodecTeam *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *Team) GetAvroCodec() *goavro.Codec {
	if cachedCodecTeam == nil {
		c, err := GetTeamAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecTeam = c
	}
	return cachedCodecTeam
}

// ToAvroBinary returns the data as Avro binary data
func (o *Team) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *Team) FromAvroBinary(value []byte) error {
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
func (o *Team) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Team objects are equal
func (o *Team) IsEqual(other *Team) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Team) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
		if o.ChildrenIds == nil {
			o.ChildrenIds = make([]string, 0)
		}
		if o.ParentIds == nil {
			o.ParentIds = make([]string, 0)
		}
	}
	o.setDefaults()
	return map[string]interface{}{
		"active":       toTeamObject(o.Active, isavro, true, "boolean"),
		"children_ids": toTeamObject(o.ChildrenIds, isavro, false, "children_ids"),
		"created_ts":   toTeamObject(o.CreatedAt, isavro, false, "long"),
		"customer_id":  toTeamObject(o.CustomerID, isavro, false, "string"),
		"description":  toTeamObject(o.Description, isavro, false, "string"),
		"id":           toTeamObject(o.ID, isavro, false, "string"),
		"leaf":         toTeamObject(o.Leaf, isavro, false, "boolean"),
		"name":         toTeamObject(o.Name, isavro, false, "string"),
		"parent_ids":   toTeamObject(o.ParentIds, isavro, false, "parent_ids"),
		"ref_id":       toTeamObject(o.RefID, isavro, false, "string"),
		"ref_type":     toTeamObject(o.RefType, isavro, false, "string"),
		"updated_ts":   toTeamObject(o.UpdatedAt, isavro, false, "long"),
		"hashcode":     toTeamObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *Team) FromMap(kv map[string]interface{}) {
	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["active"].(*bool); ok {
		o.Active = val
	} else if val, ok := kv["active"].(bool); ok {
		o.Active = &val
	} else {
		val := kv["active"]
		if val == nil {
			o.Active = number.BoolPointer(number.ToBoolAny(nil))
		} else {
			// if coming in as avro union, convert it back
			if kv, ok := val.(map[string]interface{}); ok {
				val = kv["bool"]
			}
			o.Active = number.BoolPointer(number.ToBoolAny(val))
		}
	}
	if val := kv["children_ids"]; val != nil {
		na := make([]string, 0)
		if a, ok := val.([]string); ok {
			na = append(na, a...)
		} else {
			if a, ok := val.([]interface{}); ok {
				for _, ae := range a {
					if av, ok := ae.(string); ok {
						na = append(na, av)
					} else {
						b, _ := json.Marshal(ae)
						var av string
						if err := json.Unmarshal(b, &av); err != nil {
							panic("unsupported type for children_ids field entry: " + reflect.TypeOf(ae).String())
						}
						na = append(na, av)
					}
				}
			} else if s, ok := val.(string); ok {
				for _, sv := range strings.Split(s, ",") {
					na = append(na, strings.TrimSpace(sv))
				}
			} else {
				fmt.Println(reflect.TypeOf(val).String())
				panic("unsupported type for children_ids field")
			}
		}
		o.ChildrenIds = na
	} else {
		o.ChildrenIds = []string{}
	}
	if o.ChildrenIds == nil {
		o.ChildrenIds = make([]string, 0)
	}
	if val, ok := kv["created_ts"].(int64); ok {
		o.CreatedAt = val
	} else {
		val := kv["created_ts"]
		if val == nil {
			o.CreatedAt = number.ToInt64Any(nil)
		} else {
			if tv, ok := val.(time.Time); ok {
				val = datetime.TimeToEpoch(tv)
			}
			o.CreatedAt = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["customer_id"].(string); ok {
		o.CustomerID = val
	} else {
		val := kv["customer_id"]
		if val == nil {
			o.CustomerID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.CustomerID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["description"].(string); ok {
		o.Description = val
	} else {
		val := kv["description"]
		if val == nil {
			o.Description = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Description = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["id"].(string); ok {
		o.ID = val
	} else {
		val := kv["id"]
		if val == nil {
			o.ID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.ID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["leaf"].(bool); ok {
		o.Leaf = val
	} else {
		val := kv["leaf"]
		if val == nil {
			o.Leaf = number.ToBoolAny(nil)
		} else {
			o.Leaf = number.ToBoolAny(val)
		}
	}
	if val, ok := kv["name"].(string); ok {
		o.Name = val
	} else {
		val := kv["name"]
		if val == nil {
			o.Name = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Name = fmt.Sprintf("%v", val)
		}
	}
	if val := kv["parent_ids"]; val != nil {
		na := make([]string, 0)
		if a, ok := val.([]string); ok {
			na = append(na, a...)
		} else {
			if a, ok := val.([]interface{}); ok {
				for _, ae := range a {
					if av, ok := ae.(string); ok {
						na = append(na, av)
					} else {
						b, _ := json.Marshal(ae)
						var av string
						if err := json.Unmarshal(b, &av); err != nil {
							panic("unsupported type for parent_ids field entry: " + reflect.TypeOf(ae).String())
						}
						na = append(na, av)
					}
				}
			} else if s, ok := val.(string); ok {
				for _, sv := range strings.Split(s, ",") {
					na = append(na, strings.TrimSpace(sv))
				}
			} else {
				fmt.Println(reflect.TypeOf(val).String())
				panic("unsupported type for parent_ids field")
			}
		}
		o.ParentIds = na
	} else {
		o.ParentIds = []string{}
	}
	if o.ParentIds == nil {
		o.ParentIds = make([]string, 0)
	}
	if val, ok := kv["ref_id"].(string); ok {
		o.RefID = val
	} else {
		val := kv["ref_id"]
		if val == nil {
			o.RefID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.RefID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["ref_type"].(string); ok {
		o.RefType = val
	} else {
		val := kv["ref_type"]
		if val == nil {
			o.RefType = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.RefType = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["updated_ts"].(int64); ok {
		o.UpdatedAt = val
	} else {
		val := kv["updated_ts"]
		if val == nil {
			o.UpdatedAt = number.ToInt64Any(nil)
		} else {
			if tv, ok := val.(time.Time); ok {
				val = datetime.TimeToEpoch(tv)
			}
			o.UpdatedAt = number.ToInt64Any(val)
		}
	}
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *Team) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Active)
	args = append(args, o.ChildrenIds)
	args = append(args, o.CreatedAt)
	args = append(args, o.CustomerID)
	args = append(args, o.Description)
	args = append(args, o.ID)
	args = append(args, o.Leaf)
	args = append(args, o.Name)
	args = append(args, o.ParentIds)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.UpdatedAt)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// CreateTeam creates a new Team in the database
func CreateTeam(ctx context.Context, db datamodel.Storage, o *Team) error {
	o.setDefaults()
	return db.Create(ctx, o)
}

// DeleteTeam deletes a Team in the database
func DeleteTeam(ctx context.Context, db datamodel.Storage, o *Team) error {
	o.setDefaults()
	return db.Delete(ctx, o)
}

// UpdateTeam updates a Team in the database
func UpdateTeam(ctx context.Context, db datamodel.Storage, o *Team) error {
	o.setDefaults()
	return db.Update(ctx, o)
}

// FindTeam returns a Team from the database
func FindTeam(ctx context.Context, db datamodel.Storage, id string) (*Team, error) {
	kv, err := db.FindOne(ctx, TeamModelName, id)
	if err != nil {
		return nil, err
	}
	if kv == nil {
		return nil, nil
	}
	return kv.(*Team), nil
}

// FindTeams returns all Team from the database matching keys
func FindTeams(ctx context.Context, db datamodel.Storage, kv map[string]interface{}) ([]*Team, error) {
	res, err := db.Find(ctx, TeamModelName, kv)
	if err != nil {
		return nil, err
	}
	if res != nil {
		arr := make([]*Team, 0)
		for _, m := range res {
			arr = append(arr, m.(*Team))
		}
		return arr, nil
	}
	return nil, nil
}

// GetTeamAvroSchemaSpec creates the avro schema specification for Team
func GetTeamAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "customer",
		"name":      "Team",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "hashcode",
				"type": "string",
			},
			map[string]interface{}{
				"name":    "active",
				"type":    []interface{}{"null", "boolean"},
				"default": nil,
			},
			map[string]interface{}{
				"name": "children_ids",
				"type": map[string]interface{}{"items": "string", "type": "array", "name": "children_ids"},
			},
			map[string]interface{}{
				"name": "created_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "customer_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "description",
				"type": "string",
			},
			map[string]interface{}{
				"name": "id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "leaf",
				"type": "boolean",
			},
			map[string]interface{}{
				"name": "name",
				"type": "string",
			},
			map[string]interface{}{
				"name": "parent_ids",
				"type": map[string]interface{}{"type": "array", "name": "parent_ids", "items": "string"},
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
				"name": "updated_ts",
				"type": "long",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetTeamAvroSchema creates the avro schema for Team
func GetTeamAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetTeamAvroSchemaSpec())
}

// TransformTeamFunc is a function for transforming Team during processing
type TransformTeamFunc func(input *Team) (*Team, error)

// NewTeamPipe creates a pipe for processing Team items
func NewTeamPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformTeamFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewTeamInputStream(input, errors)
	var stream chan Team
	if len(transforms) > 0 {
		stream = make(chan Team, 1000)
	} else {
		stream = inch
	}
	outdone := NewTeamOutputStream(output, stream, errors)
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

// NewTeamInputStreamDir creates a channel for reading Team as JSON newlines from a directory of files
func NewTeamInputStreamDir(dir string, errors chan<- error, transforms ...TransformTeamFunc) (chan Team, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/customer/team\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan Team)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for team")
		ch := make(chan Team)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewTeamInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan Team)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewTeamInputStreamFile creates an channel for reading Team as JSON newlines from filename
func NewTeamInputStreamFile(filename string, errors chan<- error, transforms ...TransformTeamFunc) (chan Team, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan Team)
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
			ch := make(chan Team)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewTeamInputStream(f, errors, transforms...)
}

// NewTeamInputStream creates an channel for reading Team as JSON newlines from stream
func NewTeamInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformTeamFunc) (chan Team, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan Team, 1000)
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
			var item Team
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

// NewTeamOutputStreamDir will output json newlines from channel and save in dir
func NewTeamOutputStreamDir(dir string, ch chan Team, errors chan<- error, transforms ...TransformTeamFunc) <-chan bool {
	fp := filepath.Join(dir, "/customer/team\\.json(\\.gz)?$")
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
	return NewTeamOutputStream(gz, ch, errors, transforms...)
}

// NewTeamOutputStream will output json newlines from channel to the stream
func NewTeamOutputStream(stream io.WriteCloser, ch chan Team, errors chan<- error, transforms ...TransformTeamFunc) <-chan bool {
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

// TeamSendEvent is an event detail for sending data
type TeamSendEvent struct {
	Team    *Team
	headers map[string]string
	time    time.Time
	key     string
}

var _ datamodel.ModelSendEvent = (*TeamSendEvent)(nil)

// Key is the key to use for the message
func (e *TeamSendEvent) Key() string {
	if e.key == "" {
		return e.Team.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *TeamSendEvent) Object() datamodel.Model {
	return e.Team
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *TeamSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *TeamSendEvent) Timestamp() time.Time {
	return e.time
}

// TeamSendEventOpts is a function handler for setting opts
type TeamSendEventOpts func(o *TeamSendEvent)

// WithTeamSendEventKey sets the key value to a value different than the object ID
func WithTeamSendEventKey(key string) TeamSendEventOpts {
	return func(o *TeamSendEvent) {
		o.key = key
	}
}

// WithTeamSendEventTimestamp sets the timestamp value
func WithTeamSendEventTimestamp(tv time.Time) TeamSendEventOpts {
	return func(o *TeamSendEvent) {
		o.time = tv
	}
}

// WithTeamSendEventHeader sets the timestamp value
func WithTeamSendEventHeader(key, value string) TeamSendEventOpts {
	return func(o *TeamSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewTeamSendEvent returns a new TeamSendEvent instance
func NewTeamSendEvent(o *Team, opts ...TeamSendEventOpts) *TeamSendEvent {
	res := &TeamSendEvent{
		Team: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewTeamProducer will stream data from the channel
func NewTeamProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
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
				if object, ok := item.Object().(*Team); ok {
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
					errors <- fmt.Errorf("invalid event received. expected an object of type customer.Team but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewTeamConsumer will stream data from the topic into the provided channel
func NewTeamConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object Team
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into customer.Team: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avri data into customer.Team: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for customer.Team")
			}
			msg.Codec = object.GetAvroCodec() // match the codec

			//ignore messages that have exceeded the TTL
			cfg := object.GetTopicConfig()
			if cfg != nil && cfg.TTL != 0 && msg.Timestamp.Add(cfg.TTL).Sub(time.Now()) < 0 {
				return nil
			}
			ch <- &TeamReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object Team
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &TeamReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// TeamReceiveEvent is an event detail for receiving data
type TeamReceiveEvent struct {
	Team    *Team
	message eventing.Message
	eof     bool
}

var _ datamodel.ModelReceiveEvent = (*TeamReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *TeamReceiveEvent) Object() datamodel.Model {
	return e.Team
}

// Message returns the underlying message data for the event
func (e *TeamReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *TeamReceiveEvent) EOF() bool {
	return e.eof
}

// TeamProducer implements the datamodel.ModelEventProducer
type TeamProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*TeamProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *TeamProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *TeamProducer) Close() error {
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
func (o *Team) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *Team) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &TeamProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewTeamProducer(newctx, producer, ch, errors, empty),
	}
}

// NewTeamProducerChannel returns a channel which can be used for producing Model events
func NewTeamProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewTeamProducerChannelSize(producer, 0, errors)
}

// NewTeamProducerChannelSize returns a channel which can be used for producing Model events
func NewTeamProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &TeamProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewTeamProducer(newctx, producer, ch, errors, empty),
	}
}

// TeamConsumer implements the datamodel.ModelEventConsumer
type TeamConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*TeamConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *TeamConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *TeamConsumer) Close() error {
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
func (o *Team) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &TeamConsumer{
		ch:       ch,
		callback: NewTeamConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewTeamConsumerChannel returns a consumer channel which can be used to consume Model events
func NewTeamConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &TeamConsumer{
		ch:       ch,
		callback: NewTeamConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
