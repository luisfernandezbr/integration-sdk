// DO NOT EDIT -- generated code

// Package agent -
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

	// ExportCompleteTable is the default table name
	ExportCompleteTable datamodel.ModelNameType = "agent_exportcomplete"

	// ExportCompleteModelName is the model name
	ExportCompleteModelName datamodel.ModelNameType = "agent.ExportComplete"
)

const (
	// ExportCompleteModelCreatedAtColumn is the column json value created_ts
	ExportCompleteModelCreatedAtColumn = "created_ts"
	// ExportCompleteModelCustomerIDColumn is the column json value customer_id
	ExportCompleteModelCustomerIDColumn = "customer_id"
	// ExportCompleteModelEndedAtColumn is the column json value ended_ts
	ExportCompleteModelEndedAtColumn = "ended_ts"
	// ExportCompleteModelErrorColumn is the column json value error
	ExportCompleteModelErrorColumn = "error"
	// ExportCompleteModelHistoricalColumn is the column json value historical
	ExportCompleteModelHistoricalColumn = "historical"
	// ExportCompleteModelIDColumn is the column json value id
	ExportCompleteModelIDColumn = "id"
	// ExportCompleteModelIntegrationIDColumn is the column json value integration_id
	ExportCompleteModelIntegrationIDColumn = "integration_id"
	// ExportCompleteModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	ExportCompleteModelIntegrationInstanceIDColumn = "integration_instance_id"
	// ExportCompleteModelJobIDColumn is the column json value job_id
	ExportCompleteModelJobIDColumn = "job_id"
	// ExportCompleteModelRefIDColumn is the column json value ref_id
	ExportCompleteModelRefIDColumn = "ref_id"
	// ExportCompleteModelRefTypeColumn is the column json value ref_type
	ExportCompleteModelRefTypeColumn = "ref_type"
	// ExportCompleteModelStartedAtColumn is the column json value started_ts
	ExportCompleteModelStartedAtColumn = "started_ts"
	// ExportCompleteModelStatsColumn is the column json value stats
	ExportCompleteModelStatsColumn = "stats"
	// ExportCompleteModelSuccessColumn is the column json value success
	ExportCompleteModelSuccessColumn = "success"
	// ExportCompleteModelUpdatedAtColumn is the column json value updated_ts
	ExportCompleteModelUpdatedAtColumn = "updated_ts"
)

