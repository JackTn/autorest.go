// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azalias

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// AliasListResponsePager provides iteration over AliasListResponse pages.
type AliasListResponsePager interface {
	azcore.Pager

	// PageResponse returns the current AliasListResponseResponse.
	PageResponse() AliasListResponseResponse
}

type aliasListResponseCreateRequest func(context.Context) (*azcore.Request, error)

type aliasListResponseHandleError func(*azcore.Response) error

type aliasListResponseHandleResponse func(*azcore.Response) (AliasListResponseResponse, error)

type aliasListResponseAdvancePage func(context.Context, AliasListResponseResponse) (*azcore.Request, error)

type aliasListResponsePager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// creates the initial request (non-LRO case)
	requester aliasListResponseCreateRequest
	// callback for handling response errors
	errorer aliasListResponseHandleError
	// callback for handling the HTTP response
	responder aliasListResponseHandleResponse
	// callback for advancing to the next page
	advancer aliasListResponseAdvancePage
	// contains the current response
	current AliasListResponseResponse
	// status codes for successful retrieval
	statusCodes []int
	// any error encountered
	err error
}

func (p *aliasListResponsePager) Err() error {
	return p.err
}

func (p *aliasListResponsePager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AliasListResponse.NextLink == nil || len(*p.current.AliasListResponse.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.pipeline.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(p.statusCodes...) {
		p.err = p.errorer(resp)
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *aliasListResponsePager) PageResponse() AliasListResponseResponse {
	return p.current
}
