// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// VirtualMachineExtensionsOperations contains the methods for the VirtualMachineExtensions group.
type VirtualMachineExtensionsOperations interface {
	// BeginCreateOrUpdate - The operation to create or update the extension.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, vmName string, vmExtensionName string, extensionParameters VirtualMachineExtension, options *VirtualMachineExtensionsCreateOrUpdateOptions) (*VirtualMachineExtensionPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (VirtualMachineExtensionPoller, error)
	// BeginDelete - The operation to delete the extension.
	BeginDelete(ctx context.Context, resourceGroupName string, vmName string, vmExtensionName string, options *VirtualMachineExtensionsDeleteOptions) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - The operation to get the extension.
	Get(ctx context.Context, resourceGroupName string, vmName string, vmExtensionName string, options *VirtualMachineExtensionsGetOptions) (*VirtualMachineExtensionResponse, error)
	// List - The operation to get all extensions of a Virtual Machine.
	List(ctx context.Context, resourceGroupName string, vmName string, options *VirtualMachineExtensionsListOptions) (*VirtualMachineExtensionsListResultResponse, error)
	// BeginUpdate - The operation to update the extension.
	BeginUpdate(ctx context.Context, resourceGroupName string, vmName string, vmExtensionName string, extensionParameters VirtualMachineExtensionUpdate, options *VirtualMachineExtensionsUpdateOptions) (*VirtualMachineExtensionPollerResponse, error)
	// ResumeUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeUpdate(token string) (VirtualMachineExtensionPoller, error)
}

// VirtualMachineExtensionsClient implements the VirtualMachineExtensionsOperations interface.
// Don't use this type directly, use NewVirtualMachineExtensionsClient() instead.
type VirtualMachineExtensionsClient struct {
	con            *Connection
	subscriptionID string
}

