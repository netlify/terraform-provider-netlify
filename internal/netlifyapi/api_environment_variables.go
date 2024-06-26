/*
Netlify's API documentation

Netlify is a hosting service for the programmable web. It understands your documents and provides an API to handle atomic deploys of websites, manage form submissions, inject JavaScript snippets, and much more. This is a REST-style API that uses JSON for serialization and OAuth 2 for authentication.   This document is an OpenAPI reference for the Netlify API that you can explore. For more detailed instructions for common uses, please visit the [online documentation](https://docs.netlify.com/api/get-started/). Visit our Community forum to join the conversation about [understanding and using Netlify’s API](https://community.netlify.com/t/common-issue-understanding-and-using-netlifys-api/160).   Additionally, we have two API clients for your convenience: - [Go Client](https://github.com/netlify/open-api#go-client) - [JS Client](https://github.com/netlify/js-client) 

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netlifyapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// EnvironmentVariablesAPIService EnvironmentVariablesAPI service
type EnvironmentVariablesAPIService service

type ApiCreateEnvVarsRequest struct {
	ctx context.Context
	ApiService *EnvironmentVariablesAPIService
	accountId string
	envVar *[]EnvVar
	client *string
	siteId *string
}

// The array of environment variables to create
func (r ApiCreateEnvVarsRequest) EnvVar(envVar []EnvVar) ApiCreateEnvVarsRequest {
	r.envVar = &envVar
	return r
}

// Optional parameter that identifies where the request is coming from (added originally for telemetry purposes)
func (r ApiCreateEnvVarsRequest) Client(client string) ApiCreateEnvVarsRequest {
	r.client = &client
	return r
}

// If provided, create an environment variable on the site level, not the account level
func (r ApiCreateEnvVarsRequest) SiteId(siteId string) ApiCreateEnvVarsRequest {
	r.siteId = &siteId
	return r
}

func (r ApiCreateEnvVarsRequest) Execute() ([]EnvVar, *http.Response, error) {
	return r.ApiService.CreateEnvVarsExecute(r)
}

/*
CreateEnvVars Method for CreateEnvVars

Creates new environment variables. Granular scopes are available on Pro plans and above.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId Scope response to account_id
 @return ApiCreateEnvVarsRequest
*/
func (a *EnvironmentVariablesAPIService) CreateEnvVars(ctx context.Context, accountId string) ApiCreateEnvVarsRequest {
	return ApiCreateEnvVarsRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
	}
}

// Execute executes the request
//  @return []EnvVar
func (a *EnvironmentVariablesAPIService) CreateEnvVarsExecute(r ApiCreateEnvVarsRequest) ([]EnvVar, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []EnvVar
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentVariablesAPIService.CreateEnvVars")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts/{account_id}/env"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.envVar == nil {
		return localVarReturnValue, nil, reportError("envVar is required and must be specified")
	}

	if r.client != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "client", r.client, "")
	}
	if r.siteId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "site_id", r.siteId, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.envVar
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteEnvVarRequest struct {
	ctx context.Context
	ApiService *EnvironmentVariablesAPIService
	accountId string
	envKey string
	siteId *string
}

// If provided, delete the environment variable from this site
func (r ApiDeleteEnvVarRequest) SiteId(siteId string) ApiDeleteEnvVarRequest {
	r.siteId = &siteId
	return r
}

func (r ApiDeleteEnvVarRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteEnvVarExecute(r)
}

/*
DeleteEnvVar Method for DeleteEnvVar

Deletes an environment variable.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId Scope response to account_id
 @param envKey The environment variable key (case-sensitive)
 @return ApiDeleteEnvVarRequest
*/
func (a *EnvironmentVariablesAPIService) DeleteEnvVar(ctx context.Context, accountId string, envKey string) ApiDeleteEnvVarRequest {
	return ApiDeleteEnvVarRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		envKey: envKey,
	}
}

