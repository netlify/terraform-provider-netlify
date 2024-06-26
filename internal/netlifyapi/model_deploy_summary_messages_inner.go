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

// checks if the DeploySummaryMessagesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploySummaryMessagesInner{}

// DeploySummaryMessagesInner struct for DeploySummaryMessagesInner
type DeploySummaryMessagesInner struct {
	Type *string `json:"type,omitempty"`
	Title *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Details *string `json:"details,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeploySummaryMessagesInner DeploySummaryMessagesInner

// NewDeploySummaryMessagesInner instantiates a new DeploySummaryMessagesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploySummaryMessagesInner() *DeploySummaryMessagesInner {
	this := DeploySummaryMessagesInner{}
	return &this
}

// NewDeploySummaryMessagesInnerWithDefaults instantiates a new DeploySummaryMessagesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploySummaryMessagesInnerWithDefaults() *DeploySummaryMessagesInner {
	this := DeploySummaryMessagesInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DeploySummaryMessagesInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploySummaryMessagesInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DeploySummaryMessagesInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DeploySummaryMessagesInner) SetType(v string) {
	o.Type = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *DeploySummaryMessagesInner) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploySummaryMessagesInner) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *DeploySummaryMessagesInner) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *DeploySummaryMessagesInner) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DeploySummaryMessagesInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploySummaryMessagesInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DeploySummaryMessagesInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DeploySummaryMessagesInner) SetDescription(v string) {
	o.Description = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *DeploySummaryMessagesInner) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploySummaryMessagesInner) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *DeploySummaryMessagesInner) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *DeploySummaryMessagesInner) SetDetails(v string) {
	o.Details = &v
}

func (o DeploySummaryMessagesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploySummaryMessagesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeploySummaryMessagesInner) UnmarshalJSON(data []byte) (err error) {
	varDeploySummaryMessagesInner := _DeploySummaryMessagesInner{}

	err = json.Unmarshal(data, &varDeploySummaryMessagesInner)

	if err != nil {
		return err
	}

	*o = DeploySummaryMessagesInner(varDeploySummaryMessagesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "title")
		delete(additionalProperties, "description")
		delete(additionalProperties, "details")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeploySummaryMessagesInner struct {
	value *DeploySummaryMessagesInner
	isSet bool
}

func (v NullableDeploySummaryMessagesInner) Get() *DeploySummaryMessagesInner {
	return v.value
}

func (v *NullableDeploySummaryMessagesInner) Set(val *DeploySummaryMessagesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploySummaryMessagesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploySummaryMessagesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploySummaryMessagesInner(val *DeploySummaryMessagesInner) *NullableDeploySummaryMessagesInner {
	return &NullableDeploySummaryMessagesInner{value: val, isSet: true}
}

func (v NullableDeploySummaryMessagesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploySummaryMessagesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


