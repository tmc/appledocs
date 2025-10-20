// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Text] class.
var (
	TextClass     _TextClass
	TextClassOnce sync.Once
)

func getTextClass() _TextClass {
	TextClassOnce.Do(func() {
		TextClass = _TextClass{objc.GetClass("NSText")}
	})
	return TextClass
}

type _TextClass struct {
	class objc.Class
}

// An interface definition for the [Text] class.
type IText interface {
	IView
	AlignCenter(sender objc.ID)
	AlignLeft(sender objc.ID)
	AlignRight(sender objc.ID)
	ChangeFont(sender objc.ID)
	CheckSpelling(sender objc.ID)
	Copy(sender objc.ID)
	CopyFont(sender objc.ID)
	CopyRuler(sender objc.ID)
	Cut(sender objc.ID)
	Delete(sender objc.ID)
	Paste(sender objc.ID)
	PasteFont(sender objc.ID)
	PasteRuler(sender objc.ID)
	ReadRTFDFromFile(path string) bool
	ReplaceCharactersInRangeWithString(range_ foundation.Range, string string)
	ReplaceCharactersInRangeWithRTF(range_ foundation.Range, rtfData unsafe.Pointer)
	ReplaceCharactersInRangeWithRTFD(range_ foundation.Range, rtfdData unsafe.Pointer)
	RTFFromRange(range_ foundation.Range) unsafe.Pointer
	RTFDFromRange(range_ foundation.Range) unsafe.Pointer
	ScrollRangeToVisible(range_ foundation.Range)
	SelectAll(sender objc.ID)
	SetFontRange(font unsafe.Pointer, range_ foundation.Range)
	SetTextColorRange(color unsafe.Pointer, range_ foundation.Range)
	ShowGuessPanel(sender objc.ID)
	SizeToFit()
	Subscript(sender objc.ID)
	Superscript(sender objc.ID)
	ToggleRuler(sender objc.ID)
	Underline(sender objc.ID)
	Unscript(sender objc.ID)
	WriteRTFDToFileAtomically(path string, flag bool) bool
}

// The most general programmatic interface for objects that manage text.
//
// draws text for user interface objects, provides text editing capabilities, and controls text attributes such as type size, font, and color. initialization creates an instance of a concrete subclass, such as (generically called a text object). In general, you’re more likely to use the subclass, because it extends the interface declared by and provides much more sophisticated functionality than that declared in . AppKit uses text objects wherever text appears in interface objects. For example, a text object draws the title of a window, the commands in a menu, the title of a button, and the items in a browser. Your app can also create text objects for its own purposes.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText
type Text struct {
	View
}

// TextFrom constructs a [Text] from an unsafe.Pointer.
//
// The most general programmatic interface for objects that manage text.
func TextFrom(ptr unsafe.Pointer) Text {
	return Text{
		View: ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _TextClass) Alloc() Text {
	rv := objc.Send[Text](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextClass) New() Text {
	rv := objc.Send[Text](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ Text) Init() Text {
	rv := objc.Send[Text](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ Text) Autorelease() Text {
	rv := objc.Send[Text](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewText creates a new Text instance.
func NewText() Text {
	return getTextClass().New()
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/init(coder:)
func NewTextWithCoder(coder unsafe.Pointer) Text {
	instance := getTextClass().Alloc()
	rv := objc.Send[Text](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/init(frame:)
func NewTextWithFrame(frameRect coregraphics.CGRect) Text {
	instance := getTextClass().Alloc()
	rv := objc.Send[Text](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	rv.Autorelease()
	return rv
}

// This action method applies center alignment to selected paragraphs (or all text if the receiver is a plain text object).
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/alignCenter(_:)
func (t_ Text) AlignCenter(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("alignCenter:"), sender)
}

// This action method applies left alignment to selected paragraphs (or all text if the receiver is a plain text object).
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/alignLeft(_:)
func (t_ Text) AlignLeft(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("alignLeft:"), sender)
}

// This action method applies right alignment to selected paragraphs (or all text if the receiver is a plain text object).
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/alignRight(_:)
func (t_ Text) AlignRight(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("alignRight:"), sender)
}

// This action method changes the font of the selection for a rich text object, or of all text for a plain text object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/changeFont(_:)
func (t_ Text) ChangeFont(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("changeFont:"), sender)
}

// This action method searches for a misspelled word in the receiver’s text.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/checkSpelling(_:)
func (t_ Text) CheckSpelling(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("checkSpelling:"), sender)
}

// This action method copies the selected text onto the general pasteboard, in as many formats as the receiver supports.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/copy(_:)
func (t_ Text) Copy(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("copy:"), sender)
}

// This action method copies the font information for the first character of the selection (or for the insertion point) onto the font pasteboard, as .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/copyFont(_:)
func (t_ Text) CopyFont(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("copyFont:"), sender)
}

// This action method copies the paragraph style information for first selected paragraph onto the ruler pasteboard, as , and expands the selection to paragraph boundaries.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/copyRuler(_:)
func (t_ Text) CopyRuler(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("copyRuler:"), sender)
}

// This action method deletes the selected text and places it onto the general pasteboard, in as many formats as the receiver supports.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/cut(_:)
func (t_ Text) Cut(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("cut:"), sender)
}

// This action method deletes the selected text.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/delete(_:)
func (t_ Text) Delete(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("delete:"), sender)
}

