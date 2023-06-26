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

// LoadBalancerFrontendIPConfigurationsServer is a fake server for instances of the armnetwork.LoadBalancerFrontendIPConfigurationsClient type.
type LoadBalancerFrontendIPConfigurationsServer struct {
	// Get is the fake for method LoadBalancerFrontendIPConfigurationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, loadBalancerName string, frontendIPConfigurationName string, options *armnetwork.LoadBalancerFrontendIPConfigurationsClientGetOptions) (resp azfake.Responder[armnetwork.LoadBalancerFrontendIPConfigurationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method LoadBalancerFrontendIPConfigurationsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, loadBalancerName string, options *armnetwork.LoadBalancerFrontendIPConfigurationsClientListOptions) (resp azfake.PagerResponder[armnetwork.LoadBalancerFrontendIPConfigurationsClientListResponse])
}

// NewLoadBalancerFrontendIPConfigurationsServerTransport creates a new instance of LoadBalancerFrontendIPConfigurationsServerTransport with the provided implementation.
// The returned LoadBalancerFrontendIPConfigurationsServerTransport instance is connected to an instance of armnetwork.LoadBalancerFrontendIPConfigurationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewLoadBalancerFrontendIPConfigurationsServerTransport(srv *LoadBalancerFrontendIPConfigurationsServer) *LoadBalancerFrontendIPConfigurationsServerTransport {
	return &LoadBalancerFrontendIPConfigurationsServerTransport{srv: srv}
}

// LoadBalancerFrontendIPConfigurationsServerTransport connects instances of armnetwork.LoadBalancerFrontendIPConfigurationsClient to instances of LoadBalancerFrontendIPConfigurationsServer.
// Don't use this type directly, use NewLoadBalancerFrontendIPConfigurationsServerTransport instead.
type LoadBalancerFrontendIPConfigurationsServerTransport struct {
	srv          *LoadBalancerFrontendIPConfigurationsServer
	newListPager *azfake.PagerResponder[armnetwork.LoadBalancerFrontendIPConfigurationsClientListResponse]
}

// Do implements the policy.Transporter interface for LoadBalancerFrontendIPConfigurationsServerTransport.
func (l *LoadBalancerFrontendIPConfigurationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "LoadBalancerFrontendIPConfigurationsClient.Get":
		resp, err = l.dispatchGet(req)
	case "LoadBalancerFrontendIPConfigurationsClient.NewListPager":
		resp, err = l.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (l *LoadBalancerFrontendIPConfigurationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if l.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/loadBalancers/(?P<loadBalancerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/frontendIPConfigurations/(?P<frontendIPConfigurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	frontendIPConfigurationNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("frontendIPConfigurationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := l.srv.Get(req.Context(), resourceGroupNameUnescaped, loadBalancerNameUnescaped, frontendIPConfigurationNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).FrontendIPConfiguration, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *LoadBalancerFrontendIPConfigurationsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if l.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	if l.newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/loadBalancers/(?P<loadBalancerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/frontendIPConfigurations`
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
		server.PagerResponderInjectNextLinks(l.newListPager, req, func(page *armnetwork.LoadBalancerFrontendIPConfigurationsClientListResponse, createLink func() string) {
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