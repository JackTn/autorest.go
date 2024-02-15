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

// VPNSitesConfigurationClient contains the methods for the VPNSitesConfiguration group.
// Don't use this type directly, use NewVPNSitesConfigurationClient() instead.
type VPNSitesConfigurationClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVPNSitesConfigurationClient creates a new instance of VPNSitesConfigurationClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVPNSitesConfigurationClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VPNSitesConfigurationClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VPNSitesConfigurationClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginDownload - Gives the sas-url to download the configurations for vpn-sites in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The resource group name.
//   - virtualWANName - The name of the VirtualWAN for which configuration of all vpn-sites is needed.
//   - request - Parameters supplied to download vpn-sites configuration.
//   - options - VPNSitesConfigurationClientBeginDownloadOptions contains the optional parameters for the VPNSitesConfigurationClient.BeginDownload
//     method.
func (client *VPNSitesConfigurationClient) BeginDownload(ctx context.Context, resourceGroupName string, virtualWANName string, request GetVPNSitesConfigurationRequest, options *VPNSitesConfigurationClientBeginDownloadOptions) (*runtime.Poller[VPNSitesConfigurationClientDownloadResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.download(ctx, resourceGroupName, virtualWANName, request, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VPNSitesConfigurationClientDownloadResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VPNSitesConfigurationClientDownloadResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Download - Gives the sas-url to download the configurations for vpn-sites in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *VPNSitesConfigurationClient) download(ctx context.Context, resourceGroupName string, virtualWANName string, request GetVPNSitesConfigurationRequest, options *VPNSitesConfigurationClientBeginDownloadOptions) (*http.Response, error) {
	var err error
	const operationName = "VPNSitesConfigurationClient.BeginDownload"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.downloadCreateRequest(ctx, resourceGroupName, virtualWANName, request, options)
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

// downloadCreateRequest creates the Download request.
func (client *VPNSitesConfigurationClient) downloadCreateRequest(ctx context.Context, resourceGroupName string, virtualWANName string, request GetVPNSitesConfigurationRequest, options *VPNSitesConfigurationClientBeginDownloadOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{virtualWANName}/vpnConfiguration"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualWANName == "" {
		return nil, errors.New("parameter virtualWANName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualWANName}", url.PathEscape(virtualWANName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, request); err != nil {
		return nil, err
	}
	return req, nil
}
