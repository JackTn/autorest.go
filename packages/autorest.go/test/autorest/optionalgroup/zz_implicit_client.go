// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package optionalgroup

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ImplicitClient contains the methods for the Implicit group.
// Don't use this type directly, use a constructor function instead.
type ImplicitClient struct {
	internal            *azcore.Client
	requiredGlobalPath  string
	requiredGlobalQuery string
	optionalGlobalQuery *int32
}

// GetOptionalGlobalQuery - Test implicitly optional query parameter
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - ImplicitClientGetOptionalGlobalQueryOptions contains the optional parameters for the ImplicitClient.GetOptionalGlobalQuery
//     method.
func (client *ImplicitClient) GetOptionalGlobalQuery(ctx context.Context, options *ImplicitClientGetOptionalGlobalQueryOptions) (ImplicitClientGetOptionalGlobalQueryResponse, error) {
	var err error
	const operationName = "ImplicitClient.GetOptionalGlobalQuery"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getOptionalGlobalQueryCreateRequest(ctx, options)
	if err != nil {
		return ImplicitClientGetOptionalGlobalQueryResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ImplicitClientGetOptionalGlobalQueryResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ImplicitClientGetOptionalGlobalQueryResponse{}, err
	}
	return ImplicitClientGetOptionalGlobalQueryResponse{}, nil
}

