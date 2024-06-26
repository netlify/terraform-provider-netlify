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

// checks if the CDPTicketDataTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CDPTicketDataTemplate{}

// CDPTicketDataTemplate struct for CDPTicketDataTemplate
type CDPTicketDataTemplate struct {
	CreatedAt time.Time `json:"created_at"`
	Fields map[string]interface{} `json:"fields"`
	Id string `json:"id"`
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _CDPTicketDataTemplate CDPTicketDataTemplate

// NewCDPTicketDataTemplate instantiates a new CDPTicketDataTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCDPTicketDataTemplate(createdAt time.Time, fields map[string]interface{}, id string, name string) *CDPTicketDataTemplate {
	this := CDPTicketDataTemplate{}
	this.CreatedAt = createdAt
	this.Fields = fields
	this.Id = id
	this.Name = name
	return &this
}

// NewCDPTicketDataTemplateWithDefaults instantiates a new CDPTicketDataTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCDPTicketDataTemplateWithDefaults() *CDPTicketDataTemplate {
	this := CDPTicketDataTemplate{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *CDPTicketDataTemplate) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CDPTicketDataTemplate) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CDPTicketDataTemplate) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetFields returns the Fields field value
func (o *CDPTicketDataTemplate) GetFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *CDPTicketDataTemplate) GetFieldsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Fields, true
}

// SetFields sets field value
func (o *CDPTicketDataTemplate) SetFields(v map[string]interface{}) {
	o.Fields = v
}

// GetId returns the Id field value
func (o *CDPTicketDataTemplate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CDPTicketDataTemplate) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CDPTicketDataTemplate) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *CDPTicketDataTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CDPTicketDataTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CDPTicketDataTemplate) SetName(v string) {
	o.Name = v
}

func (o CDPTicketDataTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CDPTicketDataTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["fields"] = o.Fields
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CDPTicketDataTemplate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"created_at",
		"fields",
		"id",
		"name",
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

	varCDPTicketDataTemplate := _CDPTicketDataTemplate{}

	err = json.Unmarshal(data, &varCDPTicketDataTemplate)

	if err != nil {
		return err
	}

	*o = CDPTicketDataTemplate(varCDPTicketDataTemplate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "fields")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCDPTicketDataTemplate struct {
	value *CDPTicketDataTemplate
	isSet bool
}

func (v NullableCDPTicketDataTemplate) Get() *CDPTicketDataTemplate {
	return v.value
}

func (v *NullableCDPTicketDataTemplate) Set(val *CDPTicketDataTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableCDPTicketDataTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableCDPTicketDataTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCDPTicketDataTemplate(val *CDPTicketDataTemplate) *NullableCDPTicketDataTemplate {
	return &NullableCDPTicketDataTemplate{value: val, isSet: true}
}

func (v NullableCDPTicketDataTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCDPTicketDataTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


