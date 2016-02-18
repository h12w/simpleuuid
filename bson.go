package uuid

import (
	"gopkg.in/mgo.v2/bson"
)

func (id *UUID) SetBSON(raw bson.Raw) error {
	var bsonID bson.ObjectId
	if err := raw.Unmarshal(&bsonID); err != nil {
		return err
	}
	return id.UnmarshalText([]byte(bsonID))
}

func (id UUID) GetBSON() (interface{}, error) {
	return bson.ObjectId(id.String()), nil
}
