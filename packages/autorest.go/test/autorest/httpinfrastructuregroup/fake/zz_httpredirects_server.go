// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	"generatortests/httpinfrastructuregroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// HTTPRedirectsServer is a fake server for instances of the httpinfrastructuregroup.HTTPRedirectsClient type.
type HTTPRedirectsServer struct {
	// Delete307 is the fake for method HTTPRedirectsClient.Delete307
	// HTTP status codes to indicate success: http.StatusOK
	Delete307 func(ctx context.Context, options *httpinfrastructuregroup.HTTPRedirectsClientDelete307Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPRedirectsClientDelete307Response], errResp azfake.ErrorResponder)

	// Get300 is the fake for method HTTPRedirectsClient.Get300
	// HTTP status codes to indicate success: http.StatusOK, http.StatusMultipleChoices
	Get300 func(ctx context.Context, options *httpinfrastructuregroup.HTTPRedirectsClientGet300Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPRedirectsClientGet300Response], errResp azfake.ErrorResponder)

	// Get301 is the fake for method HTTPRedirectsClient.Get301
	// HTTP status codes to indicate success: http.StatusOK
	Get301 func(ctx context.Context, options *httpinfrastructuregroup.HTTPRedirectsClientGet301Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPRedirectsClientGet301Response], errResp azfake.ErrorResponder)

	// Get302 is the fake for method HTTPRedirectsClient.Get302
	// HTTP status codes to indicate success: http.StatusOK
	Get302 func(ctx context.Context, options *httpinfrastructuregroup.HTTPRedirectsClientGet302Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPRedirectsClientGet302Response], errResp azfake.ErrorResponder)

	// Get307 is the fake for method HTTPRedirectsClient.Get307
	// HTTP status codes to indicate success: http.StatusOK
	Get307 func(ctx context.Context, options *httpinfrastructuregroup.HTTPRedirectsClientGet307Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPRedirectsClientGet307Response], errResp azfake.ErrorResponder)

	// Head300 is the fake for method HTTPRedirectsClient.Head300
	// HTTP status codes to indicate success: http.StatusOK, http.StatusMultipleChoices
	Head300 func(ctx context.Context, options *httpinfrastructuregroup.HTTPRedirectsClientHead300Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPRedirectsClientHead300Response], errResp azfake.ErrorResponder)

	// Head301 is the fake for method HTTPRedirectsClient.Head301
	// HTTP status codes to indicate success: http.StatusOK
	Head301 func(ctx context.Context, options *httpinfrastructuregroup.HTTPRedirectsClientHead301Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPRedirectsClientHead301Response], errResp azfake.ErrorResponder)

	// Head302 is the fake for method HTTPRedirectsClient.Head302
	// HTTP status codes to indicate success: http.StatusOK
	Head302 func(ctx context.Context, options *httpinfrastructuregroup.HTTPRedirectsClientHead302Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPRedirectsClientHead302Response], errResp azfake.ErrorResponder)

	// Head307 is the fake for method HTTPRedirectsClient.Head307
	// HTTP status codes to indicate success: http.StatusOK
	Head307 func(ctx context.Context, options *httpinfrastructuregroup.HTTPRedirectsClientHead307Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPRedirectsClientHead307Response], errResp azfake.ErrorResponder)

	// Options307 is the fake for method HTTPRedirectsClient.Options307
	// HTTP status codes to indicate success: http.StatusOK
	Options307 func(ctx context.Context, options *httpinfrastructuregroup.HTTPRedirectsClientOptions307Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPRedirectsClientOptions307Response], errResp azfake.ErrorResponder)

	// Patch302 is the fake for method HTTPRedirectsClient.Patch302
	// HTTP status codes to indicate success: http.StatusFound
	Patch302 func(ctx context.Context, options *httpinfrastructuregroup.HTTPRedirectsClientPatch302Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPRedirectsClientPatch302Response], errResp azfake.ErrorResponder)

	// Patch307 is the fake for method HTTPRedirectsClient.Patch307
	// HTTP status codes to indicate success: http.StatusOK
	Patch307 func(ctx context.Context, options *httpinfrastructuregroup.HTTPRedirectsClientPatch307Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPRedirectsClientPatch307Response], errResp azfake.ErrorResponder)

	// Post303 is the fake for method HTTPRedirectsClient.Post303
	// HTTP status codes to indicate success: http.StatusOK, http.StatusSeeOther
	Post303 func(ctx context.Context, options *httpinfrastructuregroup.HTTPRedirectsClientPost303Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPRedirectsClientPost303Response], errResp azfake.ErrorResponder)

	// Post307 is the fake for method HTTPRedirectsClient.Post307
	// HTTP status codes to indicate success: http.StatusOK
	Post307 func(ctx context.Context, options *httpinfrastructuregroup.HTTPRedirectsClientPost307Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPRedirectsClientPost307Response], errResp azfake.ErrorResponder)

	// Put301 is the fake for method HTTPRedirectsClient.Put301
	// HTTP status codes to indicate success: http.StatusMovedPermanently
	Put301 func(ctx context.Context, options *httpinfrastructuregroup.HTTPRedirectsClientPut301Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPRedirectsClientPut301Response], errResp azfake.ErrorResponder)

	// Put307 is the fake for method HTTPRedirectsClient.Put307
	// HTTP status codes to indicate success: http.StatusOK
	Put307 func(ctx context.Context, options *httpinfrastructuregroup.HTTPRedirectsClientPut307Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPRedirectsClientPut307Response], errResp azfake.ErrorResponder)
}

