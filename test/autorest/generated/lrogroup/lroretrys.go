// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package lrogroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"strconv"
)

// LroRetrysOperations contains the methods for the LroRetrys group.
type LroRetrysOperations interface {
	// BeginDelete202Retry200 - Long running delete request, service returns a 500, then a 202 to the initial request. Polls return this value until the last poll returns a ‘200’ with ProvisioningState=’Succeeded’
	BeginDelete202Retry200(ctx context.Context) (LroRetrysDelete202Retry200Poller, error)
	// BeginDeleteAsyncRelativeRetrySucceeded - Long running delete request, service returns a 500, then a 202 to the initial request. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
	BeginDeleteAsyncRelativeRetrySucceeded(ctx context.Context) (LroRetrysDeleteAsyncRelativeRetrySucceededPoller, error)
	// BeginDeleteProvisioning202Accepted200Succeeded - Long running delete request, service returns a 500, then a  202 to the initial request, with an entity that contains ProvisioningState=’Accepted’.  Polls return this value until the last poll returns a ‘200’ with ProvisioningState=’Succeeded’
	BeginDeleteProvisioning202Accepted200Succeeded(ctx context.Context) (LroRetrysDeleteProvisioning202Accepted200SucceededPoller, error)
	// BeginPost202Retry200 - Long running post request, service returns a 500, then a 202 to the initial request, with 'Location' and 'Retry-After' headers, Polls return a 200 with a response body after success
	BeginPost202Retry200(ctx context.Context, options *LroRetrysPost202Retry200Options) (LroRetrysPost202Retry200Poller, error)
	// BeginPostAsyncRelativeRetrySucceeded - Long running post request, service returns a 500, then a 202 to the initial request, with an entity that contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
	BeginPostAsyncRelativeRetrySucceeded(ctx context.Context, options *LroRetrysPostAsyncRelativeRetrySucceededOptions) (LroRetrysPostAsyncRelativeRetrySucceededPoller, error)
	// BeginPut201CreatingSucceeded200 - Long running put request, service returns a 500, then a 201 to the initial request, with an entity that contains ProvisioningState=’Creating’.  Polls return this value until the last poll returns a ‘200’ with ProvisioningState=’Succeeded’
	BeginPut201CreatingSucceeded200(ctx context.Context, options *LroRetrysPut201CreatingSucceeded200Options) (LroRetrysPut201CreatingSucceeded200Poller, error)
	// BeginPutAsyncRelativeRetrySucceeded - Long running put request, service returns a 500, then a 200 to the initial request, with an entity that contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
	BeginPutAsyncRelativeRetrySucceeded(ctx context.Context, options *LroRetrysPutAsyncRelativeRetrySucceededOptions) (LroRetrysPutAsyncRelativeRetrySucceededPoller, error)
}

// lroRetrysOperations implements the LroRetrysOperations interface.
type lroRetrysOperations struct {
	*Client
}

// Delete202Retry200 - Long running delete request, service returns a 500, then a 202 to the initial request. Polls return this value until the last poll returns a ‘200’ with ProvisioningState=’Succeeded’
func (client *lroRetrysOperations) BeginDelete202Retry200(ctx context.Context) (LroRetrysDelete202Retry200Poller, error) {
	return &lroRetrysDelete202Retry200Poller{
		client: client,
	}, nil
}

// delete202Retry200CreateRequest creates the Delete202Retry200 request.
func (client *lroRetrysOperations) delete202Retry200CreateRequest() (*azcore.Request, error) {
	urlPath := "/lro/retryerror/delete/202/retry/200"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodDelete, *u)
	return req, nil
}

// delete202Retry200HandleResponse handles the Delete202Retry200 response.
func (client *lroRetrysOperations) delete202Retry200HandleResponse(resp *azcore.Response) (*LroRetrysDelete202Retry200Response, error) {
	if !resp.HasStatusCode(http.StatusAccepted) {
		return nil, newCloudError(resp)
	}
	result := LroRetrysDelete202Retry200Response{RawResponse: resp.Response}
	if val := resp.Header.Get("Location"); val != "" {
		result.Location = &val
	}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return nil, err
		}
		result.RetryAfter = &retryAfter
	}
	return &result, nil
}

