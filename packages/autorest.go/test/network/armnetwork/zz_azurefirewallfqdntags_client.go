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

// AzureFirewallFqdnTagsClient contains the methods for the AzureFirewallFqdnTags group.
// Don't use this type directly, use NewAzureFirewallFqdnTagsClient() instead.
type AzureFirewallFqdnTagsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAzureFirewallFqdnTagsClient creates a new instance of AzureFirewallFqdnTagsClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAzureFirewallFqdnTagsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AzureFirewallFqdnTagsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AzureFirewallFqdnTagsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListAllPager - Gets all the Azure Firewall FQDN Tags in a subscription.
//
// Generated from API version 2022-09-01
//   - options - AzureFirewallFqdnTagsClientListAllOptions contains the optional parameters for the AzureFirewallFqdnTagsClient.NewListAllPager
//     method.
func (client *AzureFirewallFqdnTagsClient) NewListAllPager(options *AzureFirewallFqdnTagsClientListAllOptions) *runtime.Pager[AzureFirewallFqdnTagsClientListAllResponse] {
	return runtime.NewPager(runtime.PagingHandler[AzureFirewallFqdnTagsClientListAllResponse]{
		More: func(page AzureFirewallFqdnTagsClientListAllResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AzureFirewallFqdnTagsClientListAllResponse) (AzureFirewallFqdnTagsClientListAllResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AzureFirewallFqdnTagsClient.NewListAllPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listAllCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return AzureFirewallFqdnTagsClientListAllResponse{}, err
			}
			return client.listAllHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listAllCreateRequest creates the ListAll request.
func (client *AzureFirewallFqdnTagsClient) listAllCreateRequest(ctx context.Context, options *AzureFirewallFqdnTagsClientListAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/azureFirewallFqdnTags"
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

// listAllHandleResponse handles the ListAll response.
func (client *AzureFirewallFqdnTagsClient) listAllHandleResponse(resp *http.Response) (AzureFirewallFqdnTagsClientListAllResponse, error) {
	result := AzureFirewallFqdnTagsClientListAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureFirewallFqdnTagListResult); err != nil {
		return AzureFirewallFqdnTagsClientListAllResponse{}, err
	}
	return result, nil
}
