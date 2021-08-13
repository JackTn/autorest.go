// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package aztables

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// TableClient contains the methods for the Table group.
// Don't use this type directly, use NewTableClient() instead.
type TableClient struct {
	con *Connection
}

// NewTableClient creates a new instance of TableClient with the specified values.
func NewTableClient(con *Connection) *TableClient {
	return &TableClient{con: con}
}

// Create - Creates a new table under the given account.
// If the operation fails it returns the *TableServiceError error type.
func (client *TableClient) Create(ctx context.Context, tableProperties TableProperties, options *TableCreateOptions) (TableCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, tableProperties, options)
	if err != nil {
		return TableCreateResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TableCreateResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusCreated, http.StatusNoContent) {
		return TableCreateResponse{}, client.createHandleError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *TableClient) createCreateRequest(ctx context.Context, tableProperties TableProperties, options *TableCreateOptions) (*azcore.Request, error) {
	urlPath := "/Tables"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if options != nil && options.Format != nil {
		reqQP.Set("$format", string(*options.Format))
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("x-ms-version", "2019-02-02")
	if options != nil && options.RequestID != nil {
		req.Header.Set("x-ms-client-request-id", *options.RequestID)
	}
	req.Header.Set("DataServiceVersion", "3.0")
	if options != nil && options.ResponsePreference != nil {
		req.Header.Set("Prefer", string(*options.ResponsePreference))
	}
	req.Header.Set("Accept", "application/json;odata=minimalmetadata")
	return req, req.MarshalAsJSON(tableProperties)
}

// createHandleResponse handles the Create response.
func (client *TableClient) createHandleResponse(resp *azcore.Response) (TableCreateResponse, error) {
	result := TableCreateResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return TableCreateResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("Preference-Applied"); val != "" {
		result.PreferenceApplied = &val
	}
	if err := resp.UnmarshalAsJSON(&result.TableResponse); err != nil {
		return TableCreateResponse{}, err
	}
	return result, nil
}

// createHandleError handles the Create error response.
func (client *TableClient) createHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := TableServiceError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Delete - Operation permanently deletes the specified table.
// If the operation fails it returns the *TableServiceError error type.
func (client *TableClient) Delete(ctx context.Context, table string, options *TableDeleteOptions) (TableDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, table, options)
	if err != nil {
		return TableDeleteResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TableDeleteResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusNoContent) {
		return TableDeleteResponse{}, client.deleteHandleError(resp)
	}
	return client.deleteHandleResponse(resp)
}

// deleteCreateRequest creates the Delete request.
func (client *TableClient) deleteCreateRequest(ctx context.Context, table string, options *TableDeleteOptions) (*azcore.Request, error) {
	urlPath := "/Tables('{table}')"
	if table == "" {
		return nil, errors.New("parameter table cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{table}", url.PathEscape(table))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("x-ms-version", "2019-02-02")
	if options != nil && options.RequestID != nil {
		req.Header.Set("x-ms-client-request-id", *options.RequestID)
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *TableClient) deleteHandleResponse(resp *azcore.Response) (TableDeleteResponse, error) {
	result := TableDeleteResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return TableDeleteResponse{}, err
		}
		result.Date = &date
	}
	return result, nil
}

