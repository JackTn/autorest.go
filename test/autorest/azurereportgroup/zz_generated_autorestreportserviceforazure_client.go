//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azurereportgroup

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// AutoRestReportServiceForAzureClient contains the methods for the AutoRestReportServiceForAzure group.
// Don't use this type directly, use NewAutoRestReportServiceForAzureClient() instead.
type AutoRestReportServiceForAzureClient struct {
	pl runtime.Pipeline
}

// NewAutoRestReportServiceForAzureClient creates a new instance of AutoRestReportServiceForAzureClient with the specified values.
// options - pass nil to accept the default values.
func NewAutoRestReportServiceForAzureClient(options *azcore.ClientOptions) *AutoRestReportServiceForAzureClient {
	cp := azcore.ClientOptions{}
	if options != nil {
		cp = *options
	}
	client := &AutoRestReportServiceForAzureClient{
		pl: runtime.NewPipeline(module, version, nil, nil, &cp),
	}
	return client
}

// GetReport - Get test coverage report
// If the operation fails it returns the *Error error type.
// options - AutoRestReportServiceForAzureClientGetReportOptions contains the optional parameters for the AutoRestReportServiceForAzureClient.GetReport
// method.
func (client *AutoRestReportServiceForAzureClient) GetReport(ctx context.Context, options *AutoRestReportServiceForAzureClientGetReportOptions) (AutoRestReportServiceForAzureClientGetReportResponse, error) {
	req, err := client.getReportCreateRequest(ctx, options)
	if err != nil {
		return AutoRestReportServiceForAzureClientGetReportResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AutoRestReportServiceForAzureClientGetReportResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AutoRestReportServiceForAzureClientGetReportResponse{}, client.getReportHandleError(resp)
	}
	return client.getReportHandleResponse(resp)
}

// getReportCreateRequest creates the GetReport request.
func (client *AutoRestReportServiceForAzureClient) getReportCreateRequest(ctx context.Context, options *AutoRestReportServiceForAzureClientGetReportOptions) (*policy.Request, error) {
	urlPath := "/report/azure"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Qualifier != nil {
		reqQP.Set("qualifier", *options.Qualifier)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getReportHandleResponse handles the GetReport response.
func (client *AutoRestReportServiceForAzureClient) getReportHandleResponse(resp *http.Response) (AutoRestReportServiceForAzureClientGetReportResponse, error) {
	result := AutoRestReportServiceForAzureClientGetReportResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return AutoRestReportServiceForAzureClientGetReportResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getReportHandleError handles the GetReport error response.
func (client *AutoRestReportServiceForAzureClient) getReportHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
