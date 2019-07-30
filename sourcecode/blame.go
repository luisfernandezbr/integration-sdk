// DO NOT EDIT -- generated code

// Package sourcecode - the system which contains source code
package sourcecode

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
	pstrings "github.com/pinpt/go-common/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// BlameTopic is the default topic name
	BlameTopic datamodel.TopicNameType = "sourcecode_Blame_topic"

	// BlameStream is the default stream name
	BlameStream datamodel.TopicNameType = "sourcecode_Blame_stream"

	// BlameTable is the default table name
	BlameTable datamodel.TopicNameType = "sourcecode_Blame"

	// BlameModelName is the model name
	BlameModelName datamodel.ModelNameType = "sourcecode.Blame"
)

const (
	// BlameBlanksColumn is the blanks column name
	BlameBlanksColumn = "blanks"
	// BlameChangeDateColumn is the change_date column name
	BlameChangeDateColumn = "change_date"
	// BlameChangeDateColumnEpochColumn is the epoch column property of the ChangeDate name
	BlameChangeDateColumnEpochColumn = "change_date->epoch"
	// BlameChangeDateColumnOffsetColumn is the offset column property of the ChangeDate name
	BlameChangeDateColumnOffsetColumn = "change_date->offset"
	// BlameChangeDateColumnRfc3339Column is the rfc3339 column property of the ChangeDate name
	BlameChangeDateColumnRfc3339Column = "change_date->rfc3339"
	// BlameCommentsColumn is the comments column name
	BlameCommentsColumn = "comments"
	// BlameCommitIDColumn is the commit_id column name
	BlameCommitIDColumn = "commit_id"
	// BlameComplexityColumn is the complexity column name
	BlameComplexityColumn = "complexity"
	// BlameCustomerIDColumn is the customer_id column name
	BlameCustomerIDColumn = "customer_id"
	// BlameExcludedColumn is the excluded column name
	BlameExcludedColumn = "excluded"
	// BlameExcludedReasonColumn is the excluded_reason column name
	BlameExcludedReasonColumn = "excluded_reason"
	// BlameFilenameColumn is the filename column name
	BlameFilenameColumn = "filename"
	// BlameIDColumn is the id column name
	BlameIDColumn = "id"
	// BlameLanguageColumn is the language column name
	BlameLanguageColumn = "language"
	// BlameLicenseColumn is the license column name
	BlameLicenseColumn = "license"
	// BlameLinesColumn is the lines column name
	BlameLinesColumn = "lines"
	// BlameLinesColumnAuthorRefIDColumn is the author_ref_id column property of the Lines name
	BlameLinesColumnAuthorRefIDColumn = "lines->author_ref_id"
	// BlameLinesColumnBlankColumn is the blank column property of the Lines name
	BlameLinesColumnBlankColumn = "lines->blank"
	// BlameLinesColumnCodeColumn is the code column property of the Lines name
	BlameLinesColumnCodeColumn = "lines->code"
	// BlameLinesColumnCommentColumn is the comment column property of the Lines name
	BlameLinesColumnCommentColumn = "lines->comment"
	// BlameLinesColumnDateColumn is the date column property of the Lines name
	BlameLinesColumnDateColumn = "lines->date"
	// BlameLinesColumnShaColumn is the sha column property of the Lines name
	BlameLinesColumnShaColumn = "lines->sha"
	// BlameLocColumn is the loc column name
	BlameLocColumn = "loc"
	// BlameRefIDColumn is the ref_id column name
	BlameRefIDColumn = "ref_id"
	// BlameRefTypeColumn is the ref_type column name
	BlameRefTypeColumn = "ref_type"
	// BlameRepoIDColumn is the repo_id column name
	BlameRepoIDColumn = "repo_id"
	// BlameShaColumn is the sha column name
	BlameShaColumn = "sha"
	// BlameSizeColumn is the size column name
	BlameSizeColumn = "size"
	// BlameSlocColumn is the sloc column name
	BlameSlocColumn = "sloc"
	// BlameStatusColumn is the status column name
	BlameStatusColumn = "status"
)

