// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v28/common"
)

// LogAnalyticsCollectionWarning Defines the resource kind for collection warning.
type LogAnalyticsCollectionWarning struct {

	// The id of the collection warning
	Id *string `mandatory:"false" json:"id"`
}

func (m LogAnalyticsCollectionWarning) String() string {
	return common.PointerString(m)
}
