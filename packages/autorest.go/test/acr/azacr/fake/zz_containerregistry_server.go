//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"azacr"
	"bytes"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// ContainerRegistryServer is a fake server for instances of the azacr.ContainerRegistryClient type.
type ContainerRegistryServer struct {
	// CheckDockerV2Support is the fake for method ContainerRegistryClient.CheckDockerV2Support
	// HTTP status codes to indicate success: http.StatusOK
	CheckDockerV2Support func(ctx context.Context, options *azacr.ContainerRegistryClientCheckDockerV2SupportOptions) (resp azfake.Responder[azacr.ContainerRegistryClientCheckDockerV2SupportResponse], errResp azfake.ErrorResponder)

	// CreateManifest is the fake for method ContainerRegistryClient.CreateManifest
	// HTTP status codes to indicate success: http.StatusCreated
	CreateManifest func(ctx context.Context, name string, reference string, payload azacr.Manifest, options *azacr.ContainerRegistryClientCreateManifestOptions) (resp azfake.Responder[azacr.ContainerRegistryClientCreateManifestResponse], errResp azfake.ErrorResponder)

	// DeleteManifest is the fake for method ContainerRegistryClient.DeleteManifest
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNotFound
	DeleteManifest func(ctx context.Context, name string, reference string, options *azacr.ContainerRegistryClientDeleteManifestOptions) (resp azfake.Responder[azacr.ContainerRegistryClientDeleteManifestResponse], errResp azfake.ErrorResponder)

	// DeleteRepository is the fake for method ContainerRegistryClient.DeleteRepository
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNotFound
	DeleteRepository func(ctx context.Context, name string, options *azacr.ContainerRegistryClientDeleteRepositoryOptions) (resp azfake.Responder[azacr.ContainerRegistryClientDeleteRepositoryResponse], errResp azfake.ErrorResponder)

	// DeleteTag is the fake for method ContainerRegistryClient.DeleteTag
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNotFound
	DeleteTag func(ctx context.Context, name string, reference string, options *azacr.ContainerRegistryClientDeleteTagOptions) (resp azfake.Responder[azacr.ContainerRegistryClientDeleteTagResponse], errResp azfake.ErrorResponder)

	// GetManifest is the fake for method ContainerRegistryClient.GetManifest
	// HTTP status codes to indicate success: http.StatusOK
	GetManifest func(ctx context.Context, name string, reference string, options *azacr.ContainerRegistryClientGetManifestOptions) (resp azfake.Responder[azacr.ContainerRegistryClientGetManifestResponse], errResp azfake.ErrorResponder)

	// GetManifestProperties is the fake for method ContainerRegistryClient.GetManifestProperties
	// HTTP status codes to indicate success: http.StatusOK
	GetManifestProperties func(ctx context.Context, name string, digest string, options *azacr.ContainerRegistryClientGetManifestPropertiesOptions) (resp azfake.Responder[azacr.ContainerRegistryClientGetManifestPropertiesResponse], errResp azfake.ErrorResponder)

	// NewGetManifestsPager is the fake for method ContainerRegistryClient.NewGetManifestsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewGetManifestsPager func(name string, options *azacr.ContainerRegistryClientGetManifestsOptions) (resp azfake.PagerResponder[azacr.ContainerRegistryClientGetManifestsResponse])

	// GetProperties is the fake for method ContainerRegistryClient.GetProperties
	// HTTP status codes to indicate success: http.StatusOK
	GetProperties func(ctx context.Context, name string, options *azacr.ContainerRegistryClientGetPropertiesOptions) (resp azfake.Responder[azacr.ContainerRegistryClientGetPropertiesResponse], errResp azfake.ErrorResponder)

	// NewGetRepositoriesPager is the fake for method ContainerRegistryClient.NewGetRepositoriesPager
	// HTTP status codes to indicate success: http.StatusOK
	NewGetRepositoriesPager func(options *azacr.ContainerRegistryClientGetRepositoriesOptions) (resp azfake.PagerResponder[azacr.ContainerRegistryClientGetRepositoriesResponse])

	// GetTagProperties is the fake for method ContainerRegistryClient.GetTagProperties
	// HTTP status codes to indicate success: http.StatusOK
	GetTagProperties func(ctx context.Context, name string, reference string, options *azacr.ContainerRegistryClientGetTagPropertiesOptions) (resp azfake.Responder[azacr.ContainerRegistryClientGetTagPropertiesResponse], errResp azfake.ErrorResponder)

	// NewGetTagsPager is the fake for method ContainerRegistryClient.NewGetTagsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewGetTagsPager func(name string, options *azacr.ContainerRegistryClientGetTagsOptions) (resp azfake.PagerResponder[azacr.ContainerRegistryClientGetTagsResponse])

	// UpdateManifestProperties is the fake for method ContainerRegistryClient.UpdateManifestProperties
	// HTTP status codes to indicate success: http.StatusOK
	UpdateManifestProperties func(ctx context.Context, name string, digest string, value azacr.ManifestWriteableProperties, options *azacr.ContainerRegistryClientUpdateManifestPropertiesOptions) (resp azfake.Responder[azacr.ContainerRegistryClientUpdateManifestPropertiesResponse], errResp azfake.ErrorResponder)

	// UpdateProperties is the fake for method ContainerRegistryClient.UpdateProperties
	// HTTP status codes to indicate success: http.StatusOK
	UpdateProperties func(ctx context.Context, name string, value azacr.RepositoryWriteableProperties, options *azacr.ContainerRegistryClientUpdatePropertiesOptions) (resp azfake.Responder[azacr.ContainerRegistryClientUpdatePropertiesResponse], errResp azfake.ErrorResponder)

	// UpdateTagAttributes is the fake for method ContainerRegistryClient.UpdateTagAttributes
	// HTTP status codes to indicate success: http.StatusOK
	UpdateTagAttributes func(ctx context.Context, name string, reference string, value azacr.TagWriteableProperties, options *azacr.ContainerRegistryClientUpdateTagAttributesOptions) (resp azfake.Responder[azacr.ContainerRegistryClientUpdateTagAttributesResponse], errResp azfake.ErrorResponder)
}

