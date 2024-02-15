// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	"generatortests/lrogroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// LROsCustomHeaderServer is a fake server for instances of the lrogroup.LROsCustomHeaderClient type.
type LROsCustomHeaderServer struct {
	// BeginPost202Retry200 is the fake for method LROsCustomHeaderClient.BeginPost202Retry200
	// HTTP status codes to indicate success: http.StatusAccepted
	BeginPost202Retry200 func(ctx context.Context, options *lrogroup.LROsCustomHeaderClientBeginPost202Retry200Options) (resp azfake.PollerResponder[lrogroup.LROsCustomHeaderClientPost202Retry200Response], errResp azfake.ErrorResponder)

	// BeginPostAsyncRetrySucceeded is the fake for method LROsCustomHeaderClient.BeginPostAsyncRetrySucceeded
	// HTTP status codes to indicate success: http.StatusAccepted
	BeginPostAsyncRetrySucceeded func(ctx context.Context, options *lrogroup.LROsCustomHeaderClientBeginPostAsyncRetrySucceededOptions) (resp azfake.PollerResponder[lrogroup.LROsCustomHeaderClientPostAsyncRetrySucceededResponse], errResp azfake.ErrorResponder)

	// BeginPut201CreatingSucceeded200 is the fake for method LROsCustomHeaderClient.BeginPut201CreatingSucceeded200
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginPut201CreatingSucceeded200 func(ctx context.Context, product lrogroup.Product, options *lrogroup.LROsCustomHeaderClientBeginPut201CreatingSucceeded200Options) (resp azfake.PollerResponder[lrogroup.LROsCustomHeaderClientPut201CreatingSucceeded200Response], errResp azfake.ErrorResponder)

	// BeginPutAsyncRetrySucceeded is the fake for method LROsCustomHeaderClient.BeginPutAsyncRetrySucceeded
	// HTTP status codes to indicate success: http.StatusOK
	BeginPutAsyncRetrySucceeded func(ctx context.Context, product lrogroup.Product, options *lrogroup.LROsCustomHeaderClientBeginPutAsyncRetrySucceededOptions) (resp azfake.PollerResponder[lrogroup.LROsCustomHeaderClientPutAsyncRetrySucceededResponse], errResp azfake.ErrorResponder)
}

// NewLROsCustomHeaderServerTransport creates a new instance of LROsCustomHeaderServerTransport with the provided implementation.
// The returned LROsCustomHeaderServerTransport instance is connected to an instance of lrogroup.LROsCustomHeaderClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewLROsCustomHeaderServerTransport(srv *LROsCustomHeaderServer) *LROsCustomHeaderServerTransport {
	return &LROsCustomHeaderServerTransport{
		srv:                             srv,
		beginPost202Retry200:            newTracker[azfake.PollerResponder[lrogroup.LROsCustomHeaderClientPost202Retry200Response]](),
		beginPostAsyncRetrySucceeded:    newTracker[azfake.PollerResponder[lrogroup.LROsCustomHeaderClientPostAsyncRetrySucceededResponse]](),
		beginPut201CreatingSucceeded200: newTracker[azfake.PollerResponder[lrogroup.LROsCustomHeaderClientPut201CreatingSucceeded200Response]](),
		beginPutAsyncRetrySucceeded:     newTracker[azfake.PollerResponder[lrogroup.LROsCustomHeaderClientPutAsyncRetrySucceededResponse]](),
	}
}

// LROsCustomHeaderServerTransport connects instances of lrogroup.LROsCustomHeaderClient to instances of LROsCustomHeaderServer.
// Don't use this type directly, use NewLROsCustomHeaderServerTransport instead.
type LROsCustomHeaderServerTransport struct {
	srv                             *LROsCustomHeaderServer
	beginPost202Retry200            *tracker[azfake.PollerResponder[lrogroup.LROsCustomHeaderClientPost202Retry200Response]]
	beginPostAsyncRetrySucceeded    *tracker[azfake.PollerResponder[lrogroup.LROsCustomHeaderClientPostAsyncRetrySucceededResponse]]
	beginPut201CreatingSucceeded200 *tracker[azfake.PollerResponder[lrogroup.LROsCustomHeaderClientPut201CreatingSucceeded200Response]]
	beginPutAsyncRetrySucceeded     *tracker[azfake.PollerResponder[lrogroup.LROsCustomHeaderClientPutAsyncRetrySucceededResponse]]
}

