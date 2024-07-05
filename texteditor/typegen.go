// Code generated by "core generate"; DO NOT EDIT.

package texteditor

import (
	"image"

	"cogentcore.org/core/core"
	"cogentcore.org/core/paint"
	"cogentcore.org/core/styles/units"
	"cogentcore.org/core/texteditor/textbuf"
	"cogentcore.org/core/tree"
	"cogentcore.org/core/types"
)

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/texteditor.Spell", IDName: "spell", Doc: "Spell has all the texteditor spell check state", Directives: []types.Directive{{Tool: "types", Directive: "add", Args: []string{"-setters"}}}, Fields: []types.Field{{Name: "SrcLn", Doc: "line number in source that spelling is operating on, if relevant"}, {Name: "SrcCh", Doc: "character position in source that spelling is operating on (start of word to be corrected)"}, {Name: "Suggest", Doc: "list of suggested corrections"}, {Name: "Word", Doc: "word being checked"}, {Name: "LastLearned", Doc: "last word learned -- can be undone -- stored in lowercase format"}, {Name: "Correction", Doc: "the user's correction selection"}, {Name: "Listeners", Doc: "the event listeners for the spell (it sends Select events)"}, {Name: "Stage", Doc: "Stage is the [PopupStage] associated with the [Spell]"}, {Name: "ShowMu"}}})

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

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/texteditor.Buffer", IDName: "buffer", Doc: "Buffer is a buffer of text, which can be viewed by [Editor](s).\nIt holds the raw text lines (in original string and rune formats,\nand marked-up from syntax highlighting), and sends signals for making\nedits to the text and coordinating those edits across multiple views.\nViews always only view a single buffer, so they directly call methods\non the buffer to drive updates, which are then broadcast.\nIt also has methods for loading and saving buffers to files.\nUnlike GUI Widgets, its methods generally send events, without an\nexplicit Action suffix.\nInternally, the buffer represents new lines using \\n = LF, but saving\nand loading can deal with Windows/DOS CRLF format.", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Methods: []types.Method{{Name: "Open", Doc: "Open loads the given file into the buffer.", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Args: []string{"filename"}, Returns: []string{"error"}}, {Name: "Revert", Doc: "Revert re-opens text from current file, if filename set -- returns false if\nnot -- uses an optimized diff-based update to preserve existing formatting\n-- very fast if not very different", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Returns: []string{"bool"}}, {Name: "SaveAs", Doc: "SaveAs saves the current text into given file -- does an EditDone first to save edits\nand checks for an existing file -- if it does exist then prompts to overwrite or not.", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Args: []string{"filename"}}, {Name: "Save", Doc: "Save saves the current text into current Filename associated with this\nbuffer", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Returns: []string{"error"}}}, Fields: []types.Field{{Name: "Filename", Doc: "Filename is the filename of the file that was last loaded or saved. It is used when highlighting code."}, {Name: "Txt", Doc: "Txt is the current value of the entire text being edited, represented as a byte slice for efficiency."}, {Name: "Autosave", Doc: "Autosave specifies whether the file should be automatically saved after changes are made."}, {Name: "Options", Doc: "Options are the options for how text editing and viewing works."}, {Name: "Info", Doc: "Info is the full information about the file."}, {Name: "ParseState", Doc: "ParseState is the parsing state information for the file."}, {Name: "Hi", Doc: "Hi is the syntax highlighting markup parameters, such as the language and style."}, {Name: "NLines", Doc: "NLines is the number of lines in the buffer."}, {Name: "LineColors", Doc: "LineColors are the colors to use for rendering circles next to the line numbers of certain lines."}, {Name: "Lines", Doc: "Lines are the live lines of text being edited, with the latest modifications. They are encoded as runes per line, which is necessary for one-to-one rune/glyph rendering correspondence. All TextPos positions are in rune indexes, not byte indexes."}, {Name: "LineBytes", Doc: "LineBytes are the live lines of text being edited, with the latest modifications. They are encoded in bytes per line, translated from Lines, and used for input to markup. It is essential to use Lines and not LineBytes when dealing with TextPos positions, which are in runes."}, {Name: "Tags", Doc: "Tags are the extra custom tagged regions for each line."}, {Name: "HiTags", Doc: "HiTags are the syntax highlighting tags, which are auto-generated."}, {Name: "Markup", Doc: "Markup is the marked-up version of the edited text lines, after being run through the syntax highlighting process. This is what is actually rendered."}, {Name: "MarkupEdits", Doc: "MarkupEdits are the edits that have been made since the last full markup."}, {Name: "ByteOffs", Doc: "ByteOffs are the offsets for the start of each line in the Txt byte slice. This is not updated with edits. Call SetByteOffs to set it when needed. It is used for re-generating the Txt in LinesToBytes and set on initial open in BytesToLines."}, {Name: "TotalBytes", Doc: "TotalBytes is the total number of bytes in the document. See ByteOffs for when it is updated."}, {Name: "LinesMu", Doc: "LinesMu is the mutex for updating lines."}, {Name: "MarkupMu", Doc: "MarkupMu is the mutex for updating markup."}, {Name: "MarkupDelayTimer", Doc: "MarkupDelayTimer is the markup delay timer."}, {Name: "MarkupDelayMu", Doc: "MarkupDelayMu is the mutex for updating the markup delay timer."}, {Name: "Editors", Doc: "Editors are the editors that are currently viewing this buffer."}, {Name: "Undos", Doc: "Undos is the undo manager."}, {Name: "PosHistory", Doc: "PosHistory is the history of cursor positions. It can be used to move back through them."}, {Name: "Complete", Doc: "Complete is the functions and data for text completion."}, {Name: "Spell", Doc: "Spell is the functions and data for spelling correction."}, {Name: "CurrentEditor", Doc: "CurrentEditor is the current text editor, such as the one that initiated the Complete or Correct process. The cursor position in this view is updated, and it is reset to nil after usage."}, {Name: "Listeners", Doc: "Listeners is used for sending standard system events. Change is sent for BufDone, BufInsert, and BufDelete."}, {Name: "autoSaving", Doc: "autoSaving is used in atomically safe way to protect autosaving"}, {Name: "markingUp", Doc: "markingUp indicates current markup operation in progress -- don't redo"}, {Name: "Changed", Doc: "Changed indicates if the text has been Changed (edited) relative to the\noriginal, since last EditDone"}, {Name: "NotSaved", Doc: "NotSaved indicates if the text has been changed (edited) relative to the\noriginal, since last Save"}, {Name: "fileModOK", Doc: "fileModOK have already asked about fact that file has changed since being\nopened, user is ok"}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/texteditor.DiffEditor", IDName: "diff-editor", Doc: "DiffEditor presents two side-by-side [Editor]s showing the differences\nbetween two files (represented as lines of strings).", Methods: []types.Method{{Name: "SaveFileA", Doc: "SaveFileA saves the current state of file A to given filename", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Args: []string{"fname"}}, {Name: "SaveFileB", Doc: "SaveFileB saves the current state of file B to given filename", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Args: []string{"fname"}}}, Embeds: []types.Field{{Name: "Frame"}}, Fields: []types.Field{{Name: "FileA", Doc: "first file name being compared"}, {Name: "FileB", Doc: "second file name being compared"}, {Name: "RevA", Doc: "revision for first file, if relevant"}, {Name: "RevB", Doc: "revision for second file, if relevant"}, {Name: "BufA", Doc: "textbuf for A showing the aligned edit view"}, {Name: "BufB", Doc: "textbuf for B showing the aligned edit view"}, {Name: "AlignD", Doc: "aligned diffs records diff for aligned lines"}, {Name: "Diffs", Doc: "Diffs applied"}, {Name: "inInputEvent"}}})

