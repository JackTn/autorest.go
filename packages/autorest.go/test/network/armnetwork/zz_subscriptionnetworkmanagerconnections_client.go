// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

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

// SubscriptionNetworkManagerConnectionsClient contains the methods for the SubscriptionNetworkManagerConnections group.
// Don't use this type directly, use NewSubscriptionNetworkManagerConnectionsClient() instead.
type SubscriptionNetworkManagerConnectionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSubscriptionNetworkManagerConnectionsClient creates a new instance of SubscriptionNetworkManagerConnectionsClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSubscriptionNetworkManagerConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SubscriptionNetworkManagerConnectionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SubscriptionNetworkManagerConnectionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create a network manager connection on this subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - networkManagerConnectionName - Name for the network manager connection.
//   - parameters - Network manager connection to be created/updated.
//   - options - SubscriptionNetworkManagerConnectionsClientCreateOrUpdateOptions contains the optional parameters for the SubscriptionNetworkManagerConnectionsClient.CreateOrUpdate
//     method.
func (client *SubscriptionNetworkManagerConnectionsClient) CreateOrUpdate(ctx context.Context, networkManagerConnectionName string, parameters ManagerConnection, options *SubscriptionNetworkManagerConnectionsClientCreateOrUpdateOptions) (SubscriptionNetworkManagerConnectionsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "SubscriptionNetworkManagerConnectionsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, networkManagerConnectionName, parameters, options)
	if err != nil {
		return SubscriptionNetworkManagerConnectionsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SubscriptionNetworkManagerConnectionsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return SubscriptionNetworkManagerConnectionsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SubscriptionNetworkManagerConnectionsClient) createOrUpdateCreateRequest(ctx context.Context, networkManagerConnectionName string, parameters ManagerConnection, _ *SubscriptionNetworkManagerConnectionsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkManagerConnections/{networkManagerConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if networkManagerConnectionName == "" {
		return nil, errors.New("parameter networkManagerConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkManagerConnectionName}", url.PathEscape(networkManagerConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *SubscriptionNetworkManagerConnectionsClient) createOrUpdateHandleResponse(resp *http.Response) (SubscriptionNetworkManagerConnectionsClientCreateOrUpdateResponse, error) {
	result := SubscriptionNetworkManagerConnectionsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagerConnection); err != nil {
		return SubscriptionNetworkManagerConnectionsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete specified connection created by this subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - networkManagerConnectionName - Name for the network manager connection.
//   - options - SubscriptionNetworkManagerConnectionsClientDeleteOptions contains the optional parameters for the SubscriptionNetworkManagerConnectionsClient.Delete
//     method.
func (client *SubscriptionNetworkManagerConnectionsClient) Delete(ctx context.Context, networkManagerConnectionName string, options *SubscriptionNetworkManagerConnectionsClientDeleteOptions) (SubscriptionNetworkManagerConnectionsClientDeleteResponse, error) {
	var err error
	const operationName = "SubscriptionNetworkManagerConnectionsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, networkManagerConnectionName, options)
	if err != nil {
		return SubscriptionNetworkManagerConnectionsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SubscriptionNetworkManagerConnectionsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return SubscriptionNetworkManagerConnectionsClientDeleteResponse{}, err
	}
	return SubscriptionNetworkManagerConnectionsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SubscriptionNetworkManagerConnectionsClient) deleteCreateRequest(ctx context.Context, networkManagerConnectionName string, _ *SubscriptionNetworkManagerConnectionsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkManagerConnections/{networkManagerConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if networkManagerConnectionName == "" {
		return nil, errors.New("parameter networkManagerConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkManagerConnectionName}", url.PathEscape(networkManagerConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a specified connection created by this subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - networkManagerConnectionName - Name for the network manager connection.
//   - options - SubscriptionNetworkManagerConnectionsClientGetOptions contains the optional parameters for the SubscriptionNetworkManagerConnectionsClient.Get
//     method.
func (client *SubscriptionNetworkManagerConnectionsClient) Get(ctx context.Context, networkManagerConnectionName string, options *SubscriptionNetworkManagerConnectionsClientGetOptions) (SubscriptionNetworkManagerConnectionsClientGetResponse, error) {
	var err error
	const operationName = "SubscriptionNetworkManagerConnectionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, networkManagerConnectionName, options)
	if err != nil {
		return SubscriptionNetworkManagerConnectionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SubscriptionNetworkManagerConnectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SubscriptionNetworkManagerConnectionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SubscriptionNetworkManagerConnectionsClient) getCreateRequest(ctx context.Context, networkManagerConnectionName string, _ *SubscriptionNetworkManagerConnectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkManagerConnections/{networkManagerConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if networkManagerConnectionName == "" {
		return nil, errors.New("parameter networkManagerConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkManagerConnectionName}", url.PathEscape(networkManagerConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SubscriptionNetworkManagerConnectionsClient) getHandleResponse(resp *http.Response) (SubscriptionNetworkManagerConnectionsClientGetResponse, error) {
	result := SubscriptionNetworkManagerConnectionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagerConnection); err != nil {
		return SubscriptionNetworkManagerConnectionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List all network manager connections created by this subscription.
//
// Generated from API version 2022-09-01
//   - options - SubscriptionNetworkManagerConnectionsClientListOptions contains the optional parameters for the SubscriptionNetworkManagerConnectionsClient.NewListPager
//     method.
func (client *SubscriptionNetworkManagerConnectionsClient) NewListPager(options *SubscriptionNetworkManagerConnectionsClientListOptions) *runtime.Pager[SubscriptionNetworkManagerConnectionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SubscriptionNetworkManagerConnectionsClientListResponse]{
		More: func(page SubscriptionNetworkManagerConnectionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SubscriptionNetworkManagerConnectionsClientListResponse) (SubscriptionNetworkManagerConnectionsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SubscriptionNetworkManagerConnectionsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return SubscriptionNetworkManagerConnectionsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *SubscriptionNetworkManagerConnectionsClient) listCreateRequest(ctx context.Context, options *SubscriptionNetworkManagerConnectionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkManagerConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SubscriptionNetworkManagerConnectionsClient) listHandleResponse(resp *http.Response) (SubscriptionNetworkManagerConnectionsClientListResponse, error) {
	result := SubscriptionNetworkManagerConnectionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagerConnectionListResult); err != nil {
		return SubscriptionNetworkManagerConnectionsClientListResponse{}, err
	}
	return result, nil
}
