// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	"generatortests/mediatypesgroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"io"
	"net/http"
	"reflect"
)

// MediaTypesServer is a fake server for instances of the mediatypesgroup.MediaTypesClient type.
type MediaTypesServer struct {
	// AnalyzeBody is the fake for method MediaTypesClient.AnalyzeBody
	// HTTP status codes to indicate success: http.StatusOK
	AnalyzeBody func(ctx context.Context, contentType mediatypesgroup.ContentType, options *mediatypesgroup.MediaTypesClientAnalyzeBodyOptions) (resp azfake.Responder[mediatypesgroup.MediaTypesClientAnalyzeBodyResponse], errResp azfake.ErrorResponder)

	// AnalyzeBodyNoAcceptHeader is the fake for method MediaTypesClient.AnalyzeBodyNoAcceptHeader
	// HTTP status codes to indicate success: http.StatusAccepted
	AnalyzeBodyNoAcceptHeader func(ctx context.Context, contentType mediatypesgroup.ContentType, options *mediatypesgroup.MediaTypesClientAnalyzeBodyNoAcceptHeaderOptions) (resp azfake.Responder[mediatypesgroup.MediaTypesClientAnalyzeBodyNoAcceptHeaderResponse], errResp azfake.ErrorResponder)

	// AnalyzeBodyNoAcceptHeaderWithJSON is the fake for method MediaTypesClient.AnalyzeBodyNoAcceptHeaderWithJSON
	// HTTP status codes to indicate success: http.StatusAccepted
	AnalyzeBodyNoAcceptHeaderWithJSON func(ctx context.Context, options *mediatypesgroup.MediaTypesClientAnalyzeBodyNoAcceptHeaderWithJSONOptions) (resp azfake.Responder[mediatypesgroup.MediaTypesClientAnalyzeBodyNoAcceptHeaderWithJSONResponse], errResp azfake.ErrorResponder)

	// AnalyzeBodyWithJSON is the fake for method MediaTypesClient.AnalyzeBodyWithJSON
	// HTTP status codes to indicate success: http.StatusOK
	AnalyzeBodyWithJSON func(ctx context.Context, options *mediatypesgroup.MediaTypesClientAnalyzeBodyWithJSONOptions) (resp azfake.Responder[mediatypesgroup.MediaTypesClientAnalyzeBodyWithJSONResponse], errResp azfake.ErrorResponder)

	// BinaryBodyWithThreeContentTypes is the fake for method MediaTypesClient.BinaryBodyWithThreeContentTypes
	// HTTP status codes to indicate success: http.StatusOK
	BinaryBodyWithThreeContentTypes func(ctx context.Context, contentType mediatypesgroup.ContentType2, message io.ReadSeekCloser, options *mediatypesgroup.MediaTypesClientBinaryBodyWithThreeContentTypesOptions) (resp azfake.Responder[mediatypesgroup.MediaTypesClientBinaryBodyWithThreeContentTypesResponse], errResp azfake.ErrorResponder)

	// BinaryBodyWithTwoContentTypes is the fake for method MediaTypesClient.BinaryBodyWithTwoContentTypes
	// HTTP status codes to indicate success: http.StatusOK
	BinaryBodyWithTwoContentTypes func(ctx context.Context, contentType mediatypesgroup.ContentType1, message io.ReadSeekCloser, options *mediatypesgroup.MediaTypesClientBinaryBodyWithTwoContentTypesOptions) (resp azfake.Responder[mediatypesgroup.MediaTypesClientBinaryBodyWithTwoContentTypesResponse], errResp azfake.ErrorResponder)

	// ContentTypeWithEncoding is the fake for method MediaTypesClient.ContentTypeWithEncoding
	// HTTP status codes to indicate success: http.StatusOK
	ContentTypeWithEncoding func(ctx context.Context, options *mediatypesgroup.MediaTypesClientContentTypeWithEncodingOptions) (resp azfake.Responder[mediatypesgroup.MediaTypesClientContentTypeWithEncodingResponse], errResp azfake.ErrorResponder)

	// PutTextAndJSONBody is the fake for method MediaTypesClient.PutTextAndJSONBody
	// HTTP status codes to indicate success: http.StatusOK
	PutTextAndJSONBody func(ctx context.Context, contentType mediatypesgroup.ContentType3, message string, options *mediatypesgroup.MediaTypesClientPutTextAndJSONBodyOptions) (resp azfake.Responder[mediatypesgroup.MediaTypesClientPutTextAndJSONBodyResponse], errResp azfake.ErrorResponder)
}

