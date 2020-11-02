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

	// ExportStatTable is the default table name
	ExportStatTable datamodel.ModelNameType = "agent_exportstat"

	// ExportStatModelName is the model name
	ExportStatModelName datamodel.ModelNameType = "agent.ExportStat"
)

const (
	// ExportStatModelCreatedDateColumn is the column json value created_date
	ExportStatModelCreatedDateColumn = "created_date"
	// ExportStatModelCreatedDateEpochColumn is the column json value epoch
	ExportStatModelCreatedDateEpochColumn = "epoch"
	// ExportStatModelCreatedDateOffsetColumn is the column json value offset
	ExportStatModelCreatedDateOffsetColumn = "offset"
	// ExportStatModelCreatedDateRfc3339Column is the column json value rfc3339
	ExportStatModelCreatedDateRfc3339Column = "rfc3339"
	// ExportStatModelCreatedAtColumn is the column json value created_ts
	ExportStatModelCreatedAtColumn = "created_ts"
	// ExportStatModelCustomerIDColumn is the column json value customer_id
	ExportStatModelCustomerIDColumn = "customer_id"
	// ExportStatModelIDColumn is the column json value id
	ExportStatModelIDColumn = "id"
	// ExportStatModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	ExportStatModelIntegrationInstanceIDColumn = "integration_instance_id"
	// ExportStatModelJobIDColumn is the column json value job_id
	ExportStatModelJobIDColumn = "job_id"
	// ExportStatModelRefIDColumn is the column json value ref_id
	ExportStatModelRefIDColumn = "ref_id"
	// ExportStatModelRefTypeColumn is the column json value ref_type
	ExportStatModelRefTypeColumn = "ref_type"
	// ExportStatModelTTLStartsAtColumn is the column json value ttl_starts_at
	ExportStatModelTTLStartsAtColumn = "ttl_starts_at"
	// ExportStatModelUpdatedAtColumn is the column json value updated_ts
	ExportStatModelUpdatedAtColumn = "updated_ts"
)

// ExportStatCreatedDate represents the object structure for created_date
type ExportStatCreatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toExportStatCreatedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ExportStatCreatedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *ExportStatCreatedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toExportStatCreatedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toExportStatCreatedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toExportStatCreatedDateObject(o.Rfc3339, false),
	}
}

func (o *ExportStatCreatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ExportStatCreatedDate) FromMap(kv map[string]interface{}) {

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

// ExportStat A record that is created if an agent is currently exporting
type ExportStat struct {
	// CreatedDate when the export was created
	CreatedDate ExportStatCreatedDate `json:"created_date" codec:"created_date" bson:"created_date" yaml:"created_date" faker:"-"`
	// CreatedAt the date the record was created in Epoch time
	CreatedAt int64 `json:"created_ts" codec:"created_ts" bson:"created_ts" yaml:"created_ts" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the id of the integration instance.
	IntegrationInstanceID string `json:"integration_instance_id" codec:"integration_instance_id" bson:"integration_instance_id" yaml:"integration_instance_id" faker:"-"`
	// JobID the job_id of the export
	JobID string `json:"job_id" codec:"job_id" bson:"job_id" yaml:"job_id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// TTLStartsAt the date the record will expire
	TTLStartsAt *string `json:"ttl_starts_at,omitempty" codec:"ttl_starts_at,omitempty" bson:"ttl_starts_at" yaml:"ttl_starts_at,omitempty" faker:"-"`
	// UpdatedAt the date the record was updated in Epoch time
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*ExportStat)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*ExportStat)(nil)

func toExportStatObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ExportStat:
		return v.ToMap()

	case ExportStatCreatedDate:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of ExportStat
