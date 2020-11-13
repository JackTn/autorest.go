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

type sqlScriptClient struct {
	con *connection
}

// Pipeline returns the pipeline associated with this client.
func (client *sqlScriptClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// CreateOrUpdateSQLScript - Creates or updates a Sql Script.
func (client *sqlScriptClient) CreateOrUpdateSQLScript(ctx context.Context, sqlScriptName string, sqlScript SQLScriptResource, options *SQLScriptCreateOrUpdateSQLScriptOptions) (*SQLScriptResourceResponse, error) {
	req, err := client.CreateOrUpdateSQLScriptCreateRequest(ctx, sqlScriptName, sqlScript, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.CreateOrUpdateSQLScriptHandleError(resp)
	}
	result, err := client.CreateOrUpdateSQLScriptHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateOrUpdateSQLScriptCreateRequest creates the CreateOrUpdateSQLScript request.
func (client *sqlScriptClient) CreateOrUpdateSQLScriptCreateRequest(ctx context.Context, sqlScriptName string, sqlScript SQLScriptResource, options *SQLScriptCreateOrUpdateSQLScriptOptions) (*azcore.Request, error) {
	urlPath := "/sqlScripts/{sqlScriptName}"
	urlPath = strings.ReplaceAll(urlPath, "{sqlScriptName}", url.PathEscape(sqlScriptName))
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
	return req, req.MarshalAsJSON(sqlScript)
}

// CreateOrUpdateSQLScriptHandleResponse handles the CreateOrUpdateSQLScript response.
func (client *sqlScriptClient) CreateOrUpdateSQLScriptHandleResponse(resp *azcore.Response) (*SQLScriptResourceResponse, error) {
	result := SQLScriptResourceResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.SQLScriptResource)
}

// CreateOrUpdateSQLScriptHandleError handles the CreateOrUpdateSQLScript error response.
func (client *sqlScriptClient) CreateOrUpdateSQLScriptHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// DeleteSQLScript - Deletes a Sql Script.
func (client *sqlScriptClient) DeleteSQLScript(ctx context.Context, sqlScriptName string, options *SQLScriptDeleteSQLScriptOptions) (*http.Response, error) {
	req, err := client.DeleteSQLScriptCreateRequest(ctx, sqlScriptName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return nil, client.DeleteSQLScriptHandleError(resp)
	}
	return resp.Response, nil
}

// DeleteSQLScriptCreateRequest creates the DeleteSQLScript request.
func (client *sqlScriptClient) DeleteSQLScriptCreateRequest(ctx context.Context, sqlScriptName string, options *SQLScriptDeleteSQLScriptOptions) (*azcore.Request, error) {
	urlPath := "/sqlScripts/{sqlScriptName}"
	urlPath = strings.ReplaceAll(urlPath, "{sqlScriptName}", url.PathEscape(sqlScriptName))
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

// DeleteSQLScriptHandleError handles the DeleteSQLScript error response.
func (client *sqlScriptClient) DeleteSQLScriptHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetSQLScript - Gets a sql script.
func (client *sqlScriptClient) GetSQLScript(ctx context.Context, sqlScriptName string, options *SQLScriptGetSQLScriptOptions) (*SQLScriptResourceResponse, error) {
	req, err := client.GetSQLScriptCreateRequest(ctx, sqlScriptName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNotModified) {
		return nil, client.GetSQLScriptHandleError(resp)
	}
	result, err := client.GetSQLScriptHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetSQLScriptCreateRequest creates the GetSQLScript request.
func (client *sqlScriptClient) GetSQLScriptCreateRequest(ctx context.Context, sqlScriptName string, options *SQLScriptGetSQLScriptOptions) (*azcore.Request, error) {
	urlPath := "/sqlScripts/{sqlScriptName}"
	urlPath = strings.ReplaceAll(urlPath, "{sqlScriptName}", url.PathEscape(sqlScriptName))
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

// GetSQLScriptHandleResponse handles the GetSQLScript response.
func (client *sqlScriptClient) GetSQLScriptHandleResponse(resp *azcore.Response) (*SQLScriptResourceResponse, error) {
	result := SQLScriptResourceResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.SQLScriptResource)
}

// GetSQLScriptHandleError handles the GetSQLScript error response.
func (client *sqlScriptClient) GetSQLScriptHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetSQLScriptsByWorkspace - Lists sql scripts.
func (client *sqlScriptClient) GetSQLScriptsByWorkspace(options *SQLScriptGetSQLScriptsByWorkspaceOptions) SQLScriptsListResponsePager {
	return &sqlScriptsListResponsePager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.GetSQLScriptsByWorkspaceCreateRequest(ctx, options)
		},
		responder: client.GetSQLScriptsByWorkspaceHandleResponse,
		errorer:   client.GetSQLScriptsByWorkspaceHandleError,
		advancer: func(ctx context.Context, resp *SQLScriptsListResponseResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.SQLScriptsListResponse.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// GetSQLScriptsByWorkspaceCreateRequest creates the GetSQLScriptsByWorkspace request.
func (client *sqlScriptClient) GetSQLScriptsByWorkspaceCreateRequest(ctx context.Context, options *SQLScriptGetSQLScriptsByWorkspaceOptions) (*azcore.Request, error) {
	urlPath := "/sqlScripts"
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

// GetSQLScriptsByWorkspaceHandleResponse handles the GetSQLScriptsByWorkspace response.
func (client *sqlScriptClient) GetSQLScriptsByWorkspaceHandleResponse(resp *azcore.Response) (*SQLScriptsListResponseResponse, error) {
	result := SQLScriptsListResponseResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.SQLScriptsListResponse)
}

// GetSQLScriptsByWorkspaceHandleError handles the GetSQLScriptsByWorkspace error response.
func (client *sqlScriptClient) GetSQLScriptsByWorkspaceHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// RenameSQLScript - Renames a sqlScript.
func (client *sqlScriptClient) RenameSQLScript(ctx context.Context, sqlScriptName string, request ArtifactRenameRequest, options *SQLScriptRenameSQLScriptOptions) (*azcore.Response, error) {
	req, err := client.RenameSQLScriptCreateRequest(ctx, sqlScriptName, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.RenameSQLScriptHandleError(resp)
	}
	return resp, nil
}

// RenameSQLScriptCreateRequest creates the RenameSQLScript request.
func (client *sqlScriptClient) RenameSQLScriptCreateRequest(ctx context.Context, sqlScriptName string, request ArtifactRenameRequest, options *SQLScriptRenameSQLScriptOptions) (*azcore.Request, error) {
	urlPath := "/sqlScripts/{sqlScriptName}/rename"
	urlPath = strings.ReplaceAll(urlPath, "{sqlScriptName}", url.PathEscape(sqlScriptName))
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

// RenameSQLScriptHandleError handles the RenameSQLScript error response.
func (client *sqlScriptClient) RenameSQLScriptHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
