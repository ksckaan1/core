// Copyright (c) 2018, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gi

import (
	"image"
	"log"

	"goki.dev/girl/gist"
	"goki.dev/girl/units"
	"goki.dev/goosi"
	"goki.dev/goosi/key"
	"goki.dev/icons"
	"goki.dev/ki/v2"
	"goki.dev/mat32/v2"
)

// Menu is a slice list of Node2D actions, which can contain sub-actions
// (though it can contain anything -- it is just added to a column layout and
// displayed in a popup) -- don't use stretchy sizes in general for these
// items!
type Menu ki.Slice

func (m Menu) MarshalJSON() ([]byte, error) {
	ks := (ki.Slice)(m)
	_ = ks
	// return ks.MarshalJSON()
	return nil, nil
}

func (m *Menu) UnmarshalJSON(b []byte) error {
	ks := (*ki.Slice)(m)
	_ = ks
	// return ks.UnmarshalJSON(b)
	return nil
}

func (m *Menu) CopyFrom(men *Menu) {
	ks := (*ki.Slice)(m)
	ks.CopyFrom((ki.Slice)(*men))
}

// MakeMenuFunc is a callback for making a menu on demand, receives the object
// calling this function (typically an Action or Button) and the menu
type MakeMenuFunc func(obj ki.Ki, m *Menu)

// SetAction sets properties of given action
func (m *Menu) SetAction(ac *Action, opts ActOpts, sigTo ki.Ki, fun ki.RecvFunc) {
	nm := opts.Name
	if nm == "" {
		nm = opts.Label
	}
	if nm == "" {
		nm = string(opts.Icon)
	}
	ac.InitName(ac, nm)
	ac.Text = opts.Label
	ac.Tooltip = opts.Tooltip
	ac.Icon = icons.Icon(opts.Icon)
	ac.Shortcut = key.Chord(opts.Shortcut).OSShortcut()
	if opts.ShortcutKey != KeyFunNil {
		ac.Shortcut = ShortcutForFun(opts.ShortcutKey)
		// todo: need a flag for menu-based?
	}
	ac.Data = opts.Data
	ac.UpdateFunc = opts.UpdateFunc
	ac.SetAsMenu()
	if sigTo != nil && fun != nil {
		ac.ActionSig.Connect(sigTo, fun)
	}
}

// AddAction adds an action to the menu using given options, and connects the
// action signal to given receiver object and function, along with given data
// which is stored on the action and then passed in the action signal.
// Optional updateFunc is a function called prior to showing the menu to
// update the actions (enabled or not typically).
func (m *Menu) AddAction(opts ActOpts, sigTo ki.Ki, fun ki.RecvFunc) *Action {
	if m == nil {
		*m = make(Menu, 0, 10)
	}
	ac := &Action{}
	m.SetAction(ac, opts, sigTo, fun)
	*m = append(*m, ac.This().(Widget))
	return ac
}

// InsertActionBefore adds an action to the menu before existing item of given
// name, using given options, and connects the action signal to given receiver
// object and function, along with given data which is stored on the action
// and then passed in the action signal.  Optional updateFunc is a function
// called prior to showing the menu to update the actions (enabled or not
// typically).  If name not found, adds to end of list..
func (m *Menu) InsertActionBefore(before string, opts ActOpts, sigTo ki.Ki, fun ki.RecvFunc) *Action {
	sl := (*[]ki.Ki)(m)
	if idx, got := ki.SliceIndexByName(sl, before, 0); got {
		ac := &Action{}
		m.SetAction(ac, opts, sigTo, fun)
		ki.SliceInsert(sl, ac.This(), idx)
		return ac
	} else {
		return m.AddAction(opts, sigTo, fun)
	}
}