// Do implements the policy.Transporter interface for LROsCustomHeaderServerTransport.
func (l *LROsCustomHeaderServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "LROsCustomHeaderClient.BeginPost202Retry200":
		resp, err = l.dispatchBeginPost202Retry200(req)
	case "LROsCustomHeaderClient.BeginPostAsyncRetrySucceeded":
		resp, err = l.dispatchBeginPostAsyncRetrySucceeded(req)
	case "LROsCustomHeaderClient.BeginPut201CreatingSucceeded200":
		resp, err = l.dispatchBeginPut201CreatingSucceeded200(req)
	case "LROsCustomHeaderClient.BeginPutAsyncRetrySucceeded":
		resp, err = l.dispatchBeginPutAsyncRetrySucceeded(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (l *LROsCustomHeaderServerTransport) dispatchBeginPost202Retry200(req *http.Request) (*http.Response, error) {
	if l.srv.BeginPost202Retry200 == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginPost202Retry200 not implemented")}
	}
	beginPost202Retry200 := l.beginPost202Retry200.get(req)
	if beginPost202Retry200 == nil {
		body, err := server.UnmarshalRequestAsJSON[lrogroup.Product](req)
		if err != nil {
			return nil, err
		}
		var options *lrogroup.LROsCustomHeaderClientBeginPost202Retry200Options
		if !reflect.ValueOf(body).IsZero() {
			options = &lrogroup.LROsCustomHeaderClientBeginPost202Retry200Options{
				Product: &body,
			}
		}
		respr, errRespr := l.srv.BeginPost202Retry200(req.Context(), options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginPost202Retry200 = &respr
		l.beginPost202Retry200.add(req, beginPost202Retry200)
	}

	resp, err := server.PollerResponderNext(beginPost202Retry200, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted}, resp.StatusCode) {
		l.beginPost202Retry200.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginPost202Retry200) {
		l.beginPost202Retry200.remove(req)
	}

	return resp, nil
}

func (l *LROsCustomHeaderServerTransport) dispatchBeginPostAsyncRetrySucceeded(req *http.Request) (*http.Response, error) {
	if l.srv.BeginPostAsyncRetrySucceeded == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginPostAsyncRetrySucceeded not implemented")}
	}
	beginPostAsyncRetrySucceeded := l.beginPostAsyncRetrySucceeded.get(req)
	if beginPostAsyncRetrySucceeded == nil {
		body, err := server.UnmarshalRequestAsJSON[lrogroup.Product](req)
		if err != nil {
			return nil, err
		}
		var options *lrogroup.LROsCustomHeaderClientBeginPostAsyncRetrySucceededOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &lrogroup.LROsCustomHeaderClientBeginPostAsyncRetrySucceededOptions{
				Product: &body,
			}
		}
		respr, errRespr := l.srv.BeginPostAsyncRetrySucceeded(req.Context(), options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginPostAsyncRetrySucceeded = &respr
		l.beginPostAsyncRetrySucceeded.add(req, beginPostAsyncRetrySucceeded)
	}

	resp, err := server.PollerResponderNext(beginPostAsyncRetrySucceeded, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted}, resp.StatusCode) {
		l.beginPostAsyncRetrySucceeded.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginPostAsyncRetrySucceeded) {
		l.beginPostAsyncRetrySucceeded.remove(req)
	}

	return resp, nil
}

func (l *LROsCustomHeaderServerTransport) dispatchBeginPut201CreatingSucceeded200(req *http.Request) (*http.Response, error) {
	if l.srv.BeginPut201CreatingSucceeded200 == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginPut201CreatingSucceeded200 not implemented")}
	}
	beginPut201CreatingSucceeded200 := l.beginPut201CreatingSucceeded200.get(req)
	if beginPut201CreatingSucceeded200 == nil {
		body, err := server.UnmarshalRequestAsJSON[lrogroup.Product](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := l.srv.BeginPut201CreatingSucceeded200(req.Context(), body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginPut201CreatingSucceeded200 = &respr
		l.beginPut201CreatingSucceeded200.add(req, beginPut201CreatingSucceeded200)
	}

	resp, err := server.PollerResponderNext(beginPut201CreatingSucceeded200, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		l.beginPut201CreatingSucceeded200.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginPut201CreatingSucceeded200) {
		l.beginPut201CreatingSucceeded200.remove(req)
	}

	return resp, nil
}

func (l *LROsCustomHeaderServerTransport) dispatchBeginPutAsyncRetrySucceeded(req *http.Request) (*http.Response, error) {
	if l.srv.BeginPutAsyncRetrySucceeded == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginPutAsyncRetrySucceeded not implemented")}
	}
	beginPutAsyncRetrySucceeded := l.beginPutAsyncRetrySucceeded.get(req)
	if beginPutAsyncRetrySucceeded == nil {
		body, err := server.UnmarshalRequestAsJSON[lrogroup.Product](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := l.srv.BeginPutAsyncRetrySucceeded(req.Context(), body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginPutAsyncRetrySucceeded = &respr
		l.beginPutAsyncRetrySucceeded.add(req, beginPutAsyncRetrySucceeded)
	}

	resp, err := server.PollerResponderNext(beginPutAsyncRetrySucceeded, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		l.beginPutAsyncRetrySucceeded.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginPutAsyncRetrySucceeded) {
		l.beginPutAsyncRetrySucceeded.remove(req)
	}

	return resp, nil
}
