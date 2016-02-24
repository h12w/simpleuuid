package hexid

import (
	"encoding/hex"
	"errors"
	"strings"
	"time"
)

var (
	errNotStringMap = errors.New("not a string map")
	errNotTime      = errors.New("not time.Time")
)

type hexString []byte

func (s hexString) MarshalJSON() ([]byte, error) {
	return []byte(`"` + hex.EncodeToString([]byte(s)) + `"`), nil
}

func (s *hexString) UnmarshalJSON(bs []byte) error {
	bs, err := hex.DecodeString(strings.Trim(string(bs), `"`))
	if err != nil {
		return err
	}
	*s = hexString(bs)
	return nil
}

func Restore(any interface{}) interface{} {
	switch o := any.(type) {
	case map[string]interface{}:
		for key, value := range o {
			o[key] = Restore(value)
		}
	case map[interface{}]interface{}:
		m, err := tryStringMap(o)
		if err != nil {
			for key, value := range o {
				o[key] = Restore(value)
			}
		}
		return Restore(m)
	case []interface{}:
		t, err := tryTime(o)
		if err != nil {
			for i := range o {
				o[i] = Restore(o[i])
			}
			return o
		}
		return t
	case time.Time:
		return o.Format(time.RFC3339Nano)
	case []byte: // UUID like ID
		if len(o)%4 == 0 && len(o) <= 16 {
			return hexString(o)
		}
	}
	return any
}
func tryStringMap(m map[interface{}]interface{}) (map[string]interface{}, error) {
	for key := range m {
		if _, isStr := key.(string); !isStr {
			return nil, errNotStringMap
		}
	}
	strMap := make(map[string]interface{})
	for key, value := range m {
		strMap[key.(string)] = value
	}
	return strMap, nil
}
func tryTime(a []interface{}) (time.Time, error) {
	if len(a) != 2 {
		return time.Time{}, errNotTime
	}
	sec, ok := a[0].(uint64)
	if !ok {
		return time.Time{}, errNotTime
	}
	nsec, ok := a[1].(uint64)
	if !ok {
		return time.Time{}, errNotTime
	}
	return time.Unix(int64(sec), int64(nsec)).UTC(), nil
}
