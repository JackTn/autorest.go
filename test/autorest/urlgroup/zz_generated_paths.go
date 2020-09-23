// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package urlgroup

import (
	"context"
	"encoding/base64"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// PathsOperations contains the methods for the Paths group.
type PathsOperations interface {
	// ArrayCSVInPath - Get an array of string ['ArrayPath1', 'begin!*'();:@ &=+$,/?#[]end' , null, ''] using the csv-array format
	ArrayCSVInPath(ctx context.Context, arrayPath []string) (*http.Response, error)
	// Base64URL - Get 'lorem' encoded value as 'bG9yZW0' (base64url)
	Base64URL(ctx context.Context, base64UrlPath []byte) (*http.Response, error)
	// ByteEmpty - Get '' as byte array
	ByteEmpty(ctx context.Context) (*http.Response, error)
	// ByteMultiByte - Get '啊齄丂狛狜隣郎隣兀﨩' multibyte value as utf-8 encoded byte array
	ByteMultiByte(ctx context.Context, bytePath []byte) (*http.Response, error)
	// ByteNull - Get null as byte array (should throw)
	ByteNull(ctx context.Context, bytePath []byte) (*http.Response, error)
	// DateNull - Get null as date - this should throw or be unusable on the client side, depending on date representation
	DateNull(ctx context.Context, datePath time.Time) (*http.Response, error)
	// DateTimeNull - Get null as date-time, should be disallowed or throw depending on representation of date-time
	DateTimeNull(ctx context.Context, dateTimePath time.Time) (*http.Response, error)
	// DateTimeValid - Get '2012-01-01T01:01:01Z' as date-time
	DateTimeValid(ctx context.Context) (*http.Response, error)
	// DateValid - Get '2012-01-01' as date
	DateValid(ctx context.Context) (*http.Response, error)
	// DoubleDecimalNegative - Get '-9999999.999' numeric value
	DoubleDecimalNegative(ctx context.Context) (*http.Response, error)
	// DoubleDecimalPositive - Get '9999999.999' numeric value
	DoubleDecimalPositive(ctx context.Context) (*http.Response, error)
	// EnumNull - Get null (should throw on the client before the request is sent on wire)
	EnumNull(ctx context.Context, enumPath URIColor) (*http.Response, error)
	// EnumValid - Get using uri with 'green color' in path parameter
	EnumValid(ctx context.Context, enumPath URIColor) (*http.Response, error)
	// FloatScientificNegative - Get '-1.034E-20' numeric value
	FloatScientificNegative(ctx context.Context) (*http.Response, error)
	// FloatScientificPositive - Get '1.034E+20' numeric value
	FloatScientificPositive(ctx context.Context) (*http.Response, error)
	// GetBooleanFalse - Get false Boolean value on path
	GetBooleanFalse(ctx context.Context) (*http.Response, error)
	// GetBooleanTrue - Get true Boolean value on path
	GetBooleanTrue(ctx context.Context) (*http.Response, error)
	// GetIntNegativeOneMillion - Get '-1000000' integer value
	GetIntNegativeOneMillion(ctx context.Context) (*http.Response, error)
	// GetIntOneMillion - Get '1000000' integer value
	GetIntOneMillion(ctx context.Context) (*http.Response, error)
	// GetNegativeTenBillion - Get '-10000000000' 64 bit integer value
	GetNegativeTenBillion(ctx context.Context) (*http.Response, error)
	// GetTenBillion - Get '10000000000' 64 bit integer value
	GetTenBillion(ctx context.Context) (*http.Response, error)
	// StringEmpty - Get ''
	StringEmpty(ctx context.Context) (*http.Response, error)
	// StringNull - Get null (should throw)
	StringNull(ctx context.Context, stringPath string) (*http.Response, error)
	// StringURLEncoded - Get 'begin!*'();:@ &=+$,/?#[]end
	StringURLEncoded(ctx context.Context) (*http.Response, error)
	// StringURLNonEncoded - https://tools.ietf.org/html/rfc3986#appendix-A 'path' accept any 'pchar' not encoded
	StringURLNonEncoded(ctx context.Context) (*http.Response, error)
	// StringUnicode - Get '啊齄丂狛狜隣郎隣兀﨩' multi-byte string value
	StringUnicode(ctx context.Context) (*http.Response, error)
	// UnixTimeURL - Get the date 2016-04-13 encoded value as '1460505600' (Unix time)
	UnixTimeURL(ctx context.Context, unixTimeUrlPath time.Time) (*http.Response, error)
}

// PathsClient implements the PathsOperations interface.
// Don't use this type directly, use NewPathsClient() instead.
type PathsClient struct {
	*Client
}

// NewPathsClient creates a new instance of PathsClient with the specified values.
func NewPathsClient(c *Client) PathsOperations {
	return &PathsClient{Client: c}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *PathsClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// ArrayCSVInPath - Get an array of string ['ArrayPath1', 'begin!*'();:@ &=+$,/?#[]end' , null, ''] using the csv-array format
func (client *PathsClient) ArrayCSVInPath(ctx context.Context, arrayPath []string) (*http.Response, error) {
	req, err := client.ArrayCSVInPathCreateRequest(ctx, arrayPath)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ArrayCSVInPathHandleError(resp)
	}
	return resp.Response, nil
}

// ArrayCSVInPathCreateRequest creates the ArrayCSVInPath request.
func (client *PathsClient) ArrayCSVInPathCreateRequest(ctx context.Context, arrayPath []string) (*azcore.Request, error) {
	urlPath := "/paths/array/ArrayPath1%2cbegin%21%2A%27%28%29%3B%3A%40%20%26%3D%2B%24%2C%2F%3F%23%5B%5Dend%2c%2c/{arrayPath}"
	urlPath = strings.ReplaceAll(urlPath, "{arrayPath}", url.PathEscape(strings.Join(arrayPath, ",")))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ArrayCSVInPathHandleError handles the ArrayCSVInPath error response.
func (client *PathsClient) ArrayCSVInPathHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Base64URL - Get 'lorem' encoded value as 'bG9yZW0' (base64url)
func (client *PathsClient) Base64URL(ctx context.Context, base64UrlPath []byte) (*http.Response, error) {
	req, err := client.Base64URLCreateRequest(ctx, base64UrlPath)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.Base64URLHandleError(resp)
	}
	return resp.Response, nil
}

// Base64URLCreateRequest creates the Base64URL request.
func (client *PathsClient) Base64URLCreateRequest(ctx context.Context, base64UrlPath []byte) (*azcore.Request, error) {
	urlPath := "/paths/string/bG9yZW0/{base64UrlPath}"
	urlPath = strings.ReplaceAll(urlPath, "{base64UrlPath}", url.PathEscape(base64.RawURLEncoding.EncodeToString(base64UrlPath)))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// Base64URLHandleError handles the Base64URL error response.
func (client *PathsClient) Base64URLHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ByteEmpty - Get '' as byte array
func (client *PathsClient) ByteEmpty(ctx context.Context) (*http.Response, error) {
	req, err := client.ByteEmptyCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ByteEmptyHandleError(resp)
	}
	return resp.Response, nil
}

// ByteEmptyCreateRequest creates the ByteEmpty request.
func (client *PathsClient) ByteEmptyCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/paths/byte/empty/{bytePath}"
	urlPath = strings.ReplaceAll(urlPath, "{bytePath}", url.PathEscape(""))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ByteEmptyHandleError handles the ByteEmpty error response.
func (client *PathsClient) ByteEmptyHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ByteMultiByte - Get '啊齄丂狛狜隣郎隣兀﨩' multibyte value as utf-8 encoded byte array
func (client *PathsClient) ByteMultiByte(ctx context.Context, bytePath []byte) (*http.Response, error) {
	req, err := client.ByteMultiByteCreateRequest(ctx, bytePath)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ByteMultiByteHandleError(resp)
	}
	return resp.Response, nil
}

