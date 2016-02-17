package hexid

import (
	"encoding/hex"
	"errors"
	"strings"
)

var (
	errNotStringMap = errors.New("not a string map")
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
