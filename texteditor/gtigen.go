// Code generated by "core generate"; DO NOT EDIT.

package texteditor

import (
	"image"

	"cogentcore.org/core/core"
	"cogentcore.org/core/gti"
	"cogentcore.org/core/tree"
	"cogentcore.org/core/paint"
	"cogentcore.org/core/texteditor/textbuf"
	"cogentcore.org/core/units"
)

var _ = gti.AddType(&gti.Type{Name: "cogentcore.org/core/texteditor.Spell", IDName: "spell", Doc: "Spell", Directives: []gti.Directive{{Tool: "gti", Directive: "add", Args: []string{"-setters"}}}, Fields: []gti.Field{{Name: "SrcLn", Doc: "line number in source that spelling is operating on, if relevant"}, {Name: "SrcCh", Doc: "character position in source that spelling is operating on (start of word to be corrected)"}, {Name: "Suggest", Doc: "list of suggested corrections"}, {Name: "Word", Doc: "word being checked"}, {Name: "LastLearned", Doc: "last word learned -- can be undone -- stored in lowercase format"}, {Name: "Correction", Doc: "the user's correction selection"}, {Name: "Listeners", Doc: "the event listeners for the spell (it sends Select events)"}, {Name: "Stage", Doc: "Stage is the [PopupStage] associated with the [Spell]"}, {Name: "ShowMu"}}})

// SetSrcLn sets the [Spell.SrcLn]:
// line number in source that spelling is operating on, if relevant
func (t *Spell) SetSrcLn(v int) *Spell { t.SrcLn = v; return t }

// SetSrcCh sets the [Spell.SrcCh]:
// character position in source that spelling is operating on (start of word to be corrected)
func (t *Spell) SetSrcCh(v int) *Spell { t.SrcCh = v; return t }

// SetSuggest sets the [Spell.Suggest]:
// list of suggested corrections
func (t *Spell) SetSuggest(v ...string) *Spell { t.Suggest = v; return t }

// SetStage sets the [Spell.Stage]:
// Stage is the [PopupStage] associated with the [Spell]
func (t *Spell) SetStage(v *core.Stage) *Spell { t.Stage = v; return t }

// DiffViewType is the [gti.Type] for [DiffView]
var DiffViewType = gti.AddType(&gti.Type{Name: "cogentcore.org/core/texteditor.DiffView", IDName: "diff-view", Doc: "DiffView presents two side-by-side TextEditor windows showing the differences\nbetween two files (represented as lines of strings).", Methods: []gti.Method{{Name: "SaveFileA", Doc: "SaveFileA saves the current state of file A to given filename", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Args: []string{"fname"}}, {Name: "SaveFileB", Doc: "SaveFileB saves the current state of file B to given filename", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Args: []string{"fname"}}}, Embeds: []gti.Field{{Name: "Frame"}}, Fields: []gti.Field{{Name: "FileA", Doc: "first file name being compared"}, {Name: "FileB", Doc: "second file name being compared"}, {Name: "RevA", Doc: "revision for first file, if relevant"}, {Name: "RevB", Doc: "revision for second file, if relevant"}, {Name: "BufA", Doc: "textbuf for A showing the aligned edit view"}, {Name: "BufB", Doc: "textbuf for B showing the aligned edit view"}, {Name: "AlignD", Doc: "aligned diffs records diff for aligned lines"}, {Name: "Diffs", Doc: "Diffs applied"}}, Instance: &DiffView{}})

// NewDiffView adds a new [DiffView] with the given name to the given parent:
// DiffView presents two side-by-side TextEditor windows showing the differences
// between two files (represented as lines of strings).
func NewDiffView(parent tree.Node, name ...string) *DiffView {
	return parent.NewChild(DiffViewType, name...).(*DiffView)
}

// NodeType returns the [*gti.Type] of [DiffView]
func (t *DiffView) NodeType() *gti.Type { return DiffViewType }

// New returns a new [*DiffView] value
func (t *DiffView) New() tree.Node { return &DiffView{} }

// SetFileA sets the [DiffView.FileA]:
// first file name being compared
func (t *DiffView) SetFileA(v string) *DiffView { t.FileA = v; return t }

// SetFileB sets the [DiffView.FileB]:
// second file name being compared
func (t *DiffView) SetFileB(v string) *DiffView { t.FileB = v; return t }

