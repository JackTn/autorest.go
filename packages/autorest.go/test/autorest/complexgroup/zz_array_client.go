// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ArrayClient contains the methods for the Array group.
// Don't use this type directly, use a constructor function instead.
type ArrayClient struct {
	internal *azcore.Client
}

// GetEmpty - Get complex types with array property which is empty
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-02-29
//   - options - ArrayClientGetEmptyOptions contains the optional parameters for the ArrayClient.GetEmpty method.
func (client *ArrayClient) GetEmpty(ctx context.Context, options *ArrayClientGetEmptyOptions) (ArrayClientGetEmptyResponse, error) {
	var err error
	const operationName = "ArrayClient.GetEmpty"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getEmptyCreateRequest(ctx, options)
	if err != nil {
		return ArrayClientGetEmptyResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ArrayClientGetEmptyResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ArrayClientGetEmptyResponse{}, err
	}
	resp, err := client.getEmptyHandleResponse(httpResp)
	return resp, err
}

// getEmptyCreateRequest creates the GetEmpty request.
func (client *ArrayClient) getEmptyCreateRequest(ctx context.Context, _ *ArrayClientGetEmptyOptions) (*policy.Request, error) {
	urlPath := "/complex/array/empty"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getEmptyHandleResponse handles the GetEmpty response.
func (client *ArrayClient) getEmptyHandleResponse(resp *http.Response) (ArrayClientGetEmptyResponse, error) {
	result := ArrayClientGetEmptyResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArrayWrapper); err != nil {
		return ArrayClientGetEmptyResponse{}, err
	}
	return result, nil
}

// GetNotProvided - Get complex types with array property while server doesn't provide a response payload
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-02-29
//   - options - ArrayClientGetNotProvidedOptions contains the optional parameters for the ArrayClient.GetNotProvided method.
func (client *ArrayClient) GetNotProvided(ctx context.Context, options *ArrayClientGetNotProvidedOptions) (ArrayClientGetNotProvidedResponse, error) {
	var err error
	const operationName = "ArrayClient.GetNotProvided"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getNotProvidedCreateRequest(ctx, options)
	if err != nil {
		return ArrayClientGetNotProvidedResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ArrayClientGetNotProvidedResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ArrayClientGetNotProvidedResponse{}, err
	}
	resp, err := client.getNotProvidedHandleResponse(httpResp)
	return resp, err
}

// getNotProvidedCreateRequest creates the GetNotProvided request.
func (client *ArrayClient) getNotProvidedCreateRequest(ctx context.Context, _ *ArrayClientGetNotProvidedOptions) (*policy.Request, error) {
	urlPath := "/complex/array/notprovided"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getNotProvidedHandleResponse handles the GetNotProvided response.
func (client *ArrayClient) getNotProvidedHandleResponse(resp *http.Response) (ArrayClientGetNotProvidedResponse, error) {
	result := ArrayClientGetNotProvidedResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArrayWrapper); err != nil {
		return ArrayClientGetNotProvidedResponse{}, err
	}
	return result, nil
}

// GetValid - Get complex types with array property
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-02-29
//   - options - ArrayClientGetValidOptions contains the optional parameters for the ArrayClient.GetValid method.
func (client *ArrayClient) GetValid(ctx context.Context, options *ArrayClientGetValidOptions) (ArrayClientGetValidResponse, error) {
	var err error
	const operationName = "ArrayClient.GetValid"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getValidCreateRequest(ctx, options)
	if err != nil {
		return ArrayClientGetValidResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ArrayClientGetValidResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ArrayClientGetValidResponse{}, err
	}
	resp, err := client.getValidHandleResponse(httpResp)
	return resp, err
}

// getValidCreateRequest creates the GetValid request.
func (client *ArrayClient) getValidCreateRequest(ctx context.Context, _ *ArrayClientGetValidOptions) (*policy.Request, error) {
	urlPath := "/complex/array/valid"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getValidHandleResponse handles the GetValid response.
func (client *ArrayClient) getValidHandleResponse(resp *http.Response) (ArrayClientGetValidResponse, error) {
	result := ArrayClientGetValidResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArrayWrapper); err != nil {
		return ArrayClientGetValidResponse{}, err
	}
	return result, nil
}

// PutEmpty - Put complex types with array property which is empty
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-02-29
//   - complexBody - Please put an empty array
//   - options - ArrayClientPutEmptyOptions contains the optional parameters for the ArrayClient.PutEmpty method.
func (client *ArrayClient) PutEmpty(ctx context.Context, complexBody ArrayWrapper, options *ArrayClientPutEmptyOptions) (ArrayClientPutEmptyResponse, error) {
	var err error
	const operationName = "ArrayClient.PutEmpty"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putEmptyCreateRequest(ctx, complexBody, options)
	if err != nil {
		return ArrayClientPutEmptyResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ArrayClientPutEmptyResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ArrayClientPutEmptyResponse{}, err
	}
	return ArrayClientPutEmptyResponse{}, nil
}

// putEmptyCreateRequest creates the PutEmpty request.
func (client *ArrayClient) putEmptyCreateRequest(ctx context.Context, complexBody ArrayWrapper, _ *ArrayClientPutEmptyOptions) (*policy.Request, error) {
	urlPath := "/complex/array/empty"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, complexBody); err != nil {
		return nil, err
	}
	return req, nil
}

// PutValid - Put complex types with array property
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-02-29
//   - complexBody - Please put an array with 4 items: "1, 2, 3, 4", "", null, "&S#$(*Y", "The quick brown fox jumps over the
//     lazy dog"
//   - options - ArrayClientPutValidOptions contains the optional parameters for the ArrayClient.PutValid method.
func (client *ArrayClient) PutValid(ctx context.Context, complexBody ArrayWrapper, options *ArrayClientPutValidOptions) (ArrayClientPutValidResponse, error) {
	var err error
	const operationName = "ArrayClient.PutValid"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putValidCreateRequest(ctx, complexBody, options)
	if err != nil {
		return ArrayClientPutValidResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ArrayClientPutValidResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ArrayClientPutValidResponse{}, err
	}
	return ArrayClientPutValidResponse{}, nil
}

// putValidCreateRequest creates the PutValid request.
func (client *ArrayClient) putValidCreateRequest(ctx context.Context, complexBody ArrayWrapper, _ *ArrayClientPutValidOptions) (*policy.Request, error) {
	urlPath := "/complex/array/valid"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, complexBody); err != nil {
		return nil, err
	}
	return req, nil
}
