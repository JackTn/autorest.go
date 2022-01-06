//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package bytegroup

// ByteClientGetEmptyOptions contains the optional parameters for the ByteClient.GetEmpty method.
type ByteClientGetEmptyOptions struct {
	// placeholder for future optional parameters
}

// ByteClientGetInvalidOptions contains the optional parameters for the ByteClient.GetInvalid method.
type ByteClientGetInvalidOptions struct {
	// placeholder for future optional parameters
}

// ByteClientGetNonASCIIOptions contains the optional parameters for the ByteClient.GetNonASCII method.
type ByteClientGetNonASCIIOptions struct {
	// placeholder for future optional parameters
}

// ByteClientGetNullOptions contains the optional parameters for the ByteClient.GetNull method.
type ByteClientGetNullOptions struct {
	// placeholder for future optional parameters
}

// ByteClientPutNonASCIIOptions contains the optional parameters for the ByteClient.PutNonASCII method.
type ByteClientPutNonASCIIOptions struct {
	// placeholder for future optional parameters
}

// Implements the error and azcore.HTTPResponse interfaces.
type Error struct {
	raw     string
	Message *string `json:"message,omitempty"`
	Status  *int32  `json:"status,omitempty"`
}

// Error implements the error interface for type Error.
// The contents of the error text are not contractual and subject to change.
func (e Error) Error() string {
	return e.raw
}
