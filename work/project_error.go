// DO NOT EDIT -- generated code

// Package work - the system which contains project work
package work

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
	// ProjectErrorTopic is the default topic name
	ProjectErrorTopic datamodel.TopicNameType = "work_ProjectError"

	// ProjectErrorTable is the default table name
	ProjectErrorTable datamodel.ModelNameType = "work_projecterror"

	// ProjectErrorModelName is the model name
	ProjectErrorModelName datamodel.ModelNameType = "work.ProjectError"
)

const (
	// ProjectErrorModelCreatedAtColumn is the column json value created_ts
	ProjectErrorModelCreatedAtColumn = "created_ts"
	// ProjectErrorModelCustomerIDColumn is the column json value customer_id
	ProjectErrorModelCustomerIDColumn = "customer_id"
	// ProjectErrorModelErrorMessageColumn is the column json value error_message
	ProjectErrorModelErrorMessageColumn = "error_message"
	// ProjectErrorModelErroredColumn is the column json value errored
	ProjectErrorModelErroredColumn = "errored"
	// ProjectErrorModelIDColumn is the column json value id
	ProjectErrorModelIDColumn = "id"
	// ProjectErrorModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	ProjectErrorModelIntegrationInstanceIDColumn = "integration_instance_id"
	// ProjectErrorModelProjectIDColumn is the column json value project_id
	ProjectErrorModelProjectIDColumn = "project_id"
	// ProjectErrorModelRefIDColumn is the column json value ref_id
	ProjectErrorModelRefIDColumn = "ref_id"
	// ProjectErrorModelRefTypeColumn is the column json value ref_type
	ProjectErrorModelRefTypeColumn = "ref_type"
	// ProjectErrorModelUpdatedAtColumn is the column json value updated_ts
	ProjectErrorModelUpdatedAtColumn = "updated_ts"
)

// ProjectError the project error contains details about any errors found during processing
type ProjectError struct {
	// CreatedAt the date the record was created in Epoch time
	CreatedAt int64 `json:"created_ts" codec:"created_ts" bson:"created_ts" yaml:"created_ts" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ErrorMessage the error message
	ErrorMessage *string `json:"error_message,omitempty" codec:"error_message,omitempty" bson:"error_message" yaml:"error_message,omitempty" faker:"-"`
	// Errored if the project is in an errored state
	Errored bool `json:"errored" codec:"errored" bson:"errored" yaml:"errored" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// ProjectID the project id
	ProjectID string `json:"project_id" codec:"project_id" bson:"project_id" yaml:"project_id" faker:"-"`
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
var _ datamodel.Model = (*ProjectError)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*ProjectError)(nil)

func toProjectErrorObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ProjectError:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of ProjectError
func (o *ProjectError) String() string {
	return fmt.Sprintf("work.ProjectError<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *ProjectError) GetTopicName() datamodel.TopicNameType {
	return ProjectErrorTopic
}

// GetStreamName returns the name of the stream
func (o *ProjectError) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *ProjectError) GetTableName() string {
	return ProjectErrorTable.String()
}

// GetModelName returns the name of the model
func (o *ProjectError) GetModelName() datamodel.ModelNameType {
	return ProjectErrorModelName
}

// NewProjectErrorID provides a template for generating an ID field for ProjectError
func NewProjectErrorID(customerID string, ProjectID string) string {
	return hash.Values(customerID, ProjectID)
}

func (o *ProjectError) setDefaults(frommap bool) {

	if o.ID == "" {
		// set the id from the spec provided in the model
		o.ID = hash.Values(o.CustomerID, o.ProjectID)
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *ProjectError) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *ProjectError) GetTopicKey() string {
	var i interface{} = o.ID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *ProjectError) GetTimestamp() time.Time {
	var dt interface{} = o.UpdatedAt
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
	}
	panic("not sure how to handle the date time format for ProjectError")
}

