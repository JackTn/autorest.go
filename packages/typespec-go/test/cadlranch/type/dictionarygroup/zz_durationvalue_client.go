//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package dictionarygroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// DurationValueClient contains the methods for the Type.Dictionary group.
// Don't use this type directly, use a constructor function instead.
type DurationValueClient struct {
	internal *azcore.Client
}

func (client *DurationValueClient) Get(ctx context.Context, options *DurationValueClientGetOptions) (DurationValueClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return DurationValueClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DurationValueClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DurationValueClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DurationValueClient) getCreateRequest(ctx context.Context, options *DurationValueClientGetOptions) (*policy.Request, error) {
	urlPath := "/type/dictionary/duration"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DurationValueClient) getHandleResponse(resp *http.Response) (DurationValueClientGetResponse, error) {
	result := DurationValueClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return DurationValueClientGetResponse{}, err
	}
	return result, nil
}

func (client *DurationValueClient) Put(ctx context.Context, body map[string]*string, options *DurationValueClientPutOptions) (DurationValueClientPutResponse, error) {
	var err error
	req, err := client.putCreateRequest(ctx, body, options)
	if err != nil {
		return DurationValueClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DurationValueClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return DurationValueClientPutResponse{}, err
	}
	return DurationValueClientPutResponse{}, nil
}

// putCreateRequest creates the Put request.
func (client *DurationValueClient) putCreateRequest(ctx context.Context, body map[string]*string, options *DurationValueClientPutOptions) (*policy.Request, error) {
	urlPath := "/type/dictionary/duration"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}