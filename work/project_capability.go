// DO NOT EDIT -- generated code

// Package work - the system which contains project work
package work

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
	// ProjectCapabilityTopic is the default topic name
	ProjectCapabilityTopic datamodel.TopicNameType = "work_ProjectCapability"

	// ProjectCapabilityTable is the default table name
	ProjectCapabilityTable datamodel.ModelNameType = "work_projectcapability"

	// ProjectCapabilityModelName is the model name
	ProjectCapabilityModelName datamodel.ModelNameType = "work.ProjectCapability"
)

const (
	// ProjectCapabilityModelActiveColumn is the column json value active
	ProjectCapabilityModelActiveColumn = "active"
	// ProjectCapabilityModelAttachmentsColumn is the column json value attachments
	ProjectCapabilityModelAttachmentsColumn = "attachments"
	// ProjectCapabilityModelChangeLogsColumn is the column json value change_logs
	ProjectCapabilityModelChangeLogsColumn = "change_logs"
	// ProjectCapabilityModelCustomerIDColumn is the column json value customer_id
	ProjectCapabilityModelCustomerIDColumn = "customer_id"
	// ProjectCapabilityModelDueDatesColumn is the column json value due_dates
	ProjectCapabilityModelDueDatesColumn = "due_dates"
	// ProjectCapabilityModelEpicsColumn is the column json value epics
	ProjectCapabilityModelEpicsColumn = "epics"
	// ProjectCapabilityModelIDColumn is the column json value id
	ProjectCapabilityModelIDColumn = "id"
	// ProjectCapabilityModelInProgressStatesColumn is the column json value in_progress_states
	ProjectCapabilityModelInProgressStatesColumn = "in_progress_states"
	// ProjectCapabilityModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	ProjectCapabilityModelIntegrationInstanceIDColumn = "integration_instance_id"
	// ProjectCapabilityModelKanbanBoardsColumn is the column json value kanban_boards
	ProjectCapabilityModelKanbanBoardsColumn = "kanban_boards"
	// ProjectCapabilityModelLinkedIssuesColumn is the column json value linked_issues
	ProjectCapabilityModelLinkedIssuesColumn = "linked_issues"
	// ProjectCapabilityModelParentsColumn is the column json value parents
	ProjectCapabilityModelParentsColumn = "parents"
	// ProjectCapabilityModelPrioritiesColumn is the column json value priorities
	ProjectCapabilityModelPrioritiesColumn = "priorities"
	// ProjectCapabilityModelProjectIDColumn is the column json value project_id
	ProjectCapabilityModelProjectIDColumn = "project_id"
	// ProjectCapabilityModelRefIDColumn is the column json value ref_id
	ProjectCapabilityModelRefIDColumn = "ref_id"
	// ProjectCapabilityModelRefTypeColumn is the column json value ref_type
	ProjectCapabilityModelRefTypeColumn = "ref_type"
	// ProjectCapabilityModelResolutionsColumn is the column json value resolutions
	ProjectCapabilityModelResolutionsColumn = "resolutions"
	// ProjectCapabilityModelSprintsColumn is the column json value sprints
	ProjectCapabilityModelSprintsColumn = "sprints"
	// ProjectCapabilityModelStoryPointsColumn is the column json value story_points
	ProjectCapabilityModelStoryPointsColumn = "story_points"
	// ProjectCapabilityModelUpdatedAtColumn is the column json value updated_ts
	ProjectCapabilityModelUpdatedAtColumn = "updated_ts"
)

