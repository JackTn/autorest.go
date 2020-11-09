// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// VirtualMachineRunCommandsOperations contains the methods for the VirtualMachineRunCommands group.
type VirtualMachineRunCommandsOperations interface {
	// Get - Gets specific run command for a subscription in a location.
	Get(ctx context.Context, location string, commandId string, options *VirtualMachineRunCommandsGetOptions) (*RunCommandDocumentResponse, error)
	// List - Lists all available run commands for a subscription in a location.
	List(location string, options *VirtualMachineRunCommandsListOptions) RunCommandListResultPager
}

// VirtualMachineRunCommandsClient implements the VirtualMachineRunCommandsOperations interface.
// Don't use this type directly, use NewVirtualMachineRunCommandsClient() instead.
type VirtualMachineRunCommandsClient struct {
	con            *Connection
	subscriptionID string
}

// NewVirtualMachineRunCommandsClient creates a new instance of VirtualMachineRunCommandsClient with the specified values.
func NewVirtualMachineRunCommandsClient(con *Connection, subscriptionID string) VirtualMachineRunCommandsOperations {
	return &VirtualMachineRunCommandsClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *VirtualMachineRunCommandsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// Get - Gets specific run command for a subscription in a location.
func (client *VirtualMachineRunCommandsClient) Get(ctx context.Context, location string, commandId string, options *VirtualMachineRunCommandsGetOptions) (*RunCommandDocumentResponse, error) {
	req, err := client.GetCreateRequest(ctx, location, commandId, options)
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
func (client *VirtualMachineRunCommandsClient) GetCreateRequest(ctx context.Context, location string, commandId string, options *VirtualMachineRunCommandsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/runCommands/{commandId}"
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	urlPath = strings.ReplaceAll(urlPath, "{commandId}", url.PathEscape(commandId))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json, text/json")
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *VirtualMachineRunCommandsClient) GetHandleResponse(resp *azcore.Response) (*RunCommandDocumentResponse, error) {
	result := RunCommandDocumentResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.RunCommandDocument)
}

// GetHandleError handles the Get error response.
func (client *VirtualMachineRunCommandsClient) GetHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// List - Lists all available run commands for a subscription in a location.
func (client *VirtualMachineRunCommandsClient) List(location string, options *VirtualMachineRunCommandsListOptions) RunCommandListResultPager {
	return &runCommandListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, location, options)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *RunCommandListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.RunCommandListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListCreateRequest creates the List request.
func (client *VirtualMachineRunCommandsClient) ListCreateRequest(ctx context.Context, location string, options *VirtualMachineRunCommandsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/runCommands"
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json, text/json")
	return req, nil
}

// ListHandleResponse handles the List response.
func (client *VirtualMachineRunCommandsClient) ListHandleResponse(resp *azcore.Response) (*RunCommandListResultResponse, error) {
	result := RunCommandListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.RunCommandListResult)
}

// ListHandleError handles the List error response.
func (client *VirtualMachineRunCommandsClient) ListHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
