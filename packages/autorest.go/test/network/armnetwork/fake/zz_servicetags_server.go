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
	"net/http"
	"net/url"
	"regexp"
)

// ServiceTagsServer is a fake server for instances of the armnetwork.ServiceTagsClient type.
type ServiceTagsServer struct {
	// List is the fake for method ServiceTagsClient.List
	// HTTP status codes to indicate success: http.StatusOK
	List func(ctx context.Context, location string, options *armnetwork.ServiceTagsClientListOptions) (resp azfake.Responder[armnetwork.ServiceTagsClientListResponse], errResp azfake.ErrorResponder)
}

// NewServiceTagsServerTransport creates a new instance of ServiceTagsServerTransport with the provided implementation.
// The returned ServiceTagsServerTransport instance is connected to an instance of armnetwork.ServiceTagsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServiceTagsServerTransport(srv *ServiceTagsServer) *ServiceTagsServerTransport {
	return &ServiceTagsServerTransport{srv: srv}
}

// ServiceTagsServerTransport connects instances of armnetwork.ServiceTagsClient to instances of ServiceTagsServer.
// Don't use this type directly, use NewServiceTagsServerTransport instead.
type ServiceTagsServerTransport struct {
	srv *ServiceTagsServer
}

// Do implements the policy.Transporter interface for ServiceTagsServerTransport.
func (s *ServiceTagsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ServiceTagsClient.List":
		resp, err = s.dispatchList(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ServiceTagsServerTransport) dispatchList(req *http.Request) (*http.Response, error) {
	if s.srv.List == nil {
		return nil, &nonRetriableError{errors.New("fake for method List not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/serviceTags`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.List(req.Context(), locationUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ServiceTagsListResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
