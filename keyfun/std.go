// Copyright (c) 2023, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package keyfun

// StdMaps is the original compiled-in set of standard keymaps that have
// the lastest key functions bound to standard key chords.
var StdMaps = Maps{
	{"MacStd", "Standard Mac KeyMap", Map{
		"UpArrow":              MoveUp,
		"Shift+UpArrow":        MoveUp,
		"Meta+UpArrow":         MoveUp,
		"Control+P":            MoveUp,
		"Shift+Control+P":      MoveUp,
		"Meta+Control+P":       MoveUp,
		"DownArrow":            MoveDown,
		"Shift+DownArrow":      MoveDown,
		"Meta+DownArrow":       MoveDown,
		"Control+N":            MoveDown,
		"Shift+Control+N":      MoveDown,
		"Meta+Control+N":       MoveDown,
		"RightArrow":           MoveRight,
		"Shift+RightArrow":     MoveRight,
		"Meta+RightArrow":      End,
		"Control+F":            MoveRight,
		"Shift+Control+F":      MoveRight,
		"Meta+Control+F":       MoveRight,
		"LeftArrow":            MoveLeft,
		"Shift+LeftArrow":      MoveLeft,
		"Meta+LeftArrow":       Home,
		"Control+B":            MoveLeft,
		"Shift+Control+B":      MoveLeft,
		"Meta+Control+B":       MoveLeft,
		"PageUp":               PageUp,
		"Shift+PageUp":         PageUp,
		"Control+UpArrow":      PageUp,
		"Control+U":            PageUp,
		"PageDown":             PageDown,
		"Shift+PageDown":       PageDown,
		"Control+DownArrow":    PageDown,
		"Shift+Control+V":      PageDown,
		"Alt+√":                PageDown,
		"Alt+V":                PageDown,
		"Meta+Home":            DocHome,
		"Shift+Home":           DocHome,
		"Meta+H":               DocHome,
		"Meta+End":             DocEnd,
		"Shift+End":            DocEnd,
		"Meta+L":               DocEnd,
		"Control+RightArrow":   WordRight,
		"Control+LeftArrow":    WordLeft,
		"Alt+RightArrow":       WordRight,
		"Shift+Alt+RightArrow": WordRight,
		"Alt+LeftArrow":        WordLeft,
		"Shift+Alt+LeftArrow":  WordLeft,
		"Home":                 Home,
		"Control+A":            Home,
		"Shift+Control+A":      Home,
		"End":                  End,
		"Control+E":            End,
		"Shift+Control+E":      End,
		"Tab":                  FocusNext,
		"Shift+Tab":            FocusPrev,
		"ReturnEnter":          Enter,
		"KeypadEnter":          Enter,
		"Meta+A":               SelectAll,
		"Control+G":            CancelSelect,
		"Control+Spacebar":     SelectMode,
		"Control+ReturnEnter":  Accept,
		"Escape":               Abort,
		"Backspace":            Backspace,
		"Control+Backspace":    BackspaceWord,
		"Alt+Backspace":        BackspaceWord,
		"Delete":               Delete,
		"Control+Delete":       DeleteWord,
		"Alt+Delete":           DeleteWord,
		"Control+D":            Delete,
		"Control+K":            Kill,
		"Alt+∑":                Copy,
		"Alt+C":                Copy,
		"Meta+C":               Copy,
		"Control+W":            Cut,
		"Meta+X":               Cut,
		"Control+Y":            Paste,
		"Control+V":            Paste,
		"Meta+V":               Paste,
		"Shift+Meta+V":         PasteHist,
		"Alt+D":                Duplicate,
		"Control+T":            Transpose,
		"Alt+T":                TransposeWord,
		"Control+Z":            Undo,
		"Meta+Z":               Undo,
		"Shift+Control+Z":      Redo,
		"Shift+Meta+Z":         Redo,
		"Control+I":            Insert,
		"Control+O":            InsertAfter,
		"Shift+Meta+=":         ZoomIn,
		"Meta+=":               ZoomIn,
		"Meta+-":               ZoomOut,
		"Control+=":            ZoomIn,
		"Shift+Control++":      ZoomIn,
		"Shift+Meta+-":         ZoomOut,
		"Control+-":            ZoomOut,
		"Shift+Control+_":      ZoomOut,
		"Control+Alt+P":        Prefs,
		"F5":                   Refresh,
		"Control+L":            Recenter,
		"Control+.":            Complete,
		"Control+,":            Lookup,
		"Control+S":            Search,
		"Meta+F":               Find,
		"Meta+R":               Replace,
		"Control+J":            Jump,
		"Control+[":            HistPrev,
		"Control+]":            HistNext,
		"Meta+[":               HistPrev,
		"Meta+]":               HistNext,
		"F10":                  Menu,
		"Control+M":            Menu,
		"Meta+`":               WinFocusNext,
		"Meta+W":               WinClose,
		"Control+Alt+G":        WinSnapshot,
		"Shift+Control+G":      WinSnapshot,
		"Control+Alt+I":        Inspector,
		"Shift+Control+I":      Inspector,
		"Meta+N":               New,
		"Shift+Meta+N":         NewAlt1,
		"Alt+Meta+N":           NewAlt2,
		"Meta+O":               Open,
		"Shift+Meta+O":         OpenAlt1,
		"Alt+Meta+O":           OpenAlt2,
		"Meta+S":               Save,
		"Shift+Meta+S":         SaveAs,
		"Alt+Meta+S":           SaveAlt,
		"Shift+Meta+W":         CloseAlt1,
		"Alt+Meta+W":           CloseAlt2,
	}},
	{"MacEmacs", "Mac with emacs-style navigation -- emacs wins in conflicts", Map{
		"UpArrow":              MoveUp,
		"Shift+UpArrow":        MoveUp,
		"Meta+UpArrow":         MoveUp,
		"Control+P":            MoveUp,
		"Shift+Control+P":      MoveUp,
		"Meta+Control+P":       MoveUp,
		"DownArrow":            MoveDown,
		"Shift+DownArrow":      MoveDown,
		"Meta+DownArrow":       MoveDown,
		"Control+N":            MoveDown,
		"Shift+Control+N":      MoveDown,
		"Meta+Control+N":       MoveDown,
		"RightArrow":           MoveRight,
		"Shift+RightArrow":     MoveRight,
		"Meta+RightArrow":      End,
		"Control+F":            MoveRight,
		"Shift+Control+F":      MoveRight,
		"Meta+Control+F":       MoveRight,
		"LeftArrow":            MoveLeft,
		"Shift+LeftArrow":      MoveLeft,
		"Meta+LeftArrow":       Home,
		"Control+B":            MoveLeft,
		"Shift+Control+B":      MoveLeft,
		"Meta+Control+B":       MoveLeft,
		"PageUp":               PageUp,
		"Shift+PageUp":         PageUp,
		"Control+UpArrow":      PageUp,
		"Control+U":            PageUp,
		"PageDown":             PageDown,
		"Shift+PageDown":       PageDown,
		"Control+DownArrow":    PageDown,
		"Shift+Control+V":      PageDown,
		"Alt+√":                PageDown,
		"Alt+V":                PageDown,
		"Control+V":            PageDown,
		"Control+RightArrow":   WordRight,
		"Control+LeftArrow":    WordLeft,
		"Alt+RightArrow":       WordRight,
		"Shift+Alt+RightArrow": WordRight,
		"Alt+LeftArrow":        WordLeft,
		"Shift+Alt+LeftArrow":  WordLeft,
		"Home":                 Home,
		"Control+A":            Home,
		"Shift+Control+A":      Home,
		"End":                  End,
		"Control+E":            End,
		"Shift+Control+E":      End,
		"Meta+Home":            DocHome,
		"Shift+Home":           DocHome,
		"Meta+H":               DocHome,
		"Control+H":            DocHome,
		"Control+Alt+A":        DocHome,
		"Meta+End":             DocEnd,
		"Shift+End":            DocEnd,
		"Meta+L":               DocEnd,
		"Control+Alt+E":        DocEnd,
		"Alt+Ƒ":                WordRight,
		"Alt+F":                WordRight,
		"Alt+∫":                WordLeft,
		"Alt+B":                WordLeft,
		"Tab":                  FocusNext,
		"Shift+Tab":            FocusPrev,
		"ReturnEnter":          Enter,
		"KeypadEnter":          Enter,
		"Meta+A":               SelectAll,
		"Control+G":            CancelSelect,
		"Control+Spacebar":     SelectMode,
		"Control+ReturnEnter":  Accept,
		"Escape":               Abort,
		"Backspace":            Backspace,
		"Control+Backspace":    BackspaceWord,
		"Alt+Backspace":        BackspaceWord,
		"Delete":               Delete,
		"Control+Delete":       DeleteWord,
		"Alt+Delete":           DeleteWord,
		"Control+D":            Delete,
		"Control+K":            Kill,
		"Alt+∑":                Copy,
		"Alt+C":                Copy,
		"Meta+C":               Copy,
		"Control+W":            Cut,
		"Meta+X":               Cut,
		"Control+Y":            Paste,
		"Meta+V":               Paste,
		"Shift+Meta+V":         PasteHist,
		"Shift+Control+Y":      PasteHist,
		"Alt+∂":                Duplicate,
		"Alt+D":                Duplicate,
		"Control+T":            Transpose,
		"Alt+T":                TransposeWord,
		"Control+Z":            Undo,
		"Meta+Z":               Undo,
		"Control+/":            Undo,
		"Shift+Control+Z":      Redo,
		"Shift+Meta+Z":         Redo,
		"Control+I":            Insert,
		"Control+O":            InsertAfter,
		"Shift+Meta+=":         ZoomIn,
		"Meta+=":               ZoomIn,
		"Meta+-":               ZoomOut,
		"Control+=":            ZoomIn,
		"Shift+Control++":      ZoomIn,
		"Shift+Meta+-":         ZoomOut,
		"Control+-":            ZoomOut,
		"Shift+Control+_":      ZoomOut,
		"Control+Alt+P":        Prefs,
		"F5":                   Refresh,
		"Control+L":            Recenter,
		"Control+.":            Complete,
		"Control+,":            Lookup,
		"Control+S":            Search,
		"Meta+F":               Find,
		"Meta+R":               Replace,
		"Control+R":            Replace,
		"Control+J":            Jump,
		"Control+[":            HistPrev,
		"Control+]":            HistNext,
		"Meta+[":               HistPrev,
		"Meta+]":               HistNext,
		"F10":                  Menu,
		"Control+M":            Menu,
		"Meta+`":               WinFocusNext,
		"Meta+W":               WinClose,
		"Control+Alt+G":        WinSnapshot,
		"Shift+Control+G":      WinSnapshot,
		"Control+Alt+I":        Inspector,
		"Shift+Control+I":      Inspector,
		"Meta+N":               New,
		"Shift+Meta+N":         NewAlt1,
		"Alt+Meta+N":           NewAlt2,
		"Meta+O":               Open,
		"Shift+Meta+O":         OpenAlt1,
		"Alt+Meta+O":           OpenAlt2,
		"Meta+S":               Save,
		"Shift+Meta+S":         SaveAs,
		"Alt+Meta+S":           SaveAlt,
		"Shift+Meta+W":         CloseAlt1,
		"Alt+Meta+W":           CloseAlt2,
	}},
	{"LinuxEmacs", "Linux with emacs-style navigation -- emacs wins in conflicts", Map{
		"UpArrow":             MoveUp,
		"Shift+UpArrow":       MoveUp,
		"Alt+UpArrow":         MoveUp,
		"Control+P":           MoveUp,
		"Shift+Control+P":     MoveUp,
		"Alt+Control+P":       MoveUp,
		"DownArrow":           MoveDown,
		"Shift+DownArrow":     MoveDown,
		"Alt+DownArrow":       MoveDown,
		"Control+N":           MoveDown,
		"Shift+Control+N":     MoveDown,
		"Alt+Control+N":       MoveDown,
		"RightArrow":          MoveRight,
		"Shift+RightArrow":    MoveRight,
		"Alt+RightArrow":      End,
		"Control+F":           MoveRight,
		"Shift+Control+F":     MoveRight,
		"Alt+Control+F":       MoveRight,
		"LeftArrow":           MoveLeft,
		"Shift+LeftArrow":     MoveLeft,
		"Alt+LeftArrow":       Home,
		"Control+B":           MoveLeft,
		"Shift+Control+B":     MoveLeft,
		"Alt+Control+B":       MoveLeft,
		"PageUp":              PageUp,
		"Shift+PageUp":        PageUp,
		"Control+UpArrow":     PageUp,
		"Control+U":           PageUp,
		"Shift+Control+U":     PageUp,
		"Alt+Control+U":       PageUp,
		"PageDown":            PageDown,
		"Shift+PageDown":      PageDown,
		"Control+DownArrow":   PageDown,
		"Control+V":           PageDown,
		"Shift+Control+V":     PageDown,
		"Alt+Control+V":       PageDown,
		"Alt+Home":            DocHome,
		"Shift+Home":          DocHome,
		"Alt+H":               DocHome,
		"Control+Alt+A":       DocHome,
		"Alt+End":             DocEnd,
		"Shift+End":           DocEnd,
		"Alt+L":               DocEnd,
		"Control+Alt+E":       DocEnd,
		"Control+RightArrow":  WordRight,
		"Control+LeftArrow":   WordLeft,
		"Home":                Home,
		"Control+A":           Home,
		"Shift+Control+A":     Home,
		"End":                 End,
		"Control+E":           End,
		"Shift+Control+E":     End,
		"Tab":                 FocusNext,
		"Shift+Tab":           FocusPrev,
		"ReturnEnter":         Enter,
		"KeypadEnter":         Enter,
		"Alt+A":               SelectAll,
		"Control+G":           CancelSelect,
		"Control+Spacebar":    SelectMode,
		"Control+ReturnEnter": Accept,
		"Escape":              Abort,
		"Backspace":           Backspace,
		"Control+Backspace":   BackspaceWord,
		"Delete":              Delete,
		"Control+D":           Delete,
		"Control+Delete":      DeleteWord,
		"Alt+Delete":          DeleteWord,
		"Control+K":           Kill,
		"Alt+W":               Copy,
		"Alt+C":               Copy,
		"Control+W":           Cut,
		"Alt+X":               Cut,
		"Control+Y":           Paste,
		"Alt+V":               Paste,
		"Shift+Alt+V":         PasteHist,
		"Shift+Control+Y":     PasteHist,
		"Alt+D":               Duplicate,
		"Control+T":           Transpose,
		"Alt+T":               TransposeWord,
		"Control+Z":           Undo,
		"Control+/":           Undo,
		"Shift+Control+Z":     Redo,
		"Control+I":           Insert,
		"Control+O":           InsertAfter,
		"Control+=":           ZoomIn,
		"Shift+Control++":     ZoomIn,
		"Control+-":           ZoomOut,
		"Shift+Control+_":     ZoomOut,
		"Control+Alt+P":       Prefs,
		"F5":                  Refresh,
		"Control+L":           Recenter,
		"Control+.":           Complete,
		"Control+,":           Lookup,
		"Control+S":           Search,
		"Alt+F":               Find,
		"Control+R":           Replace,
		"Control+J":           Jump,
		"Control+[":           HistPrev,
		"Control+]":           HistNext,
		"F10":                 Menu,
		"Control+M":           Menu,
		"Alt+F6":              WinFocusNext,
		"Shift+Control+W":     WinClose,
		"Control+Alt+G":       WinSnapshot,
		"Shift+Control+G":     WinSnapshot,
		"Control+Alt+I":       Inspector,
		"Shift+Control+I":     Inspector,
		"Alt+N":               New, // ctrl keys conflict..
		"Shift+Alt+N":         NewAlt1,
		"Control+Alt+N":       NewAlt2,
		"Alt+O":               Open,
		"Shift+Alt+O":         OpenAlt1,
		"Control+Alt+O":       OpenAlt2,
		"Alt+S":               Save,
		"Shift+Alt+S":         SaveAs,
		"Control+Alt+S":       SaveAlt,
		"Shift+Alt+W":         CloseAlt1,
		"Control+Alt+W":       CloseAlt2,
	}},
	{"LinuxStd", "Standard Linux KeyMap", Map{
		"UpArrow":             MoveUp,
		"Shift+UpArrow":       MoveUp,
		"DownArrow":           MoveDown,
		"Shift+DownArrow":     MoveDown,
		"RightArrow":          MoveRight,
		"Shift+RightArrow":    MoveRight,
		"LeftArrow":           MoveLeft,
		"Shift+LeftArrow":     MoveLeft,
		"PageUp":              PageUp,
		"Shift+PageUp":        PageUp,
		"Control+UpArrow":     PageUp,
		"PageDown":            PageDown,
		"Shift+PageDown":      PageDown,
		"Control+DownArrow":   PageDown,
		"Home":                Home,
		"Alt+LeftArrow":       Home,
		"End":                 End,
		"Alt+Home":            DocHome,
		"Shift+Home":          DocHome,
		"Alt+End":             DocEnd,
		"Shift+End":           DocEnd,
		"Control+RightArrow":  WordRight,
		"Control+LeftArrow":   WordLeft,
		"Alt+RightArrow":      End,
		"Tab":                 FocusNext,
		"Shift+Tab":           FocusPrev,
		"ReturnEnter":         Enter,
		"KeypadEnter":         Enter,
		"Control+A":           SelectAll,
		"Shift+Control+A":     CancelSelect,
		"Control+G":           CancelSelect,
		"Control+Spacebar":    SelectMode, // change input method / keyboard
		"Control+ReturnEnter": Accept,
		"Escape":              Abort,
		"Backspace":           Backspace,
		"Control+Backspace":   BackspaceWord,
		"Delete":              Delete,
		"Control+Delete":      DeleteWord,
		"Alt+Delete":          DeleteWord,
		"Control+K":           Kill,
		"Control+C":           Copy,
		"Control+X":           Cut,
		"Control+V":           Paste,
		"Shift+Control+V":     PasteHist,
		"Alt+D":               Duplicate,
		"Control+T":           Transpose,
		"Alt+T":               TransposeWord,
		"Control+Z":           Undo,
		"Control+Y":           Redo,
		"Shift+Control+Z":     Redo,
		"Control+Alt+I":       Insert,
		"Control+Alt+O":       InsertAfter,
		"Control+=":           ZoomIn,
		"Shift+Control++":     ZoomIn,
		"Control+-":           ZoomOut,
		"Shift+Control+_":     ZoomOut,
		"Shift+Control+P":     Prefs,
		"Control+Alt+P":       Prefs,
		"F5":                  Refresh,
		"Control+L":           Recenter,
		"Control+.":           Complete,
		"Control+,":           Lookup,
		"Alt+S":               Search,
		"Control+F":           Find,
		"Control+H":           Replace,
		"Control+R":           Replace,
		"Control+J":           Jump,
		"Control+[":           HistPrev,
		"Control+]":           HistNext,
		"Control+N":           New,
		"F10":                 Menu,
		"Control+M":           Menu,
		"Alt+F6":              WinFocusNext,
		"Control+W":           WinClose,
		"Control+Alt+G":       WinSnapshot,
		"Shift+Control+G":     WinSnapshot,
		"Shift+Control+I":     Inspector,
		"Shift+Control+N":     NewAlt1,
		"Control+Alt+N":       NewAlt2,
		"Control+O":           Open,
		"Shift+Control+O":     OpenAlt1,
		"Shift+Alt+O":         OpenAlt2,
		"Control+S":           Save,
		"Shift+Control+S":     SaveAs,
		"Control+Alt+S":       SaveAlt,
		"Shift+Control+W":     CloseAlt1,
		"Control+Alt+W":       CloseAlt2,
	}},
	{"WindowsStd", "Standard Windows KeyMap", Map{
		"UpArrow":             MoveUp,
		"Shift+UpArrow":       MoveUp,
		"DownArrow":           MoveDown,
		"Shift+DownArrow":     MoveDown,
		"RightArrow":          MoveRight,
		"Shift+RightArrow":    MoveRight,
		"LeftArrow":           MoveLeft,
		"Shift+LeftArrow":     MoveLeft,
		"PageUp":              PageUp,
		"Shift+PageUp":        PageUp,
		"Control+UpArrow":     PageUp,
		"PageDown":            PageDown,
		"Shift+PageDown":      PageDown,
		"Control+DownArrow":   PageDown,
		"Home":                Home,
		"Alt+LeftArrow":       Home,
		"End":                 End,
		"Alt+RightArrow":      End,
		"Alt+Home":            DocHome,
		"Shift+Home":          DocHome,
		"Alt+End":             DocEnd,
		"Shift+End":           DocEnd,
		"Control+RightArrow":  WordRight,
		"Control+LeftArrow":   WordLeft,
		"Tab":                 FocusNext,
		"Shift+Tab":           FocusPrev,
		"ReturnEnter":         Enter,
		"KeypadEnter":         Enter,
		"Control+A":           SelectAll,
		"Shift+Control+A":     CancelSelect,
		"Control+G":           CancelSelect,
		"Control+Spacebar":    SelectMode, // change input method / keyboard
		"Control+ReturnEnter": Accept,
		"Escape":              Abort,
		"Backspace":           Backspace,
		"Control+Backspace":   BackspaceWord,
		"Delete":              Delete,
		"Control+Delete":      DeleteWord,
		"Alt+Delete":          DeleteWord,
		"Control+K":           Kill,
		"Control+C":           Copy,
		"Control+X":           Cut,
		"Control+V":           Paste,
		"Shift+Control+V":     PasteHist,
		"Alt+D":               Duplicate,
		"Control+T":           Transpose,
		"Alt+T":               TransposeWord,
		"Control+Z":           Undo,
		"Control+Y":           Redo,
		"Shift+Control+Z":     Redo,
		"Control+Alt+I":       Insert,
		"Control+Alt+O":       InsertAfter,
		"Control+=":           ZoomIn,
		"Shift+Control++":     ZoomIn,
		"Control+-":           ZoomOut,
		"Shift+Control+_":     ZoomOut,
		"Shift+Control+P":     Prefs,
		"Control+Alt+P":       Prefs,
		"F5":                  Refresh,
		"Control+L":           Recenter,
		"Control+.":           Complete,
		"Control+,":           Lookup,
		"Alt+S":               Search,
		"Control+F":           Find,
		"Control+H":           Replace,
		"Control+R":           Replace,
		"Control+J":           Jump,
		"Control+[":           HistPrev,
		"Control+]":           HistNext,
		"F10":                 Menu,
		"Control+M":           Menu,
		"Alt+F6":              WinFocusNext,
		"Control+W":           WinClose,
		"Control+Alt+G":       WinSnapshot,
		"Shift+Control+G":     WinSnapshot,
		"Shift+Control+I":     Inspector,
		"Control+N":           New,
		"Shift+Control+N":     NewAlt1,
		"Control+Alt+N":       NewAlt2,
		"Control+O":           Open,
		"Shift+Control+O":     OpenAlt1,
		"Shift+Alt+O":         OpenAlt2,
		"Control+S":           Save,
		"Shift+Control+S":     SaveAs,
		"Control+Alt+S":       SaveAlt,
		"Shift+Control+W":     CloseAlt1,
		"Control+Alt+W":       CloseAlt2,
	}},
	{"ChromeStd", "Standard chrome-browser and linux-under-chrome bindings", Map{
		"UpArrow":             MoveUp,
		"Shift+UpArrow":       MoveUp,
		"DownArrow":           MoveDown,
		"Shift+DownArrow":     MoveDown,
		"RightArrow":          MoveRight,
		"Shift+RightArrow":    MoveRight,
		"LeftArrow":           MoveLeft,
		"Shift+LeftArrow":     MoveLeft,
		"PageUp":              PageUp,
		"Shift+PageUp":        PageUp,
		"Control+UpArrow":     PageUp,
		"PageDown":            PageDown,
		"Shift+PageDown":      PageDown,
		"Control+DownArrow":   PageDown,
		"Home":                Home,
		"Alt+LeftArrow":       Home,
		"End":                 End,
		"Alt+Home":            DocHome,
		"Shift+Home":          DocHome,
		"Alt+End":             DocEnd,
		"Shift+End":           DocEnd,
		"Control+RightArrow":  WordRight,
		"Control+LeftArrow":   WordLeft,
		"Alt+RightArrow":      End,
		"Tab":                 FocusNext,
		"Shift+Tab":           FocusPrev,
		"ReturnEnter":         Enter,
		"KeypadEnter":         Enter,
		"Control+A":           SelectAll,
		"Shift+Control+A":     CancelSelect,
		"Control+G":           CancelSelect,
		"Control+Spacebar":    SelectMode, // change input method / keyboard
		"Control+ReturnEnter": Accept,
		"Escape":              Abort,
		"Backspace":           Backspace,
		"Control+Backspace":   BackspaceWord,
		"Delete":              Delete,
		"Control+Delete":      DeleteWord,
		"Alt+Delete":          DeleteWord,
		"Control+K":           Kill,
		"Control+C":           Copy,
		"Control+X":           Cut,
		"Control+V":           Paste,
		"Shift+Control+V":     PasteHist,
		"Alt+D":               Duplicate,
		"Control+T":           Transpose,
		"Alt+T":               TransposeWord,
		"Control+Z":           Undo,
		"Control+Y":           Redo,
		"Shift+Control+Z":     Redo,
		"Control+Alt+I":       Insert,
		"Control+Alt+O":       InsertAfter,
		"Control+=":           ZoomIn,
		"Shift+Control++":     ZoomIn,
		"Control+-":           ZoomOut,
		"Shift+Control+_":     ZoomOut,
		"Shift+Control+P":     Prefs,
		"Control+Alt+P":       Prefs,
		"F5":                  Refresh,
		"Control+L":           Recenter,
		"Control+.":           Complete,
		"Control+,":           Lookup,
		"Alt+S":               Search,
		"Control+F":           Find,
		"Control+H":           Replace,
		"Control+R":           Replace,
		"Control+J":           Jump,
		"Control+[":           HistPrev,
		"Control+]":           HistNext,
		"F10":                 Menu,
		"Control+M":           Menu,
		"Alt+F6":              WinFocusNext,
		"Control+W":           WinClose,
		"Control+Alt+G":       WinSnapshot,
		"Shift+Control+G":     WinSnapshot,
		"Shift+Control+I":     Inspector,
		"Control+N":           New,
		"Shift+Control+N":     NewAlt1,
		"Control+Alt+N":       NewAlt2,
		"Control+O":           Open,
		"Shift+Control+O":     OpenAlt1,
		"Shift+Alt+O":         OpenAlt2,
		"Control+S":           Save,
		"Shift+Control+S":     SaveAs,
		"Control+Alt+S":       SaveAlt,
		"Shift+Control+W":     CloseAlt1,
		"Control+Alt+W":       CloseAlt2,
	}},
}
