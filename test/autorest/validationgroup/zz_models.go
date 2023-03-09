//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package validationgroup

// AutoRestValidationTestClientGetWithConstantInPathOptions contains the optional parameters for the AutoRestValidationTestClient.GetWithConstantInPath
// method.
type AutoRestValidationTestClientGetWithConstantInPathOptions struct {
	// placeholder for future optional parameters
}

// AutoRestValidationTestClientPostWithConstantInBodyOptions contains the optional parameters for the AutoRestValidationTestClient.PostWithConstantInBody
// method.
type AutoRestValidationTestClientPostWithConstantInBodyOptions struct {
	Body *Product
}

// AutoRestValidationTestClientValidationOfBodyOptions contains the optional parameters for the AutoRestValidationTestClient.ValidationOfBody
// method.
type AutoRestValidationTestClientValidationOfBodyOptions struct {
	// placeholder for future optional parameters
}

// AutoRestValidationTestClientValidationOfMethodParametersOptions contains the optional parameters for the AutoRestValidationTestClient.ValidationOfMethodParameters
// method.
type AutoRestValidationTestClientValidationOfMethodParametersOptions struct {
	// placeholder for future optional parameters
}

// ChildProduct - The product documentation.
type ChildProduct struct {
	// CONSTANT; Constant string
	// Field has constant value "constant", any specified value is ignored.
	ConstProperty *string `json:"constProperty,omitempty"`

	// Count
	Count *int32 `json:"count,omitempty"`
}

// ConstantProduct - The product documentation.
type ConstantProduct struct {
	// CONSTANT; Constant string
	// Field has constant value "constant", any specified value is ignored.
	ConstProperty *string `json:"constProperty,omitempty"`

	// CONSTANT; Constant string2
	// Field has constant value "constant2", any specified value is ignored.
	ConstProperty2 *string `json:"constProperty2,omitempty"`
}

// Product - The product documentation.
type Product struct {
	// REQUIRED; The product documentation.
	Child *ChildProduct `json:"child,omitempty"`

	// REQUIRED; The product documentation.
	ConstChild *ConstantProduct `json:"constChild,omitempty"`

	// CONSTANT; Constant int
	// Field has constant value 0, any specified value is ignored.
	ConstInt *int32 `json:"constInt,omitempty"`

	// CONSTANT; Constant string
	// Field has constant value "constant", any specified value is ignored.
	ConstString *string `json:"constString,omitempty"`

	// Non required int betwen 0 and 100 exclusive.
	Capacity *int32 `json:"capacity,omitempty"`

	// Constant string as Enum
	ConstStringAsEnum *string `json:"constStringAsEnum,omitempty"`

	// Non required array of unique items from 0 to 6 elements.
	DisplayNames []*string `json:"display_names,omitempty"`

	// Image URL representing the product.
	Image *string `json:"image,omitempty"`
}