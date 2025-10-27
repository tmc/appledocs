// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
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
	

	// properties:
	Alignment() TextAlignment
	SetAlignment(value TextAlignment)
	Font() IFont
	SetFont(value IFont)
	HorizontallyResizable() bool
	SetHorizontallyResizable(value bool)
	RulerVisible() bool
	VerticallyResizable() bool
	SetVerticallyResizable(value bool)
	MaxSize() corefoundation.CGSize
	SetMaxSize(value corefoundation.CGSize)
	MinSize() corefoundation.CGSize
	SetMinSize(value corefoundation.CGSize)
	UsesFontPanel() bool
	SetUsesFontPanel(value bool)
	BackgroundColor() IColor
	SetBackgroundColor(value IColor)
	BaseWritingDirection() WritingDirection
	SetBaseWritingDirection(value WritingDirection)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
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
	SelectedRange() foundation.Range
	SetSelectedRange(value foundation.Range)
	String() foundation.foundation.INSString
	SetString(value foundation.foundation.INSString)
	TextColor() IColor
	SetTextColor(value IColor)


	

	// methods:
	AlignCenter(sender objectivec.IObject)
	AlignLeft(sender objectivec.IObject)
	AlignRight(sender objectivec.IObject)
	ChangeFont(sender objectivec.IObject)
	Copy(sender objectivec.IObject)
	CopyFont(sender objectivec.IObject)
	CopyRuler(sender objectivec.IObject)
	Cut(sender objectivec.IObject)
	Delete(sender objectivec.IObject)
	Paste(sender objectivec.IObject)
	PasteFont(sender objectivec.IObject)
	PasteRuler(sender objectivec.IObject)
	SelectAll(sender objectivec.IObject)
	SetFontRange(font IFont, range_ foundation.Range)
	SizeToFit()
	Subscript(sender objectivec.IObject)
	Superscript(sender objectivec.IObject)
	ToggleRuler(sender objectivec.IObject)
	Unscript(sender objectivec.IObject)


}





// Alloc allocates a new instance without initialization.
func (tc _TextClass) Alloc() Text {
	rv := objc.Send[Text](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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




















// This action method applies center alignment to selected paragraphs (or all text if the receiver is a plain text object).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/alignCenter(_:)
func (t_ Text) AlignCenter(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("alignCenter:"), sender)
}


// This action method applies left alignment to selected paragraphs (or all text if the receiver is a plain text object).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/alignLeft(_:)
func (t_ Text) AlignLeft(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("alignLeft:"), sender)
}


// This action method applies right alignment to selected paragraphs (or all text if the receiver is a plain text object).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/alignRight(_:)
func (t_ Text) AlignRight(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("alignRight:"), sender)
}


// This action method changes the font of the selection for a rich text object, or of all text for a plain text object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/changeFont(_:)
func (t_ Text) ChangeFont(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("changeFont:"), sender)
}


// This action method copies the selected text onto the general pasteboard, in as many formats as the receiver supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/copy(_:)
func (t_ Text) Copy(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("copy:"), sender)
}


// This action method copies the font information for the first character of the selection (or for the insertion point) onto the font pasteboard, as .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/copyFont(_:)
func (t_ Text) CopyFont(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("copyFont:"), sender)
}


// This action method copies the paragraph style information for first selected paragraph onto the ruler pasteboard, as , and expands the selection to paragraph boundaries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/copyRuler(_:)
func (t_ Text) CopyRuler(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("copyRuler:"), sender)
}


// This action method deletes the selected text and places it onto the general pasteboard, in as many formats as the receiver supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/cut(_:)
func (t_ Text) Cut(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("cut:"), sender)
}


// This action method deletes the selected text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/delete(_:)
func (t_ Text) Delete(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("delete:"), sender)
}


// This action method pastes text from the general pasteboard at the insertion point or over the selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/paste(_:)
func (t_ Text) Paste(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("paste:"), sender)
}


// This action method pastes font information from the font pasteboard onto the selected text or insertion point of a rich text object, or over all text of a plain text object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/pasteFont(_:)
func (t_ Text) PasteFont(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("pasteFont:"), sender)
}


// This action method pastes paragraph style information from the ruler pasteboard onto the selected paragraphs of a rich text object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/pasteRuler(_:)
func (t_ Text) PasteRuler(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("pasteRuler:"), sender)
}


