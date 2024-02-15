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

// DiskAccessesServer is a fake server for instances of the armcompute.DiskAccessesClient type.
type DiskAccessesServer struct {
	// BeginCreateOrUpdate is the fake for method DiskAccessesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, diskAccessName string, diskAccess armcompute.DiskAccess, options *armcompute.DiskAccessesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcompute.DiskAccessesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method DiskAccessesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, diskAccessName string, options *armcompute.DiskAccessesClientBeginDeleteOptions) (resp azfake.PollerResponder[armcompute.DiskAccessesClientDeleteResponse], errResp azfake.ErrorResponder)

	// BeginDeleteAPrivateEndpointConnection is the fake for method DiskAccessesClient.BeginDeleteAPrivateEndpointConnection
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDeleteAPrivateEndpointConnection func(ctx context.Context, resourceGroupName string, diskAccessName string, privateEndpointConnectionName string, options *armcompute.DiskAccessesClientBeginDeleteAPrivateEndpointConnectionOptions) (resp azfake.PollerResponder[armcompute.DiskAccessesClientDeleteAPrivateEndpointConnectionResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method DiskAccessesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, diskAccessName string, options *armcompute.DiskAccessesClientGetOptions) (resp azfake.Responder[armcompute.DiskAccessesClientGetResponse], errResp azfake.ErrorResponder)

	// GetAPrivateEndpointConnection is the fake for method DiskAccessesClient.GetAPrivateEndpointConnection
	// HTTP status codes to indicate success: http.StatusOK
	GetAPrivateEndpointConnection func(ctx context.Context, resourceGroupName string, diskAccessName string, privateEndpointConnectionName string, options *armcompute.DiskAccessesClientGetAPrivateEndpointConnectionOptions) (resp azfake.Responder[armcompute.DiskAccessesClientGetAPrivateEndpointConnectionResponse], errResp azfake.ErrorResponder)

	// GetPrivateLinkResources is the fake for method DiskAccessesClient.GetPrivateLinkResources
	// HTTP status codes to indicate success: http.StatusOK
	GetPrivateLinkResources func(ctx context.Context, resourceGroupName string, diskAccessName string, options *armcompute.DiskAccessesClientGetPrivateLinkResourcesOptions) (resp azfake.Responder[armcompute.DiskAccessesClientGetPrivateLinkResourcesResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method DiskAccessesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armcompute.DiskAccessesClientListOptions) (resp azfake.PagerResponder[armcompute.DiskAccessesClientListResponse])

	// NewListByResourceGroupPager is the fake for method DiskAccessesClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armcompute.DiskAccessesClientListByResourceGroupOptions) (resp azfake.PagerResponder[armcompute.DiskAccessesClientListByResourceGroupResponse])

	// NewListPrivateEndpointConnectionsPager is the fake for method DiskAccessesClient.NewListPrivateEndpointConnectionsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPrivateEndpointConnectionsPager func(resourceGroupName string, diskAccessName string, options *armcompute.DiskAccessesClientListPrivateEndpointConnectionsOptions) (resp azfake.PagerResponder[armcompute.DiskAccessesClientListPrivateEndpointConnectionsResponse])

	// BeginUpdate is the fake for method DiskAccessesClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, diskAccessName string, diskAccess armcompute.DiskAccessUpdate, options *armcompute.DiskAccessesClientBeginUpdateOptions) (resp azfake.PollerResponder[armcompute.DiskAccessesClientUpdateResponse], errResp azfake.ErrorResponder)

	// BeginUpdateAPrivateEndpointConnection is the fake for method DiskAccessesClient.BeginUpdateAPrivateEndpointConnection
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdateAPrivateEndpointConnection func(ctx context.Context, resourceGroupName string, diskAccessName string, privateEndpointConnectionName string, privateEndpointConnection armcompute.PrivateEndpointConnection, options *armcompute.DiskAccessesClientBeginUpdateAPrivateEndpointConnectionOptions) (resp azfake.PollerResponder[armcompute.DiskAccessesClientUpdateAPrivateEndpointConnectionResponse], errResp azfake.ErrorResponder)
}

// NewDiskAccessesServerTransport creates a new instance of DiskAccessesServerTransport with the provided implementation.
// The returned DiskAccessesServerTransport instance is connected to an instance of armcompute.DiskAccessesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDiskAccessesServerTransport(srv *DiskAccessesServer) *DiskAccessesServerTransport {
	return &DiskAccessesServerTransport{
		srv:                                    srv,
		beginCreateOrUpdate:                    newTracker[azfake.PollerResponder[armcompute.DiskAccessesClientCreateOrUpdateResponse]](),
		beginDelete:                            newTracker[azfake.PollerResponder[armcompute.DiskAccessesClientDeleteResponse]](),
		beginDeleteAPrivateEndpointConnection:  newTracker[azfake.PollerResponder[armcompute.DiskAccessesClientDeleteAPrivateEndpointConnectionResponse]](),
		newListPager:                           newTracker[azfake.PagerResponder[armcompute.DiskAccessesClientListResponse]](),
		newListByResourceGroupPager:            newTracker[azfake.PagerResponder[armcompute.DiskAccessesClientListByResourceGroupResponse]](),
		newListPrivateEndpointConnectionsPager: newTracker[azfake.PagerResponder[armcompute.DiskAccessesClientListPrivateEndpointConnectionsResponse]](),
		beginUpdate:                            newTracker[azfake.PollerResponder[armcompute.DiskAccessesClientUpdateResponse]](),
		beginUpdateAPrivateEndpointConnection:  newTracker[azfake.PollerResponder[armcompute.DiskAccessesClientUpdateAPrivateEndpointConnectionResponse]](),
	}
}

// DiskAccessesServerTransport connects instances of armcompute.DiskAccessesClient to instances of DiskAccessesServer.
// Don't use this type directly, use NewDiskAccessesServerTransport instead.
type DiskAccessesServerTransport struct {
	srv                                    *DiskAccessesServer
	beginCreateOrUpdate                    *tracker[azfake.PollerResponder[armcompute.DiskAccessesClientCreateOrUpdateResponse]]
	beginDelete                            *tracker[azfake.PollerResponder[armcompute.DiskAccessesClientDeleteResponse]]
	beginDeleteAPrivateEndpointConnection  *tracker[azfake.PollerResponder[armcompute.DiskAccessesClientDeleteAPrivateEndpointConnectionResponse]]
	newListPager                           *tracker[azfake.PagerResponder[armcompute.DiskAccessesClientListResponse]]
	newListByResourceGroupPager            *tracker[azfake.PagerResponder[armcompute.DiskAccessesClientListByResourceGroupResponse]]
	newListPrivateEndpointConnectionsPager *tracker[azfake.PagerResponder[armcompute.DiskAccessesClientListPrivateEndpointConnectionsResponse]]
	beginUpdate                            *tracker[azfake.PollerResponder[armcompute.DiskAccessesClientUpdateResponse]]
	beginUpdateAPrivateEndpointConnection  *tracker[azfake.PollerResponder[armcompute.DiskAccessesClientUpdateAPrivateEndpointConnectionResponse]]
}

// Do implements the policy.Transporter interface for DiskAccessesServerTransport.
func (d *DiskAccessesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "DiskAccessesClient.BeginCreateOrUpdate":
		resp, err = d.dispatchBeginCreateOrUpdate(req)
	case "DiskAccessesClient.BeginDelete":
		resp, err = d.dispatchBeginDelete(req)
	case "DiskAccessesClient.BeginDeleteAPrivateEndpointConnection":
		resp, err = d.dispatchBeginDeleteAPrivateEndpointConnection(req)
	case "DiskAccessesClient.Get":
		resp, err = d.dispatchGet(req)
	case "DiskAccessesClient.GetAPrivateEndpointConnection":
		resp, err = d.dispatchGetAPrivateEndpointConnection(req)
	case "DiskAccessesClient.GetPrivateLinkResources":
		resp, err = d.dispatchGetPrivateLinkResources(req)
	case "DiskAccessesClient.NewListPager":
		resp, err = d.dispatchNewListPager(req)
	case "DiskAccessesClient.NewListByResourceGroupPager":
		resp, err = d.dispatchNewListByResourceGroupPager(req)
	case "DiskAccessesClient.NewListPrivateEndpointConnectionsPager":
		resp, err = d.dispatchNewListPrivateEndpointConnectionsPager(req)
	case "DiskAccessesClient.BeginUpdate":
		resp, err = d.dispatchBeginUpdate(req)
	case "DiskAccessesClient.BeginUpdateAPrivateEndpointConnection":
		resp, err = d.dispatchBeginUpdateAPrivateEndpointConnection(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DiskAccessesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if d.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := d.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/diskAccesses/(?P<diskAccessName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.DiskAccess](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		diskAccessNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("diskAccessName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, diskAccessNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		d.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		d.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		d.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (d *DiskAccessesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if d.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := d.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/diskAccesses/(?P<diskAccessName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		diskAccessNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("diskAccessName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginDelete(req.Context(), resourceGroupNameParam, diskAccessNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		d.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		d.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		d.beginDelete.remove(req)
	}

	return resp, nil
}

func (d *DiskAccessesServerTransport) dispatchBeginDeleteAPrivateEndpointConnection(req *http.Request) (*http.Response, error) {
	if d.srv.BeginDeleteAPrivateEndpointConnection == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDeleteAPrivateEndpointConnection not implemented")}
	}
	beginDeleteAPrivateEndpointConnection := d.beginDeleteAPrivateEndpointConnection.get(req)
	if beginDeleteAPrivateEndpointConnection == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/diskAccesses/(?P<diskAccessName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<privateEndpointConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		diskAccessNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("diskAccessName")])
		if err != nil {
			return nil, err
		}
		privateEndpointConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointConnectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginDeleteAPrivateEndpointConnection(req.Context(), resourceGroupNameParam, diskAccessNameParam, privateEndpointConnectionNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDeleteAPrivateEndpointConnection = &respr
		d.beginDeleteAPrivateEndpointConnection.add(req, beginDeleteAPrivateEndpointConnection)
	}

	resp, err := server.PollerResponderNext(beginDeleteAPrivateEndpointConnection, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		d.beginDeleteAPrivateEndpointConnection.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDeleteAPrivateEndpointConnection) {
		d.beginDeleteAPrivateEndpointConnection.remove(req)
	}

	return resp, nil
}

func (d *DiskAccessesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/diskAccesses/(?P<diskAccessName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	diskAccessNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("diskAccessName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Get(req.Context(), resourceGroupNameParam, diskAccessNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DiskAccess, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DiskAccessesServerTransport) dispatchGetAPrivateEndpointConnection(req *http.Request) (*http.Response, error) {
	if d.srv.GetAPrivateEndpointConnection == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetAPrivateEndpointConnection not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/diskAccesses/(?P<diskAccessName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<privateEndpointConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	diskAccessNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("diskAccessName")])
	if err != nil {
		return nil, err
	}
	privateEndpointConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.GetAPrivateEndpointConnection(req.Context(), resourceGroupNameParam, diskAccessNameParam, privateEndpointConnectionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PrivateEndpointConnection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DiskAccessesServerTransport) dispatchGetPrivateLinkResources(req *http.Request) (*http.Response, error) {
	if d.srv.GetPrivateLinkResources == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetPrivateLinkResources not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/diskAccesses/(?P<diskAccessName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateLinkResources`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	diskAccessNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("diskAccessName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.GetPrivateLinkResources(req.Context(), resourceGroupNameParam, diskAccessNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PrivateLinkResourceListResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DiskAccessesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := d.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/diskAccesses`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := d.srv.NewListPager(nil)
		newListPager = &resp
		d.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armcompute.DiskAccessesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		d.newListPager.remove(req)
	}
	return resp, nil
}

func (d *DiskAccessesServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := d.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/diskAccesses`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := d.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		d.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armcompute.DiskAccessesClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		d.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (d *DiskAccessesServerTransport) dispatchNewListPrivateEndpointConnectionsPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListPrivateEndpointConnectionsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPrivateEndpointConnectionsPager not implemented")}
	}
	newListPrivateEndpointConnectionsPager := d.newListPrivateEndpointConnectionsPager.get(req)
	if newListPrivateEndpointConnectionsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/diskAccesses/(?P<diskAccessName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		diskAccessNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("diskAccessName")])
		if err != nil {
			return nil, err
		}
		resp := d.srv.NewListPrivateEndpointConnectionsPager(resourceGroupNameParam, diskAccessNameParam, nil)
		newListPrivateEndpointConnectionsPager = &resp
		d.newListPrivateEndpointConnectionsPager.add(req, newListPrivateEndpointConnectionsPager)
		server.PagerResponderInjectNextLinks(newListPrivateEndpointConnectionsPager, req, func(page *armcompute.DiskAccessesClientListPrivateEndpointConnectionsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPrivateEndpointConnectionsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListPrivateEndpointConnectionsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPrivateEndpointConnectionsPager) {
		d.newListPrivateEndpointConnectionsPager.remove(req)
	}
	return resp, nil
}

func (d *DiskAccessesServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if d.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := d.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/diskAccesses/(?P<diskAccessName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.DiskAccessUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		diskAccessNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("diskAccessName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginUpdate(req.Context(), resourceGroupNameParam, diskAccessNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		d.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		d.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		d.beginUpdate.remove(req)
	}

	return resp, nil
}

func (d *DiskAccessesServerTransport) dispatchBeginUpdateAPrivateEndpointConnection(req *http.Request) (*http.Response, error) {
	if d.srv.BeginUpdateAPrivateEndpointConnection == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdateAPrivateEndpointConnection not implemented")}
	}
	beginUpdateAPrivateEndpointConnection := d.beginUpdateAPrivateEndpointConnection.get(req)
	if beginUpdateAPrivateEndpointConnection == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/diskAccesses/(?P<diskAccessName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<privateEndpointConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.PrivateEndpointConnection](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		diskAccessNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("diskAccessName")])
		if err != nil {
			return nil, err
		}
		privateEndpointConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointConnectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginUpdateAPrivateEndpointConnection(req.Context(), resourceGroupNameParam, diskAccessNameParam, privateEndpointConnectionNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdateAPrivateEndpointConnection = &respr
		d.beginUpdateAPrivateEndpointConnection.add(req, beginUpdateAPrivateEndpointConnection)
	}

	resp, err := server.PollerResponderNext(beginUpdateAPrivateEndpointConnection, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		d.beginUpdateAPrivateEndpointConnection.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdateAPrivateEndpointConnection) {
		d.beginUpdateAPrivateEndpointConnection.remove(req)
	}

	return resp, nil
}
