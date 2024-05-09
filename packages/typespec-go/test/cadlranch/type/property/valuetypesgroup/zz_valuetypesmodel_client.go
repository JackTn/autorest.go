// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package valuetypesgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ValueTypesModelClient contains the methods for the Type.Property.ValueTypes namespace.
// Don't use this type directly, use [ValueTypesClient.NewValueTypesModelClient] instead.
type ValueTypesModelClient struct {
	internal *azcore.Client
}

// Get - Get call
//   - options - ValueTypesModelClientGetOptions contains the optional parameters for the ValueTypesModelClient.Get method.
func (client *ValueTypesModelClient) Get(ctx context.Context, options *ValueTypesModelClientGetOptions) (ValueTypesModelClientGetResponse, error) {
	var err error
	const operationName = "ValueTypesModelClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return ValueTypesModelClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ValueTypesModelClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ValueTypesModelClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ValueTypesModelClient) getCreateRequest(ctx context.Context, _ *ValueTypesModelClientGetOptions) (*policy.Request, error) {
	urlPath := "/type/property/value-types/model"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ValueTypesModelClient) getHandleResponse(resp *http.Response) (ValueTypesModelClientGetResponse, error) {
	result := ValueTypesModelClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ModelProperty); err != nil {
		return ValueTypesModelClientGetResponse{}, err
	}
	return result, nil
}

// Put - Put operation
//   - body - body
//   - options - ValueTypesModelClientPutOptions contains the optional parameters for the ValueTypesModelClient.Put method.
func (client *ValueTypesModelClient) Put(ctx context.Context, body ModelProperty, options *ValueTypesModelClientPutOptions) (ValueTypesModelClientPutResponse, error) {
	var err error
	const operationName = "ValueTypesModelClient.Put"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putCreateRequest(ctx, body, options)
	if err != nil {
		return ValueTypesModelClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ValueTypesModelClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ValueTypesModelClientPutResponse{}, err
	}
	return ValueTypesModelClientPutResponse{}, nil
}

// putCreateRequest creates the Put request.
func (client *ValueTypesModelClient) putCreateRequest(ctx context.Context, body ModelProperty, _ *ValueTypesModelClientPutOptions) (*policy.Request, error) {
	urlPath := "/type/property/value-types/model"
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
