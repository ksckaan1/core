// Copyright 2023 The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Based on golang.org/x/exp/shiny:
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package base

import (
	"image"

	"goki.dev/goosi"
)

// WindowSingle contains the data and logic common to all implementations of [goosi.Window]
// on single-window platforms (mobile, web, and offscreen), as opposed to multi-window
// platforms (desktop), for which you should use [WindowSingle].
// A WindowSingle is associated with a corresponding [goosi.App] type.
// The [goosi.App] type should embed [AppSingle].
type WindowSingle[A goosi.App] struct {
	Window[A]
}

func (w *WindowSingle[A]) GetScreen() *goosi.Screen {
	return w.App.Screen(0)
}

func (w *WindowSingle[A]) Size() image.Point {
	// w.Mu.Lock() // this prevents race conditions but also locks up
	// defer w.Mu.Unlock()
	return w.GetScreen().PixSize
}

func (w *WindowSingle[A]) WinSize() image.Point {
	// w.Mu.Lock() // this prevents race conditions but also locks up
	// defer w.Mu.Unlock()
	return w.GetScreen().PixSize
}

func (w *WindowSingle[A]) Position() image.Point {
	// w.Mu.Lock()
	// defer w.Mu.Unlock()
	return image.Point{}
}

func (w *WindowSingle[A]) PhysicalDPI() float32 {
	w.Mu.Lock()
	defer w.Mu.Unlock()
	return w.GetScreen().PhysicalDPI
}

func (w *WindowSingle[A]) LogicalDPI() float32 {
	w.Mu.Lock()
	defer w.Mu.Unlock()
	return w.GetScreen().LogicalDPI
}

func (w *WindowSingle[A]) SetLogicalDPI(dpi float32) {
	w.Mu.Lock()
	defer w.Mu.Unlock()
	w.GetScreen().LogicalDPI = dpi
}

func (w *WindowSingle[A]) SetWinSize(sz image.Point) {
	if w.This.IsClosed() {
		return
	}
	w.GetScreen().PixSize = sz
}

func (w *WindowSingle[A]) SetSize(sz image.Point) {
	if w.This.IsClosed() {
		return
	}
	w.GetScreen().PixSize = sz
}

func (w *WindowSingle[A]) SetPos(pos image.Point) {
	// no-op
}

func (w *WindowSingle[A]) SetGeom(pos image.Point, sz image.Point) {
	if w.This.IsClosed() {
		return
	}
	w.GetScreen().PixSize = sz
}
