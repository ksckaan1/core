// Copyright (c) 2018, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package giv

import (
	"fmt"
	"log/slog"
	"reflect"
	"strings"

	"cogentcore.org/core/enums"
	"cogentcore.org/core/events"
	"cogentcore.org/core/gi"
	"cogentcore.org/core/grr"
	"cogentcore.org/core/gti"
	"cogentcore.org/core/icons"
	"cogentcore.org/core/ki"
	"cogentcore.org/core/laser"
	"cogentcore.org/core/paint"
	"cogentcore.org/core/strcase"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/units"
)

// This file contains the standard [Value]s built into giv.

// StringValue represents any value with a text field.
type StringValue struct {
	ValueBase[*gi.TextField]
}

func (v *StringValue) Config() {
	if vtag, _ := v.Tag("view"); vtag == "password" {
		v.Widget.SetTypePassword()
	}
	if vl, ok := v.Value.Interface().(gi.Validator); ok {
		v.Widget.SetValidator(vl.Validate)
	}
	if fv, ok := v.Owner.(gi.FieldValidator); ok {
		v.Widget.SetValidator(func() error {
			return fv.ValidateField(v.Field.Name)
		})
	}

	v.Widget.OnFinal(events.Change, func(e events.Event) {
		if v.SetValue(v.Widget.Text()) {
			v.Update()
		}
	})
}

func (v *StringValue) Update() {
	npv := laser.NonPtrValue(v.Value)
	if npv.Kind() == reflect.Interface && npv.IsZero() {
		v.Widget.SetText("None")
	} else {
		txt := laser.ToString(v.Value.Interface())
		v.Widget.SetText(txt)
	}
}

// BoolValue represents a bool value with a switch.
type BoolValue struct {
	ValueBase[*gi.Switch]
}

func (v *BoolValue) Config() {
	v.Widget.OnChange(func(e events.Event) {
		v.SetValue(v.Widget.IsChecked())
	})
}

func (v *BoolValue) Update() {
	npv := laser.NonPtrValue(v.Value)
	bv, err := laser.ToBool(npv.Interface())
	if grr.Log(err) == nil {
		v.Widget.SetChecked(bv)
	}
}

// NumberValue represents an integer or float value with a spinner.
type NumberValue struct {
	ValueBase[*gi.Spinner]
}

func (v *NumberValue) Config() {
	vk := laser.NonPtrType(v.Value.Type()).Kind()
	if vk >= reflect.Int && vk <= reflect.Uintptr {
		v.Widget.SetStep(1).SetPageStep(10)
	}
	if vk >= reflect.Uint && vk <= reflect.Uintptr {
		v.Widget.SetMin(0)
	}
	if min, ok := v.Tag("min"); ok {
		minv, err := laser.ToFloat32(min)
		if grr.Log(err) == nil {
			v.Widget.SetMin(minv)
		}
	}
	if max, ok := v.Tag("max"); ok {
		maxv, err := laser.ToFloat32(max)
		if grr.Log(err) == nil {
			v.Widget.SetMax(maxv)
		}
	}
	if step, ok := v.Tag("step"); ok {
		step, err := laser.ToFloat32(step)
		if grr.Log(err) == nil {
			v.Widget.SetStep(step)
		}
	}
	if format, ok := v.Tag("format"); ok {
		v.Widget.SetFormat(format)
	}
	v.Widget.OnChange(func(e events.Event) {
		v.SetValue(v.Widget.Value)
	})
}

func (v *NumberValue) Update() {
	npv := laser.NonPtrValue(v.Value)
	fv, err := laser.ToFloat32(npv.Interface())
	if grr.Log(err) == nil {
		v.Widget.SetValue(fv)
	}
}

// SliderValue represents an integer or float value with a slider.
type SliderValue struct {
	ValueBase[*gi.Slider]
}

