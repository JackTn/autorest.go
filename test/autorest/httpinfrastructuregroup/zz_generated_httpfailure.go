// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package httpinfrastructuregroup

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
)

// HTTPFailureOperations contains the methods for the HTTPFailure group.
type HTTPFailureOperations interface {
	// GetEmptyError - Get empty error form server
	GetEmptyError(ctx context.Context, options *HTTPFailureGetEmptyErrorOptions) (*BoolResponse, error)
	// GetNoModelEmpty - Get empty response from server
	GetNoModelEmpty(ctx context.Context, options *HTTPFailureGetNoModelEmptyOptions) (*BoolResponse, error)
	// GetNoModelError - Get empty error form server
	GetNoModelError(ctx context.Context, options *HTTPFailureGetNoModelErrorOptions) (*BoolResponse, error)
}

// HTTPFailureClient implements the HTTPFailureOperations interface.
// Don't use this type directly, use NewHTTPFailureClient() instead.
type HTTPFailureClient struct {
	con *Connection
}

// NewHTTPFailureClient creates a new instance of HTTPFailureClient with the specified values.
func NewHTTPFailureClient(con *Connection) HTTPFailureOperations {
	return &HTTPFailureClient{con: con}
}

// Pipeline returns the pipeline associated with this client.
func (client *HTTPFailureClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// GetEmptyError - Get empty error form server
func (client *HTTPFailureClient) GetEmptyError(ctx context.Context, options *HTTPFailureGetEmptyErrorOptions) (*BoolResponse, error) {
	req, err := client.GetEmptyErrorCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetEmptyErrorHandleError(resp)
	}
	result, err := client.GetEmptyErrorHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetEmptyErrorCreateRequest creates the GetEmptyError request.
func (client *HTTPFailureClient) GetEmptyErrorCreateRequest(ctx context.Context, options *HTTPFailureGetEmptyErrorOptions) (*azcore.Request, error) {
	urlPath := "/http/failure/emptybody/error"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetEmptyErrorHandleResponse handles the GetEmptyError response.
func (client *HTTPFailureClient) GetEmptyErrorHandleResponse(resp *azcore.Response) (*BoolResponse, error) {
	result := BoolResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetEmptyErrorHandleError handles the GetEmptyError error response.
func (client *HTTPFailureClient) GetEmptyErrorHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetNoModelEmpty - Get empty response from server
func (client *HTTPFailureClient) GetNoModelEmpty(ctx context.Context, options *HTTPFailureGetNoModelEmptyOptions) (*BoolResponse, error) {
	req, err := client.GetNoModelEmptyCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetNoModelEmptyHandleError(resp)
	}
	result, err := client.GetNoModelEmptyHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetNoModelEmptyCreateRequest creates the GetNoModelEmpty request.
func (client *HTTPFailureClient) GetNoModelEmptyCreateRequest(ctx context.Context, options *HTTPFailureGetNoModelEmptyOptions) (*azcore.Request, error) {
	urlPath := "/http/failure/nomodel/empty"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetNoModelEmptyHandleResponse handles the GetNoModelEmpty response.
func (client *HTTPFailureClient) GetNoModelEmptyHandleResponse(resp *azcore.Response) (*BoolResponse, error) {
	result := BoolResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetNoModelEmptyHandleError handles the GetNoModelEmpty error response.
func (client *HTTPFailureClient) GetNoModelEmptyHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// GetNoModelError - Get empty error form server
func (client *HTTPFailureClient) GetNoModelError(ctx context.Context, options *HTTPFailureGetNoModelErrorOptions) (*BoolResponse, error) {
	req, err := client.GetNoModelErrorCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetNoModelErrorHandleError(resp)
	}
	result, err := client.GetNoModelErrorHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetNoModelErrorCreateRequest creates the GetNoModelError request.
func (client *HTTPFailureClient) GetNoModelErrorCreateRequest(ctx context.Context, options *HTTPFailureGetNoModelErrorOptions) (*azcore.Request, error) {
	urlPath := "/http/failure/nomodel/error"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetNoModelErrorHandleResponse handles the GetNoModelError response.
func (client *HTTPFailureClient) GetNoModelErrorHandleResponse(resp *azcore.Response) (*BoolResponse, error) {
	result := BoolResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetNoModelErrorHandleError handles the GetNoModelError error response.
func (client *HTTPFailureClient) GetNoModelErrorHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