// NewDiffEditor returns a new [DiffEditor] with the given optional parent:
// DiffEditor presents two side-by-side [Editor]s showing the differences
// between two files (represented as lines of strings).
func NewDiffEditor(parent ...tree.Node) *DiffEditor { return tree.New[DiffEditor](parent...) }

// SetFileA sets the [DiffEditor.FileA]:
// first file name being compared
func (t *DiffEditor) SetFileA(v string) *DiffEditor { t.FileA = v; return t }

// SetFileB sets the [DiffEditor.FileB]:
// second file name being compared
func (t *DiffEditor) SetFileB(v string) *DiffEditor { t.FileB = v; return t }

// SetRevA sets the [DiffEditor.RevA]:
// revision for first file, if relevant
func (t *DiffEditor) SetRevA(v string) *DiffEditor { t.RevA = v; return t }

// SetRevB sets the [DiffEditor.RevB]:
// revision for second file, if relevant
func (t *DiffEditor) SetRevB(v string) *DiffEditor { t.RevB = v; return t }

// SetDiffs sets the [DiffEditor.Diffs]:
// Diffs applied
func (t *DiffEditor) SetDiffs(v textbuf.DiffSelected) *DiffEditor { t.Diffs = v; return t }

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/texteditor.DiffTextEditor", IDName: "diff-text-editor", Doc: "DiffTextEditor supports double-click based application of edits from one\nbuffer to the other.", Embeds: []types.Field{{Name: "Editor"}}})

