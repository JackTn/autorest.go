// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package additionalpropsgroup

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"reflect"
)

type CatAPTrue struct {
	PetAPTrue
	Friendly *bool `json:"friendly,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type CatAPTrue.
func (c CatAPTrue) MarshalJSON() ([]byte, error) {
	objectMap := c.PetAPTrue.marshalInternal()
	populate(objectMap, "friendly", c.Friendly)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CatAPTrue.
func (c *CatAPTrue) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "friendly":
			err = unpopulate(val, &c.Friendly)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return c.PetAPTrue.unmarshalInternal(rawMsg)
}

// CatAPTrueResponse is the response envelope for operations that return a CatAPTrue type.
type CatAPTrueResponse struct {
	CatAPTrue *CatAPTrue

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type Error struct {
	Message *string `json:"message,omitempty"`
	Status  *int32  `json:"status,omitempty"`
}

// Error implements the error interface for type Error.
func (e Error) Error() string {
	msg := ""
	if e.Message != nil {
		msg += fmt.Sprintf("Message: %v\n", *e.Message)
	}
	if e.Status != nil {
		msg += fmt.Sprintf("Status: %v\n", *e.Status)
	}
	if msg == "" {
		msg = "missing error info"
	}
	return msg
}

type PetAPInProperties struct {
	// Dictionary of
	AdditionalProperties *map[string]*float32 `json:"additionalProperties,omitempty"`
	ID                   *int32               `json:"id,omitempty"`
	Name                 *string              `json:"name,omitempty"`

	// Status - READ-ONLY
	Status *bool `json:"status,omitempty" azure:"ro"`
}

// PetAPInPropertiesResponse is the response envelope for operations that return a PetAPInProperties type.
type PetAPInPropertiesResponse struct {
	PetAPInProperties *PetAPInProperties

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type PetAPInPropertiesWithAPString struct {
	// Contains additional key/value pairs not defined in the schema.
	AdditionalProperties *map[string]*string

	// Dictionary of
	AdditionalProperties1 *map[string]*float32 `json:"additionalProperties,omitempty"`
	ID                    *int32               `json:"id,omitempty"`
	Name                  *string              `json:"name,omitempty"`
	OdataLocation         *string              `json:"@odata.location,omitempty"`

	// Status - READ-ONLY
	Status *bool `json:"status,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type PetAPInPropertiesWithAPString.
func (p PetAPInPropertiesWithAPString) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "additionalProperties", p.AdditionalProperties1)
	populate(objectMap, "id", p.ID)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "@odata.location", p.OdataLocation)
	populate(objectMap, "status", p.Status)
	if p.AdditionalProperties != nil {
		for key, val := range *p.AdditionalProperties {
			objectMap[key] = val
		}
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PetAPInPropertiesWithAPString.
func (p *PetAPInPropertiesWithAPString) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "additionalProperties":
			err = unpopulate(val, &p.AdditionalProperties1)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, &p.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, &p.Name)
			delete(rawMsg, key)
		case "@odata.location":
			err = unpopulate(val, &p.OdataLocation)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, &p.Status)
			delete(rawMsg, key)
		default:
			if p.AdditionalProperties == nil {
				p.AdditionalProperties = &map[string]*string{}
			}
			if val != nil {
				var aux string
				err = json.Unmarshal(*val, &aux)
				(*p.AdditionalProperties)[key] = &aux
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// PetAPInPropertiesWithAPStringResponse is the response envelope for operations that return a PetAPInPropertiesWithAPString type.
type PetAPInPropertiesWithAPStringResponse struct {
	PetAPInPropertiesWithAPString *PetAPInPropertiesWithAPString

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type PetAPObject struct {
	// Contains additional key/value pairs not defined in the schema.
	AdditionalProperties *map[string]interface{}
	ID                   *int32  `json:"id,omitempty"`
	Name                 *string `json:"name,omitempty"`

	// Status - READ-ONLY
	Status *bool `json:"status,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type PetAPObject.
func (p PetAPObject) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", p.ID)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "status", p.Status)
	if p.AdditionalProperties != nil {
		for key, val := range *p.AdditionalProperties {
			objectMap[key] = val
		}
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PetAPObject.
func (p *PetAPObject) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, &p.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, &p.Name)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, &p.Status)
			delete(rawMsg, key)
		default:
			if p.AdditionalProperties == nil {
				p.AdditionalProperties = &map[string]interface{}{}
			}
			if val != nil {
				var aux interface{}
				err = json.Unmarshal(*val, &aux)
				(*p.AdditionalProperties)[key] = aux
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// PetAPObjectResponse is the response envelope for operations that return a PetAPObject type.
type PetAPObjectResponse struct {
	PetAPObject *PetAPObject

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type PetAPString struct {
	// Contains additional key/value pairs not defined in the schema.
	AdditionalProperties *map[string]*string
	ID                   *int32  `json:"id,omitempty"`
	Name                 *string `json:"name,omitempty"`

	// Status - READ-ONLY
	Status *bool `json:"status,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type PetAPString.
func (p PetAPString) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", p.ID)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "status", p.Status)
	if p.AdditionalProperties != nil {
		for key, val := range *p.AdditionalProperties {
			objectMap[key] = val
		}
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PetAPString.
func (p *PetAPString) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, &p.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, &p.Name)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, &p.Status)
			delete(rawMsg, key)
		default:
			if p.AdditionalProperties == nil {
				p.AdditionalProperties = &map[string]*string{}
			}
			if val != nil {
				var aux string
				err = json.Unmarshal(*val, &aux)
				(*p.AdditionalProperties)[key] = &aux
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// PetAPStringResponse is the response envelope for operations that return a PetAPString type.
type PetAPStringResponse struct {
	PetAPString *PetAPString

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type PetAPTrue struct {
	// Contains additional key/value pairs not defined in the schema.
	AdditionalProperties *map[string]interface{}
	ID                   *int32  `json:"id,omitempty"`
	Name                 *string `json:"name,omitempty"`

	// Status - READ-ONLY
	Status *bool `json:"status,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type PetAPTrue.
func (p PetAPTrue) MarshalJSON() ([]byte, error) {
	objectMap := p.marshalInternal()
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PetAPTrue.
func (p *PetAPTrue) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	return p.unmarshalInternal(rawMsg)
}

func (p PetAPTrue) marshalInternal() map[string]interface{} {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", p.ID)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "status", p.Status)
	if p.AdditionalProperties != nil {
		for key, val := range *p.AdditionalProperties {
			objectMap[key] = val
		}
	}
	return objectMap
}

func (p *PetAPTrue) unmarshalInternal(rawMsg map[string]*json.RawMessage) error {
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, &p.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, &p.Name)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, &p.Status)
			delete(rawMsg, key)
		default:
			if p.AdditionalProperties == nil {
				p.AdditionalProperties = &map[string]interface{}{}
			}
			if val != nil {
				var aux interface{}
				err = json.Unmarshal(*val, &aux)
				(*p.AdditionalProperties)[key] = aux
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// PetAPTrueResponse is the response envelope for operations that return a PetAPTrue type.
type PetAPTrueResponse struct {
	PetAPTrue *PetAPTrue

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PetsCreateAPInPropertiesOptions contains the optional parameters for the Pets.CreateAPInProperties method.
type PetsCreateAPInPropertiesOptions struct {
	// placeholder for future optional parameters
}

// PetsCreateAPInPropertiesWithAPStringOptions contains the optional parameters for the Pets.CreateAPInPropertiesWithAPString method.
type PetsCreateAPInPropertiesWithAPStringOptions struct {
	// placeholder for future optional parameters
}

// PetsCreateAPObjectOptions contains the optional parameters for the Pets.CreateAPObject method.
type PetsCreateAPObjectOptions struct {
	// placeholder for future optional parameters
}

// PetsCreateAPStringOptions contains the optional parameters for the Pets.CreateAPString method.
type PetsCreateAPStringOptions struct {
	// placeholder for future optional parameters
}

// PetsCreateAPTrueOptions contains the optional parameters for the Pets.CreateAPTrue method.
type PetsCreateAPTrueOptions struct {
	// placeholder for future optional parameters
}

// PetsCreateCatAPTrueOptions contains the optional parameters for the Pets.CreateCatAPTrue method.
type PetsCreateCatAPTrueOptions struct {
	// placeholder for future optional parameters
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data *json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(*data, v)
}
