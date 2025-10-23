// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	BackgroundColor() IColor
	SetBackgroundColor(value IColor)
	Editable() bool
	SetEditable(value bool)
	RichText() bool
	SetRichText(value bool)
	Alignment() unsafe.Pointer
	SetAlignment(value unsafe.Pointer)
	BaseWritingDirection() unsafe.Pointer
	SetBaseWritingDirection(value unsafe.Pointer)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	Font() IFont
	SetFont(value IFont)
	ImportsGraphics() bool
	SetImportsGraphics(value bool)
	IsEditable() bool
	SetIsEditable(value bool)
	IsFieldEditor() bool
	SetIsFieldEditor(value bool)
	IsHorizontallyResizable() bool
	SetIsHorizontallyResizable(value bool)
	IsRichText() bool
	SetIsRichText(value bool)
	IsRulerVisible() bool
	SetIsRulerVisible(value bool)
	IsSelectable() bool
	SetIsSelectable(value bool)
	IsVerticallyResizable() bool
	SetIsVerticallyResizable(value bool)
	MaxSize() coregraphics.CGSize
	SetMaxSize(value coregraphics.CGSize)
	MinSize() coregraphics.CGSize
	SetMinSize(value coregraphics.CGSize)
	SelectedRange() foundation.Range
	SetSelectedRange(value foundation.Range)
	String() string
	SetString(value string)
	TextColor() IColor
	SetTextColor(value IColor)
	UsesFontPanel() bool
	SetUsesFontPanel(value bool)
	AlignLeft(sender objectivec.IObject)
	Copy(sender objectivec.IObject)
	CopyRuler(sender objectivec.IObject)
	Delete(sender objectivec.IObject)
	PasteRuler(sender objectivec.IObject)
	ReplaceCharactersInRangeWithRTFD(range_ foundation.Range, rtfdData foundation.NSData)
	SetTextColorRange(color IColor, range_ foundation.Range)
	SizeToFit()
	Superscript(sender objectivec.IObject)
	ToggleRuler(sender objectivec.IObject)
	Unscript(sender objectivec.IObject)
	WriteRTFDToFileAtomically(path string, flag bool) bool
}

// The most general programmatic interface for objects that manage text.
//
// draws text for user interface objects, provides text editing capabilities, and controls text attributes such as type size, font, and color. initialization creates an instance of a concrete subclass, such as (generically called a text object). In general, you’re more likely to use the subclass, because it extends the interface declared by and provides much more sophisticated functionality than that declared in . AppKit uses text objects wherever text appears in interface objects. For example, a text object draws the title of a window, the commands in a menu, the title of a button, and the items in a browser. Your app can also create text objects for its own purposes.


// The most general programmatic interface for objects that manage text.
//
// [Full Topic]
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



// This action method applies left alignment to selected paragraphs (or all text if the receiver is a plain text object).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/alignLeft(_:)
func (t_ Text) AlignLeft(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("alignLeft:"), sender)
}


// This action method copies the selected text onto the general pasteboard, in as many formats as the receiver supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/copy(_:)
func (t_ Text) Copy(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("copy:"), sender)
}


// This action method copies the paragraph style information for first selected paragraph onto the ruler pasteboard, as , and expands the selection to paragraph boundaries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/copyRuler(_:)
func (t_ Text) CopyRuler(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("copyRuler:"), sender)
}


// This action method deletes the selected text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/delete(_:)
func (t_ Text) Delete(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("delete:"), sender)
}


// This action method pastes paragraph style information from the ruler pasteboard onto the selected paragraphs of a rich text object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/pasteRuler(_:)
func (t_ Text) PasteRuler(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("pasteRuler:"), sender)
}


// Replaces the characters in the given range with RTFD text interpreted from the given RTFD data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/replaceCharacters(in:withRTFD:)
func (t_ Text) ReplaceCharactersInRangeWithRTFD(range_ foundation.Range, rtfdData foundation.NSData) {
	objc.Send[objc.ID](t_.ID, objc.Sel("replaceCharactersInRange:withRTFD:"), range_, rtfdData)
}