// NewDiffTextEditor returns a new [DiffTextEditor] with the given optional parent:
// DiffTextEditor supports double-click based application of edits from one
// buffer to the other.
func NewDiffTextEditor(parent ...tree.Node) *DiffTextEditor {
	return tree.New[DiffTextEditor](parent...)
}

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/texteditor.Editor", IDName: "editor", Doc: "Editor is a widget for editing multiple lines of complicated text (as compared to\n[core.TextField] for a single line of simple text).  The Editor is driven by a [Buffer]\nbuffer which contains all the text, and manages all the edits,\nsending update events out to the editors.\n\nUse NeedsRender to drive an render update for any change that does\nnot change the line-level layout of the text.\nUse NeedsLayout whenever there are changes across lines that require\nre-layout of the text.  This sets the Widget NeedsRender flag and triggers\nlayout during that render.\n\nMultiple editors can be attached to a given buffer.  All updating in the\nEditor should be within a single goroutine, as it would require\nextensive protections throughout code otherwise.", Directives: []types.Directive{{Tool: "core", Directive: "embedder"}}, Methods: []types.Method{{Name: "Lookup", Doc: "Lookup attempts to lookup symbol at current location, popping up a window\nif something is found.", Directives: []types.Directive{{Tool: "types", Directive: "add"}}}}, Embeds: []types.Field{{Name: "Frame"}}, Fields: []types.Field{{Name: "Buffer", Doc: "Buffer is the text buffer being edited."}, {Name: "CursorWidth", Doc: "CursorWidth is the width of the cursor."}, {Name: "LineNumberColor", Doc: "LineNumberColor is the color used for the side bar containing the line numbers.\nThis should be set in Stylers like all other style properties."}, {Name: "SelectColor", Doc: "SelectColor is the color used for the user text selection background color.\nThis should be set in Stylers like all other style properties."}, {Name: "HighlightColor", Doc: "HighlightColor is the color used for the text highlight background color (like in find).\nThis should be set in Stylers like all other style properties."}, {Name: "CursorColor", Doc: "CursorColor is the color used for the text editor cursor bar.\nThis should be set in Stylers like all other style properties."}, {Name: "NLines", Doc: "NLines is the number of lines in the view, synced with the Buf after edits,\nbut always reflects the storage size of Renders etc."}, {Name: "Renders", Doc: "Renders is a slice of paint.Text representing the renders of the text lines,\nwith one render per line (each line could visibly wrap-around, so these are logical lines, not display lines)."}, {Name: "Offsets", Doc: "Offsets is a slice of float32 representing the starting render offsets for the top of each line."}, {Name: "LineNumberDigits", Doc: "LineNumberDigits is the number of line number digits needed."}, {Name: "LineNumberOffset", Doc: "LineNumberOffset is the horizontal offset for the start of text after line numbers."}, {Name: "LineNumberRender", Doc: "LineNumberRender is the render for line numbers."}, {Name: "CursorPos", Doc: "CursorPos is the current cursor position."}, {Name: "CursorTarget", Doc: "CursorTarget is the target cursor position for externally set targets.\nIt ensures that the target position is visible."}, {Name: "CursorCol", Doc: "CursorCol is the desired cursor column, where the cursor was last when moved using left / right arrows.\nIt is used when doing up / down to not always go to short line columns."}, {Name: "PosHistIndex", Doc: "PosHistIndex is the current index within PosHistory."}, {Name: "SelectStart", Doc: "SelectStart is the starting point for selection, which will either be the start or end of selected region\ndepending on subsequent selection."}, {Name: "SelectRegion", Doc: "SelectRegion is the current selection region."}, {Name: "PreviousSelectRegion", Doc: "PreviousSelectRegion is the previous selection region that was actually rendered.\nIt is needed to update the render."}, {Name: "Highlights", Doc: "Highlights is a slice of regions representing the highlighted regions, e.g., for search results."}, {Name: "Scopelights", Doc: "Scopelights is a slice of regions representing the highlighted regions specific to scope markers."}, {Name: "LinkHandler", Doc: "LinkHandler handles link clicks.\nIf it is nil, they are sent to the standard web URL handler."}, {Name: "ISearch", Doc: "ISearch is the interactive search data."}, {Name: "QReplace", Doc: "QReplace is the query replace data."}, {Name: "selectMode", Doc: "selectMode is a boolean indicating whether to select text as the cursor moves."}, {Name: "fontHeight", Doc: "fontHeight is the font height, cached during styling."}, {Name: "lineHeight", Doc: "lineHeight is the line height, cached during styling."}, {Name: "fontAscent", Doc: "fontAscent is the font ascent, cached during styling."}, {Name: "fontDescent", Doc: "fontDescent is the font descent, cached during styling."}, {Name: "nLinesChars", Doc: "nLinesChars is the height in lines and width in chars of the visible area."}, {Name: "linesSize", Doc: "linesSize is the total size of all lines as rendered."}, {Name: "totalSize", Doc: "totalSize is the LinesSize plus extra space and line numbers etc."}, {Name: "lineLayoutSize", Doc: "lineLayoutSize is the Geom.Size.Actual.Total subtracting extra space and line numbers.\nThis is what LayoutStdLR sees for laying out each line."}, {Name: "lastlineLayoutSize", Doc: "lastlineLayoutSize is the last LineLayoutSize used in laying out lines.\nIt is used to trigger a new layout only when needed."}, {Name: "blinkOn", Doc: "blinkOn oscillates between on and off for blinking."}, {Name: "cursorMu", Doc: "cursorMu is a mutex protecting cursor rendering, shared between blink and main code."}, {Name: "hasLinks", Doc: "hasLinks is a boolean indicating if at least one of the renders has links.\nIt determines if we set the cursor for hand movements."}, {Name: "hasLineNumbers", Doc: "hasLineNumbers indicates that this editor has line numbers\n(per [Buffer] option)"}, {Name: "needsLayout", Doc: "needsLayout is set by NeedsLayout: Editor does significant\ninternal layout in LayoutAllLines, and its layout is simply based\non what it gets allocated, so it does not affect the rest\nof the Scene."}, {Name: "lastWasTabAI", Doc: "lastWasTabAI indicates that last key was a Tab auto-indent"}, {Name: "lastWasUndo", Doc: "lastWasUndo indicates that last key was an undo"}, {Name: "targetSet", Doc: "targetSet indicates that the CursorTarget is set"}, {Name: "lastRecenter"}, {Name: "lastAutoInsert"}, {Name: "lastFilename"}}})

