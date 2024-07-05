// Code generated by 'yaegi extract cogentcore.org/core/events'. DO NOT EDIT.

package symbols

import (
	"cogentcore.org/core/enums"
	"cogentcore.org/core/events"
	"cogentcore.org/core/events/key"
	"image"
	"reflect"
	"time"
)

func init() {
	Symbols["cogentcore.org/core/events/events"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"ButtonsN":              reflect.ValueOf(events.ButtonsN),
		"ButtonsValues":         reflect.ValueOf(events.ButtonsValues),
		"Change":                reflect.ValueOf(events.Change),
		"Click":                 reflect.ValueOf(events.Click),
		"Close":                 reflect.ValueOf(events.Close),
		"ContextMenu":           reflect.ValueOf(events.ContextMenu),
		"Custom":                reflect.ValueOf(events.Custom),
		"DefaultModBits":        reflect.ValueOf(events.DefaultModBits),
		"DoubleClick":           reflect.ValueOf(events.DoubleClick),
		"DragEnter":             reflect.ValueOf(events.DragEnter),
		"DragLeave":             reflect.ValueOf(events.DragLeave),
		"DragMove":              reflect.ValueOf(events.DragMove),
		"DragStart":             reflect.ValueOf(events.DragStart),
		"Drop":                  reflect.ValueOf(events.Drop),
		"DropCopy":              reflect.ValueOf(events.DropCopy),
		"DropDeleteSource":      reflect.ValueOf(events.DropDeleteSource),
		"DropIgnore":            reflect.ValueOf(events.DropIgnore),
		"DropLink":              reflect.ValueOf(events.DropLink),
		"DropModsN":             reflect.ValueOf(events.DropModsN),
		"DropModsValues":        reflect.ValueOf(events.DropModsValues),
		"DropMove":              reflect.ValueOf(events.DropMove),
		"EventFlagsN":           reflect.ValueOf(events.EventFlagsN),
		"EventFlagsValues":      reflect.ValueOf(events.EventFlagsValues),
		"ExtendContinuous":      reflect.ValueOf(events.ExtendContinuous),
		"ExtendOne":             reflect.ValueOf(events.ExtendOne),
		"Focus":                 reflect.ValueOf(events.Focus),
		"FocusLost":             reflect.ValueOf(events.FocusLost),
		"Handled":               reflect.ValueOf(events.Handled),
		"Input":                 reflect.ValueOf(events.Input),
		"KeyChord":              reflect.ValueOf(events.KeyChord),
		"KeyDown":               reflect.ValueOf(events.KeyDown),
		"KeyUp":                 reflect.ValueOf(events.KeyUp),
		"Left":                  reflect.ValueOf(events.Left),
		"LongHoverEnd":          reflect.ValueOf(events.LongHoverEnd),
		"LongHoverStart":        reflect.ValueOf(events.LongHoverStart),
		"LongPressEnd":          reflect.ValueOf(events.LongPressEnd),
		"LongPressStart":        reflect.ValueOf(events.LongPressStart),
		"Magnify":               reflect.ValueOf(events.Magnify),
		"Middle":                reflect.ValueOf(events.Middle),
		"MouseDown":             reflect.ValueOf(events.MouseDown),
		"MouseDrag":             reflect.ValueOf(events.MouseDrag),
		"MouseEnter":            reflect.ValueOf(events.MouseEnter),
		"MouseLeave":            reflect.ValueOf(events.MouseLeave),
		"MouseMove":             reflect.ValueOf(events.MouseMove),
		"MouseUp":               reflect.ValueOf(events.MouseUp),
		"NewDragDrop":           reflect.ValueOf(events.NewDragDrop),
		"NewExternalDrop":       reflect.ValueOf(events.NewExternalDrop),
		"NewKey":                reflect.ValueOf(events.NewKey),
		"NewMagnify":            reflect.ValueOf(events.NewMagnify),
		"NewMouse":              reflect.ValueOf(events.NewMouse),
		"NewMouseDrag":          reflect.ValueOf(events.NewMouseDrag),
		"NewMouseMove":          reflect.ValueOf(events.NewMouseMove),
		"NewOSEvent":            reflect.ValueOf(events.NewOSEvent),
		"NewOSFiles":            reflect.ValueOf(events.NewOSFiles),
		"NewScroll":             reflect.ValueOf(events.NewScroll),
		"NewTouch":              reflect.ValueOf(events.NewTouch),
		"NewWindow":             reflect.ValueOf(events.NewWindow),
		"NewWindowPaint":        reflect.ValueOf(events.NewWindowPaint),
		"NewWindowResize":       reflect.ValueOf(events.NewWindowResize),
		"NoButton":              reflect.ValueOf(events.NoButton),
		"NoDropMod":             reflect.ValueOf(events.NoDropMod),
		"NoSelect":              reflect.ValueOf(events.NoSelect),
		"NoWinAction":           reflect.ValueOf(events.NoWinAction),
		"OS":                    reflect.ValueOf(events.OS),
		"OSOpenFiles":           reflect.ValueOf(events.OSOpenFiles),
		"Right":                 reflect.ValueOf(events.Right),
		"Rotate":                reflect.ValueOf(events.Rotate),
		"ScreenUpdate":          reflect.ValueOf(events.ScreenUpdate),
		"Scroll":                reflect.ValueOf(events.Scroll),
		"ScrollWheelSpeed":      reflect.ValueOf(&events.ScrollWheelSpeed).Elem(),
		"Select":                reflect.ValueOf(events.Select),
		"SelectModeBits":        reflect.ValueOf(events.SelectModeBits),
		"SelectModesN":          reflect.ValueOf(events.SelectModesN),
		"SelectModesValues":     reflect.ValueOf(events.SelectModesValues),
		"SelectOne":             reflect.ValueOf(events.SelectOne),
		"SelectQuiet":           reflect.ValueOf(events.SelectQuiet),
		"Show":                  reflect.ValueOf(events.Show),
		"SlideMove":             reflect.ValueOf(events.SlideMove),
		"SlideStart":            reflect.ValueOf(events.SlideStart),
		"SlideStop":             reflect.ValueOf(events.SlideStop),
		"TouchEnd":              reflect.ValueOf(events.TouchEnd),
		"TouchMove":             reflect.ValueOf(events.TouchMove),
		"TouchStart":            reflect.ValueOf(events.TouchStart),
		"TraceEventCompression": reflect.ValueOf(&events.TraceEventCompression).Elem(),
		"TraceWindowPaint":      reflect.ValueOf(&events.TraceWindowPaint).Elem(),
		"TripleClick":           reflect.ValueOf(events.TripleClick),
		"TypesN":                reflect.ValueOf(events.TypesN),
		"TypesValues":           reflect.ValueOf(events.TypesValues),
		"Unique":                reflect.ValueOf(events.Unique),
		"UnknownType":           reflect.ValueOf(events.UnknownType),
		"Unselect":              reflect.ValueOf(events.Unselect),
		"UnselectQuiet":         reflect.ValueOf(events.UnselectQuiet),
		"WinActionsN":           reflect.ValueOf(events.WinActionsN),
		"WinActionsValues":      reflect.ValueOf(events.WinActionsValues),
		"WinClose":              reflect.ValueOf(events.WinClose),
		"WinFocus":              reflect.ValueOf(events.WinFocus),
		"WinFocusLost":          reflect.ValueOf(events.WinFocusLost),
		"WinMinimize":           reflect.ValueOf(events.WinMinimize),
		"WinMove":               reflect.ValueOf(events.WinMove),
		"WinShow":               reflect.ValueOf(events.WinShow),
		"Window":                reflect.ValueOf(events.Window),
		"WindowPaint":           reflect.ValueOf(events.WindowPaint),
		"WindowResize":          reflect.ValueOf(events.WindowResize),

		// type definitions
		"Base":         reflect.ValueOf((*events.Base)(nil)),
		"Buttons":      reflect.ValueOf((*events.Buttons)(nil)),
		"CustomEvent":  reflect.ValueOf((*events.CustomEvent)(nil)),
		"Deque":        reflect.ValueOf((*events.Deque)(nil)),
		"DragDrop":     reflect.ValueOf((*events.DragDrop)(nil)),
		"DropMods":     reflect.ValueOf((*events.DropMods)(nil)),
		"Event":        reflect.ValueOf((*events.Event)(nil)),
		"EventFlags":   reflect.ValueOf((*events.EventFlags)(nil)),
		"Key":          reflect.ValueOf((*events.Key)(nil)),
		"Listeners":    reflect.ValueOf((*events.Listeners)(nil)),
		"Mouse":        reflect.ValueOf((*events.Mouse)(nil)),
		"MouseScroll":  reflect.ValueOf((*events.MouseScroll)(nil)),
		"OSEvent":      reflect.ValueOf((*events.OSEvent)(nil)),
		"OSFiles":      reflect.ValueOf((*events.OSFiles)(nil)),
		"SelectModes":  reflect.ValueOf((*events.SelectModes)(nil)),
		"Sequence":     reflect.ValueOf((*events.Sequence)(nil)),
		"Source":       reflect.ValueOf((*events.Source)(nil)),
		"SourceState":  reflect.ValueOf((*events.SourceState)(nil)),
		"Touch":        reflect.ValueOf((*events.Touch)(nil)),
		"TouchMagnify": reflect.ValueOf((*events.TouchMagnify)(nil)),
		"Types":        reflect.ValueOf((*events.Types)(nil)),
		"WinActions":   reflect.ValueOf((*events.WinActions)(nil)),
		"WindowEvent":  reflect.ValueOf((*events.WindowEvent)(nil)),

		// interface wrapper definitions
		"_Event": reflect.ValueOf((*_cogentcore_org_core_events_Event)(nil)),
	}
}

