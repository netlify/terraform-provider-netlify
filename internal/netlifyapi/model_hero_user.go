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

// checks if the HeroUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HeroUser{}

// HeroUser HeroUser model definition, see UserSerializer for other fields
type HeroUser struct {
	Disabled bool `json:"disabled"`
	DisabledReason string `json:"disabled_reason"`
	Spam bool `json:"spam"`
	SpamScore float32 `json:"spam_score"`
	BillingDetails string `json:"billing_details"`
	PaymentsGatewayName string `json:"payments_gateway_name"`
	GithubSlug string `json:"github_slug"`
	BitbucketSlug string `json:"bitbucket_slug"`
	AllSites int64 `json:"all_sites"`
	SupportPriority int64 `json:"support_priority"`
	ZuoraUrl string `json:"zuora_url"`
	AllAccounts []string `json:"all_accounts"`
	AllOrganizations []string `json:"all_organizations"`
	SiftSpamScore float32 `json:"sift_spam_score"`
	SafeToSpam bool `json:"safe_to_spam"`
	AdditionalProperties map[string]interface{}
}

type _HeroUser HeroUser

// NewHeroUser instantiates a new HeroUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeroUser(disabled bool, disabledReason string, spam bool, spamScore float32, billingDetails string, paymentsGatewayName string, githubSlug string, bitbucketSlug string, allSites int64, supportPriority int64, zuoraUrl string, allAccounts []string, allOrganizations []string, siftSpamScore float32, safeToSpam bool) *HeroUser {
	this := HeroUser{}
	this.Disabled = disabled
	this.DisabledReason = disabledReason
	this.Spam = spam
	this.SpamScore = spamScore
	this.BillingDetails = billingDetails
	this.PaymentsGatewayName = paymentsGatewayName
	this.GithubSlug = githubSlug
	this.BitbucketSlug = bitbucketSlug
	this.AllSites = allSites
	this.SupportPriority = supportPriority
	this.ZuoraUrl = zuoraUrl
	this.AllAccounts = allAccounts
	this.AllOrganizations = allOrganizations
	this.SiftSpamScore = siftSpamScore
	this.SafeToSpam = safeToSpam
	return &this
}

// NewHeroUserWithDefaults instantiates a new HeroUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeroUserWithDefaults() *HeroUser {
	this := HeroUser{}
	return &this
}

// GetDisabled returns the Disabled field value
func (o *HeroUser) GetDisabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value
// and a boolean to check if the value has been set.
func (o *HeroUser) GetDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Disabled, true
}

// SetDisabled sets field value
func (o *HeroUser) SetDisabled(v bool) {
	o.Disabled = v
}

// GetDisabledReason returns the DisabledReason field value
func (o *HeroUser) GetDisabledReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisabledReason
}

// GetDisabledReasonOk returns a tuple with the DisabledReason field value
// and a boolean to check if the value has been set.
func (o *HeroUser) GetDisabledReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisabledReason, true
}

// SetDisabledReason sets field value
func (o *HeroUser) SetDisabledReason(v string) {
	o.DisabledReason = v
}

// GetSpam returns the Spam field value
func (o *HeroUser) GetSpam() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Spam
}

// GetSpamOk returns a tuple with the Spam field value
// and a boolean to check if the value has been set.
func (o *HeroUser) GetSpamOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spam, true
}

// SetSpam sets field value
func (o *HeroUser) SetSpam(v bool) {
	o.Spam = v
}

// GetSpamScore returns the SpamScore field value
func (o *HeroUser) GetSpamScore() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SpamScore
}

// GetSpamScoreOk returns a tuple with the SpamScore field value
// and a boolean to check if the value has been set.
func (o *HeroUser) GetSpamScoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpamScore, true
}

// SetSpamScore sets field value
func (o *HeroUser) SetSpamScore(v float32) {
	o.SpamScore = v
}

