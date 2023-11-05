// Copyright (c) 2023, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goosi

// Drawer is an interface representing a type capable of high-performance
// rendering to a window surface. It is implemented by [*goki.dev/vgpu/v2/vdraw.Drawer]
// and an internal web driver.
type Drawer interface {
	// SetMaxTextures updates the max number of textures for drawing
	// Must call this prior to doing any allocation of images.
	SetMaxTextures(maxTextures int)
	// MaxTextures returns the max number of textures for drawing
	MaxTextures() int
}
