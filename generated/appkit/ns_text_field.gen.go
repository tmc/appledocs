// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextField] class.
var (
	TextFieldClass     _TextFieldClass
	TextFieldClassOnce sync.Once
)

func getTextFieldClass() _TextFieldClass {
	TextFieldClassOnce.Do(func() {
		TextFieldClass = _TextFieldClass{objc.GetClass("NSTextField")}
	})
	return TextFieldClass
}

type _TextFieldClass struct {
	class objc.Class
}

// An interface definition for the [TextField] class.
type ITextField interface {
	IControl
	// properties:
	AcceptsFirstResponder() bool
	AllowsCharacterPickerTouchBarItem() bool
	SetAllowsCharacterPickerTouchBarItem(value bool)
	AllowsDefaultTighteningForTruncation() bool
	SetAllowsDefaultTighteningForTruncation(value bool)
	AllowsEditingTextAttributes() bool
	SetAllowsEditingTextAttributes(value bool)
	AllowsWritingTools() bool
	SetAllowsWritingTools(value bool)
	AllowsWritingToolsAffordance() bool
	SetAllowsWritingToolsAffordance(value bool)
	BackgroundColor() IColor
	SetBackgroundColor(value IColor)
	BezelStyle() TextFieldBezelStyle
	SetBezelStyle(value TextFieldBezelStyle)
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	ImportsGraphics() bool
	SetImportsGraphics(value bool)
	AutomaticTextCompletionEnabled() bool
	SetAutomaticTextCompletionEnabled(value bool)
	Bezeled() bool
	SetBezeled(value bool)
	Bordered() bool
	SetBordered(value bool)
	Editable() bool
	SetEditable(value bool)
	Selectable() bool
	SetSelectable(value bool)
	LineBreakStrategy() LineBreakStrategy
	SetLineBreakStrategy(value LineBreakStrategy)
	MaximumNumberOfLines() int
	SetMaximumNumberOfLines(value int)
	PlaceholderAttributedString() foundation.AttributedString
	SetPlaceholderAttributedString(value foundation.AttributedString)
	PlaceholderAttributedStrings() []foundation.AttributedString
	SetPlaceholderAttributedStrings(value []foundation.AttributedString)
	PlaceholderString() objc.IObject /* cross-framework: NSString */
	SetPlaceholderString(value objc.IObject /* cross-framework: NSString */)
	PlaceholderStrings() []string
	SetPlaceholderStrings(value []string)
	PreferredMaxLayoutWidth() float64
	SetPreferredMaxLayoutWidth(value float64)
	ResolvesNaturalAlignmentWithBaseWritingDirection() bool
	SetResolvesNaturalAlignmentWithBaseWritingDirection(value bool)
	TextColor() IColor
	SetTextColor(value IColor)
	DoubleValue() float64
	SetDoubleValue(value float64)
	StringValue() objc.IObject /* cross-framework: NSString */
	SetStringValue(value objc.IObject /* cross-framework: NSString */)
	IsAutomaticTextCompletionEnabled() bool
	SetIsAutomaticTextCompletionEnabled(value bool)
	IsBezeled() bool
	SetIsBezeled(value bool)
	IsBordered() bool
	SetIsBordered(value bool)
	IsEditable() bool
	SetIsEditable(value bool)
	IsSelectable() bool
	SetIsSelectable(value bool)
	SuggestionsDelegate() TextSuggestionsDelegate /* not a class type */
	SetSuggestionsDelegate(value TextSuggestionsDelegate /* not a class type */)
	// methods:
	SelectText(sender objc.IObject)
	TextDidBeginEditing(notification foundation.Notification)
	TextDidChange(notification foundation.Notification)
	TextDidEndEditing(notification foundation.Notification)
	TextShouldBeginEditing(textObject IText) bool
	TextShouldEndEditing(textObject IText) bool
}

// Text the user can select or edit to send an action message to a target when the user presses the Return key.
//
// The class uses the class to implement its user interface. Text fields display text either as a static label or as an editable input field. The content of a text field is either plain text or a rich-text attributed string. Text fields also support line wrapping to display multiline text, and a variety of truncation styles if the content doesn’t fit the available space. The parent class, , provides the methods for setting the values of the text field, such as and . There are corresponding methods to retrieve values. In macOS 12 and later, if you explicitly call the property on your text field, the framework will revert to a compatibility mode that uses . The text view also switches to this compatibility mode when it encounters text content that’s not yet supported.


