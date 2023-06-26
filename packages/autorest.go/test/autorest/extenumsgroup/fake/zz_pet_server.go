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
	"generatortests/extenumsgroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
)

// PetServer is a fake server for instances of the extenumsgroup.PetClient type.
type PetServer struct {
	// AddPet is the fake for method PetClient.AddPet
	// HTTP status codes to indicate success: http.StatusOK
	AddPet func(ctx context.Context, options *extenumsgroup.PetClientAddPetOptions) (resp azfake.Responder[extenumsgroup.PetClientAddPetResponse], errResp azfake.ErrorResponder)

	// GetByPetID is the fake for method PetClient.GetByPetID
	// HTTP status codes to indicate success: http.StatusOK
	GetByPetID func(ctx context.Context, petID string, options *extenumsgroup.PetClientGetByPetIDOptions) (resp azfake.Responder[extenumsgroup.PetClientGetByPetIDResponse], errResp azfake.ErrorResponder)
}

// NewPetServerTransport creates a new instance of PetServerTransport with the provided implementation.
// The returned PetServerTransport instance is connected to an instance of extenumsgroup.PetClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPetServerTransport(srv *PetServer) *PetServerTransport {
	return &PetServerTransport{srv: srv}
}

// PetServerTransport connects instances of extenumsgroup.PetClient to instances of PetServer.
// Don't use this type directly, use NewPetServerTransport instead.
type PetServerTransport struct {
	srv *PetServer
}

// Do implements the policy.Transporter interface for PetServerTransport.
func (p *PetServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PetClient.AddPet":
		resp, err = p.dispatchAddPet(req)
	case "PetClient.GetByPetID":
		resp, err = p.dispatchGetByPetID(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PetServerTransport) dispatchAddPet(req *http.Request) (*http.Response, error) {
	if p.srv.AddPet == nil {
		return nil, &nonRetriableError{errors.New("fake for method AddPet not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[extenumsgroup.Pet](req)
	if err != nil {
		return nil, err
	}
	var options *extenumsgroup.PetClientAddPetOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &extenumsgroup.PetClientAddPetOptions{
			PetParam: &body,
		}
	}
	respr, errRespr := p.srv.AddPet(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Pet, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PetServerTransport) dispatchGetByPetID(req *http.Request) (*http.Response, error) {
	if p.srv.GetByPetID == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetByPetID not implemented")}
	}
	const regexStr = `/extensibleenums/pet/(?P<petId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	petIDUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("petId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.GetByPetID(req.Context(), petIDUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Pet, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}