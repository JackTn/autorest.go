// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package accessgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// AccessPublicOperationClient contains the methods for the _Specs_.Azure.ClientGenerator.Core.Access namespace.
// Don't use this type directly, use [AccessClient.NewAccessPublicOperationClient] instead.
type AccessPublicOperationClient struct {
	internal *azcore.Client
}

//   - options - AccessPublicOperationClientNoDecoratorInPublicOptions contains the optional parameters for the AccessPublicOperationClient.NoDecoratorInPublic
//     method.
func (client *AccessPublicOperationClient) NoDecoratorInPublic(ctx context.Context, name string, options *AccessPublicOperationClientNoDecoratorInPublicOptions) (AccessPublicOperationClientNoDecoratorInPublicResponse, error) {
	var err error
	const operationName = "AccessPublicOperationClient.NoDecoratorInPublic"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.noDecoratorInPublicCreateRequest(ctx, name, options)
	if err != nil {
		return AccessPublicOperationClientNoDecoratorInPublicResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AccessPublicOperationClientNoDecoratorInPublicResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AccessPublicOperationClientNoDecoratorInPublicResponse{}, err
	}
	resp, err := client.noDecoratorInPublicHandleResponse(httpResp)
	return resp, err
}

// noDecoratorInPublicCreateRequest creates the NoDecoratorInPublic request.
func (client *AccessPublicOperationClient) noDecoratorInPublicCreateRequest(ctx context.Context, name string, _ *AccessPublicOperationClientNoDecoratorInPublicOptions) (*policy.Request, error) {
	urlPath := "/azure/client-generator-core/access/publicOperation/noDecoratorInPublic"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("name", name)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// noDecoratorInPublicHandleResponse handles the NoDecoratorInPublic response.
func (client *AccessPublicOperationClient) noDecoratorInPublicHandleResponse(resp *http.Response) (AccessPublicOperationClientNoDecoratorInPublicResponse, error) {
	result := AccessPublicOperationClientNoDecoratorInPublicResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NoDecoratorModelInPublic); err != nil {
		return AccessPublicOperationClientNoDecoratorInPublicResponse{}, err
	}
	return result, nil
}

//   - options - AccessPublicOperationClientPublicDecoratorInPublicOptions contains the optional parameters for the AccessPublicOperationClient.PublicDecoratorInPublic
//     method.
func (client *AccessPublicOperationClient) PublicDecoratorInPublic(ctx context.Context, name string, options *AccessPublicOperationClientPublicDecoratorInPublicOptions) (AccessPublicOperationClientPublicDecoratorInPublicResponse, error) {
	var err error
	const operationName = "AccessPublicOperationClient.PublicDecoratorInPublic"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.publicDecoratorInPublicCreateRequest(ctx, name, options)
	if err != nil {
		return AccessPublicOperationClientPublicDecoratorInPublicResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AccessPublicOperationClientPublicDecoratorInPublicResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AccessPublicOperationClientPublicDecoratorInPublicResponse{}, err
	}
	resp, err := client.publicDecoratorInPublicHandleResponse(httpResp)
	return resp, err
}

// publicDecoratorInPublicCreateRequest creates the PublicDecoratorInPublic request.
func (client *AccessPublicOperationClient) publicDecoratorInPublicCreateRequest(ctx context.Context, name string, _ *AccessPublicOperationClientPublicDecoratorInPublicOptions) (*policy.Request, error) {
	urlPath := "/azure/client-generator-core/access/publicOperation/publicDecoratorInPublic"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("name", name)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// publicDecoratorInPublicHandleResponse handles the PublicDecoratorInPublic response.
func (client *AccessPublicOperationClient) publicDecoratorInPublicHandleResponse(resp *http.Response) (AccessPublicOperationClientPublicDecoratorInPublicResponse, error) {
	result := AccessPublicOperationClientPublicDecoratorInPublicResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PublicDecoratorModelInPublic); err != nil {
		return AccessPublicOperationClientPublicDecoratorInPublicResponse{}, err
	}
	return result, nil
}