// InsertActionAfter adds an action to the menu after existing item of given
// name, using given options, and connects the action signal to given receiver
// object and function, along with given data which is stored on the action
// and then passed in the action signal.  Optional updateFunc is a function
// called prior to showing the menu to update the actions (enabled or not
// typically).  If name not found, adds to end of list..
func (m *Menu) InsertActionAfter(after string, opts ActOpts, sigTo ki.Ki, fun ki.RecvFunc) *Action {
	sl := (*[]ki.Ki)(m)
	if idx, got := ki.SliceIndexByName(sl, after, 0); got {
		ac := &Action{}
		m.SetAction(ac, opts, sigTo, fun)
		ki.SliceInsert(sl, ac.This(), idx+1)
		return ac
	} else {
		return m.AddAction(opts, sigTo, fun)
	}
}

// AddSeparator adds a separator at the next point in the menu (name is just
// internal label of element, defaults to 'sep' if empty)
func (m *Menu) AddSeparator(name string) *Separator {
	if m == nil {
		*m = make(Menu, 0, 10)
	}
	sp := &Separator{}
	if name == "" {
		name = "sep"
	}
	sp.InitName(sp, name)
	sp.Horiz = true
	*m = append(*m, sp.This().(Widget))
	return sp
}

// AddLabel adds a label to the menu
func (m *Menu) AddLabel(lbl string) *Label {
	if m == nil {
		*m = make(Menu, 0, 10)
	}
	lb := &Label{}
	lb.InitName(lb, lbl)
	lb.SetText(lbl)
	*m = append(*m, lb.This().(Widget))
	return lb
}

// SetShortcuts sets the shortcuts to given window -- call when the menu has
// been attached to a window
func (m *Menu) SetShortcuts(win *Window) {
	if win == nil {
		return
	}
	for _, mi := range *m {
		if ac := AsAction(mi); ac != nil {
			win.AddShortcut(ac.Shortcut, ac)
		}
	}
}

// DeleteShortcuts deletes the shortcuts in given window
func (m *Menu) DeleteShortcuts(win *Window) {
	if win == nil {
		return
	}
	for _, mi := range *m {
		if ac := AsAction(mi); ac != nil {
			win.DeleteShortcut(ac.Shortcut, ac)
		}
	}
}

// UpdateActions calls update function on all the actions in the menu, and any
// of their sub-actions
func (m *Menu) UpdateActions() {
	for _, mi := range *m {
		if ac := AsAction(mi); ac != nil {
			ac.UpdateActions()
		}
	}
}

// FindActionByName finds an action on the menu, or any sub-menu, with given
// name (exact match) -- this is not the Text label but the Name of the
// element (for AddAction items, this is the same as Label or Icon (if Label
// is empty)) -- returns false if not found
func (m *Menu) FindActionByName(name string) (*Action, bool) {
	for _, mi := range *m {
		if ac := AsAction(mi); ac != nil {
			if ac.Name() == name {
				return ac, true
			}
			if ac.Menu != nil {
				if sac, ok := ac.Menu.FindActionByName(name); ok {
					return sac, ok
				}
			}
		}
	}
	return nil, false
}

////////////////////////////////////////////////////////////////////////////////////////
// Standard menu elements

// AddCopyCutPaste adds a Copy, Cut, Paste actions that just emit the
// corresponding keyboard shortcut.  Paste is automatically enabled by
// clipboard having something in it.
func (m *Menu) AddCopyCutPaste(win *Window) {
	m.AddAction(ActOpts{Label: "Copy", ShortcutKey: KeyFunCopy},
		nil, func(recv, send ki.Ki, sig int64, data any) {
			win.EventMgr.SendKeyFunEvent(KeyFunCopy, false) // false = ignore popups -- don't send to menu
		})
	m.AddAction(ActOpts{Label: "Cut", ShortcutKey: KeyFunCut},
		nil, func(recv, send ki.Ki, sig int64, data any) {
			win.EventMgr.SendKeyFunEvent(KeyFunCut, false) // false = ignore popups -- don't send to menu
		})
	m.AddAction(ActOpts{Label: "Paste", ShortcutKey: KeyFunPaste,
		UpdateFunc: func(ac *Action) {
			ac.SetEnabledState(!goosi.TheApp.ClipBoard(win.OSWin).IsEmpty())
		}}, nil, func(recv, send ki.Ki, sig int64, data any) {
		win.EventMgr.SendKeyFunEvent(KeyFunPaste, false) // false = ignore popups -- don't send to menu
	})
}

