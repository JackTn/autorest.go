// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"generatortests/urlgroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"time"
)

// PathsServer is a fake server for instances of the urlgroup.PathsClient type.
type PathsServer struct {
	// ArrayCSVInPath is the fake for method PathsClient.ArrayCSVInPath
	// HTTP status codes to indicate success: http.StatusOK
	ArrayCSVInPath func(ctx context.Context, arrayPath []string, options *urlgroup.PathsClientArrayCSVInPathOptions) (resp azfake.Responder[urlgroup.PathsClientArrayCSVInPathResponse], errResp azfake.ErrorResponder)

	// Base64URL is the fake for method PathsClient.Base64URL
	// HTTP status codes to indicate success: http.StatusOK
	Base64URL func(ctx context.Context, base64URLPath []byte, options *urlgroup.PathsClientBase64URLOptions) (resp azfake.Responder[urlgroup.PathsClientBase64URLResponse], errResp azfake.ErrorResponder)

	// ByteEmpty is the fake for method PathsClient.ByteEmpty
	// HTTP status codes to indicate success: http.StatusOK
	ByteEmpty func(ctx context.Context, options *urlgroup.PathsClientByteEmptyOptions) (resp azfake.Responder[urlgroup.PathsClientByteEmptyResponse], errResp azfake.ErrorResponder)

	// ByteMultiByte is the fake for method PathsClient.ByteMultiByte
	// HTTP status codes to indicate success: http.StatusOK
	ByteMultiByte func(ctx context.Context, bytePath []byte, options *urlgroup.PathsClientByteMultiByteOptions) (resp azfake.Responder[urlgroup.PathsClientByteMultiByteResponse], errResp azfake.ErrorResponder)

	// ByteNull is the fake for method PathsClient.ByteNull
	// HTTP status codes to indicate success: http.StatusBadRequest
	ByteNull func(ctx context.Context, bytePath []byte, options *urlgroup.PathsClientByteNullOptions) (resp azfake.Responder[urlgroup.PathsClientByteNullResponse], errResp azfake.ErrorResponder)

	// DateNull is the fake for method PathsClient.DateNull
	// HTTP status codes to indicate success: http.StatusBadRequest
	DateNull func(ctx context.Context, datePath time.Time, options *urlgroup.PathsClientDateNullOptions) (resp azfake.Responder[urlgroup.PathsClientDateNullResponse], errResp azfake.ErrorResponder)

	// DateTimeNull is the fake for method PathsClient.DateTimeNull
	// HTTP status codes to indicate success: http.StatusBadRequest
	DateTimeNull func(ctx context.Context, dateTimePath time.Time, options *urlgroup.PathsClientDateTimeNullOptions) (resp azfake.Responder[urlgroup.PathsClientDateTimeNullResponse], errResp azfake.ErrorResponder)

	// DateTimeValid is the fake for method PathsClient.DateTimeValid
	// HTTP status codes to indicate success: http.StatusOK
	DateTimeValid func(ctx context.Context, options *urlgroup.PathsClientDateTimeValidOptions) (resp azfake.Responder[urlgroup.PathsClientDateTimeValidResponse], errResp azfake.ErrorResponder)

	// DateValid is the fake for method PathsClient.DateValid
	// HTTP status codes to indicate success: http.StatusOK
	DateValid func(ctx context.Context, options *urlgroup.PathsClientDateValidOptions) (resp azfake.Responder[urlgroup.PathsClientDateValidResponse], errResp azfake.ErrorResponder)

	// DoubleDecimalNegative is the fake for method PathsClient.DoubleDecimalNegative
	// HTTP status codes to indicate success: http.StatusOK
	DoubleDecimalNegative func(ctx context.Context, options *urlgroup.PathsClientDoubleDecimalNegativeOptions) (resp azfake.Responder[urlgroup.PathsClientDoubleDecimalNegativeResponse], errResp azfake.ErrorResponder)

	// DoubleDecimalPositive is the fake for method PathsClient.DoubleDecimalPositive
	// HTTP status codes to indicate success: http.StatusOK
	DoubleDecimalPositive func(ctx context.Context, options *urlgroup.PathsClientDoubleDecimalPositiveOptions) (resp azfake.Responder[urlgroup.PathsClientDoubleDecimalPositiveResponse], errResp azfake.ErrorResponder)

	// EnumNull is the fake for method PathsClient.EnumNull
	// HTTP status codes to indicate success: http.StatusBadRequest
	EnumNull func(ctx context.Context, enumPath urlgroup.URIColor, options *urlgroup.PathsClientEnumNullOptions) (resp azfake.Responder[urlgroup.PathsClientEnumNullResponse], errResp azfake.ErrorResponder)

	// EnumValid is the fake for method PathsClient.EnumValid
	// HTTP status codes to indicate success: http.StatusOK
	EnumValid func(ctx context.Context, enumPath urlgroup.URIColor, options *urlgroup.PathsClientEnumValidOptions) (resp azfake.Responder[urlgroup.PathsClientEnumValidResponse], errResp azfake.ErrorResponder)

	// FloatScientificNegative is the fake for method PathsClient.FloatScientificNegative
	// HTTP status codes to indicate success: http.StatusOK
	FloatScientificNegative func(ctx context.Context, options *urlgroup.PathsClientFloatScientificNegativeOptions) (resp azfake.Responder[urlgroup.PathsClientFloatScientificNegativeResponse], errResp azfake.ErrorResponder)

	// FloatScientificPositive is the fake for method PathsClient.FloatScientificPositive
	// HTTP status codes to indicate success: http.StatusOK
	FloatScientificPositive func(ctx context.Context, options *urlgroup.PathsClientFloatScientificPositiveOptions) (resp azfake.Responder[urlgroup.PathsClientFloatScientificPositiveResponse], errResp azfake.ErrorResponder)

	// GetBooleanFalse is the fake for method PathsClient.GetBooleanFalse
	// HTTP status codes to indicate success: http.StatusOK
	GetBooleanFalse func(ctx context.Context, options *urlgroup.PathsClientGetBooleanFalseOptions) (resp azfake.Responder[urlgroup.PathsClientGetBooleanFalseResponse], errResp azfake.ErrorResponder)

	// GetBooleanTrue is the fake for method PathsClient.GetBooleanTrue
	// HTTP status codes to indicate success: http.StatusOK
	GetBooleanTrue func(ctx context.Context, options *urlgroup.PathsClientGetBooleanTrueOptions) (resp azfake.Responder[urlgroup.PathsClientGetBooleanTrueResponse], errResp azfake.ErrorResponder)

	// GetIntNegativeOneMillion is the fake for method PathsClient.GetIntNegativeOneMillion
	// HTTP status codes to indicate success: http.StatusOK
	GetIntNegativeOneMillion func(ctx context.Context, options *urlgroup.PathsClientGetIntNegativeOneMillionOptions) (resp azfake.Responder[urlgroup.PathsClientGetIntNegativeOneMillionResponse], errResp azfake.ErrorResponder)

	// GetIntOneMillion is the fake for method PathsClient.GetIntOneMillion
	// HTTP status codes to indicate success: http.StatusOK
	GetIntOneMillion func(ctx context.Context, options *urlgroup.PathsClientGetIntOneMillionOptions) (resp azfake.Responder[urlgroup.PathsClientGetIntOneMillionResponse], errResp azfake.ErrorResponder)

	// GetNegativeTenBillion is the fake for method PathsClient.GetNegativeTenBillion
	// HTTP status codes to indicate success: http.StatusOK
	GetNegativeTenBillion func(ctx context.Context, options *urlgroup.PathsClientGetNegativeTenBillionOptions) (resp azfake.Responder[urlgroup.PathsClientGetNegativeTenBillionResponse], errResp azfake.ErrorResponder)

	// GetTenBillion is the fake for method PathsClient.GetTenBillion
	// HTTP status codes to indicate success: http.StatusOK
	GetTenBillion func(ctx context.Context, options *urlgroup.PathsClientGetTenBillionOptions) (resp azfake.Responder[urlgroup.PathsClientGetTenBillionResponse], errResp azfake.ErrorResponder)

	// StringEmpty is the fake for method PathsClient.StringEmpty
	// HTTP status codes to indicate success: http.StatusOK
	StringEmpty func(ctx context.Context, options *urlgroup.PathsClientStringEmptyOptions) (resp azfake.Responder[urlgroup.PathsClientStringEmptyResponse], errResp azfake.ErrorResponder)

	// StringNull is the fake for method PathsClient.StringNull
	// HTTP status codes to indicate success: http.StatusBadRequest
	StringNull func(ctx context.Context, stringPath string, options *urlgroup.PathsClientStringNullOptions) (resp azfake.Responder[urlgroup.PathsClientStringNullResponse], errResp azfake.ErrorResponder)

	// StringURLEncoded is the fake for method PathsClient.StringURLEncoded
	// HTTP status codes to indicate success: http.StatusOK
	StringURLEncoded func(ctx context.Context, options *urlgroup.PathsClientStringURLEncodedOptions) (resp azfake.Responder[urlgroup.PathsClientStringURLEncodedResponse], errResp azfake.ErrorResponder)

	// StringURLNonEncoded is the fake for method PathsClient.StringURLNonEncoded
	// HTTP status codes to indicate success: http.StatusOK
	StringURLNonEncoded func(ctx context.Context, options *urlgroup.PathsClientStringURLNonEncodedOptions) (resp azfake.Responder[urlgroup.PathsClientStringURLNonEncodedResponse], errResp azfake.ErrorResponder)

	// StringUnicode is the fake for method PathsClient.StringUnicode
	// HTTP status codes to indicate success: http.StatusOK
	StringUnicode func(ctx context.Context, options *urlgroup.PathsClientStringUnicodeOptions) (resp azfake.Responder[urlgroup.PathsClientStringUnicodeResponse], errResp azfake.ErrorResponder)

	// UnixTimeURL is the fake for method PathsClient.UnixTimeURL
	// HTTP status codes to indicate success: http.StatusOK
	UnixTimeURL func(ctx context.Context, unixTimeURLPath time.Time, options *urlgroup.PathsClientUnixTimeURLOptions) (resp azfake.Responder[urlgroup.PathsClientUnixTimeURLResponse], errResp azfake.ErrorResponder)
}

