// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	DoubleValue() float64
	SetDoubleValue(value float64)
	StringValue() objc.IObject /* cross-framework: NSString */
	SetStringValue(value objc.IObject /* cross-framework: NSString */)
	AcceptsFirstResponder() bool
	SetAcceptsFirstResponder(value bool)
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
	BackgroundColor() objc.IObject /* cross-framework: Color */
	SetBackgroundColor(value objc.IObject /* cross-framework: Color */)
	BezelStyle() unsafe.Pointer
	SetBezelStyle(value unsafe.Pointer)
	Delegate() TextFieldDelegate /* not a class type */
	SetDelegate(value TextFieldDelegate /* not a class type */)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	ImportsGraphics() bool
	SetImportsGraphics(value bool)
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
	LineBreakStrategy() unsafe.Pointer
	SetLineBreakStrategy(value unsafe.Pointer)
	MaximumNumberOfLines() int
	SetMaximumNumberOfLines(value int)
	PlaceholderAttributedString() objc.IObject /* cross-framework: AttributedString */
	SetPlaceholderAttributedString(value objc.IObject /* cross-framework: AttributedString */)
	PlaceholderAttributedStrings() objc.IObject /* cross-framework: AttributedString */
	SetPlaceholderAttributedStrings(value objc.IObject /* cross-framework: AttributedString */)
	PlaceholderString() objc.IObject /* cross-framework: NSString */
	SetPlaceholderString(value objc.IObject /* cross-framework: NSString */)
	PlaceholderStrings() objc.IObject /* cross-framework: NSString */
	SetPlaceholderStrings(value objc.IObject /* cross-framework: NSString */)
	PreferredMaxLayoutWidth() float64
	SetPreferredMaxLayoutWidth(value float64)
	ResolvesNaturalAlignmentWithBaseWritingDirection() bool
	SetResolvesNaturalAlignmentWithBaseWritingDirection(value bool)
	SuggestionsDelegate() TextSuggestionsDelegate /* not a class type */
	SetSuggestionsDelegate(value TextSuggestionsDelegate /* not a class type */)
	TextColor() objc.IObject /* cross-framework: Color */
	SetTextColor(value objc.IObject /* cross-framework: Color */)
	// methods:
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


// A Boolean value that indicates whether the text field is editable and accepts first responder status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/acceptsfirstresponder
func (t_ TextField) AcceptsFirstResponder() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("acceptsFirstResponder"))
	return rv
}


// A Boolean value that indicates whether the text field is editable and accepts first responder status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/acceptsfirstresponder
func (t_ TextField) SetAcceptsFirstResponder(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAcceptsFirstResponder:"), value)
}


// A Boolean value that controls whether the Touch Bar displays the character picker item for rich text fields.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/allowscharacterpickertouchbaritem
func (t_ TextField) AllowsCharacterPickerTouchBarItem() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsCharacterPickerTouchBarItem"))
	return rv
}


// A Boolean value that controls whether the Touch Bar displays the character picker item for rich text fields.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/allowscharacterpickertouchbaritem
func (t_ TextField) SetAllowsCharacterPickerTouchBarItem(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsCharacterPickerTouchBarItem:"), value)
}


// A Boolean value that controls whether single-line text fields tighten intercharacter spacing before truncating the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/allowsdefaulttighteningfortruncation
func (t_ TextField) AllowsDefaultTighteningForTruncation() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsDefaultTighteningForTruncation"))
	return rv
}


// A Boolean value that controls whether single-line text fields tighten intercharacter spacing before truncating the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/allowsdefaulttighteningfortruncation
func (t_ TextField) SetAllowsDefaultTighteningForTruncation(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsDefaultTighteningForTruncation:"), value)
}


// A Boolean value that controls whether the user can change font attributes of the text field’s string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/allowseditingtextattributes
func (t_ TextField) AllowsEditingTextAttributes() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsEditingTextAttributes"))
	return rv
}


// A Boolean value that controls whether the user can change font attributes of the text field’s string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/allowseditingtextattributes
func (t_ TextField) SetAllowsEditingTextAttributes(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsEditingTextAttributes:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/allowswritingtools
func (t_ TextField) AllowsWritingTools() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsWritingTools"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/allowswritingtools
func (t_ TextField) SetAllowsWritingTools(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsWritingTools:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/allowswritingtoolsaffordance
func (t_ TextField) AllowsWritingToolsAffordance() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsWritingToolsAffordance"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/allowswritingtoolsaffordance
func (t_ TextField) SetAllowsWritingToolsAffordance(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsWritingToolsAffordance:"), value)
}


// The color of the background the text field’s cell draws behind the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/backgroundcolor
func (t_ TextField) BackgroundColor() objc.IObject /* cross-framework: Color */ {
	rv := objc.Send[Color](t_.ID, objc.Sel("backgroundColor"))
	return rv
}


// The color of the background the text field’s cell draws behind the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/backgroundcolor
func (t_ TextField) SetBackgroundColor(value objc.IObject /* cross-framework: Color */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBackgroundColor:"), value)
}


// The text field’s bezel style, square or rounded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/bezelstyle-swift.property
func (t_ TextField) BezelStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("bezelStyle"))
	return rv
}


