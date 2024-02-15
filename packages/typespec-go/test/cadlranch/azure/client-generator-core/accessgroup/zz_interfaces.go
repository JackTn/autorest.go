// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package accessgroup

// AbstractModelClassification provides polymorphic access to related types.
// Call the interface's GetAbstractModel() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AbstractModel, *RealModel
type AbstractModelClassification interface {
	// GetAbstractModel returns the AbstractModel content of the underlying type.
	GetAbstractModel() *AbstractModel
}
