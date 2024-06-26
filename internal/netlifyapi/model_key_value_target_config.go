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

// checks if the KeyValueTargetConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyValueTargetConfig{}

// KeyValueTargetConfig struct for KeyValueTargetConfig
type KeyValueTargetConfig struct {
	Key string `json:"key"`
	Value string `json:"value"`
	Regex *bool `json:"regex,omitempty"`
	Exclude *bool `json:"exclude,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KeyValueTargetConfig KeyValueTargetConfig

// NewKeyValueTargetConfig instantiates a new KeyValueTargetConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyValueTargetConfig(key string, value string) *KeyValueTargetConfig {
	this := KeyValueTargetConfig{}
	this.Key = key
	this.Value = value
	return &this
}

// NewKeyValueTargetConfigWithDefaults instantiates a new KeyValueTargetConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyValueTargetConfigWithDefaults() *KeyValueTargetConfig {
	this := KeyValueTargetConfig{}
	return &this
}

// GetKey returns the Key field value
func (o *KeyValueTargetConfig) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *KeyValueTargetConfig) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *KeyValueTargetConfig) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *KeyValueTargetConfig) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *KeyValueTargetConfig) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *KeyValueTargetConfig) SetValue(v string) {
	o.Value = v
}

// GetRegex returns the Regex field value if set, zero value otherwise.
func (o *KeyValueTargetConfig) GetRegex() bool {
	if o == nil || IsNil(o.Regex) {
		var ret bool
		return ret
	}
	return *o.Regex
}

// GetRegexOk returns a tuple with the Regex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyValueTargetConfig) GetRegexOk() (*bool, bool) {
	if o == nil || IsNil(o.Regex) {
		return nil, false
	}
	return o.Regex, true
}

// HasRegex returns a boolean if a field has been set.
func (o *KeyValueTargetConfig) HasRegex() bool {
	if o != nil && !IsNil(o.Regex) {
		return true
	}

	return false
}

// SetRegex gets a reference to the given bool and assigns it to the Regex field.
func (o *KeyValueTargetConfig) SetRegex(v bool) {
	o.Regex = &v
}

// GetExclude returns the Exclude field value if set, zero value otherwise.
func (o *KeyValueTargetConfig) GetExclude() bool {
	if o == nil || IsNil(o.Exclude) {
		var ret bool
		return ret
	}
	return *o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyValueTargetConfig) GetExcludeOk() (*bool, bool) {
	if o == nil || IsNil(o.Exclude) {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *KeyValueTargetConfig) HasExclude() bool {
	if o != nil && !IsNil(o.Exclude) {
		return true
	}

	return false
}

// SetExclude gets a reference to the given bool and assigns it to the Exclude field.
func (o *KeyValueTargetConfig) SetExclude(v bool) {
	o.Exclude = &v
}

func (o KeyValueTargetConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyValueTargetConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["value"] = o.Value
	if !IsNil(o.Regex) {
		toSerialize["regex"] = o.Regex
	}
	if !IsNil(o.Exclude) {
		toSerialize["exclude"] = o.Exclude
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KeyValueTargetConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"value",
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

	varKeyValueTargetConfig := _KeyValueTargetConfig{}

	err = json.Unmarshal(data, &varKeyValueTargetConfig)

	if err != nil {
		return err
	}

	*o = KeyValueTargetConfig(varKeyValueTargetConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "key")
		delete(additionalProperties, "value")
		delete(additionalProperties, "regex")
		delete(additionalProperties, "exclude")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKeyValueTargetConfig struct {
	value *KeyValueTargetConfig
	isSet bool
}

func (v NullableKeyValueTargetConfig) Get() *KeyValueTargetConfig {
	return v.value
}

func (v *NullableKeyValueTargetConfig) Set(val *KeyValueTargetConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyValueTargetConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyValueTargetConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyValueTargetConfig(val *KeyValueTargetConfig) *NullableKeyValueTargetConfig {
	return &NullableKeyValueTargetConfig{value: val, isSet: true}
}

func (v NullableKeyValueTargetConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyValueTargetConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


