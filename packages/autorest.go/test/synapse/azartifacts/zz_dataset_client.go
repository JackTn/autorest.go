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

// DatasetClient contains the methods for the Dataset group.
// Don't use this type directly, use a constructor function instead.
type DatasetClient struct {
	internal *azcore.Client
	endpoint string
}

// BeginCreateOrUpdateDataset - Creates or updates a dataset.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - datasetName - The dataset name.
//   - dataset - Dataset resource definition.
//   - options - DatasetClientBeginCreateOrUpdateDatasetOptions contains the optional parameters for the DatasetClient.BeginCreateOrUpdateDataset
//     method.
func (client *DatasetClient) BeginCreateOrUpdateDataset(ctx context.Context, datasetName string, dataset DatasetResource, options *DatasetClientBeginCreateOrUpdateDatasetOptions) (*runtime.Poller[DatasetClientCreateOrUpdateDatasetResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdateDataset(ctx, datasetName, dataset, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[DatasetClientCreateOrUpdateDatasetResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[DatasetClientCreateOrUpdateDatasetResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdateDataset - Creates or updates a dataset.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
func (client *DatasetClient) createOrUpdateDataset(ctx context.Context, datasetName string, dataset DatasetResource, options *DatasetClientBeginCreateOrUpdateDatasetOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateDatasetCreateRequest(ctx, datasetName, dataset, options)
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

// createOrUpdateDatasetCreateRequest creates the CreateOrUpdateDataset request.
func (client *DatasetClient) createOrUpdateDatasetCreateRequest(ctx context.Context, datasetName string, dataset DatasetResource, options *DatasetClientBeginCreateOrUpdateDatasetOptions) (*policy.Request, error) {
	urlPath := "/datasets/{datasetName}"
	if datasetName == "" {
		return nil, errors.New("parameter datasetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{datasetName}", url.PathEscape(datasetName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	if err := runtime.MarshalAsJSON(req, dataset); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDeleteDataset - Deletes a dataset.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - datasetName - The dataset name.
//   - options - DatasetClientBeginDeleteDatasetOptions contains the optional parameters for the DatasetClient.BeginDeleteDataset
//     method.
func (client *DatasetClient) BeginDeleteDataset(ctx context.Context, datasetName string, options *DatasetClientBeginDeleteDatasetOptions) (*runtime.Poller[DatasetClientDeleteDatasetResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteDataset(ctx, datasetName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[DatasetClientDeleteDatasetResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[DatasetClientDeleteDatasetResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// DeleteDataset - Deletes a dataset.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
func (client *DatasetClient) deleteDataset(ctx context.Context, datasetName string, options *DatasetClientBeginDeleteDatasetOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteDatasetCreateRequest(ctx, datasetName, options)
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

// deleteDatasetCreateRequest creates the DeleteDataset request.
func (client *DatasetClient) deleteDatasetCreateRequest(ctx context.Context, datasetName string, _ *DatasetClientBeginDeleteDatasetOptions) (*policy.Request, error) {
	urlPath := "/datasets/{datasetName}"
	if datasetName == "" {
		return nil, errors.New("parameter datasetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{datasetName}", url.PathEscape(datasetName))
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

// GetDataset - Gets a dataset.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - datasetName - The dataset name.
//   - options - DatasetClientGetDatasetOptions contains the optional parameters for the DatasetClient.GetDataset method.
func (client *DatasetClient) GetDataset(ctx context.Context, datasetName string, options *DatasetClientGetDatasetOptions) (DatasetClientGetDatasetResponse, error) {
	var err error
	req, err := client.getDatasetCreateRequest(ctx, datasetName, options)
	if err != nil {
		return DatasetClientGetDatasetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatasetClientGetDatasetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNotModified) {
		err = runtime.NewResponseError(httpResp)
		return DatasetClientGetDatasetResponse{}, err
	}
	resp, err := client.getDatasetHandleResponse(httpResp)
	return resp, err
}

// getDatasetCreateRequest creates the GetDataset request.
func (client *DatasetClient) getDatasetCreateRequest(ctx context.Context, datasetName string, options *DatasetClientGetDatasetOptions) (*policy.Request, error) {
	urlPath := "/datasets/{datasetName}"
	if datasetName == "" {
		return nil, errors.New("parameter datasetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{datasetName}", url.PathEscape(datasetName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	return req, nil
}

// getDatasetHandleResponse handles the GetDataset response.
func (client *DatasetClient) getDatasetHandleResponse(resp *http.Response) (DatasetClientGetDatasetResponse, error) {
	result := DatasetClientGetDatasetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatasetResource); err != nil {
		return DatasetClientGetDatasetResponse{}, err
	}
	return result, nil
}

// NewGetDatasetsByWorkspacePager - Lists datasets.
//
// Generated from API version 2020-12-01
//   - options - DatasetClientGetDatasetsByWorkspaceOptions contains the optional parameters for the DatasetClient.NewGetDatasetsByWorkspacePager
//     method.
func (client *DatasetClient) NewGetDatasetsByWorkspacePager(options *DatasetClientGetDatasetsByWorkspaceOptions) *runtime.Pager[DatasetClientGetDatasetsByWorkspaceResponse] {
	return runtime.NewPager(runtime.PagingHandler[DatasetClientGetDatasetsByWorkspaceResponse]{
		More: func(page DatasetClientGetDatasetsByWorkspaceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DatasetClientGetDatasetsByWorkspaceResponse) (DatasetClientGetDatasetsByWorkspaceResponse, error) {
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.getDatasetsByWorkspaceCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return DatasetClientGetDatasetsByWorkspaceResponse{}, err
			}
			return client.getDatasetsByWorkspaceHandleResponse(resp)
		},
	})
}

// getDatasetsByWorkspaceCreateRequest creates the GetDatasetsByWorkspace request.
func (client *DatasetClient) getDatasetsByWorkspaceCreateRequest(ctx context.Context, _ *DatasetClientGetDatasetsByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/datasets"
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

// getDatasetsByWorkspaceHandleResponse handles the GetDatasetsByWorkspace response.
func (client *DatasetClient) getDatasetsByWorkspaceHandleResponse(resp *http.Response) (DatasetClientGetDatasetsByWorkspaceResponse, error) {
	result := DatasetClientGetDatasetsByWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatasetListResponse); err != nil {
		return DatasetClientGetDatasetsByWorkspaceResponse{}, err
	}
	return result, nil
}

// BeginRenameDataset - Renames a dataset.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - datasetName - The dataset name.
//   - request - proposed new name.
//   - options - DatasetClientBeginRenameDatasetOptions contains the optional parameters for the DatasetClient.BeginRenameDataset
//     method.
func (client *DatasetClient) BeginRenameDataset(ctx context.Context, datasetName string, request ArtifactRenameRequest, options *DatasetClientBeginRenameDatasetOptions) (*runtime.Poller[DatasetClientRenameDatasetResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.renameDataset(ctx, datasetName, request, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[DatasetClientRenameDatasetResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[DatasetClientRenameDatasetResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// RenameDataset - Renames a dataset.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
func (client *DatasetClient) renameDataset(ctx context.Context, datasetName string, request ArtifactRenameRequest, options *DatasetClientBeginRenameDatasetOptions) (*http.Response, error) {
	var err error
	req, err := client.renameDatasetCreateRequest(ctx, datasetName, request, options)
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

// renameDatasetCreateRequest creates the RenameDataset request.
func (client *DatasetClient) renameDatasetCreateRequest(ctx context.Context, datasetName string, request ArtifactRenameRequest, _ *DatasetClientBeginRenameDatasetOptions) (*policy.Request, error) {
	urlPath := "/datasets/{datasetName}/rename"
	if datasetName == "" {
		return nil, errors.New("parameter datasetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{datasetName}", url.PathEscape(datasetName))
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
