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

// OpportunityScore An enumeration.
type OpportunityScore string

// List of OpportunityScore
const (
	NONE OpportunityScore = "none"
	INFO OpportunityScore = "info"
	LOW OpportunityScore = "low"
	MEDIUM OpportunityScore = "medium"
	HIGH OpportunityScore = "high"
	CRITICAL OpportunityScore = "critical"
	UNKNOWN OpportunityScore = "unknown"
)

// All allowed values of OpportunityScore enum
var AllowedOpportunityScoreEnumValues = []OpportunityScore{
	"none",
	"info",
	"low",
	"medium",
	"high",
	"critical",
	"unknown",
}

func (v *OpportunityScore) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OpportunityScore(value)
	for _, existing := range AllowedOpportunityScoreEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OpportunityScore", value)
}

// NewOpportunityScoreFromValue returns a pointer to a valid OpportunityScore
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOpportunityScoreFromValue(v string) (*OpportunityScore, error) {
	ev := OpportunityScore(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OpportunityScore: valid values are %v", v, AllowedOpportunityScoreEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OpportunityScore) IsValid() bool {
	for _, existing := range AllowedOpportunityScoreEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OpportunityScore value
func (v OpportunityScore) Ptr() *OpportunityScore {
	return &v
}

type NullableOpportunityScore struct {
	value *OpportunityScore
	isSet bool
}

func (v NullableOpportunityScore) Get() *OpportunityScore {
	return v.value
}

func (v *NullableOpportunityScore) Set(val *OpportunityScore) {
	v.value = val
	v.isSet = true
}

func (v NullableOpportunityScore) IsSet() bool {
	return v.isSet
}

func (v *NullableOpportunityScore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpportunityScore(val *OpportunityScore) *NullableOpportunityScore {
	return &NullableOpportunityScore{value: val, isSet: true}
}

func (v NullableOpportunityScore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpportunityScore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

