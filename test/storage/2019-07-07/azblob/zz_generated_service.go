// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azblob

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"strconv"
	"time"
)

type serviceClient struct {
	con *connection
}

// Pipeline returns the pipeline associated with this client.
func (client *serviceClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// GetAccountInfo - Returns the sku name and account kind
func (client *serviceClient) GetAccountInfo(ctx context.Context, options *ServiceGetAccountInfoOptions) (*ServiceGetAccountInfoResponse, error) {
	req, err := client.GetAccountInfoCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetAccountInfoHandleError(resp)
	}
	result, err := client.GetAccountInfoHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetAccountInfoCreateRequest creates the GetAccountInfo request.
func (client *serviceClient) GetAccountInfoCreateRequest(ctx context.Context, options *ServiceGetAccountInfoOptions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodGet, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("restype", "account")
	query.Set("comp", "properties")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("x-ms-version", "2019-07-07")
	req.Header.Set("Accept", "application/xml")
	return req, nil
}

// GetAccountInfoHandleResponse handles the GetAccountInfo response.
func (client *serviceClient) GetAccountInfoHandleResponse(resp *azcore.Response) (*ServiceGetAccountInfoResponse, error) {
	result := ServiceGetAccountInfoResponse{RawResponse: resp.Response}
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
			return nil, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("x-ms-sku-name"); val != "" {
		result.SKUName = (*SKUName)(&val)
	}
	if val := resp.Header.Get("x-ms-account-kind"); val != "" {
		result.AccountKind = (*AccountKind)(&val)
	}
	return &result, nil
}

// GetAccountInfoHandleError handles the GetAccountInfo error response.
func (client *serviceClient) GetAccountInfoHandleError(resp *azcore.Response) error {
	var err StorageError
	if err := resp.UnmarshalAsXML(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetProperties - gets the properties of a storage account's Blob service, including properties for Storage Analytics and CORS (Cross-Origin Resource Sharing)
// rules.
func (client *serviceClient) GetProperties(ctx context.Context, options *ServiceGetPropertiesOptions) (*StorageServicePropertiesResponse, error) {
	req, err := client.GetPropertiesCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetPropertiesHandleError(resp)
	}
	result, err := client.GetPropertiesHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetPropertiesCreateRequest creates the GetProperties request.
func (client *serviceClient) GetPropertiesCreateRequest(ctx context.Context, options *ServiceGetPropertiesOptions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodGet, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("restype", "service")
	query.Set("comp", "properties")
	if options != nil && options.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("x-ms-version", "2019-07-07")
	if options != nil && options.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *options.RequestId)
	}
	req.Header.Set("Accept", "application/xml")
	return req, nil
}

// GetPropertiesHandleResponse handles the GetProperties response.
func (client *serviceClient) GetPropertiesHandleResponse(resp *azcore.Response) (*StorageServicePropertiesResponse, error) {
	result := StorageServicePropertiesResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	return &result, resp.UnmarshalAsXML(&result.StorageServiceProperties)
}

// GetPropertiesHandleError handles the GetProperties error response.
func (client *serviceClient) GetPropertiesHandleError(resp *azcore.Response) error {
	var err StorageError
	if err := resp.UnmarshalAsXML(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetStatistics - Retrieves statistics related to replication for the Blob service. It is only available on the secondary location endpoint when read-access
// geo-redundant replication is enabled for the storage account.
func (client *serviceClient) GetStatistics(ctx context.Context, options *ServiceGetStatisticsOptions) (*StorageServiceStatsResponse, error) {
	req, err := client.GetStatisticsCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetStatisticsHandleError(resp)
	}
	result, err := client.GetStatisticsHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetStatisticsCreateRequest creates the GetStatistics request.
func (client *serviceClient) GetStatisticsCreateRequest(ctx context.Context, options *ServiceGetStatisticsOptions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodGet, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("restype", "service")
	query.Set("comp", "stats")
	if options != nil && options.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("x-ms-version", "2019-07-07")
	if options != nil && options.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *options.RequestId)
	}
	req.Header.Set("Accept", "application/xml")
	return req, nil
}

// GetStatisticsHandleResponse handles the GetStatistics response.
func (client *serviceClient) GetStatisticsHandleResponse(resp *azcore.Response) (*StorageServiceStatsResponse, error) {
	result := StorageServiceStatsResponse{RawResponse: resp.Response}
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
			return nil, err
		}
		result.Date = &date
	}
	return &result, resp.UnmarshalAsXML(&result.StorageServiceStats)
}

