// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// LocalNetworkGatewaysOperations contains the methods for the LocalNetworkGateways group.
type LocalNetworkGatewaysOperations interface {
	// BeginCreateOrUpdate - Creates or updates a local network gateway in the specified resource group.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, parameters LocalNetworkGateway, options *LocalNetworkGatewaysCreateOrUpdateOptions) (*LocalNetworkGatewayPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (LocalNetworkGatewayPoller, error)
	// BeginDelete - Deletes the specified local network gateway.
	BeginDelete(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, options *LocalNetworkGatewaysDeleteOptions) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets the specified local network gateway in a resource group.
	Get(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, options *LocalNetworkGatewaysGetOptions) (*LocalNetworkGatewayResponse, error)
	// List - Gets all the local network gateways in a resource group.
	List(resourceGroupName string, options *LocalNetworkGatewaysListOptions) LocalNetworkGatewayListResultPager
	// UpdateTags - Updates a local network gateway tags.
	UpdateTags(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, parameters TagsObject, options *LocalNetworkGatewaysUpdateTagsOptions) (*LocalNetworkGatewayResponse, error)
}

// LocalNetworkGatewaysClient implements the LocalNetworkGatewaysOperations interface.
// Don't use this type directly, use NewLocalNetworkGatewaysClient() instead.
type LocalNetworkGatewaysClient struct {
	con            *Connection
	subscriptionID string
}

// NewLocalNetworkGatewaysClient creates a new instance of LocalNetworkGatewaysClient with the specified values.
func NewLocalNetworkGatewaysClient(con *Connection, subscriptionID string) LocalNetworkGatewaysOperations {
	return &LocalNetworkGatewaysClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *LocalNetworkGatewaysClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

func (client *LocalNetworkGatewaysClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, parameters LocalNetworkGateway, options *LocalNetworkGatewaysCreateOrUpdateOptions) (*LocalNetworkGatewayPollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, localNetworkGatewayName, parameters, options)
	if err != nil {
		return nil, err
	}
	result := &LocalNetworkGatewayPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("LocalNetworkGatewaysClient.CreateOrUpdate", "azure-async-operation", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &localNetworkGatewayPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*LocalNetworkGatewayResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *LocalNetworkGatewaysClient) ResumeCreateOrUpdate(token string) (LocalNetworkGatewayPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("LocalNetworkGatewaysClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &localNetworkGatewayPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates a local network gateway in the specified resource group.
func (client *LocalNetworkGatewaysClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, parameters LocalNetworkGateway, options *LocalNetworkGatewaysCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, localNetworkGatewayName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.CreateOrUpdateHandleError(resp)
	}
	return resp, nil
}

// CreateOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *LocalNetworkGatewaysClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, parameters LocalNetworkGateway, options *LocalNetworkGatewaysCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways/{localNetworkGatewayName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{localNetworkGatewayName}", url.PathEscape(localNetworkGatewayName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// CreateOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *LocalNetworkGatewaysClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*LocalNetworkGatewayResponse, error) {
	result := LocalNetworkGatewayResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.LocalNetworkGateway)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *LocalNetworkGatewaysClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *LocalNetworkGatewaysClient) BeginDelete(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, options *LocalNetworkGatewaysDeleteOptions) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, localNetworkGatewayName, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("LocalNetworkGatewaysClient.Delete", "location", resp, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *LocalNetworkGatewaysClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("LocalNetworkGatewaysClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified local network gateway.
func (client *LocalNetworkGatewaysClient) Delete(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, options *LocalNetworkGatewaysDeleteOptions) (*azcore.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, localNetworkGatewayName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.DeleteHandleError(resp)
	}
	return resp, nil
}

// DeleteCreateRequest creates the Delete request.
func (client *LocalNetworkGatewaysClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, options *LocalNetworkGatewaysDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways/{localNetworkGatewayName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{localNetworkGatewayName}", url.PathEscape(localNetworkGatewayName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// DeleteHandleError handles the Delete error response.
func (client *LocalNetworkGatewaysClient) DeleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the specified local network gateway in a resource group.
func (client *LocalNetworkGatewaysClient) Get(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, options *LocalNetworkGatewaysGetOptions) (*LocalNetworkGatewayResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, localNetworkGatewayName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetHandleError(resp)
	}
	result, err := client.GetHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCreateRequest creates the Get request.
func (client *LocalNetworkGatewaysClient) GetCreateRequest(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, options *LocalNetworkGatewaysGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways/{localNetworkGatewayName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{localNetworkGatewayName}", url.PathEscape(localNetworkGatewayName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *LocalNetworkGatewaysClient) GetHandleResponse(resp *azcore.Response) (*LocalNetworkGatewayResponse, error) {
	result := LocalNetworkGatewayResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.LocalNetworkGateway)
}

// GetHandleError handles the Get error response.
func (client *LocalNetworkGatewaysClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Gets all the local network gateways in a resource group.
func (client *LocalNetworkGatewaysClient) List(resourceGroupName string, options *LocalNetworkGatewaysListOptions) LocalNetworkGatewayListResultPager {
	return &localNetworkGatewayListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *LocalNetworkGatewayListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.LocalNetworkGatewayListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListCreateRequest creates the List request.
func (client *LocalNetworkGatewaysClient) ListCreateRequest(ctx context.Context, resourceGroupName string, options *LocalNetworkGatewaysListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListHandleResponse handles the List response.
func (client *LocalNetworkGatewaysClient) ListHandleResponse(resp *azcore.Response) (*LocalNetworkGatewayListResultResponse, error) {
	result := LocalNetworkGatewayListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.LocalNetworkGatewayListResult)
}

// ListHandleError handles the List error response.
func (client *LocalNetworkGatewaysClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// UpdateTags - Updates a local network gateway tags.
func (client *LocalNetworkGatewaysClient) UpdateTags(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, parameters TagsObject, options *LocalNetworkGatewaysUpdateTagsOptions) (*LocalNetworkGatewayResponse, error) {
	req, err := client.UpdateTagsCreateRequest(ctx, resourceGroupName, localNetworkGatewayName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.UpdateTagsHandleError(resp)
	}
	result, err := client.UpdateTagsHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateTagsCreateRequest creates the UpdateTags request.
func (client *LocalNetworkGatewaysClient) UpdateTagsCreateRequest(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, parameters TagsObject, options *LocalNetworkGatewaysUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways/{localNetworkGatewayName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{localNetworkGatewayName}", url.PathEscape(localNetworkGatewayName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// UpdateTagsHandleResponse handles the UpdateTags response.
func (client *LocalNetworkGatewaysClient) UpdateTagsHandleResponse(resp *azcore.Response) (*LocalNetworkGatewayResponse, error) {
	result := LocalNetworkGatewayResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.LocalNetworkGateway)
}

// UpdateTagsHandleError handles the UpdateTags error response.
func (client *LocalNetworkGatewaysClient) UpdateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
