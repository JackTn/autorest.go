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

// LinkedServiceClient contains the methods for the LinkedService group.
// Don't use this type directly, use a constructor function instead.
type LinkedServiceClient struct {
	internal *azcore.Client
	endpoint string
}

// BeginCreateOrUpdateLinkedService - Creates or updates a linked service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - linkedServiceName - The linked service name.
//   - linkedService - Linked service resource definition.
//   - options - LinkedServiceClientBeginCreateOrUpdateLinkedServiceOptions contains the optional parameters for the LinkedServiceClient.BeginCreateOrUpdateLinkedService
//     method.
func (client *LinkedServiceClient) BeginCreateOrUpdateLinkedService(ctx context.Context, linkedServiceName string, linkedService LinkedServiceResource, options *LinkedServiceClientBeginCreateOrUpdateLinkedServiceOptions) (*runtime.Poller[LinkedServiceClientCreateOrUpdateLinkedServiceResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdateLinkedService(ctx, linkedServiceName, linkedService, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[LinkedServiceClientCreateOrUpdateLinkedServiceResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[LinkedServiceClientCreateOrUpdateLinkedServiceResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdateLinkedService - Creates or updates a linked service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
func (client *LinkedServiceClient) createOrUpdateLinkedService(ctx context.Context, linkedServiceName string, linkedService LinkedServiceResource, options *LinkedServiceClientBeginCreateOrUpdateLinkedServiceOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateLinkedServiceCreateRequest(ctx, linkedServiceName, linkedService, options)
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

// createOrUpdateLinkedServiceCreateRequest creates the CreateOrUpdateLinkedService request.
func (client *LinkedServiceClient) createOrUpdateLinkedServiceCreateRequest(ctx context.Context, linkedServiceName string, linkedService LinkedServiceResource, options *LinkedServiceClientBeginCreateOrUpdateLinkedServiceOptions) (*policy.Request, error) {
	urlPath := "/linkedservices/{linkedServiceName}"
	if linkedServiceName == "" {
		return nil, errors.New("parameter linkedServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkedServiceName}", url.PathEscape(linkedServiceName))
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
	if err := runtime.MarshalAsJSON(req, linkedService); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDeleteLinkedService - Deletes a linked service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - linkedServiceName - The linked service name.
//   - options - LinkedServiceClientBeginDeleteLinkedServiceOptions contains the optional parameters for the LinkedServiceClient.BeginDeleteLinkedService
//     method.
func (client *LinkedServiceClient) BeginDeleteLinkedService(ctx context.Context, linkedServiceName string, options *LinkedServiceClientBeginDeleteLinkedServiceOptions) (*runtime.Poller[LinkedServiceClientDeleteLinkedServiceResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteLinkedService(ctx, linkedServiceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[LinkedServiceClientDeleteLinkedServiceResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[LinkedServiceClientDeleteLinkedServiceResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// DeleteLinkedService - Deletes a linked service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
func (client *LinkedServiceClient) deleteLinkedService(ctx context.Context, linkedServiceName string, options *LinkedServiceClientBeginDeleteLinkedServiceOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteLinkedServiceCreateRequest(ctx, linkedServiceName, options)
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

// deleteLinkedServiceCreateRequest creates the DeleteLinkedService request.
func (client *LinkedServiceClient) deleteLinkedServiceCreateRequest(ctx context.Context, linkedServiceName string, _ *LinkedServiceClientBeginDeleteLinkedServiceOptions) (*policy.Request, error) {
	urlPath := "/linkedservices/{linkedServiceName}"
	if linkedServiceName == "" {
		return nil, errors.New("parameter linkedServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkedServiceName}", url.PathEscape(linkedServiceName))
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

// GetLinkedService - Gets a linked service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - linkedServiceName - The linked service name.
//   - options - LinkedServiceClientGetLinkedServiceOptions contains the optional parameters for the LinkedServiceClient.GetLinkedService
//     method.
func (client *LinkedServiceClient) GetLinkedService(ctx context.Context, linkedServiceName string, options *LinkedServiceClientGetLinkedServiceOptions) (LinkedServiceClientGetLinkedServiceResponse, error) {
	var err error
	req, err := client.getLinkedServiceCreateRequest(ctx, linkedServiceName, options)
	if err != nil {
		return LinkedServiceClientGetLinkedServiceResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LinkedServiceClientGetLinkedServiceResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNotModified) {
		err = runtime.NewResponseError(httpResp)
		return LinkedServiceClientGetLinkedServiceResponse{}, err
	}
	resp, err := client.getLinkedServiceHandleResponse(httpResp)
	return resp, err
}

// getLinkedServiceCreateRequest creates the GetLinkedService request.
func (client *LinkedServiceClient) getLinkedServiceCreateRequest(ctx context.Context, linkedServiceName string, options *LinkedServiceClientGetLinkedServiceOptions) (*policy.Request, error) {
	urlPath := "/linkedservices/{linkedServiceName}"
	if linkedServiceName == "" {
		return nil, errors.New("parameter linkedServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkedServiceName}", url.PathEscape(linkedServiceName))
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

// getLinkedServiceHandleResponse handles the GetLinkedService response.
func (client *LinkedServiceClient) getLinkedServiceHandleResponse(resp *http.Response) (LinkedServiceClientGetLinkedServiceResponse, error) {
	result := LinkedServiceClientGetLinkedServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LinkedServiceResource); err != nil {
		return LinkedServiceClientGetLinkedServiceResponse{}, err
	}
	return result, nil
}

// NewGetLinkedServicesByWorkspacePager - Lists linked services.
//
// Generated from API version 2020-12-01
//   - options - LinkedServiceClientGetLinkedServicesByWorkspaceOptions contains the optional parameters for the LinkedServiceClient.NewGetLinkedServicesByWorkspacePager
//     method.
func (client *LinkedServiceClient) NewGetLinkedServicesByWorkspacePager(options *LinkedServiceClientGetLinkedServicesByWorkspaceOptions) *runtime.Pager[LinkedServiceClientGetLinkedServicesByWorkspaceResponse] {
	return runtime.NewPager(runtime.PagingHandler[LinkedServiceClientGetLinkedServicesByWorkspaceResponse]{
		More: func(page LinkedServiceClientGetLinkedServicesByWorkspaceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LinkedServiceClientGetLinkedServicesByWorkspaceResponse) (LinkedServiceClientGetLinkedServicesByWorkspaceResponse, error) {
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.getLinkedServicesByWorkspaceCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return LinkedServiceClientGetLinkedServicesByWorkspaceResponse{}, err
			}
			return client.getLinkedServicesByWorkspaceHandleResponse(resp)
		},
	})
}

// getLinkedServicesByWorkspaceCreateRequest creates the GetLinkedServicesByWorkspace request.
func (client *LinkedServiceClient) getLinkedServicesByWorkspaceCreateRequest(ctx context.Context, _ *LinkedServiceClientGetLinkedServicesByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/linkedservices"
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

// getLinkedServicesByWorkspaceHandleResponse handles the GetLinkedServicesByWorkspace response.
func (client *LinkedServiceClient) getLinkedServicesByWorkspaceHandleResponse(resp *http.Response) (LinkedServiceClientGetLinkedServicesByWorkspaceResponse, error) {
	result := LinkedServiceClientGetLinkedServicesByWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LinkedServiceListResponse); err != nil {
		return LinkedServiceClientGetLinkedServicesByWorkspaceResponse{}, err
	}
	return result, nil
}

// BeginRenameLinkedService - Renames a linked service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - linkedServiceName - The linked service name.
//   - request - proposed new name.
//   - options - LinkedServiceClientBeginRenameLinkedServiceOptions contains the optional parameters for the LinkedServiceClient.BeginRenameLinkedService
//     method.
func (client *LinkedServiceClient) BeginRenameLinkedService(ctx context.Context, linkedServiceName string, request ArtifactRenameRequest, options *LinkedServiceClientBeginRenameLinkedServiceOptions) (*runtime.Poller[LinkedServiceClientRenameLinkedServiceResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.renameLinkedService(ctx, linkedServiceName, request, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[LinkedServiceClientRenameLinkedServiceResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[LinkedServiceClientRenameLinkedServiceResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// RenameLinkedService - Renames a linked service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
func (client *LinkedServiceClient) renameLinkedService(ctx context.Context, linkedServiceName string, request ArtifactRenameRequest, options *LinkedServiceClientBeginRenameLinkedServiceOptions) (*http.Response, error) {
	var err error
	req, err := client.renameLinkedServiceCreateRequest(ctx, linkedServiceName, request, options)
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

// renameLinkedServiceCreateRequest creates the RenameLinkedService request.
func (client *LinkedServiceClient) renameLinkedServiceCreateRequest(ctx context.Context, linkedServiceName string, request ArtifactRenameRequest, _ *LinkedServiceClientBeginRenameLinkedServiceOptions) (*policy.Request, error) {
	urlPath := "/linkedservices/{linkedServiceName}/rename"
	if linkedServiceName == "" {
		return nil, errors.New("parameter linkedServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkedServiceName}", url.PathEscape(linkedServiceName))
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
