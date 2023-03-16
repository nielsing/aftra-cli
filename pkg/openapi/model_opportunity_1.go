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

// Opportunity1 - struct for Opportunity1
type Opportunity1 struct {
	AccountDiscoverability *AccountDiscoverability
	BurpVulnerability *BurpVulnerability
	ExposedFileSecrets *ExposedFileSecrets
	ExposedServices *ExposedServices
	GreenboneVulnerability *GreenboneVulnerability
	InternalOpportunity *InternalOpportunity
	LeakedPassword *LeakedPassword
	SuspiciousDomain *SuspiciousDomain
	WebApplicationVulnerability *WebApplicationVulnerability
}

// AccountDiscoverabilityAsOpportunity1 is a convenience function that returns AccountDiscoverability wrapped in Opportunity1
func AccountDiscoverabilityAsOpportunity1(v *AccountDiscoverability) Opportunity1 {
	return Opportunity1{
		AccountDiscoverability: v,
	}
}

// BurpVulnerabilityAsOpportunity1 is a convenience function that returns BurpVulnerability wrapped in Opportunity1
func BurpVulnerabilityAsOpportunity1(v *BurpVulnerability) Opportunity1 {
	return Opportunity1{
		BurpVulnerability: v,
	}
}

// ExposedFileSecretsAsOpportunity1 is a convenience function that returns ExposedFileSecrets wrapped in Opportunity1
func ExposedFileSecretsAsOpportunity1(v *ExposedFileSecrets) Opportunity1 {
	return Opportunity1{
		ExposedFileSecrets: v,
	}
}

// ExposedServicesAsOpportunity1 is a convenience function that returns ExposedServices wrapped in Opportunity1
func ExposedServicesAsOpportunity1(v *ExposedServices) Opportunity1 {
	return Opportunity1{
		ExposedServices: v,
	}
}

// GreenboneVulnerabilityAsOpportunity1 is a convenience function that returns GreenboneVulnerability wrapped in Opportunity1
func GreenboneVulnerabilityAsOpportunity1(v *GreenboneVulnerability) Opportunity1 {
	return Opportunity1{
		GreenboneVulnerability: v,
	}
}

// InternalOpportunityAsOpportunity1 is a convenience function that returns InternalOpportunity wrapped in Opportunity1
func InternalOpportunityAsOpportunity1(v *InternalOpportunity) Opportunity1 {
	return Opportunity1{
		InternalOpportunity: v,
	}
}

// LeakedPasswordAsOpportunity1 is a convenience function that returns LeakedPassword wrapped in Opportunity1
func LeakedPasswordAsOpportunity1(v *LeakedPassword) Opportunity1 {
	return Opportunity1{
		LeakedPassword: v,
	}
}

// SuspiciousDomainAsOpportunity1 is a convenience function that returns SuspiciousDomain wrapped in Opportunity1
func SuspiciousDomainAsOpportunity1(v *SuspiciousDomain) Opportunity1 {
	return Opportunity1{
		SuspiciousDomain: v,
	}
}

