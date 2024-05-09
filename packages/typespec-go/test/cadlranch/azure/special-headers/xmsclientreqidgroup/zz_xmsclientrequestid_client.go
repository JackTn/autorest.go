// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package xmsclientreqidgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// XMSClientRequestIDClient - Azure client request id header configurations.
// Don't use this type directly, use a constructor function instead.
type XMSClientRequestIDClient struct {
	internal *azcore.Client
}

// Get - Get operation with azure `x-ms-client-request-id` header.
//   - options - XMSClientRequestIDClientGetOptions contains the optional parameters for the XMSClientRequestIDClient.Get method.
func (client *XMSClientRequestIDClient) Get(ctx context.Context, options *XMSClientRequestIDClientGetOptions) (XMSClientRequestIDClientGetResponse, error) {
	var err error
	const operationName = "XMSClientRequestIDClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return XMSClientRequestIDClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return XMSClientRequestIDClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return XMSClientRequestIDClientGetResponse{}, err
	}
	return XMSClientRequestIDClientGetResponse{}, nil
}

// getCreateRequest creates the Get request.
func (client *XMSClientRequestIDClient) getCreateRequest(ctx context.Context, options *XMSClientRequestIDClientGetOptions) (*policy.Request, error) {
	urlPath := "/azure/special-headers/x-ms-client-request-id"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	if options != nil && options.ClientRequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.ClientRequestID}
	}
	return req, nil
}