// ByteMultiByteCreateRequest creates the ByteMultiByte request.
func (client *PathsClient) ByteMultiByteCreateRequest(ctx context.Context, bytePath []byte) (*azcore.Request, error) {
	urlPath := "/paths/byte/multibyte/{bytePath}"
	urlPath = strings.ReplaceAll(urlPath, "{bytePath}", url.PathEscape(base64.StdEncoding.EncodeToString(bytePath)))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ByteMultiByteHandleError handles the ByteMultiByte error response.
func (client *PathsClient) ByteMultiByteHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ByteNull - Get null as byte array (should throw)
func (client *PathsClient) ByteNull(ctx context.Context, bytePath []byte) (*http.Response, error) {
	req, err := client.ByteNullCreateRequest(ctx, bytePath)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusBadRequest) {
		return nil, client.ByteNullHandleError(resp)
	}
	return resp.Response, nil
}

// ByteNullCreateRequest creates the ByteNull request.
func (client *PathsClient) ByteNullCreateRequest(ctx context.Context, bytePath []byte) (*azcore.Request, error) {
	urlPath := "/paths/byte/null/{bytePath}"
	urlPath = strings.ReplaceAll(urlPath, "{bytePath}", url.PathEscape(base64.StdEncoding.EncodeToString(bytePath)))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ByteNullHandleError handles the ByteNull error response.
func (client *PathsClient) ByteNullHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// DateNull - Get null as date - this should throw or be unusable on the client side, depending on date representation
func (client *PathsClient) DateNull(ctx context.Context, datePath time.Time) (*http.Response, error) {
	req, err := client.DateNullCreateRequest(ctx, datePath)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusBadRequest) {
		return nil, client.DateNullHandleError(resp)
	}
	return resp.Response, nil
}

