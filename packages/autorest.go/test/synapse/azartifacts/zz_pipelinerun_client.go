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
	"strconv"
	"strings"
)

// PipelineRunClient contains the methods for the PipelineRun group.
// Don't use this type directly, use a constructor function instead.
type PipelineRunClient struct {
	internal *azcore.Client
	endpoint string
}

// CancelPipelineRun - Cancel a pipeline run by its run ID.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - runID - The pipeline run identifier.
//   - options - PipelineRunClientCancelPipelineRunOptions contains the optional parameters for the PipelineRunClient.CancelPipelineRun
//     method.
func (client *PipelineRunClient) CancelPipelineRun(ctx context.Context, runID string, options *PipelineRunClientCancelPipelineRunOptions) (PipelineRunClientCancelPipelineRunResponse, error) {
	var err error
	req, err := client.cancelPipelineRunCreateRequest(ctx, runID, options)
	if err != nil {
		return PipelineRunClientCancelPipelineRunResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PipelineRunClientCancelPipelineRunResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PipelineRunClientCancelPipelineRunResponse{}, err
	}
	return PipelineRunClientCancelPipelineRunResponse{}, nil
}

// cancelPipelineRunCreateRequest creates the CancelPipelineRun request.
func (client *PipelineRunClient) cancelPipelineRunCreateRequest(ctx context.Context, runID string, options *PipelineRunClientCancelPipelineRunOptions) (*policy.Request, error) {
	urlPath := "/pipelineruns/{runId}/cancel"
	if runID == "" {
		return nil, errors.New("parameter runID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runId}", url.PathEscape(runID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	if options != nil && options.IsRecursive != nil {
		reqQP.Set("isRecursive", strconv.FormatBool(*options.IsRecursive))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetPipelineRun - Get a pipeline run by its run ID.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - runID - The pipeline run identifier.
//   - options - PipelineRunClientGetPipelineRunOptions contains the optional parameters for the PipelineRunClient.GetPipelineRun
//     method.
func (client *PipelineRunClient) GetPipelineRun(ctx context.Context, runID string, options *PipelineRunClientGetPipelineRunOptions) (PipelineRunClientGetPipelineRunResponse, error) {
	var err error
	req, err := client.getPipelineRunCreateRequest(ctx, runID, options)
	if err != nil {
		return PipelineRunClientGetPipelineRunResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PipelineRunClientGetPipelineRunResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PipelineRunClientGetPipelineRunResponse{}, err
	}
	resp, err := client.getPipelineRunHandleResponse(httpResp)
	return resp, err
}

// getPipelineRunCreateRequest creates the GetPipelineRun request.
func (client *PipelineRunClient) getPipelineRunCreateRequest(ctx context.Context, runID string, _ *PipelineRunClientGetPipelineRunOptions) (*policy.Request, error) {
	urlPath := "/pipelineruns/{runId}"
	if runID == "" {
		return nil, errors.New("parameter runID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runId}", url.PathEscape(runID))
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

// getPipelineRunHandleResponse handles the GetPipelineRun response.
func (client *PipelineRunClient) getPipelineRunHandleResponse(resp *http.Response) (PipelineRunClientGetPipelineRunResponse, error) {
	result := PipelineRunClientGetPipelineRunResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PipelineRun); err != nil {
		return PipelineRunClientGetPipelineRunResponse{}, err
	}
	return result, nil
}

// QueryActivityRuns - Query activity runs based on input filter conditions.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - pipelineName - The pipeline name.
//   - runID - The pipeline run identifier.
//   - filterParameters - Parameters to filter the activity runs.
//   - options - PipelineRunClientQueryActivityRunsOptions contains the optional parameters for the PipelineRunClient.QueryActivityRuns
//     method.
func (client *PipelineRunClient) QueryActivityRuns(ctx context.Context, pipelineName string, runID string, filterParameters RunFilterParameters, options *PipelineRunClientQueryActivityRunsOptions) (PipelineRunClientQueryActivityRunsResponse, error) {
	var err error
	req, err := client.queryActivityRunsCreateRequest(ctx, pipelineName, runID, filterParameters, options)
	if err != nil {
		return PipelineRunClientQueryActivityRunsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PipelineRunClientQueryActivityRunsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PipelineRunClientQueryActivityRunsResponse{}, err
	}
	resp, err := client.queryActivityRunsHandleResponse(httpResp)
	return resp, err
}

// queryActivityRunsCreateRequest creates the QueryActivityRuns request.
func (client *PipelineRunClient) queryActivityRunsCreateRequest(ctx context.Context, pipelineName string, runID string, filterParameters RunFilterParameters, _ *PipelineRunClientQueryActivityRunsOptions) (*policy.Request, error) {
	urlPath := "/pipelines/{pipelineName}/pipelineruns/{runId}/queryActivityruns"
	if pipelineName == "" {
		return nil, errors.New("parameter pipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	if runID == "" {
		return nil, errors.New("parameter runID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runId}", url.PathEscape(runID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, filterParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// queryActivityRunsHandleResponse handles the QueryActivityRuns response.
func (client *PipelineRunClient) queryActivityRunsHandleResponse(resp *http.Response) (PipelineRunClientQueryActivityRunsResponse, error) {
	result := PipelineRunClientQueryActivityRunsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ActivityRunsQueryResponse); err != nil {
		return PipelineRunClientQueryActivityRunsResponse{}, err
	}
	return result, nil
}

// QueryPipelineRunsByWorkspace - Query pipeline runs in the workspace based on input filter conditions.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - filterParameters - Parameters to filter the pipeline run.
//   - options - PipelineRunClientQueryPipelineRunsByWorkspaceOptions contains the optional parameters for the PipelineRunClient.QueryPipelineRunsByWorkspace
//     method.
func (client *PipelineRunClient) QueryPipelineRunsByWorkspace(ctx context.Context, filterParameters RunFilterParameters, options *PipelineRunClientQueryPipelineRunsByWorkspaceOptions) (PipelineRunClientQueryPipelineRunsByWorkspaceResponse, error) {
	var err error
	req, err := client.queryPipelineRunsByWorkspaceCreateRequest(ctx, filterParameters, options)
	if err != nil {
		return PipelineRunClientQueryPipelineRunsByWorkspaceResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PipelineRunClientQueryPipelineRunsByWorkspaceResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PipelineRunClientQueryPipelineRunsByWorkspaceResponse{}, err
	}
	resp, err := client.queryPipelineRunsByWorkspaceHandleResponse(httpResp)
	return resp, err
}

// queryPipelineRunsByWorkspaceCreateRequest creates the QueryPipelineRunsByWorkspace request.
func (client *PipelineRunClient) queryPipelineRunsByWorkspaceCreateRequest(ctx context.Context, filterParameters RunFilterParameters, _ *PipelineRunClientQueryPipelineRunsByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/queryPipelineRuns"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, filterParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// queryPipelineRunsByWorkspaceHandleResponse handles the QueryPipelineRunsByWorkspace response.
func (client *PipelineRunClient) queryPipelineRunsByWorkspaceHandleResponse(resp *http.Response) (PipelineRunClientQueryPipelineRunsByWorkspaceResponse, error) {
	result := PipelineRunClientQueryPipelineRunsByWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PipelineRunsQueryResponse); err != nil {
		return PipelineRunClientQueryPipelineRunsByWorkspaceResponse{}, err
	}
	return result, nil
}
