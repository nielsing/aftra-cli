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

// PaginatedEntityCollectionDomainEntity struct for PaginatedEntityCollectionDomainEntity
type PaginatedEntityCollectionDomainEntity struct {
	Entities []DomainEntity `json:"entities"`
	Marker *string `json:"marker,omitempty"`
}

// NewPaginatedEntityCollectionDomainEntity instantiates a new PaginatedEntityCollectionDomainEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedEntityCollectionDomainEntity(entities []DomainEntity) *PaginatedEntityCollectionDomainEntity {
	this := PaginatedEntityCollectionDomainEntity{}
	this.Entities = entities
	return &this
}

// NewPaginatedEntityCollectionDomainEntityWithDefaults instantiates a new PaginatedEntityCollectionDomainEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedEntityCollectionDomainEntityWithDefaults() *PaginatedEntityCollectionDomainEntity {
	this := PaginatedEntityCollectionDomainEntity{}
	return &this
}

// GetEntities returns the Entities field value
func (o *PaginatedEntityCollectionDomainEntity) GetEntities() []DomainEntity {
	if o == nil {
		var ret []DomainEntity
		return ret
	}

	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value
// and a boolean to check if the value has been set.
func (o *PaginatedEntityCollectionDomainEntity) GetEntitiesOk() ([]DomainEntity, bool) {
	if o == nil {
		return nil, false
	}
	return o.Entities, true
}

// SetEntities sets field value
func (o *PaginatedEntityCollectionDomainEntity) SetEntities(v []DomainEntity) {
	o.Entities = v
}

// GetMarker returns the Marker field value if set, zero value otherwise.
func (o *PaginatedEntityCollectionDomainEntity) GetMarker() string {
	if o == nil || o.Marker == nil {
		var ret string
		return ret
	}
	return *o.Marker
}

// GetMarkerOk returns a tuple with the Marker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedEntityCollectionDomainEntity) GetMarkerOk() (*string, bool) {
	if o == nil || o.Marker == nil {
		return nil, false
	}
	return o.Marker, true
}

// HasMarker returns a boolean if a field has been set.
func (o *PaginatedEntityCollectionDomainEntity) HasMarker() bool {
	if o != nil && o.Marker != nil {
		return true
	}

	return false
}

// SetMarker gets a reference to the given string and assigns it to the Marker field.
func (o *PaginatedEntityCollectionDomainEntity) SetMarker(v string) {
	o.Marker = &v
}

func (o PaginatedEntityCollectionDomainEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["entities"] = o.Entities
	}
	if o.Marker != nil {
		toSerialize["marker"] = o.Marker
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedEntityCollectionDomainEntity struct {
	value *PaginatedEntityCollectionDomainEntity
	isSet bool
}

func (v NullablePaginatedEntityCollectionDomainEntity) Get() *PaginatedEntityCollectionDomainEntity {
	return v.value
}

func (v *NullablePaginatedEntityCollectionDomainEntity) Set(val *PaginatedEntityCollectionDomainEntity) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedEntityCollectionDomainEntity) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedEntityCollectionDomainEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedEntityCollectionDomainEntity(val *PaginatedEntityCollectionDomainEntity) *NullablePaginatedEntityCollectionDomainEntity {
	return &NullablePaginatedEntityCollectionDomainEntity{value: val, isSet: true}
}

func (v NullablePaginatedEntityCollectionDomainEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedEntityCollectionDomainEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


