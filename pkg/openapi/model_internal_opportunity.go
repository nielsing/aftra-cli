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

// InternalOpportunity struct for InternalOpportunity
type InternalOpportunity struct {
	Company string `json:"company"`
	Uid string `json:"uid"`
	Type *string `json:"type,omitempty"`
	Group *string `json:"group,omitempty"`
	Entity *string `json:"entity,omitempty"`
	Score *OpportunityScore `json:"score,omitempty"`
	Resolution *OpportunityResolution `json:"resolution,omitempty"`
	Name string `json:"name"`
	Details map[string]string `json:"details"`
}

// NewInternalOpportunity instantiates a new InternalOpportunity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalOpportunity(company string, uid string, name string, details map[string]string) *InternalOpportunity {
	this := InternalOpportunity{}
	this.Company = company
	this.Uid = uid
	var type_ string = "internal"
	this.Type = &type_
	var group string = "internal"
	this.Group = &group
	this.Name = name
	this.Details = details
	return &this
}

// NewInternalOpportunityWithDefaults instantiates a new InternalOpportunity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalOpportunityWithDefaults() *InternalOpportunity {
	this := InternalOpportunity{}
	var type_ string = "internal"
	this.Type = &type_
	var group string = "internal"
	this.Group = &group
	return &this
}

// GetCompany returns the Company field value
func (o *InternalOpportunity) GetCompany() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Company
}

// GetCompanyOk returns a tuple with the Company field value
// and a boolean to check if the value has been set.
func (o *InternalOpportunity) GetCompanyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Company, true
}

// SetCompany sets field value
func (o *InternalOpportunity) SetCompany(v string) {
	o.Company = v
}

// GetUid returns the Uid field value
func (o *InternalOpportunity) GetUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uid
}

// GetUidOk returns a tuple with the Uid field value
// and a boolean to check if the value has been set.
func (o *InternalOpportunity) GetUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uid, true
}

// SetUid sets field value
func (o *InternalOpportunity) SetUid(v string) {
	o.Uid = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InternalOpportunity) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalOpportunity) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InternalOpportunity) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InternalOpportunity) SetType(v string) {
	o.Type = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *InternalOpportunity) GetGroup() string {
	if o == nil || o.Group == nil {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalOpportunity) GetGroupOk() (*string, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *InternalOpportunity) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *InternalOpportunity) SetGroup(v string) {
	o.Group = &v
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *InternalOpportunity) GetEntity() string {
	if o == nil || o.Entity == nil {
		var ret string
		return ret
	}
	return *o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalOpportunity) GetEntityOk() (*string, bool) {
	if o == nil || o.Entity == nil {
		return nil, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *InternalOpportunity) HasEntity() bool {
	if o != nil && o.Entity != nil {
		return true
	}

	return false
}

// SetEntity gets a reference to the given string and assigns it to the Entity field.
func (o *InternalOpportunity) SetEntity(v string) {
	o.Entity = &v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *InternalOpportunity) GetScore() OpportunityScore {
	if o == nil || o.Score == nil {
		var ret OpportunityScore
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalOpportunity) GetScoreOk() (*OpportunityScore, bool) {
	if o == nil || o.Score == nil {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *InternalOpportunity) HasScore() bool {
	if o != nil && o.Score != nil {
		return true
	}

	return false
}

// SetScore gets a reference to the given OpportunityScore and assigns it to the Score field.
func (o *InternalOpportunity) SetScore(v OpportunityScore) {
	o.Score = &v
}

// GetResolution returns the Resolution field value if set, zero value otherwise.
func (o *InternalOpportunity) GetResolution() OpportunityResolution {
	if o == nil || o.Resolution == nil {
		var ret OpportunityResolution
		return ret
	}
	return *o.Resolution
}

// GetResolutionOk returns a tuple with the Resolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalOpportunity) GetResolutionOk() (*OpportunityResolution, bool) {
	if o == nil || o.Resolution == nil {
		return nil, false
	}
	return o.Resolution, true
}

// HasResolution returns a boolean if a field has been set.
func (o *InternalOpportunity) HasResolution() bool {
	if o != nil && o.Resolution != nil {
		return true
	}

	return false
}

// SetResolution gets a reference to the given OpportunityResolution and assigns it to the Resolution field.
func (o *InternalOpportunity) SetResolution(v OpportunityResolution) {
	o.Resolution = &v
}

// GetName returns the Name field value
func (o *InternalOpportunity) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InternalOpportunity) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InternalOpportunity) SetName(v string) {
	o.Name = v
}

// GetDetails returns the Details field value
func (o *InternalOpportunity) GetDetails() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *InternalOpportunity) GetDetailsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *InternalOpportunity) SetDetails(v map[string]string) {
	o.Details = v
}

func (o InternalOpportunity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["company"] = o.Company
	}
	if true {
		toSerialize["uid"] = o.Uid
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Group != nil {
		toSerialize["group"] = o.Group
	}
	if o.Entity != nil {
		toSerialize["entity"] = o.Entity
	}
	if o.Score != nil {
		toSerialize["score"] = o.Score
	}
	if o.Resolution != nil {
		toSerialize["resolution"] = o.Resolution
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["details"] = o.Details
	}
	return json.Marshal(toSerialize)
}

type NullableInternalOpportunity struct {
	value *InternalOpportunity
	isSet bool
}

func (v NullableInternalOpportunity) Get() *InternalOpportunity {
	return v.value
}

func (v *NullableInternalOpportunity) Set(val *InternalOpportunity) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalOpportunity) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalOpportunity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalOpportunity(val *InternalOpportunity) *NullableInternalOpportunity {
	return &NullableInternalOpportunity{value: val, isSet: true}
}

func (v NullableInternalOpportunity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalOpportunity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


