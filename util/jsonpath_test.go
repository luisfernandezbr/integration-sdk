package util

import (
	"testing"
	"time"

	"github.com/pinpt/go-common/datetime"

	"github.com/oliveagle/jsonpath"
	"github.com/stretchr/testify/assert"
)

func TestInvokeActionStatic(t *testing.T) {
	assert := assert.New(t)
	result, err := invokeAction("hi", map[string]interface{}{})
	assert.NoError(err)
	assert.Equal("hi", result)
}

func TestInvokeActionEpoch(t *testing.T) {
	assert := assert.New(t)
	result, err := invokeAction("epoch($.date)", map[string]interface{}{"date": "2019-01-23T22:09:58Z"})
	assert.NoError(err)
	assert.Equal(int64(1548281398000), result)
	result, err = invokeAction(`epoch($.date, "2006-01-02T15:04:05Z")`, map[string]interface{}{"date": "2019-01-23T22:09:58Z"})
	assert.NoError(err)
	assert.Equal(int64(1548281398000), result)
	result, err = invokeAction(`epoch()`, map[string]interface{}{})
	assert.NoError(err)
	assert.NotNil(result)
	assert.WithinDuration(time.Now(), datetime.DateFromEpoch(result.(int64)), time.Millisecond)
}

func TestInvokeActionString(t *testing.T) {
	assert := assert.New(t)
	result, err := invokeAction("string($.date)", map[string]interface{}{"date": 1548281398000})
	assert.NoError(err)
	assert.Equal("1548281398000", result)
}

func TestInvokeActionHash(t *testing.T) {
	assert := assert.New(t)
	result, err := invokeAction("hash($.date, $.id)", map[string]interface{}{"date": 1548281398000, "id": "234"})
	assert.NoError(err)
	assert.Equal("18121570b4652185", result)
}

func TestInvokeActionStringStaticArg(t *testing.T) {
	assert := assert.New(t)
	result, err := invokeAction("string(1548281398000)", map[string]interface{}{})
	assert.NoError(err)
	assert.Equal("1548281398000", result)
}

func TestInvokeActionNoAction(t *testing.T) {
	assert := assert.New(t)
	result, err := invokeAction("$.foo", map[string]interface{}{"foo": "bar"})
	assert.NoError(err)
	assert.Equal("bar", result)
}

func TestJSONPathWithNil(t *testing.T) {
	assert := assert.New(t)
	val, err := jsonpath.JsonPathLookup(map[string]interface{}{}, "$.foo")
	assert.True(isJSONPathNotFound(err))
	assert.True(isJSONPathNil(val))
	val, err = jsonpath.JsonPathLookup(map[string]interface{}{}, "$.foo.bar")
	assert.True(isJSONPathNotFound(err))
	assert.True(isJSONPathNil(val))
}
