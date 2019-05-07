package gen

import (
	"bytes"
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"html"
	"html/template"
	"io"
	"strings"

	"github.com/serenize/snaker"
	"golang.org/x/tools/imports"
)

var inittmpl = `// DO NOT EDIT -- generated code

// Package {{ .schema.Name }} - {{ .schema.Description }}
package {{ .schema.Name }}

{{ $schema := .schema -}}

// GenerateKQL will generate SQL statements for all models into file in directory
func GenerateKQL(dir string) error {
	var builder strings.Builder
{{ range $colkey, $colvalue := .schema.Models -}}
	builder.WriteString(Create{{ proper $colkey }}KQLStreamSQL())
	builder.WriteString("\n")
	builder.WriteString(Create{{ proper $colkey }}KQLTableSQL())
	builder.WriteString("\n")
{{ end -}}
	return ioutil.WriteFile(filepath.Join(dir,  "{{ $schema.Name }}.v{{ $schema.Version }}.sql"), []byte(builder.String()), 0644)
}

// GenerateAvroSchemaSpec will generate the Avro schema to directory for each model
func GenerateAvroSchemaSpec(dir string) error {
{{ range $colkey, $colvalue := .schema.Models -}}
	if err := ioutil.WriteFile(filepath.Join(dir,  "{{ $schema.Name }}.v{{ $schema.Version }}.{{  proper $colkey }}.avro"), []byte(Create{{ proper $colkey }}AvroSchemaSpec()), 0644); err != nil {
		return err
	}
{{ end -}}
	return nil
}

func init() {
}
`

