//go:build go1.18
// +build go1.18

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

// LRORetrysServer is a fake server for instances of the lrogroup.LRORetrysClient type.
type LRORetrysServer struct {
	// BeginDelete202Retry200 is the fake for method LRORetrysClient.BeginDelete202Retry200
	// HTTP status codes to indicate success: http.StatusAccepted
	BeginDelete202Retry200 func(ctx context.Context, options *lrogroup.LRORetrysClientBeginDelete202Retry200Options) (resp azfake.PollerResponder[lrogroup.LRORetrysClientDelete202Retry200Response], errResp azfake.ErrorResponder)

	// BeginDeleteAsyncRelativeRetrySucceeded is the fake for method LRORetrysClient.BeginDeleteAsyncRelativeRetrySucceeded
	// HTTP status codes to indicate success: http.StatusAccepted
	BeginDeleteAsyncRelativeRetrySucceeded func(ctx context.Context, options *lrogroup.LRORetrysClientBeginDeleteAsyncRelativeRetrySucceededOptions) (resp azfake.PollerResponder[lrogroup.LRORetrysClientDeleteAsyncRelativeRetrySucceededResponse], errResp azfake.ErrorResponder)

	// BeginDeleteProvisioning202Accepted200Succeeded is the fake for method LRORetrysClient.BeginDeleteProvisioning202Accepted200Succeeded
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginDeleteProvisioning202Accepted200Succeeded func(ctx context.Context, options *lrogroup.LRORetrysClientBeginDeleteProvisioning202Accepted200SucceededOptions) (resp azfake.PollerResponder[lrogroup.LRORetrysClientDeleteProvisioning202Accepted200SucceededResponse], errResp azfake.ErrorResponder)

	// BeginPost202Retry200 is the fake for method LRORetrysClient.BeginPost202Retry200
	// HTTP status codes to indicate success: http.StatusAccepted
	BeginPost202Retry200 func(ctx context.Context, options *lrogroup.LRORetrysClientBeginPost202Retry200Options) (resp azfake.PollerResponder[lrogroup.LRORetrysClientPost202Retry200Response], errResp azfake.ErrorResponder)

	// BeginPostAsyncRelativeRetrySucceeded is the fake for method LRORetrysClient.BeginPostAsyncRelativeRetrySucceeded
	// HTTP status codes to indicate success: http.StatusAccepted
	BeginPostAsyncRelativeRetrySucceeded func(ctx context.Context, options *lrogroup.LRORetrysClientBeginPostAsyncRelativeRetrySucceededOptions) (resp azfake.PollerResponder[lrogroup.LRORetrysClientPostAsyncRelativeRetrySucceededResponse], errResp azfake.ErrorResponder)

	// BeginPut201CreatingSucceeded200 is the fake for method LRORetrysClient.BeginPut201CreatingSucceeded200
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginPut201CreatingSucceeded200 func(ctx context.Context, product lrogroup.Product, options *lrogroup.LRORetrysClientBeginPut201CreatingSucceeded200Options) (resp azfake.PollerResponder[lrogroup.LRORetrysClientPut201CreatingSucceeded200Response], errResp azfake.ErrorResponder)

	// BeginPutAsyncRelativeRetrySucceeded is the fake for method LRORetrysClient.BeginPutAsyncRelativeRetrySucceeded
	// HTTP status codes to indicate success: http.StatusOK
	BeginPutAsyncRelativeRetrySucceeded func(ctx context.Context, product lrogroup.Product, options *lrogroup.LRORetrysClientBeginPutAsyncRelativeRetrySucceededOptions) (resp azfake.PollerResponder[lrogroup.LRORetrysClientPutAsyncRelativeRetrySucceededResponse], errResp azfake.ErrorResponder)
}

// NewLRORetrysServerTransport creates a new instance of LRORetrysServerTransport with the provided implementation.
// The returned LRORetrysServerTransport instance is connected to an instance of lrogroup.LRORetrysClient by way of the
// undefined.Transporter field.
func NewLRORetrysServerTransport(srv *LRORetrysServer) *LRORetrysServerTransport {
	return &LRORetrysServerTransport{srv: srv}
}

