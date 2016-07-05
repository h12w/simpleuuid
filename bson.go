package uuid

import (
	"gopkg.in/mgo.v2/bson"
)

func (id *UUID) SetBSON(raw bson.Raw) error {
	var s string
	if err := raw.Unmarshal(&s); err != nil {
		return err
	}
	return id.UnmarshalText([]byte(s))
}

func (id UUID) GetBSON() (interface{}, error) {
	return id.String(), nil
}
