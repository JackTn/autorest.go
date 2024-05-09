// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package nullablegroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// NullableBytesClient contains the methods for the Type.Property.Nullable namespace.
// Don't use this type directly, use [NullableClient.NewNullableBytesClient] instead.
type NullableBytesClient struct {
	internal *azcore.Client
}

// GetNonNull - Get models that will return all properties in the model
//   - options - NullableBytesClientGetNonNullOptions contains the optional parameters for the NullableBytesClient.GetNonNull
//     method.
func (client *NullableBytesClient) GetNonNull(ctx context.Context, options *NullableBytesClientGetNonNullOptions) (NullableBytesClientGetNonNullResponse, error) {
	var err error
	const operationName = "NullableBytesClient.GetNonNull"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getNonNullCreateRequest(ctx, options)
	if err != nil {
		return NullableBytesClientGetNonNullResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NullableBytesClientGetNonNullResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NullableBytesClientGetNonNullResponse{}, err
	}
	resp, err := client.getNonNullHandleResponse(httpResp)
	return resp, err
}

// getNonNullCreateRequest creates the GetNonNull request.
func (client *NullableBytesClient) getNonNullCreateRequest(ctx context.Context, _ *NullableBytesClientGetNonNullOptions) (*policy.Request, error) {
	urlPath := "/type/property/nullable/bytes/non-null"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getNonNullHandleResponse handles the GetNonNull response.
func (client *NullableBytesClient) getNonNullHandleResponse(resp *http.Response) (NullableBytesClientGetNonNullResponse, error) {
	result := NullableBytesClientGetNonNullResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BytesProperty); err != nil {
		return NullableBytesClientGetNonNullResponse{}, err
	}
	return result, nil
}

// GetNull - Get models that will return the default object
//   - options - NullableBytesClientGetNullOptions contains the optional parameters for the NullableBytesClient.GetNull method.
func (client *NullableBytesClient) GetNull(ctx context.Context, options *NullableBytesClientGetNullOptions) (NullableBytesClientGetNullResponse, error) {
	var err error
	const operationName = "NullableBytesClient.GetNull"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getNullCreateRequest(ctx, options)
	if err != nil {
		return NullableBytesClientGetNullResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NullableBytesClientGetNullResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NullableBytesClientGetNullResponse{}, err
	}
	resp, err := client.getNullHandleResponse(httpResp)
	return resp, err
}

// getNullCreateRequest creates the GetNull request.
func (client *NullableBytesClient) getNullCreateRequest(ctx context.Context, _ *NullableBytesClientGetNullOptions) (*policy.Request, error) {
	urlPath := "/type/property/nullable/bytes/null"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getNullHandleResponse handles the GetNull response.
func (client *NullableBytesClient) getNullHandleResponse(resp *http.Response) (NullableBytesClientGetNullResponse, error) {
	result := NullableBytesClientGetNullResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BytesProperty); err != nil {
		return NullableBytesClientGetNullResponse{}, err
	}
	return result, nil
}

// PatchNonNull - Put a body with all properties present.
//   - options - NullableBytesClientPatchNonNullOptions contains the optional parameters for the NullableBytesClient.PatchNonNull
//     method.
func (client *NullableBytesClient) PatchNonNull(ctx context.Context, body BytesProperty, options *NullableBytesClientPatchNonNullOptions) (NullableBytesClientPatchNonNullResponse, error) {
	var err error
	const operationName = "NullableBytesClient.PatchNonNull"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.patchNonNullCreateRequest(ctx, body, options)
	if err != nil {
		return NullableBytesClientPatchNonNullResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NullableBytesClientPatchNonNullResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return NullableBytesClientPatchNonNullResponse{}, err
	}
	return NullableBytesClientPatchNonNullResponse{}, nil
}

// patchNonNullCreateRequest creates the PatchNonNull request.
func (client *NullableBytesClient) patchNonNullCreateRequest(ctx context.Context, body BytesProperty, _ *NullableBytesClientPatchNonNullOptions) (*policy.Request, error) {
	urlPath := "/type/property/nullable/bytes/non-null"
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/merge-patch+json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// PatchNull - Put a body with default properties.
//   - options - NullableBytesClientPatchNullOptions contains the optional parameters for the NullableBytesClient.PatchNull method.
func (client *NullableBytesClient) PatchNull(ctx context.Context, body BytesProperty, options *NullableBytesClientPatchNullOptions) (NullableBytesClientPatchNullResponse, error) {
	var err error
	const operationName = "NullableBytesClient.PatchNull"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.patchNullCreateRequest(ctx, body, options)
	if err != nil {
		return NullableBytesClientPatchNullResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NullableBytesClientPatchNullResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return NullableBytesClientPatchNullResponse{}, err
	}
	return NullableBytesClientPatchNullResponse{}, nil
}

// patchNullCreateRequest creates the PatchNull request.
func (client *NullableBytesClient) patchNullCreateRequest(ctx context.Context, body BytesProperty, _ *NullableBytesClientPatchNullOptions) (*policy.Request, error) {
	urlPath := "/type/property/nullable/bytes/null"
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/merge-patch+json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}
