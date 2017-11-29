package dategroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
)

// DateModel
type DateModel struct {
	autorest.Response `json:"-"`
	Value             *date.Date `json:"value,omitempty"`
}

// Error
type Error struct {
	Status  *int32  `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
}
