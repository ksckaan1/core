// Copyright (c) 2023, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gi

import (
	"goki.dev/colors"
	"goki.dev/cursors"
	"goki.dev/girl/abilities"
	"goki.dev/girl/states"
	"goki.dev/girl/styles"
	"goki.dev/girl/units"
	"goki.dev/goosi/events"
	"goki.dev/icons"
	"goki.dev/ki/v2"
)

// Switch is a widget that can toggle between an on and off state.
// It can be displayed as a switch, checkbox, or radio button.
type Switch struct {
	WidgetBase

	// the type of switch that this is
	Type SwitchTypes `set:"-"`

	// the label text for the switch
	Text string

	// icon to use for the on, checked state of the switch
	IconOn icons.Icon `view:"show-name"`

	// icon to use for the off, unchecked state of the switch
	IconOff icons.Icon `view:"show-name"`

	// icon to use for the disabled state of the switch
	IconDisab icons.Icon `view:"show-name"`
}

// SwitchTypes contains the different types of [Switch]es
type SwitchTypes int32 //enums:enum -trimprefix Switch

const (
	// SwitchSwitch indicates to display a switch as a switch (toggle slider)
	SwitchSwitch SwitchTypes = iota
	// SwitchCheckbox indicates to display a switch as a checkbox
	SwitchCheckbox
	// SwitchRadioButton indicates to display a switch as a radio button
	SwitchRadioButton
)

func (sw *Switch) CopyFieldsFrom(frm any) {
	fr := frm.(*Switch)
	sw.WidgetBase.CopyFieldsFrom(&fr.WidgetBase)
	sw.Type = fr.Type
	sw.Text = fr.Text
	sw.IconOn = fr.IconOn
	sw.IconOff = fr.IconOff
	sw.IconDisab = fr.IconDisab
}

func (sw *Switch) OnInit() {
	sw.IconDisab = icons.Blank
	sw.HandleSwitchEvents()
	sw.SwitchStyles()
}

func (sw *Switch) SetIconFromState() {
	if sw.Parts == nil {
		return
	}
	ist := sw.Parts.ChildByName("stack", 0)
	if ist == nil {
		return
	}
	st := ist.(*Layout)
	switch {
	case sw.IsReadOnly():
		st.StackTop = 2
	case sw.StateIs(states.Checked):
		st.StackTop = 0
	default:
		st.StackTop = 1
	}
}

func (sw *Switch) HandleSwitchEvents() {
	sw.HandleWidgetEvents()
	sw.HandleSelectToggle()
	sw.HandleClickOnEnterSpace()
	sw.OnClick(func(e events.Event) {
		e.SetHandled()
		sw.SetState(!sw.StateIs(states.Checked), states.Checked)
		if sw.Parts != nil && sw.Parts.HasChildren() {
			sw.SetIconFromState()
		}
		sw.Send(events.Change, e)
	})
}

func (sw *Switch) SwitchStyles() {
	sw.Style(func(s *styles.Style) {
		s.SetAbilities(true, abilities.Activatable, abilities.Focusable, abilities.Hoverable, abilities.Checkable)
		if !sw.IsReadOnly() {
			s.Cursor = cursors.Pointer
		}
		s.Text.Align = styles.AlignLeft
		s.Color = colors.Scheme.OnBackground
		s.Margin.Set(units.Dp(1))
		s.Padding.Set(units.Dp(1))
		s.Border.Radius = styles.BorderRadiusExtraSmall

		if s.Is(states.Selected) {
			s.BackgroundColor.SetSolid(colors.Scheme.Select.Container)
		}
	})
	sw.OnWidgetAdded(func(w Widget) {
		switch w.PathFrom(sw.This()) {
		case "parts/stack/icon0": // on
			w.Style(func(s *styles.Style) {
				s.Color = colors.Scheme.Primary.Base
				// switches need to be bigger
				if sw.Type == SwitchSwitch {
					s.Width.SetEm(3)
					s.Height.SetEm(3)
				} else {
					s.Width.SetEm(1.5)
					s.Height.SetEm(1.5)
				}
			})
		case "parts/stack/icon1": // off
			w.Style(func(s *styles.Style) {
				// switches need to be bigger
				if sw.Type == SwitchSwitch {
					s.Width.SetEm(3)
					s.Height.SetEm(3)
				} else {
					s.Width.SetEm(1.5)
					s.Height.SetEm(1.5)
				}
			})
		// todo: disabled?
		case "parts/space":
			w.Style(func(s *styles.Style) {
				s.Width.SetCh(0.1)
			})
		case "parts/label":
			w.Style(func(s *styles.Style) {
				s.SetAbilities(false, abilities.Selectable, abilities.DoubleClickable)
				s.Cursor = cursors.None
				s.Margin.Set()
				s.Padding.Set()
				s.AlignV = styles.AlignMiddle
			})
		}
	})
}

