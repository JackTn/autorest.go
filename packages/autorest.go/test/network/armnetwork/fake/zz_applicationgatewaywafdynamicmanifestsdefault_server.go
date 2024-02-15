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

// ApplicationGatewayWafDynamicManifestsDefaultServer is a fake server for instances of the armnetwork.ApplicationGatewayWafDynamicManifestsDefaultClient type.
type ApplicationGatewayWafDynamicManifestsDefaultServer struct {
	// Get is the fake for method ApplicationGatewayWafDynamicManifestsDefaultClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, location string, options *armnetwork.ApplicationGatewayWafDynamicManifestsDefaultClientGetOptions) (resp azfake.Responder[armnetwork.ApplicationGatewayWafDynamicManifestsDefaultClientGetResponse], errResp azfake.ErrorResponder)
}

// NewApplicationGatewayWafDynamicManifestsDefaultServerTransport creates a new instance of ApplicationGatewayWafDynamicManifestsDefaultServerTransport with the provided implementation.
// The returned ApplicationGatewayWafDynamicManifestsDefaultServerTransport instance is connected to an instance of armnetwork.ApplicationGatewayWafDynamicManifestsDefaultClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewApplicationGatewayWafDynamicManifestsDefaultServerTransport(srv *ApplicationGatewayWafDynamicManifestsDefaultServer) *ApplicationGatewayWafDynamicManifestsDefaultServerTransport {
	return &ApplicationGatewayWafDynamicManifestsDefaultServerTransport{srv: srv}
}

// ApplicationGatewayWafDynamicManifestsDefaultServerTransport connects instances of armnetwork.ApplicationGatewayWafDynamicManifestsDefaultClient to instances of ApplicationGatewayWafDynamicManifestsDefaultServer.
// Don't use this type directly, use NewApplicationGatewayWafDynamicManifestsDefaultServerTransport instead.
type ApplicationGatewayWafDynamicManifestsDefaultServerTransport struct {
	srv *ApplicationGatewayWafDynamicManifestsDefaultServer
}

// Do implements the policy.Transporter interface for ApplicationGatewayWafDynamicManifestsDefaultServerTransport.
func (a *ApplicationGatewayWafDynamicManifestsDefaultServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ApplicationGatewayWafDynamicManifestsDefaultClient.Get":
		resp, err = a.dispatchGet(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *ApplicationGatewayWafDynamicManifestsDefaultServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applicationGatewayWafDynamicManifests/dafault`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), locationParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ApplicationGatewayWafDynamicManifestResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
