// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Text] class.
var textClass = _TextClass{objc.GetClass("NSText")}

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
	ReplaceCharactersInRangeWithString(range_ unsafe.Pointer, string string)
	ReplaceCharactersInRangeWithRTF(range_ unsafe.Pointer, rtfData unsafe.Pointer)
	ReplaceCharactersInRangeWithRTFD(range_ unsafe.Pointer, rtfdData unsafe.Pointer)
	RTFFromRange(range_ unsafe.Pointer) unsafe.Pointer
	RTFDFromRange(range_ unsafe.Pointer) unsafe.Pointer
	ScrollRangeToVisible(range_ unsafe.Pointer)
	SelectAll(sender objc.ID)
	SetFontRange(font unsafe.Pointer, range_ unsafe.Pointer)
	SetTextColorRange(color unsafe.Pointer, range_ unsafe.Pointer)
	ShowGuessPanel(sender objc.ID)
	SizeToFit()
	Subscript(sender objc.ID)
	Superscript(sender objc.ID)
	ToggleRuler(sender objc.ID)
	Underline(sender objc.ID)
	Unscript(sender objc.ID)
	WriteRTFDToFileAtomically(path string, flag bool) bool
}

// The most general programmatic interface for objects that manage text. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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
	return textClass.New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/init(frame:)
