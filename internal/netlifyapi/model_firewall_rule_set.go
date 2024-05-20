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

// checks if the FirewallRuleSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FirewallRuleSet{}

// FirewallRuleSet struct for FirewallRuleSet
type FirewallRuleSet struct {
	Default string `json:"default"`
	Rules []FirewallRule `json:"rules"`
	AdditionalProperties map[string]interface{}
}

type _FirewallRuleSet FirewallRuleSet

// NewFirewallRuleSet instantiates a new FirewallRuleSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirewallRuleSet(default_ string, rules []FirewallRule) *FirewallRuleSet {
	this := FirewallRuleSet{}
	this.Default = default_
	this.Rules = rules
	return &this
}

// NewFirewallRuleSetWithDefaults instantiates a new FirewallRuleSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirewallRuleSetWithDefaults() *FirewallRuleSet {
	this := FirewallRuleSet{}
	return &this
}

// GetDefault returns the Default field value
func (o *FirewallRuleSet) GetDefault() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value
// and a boolean to check if the value has been set.
func (o *FirewallRuleSet) GetDefaultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Default, true
}

// SetDefault sets field value
func (o *FirewallRuleSet) SetDefault(v string) {
	o.Default = v
}

// GetRules returns the Rules field value
func (o *FirewallRuleSet) GetRules() []FirewallRule {
	if o == nil {
		var ret []FirewallRule
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *FirewallRuleSet) GetRulesOk() ([]FirewallRule, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *FirewallRuleSet) SetRules(v []FirewallRule) {
	o.Rules = v
}

func (o FirewallRuleSet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FirewallRuleSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["default"] = o.Default
	toSerialize["rules"] = o.Rules

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FirewallRuleSet) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"default",
		"rules",
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

	varFirewallRuleSet := _FirewallRuleSet{}

	err = json.Unmarshal(data, &varFirewallRuleSet)

	if err != nil {
		return err
	}

	*o = FirewallRuleSet(varFirewallRuleSet)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "default")
		delete(additionalProperties, "rules")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFirewallRuleSet struct {
	value *FirewallRuleSet
	isSet bool
}

func (v NullableFirewallRuleSet) Get() *FirewallRuleSet {
	return v.value
}

func (v *NullableFirewallRuleSet) Set(val *FirewallRuleSet) {
	v.value = val
	v.isSet = true
}

func (v NullableFirewallRuleSet) IsSet() bool {
	return v.isSet
}

func (v *NullableFirewallRuleSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirewallRuleSet(val *FirewallRuleSet) *NullableFirewallRuleSet {
	return &NullableFirewallRuleSet{value: val, isSet: true}
}

func (v NullableFirewallRuleSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirewallRuleSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


