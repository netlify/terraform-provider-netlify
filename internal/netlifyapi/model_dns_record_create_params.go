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

// checks if the DnsRecordCreateParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnsRecordCreateParams{}

// DnsRecordCreateParams struct for DnsRecordCreateParams
type DnsRecordCreateParams struct {
	// The site ID
	SiteId *string `json:"site_id,omitempty"`
	// Whether Netlify created (managed) record or the user created
	Managed *bool `json:"managed,omitempty"`
	// The value of the DNS record
	Value *string `json:"value,omitempty"`
	// The hostname of the DNS record
	Hostname *string `json:"hostname,omitempty"`
	// The type of the DNS record
	Type *string `json:"type,omitempty"`
	// The TTL of the DNS record
	Ttl *int64 `json:"ttl,omitempty"`
	// The priority of the DNS record
	Priority *int64 `json:"priority,omitempty"`
	// The weight of the DNS record (for SRV type record)
	Weight *int64 `json:"weight,omitempty"`
	// The port of the DNS record (for SRV type record)
	Port *int64 `json:"port,omitempty"`
	// The flag of the DNS record (for CAA type record)
	Flag *int64 `json:"flag,omitempty"`
	// The tag of the DNS record (for CAA type record)
	Tag *string `json:"tag,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DnsRecordCreateParams DnsRecordCreateParams

// NewDnsRecordCreateParams instantiates a new DnsRecordCreateParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsRecordCreateParams() *DnsRecordCreateParams {
	this := DnsRecordCreateParams{}
	return &this
}

// NewDnsRecordCreateParamsWithDefaults instantiates a new DnsRecordCreateParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsRecordCreateParamsWithDefaults() *DnsRecordCreateParams {
	this := DnsRecordCreateParams{}
	return &this
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *DnsRecordCreateParams) GetSiteId() string {
	if o == nil || IsNil(o.SiteId) {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRecordCreateParams) GetSiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *DnsRecordCreateParams) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *DnsRecordCreateParams) SetSiteId(v string) {
	o.SiteId = &v
}

// GetManaged returns the Managed field value if set, zero value otherwise.
func (o *DnsRecordCreateParams) GetManaged() bool {
	if o == nil || IsNil(o.Managed) {
		var ret bool
		return ret
	}
	return *o.Managed
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRecordCreateParams) GetManagedOk() (*bool, bool) {
	if o == nil || IsNil(o.Managed) {
		return nil, false
	}
	return o.Managed, true
}

// HasManaged returns a boolean if a field has been set.
func (o *DnsRecordCreateParams) HasManaged() bool {
	if o != nil && !IsNil(o.Managed) {
		return true
	}

	return false
}

// SetManaged gets a reference to the given bool and assigns it to the Managed field.
func (o *DnsRecordCreateParams) SetManaged(v bool) {
	o.Managed = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DnsRecordCreateParams) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRecordCreateParams) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DnsRecordCreateParams) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *DnsRecordCreateParams) SetValue(v string) {
	o.Value = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *DnsRecordCreateParams) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRecordCreateParams) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *DnsRecordCreateParams) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *DnsRecordCreateParams) SetHostname(v string) {
	o.Hostname = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DnsRecordCreateParams) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRecordCreateParams) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DnsRecordCreateParams) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DnsRecordCreateParams) SetType(v string) {
	o.Type = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *DnsRecordCreateParams) GetTtl() int64 {
	if o == nil || IsNil(o.Ttl) {
		var ret int64
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRecordCreateParams) GetTtlOk() (*int64, bool) {
	if o == nil || IsNil(o.Ttl) {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *DnsRecordCreateParams) HasTtl() bool {
	if o != nil && !IsNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int64 and assigns it to the Ttl field.
func (o *DnsRecordCreateParams) SetTtl(v int64) {
	o.Ttl = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *DnsRecordCreateParams) GetPriority() int64 {
	if o == nil || IsNil(o.Priority) {
		var ret int64
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRecordCreateParams) GetPriorityOk() (*int64, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *DnsRecordCreateParams) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int64 and assigns it to the Priority field.
func (o *DnsRecordCreateParams) SetPriority(v int64) {
	o.Priority = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *DnsRecordCreateParams) GetWeight() int64 {
	if o == nil || IsNil(o.Weight) {
		var ret int64
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRecordCreateParams) GetWeightOk() (*int64, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *DnsRecordCreateParams) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int64 and assigns it to the Weight field.
func (o *DnsRecordCreateParams) SetWeight(v int64) {
	o.Weight = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *DnsRecordCreateParams) GetPort() int64 {
	if o == nil || IsNil(o.Port) {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRecordCreateParams) GetPortOk() (*int64, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *DnsRecordCreateParams) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *DnsRecordCreateParams) SetPort(v int64) {
	o.Port = &v
}

// GetFlag returns the Flag field value if set, zero value otherwise.
func (o *DnsRecordCreateParams) GetFlag() int64 {
	if o == nil || IsNil(o.Flag) {
		var ret int64
		return ret
	}
	return *o.Flag
}

// GetFlagOk returns a tuple with the Flag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRecordCreateParams) GetFlagOk() (*int64, bool) {
	if o == nil || IsNil(o.Flag) {
		return nil, false
	}
	return o.Flag, true
}

// HasFlag returns a boolean if a field has been set.
func (o *DnsRecordCreateParams) HasFlag() bool {
	if o != nil && !IsNil(o.Flag) {
		return true
	}

	return false
}

// SetFlag gets a reference to the given int64 and assigns it to the Flag field.
func (o *DnsRecordCreateParams) SetFlag(v int64) {
	o.Flag = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *DnsRecordCreateParams) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRecordCreateParams) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *DnsRecordCreateParams) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *DnsRecordCreateParams) SetTag(v string) {
	o.Tag = &v
}

func (o DnsRecordCreateParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnsRecordCreateParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SiteId) {
		toSerialize["site_id"] = o.SiteId
	}
	if !IsNil(o.Managed) {
		toSerialize["managed"] = o.Managed
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Ttl) {
		toSerialize["ttl"] = o.Ttl
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Flag) {
		toSerialize["flag"] = o.Flag
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DnsRecordCreateParams) UnmarshalJSON(data []byte) (err error) {
	varDnsRecordCreateParams := _DnsRecordCreateParams{}

	err = json.Unmarshal(data, &varDnsRecordCreateParams)

	if err != nil {
		return err
	}

	*o = DnsRecordCreateParams(varDnsRecordCreateParams)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "site_id")
		delete(additionalProperties, "managed")
		delete(additionalProperties, "value")
		delete(additionalProperties, "hostname")
		delete(additionalProperties, "type")
		delete(additionalProperties, "ttl")
		delete(additionalProperties, "priority")
		delete(additionalProperties, "weight")
		delete(additionalProperties, "port")
		delete(additionalProperties, "flag")
		delete(additionalProperties, "tag")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDnsRecordCreateParams struct {
	value *DnsRecordCreateParams
	isSet bool
}

func (v NullableDnsRecordCreateParams) Get() *DnsRecordCreateParams {
	return v.value
}

func (v *NullableDnsRecordCreateParams) Set(val *DnsRecordCreateParams) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsRecordCreateParams) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsRecordCreateParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsRecordCreateParams(val *DnsRecordCreateParams) *NullableDnsRecordCreateParams {
	return &NullableDnsRecordCreateParams{value: val, isSet: true}
}

func (v NullableDnsRecordCreateParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsRecordCreateParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


