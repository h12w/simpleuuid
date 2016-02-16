package uuid

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"gopkg.in/vmihailenco/msgpack.v2"
)

func TestMsgPack(t *testing.T) {
	u, err := NewTime(time.Now())
	if err != nil {
		t.Fatal(err)
	}
	buf, err := msgpack.Marshal(u)
	if err != nil {
		t.Fatal(err)
	}
	var v UUID
	if err := msgpack.Unmarshal(buf, &v); err != nil {
		t.Fatal(err)
	}
	if u != v {
		t.Fatalf("expect %v, got %v", u, v)
	}
}

func TestMsgPackPtr(t *testing.T) {
	u, err := NewTime(time.Now())
	if err != nil {
		t.Fatal(err)
	}
	buf, err := msgpack.Marshal(&u)
	if err != nil {
		t.Fatal(err)
	}
	var v UUID
	if err := msgpack.Unmarshal(buf, &v); err != nil {
		t.Fatal(err)
	}
	if u != v {
		t.Fatalf("expect %v, got %v", u, v)
	}
}

type S struct {
	ID UUID
}

type HexString []byte

func (s HexString) MarshalJSON() ([]byte, error) {
	return []byte(`"` + hex.EncodeToString([]byte(s)) + `"`), nil
}

func (s *HexString) UnmarshalJSON(bs []byte) error {
	bs, err := hex.DecodeString(strings.Trim(string(bs), `"`))
	if err != nil {
		return err
	}
	*s = HexString(bs)
	return nil
}

func RestoreHexString(any interface{}) interface{} {
	switch o := any.(type) {
	case map[string]interface{}:
		for key, value := range o {
			o[key] = RestoreHexString(value)
		}
	case []byte:
		if len(o) <= size {
			return HexString(o)
		}
	default:
		fmt.Println(reflect.TypeOf(any))
	}
	return any
}

func TestToJSON(t *testing.T) {
	var s S
	s.ID, _ = NewTime(time.Now())
	buf, err := msgpack.Marshal(&s)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(buf))
	fmt.Printf("%#v\n", buf)
	m := make(map[string]interface{})
	if err := msgpack.Unmarshal(buf, &m); err != nil {
		t.Fatal(err)
	}
	buf, err = json.Marshal(RestoreHexString(m))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(buf))
	fmt.Println(string(buf))

}