func (o *ExportStat) String() string {
	return fmt.Sprintf("agent.ExportStat<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *ExportStat) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *ExportStat) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *ExportStat) GetTableName() string {
	return ExportStatTable.String()
}

// GetModelName returns the name of the model
func (o *ExportStat) GetModelName() datamodel.ModelNameType {
	return ExportStatModelName
}

// NewExportStatID provides a template for generating an ID field for ExportStat
func NewExportStatID(IntegrationInstanceID string) string {
	return hash.Values(IntegrationInstanceID)
}

func (o *ExportStat) setDefaults(frommap bool) {

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
func (o *ExportStat) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *ExportStat) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *ExportStat) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *ExportStat) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *ExportStat) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *ExportStat) IsMutable() bool {
	return true
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *ExportStat) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *ExportStat) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *ExportStat) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = ExportStatModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *ExportStat) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *ExportStat) GetCustomerID() string {
	return o.CustomerID
}

// SetCustomerID will return the customer_id
func (o *ExportStat) SetCustomerID(id string) {
	o.CustomerID = id
}

// GetRefType will return the ref_type
func (o *ExportStat) GetRefType() string {
	return o.RefType
}

// SetRefType will return the ref_type
func (o *ExportStat) SetRefType(t string) {
	o.RefType = t
}

