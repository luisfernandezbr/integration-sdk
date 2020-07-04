// DO NOT EDIT -- generated code

// Package sourcecode - the system which contains source code
package sourcecode

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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (

	// BlameTable is the default table name
	BlameTable datamodel.ModelNameType = "sourcecode_blame"

	// BlameModelName is the model name
	BlameModelName datamodel.ModelNameType = "sourcecode.Blame"
)

const (
	// BlameModelBlanksColumn is the column json value blanks
	BlameModelBlanksColumn = "blanks"
	// BlameModelChangeDateColumn is the column json value change_date
	BlameModelChangeDateColumn = "change_date"
	// BlameModelChangeDateEpochColumn is the column json value epoch
	BlameModelChangeDateEpochColumn = "epoch"
	// BlameModelChangeDateOffsetColumn is the column json value offset
	BlameModelChangeDateOffsetColumn = "offset"
	// BlameModelChangeDateRfc3339Column is the column json value rfc3339
	BlameModelChangeDateRfc3339Column = "rfc3339"
	// BlameModelCommentsColumn is the column json value comments
	BlameModelCommentsColumn = "comments"
	// BlameModelCommitIDColumn is the column json value commit_id
	BlameModelCommitIDColumn = "commit_id"
	// BlameModelComplexityColumn is the column json value complexity
	BlameModelComplexityColumn = "complexity"
	// BlameModelCustomerIDColumn is the column json value customer_id
	BlameModelCustomerIDColumn = "customer_id"
	// BlameModelExcludedColumn is the column json value excluded
	BlameModelExcludedColumn = "excluded"
	// BlameModelExcludedReasonColumn is the column json value excluded_reason
	BlameModelExcludedReasonColumn = "excluded_reason"
	// BlameModelFilenameColumn is the column json value filename
	BlameModelFilenameColumn = "filename"
	// BlameModelIDColumn is the column json value id
	BlameModelIDColumn = "id"
	// BlameModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	BlameModelIntegrationInstanceIDColumn = "integration_instance_id"
	// BlameModelLanguageColumn is the column json value language
	BlameModelLanguageColumn = "language"
	// BlameModelLicenseColumn is the column json value license
	BlameModelLicenseColumn = "license"
	// BlameModelLinesColumn is the column json value lines
	BlameModelLinesColumn = "lines"
	// BlameModelLinesAuthorRefIDColumn is the column json value author_ref_id
	BlameModelLinesAuthorRefIDColumn = "author_ref_id"
	// BlameModelLinesBlankColumn is the column json value blank
	BlameModelLinesBlankColumn = "blank"
	// BlameModelLinesCodeColumn is the column json value code
	BlameModelLinesCodeColumn = "code"
	// BlameModelLinesCommentColumn is the column json value comment
	BlameModelLinesCommentColumn = "comment"
	// BlameModelLinesDateColumn is the column json value date
	BlameModelLinesDateColumn = "date"
	// BlameModelLinesShaColumn is the column json value sha
	BlameModelLinesShaColumn = "sha"
	// BlameModelLocColumn is the column json value loc
	BlameModelLocColumn = "loc"
	// BlameModelRefIDColumn is the column json value ref_id
	BlameModelRefIDColumn = "ref_id"
	// BlameModelRefTypeColumn is the column json value ref_type
	BlameModelRefTypeColumn = "ref_type"
	// BlameModelRepoIDColumn is the column json value repo_id
	BlameModelRepoIDColumn = "repo_id"
	// BlameModelShaColumn is the column json value sha
	BlameModelShaColumn = "sha"
	// BlameModelSizeColumn is the column json value size
	BlameModelSizeColumn = "size"
	// BlameModelSlocColumn is the column json value sloc
	BlameModelSlocColumn = "sloc"
	// BlameModelStatusColumn is the column json value status
	BlameModelStatusColumn = "status"
)

// BlameChangeDate represents the object structure for change_date
type BlameChangeDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toBlameChangeDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *BlameChangeDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *BlameChangeDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toBlameChangeDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toBlameChangeDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toBlameChangeDateObject(o.Rfc3339, false),
	}
}

