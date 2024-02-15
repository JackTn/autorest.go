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

// VirtualRouterPeeringsServer is a fake server for instances of the armnetwork.VirtualRouterPeeringsClient type.
type VirtualRouterPeeringsServer struct {
	// BeginCreateOrUpdate is the fake for method VirtualRouterPeeringsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, virtualRouterName string, peeringName string, parameters armnetwork.VirtualRouterPeering, options *armnetwork.VirtualRouterPeeringsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.VirtualRouterPeeringsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method VirtualRouterPeeringsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, virtualRouterName string, peeringName string, options *armnetwork.VirtualRouterPeeringsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.VirtualRouterPeeringsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method VirtualRouterPeeringsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, virtualRouterName string, peeringName string, options *armnetwork.VirtualRouterPeeringsClientGetOptions) (resp azfake.Responder[armnetwork.VirtualRouterPeeringsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method VirtualRouterPeeringsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, virtualRouterName string, options *armnetwork.VirtualRouterPeeringsClientListOptions) (resp azfake.PagerResponder[armnetwork.VirtualRouterPeeringsClientListResponse])
}

// NewVirtualRouterPeeringsServerTransport creates a new instance of VirtualRouterPeeringsServerTransport with the provided implementation.
// The returned VirtualRouterPeeringsServerTransport instance is connected to an instance of armnetwork.VirtualRouterPeeringsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVirtualRouterPeeringsServerTransport(srv *VirtualRouterPeeringsServer) *VirtualRouterPeeringsServerTransport {
	return &VirtualRouterPeeringsServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armnetwork.VirtualRouterPeeringsClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armnetwork.VirtualRouterPeeringsClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armnetwork.VirtualRouterPeeringsClientListResponse]](),
	}
}

// VirtualRouterPeeringsServerTransport connects instances of armnetwork.VirtualRouterPeeringsClient to instances of VirtualRouterPeeringsServer.
// Don't use this type directly, use NewVirtualRouterPeeringsServerTransport instead.
type VirtualRouterPeeringsServerTransport struct {
	srv                 *VirtualRouterPeeringsServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armnetwork.VirtualRouterPeeringsClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armnetwork.VirtualRouterPeeringsClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armnetwork.VirtualRouterPeeringsClientListResponse]]
}

// Do implements the policy.Transporter interface for VirtualRouterPeeringsServerTransport.
func (v *VirtualRouterPeeringsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VirtualRouterPeeringsClient.BeginCreateOrUpdate":
		resp, err = v.dispatchBeginCreateOrUpdate(req)
	case "VirtualRouterPeeringsClient.BeginDelete":
		resp, err = v.dispatchBeginDelete(req)
	case "VirtualRouterPeeringsClient.Get":
		resp, err = v.dispatchGet(req)
	case "VirtualRouterPeeringsClient.NewListPager":
		resp, err = v.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VirtualRouterPeeringsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := v.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/virtualRouters/(?P<virtualRouterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/peerings/(?P<peeringName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.VirtualRouterPeering](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualRouterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualRouterName")])
		if err != nil {
			return nil, err
		}
		peeringNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("peeringName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, virtualRouterNameParam, peeringNameParam, body, nil)
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

func (v *VirtualRouterPeeringsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if v.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := v.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/virtualRouters/(?P<virtualRouterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/peerings/(?P<peeringName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualRouterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualRouterName")])
		if err != nil {
			return nil, err
		}
		peeringNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("peeringName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginDelete(req.Context(), resourceGroupNameParam, virtualRouterNameParam, peeringNameParam, nil)
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

func (v *VirtualRouterPeeringsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/virtualRouters/(?P<virtualRouterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/peerings/(?P<peeringName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	virtualRouterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualRouterName")])
	if err != nil {
		return nil, err
	}
	peeringNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("peeringName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.Get(req.Context(), resourceGroupNameParam, virtualRouterNameParam, peeringNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualRouterPeering, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualRouterPeeringsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := v.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/virtualRouters/(?P<virtualRouterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/peerings`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualRouterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualRouterName")])
		if err != nil {
			return nil, err
		}
		resp := v.srv.NewListPager(resourceGroupNameParam, virtualRouterNameParam, nil)
		newListPager = &resp
		v.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.VirtualRouterPeeringsClientListResponse, createLink func() string) {
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
