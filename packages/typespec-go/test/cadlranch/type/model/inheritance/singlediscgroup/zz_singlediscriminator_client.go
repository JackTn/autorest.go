// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package singlediscgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// SingleDiscriminatorClient - Illustrates inheritance with single discriminator.
// Don't use this type directly, use a constructor function instead.
type SingleDiscriminatorClient struct {
	internal *azcore.Client
}

//   - options - SingleDiscriminatorClientGetLegacyModelOptions contains the optional parameters for the SingleDiscriminatorClient.GetLegacyModel
//     method.
func (client *SingleDiscriminatorClient) GetLegacyModel(ctx context.Context, options *SingleDiscriminatorClientGetLegacyModelOptions) (SingleDiscriminatorClientGetLegacyModelResponse, error) {
	var err error
	req, err := client.getLegacyModelCreateRequest(ctx, options)
	if err != nil {
		return SingleDiscriminatorClientGetLegacyModelResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SingleDiscriminatorClientGetLegacyModelResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SingleDiscriminatorClientGetLegacyModelResponse{}, err
	}
	resp, err := client.getLegacyModelHandleResponse(httpResp)
	return resp, err
}

// getLegacyModelCreateRequest creates the GetLegacyModel request.
func (client *SingleDiscriminatorClient) getLegacyModelCreateRequest(ctx context.Context, options *SingleDiscriminatorClientGetLegacyModelOptions) (*policy.Request, error) {
	urlPath := "/type/model/inheritance/single-discriminator/legacy-model"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getLegacyModelHandleResponse handles the GetLegacyModel response.
func (client *SingleDiscriminatorClient) getLegacyModelHandleResponse(resp *http.Response) (SingleDiscriminatorClientGetLegacyModelResponse, error) {
	result := SingleDiscriminatorClientGetLegacyModelResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return SingleDiscriminatorClientGetLegacyModelResponse{}, err
	}
	return result, nil
}

//   - options - SingleDiscriminatorClientGetMissingDiscriminatorOptions contains the optional parameters for the SingleDiscriminatorClient.GetMissingDiscriminator
//     method.
func (client *SingleDiscriminatorClient) GetMissingDiscriminator(ctx context.Context, options *SingleDiscriminatorClientGetMissingDiscriminatorOptions) (SingleDiscriminatorClientGetMissingDiscriminatorResponse, error) {
	var err error
	req, err := client.getMissingDiscriminatorCreateRequest(ctx, options)
	if err != nil {
		return SingleDiscriminatorClientGetMissingDiscriminatorResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SingleDiscriminatorClientGetMissingDiscriminatorResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SingleDiscriminatorClientGetMissingDiscriminatorResponse{}, err
	}
	resp, err := client.getMissingDiscriminatorHandleResponse(httpResp)
	return resp, err
}

// getMissingDiscriminatorCreateRequest creates the GetMissingDiscriminator request.
func (client *SingleDiscriminatorClient) getMissingDiscriminatorCreateRequest(ctx context.Context, options *SingleDiscriminatorClientGetMissingDiscriminatorOptions) (*policy.Request, error) {
	urlPath := "/type/model/inheritance/single-discriminator/missingdiscriminator"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getMissingDiscriminatorHandleResponse handles the GetMissingDiscriminator response.
func (client *SingleDiscriminatorClient) getMissingDiscriminatorHandleResponse(resp *http.Response) (SingleDiscriminatorClientGetMissingDiscriminatorResponse, error) {
	result := SingleDiscriminatorClientGetMissingDiscriminatorResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return SingleDiscriminatorClientGetMissingDiscriminatorResponse{}, err
	}
	return result, nil
}

//   - options - SingleDiscriminatorClientGetModelOptions contains the optional parameters for the SingleDiscriminatorClient.GetModel
//     method.
func (client *SingleDiscriminatorClient) GetModel(ctx context.Context, options *SingleDiscriminatorClientGetModelOptions) (SingleDiscriminatorClientGetModelResponse, error) {
	var err error
	req, err := client.getModelCreateRequest(ctx, options)
	if err != nil {
		return SingleDiscriminatorClientGetModelResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SingleDiscriminatorClientGetModelResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SingleDiscriminatorClientGetModelResponse{}, err
	}
	resp, err := client.getModelHandleResponse(httpResp)
	return resp, err
}

// getModelCreateRequest creates the GetModel request.
func (client *SingleDiscriminatorClient) getModelCreateRequest(ctx context.Context, options *SingleDiscriminatorClientGetModelOptions) (*policy.Request, error) {
	urlPath := "/type/model/inheritance/single-discriminator/model"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getModelHandleResponse handles the GetModel response.
func (client *SingleDiscriminatorClient) getModelHandleResponse(resp *http.Response) (SingleDiscriminatorClientGetModelResponse, error) {
	result := SingleDiscriminatorClientGetModelResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return SingleDiscriminatorClientGetModelResponse{}, err
	}
	return result, nil
}

