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

// VirtualHubRouteTableV2SClient contains the methods for the VirtualHubRouteTableV2S group.
// Don't use this type directly, use NewVirtualHubRouteTableV2SClient() instead.
type VirtualHubRouteTableV2SClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVirtualHubRouteTableV2SClient creates a new instance of VirtualHubRouteTableV2SClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVirtualHubRouteTableV2SClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *VirtualHubRouteTableV2SClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	client := &VirtualHubRouteTableV2SClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Host),
		pl:             armruntime.NewPipeline(module, version, credential, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Creates a VirtualHubRouteTableV2 resource if it doesn't exist else updates the existing VirtualHubRouteTableV2.
// If the operation fails it returns the *Error error type.
// resourceGroupName - The resource group name of the VirtualHub.
// virtualHubName - The name of the VirtualHub.
// routeTableName - The name of the VirtualHubRouteTableV2.
// virtualHubRouteTableV2Parameters - Parameters supplied to create or update VirtualHubRouteTableV2.
// options - VirtualHubRouteTableV2SClientBeginCreateOrUpdateOptions contains the optional parameters for the VirtualHubRouteTableV2SClient.BeginCreateOrUpdate
// method.
func (client *VirtualHubRouteTableV2SClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, virtualHubRouteTableV2Parameters VirtualHubRouteTableV2, options *VirtualHubRouteTableV2SClientBeginCreateOrUpdateOptions) (VirtualHubRouteTableV2SClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, virtualHubName, routeTableName, virtualHubRouteTableV2Parameters, options)
	if err != nil {
		return VirtualHubRouteTableV2SClientCreateOrUpdatePollerResponse{}, err
	}
	result := VirtualHubRouteTableV2SClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualHubRouteTableV2SClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return VirtualHubRouteTableV2SClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &VirtualHubRouteTableV2SClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates a VirtualHubRouteTableV2 resource if it doesn't exist else updates the existing VirtualHubRouteTableV2.
// If the operation fails it returns the *Error error type.
func (client *VirtualHubRouteTableV2SClient) createOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, virtualHubRouteTableV2Parameters VirtualHubRouteTableV2, options *VirtualHubRouteTableV2SClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, virtualHubName, routeTableName, virtualHubRouteTableV2Parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VirtualHubRouteTableV2SClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, virtualHubRouteTableV2Parameters VirtualHubRouteTableV2, options *VirtualHubRouteTableV2SClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/routeTables/{routeTableName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	if routeTableName == "" {
		return nil, errors.New("parameter routeTableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeTableName}", url.PathEscape(routeTableName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, virtualHubRouteTableV2Parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *VirtualHubRouteTableV2SClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes a VirtualHubRouteTableV2.
// If the operation fails it returns the *Error error type.
// resourceGroupName - The resource group name of the VirtualHubRouteTableV2.
// virtualHubName - The name of the VirtualHub.
// routeTableName - The name of the VirtualHubRouteTableV2.
// options - VirtualHubRouteTableV2SClientBeginDeleteOptions contains the optional parameters for the VirtualHubRouteTableV2SClient.BeginDelete
// method.
func (client *VirtualHubRouteTableV2SClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, options *VirtualHubRouteTableV2SClientBeginDeleteOptions) (VirtualHubRouteTableV2SClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, virtualHubName, routeTableName, options)
	if err != nil {
		return VirtualHubRouteTableV2SClientDeletePollerResponse{}, err
	}
	result := VirtualHubRouteTableV2SClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualHubRouteTableV2SClient.Delete", "location", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return VirtualHubRouteTableV2SClientDeletePollerResponse{}, err
	}
	result.Poller = &VirtualHubRouteTableV2SClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a VirtualHubRouteTableV2.
// If the operation fails it returns the *Error error type.
func (client *VirtualHubRouteTableV2SClient) deleteOperation(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, options *VirtualHubRouteTableV2SClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualHubName, routeTableName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VirtualHubRouteTableV2SClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, options *VirtualHubRouteTableV2SClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/routeTables/{routeTableName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	if routeTableName == "" {
		return nil, errors.New("parameter routeTableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeTableName}", url.PathEscape(routeTableName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *VirtualHubRouteTableV2SClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Retrieves the details of a VirtualHubRouteTableV2.
// If the operation fails it returns the *Error error type.
// resourceGroupName - The resource group name of the VirtualHubRouteTableV2.
// virtualHubName - The name of the VirtualHub.
// routeTableName - The name of the VirtualHubRouteTableV2.
// options - VirtualHubRouteTableV2SClientGetOptions contains the optional parameters for the VirtualHubRouteTableV2SClient.Get
// method.
func (client *VirtualHubRouteTableV2SClient) Get(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, options *VirtualHubRouteTableV2SClientGetOptions) (VirtualHubRouteTableV2SClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualHubName, routeTableName, options)
	if err != nil {
		return VirtualHubRouteTableV2SClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualHubRouteTableV2SClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualHubRouteTableV2SClientGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualHubRouteTableV2SClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, options *VirtualHubRouteTableV2SClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/routeTables/{routeTableName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	if routeTableName == "" {
		return nil, errors.New("parameter routeTableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeTableName}", url.PathEscape(routeTableName))
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
func (client *VirtualHubRouteTableV2SClient) getHandleResponse(resp *http.Response) (VirtualHubRouteTableV2SClientGetResponse, error) {
	result := VirtualHubRouteTableV2SClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualHubRouteTableV2); err != nil {
		return VirtualHubRouteTableV2SClientGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *VirtualHubRouteTableV2SClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Retrieves the details of all VirtualHubRouteTableV2s.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The resource group name of the VirtualHub.
// virtualHubName - The name of the VirtualHub.
// options - VirtualHubRouteTableV2SClientListOptions contains the optional parameters for the VirtualHubRouteTableV2SClient.List
// method.
func (client *VirtualHubRouteTableV2SClient) List(resourceGroupName string, virtualHubName string, options *VirtualHubRouteTableV2SClientListOptions) *VirtualHubRouteTableV2SClientListPager {
	return &VirtualHubRouteTableV2SClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, virtualHubName, options)
		},
		advancer: func(ctx context.Context, resp VirtualHubRouteTableV2SClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListVirtualHubRouteTableV2SResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *VirtualHubRouteTableV2SClient) listCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, options *VirtualHubRouteTableV2SClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/routeTables"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
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
func (client *VirtualHubRouteTableV2SClient) listHandleResponse(resp *http.Response) (VirtualHubRouteTableV2SClientListResponse, error) {
	result := VirtualHubRouteTableV2SClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListVirtualHubRouteTableV2SResult); err != nil {
		return VirtualHubRouteTableV2SClientListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *VirtualHubRouteTableV2SClient) listHandleError(resp *http.Response) error {
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
