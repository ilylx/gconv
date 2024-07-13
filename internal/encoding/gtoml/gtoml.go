// Package gtoml provides accessing and converting for TOML content.
package gtoml

import (
	"bytes"
	"github.com/BurntSushi/toml"
	"github.com/ilylx/gconv/internal/json"
)

func Encode(v interface{}) ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	if err := toml.NewEncoder(buffer).Encode(v); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func Decode(v []byte) (interface{}, error) {
	var result interface{}
	if err := toml.Unmarshal(v, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func DecodeTo(v []byte, result interface{}) error {
	return toml.Unmarshal(v, result)
}

func ToJson(v []byte) ([]byte, error) {
	if r, err := Decode(v); err != nil {
		return nil, err
	} else {
		return json.Marshal(r)
	}
}
