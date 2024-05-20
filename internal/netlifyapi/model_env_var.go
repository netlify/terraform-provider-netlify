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

// checks if the EnvVar type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvVar{}

// EnvVar Environment variable model definition
type EnvVar struct {
	// The environment variable key, like ALGOLIA_ID (case-sensitive)
	Key string `json:"key"`
	// The scopes that this environment variable is set to
	Scopes []string `json:"scopes"`
	// An array of Value objects containing values and metadata
	Values []EnvVarValue `json:"values"`
	// The timestamp of when the value was last updated
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy EnvVarUser `json:"updated_by"`
	// Should this environment variable be treated as a write-only variable and not accesible outside of Netlify runtimes
	IsSecret *bool `json:"is_secret,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EnvVar EnvVar

// NewEnvVar instantiates a new EnvVar object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvVar(key string, scopes []string, values []EnvVarValue, updatedAt time.Time, updatedBy EnvVarUser) *EnvVar {
	this := EnvVar{}
	this.Key = key
	this.Scopes = scopes
	this.Values = values
	this.UpdatedAt = updatedAt
	this.UpdatedBy = updatedBy
	return &this
}

// NewEnvVarWithDefaults instantiates a new EnvVar object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvVarWithDefaults() *EnvVar {
	this := EnvVar{}
	return &this
}

// GetKey returns the Key field value
func (o *EnvVar) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *EnvVar) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *EnvVar) SetKey(v string) {
	o.Key = v
}

// GetScopes returns the Scopes field value
func (o *EnvVar) GetScopes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *EnvVar) GetScopesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scopes, true
}

// SetScopes sets field value
func (o *EnvVar) SetScopes(v []string) {
	o.Scopes = v
}

// GetValues returns the Values field value
func (o *EnvVar) GetValues() []EnvVarValue {
	if o == nil {
		var ret []EnvVarValue
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *EnvVar) GetValuesOk() ([]EnvVarValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *EnvVar) SetValues(v []EnvVarValue) {
	o.Values = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *EnvVar) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *EnvVar) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *EnvVar) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetUpdatedBy returns the UpdatedBy field value
func (o *EnvVar) GetUpdatedBy() EnvVarUser {
	if o == nil {
		var ret EnvVarUser
		return ret
	}

	return o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value
// and a boolean to check if the value has been set.
func (o *EnvVar) GetUpdatedByOk() (*EnvVarUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedBy, true
}

// SetUpdatedBy sets field value
func (o *EnvVar) SetUpdatedBy(v EnvVarUser) {
	o.UpdatedBy = v
}

// GetIsSecret returns the IsSecret field value if set, zero value otherwise.
func (o *EnvVar) GetIsSecret() bool {
	if o == nil || IsNil(o.IsSecret) {
		var ret bool
		return ret
	}
	return *o.IsSecret
}

// GetIsSecretOk returns a tuple with the IsSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvVar) GetIsSecretOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSecret) {
		return nil, false
	}
	return o.IsSecret, true
}

// HasIsSecret returns a boolean if a field has been set.
func (o *EnvVar) HasIsSecret() bool {
	if o != nil && !IsNil(o.IsSecret) {
		return true
	}

	return false
}

// SetIsSecret gets a reference to the given bool and assigns it to the IsSecret field.
func (o *EnvVar) SetIsSecret(v bool) {
	o.IsSecret = &v
}

func (o EnvVar) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvVar) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["scopes"] = o.Scopes
	toSerialize["values"] = o.Values
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["updated_by"] = o.UpdatedBy
	if !IsNil(o.IsSecret) {
		toSerialize["is_secret"] = o.IsSecret
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EnvVar) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"scopes",
		"values",
		"updated_at",
		"updated_by",
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

	varEnvVar := _EnvVar{}

	err = json.Unmarshal(data, &varEnvVar)

	if err != nil {
		return err
	}

	*o = EnvVar(varEnvVar)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "key")
		delete(additionalProperties, "scopes")
		delete(additionalProperties, "values")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "updated_by")
		delete(additionalProperties, "is_secret")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEnvVar struct {
	value *EnvVar
	isSet bool
}

func (v NullableEnvVar) Get() *EnvVar {
	return v.value
}

func (v *NullableEnvVar) Set(val *EnvVar) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvVar) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvVar) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvVar(val *EnvVar) *NullableEnvVar {
	return &NullableEnvVar{value: val, isSet: true}
}

func (v NullableEnvVar) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvVar) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