// DeleteAsyncRelativeRetrySucceeded - Long running delete request, service returns a 500, then a 202 to the initial request. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
func (client *lroRetrysOperations) BeginDeleteAsyncRelativeRetrySucceeded(ctx context.Context) (LroRetrysDeleteAsyncRelativeRetrySucceededPoller, error) {
	return &lroRetrysDeleteAsyncRelativeRetrySucceededPoller{
		client: client,
	}, nil
}

// deleteAsyncRelativeRetrySucceededCreateRequest creates the DeleteAsyncRelativeRetrySucceeded request.
func (client *lroRetrysOperations) deleteAsyncRelativeRetrySucceededCreateRequest() (*azcore.Request, error) {
	urlPath := "/lro/retryerror/deleteasync/retry/succeeded"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodDelete, *u)
	return req, nil
}

// deleteAsyncRelativeRetrySucceededHandleResponse handles the DeleteAsyncRelativeRetrySucceeded response.
func (client *lroRetrysOperations) deleteAsyncRelativeRetrySucceededHandleResponse(resp *azcore.Response) (*LroRetrysDeleteAsyncRelativeRetrySucceededResponse, error) {
	if !resp.HasStatusCode(http.StatusAccepted) {
		return nil, newCloudError(resp)
	}
	result := LroRetrysDeleteAsyncRelativeRetrySucceededResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("Azure-AsyncOperation"); val != "" {
		result.AzureAsyncOperation = &val
	}
	if val := resp.Header.Get("Location"); val != "" {
		result.Location = &val
	}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return nil, err
		}
		result.RetryAfter = &retryAfter
	}
	return &result, nil
}

// DeleteProvisioning202Accepted200Succeeded - Long running delete request, service returns a 500, then a  202 to the initial request, with an entity that contains ProvisioningState=’Accepted’.  Polls return this value until the last poll returns a ‘200’ with ProvisioningState=’Succeeded’
func (client *lroRetrysOperations) BeginDeleteProvisioning202Accepted200Succeeded(ctx context.Context) (LroRetrysDeleteProvisioning202Accepted200SucceededPoller, error) {
	return &lroRetrysDeleteProvisioning202Accepted200SucceededPoller{
		client: client,
	}, nil
}

// deleteProvisioning202Accepted200SucceededCreateRequest creates the DeleteProvisioning202Accepted200Succeeded request.
func (client *lroRetrysOperations) deleteProvisioning202Accepted200SucceededCreateRequest() (*azcore.Request, error) {
	urlPath := "/lro/retryerror/delete/provisioning/202/accepted/200/succeeded"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodDelete, *u)
	return req, nil
}

// deleteProvisioning202Accepted200SucceededHandleResponse handles the DeleteProvisioning202Accepted200Succeeded response.
func (client *lroRetrysOperations) deleteProvisioning202Accepted200SucceededHandleResponse(resp *azcore.Response) (*ProductResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newCloudError(resp)
	}
	result := ProductResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Product)
}

// Post202Retry200 - Long running post request, service returns a 500, then a 202 to the initial request, with 'Location' and 'Retry-After' headers, Polls return a 200 with a response body after success
func (client *lroRetrysOperations) BeginPost202Retry200(ctx context.Context, options *LroRetrysPost202Retry200Options) (LroRetrysPost202Retry200Poller, error) {
	return &lroRetrysPost202Retry200Poller{
		client: client,
	}, nil
}

// post202Retry200CreateRequest creates the Post202Retry200 request.
func (client *lroRetrysOperations) post202Retry200CreateRequest(options *LroRetrysPost202Retry200Options) (*azcore.Request, error) {
	urlPath := "/lro/retryerror/post/202/retry/200"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPost, *u)
	if options != nil {
		return req, req.MarshalAsJSON(options.Product)
	}
	return req, nil
}

// post202Retry200HandleResponse handles the Post202Retry200 response.
func (client *lroRetrysOperations) post202Retry200HandleResponse(resp *azcore.Response) (*LroRetrysPost202Retry200Response, error) {
	if !resp.HasStatusCode(http.StatusAccepted) {
		return nil, newCloudError(resp)
	}
	result := LroRetrysPost202Retry200Response{RawResponse: resp.Response}
	if val := resp.Header.Get("Location"); val != "" {
		result.Location = &val
	}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return nil, err
		}
		result.RetryAfter = &retryAfter
	}
	return &result, nil
}

