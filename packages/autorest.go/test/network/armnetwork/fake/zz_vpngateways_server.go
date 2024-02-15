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

// VPNGatewaysServer is a fake server for instances of the armnetwork.VPNGatewaysClient type.
type VPNGatewaysServer struct {
	// BeginCreateOrUpdate is the fake for method VPNGatewaysClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, gatewayName string, vpnGatewayParameters armnetwork.VPNGateway, options *armnetwork.VPNGatewaysClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.VPNGatewaysClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method VPNGatewaysClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, gatewayName string, options *armnetwork.VPNGatewaysClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.VPNGatewaysClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method VPNGatewaysClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, gatewayName string, options *armnetwork.VPNGatewaysClientGetOptions) (resp azfake.Responder[armnetwork.VPNGatewaysClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method VPNGatewaysClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armnetwork.VPNGatewaysClientListOptions) (resp azfake.PagerResponder[armnetwork.VPNGatewaysClientListResponse])

	// NewListByResourceGroupPager is the fake for method VPNGatewaysClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armnetwork.VPNGatewaysClientListByResourceGroupOptions) (resp azfake.PagerResponder[armnetwork.VPNGatewaysClientListByResourceGroupResponse])

	// BeginReset is the fake for method VPNGatewaysClient.BeginReset
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginReset func(ctx context.Context, resourceGroupName string, gatewayName string, options *armnetwork.VPNGatewaysClientBeginResetOptions) (resp azfake.PollerResponder[armnetwork.VPNGatewaysClientResetResponse], errResp azfake.ErrorResponder)

	// BeginStartPacketCapture is the fake for method VPNGatewaysClient.BeginStartPacketCapture
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginStartPacketCapture func(ctx context.Context, resourceGroupName string, gatewayName string, options *armnetwork.VPNGatewaysClientBeginStartPacketCaptureOptions) (resp azfake.PollerResponder[armnetwork.VPNGatewaysClientStartPacketCaptureResponse], errResp azfake.ErrorResponder)

	// BeginStopPacketCapture is the fake for method VPNGatewaysClient.BeginStopPacketCapture
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginStopPacketCapture func(ctx context.Context, resourceGroupName string, gatewayName string, options *armnetwork.VPNGatewaysClientBeginStopPacketCaptureOptions) (resp azfake.PollerResponder[armnetwork.VPNGatewaysClientStopPacketCaptureResponse], errResp azfake.ErrorResponder)

	// BeginUpdateTags is the fake for method VPNGatewaysClient.BeginUpdateTags
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdateTags func(ctx context.Context, resourceGroupName string, gatewayName string, vpnGatewayParameters armnetwork.TagsObject, options *armnetwork.VPNGatewaysClientBeginUpdateTagsOptions) (resp azfake.PollerResponder[armnetwork.VPNGatewaysClientUpdateTagsResponse], errResp azfake.ErrorResponder)
}

// NewVPNGatewaysServerTransport creates a new instance of VPNGatewaysServerTransport with the provided implementation.
// The returned VPNGatewaysServerTransport instance is connected to an instance of armnetwork.VPNGatewaysClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVPNGatewaysServerTransport(srv *VPNGatewaysServer) *VPNGatewaysServerTransport {
	return &VPNGatewaysServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armnetwork.VPNGatewaysClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armnetwork.VPNGatewaysClientDeleteResponse]](),
		newListPager:                newTracker[azfake.PagerResponder[armnetwork.VPNGatewaysClientListResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armnetwork.VPNGatewaysClientListByResourceGroupResponse]](),
		beginReset:                  newTracker[azfake.PollerResponder[armnetwork.VPNGatewaysClientResetResponse]](),
		beginStartPacketCapture:     newTracker[azfake.PollerResponder[armnetwork.VPNGatewaysClientStartPacketCaptureResponse]](),
		beginStopPacketCapture:      newTracker[azfake.PollerResponder[armnetwork.VPNGatewaysClientStopPacketCaptureResponse]](),
		beginUpdateTags:             newTracker[azfake.PollerResponder[armnetwork.VPNGatewaysClientUpdateTagsResponse]](),
	}
}

// VPNGatewaysServerTransport connects instances of armnetwork.VPNGatewaysClient to instances of VPNGatewaysServer.
// Don't use this type directly, use NewVPNGatewaysServerTransport instead.
type VPNGatewaysServerTransport struct {
	srv                         *VPNGatewaysServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armnetwork.VPNGatewaysClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armnetwork.VPNGatewaysClientDeleteResponse]]
	newListPager                *tracker[azfake.PagerResponder[armnetwork.VPNGatewaysClientListResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armnetwork.VPNGatewaysClientListByResourceGroupResponse]]
	beginReset                  *tracker[azfake.PollerResponder[armnetwork.VPNGatewaysClientResetResponse]]
	beginStartPacketCapture     *tracker[azfake.PollerResponder[armnetwork.VPNGatewaysClientStartPacketCaptureResponse]]
	beginStopPacketCapture      *tracker[azfake.PollerResponder[armnetwork.VPNGatewaysClientStopPacketCaptureResponse]]
	beginUpdateTags             *tracker[azfake.PollerResponder[armnetwork.VPNGatewaysClientUpdateTagsResponse]]
}

