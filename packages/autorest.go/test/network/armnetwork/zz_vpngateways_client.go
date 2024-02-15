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

// VPNGatewaysClient contains the methods for the VPNGateways group.
// Don't use this type directly, use NewVPNGatewaysClient() instead.
type VPNGatewaysClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVPNGatewaysClient creates a new instance of VPNGatewaysClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVPNGatewaysClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VPNGatewaysClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VPNGatewaysClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates a virtual wan vpn gateway if it doesn't exist else updates the existing gateway.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The resource group name of the VpnGateway.
//   - gatewayName - The name of the gateway.
//   - vpnGatewayParameters - Parameters supplied to create or Update a virtual wan vpn gateway.
//   - options - VPNGatewaysClientBeginCreateOrUpdateOptions contains the optional parameters for the VPNGatewaysClient.BeginCreateOrUpdate
//     method.
func (client *VPNGatewaysClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, gatewayName string, vpnGatewayParameters VPNGateway, options *VPNGatewaysClientBeginCreateOrUpdateOptions) (*runtime.Poller[VPNGatewaysClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, gatewayName, vpnGatewayParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VPNGatewaysClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VPNGatewaysClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates a virtual wan vpn gateway if it doesn't exist else updates the existing gateway.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *VPNGatewaysClient) createOrUpdate(ctx context.Context, resourceGroupName string, gatewayName string, vpnGatewayParameters VPNGateway, options *VPNGatewaysClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "VPNGatewaysClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, gatewayName, vpnGatewayParameters, options)
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
func (client *VPNGatewaysClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, vpnGatewayParameters VPNGateway, options *VPNGatewaysClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, vpnGatewayParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes a virtual wan vpn gateway.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The resource group name of the VpnGateway.
//   - gatewayName - The name of the gateway.
//   - options - VPNGatewaysClientBeginDeleteOptions contains the optional parameters for the VPNGatewaysClient.BeginDelete method.
func (client *VPNGatewaysClient) BeginDelete(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysClientBeginDeleteOptions) (*runtime.Poller[VPNGatewaysClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, gatewayName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VPNGatewaysClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VPNGatewaysClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes a virtual wan vpn gateway.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *VPNGatewaysClient) deleteOperation(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "VPNGatewaysClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, gatewayName, options)
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
func (client *VPNGatewaysClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
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

// Get - Retrieves the details of a virtual wan vpn gateway.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The resource group name of the VpnGateway.
//   - gatewayName - The name of the gateway.
//   - options - VPNGatewaysClientGetOptions contains the optional parameters for the VPNGatewaysClient.Get method.
func (client *VPNGatewaysClient) Get(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysClientGetOptions) (VPNGatewaysClientGetResponse, error) {
	var err error
	const operationName = "VPNGatewaysClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, gatewayName, options)
	if err != nil {
		return VPNGatewaysClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VPNGatewaysClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VPNGatewaysClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *VPNGatewaysClient) getCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
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

// getHandleResponse handles the Get response.
func (client *VPNGatewaysClient) getHandleResponse(resp *http.Response) (VPNGatewaysClientGetResponse, error) {
	result := VPNGatewaysClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VPNGateway); err != nil {
		return VPNGatewaysClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all the VpnGateways in a subscription.
//
// Generated from API version 2022-09-01
//   - options - VPNGatewaysClientListOptions contains the optional parameters for the VPNGatewaysClient.NewListPager method.
func (client *VPNGatewaysClient) NewListPager(options *VPNGatewaysClientListOptions) *runtime.Pager[VPNGatewaysClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[VPNGatewaysClientListResponse]{
		More: func(page VPNGatewaysClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VPNGatewaysClientListResponse) (VPNGatewaysClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "VPNGatewaysClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return VPNGatewaysClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *VPNGatewaysClient) listCreateRequest(ctx context.Context, options *VPNGatewaysClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/vpnGateways"
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
func (client *VPNGatewaysClient) listHandleResponse(resp *http.Response) (VPNGatewaysClientListResponse, error) {
	result := VPNGatewaysClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListVPNGatewaysResult); err != nil {
		return VPNGatewaysClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists all the VpnGateways in a resource group.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The resource group name of the VpnGateway.
//   - options - VPNGatewaysClientListByResourceGroupOptions contains the optional parameters for the VPNGatewaysClient.NewListByResourceGroupPager
//     method.
func (client *VPNGatewaysClient) NewListByResourceGroupPager(resourceGroupName string, options *VPNGatewaysClientListByResourceGroupOptions) *runtime.Pager[VPNGatewaysClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[VPNGatewaysClientListByResourceGroupResponse]{
		More: func(page VPNGatewaysClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VPNGatewaysClientListByResourceGroupResponse) (VPNGatewaysClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "VPNGatewaysClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return VPNGatewaysClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *VPNGatewaysClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *VPNGatewaysClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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
func (client *VPNGatewaysClient) listByResourceGroupHandleResponse(resp *http.Response) (VPNGatewaysClientListByResourceGroupResponse, error) {
	result := VPNGatewaysClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListVPNGatewaysResult); err != nil {
		return VPNGatewaysClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// BeginReset - Resets the primary of the vpn gateway in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The resource group name of the VpnGateway.
//   - gatewayName - The name of the gateway.
//   - options - VPNGatewaysClientBeginResetOptions contains the optional parameters for the VPNGatewaysClient.BeginReset method.
func (client *VPNGatewaysClient) BeginReset(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysClientBeginResetOptions) (*runtime.Poller[VPNGatewaysClientResetResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.reset(ctx, resourceGroupName, gatewayName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VPNGatewaysClientResetResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VPNGatewaysClientResetResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Reset - Resets the primary of the vpn gateway in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *VPNGatewaysClient) reset(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysClientBeginResetOptions) (*http.Response, error) {
	var err error
	const operationName = "VPNGatewaysClient.BeginReset"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.resetCreateRequest(ctx, resourceGroupName, gatewayName, options)
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

// resetCreateRequest creates the Reset request.
func (client *VPNGatewaysClient) resetCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysClientBeginResetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/reset"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	if options != nil && options.IPConfigurationID != nil {
		reqQP.Set("ipConfigurationId", *options.IPConfigurationID)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginStartPacketCapture - Starts packet capture on vpn gateway in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The resource group name of the VpnGateway.
//   - gatewayName - The name of the gateway.
//   - options - VPNGatewaysClientBeginStartPacketCaptureOptions contains the optional parameters for the VPNGatewaysClient.BeginStartPacketCapture
//     method.
func (client *VPNGatewaysClient) BeginStartPacketCapture(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysClientBeginStartPacketCaptureOptions) (*runtime.Poller[VPNGatewaysClientStartPacketCaptureResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.startPacketCapture(ctx, resourceGroupName, gatewayName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VPNGatewaysClientStartPacketCaptureResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VPNGatewaysClientStartPacketCaptureResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// StartPacketCapture - Starts packet capture on vpn gateway in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *VPNGatewaysClient) startPacketCapture(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysClientBeginStartPacketCaptureOptions) (*http.Response, error) {
	var err error
	const operationName = "VPNGatewaysClient.BeginStartPacketCapture"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.startPacketCaptureCreateRequest(ctx, resourceGroupName, gatewayName, options)
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

// startPacketCaptureCreateRequest creates the StartPacketCapture request.
func (client *VPNGatewaysClient) startPacketCaptureCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysClientBeginStartPacketCaptureOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/startpacketcapture"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Parameters != nil {
		if err := runtime.MarshalAsJSON(req, *options.Parameters); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// BeginStopPacketCapture - Stops packet capture on vpn gateway in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The resource group name of the VpnGateway.
//   - gatewayName - The name of the gateway.
//   - options - VPNGatewaysClientBeginStopPacketCaptureOptions contains the optional parameters for the VPNGatewaysClient.BeginStopPacketCapture
//     method.
func (client *VPNGatewaysClient) BeginStopPacketCapture(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysClientBeginStopPacketCaptureOptions) (*runtime.Poller[VPNGatewaysClientStopPacketCaptureResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.stopPacketCapture(ctx, resourceGroupName, gatewayName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VPNGatewaysClientStopPacketCaptureResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VPNGatewaysClientStopPacketCaptureResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// StopPacketCapture - Stops packet capture on vpn gateway in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *VPNGatewaysClient) stopPacketCapture(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysClientBeginStopPacketCaptureOptions) (*http.Response, error) {
	var err error
	const operationName = "VPNGatewaysClient.BeginStopPacketCapture"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.stopPacketCaptureCreateRequest(ctx, resourceGroupName, gatewayName, options)
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

// stopPacketCaptureCreateRequest creates the StopPacketCapture request.
func (client *VPNGatewaysClient) stopPacketCaptureCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysClientBeginStopPacketCaptureOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/stoppacketcapture"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Parameters != nil {
		if err := runtime.MarshalAsJSON(req, *options.Parameters); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// BeginUpdateTags - Updates virtual wan vpn gateway tags.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The resource group name of the VpnGateway.
//   - gatewayName - The name of the gateway.
//   - vpnGatewayParameters - Parameters supplied to update a virtual wan vpn gateway tags.
//   - options - VPNGatewaysClientBeginUpdateTagsOptions contains the optional parameters for the VPNGatewaysClient.BeginUpdateTags
//     method.
func (client *VPNGatewaysClient) BeginUpdateTags(ctx context.Context, resourceGroupName string, gatewayName string, vpnGatewayParameters TagsObject, options *VPNGatewaysClientBeginUpdateTagsOptions) (*runtime.Poller[VPNGatewaysClientUpdateTagsResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updateTags(ctx, resourceGroupName, gatewayName, vpnGatewayParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VPNGatewaysClientUpdateTagsResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VPNGatewaysClientUpdateTagsResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// UpdateTags - Updates virtual wan vpn gateway tags.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *VPNGatewaysClient) updateTags(ctx context.Context, resourceGroupName string, gatewayName string, vpnGatewayParameters TagsObject, options *VPNGatewaysClientBeginUpdateTagsOptions) (*http.Response, error) {
	var err error
	const operationName = "VPNGatewaysClient.BeginUpdateTags"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, gatewayName, vpnGatewayParameters, options)
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

// updateTagsCreateRequest creates the UpdateTags request.
func (client *VPNGatewaysClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, vpnGatewayParameters TagsObject, options *VPNGatewaysClientBeginUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, vpnGatewayParameters); err != nil {
		return nil, err
	}
	return req, nil
}