// NewContainerRegistryServerTransport creates a new instance of ContainerRegistryServerTransport with the provided implementation.
// The returned ContainerRegistryServerTransport instance is connected to an instance of azacr.ContainerRegistryClient by way of the
// undefined.Transporter field.
func NewContainerRegistryServerTransport(srv *ContainerRegistryServer) *ContainerRegistryServerTransport {
	return &ContainerRegistryServerTransport{srv: srv}
}

// ContainerRegistryServerTransport connects instances of azacr.ContainerRegistryClient to instances of ContainerRegistryServer.
// Don't use this type directly, use NewContainerRegistryServerTransport instead.
type ContainerRegistryServerTransport struct {
	srv                     *ContainerRegistryServer
	newGetManifestsPager    *azfake.PagerResponder[azacr.ContainerRegistryClientGetManifestsResponse]
	newGetRepositoriesPager *azfake.PagerResponder[azacr.ContainerRegistryClientGetRepositoriesResponse]
	newGetTagsPager         *azfake.PagerResponder[azacr.ContainerRegistryClientGetTagsResponse]
}

// Do implements the policy.Transporter interface for ContainerRegistryServerTransport.
func (c *ContainerRegistryServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ContainerRegistryClient.CheckDockerV2Support":
		resp, err = c.dispatchCheckDockerV2Support(req)
	case "ContainerRegistryClient.CreateManifest":
		resp, err = c.dispatchCreateManifest(req)
	case "ContainerRegistryClient.DeleteManifest":
		resp, err = c.dispatchDeleteManifest(req)
	case "ContainerRegistryClient.DeleteRepository":
		resp, err = c.dispatchDeleteRepository(req)
	case "ContainerRegistryClient.DeleteTag":
		resp, err = c.dispatchDeleteTag(req)
	case "ContainerRegistryClient.GetManifest":
		resp, err = c.dispatchGetManifest(req)
	case "ContainerRegistryClient.GetManifestProperties":
		resp, err = c.dispatchGetManifestProperties(req)
	case "ContainerRegistryClient.NewGetManifestsPager":
		resp, err = c.dispatchNewGetManifestsPager(req)
	case "ContainerRegistryClient.GetProperties":
		resp, err = c.dispatchGetProperties(req)
	case "ContainerRegistryClient.NewGetRepositoriesPager":
		resp, err = c.dispatchNewGetRepositoriesPager(req)
	case "ContainerRegistryClient.GetTagProperties":
		resp, err = c.dispatchGetTagProperties(req)
	case "ContainerRegistryClient.NewGetTagsPager":
		resp, err = c.dispatchNewGetTagsPager(req)
	case "ContainerRegistryClient.UpdateManifestProperties":
		resp, err = c.dispatchUpdateManifestProperties(req)
	case "ContainerRegistryClient.UpdateProperties":
		resp, err = c.dispatchUpdateProperties(req)
	case "ContainerRegistryClient.UpdateTagAttributes":
		resp, err = c.dispatchUpdateTagAttributes(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *ContainerRegistryServerTransport) dispatchCheckDockerV2Support(req *http.Request) (*http.Response, error) {
	if c.srv.CheckDockerV2Support == nil {
		return nil, &nonRetriableError{errors.New("fake for method CheckDockerV2Support not implemented")}
	}
	respr, errRespr := c.srv.CheckDockerV2Support(req.Context(), nil)
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

func (c *ContainerRegistryServerTransport) dispatchCreateManifest(req *http.Request) (*http.Response, error) {
	if c.srv.CreateManifest == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateManifest not implemented")}
	}
	const regexStr = `/v2/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/manifests/(?P<reference>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[azacr.Manifest](req)
	if err != nil {
		return nil, err
	}
	nameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	referenceUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("reference")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.CreateManifest(req.Context(), nameUnescaped, referenceUnescaped, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, &server.ResponseOptions{
		Body:        io.NopCloser(bytes.NewReader(server.GetResponse(respr).RawJSON)),
		ContentType: "application/json",
	})
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).DockerContentDigest; val != nil {
		resp.Header.Set("Docker-Content-Digest", *val)
	}
	if val := server.GetResponse(respr).Location; val != nil {
		resp.Header.Set("Location", *val)
	}
	if val := server.GetResponse(respr).ContentLength; val != nil {
		resp.Header.Set("Content-Length", strconv.FormatInt(*val, 10))
	}
	return resp, nil
}

func (c *ContainerRegistryServerTransport) dispatchDeleteManifest(req *http.Request) (*http.Response, error) {
	if c.srv.DeleteManifest == nil {
		return nil, &nonRetriableError{errors.New("fake for method DeleteManifest not implemented")}
	}
	const regexStr = `/v2/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/manifests/(?P<reference>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	nameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	referenceUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("reference")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.DeleteManifest(req.Context(), nameUnescaped, referenceUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusAccepted, http.StatusNotFound}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNotFound", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ContainerRegistryServerTransport) dispatchDeleteRepository(req *http.Request) (*http.Response, error) {
	if c.srv.DeleteRepository == nil {
		return nil, &nonRetriableError{errors.New("fake for method DeleteRepository not implemented")}
	}
	const regexStr = `/acr/v1/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	nameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.DeleteRepository(req.Context(), nameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusAccepted, http.StatusNotFound}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNotFound", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DeleteRepositoryResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ContainerRegistryServerTransport) dispatchDeleteTag(req *http.Request) (*http.Response, error) {
	if c.srv.DeleteTag == nil {
		return nil, &nonRetriableError{errors.New("fake for method DeleteTag not implemented")}
	}
	const regexStr = `/acr/v1/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/_tags/(?P<reference>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	nameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	referenceUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("reference")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.DeleteTag(req.Context(), nameUnescaped, referenceUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusAccepted, http.StatusNotFound}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNotFound", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ContainerRegistryServerTransport) dispatchGetManifest(req *http.Request) (*http.Response, error) {
	if c.srv.GetManifest == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetManifest not implemented")}
	}
	const regexStr = `/v2/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/manifests/(?P<reference>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	nameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	referenceUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("reference")])
	if err != nil {
		return nil, err
	}
	acceptParam := getOptional(getHeaderValue(req.Header, "accept"))
	var options *azacr.ContainerRegistryClientGetManifestOptions
	if acceptParam != nil {
		options = &azacr.ContainerRegistryClientGetManifestOptions{
			Accept: acceptParam,
		}
	}
	respr, errRespr := c.srv.GetManifest(req.Context(), nameUnescaped, referenceUnescaped, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ManifestWrapper, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ContainerRegistryServerTransport) dispatchGetManifestProperties(req *http.Request) (*http.Response, error) {
	if c.srv.GetManifestProperties == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetManifestProperties not implemented")}
	}
	const regexStr = `/acr/v1/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/_manifests/(?P<digest>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	nameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	digestUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("digest")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.GetManifestProperties(req.Context(), nameUnescaped, digestUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ArtifactManifestProperties, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ContainerRegistryServerTransport) dispatchNewGetManifestsPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewGetManifestsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewGetManifestsPager not implemented")}
	}
	if c.newGetManifestsPager == nil {
		const regexStr = `/acr/v1/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/_manifests`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		nameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
		if err != nil {
			return nil, err
		}
		lastUnescaped, err := url.QueryUnescape(qp.Get("last"))
		if err != nil {
			return nil, err
		}
		lastParam := getOptional(lastUnescaped)
		nUnescaped, err := url.QueryUnescape(qp.Get("n"))
		if err != nil {
			return nil, err
		}
		nParam, err := parseOptional(nUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		orderbyUnescaped, err := url.QueryUnescape(qp.Get("orderby"))
		if err != nil {
			return nil, err
		}
		orderbyParam := getOptional(orderbyUnescaped)
		var options *azacr.ContainerRegistryClientGetManifestsOptions
		if lastParam != nil || nParam != nil || orderbyParam != nil {
			options = &azacr.ContainerRegistryClientGetManifestsOptions{
				Last:    lastParam,
				N:       nParam,
				Orderby: orderbyParam,
			}
		}
		resp := c.srv.NewGetManifestsPager(nameUnescaped, options)
		c.newGetManifestsPager = &resp
		server.PagerResponderInjectNextLinks(c.newGetManifestsPager, req, func(page *azacr.ContainerRegistryClientGetManifestsResponse, createLink func() string) {
			page.Link = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(c.newGetManifestsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(c.newGetManifestsPager) {
		c.newGetManifestsPager = nil
	}
	return resp, nil
}

func (c *ContainerRegistryServerTransport) dispatchGetProperties(req *http.Request) (*http.Response, error) {
	if c.srv.GetProperties == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetProperties not implemented")}
	}
	const regexStr = `/acr/v1/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	nameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.GetProperties(req.Context(), nameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ContainerRepositoryProperties, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ContainerRegistryServerTransport) dispatchNewGetRepositoriesPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewGetRepositoriesPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewGetRepositoriesPager not implemented")}
	}
	if c.newGetRepositoriesPager == nil {
		qp := req.URL.Query()
		lastUnescaped, err := url.QueryUnescape(qp.Get("last"))
		if err != nil {
			return nil, err
		}
		lastParam := getOptional(lastUnescaped)
		nUnescaped, err := url.QueryUnescape(qp.Get("n"))
		if err != nil {
			return nil, err
		}
		nParam, err := parseOptional(nUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		var options *azacr.ContainerRegistryClientGetRepositoriesOptions
		if lastParam != nil || nParam != nil {
			options = &azacr.ContainerRegistryClientGetRepositoriesOptions{
				Last: lastParam,
				N:    nParam,
			}
		}
		resp := c.srv.NewGetRepositoriesPager(options)
		c.newGetRepositoriesPager = &resp
		server.PagerResponderInjectNextLinks(c.newGetRepositoriesPager, req, func(page *azacr.ContainerRegistryClientGetRepositoriesResponse, createLink func() string) {
			page.Link = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(c.newGetRepositoriesPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(c.newGetRepositoriesPager) {
		c.newGetRepositoriesPager = nil
	}
	return resp, nil
}

func (c *ContainerRegistryServerTransport) dispatchGetTagProperties(req *http.Request) (*http.Response, error) {
	if c.srv.GetTagProperties == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetTagProperties not implemented")}
	}
	const regexStr = `/acr/v1/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/_tags/(?P<reference>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	nameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	referenceUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("reference")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.GetTagProperties(req.Context(), nameUnescaped, referenceUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ArtifactTagProperties, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ContainerRegistryServerTransport) dispatchNewGetTagsPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewGetTagsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewGetTagsPager not implemented")}
	}
	if c.newGetTagsPager == nil {
		const regexStr = `/acr/v1/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/_tags`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		nameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
		if err != nil {
			return nil, err
		}
		lastUnescaped, err := url.QueryUnescape(qp.Get("last"))
		if err != nil {
			return nil, err
		}
		lastParam := getOptional(lastUnescaped)
		nUnescaped, err := url.QueryUnescape(qp.Get("n"))
		if err != nil {
			return nil, err
		}
		nParam, err := parseOptional(nUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		orderbyUnescaped, err := url.QueryUnescape(qp.Get("orderby"))
		if err != nil {
			return nil, err
		}
		orderbyParam := getOptional(orderbyUnescaped)
		digestUnescaped, err := url.QueryUnescape(qp.Get("digest"))
		if err != nil {
			return nil, err
		}
		digestParam := getOptional(digestUnescaped)
		var options *azacr.ContainerRegistryClientGetTagsOptions
		if lastParam != nil || nParam != nil || orderbyParam != nil || digestParam != nil {
			options = &azacr.ContainerRegistryClientGetTagsOptions{
				Last:    lastParam,
				N:       nParam,
				Orderby: orderbyParam,
				Digest:  digestParam,
			}
		}
		resp := c.srv.NewGetTagsPager(nameUnescaped, options)
		c.newGetTagsPager = &resp
		server.PagerResponderInjectNextLinks(c.newGetTagsPager, req, func(page *azacr.ContainerRegistryClientGetTagsResponse, createLink func() string) {
			page.Link = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(c.newGetTagsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(c.newGetTagsPager) {
		c.newGetTagsPager = nil
	}
	return resp, nil
}

func (c *ContainerRegistryServerTransport) dispatchUpdateManifestProperties(req *http.Request) (*http.Response, error) {
	if c.srv.UpdateManifestProperties == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateManifestProperties not implemented")}
	}
	const regexStr = `/acr/v1/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/_manifests/(?P<digest>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[azacr.ManifestWriteableProperties](req)
	if err != nil {
		return nil, err
	}
	nameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	digestUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("digest")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.UpdateManifestProperties(req.Context(), nameUnescaped, digestUnescaped, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ArtifactManifestProperties, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ContainerRegistryServerTransport) dispatchUpdateProperties(req *http.Request) (*http.Response, error) {
	if c.srv.UpdateProperties == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateProperties not implemented")}
	}
	const regexStr = `/acr/v1/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[azacr.RepositoryWriteableProperties](req)
	if err != nil {
		return nil, err
	}
	nameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.UpdateProperties(req.Context(), nameUnescaped, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ContainerRepositoryProperties, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ContainerRegistryServerTransport) dispatchUpdateTagAttributes(req *http.Request) (*http.Response, error) {
	if c.srv.UpdateTagAttributes == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateTagAttributes not implemented")}
	}
	const regexStr = `/acr/v1/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/_tags/(?P<reference>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[azacr.TagWriteableProperties](req)
	if err != nil {
		return nil, err
	}
	nameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	referenceUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("reference")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.UpdateTagAttributes(req.Context(), nameUnescaped, referenceUnescaped, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ArtifactTagProperties, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