// SetRevA sets the [DiffView.RevA]:
// revision for first file, if relevant
func (t *DiffView) SetRevA(v string) *DiffView { t.RevA = v; return t }

// SetRevB sets the [DiffView.RevB]:
// revision for second file, if relevant
func (t *DiffView) SetRevB(v string) *DiffView { t.RevB = v; return t }

// SetDiffs sets the [DiffView.Diffs]:
// Diffs applied
func (t *DiffView) SetDiffs(v textbuf.DiffSelected) *DiffView { t.Diffs = v; return t }

// SetTooltip sets the [DiffView.Tooltip]
func (t *DiffView) SetTooltip(v string) *DiffView { t.Tooltip = v; return t }

// DiffTextEditorType is the [gti.Type] for [DiffTextEditor]
var DiffTextEditorType = gti.AddType(&gti.Type{Name: "cogentcore.org/core/texteditor.DiffTextEditor", IDName: "diff-text-editor", Doc: "DiffTextEditor supports double-click based application of edits from one\nbuffer to the other.", Embeds: []gti.Field{{Name: "Editor"}}, Instance: &DiffTextEditor{}})

// NewDiffTextEditor adds a new [DiffTextEditor] with the given name to the given parent:
// DiffTextEditor supports double-click based application of edits from one
// buffer to the other.
func NewDiffTextEditor(parent tree.Node, name ...string) *DiffTextEditor {
	return parent.NewChild(DiffTextEditorType, name...).(*DiffTextEditor)
}

// NodeType returns the [*gti.Type] of [DiffTextEditor]
func (t *DiffTextEditor) NodeType() *gti.Type { return DiffTextEditorType }

// New returns a new [*DiffTextEditor] value
func (t *DiffTextEditor) New() tree.Node { return &DiffTextEditor{} }

// SetTooltip sets the [DiffTextEditor.Tooltip]
func (t *DiffTextEditor) SetTooltip(v string) *DiffTextEditor { t.Tooltip = v; return t }

// SetCursorWidth sets the [DiffTextEditor.CursorWidth]
func (t *DiffTextEditor) SetCursorWidth(v units.Value) *DiffTextEditor { t.CursorWidth = v; return t }

// SetLineNumberColor sets the [DiffTextEditor.LineNumberColor]
func (t *DiffTextEditor) SetLineNumberColor(v image.Image) *DiffTextEditor {
	t.LineNumberColor = v
	return t
}

// SetSelectColor sets the [DiffTextEditor.SelectColor]
func (t *DiffTextEditor) SetSelectColor(v image.Image) *DiffTextEditor { t.SelectColor = v; return t }

// SetHighlightColor sets the [DiffTextEditor.HighlightColor]
func (t *DiffTextEditor) SetHighlightColor(v image.Image) *DiffTextEditor {
	t.HighlightColor = v
	return t
}

// SetCursorColor sets the [DiffTextEditor.CursorColor]
func (t *DiffTextEditor) SetCursorColor(v image.Image) *DiffTextEditor { t.CursorColor = v; return t }

// SetLinkHandler sets the [DiffTextEditor.LinkHandler]
func (t *DiffTextEditor) SetLinkHandler(v func(tl *paint.TextLink)) *DiffTextEditor {
	t.LinkHandler = v
	return t
}

