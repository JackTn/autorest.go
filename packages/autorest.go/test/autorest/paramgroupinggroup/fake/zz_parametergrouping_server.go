// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	"generatortests/paramgroupinggroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// ParameterGroupingServer is a fake server for instances of the paramgroupinggroup.ParameterGroupingClient type.
type ParameterGroupingServer struct {
	// PostMultiParamGroups is the fake for method ParameterGroupingClient.PostMultiParamGroups
	// HTTP status codes to indicate success: http.StatusOK
	PostMultiParamGroups func(ctx context.Context, firstParameterGroup *paramgroupinggroup.FirstParameterGroup, parameterGroupingClientPostMultiParamGroupsSecondParamGroup *paramgroupinggroup.ParameterGroupingClientPostMultiParamGroupsSecondParamGroup, options *paramgroupinggroup.ParameterGroupingClientPostMultiParamGroupsOptions) (resp azfake.Responder[paramgroupinggroup.ParameterGroupingClientPostMultiParamGroupsResponse], errResp azfake.ErrorResponder)

	// PostOptional is the fake for method ParameterGroupingClient.PostOptional
	// HTTP status codes to indicate success: http.StatusOK
	PostOptional func(ctx context.Context, parameterGroupingClientPostOptionalParameters *paramgroupinggroup.ParameterGroupingClientPostOptionalParameters, options *paramgroupinggroup.ParameterGroupingClientPostOptionalOptions) (resp azfake.Responder[paramgroupinggroup.ParameterGroupingClientPostOptionalResponse], errResp azfake.ErrorResponder)

	// PostRequired is the fake for method ParameterGroupingClient.PostRequired
	// HTTP status codes to indicate success: http.StatusOK
	PostRequired func(ctx context.Context, parameterGroupingClientPostRequiredParameters paramgroupinggroup.ParameterGroupingClientPostRequiredParameters, options *paramgroupinggroup.ParameterGroupingClientPostRequiredOptions) (resp azfake.Responder[paramgroupinggroup.ParameterGroupingClientPostRequiredResponse], errResp azfake.ErrorResponder)

	// PostReservedWords is the fake for method ParameterGroupingClient.PostReservedWords
	// HTTP status codes to indicate success: http.StatusOK
	PostReservedWords func(ctx context.Context, parameterGroupingClientPostReservedWordsParameters *paramgroupinggroup.ParameterGroupingClientPostReservedWordsParameters, options *paramgroupinggroup.ParameterGroupingClientPostReservedWordsOptions) (resp azfake.Responder[paramgroupinggroup.ParameterGroupingClientPostReservedWordsResponse], errResp azfake.ErrorResponder)

	// PostSharedParameterGroupObject is the fake for method ParameterGroupingClient.PostSharedParameterGroupObject
	// HTTP status codes to indicate success: http.StatusOK
	PostSharedParameterGroupObject func(ctx context.Context, firstParameterGroup *paramgroupinggroup.FirstParameterGroup, options *paramgroupinggroup.ParameterGroupingClientPostSharedParameterGroupObjectOptions) (resp azfake.Responder[paramgroupinggroup.ParameterGroupingClientPostSharedParameterGroupObjectResponse], errResp azfake.ErrorResponder)
}

// NewParameterGroupingServerTransport creates a new instance of ParameterGroupingServerTransport with the provided implementation.
// The returned ParameterGroupingServerTransport instance is connected to an instance of paramgroupinggroup.ParameterGroupingClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewParameterGroupingServerTransport(srv *ParameterGroupingServer) *ParameterGroupingServerTransport {
	return &ParameterGroupingServerTransport{srv: srv}
}

// ParameterGroupingServerTransport connects instances of paramgroupinggroup.ParameterGroupingClient to instances of ParameterGroupingServer.
// Don't use this type directly, use NewParameterGroupingServerTransport instead.
type ParameterGroupingServerTransport struct {
	srv *ParameterGroupingServer
}