// DateNullCreateRequest creates the DateNull request.
func (client *PathsClient) DateNullCreateRequest(ctx context.Context, datePath time.Time) (*azcore.Request, error) {
	urlPath := "/paths/date/null/{datePath}"
	urlPath = strings.ReplaceAll(urlPath, "{datePath}", url.PathEscape(datePath.Format("2006-01-02")))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// DateNullHandleError handles the DateNull error response.
func (client *PathsClient) DateNullHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// DateTimeNull - Get null as date-time, should be disallowed or throw depending on representation of date-time
func (client *PathsClient) DateTimeNull(ctx context.Context, dateTimePath time.Time) (*http.Response, error) {
	req, err := client.DateTimeNullCreateRequest(ctx, dateTimePath)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusBadRequest) {
		return nil, client.DateTimeNullHandleError(resp)
	}
	return resp.Response, nil
}

// DateTimeNullCreateRequest creates the DateTimeNull request.
func (client *PathsClient) DateTimeNullCreateRequest(ctx context.Context, dateTimePath time.Time) (*azcore.Request, error) {
	urlPath := "/paths/datetime/null/{dateTimePath}"
	urlPath = strings.ReplaceAll(urlPath, "{dateTimePath}", url.PathEscape(dateTimePath.Format(time.RFC3339Nano)))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// DateTimeNullHandleError handles the DateTimeNull error response.
func (client *PathsClient) DateTimeNullHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// DateTimeValid - Get '2012-01-01T01:01:01Z' as date-time
func (client *PathsClient) DateTimeValid(ctx context.Context) (*http.Response, error) {
	req, err := client.DateTimeValidCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.DateTimeValidHandleError(resp)
	}
	return resp.Response, nil
}

// DateTimeValidCreateRequest creates the DateTimeValid request.
func (client *PathsClient) DateTimeValidCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/paths/datetime/2012-01-01T01%3A01%3A01Z/{dateTimePath}"
	urlPath = strings.ReplaceAll(urlPath, "{dateTimePath}", url.PathEscape("2012-01-01T01:01:01Z"))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// DateTimeValidHandleError handles the DateTimeValid error response.
func (client *PathsClient) DateTimeValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// DateValid - Get '2012-01-01' as date
func (client *PathsClient) DateValid(ctx context.Context) (*http.Response, error) {
	req, err := client.DateValidCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.DateValidHandleError(resp)
	}
	return resp.Response, nil
}

// DateValidCreateRequest creates the DateValid request.
func (client *PathsClient) DateValidCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/paths/date/2012-01-01/{datePath}"
	urlPath = strings.ReplaceAll(urlPath, "{datePath}", url.PathEscape("2012-01-01"))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// DateValidHandleError handles the DateValid error response.
func (client *PathsClient) DateValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// DoubleDecimalNegative - Get '-9999999.999' numeric value
func (client *PathsClient) DoubleDecimalNegative(ctx context.Context) (*http.Response, error) {
	req, err := client.DoubleDecimalNegativeCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.DoubleDecimalNegativeHandleError(resp)
	}
	return resp.Response, nil
}

// DoubleDecimalNegativeCreateRequest creates the DoubleDecimalNegative request.
func (client *PathsClient) DoubleDecimalNegativeCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/paths/double/-9999999.999/{doublePath}"
	urlPath = strings.ReplaceAll(urlPath, "{doublePath}", url.PathEscape("-9999999.999"))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// DoubleDecimalNegativeHandleError handles the DoubleDecimalNegative error response.