func (o *BlameChangeDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *BlameChangeDate) FromMap(kv map[string]interface{}) {

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

// BlameLines represents the object structure for lines
type BlameLines struct {
	// AuthorRefID the author ref_id of this line when last changed
	AuthorRefID string `json:"author_ref_id" codec:"author_ref_id" bson:"author_ref_id" yaml:"author_ref_id" faker:"-"`
	// Blank if the line is a blank line
	Blank bool `json:"blank" codec:"blank" bson:"blank" yaml:"blank" faker:"-"`
	// Code if the line is sourcecode
	Code bool `json:"code" codec:"code" bson:"code" yaml:"code" faker:"-"`
	// Comment if the line is a comment
	Comment bool `json:"comment" codec:"comment" bson:"comment" yaml:"comment" faker:"-"`
	// Date the change date in RFC3339 format of this line when last changed
	Date string `json:"date" codec:"date" bson:"date" yaml:"date" faker:"-"`
	// Sha the sha when this line was last changed
	Sha string `json:"sha" codec:"sha" bson:"sha" yaml:"sha" faker:"-"`
}

func toBlameLinesObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *BlameLines:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *BlameLines) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// AuthorRefID the author ref_id of this line when last changed
		"author_ref_id": toBlameLinesObject(o.AuthorRefID, false),
		// Blank if the line is a blank line
		"blank": toBlameLinesObject(o.Blank, false),
		// Code if the line is sourcecode
		"code": toBlameLinesObject(o.Code, false),
		// Comment if the line is a comment
		"comment": toBlameLinesObject(o.Comment, false),
		// Date the change date in RFC3339 format of this line when last changed
		"date": toBlameLinesObject(o.Date, false),
		// Sha the sha when this line was last changed
		"sha": toBlameLinesObject(o.Sha, false),
	}
}

