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

// checks if the SiteFileUploadResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteFileUploadResponse{}

// SiteFileUploadResponse struct for SiteFileUploadResponse
type SiteFileUploadResponse struct {
	// The path of the file
	Id *string `json:"id,omitempty"`
	// The path of the file
	Path *string `json:"path,omitempty"`
	// The sha of the file
	Sha *string `json:"sha,omitempty"`
	// The mime_type of the file
	MimeType *string `json:"mime_type,omitempty"`
	// The size of the file in bytes
	Size *int64 `json:"size,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SiteFileUploadResponse SiteFileUploadResponse

// NewSiteFileUploadResponse instantiates a new SiteFileUploadResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteFileUploadResponse() *SiteFileUploadResponse {
	this := SiteFileUploadResponse{}
	return &this
}

// NewSiteFileUploadResponseWithDefaults instantiates a new SiteFileUploadResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteFileUploadResponseWithDefaults() *SiteFileUploadResponse {
	this := SiteFileUploadResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SiteFileUploadResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteFileUploadResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SiteFileUploadResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SiteFileUploadResponse) SetId(v string) {
	o.Id = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *SiteFileUploadResponse) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteFileUploadResponse) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *SiteFileUploadResponse) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *SiteFileUploadResponse) SetPath(v string) {
	o.Path = &v
}

// GetSha returns the Sha field value if set, zero value otherwise.
func (o *SiteFileUploadResponse) GetSha() string {
	if o == nil || IsNil(o.Sha) {
		var ret string
		return ret
	}
	return *o.Sha
}

// GetShaOk returns a tuple with the Sha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteFileUploadResponse) GetShaOk() (*string, bool) {
	if o == nil || IsNil(o.Sha) {
		return nil, false
	}
	return o.Sha, true
}

// HasSha returns a boolean if a field has been set.
func (o *SiteFileUploadResponse) HasSha() bool {
	if o != nil && !IsNil(o.Sha) {
		return true
	}

	return false
}

// SetSha gets a reference to the given string and assigns it to the Sha field.
func (o *SiteFileUploadResponse) SetSha(v string) {
	o.Sha = &v
}

// GetMimeType returns the MimeType field value if set, zero value otherwise.
func (o *SiteFileUploadResponse) GetMimeType() string {
	if o == nil || IsNil(o.MimeType) {
		var ret string
		return ret
	}
	return *o.MimeType
}

// GetMimeTypeOk returns a tuple with the MimeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteFileUploadResponse) GetMimeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MimeType) {
		return nil, false
	}
	return o.MimeType, true
}

// HasMimeType returns a boolean if a field has been set.
func (o *SiteFileUploadResponse) HasMimeType() bool {
	if o != nil && !IsNil(o.MimeType) {
		return true
	}

	return false
}

// SetMimeType gets a reference to the given string and assigns it to the MimeType field.
func (o *SiteFileUploadResponse) SetMimeType(v string) {
	o.MimeType = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *SiteFileUploadResponse) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteFileUploadResponse) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *SiteFileUploadResponse) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *SiteFileUploadResponse) SetSize(v int64) {
	o.Size = &v
}

func (o SiteFileUploadResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteFileUploadResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Sha) {
		toSerialize["sha"] = o.Sha
	}
	if !IsNil(o.MimeType) {
		toSerialize["mime_type"] = o.MimeType
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SiteFileUploadResponse) UnmarshalJSON(data []byte) (err error) {
	varSiteFileUploadResponse := _SiteFileUploadResponse{}

	err = json.Unmarshal(data, &varSiteFileUploadResponse)

	if err != nil {
		return err
	}

	*o = SiteFileUploadResponse(varSiteFileUploadResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "path")
		delete(additionalProperties, "sha")
		delete(additionalProperties, "mime_type")
		delete(additionalProperties, "size")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSiteFileUploadResponse struct {
	value *SiteFileUploadResponse
	isSet bool
}

func (v NullableSiteFileUploadResponse) Get() *SiteFileUploadResponse {
	return v.value
}

func (v *NullableSiteFileUploadResponse) Set(val *SiteFileUploadResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteFileUploadResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteFileUploadResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteFileUploadResponse(val *SiteFileUploadResponse) *NullableSiteFileUploadResponse {
	return &NullableSiteFileUploadResponse{value: val, isSet: true}
}

func (v NullableSiteFileUploadResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteFileUploadResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