// NewPathsServerTransport creates a new instance of PathsServerTransport with the provided implementation.
// The returned PathsServerTransport instance is connected to an instance of urlgroup.PathsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPathsServerTransport(srv *PathsServer) *PathsServerTransport {
	return &PathsServerTransport{srv: srv}
}

// PathsServerTransport connects instances of urlgroup.PathsClient to instances of PathsServer.
// Don't use this type directly, use NewPathsServerTransport instead.
type PathsServerTransport struct {
	srv *PathsServer
}

// Do implements the policy.Transporter interface for PathsServerTransport.
func (p *PathsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PathsClient.ArrayCSVInPath":
		resp, err = p.dispatchArrayCSVInPath(req)
	case "PathsClient.Base64URL":
		resp, err = p.dispatchBase64URL(req)
	case "PathsClient.ByteEmpty":
		resp, err = p.dispatchByteEmpty(req)
	case "PathsClient.ByteMultiByte":
		resp, err = p.dispatchByteMultiByte(req)
	case "PathsClient.ByteNull":
		resp, err = p.dispatchByteNull(req)
	case "PathsClient.DateNull":
		resp, err = p.dispatchDateNull(req)
	case "PathsClient.DateTimeNull":
		resp, err = p.dispatchDateTimeNull(req)
	case "PathsClient.DateTimeValid":
		resp, err = p.dispatchDateTimeValid(req)
	case "PathsClient.DateValid":
		resp, err = p.dispatchDateValid(req)
	case "PathsClient.DoubleDecimalNegative":
		resp, err = p.dispatchDoubleDecimalNegative(req)
	case "PathsClient.DoubleDecimalPositive":
		resp, err = p.dispatchDoubleDecimalPositive(req)
	case "PathsClient.EnumNull":
		resp, err = p.dispatchEnumNull(req)
	case "PathsClient.EnumValid":
		resp, err = p.dispatchEnumValid(req)
	case "PathsClient.FloatScientificNegative":
		resp, err = p.dispatchFloatScientificNegative(req)
	case "PathsClient.FloatScientificPositive":
		resp, err = p.dispatchFloatScientificPositive(req)
	case "PathsClient.GetBooleanFalse":
		resp, err = p.dispatchGetBooleanFalse(req)
	case "PathsClient.GetBooleanTrue":
		resp, err = p.dispatchGetBooleanTrue(req)
	case "PathsClient.GetIntNegativeOneMillion":
		resp, err = p.dispatchGetIntNegativeOneMillion(req)
	case "PathsClient.GetIntOneMillion":
		resp, err = p.dispatchGetIntOneMillion(req)
	case "PathsClient.GetNegativeTenBillion":
		resp, err = p.dispatchGetNegativeTenBillion(req)
	case "PathsClient.GetTenBillion":
		resp, err = p.dispatchGetTenBillion(req)
	case "PathsClient.StringEmpty":
		resp, err = p.dispatchStringEmpty(req)
	case "PathsClient.StringNull":
		resp, err = p.dispatchStringNull(req)
	case "PathsClient.StringURLEncoded":
		resp, err = p.dispatchStringURLEncoded(req)
	case "PathsClient.StringURLNonEncoded":
		resp, err = p.dispatchStringURLNonEncoded(req)
	case "PathsClient.StringUnicode":
		resp, err = p.dispatchStringUnicode(req)
	case "PathsClient.UnixTimeURL":
		resp, err = p.dispatchUnixTimeURL(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PathsServerTransport) dispatchArrayCSVInPath(req *http.Request) (*http.Response, error) {
	if p.srv.ArrayCSVInPath == nil {
		return nil, &nonRetriableError{errors.New("fake for method ArrayCSVInPath not implemented")}
	}
	const regexStr = `/paths/array/ArrayPath1%2cbegin%21%2A%27%28%29%3B%3A%40%20%26%3D%2B%24%2C%2F%3F%23%5B%5Dend%2c%2c/(?P<arrayPath>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	arrayPathParam, err := url.PathUnescape(matches[regex.SubexpIndex("arrayPath")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.ArrayCSVInPath(req.Context(), splitHelper(arrayPathParam, ","), nil)
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

func (p *PathsServerTransport) dispatchBase64URL(req *http.Request) (*http.Response, error) {
	if p.srv.Base64URL == nil {
		return nil, &nonRetriableError{errors.New("fake for method Base64URL not implemented")}
	}
	const regexStr = `/paths/string/bG9yZW0/(?P<base64UrlPath>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	base64URLPathUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("base64UrlPath")])
	if err != nil {
		return nil, err
	}
	base64URLPathParam, err := base64.StdEncoding.DecodeString(base64URLPathUnescaped)
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Base64URL(req.Context(), base64URLPathParam, nil)
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

func (p *PathsServerTransport) dispatchByteEmpty(req *http.Request) (*http.Response, error) {
	if p.srv.ByteEmpty == nil {
		return nil, &nonRetriableError{errors.New("fake for method ByteEmpty not implemented")}
	}
	respr, errRespr := p.srv.ByteEmpty(req.Context(), nil)
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

func (p *PathsServerTransport) dispatchByteMultiByte(req *http.Request) (*http.Response, error) {
	if p.srv.ByteMultiByte == nil {
		return nil, &nonRetriableError{errors.New("fake for method ByteMultiByte not implemented")}
	}
	const regexStr = `/paths/byte/multibyte/(?P<bytePath>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	bytePathUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("bytePath")])
	if err != nil {
		return nil, err
	}
	bytePathParam, err := base64.StdEncoding.DecodeString(bytePathUnescaped)
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.ByteMultiByte(req.Context(), bytePathParam, nil)
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

func (p *PathsServerTransport) dispatchByteNull(req *http.Request) (*http.Response, error) {
	if p.srv.ByteNull == nil {
		return nil, &nonRetriableError{errors.New("fake for method ByteNull not implemented")}
	}
	const regexStr = `/paths/byte/null/(?P<bytePath>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	bytePathUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("bytePath")])
	if err != nil {
		return nil, err
	}
	bytePathParam, err := base64.StdEncoding.DecodeString(bytePathUnescaped)
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.ByteNull(req.Context(), bytePathParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusBadRequest}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusBadRequest", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PathsServerTransport) dispatchDateNull(req *http.Request) (*http.Response, error) {
	if p.srv.DateNull == nil {
		return nil, &nonRetriableError{errors.New("fake for method DateNull not implemented")}
	}
	const regexStr = `/paths/date/null/(?P<datePath>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	datePathUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("datePath")])
	if err != nil {
		return nil, err
	}
	datePathParam, err := time.Parse("2006-01-02", datePathUnescaped)
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.DateNull(req.Context(), datePathParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusBadRequest}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusBadRequest", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PathsServerTransport) dispatchDateTimeNull(req *http.Request) (*http.Response, error) {
	if p.srv.DateTimeNull == nil {
		return nil, &nonRetriableError{errors.New("fake for method DateTimeNull not implemented")}
	}
	const regexStr = `/paths/datetime/null/(?P<dateTimePath>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	dateTimePathUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("dateTimePath")])
	if err != nil {
		return nil, err
	}
	dateTimePathParam, err := time.Parse(time.RFC3339Nano, dateTimePathUnescaped)
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.DateTimeNull(req.Context(), dateTimePathParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusBadRequest}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusBadRequest", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PathsServerTransport) dispatchDateTimeValid(req *http.Request) (*http.Response, error) {
	if p.srv.DateTimeValid == nil {
		return nil, &nonRetriableError{errors.New("fake for method DateTimeValid not implemented")}
	}
	respr, errRespr := p.srv.DateTimeValid(req.Context(), nil)
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

func (p *PathsServerTransport) dispatchDateValid(req *http.Request) (*http.Response, error) {
	if p.srv.DateValid == nil {
		return nil, &nonRetriableError{errors.New("fake for method DateValid not implemented")}
	}
	respr, errRespr := p.srv.DateValid(req.Context(), nil)
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

func (p *PathsServerTransport) dispatchDoubleDecimalNegative(req *http.Request) (*http.Response, error) {
	if p.srv.DoubleDecimalNegative == nil {
		return nil, &nonRetriableError{errors.New("fake for method DoubleDecimalNegative not implemented")}
	}
	respr, errRespr := p.srv.DoubleDecimalNegative(req.Context(), nil)
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

func (p *PathsServerTransport) dispatchDoubleDecimalPositive(req *http.Request) (*http.Response, error) {
	if p.srv.DoubleDecimalPositive == nil {
		return nil, &nonRetriableError{errors.New("fake for method DoubleDecimalPositive not implemented")}
	}
	respr, errRespr := p.srv.DoubleDecimalPositive(req.Context(), nil)
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

func (p *PathsServerTransport) dispatchEnumNull(req *http.Request) (*http.Response, error) {
	if p.srv.EnumNull == nil {
		return nil, &nonRetriableError{errors.New("fake for method EnumNull not implemented")}
	}
	const regexStr = `/paths/string/null/(?P<enumPath>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	enumPathParam, err := parseWithCast(matches[regex.SubexpIndex("enumPath")], func(v string) (urlgroup.URIColor, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return urlgroup.URIColor(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.EnumNull(req.Context(), enumPathParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusBadRequest}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusBadRequest", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PathsServerTransport) dispatchEnumValid(req *http.Request) (*http.Response, error) {
	if p.srv.EnumValid == nil {
		return nil, &nonRetriableError{errors.New("fake for method EnumValid not implemented")}
	}
	const regexStr = `/paths/enum/green%20color/(?P<enumPath>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	enumPathParam, err := parseWithCast(matches[regex.SubexpIndex("enumPath")], func(v string) (urlgroup.URIColor, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return urlgroup.URIColor(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.EnumValid(req.Context(), enumPathParam, nil)
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

func (p *PathsServerTransport) dispatchFloatScientificNegative(req *http.Request) (*http.Response, error) {
	if p.srv.FloatScientificNegative == nil {
		return nil, &nonRetriableError{errors.New("fake for method FloatScientificNegative not implemented")}
	}
	respr, errRespr := p.srv.FloatScientificNegative(req.Context(), nil)
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

func (p *PathsServerTransport) dispatchFloatScientificPositive(req *http.Request) (*http.Response, error) {
	if p.srv.FloatScientificPositive == nil {
		return nil, &nonRetriableError{errors.New("fake for method FloatScientificPositive not implemented")}
	}
	respr, errRespr := p.srv.FloatScientificPositive(req.Context(), nil)
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

func (p *PathsServerTransport) dispatchGetBooleanFalse(req *http.Request) (*http.Response, error) {
	if p.srv.GetBooleanFalse == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetBooleanFalse not implemented")}
	}
	respr, errRespr := p.srv.GetBooleanFalse(req.Context(), nil)
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

func (p *PathsServerTransport) dispatchGetBooleanTrue(req *http.Request) (*http.Response, error) {
	if p.srv.GetBooleanTrue == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetBooleanTrue not implemented")}
	}
	respr, errRespr := p.srv.GetBooleanTrue(req.Context(), nil)
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

func (p *PathsServerTransport) dispatchGetIntNegativeOneMillion(req *http.Request) (*http.Response, error) {
	if p.srv.GetIntNegativeOneMillion == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetIntNegativeOneMillion not implemented")}
	}
	respr, errRespr := p.srv.GetIntNegativeOneMillion(req.Context(), nil)
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

func (p *PathsServerTransport) dispatchGetIntOneMillion(req *http.Request) (*http.Response, error) {
	if p.srv.GetIntOneMillion == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetIntOneMillion not implemented")}
	}
	respr, errRespr := p.srv.GetIntOneMillion(req.Context(), nil)
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

func (p *PathsServerTransport) dispatchGetNegativeTenBillion(req *http.Request) (*http.Response, error) {
	if p.srv.GetNegativeTenBillion == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetNegativeTenBillion not implemented")}
	}
	respr, errRespr := p.srv.GetNegativeTenBillion(req.Context(), nil)
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

func (p *PathsServerTransport) dispatchGetTenBillion(req *http.Request) (*http.Response, error) {
	if p.srv.GetTenBillion == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetTenBillion not implemented")}
	}
	respr, errRespr := p.srv.GetTenBillion(req.Context(), nil)
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

func (p *PathsServerTransport) dispatchStringEmpty(req *http.Request) (*http.Response, error) {
	if p.srv.StringEmpty == nil {
		return nil, &nonRetriableError{errors.New("fake for method StringEmpty not implemented")}
	}
	respr, errRespr := p.srv.StringEmpty(req.Context(), nil)
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

func (p *PathsServerTransport) dispatchStringNull(req *http.Request) (*http.Response, error) {
	if p.srv.StringNull == nil {
		return nil, &nonRetriableError{errors.New("fake for method StringNull not implemented")}
	}
	const regexStr = `/paths/string/null/(?P<stringPath>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	stringPathParam, err := url.PathUnescape(matches[regex.SubexpIndex("stringPath")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.StringNull(req.Context(), stringPathParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusBadRequest}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusBadRequest", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PathsServerTransport) dispatchStringURLEncoded(req *http.Request) (*http.Response, error) {
	if p.srv.StringURLEncoded == nil {
		return nil, &nonRetriableError{errors.New("fake for method StringURLEncoded not implemented")}
	}
	respr, errRespr := p.srv.StringURLEncoded(req.Context(), nil)
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

func (p *PathsServerTransport) dispatchStringURLNonEncoded(req *http.Request) (*http.Response, error) {
	if p.srv.StringURLNonEncoded == nil {
		return nil, &nonRetriableError{errors.New("fake for method StringURLNonEncoded not implemented")}
	}
	respr, errRespr := p.srv.StringURLNonEncoded(req.Context(), nil)
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

func (p *PathsServerTransport) dispatchStringUnicode(req *http.Request) (*http.Response, error) {
	if p.srv.StringUnicode == nil {
		return nil, &nonRetriableError{errors.New("fake for method StringUnicode not implemented")}
	}
	respr, errRespr := p.srv.StringUnicode(req.Context(), nil)
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

func (p *PathsServerTransport) dispatchUnixTimeURL(req *http.Request) (*http.Response, error) {
	if p.srv.UnixTimeURL == nil {
		return nil, &nonRetriableError{errors.New("fake for method UnixTimeURL not implemented")}
	}
	const regexStr = `/paths/int/1460505600/(?P<unixTimeUrlPath>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	unixTimeURLPathUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("unixTimeUrlPath")])
	if err != nil {
		return nil, err
	}
	unixTimeURLPathParam, err := parseWithCast(unixTimeURLPathUnescaped, func(v string) (time.Time, error) {
		p, parseErr := strconv.ParseInt(v, 10, 64)
		if parseErr != nil {
			return time.Time{}, parseErr
		}
		return time.Unix(p, 0).UTC(), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.UnixTimeURL(req.Context(), unixTimeURLPathParam, nil)
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