// getOptionalGlobalQueryCreateRequest creates the GetOptionalGlobalQuery request.
func (client *ImplicitClient) getOptionalGlobalQueryCreateRequest(ctx context.Context, _ *ImplicitClientGetOptionalGlobalQueryOptions) (*policy.Request, error) {
	urlPath := "/reqopt/global/optional/query"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if client.optionalGlobalQuery != nil {
		reqQP.Set("optional-global-query", strconv.FormatInt(int64(*client.optionalGlobalQuery), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetRequiredGlobalPath - Test implicitly required path parameter
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - ImplicitClientGetRequiredGlobalPathOptions contains the optional parameters for the ImplicitClient.GetRequiredGlobalPath
//     method.
func (client *ImplicitClient) GetRequiredGlobalPath(ctx context.Context, options *ImplicitClientGetRequiredGlobalPathOptions) (ImplicitClientGetRequiredGlobalPathResponse, error) {
	var err error
	const operationName = "ImplicitClient.GetRequiredGlobalPath"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getRequiredGlobalPathCreateRequest(ctx, options)
	if err != nil {
		return ImplicitClientGetRequiredGlobalPathResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ImplicitClientGetRequiredGlobalPathResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ImplicitClientGetRequiredGlobalPathResponse{}, err
	}
	return ImplicitClientGetRequiredGlobalPathResponse{}, nil
}

// getRequiredGlobalPathCreateRequest creates the GetRequiredGlobalPath request.
func (client *ImplicitClient) getRequiredGlobalPathCreateRequest(ctx context.Context, _ *ImplicitClientGetRequiredGlobalPathOptions) (*policy.Request, error) {
	urlPath := "/reqopt/global/required/path/{required-global-path}"
	if client.requiredGlobalPath == "" {
		return nil, errors.New("parameter client.requiredGlobalPath cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{required-global-path}", url.PathEscape(client.requiredGlobalPath))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetRequiredGlobalQuery - Test implicitly required query parameter
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - ImplicitClientGetRequiredGlobalQueryOptions contains the optional parameters for the ImplicitClient.GetRequiredGlobalQuery
//     method.
func (client *ImplicitClient) GetRequiredGlobalQuery(ctx context.Context, options *ImplicitClientGetRequiredGlobalQueryOptions) (ImplicitClientGetRequiredGlobalQueryResponse, error) {
	var err error
	const operationName = "ImplicitClient.GetRequiredGlobalQuery"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getRequiredGlobalQueryCreateRequest(ctx, options)
	if err != nil {
		return ImplicitClientGetRequiredGlobalQueryResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ImplicitClientGetRequiredGlobalQueryResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ImplicitClientGetRequiredGlobalQueryResponse{}, err
	}
	return ImplicitClientGetRequiredGlobalQueryResponse{}, nil
}

// getRequiredGlobalQueryCreateRequest creates the GetRequiredGlobalQuery request.
func (client *ImplicitClient) getRequiredGlobalQueryCreateRequest(ctx context.Context, _ *ImplicitClientGetRequiredGlobalQueryOptions) (*policy.Request, error) {
	urlPath := "/reqopt/global/required/query"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("required-global-query", client.requiredGlobalQuery)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetRequiredPath - Test implicitly required path parameter
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - ImplicitClientGetRequiredPathOptions contains the optional parameters for the ImplicitClient.GetRequiredPath
//     method.
func (client *ImplicitClient) GetRequiredPath(ctx context.Context, pathParameter string, options *ImplicitClientGetRequiredPathOptions) (ImplicitClientGetRequiredPathResponse, error) {
	var err error
	const operationName = "ImplicitClient.GetRequiredPath"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getRequiredPathCreateRequest(ctx, pathParameter, options)
	if err != nil {
		return ImplicitClientGetRequiredPathResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ImplicitClientGetRequiredPathResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ImplicitClientGetRequiredPathResponse{}, err
	}
	return ImplicitClientGetRequiredPathResponse{}, nil
}

// getRequiredPathCreateRequest creates the GetRequiredPath request.
func (client *ImplicitClient) getRequiredPathCreateRequest(ctx context.Context, pathParameter string, _ *ImplicitClientGetRequiredPathOptions) (*policy.Request, error) {
	urlPath := "/reqopt/implicit/required/path/{pathParameter}"
	if pathParameter == "" {
		return nil, errors.New("parameter pathParameter cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pathParameter}", url.PathEscape(pathParameter))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// PutOptionalBinaryBody - Test implicitly optional body parameter
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - ImplicitClientPutOptionalBinaryBodyOptions contains the optional parameters for the ImplicitClient.PutOptionalBinaryBody
//     method.
func (client *ImplicitClient) PutOptionalBinaryBody(ctx context.Context, options *ImplicitClientPutOptionalBinaryBodyOptions) (ImplicitClientPutOptionalBinaryBodyResponse, error) {
	var err error
	const operationName = "ImplicitClient.PutOptionalBinaryBody"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putOptionalBinaryBodyCreateRequest(ctx, options)
	if err != nil {
		return ImplicitClientPutOptionalBinaryBodyResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ImplicitClientPutOptionalBinaryBodyResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ImplicitClientPutOptionalBinaryBodyResponse{}, err
	}
	return ImplicitClientPutOptionalBinaryBodyResponse{}, nil
}

// putOptionalBinaryBodyCreateRequest creates the PutOptionalBinaryBody request.
func (client *ImplicitClient) putOptionalBinaryBodyCreateRequest(ctx context.Context, options *ImplicitClientPutOptionalBinaryBodyOptions) (*policy.Request, error) {
	urlPath := "/reqopt/implicit/optional/binary-body"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.BodyParameter != nil {
		if err := req.SetBody(options.BodyParameter, "application/octet-stream"); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// PutOptionalBody - Test implicitly optional body parameter
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - ImplicitClientPutOptionalBodyOptions contains the optional parameters for the ImplicitClient.PutOptionalBody
//     method.
func (client *ImplicitClient) PutOptionalBody(ctx context.Context, options *ImplicitClientPutOptionalBodyOptions) (ImplicitClientPutOptionalBodyResponse, error) {
	var err error
	const operationName = "ImplicitClient.PutOptionalBody"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putOptionalBodyCreateRequest(ctx, options)
	if err != nil {
		return ImplicitClientPutOptionalBodyResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ImplicitClientPutOptionalBodyResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ImplicitClientPutOptionalBodyResponse{}, err
	}
	return ImplicitClientPutOptionalBodyResponse{}, nil
}

// putOptionalBodyCreateRequest creates the PutOptionalBody request.
func (client *ImplicitClient) putOptionalBodyCreateRequest(ctx context.Context, options *ImplicitClientPutOptionalBodyOptions) (*policy.Request, error) {
	urlPath := "/reqopt/implicit/optional/body"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.BodyParameter != nil {
		if err := runtime.MarshalAsJSON(req, *options.BodyParameter); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// PutOptionalHeader - Test implicitly optional header parameter
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - ImplicitClientPutOptionalHeaderOptions contains the optional parameters for the ImplicitClient.PutOptionalHeader
//     method.
func (client *ImplicitClient) PutOptionalHeader(ctx context.Context, options *ImplicitClientPutOptionalHeaderOptions) (ImplicitClientPutOptionalHeaderResponse, error) {
	var err error
	const operationName = "ImplicitClient.PutOptionalHeader"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putOptionalHeaderCreateRequest(ctx, options)
	if err != nil {
		return ImplicitClientPutOptionalHeaderResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ImplicitClientPutOptionalHeaderResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ImplicitClientPutOptionalHeaderResponse{}, err
	}
	return ImplicitClientPutOptionalHeaderResponse{}, nil
}

// putOptionalHeaderCreateRequest creates the PutOptionalHeader request.
func (client *ImplicitClient) putOptionalHeaderCreateRequest(ctx context.Context, options *ImplicitClientPutOptionalHeaderOptions) (*policy.Request, error) {
	urlPath := "/reqopt/implicit/optional/header"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.QueryParameter != nil {
		req.Raw().Header["queryParameter"] = []string{*options.QueryParameter}
	}
	return req, nil
}

// PutOptionalQuery - Test implicitly optional query parameter
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - ImplicitClientPutOptionalQueryOptions contains the optional parameters for the ImplicitClient.PutOptionalQuery
//     method.
func (client *ImplicitClient) PutOptionalQuery(ctx context.Context, options *ImplicitClientPutOptionalQueryOptions) (ImplicitClientPutOptionalQueryResponse, error) {
	var err error
	const operationName = "ImplicitClient.PutOptionalQuery"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putOptionalQueryCreateRequest(ctx, options)
	if err != nil {
		return ImplicitClientPutOptionalQueryResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ImplicitClientPutOptionalQueryResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ImplicitClientPutOptionalQueryResponse{}, err
	}
	return ImplicitClientPutOptionalQueryResponse{}, nil
}

// putOptionalQueryCreateRequest creates the PutOptionalQuery request.
func (client *ImplicitClient) putOptionalQueryCreateRequest(ctx context.Context, options *ImplicitClientPutOptionalQueryOptions) (*policy.Request, error) {
	urlPath := "/reqopt/implicit/optional/query"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.QueryParameter != nil {
		reqQP.Set("queryParameter", *options.QueryParameter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