// Execute executes the request
func (a *EnvironmentVariablesAPIService) DeleteEnvVarExecute(r ApiDeleteEnvVarRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentVariablesAPIService.DeleteEnvVar")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts/{account_id}/env/{env_key}"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"env_key"+"}", url.PathEscape(parameterValueToString(r.envKey, "envKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.siteId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "site_id", r.siteId, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeleteEnvVarValueRequest struct {
	ctx context.Context
	ApiService *EnvironmentVariablesAPIService
	accountId string
	envKey string
	id string
	siteId *string
}

// If provided, delete the value from an environment variable on this site
func (r ApiDeleteEnvVarValueRequest) SiteId(siteId string) ApiDeleteEnvVarValueRequest {
	r.siteId = &siteId
	return r
}

func (r ApiDeleteEnvVarValueRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteEnvVarValueExecute(r)
}

/*
DeleteEnvVarValue Method for DeleteEnvVarValue

Deletes a specific environment variable value.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId Scope response to account_id
 @param envKey The environment variable key name (case-sensitive)
 @param id The environment variable value's ID
 @return ApiDeleteEnvVarValueRequest
*/
func (a *EnvironmentVariablesAPIService) DeleteEnvVarValue(ctx context.Context, accountId string, envKey string, id string) ApiDeleteEnvVarValueRequest {
	return ApiDeleteEnvVarValueRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		envKey: envKey,
		id: id,
	}
}

// Execute executes the request
func (a *EnvironmentVariablesAPIService) DeleteEnvVarValueExecute(r ApiDeleteEnvVarValueRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentVariablesAPIService.DeleteEnvVarValue")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts/{account_id}/env/{env_key}/value/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"env_key"+"}", url.PathEscape(parameterValueToString(r.envKey, "envKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.siteId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "site_id", r.siteId, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetEnvVarRequest struct {
	ctx context.Context
	ApiService *EnvironmentVariablesAPIService
	accountId string
	envKey string
	siteId *string
}

// If provided, return the environment variable for a specific site (no merging is performed)
func (r ApiGetEnvVarRequest) SiteId(siteId string) ApiGetEnvVarRequest {
	r.siteId = &siteId
	return r
}

func (r ApiGetEnvVarRequest) Execute() (*EnvVar, *http.Response, error) {
	return r.ApiService.GetEnvVarExecute(r)
}

/*
GetEnvVar Method for GetEnvVar

Returns an individual environment variable.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId Scope response to account_id
 @param envKey The environment variable key (case-sensitive)
 @return ApiGetEnvVarRequest
*/
func (a *EnvironmentVariablesAPIService) GetEnvVar(ctx context.Context, accountId string, envKey string) ApiGetEnvVarRequest {
	return ApiGetEnvVarRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		envKey: envKey,
	}
}

// Execute executes the request
//  @return EnvVar
func (a *EnvironmentVariablesAPIService) GetEnvVarExecute(r ApiGetEnvVarRequest) (*EnvVar, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EnvVar
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentVariablesAPIService.GetEnvVar")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts/{account_id}/env/{env_key}"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"env_key"+"}", url.PathEscape(parameterValueToString(r.envKey, "envKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.siteId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "site_id", r.siteId, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetEnvVarsRequest struct {
	ctx context.Context
	ApiService *EnvironmentVariablesAPIService
	accountId string
	context *string
	scope *string
	siteId *string
}

// Filter by deploy context
func (r ApiGetEnvVarsRequest) Context(context string) ApiGetEnvVarsRequest {
	r.context = &context
	return r
}

// Filter by scope
func (r ApiGetEnvVarsRequest) Scope(scope string) ApiGetEnvVarsRequest {
	r.scope = &scope
	return r
}

// If specified, only return environment variables set on this site
func (r ApiGetEnvVarsRequest) SiteId(siteId string) ApiGetEnvVarsRequest {
	r.siteId = &siteId
	return r
}

func (r ApiGetEnvVarsRequest) Execute() ([]EnvVar, *http.Response, error) {
	return r.ApiService.GetEnvVarsExecute(r)
}

/*
GetEnvVars Method for GetEnvVars

Returns all environment variables for an account or site. An account corresponds to a team in the Netlify UI.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId Scope response to account_id
 @return ApiGetEnvVarsRequest
*/
func (a *EnvironmentVariablesAPIService) GetEnvVars(ctx context.Context, accountId string) ApiGetEnvVarsRequest {
	return ApiGetEnvVarsRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
	}
}

// Execute executes the request
//  @return []EnvVar
func (a *EnvironmentVariablesAPIService) GetEnvVarsExecute(r ApiGetEnvVarsRequest) ([]EnvVar, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []EnvVar
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentVariablesAPIService.GetEnvVars")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts/{account_id}/env"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.context != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "context", r.context, "")
	}
	if r.scope != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "scope", r.scope, "")
	}
	if r.siteId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "site_id", r.siteId, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetSiteEnvVarsRequest struct {
	ctx context.Context
	ApiService *EnvironmentVariablesAPIService
	siteId string
	context *string
	scope *string
}

// Filter by deploy context
func (r ApiGetSiteEnvVarsRequest) Context(context string) ApiGetSiteEnvVarsRequest {
	r.context = &context
	return r
}

// Filter by scope
func (r ApiGetSiteEnvVarsRequest) Scope(scope string) ApiGetSiteEnvVarsRequest {
	r.scope = &scope
	return r
}

func (r ApiGetSiteEnvVarsRequest) Execute() ([]EnvVar, *http.Response, error) {
	return r.ApiService.GetSiteEnvVarsExecute(r)
}

/*
GetSiteEnvVars Method for GetSiteEnvVars

Returns all environment variables for a site.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param siteId Scope response to site_id
 @return ApiGetSiteEnvVarsRequest
*/
func (a *EnvironmentVariablesAPIService) GetSiteEnvVars(ctx context.Context, siteId string) ApiGetSiteEnvVarsRequest {
	return ApiGetSiteEnvVarsRequest{
		ApiService: a,
		ctx: ctx,
		siteId: siteId,
	}
}