// This action method pastes text from the general pasteboard at the insertion point or over the selection.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/paste(_:)
func (t_ Text) Paste(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("paste:"), sender)
}

// This action method pastes font information from the font pasteboard onto the selected text or insertion point of a rich text object, or over all text of a plain text object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/pasteFont(_:)
func (t_ Text) PasteFont(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("pasteFont:"), sender)
}

// This action method pastes paragraph style information from the ruler pasteboard onto the selected paragraphs of a rich text object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/pasteRuler(_:)
func (t_ Text) PasteRuler(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("pasteRuler:"), sender)
}

// Attempts to read the RTFD file at , returning if successful and if not.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/readRTFD(fromFile:)
func (t_ Text) ReadRTFDFromFile(path string) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("readRTFDFromFile:"), objc.String(path))
	return rv
}

// Replaces the characters in the given range with those in the given string.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/replaceCharacters(in:with:)
func (t_ Text) ReplaceCharactersInRangeWithString(range_ foundation.Range, string string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("replaceCharactersInRange:withString:"), range_, objc.String(string))
}

// Replaces the characters in the given range with RTF text interpreted from the given RTF data.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/replaceCharacters(in:withRTF:)
func (t_ Text) ReplaceCharactersInRangeWithRTF(range_ foundation.Range, rtfData unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("replaceCharactersInRange:withRTF:"), range_, rtfData)
}

// Replaces the characters in the given range with RTFD text interpreted from the given RTFD data.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/replaceCharacters(in:withRTFD:)
func (t_ Text) ReplaceCharactersInRangeWithRTFD(range_ foundation.Range, rtfdData unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("replaceCharactersInRange:withRTFD:"), range_, rtfdData)
}

// Returns an NSData object that contains an RTF stream corresponding to the characters and attributes within , omitting any attachment characters and attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/rtf(from:)
func (t_ Text) RTFFromRange(range_ foundation.Range) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("RTFFromRange:"), range_)
	return rv
}

// Returns an NSData object that contains an RTFD stream corresponding to the characters and attributes within .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/rtfd(from:)
func (t_ Text) RTFDFromRange(range_ foundation.Range) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("RTFDFromRange:"), range_)
	return rv
}

// Scrolls the receiver in its enclosing scroll view so the first characters of are visible.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/scrollRangeToVisible(_:)
func (t_ Text) ScrollRangeToVisible(range_ foundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("scrollRangeToVisible:"), range_)
}

// This action method selects all of the receiver’s text.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/selectAll(_:)
func (t_ Text) SelectAll(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("selectAll:"), sender)
}

// Sets the font of characters within to .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/setFont(_:range:)
func (t_ Text) SetFontRange(font unsafe.Pointer, range_ foundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFont:range:"), font, range_)
}

// Sets the text color of characters within the specified range to the specified color.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/setTextColor(_:range:)
func (t_ Text) SetTextColorRange(color unsafe.Pointer, range_ foundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextColor:range:"), color, range_)
}

// This action method opens the Spelling panel, allowing the user to make a correction during spell checking.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/showGuessPanel(_:)
func (t_ Text) ShowGuessPanel(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("showGuessPanel:"), sender)
}

