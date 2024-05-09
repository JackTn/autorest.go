// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package nonstringenumgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// FloatClient contains the methods for the Float group.
// Don't use this type directly, use a constructor function instead.
type FloatClient struct {
	internal *azcore.Client
}

// Get - Get a float enum
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2.0-preview
//   - options - FloatClientGetOptions contains the optional parameters for the FloatClient.Get method.
func (client *FloatClient) Get(ctx context.Context, options *FloatClientGetOptions) (FloatClientGetResponse, error) {
	var err error
	const operationName = "FloatClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return FloatClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FloatClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return FloatClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *FloatClient) getCreateRequest(ctx context.Context, _ *FloatClientGetOptions) (*policy.Request, error) {
	urlPath := "/nonStringEnums/float/get"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FloatClient) getHandleResponse(resp *http.Response) (FloatClientGetResponse, error) {
	result := FloatClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return FloatClientGetResponse{}, err
	}
	return result, nil
}

// Put - Put a float enum
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2.0-preview
//   - input - Input float enum.
//   - options - FloatClientPutOptions contains the optional parameters for the FloatClient.Put method.
func (client *FloatClient) Put(ctx context.Context, input FloatEnum, options *FloatClientPutOptions) (FloatClientPutResponse, error) {
	var err error
	const operationName = "FloatClient.Put"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putCreateRequest(ctx, input, options)
	if err != nil {
		return FloatClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FloatClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return FloatClientPutResponse{}, err
	}
	resp, err := client.putHandleResponse(httpResp)
	return resp, err
}

// putCreateRequest creates the Put request.
func (client *FloatClient) putCreateRequest(ctx context.Context, input FloatEnum, _ *FloatClientPutOptions) (*policy.Request, error) {
	urlPath := "/nonStringEnums/float/put"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, input); err != nil {
		return nil, err
	}
	return req, nil
}

// putHandleResponse handles the Put response.
func (client *FloatClient) putHandleResponse(resp *http.Response) (FloatClientPutResponse, error) {
	result := FloatClientPutResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return FloatClientPutResponse{}, err
	}
	return result, nil
}