// BlameChangeDate represents the object structure for change_date
type BlameChangeDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toBlameChangeDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toBlameChangeDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *BlameChangeDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *BlameChangeDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toBlameChangeDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toBlameChangeDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toBlameChangeDateObject(o.Rfc3339, isavro, false, "string"),
	}
}

func (o *BlameChangeDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *BlameChangeDate) FromMap(kv map[string]interface{}) {

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
	AuthorRefID string `json:"author_ref_id" bson:"author_ref_id" yaml:"author_ref_id" faker:"-"`
	// Blank if the line is a blank line
	Blank bool `json:"blank" bson:"blank" yaml:"blank" faker:"-"`
	// Code if the line is sourcecode
	Code bool `json:"code" bson:"code" yaml:"code" faker:"-"`
	// Comment if the line is a comment
	Comment bool `json:"comment" bson:"comment" yaml:"comment" faker:"-"`
	// Date the change date in RFC3339 format of this line when last changed
	Date string `json:"date" bson:"date" yaml:"date" faker:"-"`
	// Sha the sha when this line was last changed
	Sha string `json:"sha" bson:"sha" yaml:"sha" faker:"-"`
}

func toBlameLinesObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toBlameLinesObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *BlameLines:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *BlameLines) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// AuthorRefID the author ref_id of this line when last changed
		"author_ref_id": toBlameLinesObject(o.AuthorRefID, isavro, false, "string"),
		// Blank if the line is a blank line
		"blank": toBlameLinesObject(o.Blank, isavro, false, "boolean"),
		// Code if the line is sourcecode
		"code": toBlameLinesObject(o.Code, isavro, false, "boolean"),
		// Comment if the line is a comment
		"comment": toBlameLinesObject(o.Comment, isavro, false, "boolean"),
		// Date the change date in RFC3339 format of this line when last changed
		"date": toBlameLinesObject(o.Date, isavro, false, "string"),
		// Sha the sha when this line was last changed
		"sha": toBlameLinesObject(o.Sha, isavro, false, "string"),
	}
}

func (o *BlameLines) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *BlameLines) FromMap(kv map[string]interface{}) {

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
	Blanks int64 `json:"blanks" bson:"blanks" yaml:"blanks" faker:"-"`
	// ChangeDate the date of the change
	ChangeDate BlameChangeDate `json:"change_date" bson:"change_date" yaml:"change_date" faker:"-"`
	// Comments the count of comment lines in the file based on language rules
	Comments int64 `json:"comments" bson:"comments" yaml:"comments" faker:"-"`
	// CommitID the commit ID
	CommitID string `json:"commit_id" bson:"commit_id" yaml:"commit_id" faker:"-"`
	// Complexity the cyclomatic complexity for the change
	Complexity int64 `json:"complexity" bson:"complexity" yaml:"complexity" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Excluded if the result was excluded
	Excluded bool `json:"excluded" bson:"excluded" yaml:"excluded" faker:"-"`
	// ExcludedReason why the result was excluded
	ExcludedReason string `json:"excluded_reason" bson:"excluded_reason" yaml:"excluded_reason" faker:"-"`
	// Filename the filename
	Filename string `json:"filename" bson:"filename" yaml:"filename" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// Language the detected language
	Language string `json:"language" bson:"language" yaml:"language" faker:"-"`
	// License if a license was detected in the file, what was the license SPDX
	License *string `json:"license" bson:"license" yaml:"license" faker:"-"`
	// Lines the individual line attributions
	Lines []BlameLines `json:"lines" bson:"lines" yaml:"lines" faker:"-"`
	// Loc the count of lines in the file
	Loc int64 `json:"loc" bson:"loc" yaml:"loc" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RepoID the unique id for the repo
	RepoID string `json:"repo_id" bson:"repo_id" yaml:"repo_id" faker:"-"`
	// Sha the commit SHA
	Sha string `json:"sha" bson:"sha" yaml:"sha" faker:"-"`
	// Size the size of the file
	Size int64 `json:"size" bson:"size" yaml:"size" faker:"-"`
	// Sloc the count of source lines in the file based on language rules
	Sloc int64 `json:"sloc" bson:"sloc" yaml:"sloc" faker:"-"`
	// Status the status of the change
	Status BlameStatus `json:"status" bson:"status" yaml:"status" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Blame)(nil)

func toBlameObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toBlameObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *Blame:
		return v.ToMap(isavro)

	case BlameChangeDate:
		return v.ToMap(isavro)

	case []BlameLines:
		return v

	case BlameStatus:
		return v.String()
	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
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

// GetModelName returns the name of the model
func (o *Blame) GetModelName() datamodel.ModelNameType {
	return BlameModelName
}

func (o *Blame) setDefaults(frommap bool) {
	if o.License == nil {
		o.License = &emptyString
	}
	if o.Lines == nil {
		o.Lines = make([]BlameLines, 0)
	}

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Blame", o.CustomerID, o.RefType, o.GetRefID())
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
	retention, err := time.ParseDuration("168h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 168h0m0s. " + err.Error())
	}

	ttl, err := time.ParseDuration("0s")
	if err != nil {
		ttl = 0
	}
	return &datamodel.ModelTopicConfig{
		Key:               "repo_id",
		Timestamp:         "change_date",
		NumPartitions:     8,
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetStateKey returns a key for use in state store
func (o *Blame) GetStateKey() string {
	key := "repo_id"
	return fmt.Sprintf("%s_%s", key, o.GetID())
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

var cachedCodecBlame *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *Blame) GetAvroCodec() *goavro.Codec {
	if cachedCodecBlame == nil {
		c, err := GetBlameAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecBlame = c
	}
	return cachedCodecBlame
}

// ToAvroBinary returns the data as Avro binary data
func (o *Blame) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *Blame) FromAvroBinary(value []byte) error {
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
func (o *Blame) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Blame objects are equal
func (o *Blame) IsEqual(other *Blame) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Blame) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
		if o.Lines == nil {
			o.Lines = make([]BlameLines, 0)
		}
	}
	o.setDefaults(false)
	return map[string]interface{}{
		"blanks":          toBlameObject(o.Blanks, isavro, false, "long"),
		"change_date":     toBlameObject(o.ChangeDate, isavro, false, "change_date"),
		"comments":        toBlameObject(o.Comments, isavro, false, "long"),
		"commit_id":       toBlameObject(o.CommitID, isavro, false, "string"),
		"complexity":      toBlameObject(o.Complexity, isavro, false, "long"),
		"customer_id":     toBlameObject(o.CustomerID, isavro, false, "string"),
		"excluded":        toBlameObject(o.Excluded, isavro, false, "boolean"),
		"excluded_reason": toBlameObject(o.ExcludedReason, isavro, false, "string"),
		"filename":        toBlameObject(o.Filename, isavro, false, "string"),
		"id":              toBlameObject(o.ID, isavro, false, "string"),
		"language":        toBlameObject(o.Language, isavro, false, "string"),
		"license":         toBlameObject(o.License, isavro, true, "string"),
		"lines":           toBlameObject(o.Lines, isavro, false, "lines"),
		"loc":             toBlameObject(o.Loc, isavro, false, "long"),
		"ref_id":          toBlameObject(o.RefID, isavro, false, "string"),
		"ref_type":        toBlameObject(o.RefType, isavro, false, "string"),
		"repo_id":         toBlameObject(o.RepoID, isavro, false, "string"),
		"sha":             toBlameObject(o.Sha, isavro, false, "string"),
		"size":            toBlameObject(o.Size, isavro, false, "long"),
		"sloc":            toBlameObject(o.Sloc, isavro, false, "long"),
		"status":          toBlameObject(o.Status, isavro, false, "status"),
		"hashcode":        toBlameObject(o.Hashcode, isavro, false, "string"),
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
				// if coming in as avro union, convert it back
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
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetBlameAvroSchemaSpec creates the avro schema specification for Blame
func GetBlameAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "sourcecode",
		"name":      "Blame",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "hashcode",
				"type": "string",
			},
			map[string]interface{}{
				"name": "blanks",
				"type": "long",
			},
			map[string]interface{}{
				"name": "change_date",
				"type": map[string]interface{}{"type": "record", "name": "change_date", "fields": []interface{}{map[string]interface{}{"type": "long", "name": "epoch", "doc": "the date in epoch format"}, map[string]interface{}{"doc": "the timezone offset from GMT", "type": "long", "name": "offset"}, map[string]interface{}{"type": "string", "name": "rfc3339", "doc": "the date in RFC3339 format"}}, "doc": "the date of the change"},
			},
			map[string]interface{}{
				"name": "comments",
				"type": "long",
			},
			map[string]interface{}{
				"name": "commit_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "complexity",
				"type": "long",
			},
			map[string]interface{}{
				"name": "customer_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "excluded",
				"type": "boolean",
			},
			map[string]interface{}{
				"name": "excluded_reason",
				"type": "string",
			},
			map[string]interface{}{
				"name": "filename",
				"type": "string",
			},
			map[string]interface{}{
				"name": "id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "language",
				"type": "string",
			},
			map[string]interface{}{
				"name":    "license",
				"type":    []interface{}{"null", "string"},
				"default": nil,
			},
			map[string]interface{}{
				"name": "lines",
				"type": map[string]interface{}{"type": "array", "name": "lines", "items": map[string]interface{}{"type": "record", "name": "lines", "fields": []interface{}{map[string]interface{}{"type": "string", "name": "author_ref_id", "doc": "the author ref_id of this line when last changed"}, map[string]interface{}{"type": "boolean", "name": "blank", "doc": "if the line is a blank line"}, map[string]interface{}{"name": "code", "doc": "if the line is sourcecode", "type": "boolean"}, map[string]interface{}{"type": "boolean", "name": "comment", "doc": "if the line is a comment"}, map[string]interface{}{"type": "string", "name": "date", "doc": "the change date in RFC3339 format of this line when last changed"}, map[string]interface{}{"type": "string", "name": "sha", "doc": "the sha when this line was last changed"}}, "doc": "the individual line attributions"}},
			},
			map[string]interface{}{
				"name": "loc",
				"type": "long",
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
				"name": "repo_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "sha",
				"type": "string",
			},
			map[string]interface{}{
				"name": "size",
				"type": "long",
			},
			map[string]interface{}{
				"name": "sloc",
				"type": "long",
			},
			map[string]interface{}{
				"name": "status",
				"type": map[string]interface{}{
					"type":    "enum",
					"name":    "status",
					"symbols": []interface{}{"ADDED", "MODIFIED", "REMOVED"},
				},
			},
		},
	}
	return pjson.Stringify(spec, true)
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

