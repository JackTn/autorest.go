// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

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

// VirtualAppliancesClient contains the methods for the NetworkVirtualAppliances group.
// Don't use this type directly, use NewVirtualAppliancesClient() instead.
type VirtualAppliancesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVirtualAppliancesClient creates a new instance of VirtualAppliancesClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVirtualAppliancesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VirtualAppliancesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VirtualAppliancesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates the specified Network Virtual Appliance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - networkVirtualApplianceName - The name of Network Virtual Appliance.
//   - parameters - Parameters supplied to the create or update Network Virtual Appliance.
//   - options - VirtualAppliancesClientBeginCreateOrUpdateOptions contains the optional parameters for the VirtualAppliancesClient.BeginCreateOrUpdate
//     method.
func (client *VirtualAppliancesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, parameters VirtualAppliance, options *VirtualAppliancesClientBeginCreateOrUpdateOptions) (*runtime.Poller[VirtualAppliancesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, networkVirtualApplianceName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualAppliancesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VirtualAppliancesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates the specified Network Virtual Appliance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *VirtualAppliancesClient) createOrUpdate(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, parameters VirtualAppliance, options *VirtualAppliancesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "VirtualAppliancesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, networkVirtualApplianceName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VirtualAppliancesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, parameters VirtualAppliance, options *VirtualAppliancesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkVirtualApplianceName == "" {
		return nil, errors.New("parameter networkVirtualApplianceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkVirtualApplianceName}", url.PathEscape(networkVirtualApplianceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the specified Network Virtual Appliance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - networkVirtualApplianceName - The name of Network Virtual Appliance.
//   - options - VirtualAppliancesClientBeginDeleteOptions contains the optional parameters for the VirtualAppliancesClient.BeginDelete
//     method.
func (client *VirtualAppliancesClient) BeginDelete(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, options *VirtualAppliancesClientBeginDeleteOptions) (*runtime.Poller[VirtualAppliancesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, networkVirtualApplianceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualAppliancesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VirtualAppliancesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes the specified Network Virtual Appliance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *VirtualAppliancesClient) deleteOperation(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, options *VirtualAppliancesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "VirtualAppliancesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkVirtualApplianceName, options)
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

// deleteCreateRequest creates the Delete request.
func (client *VirtualAppliancesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, options *VirtualAppliancesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkVirtualApplianceName == "" {
		return nil, errors.New("parameter networkVirtualApplianceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkVirtualApplianceName}", url.PathEscape(networkVirtualApplianceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the specified Network Virtual Appliance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - networkVirtualApplianceName - The name of Network Virtual Appliance.
//   - options - VirtualAppliancesClientGetOptions contains the optional parameters for the VirtualAppliancesClient.Get method.
func (client *VirtualAppliancesClient) Get(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, options *VirtualAppliancesClientGetOptions) (VirtualAppliancesClientGetResponse, error) {
	var err error
	const operationName = "VirtualAppliancesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkVirtualApplianceName, options)
	if err != nil {
		return VirtualAppliancesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualAppliancesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VirtualAppliancesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *VirtualAppliancesClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, options *VirtualAppliancesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkVirtualApplianceName == "" {
		return nil, errors.New("parameter networkVirtualApplianceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkVirtualApplianceName}", url.PathEscape(networkVirtualApplianceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualAppliancesClient) getHandleResponse(resp *http.Response) (VirtualAppliancesClientGetResponse, error) {
	result := VirtualAppliancesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualAppliance); err != nil {
		return VirtualAppliancesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all Network Virtual Appliances in a subscription.
//
// Generated from API version 2022-09-01
//   - options - VirtualAppliancesClientListOptions contains the optional parameters for the VirtualAppliancesClient.NewListPager
//     method.
func (client *VirtualAppliancesClient) NewListPager(options *VirtualAppliancesClientListOptions) *runtime.Pager[VirtualAppliancesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualAppliancesClientListResponse]{
		More: func(page VirtualAppliancesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualAppliancesClientListResponse) (VirtualAppliancesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "VirtualAppliancesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return VirtualAppliancesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *VirtualAppliancesClient) listCreateRequest(ctx context.Context, options *VirtualAppliancesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkVirtualAppliances"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualAppliancesClient) listHandleResponse(resp *http.Response) (VirtualAppliancesClientListResponse, error) {
	result := VirtualAppliancesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualApplianceListResult); err != nil {
		return VirtualAppliancesClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists all Network Virtual Appliances in a resource group.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - options - VirtualAppliancesClientListByResourceGroupOptions contains the optional parameters for the VirtualAppliancesClient.NewListByResourceGroupPager
//     method.
func (client *VirtualAppliancesClient) NewListByResourceGroupPager(resourceGroupName string, options *VirtualAppliancesClientListByResourceGroupOptions) *runtime.Pager[VirtualAppliancesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualAppliancesClientListByResourceGroupResponse]{
		More: func(page VirtualAppliancesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualAppliancesClientListByResourceGroupResponse) (VirtualAppliancesClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "VirtualAppliancesClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return VirtualAppliancesClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *VirtualAppliancesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *VirtualAppliancesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *VirtualAppliancesClient) listByResourceGroupHandleResponse(resp *http.Response) (VirtualAppliancesClientListByResourceGroupResponse, error) {
	result := VirtualAppliancesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualApplianceListResult); err != nil {
		return VirtualAppliancesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates a Network Virtual Appliance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The resource group name of Network Virtual Appliance.
//   - networkVirtualApplianceName - The name of Network Virtual Appliance being updated.
//   - parameters - Parameters supplied to Update Network Virtual Appliance Tags.
//   - options - VirtualAppliancesClientUpdateTagsOptions contains the optional parameters for the VirtualAppliancesClient.UpdateTags
//     method.
func (client *VirtualAppliancesClient) UpdateTags(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, parameters TagsObject, options *VirtualAppliancesClientUpdateTagsOptions) (VirtualAppliancesClientUpdateTagsResponse, error) {
	var err error
	const operationName = "VirtualAppliancesClient.UpdateTags"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, networkVirtualApplianceName, parameters, options)
	if err != nil {
		return VirtualAppliancesClientUpdateTagsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualAppliancesClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VirtualAppliancesClientUpdateTagsResponse{}, err
	}
	resp, err := client.updateTagsHandleResponse(httpResp)
	return resp, err
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *VirtualAppliancesClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, parameters TagsObject, options *VirtualAppliancesClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkVirtualApplianceName == "" {
		return nil, errors.New("parameter networkVirtualApplianceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkVirtualApplianceName}", url.PathEscape(networkVirtualApplianceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *VirtualAppliancesClient) updateTagsHandleResponse(resp *http.Response) (VirtualAppliancesClientUpdateTagsResponse, error) {
	result := VirtualAppliancesClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualAppliance); err != nil {
		return VirtualAppliancesClientUpdateTagsResponse{}, err
	}
	return result, nil
}
