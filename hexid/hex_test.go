package hexid

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"gopkg.in/vmihailenco/msgpack.v2"
	"h12.me/uuid"
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
