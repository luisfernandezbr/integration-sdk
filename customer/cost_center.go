// DO NOT EDIT -- generated code

// Package customer - customer specific data
package customer

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

	// CostCenterTable is the default table name
	CostCenterTable datamodel.ModelNameType = "customer_costcenter"

	// CostCenterModelName is the model name
	CostCenterModelName datamodel.ModelNameType = "customer.CostCenter"
)

const (
	// CostCenterModelActiveColumn is the column json value active
	CostCenterModelActiveColumn = "active"
	// CostCenterModelCostColumn is the column json value cost
	CostCenterModelCostColumn = "cost"
	// CostCenterModelCreatedAtColumn is the column json value created_ts
	CostCenterModelCreatedAtColumn = "created_ts"
	// CostCenterModelCustomerIDColumn is the column json value customer_id
	CostCenterModelCustomerIDColumn = "customer_id"
	// CostCenterModelDescriptionColumn is the column json value description
	CostCenterModelDescriptionColumn = "description"
	// CostCenterModelIDColumn is the column json value id
	CostCenterModelIDColumn = "id"
	// CostCenterModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	CostCenterModelIntegrationInstanceIDColumn = "integration_instance_id"
	// CostCenterModelNameColumn is the column json value name
	CostCenterModelNameColumn = "name"
	// CostCenterModelRefIDColumn is the column json value ref_id
	CostCenterModelRefIDColumn = "ref_id"
	// CostCenterModelRefTypeColumn is the column json value ref_type
	CostCenterModelRefTypeColumn = "ref_type"
	// CostCenterModelUpdatedAtColumn is the column json value updated_ts
	CostCenterModelUpdatedAtColumn = "updated_ts"
)

