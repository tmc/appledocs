// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Text] class.
var TextClass objc.Class

func init() {
	TextClass = objc.GetClass("NSText")
}

type Text struct {
	objc.ID
}

func TextFrom(ptr unsafe.Pointer) Text {
	return Text{
		ID: objc.ID(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc Text) Alloc() Text {
	ret := objc.ID(TextClass).Send(objc.RegisterName("alloc"))
	return Text{ret}
}

// Init initializes the instance.
func (t_ Text) Init() Text {
	ret := t_.ID.Send(objc.RegisterName("init"))
	return Text{ret}
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSText/init(coder:)
func NewTextWithCoder(coder unsafe.Pointer) Text {
	instance := Text{}.Alloc()
	sel := objc.RegisterName("initWithCoder:")
	ret := instance.ID.Send(sel, coder)
	instance = Text{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSText/init(frame:)
func NewTextWithFrame(frameRect unsafe.Pointer) Text {
	instance := Text{}.Alloc()
	sel := objc.RegisterName("initWithFrame:")
	ret := instance.ID.Send(sel, frameRect)
	instance = Text{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}


// This action method applies center alignment to selected paragraphs (or all text if the receiver is a plain text object). [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSText/alignCenter(_:)
func (t_ Text) AlignCenter(sender objc.ID) {
	sel := objc.RegisterName("alignCenter:")
	t_.ID.Send(sel, sender)
}
// This action method applies left alignment to selected paragraphs (or all text if the receiver is a plain text object). [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSText/alignLeft(_:)
func (t_ Text) AlignLeft(sender objc.ID) {
	sel := objc.RegisterName("alignLeft:")
	t_.ID.Send(sel, sender)
}
// This action method applies right alignment to selected paragraphs (or all text if the receiver is a plain text object). [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSText/alignRight(_:)
func (t_ Text) AlignRight(sender objc.ID) {
	sel := objc.RegisterName("alignRight:")
	t_.ID.Send(sel, sender)
}
// This action method changes the font of the selection for a rich text object, or of all text for a plain text object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSText/changeFont(_:)
func (t_ Text) ChangeFont(sender objc.ID) {
	sel := objc.RegisterName("changeFont:")
	t_.ID.Send(sel, sender)
}
// This action method searches for a misspelled word in the receiver’s text. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSText/checkSpelling(_:)
func (t_ Text) CheckSpelling(sender objc.ID) {
	sel := objc.RegisterName("checkSpelling:")
	t_.ID.Send(sel, sender)
}
// This action method copies the selected text onto the general pasteboard, in as many formats as the receiver supports. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSText/copy(_:)
func (t_ Text) Copy(sender objc.ID) {
	sel := objc.RegisterName("copy:")
	t_.ID.Send(sel, sender)
}
// This action method copies the font information for the first character of the selection (or for the insertion point) onto the font pasteboard, as  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSText/copyFont(_:)
func (t_ Text) CopyFont(sender objc.ID) {
	sel := objc.RegisterName("copyFont:")
	t_.ID.Send(sel, sender)
}
// This action method copies the paragraph style information for first selected paragraph onto the ruler pasteboard, as  , and expands the selection to paragraph boundaries. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSText/copyRuler(_:)
func (t_ Text) CopyRuler(sender objc.ID) {
	sel := objc.RegisterName("copyRuler:")
	t_.ID.Send(sel, sender)
}
// This action method deletes the selected text and places it onto the general pasteboard, in as many formats as the receiver supports. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSText/cut(_:)
func (t_ Text) Cut(sender objc.ID) {
	sel := objc.RegisterName("cut:")
	t_.ID.Send(sel, sender)
}
// This action method deletes the selected text. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSText/delete(_:)
func (t_ Text) Delete(sender objc.ID) {
	sel := objc.RegisterName("delete:")
	t_.ID.Send(sel, sender)
}
// This action method pastes text from the general pasteboard at the insertion point or over the selection. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSText/paste(_:)
func (t_ Text) Paste(sender objc.ID) {
	sel := objc.RegisterName("paste:")
	t_.ID.Send(sel, sender)
}
// This action method pastes font information from the font pasteboard onto the selected text or insertion point of a rich text object, or over all text of a plain text object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSText/pasteFont(_:)
func (t_ Text) PasteFont(sender objc.ID) {
	sel := objc.RegisterName("pasteFont:")
	t_.ID.Send(sel, sender)
}
// This action method pastes paragraph style information from the ruler pasteboard onto the selected paragraphs of a rich text object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSText/pasteRuler(_:)
func (t_ Text) PasteRuler(sender objc.ID) {
	sel := objc.RegisterName("pasteRuler:")
	t_.ID.Send(sel, sender)
}
// Attempts to read the RTFD file at  , returning   if successful and   if not. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSText/readRTFD(fromFile:)
func (t_ Text) ReadRTFDFromFile(path unsafe.Pointer) bool {
	sel := objc.RegisterName("readRTFDFromFile:")
	ret := t_.ID.Send(sel, path)
	return ret != 0
}
// Replaces the characters in the given range with those in the given string. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSText/replaceCharacters(in:with:)
func (t_ Text) ReplaceCharactersInRangeWithString(range_ unsafe.Pointer, string unsafe.Pointer) {
	sel := objc.RegisterName("replaceCharactersInRange:withString:")
	t_.ID.Send(sel, range_, string)
}
// Replaces the characters in the given range with RTF text interpreted from the given RTF data. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSText/replaceCharacters(in:withRTF:)
func (t_ Text) ReplaceCharactersInRangeWithRTF(range_ unsafe.Pointer, rtfData unsafe.Pointer) {
	sel := objc.RegisterName("replaceCharactersInRange:withRTF:")
	t_.ID.Send(sel, range_, rtfData)
}
// Replaces the characters in the given range with RTFD text interpreted from the given RTFD data. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSText/replaceCharacters(in:withRTFD:)
func (t_ Text) ReplaceCharactersInRangeWithRTFD(range_ unsafe.Pointer, rtfdData unsafe.Pointer) {
	sel := objc.RegisterName("replaceCharactersInRange:withRTFD:")
	t_.ID.Send(sel, range_, rtfdData)
}
// Returns an NSData object that contains an RTF stream corresponding to the characters and attributes within  , omitting any attachment characters and attributes. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSText/rtf(from:)
func (t_ Text) RTFFromRange(range_ unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("RTFFromRange:")
	ret := t_.ID.Send(sel, range_)
	return unsafe.Pointer(ret)
}
// Returns an NSData object that contains an RTFD stream corresponding to the characters and attributes within  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSText/rtfd(from:)
func (t_ Text) RTFDFromRange(range_ unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("RTFDFromRange:")
	ret := t_.ID.Send(sel, range_)
	return unsafe.Pointer(ret)
}
// Scrolls the receiver in its enclosing scroll view so the first characters of   are visible. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSText/scrollRangeToVisible(_:)
func (t_ Text) ScrollRangeToVisible(range_ unsafe.Pointer) {
	sel := objc.RegisterName("scrollRangeToVisible:")
	t_.ID.Send(sel, range_)
}
// This action method selects all of the receiver’s text. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSText/selectAll(_:)
func (t_ Text) SelectAll(sender objc.ID) {
	sel := objc.RegisterName("selectAll:")
	t_.ID.Send(sel, sender)
}
// Sets the font of characters within   to  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSText/setFont(_:range:)
func (t_ Text) SetFontRange(font unsafe.Pointer, range_ unsafe.Pointer) {
	sel := objc.RegisterName("setFont:range:")
	t_.ID.Send(sel, font, range_)
}
// Sets the text color of characters within the specified range to the specified color. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSText/setTextColor(_:range:)
func (t_ Text) SetTextColorRange(color unsafe.Pointer, range_ unsafe.Pointer) {
	sel := objc.RegisterName("setTextColor:range:")
	t_.ID.Send(sel, color, range_)
}
// This action method opens the Spelling panel, allowing the user to make a correction during spell checking. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSText/showGuessPanel(_:)
func (t_ Text) ShowGuessPanel(sender objc.ID) {
	sel := objc.RegisterName("showGuessPanel:")
	t_.ID.Send(sel, sender)
}
// Resizes the receiver to fit its text. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSText/sizeToFit()
func (t_ Text) SizeToFit() {
	sel := objc.RegisterName("sizeToFit")
	t_.ID.Send(sel)
}
// This action method applies a subscript attribute to selected text (or all text if the receiver is a plain text object), lowering its baseline offset by a predefined amount. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSText/subscript(_:)
func (t_ Text) Subscript(sender objc.ID) {
	sel := objc.RegisterName("subscript:")
	t_.ID.Send(sel, sender)
}
// This action method applies a superscript attribute to selected text (or all text if the receiver is a plain text object), raising its baseline offset by a predefined amount. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSText/superscript(_:)
func (t_ Text) Superscript(sender objc.ID) {
	sel := objc.RegisterName("superscript:")
	t_.ID.Send(sel, sender)
}
// This action method shows or hides the ruler, if the receiver is enclosed in a scroll view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSText/toggleRuler(_:)
func (t_ Text) ToggleRuler(sender objc.ID) {
	sel := objc.RegisterName("toggleRuler:")
	t_.ID.Send(sel, sender)
}
// Adds the underline attribute to the selected text attributes if absent; removes the attribute if present. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSText/underline(_:)
func (t_ Text) Underline(sender objc.ID) {
	sel := objc.RegisterName("underline:")
	t_.ID.Send(sel, sender)
}
// This action method removes any superscripting or subscripting from selected text (or all text if the receiver is a plain text object). [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSText/unscript(_:)
func (t_ Text) Unscript(sender objc.ID) {
	sel := objc.RegisterName("unscript:")
	t_.ID.Send(sel, sender)
}
// Writes the receiver’s text as RTF with attachments to a file or directory at  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSText/writeRTFD(toFile:atomically:)
func (t_ Text) WriteRTFDToFileAtomically(path unsafe.Pointer, flag bool) bool {
	sel := objc.RegisterName("writeRTFDToFile:atomically:")
	ret := t_.ID.Send(sel, path, flag)
	return ret != 0
}

