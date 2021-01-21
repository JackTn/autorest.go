// +build go1.13

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
	"strings"
)

type sparkJobDefinitionClient struct {
	con *connection
}

// CreateOrUpdateSparkJobDefinition - Creates or updates a Spark Job Definition.
func (client *sparkJobDefinitionClient) createOrUpdateSparkJobDefinition(ctx context.Context, sparkJobDefinitionName string, sparkJobDefinition SparkJobDefinitionResource, options *SparkJobDefinitionBeginCreateOrUpdateSparkJobDefinitionOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateSparkJobDefinitionCreateRequest(ctx, sparkJobDefinitionName, sparkJobDefinition, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.createOrUpdateSparkJobDefinitionHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateSparkJobDefinitionCreateRequest creates the CreateOrUpdateSparkJobDefinition request.
func (client *sparkJobDefinitionClient) createOrUpdateSparkJobDefinitionCreateRequest(ctx context.Context, sparkJobDefinitionName string, sparkJobDefinition SparkJobDefinitionResource, options *SparkJobDefinitionBeginCreateOrUpdateSparkJobDefinitionOptions) (*azcore.Request, error) {
	urlPath := "/sparkJobDefinitions/{sparkJobDefinitionName}"
	urlPath = strings.ReplaceAll(urlPath, "{sparkJobDefinitionName}", url.PathEscape(sparkJobDefinitionName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	if options != nil && options.IfMatch != nil {
		req.Header.Set("If-Match", *options.IfMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(sparkJobDefinition)
}

// createOrUpdateSparkJobDefinitionHandleResponse handles the CreateOrUpdateSparkJobDefinition response.
func (client *sparkJobDefinitionClient) createOrUpdateSparkJobDefinitionHandleResponse(resp *azcore.Response) (SparkJobDefinitionResourceResponse, error) {
	var val *SparkJobDefinitionResource
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return SparkJobDefinitionResourceResponse{}, err
	}
	return SparkJobDefinitionResourceResponse{RawResponse: resp.Response, SparkJobDefinitionResource: val}, nil
}

// createOrUpdateSparkJobDefinitionHandleError handles the CreateOrUpdateSparkJobDefinition error response.
func (client *sparkJobDefinitionClient) createOrUpdateSparkJobDefinitionHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// DebugSparkJobDefinition - Debug the spark job definition.
func (client *sparkJobDefinitionClient) debugSparkJobDefinition(ctx context.Context, sparkJobDefinitionAzureResource SparkJobDefinitionResource, options *SparkJobDefinitionBeginDebugSparkJobDefinitionOptions) (*azcore.Response, error) {
	req, err := client.debugSparkJobDefinitionCreateRequest(ctx, sparkJobDefinitionAzureResource, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.debugSparkJobDefinitionHandleError(resp)
	}
	return resp, nil
}

// debugSparkJobDefinitionCreateRequest creates the DebugSparkJobDefinition request.
func (client *sparkJobDefinitionClient) debugSparkJobDefinitionCreateRequest(ctx context.Context, sparkJobDefinitionAzureResource SparkJobDefinitionResource, options *SparkJobDefinitionBeginDebugSparkJobDefinitionOptions) (*azcore.Request, error) {
	urlPath := "/debugSparkJobDefinition"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(sparkJobDefinitionAzureResource)
}

// debugSparkJobDefinitionHandleResponse handles the DebugSparkJobDefinition response.
func (client *sparkJobDefinitionClient) debugSparkJobDefinitionHandleResponse(resp *azcore.Response) (SparkBatchJobResponse, error) {
	var val *SparkBatchJob
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return SparkBatchJobResponse{}, err
	}
	return SparkBatchJobResponse{RawResponse: resp.Response, SparkBatchJob: val}, nil
}

// debugSparkJobDefinitionHandleError handles the DebugSparkJobDefinition error response.
func (client *sparkJobDefinitionClient) debugSparkJobDefinitionHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// DeleteSparkJobDefinition - Deletes a Spark Job Definition.
func (client *sparkJobDefinitionClient) deleteSparkJobDefinition(ctx context.Context, sparkJobDefinitionName string, options *SparkJobDefinitionBeginDeleteSparkJobDefinitionOptions) (*azcore.Response, error) {
	req, err := client.deleteSparkJobDefinitionCreateRequest(ctx, sparkJobDefinitionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteSparkJobDefinitionHandleError(resp)
	}
	return resp, nil
}

// deleteSparkJobDefinitionCreateRequest creates the DeleteSparkJobDefinition request.
func (client *sparkJobDefinitionClient) deleteSparkJobDefinitionCreateRequest(ctx context.Context, sparkJobDefinitionName string, options *SparkJobDefinitionBeginDeleteSparkJobDefinitionOptions) (*azcore.Request, error) {
	urlPath := "/sparkJobDefinitions/{sparkJobDefinitionName}"
	urlPath = strings.ReplaceAll(urlPath, "{sparkJobDefinitionName}", url.PathEscape(sparkJobDefinitionName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteSparkJobDefinitionHandleError handles the DeleteSparkJobDefinition error response.
func (client *sparkJobDefinitionClient) deleteSparkJobDefinitionHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ExecuteSparkJobDefinition - Executes the spark job definition.
func (client *sparkJobDefinitionClient) executeSparkJobDefinition(ctx context.Context, sparkJobDefinitionName string, options *SparkJobDefinitionBeginExecuteSparkJobDefinitionOptions) (*azcore.Response, error) {
	req, err := client.executeSparkJobDefinitionCreateRequest(ctx, sparkJobDefinitionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.executeSparkJobDefinitionHandleError(resp)
	}
	return resp, nil
}

// executeSparkJobDefinitionCreateRequest creates the ExecuteSparkJobDefinition request.
func (client *sparkJobDefinitionClient) executeSparkJobDefinitionCreateRequest(ctx context.Context, sparkJobDefinitionName string, options *SparkJobDefinitionBeginExecuteSparkJobDefinitionOptions) (*azcore.Request, error) {
	urlPath := "/sparkJobDefinitions/{sparkJobDefinitionName}/execute"
	urlPath = strings.ReplaceAll(urlPath, "{sparkJobDefinitionName}", url.PathEscape(sparkJobDefinitionName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// executeSparkJobDefinitionHandleResponse handles the ExecuteSparkJobDefinition response.
func (client *sparkJobDefinitionClient) executeSparkJobDefinitionHandleResponse(resp *azcore.Response) (SparkBatchJobResponse, error) {
	var val *SparkBatchJob
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return SparkBatchJobResponse{}, err
	}
	return SparkBatchJobResponse{RawResponse: resp.Response, SparkBatchJob: val}, nil
}

// executeSparkJobDefinitionHandleError handles the ExecuteSparkJobDefinition error response.
func (client *sparkJobDefinitionClient) executeSparkJobDefinitionHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetSparkJobDefinition - Gets a Spark Job Definition.
func (client *sparkJobDefinitionClient) GetSparkJobDefinition(ctx context.Context, sparkJobDefinitionName string, options *SparkJobDefinitionGetSparkJobDefinitionOptions) (SparkJobDefinitionResourceResponse, error) {
	req, err := client.getSparkJobDefinitionCreateRequest(ctx, sparkJobDefinitionName, options)
	if err != nil {
		return SparkJobDefinitionResourceResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SparkJobDefinitionResourceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNotModified) {
		return SparkJobDefinitionResourceResponse{}, client.getSparkJobDefinitionHandleError(resp)
	}
	return client.getSparkJobDefinitionHandleResponse(resp)
}

// getSparkJobDefinitionCreateRequest creates the GetSparkJobDefinition request.
func (client *sparkJobDefinitionClient) getSparkJobDefinitionCreateRequest(ctx context.Context, sparkJobDefinitionName string, options *SparkJobDefinitionGetSparkJobDefinitionOptions) (*azcore.Request, error) {
	urlPath := "/sparkJobDefinitions/{sparkJobDefinitionName}"
	urlPath = strings.ReplaceAll(urlPath, "{sparkJobDefinitionName}", url.PathEscape(sparkJobDefinitionName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *options.IfNoneMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getSparkJobDefinitionHandleResponse handles the GetSparkJobDefinition response.
func (client *sparkJobDefinitionClient) getSparkJobDefinitionHandleResponse(resp *azcore.Response) (SparkJobDefinitionResourceResponse, error) {
	var val *SparkJobDefinitionResource
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return SparkJobDefinitionResourceResponse{}, err
	}
	return SparkJobDefinitionResourceResponse{RawResponse: resp.Response, SparkJobDefinitionResource: val}, nil
}

// getSparkJobDefinitionHandleError handles the GetSparkJobDefinition error response.
func (client *sparkJobDefinitionClient) getSparkJobDefinitionHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetSparkJobDefinitionsByWorkspace - Lists spark job definitions.
func (client *sparkJobDefinitionClient) GetSparkJobDefinitionsByWorkspace(options *SparkJobDefinitionGetSparkJobDefinitionsByWorkspaceOptions) SparkJobDefinitionsListResponsePager {
	return &sparkJobDefinitionsListResponsePager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.getSparkJobDefinitionsByWorkspaceCreateRequest(ctx, options)
		},
		responder: client.getSparkJobDefinitionsByWorkspaceHandleResponse,
		errorer:   client.getSparkJobDefinitionsByWorkspaceHandleError,
		advancer: func(ctx context.Context, resp SparkJobDefinitionsListResponseResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.SparkJobDefinitionsListResponse.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// getSparkJobDefinitionsByWorkspaceCreateRequest creates the GetSparkJobDefinitionsByWorkspace request.
func (client *sparkJobDefinitionClient) getSparkJobDefinitionsByWorkspaceCreateRequest(ctx context.Context, options *SparkJobDefinitionGetSparkJobDefinitionsByWorkspaceOptions) (*azcore.Request, error) {
	urlPath := "/sparkJobDefinitions"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getSparkJobDefinitionsByWorkspaceHandleResponse handles the GetSparkJobDefinitionsByWorkspace response.
func (client *sparkJobDefinitionClient) getSparkJobDefinitionsByWorkspaceHandleResponse(resp *azcore.Response) (SparkJobDefinitionsListResponseResponse, error) {
	var val *SparkJobDefinitionsListResponse
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return SparkJobDefinitionsListResponseResponse{}, err
	}
	return SparkJobDefinitionsListResponseResponse{RawResponse: resp.Response, SparkJobDefinitionsListResponse: val}, nil
}

// getSparkJobDefinitionsByWorkspaceHandleError handles the GetSparkJobDefinitionsByWorkspace error response.
func (client *sparkJobDefinitionClient) getSparkJobDefinitionsByWorkspaceHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// RenameSparkJobDefinition - Renames a sparkJobDefinition.
func (client *sparkJobDefinitionClient) renameSparkJobDefinition(ctx context.Context, sparkJobDefinitionName string, request ArtifactRenameRequest, options *SparkJobDefinitionBeginRenameSparkJobDefinitionOptions) (*azcore.Response, error) {
	req, err := client.renameSparkJobDefinitionCreateRequest(ctx, sparkJobDefinitionName, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.renameSparkJobDefinitionHandleError(resp)
	}
	return resp, nil
}

// renameSparkJobDefinitionCreateRequest creates the RenameSparkJobDefinition request.
func (client *sparkJobDefinitionClient) renameSparkJobDefinitionCreateRequest(ctx context.Context, sparkJobDefinitionName string, request ArtifactRenameRequest, options *SparkJobDefinitionBeginRenameSparkJobDefinitionOptions) (*azcore.Request, error) {
	urlPath := "/sparkJobDefinitions/{sparkJobDefinitionName}/rename"
	urlPath = strings.ReplaceAll(urlPath, "{sparkJobDefinitionName}", url.PathEscape(sparkJobDefinitionName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(request)
}

// renameSparkJobDefinitionHandleError handles the RenameSparkJobDefinition error response.
func (client *sparkJobDefinitionClient) renameSparkJobDefinitionHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
