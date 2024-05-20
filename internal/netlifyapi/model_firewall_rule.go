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

// checks if the FirewallRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FirewallRule{}

// FirewallRule struct for FirewallRule
type FirewallRule struct {
	Action string `json:"action"`
	Type string `json:"type"`
	Data map[string][]string `json:"data"`
	Description *string `json:"description,omitempty"`
	Disabled *bool `json:"disabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirewallRule FirewallRule

// NewFirewallRule instantiates a new FirewallRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirewallRule(action string, type_ string, data map[string][]string) *FirewallRule {
	this := FirewallRule{}
	this.Action = action
	this.Type = type_
	this.Data = data
	return &this
}

// NewFirewallRuleWithDefaults instantiates a new FirewallRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirewallRuleWithDefaults() *FirewallRule {
	this := FirewallRule{}
	return &this
}

// GetAction returns the Action field value
func (o *FirewallRule) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *FirewallRule) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *FirewallRule) SetAction(v string) {
	o.Action = v
}

// GetType returns the Type field value
func (o *FirewallRule) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FirewallRule) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FirewallRule) SetType(v string) {
	o.Type = v
}

// GetData returns the Data field value
func (o *FirewallRule) GetData() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *FirewallRule) GetDataOk() (map[string][]string, bool) {
	if o == nil {
		return map[string][]string{}, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *FirewallRule) SetData(v map[string][]string) {
	o.Data = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FirewallRule) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallRule) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FirewallRule) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FirewallRule) SetDescription(v string) {
	o.Description = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *FirewallRule) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallRule) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *FirewallRule) HasDisabled() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *FirewallRule) SetDisabled(v bool) {
	o.Disabled = &v
}

func (o FirewallRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FirewallRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	toSerialize["type"] = o.Type
	toSerialize["data"] = o.Data
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FirewallRule) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
		"type",
		"data",
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

	varFirewallRule := _FirewallRule{}

	err = json.Unmarshal(data, &varFirewallRule)

	if err != nil {
		return err
	}

	*o = FirewallRule(varFirewallRule)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "type")
		delete(additionalProperties, "data")
		delete(additionalProperties, "description")
		delete(additionalProperties, "disabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFirewallRule struct {
	value *FirewallRule
	isSet bool
}

func (v NullableFirewallRule) Get() *FirewallRule {
	return v.value
}

func (v *NullableFirewallRule) Set(val *FirewallRule) {
	v.value = val
	v.isSet = true
}

func (v NullableFirewallRule) IsSet() bool {
	return v.isSet
}

func (v *NullableFirewallRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirewallRule(val *FirewallRule) *NullableFirewallRule {
	return &NullableFirewallRule{value: val, isSet: true}
}

func (v NullableFirewallRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirewallRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