// Sets the text color of characters within the specified range to the specified color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/setTextColor(_:range:)
func (t_ Text) SetTextColorRange(color IColor, range_ foundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextColor:range:"), color, range_)
}


// Resizes the receiver to fit its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/sizeToFit()
func (t_ Text) SizeToFit() {
	objc.Send[objc.ID](t_.ID, objc.Sel("sizeToFit"))
}


// This action method applies a superscript attribute to selected text (or all text if the receiver is a plain text object), raising its baseline offset by a predefined amount.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/superscript(_:)
func (t_ Text) Superscript(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("superscript:"), sender)
}


// This action method shows or hides the ruler, if the receiver is enclosed in a scroll view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/toggleRuler(_:)
func (t_ Text) ToggleRuler(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("toggleRuler:"), sender)
}


// This action method removes any superscripting or subscripting from selected text (or all text if the receiver is a plain text object).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/unscript(_:)
func (t_ Text) Unscript(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("unscript:"), sender)
}


// Writes the receiver’s text as RTF with attachments to a file or directory at .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/writeRTFD(toFile:atomically:)
func (t_ Text) WriteRTFDToFileAtomically(path string, flag bool) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("writeRTFDToFile:atomically:"), objc.String(path), flag)
	return rv
}


// The receiver’s background color to a given color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/backgroundColor
func (t_ Text) BackgroundColor() IColor {
	rv := objc.Send[Color](t_.ID, objc.Sel("backgroundColor"))
	return rv
}


// The receiver’s background color to a given color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/backgroundColor
func (t_ Text) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBackgroundColor:"), value)
}


// A Boolean that controls whether the receiver allows the user to edit its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/isEditable
func (t_ Text) Editable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("editable"))
	return rv
}


// A Boolean that controls whether the receiver allows the user to edit its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/isEditable
func (t_ Text) SetEditable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setEditable:"), value)
}


// A Boolean that controls whether the receiver allows the user to apply attributes to specific ranges of the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/isRichText
func (t_ Text) RichText() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("richText"))
	return rv
}


// A Boolean that controls whether the receiver allows the user to apply attributes to specific ranges of the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/isRichText
func (t_ Text) SetRichText(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRichText:"), value)
}


// The alignment of all the receiver’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/alignment
func (t_ Text) Alignment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("alignment"))
	return rv
}


// The alignment of all the receiver’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/alignment
func (t_ Text) SetAlignment(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAlignment:"), value)
}


// The initial writing direction used to determine the actual writing direction for text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/basewritingdirection
func (t_ Text) BaseWritingDirection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("baseWritingDirection"))
	return rv
}


// The initial writing direction used to determine the actual writing direction for text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/basewritingdirection
func (t_ Text) SetBaseWritingDirection(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBaseWritingDirection:"), value)
}


// The receiver’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/delegate
func (t_ Text) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("delegate"))
	return rv
}


// The receiver’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/delegate
func (t_ Text) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}


// A Boolean that controls whether the receiver draws its background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/drawsbackground
func (t_ Text) DrawsBackground() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("drawsBackground"))
	return rv
}


// A Boolean that controls whether the receiver draws its background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/drawsbackground
func (t_ Text) SetDrawsBackground(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDrawsBackground:"), value)
}


// The font of all the receiver’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/font
func (t_ Text) Font() IFont {
	rv := objc.Send[Font](t_.ID, objc.Sel("font"))
	return rv
}


// The font of all the receiver’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/font
func (t_ Text) SetFont(value IFont) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFont:"), value)
}


// A Boolean that controls whether the receiver allows the user to import files by dragging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/importsgraphics
func (t_ Text) ImportsGraphics() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("importsGraphics"))
	return rv
}


// A Boolean that controls whether the receiver allows the user to import files by dragging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/importsgraphics
func (t_ Text) SetImportsGraphics(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setImportsGraphics:"), value)
}


// A Boolean that controls whether the receiver allows the user to edit its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/iseditable
func (t_ Text) IsEditable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isEditable"))
	return rv
}


// A Boolean that controls whether the receiver allows the user to edit its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/iseditable
func (t_ Text) SetIsEditable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsEditable:"), value)
}


