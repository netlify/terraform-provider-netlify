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

// checks if the DeploySummaryMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploySummaryMessage{}

// DeploySummaryMessage Deploy summary message
type DeploySummaryMessage struct {
	// The description of the message in Markdown format
	Description *string `json:"description,omitempty"`
	// The details of the message in Markdown format
	Details *string `json:"details,omitempty"`
	// The title of the message
	Title *string `json:"title,omitempty"`
	// The type of the message
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeploySummaryMessage DeploySummaryMessage

// NewDeploySummaryMessage instantiates a new DeploySummaryMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploySummaryMessage() *DeploySummaryMessage {
	this := DeploySummaryMessage{}
	return &this
}

// NewDeploySummaryMessageWithDefaults instantiates a new DeploySummaryMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploySummaryMessageWithDefaults() *DeploySummaryMessage {
	this := DeploySummaryMessage{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DeploySummaryMessage) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploySummaryMessage) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DeploySummaryMessage) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DeploySummaryMessage) SetDescription(v string) {
	o.Description = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *DeploySummaryMessage) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploySummaryMessage) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *DeploySummaryMessage) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *DeploySummaryMessage) SetDetails(v string) {
	o.Details = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *DeploySummaryMessage) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploySummaryMessage) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *DeploySummaryMessage) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *DeploySummaryMessage) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DeploySummaryMessage) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploySummaryMessage) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DeploySummaryMessage) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DeploySummaryMessage) SetType(v string) {
	o.Type = &v
}

func (o DeploySummaryMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploySummaryMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeploySummaryMessage) UnmarshalJSON(data []byte) (err error) {
	varDeploySummaryMessage := _DeploySummaryMessage{}

	err = json.Unmarshal(data, &varDeploySummaryMessage)

	if err != nil {
		return err
	}

	*o = DeploySummaryMessage(varDeploySummaryMessage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "details")
		delete(additionalProperties, "title")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeploySummaryMessage struct {
	value *DeploySummaryMessage
	isSet bool
}

func (v NullableDeploySummaryMessage) Get() *DeploySummaryMessage {
	return v.value
}

func (v *NullableDeploySummaryMessage) Set(val *DeploySummaryMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploySummaryMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploySummaryMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploySummaryMessage(val *DeploySummaryMessage) *NullableDeploySummaryMessage {
	return &NullableDeploySummaryMessage{value: val, isSet: true}
}

func (v NullableDeploySummaryMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploySummaryMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


