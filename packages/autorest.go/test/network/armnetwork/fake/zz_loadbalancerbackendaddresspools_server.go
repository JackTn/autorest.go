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

// LoadBalancerBackendAddressPoolsServer is a fake server for instances of the armnetwork.LoadBalancerBackendAddressPoolsClient type.
type LoadBalancerBackendAddressPoolsServer struct {
	// BeginCreateOrUpdate is the fake for method LoadBalancerBackendAddressPoolsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, loadBalancerName string, backendAddressPoolName string, parameters armnetwork.BackendAddressPool, options *armnetwork.LoadBalancerBackendAddressPoolsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.LoadBalancerBackendAddressPoolsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method LoadBalancerBackendAddressPoolsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, loadBalancerName string, backendAddressPoolName string, options *armnetwork.LoadBalancerBackendAddressPoolsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.LoadBalancerBackendAddressPoolsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method LoadBalancerBackendAddressPoolsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, loadBalancerName string, backendAddressPoolName string, options *armnetwork.LoadBalancerBackendAddressPoolsClientGetOptions) (resp azfake.Responder[armnetwork.LoadBalancerBackendAddressPoolsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method LoadBalancerBackendAddressPoolsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, loadBalancerName string, options *armnetwork.LoadBalancerBackendAddressPoolsClientListOptions) (resp azfake.PagerResponder[armnetwork.LoadBalancerBackendAddressPoolsClientListResponse])
}

// NewLoadBalancerBackendAddressPoolsServerTransport creates a new instance of LoadBalancerBackendAddressPoolsServerTransport with the provided implementation.
// The returned LoadBalancerBackendAddressPoolsServerTransport instance is connected to an instance of armnetwork.LoadBalancerBackendAddressPoolsClient by way of the
// undefined.Transporter field.
func NewLoadBalancerBackendAddressPoolsServerTransport(srv *LoadBalancerBackendAddressPoolsServer) *LoadBalancerBackendAddressPoolsServerTransport {
	return &LoadBalancerBackendAddressPoolsServerTransport{srv: srv}
}

// LoadBalancerBackendAddressPoolsServerTransport connects instances of armnetwork.LoadBalancerBackendAddressPoolsClient to instances of LoadBalancerBackendAddressPoolsServer.
// Don't use this type directly, use NewLoadBalancerBackendAddressPoolsServerTransport instead.
type LoadBalancerBackendAddressPoolsServerTransport struct {
	srv                 *LoadBalancerBackendAddressPoolsServer
	beginCreateOrUpdate *azfake.PollerResponder[armnetwork.LoadBalancerBackendAddressPoolsClientCreateOrUpdateResponse]
	beginDelete         *azfake.PollerResponder[armnetwork.LoadBalancerBackendAddressPoolsClientDeleteResponse]
	newListPager        *azfake.PagerResponder[armnetwork.LoadBalancerBackendAddressPoolsClientListResponse]
}

// Do implements the policy.Transporter interface for LoadBalancerBackendAddressPoolsServerTransport.
func (l *LoadBalancerBackendAddressPoolsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "LoadBalancerBackendAddressPoolsClient.BeginCreateOrUpdate":
		resp, err = l.dispatchBeginCreateOrUpdate(req)
	case "LoadBalancerBackendAddressPoolsClient.BeginDelete":
		resp, err = l.dispatchBeginDelete(req)
	case "LoadBalancerBackendAddressPoolsClient.Get":
		resp, err = l.dispatchGet(req)
	case "LoadBalancerBackendAddressPoolsClient.NewListPager":
		resp, err = l.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (l *LoadBalancerBackendAddressPoolsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if l.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	if l.beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/loadBalancers/(?P<loadBalancerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backendAddressPools/(?P<backendAddressPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.BackendAddressPool](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		loadBalancerNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("loadBalancerName")])
		if err != nil {
			return nil, err
		}
		backendAddressPoolNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("backendAddressPoolName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := l.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameUnescaped, loadBalancerNameUnescaped, backendAddressPoolNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		l.beginCreateOrUpdate = &respr
	}

	resp, err := server.PollerResponderNext(l.beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(l.beginCreateOrUpdate) {
		l.beginCreateOrUpdate = nil
	}

	return resp, nil
}

func (l *LoadBalancerBackendAddressPoolsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if l.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	if l.beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/loadBalancers/(?P<loadBalancerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backendAddressPools/(?P<backendAddressPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		loadBalancerNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("loadBalancerName")])
		if err != nil {
			return nil, err
		}
		backendAddressPoolNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("backendAddressPoolName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := l.srv.BeginDelete(req.Context(), resourceGroupNameUnescaped, loadBalancerNameUnescaped, backendAddressPoolNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		l.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(l.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(l.beginDelete) {
		l.beginDelete = nil
	}

	return resp, nil
}

func (l *LoadBalancerBackendAddressPoolsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if l.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/loadBalancers/(?P<loadBalancerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backendAddressPools/(?P<backendAddressPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	loadBalancerNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("loadBalancerName")])
	if err != nil {
		return nil, err
	}
	backendAddressPoolNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("backendAddressPoolName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := l.srv.Get(req.Context(), resourceGroupNameUnescaped, loadBalancerNameUnescaped, backendAddressPoolNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).BackendAddressPool, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *LoadBalancerBackendAddressPoolsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if l.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	if l.newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/loadBalancers/(?P<loadBalancerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backendAddressPools`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		loadBalancerNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("loadBalancerName")])
		if err != nil {
			return nil, err
		}
		resp := l.srv.NewListPager(resourceGroupNameUnescaped, loadBalancerNameUnescaped, nil)
		l.newListPager = &resp
		server.PagerResponderInjectNextLinks(l.newListPager, req, func(page *armnetwork.LoadBalancerBackendAddressPoolsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(l.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(l.newListPager) {
		l.newListPager = nil
	}
	return resp, nil
}
