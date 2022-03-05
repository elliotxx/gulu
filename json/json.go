package json

import (
	"encoding/json"
)

// MustMarshal marshal to json string and panic on error
func MustMarshal(v interface{}) []byte {
	r, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return r
}

// MustPrettyMarshal marshal to pretty json string and panic on error
func MustPrettyMarshal(v interface{}) []byte {
	r, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}

	return r
}
