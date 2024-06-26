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

// checks if the FormSubmission type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormSubmission{}

// FormSubmission struct for FormSubmission
type FormSubmission struct {
	Id string `json:"id"`
	FormId string `json:"form_id"`
	FormName string `json:"form_name"`
	SiteUrl string `json:"site_url"`
	SiteName string `json:"site_name"`
	Number int64 `json:"number"`
	Email string `json:"email"`
	Name string `json:"name"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Company string `json:"company"`
	Summary string `json:"summary"`
	Body string `json:"body"`
	Data map[string]interface{} `json:"data"`
	CreatedAt time.Time `json:"created_at"`
	AdditionalProperties map[string]interface{}
}

type _FormSubmission FormSubmission

// NewFormSubmission instantiates a new FormSubmission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormSubmission(id string, formId string, formName string, siteUrl string, siteName string, number int64, email string, name string, firstName string, lastName string, company string, summary string, body string, data map[string]interface{}, createdAt time.Time) *FormSubmission {
	this := FormSubmission{}
	this.Id = id
	this.FormId = formId
	this.FormName = formName
	this.SiteUrl = siteUrl
	this.SiteName = siteName
	this.Number = number
	this.Email = email
	this.Name = name
	this.FirstName = firstName
	this.LastName = lastName
	this.Company = company
	this.Summary = summary
	this.Body = body
	this.Data = data
	this.CreatedAt = createdAt
	return &this
}

// NewFormSubmissionWithDefaults instantiates a new FormSubmission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormSubmissionWithDefaults() *FormSubmission {
	this := FormSubmission{}
	return &this
}

// GetId returns the Id field value
func (o *FormSubmission) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FormSubmission) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FormSubmission) SetId(v string) {
	o.Id = v
}

// GetFormId returns the FormId field value
func (o *FormSubmission) GetFormId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FormId
}

// GetFormIdOk returns a tuple with the FormId field value
// and a boolean to check if the value has been set.
func (o *FormSubmission) GetFormIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FormId, true
}

// SetFormId sets field value
func (o *FormSubmission) SetFormId(v string) {
	o.FormId = v
}

// GetFormName returns the FormName field value
func (o *FormSubmission) GetFormName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FormName
}

// GetFormNameOk returns a tuple with the FormName field value
// and a boolean to check if the value has been set.
func (o *FormSubmission) GetFormNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FormName, true
}

// SetFormName sets field value
func (o *FormSubmission) SetFormName(v string) {
	o.FormName = v
}

// GetSiteUrl returns the SiteUrl field value
func (o *FormSubmission) GetSiteUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SiteUrl
}

// GetSiteUrlOk returns a tuple with the SiteUrl field value
// and a boolean to check if the value has been set.
func (o *FormSubmission) GetSiteUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteUrl, true
}

// SetSiteUrl sets field value
func (o *FormSubmission) SetSiteUrl(v string) {
	o.SiteUrl = v
}

// GetSiteName returns the SiteName field value
func (o *FormSubmission) GetSiteName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value
// and a boolean to check if the value has been set.
func (o *FormSubmission) GetSiteNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteName, true
}

// SetSiteName sets field value
func (o *FormSubmission) SetSiteName(v string) {
	o.SiteName = v
}

// GetNumber returns the Number field value
func (o *FormSubmission) GetNumber() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *FormSubmission) GetNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *FormSubmission) SetNumber(v int64) {
	o.Number = v
}

// GetEmail returns the Email field value
func (o *FormSubmission) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *FormSubmission) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *FormSubmission) SetEmail(v string) {
	o.Email = v
}

// GetName returns the Name field value
func (o *FormSubmission) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FormSubmission) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FormSubmission) SetName(v string) {
	o.Name = v
}

// GetFirstName returns the FirstName field value
func (o *FormSubmission) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *FormSubmission) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *FormSubmission) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *FormSubmission) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *FormSubmission) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *FormSubmission) SetLastName(v string) {
	o.LastName = v
}

// GetCompany returns the Company field value
func (o *FormSubmission) GetCompany() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Company
}

// GetCompanyOk returns a tuple with the Company field value
// and a boolean to check if the value has been set.
func (o *FormSubmission) GetCompanyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Company, true
}

// SetCompany sets field value
func (o *FormSubmission) SetCompany(v string) {
	o.Company = v
}

// GetSummary returns the Summary field value
func (o *FormSubmission) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *FormSubmission) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value
func (o *FormSubmission) SetSummary(v string) {
	o.Summary = v
}

// GetBody returns the Body field value
func (o *FormSubmission) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *FormSubmission) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *FormSubmission) SetBody(v string) {
	o.Body = v
}

// GetData returns the Data field value
func (o *FormSubmission) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *FormSubmission) GetDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *FormSubmission) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *FormSubmission) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *FormSubmission) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *FormSubmission) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o FormSubmission) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormSubmission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["form_id"] = o.FormId
	toSerialize["form_name"] = o.FormName
	toSerialize["site_url"] = o.SiteUrl
	toSerialize["site_name"] = o.SiteName
	toSerialize["number"] = o.Number
	toSerialize["email"] = o.Email
	toSerialize["name"] = o.Name
	toSerialize["first_name"] = o.FirstName
	toSerialize["last_name"] = o.LastName
	toSerialize["company"] = o.Company
	toSerialize["summary"] = o.Summary
	toSerialize["body"] = o.Body
	toSerialize["data"] = o.Data
	toSerialize["created_at"] = o.CreatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FormSubmission) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"form_id",
		"form_name",
		"site_url",
		"site_name",
		"number",
		"email",
		"name",
		"first_name",
		"last_name",
		"company",
		"summary",
		"body",
		"data",
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

	varFormSubmission := _FormSubmission{}

	err = json.Unmarshal(data, &varFormSubmission)

	if err != nil {
		return err
	}

	*o = FormSubmission(varFormSubmission)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "form_id")
		delete(additionalProperties, "form_name")
		delete(additionalProperties, "site_url")
		delete(additionalProperties, "site_name")
		delete(additionalProperties, "number")
		delete(additionalProperties, "email")
		delete(additionalProperties, "name")
		delete(additionalProperties, "first_name")
		delete(additionalProperties, "last_name")
		delete(additionalProperties, "company")
		delete(additionalProperties, "summary")
		delete(additionalProperties, "body")
		delete(additionalProperties, "data")
		delete(additionalProperties, "created_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFormSubmission struct {
	value *FormSubmission
	isSet bool
}

func (v NullableFormSubmission) Get() *FormSubmission {
	return v.value
}

func (v *NullableFormSubmission) Set(val *FormSubmission) {
	v.value = val
	v.isSet = true
}

func (v NullableFormSubmission) IsSet() bool {
	return v.isSet
}

func (v *NullableFormSubmission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormSubmission(val *FormSubmission) *NullableFormSubmission {
	return &NullableFormSubmission{value: val, isSet: true}
}

func (v NullableFormSubmission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormSubmission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


