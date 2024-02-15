// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package enumdiscgroup

const host = "http://localhost:3000"

// DogKind - extensible enum type for discriminator
type DogKind string

const (
	// DogKindGolden - Species golden
	DogKindGolden DogKind = "golden"
)

// PossibleDogKindValues returns the possible values for the DogKind const type.
func PossibleDogKindValues() []DogKind {
	return []DogKind{
		DogKindGolden,
	}
}

// SnakeKind - fixed enum type for discriminator
type SnakeKind string

const (
	// SnakeKindCobra - Species cobra
	SnakeKindCobra SnakeKind = "cobra"
)

// PossibleSnakeKindValues returns the possible values for the SnakeKind const type.
func PossibleSnakeKindValues() []SnakeKind {
	return []SnakeKind{
		SnakeKindCobra,
	}
}