// Resizes the receiver to fit its text.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/sizeToFit()
func (t_ Text) SizeToFit() {
	objc.Send[objc.ID](t_.ID, objc.Sel("sizeToFit"))
}

// This action method applies a subscript attribute to selected text (or all text if the receiver is a plain text object), lowering its baseline offset by a predefined amount.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/subscript(_:)
func (t_ Text) Subscript(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("subscript:"), sender)
}

// This action method applies a superscript attribute to selected text (or all text if the receiver is a plain text object), raising its baseline offset by a predefined amount.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/superscript(_:)
func (t_ Text) Superscript(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("superscript:"), sender)
}

// This action method shows or hides the ruler, if the receiver is enclosed in a scroll view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/toggleRuler(_:)
func (t_ Text) ToggleRuler(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("toggleRuler:"), sender)
}

// Adds the underline attribute to the selected text attributes if absent; removes the attribute if present.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/underline(_:)
func (t_ Text) Underline(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("underline:"), sender)
}

// This action method removes any superscripting or subscripting from selected text (or all text if the receiver is a plain text object).
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/unscript(_:)
func (t_ Text) Unscript(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("unscript:"), sender)
}

// Writes the receiver’s text as RTF with attachments to a file or directory at .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/writeRTFD(toFile:atomically:)
func (t_ Text) WriteRTFDToFileAtomically(path string, flag bool) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("writeRTFDToFile:atomically:"), objc.String(path), flag)
	return rv
}

// The alignment of all the receiver’s text.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/alignment
func (t_ Text) Alignment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("alignment"))
	return rv
}

// SetAlignment sets the value of the alignment property.
// The alignment of all the receiver’s text.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/alignment
func (t_ Text) SetAlignment(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAlignment:"), value)
}

// The receiver’s background color to a given color.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/backgroundColor
func (t_ Text) BackgroundColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("backgroundColor"))
	return rv
}

// SetBackgroundColor sets the value of the backgroundColor property.
// The receiver’s background color to a given color.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/backgroundColor
func (t_ Text) SetBackgroundColor(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBackgroundColor:"), value)
}

// The initial writing direction used to determine the actual writing direction for text.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/baseWritingDirection
func (t_ Text) BaseWritingDirection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("baseWritingDirection"))
	return rv
}

// SetBaseWritingDirection sets the value of the baseWritingDirection property.
// The initial writing direction used to determine the actual writing direction for text.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/baseWritingDirection
func (t_ Text) SetBaseWritingDirection(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBaseWritingDirection:"), value)
}

// The receiver’s delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/delegate
func (t_ Text) Delegate() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("delegate"))
	return rv
}

// SetDelegate sets the value of the delegate property.
// The receiver’s delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/delegate
func (t_ Text) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean that controls whether the receiver draws its background.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/drawsBackground
func (t_ Text) DrawsBackground() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("drawsBackground"))
	return rv
}

// SetDrawsBackground sets the value of the drawsBackground property.
// A Boolean that controls whether the receiver draws its background.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/drawsBackground
func (t_ Text) SetDrawsBackground(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDrawsBackground:"), value)
}

// The font of all the receiver’s text.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/font
func (t_ Text) Font() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("font"))
	return rv
}

// SetFont sets the value of the font property.
// The font of all the receiver’s text.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/font
func (t_ Text) SetFont(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFont:"), value)
}

// A Boolean that controls whether the receiver allows the user to import files by dragging.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/importsGraphics
func (t_ Text) ImportsGraphics() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("importsGraphics"))
	return rv
}

// SetImportsGraphics sets the value of the importsGraphics property.
// A Boolean that controls whether the receiver allows the user to import files by dragging.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/importsGraphics
func (t_ Text) SetImportsGraphics(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setImportsGraphics:"), value)
}

// A Boolean that controls whether the receiver allows the user to edit its text.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/isEditable
func (t_ Text) Editable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("editable"))
	return rv
}

// SetEditable sets the value of the editable property.
// A Boolean that controls whether the receiver allows the user to edit its text.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/isEditable
func (t_ Text) SetEditable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setEditable:"), value)
}

