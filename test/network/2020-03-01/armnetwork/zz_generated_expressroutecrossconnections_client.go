//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ExpressRouteCrossConnectionsClient contains the methods for the ExpressRouteCrossConnections group.
// Don't use this type directly, use NewExpressRouteCrossConnectionsClient() instead.
type ExpressRouteCrossConnectionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewExpressRouteCrossConnectionsClient creates a new instance of ExpressRouteCrossConnectionsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewExpressRouteCrossConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ExpressRouteCrossConnectionsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	client := &ExpressRouteCrossConnectionsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Host),
		pl:             armruntime.NewPipeline(module, version, credential, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Update the specified ExpressRouteCrossConnection.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// crossConnectionName - The name of the ExpressRouteCrossConnection.
// parameters - Parameters supplied to the update express route crossConnection operation.
// options - ExpressRouteCrossConnectionsClientBeginCreateOrUpdateOptions contains the optional parameters for the ExpressRouteCrossConnectionsClient.BeginCreateOrUpdate
// method.
func (client *ExpressRouteCrossConnectionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, crossConnectionName string, parameters ExpressRouteCrossConnection, options *ExpressRouteCrossConnectionsClientBeginCreateOrUpdateOptions) (ExpressRouteCrossConnectionsClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, crossConnectionName, parameters, options)
	if err != nil {
		return ExpressRouteCrossConnectionsClientCreateOrUpdatePollerResponse{}, err
	}
	result := ExpressRouteCrossConnectionsClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ExpressRouteCrossConnectionsClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return ExpressRouteCrossConnectionsClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &ExpressRouteCrossConnectionsClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Update the specified ExpressRouteCrossConnection.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteCrossConnectionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, crossConnectionName string, parameters ExpressRouteCrossConnection, options *ExpressRouteCrossConnectionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, crossConnectionName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ExpressRouteCrossConnectionsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets details about the specified ExpressRouteCrossConnection.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group (peering location of the circuit).
// crossConnectionName - The name of the ExpressRouteCrossConnection (service key of the circuit).
// options - ExpressRouteCrossConnectionsClientGetOptions contains the optional parameters for the ExpressRouteCrossConnectionsClient.Get
// method.
func (client *ExpressRouteCrossConnectionsClient) Get(ctx context.Context, resourceGroupName string, crossConnectionName string, options *ExpressRouteCrossConnectionsClientGetOptions) (ExpressRouteCrossConnectionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, crossConnectionName, options)
	if err != nil {
		return ExpressRouteCrossConnectionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExpressRouteCrossConnectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExpressRouteCrossConnectionsClientGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExpressRouteCrossConnectionsClient) getHandleResponse(resp *http.Response) (ExpressRouteCrossConnectionsClientGetResponse, error) {
	result := ExpressRouteCrossConnectionsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteCrossConnection); err != nil {
		return ExpressRouteCrossConnectionsClientGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ExpressRouteCrossConnectionsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Retrieves all the ExpressRouteCrossConnections in a subscription.
// If the operation fails it returns the *CloudError error type.
// options - ExpressRouteCrossConnectionsClientListOptions contains the optional parameters for the ExpressRouteCrossConnectionsClient.List
// method.
func (client *ExpressRouteCrossConnectionsClient) List(options *ExpressRouteCrossConnectionsClientListOptions) *ExpressRouteCrossConnectionsClientListPager {
	return &ExpressRouteCrossConnectionsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ExpressRouteCrossConnectionsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ExpressRouteCrossConnectionListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ExpressRouteCrossConnectionsClient) listCreateRequest(ctx context.Context, options *ExpressRouteCrossConnectionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteCrossConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ExpressRouteCrossConnectionsClient) listHandleResponse(resp *http.Response) (ExpressRouteCrossConnectionsClientListResponse, error) {
	result := ExpressRouteCrossConnectionsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteCrossConnectionListResult); err != nil {
		return ExpressRouteCrossConnectionsClientListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ExpressRouteCrossConnectionsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginListArpTable - Gets the currently advertised ARP table associated with the express route cross connection in a resource
// group.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// crossConnectionName - The name of the ExpressRouteCrossConnection.
// peeringName - The name of the peering.
// devicePath - The path of the device.
// options - ExpressRouteCrossConnectionsClientBeginListArpTableOptions contains the optional parameters for the ExpressRouteCrossConnectionsClient.BeginListArpTable
// method.
func (client *ExpressRouteCrossConnectionsClient) BeginListArpTable(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsClientBeginListArpTableOptions) (ExpressRouteCrossConnectionsClientListArpTablePollerResponse, error) {
	resp, err := client.listArpTable(ctx, resourceGroupName, crossConnectionName, peeringName, devicePath, options)
	if err != nil {
		return ExpressRouteCrossConnectionsClientListArpTablePollerResponse{}, err
	}
	result := ExpressRouteCrossConnectionsClientListArpTablePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ExpressRouteCrossConnectionsClient.ListArpTable", "location", resp, client.pl, client.listArpTableHandleError)
	if err != nil {
		return ExpressRouteCrossConnectionsClientListArpTablePollerResponse{}, err
	}
	result.Poller = &ExpressRouteCrossConnectionsClientListArpTablePoller{
		pt: pt,
	}
	return result, nil
}

// ListArpTable - Gets the currently advertised ARP table associated with the express route cross connection in a resource
// group.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteCrossConnectionsClient) listArpTable(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsClientBeginListArpTableOptions) (*http.Response, error) {
	req, err := client.listArpTableCreateRequest(ctx, resourceGroupName, crossConnectionName, peeringName, devicePath, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.listArpTableHandleError(resp)
	}
	return resp, nil
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listArpTableHandleError handles the ListArpTable error response.
func (client *ExpressRouteCrossConnectionsClient) listArpTableHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByResourceGroup - Retrieves all the ExpressRouteCrossConnections in a resource group.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// options - ExpressRouteCrossConnectionsClientListByResourceGroupOptions contains the optional parameters for the ExpressRouteCrossConnectionsClient.ListByResourceGroup
// method.
func (client *ExpressRouteCrossConnectionsClient) ListByResourceGroup(resourceGroupName string, options *ExpressRouteCrossConnectionsClientListByResourceGroupOptions) *ExpressRouteCrossConnectionsClientListByResourceGroupPager {
	return &ExpressRouteCrossConnectionsClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp ExpressRouteCrossConnectionsClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ExpressRouteCrossConnectionListResult.NextLink)
		},
	}
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ExpressRouteCrossConnectionsClient) listByResourceGroupHandleResponse(resp *http.Response) (ExpressRouteCrossConnectionsClientListByResourceGroupResponse, error) {
	result := ExpressRouteCrossConnectionsClientListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteCrossConnectionListResult); err != nil {
		return ExpressRouteCrossConnectionsClientListByResourceGroupResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *ExpressRouteCrossConnectionsClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginListRoutesTable - Gets the currently advertised routes table associated with the express route cross connection in
// a resource group.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// crossConnectionName - The name of the ExpressRouteCrossConnection.
// peeringName - The name of the peering.
// devicePath - The path of the device.
// options - ExpressRouteCrossConnectionsClientBeginListRoutesTableOptions contains the optional parameters for the ExpressRouteCrossConnectionsClient.BeginListRoutesTable
// method.
func (client *ExpressRouteCrossConnectionsClient) BeginListRoutesTable(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsClientBeginListRoutesTableOptions) (ExpressRouteCrossConnectionsClientListRoutesTablePollerResponse, error) {
	resp, err := client.listRoutesTable(ctx, resourceGroupName, crossConnectionName, peeringName, devicePath, options)
	if err != nil {
		return ExpressRouteCrossConnectionsClientListRoutesTablePollerResponse{}, err
	}
	result := ExpressRouteCrossConnectionsClientListRoutesTablePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ExpressRouteCrossConnectionsClient.ListRoutesTable", "location", resp, client.pl, client.listRoutesTableHandleError)
	if err != nil {
		return ExpressRouteCrossConnectionsClientListRoutesTablePollerResponse{}, err
	}
	result.Poller = &ExpressRouteCrossConnectionsClientListRoutesTablePoller{
		pt: pt,
	}
	return result, nil
}

// ListRoutesTable - Gets the currently advertised routes table associated with the express route cross connection in a resource
// group.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteCrossConnectionsClient) listRoutesTable(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsClientBeginListRoutesTableOptions) (*http.Response, error) {
	req, err := client.listRoutesTableCreateRequest(ctx, resourceGroupName, crossConnectionName, peeringName, devicePath, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.listRoutesTableHandleError(resp)
	}
	return resp, nil
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listRoutesTableHandleError handles the ListRoutesTable error response.
func (client *ExpressRouteCrossConnectionsClient) listRoutesTableHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginListRoutesTableSummary - Gets the route table summary associated with the express route cross connection in a resource
// group.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// crossConnectionName - The name of the ExpressRouteCrossConnection.
// peeringName - The name of the peering.
// devicePath - The path of the device.
// options - ExpressRouteCrossConnectionsClientBeginListRoutesTableSummaryOptions contains the optional parameters for the
// ExpressRouteCrossConnectionsClient.BeginListRoutesTableSummary method.
func (client *ExpressRouteCrossConnectionsClient) BeginListRoutesTableSummary(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsClientBeginListRoutesTableSummaryOptions) (ExpressRouteCrossConnectionsClientListRoutesTableSummaryPollerResponse, error) {
	resp, err := client.listRoutesTableSummary(ctx, resourceGroupName, crossConnectionName, peeringName, devicePath, options)
	if err != nil {
		return ExpressRouteCrossConnectionsClientListRoutesTableSummaryPollerResponse{}, err
	}
	result := ExpressRouteCrossConnectionsClientListRoutesTableSummaryPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ExpressRouteCrossConnectionsClient.ListRoutesTableSummary", "location", resp, client.pl, client.listRoutesTableSummaryHandleError)
	if err != nil {
		return ExpressRouteCrossConnectionsClientListRoutesTableSummaryPollerResponse{}, err
	}
	result.Poller = &ExpressRouteCrossConnectionsClientListRoutesTableSummaryPoller{
		pt: pt,
	}
	return result, nil
}

// ListRoutesTableSummary - Gets the route table summary associated with the express route cross connection in a resource
// group.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteCrossConnectionsClient) listRoutesTableSummary(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsClientBeginListRoutesTableSummaryOptions) (*http.Response, error) {
	req, err := client.listRoutesTableSummaryCreateRequest(ctx, resourceGroupName, crossConnectionName, peeringName, devicePath, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.listRoutesTableSummaryHandleError(resp)
	}
	return resp, nil
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listRoutesTableSummaryHandleError handles the ListRoutesTableSummary error response.
func (client *ExpressRouteCrossConnectionsClient) listRoutesTableSummaryHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// UpdateTags - Updates an express route cross connection tags.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// crossConnectionName - The name of the cross connection.
// crossConnectionParameters - Parameters supplied to update express route cross connection tags.
// options - ExpressRouteCrossConnectionsClientUpdateTagsOptions contains the optional parameters for the ExpressRouteCrossConnectionsClient.UpdateTags
// method.
func (client *ExpressRouteCrossConnectionsClient) UpdateTags(ctx context.Context, resourceGroupName string, crossConnectionName string, crossConnectionParameters TagsObject, options *ExpressRouteCrossConnectionsClientUpdateTagsOptions) (ExpressRouteCrossConnectionsClientUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, crossConnectionName, crossConnectionParameters, options)
	if err != nil {
		return ExpressRouteCrossConnectionsClientUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExpressRouteCrossConnectionsClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExpressRouteCrossConnectionsClientUpdateTagsResponse{}, client.updateTagsHandleError(resp)
	}
	return client.updateTagsHandleResponse(resp)
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
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, crossConnectionParameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *ExpressRouteCrossConnectionsClient) updateTagsHandleResponse(resp *http.Response) (ExpressRouteCrossConnectionsClientUpdateTagsResponse, error) {
	result := ExpressRouteCrossConnectionsClientUpdateTagsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteCrossConnection); err != nil {
		return ExpressRouteCrossConnectionsClientUpdateTagsResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *ExpressRouteCrossConnectionsClient) updateTagsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