//   - options - SingleDiscriminatorClientGetRecursiveModelOptions contains the optional parameters for the SingleDiscriminatorClient.GetRecursiveModel
//     method.
func (client *SingleDiscriminatorClient) GetRecursiveModel(ctx context.Context, options *SingleDiscriminatorClientGetRecursiveModelOptions) (SingleDiscriminatorClientGetRecursiveModelResponse, error) {
	var err error
	req, err := client.getRecursiveModelCreateRequest(ctx, options)
	if err != nil {
		return SingleDiscriminatorClientGetRecursiveModelResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SingleDiscriminatorClientGetRecursiveModelResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SingleDiscriminatorClientGetRecursiveModelResponse{}, err
	}
	resp, err := client.getRecursiveModelHandleResponse(httpResp)
	return resp, err
}

// getRecursiveModelCreateRequest creates the GetRecursiveModel request.
func (client *SingleDiscriminatorClient) getRecursiveModelCreateRequest(ctx context.Context, options *SingleDiscriminatorClientGetRecursiveModelOptions) (*policy.Request, error) {
	urlPath := "/type/model/inheritance/single-discriminator/recursivemodel"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getRecursiveModelHandleResponse handles the GetRecursiveModel response.
func (client *SingleDiscriminatorClient) getRecursiveModelHandleResponse(resp *http.Response) (SingleDiscriminatorClientGetRecursiveModelResponse, error) {
	result := SingleDiscriminatorClientGetRecursiveModelResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return SingleDiscriminatorClientGetRecursiveModelResponse{}, err
	}
	return result, nil
}

//   - options - SingleDiscriminatorClientGetWrongDiscriminatorOptions contains the optional parameters for the SingleDiscriminatorClient.GetWrongDiscriminator
//     method.
func (client *SingleDiscriminatorClient) GetWrongDiscriminator(ctx context.Context, options *SingleDiscriminatorClientGetWrongDiscriminatorOptions) (SingleDiscriminatorClientGetWrongDiscriminatorResponse, error) {
	var err error
	req, err := client.getWrongDiscriminatorCreateRequest(ctx, options)
	if err != nil {
		return SingleDiscriminatorClientGetWrongDiscriminatorResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SingleDiscriminatorClientGetWrongDiscriminatorResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SingleDiscriminatorClientGetWrongDiscriminatorResponse{}, err
	}
	resp, err := client.getWrongDiscriminatorHandleResponse(httpResp)
	return resp, err
}

// getWrongDiscriminatorCreateRequest creates the GetWrongDiscriminator request.
func (client *SingleDiscriminatorClient) getWrongDiscriminatorCreateRequest(ctx context.Context, options *SingleDiscriminatorClientGetWrongDiscriminatorOptions) (*policy.Request, error) {
	urlPath := "/type/model/inheritance/single-discriminator/wrongdiscriminator"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getWrongDiscriminatorHandleResponse handles the GetWrongDiscriminator response.
func (client *SingleDiscriminatorClient) getWrongDiscriminatorHandleResponse(resp *http.Response) (SingleDiscriminatorClientGetWrongDiscriminatorResponse, error) {
	result := SingleDiscriminatorClientGetWrongDiscriminatorResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return SingleDiscriminatorClientGetWrongDiscriminatorResponse{}, err
	}
	return result, nil
}

//   - options - SingleDiscriminatorClientPutModelOptions contains the optional parameters for the SingleDiscriminatorClient.PutModel
//     method.
func (client *SingleDiscriminatorClient) PutModel(ctx context.Context, input BirdClassification, options *SingleDiscriminatorClientPutModelOptions) (SingleDiscriminatorClientPutModelResponse, error) {
	var err error
	req, err := client.putModelCreateRequest(ctx, input, options)
	if err != nil {
		return SingleDiscriminatorClientPutModelResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SingleDiscriminatorClientPutModelResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return SingleDiscriminatorClientPutModelResponse{}, err
	}
	return SingleDiscriminatorClientPutModelResponse{}, nil
}

// putModelCreateRequest creates the PutModel request.
func (client *SingleDiscriminatorClient) putModelCreateRequest(ctx context.Context, input BirdClassification, options *SingleDiscriminatorClientPutModelOptions) (*policy.Request, error) {
	urlPath := "/type/model/inheritance/single-discriminator/model"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, input); err != nil {
		return nil, err
	}
	return req, nil
}

//   - options - SingleDiscriminatorClientPutRecursiveModelOptions contains the optional parameters for the SingleDiscriminatorClient.PutRecursiveModel
//     method.
func (client *SingleDiscriminatorClient) PutRecursiveModel(ctx context.Context, input BirdClassification, options *SingleDiscriminatorClientPutRecursiveModelOptions) (SingleDiscriminatorClientPutRecursiveModelResponse, error) {
	var err error
	req, err := client.putRecursiveModelCreateRequest(ctx, input, options)
	if err != nil {
		return SingleDiscriminatorClientPutRecursiveModelResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SingleDiscriminatorClientPutRecursiveModelResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return SingleDiscriminatorClientPutRecursiveModelResponse{}, err
	}
	return SingleDiscriminatorClientPutRecursiveModelResponse{}, nil
}

// putRecursiveModelCreateRequest creates the PutRecursiveModel request.
func (client *SingleDiscriminatorClient) putRecursiveModelCreateRequest(ctx context.Context, input BirdClassification, options *SingleDiscriminatorClientPutRecursiveModelOptions) (*policy.Request, error) {
	urlPath := "/type/model/inheritance/single-discriminator/recursivemodel"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, input); err != nil {
		return nil, err
	}
	return req, nil
}
