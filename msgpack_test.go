package uuid

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"gopkg.in/vmihailenco/msgpack.v2"
)

func TestMsgPack(t *testing.T) {
	u, err := NewTime(time.Now())
	if err != nil {
		t.Fatal(err)
	}
	// error case
	//buf, err := msgpack.Marshal(&u)
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

func TestFailed(t *testing.T) {
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

func TestToJSON(t *testing.T) {
	var s S
	s.ID, _ = NewTime(time.Now())
	buf, err := msgpack.Marshal(&s)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(s.ID)
	fmt.Println(len(buf))
	m := make(map[string]interface{})
	if err := msgpack.Unmarshal(buf, &m); err != nil {
		t.Fatal(err)
	}
	buf, err = json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(buf))

}