// This action method selects all of the receiver’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/selectAll(_:)
func (t_ Text) SelectAll(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("selectAll:"), sender)
}


// Sets the font of characters within to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/setFont(_:range:)
func (t_ Text) SetFontRange(font IFont, range_ foundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFont:range:"), font, range_)
}


// Resizes the receiver to fit its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/sizeToFit()
func (t_ Text) SizeToFit() {
	objc.Send[objc.ID](t_.ID, objc.Sel("sizeToFit"))
}


// This action method applies a subscript attribute to selected text (or all text if the receiver is a plain text object), lowering its baseline offset by a predefined amount.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/subscript(_:)
func (t_ Text) Subscript(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("subscript:"), sender)
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







// The alignment of all the receiver’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/alignment
func (t_ Text) Alignment() TextAlignment {
	rv := objc.Send[TextAlignment](t_.ID, objc.Sel("alignment"))
	return rv
}


// The alignment of all the receiver’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/alignment
func (t_ Text) SetAlignment(value TextAlignment) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAlignment:"), value)
}


// The font of all the receiver’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/font
func (t_ Text) Font() IFont {
	rv := objc.Send[Font](t_.ID, objc.Sel("font"))
	return rv
}


// The font of all the receiver’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/font
func (t_ Text) SetFont(value IFont) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFont:"), value)
}


// A Boolean that controls whether the receiver changes its width to fit the width of its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/isHorizontallyResizable
func (t_ Text) HorizontallyResizable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("horizontallyResizable"))
	return rv
}


// A Boolean that controls whether the receiver changes its width to fit the width of its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/isHorizontallyResizable
func (t_ Text) SetHorizontallyResizable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setHorizontallyResizable:"), value)
}


// A Boolean value that indicates whether the receiver’s enclosing scroll view shows its ruler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/isRulerVisible
func (t_ Text) RulerVisible() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("rulerVisible"))
	return rv
}


// A Boolean that controls whether the receiver changes its height to fit the height of its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/isVerticallyResizable
func (t_ Text) VerticallyResizable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("verticallyResizable"))
	return rv
}


// A Boolean that controls whether the receiver changes its height to fit the height of its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/isVerticallyResizable
func (t_ Text) SetVerticallyResizable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setVerticallyResizable:"), value)
}


// The receiver’s maximum size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/maxSize
func (t_ Text) MaxSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](t_.ID, objc.Sel("maxSize"))
	return rv
}


// The receiver’s maximum size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/maxSize
func (t_ Text) SetMaxSize(value corefoundation.CGSize) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMaxSize:"), value)
}


// The receiver’s minimum size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/minSize
func (t_ Text) MinSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](t_.ID, objc.Sel("minSize"))
	return rv
}


// The receiver’s minimum size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/minSize
func (t_ Text) SetMinSize(value corefoundation.CGSize) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMinSize:"), value)
}


// A Boolean that controls whether the receiver uses the Font panel and Font menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/usesFontPanel
func (t_ Text) UsesFontPanel() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesFontPanel"))
	return rv
}


// A Boolean that controls whether the receiver uses the Font panel and Font menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/usesFontPanel
func (t_ Text) SetUsesFontPanel(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesFontPanel:"), value)
}


// The receiver’s background color to a given color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/backgroundcolor
func (t_ Text) BackgroundColor() IColor {
	rv := objc.Send[Color](t_.ID, objc.Sel("backgroundColor"))
	return rv
}


// The receiver’s background color to a given color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/backgroundcolor
func (t_ Text) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBackgroundColor:"), value)
}


// The initial writing direction used to determine the actual writing direction for text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/basewritingdirection
func (t_ Text) BaseWritingDirection() WritingDirection {
	rv := objc.Send[WritingDirection](t_.ID, objc.Sel("baseWritingDirection"))
	return rv
}


// The initial writing direction used to determine the actual writing direction for text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/basewritingdirection
func (t_ Text) SetBaseWritingDirection(value WritingDirection) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBaseWritingDirection:"), value)
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
func (t_ Text) String() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("string"))
	return rv
}


// The characters of the receiver’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/string
func (t_ Text) SetString(value foundation.foundation.INSString) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setString:"), value)
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








