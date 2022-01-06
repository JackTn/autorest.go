//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package arraygroup

import (
	"net/http"
	"time"
)

// ArrayClientGetArrayEmptyResponse contains the response from method ArrayClient.GetArrayEmpty.
type ArrayClientGetArrayEmptyResponse struct {
	ArrayClientGetArrayEmptyResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetArrayEmptyResult contains the result from method ArrayClient.GetArrayEmpty.
type ArrayClientGetArrayEmptyResult struct {
	// An empty array []
	StringArrayArray [][]*string
}

// ArrayClientGetArrayItemEmptyResponse contains the response from method ArrayClient.GetArrayItemEmpty.
type ArrayClientGetArrayItemEmptyResponse struct {
	ArrayClientGetArrayItemEmptyResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetArrayItemEmptyResult contains the result from method ArrayClient.GetArrayItemEmpty.
type ArrayClientGetArrayItemEmptyResult struct {
	// An array of array of strings [['1', '2', '3'], [], ['7', '8', '9']]
	StringArrayArray [][]*string
}

// ArrayClientGetArrayItemNullResponse contains the response from method ArrayClient.GetArrayItemNull.
type ArrayClientGetArrayItemNullResponse struct {
	ArrayClientGetArrayItemNullResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetArrayItemNullResult contains the result from method ArrayClient.GetArrayItemNull.
type ArrayClientGetArrayItemNullResult struct {
	// An array of array of strings [['1', '2', '3'], null, ['7', '8', '9']]
	StringArrayArray [][]*string
}

// ArrayClientGetArrayNullResponse contains the response from method ArrayClient.GetArrayNull.
type ArrayClientGetArrayNullResponse struct {
	ArrayClientGetArrayNullResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetArrayNullResult contains the result from method ArrayClient.GetArrayNull.
type ArrayClientGetArrayNullResult struct {
	// a null array
	StringArrayArray [][]*string
}

// ArrayClientGetArrayValidResponse contains the response from method ArrayClient.GetArrayValid.
type ArrayClientGetArrayValidResponse struct {
	ArrayClientGetArrayValidResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetArrayValidResult contains the result from method ArrayClient.GetArrayValid.
type ArrayClientGetArrayValidResult struct {
	// An array of array of strings [['1', '2', '3'], ['4', '5', '6'], ['7', '8', '9']]
	StringArrayArray [][]*string
}

// ArrayClientGetBase64URLResponse contains the response from method ArrayClient.GetBase64URL.
type ArrayClientGetBase64URLResponse struct {
	ArrayClientGetBase64URLResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetBase64URLResult contains the result from method ArrayClient.GetBase64URL.
type ArrayClientGetBase64URLResult struct {
	// Get array value ['a string that gets encoded with base64url', 'test string' 'Lorem ipsum'] with the items base64url encoded
	ByteArrayArray [][]byte
}

// ArrayClientGetBooleanInvalidNullResponse contains the response from method ArrayClient.GetBooleanInvalidNull.
type ArrayClientGetBooleanInvalidNullResponse struct {
	ArrayClientGetBooleanInvalidNullResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetBooleanInvalidNullResult contains the result from method ArrayClient.GetBooleanInvalidNull.
type ArrayClientGetBooleanInvalidNullResult struct {
	// The array value [true, null, false]
	BoolArray []*bool
}

// ArrayClientGetBooleanInvalidStringResponse contains the response from method ArrayClient.GetBooleanInvalidString.
type ArrayClientGetBooleanInvalidStringResponse struct {
	ArrayClientGetBooleanInvalidStringResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetBooleanInvalidStringResult contains the result from method ArrayClient.GetBooleanInvalidString.
type ArrayClientGetBooleanInvalidStringResult struct {
	// The array value [true, 'boolean', false]
	BoolArray []*bool
}

// ArrayClientGetBooleanTfftResponse contains the response from method ArrayClient.GetBooleanTfft.
type ArrayClientGetBooleanTfftResponse struct {
	ArrayClientGetBooleanTfftResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetBooleanTfftResult contains the result from method ArrayClient.GetBooleanTfft.
type ArrayClientGetBooleanTfftResult struct {
	// The array value [true, false, false, true]
	BoolArray []*bool
}

// ArrayClientGetByteInvalidNullResponse contains the response from method ArrayClient.GetByteInvalidNull.
type ArrayClientGetByteInvalidNullResponse struct {
	ArrayClientGetByteInvalidNullResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetByteInvalidNullResult contains the result from method ArrayClient.GetByteInvalidNull.
type ArrayClientGetByteInvalidNullResult struct {
	// The byte array value [hex(AB, AC, AD), null] with the first item base64 encoded
	ByteArrayArray [][]byte
}

// ArrayClientGetByteValidResponse contains the response from method ArrayClient.GetByteValid.
type ArrayClientGetByteValidResponse struct {
	ArrayClientGetByteValidResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetByteValidResult contains the result from method ArrayClient.GetByteValid.
type ArrayClientGetByteValidResult struct {
	// The array value [hex(FF FF FF FA), hex(01 02 03), hex (25, 29, 43)] with each elementencoded in base 64
	ByteArrayArray [][]byte
}

// ArrayClientGetComplexEmptyResponse contains the response from method ArrayClient.GetComplexEmpty.
type ArrayClientGetComplexEmptyResponse struct {
	ArrayClientGetComplexEmptyResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetComplexEmptyResult contains the result from method ArrayClient.GetComplexEmpty.
type ArrayClientGetComplexEmptyResult struct {
	// Empty array of complex type []
	ProductArray []*Product
}

// ArrayClientGetComplexItemEmptyResponse contains the response from method ArrayClient.GetComplexItemEmpty.
type ArrayClientGetComplexItemEmptyResponse struct {
	ArrayClientGetComplexItemEmptyResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetComplexItemEmptyResult contains the result from method ArrayClient.GetComplexItemEmpty.
type ArrayClientGetComplexItemEmptyResult struct {
	// Array of complex type with empty item [{'integer': 1 'string': '2'}, {}, {'integer': 5, 'string': '6'}]
	ProductArray []*Product
}

// ArrayClientGetComplexItemNullResponse contains the response from method ArrayClient.GetComplexItemNull.
type ArrayClientGetComplexItemNullResponse struct {
	ArrayClientGetComplexItemNullResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetComplexItemNullResult contains the result from method ArrayClient.GetComplexItemNull.
type ArrayClientGetComplexItemNullResult struct {
	// Array of complex type with null item [{'integer': 1 'string': '2'}, null, {'integer': 5, 'string': '6'}]
	ProductArray []*Product
}

// ArrayClientGetComplexNullResponse contains the response from method ArrayClient.GetComplexNull.
type ArrayClientGetComplexNullResponse struct {
	ArrayClientGetComplexNullResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetComplexNullResult contains the result from method ArrayClient.GetComplexNull.
type ArrayClientGetComplexNullResult struct {
	// array of complex type with null value
	ProductArray []*Product
}

// ArrayClientGetComplexValidResponse contains the response from method ArrayClient.GetComplexValid.
type ArrayClientGetComplexValidResponse struct {
	ArrayClientGetComplexValidResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetComplexValidResult contains the result from method ArrayClient.GetComplexValid.
type ArrayClientGetComplexValidResult struct {
	// array of complex type with [{'integer': 1 'string': '2'}, {'integer': 3, 'string': '4'}, {'integer': 5, 'string': '6'}]
	ProductArray []*Product
}

// ArrayClientGetDateInvalidCharsResponse contains the response from method ArrayClient.GetDateInvalidChars.
type ArrayClientGetDateInvalidCharsResponse struct {
	ArrayClientGetDateInvalidCharsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetDateInvalidCharsResult contains the result from method ArrayClient.GetDateInvalidChars.
type ArrayClientGetDateInvalidCharsResult struct {
	// The array value ['2011-03-22', 'date']
	TimeArray []*time.Time
}

// ArrayClientGetDateInvalidNullResponse contains the response from method ArrayClient.GetDateInvalidNull.
type ArrayClientGetDateInvalidNullResponse struct {
	ArrayClientGetDateInvalidNullResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetDateInvalidNullResult contains the result from method ArrayClient.GetDateInvalidNull.
type ArrayClientGetDateInvalidNullResult struct {
	// The array value ['2012-01-01', null, '1776-07-04']
	TimeArray []*time.Time
}

// ArrayClientGetDateTimeInvalidCharsResponse contains the response from method ArrayClient.GetDateTimeInvalidChars.
type ArrayClientGetDateTimeInvalidCharsResponse struct {
	ArrayClientGetDateTimeInvalidCharsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetDateTimeInvalidCharsResult contains the result from method ArrayClient.GetDateTimeInvalidChars.
type ArrayClientGetDateTimeInvalidCharsResult struct {
	// The array value ['2000-12-01t00:00:01z', 'date-time']
	TimeArray []*time.Time
}

// ArrayClientGetDateTimeInvalidNullResponse contains the response from method ArrayClient.GetDateTimeInvalidNull.
type ArrayClientGetDateTimeInvalidNullResponse struct {
	ArrayClientGetDateTimeInvalidNullResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetDateTimeInvalidNullResult contains the result from method ArrayClient.GetDateTimeInvalidNull.
type ArrayClientGetDateTimeInvalidNullResult struct {
	// The array value ['2000-12-01t00:00:01z', null]
	TimeArray []*time.Time
}

// ArrayClientGetDateTimeRFC1123ValidResponse contains the response from method ArrayClient.GetDateTimeRFC1123Valid.
type ArrayClientGetDateTimeRFC1123ValidResponse struct {
	ArrayClientGetDateTimeRFC1123ValidResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetDateTimeRFC1123ValidResult contains the result from method ArrayClient.GetDateTimeRFC1123Valid.
type ArrayClientGetDateTimeRFC1123ValidResult struct {
	// The array value ['Fri, 01 Dec 2000 00:00:01 GMT', 'Wed, 02 Jan 1980 00:11:35 GMT', 'Wed, 12 Oct 1492 10:15:01 GMT']
	TimeArray []*time.Time
}

// ArrayClientGetDateTimeValidResponse contains the response from method ArrayClient.GetDateTimeValid.
type ArrayClientGetDateTimeValidResponse struct {
	ArrayClientGetDateTimeValidResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetDateTimeValidResult contains the result from method ArrayClient.GetDateTimeValid.
type ArrayClientGetDateTimeValidResult struct {
	// The array value ['2000-12-01t00:00:01z', '1980-01-02T00:11:35+01:00', '1492-10-12T10:15:01-08:00']
	TimeArray []*time.Time
}

// ArrayClientGetDateValidResponse contains the response from method ArrayClient.GetDateValid.
type ArrayClientGetDateValidResponse struct {
	ArrayClientGetDateValidResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetDateValidResult contains the result from method ArrayClient.GetDateValid.
type ArrayClientGetDateValidResult struct {
	// The array value ['2000-12-01', '1980-01-02', '1492-10-12']
	TimeArray []*time.Time
}

// ArrayClientGetDictionaryEmptyResponse contains the response from method ArrayClient.GetDictionaryEmpty.
type ArrayClientGetDictionaryEmptyResponse struct {
	ArrayClientGetDictionaryEmptyResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetDictionaryEmptyResult contains the result from method ArrayClient.GetDictionaryEmpty.
type ArrayClientGetDictionaryEmptyResult struct {
	// An array of Dictionaries of type <string, string> with value []
	MapOfStringArray []map[string]*string
}

// ArrayClientGetDictionaryItemEmptyResponse contains the response from method ArrayClient.GetDictionaryItemEmpty.
type ArrayClientGetDictionaryItemEmptyResponse struct {
	ArrayClientGetDictionaryItemEmptyResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetDictionaryItemEmptyResult contains the result from method ArrayClient.GetDictionaryItemEmpty.
type ArrayClientGetDictionaryItemEmptyResult struct {
	// An array of Dictionaries of type <string, string> with value [{'1': 'one', '2': 'two', '3': 'three'}, {}, {'7': 'seven',
	// '8': 'eight', '9': 'nine'}]
	MapOfStringArray []map[string]*string
}

// ArrayClientGetDictionaryItemNullResponse contains the response from method ArrayClient.GetDictionaryItemNull.
type ArrayClientGetDictionaryItemNullResponse struct {
	ArrayClientGetDictionaryItemNullResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetDictionaryItemNullResult contains the result from method ArrayClient.GetDictionaryItemNull.
type ArrayClientGetDictionaryItemNullResult struct {
	// An array of Dictionaries of type <string, string> with value [{'1': 'one', '2': 'two', '3': 'three'}, null, {'7': 'seven',
	// '8': 'eight', '9': 'nine'}]
	MapOfStringArray []map[string]*string
}

// ArrayClientGetDictionaryNullResponse contains the response from method ArrayClient.GetDictionaryNull.
type ArrayClientGetDictionaryNullResponse struct {
	ArrayClientGetDictionaryNullResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetDictionaryNullResult contains the result from method ArrayClient.GetDictionaryNull.
type ArrayClientGetDictionaryNullResult struct {
	// An array of Dictionaries with value null
	MapOfStringArray []map[string]*string
}

// ArrayClientGetDictionaryValidResponse contains the response from method ArrayClient.GetDictionaryValid.
type ArrayClientGetDictionaryValidResponse struct {
	ArrayClientGetDictionaryValidResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetDictionaryValidResult contains the result from method ArrayClient.GetDictionaryValid.
type ArrayClientGetDictionaryValidResult struct {
	// An array of Dictionaries of type <string, string> with value [{'1': 'one', '2': 'two', '3': 'three'}, {'4': 'four', '5':
	// 'five', '6': 'six'}, {'7': 'seven', '8': 'eight', '9': 'nine'}]
	MapOfStringArray []map[string]*string
}

// ArrayClientGetDoubleInvalidNullResponse contains the response from method ArrayClient.GetDoubleInvalidNull.
type ArrayClientGetDoubleInvalidNullResponse struct {
	ArrayClientGetDoubleInvalidNullResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetDoubleInvalidNullResult contains the result from method ArrayClient.GetDoubleInvalidNull.
type ArrayClientGetDoubleInvalidNullResult struct {
	// The array value [0.0, null, -1.2e20]
	Float64Array []*float64
}

// ArrayClientGetDoubleInvalidStringResponse contains the response from method ArrayClient.GetDoubleInvalidString.
type ArrayClientGetDoubleInvalidStringResponse struct {
	ArrayClientGetDoubleInvalidStringResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetDoubleInvalidStringResult contains the result from method ArrayClient.GetDoubleInvalidString.
type ArrayClientGetDoubleInvalidStringResult struct {
	// The array value [1.0, 'number', 0.0]
	Float64Array []*float64
}

// ArrayClientGetDoubleValidResponse contains the response from method ArrayClient.GetDoubleValid.
type ArrayClientGetDoubleValidResponse struct {
	ArrayClientGetDoubleValidResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetDoubleValidResult contains the result from method ArrayClient.GetDoubleValid.
type ArrayClientGetDoubleValidResult struct {
	// The array value [0, -0.01, 1.2e20]
	Float64Array []*float64
}

// ArrayClientGetDurationValidResponse contains the response from method ArrayClient.GetDurationValid.
type ArrayClientGetDurationValidResponse struct {
	ArrayClientGetDurationValidResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetDurationValidResult contains the result from method ArrayClient.GetDurationValid.
type ArrayClientGetDurationValidResult struct {
	// The array value ['P123DT22H14M12.011S', 'P5DT1H0M0S']
	StringArray []*string
}

// ArrayClientGetEmptyResponse contains the response from method ArrayClient.GetEmpty.
type ArrayClientGetEmptyResponse struct {
	ArrayClientGetEmptyResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetEmptyResult contains the result from method ArrayClient.GetEmpty.
type ArrayClientGetEmptyResult struct {
	// The empty array value []
	Int32Array []*int32
}

// ArrayClientGetEnumValidResponse contains the response from method ArrayClient.GetEnumValid.
type ArrayClientGetEnumValidResponse struct {
	ArrayClientGetEnumValidResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetEnumValidResult contains the result from method ArrayClient.GetEnumValid.
type ArrayClientGetEnumValidResult struct {
	// The array value ['foo1', 'foo2', 'foo3']
	FooEnumArray []*FooEnum
}

// ArrayClientGetFloatInvalidNullResponse contains the response from method ArrayClient.GetFloatInvalidNull.
type ArrayClientGetFloatInvalidNullResponse struct {
	ArrayClientGetFloatInvalidNullResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetFloatInvalidNullResult contains the result from method ArrayClient.GetFloatInvalidNull.
type ArrayClientGetFloatInvalidNullResult struct {
	// The array value [0.0, null, -1.2e20]
	Float32Array []*float32
}

// ArrayClientGetFloatInvalidStringResponse contains the response from method ArrayClient.GetFloatInvalidString.
type ArrayClientGetFloatInvalidStringResponse struct {
	ArrayClientGetFloatInvalidStringResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetFloatInvalidStringResult contains the result from method ArrayClient.GetFloatInvalidString.
type ArrayClientGetFloatInvalidStringResult struct {
	// The array value [1.0, 'number', 0.0]
	Float32Array []*float32
}

// ArrayClientGetFloatValidResponse contains the response from method ArrayClient.GetFloatValid.
type ArrayClientGetFloatValidResponse struct {
	ArrayClientGetFloatValidResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetFloatValidResult contains the result from method ArrayClient.GetFloatValid.
type ArrayClientGetFloatValidResult struct {
	// The array value [0, -0.01, 1.2e20]
	Float32Array []*float32
}

// ArrayClientGetIntInvalidNullResponse contains the response from method ArrayClient.GetIntInvalidNull.
type ArrayClientGetIntInvalidNullResponse struct {
	ArrayClientGetIntInvalidNullResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetIntInvalidNullResult contains the result from method ArrayClient.GetIntInvalidNull.
type ArrayClientGetIntInvalidNullResult struct {
	// The array value [1, null, 0]
	Int32Array []*int32
}

// ArrayClientGetIntInvalidStringResponse contains the response from method ArrayClient.GetIntInvalidString.
type ArrayClientGetIntInvalidStringResponse struct {
	ArrayClientGetIntInvalidStringResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetIntInvalidStringResult contains the result from method ArrayClient.GetIntInvalidString.
type ArrayClientGetIntInvalidStringResult struct {
	// The array value [1, 'integer', 0]
	Int32Array []*int32
}

// ArrayClientGetIntegerValidResponse contains the response from method ArrayClient.GetIntegerValid.
type ArrayClientGetIntegerValidResponse struct {
	ArrayClientGetIntegerValidResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetIntegerValidResult contains the result from method ArrayClient.GetIntegerValid.
type ArrayClientGetIntegerValidResult struct {
	// The array value [1, -1, 3, 300]
	Int32Array []*int32
}

// ArrayClientGetInvalidResponse contains the response from method ArrayClient.GetInvalid.
type ArrayClientGetInvalidResponse struct {
	ArrayClientGetInvalidResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetInvalidResult contains the result from method ArrayClient.GetInvalid.
type ArrayClientGetInvalidResult struct {
	// The invalid Array [1, 2, 3
	Int32Array []*int32
}

// ArrayClientGetLongInvalidNullResponse contains the response from method ArrayClient.GetLongInvalidNull.
type ArrayClientGetLongInvalidNullResponse struct {
	ArrayClientGetLongInvalidNullResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetLongInvalidNullResult contains the result from method ArrayClient.GetLongInvalidNull.
type ArrayClientGetLongInvalidNullResult struct {
	// The array value [1, null, 0]
	Int64Array []*int64
}

// ArrayClientGetLongInvalidStringResponse contains the response from method ArrayClient.GetLongInvalidString.
type ArrayClientGetLongInvalidStringResponse struct {
	ArrayClientGetLongInvalidStringResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetLongInvalidStringResult contains the result from method ArrayClient.GetLongInvalidString.
type ArrayClientGetLongInvalidStringResult struct {
	// The array value [1, 'integer', 0]
	Int64Array []*int64
}

// ArrayClientGetLongValidResponse contains the response from method ArrayClient.GetLongValid.
type ArrayClientGetLongValidResponse struct {
	ArrayClientGetLongValidResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetLongValidResult contains the result from method ArrayClient.GetLongValid.
type ArrayClientGetLongValidResult struct {
	// The array value [1, -1, 3, 300]
	Int64Array []*int64
}

// ArrayClientGetNullResponse contains the response from method ArrayClient.GetNull.
type ArrayClientGetNullResponse struct {
	ArrayClientGetNullResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetNullResult contains the result from method ArrayClient.GetNull.
type ArrayClientGetNullResult struct {
	// The null Array value
	Int32Array []*int32
}

// ArrayClientGetStringEnumValidResponse contains the response from method ArrayClient.GetStringEnumValid.
type ArrayClientGetStringEnumValidResponse struct {
	ArrayClientGetStringEnumValidResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetStringEnumValidResult contains the result from method ArrayClient.GetStringEnumValid.
type ArrayClientGetStringEnumValidResult struct {
	// The array value ['foo1', 'foo2', 'foo3']
	Enum0Array []*Enum0
}

// ArrayClientGetStringValidResponse contains the response from method ArrayClient.GetStringValid.
type ArrayClientGetStringValidResponse struct {
	ArrayClientGetStringValidResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetStringValidResult contains the result from method ArrayClient.GetStringValid.
type ArrayClientGetStringValidResult struct {
	// The array value ['foo1', 'foo2', 'foo3']
	StringArray []*string
}

// ArrayClientGetStringWithInvalidResponse contains the response from method ArrayClient.GetStringWithInvalid.
type ArrayClientGetStringWithInvalidResponse struct {
	ArrayClientGetStringWithInvalidResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetStringWithInvalidResult contains the result from method ArrayClient.GetStringWithInvalid.
type ArrayClientGetStringWithInvalidResult struct {
	// The array value ['foo', 123, 'foo2']
	StringArray []*string
}

// ArrayClientGetStringWithNullResponse contains the response from method ArrayClient.GetStringWithNull.
type ArrayClientGetStringWithNullResponse struct {
	ArrayClientGetStringWithNullResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetStringWithNullResult contains the result from method ArrayClient.GetStringWithNull.
type ArrayClientGetStringWithNullResult struct {
	// The array value ['foo', null, 'foo2']
	StringArray []*string
}

// ArrayClientGetUUIDInvalidCharsResponse contains the response from method ArrayClient.GetUUIDInvalidChars.
type ArrayClientGetUUIDInvalidCharsResponse struct {
	ArrayClientGetUUIDInvalidCharsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetUUIDInvalidCharsResult contains the result from method ArrayClient.GetUUIDInvalidChars.
type ArrayClientGetUUIDInvalidCharsResult struct {
	// The array value ['6dcc7237-45fe-45c4-8a6b-3a8a3f625652', 'foo']
	StringArray []*string
}

// ArrayClientGetUUIDValidResponse contains the response from method ArrayClient.GetUUIDValid.
type ArrayClientGetUUIDValidResponse struct {
	ArrayClientGetUUIDValidResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientGetUUIDValidResult contains the result from method ArrayClient.GetUUIDValid.
type ArrayClientGetUUIDValidResult struct {
	// The array value ['6dcc7237-45fe-45c4-8a6b-3a8a3f625652', 'd1399005-30f7-40d6-8da6-dd7c89ad34db', 'f42f6aa1-a5bc-4ddf-907e-5f915de43205']
	StringArray []*string
}

// ArrayClientPutArrayValidResponse contains the response from method ArrayClient.PutArrayValid.
type ArrayClientPutArrayValidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientPutBooleanTfftResponse contains the response from method ArrayClient.PutBooleanTfft.
type ArrayClientPutBooleanTfftResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientPutByteValidResponse contains the response from method ArrayClient.PutByteValid.
type ArrayClientPutByteValidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientPutComplexValidResponse contains the response from method ArrayClient.PutComplexValid.
type ArrayClientPutComplexValidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientPutDateTimeRFC1123ValidResponse contains the response from method ArrayClient.PutDateTimeRFC1123Valid.
type ArrayClientPutDateTimeRFC1123ValidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientPutDateTimeValidResponse contains the response from method ArrayClient.PutDateTimeValid.
type ArrayClientPutDateTimeValidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientPutDateValidResponse contains the response from method ArrayClient.PutDateValid.
type ArrayClientPutDateValidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientPutDictionaryValidResponse contains the response from method ArrayClient.PutDictionaryValid.
type ArrayClientPutDictionaryValidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientPutDoubleValidResponse contains the response from method ArrayClient.PutDoubleValid.
type ArrayClientPutDoubleValidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientPutDurationValidResponse contains the response from method ArrayClient.PutDurationValid.
type ArrayClientPutDurationValidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientPutEmptyResponse contains the response from method ArrayClient.PutEmpty.
type ArrayClientPutEmptyResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientPutEnumValidResponse contains the response from method ArrayClient.PutEnumValid.
type ArrayClientPutEnumValidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientPutFloatValidResponse contains the response from method ArrayClient.PutFloatValid.
type ArrayClientPutFloatValidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientPutIntegerValidResponse contains the response from method ArrayClient.PutIntegerValid.
type ArrayClientPutIntegerValidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientPutLongValidResponse contains the response from method ArrayClient.PutLongValid.
type ArrayClientPutLongValidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientPutStringEnumValidResponse contains the response from method ArrayClient.PutStringEnumValid.
type ArrayClientPutStringEnumValidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientPutStringValidResponse contains the response from method ArrayClient.PutStringValid.
type ArrayClientPutStringValidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArrayClientPutUUIDValidResponse contains the response from method ArrayClient.PutUUIDValid.
type ArrayClientPutUUIDValidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}
