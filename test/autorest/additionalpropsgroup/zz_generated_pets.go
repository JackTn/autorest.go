// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package additionalpropsgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// PetsOperations contains the methods for the Pets group.
type PetsOperations interface {
	// CreateApInProperties - Create a Pet which contains more properties than what is defined.
	CreateApInProperties(ctx context.Context, createParameters PetApInProperties, options *PetsCreateApInPropertiesOptions) (*PetApInPropertiesResponse, error)
	// CreateApInPropertiesWithApstring - Create a Pet which contains more properties than what is defined.
	CreateApInPropertiesWithApstring(ctx context.Context, createParameters PetApInPropertiesWithApstring, options *PetsCreateApInPropertiesWithApstringOptions) (*PetApInPropertiesWithApstringResponse, error)
	// CreateApObject - Create a Pet which contains more properties than what is defined.
	CreateApObject(ctx context.Context, createParameters PetApObject, options *PetsCreateApObjectOptions) (*PetApObjectResponse, error)
	// CreateApString - Create a Pet which contains more properties than what is defined.
	CreateApString(ctx context.Context, createParameters PetApString, options *PetsCreateApStringOptions) (*PetApStringResponse, error)
	// CreateApTrue - Create a Pet which contains more properties than what is defined.
	CreateApTrue(ctx context.Context, createParameters PetApTrue, options *PetsCreateApTrueOptions) (*PetApTrueResponse, error)
	// CreateCatApTrue - Create a CatAPTrue which contains more properties than what is defined.
	CreateCatApTrue(ctx context.Context, createParameters CatApTrue, options *PetsCreateCatApTrueOptions) (*CatApTrueResponse, error)
}

// PetsClient implements the PetsOperations interface.
// Don't use this type directly, use NewPetsClient() instead.
type PetsClient struct {
	con *Connection
}

// NewPetsClient creates a new instance of PetsClient with the specified values.
func NewPetsClient(con *Connection) PetsOperations {
	return &PetsClient{con: con}
}

