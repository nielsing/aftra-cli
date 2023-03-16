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

// ChallengeTypes An enumeration.
type ChallengeTypes string

// List of ChallengeTypes
const (
	MFA_SETUP ChallengeTypes = "MFA_SETUP"
	SOFTWARE_TOKEN_MFA ChallengeTypes = "SOFTWARE_TOKEN_MFA"
)

// All allowed values of ChallengeTypes enum
var AllowedChallengeTypesEnumValues = []ChallengeTypes{
	"MFA_SETUP",
	"SOFTWARE_TOKEN_MFA",
}

func (v *ChallengeTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ChallengeTypes(value)
	for _, existing := range AllowedChallengeTypesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ChallengeTypes", value)
}

// NewChallengeTypesFromValue returns a pointer to a valid ChallengeTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewChallengeTypesFromValue(v string) (*ChallengeTypes, error) {
	ev := ChallengeTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ChallengeTypes: valid values are %v", v, AllowedChallengeTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ChallengeTypes) IsValid() bool {
	for _, existing := range AllowedChallengeTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ChallengeTypes value
func (v ChallengeTypes) Ptr() *ChallengeTypes {
	return &v
}

type NullableChallengeTypes struct {
	value *ChallengeTypes
	isSet bool
}

func (v NullableChallengeTypes) Get() *ChallengeTypes {
	return v.value
}

func (v *NullableChallengeTypes) Set(val *ChallengeTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableChallengeTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableChallengeTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChallengeTypes(val *ChallengeTypes) *NullableChallengeTypes {
	return &NullableChallengeTypes{value: val, isSet: true}
}

func (v NullableChallengeTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChallengeTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

