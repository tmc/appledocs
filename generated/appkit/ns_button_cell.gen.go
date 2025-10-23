// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ButtonCell] class.
var (
	ButtonCellClass     _ButtonCellClass
	ButtonCellClassOnce sync.Once
)

func getButtonCellClass() _ButtonCellClass {
	ButtonCellClassOnce.Do(func() {
		ButtonCellClass = _ButtonCellClass{objc.GetClass("NSButtonCell")}
	})
	return ButtonCellClass
}

type _ButtonCellClass struct {
	class objc.Class
}

// An interface definition for the [ButtonCell] class.
type IButtonCell interface {
	IActionCell
	// properties:
	ImageScaling() ImageScaling
	SetImageScaling(value ImageScaling)
	AlternateImage() IImage
	SetAlternateImage(value IImage)
	AlternateTitle() string /* primitive/slice/pointer. */
	SetAlternateTitle(value string /* primitive/slice/pointer. */)
	AttributedAlternateTitle() objc.IObject /* cross-framework: AttributedString */
	SetAttributedAlternateTitle(value objc.IObject /* cross-framework: AttributedString */)
	AttributedTitle() objc.IObject /* cross-framework: AttributedString */
	SetAttributedTitle(value objc.IObject /* cross-framework: AttributedString */)
	BackgroundColor() IColor
	SetBackgroundColor(value IColor)
	BezelStyle() unsafe.Pointer
	SetBezelStyle(value unsafe.Pointer)
	GradientType() unsafe.Pointer
	SetGradientType(value unsafe.Pointer)
	HighlightsBy() unsafe.Pointer
	SetHighlightsBy(value unsafe.Pointer)
	ImageDimsWhenDisabled() bool /* primitive/slice/pointer. */
	SetImageDimsWhenDisabled(value bool /* primitive/slice/pointer. */)
	ImagePosition() unsafe.Pointer
	SetImagePosition(value unsafe.Pointer)
	IsOpaque() bool /* primitive/slice/pointer. */
	SetIsOpaque(value bool /* primitive/slice/pointer. */)
	IsTransparent() bool /* primitive/slice/pointer. */
	SetIsTransparent(value bool /* primitive/slice/pointer. */)
	KeyEquivalent() string /* primitive/slice/pointer. */
	SetKeyEquivalent(value string /* primitive/slice/pointer. */)
	KeyEquivalentFont() IFont
	SetKeyEquivalentFont(value IFont)
	KeyEquivalentModifierMask() unsafe.Pointer
	SetKeyEquivalentModifierMask(value unsafe.Pointer)
	ShowsBorderOnlyWhileMouseInside() bool /* primitive/slice/pointer. */
	SetShowsBorderOnlyWhileMouseInside(value bool /* primitive/slice/pointer. */)
	ShowsStateBy() unsafe.Pointer
	SetShowsStateBy(value unsafe.Pointer)
	Sound() ISound
	SetSound(value ISound)
	Title() string /* primitive/slice/pointer. */
	SetTitle(value string /* primitive/slice/pointer. */)
	Font() IFont
	SetFont(value IFont)
	IntValue() unsafe.Pointer
	SetIntValue(value unsafe.Pointer)
	ObjectValue() unsafe.Pointer
	SetObjectValue(value unsafe.Pointer)
	State() unsafe.Pointer
	SetState(value unsafe.Pointer)
	DoubleValue() float64 /* primitive/slice/pointer. */
	SetDoubleValue(value float64 /* primitive/slice/pointer. */)
	// methods:
}

// An object that defines the user interface of a button or other clickable region of a view.
//
// Setting the integer, float, double, or object value of an object results in a call to with the value converted to integer. In the case of , is equivalent to , and a non- object that doesn’t respond to sets the state to . Otherwise, the state is set to the object’s . Similarly, for most button types, querying the integer, float, double, or object value of an returns the current state in the requested representation. In the case of , this is an containing for on, for off, and integer value for the mixed state. For accelerator buttons (type or ) on systems that support pressure sensitivity, querying returns the amount of pressure applied while pressing the button. The configuration of an object controls how the button object appears and behaves, but it’s that sends a message when the control is clicked. For more information on the behavior of , see the and class specifications, and .


// An object that defines the user interface of a button or other clickable region of a view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell
type ButtonCell struct {
	ActionCell
}

