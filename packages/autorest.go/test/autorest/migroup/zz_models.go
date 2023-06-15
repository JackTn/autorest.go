//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package migroup

type Cat struct {
	// REQUIRED
	Name      *string
	Hisses    *bool
	LikesMilk *bool
	Meows     *bool
}

type Feline struct {
	Hisses *bool
	Meows  *bool
}

type Horse struct {
	// REQUIRED
	Name         *string
	IsAShowHorse *bool
}

type Kitten struct {
	// REQUIRED
	Name        *string
	EatsMiceYet *bool
	Hisses      *bool
	LikesMilk   *bool
	Meows       *bool
}

type Pet struct {
	// REQUIRED
	Name *string
}