// Text the user can select or edit to send an action message to a target when the user presses the Return key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField
type TextField struct {
	Control
}

// TextFieldFrom constructs a [TextField] from an unsafe.Pointer.
//
// Text the user can select or edit to send an action message to a target when the user presses the Return key.
func TextFieldFrom(ptr unsafe.Pointer) TextField {
	return TextField{
		Control: ControlFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _TextFieldClass) Alloc() TextField {
	rv := objc.Send[TextField](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextFieldClass) New() TextField {
	rv := objc.Send[TextField](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextField) Init() TextField {
	rv := objc.Send[TextField](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextField) Autorelease() TextField {
	rv := objc.Send[TextField](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextField creates a new TextField instance.
func NewTextField() TextField {
	return getTextFieldClass().New()
}



// Creates a text field for use as a static label that displays styled text, doesn’t wrap, and doesn’t have selectable text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/init(labelWithAttributedString:)
func NewTextFieldLabelWithAttributedString(attributedStringValue foundation.AttributedString) TextField {
	rv := objc.Send[TextField](objc.ID(getTextFieldClass().class), objc.Sel("labelWithAttributedString:"), attributedStringValue)
	return rv
}


// Initializes a text field for use as a static label that uses the system default font, doesn’t wrap, and doesn’t have selectable text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/init(labelWithString:)
func NewTextFieldLabelWithString(stringValue objc.IObject /* cross-framework: NSString */) TextField {
	rv := objc.Send[TextField](objc.ID(getTextFieldClass().class), objc.Sel("labelWithString:"), stringValue)
	return rv
}


// Initializes a single-line editable text field for user input using the system default font and standard visual appearance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/init(string:)
func NewTextFieldWithString(stringValue objc.IObject /* cross-framework: NSString */) TextField {
	rv := objc.Send[TextField](objc.ID(getTextFieldClass().class), objc.Sel("textFieldWithString:"), stringValue)
	return rv
}


// Initializes a text field for use as a multiline static label with selectable text that uses the system default font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/init(wrappingLabelWithString:)
func NewTextFieldWrappingLabelWithString(stringValue objc.IObject /* cross-framework: NSString */) TextField {
	rv := objc.Send[TextField](objc.ID(getTextFieldClass().class), objc.Sel("wrappingLabelWithString:"), stringValue)
	return rv
}



// Creates a text field for use as a static label that displays styled text, doesn’t wrap, and doesn’t have selectable text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/init(labelWithAttributedString:)
func (tc _TextFieldClass) LabelWithAttributedString(attributedStringValue foundation.AttributedString) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("labelWithAttributedString:"), attributedStringValue)
	return rv
}


// Initializes a text field for use as a static label that uses the system default font, doesn’t wrap, and doesn’t have selectable text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/init(labelWithString:)
func (tc _TextFieldClass) LabelWithString(stringValue objc.IObject /* cross-framework: NSString */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("labelWithString:"), stringValue)
	return rv
}


// Initializes a single-line editable text field for user input using the system default font and standard visual appearance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/init(string:)
func (tc _TextFieldClass) TextFieldWithString(stringValue objc.IObject /* cross-framework: NSString */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("textFieldWithString:"), stringValue)
	return rv
}


// Initializes a text field for use as a multiline static label with selectable text that uses the system default font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/init(wrappingLabelWithString:)
func (tc _TextFieldClass) WrappingLabelWithString(stringValue objc.IObject /* cross-framework: NSString */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("wrappingLabelWithString:"), stringValue)
	return rv
}


// Ends editing in the text field and, if it’s selectable, selects the entire text content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/selectText(_:)
func (t_ TextField) SelectText(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("selectText:"), sender)
}


// Posts a notification to the default notification center that the text is about to go into edit mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/textDidBeginEditing(_:)
func (t_ TextField) TextDidBeginEditing(notification foundation.Notification) {
	objc.Send[objc.ID](t_.ID, objc.Sel("textDidBeginEditing:"), notification)
}


// Posts a notification when the text changes, and forwards the message to the text field’s cell if it responds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/textDidChange(_:)
func (t_ TextField) TextDidChange(notification foundation.Notification) {
	objc.Send[objc.ID](t_.ID, objc.Sel("textDidChange:"), notification)
}


