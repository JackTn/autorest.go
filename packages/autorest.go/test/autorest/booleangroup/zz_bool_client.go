// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package booleangroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// BoolClient contains the methods for the Bool group.
// Don't use this type directly, use a constructor function instead.
type BoolClient struct {
	internal *azcore.Client
}

// GetFalse - Get false Boolean value
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - BoolClientGetFalseOptions contains the optional parameters for the BoolClient.GetFalse method.
func (client *BoolClient) GetFalse(ctx context.Context, options *BoolClientGetFalseOptions) (BoolClientGetFalseResponse, error) {
	var err error
	const operationName = "BoolClient.GetFalse"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getFalseCreateRequest(ctx, options)
	if err != nil {
		return BoolClientGetFalseResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BoolClientGetFalseResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BoolClientGetFalseResponse{}, err
	}
	resp, err := client.getFalseHandleResponse(httpResp)
	return resp, err
}

// getFalseCreateRequest creates the GetFalse request.
func (client *BoolClient) getFalseCreateRequest(ctx context.Context, _ *BoolClientGetFalseOptions) (*policy.Request, error) {
	urlPath := "/bool/false"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getFalseHandleResponse handles the GetFalse response.
func (client *BoolClient) getFalseHandleResponse(resp *http.Response) (BoolClientGetFalseResponse, error) {
	result := BoolClientGetFalseResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return BoolClientGetFalseResponse{}, err
	}
	return result, nil
}

// GetInvalid - Get invalid Boolean value
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - BoolClientGetInvalidOptions contains the optional parameters for the BoolClient.GetInvalid method.
func (client *BoolClient) GetInvalid(ctx context.Context, options *BoolClientGetInvalidOptions) (BoolClientGetInvalidResponse, error) {
	var err error
	const operationName = "BoolClient.GetInvalid"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getInvalidCreateRequest(ctx, options)
	if err != nil {
		return BoolClientGetInvalidResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BoolClientGetInvalidResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BoolClientGetInvalidResponse{}, err
	}
	resp, err := client.getInvalidHandleResponse(httpResp)
	return resp, err
}

// getInvalidCreateRequest creates the GetInvalid request.
func (client *BoolClient) getInvalidCreateRequest(ctx context.Context, _ *BoolClientGetInvalidOptions) (*policy.Request, error) {
	urlPath := "/bool/invalid"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getInvalidHandleResponse handles the GetInvalid response.
func (client *BoolClient) getInvalidHandleResponse(resp *http.Response) (BoolClientGetInvalidResponse, error) {
	result := BoolClientGetInvalidResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return BoolClientGetInvalidResponse{}, err
	}
	return result, nil
}

// GetNull - Get null Boolean value
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - BoolClientGetNullOptions contains the optional parameters for the BoolClient.GetNull method.
func (client *BoolClient) GetNull(ctx context.Context, options *BoolClientGetNullOptions) (BoolClientGetNullResponse, error) {
	var err error
	const operationName = "BoolClient.GetNull"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getNullCreateRequest(ctx, options)
	if err != nil {
		return BoolClientGetNullResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BoolClientGetNullResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BoolClientGetNullResponse{}, err
	}
	resp, err := client.getNullHandleResponse(httpResp)
	return resp, err
}

// getNullCreateRequest creates the GetNull request.
func (client *BoolClient) getNullCreateRequest(ctx context.Context, _ *BoolClientGetNullOptions) (*policy.Request, error) {
	urlPath := "/bool/null"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getNullHandleResponse handles the GetNull response.
func (client *BoolClient) getNullHandleResponse(resp *http.Response) (BoolClientGetNullResponse, error) {
	result := BoolClientGetNullResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return BoolClientGetNullResponse{}, err
	}
	return result, nil
}

// GetTrue - Get true Boolean value
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - BoolClientGetTrueOptions contains the optional parameters for the BoolClient.GetTrue method.
func (client *BoolClient) GetTrue(ctx context.Context, options *BoolClientGetTrueOptions) (BoolClientGetTrueResponse, error) {
	var err error
	const operationName = "BoolClient.GetTrue"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getTrueCreateRequest(ctx, options)
	if err != nil {
		return BoolClientGetTrueResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BoolClientGetTrueResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BoolClientGetTrueResponse{}, err
	}
	resp, err := client.getTrueHandleResponse(httpResp)
	return resp, err
}

// getTrueCreateRequest creates the GetTrue request.
func (client *BoolClient) getTrueCreateRequest(ctx context.Context, _ *BoolClientGetTrueOptions) (*policy.Request, error) {
	urlPath := "/bool/true"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getTrueHandleResponse handles the GetTrue response.
func (client *BoolClient) getTrueHandleResponse(resp *http.Response) (BoolClientGetTrueResponse, error) {
	result := BoolClientGetTrueResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return BoolClientGetTrueResponse{}, err
	}
	return result, nil
}

// PutFalse - Set Boolean value false
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - BoolClientPutFalseOptions contains the optional parameters for the BoolClient.PutFalse method.
func (client *BoolClient) PutFalse(ctx context.Context, options *BoolClientPutFalseOptions) (BoolClientPutFalseResponse, error) {
	var err error
	const operationName = "BoolClient.PutFalse"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putFalseCreateRequest(ctx, options)
	if err != nil {
		return BoolClientPutFalseResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BoolClientPutFalseResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BoolClientPutFalseResponse{}, err
	}
	return BoolClientPutFalseResponse{}, nil
}

// putFalseCreateRequest creates the PutFalse request.
func (client *BoolClient) putFalseCreateRequest(ctx context.Context, _ *BoolClientPutFalseOptions) (*policy.Request, error) {
	urlPath := "/bool/false"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, false); err != nil {
		return nil, err
	}
	return req, nil
}

// PutTrue - Set Boolean value true
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - BoolClientPutTrueOptions contains the optional parameters for the BoolClient.PutTrue method.
func (client *BoolClient) PutTrue(ctx context.Context, options *BoolClientPutTrueOptions) (BoolClientPutTrueResponse, error) {
	var err error
	const operationName = "BoolClient.PutTrue"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putTrueCreateRequest(ctx, options)
	if err != nil {
		return BoolClientPutTrueResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BoolClientPutTrueResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BoolClientPutTrueResponse{}, err
	}
	return BoolClientPutTrueResponse{}, nil
}

// putTrueCreateRequest creates the PutTrue request.
func (client *BoolClient) putTrueCreateRequest(ctx context.Context, _ *BoolClientPutTrueOptions) (*policy.Request, error) {
	urlPath := "/bool/true"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, true); err != nil {
		return nil, err
	}
	return req, nil
}
