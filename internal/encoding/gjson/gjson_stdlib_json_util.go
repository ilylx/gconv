package gjson

import (
	"bytes"
	"github.com/ilylx/gconv"
	"github.com/ilylx/gconv/internal/json"
)

// Valid checks whether <data> is a valid JSON data type.
// The parameter <data> specifies the json format data, which can be either
// bytes or string type.
func Valid(data interface{}) bool {
	return json.Valid(gconv.Bytes(data))
}

// Encode encodes any golang variable <value> to JSON bytes.
func Encode(value interface{}) ([]byte, error) {
	return json.Marshal(value)
}

// Decode decodes json format <data> to golang variable.
// The parameter <data> can be either bytes or string type.
func Decode(data interface{}) (interface{}, error) {
	var value interface{}
	if err := DecodeTo(gconv.Bytes(data), &value); err != nil {
		return nil, err
	} else {
		return value, nil
	}
}

// Decode decodes json format <data> to specified golang variable <v>.
// The parameter <data> can be either bytes or string type.
// The parameter <v> should be a pointer type.
func DecodeTo(data interface{}, v interface{}) error {
	decoder := json.NewDecoder(bytes.NewReader(gconv.Bytes(data)))
	// Do not use number, it converts float64 to json.Number type,
	// which actually a string type. It causes converting issue for other data formats,
	// for example: yaml.
	//decoder.UseNumber()
	return decoder.Decode(v)
}

// DecodeToJson codes json format <data> to a Json object.
// The parameter <data> can be either bytes or string type.
func DecodeToJson(data interface{}, safe ...bool) (*Json, error) {
	if v, err := Decode(gconv.Bytes(data)); err != nil {
		return nil, err
	} else {
		return New(v, safe...), nil
	}
}
