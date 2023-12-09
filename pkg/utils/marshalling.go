package utils

import (
	"encoding/json"
	"strings"
)

type Marshaller interface {
	Marshal(interface{}) (string, error)
}

type Unmarshaller interface {
	Unmarshal(string, interface{}) error
}

type JSONMarshaller struct{}

func (j *JSONMarshaller) Unmarshal(obj string, result interface{}) error {
	obj = strings.TrimSpace(obj)
	err := json.Unmarshal([]byte(obj), result)
	if err != nil {
		switch err.(type) {
		case *json.SyntaxError:
			return err
		case *json.UnmarshalTypeError:
			return err
		case *json.UnsupportedTypeError:
			return err
		case *json.UnsupportedValueError:
			return err
		case *json.InvalidUnmarshalError:
			return err
		default:
			return err
		}
	}
	return nil
}

func (j *JSONMarshaller) Marshal(obj interface{}) (string, error) {
	result, err := json.Marshal(obj)
	if err != nil {
		return " ", err
	}
	return string(result), nil
}
