package uuid

import (
	"gopkg.in/vmihailenco/msgpack.v2"
)

var (
	MsgPackID int8 = 'u'
)

func init() {
	msgpack.RegisterExt(MsgPackID, UUID{})
}