func (v *SliderValue) Config() {
	if min, ok := v.Tag("min"); ok {
		minv, err := laser.ToFloat32(min)
		if grr.Log(err) == nil {
			v.Widget.SetMin(minv)
		}
	}
	if max, ok := v.Tag("max"); ok {
		maxv, err := laser.ToFloat32(max)
		if grr.Log(err) == nil {
			v.Widget.SetMax(maxv)
		}
	}
	if step, ok := v.Tag("step"); ok {
		stepv, err := laser.ToFloat32(step)
		if grr.Log(err) == nil {
			v.Widget.SetStep(stepv)
		}
	}
	v.Widget.OnChange(func(e events.Event) {
		v.SetValue(v.Widget.Value)
	})
}

func (v *SliderValue) Update() {
	npv := laser.NonPtrValue(v.Value)
	fv, err := laser.ToFloat32(npv.Interface())
	if grr.Log(err) == nil {
		v.Widget.SetValue(fv)
	}
}

// StructValue represents a struct value with a button.
type StructValue struct {
	ValueBase[*gi.Button]
}

func (v *StructValue) Config() {
	v.Widget.SetType(gi.ButtonTonal).SetIcon(icons.Edit)
	ConfigDialogWidget(v, true)
}

func (v *StructValue) Update() {
	npv := laser.NonPtrValue(v.Value)
	if v.Value.IsZero() || npv.IsZero() {
		v.Widget.SetText("None")
	} else {
		opv := laser.OnePtrUnderlyingValue(v.Value)
		if lbler, ok := opv.Interface().(gi.Labeler); ok {
			v.Widget.SetText(lbler.Label())
		} else {
			v.Widget.SetText(laser.FriendlyTypeName(npv.Type()))
		}
	}
	v.Widget.Update()
}

func (v *StructValue) ConfigDialog(d *gi.Body) (bool, func()) {
	if v.Value.IsZero() || laser.NonPtrValue(v.Value).IsZero() {
		return false, nil
	}
	opv := laser.OnePtrUnderlyingValue(v.Value)
	str := opv.Interface()
	NewStructView(d).SetStruct(str).SetViewPath(v.ViewPath).SetTmpSave(v.TmpSave).
		SetReadOnly(v.IsReadOnly())
	if tb, ok := str.(gi.Toolbarer); ok {
		d.AddAppBar(tb.ConfigToolbar)
	}
	return true, nil
}

// StructInlineValue represents a struct value with a [StructViewInline].
type StructInlineValue struct {
	ValueBase[*StructViewInline]
}

func (v *StructInlineValue) Config() {
	v.Widget.StructValue = v
	v.Widget.ViewPath = v.ViewPath
	v.Widget.TmpSave = v.TmpSave
	v.Widget.SetStruct(v.Value.Interface())
	v.Widget.OnChange(func(e events.Event) {
		v.SendChange(e)
	})
}

func (v *StructInlineValue) Update() {
	v.Widget.SetStruct(v.Value.Interface())
}

// SliceValue represents a slice or array value with a button.
type SliceValue struct {
	ValueBase[*gi.Button]
}

func (v *SliceValue) Config() {
	v.Widget.SetType(gi.ButtonTonal).SetIcon(icons.Edit)
	ConfigDialogWidget(v, true)
}

func (v *SliceValue) Update() {
	npv := laser.OnePtrUnderlyingValue(v.Value).Elem()
	txt := ""
	if !npv.IsValid() {
		txt = "None"
	} else {
		if npv.Kind() == reflect.Array || !npv.IsNil() {
			bnm := laser.FriendlyTypeName(laser.SliceElType(v.Value.Interface()))
			if strings.HasSuffix(bnm, "s") {
				txt = strcase.ToSentence(fmt.Sprintf("%d lists of %s", npv.Len(), bnm))
			} else {
				txt = strcase.ToSentence(fmt.Sprintf("%d %ss", npv.Len(), bnm))
			}
		} else {
			txt = "None"
		}
	}
	v.Widget.SetText(txt).Update()
}

