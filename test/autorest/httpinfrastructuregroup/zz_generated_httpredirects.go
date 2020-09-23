// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package httpinfrastructuregroup

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// HTTPRedirectsOperations contains the methods for the HTTPRedirects group.
type HTTPRedirectsOperations interface {
	// Delete307 - Delete redirected with 307, resulting in a 200 after redirect
	Delete307(ctx context.Context) (*http.Response, error)
	// Get300 - Return 300 status code and redirect to /http/success/200
	// Possible return types are *HTTPRedirectsGet300Response, *StringArrayResponse
	Get300(ctx context.Context) (interface{}, error)
	// Get301 - Return 301 status code and redirect to /http/success/200
	Get301(ctx context.Context) (*http.Response, error)
	// Get302 - Return 302 status code and redirect to /http/success/200
	Get302(ctx context.Context) (*http.Response, error)
	// Get307 - Redirect get with 307, resulting in a 200 success
	Get307(ctx context.Context) (*http.Response, error)
	// Head300 - Return 300 status code and redirect to /http/success/200
	Head300(ctx context.Context) (*HTTPRedirectsHead300Response, error)
	// Head301 - Return 301 status code and redirect to /http/success/200
	Head301(ctx context.Context) (*http.Response, error)
	// Head302 - Return 302 status code and redirect to /http/success/200
	Head302(ctx context.Context) (*http.Response, error)
	// Head307 - Redirect with 307, resulting in a 200 success
	Head307(ctx context.Context) (*http.Response, error)
	// Options307 - options redirected with 307, resulting in a 200 after redirect
	Options307(ctx context.Context) (*http.Response, error)
	// Patch302 - Patch true Boolean value in request returns 302.  This request should not be automatically redirected, but should return the received 302 to the caller for evaluation
	Patch302(ctx context.Context) (*HTTPRedirectsPatch302Response, error)
	// Patch307 - Patch redirected with 307, resulting in a 200 after redirect
	Patch307(ctx context.Context) (*http.Response, error)
	// Post303 - Post true Boolean value in request returns 303.  This request should be automatically redirected usign a get, ultimately returning a 200 status code
	Post303(ctx context.Context) (*HTTPRedirectsPost303Response, error)
	// Post307 - Post redirected with 307, resulting in a 200 after redirect
	Post307(ctx context.Context) (*http.Response, error)
	// Put301 - Put true Boolean value in request returns 301.  This request should not be automatically redirected, but should return the received 301 to the caller for evaluation
	Put301(ctx context.Context) (*HTTPRedirectsPut301Response, error)
	// Put307 - Put redirected with 307, resulting in a 200 after redirect
	Put307(ctx context.Context) (*http.Response, error)
}

// HTTPRedirectsClient implements the HTTPRedirectsOperations interface.
// Don't use this type directly, use NewHTTPRedirectsClient() instead.
type HTTPRedirectsClient struct {
	*Client
}

// NewHTTPRedirectsClient creates a new instance of HTTPRedirectsClient with the specified values.
func NewHTTPRedirectsClient(c *Client) HTTPRedirectsOperations {
	return &HTTPRedirectsClient{Client: c}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *HTTPRedirectsClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// Delete307 - Delete redirected with 307, resulting in a 200 after redirect
func (client *HTTPRedirectsClient) Delete307(ctx context.Context) (*http.Response, error) {
	req, err := client.Delete307CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.Delete307HandleError(resp)
	}
	return resp.Response, nil
}

// Delete307CreateRequest creates the Delete307 request.
func (client *HTTPRedirectsClient) Delete307CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/redirect/307"
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// Delete307HandleError handles the Delete307 error response.
func (client *HTTPRedirectsClient) Delete307HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get300 - Return 300 status code and redirect to /http/success/200
// Possible return types are *HTTPRedirectsGet300Response, *StringArrayResponse
func (client *HTTPRedirectsClient) Get300(ctx context.Context) (interface{}, error) {
	req, err := client.Get300CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusMultipleChoices) {
		return nil, client.Get300HandleError(resp)
	}
	result, err := client.Get300HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Get300CreateRequest creates the Get300 request.
func (client *HTTPRedirectsClient) Get300CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/redirect/300"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// Get300HandleResponse handles the Get300 response.
func (client *HTTPRedirectsClient) Get300HandleResponse(resp *azcore.Response) (interface{}, error) {
	switch resp.StatusCode {
	case http.StatusOK:
		result := HTTPRedirectsGet300Response{RawResponse: resp.Response}
		if val := resp.Header.Get("Location"); val != "" {
			result.Location = &val
		}
		return &result, nil
	case http.StatusMultipleChoices:
		result := StringArrayResponse{RawResponse: resp.Response}
		if val := resp.Header.Get("Location"); val != "" {
			result.Location = &val
		}
		return &result, resp.UnmarshalAsJSON(&result.StringArray)
	default:
		return nil, fmt.Errorf("unhandled HTTP status code %d", resp.StatusCode)
	}
}

