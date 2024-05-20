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

// checks if the DeploySimpleState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploySimpleState{}

// DeploySimpleState struct for DeploySimpleState
type DeploySimpleState struct {
	Id string `json:"id"`
	SiteId string `json:"site_id"`
	State string `json:"state"`
	// When the deploy was created
	CreatedAt time.Time `json:"created_at"`
	// The total time, in seconds, it took to deploy
	DeployTime int64 `json:"deploy_time"`
	// The deploy context
	Context string `json:"context"`
	ConcurrentUploadLimit int64 `json:"concurrent_upload_limit"`
	AdditionalProperties map[string]interface{}
}

type _DeploySimpleState DeploySimpleState

// NewDeploySimpleState instantiates a new DeploySimpleState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploySimpleState(id string, siteId string, state string, createdAt time.Time, deployTime int64, context string, concurrentUploadLimit int64) *DeploySimpleState {
	this := DeploySimpleState{}
	this.Id = id
	this.SiteId = siteId
	this.State = state
	this.CreatedAt = createdAt
	this.DeployTime = deployTime
	this.Context = context
	this.ConcurrentUploadLimit = concurrentUploadLimit
	return &this
}

// NewDeploySimpleStateWithDefaults instantiates a new DeploySimpleState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploySimpleStateWithDefaults() *DeploySimpleState {
	this := DeploySimpleState{}
	return &this
}

// GetId returns the Id field value
func (o *DeploySimpleState) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeploySimpleState) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeploySimpleState) SetId(v string) {
	o.Id = v
}

// GetSiteId returns the SiteId field value
func (o *DeploySimpleState) GetSiteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value
// and a boolean to check if the value has been set.
func (o *DeploySimpleState) GetSiteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteId, true
}

// SetSiteId sets field value
func (o *DeploySimpleState) SetSiteId(v string) {
	o.SiteId = v
}

// GetState returns the State field value
func (o *DeploySimpleState) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *DeploySimpleState) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *DeploySimpleState) SetState(v string) {
	o.State = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *DeploySimpleState) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DeploySimpleState) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *DeploySimpleState) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDeployTime returns the DeployTime field value
func (o *DeploySimpleState) GetDeployTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DeployTime
}

// GetDeployTimeOk returns a tuple with the DeployTime field value
// and a boolean to check if the value has been set.
func (o *DeploySimpleState) GetDeployTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeployTime, true
}

// SetDeployTime sets field value
func (o *DeploySimpleState) SetDeployTime(v int64) {
	o.DeployTime = v
}

// GetContext returns the Context field value
func (o *DeploySimpleState) GetContext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *DeploySimpleState) GetContextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value
func (o *DeploySimpleState) SetContext(v string) {
	o.Context = v
}

// GetConcurrentUploadLimit returns the ConcurrentUploadLimit field value
func (o *DeploySimpleState) GetConcurrentUploadLimit() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ConcurrentUploadLimit
}

// GetConcurrentUploadLimitOk returns a tuple with the ConcurrentUploadLimit field value
// and a boolean to check if the value has been set.
func (o *DeploySimpleState) GetConcurrentUploadLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConcurrentUploadLimit, true
}

// SetConcurrentUploadLimit sets field value
func (o *DeploySimpleState) SetConcurrentUploadLimit(v int64) {
	o.ConcurrentUploadLimit = v
}

func (o DeploySimpleState) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploySimpleState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["site_id"] = o.SiteId
	toSerialize["state"] = o.State
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["deploy_time"] = o.DeployTime
	toSerialize["context"] = o.Context
	toSerialize["concurrent_upload_limit"] = o.ConcurrentUploadLimit

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeploySimpleState) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"site_id",
		"state",
		"created_at",
		"deploy_time",
		"context",
		"concurrent_upload_limit",
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

	varDeploySimpleState := _DeploySimpleState{}

	err = json.Unmarshal(data, &varDeploySimpleState)

	if err != nil {
		return err
	}

	*o = DeploySimpleState(varDeploySimpleState)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "site_id")
		delete(additionalProperties, "state")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deploy_time")
		delete(additionalProperties, "context")
		delete(additionalProperties, "concurrent_upload_limit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeploySimpleState struct {
	value *DeploySimpleState
	isSet bool
}

func (v NullableDeploySimpleState) Get() *DeploySimpleState {
	return v.value
}

func (v *NullableDeploySimpleState) Set(val *DeploySimpleState) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploySimpleState) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploySimpleState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploySimpleState(val *DeploySimpleState) *NullableDeploySimpleState {
	return &NullableDeploySimpleState{value: val, isSet: true}
}

func (v NullableDeploySimpleState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploySimpleState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