func NewTextWithFrame(frameRect unsafe.Pointer) Text {
	instance := textClass.Alloc()
	rv := objc.Send[Text](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	rv.Autorelease()
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/init(coder:)
func NewTextWithCoder(coder unsafe.Pointer) Text {
	instance := textClass.Alloc()
	rv := objc.Send[Text](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}


// This action method applies center alignment to selected paragraphs (or all text if the receiver is a plain text object). [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/alignCenter(_:)
func (t_ Text) AlignCenter(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("alignCenter:"), sender)
}
// This action method applies left alignment to selected paragraphs (or all text if the receiver is a plain text object). [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/alignLeft(_:)
func (t_ Text) AlignLeft(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("alignLeft:"), sender)
}
// This action method applies right alignment to selected paragraphs (or all text if the receiver is a plain text object). [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/alignRight(_:)
func (t_ Text) AlignRight(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("alignRight:"), sender)
}
// This action method changes the font of the selection for a rich text object, or of all text for a plain text object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/changeFont(_:)
func (t_ Text) ChangeFont(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("changeFont:"), sender)
}
// This action method searches for a misspelled word in the receiver’s text. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/checkSpelling(_:)
func (t_ Text) CheckSpelling(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("checkSpelling:"), sender)
}
// This action method copies the selected text onto the general pasteboard, in as many formats as the receiver supports. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/copy(_:)
func (t_ Text) Copy(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("copy:"), sender)
}
// This action method copies the font information for the first character of the selection (or for the insertion point) onto the font pasteboard, as . [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/copyFont(_:)
func (t_ Text) CopyFont(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("copyFont:"), sender)
}
// This action method copies the paragraph style information for first selected paragraph onto the ruler pasteboard, as , and expands the selection to paragraph boundaries. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/copyRuler(_:)
func (t_ Text) CopyRuler(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("copyRuler:"), sender)
}
// This action method deletes the selected text and places it onto the general pasteboard, in as many formats as the receiver supports. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/cut(_:)
func (t_ Text) Cut(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("cut:"), sender)
}
// This action method deletes the selected text. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/delete(_:)
func (t_ Text) Delete(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("delete:"), sender)
}
// This action method pastes text from the general pasteboard at the insertion point or over the selection. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/paste(_:)
func (t_ Text) Paste(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("paste:"), sender)
}
// This action method pastes font information from the font pasteboard onto the selected text or insertion point of a rich text object, or over all text of a plain text object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/pasteFont(_:)
func (t_ Text) PasteFont(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("pasteFont:"), sender)
}
// This action method pastes paragraph style information from the ruler pasteboard onto the selected paragraphs of a rich text object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/pasteRuler(_:)
func (t_ Text) PasteRuler(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("pasteRuler:"), sender)
}
// Attempts to read the RTFD file at , returning if successful and if not. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/readRTFD(fromFile:)
func (t_ Text) ReadRTFDFromFile(path string) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("readRTFDFromFile:"), objc.String(path))
	return rv
}
// Replaces the characters in the given range with those in the given string. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/replaceCharacters(in:with:)
func (t_ Text) ReplaceCharactersInRangeWithString(range_ unsafe.Pointer, string string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("replaceCharactersInRange:withString:"), range_, objc.String(string))
}
// Replaces the characters in the given range with RTF text interpreted from the given RTF data. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/replaceCharacters(in:withRTF:)
func (t_ Text) ReplaceCharactersInRangeWithRTF(range_ unsafe.Pointer, rtfData unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("replaceCharactersInRange:withRTF:"), range_, rtfData)
}
// Replaces the characters in the given range with RTFD text interpreted from the given RTFD data. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/replaceCharacters(in:withRTFD:)
func (t_ Text) ReplaceCharactersInRangeWithRTFD(range_ unsafe.Pointer, rtfdData unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("replaceCharactersInRange:withRTFD:"), range_, rtfdData)
}
// Returns an NSData object that contains an RTF stream corresponding to the characters and attributes within , omitting any attachment characters and attributes. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/rtf(from:)
func (t_ Text) RTFFromRange(range_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("RTFFromRange:"), range_)
	return rv
}
// Returns an NSData object that contains an RTFD stream corresponding to the characters and attributes within . [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/rtfd(from:)
func (t_ Text) RTFDFromRange(range_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("RTFDFromRange:"), range_)
	return rv
}
// Scrolls the receiver in its enclosing scroll view so the first characters of are visible. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/scrollRangeToVisible(_:)
func (t_ Text) ScrollRangeToVisible(range_ unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("scrollRangeToVisible:"), range_)
}
// This action method selects all of the receiver’s text. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/selectAll(_:)
func (t_ Text) SelectAll(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("selectAll:"), sender)
}
// Sets the font of characters within to . [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/setFont(_:range:)
func (t_ Text) SetFontRange(font unsafe.Pointer, range_ unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFont:range:"), font, range_)
}
// Sets the text color of characters within the specified range to the specified color. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/setTextColor(_:range:)
func (t_ Text) SetTextColorRange(color unsafe.Pointer, range_ unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextColor:range:"), color, range_)
}
// This action method opens the Spelling panel, allowing the user to make a correction during spell checking. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/showGuessPanel(_:)
func (t_ Text) ShowGuessPanel(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("showGuessPanel:"), sender)
}
// Resizes the receiver to fit its text. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/sizeToFit()
func (t_ Text) SizeToFit() {
	objc.Send[objc.ID](t_.ID, objc.Sel("sizeToFit"))
}
// This action method applies a subscript attribute to selected text (or all text if the receiver is a plain text object), lowering its baseline offset by a predefined amount. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/subscript(_:)
func (t_ Text) Subscript(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("subscript:"), sender)
}
// This action method applies a superscript attribute to selected text (or all text if the receiver is a plain text object), raising its baseline offset by a predefined amount. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/superscript(_:)
func (t_ Text) Superscript(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("superscript:"), sender)
}
// This action method shows or hides the ruler, if the receiver is enclosed in a scroll view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/toggleRuler(_:)
func (t_ Text) ToggleRuler(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("toggleRuler:"), sender)
}
// Adds the underline attribute to the selected text attributes if absent; removes the attribute if present. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/underline(_:)
func (t_ Text) Underline(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("underline:"), sender)
}
// This action method removes any superscripting or subscripting from selected text (or all text if the receiver is a plain text object). [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/unscript(_:)
func (t_ Text) Unscript(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("unscript:"), sender)
}
// Writes the receiver’s text as RTF with attachments to a file or directory at . [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/writeRTFD(toFile:atomically:)
func (t_ Text) WriteRTFDToFileAtomically(path string, flag bool) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("writeRTFDToFile:atomically:"), objc.String(path), flag)
	return rv
}

