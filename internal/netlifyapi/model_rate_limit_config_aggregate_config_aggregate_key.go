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

// checks if the RateLimitConfigAggregateConfigAggregateKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RateLimitConfigAggregateConfigAggregateKey{}

// RateLimitConfigAggregateConfigAggregateKey struct for RateLimitConfigAggregateConfigAggregateKey
type RateLimitConfigAggregateConfigAggregateKey struct {
	Type string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _RateLimitConfigAggregateConfigAggregateKey RateLimitConfigAggregateConfigAggregateKey

// NewRateLimitConfigAggregateConfigAggregateKey instantiates a new RateLimitConfigAggregateConfigAggregateKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRateLimitConfigAggregateConfigAggregateKey(type_ string) *RateLimitConfigAggregateConfigAggregateKey {
	this := RateLimitConfigAggregateConfigAggregateKey{}
	this.Type = type_
	return &this
}

// NewRateLimitConfigAggregateConfigAggregateKeyWithDefaults instantiates a new RateLimitConfigAggregateConfigAggregateKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateLimitConfigAggregateConfigAggregateKeyWithDefaults() *RateLimitConfigAggregateConfigAggregateKey {
	this := RateLimitConfigAggregateConfigAggregateKey{}
	return &this
}

// GetType returns the Type field value
func (o *RateLimitConfigAggregateConfigAggregateKey) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RateLimitConfigAggregateConfigAggregateKey) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RateLimitConfigAggregateConfigAggregateKey) SetType(v string) {
	o.Type = v
}

func (o RateLimitConfigAggregateConfigAggregateKey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RateLimitConfigAggregateConfigAggregateKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RateLimitConfigAggregateConfigAggregateKey) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varRateLimitConfigAggregateConfigAggregateKey := _RateLimitConfigAggregateConfigAggregateKey{}

	err = json.Unmarshal(data, &varRateLimitConfigAggregateConfigAggregateKey)

	if err != nil {
		return err
	}

	*o = RateLimitConfigAggregateConfigAggregateKey(varRateLimitConfigAggregateConfigAggregateKey)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRateLimitConfigAggregateConfigAggregateKey struct {
	value *RateLimitConfigAggregateConfigAggregateKey
	isSet bool
}

func (v NullableRateLimitConfigAggregateConfigAggregateKey) Get() *RateLimitConfigAggregateConfigAggregateKey {
	return v.value
}

func (v *NullableRateLimitConfigAggregateConfigAggregateKey) Set(val *RateLimitConfigAggregateConfigAggregateKey) {
	v.value = val
	v.isSet = true
}

func (v NullableRateLimitConfigAggregateConfigAggregateKey) IsSet() bool {
	return v.isSet
}

func (v *NullableRateLimitConfigAggregateConfigAggregateKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRateLimitConfigAggregateConfigAggregateKey(val *RateLimitConfigAggregateConfigAggregateKey) *NullableRateLimitConfigAggregateConfigAggregateKey {
	return &NullableRateLimitConfigAggregateConfigAggregateKey{value: val, isSet: true}
}

func (v NullableRateLimitConfigAggregateConfigAggregateKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRateLimitConfigAggregateConfigAggregateKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