var codetmpl = `// DO NOT EDIT -- generated code

// Package {{ .schema.Name }} - {{ .schema.Description }}
package {{ .schema.Name }}

import (
	pjson "github.com/pinpt/go-common/json"
	pstrings "github.com/pinpt/go-common/strings"
	kafka "github.com/dangkaka/go-kafka-avro"
)

{{ $name := proper .modelname -}}
// {{ $name }}DefaultTopic is the default topic name
const {{ $name }}DefaultTopic = "{{ .schema.Name }}_{{ $name }}_topic"

// {{ $name }}DefaultStream is the default stream name
const {{ $name }}DefaultStream = "{{ .schema.Name }}_{{ $name }}_stream"

// {{ $name }}DefaultTable is the default table name
const {{ $name }}DefaultTable = "{{ .schema.Name }}_{{ $name }}"

// {{ $name }} {{ .model.Description }}
type {{ $name }} struct {
	// built in types
	
	ID string {{ tag "id" }}
	RefID string {{ tag "ref_id" }}
	RefType string {{ tag "ref_type" }}
	CustomerID string {{ tag "customer_id" }}
	Hashcode string {{ tag "hashcode" }}
	
	// custom types

{{ range $colkey, $colvalue := .model.Columns -}}
	{{ proper $colvalue.Name }} {{ gotype $colvalue.Type $colvalue.Optional }} {{ tag $colvalue.Name }}
{{ end -}}
}

// String returns a string representation of {{ $name }}
func (o *{{ $name }}) String() string {
	return fmt.Sprintf("{{ .schema.Name }}.v{{ .schema.Version }}.{{ $name }}<%s>", o.ID)
}

func (o *{{ $name }}) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *{{ $name }}) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("{{ $name }}", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetRefID returns the RefID for the object
func (o *{{ $name }}) GetRefID() string {
{{ range $colkey, $colvalue := .model.Columns -}}
{{ if $colvalue.RefID -}}
	o.RefID = o.{{ proper $colvalue.Name }}
{{ end -}}
{{ end -}}
	return o.RefID
}

// MarshalJSON returns the bytes for marshaling to json
func (o *{{ $name }}) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *{{ $name }}) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	// make sure that these have values if empty
	o.setDefaults()
	return nil
}

// Stringify returns the object in JSON format as a string
func (o *{{ $name }}) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two {{ $name }} objects are equal
func (o *{{ $name }}) IsEqual(other *{{ $name }}) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *{{ $name }}) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id": o.GetID(),
		"ref_id": o.GetRefID(),
		"ref_type": o.RefType,
		"customer_id": o.CustomerID,
		"hashcode": o.Hash(),
	{{ range $colkey, $colvalue := .model.Columns -}}
		"{{ $colvalue.Name }}": o.{{ proper $colvalue.Name }},
	{{ end -}}
	}
}

// FromMap attempts to load data into object from a map
func (o *{{ $name }}) FromMap(kv map[string]interface{}) {
	if val, ok := kv["id"].(string); ok {
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
{{ range $colkey, $colvalue := .model.Columns -}}
{{ if $colvalue.Optional -}}
	if val, ok := kv["{{ $colvalue.Name }}"].({{ gotype $colvalue.Type true }}); ok {
		o.{{ proper $colvalue.Name }} = val
	} else if val, ok := kv["{{ $colvalue.Name }}"].({{ gotype $colvalue.Type false }}); ok {
		o.{{ proper $colvalue.Name }} = &val
	} else {
		val := kv["{{ $colvalue.Name }}"]
		if val == nil {
			o.{{ proper $colvalue.Name }} = {{ gotypeconvempty $colvalue }}
		} else {
			o.{{ proper $colvalue.Name }} = {{ gotypeconv $colvalue "val" }}
		}
	}
{{ else -}}
	if val, ok := kv["{{ $colvalue.Name }}"].({{ gotype $colvalue.Type $colvalue.Optional }}); ok {
		o.{{ proper $colvalue.Name }} = val
	} else {
		val := kv["{{ $colvalue.Name }}"]
		if val == nil {
			o.{{ proper $colvalue.Name }} = {{ gotypeconvempty $colvalue }}
		} else {
			o.{{ proper $colvalue.Name }} = {{ gotypeconv $colvalue "val" }}
		}
	}
{{ end -}}
{{ end -}}
	// make sure that these have values if empty
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *{{ $name }}) Hash() string {
	if o.Hashcode == "" {
		args := make([]interface{}, 0)
		args = append(args, o.GetID())
		args = append(args, o.GetRefID())
		args = append(args, o.RefType)
	{{ range $colkey, $colvalue := .model.Columns -}}
		args = append(args, o.{{ proper $colvalue.Name }})
	{{ end -}}	
		o.Hashcode = hash.Values(args...)
	}
	return o.Hashcode
}

// Create{{ $name }}AvroSchemaSpec creates the avro schema specification for {{ $name }}
func Create{{ $name }}AvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type": "record",
		"namespace": "{{ .schema.Name }}.v{{ .schema.Version }}",
		"name": "{{ $name }}",
		"connect.name": "{{ .schema.Name }}.v{{ .schema.Version }}.{{ $name }}",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "id",
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
		{{ range $index, $colvalue := .model.Columns -}}
			map[string]interface{}{
				"name": "{{ $colvalue.Name }}",
				{{ if $colvalue.Optional -}}
				"type": []string{"null", "{{ avrotype $colvalue }}"},
				{{ else -}}
				"type": "{{ avrotype $colvalue }}",
				{{ end -}}
			},
		{{ end -}}		
		},
	}
	return pjson.Stringify(spec, true)
}

// Create{{ $name }}AvroSchema creates the avro schema for {{ $name }}
func Create{{ $name }}AvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(Create{{ $name }}AvroSchemaSpec())
}

// Create{{ $name }}KQLStreamSQL creates KQL Stream SQL for {{ $name }}
func Create{{ $name }}KQLStreamSQL() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("CREATE STREAM %s ", {{ $name }}DefaultStream))
	builder.WriteString(fmt.Sprintf("WITH (KAFKA_TOPIC='%s', VALUE_FORMAT='AVRO', KEY='id'", {{ $name }}DefaultTopic))
	{{ range $colkey, $colvalue := .model.Columns -}}
	{{ if $colvalue.Timestamp -}}
	builder.WriteString(", TIMESTAMP='{{ $colvalue.Name }}'")
	{{ end -}}
	{{ end -}}
	builder.WriteString(");")
	return builder.String()
}

// Create{{ $name }}KQLTableSQL creates KQL Table SQL for {{ $name }}
func Create{{ $name }}KQLTableSQL() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("CREATE TABLE %s ", {{ $name }}DefaultTable))
	builder.WriteString(fmt.Sprintf("WITH (KAFKA_TOPIC='%s', VALUE_FORMAT='AVRO', KEY='id'", {{ $name }}DefaultTopic))
	{{ range $colkey, $colvalue := .model.Columns -}}
	{{ if $colvalue.Timestamp -}}
	builder.WriteString(", TIMESTAMP='{{ $colvalue.Name }}'")
	{{ end -}}
	{{ end -}}
	builder.WriteString(");")
	return builder.String()
}

// Transform{{ $name }}Func is a function for transforming {{ $name }} during processing
type Transform{{ $name }}Func func(input *{{ $name }}) (*{{ $name }}, error)

// Create{{ $name }}Pipe creates a pipe for processing {{ $name }} items
func Create{{ $name }}Pipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...Transform{{ $name }}Func) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := Create{{ $name }}InputStream(input, errors)
	var stream chan {{ $name }}
	if len(transforms) > 0 {
		stream = make(chan {{ $name }}, 1000)
	} else {
		stream = inch
	}
	outdone := Create{{ $name }}OutputStream(output, stream, errors)
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

// Create{{ $name }}InputStreamDir creates a channel for reading {{ $name }} as JSON newlines from a directory of files
func Create{{ $name }}InputStreamDir(dir string, errors chan<- error, transforms ...Transform{{ $name }}Func) (chan {{ $name }}, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/{{ .schema.Name }}/v{{ .schema.Version }}/{{ .modelname }}\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan {{ $name }})
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for {{ .modelname }}")
		ch := make(chan {{ $name }})
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return Create{{ $name }}InputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan {{ $name }})
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// Create{{ $name }}InputStreamFile creates an channel for reading {{ $name }} as JSON newlines from filename
func Create{{ $name }}InputStreamFile(filename string, errors chan<- error, transforms ...Transform{{ $name }}Func) (chan {{ $name }}, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan {{ $name }})
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
			ch := make(chan {{ $name }})
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done	
		}
		f = gz
	}
	return Create{{ $name }}InputStream(f, errors, transforms...)
}

// Create{{ $name }}InputStream creates an channel for reading {{ $name }} as JSON newlines from stream
func Create{{ $name }}InputStream(stream io.ReadCloser, errors chan<- error, transforms ...Transform{{ $name }}Func) (chan {{ $name }}, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan {{ $name }}, 1000)
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
			var item {{ $name }}
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

// Create{{ $name }}OutputStreamDir will output json newlines from channel and save in dir
func Create{{ $name }}OutputStreamDir(dir string, ch chan {{ $name }}, errors chan<- error, transforms ...Transform{{ $name }}Func) <-chan bool {
	fp := filepath.Join(dir, "/{{ .schema.Name }}/v{{ .schema.Version }}/{{ .modelname }}\\.json(\\.gz)?$")
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
	return Create{{ $name }}OutputStream(gz, ch, errors, transforms...)
}

// Create{{ $name }}OutputStream will output json newlines from channel to the stream
func Create{{ $name }}OutputStream(stream io.WriteCloser, ch chan {{ $name }}, errors chan<- error, transforms ...Transform{{ $name }}Func) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() {
			if gz, ok := stream.(*gzip.Writer); ok {
				gz.Flush();
				gz.Close();
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

// Create{{ $name }}Producer will stream data from the channel
func Create{{ $name }}Producer(producer util.Producer, ch chan {{ $name }}, errors chan<- error) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		schemaspec := Create{{ $name }}AvroSchemaSpec()
		ctx := context.Background()
		for item := range ch {
			if err := producer.Send(ctx, schemaspec, []byte(item.ID), []byte(item.Stringify())); err != nil {
				errors <- fmt.Errorf("error sending %s. %v", item.String(), err)
			}
		}
	}()
	return done
}

// Create{{ $name }}Consumer will stream data from the default topic into the provided channel
func Create{{ $name }}Consumer(factory util.ConsumerFactory, topic string, ch chan {{ $name }}, errors chan<- error) (<-chan bool, chan<- bool) {
	return Create{{ $name }}ConsumerForTopic(factory, {{ $name }}DefaultTopic, ch, errors)
}

// Create{{ $name }}ConsumerForTopic will stream data from the topic into the provided channel
func Create{{ $name }}ConsumerForTopic(factory util.ConsumerFactory, topic string, ch chan {{ $name }}, errors chan<- error) (<-chan bool, chan<- bool) {
	done := make(chan bool, 1)
	closed := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		callback := util.ConsumerCallback{
			OnDataReceived: func(key []byte, value []byte) error {
				var object {{ $name }}
				if err := json.Unmarshal(value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into {{ $name }}: %s", err)
				}
				ch <- object
				return nil
			},
			OnErrorReceived: func(err error) {
				errors <- err
			},
		}
		consumer, err := factory.CreateConsumer(topic, callback)
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

`

