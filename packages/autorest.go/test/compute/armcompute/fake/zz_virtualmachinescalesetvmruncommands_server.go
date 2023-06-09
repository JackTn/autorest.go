//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"armcompute"
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

// VirtualMachineScaleSetVMRunCommandsServer is a fake server for instances of the armcompute.VirtualMachineScaleSetVMRunCommandsClient type.
type VirtualMachineScaleSetVMRunCommandsServer struct {
	// BeginCreateOrUpdate is the fake for method VirtualMachineScaleSetVMRunCommandsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, runCommandName string, runCommand armcompute.VirtualMachineRunCommand, options *armcompute.VirtualMachineScaleSetVMRunCommandsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcompute.VirtualMachineScaleSetVMRunCommandsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method VirtualMachineScaleSetVMRunCommandsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, runCommandName string, options *armcompute.VirtualMachineScaleSetVMRunCommandsClientBeginDeleteOptions) (resp azfake.PollerResponder[armcompute.VirtualMachineScaleSetVMRunCommandsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method VirtualMachineScaleSetVMRunCommandsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, runCommandName string, options *armcompute.VirtualMachineScaleSetVMRunCommandsClientGetOptions) (resp azfake.Responder[armcompute.VirtualMachineScaleSetVMRunCommandsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method VirtualMachineScaleSetVMRunCommandsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, vmScaleSetName string, instanceID string, options *armcompute.VirtualMachineScaleSetVMRunCommandsClientListOptions) (resp azfake.PagerResponder[armcompute.VirtualMachineScaleSetVMRunCommandsClientListResponse])

	// BeginUpdate is the fake for method VirtualMachineScaleSetVMRunCommandsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK
	BeginUpdate func(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, runCommandName string, runCommand armcompute.VirtualMachineRunCommandUpdate, options *armcompute.VirtualMachineScaleSetVMRunCommandsClientBeginUpdateOptions) (resp azfake.PollerResponder[armcompute.VirtualMachineScaleSetVMRunCommandsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewVirtualMachineScaleSetVMRunCommandsServerTransport creates a new instance of VirtualMachineScaleSetVMRunCommandsServerTransport with the provided implementation.
// The returned VirtualMachineScaleSetVMRunCommandsServerTransport instance is connected to an instance of armcompute.VirtualMachineScaleSetVMRunCommandsClient by way of the
// undefined.Transporter field.
func NewVirtualMachineScaleSetVMRunCommandsServerTransport(srv *VirtualMachineScaleSetVMRunCommandsServer) *VirtualMachineScaleSetVMRunCommandsServerTransport {
	return &VirtualMachineScaleSetVMRunCommandsServerTransport{srv: srv}
}

// VirtualMachineScaleSetVMRunCommandsServerTransport connects instances of armcompute.VirtualMachineScaleSetVMRunCommandsClient to instances of VirtualMachineScaleSetVMRunCommandsServer.
// Don't use this type directly, use NewVirtualMachineScaleSetVMRunCommandsServerTransport instead.
type VirtualMachineScaleSetVMRunCommandsServerTransport struct {
	srv                 *VirtualMachineScaleSetVMRunCommandsServer
	beginCreateOrUpdate *azfake.PollerResponder[armcompute.VirtualMachineScaleSetVMRunCommandsClientCreateOrUpdateResponse]
	beginDelete         *azfake.PollerResponder[armcompute.VirtualMachineScaleSetVMRunCommandsClientDeleteResponse]
	newListPager        *azfake.PagerResponder[armcompute.VirtualMachineScaleSetVMRunCommandsClientListResponse]
	beginUpdate         *azfake.PollerResponder[armcompute.VirtualMachineScaleSetVMRunCommandsClientUpdateResponse]
}

// Do implements the policy.Transporter interface for VirtualMachineScaleSetVMRunCommandsServerTransport.
func (v *VirtualMachineScaleSetVMRunCommandsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VirtualMachineScaleSetVMRunCommandsClient.BeginCreateOrUpdate":
		resp, err = v.dispatchBeginCreateOrUpdate(req)
	case "VirtualMachineScaleSetVMRunCommandsClient.BeginDelete":
		resp, err = v.dispatchBeginDelete(req)
	case "VirtualMachineScaleSetVMRunCommandsClient.Get":
		resp, err = v.dispatchGet(req)
	case "VirtualMachineScaleSetVMRunCommandsClient.NewListPager":
		resp, err = v.dispatchNewListPager(req)
	case "VirtualMachineScaleSetVMRunCommandsClient.BeginUpdate":
		resp, err = v.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VirtualMachineScaleSetVMRunCommandsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	if v.beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/virtualMachineScaleSets/(?P<vmScaleSetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualMachines/(?P<instanceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runCommands/(?P<runCommandName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.VirtualMachineRunCommand](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		vmScaleSetNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("vmScaleSetName")])
		if err != nil {
			return nil, err
		}
		instanceIDUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("instanceId")])
		if err != nil {
			return nil, err
		}
		runCommandNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("runCommandName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameUnescaped, vmScaleSetNameUnescaped, instanceIDUnescaped, runCommandNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		v.beginCreateOrUpdate = &respr
	}

	resp, err := server.PollerResponderNext(v.beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(v.beginCreateOrUpdate) {
		v.beginCreateOrUpdate = nil
	}

	return resp, nil
}

func (v *VirtualMachineScaleSetVMRunCommandsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if v.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	if v.beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/virtualMachineScaleSets/(?P<vmScaleSetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualMachines/(?P<instanceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runCommands/(?P<runCommandName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		vmScaleSetNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("vmScaleSetName")])
		if err != nil {
			return nil, err
		}
		instanceIDUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("instanceId")])
		if err != nil {
			return nil, err
		}
		runCommandNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("runCommandName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginDelete(req.Context(), resourceGroupNameUnescaped, vmScaleSetNameUnescaped, instanceIDUnescaped, runCommandNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		v.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(v.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(v.beginDelete) {
		v.beginDelete = nil
	}

	return resp, nil
}

func (v *VirtualMachineScaleSetVMRunCommandsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/virtualMachineScaleSets/(?P<vmScaleSetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualMachines/(?P<instanceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runCommands/(?P<runCommandName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	vmScaleSetNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("vmScaleSetName")])
	if err != nil {
		return nil, err
	}
	instanceIDUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("instanceId")])
	if err != nil {
		return nil, err
	}
	runCommandNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("runCommandName")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
	var options *armcompute.VirtualMachineScaleSetVMRunCommandsClientGetOptions
	if expandParam != nil {
		options = &armcompute.VirtualMachineScaleSetVMRunCommandsClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := v.srv.Get(req.Context(), resourceGroupNameUnescaped, vmScaleSetNameUnescaped, instanceIDUnescaped, runCommandNameUnescaped, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualMachineRunCommand, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualMachineScaleSetVMRunCommandsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	if v.newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/virtualMachineScaleSets/(?P<vmScaleSetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualMachines/(?P<instanceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runCommands`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		vmScaleSetNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("vmScaleSetName")])
		if err != nil {
			return nil, err
		}
		instanceIDUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("instanceId")])
		if err != nil {
			return nil, err
		}
		expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
		if err != nil {
			return nil, err
		}
		expandParam := getOptional(expandUnescaped)
		var options *armcompute.VirtualMachineScaleSetVMRunCommandsClientListOptions
		if expandParam != nil {
			options = &armcompute.VirtualMachineScaleSetVMRunCommandsClientListOptions{
				Expand: expandParam,
			}
		}
		resp := v.srv.NewListPager(resourceGroupNameUnescaped, vmScaleSetNameUnescaped, instanceIDUnescaped, options)
		v.newListPager = &resp
		server.PagerResponderInjectNextLinks(v.newListPager, req, func(page *armcompute.VirtualMachineScaleSetVMRunCommandsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(v.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(v.newListPager) {
		v.newListPager = nil
	}
	return resp, nil
}

func (v *VirtualMachineScaleSetVMRunCommandsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	if v.beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/virtualMachineScaleSets/(?P<vmScaleSetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualMachines/(?P<instanceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runCommands/(?P<runCommandName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.VirtualMachineRunCommandUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		vmScaleSetNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("vmScaleSetName")])
		if err != nil {
			return nil, err
		}
		instanceIDUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("instanceId")])
		if err != nil {
			return nil, err
		}
		runCommandNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("runCommandName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginUpdate(req.Context(), resourceGroupNameUnescaped, vmScaleSetNameUnescaped, instanceIDUnescaped, runCommandNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		v.beginUpdate = &respr
	}

	resp, err := server.PollerResponderNext(v.beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PollerResponderMore(v.beginUpdate) {
		v.beginUpdate = nil
	}

	return resp, nil
}