// AddCopyCutPasteDupe adds a Copy, Cut, Paste, and Duplicate actions that
// just emit the corresponding keyboard shortcut.  Paste is automatically
// enabled by clipboard having something in it.
func (m *Menu) AddCopyCutPasteDupe(win *Window) {
	m.AddCopyCutPaste(win)
	dpsc := ActiveKeyMap.ChordForFun(KeyFunDuplicate)
	m.AddAction(ActOpts{Label: "Duplicate", Shortcut: dpsc},
		nil, func(recv, send ki.Ki, sig int64, data any) {
			win.EventMgr.SendKeyFunEvent(KeyFunDuplicate, false) // false = ignore popups -- don't send to menu
		})
}

// CustomAppMenuFunc is a function called by AddAppMenu after the
// AddStdAppMenu is called -- apps can set this function to add / modify / etc
// the menu
var CustomAppMenuFunc = (func(m *Menu, win *Window))(nil)

// AddAppMenu adds an "app" menu to the menu -- calls AddStdAppMenu and then
// CustomAppMenuFunc if non-nil
func (m *Menu) AddAppMenu(win *Window) {
	m.AddStdAppMenu(win)
	if CustomAppMenuFunc != nil {
		CustomAppMenuFunc(m, win)
	}
}

// AddStdAppMenu adds a standard set of menu items for application-level control.
func (m *Menu) AddStdAppMenu(win *Window) {
	aboutitle := "About " + goosi.TheApp.Name()
	m.AddAction(ActOpts{Label: aboutitle},
		nil, func(recv, send ki.Ki, sig int64, data any) {
			PromptDialog(win.Scene, DlgOpts{Title: aboutitle, Prompt: goosi.TheApp.About()}, AddOk, NoCancel, nil, nil)
		})
	m.AddAction(ActOpts{Label: "GoGi Preferences...", Shortcut: "Command+P"},
		nil, func(recv, send ki.Ki, sig int64, data any) {
			TheViewIFace.PrefsView(&Prefs)
		})
	m.AddSeparator("sepq")
	m.AddAction(ActOpts{Label: "Quit", Shortcut: "Command+Q"},
		nil, func(recv, send ki.Ki, sig int64, data any) {
			goosi.TheApp.QuitReq()
		})
}

// AddWindowsMenu adds menu items for current main and dialog windows.
// must be called under WindowGlobalMu mutex lock!
func (m *Menu) AddWindowsMenu(win *Window) {
	m.AddAction(ActOpts{Label: "Minimize"},
		nil, func(recv, send ki.Ki, sig int64, data any) {
			win.OSWin.Minimize()
		})
	m.AddAction(ActOpts{Label: "Focus Next", ShortcutKey: KeyFunWinFocusNext},
		nil, func(recv, send ki.Ki, sig int64, data any) {
			AllWindows.FocusNext()
		})
	m.AddSeparator("sepa")
	for _, w := range MainWindows {
		if w != nil {
			m.AddAction(ActOpts{Label: w.Title},
				nil, func(recv, send ki.Ki, sig int64, data any) {
					w.OSWin.Raise()
				})
		}
	}
	if len(DialogWindows) > 0 {
		m.AddSeparator("sepw")
		for _, w := range DialogWindows {
			if w != nil {
				m.AddAction(ActOpts{Label: w.Title},
					nil, func(recv, send ki.Ki, sig int64, data any) {
						w.OSWin.Raise()
					})
			}
		}
	}
}

////////////////////////////////////////////////////////////////////////////////////////
// PopupMenu function

// MenuFrameConfigStyles configures the default styles
// for the given pop-up menu frame with the given parent.
// It should be called on menu frames when they are created.
func MenuFrameConfigStyles(frame *Frame) {
	frame.AddStyler(func(w *WidgetBase, s *gist.Style) {
		s.Border.Style.Set(gist.BorderNone)
		s.Border.Radius = gist.BorderRadiusExtraSmall
		s.BackgroundColor.SetSolid(ColorScheme.SurfaceContainer)
		s.BoxShadow = BoxShadow2
	})
}