func pretty(expr string) (string, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "filename.go", expr, parser.ParseComments)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = format.Node(&buf, fset, node)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

// Generate any setup files
func Generate(schema Schema, w io.Writer) error {
	var err error
	tmpl := template.New("schema")
	tmpl = tmpl.Funcs(getFuncs())
	tmpl, err = tmpl.Parse(inittmpl)
	if err != nil {
		return err
	}
	var buf strings.Builder
	if err := tmpl.Execute(&buf, map[string]interface{}{
		"schema": schema,
	}); err != nil {
		return err
	}
	// fmt.Println(html.UnescapeString(buf.String()))
	sbuf, err := pretty(html.UnescapeString(buf.String()))
	if err != nil {
		return err
	}
	options := &imports.Options{
		TabWidth:  8,
		TabIndent: true,
		Comments:  true,
		Fragment:  true,
	}
	nbuf, err := imports.Process("init.go", []byte(sbuf), options)
	if err != nil {
		return err
	}
	if _, err := w.Write(nbuf); err != nil {
		return err
	}
	return nil
}

// ModelName returns a formatted model name
func ModelName(val string) string {
	val = snaker.SnakeToCamel(val)
	return strings.ToUpper(val[0:1]) + val[1:]
}

func getFuncs() template.FuncMap {
	return template.FuncMap{
		"proper": func(val string) string {
			if strings.HasSuffix(val, "_ts") {
				val = strings.Replace(val, "_ts", "At", 1)
			}
			return ModelName(val)
		},
		"tick": func(val string) template.HTML {
			return template.HTML("`" + val + "`")
		},
		"comma": func(i int, l int) string {
			if i+1 < l {
				return ","
			}
			return ""
		},
		"tag": func(val string) template.HTML {
			return template.HTML("`json:" + `"` + val + `"` + " yaml:" + `"` + val + `"` + "`")
		},
		"gotypeconvempty": func(val Column) template.HTML {
			str := `""`
			switch val.Type {
			case StringColumnType:
				if val.Optional {
					str = "pstrings.Pointer(" + str + ")"
				}
			case BooleanColumnType:
				str = `number.ToBoolAny(nil)`
				if val.Optional {
					str = "number.BoolPointer(" + str + ")"
				}
			case IntColumnType, LongColumnType:
				str = `number.ToInt64Any(nil)`
				if val.Optional {
					str = "number.Int64Pointer(" + str + ")"
				}
			case DoubleColumnType, FloatColumnType:
				str = `number.ToFloat64Any(nil)`
				if val.Optional {
					str = "number.Float64Pointer(" + str + ")"
				}
			case NullColumnType, BytesColumnType:
				str = `nil`
			default:
				panic("unusure of the type specified for column type value: " + val.Type)
			}
			return template.HTML(str)
		},
		"gotypeconv": func(val Column, name string) template.HTML {
			str := `""`
			switch val.Type {
			case StringColumnType:
				str = fmt.Sprintf(`fmt.Sprintf("%%v",%s)`, name)
				if val.Optional {
					str = "pstrings.Pointer(" + str + ")"
				}
			case BooleanColumnType:
				str = fmt.Sprintf(`number.ToBoolAny(%s)`, name)
				if val.Optional {
					str = "number.BoolPointer(" + str + ")"
				}
			case IntColumnType, LongColumnType:
				str = fmt.Sprintf(`number.ToInt64Any(%s)`, name)
				if val.Optional {
					str = "number.Int64Pointer(" + str + ")"
				}
			case DoubleColumnType, FloatColumnType:
				str = fmt.Sprintf(`number.ToFloat64Any(%s)`, name)
				if val.Optional {
					str = "number.Float64Pointer(" + str + ")"
				}
			case NullColumnType:
			case BytesColumnType:
			default:
				panic("unusure of the type specified for column type value: " + val.Type)
			}
			return template.HTML(str)
		},
		"avrotype": func(col Column) template.HTML {
			var val string
			switch col.Type {
			case StringColumnType:
				val = "string"
			case BooleanColumnType:
				val = "boolean"
			case IntColumnType, LongColumnType:
				val = "long"
			case DoubleColumnType:
				val = "double"
			case FloatColumnType:
				val = "float"
			case NullColumnType:
				val = "null"
			case BytesColumnType:
				val = "bytes"
			default:
				panic("unusure of the type specified for column type value: " + val)
			}
			return template.HTML(string(val))
		},
		"gotype": func(val ColumnType, optional bool) template.HTML {
			switch val {
			case StringColumnType:
				val = "string"
			case BooleanColumnType:
				val = "bool"
			case IntColumnType, LongColumnType:
				val = "int64"
			case DoubleColumnType, FloatColumnType:
				val = "float64"
			case NullColumnType:
				val = "interface{}"
			case BytesColumnType:
				val = "[]byte"
			default:
				panic("unusure of the type specified for column type value: " + val)
			}
			if optional {
				return template.HTML("*" + string(val))
			}
			return template.HTML(string(val))
		},
	}
}

// GenerateModel a set of schemas
func GenerateModel(schema Schema, model string, w io.Writer) error {
	var err error
	tmpl := template.New("model")
	tmpl = tmpl.Funcs(getFuncs())
	tmpl, err = tmpl.Parse(codetmpl)
	if err != nil {
		return err
	}
	var buf strings.Builder
	if err := tmpl.Execute(&buf, map[string]interface{}{
		"schema":    schema,
		"modelname": model,
		"model":     schema.Models[model],
	}); err != nil {
		return err
	}
	// fmt.Println(html.UnescapeString(buf.String()))
	sbuf, err := pretty(html.UnescapeString(buf.String()))
	if err != nil {
		return err
	}
	options := &imports.Options{
		TabWidth:  8,
		TabIndent: true,
		Comments:  true,
		Fragment:  true,
	}
	nbuf, err := imports.Process(model+".go", []byte(sbuf), options)
	if err != nil {
		return err
	}
	if _, err := w.Write(nbuf); err != nil {
		return err
	}
	return nil
}
