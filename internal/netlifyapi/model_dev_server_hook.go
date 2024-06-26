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

// checks if the DevServerHook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DevServerHook{}

// DevServerHook struct for DevServerHook
type DevServerHook struct {
	// The ID of the dev server hook
	Id string `json:"id"`
	// The site ID of the dev server hook
	SiteId string `json:"site_id"`
	// The title of the dev server hook
	Title string `json:"title"`
	// The branch of the dev server hook
	Branch string `json:"branch"`
	// The URL of the dev server hook
	Url string `json:"url"`
	// The message of the dev server hook
	Msg string `json:"msg"`
	// The type of the dev server hook
	Type string `json:"type"`
	// When the dev server hook was created
	CreatedAt time.Time `json:"created_at"`
	AdditionalProperties map[string]interface{}
}

type _DevServerHook DevServerHook

// NewDevServerHook instantiates a new DevServerHook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevServerHook(id string, siteId string, title string, branch string, url string, msg string, type_ string, createdAt time.Time) *DevServerHook {
	this := DevServerHook{}
	this.Id = id
	this.SiteId = siteId
	this.Title = title
	this.Branch = branch
	this.Url = url
	this.Msg = msg
	this.Type = type_
	this.CreatedAt = createdAt
	return &this
}

// NewDevServerHookWithDefaults instantiates a new DevServerHook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevServerHookWithDefaults() *DevServerHook {
	this := DevServerHook{}
	return &this
}

// GetId returns the Id field value
func (o *DevServerHook) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DevServerHook) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DevServerHook) SetId(v string) {
	o.Id = v
}

// GetSiteId returns the SiteId field value
func (o *DevServerHook) GetSiteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value
// and a boolean to check if the value has been set.
func (o *DevServerHook) GetSiteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteId, true
}

// SetSiteId sets field value
func (o *DevServerHook) SetSiteId(v string) {
	o.SiteId = v
}

// GetTitle returns the Title field value
func (o *DevServerHook) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *DevServerHook) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *DevServerHook) SetTitle(v string) {
	o.Title = v
}

// GetBranch returns the Branch field value
func (o *DevServerHook) GetBranch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Branch
}

// GetBranchOk returns a tuple with the Branch field value
// and a boolean to check if the value has been set.
func (o *DevServerHook) GetBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Branch, true
}

// SetBranch sets field value
func (o *DevServerHook) SetBranch(v string) {
	o.Branch = v
}

// GetUrl returns the Url field value
func (o *DevServerHook) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *DevServerHook) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *DevServerHook) SetUrl(v string) {
	o.Url = v
}

// GetMsg returns the Msg field value
func (o *DevServerHook) GetMsg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Msg
}

// GetMsgOk returns a tuple with the Msg field value
// and a boolean to check if the value has been set.
func (o *DevServerHook) GetMsgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Msg, true
}

// SetMsg sets field value
func (o *DevServerHook) SetMsg(v string) {
	o.Msg = v
}

// GetType returns the Type field value
func (o *DevServerHook) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DevServerHook) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DevServerHook) SetType(v string) {
	o.Type = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *DevServerHook) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DevServerHook) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *DevServerHook) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o DevServerHook) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DevServerHook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["site_id"] = o.SiteId
	toSerialize["title"] = o.Title
	toSerialize["branch"] = o.Branch
	toSerialize["url"] = o.Url
	toSerialize["msg"] = o.Msg
	toSerialize["type"] = o.Type
	toSerialize["created_at"] = o.CreatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DevServerHook) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"site_id",
		"title",
		"branch",
		"url",
		"msg",
		"type",
		"created_at",
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

	varDevServerHook := _DevServerHook{}

	err = json.Unmarshal(data, &varDevServerHook)

	if err != nil {
		return err
	}

	*o = DevServerHook(varDevServerHook)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "site_id")
		delete(additionalProperties, "title")
		delete(additionalProperties, "branch")
		delete(additionalProperties, "url")
		delete(additionalProperties, "msg")
		delete(additionalProperties, "type")
		delete(additionalProperties, "created_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDevServerHook struct {
	value *DevServerHook
	isSet bool
}

func (v NullableDevServerHook) Get() *DevServerHook {
	return v.value
}

func (v *NullableDevServerHook) Set(val *DevServerHook) {
	v.value = val
	v.isSet = true
}

func (v NullableDevServerHook) IsSet() bool {
	return v.isSet
}

func (v *NullableDevServerHook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevServerHook(val *DevServerHook) *NullableDevServerHook {
	return &NullableDevServerHook{value: val, isSet: true}
}

func (v NullableDevServerHook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevServerHook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