// Do implements the policy.Transporter interface for ParameterGroupingServerTransport.
func (p *ParameterGroupingServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ParameterGroupingClient.PostMultiParamGroups":
		resp, err = p.dispatchPostMultiParamGroups(req)
	case "ParameterGroupingClient.PostOptional":
		resp, err = p.dispatchPostOptional(req)
	case "ParameterGroupingClient.PostRequired":
		resp, err = p.dispatchPostRequired(req)
	case "ParameterGroupingClient.PostReservedWords":
		resp, err = p.dispatchPostReservedWords(req)
	case "ParameterGroupingClient.PostSharedParameterGroupObject":
		resp, err = p.dispatchPostSharedParameterGroupObject(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *ParameterGroupingServerTransport) dispatchPostMultiParamGroups(req *http.Request) (*http.Response, error) {
	if p.srv.PostMultiParamGroups == nil {
		return nil, &nonRetriableError{errors.New("fake for method PostMultiParamGroups not implemented")}
	}
	qp := req.URL.Query()
	headerOneParam := getOptional(getHeaderValue(req.Header, "header-one"))
	queryOneUnescaped, err := url.QueryUnescape(qp.Get("query-one"))
	if err != nil {
		return nil, err
	}
	queryOneParam, err := parseOptional(queryOneUnescaped, func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	headerTwoParam := getOptional(getHeaderValue(req.Header, "header-two"))
	queryTwoUnescaped, err := url.QueryUnescape(qp.Get("query-two"))
	if err != nil {
		return nil, err
	}
	queryTwoParam, err := parseOptional(queryTwoUnescaped, func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	var firstParameterGroup *paramgroupinggroup.FirstParameterGroup
	if headerOneParam != nil || queryOneParam != nil {
		firstParameterGroup = &paramgroupinggroup.FirstParameterGroup{
			HeaderOne: headerOneParam,
			QueryOne:  queryOneParam,
		}
	}
	var parameterGroupingClientPostMultiParamGroupsSecondParamGroup *paramgroupinggroup.ParameterGroupingClientPostMultiParamGroupsSecondParamGroup
	if headerTwoParam != nil || queryTwoParam != nil {
		parameterGroupingClientPostMultiParamGroupsSecondParamGroup = &paramgroupinggroup.ParameterGroupingClientPostMultiParamGroupsSecondParamGroup{
			HeaderTwo: headerTwoParam,
			QueryTwo:  queryTwoParam,
		}
	}
	respr, errRespr := p.srv.PostMultiParamGroups(req.Context(), firstParameterGroup, parameterGroupingClientPostMultiParamGroupsSecondParamGroup, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *ParameterGroupingServerTransport) dispatchPostOptional(req *http.Request) (*http.Response, error) {
	if p.srv.PostOptional == nil {
		return nil, &nonRetriableError{errors.New("fake for method PostOptional not implemented")}
	}
	qp := req.URL.Query()
	customHeaderParam := getOptional(getHeaderValue(req.Header, "customHeader"))
	queryUnescaped, err := url.QueryUnescape(qp.Get("query"))
	if err != nil {
		return nil, err
	}
	queryParam, err := parseOptional(queryUnescaped, func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	var parameterGroupingClientPostOptionalParameters *paramgroupinggroup.ParameterGroupingClientPostOptionalParameters
	if customHeaderParam != nil || queryParam != nil {
		parameterGroupingClientPostOptionalParameters = &paramgroupinggroup.ParameterGroupingClientPostOptionalParameters{
			CustomHeader: customHeaderParam,
			Query:        queryParam,
		}
	}
	respr, errRespr := p.srv.PostOptional(req.Context(), parameterGroupingClientPostOptionalParameters, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *ParameterGroupingServerTransport) dispatchPostRequired(req *http.Request) (*http.Response, error) {
	if p.srv.PostRequired == nil {
		return nil, &nonRetriableError{errors.New("fake for method PostRequired not implemented")}
	}
	const regexStr = `/parameterGrouping/postRequired/(?P<path>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	body, err := server.UnmarshalRequestAsJSON[int32](req)
	if err != nil {
		return nil, err
	}
	customHeaderParam := getOptional(getHeaderValue(req.Header, "customHeader"))
	queryUnescaped, err := url.QueryUnescape(qp.Get("query"))
	if err != nil {
		return nil, err
	}
	queryParam, err := parseOptional(queryUnescaped, func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	pathParam, err := url.PathUnescape(matches[regex.SubexpIndex("path")])
	if err != nil {
		return nil, err
	}
	parameterGroupingClientPostRequiredParameters := paramgroupinggroup.ParameterGroupingClientPostRequiredParameters{
		CustomHeader: customHeaderParam,
		Query:        queryParam,
		Path:         pathParam,
		Body:         body,
	}
	respr, errRespr := p.srv.PostRequired(req.Context(), parameterGroupingClientPostRequiredParameters, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *ParameterGroupingServerTransport) dispatchPostReservedWords(req *http.Request) (*http.Response, error) {
	if p.srv.PostReservedWords == nil {
		return nil, &nonRetriableError{errors.New("fake for method PostReservedWords not implemented")}
	}
	qp := req.URL.Query()
	fromUnescaped, err := url.QueryUnescape(qp.Get("from"))
	if err != nil {
		return nil, err
	}
	fromParam := getOptional(fromUnescaped)
	acceptUnescaped, err := url.QueryUnescape(qp.Get("accept"))
	if err != nil {
		return nil, err
	}
	acceptParam := getOptional(acceptUnescaped)
	var parameterGroupingClientPostReservedWordsParameters *paramgroupinggroup.ParameterGroupingClientPostReservedWordsParameters
	if fromParam != nil || acceptParam != nil {
		parameterGroupingClientPostReservedWordsParameters = &paramgroupinggroup.ParameterGroupingClientPostReservedWordsParameters{
			From:   fromParam,
			Accept: acceptParam,
		}
	}
	respr, errRespr := p.srv.PostReservedWords(req.Context(), parameterGroupingClientPostReservedWordsParameters, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *ParameterGroupingServerTransport) dispatchPostSharedParameterGroupObject(req *http.Request) (*http.Response, error) {
	if p.srv.PostSharedParameterGroupObject == nil {
		return nil, &nonRetriableError{errors.New("fake for method PostSharedParameterGroupObject not implemented")}
	}
	qp := req.URL.Query()
	headerOneParam := getOptional(getHeaderValue(req.Header, "header-one"))
	queryOneUnescaped, err := url.QueryUnescape(qp.Get("query-one"))
	if err != nil {
		return nil, err
	}
	queryOneParam, err := parseOptional(queryOneUnescaped, func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	var firstParameterGroup *paramgroupinggroup.FirstParameterGroup
	if headerOneParam != nil || queryOneParam != nil {
		firstParameterGroup = &paramgroupinggroup.FirstParameterGroup{
			HeaderOne: headerOneParam,
			QueryOne:  queryOneParam,
		}
	}
	respr, errRespr := p.srv.PostSharedParameterGroupObject(req.Context(), firstParameterGroup, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
