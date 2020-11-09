// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package paramgroupinggroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ParameterGroupingOperations contains the methods for the ParameterGrouping group.
type ParameterGroupingOperations interface {
	// PostMultiParamGroups - Post parameters from multiple different parameter groups
	PostMultiParamGroups(ctx context.Context, options *FirstParameterGroup, parameterGroupingPostMultiParamGroupsSecondParamGroup *ParameterGroupingPostMultiParamGroupsSecondParamGroup) (*http.Response, error)
	// PostOptional - Post a bunch of optional parameters grouped
	PostOptional(ctx context.Context, options *ParameterGroupingPostOptionalParameters) (*http.Response, error)
	// PostRequired - Post a bunch of required parameters grouped
	PostRequired(ctx context.Context, parameterGroupingPostRequiredParameters ParameterGroupingPostRequiredParameters) (*http.Response, error)
	// PostSharedParameterGroupObject - Post parameters with a shared parameter group object
	PostSharedParameterGroupObject(ctx context.Context, options *FirstParameterGroup) (*http.Response, error)
}

// ParameterGroupingClient implements the ParameterGroupingOperations interface.
// Don't use this type directly, use NewParameterGroupingClient() instead.
type ParameterGroupingClient struct {
	con *Connection
}

// NewParameterGroupingClient creates a new instance of ParameterGroupingClient with the specified values.
func NewParameterGroupingClient(con *Connection) ParameterGroupingOperations {
	return &ParameterGroupingClient{con: con}
}

// Pipeline returns the pipeline associated with this client.
func (client *ParameterGroupingClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// PostMultiParamGroups - Post parameters from multiple different parameter groups
func (client *ParameterGroupingClient) PostMultiParamGroups(ctx context.Context, firstParameterGroup *FirstParameterGroup, parameterGroupingPostMultiParamGroupsSecondParamGroup *ParameterGroupingPostMultiParamGroupsSecondParamGroup) (*http.Response, error) {
	req, err := client.PostMultiParamGroupsCreateRequest(ctx, firstParameterGroup, parameterGroupingPostMultiParamGroupsSecondParamGroup)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PostMultiParamGroupsHandleError(resp)
	}
	return resp.Response, nil
}

// PostMultiParamGroupsCreateRequest creates the PostMultiParamGroups request.
func (client *ParameterGroupingClient) PostMultiParamGroupsCreateRequest(ctx context.Context, firstParameterGroup *FirstParameterGroup, parameterGroupingPostMultiParamGroupsSecondParamGroup *ParameterGroupingPostMultiParamGroupsSecondParamGroup) (*azcore.Request, error) {
	urlPath := "/parameterGrouping/postMultipleParameterGroups"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	if firstParameterGroup != nil && firstParameterGroup.QueryOne != nil {
		query.Set("query-one", strconv.FormatInt(int64(*firstParameterGroup.QueryOne), 10))
	}
	if parameterGroupingPostMultiParamGroupsSecondParamGroup != nil && parameterGroupingPostMultiParamGroupsSecondParamGroup.QueryTwo != nil {
		query.Set("query-two", strconv.FormatInt(int64(*parameterGroupingPostMultiParamGroupsSecondParamGroup.QueryTwo), 10))
	}
	req.URL.RawQuery = query.Encode()
	if firstParameterGroup != nil && firstParameterGroup.HeaderOne != nil {
		req.Header.Set("header-one", *firstParameterGroup.HeaderOne)
	}
	if parameterGroupingPostMultiParamGroupsSecondParamGroup != nil && parameterGroupingPostMultiParamGroupsSecondParamGroup.HeaderTwo != nil {
		req.Header.Set("header-two", *parameterGroupingPostMultiParamGroupsSecondParamGroup.HeaderTwo)
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// PostMultiParamGroupsHandleError handles the PostMultiParamGroups error response.
func (client *ParameterGroupingClient) PostMultiParamGroupsHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PostOptional - Post a bunch of optional parameters grouped
func (client *ParameterGroupingClient) PostOptional(ctx context.Context, options *ParameterGroupingPostOptionalParameters) (*http.Response, error) {
	req, err := client.PostOptionalCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PostOptionalHandleError(resp)
	}
	return resp.Response, nil
}

// PostOptionalCreateRequest creates the PostOptional request.
func (client *ParameterGroupingClient) PostOptionalCreateRequest(ctx context.Context, options *ParameterGroupingPostOptionalParameters) (*azcore.Request, error) {
	urlPath := "/parameterGrouping/postOptional"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	if options != nil && options.Query != nil {
		query.Set("query", strconv.FormatInt(int64(*options.Query), 10))
	}
	req.URL.RawQuery = query.Encode()
	if options != nil && options.CustomHeader != nil {
		req.Header.Set("customHeader", *options.CustomHeader)
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// PostOptionalHandleError handles the PostOptional error response.
func (client *ParameterGroupingClient) PostOptionalHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PostRequired - Post a bunch of required parameters grouped
func (client *ParameterGroupingClient) PostRequired(ctx context.Context, parameterGroupingPostRequiredParameters ParameterGroupingPostRequiredParameters) (*http.Response, error) {
	req, err := client.PostRequiredCreateRequest(ctx, parameterGroupingPostRequiredParameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PostRequiredHandleError(resp)
	}
	return resp.Response, nil
}

// PostRequiredCreateRequest creates the PostRequired request.
func (client *ParameterGroupingClient) PostRequiredCreateRequest(ctx context.Context, parameterGroupingPostRequiredParameters ParameterGroupingPostRequiredParameters) (*azcore.Request, error) {
	urlPath := "/parameterGrouping/postRequired/{path}"
	urlPath = strings.ReplaceAll(urlPath, "{path}", url.PathEscape(parameterGroupingPostRequiredParameters.PathParameter))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	if parameterGroupingPostRequiredParameters.Query != nil {
		query.Set("query", strconv.FormatInt(int64(*parameterGroupingPostRequiredParameters.Query), 10))
	}
	req.URL.RawQuery = query.Encode()
	if parameterGroupingPostRequiredParameters.CustomHeader != nil {
		req.Header.Set("customHeader", *parameterGroupingPostRequiredParameters.CustomHeader)
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameterGroupingPostRequiredParameters.Body)
}

// PostRequiredHandleError handles the PostRequired error response.
func (client *ParameterGroupingClient) PostRequiredHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PostSharedParameterGroupObject - Post parameters with a shared parameter group object
func (client *ParameterGroupingClient) PostSharedParameterGroupObject(ctx context.Context, options *FirstParameterGroup) (*http.Response, error) {
	req, err := client.PostSharedParameterGroupObjectCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PostSharedParameterGroupObjectHandleError(resp)
	}
	return resp.Response, nil
}

// PostSharedParameterGroupObjectCreateRequest creates the PostSharedParameterGroupObject request.
func (client *ParameterGroupingClient) PostSharedParameterGroupObjectCreateRequest(ctx context.Context, options *FirstParameterGroup) (*azcore.Request, error) {
	urlPath := "/parameterGrouping/sharedParameterGroupObject"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	if options != nil && options.QueryOne != nil {
		query.Set("query-one", strconv.FormatInt(int64(*options.QueryOne), 10))
	}
	req.URL.RawQuery = query.Encode()
	if options != nil && options.HeaderOne != nil {
		req.Header.Set("header-one", *options.HeaderOne)
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// PostSharedParameterGroupObjectHandleError handles the PostSharedParameterGroupObject error response.
func (client *ParameterGroupingClient) PostSharedParameterGroupObjectHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