// LRORetrysServerTransport connects instances of lrogroup.LRORetrysClient to instances of LRORetrysServer.
// Don't use this type directly, use NewLRORetrysServerTransport instead.
type LRORetrysServerTransport struct {
	srv                                            *LRORetrysServer
	beginDelete202Retry200                         *azfake.PollerResponder[lrogroup.LRORetrysClientDelete202Retry200Response]
	beginDeleteAsyncRelativeRetrySucceeded         *azfake.PollerResponder[lrogroup.LRORetrysClientDeleteAsyncRelativeRetrySucceededResponse]
	beginDeleteProvisioning202Accepted200Succeeded *azfake.PollerResponder[lrogroup.LRORetrysClientDeleteProvisioning202Accepted200SucceededResponse]
	beginPost202Retry200                           *azfake.PollerResponder[lrogroup.LRORetrysClientPost202Retry200Response]
	beginPostAsyncRelativeRetrySucceeded           *azfake.PollerResponder[lrogroup.LRORetrysClientPostAsyncRelativeRetrySucceededResponse]
	beginPut201CreatingSucceeded200                *azfake.PollerResponder[lrogroup.LRORetrysClientPut201CreatingSucceeded200Response]
	beginPutAsyncRelativeRetrySucceeded            *azfake.PollerResponder[lrogroup.LRORetrysClientPutAsyncRelativeRetrySucceededResponse]
}

