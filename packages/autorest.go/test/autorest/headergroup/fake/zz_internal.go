// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"net/http"
	"reflect"
)

type nonRetriableError struct {
	error
}

func (nonRetriableError) NonRetriable() {
	// marker method
}

func contains[T comparable](s []T, v T) bool {
	for _, vv := range s {
		if vv == v {
			return true
		}
	}
	return false
}

func getHeaderValue(h http.Header, k string) string {
	v := h[k]
	if len(v) == 0 {
		return ""
	}
	return v[0]
}

func getOptional[T any](v T) *T {
	if reflect.ValueOf(v).IsZero() {
		return nil
	}
	return &v
}

func parseOptional[T any](v string, parse func(v string) (T, error)) (*T, error) {
	if v == "" {
		return nil, nil
	}
	t, err := parse(v)
	if err != nil {
		return nil, err
	}
	return &t, err
}

func parseWithCast[T any](v string, parse func(v string) (T, error)) (T, error) {
	t, err := parse(v)
	if err != nil {
		return *new(T), err
	}
	return t, err
}
