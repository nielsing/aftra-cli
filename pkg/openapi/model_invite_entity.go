/*
Web backend for Vikingr

The main api for Vikingr

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// InviteEntity struct for InviteEntity
type InviteEntity struct {
	Pk string `json:"pk"`
	Sk string `json:"sk"`
	EntityType *string `json:"entityType,omitempty"`
	Created *time.Time `json:"created,omitempty"`
	Updated *time.Time `json:"updated,omitempty"`
	Name string `json:"name"`
	AccessType CompanyAccessEnum `json:"accessType"`
}

// NewInviteEntity instantiates a new InviteEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInviteEntity(pk string, sk string, name string, accessType CompanyAccessEnum) *InviteEntity {
	this := InviteEntity{}
	this.Pk = pk
	this.Sk = sk
	this.Name = name
	this.AccessType = accessType
	return &this
}

// NewInviteEntityWithDefaults instantiates a new InviteEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInviteEntityWithDefaults() *InviteEntity {
	this := InviteEntity{}
	return &this
}

// GetPk returns the Pk field value
func (o *InviteEntity) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *InviteEntity) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *InviteEntity) SetPk(v string) {
	o.Pk = v
}

// GetSk returns the Sk field value
func (o *InviteEntity) GetSk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sk
}

// GetSkOk returns a tuple with the Sk field value
// and a boolean to check if the value has been set.
func (o *InviteEntity) GetSkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sk, true
}

// SetSk sets field value
func (o *InviteEntity) SetSk(v string) {
	o.Sk = v
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *InviteEntity) GetEntityType() string {
	if o == nil || o.EntityType == nil {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteEntity) GetEntityTypeOk() (*string, bool) {
	if o == nil || o.EntityType == nil {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *InviteEntity) HasEntityType() bool {
	if o != nil && o.EntityType != nil {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *InviteEntity) SetEntityType(v string) {
	o.EntityType = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *InviteEntity) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteEntity) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *InviteEntity) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *InviteEntity) SetCreated(v time.Time) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *InviteEntity) GetUpdated() time.Time {
	if o == nil || o.Updated == nil {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteEntity) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || o.Updated == nil {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *InviteEntity) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *InviteEntity) SetUpdated(v time.Time) {
	o.Updated = &v
}

// GetName returns the Name field value
func (o *InviteEntity) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InviteEntity) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InviteEntity) SetName(v string) {
	o.Name = v
}

// GetAccessType returns the AccessType field value
func (o *InviteEntity) GetAccessType() CompanyAccessEnum {
	if o == nil {
		var ret CompanyAccessEnum
		return ret
	}

	return o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value
// and a boolean to check if the value has been set.
func (o *InviteEntity) GetAccessTypeOk() (*CompanyAccessEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessType, true
}

// SetAccessType sets field value
func (o *InviteEntity) SetAccessType(v CompanyAccessEnum) {
	o.AccessType = v
}

func (o InviteEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["sk"] = o.Sk
	}
	if o.EntityType != nil {
		toSerialize["entityType"] = o.EntityType
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Updated != nil {
		toSerialize["updated"] = o.Updated
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["accessType"] = o.AccessType
	}
	return json.Marshal(toSerialize)
}

type NullableInviteEntity struct {
	value *InviteEntity
	isSet bool
}

func (v NullableInviteEntity) Get() *InviteEntity {
	return v.value
}

func (v *NullableInviteEntity) Set(val *InviteEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableInviteEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableInviteEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInviteEntity(val *InviteEntity) *NullableInviteEntity {
	return &NullableInviteEntity{value: val, isSet: true}
}

func (v NullableInviteEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInviteEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