// Posts a notification when the text is no longer in edit mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/textDidEndEditing(_:)
func (t_ TextField) TextDidEndEditing(notification foundation.Notification) {
	objc.Send[objc.ID](t_.ID, objc.Sel("textDidEndEditing:"), notification)
}


// Requests permission to begin editing a text object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/textShouldBeginEditing(_:)
func (t_ TextField) TextShouldBeginEditing(textObject IText) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("textShouldBeginEditing:"), textObject)
	return rv
}


// Performs validation on the text field’s new value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/textShouldEndEditing(_:)
func (t_ TextField) TextShouldEndEditing(textObject IText) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("textShouldEndEditing:"), textObject)
	return rv
}


// A Boolean value that indicates whether the text field is editable and accepts first responder status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/acceptsFirstResponder
func (t_ TextField) AcceptsFirstResponder() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("acceptsFirstResponder"))
	return rv
}


// A Boolean value that controls whether the Touch Bar displays the character picker item for rich text fields.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/allowsCharacterPickerTouchBarItem
func (t_ TextField) AllowsCharacterPickerTouchBarItem() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsCharacterPickerTouchBarItem"))
	return rv
}


// A Boolean value that controls whether the Touch Bar displays the character picker item for rich text fields.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/allowsCharacterPickerTouchBarItem
func (t_ TextField) SetAllowsCharacterPickerTouchBarItem(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsCharacterPickerTouchBarItem:"), value)
}


// A Boolean value that controls whether single-line text fields tighten intercharacter spacing before truncating the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/allowsDefaultTighteningForTruncation
func (t_ TextField) AllowsDefaultTighteningForTruncation() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsDefaultTighteningForTruncation"))
	return rv
}


// A Boolean value that controls whether single-line text fields tighten intercharacter spacing before truncating the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/allowsDefaultTighteningForTruncation
func (t_ TextField) SetAllowsDefaultTighteningForTruncation(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsDefaultTighteningForTruncation:"), value)
}


// A Boolean value that controls whether the user can change font attributes of the text field’s string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/allowsEditingTextAttributes
func (t_ TextField) AllowsEditingTextAttributes() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsEditingTextAttributes"))
	return rv
}


// A Boolean value that controls whether the user can change font attributes of the text field’s string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/allowsEditingTextAttributes
func (t_ TextField) SetAllowsEditingTextAttributes(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsEditingTextAttributes:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/allowsWritingTools
func (t_ TextField) AllowsWritingTools() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsWritingTools"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/allowsWritingTools
func (t_ TextField) SetAllowsWritingTools(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsWritingTools:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/allowsWritingToolsAffordance
func (t_ TextField) AllowsWritingToolsAffordance() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsWritingToolsAffordance"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/allowsWritingToolsAffordance
func (t_ TextField) SetAllowsWritingToolsAffordance(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsWritingToolsAffordance:"), value)
}


// The color of the background the text field’s cell draws behind the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/backgroundColor
func (t_ TextField) BackgroundColor() IColor {
	rv := objc.Send[Color](t_.ID, objc.Sel("backgroundColor"))
	return rv
}


// The color of the background the text field’s cell draws behind the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/backgroundColor
func (t_ TextField) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBackgroundColor:"), value)
}


// The text field’s bezel style, square or rounded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/bezelStyle-swift.property
func (t_ TextField) BezelStyle() TextFieldBezelStyle {
	rv := objc.Send[TextFieldBezelStyle](t_.ID, objc.Sel("bezelStyle"))
	return rv
}


// The text field’s bezel style, square or rounded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/bezelStyle-swift.property
func (t_ TextField) SetBezelStyle(value TextFieldBezelStyle) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBezelStyle:"), value)
}


// The text field’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/delegate
func (t_ TextField) Delegate() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("delegate"))
	return rv
}


// The text field’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/delegate
func (t_ TextField) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}


// A Boolean value that controls whether the text field’s cell draws a background color behind the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/drawsBackground
func (t_ TextField) DrawsBackground() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("drawsBackground"))
	return rv
}


// A Boolean value that controls whether the text field’s cell draws a background color behind the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/drawsBackground
func (t_ TextField) SetDrawsBackground(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDrawsBackground:"), value)
}


// A Boolean value that controls whether the user can drag image files into the text field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/importsGraphics
func (t_ TextField) ImportsGraphics() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("importsGraphics"))
	return rv
}


