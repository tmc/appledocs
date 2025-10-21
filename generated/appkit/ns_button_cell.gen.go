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
	SetKeyEquivalentFontSize(fontName string, fontSize float64)
}

// An object that defines the user interface of a button or other clickable region of a view.
//
// Setting the integer, float, double, or object value of an object results in a call to with the value converted to integer. In the case of , is equivalent to , and a non- object that doesn’t respond to sets the state to . Otherwise, the state is set to the object’s . Similarly, for most button types, querying the integer, float, double, or object value of an returns the current state in the requested representation. In the case of , this is an containing for on, for off, and integer value for the mixed state. For accelerator buttons (type or ) on systems that support pressure sensitivity, querying returns the amount of pressure applied while pressing the button. The configuration of an object controls how the button object appears and behaves, but it’s that sends a message when the control is clicked. For more information on the behavior of , see the and class specifications, and .
//
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


// Sets by name and size of the font used to draw the key equivalent.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/setKeyEquivalentFont(_:size:)
func (b_ ButtonCell) SetKeyEquivalentFontSize(fontName string, fontSize float64) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setKeyEquivalentFont:size:"), objc.String(fontName), fontSize)
}

// The gradient of the button’s border.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/gradientType
func (b_ ButtonCell) GradientType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("gradientType"))
	return rv
}


// SetGradientType sets the value of the gradientType property.
// The gradient of the button’s border.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/gradientType
func (b_ ButtonCell) SetGradientType(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setGradientType:"), value)
}

// The button’s key-equivalent character.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/keyEquivalent
func (b_ ButtonCell) KeyEquivalent() string {
	rv := objc.Send[string](b_.ID, objc.Sel("keyEquivalent"))
	return rv
}


// SetKeyEquivalent sets the value of the keyEquivalent property.
// The button’s key-equivalent character.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/keyEquivalent
func (b_ ButtonCell) SetKeyEquivalent(value string) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setKeyEquivalent:"), objc.String(value))
}

// The font used to draw the button’s key equivalent.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/keyEquivalentFont
func (b_ ButtonCell) KeyEquivalentFont() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("keyEquivalentFont"))
	return rv
}


// SetKeyEquivalentFont sets the value of the keyEquivalentFont property.
// The font used to draw the button’s key equivalent.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/keyEquivalentFont
func (b_ ButtonCell) SetKeyEquivalentFont(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setKeyEquivalentFont:"), value)
}

// The mask that identifies the modifier keys for the button’s key equivalent.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/keyEquivalentModifierMask
func (b_ ButtonCell) KeyEquivalentModifierMask() EventModifierFlags {
	rv := objc.Send[EventModifierFlags](b_.ID, objc.Sel("keyEquivalentModifierMask"))
	return rv
}


// SetKeyEquivalentModifierMask sets the value of the keyEquivalentModifierMask property.
// The mask that identifies the modifier keys for the button’s key equivalent.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/keyEquivalentModifierMask
func (b_ ButtonCell) SetKeyEquivalentModifierMask(value EventModifierFlags) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setKeyEquivalentModifierMask:"), value)
}

// The image the button displays in its alternate state.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/alternateimage
func (b_ ButtonCell) AlternateImage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("alternateImage"))
	return rv
}


// SetAlternateImage sets the value of the alternateImage property.
// The image the button displays in its alternate state.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/alternateimage
func (b_ ButtonCell) SetAlternateImage(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAlternateImage:"), value)
}

// The string displayed by the button when it’s in its alternate state.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/alternatetitle
func (b_ ButtonCell) AlternateTitle() string {
	rv := objc.Send[string](b_.ID, objc.Sel("alternateTitle"))
	return rv
}


// SetAlternateTitle sets the value of the alternateTitle property.
// The string displayed by the button when it’s in its alternate state.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/alternatetitle
func (b_ ButtonCell) SetAlternateTitle(value string) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAlternateTitle:"), objc.String(value))
}

// The title displayed by the button when it’s in its alternate state, as an attributed string.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/attributedalternatetitle
func (b_ ButtonCell) AttributedAlternateTitle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("attributedAlternateTitle"))
	return rv
}


// SetAttributedAlternateTitle sets the value of the attributedAlternateTitle property.
// The title displayed by the button when it’s in its alternate state, as an attributed string.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/attributedalternatetitle
func (b_ ButtonCell) SetAttributedAlternateTitle(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAttributedAlternateTitle:"), value)
}

// The title displayed by the button when it’s in its normal state as an attributed string.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/attributedtitle
func (b_ ButtonCell) AttributedTitle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("attributedTitle"))
	return rv
}


// SetAttributedTitle sets the value of the attributedTitle property.
// The title displayed by the button when it’s in its normal state as an attributed string.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/attributedtitle
func (b_ ButtonCell) SetAttributedTitle(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAttributedTitle:"), value)
}

// The background color of the button.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/backgroundcolor
func (b_ ButtonCell) BackgroundColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("backgroundColor"))
	return rv
}


// SetBackgroundColor sets the value of the backgroundColor property.
// The background color of the button.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/backgroundcolor
func (b_ ButtonCell) SetBackgroundColor(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBackgroundColor:"), value)
}

// The appearance of the button’s border, if it has one.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/bezelstyle
func (b_ ButtonCell) BezelStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("bezelStyle"))
	return rv
}


// SetBezelStyle sets the value of the bezelStyle property.
// The appearance of the button’s border, if it has one.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/bezelstyle
func (b_ ButtonCell) SetBezelStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBezelStyle:"), value)
}

