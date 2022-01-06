//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package integergroup

import (
	"net/http"
	"time"
)

// IntClientGetInvalidResponse contains the response from method IntClient.GetInvalid.
type IntClientGetInvalidResponse struct {
	IntClientGetInvalidResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// IntClientGetInvalidResult contains the result from method IntClient.GetInvalid.
type IntClientGetInvalidResult struct {
	Value *int32
}

// IntClientGetInvalidUnixTimeResponse contains the response from method IntClient.GetInvalidUnixTime.
type IntClientGetInvalidUnixTimeResponse struct {
	IntClientGetInvalidUnixTimeResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// IntClientGetInvalidUnixTimeResult contains the result from method IntClient.GetInvalidUnixTime.
type IntClientGetInvalidUnixTimeResult struct {
	// date in seconds since 1970-01-01T00:00:00Z.
	Value *time.Time
}

// IntClientGetNullResponse contains the response from method IntClient.GetNull.
type IntClientGetNullResponse struct {
	IntClientGetNullResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// IntClientGetNullResult contains the result from method IntClient.GetNull.
type IntClientGetNullResult struct {
	Value *int32
}

// IntClientGetNullUnixTimeResponse contains the response from method IntClient.GetNullUnixTime.
type IntClientGetNullUnixTimeResponse struct {
	IntClientGetNullUnixTimeResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// IntClientGetNullUnixTimeResult contains the result from method IntClient.GetNullUnixTime.
type IntClientGetNullUnixTimeResult struct {
	// date in seconds since 1970-01-01T00:00:00Z.
	Value *time.Time
}

// IntClientGetOverflowInt32Response contains the response from method IntClient.GetOverflowInt32.
type IntClientGetOverflowInt32Response struct {
	IntClientGetOverflowInt32Result
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// IntClientGetOverflowInt32Result contains the result from method IntClient.GetOverflowInt32.
type IntClientGetOverflowInt32Result struct {
	Value *int32
}

// IntClientGetOverflowInt64Response contains the response from method IntClient.GetOverflowInt64.
type IntClientGetOverflowInt64Response struct {
	IntClientGetOverflowInt64Result
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// IntClientGetOverflowInt64Result contains the result from method IntClient.GetOverflowInt64.
type IntClientGetOverflowInt64Result struct {
	Value *int64
}

// IntClientGetUnderflowInt32Response contains the response from method IntClient.GetUnderflowInt32.
type IntClientGetUnderflowInt32Response struct {
	IntClientGetUnderflowInt32Result
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// IntClientGetUnderflowInt32Result contains the result from method IntClient.GetUnderflowInt32.
type IntClientGetUnderflowInt32Result struct {
	Value *int32
}

// IntClientGetUnderflowInt64Response contains the response from method IntClient.GetUnderflowInt64.
type IntClientGetUnderflowInt64Response struct {
	IntClientGetUnderflowInt64Result
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// IntClientGetUnderflowInt64Result contains the result from method IntClient.GetUnderflowInt64.
type IntClientGetUnderflowInt64Result struct {
	Value *int64
}

// IntClientGetUnixTimeResponse contains the response from method IntClient.GetUnixTime.
type IntClientGetUnixTimeResponse struct {
	IntClientGetUnixTimeResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// IntClientGetUnixTimeResult contains the result from method IntClient.GetUnixTime.
type IntClientGetUnixTimeResult struct {
	// date in seconds since 1970-01-01T00:00:00Z.
	Value *time.Time
}

// IntClientPutMax32Response contains the response from method IntClient.PutMax32.
type IntClientPutMax32Response struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// IntClientPutMax64Response contains the response from method IntClient.PutMax64.
type IntClientPutMax64Response struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// IntClientPutMin32Response contains the response from method IntClient.PutMin32.
type IntClientPutMin32Response struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// IntClientPutMin64Response contains the response from method IntClient.PutMin64.
type IntClientPutMin64Response struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// IntClientPutUnixTimeDateResponse contains the response from method IntClient.PutUnixTimeDate.
type IntClientPutUnixTimeDateResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}