// ExportComplete a record of every export completed by the agent
type ExportComplete struct {
	// CreatedAt the date the record was created in Epoch time
	CreatedAt int64 `json:"created_ts" codec:"created_ts" bson:"created_ts" yaml:"created_ts" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// EndedAt the timestamp when the export finished
	EndedAt int64 `json:"ended_ts" codec:"ended_ts" bson:"ended_ts" yaml:"ended_ts" faker:"-"`
	// Error if not successfully, will be set to the error message from the integration
	Error *string `json:"error,omitempty" codec:"error,omitempty" bson:"error" yaml:"error,omitempty" faker:"-"`
	// Historical true if historical
	Historical bool `json:"historical" codec:"historical" bson:"historical" yaml:"historical" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationID The ID of the integration
	IntegrationID string `json:"integration_id" codec:"integration_id" bson:"integration_id" yaml:"integration_id" faker:"-"`
	// IntegrationInstanceID The ID of the integration instance
	IntegrationInstanceID string `json:"integration_instance_id" codec:"integration_instance_id" bson:"integration_instance_id" yaml:"integration_instance_id" faker:"-"`
	// JobID The job ID
	JobID string `json:"job_id" codec:"job_id" bson:"job_id" yaml:"job_id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// StartedAt the timestamp when the export started
	StartedAt int64 `json:"started_ts" codec:"started_ts" bson:"started_ts" yaml:"started_ts" faker:"-"`
	// Stats a JSON string of any integration specific stats that were collected
	Stats *string `json:"stats,omitempty" codec:"stats,omitempty" bson:"stats" yaml:"stats,omitempty" faker:"-"`
	// Success if successful, will be set to true. if not, will be set to false with an error message
	Success bool `json:"success" codec:"success" bson:"success" yaml:"success" faker:"-"`
	// UpdatedAt the date the record was updated in Epoch time
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*ExportComplete)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*ExportComplete)(nil)

func toExportCompleteObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ExportComplete:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of ExportComplete
func (o *ExportComplete) String() string {
	return fmt.Sprintf("agent.ExportComplete<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *ExportComplete) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *ExportComplete) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *ExportComplete) GetTableName() string {
	return ExportCompleteTable.String()
}

// GetModelName returns the name of the model
func (o *ExportComplete) GetModelName() datamodel.ModelNameType {
	return ExportCompleteModelName
}

// NewExportCompleteID provides a template for generating an ID field for ExportComplete
func NewExportCompleteID(customerID string, JobID string, IntegrationID string) string {
	return hash.Values(customerID, JobID, IntegrationID)
}

func (o *ExportComplete) setDefaults(frommap bool) {

	if o.ID == "" {
		// set the id from the spec provided in the model
		o.ID = hash.Values(o.CustomerID, o.JobID, o.IntegrationID)
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *ExportComplete) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *ExportComplete) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *ExportComplete) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *ExportComplete) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *ExportComplete) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *ExportComplete) IsMutable() bool {
	return true
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *ExportComplete) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *ExportComplete) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *ExportComplete) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = ExportCompleteModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *ExportComplete) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *ExportComplete) GetCustomerID() string {

	return o.CustomerID

}

// SetCustomerID will return the customer_id
func (o *ExportComplete) SetCustomerID(id string) {

	o.CustomerID = id

}

// GetRefType will return the ref_type
func (o *ExportComplete) GetRefType() string {
	return o.RefType
}

// SetRefType will return the ref_type
func (o *ExportComplete) SetRefType(t string) {
	o.RefType = t
}

// Clone returns an exact copy of ExportComplete
func (o *ExportComplete) Clone() datamodel.Model {
	c := new(ExportComplete)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *ExportComplete) Anon() datamodel.Model {
	c := new(ExportComplete)
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
func (o *ExportComplete) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *ExportComplete) UnmarshalJSON(data []byte) error {
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
func (o *ExportComplete) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *ExportComplete) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two ExportComplete objects are equal
func (o *ExportComplete) IsEqual(other *ExportComplete) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *ExportComplete) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"created_ts":              toExportCompleteObject(o.CreatedAt, false),
		"customer_id":             toExportCompleteObject(o.CustomerID, false),
		"ended_ts":                toExportCompleteObject(o.EndedAt, false),
		"error":                   toExportCompleteObject(o.Error, true),
		"historical":              toExportCompleteObject(o.Historical, false),
		"id":                      toExportCompleteObject(o.ID, false),
		"integration_id":          toExportCompleteObject(o.IntegrationID, false),
		"integration_instance_id": toExportCompleteObject(o.IntegrationInstanceID, false),
		"job_id":                  toExportCompleteObject(o.JobID, false),
		"ref_id":                  toExportCompleteObject(o.RefID, false),
		"ref_type":                toExportCompleteObject(o.RefType, false),
		"started_ts":              toExportCompleteObject(o.StartedAt, false),
		"stats":                   toExportCompleteObject(o.Stats, true),
		"success":                 toExportCompleteObject(o.Success, false),
		"updated_ts":              toExportCompleteObject(o.UpdatedAt, false),
		"hashcode":                toExportCompleteObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *ExportComplete) FromMap(kv map[string]interface{}) {

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
	if val, ok := kv["ended_ts"].(int64); ok {
		o.EndedAt = val
	} else {
		if val, ok := kv["ended_ts"]; ok {
			if val == nil {
				o.EndedAt = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.EndedAt = number.ToInt64Any(val)
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
	if val, ok := kv["historical"].(bool); ok {
		o.Historical = val
	} else {
		if val, ok := kv["historical"]; ok {
			if val == nil {
				o.Historical = false
			} else {
				o.Historical = number.ToBoolAny(val)
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
	if val, ok := kv["started_ts"].(int64); ok {
		o.StartedAt = val
	} else {
		if val, ok := kv["started_ts"]; ok {
			if val == nil {
				o.StartedAt = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.StartedAt = number.ToInt64Any(val)
			}
		}
	}
	if val, ok := kv["stats"].(*string); ok {
		o.Stats = val
	} else if val, ok := kv["stats"].(string); ok {
		o.Stats = &val
	} else {
		if val, ok := kv["stats"]; ok {
			if val == nil {
				o.Stats = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Stats = pstrings.Pointer(fmt.Sprintf("%v", val))
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
func (o *ExportComplete) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CreatedAt)
	args = append(args, o.CustomerID)
	args = append(args, o.EndedAt)
	args = append(args, o.Error)
	args = append(args, o.Historical)
	args = append(args, o.ID)
	args = append(args, o.IntegrationID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.JobID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.StartedAt)
	args = append(args, o.Stats)
	args = append(args, o.Success)
	args = append(args, o.UpdatedAt)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// SetIntegrationInstanceID will set the integration instance ID
func (o *ExportComplete) SetIntegrationInstanceID(id string) {
	o.IntegrationInstanceID = id
}

// GetIntegrationInstanceID will return the integration instance ID
func (o *ExportComplete) GetIntegrationInstanceID() *string {
	return &o.IntegrationInstanceID
}

// GetHydrationQuery returns a query for all fields, and one level deep of relations.
// This query requires "id" to be in the query variables.
func (o *ExportComplete) GetHydrationQuery() string {
	return `query GoExportCompleteQuery($id: ID) {
	agent {
		ExportComplete(_id: $id) {
			created_ts
			customer_id
			ended_ts
			error
			historical
			_id
			integration_id
			integration_instance_id
			job_id
			ref_id
			ref_type
			started_ts
			stats
			success
			updated_ts
		}
	}
}
`
}

func getExportCompleteQueryFields() string {
	var sb strings.Builder

	// scalar
	sb.WriteString("\t\t\tcreated_ts\n")
	// scalar
	sb.WriteString("\t\t\tcustomer_id\n")
	// scalar
	sb.WriteString("\t\t\tended_ts\n")
	// scalar
	sb.WriteString("\t\t\terror\n")
	// scalar
	sb.WriteString("\t\t\thistorical\n")
	// id
	sb.WriteString("\t\t\t_id\n")
	// scalar
	sb.WriteString("\t\t\tintegration_id\n")
	// scalar
	sb.WriteString("\t\t\tintegration_instance_id\n")
	// scalar
	sb.WriteString("\t\t\tjob_id\n")
	// scalar
	sb.WriteString("\t\t\tref_id\n")
	// scalar
	sb.WriteString("\t\t\tref_type\n")
	// scalar
	sb.WriteString("\t\t\tstarted_ts\n")
	// scalar
	sb.WriteString("\t\t\tstats\n")
	// scalar
	sb.WriteString("\t\t\tsuccess\n")
	// scalar
	sb.WriteString("\t\t\tupdated_ts\n")
	return sb.String()
}

// ExportCompletePageInfo is a grapqhl PageInfo
type ExportCompletePageInfo struct {
	StartCursor     string `json:"startCursor,omitempty"`
	EndCursor       string `json:"endCursor,omitempty"`
	HasNextPage     bool   `json:"hasNextPage,omitempty"`
	HasPreviousPage bool   `json:"hasPreviousPage,omitempty"`
}

// ExportCompleteConnection is a grapqhl connection
type ExportCompleteConnection struct {
	Edges      []*ExportCompleteEdge  `json:"edges,omitempty"`
	PageInfo   ExportCompletePageInfo `json:"pageInfo,omitempty"`
	TotalCount *int64                 `json:"totalCount,omitempty"`
}

// ExportCompleteEdge is a grapqhl edge
type ExportCompleteEdge struct {
	Node *ExportComplete `json:"node,omitempty"`
}

// QueryManyExportCompleteNode is a grapqhl query many node
type QueryManyExportCompleteNode struct {
	Object *ExportCompleteConnection `json:"ExportCompletes,omitempty"`
}

// QueryManyExportCompleteData is a grapqhl query many data node
type QueryManyExportCompleteData struct {
	Data *QueryManyExportCompleteNode `json:"agent,omitempty"`
}

// QueryOneExportCompleteNode is a grapqhl query one node
type QueryOneExportCompleteNode struct {
	Object *ExportComplete `json:"ExportComplete,omitempty"`
}

// QueryOneExportCompleteData is a grapqhl query one data node
type QueryOneExportCompleteData struct {
	Data *QueryOneExportCompleteNode `json:"agent,omitempty"`
}

// ExportCompleteQuery is query struct
type ExportCompleteQuery struct {
	Filters []string      `json:"filters,omitempty"`
	Params  []interface{} `json:"params,omitempty"`
}

// ExportCompleteOrder is the order direction
type ExportCompleteOrder *string

var (
	// ExportCompleteOrderDesc is descending
	ExportCompleteOrderDesc ExportCompleteOrder = pstrings.Pointer("DESC")
	// ExportCompleteOrderAsc is ascending
	ExportCompleteOrderAsc ExportCompleteOrder = pstrings.Pointer("ASC")
)

// ExportCompleteQueryInput is query input struct
type ExportCompleteQueryInput struct {
	First   *int64               `json:"first,omitempty"`
	Last    *int64               `json:"last,omitempty"`
	Before  *string              `json:"before,omitempty"`
	After   *string              `json:"after,omitempty"`
	Query   *ExportCompleteQuery `json:"query,omitempty"`
	OrderBy *string              `json:"orderBy,omitempty"`
	Order   ExportCompleteOrder  `json:"order,omitempty"`
}

// NewExportCompleteQuery is a convenience for building a *ExportCompleteQuery
func NewExportCompleteQuery(params ...interface{}) *ExportCompleteQueryInput {
	if len(params)%2 != 0 {
		panic("incorrect number of arguments passed")
	}
	q := &ExportCompleteQuery{
		Filters: make([]string, 0),
		Params:  make([]interface{}, 0),
	}
	for i := 0; i < len(params); i += 2 {
		q.Filters = append(q.Filters, params[i].(string))
		q.Params = append(q.Params, params[i+1])
	}
	return &ExportCompleteQueryInput{
		Query: q,
	}
}

// FindExportComplete will query an ExportComplete by id
func FindExportComplete(client graphql.Client, id string) (*ExportComplete, error) {
	variables := make(graphql.Variables)
	variables["id"] = id
	var sb strings.Builder
	sb.WriteString("query GoExportCompleteQuery($id: ID) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tExportComplete(_id: $id) {\n")
	sb.WriteString(getExportCompleteQueryFields())
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res QueryOneExportCompleteData
	if err := client.Query(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	if res.Data != nil {
		return res.Data.Object, nil
	}
	return nil, nil
}

// FindExportCompletes will query for any ExportCompletes matching the query
func FindExportCompletes(client graphql.Client, input *ExportCompleteQueryInput) (*ExportCompleteConnection, error) {
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
	sb.WriteString("query GoExportCompleteQueryMany($first: Int, $last: Int, $before: Cursor, $after: Cursor, $query: QueryInput, $order: SortOrderEnum, $orderBy: AgentExportCompleteColumnEnum) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tExportCompletes(first: $first last: $last before: $before after: $after query: $query orderBy: $orderBy order: $order) {\n")
	sb.WriteString("\t\t\tpageInfo {\n")
	sb.WriteString("\t\t\t\tstartCursor\n")
	sb.WriteString("\t\t\t\tendCursor\n")
	sb.WriteString("\t\t\t\thasNextPage\n")
	sb.WriteString("\t\t\t\thasPreviousPage\n")
	sb.WriteString("\t\t\t}\n")
	sb.WriteString("\t\t\ttotalCount\n")
	sb.WriteString("\t\t\tedges {\n")
	sb.WriteString("\t\t\t\tnode {\n")
	sb.WriteString(getExportCompleteQueryFields())
	sb.WriteString("\t\t\t\t}\n")
	sb.WriteString("\t\t\t}\n")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res QueryManyExportCompleteData
	if err := client.Query(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	return res.Data.Object, nil
}

// FindExportCompletesPaginatedCallback is a callback function for handling each page
type FindExportCompletesPaginatedCallback func(conn *ExportCompleteConnection) (bool, error)

// FindExportCompletesPaginated will query for any ExportCompletes matching the query and return each page callback
func FindExportCompletesPaginated(client graphql.Client, query *ExportCompleteQuery, pageSize int64, callback FindExportCompletesPaginatedCallback) error {
	input := &ExportCompleteQueryInput{
		First: &pageSize,
		Query: query,
	}
	for {
		res, err := FindExportCompletes(client, input)
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

// UpdateExportCompleteNode is a grapqhl update node
type UpdateExportCompleteNode struct {
	Object *ExportComplete `json:"updateExportComplete,omitempty"`
}

// UpdateExportCompleteData is a grapqhl update data node
type UpdateExportCompleteData struct {
	Data *UpdateExportCompleteNode `json:"agent,omitempty"`
}

// ExecExportCompleteUpdateMutation returns a graphql update mutation result for ExportComplete
func ExecExportCompleteUpdateMutation(client graphql.Client, id string, input graphql.Variables, upsert bool) (*ExportComplete, error) {
	if !upsert && id == "" {
		return nil, fmt.Errorf("id is required with upsert false")
	}
	variables := make(graphql.Variables)
	variables["id"] = id
	variables["upsert"] = upsert
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation GoExportCompleteUpdateMutation($id: String, $input: UpdateAgentExportCompleteInput, $upsert: Boolean) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tupdateExportComplete(_id: $id, input: $input, upsert: $upsert) {\n")
	sb.WriteString(getExportCompleteQueryFields())
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateExportCompleteData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	return res.Data.Object, nil
}

// ExecExportCompleteSilentUpdateMutation returns a graphql update mutation result for ExportComplete
func ExecExportCompleteSilentUpdateMutation(client graphql.Client, id string, input graphql.Variables, upsert bool) error {
	if !upsert && id == "" {
		return fmt.Errorf("id is required with upsert false")
	}
	variables := make(graphql.Variables)
	variables["id"] = id
	variables["upsert"] = upsert
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation GoExportCompleteUpdateMutation($id: String, $input: UpdateAgentExportCompleteInput, $upsert: Boolean) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tupdateExportComplete(_id: $id, input: $input, upsert: $upsert) {\n")
	sb.WriteString("\t\t\t_id")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateExportCompleteData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return err
	}
	return nil
}

// ExecExportCompleteDeleteMutation executes a graphql delete mutation for ExportComplete
func ExecExportCompleteDeleteMutation(client graphql.Client, id string) error {
	variables := make(graphql.Variables)
	variables["id"] = id
	var sb strings.Builder
	sb.WriteString("mutation GoExportCompleteDeleteMutation($id: String!) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tdeleteExportComplete(_id: $id)\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res interface{}
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return err
	}
	return nil
}

func CreateExportComplete(client graphql.Client, model ExportComplete) error {
	variables := make(graphql.Variables)
	input := model.ToMap()
	delete(input, "hashcode")
	delete(input, "customer_id")
	delete(input, "created_ts")
	delete(input, "id")
	input["_id"] = model.GetID()
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation GoCreateExportComplete($input: CreateAgentExportCompleteInput!) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tcreateExportComplete(input: $input) {\n")
	sb.WriteString("\t\t\t_id")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateExportCompleteData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return err
	}
	return nil
}

// ExportCompletePartial is a partial struct for upsert mutations for ExportComplete
type ExportCompletePartial struct {
	// EndedAt the timestamp when the export finished
	EndedAt *int64 `json:"ended_ts,omitempty"`
	// Error if not successfully, will be set to the error message from the integration
	Error *string `json:"error,omitempty"`
	// Historical true if historical
	Historical *bool `json:"historical,omitempty"`
	// IntegrationID The ID of the integration
	IntegrationID *string `json:"integration_id,omitempty"`
	// JobID The job ID
	JobID *string `json:"job_id,omitempty"`
	// StartedAt the timestamp when the export started
	StartedAt *int64 `json:"started_ts,omitempty"`
	// Stats a JSON string of any integration specific stats that were collected
	Stats *string `json:"stats,omitempty"`
	// Success if successful, will be set to true. if not, will be set to false with an error message
	Success *bool `json:"success,omitempty"`
}

var _ datamodel.PartialModel = (*ExportCompletePartial)(nil)

// GetModelName returns the name of the model
func (o *ExportCompletePartial) GetModelName() datamodel.ModelNameType {
	return ExportCompleteModelName
}

// ToMap returns the object as a map
func (o *ExportCompletePartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"ended_ts":       toExportCompleteObject(o.EndedAt, true),
		"error":          toExportCompleteObject(o.Error, true),
		"historical":     toExportCompleteObject(o.Historical, true),
		"integration_id": toExportCompleteObject(o.IntegrationID, true),
		"job_id":         toExportCompleteObject(o.JobID, true),
		"started_ts":     toExportCompleteObject(o.StartedAt, true),
		"stats":          toExportCompleteObject(o.Stats, true),
		"success":        toExportCompleteObject(o.Success, true),
	}
	for k, v := range kv {
		if v == nil || reflect.ValueOf(v).IsZero() {
			delete(kv, k)
		} else {
		}
	}
	return kv
}

// Stringify returns the object in JSON format as a string
func (o *ExportCompletePartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *ExportCompletePartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *ExportCompletePartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *ExportCompletePartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *ExportCompletePartial) FromMap(kv map[string]interface{}) {
	if val, ok := kv["ended_ts"].(*int64); ok {
		o.EndedAt = val
	} else if val, ok := kv["ended_ts"].(int64); ok {
		o.EndedAt = &val
	} else {
		if val, ok := kv["ended_ts"]; ok {
			if val == nil {
				o.EndedAt = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["long"]
				}
				o.EndedAt = number.Int64Pointer(number.ToInt64Any(val))
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
	if val, ok := kv["historical"].(*bool); ok {
		o.Historical = val
	} else if val, ok := kv["historical"].(bool); ok {
		o.Historical = &val
	} else {
		if val, ok := kv["historical"]; ok {
			if val == nil {
				o.Historical = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Historical = number.BoolPointer(number.ToBoolAny(val))
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
	if val, ok := kv["started_ts"].(*int64); ok {
		o.StartedAt = val
	} else if val, ok := kv["started_ts"].(int64); ok {
		o.StartedAt = &val
	} else {
		if val, ok := kv["started_ts"]; ok {
			if val == nil {
				o.StartedAt = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["long"]
				}
				o.StartedAt = number.Int64Pointer(number.ToInt64Any(val))
			}
		}
	}
	if val, ok := kv["stats"].(*string); ok {
		o.Stats = val
	} else if val, ok := kv["stats"].(string); ok {
		o.Stats = &val
	} else {
		if val, ok := kv["stats"]; ok {
			if val == nil {
				o.Stats = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Stats = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["success"].(*bool); ok {
		o.Success = val
	} else if val, ok := kv["success"].(bool); ok {
		o.Success = &val
	} else {
		if val, ok := kv["success"]; ok {
			if val == nil {
				o.Success = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Success = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}
	o.setDefaults(false)
}
