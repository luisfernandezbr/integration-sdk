// DO NOT EDIT -- generated code

// Package agent - the agent communicates with pinpoint cloud to send integration data
package agent

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
	"github.com/pinpt/go-common/slice"
	pstrings "github.com/pinpt/go-common/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// CodequalityRequestModelName is the model name
	CodequalityRequestModelName datamodel.ModelNameType = "agent.CodequalityRequest"
)

// CodequalityRequestIntegrationAuthorization represents the object structure for authorization
type CodequalityRequestIntegrationAuthorization struct {
	// AccessToken Access token
	AccessToken *string `json:"access_token,omitempty" codec:"access_token,omitempty" bson:"access_token" yaml:"access_token,omitempty" faker:"-"`
	// RefreshToken Refresh token
	RefreshToken *string `json:"refresh_token,omitempty" codec:"refresh_token,omitempty" bson:"refresh_token" yaml:"refresh_token,omitempty" faker:"-"`
	// URL URL of instance if relevant
	URL *string `json:"url,omitempty" codec:"url,omitempty" bson:"url" yaml:"url,omitempty" faker:"-"`
	// Username Username for instance, if relevant
	Username *string `json:"username,omitempty" codec:"username,omitempty" bson:"username" yaml:"username,omitempty" faker:"-"`
	// Password Password for instance, if relevant
	Password *string `json:"password,omitempty" codec:"password,omitempty" bson:"password" yaml:"password,omitempty" faker:"-"`
	// APIToken API Token for instance, if relevant
	APIToken *string `json:"api_token,omitempty" codec:"api_token,omitempty" bson:"api_token" yaml:"api_token,omitempty" faker:"-"`
	// Collection Collection for instance, if relevant
	Collection *string `json:"collection,omitempty" codec:"collection,omitempty" bson:"collection" yaml:"collection,omitempty" faker:"-"`
	// APIKey API Key for instance, if relevant
	APIKey *string `json:"api_key,omitempty" codec:"api_key,omitempty" bson:"api_key" yaml:"api_key,omitempty" faker:"-"`
	// Authorization the agents encrypted authorization
	Authorization *string `json:"authorization,omitempty" codec:"authorization,omitempty" bson:"authorization" yaml:"authorization,omitempty" faker:"-"`
	// Hostname Hostname for instance, if relevant
	Hostname *string `json:"hostname,omitempty" codec:"hostname,omitempty" bson:"hostname" yaml:"hostname,omitempty" faker:"-"`
}

func toCodequalityRequestIntegrationAuthorizationObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CodequalityRequestIntegrationAuthorization:
		return v.ToMap()

	default:
		return o
	}
}

func (o *CodequalityRequestIntegrationAuthorization) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// AccessToken Access token
		"access_token": toCodequalityRequestIntegrationAuthorizationObject(o.AccessToken, true),
		// RefreshToken Refresh token
		"refresh_token": toCodequalityRequestIntegrationAuthorizationObject(o.RefreshToken, true),
		// URL URL of instance if relevant
		"url": toCodequalityRequestIntegrationAuthorizationObject(o.URL, true),
		// Username Username for instance, if relevant
		"username": toCodequalityRequestIntegrationAuthorizationObject(o.Username, true),
		// Password Password for instance, if relevant
		"password": toCodequalityRequestIntegrationAuthorizationObject(o.Password, true),
		// APIToken API Token for instance, if relevant
		"api_token": toCodequalityRequestIntegrationAuthorizationObject(o.APIToken, true),
		// Collection Collection for instance, if relevant
		"collection": toCodequalityRequestIntegrationAuthorizationObject(o.Collection, true),
		// APIKey API Key for instance, if relevant
		"api_key": toCodequalityRequestIntegrationAuthorizationObject(o.APIKey, true),
		// Authorization the agents encrypted authorization
		"authorization": toCodequalityRequestIntegrationAuthorizationObject(o.Authorization, true),
		// Hostname Hostname for instance, if relevant
		"hostname": toCodequalityRequestIntegrationAuthorizationObject(o.Hostname, true),
	}
}

