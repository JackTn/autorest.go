// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"armdatabasewatcher"
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

// TargetsServer is a fake server for instances of the armdatabasewatcher.TargetsClient type.
type TargetsServer struct {
	// CreateOrUpdate is the fake for method TargetsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, watcherName string, targetName string, resource armdatabasewatcher.Target, options *armdatabasewatcher.TargetsClientCreateOrUpdateOptions) (resp azfake.Responder[armdatabasewatcher.TargetsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method TargetsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, watcherName string, targetName string, options *armdatabasewatcher.TargetsClientDeleteOptions) (resp azfake.Responder[armdatabasewatcher.TargetsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method TargetsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, watcherName string, targetName string, options *armdatabasewatcher.TargetsClientGetOptions) (resp azfake.Responder[armdatabasewatcher.TargetsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByWatcherPager is the fake for method TargetsClient.NewListByWatcherPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByWatcherPager func(resourceGroupName string, watcherName string, options *armdatabasewatcher.TargetsClientListByWatcherOptions) (resp azfake.PagerResponder[armdatabasewatcher.TargetsClientListByWatcherResponse])
}

// NewTargetsServerTransport creates a new instance of TargetsServerTransport with the provided implementation.
// The returned TargetsServerTransport instance is connected to an instance of armdatabasewatcher.TargetsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewTargetsServerTransport(srv *TargetsServer) *TargetsServerTransport {
	return &TargetsServerTransport{
		srv:                   srv,
		newListByWatcherPager: newTracker[azfake.PagerResponder[armdatabasewatcher.TargetsClientListByWatcherResponse]](),
	}
}

// TargetsServerTransport connects instances of armdatabasewatcher.TargetsClient to instances of TargetsServer.
// Don't use this type directly, use NewTargetsServerTransport instead.
type TargetsServerTransport struct {
	srv                   *TargetsServer
	newListByWatcherPager *tracker[azfake.PagerResponder[armdatabasewatcher.TargetsClientListByWatcherResponse]]
}

// Do implements the policy.Transporter interface for TargetsServerTransport.
func (t *TargetsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return t.dispatchToMethodFake(req, method)
}

func (t *TargetsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "TargetsClient.CreateOrUpdate":
		resp, err = t.dispatchCreateOrUpdate(req)
	case "TargetsClient.Delete":
		resp, err = t.dispatchDelete(req)
	case "TargetsClient.Get":
		resp, err = t.dispatchGet(req)
	case "TargetsClient.NewListByWatcherPager":
		resp, err = t.dispatchNewListByWatcherPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (t *TargetsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if t.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DatabaseWatcher/watchers/(?P<watcherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/targets/(?P<targetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdatabasewatcher.Target](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	watcherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("watcherName")])
	if err != nil {
		return nil, err
	}
	targetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("targetName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, watcherNameParam, targetNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Target, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TargetsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if t.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DatabaseWatcher/watchers/(?P<watcherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/targets/(?P<targetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	watcherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("watcherName")])
	if err != nil {
		return nil, err
	}
	targetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("targetName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.Delete(req.Context(), resourceGroupNameParam, watcherNameParam, targetNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TargetsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if t.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DatabaseWatcher/watchers/(?P<watcherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/targets/(?P<targetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	watcherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("watcherName")])
	if err != nil {
		return nil, err
	}
	targetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("targetName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.Get(req.Context(), resourceGroupNameParam, watcherNameParam, targetNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Target, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TargetsServerTransport) dispatchNewListByWatcherPager(req *http.Request) (*http.Response, error) {
	if t.srv.NewListByWatcherPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByWatcherPager not implemented")}
	}
	newListByWatcherPager := t.newListByWatcherPager.get(req)
	if newListByWatcherPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DatabaseWatcher/watchers/(?P<watcherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/targets`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		watcherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("watcherName")])
		if err != nil {
			return nil, err
		}
		resp := t.srv.NewListByWatcherPager(resourceGroupNameParam, watcherNameParam, nil)
		newListByWatcherPager = &resp
		t.newListByWatcherPager.add(req, newListByWatcherPager)
		server.PagerResponderInjectNextLinks(newListByWatcherPager, req, func(page *armdatabasewatcher.TargetsClientListByWatcherResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByWatcherPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		t.newListByWatcherPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByWatcherPager) {
		t.newListByWatcherPager.remove(req)
	}
	return resp, nil
}