// Get300HandleError handles the Get300 error response.
func (client *HTTPRedirectsClient) Get300HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get301 - Return 301 status code and redirect to /http/success/200
func (client *HTTPRedirectsClient) Get301(ctx context.Context) (*http.Response, error) {
	req, err := client.Get301CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.Get301HandleError(resp)
	}
	return resp.Response, nil
}

// Get301CreateRequest creates the Get301 request.
func (client *HTTPRedirectsClient) Get301CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/redirect/301"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// Get301HandleError handles the Get301 error response.
func (client *HTTPRedirectsClient) Get301HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get302 - Return 302 status code and redirect to /http/success/200
func (client *HTTPRedirectsClient) Get302(ctx context.Context) (*http.Response, error) {
	req, err := client.Get302CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.Get302HandleError(resp)
	}
	return resp.Response, nil
}

// Get302CreateRequest creates the Get302 request.
func (client *HTTPRedirectsClient) Get302CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/redirect/302"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// Get302HandleError handles the Get302 error response.
func (client *HTTPRedirectsClient) Get302HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get307 - Redirect get with 307, resulting in a 200 success
func (client *HTTPRedirectsClient) Get307(ctx context.Context) (*http.Response, error) {
	req, err := client.Get307CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.Get307HandleError(resp)
	}
	return resp.Response, nil
}

// Get307CreateRequest creates the Get307 request.
func (client *HTTPRedirectsClient) Get307CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/redirect/307"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// Get307HandleError handles the Get307 error response.
func (client *HTTPRedirectsClient) Get307HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Head300 - Return 300 status code and redirect to /http/success/200
func (client *HTTPRedirectsClient) Head300(ctx context.Context) (*HTTPRedirectsHead300Response, error) {
	req, err := client.Head300CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusMultipleChoices) {
		return nil, client.Head300HandleError(resp)
	}
	result, err := client.Head300HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Head300CreateRequest creates the Head300 request.
func (client *HTTPRedirectsClient) Head300CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/redirect/300"
	req, err := azcore.NewRequest(ctx, http.MethodHead, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// Head300HandleResponse handles the Head300 response.
func (client *HTTPRedirectsClient) Head300HandleResponse(resp *azcore.Response) (*HTTPRedirectsHead300Response, error) {
	result := HTTPRedirectsHead300Response{RawResponse: resp.Response}
	if val := resp.Header.Get("Location"); val != "" {
		result.Location = &val
	}
	return &result, nil
}

// Head300HandleError handles the Head300 error response.
func (client *HTTPRedirectsClient) Head300HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Head301 - Return 301 status code and redirect to /http/success/200
func (client *HTTPRedirectsClient) Head301(ctx context.Context) (*http.Response, error) {
	req, err := client.Head301CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.Head301HandleError(resp)
	}
	return resp.Response, nil
}

// Head301CreateRequest creates the Head301 request.
func (client *HTTPRedirectsClient) Head301CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/redirect/301"
	req, err := azcore.NewRequest(ctx, http.MethodHead, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// Head301HandleError handles the Head301 error response.
func (client *HTTPRedirectsClient) Head301HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Head302 - Return 302 status code and redirect to /http/success/200
func (client *HTTPRedirectsClient) Head302(ctx context.Context) (*http.Response, error) {
	req, err := client.Head302CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.Head302HandleError(resp)
	}
	return resp.Response, nil
}

// Head302CreateRequest creates the Head302 request.
func (client *HTTPRedirectsClient) Head302CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/redirect/302"
	req, err := azcore.NewRequest(ctx, http.MethodHead, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// Head302HandleError handles the Head302 error response.
func (client *HTTPRedirectsClient) Head302HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Head307 - Redirect with 307, resulting in a 200 success
func (client *HTTPRedirectsClient) Head307(ctx context.Context) (*http.Response, error) {
	req, err := client.Head307CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.Head307HandleError(resp)
	}
	return resp.Response, nil
}

// Head307CreateRequest creates the Head307 request.
func (client *HTTPRedirectsClient) Head307CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/redirect/307"
	req, err := azcore.NewRequest(ctx, http.MethodHead, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// Head307HandleError handles the Head307 error response.
func (client *HTTPRedirectsClient) Head307HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Options307 - options redirected with 307, resulting in a 200 after redirect
func (client *HTTPRedirectsClient) Options307(ctx context.Context) (*http.Response, error) {
	req, err := client.Options307CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.Options307HandleError(resp)
	}
	return resp.Response, nil
}