// MenuMaxHeight is the maximum height of any menu popup panel in units of font height
// scroll bars are enforced beyond that size.
var MenuMaxHeight = 30

// PopupMenu pops up a viewport with a layout that draws the supplied actions
// positions are relative to given viewport -- name is relevant base name to
// which Menu is appended
func PopupMenu(menu Menu, x, y int, parVp *Scene, name string) *Scene {
	win := parVp.Win
	mainSc := win.Scene
	if len(menu) == 0 {
		log.Printf("GoGi PopupMenu: empty menu given\n")
		return nil
	}

	menu.UpdateActions()

	psc := &Scene{}
	psc.Name = name + "Menu"
	psc.Win = win
	psc.Type = VpMenu

	psc.Geom.Pos = image.Point{x, y}
	frame := &psc.Frame
	MenuFrameConfigStyles(frame)
	var focus ki.Ki
	for _, ac := range menu {
		acn, ac := AsWidget(ac)
		if acn != nil {
			frame.AddChild(acn)
			if ac.IsSelected() {
				focus = acn
			}
		}
	}
	frame.ConfigTree(psc)
	frame.SetStyleTree(psc) // sufficient to get sizes
	mainSz := mat32.NewVec2FmPoint(mainSc.Geom.Size)
	frame.LayState.Alloc.Size = mainSz // give it the whole vp initially
	frame.GetSizeTree(psc, 0)          // collect sizes
	psc.Win = nil
	scextra := frame.Style.ScrollBarWidth.Dots
	frame.LayState.Size.Pref.X += scextra // make room for scrollbar..
	vpsz := frame.LayState.Size.Pref.Min(mainSz.MulScalar(2)).ToPoint()
	maxht := int(32 * frame.Style.Font.Face.Metrics.Height)
	vpsz.Y = min(maxht, vpsz.Y)
	x = max(0, x)
	y = max(0, y)
	x = min(x, mainSc.Geom.Size.X-vpsz.X) // fit
	y = min(y, mainSc.Geom.Size.Y-vpsz.Y) // fit
	psc.Resize(vpsz)
	psc.Geom.Pos = image.Point{x, y}
	win.SetNextPopup(psc, focus)
	return psc
}

// TODO: not working; need to get working.
// RecyclePopupMenu reuses the already existing popup to display
// a viewport with a layout that draws the supplied actions
// positions are relative to given viewport -- name is relevant base name to
// which Menu is appended
func RecyclePopupMenu(menu Menu, x, y int, parVp *Scene, name string) *Scene {
	win := parVp.Win
	mainSc := win.Scene
	if len(menu) == 0 {
		log.Printf("GoGi PopupMenu: empty menu given\n")
		return nil
	}

	menu.UpdateActions()

	psc, ok := win.CurPopup()
	if !ok {
		return PopupMenu(menu, x, y, parVp, name)
	}
	// psc.InitName(psc, name+"Menu")
	psc.Win = win
	psc.Type = VpMenu

	psc.Geom.Pos = image.Point{x, y}
	// note: not setting VpFlagPopupDestroyAll -- we keep the menu list intact
	frame := &psc.Frame
	frame.DeleteChildren(ki.NoDestroyKids)
	// frame.Properties().CopyFrom(MenuFrameProps, ki.DeepCopy)
	var focus ki.Ki
	_ = focus
	for _, ac := range menu {
		acn, ac := AsWidget(ac)
		if acn != nil {
			frame.AddChild(acn)
			if ac.IsSelected() {
				focus = acn
			}
		}
	}
	frame.ConfigTree(psc)
	frame.SetStyleTree(psc) // sufficient to get sizes
	mainSz := mat32.NewVec2FmPoint(mainSc.Geom.Size)
	frame.LayState.Alloc.Size = mainSz // give it the whole vp initially
	frame.GetSizeTree(psc, 0)          // collect sizes
	psc.Win = nil
	scextra := frame.Style.ScrollBarWidth.Dots
	frame.LayState.Size.Pref.X += scextra // make room for scrollbar..
	vpsz := frame.LayState.Size.Pref.Min(mainSz.MulScalar(2)).ToPoint()
	maxht := int(32 * frame.Style.Font.Face.Metrics.Height)
	vpsz.Y = min(maxht, vpsz.Y)
	x = max(0, x)
	y = max(0, y)
	x = min(x, mainSc.Geom.Size.X-vpsz.X) // fit
	y = min(y, mainSc.Geom.Size.Y-vpsz.Y) // fit
	psc.Resize(vpsz)
	psc.Geom.Pos = image.Point{x, y}
	psc.SetFullReRender()
	win.SetNextPopup(psc, focus)
	return psc
}

