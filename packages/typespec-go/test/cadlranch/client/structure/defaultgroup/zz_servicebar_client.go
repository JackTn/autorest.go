// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package defaultgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// ServiceBarClient contains the methods for the Client.Structure.Service namespace.
// Don't use this type directly, use [ServiceClient.NewServiceBarClient] instead.
type ServiceBarClient struct {
	internal *azcore.Client
	endpoint string
	client   ClientType
}

// - options - ServiceBarClientFiveOptions contains the optional parameters for the ServiceBarClient.Five method.
func (client *ServiceBarClient) Five(ctx context.Context, options *ServiceBarClientFiveOptions) (ServiceBarClientFiveResponse, error) {
	var err error
	const operationName = "ServiceBarClient.Five"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.fiveCreateRequest(ctx, options)
	if err != nil {
		return ServiceBarClientFiveResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceBarClientFiveResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ServiceBarClientFiveResponse{}, err
	}
	return ServiceBarClientFiveResponse{}, nil
}

// fiveCreateRequest creates the Five request.
func (client *ServiceBarClient) fiveCreateRequest(ctx context.Context, _ *ServiceBarClientFiveOptions) (*policy.Request, error) {
	host := "{endpoint}/client/structure/{client}"
	host = strings.ReplaceAll(host, "{endpoint}", client.endpoint)
	host = strings.ReplaceAll(host, "{client}", string(client.client))
	urlPath := "/five"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// - options - ServiceBarClientSixOptions contains the optional parameters for the ServiceBarClient.Six method.
func (client *ServiceBarClient) Six(ctx context.Context, options *ServiceBarClientSixOptions) (ServiceBarClientSixResponse, error) {
	var err error
	const operationName = "ServiceBarClient.Six"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.sixCreateRequest(ctx, options)
	if err != nil {
		return ServiceBarClientSixResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceBarClientSixResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ServiceBarClientSixResponse{}, err
	}
	return ServiceBarClientSixResponse{}, nil
}

// sixCreateRequest creates the Six request.
func (client *ServiceBarClient) sixCreateRequest(ctx context.Context, _ *ServiceBarClientSixOptions) (*policy.Request, error) {
	host := "{endpoint}/client/structure/{client}"
	host = strings.ReplaceAll(host, "{endpoint}", client.endpoint)
	host = strings.ReplaceAll(host, "{client}", string(client.client))
	urlPath := "/six"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}