func (v *SliceValue) ConfigDialog(d *gi.Body) (bool, func()) {
	npv := laser.NonPtrValue(v.Value)
	if v.Value.IsZero() || npv.IsZero() {
		return false, nil
	}
	vvp := laser.OnePtrValue(v.Value)
	if vvp.Kind() != reflect.Ptr {
		slog.Error("giv.SliceValue: Cannot view unadressable (non-pointer) slices", "type", v.Value.Type())
		return false, nil
	}
	slci := vvp.Interface()
	if npv.Kind() != reflect.Array && laser.NonPtrType(laser.SliceElType(v.Value.Interface())).Kind() == reflect.Struct {
		tv := NewTableView(d).SetSlice(slci).SetTmpSave(v.TmpSave).SetViewPath(v.ViewPath)
		tv.SetReadOnly(v.IsReadOnly())
		d.AddAppBar(tv.ConfigToolbar)
	} else {
		sv := NewSliceView(d).SetSlice(slci).SetTmpSave(v.TmpSave).SetViewPath(v.ViewPath)
		sv.SetReadOnly(v.IsReadOnly())
		d.AddAppBar(sv.ConfigToolbar)
	}
	return true, nil
}

// SliceInlineValue represents a slice or array value with a [SliceViewInline].
type SliceInlineValue struct {
	ValueBase[*SliceViewInline]
}

func (v *SliceInlineValue) Config() {
	v.Widget.SliceValue = v
	v.Widget.ViewPath = v.ViewPath
	v.Widget.TmpSave = v.TmpSave
	v.Widget.SetSlice(v.Value.Interface())
	v.Widget.OnChange(func(e events.Event) {
		v.SendChange(e)
	})
}

func (v *SliceInlineValue) Update() {
	csl := v.Value.Interface()
	newslc := false
	if reflect.TypeOf(v.Value).Kind() != reflect.Pointer { // prevent crash on non-comparable
		newslc = true
	} else {
		newslc = v.Widget.Slice != csl
	}
	if newslc {
		v.Widget.SetSlice(csl)
	} else {
		v.Widget.Update()
	}
}

// MapValue represents a map value with a button.
type MapValue struct {
	ValueBase[*gi.Button]
}

func (v *MapValue) Config() {
	v.Widget.SetType(gi.ButtonTonal).SetIcon(icons.Edit)
	ConfigDialogWidget(v, true)
}

func (v *MapValue) Update() {
	npv := laser.NonPtrValue(v.Value)
	mpi := v.Value.Interface()
	txt := ""
	if !npv.IsValid() || npv.IsNil() {
		txt = "None"
	} else {
		bnm := laser.FriendlyTypeName(laser.MapValueType(mpi))
		if strings.HasSuffix(bnm, "s") {
			txt = strcase.ToSentence(fmt.Sprintf("%d lists of %s", npv.Len(), bnm))
		} else {
			txt = strcase.ToSentence(fmt.Sprintf("%d %ss", npv.Len(), bnm))
		}
	}
	v.Widget.SetText(txt).Update()
}

func (v *MapValue) ConfigDialog(d *gi.Body) (bool, func()) {
	if v.Value.IsZero() || laser.NonPtrValue(v.Value).IsZero() {
		return false, nil
	}
	mpi := v.Value.Interface()
	mv := NewMapView(d).SetMap(mpi)
	mv.SetViewPath(v.ViewPath).SetTmpSave(v.TmpSave).SetReadOnly(v.IsReadOnly())
	d.AddAppBar(mv.ConfigToolbar)
	return true, nil
}

// MapInlineValue represents a map value with a [MapViewInline].
type MapInlineValue struct {
	ValueBase[*MapViewInline]
}

func (v *MapInlineValue) Config() {
	v.Widget.MapValue = v
	v.Widget.ViewPath = v.ViewPath
	v.Widget.TmpSave = v.TmpSave
	v.Widget.SetMap(v.Value.Interface())
	v.Widget.OnChange(func(e events.Event) {
		v.SendChange(e)
	})
}