// NewHTTPRedirectsServerTransport creates a new instance of HTTPRedirectsServerTransport with the provided implementation.
// The returned HTTPRedirectsServerTransport instance is connected to an instance of httpinfrastructuregroup.HTTPRedirectsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewHTTPRedirectsServerTransport(srv *HTTPRedirectsServer) *HTTPRedirectsServerTransport {
	return &HTTPRedirectsServerTransport{srv: srv}
}

// HTTPRedirectsServerTransport connects instances of httpinfrastructuregroup.HTTPRedirectsClient to instances of HTTPRedirectsServer.
// Don't use this type directly, use NewHTTPRedirectsServerTransport instead.
type HTTPRedirectsServerTransport struct {
	srv *HTTPRedirectsServer
}

// Do implements the policy.Transporter interface for HTTPRedirectsServerTransport.
func (h *HTTPRedirectsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "HTTPRedirectsClient.Delete307":
		resp, err = h.dispatchDelete307(req)
	case "HTTPRedirectsClient.Get300":
		resp, err = h.dispatchGet300(req)
	case "HTTPRedirectsClient.Get301":
		resp, err = h.dispatchGet301(req)
	case "HTTPRedirectsClient.Get302":
		resp, err = h.dispatchGet302(req)
	case "HTTPRedirectsClient.Get307":
		resp, err = h.dispatchGet307(req)
	case "HTTPRedirectsClient.Head300":
		resp, err = h.dispatchHead300(req)
	case "HTTPRedirectsClient.Head301":
		resp, err = h.dispatchHead301(req)
	case "HTTPRedirectsClient.Head302":
		resp, err = h.dispatchHead302(req)
	case "HTTPRedirectsClient.Head307":
		resp, err = h.dispatchHead307(req)
	case "HTTPRedirectsClient.Options307":
		resp, err = h.dispatchOptions307(req)
	case "HTTPRedirectsClient.Patch302":
		resp, err = h.dispatchPatch302(req)
	case "HTTPRedirectsClient.Patch307":
		resp, err = h.dispatchPatch307(req)
	case "HTTPRedirectsClient.Post303":
		resp, err = h.dispatchPost303(req)
	case "HTTPRedirectsClient.Post307":
		resp, err = h.dispatchPost307(req)
	case "HTTPRedirectsClient.Put301":
		resp, err = h.dispatchPut301(req)
	case "HTTPRedirectsClient.Put307":
		resp, err = h.dispatchPut307(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *HTTPRedirectsServerTransport) dispatchDelete307(req *http.Request) (*http.Response, error) {
	if h.srv.Delete307 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete307 not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[bool](req)
	if err != nil {
		return nil, err
	}
	var options *httpinfrastructuregroup.HTTPRedirectsClientDelete307Options
	if !reflect.ValueOf(body).IsZero() {
		options = &httpinfrastructuregroup.HTTPRedirectsClientDelete307Options{
			BooleanValue: &body,
		}
	}
	respr, errRespr := h.srv.Delete307(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HTTPRedirectsServerTransport) dispatchGet300(req *http.Request) (*http.Response, error) {
	if h.srv.Get300 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get300 not implemented")}
	}
	respr, errRespr := h.srv.Get300(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusMultipleChoices}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusMultipleChoices", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).StringArray, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).Location; val != nil {
		resp.Header.Set("Location", "/http/success/get/200")
	}
	return resp, nil
}

