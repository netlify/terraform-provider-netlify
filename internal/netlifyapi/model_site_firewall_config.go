/*
Netlify's API documentation

Netlify is a hosting service for the programmable web. It understands your documents and provides an API to handle atomic deploys of websites, manage form submissions, inject JavaScript snippets, and much more. This is a REST-style API that uses JSON for serialization and OAuth 2 for authentication.   This document is an OpenAPI reference for the Netlify API that you can explore. For more detailed instructions for common uses, please visit the [online documentation](https://docs.netlify.com/api/get-started/). Visit our Community forum to join the conversation about [understanding and using Netlify’s API](https://community.netlify.com/t/common-issue-understanding-and-using-netlifys-api/160).   Additionally, we have two API clients for your convenience: - [Go Client](https://github.com/netlify/open-api#go-client) - [JS Client](https://github.com/netlify/js-client) 

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netlifyapi

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the SiteFirewallConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteFirewallConfig{}

// SiteFirewallConfig struct for SiteFirewallConfig
type SiteFirewallConfig struct {
	Id string `json:"id"`
	UnpublishedRules FirewallRuleSet `json:"unpublished_rules"`
	PublishedRules FirewallRuleSet `json:"published_rules"`
	// When the deployed branch was created
	CreatedAt time.Time `json:"created_at"`
	// When the deployed branch was updated
	UpdatedAt time.Time `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _SiteFirewallConfig SiteFirewallConfig

// NewSiteFirewallConfig instantiates a new SiteFirewallConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteFirewallConfig(id string, unpublishedRules FirewallRuleSet, publishedRules FirewallRuleSet, createdAt time.Time, updatedAt time.Time) *SiteFirewallConfig {
	this := SiteFirewallConfig{}
	this.Id = id
	this.UnpublishedRules = unpublishedRules
	this.PublishedRules = publishedRules
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewSiteFirewallConfigWithDefaults instantiates a new SiteFirewallConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteFirewallConfigWithDefaults() *SiteFirewallConfig {
	this := SiteFirewallConfig{}
	return &this
}

// GetId returns the Id field value
func (o *SiteFirewallConfig) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SiteFirewallConfig) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SiteFirewallConfig) SetId(v string) {
	o.Id = v
}

// GetUnpublishedRules returns the UnpublishedRules field value
func (o *SiteFirewallConfig) GetUnpublishedRules() FirewallRuleSet {
	if o == nil {
		var ret FirewallRuleSet
		return ret
	}

	return o.UnpublishedRules
}

// GetUnpublishedRulesOk returns a tuple with the UnpublishedRules field value
// and a boolean to check if the value has been set.
func (o *SiteFirewallConfig) GetUnpublishedRulesOk() (*FirewallRuleSet, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnpublishedRules, true
}

// SetUnpublishedRules sets field value
func (o *SiteFirewallConfig) SetUnpublishedRules(v FirewallRuleSet) {
	o.UnpublishedRules = v
}

// GetPublishedRules returns the PublishedRules field value
func (o *SiteFirewallConfig) GetPublishedRules() FirewallRuleSet {
	if o == nil {
		var ret FirewallRuleSet
		return ret
	}

	return o.PublishedRules
}

// GetPublishedRulesOk returns a tuple with the PublishedRules field value
// and a boolean to check if the value has been set.
func (o *SiteFirewallConfig) GetPublishedRulesOk() (*FirewallRuleSet, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublishedRules, true
}

// SetPublishedRules sets field value
func (o *SiteFirewallConfig) SetPublishedRules(v FirewallRuleSet) {
	o.PublishedRules = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *SiteFirewallConfig) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SiteFirewallConfig) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SiteFirewallConfig) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *SiteFirewallConfig) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *SiteFirewallConfig) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *SiteFirewallConfig) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o SiteFirewallConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteFirewallConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["unpublished_rules"] = o.UnpublishedRules
	toSerialize["published_rules"] = o.PublishedRules
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SiteFirewallConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"unpublished_rules",
		"published_rules",
		"created_at",
		"updated_at",
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

	varSiteFirewallConfig := _SiteFirewallConfig{}

	err = json.Unmarshal(data, &varSiteFirewallConfig)

	if err != nil {
		return err
	}

	*o = SiteFirewallConfig(varSiteFirewallConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "unpublished_rules")
		delete(additionalProperties, "published_rules")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSiteFirewallConfig struct {
	value *SiteFirewallConfig
	isSet bool
}

func (v NullableSiteFirewallConfig) Get() *SiteFirewallConfig {
	return v.value
}

func (v *NullableSiteFirewallConfig) Set(val *SiteFirewallConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteFirewallConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteFirewallConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteFirewallConfig(val *SiteFirewallConfig) *NullableSiteFirewallConfig {
	return &NullableSiteFirewallConfig{value: val, isSet: true}
}

func (v NullableSiteFirewallConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteFirewallConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

