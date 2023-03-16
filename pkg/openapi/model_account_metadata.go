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

// AccountMetadata struct for AccountMetadata
type AccountMetadata struct {
	ServiceRunningCount *ModelMap `json:"serviceRunningCount,omitempty"`
	RiskScore *float32 `json:"riskScore,omitempty"`
	CampaignConclusion *AccountCampaignConclusion `json:"campaignConclusion,omitempty"`
}

// NewAccountMetadata instantiates a new AccountMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountMetadata() *AccountMetadata {
	this := AccountMetadata{}
	var riskScore float32 = 0
	this.RiskScore = &riskScore
	return &this
}

// NewAccountMetadataWithDefaults instantiates a new AccountMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountMetadataWithDefaults() *AccountMetadata {
	this := AccountMetadata{}
	var riskScore float32 = 0
	this.RiskScore = &riskScore
	return &this
}

// GetServiceRunningCount returns the ServiceRunningCount field value if set, zero value otherwise.
func (o *AccountMetadata) GetServiceRunningCount() ModelMap {
	if o == nil || o.ServiceRunningCount == nil {
		var ret ModelMap
		return ret
	}
	return *o.ServiceRunningCount
}

// GetServiceRunningCountOk returns a tuple with the ServiceRunningCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountMetadata) GetServiceRunningCountOk() (*ModelMap, bool) {
	if o == nil || o.ServiceRunningCount == nil {
		return nil, false
	}
	return o.ServiceRunningCount, true
}

// HasServiceRunningCount returns a boolean if a field has been set.
func (o *AccountMetadata) HasServiceRunningCount() bool {
	if o != nil && o.ServiceRunningCount != nil {
		return true
	}

	return false
}

// SetServiceRunningCount gets a reference to the given ModelMap and assigns it to the ServiceRunningCount field.
func (o *AccountMetadata) SetServiceRunningCount(v ModelMap) {
	o.ServiceRunningCount = &v
}

// GetRiskScore returns the RiskScore field value if set, zero value otherwise.
func (o *AccountMetadata) GetRiskScore() float32 {
	if o == nil || o.RiskScore == nil {
		var ret float32
		return ret
	}
	return *o.RiskScore
}

// GetRiskScoreOk returns a tuple with the RiskScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountMetadata) GetRiskScoreOk() (*float32, bool) {
	if o == nil || o.RiskScore == nil {
		return nil, false
	}
	return o.RiskScore, true
}

// HasRiskScore returns a boolean if a field has been set.
func (o *AccountMetadata) HasRiskScore() bool {
	if o != nil && o.RiskScore != nil {
		return true
	}

	return false
}

// SetRiskScore gets a reference to the given float32 and assigns it to the RiskScore field.
func (o *AccountMetadata) SetRiskScore(v float32) {
	o.RiskScore = &v
}

// GetCampaignConclusion returns the CampaignConclusion field value if set, zero value otherwise.
func (o *AccountMetadata) GetCampaignConclusion() AccountCampaignConclusion {
	if o == nil || o.CampaignConclusion == nil {
		var ret AccountCampaignConclusion
		return ret
	}
	return *o.CampaignConclusion
}

// GetCampaignConclusionOk returns a tuple with the CampaignConclusion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountMetadata) GetCampaignConclusionOk() (*AccountCampaignConclusion, bool) {
	if o == nil || o.CampaignConclusion == nil {
		return nil, false
	}
	return o.CampaignConclusion, true
}

// HasCampaignConclusion returns a boolean if a field has been set.
func (o *AccountMetadata) HasCampaignConclusion() bool {
	if o != nil && o.CampaignConclusion != nil {
		return true
	}

	return false
}

// SetCampaignConclusion gets a reference to the given AccountCampaignConclusion and assigns it to the CampaignConclusion field.
func (o *AccountMetadata) SetCampaignConclusion(v AccountCampaignConclusion) {
	o.CampaignConclusion = &v
}

func (o AccountMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ServiceRunningCount != nil {
		toSerialize["serviceRunningCount"] = o.ServiceRunningCount
	}
	if o.RiskScore != nil {
		toSerialize["riskScore"] = o.RiskScore
	}
	if o.CampaignConclusion != nil {
		toSerialize["campaignConclusion"] = o.CampaignConclusion
	}
	return json.Marshal(toSerialize)
}

type NullableAccountMetadata struct {
	value *AccountMetadata
	isSet bool
}

func (v NullableAccountMetadata) Get() *AccountMetadata {
	return v.value
}

func (v *NullableAccountMetadata) Set(val *AccountMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountMetadata(val *AccountMetadata) *NullableAccountMetadata {
	return &NullableAccountMetadata{value: val, isSet: true}
}

func (v NullableAccountMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