// A Boolean that controls whether the receiver interprets Tab, Shift-Tab, and Return (Enter) as cues to end editing and possibly to change the first responder.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/isFieldEditor
func (t_ Text) FieldEditor() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("fieldEditor"))
	return rv
}

// SetFieldEditor sets the value of the fieldEditor property.
// A Boolean that controls whether the receiver interprets Tab, Shift-Tab, and Return (Enter) as cues to end editing and possibly to change the first responder.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/isFieldEditor
func (t_ Text) SetFieldEditor(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFieldEditor:"), value)
}

// A Boolean that controls whether the receiver changes its width to fit the width of its text.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/isHorizontallyResizable
func (t_ Text) HorizontallyResizable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("horizontallyResizable"))
	return rv
}

// SetHorizontallyResizable sets the value of the horizontallyResizable property.
// A Boolean that controls whether the receiver changes its width to fit the width of its text.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/isHorizontallyResizable
func (t_ Text) SetHorizontallyResizable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setHorizontallyResizable:"), value)
}

// A Boolean that controls whether the receiver allows the user to apply attributes to specific ranges of the text.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/isRichText
func (t_ Text) RichText() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("richText"))
	return rv
}

// SetRichText sets the value of the richText property.
// A Boolean that controls whether the receiver allows the user to apply attributes to specific ranges of the text.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/isRichText
func (t_ Text) SetRichText(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRichText:"), value)
}

// A Boolean value that indicates whether the receiver’s enclosing scroll view shows its ruler.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/isRulerVisible
func (t_ Text) RulerVisible() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("rulerVisible"))
	return rv
}

// A Boolean that controls whether the receiver allows the user to select its text.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/isSelectable
func (t_ Text) Selectable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("selectable"))
	return rv
}

// SetSelectable sets the value of the selectable property.
// A Boolean that controls whether the receiver allows the user to select its text.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/isSelectable
func (t_ Text) SetSelectable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectable:"), value)
}

// A Boolean that controls whether the receiver changes its height to fit the height of its text.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/isVerticallyResizable
func (t_ Text) VerticallyResizable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("verticallyResizable"))
	return rv
}

// SetVerticallyResizable sets the value of the verticallyResizable property.
// A Boolean that controls whether the receiver changes its height to fit the height of its text.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/isVerticallyResizable
func (t_ Text) SetVerticallyResizable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setVerticallyResizable:"), value)
}

// The receiver’s maximum size.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/maxSize
func (t_ Text) MaxSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](t_.ID, objc.Sel("maxSize"))
	return rv
}

// SetMaxSize sets the value of the maxSize property.
// The receiver’s maximum size.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/maxSize
func (t_ Text) SetMaxSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMaxSize:"), value)
}

// The receiver’s minimum size.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/minSize
func (t_ Text) MinSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](t_.ID, objc.Sel("minSize"))
	return rv
}

// SetMinSize sets the value of the minSize property.
// The receiver’s minimum size.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/minSize
func (t_ Text) SetMinSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMinSize:"), value)
}

// The receiver’s characters within .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/selectedRange
func (t_ Text) SelectedRange() foundation.Range {
	rv := objc.Send[foundation.Range](t_.ID, objc.Sel("selectedRange"))
	return rv
}

// SetSelectedRange sets the value of the selectedRange property.
// The receiver’s characters within .

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/selectedRange
func (t_ Text) SetSelectedRange(value foundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectedRange:"), value)
}

// The characters of the receiver’s text.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/string
func (t_ Text) String() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("string"))
	return rv
}

// SetString sets the value of the string property.
// The characters of the receiver’s text.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/string
func (t_ Text) SetString(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setString:"), value)
}

// The text color of all characters in the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/textColor
func (t_ Text) TextColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("textColor"))
	return rv
}

// SetTextColor sets the value of the textColor property.
// The text color of all characters in the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/textColor
func (t_ Text) SetTextColor(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextColor:"), value)
}

// A Boolean that controls whether the receiver uses the Font panel and Font menu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/usesFontPanel
func (t_ Text) UsesFontPanel() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesFontPanel"))
	return rv
}

// SetUsesFontPanel sets the value of the usesFontPanel property.
// A Boolean that controls whether the receiver uses the Font panel and Font menu.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/usesFontPanel
func (t_ Text) SetUsesFontPanel(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesFontPanel:"), value)
}
