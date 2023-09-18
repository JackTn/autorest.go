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
	"reflect"
	"regexp"
)

// VirtualNetworkGatewayConnectionsServer is a fake server for instances of the armnetwork.VirtualNetworkGatewayConnectionsClient type.
type VirtualNetworkGatewayConnectionsServer struct {
	// BeginCreateOrUpdate is the fake for method VirtualNetworkGatewayConnectionsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters armnetwork.VirtualNetworkGatewayConnection, options *armnetwork.VirtualNetworkGatewayConnectionsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.VirtualNetworkGatewayConnectionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method VirtualNetworkGatewayConnectionsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, options *armnetwork.VirtualNetworkGatewayConnectionsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.VirtualNetworkGatewayConnectionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method VirtualNetworkGatewayConnectionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, options *armnetwork.VirtualNetworkGatewayConnectionsClientGetOptions) (resp azfake.Responder[armnetwork.VirtualNetworkGatewayConnectionsClientGetResponse], errResp azfake.ErrorResponder)

	// BeginGetIkeSas is the fake for method VirtualNetworkGatewayConnectionsClient.BeginGetIkeSas
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginGetIkeSas func(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, options *armnetwork.VirtualNetworkGatewayConnectionsClientBeginGetIkeSasOptions) (resp azfake.PollerResponder[armnetwork.VirtualNetworkGatewayConnectionsClientGetIkeSasResponse], errResp azfake.ErrorResponder)

	// GetSharedKey is the fake for method VirtualNetworkGatewayConnectionsClient.GetSharedKey
	// HTTP status codes to indicate success: http.StatusOK
	GetSharedKey func(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, options *armnetwork.VirtualNetworkGatewayConnectionsClientGetSharedKeyOptions) (resp azfake.Responder[armnetwork.VirtualNetworkGatewayConnectionsClientGetSharedKeyResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method VirtualNetworkGatewayConnectionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, options *armnetwork.VirtualNetworkGatewayConnectionsClientListOptions) (resp azfake.PagerResponder[armnetwork.VirtualNetworkGatewayConnectionsClientListResponse])

	// BeginResetConnection is the fake for method VirtualNetworkGatewayConnectionsClient.BeginResetConnection
	// HTTP status codes to indicate success: http.StatusAccepted
	BeginResetConnection func(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, options *armnetwork.VirtualNetworkGatewayConnectionsClientBeginResetConnectionOptions) (resp azfake.PollerResponder[armnetwork.VirtualNetworkGatewayConnectionsClientResetConnectionResponse], errResp azfake.ErrorResponder)

	// BeginResetSharedKey is the fake for method VirtualNetworkGatewayConnectionsClient.BeginResetSharedKey
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginResetSharedKey func(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters armnetwork.ConnectionResetSharedKey, options *armnetwork.VirtualNetworkGatewayConnectionsClientBeginResetSharedKeyOptions) (resp azfake.PollerResponder[armnetwork.VirtualNetworkGatewayConnectionsClientResetSharedKeyResponse], errResp azfake.ErrorResponder)

	// BeginSetSharedKey is the fake for method VirtualNetworkGatewayConnectionsClient.BeginSetSharedKey
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginSetSharedKey func(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters armnetwork.ConnectionSharedKey, options *armnetwork.VirtualNetworkGatewayConnectionsClientBeginSetSharedKeyOptions) (resp azfake.PollerResponder[armnetwork.VirtualNetworkGatewayConnectionsClientSetSharedKeyResponse], errResp azfake.ErrorResponder)

	// BeginStartPacketCapture is the fake for method VirtualNetworkGatewayConnectionsClient.BeginStartPacketCapture
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginStartPacketCapture func(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, options *armnetwork.VirtualNetworkGatewayConnectionsClientBeginStartPacketCaptureOptions) (resp azfake.PollerResponder[armnetwork.VirtualNetworkGatewayConnectionsClientStartPacketCaptureResponse], errResp azfake.ErrorResponder)

	// BeginStopPacketCapture is the fake for method VirtualNetworkGatewayConnectionsClient.BeginStopPacketCapture
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginStopPacketCapture func(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters armnetwork.VPNPacketCaptureStopParameters, options *armnetwork.VirtualNetworkGatewayConnectionsClientBeginStopPacketCaptureOptions) (resp azfake.PollerResponder[armnetwork.VirtualNetworkGatewayConnectionsClientStopPacketCaptureResponse], errResp azfake.ErrorResponder)

	// BeginUpdateTags is the fake for method VirtualNetworkGatewayConnectionsClient.BeginUpdateTags
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdateTags func(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters armnetwork.TagsObject, options *armnetwork.VirtualNetworkGatewayConnectionsClientBeginUpdateTagsOptions) (resp azfake.PollerResponder[armnetwork.VirtualNetworkGatewayConnectionsClientUpdateTagsResponse], errResp azfake.ErrorResponder)
}

// NewVirtualNetworkGatewayConnectionsServerTransport creates a new instance of VirtualNetworkGatewayConnectionsServerTransport with the provided implementation.
// The returned VirtualNetworkGatewayConnectionsServerTransport instance is connected to an instance of armnetwork.VirtualNetworkGatewayConnectionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVirtualNetworkGatewayConnectionsServerTransport(srv *VirtualNetworkGatewayConnectionsServer) *VirtualNetworkGatewayConnectionsServerTransport {
	return &VirtualNetworkGatewayConnectionsServerTransport{
		srv:                     srv,
		beginCreateOrUpdate:     newTracker[azfake.PollerResponder[armnetwork.VirtualNetworkGatewayConnectionsClientCreateOrUpdateResponse]](),
		beginDelete:             newTracker[azfake.PollerResponder[armnetwork.VirtualNetworkGatewayConnectionsClientDeleteResponse]](),
		beginGetIkeSas:          newTracker[azfake.PollerResponder[armnetwork.VirtualNetworkGatewayConnectionsClientGetIkeSasResponse]](),
		newListPager:            newTracker[azfake.PagerResponder[armnetwork.VirtualNetworkGatewayConnectionsClientListResponse]](),
		beginResetConnection:    newTracker[azfake.PollerResponder[armnetwork.VirtualNetworkGatewayConnectionsClientResetConnectionResponse]](),
		beginResetSharedKey:     newTracker[azfake.PollerResponder[armnetwork.VirtualNetworkGatewayConnectionsClientResetSharedKeyResponse]](),
		beginSetSharedKey:       newTracker[azfake.PollerResponder[armnetwork.VirtualNetworkGatewayConnectionsClientSetSharedKeyResponse]](),
		beginStartPacketCapture: newTracker[azfake.PollerResponder[armnetwork.VirtualNetworkGatewayConnectionsClientStartPacketCaptureResponse]](),
		beginStopPacketCapture:  newTracker[azfake.PollerResponder[armnetwork.VirtualNetworkGatewayConnectionsClientStopPacketCaptureResponse]](),
		beginUpdateTags:         newTracker[azfake.PollerResponder[armnetwork.VirtualNetworkGatewayConnectionsClientUpdateTagsResponse]](),
	}
}

