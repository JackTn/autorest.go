//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

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

// LinkConnectionClient contains the methods for the LinkConnection group.
// Don't use this type directly, use a constructor function instead.
type LinkConnectionClient struct {
	internal *azcore.Client
	endpoint string
}

// CreateOrUpdate - Creates or updates a link connection
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-12-01-preview
//   - linkConnectionName - The link connection name
//   - linkConnection - Link connection resource definition
//   - options - LinkConnectionClientCreateOrUpdateOptions contains the optional parameters for the LinkConnectionClient.CreateOrUpdate
//     method.
func (client *LinkConnectionClient) CreateOrUpdate(ctx context.Context, linkConnectionName string, linkConnection LinkConnectionResource, options *LinkConnectionClientCreateOrUpdateOptions) (LinkConnectionClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, linkConnectionName, linkConnection, options)
	if err != nil {
		return LinkConnectionClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LinkConnectionClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LinkConnectionClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *LinkConnectionClient) createOrUpdateCreateRequest(ctx context.Context, linkConnectionName string, linkConnection LinkConnectionResource, options *LinkConnectionClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/linkconnections/{linkConnectionName}"
	if linkConnectionName == "" {
		return nil, errors.New("parameter linkConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkConnectionName}", url.PathEscape(linkConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, linkConnection)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *LinkConnectionClient) createOrUpdateHandleResponse(resp *http.Response) (LinkConnectionClientCreateOrUpdateResponse, error) {
	result := LinkConnectionClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LinkConnectionResource); err != nil {
		return LinkConnectionClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a link connection
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-12-01-preview
//   - linkConnectionName - The link connection name
//   - options - LinkConnectionClientDeleteOptions contains the optional parameters for the LinkConnectionClient.Delete method.
func (client *LinkConnectionClient) Delete(ctx context.Context, linkConnectionName string, options *LinkConnectionClientDeleteOptions) (LinkConnectionClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, linkConnectionName, options)
	if err != nil {
		return LinkConnectionClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LinkConnectionClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return LinkConnectionClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return LinkConnectionClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *LinkConnectionClient) deleteCreateRequest(ctx context.Context, linkConnectionName string, options *LinkConnectionClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/linkconnections/{linkConnectionName}"
	if linkConnectionName == "" {
		return nil, errors.New("parameter linkConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkConnectionName}", url.PathEscape(linkConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// EditTables - Edit tables for a link connection
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-12-01-preview
//   - linkConnectionName - The link connection name
//   - editTablesRequest - Edit tables request
//   - options - LinkConnectionClientEditTablesOptions contains the optional parameters for the LinkConnectionClient.EditTables
//     method.
func (client *LinkConnectionClient) EditTables(ctx context.Context, linkConnectionName string, editTablesRequest EditTablesRequest, options *LinkConnectionClientEditTablesOptions) (LinkConnectionClientEditTablesResponse, error) {
	req, err := client.editTablesCreateRequest(ctx, linkConnectionName, editTablesRequest, options)
	if err != nil {
		return LinkConnectionClientEditTablesResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LinkConnectionClientEditTablesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LinkConnectionClientEditTablesResponse{}, runtime.NewResponseError(resp)
	}
	return LinkConnectionClientEditTablesResponse{}, nil
}

// editTablesCreateRequest creates the EditTables request.
func (client *LinkConnectionClient) editTablesCreateRequest(ctx context.Context, linkConnectionName string, editTablesRequest EditTablesRequest, options *LinkConnectionClientEditTablesOptions) (*policy.Request, error) {
	urlPath := "/linkconnections/{linkConnectionName}/edittables"
	if linkConnectionName == "" {
		return nil, errors.New("parameter linkConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkConnectionName}", url.PathEscape(linkConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, editTablesRequest)
}

// Get - Get a link connection
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-12-01-preview
//   - linkConnectionName - The link connection name
//   - options - LinkConnectionClientGetOptions contains the optional parameters for the LinkConnectionClient.Get method.
func (client *LinkConnectionClient) Get(ctx context.Context, linkConnectionName string, options *LinkConnectionClientGetOptions) (LinkConnectionClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, linkConnectionName, options)
	if err != nil {
		return LinkConnectionClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LinkConnectionClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LinkConnectionClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LinkConnectionClient) getCreateRequest(ctx context.Context, linkConnectionName string, options *LinkConnectionClientGetOptions) (*policy.Request, error) {
	urlPath := "/linkconnections/{linkConnectionName}"
	if linkConnectionName == "" {
		return nil, errors.New("parameter linkConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkConnectionName}", url.PathEscape(linkConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LinkConnectionClient) getHandleResponse(resp *http.Response) (LinkConnectionClientGetResponse, error) {
	result := LinkConnectionClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LinkConnectionResource); err != nil {
		return LinkConnectionClientGetResponse{}, err
	}
	return result, nil
}

// GetDetailedStatus - Get the detailed status of a link connection
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-12-01-preview
//   - linkConnectionName - The link connection name
//   - options - LinkConnectionClientGetDetailedStatusOptions contains the optional parameters for the LinkConnectionClient.GetDetailedStatus
//     method.
func (client *LinkConnectionClient) GetDetailedStatus(ctx context.Context, linkConnectionName string, options *LinkConnectionClientGetDetailedStatusOptions) (LinkConnectionClientGetDetailedStatusResponse, error) {
	req, err := client.getDetailedStatusCreateRequest(ctx, linkConnectionName, options)
	if err != nil {
		return LinkConnectionClientGetDetailedStatusResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LinkConnectionClientGetDetailedStatusResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LinkConnectionClientGetDetailedStatusResponse{}, runtime.NewResponseError(resp)
	}
	return client.getDetailedStatusHandleResponse(resp)
}

// getDetailedStatusCreateRequest creates the GetDetailedStatus request.
func (client *LinkConnectionClient) getDetailedStatusCreateRequest(ctx context.Context, linkConnectionName string, options *LinkConnectionClientGetDetailedStatusOptions) (*policy.Request, error) {
	urlPath := "/linkconnections/{linkConnectionName}/detailedstatus"
	if linkConnectionName == "" {
		return nil, errors.New("parameter linkConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkConnectionName}", url.PathEscape(linkConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDetailedStatusHandleResponse handles the GetDetailedStatus response.
func (client *LinkConnectionClient) getDetailedStatusHandleResponse(resp *http.Response) (LinkConnectionClientGetDetailedStatusResponse, error) {
	result := LinkConnectionClientGetDetailedStatusResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LinkConnectionDetailedStatus); err != nil {
		return LinkConnectionClientGetDetailedStatusResponse{}, err
	}
	return result, nil
}

// NewListByWorkspacePager - List link connections
//
// Generated from API version 2022-12-01-preview
//   - options - LinkConnectionClientListByWorkspaceOptions contains the optional parameters for the LinkConnectionClient.NewListByWorkspacePager
//     method.
func (client *LinkConnectionClient) NewListByWorkspacePager(options *LinkConnectionClientListByWorkspaceOptions) *runtime.Pager[LinkConnectionClientListByWorkspaceResponse] {
	return runtime.NewPager(runtime.PagingHandler[LinkConnectionClientListByWorkspaceResponse]{
		More: func(page LinkConnectionClientListByWorkspaceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LinkConnectionClientListByWorkspaceResponse) (LinkConnectionClientListByWorkspaceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByWorkspaceCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return LinkConnectionClientListByWorkspaceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return LinkConnectionClientListByWorkspaceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return LinkConnectionClientListByWorkspaceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByWorkspaceHandleResponse(resp)
		},
	})
}

// listByWorkspaceCreateRequest creates the ListByWorkspace request.
func (client *LinkConnectionClient) listByWorkspaceCreateRequest(ctx context.Context, options *LinkConnectionClientListByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/linkconnections"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByWorkspaceHandleResponse handles the ListByWorkspace response.
func (client *LinkConnectionClient) listByWorkspaceHandleResponse(resp *http.Response) (LinkConnectionClientListByWorkspaceResponse, error) {
	result := LinkConnectionClientListByWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LinkConnectionListResponse); err != nil {
		return LinkConnectionClientListByWorkspaceResponse{}, err
	}
	return result, nil
}

// ListLinkTables - List the link tables of a link connection
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-12-01-preview
//   - linkConnectionName - The link connection name
//   - options - LinkConnectionClientListLinkTablesOptions contains the optional parameters for the LinkConnectionClient.ListLinkTables
//     method.
func (client *LinkConnectionClient) ListLinkTables(ctx context.Context, linkConnectionName string, options *LinkConnectionClientListLinkTablesOptions) (LinkConnectionClientListLinkTablesResponse, error) {
	req, err := client.listLinkTablesCreateRequest(ctx, linkConnectionName, options)
	if err != nil {
		return LinkConnectionClientListLinkTablesResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LinkConnectionClientListLinkTablesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LinkConnectionClientListLinkTablesResponse{}, runtime.NewResponseError(resp)
	}
	return client.listLinkTablesHandleResponse(resp)
}

// listLinkTablesCreateRequest creates the ListLinkTables request.
func (client *LinkConnectionClient) listLinkTablesCreateRequest(ctx context.Context, linkConnectionName string, options *LinkConnectionClientListLinkTablesOptions) (*policy.Request, error) {
	urlPath := "/linkconnections/{linkConnectionName}/linktables"
	if linkConnectionName == "" {
		return nil, errors.New("parameter linkConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkConnectionName}", url.PathEscape(linkConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listLinkTablesHandleResponse handles the ListLinkTables response.
func (client *LinkConnectionClient) listLinkTablesHandleResponse(resp *http.Response) (LinkConnectionClientListLinkTablesResponse, error) {
	result := LinkConnectionClientListLinkTablesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LinkTableListResponse); err != nil {
		return LinkConnectionClientListLinkTablesResponse{}, err
	}
	return result, nil
}

// Pause - Pause a link connection. It may take a few minutes from Pausing to Paused, monitor the status with LinkConnection_GetDetailedStatus.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-12-01-preview
//   - linkConnectionName - The link connection name
//   - options - LinkConnectionClientPauseOptions contains the optional parameters for the LinkConnectionClient.Pause method.
func (client *LinkConnectionClient) Pause(ctx context.Context, linkConnectionName string, options *LinkConnectionClientPauseOptions) (LinkConnectionClientPauseResponse, error) {
	req, err := client.pauseCreateRequest(ctx, linkConnectionName, options)
	if err != nil {
		return LinkConnectionClientPauseResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LinkConnectionClientPauseResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LinkConnectionClientPauseResponse{}, runtime.NewResponseError(resp)
	}
	return LinkConnectionClientPauseResponse{}, nil
}

// pauseCreateRequest creates the Pause request.
func (client *LinkConnectionClient) pauseCreateRequest(ctx context.Context, linkConnectionName string, options *LinkConnectionClientPauseOptions) (*policy.Request, error) {
	urlPath := "/linkconnections/{linkConnectionName}/pause"
	if linkConnectionName == "" {
		return nil, errors.New("parameter linkConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkConnectionName}", url.PathEscape(linkConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// QueryTableStatus - Query the link table status of a link connection
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-12-01-preview
//   - linkConnectionName - The link connection name
//   - queryTableStatusRequest - Query table status request
//   - options - LinkConnectionClientQueryTableStatusOptions contains the optional parameters for the LinkConnectionClient.QueryTableStatus
//     method.
func (client *LinkConnectionClient) QueryTableStatus(ctx context.Context, linkConnectionName string, queryTableStatusRequest QueryTableStatusRequest, options *LinkConnectionClientQueryTableStatusOptions) (LinkConnectionClientQueryTableStatusResponse, error) {
	req, err := client.queryTableStatusCreateRequest(ctx, linkConnectionName, queryTableStatusRequest, options)
	if err != nil {
		return LinkConnectionClientQueryTableStatusResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LinkConnectionClientQueryTableStatusResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LinkConnectionClientQueryTableStatusResponse{}, runtime.NewResponseError(resp)
	}
	return client.queryTableStatusHandleResponse(resp)
}

// queryTableStatusCreateRequest creates the QueryTableStatus request.
func (client *LinkConnectionClient) queryTableStatusCreateRequest(ctx context.Context, linkConnectionName string, queryTableStatusRequest QueryTableStatusRequest, options *LinkConnectionClientQueryTableStatusOptions) (*policy.Request, error) {
	urlPath := "/linkconnections/{linkConnectionName}/querytablestatus"
	if linkConnectionName == "" {
		return nil, errors.New("parameter linkConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkConnectionName}", url.PathEscape(linkConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, queryTableStatusRequest)
}

// queryTableStatusHandleResponse handles the QueryTableStatus response.
func (client *LinkConnectionClient) queryTableStatusHandleResponse(resp *http.Response) (LinkConnectionClientQueryTableStatusResponse, error) {
	result := LinkConnectionClientQueryTableStatusResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LinkConnectionQueryTableStatus); err != nil {
		return LinkConnectionClientQueryTableStatusResponse{}, err
	}
	return result, nil
}

// Resume - Resume a link connection. It may take a few minutes from Resuming to Running, monitor the status with LinkConnection_GetDetailedStatus.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-12-01-preview
//   - linkConnectionName - The link connection name
//   - options - LinkConnectionClientResumeOptions contains the optional parameters for the LinkConnectionClient.Resume method.
func (client *LinkConnectionClient) Resume(ctx context.Context, linkConnectionName string, options *LinkConnectionClientResumeOptions) (LinkConnectionClientResumeResponse, error) {
	req, err := client.resumeCreateRequest(ctx, linkConnectionName, options)
	if err != nil {
		return LinkConnectionClientResumeResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LinkConnectionClientResumeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LinkConnectionClientResumeResponse{}, runtime.NewResponseError(resp)
	}
	return LinkConnectionClientResumeResponse{}, nil
}

// resumeCreateRequest creates the Resume request.
func (client *LinkConnectionClient) resumeCreateRequest(ctx context.Context, linkConnectionName string, options *LinkConnectionClientResumeOptions) (*policy.Request, error) {
	urlPath := "/linkconnections/{linkConnectionName}/resume"
	if linkConnectionName == "" {
		return nil, errors.New("parameter linkConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkConnectionName}", url.PathEscape(linkConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Start - Start a link connection. It may take a few minutes from Starting to Running, monitor the status with LinkConnection_GetDetailedStatus.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-12-01-preview
//   - linkConnectionName - The link connection name
//   - options - LinkConnectionClientStartOptions contains the optional parameters for the LinkConnectionClient.Start method.
func (client *LinkConnectionClient) Start(ctx context.Context, linkConnectionName string, options *LinkConnectionClientStartOptions) (LinkConnectionClientStartResponse, error) {
	req, err := client.startCreateRequest(ctx, linkConnectionName, options)
	if err != nil {
		return LinkConnectionClientStartResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LinkConnectionClientStartResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LinkConnectionClientStartResponse{}, runtime.NewResponseError(resp)
	}
	return LinkConnectionClientStartResponse{}, nil
}

// startCreateRequest creates the Start request.
func (client *LinkConnectionClient) startCreateRequest(ctx context.Context, linkConnectionName string, options *LinkConnectionClientStartOptions) (*policy.Request, error) {
	urlPath := "/linkconnections/{linkConnectionName}/start"
	if linkConnectionName == "" {
		return nil, errors.New("parameter linkConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkConnectionName}", url.PathEscape(linkConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Stop - Stop a link connection. It may take a few minutes from Stopping to stopped, monitor the status with LinkConnection_GetDetailedStatus.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-12-01-preview
//   - linkConnectionName - The link connection name
//   - options - LinkConnectionClientStopOptions contains the optional parameters for the LinkConnectionClient.Stop method.
func (client *LinkConnectionClient) Stop(ctx context.Context, linkConnectionName string, options *LinkConnectionClientStopOptions) (LinkConnectionClientStopResponse, error) {
	req, err := client.stopCreateRequest(ctx, linkConnectionName, options)
	if err != nil {
		return LinkConnectionClientStopResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LinkConnectionClientStopResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LinkConnectionClientStopResponse{}, runtime.NewResponseError(resp)
	}
	return LinkConnectionClientStopResponse{}, nil
}

// stopCreateRequest creates the Stop request.
func (client *LinkConnectionClient) stopCreateRequest(ctx context.Context, linkConnectionName string, options *LinkConnectionClientStopOptions) (*policy.Request, error) {
	urlPath := "/linkconnections/{linkConnectionName}/stop"
	if linkConnectionName == "" {
		return nil, errors.New("parameter linkConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkConnectionName}", url.PathEscape(linkConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// UpdateLandingZoneCredential - Update landing zone credential of a link connection
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-12-01-preview
//   - linkConnectionName - The link connection name
//   - updateLandingZoneCredentialRequest - update landing zone credential request
//   - options - LinkConnectionClientUpdateLandingZoneCredentialOptions contains the optional parameters for the LinkConnectionClient.UpdateLandingZoneCredential
//     method.
func (client *LinkConnectionClient) UpdateLandingZoneCredential(ctx context.Context, linkConnectionName string, updateLandingZoneCredentialRequest UpdateLandingZoneCredential, options *LinkConnectionClientUpdateLandingZoneCredentialOptions) (LinkConnectionClientUpdateLandingZoneCredentialResponse, error) {
	req, err := client.updateLandingZoneCredentialCreateRequest(ctx, linkConnectionName, updateLandingZoneCredentialRequest, options)
	if err != nil {
		return LinkConnectionClientUpdateLandingZoneCredentialResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LinkConnectionClientUpdateLandingZoneCredentialResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LinkConnectionClientUpdateLandingZoneCredentialResponse{}, runtime.NewResponseError(resp)
	}
	return LinkConnectionClientUpdateLandingZoneCredentialResponse{}, nil
}

// updateLandingZoneCredentialCreateRequest creates the UpdateLandingZoneCredential request.
func (client *LinkConnectionClient) updateLandingZoneCredentialCreateRequest(ctx context.Context, linkConnectionName string, updateLandingZoneCredentialRequest UpdateLandingZoneCredential, options *LinkConnectionClientUpdateLandingZoneCredentialOptions) (*policy.Request, error) {
	urlPath := "/linkconnections/{linkConnectionName}/updateLandingZoneCredential"
	if linkConnectionName == "" {
		return nil, errors.New("parameter linkConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkConnectionName}", url.PathEscape(linkConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, updateLandingZoneCredentialRequest)
}
