/*
Netlify's API documentation

Netlify is a hosting service for the programmable web. It understands your documents and provides an API to handle atomic deploys of websites, manage form submissions, inject JavaScript snippets, and much more. This is a REST-style API that uses JSON for serialization and OAuth 2 for authentication.   This document is an OpenAPI reference for the Netlify API that you can explore. For more detailed instructions for common uses, please visit the [online documentation](https://docs.netlify.com/api/get-started/). Visit our Community forum to join the conversation about [understanding and using Netlify’s API](https://community.netlify.com/t/common-issue-understanding-and-using-netlifys-api/160).   Additionally, we have two API clients for your convenience: - [Go Client](https://github.com/netlify/open-api#go-client) - [JS Client](https://github.com/netlify/js-client) 

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netlifyapi

import (
	"encoding/json"
	"fmt"
)

// checks if the CDPTicketDataTeam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CDPTicketDataTeam{}

// CDPTicketDataTeam struct for CDPTicketDataTeam
type CDPTicketDataTeam struct {
	Archived bool `json:"archived"`
	Color string `json:"color"`
	Id string `json:"id"`
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _CDPTicketDataTeam CDPTicketDataTeam

// NewCDPTicketDataTeam instantiates a new CDPTicketDataTeam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCDPTicketDataTeam(archived bool, color string, id string, name string) *CDPTicketDataTeam {
	this := CDPTicketDataTeam{}
	this.Archived = archived
	this.Color = color
	this.Id = id
	this.Name = name
	return &this
}

// NewCDPTicketDataTeamWithDefaults instantiates a new CDPTicketDataTeam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCDPTicketDataTeamWithDefaults() *CDPTicketDataTeam {
	this := CDPTicketDataTeam{}
	return &this
}

// GetArchived returns the Archived field value
func (o *CDPTicketDataTeam) GetArchived() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value
// and a boolean to check if the value has been set.
func (o *CDPTicketDataTeam) GetArchivedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Archived, true
}

// SetArchived sets field value
func (o *CDPTicketDataTeam) SetArchived(v bool) {
	o.Archived = v
}

// GetColor returns the Color field value
func (o *CDPTicketDataTeam) GetColor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Color
}

// GetColorOk returns a tuple with the Color field value
// and a boolean to check if the value has been set.
func (o *CDPTicketDataTeam) GetColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Color, true
}

// SetColor sets field value
func (o *CDPTicketDataTeam) SetColor(v string) {
	o.Color = v
}

// GetId returns the Id field value
func (o *CDPTicketDataTeam) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CDPTicketDataTeam) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CDPTicketDataTeam) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *CDPTicketDataTeam) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CDPTicketDataTeam) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CDPTicketDataTeam) SetName(v string) {
	o.Name = v
}

func (o CDPTicketDataTeam) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CDPTicketDataTeam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["archived"] = o.Archived
	toSerialize["color"] = o.Color
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CDPTicketDataTeam) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"archived",
		"color",
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

	varCDPTicketDataTeam := _CDPTicketDataTeam{}

	err = json.Unmarshal(data, &varCDPTicketDataTeam)

	if err != nil {
		return err
	}

	*o = CDPTicketDataTeam(varCDPTicketDataTeam)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "archived")
		delete(additionalProperties, "color")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCDPTicketDataTeam struct {
	value *CDPTicketDataTeam
	isSet bool
}

func (v NullableCDPTicketDataTeam) Get() *CDPTicketDataTeam {
	return v.value
}

func (v *NullableCDPTicketDataTeam) Set(val *CDPTicketDataTeam) {
	v.value = val
	v.isSet = true
}

func (v NullableCDPTicketDataTeam) IsSet() bool {
	return v.isSet
}

func (v *NullableCDPTicketDataTeam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCDPTicketDataTeam(val *CDPTicketDataTeam) *NullableCDPTicketDataTeam {
	return &NullableCDPTicketDataTeam{value: val, isSet: true}
}

func (v NullableCDPTicketDataTeam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCDPTicketDataTeam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