// Do implements the policy.Transporter interface for LRORetrysServerTransport.
func (l *LRORetrysServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "LRORetrysClient.BeginDelete202Retry200":
		resp, err = l.dispatchBeginDelete202Retry200(req)
	case "LRORetrysClient.BeginDeleteAsyncRelativeRetrySucceeded":
		resp, err = l.dispatchBeginDeleteAsyncRelativeRetrySucceeded(req)
	case "LRORetrysClient.BeginDeleteProvisioning202Accepted200Succeeded":
		resp, err = l.dispatchBeginDeleteProvisioning202Accepted200Succeeded(req)
	case "LRORetrysClient.BeginPost202Retry200":
		resp, err = l.dispatchBeginPost202Retry200(req)
	case "LRORetrysClient.BeginPostAsyncRelativeRetrySucceeded":
		resp, err = l.dispatchBeginPostAsyncRelativeRetrySucceeded(req)
	case "LRORetrysClient.BeginPut201CreatingSucceeded200":
		resp, err = l.dispatchBeginPut201CreatingSucceeded200(req)
	case "LRORetrysClient.BeginPutAsyncRelativeRetrySucceeded":
		resp, err = l.dispatchBeginPutAsyncRelativeRetrySucceeded(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (l *LRORetrysServerTransport) dispatchBeginDelete202Retry200(req *http.Request) (*http.Response, error) {
	if l.srv.BeginDelete202Retry200 == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete202Retry200 not implemented")}
	}
	if l.beginDelete202Retry200 == nil {
		respr, errRespr := l.srv.BeginDelete202Retry200(req.Context(), nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		l.beginDelete202Retry200 = &respr
	}

	resp, err := server.PollerResponderNext(l.beginDelete202Retry200, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(l.beginDelete202Retry200) {
		l.beginDelete202Retry200 = nil
	}

	return resp, nil
}

func (l *LRORetrysServerTransport) dispatchBeginDeleteAsyncRelativeRetrySucceeded(req *http.Request) (*http.Response, error) {
	if l.srv.BeginDeleteAsyncRelativeRetrySucceeded == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDeleteAsyncRelativeRetrySucceeded not implemented")}
	}
	if l.beginDeleteAsyncRelativeRetrySucceeded == nil {
		respr, errRespr := l.srv.BeginDeleteAsyncRelativeRetrySucceeded(req.Context(), nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		l.beginDeleteAsyncRelativeRetrySucceeded = &respr
	}

	resp, err := server.PollerResponderNext(l.beginDeleteAsyncRelativeRetrySucceeded, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(l.beginDeleteAsyncRelativeRetrySucceeded) {
		l.beginDeleteAsyncRelativeRetrySucceeded = nil
	}

	return resp, nil
}

func (l *LRORetrysServerTransport) dispatchBeginDeleteProvisioning202Accepted200Succeeded(req *http.Request) (*http.Response, error) {
	if l.srv.BeginDeleteProvisioning202Accepted200Succeeded == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDeleteProvisioning202Accepted200Succeeded not implemented")}
	}
	if l.beginDeleteProvisioning202Accepted200Succeeded == nil {
		respr, errRespr := l.srv.BeginDeleteProvisioning202Accepted200Succeeded(req.Context(), nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		l.beginDeleteProvisioning202Accepted200Succeeded = &respr
	}

	resp, err := server.PollerResponderNext(l.beginDeleteProvisioning202Accepted200Succeeded, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(l.beginDeleteProvisioning202Accepted200Succeeded) {
		l.beginDeleteProvisioning202Accepted200Succeeded = nil
	}

	return resp, nil
}

func (l *LRORetrysServerTransport) dispatchBeginPost202Retry200(req *http.Request) (*http.Response, error) {
	if l.srv.BeginPost202Retry200 == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginPost202Retry200 not implemented")}
	}
	if l.beginPost202Retry200 == nil {
		body, err := server.UnmarshalRequestAsJSON[lrogroup.Product](req)
		if err != nil {
			return nil, err
		}
		var options *lrogroup.LRORetrysClientBeginPost202Retry200Options
		if !reflect.ValueOf(body).IsZero() {
			options = &lrogroup.LRORetrysClientBeginPost202Retry200Options{
				Product: &body,
			}
		}
		respr, errRespr := l.srv.BeginPost202Retry200(req.Context(), options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		l.beginPost202Retry200 = &respr
	}

	resp, err := server.PollerResponderNext(l.beginPost202Retry200, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(l.beginPost202Retry200) {
		l.beginPost202Retry200 = nil
	}

	return resp, nil
}

func (l *LRORetrysServerTransport) dispatchBeginPostAsyncRelativeRetrySucceeded(req *http.Request) (*http.Response, error) {
	if l.srv.BeginPostAsyncRelativeRetrySucceeded == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginPostAsyncRelativeRetrySucceeded not implemented")}
	}
	if l.beginPostAsyncRelativeRetrySucceeded == nil {
		body, err := server.UnmarshalRequestAsJSON[lrogroup.Product](req)
		if err != nil {
			return nil, err
		}
		var options *lrogroup.LRORetrysClientBeginPostAsyncRelativeRetrySucceededOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &lrogroup.LRORetrysClientBeginPostAsyncRelativeRetrySucceededOptions{
				Product: &body,
			}
		}
		respr, errRespr := l.srv.BeginPostAsyncRelativeRetrySucceeded(req.Context(), options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		l.beginPostAsyncRelativeRetrySucceeded = &respr
	}

	resp, err := server.PollerResponderNext(l.beginPostAsyncRelativeRetrySucceeded, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(l.beginPostAsyncRelativeRetrySucceeded) {
		l.beginPostAsyncRelativeRetrySucceeded = nil
	}

	return resp, nil
}

func (l *LRORetrysServerTransport) dispatchBeginPut201CreatingSucceeded200(req *http.Request) (*http.Response, error) {
	if l.srv.BeginPut201CreatingSucceeded200 == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginPut201CreatingSucceeded200 not implemented")}
	}
	if l.beginPut201CreatingSucceeded200 == nil {
		body, err := server.UnmarshalRequestAsJSON[lrogroup.Product](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := l.srv.BeginPut201CreatingSucceeded200(req.Context(), body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		l.beginPut201CreatingSucceeded200 = &respr
	}

	resp, err := server.PollerResponderNext(l.beginPut201CreatingSucceeded200, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(l.beginPut201CreatingSucceeded200) {
		l.beginPut201CreatingSucceeded200 = nil
	}

	return resp, nil
}

func (l *LRORetrysServerTransport) dispatchBeginPutAsyncRelativeRetrySucceeded(req *http.Request) (*http.Response, error) {
	if l.srv.BeginPutAsyncRelativeRetrySucceeded == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginPutAsyncRelativeRetrySucceeded not implemented")}
	}
	if l.beginPutAsyncRelativeRetrySucceeded == nil {
		body, err := server.UnmarshalRequestAsJSON[lrogroup.Product](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := l.srv.BeginPutAsyncRelativeRetrySucceeded(req.Context(), body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		l.beginPutAsyncRelativeRetrySucceeded = &respr
	}

	resp, err := server.PollerResponderNext(l.beginPutAsyncRelativeRetrySucceeded, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PollerResponderMore(l.beginPutAsyncRelativeRetrySucceeded) {
		l.beginPutAsyncRelativeRetrySucceeded = nil
	}

	return resp, nil
}
