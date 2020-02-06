// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package additionalproperties

import (
	"net/http"
	"net/url"
	"path"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

type PetsOperations struct{}

// CreateApInPropertiesCreateRequest creates the CreateApInProperties request.
func (PetsOperations) CreateApInPropertiesCreateRequest(u url.URL, createParameters PetApInProperties) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/additionalProperties/in/properties")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(createParameters)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// CreateApInPropertiesHandleResponse handles the CreateApInProperties response.
func (PetsOperations) CreateApInPropertiesHandleResponse(resp *azcore.Response) (*PetsCreateApInPropertiesResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := PetsCreateApInPropertiesResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// CreateApInPropertiesWithApstringCreateRequest creates the CreateApInPropertiesWithApstring request.
func (PetsOperations) CreateApInPropertiesWithApstringCreateRequest(u url.URL, createParameters PetApInPropertiesWithApstring) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/additionalProperties/in/properties/with/additionalProperties/string")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(createParameters)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// CreateApInPropertiesWithApstringHandleResponse handles the CreateApInPropertiesWithApstring response.
func (PetsOperations) CreateApInPropertiesWithApstringHandleResponse(resp *azcore.Response) (*PetsCreateApInPropertiesWithApstringResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := PetsCreateApInPropertiesWithApstringResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// CreateApObjectCreateRequest creates the CreateApObject request.
func (PetsOperations) CreateApObjectCreateRequest(u url.URL, createParameters PetApObject) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/additionalProperties/type/object")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(createParameters)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// CreateApObjectHandleResponse handles the CreateApObject response.
func (PetsOperations) CreateApObjectHandleResponse(resp *azcore.Response) (*PetsCreateApObjectResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := PetsCreateApObjectResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// CreateApStringCreateRequest creates the CreateApString request.
func (PetsOperations) CreateApStringCreateRequest(u url.URL, createParameters PetApString) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/additionalProperties/type/string")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(createParameters)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// CreateApStringHandleResponse handles the CreateApString response.
func (PetsOperations) CreateApStringHandleResponse(resp *azcore.Response) (*PetsCreateApStringResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := PetsCreateApStringResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// CreateApTrueCreateRequest creates the CreateApTrue request.
func (PetsOperations) CreateApTrueCreateRequest(u url.URL, createParameters PetApTrue) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/additionalProperties/true")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(createParameters)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// CreateApTrueHandleResponse handles the CreateApTrue response.
func (PetsOperations) CreateApTrueHandleResponse(resp *azcore.Response) (*PetsCreateApTrueResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := PetsCreateApTrueResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// CreateCatApTrueCreateRequest creates the CreateCatApTrue request.
func (PetsOperations) CreateCatApTrueCreateRequest(u url.URL, createParameters CatApTrue) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/additionalProperties/true-subclass")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(createParameters)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// CreateCatApTrueHandleResponse handles the CreateCatApTrue response.
func (PetsOperations) CreateCatApTrueHandleResponse(resp *azcore.Response) (*PetsCreateCatApTrueResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := PetsCreateCatApTrueResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}
