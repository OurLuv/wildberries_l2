package model

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type User struct {
	ID     uint64  `json:"id,omitempty"`
	Events []Event `json:"events,omitempty"`
}

func (u *User) Serialize() ([]byte, error) {
	return json.Marshal(u)
}

func (u *User) ParseAndValidate(data url.Values) error {
	idStr := data.Get("user_id")
	if idStr == "" {
		return fmt.Errorf("missing ID parameter")
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("invalid ID parameter: %v", err)
	}
	u.ID = uint64(id)

	return nil
}
