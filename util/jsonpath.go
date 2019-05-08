package util

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/oliveagle/jsonpath"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/hash"
)

// Action is the signature for handling registered actions
type Action func(args ...interface{}) (interface{}, error)

var (
	actionRe = regexp.MustCompile("^(\\w+)\\((.*)\\)$")
	actions  = make(map[string]Action)
)

// RegisterAction will register a new action
func RegisterAction(name string, action Action) {
	actions[name] = action
}

func isJSONPathNil(val interface{}) bool {
	return val == nil
}

func isJSONPathNotFound(err error) bool {
	return strings.Contains(err.Error(), "not found in object") ||
		strings.Contains(err.Error(), "get attribute from null object")
}

func invokeAction(val string, o interface{}) (interface{}, error) {
	tok := actionRe.FindStringSubmatch(val)
	if tok != nil && len(tok) > 1 {
		actionName := tok[1]
		action := actions[actionName]
		if action == nil {
			return nil, fmt.Errorf("invalid action: %s", actionName)
		}
		args := []interface{}{}
		if len(tok) > 2 {
			if tok[2] != "" {
				for _, arg := range strings.Split(tok[2], ",") {
					arg = strings.TrimSpace(arg)
					if arg[0:1] != "@" && arg[0:1] != "$" {
						args = append(args, arg)
					} else {
						val, err := jsonpath.JsonPathLookup(o, arg)
						if err != nil && !isJSONPathNotFound(err) {
							return nil, fmt.Errorf("error fetching json path: %v. %v", arg, err)
						}
						if isJSONPathNil(val) {
							val = nil
						}
						args = append(args, val)
					}
				}
			}
		}
		return action(args...)
	}
	if val[0:1] != "@" && val[0:1] != "$" {
		return val, nil
	}
	newval, err := jsonpath.JsonPathLookup(o, val)
	if err != nil && !isJSONPathNotFound(err) {
		return nil, err
	}
	if isJSONPathNil(val) {
		return nil, nil
	}
	return newval, nil
}

func dequote(val string) string {
	if val[0:1] == `"` || val[0:1] == `'` {
		return val[1 : len(val)-1]
	}
	return val
}

type actionFunc func(customerid string, val ...interface{}) (interface{}, error)

func init() {
	RegisterAction("epoch", func(args ...interface{}) (interface{}, error) {
		if len(args) == 0 {
			return datetime.EpochNow(), nil
		} else if len(args) == 1 {
			return datetime.ISODateToEpoch(args[0].(string))
		}
		tv, err := time.Parse(dequote(args[1].(string)), args[0].(string))
		if err != nil {
			return nil, err
		}
		return datetime.TimeToEpoch(tv), nil
	})
	RegisterAction("hash", func(args ...interface{}) (interface{}, error) {
		return hash.Values(args...), nil
	})
	RegisterAction("string", func(args ...interface{}) (interface{}, error) {
		return fmt.Sprintf("%s", args[0]), nil
	})
	RegisterAction("len", func(args ...interface{}) (interface{}, error) {
		if val, ok := args[0].([]interface{}); ok {
			return len(val), nil
		}
		if val, ok := args[0].(map[string]interface{}); ok {
			return len(val), nil
		}
		return 0, nil
	})
	RegisterAction("concat", func(args ...interface{}) (interface{}, error) {
		var res string
		for _, e := range args {
			res = res + fmt.Sprint(e)
		}
		return res, nil
	})
}
