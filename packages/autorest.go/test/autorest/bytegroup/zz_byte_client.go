// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package bytegroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ByteClient contains the methods for the Byte group.
// Don't use this type directly, use a constructor function instead.
type ByteClient struct {
	internal *azcore.Client
}

// GetEmpty - Get empty byte value ”
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - ByteClientGetEmptyOptions contains the optional parameters for the ByteClient.GetEmpty method.
func (client *ByteClient) GetEmpty(ctx context.Context, options *ByteClientGetEmptyOptions) (ByteClientGetEmptyResponse, error) {
	var err error
	const operationName = "ByteClient.GetEmpty"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getEmptyCreateRequest(ctx, options)
	if err != nil {
		return ByteClientGetEmptyResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ByteClientGetEmptyResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ByteClientGetEmptyResponse{}, err
	}
	resp, err := client.getEmptyHandleResponse(httpResp)
	return resp, err
}

// getEmptyCreateRequest creates the GetEmpty request.
func (client *ByteClient) getEmptyCreateRequest(ctx context.Context, options *ByteClientGetEmptyOptions) (*policy.Request, error) {
	urlPath := "/byte/empty"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getEmptyHandleResponse handles the GetEmpty response.
func (client *ByteClient) getEmptyHandleResponse(resp *http.Response) (ByteClientGetEmptyResponse, error) {
	result := ByteClientGetEmptyResponse{}
	if err := runtime.UnmarshalAsByteArray(resp, &result.Value, runtime.Base64StdFormat); err != nil {
		return ByteClientGetEmptyResponse{}, err
	}
	return result, nil
}

// GetInvalid - Get invalid byte value ':::SWAGGER::::'
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - ByteClientGetInvalidOptions contains the optional parameters for the ByteClient.GetInvalid method.
func (client *ByteClient) GetInvalid(ctx context.Context, options *ByteClientGetInvalidOptions) (ByteClientGetInvalidResponse, error) {
	var err error
	const operationName = "ByteClient.GetInvalid"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getInvalidCreateRequest(ctx, options)
	if err != nil {
		return ByteClientGetInvalidResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ByteClientGetInvalidResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ByteClientGetInvalidResponse{}, err
	}
	resp, err := client.getInvalidHandleResponse(httpResp)
	return resp, err
}

// getInvalidCreateRequest creates the GetInvalid request.
func (client *ByteClient) getInvalidCreateRequest(ctx context.Context, options *ByteClientGetInvalidOptions) (*policy.Request, error) {
	urlPath := "/byte/invalid"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getInvalidHandleResponse handles the GetInvalid response.
func (client *ByteClient) getInvalidHandleResponse(resp *http.Response) (ByteClientGetInvalidResponse, error) {
	result := ByteClientGetInvalidResponse{}
	if err := runtime.UnmarshalAsByteArray(resp, &result.Value, runtime.Base64StdFormat); err != nil {
		return ByteClientGetInvalidResponse{}, err
	}
	return result, nil
}

// GetNonASCII - Get non-ascii byte string hex(FF FE FD FC FB FA F9 F8 F7 F6)
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - ByteClientGetNonASCIIOptions contains the optional parameters for the ByteClient.GetNonASCII method.
func (client *ByteClient) GetNonASCII(ctx context.Context, options *ByteClientGetNonASCIIOptions) (ByteClientGetNonASCIIResponse, error) {
	var err error
	const operationName = "ByteClient.GetNonASCII"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getNonASCIICreateRequest(ctx, options)
	if err != nil {
		return ByteClientGetNonASCIIResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ByteClientGetNonASCIIResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ByteClientGetNonASCIIResponse{}, err
	}
	resp, err := client.getNonASCIIHandleResponse(httpResp)
	return resp, err
}

// getNonASCIICreateRequest creates the GetNonASCII request.
func (client *ByteClient) getNonASCIICreateRequest(ctx context.Context, options *ByteClientGetNonASCIIOptions) (*policy.Request, error) {
	urlPath := "/byte/nonAscii"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getNonASCIIHandleResponse handles the GetNonASCII response.
func (client *ByteClient) getNonASCIIHandleResponse(resp *http.Response) (ByteClientGetNonASCIIResponse, error) {
	result := ByteClientGetNonASCIIResponse{}
	if err := runtime.UnmarshalAsByteArray(resp, &result.Value, runtime.Base64StdFormat); err != nil {
		return ByteClientGetNonASCIIResponse{}, err
	}
	return result, nil
}

// GetNull - Get null byte value
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - ByteClientGetNullOptions contains the optional parameters for the ByteClient.GetNull method.
func (client *ByteClient) GetNull(ctx context.Context, options *ByteClientGetNullOptions) (ByteClientGetNullResponse, error) {
	var err error
	const operationName = "ByteClient.GetNull"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getNullCreateRequest(ctx, options)
	if err != nil {
		return ByteClientGetNullResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ByteClientGetNullResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ByteClientGetNullResponse{}, err
	}
	resp, err := client.getNullHandleResponse(httpResp)
	return resp, err
}

// getNullCreateRequest creates the GetNull request.
func (client *ByteClient) getNullCreateRequest(ctx context.Context, options *ByteClientGetNullOptions) (*policy.Request, error) {
	urlPath := "/byte/null"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getNullHandleResponse handles the GetNull response.
func (client *ByteClient) getNullHandleResponse(resp *http.Response) (ByteClientGetNullResponse, error) {
	result := ByteClientGetNullResponse{}
	if err := runtime.UnmarshalAsByteArray(resp, &result.Value, runtime.Base64StdFormat); err != nil {
		return ByteClientGetNullResponse{}, err
	}
	return result, nil
}

// PutNonASCII - Put non-ascii byte string hex(FF FE FD FC FB FA F9 F8 F7 F6)
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - byteBody - Base64-encoded non-ascii byte string hex(FF FE FD FC FB FA F9 F8 F7 F6)
//   - options - ByteClientPutNonASCIIOptions contains the optional parameters for the ByteClient.PutNonASCII method.
func (client *ByteClient) PutNonASCII(ctx context.Context, byteBody []byte, options *ByteClientPutNonASCIIOptions) (ByteClientPutNonASCIIResponse, error) {
	var err error
	const operationName = "ByteClient.PutNonASCII"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putNonASCIICreateRequest(ctx, byteBody, options)
	if err != nil {
		return ByteClientPutNonASCIIResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ByteClientPutNonASCIIResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ByteClientPutNonASCIIResponse{}, err
	}
	return ByteClientPutNonASCIIResponse{}, nil
}

// putNonASCIICreateRequest creates the PutNonASCII request.
func (client *ByteClient) putNonASCIICreateRequest(ctx context.Context, byteBody []byte, options *ByteClientPutNonASCIIOptions) (*policy.Request, error) {
	urlPath := "/byte/nonAscii"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsByteArray(req, byteBody, runtime.Base64StdFormat); err != nil {
		return nil, err
	}
	return req, nil
}
