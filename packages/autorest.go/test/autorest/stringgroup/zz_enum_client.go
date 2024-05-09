// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package stringgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// EnumClient contains the methods for the Enum group.
// Don't use this type directly, use a constructor function instead.
type EnumClient struct {
	internal *azcore.Client
}

// GetNotExpandable - Get enum value 'red color' from enumeration of 'red color', 'green-color', 'blue_color'.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - EnumClientGetNotExpandableOptions contains the optional parameters for the EnumClient.GetNotExpandable method.
func (client *EnumClient) GetNotExpandable(ctx context.Context, options *EnumClientGetNotExpandableOptions) (EnumClientGetNotExpandableResponse, error) {
	var err error
	const operationName = "EnumClient.GetNotExpandable"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getNotExpandableCreateRequest(ctx, options)
	if err != nil {
		return EnumClientGetNotExpandableResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EnumClientGetNotExpandableResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EnumClientGetNotExpandableResponse{}, err
	}
	resp, err := client.getNotExpandableHandleResponse(httpResp)
	return resp, err
}

// getNotExpandableCreateRequest creates the GetNotExpandable request.
func (client *EnumClient) getNotExpandableCreateRequest(ctx context.Context, _ *EnumClientGetNotExpandableOptions) (*policy.Request, error) {
	urlPath := "/string/enum/notExpandable"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getNotExpandableHandleResponse handles the GetNotExpandable response.
func (client *EnumClient) getNotExpandableHandleResponse(resp *http.Response) (EnumClientGetNotExpandableResponse, error) {
	result := EnumClientGetNotExpandableResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return EnumClientGetNotExpandableResponse{}, err
	}
	return result, nil
}

// GetReferenced - Get enum value 'red color' from enumeration of 'red color', 'green-color', 'blue_color'.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - EnumClientGetReferencedOptions contains the optional parameters for the EnumClient.GetReferenced method.
func (client *EnumClient) GetReferenced(ctx context.Context, options *EnumClientGetReferencedOptions) (EnumClientGetReferencedResponse, error) {
	var err error
	const operationName = "EnumClient.GetReferenced"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getReferencedCreateRequest(ctx, options)
	if err != nil {
		return EnumClientGetReferencedResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EnumClientGetReferencedResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EnumClientGetReferencedResponse{}, err
	}
	resp, err := client.getReferencedHandleResponse(httpResp)
	return resp, err
}

// getReferencedCreateRequest creates the GetReferenced request.
func (client *EnumClient) getReferencedCreateRequest(ctx context.Context, _ *EnumClientGetReferencedOptions) (*policy.Request, error) {
	urlPath := "/string/enum/Referenced"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getReferencedHandleResponse handles the GetReferenced response.
func (client *EnumClient) getReferencedHandleResponse(resp *http.Response) (EnumClientGetReferencedResponse, error) {
	result := EnumClientGetReferencedResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return EnumClientGetReferencedResponse{}, err
	}
	return result, nil
}

// GetReferencedConstant - Get value 'green-color' from the constant.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - EnumClientGetReferencedConstantOptions contains the optional parameters for the EnumClient.GetReferencedConstant
//     method.
func (client *EnumClient) GetReferencedConstant(ctx context.Context, options *EnumClientGetReferencedConstantOptions) (EnumClientGetReferencedConstantResponse, error) {
	var err error
	const operationName = "EnumClient.GetReferencedConstant"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getReferencedConstantCreateRequest(ctx, options)
	if err != nil {
		return EnumClientGetReferencedConstantResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EnumClientGetReferencedConstantResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EnumClientGetReferencedConstantResponse{}, err
	}
	resp, err := client.getReferencedConstantHandleResponse(httpResp)
	return resp, err
}

