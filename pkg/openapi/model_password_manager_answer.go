/*
Web backend for Vikingr

The main api for Vikingr

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// PasswordManagerAnswer An enumeration.
type PasswordManagerAnswer string

// List of PasswordManagerAnswer
const (
	PROVIDED PasswordManagerAnswer = "provided"
	ADVISED PasswordManagerAnswer = "advised"
	NONE PasswordManagerAnswer = "none"
)

// All allowed values of PasswordManagerAnswer enum
var AllowedPasswordManagerAnswerEnumValues = []PasswordManagerAnswer{
	"provided",
	"advised",
	"none",
}

func (v *PasswordManagerAnswer) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PasswordManagerAnswer(value)
	for _, existing := range AllowedPasswordManagerAnswerEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PasswordManagerAnswer", value)
}

// NewPasswordManagerAnswerFromValue returns a pointer to a valid PasswordManagerAnswer
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPasswordManagerAnswerFromValue(v string) (*PasswordManagerAnswer, error) {
	ev := PasswordManagerAnswer(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PasswordManagerAnswer: valid values are %v", v, AllowedPasswordManagerAnswerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PasswordManagerAnswer) IsValid() bool {
	for _, existing := range AllowedPasswordManagerAnswerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PasswordManagerAnswer value
func (v PasswordManagerAnswer) Ptr() *PasswordManagerAnswer {
	return &v
}

type NullablePasswordManagerAnswer struct {
	value *PasswordManagerAnswer
	isSet bool
}

func (v NullablePasswordManagerAnswer) Get() *PasswordManagerAnswer {
	return v.value
}

func (v *NullablePasswordManagerAnswer) Set(val *PasswordManagerAnswer) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordManagerAnswer) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordManagerAnswer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordManagerAnswer(val *PasswordManagerAnswer) *NullablePasswordManagerAnswer {
	return &NullablePasswordManagerAnswer{value: val, isSet: true}
}

func (v NullablePasswordManagerAnswer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordManagerAnswer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

