// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azurespecialsgroup

import (
	"context"
	"generatortests/helpers"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
)

func newOdataClient() OdataOperations {
	return NewOdataClient(NewDefaultConnection(nil))
}

// GetWithFilter - Specify filter parameter with value '$filter=id gt 5 and name eq 'foo'&$orderby=id&$top=10'
func TestGetWithFilter(t *testing.T) {
	client := newOdataClient()
	result, err := client.GetWithFilter(context.Background(), &OdataGetWithFilterOptions{
		Filter:  to.StringPtr("id gt 5 and name eq 'foo'"),
		Orderby: to.StringPtr("id"),
		Top:     to.Int32Ptr(10),
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}
