/*
Web backend for Vikingr

The main api for Vikingr

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// CognitoIdentity struct for CognitoIdentity
type CognitoIdentity struct {
	Sub string `json:"sub"`
	Name string `json:"name"`
	Email string `json:"email"`
}

// NewCognitoIdentity instantiates a new CognitoIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCognitoIdentity(sub string, name string, email string) *CognitoIdentity {
	this := CognitoIdentity{}
	this.Sub = sub
	this.Name = name
	this.Email = email
	return &this
}

// NewCognitoIdentityWithDefaults instantiates a new CognitoIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCognitoIdentityWithDefaults() *CognitoIdentity {
	this := CognitoIdentity{}
	return &this
}

// GetSub returns the Sub field value
func (o *CognitoIdentity) GetSub() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sub
}

// GetSubOk returns a tuple with the Sub field value
// and a boolean to check if the value has been set.
func (o *CognitoIdentity) GetSubOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sub, true
}

// SetSub sets field value
func (o *CognitoIdentity) SetSub(v string) {
	o.Sub = v
}

// GetName returns the Name field value
func (o *CognitoIdentity) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CognitoIdentity) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CognitoIdentity) SetName(v string) {
	o.Name = v
}

// GetEmail returns the Email field value
func (o *CognitoIdentity) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *CognitoIdentity) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *CognitoIdentity) SetEmail(v string) {
	o.Email = v
}

func (o CognitoIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sub"] = o.Sub
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["email"] = o.Email
	}
	return json.Marshal(toSerialize)
}

type NullableCognitoIdentity struct {
	value *CognitoIdentity
	isSet bool
}

func (v NullableCognitoIdentity) Get() *CognitoIdentity {
	return v.value
}

func (v *NullableCognitoIdentity) Set(val *CognitoIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableCognitoIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableCognitoIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCognitoIdentity(val *CognitoIdentity) *NullableCognitoIdentity {
	return &NullableCognitoIdentity{value: val, isSet: true}
}

func (v NullableCognitoIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCognitoIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


