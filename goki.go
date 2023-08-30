// Copyright (c) 2023, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"goki.dev/goki/cmd"
	"goki.dev/grease"
)

func main() {
	grease.AppName = "goki"
	grease.AppAbout = "Command line and GUI tools for developing apps and libraries using the GoKi framework."
	err := grease.Run(cmd.TheApp, ".goki/config.toml")
	if err != nil {
		fmt.Println(err)
	}
}