// Pipeline returns the pipeline associated with this client.
func (client *PetsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// CreateApInProperties - Create a Pet which contains more properties than what is defined.
func (client *PetsClient) CreateApInProperties(ctx context.Context, createParameters PetApInProperties, options *PetsCreateApInPropertiesOptions) (*PetApInPropertiesResponse, error) {
	req, err := client.CreateApInPropertiesCreateRequest(ctx, createParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.CreateApInPropertiesHandleError(resp)
	}
	result, err := client.CreateApInPropertiesHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateApInPropertiesCreateRequest creates the CreateApInProperties request.
func (client *PetsClient) CreateApInPropertiesCreateRequest(ctx context.Context, createParameters PetApInProperties, options *PetsCreateApInPropertiesOptions) (*azcore.Request, error) {
	urlPath := "/additionalProperties/in/properties"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(createParameters)
}

// CreateApInPropertiesHandleResponse handles the CreateApInProperties response.
func (client *PetsClient) CreateApInPropertiesHandleResponse(resp *azcore.Response) (*PetApInPropertiesResponse, error) {
	result := PetApInPropertiesResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PetApInProperties)
}

// CreateApInPropertiesHandleError handles the CreateApInProperties error response.
func (client *PetsClient) CreateApInPropertiesHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// CreateApInPropertiesWithApstring - Create a Pet which contains more properties than what is defined.
func (client *PetsClient) CreateApInPropertiesWithApstring(ctx context.Context, createParameters PetApInPropertiesWithApstring, options *PetsCreateApInPropertiesWithApstringOptions) (*PetApInPropertiesWithApstringResponse, error) {
	req, err := client.CreateApInPropertiesWithApstringCreateRequest(ctx, createParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.CreateApInPropertiesWithApstringHandleError(resp)
	}
	result, err := client.CreateApInPropertiesWithApstringHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateApInPropertiesWithApstringCreateRequest creates the CreateApInPropertiesWithApstring request.
func (client *PetsClient) CreateApInPropertiesWithApstringCreateRequest(ctx context.Context, createParameters PetApInPropertiesWithApstring, options *PetsCreateApInPropertiesWithApstringOptions) (*azcore.Request, error) {
	urlPath := "/additionalProperties/in/properties/with/additionalProperties/string"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(createParameters)
}

// CreateApInPropertiesWithApstringHandleResponse handles the CreateApInPropertiesWithApstring response.
func (client *PetsClient) CreateApInPropertiesWithApstringHandleResponse(resp *azcore.Response) (*PetApInPropertiesWithApstringResponse, error) {
	result := PetApInPropertiesWithApstringResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PetApInPropertiesWithApstring)
}

// CreateApInPropertiesWithApstringHandleError handles the CreateApInPropertiesWithApstring error response.
func (client *PetsClient) CreateApInPropertiesWithApstringHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// CreateApObject - Create a Pet which contains more properties than what is defined.
func (client *PetsClient) CreateApObject(ctx context.Context, createParameters PetApObject, options *PetsCreateApObjectOptions) (*PetApObjectResponse, error) {
	req, err := client.CreateApObjectCreateRequest(ctx, createParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.CreateApObjectHandleError(resp)
	}
	result, err := client.CreateApObjectHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateApObjectCreateRequest creates the CreateApObject request.
func (client *PetsClient) CreateApObjectCreateRequest(ctx context.Context, createParameters PetApObject, options *PetsCreateApObjectOptions) (*azcore.Request, error) {
	urlPath := "/additionalProperties/type/object"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(createParameters)
}

// CreateApObjectHandleResponse handles the CreateApObject response.
func (client *PetsClient) CreateApObjectHandleResponse(resp *azcore.Response) (*PetApObjectResponse, error) {
	result := PetApObjectResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PetApObject)
}

// CreateApObjectHandleError handles the CreateApObject error response.
func (client *PetsClient) CreateApObjectHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// CreateApString - Create a Pet which contains more properties than what is defined.
func (client *PetsClient) CreateApString(ctx context.Context, createParameters PetApString, options *PetsCreateApStringOptions) (*PetApStringResponse, error) {
	req, err := client.CreateApStringCreateRequest(ctx, createParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.CreateApStringHandleError(resp)
	}
	result, err := client.CreateApStringHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateApStringCreateRequest creates the CreateApString request.
func (client *PetsClient) CreateApStringCreateRequest(ctx context.Context, createParameters PetApString, options *PetsCreateApStringOptions) (*azcore.Request, error) {
	urlPath := "/additionalProperties/type/string"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(createParameters)
}

// CreateApStringHandleResponse handles the CreateApString response.
func (client *PetsClient) CreateApStringHandleResponse(resp *azcore.Response) (*PetApStringResponse, error) {
	result := PetApStringResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PetApString)
}

// CreateApStringHandleError handles the CreateApString error response.
func (client *PetsClient) CreateApStringHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// CreateApTrue - Create a Pet which contains more properties than what is defined.
func (client *PetsClient) CreateApTrue(ctx context.Context, createParameters PetApTrue, options *PetsCreateApTrueOptions) (*PetApTrueResponse, error) {
	req, err := client.CreateApTrueCreateRequest(ctx, createParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.CreateApTrueHandleError(resp)
	}
	result, err := client.CreateApTrueHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateApTrueCreateRequest creates the CreateApTrue request.
func (client *PetsClient) CreateApTrueCreateRequest(ctx context.Context, createParameters PetApTrue, options *PetsCreateApTrueOptions) (*azcore.Request, error) {
	urlPath := "/additionalProperties/true"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(createParameters)
}

// CreateApTrueHandleResponse handles the CreateApTrue response.
func (client *PetsClient) CreateApTrueHandleResponse(resp *azcore.Response) (*PetApTrueResponse, error) {
	result := PetApTrueResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PetApTrue)
}

// CreateApTrueHandleError handles the CreateApTrue error response.
func (client *PetsClient) CreateApTrueHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// CreateCatApTrue - Create a CatAPTrue which contains more properties than what is defined.
func (client *PetsClient) CreateCatApTrue(ctx context.Context, createParameters CatApTrue, options *PetsCreateCatApTrueOptions) (*CatApTrueResponse, error) {
	req, err := client.CreateCatApTrueCreateRequest(ctx, createParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.CreateCatApTrueHandleError(resp)
	}
	result, err := client.CreateCatApTrueHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateCatApTrueCreateRequest creates the CreateCatApTrue request.
func (client *PetsClient) CreateCatApTrueCreateRequest(ctx context.Context, createParameters CatApTrue, options *PetsCreateCatApTrueOptions) (*azcore.Request, error) {
	urlPath := "/additionalProperties/true-subclass"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(createParameters)
}

// CreateCatApTrueHandleResponse handles the CreateCatApTrue response.
func (client *PetsClient) CreateCatApTrueHandleResponse(resp *azcore.Response) (*CatApTrueResponse, error) {
	result := CatApTrueResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.CatApTrue)
}

// CreateCatApTrueHandleError handles the CreateCatApTrue error response.
func (client *PetsClient) CreateCatApTrueHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
