// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package arraygroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// Int64ValueClient - Array of int64 values
// Don't use this type directly, use [ArrayClient.NewInt64ValueClient] instead.
type Int64ValueClient struct {
	internal *azcore.Client
}

// - options - Int64ValueClientGetOptions contains the optional parameters for the Int64ValueClient.Get method.
func (client *Int64ValueClient) Get(ctx context.Context, options *Int64ValueClientGetOptions) (Int64ValueClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return Int64ValueClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return Int64ValueClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return Int64ValueClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *Int64ValueClient) getCreateRequest(ctx context.Context, options *Int64ValueClientGetOptions) (*policy.Request, error) {
	urlPath := "/type/array/int64"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *Int64ValueClient) getHandleResponse(resp *http.Response) (Int64ValueClientGetResponse, error) {
	result := Int64ValueClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return Int64ValueClientGetResponse{}, err
	}
	return result, nil
}

// - options - Int64ValueClientPutOptions contains the optional parameters for the Int64ValueClient.Put method.
func (client *Int64ValueClient) Put(ctx context.Context, body []int64, options *Int64ValueClientPutOptions) (Int64ValueClientPutResponse, error) {
	var err error
	req, err := client.putCreateRequest(ctx, body, options)
	if err != nil {
		return Int64ValueClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return Int64ValueClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return Int64ValueClientPutResponse{}, err
	}
	return Int64ValueClientPutResponse{}, nil
}

// putCreateRequest creates the Put request.
func (client *Int64ValueClient) putCreateRequest(ctx context.Context, body []int64, options *Int64ValueClientPutOptions) (*policy.Request, error) {
	urlPath := "/type/array/int64"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}
