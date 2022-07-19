package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

var (
	ErrWithFirstParam  = errors.New("cannot transform first parameter")
	ErrWithSecondParam = errors.New("cannot load first parameter into second parameter")
)

func ConvertEntity(in, out interface{}) error {
	str, err := json.Marshal(in)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrWithFirstParam, err.Error())
	}

	err = json.Unmarshal(str, out)

	if err != nil {
		return fmt.Errorf("%w: %s", ErrWithSecondParam, err.Error())
	}
	return nil
}

func EntityToJson(entity interface{}) string {
	if IsNilFixed(entity) {
		return "{}"
	}
	str, err := json.Marshal(entity)
	if err != nil {
		return "{}"
	}
	return string(str)
}

func EntityToJsonEscape(entity interface{}) string {
	str, err := json.Marshal(entity)

	buffer := new(bytes.Buffer)
	json.HTMLEscape(buffer, str)

	if err != nil {
		return "{}"
	}
	return string(str)
}

func JsonToEntity(jsonIn string, entity interface{}) {
	err := json.Unmarshal([]byte(jsonIn), entity)

	if err != nil {
		entity = nil
	}
}

func IsNilFixed(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}
