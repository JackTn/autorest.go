// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"nullablegroup"
)

// DurationServer is a fake server for instances of the nullablegroup.DurationClient type.
type DurationServer struct {
	// GetNonNull is the fake for method DurationClient.GetNonNull
	// HTTP status codes to indicate success: http.StatusOK
	GetNonNull func(ctx context.Context, options *nullablegroup.DurationClientGetNonNullOptions) (resp azfake.Responder[nullablegroup.DurationClientGetNonNullResponse], errResp azfake.ErrorResponder)

	// GetNull is the fake for method DurationClient.GetNull
	// HTTP status codes to indicate success: http.StatusOK
	GetNull func(ctx context.Context, options *nullablegroup.DurationClientGetNullOptions) (resp azfake.Responder[nullablegroup.DurationClientGetNullResponse], errResp azfake.ErrorResponder)

	// PatchNonNull is the fake for method DurationClient.PatchNonNull
	// HTTP status codes to indicate success: http.StatusNoContent
	PatchNonNull func(ctx context.Context, body nullablegroup.DurationProperty, options *nullablegroup.DurationClientPatchNonNullOptions) (resp azfake.Responder[nullablegroup.DurationClientPatchNonNullResponse], errResp azfake.ErrorResponder)

	// PatchNull is the fake for method DurationClient.PatchNull
	// HTTP status codes to indicate success: http.StatusNoContent
	PatchNull func(ctx context.Context, body nullablegroup.DurationProperty, options *nullablegroup.DurationClientPatchNullOptions) (resp azfake.Responder[nullablegroup.DurationClientPatchNullResponse], errResp azfake.ErrorResponder)
}

// NewDurationServerTransport creates a new instance of DurationServerTransport with the provided implementation.
// The returned DurationServerTransport instance is connected to an instance of nullablegroup.DurationClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDurationServerTransport(srv *DurationServer) *DurationServerTransport {
	return &DurationServerTransport{srv: srv}
}

// DurationServerTransport connects instances of nullablegroup.DurationClient to instances of DurationServer.
// Don't use this type directly, use NewDurationServerTransport instead.
type DurationServerTransport struct {
	srv *DurationServer
}

// Do implements the policy.Transporter interface for DurationServerTransport.
func (d *DurationServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return d.dispatchToMethodFake(req, method)
}

func (d *DurationServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "DurationClient.GetNonNull":
		resp, err = d.dispatchGetNonNull(req)
	case "DurationClient.GetNull":
		resp, err = d.dispatchGetNull(req)
	case "DurationClient.PatchNonNull":
		resp, err = d.dispatchPatchNonNull(req)
	case "DurationClient.PatchNull":
		resp, err = d.dispatchPatchNull(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (d *DurationServerTransport) dispatchGetNonNull(req *http.Request) (*http.Response, error) {
	if d.srv.GetNonNull == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetNonNull not implemented")}
	}
	respr, errRespr := d.srv.GetNonNull(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DurationProperty, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DurationServerTransport) dispatchGetNull(req *http.Request) (*http.Response, error) {
	if d.srv.GetNull == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetNull not implemented")}
	}
	respr, errRespr := d.srv.GetNull(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DurationProperty, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DurationServerTransport) dispatchPatchNonNull(req *http.Request) (*http.Response, error) {
	if d.srv.PatchNonNull == nil {
		return nil, &nonRetriableError{errors.New("fake for method PatchNonNull not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[nullablegroup.DurationProperty](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.PatchNonNull(req.Context(), body, nil)
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

func (d *DurationServerTransport) dispatchPatchNull(req *http.Request) (*http.Response, error) {
	if d.srv.PatchNull == nil {
		return nil, &nonRetriableError{errors.New("fake for method PatchNull not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[nullablegroup.DurationProperty](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.PatchNull(req.Context(), body, nil)
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