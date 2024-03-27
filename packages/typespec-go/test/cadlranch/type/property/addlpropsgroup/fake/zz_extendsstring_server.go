// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"addlpropsgroup"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ExtendsStringServer is a fake server for instances of the addlpropsgroup.ExtendsStringClient type.
type ExtendsStringServer struct {
	// Get is the fake for method ExtendsStringClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, options *addlpropsgroup.ExtendsStringClientGetOptions) (resp azfake.Responder[addlpropsgroup.ExtendsStringClientGetResponse], errResp azfake.ErrorResponder)

	// Put is the fake for method ExtendsStringClient.Put
	// HTTP status codes to indicate success: http.StatusNoContent
	Put func(ctx context.Context, body addlpropsgroup.ExtendsStringAdditionalProperties, options *addlpropsgroup.ExtendsStringClientPutOptions) (resp azfake.Responder[addlpropsgroup.ExtendsStringClientPutResponse], errResp azfake.ErrorResponder)
}

// NewExtendsStringServerTransport creates a new instance of ExtendsStringServerTransport with the provided implementation.
// The returned ExtendsStringServerTransport instance is connected to an instance of addlpropsgroup.ExtendsStringClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewExtendsStringServerTransport(srv *ExtendsStringServer) *ExtendsStringServerTransport {
	return &ExtendsStringServerTransport{srv: srv}
}

// ExtendsStringServerTransport connects instances of addlpropsgroup.ExtendsStringClient to instances of ExtendsStringServer.
// Don't use this type directly, use NewExtendsStringServerTransport instead.
type ExtendsStringServerTransport struct {
	srv *ExtendsStringServer
}

// Do implements the policy.Transporter interface for ExtendsStringServerTransport.
func (e *ExtendsStringServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return e.dispatchToMethodFake(req, method)
}

func (e *ExtendsStringServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "ExtendsStringClient.Get":
		resp, err = e.dispatchGet(req)
	case "ExtendsStringClient.Put":
		resp, err = e.dispatchPut(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (e *ExtendsStringServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if e.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	respr, errRespr := e.srv.Get(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ExtendsStringAdditionalProperties, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExtendsStringServerTransport) dispatchPut(req *http.Request) (*http.Response, error) {
	if e.srv.Put == nil {
		return nil, &nonRetriableError{errors.New("fake for method Put not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[addlpropsgroup.ExtendsStringAdditionalProperties](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Put(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}