// GetStatisticsHandleError handles the GetStatistics error response.
func (client *serviceClient) GetStatisticsHandleError(resp *azcore.Response) error {
	var err StorageError
	if err := resp.UnmarshalAsXML(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetUserDelegationKey - Retrieves a user delegation key for the Blob service. This is only a valid operation when using bearer token authentication.
func (client *serviceClient) GetUserDelegationKey(ctx context.Context, keyInfo KeyInfo, options *ServiceGetUserDelegationKeyOptions) (*UserDelegationKeyResponse, error) {
	req, err := client.GetUserDelegationKeyCreateRequest(ctx, keyInfo, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetUserDelegationKeyHandleError(resp)
	}
	result, err := client.GetUserDelegationKeyHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUserDelegationKeyCreateRequest creates the GetUserDelegationKey request.
func (client *serviceClient) GetUserDelegationKeyCreateRequest(ctx context.Context, keyInfo KeyInfo, options *ServiceGetUserDelegationKeyOptions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodPost, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("restype", "service")
	query.Set("comp", "userdelegationkey")
	if options != nil && options.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("x-ms-version", "2019-07-07")
	if options != nil && options.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *options.RequestId)
	}
	req.Header.Set("Accept", "application/xml")
	return req, req.MarshalAsXML(keyInfo)
}

// GetUserDelegationKeyHandleResponse handles the GetUserDelegationKey response.
func (client *serviceClient) GetUserDelegationKeyHandleResponse(resp *azcore.Response) (*UserDelegationKeyResponse, error) {
	result := UserDelegationKeyResponse{RawResponse: resp.Response}
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
			return nil, err
		}
		result.Date = &date
	}
	return &result, resp.UnmarshalAsXML(&result.UserDelegationKey)
}

// GetUserDelegationKeyHandleError handles the GetUserDelegationKey error response.
func (client *serviceClient) GetUserDelegationKeyHandleError(resp *azcore.Response) error {
	var err StorageError
	if err := resp.UnmarshalAsXML(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListContainersSegment - The List Containers Segment operation returns a list of the containers under the specified account
func (client *serviceClient) ListContainersSegment(options *ServiceListContainersSegmentOptions) ListContainersSegmentResponsePager {
	return &listContainersSegmentResponsePager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListContainersSegmentCreateRequest(ctx, options)
		},
		responder: client.ListContainersSegmentHandleResponse,
		errorer:   client.ListContainersSegmentHandleError,
		advancer: func(ctx context.Context, resp *ListContainersSegmentResponseResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.EnumerationResults.NextMarker)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListContainersSegmentCreateRequest creates the ListContainersSegment request.
func (client *serviceClient) ListContainersSegmentCreateRequest(ctx context.Context, options *ServiceListContainersSegmentOptions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodGet, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("comp", "list")
	if options != nil && options.Prefix != nil {
		query.Set("prefix", *options.Prefix)
	}
	if options != nil && options.Marker != nil {
		query.Set("marker", *options.Marker)
	}
	if options != nil && options.Maxresults != nil {
		query.Set("maxresults", strconv.FormatInt(int64(*options.Maxresults), 10))
	}
	if options != nil && options.Include != nil {
		query.Set("include", "metadata")
	}
	if options != nil && options.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("x-ms-version", "2019-07-07")
	if options != nil && options.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *options.RequestId)
	}
	req.Header.Set("Accept", "application/xml")
	return req, nil
}

