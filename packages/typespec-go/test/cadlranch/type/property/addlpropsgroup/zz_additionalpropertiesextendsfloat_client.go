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

// AdditionalPropertiesExtendsFloatClient contains the methods for the Type.Property.AdditionalProperties namespace.
// Don't use this type directly, use [AdditionalPropertiesClient.NewAdditionalPropertiesExtendsFloatClient] instead.
type AdditionalPropertiesExtendsFloatClient struct {
	internal *azcore.Client
}

// Get - Get call
//   - options - AdditionalPropertiesExtendsFloatClientGetOptions contains the optional parameters for the AdditionalPropertiesExtendsFloatClient.Get
//     method.
func (client *AdditionalPropertiesExtendsFloatClient) Get(ctx context.Context, options *AdditionalPropertiesExtendsFloatClientGetOptions) (AdditionalPropertiesExtendsFloatClientGetResponse, error) {
	var err error
	const operationName = "AdditionalPropertiesExtendsFloatClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return AdditionalPropertiesExtendsFloatClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AdditionalPropertiesExtendsFloatClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AdditionalPropertiesExtendsFloatClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AdditionalPropertiesExtendsFloatClient) getCreateRequest(ctx context.Context, options *AdditionalPropertiesExtendsFloatClientGetOptions) (*policy.Request, error) {
	urlPath := "/type/property/additionalProperties/extendsRecordFloat"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AdditionalPropertiesExtendsFloatClient) getHandleResponse(resp *http.Response) (AdditionalPropertiesExtendsFloatClientGetResponse, error) {
	result := AdditionalPropertiesExtendsFloatClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtendsFloatAdditionalProperties); err != nil {
		return AdditionalPropertiesExtendsFloatClientGetResponse{}, err
	}
	return result, nil
}

// Put - Put operation
//   - body - body
//   - options - AdditionalPropertiesExtendsFloatClientPutOptions contains the optional parameters for the AdditionalPropertiesExtendsFloatClient.Put
//     method.
func (client *AdditionalPropertiesExtendsFloatClient) Put(ctx context.Context, body ExtendsFloatAdditionalProperties, options *AdditionalPropertiesExtendsFloatClientPutOptions) (AdditionalPropertiesExtendsFloatClientPutResponse, error) {
	var err error
	const operationName = "AdditionalPropertiesExtendsFloatClient.Put"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putCreateRequest(ctx, body, options)
	if err != nil {
		return AdditionalPropertiesExtendsFloatClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AdditionalPropertiesExtendsFloatClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return AdditionalPropertiesExtendsFloatClientPutResponse{}, err
	}
	return AdditionalPropertiesExtendsFloatClientPutResponse{}, nil
}

// putCreateRequest creates the Put request.
func (client *AdditionalPropertiesExtendsFloatClient) putCreateRequest(ctx context.Context, body ExtendsFloatAdditionalProperties, options *AdditionalPropertiesExtendsFloatClientPutOptions) (*policy.Request, error) {
	urlPath := "/type/property/additionalProperties/extendsRecordFloat"
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