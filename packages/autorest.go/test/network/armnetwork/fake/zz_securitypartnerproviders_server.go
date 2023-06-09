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

// SecurityPartnerProvidersServer is a fake server for instances of the armnetwork.SecurityPartnerProvidersClient type.
type SecurityPartnerProvidersServer struct {
	// BeginCreateOrUpdate is the fake for method SecurityPartnerProvidersClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, securityPartnerProviderName string, parameters armnetwork.SecurityPartnerProvider, options *armnetwork.SecurityPartnerProvidersClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.SecurityPartnerProvidersClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method SecurityPartnerProvidersClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, securityPartnerProviderName string, options *armnetwork.SecurityPartnerProvidersClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.SecurityPartnerProvidersClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method SecurityPartnerProvidersClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, securityPartnerProviderName string, options *armnetwork.SecurityPartnerProvidersClientGetOptions) (resp azfake.Responder[armnetwork.SecurityPartnerProvidersClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method SecurityPartnerProvidersClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armnetwork.SecurityPartnerProvidersClientListOptions) (resp azfake.PagerResponder[armnetwork.SecurityPartnerProvidersClientListResponse])

	// NewListByResourceGroupPager is the fake for method SecurityPartnerProvidersClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armnetwork.SecurityPartnerProvidersClientListByResourceGroupOptions) (resp azfake.PagerResponder[armnetwork.SecurityPartnerProvidersClientListByResourceGroupResponse])

	// UpdateTags is the fake for method SecurityPartnerProvidersClient.UpdateTags
	// HTTP status codes to indicate success: http.StatusOK
	UpdateTags func(ctx context.Context, resourceGroupName string, securityPartnerProviderName string, parameters armnetwork.TagsObject, options *armnetwork.SecurityPartnerProvidersClientUpdateTagsOptions) (resp azfake.Responder[armnetwork.SecurityPartnerProvidersClientUpdateTagsResponse], errResp azfake.ErrorResponder)
}

// NewSecurityPartnerProvidersServerTransport creates a new instance of SecurityPartnerProvidersServerTransport with the provided implementation.
// The returned SecurityPartnerProvidersServerTransport instance is connected to an instance of armnetwork.SecurityPartnerProvidersClient by way of the
// undefined.Transporter field.
func NewSecurityPartnerProvidersServerTransport(srv *SecurityPartnerProvidersServer) *SecurityPartnerProvidersServerTransport {
	return &SecurityPartnerProvidersServerTransport{srv: srv}
}

// SecurityPartnerProvidersServerTransport connects instances of armnetwork.SecurityPartnerProvidersClient to instances of SecurityPartnerProvidersServer.
// Don't use this type directly, use NewSecurityPartnerProvidersServerTransport instead.
type SecurityPartnerProvidersServerTransport struct {
	srv                         *SecurityPartnerProvidersServer
	beginCreateOrUpdate         *azfake.PollerResponder[armnetwork.SecurityPartnerProvidersClientCreateOrUpdateResponse]
	beginDelete                 *azfake.PollerResponder[armnetwork.SecurityPartnerProvidersClientDeleteResponse]
	newListPager                *azfake.PagerResponder[armnetwork.SecurityPartnerProvidersClientListResponse]
	newListByResourceGroupPager *azfake.PagerResponder[armnetwork.SecurityPartnerProvidersClientListByResourceGroupResponse]
}

// Do implements the policy.Transporter interface for SecurityPartnerProvidersServerTransport.
func (s *SecurityPartnerProvidersServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "SecurityPartnerProvidersClient.BeginCreateOrUpdate":
		resp, err = s.dispatchBeginCreateOrUpdate(req)
	case "SecurityPartnerProvidersClient.BeginDelete":
		resp, err = s.dispatchBeginDelete(req)
	case "SecurityPartnerProvidersClient.Get":
		resp, err = s.dispatchGet(req)
	case "SecurityPartnerProvidersClient.NewListPager":
		resp, err = s.dispatchNewListPager(req)
	case "SecurityPartnerProvidersClient.NewListByResourceGroupPager":
		resp, err = s.dispatchNewListByResourceGroupPager(req)
	case "SecurityPartnerProvidersClient.UpdateTags":
		resp, err = s.dispatchUpdateTags(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SecurityPartnerProvidersServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	if s.beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/securityPartnerProviders/(?P<securityPartnerProviderName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.SecurityPartnerProvider](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		securityPartnerProviderNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("securityPartnerProviderName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameUnescaped, securityPartnerProviderNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		s.beginCreateOrUpdate = &respr
	}

	resp, err := server.PollerResponderNext(s.beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(s.beginCreateOrUpdate) {
		s.beginCreateOrUpdate = nil
	}

	return resp, nil
}

func (s *SecurityPartnerProvidersServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if s.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	if s.beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/securityPartnerProviders/(?P<securityPartnerProviderName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		securityPartnerProviderNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("securityPartnerProviderName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginDelete(req.Context(), resourceGroupNameUnescaped, securityPartnerProviderNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		s.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(s.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(s.beginDelete) {
		s.beginDelete = nil
	}

	return resp, nil
}

func (s *SecurityPartnerProvidersServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/securityPartnerProviders/(?P<securityPartnerProviderName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	securityPartnerProviderNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("securityPartnerProviderName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameUnescaped, securityPartnerProviderNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SecurityPartnerProvider, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SecurityPartnerProvidersServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	if s.newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/securityPartnerProviders`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := s.srv.NewListPager(nil)
		s.newListPager = &resp
		server.PagerResponderInjectNextLinks(s.newListPager, req, func(page *armnetwork.SecurityPartnerProvidersClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(s.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(s.newListPager) {
		s.newListPager = nil
	}
	return resp, nil
}

func (s *SecurityPartnerProvidersServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	if s.newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/securityPartnerProviders`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListByResourceGroupPager(resourceGroupNameUnescaped, nil)
		s.newListByResourceGroupPager = &resp
		server.PagerResponderInjectNextLinks(s.newListByResourceGroupPager, req, func(page *armnetwork.SecurityPartnerProvidersClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(s.newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(s.newListByResourceGroupPager) {
		s.newListByResourceGroupPager = nil
	}
	return resp, nil
}

func (s *SecurityPartnerProvidersServerTransport) dispatchUpdateTags(req *http.Request) (*http.Response, error) {
	if s.srv.UpdateTags == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateTags not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/securityPartnerProviders/(?P<securityPartnerProviderName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	securityPartnerProviderNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("securityPartnerProviderName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.UpdateTags(req.Context(), resourceGroupNameUnescaped, securityPartnerProviderNameUnescaped, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SecurityPartnerProvider, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