// deleteHandleError handles the Delete error response.
func (client *TableClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := TableServiceError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// DeleteEntity - Deletes the specified entity in a table.
// If the operation fails it returns the *TableServiceError error type.
func (client *TableClient) DeleteEntity(ctx context.Context, table string, partitionKey string, rowKey string, ifMatch string, options *TableDeleteEntityOptions) (TableDeleteEntityResponse, error) {
	req, err := client.deleteEntityCreateRequest(ctx, table, partitionKey, rowKey, ifMatch, options)
	if err != nil {
		return TableDeleteEntityResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TableDeleteEntityResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusNoContent) {
		return TableDeleteEntityResponse{}, client.deleteEntityHandleError(resp)
	}
	return client.deleteEntityHandleResponse(resp)
}

// deleteEntityCreateRequest creates the DeleteEntity request.
func (client *TableClient) deleteEntityCreateRequest(ctx context.Context, table string, partitionKey string, rowKey string, ifMatch string, options *TableDeleteEntityOptions) (*azcore.Request, error) {
	urlPath := "/{table}(PartitionKey='{partitionKey}',RowKey='{rowKey}')"
	if table == "" {
		return nil, errors.New("parameter table cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{table}", url.PathEscape(table))
	if partitionKey == "" {
		return nil, errors.New("parameter partitionKey cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partitionKey}", url.PathEscape(partitionKey))
	if rowKey == "" {
		return nil, errors.New("parameter rowKey cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rowKey}", url.PathEscape(rowKey))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	if options != nil && options.Format != nil {
		reqQP.Set("$format", string(*options.Format))
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("x-ms-version", "2019-02-02")
	if options != nil && options.RequestID != nil {
		req.Header.Set("x-ms-client-request-id", *options.RequestID)
	}
	req.Header.Set("DataServiceVersion", "3.0")
	req.Header.Set("If-Match", ifMatch)
	req.Header.Set("Accept", "application/json;odata=minimalmetadata")
	return req, nil
}

// deleteEntityHandleResponse handles the DeleteEntity response.
func (client *TableClient) deleteEntityHandleResponse(resp *azcore.Response) (TableDeleteEntityResponse, error) {
	result := TableDeleteEntityResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return TableDeleteEntityResponse{}, err
		}
		result.Date = &date
	}
	return result, nil
}

// deleteEntityHandleError handles the DeleteEntity error response.
func (client *TableClient) deleteEntityHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := TableServiceError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// GetAccessPolicy - Retrieves details about any stored access policies specified on the table that may be used with Shared Access Signatures.
// If the operation fails it returns the *TableServiceError error type.
func (client *TableClient) GetAccessPolicy(ctx context.Context, table string, options *TableGetAccessPolicyOptions) (TableGetAccessPolicyResponse, error) {
	req, err := client.getAccessPolicyCreateRequest(ctx, table, options)
	if err != nil {
		return TableGetAccessPolicyResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TableGetAccessPolicyResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return TableGetAccessPolicyResponse{}, client.getAccessPolicyHandleError(resp)
	}
	return client.getAccessPolicyHandleResponse(resp)
}

// getAccessPolicyCreateRequest creates the GetAccessPolicy request.
func (client *TableClient) getAccessPolicyCreateRequest(ctx context.Context, table string, options *TableGetAccessPolicyOptions) (*azcore.Request, error) {
	urlPath := "/{table}"
	if table == "" {
		return nil, errors.New("parameter table cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{table}", url.PathEscape(table))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	reqQP.Set("comp", "acl")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("x-ms-version", "2019-02-02")
	if options != nil && options.RequestID != nil {
		req.Header.Set("x-ms-client-request-id", *options.RequestID)
	}
	req.Header.Set("Accept", "application/xml")
	return req, nil
}

// getAccessPolicyHandleResponse handles the GetAccessPolicy response.
func (client *TableClient) getAccessPolicyHandleResponse(resp *azcore.Response) (TableGetAccessPolicyResponse, error) {
	result := TableGetAccessPolicyResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return TableGetAccessPolicyResponse{}, err
		}
		result.Date = &date
	}
	if err := resp.UnmarshalAsXML(&result); err != nil {
		return TableGetAccessPolicyResponse{}, err
	}
	return result, nil
}

// getAccessPolicyHandleError handles the GetAccessPolicy error response.
func (client *TableClient) getAccessPolicyHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := TableServiceError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// InsertEntity - Insert entity in a table.
// If the operation fails it returns the *TableServiceError error type.
func (client *TableClient) InsertEntity(ctx context.Context, table string, options *TableInsertEntityOptions) (TableInsertEntityResponse, error) {
	req, err := client.insertEntityCreateRequest(ctx, table, options)
	if err != nil {
		return TableInsertEntityResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TableInsertEntityResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusCreated, http.StatusNoContent) {
		return TableInsertEntityResponse{}, client.insertEntityHandleError(resp)
	}
	return client.insertEntityHandleResponse(resp)
}

// insertEntityCreateRequest creates the InsertEntity request.
func (client *TableClient) insertEntityCreateRequest(ctx context.Context, table string, options *TableInsertEntityOptions) (*azcore.Request, error) {
	urlPath := "/{table}"
	if table == "" {
		return nil, errors.New("parameter table cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{table}", url.PathEscape(table))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	if options != nil && options.Format != nil {
		reqQP.Set("$format", string(*options.Format))
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("x-ms-version", "2019-02-02")
	if options != nil && options.RequestID != nil {
		req.Header.Set("x-ms-client-request-id", *options.RequestID)
	}
	req.Header.Set("DataServiceVersion", "3.0")
	if options != nil && options.ResponsePreference != nil {
		req.Header.Set("Prefer", string(*options.ResponsePreference))
	}
	req.Header.Set("Accept", "application/json;odata=minimalmetadata")
	if options != nil && options.TableEntityProperties != nil {
		return req, req.MarshalAsJSON(options.TableEntityProperties)
	}
	return req, nil
}

// insertEntityHandleResponse handles the InsertEntity response.
func (client *TableClient) insertEntityHandleResponse(resp *azcore.Response) (TableInsertEntityResponse, error) {
	result := TableInsertEntityResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return TableInsertEntityResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if val := resp.Header.Get("Preference-Applied"); val != "" {
		result.PreferenceApplied = &val
	}
	if val := resp.Header.Get("Content-Type"); val != "" {
		result.ContentType = &val
	}
	if err := resp.UnmarshalAsJSON(&result.Value); err != nil {
		return TableInsertEntityResponse{}, err
	}
	return result, nil
}

// insertEntityHandleError handles the InsertEntity error response.
func (client *TableClient) insertEntityHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := TableServiceError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// MergeEntity - Merge entity in a table.
// If the operation fails it returns the *TableServiceError error type.
func (client *TableClient) MergeEntity(ctx context.Context, table string, partitionKey string, rowKey string, options *TableMergeEntityOptions) (TableMergeEntityResponse, error) {
	req, err := client.mergeEntityCreateRequest(ctx, table, partitionKey, rowKey, options)
	if err != nil {
		return TableMergeEntityResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TableMergeEntityResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusNoContent) {
		return TableMergeEntityResponse{}, client.mergeEntityHandleError(resp)
	}
	return client.mergeEntityHandleResponse(resp)
}

// mergeEntityCreateRequest creates the MergeEntity request.
func (client *TableClient) mergeEntityCreateRequest(ctx context.Context, table string, partitionKey string, rowKey string, options *TableMergeEntityOptions) (*azcore.Request, error) {
	urlPath := "/{table}(PartitionKey='{partitionKey}',RowKey='{rowKey}')"
	if table == "" {
		return nil, errors.New("parameter table cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{table}", url.PathEscape(table))
	if partitionKey == "" {
		return nil, errors.New("parameter partitionKey cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partitionKey}", url.PathEscape(partitionKey))
	if rowKey == "" {
		return nil, errors.New("parameter rowKey cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rowKey}", url.PathEscape(rowKey))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	if options != nil && options.Format != nil {
		reqQP.Set("$format", string(*options.Format))
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("x-ms-version", "2019-02-02")
	if options != nil && options.RequestID != nil {
		req.Header.Set("x-ms-client-request-id", *options.RequestID)
	}
	req.Header.Set("DataServiceVersion", "3.0")
	if options != nil && options.IfMatch != nil {
		req.Header.Set("If-Match", *options.IfMatch)
	}
	req.Header.Set("Accept", "application/json")
	if options != nil && options.TableEntityProperties != nil {
		return req, req.MarshalAsJSON(options.TableEntityProperties)
	}
	return req, nil
}

// mergeEntityHandleResponse handles the MergeEntity response.
func (client *TableClient) mergeEntityHandleResponse(resp *azcore.Response) (TableMergeEntityResponse, error) {
	result := TableMergeEntityResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return TableMergeEntityResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	return result, nil
}

// mergeEntityHandleError handles the MergeEntity error response.
func (client *TableClient) mergeEntityHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := TableServiceError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Query - Queries tables under the given account.
// If the operation fails it returns a generic error.
func (client *TableClient) Query(ctx context.Context, options *TableQueryOptions) (TableQueryResponseEnvelope, error) {
	req, err := client.queryCreateRequest(ctx, options)
	if err != nil {
		return TableQueryResponseEnvelope{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TableQueryResponseEnvelope{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return TableQueryResponseEnvelope{}, client.queryHandleError(resp)
	}
	return client.queryHandleResponse(resp)
}

// queryCreateRequest creates the Query request.
func (client *TableClient) queryCreateRequest(ctx context.Context, options *TableQueryOptions) (*azcore.Request, error) {
	urlPath := "/Tables"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if options != nil && options.Format != nil {
		reqQP.Set("$format", string(*options.Format))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Select != nil {
		reqQP.Set("$select", *options.Select)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.NextTableName != nil {
		reqQP.Set("NextTableName", *options.NextTableName)
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("x-ms-version", "2019-02-02")
	if options != nil && options.RequestID != nil {
		req.Header.Set("x-ms-client-request-id", *options.RequestID)
	}
	req.Header.Set("DataServiceVersion", "3.0")
	req.Header.Set("Accept", "application/json;odata=minimalmetadata")
	return req, nil
}

// queryHandleResponse handles the Query response.
func (client *TableClient) queryHandleResponse(resp *azcore.Response) (TableQueryResponseEnvelope, error) {
	result := TableQueryResponseEnvelope{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return TableQueryResponseEnvelope{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("x-ms-continuation-NextTableName"); val != "" {
		result.XMSContinuationNextTableName = &val
	}
	if err := resp.UnmarshalAsJSON(&result.TableQueryResponse); err != nil {
		return TableQueryResponseEnvelope{}, err
	}
	return result, nil
}

// queryHandleError handles the Query error response.
func (client *TableClient) queryHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// QueryEntities - Queries entities in a table.
// If the operation fails it returns the *TableServiceError error type.
func (client *TableClient) QueryEntities(ctx context.Context, table string, options *TableQueryEntitiesOptions) (TableQueryEntitiesResponse, error) {
	req, err := client.queryEntitiesCreateRequest(ctx, table, options)
	if err != nil {
		return TableQueryEntitiesResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TableQueryEntitiesResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return TableQueryEntitiesResponse{}, client.queryEntitiesHandleError(resp)
	}
	return client.queryEntitiesHandleResponse(resp)
}

// queryEntitiesCreateRequest creates the QueryEntities request.
func (client *TableClient) queryEntitiesCreateRequest(ctx context.Context, table string, options *TableQueryEntitiesOptions) (*azcore.Request, error) {
	urlPath := "/{table}()"
	if table == "" {
		return nil, errors.New("parameter table cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{table}", url.PathEscape(table))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	if options != nil && options.Format != nil {
		reqQP.Set("$format", string(*options.Format))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Select != nil {
		reqQP.Set("$select", *options.Select)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.NextPartitionKey != nil {
		reqQP.Set("NextPartitionKey", *options.NextPartitionKey)
	}
	if options != nil && options.NextRowKey != nil {
		reqQP.Set("NextRowKey", *options.NextRowKey)
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("x-ms-version", "2019-02-02")
	if options != nil && options.RequestID != nil {
		req.Header.Set("x-ms-client-request-id", *options.RequestID)
	}
	req.Header.Set("DataServiceVersion", "3.0")
	req.Header.Set("Accept", "application/json;odata=minimalmetadata")
	return req, nil
}

// queryEntitiesHandleResponse handles the QueryEntities response.
func (client *TableClient) queryEntitiesHandleResponse(resp *azcore.Response) (TableQueryEntitiesResponse, error) {
	result := TableQueryEntitiesResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return TableQueryEntitiesResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("x-ms-continuation-NextPartitionKey"); val != "" {
		result.XMSContinuationNextPartitionKey = &val
	}
	if val := resp.Header.Get("x-ms-continuation-NextRowKey"); val != "" {
		result.XMSContinuationNextRowKey = &val
	}
	if err := resp.UnmarshalAsJSON(&result.TableEntityQueryResponse); err != nil {
		return TableQueryEntitiesResponse{}, err
	}
	return result, nil
}

// queryEntitiesHandleError handles the QueryEntities error response.
func (client *TableClient) queryEntitiesHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := TableServiceError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// QueryEntityWithPartitionAndRowKey - Queries a single entity in a table.
// If the operation fails it returns the *TableServiceError error type.
func (client *TableClient) QueryEntityWithPartitionAndRowKey(ctx context.Context, table string, partitionKey string, rowKey string, options *TableQueryEntityWithPartitionAndRowKeyOptions) (TableQueryEntityWithPartitionAndRowKeyResponse, error) {
	req, err := client.queryEntityWithPartitionAndRowKeyCreateRequest(ctx, table, partitionKey, rowKey, options)
	if err != nil {
		return TableQueryEntityWithPartitionAndRowKeyResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TableQueryEntityWithPartitionAndRowKeyResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return TableQueryEntityWithPartitionAndRowKeyResponse{}, client.queryEntityWithPartitionAndRowKeyHandleError(resp)
	}
	return client.queryEntityWithPartitionAndRowKeyHandleResponse(resp)
}

// queryEntityWithPartitionAndRowKeyCreateRequest creates the QueryEntityWithPartitionAndRowKey request.
func (client *TableClient) queryEntityWithPartitionAndRowKeyCreateRequest(ctx context.Context, table string, partitionKey string, rowKey string, options *TableQueryEntityWithPartitionAndRowKeyOptions) (*azcore.Request, error) {
	urlPath := "/{table}(PartitionKey='{partitionKey}',RowKey='{rowKey}')"
	if table == "" {
		return nil, errors.New("parameter table cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{table}", url.PathEscape(table))
	if partitionKey == "" {
		return nil, errors.New("parameter partitionKey cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partitionKey}", url.PathEscape(partitionKey))
	if rowKey == "" {
		return nil, errors.New("parameter rowKey cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rowKey}", url.PathEscape(rowKey))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	if options != nil && options.Format != nil {
		reqQP.Set("$format", string(*options.Format))
	}
	if options != nil && options.Select != nil {
		reqQP.Set("$select", *options.Select)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("x-ms-version", "2019-02-02")
	if options != nil && options.RequestID != nil {
		req.Header.Set("x-ms-client-request-id", *options.RequestID)
	}
	req.Header.Set("DataServiceVersion", "3.0")
	req.Header.Set("Accept", "application/json;odata=minimalmetadata")
	return req, nil
}

// queryEntityWithPartitionAndRowKeyHandleResponse handles the QueryEntityWithPartitionAndRowKey response.
func (client *TableClient) queryEntityWithPartitionAndRowKeyHandleResponse(resp *azcore.Response) (TableQueryEntityWithPartitionAndRowKeyResponse, error) {
	result := TableQueryEntityWithPartitionAndRowKeyResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return TableQueryEntityWithPartitionAndRowKeyResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if val := resp.Header.Get("x-ms-continuation-NextPartitionKey"); val != "" {
		result.XMSContinuationNextPartitionKey = &val
	}
	if val := resp.Header.Get("x-ms-continuation-NextRowKey"); val != "" {
		result.XMSContinuationNextRowKey = &val
	}
	if err := resp.UnmarshalAsJSON(&result.Value); err != nil {
		return TableQueryEntityWithPartitionAndRowKeyResponse{}, err
	}
	return result, nil
}

// queryEntityWithPartitionAndRowKeyHandleError handles the QueryEntityWithPartitionAndRowKey error response.
func (client *TableClient) queryEntityWithPartitionAndRowKeyHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := TableServiceError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// SetAccessPolicy - Sets stored access policies for the table that may be used with Shared Access Signatures.
// If the operation fails it returns the *TableServiceError error type.
func (client *TableClient) SetAccessPolicy(ctx context.Context, table string, options *TableSetAccessPolicyOptions) (TableSetAccessPolicyResponse, error) {
	req, err := client.setAccessPolicyCreateRequest(ctx, table, options)
	if err != nil {
		return TableSetAccessPolicyResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TableSetAccessPolicyResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusNoContent) {
		return TableSetAccessPolicyResponse{}, client.setAccessPolicyHandleError(resp)
	}
	return client.setAccessPolicyHandleResponse(resp)
}

// setAccessPolicyCreateRequest creates the SetAccessPolicy request.
func (client *TableClient) setAccessPolicyCreateRequest(ctx context.Context, table string, options *TableSetAccessPolicyOptions) (*azcore.Request, error) {
	urlPath := "/{table}"
	if table == "" {
		return nil, errors.New("parameter table cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{table}", url.PathEscape(table))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	reqQP.Set("comp", "acl")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("x-ms-version", "2019-02-02")
	if options != nil && options.RequestID != nil {
		req.Header.Set("x-ms-client-request-id", *options.RequestID)
	}
	req.Header.Set("Accept", "application/xml")
	type wrapper struct {
		XMLName  xml.Name             `xml:"SignedIdentifiers"`
		TableACL *[]*SignedIdentifier `xml:"SignedIdentifier"`
	}
	if options != nil && options.TableACL != nil {
		return req, req.MarshalAsXML(wrapper{TableACL: &options.TableACL})
	}
	return req, nil
}

// setAccessPolicyHandleResponse handles the SetAccessPolicy response.
func (client *TableClient) setAccessPolicyHandleResponse(resp *azcore.Response) (TableSetAccessPolicyResponse, error) {
	result := TableSetAccessPolicyResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return TableSetAccessPolicyResponse{}, err
		}
		result.Date = &date
	}
	return result, nil
}

// setAccessPolicyHandleError handles the SetAccessPolicy error response.
func (client *TableClient) setAccessPolicyHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := TableServiceError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// UpdateEntity - Update entity in a table.
// If the operation fails it returns the *TableServiceError error type.
func (client *TableClient) UpdateEntity(ctx context.Context, table string, partitionKey string, rowKey string, options *TableUpdateEntityOptions) (TableUpdateEntityResponse, error) {
	req, err := client.updateEntityCreateRequest(ctx, table, partitionKey, rowKey, options)
	if err != nil {
		return TableUpdateEntityResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TableUpdateEntityResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusNoContent) {
		return TableUpdateEntityResponse{}, client.updateEntityHandleError(resp)
	}
	return client.updateEntityHandleResponse(resp)
}

// updateEntityCreateRequest creates the UpdateEntity request.
func (client *TableClient) updateEntityCreateRequest(ctx context.Context, table string, partitionKey string, rowKey string, options *TableUpdateEntityOptions) (*azcore.Request, error) {
	urlPath := "/{table}(PartitionKey='{partitionKey}',RowKey='{rowKey}')"
	if table == "" {
		return nil, errors.New("parameter table cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{table}", url.PathEscape(table))
	if partitionKey == "" {
		return nil, errors.New("parameter partitionKey cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partitionKey}", url.PathEscape(partitionKey))
	if rowKey == "" {
		return nil, errors.New("parameter rowKey cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rowKey}", url.PathEscape(rowKey))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	if options != nil && options.Format != nil {
		reqQP.Set("$format", string(*options.Format))
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("x-ms-version", "2019-02-02")
	if options != nil && options.RequestID != nil {
		req.Header.Set("x-ms-client-request-id", *options.RequestID)
	}
	req.Header.Set("DataServiceVersion", "3.0")
	if options != nil && options.IfMatch != nil {
		req.Header.Set("If-Match", *options.IfMatch)
	}
	req.Header.Set("Accept", "application/json")
	if options != nil && options.TableEntityProperties != nil {
		return req, req.MarshalAsJSON(options.TableEntityProperties)
	}
	return req, nil
}

// updateEntityHandleResponse handles the UpdateEntity response.
func (client *TableClient) updateEntityHandleResponse(resp *azcore.Response) (TableUpdateEntityResponse, error) {
	result := TableUpdateEntityResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return TableUpdateEntityResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	return result, nil
}

// updateEntityHandleError handles the UpdateEntity error response.
func (client *TableClient) updateEntityHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := TableServiceError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
