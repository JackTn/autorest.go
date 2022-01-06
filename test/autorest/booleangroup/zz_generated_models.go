//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package booleangroup

// BoolClientGetFalseOptions contains the optional parameters for the BoolClient.GetFalse method.
type BoolClientGetFalseOptions struct {
	// placeholder for future optional parameters
}

// BoolClientGetInvalidOptions contains the optional parameters for the BoolClient.GetInvalid method.
type BoolClientGetInvalidOptions struct {
	// placeholder for future optional parameters
}

// BoolClientGetNullOptions contains the optional parameters for the BoolClient.GetNull method.
type BoolClientGetNullOptions struct {
	// placeholder for future optional parameters
}

// BoolClientGetTrueOptions contains the optional parameters for the BoolClient.GetTrue method.
type BoolClientGetTrueOptions struct {
	// placeholder for future optional parameters
}

// BoolClientPutFalseOptions contains the optional parameters for the BoolClient.PutFalse method.
type BoolClientPutFalseOptions struct {
	// placeholder for future optional parameters
}

// BoolClientPutTrueOptions contains the optional parameters for the BoolClient.PutTrue method.
type BoolClientPutTrueOptions struct {
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
