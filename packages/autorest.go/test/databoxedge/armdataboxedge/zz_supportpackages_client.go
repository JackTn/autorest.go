// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdataboxedge

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

// SupportPackagesClient contains the methods for the SupportPackages group.
// Don't use this type directly, use NewSupportPackagesClient() instead.
type SupportPackagesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSupportPackagesClient creates a new instance of SupportPackagesClient with the specified values.
//   - subscriptionID - The subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSupportPackagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SupportPackagesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SupportPackagesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginTriggerSupportPackage - Triggers support package on the device
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01
//   - deviceName - The device name.
//   - resourceGroupName - The resource group name.
//   - triggerSupportPackageRequest - The trigger support package request object
//   - options - SupportPackagesClientBeginTriggerSupportPackageOptions contains the optional parameters for the SupportPackagesClient.BeginTriggerSupportPackage
//     method.
func (client *SupportPackagesClient) BeginTriggerSupportPackage(ctx context.Context, deviceName string, resourceGroupName string, triggerSupportPackageRequest TriggerSupportPackageRequest, options *SupportPackagesClientBeginTriggerSupportPackageOptions) (*runtime.Poller[SupportPackagesClientTriggerSupportPackageResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.triggerSupportPackage(ctx, deviceName, resourceGroupName, triggerSupportPackageRequest, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[SupportPackagesClientTriggerSupportPackageResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[SupportPackagesClientTriggerSupportPackageResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// TriggerSupportPackage - Triggers support package on the device
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01
func (client *SupportPackagesClient) triggerSupportPackage(ctx context.Context, deviceName string, resourceGroupName string, triggerSupportPackageRequest TriggerSupportPackageRequest, options *SupportPackagesClientBeginTriggerSupportPackageOptions) (*http.Response, error) {
	var err error
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SupportPackagesClient.BeginTriggerSupportPackage")
	req, err := client.triggerSupportPackageCreateRequest(ctx, deviceName, resourceGroupName, triggerSupportPackageRequest, options)
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

// triggerSupportPackageCreateRequest creates the TriggerSupportPackage request.
func (client *SupportPackagesClient) triggerSupportPackageCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, triggerSupportPackageRequest TriggerSupportPackageRequest, _ *SupportPackagesClientBeginTriggerSupportPackageOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/triggerSupportPackage"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, triggerSupportPackageRequest); err != nil {
		return nil, err
	}
	return req, nil
}
