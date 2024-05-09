// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package optionalitygroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// OptionalIntLiteralClient contains the methods for the Type.Property.Optional namespace.
// Don't use this type directly, use [OptionalClient.NewOptionalIntLiteralClient] instead.
type OptionalIntLiteralClient struct {
	internal *azcore.Client
}

// GetAll - Get models that will return all properties in the model
//   - options - OptionalIntLiteralClientGetAllOptions contains the optional parameters for the OptionalIntLiteralClient.GetAll
//     method.
func (client *OptionalIntLiteralClient) GetAll(ctx context.Context, options *OptionalIntLiteralClientGetAllOptions) (OptionalIntLiteralClientGetAllResponse, error) {
	var err error
	const operationName = "OptionalIntLiteralClient.GetAll"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getAllCreateRequest(ctx, options)
	if err != nil {
		return OptionalIntLiteralClientGetAllResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OptionalIntLiteralClientGetAllResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OptionalIntLiteralClientGetAllResponse{}, err
	}
	resp, err := client.getAllHandleResponse(httpResp)
	return resp, err
}

// getAllCreateRequest creates the GetAll request.
func (client *OptionalIntLiteralClient) getAllCreateRequest(ctx context.Context, _ *OptionalIntLiteralClientGetAllOptions) (*policy.Request, error) {
	urlPath := "/type/property/optional/int/literal/all"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getAllHandleResponse handles the GetAll response.
func (client *OptionalIntLiteralClient) getAllHandleResponse(resp *http.Response) (OptionalIntLiteralClientGetAllResponse, error) {
	result := OptionalIntLiteralClientGetAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntLiteralProperty); err != nil {
		return OptionalIntLiteralClientGetAllResponse{}, err
	}
	return result, nil
}

// GetDefault - Get models that will return the default object
//   - options - OptionalIntLiteralClientGetDefaultOptions contains the optional parameters for the OptionalIntLiteralClient.GetDefault
//     method.
func (client *OptionalIntLiteralClient) GetDefault(ctx context.Context, options *OptionalIntLiteralClientGetDefaultOptions) (OptionalIntLiteralClientGetDefaultResponse, error) {
	var err error
	const operationName = "OptionalIntLiteralClient.GetDefault"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getDefaultCreateRequest(ctx, options)
	if err != nil {
		return OptionalIntLiteralClientGetDefaultResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OptionalIntLiteralClientGetDefaultResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OptionalIntLiteralClientGetDefaultResponse{}, err
	}
	resp, err := client.getDefaultHandleResponse(httpResp)
	return resp, err
}

// getDefaultCreateRequest creates the GetDefault request.
func (client *OptionalIntLiteralClient) getDefaultCreateRequest(ctx context.Context, _ *OptionalIntLiteralClientGetDefaultOptions) (*policy.Request, error) {
	urlPath := "/type/property/optional/int/literal/default"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDefaultHandleResponse handles the GetDefault response.
func (client *OptionalIntLiteralClient) getDefaultHandleResponse(resp *http.Response) (OptionalIntLiteralClientGetDefaultResponse, error) {
	result := OptionalIntLiteralClientGetDefaultResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntLiteralProperty); err != nil {
		return OptionalIntLiteralClientGetDefaultResponse{}, err
	}
	return result, nil
}

// PutAll - Put a body with all properties present.
//   - options - OptionalIntLiteralClientPutAllOptions contains the optional parameters for the OptionalIntLiteralClient.PutAll
//     method.
func (client *OptionalIntLiteralClient) PutAll(ctx context.Context, body IntLiteralProperty, options *OptionalIntLiteralClientPutAllOptions) (OptionalIntLiteralClientPutAllResponse, error) {
	var err error
	const operationName = "OptionalIntLiteralClient.PutAll"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putAllCreateRequest(ctx, body, options)
	if err != nil {
		return OptionalIntLiteralClientPutAllResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OptionalIntLiteralClientPutAllResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OptionalIntLiteralClientPutAllResponse{}, err
	}
	return OptionalIntLiteralClientPutAllResponse{}, nil
}

// putAllCreateRequest creates the PutAll request.
func (client *OptionalIntLiteralClient) putAllCreateRequest(ctx context.Context, body IntLiteralProperty, _ *OptionalIntLiteralClientPutAllOptions) (*policy.Request, error) {
	urlPath := "/type/property/optional/int/literal/all"
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

// PutDefault - Put a body with default properties.
//   - options - OptionalIntLiteralClientPutDefaultOptions contains the optional parameters for the OptionalIntLiteralClient.PutDefault
//     method.
func (client *OptionalIntLiteralClient) PutDefault(ctx context.Context, body IntLiteralProperty, options *OptionalIntLiteralClientPutDefaultOptions) (OptionalIntLiteralClientPutDefaultResponse, error) {
	var err error
	const operationName = "OptionalIntLiteralClient.PutDefault"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putDefaultCreateRequest(ctx, body, options)
	if err != nil {
		return OptionalIntLiteralClientPutDefaultResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OptionalIntLiteralClientPutDefaultResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OptionalIntLiteralClientPutDefaultResponse{}, err
	}
	return OptionalIntLiteralClientPutDefaultResponse{}, nil
}

// putDefaultCreateRequest creates the PutDefault request.
func (client *OptionalIntLiteralClient) putDefaultCreateRequest(ctx context.Context, body IntLiteralProperty, _ *OptionalIntLiteralClientPutDefaultOptions) (*policy.Request, error) {
	urlPath := "/type/property/optional/int/literal/default"
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
