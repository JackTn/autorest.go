// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package renamedopgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// RenamedOperationClient contains the methods for the Client.Structure.Service namespace.
// Don't use this type directly, use a constructor function instead.
type RenamedOperationClient struct {
	internal *azcore.Client
	endpoint string
	client   ClientType
}

// NewRenamedOperationGroupClient creates a new instance of [RenamedOperationGroupClient].
func (client *RenamedOperationClient) NewRenamedOperationGroupClient() *RenamedOperationGroupClient {
	return &RenamedOperationGroupClient{
		internal: client.internal,
		endpoint: client.endpoint,
		client:   client.client,
	}
}

//   - options - RenamedOperationClientRenamedFiveOptions contains the optional parameters for the RenamedOperationClient.RenamedFive
//     method.
func (client *RenamedOperationClient) RenamedFive(ctx context.Context, options *RenamedOperationClientRenamedFiveOptions) (RenamedOperationClientRenamedFiveResponse, error) {
	var err error
	const operationName = "RenamedOperationClient.RenamedFive"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.renamedFiveCreateRequest(ctx, options)
	if err != nil {
		return RenamedOperationClientRenamedFiveResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RenamedOperationClientRenamedFiveResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return RenamedOperationClientRenamedFiveResponse{}, err
	}
	return RenamedOperationClientRenamedFiveResponse{}, nil
}

// renamedFiveCreateRequest creates the RenamedFive request.
func (client *RenamedOperationClient) renamedFiveCreateRequest(ctx context.Context, _ *RenamedOperationClientRenamedFiveOptions) (*policy.Request, error) {
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

//   - options - RenamedOperationClientRenamedOneOptions contains the optional parameters for the RenamedOperationClient.RenamedOne
//     method.
func (client *RenamedOperationClient) RenamedOne(ctx context.Context, options *RenamedOperationClientRenamedOneOptions) (RenamedOperationClientRenamedOneResponse, error) {
	var err error
	const operationName = "RenamedOperationClient.RenamedOne"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.renamedOneCreateRequest(ctx, options)
	if err != nil {
		return RenamedOperationClientRenamedOneResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RenamedOperationClientRenamedOneResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return RenamedOperationClientRenamedOneResponse{}, err
	}
	return RenamedOperationClientRenamedOneResponse{}, nil
}

// renamedOneCreateRequest creates the RenamedOne request.
func (client *RenamedOperationClient) renamedOneCreateRequest(ctx context.Context, _ *RenamedOperationClientRenamedOneOptions) (*policy.Request, error) {
	host := "{endpoint}/client/structure/{client}"
	host = strings.ReplaceAll(host, "{endpoint}", client.endpoint)
	host = strings.ReplaceAll(host, "{client}", string(client.client))
	urlPath := "/one"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

//   - options - RenamedOperationClientRenamedThreeOptions contains the optional parameters for the RenamedOperationClient.RenamedThree
//     method.
func (client *RenamedOperationClient) RenamedThree(ctx context.Context, options *RenamedOperationClientRenamedThreeOptions) (RenamedOperationClientRenamedThreeResponse, error) {
	var err error
	const operationName = "RenamedOperationClient.RenamedThree"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.renamedThreeCreateRequest(ctx, options)
	if err != nil {
		return RenamedOperationClientRenamedThreeResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RenamedOperationClientRenamedThreeResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return RenamedOperationClientRenamedThreeResponse{}, err
	}
	return RenamedOperationClientRenamedThreeResponse{}, nil
}

// renamedThreeCreateRequest creates the RenamedThree request.
func (client *RenamedOperationClient) renamedThreeCreateRequest(ctx context.Context, _ *RenamedOperationClientRenamedThreeOptions) (*policy.Request, error) {
	host := "{endpoint}/client/structure/{client}"
	host = strings.ReplaceAll(host, "{endpoint}", client.endpoint)
	host = strings.ReplaceAll(host, "{client}", string(client.client))
	urlPath := "/three"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}
