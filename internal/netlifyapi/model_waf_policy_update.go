/*
Netlify's API documentation

Netlify is a hosting service for the programmable web. It understands your documents and provides an API to handle atomic deploys of websites, manage form submissions, inject JavaScript snippets, and much more. This is a REST-style API that uses JSON for serialization and OAuth 2 for authentication.   This document is an OpenAPI reference for the Netlify API that you can explore. For more detailed instructions for common uses, please visit the [online documentation](https://docs.netlify.com/api/get-started/). Visit our Community forum to join the conversation about [understanding and using Netlify’s API](https://community.netlify.com/t/common-issue-understanding-and-using-netlifys-api/160).   Additionally, we have two API clients for your convenience: - [Go Client](https://github.com/netlify/open-api#go-client) - [JS Client](https://github.com/netlify/js-client) 

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netlifyapi

import (
	"encoding/json"
	"fmt"
)

// checks if the WafPolicyUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WafPolicyUpdate{}

// WafPolicyUpdate struct for WafPolicyUpdate
type WafPolicyUpdate struct {
	PolicyId string `json:"policy_id"`
	AdditionalProperties map[string]interface{}
}

type _WafPolicyUpdate WafPolicyUpdate

// NewWafPolicyUpdate instantiates a new WafPolicyUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafPolicyUpdate(policyId string) *WafPolicyUpdate {
	this := WafPolicyUpdate{}
	this.PolicyId = policyId
	return &this
}

// NewWafPolicyUpdateWithDefaults instantiates a new WafPolicyUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafPolicyUpdateWithDefaults() *WafPolicyUpdate {
	this := WafPolicyUpdate{}
	return &this
}

// GetPolicyId returns the PolicyId field value
func (o *WafPolicyUpdate) GetPolicyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value
// and a boolean to check if the value has been set.
func (o *WafPolicyUpdate) GetPolicyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyId, true
}

// SetPolicyId sets field value
func (o *WafPolicyUpdate) SetPolicyId(v string) {
	o.PolicyId = v
}

func (o WafPolicyUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WafPolicyUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["policy_id"] = o.PolicyId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WafPolicyUpdate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"policy_id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varWafPolicyUpdate := _WafPolicyUpdate{}

	err = json.Unmarshal(data, &varWafPolicyUpdate)

	if err != nil {
		return err
	}

	*o = WafPolicyUpdate(varWafPolicyUpdate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "policy_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWafPolicyUpdate struct {
	value *WafPolicyUpdate
	isSet bool
}

func (v NullableWafPolicyUpdate) Get() *WafPolicyUpdate {
	return v.value
}

func (v *NullableWafPolicyUpdate) Set(val *WafPolicyUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableWafPolicyUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableWafPolicyUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWafPolicyUpdate(val *WafPolicyUpdate) *NullableWafPolicyUpdate {
	return &NullableWafPolicyUpdate{value: val, isSet: true}
}

func (v NullableWafPolicyUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWafPolicyUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


