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

// ExtendsModelArrayClient contains the methods for the Type.Property.AdditionalProperties group.
// Don't use this type directly, use a constructor function instead.
type ExtendsModelArrayClient struct {
	internal *azcore.Client
}

// Get - Get call
func (client *ExtendsModelArrayClient) Get(ctx context.Context, options *ExtendsModelArrayClientGetOptions) (ExtendsModelArrayClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return ExtendsModelArrayClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExtendsModelArrayClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ExtendsModelArrayClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ExtendsModelArrayClient) getCreateRequest(ctx context.Context, options *ExtendsModelArrayClientGetOptions) (*policy.Request, error) {
	urlPath := "/type/property/additionalProperties/extendsRecordModelArray"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExtendsModelArrayClient) getHandleResponse(resp *http.Response) (ExtendsModelArrayClientGetResponse, error) {
	result := ExtendsModelArrayClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtendsModelArrayAdditionalProperties); err != nil {
		return ExtendsModelArrayClientGetResponse{}, err
	}
	return result, nil
}

// Put - Put operation
//   - body - body
func (client *ExtendsModelArrayClient) Put(ctx context.Context, body ExtendsModelArrayAdditionalProperties, options *ExtendsModelArrayClientPutOptions) (ExtendsModelArrayClientPutResponse, error) {
	var err error
	req, err := client.putCreateRequest(ctx, body, options)
	if err != nil {
		return ExtendsModelArrayClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExtendsModelArrayClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ExtendsModelArrayClientPutResponse{}, err
	}
	return ExtendsModelArrayClientPutResponse{}, nil
}

// putCreateRequest creates the Put request.
func (client *ExtendsModelArrayClient) putCreateRequest(ctx context.Context, body ExtendsModelArrayAdditionalProperties, options *ExtendsModelArrayClientPutOptions) (*policy.Request, error) {
	urlPath := "/type/property/additionalProperties/extendsRecordModelArray"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}
