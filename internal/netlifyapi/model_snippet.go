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

// checks if the Snippet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Snippet{}

// Snippet struct for Snippet
type Snippet struct {
	// The ID of the snippet
	Id int64 `json:"id"`
	// The title of the snippet
	Title string `json:"title"`
	// The general snippet content
	General string `json:"general"`
	// The position to inject the snippet. Uses footer if not specified
	GeneralPosition string `json:"general_position"`
	// The goal snippet content
	Goal string `json:"goal"`
	// The position to inject the snippet. Uses footer if not specified
	GoalPosition string `json:"goal_position"`
	AdditionalProperties map[string]interface{}
}

type _Snippet Snippet

// NewSnippet instantiates a new Snippet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnippet(id int64, title string, general string, generalPosition string, goal string, goalPosition string) *Snippet {
	this := Snippet{}
	this.Id = id
	this.Title = title
	this.General = general
	this.GeneralPosition = generalPosition
	this.Goal = goal
	this.GoalPosition = goalPosition
	return &this
}

// NewSnippetWithDefaults instantiates a new Snippet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnippetWithDefaults() *Snippet {
	this := Snippet{}
	return &this
}

// GetId returns the Id field value
func (o *Snippet) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Snippet) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Snippet) SetId(v int64) {
	o.Id = v
}

// GetTitle returns the Title field value
func (o *Snippet) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *Snippet) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *Snippet) SetTitle(v string) {
	o.Title = v
}

// GetGeneral returns the General field value
func (o *Snippet) GetGeneral() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.General
}

// GetGeneralOk returns a tuple with the General field value
// and a boolean to check if the value has been set.
func (o *Snippet) GetGeneralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.General, true
}

// SetGeneral sets field value
func (o *Snippet) SetGeneral(v string) {
	o.General = v
}

// GetGeneralPosition returns the GeneralPosition field value
func (o *Snippet) GetGeneralPosition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GeneralPosition
}

// GetGeneralPositionOk returns a tuple with the GeneralPosition field value
// and a boolean to check if the value has been set.
func (o *Snippet) GetGeneralPositionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GeneralPosition, true
}

// SetGeneralPosition sets field value
func (o *Snippet) SetGeneralPosition(v string) {
	o.GeneralPosition = v
}

// GetGoal returns the Goal field value
func (o *Snippet) GetGoal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Goal
}

// GetGoalOk returns a tuple with the Goal field value
// and a boolean to check if the value has been set.
func (o *Snippet) GetGoalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Goal, true
}

// SetGoal sets field value
func (o *Snippet) SetGoal(v string) {
	o.Goal = v
}

// GetGoalPosition returns the GoalPosition field value
func (o *Snippet) GetGoalPosition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GoalPosition
}

// GetGoalPositionOk returns a tuple with the GoalPosition field value
// and a boolean to check if the value has been set.
func (o *Snippet) GetGoalPositionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GoalPosition, true
}

// SetGoalPosition sets field value
func (o *Snippet) SetGoalPosition(v string) {
	o.GoalPosition = v
}

func (o Snippet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Snippet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["title"] = o.Title
	toSerialize["general"] = o.General
	toSerialize["general_position"] = o.GeneralPosition
	toSerialize["goal"] = o.Goal
	toSerialize["goal_position"] = o.GoalPosition

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Snippet) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"title",
		"general",
		"general_position",
		"goal",
		"goal_position",
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

	varSnippet := _Snippet{}

	err = json.Unmarshal(data, &varSnippet)

	if err != nil {
		return err
	}

	*o = Snippet(varSnippet)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "title")
		delete(additionalProperties, "general")
		delete(additionalProperties, "general_position")
		delete(additionalProperties, "goal")
		delete(additionalProperties, "goal_position")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSnippet struct {
	value *Snippet
	isSet bool
}

func (v NullableSnippet) Get() *Snippet {
	return v.value
}

func (v *NullableSnippet) Set(val *Snippet) {
	v.value = val
	v.isSet = true
}

func (v NullableSnippet) IsSet() bool {
	return v.isSet
}

func (v *NullableSnippet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnippet(val *Snippet) *NullableSnippet {
	return &NullableSnippet{value: val, isSet: true}
}

func (v NullableSnippet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnippet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


