//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmachinelearning

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// DatastoresClient contains the methods for the Datastores group.
// Don't use this type directly, use NewDatastoresClient() instead.
type DatastoresClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDatastoresClient creates a new instance of DatastoresClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDatastoresClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DatastoresClient, error) {
	cl, err := arm.NewClient(moduleName+".DatastoresClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DatastoresClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update datastore.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-02-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - name - Datastore name.
//   - body - Datastore entity to create or update.
//   - options - DatastoresClientCreateOrUpdateOptions contains the optional parameters for the DatastoresClient.CreateOrUpdate
//     method.
func (client *DatastoresClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, name string, body DatastoreData, options *DatastoresClientCreateOrUpdateOptions) (DatastoresClientCreateOrUpdateResponse, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, name, body, options)
	if err != nil {
		return DatastoresClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatastoresClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return DatastoresClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DatastoresClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, name string, body DatastoreData, options *DatastoresClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/datastores/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01-preview")
	if options != nil && options.SkipValidation != nil {
		reqQP.Set("skipValidation", strconv.FormatBool(*options.SkipValidation))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DatastoresClient) createOrUpdateHandleResponse(resp *http.Response) (DatastoresClientCreateOrUpdateResponse, error) {
	result := DatastoresClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatastoreData); err != nil {
		return DatastoresClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete datastore.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-02-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - name - Datastore name.
//   - options - DatastoresClientDeleteOptions contains the optional parameters for the DatastoresClient.Delete method.
func (client *DatastoresClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, name string, options *DatastoresClientDeleteOptions) (DatastoresClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, name, options)
	if err != nil {
		return DatastoresClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatastoresClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return DatastoresClientDeleteResponse{}, err
	}
	return DatastoresClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DatastoresClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, name string, options *DatastoresClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/datastores/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get datastore.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-02-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - name - Datastore name.
//   - options - DatastoresClientGetOptions contains the optional parameters for the DatastoresClient.Get method.
func (client *DatastoresClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, name string, options *DatastoresClientGetOptions) (DatastoresClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, name, options)
	if err != nil {
		return DatastoresClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatastoresClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DatastoresClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DatastoresClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, name string, options *DatastoresClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/datastores/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DatastoresClient) getHandleResponse(resp *http.Response) (DatastoresClientGetResponse, error) {
	result := DatastoresClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatastoreData); err != nil {
		return DatastoresClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List datastores.
//
// Generated from API version 2022-02-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - options - DatastoresClientListOptions contains the optional parameters for the DatastoresClient.NewListPager method.
func (client *DatastoresClient) NewListPager(resourceGroupName string, workspaceName string, options *DatastoresClientListOptions) *runtime.Pager[DatastoresClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DatastoresClientListResponse]{
		More: func(page DatastoresClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DatastoresClientListResponse) (DatastoresClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DatastoresClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DatastoresClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DatastoresClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *DatastoresClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *DatastoresClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/datastores"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01-preview")
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", *options.Skip)
	}
	if options != nil && options.Count != nil {
		reqQP.Set("count", strconv.FormatInt(int64(*options.Count), 10))
	}
	if options != nil && options.IsDefault != nil {
		reqQP.Set("isDefault", strconv.FormatBool(*options.IsDefault))
	}
	if options != nil && options.Names != nil {
		reqQP.Set("names", strings.Join(options.Names, ","))
	}
	if options != nil && options.SearchText != nil {
		reqQP.Set("searchText", *options.SearchText)
	}
	if options != nil && options.OrderBy != nil {
		reqQP.Set("orderBy", *options.OrderBy)
	}
	if options != nil && options.OrderByAsc != nil {
		reqQP.Set("orderByAsc", strconv.FormatBool(*options.OrderByAsc))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DatastoresClient) listHandleResponse(resp *http.Response) (DatastoresClientListResponse, error) {
	result := DatastoresClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatastoreResourceArmPaginatedResult); err != nil {
		return DatastoresClientListResponse{}, err
	}
	return result, nil
}

// ListSecrets - Get datastore secrets.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-02-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - name - Datastore name.
//   - options - DatastoresClientListSecretsOptions contains the optional parameters for the DatastoresClient.ListSecrets method.
func (client *DatastoresClient) ListSecrets(ctx context.Context, resourceGroupName string, workspaceName string, name string, options *DatastoresClientListSecretsOptions) (DatastoresClientListSecretsResponse, error) {
	var err error
	req, err := client.listSecretsCreateRequest(ctx, resourceGroupName, workspaceName, name, options)
	if err != nil {
		return DatastoresClientListSecretsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatastoresClientListSecretsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DatastoresClientListSecretsResponse{}, err
	}
	resp, err := client.listSecretsHandleResponse(httpResp)
	return resp, err
}

// listSecretsCreateRequest creates the ListSecrets request.
func (client *DatastoresClient) listSecretsCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, name string, options *DatastoresClientListSecretsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/datastores/{name}/listSecrets"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listSecretsHandleResponse handles the ListSecrets response.
func (client *DatastoresClient) listSecretsHandleResponse(resp *http.Response) (DatastoresClientListSecretsResponse, error) {
	result := DatastoresClientListSecretsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return DatastoresClientListSecretsResponse{}, err
	}
	return result, nil
}