// ListContainersSegmentHandleResponse handles the ListContainersSegment response.
func (client *serviceClient) ListContainersSegmentHandleResponse(resp *azcore.Response) (*ListContainersSegmentResponseResponse, error) {
	result := ListContainersSegmentResponseResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	return &result, resp.UnmarshalAsXML(&result.EnumerationResults)
}

// ListContainersSegmentHandleError handles the ListContainersSegment error response.
func (client *serviceClient) ListContainersSegmentHandleError(resp *azcore.Response) error {
	var err StorageError
	if err := resp.UnmarshalAsXML(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// SetProperties - Sets properties for a storage account's Blob service endpoint, including properties for Storage Analytics and CORS (Cross-Origin Resource
// Sharing) rules
func (client *serviceClient) SetProperties(ctx context.Context, storageServiceProperties StorageServiceProperties, options *ServiceSetPropertiesOptions) (*ServiceSetPropertiesResponse, error) {
	req, err := client.SetPropertiesCreateRequest(ctx, storageServiceProperties, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusAccepted) {
		return nil, client.SetPropertiesHandleError(resp)
	}
	result, err := client.SetPropertiesHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SetPropertiesCreateRequest creates the SetProperties request.
func (client *serviceClient) SetPropertiesCreateRequest(ctx context.Context, storageServiceProperties StorageServiceProperties, options *ServiceSetPropertiesOptions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodPut, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("restype", "service")
	query.Set("comp", "properties")
	if options != nil && options.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("x-ms-version", "2019-07-07")
	if options != nil && options.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *options.RequestId)
	}
	req.Header.Set("Accept", "application/xml")
	return req, req.MarshalAsXML(storageServiceProperties)
}

// SetPropertiesHandleResponse handles the SetProperties response.
func (client *serviceClient) SetPropertiesHandleResponse(resp *azcore.Response) (*ServiceSetPropertiesResponse, error) {
	result := ServiceSetPropertiesResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	return &result, nil
}

// SetPropertiesHandleError handles the SetProperties error response.
func (client *serviceClient) SetPropertiesHandleError(resp *azcore.Response) error {
	var err StorageError
	if err := resp.UnmarshalAsXML(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// SubmitBatch - The Batch operation allows multiple API calls to be embedded into a single HTTP request.
func (client *serviceClient) SubmitBatch(ctx context.Context, contentLength int64, multipartContentType string, body azcore.ReadSeekCloser, options *ServiceSubmitBatchOptions) (*ServiceSubmitBatchResponse, error) {
	req, err := client.SubmitBatchCreateRequest(ctx, contentLength, multipartContentType, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.SubmitBatchHandleError(resp)
	}
	result, err := client.SubmitBatchHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SubmitBatchCreateRequest creates the SubmitBatch request.
func (client *serviceClient) SubmitBatchCreateRequest(ctx context.Context, contentLength int64, multipartContentType string, body azcore.ReadSeekCloser, options *ServiceSubmitBatchOptions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodPost, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("comp", "batch")
	if options != nil && options.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.URL.RawQuery = query.Encode()
	req.SkipBodyDownload()
	req.Header.Set("Content-Length", strconv.FormatInt(contentLength, 10))
	req.Header.Set("Content-Type", multipartContentType)
	req.Header.Set("x-ms-version", "2019-07-07")
	if options != nil && options.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *options.RequestId)
	}
	req.Header.Set("Accept", "application/xml")
	return req, req.MarshalAsXML(body)
}

// SubmitBatchHandleResponse handles the SubmitBatch response.
func (client *serviceClient) SubmitBatchHandleResponse(resp *azcore.Response) (*ServiceSubmitBatchResponse, error) {
	result := ServiceSubmitBatchResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("Content-Type"); val != "" {
		result.ContentType = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	return &result, nil
}

// SubmitBatchHandleError handles the SubmitBatch error response.
func (client *serviceClient) SubmitBatchHandleError(resp *azcore.Response) error {
	var err StorageError
	if err := resp.UnmarshalAsXML(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
