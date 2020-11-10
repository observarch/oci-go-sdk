// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataflow

import (
	"github.com/oracle/oci-go-sdk/v28/common"
	"net/http"
)

// ListPrivateEndpointsRequest wrapper for the ListPrivateEndpoints operation
type ListPrivateEndpointsRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique identifier for the request. If provided, the returned request ID will include this value.
	// Otherwise, a random request ID will be generated by the service.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of results to return in a paginated `List` call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` or `opc-prev-page` response header from the last `List` call
	// to sent back to server for getting the next page of results.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The LifecycleState of the private endpoint.
	LifecycleState ListPrivateEndpointsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The field used to sort the results. Multiple fields are not supported.
	SortBy ListPrivateEndpointsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The ordering of results in ascending or descending order.
	SortOrder ListPrivateEndpointsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The query parameter for the Spark application name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID of the user who created the resource.
	OwnerPrincipalId *string `mandatory:"false" contributesTo:"query" name:"ownerPrincipalId"`

	// The displayName prefix.
	DisplayNameStartsWith *string `mandatory:"false" contributesTo:"query" name:"displayNameStartsWith"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPrivateEndpointsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPrivateEndpointsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPrivateEndpointsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListPrivateEndpointsResponse wrapper for the ListPrivateEndpoints operation
type ListPrivateEndpointsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PrivateEndpointCollection instances
	PrivateEndpointCollection `presentIn:"body"`

	// Retrieves the previous page of results.
	// When this header appears in the response, previous pages of results exist.
	// See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Retrieves the next page of results. When this header appears in the response,
	// additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListPrivateEndpointsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPrivateEndpointsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPrivateEndpointsLifecycleStateEnum Enum with underlying type: string
type ListPrivateEndpointsLifecycleStateEnum string

// Set of constants representing the allowable values for ListPrivateEndpointsLifecycleStateEnum
const (
	ListPrivateEndpointsLifecycleStateCreating ListPrivateEndpointsLifecycleStateEnum = "CREATING"
	ListPrivateEndpointsLifecycleStateActive   ListPrivateEndpointsLifecycleStateEnum = "ACTIVE"
	ListPrivateEndpointsLifecycleStateInactive ListPrivateEndpointsLifecycleStateEnum = "INACTIVE"
	ListPrivateEndpointsLifecycleStateUpdating ListPrivateEndpointsLifecycleStateEnum = "UPDATING"
	ListPrivateEndpointsLifecycleStateDeleting ListPrivateEndpointsLifecycleStateEnum = "DELETING"
	ListPrivateEndpointsLifecycleStateDeleted  ListPrivateEndpointsLifecycleStateEnum = "DELETED"
	ListPrivateEndpointsLifecycleStateFailed   ListPrivateEndpointsLifecycleStateEnum = "FAILED"
)

var mappingListPrivateEndpointsLifecycleState = map[string]ListPrivateEndpointsLifecycleStateEnum{
	"CREATING": ListPrivateEndpointsLifecycleStateCreating,
	"ACTIVE":   ListPrivateEndpointsLifecycleStateActive,
	"INACTIVE": ListPrivateEndpointsLifecycleStateInactive,
	"UPDATING": ListPrivateEndpointsLifecycleStateUpdating,
	"DELETING": ListPrivateEndpointsLifecycleStateDeleting,
	"DELETED":  ListPrivateEndpointsLifecycleStateDeleted,
	"FAILED":   ListPrivateEndpointsLifecycleStateFailed,
}

// GetListPrivateEndpointsLifecycleStateEnumValues Enumerates the set of values for ListPrivateEndpointsLifecycleStateEnum
func GetListPrivateEndpointsLifecycleStateEnumValues() []ListPrivateEndpointsLifecycleStateEnum {
	values := make([]ListPrivateEndpointsLifecycleStateEnum, 0)
	for _, v := range mappingListPrivateEndpointsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListPrivateEndpointsSortByEnum Enum with underlying type: string
type ListPrivateEndpointsSortByEnum string

// Set of constants representing the allowable values for ListPrivateEndpointsSortByEnum
const (
	ListPrivateEndpointsSortByTimecreated ListPrivateEndpointsSortByEnum = "timeCreated"
)

var mappingListPrivateEndpointsSortBy = map[string]ListPrivateEndpointsSortByEnum{
	"timeCreated": ListPrivateEndpointsSortByTimecreated,
}

// GetListPrivateEndpointsSortByEnumValues Enumerates the set of values for ListPrivateEndpointsSortByEnum
func GetListPrivateEndpointsSortByEnumValues() []ListPrivateEndpointsSortByEnum {
	values := make([]ListPrivateEndpointsSortByEnum, 0)
	for _, v := range mappingListPrivateEndpointsSortBy {
		values = append(values, v)
	}
	return values
}

// ListPrivateEndpointsSortOrderEnum Enum with underlying type: string
type ListPrivateEndpointsSortOrderEnum string

// Set of constants representing the allowable values for ListPrivateEndpointsSortOrderEnum
const (
	ListPrivateEndpointsSortOrderAsc  ListPrivateEndpointsSortOrderEnum = "ASC"
	ListPrivateEndpointsSortOrderDesc ListPrivateEndpointsSortOrderEnum = "DESC"
)

var mappingListPrivateEndpointsSortOrder = map[string]ListPrivateEndpointsSortOrderEnum{
	"ASC":  ListPrivateEndpointsSortOrderAsc,
	"DESC": ListPrivateEndpointsSortOrderDesc,
}

// GetListPrivateEndpointsSortOrderEnumValues Enumerates the set of values for ListPrivateEndpointsSortOrderEnum
func GetListPrivateEndpointsSortOrderEnumValues() []ListPrivateEndpointsSortOrderEnum {
	values := make([]ListPrivateEndpointsSortOrderEnum, 0)
	for _, v := range mappingListPrivateEndpointsSortOrder {
		values = append(values, v)
	}
	return values
}
