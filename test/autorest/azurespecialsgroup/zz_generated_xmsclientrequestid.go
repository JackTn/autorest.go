// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azurespecialsgroup

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
)

// XMSClientRequestIDOperations contains the methods for the XMSClientRequestID group.
type XMSClientRequestIDOperations interface {
	// Get - Get method that overwrites x-ms-client-request header with value 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0.
	Get(ctx context.Context, options *XMSClientRequestIDGetOptions) (*http.Response, error)
	// ParamGet - Get method that overwrites x-ms-client-request header with value 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0.
	ParamGet(ctx context.Context, xMSClientRequestId string, options *XMSClientRequestIDParamGetOptions) (*http.Response, error)
}

// XMSClientRequestIDClient implements the XMSClientRequestIDOperations interface.
// Don't use this type directly, use NewXMSClientRequestIDClient() instead.
type XMSClientRequestIDClient struct {
	con *Connection
}

// NewXMSClientRequestIDClient creates a new instance of XMSClientRequestIDClient with the specified values.
func NewXMSClientRequestIDClient(con *Connection) XMSClientRequestIDOperations {
	return &XMSClientRequestIDClient{con: con}
}

// Pipeline returns the pipeline associated with this client.
func (client *XMSClientRequestIDClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// Get - Get method that overwrites x-ms-client-request header with value 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0.
func (client *XMSClientRequestIDClient) Get(ctx context.Context, options *XMSClientRequestIDGetOptions) (*http.Response, error) {
	req, err := client.GetCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetHandleError(resp)
	}
	return resp.Response, nil
}

// GetCreateRequest creates the Get request.
func (client *XMSClientRequestIDClient) GetCreateRequest(ctx context.Context, options *XMSClientRequestIDGetOptions) (*azcore.Request, error) {
	urlPath := "/azurespecials/overwrite/x-ms-client-request-id/method/"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// GetHandleError handles the Get error response.
func (client *XMSClientRequestIDClient) GetHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ParamGet - Get method that overwrites x-ms-client-request header with value 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0.
func (client *XMSClientRequestIDClient) ParamGet(ctx context.Context, xMSClientRequestId string, options *XMSClientRequestIDParamGetOptions) (*http.Response, error) {
	req, err := client.ParamGetCreateRequest(ctx, xMSClientRequestId, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ParamGetHandleError(resp)
	}
	return resp.Response, nil
}

// ParamGetCreateRequest creates the ParamGet request.
func (client *XMSClientRequestIDClient) ParamGetCreateRequest(ctx context.Context, xMSClientRequestId string, options *XMSClientRequestIDParamGetOptions) (*azcore.Request, error) {
	urlPath := "/azurespecials/overwrite/x-ms-client-request-id/via-param/method/"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-ms-client-request-id", xMSClientRequestId)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ParamGetHandleError handles the ParamGet error response.
func (client *XMSClientRequestIDClient) ParamGetHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