func (v *MapInlineValue) Update() {
	cmp := v.Value.Interface()
	if v.Widget.Map != cmp {
		v.Widget.SetMap(cmp)
	} else {
		v.Widget.UpdateValues()
	}
}

// KiValue represents a [ki.Ki] value with a button.
type KiValue struct {
	ValueBase[*gi.Button]
}

func (v *KiValue) Config() {
	v.Widget.SetType(gi.ButtonTonal).SetIcon(icons.Edit)
	ConfigDialogWidget(v, true)
}

func (v *KiValue) Update() {
	path := "None"
	k := v.KiValue()
	if k != nil && k.This() != nil {
		path = k.AsKi().String()
	}
	v.Widget.SetText(path).Update()
}

func (v *KiValue) ConfigDialog(d *gi.Body) (bool, func()) {
	k := v.KiValue()
	if k == nil {
		return false, nil
	}
	InspectorView(d, k)
	return true, nil
}

// KiValue returns the actual underlying [ki.Ki] value, or nil.
func (vv *KiValue) KiValue() ki.Ki {
	if !vv.Value.IsValid() || vv.Value.IsNil() {
		return nil
	}
	opv := laser.OnePtrValue(vv.Value)
	if opv.IsNil() {
		return nil
	}
	k, _ := opv.Interface().(ki.Ki)
	return k
}

// EnumValue represents an [enums.Enum] value with a chooser.
type EnumValue struct {
	ValueBase[*gi.Chooser]
}

func (v *EnumValue) Config() {
	e, _ := laser.OnePtrUnderlyingValue(v.Value).Interface().(enums.Enum)
	v.Widget.SetEnum(e)
	v.Widget.OnChange(func(e events.Event) {
		v.SetValue(v.Widget.CurrentItem.Value)
	})
}

func (v *EnumValue) Update() {
	npv := laser.NonPtrValue(v.Value)
	v.Widget.SetCurrentValue(npv.Interface())
}

// BitFlagValue represents an [enums.BitFlag] value with chip switches.
type BitFlagValue struct {
	ValueBase[*gi.Switches]
}

func (v *BitFlagValue) Config() {
	v.Widget.SetType(gi.SwitchChip).SetEnum(v.EnumValue())
	v.Widget.OnChange(func(e events.Event) {
		v.Widget.BitFlagValue(v.EnumValue())
	})
}

func (v *BitFlagValue) Update() {
	v.Widget.UpdateFromBitFlag(v.EnumValue())
}

// EnumValue returns the underlying [enums.BitFlagSetter] value.
func (v *BitFlagValue) EnumValue() enums.BitFlagSetter {
	// special case to use [ki.Ki.FlagType] if we are the Flags field
	if v.Field != nil && v.Field.Name == "Flags" {
		if k, ok := v.Owner.(ki.Ki); ok {
			return k.FlagType()
		}
	}
	e, _ := v.Value.Interface().(enums.BitFlagSetter)
	return e
}

// TypeValue represents a [gti.Type] value with a chooser.
type TypeValue struct {
	ValueBase[*gi.Chooser]
}

func (v *TypeValue) Config() {
	typEmbeds := gi.WidgetBaseType
	if tetag, ok := v.Tag("type-embeds"); ok {
		typ := gti.TypeByName(tetag)
		if typ != nil {
			typEmbeds = typ
		}
	}

	tl := gti.AllEmbeddersOf(typEmbeds)
	v.Widget.SetTypes(tl)
	v.Widget.OnChange(func(e events.Event) {
		tval := v.Widget.CurrentItem.Value.(*gti.Type)
		v.SetValue(tval)
	})
}

func (v *TypeValue) Update() {
	opv := laser.OnePtrValue(v.Value)
	typ, _ := opv.Interface().(*gti.Type)
	v.Widget.SetCurrentValue(typ)
}

//////////////////////////////////////////////////////////////////////////////
//  ByteSliceValue

// ByteSliceValue presents a textfield of the bytes
type ByteSliceValue struct {
	ValueBase
}