// A Boolean value that controls whether the user can drag image files into the text field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/importsGraphics
func (t_ TextField) SetImportsGraphics(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setImportsGraphics:"), value)
}


// A Boolean value that indicates whether the text field automatically completes text as the user types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/isAutomaticTextCompletionEnabled
func (t_ TextField) AutomaticTextCompletionEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("automaticTextCompletionEnabled"))
	return rv
}


// A Boolean value that indicates whether the text field automatically completes text as the user types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/isAutomaticTextCompletionEnabled
func (t_ TextField) SetAutomaticTextCompletionEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAutomaticTextCompletionEnabled:"), value)
}


// A Boolean value that controls whether the text field draws a bezeled background around its contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/isBezeled
func (t_ TextField) Bezeled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("bezeled"))
	return rv
}


// A Boolean value that controls whether the text field draws a bezeled background around its contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/isBezeled
func (t_ TextField) SetBezeled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBezeled:"), value)
}


// A Boolean value that controls whether the text field draws a solid black border around its contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/isBordered
func (t_ TextField) Bordered() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("bordered"))
	return rv
}


// A Boolean value that controls whether the text field draws a solid black border around its contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/isBordered
func (t_ TextField) SetBordered(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBordered:"), value)
}


// A Boolean value that controls whether the user can edit the value in the text field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/isEditable
func (t_ TextField) Editable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("editable"))
	return rv
}


// A Boolean value that controls whether the user can edit the value in the text field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/isEditable
func (t_ TextField) SetEditable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setEditable:"), value)
}


// A Boolean value that determines whether the user can select the content of the text field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/isSelectable
func (t_ TextField) Selectable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("selectable"))
	return rv
}


// A Boolean value that determines whether the user can select the content of the text field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/isSelectable
func (t_ TextField) SetSelectable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectable:"), value)
}


// The strategy that the system uses to break lines when laying out multiple lines of text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/lineBreakStrategy
func (t_ TextField) LineBreakStrategy() LineBreakStrategy {
	rv := objc.Send[LineBreakStrategy](t_.ID, objc.Sel("lineBreakStrategy"))
	return rv
}


// The strategy that the system uses to break lines when laying out multiple lines of text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/lineBreakStrategy
func (t_ TextField) SetLineBreakStrategy(value LineBreakStrategy) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLineBreakStrategy:"), value)
}


// The maximum number of lines a wrapping text field displays before clipping or truncating the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/maximumNumberOfLines
func (t_ TextField) MaximumNumberOfLines() int {
	rv := objc.Send[int](t_.ID, objc.Sel("maximumNumberOfLines"))
	return rv
}


// The maximum number of lines a wrapping text field displays before clipping or truncating the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/maximumNumberOfLines
func (t_ TextField) SetMaximumNumberOfLines(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMaximumNumberOfLines:"), value)
}


// The attributed string the text field displays when empty to help the user understand the text field’s purpose.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/placeholderAttributedString
func (t_ TextField) PlaceholderAttributedString() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](t_.ID, objc.Sel("placeholderAttributedString"))
	return rv
}


// The attributed string the text field displays when empty to help the user understand the text field’s purpose.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/placeholderAttributedString
func (t_ TextField) SetPlaceholderAttributedString(value foundation.AttributedString) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPlaceholderAttributedString:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/placeholderAttributedStrings
func (t_ TextField) PlaceholderAttributedStrings() []foundation.AttributedString {
	rv := objc.Send[[]foundation.AttributedString](t_.ID, objc.Sel("placeholderAttributedStrings"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/placeholderAttributedStrings
func (t_ TextField) SetPlaceholderAttributedStrings(value []foundation.AttributedString) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](t_.ID, objc.Sel("setPlaceholderAttributedStrings:"), nsArray)
}


// The string the text field displays when empty to help the user understand the text field’s purpose.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/placeholderString
func (t_ TextField) PlaceholderString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("placeholderString"))
	return rv
}


// The string the text field displays when empty to help the user understand the text field’s purpose.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/placeholderString
func (t_ TextField) SetPlaceholderString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPlaceholderString:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/placeholderStrings
func (t_ TextField) PlaceholderStrings() []string {
	rv := objc.Send[[]string](t_.ID, objc.Sel("placeholderStrings"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/placeholderStrings
func (t_ TextField) SetPlaceholderStrings(value []string) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](t_.ID, objc.Sel("setPlaceholderStrings:"), nsArray)
}


// The maximum width of the text field’s intrinsic content size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/preferredMaxLayoutWidth
func (t_ TextField) PreferredMaxLayoutWidth() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("preferredMaxLayoutWidth"))
	return rv
}


