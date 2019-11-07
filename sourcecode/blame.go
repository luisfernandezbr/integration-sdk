// DO NOT EDIT -- generated code

// Package sourcecode - the system which contains source code
package sourcecode

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/bxcodec/faker"
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
	pstrings "github.com/pinpt/go-common/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// BlameTopic is the default topic name
	BlameTopic datamodel.TopicNameType = "sourcecode_Blame_topic"

	// BlameModelName is the model name
	BlameModelName datamodel.ModelNameType = "sourcecode.Blame"
)

const (
	// BlameBlanksColumn is the blanks column name
	BlameBlanksColumn = "Blanks"
	// BlameChangeDateColumn is the change_date column name
	BlameChangeDateColumn = "ChangeDate"
	// BlameChangeDateColumnEpochColumn is the epoch column property of the ChangeDate name
	BlameChangeDateColumnEpochColumn = "ChangeDate.Epoch"
	// BlameChangeDateColumnOffsetColumn is the offset column property of the ChangeDate name
	BlameChangeDateColumnOffsetColumn = "ChangeDate.Offset"
	// BlameChangeDateColumnRfc3339Column is the rfc3339 column property of the ChangeDate name
	BlameChangeDateColumnRfc3339Column = "ChangeDate.Rfc3339"
	// BlameCommentsColumn is the comments column name
	BlameCommentsColumn = "Comments"
	// BlameCommitIDColumn is the commit_id column name
	BlameCommitIDColumn = "CommitID"
	// BlameComplexityColumn is the complexity column name
	BlameComplexityColumn = "Complexity"
	// BlameCustomerIDColumn is the customer_id column name
	BlameCustomerIDColumn = "CustomerID"
	// BlameExcludedColumn is the excluded column name
	BlameExcludedColumn = "Excluded"
	// BlameExcludedReasonColumn is the excluded_reason column name
	BlameExcludedReasonColumn = "ExcludedReason"
	// BlameFilenameColumn is the filename column name
	BlameFilenameColumn = "Filename"
	// BlameIDColumn is the id column name
	BlameIDColumn = "ID"
	// BlameLanguageColumn is the language column name
	BlameLanguageColumn = "Language"
	// BlameLicenseColumn is the license column name
	BlameLicenseColumn = "License"
	// BlameLinesColumn is the lines column name
	BlameLinesColumn = "Lines"
	// BlameLinesColumnAuthorRefIDColumn is the author_ref_id column property of the Lines name
	BlameLinesColumnAuthorRefIDColumn = "Lines.AuthorRefID"
	// BlameLinesColumnBlankColumn is the blank column property of the Lines name
	BlameLinesColumnBlankColumn = "Lines.Blank"
	// BlameLinesColumnCodeColumn is the code column property of the Lines name
	BlameLinesColumnCodeColumn = "Lines.Code"
	// BlameLinesColumnCommentColumn is the comment column property of the Lines name
	BlameLinesColumnCommentColumn = "Lines.Comment"
	// BlameLinesColumnDateColumn is the date column property of the Lines name
	BlameLinesColumnDateColumn = "Lines.Date"
	// BlameLinesColumnShaColumn is the sha column property of the Lines name
	BlameLinesColumnShaColumn = "Lines.Sha"
	// BlameLocColumn is the loc column name
	BlameLocColumn = "Loc"
	// BlameRefIDColumn is the ref_id column name
	BlameRefIDColumn = "RefID"
	// BlameRefTypeColumn is the ref_type column name
	BlameRefTypeColumn = "RefType"
	// BlameRepoIDColumn is the repo_id column name
	BlameRepoIDColumn = "RepoID"
	// BlameShaColumn is the sha column name
	BlameShaColumn = "Sha"
	// BlameSizeColumn is the size column name
	BlameSizeColumn = "Size"
	// BlameSlocColumn is the sloc column name
	BlameSlocColumn = "Sloc"
	// BlameStatusColumn is the status column name
	BlameStatusColumn = "Status"
	// BlameUpdatedAtColumn is the updated_ts column name
	BlameUpdatedAtColumn = "UpdatedAt"
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				o.Blank = number.ToBoolAny(nil)
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
				o.Code = number.ToBoolAny(nil)
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
				o.Comment = number.ToBoolAny(nil)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Sha = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// BlameStatus is the enumeration type for status
type BlameStatus int32

// UnmarshalBSON unmarshals the enum value
func (v BlameStatus) UnmarshalBSON(buf []byte) error {
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

const (
	// StatusAdded is the enumeration value for added
	BlameStatusAdded BlameStatus = 0
	// StatusModified is the enumeration value for modified
	BlameStatusModified BlameStatus = 1
	// StatusRemoved is the enumeration value for removed
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
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
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
	return BlameTopic
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
	if o.License == nil {
		o.License = &emptyString
	}
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
	var i interface{} = o.RepoID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Blame) GetTimestamp() time.Time {
	var dt interface{} = o.ChangeDate
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
	case BlameChangeDate:
		return datetime.DateFromEpoch(v.Epoch)
	}
	panic("not sure how to handle the date time format for Blame")
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
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *Blame) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = BlameModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Blame) GetTopicConfig() *datamodel.ModelTopicConfig {
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
		Key:               "repo_id",
		Timestamp:         "change_date",
		NumPartitions:     8,
		CleanupPolicy:     datamodel.CleanupPolicy("compact"),
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
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

// IsEqual returns true if the two Blame objects are equal
func (o *Blame) IsEqual(other *Blame) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Blame) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"blanks":          toBlameObject(o.Blanks, false),
		"change_date":     toBlameObject(o.ChangeDate, false),
		"comments":        toBlameObject(o.Comments, false),
		"commit_id":       toBlameObject(o.CommitID, false),
		"complexity":      toBlameObject(o.Complexity, false),
		"customer_id":     toBlameObject(o.CustomerID, false),
		"excluded":        toBlameObject(o.Excluded, false),
		"excluded_reason": toBlameObject(o.ExcludedReason, false),
		"filename":        toBlameObject(o.Filename, false),
		"id":              toBlameObject(o.ID, false),
		"language":        toBlameObject(o.Language, false),
		"license":         toBlameObject(o.License, true),
		"lines":           toBlameObject(o.Lines, false),
		"loc":             toBlameObject(o.Loc, false),
		"ref_id":          toBlameObject(o.RefID, false),
		"ref_type":        toBlameObject(o.RefType, false),
		"repo_id":         toBlameObject(o.RepoID, false),
		"sha":             toBlameObject(o.Sha, false),
		"size":            toBlameObject(o.Size, false),
		"sloc":            toBlameObject(o.Sloc, false),
		"status":          toBlameObject(o.Status, false),
		"updated_ts":      toBlameObject(o.UpdatedAt, false),
		"hashcode":        toBlameObject(o.Hashcode, false),
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
				o.Blanks = number.ToInt64Any(nil)
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
				o.Comments = number.ToInt64Any(nil)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				o.Complexity = number.ToInt64Any(nil)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				o.Excluded = number.ToBoolAny(nil)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.ID = fmt.Sprintf("%v", val)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
					var av BlameLines
					if err := json.Unmarshal(b, &av); err != nil {
						panic("unsupported type for lines field entry: " + reflect.TypeOf(ae).String())
					}
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
				o.Loc = number.ToInt64Any(nil)
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

	if val, ok := kv["repo_id"].(string); ok {
		o.RepoID = val
	} else {
		if val, ok := kv["repo_id"]; ok {
			if val == nil {
				o.RepoID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				o.Size = number.ToInt64Any(nil)
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
				o.Sloc = number.ToInt64Any(nil)
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
	args = append(args, o.UpdatedAt)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *Blame) GetEventAPIConfig() datamodel.EventAPIConfig {
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
