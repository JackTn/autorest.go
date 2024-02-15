// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package pageablegroup

// Paged collection of User items
type PagedUser struct {
	// REQUIRED; The User items on this page
	Value []*User

	// The link to the next page of items
	NextLink *string
}

// User model
type User struct {
	// REQUIRED; User name
	Name *string
}
