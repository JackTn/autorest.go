// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package bytegroup

// ByteClientGetEmptyResponse contains the response from method ByteClient.GetEmpty.
type ByteClientGetEmptyResponse struct {
	// The empty byte value ''
	Value []byte
}

// ByteClientGetInvalidResponse contains the response from method ByteClient.GetInvalid.
type ByteClientGetInvalidResponse struct {
	// The invalid byte value '::::SWAGGER::::'
	Value []byte
}

// ByteClientGetNonASCIIResponse contains the response from method ByteClient.GetNonASCII.
type ByteClientGetNonASCIIResponse struct {
	// Non-ascii base-64 encoded byte string hex(FF FE FD FC FB FA F9 F8 F7 F6)
	Value []byte
}

// ByteClientGetNullResponse contains the response from method ByteClient.GetNull.
type ByteClientGetNullResponse struct {
	// The null byte value
	Value []byte
}

// ByteClientPutNonASCIIResponse contains the response from method ByteClient.PutNonASCII.
type ByteClientPutNonASCIIResponse struct {
	// placeholder for future response values
}