// ProjectCapability the project capabilities for a given integration
type ProjectCapability struct {
	// Active indicates that this model is displayed in a source system, false if the model is deleted
	Active bool `json:"active" codec:"active" bson:"active" yaml:"active" faker:"-"`
	// Attachments if the project supports issue attachments
	Attachments bool `json:"attachments" codec:"attachments" bson:"attachments" yaml:"attachments" faker:"-"`
	// ChangeLogs if the project supports issue change logs for issue state transitions
	ChangeLogs bool `json:"change_logs" codec:"change_logs" bson:"change_logs" yaml:"change_logs" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// DueDates if the project supports the notion of due dates for an issue
	DueDates bool `json:"due_dates" codec:"due_dates" bson:"due_dates" yaml:"due_dates" faker:"-"`
	// Epics if the project supports the notion of agile epics
	Epics bool `json:"epics" codec:"epics" bson:"epics" yaml:"epics" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// InProgressStates if the project supports the ability to transition from in progress states
	InProgressStates bool `json:"in_progress_states" codec:"in_progress_states" bson:"in_progress_states" yaml:"in_progress_states" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// KanbanBoards if the project supports agile kanban boards
	KanbanBoards bool `json:"kanban_boards" codec:"kanban_boards" bson:"kanban_boards" yaml:"kanban_boards" faker:"-"`
	// LinkedIssues if the project supports issue linking such as blocked issues, related issues, etc
	LinkedIssues bool `json:"linked_issues" codec:"linked_issues" bson:"linked_issues" yaml:"linked_issues" faker:"-"`
	// Parents if the project supports parent-child relationships between issues
	Parents bool `json:"parents" codec:"parents" bson:"parents" yaml:"parents" faker:"-"`
	// Priorities if the project supports requiring a priority when creating an issue
	Priorities bool `json:"priorities" codec:"priorities" bson:"priorities" yaml:"priorities" faker:"-"`
	// ProjectID the project id
	ProjectID string `json:"project_id" codec:"project_id" bson:"project_id" yaml:"project_id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// Resolutions if the project supports requiring a resolution when closing an issue
	Resolutions bool `json:"resolutions" codec:"resolutions" bson:"resolutions" yaml:"resolutions" faker:"-"`
	// Sprints if the project supports agile sprints
	Sprints bool `json:"sprints" codec:"sprints" bson:"sprints" yaml:"sprints" faker:"-"`
	// StoryPoints if the project supports the notion of agile story points
	StoryPoints bool `json:"story_points" codec:"story_points" bson:"story_points" yaml:"story_points" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*ProjectCapability)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*ProjectCapability)(nil)

func toProjectCapabilityObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ProjectCapability:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of ProjectCapability
func (o *ProjectCapability) String() string {
	return fmt.Sprintf("work.ProjectCapability<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *ProjectCapability) GetTopicName() datamodel.TopicNameType {
	return ProjectCapabilityTopic
}

// GetStreamName returns the name of the stream
func (o *ProjectCapability) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *ProjectCapability) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *ProjectCapability) GetModelName() datamodel.ModelNameType {
	return ProjectCapabilityModelName
}

// NewProjectCapabilityID provides a template for generating an ID field for ProjectCapability
func NewProjectCapabilityID(customerID string, ProjectID string) string {
	return hash.Values(customerID, ProjectID)
}

func (o *ProjectCapability) setDefaults(frommap bool) {

	if o.ID == "" {
		o.ID = hash.Values(o.CustomerID, o.ProjectID)
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *ProjectCapability) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *ProjectCapability) GetTopicKey() string {
	var i interface{} = o.ID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *ProjectCapability) GetTimestamp() time.Time {
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
	panic("not sure how to handle the date time format for ProjectCapability")
}

// GetRefID returns the RefID for the object
func (o *ProjectCapability) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *ProjectCapability) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *ProjectCapability) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *ProjectCapability) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *ProjectCapability) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *ProjectCapability) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = ProjectCapabilityModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *ProjectCapability) GetTopicConfig() *datamodel.ModelTopicConfig {
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
func (o *ProjectCapability) GetCustomerID() string {

	return o.CustomerID

}

// SetCustomerID will return the customer_id
func (o *ProjectCapability) SetCustomerID(id string) {

	o.CustomerID = id

}

// GetRefType will return the ref_type
func (o *ProjectCapability) GetRefType() string {

	return o.RefType

}

// SetRefType will return the ref_type
func (o *ProjectCapability) SetRefType(t string) {

	o.RefType = t

}

// Clone returns an exact copy of ProjectCapability
func (o *ProjectCapability) Clone() datamodel.Model {
	c := new(ProjectCapability)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *ProjectCapability) Anon() datamodel.Model {
	c := new(ProjectCapability)
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
func (o *ProjectCapability) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *ProjectCapability) UnmarshalJSON(data []byte) error {
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
func (o *ProjectCapability) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *ProjectCapability) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two ProjectCapability objects are equal
func (o *ProjectCapability) IsEqual(other *ProjectCapability) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *ProjectCapability) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"active":                  toProjectCapabilityObject(o.Active, false),
		"attachments":             toProjectCapabilityObject(o.Attachments, false),
		"change_logs":             toProjectCapabilityObject(o.ChangeLogs, false),
		"customer_id":             toProjectCapabilityObject(o.CustomerID, false),
		"due_dates":               toProjectCapabilityObject(o.DueDates, false),
		"epics":                   toProjectCapabilityObject(o.Epics, false),
		"id":                      toProjectCapabilityObject(o.ID, false),
		"in_progress_states":      toProjectCapabilityObject(o.InProgressStates, false),
		"integration_instance_id": toProjectCapabilityObject(o.IntegrationInstanceID, true),
		"kanban_boards":           toProjectCapabilityObject(o.KanbanBoards, false),
		"linked_issues":           toProjectCapabilityObject(o.LinkedIssues, false),
		"parents":                 toProjectCapabilityObject(o.Parents, false),
		"priorities":              toProjectCapabilityObject(o.Priorities, false),
		"project_id":              toProjectCapabilityObject(o.ProjectID, false),
		"ref_id":                  toProjectCapabilityObject(o.RefID, false),
		"ref_type":                toProjectCapabilityObject(o.RefType, false),
		"resolutions":             toProjectCapabilityObject(o.Resolutions, false),
		"sprints":                 toProjectCapabilityObject(o.Sprints, false),
		"story_points":            toProjectCapabilityObject(o.StoryPoints, false),
		"updated_ts":              toProjectCapabilityObject(o.UpdatedAt, false),
		"hashcode":                toProjectCapabilityObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *ProjectCapability) FromMap(kv map[string]interface{}) {

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
	if val, ok := kv["attachments"].(bool); ok {
		o.Attachments = val
	} else {
		if val, ok := kv["attachments"]; ok {
			if val == nil {
				o.Attachments = false
			} else {
				o.Attachments = number.ToBoolAny(val)
			}
		}
	}
	if val, ok := kv["change_logs"].(bool); ok {
		o.ChangeLogs = val
	} else {
		if val, ok := kv["change_logs"]; ok {
			if val == nil {
				o.ChangeLogs = false
			} else {
				o.ChangeLogs = number.ToBoolAny(val)
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
	if val, ok := kv["due_dates"].(bool); ok {
		o.DueDates = val
	} else {
		if val, ok := kv["due_dates"]; ok {
			if val == nil {
				o.DueDates = false
			} else {
				o.DueDates = number.ToBoolAny(val)
			}
		}
	}
	if val, ok := kv["epics"].(bool); ok {
		o.Epics = val
	} else {
		if val, ok := kv["epics"]; ok {
			if val == nil {
				o.Epics = false
			} else {
				o.Epics = number.ToBoolAny(val)
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
	if val, ok := kv["in_progress_states"].(bool); ok {
		o.InProgressStates = val
	} else {
		if val, ok := kv["in_progress_states"]; ok {
			if val == nil {
				o.InProgressStates = false
			} else {
				o.InProgressStates = number.ToBoolAny(val)
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
	if val, ok := kv["kanban_boards"].(bool); ok {
		o.KanbanBoards = val
	} else {
		if val, ok := kv["kanban_boards"]; ok {
			if val == nil {
				o.KanbanBoards = false
			} else {
				o.KanbanBoards = number.ToBoolAny(val)
			}
		}
	}
	if val, ok := kv["linked_issues"].(bool); ok {
		o.LinkedIssues = val
	} else {
		if val, ok := kv["linked_issues"]; ok {
			if val == nil {
				o.LinkedIssues = false
			} else {
				o.LinkedIssues = number.ToBoolAny(val)
			}
		}
	}
	if val, ok := kv["parents"].(bool); ok {
		o.Parents = val
	} else {
		if val, ok := kv["parents"]; ok {
			if val == nil {
				o.Parents = false
			} else {
				o.Parents = number.ToBoolAny(val)
			}
		}
	}
	if val, ok := kv["priorities"].(bool); ok {
		o.Priorities = val
	} else {
		if val, ok := kv["priorities"]; ok {
			if val == nil {
				o.Priorities = false
			} else {
				o.Priorities = number.ToBoolAny(val)
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
	if val, ok := kv["resolutions"].(bool); ok {
		o.Resolutions = val
	} else {
		if val, ok := kv["resolutions"]; ok {
			if val == nil {
				o.Resolutions = false
			} else {
				o.Resolutions = number.ToBoolAny(val)
			}
		}
	}
	if val, ok := kv["sprints"].(bool); ok {
		o.Sprints = val
	} else {
		if val, ok := kv["sprints"]; ok {
			if val == nil {
				o.Sprints = false
			} else {
				o.Sprints = number.ToBoolAny(val)
			}
		}
	}
	if val, ok := kv["story_points"].(bool); ok {
		o.StoryPoints = val
	} else {
		if val, ok := kv["story_points"]; ok {
			if val == nil {
				o.StoryPoints = false
			} else {
				o.StoryPoints = number.ToBoolAny(val)
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
func (o *ProjectCapability) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Active)
	args = append(args, o.Attachments)
	args = append(args, o.ChangeLogs)
	args = append(args, o.CustomerID)
	args = append(args, o.DueDates)
	args = append(args, o.Epics)
	args = append(args, o.ID)
	args = append(args, o.InProgressStates)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.KanbanBoards)
	args = append(args, o.LinkedIssues)
	args = append(args, o.Parents)
	args = append(args, o.Priorities)
	args = append(args, o.ProjectID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.Resolutions)
	args = append(args, o.Sprints)
	args = append(args, o.StoryPoints)
	args = append(args, o.UpdatedAt)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// SetIntegrationInstanceID will set the integration instance ID
func (o *ProjectCapability) SetIntegrationInstanceID(id string) {
	if id == "" {
		o.IntegrationInstanceID = nil
	} else {
		o.IntegrationInstanceID = &id
	}
}

// GetIntegrationInstanceID will return the integration instance ID
func (o *ProjectCapability) GetIntegrationInstanceID() *string {
	return o.IntegrationInstanceID
}

// ProjectCapabilityPartial is a partial struct for upsert mutations for ProjectCapability
type ProjectCapabilityPartial struct {
	// Active indicates that this model is displayed in a source system, false if the model is deleted
	Active *bool `json:"active,omitempty"`
	// Attachments if the project supports issue attachments
	Attachments *bool `json:"attachments,omitempty"`
	// ChangeLogs if the project supports issue change logs for issue state transitions
	ChangeLogs *bool `json:"change_logs,omitempty"`
	// DueDates if the project supports the notion of due dates for an issue
	DueDates *bool `json:"due_dates,omitempty"`
	// Epics if the project supports the notion of agile epics
	Epics *bool `json:"epics,omitempty"`
	// InProgressStates if the project supports the ability to transition from in progress states
	InProgressStates *bool `json:"in_progress_states,omitempty"`
	// KanbanBoards if the project supports agile kanban boards
	KanbanBoards *bool `json:"kanban_boards,omitempty"`
	// LinkedIssues if the project supports issue linking such as blocked issues, related issues, etc
	LinkedIssues *bool `json:"linked_issues,omitempty"`
	// Parents if the project supports parent-child relationships between issues
	Parents *bool `json:"parents,omitempty"`
	// Priorities if the project supports requiring a priority when creating an issue
	Priorities *bool `json:"priorities,omitempty"`
	// ProjectID the project id
	ProjectID *string `json:"project_id,omitempty"`
	// Resolutions if the project supports requiring a resolution when closing an issue
	Resolutions *bool `json:"resolutions,omitempty"`
	// Sprints if the project supports agile sprints
	Sprints *bool `json:"sprints,omitempty"`
	// StoryPoints if the project supports the notion of agile story points
	StoryPoints *bool `json:"story_points,omitempty"`
}

var _ datamodel.PartialModel = (*ProjectCapabilityPartial)(nil)

// GetModelName returns the name of the model
func (o *ProjectCapabilityPartial) GetModelName() datamodel.ModelNameType {
	return ProjectCapabilityModelName
}

// ToMap returns the object as a map
func (o *ProjectCapabilityPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"active":             toProjectCapabilityObject(o.Active, true),
		"attachments":        toProjectCapabilityObject(o.Attachments, true),
		"change_logs":        toProjectCapabilityObject(o.ChangeLogs, true),
		"due_dates":          toProjectCapabilityObject(o.DueDates, true),
		"epics":              toProjectCapabilityObject(o.Epics, true),
		"in_progress_states": toProjectCapabilityObject(o.InProgressStates, true),
		"kanban_boards":      toProjectCapabilityObject(o.KanbanBoards, true),
		"linked_issues":      toProjectCapabilityObject(o.LinkedIssues, true),
		"parents":            toProjectCapabilityObject(o.Parents, true),
		"priorities":         toProjectCapabilityObject(o.Priorities, true),
		"project_id":         toProjectCapabilityObject(o.ProjectID, true),
		"resolutions":        toProjectCapabilityObject(o.Resolutions, true),
		"sprints":            toProjectCapabilityObject(o.Sprints, true),
		"story_points":       toProjectCapabilityObject(o.StoryPoints, true),
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
func (o *ProjectCapabilityPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *ProjectCapabilityPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *ProjectCapabilityPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *ProjectCapabilityPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *ProjectCapabilityPartial) FromMap(kv map[string]interface{}) {
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
	if val, ok := kv["attachments"].(*bool); ok {
		o.Attachments = val
	} else if val, ok := kv["attachments"].(bool); ok {
		o.Attachments = &val
	} else {
		if val, ok := kv["attachments"]; ok {
			if val == nil {
				o.Attachments = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Attachments = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}
	if val, ok := kv["change_logs"].(*bool); ok {
		o.ChangeLogs = val
	} else if val, ok := kv["change_logs"].(bool); ok {
		o.ChangeLogs = &val
	} else {
		if val, ok := kv["change_logs"]; ok {
			if val == nil {
				o.ChangeLogs = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.ChangeLogs = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}
	if val, ok := kv["due_dates"].(*bool); ok {
		o.DueDates = val
	} else if val, ok := kv["due_dates"].(bool); ok {
		o.DueDates = &val
	} else {
		if val, ok := kv["due_dates"]; ok {
			if val == nil {
				o.DueDates = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.DueDates = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}
	if val, ok := kv["epics"].(*bool); ok {
		o.Epics = val
	} else if val, ok := kv["epics"].(bool); ok {
		o.Epics = &val
	} else {
		if val, ok := kv["epics"]; ok {
			if val == nil {
				o.Epics = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Epics = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}
	if val, ok := kv["in_progress_states"].(*bool); ok {
		o.InProgressStates = val
	} else if val, ok := kv["in_progress_states"].(bool); ok {
		o.InProgressStates = &val
	} else {
		if val, ok := kv["in_progress_states"]; ok {
			if val == nil {
				o.InProgressStates = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.InProgressStates = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}
	if val, ok := kv["kanban_boards"].(*bool); ok {
		o.KanbanBoards = val
	} else if val, ok := kv["kanban_boards"].(bool); ok {
		o.KanbanBoards = &val
	} else {
		if val, ok := kv["kanban_boards"]; ok {
			if val == nil {
				o.KanbanBoards = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.KanbanBoards = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}
	if val, ok := kv["linked_issues"].(*bool); ok {
		o.LinkedIssues = val
	} else if val, ok := kv["linked_issues"].(bool); ok {
		o.LinkedIssues = &val
	} else {
		if val, ok := kv["linked_issues"]; ok {
			if val == nil {
				o.LinkedIssues = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.LinkedIssues = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}
	if val, ok := kv["parents"].(*bool); ok {
		o.Parents = val
	} else if val, ok := kv["parents"].(bool); ok {
		o.Parents = &val
	} else {
		if val, ok := kv["parents"]; ok {
			if val == nil {
				o.Parents = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Parents = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}
	if val, ok := kv["priorities"].(*bool); ok {
		o.Priorities = val
	} else if val, ok := kv["priorities"].(bool); ok {
		o.Priorities = &val
	} else {
		if val, ok := kv["priorities"]; ok {
			if val == nil {
				o.Priorities = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Priorities = number.BoolPointer(number.ToBoolAny(val))
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
	if val, ok := kv["resolutions"].(*bool); ok {
		o.Resolutions = val
	} else if val, ok := kv["resolutions"].(bool); ok {
		o.Resolutions = &val
	} else {
		if val, ok := kv["resolutions"]; ok {
			if val == nil {
				o.Resolutions = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Resolutions = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}
	if val, ok := kv["sprints"].(*bool); ok {
		o.Sprints = val
	} else if val, ok := kv["sprints"].(bool); ok {
		o.Sprints = &val
	} else {
		if val, ok := kv["sprints"]; ok {
			if val == nil {
				o.Sprints = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Sprints = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}
	if val, ok := kv["story_points"].(*bool); ok {
		o.StoryPoints = val
	} else if val, ok := kv["story_points"].(bool); ok {
		o.StoryPoints = &val
	} else {
		if val, ok := kv["story_points"]; ok {
			if val == nil {
				o.StoryPoints = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.StoryPoints = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}
	o.setDefaults(false)
}
