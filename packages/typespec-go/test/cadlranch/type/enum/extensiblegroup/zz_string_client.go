// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package extensiblegroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// StringClient contains the methods for the Type.Enum.Extensible namespace.
// Don't use this type directly, use [ExtensibleClient.NewStringClient] instead.
type StringClient struct {
	internal *azcore.Client
}

// - options - StringClientGetKnownValueOptions contains the optional parameters for the StringClient.GetKnownValue method.
func (client *StringClient) GetKnownValue(ctx context.Context, options *StringClientGetKnownValueOptions) (StringClientGetKnownValueResponse, error) {
	var err error
	req, err := client.getKnownValueCreateRequest(ctx, options)
	if err != nil {
		return StringClientGetKnownValueResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return StringClientGetKnownValueResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return StringClientGetKnownValueResponse{}, err
	}
	resp, err := client.getKnownValueHandleResponse(httpResp)
	return resp, err
}

// getKnownValueCreateRequest creates the GetKnownValue request.
func (client *StringClient) getKnownValueCreateRequest(ctx context.Context, options *StringClientGetKnownValueOptions) (*policy.Request, error) {
	urlPath := "/type/enum/extensible/string/known-value"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getKnownValueHandleResponse handles the GetKnownValue response.
func (client *StringClient) getKnownValueHandleResponse(resp *http.Response) (StringClientGetKnownValueResponse, error) {
	result := StringClientGetKnownValueResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return StringClientGetKnownValueResponse{}, err
	}
	return result, nil
}

// - options - StringClientGetUnknownValueOptions contains the optional parameters for the StringClient.GetUnknownValue method.
func (client *StringClient) GetUnknownValue(ctx context.Context, options *StringClientGetUnknownValueOptions) (StringClientGetUnknownValueResponse, error) {
	var err error
	req, err := client.getUnknownValueCreateRequest(ctx, options)
	if err != nil {
		return StringClientGetUnknownValueResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return StringClientGetUnknownValueResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return StringClientGetUnknownValueResponse{}, err
	}
	resp, err := client.getUnknownValueHandleResponse(httpResp)
	return resp, err
}

// getUnknownValueCreateRequest creates the GetUnknownValue request.
func (client *StringClient) getUnknownValueCreateRequest(ctx context.Context, options *StringClientGetUnknownValueOptions) (*policy.Request, error) {
	urlPath := "/type/enum/extensible/string/unknown-value"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getUnknownValueHandleResponse handles the GetUnknownValue response.
func (client *StringClient) getUnknownValueHandleResponse(resp *http.Response) (StringClientGetUnknownValueResponse, error) {
	result := StringClientGetUnknownValueResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return StringClientGetUnknownValueResponse{}, err
	}
	return result, nil
}

// - options - StringClientPutKnownValueOptions contains the optional parameters for the StringClient.PutKnownValue method.
func (client *StringClient) PutKnownValue(ctx context.Context, body DaysOfWeekExtensibleEnum, options *StringClientPutKnownValueOptions) (StringClientPutKnownValueResponse, error) {
	var err error
	req, err := client.putKnownValueCreateRequest(ctx, body, options)
	if err != nil {
		return StringClientPutKnownValueResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return StringClientPutKnownValueResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return StringClientPutKnownValueResponse{}, err
	}
	return StringClientPutKnownValueResponse{}, nil
}

// putKnownValueCreateRequest creates the PutKnownValue request.
func (client *StringClient) putKnownValueCreateRequest(ctx context.Context, body DaysOfWeekExtensibleEnum, options *StringClientPutKnownValueOptions) (*policy.Request, error) {
	urlPath := "/type/enum/extensible/string/known-value"
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

// - options - StringClientPutUnknownValueOptions contains the optional parameters for the StringClient.PutUnknownValue method.
func (client *StringClient) PutUnknownValue(ctx context.Context, body DaysOfWeekExtensibleEnum, options *StringClientPutUnknownValueOptions) (StringClientPutUnknownValueResponse, error) {
	var err error
	req, err := client.putUnknownValueCreateRequest(ctx, body, options)
	if err != nil {
		return StringClientPutUnknownValueResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return StringClientPutUnknownValueResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return StringClientPutUnknownValueResponse{}, err
	}
	return StringClientPutUnknownValueResponse{}, nil
}

// putUnknownValueCreateRequest creates the PutUnknownValue request.
func (client *StringClient) putUnknownValueCreateRequest(ctx context.Context, body DaysOfWeekExtensibleEnum, options *StringClientPutUnknownValueOptions) (*policy.Request, error) {
	urlPath := "/type/enum/extensible/string/unknown-value"
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