// The maximum width of the text field’s intrinsic content size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/preferredMaxLayoutWidth
func (t_ TextField) SetPreferredMaxLayoutWidth(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPreferredMaxLayoutWidth:"), value)
}


// Specifies the behavior for resolving to the visual alignment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/resolvesNaturalAlignmentWithBaseWritingDirection
func (t_ TextField) ResolvesNaturalAlignmentWithBaseWritingDirection() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("resolvesNaturalAlignmentWithBaseWritingDirection"))
	return rv
}


// Specifies the behavior for resolving to the visual alignment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/resolvesNaturalAlignmentWithBaseWritingDirection
func (t_ TextField) SetResolvesNaturalAlignmentWithBaseWritingDirection(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setResolvesNaturalAlignmentWithBaseWritingDirection:"), value)
}


// The color of the text field’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/textColor
func (t_ TextField) TextColor() IColor {
	rv := objc.Send[Color](t_.ID, objc.Sel("textColor"))
	return rv
}


// The color of the text field’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/textColor
func (t_ TextField) SetTextColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextColor:"), value)
}


// The value of the receiver’s cell as a double-precision floating-point number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/doublevalue
func (t_ TextField) DoubleValue() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("doubleValue"))
	return rv
}


// The value of the receiver’s cell as a double-precision floating-point number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/doublevalue
func (t_ TextField) SetDoubleValue(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDoubleValue:"), value)
}


// The value of the receiver’s cell as an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/stringvalue
func (t_ TextField) StringValue() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("stringValue"))
	return rv
}


// The value of the receiver’s cell as an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/stringvalue
func (t_ TextField) SetStringValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStringValue:"), value)
}


// A Boolean value that indicates whether the text field automatically completes text as the user types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/isautomatictextcompletionenabled
func (t_ TextField) IsAutomaticTextCompletionEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isAutomaticTextCompletionEnabled"))
	return rv
}


// A Boolean value that indicates whether the text field automatically completes text as the user types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/isautomatictextcompletionenabled
func (t_ TextField) SetIsAutomaticTextCompletionEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsAutomaticTextCompletionEnabled:"), value)
}


// A Boolean value that controls whether the text field draws a bezeled background around its contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/isbezeled
func (t_ TextField) IsBezeled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isBezeled"))
	return rv
}


// A Boolean value that controls whether the text field draws a bezeled background around its contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/isbezeled
func (t_ TextField) SetIsBezeled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsBezeled:"), value)
}


// A Boolean value that controls whether the text field draws a solid black border around its contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/isbordered
func (t_ TextField) IsBordered() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isBordered"))
	return rv
}


// A Boolean value that controls whether the text field draws a solid black border around its contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/isbordered
func (t_ TextField) SetIsBordered(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsBordered:"), value)
}


// A Boolean value that controls whether the user can edit the value in the text field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/iseditable
func (t_ TextField) IsEditable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isEditable"))
	return rv
}


// A Boolean value that controls whether the user can edit the value in the text field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/iseditable
func (t_ TextField) SetIsEditable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsEditable:"), value)
}


// A Boolean value that determines whether the user can select the content of the text field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/isselectable
func (t_ TextField) IsSelectable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isSelectable"))
	return rv
}


// A Boolean value that determines whether the user can select the content of the text field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/isselectable
func (t_ TextField) SetIsSelectable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsSelectable:"), value)
}


// The delegate that provides text suggestions for the receiving text field and responds to the user highlighting and selecting items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/suggestionsdelegate
func (t_ TextField) SuggestionsDelegate() TextSuggestionsDelegate /* not a class type */ {
	rv := objc.Send[TextSuggestionsDelegate](t_.ID, objc.Sel("suggestionsDelegate"))
	return rv
}


// The delegate that provides text suggestions for the receiving text field and responds to the user highlighting and selecting items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/suggestionsdelegate
func (t_ TextField) SetSuggestionsDelegate(value TextSuggestionsDelegate /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSuggestionsDelegate:"), value)
}


