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

// UpdateLogAnalyticsEntityTypeDetails Log analytics entity type definition to be updated.
type UpdateLogAnalyticsEntityTypeDetails struct {

	// Log analytics entity type category. Category will be used for grouping and filtering.
	Category *string `mandatory:"false" json:"category"`

	// A single log analytics entity type property definition.
	Properties []EntityTypeProperty `mandatory:"false" json:"properties"`
}

func (m UpdateLogAnalyticsEntityTypeDetails) String() string {
	return common.PointerString(m)
}