// ButtonCellFrom constructs a [ButtonCell] from an unsafe.Pointer.
//
// An object that defines the user interface of a button or other clickable region of a view.
func ButtonCellFrom(ptr unsafe.Pointer) ButtonCell {
	return ButtonCell{
		ActionCell: ActionCellFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (bc _ButtonCellClass) Alloc() ButtonCell {
	rv := objc.Send[ButtonCell](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _ButtonCellClass) New() ButtonCell {
	rv := objc.Send[ButtonCell](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ ButtonCell) Init() ButtonCell {
	rv := objc.Send[ButtonCell](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ ButtonCell) Autorelease() ButtonCell {
	rv := objc.Send[ButtonCell](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewButtonCell creates a new ButtonCell instance.
func NewButtonCell() ButtonCell {
	return getButtonCellClass().New()
}



// The scale factor for the button’s image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/imageScaling
func (b_ ButtonCell) ImageScaling() ImageScaling {
	rv := objc.Send[ImageScaling](b_.ID, objc.Sel("imageScaling"))
	return rv
}


// The scale factor for the button’s image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/imageScaling
func (b_ ButtonCell) SetImageScaling(value ImageScaling) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setImageScaling:"), value)
}


// The image the button displays in its alternate state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/alternateimage
func (b_ ButtonCell) AlternateImage() IImage {
	rv := objc.Send[Image](b_.ID, objc.Sel("alternateImage"))
	return rv
}


// The image the button displays in its alternate state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/alternateimage
func (b_ ButtonCell) SetAlternateImage(value IImage) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAlternateImage:"), value)
}


// The string displayed by the button when it’s in its alternate state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/alternatetitle
func (b_ ButtonCell) AlternateTitle() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](b_.ID, objc.Sel("alternateTitle"))
	return rv
}


// The string displayed by the button when it’s in its alternate state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/alternatetitle
func (b_ ButtonCell) SetAlternateTitle(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAlternateTitle:"), objc.String(value))
}


// The title displayed by the button when it’s in its alternate state, as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/attributedalternatetitle
func (b_ ButtonCell) AttributedAlternateTitle() objc.IObject /* cross-framework: AttributedString */ {
	rv := objc.Send[AttributedString](b_.ID, objc.Sel("attributedAlternateTitle"))
	return rv
}


// The title displayed by the button when it’s in its alternate state, as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/attributedalternatetitle
func (b_ ButtonCell) SetAttributedAlternateTitle(value objc.IObject /* cross-framework: AttributedString */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAttributedAlternateTitle:"), value)
}


// The title displayed by the button when it’s in its normal state as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/attributedtitle
func (b_ ButtonCell) AttributedTitle() objc.IObject /* cross-framework: AttributedString */ {
	rv := objc.Send[AttributedString](b_.ID, objc.Sel("attributedTitle"))
	return rv
}


// The title displayed by the button when it’s in its normal state as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/attributedtitle
func (b_ ButtonCell) SetAttributedTitle(value objc.IObject /* cross-framework: AttributedString */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAttributedTitle:"), value)
}


// The background color of the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/backgroundcolor
func (b_ ButtonCell) BackgroundColor() IColor {
	rv := objc.Send[Color](b_.ID, objc.Sel("backgroundColor"))
	return rv
}


// The background color of the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/backgroundcolor
func (b_ ButtonCell) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBackgroundColor:"), value)
}


// The appearance of the button’s border, if it has one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/bezelstyle
func (b_ ButtonCell) BezelStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("bezelStyle"))
	return rv
}


// The appearance of the button’s border, if it has one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/bezelstyle
func (b_ ButtonCell) SetBezelStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBezelStyle:"), value)
}


// The gradient of the button’s border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/gradienttype
func (b_ ButtonCell) GradientType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("gradientType"))
	return rv
}


// The gradient of the button’s border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/gradienttype
func (b_ ButtonCell) SetGradientType(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setGradientType:"), value)
}


// A set of flags that indicate how the button highlights when it receives a mouse-down event (that is, when the button is pressed).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/highlightsby
func (b_ ButtonCell) HighlightsBy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("highlightsBy"))
	return rv
}


// A set of flags that indicate how the button highlights when it receives a mouse-down event (that is, when the button is pressed).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/highlightsby
func (b_ ButtonCell) SetHighlightsBy(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setHighlightsBy:"), value)
}


// A Boolean value that indicates if the button’s image and text appear “dim” when the button is disabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/imagedimswhendisabled
func (b_ ButtonCell) ImageDimsWhenDisabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](b_.ID, objc.Sel("imageDimsWhenDisabled"))
	return rv
}


// A Boolean value that indicates if the button’s image and text appear “dim” when the button is disabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/imagedimswhendisabled
func (b_ ButtonCell) SetImageDimsWhenDisabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setImageDimsWhenDisabled:"), value)
}


// The position of the button’s image relative to its title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/imageposition
func (b_ ButtonCell) ImagePosition() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("imagePosition"))
	return rv
}


// The position of the button’s image relative to its title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/imageposition
func (b_ ButtonCell) SetImagePosition(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setImagePosition:"), value)
}


// A Boolean value that indicates if the button is opaque.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/isopaque
func (b_ ButtonCell) IsOpaque() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](b_.ID, objc.Sel("isOpaque"))
	return rv
}


