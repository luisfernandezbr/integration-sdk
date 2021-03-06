// DO NOT EDIT -- generated code

// Package agent -
package agent

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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

const (

	// IntegrationMutationRequestTable is the default table name
	IntegrationMutationRequestTable datamodel.ModelNameType = "agent_integrationmutationrequest"

	// IntegrationMutationRequestModelName is the model name
	IntegrationMutationRequestModelName datamodel.ModelNameType = "agent.IntegrationMutationRequest"
)

const (
	// IntegrationMutationRequestModelActionColumn is the column json value action
	IntegrationMutationRequestModelActionColumn = "action"
	// IntegrationMutationRequestModelAgentRequestSentDateColumn is the column json value agent_request_sent_date
	IntegrationMutationRequestModelAgentRequestSentDateColumn = "agent_request_sent_date"
	// IntegrationMutationRequestModelAgentRequestSentDateEpochColumn is the column json value epoch
	IntegrationMutationRequestModelAgentRequestSentDateEpochColumn = "epoch"
	// IntegrationMutationRequestModelAgentRequestSentDateOffsetColumn is the column json value offset
	IntegrationMutationRequestModelAgentRequestSentDateOffsetColumn = "offset"
	// IntegrationMutationRequestModelAgentRequestSentDateRfc3339Column is the column json value rfc3339
	IntegrationMutationRequestModelAgentRequestSentDateRfc3339Column = "rfc3339"
	// IntegrationMutationRequestModelAuthorizationColumn is the column json value authorization
	IntegrationMutationRequestModelAuthorizationColumn = "authorization"
	// IntegrationMutationRequestModelAuthorizationAccessTokenColumn is the column json value access_token
	IntegrationMutationRequestModelAuthorizationAccessTokenColumn = "access_token"
	// IntegrationMutationRequestModelAuthorizationRefreshTokenColumn is the column json value refresh_token
	IntegrationMutationRequestModelAuthorizationRefreshTokenColumn = "refresh_token"
	// IntegrationMutationRequestModelAuthorizationURLColumn is the column json value url
	IntegrationMutationRequestModelAuthorizationURLColumn = "url"
	// IntegrationMutationRequestModelCustomerIDColumn is the column json value customer_id
	IntegrationMutationRequestModelCustomerIDColumn = "customer_id"
	// IntegrationMutationRequestModelDataColumn is the column json value data
	IntegrationMutationRequestModelDataColumn = "data"
	// IntegrationMutationRequestModelIDColumn is the column json value id
	IntegrationMutationRequestModelIDColumn = "id"
	// IntegrationMutationRequestModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	IntegrationMutationRequestModelIntegrationInstanceIDColumn = "integration_instance_id"
	// IntegrationMutationRequestModelIntegrationNameColumn is the column json value integration_name
	IntegrationMutationRequestModelIntegrationNameColumn = "integration_name"
	// IntegrationMutationRequestModelJobIDColumn is the column json value job_id
	IntegrationMutationRequestModelJobIDColumn = "job_id"
	// IntegrationMutationRequestModelRefIDColumn is the column json value ref_id
	IntegrationMutationRequestModelRefIDColumn = "ref_id"
	// IntegrationMutationRequestModelRefTypeColumn is the column json value ref_type
	IntegrationMutationRequestModelRefTypeColumn = "ref_type"
	// IntegrationMutationRequestModelRequestDateColumn is the column json value request_date
	IntegrationMutationRequestModelRequestDateColumn = "request_date"
	// IntegrationMutationRequestModelRequestDateEpochColumn is the column json value epoch
	IntegrationMutationRequestModelRequestDateEpochColumn = "epoch"
	// IntegrationMutationRequestModelRequestDateOffsetColumn is the column json value offset
	IntegrationMutationRequestModelRequestDateOffsetColumn = "offset"
	// IntegrationMutationRequestModelRequestDateRfc3339Column is the column json value rfc3339
	IntegrationMutationRequestModelRequestDateRfc3339Column = "rfc3339"
	// IntegrationMutationRequestModelSystemTypeColumn is the column json value system_type
	IntegrationMutationRequestModelSystemTypeColumn = "system_type"
	// IntegrationMutationRequestModelUUIDColumn is the column json value uuid
	IntegrationMutationRequestModelUUIDColumn = "uuid"
	// IntegrationMutationRequestModelWebappRequestDateColumn is the column json value webapp_request_date
	IntegrationMutationRequestModelWebappRequestDateColumn = "webapp_request_date"
	// IntegrationMutationRequestModelWebappRequestDateEpochColumn is the column json value epoch
	IntegrationMutationRequestModelWebappRequestDateEpochColumn = "epoch"
	// IntegrationMutationRequestModelWebappRequestDateOffsetColumn is the column json value offset
	IntegrationMutationRequestModelWebappRequestDateOffsetColumn = "offset"
	// IntegrationMutationRequestModelWebappRequestDateRfc3339Column is the column json value rfc3339
	IntegrationMutationRequestModelWebappRequestDateRfc3339Column = "rfc3339"
)

// IntegrationMutationRequestAction is the enumeration type for action
type IntegrationMutationRequestAction int32

