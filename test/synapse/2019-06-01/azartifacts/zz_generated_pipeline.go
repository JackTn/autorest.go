// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type pipelineClient struct {
	con *connection
}

// Pipeline returns the pipeline associated with this client.
func (client *pipelineClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// CreateOrUpdatePipeline - Creates or updates a pipeline.
func (client *pipelineClient) CreateOrUpdatePipeline(ctx context.Context, pipelineName string, pipeline PipelineResource, options *PipelineCreateOrUpdatePipelineOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdatePipelineCreateRequest(ctx, pipelineName, pipeline, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.CreateOrUpdatePipelineHandleError(resp)
	}
	return resp, nil
}

// CreateOrUpdatePipelineCreateRequest creates the CreateOrUpdatePipeline request.
func (client *pipelineClient) CreateOrUpdatePipelineCreateRequest(ctx context.Context, pipelineName string, pipeline PipelineResource, options *PipelineCreateOrUpdatePipelineOptions) (*azcore.Request, error) {
	urlPath := "/pipelines/{pipelineName}"
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	if options != nil && options.IfMatch != nil {
		req.Header.Set("If-Match", *options.IfMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(pipeline)
}

// CreateOrUpdatePipelineHandleResponse handles the CreateOrUpdatePipeline response.
func (client *pipelineClient) CreateOrUpdatePipelineHandleResponse(resp *azcore.Response) (*PipelineResourceResponse, error) {
	result := PipelineResourceResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PipelineResource)
}

// CreateOrUpdatePipelineHandleError handles the CreateOrUpdatePipeline error response.
func (client *pipelineClient) CreateOrUpdatePipelineHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// CreatePipelineRun - Creates a run of a pipeline.
func (client *pipelineClient) CreatePipelineRun(ctx context.Context, pipelineName string, options *PipelineCreatePipelineRunOptions) (*CreateRunResponseResponse, error) {
	req, err := client.CreatePipelineRunCreateRequest(ctx, pipelineName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusAccepted) {
		return nil, client.CreatePipelineRunHandleError(resp)
	}
	result, err := client.CreatePipelineRunHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreatePipelineRunCreateRequest creates the CreatePipelineRun request.
func (client *pipelineClient) CreatePipelineRunCreateRequest(ctx context.Context, pipelineName string, options *PipelineCreatePipelineRunOptions) (*azcore.Request, error) {
	urlPath := "/pipelines/{pipelineName}/createRun"
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	if options != nil && options.ReferencePipelineRunId != nil {
		query.Set("referencePipelineRunId", *options.ReferencePipelineRunId)
	}
	if options != nil && options.IsRecovery != nil {
		query.Set("isRecovery", strconv.FormatBool(*options.IsRecovery))
	}
	if options != nil && options.StartActivityName != nil {
		query.Set("startActivityName", *options.StartActivityName)
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	if options != nil {
		return req, req.MarshalAsJSON(options.Parameters)
	}
	return req, nil
}

// CreatePipelineRunHandleResponse handles the CreatePipelineRun response.
func (client *pipelineClient) CreatePipelineRunHandleResponse(resp *azcore.Response) (*CreateRunResponseResponse, error) {
	result := CreateRunResponseResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.CreateRunResponse)
}

// CreatePipelineRunHandleError handles the CreatePipelineRun error response.
func (client *pipelineClient) CreatePipelineRunHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// DeletePipeline - Deletes a pipeline.
func (client *pipelineClient) DeletePipeline(ctx context.Context, pipelineName string, options *PipelineDeletePipelineOptions) (*azcore.Response, error) {
	req, err := client.DeletePipelineCreateRequest(ctx, pipelineName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.DeletePipelineHandleError(resp)
	}
	return resp, nil
}

// DeletePipelineCreateRequest creates the DeletePipeline request.
func (client *pipelineClient) DeletePipelineCreateRequest(ctx context.Context, pipelineName string, options *PipelineDeletePipelineOptions) (*azcore.Request, error) {
	urlPath := "/pipelines/{pipelineName}"
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// DeletePipelineHandleError handles the DeletePipeline error response.
func (client *pipelineClient) DeletePipelineHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetPipeline - Gets a pipeline.
func (client *pipelineClient) GetPipeline(ctx context.Context, pipelineName string, options *PipelineGetPipelineOptions) (*PipelineResourceResponse, error) {
	req, err := client.GetPipelineCreateRequest(ctx, pipelineName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNotModified) {
		return nil, client.GetPipelineHandleError(resp)
	}
	result, err := client.GetPipelineHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetPipelineCreateRequest creates the GetPipeline request.
func (client *pipelineClient) GetPipelineCreateRequest(ctx context.Context, pipelineName string, options *PipelineGetPipelineOptions) (*azcore.Request, error) {
	urlPath := "/pipelines/{pipelineName}"
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *options.IfNoneMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetPipelineHandleResponse handles the GetPipeline response.
func (client *pipelineClient) GetPipelineHandleResponse(resp *azcore.Response) (*PipelineResourceResponse, error) {
	result := PipelineResourceResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PipelineResource)
}

// GetPipelineHandleError handles the GetPipeline error response.
func (client *pipelineClient) GetPipelineHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetPipelinesByWorkspace - Lists pipelines.
func (client *pipelineClient) GetPipelinesByWorkspace(options *PipelineGetPipelinesByWorkspaceOptions) PipelineListResponsePager {
	return &pipelineListResponsePager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.GetPipelinesByWorkspaceCreateRequest(ctx, options)
		},
		responder: client.GetPipelinesByWorkspaceHandleResponse,
		errorer:   client.GetPipelinesByWorkspaceHandleError,
		advancer: func(ctx context.Context, resp *PipelineListResponseResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.PipelineListResponse.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// GetPipelinesByWorkspaceCreateRequest creates the GetPipelinesByWorkspace request.
func (client *pipelineClient) GetPipelinesByWorkspaceCreateRequest(ctx context.Context, options *PipelineGetPipelinesByWorkspaceOptions) (*azcore.Request, error) {
	urlPath := "/pipelines"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetPipelinesByWorkspaceHandleResponse handles the GetPipelinesByWorkspace response.
func (client *pipelineClient) GetPipelinesByWorkspaceHandleResponse(resp *azcore.Response) (*PipelineListResponseResponse, error) {
	result := PipelineListResponseResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PipelineListResponse)
}

// GetPipelinesByWorkspaceHandleError handles the GetPipelinesByWorkspace error response.
func (client *pipelineClient) GetPipelinesByWorkspaceHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