// StringsChooserPopup creates a menu of the strings in the given string
// slice, and calls the given function on receiver when the user selects --
// this is the ActionSig signal, coming from the Action for the given menu
// item -- the name of the Action is the string value, and the data will be
// the index in the slice.  A string equal to curSel will be marked as
// selected.  Location is from the ContextMenuPos of recv node.
func StringsChooserPopup(strs []string, curSel string, recv Widget, fun ki.RecvFunc) *Scene {
	var menu Menu
	for i, it := range strs {
		ac := menu.AddAction(ActOpts{Label: it, Data: i}, recv, fun)
		ac.SetSelected(it == curSel)
	}
	wb := recv.AsWidget()
	pos := recv.ContextMenuPos()
	sc := wb.Sc
	return PopupMenu(menu, pos.X, pos.Y, sc, recv.Name())
}

// SubStringsChooserPopup creates a menu of the sub-strings in the given
// slice of string slices, and calls the given function on receiver when
// the user selects.  This is the ActionSig signal, coming from the Action
// for the given menu item.
// The sub-menu name is the first element of each sub-slice.
// The name of the Action is the string value, and the data is an
// []int{s,i} slice of submenu and item indexes.
// A string of subMenu: item equal to curSel will be marked as selected.
// Location is from the ContextMenuPos of recv node.
func SubStringsChooserPopup(strs [][]string, curSel string, recv Widget, fun ki.RecvFunc) *Scene {
	var menu Menu
	for si, ss := range strs {
		sz := len(ss)
		if sz < 2 {
			continue
		}
		s1 := ss[0]
		sm := menu.AddAction(ActOpts{Label: s1}, nil, nil)
		sm.SetAsMenu()
		for i := 1; i < sz; i++ {
			it := ss[i]
			cnm := s1 + ": " + it
			ac := sm.Menu.AddAction(ActOpts{Label: it, Data: []int{si, i}}, recv, fun)
			ac.SetSelected(cnm == curSel)
		}
	}
	wb := recv.AsWidget()
	pos := recv.ContextMenuPos()
	sc := wb.Sc
	return PopupMenu(menu, pos.X, pos.Y, sc, recv.Name())
}

// StringsInsertFirst inserts the given string at start of a string slice,
// while keeping overall length to given max value
// useful for a "recents" kind of string list
func StringsInsertFirst(strs *[]string, str string, max int) {
	if strs == nil {
		*strs = make([]string, 0, max)
	}
	sz := len(*strs)
	if sz > max {
		*strs = (*strs)[:max]
	}
	if sz >= max {
		copy((*strs)[1:max], (*strs)[0:max-1])
		(*strs)[0] = str
	} else {
		*strs = append(*strs, "")
		if sz > 0 {
			copy((*strs)[1:], (*strs)[0:sz])
		}
		(*strs)[0] = str
	}
}

// StringsInsertFirstUnique inserts the given string at start of a string slice,
// while keeping overall length to given max value.
// if item is already on the list, then it is moved to the top and not re-added (unique items only)
// useful for a "recents" kind of string list
func StringsInsertFirstUnique(strs *[]string, str string, max int) {
	if strs == nil {
		*strs = make([]string, 0, max)
	}
	sz := len(*strs)
	if sz > max {
		*strs = (*strs)[:max]
	}
	for i, s := range *strs {
		if s == str {
			if i == 0 {
				return
			}
			copy((*strs)[1:i+1], (*strs)[0:i])
			(*strs)[0] = str
			return
		}
	}
	if sz >= max {
		copy((*strs)[1:max], (*strs)[0:max-1])
		(*strs)[0] = str
	} else {
		*strs = append(*strs, "")
		if sz > 0 {
			copy((*strs)[1:], (*strs)[0:sz])
		}
		(*strs)[0] = str
	}
}

