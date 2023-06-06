//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DataFlowClient contains the methods for the DataFlow group.
// Don't use this type directly, use a constructor function instead.
type DataFlowClient struct {
	internal *azcore.Client
	endpoint string
}

// BeginCreateOrUpdateDataFlow - Creates or updates a data flow.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - dataFlowName - The data flow name.
//   - dataFlow - Data flow resource definition.
//   - options - DataFlowClientBeginCreateOrUpdateDataFlowOptions contains the optional parameters for the DataFlowClient.BeginCreateOrUpdateDataFlow
//     method.
func (client *DataFlowClient) BeginCreateOrUpdateDataFlow(ctx context.Context, dataFlowName string, dataFlow DataFlowResource, options *DataFlowClientBeginCreateOrUpdateDataFlowOptions) (*runtime.Poller[DataFlowClientCreateOrUpdateDataFlowResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdateDataFlow(ctx, dataFlowName, dataFlow, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[DataFlowClientCreateOrUpdateDataFlowResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[DataFlowClientCreateOrUpdateDataFlowResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdateDataFlow - Creates or updates a data flow.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
func (client *DataFlowClient) createOrUpdateDataFlow(ctx context.Context, dataFlowName string, dataFlow DataFlowResource, options *DataFlowClientBeginCreateOrUpdateDataFlowOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateDataFlowCreateRequest(ctx, dataFlowName, dataFlow, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateDataFlowCreateRequest creates the CreateOrUpdateDataFlow request.
func (client *DataFlowClient) createOrUpdateDataFlowCreateRequest(ctx context.Context, dataFlowName string, dataFlow DataFlowResource, options *DataFlowClientBeginCreateOrUpdateDataFlowOptions) (*policy.Request, error) {
	urlPath := "/dataflows/{dataFlowName}"
	if dataFlowName == "" {
		return nil, errors.New("parameter dataFlowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataFlowName}", url.PathEscape(dataFlowName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, dataFlow); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDeleteDataFlow - Deletes a data flow.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - dataFlowName - The data flow name.
//   - options - DataFlowClientBeginDeleteDataFlowOptions contains the optional parameters for the DataFlowClient.BeginDeleteDataFlow
//     method.
func (client *DataFlowClient) BeginDeleteDataFlow(ctx context.Context, dataFlowName string, options *DataFlowClientBeginDeleteDataFlowOptions) (*runtime.Poller[DataFlowClientDeleteDataFlowResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteDataFlow(ctx, dataFlowName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[DataFlowClientDeleteDataFlowResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[DataFlowClientDeleteDataFlowResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// DeleteDataFlow - Deletes a data flow.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
func (client *DataFlowClient) deleteDataFlow(ctx context.Context, dataFlowName string, options *DataFlowClientBeginDeleteDataFlowOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteDataFlowCreateRequest(ctx, dataFlowName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteDataFlowCreateRequest creates the DeleteDataFlow request.
func (client *DataFlowClient) deleteDataFlowCreateRequest(ctx context.Context, dataFlowName string, options *DataFlowClientBeginDeleteDataFlowOptions) (*policy.Request, error) {
	urlPath := "/dataflows/{dataFlowName}"
	if dataFlowName == "" {
		return nil, errors.New("parameter dataFlowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataFlowName}", url.PathEscape(dataFlowName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetDataFlow - Gets a data flow.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - dataFlowName - The data flow name.
//   - options - DataFlowClientGetDataFlowOptions contains the optional parameters for the DataFlowClient.GetDataFlow method.
func (client *DataFlowClient) GetDataFlow(ctx context.Context, dataFlowName string, options *DataFlowClientGetDataFlowOptions) (DataFlowClientGetDataFlowResponse, error) {
	var err error
	req, err := client.getDataFlowCreateRequest(ctx, dataFlowName, options)
	if err != nil {
		return DataFlowClientGetDataFlowResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DataFlowClientGetDataFlowResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DataFlowClientGetDataFlowResponse{}, err
	}
	resp, err := client.getDataFlowHandleResponse(httpResp)
	return resp, err
}

// getDataFlowCreateRequest creates the GetDataFlow request.
func (client *DataFlowClient) getDataFlowCreateRequest(ctx context.Context, dataFlowName string, options *DataFlowClientGetDataFlowOptions) (*policy.Request, error) {
	urlPath := "/dataflows/{dataFlowName}"
	if dataFlowName == "" {
		return nil, errors.New("parameter dataFlowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataFlowName}", url.PathEscape(dataFlowName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDataFlowHandleResponse handles the GetDataFlow response.
func (client *DataFlowClient) getDataFlowHandleResponse(resp *http.Response) (DataFlowClientGetDataFlowResponse, error) {
	result := DataFlowClientGetDataFlowResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataFlowResource); err != nil {
		return DataFlowClientGetDataFlowResponse{}, err
	}
	return result, nil
}

// NewGetDataFlowsByWorkspacePager - Lists data flows.
//
// Generated from API version 2020-12-01
//   - options - DataFlowClientGetDataFlowsByWorkspaceOptions contains the optional parameters for the DataFlowClient.NewGetDataFlowsByWorkspacePager
//     method.
func (client *DataFlowClient) NewGetDataFlowsByWorkspacePager(options *DataFlowClientGetDataFlowsByWorkspaceOptions) *runtime.Pager[DataFlowClientGetDataFlowsByWorkspaceResponse] {
	return runtime.NewPager(runtime.PagingHandler[DataFlowClientGetDataFlowsByWorkspaceResponse]{
		More: func(page DataFlowClientGetDataFlowsByWorkspaceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DataFlowClientGetDataFlowsByWorkspaceResponse) (DataFlowClientGetDataFlowsByWorkspaceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.getDataFlowsByWorkspaceCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DataFlowClientGetDataFlowsByWorkspaceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DataFlowClientGetDataFlowsByWorkspaceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DataFlowClientGetDataFlowsByWorkspaceResponse{}, runtime.NewResponseError(resp)
			}
			return client.getDataFlowsByWorkspaceHandleResponse(resp)
		},
	})
}

// getDataFlowsByWorkspaceCreateRequest creates the GetDataFlowsByWorkspace request.
func (client *DataFlowClient) getDataFlowsByWorkspaceCreateRequest(ctx context.Context, options *DataFlowClientGetDataFlowsByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/dataflows"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDataFlowsByWorkspaceHandleResponse handles the GetDataFlowsByWorkspace response.
func (client *DataFlowClient) getDataFlowsByWorkspaceHandleResponse(resp *http.Response) (DataFlowClientGetDataFlowsByWorkspaceResponse, error) {
	result := DataFlowClientGetDataFlowsByWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataFlowListResponse); err != nil {
		return DataFlowClientGetDataFlowsByWorkspaceResponse{}, err
	}
	return result, nil
}

// BeginRenameDataFlow - Renames a dataflow.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - dataFlowName - The data flow name.
//   - request - proposed new name.
//   - options - DataFlowClientBeginRenameDataFlowOptions contains the optional parameters for the DataFlowClient.BeginRenameDataFlow
//     method.
func (client *DataFlowClient) BeginRenameDataFlow(ctx context.Context, dataFlowName string, request ArtifactRenameRequest, options *DataFlowClientBeginRenameDataFlowOptions) (*runtime.Poller[DataFlowClientRenameDataFlowResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.renameDataFlow(ctx, dataFlowName, request, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[DataFlowClientRenameDataFlowResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[DataFlowClientRenameDataFlowResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// RenameDataFlow - Renames a dataflow.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
func (client *DataFlowClient) renameDataFlow(ctx context.Context, dataFlowName string, request ArtifactRenameRequest, options *DataFlowClientBeginRenameDataFlowOptions) (*http.Response, error) {
	var err error
	req, err := client.renameDataFlowCreateRequest(ctx, dataFlowName, request, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// renameDataFlowCreateRequest creates the RenameDataFlow request.
func (client *DataFlowClient) renameDataFlowCreateRequest(ctx context.Context, dataFlowName string, request ArtifactRenameRequest, options *DataFlowClientBeginRenameDataFlowOptions) (*policy.Request, error) {
	urlPath := "/dataflows/{dataFlowName}/rename"
	if dataFlowName == "" {
		return nil, errors.New("parameter dataFlowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataFlowName}", url.PathEscape(dataFlowName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, request); err != nil {
		return nil, err
	}
	return req, nil
}
