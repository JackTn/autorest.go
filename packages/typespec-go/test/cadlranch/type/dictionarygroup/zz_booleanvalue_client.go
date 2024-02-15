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

// BooleanValueClient - Dictionary of boolean values
// Don't use this type directly, use [DictionaryClient.NewBooleanValueClient] instead.
type BooleanValueClient struct {
	internal *azcore.Client
}

// - options - BooleanValueClientGetOptions contains the optional parameters for the BooleanValueClient.Get method.
func (client *BooleanValueClient) Get(ctx context.Context, options *BooleanValueClientGetOptions) (BooleanValueClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return BooleanValueClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BooleanValueClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BooleanValueClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *BooleanValueClient) getCreateRequest(ctx context.Context, options *BooleanValueClientGetOptions) (*policy.Request, error) {
	urlPath := "/type/dictionary/boolean"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BooleanValueClient) getHandleResponse(resp *http.Response) (BooleanValueClientGetResponse, error) {
	result := BooleanValueClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return BooleanValueClientGetResponse{}, err
	}
	return result, nil
}

// - options - BooleanValueClientPutOptions contains the optional parameters for the BooleanValueClient.Put method.
func (client *BooleanValueClient) Put(ctx context.Context, body map[string]*bool, options *BooleanValueClientPutOptions) (BooleanValueClientPutResponse, error) {
	var err error
	req, err := client.putCreateRequest(ctx, body, options)
	if err != nil {
		return BooleanValueClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BooleanValueClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return BooleanValueClientPutResponse{}, err
	}
	return BooleanValueClientPutResponse{}, nil
}

// putCreateRequest creates the Put request.
func (client *BooleanValueClient) putCreateRequest(ctx context.Context, body map[string]*bool, options *BooleanValueClientPutOptions) (*policy.Request, error) {
	urlPath := "/type/dictionary/boolean"
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
