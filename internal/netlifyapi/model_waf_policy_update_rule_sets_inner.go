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

// checks if the WafPolicyUpdateRuleSetsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WafPolicyUpdateRuleSetsInner{}

// WafPolicyUpdateRuleSetsInner struct for WafPolicyUpdateRuleSetsInner
type WafPolicyUpdateRuleSetsInner struct {
	ManagedId *string `json:"managed_id,omitempty"`
	ExcludedPatterns []string `json:"excluded_patterns,omitempty"`
	PassiveMode *bool `json:"passive_mode,omitempty"`
	OverallThreshold *int64 `json:"overall_threshold,omitempty"`
	CategoryThresholds map[string]int64 `json:"category_thresholds,omitempty"`
	RuleOverrides map[string]interface{} `json:"rule_overrides,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WafPolicyUpdateRuleSetsInner WafPolicyUpdateRuleSetsInner

// NewWafPolicyUpdateRuleSetsInner instantiates a new WafPolicyUpdateRuleSetsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafPolicyUpdateRuleSetsInner() *WafPolicyUpdateRuleSetsInner {
	this := WafPolicyUpdateRuleSetsInner{}
	return &this
}

// NewWafPolicyUpdateRuleSetsInnerWithDefaults instantiates a new WafPolicyUpdateRuleSetsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafPolicyUpdateRuleSetsInnerWithDefaults() *WafPolicyUpdateRuleSetsInner {
	this := WafPolicyUpdateRuleSetsInner{}
	return &this
}

// GetManagedId returns the ManagedId field value if set, zero value otherwise.
func (o *WafPolicyUpdateRuleSetsInner) GetManagedId() string {
	if o == nil || IsNil(o.ManagedId) {
		var ret string
		return ret
	}
	return *o.ManagedId
}

// GetManagedIdOk returns a tuple with the ManagedId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafPolicyUpdateRuleSetsInner) GetManagedIdOk() (*string, bool) {
	if o == nil || IsNil(o.ManagedId) {
		return nil, false
	}
	return o.ManagedId, true
}

// HasManagedId returns a boolean if a field has been set.
func (o *WafPolicyUpdateRuleSetsInner) HasManagedId() bool {
	if o != nil && !IsNil(o.ManagedId) {
		return true
	}

	return false
}

// SetManagedId gets a reference to the given string and assigns it to the ManagedId field.
func (o *WafPolicyUpdateRuleSetsInner) SetManagedId(v string) {
	o.ManagedId = &v
}

// GetExcludedPatterns returns the ExcludedPatterns field value if set, zero value otherwise.
func (o *WafPolicyUpdateRuleSetsInner) GetExcludedPatterns() []string {
	if o == nil || IsNil(o.ExcludedPatterns) {
		var ret []string
		return ret
	}
	return o.ExcludedPatterns
}

// GetExcludedPatternsOk returns a tuple with the ExcludedPatterns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafPolicyUpdateRuleSetsInner) GetExcludedPatternsOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludedPatterns) {
		return nil, false
	}
	return o.ExcludedPatterns, true
}

// HasExcludedPatterns returns a boolean if a field has been set.
func (o *WafPolicyUpdateRuleSetsInner) HasExcludedPatterns() bool {
	if o != nil && !IsNil(o.ExcludedPatterns) {
		return true
	}

	return false
}

// SetExcludedPatterns gets a reference to the given []string and assigns it to the ExcludedPatterns field.
func (o *WafPolicyUpdateRuleSetsInner) SetExcludedPatterns(v []string) {
	o.ExcludedPatterns = v
}

// GetPassiveMode returns the PassiveMode field value if set, zero value otherwise.
func (o *WafPolicyUpdateRuleSetsInner) GetPassiveMode() bool {
	if o == nil || IsNil(o.PassiveMode) {
		var ret bool
		return ret
	}
	return *o.PassiveMode
}

// GetPassiveModeOk returns a tuple with the PassiveMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafPolicyUpdateRuleSetsInner) GetPassiveModeOk() (*bool, bool) {
	if o == nil || IsNil(o.PassiveMode) {
		return nil, false
	}
	return o.PassiveMode, true
}

// HasPassiveMode returns a boolean if a field has been set.
func (o *WafPolicyUpdateRuleSetsInner) HasPassiveMode() bool {
	if o != nil && !IsNil(o.PassiveMode) {
		return true
	}

	return false
}

// SetPassiveMode gets a reference to the given bool and assigns it to the PassiveMode field.
func (o *WafPolicyUpdateRuleSetsInner) SetPassiveMode(v bool) {
	o.PassiveMode = &v
}

// GetOverallThreshold returns the OverallThreshold field value if set, zero value otherwise.
func (o *WafPolicyUpdateRuleSetsInner) GetOverallThreshold() int64 {
	if o == nil || IsNil(o.OverallThreshold) {
		var ret int64
		return ret
	}
	return *o.OverallThreshold
}

// GetOverallThresholdOk returns a tuple with the OverallThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafPolicyUpdateRuleSetsInner) GetOverallThresholdOk() (*int64, bool) {
	if o == nil || IsNil(o.OverallThreshold) {
		return nil, false
	}
	return o.OverallThreshold, true
}

// HasOverallThreshold returns a boolean if a field has been set.
func (o *WafPolicyUpdateRuleSetsInner) HasOverallThreshold() bool {
	if o != nil && !IsNil(o.OverallThreshold) {
		return true
	}

	return false
}

// SetOverallThreshold gets a reference to the given int64 and assigns it to the OverallThreshold field.
func (o *WafPolicyUpdateRuleSetsInner) SetOverallThreshold(v int64) {
	o.OverallThreshold = &v
}

// GetCategoryThresholds returns the CategoryThresholds field value if set, zero value otherwise.
func (o *WafPolicyUpdateRuleSetsInner) GetCategoryThresholds() map[string]int64 {
	if o == nil || IsNil(o.CategoryThresholds) {
		var ret map[string]int64
		return ret
	}
	return o.CategoryThresholds
}

// GetCategoryThresholdsOk returns a tuple with the CategoryThresholds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafPolicyUpdateRuleSetsInner) GetCategoryThresholdsOk() (map[string]int64, bool) {
	if o == nil || IsNil(o.CategoryThresholds) {
		return map[string]int64{}, false
	}
	return o.CategoryThresholds, true
}

// HasCategoryThresholds returns a boolean if a field has been set.
func (o *WafPolicyUpdateRuleSetsInner) HasCategoryThresholds() bool {
	if o != nil && !IsNil(o.CategoryThresholds) {
		return true
	}

	return false
}

// SetCategoryThresholds gets a reference to the given map[string]int64 and assigns it to the CategoryThresholds field.
func (o *WafPolicyUpdateRuleSetsInner) SetCategoryThresholds(v map[string]int64) {
	o.CategoryThresholds = v
}

// GetRuleOverrides returns the RuleOverrides field value if set, zero value otherwise.
func (o *WafPolicyUpdateRuleSetsInner) GetRuleOverrides() map[string]interface{} {
	if o == nil || IsNil(o.RuleOverrides) {
		var ret map[string]interface{}
		return ret
	}
	return o.RuleOverrides
}

// GetRuleOverridesOk returns a tuple with the RuleOverrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafPolicyUpdateRuleSetsInner) GetRuleOverridesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.RuleOverrides) {
		return map[string]interface{}{}, false
	}
	return o.RuleOverrides, true
}

// HasRuleOverrides returns a boolean if a field has been set.
func (o *WafPolicyUpdateRuleSetsInner) HasRuleOverrides() bool {
	if o != nil && !IsNil(o.RuleOverrides) {
		return true
	}

	return false
}

// SetRuleOverrides gets a reference to the given map[string]interface{} and assigns it to the RuleOverrides field.
func (o *WafPolicyUpdateRuleSetsInner) SetRuleOverrides(v map[string]interface{}) {
	o.RuleOverrides = v
}

func (o WafPolicyUpdateRuleSetsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WafPolicyUpdateRuleSetsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ManagedId) {
		toSerialize["managed_id"] = o.ManagedId
	}
	if !IsNil(o.ExcludedPatterns) {
		toSerialize["excluded_patterns"] = o.ExcludedPatterns
	}
	if !IsNil(o.PassiveMode) {
		toSerialize["passive_mode"] = o.PassiveMode
	}
	if !IsNil(o.OverallThreshold) {
		toSerialize["overall_threshold"] = o.OverallThreshold
	}
	if !IsNil(o.CategoryThresholds) {
		toSerialize["category_thresholds"] = o.CategoryThresholds
	}
	if !IsNil(o.RuleOverrides) {
		toSerialize["rule_overrides"] = o.RuleOverrides
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WafPolicyUpdateRuleSetsInner) UnmarshalJSON(data []byte) (err error) {
	varWafPolicyUpdateRuleSetsInner := _WafPolicyUpdateRuleSetsInner{}

	err = json.Unmarshal(data, &varWafPolicyUpdateRuleSetsInner)

	if err != nil {
		return err
	}

	*o = WafPolicyUpdateRuleSetsInner(varWafPolicyUpdateRuleSetsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "managed_id")
		delete(additionalProperties, "excluded_patterns")
		delete(additionalProperties, "passive_mode")
		delete(additionalProperties, "overall_threshold")
		delete(additionalProperties, "category_thresholds")
		delete(additionalProperties, "rule_overrides")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWafPolicyUpdateRuleSetsInner struct {
	value *WafPolicyUpdateRuleSetsInner
	isSet bool
}

func (v NullableWafPolicyUpdateRuleSetsInner) Get() *WafPolicyUpdateRuleSetsInner {
	return v.value
}

func (v *NullableWafPolicyUpdateRuleSetsInner) Set(val *WafPolicyUpdateRuleSetsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableWafPolicyUpdateRuleSetsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableWafPolicyUpdateRuleSetsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWafPolicyUpdateRuleSetsInner(val *WafPolicyUpdateRuleSetsInner) *NullableWafPolicyUpdateRuleSetsInner {
	return &NullableWafPolicyUpdateRuleSetsInner{value: val, isSet: true}
}

func (v NullableWafPolicyUpdateRuleSetsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWafPolicyUpdateRuleSetsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


