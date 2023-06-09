//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"armnetwork"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"net/http"
	"net/url"
	"regexp"
)

// PublicIPAddressesServer is a fake server for instances of the armnetwork.PublicIPAddressesClient type.
type PublicIPAddressesServer struct {
	// BeginCreateOrUpdate is the fake for method PublicIPAddressesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, publicIPAddressName string, parameters armnetwork.PublicIPAddress, options *armnetwork.PublicIPAddressesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.PublicIPAddressesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDdosProtectionStatus is the fake for method PublicIPAddressesClient.BeginDdosProtectionStatus
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginDdosProtectionStatus func(ctx context.Context, resourceGroupName string, publicIPAddressName string, options *armnetwork.PublicIPAddressesClientBeginDdosProtectionStatusOptions) (resp azfake.PollerResponder[armnetwork.PublicIPAddressesClientDdosProtectionStatusResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method PublicIPAddressesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, publicIPAddressName string, options *armnetwork.PublicIPAddressesClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.PublicIPAddressesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method PublicIPAddressesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, publicIPAddressName string, options *armnetwork.PublicIPAddressesClientGetOptions) (resp azfake.Responder[armnetwork.PublicIPAddressesClientGetResponse], errResp azfake.ErrorResponder)

	// GetCloudServicePublicIPAddress is the fake for method PublicIPAddressesClient.GetCloudServicePublicIPAddress
	// HTTP status codes to indicate success: http.StatusOK
	GetCloudServicePublicIPAddress func(ctx context.Context, resourceGroupName string, cloudServiceName string, roleInstanceName string, networkInterfaceName string, ipConfigurationName string, publicIPAddressName string, options *armnetwork.PublicIPAddressesClientGetCloudServicePublicIPAddressOptions) (resp azfake.Responder[armnetwork.PublicIPAddressesClientGetCloudServicePublicIPAddressResponse], errResp azfake.ErrorResponder)

	// GetVirtualMachineScaleSetPublicIPAddress is the fake for method PublicIPAddressesClient.GetVirtualMachineScaleSetPublicIPAddress
	// HTTP status codes to indicate success: http.StatusOK
	GetVirtualMachineScaleSetPublicIPAddress func(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, ipConfigurationName string, publicIPAddressName string, options *armnetwork.PublicIPAddressesClientGetVirtualMachineScaleSetPublicIPAddressOptions) (resp azfake.Responder[armnetwork.PublicIPAddressesClientGetVirtualMachineScaleSetPublicIPAddressResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method PublicIPAddressesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, options *armnetwork.PublicIPAddressesClientListOptions) (resp azfake.PagerResponder[armnetwork.PublicIPAddressesClientListResponse])

	// NewListAllPager is the fake for method PublicIPAddressesClient.NewListAllPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListAllPager func(options *armnetwork.PublicIPAddressesClientListAllOptions) (resp azfake.PagerResponder[armnetwork.PublicIPAddressesClientListAllResponse])

	// NewListCloudServicePublicIPAddressesPager is the fake for method PublicIPAddressesClient.NewListCloudServicePublicIPAddressesPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListCloudServicePublicIPAddressesPager func(resourceGroupName string, cloudServiceName string, options *armnetwork.PublicIPAddressesClientListCloudServicePublicIPAddressesOptions) (resp azfake.PagerResponder[armnetwork.PublicIPAddressesClientListCloudServicePublicIPAddressesResponse])

	// NewListCloudServiceRoleInstancePublicIPAddressesPager is the fake for method PublicIPAddressesClient.NewListCloudServiceRoleInstancePublicIPAddressesPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListCloudServiceRoleInstancePublicIPAddressesPager func(resourceGroupName string, cloudServiceName string, roleInstanceName string, networkInterfaceName string, ipConfigurationName string, options *armnetwork.PublicIPAddressesClientListCloudServiceRoleInstancePublicIPAddressesOptions) (resp azfake.PagerResponder[armnetwork.PublicIPAddressesClientListCloudServiceRoleInstancePublicIPAddressesResponse])

	// NewListVirtualMachineScaleSetPublicIPAddressesPager is the fake for method PublicIPAddressesClient.NewListVirtualMachineScaleSetPublicIPAddressesPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListVirtualMachineScaleSetPublicIPAddressesPager func(resourceGroupName string, virtualMachineScaleSetName string, options *armnetwork.PublicIPAddressesClientListVirtualMachineScaleSetPublicIPAddressesOptions) (resp azfake.PagerResponder[armnetwork.PublicIPAddressesClientListVirtualMachineScaleSetPublicIPAddressesResponse])

	// NewListVirtualMachineScaleSetVMPublicIPAddressesPager is the fake for method PublicIPAddressesClient.NewListVirtualMachineScaleSetVMPublicIPAddressesPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListVirtualMachineScaleSetVMPublicIPAddressesPager func(resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, ipConfigurationName string, options *armnetwork.PublicIPAddressesClientListVirtualMachineScaleSetVMPublicIPAddressesOptions) (resp azfake.PagerResponder[armnetwork.PublicIPAddressesClientListVirtualMachineScaleSetVMPublicIPAddressesResponse])

	// UpdateTags is the fake for method PublicIPAddressesClient.UpdateTags
	// HTTP status codes to indicate success: http.StatusOK
	UpdateTags func(ctx context.Context, resourceGroupName string, publicIPAddressName string, parameters armnetwork.TagsObject, options *armnetwork.PublicIPAddressesClientUpdateTagsOptions) (resp azfake.Responder[armnetwork.PublicIPAddressesClientUpdateTagsResponse], errResp azfake.ErrorResponder)
}

// NewPublicIPAddressesServerTransport creates a new instance of PublicIPAddressesServerTransport with the provided implementation.
// The returned PublicIPAddressesServerTransport instance is connected to an instance of armnetwork.PublicIPAddressesClient by way of the
// undefined.Transporter field.
func NewPublicIPAddressesServerTransport(srv *PublicIPAddressesServer) *PublicIPAddressesServerTransport {
	return &PublicIPAddressesServerTransport{srv: srv}
}

// PublicIPAddressesServerTransport connects instances of armnetwork.PublicIPAddressesClient to instances of PublicIPAddressesServer.
// Don't use this type directly, use NewPublicIPAddressesServerTransport instead.
type PublicIPAddressesServerTransport struct {
	srv                                                   *PublicIPAddressesServer
	beginCreateOrUpdate                                   *azfake.PollerResponder[armnetwork.PublicIPAddressesClientCreateOrUpdateResponse]
	beginDdosProtectionStatus                             *azfake.PollerResponder[armnetwork.PublicIPAddressesClientDdosProtectionStatusResponse]
	beginDelete                                           *azfake.PollerResponder[armnetwork.PublicIPAddressesClientDeleteResponse]
	newListPager                                          *azfake.PagerResponder[armnetwork.PublicIPAddressesClientListResponse]
	newListAllPager                                       *azfake.PagerResponder[armnetwork.PublicIPAddressesClientListAllResponse]
	newListCloudServicePublicIPAddressesPager             *azfake.PagerResponder[armnetwork.PublicIPAddressesClientListCloudServicePublicIPAddressesResponse]
	newListCloudServiceRoleInstancePublicIPAddressesPager *azfake.PagerResponder[armnetwork.PublicIPAddressesClientListCloudServiceRoleInstancePublicIPAddressesResponse]
	newListVirtualMachineScaleSetPublicIPAddressesPager   *azfake.PagerResponder[armnetwork.PublicIPAddressesClientListVirtualMachineScaleSetPublicIPAddressesResponse]
	newListVirtualMachineScaleSetVMPublicIPAddressesPager *azfake.PagerResponder[armnetwork.PublicIPAddressesClientListVirtualMachineScaleSetVMPublicIPAddressesResponse]
}

// Do implements the policy.Transporter interface for PublicIPAddressesServerTransport.
func (p *PublicIPAddressesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PublicIPAddressesClient.BeginCreateOrUpdate":
		resp, err = p.dispatchBeginCreateOrUpdate(req)
	case "PublicIPAddressesClient.BeginDdosProtectionStatus":
		resp, err = p.dispatchBeginDdosProtectionStatus(req)
	case "PublicIPAddressesClient.BeginDelete":
		resp, err = p.dispatchBeginDelete(req)
	case "PublicIPAddressesClient.Get":
		resp, err = p.dispatchGet(req)
	case "PublicIPAddressesClient.GetCloudServicePublicIPAddress":
		resp, err = p.dispatchGetCloudServicePublicIPAddress(req)
	case "PublicIPAddressesClient.GetVirtualMachineScaleSetPublicIPAddress":
		resp, err = p.dispatchGetVirtualMachineScaleSetPublicIPAddress(req)
	case "PublicIPAddressesClient.NewListPager":
		resp, err = p.dispatchNewListPager(req)
	case "PublicIPAddressesClient.NewListAllPager":
		resp, err = p.dispatchNewListAllPager(req)
	case "PublicIPAddressesClient.NewListCloudServicePublicIPAddressesPager":
		resp, err = p.dispatchNewListCloudServicePublicIPAddressesPager(req)
	case "PublicIPAddressesClient.NewListCloudServiceRoleInstancePublicIPAddressesPager":
		resp, err = p.dispatchNewListCloudServiceRoleInstancePublicIPAddressesPager(req)
	case "PublicIPAddressesClient.NewListVirtualMachineScaleSetPublicIPAddressesPager":
		resp, err = p.dispatchNewListVirtualMachineScaleSetPublicIPAddressesPager(req)
	case "PublicIPAddressesClient.NewListVirtualMachineScaleSetVMPublicIPAddressesPager":
		resp, err = p.dispatchNewListVirtualMachineScaleSetVMPublicIPAddressesPager(req)
	case "PublicIPAddressesClient.UpdateTags":
		resp, err = p.dispatchUpdateTags(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PublicIPAddressesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	if p.beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/publicIPAddresses/(?P<publicIpAddressName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.PublicIPAddress](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		publicIPAddressNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("publicIpAddressName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameUnescaped, publicIPAddressNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		p.beginCreateOrUpdate = &respr
	}

	resp, err := server.PollerResponderNext(p.beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(p.beginCreateOrUpdate) {
		p.beginCreateOrUpdate = nil
	}

	return resp, nil
}

func (p *PublicIPAddressesServerTransport) dispatchBeginDdosProtectionStatus(req *http.Request) (*http.Response, error) {
	if p.srv.BeginDdosProtectionStatus == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDdosProtectionStatus not implemented")}
	}
	if p.beginDdosProtectionStatus == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/publicIPAddresses/(?P<publicIpAddressName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ddosProtectionStatus`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		publicIPAddressNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("publicIpAddressName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginDdosProtectionStatus(req.Context(), resourceGroupNameUnescaped, publicIPAddressNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		p.beginDdosProtectionStatus = &respr
	}

	resp, err := server.PollerResponderNext(p.beginDdosProtectionStatus, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(p.beginDdosProtectionStatus) {
		p.beginDdosProtectionStatus = nil
	}

	return resp, nil
}

func (p *PublicIPAddressesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if p.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	if p.beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/publicIPAddresses/(?P<publicIpAddressName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		publicIPAddressNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("publicIpAddressName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginDelete(req.Context(), resourceGroupNameUnescaped, publicIPAddressNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		p.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(p.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(p.beginDelete) {
		p.beginDelete = nil
	}

	return resp, nil
}

func (p *PublicIPAddressesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/publicIPAddresses/(?P<publicIpAddressName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	publicIPAddressNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("publicIpAddressName")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
	var options *armnetwork.PublicIPAddressesClientGetOptions
	if expandParam != nil {
		options = &armnetwork.PublicIPAddressesClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameUnescaped, publicIPAddressNameUnescaped, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PublicIPAddress, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PublicIPAddressesServerTransport) dispatchGetCloudServicePublicIPAddress(req *http.Request) (*http.Response, error) {
	if p.srv.GetCloudServicePublicIPAddress == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetCloudServicePublicIPAddress not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/roleInstances/(?P<roleInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkInterfaces/(?P<networkInterfaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ipconfigurations/(?P<ipConfigurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/publicipaddresses/(?P<publicIpAddressName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 7 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	cloudServiceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("cloudServiceName")])
	if err != nil {
		return nil, err
	}
	roleInstanceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("roleInstanceName")])
	if err != nil {
		return nil, err
	}
	networkInterfaceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("networkInterfaceName")])
	if err != nil {
		return nil, err
	}
	ipConfigurationNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("ipConfigurationName")])
	if err != nil {
		return nil, err
	}
	publicIPAddressNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("publicIpAddressName")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
	var options *armnetwork.PublicIPAddressesClientGetCloudServicePublicIPAddressOptions
	if expandParam != nil {
		options = &armnetwork.PublicIPAddressesClientGetCloudServicePublicIPAddressOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := p.srv.GetCloudServicePublicIPAddress(req.Context(), resourceGroupNameUnescaped, cloudServiceNameUnescaped, roleInstanceNameUnescaped, networkInterfaceNameUnescaped, ipConfigurationNameUnescaped, publicIPAddressNameUnescaped, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PublicIPAddress, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PublicIPAddressesServerTransport) dispatchGetVirtualMachineScaleSetPublicIPAddress(req *http.Request) (*http.Response, error) {
	if p.srv.GetVirtualMachineScaleSetPublicIPAddress == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetVirtualMachineScaleSetPublicIPAddress not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/virtualMachineScaleSets/(?P<virtualMachineScaleSetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualMachines/(?P<virtualmachineIndex>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkInterfaces/(?P<networkInterfaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ipconfigurations/(?P<ipConfigurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/publicipaddresses/(?P<publicIpAddressName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 7 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	virtualMachineScaleSetNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualMachineScaleSetName")])
	if err != nil {
		return nil, err
	}
	virtualmachineIndexUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualmachineIndex")])
	if err != nil {
		return nil, err
	}
	networkInterfaceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("networkInterfaceName")])
	if err != nil {
		return nil, err
	}
	ipConfigurationNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("ipConfigurationName")])
	if err != nil {
		return nil, err
	}
	publicIPAddressNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("publicIpAddressName")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
	var options *armnetwork.PublicIPAddressesClientGetVirtualMachineScaleSetPublicIPAddressOptions
	if expandParam != nil {
		options = &armnetwork.PublicIPAddressesClientGetVirtualMachineScaleSetPublicIPAddressOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := p.srv.GetVirtualMachineScaleSetPublicIPAddress(req.Context(), resourceGroupNameUnescaped, virtualMachineScaleSetNameUnescaped, virtualmachineIndexUnescaped, networkInterfaceNameUnescaped, ipConfigurationNameUnescaped, publicIPAddressNameUnescaped, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PublicIPAddress, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PublicIPAddressesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	if p.newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/publicIPAddresses`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListPager(resourceGroupNameUnescaped, nil)
		p.newListPager = &resp
		server.PagerResponderInjectNextLinks(p.newListPager, req, func(page *armnetwork.PublicIPAddressesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(p.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(p.newListPager) {
		p.newListPager = nil
	}
	return resp, nil
}

func (p *PublicIPAddressesServerTransport) dispatchNewListAllPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListAllPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListAllPager not implemented")}
	}
	if p.newListAllPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/publicIPAddresses`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := p.srv.NewListAllPager(nil)
		p.newListAllPager = &resp
		server.PagerResponderInjectNextLinks(p.newListAllPager, req, func(page *armnetwork.PublicIPAddressesClientListAllResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(p.newListAllPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(p.newListAllPager) {
		p.newListAllPager = nil
	}
	return resp, nil
}

func (p *PublicIPAddressesServerTransport) dispatchNewListCloudServicePublicIPAddressesPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListCloudServicePublicIPAddressesPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListCloudServicePublicIPAddressesPager not implemented")}
	}
	if p.newListCloudServicePublicIPAddressesPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/publicipaddresses`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		cloudServiceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("cloudServiceName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListCloudServicePublicIPAddressesPager(resourceGroupNameUnescaped, cloudServiceNameUnescaped, nil)
		p.newListCloudServicePublicIPAddressesPager = &resp
		server.PagerResponderInjectNextLinks(p.newListCloudServicePublicIPAddressesPager, req, func(page *armnetwork.PublicIPAddressesClientListCloudServicePublicIPAddressesResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(p.newListCloudServicePublicIPAddressesPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(p.newListCloudServicePublicIPAddressesPager) {
		p.newListCloudServicePublicIPAddressesPager = nil
	}
	return resp, nil
}

func (p *PublicIPAddressesServerTransport) dispatchNewListCloudServiceRoleInstancePublicIPAddressesPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListCloudServiceRoleInstancePublicIPAddressesPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListCloudServiceRoleInstancePublicIPAddressesPager not implemented")}
	}
	if p.newListCloudServiceRoleInstancePublicIPAddressesPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/roleInstances/(?P<roleInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkInterfaces/(?P<networkInterfaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ipconfigurations/(?P<ipConfigurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/publicipaddresses`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 6 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		cloudServiceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("cloudServiceName")])
		if err != nil {
			return nil, err
		}
		roleInstanceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("roleInstanceName")])
		if err != nil {
			return nil, err
		}
		networkInterfaceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("networkInterfaceName")])
		if err != nil {
			return nil, err
		}
		ipConfigurationNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("ipConfigurationName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListCloudServiceRoleInstancePublicIPAddressesPager(resourceGroupNameUnescaped, cloudServiceNameUnescaped, roleInstanceNameUnescaped, networkInterfaceNameUnescaped, ipConfigurationNameUnescaped, nil)
		p.newListCloudServiceRoleInstancePublicIPAddressesPager = &resp
		server.PagerResponderInjectNextLinks(p.newListCloudServiceRoleInstancePublicIPAddressesPager, req, func(page *armnetwork.PublicIPAddressesClientListCloudServiceRoleInstancePublicIPAddressesResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(p.newListCloudServiceRoleInstancePublicIPAddressesPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(p.newListCloudServiceRoleInstancePublicIPAddressesPager) {
		p.newListCloudServiceRoleInstancePublicIPAddressesPager = nil
	}
	return resp, nil
}

func (p *PublicIPAddressesServerTransport) dispatchNewListVirtualMachineScaleSetPublicIPAddressesPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListVirtualMachineScaleSetPublicIPAddressesPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListVirtualMachineScaleSetPublicIPAddressesPager not implemented")}
	}
	if p.newListVirtualMachineScaleSetPublicIPAddressesPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/virtualMachineScaleSets/(?P<virtualMachineScaleSetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/publicipaddresses`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualMachineScaleSetNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualMachineScaleSetName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListVirtualMachineScaleSetPublicIPAddressesPager(resourceGroupNameUnescaped, virtualMachineScaleSetNameUnescaped, nil)
		p.newListVirtualMachineScaleSetPublicIPAddressesPager = &resp
		server.PagerResponderInjectNextLinks(p.newListVirtualMachineScaleSetPublicIPAddressesPager, req, func(page *armnetwork.PublicIPAddressesClientListVirtualMachineScaleSetPublicIPAddressesResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(p.newListVirtualMachineScaleSetPublicIPAddressesPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(p.newListVirtualMachineScaleSetPublicIPAddressesPager) {
		p.newListVirtualMachineScaleSetPublicIPAddressesPager = nil
	}
	return resp, nil
}

func (p *PublicIPAddressesServerTransport) dispatchNewListVirtualMachineScaleSetVMPublicIPAddressesPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListVirtualMachineScaleSetVMPublicIPAddressesPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListVirtualMachineScaleSetVMPublicIPAddressesPager not implemented")}
	}
	if p.newListVirtualMachineScaleSetVMPublicIPAddressesPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/virtualMachineScaleSets/(?P<virtualMachineScaleSetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualMachines/(?P<virtualmachineIndex>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkInterfaces/(?P<networkInterfaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ipconfigurations/(?P<ipConfigurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/publicipaddresses`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 6 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualMachineScaleSetNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualMachineScaleSetName")])
		if err != nil {
			return nil, err
		}
		virtualmachineIndexUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualmachineIndex")])
		if err != nil {
			return nil, err
		}
		networkInterfaceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("networkInterfaceName")])
		if err != nil {
			return nil, err
		}
		ipConfigurationNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("ipConfigurationName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListVirtualMachineScaleSetVMPublicIPAddressesPager(resourceGroupNameUnescaped, virtualMachineScaleSetNameUnescaped, virtualmachineIndexUnescaped, networkInterfaceNameUnescaped, ipConfigurationNameUnescaped, nil)
		p.newListVirtualMachineScaleSetVMPublicIPAddressesPager = &resp
		server.PagerResponderInjectNextLinks(p.newListVirtualMachineScaleSetVMPublicIPAddressesPager, req, func(page *armnetwork.PublicIPAddressesClientListVirtualMachineScaleSetVMPublicIPAddressesResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(p.newListVirtualMachineScaleSetVMPublicIPAddressesPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(p.newListVirtualMachineScaleSetVMPublicIPAddressesPager) {
		p.newListVirtualMachineScaleSetVMPublicIPAddressesPager = nil
	}
	return resp, nil
}

func (p *PublicIPAddressesServerTransport) dispatchUpdateTags(req *http.Request) (*http.Response, error) {
	if p.srv.UpdateTags == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateTags not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/publicIPAddresses/(?P<publicIpAddressName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetwork.TagsObject](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	publicIPAddressNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("publicIpAddressName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.UpdateTags(req.Context(), resourceGroupNameUnescaped, publicIPAddressNameUnescaped, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PublicIPAddress, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
