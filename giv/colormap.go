// Copyright (c) 2019, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package giv

import (
	"log/slog"

	"cogentcore.org/core/abilities"
	"cogentcore.org/core/colors"
	"cogentcore.org/core/colors/colormap"
	"cogentcore.org/core/colors/gradient"
	"cogentcore.org/core/cursors"
	"cogentcore.org/core/gi"
	"cogentcore.org/core/gti"
	"cogentcore.org/core/laser"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/units"
)

// ColorMapName represents the name of a color map, which can be edited using a [ColorMapValue].
type ColorMapName string

func (cmn ColorMapName) Value() Value {
	return &ColorMapValue{}
}

// ColorMapValue displays a color map spectrum and can be clicked on
// to display a dialog for selecting different color map options.
// It represents a [ColorMapName] value.
type ColorMapValue struct {
	ValueBase
}

func (vv *ColorMapValue) WidgetType() *gti.Type {
	vv.WidgetTyp = gi.FrameType
	return vv.WidgetTyp
}

func (vv *ColorMapValue) UpdateWidget() {
	if vv.Widget == nil {
		return
	}
	vv.Widget.ApplyStyle()
	vv.AsWidgetBase().NeedsRender()
}

func (vv *ColorMapValue) ConfigWidget(w gi.Widget) {
	if vv.Widget == w {
		vv.UpdateWidget()
		return
	}
	vv.Widget = w
	vv.StdConfigWidget(w)
	fr := vv.Widget.(*gi.Frame)
	fr.HandleClickOnEnterSpace()
	ConfigDialogWidget(vv, fr, false)
	fr.Style(func(s *styles.Style) {
		s.SetAbilities(true, abilities.Hoverable, abilities.Clickable, abilities.Focusable)
		s.Cursor = cursors.Pointer
		s.Border.Radius = styles.BorderRadiusMedium

		s.Grow.Set(0, 0)
		s.Min.Set(units.Em(10), units.Em(1.5))

		cmn, ok := laser.NonPtrValue(vv.Value).Interface().(ColorMapName)
		if !ok || cmn == "" {
			s.Background = colors.C(colors.Scheme.OutlineVariant)
			return
		}
		cm, ok := colormap.AvailMaps[string(cmn)]
		if !ok {
			slog.Error("got invalid color map name", "name", cmn)
			s.Background = colors.C(colors.Scheme.OutlineVariant)
			return
		}
		g := gradient.NewLinear()
		for i := float32(0); i < 1; i += 0.01 {
			gc := cm.Map(i)
			g.AddStop(gc, i)
		}
		s.Background = g
	})
	vv.UpdateWidget()
}

func (vv *ColorMapValue) HasDialog() bool { return true }
func (vv *ColorMapValue) OpenDialog(ctx gi.Widget, fun func()) {
	OpenValueDialog(vv, ctx, fun, "Select a color map")
}

func (vv *ColorMapValue) ConfigDialog(d *gi.Body) (bool, func()) {
	sl := colormap.AvailMapsList()
	cur := laser.ToString(vv.Value.Interface())
	si := 0
	NewSliceView(d).SetSlice(&sl).SetSelVal(cur).BindSelect(&si)
	return true, func() {
		if si >= 0 {
			vv.SetValue(sl[si])
			vv.UpdateWidget()
		}
	}
}
