package util

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/jhaynie/go-gator/orm"
	"github.com/oliveagle/jsonpath"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/hash"
)

// StreamToMap will stream a file into results
func StreamToMap(fp string, keypath string, results map[string]map[string]interface{}, insert bool) error {
	of, err := os.Open(fp)
	if err != nil {
		return err
	}
	defer of.Close()
	var f io.ReadCloser = of
	if filepath.Ext(fp) == ".gz" {
		gf, err := gzip.NewReader(of)
		if err != nil {
			return err
		}
		f = gf
		defer gf.Close()
	}
	keys := make(map[string]bool)
	// mark all keys as not found
	for k := range results {
		keys[k] = false
	}
	if err := orm.Deserialize(f, func(buf json.RawMessage) error {
		kv := make(map[string]interface{})
		if err := json.Unmarshal(buf, &kv); err != nil {
			return err
		}
		key, err := jsonpath.JsonPathLookup(kv, keypath)
		if err != nil {
			return err
		}
		keystr := key.(string)
		if insert {
			results[keystr] = kv
			keys[keystr] = true
		} else {
			found := results[keystr]
			if found != nil {
				keys[keystr] = true
				for k, v := range kv {
					found[k] = v
				}
			}
		}
		return nil
	}); err != nil {
		return err
	}
	// delete any keys not found in the original source (since this is a join)
	for k, v := range keys {
		if !v {
			delete(results, k)
		}
	}
	return nil
}

// JoinTable will join 2 tables from dir into a map of joined maps
func JoinTable(dir string, fromtable string, frompath string, totable string, topath string, kv map[string]interface{}) (map[string]map[string]interface{}, error) {
	results := make(map[string]map[string]interface{})
	if err := StreamToMap(filepath.Join(dir, fromtable+".json.gz"), frompath, results, true); err != nil {
		return nil, err
	}
	if err := StreamToMap(filepath.Join(dir, totable+".json.gz"), topath, results, false); err != nil {
		return nil, err
	}
	if kv != nil {
		for _, val := range results {
			for k, v := range kv {
				val[k] = v
			}
		}
	}
	return results, nil
}

// Object is a schema model object interface
type Object interface {
	FromMap(kv map[string]interface{})
}

type actionFunc func(customerid string, val interface{}) (interface{}, error)

var actionMap = map[string]actionFunc{
	"epoch()": func(customerid string, val interface{}) (interface{}, error) {
		return datetime.ISODateToEpoch(val.(string))
	},
	"hash()": func(customerid string, val interface{}) (interface{}, error) {
		return hash.Values(customerid, val), nil
	},
	"string()": func(customerid string, val interface{}) (interface{}, error) {
		return fmt.Sprintf("%v", val), nil
	},
}

// CreateObject will create a new model object from kv
func CreateObject(object Object, kv map[string]interface{}, mapping map[string][]string) error {
	newkv := make(map[string]interface{})
	for k, vals := range mapping {
		var lasterr error
		for _, v := range vals {
			lasterr = nil
			res, err := invokeAction(v, kv)
			if err != nil {
				if !strings.Contains(err.Error(), "not found in object") {
					// this is the case where an object value might be null and not found
					lasterr = err
				}
				continue
			}
			newkv[k] = res
		}
		if lasterr != nil {
			return lasterr
		}
	}
	object.FromMap(newkv)
	return nil
}

// CloneMap makes a copy of map
func CloneMap(kv map[string]interface{}) map[string]interface{} {
	newkv := make(map[string]interface{})
	for k, v := range kv {
		newkv[k] = v
	}
	return newkv
}