// GetBlameAvroSchema creates the avro schema for Blame
func GetBlameAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetBlameAvroSchemaSpec())
}

// TransformBlameFunc is a function for transforming Blame during processing
type TransformBlameFunc func(input *Blame) (*Blame, error)

// NewBlamePipe creates a pipe for processing Blame items
func NewBlamePipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformBlameFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewBlameInputStream(input, errors)
	var stream chan Blame
	if len(transforms) > 0 {
		stream = make(chan Blame, 1000)
	} else {
		stream = inch
	}
	outdone := NewBlameOutputStream(output, stream, errors)
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

// NewBlameInputStreamDir creates a channel for reading Blame as JSON newlines from a directory of files
func NewBlameInputStreamDir(dir string, errors chan<- error, transforms ...TransformBlameFunc) (chan Blame, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/sourcecode/blame\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan Blame)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for blame")
		ch := make(chan Blame)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewBlameInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan Blame)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewBlameInputStreamFile creates an channel for reading Blame as JSON newlines from filename
func NewBlameInputStreamFile(filename string, errors chan<- error, transforms ...TransformBlameFunc) (chan Blame, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan Blame)
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
			ch := make(chan Blame)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewBlameInputStream(f, errors, transforms...)
}

// NewBlameInputStream creates an channel for reading Blame as JSON newlines from stream
func NewBlameInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformBlameFunc) (chan Blame, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan Blame, 1000)
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
			var item Blame
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

// NewBlameOutputStreamDir will output json newlines from channel and save in dir
func NewBlameOutputStreamDir(dir string, ch chan Blame, errors chan<- error, transforms ...TransformBlameFunc) <-chan bool {
	fp := filepath.Join(dir, "/sourcecode/blame\\.json(\\.gz)?$")
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
	return NewBlameOutputStream(gz, ch, errors, transforms...)
}

