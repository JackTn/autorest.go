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

// ExpressRoutePortsLocationsServer is a fake server for instances of the armnetwork.ExpressRoutePortsLocationsClient type.
type ExpressRoutePortsLocationsServer struct {
	// Get is the fake for method ExpressRoutePortsLocationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, locationName string, options *armnetwork.ExpressRoutePortsLocationsClientGetOptions) (resp azfake.Responder[armnetwork.ExpressRoutePortsLocationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ExpressRoutePortsLocationsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armnetwork.ExpressRoutePortsLocationsClientListOptions) (resp azfake.PagerResponder[armnetwork.ExpressRoutePortsLocationsClientListResponse])
}

// NewExpressRoutePortsLocationsServerTransport creates a new instance of ExpressRoutePortsLocationsServerTransport with the provided implementation.
// The returned ExpressRoutePortsLocationsServerTransport instance is connected to an instance of armnetwork.ExpressRoutePortsLocationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewExpressRoutePortsLocationsServerTransport(srv *ExpressRoutePortsLocationsServer) *ExpressRoutePortsLocationsServerTransport {
	return &ExpressRoutePortsLocationsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armnetwork.ExpressRoutePortsLocationsClientListResponse]](),
	}
}

// ExpressRoutePortsLocationsServerTransport connects instances of armnetwork.ExpressRoutePortsLocationsClient to instances of ExpressRoutePortsLocationsServer.
// Don't use this type directly, use NewExpressRoutePortsLocationsServerTransport instead.
type ExpressRoutePortsLocationsServerTransport struct {
	srv          *ExpressRoutePortsLocationsServer
	newListPager *tracker[azfake.PagerResponder[armnetwork.ExpressRoutePortsLocationsClientListResponse]]
}

// Do implements the policy.Transporter interface for ExpressRoutePortsLocationsServerTransport.
func (e *ExpressRoutePortsLocationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ExpressRoutePortsLocationsClient.Get":
		resp, err = e.dispatchGet(req)
	case "ExpressRoutePortsLocationsClient.NewListPager":
		resp, err = e.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (e *ExpressRoutePortsLocationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if e.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/ExpressRoutePortsLocations/(?P<locationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("locationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Get(req.Context(), locationNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ExpressRoutePortsLocation, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExpressRoutePortsLocationsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if e.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := e.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/ExpressRoutePortsLocations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := e.srv.NewListPager(nil)
		newListPager = &resp
		e.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.ExpressRoutePortsLocationsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		e.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		e.newListPager.remove(req)
	}
	return resp, nil
}
