// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
)

// InstanceCredentials. The credentials for a particular instance.
type InstanceCredentials struct {

	// The password for the username.
	Password *string `mandatory:"true" json:"password,omitempty"`

	// The username.
	Username *string `mandatory:"true" json:"username,omitempty"`
}

func (model InstanceCredentials) String() string {
	return common.PointerString(model)
}