// A set of flags that indicate how the button highlights when it receives a mouse-down event (that is, when the button is pressed).
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/highlightsby
func (b_ ButtonCell) HighlightsBy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("highlightsBy"))
	return rv
}


// SetHighlightsBy sets the value of the highlightsBy property.
// A set of flags that indicate how the button highlights when it receives a mouse-down event (that is, when the button is pressed).

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/highlightsby
func (b_ ButtonCell) SetHighlightsBy(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setHighlightsBy:"), value)
}

// A Boolean value that indicates if the button’s image and text appear “dim” when the button is disabled.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/imagedimswhendisabled
func (b_ ButtonCell) ImageDimsWhenDisabled() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("imageDimsWhenDisabled"))
	return rv
}


// SetImageDimsWhenDisabled sets the value of the imageDimsWhenDisabled property.
// A Boolean value that indicates if the button’s image and text appear “dim” when the button is disabled.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/imagedimswhendisabled
func (b_ ButtonCell) SetImageDimsWhenDisabled(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setImageDimsWhenDisabled:"), value)
}

// The position of the button’s image relative to its title.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/imageposition
func (b_ ButtonCell) ImagePosition() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("imagePosition"))
	return rv
}


// SetImagePosition sets the value of the imagePosition property.
// The position of the button’s image relative to its title.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/imageposition
func (b_ ButtonCell) SetImagePosition(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setImagePosition:"), value)
}

// The scale factor for the button’s image.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/imagescaling
func (b_ ButtonCell) ImageScaling() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("imageScaling"))
	return rv
}


// SetImageScaling sets the value of the imageScaling property.
// The scale factor for the button’s image.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/imagescaling
func (b_ ButtonCell) SetImageScaling(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setImageScaling:"), value)
}

// A Boolean value that indicates if the button is opaque.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/isopaque
func (b_ ButtonCell) IsOpaque() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isOpaque"))
	return rv
}


// SetIsOpaque sets the value of the isOpaque property.
// A Boolean value that indicates if the button is opaque.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/isopaque
func (b_ ButtonCell) SetIsOpaque(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsOpaque:"), value)
}

// A Boolean value that indicates if the button is transparent.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/istransparent
func (b_ ButtonCell) IsTransparent() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isTransparent"))
	return rv
}


// SetIsTransparent sets the value of the isTransparent property.
// A Boolean value that indicates if the button is transparent.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/istransparent
func (b_ ButtonCell) SetIsTransparent(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsTransparent:"), value)
}

// A Boolean value that indicates if the button displays its border only when the pointer is over it.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/showsborderonlywhilemouseinside
func (b_ ButtonCell) ShowsBorderOnlyWhileMouseInside() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("showsBorderOnlyWhileMouseInside"))
	return rv
}


// SetShowsBorderOnlyWhileMouseInside sets the value of the showsBorderOnlyWhileMouseInside property.
// A Boolean value that indicates if the button displays its border only when the pointer is over it.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/showsborderonlywhilemouseinside
func (b_ ButtonCell) SetShowsBorderOnlyWhileMouseInside(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setShowsBorderOnlyWhileMouseInside:"), value)
}

// The flags that indicate how the button cell shows its alternate state.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/showsstateby
func (b_ ButtonCell) ShowsStateBy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("showsStateBy"))
	return rv
}


// SetShowsStateBy sets the value of the showsStateBy property.
// The flags that indicate how the button cell shows its alternate state.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/showsstateby
func (b_ ButtonCell) SetShowsStateBy(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setShowsStateBy:"), value)
}

// The sound that’s played when the user presses the button (that is during a mouse-down event).
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/sound
func (b_ ButtonCell) Sound() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("sound"))
	return rv
}


// SetSound sets the value of the sound property.
// The sound that’s played when the user presses the button (that is during a mouse-down event).

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/sound
func (b_ ButtonCell) SetSound(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSound:"), value)
}

// The title displayed on the button when it’s in its normal state.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/title
func (b_ ButtonCell) Title() string {
	rv := objc.Send[string](b_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
// The title displayed on the button when it’s in its normal state.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/title
func (b_ ButtonCell) SetTitle(value string) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTitle:"), objc.String(value))
}

// The font that the cell uses to display text.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/font
func (b_ ButtonCell) Font() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("font"))
	return rv
}


// SetFont sets the value of the font property.
// The font that the cell uses to display text.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/font
func (b_ ButtonCell) SetFont(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setFont:"), value)
}

// The cell’s value as an integer.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/intvalue
func (b_ ButtonCell) IntValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("intValue"))
	return rv
}


// SetIntValue sets the value of the intValue property.
// The cell’s value as an integer.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/intvalue
func (b_ ButtonCell) SetIntValue(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIntValue:"), value)
}

// The cell’s value as an Objective-C object.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/objectvalue
func (b_ ButtonCell) ObjectValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("objectValue"))
	return rv
}


// SetObjectValue sets the value of the objectValue property.
// The cell’s value as an Objective-C object.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/objectvalue
func (b_ ButtonCell) SetObjectValue(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setObjectValue:"), value)
}

// The cell’s current state.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/state
func (b_ ButtonCell) State() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("state"))
	return rv
}


// SetState sets the value of the state property.
// The cell’s current state.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/state
func (b_ ButtonCell) SetState(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setState:"), value)
}

// The value of the receiver’s cell as a double-precision floating-point number.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/doublevalue
func (b_ ButtonCell) DoubleValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("doubleValue"))
	return rv
}


// SetDoubleValue sets the value of the doubleValue property.
// The value of the receiver’s cell as a double-precision floating-point number.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/doublevalue
func (b_ ButtonCell) SetDoubleValue(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setDoubleValue:"), value)
}



