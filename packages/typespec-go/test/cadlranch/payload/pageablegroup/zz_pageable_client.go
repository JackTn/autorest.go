// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package pageablegroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strconv"
)

// PageableClient - Test describing pageable.
// Don't use this type directly, use a constructor function instead.
type PageableClient struct {
	internal *azcore.Client
}

// NewListPager - List users
//   - options - PageableClientListOptions contains the optional parameters for the PageableClient.NewListPager method.
func (client *PageableClient) NewListPager(options *PageableClientListOptions) *runtime.Pager[PageableClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[PageableClientListResponse]{
		More: func(page PageableClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PageableClientListResponse) (PageableClientListResponse, error) {
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return PageableClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *PageableClient) listCreateRequest(ctx context.Context, options *PageableClientListOptions) (*policy.Request, error) {
	urlPath := "/payload/pageable"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Maxpagesize != nil {
		reqQP.Set("maxpagesize", strconv.FormatInt(int64(*options.Maxpagesize), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PageableClient) listHandleResponse(resp *http.Response) (PageableClientListResponse, error) {
	result := PageableClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PagedUser); err != nil {
		return PageableClientListResponse{}, err
	}
	return result, nil
}
