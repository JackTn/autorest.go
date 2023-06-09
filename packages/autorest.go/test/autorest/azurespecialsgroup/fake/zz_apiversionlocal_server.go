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
	"generatortests/azurespecialsgroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
)

// APIVersionLocalServer is a fake server for instances of the azurespecialsgroup.APIVersionLocalClient type.
type APIVersionLocalServer struct {
	// GetMethodLocalNull is the fake for method APIVersionLocalClient.GetMethodLocalNull
	// HTTP status codes to indicate success: http.StatusOK
	GetMethodLocalNull func(ctx context.Context, options *azurespecialsgroup.APIVersionLocalClientGetMethodLocalNullOptions) (resp azfake.Responder[azurespecialsgroup.APIVersionLocalClientGetMethodLocalNullResponse], errResp azfake.ErrorResponder)

	// GetMethodLocalValid is the fake for method APIVersionLocalClient.GetMethodLocalValid
	// HTTP status codes to indicate success: http.StatusOK
	GetMethodLocalValid func(ctx context.Context, options *azurespecialsgroup.APIVersionLocalClientGetMethodLocalValidOptions) (resp azfake.Responder[azurespecialsgroup.APIVersionLocalClientGetMethodLocalValidResponse], errResp azfake.ErrorResponder)

	// GetPathLocalValid is the fake for method APIVersionLocalClient.GetPathLocalValid
	// HTTP status codes to indicate success: http.StatusOK
	GetPathLocalValid func(ctx context.Context, options *azurespecialsgroup.APIVersionLocalClientGetPathLocalValidOptions) (resp azfake.Responder[azurespecialsgroup.APIVersionLocalClientGetPathLocalValidResponse], errResp azfake.ErrorResponder)

	// GetSwaggerLocalValid is the fake for method APIVersionLocalClient.GetSwaggerLocalValid
	// HTTP status codes to indicate success: http.StatusOK
	GetSwaggerLocalValid func(ctx context.Context, options *azurespecialsgroup.APIVersionLocalClientGetSwaggerLocalValidOptions) (resp azfake.Responder[azurespecialsgroup.APIVersionLocalClientGetSwaggerLocalValidResponse], errResp azfake.ErrorResponder)
}

// NewAPIVersionLocalServerTransport creates a new instance of APIVersionLocalServerTransport with the provided implementation.
// The returned APIVersionLocalServerTransport instance is connected to an instance of azurespecialsgroup.APIVersionLocalClient by way of the
// undefined.Transporter field.
func NewAPIVersionLocalServerTransport(srv *APIVersionLocalServer) *APIVersionLocalServerTransport {
	return &APIVersionLocalServerTransport{srv: srv}
}

// APIVersionLocalServerTransport connects instances of azurespecialsgroup.APIVersionLocalClient to instances of APIVersionLocalServer.
// Don't use this type directly, use NewAPIVersionLocalServerTransport instead.
type APIVersionLocalServerTransport struct {
	srv *APIVersionLocalServer
}

// Do implements the policy.Transporter interface for APIVersionLocalServerTransport.
func (a *APIVersionLocalServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "APIVersionLocalClient.GetMethodLocalNull":
		resp, err = a.dispatchGetMethodLocalNull(req)
	case "APIVersionLocalClient.GetMethodLocalValid":
		resp, err = a.dispatchGetMethodLocalValid(req)
	case "APIVersionLocalClient.GetPathLocalValid":
		resp, err = a.dispatchGetPathLocalValid(req)
	case "APIVersionLocalClient.GetSwaggerLocalValid":
		resp, err = a.dispatchGetSwaggerLocalValid(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *APIVersionLocalServerTransport) dispatchGetMethodLocalNull(req *http.Request) (*http.Response, error) {
	if a.srv.GetMethodLocalNull == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetMethodLocalNull not implemented")}
	}
	qp := req.URL.Query()
	aPIVersionUnescaped, err := url.QueryUnescape(qp.Get("api-version"))
	if err != nil {
		return nil, err
	}
	aPIVersionParam := getOptional(aPIVersionUnescaped)
	var options *azurespecialsgroup.APIVersionLocalClientGetMethodLocalNullOptions
	if aPIVersionParam != nil {
		options = &azurespecialsgroup.APIVersionLocalClientGetMethodLocalNullOptions{
			APIVersion: aPIVersionParam,
		}
	}
	respr, errRespr := a.srv.GetMethodLocalNull(req.Context(), options)
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

func (a *APIVersionLocalServerTransport) dispatchGetMethodLocalValid(req *http.Request) (*http.Response, error) {
	if a.srv.GetMethodLocalValid == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetMethodLocalValid not implemented")}
	}
	respr, errRespr := a.srv.GetMethodLocalValid(req.Context(), nil)
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

func (a *APIVersionLocalServerTransport) dispatchGetPathLocalValid(req *http.Request) (*http.Response, error) {
	if a.srv.GetPathLocalValid == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetPathLocalValid not implemented")}
	}
	respr, errRespr := a.srv.GetPathLocalValid(req.Context(), nil)
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

func (a *APIVersionLocalServerTransport) dispatchGetSwaggerLocalValid(req *http.Request) (*http.Response, error) {
	if a.srv.GetSwaggerLocalValid == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetSwaggerLocalValid not implemented")}
	}
	respr, errRespr := a.srv.GetSwaggerLocalValid(req.Context(), nil)
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
