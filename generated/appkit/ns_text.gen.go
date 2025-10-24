// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSText */


/* debug [class_header]: Header for NSText */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Text */
// An interface definition for the [Text] class.
type IText interface {
	IView
	
/* debug [class_interface_properties]: Properties for Text */
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
	MaxSize() Size /* not a class type */
	SetMaxSize(value Size /* not a class type */)
	MinSize() Size /* not a class type */
	SetMinSize(value Size /* not a class type */)
	UsesFontPanel() bool
	SetUsesFontPanel(value bool)
	BackgroundColor() IColor
	SetBackgroundColor(value IColor)
	BaseWritingDirection() WritingDirection
	SetBaseWritingDirection(value WritingDirection)
	Delegate() objc.IObject /* cross-framework: TextDelegate */
	SetDelegate(value objc.IObject /* cross-framework: TextDelegate */)
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
	SelectedRange() corefoundation.Range
	SetSelectedRange(value corefoundation.Range)
	String() objc.IObject /* cross-framework: NSString */
	SetString(value objc.IObject /* cross-framework: NSString */)
	TextColor() IColor
	SetTextColor(value IColor)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Text */
	// methods:
	AlignCenter(sender objc.IObject)
	AlignLeft(sender objc.IObject)
	AlignRight(sender objc.IObject)
	ChangeFont(sender objc.IObject)
	Copy(sender objc.IObject)
	CopyFont(sender objc.IObject)
	CopyRuler(sender objc.IObject)
	Cut(sender objc.IObject)
	Delete(sender objc.IObject)
	Paste(sender objc.IObject)
	PasteFont(sender objc.IObject)
	PasteRuler(sender objc.IObject)
	SelectAll(sender objc.IObject)
	SetFontRange(font IFont, range_ corefoundation.Range)
	SizeToFit()
	Subscript(sender objc.IObject)
	Superscript(sender objc.IObject)
	ToggleRuler(sender objc.IObject)
	Unscript(sender objc.IObject)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Text */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Text */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Text *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Text */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Text */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Text */

// This action method applies center alignment to selected paragraphs (or all text if the receiver is a plain text object).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/alignCenter(_:)
func (t_ Text) AlignCenter(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("alignCenter:"), sender)
}/* debug [instance_methods/method]: AlignCenter */


// This action method applies left alignment to selected paragraphs (or all text if the receiver is a plain text object).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/alignLeft(_:)
func (t_ Text) AlignLeft(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("alignLeft:"), sender)
}/* debug [instance_methods/method]: AlignLeft */


// This action method applies right alignment to selected paragraphs (or all text if the receiver is a plain text object).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/alignRight(_:)
func (t_ Text) AlignRight(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("alignRight:"), sender)
}/* debug [instance_methods/method]: AlignRight */


// This action method changes the font of the selection for a rich text object, or of all text for a plain text object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/changeFont(_:)
func (t_ Text) ChangeFont(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("changeFont:"), sender)
}/* debug [instance_methods/method]: ChangeFont */


// This action method copies the selected text onto the general pasteboard, in as many formats as the receiver supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/copy(_:)
func (t_ Text) Copy(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("copy:"), sender)
}/* debug [instance_methods/method]: Copy */


// This action method copies the font information for the first character of the selection (or for the insertion point) onto the font pasteboard, as .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/copyFont(_:)
func (t_ Text) CopyFont(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("copyFont:"), sender)
}/* debug [instance_methods/method]: CopyFont */


// This action method copies the paragraph style information for first selected paragraph onto the ruler pasteboard, as , and expands the selection to paragraph boundaries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/copyRuler(_:)
func (t_ Text) CopyRuler(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("copyRuler:"), sender)
}/* debug [instance_methods/method]: CopyRuler */


// This action method deletes the selected text and places it onto the general pasteboard, in as many formats as the receiver supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/cut(_:)
func (t_ Text) Cut(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("cut:"), sender)
}/* debug [instance_methods/method]: Cut */


// This action method deletes the selected text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/delete(_:)
func (t_ Text) Delete(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("delete:"), sender)
}/* debug [instance_methods/method]: Delete */