func (vv *ByteSliceValue) WidgetType() *gti.Type {
	vv.WidgetTyp = gi.TextFieldType
	return vv.WidgetTyp
}

func (vv *ByteSliceValue) UpdateWidget() {
	if vv.Widget == nil {
		return
	}
	tf := vv.Widget.(*gi.TextField)
	npv := laser.NonPtrValue(vv.Value)
	bv, ok := npv.Interface().([]byte)
	if ok {
		tf.SetText(string(bv))
	}
}

func (vv *ByteSliceValue) Config(w gi.Widget) {
	if vv.Widget == w {
		vv.UpdateWidget()
		return
	}
	vv.Widget = w
	tf := vv.Widget.(*gi.TextField)
	tf.Tooltip = vv.Doc()
	// STYTODO: figure out how how to handle these kinds of styles
	tf.Style(func(s *styles.Style) {
		s.Min.X.Ch(16)
	})
	vv.StdConfig(w)

	tf.OnFinal(events.Change, func(e events.Event) {
		vv.SetValue(tf.Text())
	})
	vv.UpdateWidget()
}

//////////////////////////////////////////////////////////////////////////////
//  RuneSliceValue

// RuneSliceValue presents a textfield of the bytes
type RuneSliceValue struct {
	ValueBase
}

func (vv *RuneSliceValue) WidgetType() *gti.Type {
	vv.WidgetTyp = gi.TextFieldType
	return vv.WidgetTyp
}

func (vv *RuneSliceValue) UpdateWidget() {
	if vv.Widget == nil {
		return
	}
	tf := vv.Widget.(*gi.TextField)
	npv := laser.NonPtrValue(vv.Value)
	rv, ok := npv.Interface().([]rune)
	if ok {
		tf.SetText(string(rv))
	}
}

func (vv *RuneSliceValue) Config(w gi.Widget) {
	if vv.Widget == w {
		vv.UpdateWidget()
		return
	}
	vv.Widget = w
	tf := vv.Widget.(*gi.TextField)
	tf.Tooltip = vv.Doc()
	tf.Style(func(s *styles.Style) {
		s.Min.X.Ch(16)
	})
	vv.StdConfig(w)

	tf.OnFinal(events.Change, func(e events.Event) {
		vv.SetValue(tf.Text())
	})
	vv.UpdateWidget()
}

//////////////////////////////////////////////////////////////////////////////
//  NilValue

// NilValue presents a label saying 'nil' -- for any nil or otherwise unrepresentable items
type NilValue struct {
	ValueBase
}

func (vv *NilValue) WidgetType() *gti.Type {
	vv.WidgetTyp = gi.LabelType
	return vv.WidgetTyp
}

func (vv *NilValue) UpdateWidget() {
	if vv.Widget == nil {
		return
	}
	lb := vv.Widget.(*gi.Label)
	lb.SetText("None")
}

func (vv *NilValue) Config(w gi.Widget) {
	if vv.Widget == w {
		vv.UpdateWidget()
		return
	}
	vv.Widget = w
	vv.StdConfig(w)
	lb := vv.Widget.(*gi.Label)
	lb.Tooltip = vv.Doc()
	vv.UpdateWidget()
}

//////////////////////////////////////////////////////////////////////////////
//  IconValue

// IconValue presents an action for displaying an IconName and selecting
// icons from IconChooserDialog
type IconValue struct {
	ValueBase
}

func (vv *IconValue) WidgetType() *gti.Type {
	vv.WidgetTyp = gi.ButtonType
	return vv.WidgetTyp
}

func (vv *IconValue) UpdateWidget() {
	if vv.Widget == nil {
		return
	}
	bt := vv.Widget.(*gi.Button)
	txt := laser.ToString(vv.Value.Interface())
	if icons.Icon(txt).IsNil() {
		bt.SetIcon(icons.Blank)
	} else {
		bt.SetIcon(icons.Icon(txt))
	}
	if sntag, ok := vv.Tag("view"); ok {
		if strings.Contains(sntag, "show-name") {
			if txt == "" {
				txt = "None"
			}
			bt.SetText(strcase.ToSentence(txt))
		}
	}
	bt.Update()
}

