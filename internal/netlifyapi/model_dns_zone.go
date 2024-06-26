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

// checks if the DnsZone type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnsZone{}

// DnsZone struct for DnsZone
type DnsZone struct {
	// The ID of the DNS zone
	Id string `json:"id"`
	// The name of the DNS zone
	Name string `json:"name"`
	// An array of error messages
	Errors []string `json:"errors"`
	// An array of supported DNS record types
	SupportedRecordTypes []string `json:"supported_record_types"`
	// The user ID of the DNS zone creator
	UserId string `json:"user_id"`
	// An array of DNS records for this DNS zone
	Records []DnsRecord `json:"records"`
	// The name servers of the DNS zone
	DnsServers []string `json:"dns_servers"`
	// The account ID
	AccountId string `json:"account_id"`
	// The site ID
	SiteId string `json:"site_id"`
	// The account slug
	AccountSlug string `json:"account_slug"`
	// The account name
	AccountName string `json:"account_name"`
	Domain *Domain `json:"domain,omitempty"`
	// Whether IPv6 is enabled
	Ipv6Enabled bool `json:"ipv6_enabled"`
	// Whether using a dedicated network
	Dedicated bool `json:"dedicated"`
	// When the DNS zone was created
	CreatedAt time.Time `json:"created_at"`
	// When the DNS zone was updated
	UpdatedAt time.Time `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _DnsZone DnsZone

// NewDnsZone instantiates a new DnsZone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsZone(id string, name string, errors []string, supportedRecordTypes []string, userId string, records []DnsRecord, dnsServers []string, accountId string, siteId string, accountSlug string, accountName string, ipv6Enabled bool, dedicated bool, createdAt time.Time, updatedAt time.Time) *DnsZone {
	this := DnsZone{}
	this.Id = id
	this.Name = name
	this.Errors = errors
	this.SupportedRecordTypes = supportedRecordTypes
	this.UserId = userId
	this.Records = records
	this.DnsServers = dnsServers
	this.AccountId = accountId
	this.SiteId = siteId
	this.AccountSlug = accountSlug
	this.AccountName = accountName
	this.Ipv6Enabled = ipv6Enabled
	this.Dedicated = dedicated
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewDnsZoneWithDefaults instantiates a new DnsZone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsZoneWithDefaults() *DnsZone {
	this := DnsZone{}
	return &this
}

// GetId returns the Id field value
func (o *DnsZone) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DnsZone) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DnsZone) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *DnsZone) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DnsZone) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DnsZone) SetName(v string) {
	o.Name = v
}

// GetErrors returns the Errors field value
func (o *DnsZone) GetErrors() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *DnsZone) GetErrorsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Errors, true
}

// SetErrors sets field value
func (o *DnsZone) SetErrors(v []string) {
	o.Errors = v
}

// GetSupportedRecordTypes returns the SupportedRecordTypes field value
func (o *DnsZone) GetSupportedRecordTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SupportedRecordTypes
}

// GetSupportedRecordTypesOk returns a tuple with the SupportedRecordTypes field value
// and a boolean to check if the value has been set.
func (o *DnsZone) GetSupportedRecordTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SupportedRecordTypes, true
}

// SetSupportedRecordTypes sets field value
func (o *DnsZone) SetSupportedRecordTypes(v []string) {
	o.SupportedRecordTypes = v
}

// GetUserId returns the UserId field value
func (o *DnsZone) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *DnsZone) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *DnsZone) SetUserId(v string) {
	o.UserId = v
}

// GetRecords returns the Records field value
func (o *DnsZone) GetRecords() []DnsRecord {
	if o == nil {
		var ret []DnsRecord
		return ret
	}

	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value
// and a boolean to check if the value has been set.
func (o *DnsZone) GetRecordsOk() ([]DnsRecord, bool) {
	if o == nil {
		return nil, false
	}
	return o.Records, true
}

// SetRecords sets field value
func (o *DnsZone) SetRecords(v []DnsRecord) {
	o.Records = v
}

// GetDnsServers returns the DnsServers field value
func (o *DnsZone) GetDnsServers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DnsServers
}

// GetDnsServersOk returns a tuple with the DnsServers field value
// and a boolean to check if the value has been set.
func (o *DnsZone) GetDnsServersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DnsServers, true
}

// SetDnsServers sets field value
func (o *DnsZone) SetDnsServers(v []string) {
	o.DnsServers = v
}

// GetAccountId returns the AccountId field value
func (o *DnsZone) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *DnsZone) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *DnsZone) SetAccountId(v string) {
	o.AccountId = v
}

// GetSiteId returns the SiteId field value
func (o *DnsZone) GetSiteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value
// and a boolean to check if the value has been set.
func (o *DnsZone) GetSiteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteId, true
}

// SetSiteId sets field value
func (o *DnsZone) SetSiteId(v string) {
	o.SiteId = v
}

// GetAccountSlug returns the AccountSlug field value
func (o *DnsZone) GetAccountSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountSlug
}

// GetAccountSlugOk returns a tuple with the AccountSlug field value
// and a boolean to check if the value has been set.
func (o *DnsZone) GetAccountSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountSlug, true
}

// SetAccountSlug sets field value
func (o *DnsZone) SetAccountSlug(v string) {
	o.AccountSlug = v
}

// GetAccountName returns the AccountName field value
func (o *DnsZone) GetAccountName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value
// and a boolean to check if the value has been set.
func (o *DnsZone) GetAccountNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountName, true
}

// SetAccountName sets field value
func (o *DnsZone) SetAccountName(v string) {
	o.AccountName = v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *DnsZone) GetDomain() Domain {
	if o == nil || IsNil(o.Domain) {
		var ret Domain
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsZone) GetDomainOk() (*Domain, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *DnsZone) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given Domain and assigns it to the Domain field.
func (o *DnsZone) SetDomain(v Domain) {
	o.Domain = &v
}

// GetIpv6Enabled returns the Ipv6Enabled field value
func (o *DnsZone) GetIpv6Enabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Ipv6Enabled
}

// GetIpv6EnabledOk returns a tuple with the Ipv6Enabled field value
// and a boolean to check if the value has been set.
func (o *DnsZone) GetIpv6EnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ipv6Enabled, true
}

// SetIpv6Enabled sets field value
func (o *DnsZone) SetIpv6Enabled(v bool) {
	o.Ipv6Enabled = v
}

// GetDedicated returns the Dedicated field value
func (o *DnsZone) GetDedicated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Dedicated
}

// GetDedicatedOk returns a tuple with the Dedicated field value
// and a boolean to check if the value has been set.
func (o *DnsZone) GetDedicatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dedicated, true
}

// SetDedicated sets field value
func (o *DnsZone) SetDedicated(v bool) {
	o.Dedicated = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *DnsZone) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DnsZone) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *DnsZone) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *DnsZone) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *DnsZone) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *DnsZone) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o DnsZone) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnsZone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["errors"] = o.Errors
	toSerialize["supported_record_types"] = o.SupportedRecordTypes
	toSerialize["user_id"] = o.UserId
	toSerialize["records"] = o.Records
	toSerialize["dns_servers"] = o.DnsServers
	toSerialize["account_id"] = o.AccountId
	toSerialize["site_id"] = o.SiteId
	toSerialize["account_slug"] = o.AccountSlug
	toSerialize["account_name"] = o.AccountName
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	toSerialize["ipv6_enabled"] = o.Ipv6Enabled
	toSerialize["dedicated"] = o.Dedicated
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DnsZone) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"errors",
		"supported_record_types",
		"user_id",
		"records",
		"dns_servers",
		"account_id",
		"site_id",
		"account_slug",
		"account_name",
		"ipv6_enabled",
		"dedicated",
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

	varDnsZone := _DnsZone{}

	err = json.Unmarshal(data, &varDnsZone)

	if err != nil {
		return err
	}

	*o = DnsZone(varDnsZone)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "errors")
		delete(additionalProperties, "supported_record_types")
		delete(additionalProperties, "user_id")
		delete(additionalProperties, "records")
		delete(additionalProperties, "dns_servers")
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "site_id")
		delete(additionalProperties, "account_slug")
		delete(additionalProperties, "account_name")
		delete(additionalProperties, "domain")
		delete(additionalProperties, "ipv6_enabled")
		delete(additionalProperties, "dedicated")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDnsZone struct {
	value *DnsZone
	isSet bool
}

func (v NullableDnsZone) Get() *DnsZone {
	return v.value
}

func (v *NullableDnsZone) Set(val *DnsZone) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsZone) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsZone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsZone(val *DnsZone) *NullableDnsZone {
	return &NullableDnsZone{value: val, isSet: true}
}

func (v NullableDnsZone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsZone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