// This action method pastes text from the general pasteboard at the insertion point or over the selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/paste(_:)
func (t_ Text) Paste(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("paste:"), sender)
}/* debug [instance_methods/method]: Paste */


// This action method pastes font information from the font pasteboard onto the selected text or insertion point of a rich text object, or over all text of a plain text object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/pasteFont(_:)
func (t_ Text) PasteFont(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("pasteFont:"), sender)
}/* debug [instance_methods/method]: PasteFont */


// This action method pastes paragraph style information from the ruler pasteboard onto the selected paragraphs of a rich text object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/pasteRuler(_:)
func (t_ Text) PasteRuler(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("pasteRuler:"), sender)
}/* debug [instance_methods/method]: PasteRuler */


// This action method selects all of the receiver’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/selectAll(_:)
func (t_ Text) SelectAll(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("selectAll:"), sender)
}/* debug [instance_methods/method]: SelectAll */


// Sets the font of characters within to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/setFont(_:range:)
func (t_ Text) SetFontRange(font IFont, range_ corefoundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFont:range:"), font, range_)
}/* debug [instance_methods/method]: SetFontRange */


// Resizes the receiver to fit its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/sizeToFit()
func (t_ Text) SizeToFit() {
	objc.Send[objc.ID](t_.ID, objc.Sel("sizeToFit"))
}/* debug [instance_methods/method]: SizeToFit */


// This action method applies a subscript attribute to selected text (or all text if the receiver is a plain text object), lowering its baseline offset by a predefined amount.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/subscript(_:)
func (t_ Text) Subscript(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("subscript:"), sender)
}/* debug [instance_methods/method]: Subscript */


// This action method applies a superscript attribute to selected text (or all text if the receiver is a plain text object), raising its baseline offset by a predefined amount.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/superscript(_:)
func (t_ Text) Superscript(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("superscript:"), sender)
}/* debug [instance_methods/method]: Superscript */


// This action method shows or hides the ruler, if the receiver is enclosed in a scroll view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/toggleRuler(_:)
func (t_ Text) ToggleRuler(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("toggleRuler:"), sender)
}/* debug [instance_methods/method]: ToggleRuler */


// This action method removes any superscripting or subscripting from selected text (or all text if the receiver is a plain text object).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/unscript(_:)
func (t_ Text) Unscript(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("unscript:"), sender)
}/* debug [instance_methods/method]: Unscript */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Text */

// The alignment of all the receiver’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/alignment
func (t_ Text) Alignment() TextAlignment {
	rv := objc.Send[TextAlignment](t_.ID, objc.Sel("alignment"))
	return rv
}/* debug [instance_properties/getter]: alignment */


// The alignment of all the receiver’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/alignment
func (t_ Text) SetAlignment(value TextAlignment) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAlignment:"), value)
}/* debug [instance_properties/setter]: alignment */


// The font of all the receiver’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/font
func (t_ Text) Font() IFont {
	rv := objc.Send[Font](t_.ID, objc.Sel("font"))
	return rv
}/* debug [instance_properties/getter]: font */


// The font of all the receiver’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/font
func (t_ Text) SetFont(value IFont) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFont:"), value)
}/* debug [instance_properties/setter]: font */


// A Boolean that controls whether the receiver changes its width to fit the width of its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/isHorizontallyResizable
func (t_ Text) HorizontallyResizable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("horizontallyResizable"))
	return rv
}/* debug [instance_properties/getter]: horizontallyResizable */


// A Boolean that controls whether the receiver changes its width to fit the width of its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/isHorizontallyResizable
func (t_ Text) SetHorizontallyResizable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setHorizontallyResizable:"), value)
}/* debug [instance_properties/setter]: horizontallyResizable */


// A Boolean value that indicates whether the receiver’s enclosing scroll view shows its ruler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/isRulerVisible
func (t_ Text) RulerVisible() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("rulerVisible"))
	return rv
}/* debug [instance_properties/getter]: rulerVisible */


// A Boolean that controls whether the receiver changes its height to fit the height of its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/isVerticallyResizable
func (t_ Text) VerticallyResizable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("verticallyResizable"))
	return rv
}/* debug [instance_properties/getter]: verticallyResizable */