func (vv *IconValue) Config(w gi.Widget) {
	if vv.Widget == w {
		vv.UpdateWidget()
		return
	}
	vv.Widget = w
	vv.StdConfig(w)
	bt := vv.Widget.(*gi.Button)
	bt.SetType(gi.ButtonTonal)
	ConfigDialogWidget(vv, bt, false)
	vv.UpdateWidget()
}

func (vv *IconValue) HasDialog() bool { return true }
func (vv *IconValue) OpenDialog(ctx gi.Widget, fun func()) {
	OpenValueDialog(vv, ctx, fun, "Select an icon")
}

func (vv *IconValue) ConfigDialog(d *gi.Body) (bool, func()) {
	si := 0
	ics := icons.All()
	cur := icons.Icon(laser.ToString(vv.Value.Interface()))
	NewSliceView(d).SetStyleFunc(func(w gi.Widget, s *styles.Style, row int) {
		w.(*gi.Button).SetText(strcase.ToSentence(string(ics[row])))
	}).SetSlice(&ics).SetSelVal(cur).BindSelect(&si)
	return true, func() {
		if si >= 0 {
			ic := icons.AllIcons[si]
			vv.SetValue(ic)
			vv.UpdateWidget()
		}
	}
}

//////////////////////////////////////////////////////////////////////////////
//  FontValue

// FontValue presents an action for displaying a FontName and selecting
// fonts from FontChooserDialog
type FontValue struct {
	ValueBase
}

func (vv *FontValue) WidgetType() *gti.Type {
	vv.WidgetTyp = gi.ButtonType
	return vv.WidgetTyp
}

func (vv *FontValue) UpdateWidget() {
	if vv.Widget == nil {
		return
	}
	bt := vv.Widget.(*gi.Button)
	txt := laser.ToString(vv.Value.Interface())
	bt.SetText(txt).Update()
}

func (vv *FontValue) Config(w gi.Widget) {
	if vv.Widget == w {
		vv.UpdateWidget()
		return
	}
	vv.Widget = w
	vv.StdConfig(w)
	bt := vv.Widget.(*gi.Button)
	bt.SetType(gi.ButtonTonal)
	bt.Style(func(s *styles.Style) {
		// TODO(kai): fix this not working (probably due to medium font weight)
		s.Font.Family = laser.ToString(vv.Value.Interface())
	})
	ConfigDialogWidget(vv, bt, false)
	vv.UpdateWidget()
}

func (vv *FontValue) HasDialog() bool { return true }
func (vv *FontValue) OpenDialog(ctx gi.Widget, fun func()) {
	OpenValueDialog(vv, ctx, fun, "Select a font")
}

// show fonts in a bigger size so you can actually see the differences
var FontChooserSize = units.Pt(18)

func (vv *FontValue) ConfigDialog(d *gi.Body) (bool, func()) {
	si := 0
	wb := vv.Widget.AsWidget()
	FontChooserSize.ToDots(&wb.Styles.UnitContext)
	paint.FontLibrary.OpenAllFonts(int(FontChooserSize.Dots))
	fi := paint.FontLibrary.FontInfo
	cur := gi.FontName(laser.ToString(vv.Value.Interface()))
	NewTableView(d).SetStyleFunc(func(w gi.Widget, s *styles.Style, row, col int) {
		if col != 4 {
			return
		}
		s.Font.Family = fi[row].Name
		s.Font.Stretch = fi[row].Stretch
		s.Font.Weight = fi[row].Weight
		s.Font.Style = fi[row].Style
		s.Font.Size = FontChooserSize
	}).SetSlice(&fi).SetSelVal(cur).SetSelField("Name").BindSelect(&si)

	return true, func() {
		if si >= 0 {
			fi := paint.FontLibrary.FontInfo[si]
			vv.SetValue(fi.Name)
			vv.UpdateWidget()
		}
	}
}