// toIntegrationMutationRequestActionPointer is the enumeration pointer type for action
func toIntegrationMutationRequestActionPointer(v int32) *IntegrationMutationRequestAction {
	nv := IntegrationMutationRequestAction(v)
	return &nv
}

// toIntegrationMutationRequestActionEnum is the enumeration pointer wrapper for action
func toIntegrationMutationRequestActionEnum(v *IntegrationMutationRequestAction) string {
	if v == nil {
		return toIntegrationMutationRequestActionPointer(0).String()
	}
	return v.String()
}

// UnmarshalBSONValue for unmarshaling value
func (v *IntegrationMutationRequestAction) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = IntegrationMutationRequestAction(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "ISSUE_ADD_COMMENT":
			*v = IntegrationMutationRequestAction(0)
		case "ISSUE_SET_TITLE":
			*v = IntegrationMutationRequestAction(1)
		case "ISSUE_SET_STATUS":
			*v = IntegrationMutationRequestAction(2)
		case "ISSUE_SET_PRIORITY":
			*v = IntegrationMutationRequestAction(3)
		case "ISSUE_SET_ASSIGNEE":
			*v = IntegrationMutationRequestAction(4)
		case "ISSUE_GET_TRANSITIONS":
			*v = IntegrationMutationRequestAction(5)
		case "PR_SET_TITLE":
			*v = IntegrationMutationRequestAction(6)
		case "PR_SET_DESCRIPTION":
			*v = IntegrationMutationRequestAction(7)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v *IntegrationMutationRequestAction) UnmarshalJSON(buf []byte) error {
	var val string
	if err := json.Unmarshal(buf, &val); err != nil {
		return err
	}
	switch val {
	case "ISSUE_ADD_COMMENT":
		*v = 0
	case "ISSUE_SET_TITLE":
		*v = 1
	case "ISSUE_SET_STATUS":
		*v = 2
	case "ISSUE_SET_PRIORITY":
		*v = 3
	case "ISSUE_SET_ASSIGNEE":
		*v = 4
	case "ISSUE_GET_TRANSITIONS":
		*v = 5
	case "PR_SET_TITLE":
		*v = 6
	case "PR_SET_DESCRIPTION":
		*v = 7
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v IntegrationMutationRequestAction) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("ISSUE_ADD_COMMENT")
	case 1:
		return json.Marshal("ISSUE_SET_TITLE")
	case 2:
		return json.Marshal("ISSUE_SET_STATUS")
	case 3:
		return json.Marshal("ISSUE_SET_PRIORITY")
	case 4:
		return json.Marshal("ISSUE_SET_ASSIGNEE")
	case 5:
		return json.Marshal("ISSUE_GET_TRANSITIONS")
	case 6:
		return json.Marshal("PR_SET_TITLE")
	case 7:
		return json.Marshal("PR_SET_DESCRIPTION")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for Action
func (v IntegrationMutationRequestAction) String() string {
	switch int32(v) {
	case 0:
		return "ISSUE_ADD_COMMENT"
	case 1:
		return "ISSUE_SET_TITLE"
	case 2:
		return "ISSUE_SET_STATUS"
	case 3:
		return "ISSUE_SET_PRIORITY"
	case 4:
		return "ISSUE_SET_ASSIGNEE"
	case 5:
		return "ISSUE_GET_TRANSITIONS"
	case 6:
		return "PR_SET_TITLE"
	case 7:
		return "PR_SET_DESCRIPTION"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *IntegrationMutationRequestAction) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case IntegrationMutationRequestAction:
		*v = val
	case int32:
		*v = IntegrationMutationRequestAction(int32(val))
	case int:
		*v = IntegrationMutationRequestAction(int32(val))
	case string:
		switch val {
		case "ISSUE_ADD_COMMENT":
			*v = IntegrationMutationRequestAction(0)
		case "ISSUE_SET_TITLE":
			*v = IntegrationMutationRequestAction(1)
		case "ISSUE_SET_STATUS":
			*v = IntegrationMutationRequestAction(2)
		case "ISSUE_SET_PRIORITY":
			*v = IntegrationMutationRequestAction(3)
		case "ISSUE_SET_ASSIGNEE":
			*v = IntegrationMutationRequestAction(4)
		case "ISSUE_GET_TRANSITIONS":
			*v = IntegrationMutationRequestAction(5)
		case "PR_SET_TITLE":
			*v = IntegrationMutationRequestAction(6)
		case "PR_SET_DESCRIPTION":
			*v = IntegrationMutationRequestAction(7)
		}
	}
	return nil
}

