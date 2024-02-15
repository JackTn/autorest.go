// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package bytesgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ResponseBodyClient contains the methods for the Encode.Bytes namespace.
// Don't use this type directly, use [BytesClient.NewResponseBodyClient] instead.
type ResponseBodyClient struct {
	internal *azcore.Client
}

// - options - ResponseBodyClientBase64Options contains the optional parameters for the ResponseBodyClient.Base64 method.
func (client *ResponseBodyClient) Base64(ctx context.Context, options *ResponseBodyClientBase64Options) (ResponseBodyClientBase64Response, error) {
	var err error
	req, err := client.base64CreateRequest(ctx, options)
	if err != nil {
		return ResponseBodyClientBase64Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ResponseBodyClientBase64Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ResponseBodyClientBase64Response{}, err
	}
	resp, err := client.base64HandleResponse(httpResp)
	return resp, err
}

// base64CreateRequest creates the Base64 request.
func (client *ResponseBodyClient) base64CreateRequest(ctx context.Context, options *ResponseBodyClientBase64Options) (*policy.Request, error) {
	urlPath := "/encode/bytes/body/response/base64"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// base64HandleResponse handles the Base64 response.
func (client *ResponseBodyClient) base64HandleResponse(resp *http.Response) (ResponseBodyClientBase64Response, error) {
	result := ResponseBodyClientBase64Response{}
	if err := runtime.UnmarshalAsByteArray(resp, &result.Value, runtime.Base64StdFormat); err != nil {
		return ResponseBodyClientBase64Response{}, err
	}
	return result, nil
}

// - options - ResponseBodyClientBase64URLOptions contains the optional parameters for the ResponseBodyClient.Base64URL method.
func (client *ResponseBodyClient) Base64URL(ctx context.Context, options *ResponseBodyClientBase64URLOptions) (ResponseBodyClientBase64URLResponse, error) {
	var err error
	req, err := client.base64URLCreateRequest(ctx, options)
	if err != nil {
		return ResponseBodyClientBase64URLResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ResponseBodyClientBase64URLResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ResponseBodyClientBase64URLResponse{}, err
	}
	resp, err := client.base64URLHandleResponse(httpResp)
	return resp, err
}

// base64URLCreateRequest creates the Base64URL request.
func (client *ResponseBodyClient) base64URLCreateRequest(ctx context.Context, options *ResponseBodyClientBase64URLOptions) (*policy.Request, error) {
	urlPath := "/encode/bytes/body/response/base64url"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// base64URLHandleResponse handles the Base64URL response.
func (client *ResponseBodyClient) base64URLHandleResponse(resp *http.Response) (ResponseBodyClientBase64URLResponse, error) {
	result := ResponseBodyClientBase64URLResponse{}
	if err := runtime.UnmarshalAsByteArray(resp, &result.Value, runtime.Base64URLFormat); err != nil {
		return ResponseBodyClientBase64URLResponse{}, err
	}
	return result, nil
}

//   - options - ResponseBodyClientCustomContentTypeOptions contains the optional parameters for the ResponseBodyClient.CustomContentType
//     method.
func (client *ResponseBodyClient) CustomContentType(ctx context.Context, options *ResponseBodyClientCustomContentTypeOptions) (ResponseBodyClientCustomContentTypeResponse, error) {
	var err error
	req, err := client.customContentTypeCreateRequest(ctx, options)
	if err != nil {
		return ResponseBodyClientCustomContentTypeResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ResponseBodyClientCustomContentTypeResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ResponseBodyClientCustomContentTypeResponse{}, err
	}
	resp, err := client.customContentTypeHandleResponse(httpResp)
	return resp, err
}

// customContentTypeCreateRequest creates the CustomContentType request.
func (client *ResponseBodyClient) customContentTypeCreateRequest(ctx context.Context, options *ResponseBodyClientCustomContentTypeOptions) (*policy.Request, error) {
	urlPath := "/encode/bytes/body/response/custom-content-type"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"image/png"}
	return req, nil
}

// customContentTypeHandleResponse handles the CustomContentType response.
func (client *ResponseBodyClient) customContentTypeHandleResponse(resp *http.Response) (ResponseBodyClientCustomContentTypeResponse, error) {
	result := ResponseBodyClientCustomContentTypeResponse{}
	if err := runtime.UnmarshalAsByteArray(resp, &result.Value, runtime.Base64StdFormat); err != nil {
		return ResponseBodyClientCustomContentTypeResponse{}, err
	}
	return result, nil
}

// - options - ResponseBodyClientDefaultOptions contains the optional parameters for the ResponseBodyClient.Default method.
func (client *ResponseBodyClient) Default(ctx context.Context, options *ResponseBodyClientDefaultOptions) (ResponseBodyClientDefaultResponse, error) {
	var err error
	req, err := client.defaultCreateRequest(ctx, options)
	if err != nil {
		return ResponseBodyClientDefaultResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ResponseBodyClientDefaultResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ResponseBodyClientDefaultResponse{}, err
	}
	resp, err := client.defaultHandleResponse(httpResp)
	return resp, err
}

// defaultCreateRequest creates the Default request.
func (client *ResponseBodyClient) defaultCreateRequest(ctx context.Context, options *ResponseBodyClientDefaultOptions) (*policy.Request, error) {
	urlPath := "/encode/bytes/body/response/default"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// defaultHandleResponse handles the Default response.
func (client *ResponseBodyClient) defaultHandleResponse(resp *http.Response) (ResponseBodyClientDefaultResponse, error) {
	result := ResponseBodyClientDefaultResponse{}
	if err := runtime.UnmarshalAsByteArray(resp, &result.Value, runtime.Base64StdFormat); err != nil {
		return ResponseBodyClientDefaultResponse{}, err
	}
	return result, nil
}

//   - options - ResponseBodyClientOctetStreamOptions contains the optional parameters for the ResponseBodyClient.OctetStream
//     method.
func (client *ResponseBodyClient) OctetStream(ctx context.Context, options *ResponseBodyClientOctetStreamOptions) (ResponseBodyClientOctetStreamResponse, error) {
	var err error
	req, err := client.octetStreamCreateRequest(ctx, options)
	if err != nil {
		return ResponseBodyClientOctetStreamResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ResponseBodyClientOctetStreamResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ResponseBodyClientOctetStreamResponse{}, err
	}
	resp, err := client.octetStreamHandleResponse(httpResp)
	return resp, err
}

// octetStreamCreateRequest creates the OctetStream request.
func (client *ResponseBodyClient) octetStreamCreateRequest(ctx context.Context, options *ResponseBodyClientOctetStreamOptions) (*policy.Request, error) {
	urlPath := "/encode/bytes/body/response/octet-stream"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/octet-stream"}
	return req, nil
}

// octetStreamHandleResponse handles the OctetStream response.
func (client *ResponseBodyClient) octetStreamHandleResponse(resp *http.Response) (ResponseBodyClientOctetStreamResponse, error) {
	result := ResponseBodyClientOctetStreamResponse{}
	if err := runtime.UnmarshalAsByteArray(resp, &result.Value, runtime.Base64StdFormat); err != nil {
		return ResponseBodyClientOctetStreamResponse{}, err
	}
	return result, nil
}