// EditorType is the [gti.Type] for [Editor]
var EditorType = gti.AddType(&gti.Type{Name: "cogentcore.org/core/texteditor.Editor", IDName: "editor", Doc: "Editor is a widget for editing multiple lines of complicated text (as compared to\n[gi.TextField] for a single line of simple text).  The Editor is driven by a [Buffer]\nbuffer which contains all the text, and manages all the edits,\nsending update events out to the editors.\n\nUse NeedsRender to drive an render update for any change that does\nnot change the line-level layout of the text.\nUse NeedsLayout whenever there are changes across lines that require\nre-layout of the text.  This sets the Widget NeedsRender flag and triggers\nlayout during that render.\n\nMultiple editors can be attached to a given buffer.  All updating in the\nEditor should be within a single goroutine, as it would require\nextensive protections throughout code otherwise.", Directives: []gti.Directive{{Tool: "core", Directive: "embedder"}}, Methods: []gti.Method{{Name: "Lookup", Doc: "Lookup attempts to lookup symbol at current location, popping up a window\nif something is found.", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}}}, Embeds: []gti.Field{{Name: "Layout"}}, Fields: []gti.Field{{Name: "Buffer", Doc: "Buffer is the text buffer being edited."}, {Name: "CursorWidth", Doc: "width of cursor -- set from cursor-width property (inherited)"}, {Name: "LineNumberColor", Doc: "the color used for the side bar containing the line numbers; this should be set in Stylers like all other style properties"}, {Name: "SelectColor", Doc: "the color used for the user text selection background color; this should be set in Stylers like all other style properties"}, {Name: "HighlightColor", Doc: "the color used for the text highlight background color (like in find); this should be set in Stylers like all other style properties"}, {Name: "CursorColor", Doc: "the color used for the text field cursor (caret); this should be set in Stylers like all other style properties"}, {Name: "NLines", Doc: "number of lines in the view -- sync'd with the Buf after edits, but always reflects storage size of Renders etc"}, {Name: "Renders", Doc: "renders of the text lines, with one render per line (each line could visibly wrap-around, so these are logical lines, not display lines)"}, {Name: "Offs", Doc: "starting render offsets for top of each line"}, {Name: "LineNoDigs", Doc: "number of line number digits needed"}, {Name: "LineNoOff", Doc: "horizontal offset for start of text after line numbers"}, {Name: "LineNoRender", Doc: "render for line numbers"}, {Name: "CursorPos", Doc: "current cursor position"}, {Name: "CursorTarg", Doc: "target cursor position for externally-set targets: ensures that it is visible"}, {Name: "CursorCol", Doc: "desired cursor column -- where the cursor was last when moved using left / right arrows -- used when doing up / down to not always go to short line columns"}, {Name: "PosHistIndex", Doc: "current index within PosHistory"}, {Name: "SelectStart", Doc: "starting point for selection -- will either be the start or end of selected region depending on subsequent selection."}, {Name: "SelectReg", Doc: "current selection region"}, {Name: "PrevSelectReg", Doc: "previous selection region, that was actually rendered -- needed to update render"}, {Name: "Highlights", Doc: "highlighted regions, e.g., for search results"}, {Name: "Scopelights", Doc: "highlighted regions, specific to scope markers"}, {Name: "SelectMode", Doc: "if true, select text as cursor moves"}, {Name: "ISearch", Doc: "interactive search data"}, {Name: "QReplace", Doc: "query replace data"}, {Name: "FontHeight", Doc: "font height, cached during styling"}, {Name: "LineHeight", Doc: "line height, cached during styling"}, {Name: "FontAscent", Doc: "font ascent, cached during styling"}, {Name: "FontDescent", Doc: "font descent, cached during styling"}, {Name: "NLinesChars", Doc: "height in lines and width in chars of the visible area"}, {Name: "LinesSize", Doc: "total size of all lines as rendered"}, {Name: "TotalSize", Doc: "the LinesSize plus extra space and line numbers etc"}, {Name: "LineLayoutSize", Doc: "the Geom.Size.Actual.Total subtracting\nextra space and line numbers -- this is what\nLayoutStdLR sees for laying out each line"}, {Name: "lastlineLayoutSize", Doc: "the last LineLayoutSize used in laying out lines.\nUsed to trigger a new layout only when needed."}, {Name: "BlinkOn", Doc: "oscillates between on and off for blinking"}, {Name: "CursorMu", Doc: "mutex protecting cursor rendering -- shared between blink and main code"}, {Name: "HasLinks", Doc: "at least one of the renders has links -- determines if we set the cursor for hand movements"}, {Name: "LinkHandler", Doc: "handles link clicks -- if nil, they are sent to the standard web URL handler"}, {Name: "lastRecenter"}, {Name: "lastAutoInsert"}, {Name: "lastFilename"}}, Instance: &Editor{}})

// NewEditor adds a new [Editor] with the given name to the given parent:
// Editor is a widget for editing multiple lines of complicated text (as compared to
// [core.TextField] for a single line of simple text).  The Editor is driven by a [Buffer]
// buffer which contains all the text, and manages all the edits,
// sending update events out to the editors.
//
// Use NeedsRender to drive an render update for any change that does
// not change the line-level layout of the text.
// Use NeedsLayout whenever there are changes across lines that require
// re-layout of the text.  This sets the Widget NeedsRender flag and triggers
// layout during that render.
//
// Multiple editors can be attached to a given buffer.  All updating in the
// Editor should be within a single goroutine, as it would require
// extensive protections throughout code otherwise.
func NewEditor(parent tree.Node, name ...string) *Editor {
	return parent.NewChild(EditorType, name...).(*Editor)
}