// GetBillingDetails returns the BillingDetails field value
func (o *HeroUser) GetBillingDetails() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BillingDetails
}

// GetBillingDetailsOk returns a tuple with the BillingDetails field value
// and a boolean to check if the value has been set.
func (o *HeroUser) GetBillingDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillingDetails, true
}

// SetBillingDetails sets field value
func (o *HeroUser) SetBillingDetails(v string) {
	o.BillingDetails = v
}

// GetPaymentsGatewayName returns the PaymentsGatewayName field value
func (o *HeroUser) GetPaymentsGatewayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentsGatewayName
}

// GetPaymentsGatewayNameOk returns a tuple with the PaymentsGatewayName field value
// and a boolean to check if the value has been set.
func (o *HeroUser) GetPaymentsGatewayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentsGatewayName, true
}

// SetPaymentsGatewayName sets field value
func (o *HeroUser) SetPaymentsGatewayName(v string) {
	o.PaymentsGatewayName = v
}

// GetGithubSlug returns the GithubSlug field value
func (o *HeroUser) GetGithubSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GithubSlug
}

// GetGithubSlugOk returns a tuple with the GithubSlug field value
// and a boolean to check if the value has been set.
func (o *HeroUser) GetGithubSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GithubSlug, true
}

// SetGithubSlug sets field value
func (o *HeroUser) SetGithubSlug(v string) {
	o.GithubSlug = v
}

// GetBitbucketSlug returns the BitbucketSlug field value
func (o *HeroUser) GetBitbucketSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BitbucketSlug
}

// GetBitbucketSlugOk returns a tuple with the BitbucketSlug field value
// and a boolean to check if the value has been set.
func (o *HeroUser) GetBitbucketSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BitbucketSlug, true
}

// SetBitbucketSlug sets field value
func (o *HeroUser) SetBitbucketSlug(v string) {
	o.BitbucketSlug = v
}

// GetAllSites returns the AllSites field value
func (o *HeroUser) GetAllSites() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AllSites
}

// GetAllSitesOk returns a tuple with the AllSites field value
// and a boolean to check if the value has been set.
func (o *HeroUser) GetAllSitesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllSites, true
}

// SetAllSites sets field value
func (o *HeroUser) SetAllSites(v int64) {
	o.AllSites = v
}

// GetSupportPriority returns the SupportPriority field value
func (o *HeroUser) GetSupportPriority() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SupportPriority
}

// GetSupportPriorityOk returns a tuple with the SupportPriority field value
// and a boolean to check if the value has been set.
func (o *HeroUser) GetSupportPriorityOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportPriority, true
}

// SetSupportPriority sets field value
func (o *HeroUser) SetSupportPriority(v int64) {
	o.SupportPriority = v
}

// GetZuoraUrl returns the ZuoraUrl field value
func (o *HeroUser) GetZuoraUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ZuoraUrl
}

// GetZuoraUrlOk returns a tuple with the ZuoraUrl field value
// and a boolean to check if the value has been set.
func (o *HeroUser) GetZuoraUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ZuoraUrl, true
}

// SetZuoraUrl sets field value
func (o *HeroUser) SetZuoraUrl(v string) {
	o.ZuoraUrl = v
}

// GetAllAccounts returns the AllAccounts field value
func (o *HeroUser) GetAllAccounts() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AllAccounts
}

// GetAllAccountsOk returns a tuple with the AllAccounts field value
// and a boolean to check if the value has been set.
func (o *HeroUser) GetAllAccountsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllAccounts, true
}

// SetAllAccounts sets field value
func (o *HeroUser) SetAllAccounts(v []string) {
	o.AllAccounts = v
}

// GetAllOrganizations returns the AllOrganizations field value
func (o *HeroUser) GetAllOrganizations() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AllOrganizations
}

// GetAllOrganizationsOk returns a tuple with the AllOrganizations field value
// and a boolean to check if the value has been set.
func (o *HeroUser) GetAllOrganizationsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllOrganizations, true
}

