// Copyright (c) 2018, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"fmt"
	"log/slog"
	"time"
	"unicode"

	"cogentcore.org/core/enums"
	"cogentcore.org/core/events"
	"cogentcore.org/core/keymap"
	"cogentcore.org/core/math32"
	"cogentcore.org/core/parse/complete"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/styles/abilities"
	"cogentcore.org/core/tree"
)

// FrameFlags has state bit flags for [Frame].
type FrameFlags WidgetFlags //enums:bitflag -trim-prefix Frame

const (
	// FrameStackTopOnly is whether to only layout the top widget for a stacked
	// frame layout. This is appropriate for e.g., tab layout, which does a full
	// redraw on stack changes, but not for e.g., check boxes which don't.
	FrameStackTopOnly FrameFlags = FrameFlags(WidgetFlagsN) + iota
)

// Frame is the primary node type responsible for organizing the sizes
// and positions of child widgets. It also renders the standard box model.
// All collections of widgets should generally be contained within a [Frame];
// otherwise, the parent widget must take over responsibility for positioning.
// Frames automatically can add scrollbars depending on the [styles.Style.Overflow].
//
// For a [styles.Grid] layout, the [styles.Style.Columns] property should
// generally be set to the desired number of columns, from which the number of rows
// is computed; otherwise, it uses the square root of number of
// elements.
type Frame struct {
	WidgetBase

	// StackTop, for a [styles.Stacked] layout, is the index of the node to use as the top of the stack.
	// Only the node at this index is rendered; if not a valid index, nothing is rendered.
	StackTop int `set:"-"`

	// LayImpl contains implementation state info for doing layout
	LayImpl LayImplState `edit:"-" copier:"-" json:"-" xml:"-" set:"-"`

	// whether scrollbar is used for given dim
	HasScroll [2]bool `edit:"-" copier:"-" json:"-" xml:"-" set:"-"`

	// scroll bars -- we fully manage them as needed
	Scrolls [2]*Slider `edit:"-" copier:"-" json:"-" xml:"-" set:"-"`

	// accumulated name to search for when keys are typed
	FocusName string `edit:"-" copier:"-" json:"-" xml:"-" set:"-"`

	// time of last focus name event -- for timeout
	FocusNameTime time.Time `edit:"-" copier:"-" json:"-" xml:"-" set:"-"`

	// last element focused on -- used as a starting point if name is the same
	FocusNameLast tree.Node `edit:"-" copier:"-" json:"-" xml:"-" set:"-"`
}

func (ly *Frame) FlagType() enums.BitFlagSetter {
	return (*FrameFlags)(&ly.Flags)
}

func (ly *Frame) OnInit() {
	ly.WidgetBase.OnInit()
	ly.SetStyles()
	ly.HandleEvents()
}

func (ly *Frame) ApplyStyle() {
	ly.ApplyStyleWidget()
	for d := math32.X; d <= math32.Y; d++ {
		if ly.HasScroll[d] && ly.Scrolls[d] != nil {
			ly.Scrolls[d].ApplyStyle()
		}
	}
}

func (ly *Frame) SetStyles() {
	ly.Style(func(s *styles.Style) {
		// we never want borders on layouts
		s.MaxBorder = styles.Border{}
	})
	ly.StyleFinal(func(s *styles.Style) {
		s.SetAbilities(s.Overflow.X == styles.OverflowAuto || s.Overflow.Y == styles.OverflowAuto, abilities.Scrollable, abilities.Slideable)
	})
}

func (ly *Frame) Destroy() {
	for d := math32.X; d <= math32.Y; d++ {
		ly.DeleteScroll(d)
	}
	ly.WidgetBase.Destroy()
}

// DeleteScroll deletes scrollbar along given dimesion.
func (ly *Frame) DeleteScroll(d math32.Dims) {
	if ly.Scrolls[d] == nil {
		return
	}
	sb := ly.Scrolls[d]
	sb.This().Destroy()
	ly.Scrolls[d] = nil
}

func (ly *Frame) RenderChildren() {
	if ly.Styles.Display == styles.Stacked {
		kwi, _ := ly.StackTopWidget()
		if kwi != nil {
			kwi.RenderWidget()
		}
		return
	}
	ly.WidgetKidsIter(func(i int, kwi Widget, kwb *WidgetBase) bool {
		kwi.RenderWidget()
		return tree.Continue
	})
}

func (ly *Frame) RenderWidget() {
	if ly.PushBounds() {
		ly.This().(Widget).Render()
		ly.RenderParts()
		ly.RenderChildren()
		ly.RenderScrolls()
		ly.PopBounds()
	}
}

