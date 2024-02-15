// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package mediatypesgroup

const host = "http://localhost:3000"

// ContentType - Content type for upload
type ContentType string

const (
	// ContentTypeApplicationPDF - Content Type 'application/pdf'
	ContentTypeApplicationPDF ContentType = "application/pdf"
	// ContentTypeImageJPEG - Content Type 'image/jpeg'
	ContentTypeImageJPEG ContentType = "image/jpeg"
	// ContentTypeImagePNG - Content Type 'image/png'
	ContentTypeImagePNG ContentType = "image/png"
	// ContentTypeImageTIFF - Content Type 'image/tiff'
	ContentTypeImageTIFF ContentType = "image/tiff"
)

// PossibleContentTypeValues returns the possible values for the ContentType const type.
func PossibleContentTypeValues() []ContentType {
	return []ContentType{
		ContentTypeApplicationPDF,
		ContentTypeImageJPEG,
		ContentTypeImagePNG,
		ContentTypeImageTIFF,
	}
}

// ContentType1 - Content type for upload
type ContentType1 string

const (
	// ContentType1ApplicationJSON - Content Type 'application/json'
	ContentType1ApplicationJSON ContentType1 = "application/json"
	// ContentType1ApplicationOctetStream - Content Type 'application/octet-stream'
	ContentType1ApplicationOctetStream ContentType1 = "application/octet-stream"
)

// PossibleContentType1Values returns the possible values for the ContentType1 const type.
func PossibleContentType1Values() []ContentType1 {
	return []ContentType1{
		ContentType1ApplicationJSON,
		ContentType1ApplicationOctetStream,
	}
}

// ContentType2 - Content type for upload
type ContentType2 string

const (
	// ContentType2ApplicationJSON - Content Type 'application/json'
	ContentType2ApplicationJSON ContentType2 = "application/json"
	// ContentType2ApplicationOctetStream - Content Type 'application/octet-stream'
	ContentType2ApplicationOctetStream ContentType2 = "application/octet-stream"
	// ContentType2TextPlain - Content Type 'text/plain'
	ContentType2TextPlain ContentType2 = "text/plain"
)

// PossibleContentType2Values returns the possible values for the ContentType2 const type.
func PossibleContentType2Values() []ContentType2 {
	return []ContentType2{
		ContentType2ApplicationJSON,
		ContentType2ApplicationOctetStream,
		ContentType2TextPlain,
	}
}

// ContentType3 - Content type for upload
type ContentType3 string

const (
	// ContentType3ApplicationJSON - Content Type 'application/json'
	ContentType3ApplicationJSON ContentType3 = "application/json"
	// ContentType3TextPlain - Content Type 'text/plain'
	ContentType3TextPlain ContentType3 = "text/plain"
)

// PossibleContentType3Values returns the possible values for the ContentType3 const type.
func PossibleContentType3Values() []ContentType3 {
	return []ContentType3{
		ContentType3ApplicationJSON,
		ContentType3TextPlain,
	}
}
