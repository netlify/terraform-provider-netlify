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

// checks if the CreateSiteDeployRequestFunctionsConfigNameRoutesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSiteDeployRequestFunctionsConfigNameRoutesInner{}

// CreateSiteDeployRequestFunctionsConfigNameRoutesInner struct for CreateSiteDeployRequestFunctionsConfigNameRoutesInner
type CreateSiteDeployRequestFunctionsConfigNameRoutesInner struct {
	Pattern *string `json:"pattern,omitempty"`
	Literal *string `json:"literal,omitempty"`
	Expression *string `json:"expression,omitempty"`
	Methods []string `json:"methods,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateSiteDeployRequestFunctionsConfigNameRoutesInner CreateSiteDeployRequestFunctionsConfigNameRoutesInner

// NewCreateSiteDeployRequestFunctionsConfigNameRoutesInner instantiates a new CreateSiteDeployRequestFunctionsConfigNameRoutesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSiteDeployRequestFunctionsConfigNameRoutesInner() *CreateSiteDeployRequestFunctionsConfigNameRoutesInner {
	this := CreateSiteDeployRequestFunctionsConfigNameRoutesInner{}
	return &this
}

// NewCreateSiteDeployRequestFunctionsConfigNameRoutesInnerWithDefaults instantiates a new CreateSiteDeployRequestFunctionsConfigNameRoutesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSiteDeployRequestFunctionsConfigNameRoutesInnerWithDefaults() *CreateSiteDeployRequestFunctionsConfigNameRoutesInner {
	this := CreateSiteDeployRequestFunctionsConfigNameRoutesInner{}
	return &this
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *CreateSiteDeployRequestFunctionsConfigNameRoutesInner) GetPattern() string {
	if o == nil || IsNil(o.Pattern) {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSiteDeployRequestFunctionsConfigNameRoutesInner) GetPatternOk() (*string, bool) {
	if o == nil || IsNil(o.Pattern) {
		return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *CreateSiteDeployRequestFunctionsConfigNameRoutesInner) HasPattern() bool {
	if o != nil && !IsNil(o.Pattern) {
		return true
	}

	return false
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *CreateSiteDeployRequestFunctionsConfigNameRoutesInner) SetPattern(v string) {
	o.Pattern = &v
}

// GetLiteral returns the Literal field value if set, zero value otherwise.
func (o *CreateSiteDeployRequestFunctionsConfigNameRoutesInner) GetLiteral() string {
	if o == nil || IsNil(o.Literal) {
		var ret string
		return ret
	}
	return *o.Literal
}

// GetLiteralOk returns a tuple with the Literal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSiteDeployRequestFunctionsConfigNameRoutesInner) GetLiteralOk() (*string, bool) {
	if o == nil || IsNil(o.Literal) {
		return nil, false
	}
	return o.Literal, true
}

// HasLiteral returns a boolean if a field has been set.
func (o *CreateSiteDeployRequestFunctionsConfigNameRoutesInner) HasLiteral() bool {
	if o != nil && !IsNil(o.Literal) {
		return true
	}

	return false
}

// SetLiteral gets a reference to the given string and assigns it to the Literal field.
func (o *CreateSiteDeployRequestFunctionsConfigNameRoutesInner) SetLiteral(v string) {
	o.Literal = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *CreateSiteDeployRequestFunctionsConfigNameRoutesInner) GetExpression() string {
	if o == nil || IsNil(o.Expression) {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSiteDeployRequestFunctionsConfigNameRoutesInner) GetExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.Expression) {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *CreateSiteDeployRequestFunctionsConfigNameRoutesInner) HasExpression() bool {
	if o != nil && !IsNil(o.Expression) {
		return true
	}

	return false
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *CreateSiteDeployRequestFunctionsConfigNameRoutesInner) SetExpression(v string) {
	o.Expression = &v
}

// GetMethods returns the Methods field value if set, zero value otherwise.
func (o *CreateSiteDeployRequestFunctionsConfigNameRoutesInner) GetMethods() []string {
	if o == nil || IsNil(o.Methods) {
		var ret []string
		return ret
	}
	return o.Methods
}

// GetMethodsOk returns a tuple with the Methods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSiteDeployRequestFunctionsConfigNameRoutesInner) GetMethodsOk() ([]string, bool) {
	if o == nil || IsNil(o.Methods) {
		return nil, false
	}
	return o.Methods, true
}

// HasMethods returns a boolean if a field has been set.
func (o *CreateSiteDeployRequestFunctionsConfigNameRoutesInner) HasMethods() bool {
	if o != nil && !IsNil(o.Methods) {
		return true
	}

	return false
}

// SetMethods gets a reference to the given []string and assigns it to the Methods field.
func (o *CreateSiteDeployRequestFunctionsConfigNameRoutesInner) SetMethods(v []string) {
	o.Methods = v
}

func (o CreateSiteDeployRequestFunctionsConfigNameRoutesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSiteDeployRequestFunctionsConfigNameRoutesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pattern) {
		toSerialize["pattern"] = o.Pattern
	}
	if !IsNil(o.Literal) {
		toSerialize["literal"] = o.Literal
	}
	if !IsNil(o.Expression) {
		toSerialize["expression"] = o.Expression
	}
	if !IsNil(o.Methods) {
		toSerialize["methods"] = o.Methods
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateSiteDeployRequestFunctionsConfigNameRoutesInner) UnmarshalJSON(data []byte) (err error) {
	varCreateSiteDeployRequestFunctionsConfigNameRoutesInner := _CreateSiteDeployRequestFunctionsConfigNameRoutesInner{}

	err = json.Unmarshal(data, &varCreateSiteDeployRequestFunctionsConfigNameRoutesInner)

	if err != nil {
		return err
	}

	*o = CreateSiteDeployRequestFunctionsConfigNameRoutesInner(varCreateSiteDeployRequestFunctionsConfigNameRoutesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pattern")
		delete(additionalProperties, "literal")
		delete(additionalProperties, "expression")
		delete(additionalProperties, "methods")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateSiteDeployRequestFunctionsConfigNameRoutesInner struct {
	value *CreateSiteDeployRequestFunctionsConfigNameRoutesInner
	isSet bool
}

func (v NullableCreateSiteDeployRequestFunctionsConfigNameRoutesInner) Get() *CreateSiteDeployRequestFunctionsConfigNameRoutesInner {
	return v.value
}

func (v *NullableCreateSiteDeployRequestFunctionsConfigNameRoutesInner) Set(val *CreateSiteDeployRequestFunctionsConfigNameRoutesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSiteDeployRequestFunctionsConfigNameRoutesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSiteDeployRequestFunctionsConfigNameRoutesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSiteDeployRequestFunctionsConfigNameRoutesInner(val *CreateSiteDeployRequestFunctionsConfigNameRoutesInner) *NullableCreateSiteDeployRequestFunctionsConfigNameRoutesInner {
	return &NullableCreateSiteDeployRequestFunctionsConfigNameRoutesInner{value: val, isSet: true}
}

func (v NullableCreateSiteDeployRequestFunctionsConfigNameRoutesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSiteDeployRequestFunctionsConfigNameRoutesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