// ChildWithFocus returns a direct child of this layout that either is the
// current window focus item, or contains that focus item (along with its
// index) -- nil, -1 if none.
func (ly *Frame) ChildWithFocus() (Widget, int) {
	em := ly.Events()
	if em == nil {
		return nil, -1
	}
	var foc Widget
	focIndex := -1
	ly.WidgetKidsIter(func(i int, kwi Widget, kwb *WidgetBase) bool {
		if kwb.ContainsFocus() {
			foc = kwi
			focIndex = i
			return tree.Break
		}
		return tree.Continue
	})
	return foc, focIndex
}

// FocusNextChild attempts to move the focus into the next layout child
// (with wraparound to start); returns true if successful.
// if updn is true, then for Grid layouts, it moves down to next row
// instead of just the sequentially next item.
func (ly *Frame) FocusNextChild(updn bool) bool {
	sz := len(ly.Kids)
	if sz <= 1 {
		return false
	}
	foc, idx := ly.ChildWithFocus()
	if foc == nil {
		fmt.Println("no child foc")
		return false
	}
	em := ly.Events()
	if em == nil {
		return false
	}
	cur := em.Focus
	nxti := idx + 1
	if ly.Styles.Display == styles.Grid && updn {
		nxti = idx + ly.Styles.Columns
	}
	did := false
	if nxti < sz {
		nx := ly.Child(nxti).(Widget)
		did = em.FocusOnOrNext(nx)
	} else {
		nx := ly.Child(0).(Widget)
		did = em.FocusOnOrNext(nx)
	}
	if !did || em.Focus == cur {
		return false
	}
	return true
}

// FocusPreviousChild attempts to move the focus into the previous layout child
// (with wraparound to end); returns true if successful.
// If updn is true, then for Grid layouts, it moves up to next row
// instead of just the sequentially next item.
func (ly *Frame) FocusPreviousChild(updn bool) bool {
	sz := len(ly.Kids)
	if sz <= 1 {
		return false
	}
	foc, idx := ly.ChildWithFocus()
	if foc == nil {
		return false
	}
	em := ly.Events()
	if em == nil {
		return false
	}
	cur := em.Focus
	nxti := idx - 1
	if ly.Styles.Display == styles.Grid && updn {
		nxti = idx - ly.Styles.Columns
	}
	did := false
	if nxti >= 0 {
		did = em.FocusOnOrPrev(ly.Child(nxti).(Widget))
	} else {
		did = em.FocusOnOrPrev(ly.Child(sz - 1).(Widget))
	}
	if !did || em.Focus == cur {
		return false
	}
	return true
}

func (ly *Frame) HandleEvents() {
	ly.WidgetBase.HandleEvents()
	ly.HandleKeys()
	ly.On(events.Scroll, func(e events.Event) {
		ly.ScrollDelta(e)
	})
	// we treat slide events on layouts as scroll events
	// we must reverse the delta for "natural" scrolling behavior
	ly.On(events.SlideMove, func(e events.Event) {
		del := math32.Vector2FromPoint(e.PrevDelta()).MulScalar(-0.1)
		ly.ScrollDelta(events.NewScroll(e.WindowPos(), del, e.Modifiers()))
	})
}

// HandleKeys handles all key events for navigating focus within a Layout.
// Typically this is done by the parent Scene level layout, but can be
// done by default if FocusWithinable Ability is set.
func (ly *Frame) HandleKeys() {
	ly.OnFinal(events.KeyChord, func(e events.Event) {
		kf := keymap.Of(e.KeyChord())
		if DebugSettings.KeyEventTrace {
			slog.Info("Layout KeyInput", "widget", ly, "keyFunction", kf)
		}
		if kf == keymap.Abort {
			if ly.Scene.Stage.ClosePopupAndBelow() {
				e.SetHandled()
			}
			return
		}
		em := ly.Events()
		if em == nil {
			return
		}
		grid := ly.Styles.Display == styles.Grid
		if ly.Styles.Direction == styles.Row || grid {
			switch kf {
			case keymap.MoveRight:
				if ly.FocusNextChild(false) {
					e.SetHandled()
				}
				return
			case keymap.MoveLeft:
				if ly.FocusPreviousChild(false) {
					e.SetHandled()
				}
				return
			}
		}
		if ly.Styles.Direction == styles.Column || grid {
			switch kf {
			case keymap.MoveDown:
				if ly.FocusNextChild(true) {
					e.SetHandled()
				}
				return
			case keymap.MoveUp:
				if ly.FocusPreviousChild(true) {
					e.SetHandled()
				}
				return
			case keymap.PageDown:
				proc := false
				for st := 0; st < SystemSettings.LayoutPageSteps; st++ {
					if !ly.FocusNextChild(true) {
						break
					}
					proc = true
				}
				if proc {
					e.SetHandled()
				}
				return
			case keymap.PageUp:
				proc := false
				for st := 0; st < SystemSettings.LayoutPageSteps; st++ {
					if !ly.FocusPreviousChild(true) {
						break
					}
					proc = true
				}
				if proc {
					e.SetHandled()
				}
				return
			}
		}
		ly.FocusOnName(e)
	})
}