// VirtualNetworkGatewayConnectionsServerTransport connects instances of armnetwork.VirtualNetworkGatewayConnectionsClient to instances of VirtualNetworkGatewayConnectionsServer.
// Don't use this type directly, use NewVirtualNetworkGatewayConnectionsServerTransport instead.
type VirtualNetworkGatewayConnectionsServerTransport struct {
	srv                     *VirtualNetworkGatewayConnectionsServer
	beginCreateOrUpdate     *tracker[azfake.PollerResponder[armnetwork.VirtualNetworkGatewayConnectionsClientCreateOrUpdateResponse]]
	beginDelete             *tracker[azfake.PollerResponder[armnetwork.VirtualNetworkGatewayConnectionsClientDeleteResponse]]
	beginGetIkeSas          *tracker[azfake.PollerResponder[armnetwork.VirtualNetworkGatewayConnectionsClientGetIkeSasResponse]]
	newListPager            *tracker[azfake.PagerResponder[armnetwork.VirtualNetworkGatewayConnectionsClientListResponse]]
	beginResetConnection    *tracker[azfake.PollerResponder[armnetwork.VirtualNetworkGatewayConnectionsClientResetConnectionResponse]]
	beginResetSharedKey     *tracker[azfake.PollerResponder[armnetwork.VirtualNetworkGatewayConnectionsClientResetSharedKeyResponse]]
	beginSetSharedKey       *tracker[azfake.PollerResponder[armnetwork.VirtualNetworkGatewayConnectionsClientSetSharedKeyResponse]]
	beginStartPacketCapture *tracker[azfake.PollerResponder[armnetwork.VirtualNetworkGatewayConnectionsClientStartPacketCaptureResponse]]
	beginStopPacketCapture  *tracker[azfake.PollerResponder[armnetwork.VirtualNetworkGatewayConnectionsClientStopPacketCaptureResponse]]
	beginUpdateTags         *tracker[azfake.PollerResponder[armnetwork.VirtualNetworkGatewayConnectionsClientUpdateTagsResponse]]
}

