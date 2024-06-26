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

// checks if the EdgeFunctionsUsage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EdgeFunctionsUsage{}

// EdgeFunctionsUsage struct for EdgeFunctionsUsage
type EdgeFunctionsUsage struct {
	Used int64 `json:"used"`
	Included int64 `json:"included"`
	Additional int64 `json:"additional"`
	LastUpdatedAt time.Time `json:"last_updated_at"`
	PeriodStartDate time.Time `json:"period_start_date"`
	PeriodEndDate time.Time `json:"period_end_date"`
	AdditionalProperties map[string]interface{}
}

type _EdgeFunctionsUsage EdgeFunctionsUsage

// NewEdgeFunctionsUsage instantiates a new EdgeFunctionsUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgeFunctionsUsage(used int64, included int64, additional int64, lastUpdatedAt time.Time, periodStartDate time.Time, periodEndDate time.Time) *EdgeFunctionsUsage {
	this := EdgeFunctionsUsage{}
	this.Used = used
	this.Included = included
	this.Additional = additional
	this.LastUpdatedAt = lastUpdatedAt
	this.PeriodStartDate = periodStartDate
	this.PeriodEndDate = periodEndDate
	return &this
}

// NewEdgeFunctionsUsageWithDefaults instantiates a new EdgeFunctionsUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeFunctionsUsageWithDefaults() *EdgeFunctionsUsage {
	this := EdgeFunctionsUsage{}
	return &this
}

// GetUsed returns the Used field value
func (o *EdgeFunctionsUsage) GetUsed() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Used
}

// GetUsedOk returns a tuple with the Used field value
// and a boolean to check if the value has been set.
func (o *EdgeFunctionsUsage) GetUsedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Used, true
}

// SetUsed sets field value
func (o *EdgeFunctionsUsage) SetUsed(v int64) {
	o.Used = v
}

// GetIncluded returns the Included field value
func (o *EdgeFunctionsUsage) GetIncluded() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value
// and a boolean to check if the value has been set.
func (o *EdgeFunctionsUsage) GetIncludedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Included, true
}

// SetIncluded sets field value
func (o *EdgeFunctionsUsage) SetIncluded(v int64) {
	o.Included = v
}

// GetAdditional returns the Additional field value
func (o *EdgeFunctionsUsage) GetAdditional() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Additional
}

// GetAdditionalOk returns a tuple with the Additional field value
// and a boolean to check if the value has been set.
func (o *EdgeFunctionsUsage) GetAdditionalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Additional, true
}

// SetAdditional sets field value
func (o *EdgeFunctionsUsage) SetAdditional(v int64) {
	o.Additional = v
}

// GetLastUpdatedAt returns the LastUpdatedAt field value
func (o *EdgeFunctionsUsage) GetLastUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdatedAt
}

// GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field value
// and a boolean to check if the value has been set.
func (o *EdgeFunctionsUsage) GetLastUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedAt, true
}

// SetLastUpdatedAt sets field value
func (o *EdgeFunctionsUsage) SetLastUpdatedAt(v time.Time) {
	o.LastUpdatedAt = v
}

// GetPeriodStartDate returns the PeriodStartDate field value
func (o *EdgeFunctionsUsage) GetPeriodStartDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.PeriodStartDate
}

// GetPeriodStartDateOk returns a tuple with the PeriodStartDate field value
// and a boolean to check if the value has been set.
func (o *EdgeFunctionsUsage) GetPeriodStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PeriodStartDate, true
}

// SetPeriodStartDate sets field value
func (o *EdgeFunctionsUsage) SetPeriodStartDate(v time.Time) {
	o.PeriodStartDate = v
}

// GetPeriodEndDate returns the PeriodEndDate field value
func (o *EdgeFunctionsUsage) GetPeriodEndDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.PeriodEndDate
}

// GetPeriodEndDateOk returns a tuple with the PeriodEndDate field value
// and a boolean to check if the value has been set.
func (o *EdgeFunctionsUsage) GetPeriodEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PeriodEndDate, true
}

// SetPeriodEndDate sets field value
func (o *EdgeFunctionsUsage) SetPeriodEndDate(v time.Time) {
	o.PeriodEndDate = v
}

func (o EdgeFunctionsUsage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EdgeFunctionsUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["used"] = o.Used
	toSerialize["included"] = o.Included
	toSerialize["additional"] = o.Additional
	toSerialize["last_updated_at"] = o.LastUpdatedAt
	toSerialize["period_start_date"] = o.PeriodStartDate
	toSerialize["period_end_date"] = o.PeriodEndDate

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EdgeFunctionsUsage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"used",
		"included",
		"additional",
		"last_updated_at",
		"period_start_date",
		"period_end_date",
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

	varEdgeFunctionsUsage := _EdgeFunctionsUsage{}

	err = json.Unmarshal(data, &varEdgeFunctionsUsage)

	if err != nil {
		return err
	}

	*o = EdgeFunctionsUsage(varEdgeFunctionsUsage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "used")
		delete(additionalProperties, "included")
		delete(additionalProperties, "additional")
		delete(additionalProperties, "last_updated_at")
		delete(additionalProperties, "period_start_date")
		delete(additionalProperties, "period_end_date")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEdgeFunctionsUsage struct {
	value *EdgeFunctionsUsage
	isSet bool
}

func (v NullableEdgeFunctionsUsage) Get() *EdgeFunctionsUsage {
	return v.value
}

func (v *NullableEdgeFunctionsUsage) Set(val *EdgeFunctionsUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeFunctionsUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeFunctionsUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeFunctionsUsage(val *EdgeFunctionsUsage) *NullableEdgeFunctionsUsage {
	return &NullableEdgeFunctionsUsage{value: val, isSet: true}
}

func (v NullableEdgeFunctionsUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeFunctionsUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


