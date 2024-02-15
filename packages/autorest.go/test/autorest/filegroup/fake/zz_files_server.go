// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	"generatortests/filegroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// FilesServer is a fake server for instances of the filegroup.FilesClient type.
type FilesServer struct {
	// GetEmptyFile is the fake for method FilesClient.GetEmptyFile
	// HTTP status codes to indicate success: http.StatusOK
	GetEmptyFile func(ctx context.Context, options *filegroup.FilesClientGetEmptyFileOptions) (resp azfake.Responder[filegroup.FilesClientGetEmptyFileResponse], errResp azfake.ErrorResponder)

	// GetFile is the fake for method FilesClient.GetFile
	// HTTP status codes to indicate success: http.StatusOK
	GetFile func(ctx context.Context, options *filegroup.FilesClientGetFileOptions) (resp azfake.Responder[filegroup.FilesClientGetFileResponse], errResp azfake.ErrorResponder)

	// GetFileLarge is the fake for method FilesClient.GetFileLarge
	// HTTP status codes to indicate success: http.StatusOK
	GetFileLarge func(ctx context.Context, options *filegroup.FilesClientGetFileLargeOptions) (resp azfake.Responder[filegroup.FilesClientGetFileLargeResponse], errResp azfake.ErrorResponder)
}

// NewFilesServerTransport creates a new instance of FilesServerTransport with the provided implementation.
// The returned FilesServerTransport instance is connected to an instance of filegroup.FilesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewFilesServerTransport(srv *FilesServer) *FilesServerTransport {
	return &FilesServerTransport{srv: srv}
}

// FilesServerTransport connects instances of filegroup.FilesClient to instances of FilesServer.
// Don't use this type directly, use NewFilesServerTransport instead.
type FilesServerTransport struct {
	srv *FilesServer
}

// Do implements the policy.Transporter interface for FilesServerTransport.
func (f *FilesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "FilesClient.GetEmptyFile":
		resp, err = f.dispatchGetEmptyFile(req)
	case "FilesClient.GetFile":
		resp, err = f.dispatchGetFile(req)
	case "FilesClient.GetFileLarge":
		resp, err = f.dispatchGetFileLarge(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (f *FilesServerTransport) dispatchGetEmptyFile(req *http.Request) (*http.Response, error) {
	if f.srv.GetEmptyFile == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetEmptyFile not implemented")}
	}
	respr, errRespr := f.srv.GetEmptyFile(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, &server.ResponseOptions{
		Body:        server.GetResponse(respr).Body,
		ContentType: "application/octet-stream",
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (f *FilesServerTransport) dispatchGetFile(req *http.Request) (*http.Response, error) {
	if f.srv.GetFile == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetFile not implemented")}
	}
	respr, errRespr := f.srv.GetFile(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, &server.ResponseOptions{
		Body:        server.GetResponse(respr).Body,
		ContentType: "application/octet-stream",
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (f *FilesServerTransport) dispatchGetFileLarge(req *http.Request) (*http.Response, error) {
	if f.srv.GetFileLarge == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetFileLarge not implemented")}
	}
	respr, errRespr := f.srv.GetFileLarge(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, &server.ResponseOptions{
		Body:        server.GetResponse(respr).Body,
		ContentType: "application/octet-stream",
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