const (
	// IntegrationMutationRequestActionIssueAddComment is the enumeration value for issue_add_comment
	IntegrationMutationRequestActionIssueAddComment IntegrationMutationRequestAction = 0
	// IntegrationMutationRequestActionIssueSetTitle is the enumeration value for issue_set_title
	IntegrationMutationRequestActionIssueSetTitle IntegrationMutationRequestAction = 1
	// IntegrationMutationRequestActionIssueSetStatus is the enumeration value for issue_set_status
	IntegrationMutationRequestActionIssueSetStatus IntegrationMutationRequestAction = 2
	// IntegrationMutationRequestActionIssueSetPriority is the enumeration value for issue_set_priority
	IntegrationMutationRequestActionIssueSetPriority IntegrationMutationRequestAction = 3
	// IntegrationMutationRequestActionIssueSetAssignee is the enumeration value for issue_set_assignee
	IntegrationMutationRequestActionIssueSetAssignee IntegrationMutationRequestAction = 4
	// IntegrationMutationRequestActionIssueGetTransitions is the enumeration value for issue_get_transitions
	IntegrationMutationRequestActionIssueGetTransitions IntegrationMutationRequestAction = 5
	// IntegrationMutationRequestActionPrSetTitle is the enumeration value for pr_set_title
	IntegrationMutationRequestActionPrSetTitle IntegrationMutationRequestAction = 6
	// IntegrationMutationRequestActionPrSetDescription is the enumeration value for pr_set_description
	IntegrationMutationRequestActionPrSetDescription IntegrationMutationRequestAction = 7
)

// IntegrationMutationRequestAgentRequestSentDate represents the object structure for agent_request_sent_date
type IntegrationMutationRequestAgentRequestSentDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIntegrationMutationRequestAgentRequestSentDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationMutationRequestAgentRequestSentDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *IntegrationMutationRequestAgentRequestSentDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIntegrationMutationRequestAgentRequestSentDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIntegrationMutationRequestAgentRequestSentDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIntegrationMutationRequestAgentRequestSentDateObject(o.Rfc3339, false),
	}
}

func (o *IntegrationMutationRequestAgentRequestSentDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationMutationRequestAgentRequestSentDate) FromMap(kv map[string]interface{}) {

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

// IntegrationMutationRequestAuthorization represents the object structure for authorization
type IntegrationMutationRequestAuthorization struct {
	// AccessToken Access token
	AccessToken string `json:"access_token" codec:"access_token" bson:"access_token" yaml:"access_token" faker:"-"`
	// RefreshToken Refresh token
	RefreshToken string `json:"refresh_token" codec:"refresh_token" bson:"refresh_token" yaml:"refresh_token" faker:"-"`
	// URL URL of instance if relevant
	URL *string `json:"url,omitempty" codec:"url,omitempty" bson:"url" yaml:"url,omitempty" faker:"-"`
}

func toIntegrationMutationRequestAuthorizationObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationMutationRequestAuthorization:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *IntegrationMutationRequestAuthorization) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// AccessToken Access token
		"access_token": toIntegrationMutationRequestAuthorizationObject(o.AccessToken, false),
		// RefreshToken Refresh token
		"refresh_token": toIntegrationMutationRequestAuthorizationObject(o.RefreshToken, false),
		// URL URL of instance if relevant
		"url": toIntegrationMutationRequestAuthorizationObject(o.URL, true),
	}
}

