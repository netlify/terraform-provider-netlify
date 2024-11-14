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

// checks if the ManagedWafRuleSetDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagedWafRuleSetDefinition{}

// ManagedWafRuleSetDefinition struct for ManagedWafRuleSetDefinition
type ManagedWafRuleSetDefinition struct {
	Id *string `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ManagedWafRuleSetDefinition ManagedWafRuleSetDefinition

// NewManagedWafRuleSetDefinition instantiates a new ManagedWafRuleSetDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedWafRuleSetDefinition() *ManagedWafRuleSetDefinition {
	this := ManagedWafRuleSetDefinition{}
	return &this
}

// NewManagedWafRuleSetDefinitionWithDefaults instantiates a new ManagedWafRuleSetDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedWafRuleSetDefinitionWithDefaults() *ManagedWafRuleSetDefinition {
	this := ManagedWafRuleSetDefinition{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ManagedWafRuleSetDefinition) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedWafRuleSetDefinition) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ManagedWafRuleSetDefinition) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ManagedWafRuleSetDefinition) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ManagedWafRuleSetDefinition) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedWafRuleSetDefinition) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ManagedWafRuleSetDefinition) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ManagedWafRuleSetDefinition) SetType(v string) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ManagedWafRuleSetDefinition) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedWafRuleSetDefinition) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ManagedWafRuleSetDefinition) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ManagedWafRuleSetDefinition) SetVersion(v string) {
	o.Version = &v
}

func (o ManagedWafRuleSetDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagedWafRuleSetDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ManagedWafRuleSetDefinition) UnmarshalJSON(data []byte) (err error) {
	varManagedWafRuleSetDefinition := _ManagedWafRuleSetDefinition{}

	err = json.Unmarshal(data, &varManagedWafRuleSetDefinition)

	if err != nil {
		return err
	}

	*o = ManagedWafRuleSetDefinition(varManagedWafRuleSetDefinition)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableManagedWafRuleSetDefinition struct {
	value *ManagedWafRuleSetDefinition
	isSet bool
}

func (v NullableManagedWafRuleSetDefinition) Get() *ManagedWafRuleSetDefinition {
	return v.value
}

func (v *NullableManagedWafRuleSetDefinition) Set(val *ManagedWafRuleSetDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedWafRuleSetDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedWafRuleSetDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedWafRuleSetDefinition(val *ManagedWafRuleSetDefinition) *NullableManagedWafRuleSetDefinition {
	return &NullableManagedWafRuleSetDefinition{value: val, isSet: true}
}

func (v NullableManagedWafRuleSetDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedWafRuleSetDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


