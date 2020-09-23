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
	*client
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *pipelineClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// CreateOrUpdatePipeline - Creates or updates a pipeline.
func (client *pipelineClient) CreateOrUpdatePipeline(ctx context.Context, pipelineName string, pipeline PipelineResource, pipelineCreateOrUpdatePipelineOptions *PipelineCreateOrUpdatePipelineOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdatePipelineCreateRequest(ctx, pipelineName, pipeline, pipelineCreateOrUpdatePipelineOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.CreateOrUpdatePipelineHandleError(resp)
	}
	return resp, nil
}

// CreateOrUpdatePipelineCreateRequest creates the CreateOrUpdatePipeline request.
func (client *pipelineClient) CreateOrUpdatePipelineCreateRequest(ctx context.Context, pipelineName string, pipeline PipelineResource, pipelineCreateOrUpdatePipelineOptions *PipelineCreateOrUpdatePipelineOptions) (*azcore.Request, error) {
	urlPath := "/pipelines/{pipelineName}"
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	if pipelineCreateOrUpdatePipelineOptions != nil && pipelineCreateOrUpdatePipelineOptions.IfMatch != nil {
		req.Header.Set("If-Match", *pipelineCreateOrUpdatePipelineOptions.IfMatch)
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
func (client *pipelineClient) CreatePipelineRun(ctx context.Context, pipelineName string, pipelineCreatePipelineRunOptions *PipelineCreatePipelineRunOptions) (*CreateRunResponseResponse, error) {
	req, err := client.CreatePipelineRunCreateRequest(ctx, pipelineName, pipelineCreatePipelineRunOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
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
func (client *pipelineClient) CreatePipelineRunCreateRequest(ctx context.Context, pipelineName string, pipelineCreatePipelineRunOptions *PipelineCreatePipelineRunOptions) (*azcore.Request, error) {
	urlPath := "/pipelines/{pipelineName}/createRun"
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	if pipelineCreatePipelineRunOptions != nil && pipelineCreatePipelineRunOptions.ReferencePipelineRunId != nil {
		query.Set("referencePipelineRunId", *pipelineCreatePipelineRunOptions.ReferencePipelineRunId)
	}
	if pipelineCreatePipelineRunOptions != nil && pipelineCreatePipelineRunOptions.IsRecovery != nil {
		query.Set("isRecovery", strconv.FormatBool(*pipelineCreatePipelineRunOptions.IsRecovery))
	}
	if pipelineCreatePipelineRunOptions != nil && pipelineCreatePipelineRunOptions.StartActivityName != nil {
		query.Set("startActivityName", *pipelineCreatePipelineRunOptions.StartActivityName)
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	if pipelineCreatePipelineRunOptions != nil {
		return req, req.MarshalAsJSON(pipelineCreatePipelineRunOptions.Parameters)
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
func (client *pipelineClient) DeletePipeline(ctx context.Context, pipelineName string) (*azcore.Response, error) {
	req, err := client.DeletePipelineCreateRequest(ctx, pipelineName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.DeletePipelineHandleError(resp)
	}
	return resp, nil
}

// DeletePipelineCreateRequest creates the DeletePipeline request.
func (client *pipelineClient) DeletePipelineCreateRequest(ctx context.Context, pipelineName string) (*azcore.Request, error) {
	urlPath := "/pipelines/{pipelineName}"
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.u, urlPath))
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
func (client *pipelineClient) GetPipeline(ctx context.Context, pipelineName string, pipelineGetPipelineOptions *PipelineGetPipelineOptions) (*PipelineResourceResponse, error) {
	req, err := client.GetPipelineCreateRequest(ctx, pipelineName, pipelineGetPipelineOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
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
func (client *pipelineClient) GetPipelineCreateRequest(ctx context.Context, pipelineName string, pipelineGetPipelineOptions *PipelineGetPipelineOptions) (*azcore.Request, error) {
	urlPath := "/pipelines/{pipelineName}"
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	if pipelineGetPipelineOptions != nil && pipelineGetPipelineOptions.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *pipelineGetPipelineOptions.IfNoneMatch)
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
func (client *pipelineClient) GetPipelinesByWorkspace() PipelineListResponsePager {
	return &pipelineListResponsePager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.GetPipelinesByWorkspaceCreateRequest(ctx)
		},
		responder: client.GetPipelinesByWorkspaceHandleResponse,
		errorer:   client.GetPipelinesByWorkspaceHandleError,
		advancer: func(ctx context.Context, resp *PipelineListResponseResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.PipelineListResponse.NextLink)
		},
	}
}

// GetPipelinesByWorkspaceCreateRequest creates the GetPipelinesByWorkspace request.
func (client *pipelineClient) GetPipelinesByWorkspaceCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/pipelines"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
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