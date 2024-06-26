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

// checks if the BitbucketSelfHostedConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BitbucketSelfHostedConfig{}

// BitbucketSelfHostedConfig struct for BitbucketSelfHostedConfig
type BitbucketSelfHostedConfig struct {
	// the url for the bitbucket instance
	InstanceUrl *string `json:"instance_url,omitempty"`
	// the clone url for the bitbucket instance
	CloneUrl *string `json:"clone_url,omitempty"`
	// the client id of the bitbucket application
	ClientId *string `json:"client_id,omitempty"`
	// the client secret of the bitbucket application
	ClientSecret *string `json:"client_secret,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BitbucketSelfHostedConfig BitbucketSelfHostedConfig

// NewBitbucketSelfHostedConfig instantiates a new BitbucketSelfHostedConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBitbucketSelfHostedConfig() *BitbucketSelfHostedConfig {
	this := BitbucketSelfHostedConfig{}
	return &this
}

// NewBitbucketSelfHostedConfigWithDefaults instantiates a new BitbucketSelfHostedConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBitbucketSelfHostedConfigWithDefaults() *BitbucketSelfHostedConfig {
	this := BitbucketSelfHostedConfig{}
	return &this
}

// GetInstanceUrl returns the InstanceUrl field value if set, zero value otherwise.
func (o *BitbucketSelfHostedConfig) GetInstanceUrl() string {
	if o == nil || IsNil(o.InstanceUrl) {
		var ret string
		return ret
	}
	return *o.InstanceUrl
}

// GetInstanceUrlOk returns a tuple with the InstanceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BitbucketSelfHostedConfig) GetInstanceUrlOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceUrl) {
		return nil, false
	}
	return o.InstanceUrl, true
}

// HasInstanceUrl returns a boolean if a field has been set.
func (o *BitbucketSelfHostedConfig) HasInstanceUrl() bool {
	if o != nil && !IsNil(o.InstanceUrl) {
		return true
	}

	return false
}

// SetInstanceUrl gets a reference to the given string and assigns it to the InstanceUrl field.
func (o *BitbucketSelfHostedConfig) SetInstanceUrl(v string) {
	o.InstanceUrl = &v
}

// GetCloneUrl returns the CloneUrl field value if set, zero value otherwise.
func (o *BitbucketSelfHostedConfig) GetCloneUrl() string {
	if o == nil || IsNil(o.CloneUrl) {
		var ret string
		return ret
	}
	return *o.CloneUrl
}

// GetCloneUrlOk returns a tuple with the CloneUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BitbucketSelfHostedConfig) GetCloneUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CloneUrl) {
		return nil, false
	}
	return o.CloneUrl, true
}

// HasCloneUrl returns a boolean if a field has been set.
func (o *BitbucketSelfHostedConfig) HasCloneUrl() bool {
	if o != nil && !IsNil(o.CloneUrl) {
		return true
	}

	return false
}

// SetCloneUrl gets a reference to the given string and assigns it to the CloneUrl field.
func (o *BitbucketSelfHostedConfig) SetCloneUrl(v string) {
	o.CloneUrl = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *BitbucketSelfHostedConfig) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BitbucketSelfHostedConfig) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *BitbucketSelfHostedConfig) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *BitbucketSelfHostedConfig) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *BitbucketSelfHostedConfig) GetClientSecret() string {
	if o == nil || IsNil(o.ClientSecret) {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BitbucketSelfHostedConfig) GetClientSecretOk() (*string, bool) {
	if o == nil || IsNil(o.ClientSecret) {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *BitbucketSelfHostedConfig) HasClientSecret() bool {
	if o != nil && !IsNil(o.ClientSecret) {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *BitbucketSelfHostedConfig) SetClientSecret(v string) {
	o.ClientSecret = &v
}

func (o BitbucketSelfHostedConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BitbucketSelfHostedConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InstanceUrl) {
		toSerialize["instance_url"] = o.InstanceUrl
	}
	if !IsNil(o.CloneUrl) {
		toSerialize["clone_url"] = o.CloneUrl
	}
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}
	if !IsNil(o.ClientSecret) {
		toSerialize["client_secret"] = o.ClientSecret
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BitbucketSelfHostedConfig) UnmarshalJSON(data []byte) (err error) {
	varBitbucketSelfHostedConfig := _BitbucketSelfHostedConfig{}

	err = json.Unmarshal(data, &varBitbucketSelfHostedConfig)

	if err != nil {
		return err
	}

	*o = BitbucketSelfHostedConfig(varBitbucketSelfHostedConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "instance_url")
		delete(additionalProperties, "clone_url")
		delete(additionalProperties, "client_id")
		delete(additionalProperties, "client_secret")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBitbucketSelfHostedConfig struct {
	value *BitbucketSelfHostedConfig
	isSet bool
}

func (v NullableBitbucketSelfHostedConfig) Get() *BitbucketSelfHostedConfig {
	return v.value
}

func (v *NullableBitbucketSelfHostedConfig) Set(val *BitbucketSelfHostedConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableBitbucketSelfHostedConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableBitbucketSelfHostedConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBitbucketSelfHostedConfig(val *BitbucketSelfHostedConfig) *NullableBitbucketSelfHostedConfig {
	return &NullableBitbucketSelfHostedConfig{value: val, isSet: true}
}

func (v NullableBitbucketSelfHostedConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBitbucketSelfHostedConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