// Do implements the policy.Transporter interface for VPNGatewaysServerTransport.
func (v *VPNGatewaysServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VPNGatewaysClient.BeginCreateOrUpdate":
		resp, err = v.dispatchBeginCreateOrUpdate(req)
	case "VPNGatewaysClient.BeginDelete":
		resp, err = v.dispatchBeginDelete(req)
	case "VPNGatewaysClient.Get":
		resp, err = v.dispatchGet(req)
	case "VPNGatewaysClient.NewListPager":
		resp, err = v.dispatchNewListPager(req)
	case "VPNGatewaysClient.NewListByResourceGroupPager":
		resp, err = v.dispatchNewListByResourceGroupPager(req)
	case "VPNGatewaysClient.BeginReset":
		resp, err = v.dispatchBeginReset(req)
	case "VPNGatewaysClient.BeginStartPacketCapture":
		resp, err = v.dispatchBeginStartPacketCapture(req)
	case "VPNGatewaysClient.BeginStopPacketCapture":
		resp, err = v.dispatchBeginStopPacketCapture(req)
	case "VPNGatewaysClient.BeginUpdateTags":
		resp, err = v.dispatchBeginUpdateTags(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VPNGatewaysServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := v.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/vpnGateways/(?P<gatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.VPNGateway](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		gatewayNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("gatewayName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, gatewayNameParam, body, nil)
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

func (v *VPNGatewaysServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if v.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := v.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/vpnGateways/(?P<gatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		gatewayNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("gatewayName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginDelete(req.Context(), resourceGroupNameParam, gatewayNameParam, nil)
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

func (v *VPNGatewaysServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/vpnGateways/(?P<gatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	gatewayNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("gatewayName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.Get(req.Context(), resourceGroupNameParam, gatewayNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VPNGateway, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VPNGatewaysServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := v.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/vpnGateways`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := v.srv.NewListPager(nil)
		newListPager = &resp
		v.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.VPNGatewaysClientListResponse, createLink func() string) {
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

func (v *VPNGatewaysServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := v.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/vpnGateways`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := v.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		v.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armnetwork.VPNGatewaysClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		v.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		v.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (v *VPNGatewaysServerTransport) dispatchBeginReset(req *http.Request) (*http.Response, error) {
	if v.srv.BeginReset == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginReset not implemented")}
	}
	beginReset := v.beginReset.get(req)
	if beginReset == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/vpnGateways/(?P<gatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reset`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		gatewayNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("gatewayName")])
		if err != nil {
			return nil, err
		}
		iPConfigurationIDUnescaped, err := url.QueryUnescape(qp.Get("ipConfigurationId"))
		if err != nil {
			return nil, err
		}
		iPConfigurationIDParam := getOptional(iPConfigurationIDUnescaped)
		var options *armnetwork.VPNGatewaysClientBeginResetOptions
		if iPConfigurationIDParam != nil {
			options = &armnetwork.VPNGatewaysClientBeginResetOptions{
				IPConfigurationID: iPConfigurationIDParam,
			}
		}
		respr, errRespr := v.srv.BeginReset(req.Context(), resourceGroupNameParam, gatewayNameParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginReset = &respr
		v.beginReset.add(req, beginReset)
	}

	resp, err := server.PollerResponderNext(beginReset, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		v.beginReset.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginReset) {
		v.beginReset.remove(req)
	}

	return resp, nil
}

func (v *VPNGatewaysServerTransport) dispatchBeginStartPacketCapture(req *http.Request) (*http.Response, error) {
	if v.srv.BeginStartPacketCapture == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginStartPacketCapture not implemented")}
	}
	beginStartPacketCapture := v.beginStartPacketCapture.get(req)
	if beginStartPacketCapture == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/vpnGateways/(?P<gatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/startpacketcapture`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.VPNGatewayPacketCaptureStartParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		gatewayNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("gatewayName")])
		if err != nil {
			return nil, err
		}
		var options *armnetwork.VPNGatewaysClientBeginStartPacketCaptureOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &armnetwork.VPNGatewaysClientBeginStartPacketCaptureOptions{
				Parameters: &body,
			}
		}
		respr, errRespr := v.srv.BeginStartPacketCapture(req.Context(), resourceGroupNameParam, gatewayNameParam, options)
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

func (v *VPNGatewaysServerTransport) dispatchBeginStopPacketCapture(req *http.Request) (*http.Response, error) {
	if v.srv.BeginStopPacketCapture == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginStopPacketCapture not implemented")}
	}
	beginStopPacketCapture := v.beginStopPacketCapture.get(req)
	if beginStopPacketCapture == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/vpnGateways/(?P<gatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/stoppacketcapture`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.VPNGatewayPacketCaptureStopParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		gatewayNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("gatewayName")])
		if err != nil {
			return nil, err
		}
		var options *armnetwork.VPNGatewaysClientBeginStopPacketCaptureOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &armnetwork.VPNGatewaysClientBeginStopPacketCaptureOptions{
				Parameters: &body,
			}
		}
		respr, errRespr := v.srv.BeginStopPacketCapture(req.Context(), resourceGroupNameParam, gatewayNameParam, options)
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

func (v *VPNGatewaysServerTransport) dispatchBeginUpdateTags(req *http.Request) (*http.Response, error) {
	if v.srv.BeginUpdateTags == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdateTags not implemented")}
	}
	beginUpdateTags := v.beginUpdateTags.get(req)
	if beginUpdateTags == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/vpnGateways/(?P<gatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.TagsObject](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		gatewayNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("gatewayName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginUpdateTags(req.Context(), resourceGroupNameParam, gatewayNameParam, body, nil)
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
