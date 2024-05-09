// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package datetimegroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"time"
)

// DatetimeHeaderClient contains the methods for the Encode.Datetime namespace.
// Don't use this type directly, use [DatetimeClient.NewDatetimeHeaderClient] instead.
type DatetimeHeaderClient struct {
	internal *azcore.Client
}

// - options - DatetimeHeaderClientDefaultOptions contains the optional parameters for the DatetimeHeaderClient.Default method.
func (client *DatetimeHeaderClient) Default(ctx context.Context, value time.Time, options *DatetimeHeaderClientDefaultOptions) (DatetimeHeaderClientDefaultResponse, error) {
	var err error
	const operationName = "DatetimeHeaderClient.Default"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.defaultCreateRequest(ctx, value, options)
	if err != nil {
		return DatetimeHeaderClientDefaultResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatetimeHeaderClientDefaultResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return DatetimeHeaderClientDefaultResponse{}, err
	}
	return DatetimeHeaderClientDefaultResponse{}, nil
}

// defaultCreateRequest creates the Default request.
func (client *DatetimeHeaderClient) defaultCreateRequest(ctx context.Context, value time.Time, _ *DatetimeHeaderClientDefaultOptions) (*policy.Request, error) {
	urlPath := "/encode/datetime/header/default"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["value"] = []string{value.Format(time.RFC1123)}
	return req, nil
}

// - options - DatetimeHeaderClientRFC3339Options contains the optional parameters for the DatetimeHeaderClient.RFC3339 method.
func (client *DatetimeHeaderClient) RFC3339(ctx context.Context, value time.Time, options *DatetimeHeaderClientRFC3339Options) (DatetimeHeaderClientRFC3339Response, error) {
	var err error
	const operationName = "DatetimeHeaderClient.RFC3339"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.rfc3339CreateRequest(ctx, value, options)
	if err != nil {
		return DatetimeHeaderClientRFC3339Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatetimeHeaderClientRFC3339Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return DatetimeHeaderClientRFC3339Response{}, err
	}
	return DatetimeHeaderClientRFC3339Response{}, nil
}

// rfc3339CreateRequest creates the RFC3339 request.
func (client *DatetimeHeaderClient) rfc3339CreateRequest(ctx context.Context, value time.Time, _ *DatetimeHeaderClientRFC3339Options) (*policy.Request, error) {
	urlPath := "/encode/datetime/header/rfc3339"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["value"] = []string{value.Format(time.RFC3339Nano)}
	return req, nil
}

// - options - DatetimeHeaderClientRFC7231Options contains the optional parameters for the DatetimeHeaderClient.RFC7231 method.
func (client *DatetimeHeaderClient) RFC7231(ctx context.Context, value time.Time, options *DatetimeHeaderClientRFC7231Options) (DatetimeHeaderClientRFC7231Response, error) {
	var err error
	const operationName = "DatetimeHeaderClient.RFC7231"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.rfc7231CreateRequest(ctx, value, options)
	if err != nil {
		return DatetimeHeaderClientRFC7231Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatetimeHeaderClientRFC7231Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return DatetimeHeaderClientRFC7231Response{}, err
	}
	return DatetimeHeaderClientRFC7231Response{}, nil
}

// rfc7231CreateRequest creates the RFC7231 request.
func (client *DatetimeHeaderClient) rfc7231CreateRequest(ctx context.Context, value time.Time, _ *DatetimeHeaderClientRFC7231Options) (*policy.Request, error) {
	urlPath := "/encode/datetime/header/rfc7231"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["value"] = []string{value.Format(time.RFC1123)}
	return req, nil
}

//   - options - DatetimeHeaderClientUnixTimestampOptions contains the optional parameters for the DatetimeHeaderClient.UnixTimestamp
//     method.
func (client *DatetimeHeaderClient) UnixTimestamp(ctx context.Context, value time.Time, options *DatetimeHeaderClientUnixTimestampOptions) (DatetimeHeaderClientUnixTimestampResponse, error) {
	var err error
	const operationName = "DatetimeHeaderClient.UnixTimestamp"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.unixTimestampCreateRequest(ctx, value, options)
	if err != nil {
		return DatetimeHeaderClientUnixTimestampResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatetimeHeaderClientUnixTimestampResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return DatetimeHeaderClientUnixTimestampResponse{}, err
	}
	return DatetimeHeaderClientUnixTimestampResponse{}, nil
}

// unixTimestampCreateRequest creates the UnixTimestamp request.
func (client *DatetimeHeaderClient) unixTimestampCreateRequest(ctx context.Context, value time.Time, _ *DatetimeHeaderClientUnixTimestampOptions) (*policy.Request, error) {
	urlPath := "/encode/datetime/header/unix-timestamp"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["value"] = []string{timeUnix(value).String()}
	return req, nil
}

//   - options - DatetimeHeaderClientUnixTimestampArrayOptions contains the optional parameters for the DatetimeHeaderClient.UnixTimestampArray
//     method.
func (client *DatetimeHeaderClient) UnixTimestampArray(ctx context.Context, value []time.Time, options *DatetimeHeaderClientUnixTimestampArrayOptions) (DatetimeHeaderClientUnixTimestampArrayResponse, error) {
	var err error
	const operationName = "DatetimeHeaderClient.UnixTimestampArray"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.unixTimestampArrayCreateRequest(ctx, value, options)
	if err != nil {
		return DatetimeHeaderClientUnixTimestampArrayResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatetimeHeaderClientUnixTimestampArrayResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return DatetimeHeaderClientUnixTimestampArrayResponse{}, err
	}
	return DatetimeHeaderClientUnixTimestampArrayResponse{}, nil
}

// unixTimestampArrayCreateRequest creates the UnixTimestampArray request.
func (client *DatetimeHeaderClient) unixTimestampArrayCreateRequest(ctx context.Context, value []time.Time, _ *DatetimeHeaderClientUnixTimestampArrayOptions) (*policy.Request, error) {
	urlPath := "/encode/datetime/header/unix-timestamp-array"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["value"] = []string{strings.Join(func() []string {
		encodedValue := make([]string, len(value))
		for i := 0; i < len(value); i++ {
			encodedValue[i] = timeUnix(value[i]).String()
		}
		return encodedValue
	}(), ",")}
	return req, nil
}