// Do implements the policy.Transporter interface for VirtualNetworkGatewayConnectionsServerTransport.
func (v *VirtualNetworkGatewayConnectionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VirtualNetworkGatewayConnectionsClient.BeginCreateOrUpdate":
		resp, err = v.dispatchBeginCreateOrUpdate(req)
	case "VirtualNetworkGatewayConnectionsClient.BeginDelete":
		resp, err = v.dispatchBeginDelete(req)
	case "VirtualNetworkGatewayConnectionsClient.Get":
		resp, err = v.dispatchGet(req)
	case "VirtualNetworkGatewayConnectionsClient.BeginGetIkeSas":
		resp, err = v.dispatchBeginGetIkeSas(req)
	case "VirtualNetworkGatewayConnectionsClient.GetSharedKey":
		resp, err = v.dispatchGetSharedKey(req)
	case "VirtualNetworkGatewayConnectionsClient.NewListPager":
		resp, err = v.dispatchNewListPager(req)
	case "VirtualNetworkGatewayConnectionsClient.BeginResetConnection":
		resp, err = v.dispatchBeginResetConnection(req)
	case "VirtualNetworkGatewayConnectionsClient.BeginResetSharedKey":
		resp, err = v.dispatchBeginResetSharedKey(req)
	case "VirtualNetworkGatewayConnectionsClient.BeginSetSharedKey":
		resp, err = v.dispatchBeginSetSharedKey(req)
	case "VirtualNetworkGatewayConnectionsClient.BeginStartPacketCapture":
		resp, err = v.dispatchBeginStartPacketCapture(req)
	case "VirtualNetworkGatewayConnectionsClient.BeginStopPacketCapture":
		resp, err = v.dispatchBeginStopPacketCapture(req)
	case "VirtualNetworkGatewayConnectionsClient.BeginUpdateTags":
		resp, err = v.dispatchBeginUpdateTags(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VirtualNetworkGatewayConnectionsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := v.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/connections/(?P<virtualNetworkGatewayConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.VirtualNetworkGatewayConnection](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualNetworkGatewayConnectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkGatewayConnectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameUnescaped, virtualNetworkGatewayConnectionNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		v.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		v.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		v.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (v *VirtualNetworkGatewayConnectionsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if v.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := v.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/connections/(?P<virtualNetworkGatewayConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualNetworkGatewayConnectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkGatewayConnectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginDelete(req.Context(), resourceGroupNameUnescaped, virtualNetworkGatewayConnectionNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		v.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		v.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		v.beginDelete.remove(req)
	}

	return resp, nil
}

func (v *VirtualNetworkGatewayConnectionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/connections/(?P<virtualNetworkGatewayConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	virtualNetworkGatewayConnectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkGatewayConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.Get(req.Context(), resourceGroupNameUnescaped, virtualNetworkGatewayConnectionNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualNetworkGatewayConnection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualNetworkGatewayConnectionsServerTransport) dispatchBeginGetIkeSas(req *http.Request) (*http.Response, error) {
	if v.srv.BeginGetIkeSas == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginGetIkeSas not implemented")}
	}
	beginGetIkeSas := v.beginGetIkeSas.get(req)
	if beginGetIkeSas == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/connections/(?P<virtualNetworkGatewayConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/getikesas`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualNetworkGatewayConnectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkGatewayConnectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginGetIkeSas(req.Context(), resourceGroupNameUnescaped, virtualNetworkGatewayConnectionNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginGetIkeSas = &respr
		v.beginGetIkeSas.add(req, beginGetIkeSas)
	}

	resp, err := server.PollerResponderNext(beginGetIkeSas, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		v.beginGetIkeSas.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginGetIkeSas) {
		v.beginGetIkeSas.remove(req)
	}

	return resp, nil
}

func (v *VirtualNetworkGatewayConnectionsServerTransport) dispatchGetSharedKey(req *http.Request) (*http.Response, error) {
	if v.srv.GetSharedKey == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetSharedKey not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/connections/(?P<virtualNetworkGatewayConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sharedkey`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	virtualNetworkGatewayConnectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkGatewayConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.GetSharedKey(req.Context(), resourceGroupNameUnescaped, virtualNetworkGatewayConnectionNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ConnectionSharedKey, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualNetworkGatewayConnectionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := v.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/connections`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := v.srv.NewListPager(resourceGroupNameUnescaped, nil)
		newListPager = &resp
		v.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.VirtualNetworkGatewayConnectionsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		v.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		v.newListPager.remove(req)
	}
	return resp, nil
}

func (v *VirtualNetworkGatewayConnectionsServerTransport) dispatchBeginResetConnection(req *http.Request) (*http.Response, error) {
	if v.srv.BeginResetConnection == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginResetConnection not implemented")}
	}
	beginResetConnection := v.beginResetConnection.get(req)
	if beginResetConnection == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/connections/(?P<virtualNetworkGatewayConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resetconnection`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualNetworkGatewayConnectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkGatewayConnectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginResetConnection(req.Context(), resourceGroupNameUnescaped, virtualNetworkGatewayConnectionNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginResetConnection = &respr
		v.beginResetConnection.add(req, beginResetConnection)
	}

	resp, err := server.PollerResponderNext(beginResetConnection, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted}, resp.StatusCode) {
		v.beginResetConnection.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginResetConnection) {
		v.beginResetConnection.remove(req)
	}

	return resp, nil
}

func (v *VirtualNetworkGatewayConnectionsServerTransport) dispatchBeginResetSharedKey(req *http.Request) (*http.Response, error) {
	if v.srv.BeginResetSharedKey == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginResetSharedKey not implemented")}
	}
	beginResetSharedKey := v.beginResetSharedKey.get(req)
	if beginResetSharedKey == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/connections/(?P<virtualNetworkGatewayConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sharedkey/reset`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.ConnectionResetSharedKey](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualNetworkGatewayConnectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkGatewayConnectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginResetSharedKey(req.Context(), resourceGroupNameUnescaped, virtualNetworkGatewayConnectionNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginResetSharedKey = &respr
		v.beginResetSharedKey.add(req, beginResetSharedKey)
	}

	resp, err := server.PollerResponderNext(beginResetSharedKey, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		v.beginResetSharedKey.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginResetSharedKey) {
		v.beginResetSharedKey.remove(req)
	}

	return resp, nil
}

func (v *VirtualNetworkGatewayConnectionsServerTransport) dispatchBeginSetSharedKey(req *http.Request) (*http.Response, error) {
	if v.srv.BeginSetSharedKey == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginSetSharedKey not implemented")}
	}
	beginSetSharedKey := v.beginSetSharedKey.get(req)
	if beginSetSharedKey == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/connections/(?P<virtualNetworkGatewayConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sharedkey`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.ConnectionSharedKey](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualNetworkGatewayConnectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkGatewayConnectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginSetSharedKey(req.Context(), resourceGroupNameUnescaped, virtualNetworkGatewayConnectionNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginSetSharedKey = &respr
		v.beginSetSharedKey.add(req, beginSetSharedKey)
	}

	resp, err := server.PollerResponderNext(beginSetSharedKey, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		v.beginSetSharedKey.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginSetSharedKey) {
		v.beginSetSharedKey.remove(req)
	}

	return resp, nil
}

func (v *VirtualNetworkGatewayConnectionsServerTransport) dispatchBeginStartPacketCapture(req *http.Request) (*http.Response, error) {
	if v.srv.BeginStartPacketCapture == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginStartPacketCapture not implemented")}
	}
	beginStartPacketCapture := v.beginStartPacketCapture.get(req)
	if beginStartPacketCapture == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/connections/(?P<virtualNetworkGatewayConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/startPacketCapture`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.VPNPacketCaptureStartParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualNetworkGatewayConnectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkGatewayConnectionName")])
		if err != nil {
			return nil, err
		}
		var options *armnetwork.VirtualNetworkGatewayConnectionsClientBeginStartPacketCaptureOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &armnetwork.VirtualNetworkGatewayConnectionsClientBeginStartPacketCaptureOptions{
				Parameters: &body,
			}
		}
		respr, errRespr := v.srv.BeginStartPacketCapture(req.Context(), resourceGroupNameUnescaped, virtualNetworkGatewayConnectionNameUnescaped, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginStartPacketCapture = &respr
		v.beginStartPacketCapture.add(req, beginStartPacketCapture)
	}

	resp, err := server.PollerResponderNext(beginStartPacketCapture, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		v.beginStartPacketCapture.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginStartPacketCapture) {
		v.beginStartPacketCapture.remove(req)
	}

	return resp, nil
}

func (v *VirtualNetworkGatewayConnectionsServerTransport) dispatchBeginStopPacketCapture(req *http.Request) (*http.Response, error) {
	if v.srv.BeginStopPacketCapture == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginStopPacketCapture not implemented")}
	}
	beginStopPacketCapture := v.beginStopPacketCapture.get(req)
	if beginStopPacketCapture == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/connections/(?P<virtualNetworkGatewayConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/stopPacketCapture`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.VPNPacketCaptureStopParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualNetworkGatewayConnectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkGatewayConnectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginStopPacketCapture(req.Context(), resourceGroupNameUnescaped, virtualNetworkGatewayConnectionNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginStopPacketCapture = &respr
		v.beginStopPacketCapture.add(req, beginStopPacketCapture)
	}

	resp, err := server.PollerResponderNext(beginStopPacketCapture, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		v.beginStopPacketCapture.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginStopPacketCapture) {
		v.beginStopPacketCapture.remove(req)
	}

	return resp, nil
}

func (v *VirtualNetworkGatewayConnectionsServerTransport) dispatchBeginUpdateTags(req *http.Request) (*http.Response, error) {
	if v.srv.BeginUpdateTags == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdateTags not implemented")}
	}
	beginUpdateTags := v.beginUpdateTags.get(req)
	if beginUpdateTags == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/connections/(?P<virtualNetworkGatewayConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
		virtualNetworkGatewayConnectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkGatewayConnectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginUpdateTags(req.Context(), resourceGroupNameUnescaped, virtualNetworkGatewayConnectionNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdateTags = &respr
		v.beginUpdateTags.add(req, beginUpdateTags)
	}

	resp, err := server.PollerResponderNext(beginUpdateTags, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		v.beginUpdateTags.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdateTags) {
		v.beginUpdateTags.remove(req)
	}

	return resp, nil
}