// Execute executes the request
//  @return []EnvVar
func (a *EnvironmentVariablesAPIService) GetSiteEnvVarsExecute(r ApiGetSiteEnvVarsRequest) ([]EnvVar, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []EnvVar
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentVariablesAPIService.GetSiteEnvVars")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/sites/{site_id}/env"
	localVarPath = strings.Replace(localVarPath, "{"+"site_id"+"}", url.PathEscape(parameterValueToString(r.siteId, "siteId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.context != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "context", r.context, "")
	}
	if r.scope != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "scope", r.scope, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSetEnvVarValueRequest struct {
	ctx context.Context
	ApiService *EnvironmentVariablesAPIService
	accountId string
	envKey string
	envVarSet *EnvVarSet
	siteId *string
}

// 
func (r ApiSetEnvVarValueRequest) EnvVarSet(envVarSet EnvVarSet) ApiSetEnvVarValueRequest {
	r.envVarSet = &envVarSet
	return r
}

// If provided, update an environment variable set on this site
func (r ApiSetEnvVarValueRequest) SiteId(siteId string) ApiSetEnvVarValueRequest {
	r.siteId = &siteId
	return r
}

func (r ApiSetEnvVarValueRequest) Execute() (*EnvVar, *http.Response, error) {
	return r.ApiService.SetEnvVarValueExecute(r)
}

/*
SetEnvVarValue Method for SetEnvVarValue

(SET) Updates or creates a value on an environment variable.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId Scope response to account_id
 @param envKey The existing environment variable key name (case-sensitive)
 @return ApiSetEnvVarValueRequest
*/
func (a *EnvironmentVariablesAPIService) SetEnvVarValue(ctx context.Context, accountId string, envKey string) ApiSetEnvVarValueRequest {
	return ApiSetEnvVarValueRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		envKey: envKey,
	}
}

// Execute executes the request
//  @return EnvVar
func (a *EnvironmentVariablesAPIService) SetEnvVarValueExecute(r ApiSetEnvVarValueRequest) (*EnvVar, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EnvVar
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentVariablesAPIService.SetEnvVarValue")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts/{account_id}/env/{env_key}"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"env_key"+"}", url.PathEscape(parameterValueToString(r.envKey, "envKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.envVarSet == nil {
		return localVarReturnValue, nil, reportError("envVarSet is required and must be specified")
	}

	if r.siteId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "site_id", r.siteId, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.envVarSet
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateEnvVarRequest struct {
	ctx context.Context
	ApiService *EnvironmentVariablesAPIService
	accountId string
	envKey string
	key *string
	updateEnvVarRequest *UpdateEnvVarRequest
	siteId *string
}

// The existing or new name of the key, if you wish to rename it (case-sensitive)
func (r ApiUpdateEnvVarRequest) Key(key string) ApiUpdateEnvVarRequest {
	r.key = &key
	return r
}

// 
func (r ApiUpdateEnvVarRequest) UpdateEnvVarRequest(updateEnvVarRequest UpdateEnvVarRequest) ApiUpdateEnvVarRequest {
	r.updateEnvVarRequest = &updateEnvVarRequest
	return r
}

// If provided, update an environment variable set on this site
func (r ApiUpdateEnvVarRequest) SiteId(siteId string) ApiUpdateEnvVarRequest {
	r.siteId = &siteId
	return r
}

func (r ApiUpdateEnvVarRequest) Execute() (*EnvVar, *http.Response, error) {
	return r.ApiService.UpdateEnvVarExecute(r)
}

/*
UpdateEnvVar Method for UpdateEnvVar

Updates an existing environment variable and all of its values. Existing values will be replaced by values provided.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId Scope response to account_id
 @param envKey The existing environment variable key name (case-sensitive)
 @return ApiUpdateEnvVarRequest
*/
func (a *EnvironmentVariablesAPIService) UpdateEnvVar(ctx context.Context, accountId string, envKey string) ApiUpdateEnvVarRequest {
	return ApiUpdateEnvVarRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		envKey: envKey,
	}
}

// Execute executes the request
//  @return EnvVar
func (a *EnvironmentVariablesAPIService) UpdateEnvVarExecute(r ApiUpdateEnvVarRequest) (*EnvVar, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EnvVar
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentVariablesAPIService.UpdateEnvVar")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts/{account_id}/env/{env_key}"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"env_key"+"}", url.PathEscape(parameterValueToString(r.envKey, "envKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.key == nil {
		return localVarReturnValue, nil, reportError("key is required and must be specified")
	}
	if r.updateEnvVarRequest == nil {
		return localVarReturnValue, nil, reportError("updateEnvVarRequest is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "key", r.key, "")
	if r.siteId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "site_id", r.siteId, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateEnvVarRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
