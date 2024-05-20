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

// checks if the GetDirectory200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDirectory200Response{}

// GetDirectory200Response struct for GetDirectory200Response
type GetDirectory200Response struct {
	Name *string `json:"name,omitempty"`
	DirectoryType *string `json:"directory_type,omitempty"`
	ScimConfigured *bool `json:"scim_configured,omitempty"`
	Status *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetDirectory200Response GetDirectory200Response

// NewGetDirectory200Response instantiates a new GetDirectory200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDirectory200Response() *GetDirectory200Response {
	this := GetDirectory200Response{}
	return &this
}

// NewGetDirectory200ResponseWithDefaults instantiates a new GetDirectory200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDirectory200ResponseWithDefaults() *GetDirectory200Response {
	this := GetDirectory200Response{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetDirectory200Response) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDirectory200Response) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetDirectory200Response) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetDirectory200Response) SetName(v string) {
	o.Name = &v
}

// GetDirectoryType returns the DirectoryType field value if set, zero value otherwise.
func (o *GetDirectory200Response) GetDirectoryType() string {
	if o == nil || IsNil(o.DirectoryType) {
		var ret string
		return ret
	}
	return *o.DirectoryType
}

// GetDirectoryTypeOk returns a tuple with the DirectoryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDirectory200Response) GetDirectoryTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DirectoryType) {
		return nil, false
	}
	return o.DirectoryType, true
}

// HasDirectoryType returns a boolean if a field has been set.
func (o *GetDirectory200Response) HasDirectoryType() bool {
	if o != nil && !IsNil(o.DirectoryType) {
		return true
	}

	return false
}

// SetDirectoryType gets a reference to the given string and assigns it to the DirectoryType field.
func (o *GetDirectory200Response) SetDirectoryType(v string) {
	o.DirectoryType = &v
}

// GetScimConfigured returns the ScimConfigured field value if set, zero value otherwise.
func (o *GetDirectory200Response) GetScimConfigured() bool {
	if o == nil || IsNil(o.ScimConfigured) {
		var ret bool
		return ret
	}
	return *o.ScimConfigured
}

// GetScimConfiguredOk returns a tuple with the ScimConfigured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDirectory200Response) GetScimConfiguredOk() (*bool, bool) {
	if o == nil || IsNil(o.ScimConfigured) {
		return nil, false
	}
	return o.ScimConfigured, true
}

// HasScimConfigured returns a boolean if a field has been set.
func (o *GetDirectory200Response) HasScimConfigured() bool {
	if o != nil && !IsNil(o.ScimConfigured) {
		return true
	}

	return false
}

// SetScimConfigured gets a reference to the given bool and assigns it to the ScimConfigured field.
func (o *GetDirectory200Response) SetScimConfigured(v bool) {
	o.ScimConfigured = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetDirectory200Response) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDirectory200Response) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetDirectory200Response) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetDirectory200Response) SetStatus(v string) {
	o.Status = &v
}

func (o GetDirectory200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDirectory200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.DirectoryType) {
		toSerialize["directory_type"] = o.DirectoryType
	}
	if !IsNil(o.ScimConfigured) {
		toSerialize["scim_configured"] = o.ScimConfigured
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetDirectory200Response) UnmarshalJSON(data []byte) (err error) {
	varGetDirectory200Response := _GetDirectory200Response{}

	err = json.Unmarshal(data, &varGetDirectory200Response)

	if err != nil {
		return err
	}

	*o = GetDirectory200Response(varGetDirectory200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "directory_type")
		delete(additionalProperties, "scim_configured")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetDirectory200Response struct {
	value *GetDirectory200Response
	isSet bool
}

func (v NullableGetDirectory200Response) Get() *GetDirectory200Response {
	return v.value
}

func (v *NullableGetDirectory200Response) Set(val *GetDirectory200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDirectory200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDirectory200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDirectory200Response(val *GetDirectory200Response) *NullableGetDirectory200Response {
	return &NullableGetDirectory200Response{value: val, isSet: true}
}

func (v NullableGetDirectory200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDirectory200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