// FocusOnName processes key events to look for an element starting with given name
func (ly *Frame) FocusOnName(e events.Event) bool {
	kf := keymap.Of(e.KeyChord())
	if DebugSettings.KeyEventTrace {
		slog.Info("Layout FocusOnName", "widget", ly, "keyFunction", kf)
	}
	delay := e.Time().Sub(ly.FocusNameTime)
	ly.FocusNameTime = e.Time()
	if kf == keymap.FocusNext { // tab means go to next match -- don't worry about time
		if ly.FocusName == "" || delay > SystemSettings.LayoutFocusNameTabTime {
			ly.FocusName = ""
			ly.FocusNameLast = nil
			return false
		}
	} else {
		if delay > SystemSettings.LayoutFocusNameTimeout {
			ly.FocusName = ""
		}
		if !unicode.IsPrint(e.KeyRune()) || e.Modifiers() != 0 {
			return false
		}
		sr := string(e.KeyRune())
		if ly.FocusName == sr {
			// re-search same letter
		} else {
			ly.FocusName += sr
			ly.FocusNameLast = nil // only use last if tabbing
		}
	}
	// e.SetHandled()
	// fmt.Printf("searching for: %v  last: %v\n", ly.FocusName, ly.FocusNameLast)
	focel := ChildByLabelCanFocus(ly, ly.FocusName, ly.FocusNameLast)
	if focel != nil {
		focel = focel.This()
		em := ly.Events()
		if em != nil {
			em.SetFocusEvent(focel.(Widget)) // this will also scroll by default!
		}
		ly.FocusNameLast = focel
		return true
	} else {
		if ly.FocusNameLast == nil {
			ly.FocusName = "" // nothing being found
		}
		ly.FocusNameLast = nil // start over
	}
	return false
}

// ChildByLabelCanFocus uses breadth-first search to find
// the first focusable element within the layout whose Label (using
// [ToLabel]) matches the given name using [complete.IsSeedMatching].
// If after is non-nil, it only finds after that element.
func ChildByLabelCanFocus(ly *Frame, name string, after tree.Node) tree.Node {
	gotAfter := false
	completions := []complete.Completion{}
	ly.WalkDownBreadth(func(k tree.Node) bool {
		if k == ly.This() { // skip us
			return tree.Continue
		}
		_, ni := AsWidget(k)
		if ni == nil || !ni.CanFocus() { // don't go any further
			return tree.Continue
		}
		if after != nil && !gotAfter {
			if k == after {
				gotAfter = true
			}
			return tree.Continue // skip to next
		}
		completions = append(completions, complete.Completion{
			Text: ToLabel(k),
			Desc: k.PathFrom(ly),
		})
		return tree.Continue
	})
	matches := complete.MatchSeedCompletion(completions, name)
	if len(matches) > 0 {
		return ly.FindPath(matches[0].Desc)
	}
	return nil
}

// Stretch and Space: spacing elements for layouts

// Stretch adds a stretchy element that grows to fill all
// available space. You can set [styles.Style.Grow] to change
// how much it grows relative to other growing elements.
// It does not render anything.
type Stretch struct {
	WidgetBase
}

func (st *Stretch) OnInit() {
	st.WidgetBase.SetStyles()
	// note: not getting base events
	st.Style(func(s *styles.Style) {
		s.Min.X.Ch(1)
		s.Min.Y.Em(1)
		s.Grow.Set(1, 1)
	})
}

func (st *Stretch) Render() {}

// Space is a fixed size blank space, with
// a default width of 1ch and a height of 1em.
// You can set [styles.Style.Min] to change its size.
// It does not render anything.
type Space struct {
	WidgetBase
}

func (sp *Space) OnInit() {
	sp.WidgetBase.SetStyles()
	// note: not getting base events
	sp.Style(func(s *styles.Style) {
		s.Min.X.Ch(1)
		s.Min.Y.Em(1)
		s.Padding.Zero()
		s.Margin.Zero()
		s.MaxBorder.Width.Zero()
		s.Border.Width.Zero()
	})
}

func (sp *Space) Render() {}