// A Boolean value that indicates if the button is opaque.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/isopaque
func (b_ ButtonCell) SetIsOpaque(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsOpaque:"), value)
}


// A Boolean value that indicates if the button is transparent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/istransparent
func (b_ ButtonCell) IsTransparent() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](b_.ID, objc.Sel("isTransparent"))
	return rv
}


// A Boolean value that indicates if the button is transparent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/istransparent
func (b_ ButtonCell) SetIsTransparent(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsTransparent:"), value)
}


// The button’s key-equivalent character.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/keyequivalent
func (b_ ButtonCell) KeyEquivalent() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](b_.ID, objc.Sel("keyEquivalent"))
	return rv
}


// The button’s key-equivalent character.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/keyequivalent
func (b_ ButtonCell) SetKeyEquivalent(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setKeyEquivalent:"), objc.String(value))
}


// The font used to draw the button’s key equivalent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/keyequivalentfont
func (b_ ButtonCell) KeyEquivalentFont() IFont {
	rv := objc.Send[Font](b_.ID, objc.Sel("keyEquivalentFont"))
	return rv
}


// The font used to draw the button’s key equivalent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/keyequivalentfont
func (b_ ButtonCell) SetKeyEquivalentFont(value IFont) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setKeyEquivalentFont:"), value)
}


// The mask that identifies the modifier keys for the button’s key equivalent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/keyequivalentmodifiermask
func (b_ ButtonCell) KeyEquivalentModifierMask() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("keyEquivalentModifierMask"))
	return rv
}


// The mask that identifies the modifier keys for the button’s key equivalent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/keyequivalentmodifiermask
func (b_ ButtonCell) SetKeyEquivalentModifierMask(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setKeyEquivalentModifierMask:"), value)
}


// A Boolean value that indicates if the button displays its border only when the pointer is over it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/showsborderonlywhilemouseinside
func (b_ ButtonCell) ShowsBorderOnlyWhileMouseInside() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](b_.ID, objc.Sel("showsBorderOnlyWhileMouseInside"))
	return rv
}


// A Boolean value that indicates if the button displays its border only when the pointer is over it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/showsborderonlywhilemouseinside
func (b_ ButtonCell) SetShowsBorderOnlyWhileMouseInside(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setShowsBorderOnlyWhileMouseInside:"), value)
}


// The flags that indicate how the button cell shows its alternate state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/showsstateby
func (b_ ButtonCell) ShowsStateBy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("showsStateBy"))
	return rv
}


// The flags that indicate how the button cell shows its alternate state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/showsstateby
func (b_ ButtonCell) SetShowsStateBy(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setShowsStateBy:"), value)
}


// The sound that’s played when the user presses the button (that is during a mouse-down event).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/sound
func (b_ ButtonCell) Sound() ISound {
	rv := objc.Send[Sound](b_.ID, objc.Sel("sound"))
	return rv
}


// The sound that’s played when the user presses the button (that is during a mouse-down event).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/sound
func (b_ ButtonCell) SetSound(value ISound) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSound:"), value)
}


// The title displayed on the button when it’s in its normal state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/title
func (b_ ButtonCell) Title() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](b_.ID, objc.Sel("title"))
	return rv
}


// The title displayed on the button when it’s in its normal state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/title
func (b_ ButtonCell) SetTitle(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTitle:"), objc.String(value))
}


// The font that the cell uses to display text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/font
func (b_ ButtonCell) Font() IFont {
	rv := objc.Send[Font](b_.ID, objc.Sel("font"))
	return rv
}


// The font that the cell uses to display text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/font
func (b_ ButtonCell) SetFont(value IFont) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setFont:"), value)
}


// The cell’s value as an integer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/intvalue
func (b_ ButtonCell) IntValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("intValue"))
	return rv
}


// The cell’s value as an integer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/intvalue
func (b_ ButtonCell) SetIntValue(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIntValue:"), value)
}


// The cell’s value as an Objective-C object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/objectvalue
func (b_ ButtonCell) ObjectValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("objectValue"))
	return rv
}


// The cell’s value as an Objective-C object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/objectvalue
func (b_ ButtonCell) SetObjectValue(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setObjectValue:"), value)
}


// The cell’s current state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/state
func (b_ ButtonCell) State() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("state"))
	return rv
}


// The cell’s current state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/state
func (b_ ButtonCell) SetState(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setState:"), value)
}


// The value of the receiver’s cell as a double-precision floating-point number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/doublevalue
func (b_ ButtonCell) DoubleValue() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](b_.ID, objc.Sel("doubleValue"))
	return rv
}


// The value of the receiver’s cell as a double-precision floating-point number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/doublevalue
func (b_ ButtonCell) SetDoubleValue(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setDoubleValue:"), value)
}



