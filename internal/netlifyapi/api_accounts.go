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
	"reflect"
)


// AccountsAPIService AccountsAPI service
type AccountsAPIService service

type ApiCancelAccountRequest struct {
	ctx context.Context
	ApiService *AccountsAPIService
	accountId string
}

func (r ApiCancelAccountRequest) Execute() (*http.Response, error) {
	return r.ApiService.CancelAccountExecute(r)
}

/*
CancelAccount Method for CancelAccount

Cancels an account (team).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The ID of the account
 @return ApiCancelAccountRequest
*/
func (a *AccountsAPIService) CancelAccount(ctx context.Context, accountId string) ApiCancelAccountRequest {
	return ApiCancelAccountRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
	}
}

// Execute executes the request
func (a *AccountsAPIService) CancelAccountExecute(r ApiCancelAccountRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsAPIService.CancelAccount")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts/{account_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiCreateAccountRequest struct {
	ctx context.Context
	ApiService *AccountsAPIService
	createAccountRequest *CreateAccountRequest
}

// 
func (r ApiCreateAccountRequest) CreateAccountRequest(createAccountRequest CreateAccountRequest) ApiCreateAccountRequest {
	r.createAccountRequest = &createAccountRequest
	return r
}

func (r ApiCreateAccountRequest) Execute() (*Account, *http.Response, error) {
	return r.ApiService.CreateAccountExecute(r)
}

