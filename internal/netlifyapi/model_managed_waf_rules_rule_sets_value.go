/*
Netlify's API documentation

Netlify is a hosting service for the programmable web. It understands your documents and provides an API to handle atomic deploys of websites, manage form submissions, inject JavaScript snippets, and much more. This is a REST-style API that uses JSON for serialization and OAuth 2 for authentication.   This document is an OpenAPI reference for the Netlify API that you can explore. For more detailed instructions for common uses, please visit the [online documentation](https://docs.netlify.com/api/get-started/). Visit our Community forum to join the conversation about [understanding and using Netlify’s API](https://community.netlify.com/t/common-issue-understanding-and-using-netlifys-api/160).   Additionally, we have two API clients for your convenience: - [Go Client](https://github.com/netlify/open-api#go-client) - [JS Client](https://github.com/netlify/js-client) 

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netlifyapi

import (
	"encoding/json"
)

// checks if the ManagedWafRulesRuleSetsValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagedWafRulesRuleSetsValue{}

// ManagedWafRulesRuleSetsValue struct for ManagedWafRulesRuleSetsValue
type ManagedWafRulesRuleSetsValue struct {
	Definition *ManagedWafRulesRuleSetsValueDefinition `json:"definition,omitempty"`
	Rules []ManagedWafRulesRuleSetsValueRulesInner `json:"rules,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ManagedWafRulesRuleSetsValue ManagedWafRulesRuleSetsValue

// NewManagedWafRulesRuleSetsValue instantiates a new ManagedWafRulesRuleSetsValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedWafRulesRuleSetsValue() *ManagedWafRulesRuleSetsValue {
	this := ManagedWafRulesRuleSetsValue{}
	return &this
}

// NewManagedWafRulesRuleSetsValueWithDefaults instantiates a new ManagedWafRulesRuleSetsValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedWafRulesRuleSetsValueWithDefaults() *ManagedWafRulesRuleSetsValue {
	this := ManagedWafRulesRuleSetsValue{}
	return &this
}

// GetDefinition returns the Definition field value if set, zero value otherwise.
func (o *ManagedWafRulesRuleSetsValue) GetDefinition() ManagedWafRulesRuleSetsValueDefinition {
	if o == nil || IsNil(o.Definition) {
		var ret ManagedWafRulesRuleSetsValueDefinition
		return ret
	}
	return *o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedWafRulesRuleSetsValue) GetDefinitionOk() (*ManagedWafRulesRuleSetsValueDefinition, bool) {
	if o == nil || IsNil(o.Definition) {
		return nil, false
	}
	return o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *ManagedWafRulesRuleSetsValue) HasDefinition() bool {
	if o != nil && !IsNil(o.Definition) {
		return true
	}

	return false
}

// SetDefinition gets a reference to the given ManagedWafRulesRuleSetsValueDefinition and assigns it to the Definition field.
func (o *ManagedWafRulesRuleSetsValue) SetDefinition(v ManagedWafRulesRuleSetsValueDefinition) {
	o.Definition = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *ManagedWafRulesRuleSetsValue) GetRules() []ManagedWafRulesRuleSetsValueRulesInner {
	if o == nil || IsNil(o.Rules) {
		var ret []ManagedWafRulesRuleSetsValueRulesInner
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedWafRulesRuleSetsValue) GetRulesOk() ([]ManagedWafRulesRuleSetsValueRulesInner, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *ManagedWafRulesRuleSetsValue) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []ManagedWafRulesRuleSetsValueRulesInner and assigns it to the Rules field.
func (o *ManagedWafRulesRuleSetsValue) SetRules(v []ManagedWafRulesRuleSetsValueRulesInner) {
	o.Rules = v
}

func (o ManagedWafRulesRuleSetsValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagedWafRulesRuleSetsValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Definition) {
		toSerialize["definition"] = o.Definition
	}
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ManagedWafRulesRuleSetsValue) UnmarshalJSON(data []byte) (err error) {
	varManagedWafRulesRuleSetsValue := _ManagedWafRulesRuleSetsValue{}

	err = json.Unmarshal(data, &varManagedWafRulesRuleSetsValue)

	if err != nil {
		return err
	}

	*o = ManagedWafRulesRuleSetsValue(varManagedWafRulesRuleSetsValue)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "definition")
		delete(additionalProperties, "rules")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableManagedWafRulesRuleSetsValue struct {
	value *ManagedWafRulesRuleSetsValue
	isSet bool
}

func (v NullableManagedWafRulesRuleSetsValue) Get() *ManagedWafRulesRuleSetsValue {
	return v.value
}

func (v *NullableManagedWafRulesRuleSetsValue) Set(val *ManagedWafRulesRuleSetsValue) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedWafRulesRuleSetsValue) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedWafRulesRuleSetsValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedWafRulesRuleSetsValue(val *ManagedWafRulesRuleSetsValue) *NullableManagedWafRulesRuleSetsValue {
	return &NullableManagedWafRulesRuleSetsValue{value: val, isSet: true}
}

func (v NullableManagedWafRulesRuleSetsValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedWafRulesRuleSetsValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


