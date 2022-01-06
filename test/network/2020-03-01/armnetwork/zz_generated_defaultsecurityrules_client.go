//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DefaultSecurityRulesClient contains the methods for the DefaultSecurityRules group.
// Don't use this type directly, use NewDefaultSecurityRulesClient() instead.
type DefaultSecurityRulesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDefaultSecurityRulesClient creates a new instance of DefaultSecurityRulesClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDefaultSecurityRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *DefaultSecurityRulesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	client := &DefaultSecurityRulesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Host),
		pl:             armruntime.NewPipeline(module, version, credential, &cp),
	}
	return client
}

// Get - Get the specified default network security rule.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// networkSecurityGroupName - The name of the network security group.
// defaultSecurityRuleName - The name of the default security rule.
// options - DefaultSecurityRulesClientGetOptions contains the optional parameters for the DefaultSecurityRulesClient.Get
// method.
func (client *DefaultSecurityRulesClient) Get(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, defaultSecurityRuleName string, options *DefaultSecurityRulesClientGetOptions) (DefaultSecurityRulesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkSecurityGroupName, defaultSecurityRuleName, options)
	if err != nil {
		return DefaultSecurityRulesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DefaultSecurityRulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DefaultSecurityRulesClientGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DefaultSecurityRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, defaultSecurityRuleName string, options *DefaultSecurityRulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/defaultSecurityRules/{defaultSecurityRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkSecurityGroupName == "" {
		return nil, errors.New("parameter networkSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
	if defaultSecurityRuleName == "" {
		return nil, errors.New("parameter defaultSecurityRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{defaultSecurityRuleName}", url.PathEscape(defaultSecurityRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DefaultSecurityRulesClient) getHandleResponse(resp *http.Response) (DefaultSecurityRulesClientGetResponse, error) {
	result := DefaultSecurityRulesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityRule); err != nil {
		return DefaultSecurityRulesClientGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *DefaultSecurityRulesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Gets all default security rules in a network security group.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// networkSecurityGroupName - The name of the network security group.
// options - DefaultSecurityRulesClientListOptions contains the optional parameters for the DefaultSecurityRulesClient.List
// method.
func (client *DefaultSecurityRulesClient) List(resourceGroupName string, networkSecurityGroupName string, options *DefaultSecurityRulesClientListOptions) *DefaultSecurityRulesClientListPager {
	return &DefaultSecurityRulesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, networkSecurityGroupName, options)
		},
		advancer: func(ctx context.Context, resp DefaultSecurityRulesClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SecurityRuleListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *DefaultSecurityRulesClient) listCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, options *DefaultSecurityRulesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/defaultSecurityRules"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkSecurityGroupName == "" {
		return nil, errors.New("parameter networkSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DefaultSecurityRulesClient) listHandleResponse(resp *http.Response) (DefaultSecurityRulesClientListResponse, error) {
	result := DefaultSecurityRulesClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityRuleListResult); err != nil {
		return DefaultSecurityRulesClientListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *DefaultSecurityRulesClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
