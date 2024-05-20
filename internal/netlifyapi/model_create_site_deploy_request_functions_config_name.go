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

// checks if the CreateSiteDeployRequestFunctionsConfigName type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSiteDeployRequestFunctionsConfigName{}

// CreateSiteDeployRequestFunctionsConfigName struct for CreateSiteDeployRequestFunctionsConfigName
type CreateSiteDeployRequestFunctionsConfigName struct {
	Priority *int64 `json:"priority,omitempty"`
	BuildData map[string]interface{} `json:"build_data,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Generator *string `json:"generator,omitempty"`
	Routes []CreateSiteDeployRequestFunctionsConfigNameRoutesInner `json:"routes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateSiteDeployRequestFunctionsConfigName CreateSiteDeployRequestFunctionsConfigName

// NewCreateSiteDeployRequestFunctionsConfigName instantiates a new CreateSiteDeployRequestFunctionsConfigName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSiteDeployRequestFunctionsConfigName() *CreateSiteDeployRequestFunctionsConfigName {
	this := CreateSiteDeployRequestFunctionsConfigName{}
	return &this
}

// NewCreateSiteDeployRequestFunctionsConfigNameWithDefaults instantiates a new CreateSiteDeployRequestFunctionsConfigName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSiteDeployRequestFunctionsConfigNameWithDefaults() *CreateSiteDeployRequestFunctionsConfigName {
	this := CreateSiteDeployRequestFunctionsConfigName{}
	return &this
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *CreateSiteDeployRequestFunctionsConfigName) GetPriority() int64 {
	if o == nil || IsNil(o.Priority) {
		var ret int64
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSiteDeployRequestFunctionsConfigName) GetPriorityOk() (*int64, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *CreateSiteDeployRequestFunctionsConfigName) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int64 and assigns it to the Priority field.
func (o *CreateSiteDeployRequestFunctionsConfigName) SetPriority(v int64) {
	o.Priority = &v
}

// GetBuildData returns the BuildData field value if set, zero value otherwise.
func (o *CreateSiteDeployRequestFunctionsConfigName) GetBuildData() map[string]interface{} {
	if o == nil || IsNil(o.BuildData) {
		var ret map[string]interface{}
		return ret
	}
	return o.BuildData
}

// GetBuildDataOk returns a tuple with the BuildData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSiteDeployRequestFunctionsConfigName) GetBuildDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.BuildData) {
		return map[string]interface{}{}, false
	}
	return o.BuildData, true
}

// HasBuildData returns a boolean if a field has been set.
func (o *CreateSiteDeployRequestFunctionsConfigName) HasBuildData() bool {
	if o != nil && !IsNil(o.BuildData) {
		return true
	}

	return false
}

// SetBuildData gets a reference to the given map[string]interface{} and assigns it to the BuildData field.
func (o *CreateSiteDeployRequestFunctionsConfigName) SetBuildData(v map[string]interface{}) {
	o.BuildData = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *CreateSiteDeployRequestFunctionsConfigName) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSiteDeployRequestFunctionsConfigName) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CreateSiteDeployRequestFunctionsConfigName) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *CreateSiteDeployRequestFunctionsConfigName) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetGenerator returns the Generator field value if set, zero value otherwise.
func (o *CreateSiteDeployRequestFunctionsConfigName) GetGenerator() string {
	if o == nil || IsNil(o.Generator) {
		var ret string
		return ret
	}
	return *o.Generator
}

// GetGeneratorOk returns a tuple with the Generator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSiteDeployRequestFunctionsConfigName) GetGeneratorOk() (*string, bool) {
	if o == nil || IsNil(o.Generator) {
		return nil, false
	}
	return o.Generator, true
}

// HasGenerator returns a boolean if a field has been set.
func (o *CreateSiteDeployRequestFunctionsConfigName) HasGenerator() bool {
	if o != nil && !IsNil(o.Generator) {
		return true
	}

	return false
}

// SetGenerator gets a reference to the given string and assigns it to the Generator field.
func (o *CreateSiteDeployRequestFunctionsConfigName) SetGenerator(v string) {
	o.Generator = &v
}

// GetRoutes returns the Routes field value if set, zero value otherwise.
func (o *CreateSiteDeployRequestFunctionsConfigName) GetRoutes() []CreateSiteDeployRequestFunctionsConfigNameRoutesInner {
	if o == nil || IsNil(o.Routes) {
		var ret []CreateSiteDeployRequestFunctionsConfigNameRoutesInner
		return ret
	}
	return o.Routes
}

// GetRoutesOk returns a tuple with the Routes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSiteDeployRequestFunctionsConfigName) GetRoutesOk() ([]CreateSiteDeployRequestFunctionsConfigNameRoutesInner, bool) {
	if o == nil || IsNil(o.Routes) {
		return nil, false
	}
	return o.Routes, true
}

// HasRoutes returns a boolean if a field has been set.
func (o *CreateSiteDeployRequestFunctionsConfigName) HasRoutes() bool {
	if o != nil && !IsNil(o.Routes) {
		return true
	}

	return false
}

// SetRoutes gets a reference to the given []CreateSiteDeployRequestFunctionsConfigNameRoutesInner and assigns it to the Routes field.
func (o *CreateSiteDeployRequestFunctionsConfigName) SetRoutes(v []CreateSiteDeployRequestFunctionsConfigNameRoutesInner) {
	o.Routes = v
}

func (o CreateSiteDeployRequestFunctionsConfigName) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSiteDeployRequestFunctionsConfigName) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.BuildData) {
		toSerialize["build_data"] = o.BuildData
	}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.Generator) {
		toSerialize["generator"] = o.Generator
	}
	if !IsNil(o.Routes) {
		toSerialize["routes"] = o.Routes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateSiteDeployRequestFunctionsConfigName) UnmarshalJSON(data []byte) (err error) {
	varCreateSiteDeployRequestFunctionsConfigName := _CreateSiteDeployRequestFunctionsConfigName{}

	err = json.Unmarshal(data, &varCreateSiteDeployRequestFunctionsConfigName)

	if err != nil {
		return err
	}

	*o = CreateSiteDeployRequestFunctionsConfigName(varCreateSiteDeployRequestFunctionsConfigName)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "priority")
		delete(additionalProperties, "build_data")
		delete(additionalProperties, "display_name")
		delete(additionalProperties, "generator")
		delete(additionalProperties, "routes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateSiteDeployRequestFunctionsConfigName struct {
	value *CreateSiteDeployRequestFunctionsConfigName
	isSet bool
}

func (v NullableCreateSiteDeployRequestFunctionsConfigName) Get() *CreateSiteDeployRequestFunctionsConfigName {
	return v.value
}

func (v *NullableCreateSiteDeployRequestFunctionsConfigName) Set(val *CreateSiteDeployRequestFunctionsConfigName) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSiteDeployRequestFunctionsConfigName) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSiteDeployRequestFunctionsConfigName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSiteDeployRequestFunctionsConfigName(val *CreateSiteDeployRequestFunctionsConfigName) *NullableCreateSiteDeployRequestFunctionsConfigName {
	return &NullableCreateSiteDeployRequestFunctionsConfigName{value: val, isSet: true}
}

func (v NullableCreateSiteDeployRequestFunctionsConfigName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSiteDeployRequestFunctionsConfigName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

