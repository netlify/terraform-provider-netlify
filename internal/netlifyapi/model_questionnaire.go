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

// checks if the Questionnaire type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Questionnaire{}

// Questionnaire struct for Questionnaire
type Questionnaire struct {
	CompanySizeC *string `json:"company_size_c,omitempty"`
	JobLeadershipTitle *string `json:"job_leadership_title,omitempty"`
	JobRole *string `json:"job_role,omitempty"`
	UseCase *string `json:"use_case,omitempty"`
	UseCaseContext *string `json:"use_case_context,omitempty"`
	UseCaseSomethingElse *string `json:"use_case_something_else,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Questionnaire Questionnaire

// NewQuestionnaire instantiates a new Questionnaire object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuestionnaire() *Questionnaire {
	this := Questionnaire{}
	return &this
}

// NewQuestionnaireWithDefaults instantiates a new Questionnaire object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuestionnaireWithDefaults() *Questionnaire {
	this := Questionnaire{}
	return &this
}

// GetCompanySizeC returns the CompanySizeC field value if set, zero value otherwise.
func (o *Questionnaire) GetCompanySizeC() string {
	if o == nil || IsNil(o.CompanySizeC) {
		var ret string
		return ret
	}
	return *o.CompanySizeC
}

// GetCompanySizeCOk returns a tuple with the CompanySizeC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Questionnaire) GetCompanySizeCOk() (*string, bool) {
	if o == nil || IsNil(o.CompanySizeC) {
		return nil, false
	}
	return o.CompanySizeC, true
}

// HasCompanySizeC returns a boolean if a field has been set.
func (o *Questionnaire) HasCompanySizeC() bool {
	if o != nil && !IsNil(o.CompanySizeC) {
		return true
	}

	return false
}

// SetCompanySizeC gets a reference to the given string and assigns it to the CompanySizeC field.
func (o *Questionnaire) SetCompanySizeC(v string) {
	o.CompanySizeC = &v
}

// GetJobLeadershipTitle returns the JobLeadershipTitle field value if set, zero value otherwise.
func (o *Questionnaire) GetJobLeadershipTitle() string {
	if o == nil || IsNil(o.JobLeadershipTitle) {
		var ret string
		return ret
	}
	return *o.JobLeadershipTitle
}

// GetJobLeadershipTitleOk returns a tuple with the JobLeadershipTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Questionnaire) GetJobLeadershipTitleOk() (*string, bool) {
	if o == nil || IsNil(o.JobLeadershipTitle) {
		return nil, false
	}
	return o.JobLeadershipTitle, true
}

// HasJobLeadershipTitle returns a boolean if a field has been set.
func (o *Questionnaire) HasJobLeadershipTitle() bool {
	if o != nil && !IsNil(o.JobLeadershipTitle) {
		return true
	}

	return false
}

// SetJobLeadershipTitle gets a reference to the given string and assigns it to the JobLeadershipTitle field.
func (o *Questionnaire) SetJobLeadershipTitle(v string) {
	o.JobLeadershipTitle = &v
}

// GetJobRole returns the JobRole field value if set, zero value otherwise.
func (o *Questionnaire) GetJobRole() string {
	if o == nil || IsNil(o.JobRole) {
		var ret string
		return ret
	}
	return *o.JobRole
}

// GetJobRoleOk returns a tuple with the JobRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Questionnaire) GetJobRoleOk() (*string, bool) {
	if o == nil || IsNil(o.JobRole) {
		return nil, false
	}
	return o.JobRole, true
}

// HasJobRole returns a boolean if a field has been set.
func (o *Questionnaire) HasJobRole() bool {
	if o != nil && !IsNil(o.JobRole) {
		return true
	}

	return false
}

// SetJobRole gets a reference to the given string and assigns it to the JobRole field.
func (o *Questionnaire) SetJobRole(v string) {
	o.JobRole = &v
}

// GetUseCase returns the UseCase field value if set, zero value otherwise.
func (o *Questionnaire) GetUseCase() string {
	if o == nil || IsNil(o.UseCase) {
		var ret string
		return ret
	}
	return *o.UseCase
}

// GetUseCaseOk returns a tuple with the UseCase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Questionnaire) GetUseCaseOk() (*string, bool) {
	if o == nil || IsNil(o.UseCase) {
		return nil, false
	}
	return o.UseCase, true
}

// HasUseCase returns a boolean if a field has been set.
func (o *Questionnaire) HasUseCase() bool {
	if o != nil && !IsNil(o.UseCase) {
		return true
	}

	return false
}

// SetUseCase gets a reference to the given string and assigns it to the UseCase field.
func (o *Questionnaire) SetUseCase(v string) {
	o.UseCase = &v
}

// GetUseCaseContext returns the UseCaseContext field value if set, zero value otherwise.
func (o *Questionnaire) GetUseCaseContext() string {
	if o == nil || IsNil(o.UseCaseContext) {
		var ret string
		return ret
	}
	return *o.UseCaseContext
}

// GetUseCaseContextOk returns a tuple with the UseCaseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Questionnaire) GetUseCaseContextOk() (*string, bool) {
	if o == nil || IsNil(o.UseCaseContext) {
		return nil, false
	}
	return o.UseCaseContext, true
}

// HasUseCaseContext returns a boolean if a field has been set.
func (o *Questionnaire) HasUseCaseContext() bool {
	if o != nil && !IsNil(o.UseCaseContext) {
		return true
	}

	return false
}

// SetUseCaseContext gets a reference to the given string and assigns it to the UseCaseContext field.
func (o *Questionnaire) SetUseCaseContext(v string) {
	o.UseCaseContext = &v
}

// GetUseCaseSomethingElse returns the UseCaseSomethingElse field value if set, zero value otherwise.
func (o *Questionnaire) GetUseCaseSomethingElse() string {
	if o == nil || IsNil(o.UseCaseSomethingElse) {
		var ret string
		return ret
	}
	return *o.UseCaseSomethingElse
}

// GetUseCaseSomethingElseOk returns a tuple with the UseCaseSomethingElse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Questionnaire) GetUseCaseSomethingElseOk() (*string, bool) {
	if o == nil || IsNil(o.UseCaseSomethingElse) {
		return nil, false
	}
	return o.UseCaseSomethingElse, true
}

// HasUseCaseSomethingElse returns a boolean if a field has been set.
func (o *Questionnaire) HasUseCaseSomethingElse() bool {
	if o != nil && !IsNil(o.UseCaseSomethingElse) {
		return true
	}

	return false
}

// SetUseCaseSomethingElse gets a reference to the given string and assigns it to the UseCaseSomethingElse field.
func (o *Questionnaire) SetUseCaseSomethingElse(v string) {
	o.UseCaseSomethingElse = &v
}

func (o Questionnaire) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Questionnaire) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CompanySizeC) {
		toSerialize["company_size_c"] = o.CompanySizeC
	}
	if !IsNil(o.JobLeadershipTitle) {
		toSerialize["job_leadership_title"] = o.JobLeadershipTitle
	}
	if !IsNil(o.JobRole) {
		toSerialize["job_role"] = o.JobRole
	}
	if !IsNil(o.UseCase) {
		toSerialize["use_case"] = o.UseCase
	}
	if !IsNil(o.UseCaseContext) {
		toSerialize["use_case_context"] = o.UseCaseContext
	}
	if !IsNil(o.UseCaseSomethingElse) {
		toSerialize["use_case_something_else"] = o.UseCaseSomethingElse
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Questionnaire) UnmarshalJSON(data []byte) (err error) {
	varQuestionnaire := _Questionnaire{}

	err = json.Unmarshal(data, &varQuestionnaire)

	if err != nil {
		return err
	}

	*o = Questionnaire(varQuestionnaire)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "company_size_c")
		delete(additionalProperties, "job_leadership_title")
		delete(additionalProperties, "job_role")
		delete(additionalProperties, "use_case")
		delete(additionalProperties, "use_case_context")
		delete(additionalProperties, "use_case_something_else")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuestionnaire struct {
	value *Questionnaire
	isSet bool
}

func (v NullableQuestionnaire) Get() *Questionnaire {
	return v.value
}

func (v *NullableQuestionnaire) Set(val *Questionnaire) {
	v.value = val
	v.isSet = true
}

func (v NullableQuestionnaire) IsSet() bool {
	return v.isSet
}

func (v *NullableQuestionnaire) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuestionnaire(val *Questionnaire) *NullableQuestionnaire {
	return &NullableQuestionnaire{value: val, isSet: true}
}

func (v NullableQuestionnaire) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuestionnaire) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

