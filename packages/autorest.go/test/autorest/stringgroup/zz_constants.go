// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package stringgroup

const host = "http://localhost:3000"

// Colors - Referenced Color Enum Description.
type Colors string

const (
	ColorsBlueColor  Colors = "blue_color"
	ColorsGreenColor Colors = "green-color"
	ColorsRedColor   Colors = "red color"
)

// PossibleColorsValues returns the possible values for the Colors const type.
func PossibleColorsValues() []Colors {
	return []Colors{
		ColorsBlueColor,
		ColorsGreenColor,
		ColorsRedColor,
	}
}
