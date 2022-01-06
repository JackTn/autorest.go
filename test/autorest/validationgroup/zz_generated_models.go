//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package validationgroup

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

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
	Body *Product
}

// AutoRestValidationTestClientValidationOfMethodParametersOptions contains the optional parameters for the AutoRestValidationTestClient.ValidationOfMethodParameters
// method.
type AutoRestValidationTestClientValidationOfMethodParametersOptions struct {
	// placeholder for future optional parameters
}

// ChildProduct - The product documentation.
type ChildProduct struct {
	// REQUIRED; Constant string
	ConstProperty *string `json:"constProperty,omitempty"`

	// Count
	Count *int32 `json:"count,omitempty"`
}

// ConstantProduct - The product documentation.
type ConstantProduct struct {
	// REQUIRED; Constant string
	ConstProperty *string `json:"constProperty,omitempty"`

	// REQUIRED; Constant string2
	ConstProperty2 *string `json:"constProperty2,omitempty"`
}

// Implements the error and azcore.HTTPResponse interfaces.
type Error struct {
	raw     string
	Code    *int32  `json:"code,omitempty"`
	Fields  *string `json:"fields,omitempty"`
	Message *string `json:"message,omitempty"`
}

// Error implements the error interface for type Error.
// The contents of the error text are not contractual and subject to change.
func (e Error) Error() string {
	return e.raw
}

// Product - The product documentation.
type Product struct {
	// REQUIRED; The product documentation.
	Child *ChildProduct `json:"child,omitempty"`

	// REQUIRED; The product documentation.
	ConstChild *ConstantProduct `json:"constChild,omitempty"`

	// REQUIRED; Constant int
	ConstInt *int32 `json:"constInt,omitempty"`

	// REQUIRED; Constant string
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

// MarshalJSON implements the json.Marshaller interface for type Product.
func (p Product) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "capacity", p.Capacity)
	populate(objectMap, "child", p.Child)
	populate(objectMap, "constChild", p.ConstChild)
	populate(objectMap, "constInt", p.ConstInt)
	populate(objectMap, "constString", p.ConstString)
	populate(objectMap, "constStringAsEnum", p.ConstStringAsEnum)
	populate(objectMap, "display_names", p.DisplayNames)
	populate(objectMap, "image", p.Image)
	return json.Marshal(objectMap)
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}
