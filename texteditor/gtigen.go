// Code generated by "goki generate ./..."; DO NOT EDIT.

package texteditor

import (
	"goki.dev/colors"
	"goki.dev/gi/v2/gi"
	"goki.dev/girl/units"
	"goki.dev/gti"
	"goki.dev/ki/v2"
	"goki.dev/mat32/v2"
	"goki.dev/ordmap"
)

// EditorType is the [gti.Type] for [Editor]
var EditorType = gti.AddType(&gti.Type{
	Name:      "goki.dev/gi/v2/texteditor.Editor",
	ShortName: "texteditor.Editor",
	IDName:    "editor",
	Doc:       "Editor is a widget for editing multiple lines of text (as compared to\n[gi.TextField] for a single line).  The Editor is driven by a Buf buffer which\ncontains all the text, and manages all the edits, sending update signals\nout to the views -- multiple views can be attached to a given buffer.  All\nupdating in the Editor should be within a single goroutine -- it would\nrequire extensive protections throughout code otherwise.",
	Directives: gti.Directives{
		&gti.Directive{Tool: "goki", Directive: "embedder", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Buf", &gti.Field{Name: "Buf", Type: "*goki.dev/gi/v2/texteditor.Buf", LocalType: "*Buf", Doc: "the text buffer that we're editing", Directives: gti.Directives{}, Tag: "set:\"-\" json:\"-\" xml:\"-\""}},
		{"Placeholder", &gti.Field{Name: "Placeholder", Type: "string", LocalType: "string", Doc: "text that is displayed when the field is empty, in a lower-contrast manner", Directives: gti.Directives{}, Tag: "json:\"-\" xml:\"placeholder\""}},
		{"CursorWidth", &gti.Field{Name: "CursorWidth", Type: "goki.dev/girl/units.Value", LocalType: "units.Value", Doc: "width of cursor -- set from cursor-width property (inherited)", Directives: gti.Directives{}, Tag: "xml:\"cursor-width\""}},
		{"LineNumberColor", &gti.Field{Name: "LineNumberColor", Type: "goki.dev/colors.Full", LocalType: "colors.Full", Doc: "the color used for the side bar containing the line numbers; this should be set in Stylers like all other style properties", Directives: gti.Directives{}, Tag: ""}},
		{"SelectColor", &gti.Field{Name: "SelectColor", Type: "goki.dev/colors.Full", LocalType: "colors.Full", Doc: "the color used for the user text selection background color; this should be set in Stylers like all other style properties", Directives: gti.Directives{}, Tag: ""}},
		{"HighlightColor", &gti.Field{Name: "HighlightColor", Type: "goki.dev/colors.Full", LocalType: "colors.Full", Doc: "the color used for the text highlight background color (like in find); this should be set in Stylers like all other style properties", Directives: gti.Directives{}, Tag: ""}},
		{"CursorColor", &gti.Field{Name: "CursorColor", Type: "goki.dev/colors.Full", LocalType: "colors.Full", Doc: "the color used for the text field cursor (caret); this should be set in Stylers like all other style properties", Directives: gti.Directives{}, Tag: ""}},
		{"NLines", &gti.Field{Name: "NLines", Type: "int", LocalType: "int", Doc: "number of lines in the view -- sync'd with the Buf after edits, but always reflects storage size of Renders etc", Directives: gti.Directives{}, Tag: "set:\"-\" view:\"-\" json:\"-\" xml:\"-\""}},
		{"Renders", &gti.Field{Name: "Renders", Type: "[]goki.dev/girl/paint.Text", LocalType: "[]paint.Text", Doc: "renders of the text lines, with one render per line (each line could visibly wrap-around, so these are logical lines, not display lines)", Directives: gti.Directives{}, Tag: "set:\"-\" view:\"-\" json:\"-\" xml:\"-\""}},
		{"Offs", &gti.Field{Name: "Offs", Type: "[]float32", LocalType: "[]float32", Doc: "starting render offsets for top of each line", Directives: gti.Directives{}, Tag: "set:\"-\" view:\"-\" json:\"-\" xml:\"-\""}},
		{"LineNoDigs", &gti.Field{Name: "LineNoDigs", Type: "int", LocalType: "int", Doc: "number of line number digits needed", Directives: gti.Directives{}, Tag: "set:\"-\" view:\"-\" json:\"-\" xml:\"-\""}},
		{"LineNoOff", &gti.Field{Name: "LineNoOff", Type: "float32", LocalType: "float32", Doc: "horizontal offset for start of text after line numbers", Directives: gti.Directives{}, Tag: "set:\"-\" view:\"-\" json:\"-\" xml:\"-\""}},
		{"LineNoRender", &gti.Field{Name: "LineNoRender", Type: "goki.dev/girl/paint.Text", LocalType: "paint.Text", Doc: "render for line numbers", Directives: gti.Directives{}, Tag: "set:\"-\" view:\"-\" json:\"-\" xml:\"-\""}},
		{"CursorPos", &gti.Field{Name: "CursorPos", Type: "goki.dev/pi/v2/lex.Pos", LocalType: "lex.Pos", Doc: "current cursor position", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"CursorCol", &gti.Field{Name: "CursorCol", Type: "int", LocalType: "int", Doc: "desired cursor column -- where the cursor was last when moved using left / right arrows -- used when doing up / down to not always go to short line columns", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"ScrollToCursorOnRender", &gti.Field{Name: "ScrollToCursorOnRender", Type: "bool", LocalType: "bool", Doc: "if true, scroll screen to cursor on next render", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"ScrollToCursorPos", &gti.Field{Name: "ScrollToCursorPos", Type: "goki.dev/pi/v2/lex.Pos", LocalType: "lex.Pos", Doc: "cursor position to scroll to", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"PosHistIdx", &gti.Field{Name: "PosHistIdx", Type: "int", LocalType: "int", Doc: "current index within PosHistory", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"SelectStart", &gti.Field{Name: "SelectStart", Type: "goki.dev/pi/v2/lex.Pos", LocalType: "lex.Pos", Doc: "starting point for selection -- will either be the start or end of selected region depending on subsequent selection.", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"SelectReg", &gti.Field{Name: "SelectReg", Type: "goki.dev/gi/v2/texteditor/textbuf.Region", LocalType: "textbuf.Region", Doc: "current selection region", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"PrevSelectReg", &gti.Field{Name: "PrevSelectReg", Type: "goki.dev/gi/v2/texteditor/textbuf.Region", LocalType: "textbuf.Region", Doc: "previous selection region, that was actually rendered -- needed to update render", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"Highlights", &gti.Field{Name: "Highlights", Type: "[]goki.dev/gi/v2/texteditor/textbuf.Region", LocalType: "[]textbuf.Region", Doc: "highlighted regions, e.g., for search results", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"Scopelights", &gti.Field{Name: "Scopelights", Type: "[]goki.dev/gi/v2/texteditor/textbuf.Region", LocalType: "[]textbuf.Region", Doc: "highlighted regions, specific to scope markers", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"SelectMode", &gti.Field{Name: "SelectMode", Type: "bool", LocalType: "bool", Doc: "if true, select text as cursor moves", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"ForceComplete", &gti.Field{Name: "ForceComplete", Type: "bool", LocalType: "bool", Doc: "if true, complete regardless of any disqualifying reasons", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"ISearch", &gti.Field{Name: "ISearch", Type: "goki.dev/gi/v2/texteditor.ISearch", LocalType: "ISearch", Doc: "interactive search data", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"QReplace", &gti.Field{Name: "QReplace", Type: "goki.dev/gi/v2/texteditor.QReplace", LocalType: "QReplace", Doc: "query replace data", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"FontHeight", &gti.Field{Name: "FontHeight", Type: "float32", LocalType: "float32", Doc: "font height, cached during styling", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"LineHeight", &gti.Field{Name: "LineHeight", Type: "float32", LocalType: "float32", Doc: "line height, cached during styling", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"NLinesChars", &gti.Field{Name: "NLinesChars", Type: "image.Point", LocalType: "image.Point", Doc: "height in lines and width in chars of the visible area", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"LinesSize", &gti.Field{Name: "LinesSize", Type: "goki.dev/mat32/v2.Vec2", LocalType: "mat32.Vec2", Doc: "total size of all lines as rendered", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"TotalSize", &gti.Field{Name: "TotalSize", Type: "goki.dev/mat32/v2.Vec2", LocalType: "mat32.Vec2", Doc: "TotalSize = LinesSize plus extra space and line numbers etc", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"LineLayoutSize", &gti.Field{Name: "LineLayoutSize", Type: "goki.dev/mat32/v2.Vec2", LocalType: "mat32.Vec2", Doc: "LineLayoutSize is Alloc.Size.Total subtracting\nextra space and line numbers -- this is what\nLayoutStdLR sees for laying out each line", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"BlinkOn", &gti.Field{Name: "BlinkOn", Type: "bool", LocalType: "bool", Doc: "oscillates between on and off for blinking", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"CursorMu", &gti.Field{Name: "CursorMu", Type: "sync.Mutex", LocalType: "sync.Mutex", Doc: "mutex protecting cursor rendering -- shared between blink and main code", Directives: gti.Directives{}, Tag: "set:\"-\" json:\"-\" xml:\"-\" view:\"-\""}},
		{"HasLinks", &gti.Field{Name: "HasLinks", Type: "bool", LocalType: "bool", Doc: "at least one of the renders has links -- determines if we set the cursor for hand movements", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"lastRecenter", &gti.Field{Name: "lastRecenter", Type: "int", LocalType: "int", Doc: "", Directives: gti.Directives{}, Tag: "set:\"-\""}},
		{"lastAutoInsert", &gti.Field{Name: "lastAutoInsert", Type: "rune", LocalType: "rune", Doc: "", Directives: gti.Directives{}, Tag: "set:\"-\""}},
		{"lastFilename", &gti.Field{Name: "lastFilename", Type: "goki.dev/gi/v2/gi.FileName", LocalType: "gi.FileName", Doc: "", Directives: gti.Directives{}, Tag: "set:\"-\""}},
	}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Layout", &gti.Field{Name: "Layout", Type: "goki.dev/gi/v2/gi.Layout", LocalType: "gi.Layout", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods:  ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
	Instance: &Editor{},
})

// NewEditor adds a new [Editor] with the given name
// to the given parent. If the name is unspecified, it defaults
// to the ID (kebab-case) name of the type, plus the
// [ki.Ki.NumLifetimeChildren] of the given parent.
func NewEditor(par ki.Ki, name ...string) *Editor {
	return par.NewChild(EditorType, name...).(*Editor)
}

// KiType returns the [*gti.Type] of [Editor]
func (t *Editor) KiType() *gti.Type {
	return EditorType
}

// New returns a new [*Editor] value
func (t *Editor) New() ki.Ki {
	return &Editor{}
}

// EditorEmbedder is an interface that all types that embed Editor satisfy
type EditorEmbedder interface {
	AsEditor() *Editor
}

// AsEditor returns the given value as a value of type Editor if the type
// of the given value embeds Editor, or nil otherwise
func AsEditor(k ki.Ki) *Editor {
	if k == nil || k.This() == nil {
		return nil
	}
	if t, ok := k.(EditorEmbedder); ok {
		return t.AsEditor()
	}
	return nil
}

// AsEditor satisfies the [EditorEmbedder] interface
func (t *Editor) AsEditor() *Editor {
	return t
}

// SetPlaceholder sets the [Editor.Placeholder]:
// text that is displayed when the field is empty, in a lower-contrast manner
func (t *Editor) SetPlaceholder(v string) *Editor {
	t.Placeholder = v
	return t
}

// SetCursorWidth sets the [Editor.CursorWidth]:
// width of cursor -- set from cursor-width property (inherited)
func (t *Editor) SetCursorWidth(v units.Value) *Editor {
	t.CursorWidth = v
	return t
}

// SetLineNumberColor sets the [Editor.LineNumberColor]:
// the color used for the side bar containing the line numbers; this should be set in Stylers like all other style properties
func (t *Editor) SetLineNumberColor(v colors.Full) *Editor {
	t.LineNumberColor = v
	return t
}

// SetSelectColor sets the [Editor.SelectColor]:
// the color used for the user text selection background color; this should be set in Stylers like all other style properties
func (t *Editor) SetSelectColor(v colors.Full) *Editor {
	t.SelectColor = v
	return t
}

// SetHighlightColor sets the [Editor.HighlightColor]:
// the color used for the text highlight background color (like in find); this should be set in Stylers like all other style properties
func (t *Editor) SetHighlightColor(v colors.Full) *Editor {
	t.HighlightColor = v
	return t
}

// SetCursorColor sets the [Editor.CursorColor]:
// the color used for the text field cursor (caret); this should be set in Stylers like all other style properties
func (t *Editor) SetCursorColor(v colors.Full) *Editor {
	t.CursorColor = v
	return t
}

// SetTooltip sets the [Editor.Tooltip]
func (t *Editor) SetTooltip(v string) *Editor {
	t.Tooltip = v
	return t
}

// SetClass sets the [Editor.Class]
func (t *Editor) SetClass(v string) *Editor {
	t.Class = v
	return t
}

// SetCustomContextMenu sets the [Editor.CustomContextMenu]
func (t *Editor) SetCustomContextMenu(v func(m *gi.Scene)) *Editor {
	t.CustomContextMenu = v
	return t
}

// SetLayout sets the [Editor.Lay]
func (t *Editor) SetLayout(v gi.Layouts) *Editor {
	t.Lay = v
	return t
}

// SetStackTop sets the [Editor.StackTop]
func (t *Editor) SetStackTop(v int) *Editor {
	t.StackTop = v
	return t
}

// TwinEditorsType is the [gti.Type] for [TwinEditors]
var TwinEditorsType = gti.AddType(&gti.Type{
	Name:       "goki.dev/gi/v2/texteditor.TwinEditors",
	ShortName:  "texteditor.TwinEditors",
	IDName:     "twin-editors",
	Doc:        "TwinEditors presents two side-by-side [Editor]s in [gi.Splits]\nthat scroll in sync with each other.",
	Directives: gti.Directives{},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"BufA", &gti.Field{Name: "BufA", Type: "*goki.dev/gi/v2/texteditor.Buf", LocalType: "*Buf", Doc: "textbuf for A", Directives: gti.Directives{}, Tag: "json:\"-\" xml:\"-\""}},
		{"BufB", &gti.Field{Name: "BufB", Type: "*goki.dev/gi/v2/texteditor.Buf", LocalType: "*Buf", Doc: "textbuf for B", Directives: gti.Directives{}, Tag: "json:\"-\" xml:\"-\""}},
	}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Splits", &gti.Field{Name: "Splits", Type: "goki.dev/gi/v2/gi.Splits", LocalType: "gi.Splits", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods:  ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
	Instance: &TwinEditors{},
})

// NewTwinEditors adds a new [TwinEditors] with the given name
// to the given parent. If the name is unspecified, it defaults
// to the ID (kebab-case) name of the type, plus the
// [ki.Ki.NumLifetimeChildren] of the given parent.
func NewTwinEditors(par ki.Ki, name ...string) *TwinEditors {
	return par.NewChild(TwinEditorsType, name...).(*TwinEditors)
}

// KiType returns the [*gti.Type] of [TwinEditors]
func (t *TwinEditors) KiType() *gti.Type {
	return TwinEditorsType
}

// New returns a new [*TwinEditors] value
func (t *TwinEditors) New() ki.Ki {
	return &TwinEditors{}
}

// SetBufA sets the [TwinEditors.BufA]:
// textbuf for A
func (t *TwinEditors) SetBufA(v *Buf) *TwinEditors {
	t.BufA = v
	return t
}

// SetBufB sets the [TwinEditors.BufB]:
// textbuf for B
func (t *TwinEditors) SetBufB(v *Buf) *TwinEditors {
	t.BufB = v
	return t
}

// SetTooltip sets the [TwinEditors.Tooltip]
func (t *TwinEditors) SetTooltip(v string) *TwinEditors {
	t.Tooltip = v
	return t
}

// SetClass sets the [TwinEditors.Class]
func (t *TwinEditors) SetClass(v string) *TwinEditors {
	t.Class = v
	return t
}

// SetCustomContextMenu sets the [TwinEditors.CustomContextMenu]
func (t *TwinEditors) SetCustomContextMenu(v func(m *gi.Scene)) *TwinEditors {
	t.CustomContextMenu = v
	return t
}

// SetDim sets the [TwinEditors.Dim]
func (t *TwinEditors) SetDim(v mat32.Dims) *TwinEditors {
	t.Dim = v
	return t
}
