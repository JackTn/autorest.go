//go:build go1.18
// +build go1.18

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

// CloudServiceRolesServer is a fake server for instances of the armcompute.CloudServiceRolesClient type.
type CloudServiceRolesServer struct {
	// Get is the fake for method CloudServiceRolesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, roleName string, resourceGroupName string, cloudServiceName string, options *armcompute.CloudServiceRolesClientGetOptions) (resp azfake.Responder[armcompute.CloudServiceRolesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method CloudServiceRolesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, cloudServiceName string, options *armcompute.CloudServiceRolesClientListOptions) (resp azfake.PagerResponder[armcompute.CloudServiceRolesClientListResponse])
}

// NewCloudServiceRolesServerTransport creates a new instance of CloudServiceRolesServerTransport with the provided implementation.
// The returned CloudServiceRolesServerTransport instance is connected to an instance of armcompute.CloudServiceRolesClient by way of the
// undefined.Transporter field.
func NewCloudServiceRolesServerTransport(srv *CloudServiceRolesServer) *CloudServiceRolesServerTransport {
	return &CloudServiceRolesServerTransport{srv: srv}
}

// CloudServiceRolesServerTransport connects instances of armcompute.CloudServiceRolesClient to instances of CloudServiceRolesServer.
// Don't use this type directly, use NewCloudServiceRolesServerTransport instead.
type CloudServiceRolesServerTransport struct {
	srv          *CloudServiceRolesServer
	newListPager *azfake.PagerResponder[armcompute.CloudServiceRolesClientListResponse]
}

// Do implements the policy.Transporter interface for CloudServiceRolesServerTransport.
func (c *CloudServiceRolesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "CloudServiceRolesClient.Get":
		resp, err = c.dispatchGet(req)
	case "CloudServiceRolesClient.NewListPager":
		resp, err = c.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CloudServiceRolesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/roles/(?P<roleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	roleNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("roleName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	cloudServiceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("cloudServiceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), roleNameUnescaped, resourceGroupNameUnescaped, cloudServiceNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CloudServiceRole, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CloudServiceRolesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	if c.newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/roles`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		cloudServiceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("cloudServiceName")])
		if err != nil {
			return nil, err
		}
		resp := c.srv.NewListPager(resourceGroupNameUnescaped, cloudServiceNameUnescaped, nil)
		c.newListPager = &resp
		server.PagerResponderInjectNextLinks(c.newListPager, req, func(page *armcompute.CloudServiceRolesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(c.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(c.newListPager) {
		c.newListPager = nil
	}
	return resp, nil
}
