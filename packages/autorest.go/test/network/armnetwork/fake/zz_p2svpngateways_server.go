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

// P2SVPNGatewaysServer is a fake server for instances of the armnetwork.P2SVPNGatewaysClient type.
type P2SVPNGatewaysServer struct {
	// BeginCreateOrUpdate is the fake for method P2SVPNGatewaysClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, gatewayName string, p2SVPNGatewayParameters armnetwork.P2SVPNGateway, options *armnetwork.P2SVPNGatewaysClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.P2SVPNGatewaysClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method P2SVPNGatewaysClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, gatewayName string, options *armnetwork.P2SVPNGatewaysClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.P2SVPNGatewaysClientDeleteResponse], errResp azfake.ErrorResponder)

	// BeginDisconnectP2SVPNConnections is the fake for method P2SVPNGatewaysClient.BeginDisconnectP2SVPNConnections
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginDisconnectP2SVPNConnections func(ctx context.Context, resourceGroupName string, p2SVPNGatewayName string, request armnetwork.P2SVPNConnectionRequest, options *armnetwork.P2SVPNGatewaysClientBeginDisconnectP2SVPNConnectionsOptions) (resp azfake.PollerResponder[armnetwork.P2SVPNGatewaysClientDisconnectP2SVPNConnectionsResponse], errResp azfake.ErrorResponder)

	// BeginGenerateVPNProfile is the fake for method P2SVPNGatewaysClient.BeginGenerateVPNProfile
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginGenerateVPNProfile func(ctx context.Context, resourceGroupName string, gatewayName string, parameters armnetwork.P2SVPNProfileParameters, options *armnetwork.P2SVPNGatewaysClientBeginGenerateVPNProfileOptions) (resp azfake.PollerResponder[armnetwork.P2SVPNGatewaysClientGenerateVPNProfileResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method P2SVPNGatewaysClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, gatewayName string, options *armnetwork.P2SVPNGatewaysClientGetOptions) (resp azfake.Responder[armnetwork.P2SVPNGatewaysClientGetResponse], errResp azfake.ErrorResponder)

	// BeginGetP2SVPNConnectionHealth is the fake for method P2SVPNGatewaysClient.BeginGetP2SVPNConnectionHealth
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginGetP2SVPNConnectionHealth func(ctx context.Context, resourceGroupName string, gatewayName string, options *armnetwork.P2SVPNGatewaysClientBeginGetP2SVPNConnectionHealthOptions) (resp azfake.PollerResponder[armnetwork.P2SVPNGatewaysClientGetP2SVPNConnectionHealthResponse], errResp azfake.ErrorResponder)

	// BeginGetP2SVPNConnectionHealthDetailed is the fake for method P2SVPNGatewaysClient.BeginGetP2SVPNConnectionHealthDetailed
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginGetP2SVPNConnectionHealthDetailed func(ctx context.Context, resourceGroupName string, gatewayName string, request armnetwork.P2SVPNConnectionHealthRequest, options *armnetwork.P2SVPNGatewaysClientBeginGetP2SVPNConnectionHealthDetailedOptions) (resp azfake.PollerResponder[armnetwork.P2SVPNGatewaysClientGetP2SVPNConnectionHealthDetailedResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method P2SVPNGatewaysClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armnetwork.P2SVPNGatewaysClientListOptions) (resp azfake.PagerResponder[armnetwork.P2SVPNGatewaysClientListResponse])

	// NewListByResourceGroupPager is the fake for method P2SVPNGatewaysClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armnetwork.P2SVPNGatewaysClientListByResourceGroupOptions) (resp azfake.PagerResponder[armnetwork.P2SVPNGatewaysClientListByResourceGroupResponse])

	// BeginReset is the fake for method P2SVPNGatewaysClient.BeginReset
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginReset func(ctx context.Context, resourceGroupName string, gatewayName string, options *armnetwork.P2SVPNGatewaysClientBeginResetOptions) (resp azfake.PollerResponder[armnetwork.P2SVPNGatewaysClientResetResponse], errResp azfake.ErrorResponder)

	// BeginUpdateTags is the fake for method P2SVPNGatewaysClient.BeginUpdateTags
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdateTags func(ctx context.Context, resourceGroupName string, gatewayName string, p2SVPNGatewayParameters armnetwork.TagsObject, options *armnetwork.P2SVPNGatewaysClientBeginUpdateTagsOptions) (resp azfake.PollerResponder[armnetwork.P2SVPNGatewaysClientUpdateTagsResponse], errResp azfake.ErrorResponder)
}

// NewP2SVPNGatewaysServerTransport creates a new instance of P2SVPNGatewaysServerTransport with the provided implementation.
// The returned P2SVPNGatewaysServerTransport instance is connected to an instance of armnetwork.P2SVPNGatewaysClient by way of the
// undefined.Transporter field.
func NewP2SVPNGatewaysServerTransport(srv *P2SVPNGatewaysServer) *P2SVPNGatewaysServerTransport {
	return &P2SVPNGatewaysServerTransport{srv: srv}
}

// P2SVPNGatewaysServerTransport connects instances of armnetwork.P2SVPNGatewaysClient to instances of P2SVPNGatewaysServer.
// Don't use this type directly, use NewP2SVPNGatewaysServerTransport instead.
type P2SVPNGatewaysServerTransport struct {
	srv                                    *P2SVPNGatewaysServer
	beginCreateOrUpdate                    *azfake.PollerResponder[armnetwork.P2SVPNGatewaysClientCreateOrUpdateResponse]
	beginDelete                            *azfake.PollerResponder[armnetwork.P2SVPNGatewaysClientDeleteResponse]
	beginDisconnectP2SVPNConnections       *azfake.PollerResponder[armnetwork.P2SVPNGatewaysClientDisconnectP2SVPNConnectionsResponse]
	beginGenerateVPNProfile                *azfake.PollerResponder[armnetwork.P2SVPNGatewaysClientGenerateVPNProfileResponse]
	beginGetP2SVPNConnectionHealth         *azfake.PollerResponder[armnetwork.P2SVPNGatewaysClientGetP2SVPNConnectionHealthResponse]
	beginGetP2SVPNConnectionHealthDetailed *azfake.PollerResponder[armnetwork.P2SVPNGatewaysClientGetP2SVPNConnectionHealthDetailedResponse]
	newListPager                           *azfake.PagerResponder[armnetwork.P2SVPNGatewaysClientListResponse]
	newListByResourceGroupPager            *azfake.PagerResponder[armnetwork.P2SVPNGatewaysClientListByResourceGroupResponse]
	beginReset                             *azfake.PollerResponder[armnetwork.P2SVPNGatewaysClientResetResponse]
	beginUpdateTags                        *azfake.PollerResponder[armnetwork.P2SVPNGatewaysClientUpdateTagsResponse]
}

// Do implements the policy.Transporter interface for P2SVPNGatewaysServerTransport.
func (p *P2SVPNGatewaysServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "P2SVPNGatewaysClient.BeginCreateOrUpdate":
		resp, err = p.dispatchBeginCreateOrUpdate(req)
	case "P2SVPNGatewaysClient.BeginDelete":
		resp, err = p.dispatchBeginDelete(req)
	case "P2SVPNGatewaysClient.BeginDisconnectP2SVPNConnections":
		resp, err = p.dispatchBeginDisconnectP2SVPNConnections(req)
	case "P2SVPNGatewaysClient.BeginGenerateVPNProfile":
		resp, err = p.dispatchBeginGenerateVPNProfile(req)
	case "P2SVPNGatewaysClient.Get":
		resp, err = p.dispatchGet(req)
	case "P2SVPNGatewaysClient.BeginGetP2SVPNConnectionHealth":
		resp, err = p.dispatchBeginGetP2SVPNConnectionHealth(req)
	case "P2SVPNGatewaysClient.BeginGetP2SVPNConnectionHealthDetailed":
		resp, err = p.dispatchBeginGetP2SVPNConnectionHealthDetailed(req)
	case "P2SVPNGatewaysClient.NewListPager":
		resp, err = p.dispatchNewListPager(req)
	case "P2SVPNGatewaysClient.NewListByResourceGroupPager":
		resp, err = p.dispatchNewListByResourceGroupPager(req)
	case "P2SVPNGatewaysClient.BeginReset":
		resp, err = p.dispatchBeginReset(req)
	case "P2SVPNGatewaysClient.BeginUpdateTags":
		resp, err = p.dispatchBeginUpdateTags(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *P2SVPNGatewaysServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	if p.beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/p2svpnGateways/(?P<gatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.P2SVPNGateway](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		gatewayNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("gatewayName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameUnescaped, gatewayNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		p.beginCreateOrUpdate = &respr
	}

	resp, err := server.PollerResponderNext(p.beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(p.beginCreateOrUpdate) {
		p.beginCreateOrUpdate = nil
	}

	return resp, nil
}

func (p *P2SVPNGatewaysServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if p.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	if p.beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/p2svpnGateways/(?P<gatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		gatewayNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("gatewayName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginDelete(req.Context(), resourceGroupNameUnescaped, gatewayNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		p.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(p.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(p.beginDelete) {
		p.beginDelete = nil
	}

	return resp, nil
}

func (p *P2SVPNGatewaysServerTransport) dispatchBeginDisconnectP2SVPNConnections(req *http.Request) (*http.Response, error) {
	if p.srv.BeginDisconnectP2SVPNConnections == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDisconnectP2SVPNConnections not implemented")}
	}
	if p.beginDisconnectP2SVPNConnections == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/p2svpnGateways/(?P<p2sVpnGatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/disconnectP2sVpnConnections`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.P2SVPNConnectionRequest](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		p2SVPNGatewayNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("p2sVpnGatewayName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginDisconnectP2SVPNConnections(req.Context(), resourceGroupNameUnescaped, p2SVPNGatewayNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		p.beginDisconnectP2SVPNConnections = &respr
	}

	resp, err := server.PollerResponderNext(p.beginDisconnectP2SVPNConnections, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(p.beginDisconnectP2SVPNConnections) {
		p.beginDisconnectP2SVPNConnections = nil
	}

	return resp, nil
}

func (p *P2SVPNGatewaysServerTransport) dispatchBeginGenerateVPNProfile(req *http.Request) (*http.Response, error) {
	if p.srv.BeginGenerateVPNProfile == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginGenerateVPNProfile not implemented")}
	}
	if p.beginGenerateVPNProfile == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/p2svpnGateways/(?P<gatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/generatevpnprofile`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.P2SVPNProfileParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		gatewayNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("gatewayName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginGenerateVPNProfile(req.Context(), resourceGroupNameUnescaped, gatewayNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		p.beginGenerateVPNProfile = &respr
	}

	resp, err := server.PollerResponderNext(p.beginGenerateVPNProfile, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(p.beginGenerateVPNProfile) {
		p.beginGenerateVPNProfile = nil
	}

	return resp, nil
}

func (p *P2SVPNGatewaysServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/p2svpnGateways/(?P<gatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	gatewayNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("gatewayName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameUnescaped, gatewayNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).P2SVPNGateway, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *P2SVPNGatewaysServerTransport) dispatchBeginGetP2SVPNConnectionHealth(req *http.Request) (*http.Response, error) {
	if p.srv.BeginGetP2SVPNConnectionHealth == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginGetP2SVPNConnectionHealth not implemented")}
	}
	if p.beginGetP2SVPNConnectionHealth == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/p2svpnGateways/(?P<gatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/getP2sVpnConnectionHealth`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		gatewayNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("gatewayName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginGetP2SVPNConnectionHealth(req.Context(), resourceGroupNameUnescaped, gatewayNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		p.beginGetP2SVPNConnectionHealth = &respr
	}

	resp, err := server.PollerResponderNext(p.beginGetP2SVPNConnectionHealth, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(p.beginGetP2SVPNConnectionHealth) {
		p.beginGetP2SVPNConnectionHealth = nil
	}

	return resp, nil
}

func (p *P2SVPNGatewaysServerTransport) dispatchBeginGetP2SVPNConnectionHealthDetailed(req *http.Request) (*http.Response, error) {
	if p.srv.BeginGetP2SVPNConnectionHealthDetailed == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginGetP2SVPNConnectionHealthDetailed not implemented")}
	}
	if p.beginGetP2SVPNConnectionHealthDetailed == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/p2svpnGateways/(?P<gatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/getP2sVpnConnectionHealthDetailed`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.P2SVPNConnectionHealthRequest](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		gatewayNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("gatewayName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginGetP2SVPNConnectionHealthDetailed(req.Context(), resourceGroupNameUnescaped, gatewayNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		p.beginGetP2SVPNConnectionHealthDetailed = &respr
	}

	resp, err := server.PollerResponderNext(p.beginGetP2SVPNConnectionHealthDetailed, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(p.beginGetP2SVPNConnectionHealthDetailed) {
		p.beginGetP2SVPNConnectionHealthDetailed = nil
	}

	return resp, nil
}

func (p *P2SVPNGatewaysServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	if p.newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/p2svpnGateways`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := p.srv.NewListPager(nil)
		p.newListPager = &resp
		server.PagerResponderInjectNextLinks(p.newListPager, req, func(page *armnetwork.P2SVPNGatewaysClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(p.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(p.newListPager) {
		p.newListPager = nil
	}
	return resp, nil
}

func (p *P2SVPNGatewaysServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	if p.newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/p2svpnGateways`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListByResourceGroupPager(resourceGroupNameUnescaped, nil)
		p.newListByResourceGroupPager = &resp
		server.PagerResponderInjectNextLinks(p.newListByResourceGroupPager, req, func(page *armnetwork.P2SVPNGatewaysClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(p.newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(p.newListByResourceGroupPager) {
		p.newListByResourceGroupPager = nil
	}
	return resp, nil
}

func (p *P2SVPNGatewaysServerTransport) dispatchBeginReset(req *http.Request) (*http.Response, error) {
	if p.srv.BeginReset == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginReset not implemented")}
	}
	if p.beginReset == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/p2svpnGateways/(?P<gatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reset`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		gatewayNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("gatewayName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginReset(req.Context(), resourceGroupNameUnescaped, gatewayNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		p.beginReset = &respr
	}

	resp, err := server.PollerResponderNext(p.beginReset, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(p.beginReset) {
		p.beginReset = nil
	}

	return resp, nil
}

func (p *P2SVPNGatewaysServerTransport) dispatchBeginUpdateTags(req *http.Request) (*http.Response, error) {
	if p.srv.BeginUpdateTags == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdateTags not implemented")}
	}
	if p.beginUpdateTags == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/p2svpnGateways/(?P<gatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
		gatewayNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("gatewayName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginUpdateTags(req.Context(), resourceGroupNameUnescaped, gatewayNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		p.beginUpdateTags = &respr
	}

	resp, err := server.PollerResponderNext(p.beginUpdateTags, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(p.beginUpdateTags) {
		p.beginUpdateTags = nil
	}

	return resp, nil
}
