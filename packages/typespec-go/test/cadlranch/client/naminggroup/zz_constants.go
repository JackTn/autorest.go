// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package naminggroup

const host = "http://localhost:3000"

type ClientExtensibleEnum string

const (
	ClientExtensibleEnumEnumValue1 ClientExtensibleEnum = "value1"
)

// PossibleClientExtensibleEnumValues returns the possible values for the ClientExtensibleEnum const type.
func PossibleClientExtensibleEnumValues() []ClientExtensibleEnum {
	return []ClientExtensibleEnum{
		ClientExtensibleEnumEnumValue1,
	}
}

type ExtensibleEnum string

const (
	ExtensibleEnumClientEnumValue1 ExtensibleEnum = "value1"
	ExtensibleEnumClientEnumValue2 ExtensibleEnum = "value2"
)

// PossibleExtensibleEnumValues returns the possible values for the ExtensibleEnum const type.
func PossibleExtensibleEnumValues() []ExtensibleEnum {
	return []ExtensibleEnum{
		ExtensibleEnumClientEnumValue1,
		ExtensibleEnumClientEnumValue2,
	}
}
