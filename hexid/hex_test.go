package hexid

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"gopkg.in/mgo.v2/bson"
	"gopkg.in/vmihailenco/msgpack.v2"
	"h12.me/decimal"
	"h12.me/uuid"
)

var (
	testTime, _ = time.Parse(time.RFC3339Nano, "2016-02-19T19:18:56.189Z")
)

type S struct {
	ID uuid.UUID
}

func TestToJSON(t *testing.T) {
	var s S
	s.ID, _ = uuid.NewTime(time.Now())
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
	buf, err = json.Marshal(Restore(m))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(buf))
	fmt.Println(string(buf))

}

func TestRestore(t *testing.T) {
	m := map[string]interface{}{"site_oid": []uint8{0x54, 0xaa, 0x4d, 0xbc, 0x4a, 0x2a, 0x38, 0x2f, 0x1b, 0x8b, 0x45, 0x68}, "slot_index": 0x6, "count": 0x1, "site_sid": "spSExzNg", "device": map[interface{}]interface{}{"geo": map[interface{}]interface{}{"country": "CN"}, "os": map[interface{}]interface{}{"name": "android"}, "android": map[interface{}]interface{}{}, "ios": map[interface{}]interface{}{}, "ip": "116.246.13.86", "network": "wifi"}, "slots": []interface{}{map[interface{}]interface{}{"ad_count": 0x1, "index": 0x6}}, "rid": []uint8{0x76, 0xe, 0xdd, 0x9a, 0xd6, 0x12, 0x11, 0xe5, 0x8e, 0xd, 0x8b, 0x42, 0xcd, 0x88, 0xb5, 0xa6}, "created_at": []interface{}{0x56c574b8, 0x1288602f}}
	n := Restore(m)
	fmt.Printf("%#v\n", n)
	_, err := json.Marshal(n)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRestoreTime(t *testing.T) {
	buf, err := msgpack.Marshal(testTime)
	if err != nil {
		t.Fatal(err)
	}
	var v interface{}
	if err := msgpack.Unmarshal(buf, &v); err != nil {
		t.Fatal(err)
	}
	v = Restore(v)
	if !v.(time.Time).Equal(testTime) {
		t.Fatalf("expect %v, got %v", testTime, v)
	}
}

func TestRestoreDecimal(t *testing.T) {
	d := struct {
		D decimal.D `bson:"d"`
	}{D: decimal.Float(1.2)}
	buf, err := bson.Marshal(d)
	if err != nil {
		t.Fatal(err)
	}
	var v interface{}
	if err := bson.Unmarshal(buf, &v); err != nil {
		t.Fatal(err)
	}
	v = Restore(v)
	if !v.(bson.M)["d"].(decimal.D).Equal(d.D) {
		t.Fatalf("expect %v, got %v", testTime, v)
	}
}
