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

// WebApplicationFirewallPoliciesServer is a fake server for instances of the armnetwork.WebApplicationFirewallPoliciesClient type.
type WebApplicationFirewallPoliciesServer struct {
	// CreateOrUpdate is the fake for method WebApplicationFirewallPoliciesClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, policyName string, parameters armnetwork.WebApplicationFirewallPolicy, options *armnetwork.WebApplicationFirewallPoliciesClientCreateOrUpdateOptions) (resp azfake.Responder[armnetwork.WebApplicationFirewallPoliciesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method WebApplicationFirewallPoliciesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, policyName string, options *armnetwork.WebApplicationFirewallPoliciesClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.WebApplicationFirewallPoliciesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method WebApplicationFirewallPoliciesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, policyName string, options *armnetwork.WebApplicationFirewallPoliciesClientGetOptions) (resp azfake.Responder[armnetwork.WebApplicationFirewallPoliciesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method WebApplicationFirewallPoliciesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, options *armnetwork.WebApplicationFirewallPoliciesClientListOptions) (resp azfake.PagerResponder[armnetwork.WebApplicationFirewallPoliciesClientListResponse])

	// NewListAllPager is the fake for method WebApplicationFirewallPoliciesClient.NewListAllPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListAllPager func(options *armnetwork.WebApplicationFirewallPoliciesClientListAllOptions) (resp azfake.PagerResponder[armnetwork.WebApplicationFirewallPoliciesClientListAllResponse])
}

// NewWebApplicationFirewallPoliciesServerTransport creates a new instance of WebApplicationFirewallPoliciesServerTransport with the provided implementation.
// The returned WebApplicationFirewallPoliciesServerTransport instance is connected to an instance of armnetwork.WebApplicationFirewallPoliciesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewWebApplicationFirewallPoliciesServerTransport(srv *WebApplicationFirewallPoliciesServer) *WebApplicationFirewallPoliciesServerTransport {
	return &WebApplicationFirewallPoliciesServerTransport{
		srv:             srv,
		beginDelete:     newTracker[azfake.PollerResponder[armnetwork.WebApplicationFirewallPoliciesClientDeleteResponse]](),
		newListPager:    newTracker[azfake.PagerResponder[armnetwork.WebApplicationFirewallPoliciesClientListResponse]](),
		newListAllPager: newTracker[azfake.PagerResponder[armnetwork.WebApplicationFirewallPoliciesClientListAllResponse]](),
	}
}

// WebApplicationFirewallPoliciesServerTransport connects instances of armnetwork.WebApplicationFirewallPoliciesClient to instances of WebApplicationFirewallPoliciesServer.
// Don't use this type directly, use NewWebApplicationFirewallPoliciesServerTransport instead.
type WebApplicationFirewallPoliciesServerTransport struct {
	srv             *WebApplicationFirewallPoliciesServer
	beginDelete     *tracker[azfake.PollerResponder[armnetwork.WebApplicationFirewallPoliciesClientDeleteResponse]]
	newListPager    *tracker[azfake.PagerResponder[armnetwork.WebApplicationFirewallPoliciesClientListResponse]]
	newListAllPager *tracker[azfake.PagerResponder[armnetwork.WebApplicationFirewallPoliciesClientListAllResponse]]
}

// Do implements the policy.Transporter interface for WebApplicationFirewallPoliciesServerTransport.
func (w *WebApplicationFirewallPoliciesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "WebApplicationFirewallPoliciesClient.CreateOrUpdate":
		resp, err = w.dispatchCreateOrUpdate(req)
	case "WebApplicationFirewallPoliciesClient.BeginDelete":
		resp, err = w.dispatchBeginDelete(req)
	case "WebApplicationFirewallPoliciesClient.Get":
		resp, err = w.dispatchGet(req)
	case "WebApplicationFirewallPoliciesClient.NewListPager":
		resp, err = w.dispatchNewListPager(req)
	case "WebApplicationFirewallPoliciesClient.NewListAllPager":
		resp, err = w.dispatchNewListAllPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w *WebApplicationFirewallPoliciesServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if w.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/ApplicationGatewayWebApplicationFirewallPolicies/(?P<policyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetwork.WebApplicationFirewallPolicy](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	policyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, policyNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).WebApplicationFirewallPolicy, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WebApplicationFirewallPoliciesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if w.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := w.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/ApplicationGatewayWebApplicationFirewallPolicies/(?P<policyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		policyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := w.srv.BeginDelete(req.Context(), resourceGroupNameParam, policyNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		w.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		w.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		w.beginDelete.remove(req)
	}

	return resp, nil
}

func (w *WebApplicationFirewallPoliciesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if w.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/ApplicationGatewayWebApplicationFirewallPolicies/(?P<policyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	policyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Get(req.Context(), resourceGroupNameParam, policyNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).WebApplicationFirewallPolicy, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WebApplicationFirewallPoliciesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if w.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := w.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/ApplicationGatewayWebApplicationFirewallPolicies`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := w.srv.NewListPager(resourceGroupNameParam, nil)
		newListPager = &resp
		w.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.WebApplicationFirewallPoliciesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		w.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		w.newListPager.remove(req)
	}
	return resp, nil
}

func (w *WebApplicationFirewallPoliciesServerTransport) dispatchNewListAllPager(req *http.Request) (*http.Response, error) {
	if w.srv.NewListAllPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListAllPager not implemented")}
	}
	newListAllPager := w.newListAllPager.get(req)
	if newListAllPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/ApplicationGatewayWebApplicationFirewallPolicies`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := w.srv.NewListAllPager(nil)
		newListAllPager = &resp
		w.newListAllPager.add(req, newListAllPager)
		server.PagerResponderInjectNextLinks(newListAllPager, req, func(page *armnetwork.WebApplicationFirewallPoliciesClientListAllResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListAllPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		w.newListAllPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListAllPager) {
		w.newListAllPager.remove(req)
	}
	return resp, nil
}