// _cogentcore_org_core_events_Event is an interface wrapper for Event type
type _cogentcore_org_core_events_Event struct {
	IValue           interface{}
	WAsBase          func() *events.Base
	WClearHandled    func()
	WClone           func() events.Event
	WHasAllModifiers func(mods ...enums.BitFlag) bool
	WHasAnyModifier  func(mods ...enums.BitFlag) bool
	WHasPos          func() bool
	WInit            func()
	WIsHandled       func() bool
	WIsSame          func(oth events.Event) bool
	WIsUnique        func() bool
	WKeyChord        func() key.Chord
	WKeyCode         func() key.Codes
	WKeyRune         func() rune
	WLocalOff        func() image.Point
	WModifiers       func() key.Modifiers
	WMouseButton     func() events.Buttons
	WNeedsFocus      func() bool
	WNewFromClone    func(typ events.Types) events.Event
	WPos             func() image.Point
	WPrevDelta       func() image.Point
	WPrevPos         func() image.Point
	WPrevTime        func() time.Time
	WSelectMode      func() events.SelectModes
	WSetHandled      func()
	WSetLocalOff     func(off image.Point)
	WSetTime         func()
	WSincePrev       func() time.Duration
	WSinceStart      func() time.Duration
	WStartDelta      func() image.Point
	WStartPos        func() image.Point
	WStartTime       func() time.Time
	WString          func() string
	WTime            func() time.Time
	WType            func() events.Types
	WWindowPos       func() image.Point
	WWindowPrevPos   func() image.Point
	WWindowStartPos  func() image.Point
}