func (o *IntegrationMutationRequestAuthorization) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationMutationRequestAuthorization) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["access_token"].(string); ok {
		o.AccessToken = val
	} else {
		if val, ok := kv["access_token"]; ok {
			if val == nil {
				o.AccessToken = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.AccessToken = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["refresh_token"].(string); ok {
		o.RefreshToken = val
	} else {
		if val, ok := kv["refresh_token"]; ok {
			if val == nil {
				o.RefreshToken = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.RefreshToken = fmt.Sprintf("%v", val)
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

// IntegrationMutationRequestRequestDate represents the object structure for request_date
type IntegrationMutationRequestRequestDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIntegrationMutationRequestRequestDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationMutationRequestRequestDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *IntegrationMutationRequestRequestDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIntegrationMutationRequestRequestDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIntegrationMutationRequestRequestDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIntegrationMutationRequestRequestDateObject(o.Rfc3339, false),
	}
}

func (o *IntegrationMutationRequestRequestDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationMutationRequestRequestDate) FromMap(kv map[string]interface{}) {

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

// IntegrationMutationRequestSystemType is the enumeration type for system_type
type IntegrationMutationRequestSystemType int32

// toIntegrationMutationRequestSystemTypePointer is the enumeration pointer type for system_type
func toIntegrationMutationRequestSystemTypePointer(v int32) *IntegrationMutationRequestSystemType {
	nv := IntegrationMutationRequestSystemType(v)
	return &nv
}

// toIntegrationMutationRequestSystemTypeEnum is the enumeration pointer wrapper for system_type
func toIntegrationMutationRequestSystemTypeEnum(v *IntegrationMutationRequestSystemType) string {
	if v == nil {
		return toIntegrationMutationRequestSystemTypePointer(0).String()
	}
	return v.String()
}

// UnmarshalBSONValue for unmarshaling value
func (v *IntegrationMutationRequestSystemType) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = IntegrationMutationRequestSystemType(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "WORK":
			*v = IntegrationMutationRequestSystemType(0)
		case "SOURCECODE":
			*v = IntegrationMutationRequestSystemType(1)
		case "CODEQUALITY":
			*v = IntegrationMutationRequestSystemType(2)
		case "USER":
			*v = IntegrationMutationRequestSystemType(3)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v *IntegrationMutationRequestSystemType) UnmarshalJSON(buf []byte) error {
	var val string
	if err := json.Unmarshal(buf, &val); err != nil {
		return err
	}
	switch val {
	case "WORK":
		*v = 0
	case "SOURCECODE":
		*v = 1
	case "CODEQUALITY":
		*v = 2
	case "USER":
		*v = 3
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v IntegrationMutationRequestSystemType) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("WORK")
	case 1:
		return json.Marshal("SOURCECODE")
	case 2:
		return json.Marshal("CODEQUALITY")
	case 3:
		return json.Marshal("USER")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for SystemType
func (v IntegrationMutationRequestSystemType) String() string {
	switch int32(v) {
	case 0:
		return "WORK"
	case 1:
		return "SOURCECODE"
	case 2:
		return "CODEQUALITY"
	case 3:
		return "USER"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *IntegrationMutationRequestSystemType) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case IntegrationMutationRequestSystemType:
		*v = val
	case int32:
		*v = IntegrationMutationRequestSystemType(int32(val))
	case int:
		*v = IntegrationMutationRequestSystemType(int32(val))
	case string:
		switch val {
		case "WORK":
			*v = IntegrationMutationRequestSystemType(0)
		case "SOURCECODE":
			*v = IntegrationMutationRequestSystemType(1)
		case "CODEQUALITY":
			*v = IntegrationMutationRequestSystemType(2)
		case "USER":
			*v = IntegrationMutationRequestSystemType(3)
		}
	}
	return nil
}

const (
	// IntegrationMutationRequestSystemTypeWork is the enumeration value for work
	IntegrationMutationRequestSystemTypeWork IntegrationMutationRequestSystemType = 0
	// IntegrationMutationRequestSystemTypeSourcecode is the enumeration value for sourcecode
	IntegrationMutationRequestSystemTypeSourcecode IntegrationMutationRequestSystemType = 1
	// IntegrationMutationRequestSystemTypeCodequality is the enumeration value for codequality
	IntegrationMutationRequestSystemTypeCodequality IntegrationMutationRequestSystemType = 2
	// IntegrationMutationRequestSystemTypeUser is the enumeration value for user
	IntegrationMutationRequestSystemTypeUser IntegrationMutationRequestSystemType = 3
)

// IntegrationMutationRequestWebappRequestDate represents the object structure for webapp_request_date
type IntegrationMutationRequestWebappRequestDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIntegrationMutationRequestWebappRequestDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationMutationRequestWebappRequestDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *IntegrationMutationRequestWebappRequestDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIntegrationMutationRequestWebappRequestDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIntegrationMutationRequestWebappRequestDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIntegrationMutationRequestWebappRequestDateObject(o.Rfc3339, false),
	}
}

func (o *IntegrationMutationRequestWebappRequestDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationMutationRequestWebappRequestDate) FromMap(kv map[string]interface{}) {

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

// IntegrationMutationRequest request to change data in integration going from agent service to agent
type IntegrationMutationRequest struct {
	// Action Action to perform
	Action IntegrationMutationRequestAction `json:"action" codec:"action" bson:"action" yaml:"action" faker:"-"`
	// AgentRequestSentDate set by agent operator when sending event to agent
	AgentRequestSentDate IntegrationMutationRequestAgentRequestSentDate `json:"agent_request_sent_date" codec:"agent_request_sent_date" bson:"agent_request_sent_date" yaml:"agent_request_sent_date" faker:"-"`
	// Authorization Authorization information
	Authorization IntegrationMutationRequestAuthorization `json:"authorization" codec:"authorization" bson:"authorization" yaml:"authorization" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Data Mutation parameters that vary based on action as JSON string.
	Data string `json:"data" codec:"data" bson:"data" yaml:"data" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// IntegrationName Name of the integration binary
	IntegrationName string `json:"integration_name" codec:"integration_name" bson:"integration_name" yaml:"integration_name" faker:"-"`
	// JobID The job ID
	JobID string `json:"job_id" codec:"job_id" bson:"job_id" yaml:"job_id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RequestDate the date when the request was made
	RequestDate IntegrationMutationRequestRequestDate `json:"request_date" codec:"request_date" bson:"request_date" yaml:"request_date" faker:"-"`
	// SystemType The system type of the integration
	SystemType IntegrationMutationRequestSystemType `json:"system_type" codec:"system_type" bson:"system_type" yaml:"system_type" faker:"-"`
	// UUID the agent unique identifier
	UUID string `json:"uuid" codec:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
	// WebappRequestDate set by webapp using browser time
	WebappRequestDate IntegrationMutationRequestWebappRequestDate `json:"webapp_request_date" codec:"webapp_request_date" bson:"webapp_request_date" yaml:"webapp_request_date" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*IntegrationMutationRequest)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*IntegrationMutationRequest)(nil)

func toIntegrationMutationRequestObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationMutationRequest:
		return v.ToMap()

	case IntegrationMutationRequestAction:
		return v.String()

	case IntegrationMutationRequestAgentRequestSentDate:
		return v.ToMap()

	case IntegrationMutationRequestAuthorization:
		return v.ToMap()

	case IntegrationMutationRequestRequestDate:
		return v.ToMap()

	case IntegrationMutationRequestSystemType:
		return v.String()

	case IntegrationMutationRequestWebappRequestDate:
		return v.ToMap()
	default:
		return o
	}
}

// String returns a string representation of IntegrationMutationRequest
func (o *IntegrationMutationRequest) String() string {
	return fmt.Sprintf("agent.IntegrationMutationRequest<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *IntegrationMutationRequest) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *IntegrationMutationRequest) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *IntegrationMutationRequest) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *IntegrationMutationRequest) GetModelName() datamodel.ModelNameType {
	return IntegrationMutationRequestModelName
}

// NewIntegrationMutationRequestID provides a template for generating an ID field for IntegrationMutationRequest
func NewIntegrationMutationRequestID(customerID string, refType string, refID string) string {
	return hash.Values("IntegrationMutationRequest", customerID, refType, refID)
}

func (o *IntegrationMutationRequest) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("IntegrationMutationRequest", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *IntegrationMutationRequest) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *IntegrationMutationRequest) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *IntegrationMutationRequest) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *IntegrationMutationRequest) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *IntegrationMutationRequest) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *IntegrationMutationRequest) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *IntegrationMutationRequest) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *IntegrationMutationRequest) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *IntegrationMutationRequest) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = IntegrationMutationRequestModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *IntegrationMutationRequest) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *IntegrationMutationRequest) GetCustomerID() string {
	return o.CustomerID
}

// SetCustomerID will return the customer_id
func (o *IntegrationMutationRequest) SetCustomerID(id string) {
	o.CustomerID = id
}

// GetRefType will return the ref_type
func (o *IntegrationMutationRequest) GetRefType() string {
	return o.RefType
}

// SetRefType will return the ref_type
func (o *IntegrationMutationRequest) SetRefType(t string) {
	o.RefType = t
}

// Clone returns an exact copy of IntegrationMutationRequest
func (o *IntegrationMutationRequest) Clone() datamodel.Model {
	c := new(IntegrationMutationRequest)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *IntegrationMutationRequest) Anon() datamodel.Model {
	c := new(IntegrationMutationRequest)
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
func (o *IntegrationMutationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *IntegrationMutationRequest) UnmarshalJSON(data []byte) error {
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
func (o *IntegrationMutationRequest) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *IntegrationMutationRequest) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two IntegrationMutationRequest objects are equal
func (o *IntegrationMutationRequest) IsEqual(other *IntegrationMutationRequest) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *IntegrationMutationRequest) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{

		"action":                  o.Action.String(),
		"agent_request_sent_date": toIntegrationMutationRequestObject(o.AgentRequestSentDate, false),
		"authorization":           toIntegrationMutationRequestObject(o.Authorization, false),
		"customer_id":             toIntegrationMutationRequestObject(o.CustomerID, false),
		"data":                    toIntegrationMutationRequestObject(o.Data, false),
		"id":                      toIntegrationMutationRequestObject(o.ID, false),
		"integration_instance_id": toIntegrationMutationRequestObject(o.IntegrationInstanceID, true),
		"integration_name":        toIntegrationMutationRequestObject(o.IntegrationName, false),
		"job_id":                  toIntegrationMutationRequestObject(o.JobID, false),
		"ref_id":                  toIntegrationMutationRequestObject(o.RefID, false),
		"ref_type":                toIntegrationMutationRequestObject(o.RefType, false),
		"request_date":            toIntegrationMutationRequestObject(o.RequestDate, false),

		"system_type":         o.SystemType.String(),
		"uuid":                toIntegrationMutationRequestObject(o.UUID, false),
		"webapp_request_date": toIntegrationMutationRequestObject(o.WebappRequestDate, false),
		"hashcode":            toIntegrationMutationRequestObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationMutationRequest) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["action"].(IntegrationMutationRequestAction); ok {
		o.Action = val
	} else {
		if em, ok := kv["action"].(map[string]interface{}); ok {

			ev := em["agent.action"].(string)
			switch ev {
			case "issue_add_comment", "ISSUE_ADD_COMMENT":
				o.Action = 0
			case "issue_set_title", "ISSUE_SET_TITLE":
				o.Action = 1
			case "issue_set_status", "ISSUE_SET_STATUS":
				o.Action = 2
			case "issue_set_priority", "ISSUE_SET_PRIORITY":
				o.Action = 3
			case "issue_set_assignee", "ISSUE_SET_ASSIGNEE":
				o.Action = 4
			case "issue_get_transitions", "ISSUE_GET_TRANSITIONS":
				o.Action = 5
			case "pr_set_title", "PR_SET_TITLE":
				o.Action = 6
			case "pr_set_description", "PR_SET_DESCRIPTION":
				o.Action = 7
			}
		}
		if em, ok := kv["action"].(string); ok {
			switch em {
			case "issue_add_comment", "ISSUE_ADD_COMMENT":
				o.Action = 0
			case "issue_set_title", "ISSUE_SET_TITLE":
				o.Action = 1
			case "issue_set_status", "ISSUE_SET_STATUS":
				o.Action = 2
			case "issue_set_priority", "ISSUE_SET_PRIORITY":
				o.Action = 3
			case "issue_set_assignee", "ISSUE_SET_ASSIGNEE":
				o.Action = 4
			case "issue_get_transitions", "ISSUE_GET_TRANSITIONS":
				o.Action = 5
			case "pr_set_title", "PR_SET_TITLE":
				o.Action = 6
			case "pr_set_description", "PR_SET_DESCRIPTION":
				o.Action = 7
			}
		}
	}

	if val, ok := kv["agent_request_sent_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.AgentRequestSentDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationMutationRequestAgentRequestSentDate); ok {
			// struct
			o.AgentRequestSentDate = sv
		} else if sp, ok := val.(*IntegrationMutationRequestAgentRequestSentDate); ok {
			// struct pointer
			o.AgentRequestSentDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.AgentRequestSentDate.Epoch = dt.Epoch
			o.AgentRequestSentDate.Rfc3339 = dt.Rfc3339
			o.AgentRequestSentDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.AgentRequestSentDate.Epoch = dt.Epoch
			o.AgentRequestSentDate.Rfc3339 = dt.Rfc3339
			o.AgentRequestSentDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.AgentRequestSentDate.Epoch = dt.Epoch
				o.AgentRequestSentDate.Rfc3339 = dt.Rfc3339
				o.AgentRequestSentDate.Offset = dt.Offset
			}
		}
	} else {
		o.AgentRequestSentDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["authorization"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.Authorization.FromMap(kv)
		} else if sv, ok := val.(IntegrationMutationRequestAuthorization); ok {
			// struct
			o.Authorization = sv
		} else if sp, ok := val.(*IntegrationMutationRequestAuthorization); ok {
			// struct pointer
			o.Authorization = *sp
		}
	} else {
		o.Authorization.FromMap(map[string]interface{}{})
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
	if val, ok := kv["integration_name"].(string); ok {
		o.IntegrationName = val
	} else {
		if val, ok := kv["integration_name"]; ok {
			if val == nil {
				o.IntegrationName = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.IntegrationName = fmt.Sprintf("%v", val)
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

	if val, ok := kv["request_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.RequestDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationMutationRequestRequestDate); ok {
			// struct
			o.RequestDate = sv
		} else if sp, ok := val.(*IntegrationMutationRequestRequestDate); ok {
			// struct pointer
			o.RequestDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.RequestDate.Epoch = dt.Epoch
			o.RequestDate.Rfc3339 = dt.Rfc3339
			o.RequestDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.RequestDate.Epoch = dt.Epoch
			o.RequestDate.Rfc3339 = dt.Rfc3339
			o.RequestDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.RequestDate.Epoch = dt.Epoch
				o.RequestDate.Rfc3339 = dt.Rfc3339
				o.RequestDate.Offset = dt.Offset
			}
		}
	} else {
		o.RequestDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["system_type"].(IntegrationMutationRequestSystemType); ok {
		o.SystemType = val
	} else {
		if em, ok := kv["system_type"].(map[string]interface{}); ok {

			ev := em["agent.system_type"].(string)
			switch ev {
			case "work", "WORK":
				o.SystemType = 0
			case "sourcecode", "SOURCECODE":
				o.SystemType = 1
			case "codequality", "CODEQUALITY":
				o.SystemType = 2
			case "user", "USER":
				o.SystemType = 3
			}
		}
		if em, ok := kv["system_type"].(string); ok {
			switch em {
			case "work", "WORK":
				o.SystemType = 0
			case "sourcecode", "SOURCECODE":
				o.SystemType = 1
			case "codequality", "CODEQUALITY":
				o.SystemType = 2
			case "user", "USER":
				o.SystemType = 3
			}
		}
	}
	if val, ok := kv["uuid"].(string); ok {
		o.UUID = val
	} else {
		if val, ok := kv["uuid"]; ok {
			if val == nil {
				o.UUID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.UUID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["webapp_request_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.WebappRequestDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationMutationRequestWebappRequestDate); ok {
			// struct
			o.WebappRequestDate = sv
		} else if sp, ok := val.(*IntegrationMutationRequestWebappRequestDate); ok {
			// struct pointer
			o.WebappRequestDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.WebappRequestDate.Epoch = dt.Epoch
			o.WebappRequestDate.Rfc3339 = dt.Rfc3339
			o.WebappRequestDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.WebappRequestDate.Epoch = dt.Epoch
			o.WebappRequestDate.Rfc3339 = dt.Rfc3339
			o.WebappRequestDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.WebappRequestDate.Epoch = dt.Epoch
				o.WebappRequestDate.Rfc3339 = dt.Rfc3339
				o.WebappRequestDate.Offset = dt.Offset
			}
		}
	} else {
		o.WebappRequestDate.FromMap(map[string]interface{}{})
	}

	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *IntegrationMutationRequest) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Action)
	args = append(args, o.AgentRequestSentDate)
	args = append(args, o.Authorization)
	args = append(args, o.CustomerID)
	args = append(args, o.Data)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.IntegrationName)
	args = append(args, o.JobID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RequestDate)
	args = append(args, o.SystemType)
	args = append(args, o.UUID)
	args = append(args, o.WebappRequestDate)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// SetIntegrationInstanceID will set the integration instance ID
func (o *IntegrationMutationRequest) SetIntegrationInstanceID(id string) {
	if id == "" {
		o.IntegrationInstanceID = nil
	} else {
		o.IntegrationInstanceID = &id
	}
}

// GetIntegrationInstanceID will return the integration instance ID
func (o *IntegrationMutationRequest) GetIntegrationInstanceID() *string {
	return o.IntegrationInstanceID
}

// IntegrationMutationRequestPartial is a partial struct for upsert mutations for IntegrationMutationRequest
type IntegrationMutationRequestPartial struct {
	// Action Action to perform
	Action *IntegrationMutationRequestAction `json:"action,omitempty"`
	// AgentRequestSentDate set by agent operator when sending event to agent
	AgentRequestSentDate *IntegrationMutationRequestAgentRequestSentDate `json:"agent_request_sent_date,omitempty"`
	// Authorization Authorization information
	Authorization *IntegrationMutationRequestAuthorization `json:"authorization,omitempty"`
	// Data Mutation parameters that vary based on action as JSON string.
	Data *string `json:"data,omitempty"`
	// IntegrationName Name of the integration binary
	IntegrationName *string `json:"integration_name,omitempty"`
	// JobID The job ID
	JobID *string `json:"job_id,omitempty"`
	// RequestDate the date when the request was made
	RequestDate *IntegrationMutationRequestRequestDate `json:"request_date,omitempty"`
	// SystemType The system type of the integration
	SystemType *IntegrationMutationRequestSystemType `json:"system_type,omitempty"`
	// UUID the agent unique identifier
	UUID *string `json:"uuid,omitempty"`
	// WebappRequestDate set by webapp using browser time
	WebappRequestDate *IntegrationMutationRequestWebappRequestDate `json:"webapp_request_date,omitempty"`
}

var _ datamodel.PartialModel = (*IntegrationMutationRequestPartial)(nil)

// GetModelName returns the name of the model
func (o *IntegrationMutationRequestPartial) GetModelName() datamodel.ModelNameType {
	return IntegrationMutationRequestModelName
}

// ToMap returns the object as a map
func (o *IntegrationMutationRequestPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{

		"action":                  toIntegrationMutationRequestActionEnum(o.Action),
		"agent_request_sent_date": toIntegrationMutationRequestObject(o.AgentRequestSentDate, true),
		"authorization":           toIntegrationMutationRequestObject(o.Authorization, true),
		"data":                    toIntegrationMutationRequestObject(o.Data, true),
		"integration_name":        toIntegrationMutationRequestObject(o.IntegrationName, true),
		"job_id":                  toIntegrationMutationRequestObject(o.JobID, true),
		"request_date":            toIntegrationMutationRequestObject(o.RequestDate, true),

		"system_type":         toIntegrationMutationRequestSystemTypeEnum(o.SystemType),
		"uuid":                toIntegrationMutationRequestObject(o.UUID, true),
		"webapp_request_date": toIntegrationMutationRequestObject(o.WebappRequestDate, true),
	}
	for k, v := range kv {
		if v == nil || reflect.ValueOf(v).IsZero() {
			delete(kv, k)
		} else {
			if k == "agent_request_sent_date" {
				if dt, ok := v.(*IntegrationMutationRequestAgentRequestSentDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}
			if k == "request_date" {
				if dt, ok := v.(*IntegrationMutationRequestRequestDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}
			if k == "webapp_request_date" {
				if dt, ok := v.(*IntegrationMutationRequestWebappRequestDate); ok {
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
func (o *IntegrationMutationRequestPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *IntegrationMutationRequestPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *IntegrationMutationRequestPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *IntegrationMutationRequestPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *IntegrationMutationRequestPartial) FromMap(kv map[string]interface{}) {
	if val, ok := kv["action"].(*IntegrationMutationRequestAction); ok {
		o.Action = val
	} else if val, ok := kv["action"].(IntegrationMutationRequestAction); ok {
		o.Action = &val
	} else {
		if val, ok := kv["action"]; ok {
			if val == nil {
				o.Action = toIntegrationMutationRequestActionPointer(0)
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["IntegrationMutationRequestAction"]
				}
				// this is an enum pointer
				if em, ok := val.(string); ok {
					switch em {
					case "issue_add_comment", "ISSUE_ADD_COMMENT":
						o.Action = toIntegrationMutationRequestActionPointer(0)
					case "issue_set_title", "ISSUE_SET_TITLE":
						o.Action = toIntegrationMutationRequestActionPointer(1)
					case "issue_set_status", "ISSUE_SET_STATUS":
						o.Action = toIntegrationMutationRequestActionPointer(2)
					case "issue_set_priority", "ISSUE_SET_PRIORITY":
						o.Action = toIntegrationMutationRequestActionPointer(3)
					case "issue_set_assignee", "ISSUE_SET_ASSIGNEE":
						o.Action = toIntegrationMutationRequestActionPointer(4)
					case "issue_get_transitions", "ISSUE_GET_TRANSITIONS":
						o.Action = toIntegrationMutationRequestActionPointer(5)
					case "pr_set_title", "PR_SET_TITLE":
						o.Action = toIntegrationMutationRequestActionPointer(6)
					case "pr_set_description", "PR_SET_DESCRIPTION":
						o.Action = toIntegrationMutationRequestActionPointer(7)
					}
				}
			}
		}
	}

	if o.AgentRequestSentDate == nil {
		o.AgentRequestSentDate = &IntegrationMutationRequestAgentRequestSentDate{}
	}

	if val, ok := kv["agent_request_sent_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.AgentRequestSentDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationMutationRequestAgentRequestSentDate); ok {
			// struct
			o.AgentRequestSentDate = &sv
		} else if sp, ok := val.(*IntegrationMutationRequestAgentRequestSentDate); ok {
			// struct pointer
			o.AgentRequestSentDate = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.AgentRequestSentDate.Epoch = dt.Epoch
			o.AgentRequestSentDate.Rfc3339 = dt.Rfc3339
			o.AgentRequestSentDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.AgentRequestSentDate.Epoch = dt.Epoch
			o.AgentRequestSentDate.Rfc3339 = dt.Rfc3339
			o.AgentRequestSentDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.AgentRequestSentDate.Epoch = dt.Epoch
				o.AgentRequestSentDate.Rfc3339 = dt.Rfc3339
				o.AgentRequestSentDate.Offset = dt.Offset
			}
		}
	} else {
		o.AgentRequestSentDate.FromMap(map[string]interface{}{})
	}

	if o.Authorization == nil {
		o.Authorization = &IntegrationMutationRequestAuthorization{}
	}

	if val, ok := kv["authorization"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.Authorization.FromMap(kv)
		} else if sv, ok := val.(IntegrationMutationRequestAuthorization); ok {
			// struct
			o.Authorization = &sv
		} else if sp, ok := val.(*IntegrationMutationRequestAuthorization); ok {
			// struct pointer
			o.Authorization = sp
		}
	} else {
		o.Authorization.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["data"].(*string); ok {
		o.Data = val
	} else if val, ok := kv["data"].(string); ok {
		o.Data = &val
	} else {
		if val, ok := kv["data"]; ok {
			if val == nil {
				o.Data = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Data = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["integration_name"].(*string); ok {
		o.IntegrationName = val
	} else if val, ok := kv["integration_name"].(string); ok {
		o.IntegrationName = &val
	} else {
		if val, ok := kv["integration_name"]; ok {
			if val == nil {
				o.IntegrationName = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.IntegrationName = pstrings.Pointer(fmt.Sprintf("%v", val))
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

	if o.RequestDate == nil {
		o.RequestDate = &IntegrationMutationRequestRequestDate{}
	}

	if val, ok := kv["request_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.RequestDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationMutationRequestRequestDate); ok {
			// struct
			o.RequestDate = &sv
		} else if sp, ok := val.(*IntegrationMutationRequestRequestDate); ok {
			// struct pointer
			o.RequestDate = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.RequestDate.Epoch = dt.Epoch
			o.RequestDate.Rfc3339 = dt.Rfc3339
			o.RequestDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.RequestDate.Epoch = dt.Epoch
			o.RequestDate.Rfc3339 = dt.Rfc3339
			o.RequestDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.RequestDate.Epoch = dt.Epoch
				o.RequestDate.Rfc3339 = dt.Rfc3339
				o.RequestDate.Offset = dt.Offset
			}
		}
	} else {
		o.RequestDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["system_type"].(*IntegrationMutationRequestSystemType); ok {
		o.SystemType = val
	} else if val, ok := kv["system_type"].(IntegrationMutationRequestSystemType); ok {
		o.SystemType = &val
	} else {
		if val, ok := kv["system_type"]; ok {
			if val == nil {
				o.SystemType = toIntegrationMutationRequestSystemTypePointer(0)
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["IntegrationMutationRequestSystemType"]
				}
				// this is an enum pointer
				if em, ok := val.(string); ok {
					switch em {
					case "work", "WORK":
						o.SystemType = toIntegrationMutationRequestSystemTypePointer(0)
					case "sourcecode", "SOURCECODE":
						o.SystemType = toIntegrationMutationRequestSystemTypePointer(1)
					case "codequality", "CODEQUALITY":
						o.SystemType = toIntegrationMutationRequestSystemTypePointer(2)
					case "user", "USER":
						o.SystemType = toIntegrationMutationRequestSystemTypePointer(3)
					}
				}
			}
		}
	}
	if val, ok := kv["uuid"].(*string); ok {
		o.UUID = val
	} else if val, ok := kv["uuid"].(string); ok {
		o.UUID = &val
	} else {
		if val, ok := kv["uuid"]; ok {
			if val == nil {
				o.UUID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.UUID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if o.WebappRequestDate == nil {
		o.WebappRequestDate = &IntegrationMutationRequestWebappRequestDate{}
	}

	if val, ok := kv["webapp_request_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.WebappRequestDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationMutationRequestWebappRequestDate); ok {
			// struct
			o.WebappRequestDate = &sv
		} else if sp, ok := val.(*IntegrationMutationRequestWebappRequestDate); ok {
			// struct pointer
			o.WebappRequestDate = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.WebappRequestDate.Epoch = dt.Epoch
			o.WebappRequestDate.Rfc3339 = dt.Rfc3339
			o.WebappRequestDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.WebappRequestDate.Epoch = dt.Epoch
			o.WebappRequestDate.Rfc3339 = dt.Rfc3339
			o.WebappRequestDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.WebappRequestDate.Epoch = dt.Epoch
				o.WebappRequestDate.Rfc3339 = dt.Rfc3339
				o.WebappRequestDate.Offset = dt.Offset
			}
		}
	} else {
		o.WebappRequestDate.FromMap(map[string]interface{}{})
	}

	o.setDefaults(false)
}