// CostCenter a cost center represents information about users and their cost
type CostCenter struct {
	// Active whether the cost center is tracked in pinpoint
	Active bool `json:"active" codec:"active" bson:"active" yaml:"active" faker:"-"`
	// Cost the cost value of the cost center
	Cost float64 `json:"cost" codec:"cost" bson:"cost" yaml:"cost" faker:"salary"`
	// CreatedAt the date the record was created in Epoch time
	CreatedAt int64 `json:"created_ts" codec:"created_ts" bson:"created_ts" yaml:"created_ts" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Description the description for the cost center
	Description string `json:"description" codec:"description" bson:"description" yaml:"description" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// Name the name of the cost center
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"costcenter"`
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
var _ datamodel.Model = (*CostCenter)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*CostCenter)(nil)

func toCostCenterObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CostCenter:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of CostCenter
func (o *CostCenter) String() string {
	return fmt.Sprintf("customer.CostCenter<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *CostCenter) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *CostCenter) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *CostCenter) GetTableName() string {
	return CostCenterTable.String()
}

// GetModelName returns the name of the model
func (o *CostCenter) GetModelName() datamodel.ModelNameType {
	return CostCenterModelName
}

// NewCostCenterID provides a template for generating an ID field for CostCenter
func NewCostCenterID(customerID string) string {
	return hash.Values(customerID, randomString(64))
}

func (o *CostCenter) setDefaults(frommap bool) {

	if o.ID == "" {
		// set the id from the spec provided in the model
		o.ID = hash.Values(o.CustomerID, randomString(64))
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *CostCenter) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *CostCenter) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *CostCenter) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *CostCenter) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *CostCenter) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *CostCenter) IsMutable() bool {
	return true
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *CostCenter) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *CostCenter) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *CostCenter) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = CostCenterModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *CostCenter) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *CostCenter) GetCustomerID() string {
	return o.CustomerID
}

// SetCustomerID will return the customer_id
func (o *CostCenter) SetCustomerID(id string) {
	o.CustomerID = id
}

// GetRefType will return the ref_type
func (o *CostCenter) GetRefType() string {
	return o.RefType
}

// SetRefType will return the ref_type
func (o *CostCenter) SetRefType(t string) {
	o.RefType = t
}

// Clone returns an exact copy of CostCenter
func (o *CostCenter) Clone() datamodel.Model {
	c := new(CostCenter)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *CostCenter) Anon() datamodel.Model {
	c := new(CostCenter)
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
func (o *CostCenter) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *CostCenter) UnmarshalJSON(data []byte) error {
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
func (o *CostCenter) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *CostCenter) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two CostCenter objects are equal
func (o *CostCenter) IsEqual(other *CostCenter) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *CostCenter) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"active":                  toCostCenterObject(o.Active, false),
		"cost":                    toCostCenterObject(o.Cost, false),
		"created_ts":              toCostCenterObject(o.CreatedAt, false),
		"customer_id":             toCostCenterObject(o.CustomerID, false),
		"description":             toCostCenterObject(o.Description, false),
		"id":                      toCostCenterObject(o.ID, false),
		"integration_instance_id": toCostCenterObject(o.IntegrationInstanceID, true),
		"name":                    toCostCenterObject(o.Name, false),
		"ref_id":                  toCostCenterObject(o.RefID, false),
		"ref_type":                toCostCenterObject(o.RefType, false),
		"updated_ts":              toCostCenterObject(o.UpdatedAt, false),
		"hashcode":                toCostCenterObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *CostCenter) FromMap(kv map[string]interface{}) {

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
	if val, ok := kv["cost"].(float64); ok {
		o.Cost = val
	} else {
		if val, ok := kv["cost"]; ok {
			if val == nil {
				o.Cost = 0.0
			} else {
				o.Cost = number.ToFloat64Any(val)
			}
		}
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
func (o *CostCenter) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Active)
	args = append(args, o.Cost)
	args = append(args, o.CreatedAt)
	args = append(args, o.CustomerID)
	args = append(args, o.Description)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.Name)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.UpdatedAt)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// SetIntegrationInstanceID will set the integration instance ID
func (o *CostCenter) SetIntegrationInstanceID(id string) {
	if id == "" {
		o.IntegrationInstanceID = nil
	} else {
		o.IntegrationInstanceID = &id
	}
}

// GetIntegrationInstanceID will return the integration instance ID
func (o *CostCenter) GetIntegrationInstanceID() *string {
	return o.IntegrationInstanceID
}

// GetHydrationQuery returns a query for all fields, and one level deep of relations.
// This query requires "id" to be in the query variables.
func (o *CostCenter) GetHydrationQuery() string {
	return `query GoCostCenterQuery($id: ID, $nocache: Boolean) {
	customer {
		CostCenter(_id: $id, nocache: $nocache) {
			active
			cost
			created_ts
			customer_id
			description
			_id
			integration_instance_id
			name
			ref_id
			ref_type
			updated_ts
		}
	}
}
`
}

func getCostCenterQueryFields() string {
	var sb strings.Builder

	// scalar
	sb.WriteString("\t\t\tactive\n")
	// scalar
	sb.WriteString("\t\t\tcost\n")
	// scalar
	sb.WriteString("\t\t\tcreated_ts\n")
	// scalar
	sb.WriteString("\t\t\tcustomer_id\n")
	// scalar
	sb.WriteString("\t\t\tdescription\n")
	// id
	sb.WriteString("\t\t\t_id\n")
	// scalar
	sb.WriteString("\t\t\tintegration_instance_id\n")
	// scalar
	sb.WriteString("\t\t\tname\n")
	// scalar
	sb.WriteString("\t\t\tref_id\n")
	// scalar
	sb.WriteString("\t\t\tref_type\n")
	// scalar
	sb.WriteString("\t\t\tupdated_ts\n")
	return sb.String()
}

// CostCenterPageInfo is a grapqhl PageInfo
type CostCenterPageInfo struct {
	StartCursor     string `json:"startCursor,omitempty"`
	EndCursor       string `json:"endCursor,omitempty"`
	HasNextPage     bool   `json:"hasNextPage,omitempty"`
	HasPreviousPage bool   `json:"hasPreviousPage,omitempty"`
}

// CostCenterCacheInfo is the grapqhl CacheInfo
type CostCenterCacheInfo struct {
	Cached bool    `json:"cached,omitempty"`
	ID     *string `json:"id,omitempty"`
	ETag   *string `json:"etag,omitempty"`
}

// CostCenterConnection is a grapqhl connection
type CostCenterConnection struct {
	Edges      []*CostCenterEdge   `json:"edges,omitempty"`
	PageInfo   CostCenterPageInfo  `json:"pageInfo,omitempty"`
	CacheInfo  CostCenterCacheInfo `json:"cacheInfo,omitempty"`
	TotalCount *int64              `json:"totalCount,omitempty"`
}

// CostCenterEdge is a grapqhl edge
type CostCenterEdge struct {
	Node *CostCenter `json:"node,omitempty"`
}

// QueryManyCostCenterNode is a grapqhl query many node
type QueryManyCostCenterNode struct {
	Object *CostCenterConnection `json:"CostCenters,omitempty"`
}

// QueryManyCostCenterData is a grapqhl query many data node
type QueryManyCostCenterData struct {
	Data *QueryManyCostCenterNode `json:"customer,omitempty"`
}

// QueryOneCostCenterNode is a grapqhl query one node
type QueryOneCostCenterNode struct {
	Object *CostCenter `json:"CostCenter,omitempty"`
}

// QueryOneCostCenterData is a grapqhl query one data node
type QueryOneCostCenterData struct {
	Data *QueryOneCostCenterNode `json:"customer,omitempty"`
}

// CostCenterQuery is query struct
type CostCenterQuery struct {
	Filters []string      `json:"filters,omitempty"`
	Params  []interface{} `json:"params,omitempty"`
}

// CostCenterOrder is the order direction
type CostCenterOrder *string

var (
	// CostCenterOrderDesc is descending
	CostCenterOrderDesc CostCenterOrder = pstrings.Pointer("DESC")
	// CostCenterOrderAsc is ascending
	CostCenterOrderAsc CostCenterOrder = pstrings.Pointer("ASC")
)

// CostCenterQueryInput is query input struct
type CostCenterQueryInput struct {
	First   *int64           `json:"first,omitempty"`
	Last    *int64           `json:"last,omitempty"`
	Before  *string          `json:"before,omitempty"`
	After   *string          `json:"after,omitempty"`
	Query   *CostCenterQuery `json:"query,omitempty"`
	OrderBy *string          `json:"orderBy,omitempty"`
	Order   CostCenterOrder  `json:"order,omitempty"`
	NoCache bool             `json:"nocache,omitempty"`
}

// NewCostCenterQuery is a convenience for building a *CostCenterQuery
func NewCostCenterQuery(params ...interface{}) *CostCenterQueryInput {
	if len(params)%2 != 0 {
		panic("incorrect number of arguments passed")
	}
	q := &CostCenterQuery{
		Filters: make([]string, 0),
		Params:  make([]interface{}, 0),
	}
	for i := 0; i < len(params); i += 2 {
		q.Filters = append(q.Filters, params[i].(string))
		q.Params = append(q.Params, params[i+1])
	}
	return &CostCenterQueryInput{
		Query: q,
	}
}

// FindCostCenter will query an CostCenter by id
func FindCostCenter(client graphql.Client, id string) (*CostCenter, error) {
	variables := make(graphql.Variables)
	variables["id"] = id
	var sb strings.Builder
	sb.WriteString("query GoCostCenterQuery($id: ID) {\n")
	sb.WriteString("\tcustomer {\n")
	sb.WriteString("\t\tCostCenter(_id: $id) {\n")
	sb.WriteString(getCostCenterQueryFields())
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res QueryOneCostCenterData
	if err := client.Query(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	if res.Data != nil {
		return res.Data.Object, nil
	}
	return nil, nil
}

// FindCostCenterWithoutCache will query an CostCenter by id avoiding the cache
func FindCostCenterWithoutCache(client graphql.Client, id string) (*CostCenter, error) {
	variables := make(graphql.Variables)
	variables["id"] = id
	variables["nocache"] = true
	var sb strings.Builder
	sb.WriteString("query GoCostCenterQuery($id: ID, $nocache: Boolean) {\n")
	sb.WriteString("\tcustomer {\n")
	sb.WriteString("\t\tCostCenter(_id: $id, nocache: $nocache) {\n")
	sb.WriteString(getCostCenterQueryFields())
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res QueryOneCostCenterData
	if err := client.Query(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	if res.Data != nil {
		return res.Data.Object, nil
	}
	return nil, nil
}

// FindCostCenters will query for any CostCenters matching the query
func FindCostCenters(client graphql.Client, input *CostCenterQueryInput) (*CostCenterConnection, error) {
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
	sb.WriteString("query GoCostCenterQueryMany($first: Int, $last: Int, $before: Cursor, $after: Cursor, $query: QueryInput, $order: SortOrderEnum, $orderBy: CustomerCostCenterColumnEnum, $nocache: Boolean) {\n")
	sb.WriteString("\tcustomer {\n")
	sb.WriteString("\t\tCostCenters(first: $first last: $last before: $before after: $after query: $query orderBy: $orderBy order: $order nocache: $nocache) {\n")
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
	sb.WriteString(getCostCenterQueryFields())
	sb.WriteString("\t\t\t\t}\n")
	sb.WriteString("\t\t\t}\n")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res QueryManyCostCenterData
	if err := client.Query(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	return res.Data.Object, nil
}

// FindCostCentersPaginatedCallback is a callback function for handling each page
type FindCostCentersPaginatedCallback func(conn *CostCenterConnection) (bool, error)

// FindCostCentersPaginated will query for any CostCenters matching the query and return each page callback
func FindCostCentersPaginated(client graphql.Client, query *CostCenterQuery, pageSize int64, callback FindCostCentersPaginatedCallback) error {
	input := &CostCenterQueryInput{
		First: &pageSize,
		Query: query,
	}
	for {
		res, err := FindCostCenters(client, input)
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

// UpdateCostCenterNode is a grapqhl update node
type UpdateCostCenterNode struct {
	Object *CostCenter `json:"updateCostCenter,omitempty"`
}

// UpdateCostCenterData is a grapqhl update data node
type UpdateCostCenterData struct {
	Data *UpdateCostCenterNode `json:"customer,omitempty"`
}

// ExecCostCenterUpdateMutation returns a graphql update mutation result for CostCenter
func ExecCostCenterUpdateMutation(client graphql.Client, id string, input graphql.Variables, upsert bool) (*CostCenter, error) {
	if !upsert && id == "" {
		return nil, fmt.Errorf("id is required with upsert false")
	}
	variables := make(graphql.Variables)
	variables["id"] = id
	variables["upsert"] = upsert
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation GoCostCenterUpdateMutation($id: String, $input: UpdateCustomerCostCenterInput, $upsert: Boolean) {\n")
	sb.WriteString("\tcustomer {\n")
	sb.WriteString("\t\tupdateCostCenter(_id: $id, input: $input, upsert: $upsert) {\n")
	sb.WriteString(getCostCenterQueryFields())
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateCostCenterData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	return res.Data.Object, nil
}

// ExecCostCenterSilentUpdateMutation returns a graphql update mutation result for CostCenter
func ExecCostCenterSilentUpdateMutation(client graphql.Client, id string, input graphql.Variables, upsert bool) error {
	if !upsert && id == "" {
		return fmt.Errorf("id is required with upsert false")
	}
	variables := make(graphql.Variables)
	variables["id"] = id
	variables["upsert"] = upsert
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation GoCostCenterUpdateMutation($id: String, $input: UpdateCustomerCostCenterInput, $upsert: Boolean) {\n")
	sb.WriteString("\tcustomer {\n")
	sb.WriteString("\t\tupdateCostCenter(_id: $id, input: $input, upsert: $upsert) {\n")
	sb.WriteString("\t\t\t_id")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateCostCenterData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return err
	}
	return nil
}

// ExecCostCenterDeleteMutation executes a graphql delete mutation for CostCenter
func ExecCostCenterDeleteMutation(client graphql.Client, id string) error {
	variables := make(graphql.Variables)
	variables["id"] = id
	var sb strings.Builder
	sb.WriteString("mutation GoCostCenterDeleteMutation($id: String!) {\n")
	sb.WriteString("\tcustomer {\n")
	sb.WriteString("\t\tdeleteCostCenter(_id: $id)\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res interface{}
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return err
	}
	return nil
}

func CreateCostCenter(client graphql.Client, model CostCenter) error {
	variables := make(graphql.Variables)
	input := model.ToMap()
	delete(input, "hashcode")
	delete(input, "customer_id")
	delete(input, "created_ts")
	delete(input, "id")
	input["_id"] = model.GetID()
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation GoCreateCostCenter($input: CreateCustomerCostCenterInput!) {\n")
	sb.WriteString("\tcustomer {\n")
	sb.WriteString("\t\tcreateCostCenter(input: $input) {\n")
	sb.WriteString("\t\t\t_id")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateCostCenterData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return err
	}
	return nil
}

// CostCenterPartial is a partial struct for upsert mutations for CostCenter
type CostCenterPartial struct {
	// Active whether the cost center is tracked in pinpoint
	Active *bool `json:"active,omitempty"`
	// Cost the cost value of the cost center
	Cost *float64 `json:"cost,omitempty"`
	// Description the description for the cost center
	Description *string `json:"description,omitempty"`
	// Name the name of the cost center
	Name *string `json:"name,omitempty"`
}

var _ datamodel.PartialModel = (*CostCenterPartial)(nil)

// GetModelName returns the name of the model
func (o *CostCenterPartial) GetModelName() datamodel.ModelNameType {
	return CostCenterModelName
}

// ToMap returns the object as a map
func (o *CostCenterPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"active":      toCostCenterObject(o.Active, true),
		"cost":        toCostCenterObject(o.Cost, true),
		"description": toCostCenterObject(o.Description, true),
		"name":        toCostCenterObject(o.Name, true),
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
func (o *CostCenterPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *CostCenterPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *CostCenterPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *CostCenterPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *CostCenterPartial) FromMap(kv map[string]interface{}) {
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
	if val, ok := kv["cost"].(*float64); ok {
		o.Cost = val
	} else if val, ok := kv["cost"].(float64); ok {
		o.Cost = &val
	} else {
		if val, ok := kv["cost"]; ok {
			if val == nil {
				o.Cost = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["float64"]
				}
				o.Cost = number.Float64Pointer(number.ToFloat64Any(val))
			}
		}
	}
	if val, ok := kv["description"].(*string); ok {
		o.Description = val
	} else if val, ok := kv["description"].(string); ok {
		o.Description = &val
	} else {
		if val, ok := kv["description"]; ok {
			if val == nil {
				o.Description = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Description = pstrings.Pointer(fmt.Sprintf("%v", val))
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
	o.setDefaults(false)
}