// NewMediaTypesServerTransport creates a new instance of MediaTypesServerTransport with the provided implementation.
// The returned MediaTypesServerTransport instance is connected to an instance of mediatypesgroup.MediaTypesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewMediaTypesServerTransport(srv *MediaTypesServer) *MediaTypesServerTransport {
	return &MediaTypesServerTransport{srv: srv}
}

// MediaTypesServerTransport connects instances of mediatypesgroup.MediaTypesClient to instances of MediaTypesServer.
// Don't use this type directly, use NewMediaTypesServerTransport instead.
type MediaTypesServerTransport struct {
	srv *MediaTypesServer
}

// Do implements the policy.Transporter interface for MediaTypesServerTransport.
func (m *MediaTypesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "MediaTypesClient.AnalyzeBody":
		resp, err = m.dispatchAnalyzeBody(req)
	case "MediaTypesClient.AnalyzeBodyNoAcceptHeader":
		resp, err = m.dispatchAnalyzeBodyNoAcceptHeader(req)
	case "MediaTypesClient.AnalyzeBodyNoAcceptHeaderWithJSON":
		resp, err = m.dispatchAnalyzeBodyNoAcceptHeaderWithJSON(req)
	case "MediaTypesClient.AnalyzeBodyWithJSON":
		resp, err = m.dispatchAnalyzeBodyWithJSON(req)
	case "MediaTypesClient.BinaryBodyWithThreeContentTypes":
		resp, err = m.dispatchBinaryBodyWithThreeContentTypes(req)
	case "MediaTypesClient.BinaryBodyWithTwoContentTypes":
		resp, err = m.dispatchBinaryBodyWithTwoContentTypes(req)
	case "MediaTypesClient.ContentTypeWithEncoding":
		resp, err = m.dispatchContentTypeWithEncoding(req)
	case "MediaTypesClient.PutTextAndJSONBody":
		resp, err = m.dispatchPutTextAndJSONBody(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *MediaTypesServerTransport) dispatchAnalyzeBody(req *http.Request) (*http.Response, error) {
	if m.srv.AnalyzeBody == nil {
		return nil, &nonRetriableError{errors.New("fake for method AnalyzeBody not implemented")}
	}
	var options *mediatypesgroup.MediaTypesClientAnalyzeBodyOptions
	if req.Body != nil {
		options = &mediatypesgroup.MediaTypesClientAnalyzeBodyOptions{
			Input: req.Body.(io.ReadSeekCloser),
		}
	}
	respr, errRespr := m.srv.AnalyzeBody(req.Context(), mediatypesgroup.ContentType(getHeaderValue(req.Header, "Content-Type")), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Value, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *MediaTypesServerTransport) dispatchAnalyzeBodyNoAcceptHeader(req *http.Request) (*http.Response, error) {
	if m.srv.AnalyzeBodyNoAcceptHeader == nil {
		return nil, &nonRetriableError{errors.New("fake for method AnalyzeBodyNoAcceptHeader not implemented")}
	}
	var options *mediatypesgroup.MediaTypesClientAnalyzeBodyNoAcceptHeaderOptions
	if req.Body != nil {
		options = &mediatypesgroup.MediaTypesClientAnalyzeBodyNoAcceptHeaderOptions{
			Input: req.Body.(io.ReadSeekCloser),
		}
	}
	respr, errRespr := m.srv.AnalyzeBodyNoAcceptHeader(req.Context(), mediatypesgroup.ContentType(getHeaderValue(req.Header, "Content-Type")), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusAccepted}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *MediaTypesServerTransport) dispatchAnalyzeBodyNoAcceptHeaderWithJSON(req *http.Request) (*http.Response, error) {
	if m.srv.AnalyzeBodyNoAcceptHeaderWithJSON == nil {
		return nil, &nonRetriableError{errors.New("fake for method AnalyzeBodyNoAcceptHeaderWithJSON not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[mediatypesgroup.SourcePath](req)
	if err != nil {
		return nil, err
	}
	var options *mediatypesgroup.MediaTypesClientAnalyzeBodyNoAcceptHeaderWithJSONOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &mediatypesgroup.MediaTypesClientAnalyzeBodyNoAcceptHeaderWithJSONOptions{
			Input: &body,
		}
	}
	respr, errRespr := m.srv.AnalyzeBodyNoAcceptHeaderWithJSON(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusAccepted}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *MediaTypesServerTransport) dispatchAnalyzeBodyWithJSON(req *http.Request) (*http.Response, error) {
	if m.srv.AnalyzeBodyWithJSON == nil {
		return nil, &nonRetriableError{errors.New("fake for method AnalyzeBodyWithJSON not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[mediatypesgroup.SourcePath](req)
	if err != nil {
		return nil, err
	}
	var options *mediatypesgroup.MediaTypesClientAnalyzeBodyWithJSONOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &mediatypesgroup.MediaTypesClientAnalyzeBodyWithJSONOptions{
			Input: &body,
		}
	}
	respr, errRespr := m.srv.AnalyzeBodyWithJSON(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Value, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *MediaTypesServerTransport) dispatchBinaryBodyWithThreeContentTypes(req *http.Request) (*http.Response, error) {
	if m.srv.BinaryBodyWithThreeContentTypes == nil {
		return nil, &nonRetriableError{errors.New("fake for method BinaryBodyWithThreeContentTypes not implemented")}
	}
	respr, errRespr := m.srv.BinaryBodyWithThreeContentTypes(req.Context(), mediatypesgroup.ContentType2(getHeaderValue(req.Header, "Content-Type")), req.Body.(io.ReadSeekCloser), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsText(respContent, server.GetResponse(respr).Value, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *MediaTypesServerTransport) dispatchBinaryBodyWithTwoContentTypes(req *http.Request) (*http.Response, error) {
	if m.srv.BinaryBodyWithTwoContentTypes == nil {
		return nil, &nonRetriableError{errors.New("fake for method BinaryBodyWithTwoContentTypes not implemented")}
	}
	respr, errRespr := m.srv.BinaryBodyWithTwoContentTypes(req.Context(), mediatypesgroup.ContentType1(getHeaderValue(req.Header, "Content-Type")), req.Body.(io.ReadSeekCloser), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsText(respContent, server.GetResponse(respr).Value, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *MediaTypesServerTransport) dispatchContentTypeWithEncoding(req *http.Request) (*http.Response, error) {
	if m.srv.ContentTypeWithEncoding == nil {
		return nil, &nonRetriableError{errors.New("fake for method ContentTypeWithEncoding not implemented")}
	}
	body, err := server.UnmarshalRequestAsText(req)
	if err != nil {
		return nil, err
	}
	var options *mediatypesgroup.MediaTypesClientContentTypeWithEncodingOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &mediatypesgroup.MediaTypesClientContentTypeWithEncodingOptions{
			Input: &body,
		}
	}
	respr, errRespr := m.srv.ContentTypeWithEncoding(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Value, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *MediaTypesServerTransport) dispatchPutTextAndJSONBody(req *http.Request) (*http.Response, error) {
	if m.srv.PutTextAndJSONBody == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutTextAndJSONBody not implemented")}
	}
	body, err := server.UnmarshalRequestAsText(req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.PutTextAndJSONBody(req.Context(), mediatypesgroup.ContentType3(getHeaderValue(req.Header, "Content-Type")), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsText(respContent, server.GetResponse(respr).Value, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
