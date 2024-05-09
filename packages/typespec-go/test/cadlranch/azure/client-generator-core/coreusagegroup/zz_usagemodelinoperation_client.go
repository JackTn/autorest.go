// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package coreusagegroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// UsageModelInOperationClient contains the methods for the _Specs_.Azure.ClientGenerator.Core.Usage namespace.
// Don't use this type directly, use [UsageClient.NewUsageModelInOperationClient] instead.
type UsageModelInOperationClient struct {
	internal *azcore.Client
}

// InputToInputOutput - Expected body parameter:
// ```json
// {
// "name": <any string>
// }
// ```
//   - options - UsageModelInOperationClientInputToInputOutputOptions contains the optional parameters for the UsageModelInOperationClient.InputToInputOutput
//     method.
func (client *UsageModelInOperationClient) InputToInputOutput(ctx context.Context, body InputModel, options *UsageModelInOperationClientInputToInputOutputOptions) (UsageModelInOperationClientInputToInputOutputResponse, error) {
	var err error
	const operationName = "UsageModelInOperationClient.InputToInputOutput"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.inputToInputOutputCreateRequest(ctx, body, options)
	if err != nil {
		return UsageModelInOperationClientInputToInputOutputResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return UsageModelInOperationClientInputToInputOutputResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return UsageModelInOperationClientInputToInputOutputResponse{}, err
	}
	return UsageModelInOperationClientInputToInputOutputResponse{}, nil
}

// inputToInputOutputCreateRequest creates the InputToInputOutput request.
func (client *UsageModelInOperationClient) inputToInputOutputCreateRequest(ctx context.Context, body InputModel, _ *UsageModelInOperationClientInputToInputOutputOptions) (*policy.Request, error) {
	urlPath := "/azure/client-generator-core/usage/inputToInputOutput"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// OutputToInputOutput - Expected response body:
// ```json
// {
// "name": <any string>
// }
// ```
//   - options - UsageModelInOperationClientOutputToInputOutputOptions contains the optional parameters for the UsageModelInOperationClient.OutputToInputOutput
//     method.
func (client *UsageModelInOperationClient) OutputToInputOutput(ctx context.Context, options *UsageModelInOperationClientOutputToInputOutputOptions) (UsageModelInOperationClientOutputToInputOutputResponse, error) {
	var err error
	const operationName = "UsageModelInOperationClient.OutputToInputOutput"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.outputToInputOutputCreateRequest(ctx, options)
	if err != nil {
		return UsageModelInOperationClientOutputToInputOutputResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return UsageModelInOperationClientOutputToInputOutputResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return UsageModelInOperationClientOutputToInputOutputResponse{}, err
	}
	resp, err := client.outputToInputOutputHandleResponse(httpResp)
	return resp, err
}

// outputToInputOutputCreateRequest creates the OutputToInputOutput request.
func (client *UsageModelInOperationClient) outputToInputOutputCreateRequest(ctx context.Context, _ *UsageModelInOperationClientOutputToInputOutputOptions) (*policy.Request, error) {
	urlPath := "/azure/client-generator-core/usage/outputToInputOutput"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// outputToInputOutputHandleResponse handles the OutputToInputOutput response.
func (client *UsageModelInOperationClient) outputToInputOutputHandleResponse(resp *http.Response) (UsageModelInOperationClientOutputToInputOutputResponse, error) {
	result := UsageModelInOperationClientOutputToInputOutputResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OutputModel); err != nil {
		return UsageModelInOperationClientOutputToInputOutputResponse{}, err
	}
	return result, nil
}
