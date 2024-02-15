// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"armnetwork"
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

// ApplicationGatewayWafDynamicManifestsServer is a fake server for instances of the armnetwork.ApplicationGatewayWafDynamicManifestsClient type.
type ApplicationGatewayWafDynamicManifestsServer struct {
	// NewGetPager is the fake for method ApplicationGatewayWafDynamicManifestsClient.NewGetPager
	// HTTP status codes to indicate success: http.StatusOK
	NewGetPager func(location string, options *armnetwork.ApplicationGatewayWafDynamicManifestsClientGetOptions) (resp azfake.PagerResponder[armnetwork.ApplicationGatewayWafDynamicManifestsClientGetResponse])
}

// NewApplicationGatewayWafDynamicManifestsServerTransport creates a new instance of ApplicationGatewayWafDynamicManifestsServerTransport with the provided implementation.
// The returned ApplicationGatewayWafDynamicManifestsServerTransport instance is connected to an instance of armnetwork.ApplicationGatewayWafDynamicManifestsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewApplicationGatewayWafDynamicManifestsServerTransport(srv *ApplicationGatewayWafDynamicManifestsServer) *ApplicationGatewayWafDynamicManifestsServerTransport {
	return &ApplicationGatewayWafDynamicManifestsServerTransport{
		srv:         srv,
		newGetPager: newTracker[azfake.PagerResponder[armnetwork.ApplicationGatewayWafDynamicManifestsClientGetResponse]](),
	}
}

// ApplicationGatewayWafDynamicManifestsServerTransport connects instances of armnetwork.ApplicationGatewayWafDynamicManifestsClient to instances of ApplicationGatewayWafDynamicManifestsServer.
// Don't use this type directly, use NewApplicationGatewayWafDynamicManifestsServerTransport instead.
type ApplicationGatewayWafDynamicManifestsServerTransport struct {
	srv         *ApplicationGatewayWafDynamicManifestsServer
	newGetPager *tracker[azfake.PagerResponder[armnetwork.ApplicationGatewayWafDynamicManifestsClientGetResponse]]
}

// Do implements the policy.Transporter interface for ApplicationGatewayWafDynamicManifestsServerTransport.
func (a *ApplicationGatewayWafDynamicManifestsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ApplicationGatewayWafDynamicManifestsClient.NewGetPager":
		resp, err = a.dispatchNewGetPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *ApplicationGatewayWafDynamicManifestsServerTransport) dispatchNewGetPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewGetPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewGetPager not implemented")}
	}
	newGetPager := a.newGetPager.get(req)
	if newGetPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applicationGatewayWafDynamicManifests`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewGetPager(locationParam, nil)
		newGetPager = &resp
		a.newGetPager.add(req, newGetPager)
		server.PagerResponderInjectNextLinks(newGetPager, req, func(page *armnetwork.ApplicationGatewayWafDynamicManifestsClientGetResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newGetPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newGetPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newGetPager) {
		a.newGetPager.remove(req)
	}
	return resp, nil
}
