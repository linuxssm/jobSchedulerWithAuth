package api

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Id string `json:"id"`
	Email string `json:"email"`
}


func NewUser(id string, email string) *User {
	return &User{Id: id, Email: email}
}

func (e *User) Bytes() []byte {
	bt, err := json.Marshal(e)
	if err != nil {
		return nil
	}
	return bt
}

func NewUserFromBytes(b []byte) (*User, error) {
	var e User
	err := json.Unmarshal(b, &e)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func NewEntryFromReq(r *http.Request) *User {
	return &User{Id: r.FormValue("id"), Email: r.FormValue("email")}
}
