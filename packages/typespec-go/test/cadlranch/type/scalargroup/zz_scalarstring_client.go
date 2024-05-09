// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package scalargroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ScalarStringClient contains the methods for the Type.Scalar namespace.
// Don't use this type directly, use [ScalarClient.NewScalarStringClient] instead.
type ScalarStringClient struct {
	internal *azcore.Client
}

// Get - get string value
//   - options - ScalarStringClientGetOptions contains the optional parameters for the ScalarStringClient.Get method.
func (client *ScalarStringClient) Get(ctx context.Context, options *ScalarStringClientGetOptions) (ScalarStringClientGetResponse, error) {
	var err error
	const operationName = "ScalarStringClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return ScalarStringClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScalarStringClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ScalarStringClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ScalarStringClient) getCreateRequest(ctx context.Context, _ *ScalarStringClientGetOptions) (*policy.Request, error) {
	urlPath := "/type/scalar/string"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ScalarStringClient) getHandleResponse(resp *http.Response) (ScalarStringClientGetResponse, error) {
	result := ScalarStringClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return ScalarStringClientGetResponse{}, err
	}
	return result, nil
}

// Put - put string value
//   - body - _
//   - options - ScalarStringClientPutOptions contains the optional parameters for the ScalarStringClient.Put method.
func (client *ScalarStringClient) Put(ctx context.Context, body string, options *ScalarStringClientPutOptions) (ScalarStringClientPutResponse, error) {
	var err error
	const operationName = "ScalarStringClient.Put"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putCreateRequest(ctx, body, options)
	if err != nil {
		return ScalarStringClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScalarStringClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ScalarStringClientPutResponse{}, err
	}
	return ScalarStringClientPutResponse{}, nil
}

// putCreateRequest creates the Put request.
func (client *ScalarStringClient) putCreateRequest(ctx context.Context, body string, _ *ScalarStringClientPutOptions) (*policy.Request, error) {
	urlPath := "/type/scalar/string"
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
