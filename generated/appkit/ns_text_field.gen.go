// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	TextDidBeginEditing(notification unsafe.Pointer)
	TextShouldEndEditing(textObject unsafe.Pointer) bool
}

// Text the user can select or edit to send an action message to a target when the user presses the Return key.
//
// The class uses the class to implement its user interface. Text fields display text either as a static label or as an editable input field. The content of a text field is either plain text or a rich-text attributed string. Text fields also support line wrapping to display multiline text, and a variety of truncation styles if the content doesn’t fit the available space. The parent class, , provides the methods for setting the values of the text field, such as and . There are corresponding methods to retrieve values. In macOS 12 and later, if you explicitly call the property on your text field, the framework will revert to a compatibility mode that uses . The text view also switches to this compatibility mode when it encounters text content that’s not yet supported.
//
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
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/init(labelWithAttributedString:)
func NewTextFieldLabelWithAttributedString(attributedStringValue unsafe.Pointer) TextField {
	rv := objc.Send[TextField](objc.ID(getTextFieldClass().class), objc.Sel("labelWithAttributedString:"), attributedStringValue)
	return rv
}

// Creates a text field for use as a static label that displays styled text, doesn’t wrap, and doesn’t have selectable text.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/init(labelWithAttributedString:)
func (tc _TextFieldClass) LabelWithAttributedString(attributedStringValue unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("labelWithAttributedString:"), attributedStringValue)
	return rv
}

// Posts a notification to the default notification center that the text is about to go into edit mode.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/textDidBeginEditing(_:)
func (t_ TextField) TextDidBeginEditing(notification unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("textDidBeginEditing:"), notification)
}

// Performs validation on the text field’s new value.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/textShouldEndEditing(_:)
func (t_ TextField) TextShouldEndEditing(textObject unsafe.Pointer) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("textShouldEndEditing:"), textObject)
	return rv
}

// A Boolean value that controls whether single-line text fields tighten intercharacter spacing before truncating the text.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/allowsDefaultTighteningForTruncation
func (t_ TextField) AllowsDefaultTighteningForTruncation() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsDefaultTighteningForTruncation"))
	return rv
}

// SetAllowsDefaultTighteningForTruncation sets the value of the allowsDefaultTighteningForTruncation property.
// A Boolean value that controls whether single-line text fields tighten intercharacter spacing before truncating the text.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/allowsDefaultTighteningForTruncation
func (t_ TextField) SetAllowsDefaultTighteningForTruncation(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsDefaultTighteningForTruncation:"), value)
}

// A Boolean value that controls whether the user can change font attributes of the text field’s string.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/allowsEditingTextAttributes
func (t_ TextField) AllowsEditingTextAttributes() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsEditingTextAttributes"))
	return rv
}

// SetAllowsEditingTextAttributes sets the value of the allowsEditingTextAttributes property.
// A Boolean value that controls whether the user can change font attributes of the text field’s string.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/allowsEditingTextAttributes
func (t_ TextField) SetAllowsEditingTextAttributes(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsEditingTextAttributes:"), value)
}

// The text field’s delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/delegate
func (t_ TextField) Delegate() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("delegate"))
	return rv
}

// SetDelegate sets the value of the delegate property.
// The text field’s delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/delegate
func (t_ TextField) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean value that controls whether the text field’s cell draws a background color behind the text.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/drawsBackground
func (t_ TextField) DrawsBackground() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("drawsBackground"))
	return rv
}

// SetDrawsBackground sets the value of the drawsBackground property.
// A Boolean value that controls whether the text field’s cell draws a background color behind the text.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/drawsBackground
func (t_ TextField) SetDrawsBackground(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDrawsBackground:"), value)
}

// A Boolean value that controls whether the user can drag image files into the text field.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/importsGraphics
func (t_ TextField) ImportsGraphics() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("importsGraphics"))
	return rv
}

// SetImportsGraphics sets the value of the importsGraphics property.
// A Boolean value that controls whether the user can drag image files into the text field.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/importsGraphics
func (t_ TextField) SetImportsGraphics(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setImportsGraphics:"), value)
}

// A Boolean value that indicates whether the text field automatically completes text as the user types.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/isAutomaticTextCompletionEnabled
func (t_ TextField) AutomaticTextCompletionEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("automaticTextCompletionEnabled"))
	return rv
}

// SetAutomaticTextCompletionEnabled sets the value of the automaticTextCompletionEnabled property.
// A Boolean value that indicates whether the text field automatically completes text as the user types.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/isAutomaticTextCompletionEnabled
func (t_ TextField) SetAutomaticTextCompletionEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAutomaticTextCompletionEnabled:"), value)
}

// A Boolean value that controls whether the text field draws a bezeled background around its contents.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/isBezeled
func (t_ TextField) Bezeled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("bezeled"))
	return rv
}

// SetBezeled sets the value of the bezeled property.
// A Boolean value that controls whether the text field draws a bezeled background around its contents.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/isBezeled
func (t_ TextField) SetBezeled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBezeled:"), value)
}

// A Boolean value that controls whether the user can edit the value in the text field.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/isEditable
func (t_ TextField) Editable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("editable"))
	return rv
}

// SetEditable sets the value of the editable property.
// A Boolean value that controls whether the user can edit the value in the text field.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/isEditable
func (t_ TextField) SetEditable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setEditable:"), value)
}

// The attributed string the text field displays when empty to help the user understand the text field’s purpose.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/placeholderAttributedString
func (t_ TextField) PlaceholderAttributedString() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("placeholderAttributedString"))
	return rv
}

// SetPlaceholderAttributedString sets the value of the placeholderAttributedString property.
// The attributed string the text field displays when empty to help the user understand the text field’s purpose.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/placeholderAttributedString
func (t_ TextField) SetPlaceholderAttributedString(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPlaceholderAttributedString:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/placeholderStrings
func (t_ TextField) PlaceholderStrings() []string {
	rv := objc.Send[[]string](t_.ID, objc.Sel("placeholderStrings"))
	return rv
}

// SetPlaceholderStrings sets the value of the placeholderStrings property.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/placeholderStrings
func (t_ TextField) SetPlaceholderStrings(value []string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPlaceholderStrings:"), value)
}

// The maximum width of the text field’s intrinsic content size.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/preferredMaxLayoutWidth
func (t_ TextField) PreferredMaxLayoutWidth() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("preferredMaxLayoutWidth"))
	return rv
}

// SetPreferredMaxLayoutWidth sets the value of the preferredMaxLayoutWidth property.
// The maximum width of the text field’s intrinsic content size.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/preferredMaxLayoutWidth
func (t_ TextField) SetPreferredMaxLayoutWidth(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPreferredMaxLayoutWidth:"), value)
}