func (W _cogentcore_org_core_events_Event) AsBase() *events.Base {
	return W.WAsBase()
}
func (W _cogentcore_org_core_events_Event) ClearHandled() {
	W.WClearHandled()
}
func (W _cogentcore_org_core_events_Event) Clone() events.Event {
	return W.WClone()
}
func (W _cogentcore_org_core_events_Event) HasAllModifiers(mods ...enums.BitFlag) bool {
	return W.WHasAllModifiers(mods...)
}
func (W _cogentcore_org_core_events_Event) HasAnyModifier(mods ...enums.BitFlag) bool {
	return W.WHasAnyModifier(mods...)
}
func (W _cogentcore_org_core_events_Event) HasPos() bool {
	return W.WHasPos()
}
func (W _cogentcore_org_core_events_Event) Init() {
	W.WInit()
}
func (W _cogentcore_org_core_events_Event) IsHandled() bool {
	return W.WIsHandled()
}
func (W _cogentcore_org_core_events_Event) IsSame(oth events.Event) bool {
	return W.WIsSame(oth)
}
func (W _cogentcore_org_core_events_Event) IsUnique() bool {
	return W.WIsUnique()
}
func (W _cogentcore_org_core_events_Event) KeyChord() key.Chord {
	return W.WKeyChord()
}
func (W _cogentcore_org_core_events_Event) KeyCode() key.Codes {
	return W.WKeyCode()
}
func (W _cogentcore_org_core_events_Event) KeyRune() rune {
	return W.WKeyRune()
}
func (W _cogentcore_org_core_events_Event) LocalOff() image.Point {
	return W.WLocalOff()
}
func (W _cogentcore_org_core_events_Event) Modifiers() key.Modifiers {
	return W.WModifiers()
}
func (W _cogentcore_org_core_events_Event) MouseButton() events.Buttons {
	return W.WMouseButton()
}
func (W _cogentcore_org_core_events_Event) NeedsFocus() bool {
	return W.WNeedsFocus()
}
func (W _cogentcore_org_core_events_Event) NewFromClone(typ events.Types) events.Event {
	return W.WNewFromClone(typ)
}
func (W _cogentcore_org_core_events_Event) Pos() image.Point {
	return W.WPos()
}
func (W _cogentcore_org_core_events_Event) PrevDelta() image.Point {
	return W.WPrevDelta()
}
func (W _cogentcore_org_core_events_Event) PrevPos() image.Point {
	return W.WPrevPos()
}
func (W _cogentcore_org_core_events_Event) PrevTime() time.Time {
	return W.WPrevTime()
}
func (W _cogentcore_org_core_events_Event) SelectMode() events.SelectModes {
	return W.WSelectMode()
}
func (W _cogentcore_org_core_events_Event) SetHandled() {
	W.WSetHandled()
}
func (W _cogentcore_org_core_events_Event) SetLocalOff(off image.Point) {
	W.WSetLocalOff(off)
}
func (W _cogentcore_org_core_events_Event) SetTime() {
	W.WSetTime()
}
func (W _cogentcore_org_core_events_Event) SincePrev() time.Duration {
	return W.WSincePrev()
}
func (W _cogentcore_org_core_events_Event) SinceStart() time.Duration {
	return W.WSinceStart()
}
func (W _cogentcore_org_core_events_Event) StartDelta() image.Point {
	return W.WStartDelta()
}
func (W _cogentcore_org_core_events_Event) StartPos() image.Point {
	return W.WStartPos()
}
func (W _cogentcore_org_core_events_Event) StartTime() time.Time {
	return W.WStartTime()
}
func (W _cogentcore_org_core_events_Event) String() string {
	if W.WString == nil {
		return ""
	}
	return W.WString()
}
func (W _cogentcore_org_core_events_Event) Time() time.Time {
	return W.WTime()
}
func (W _cogentcore_org_core_events_Event) Type() events.Types {
	return W.WType()
}
func (W _cogentcore_org_core_events_Event) WindowPos() image.Point {
	return W.WWindowPos()
}
func (W _cogentcore_org_core_events_Event) WindowPrevPos() image.Point {
	return W.WWindowPrevPos()
}
func (W _cogentcore_org_core_events_Event) WindowStartPos() image.Point {
	return W.WWindowStartPos()
}