// SetType sets the styling type of the switch
func (sw *Switch) SetType(typ SwitchTypes) *Switch {
	updt := sw.UpdateStart()
	sw.Type = typ
	sw.IconDisab = icons.Blank
	switch sw.Type {
	case SwitchSwitch:
		// TODO: material has more advanced switches with a checkmark
		// if they are turned on; we could implement that at some point
		sw.IconOn = icons.ToggleOn.Fill()
		sw.IconOff = icons.ToggleOff
	case SwitchCheckbox:
		sw.IconOn = icons.CheckBox.Fill()
		sw.IconOff = icons.CheckBoxOutlineBlank
	case SwitchRadioButton:
		sw.IconOn = icons.RadioButtonChecked
		sw.IconOff = icons.RadioButtonUnchecked
	}
	sw.UpdateEndLayout(updt)
	return sw
}

// LabelWidget returns the label widget if present
func (sw *Switch) LabelWidget() *Label {
	lbi := sw.Parts.ChildByName("label")
	if lbi == nil {
		return nil
	}
	return lbi.(*Label)
}

// SetIcons sets the icons for the on (checked)
// and off (unchecked) states, and updates the switch
func (sw *Switch) SetIcons(on, off icons.Icon) *Switch {
	updt := sw.UpdateStart()
	sw.IconOn = on
	sw.IconOff = off
	sw.UpdateEndLayout(updt)
	return sw
}

func (sw *Switch) ConfigWidget(sc *Scene) {
	sw.ConfigParts(sc)
}

func (sw *Switch) ConfigParts(sc *Scene) {
	parts := sw.NewParts(LayoutHoriz)
	if !sw.IconOn.IsValid() {
		sw.IconOn = icons.ToggleOn.Fill() // fallback
	}
	if !sw.IconOff.IsValid() {
		sw.IconOff = icons.ToggleOff // fallback
	}
	config := ki.Config{}
	icIdx := 0 // always there
	lbIdx := -1
	config.Add(LayoutType, "stack")
	if sw.Text != "" {
		config.Add(SpaceType, "space")
		lbIdx = len(config)
		config.Add(LabelType, "label")
	}
	mods, updt := parts.ConfigChildren(config)
	ist := parts.Child(icIdx).(*Layout)
	if mods || sw.NeedsRebuild() {
		ist.Lay = LayoutStacked
		ist.SetNChildren(3, IconType, "icon")
		icon := ist.Child(0).(*Icon)
		icon.SetIcon(sw.IconOn)
		icoff := ist.Child(1).(*Icon)
		icoff.SetIcon(sw.IconOff)
		icdsb := ist.Child(2).(*Icon)
		icdsb.SetIcon(sw.IconDisab)
	}
	sw.SetIconFromState()
	if lbIdx >= 0 {
		lbl := parts.Child(lbIdx).(*Label)
		if lbl.Text != sw.Text {
			lbl.SetText(sw.Text)
		}
	}
	if mods {
		parts.Update()
		parts.UpdateEnd(updt)
		sw.SetNeedsLayoutUpdate(sc, updt)
	}
}

func (sw *Switch) RenderSwitch(sc *Scene) {
	rs, _, st := sw.RenderLock(sc)
	sw.RenderStdBox(sc, st)
	sw.RenderUnlock(rs)
}

func (sw *Switch) Render(sc *Scene) {
	sw.SetIconFromState() // make sure we're always up-to-date on render
	if sw.PushBounds(sc) {
		sw.RenderSwitch(sc)
		sw.RenderParts(sc)
		sw.RenderChildren(sc)
		sw.PopBounds(sc)
	}
}