// getReferencedConstantCreateRequest creates the GetReferencedConstant request.
func (client *EnumClient) getReferencedConstantCreateRequest(ctx context.Context, _ *EnumClientGetReferencedConstantOptions) (*policy.Request, error) {
	urlPath := "/string/enum/ReferencedConstant"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getReferencedConstantHandleResponse handles the GetReferencedConstant response.
func (client *EnumClient) getReferencedConstantHandleResponse(resp *http.Response) (EnumClientGetReferencedConstantResponse, error) {
	result := EnumClientGetReferencedConstantResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RefColorConstant); err != nil {
		return EnumClientGetReferencedConstantResponse{}, err
	}
	return result, nil
}

// PutNotExpandable - Sends value 'red color' from enumeration of 'red color', 'green-color', 'blue_color'
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - stringBody - string body
//   - options - EnumClientPutNotExpandableOptions contains the optional parameters for the EnumClient.PutNotExpandable method.
func (client *EnumClient) PutNotExpandable(ctx context.Context, stringBody Colors, options *EnumClientPutNotExpandableOptions) (EnumClientPutNotExpandableResponse, error) {
	var err error
	const operationName = "EnumClient.PutNotExpandable"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putNotExpandableCreateRequest(ctx, stringBody, options)
	if err != nil {
		return EnumClientPutNotExpandableResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EnumClientPutNotExpandableResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EnumClientPutNotExpandableResponse{}, err
	}
	return EnumClientPutNotExpandableResponse{}, nil
}

// putNotExpandableCreateRequest creates the PutNotExpandable request.
func (client *EnumClient) putNotExpandableCreateRequest(ctx context.Context, stringBody Colors, _ *EnumClientPutNotExpandableOptions) (*policy.Request, error) {
	urlPath := "/string/enum/notExpandable"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, stringBody); err != nil {
		return nil, err
	}
	return req, nil
}

// PutReferenced - Sends value 'red color' from enumeration of 'red color', 'green-color', 'blue_color'
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - enumStringBody - enum string body
//   - options - EnumClientPutReferencedOptions contains the optional parameters for the EnumClient.PutReferenced method.
func (client *EnumClient) PutReferenced(ctx context.Context, enumStringBody Colors, options *EnumClientPutReferencedOptions) (EnumClientPutReferencedResponse, error) {
	var err error
	const operationName = "EnumClient.PutReferenced"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putReferencedCreateRequest(ctx, enumStringBody, options)
	if err != nil {
		return EnumClientPutReferencedResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EnumClientPutReferencedResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EnumClientPutReferencedResponse{}, err
	}
	return EnumClientPutReferencedResponse{}, nil
}

// putReferencedCreateRequest creates the PutReferenced request.
func (client *EnumClient) putReferencedCreateRequest(ctx context.Context, enumStringBody Colors, _ *EnumClientPutReferencedOptions) (*policy.Request, error) {
	urlPath := "/string/enum/Referenced"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, enumStringBody); err != nil {
		return nil, err
	}
	return req, nil
}

// PutReferencedConstant - Sends value 'green-color' from a constant
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - enumStringBody - enum string body
//   - options - EnumClientPutReferencedConstantOptions contains the optional parameters for the EnumClient.PutReferencedConstant
//     method.
func (client *EnumClient) PutReferencedConstant(ctx context.Context, enumStringBody RefColorConstant, options *EnumClientPutReferencedConstantOptions) (EnumClientPutReferencedConstantResponse, error) {
	var err error
	const operationName = "EnumClient.PutReferencedConstant"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putReferencedConstantCreateRequest(ctx, enumStringBody, options)
	if err != nil {
		return EnumClientPutReferencedConstantResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EnumClientPutReferencedConstantResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EnumClientPutReferencedConstantResponse{}, err
	}
	return EnumClientPutReferencedConstantResponse{}, nil
}

// putReferencedConstantCreateRequest creates the PutReferencedConstant request.
func (client *EnumClient) putReferencedConstantCreateRequest(ctx context.Context, enumStringBody RefColorConstant, _ *EnumClientPutReferencedConstantOptions) (*policy.Request, error) {
	urlPath := "/string/enum/ReferencedConstant"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, enumStringBody); err != nil {
		return nil, err
	}
	return req, nil
}
