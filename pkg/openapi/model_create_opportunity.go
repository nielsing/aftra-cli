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

// CreateOpportunity struct for CreateOpportunity
type CreateOpportunity struct {
	Uid string `json:"uid"`
	Name string `json:"name"`
	Details map[string]string `json:"details"`
	Score OpportunityScore `json:"score"`
}

// NewCreateOpportunity instantiates a new CreateOpportunity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOpportunity(uid string, name string, details map[string]string, score OpportunityScore) *CreateOpportunity {
	this := CreateOpportunity{}
	this.Uid = uid
	this.Name = name
	this.Details = details
	this.Score = score
	return &this
}

// NewCreateOpportunityWithDefaults instantiates a new CreateOpportunity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOpportunityWithDefaults() *CreateOpportunity {
	this := CreateOpportunity{}
	return &this
}

// GetUid returns the Uid field value
func (o *CreateOpportunity) GetUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uid
}

// GetUidOk returns a tuple with the Uid field value
// and a boolean to check if the value has been set.
func (o *CreateOpportunity) GetUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uid, true
}

// SetUid sets field value
func (o *CreateOpportunity) SetUid(v string) {
	o.Uid = v
}

// GetName returns the Name field value
func (o *CreateOpportunity) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateOpportunity) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateOpportunity) SetName(v string) {
	o.Name = v
}

// GetDetails returns the Details field value
func (o *CreateOpportunity) GetDetails() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *CreateOpportunity) GetDetailsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *CreateOpportunity) SetDetails(v map[string]string) {
	o.Details = v
}

// GetScore returns the Score field value
func (o *CreateOpportunity) GetScore() OpportunityScore {
	if o == nil {
		var ret OpportunityScore
		return ret
	}

	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *CreateOpportunity) GetScoreOk() (*OpportunityScore, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value
func (o *CreateOpportunity) SetScore(v OpportunityScore) {
	o.Score = v
}

func (o CreateOpportunity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["uid"] = o.Uid
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["details"] = o.Details
	}
	if true {
		toSerialize["score"] = o.Score
	}
	return json.Marshal(toSerialize)
}

type NullableCreateOpportunity struct {
	value *CreateOpportunity
	isSet bool
}

func (v NullableCreateOpportunity) Get() *CreateOpportunity {
	return v.value
}

func (v *NullableCreateOpportunity) Set(val *CreateOpportunity) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOpportunity) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOpportunity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOpportunity(val *CreateOpportunity) *NullableCreateOpportunity {
	return &NullableCreateOpportunity{value: val, isSet: true}
}

func (v NullableCreateOpportunity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOpportunity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


