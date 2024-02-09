//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package addlpropsgroup

// ExtendsUnknownAdditionalPropertiesDiscriminatedClassification provides polymorphic access to related types.
// Call the interface's GetExtendsUnknownAdditionalPropertiesDiscriminated() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *ExtendsUnknownAdditionalPropertiesDiscriminated, *ExtendsUnknownAdditionalPropertiesDiscriminatedDerived
type ExtendsUnknownAdditionalPropertiesDiscriminatedClassification interface {
	// GetExtendsUnknownAdditionalPropertiesDiscriminated returns the ExtendsUnknownAdditionalPropertiesDiscriminated content of the underlying type.
	GetExtendsUnknownAdditionalPropertiesDiscriminated() *ExtendsUnknownAdditionalPropertiesDiscriminated
}

// IsUnknownAdditionalPropertiesDiscriminatedClassification provides polymorphic access to related types.
// Call the interface's GetIsUnknownAdditionalPropertiesDiscriminated() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *IsUnknownAdditionalPropertiesDiscriminated, *IsUnknownAdditionalPropertiesDiscriminatedDerived
type IsUnknownAdditionalPropertiesDiscriminatedClassification interface {
	// GetIsUnknownAdditionalPropertiesDiscriminated returns the IsUnknownAdditionalPropertiesDiscriminated content of the underlying type.
	GetIsUnknownAdditionalPropertiesDiscriminated() *IsUnknownAdditionalPropertiesDiscriminated
}
