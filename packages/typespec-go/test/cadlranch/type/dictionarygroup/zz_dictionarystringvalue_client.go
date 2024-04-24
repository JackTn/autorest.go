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

// DictionaryStringValueClient - Dictionary of string values
// Don't use this type directly, use [DictionaryClient.NewDictionaryStringValueClient] instead.
type DictionaryStringValueClient struct {
	internal *azcore.Client
}

//   - options - DictionaryStringValueClientGetOptions contains the optional parameters for the DictionaryStringValueClient.Get
//     method.
func (client *DictionaryStringValueClient) Get(ctx context.Context, options *DictionaryStringValueClientGetOptions) (DictionaryStringValueClientGetResponse, error) {
	var err error
	const operationName = "DictionaryStringValueClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return DictionaryStringValueClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DictionaryStringValueClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DictionaryStringValueClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DictionaryStringValueClient) getCreateRequest(ctx context.Context, options *DictionaryStringValueClientGetOptions) (*policy.Request, error) {
	urlPath := "/type/dictionary/string"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DictionaryStringValueClient) getHandleResponse(resp *http.Response) (DictionaryStringValueClientGetResponse, error) {
	result := DictionaryStringValueClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return DictionaryStringValueClientGetResponse{}, err
	}
	return result, nil
}

//   - options - DictionaryStringValueClientPutOptions contains the optional parameters for the DictionaryStringValueClient.Put
//     method.
func (client *DictionaryStringValueClient) Put(ctx context.Context, body map[string]*string, options *DictionaryStringValueClientPutOptions) (DictionaryStringValueClientPutResponse, error) {
	var err error
	const operationName = "DictionaryStringValueClient.Put"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putCreateRequest(ctx, body, options)
	if err != nil {
		return DictionaryStringValueClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DictionaryStringValueClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return DictionaryStringValueClientPutResponse{}, err
	}
	return DictionaryStringValueClientPutResponse{}, nil
}

// putCreateRequest creates the Put request.
func (client *DictionaryStringValueClient) putCreateRequest(ctx context.Context, body map[string]*string, options *DictionaryStringValueClientPutOptions) (*policy.Request, error) {
	urlPath := "/type/dictionary/string"
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