// The text field’s bezel style, square or rounded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/bezelstyle-swift.property
func (t_ TextField) SetBezelStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBezelStyle:"), value)
}


// The text field’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/delegate
func (t_ TextField) Delegate() TextFieldDelegate /* not a class type */ {
	rv := objc.Send[TextFieldDelegate](t_.ID, objc.Sel("delegate"))
	return rv
}


// The text field’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/delegate
func (t_ TextField) SetDelegate(value TextFieldDelegate /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}


// A Boolean value that controls whether the text field’s cell draws a background color behind the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/drawsbackground
func (t_ TextField) DrawsBackground() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("drawsBackground"))
	return rv
}


// A Boolean value that controls whether the text field’s cell draws a background color behind the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/drawsbackground
func (t_ TextField) SetDrawsBackground(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDrawsBackground:"), value)
}


// A Boolean value that controls whether the user can drag image files into the text field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/importsgraphics
func (t_ TextField) ImportsGraphics() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("importsGraphics"))
	return rv
}


// A Boolean value that controls whether the user can drag image files into the text field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/importsgraphics
func (t_ TextField) SetImportsGraphics(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setImportsGraphics:"), value)
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


// The strategy that the system uses to break lines when laying out multiple lines of text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/linebreakstrategy
func (t_ TextField) LineBreakStrategy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("lineBreakStrategy"))
	return rv
}


// The strategy that the system uses to break lines when laying out multiple lines of text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/linebreakstrategy
func (t_ TextField) SetLineBreakStrategy(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLineBreakStrategy:"), value)
}


// The maximum number of lines a wrapping text field displays before clipping or truncating the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/maximumnumberoflines
func (t_ TextField) MaximumNumberOfLines() int {
	rv := objc.Send[int](t_.ID, objc.Sel("maximumNumberOfLines"))
	return rv
}


// The maximum number of lines a wrapping text field displays before clipping or truncating the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/maximumnumberoflines
func (t_ TextField) SetMaximumNumberOfLines(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMaximumNumberOfLines:"), value)
}


// The attributed string the text field displays when empty to help the user understand the text field’s purpose.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/placeholderattributedstring
func (t_ TextField) PlaceholderAttributedString() objc.IObject /* cross-framework: AttributedString */ {
	rv := objc.Send[foundation.AttributedString](t_.ID, objc.Sel("placeholderAttributedString"))
	return rv
}


// The attributed string the text field displays when empty to help the user understand the text field’s purpose.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/placeholderattributedstring
func (t_ TextField) SetPlaceholderAttributedString(value objc.IObject /* cross-framework: AttributedString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPlaceholderAttributedString:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/placeholderattributedstrings
func (t_ TextField) PlaceholderAttributedStrings() objc.IObject /* cross-framework: AttributedString */ {
	rv := objc.Send[foundation.AttributedString](t_.ID, objc.Sel("placeholderAttributedStrings"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/placeholderattributedstrings
func (t_ TextField) SetPlaceholderAttributedStrings(value objc.IObject /* cross-framework: AttributedString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPlaceholderAttributedStrings:"), value)
}


// The string the text field displays when empty to help the user understand the text field’s purpose.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/placeholderstring
func (t_ TextField) PlaceholderString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("placeholderString"))
	return rv
}


// The string the text field displays when empty to help the user understand the text field’s purpose.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/placeholderstring
func (t_ TextField) SetPlaceholderString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPlaceholderString:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/placeholderstrings
func (t_ TextField) PlaceholderStrings() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("placeholderStrings"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/placeholderstrings
func (t_ TextField) SetPlaceholderStrings(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPlaceholderStrings:"), value)
}


// The maximum width of the text field’s intrinsic content size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/preferredmaxlayoutwidth
func (t_ TextField) PreferredMaxLayoutWidth() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("preferredMaxLayoutWidth"))
	return rv
}


// The maximum width of the text field’s intrinsic content size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/preferredmaxlayoutwidth
func (t_ TextField) SetPreferredMaxLayoutWidth(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPreferredMaxLayoutWidth:"), value)
}


// Specifies the behavior for resolving
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/resolvesnaturalalignmentwithbasewritingdirection
func (t_ TextField) ResolvesNaturalAlignmentWithBaseWritingDirection() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("resolvesNaturalAlignmentWithBaseWritingDirection"))
	return rv
}


// Specifies the behavior for resolving
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/resolvesnaturalalignmentwithbasewritingdirection
func (t_ TextField) SetResolvesNaturalAlignmentWithBaseWritingDirection(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setResolvesNaturalAlignmentWithBaseWritingDirection:"), value)
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


// The color of the text field’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/textcolor
func (t_ TextField) TextColor() objc.IObject /* cross-framework: Color */ {
	rv := objc.Send[Color](t_.ID, objc.Sel("textColor"))
	return rv
}


// The color of the text field’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/textcolor
func (t_ TextField) SetTextColor(value objc.IObject /* cross-framework: Color */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextColor:"), value)
}



