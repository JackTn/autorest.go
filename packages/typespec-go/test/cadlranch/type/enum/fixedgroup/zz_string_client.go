// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fixedgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// StringClient contains the methods for the Type.Enum.Fixed namespace.
// Don't use this type directly, use [FixedClient.NewStringClient] instead.
type StringClient struct {
	internal *azcore.Client
}

// GetKnownValue - getKnownValue
//   - options - StringClientGetKnownValueOptions contains the optional parameters for the StringClient.GetKnownValue method.
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
	urlPath := "/type/enum/fixed/string/known-value"
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

// PutKnownValue - putKnownValue
//   - body - _
//   - options - StringClientPutKnownValueOptions contains the optional parameters for the StringClient.PutKnownValue method.
func (client *StringClient) PutKnownValue(ctx context.Context, body DaysOfWeekEnum, options *StringClientPutKnownValueOptions) (StringClientPutKnownValueResponse, error) {
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
func (client *StringClient) putKnownValueCreateRequest(ctx context.Context, body DaysOfWeekEnum, options *StringClientPutKnownValueOptions) (*policy.Request, error) {
	urlPath := "/type/enum/fixed/string/known-value"
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

// PutUnknownValue - putUnknownValue
//   - body - _
//   - options - StringClientPutUnknownValueOptions contains the optional parameters for the StringClient.PutUnknownValue method.
func (client *StringClient) PutUnknownValue(ctx context.Context, body DaysOfWeekEnum, options *StringClientPutUnknownValueOptions) (StringClientPutUnknownValueResponse, error) {
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
func (client *StringClient) putUnknownValueCreateRequest(ctx context.Context, body DaysOfWeekEnum, options *StringClientPutUnknownValueOptions) (*policy.Request, error) {
	urlPath := "/type/enum/fixed/string/unknown-value"
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
