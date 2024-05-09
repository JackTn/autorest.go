// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package headgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// HTTPSuccessClient contains the methods for the HTTPSuccess group.
// Don't use this type directly, use a constructor function instead.
type HTTPSuccessClient struct {
	internal *azcore.Client
}

// Head200 - Return 200 status code if successful
//
// Generated from API version 1.0.0
//   - options - HTTPSuccessClientHead200Options contains the optional parameters for the HTTPSuccessClient.Head200 method.
func (client *HTTPSuccessClient) Head200(ctx context.Context, options *HTTPSuccessClientHead200Options) (HTTPSuccessClientHead200Response, error) {
	var err error
	const operationName = "HTTPSuccessClient.Head200"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.head200CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientHead200Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPSuccessClientHead200Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNotFound) {
		err = runtime.NewResponseError(httpResp)
		return HTTPSuccessClientHead200Response{}, err
	}
	return HTTPSuccessClientHead200Response{Success: httpResp.StatusCode >= 200 && httpResp.StatusCode < 300}, nil
}

// head200CreateRequest creates the Head200 request.
func (client *HTTPSuccessClient) head200CreateRequest(ctx context.Context, _ *HTTPSuccessClientHead200Options) (*policy.Request, error) {
	urlPath := "/http/success/200"
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// Head204 - Return 204 status code if successful
//
// Generated from API version 1.0.0
//   - options - HTTPSuccessClientHead204Options contains the optional parameters for the HTTPSuccessClient.Head204 method.
func (client *HTTPSuccessClient) Head204(ctx context.Context, options *HTTPSuccessClientHead204Options) (HTTPSuccessClientHead204Response, error) {
	var err error
	const operationName = "HTTPSuccessClient.Head204"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.head204CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientHead204Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPSuccessClientHead204Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent, http.StatusNotFound) {
		err = runtime.NewResponseError(httpResp)
		return HTTPSuccessClientHead204Response{}, err
	}
	return HTTPSuccessClientHead204Response{Success: httpResp.StatusCode >= 200 && httpResp.StatusCode < 300}, nil
}

// head204CreateRequest creates the Head204 request.
func (client *HTTPSuccessClient) head204CreateRequest(ctx context.Context, _ *HTTPSuccessClientHead204Options) (*policy.Request, error) {
	urlPath := "/http/success/204"
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// Head404 - Return 404 status code if successful
//
// Generated from API version 1.0.0
//   - options - HTTPSuccessClientHead404Options contains the optional parameters for the HTTPSuccessClient.Head404 method.
func (client *HTTPSuccessClient) Head404(ctx context.Context, options *HTTPSuccessClientHead404Options) (HTTPSuccessClientHead404Response, error) {
	var err error
	const operationName = "HTTPSuccessClient.Head404"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.head404CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientHead404Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPSuccessClientHead404Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent, http.StatusNotFound) {
		err = runtime.NewResponseError(httpResp)
		return HTTPSuccessClientHead404Response{}, err
	}
	return HTTPSuccessClientHead404Response{Success: httpResp.StatusCode >= 200 && httpResp.StatusCode < 300}, nil
}

// head404CreateRequest creates the Head404 request.
func (client *HTTPSuccessClient) head404CreateRequest(ctx context.Context, _ *HTTPSuccessClientHead404Options) (*policy.Request, error) {
	urlPath := "/http/success/404"
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}
