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
	"strings"
)

// ExpressRouteCrossConnectionsClient contains the methods for the ExpressRouteCrossConnections group.
// Don't use this type directly, use NewExpressRouteCrossConnectionsClient() instead.
type ExpressRouteCrossConnectionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewExpressRouteCrossConnectionsClient creates a new instance of ExpressRouteCrossConnectionsClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewExpressRouteCrossConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ExpressRouteCrossConnectionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ExpressRouteCrossConnectionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Update the specified ExpressRouteCrossConnection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - crossConnectionName - The name of the ExpressRouteCrossConnection.
//   - parameters - Parameters supplied to the update express route crossConnection operation.
//   - options - ExpressRouteCrossConnectionsClientBeginCreateOrUpdateOptions contains the optional parameters for the ExpressRouteCrossConnectionsClient.BeginCreateOrUpdate
//     method.
func (client *ExpressRouteCrossConnectionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, crossConnectionName string, parameters ExpressRouteCrossConnection, options *ExpressRouteCrossConnectionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ExpressRouteCrossConnectionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, crossConnectionName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ExpressRouteCrossConnectionsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ExpressRouteCrossConnectionsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Update the specified ExpressRouteCrossConnection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *ExpressRouteCrossConnectionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, crossConnectionName string, parameters ExpressRouteCrossConnection, options *ExpressRouteCrossConnectionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ExpressRouteCrossConnectionsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, crossConnectionName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ExpressRouteCrossConnectionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, parameters ExpressRouteCrossConnection, options *ExpressRouteCrossConnectionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if crossConnectionName == "" {
		return nil, errors.New("parameter crossConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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

// Get - Gets details about the specified ExpressRouteCrossConnection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group (peering location of the circuit).
//   - crossConnectionName - The name of the ExpressRouteCrossConnection (service key of the circuit).
//   - options - ExpressRouteCrossConnectionsClientGetOptions contains the optional parameters for the ExpressRouteCrossConnectionsClient.Get
//     method.
func (client *ExpressRouteCrossConnectionsClient) Get(ctx context.Context, resourceGroupName string, crossConnectionName string, options *ExpressRouteCrossConnectionsClientGetOptions) (ExpressRouteCrossConnectionsClientGetResponse, error) {
	var err error
	const operationName = "ExpressRouteCrossConnectionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, crossConnectionName, options)
	if err != nil {
		return ExpressRouteCrossConnectionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExpressRouteCrossConnectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ExpressRouteCrossConnectionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ExpressRouteCrossConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, options *ExpressRouteCrossConnectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if crossConnectionName == "" {
		return nil, errors.New("parameter crossConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
func (client *ExpressRouteCrossConnectionsClient) getHandleResponse(resp *http.Response) (ExpressRouteCrossConnectionsClientGetResponse, error) {
	result := ExpressRouteCrossConnectionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteCrossConnection); err != nil {
		return ExpressRouteCrossConnectionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Retrieves all the ExpressRouteCrossConnections in a subscription.
//
// Generated from API version 2022-09-01
//   - options - ExpressRouteCrossConnectionsClientListOptions contains the optional parameters for the ExpressRouteCrossConnectionsClient.NewListPager
//     method.
func (client *ExpressRouteCrossConnectionsClient) NewListPager(options *ExpressRouteCrossConnectionsClientListOptions) *runtime.Pager[ExpressRouteCrossConnectionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ExpressRouteCrossConnectionsClientListResponse]{
		More: func(page ExpressRouteCrossConnectionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ExpressRouteCrossConnectionsClientListResponse) (ExpressRouteCrossConnectionsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ExpressRouteCrossConnectionsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return ExpressRouteCrossConnectionsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ExpressRouteCrossConnectionsClient) listCreateRequest(ctx context.Context, options *ExpressRouteCrossConnectionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteCrossConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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

// listHandleResponse handles the List response.
func (client *ExpressRouteCrossConnectionsClient) listHandleResponse(resp *http.Response) (ExpressRouteCrossConnectionsClientListResponse, error) {
	result := ExpressRouteCrossConnectionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteCrossConnectionListResult); err != nil {
		return ExpressRouteCrossConnectionsClientListResponse{}, err
	}
	return result, nil
}

// BeginListArpTable - Gets the currently advertised ARP table associated with the express route cross connection in a resource
// group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - crossConnectionName - The name of the ExpressRouteCrossConnection.
//   - peeringName - The name of the peering.
//   - devicePath - The path of the device.
//   - options - ExpressRouteCrossConnectionsClientBeginListArpTableOptions contains the optional parameters for the ExpressRouteCrossConnectionsClient.BeginListArpTable
//     method.
func (client *ExpressRouteCrossConnectionsClient) BeginListArpTable(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsClientBeginListArpTableOptions) (*runtime.Poller[ExpressRouteCrossConnectionsClientListArpTableResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.listArpTable(ctx, resourceGroupName, crossConnectionName, peeringName, devicePath, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ExpressRouteCrossConnectionsClientListArpTableResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ExpressRouteCrossConnectionsClientListArpTableResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// ListArpTable - Gets the currently advertised ARP table associated with the express route cross connection in a resource
// group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *ExpressRouteCrossConnectionsClient) listArpTable(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsClientBeginListArpTableOptions) (*http.Response, error) {
	var err error
	const operationName = "ExpressRouteCrossConnectionsClient.BeginListArpTable"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listArpTableCreateRequest(ctx, resourceGroupName, crossConnectionName, peeringName, devicePath, options)
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

// listArpTableCreateRequest creates the ListArpTable request.
func (client *ExpressRouteCrossConnectionsClient) listArpTableCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsClientBeginListArpTableOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}/arpTables/{devicePath}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if crossConnectionName == "" {
		return nil, errors.New("parameter crossConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if devicePath == "" {
		return nil, errors.New("parameter devicePath cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devicePath}", url.PathEscape(devicePath))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// NewListByResourceGroupPager - Retrieves all the ExpressRouteCrossConnections in a resource group.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - options - ExpressRouteCrossConnectionsClientListByResourceGroupOptions contains the optional parameters for the ExpressRouteCrossConnectionsClient.NewListByResourceGroupPager
//     method.
func (client *ExpressRouteCrossConnectionsClient) NewListByResourceGroupPager(resourceGroupName string, options *ExpressRouteCrossConnectionsClientListByResourceGroupOptions) *runtime.Pager[ExpressRouteCrossConnectionsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ExpressRouteCrossConnectionsClientListByResourceGroupResponse]{
		More: func(page ExpressRouteCrossConnectionsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ExpressRouteCrossConnectionsClientListByResourceGroupResponse) (ExpressRouteCrossConnectionsClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ExpressRouteCrossConnectionsClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return ExpressRouteCrossConnectionsClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ExpressRouteCrossConnectionsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ExpressRouteCrossConnectionsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ExpressRouteCrossConnectionsClient) listByResourceGroupHandleResponse(resp *http.Response) (ExpressRouteCrossConnectionsClientListByResourceGroupResponse, error) {
	result := ExpressRouteCrossConnectionsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteCrossConnectionListResult); err != nil {
		return ExpressRouteCrossConnectionsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// BeginListRoutesTable - Gets the currently advertised routes table associated with the express route cross connection in
// a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - crossConnectionName - The name of the ExpressRouteCrossConnection.
//   - peeringName - The name of the peering.
//   - devicePath - The path of the device.
//   - options - ExpressRouteCrossConnectionsClientBeginListRoutesTableOptions contains the optional parameters for the ExpressRouteCrossConnectionsClient.BeginListRoutesTable
//     method.
func (client *ExpressRouteCrossConnectionsClient) BeginListRoutesTable(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsClientBeginListRoutesTableOptions) (*runtime.Poller[ExpressRouteCrossConnectionsClientListRoutesTableResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.listRoutesTable(ctx, resourceGroupName, crossConnectionName, peeringName, devicePath, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ExpressRouteCrossConnectionsClientListRoutesTableResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ExpressRouteCrossConnectionsClientListRoutesTableResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// ListRoutesTable - Gets the currently advertised routes table associated with the express route cross connection in a resource
// group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *ExpressRouteCrossConnectionsClient) listRoutesTable(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsClientBeginListRoutesTableOptions) (*http.Response, error) {
	var err error
	const operationName = "ExpressRouteCrossConnectionsClient.BeginListRoutesTable"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listRoutesTableCreateRequest(ctx, resourceGroupName, crossConnectionName, peeringName, devicePath, options)
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

// listRoutesTableCreateRequest creates the ListRoutesTable request.
func (client *ExpressRouteCrossConnectionsClient) listRoutesTableCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsClientBeginListRoutesTableOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}/routeTables/{devicePath}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if crossConnectionName == "" {
		return nil, errors.New("parameter crossConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if devicePath == "" {
		return nil, errors.New("parameter devicePath cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devicePath}", url.PathEscape(devicePath))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginListRoutesTableSummary - Gets the route table summary associated with the express route cross connection in a resource
// group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - crossConnectionName - The name of the ExpressRouteCrossConnection.
//   - peeringName - The name of the peering.
//   - devicePath - The path of the device.
//   - options - ExpressRouteCrossConnectionsClientBeginListRoutesTableSummaryOptions contains the optional parameters for the
//     ExpressRouteCrossConnectionsClient.BeginListRoutesTableSummary method.
func (client *ExpressRouteCrossConnectionsClient) BeginListRoutesTableSummary(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsClientBeginListRoutesTableSummaryOptions) (*runtime.Poller[ExpressRouteCrossConnectionsClientListRoutesTableSummaryResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.listRoutesTableSummary(ctx, resourceGroupName, crossConnectionName, peeringName, devicePath, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ExpressRouteCrossConnectionsClientListRoutesTableSummaryResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ExpressRouteCrossConnectionsClientListRoutesTableSummaryResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// ListRoutesTableSummary - Gets the route table summary associated with the express route cross connection in a resource
// group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *ExpressRouteCrossConnectionsClient) listRoutesTableSummary(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsClientBeginListRoutesTableSummaryOptions) (*http.Response, error) {
	var err error
	const operationName = "ExpressRouteCrossConnectionsClient.BeginListRoutesTableSummary"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listRoutesTableSummaryCreateRequest(ctx, resourceGroupName, crossConnectionName, peeringName, devicePath, options)
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

// listRoutesTableSummaryCreateRequest creates the ListRoutesTableSummary request.
func (client *ExpressRouteCrossConnectionsClient) listRoutesTableSummaryCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsClientBeginListRoutesTableSummaryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}/routeTablesSummary/{devicePath}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if crossConnectionName == "" {
		return nil, errors.New("parameter crossConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if devicePath == "" {
		return nil, errors.New("parameter devicePath cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devicePath}", url.PathEscape(devicePath))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// UpdateTags - Updates an express route cross connection tags.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - crossConnectionName - The name of the cross connection.
//   - crossConnectionParameters - Parameters supplied to update express route cross connection tags.
//   - options - ExpressRouteCrossConnectionsClientUpdateTagsOptions contains the optional parameters for the ExpressRouteCrossConnectionsClient.UpdateTags
//     method.
func (client *ExpressRouteCrossConnectionsClient) UpdateTags(ctx context.Context, resourceGroupName string, crossConnectionName string, crossConnectionParameters TagsObject, options *ExpressRouteCrossConnectionsClientUpdateTagsOptions) (ExpressRouteCrossConnectionsClientUpdateTagsResponse, error) {
	var err error
	const operationName = "ExpressRouteCrossConnectionsClient.UpdateTags"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, crossConnectionName, crossConnectionParameters, options)
	if err != nil {
		return ExpressRouteCrossConnectionsClientUpdateTagsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExpressRouteCrossConnectionsClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ExpressRouteCrossConnectionsClientUpdateTagsResponse{}, err
	}
	resp, err := client.updateTagsHandleResponse(httpResp)
	return resp, err
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *ExpressRouteCrossConnectionsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, crossConnectionParameters TagsObject, options *ExpressRouteCrossConnectionsClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if crossConnectionName == "" {
		return nil, errors.New("parameter crossConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, crossConnectionParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *ExpressRouteCrossConnectionsClient) updateTagsHandleResponse(resp *http.Response) (ExpressRouteCrossConnectionsClientUpdateTagsResponse, error) {
	result := ExpressRouteCrossConnectionsClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteCrossConnection); err != nil {
		return ExpressRouteCrossConnectionsClientUpdateTagsResponse{}, err
	}
	return result, nil
}