// WebApplicationVulnerabilityAsOpportunity1 is a convenience function that returns WebApplicationVulnerability wrapped in Opportunity1
func WebApplicationVulnerabilityAsOpportunity1(v *WebApplicationVulnerability) Opportunity1 {
	return Opportunity1{
		WebApplicationVulnerability: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *Opportunity1) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AccountDiscoverability
	err = newStrictDecoder(data).Decode(&dst.AccountDiscoverability)
	if err == nil {
		jsonAccountDiscoverability, _ := json.Marshal(dst.AccountDiscoverability)
		if string(jsonAccountDiscoverability) == "{}" { // empty struct
			dst.AccountDiscoverability = nil
		} else {
			match++
		}
	} else {
		dst.AccountDiscoverability = nil
	}

	// try to unmarshal data into BurpVulnerability
	err = newStrictDecoder(data).Decode(&dst.BurpVulnerability)
	if err == nil {
		jsonBurpVulnerability, _ := json.Marshal(dst.BurpVulnerability)
		if string(jsonBurpVulnerability) == "{}" { // empty struct
			dst.BurpVulnerability = nil
		} else {
			match++
		}
	} else {
		dst.BurpVulnerability = nil
	}

	// try to unmarshal data into ExposedFileSecrets
	err = newStrictDecoder(data).Decode(&dst.ExposedFileSecrets)
	if err == nil {
		jsonExposedFileSecrets, _ := json.Marshal(dst.ExposedFileSecrets)
		if string(jsonExposedFileSecrets) == "{}" { // empty struct
			dst.ExposedFileSecrets = nil
		} else {
			match++
		}
	} else {
		dst.ExposedFileSecrets = nil
	}

	// try to unmarshal data into ExposedServices
	err = newStrictDecoder(data).Decode(&dst.ExposedServices)
	if err == nil {
		jsonExposedServices, _ := json.Marshal(dst.ExposedServices)
		if string(jsonExposedServices) == "{}" { // empty struct
			dst.ExposedServices = nil
		} else {
			match++
		}
	} else {
		dst.ExposedServices = nil
	}

	// try to unmarshal data into GreenboneVulnerability
	err = newStrictDecoder(data).Decode(&dst.GreenboneVulnerability)
	if err == nil {
		jsonGreenboneVulnerability, _ := json.Marshal(dst.GreenboneVulnerability)
		if string(jsonGreenboneVulnerability) == "{}" { // empty struct
			dst.GreenboneVulnerability = nil
		} else {
			match++
		}
	} else {
		dst.GreenboneVulnerability = nil
	}

	// try to unmarshal data into InternalOpportunity
	err = newStrictDecoder(data).Decode(&dst.InternalOpportunity)
	if err == nil {
		jsonInternalOpportunity, _ := json.Marshal(dst.InternalOpportunity)
		if string(jsonInternalOpportunity) == "{}" { // empty struct
			dst.InternalOpportunity = nil
		} else {
			match++
		}
	} else {
		dst.InternalOpportunity = nil
	}

	// try to unmarshal data into LeakedPassword
	err = newStrictDecoder(data).Decode(&dst.LeakedPassword)
	if err == nil {
		jsonLeakedPassword, _ := json.Marshal(dst.LeakedPassword)
		if string(jsonLeakedPassword) == "{}" { // empty struct
			dst.LeakedPassword = nil
		} else {
			match++
		}
	} else {
		dst.LeakedPassword = nil
	}

	// try to unmarshal data into SuspiciousDomain
	err = newStrictDecoder(data).Decode(&dst.SuspiciousDomain)
	if err == nil {
		jsonSuspiciousDomain, _ := json.Marshal(dst.SuspiciousDomain)
		if string(jsonSuspiciousDomain) == "{}" { // empty struct
			dst.SuspiciousDomain = nil
		} else {
			match++
		}
	} else {
		dst.SuspiciousDomain = nil
	}

	// try to unmarshal data into WebApplicationVulnerability
	err = newStrictDecoder(data).Decode(&dst.WebApplicationVulnerability)
	if err == nil {
		jsonWebApplicationVulnerability, _ := json.Marshal(dst.WebApplicationVulnerability)
		if string(jsonWebApplicationVulnerability) == "{}" { // empty struct
			dst.WebApplicationVulnerability = nil
		} else {
			match++
		}
	} else {
		dst.WebApplicationVulnerability = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AccountDiscoverability = nil
		dst.BurpVulnerability = nil
		dst.ExposedFileSecrets = nil
		dst.ExposedServices = nil
		dst.GreenboneVulnerability = nil
		dst.InternalOpportunity = nil
		dst.LeakedPassword = nil
		dst.SuspiciousDomain = nil
		dst.WebApplicationVulnerability = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(Opportunity1)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(Opportunity1)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Opportunity1) MarshalJSON() ([]byte, error) {
	if src.AccountDiscoverability != nil {
		return json.Marshal(&src.AccountDiscoverability)
	}

	if src.BurpVulnerability != nil {
		return json.Marshal(&src.BurpVulnerability)
	}

	if src.ExposedFileSecrets != nil {
		return json.Marshal(&src.ExposedFileSecrets)
	}

	if src.ExposedServices != nil {
		return json.Marshal(&src.ExposedServices)
	}

	if src.GreenboneVulnerability != nil {
		return json.Marshal(&src.GreenboneVulnerability)
	}

	if src.InternalOpportunity != nil {
		return json.Marshal(&src.InternalOpportunity)
	}

	if src.LeakedPassword != nil {
		return json.Marshal(&src.LeakedPassword)
	}

	if src.SuspiciousDomain != nil {
		return json.Marshal(&src.SuspiciousDomain)
	}

	if src.WebApplicationVulnerability != nil {
		return json.Marshal(&src.WebApplicationVulnerability)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Opportunity1) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AccountDiscoverability != nil {
		return obj.AccountDiscoverability
	}

	if obj.BurpVulnerability != nil {
		return obj.BurpVulnerability
	}

	if obj.ExposedFileSecrets != nil {
		return obj.ExposedFileSecrets
	}

	if obj.ExposedServices != nil {
		return obj.ExposedServices
	}

	if obj.GreenboneVulnerability != nil {
		return obj.GreenboneVulnerability
	}

	if obj.InternalOpportunity != nil {
		return obj.InternalOpportunity
	}

	if obj.LeakedPassword != nil {
		return obj.LeakedPassword
	}

	if obj.SuspiciousDomain != nil {
		return obj.SuspiciousDomain
	}

	if obj.WebApplicationVulnerability != nil {
		return obj.WebApplicationVulnerability
	}

	// all schemas are nil
	return nil
}

type NullableOpportunity1 struct {
	value *Opportunity1
	isSet bool
}

func (v NullableOpportunity1) Get() *Opportunity1 {
	return v.value
}

func (v *NullableOpportunity1) Set(val *Opportunity1) {
	v.value = val
	v.isSet = true
}

func (v NullableOpportunity1) IsSet() bool {
	return v.isSet
}

func (v *NullableOpportunity1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpportunity1(val *Opportunity1) *NullableOpportunity1 {
	return &NullableOpportunity1{value: val, isSet: true}
}

func (v NullableOpportunity1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpportunity1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


