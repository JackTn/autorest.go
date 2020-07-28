// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexgroup

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ArrayWrapper struct {
	Array *[]string `json:"array,omitempty"`
}

// ArrayWrapperResponse is the response envelope for operations that return a ArrayWrapper type.
type ArrayWrapperResponse struct {
	ArrayWrapper *ArrayWrapper

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type Basic struct {
	Color *CMYKColors `json:"color,omitempty"`

	// Basic Id
	ID *int32 `json:"id,omitempty"`

	// Name property with a very long description that does not fit on a single line and a line break.
	Name *string `json:"name,omitempty"`
}

// BasicResponse is the response envelope for operations that return a Basic type.
type BasicResponse struct {
	Basic *Basic

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type BooleanWrapper struct {
	FieldFalse *bool `json:"field_false,omitempty"`
	FieldTrue  *bool `json:"field_true,omitempty"`
}

// BooleanWrapperResponse is the response envelope for operations that return a BooleanWrapper type.
type BooleanWrapperResponse struct {
	BooleanWrapper *BooleanWrapper

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type ByteWrapper struct {
	Field *[]byte `json:"field,omitempty"`
}

// ByteWrapperResponse is the response envelope for operations that return a ByteWrapper type.
type ByteWrapperResponse struct {
	ByteWrapper *ByteWrapper

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type Cat struct {
	Pet
	Color *string `json:"color,omitempty"`
	Hates *[]Dog  `json:"hates,omitempty"`
}

type Cookiecuttershark struct {
	Shark
}

// MarshalJSON implements the json.Marshaller interface for type Cookiecuttershark.
func (c Cookiecuttershark) MarshalJSON() ([]byte, error) {
	objectMap := c.Shark.marshalInternal("cookiecuttershark")
	return json.Marshal(objectMap)
}

type DateWrapper struct {
	Field *time.Time `json:"field,omitempty"`
	Leap  *time.Time `json:"leap,omitempty"`
}

// DateWrapperResponse is the response envelope for operations that return a DateWrapper type.
type DateWrapperResponse struct {
	DateWrapper *DateWrapper

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type DatetimeWrapper struct {
	Field *time.Time `json:"field,omitempty"`
	Now   *time.Time `json:"now,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type DatetimeWrapper.
func (d DatetimeWrapper) MarshalJSON() ([]byte, error) {
	type alias DatetimeWrapper
	aux := &struct {
		*alias
		Field *timeRFC3339 `json:"field"`
		Now   *timeRFC3339 `json:"now"`
	}{
		alias: (*alias)(&d),
		Field: (*timeRFC3339)(d.Field),
		Now:   (*timeRFC3339)(d.Now),
	}
	return json.Marshal(aux)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DatetimeWrapper.
func (d *DatetimeWrapper) UnmarshalJSON(data []byte) error {
	type alias DatetimeWrapper
	aux := &struct {
		*alias
		Field *timeRFC3339 `json:"field"`
		Now   *timeRFC3339 `json:"now"`
	}{
		alias: (*alias)(d),
	}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	d.Field = (*time.Time)(aux.Field)
	d.Now = (*time.Time)(aux.Now)
	return nil
}

// DatetimeWrapperResponse is the response envelope for operations that return a DatetimeWrapper type.
type DatetimeWrapperResponse struct {
	DatetimeWrapper *DatetimeWrapper

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type Datetimerfc1123Wrapper struct {
	Field *time.Time `json:"field,omitempty"`
	Now   *time.Time `json:"now,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type Datetimerfc1123Wrapper.
func (d Datetimerfc1123Wrapper) MarshalJSON() ([]byte, error) {
	type alias Datetimerfc1123Wrapper
	aux := &struct {
		*alias
		Field *timeRFC1123 `json:"field"`
		Now   *timeRFC1123 `json:"now"`
	}{
		alias: (*alias)(&d),
		Field: (*timeRFC1123)(d.Field),
		Now:   (*timeRFC1123)(d.Now),
	}
	return json.Marshal(aux)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Datetimerfc1123Wrapper.
func (d *Datetimerfc1123Wrapper) UnmarshalJSON(data []byte) error {
	type alias Datetimerfc1123Wrapper
	aux := &struct {
		*alias
		Field *timeRFC1123 `json:"field"`
		Now   *timeRFC1123 `json:"now"`
	}{
		alias: (*alias)(d),
	}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	d.Field = (*time.Time)(aux.Field)
	d.Now = (*time.Time)(aux.Now)
	return nil
}

// Datetimerfc1123WrapperResponse is the response envelope for operations that return a Datetimerfc1123Wrapper type.
type Datetimerfc1123WrapperResponse struct {
	Datetimerfc1123Wrapper *Datetimerfc1123Wrapper

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type DictionaryWrapper struct {
	// Dictionary of <string>
	DefaultProgram *map[string]string `json:"defaultProgram,omitempty"`
}

// DictionaryWrapperResponse is the response envelope for operations that return a DictionaryWrapper type.
type DictionaryWrapperResponse struct {
	DictionaryWrapper *DictionaryWrapper

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type Dog struct {
	Pet
	Food *string `json:"food,omitempty"`
}

// DotFishClassification provides polymorphic access to related types.
type DotFishClassification interface {
	GetDotFish() *DotFish
}

type DotFish struct {
	FishType *string `json:"fish.type,omitempty"`
	Species  *string `json:"species,omitempty"`
}

// GetDotFish implements the DotFishClassification interface for type DotFish.
func (d *DotFish) GetDotFish() *DotFish { return d }

// UnmarshalJSON implements the json.Unmarshaller interface for type DotFish.
func (d *DotFish) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "fish.type":
			if val != nil {
				err = json.Unmarshal(*val, &d.FishType)
			}
		case "species":
			if val != nil {
				err = json.Unmarshal(*val, &d.Species)
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (d DotFish) marshalInternal(discValue string) map[string]interface{} {
	objectMap := make(map[string]interface{})
	d.FishType = &discValue
	objectMap["fish.type"] = d.FishType
	if d.Species != nil {
		objectMap["species"] = d.Species
	}
	return objectMap
}

type DotFishMarket struct {
	Fishes       *[]DotFishClassification `json:"fishes,omitempty"`
	Salmons      *[]DotSalmon             `json:"salmons,omitempty"`
	SampleFish   DotFishClassification    `json:"sampleFish,omitempty"`
	SampleSalmon *DotSalmon               `json:"sampleSalmon,omitempty"`
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DotFishMarket.
func (d *DotFishMarket) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "fishes":
			if val != nil {
				d.Fishes, err = unmarshalDotFishClassificationArray(*val)
			}
		case "salmons":
			if val != nil {
				err = json.Unmarshal(*val, &d.Salmons)
			}
		case "sampleFish":
			if val != nil {
				d.SampleFish, err = unmarshalDotFishClassification(*val)
			}
		case "sampleSalmon":
			if val != nil {
				err = json.Unmarshal(*val, &d.SampleSalmon)
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// DotFishMarketResponse is the response envelope for operations that return a DotFishMarket type.
type DotFishMarketResponse struct {
	DotFishMarket *DotFishMarket

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DotFishResponse is the response envelope for operations that return a DotFish type.
type DotFishResponse struct {
	DotFish DotFishClassification

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DotFishResponse.
func (d *DotFishResponse) UnmarshalJSON(data []byte) error {
	t, err := unmarshalDotFishClassification(data)
	if err != nil {
		return err
	}
	d.DotFish = t
	return nil
}

type DotSalmon struct {
	DotFish
	Iswild   *bool   `json:"iswild,omitempty"`
	Location *string `json:"location,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type DotSalmon.
func (d DotSalmon) MarshalJSON() ([]byte, error) {
	objectMap := d.DotFish.marshalInternal("DotSalmon")
	if d.Iswild != nil {
		objectMap["iswild"] = d.Iswild
	}
	if d.Location != nil {
		objectMap["location"] = d.Location
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DotSalmon.
func (d *DotSalmon) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "iswild":
			if val != nil {
				err = json.Unmarshal(*val, &d.Iswild)
			}
		case "location":
			if val != nil {
				err = json.Unmarshal(*val, &d.Location)
			}
		}
		if err != nil {
			return err
		}
	}
	return json.Unmarshal(data, &d.DotFish)
}

type DoubleWrapper struct {
	Field1                                                                          *float64 `json:"field1,omitempty"`
	Field56ZerosAfterTheDotAndNegativeZeroBeforeDotAndThisIsALongFieldNameOnPurpose *float64 `json:"field_56_zeros_after_the_dot_and_negative_zero_before_dot_and_this_is_a_long_field_name_on_purpose,omitempty"`
}

// DoubleWrapperResponse is the response envelope for operations that return a DoubleWrapper type.
type DoubleWrapperResponse struct {
	DoubleWrapper *DoubleWrapper

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type DurationWrapper struct {
	Field *time.Duration `json:"field,omitempty"`
}

// DurationWrapperResponse is the response envelope for operations that return a DurationWrapper type.
type DurationWrapperResponse struct {
	DurationWrapper *DurationWrapper

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

// FishClassification provides polymorphic access to related types.
type FishClassification interface {
	GetFish() *Fish
}

type Fish struct {
	Fishtype *string               `json:"fishtype,omitempty"`
	Length   *float32              `json:"length,omitempty"`
	Siblings *[]FishClassification `json:"siblings,omitempty"`
	Species  *string               `json:"species,omitempty"`
}

// GetFish implements the FishClassification interface for type Fish.
func (f *Fish) GetFish() *Fish { return f }

// UnmarshalJSON implements the json.Unmarshaller interface for type Fish.
func (f *Fish) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "fishtype":
			if val != nil {
				err = json.Unmarshal(*val, &f.Fishtype)
			}
		case "length":
			if val != nil {
				err = json.Unmarshal(*val, &f.Length)
			}
		case "siblings":
			if val != nil {
				f.Siblings, err = unmarshalFishClassificationArray(*val)
			}
		case "species":
			if val != nil {
				err = json.Unmarshal(*val, &f.Species)
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (f Fish) marshalInternal(discValue string) map[string]interface{} {
	objectMap := make(map[string]interface{})
	f.Fishtype = &discValue
	objectMap["fishtype"] = f.Fishtype
	if f.Length != nil {
		objectMap["length"] = f.Length
	}
	if f.Siblings != nil {
		objectMap["siblings"] = f.Siblings
	}
	if f.Species != nil {
		objectMap["species"] = f.Species
	}
	return objectMap
}

// FishResponse is the response envelope for operations that return a Fish type.
type FishResponse struct {
	Fish FishClassification

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// UnmarshalJSON implements the json.Unmarshaller interface for type FishResponse.
func (f *FishResponse) UnmarshalJSON(data []byte) error {
	t, err := unmarshalFishClassification(data)
	if err != nil {
		return err
	}
	f.Fish = t
	return nil
}

type FloatWrapper struct {
	Field1 *float32 `json:"field1,omitempty"`
	Field2 *float32 `json:"field2,omitempty"`
}

// FloatWrapperResponse is the response envelope for operations that return a FloatWrapper type.
type FloatWrapperResponse struct {
	FloatWrapper *FloatWrapper

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type Goblinshark struct {
	Shark
	// Colors possible
	Color   *GoblinSharkColor `json:"color,omitempty"`
	Jawsize *int32            `json:"jawsize,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type Goblinshark.
func (g Goblinshark) MarshalJSON() ([]byte, error) {
	objectMap := g.Shark.marshalInternal("goblin")
	if g.Color != nil {
		objectMap["color"] = g.Color
	}
	if g.Jawsize != nil {
		objectMap["jawsize"] = g.Jawsize
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Goblinshark.
func (g *Goblinshark) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "color":
			if val != nil {
				err = json.Unmarshal(*val, &g.Color)
			}
		case "jawsize":
			if val != nil {
				err = json.Unmarshal(*val, &g.Jawsize)
			}
		}
		if err != nil {
			return err
		}
	}
	return json.Unmarshal(data, &g.Shark)
}

type IntWrapper struct {
	Field1 *int32 `json:"field1,omitempty"`
	Field2 *int32 `json:"field2,omitempty"`
}

// IntWrapperResponse is the response envelope for operations that return a IntWrapper type.
type IntWrapperResponse struct {
	IntWrapper *IntWrapper

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type LongWrapper struct {
	Field1 *int64 `json:"field1,omitempty"`
	Field2 *int64 `json:"field2,omitempty"`
}

// LongWrapperResponse is the response envelope for operations that return a LongWrapper type.
type LongWrapperResponse struct {
	LongWrapper *LongWrapper

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type MyBaseHelperType struct {
	PropBh1 *string `json:"propBH1,omitempty"`
}

// MyBaseTypeClassification provides polymorphic access to related types.
type MyBaseTypeClassification interface {
	GetMyBaseType() *MyBaseType
}

type MyBaseType struct {
	Helper *MyBaseHelperType `json:"helper,omitempty"`
	Kind   *string           `json:"kind,omitempty"`
	PropB1 *string           `json:"propB1,omitempty"`
}

// GetMyBaseType implements the MyBaseTypeClassification interface for type MyBaseType.
func (m *MyBaseType) GetMyBaseType() *MyBaseType { return m }

// UnmarshalJSON implements the json.Unmarshaller interface for type MyBaseType.
func (m *MyBaseType) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "helper":
			if val != nil {
				err = json.Unmarshal(*val, &m.Helper)
			}
		case "kind":
			if val != nil {
				err = json.Unmarshal(*val, &m.Kind)
			}
		case "propB1":
			if val != nil {
				err = json.Unmarshal(*val, &m.PropB1)
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (m MyBaseType) marshalInternal(discValue string) map[string]interface{} {
	objectMap := make(map[string]interface{})
	if m.Helper != nil {
		objectMap["helper"] = m.Helper
	}
	m.Kind = &discValue
	objectMap["kind"] = m.Kind
	if m.PropB1 != nil {
		objectMap["propB1"] = m.PropB1
	}
	return objectMap
}

// MyBaseTypeResponse is the response envelope for operations that return a MyBaseType type.
type MyBaseTypeResponse struct {
	MyBaseType MyBaseTypeClassification

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MyBaseTypeResponse.
func (m *MyBaseTypeResponse) UnmarshalJSON(data []byte) error {
	t, err := unmarshalMyBaseTypeClassification(data)
	if err != nil {
		return err
	}
	m.MyBaseType = t
	return nil
}

type MyDerivedType struct {
	MyBaseType
	PropD1 *string `json:"propD1,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type MyDerivedType.
func (m MyDerivedType) MarshalJSON() ([]byte, error) {
	objectMap := m.MyBaseType.marshalInternal("Kind1")
	if m.PropD1 != nil {
		objectMap["propD1"] = m.PropD1
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MyDerivedType.
func (m *MyDerivedType) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "propD1":
			if val != nil {
				err = json.Unmarshal(*val, &m.PropD1)
			}
		}
		if err != nil {
			return err
		}
	}
	return json.Unmarshal(data, &m.MyBaseType)
}

type Pet struct {
	ID   *int32  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type ReadonlyObj struct {
	ID   *string `json:"id,omitempty" azure:"ro"`
	Size *int32  `json:"size,omitempty"`
}

// ReadonlyObjResponse is the response envelope for operations that return a ReadonlyObj type.
type ReadonlyObjResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
	ReadonlyObj *ReadonlyObj
}

// SalmonClassification provides polymorphic access to related types.
type SalmonClassification interface {
	FishClassification
	GetSalmon() *Salmon
}

type Salmon struct {
	Fish
	Iswild   *bool   `json:"iswild,omitempty"`
	Location *string `json:"location,omitempty"`
}

// GetSalmon implements the SalmonClassification interface for type Salmon.
func (s *Salmon) GetSalmon() *Salmon { return s }

// MarshalJSON implements the json.Marshaller interface for type Salmon.
func (s Salmon) MarshalJSON() ([]byte, error) {
	objectMap := s.Fish.marshalInternal("salmon")
	if s.Iswild != nil {
		objectMap["iswild"] = s.Iswild
	}
	if s.Location != nil {
		objectMap["location"] = s.Location
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Salmon.
func (s *Salmon) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "iswild":
			if val != nil {
				err = json.Unmarshal(*val, &s.Iswild)
			}
		case "location":
			if val != nil {
				err = json.Unmarshal(*val, &s.Location)
			}
		}
		if err != nil {
			return err
		}
	}
	return json.Unmarshal(data, &s.Fish)
}

func (s Salmon) marshalInternal(discValue string) map[string]interface{} {
	objectMap := s.Fish.marshalInternal(discValue)
	if s.Iswild != nil {
		objectMap["iswild"] = s.Iswild
	}
	if s.Location != nil {
		objectMap["location"] = s.Location
	}
	return objectMap
}

// SalmonResponse is the response envelope for operations that return a Salmon type.
type SalmonResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
	Salmon      SalmonClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SalmonResponse.
func (s *SalmonResponse) UnmarshalJSON(data []byte) error {
	t, err := unmarshalSalmonClassification(data)
	if err != nil {
		return err
	}
	s.Salmon = t
	return nil
}

type Sawshark struct {
	Shark
	Picture *[]byte `json:"picture,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type Sawshark.
func (s Sawshark) MarshalJSON() ([]byte, error) {
	objectMap := s.Shark.marshalInternal("sawshark")
	if s.Picture != nil {
		objectMap["picture"] = s.Picture
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Sawshark.
func (s *Sawshark) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "picture":
			if val != nil {
				err = json.Unmarshal(*val, &s.Picture)
			}
		}
		if err != nil {
			return err
		}
	}
	return json.Unmarshal(data, &s.Shark)
}

// SharkClassification provides polymorphic access to related types.
type SharkClassification interface {
	FishClassification
	GetShark() *Shark
}

type Shark struct {
	Fish
	Age      *int32     `json:"age,omitempty"`
	Birthday *time.Time `json:"birthday,omitempty"`
}

// GetShark implements the SharkClassification interface for type Shark.
func (s *Shark) GetShark() *Shark { return s }

// MarshalJSON implements the json.Marshaller interface for type Shark.
func (s Shark) MarshalJSON() ([]byte, error) {
	objectMap := s.Fish.marshalInternal("shark")
	if s.Age != nil {
		objectMap["age"] = s.Age
	}
	if s.Birthday != nil {
		objectMap["birthday"] = (*timeRFC3339)(s.Birthday)
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Shark.
func (s *Shark) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "age":
			if val != nil {
				err = json.Unmarshal(*val, &s.Age)
			}
		case "birthday":
			if val != nil {
				var aux timeRFC3339
				err = json.Unmarshal(*val, &aux)
				s.Birthday = (*time.Time)(&aux)
			}
		}
		if err != nil {
			return err
		}
	}
	return json.Unmarshal(data, &s.Fish)
}

func (s Shark) marshalInternal(discValue string) map[string]interface{} {
	objectMap := s.Fish.marshalInternal(discValue)
	if s.Age != nil {
		objectMap["age"] = s.Age
	}
	if s.Birthday != nil {
		objectMap["birthday"] = (*timeRFC3339)(s.Birthday)
	}
	return objectMap
}

type Siamese struct {
	Cat
	Breed *string `json:"breed,omitempty"`
}

// SiameseResponse is the response envelope for operations that return a Siamese type.
type SiameseResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
	Siamese     *Siamese
}

type SmartSalmon struct {
	Salmon
	// Contains additional key/value pairs not defined in the schema.
	AdditionalProperties *map[string]interface{}
	CollegeDegree        *string `json:"college_degree,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type SmartSalmon.
func (s SmartSalmon) MarshalJSON() ([]byte, error) {
	objectMap := s.Salmon.marshalInternal("smart_salmon")
	if s.CollegeDegree != nil {
		objectMap["college_degree"] = s.CollegeDegree
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SmartSalmon.
func (s *SmartSalmon) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "college_degree":
			if val != nil {
				err = json.Unmarshal(*val, &s.CollegeDegree)
			}
		}
		if err != nil {
			return err
		}
	}
	return json.Unmarshal(data, &s.Salmon)
}

type StringWrapper struct {
	Empty *string `json:"empty,omitempty"`
	Field *string `json:"field,omitempty"`
	Null  *string `json:"null,omitempty"`
}

// StringWrapperResponse is the response envelope for operations that return a StringWrapper type.
type StringWrapperResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse   *http.Response
	StringWrapper *StringWrapper
}
