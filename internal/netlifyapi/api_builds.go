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


// BuildsAPIService BuildsAPI service
type BuildsAPIService service

type ApiCreateSiteBuildRequest struct {
	ctx context.Context
	ApiService *BuildsAPIService
	siteId string
	branch *string
	clearCache *bool
	image *string
	templateId *string
}

// The branch to build; defaults to main branch
func (r ApiCreateSiteBuildRequest) Branch(branch string) ApiCreateSiteBuildRequest {
	r.branch = &branch
	return r
}

// Whether to clear the build cache before building
func (r ApiCreateSiteBuildRequest) ClearCache(clearCache bool) ApiCreateSiteBuildRequest {
	r.clearCache = &clearCache
	return r
}

// The build image tag to use for the build
func (r ApiCreateSiteBuildRequest) Image(image string) ApiCreateSiteBuildRequest {
	r.image = &image
	return r
}

// The build template to use for the build
func (r ApiCreateSiteBuildRequest) TemplateId(templateId string) ApiCreateSiteBuildRequest {
	r.templateId = &templateId
	return r
}

func (r ApiCreateSiteBuildRequest) Execute() (*Build, *http.Response, error) {
	return r.ApiService.CreateSiteBuildExecute(r)
}

/*
CreateSiteBuild Method for CreateSiteBuild

Runs a build for a site. The build will be scheduled to run at the first opportunity, but it might not start immediately if insufficient account build capacity is available.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param siteId The ID of the site
 @return ApiCreateSiteBuildRequest
*/
func (a *BuildsAPIService) CreateSiteBuild(ctx context.Context, siteId string) ApiCreateSiteBuildRequest {
	return ApiCreateSiteBuildRequest{
		ApiService: a,
		ctx: ctx,
		siteId: siteId,
	}
}

// Execute executes the request
//  @return Build
func (a *BuildsAPIService) CreateSiteBuildExecute(r ApiCreateSiteBuildRequest) (*Build, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Build
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BuildsAPIService.CreateSiteBuild")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/sites/{site_id}/builds"
	localVarPath = strings.Replace(localVarPath, "{"+"site_id"+"}", url.PathEscape(parameterValueToString(r.siteId, "siteId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.branch != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "branch", r.branch, "")
	}
	if r.clearCache != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "clear_cache", r.clearCache, "")
	}
	if r.image != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "image", r.image, "")
	}
	if r.templateId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "template_id", r.templateId, "")
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

type ApiGetAccountBuildStatusRequest struct {
	ctx context.Context
	ApiService *BuildsAPIService
	accountId string
	siteId *string
	state *string
}

// site_id
func (r ApiGetAccountBuildStatusRequest) SiteId(siteId string) ApiGetAccountBuildStatusRequest {
	r.siteId = &siteId
	return r
}

// state
func (r ApiGetAccountBuildStatusRequest) State(state string) ApiGetAccountBuildStatusRequest {
	r.state = &state
	return r
}

func (r ApiGetAccountBuildStatusRequest) Execute() ([]AccountBuild, *http.Response, error) {
	return r.ApiService.GetAccountBuildStatusExecute(r)
}

/*
GetAccountBuildStatus Method for GetAccountBuildStatus

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId account_id
 @return ApiGetAccountBuildStatusRequest
*/
func (a *BuildsAPIService) GetAccountBuildStatus(ctx context.Context, accountId string) ApiGetAccountBuildStatusRequest {
	return ApiGetAccountBuildStatusRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
	}
}

// Execute executes the request
//  @return []AccountBuild
func (a *BuildsAPIService) GetAccountBuildStatusExecute(r ApiGetAccountBuildStatusRequest) ([]AccountBuild, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []AccountBuild
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BuildsAPIService.GetAccountBuildStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts/{account_id}/builds/status"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.siteId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "site_id", r.siteId, "")
	}
	if r.state != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "state", r.state, "")
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

type ApiGetSiteBuildRequest struct {
	ctx context.Context
	ApiService *BuildsAPIService
	buildId string
}

func (r ApiGetSiteBuildRequest) Execute() (*Build, *http.Response, error) {
	return r.ApiService.GetSiteBuildExecute(r)
}

/*
GetSiteBuild Method for GetSiteBuild

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param buildId build_id
 @return ApiGetSiteBuildRequest
*/
func (a *BuildsAPIService) GetSiteBuild(ctx context.Context, buildId string) ApiGetSiteBuildRequest {
	return ApiGetSiteBuildRequest{
		ApiService: a,
		ctx: ctx,
		buildId: buildId,
	}
}

// Execute executes the request
//  @return Build
func (a *BuildsAPIService) GetSiteBuildExecute(r ApiGetSiteBuildRequest) (*Build, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Build
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BuildsAPIService.GetSiteBuild")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/builds/{build_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"build_id"+"}", url.PathEscape(parameterValueToString(r.buildId, "buildId")), -1)

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

