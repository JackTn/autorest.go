// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armapicenter

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ApiDefinitionsClient contains the methods for the Microsoft.ApiCenter namespace.
// Don't use this type directly, use NewApiDefinitionsClient() instead.
type ApiDefinitionsClient struct {
	internal *arm.Client
}

// NewApiDefinitionsClient creates a new instance of ApiDefinitionsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewApiDefinitionsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ApiDefinitionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ApiDefinitionsClient{
		internal: cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates new or updates existing API definition.
//   - subscriptionID - The ID of the target subscription.
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of Azure API Center service.
//   - workspaceName - The name of the workspace.
//   - apiName - The name of the API.
//   - versionName - The name of the API version.
//   - definitionName - The name of the API definition.
//   - payload - Resource create parameters.
//   - options - ApiDefinitionsClientCreateOrUpdateOptions contains the optional parameters for the ApiDefinitionsClient.CreateOrUpdate
//     method.
func (client *ApiDefinitionsClient) CreateOrUpdate(ctx context.Context, subscriptionID string, resourceGroupName string, serviceName string, workspaceName string, apiName string, versionName string, definitionName string, payload APIDefinition, options *ApiDefinitionsClientCreateOrUpdateOptions) (ApiDefinitionsClientCreateOrUpdateResponse, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, subscriptionID, resourceGroupName, serviceName, workspaceName, apiName, versionName, definitionName, payload, options)
	if err != nil {
		return ApiDefinitionsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ApiDefinitionsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ApiDefinitionsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ApiDefinitionsClient) createOrUpdateCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, serviceName string, workspaceName string, apiName string, versionName string, definitionName string, payload APIDefinition, options *ApiDefinitionsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}/workspaces/{workspaceName}/apis/{apiName}/versions/{versionName}/definitions/{definitionName}"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if apiName == "" {
		return nil, errors.New("parameter apiName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiName}", url.PathEscape(apiName))
	if versionName == "" {
		return nil, errors.New("parameter versionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionName}", url.PathEscape(versionName))
	if definitionName == "" {
		return nil, errors.New("parameter definitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{definitionName}", url.PathEscape(definitionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, payload); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ApiDefinitionsClient) createOrUpdateHandleResponse(resp *http.Response) (ApiDefinitionsClientCreateOrUpdateResponse, error) {
	result := ApiDefinitionsClientCreateOrUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.APIDefinition); err != nil {
		return ApiDefinitionsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes specified API definition.
//   - subscriptionID - The ID of the target subscription.
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of Azure API Center service.
//   - workspaceName - The name of the workspace.
//   - apiName - The name of the API.
//   - versionName - The name of the API version.
//   - definitionName - The name of the API definition.
//   - options - ApiDefinitionsClientDeleteOptions contains the optional parameters for the ApiDefinitionsClient.Delete method.
func (client *ApiDefinitionsClient) Delete(ctx context.Context, subscriptionID string, resourceGroupName string, serviceName string, workspaceName string, apiName string, versionName string, definitionName string, options *ApiDefinitionsClientDeleteOptions) (ApiDefinitionsClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, subscriptionID, resourceGroupName, serviceName, workspaceName, apiName, versionName, definitionName, options)
	if err != nil {
		return ApiDefinitionsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ApiDefinitionsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ApiDefinitionsClientDeleteResponse{}, err
	}
	return ApiDefinitionsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ApiDefinitionsClient) deleteCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, serviceName string, workspaceName string, apiName string, versionName string, definitionName string, options *ApiDefinitionsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}/workspaces/{workspaceName}/apis/{apiName}/versions/{versionName}/definitions/{definitionName}"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if apiName == "" {
		return nil, errors.New("parameter apiName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiName}", url.PathEscape(apiName))
	if versionName == "" {
		return nil, errors.New("parameter versionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionName}", url.PathEscape(versionName))
	if definitionName == "" {
		return nil, errors.New("parameter definitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{definitionName}", url.PathEscape(definitionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginExportSpecification - Exports the API specification.
//   - subscriptionID - The ID of the target subscription.
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of Azure API Center service.
//   - workspaceName - The name of the workspace.
//   - apiName - The name of the API.
//   - versionName - The name of the API version.
//   - definitionName - The name of the API definition.
//   - payload - The content of the action request
//   - options - ApiDefinitionsClientExportSpecificationOptions contains the optional parameters for the ApiDefinitionsClient.ExportSpecification
//     method.
func (client *ApiDefinitionsClient) BeginExportSpecification(ctx context.Context, subscriptionID string, resourceGroupName string, serviceName string, workspaceName string, apiName string, versionName string, definitionName string, payload any, options *ApiDefinitionsClientExportSpecificationOptions) (*runtime.Poller[ApiDefinitionsClientExportSpecificationResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.exportSpecification(ctx, subscriptionID, resourceGroupName, serviceName, workspaceName, apiName, versionName, definitionName, payload, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[ApiDefinitionsClientExportSpecificationResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ApiDefinitionsClientExportSpecificationResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// ExportSpecification - Exports the API specification.
func (client *ApiDefinitionsClient) exportSpecification(ctx context.Context, subscriptionID string, resourceGroupName string, serviceName string, workspaceName string, apiName string, versionName string, definitionName string, payload any, options *ApiDefinitionsClientExportSpecificationOptions) (*http.Response, error) {
	var err error
	req, err := client.exportSpecificationCreateRequest(ctx, subscriptionID, resourceGroupName, serviceName, workspaceName, apiName, versionName, definitionName, payload, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// exportSpecificationCreateRequest creates the ExportSpecification request.
func (client *ApiDefinitionsClient) exportSpecificationCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, serviceName string, workspaceName string, apiName string, versionName string, definitionName string, payload any, options *ApiDefinitionsClientExportSpecificationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}/workspaces/{workspaceName}/apis/{apiName}/versions/{versionName}/definitions/{definitionName}/exportSpecification"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if apiName == "" {
		return nil, errors.New("parameter apiName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiName}", url.PathEscape(apiName))
	if versionName == "" {
		return nil, errors.New("parameter versionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionName}", url.PathEscape(versionName))
	if definitionName == "" {
		return nil, errors.New("parameter definitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{definitionName}", url.PathEscape(definitionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, payload); err != nil {
		return nil, err
	}
	return req, nil
}

// Get - Returns details of the API definition.
//   - subscriptionID - The ID of the target subscription.
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of Azure API Center service.
//   - workspaceName - The name of the workspace.
//   - apiName - The name of the API.
//   - versionName - The name of the API version.
//   - definitionName - The name of the API definition.
//   - options - ApiDefinitionsClientGetOptions contains the optional parameters for the ApiDefinitionsClient.Get method.
func (client *ApiDefinitionsClient) Get(ctx context.Context, subscriptionID string, resourceGroupName string, serviceName string, workspaceName string, apiName string, versionName string, definitionName string, options *ApiDefinitionsClientGetOptions) (ApiDefinitionsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, subscriptionID, resourceGroupName, serviceName, workspaceName, apiName, versionName, definitionName, options)
	if err != nil {
		return ApiDefinitionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ApiDefinitionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ApiDefinitionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ApiDefinitionsClient) getCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, serviceName string, workspaceName string, apiName string, versionName string, definitionName string, options *ApiDefinitionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}/workspaces/{workspaceName}/apis/{apiName}/versions/{versionName}/definitions/{definitionName}"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if apiName == "" {
		return nil, errors.New("parameter apiName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiName}", url.PathEscape(apiName))
	if versionName == "" {
		return nil, errors.New("parameter versionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionName}", url.PathEscape(versionName))
	if definitionName == "" {
		return nil, errors.New("parameter definitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{definitionName}", url.PathEscape(definitionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ApiDefinitionsClient) getHandleResponse(resp *http.Response) (ApiDefinitionsClientGetResponse, error) {
	result := ApiDefinitionsClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.APIDefinition); err != nil {
		return ApiDefinitionsClientGetResponse{}, err
	}
	return result, nil
}

// Head - Checks if specified API definition exists.
//   - subscriptionID - The ID of the target subscription.
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of Azure API Center service.
//   - workspaceName - The name of the workspace.
//   - apiName - The name of the API.
//   - versionName - The name of the API version.
//   - definitionName - The name of the API definition.
//   - options - ApiDefinitionsClientHeadOptions contains the optional parameters for the ApiDefinitionsClient.Head method.
func (client *ApiDefinitionsClient) Head(ctx context.Context, subscriptionID string, resourceGroupName string, serviceName string, workspaceName string, apiName string, versionName string, definitionName string, options *ApiDefinitionsClientHeadOptions) (ApiDefinitionsClientHeadResponse, error) {
	var err error
	req, err := client.headCreateRequest(ctx, subscriptionID, resourceGroupName, serviceName, workspaceName, apiName, versionName, definitionName, options)
	if err != nil {
		return ApiDefinitionsClientHeadResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ApiDefinitionsClientHeadResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ApiDefinitionsClientHeadResponse{}, err
	}
	return ApiDefinitionsClientHeadResponse{}, nil
}

// headCreateRequest creates the Head request.
func (client *ApiDefinitionsClient) headCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, serviceName string, workspaceName string, apiName string, versionName string, definitionName string, options *ApiDefinitionsClientHeadOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}/workspaces/{workspaceName}/apis/{apiName}/versions/{versionName}/definitions/{definitionName}"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if apiName == "" {
		return nil, errors.New("parameter apiName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiName}", url.PathEscape(apiName))
	if versionName == "" {
		return nil, errors.New("parameter versionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionName}", url.PathEscape(versionName))
	if definitionName == "" {
		return nil, errors.New("parameter definitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{definitionName}", url.PathEscape(definitionName))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginImportSpecification - Imports the API specification.
//   - subscriptionID - The ID of the target subscription.
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of Azure API Center service.
//   - workspaceName - The name of the workspace.
//   - apiName - The name of the API.
//   - versionName - The name of the API version.
//   - definitionName - The name of the API definition.
//   - payload - The content of the action request
//   - options - ApiDefinitionsClientImportSpecificationOptions contains the optional parameters for the ApiDefinitionsClient.ImportSpecification
//     method.
func (client *ApiDefinitionsClient) BeginImportSpecification(ctx context.Context, subscriptionID string, resourceGroupName string, serviceName string, workspaceName string, apiName string, versionName string, definitionName string, payload APISpecImportRequest, options *ApiDefinitionsClientImportSpecificationOptions) (*runtime.Poller[ApiDefinitionsClientImportSpecificationResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.importSpecification(ctx, subscriptionID, resourceGroupName, serviceName, workspaceName, apiName, versionName, definitionName, payload, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[ApiDefinitionsClientImportSpecificationResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ApiDefinitionsClientImportSpecificationResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// ImportSpecification - Imports the API specification.
func (client *ApiDefinitionsClient) importSpecification(ctx context.Context, subscriptionID string, resourceGroupName string, serviceName string, workspaceName string, apiName string, versionName string, definitionName string, payload APISpecImportRequest, options *ApiDefinitionsClientImportSpecificationOptions) (*http.Response, error) {
	var err error
	req, err := client.importSpecificationCreateRequest(ctx, subscriptionID, resourceGroupName, serviceName, workspaceName, apiName, versionName, definitionName, payload, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// importSpecificationCreateRequest creates the ImportSpecification request.
func (client *ApiDefinitionsClient) importSpecificationCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, serviceName string, workspaceName string, apiName string, versionName string, definitionName string, payload APISpecImportRequest, options *ApiDefinitionsClientImportSpecificationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}/workspaces/{workspaceName}/apis/{apiName}/versions/{versionName}/definitions/{definitionName}/importSpecification"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if apiName == "" {
		return nil, errors.New("parameter apiName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiName}", url.PathEscape(apiName))
	if versionName == "" {
		return nil, errors.New("parameter versionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionName}", url.PathEscape(versionName))
	if definitionName == "" {
		return nil, errors.New("parameter definitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{definitionName}", url.PathEscape(definitionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, payload); err != nil {
		return nil, err
	}
	return req, nil
}

// NewListPager - Returns a collection of API definitions.
//   - subscriptionID - The ID of the target subscription.
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of Azure API Center service.
//   - workspaceName - The name of the workspace.
//   - apiName - The name of the API.
//   - versionName - The name of the API version.
//   - options - ApiDefinitionsClientListOptions contains the optional parameters for the ApiDefinitionsClient.NewListPager method.
func (client *ApiDefinitionsClient) NewListPager(subscriptionID string, resourceGroupName string, serviceName string, workspaceName string, apiName string, versionName string, options *ApiDefinitionsClientListOptions) *runtime.Pager[ApiDefinitionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ApiDefinitionsClientListResponse]{
		More: func(page ApiDefinitionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ApiDefinitionsClientListResponse) (ApiDefinitionsClientListResponse, error) {
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, subscriptionID, resourceGroupName, serviceName, workspaceName, apiName, versionName, options)
			}, nil)
			if err != nil {
				return ApiDefinitionsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ApiDefinitionsClient) listCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, serviceName string, workspaceName string, apiName string, versionName string, options *ApiDefinitionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}/workspaces/{workspaceName}/apis/{apiName}/versions/{versionName}/definitions"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if apiName == "" {
		return nil, errors.New("parameter apiName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiName}", url.PathEscape(apiName))
	if versionName == "" {
		return nil, errors.New("parameter versionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionName}", url.PathEscape(versionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ApiDefinitionsClient) listHandleResponse(resp *http.Response) (ApiDefinitionsClientListResponse, error) {
	result := ApiDefinitionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.APIDefinitionListResult); err != nil {
		return ApiDefinitionsClientListResponse{}, err
	}
	return result, nil
}
