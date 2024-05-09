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

// FirewallPolicyIdpsSignaturesClient contains the methods for the FirewallPolicyIdpsSignatures group.
// Don't use this type directly, use NewFirewallPolicyIdpsSignaturesClient() instead.
type FirewallPolicyIdpsSignaturesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewFirewallPolicyIdpsSignaturesClient creates a new instance of FirewallPolicyIdpsSignaturesClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewFirewallPolicyIdpsSignaturesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*FirewallPolicyIdpsSignaturesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &FirewallPolicyIdpsSignaturesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// List - Retrieves the current status of IDPS signatures for the relevant policy
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - firewallPolicyName - The name of the Firewall Policy.
//   - options - FirewallPolicyIdpsSignaturesClientListOptions contains the optional parameters for the FirewallPolicyIdpsSignaturesClient.List
//     method.
func (client *FirewallPolicyIdpsSignaturesClient) List(ctx context.Context, resourceGroupName string, firewallPolicyName string, parameters IDPSQueryObject, options *FirewallPolicyIdpsSignaturesClientListOptions) (FirewallPolicyIdpsSignaturesClientListResponse, error) {
	var err error
	const operationName = "FirewallPolicyIdpsSignaturesClient.List"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listCreateRequest(ctx, resourceGroupName, firewallPolicyName, parameters, options)
	if err != nil {
		return FirewallPolicyIdpsSignaturesClientListResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FirewallPolicyIdpsSignaturesClientListResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return FirewallPolicyIdpsSignaturesClientListResponse{}, err
	}
	resp, err := client.listHandleResponse(httpResp)
	return resp, err
}

// listCreateRequest creates the List request.
func (client *FirewallPolicyIdpsSignaturesClient) listCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, parameters IDPSQueryObject, _ *FirewallPolicyIdpsSignaturesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/listIdpsSignatures"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if firewallPolicyName == "" {
		return nil, errors.New("parameter firewallPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
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
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *FirewallPolicyIdpsSignaturesClient) listHandleResponse(resp *http.Response) (FirewallPolicyIdpsSignaturesClientListResponse, error) {
	result := FirewallPolicyIdpsSignaturesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.QueryResults); err != nil {
		return FirewallPolicyIdpsSignaturesClientListResponse{}, err
	}
	return result, nil
}
