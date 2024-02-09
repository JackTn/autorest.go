//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package addlpropsgroup

// The model extends from Record<float32> type.
type ExtendsFloatAdditionalProperties struct {
	// REQUIRED; The id property
	ID                   *float32
	AdditionalProperties map[string]*float32
}

// The model extends from Record<ModelForRecord> type.
type ExtendsModelAdditionalProperties struct {
	AdditionalProperties map[string]*ModelForRecord
}

// The model extends from Record<ModelForRecord[]> type.
type ExtendsModelArrayAdditionalProperties struct {
	AdditionalProperties map[string][]*ModelForRecord
}

// The model extends from Record<string> type.
type ExtendsStringAdditionalProperties struct {
	// REQUIRED; The name property
	Name                 *string
	AdditionalProperties map[string]*string
}

// The model extends from Record<unknown> type.
type ExtendsUnknownAdditionalProperties struct {
	// REQUIRED; The name property
	Name                 *string
	AdditionalProperties map[string]any
}

// The model extends from a type that extends from Record<unknown>.
type ExtendsUnknownAdditionalPropertiesDerived struct {
	// REQUIRED; The index property
	Index *int32

	// REQUIRED; The name property
	Name                 *string
	AdditionalProperties map[string]any

	// The age property
	Age *float32
}

// The model extends from Record<unknown> with a discriminator.
type ExtendsUnknownAdditionalPropertiesDiscriminated struct {
	// REQUIRED; The discriminator
	Kind *string

	// REQUIRED; The name property
	Name                 *string
	AdditionalProperties map[string]any
}

// GetExtendsUnknownAdditionalPropertiesDiscriminated implements the ExtendsUnknownAdditionalPropertiesDiscriminatedClassification
// interface for type ExtendsUnknownAdditionalPropertiesDiscriminated.
func (e *ExtendsUnknownAdditionalPropertiesDiscriminated) GetExtendsUnknownAdditionalPropertiesDiscriminated() *ExtendsUnknownAdditionalPropertiesDiscriminated {
	return e
}

// The derived discriminated type
type ExtendsUnknownAdditionalPropertiesDiscriminatedDerived struct {
	// REQUIRED; The index property
	Index *int32

	// CONSTANT; undefinedField has constant value "derived", any specified value is ignored.
	Kind *string

	// REQUIRED; The name property
	Name                 *string
	AdditionalProperties map[string]any

	// The age property
	Age *float32
}

// GetExtendsUnknownAdditionalPropertiesDiscriminated implements the ExtendsUnknownAdditionalPropertiesDiscriminatedClassification
// interface for type ExtendsUnknownAdditionalPropertiesDiscriminatedDerived.
func (e *ExtendsUnknownAdditionalPropertiesDiscriminatedDerived) GetExtendsUnknownAdditionalPropertiesDiscriminated() *ExtendsUnknownAdditionalPropertiesDiscriminated {
	return &ExtendsUnknownAdditionalPropertiesDiscriminated{
		AdditionalProperties: e.AdditionalProperties,
		Kind:                 e.Kind,
		Name:                 e.Name,
	}
}

// The model is from Record<float32> type.
type IsFloatAdditionalProperties struct {
	// REQUIRED; The id property
	ID                   *float32
	AdditionalProperties map[string]*float32
}

// The model is from Record<ModelForRecord> type.
type IsModelAdditionalProperties struct {
	AdditionalProperties map[string]*ModelForRecord
}

// The model is from Record<ModelForRecord[]> type.
type IsModelArrayAdditionalProperties struct {
	AdditionalProperties map[string][]*ModelForRecord
}

// The model is from Record<string> type.
type IsStringAdditionalProperties struct {
	// REQUIRED; The name property
	Name                 *string
	AdditionalProperties map[string]*string
}

// The model is from Record<unknown> type.
type IsUnknownAdditionalProperties struct {
	// REQUIRED; The name property
	Name                 *string
	AdditionalProperties map[string]any
}

// The model extends from a type that is Record<unknown> type
type IsUnknownAdditionalPropertiesDerived struct {
	// REQUIRED; The index property
	Index *int32

	// REQUIRED; The name property
	Name                 *string
	AdditionalProperties map[string]any

	// The age property
	Age *float32
}

// The model is Record<unknown> with a discriminator.
type IsUnknownAdditionalPropertiesDiscriminated struct {
	// REQUIRED; The discriminator
	Kind *string

	// REQUIRED; The name property
	Name                 *string
	AdditionalProperties map[string]any
}

// GetIsUnknownAdditionalPropertiesDiscriminated implements the IsUnknownAdditionalPropertiesDiscriminatedClassification interface
// for type IsUnknownAdditionalPropertiesDiscriminated.
func (i *IsUnknownAdditionalPropertiesDiscriminated) GetIsUnknownAdditionalPropertiesDiscriminated() *IsUnknownAdditionalPropertiesDiscriminated {
	return i
}

// The derived discriminated type
type IsUnknownAdditionalPropertiesDiscriminatedDerived struct {
	// REQUIRED; The index property
	Index *int32

	// CONSTANT; undefinedField has constant value "derived", any specified value is ignored.
	Kind *string

	// REQUIRED; The name property
	Name                 *string
	AdditionalProperties map[string]any

	// The age property
	Age *float32
}

// GetIsUnknownAdditionalPropertiesDiscriminated implements the IsUnknownAdditionalPropertiesDiscriminatedClassification interface
// for type IsUnknownAdditionalPropertiesDiscriminatedDerived.
func (i *IsUnknownAdditionalPropertiesDiscriminatedDerived) GetIsUnknownAdditionalPropertiesDiscriminated() *IsUnknownAdditionalPropertiesDiscriminated {
	return &IsUnknownAdditionalPropertiesDiscriminated{
		AdditionalProperties: i.AdditionalProperties,
		Kind:                 i.Kind,
		Name:                 i.Name,
	}
}

// model for record
type ModelForRecord struct {
	// REQUIRED; The state property
	State *string
}