func (o *CodequalityRequestIntegrationAuthorization) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *CodequalityRequestIntegrationAuthorization) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["access_token"].(*string); ok {
		o.AccessToken = val
	} else if val, ok := kv["access_token"].(string); ok {
		o.AccessToken = &val
	} else {
		if val, ok := kv["access_token"]; ok {
			if val == nil {
				o.AccessToken = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.AccessToken = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if val, ok := kv["refresh_token"].(*string); ok {
		o.RefreshToken = val
	} else if val, ok := kv["refresh_token"].(string); ok {
		o.RefreshToken = &val
	} else {
		if val, ok := kv["refresh_token"]; ok {
			if val == nil {
				o.RefreshToken = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.RefreshToken = pstrings.Pointer(fmt.Sprintf("%v", val))
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

	if val, ok := kv["username"].(*string); ok {
		o.Username = val
	} else if val, ok := kv["username"].(string); ok {
		o.Username = &val
	} else {
		if val, ok := kv["username"]; ok {
			if val == nil {
				o.Username = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Username = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if val, ok := kv["password"].(*string); ok {
		o.Password = val
	} else if val, ok := kv["password"].(string); ok {
		o.Password = &val
	} else {
		if val, ok := kv["password"]; ok {
			if val == nil {
				o.Password = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Password = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if val, ok := kv["api_token"].(*string); ok {
		o.APIToken = val
	} else if val, ok := kv["api_token"].(string); ok {
		o.APIToken = &val
	} else {
		if val, ok := kv["api_token"]; ok {
			if val == nil {
				o.APIToken = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.APIToken = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if val, ok := kv["collection"].(*string); ok {
		o.Collection = val
	} else if val, ok := kv["collection"].(string); ok {
		o.Collection = &val
	} else {
		if val, ok := kv["collection"]; ok {
			if val == nil {
				o.Collection = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Collection = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if val, ok := kv["api_key"].(*string); ok {
		o.APIKey = val
	} else if val, ok := kv["api_key"].(string); ok {
		o.APIKey = &val
	} else {
		if val, ok := kv["api_key"]; ok {
			if val == nil {
				o.APIKey = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.APIKey = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if val, ok := kv["authorization"].(*string); ok {
		o.Authorization = val
	} else if val, ok := kv["authorization"].(string); ok {
		o.Authorization = &val
	} else {
		if val, ok := kv["authorization"]; ok {
			if val == nil {
				o.Authorization = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Authorization = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if val, ok := kv["hostname"].(*string); ok {
		o.Hostname = val
	} else if val, ok := kv["hostname"].(string); ok {
		o.Hostname = &val
	} else {
		if val, ok := kv["hostname"]; ok {
			if val == nil {
				o.Hostname = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Hostname = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	o.setDefaults(false)
}

// CodequalityRequestIntegrationLocation is the enumeration type for location
type CodequalityRequestIntegrationLocation int32

// String returns the string value for IntegrationLocation
func (v CodequalityRequestIntegrationLocation) String() string {
	switch int32(v) {
	case 0:
		return "PRIVATE"
	case 1:
		return "CLOUD"
	}
	return "unset"
}

const (
	// IntegrationLocationPrivate is the enumeration value for private
	CodequalityRequestIntegrationLocationPrivate CodequalityRequestIntegrationLocation = 0
	// IntegrationLocationCloud is the enumeration value for cloud
	CodequalityRequestIntegrationLocationCloud CodequalityRequestIntegrationLocation = 1
)

// CodequalityRequestIntegrationProgress represents the object structure for progress
type CodequalityRequestIntegrationProgress struct {
	// Message Any relevant messaging during processing
	Message string `json:"message" codec:"message" bson:"message" yaml:"message" faker:"-"`
	// Total The total amount to be processed
	Total int64 `json:"total" codec:"total" bson:"total" yaml:"total" faker:"-"`
	// Completed The total amount processed thus far
	Completed int64 `json:"completed" codec:"completed" bson:"completed" yaml:"completed" faker:"-"`
}

func toCodequalityRequestIntegrationProgressObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CodequalityRequestIntegrationProgress:
		return v.ToMap()

	default:
		return o
	}
}

func (o *CodequalityRequestIntegrationProgress) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Message Any relevant messaging during processing
		"message": toCodequalityRequestIntegrationProgressObject(o.Message, false),
		// Total The total amount to be processed
		"total": toCodequalityRequestIntegrationProgressObject(o.Total, false),
		// Completed The total amount processed thus far
		"completed": toCodequalityRequestIntegrationProgressObject(o.Completed, false),
	}
}

func (o *CodequalityRequestIntegrationProgress) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *CodequalityRequestIntegrationProgress) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["message"].(string); ok {
		o.Message = val
	} else {
		if val, ok := kv["message"]; ok {
			if val == nil {
				o.Message = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Message = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["total"].(int64); ok {
		o.Total = val
	} else {
		if val, ok := kv["total"]; ok {
			if val == nil {
				o.Total = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Total = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["completed"].(int64); ok {
		o.Completed = val
	} else {
		if val, ok := kv["completed"]; ok {
			if val == nil {
				o.Completed = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Completed = number.ToInt64Any(val)
			}
		}
	}
	o.setDefaults(false)
}

// CodequalityRequestIntegrationSystemType is the enumeration type for system_type
type CodequalityRequestIntegrationSystemType int32

// String returns the string value for IntegrationSystemType
func (v CodequalityRequestIntegrationSystemType) String() string {
	switch int32(v) {
	case 0:
		return "WORK"
	case 1:
		return "SOURCECODE"
	case 2:
		return "CODEQUALITY"
	}
	return "unset"
}

const (
	// IntegrationSystemTypeWork is the enumeration value for work
	CodequalityRequestIntegrationSystemTypeWork CodequalityRequestIntegrationSystemType = 0
	// IntegrationSystemTypeSourcecode is the enumeration value for sourcecode
	CodequalityRequestIntegrationSystemTypeSourcecode CodequalityRequestIntegrationSystemType = 1
	// IntegrationSystemTypeCodequality is the enumeration value for codequality
	CodequalityRequestIntegrationSystemTypeCodequality CodequalityRequestIntegrationSystemType = 2
)

// CodequalityRequestIntegrationValidatedDate represents the object structure for validated_date
type CodequalityRequestIntegrationValidatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toCodequalityRequestIntegrationValidatedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CodequalityRequestIntegrationValidatedDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *CodequalityRequestIntegrationValidatedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toCodequalityRequestIntegrationValidatedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toCodequalityRequestIntegrationValidatedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toCodequalityRequestIntegrationValidatedDateObject(o.Rfc3339, false),
	}
}

func (o *CodequalityRequestIntegrationValidatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *CodequalityRequestIntegrationValidatedDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

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

// CodequalityRequestIntegration represents the object structure for integration
type CodequalityRequestIntegration struct {
	// Active If true, the integration is still active
	Active bool `json:"active" codec:"active" bson:"active" yaml:"active" faker:"-"`
	// Authorization Authorization information
	Authorization CodequalityRequestIntegrationAuthorization `json:"authorization" codec:"authorization" bson:"authorization" yaml:"authorization" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Errored If authorization failed by the agent
	Errored *bool `json:"errored,omitempty" codec:"errored,omitempty" bson:"errored" yaml:"errored,omitempty" faker:"-"`
	// Exclusions The exclusion list for this integration
	Exclusions []string `json:"exclusions" codec:"exclusions" bson:"exclusions" yaml:"exclusions" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// Location The location of this integration (on-premise / private or cloud)
	Location CodequalityRequestIntegrationLocation `json:"location" codec:"location" bson:"location" yaml:"location" faker:"-"`
	// Name The user friendly name of the integration
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
	// Progress Agent processing progress
	Progress CodequalityRequestIntegrationProgress `json:"progress" codec:"progress" bson:"progress" yaml:"progress" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// SystemType The system type of the integration (sourcecode / work (jira) / codequality / etc.)
	SystemType CodequalityRequestIntegrationSystemType `json:"system_type" codec:"system_type" bson:"system_type" yaml:"system_type" faker:"-"`
	// Validated If the validation has been run against this instance
	Validated *bool `json:"validated,omitempty" codec:"validated,omitempty" bson:"validated" yaml:"validated,omitempty" faker:"-"`
	// ValidatedDate Date when validated
	ValidatedDate CodequalityRequestIntegrationValidatedDate `json:"validated_date" codec:"validated_date" bson:"validated_date" yaml:"validated_date" faker:"-"`
	// ValidationMessage The validation message from the agent
	ValidationMessage *string `json:"validation_message,omitempty" codec:"validation_message,omitempty" bson:"validation_message" yaml:"validation_message,omitempty" faker:"-"`
}

func toCodequalityRequestIntegrationObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CodequalityRequestIntegration:
		return v.ToMap()

	case CodequalityRequestIntegrationAuthorization:
		return v.ToMap()

	case CodequalityRequestIntegrationLocation:
		return v.String()

	case CodequalityRequestIntegrationProgress:
		return v.ToMap()

	case CodequalityRequestIntegrationSystemType:
		return v.String()

	case CodequalityRequestIntegrationValidatedDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *CodequalityRequestIntegration) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Active If true, the integration is still active
		"active": toCodequalityRequestIntegrationObject(o.Active, false),
		// Authorization Authorization information
		"authorization": toCodequalityRequestIntegrationObject(o.Authorization, false),
		// CustomerID the customer id for the model instance
		"customer_id": toCodequalityRequestIntegrationObject(o.CustomerID, false),
		// Errored If authorization failed by the agent
		"errored": toCodequalityRequestIntegrationObject(o.Errored, true),
		// Exclusions The exclusion list for this integration
		"exclusions": toCodequalityRequestIntegrationObject(o.Exclusions, false),
		// ID the primary key for the model instance
		"id": toCodequalityRequestIntegrationObject(o.ID, false),
		// Location The location of this integration (on-premise / private or cloud)
		"location": toCodequalityRequestIntegrationObject(o.Location, false),
		// Name The user friendly name of the integration
		"name": toCodequalityRequestIntegrationObject(o.Name, false),
		// Progress Agent processing progress
		"progress": toCodequalityRequestIntegrationObject(o.Progress, false),
		// RefID the source system id for the model instance
		"ref_id": toCodequalityRequestIntegrationObject(o.RefID, false),
		// RefType the source system identifier for the model instance
		"ref_type": toCodequalityRequestIntegrationObject(o.RefType, false),
		// SystemType The system type of the integration (sourcecode / work (jira) / codequality / etc.)
		"system_type": toCodequalityRequestIntegrationObject(o.SystemType, false),
		// Validated If the validation has been run against this instance
		"validated": toCodequalityRequestIntegrationObject(o.Validated, true),
		// ValidatedDate Date when validated
		"validated_date": toCodequalityRequestIntegrationObject(o.ValidatedDate, false),
		// ValidationMessage The validation message from the agent
		"validation_message": toCodequalityRequestIntegrationObject(o.ValidationMessage, true),
	}
}

func (o *CodequalityRequestIntegration) setDefaults(frommap bool) {

	if o.Errored == nil {
		var v bool
		o.Errored = &v
	}

	if o.Validated == nil {
		var v bool
		o.Validated = &v
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *CodequalityRequestIntegration) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["active"].(bool); ok {
		o.Active = val
	} else {
		if val, ok := kv["active"]; ok {
			if val == nil {
				o.Active = number.ToBoolAny(nil)
			} else {
				o.Active = number.ToBoolAny(val)
			}
		}
	}

	if val, ok := kv["authorization"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.Authorization.FromMap(kv)
		} else if sv, ok := val.(CodequalityRequestIntegrationAuthorization); ok {
			// struct
			o.Authorization = sv
		} else if sp, ok := val.(*CodequalityRequestIntegrationAuthorization); ok {
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.CustomerID = fmt.Sprintf("%v", val)
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
				o.Errored = number.BoolPointer(number.ToBoolAny(nil))
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Errored = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}

	if val, ok := kv["exclusions"]; ok {
		if val != nil {
			na := make([]string, 0)
			if a, ok := val.([]string); ok {
				na = append(na, a...)
			} else {
				if a, ok := val.([]interface{}); ok {
					for _, ae := range a {
						if av, ok := ae.(string); ok {
							na = append(na, av)
						} else {
							if badMap, ok := ae.(map[interface{}]interface{}); ok {
								ae = slice.ConvertToStringToInterface(badMap)
							}
							b, _ := json.Marshal(ae)
							var av string
							if err := json.Unmarshal(b, &av); err != nil {
								panic("unsupported type for exclusions field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else if s, ok := val.(string); ok {
					for _, sv := range strings.Split(s, ",") {
						na = append(na, strings.TrimSpace(sv))
					}
				} else if a, ok := val.(primitive.A); ok {
					for _, ae := range a {
						if av, ok := ae.(string); ok {
							na = append(na, av)
						} else {
							b, _ := json.Marshal(ae)
							var av string
							if err := json.Unmarshal(b, &av); err != nil {
								panic("unsupported type for exclusions field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for exclusions field")
				}
			}
			o.Exclusions = na
		}
	}
	if o.Exclusions == nil {
		o.Exclusions = make([]string, 0)
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

	if val, ok := kv["location"].(CodequalityRequestIntegrationLocation); ok {
		o.Location = val
	} else {
		if em, ok := kv["location"].(map[string]interface{}); ok {
			ev := em["agent.location"].(string)
			switch ev {
			case "private", "PRIVATE":
				o.Location = 0
			case "cloud", "CLOUD":
				o.Location = 1
			}
		}
		if em, ok := kv["location"].(string); ok {
			switch em {
			case "private", "PRIVATE":
				o.Location = 0
			case "cloud", "CLOUD":
				o.Location = 1
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Name = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["progress"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.Progress.FromMap(kv)
		} else if sv, ok := val.(CodequalityRequestIntegrationProgress); ok {
			// struct
			o.Progress = sv
		} else if sp, ok := val.(*CodequalityRequestIntegrationProgress); ok {
			// struct pointer
			o.Progress = *sp
		}
	} else {
		o.Progress.FromMap(map[string]interface{}{})
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

	if val, ok := kv["system_type"].(CodequalityRequestIntegrationSystemType); ok {
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
			}
		}
	}

	if val, ok := kv["validated"].(*bool); ok {
		o.Validated = val
	} else if val, ok := kv["validated"].(bool); ok {
		o.Validated = &val
	} else {
		if val, ok := kv["validated"]; ok {
			if val == nil {
				o.Validated = number.BoolPointer(number.ToBoolAny(nil))
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Validated = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}

	if val, ok := kv["validated_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.ValidatedDate.FromMap(kv)
		} else if sv, ok := val.(CodequalityRequestIntegrationValidatedDate); ok {
			// struct
			o.ValidatedDate = sv
		} else if sp, ok := val.(*CodequalityRequestIntegrationValidatedDate); ok {
			// struct pointer
			o.ValidatedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.ValidatedDate.Epoch = dt.Epoch
			o.ValidatedDate.Rfc3339 = dt.Rfc3339
			o.ValidatedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.ValidatedDate.Epoch = dt.Epoch
			o.ValidatedDate.Rfc3339 = dt.Rfc3339
			o.ValidatedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.ValidatedDate.Epoch = dt.Epoch
				o.ValidatedDate.Rfc3339 = dt.Rfc3339
				o.ValidatedDate.Offset = dt.Offset
			}
		}
	} else {
		o.ValidatedDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["validation_message"].(*string); ok {
		o.ValidationMessage = val
	} else if val, ok := kv["validation_message"].(string); ok {
		o.ValidationMessage = &val
	} else {
		if val, ok := kv["validation_message"]; ok {
			if val == nil {
				o.ValidationMessage = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.ValidationMessage = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	o.setDefaults(false)
}

// CodequalityRequestRequestDate represents the object structure for request_date
type CodequalityRequestRequestDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toCodequalityRequestRequestDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CodequalityRequestRequestDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *CodequalityRequestRequestDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toCodequalityRequestRequestDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toCodequalityRequestRequestDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toCodequalityRequestRequestDateObject(o.Rfc3339, false),
	}
}

func (o *CodequalityRequestRequestDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *CodequalityRequestRequestDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

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

// CodequalityRequest an agent action to request adding new codequality entities
type CodequalityRequest struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// Integration the integration details to use
	Integration CodequalityRequestIntegration `json:"integration" codec:"integration" bson:"integration" yaml:"integration" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RequestDate the date when the request was made
	RequestDate CodequalityRequestRequestDate `json:"request_date" codec:"request_date" bson:"request_date" yaml:"request_date" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// UUID the agent unique identifier
	UUID string `json:"uuid" codec:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*CodequalityRequest)(nil)

func toCodequalityRequestObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CodequalityRequest:
		return v.ToMap()

	case CodequalityRequestIntegration:
		return v.ToMap()

	case CodequalityRequestRequestDate:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of CodequalityRequest
func (o *CodequalityRequest) String() string {
	return fmt.Sprintf("agent.CodequalityRequest<%s>", o.ID)
}

// GetModelName returns the name of the model
func (o *CodequalityRequest) GetModelName() datamodel.ModelNameType {
	return CodequalityRequestModelName
}

// NewCodequalityRequestID provides a template for generating an ID field for CodequalityRequest
func NewCodequalityRequestID(customerID string, refType string, refID string) string {
	return hash.Values("CodequalityRequest", customerID, refType, refID)
}

func (o *CodequalityRequest) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("CodequalityRequest", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *CodequalityRequest) GetID() string {
	return o.ID
}

// GetRefID returns the RefID for the object
func (o *CodequalityRequest) GetRefID() string {
	return o.RefID
}

// GetCustomerID will return the customer_id
func (o *CodequalityRequest) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of CodequalityRequest
func (o *CodequalityRequest) Clone() datamodel.Model {
	c := new(CodequalityRequest)
	c.FromMap(o.ToMap())
	return c
}

// MarshalJSON returns the bytes for marshaling to json
func (o *CodequalityRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *CodequalityRequest) UnmarshalJSON(data []byte) error {
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
func (o *CodequalityRequest) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two CodequalityRequest objects are equal
func (o *CodequalityRequest) IsEqual(other *CodequalityRequest) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *CodequalityRequest) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"customer_id":  toCodequalityRequestObject(o.CustomerID, false),
		"id":           toCodequalityRequestObject(o.ID, false),
		"integration":  toCodequalityRequestObject(o.Integration, false),
		"ref_id":       toCodequalityRequestObject(o.RefID, false),
		"ref_type":     toCodequalityRequestObject(o.RefType, false),
		"request_date": toCodequalityRequestObject(o.RequestDate, false),
		"updated_ts":   toCodequalityRequestObject(o.UpdatedAt, false),
		"uuid":         toCodequalityRequestObject(o.UUID, false),
		"hashcode":     toCodequalityRequestObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *CodequalityRequest) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
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

	if val, ok := kv["integration"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.Integration.FromMap(kv)
		} else if sv, ok := val.(CodequalityRequestIntegration); ok {
			// struct
			o.Integration = sv
		} else if sp, ok := val.(*CodequalityRequestIntegration); ok {
			// struct pointer
			o.Integration = *sp
		}
	} else {
		o.Integration.FromMap(map[string]interface{}{})
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

	if val, ok := kv["request_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.RequestDate.FromMap(kv)
		} else if sv, ok := val.(CodequalityRequestRequestDate); ok {
			// struct
			o.RequestDate = sv
		} else if sp, ok := val.(*CodequalityRequestRequestDate); ok {
			// struct pointer
			o.RequestDate = *sp
		}
	} else {
		o.RequestDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["updated_ts"].(int64); ok {
		o.UpdatedAt = val
	} else {
		if val, ok := kv["updated_ts"]; ok {
			if val == nil {
				o.UpdatedAt = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.UpdatedAt = number.ToInt64Any(val)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.UUID = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *CodequalityRequest) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.Integration)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RequestDate)
	args = append(args, o.UpdatedAt)
	args = append(args, o.UUID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}
