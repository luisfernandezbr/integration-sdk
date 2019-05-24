// DO NOT EDIT -- generated code

// Package customer - customer specific data
package customer

import (
	"bufio"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"regexp"

	"github.com/bxcodec/faker"
	"github.com/linkedin/goavro"
	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	pstrings "github.com/pinpt/go-common/strings"
	"github.com/pinpt/go-datamodel/datamodel"
)

// TeamTopic is the default topic name
const TeamTopic datamodel.TopicNameType = "customer_Team_topic"

// TeamStream is the default stream name
const TeamStream datamodel.TopicNameType = "customer_Team_stream"

// TeamTable is the default table name
const TeamTable datamodel.TopicNameType = "customer_Team"

// TeamModelName is the model name
const TeamModelName datamodel.ModelNameType = "customer.Team"

// Team a team is a grouping of one or more users
type Team struct {
	// built in types

	ID         string `json:"team_id" yaml:"team_id" faker:"-"`
	RefID      string `json:"ref_id" yaml:"ref_id" faker:"-"`
	RefType    string `json:"ref_type" yaml:"ref_type" faker:"-"`
	CustomerID string `json:"customer_id" yaml:"customer_id" faker:"-"`
	Hashcode   string `json:"hashcode" yaml:"hashcode" faker:"-"`
	// custom types

	// Name the name of the team
	Name string `json:"name" yaml:"name" faker:"team"`
	// ParentID the parent id of the team
	ParentID *string `json:"parent_id" yaml:"parent_id" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Team)(nil)

func toTeamObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toTeamObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if o == nil {
		return toTeamObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toTeamObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toTeamObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toTeamObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toTeamObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toTeamObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toTeamObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toTeamObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toTeamObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toTeamObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toTeamObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case map[string]interface{}:
		return o
	case *map[string]interface{}:
		return v
	case *Team:
		return v.ToMap()
	case Team:
		return v.ToMap()
	case []string, []int64, []float64, []bool:
		return o
	case *[]string:
		return (*(o.(*[]string)))
	case *[]int64:
		return (*(o.(*[]int64)))
	case *[]float64:
		return (*(o.(*[]float64)))
	case *[]bool:
		return (*(o.(*[]bool)))
	case []interface{}:
		a := o.([]interface{})
		arr := make([]interface{}, 0)
		for _, av := range a {
			arr = append(arr, toTeamObject(av, isavro, false, ""))
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of Team
func (o *Team) String() string {
	return fmt.Sprintf("customer.Team<%s>", o.ID)
}

func (o *Team) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Team) GetID() string {
	if o.ID == "" {
		// set the id from the spec provided in the model
		o.ID = hash.Values(o.CustomerID, o.Name)
	}
	return o.ID
}

// GetRefID returns the RefID for the object
func (o *Team) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Team) IsMaterialized() bool {
	return true
}

// MaterializedName returns the name of the materialized table
func (o *Team) MaterializedName() string {
	return "customer_team"
}

// Clone returns an exact copy of Team
func (o *Team) Clone() *Team {
	c := new(Team)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Team) Anon() *Team {
	c := new(Team)
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
func (o *Team) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Team) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	// make sure that these have values if empty
	o.setDefaults()
	return nil
}

var cachedCodecTeam *goavro.Codec

// ToAvroBinary returns the data as Avro binary data
func (o *Team) ToAvroBinary() ([]byte, *goavro.Codec, error) {
	if cachedCodecTeam == nil {
		c, err := CreateTeamAvroSchema()
		if err != nil {
			return nil, nil, err
		}
		cachedCodecTeam = c
	}
	kv := o.ToMap(true)
	jbuf, _ := json.Marshal(kv)
	native, _, err := cachedCodecTeam.NativeFromTextual(jbuf)
	if err != nil {
		return nil, nil, err
	}
	// Convert native Go form to binary Avro data
	buf, err := cachedCodecTeam.BinaryFromNative(nil, native)
	return buf, cachedCodecTeam, err
}

// Stringify returns the object in JSON format as a string
func (o *Team) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Team objects are equal
func (o *Team) IsEqual(other *Team) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Team) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	return map[string]interface{}{
		"team_id":     o.GetID(),
		"ref_id":      o.GetRefID(),
		"ref_type":    o.RefType,
		"customer_id": o.CustomerID,
		"hashcode":    o.Hash(),
		"name":        toTeamObject(o.Name, isavro, false, "string"),
		"parent_id":   toTeamObject(o.ParentID, isavro, true, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *Team) FromMap(kv map[string]interface{}) {
	if val, ok := kv["team_id"].(string); ok {
		o.ID = val
	}
	if val, ok := kv["ref_id"].(string); ok {
		o.RefID = val
	}
	if val, ok := kv["ref_type"].(string); ok {
		o.RefType = val
	}
	if val, ok := kv["customer_id"].(string); ok {
		o.CustomerID = val
	}
	if val, ok := kv["name"].(string); ok {
		o.Name = val
	} else {
		val := kv["name"]
		if val == nil {
			o.Name = ""
		} else {
			o.Name = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["parent_id"].(*string); ok {
		o.ParentID = val
	} else if val, ok := kv["parent_id"].(string); ok {
		o.ParentID = &val
	} else {
		val := kv["parent_id"]
		if val == nil {
			o.ParentID = pstrings.Pointer("")
		} else {
			o.ParentID = pstrings.Pointer(fmt.Sprintf("%v", val))
		}
	}
	// make sure that these have values if empty
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *Team) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.GetID())
	args = append(args, o.GetRefID())
	args = append(args, o.RefType)
	args = append(args, o.CustomerID)
	args = append(args, o.Name)
	args = append(args, o.ParentID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// CreateTeam creates a new Team in the database
func CreateTeam(ctx context.Context, db datamodel.Db, o *Team) error {
	return db.Create(ctx, o.ToMap())
}

// DeleteTeam deletes a Team in the database
func DeleteTeam(ctx context.Context, db datamodel.Db, o *Team) error {
	return db.Delete(ctx, o.GetID())
}

// UpdateTeam updates a Team in the database
func UpdateTeam(ctx context.Context, db datamodel.Db, o *Team) error {
	return db.Update(ctx, o.GetID(), o.ToMap())
}

// FindTeam returns a Team from the database
func FindTeam(ctx context.Context, db datamodel.Db, id string) (*Team, error) {
	kv, err := db.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}
	var result Team
	result.FromMap(kv)
	return &result, nil
}

// FindTeams returns all Team from the database matching keys
func FindTeams(ctx context.Context, db datamodel.Db, kv map[string]interface{}) ([]*Team, error) {
	res, err := db.Find(ctx, kv)
	if err != nil {
		return nil, err
	}
	arr := make([]*Team, 0)
	for _, kv := range res {
		var result Team
		result.FromMap(kv)
		arr = append(arr, &result)
	}
	return arr, nil
}

// CreateTeamAvroSchemaSpec creates the avro schema specification for Team
func CreateTeamAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "customer",
		"name":         "Team",
		"connect.name": "customer.Team",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "team_id",
				"type": "string",
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
				"name": "customer_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "hashcode",
				"type": "string",
			},
			map[string]interface{}{
				"name": "name",
				"type": "string",
			},
			map[string]interface{}{
				"name":    "parent_id",
				"type":    []interface{}{"null", "string"},
				"default": nil,
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// CreateTeamAvroSchema creates the avro schema for Team
func CreateTeamAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(CreateTeamAvroSchemaSpec())
}

// TransformTeamFunc is a function for transforming Team during processing
type TransformTeamFunc func(input *Team) (*Team, error)

// CreateTeamPipe creates a pipe for processing Team items
func CreateTeamPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformTeamFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := CreateTeamInputStream(input, errors)
	var stream chan Team
	if len(transforms) > 0 {
		stream = make(chan Team, 1000)
	} else {
		stream = inch
	}
	outdone := CreateTeamOutputStream(output, stream, errors)
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

// CreateTeamInputStreamDir creates a channel for reading Team as JSON newlines from a directory of files
func CreateTeamInputStreamDir(dir string, errors chan<- error, transforms ...TransformTeamFunc) (chan Team, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/customer/team\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan Team)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for team")
		ch := make(chan Team)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return CreateTeamInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan Team)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// CreateTeamInputStreamFile creates an channel for reading Team as JSON newlines from filename
func CreateTeamInputStreamFile(filename string, errors chan<- error, transforms ...TransformTeamFunc) (chan Team, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan Team)
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
			ch := make(chan Team)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return CreateTeamInputStream(f, errors, transforms...)
}

// CreateTeamInputStream creates an channel for reading Team as JSON newlines from stream
func CreateTeamInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformTeamFunc) (chan Team, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan Team, 1000)
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
			var item Team
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

// CreateTeamOutputStreamDir will output json newlines from channel and save in dir
func CreateTeamOutputStreamDir(dir string, ch chan Team, errors chan<- error, transforms ...TransformTeamFunc) <-chan bool {
	fp := filepath.Join(dir, "/customer/team\\.json(\\.gz)?$")
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
	return CreateTeamOutputStream(gz, ch, errors, transforms...)
}

// CreateTeamOutputStream will output json newlines from channel to the stream
func CreateTeamOutputStream(stream io.WriteCloser, ch chan Team, errors chan<- error, transforms ...TransformTeamFunc) <-chan bool {
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

// CreateTeamProducer will stream data from the channel
func CreateTeamProducer(producer datamodel.Producer, ch chan Team, errors chan<- error) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		ctx := context.Background()
		for item := range ch {
			binary, codec, err := item.ToAvroBinary()
			if err != nil {
				errors <- fmt.Errorf("error encoding %s to avro binary data. %v", item.String(), err)
				return
			}
			if err := producer.Send(ctx, codec, []byte(item.ID), binary); err != nil {
				errors <- fmt.Errorf("error sending %s. %v", item.String(), err)
			}
		}
	}()
	return done
}

// CreateTeamConsumer will stream data from the default topic into the provided channel
func CreateTeamConsumer(factory datamodel.ConsumerFactory, topic datamodel.TopicNameType, ch chan Team, errors chan<- error) (<-chan bool, chan<- bool) {
	return CreateTeamConsumerForTopic(factory, TeamTopic, ch, errors)
}

// CreateTeamConsumerForTopic will stream data from the topic into the provided channel
func CreateTeamConsumerForTopic(factory datamodel.ConsumerFactory, topic datamodel.TopicNameType, ch chan Team, errors chan<- error) (<-chan bool, chan<- bool) {
	done := make(chan bool, 1)
	closed := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		callback := datamodel.ConsumerCallback{
			OnDataReceived: func(key []byte, value []byte) error {
				var object Team
				if err := json.Unmarshal(value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into Team: %s", err)
				}
				ch <- object
				return nil
			},
			OnErrorReceived: func(err error) {
				errors <- err
			},
		}
		consumer, err := factory.CreateConsumer(string(topic), callback)
		if err != nil {
			errors <- err
			return
		}
		select {
		case <-closed:
			consumer.Close()
		}
	}()
	return done, closed
}