// NewBlameOutputStream will output json newlines from channel to the stream
func NewBlameOutputStream(stream io.WriteCloser, ch chan Blame, errors chan<- error, transforms ...TransformBlameFunc) <-chan bool {
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

// BlameSendEvent is an event detail for sending data
type BlameSendEvent struct {
	Blame   *Blame
	headers map[string]string
	time    time.Time
	key     string
}

var _ datamodel.ModelSendEvent = (*BlameSendEvent)(nil)

// Key is the key to use for the message
func (e *BlameSendEvent) Key() string {
	if e.key == "" {
		return e.Blame.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *BlameSendEvent) Object() datamodel.Model {
	return e.Blame
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *BlameSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *BlameSendEvent) Timestamp() time.Time {
	return e.time
}

// BlameSendEventOpts is a function handler for setting opts
type BlameSendEventOpts func(o *BlameSendEvent)

// WithBlameSendEventKey sets the key value to a value different than the object ID
func WithBlameSendEventKey(key string) BlameSendEventOpts {
	return func(o *BlameSendEvent) {
		o.key = key
	}
}

// WithBlameSendEventTimestamp sets the timestamp value
func WithBlameSendEventTimestamp(tv time.Time) BlameSendEventOpts {
	return func(o *BlameSendEvent) {
		o.time = tv
	}
}

// WithBlameSendEventHeader sets the timestamp value
func WithBlameSendEventHeader(key, value string) BlameSendEventOpts {
	return func(o *BlameSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewBlameSendEvent returns a new BlameSendEvent instance
func NewBlameSendEvent(o *Blame, opts ...BlameSendEventOpts) *BlameSendEvent {
	res := &BlameSendEvent{
		Blame: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewBlameProducer will stream data from the channel
func NewBlameProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
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
				if object, ok := item.Object().(*Blame); ok {
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
					errors <- fmt.Errorf("invalid event received. expected an object of type sourcecode.Blame but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewBlameConsumer will stream data from the topic into the provided channel
func NewBlameConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object Blame
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into sourcecode.Blame: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avro data into sourcecode.Blame: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for sourcecode.Blame")
			}

			// ignore messages that have exceeded the TTL
			cfg := object.GetTopicConfig()
			if cfg != nil && cfg.TTL != 0 && msg.Timestamp.UTC().Add(cfg.TTL).Sub(time.Now().UTC()) < 0 {
				return nil
			}
			msg.Codec = object.GetAvroCodec() // match the codec

			ch <- &BlameReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object Blame
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &BlameReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// BlameReceiveEvent is an event detail for receiving data
type BlameReceiveEvent struct {
	Blame   *Blame
	message eventing.Message
	eof     bool
}

var _ datamodel.ModelReceiveEvent = (*BlameReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *BlameReceiveEvent) Object() datamodel.Model {
	return e.Blame
}

// Message returns the underlying message data for the event
func (e *BlameReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *BlameReceiveEvent) EOF() bool {
	return e.eof
}

// BlameProducer implements the datamodel.ModelEventProducer
type BlameProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*BlameProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *BlameProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *BlameProducer) Close() error {
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
func (o *Blame) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *Blame) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &BlameProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewBlameProducer(newctx, producer, ch, errors, empty),
	}
}

// NewBlameProducerChannel returns a channel which can be used for producing Model events
func NewBlameProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewBlameProducerChannelSize(producer, 0, errors)
}

// NewBlameProducerChannelSize returns a channel which can be used for producing Model events
func NewBlameProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &BlameProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewBlameProducer(newctx, producer, ch, errors, empty),
	}
}

// BlameConsumer implements the datamodel.ModelEventConsumer
type BlameConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*BlameConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *BlameConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *BlameConsumer) Close() error {
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
func (o *Blame) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &BlameConsumer{
		ch:       ch,
		callback: NewBlameConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewBlameConsumerChannel returns a consumer channel which can be used to consume Model events
func NewBlameConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &BlameConsumer{
		ch:       ch,
		callback: NewBlameConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