// StringsDelete deletes item from strings list
func StringsDelete(strs *[]string, str string) {
	for i, s := range *strs {
		if s == str {
			*strs = append((*strs)[:i], (*strs)[i+1:]...)
			return
		}
	}
}

// StringsAppendIfUnique append str to strs if not already in slice
func StringsAppendIfUnique(strs *[]string, str string, max int) {
	if strs == nil {
		*strs = make([]string, 0, max)
	}
	for _, s := range *strs {
		if s == str {
			return
		}
	}
	*strs = append(*strs, str)
}

// StringsAddExtras is a generic function for appending a slice to a slice used to add items to menus
func StringsAddExtras(items *[]string, extras []string) {
	*items = append(*items, extras...)
}

// StringsRemoveExtras is a generic function for removing items of a slice from another slice
func StringsRemoveExtras(items *[]string, extras []string) {
	for _, extra := range extras {
		i := 0
		for _, item := range *items {
			if item != extra {
				(*items)[i] = item
				i++
			}
		}
		*items = (*items)[:i]
	}
}

////////////////////////////////////////////////////////////////////////////////////////
// Separator

// Separator defines a string to indicate a menu separator item
var MenuTextSeparator = "-------------"

// Separator draws a vertical or horizontal line
type Separator struct {
	WidgetBase

	// is this a horizontal separator -- otherwise vertical
	Horiz bool `xml:"horiz" desc:"is this a horizontal separator -- otherwise vertical"`
}

func (sp *Separator) OnInit() {
	// TODO: fix disappearing separator in menu
	sp.AddStyler(func(w *WidgetBase, s *gist.Style) {
		s.Margin.Set()
		s.Padding.Set(units.Px(8*Prefs.DensityMul()), units.Px(0))
		s.AlignV = gist.AlignCenter
		s.AlignH = gist.AlignCenter
		s.Border.Style.Set(gist.BorderSolid)
		s.Border.Width.Set(units.Px(1))
		s.Border.Color.Set(ColorScheme.OutlineVariant)
		s.BackgroundColor.SetSolid(ColorScheme.OutlineVariant)
		if sp.Horiz {
			s.MaxWidth.SetPx(-1)
			s.MinHeight.SetPx(1)
		} else {
			s.MaxHeight.SetPx(-1)
			s.MinWidth.SetPx(1)
		}
	})
}

func (sp *Separator) CopyFieldsFrom(frm any) {
	fr := frm.(*Separator)
	sp.WidgetBase.CopyFieldsFrom(&fr.WidgetBase)
	sp.Horiz = fr.Horiz
}

func (sp *Separator) RenderSeparator(sc *Scene) {
	rs, pc, st := sp.RenderLock(sc)
	defer sp.RenderUnlock(rs)

	pos := sp.LayState.Alloc.Pos.Add(st.EffMargin().Pos())
	sz := sp.LayState.Alloc.Size.Sub(st.EffMargin().Size())

	if !st.BackgroundColor.IsNil() {
		pc.FillBox(rs, pos, sz, &st.BackgroundColor)
	}
	// border-top is standard property for separators in CSS (see https://www.w3schools.com/howto/howto_css_dividers.asp)
	pc.StrokeStyle.Width = st.Border.Width.Top
	pc.StrokeStyle.SetColor(&st.Border.Color.Top)
	if sp.Horiz {
		pc.DrawLine(rs, pos.X, pos.Y+0.5*sz.Y, pos.X+sz.X, pos.Y+0.5*sz.Y)
	} else {
		pc.DrawLine(rs, pos.X+0.5*sz.X, pos.Y, pos.X+0.5*sz.X, pos.Y+sz.Y)
	}
	pc.FillStrokeClear(rs)
}

func (sp *Separator) Render(sc *Scene) {
	if sp.PushBounds(sc) {
		sp.RenderSeparator(sc)
		sp.RenderChildren(sc)
		sp.PopBounds(sc)
	}
}