// GetRefID returns the RefID for the object
func (o *ProjectError) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *ProjectError) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *ProjectError) IsMutable() bool {
	return true
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *ProjectError) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *ProjectError) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *ProjectError) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = ProjectErrorModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *ProjectError) GetTopicConfig() *datamodel.ModelTopicConfig {
	retention, err := time.ParseDuration("87360h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 87360h0m0s. " + err.Error())
	}

	ttl, err := time.ParseDuration("0s")
	if err != nil {
		ttl = 0
	}
	if ttl == 0 && retention != 0 {
		ttl = retention // they should be the same if not set
	}
	return &datamodel.ModelTopicConfig{
		Key:               "id",
		Timestamp:         "updated_ts",
		NumPartitions:     8,
		CleanupPolicy:     datamodel.CleanupPolicy("compact"),
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetCustomerID will return the customer_id
func (o *ProjectError) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of ProjectError
func (o *ProjectError) Clone() datamodel.Model {
	c := new(ProjectError)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *ProjectError) Anon() datamodel.Model {
	c := new(ProjectError)
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
func (o *ProjectError) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *ProjectError) UnmarshalJSON(data []byte) error {
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
func (o *ProjectError) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *ProjectError) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two ProjectError objects are equal
func (o *ProjectError) IsEqual(other *ProjectError) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *ProjectError) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"created_ts":              toProjectErrorObject(o.CreatedAt, false),
		"customer_id":             toProjectErrorObject(o.CustomerID, false),
		"error_message":           toProjectErrorObject(o.ErrorMessage, true),
		"errored":                 toProjectErrorObject(o.Errored, false),
		"id":                      toProjectErrorObject(o.ID, false),
		"integration_instance_id": toProjectErrorObject(o.IntegrationInstanceID, true),
		"project_id":              toProjectErrorObject(o.ProjectID, false),
		"ref_id":                  toProjectErrorObject(o.RefID, false),
		"ref_type":                toProjectErrorObject(o.RefType, false),
		"updated_ts":              toProjectErrorObject(o.UpdatedAt, false),
		"hashcode":                toProjectErrorObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *ProjectError) FromMap(kv map[string]interface{}) {

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
	if val, ok := kv["project_id"].(string); ok {
		o.ProjectID = val
	} else {
		if val, ok := kv["project_id"]; ok {
			if val == nil {
				o.ProjectID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.ProjectID = fmt.Sprintf("%v", val)
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
func (o *ProjectError) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CreatedAt)
	args = append(args, o.CustomerID)
	args = append(args, o.ErrorMessage)
	args = append(args, o.Errored)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.ProjectID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.UpdatedAt)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

func getProjectErrorQueryFields() string {
	var sb strings.Builder

	// scalar
	sb.WriteString("\t\t\tcreated_ts\n")
	// scalar
	sb.WriteString("\t\t\tcustomer_id\n")
	// scalar
	sb.WriteString("\t\t\terror_message\n")
	// scalar
	sb.WriteString("\t\t\terrored\n")
	// id
	sb.WriteString("\t\t\t_id\n")
	// scalar
	sb.WriteString("\t\t\tintegration_instance_id\n")
	// scalar
	sb.WriteString("\t\t\tproject_id\n")
	// scalar
	sb.WriteString("\t\t\tref_id\n")
	// scalar
	sb.WriteString("\t\t\tref_type\n")
	// scalar
	sb.WriteString("\t\t\tupdated_ts\n")
	return sb.String()
}

// ProjectErrorPageInfo is a grapqhl PageInfo
type ProjectErrorPageInfo struct {
	StartCursor     string `json:"startCursor,omitempty"`
	EndCursor       string `json:"endCursor,omitempty"`
	HasNextPage     bool   `json:"hasNextPage,omitempty"`
	HasPreviousPage bool   `json:"hasPreviousPage,omitempty"`
}

// ProjectErrorConnection is a grapqhl connection
type ProjectErrorConnection struct {
	Edges      []*ProjectErrorEdge  `json:"edges,omitempty"`
	PageInfo   ProjectErrorPageInfo `json:"pageInfo,omitempty"`
	TotalCount *int64               `json:"totalCount,omitempty"`
}

// ProjectErrorEdge is a grapqhl edge
type ProjectErrorEdge struct {
	Node *ProjectError `json:"node,omitempty"`
}

// QueryManyProjectErrorNode is a grapqhl query many node
type QueryManyProjectErrorNode struct {
	Object *ProjectErrorConnection `json:"ProjectErrors,omitempty"`
}

// QueryManyProjectErrorData is a grapqhl query many data node
type QueryManyProjectErrorData struct {
	Data *QueryManyProjectErrorNode `json:"work,omitempty"`
}

// QueryOneProjectErrorNode is a grapqhl query one node
type QueryOneProjectErrorNode struct {
	Object *ProjectError `json:"ProjectError,omitempty"`
}

// QueryOneProjectErrorData is a grapqhl query one data node
type QueryOneProjectErrorData struct {
	Data *QueryOneProjectErrorNode `json:"work,omitempty"`
}

// ProjectErrorQuery is query struct
type ProjectErrorQuery struct {
	Filters []string      `json:"filters,omitempty"`
	Params  []interface{} `json:"params,omitempty"`
}

// ProjectErrorOrder is the order direction
type ProjectErrorOrder *string

var (
	// ProjectErrorOrderDesc is descending
	ProjectErrorOrderDesc ProjectErrorOrder = pstrings.Pointer("DESC")
	// ProjectErrorOrderAsc is ascending
	ProjectErrorOrderAsc ProjectErrorOrder = pstrings.Pointer("ASC")
)

// ProjectErrorQueryInput is query input struct
type ProjectErrorQueryInput struct {
	First   *int64             `json:"first,omitempty"`
	Last    *int64             `json:"last,omitempty"`
	Before  *string            `json:"before,omitempty"`
	After   *string            `json:"after,omitempty"`
	Query   *ProjectErrorQuery `json:"query,omitempty"`
	OrderBy *string            `json:"orderBy,omitempty"`
	Order   ProjectErrorOrder  `json:"order,omitempty"`
}

// NewProjectErrorQuery is a convenience for building a *ProjectErrorQuery
func NewProjectErrorQuery(params ...interface{}) *ProjectErrorQueryInput {
	if len(params)%2 != 0 {
		panic("incorrect number of arguments passed")
	}
	q := &ProjectErrorQuery{
		Filters: make([]string, 0),
		Params:  make([]interface{}, 0),
	}
	for i := 0; i < len(params); i += 2 {
		q.Filters = append(q.Filters, params[i].(string))
		q.Params = append(q.Params, params[i+1])
	}
	return &ProjectErrorQueryInput{
		Query: q,
	}
}

// FindProjectError will query an ProjectError by id
func FindProjectError(client graphql.Client, id string) (*ProjectError, error) {
	variables := make(graphql.Variables)
	variables["id"] = id
	var sb strings.Builder
	sb.WriteString("query ProjectErrorQuery($id: ID) {\n")
	sb.WriteString("\twork {\n")
	sb.WriteString("\t\tProjectError(_id: $id) {\n")
	sb.WriteString(getProjectErrorQueryFields())
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res QueryOneProjectErrorData
	if err := client.Query(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	if res.Data != nil {
		return res.Data.Object, nil
	}
	return nil, nil
}

// FindProjectErrors will query for any ProjectErrors matching the query
func FindProjectErrors(client graphql.Client, input *ProjectErrorQueryInput) (*ProjectErrorConnection, error) {
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
	sb.WriteString("query ProjectErrorQueryMany($first: Int, $last: Int, $before: Cursor, $after: Cursor, $query: QueryInput, $order: SortOrderEnum, $orderBy: WorkProjectErrorColumnEnum) {\n")
	sb.WriteString("\twork {\n")
	sb.WriteString("\t\tProjectErrors(first: $first last: $last before: $before after: $after query: $query orderBy: $orderBy order: $order) {\n")
	sb.WriteString("\t\t\tpageInfo {\n")
	sb.WriteString("\t\t\t\tstartCursor\n")
	sb.WriteString("\t\t\t\tendCursor\n")
	sb.WriteString("\t\t\t\thasNextPage\n")
	sb.WriteString("\t\t\t\thasPreviousPage\n")
	sb.WriteString("\t\t\t}\n")
	sb.WriteString("\t\t\ttotalCount\n")
	sb.WriteString("\t\t\tedges {\n")
	sb.WriteString("\t\t\t\tnode {\n")
	sb.WriteString(getProjectErrorQueryFields())
	sb.WriteString("\t\t\t\t}\n")
	sb.WriteString("\t\t\t}\n")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res QueryManyProjectErrorData
	if err := client.Query(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	return res.Data.Object, nil
}

// FindProjectErrorsPaginatedCallback is a callback function for handling each page
type FindProjectErrorsPaginatedCallback func(conn *ProjectErrorConnection) (bool, error)

// FindProjectErrorsPaginated will query for any ProjectErrors matching the query and return each page callback
func FindProjectErrorsPaginated(client graphql.Client, query *ProjectErrorQuery, pageSize int64, callback FindProjectErrorsPaginatedCallback) error {
	input := &ProjectErrorQueryInput{
		First: &pageSize,
		Query: query,
	}
	for {
		res, err := FindProjectErrors(client, input)
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

// HydratedQuery returns all fields, and one level deep of relations.
const ProjectErrorHydratedQuery = `query ProjectErrorQuery($id: ID) {
	work {
		ProjectError(_id: $id) {
			created_ts
			customer_id
			customer {
			_id
			customer_id
			created_date
			updated_date
			name
			active
			apikey_id
			work_ids
			codequality_ids
			sourcecode_ids
			auth {
			provider
			}
			customer_information_id
			crm_id
			}
			error_message
			errored
			_id
			integration_instance_id
			integration_instance {
			active
			config
			created_by_profile_id
			created_by_user_id
			created_date {
			epoch
			offset
			rfc3339
			}
			created_ts
			customer_id
			deleted_by_profile_id
			deleted_by_user_id
			deleted_date {
			epoch
			offset
			rfc3339
			}
			error_message
			errored
			export_acknowledged
			_id
			integration_id
			integration_instance_id
			interval
			job_id
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
			location
			name
			paused
			processed
			ref_id
			ref_type
			requires_historical
			state
			throttled
			throttled_until {
			epoch
			offset
			rfc3339
			}
			updated_ts
			webhooks {
			enabled
			error_message
			errored
			ref_id
			url
			}
			}
			project_id
			project {
			}
			ref_id
			ref_type
			updated_ts
		}
	}
}
`

// UpdateProjectErrorNode is a grapqhl update node
type UpdateProjectErrorNode struct {
	Object *ProjectError `json:"updateProjectError,omitempty"`
}

// UpdateProjectErrorData is a grapqhl update data node
type UpdateProjectErrorData struct {
	Data *UpdateProjectErrorNode `json:"work,omitempty"`
}

// ExecProjectErrorUpdateMutation returns a graphql update mutation result for ProjectError
func ExecProjectErrorUpdateMutation(client graphql.Client, id string, input graphql.Variables, upsert bool) (*ProjectError, error) {
	variables := make(graphql.Variables)
	variables["id"] = id
	variables["upsert"] = upsert
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation ProjectErrorUpdateMutation($id: String, $input: UpdateWorkProjectErrorInput, $upsert: Boolean) {\n")
	sb.WriteString("\twork {\n")
	sb.WriteString("\t\tupdateProjectError(_id: $id, input: $input, upsert: $upsert) {\n")
	sb.WriteString(getProjectErrorQueryFields())
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateProjectErrorData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	return res.Data.Object, nil
}

// ExecProjectErrorSilentUpdateMutation returns a graphql update mutation result for ProjectError
func ExecProjectErrorSilentUpdateMutation(client graphql.Client, id string, input graphql.Variables, upsert bool) error {
	variables := make(graphql.Variables)
	variables["id"] = id
	variables["upsert"] = upsert
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation ProjectErrorUpdateMutation($id: String, $input: UpdateWorkProjectErrorInput, $upsert: Boolean) {\n")
	sb.WriteString("\twork {\n")
	sb.WriteString("\t\tupdateProjectError(_id: $id, input: $input, upsert: $upsert) {\n")
	sb.WriteString("\t\t\t_id")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateProjectErrorData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return err
	}
	return nil
}

// ExecProjectErrorDeleteMutation executes a graphql delete mutation for ProjectError
func ExecProjectErrorDeleteMutation(client graphql.Client, id string) error {
	variables := make(graphql.Variables)
	variables["id"] = id
	var sb strings.Builder
	sb.WriteString("mutation ProjectErrorDeleteMutation($id: String!) {\n")
	sb.WriteString("\twork {\n")
	sb.WriteString("\t\tdeleteProjectError(_id: $id)\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res interface{}
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return err
	}
	return nil
}

func CreateProjectError(client graphql.Client, model ProjectError) error {
	variables := make(graphql.Variables)
	input := model.ToMap()
	delete(input, "hashcode")
	delete(input, "customer_id")
	delete(input, "created_ts")
	delete(input, "id")
	input["_id"] = model.GetID()
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation CreateProjectError($input: CreateWorkProjectErrorInput!) {\n")
	sb.WriteString("\twork {\n")
	sb.WriteString("\t\tcreateProjectError(input: $input) {\n")
	sb.WriteString("\t\t\t_id")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateProjectErrorData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return err
	}
	return nil
}

// ProjectErrorPartial is a partial struct for upsert mutations for ProjectError
type ProjectErrorPartial struct {
	// ErrorMessage the error message
	ErrorMessage *string `json:"error_message,omitempty"`
	// Errored if the project is in an errored state
	Errored *bool `json:"errored,omitempty"`
	// ProjectID the project id
	ProjectID *string `json:"project_id,omitempty"`
}

var _ datamodel.PartialModel = (*ProjectErrorPartial)(nil)

// GetModelName returns the name of the model
func (o *ProjectErrorPartial) GetModelName() datamodel.ModelNameType {
	return ProjectErrorModelName
}

// ToMap returns the object as a map
func (o *ProjectErrorPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"error_message": toProjectErrorObject(o.ErrorMessage, true),
		"errored":       toProjectErrorObject(o.Errored, true),
		"project_id":    toProjectErrorObject(o.ProjectID, true),
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
func (o *ProjectErrorPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *ProjectErrorPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *ProjectErrorPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *ProjectErrorPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *ProjectErrorPartial) FromMap(kv map[string]interface{}) {
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
	if val, ok := kv["project_id"].(*string); ok {
		o.ProjectID = val
	} else if val, ok := kv["project_id"].(string); ok {
		o.ProjectID = &val
	} else {
		if val, ok := kv["project_id"]; ok {
			if val == nil {
				o.ProjectID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.ProjectID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	o.setDefaults(false)
}