func (o *BlameLines) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *BlameLines) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["author_ref_id"].(string); ok {
		o.AuthorRefID = val
	} else {
		if val, ok := kv["author_ref_id"]; ok {
			if val == nil {
				o.AuthorRefID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.AuthorRefID = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["blank"].(bool); ok {
		o.Blank = val
	} else {
		if val, ok := kv["blank"]; ok {
			if val == nil {
				o.Blank = false
			} else {
				o.Blank = number.ToBoolAny(val)
			}
		}
	}
	if val, ok := kv["code"].(bool); ok {
		o.Code = val
	} else {
		if val, ok := kv["code"]; ok {
			if val == nil {
				o.Code = false
			} else {
				o.Code = number.ToBoolAny(val)
			}
		}
	}
	if val, ok := kv["comment"].(bool); ok {
		o.Comment = val
	} else {
		if val, ok := kv["comment"]; ok {
			if val == nil {
				o.Comment = false
			} else {
				o.Comment = number.ToBoolAny(val)
			}
		}
	}
	if val, ok := kv["date"].(string); ok {
		o.Date = val
	} else {
		if val, ok := kv["date"]; ok {
			if val == nil {
				o.Date = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Date = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["sha"].(string); ok {
		o.Sha = val
	} else {
		if val, ok := kv["sha"]; ok {
			if val == nil {
				o.Sha = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Sha = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// BlameStatus is the enumeration type for status
type BlameStatus int32

// toBlameStatusPointer is the enumeration pointer type for status
func toBlameStatusPointer(v int32) *BlameStatus {
	nv := BlameStatus(v)
	return &nv
}

// toBlameStatusEnum is the enumeration pointer wrapper for status
func toBlameStatusEnum(v *BlameStatus) string {
	if v == nil {
		return toBlameStatusPointer(0).String()
	}
	return v.String()
}

// UnmarshalBSONValue for unmarshaling value
func (v *BlameStatus) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = BlameStatus(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "ADDED":
			*v = BlameStatus(0)
		case "MODIFIED":
			*v = BlameStatus(1)
		case "REMOVED":
			*v = BlameStatus(2)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v BlameStatus) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "ADDED":
		v = 0
	case "MODIFIED":
		v = 1
	case "REMOVED":
		v = 2
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v BlameStatus) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("ADDED")
	case 1:
		return json.Marshal("MODIFIED")
	case 2:
		return json.Marshal("REMOVED")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for Status
func (v BlameStatus) String() string {
	switch int32(v) {
	case 0:
		return "ADDED"
	case 1:
		return "MODIFIED"
	case 2:
		return "REMOVED"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *BlameStatus) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case BlameStatus:
		*v = val
	case int32:
		*v = BlameStatus(int32(val))
	case int:
		*v = BlameStatus(int32(val))
	case string:
		switch val {
		case "ADDED":
			*v = BlameStatus(0)
		case "MODIFIED":
			*v = BlameStatus(1)
		case "REMOVED":
			*v = BlameStatus(2)
		}
	}
	return nil
}

const (
	// BlameStatusAdded is the enumeration value for added
	BlameStatusAdded BlameStatus = 0
	// BlameStatusModified is the enumeration value for modified
	BlameStatusModified BlameStatus = 1
	// BlameStatusRemoved is the enumeration value for removed
	BlameStatusRemoved BlameStatus = 2
)

// Blame the blame detail for each commit in a repo
type Blame struct {
	// Blanks the count of blank lines in the file
	Blanks int64 `json:"blanks" codec:"blanks" bson:"blanks" yaml:"blanks" faker:"-"`
	// ChangeDate the date of the change
	ChangeDate BlameChangeDate `json:"change_date" codec:"change_date" bson:"change_date" yaml:"change_date" faker:"-"`
	// Comments the count of comment lines in the file based on language rules
	Comments int64 `json:"comments" codec:"comments" bson:"comments" yaml:"comments" faker:"-"`
	// CommitID the commit ID
	CommitID string `json:"commit_id" codec:"commit_id" bson:"commit_id" yaml:"commit_id" faker:"-"`
	// Complexity the cyclomatic complexity for the change
	Complexity int64 `json:"complexity" codec:"complexity" bson:"complexity" yaml:"complexity" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Excluded if the result was excluded
	Excluded bool `json:"excluded" codec:"excluded" bson:"excluded" yaml:"excluded" faker:"-"`
	// ExcludedReason why the result was excluded
	ExcludedReason string `json:"excluded_reason" codec:"excluded_reason" bson:"excluded_reason" yaml:"excluded_reason" faker:"-"`
	// Filename the filename
	Filename string `json:"filename" codec:"filename" bson:"filename" yaml:"filename" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// Language the detected language
	Language string `json:"language" codec:"language" bson:"language" yaml:"language" faker:"-"`
	// License if a license was detected in the file, what was the license SPDX
	License *string `json:"license,omitempty" codec:"license,omitempty" bson:"license" yaml:"license,omitempty" faker:"-"`
	// Lines the individual line attributions
	Lines []BlameLines `json:"lines" codec:"lines" bson:"lines" yaml:"lines" faker:"-"`
	// Loc the count of lines in the file
	Loc int64 `json:"loc" codec:"loc" bson:"loc" yaml:"loc" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RepoID the unique id for the repo
	RepoID string `json:"repo_id" codec:"repo_id" bson:"repo_id" yaml:"repo_id" faker:"-"`
	// Sha the commit SHA
	Sha string `json:"sha" codec:"sha" bson:"sha" yaml:"sha" faker:"-"`
	// Size the size of the file
	Size int64 `json:"size" codec:"size" bson:"size" yaml:"size" faker:"-"`
	// Sloc the count of source lines in the file based on language rules
	Sloc int64 `json:"sloc" codec:"sloc" bson:"sloc" yaml:"sloc" faker:"-"`
	// Status the status of the change
	Status BlameStatus `json:"status" codec:"status" bson:"status" yaml:"status" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Blame)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*Blame)(nil)

func toBlameObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *Blame:
		return v.ToMap()

	case BlameChangeDate:
		return v.ToMap()

	case []BlameLines:
		arr := make([]interface{}, 0)
		for _, i := range v {
			arr = append(arr, i.ToMap())
		}
		return arr

	case BlameStatus:
		return v.String()
	default:
		return o
	}
}

// String returns a string representation of Blame
func (o *Blame) String() string {
	return fmt.Sprintf("sourcecode.Blame<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Blame) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *Blame) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *Blame) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *Blame) GetModelName() datamodel.ModelNameType {
	return BlameModelName
}

// NewBlameID provides a template for generating an ID field for Blame
func NewBlameID(customerID string, refID string, refType string, RepoID string, Filename string) string {
	return hash.Values(customerID, refID, refType, RepoID, Filename)
}

func (o *Blame) setDefaults(frommap bool) {
	if o.Lines == nil {
		o.Lines = make([]BlameLines, 0)
	}

	if o.ID == "" {
		o.ID = hash.Values(o.CustomerID, o.RefID, o.RefType, o.RepoID, o.Filename)
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Blame) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Blame) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Blame) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *Blame) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Blame) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *Blame) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Blame) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Blame) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *Blame) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = BlameModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Blame) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *Blame) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of Blame
func (o *Blame) Clone() datamodel.Model {
	c := new(Blame)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Blame) Anon() datamodel.Model {
	c := new(Blame)
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
func (o *Blame) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Blame) UnmarshalJSON(data []byte) error {
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
func (o *Blame) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *Blame) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two Blame objects are equal
func (o *Blame) IsEqual(other *Blame) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Blame) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"blanks":                  toBlameObject(o.Blanks, false),
		"change_date":             toBlameObject(o.ChangeDate, false),
		"comments":                toBlameObject(o.Comments, false),
		"commit_id":               toBlameObject(o.CommitID, false),
		"complexity":              toBlameObject(o.Complexity, false),
		"customer_id":             toBlameObject(o.CustomerID, false),
		"excluded":                toBlameObject(o.Excluded, false),
		"excluded_reason":         toBlameObject(o.ExcludedReason, false),
		"filename":                toBlameObject(o.Filename, false),
		"id":                      toBlameObject(o.ID, false),
		"integration_instance_id": toBlameObject(o.IntegrationInstanceID, true),
		"language":                toBlameObject(o.Language, false),
		"license":                 toBlameObject(o.License, true),
		"lines":                   toBlameObject(o.Lines, false),
		"loc":                     toBlameObject(o.Loc, false),
		"ref_id":                  toBlameObject(o.RefID, false),
		"ref_type":                toBlameObject(o.RefType, false),
		"repo_id":                 toBlameObject(o.RepoID, false),
		"sha":                     toBlameObject(o.Sha, false),
		"size":                    toBlameObject(o.Size, false),
		"sloc":                    toBlameObject(o.Sloc, false),

		"status":   o.Status.String(),
		"hashcode": toBlameObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *Blame) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["blanks"].(int64); ok {
		o.Blanks = val
	} else {
		if val, ok := kv["blanks"]; ok {
			if val == nil {
				o.Blanks = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Blanks = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["change_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.ChangeDate.FromMap(kv)
		} else if sv, ok := val.(BlameChangeDate); ok {
			// struct
			o.ChangeDate = sv
		} else if sp, ok := val.(*BlameChangeDate); ok {
			// struct pointer
			o.ChangeDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.ChangeDate.Epoch = dt.Epoch
			o.ChangeDate.Rfc3339 = dt.Rfc3339
			o.ChangeDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.ChangeDate.Epoch = dt.Epoch
			o.ChangeDate.Rfc3339 = dt.Rfc3339
			o.ChangeDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.ChangeDate.Epoch = dt.Epoch
				o.ChangeDate.Rfc3339 = dt.Rfc3339
				o.ChangeDate.Offset = dt.Offset
			}
		}
	} else {
		o.ChangeDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["comments"].(int64); ok {
		o.Comments = val
	} else {
		if val, ok := kv["comments"]; ok {
			if val == nil {
				o.Comments = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Comments = number.ToInt64Any(val)
			}
		}
	}
	if val, ok := kv["commit_id"].(string); ok {
		o.CommitID = val
	} else {
		if val, ok := kv["commit_id"]; ok {
			if val == nil {
				o.CommitID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.CommitID = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["complexity"].(int64); ok {
		o.Complexity = val
	} else {
		if val, ok := kv["complexity"]; ok {
			if val == nil {
				o.Complexity = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Complexity = number.ToInt64Any(val)
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
	if val, ok := kv["excluded"].(bool); ok {
		o.Excluded = val
	} else {
		if val, ok := kv["excluded"]; ok {
			if val == nil {
				o.Excluded = false
			} else {
				o.Excluded = number.ToBoolAny(val)
			}
		}
	}
	if val, ok := kv["excluded_reason"].(string); ok {
		o.ExcludedReason = val
	} else {
		if val, ok := kv["excluded_reason"]; ok {
			if val == nil {
				o.ExcludedReason = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.ExcludedReason = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["filename"].(string); ok {
		o.Filename = val
	} else {
		if val, ok := kv["filename"]; ok {
			if val == nil {
				o.Filename = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Filename = fmt.Sprintf("%v", val)
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
	if val, ok := kv["language"].(string); ok {
		o.Language = val
	} else {
		if val, ok := kv["language"]; ok {
			if val == nil {
				o.Language = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Language = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["license"].(*string); ok {
		o.License = val
	} else if val, ok := kv["license"].(string); ok {
		o.License = &val
	} else {
		if val, ok := kv["license"]; ok {
			if val == nil {
				o.License = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.License = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if o == nil {

		o.Lines = make([]BlameLines, 0)

	}
	if val, ok := kv["lines"]; ok {
		if sv, ok := val.([]BlameLines); ok {
			o.Lines = sv
		} else if sp, ok := val.([]*BlameLines); ok {
			o.Lines = o.Lines[:0]
			for _, e := range sp {
				o.Lines = append(o.Lines, *e)
			}
		} else if a, ok := val.(primitive.A); ok {
			for _, ae := range a {
				if av, ok := ae.(BlameLines); ok {
					o.Lines = append(o.Lines, av)
				} else if av, ok := ae.(primitive.M); ok {
					var fm BlameLines
					fm.FromMap(av)
					o.Lines = append(o.Lines, fm)
				} else {
					b, _ := json.Marshal(ae)
					bkv := make(map[string]interface{})
					json.Unmarshal(b, &bkv)
					var av BlameLines
					av.FromMap(bkv)
					o.Lines = append(o.Lines, av)
				}
			}
		} else if arr, ok := val.([]interface{}); ok {
			for _, item := range arr {
				if r, ok := item.(BlameLines); ok {
					o.Lines = append(o.Lines, r)
				} else if r, ok := item.(map[string]interface{}); ok {
					var fm BlameLines
					fm.FromMap(r)
					o.Lines = append(o.Lines, fm)
				} else if r, ok := item.(primitive.M); ok {
					fm := BlameLines{}
					fm.FromMap(r)
					o.Lines = append(o.Lines, fm)
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
							var fm BlameLines
							fm.FromMap(m[0].Interface().(map[string]interface{}))
							o.Lines = append(o.Lines, fm)
						}
					}
				}
			}
		}
	}

	if val, ok := kv["loc"].(int64); ok {
		o.Loc = val
	} else {
		if val, ok := kv["loc"]; ok {
			if val == nil {
				o.Loc = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Loc = number.ToInt64Any(val)
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
	if val, ok := kv["repo_id"].(string); ok {
		o.RepoID = val
	} else {
		if val, ok := kv["repo_id"]; ok {
			if val == nil {
				o.RepoID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.RepoID = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["sha"].(string); ok {
		o.Sha = val
	} else {
		if val, ok := kv["sha"]; ok {
			if val == nil {
				o.Sha = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Sha = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["size"].(int64); ok {
		o.Size = val
	} else {
		if val, ok := kv["size"]; ok {
			if val == nil {
				o.Size = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Size = number.ToInt64Any(val)
			}
		}
	}
	if val, ok := kv["sloc"].(int64); ok {
		o.Sloc = val
	} else {
		if val, ok := kv["sloc"]; ok {
			if val == nil {
				o.Sloc = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Sloc = number.ToInt64Any(val)
			}
		}
	}
	if val, ok := kv["status"].(BlameStatus); ok {
		o.Status = val
	} else {
		if em, ok := kv["status"].(map[string]interface{}); ok {

			ev := em["sourcecode.status"].(string)
			switch ev {
			case "added", "ADDED":
				o.Status = 0
			case "modified", "MODIFIED":
				o.Status = 1
			case "removed", "REMOVED":
				o.Status = 2
			}
		}
		if em, ok := kv["status"].(string); ok {
			switch em {
			case "added", "ADDED":
				o.Status = 0
			case "modified", "MODIFIED":
				o.Status = 1
			case "removed", "REMOVED":
				o.Status = 2
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *Blame) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Blanks)
	args = append(args, o.ChangeDate)
	args = append(args, o.Comments)
	args = append(args, o.CommitID)
	args = append(args, o.Complexity)
	args = append(args, o.CustomerID)
	args = append(args, o.Excluded)
	args = append(args, o.ExcludedReason)
	args = append(args, o.Filename)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.Language)
	args = append(args, o.License)
	args = append(args, o.Lines)
	args = append(args, o.Loc)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RepoID)
	args = append(args, o.Sha)
	args = append(args, o.Size)
	args = append(args, o.Sloc)
	args = append(args, o.Status)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// BlamePartial is a partial struct for upsert mutations for Blame
type BlamePartial struct {
	// Blanks the count of blank lines in the file
	Blanks *int64 `json:"blanks,omitempty"`
	// ChangeDate the date of the change
	ChangeDate *BlameChangeDate `json:"change_date,omitempty"`
	// Comments the count of comment lines in the file based on language rules
	Comments *int64 `json:"comments,omitempty"`
	// CommitID the commit ID
	CommitID *string `json:"commit_id,omitempty"`
	// Complexity the cyclomatic complexity for the change
	Complexity *int64 `json:"complexity,omitempty"`
	// Excluded if the result was excluded
	Excluded *bool `json:"excluded,omitempty"`
	// ExcludedReason why the result was excluded
	ExcludedReason *string `json:"excluded_reason,omitempty"`
	// Filename the filename
	Filename *string `json:"filename,omitempty"`
	// Language the detected language
	Language *string `json:"language,omitempty"`
	// License if a license was detected in the file, what was the license SPDX
	License *string `json:"license,omitempty"`
	// Lines the individual line attributions
	Lines []BlameLines `json:"lines,omitempty"`
	// Loc the count of lines in the file
	Loc *int64 `json:"loc,omitempty"`
	// RepoID the unique id for the repo
	RepoID *string `json:"repo_id,omitempty"`
	// Sha the commit SHA
	Sha *string `json:"sha,omitempty"`
	// Size the size of the file
	Size *int64 `json:"size,omitempty"`
	// Sloc the count of source lines in the file based on language rules
	Sloc *int64 `json:"sloc,omitempty"`
	// Status the status of the change
	Status *BlameStatus `json:"status,omitempty"`
}

var _ datamodel.PartialModel = (*BlamePartial)(nil)

// GetModelName returns the name of the model
func (o *BlamePartial) GetModelName() datamodel.ModelNameType {
	return BlameModelName
}

// ToMap returns the object as a map
func (o *BlamePartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"blanks":          toBlameObject(o.Blanks, true),
		"change_date":     toBlameObject(o.ChangeDate, true),
		"comments":        toBlameObject(o.Comments, true),
		"commit_id":       toBlameObject(o.CommitID, true),
		"complexity":      toBlameObject(o.Complexity, true),
		"excluded":        toBlameObject(o.Excluded, true),
		"excluded_reason": toBlameObject(o.ExcludedReason, true),
		"filename":        toBlameObject(o.Filename, true),
		"language":        toBlameObject(o.Language, true),
		"license":         toBlameObject(o.License, true),
		"lines":           toBlameObject(o.Lines, true),
		"loc":             toBlameObject(o.Loc, true),
		"repo_id":         toBlameObject(o.RepoID, true),
		"sha":             toBlameObject(o.Sha, true),
		"size":            toBlameObject(o.Size, true),
		"sloc":            toBlameObject(o.Sloc, true),

		"status": toBlameStatusEnum(o.Status),
	}
	for k, v := range kv {
		if v == nil || reflect.ValueOf(v).IsZero() {
			delete(kv, k)
		} else {
			if k == "change_date" {
				if dt, ok := v.(*BlameChangeDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}

			if k == "lines" {
				if arr, ok := v.([]BlameLines); ok {
					if len(arr) == 0 {
						delete(kv, k)
					}
				}
			}
		}
	}
	return kv
}

// Stringify returns the object in JSON format as a string
func (o *BlamePartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *BlamePartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *BlamePartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *BlamePartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *BlamePartial) FromMap(kv map[string]interface{}) {
	if val, ok := kv["blanks"].(*int64); ok {
		o.Blanks = val
	} else if val, ok := kv["blanks"].(int64); ok {
		o.Blanks = &val
	} else {
		if val, ok := kv["blanks"]; ok {
			if val == nil {
				o.Blanks = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["long"]
				}
				o.Blanks = number.Int64Pointer(number.ToInt64Any(val))
			}
		}
	}

	if o.ChangeDate == nil {
		o.ChangeDate = &BlameChangeDate{}
	}

	if val, ok := kv["change_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.ChangeDate.FromMap(kv)
		} else if sv, ok := val.(BlameChangeDate); ok {
			// struct
			o.ChangeDate = &sv
		} else if sp, ok := val.(*BlameChangeDate); ok {
			// struct pointer
			o.ChangeDate = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.ChangeDate.Epoch = dt.Epoch
			o.ChangeDate.Rfc3339 = dt.Rfc3339
			o.ChangeDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.ChangeDate.Epoch = dt.Epoch
			o.ChangeDate.Rfc3339 = dt.Rfc3339
			o.ChangeDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.ChangeDate.Epoch = dt.Epoch
				o.ChangeDate.Rfc3339 = dt.Rfc3339
				o.ChangeDate.Offset = dt.Offset
			}
		}
	} else {
		o.ChangeDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["comments"].(*int64); ok {
		o.Comments = val
	} else if val, ok := kv["comments"].(int64); ok {
		o.Comments = &val
	} else {
		if val, ok := kv["comments"]; ok {
			if val == nil {
				o.Comments = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["long"]
				}
				o.Comments = number.Int64Pointer(number.ToInt64Any(val))
			}
		}
	}
	if val, ok := kv["commit_id"].(*string); ok {
		o.CommitID = val
	} else if val, ok := kv["commit_id"].(string); ok {
		o.CommitID = &val
	} else {
		if val, ok := kv["commit_id"]; ok {
			if val == nil {
				o.CommitID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.CommitID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["complexity"].(*int64); ok {
		o.Complexity = val
	} else if val, ok := kv["complexity"].(int64); ok {
		o.Complexity = &val
	} else {
		if val, ok := kv["complexity"]; ok {
			if val == nil {
				o.Complexity = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["long"]
				}
				o.Complexity = number.Int64Pointer(number.ToInt64Any(val))
			}
		}
	}
	if val, ok := kv["excluded"].(*bool); ok {
		o.Excluded = val
	} else if val, ok := kv["excluded"].(bool); ok {
		o.Excluded = &val
	} else {
		if val, ok := kv["excluded"]; ok {
			if val == nil {
				o.Excluded = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Excluded = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}
	if val, ok := kv["excluded_reason"].(*string); ok {
		o.ExcludedReason = val
	} else if val, ok := kv["excluded_reason"].(string); ok {
		o.ExcludedReason = &val
	} else {
		if val, ok := kv["excluded_reason"]; ok {
			if val == nil {
				o.ExcludedReason = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.ExcludedReason = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["filename"].(*string); ok {
		o.Filename = val
	} else if val, ok := kv["filename"].(string); ok {
		o.Filename = &val
	} else {
		if val, ok := kv["filename"]; ok {
			if val == nil {
				o.Filename = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Filename = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["language"].(*string); ok {
		o.Language = val
	} else if val, ok := kv["language"].(string); ok {
		o.Language = &val
	} else {
		if val, ok := kv["language"]; ok {
			if val == nil {
				o.Language = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Language = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["license"].(*string); ok {
		o.License = val
	} else if val, ok := kv["license"].(string); ok {
		o.License = &val
	} else {
		if val, ok := kv["license"]; ok {
			if val == nil {
				o.License = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.License = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if o == nil {

		o.Lines = make([]BlameLines, 0)

	}
	if val, ok := kv["lines"]; ok {
		if sv, ok := val.([]BlameLines); ok {
			o.Lines = sv
		} else if sp, ok := val.([]*BlameLines); ok {
			o.Lines = o.Lines[:0]
			for _, e := range sp {
				o.Lines = append(o.Lines, *e)
			}
		} else if a, ok := val.(primitive.A); ok {
			for _, ae := range a {
				if av, ok := ae.(BlameLines); ok {
					o.Lines = append(o.Lines, av)
				} else if av, ok := ae.(primitive.M); ok {
					var fm BlameLines
					fm.FromMap(av)
					o.Lines = append(o.Lines, fm)
				} else {
					b, _ := json.Marshal(ae)
					bkv := make(map[string]interface{})
					json.Unmarshal(b, &bkv)
					var av BlameLines
					av.FromMap(bkv)
					o.Lines = append(o.Lines, av)
				}
			}
		} else if arr, ok := val.([]interface{}); ok {
			for _, item := range arr {
				if r, ok := item.(BlameLines); ok {
					o.Lines = append(o.Lines, r)
				} else if r, ok := item.(map[string]interface{}); ok {
					var fm BlameLines
					fm.FromMap(r)
					o.Lines = append(o.Lines, fm)
				} else if r, ok := item.(primitive.M); ok {
					fm := BlameLines{}
					fm.FromMap(r)
					o.Lines = append(o.Lines, fm)
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
							var fm BlameLines
							fm.FromMap(m[0].Interface().(map[string]interface{}))
							o.Lines = append(o.Lines, fm)
						}
					}
				}
			}
		}
	}

	if val, ok := kv["loc"].(*int64); ok {
		o.Loc = val
	} else if val, ok := kv["loc"].(int64); ok {
		o.Loc = &val
	} else {
		if val, ok := kv["loc"]; ok {
			if val == nil {
				o.Loc = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["long"]
				}
				o.Loc = number.Int64Pointer(number.ToInt64Any(val))
			}
		}
	}
	if val, ok := kv["repo_id"].(*string); ok {
		o.RepoID = val
	} else if val, ok := kv["repo_id"].(string); ok {
		o.RepoID = &val
	} else {
		if val, ok := kv["repo_id"]; ok {
			if val == nil {
				o.RepoID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.RepoID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["sha"].(*string); ok {
		o.Sha = val
	} else if val, ok := kv["sha"].(string); ok {
		o.Sha = &val
	} else {
		if val, ok := kv["sha"]; ok {
			if val == nil {
				o.Sha = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Sha = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["size"].(*int64); ok {
		o.Size = val
	} else if val, ok := kv["size"].(int64); ok {
		o.Size = &val
	} else {
		if val, ok := kv["size"]; ok {
			if val == nil {
				o.Size = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["long"]
				}
				o.Size = number.Int64Pointer(number.ToInt64Any(val))
			}
		}
	}
	if val, ok := kv["sloc"].(*int64); ok {
		o.Sloc = val
	} else if val, ok := kv["sloc"].(int64); ok {
		o.Sloc = &val
	} else {
		if val, ok := kv["sloc"]; ok {
			if val == nil {
				o.Sloc = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["long"]
				}
				o.Sloc = number.Int64Pointer(number.ToInt64Any(val))
			}
		}
	}
	if val, ok := kv["status"].(*BlameStatus); ok {
		o.Status = val
	} else if val, ok := kv["status"].(BlameStatus); ok {
		o.Status = &val
	} else {
		if val, ok := kv["status"]; ok {
			if val == nil {
				o.Status = toBlameStatusPointer(0)
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["BlameStatus"]
				}
				// this is an enum pointer
				if em, ok := val.(string); ok {
					switch em {
					case "added", "ADDED":
						o.Status = toBlameStatusPointer(0)
					case "modified", "MODIFIED":
						o.Status = toBlameStatusPointer(1)
					case "removed", "REMOVED":
						o.Status = toBlameStatusPointer(2)
					}
				}
			}
		}
	}
	o.setDefaults(false)
}
