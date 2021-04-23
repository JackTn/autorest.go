// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// VPNGatewaysClient contains the methods for the VPNGateways group.
// Don't use this type directly, use NewVPNGatewaysClient() instead.
type VPNGatewaysClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVPNGatewaysClient creates a new instance of VPNGatewaysClient with the specified values.
func NewVPNGatewaysClient(con *armcore.Connection, subscriptionID string) *VPNGatewaysClient {
	return &VPNGatewaysClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates a virtual wan vpn gateway if it doesn't exist else updates the existing gateway.
func (client *VPNGatewaysClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, gatewayName string, vpnGatewayParameters VPNGateway, options *VPNGatewaysBeginCreateOrUpdateOptions) (VPNGatewayPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, gatewayName, vpnGatewayParameters, options)
	if err != nil {
		return VPNGatewayPollerResponse{}, err
	}
	result := VPNGatewayPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VPNGatewaysClient.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return VPNGatewayPollerResponse{}, err
	}
	poller := &vpnGatewayPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (VPNGatewayResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new VPNGatewayPoller from the specified resume token.
// token - The value must come from a previous call to VPNGatewayPoller.ResumeToken().
func (client *VPNGatewaysClient) ResumeCreateOrUpdate(ctx context.Context, token string) (VPNGatewayPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("VPNGatewaysClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return VPNGatewayPollerResponse{}, err
	}
	poller := &vpnGatewayPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return VPNGatewayPollerResponse{}, err
	}
	result := VPNGatewayPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (VPNGatewayResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdate - Creates a virtual wan vpn gateway if it doesn't exist else updates the existing gateway.
func (client *VPNGatewaysClient) createOrUpdate(ctx context.Context, resourceGroupName string, gatewayName string, vpnGatewayParameters VPNGateway, options *VPNGatewaysBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, gatewayName, vpnGatewayParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VPNGatewaysClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, vpnGatewayParameters VPNGateway, options *VPNGatewaysBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(vpnGatewayParameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *VPNGatewaysClient) createOrUpdateHandleResponse(resp *azcore.Response) (VPNGatewayResponse, error) {
	var val *VPNGateway
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return VPNGatewayResponse{}, err
	}
	return VPNGatewayResponse{RawResponse: resp.Response, VPNGateway: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *VPNGatewaysClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Deletes a virtual wan vpn gateway.
func (client *VPNGatewaysClient) BeginDelete(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, gatewayName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VPNGatewaysClient.Delete", "location", resp, client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *VPNGatewaysClient) ResumeDelete(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("VPNGatewaysClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Delete - Deletes a virtual wan vpn gateway.
func (client *VPNGatewaysClient) deleteOperation(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, gatewayName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VPNGatewaysClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *VPNGatewaysClient) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Retrieves the details of a virtual wan vpn gateway.
func (client *VPNGatewaysClient) Get(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysGetOptions) (VPNGatewayResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, gatewayName, options)
	if err != nil {
		return VPNGatewayResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return VPNGatewayResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VPNGatewayResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VPNGatewaysClient) getCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VPNGatewaysClient) getHandleResponse(resp *azcore.Response) (VPNGatewayResponse, error) {
	var val *VPNGateway
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return VPNGatewayResponse{}, err
	}
	return VPNGatewayResponse{RawResponse: resp.Response, VPNGateway: val}, nil
}

// getHandleError handles the Get error response.
func (client *VPNGatewaysClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Lists all the VpnGateways in a subscription.
func (client *VPNGatewaysClient) List(options *VPNGatewaysListOptions) ListVPNGatewaysResultPager {
	return &listVPNGatewaysResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp ListVPNGatewaysResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ListVPNGatewaysResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *VPNGatewaysClient) listCreateRequest(ctx context.Context, options *VPNGatewaysListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/vpnGateways"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VPNGatewaysClient) listHandleResponse(resp *azcore.Response) (ListVPNGatewaysResultResponse, error) {
	var val *ListVPNGatewaysResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ListVPNGatewaysResultResponse{}, err
	}
	return ListVPNGatewaysResultResponse{RawResponse: resp.Response, ListVPNGatewaysResult: val}, nil
}

// listHandleError handles the List error response.
func (client *VPNGatewaysClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListByResourceGroup - Lists all the VpnGateways in a resource group.
func (client *VPNGatewaysClient) ListByResourceGroup(resourceGroupName string, options *VPNGatewaysListByResourceGroupOptions) ListVPNGatewaysResultPager {
	return &listVPNGatewaysResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listByResourceGroupHandleResponse,
		errorer:   client.listByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp ListVPNGatewaysResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ListVPNGatewaysResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *VPNGatewaysClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *VPNGatewaysListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *VPNGatewaysClient) listByResourceGroupHandleResponse(resp *azcore.Response) (ListVPNGatewaysResultResponse, error) {
	var val *ListVPNGatewaysResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ListVPNGatewaysResultResponse{}, err
	}
	return ListVPNGatewaysResultResponse{RawResponse: resp.Response, ListVPNGatewaysResult: val}, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *VPNGatewaysClient) listByResourceGroupHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginReset - Resets the primary of the vpn gateway in the specified resource group.
func (client *VPNGatewaysClient) BeginReset(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysBeginResetOptions) (VPNGatewayPollerResponse, error) {
	resp, err := client.reset(ctx, resourceGroupName, gatewayName, options)
	if err != nil {
		return VPNGatewayPollerResponse{}, err
	}
	result := VPNGatewayPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VPNGatewaysClient.Reset", "location", resp, client.resetHandleError)
	if err != nil {
		return VPNGatewayPollerResponse{}, err
	}
	poller := &vpnGatewayPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (VPNGatewayResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeReset creates a new VPNGatewayPoller from the specified resume token.
// token - The value must come from a previous call to VPNGatewayPoller.ResumeToken().
func (client *VPNGatewaysClient) ResumeReset(ctx context.Context, token string) (VPNGatewayPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("VPNGatewaysClient.Reset", token, client.resetHandleError)
	if err != nil {
		return VPNGatewayPollerResponse{}, err
	}
	poller := &vpnGatewayPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return VPNGatewayPollerResponse{}, err
	}
	result := VPNGatewayPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (VPNGatewayResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Reset - Resets the primary of the vpn gateway in the specified resource group.
func (client *VPNGatewaysClient) reset(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysBeginResetOptions) (*azcore.Response, error) {
	req, err := client.resetCreateRequest(ctx, resourceGroupName, gatewayName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.resetHandleError(resp)
	}
	return resp, nil
}

// resetCreateRequest creates the Reset request.
func (client *VPNGatewaysClient) resetCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysBeginResetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/reset"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// resetHandleResponse handles the Reset response.
func (client *VPNGatewaysClient) resetHandleResponse(resp *azcore.Response) (VPNGatewayResponse, error) {
	var val *VPNGateway
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return VPNGatewayResponse{}, err
	}
	return VPNGatewayResponse{RawResponse: resp.Response, VPNGateway: val}, nil
}

// resetHandleError handles the Reset error response.
func (client *VPNGatewaysClient) resetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// UpdateTags - Updates virtual wan vpn gateway tags.
func (client *VPNGatewaysClient) UpdateTags(ctx context.Context, resourceGroupName string, gatewayName string, vpnGatewayParameters TagsObject, options *VPNGatewaysUpdateTagsOptions) (VPNGatewayResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, gatewayName, vpnGatewayParameters, options)
	if err != nil {
		return VPNGatewayResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return VPNGatewayResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VPNGatewayResponse{}, client.updateTagsHandleError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *VPNGatewaysClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, vpnGatewayParameters TagsObject, options *VPNGatewaysUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(vpnGatewayParameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *VPNGatewaysClient) updateTagsHandleResponse(resp *azcore.Response) (VPNGatewayResponse, error) {
	var val *VPNGateway
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return VPNGatewayResponse{}, err
	}
	return VPNGatewayResponse{RawResponse: resp.Response, VPNGateway: val}, nil
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *VPNGatewaysClient) updateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
