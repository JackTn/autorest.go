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

// LocalNetworkGatewaysServer is a fake server for instances of the armnetwork.LocalNetworkGatewaysClient type.
type LocalNetworkGatewaysServer struct {
	// BeginCreateOrUpdate is the fake for method LocalNetworkGatewaysClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, parameters armnetwork.LocalNetworkGateway, options *armnetwork.LocalNetworkGatewaysClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.LocalNetworkGatewaysClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method LocalNetworkGatewaysClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, options *armnetwork.LocalNetworkGatewaysClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.LocalNetworkGatewaysClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method LocalNetworkGatewaysClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, options *armnetwork.LocalNetworkGatewaysClientGetOptions) (resp azfake.Responder[armnetwork.LocalNetworkGatewaysClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method LocalNetworkGatewaysClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, options *armnetwork.LocalNetworkGatewaysClientListOptions) (resp azfake.PagerResponder[armnetwork.LocalNetworkGatewaysClientListResponse])

	// UpdateTags is the fake for method LocalNetworkGatewaysClient.UpdateTags
	// HTTP status codes to indicate success: http.StatusOK
	UpdateTags func(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, parameters armnetwork.TagsObject, options *armnetwork.LocalNetworkGatewaysClientUpdateTagsOptions) (resp azfake.Responder[armnetwork.LocalNetworkGatewaysClientUpdateTagsResponse], errResp azfake.ErrorResponder)
}

// NewLocalNetworkGatewaysServerTransport creates a new instance of LocalNetworkGatewaysServerTransport with the provided implementation.
// The returned LocalNetworkGatewaysServerTransport instance is connected to an instance of armnetwork.LocalNetworkGatewaysClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewLocalNetworkGatewaysServerTransport(srv *LocalNetworkGatewaysServer) *LocalNetworkGatewaysServerTransport {
	return &LocalNetworkGatewaysServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armnetwork.LocalNetworkGatewaysClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armnetwork.LocalNetworkGatewaysClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armnetwork.LocalNetworkGatewaysClientListResponse]](),
	}
}

// LocalNetworkGatewaysServerTransport connects instances of armnetwork.LocalNetworkGatewaysClient to instances of LocalNetworkGatewaysServer.
// Don't use this type directly, use NewLocalNetworkGatewaysServerTransport instead.
type LocalNetworkGatewaysServerTransport struct {
	srv                 *LocalNetworkGatewaysServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armnetwork.LocalNetworkGatewaysClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armnetwork.LocalNetworkGatewaysClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armnetwork.LocalNetworkGatewaysClientListResponse]]
}

// Do implements the policy.Transporter interface for LocalNetworkGatewaysServerTransport.
func (l *LocalNetworkGatewaysServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "LocalNetworkGatewaysClient.BeginCreateOrUpdate":
		resp, err = l.dispatchBeginCreateOrUpdate(req)
	case "LocalNetworkGatewaysClient.BeginDelete":
		resp, err = l.dispatchBeginDelete(req)
	case "LocalNetworkGatewaysClient.Get":
		resp, err = l.dispatchGet(req)
	case "LocalNetworkGatewaysClient.NewListPager":
		resp, err = l.dispatchNewListPager(req)
	case "LocalNetworkGatewaysClient.UpdateTags":
		resp, err = l.dispatchUpdateTags(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (l *LocalNetworkGatewaysServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if l.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := l.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/localNetworkGateways/(?P<localNetworkGatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.LocalNetworkGateway](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		localNetworkGatewayNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("localNetworkGatewayName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := l.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameUnescaped, localNetworkGatewayNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		l.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		l.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		l.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (l *LocalNetworkGatewaysServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if l.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := l.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/localNetworkGateways/(?P<localNetworkGatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		localNetworkGatewayNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("localNetworkGatewayName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := l.srv.BeginDelete(req.Context(), resourceGroupNameUnescaped, localNetworkGatewayNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		l.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		l.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		l.beginDelete.remove(req)
	}

	return resp, nil
}

func (l *LocalNetworkGatewaysServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if l.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/localNetworkGateways/(?P<localNetworkGatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	localNetworkGatewayNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("localNetworkGatewayName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := l.srv.Get(req.Context(), resourceGroupNameUnescaped, localNetworkGatewayNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).LocalNetworkGateway, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *LocalNetworkGatewaysServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if l.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := l.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/localNetworkGateways`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := l.srv.NewListPager(resourceGroupNameUnescaped, nil)
		newListPager = &resp
		l.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.LocalNetworkGatewaysClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		l.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		l.newListPager.remove(req)
	}
	return resp, nil
}

func (l *LocalNetworkGatewaysServerTransport) dispatchUpdateTags(req *http.Request) (*http.Response, error) {
	if l.srv.UpdateTags == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateTags not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/localNetworkGateways/(?P<localNetworkGatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	localNetworkGatewayNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("localNetworkGatewayName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := l.srv.UpdateTags(req.Context(), resourceGroupNameUnescaped, localNetworkGatewayNameUnescaped, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).LocalNetworkGateway, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