type ApiListSiteBuildsRequest struct {
	ctx context.Context
	ApiService *BuildsAPIService
	siteId string
}

func (r ApiListSiteBuildsRequest) Execute() ([]Build, *http.Response, error) {
	return r.ApiService.ListSiteBuildsExecute(r)
}

/*
ListSiteBuilds Method for ListSiteBuilds

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param siteId The ID of the site
 @return ApiListSiteBuildsRequest
*/
func (a *BuildsAPIService) ListSiteBuilds(ctx context.Context, siteId string) ApiListSiteBuildsRequest {
	return ApiListSiteBuildsRequest{
		ApiService: a,
		ctx: ctx,
		siteId: siteId,
	}
}

// Execute executes the request
//  @return []Build
func (a *BuildsAPIService) ListSiteBuildsExecute(r ApiListSiteBuildsRequest) ([]Build, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Build
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BuildsAPIService.ListSiteBuilds")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sites/{site_id}/builds"
	localVarPath = strings.Replace(localVarPath, "{"+"site_id"+"}", url.PathEscape(parameterValueToString(r.siteId, "siteId")), -1)

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

type ApiNotifyBuildStartRequest struct {
	ctx context.Context
	ApiService *BuildsAPIService
	buildId string
	buildVersion *string
	buildbotVersion *string
}

// build_version
func (r ApiNotifyBuildStartRequest) BuildVersion(buildVersion string) ApiNotifyBuildStartRequest {
	r.buildVersion = &buildVersion
	return r
}

// buildbot_version
func (r ApiNotifyBuildStartRequest) BuildbotVersion(buildbotVersion string) ApiNotifyBuildStartRequest {
	r.buildbotVersion = &buildbotVersion
	return r
}

func (r ApiNotifyBuildStartRequest) Execute() (*http.Response, error) {
	return r.ApiService.NotifyBuildStartExecute(r)
}

/*
NotifyBuildStart Method for NotifyBuildStart

Mark that the build has started.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param buildId build_id
 @return ApiNotifyBuildStartRequest
*/
func (a *BuildsAPIService) NotifyBuildStart(ctx context.Context, buildId string) ApiNotifyBuildStartRequest {
	return ApiNotifyBuildStartRequest{
		ApiService: a,
		ctx: ctx,
		buildId: buildId,
	}
}

// Execute executes the request
func (a *BuildsAPIService) NotifyBuildStartExecute(r ApiNotifyBuildStartRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BuildsAPIService.NotifyBuildStart")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/builds/{build_id}/start"
	localVarPath = strings.Replace(localVarPath, "{"+"build_id"+"}", url.PathEscape(parameterValueToString(r.buildId, "buildId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.buildVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "build_version", r.buildVersion, "")
	}
	if r.buildbotVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "buildbot_version", r.buildbotVersion, "")
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

type ApiUpdateSiteBuildLogRequest struct {
	ctx context.Context
	ApiService *BuildsAPIService
	buildId string
	updateSiteBuildLogRequest *UpdateSiteBuildLogRequest
}

// 
func (r ApiUpdateSiteBuildLogRequest) UpdateSiteBuildLogRequest(updateSiteBuildLogRequest UpdateSiteBuildLogRequest) ApiUpdateSiteBuildLogRequest {
	r.updateSiteBuildLogRequest = &updateSiteBuildLogRequest
	return r
}

func (r ApiUpdateSiteBuildLogRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateSiteBuildLogExecute(r)
}

/*
UpdateSiteBuildLog Method for UpdateSiteBuildLog

Add a message to the build log.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param buildId build_id
 @return ApiUpdateSiteBuildLogRequest
*/
func (a *BuildsAPIService) UpdateSiteBuildLog(ctx context.Context, buildId string) ApiUpdateSiteBuildLogRequest {
	return ApiUpdateSiteBuildLogRequest{
		ApiService: a,
		ctx: ctx,
		buildId: buildId,
	}
}

// Execute executes the request
func (a *BuildsAPIService) UpdateSiteBuildLogExecute(r ApiUpdateSiteBuildLogRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BuildsAPIService.UpdateSiteBuildLog")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/builds/{build_id}/log"
	localVarPath = strings.Replace(localVarPath, "{"+"build_id"+"}", url.PathEscape(parameterValueToString(r.buildId, "buildId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateSiteBuildLogRequest == nil {
		return nil, reportError("updateSiteBuildLogRequest is required and must be specified")
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
	localVarPostBody = r.updateSiteBuildLogRequest
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
		if localVarHTTPResponse.StatusCode == 422 {
			var v UpdateSiteBuildLog422Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
