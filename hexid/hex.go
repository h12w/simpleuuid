package hexid

import (
	"encoding/hex"
	"strings"
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
	case []byte: // UUID like ID
		if len(o)%4 == 0 && len(o) <= 16 {
			return hexString(o)
		}
	}
	return any
}
