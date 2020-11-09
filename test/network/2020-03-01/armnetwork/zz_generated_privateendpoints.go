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

// PrivateEndpointsOperations contains the methods for the PrivateEndpoints group.
type PrivateEndpointsOperations interface {
	// BeginCreateOrUpdate - Creates or updates an private endpoint in the specified resource group.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, privateEndpointName string, parameters PrivateEndpoint, options *PrivateEndpointsCreateOrUpdateOptions) (*PrivateEndpointPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (PrivateEndpointPoller, error)
	// BeginDelete - Deletes the specified private endpoint.
	BeginDelete(ctx context.Context, resourceGroupName string, privateEndpointName string, options *PrivateEndpointsDeleteOptions) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets the specified private endpoint by resource group.
	Get(ctx context.Context, resourceGroupName string, privateEndpointName string, options *PrivateEndpointsGetOptions) (*PrivateEndpointResponse, error)
	// List - Gets all private endpoints in a resource group.
	List(resourceGroupName string, options *PrivateEndpointsListOptions) PrivateEndpointListResultPager
	// ListBySubscription - Gets all private endpoints in a subscription.
	ListBySubscription(options *PrivateEndpointsListBySubscriptionOptions) PrivateEndpointListResultPager
}

// PrivateEndpointsClient implements the PrivateEndpointsOperations interface.
// Don't use this type directly, use NewPrivateEndpointsClient() instead.
type PrivateEndpointsClient struct {
	con            *Connection
	subscriptionID string
}

// NewPrivateEndpointsClient creates a new instance of PrivateEndpointsClient with the specified values.
func NewPrivateEndpointsClient(con *Connection, subscriptionID string) PrivateEndpointsOperations {
	return &PrivateEndpointsClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *PrivateEndpointsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

func (client *PrivateEndpointsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, privateEndpointName string, parameters PrivateEndpoint, options *PrivateEndpointsCreateOrUpdateOptions) (*PrivateEndpointPollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, privateEndpointName, parameters, options)
	if err != nil {
		return nil, err
	}
	result := &PrivateEndpointPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("PrivateEndpointsClient.CreateOrUpdate", "azure-async-operation", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &privateEndpointPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*PrivateEndpointResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *PrivateEndpointsClient) ResumeCreateOrUpdate(token string) (PrivateEndpointPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("PrivateEndpointsClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &privateEndpointPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates an private endpoint in the specified resource group.
func (client *PrivateEndpointsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, privateEndpointName string, parameters PrivateEndpoint, options *PrivateEndpointsCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, privateEndpointName, parameters, options)
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
func (client *PrivateEndpointsClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, privateEndpointName string, parameters PrivateEndpoint, options *PrivateEndpointsCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointName}", url.PathEscape(privateEndpointName))
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
func (client *PrivateEndpointsClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*PrivateEndpointResponse, error) {
	result := PrivateEndpointResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PrivateEndpoint)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *PrivateEndpointsClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *PrivateEndpointsClient) BeginDelete(ctx context.Context, resourceGroupName string, privateEndpointName string, options *PrivateEndpointsDeleteOptions) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, privateEndpointName, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("PrivateEndpointsClient.Delete", "location", resp, client.DeleteHandleError)
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

func (client *PrivateEndpointsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("PrivateEndpointsClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified private endpoint.
func (client *PrivateEndpointsClient) Delete(ctx context.Context, resourceGroupName string, privateEndpointName string, options *PrivateEndpointsDeleteOptions) (*azcore.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, privateEndpointName, options)
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
func (client *PrivateEndpointsClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, privateEndpointName string, options *PrivateEndpointsDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointName}", url.PathEscape(privateEndpointName))
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
func (client *PrivateEndpointsClient) DeleteHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the specified private endpoint by resource group.
func (client *PrivateEndpointsClient) Get(ctx context.Context, resourceGroupName string, privateEndpointName string, options *PrivateEndpointsGetOptions) (*PrivateEndpointResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, privateEndpointName, options)
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
func (client *PrivateEndpointsClient) GetCreateRequest(ctx context.Context, resourceGroupName string, privateEndpointName string, options *PrivateEndpointsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointName}", url.PathEscape(privateEndpointName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	if options != nil && options.Expand != nil {
		query.Set("$expand", *options.Expand)
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *PrivateEndpointsClient) GetHandleResponse(resp *azcore.Response) (*PrivateEndpointResponse, error) {
	result := PrivateEndpointResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PrivateEndpoint)
}

// GetHandleError handles the Get error response.
func (client *PrivateEndpointsClient) GetHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Gets all private endpoints in a resource group.
func (client *PrivateEndpointsClient) List(resourceGroupName string, options *PrivateEndpointsListOptions) PrivateEndpointListResultPager {
	return &privateEndpointListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *PrivateEndpointListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.PrivateEndpointListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListCreateRequest creates the List request.
func (client *PrivateEndpointsClient) ListCreateRequest(ctx context.Context, resourceGroupName string, options *PrivateEndpointsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints"
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
func (client *PrivateEndpointsClient) ListHandleResponse(resp *azcore.Response) (*PrivateEndpointListResultResponse, error) {
	result := PrivateEndpointListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PrivateEndpointListResult)
}

// ListHandleError handles the List error response.
func (client *PrivateEndpointsClient) ListHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListBySubscription - Gets all private endpoints in a subscription.
func (client *PrivateEndpointsClient) ListBySubscription(options *PrivateEndpointsListBySubscriptionOptions) PrivateEndpointListResultPager {
	return &privateEndpointListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListBySubscriptionCreateRequest(ctx, options)
		},
		responder: client.ListBySubscriptionHandleResponse,
		errorer:   client.ListBySubscriptionHandleError,
		advancer: func(ctx context.Context, resp *PrivateEndpointListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.PrivateEndpointListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *PrivateEndpointsClient) ListBySubscriptionCreateRequest(ctx context.Context, options *PrivateEndpointsListBySubscriptionOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/privateEndpoints"
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

// ListBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *PrivateEndpointsClient) ListBySubscriptionHandleResponse(resp *azcore.Response) (*PrivateEndpointListResultResponse, error) {
	result := PrivateEndpointListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PrivateEndpointListResult)
}

// ListBySubscriptionHandleError handles the ListBySubscription error response.
func (client *PrivateEndpointsClient) ListBySubscriptionHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