//////////////////////////////////////////////////////////////////////////////
//  FileValue

// FileValue presents an action for displaying a Filename and selecting
// icons from FileChooserDialog
type FileValue struct {
	ValueBase
}

func (vv *FileValue) WidgetType() *gti.Type {
	vv.WidgetTyp = gi.ButtonType
	return vv.WidgetTyp
}

func (vv *FileValue) UpdateWidget() {
	if vv.Widget == nil {
		return
	}
	bt := vv.Widget.(*gi.Button)
	txt := laser.ToString(vv.Value.Interface())
	if txt == "" {
		txt = "(click to open file chooser)"
	}
	prev := bt.Text
	bt.SetText(txt)
	if txt != prev {
		bt.Update()
	}
}

func (vv *FileValue) Config(w gi.Widget) {
	if vv.Widget == w {
		vv.UpdateWidget()
		return
	}
	vv.Widget = w
	vv.StdConfig(w)
	bt := vv.Widget.(*gi.Button)
	bt.SetType(gi.ButtonTonal)
	ConfigDialogWidget(vv, bt, false)
	vv.UpdateWidget()
}

func (vv *FileValue) HasDialog() bool                      { return true }
func (vv *FileValue) OpenDialog(ctx gi.Widget, fun func()) { OpenValueDialog(vv, ctx, fun) }

func (vv *FileValue) ConfigDialog(d *gi.Body) (bool, func()) {
	vv.SetFlag(true, ValueDialogNewWindow) // default to new window on supported platforms
	cur := laser.ToString(vv.Value.Interface())
	ext, _ := vv.Tag("ext")
	fv := NewFileView(d).SetFilename(cur, ext)
	d.AddAppBar(fv.ConfigToolbar)
	return true, func() {
		cur = fv.SelectedFile()
		vv.SetValue(cur)
		vv.UpdateWidget()
	}
}

//////////////////////////////////////////////////////////////////////////////
//  FuncValue

// FuncValue presents a [FuncButton] for viewing the information of and calling a function
type FuncValue struct {
	ValueBase
}

func (vv *FuncValue) WidgetType() *gti.Type {
	vv.WidgetTyp = FuncButtonType
	return vv.WidgetTyp
}

func (vv *FuncValue) UpdateWidget() {
	if vv.Widget == nil {
		return
	}
	fbt := vv.Widget.(*FuncButton)
	fun := laser.NonPtrValue(vv.Value).Interface()
	// if someone is viewing an arbitrary function, there is a good chance
	// that it is not added to gti (and that is out of their control)
	// (eg: in the inspector).
	fbt.SetWarnUnadded(false)
	fbt.SetFunc(fun)
}

func (vv *FuncValue) Config(w gi.Widget) {
	if vv.Widget == w {
		vv.UpdateWidget()
		return
	}
	vv.Widget = w
	vv.StdConfig(w)

	fbt := vv.Widget.(*FuncButton)
	fbt.Type = gi.ButtonTonal

	vv.UpdateWidget()
}

//////////////////////////////////////////////////////////////////////////////
//  OptionValue

// OptionValue presents an [option.Option]
type OptionValue struct {
	ValueBase
}

func (vv *OptionValue) WidgetType() *gti.Type {
	vv.WidgetTyp = gi.FrameType
	return vv.WidgetTyp
}

func (vv *OptionValue) UpdateWidget() {
	if vv.Widget == nil {
		return
	}
}

func (vv *OptionValue) Config(w gi.Widget) {
	if vv.Widget == w {
		vv.UpdateWidget()
		return
	}
	vv.Widget = w
	vv.StdConfig(w)

	fr := vv.Widget.(*gi.Frame)

	gi.NewButton(fr, "unset").SetText("Unset")
	val := vv.Value.FieldByName("Value").Interface()
	NewValue(fr, val, "value")

	vv.UpdateWidget()
}
