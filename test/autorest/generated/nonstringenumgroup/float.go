// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package nonstringenumgroup

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

// FloatOperations contains the methods for the Float group.
type FloatOperations interface {
	// Get - Get a float enum
	Get(ctx context.Context) (*FloatEnumResponse, error)
	// Put - Put a float enum
	Put(ctx context.Context, floatPutOptions *FloatPutOptions) (*StringResponse, error)
}

// floatOperations implements the FloatOperations interface.
type floatOperations struct {
	*Client
}

// Get - Get a float enum
func (client *floatOperations) Get(ctx context.Context) (*FloatEnumResponse, error) {
	req, err := client.getCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client *floatOperations) getCreateRequest() (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/nonStringEnums/float/get"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *floatOperations) getHandleResponse(resp *azcore.Response) (*FloatEnumResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := FloatEnumResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// getHandleError handles the Get error response.
func (client *floatOperations) getHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}

// Put - Put a float enum
func (client *floatOperations) Put(ctx context.Context, floatPutOptions *FloatPutOptions) (*StringResponse, error) {
	req, err := client.putCreateRequest(floatPutOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putCreateRequest creates the Put request.
func (client *floatOperations) putCreateRequest(floatPutOptions *FloatPutOptions) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/nonStringEnums/float/put"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	if floatPutOptions != nil {
		return req, req.MarshalAsJSON(floatPutOptions.Input)
	}
	return req, nil
}

// putHandleResponse handles the Put response.
func (client *floatOperations) putHandleResponse(resp *azcore.Response) (*StringResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putHandleError(resp)
	}
	result := StringResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// putHandleError handles the Put error response.
func (client *floatOperations) putHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}