// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package stringgroup

import "net/http"

// EnumGetNotExpandableResponse contains the response from method Enum.GetNotExpandable.
type EnumGetNotExpandableResponse struct {
	EnumGetNotExpandableResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EnumGetNotExpandableResult contains the result from method Enum.GetNotExpandable.
type EnumGetNotExpandableResult struct {
	// Referenced Color Enum Description.
	Value *Colors
}

// EnumGetReferencedConstantResponse contains the response from method Enum.GetReferencedConstant.
type EnumGetReferencedConstantResponse struct {
	EnumGetReferencedConstantResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EnumGetReferencedConstantResult contains the result from method Enum.GetReferencedConstant.
type EnumGetReferencedConstantResult struct {
	RefColorConstant
}

// EnumGetReferencedResponse contains the response from method Enum.GetReferenced.
type EnumGetReferencedResponse struct {
	EnumGetReferencedResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EnumGetReferencedResult contains the result from method Enum.GetReferenced.
type EnumGetReferencedResult struct {
	// Referenced Color Enum Description.
	Value *Colors
}

// EnumPutNotExpandableResponse contains the response from method Enum.PutNotExpandable.
type EnumPutNotExpandableResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EnumPutReferencedConstantResponse contains the response from method Enum.PutReferencedConstant.
type EnumPutReferencedConstantResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EnumPutReferencedResponse contains the response from method Enum.PutReferenced.
type EnumPutReferencedResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// StringGetBase64EncodedResponse contains the response from method String.GetBase64Encoded.
type StringGetBase64EncodedResponse struct {
	StringGetBase64EncodedResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// StringGetBase64EncodedResult contains the result from method String.GetBase64Encoded.
type StringGetBase64EncodedResult struct {
	Value []byte
}

// StringGetBase64URLEncodedResponse contains the response from method String.GetBase64URLEncoded.
type StringGetBase64URLEncodedResponse struct {
	StringGetBase64URLEncodedResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// StringGetBase64URLEncodedResult contains the result from method String.GetBase64URLEncoded.
type StringGetBase64URLEncodedResult struct {
	Value []byte
}

// StringGetEmptyResponse contains the response from method String.GetEmpty.
type StringGetEmptyResponse struct {
	StringGetEmptyResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// StringGetEmptyResult contains the result from method String.GetEmpty.
type StringGetEmptyResult struct {
	// simple string
	Value *string
}

// StringGetMBCSResponse contains the response from method String.GetMBCS.
type StringGetMBCSResponse struct {
	StringGetMBCSResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// StringGetMBCSResult contains the result from method String.GetMBCS.
type StringGetMBCSResult struct {
	// simple string
	Value *string
}

// StringGetNotProvidedResponse contains the response from method String.GetNotProvided.
type StringGetNotProvidedResponse struct {
	StringGetNotProvidedResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// StringGetNotProvidedResult contains the result from method String.GetNotProvided.
type StringGetNotProvidedResult struct {
	Value *string
}

// StringGetNullBase64URLEncodedResponse contains the response from method String.GetNullBase64URLEncoded.
type StringGetNullBase64URLEncodedResponse struct {
	StringGetNullBase64URLEncodedResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// StringGetNullBase64URLEncodedResult contains the result from method String.GetNullBase64URLEncoded.
type StringGetNullBase64URLEncodedResult struct {
	Value []byte
}

// StringGetNullResponse contains the response from method String.GetNull.
type StringGetNullResponse struct {
	StringGetNullResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// StringGetNullResult contains the result from method String.GetNull.
type StringGetNullResult struct {
	Value *string
}

// StringGetWhitespaceResponse contains the response from method String.GetWhitespace.
type StringGetWhitespaceResponse struct {
	StringGetWhitespaceResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// StringGetWhitespaceResult contains the result from method String.GetWhitespace.
type StringGetWhitespaceResult struct {
	// simple string
	Value *string
}

// StringPutBase64URLEncodedResponse contains the response from method String.PutBase64URLEncoded.
type StringPutBase64URLEncodedResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// StringPutEmptyResponse contains the response from method String.PutEmpty.
type StringPutEmptyResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// StringPutMBCSResponse contains the response from method String.PutMBCS.
type StringPutMBCSResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// StringPutNullResponse contains the response from method String.PutNull.
type StringPutNullResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// StringPutWhitespaceResponse contains the response from method String.PutWhitespace.
type StringPutWhitespaceResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}