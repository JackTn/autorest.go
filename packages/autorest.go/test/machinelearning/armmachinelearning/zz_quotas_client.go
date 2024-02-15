// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmachinelearning

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// QuotasClient contains the methods for the Quotas group.
// Don't use this type directly, use NewQuotasClient() instead.
type QuotasClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewQuotasClient creates a new instance of QuotasClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewQuotasClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*QuotasClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &QuotasClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Gets the currently assigned Workspace Quotas based on VMFamily.
//
// Generated from API version 2022-02-01-preview
//   - location - The location for which resource usage is queried.
//   - options - QuotasClientListOptions contains the optional parameters for the QuotasClient.NewListPager method.
func (client *QuotasClient) NewListPager(location string, options *QuotasClientListOptions) *runtime.Pager[QuotasClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[QuotasClientListResponse]{
		More: func(page QuotasClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *QuotasClientListResponse) (QuotasClientListResponse, error) {
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, location, options)
			}, nil)
			if err != nil {
				return QuotasClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *QuotasClient) listCreateRequest(ctx context.Context, location string, options *QuotasClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.MachineLearningServices/locations/{location}/quotas"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *QuotasClient) listHandleResponse(resp *http.Response) (QuotasClientListResponse, error) {
	result := QuotasClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListWorkspaceQuotas); err != nil {
		return QuotasClientListResponse{}, err
	}
	return result, nil
}

// Update - Update quota for each VM family in workspace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-02-01-preview
//   - location - The location for update quota is queried.
//   - parameters - Quota update parameters.
//   - options - QuotasClientUpdateOptions contains the optional parameters for the QuotasClient.Update method.
func (client *QuotasClient) Update(ctx context.Context, location string, parameters QuotaUpdateParameters, options *QuotasClientUpdateOptions) (QuotasClientUpdateResponse, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, location, parameters, options)
	if err != nil {
		return QuotasClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return QuotasClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return QuotasClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *QuotasClient) updateCreateRequest(ctx context.Context, location string, parameters QuotaUpdateParameters, options *QuotasClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.MachineLearningServices/locations/{location}/updateQuotas"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *QuotasClient) updateHandleResponse(resp *http.Response) (QuotasClientUpdateResponse, error) {
	result := QuotasClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UpdateWorkspaceQuotasResult); err != nil {
		return QuotasClientUpdateResponse{}, err
	}
	return result, nil
}