// SetAllOrganizations sets field value
func (o *HeroUser) SetAllOrganizations(v []string) {
	o.AllOrganizations = v
}

// GetSiftSpamScore returns the SiftSpamScore field value
func (o *HeroUser) GetSiftSpamScore() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SiftSpamScore
}

// GetSiftSpamScoreOk returns a tuple with the SiftSpamScore field value
// and a boolean to check if the value has been set.
func (o *HeroUser) GetSiftSpamScoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiftSpamScore, true
}

// SetSiftSpamScore sets field value
func (o *HeroUser) SetSiftSpamScore(v float32) {
	o.SiftSpamScore = v
}

// GetSafeToSpam returns the SafeToSpam field value
func (o *HeroUser) GetSafeToSpam() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SafeToSpam
}

// GetSafeToSpamOk returns a tuple with the SafeToSpam field value
// and a boolean to check if the value has been set.
func (o *HeroUser) GetSafeToSpamOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SafeToSpam, true
}

// SetSafeToSpam sets field value
func (o *HeroUser) SetSafeToSpam(v bool) {
	o.SafeToSpam = v
}

func (o HeroUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HeroUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["disabled"] = o.Disabled
	toSerialize["disabled_reason"] = o.DisabledReason
	toSerialize["spam"] = o.Spam
	toSerialize["spam_score"] = o.SpamScore
	toSerialize["billing_details"] = o.BillingDetails
	toSerialize["payments_gateway_name"] = o.PaymentsGatewayName
	toSerialize["github_slug"] = o.GithubSlug
	toSerialize["bitbucket_slug"] = o.BitbucketSlug
	toSerialize["all_sites"] = o.AllSites
	toSerialize["support_priority"] = o.SupportPriority
	toSerialize["zuora_url"] = o.ZuoraUrl
	toSerialize["all_accounts"] = o.AllAccounts
	toSerialize["all_organizations"] = o.AllOrganizations
	toSerialize["sift_spam_score"] = o.SiftSpamScore
	toSerialize["safe_to_spam"] = o.SafeToSpam

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HeroUser) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"disabled",
		"disabled_reason",
		"spam",
		"spam_score",
		"billing_details",
		"payments_gateway_name",
		"github_slug",
		"bitbucket_slug",
		"all_sites",
		"support_priority",
		"zuora_url",
		"all_accounts",
		"all_organizations",
		"sift_spam_score",
		"safe_to_spam",
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

	varHeroUser := _HeroUser{}

	err = json.Unmarshal(data, &varHeroUser)

	if err != nil {
		return err
	}

	*o = HeroUser(varHeroUser)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "disabled")
		delete(additionalProperties, "disabled_reason")
		delete(additionalProperties, "spam")
		delete(additionalProperties, "spam_score")
		delete(additionalProperties, "billing_details")
		delete(additionalProperties, "payments_gateway_name")
		delete(additionalProperties, "github_slug")
		delete(additionalProperties, "bitbucket_slug")
		delete(additionalProperties, "all_sites")
		delete(additionalProperties, "support_priority")
		delete(additionalProperties, "zuora_url")
		delete(additionalProperties, "all_accounts")
		delete(additionalProperties, "all_organizations")
		delete(additionalProperties, "sift_spam_score")
		delete(additionalProperties, "safe_to_spam")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHeroUser struct {
	value *HeroUser
	isSet bool
}

func (v NullableHeroUser) Get() *HeroUser {
	return v.value
}

func (v *NullableHeroUser) Set(val *HeroUser) {
	v.value = val
	v.isSet = true
}

func (v NullableHeroUser) IsSet() bool {
	return v.isSet
}

func (v *NullableHeroUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeroUser(val *HeroUser) *NullableHeroUser {
	return &NullableHeroUser{value: val, isSet: true}
}

func (v NullableHeroUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeroUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


