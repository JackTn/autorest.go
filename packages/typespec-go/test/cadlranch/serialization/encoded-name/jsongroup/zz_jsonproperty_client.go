// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package jsongroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// JSONPropertyClient contains the methods for the Serialization.EncodedName.Json namespace.
// Don't use this type directly, use [JSONClient.NewJSONPropertyClient] instead.
type JSONPropertyClient struct {
	internal *azcore.Client
}

// - options - JSONPropertyClientGetOptions contains the optional parameters for the JSONPropertyClient.Get method.
func (client *JSONPropertyClient) Get(ctx context.Context, options *JSONPropertyClientGetOptions) (JSONPropertyClientGetResponse, error) {
	var err error
	const operationName = "JSONPropertyClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return JSONPropertyClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return JSONPropertyClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return JSONPropertyClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *JSONPropertyClient) getCreateRequest(ctx context.Context, _ *JSONPropertyClientGetOptions) (*policy.Request, error) {
	urlPath := "/serialization/encoded-name/json/property"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *JSONPropertyClient) getHandleResponse(resp *http.Response) (JSONPropertyClientGetResponse, error) {
	result := JSONPropertyClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JSONEncodedNameModel); err != nil {
		return JSONPropertyClientGetResponse{}, err
	}
	return result, nil
}

// - options - JSONPropertyClientSendOptions contains the optional parameters for the JSONPropertyClient.Send method.
func (client *JSONPropertyClient) Send(ctx context.Context, jsonEncodedNameModel JSONEncodedNameModel, options *JSONPropertyClientSendOptions) (JSONPropertyClientSendResponse, error) {
	var err error
	const operationName = "JSONPropertyClient.Send"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.sendCreateRequest(ctx, jsonEncodedNameModel, options)
	if err != nil {
		return JSONPropertyClientSendResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return JSONPropertyClientSendResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return JSONPropertyClientSendResponse{}, err
	}
	return JSONPropertyClientSendResponse{}, nil
}

// sendCreateRequest creates the Send request.
func (client *JSONPropertyClient) sendCreateRequest(ctx context.Context, jsonEncodedNameModel JSONEncodedNameModel, _ *JSONPropertyClientSendOptions) (*policy.Request, error) {
	urlPath := "/serialization/encoded-name/json/property"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, jsonEncodedNameModel); err != nil {
		return nil, err
	}
	return req, nil
}
