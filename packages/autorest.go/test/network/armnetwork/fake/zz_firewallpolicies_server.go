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

// FirewallPoliciesServer is a fake server for instances of the armnetwork.FirewallPoliciesClient type.
type FirewallPoliciesServer struct {
	// BeginCreateOrUpdate is the fake for method FirewallPoliciesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, firewallPolicyName string, parameters armnetwork.FirewallPolicy, options *armnetwork.FirewallPoliciesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.FirewallPoliciesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method FirewallPoliciesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, firewallPolicyName string, options *armnetwork.FirewallPoliciesClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.FirewallPoliciesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method FirewallPoliciesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, firewallPolicyName string, options *armnetwork.FirewallPoliciesClientGetOptions) (resp azfake.Responder[armnetwork.FirewallPoliciesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method FirewallPoliciesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, options *armnetwork.FirewallPoliciesClientListOptions) (resp azfake.PagerResponder[armnetwork.FirewallPoliciesClientListResponse])

	// NewListAllPager is the fake for method FirewallPoliciesClient.NewListAllPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListAllPager func(options *armnetwork.FirewallPoliciesClientListAllOptions) (resp azfake.PagerResponder[armnetwork.FirewallPoliciesClientListAllResponse])

	// UpdateTags is the fake for method FirewallPoliciesClient.UpdateTags
	// HTTP status codes to indicate success: http.StatusOK
	UpdateTags func(ctx context.Context, resourceGroupName string, firewallPolicyName string, parameters armnetwork.TagsObject, options *armnetwork.FirewallPoliciesClientUpdateTagsOptions) (resp azfake.Responder[armnetwork.FirewallPoliciesClientUpdateTagsResponse], errResp azfake.ErrorResponder)
}

// NewFirewallPoliciesServerTransport creates a new instance of FirewallPoliciesServerTransport with the provided implementation.
// The returned FirewallPoliciesServerTransport instance is connected to an instance of armnetwork.FirewallPoliciesClient by way of the
// undefined.Transporter field.
func NewFirewallPoliciesServerTransport(srv *FirewallPoliciesServer) *FirewallPoliciesServerTransport {
	return &FirewallPoliciesServerTransport{srv: srv}
}

// FirewallPoliciesServerTransport connects instances of armnetwork.FirewallPoliciesClient to instances of FirewallPoliciesServer.
// Don't use this type directly, use NewFirewallPoliciesServerTransport instead.
type FirewallPoliciesServerTransport struct {
	srv                 *FirewallPoliciesServer
	beginCreateOrUpdate *azfake.PollerResponder[armnetwork.FirewallPoliciesClientCreateOrUpdateResponse]
	beginDelete         *azfake.PollerResponder[armnetwork.FirewallPoliciesClientDeleteResponse]
	newListPager        *azfake.PagerResponder[armnetwork.FirewallPoliciesClientListResponse]
	newListAllPager     *azfake.PagerResponder[armnetwork.FirewallPoliciesClientListAllResponse]
}

// Do implements the policy.Transporter interface for FirewallPoliciesServerTransport.
func (f *FirewallPoliciesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "FirewallPoliciesClient.BeginCreateOrUpdate":
		resp, err = f.dispatchBeginCreateOrUpdate(req)
	case "FirewallPoliciesClient.BeginDelete":
		resp, err = f.dispatchBeginDelete(req)
	case "FirewallPoliciesClient.Get":
		resp, err = f.dispatchGet(req)
	case "FirewallPoliciesClient.NewListPager":
		resp, err = f.dispatchNewListPager(req)
	case "FirewallPoliciesClient.NewListAllPager":
		resp, err = f.dispatchNewListAllPager(req)
	case "FirewallPoliciesClient.UpdateTags":
		resp, err = f.dispatchUpdateTags(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (f *FirewallPoliciesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if f.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	if f.beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/firewallPolicies/(?P<firewallPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.FirewallPolicy](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		firewallPolicyNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("firewallPolicyName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := f.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameUnescaped, firewallPolicyNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		f.beginCreateOrUpdate = &respr
	}

	resp, err := server.PollerResponderNext(f.beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(f.beginCreateOrUpdate) {
		f.beginCreateOrUpdate = nil
	}

	return resp, nil
}

func (f *FirewallPoliciesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if f.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	if f.beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/firewallPolicies/(?P<firewallPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		firewallPolicyNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("firewallPolicyName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := f.srv.BeginDelete(req.Context(), resourceGroupNameUnescaped, firewallPolicyNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		f.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(f.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(f.beginDelete) {
		f.beginDelete = nil
	}

	return resp, nil
}

func (f *FirewallPoliciesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if f.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/firewallPolicies/(?P<firewallPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	firewallPolicyNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("firewallPolicyName")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
	var options *armnetwork.FirewallPoliciesClientGetOptions
	if expandParam != nil {
		options = &armnetwork.FirewallPoliciesClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := f.srv.Get(req.Context(), resourceGroupNameUnescaped, firewallPolicyNameUnescaped, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).FirewallPolicy, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (f *FirewallPoliciesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if f.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	if f.newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/firewallPolicies`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := f.srv.NewListPager(resourceGroupNameUnescaped, nil)
		f.newListPager = &resp
		server.PagerResponderInjectNextLinks(f.newListPager, req, func(page *armnetwork.FirewallPoliciesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(f.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(f.newListPager) {
		f.newListPager = nil
	}
	return resp, nil
}

func (f *FirewallPoliciesServerTransport) dispatchNewListAllPager(req *http.Request) (*http.Response, error) {
	if f.srv.NewListAllPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListAllPager not implemented")}
	}
	if f.newListAllPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/firewallPolicies`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := f.srv.NewListAllPager(nil)
		f.newListAllPager = &resp
		server.PagerResponderInjectNextLinks(f.newListAllPager, req, func(page *armnetwork.FirewallPoliciesClientListAllResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(f.newListAllPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(f.newListAllPager) {
		f.newListAllPager = nil
	}
	return resp, nil
}

func (f *FirewallPoliciesServerTransport) dispatchUpdateTags(req *http.Request) (*http.Response, error) {
	if f.srv.UpdateTags == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateTags not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/firewallPolicies/(?P<firewallPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	firewallPolicyNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("firewallPolicyName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.UpdateTags(req.Context(), resourceGroupNameUnescaped, firewallPolicyNameUnescaped, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).FirewallPolicy, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