// A Boolean that controls whether the receiver interprets Tab, Shift-Tab, and Return (Enter) as cues to end editing and possibly to change the first responder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/isfieldeditor
func (t_ Text) IsFieldEditor() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isFieldEditor"))
	return rv
}


// A Boolean that controls whether the receiver interprets Tab, Shift-Tab, and Return (Enter) as cues to end editing and possibly to change the first responder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/isfieldeditor
func (t_ Text) SetIsFieldEditor(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsFieldEditor:"), value)
}


// A Boolean that controls whether the receiver changes its width to fit the width of its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/ishorizontallyresizable
func (t_ Text) IsHorizontallyResizable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isHorizontallyResizable"))
	return rv
}


// A Boolean that controls whether the receiver changes its width to fit the width of its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/ishorizontallyresizable
func (t_ Text) SetIsHorizontallyResizable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsHorizontallyResizable:"), value)
}


// A Boolean that controls whether the receiver allows the user to apply attributes to specific ranges of the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/isrichtext
func (t_ Text) IsRichText() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isRichText"))
	return rv
}


// A Boolean that controls whether the receiver allows the user to apply attributes to specific ranges of the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/isrichtext
func (t_ Text) SetIsRichText(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsRichText:"), value)
}


// A Boolean value that indicates whether the receiver’s enclosing scroll view shows its ruler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/isrulervisible
func (t_ Text) IsRulerVisible() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isRulerVisible"))
	return rv
}


// A Boolean value that indicates whether the receiver’s enclosing scroll view shows its ruler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/isrulervisible
func (t_ Text) SetIsRulerVisible(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsRulerVisible:"), value)
}


// A Boolean that controls whether the receiver allows the user to select its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/isselectable
func (t_ Text) IsSelectable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isSelectable"))
	return rv
}


// A Boolean that controls whether the receiver allows the user to select its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/isselectable
func (t_ Text) SetIsSelectable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsSelectable:"), value)
}


// A Boolean that controls whether the receiver changes its height to fit the height of its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/isverticallyresizable
func (t_ Text) IsVerticallyResizable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isVerticallyResizable"))
	return rv
}


// A Boolean that controls whether the receiver changes its height to fit the height of its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/isverticallyresizable
func (t_ Text) SetIsVerticallyResizable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsVerticallyResizable:"), value)
}


// The receiver’s maximum size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/maxsize
func (t_ Text) MaxSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](t_.ID, objc.Sel("maxSize"))
	return rv
}


// The receiver’s maximum size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/maxsize
func (t_ Text) SetMaxSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMaxSize:"), value)
}


// The receiver’s minimum size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/minsize
func (t_ Text) MinSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](t_.ID, objc.Sel("minSize"))
	return rv
}


// The receiver’s minimum size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/minsize
func (t_ Text) SetMinSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMinSize:"), value)
}


// The receiver’s characters within
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/selectedrange
func (t_ Text) SelectedRange() foundation.Range {
	rv := objc.Send[foundation.Range](t_.ID, objc.Sel("selectedRange"))
	return rv
}


// The receiver’s characters within
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/selectedrange
func (t_ Text) SetSelectedRange(value foundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectedRange:"), value)
}


// The characters of the receiver’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/string
func (t_ Text) String() string {
	rv := objc.Send[string](t_.ID, objc.Sel("string"))
	return rv
}


// The characters of the receiver’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/string
func (t_ Text) SetString(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setString:"), objc.String(value))
}


// The text color of all characters in the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/textcolor
func (t_ Text) TextColor() IColor {
	rv := objc.Send[Color](t_.ID, objc.Sel("textColor"))
	return rv
}


// The text color of all characters in the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/textcolor
func (t_ Text) SetTextColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextColor:"), value)
}


// A Boolean that controls whether the receiver uses the Font panel and Font menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/usesfontpanel
func (t_ Text) UsesFontPanel() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesFontPanel"))
	return rv
}


// A Boolean that controls whether the receiver uses the Font panel and Font menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/usesfontpanel
func (t_ Text) SetUsesFontPanel(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesFontPanel:"), value)
}



