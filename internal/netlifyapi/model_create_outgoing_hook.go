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

// checks if the CreateOutgoingHook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOutgoingHook{}

// CreateOutgoingHook struct for CreateOutgoingHook
type CreateOutgoingHook struct {
	// The site ID of the hook
	SiteId *string `json:"site_id,omitempty"`
	// The form ID of the hook (either form_id or form_name is required if the hook will be created for the form)
	FormId *string `json:"form_id,omitempty"`
	// The form name of the hook
	FormName *string `json:"form_name,omitempty"`
	// The type of the hook
	Type *string `json:"type,omitempty"`
	// The name of the hook event
	Event *string `json:"event,omitempty"`
	// Additional data for the hook
	Data map[string]interface{} `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateOutgoingHook CreateOutgoingHook

// NewCreateOutgoingHook instantiates a new CreateOutgoingHook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOutgoingHook() *CreateOutgoingHook {
	this := CreateOutgoingHook{}
	return &this
}

// NewCreateOutgoingHookWithDefaults instantiates a new CreateOutgoingHook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOutgoingHookWithDefaults() *CreateOutgoingHook {
	this := CreateOutgoingHook{}
	return &this
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *CreateOutgoingHook) GetSiteId() string {
	if o == nil || IsNil(o.SiteId) {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOutgoingHook) GetSiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *CreateOutgoingHook) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *CreateOutgoingHook) SetSiteId(v string) {
	o.SiteId = &v
}

// GetFormId returns the FormId field value if set, zero value otherwise.
func (o *CreateOutgoingHook) GetFormId() string {
	if o == nil || IsNil(o.FormId) {
		var ret string
		return ret
	}
	return *o.FormId
}

// GetFormIdOk returns a tuple with the FormId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOutgoingHook) GetFormIdOk() (*string, bool) {
	if o == nil || IsNil(o.FormId) {
		return nil, false
	}
	return o.FormId, true
}

// HasFormId returns a boolean if a field has been set.
func (o *CreateOutgoingHook) HasFormId() bool {
	if o != nil && !IsNil(o.FormId) {
		return true
	}

	return false
}

// SetFormId gets a reference to the given string and assigns it to the FormId field.
func (o *CreateOutgoingHook) SetFormId(v string) {
	o.FormId = &v
}

// GetFormName returns the FormName field value if set, zero value otherwise.
func (o *CreateOutgoingHook) GetFormName() string {
	if o == nil || IsNil(o.FormName) {
		var ret string
		return ret
	}
	return *o.FormName
}

// GetFormNameOk returns a tuple with the FormName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOutgoingHook) GetFormNameOk() (*string, bool) {
	if o == nil || IsNil(o.FormName) {
		return nil, false
	}
	return o.FormName, true
}

// HasFormName returns a boolean if a field has been set.
func (o *CreateOutgoingHook) HasFormName() bool {
	if o != nil && !IsNil(o.FormName) {
		return true
	}

	return false
}

// SetFormName gets a reference to the given string and assigns it to the FormName field.
func (o *CreateOutgoingHook) SetFormName(v string) {
	o.FormName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateOutgoingHook) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOutgoingHook) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateOutgoingHook) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CreateOutgoingHook) SetType(v string) {
	o.Type = &v
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *CreateOutgoingHook) GetEvent() string {
	if o == nil || IsNil(o.Event) {
		var ret string
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOutgoingHook) GetEventOk() (*string, bool) {
	if o == nil || IsNil(o.Event) {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *CreateOutgoingHook) HasEvent() bool {
	if o != nil && !IsNil(o.Event) {
		return true
	}

	return false
}

// SetEvent gets a reference to the given string and assigns it to the Event field.
func (o *CreateOutgoingHook) SetEvent(v string) {
	o.Event = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateOutgoingHook) GetData() map[string]interface{} {
	if o == nil || IsNil(o.Data) {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOutgoingHook) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateOutgoingHook) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *CreateOutgoingHook) SetData(v map[string]interface{}) {
	o.Data = v
}

func (o CreateOutgoingHook) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOutgoingHook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SiteId) {
		toSerialize["site_id"] = o.SiteId
	}
	if !IsNil(o.FormId) {
		toSerialize["form_id"] = o.FormId
	}
	if !IsNil(o.FormName) {
		toSerialize["form_name"] = o.FormName
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Event) {
		toSerialize["event"] = o.Event
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateOutgoingHook) UnmarshalJSON(data []byte) (err error) {
	varCreateOutgoingHook := _CreateOutgoingHook{}

	err = json.Unmarshal(data, &varCreateOutgoingHook)

	if err != nil {
		return err
	}

	*o = CreateOutgoingHook(varCreateOutgoingHook)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "site_id")
		delete(additionalProperties, "form_id")
		delete(additionalProperties, "form_name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "event")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateOutgoingHook struct {
	value *CreateOutgoingHook
	isSet bool
}

func (v NullableCreateOutgoingHook) Get() *CreateOutgoingHook {
	return v.value
}

func (v *NullableCreateOutgoingHook) Set(val *CreateOutgoingHook) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOutgoingHook) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOutgoingHook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOutgoingHook(val *CreateOutgoingHook) *NullableCreateOutgoingHook {
	return &NullableCreateOutgoingHook{value: val, isSet: true}
}

func (v NullableCreateOutgoingHook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOutgoingHook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