func (client *PathsClient) DoubleDecimalNegativeHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// DoubleDecimalPositive - Get '9999999.999' numeric value
func (client *PathsClient) DoubleDecimalPositive(ctx context.Context) (*http.Response, error) {
	req, err := client.DoubleDecimalPositiveCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.DoubleDecimalPositiveHandleError(resp)
	}
	return resp.Response, nil
}

// DoubleDecimalPositiveCreateRequest creates the DoubleDecimalPositive request.
func (client *PathsClient) DoubleDecimalPositiveCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/paths/double/9999999.999/{doublePath}"
	urlPath = strings.ReplaceAll(urlPath, "{doublePath}", url.PathEscape("9999999.999"))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// DoubleDecimalPositiveHandleError handles the DoubleDecimalPositive error response.
func (client *PathsClient) DoubleDecimalPositiveHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// EnumNull - Get null (should throw on the client before the request is sent on wire)
func (client *PathsClient) EnumNull(ctx context.Context, enumPath URIColor) (*http.Response, error) {
	req, err := client.EnumNullCreateRequest(ctx, enumPath)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusBadRequest) {
		return nil, client.EnumNullHandleError(resp)
	}
	return resp.Response, nil
}

// EnumNullCreateRequest creates the EnumNull request.
func (client *PathsClient) EnumNullCreateRequest(ctx context.Context, enumPath URIColor) (*azcore.Request, error) {
	urlPath := "/paths/string/null/{enumPath}"
	urlPath = strings.ReplaceAll(urlPath, "{enumPath}", url.PathEscape(string(enumPath)))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// EnumNullHandleError handles the EnumNull error response.
func (client *PathsClient) EnumNullHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// EnumValid - Get using uri with 'green color' in path parameter
func (client *PathsClient) EnumValid(ctx context.Context, enumPath URIColor) (*http.Response, error) {
	req, err := client.EnumValidCreateRequest(ctx, enumPath)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.EnumValidHandleError(resp)
	}
	return resp.Response, nil
}

// EnumValidCreateRequest creates the EnumValid request.
func (client *PathsClient) EnumValidCreateRequest(ctx context.Context, enumPath URIColor) (*azcore.Request, error) {
	urlPath := "/paths/enum/green%20color/{enumPath}"
	urlPath = strings.ReplaceAll(urlPath, "{enumPath}", url.PathEscape(string(enumPath)))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// EnumValidHandleError handles the EnumValid error response.
func (client *PathsClient) EnumValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// FloatScientificNegative - Get '-1.034E-20' numeric value
func (client *PathsClient) FloatScientificNegative(ctx context.Context) (*http.Response, error) {
	req, err := client.FloatScientificNegativeCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.FloatScientificNegativeHandleError(resp)
	}
	return resp.Response, nil
}

// FloatScientificNegativeCreateRequest creates the FloatScientificNegative request.
func (client *PathsClient) FloatScientificNegativeCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/paths/float/-1.034E-20/{floatPath}"
	urlPath = strings.ReplaceAll(urlPath, "{floatPath}", url.PathEscape("-1.034e-20"))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// FloatScientificNegativeHandleError handles the FloatScientificNegative error response.
func (client *PathsClient) FloatScientificNegativeHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// FloatScientificPositive - Get '1.034E+20' numeric value
func (client *PathsClient) FloatScientificPositive(ctx context.Context) (*http.Response, error) {
	req, err := client.FloatScientificPositiveCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.FloatScientificPositiveHandleError(resp)
	}
	return resp.Response, nil
}

// FloatScientificPositiveCreateRequest creates the FloatScientificPositive request.
func (client *PathsClient) FloatScientificPositiveCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/paths/float/1.034E+20/{floatPath}"
	urlPath = strings.ReplaceAll(urlPath, "{floatPath}", url.PathEscape("103400000000000000000"))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// FloatScientificPositiveHandleError handles the FloatScientificPositive error response.
func (client *PathsClient) FloatScientificPositiveHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetBooleanFalse - Get false Boolean value on path
func (client *PathsClient) GetBooleanFalse(ctx context.Context) (*http.Response, error) {
	req, err := client.GetBooleanFalseCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetBooleanFalseHandleError(resp)
	}
	return resp.Response, nil
}

// GetBooleanFalseCreateRequest creates the GetBooleanFalse request.
func (client *PathsClient) GetBooleanFalseCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/paths/bool/false/{boolPath}"
	urlPath = strings.ReplaceAll(urlPath, "{boolPath}", url.PathEscape("false"))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetBooleanFalseHandleError handles the GetBooleanFalse error response.
func (client *PathsClient) GetBooleanFalseHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetBooleanTrue - Get true Boolean value on path
func (client *PathsClient) GetBooleanTrue(ctx context.Context) (*http.Response, error) {
	req, err := client.GetBooleanTrueCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetBooleanTrueHandleError(resp)
	}
	return resp.Response, nil
}

// GetBooleanTrueCreateRequest creates the GetBooleanTrue request.
func (client *PathsClient) GetBooleanTrueCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/paths/bool/true/{boolPath}"
	urlPath = strings.ReplaceAll(urlPath, "{boolPath}", url.PathEscape("true"))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetBooleanTrueHandleError handles the GetBooleanTrue error response.
func (client *PathsClient) GetBooleanTrueHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetIntNegativeOneMillion - Get '-1000000' integer value
func (client *PathsClient) GetIntNegativeOneMillion(ctx context.Context) (*http.Response, error) {
	req, err := client.GetIntNegativeOneMillionCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetIntNegativeOneMillionHandleError(resp)
	}
	return resp.Response, nil
}

// GetIntNegativeOneMillionCreateRequest creates the GetIntNegativeOneMillion request.
func (client *PathsClient) GetIntNegativeOneMillionCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/paths/int/-1000000/{intPath}"
	urlPath = strings.ReplaceAll(urlPath, "{intPath}", url.PathEscape("-1000000"))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetIntNegativeOneMillionHandleError handles the GetIntNegativeOneMillion error response.
func (client *PathsClient) GetIntNegativeOneMillionHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetIntOneMillion - Get '1000000' integer value
func (client *PathsClient) GetIntOneMillion(ctx context.Context) (*http.Response, error) {
	req, err := client.GetIntOneMillionCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetIntOneMillionHandleError(resp)
	}
	return resp.Response, nil
}

// GetIntOneMillionCreateRequest creates the GetIntOneMillion request.
func (client *PathsClient) GetIntOneMillionCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/paths/int/1000000/{intPath}"
	urlPath = strings.ReplaceAll(urlPath, "{intPath}", url.PathEscape("1000000"))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetIntOneMillionHandleError handles the GetIntOneMillion error response.
func (client *PathsClient) GetIntOneMillionHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetNegativeTenBillion - Get '-10000000000' 64 bit integer value
func (client *PathsClient) GetNegativeTenBillion(ctx context.Context) (*http.Response, error) {
	req, err := client.GetNegativeTenBillionCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetNegativeTenBillionHandleError(resp)
	}
	return resp.Response, nil
}

// GetNegativeTenBillionCreateRequest creates the GetNegativeTenBillion request.
func (client *PathsClient) GetNegativeTenBillionCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/paths/long/-10000000000/{longPath}"
	urlPath = strings.ReplaceAll(urlPath, "{longPath}", url.PathEscape("-10000000000"))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetNegativeTenBillionHandleError handles the GetNegativeTenBillion error response.
func (client *PathsClient) GetNegativeTenBillionHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetTenBillion - Get '10000000000' 64 bit integer value
func (client *PathsClient) GetTenBillion(ctx context.Context) (*http.Response, error) {
	req, err := client.GetTenBillionCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetTenBillionHandleError(resp)
	}
	return resp.Response, nil
}

// GetTenBillionCreateRequest creates the GetTenBillion request.
func (client *PathsClient) GetTenBillionCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/paths/long/10000000000/{longPath}"
	urlPath = strings.ReplaceAll(urlPath, "{longPath}", url.PathEscape("10000000000"))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetTenBillionHandleError handles the GetTenBillion error response.
func (client *PathsClient) GetTenBillionHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// StringEmpty - Get ''
func (client *PathsClient) StringEmpty(ctx context.Context) (*http.Response, error) {
	req, err := client.StringEmptyCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.StringEmptyHandleError(resp)
	}
	return resp.Response, nil
}

// StringEmptyCreateRequest creates the StringEmpty request.
func (client *PathsClient) StringEmptyCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/paths/string/empty/{stringPath}"
	urlPath = strings.ReplaceAll(urlPath, "{stringPath}", url.PathEscape(""))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// StringEmptyHandleError handles the StringEmpty error response.
func (client *PathsClient) StringEmptyHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// StringNull - Get null (should throw)
func (client *PathsClient) StringNull(ctx context.Context, stringPath string) (*http.Response, error) {
	req, err := client.StringNullCreateRequest(ctx, stringPath)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusBadRequest) {
		return nil, client.StringNullHandleError(resp)
	}
	return resp.Response, nil
}

// StringNullCreateRequest creates the StringNull request.
func (client *PathsClient) StringNullCreateRequest(ctx context.Context, stringPath string) (*azcore.Request, error) {
	urlPath := "/paths/string/null/{stringPath}"
	urlPath = strings.ReplaceAll(urlPath, "{stringPath}", url.PathEscape(stringPath))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// StringNullHandleError handles the StringNull error response.
func (client *PathsClient) StringNullHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// StringURLEncoded - Get 'begin!*'();:@ &=+$,/?#[]end
func (client *PathsClient) StringURLEncoded(ctx context.Context) (*http.Response, error) {
	req, err := client.StringURLEncodedCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.StringURLEncodedHandleError(resp)
	}
	return resp.Response, nil
}

// StringURLEncodedCreateRequest creates the StringURLEncoded request.
func (client *PathsClient) StringURLEncodedCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/paths/string/begin%21%2A%27%28%29%3B%3A%40%20%26%3D%2B%24%2C%2F%3F%23%5B%5Dend/{stringPath}"
	urlPath = strings.ReplaceAll(urlPath, "{stringPath}", url.PathEscape("begin!*'();:@ &=+$,/?#[]end"))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// StringURLEncodedHandleError handles the StringURLEncoded error response.
func (client *PathsClient) StringURLEncodedHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// StringURLNonEncoded - https://tools.ietf.org/html/rfc3986#appendix-A 'path' accept any 'pchar' not encoded
func (client *PathsClient) StringURLNonEncoded(ctx context.Context) (*http.Response, error) {
	req, err := client.StringURLNonEncodedCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.StringURLNonEncodedHandleError(resp)
	}
	return resp.Response, nil
}

// StringURLNonEncodedCreateRequest creates the StringURLNonEncoded request.
func (client *PathsClient) StringURLNonEncodedCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/paths/string/begin!*'();:@&=+$,end/{stringPath}"
	urlPath = strings.ReplaceAll(urlPath, "{stringPath}", "begin!*'();:@&=+$,end")
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// StringURLNonEncodedHandleError handles the StringURLNonEncoded error response.
func (client *PathsClient) StringURLNonEncodedHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// StringUnicode - Get '啊齄丂狛狜隣郎隣兀﨩' multi-byte string value
func (client *PathsClient) StringUnicode(ctx context.Context) (*http.Response, error) {
	req, err := client.StringUnicodeCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.StringUnicodeHandleError(resp)
	}
	return resp.Response, nil
}

// StringUnicodeCreateRequest creates the StringUnicode request.
func (client *PathsClient) StringUnicodeCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/paths/string/unicode/{stringPath}"
	urlPath = strings.ReplaceAll(urlPath, "{stringPath}", url.PathEscape("啊齄丂狛狜隣郎隣兀﨩"))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// StringUnicodeHandleError handles the StringUnicode error response.
func (client *PathsClient) StringUnicodeHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// UnixTimeURL - Get the date 2016-04-13 encoded value as '1460505600' (Unix time)
func (client *PathsClient) UnixTimeURL(ctx context.Context, unixTimeUrlPath time.Time) (*http.Response, error) {
	req, err := client.UnixTimeURLCreateRequest(ctx, unixTimeUrlPath)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.UnixTimeURLHandleError(resp)
	}
	return resp.Response, nil
}

// UnixTimeURLCreateRequest creates the UnixTimeURL request.
func (client *PathsClient) UnixTimeURLCreateRequest(ctx context.Context, unixTimeUrlPath time.Time) (*azcore.Request, error) {
	urlPath := "/paths/int/1460505600/{unixTimeUrlPath}"
	urlPath = strings.ReplaceAll(urlPath, "{unixTimeUrlPath}", url.PathEscape(timeUnix(unixTimeUrlPath).String()))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// UnixTimeURLHandleError handles the UnixTimeURL error response.
func (client *PathsClient) UnixTimeURLHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}