// Options307CreateRequest creates the Options307 request.
func (client *HTTPRedirectsClient) Options307CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/redirect/307"
	req, err := azcore.NewRequest(ctx, http.MethodOptions, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// Options307HandleError handles the Options307 error response.
func (client *HTTPRedirectsClient) Options307HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Patch302 - Patch true Boolean value in request returns 302.  This request should not be automatically redirected, but should return the received 302 to the caller for evaluation
func (client *HTTPRedirectsClient) Patch302(ctx context.Context) (*HTTPRedirectsPatch302Response, error) {
	req, err := client.Patch302CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusFound) {
		return nil, client.Patch302HandleError(resp)
	}
	result, err := client.Patch302HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Patch302CreateRequest creates the Patch302 request.
func (client *HTTPRedirectsClient) Patch302CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/redirect/302"
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// Patch302HandleResponse handles the Patch302 response.
func (client *HTTPRedirectsClient) Patch302HandleResponse(resp *azcore.Response) (*HTTPRedirectsPatch302Response, error) {
	result := HTTPRedirectsPatch302Response{RawResponse: resp.Response}
	if val := resp.Header.Get("Location"); val != "" {
		result.Location = &val
	}
	return &result, nil
}

// Patch302HandleError handles the Patch302 error response.
func (client *HTTPRedirectsClient) Patch302HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Patch307 - Patch redirected with 307, resulting in a 200 after redirect
func (client *HTTPRedirectsClient) Patch307(ctx context.Context) (*http.Response, error) {
	req, err := client.Patch307CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.Patch307HandleError(resp)
	}
	return resp.Response, nil
}

// Patch307CreateRequest creates the Patch307 request.
func (client *HTTPRedirectsClient) Patch307CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/redirect/307"
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// Patch307HandleError handles the Patch307 error response.
func (client *HTTPRedirectsClient) Patch307HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Post303 - Post true Boolean value in request returns 303.  This request should be automatically redirected usign a get, ultimately returning a 200 status code
func (client *HTTPRedirectsClient) Post303(ctx context.Context) (*HTTPRedirectsPost303Response, error) {
	req, err := client.Post303CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusSeeOther) {
		return nil, client.Post303HandleError(resp)
	}
	result, err := client.Post303HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Post303CreateRequest creates the Post303 request.
func (client *HTTPRedirectsClient) Post303CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/redirect/303"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// Post303HandleResponse handles the Post303 response.
func (client *HTTPRedirectsClient) Post303HandleResponse(resp *azcore.Response) (*HTTPRedirectsPost303Response, error) {
	result := HTTPRedirectsPost303Response{RawResponse: resp.Response}
	if val := resp.Header.Get("Location"); val != "" {
		result.Location = &val
	}
	return &result, nil
}

// Post303HandleError handles the Post303 error response.
func (client *HTTPRedirectsClient) Post303HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Post307 - Post redirected with 307, resulting in a 200 after redirect
func (client *HTTPRedirectsClient) Post307(ctx context.Context) (*http.Response, error) {
	req, err := client.Post307CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.Post307HandleError(resp)
	}
	return resp.Response, nil
}

// Post307CreateRequest creates the Post307 request.
func (client *HTTPRedirectsClient) Post307CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/redirect/307"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// Post307HandleError handles the Post307 error response.
func (client *HTTPRedirectsClient) Post307HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Put301 - Put true Boolean value in request returns 301.  This request should not be automatically redirected, but should return the received 301 to the caller for evaluation
func (client *HTTPRedirectsClient) Put301(ctx context.Context) (*HTTPRedirectsPut301Response, error) {
	req, err := client.Put301CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusMovedPermanently) {
		return nil, client.Put301HandleError(resp)
	}
	result, err := client.Put301HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Put301CreateRequest creates the Put301 request.
func (client *HTTPRedirectsClient) Put301CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/redirect/301"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// Put301HandleResponse handles the Put301 response.
func (client *HTTPRedirectsClient) Put301HandleResponse(resp *azcore.Response) (*HTTPRedirectsPut301Response, error) {
	result := HTTPRedirectsPut301Response{RawResponse: resp.Response}
	if val := resp.Header.Get("Location"); val != "" {
		result.Location = &val
	}
	return &result, nil
}

// Put301HandleError handles the Put301 error response.
func (client *HTTPRedirectsClient) Put301HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Put307 - Put redirected with 307, resulting in a 200 after redirect
func (client *HTTPRedirectsClient) Put307(ctx context.Context) (*http.Response, error) {
	req, err := client.Put307CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.Put307HandleError(resp)
	}
	return resp.Response, nil
}

// Put307CreateRequest creates the Put307 request.
func (client *HTTPRedirectsClient) Put307CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/redirect/307"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// Put307HandleError handles the Put307 error response.
func (client *HTTPRedirectsClient) Put307HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}