// Clone returns an exact copy of ExportStat
func (o *ExportStat) Clone() datamodel.Model {
	c := new(ExportStat)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *ExportStat) Anon() datamodel.Model {
	c := new(ExportStat)
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
func (o *ExportStat) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *ExportStat) UnmarshalJSON(data []byte) error {
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
func (o *ExportStat) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *ExportStat) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two ExportStat objects are equal
func (o *ExportStat) IsEqual(other *ExportStat) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *ExportStat) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"created_date":            toExportStatObject(o.CreatedDate, false),
		"created_ts":              toExportStatObject(o.CreatedAt, false),
		"customer_id":             toExportStatObject(o.CustomerID, false),
		"id":                      toExportStatObject(o.ID, false),
		"integration_instance_id": toExportStatObject(o.IntegrationInstanceID, false),
		"job_id":                  toExportStatObject(o.JobID, false),
		"ref_id":                  toExportStatObject(o.RefID, false),
		"ref_type":                toExportStatObject(o.RefType, false),
		"ttl_starts_at":           toExportStatObject(o.TTLStartsAt, true),
		"updated_ts":              toExportStatObject(o.UpdatedAt, false),
		"hashcode":                toExportStatObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *ExportStat) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["created_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.CreatedDate.FromMap(kv)
		} else if sv, ok := val.(ExportStatCreatedDate); ok {
			// struct
			o.CreatedDate = sv
		} else if sp, ok := val.(*ExportStatCreatedDate); ok {
			// struct pointer
			o.CreatedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.CreatedDate.Epoch = dt.Epoch
			o.CreatedDate.Rfc3339 = dt.Rfc3339
			o.CreatedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
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
	if val, ok := kv["ttl_starts_at"].(*string); ok {
		o.TTLStartsAt = val
	} else if val, ok := kv["ttl_starts_at"].(string); ok {
		o.TTLStartsAt = &val
	} else {
		if val, ok := kv["ttl_starts_at"]; ok {
			if val == nil {
				o.TTLStartsAt = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.TTLStartsAt = pstrings.Pointer(fmt.Sprintf("%v", val))
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
func (o *ExportStat) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CreatedDate)
	args = append(args, o.CreatedAt)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.JobID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.TTLStartsAt)
	args = append(args, o.UpdatedAt)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// SetIntegrationInstanceID will set the integration instance ID
func (o *ExportStat) SetIntegrationInstanceID(id string) {
	o.IntegrationInstanceID = id
}

// GetIntegrationInstanceID will return the integration instance ID
func (o *ExportStat) GetIntegrationInstanceID() *string {
	return &o.IntegrationInstanceID
}

// GetHydrationQuery returns a query for all fields, and one level deep of relations.
// This query requires "id" to be in the query variables.
func (o *ExportStat) GetHydrationQuery() string {
	return `query GoExportStatQuery($id: ID, $nocache: Boolean) {
	agent {
		ExportStat(_id: $id, nocache: $nocache) {
			created_date {
			epoch
			offset
			rfc3339
			}
			created_ts
			customer_id
			_id
			integration_instance_id
			job_id
			ref_id
			ref_type
			ttl_starts_at
			updated_ts
		}
	}
}
`
}

func getExportStatQueryFields() string {
	var sb strings.Builder

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
	// id
	sb.WriteString("\t\t\t_id\n")
	// scalar
	sb.WriteString("\t\t\tintegration_instance_id\n")
	// scalar
	sb.WriteString("\t\t\tjob_id\n")
	// scalar
	sb.WriteString("\t\t\tref_id\n")
	// scalar
	sb.WriteString("\t\t\tref_type\n")
	// scalar
	sb.WriteString("\t\t\tttl_starts_at\n")
	// scalar
	sb.WriteString("\t\t\tupdated_ts\n")
	return sb.String()
}

// ExportStatPageInfo is a grapqhl PageInfo
type ExportStatPageInfo struct {
	StartCursor     string `json:"startCursor,omitempty"`
	EndCursor       string `json:"endCursor,omitempty"`
	HasNextPage     bool   `json:"hasNextPage,omitempty"`
	HasPreviousPage bool   `json:"hasPreviousPage,omitempty"`
}

// ExportStatCacheInfo is the grapqhl CacheInfo
type ExportStatCacheInfo struct {
	Cached bool    `json:"cached,omitempty"`
	ID     *string `json:"id,omitempty"`
	ETag   *string `json:"etag,omitempty"`
}

// ExportStatConnection is a grapqhl connection
type ExportStatConnection struct {
	Edges      []*ExportStatEdge   `json:"edges,omitempty"`
	PageInfo   ExportStatPageInfo  `json:"pageInfo,omitempty"`
	CacheInfo  ExportStatCacheInfo `json:"cacheInfo,omitempty"`
	TotalCount *int64              `json:"totalCount,omitempty"`
}

// ExportStatEdge is a grapqhl edge
type ExportStatEdge struct {
	Node *ExportStat `json:"node,omitempty"`
}

// QueryManyExportStatNode is a grapqhl query many node
type QueryManyExportStatNode struct {
	Object *ExportStatConnection `json:"ExportStats,omitempty"`
}

// QueryManyExportStatData is a grapqhl query many data node
type QueryManyExportStatData struct {
	Data *QueryManyExportStatNode `json:"agent,omitempty"`
}

// QueryOneExportStatNode is a grapqhl query one node
type QueryOneExportStatNode struct {
	Object *ExportStat `json:"ExportStat,omitempty"`
}

// QueryOneExportStatData is a grapqhl query one data node
type QueryOneExportStatData struct {
	Data *QueryOneExportStatNode `json:"agent,omitempty"`
}

// ExportStatQuery is query struct
type ExportStatQuery struct {
	Filters []string      `json:"filters,omitempty"`
	Params  []interface{} `json:"params,omitempty"`
}

// ExportStatOrder is the order direction
type ExportStatOrder *string

var (
	// ExportStatOrderDesc is descending
	ExportStatOrderDesc ExportStatOrder = pstrings.Pointer("DESC")
	// ExportStatOrderAsc is ascending
	ExportStatOrderAsc ExportStatOrder = pstrings.Pointer("ASC")
)

// ExportStatQueryInput is query input struct
type ExportStatQueryInput struct {
	First   *int64           `json:"first,omitempty"`
	Last    *int64           `json:"last,omitempty"`
	Before  *string          `json:"before,omitempty"`
	After   *string          `json:"after,omitempty"`
	Query   *ExportStatQuery `json:"query,omitempty"`
	OrderBy *string          `json:"orderBy,omitempty"`
	Order   ExportStatOrder  `json:"order,omitempty"`
	NoCache bool             `json:"nocache,omitempty"`
}

// NewExportStatQuery is a convenience for building a *ExportStatQuery
func NewExportStatQuery(params ...interface{}) *ExportStatQueryInput {
	if len(params)%2 != 0 {
		panic("incorrect number of arguments passed")
	}
	q := &ExportStatQuery{
		Filters: make([]string, 0),
		Params:  make([]interface{}, 0),
	}
	for i := 0; i < len(params); i += 2 {
		q.Filters = append(q.Filters, params[i].(string))
		q.Params = append(q.Params, params[i+1])
	}
	return &ExportStatQueryInput{
		Query: q,
	}
}

// FindExportStat will query an ExportStat by id
func FindExportStat(client graphql.Client, id string) (*ExportStat, error) {
	variables := make(graphql.Variables)
	variables["id"] = id
	var sb strings.Builder
	sb.WriteString("query GoExportStatQuery($id: ID) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tExportStat(_id: $id) {\n")
	sb.WriteString(getExportStatQueryFields())
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res QueryOneExportStatData
	if err := client.Query(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	if res.Data != nil {
		return res.Data.Object, nil
	}
	return nil, nil
}

// FindExportStatWithoutCache will query an ExportStat by id avoiding the cache
func FindExportStatWithoutCache(client graphql.Client, id string) (*ExportStat, error) {
	variables := make(graphql.Variables)
	variables["id"] = id
	variables["nocache"] = true
	var sb strings.Builder
	sb.WriteString("query GoExportStatQuery($id: ID, $nocache: Boolean) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tExportStat(_id: $id, nocache: $nocache) {\n")
	sb.WriteString(getExportStatQueryFields())
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res QueryOneExportStatData
	if err := client.Query(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	if res.Data != nil {
		return res.Data.Object, nil
	}
	return nil, nil
}

// FindExportStats will query for any ExportStats matching the query
func FindExportStats(client graphql.Client, input *ExportStatQueryInput) (*ExportStatConnection, error) {
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
	sb.WriteString("query GoExportStatQueryMany($first: Int, $last: Int, $before: Cursor, $after: Cursor, $query: QueryInput, $order: SortOrderEnum, $orderBy: AgentExportStatColumnEnum, $nocache: Boolean) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tExportStats(first: $first last: $last before: $before after: $after query: $query orderBy: $orderBy order: $order nocache: $nocache) {\n")
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
	sb.WriteString(getExportStatQueryFields())
	sb.WriteString("\t\t\t\t}\n")
	sb.WriteString("\t\t\t}\n")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res QueryManyExportStatData
	if err := client.Query(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	return res.Data.Object, nil
}

// FindExportStatsPaginatedCallback is a callback function for handling each page
type FindExportStatsPaginatedCallback func(conn *ExportStatConnection) (bool, error)

// FindExportStatsPaginated will query for any ExportStats matching the query and return each page callback
func FindExportStatsPaginated(client graphql.Client, query *ExportStatQuery, pageSize int64, callback FindExportStatsPaginatedCallback) error {
	input := &ExportStatQueryInput{
		First: &pageSize,
		Query: query,
	}
	for {
		res, err := FindExportStats(client, input)
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

// UpdateExportStatNode is a grapqhl update node
type UpdateExportStatNode struct {
	Object *ExportStat `json:"updateExportStat,omitempty"`
}

// UpdateExportStatData is a grapqhl update data node
type UpdateExportStatData struct {
	Data *UpdateExportStatNode `json:"agent,omitempty"`
}

// ExecExportStatUpdateMutation returns a graphql update mutation result for ExportStat
func ExecExportStatUpdateMutation(client graphql.Client, id string, input graphql.Variables, upsert bool) (*ExportStat, error) {
	if !upsert && id == "" {
		return nil, fmt.Errorf("id is required with upsert false")
	}
	variables := make(graphql.Variables)
	variables["id"] = id
	variables["upsert"] = upsert
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation GoExportStatUpdateMutation($id: String, $input: UpdateAgentExportStatInput, $upsert: Boolean) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tupdateExportStat(_id: $id, input: $input, upsert: $upsert) {\n")
	sb.WriteString(getExportStatQueryFields())
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateExportStatData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	return res.Data.Object, nil
}

// ExecExportStatSilentUpdateMutation returns a graphql update mutation result for ExportStat
func ExecExportStatSilentUpdateMutation(client graphql.Client, id string, input graphql.Variables, upsert bool) error {
	if !upsert && id == "" {
		return fmt.Errorf("id is required with upsert false")
	}
	variables := make(graphql.Variables)
	variables["id"] = id
	variables["upsert"] = upsert
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation GoExportStatUpdateMutation($id: String, $input: UpdateAgentExportStatInput, $upsert: Boolean) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tupdateExportStat(_id: $id, input: $input, upsert: $upsert) {\n")
	sb.WriteString("\t\t\t_id")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateExportStatData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return err
	}
	return nil
}

// ExecExportStatDeleteMutation executes a graphql delete mutation for ExportStat
func ExecExportStatDeleteMutation(client graphql.Client, id string) error {
	variables := make(graphql.Variables)
	variables["id"] = id
	var sb strings.Builder
	sb.WriteString("mutation GoExportStatDeleteMutation($id: String!) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tdeleteExportStat(_id: $id)\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res interface{}
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return err
	}
	return nil
}

func CreateExportStat(client graphql.Client, model ExportStat) error {
	variables := make(graphql.Variables)
	input := model.ToMap()
	delete(input, "hashcode")
	delete(input, "customer_id")
	delete(input, "created_ts")
	delete(input, "id")
	input["_id"] = model.GetID()
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation GoCreateExportStat($input: CreateAgentExportStatInput!) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tcreateExportStat(input: $input) {\n")
	sb.WriteString("\t\t\t_id")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateExportStatData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return err
	}
	return nil
}

// ExportStatPartial is a partial struct for upsert mutations for ExportStat
type ExportStatPartial struct {
	// CreatedDate when the export was created
	CreatedDate *ExportStatCreatedDate `json:"created_date,omitempty"`
	// JobID the job_id of the export
	JobID *string `json:"job_id,omitempty"`
}

var _ datamodel.PartialModel = (*ExportStatPartial)(nil)

// GetModelName returns the name of the model
func (o *ExportStatPartial) GetModelName() datamodel.ModelNameType {
	return ExportStatModelName
}

// ToMap returns the object as a map
func (o *ExportStatPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"created_date": toExportStatObject(o.CreatedDate, true),
		"job_id":       toExportStatObject(o.JobID, true),
	}
	for k, v := range kv {
		if v == nil || reflect.ValueOf(v).IsZero() {
			delete(kv, k)
		} else {
			if k == "created_date" {
				if dt, ok := v.(*ExportStatCreatedDate); ok {
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
func (o *ExportStatPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *ExportStatPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *ExportStatPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *ExportStatPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *ExportStatPartial) FromMap(kv map[string]interface{}) {

	if o.CreatedDate == nil {
		o.CreatedDate = &ExportStatCreatedDate{}
	}

	if val, ok := kv["created_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.CreatedDate.FromMap(kv)
		} else if sv, ok := val.(ExportStatCreatedDate); ok {
			// struct
			o.CreatedDate = &sv
		} else if sp, ok := val.(*ExportStatCreatedDate); ok {
			// struct pointer
			o.CreatedDate = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.CreatedDate.Epoch = dt.Epoch
			o.CreatedDate.Rfc3339 = dt.Rfc3339
			o.CreatedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
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

	if val, ok := kv["job_id"].(*string); ok {
		o.JobID = val
	} else if val, ok := kv["job_id"].(string); ok {
		o.JobID = &val
	} else {
		if val, ok := kv["job_id"]; ok {
			if val == nil {
				o.JobID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.JobID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	o.setDefaults(false)
}