/*
CreateAccount Method for CreateAccount

Creates an account (team).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateAccountRequest
*/
func (a *AccountsAPIService) CreateAccount(ctx context.Context) ApiCreateAccountRequest {
	return ApiCreateAccountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Account
func (a *AccountsAPIService) CreateAccountExecute(r ApiCreateAccountRequest) (*Account, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Account
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsAPIService.CreateAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createAccountRequest == nil {
		return localVarReturnValue, nil, reportError("createAccountRequest is required and must be specified")
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
	localVarPostBody = r.createAccountRequest
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

type ApiDeleteAccountFirewallRuleSetRequest struct {
	ctx context.Context
	ApiService *AccountsAPIService
	accountId string
}

func (r ApiDeleteAccountFirewallRuleSetRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteAccountFirewallRuleSetExecute(r)
}

/*
DeleteAccountFirewallRuleSet Method for DeleteAccountFirewallRuleSet

remove the rule set for the account if it exists

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The account ID
 @return ApiDeleteAccountFirewallRuleSetRequest
*/
func (a *AccountsAPIService) DeleteAccountFirewallRuleSet(ctx context.Context, accountId string) ApiDeleteAccountFirewallRuleSetRequest {
	return ApiDeleteAccountFirewallRuleSetRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
	}
}

// Execute executes the request
func (a *AccountsAPIService) DeleteAccountFirewallRuleSetExecute(r ApiDeleteAccountFirewallRuleSetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsAPIService.DeleteAccountFirewallRuleSet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts/{account_id}/firewall_rule_set"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiGetAccountRequest struct {
	ctx context.Context
	ApiService *AccountsAPIService
	accountId string
}

func (r ApiGetAccountRequest) Execute() (*Account, *http.Response, error) {
	return r.ApiService.GetAccountExecute(r)
}

/*
GetAccount Method for GetAccount

Returns the details of an account (team).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The ID of the account
 @return ApiGetAccountRequest
*/
func (a *AccountsAPIService) GetAccount(ctx context.Context, accountId string) ApiGetAccountRequest {
	return ApiGetAccountRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
	}
}

// Execute executes the request
//  @return Account
func (a *AccountsAPIService) GetAccountExecute(r ApiGetAccountRequest) (*Account, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Account
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsAPIService.GetAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts/{account_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiGetAccountFirewallRuleSetRequest struct {
	ctx context.Context
	ApiService *AccountsAPIService
	accountId string
	id *string
}

// The account ID
func (r ApiGetAccountFirewallRuleSetRequest) Id(id string) ApiGetAccountFirewallRuleSetRequest {
	r.id = &id
	return r
}

func (r ApiGetAccountFirewallRuleSetRequest) Execute() (*SiteFirewallConfig, *http.Response, error) {
	return r.ApiService.GetAccountFirewallRuleSetExecute(r)
}

/*
GetAccountFirewallRuleSet Method for GetAccountFirewallRuleSet

Returns the firewall rules for an account.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The account ID
 @return ApiGetAccountFirewallRuleSetRequest
*/
func (a *AccountsAPIService) GetAccountFirewallRuleSet(ctx context.Context, accountId string) ApiGetAccountFirewallRuleSetRequest {
	return ApiGetAccountFirewallRuleSetRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
	}
}

// Execute executes the request
//  @return SiteFirewallConfig
func (a *AccountsAPIService) GetAccountFirewallRuleSetExecute(r ApiGetAccountFirewallRuleSetRequest) (*SiteFirewallConfig, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SiteFirewallConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsAPIService.GetAccountFirewallRuleSet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts/{account_id}/firewall_rule_set"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.id == nil {
		return localVarReturnValue, nil, reportError("id is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "")
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

type ApiListAccountAuditEventsRequest struct {
	ctx context.Context
	ApiService *AccountsAPIService
	accountId string
	logType *[]string
}

// Filter logs by one of the accepted types
func (r ApiListAccountAuditEventsRequest) LogType(logType []string) ApiListAccountAuditEventsRequest {
	r.logType = &logType
	return r
}

func (r ApiListAccountAuditEventsRequest) Execute() ([]AuditLog, *http.Response, error) {
	return r.ApiService.ListAccountAuditEventsExecute(r)
}

/*
ListAccountAuditEvents Method for ListAccountAuditEvents

Return an account's audit logs.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The ID of the account
 @return ApiListAccountAuditEventsRequest
*/
func (a *AccountsAPIService) ListAccountAuditEvents(ctx context.Context, accountId string) ApiListAccountAuditEventsRequest {
	return ApiListAccountAuditEventsRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
	}
}

// Execute executes the request
//  @return []AuditLog
func (a *AccountsAPIService) ListAccountAuditEventsExecute(r ApiListAccountAuditEventsRequest) ([]AuditLog, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []AuditLog
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsAPIService.ListAccountAuditEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts/{account_id}/audit"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.logType != nil {
		t := *r.logType
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "log_type", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "log_type", t, "multi")
		}
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

type ApiListAccountTypesForUserRequest struct {
	ctx context.Context
	ApiService *AccountsAPIService
}

func (r ApiListAccountTypesForUserRequest) Execute() (*AccountType, *http.Response, error) {
	return r.ApiService.ListAccountTypesForUserExecute(r)
}

/*
ListAccountTypesForUser Method for ListAccountTypesForUser

Returns a list of account types.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListAccountTypesForUserRequest
*/
func (a *AccountsAPIService) ListAccountTypesForUser(ctx context.Context) ApiListAccountTypesForUserRequest {
	return ApiListAccountTypesForUserRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AccountType
func (a *AccountsAPIService) ListAccountTypesForUserExecute(r ApiListAccountTypesForUserRequest) (*AccountType, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccountType
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsAPIService.ListAccountTypesForUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts/types"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiListAccountsForUserRequest struct {
	ctx context.Context
	ApiService *AccountsAPIService
}

func (r ApiListAccountsForUserRequest) Execute() ([]Account, *http.Response, error) {
	return r.ApiService.ListAccountsForUserExecute(r)
}

/*
ListAccountsForUser Method for ListAccountsForUser

Returns a list of accounts for the current user.
The accounts list excludes accounts that only allow SSO if the current
user is not logged in with SSO. If the user is logged in with SSO, all
accounts are listed.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListAccountsForUserRequest
*/
func (a *AccountsAPIService) ListAccountsForUser(ctx context.Context) ApiListAccountsForUserRequest {
	return ApiListAccountsForUserRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Account
func (a *AccountsAPIService) ListAccountsForUserExecute(r ApiListAccountsForUserRequest) ([]Account, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Account
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsAPIService.ListAccountsForUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiUpdateAccountRequest struct {
	ctx context.Context
	ApiService *AccountsAPIService
	accountId string
	accountUpdateParams *AccountUpdateParams
}

// 
func (r ApiUpdateAccountRequest) AccountUpdateParams(accountUpdateParams AccountUpdateParams) ApiUpdateAccountRequest {
	r.accountUpdateParams = &accountUpdateParams
	return r
}

func (r ApiUpdateAccountRequest) Execute() (*Account, *http.Response, error) {
	return r.ApiService.UpdateAccountExecute(r)
}

/*
UpdateAccount Method for UpdateAccount

Updates an account (team).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The ID of the account
 @return ApiUpdateAccountRequest
*/
func (a *AccountsAPIService) UpdateAccount(ctx context.Context, accountId string) ApiUpdateAccountRequest {
	return ApiUpdateAccountRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
	}
}

// Execute executes the request
//  @return Account
func (a *AccountsAPIService) UpdateAccountExecute(r ApiUpdateAccountRequest) (*Account, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Account
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsAPIService.UpdateAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts/{account_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.accountUpdateParams == nil {
		return localVarReturnValue, nil, reportError("accountUpdateParams is required and must be specified")
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
	localVarPostBody = r.accountUpdateParams
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

type ApiUpdateAccountFirewallRuleSetRequest struct {
	ctx context.Context
	ApiService *AccountsAPIService
	accountId string
	createSiteFirewallConfig *CreateSiteFirewallConfig
}

// 
func (r ApiUpdateAccountFirewallRuleSetRequest) CreateSiteFirewallConfig(createSiteFirewallConfig CreateSiteFirewallConfig) ApiUpdateAccountFirewallRuleSetRequest {
	r.createSiteFirewallConfig = &createSiteFirewallConfig
	return r
}

func (r ApiUpdateAccountFirewallRuleSetRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateAccountFirewallRuleSetExecute(r)
}

/*
UpdateAccountFirewallRuleSet Method for UpdateAccountFirewallRuleSet

Takes a body containing the traffic rule set and completely replaces the full ruleset for the account/team.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The account ID
 @return ApiUpdateAccountFirewallRuleSetRequest
*/
func (a *AccountsAPIService) UpdateAccountFirewallRuleSet(ctx context.Context, accountId string) ApiUpdateAccountFirewallRuleSetRequest {
	return ApiUpdateAccountFirewallRuleSetRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
	}
}

// Execute executes the request
func (a *AccountsAPIService) UpdateAccountFirewallRuleSetExecute(r ApiUpdateAccountFirewallRuleSetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsAPIService.UpdateAccountFirewallRuleSet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts/{account_id}/firewall_rule_set"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createSiteFirewallConfig == nil {
		return nil, reportError("createSiteFirewallConfig is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.createSiteFirewallConfig
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
