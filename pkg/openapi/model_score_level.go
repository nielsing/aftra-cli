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

// ScoreLevel An enumeration.
type ScoreLevel int32

// List of ScoreLevel
const (
	_0 ScoreLevel = 0
	_1 ScoreLevel = 1
	_2 ScoreLevel = 2
	_3 ScoreLevel = 3
	_4 ScoreLevel = 4
)

// All allowed values of ScoreLevel enum
var AllowedScoreLevelEnumValues = []ScoreLevel{
	0,
	1,
	2,
	3,
	4,
}

func (v *ScoreLevel) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ScoreLevel(value)
	for _, existing := range AllowedScoreLevelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ScoreLevel", value)
}

// NewScoreLevelFromValue returns a pointer to a valid ScoreLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewScoreLevelFromValue(v int32) (*ScoreLevel, error) {
	ev := ScoreLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ScoreLevel: valid values are %v", v, AllowedScoreLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ScoreLevel) IsValid() bool {
	for _, existing := range AllowedScoreLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScoreLevel value
func (v ScoreLevel) Ptr() *ScoreLevel {
	return &v
}

type NullableScoreLevel struct {
	value *ScoreLevel
	isSet bool
}

func (v NullableScoreLevel) Get() *ScoreLevel {
	return v.value
}

func (v *NullableScoreLevel) Set(val *ScoreLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableScoreLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableScoreLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScoreLevel(val *ScoreLevel) *NullableScoreLevel {
	return &NullableScoreLevel{value: val, isSet: true}
}

func (v NullableScoreLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScoreLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

