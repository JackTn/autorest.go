// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package numbergroup

import (
	"context"
	"generatortests/helpers"
	"net/http"
	"testing"
)

func newNumberClient() NumberOperations {
	return NewNumberClient(NewDefaultConnection(nil))
}

func TestNumberGetBigDecimal(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetBigDecimal(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetBigDecimal: %v", err)
	}
	val := 2.5976931e+101
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	helpers.DeepEqualOrFatal(t, result.Value, &val)
}

func TestNumberGetBigDecimalNegativeDecimal(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetBigDecimalNegativeDecimal(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetBigDecimalNegativeDecimal: %v", err)
	}
	val := -99999999.99
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	helpers.DeepEqualOrFatal(t, result.Value, &val)
}

func TestNumberGetBigDecimalPositiveDecimal(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetBigDecimalPositiveDecimal(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetBigDecimalPositiveDecimal: %v", err)
	}
	val := 99999999.99
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	helpers.DeepEqualOrFatal(t, result.Value, &val)
}

func TestNumberGetBigDouble(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetBigDouble(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetBigDouble: %v", err)
	}
	val := 2.5976931e+101
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	helpers.DeepEqualOrFatal(t, result.Value, &val)
}

func TestNumberGetBigDoubleNegativeDecimal(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetBigDoubleNegativeDecimal(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetBigDoubleNegativeDecimal: %v", err)
	}
	val := -99999999.99
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	helpers.DeepEqualOrFatal(t, result.Value, &val)
}

func TestNumberGetBigDoublePositiveDecimal(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetBigDoublePositiveDecimal(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetBigDoublePositiveDecimal: %v", err)
	}
	val := 99999999.99
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	helpers.DeepEqualOrFatal(t, result.Value, &val)
}

func TestNumberGetBigFloat(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetBigFloat(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetBigFloat: %v", err)
	}
	val := float32(3.402823e+20)
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	helpers.DeepEqualOrFatal(t, result.Value, &val)
}

func TestNumberGetInvalidDecimal(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetInvalidDecimal(context.Background(), nil)
	if err == nil {
		t.Fatalf("unexpected nil error")
	}
	if result != nil {
		t.Fatalf("expected a nil result")
	}
}

func TestNumberGetInvalidDouble(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetInvalidDouble(context.Background(), nil)
	if err == nil {
		t.Fatalf("unexpected nil error")
	}
	if result != nil {
		t.Fatalf("expected a nil result")
	}
}

func TestNumberGetInvalidFloat(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetInvalidFloat(context.Background(), nil)
	if err == nil {
		t.Fatalf("unexpected nil error")
	}
	if result != nil {
		t.Fatalf("expected a nil result")
	}
}

func TestNumberGetNull(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetNull(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetNull: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	helpers.DeepEqualOrFatal(t, result.Value, (*float32)(nil))
}

func TestNumberGetSmallDecimal(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetSmallDecimal(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetSmallDecimal: %v", err)
	}
	val := 2.5976931e-101
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	helpers.DeepEqualOrFatal(t, result.Value, &val)
}

func TestNumberGetSmallDouble(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetSmallDouble(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetSmallDouble: %v", err)
	}
	val := 2.5976931e-101
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	helpers.DeepEqualOrFatal(t, result.Value, &val)
}

func TestNumberGetSmallFloat(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetSmallFloat(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetSmallFloat: %v", err)
	}
	val := 3.402823e-20
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	helpers.DeepEqualOrFatal(t, result.Value, &val)
}

func TestNumberPutBigDecimal(t *testing.T) {
	client := newNumberClient()
	result, err := client.PutBigDecimal(context.Background(), 2.5976931e+101, nil)
	if err != nil {
		t.Fatalf("PutBigDecimal: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestNumberPutBigDecimalNegativeDecimal(t *testing.T) {
	client := newNumberClient()
	result, err := client.PutBigDecimalNegativeDecimal(context.Background(), nil)
	if err != nil {
		t.Fatalf("PutBigDecimalNegativeDecimal: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestNumberPutBigDecimalPositiveDecimal(t *testing.T) {
	client := newNumberClient()
	result, err := client.PutBigDecimalPositiveDecimal(context.Background(), nil)
	if err != nil {
		t.Fatalf("PutBigDecimalPositiveDecimal: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestNumberPutBigDouble(t *testing.T) {
	client := newNumberClient()
	result, err := client.PutBigDouble(context.Background(), 2.5976931e+101, nil)
	if err != nil {
		t.Fatalf("PutBigDouble: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestNumberPutBigDoubleNegativeDecimal(t *testing.T) {
	client := newNumberClient()
	result, err := client.PutBigDoubleNegativeDecimal(context.Background(), nil)
	if err != nil {
		t.Fatalf("PutBigDoubleNegativeDecimal: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestNumberPutBigDoublePositiveDecimal(t *testing.T) {
	client := newNumberClient()
	result, err := client.PutBigDoublePositiveDecimal(context.Background(), nil)
	if err != nil {
		t.Fatalf("PutBigDeoublePositiveDecimal: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestNumberPutBigFloat(t *testing.T) {
	client := newNumberClient()
	result, err := client.PutBigFloat(context.Background(), 3.402823e+20, nil)
	if err != nil {
		t.Fatalf("PutBigFloat: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestNumberPutSmallDecimal(t *testing.T) {
	client := newNumberClient()
	result, err := client.PutSmallDecimal(context.Background(), 2.5976931e-101, nil)
	if err != nil {
		t.Fatalf("PutSmallDecimal: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestNumberPutSmallDouble(t *testing.T) {
	client := newNumberClient()
	result, err := client.PutSmallDouble(context.Background(), 2.5976931e-101, nil)
	if err != nil {
		t.Fatalf("PutSmallDouble: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestNumberPutSmallFloat(t *testing.T) {
	client := newNumberClient()
	result, err := client.PutSmallFloat(context.Background(), 3.402823e-20, nil)
	if err != nil {
		t.Fatalf("PutSmallFloat: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}