// NewVirtualMachineExtensionsClient creates a new instance of VirtualMachineExtensionsClient with the specified values.
func NewVirtualMachineExtensionsClient(con *Connection, subscriptionID string) VirtualMachineExtensionsOperations {
	return &VirtualMachineExtensionsClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *VirtualMachineExtensionsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

func (client *VirtualMachineExtensionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, vmName string, vmExtensionName string, extensionParameters VirtualMachineExtension, options *VirtualMachineExtensionsCreateOrUpdateOptions) (*VirtualMachineExtensionPollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, vmName, vmExtensionName, extensionParameters, options)
	if err != nil {
		return nil, err
	}
	result := &VirtualMachineExtensionPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualMachineExtensionsClient.CreateOrUpdate", "", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &virtualMachineExtensionPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*VirtualMachineExtensionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *VirtualMachineExtensionsClient) ResumeCreateOrUpdate(token string) (VirtualMachineExtensionPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualMachineExtensionsClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &virtualMachineExtensionPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - The operation to create or update the extension.
func (client *VirtualMachineExtensionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, vmName string, vmExtensionName string, extensionParameters VirtualMachineExtension, options *VirtualMachineExtensionsCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, vmName, vmExtensionName, extensionParameters, options)
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
func (client *VirtualMachineExtensionsClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, vmName string, vmExtensionName string, extensionParameters VirtualMachineExtension, options *VirtualMachineExtensionsCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/extensions/{vmExtensionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vmName}", url.PathEscape(vmName))
	urlPath = strings.ReplaceAll(urlPath, "{vmExtensionName}", url.PathEscape(vmExtensionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(extensionParameters)
}

// CreateOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *VirtualMachineExtensionsClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*VirtualMachineExtensionResponse, error) {
	result := VirtualMachineExtensionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualMachineExtension)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *VirtualMachineExtensionsClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

func (client *VirtualMachineExtensionsClient) BeginDelete(ctx context.Context, resourceGroupName string, vmName string, vmExtensionName string, options *VirtualMachineExtensionsDeleteOptions) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, vmName, vmExtensionName, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualMachineExtensionsClient.Delete", "", resp, client.DeleteHandleError)
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

func (client *VirtualMachineExtensionsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualMachineExtensionsClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - The operation to delete the extension.
func (client *VirtualMachineExtensionsClient) Delete(ctx context.Context, resourceGroupName string, vmName string, vmExtensionName string, options *VirtualMachineExtensionsDeleteOptions) (*azcore.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, vmName, vmExtensionName, options)
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
func (client *VirtualMachineExtensionsClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, vmName string, vmExtensionName string, options *VirtualMachineExtensionsDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/extensions/{vmExtensionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vmName}", url.PathEscape(vmName))
	urlPath = strings.ReplaceAll(urlPath, "{vmExtensionName}", url.PathEscape(vmExtensionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	return req, nil
}

// DeleteHandleError handles the Delete error response.
func (client *VirtualMachineExtensionsClient) DeleteHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Get - The operation to get the extension.
func (client *VirtualMachineExtensionsClient) Get(ctx context.Context, resourceGroupName string, vmName string, vmExtensionName string, options *VirtualMachineExtensionsGetOptions) (*VirtualMachineExtensionResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, vmName, vmExtensionName, options)
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
func (client *VirtualMachineExtensionsClient) GetCreateRequest(ctx context.Context, resourceGroupName string, vmName string, vmExtensionName string, options *VirtualMachineExtensionsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/extensions/{vmExtensionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vmName}", url.PathEscape(vmName))
	urlPath = strings.ReplaceAll(urlPath, "{vmExtensionName}", url.PathEscape(vmExtensionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	if options != nil && options.Expand != nil {
		query.Set("$expand", *options.Expand)
	}
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *VirtualMachineExtensionsClient) GetHandleResponse(resp *azcore.Response) (*VirtualMachineExtensionResponse, error) {
	result := VirtualMachineExtensionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualMachineExtension)
}

// GetHandleError handles the Get error response.
func (client *VirtualMachineExtensionsClient) GetHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// List - The operation to get all extensions of a Virtual Machine.
func (client *VirtualMachineExtensionsClient) List(ctx context.Context, resourceGroupName string, vmName string, options *VirtualMachineExtensionsListOptions) (*VirtualMachineExtensionsListResultResponse, error) {
	req, err := client.ListCreateRequest(ctx, resourceGroupName, vmName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ListHandleError(resp)
	}
	result, err := client.ListHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListCreateRequest creates the List request.
func (client *VirtualMachineExtensionsClient) ListCreateRequest(ctx context.Context, resourceGroupName string, vmName string, options *VirtualMachineExtensionsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/extensions"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vmName}", url.PathEscape(vmName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	if options != nil && options.Expand != nil {
		query.Set("$expand", *options.Expand)
	}
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListHandleResponse handles the List response.
func (client *VirtualMachineExtensionsClient) ListHandleResponse(resp *azcore.Response) (*VirtualMachineExtensionsListResultResponse, error) {
	result := VirtualMachineExtensionsListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualMachineExtensionsListResult)
}

// ListHandleError handles the List error response.
func (client *VirtualMachineExtensionsClient) ListHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

func (client *VirtualMachineExtensionsClient) BeginUpdate(ctx context.Context, resourceGroupName string, vmName string, vmExtensionName string, extensionParameters VirtualMachineExtensionUpdate, options *VirtualMachineExtensionsUpdateOptions) (*VirtualMachineExtensionPollerResponse, error) {
	resp, err := client.Update(ctx, resourceGroupName, vmName, vmExtensionName, extensionParameters, options)
	if err != nil {
		return nil, err
	}
	result := &VirtualMachineExtensionPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualMachineExtensionsClient.Update", "", resp, client.UpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &virtualMachineExtensionPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*VirtualMachineExtensionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *VirtualMachineExtensionsClient) ResumeUpdate(token string) (VirtualMachineExtensionPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualMachineExtensionsClient.Update", token, client.UpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &virtualMachineExtensionPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Update - The operation to update the extension.
func (client *VirtualMachineExtensionsClient) Update(ctx context.Context, resourceGroupName string, vmName string, vmExtensionName string, extensionParameters VirtualMachineExtensionUpdate, options *VirtualMachineExtensionsUpdateOptions) (*azcore.Response, error) {
	req, err := client.UpdateCreateRequest(ctx, resourceGroupName, vmName, vmExtensionName, extensionParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.UpdateHandleError(resp)
	}
	return resp, nil
}

// UpdateCreateRequest creates the Update request.
func (client *VirtualMachineExtensionsClient) UpdateCreateRequest(ctx context.Context, resourceGroupName string, vmName string, vmExtensionName string, extensionParameters VirtualMachineExtensionUpdate, options *VirtualMachineExtensionsUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/extensions/{vmExtensionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vmName}", url.PathEscape(vmName))
	urlPath = strings.ReplaceAll(urlPath, "{vmExtensionName}", url.PathEscape(vmExtensionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(extensionParameters)
}

// UpdateHandleResponse handles the Update response.
func (client *VirtualMachineExtensionsClient) UpdateHandleResponse(resp *azcore.Response) (*VirtualMachineExtensionResponse, error) {
	result := VirtualMachineExtensionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualMachineExtension)
}

// UpdateHandleError handles the Update error response.
func (client *VirtualMachineExtensionsClient) UpdateHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
