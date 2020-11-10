// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/v28/common"
	"net/http"
)

// GetCloudVmClusterUpdateHistoryEntryRequest wrapper for the GetCloudVmClusterUpdateHistoryEntry operation
type GetCloudVmClusterUpdateHistoryEntryRequest struct {

	// The cloud VM cluster OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CloudVmClusterId *string `mandatory:"true" contributesTo:"path" name:"cloudVmClusterId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the maintenance update history entry.
	UpdateHistoryEntryId *string `mandatory:"true" contributesTo:"path" name:"updateHistoryEntryId"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetCloudVmClusterUpdateHistoryEntryRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetCloudVmClusterUpdateHistoryEntryRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetCloudVmClusterUpdateHistoryEntryRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetCloudVmClusterUpdateHistoryEntryResponse wrapper for the GetCloudVmClusterUpdateHistoryEntry operation
type GetCloudVmClusterUpdateHistoryEntryResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The UpdateHistoryEntry instance
	UpdateHistoryEntry `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetCloudVmClusterUpdateHistoryEntryResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetCloudVmClusterUpdateHistoryEntryResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