func (h *HTTPRedirectsServerTransport) dispatchGet301(req *http.Request) (*http.Response, error) {
	if h.srv.Get301 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get301 not implemented")}
	}
	respr, errRespr := h.srv.Get301(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HTTPRedirectsServerTransport) dispatchGet302(req *http.Request) (*http.Response, error) {
	if h.srv.Get302 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get302 not implemented")}
	}
	respr, errRespr := h.srv.Get302(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HTTPRedirectsServerTransport) dispatchGet307(req *http.Request) (*http.Response, error) {
	if h.srv.Get307 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get307 not implemented")}
	}
	respr, errRespr := h.srv.Get307(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HTTPRedirectsServerTransport) dispatchHead300(req *http.Request) (*http.Response, error) {
	if h.srv.Head300 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Head300 not implemented")}
	}
	respr, errRespr := h.srv.Head300(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusMultipleChoices}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusMultipleChoices", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).Location; val != nil {
		resp.Header.Set("Location", "/http/success/head/200")
	}
	return resp, nil
}

func (h *HTTPRedirectsServerTransport) dispatchHead301(req *http.Request) (*http.Response, error) {
	if h.srv.Head301 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Head301 not implemented")}
	}
	respr, errRespr := h.srv.Head301(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HTTPRedirectsServerTransport) dispatchHead302(req *http.Request) (*http.Response, error) {
	if h.srv.Head302 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Head302 not implemented")}
	}
	respr, errRespr := h.srv.Head302(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HTTPRedirectsServerTransport) dispatchHead307(req *http.Request) (*http.Response, error) {
	if h.srv.Head307 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Head307 not implemented")}
	}
	respr, errRespr := h.srv.Head307(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HTTPRedirectsServerTransport) dispatchOptions307(req *http.Request) (*http.Response, error) {
	if h.srv.Options307 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Options307 not implemented")}
	}
	respr, errRespr := h.srv.Options307(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HTTPRedirectsServerTransport) dispatchPatch302(req *http.Request) (*http.Response, error) {
	if h.srv.Patch302 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Patch302 not implemented")}
	}
	respr, errRespr := h.srv.Patch302(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusFound}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusFound", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).Location; val != nil {
		resp.Header.Set("Location", "/http/failure/500")
	}
	return resp, nil
}

func (h *HTTPRedirectsServerTransport) dispatchPatch307(req *http.Request) (*http.Response, error) {
	if h.srv.Patch307 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Patch307 not implemented")}
	}
	respr, errRespr := h.srv.Patch307(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HTTPRedirectsServerTransport) dispatchPost303(req *http.Request) (*http.Response, error) {
	if h.srv.Post303 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Post303 not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[bool](req)
	if err != nil {
		return nil, err
	}
	var options *httpinfrastructuregroup.HTTPRedirectsClientPost303Options
	if !reflect.ValueOf(body).IsZero() {
		options = &httpinfrastructuregroup.HTTPRedirectsClientPost303Options{
			BooleanValue: &body,
		}
	}
	respr, errRespr := h.srv.Post303(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusSeeOther}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusSeeOther", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).Location; val != nil {
		resp.Header.Set("Location", "/http/success/get/200")
	}
	return resp, nil
}

func (h *HTTPRedirectsServerTransport) dispatchPost307(req *http.Request) (*http.Response, error) {
	if h.srv.Post307 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Post307 not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[bool](req)
	if err != nil {
		return nil, err
	}
	var options *httpinfrastructuregroup.HTTPRedirectsClientPost307Options
	if !reflect.ValueOf(body).IsZero() {
		options = &httpinfrastructuregroup.HTTPRedirectsClientPost307Options{
			BooleanValue: &body,
		}
	}
	respr, errRespr := h.srv.Post307(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HTTPRedirectsServerTransport) dispatchPut301(req *http.Request) (*http.Response, error) {
	if h.srv.Put301 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Put301 not implemented")}
	}
	respr, errRespr := h.srv.Put301(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusMovedPermanently}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusMovedPermanently", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).Location; val != nil {
		resp.Header.Set("Location", "/http/failure/500")
	}
	return resp, nil
}

func (h *HTTPRedirectsServerTransport) dispatchPut307(req *http.Request) (*http.Response, error) {
	if h.srv.Put307 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Put307 not implemented")}
	}
	respr, errRespr := h.srv.Put307(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
