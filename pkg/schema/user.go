package schema

import (
	"encoding/json"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type User struct {
	Id    string `deadline:"_id" json:"_id,omitempty"`
	Name  string `deadline:"Name" json:"name,omitempty"`
	Email string `deadline:"Email" json:"email,omitempty"`
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewUser(v map[string]interface{}) (*User, error) {
	this := new(User)

	if err := decoder.Decode(v, this); err != nil {
		return nil, err
	} else {
		return this, nil
	}
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (v *User) String() string {
	if data, err := json.MarshalIndent(v, "", "  "); err == nil {
		return string(data)
	} else {
		return err.Error()
	}
}
