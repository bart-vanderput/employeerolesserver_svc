// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package main

import (
	"github.com/bart-vanderput/employeerolesserver_svc/app"
	"github.com/pkg/errors"
)

// This is the name you will use for the NET START command
const svcName = "IamRolesWebServer"

// This is the name that will appear in the Services control panel
const svcNameLong = "IAM Employee Roles Web Server - Custom"

func svcLauncher() error {

	err := app.Run(elog, svcName)
	if err != nil {
		return errors.Wrap(err, "app.run")
	}

	return nil
}