// PostAsyncRelativeRetrySucceeded - Long running post request, service returns a 500, then a 202 to the initial request, with an entity that contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
func (client *lroRetrysOperations) BeginPostAsyncRelativeRetrySucceeded(ctx context.Context, options *LroRetrysPostAsyncRelativeRetrySucceededOptions) (LroRetrysPostAsyncRelativeRetrySucceededPoller, error) {
	return &lroRetrysPostAsyncRelativeRetrySucceededPoller{
		client: client,
	}, nil
}

// postAsyncRelativeRetrySucceededCreateRequest creates the PostAsyncRelativeRetrySucceeded request.
func (client *lroRetrysOperations) postAsyncRelativeRetrySucceededCreateRequest(options *LroRetrysPostAsyncRelativeRetrySucceededOptions) (*azcore.Request, error) {
	urlPath := "/lro/retryerror/postasync/retry/succeeded"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPost, *u)
	if options != nil {
		return req, req.MarshalAsJSON(options.Product)
	}
	return req, nil
}

// postAsyncRelativeRetrySucceededHandleResponse handles the PostAsyncRelativeRetrySucceeded response.
func (client *lroRetrysOperations) postAsyncRelativeRetrySucceededHandleResponse(resp *azcore.Response) (*LroRetrysPostAsyncRelativeRetrySucceededResponse, error) {
	if !resp.HasStatusCode(http.StatusAccepted) {
		return nil, newCloudError(resp)
	}
	result := LroRetrysPostAsyncRelativeRetrySucceededResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("Azure-AsyncOperation"); val != "" {
		result.AzureAsyncOperation = &val
	}
	if val := resp.Header.Get("Location"); val != "" {
		result.Location = &val
	}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return nil, err
		}
		result.RetryAfter = &retryAfter
	}
	return &result, nil
}

// Put201CreatingSucceeded200 - Long running put request, service returns a 500, then a 201 to the initial request, with an entity that contains ProvisioningState=’Creating’.  Polls return this value until the last poll returns a ‘200’ with ProvisioningState=’Succeeded’
func (client *lroRetrysOperations) BeginPut201CreatingSucceeded200(ctx context.Context, options *LroRetrysPut201CreatingSucceeded200Options) (LroRetrysPut201CreatingSucceeded200Poller, error) {
	return &lroRetrysPut201CreatingSucceeded200Poller{
		client: client,
	}, nil
}

// put201CreatingSucceeded200CreateRequest creates the Put201CreatingSucceeded200 request.
func (client *lroRetrysOperations) put201CreatingSucceeded200CreateRequest(options *LroRetrysPut201CreatingSucceeded200Options) (*azcore.Request, error) {
	urlPath := "/lro/retryerror/put/201/creating/succeeded/200"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	if options != nil {
		return req, req.MarshalAsJSON(options.Product)
	}
	return req, nil
}

// put201CreatingSucceeded200HandleResponse handles the Put201CreatingSucceeded200 response.
func (client *lroRetrysOperations) put201CreatingSucceeded200HandleResponse(resp *azcore.Response) (*ProductResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newCloudError(resp)
	}
	result := ProductResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Product)
}

// PutAsyncRelativeRetrySucceeded - Long running put request, service returns a 500, then a 200 to the initial request, with an entity that contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
func (client *lroRetrysOperations) BeginPutAsyncRelativeRetrySucceeded(ctx context.Context, options *LroRetrysPutAsyncRelativeRetrySucceededOptions) (LroRetrysPutAsyncRelativeRetrySucceededPoller, error) {
	return &lroRetrysPutAsyncRelativeRetrySucceededPoller{
		client: client,
	}, nil
}

// putAsyncRelativeRetrySucceededCreateRequest creates the PutAsyncRelativeRetrySucceeded request.
func (client *lroRetrysOperations) putAsyncRelativeRetrySucceededCreateRequest(options *LroRetrysPutAsyncRelativeRetrySucceededOptions) (*azcore.Request, error) {
	urlPath := "/lro/retryerror/putasync/retry/succeeded"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	if options != nil {
		return req, req.MarshalAsJSON(options.Product)
	}
	return req, nil
}

// putAsyncRelativeRetrySucceededHandleResponse handles the PutAsyncRelativeRetrySucceeded response.
func (client *lroRetrysOperations) putAsyncRelativeRetrySucceededHandleResponse(resp *azcore.Response) (*ProductResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newCloudError(resp)
	}
	result := ProductResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Product)
}