// A Boolean that controls whether the receiver changes its height to fit the height of its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/isVerticallyResizable
func (t_ Text) SetVerticallyResizable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setVerticallyResizable:"), value)
}/* debug [instance_properties/setter]: verticallyResizable */


// The receiver’s maximum size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/maxSize
func (t_ Text) MaxSize() Size /* not a class type */ {
	rv := objc.Send[Size](t_.ID, objc.Sel("maxSize"))
	return rv
}/* debug [instance_properties/getter]: maxSize */


// The receiver’s maximum size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/maxSize
func (t_ Text) SetMaxSize(value Size /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMaxSize:"), value)
}/* debug [instance_properties/setter]: maxSize */


// The receiver’s minimum size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/minSize
func (t_ Text) MinSize() Size /* not a class type */ {
	rv := objc.Send[Size](t_.ID, objc.Sel("minSize"))
	return rv
}/* debug [instance_properties/getter]: minSize */


// The receiver’s minimum size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/minSize
func (t_ Text) SetMinSize(value Size /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMinSize:"), value)
}/* debug [instance_properties/setter]: minSize */


// A Boolean that controls whether the receiver uses the Font panel and Font menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/usesFontPanel
func (t_ Text) UsesFontPanel() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesFontPanel"))
	return rv
}/* debug [instance_properties/getter]: usesFontPanel */


// A Boolean that controls whether the receiver uses the Font panel and Font menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSText/usesFontPanel
func (t_ Text) SetUsesFontPanel(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesFontPanel:"), value)
}/* debug [instance_properties/setter]: usesFontPanel */


// The receiver’s background color to a given color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/backgroundcolor
func (t_ Text) BackgroundColor() IColor {
	rv := objc.Send[Color](t_.ID, objc.Sel("backgroundColor"))
	return rv
}/* debug [instance_properties/getter]: backgroundColor */


// The receiver’s background color to a given color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/backgroundcolor
func (t_ Text) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBackgroundColor:"), value)
}/* debug [instance_properties/setter]: backgroundColor */


// The initial writing direction used to determine the actual writing direction for text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/basewritingdirection
func (t_ Text) BaseWritingDirection() WritingDirection {
	rv := objc.Send[WritingDirection](t_.ID, objc.Sel("baseWritingDirection"))
	return rv
}/* debug [instance_properties/getter]: baseWritingDirection */


// The initial writing direction used to determine the actual writing direction for text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/basewritingdirection
func (t_ Text) SetBaseWritingDirection(value WritingDirection) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBaseWritingDirection:"), value)
}/* debug [instance_properties/setter]: baseWritingDirection */


// The receiver’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/delegate
func (t_ Text) Delegate() objc.IObject /* cross-framework: TextDelegate */ {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The receiver’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/delegate
func (t_ Text) SetDelegate(value objc.IObject /* cross-framework: TextDelegate */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// A Boolean that controls whether the receiver draws its background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/drawsbackground
func (t_ Text) DrawsBackground() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("drawsBackground"))
	return rv
}/* debug [instance_properties/getter]: drawsBackground */


// A Boolean that controls whether the receiver draws its background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/drawsbackground
func (t_ Text) SetDrawsBackground(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDrawsBackground:"), value)
}/* debug [instance_properties/setter]: drawsBackground */


// A Boolean that controls whether the receiver allows the user to import files by dragging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/importsgraphics
func (t_ Text) ImportsGraphics() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("importsGraphics"))
	return rv
}/* debug [instance_properties/getter]: importsGraphics */


// A Boolean that controls whether the receiver allows the user to import files by dragging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/importsgraphics
func (t_ Text) SetImportsGraphics(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setImportsGraphics:"), value)
}/* debug [instance_properties/setter]: importsGraphics */


// A Boolean that controls whether the receiver allows the user to edit its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/iseditable
func (t_ Text) IsEditable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isEditable"))
	return rv
}/* debug [instance_properties/getter]: isEditable */


// A Boolean that controls whether the receiver allows the user to edit its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/iseditable
func (t_ Text) SetIsEditable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsEditable:"), value)
}/* debug [instance_properties/setter]: isEditable */