// NodeType returns the [*gti.Type] of [Editor]
func (t *Editor) NodeType() *gti.Type { return EditorType }

// New returns a new [*Editor] value
func (t *Editor) New() tree.Node { return &Editor{} }

// EditorEmbedder is an interface that all types that embed Editor satisfy
type EditorEmbedder interface {
	AsEditor() *Editor
}

// AsEditor returns the given value as a value of type Editor if the type
// of the given value embeds Editor, or nil otherwise
func AsEditor(k tree.Node) *Editor {
	if k == nil || k.This() == nil {
		return nil
	}
	if t, ok := k.(EditorEmbedder); ok {
		return t.AsEditor()
	}
	return nil
}

// AsEditor satisfies the [EditorEmbedder] interface
func (t *Editor) AsEditor() *Editor { return t }

// SetCursorWidth sets the [Editor.CursorWidth]:
// width of cursor -- set from cursor-width property (inherited)
func (t *Editor) SetCursorWidth(v units.Value) *Editor { t.CursorWidth = v; return t }

// SetLineNumberColor sets the [Editor.LineNumberColor]:
// the color used for the side bar containing the line numbers; this should be set in Stylers like all other style properties
func (t *Editor) SetLineNumberColor(v image.Image) *Editor { t.LineNumberColor = v; return t }

// SetSelectColor sets the [Editor.SelectColor]:
// the color used for the user text selection background color; this should be set in Stylers like all other style properties
func (t *Editor) SetSelectColor(v image.Image) *Editor { t.SelectColor = v; return t }

// SetHighlightColor sets the [Editor.HighlightColor]:
// the color used for the text highlight background color (like in find); this should be set in Stylers like all other style properties
func (t *Editor) SetHighlightColor(v image.Image) *Editor { t.HighlightColor = v; return t }

// SetCursorColor sets the [Editor.CursorColor]:
// the color used for the text field cursor (caret); this should be set in Stylers like all other style properties
func (t *Editor) SetCursorColor(v image.Image) *Editor { t.CursorColor = v; return t }

// SetLinkHandler sets the [Editor.LinkHandler]:
// handles link clicks -- if nil, they are sent to the standard web URL handler
func (t *Editor) SetLinkHandler(v func(tl *paint.TextLink)) *Editor { t.LinkHandler = v; return t }

// SetTooltip sets the [Editor.Tooltip]
func (t *Editor) SetTooltip(v string) *Editor { t.Tooltip = v; return t }

// TwinEditorsType is the [gti.Type] for [TwinEditors]
var TwinEditorsType = gti.AddType(&gti.Type{Name: "cogentcore.org/core/texteditor.TwinEditors", IDName: "twin-editors", Doc: "TwinEditors presents two side-by-side [Editor]s in [gi.Splits]\nthat scroll in sync with each other.", Embeds: []gti.Field{{Name: "Splits"}}, Fields: []gti.Field{{Name: "BufA", Doc: "textbuf for A"}, {Name: "BufB", Doc: "textbuf for B"}}, Instance: &TwinEditors{}})

// NewTwinEditors adds a new [TwinEditors] with the given name to the given parent:
// TwinEditors presents two side-by-side [Editor]s in [core.Splits]
// that scroll in sync with each other.
func NewTwinEditors(parent tree.Node, name ...string) *TwinEditors {
	return parent.NewChild(TwinEditorsType, name...).(*TwinEditors)
}

// NodeType returns the [*gti.Type] of [TwinEditors]
func (t *TwinEditors) NodeType() *gti.Type { return TwinEditorsType }

// New returns a new [*TwinEditors] value
func (t *TwinEditors) New() tree.Node { return &TwinEditors{} }

// SetBufA sets the [TwinEditors.BufA]:
// textbuf for A
func (t *TwinEditors) SetBufA(v *Buffer) *TwinEditors { t.BufA = v; return t }

// SetBufB sets the [TwinEditors.BufB]:
// textbuf for B
func (t *TwinEditors) SetBufB(v *Buffer) *TwinEditors { t.BufB = v; return t }

// SetTooltip sets the [TwinEditors.Tooltip]
func (t *TwinEditors) SetTooltip(v string) *TwinEditors { t.Tooltip = v; return t }
