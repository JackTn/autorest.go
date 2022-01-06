//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package headergroup

import (
	"net/http"
	"time"
)

// HeaderClientCustomRequestIDResponse contains the response from method HeaderClient.CustomRequestID.
type HeaderClientCustomRequestIDResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderClientParamBoolResponse contains the response from method HeaderClient.ParamBool.
type HeaderClientParamBoolResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderClientParamByteResponse contains the response from method HeaderClient.ParamByte.
type HeaderClientParamByteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderClientParamDateResponse contains the response from method HeaderClient.ParamDate.
type HeaderClientParamDateResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderClientParamDatetimeRFC1123Response contains the response from method HeaderClient.ParamDatetimeRFC1123.
type HeaderClientParamDatetimeRFC1123Response struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderClientParamDatetimeResponse contains the response from method HeaderClient.ParamDatetime.
type HeaderClientParamDatetimeResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderClientParamDoubleResponse contains the response from method HeaderClient.ParamDouble.
type HeaderClientParamDoubleResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderClientParamDurationResponse contains the response from method HeaderClient.ParamDuration.
type HeaderClientParamDurationResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderClientParamEnumResponse contains the response from method HeaderClient.ParamEnum.
type HeaderClientParamEnumResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderClientParamExistingKeyResponse contains the response from method HeaderClient.ParamExistingKey.
type HeaderClientParamExistingKeyResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderClientParamFloatResponse contains the response from method HeaderClient.ParamFloat.
type HeaderClientParamFloatResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderClientParamIntegerResponse contains the response from method HeaderClient.ParamInteger.
type HeaderClientParamIntegerResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderClientParamLongResponse contains the response from method HeaderClient.ParamLong.
type HeaderClientParamLongResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderClientParamProtectedKeyResponse contains the response from method HeaderClient.ParamProtectedKey.
type HeaderClientParamProtectedKeyResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderClientParamStringResponse contains the response from method HeaderClient.ParamString.
type HeaderClientParamStringResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderClientResponseBoolResponse contains the response from method HeaderClient.ResponseBool.
type HeaderClientResponseBoolResponse struct {
	HeaderClientResponseBoolResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderClientResponseBoolResult contains the result from method HeaderClient.ResponseBool.
type HeaderClientResponseBoolResult struct {
	// Value contains the information returned from the value header response.
	Value *bool
}

// HeaderClientResponseByteResponse contains the response from method HeaderClient.ResponseByte.
type HeaderClientResponseByteResponse struct {
	HeaderClientResponseByteResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderClientResponseByteResult contains the result from method HeaderClient.ResponseByte.
type HeaderClientResponseByteResult struct {
	// Value contains the information returned from the value header response.
	Value []byte
}

// HeaderClientResponseDateResponse contains the response from method HeaderClient.ResponseDate.
type HeaderClientResponseDateResponse struct {
	HeaderClientResponseDateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderClientResponseDateResult contains the result from method HeaderClient.ResponseDate.
type HeaderClientResponseDateResult struct {
	// Value contains the information returned from the value header response.
	Value *time.Time
}

// HeaderClientResponseDatetimeRFC1123Response contains the response from method HeaderClient.ResponseDatetimeRFC1123.
type HeaderClientResponseDatetimeRFC1123Response struct {
	HeaderClientResponseDatetimeRFC1123Result
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderClientResponseDatetimeRFC1123Result contains the result from method HeaderClient.ResponseDatetimeRFC1123.
type HeaderClientResponseDatetimeRFC1123Result struct {
	// Value contains the information returned from the value header response.
	Value *time.Time
}

// HeaderClientResponseDatetimeResponse contains the response from method HeaderClient.ResponseDatetime.
type HeaderClientResponseDatetimeResponse struct {
	HeaderClientResponseDatetimeResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderClientResponseDatetimeResult contains the result from method HeaderClient.ResponseDatetime.
type HeaderClientResponseDatetimeResult struct {
	// Value contains the information returned from the value header response.
	Value *time.Time
}

// HeaderClientResponseDoubleResponse contains the response from method HeaderClient.ResponseDouble.
type HeaderClientResponseDoubleResponse struct {
	HeaderClientResponseDoubleResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderClientResponseDoubleResult contains the result from method HeaderClient.ResponseDouble.
type HeaderClientResponseDoubleResult struct {
	// Value contains the information returned from the value header response.
	Value *float64
}

// HeaderClientResponseDurationResponse contains the response from method HeaderClient.ResponseDuration.
type HeaderClientResponseDurationResponse struct {
	HeaderClientResponseDurationResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderClientResponseDurationResult contains the result from method HeaderClient.ResponseDuration.
type HeaderClientResponseDurationResult struct {
	// Value contains the information returned from the value header response.
	Value *string
}

// HeaderClientResponseEnumResponse contains the response from method HeaderClient.ResponseEnum.
type HeaderClientResponseEnumResponse struct {
	HeaderClientResponseEnumResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderClientResponseEnumResult contains the result from method HeaderClient.ResponseEnum.
type HeaderClientResponseEnumResult struct {
	// Value contains the information returned from the value header response.
	Value *GreyscaleColors
}

// HeaderClientResponseExistingKeyResponse contains the response from method HeaderClient.ResponseExistingKey.
type HeaderClientResponseExistingKeyResponse struct {
	HeaderClientResponseExistingKeyResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderClientResponseExistingKeyResult contains the result from method HeaderClient.ResponseExistingKey.
type HeaderClientResponseExistingKeyResult struct {
	// UserAgent contains the information returned from the User-Agent header response.
	UserAgent *string
}

// HeaderClientResponseFloatResponse contains the response from method HeaderClient.ResponseFloat.
type HeaderClientResponseFloatResponse struct {
	HeaderClientResponseFloatResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderClientResponseFloatResult contains the result from method HeaderClient.ResponseFloat.
type HeaderClientResponseFloatResult struct {
	// Value contains the information returned from the value header response.
	Value *float32
}

// HeaderClientResponseIntegerResponse contains the response from method HeaderClient.ResponseInteger.
type HeaderClientResponseIntegerResponse struct {
	HeaderClientResponseIntegerResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderClientResponseIntegerResult contains the result from method HeaderClient.ResponseInteger.
type HeaderClientResponseIntegerResult struct {
	// Value contains the information returned from the value header response.
	Value *int32
}

// HeaderClientResponseLongResponse contains the response from method HeaderClient.ResponseLong.
type HeaderClientResponseLongResponse struct {
	HeaderClientResponseLongResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderClientResponseLongResult contains the result from method HeaderClient.ResponseLong.
type HeaderClientResponseLongResult struct {
	// Value contains the information returned from the value header response.
	Value *int64
}

// HeaderClientResponseProtectedKeyResponse contains the response from method HeaderClient.ResponseProtectedKey.
type HeaderClientResponseProtectedKeyResponse struct {
	HeaderClientResponseProtectedKeyResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderClientResponseProtectedKeyResult contains the result from method HeaderClient.ResponseProtectedKey.
type HeaderClientResponseProtectedKeyResult struct {
	// ContentType contains the information returned from the Content-Type header response.
	ContentType *string
}

// HeaderClientResponseStringResponse contains the response from method HeaderClient.ResponseString.
type HeaderClientResponseStringResponse struct {
	HeaderClientResponseStringResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderClientResponseStringResult contains the result from method HeaderClient.ResponseString.
type HeaderClientResponseStringResult struct {
	// Value contains the information returned from the value header response.
	Value *string
}