// A Boolean that controls whether the receiver interprets Tab, Shift-Tab, and Return (Enter) as cues to end editing and possibly to change the first responder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/isfieldeditor
func (t_ Text) IsFieldEditor() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isFieldEditor"))
	return rv
}/* debug [instance_properties/getter]: isFieldEditor */


// A Boolean that controls whether the receiver interprets Tab, Shift-Tab, and Return (Enter) as cues to end editing and possibly to change the first responder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/isfieldeditor
func (t_ Text) SetIsFieldEditor(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsFieldEditor:"), value)
}/* debug [instance_properties/setter]: isFieldEditor */


// A Boolean that controls whether the receiver changes its width to fit the width of its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/ishorizontallyresizable
func (t_ Text) IsHorizontallyResizable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isHorizontallyResizable"))
	return rv
}/* debug [instance_properties/getter]: isHorizontallyResizable */


// A Boolean that controls whether the receiver changes its width to fit the width of its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/ishorizontallyresizable
func (t_ Text) SetIsHorizontallyResizable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsHorizontallyResizable:"), value)
}/* debug [instance_properties/setter]: isHorizontallyResizable */


// A Boolean that controls whether the receiver allows the user to apply attributes to specific ranges of the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/isrichtext
func (t_ Text) IsRichText() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isRichText"))
	return rv
}/* debug [instance_properties/getter]: isRichText */


// A Boolean that controls whether the receiver allows the user to apply attributes to specific ranges of the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/isrichtext
func (t_ Text) SetIsRichText(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsRichText:"), value)
}/* debug [instance_properties/setter]: isRichText */


// A Boolean value that indicates whether the receiver’s enclosing scroll view shows its ruler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/isrulervisible
func (t_ Text) IsRulerVisible() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isRulerVisible"))
	return rv
}/* debug [instance_properties/getter]: isRulerVisible */


// A Boolean value that indicates whether the receiver’s enclosing scroll view shows its ruler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/isrulervisible
func (t_ Text) SetIsRulerVisible(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsRulerVisible:"), value)
}/* debug [instance_properties/setter]: isRulerVisible */


// A Boolean that controls whether the receiver allows the user to select its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/isselectable
func (t_ Text) IsSelectable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isSelectable"))
	return rv
}/* debug [instance_properties/getter]: isSelectable */


// A Boolean that controls whether the receiver allows the user to select its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/isselectable
func (t_ Text) SetIsSelectable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsSelectable:"), value)
}/* debug [instance_properties/setter]: isSelectable */


// A Boolean that controls whether the receiver changes its height to fit the height of its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/isverticallyresizable
func (t_ Text) IsVerticallyResizable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isVerticallyResizable"))
	return rv
}/* debug [instance_properties/getter]: isVerticallyResizable */


// A Boolean that controls whether the receiver changes its height to fit the height of its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/isverticallyresizable
func (t_ Text) SetIsVerticallyResizable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsVerticallyResizable:"), value)
}/* debug [instance_properties/setter]: isVerticallyResizable */


// The receiver’s characters within
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/selectedrange
func (t_ Text) SelectedRange() corefoundation.Range {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("selectedRange"))
	return rv
}/* debug [instance_properties/getter]: selectedRange */


// The receiver’s characters within
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/selectedrange
func (t_ Text) SetSelectedRange(value corefoundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectedRange:"), value)
}/* debug [instance_properties/setter]: selectedRange */


// The characters of the receiver’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/string
func (t_ Text) String() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("string"))
	return rv
}/* debug [instance_properties/getter]: string */


// The characters of the receiver’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/string
func (t_ Text) SetString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setString:"), value)
}/* debug [instance_properties/setter]: string */


// The text color of all characters in the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/textcolor
func (t_ Text) TextColor() IColor {
	rv := objc.Send[Color](t_.ID, objc.Sel("textColor"))
	return rv
}/* debug [instance_properties/getter]: textColor */


// The text color of all characters in the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/textcolor
func (t_ Text) SetTextColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextColor:"), value)
}/* debug [instance_properties/setter]: textColor */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSText */



