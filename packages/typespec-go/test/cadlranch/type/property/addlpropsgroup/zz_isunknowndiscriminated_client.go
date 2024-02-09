//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package addlpropsgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// IsUnknownDiscriminatedClient contains the methods for the Type.Property.AdditionalProperties group.
// Don't use this type directly, use a constructor function instead.
type IsUnknownDiscriminatedClient struct {
	internal *azcore.Client
}

// Get - Get call
//   - options - IsUnknownDiscriminatedClientGetOptions contains the optional parameters for the IsUnknownDiscriminatedClient.Get
//     method.
func (client *IsUnknownDiscriminatedClient) Get(ctx context.Context, options *IsUnknownDiscriminatedClientGetOptions) (IsUnknownDiscriminatedClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return IsUnknownDiscriminatedClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IsUnknownDiscriminatedClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IsUnknownDiscriminatedClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *IsUnknownDiscriminatedClient) getCreateRequest(ctx context.Context, options *IsUnknownDiscriminatedClientGetOptions) (*policy.Request, error) {
	urlPath := "/type/property/additionalProperties/isUnknownDiscriminated"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IsUnknownDiscriminatedClient) getHandleResponse(resp *http.Response) (IsUnknownDiscriminatedClientGetResponse, error) {
	result := IsUnknownDiscriminatedClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return IsUnknownDiscriminatedClientGetResponse{}, err
	}
	return result, nil
}

// Put - Put operation
//   - body - body
//   - options - IsUnknownDiscriminatedClientPutOptions contains the optional parameters for the IsUnknownDiscriminatedClient.Put
//     method.
func (client *IsUnknownDiscriminatedClient) Put(ctx context.Context, body IsUnknownAdditionalPropertiesDiscriminatedClassification, options *IsUnknownDiscriminatedClientPutOptions) (IsUnknownDiscriminatedClientPutResponse, error) {
	var err error
	req, err := client.putCreateRequest(ctx, body, options)
	if err != nil {
		return IsUnknownDiscriminatedClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IsUnknownDiscriminatedClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return IsUnknownDiscriminatedClientPutResponse{}, err
	}
	return IsUnknownDiscriminatedClientPutResponse{}, nil
}

// putCreateRequest creates the Put request.
func (client *IsUnknownDiscriminatedClient) putCreateRequest(ctx context.Context, body IsUnknownAdditionalPropertiesDiscriminatedClassification, options *IsUnknownDiscriminatedClientPutOptions) (*policy.Request, error) {
	urlPath := "/type/property/additionalProperties/isUnknownDiscriminated"
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
