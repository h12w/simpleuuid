package uuid

import (
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
