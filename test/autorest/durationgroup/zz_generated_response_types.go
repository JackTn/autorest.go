//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package durationgroup

import "net/http"

// DurationClientGetInvalidResponse contains the response from method DurationClient.GetInvalid.
type DurationClientGetInvalidResponse struct {
	DurationClientGetInvalidResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DurationClientGetInvalidResult contains the result from method DurationClient.GetInvalid.
type DurationClientGetInvalidResult struct {
	Value *string
}

// DurationClientGetNullResponse contains the response from method DurationClient.GetNull.
type DurationClientGetNullResponse struct {
	DurationClientGetNullResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DurationClientGetNullResult contains the result from method DurationClient.GetNull.
type DurationClientGetNullResult struct {
	Value *string
}

// DurationClientGetPositiveDurationResponse contains the response from method DurationClient.GetPositiveDuration.
type DurationClientGetPositiveDurationResponse struct {
	DurationClientGetPositiveDurationResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DurationClientGetPositiveDurationResult contains the result from method DurationClient.GetPositiveDuration.
type DurationClientGetPositiveDurationResult struct {
	Value *string
}

// DurationClientPutPositiveDurationResponse contains the response from method DurationClient.PutPositiveDuration.
type DurationClientPutPositiveDurationResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}