// NewEditor returns a new [Editor] with the given optional parent:
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
func NewEditor(parent ...tree.Node) *Editor { return tree.New[Editor](parent...) }

// EditorEmbedder is an interface that all types that embed Editor satisfy
type EditorEmbedder interface {
	AsEditor() *Editor
}

// AsEditor returns the given value as a value of type Editor if the type
// of the given value embeds Editor, or nil otherwise
func AsEditor(n tree.Node) *Editor {
	if t, ok := n.(EditorEmbedder); ok {
		return t.AsEditor()
	}
	return nil
}

// AsEditor satisfies the [EditorEmbedder] interface
func (t *Editor) AsEditor() *Editor { return t }

// SetCursorWidth sets the [Editor.CursorWidth]:
// CursorWidth is the width of the cursor.
func (t *Editor) SetCursorWidth(v units.Value) *Editor { t.CursorWidth = v; return t }

// SetLineNumberColor sets the [Editor.LineNumberColor]:
// LineNumberColor is the color used for the side bar containing the line numbers.
// This should be set in Stylers like all other style properties.
func (t *Editor) SetLineNumberColor(v image.Image) *Editor { t.LineNumberColor = v; return t }

// SetSelectColor sets the [Editor.SelectColor]:
// SelectColor is the color used for the user text selection background color.
// This should be set in Stylers like all other style properties.
func (t *Editor) SetSelectColor(v image.Image) *Editor { t.SelectColor = v; return t }

// SetHighlightColor sets the [Editor.HighlightColor]:
// HighlightColor is the color used for the text highlight background color (like in find).
// This should be set in Stylers like all other style properties.
func (t *Editor) SetHighlightColor(v image.Image) *Editor { t.HighlightColor = v; return t }

// SetCursorColor sets the [Editor.CursorColor]:
// CursorColor is the color used for the text editor cursor bar.
// This should be set in Stylers like all other style properties.
func (t *Editor) SetCursorColor(v image.Image) *Editor { t.CursorColor = v; return t }

// SetLinkHandler sets the [Editor.LinkHandler]:
// LinkHandler handles link clicks.
// If it is nil, they are sent to the standard web URL handler.
func (t *Editor) SetLinkHandler(v func(tl *paint.TextLink)) *Editor { t.LinkHandler = v; return t }

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/texteditor.TwinEditors", IDName: "twin-editors", Doc: "TwinEditors presents two side-by-side [Editor]s in [core.Splits]\nthat scroll in sync with each other.", Embeds: []types.Field{{Name: "Splits"}}, Fields: []types.Field{{Name: "BufferA", Doc: "[Buffer] for A"}, {Name: "BufferB", Doc: "[Buffer] for B"}, {Name: "inInputEvent"}}})

// NewTwinEditors returns a new [TwinEditors] with the given optional parent:
// TwinEditors presents two side-by-side [Editor]s in [core.Splits]
// that scroll in sync with each other.
func NewTwinEditors(parent ...tree.Node) *TwinEditors { return tree.New[TwinEditors](parent...) }

// SetBufferA sets the [TwinEditors.BufferA]:
// [Buffer] for A
func (t *TwinEditors) SetBufferA(v *Buffer) *TwinEditors { t.BufferA = v; return t }

// SetBufferB sets the [TwinEditors.BufferB]:
// [Buffer] for B
func (t *TwinEditors) SetBufferB(v *Buffer) *TwinEditors { t.BufferB = v; return t }
