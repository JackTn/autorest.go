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

// OptionalBooleanLiteralClient contains the methods for the Type.Property.Optional namespace.
// Don't use this type directly, use [OptionalClient.NewOptionalBooleanLiteralClient] instead.
type OptionalBooleanLiteralClient struct {
	internal *azcore.Client
}

// GetAll - Get models that will return all properties in the model
//   - options - OptionalBooleanLiteralClientGetAllOptions contains the optional parameters for the OptionalBooleanLiteralClient.GetAll
//     method.
func (client *OptionalBooleanLiteralClient) GetAll(ctx context.Context, options *OptionalBooleanLiteralClientGetAllOptions) (OptionalBooleanLiteralClientGetAllResponse, error) {
	var err error
	const operationName = "OptionalBooleanLiteralClient.GetAll"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getAllCreateRequest(ctx, options)
	if err != nil {
		return OptionalBooleanLiteralClientGetAllResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OptionalBooleanLiteralClientGetAllResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OptionalBooleanLiteralClientGetAllResponse{}, err
	}
	resp, err := client.getAllHandleResponse(httpResp)
	return resp, err
}

// getAllCreateRequest creates the GetAll request.
func (client *OptionalBooleanLiteralClient) getAllCreateRequest(ctx context.Context, _ *OptionalBooleanLiteralClientGetAllOptions) (*policy.Request, error) {
	urlPath := "/type/property/optional/boolean/literal/all"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getAllHandleResponse handles the GetAll response.
func (client *OptionalBooleanLiteralClient) getAllHandleResponse(resp *http.Response) (OptionalBooleanLiteralClientGetAllResponse, error) {
	result := OptionalBooleanLiteralClientGetAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BooleanLiteralProperty); err != nil {
		return OptionalBooleanLiteralClientGetAllResponse{}, err
	}
	return result, nil
}

// GetDefault - Get models that will return the default object
//   - options - OptionalBooleanLiteralClientGetDefaultOptions contains the optional parameters for the OptionalBooleanLiteralClient.GetDefault
//     method.
func (client *OptionalBooleanLiteralClient) GetDefault(ctx context.Context, options *OptionalBooleanLiteralClientGetDefaultOptions) (OptionalBooleanLiteralClientGetDefaultResponse, error) {
	var err error
	const operationName = "OptionalBooleanLiteralClient.GetDefault"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getDefaultCreateRequest(ctx, options)
	if err != nil {
		return OptionalBooleanLiteralClientGetDefaultResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OptionalBooleanLiteralClientGetDefaultResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OptionalBooleanLiteralClientGetDefaultResponse{}, err
	}
	resp, err := client.getDefaultHandleResponse(httpResp)
	return resp, err
}

// getDefaultCreateRequest creates the GetDefault request.
func (client *OptionalBooleanLiteralClient) getDefaultCreateRequest(ctx context.Context, _ *OptionalBooleanLiteralClientGetDefaultOptions) (*policy.Request, error) {
	urlPath := "/type/property/optional/boolean/literal/default"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDefaultHandleResponse handles the GetDefault response.
func (client *OptionalBooleanLiteralClient) getDefaultHandleResponse(resp *http.Response) (OptionalBooleanLiteralClientGetDefaultResponse, error) {
	result := OptionalBooleanLiteralClientGetDefaultResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BooleanLiteralProperty); err != nil {
		return OptionalBooleanLiteralClientGetDefaultResponse{}, err
	}
	return result, nil
}

// PutAll - Put a body with all properties present.
//   - options - OptionalBooleanLiteralClientPutAllOptions contains the optional parameters for the OptionalBooleanLiteralClient.PutAll
//     method.
func (client *OptionalBooleanLiteralClient) PutAll(ctx context.Context, body BooleanLiteralProperty, options *OptionalBooleanLiteralClientPutAllOptions) (OptionalBooleanLiteralClientPutAllResponse, error) {
	var err error
	const operationName = "OptionalBooleanLiteralClient.PutAll"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putAllCreateRequest(ctx, body, options)
	if err != nil {
		return OptionalBooleanLiteralClientPutAllResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OptionalBooleanLiteralClientPutAllResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OptionalBooleanLiteralClientPutAllResponse{}, err
	}
	return OptionalBooleanLiteralClientPutAllResponse{}, nil
}

// putAllCreateRequest creates the PutAll request.
func (client *OptionalBooleanLiteralClient) putAllCreateRequest(ctx context.Context, body BooleanLiteralProperty, _ *OptionalBooleanLiteralClientPutAllOptions) (*policy.Request, error) {
	urlPath := "/type/property/optional/boolean/literal/all"
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
//   - options - OptionalBooleanLiteralClientPutDefaultOptions contains the optional parameters for the OptionalBooleanLiteralClient.PutDefault
//     method.
func (client *OptionalBooleanLiteralClient) PutDefault(ctx context.Context, body BooleanLiteralProperty, options *OptionalBooleanLiteralClientPutDefaultOptions) (OptionalBooleanLiteralClientPutDefaultResponse, error) {
	var err error
	const operationName = "OptionalBooleanLiteralClient.PutDefault"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putDefaultCreateRequest(ctx, body, options)
	if err != nil {
		return OptionalBooleanLiteralClientPutDefaultResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OptionalBooleanLiteralClientPutDefaultResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OptionalBooleanLiteralClientPutDefaultResponse{}, err
	}
	return OptionalBooleanLiteralClientPutDefaultResponse{}, nil
}

// putDefaultCreateRequest creates the PutDefault request.
func (client *OptionalBooleanLiteralClient) putDefaultCreateRequest(ctx context.Context, body BooleanLiteralProperty, _ *OptionalBooleanLiteralClientPutDefaultOptions) (*policy.Request, error) {
	urlPath := "/type/property/optional/boolean/literal/default"
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
