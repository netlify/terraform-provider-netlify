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

// checks if the UpdateSiteDeployRequestFunctionsConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSiteDeployRequestFunctionsConfig{}

// UpdateSiteDeployRequestFunctionsConfig struct for UpdateSiteDeployRequestFunctionsConfig
type UpdateSiteDeployRequestFunctionsConfig struct {
	Name *UpdateSiteDeployRequestFunctionsConfigName `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateSiteDeployRequestFunctionsConfig UpdateSiteDeployRequestFunctionsConfig

// NewUpdateSiteDeployRequestFunctionsConfig instantiates a new UpdateSiteDeployRequestFunctionsConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSiteDeployRequestFunctionsConfig() *UpdateSiteDeployRequestFunctionsConfig {
	this := UpdateSiteDeployRequestFunctionsConfig{}
	return &this
}

// NewUpdateSiteDeployRequestFunctionsConfigWithDefaults instantiates a new UpdateSiteDeployRequestFunctionsConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSiteDeployRequestFunctionsConfigWithDefaults() *UpdateSiteDeployRequestFunctionsConfig {
	this := UpdateSiteDeployRequestFunctionsConfig{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateSiteDeployRequestFunctionsConfig) GetName() UpdateSiteDeployRequestFunctionsConfigName {
	if o == nil || IsNil(o.Name) {
		var ret UpdateSiteDeployRequestFunctionsConfigName
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSiteDeployRequestFunctionsConfig) GetNameOk() (*UpdateSiteDeployRequestFunctionsConfigName, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateSiteDeployRequestFunctionsConfig) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given UpdateSiteDeployRequestFunctionsConfigName and assigns it to the Name field.
func (o *UpdateSiteDeployRequestFunctionsConfig) SetName(v UpdateSiteDeployRequestFunctionsConfigName) {
	o.Name = &v
}

func (o UpdateSiteDeployRequestFunctionsConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSiteDeployRequestFunctionsConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateSiteDeployRequestFunctionsConfig) UnmarshalJSON(data []byte) (err error) {
	varUpdateSiteDeployRequestFunctionsConfig := _UpdateSiteDeployRequestFunctionsConfig{}

	err = json.Unmarshal(data, &varUpdateSiteDeployRequestFunctionsConfig)

	if err != nil {
		return err
	}

	*o = UpdateSiteDeployRequestFunctionsConfig(varUpdateSiteDeployRequestFunctionsConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateSiteDeployRequestFunctionsConfig struct {
	value *UpdateSiteDeployRequestFunctionsConfig
	isSet bool
}

func (v NullableUpdateSiteDeployRequestFunctionsConfig) Get() *UpdateSiteDeployRequestFunctionsConfig {
	return v.value
}

func (v *NullableUpdateSiteDeployRequestFunctionsConfig) Set(val *UpdateSiteDeployRequestFunctionsConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSiteDeployRequestFunctionsConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSiteDeployRequestFunctionsConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSiteDeployRequestFunctionsConfig(val *UpdateSiteDeployRequestFunctionsConfig) *NullableUpdateSiteDeployRequestFunctionsConfig {
	return &NullableUpdateSiteDeployRequestFunctionsConfig{value: val, isSet: true}
}

func (v NullableUpdateSiteDeployRequestFunctionsConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSiteDeployRequestFunctionsConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

