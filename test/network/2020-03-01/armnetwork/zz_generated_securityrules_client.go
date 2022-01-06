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

// SecurityRulesClient contains the methods for the SecurityRules group.
// Don't use this type directly, use NewSecurityRulesClient() instead.
type SecurityRulesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSecurityRulesClient creates a new instance of SecurityRulesClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSecurityRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *SecurityRulesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	client := &SecurityRulesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Host),
		pl:             armruntime.NewPipeline(module, version, credential, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates a security rule in the specified network security group.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// networkSecurityGroupName - The name of the network security group.
// securityRuleName - The name of the security rule.
// securityRuleParameters - Parameters supplied to the create or update network security rule operation.
// options - SecurityRulesClientBeginCreateOrUpdateOptions contains the optional parameters for the SecurityRulesClient.BeginCreateOrUpdate
// method.
func (client *SecurityRulesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string, securityRuleParameters SecurityRule, options *SecurityRulesClientBeginCreateOrUpdateOptions) (SecurityRulesClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, networkSecurityGroupName, securityRuleName, securityRuleParameters, options)
	if err != nil {
		return SecurityRulesClientCreateOrUpdatePollerResponse{}, err
	}
	result := SecurityRulesClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SecurityRulesClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return SecurityRulesClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &SecurityRulesClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a security rule in the specified network security group.
// If the operation fails it returns the *CloudError error type.
func (client *SecurityRulesClient) createOrUpdate(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string, securityRuleParameters SecurityRule, options *SecurityRulesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, networkSecurityGroupName, securityRuleName, securityRuleParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SecurityRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string, securityRuleParameters SecurityRule, options *SecurityRulesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules/{securityRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkSecurityGroupName == "" {
		return nil, errors.New("parameter networkSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
	if securityRuleName == "" {
		return nil, errors.New("parameter securityRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityRuleName}", url.PathEscape(securityRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, securityRuleParameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *SecurityRulesClient) createOrUpdateHandleError(resp *http.Response) error {
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

// BeginDelete - Deletes the specified network security rule.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// networkSecurityGroupName - The name of the network security group.
// securityRuleName - The name of the security rule.
// options - SecurityRulesClientBeginDeleteOptions contains the optional parameters for the SecurityRulesClient.BeginDelete
// method.
func (client *SecurityRulesClient) BeginDelete(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string, options *SecurityRulesClientBeginDeleteOptions) (SecurityRulesClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, networkSecurityGroupName, securityRuleName, options)
	if err != nil {
		return SecurityRulesClientDeletePollerResponse{}, err
	}
	result := SecurityRulesClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SecurityRulesClient.Delete", "location", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return SecurityRulesClientDeletePollerResponse{}, err
	}
	result.Poller = &SecurityRulesClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified network security rule.
// If the operation fails it returns the *CloudError error type.
func (client *SecurityRulesClient) deleteOperation(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string, options *SecurityRulesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkSecurityGroupName, securityRuleName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SecurityRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string, options *SecurityRulesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules/{securityRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkSecurityGroupName == "" {
		return nil, errors.New("parameter networkSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
	if securityRuleName == "" {
		return nil, errors.New("parameter securityRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityRuleName}", url.PathEscape(securityRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *SecurityRulesClient) deleteHandleError(resp *http.Response) error {
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

// Get - Get the specified network security rule.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// networkSecurityGroupName - The name of the network security group.
// securityRuleName - The name of the security rule.
// options - SecurityRulesClientGetOptions contains the optional parameters for the SecurityRulesClient.Get method.
func (client *SecurityRulesClient) Get(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string, options *SecurityRulesClientGetOptions) (SecurityRulesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkSecurityGroupName, securityRuleName, options)
	if err != nil {
		return SecurityRulesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SecurityRulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SecurityRulesClientGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SecurityRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string, options *SecurityRulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules/{securityRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkSecurityGroupName == "" {
		return nil, errors.New("parameter networkSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
	if securityRuleName == "" {
		return nil, errors.New("parameter securityRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityRuleName}", url.PathEscape(securityRuleName))
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
func (client *SecurityRulesClient) getHandleResponse(resp *http.Response) (SecurityRulesClientGetResponse, error) {
	result := SecurityRulesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityRule); err != nil {
		return SecurityRulesClientGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *SecurityRulesClient) getHandleError(resp *http.Response) error {
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

// List - Gets all security rules in a network security group.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// networkSecurityGroupName - The name of the network security group.
// options - SecurityRulesClientListOptions contains the optional parameters for the SecurityRulesClient.List method.
func (client *SecurityRulesClient) List(resourceGroupName string, networkSecurityGroupName string, options *SecurityRulesClientListOptions) *SecurityRulesClientListPager {
	return &SecurityRulesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, networkSecurityGroupName, options)
		},
		advancer: func(ctx context.Context, resp SecurityRulesClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SecurityRuleListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *SecurityRulesClient) listCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, options *SecurityRulesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules"
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
func (client *SecurityRulesClient) listHandleResponse(resp *http.Response) (SecurityRulesClientListResponse, error) {
	result := SecurityRulesClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityRuleListResult); err != nil {
		return SecurityRulesClientListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *SecurityRulesClient) listHandleError(resp *http.Response) error {
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
