// Copyright (c) 2023, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package keyfun

import "runtime"

func init() {
	AvailMaps.CopyFrom(StdMaps)
	switch runtime.GOOS {
	case "darwin":
		DefaultMap = "MacStd"
	case "windows":
		DefaultMap = "WindowsStd"
	}
	SetActiveMapName(DefaultMap)
}
