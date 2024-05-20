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

// checks if the SecurityContact type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityContact{}

// SecurityContact SecurityContact model definition
type SecurityContact struct {
	// Email address of the security contact
	Email string `json:"email"`
	// Role of the security contact
	Role string `json:"role"`
	// When the security contact was created
	CreatedAt time.Time `json:"created_at"`
	// When the security contact was updated
	UpdatedAt time.Time `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _SecurityContact SecurityContact

// NewSecurityContact instantiates a new SecurityContact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityContact(email string, role string, createdAt time.Time, updatedAt time.Time) *SecurityContact {
	this := SecurityContact{}
	this.Email = email
	this.Role = role
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewSecurityContactWithDefaults instantiates a new SecurityContact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityContactWithDefaults() *SecurityContact {
	this := SecurityContact{}
	return &this
}

// GetEmail returns the Email field value
func (o *SecurityContact) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *SecurityContact) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *SecurityContact) SetEmail(v string) {
	o.Email = v
}

// GetRole returns the Role field value
func (o *SecurityContact) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *SecurityContact) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *SecurityContact) SetRole(v string) {
	o.Role = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *SecurityContact) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SecurityContact) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SecurityContact) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *SecurityContact) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *SecurityContact) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *SecurityContact) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o SecurityContact) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityContact) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["role"] = o.Role
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SecurityContact) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"role",
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

	varSecurityContact := _SecurityContact{}

	err = json.Unmarshal(data, &varSecurityContact)

	if err != nil {
		return err
	}

	*o = SecurityContact(varSecurityContact)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		delete(additionalProperties, "role")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSecurityContact struct {
	value *SecurityContact
	isSet bool
}

func (v NullableSecurityContact) Get() *SecurityContact {
	return v.value
}

func (v *NullableSecurityContact) Set(val *SecurityContact) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityContact) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityContact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityContact(val *SecurityContact) *NullableSecurityContact {
	return &NullableSecurityContact{value: val, isSet: true}
}

func (v NullableSecurityContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityContact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


