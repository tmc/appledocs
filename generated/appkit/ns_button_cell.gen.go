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
	AlternateImage() IImage
	SetAlternateImage(value IImage)
	AlternateTitle() objc.IObject /* cross-framework: NSString */
	SetAlternateTitle(value objc.IObject /* cross-framework: NSString */)
	AttributedAlternateTitle() foundation.AttributedString
	SetAttributedAlternateTitle(value foundation.AttributedString)
	AttributedTitle() foundation.AttributedString
	SetAttributedTitle(value foundation.AttributedString)
	BackgroundColor() IColor
	SetBackgroundColor(value IColor)
	BezelStyle() BezelStyle
	SetBezelStyle(value BezelStyle)
	GradientType() GradientType
	SetGradientType(value GradientType)
	HighlightsBy() CellStyleMask
	SetHighlightsBy(value CellStyleMask)
	ImageDimsWhenDisabled() bool
	SetImageDimsWhenDisabled(value bool)
	ImagePosition() CellImagePosition
	SetImagePosition(value CellImagePosition)
	ImageScaling() ImageScaling
	SetImageScaling(value ImageScaling)
	Opaque() bool
	Transparent() bool
	SetTransparent(value bool)
	KeyEquivalent() objc.IObject /* cross-framework: NSString */
	SetKeyEquivalent(value objc.IObject /* cross-framework: NSString */)
	KeyEquivalentFont() IFont
	SetKeyEquivalentFont(value IFont)
	KeyEquivalentModifierMask() EventModifierFlags
	SetKeyEquivalentModifierMask(value EventModifierFlags)
	ShowsBorderOnlyWhileMouseInside() bool
	SetShowsBorderOnlyWhileMouseInside(value bool)
	ShowsStateBy() CellStyleMask
	SetShowsStateBy(value CellStyleMask)
	Sound() ISound
	SetSound(value ISound)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	IsOpaque() bool
	SetIsOpaque(value bool)
	IsTransparent() bool
	SetIsTransparent(value bool)
	Font() IFont
	SetFont(value IFont)
	IntValue() unsafe.Pointer
	SetIntValue(value unsafe.Pointer)
	State() unsafe.Pointer
	SetState(value unsafe.Pointer)
	DoubleValue() float64
	SetDoubleValue(value float64)
	// methods:
	DrawBezelWithFrameInView(frame objc.IObject /* cross-framework: Rect */, controlView IView)
	DrawImageWithFrameInView(image IImage, frame objc.IObject /* cross-framework: Rect */, controlView IView)
	DrawTitleWithFrameInView(title foundation.AttributedString, frame objc.IObject /* cross-framework: Rect */, controlView IView) objc.IObject /* cross-framework: Rect */
	GetPeriodicDelayInterval(delay unsafe.Pointer, interval unsafe.Pointer)
	MouseEntered(event IEvent)
	MouseExited(event IEvent)
	PerformClick(sender objc.IObject)
	SetButtonType(type_ ButtonType)
	SetPeriodicDelayInterval(delay float32, interval float32)
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/init(imageCell:)
func NewButtonCellImageCell(image IImage) ButtonCell {
	instance := getButtonCellClass().Alloc()
	rv := objc.Send[ButtonCell](instance.ID, objc.Sel("initImageCell:"), image)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/init(textCell:)
func NewButtonCellTextCell(string_ objc.IObject /* cross-framework: NSString */) ButtonCell {
	instance := getButtonCellClass().Alloc()
	rv := objc.Send[ButtonCell](instance.ID, objc.Sel("initTextCell:"), string_)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/init(coder:)
func NewButtonCellWithCoder(coder foundation.Coder) ButtonCell {
	instance := getButtonCellClass().Alloc()
	rv := objc.Send[ButtonCell](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}



// Draws the border of the button using the current bezel style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/drawBezel(withFrame:in:)
func (b_ ButtonCell) DrawBezelWithFrameInView(frame objc.IObject /* cross-framework: Rect */, controlView IView) {
	objc.Send[objc.ID](b_.ID, objc.Sel("drawBezelWithFrame:inView:"), frame, controlView)
}


// Draws the image associated with the button’s current state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/drawImage(_:withFrame:in:)
func (b_ ButtonCell) DrawImageWithFrameInView(image IImage, frame objc.IObject /* cross-framework: Rect */, controlView IView) {
	objc.Send[objc.ID](b_.ID, objc.Sel("drawImage:withFrame:inView:"), image, frame, controlView)
}


// Draws the button’s title centered vertically in a specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/drawTitle(_:withFrame:in:)
func (b_ ButtonCell) DrawTitleWithFrameInView(title foundation.AttributedString, frame objc.IObject /* cross-framework: Rect */, controlView IView) objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](b_.ID, objc.Sel("drawTitle:withFrame:inView:"), title, frame, controlView)
	return rv
}


// Returns by reference the delay and interval periods for a continuous button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/getPeriodicDelay(_:interval:)
func (b_ ButtonCell) GetPeriodicDelayInterval(delay unsafe.Pointer, interval unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("getPeriodicDelay:interval:"), delay, interval)
}


// Draws the button’s border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/mouseEntered(with:)
func (b_ ButtonCell) MouseEntered(event IEvent) {
	objc.Send[objc.ID](b_.ID, objc.Sel("mouseEntered:"), event)
}


// Erases the button’s border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/mouseExited(with:)
func (b_ ButtonCell) MouseExited(event IEvent) {
	objc.Send[objc.ID](b_.ID, objc.Sel("mouseExited:"), event)
}


// Simulates the user clicking the button with the pointer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/performClick(_:)
func (b_ ButtonCell) PerformClick(sender objc.IObject) {
	objc.Send[objc.ID](b_.ID, objc.Sel("performClick:"), sender)
}


// Sets how the button highlights while pressed and how it shows its state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/setButtonType(_:)
func (b_ ButtonCell) SetButtonType(type_ ButtonType) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setButtonType:"), type_)
}


// Sets the message delay and interval for the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/setPeriodicDelay(_:interval:)
func (b_ ButtonCell) SetPeriodicDelayInterval(delay float32, interval float32) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPeriodicDelay:interval:"), delay, interval)
}


// The image the button displays in its alternate state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/alternateImage
func (b_ ButtonCell) AlternateImage() IImage {
	rv := objc.Send[Image](b_.ID, objc.Sel("alternateImage"))
	return rv
}


// The image the button displays in its alternate state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/alternateImage
func (b_ ButtonCell) SetAlternateImage(value IImage) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAlternateImage:"), value)
}


// The string displayed by the button when it’s in its alternate state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/alternateTitle
func (b_ ButtonCell) AlternateTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](b_.ID, objc.Sel("alternateTitle"))
	return rv
}


// The string displayed by the button when it’s in its alternate state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/alternateTitle
func (b_ ButtonCell) SetAlternateTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAlternateTitle:"), value)
}


// The title displayed by the button when it’s in its alternate state, as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/attributedAlternateTitle
func (b_ ButtonCell) AttributedAlternateTitle() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](b_.ID, objc.Sel("attributedAlternateTitle"))
	return rv
}


// The title displayed by the button when it’s in its alternate state, as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/attributedAlternateTitle
func (b_ ButtonCell) SetAttributedAlternateTitle(value foundation.AttributedString) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAttributedAlternateTitle:"), value)
}


// The title displayed by the button when it’s in its normal state as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/attributedTitle
func (b_ ButtonCell) AttributedTitle() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](b_.ID, objc.Sel("attributedTitle"))
	return rv
}


// The title displayed by the button when it’s in its normal state as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/attributedTitle
func (b_ ButtonCell) SetAttributedTitle(value foundation.AttributedString) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAttributedTitle:"), value)
}


// The background color of the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/backgroundColor
func (b_ ButtonCell) BackgroundColor() IColor {
	rv := objc.Send[Color](b_.ID, objc.Sel("backgroundColor"))
	return rv
}


// The background color of the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/backgroundColor
func (b_ ButtonCell) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBackgroundColor:"), value)
}


// The appearance of the button’s border, if it has one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/bezelStyle
func (b_ ButtonCell) BezelStyle() BezelStyle {
	rv := objc.Send[BezelStyle](b_.ID, objc.Sel("bezelStyle"))
	return rv
}


// The appearance of the button’s border, if it has one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/bezelStyle
func (b_ ButtonCell) SetBezelStyle(value BezelStyle) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBezelStyle:"), value)
}


// The gradient of the button’s border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/gradientType
func (b_ ButtonCell) GradientType() GradientType {
	rv := objc.Send[GradientType](b_.ID, objc.Sel("gradientType"))
	return rv
}


// The gradient of the button’s border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/gradientType
func (b_ ButtonCell) SetGradientType(value GradientType) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setGradientType:"), value)
}


// A set of flags that indicate how the button highlights when it receives a mouse-down event (that is, when the button is pressed).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/highlightsBy
func (b_ ButtonCell) HighlightsBy() CellStyleMask {
	rv := objc.Send[CellStyleMask](b_.ID, objc.Sel("highlightsBy"))
	return rv
}


// A set of flags that indicate how the button highlights when it receives a mouse-down event (that is, when the button is pressed).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/highlightsBy
func (b_ ButtonCell) SetHighlightsBy(value CellStyleMask) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setHighlightsBy:"), value)
}


// A Boolean value that indicates if the button’s image and text appear “dim” when the button is disabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/imageDimsWhenDisabled
func (b_ ButtonCell) ImageDimsWhenDisabled() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("imageDimsWhenDisabled"))
	return rv
}


// A Boolean value that indicates if the button’s image and text appear “dim” when the button is disabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/imageDimsWhenDisabled
func (b_ ButtonCell) SetImageDimsWhenDisabled(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setImageDimsWhenDisabled:"), value)
}


// The position of the button’s image relative to its title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/imagePosition
func (b_ ButtonCell) ImagePosition() CellImagePosition {
	rv := objc.Send[CellImagePosition](b_.ID, objc.Sel("imagePosition"))
	return rv
}


// The position of the button’s image relative to its title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/imagePosition
func (b_ ButtonCell) SetImagePosition(value CellImagePosition) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setImagePosition:"), value)
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


// A Boolean value that indicates if the button is opaque.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/isOpaque
func (b_ ButtonCell) Opaque() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("opaque"))
	return rv
}


// A Boolean value that indicates if the button is transparent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/isTransparent
func (b_ ButtonCell) Transparent() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("transparent"))
	return rv
}


// A Boolean value that indicates if the button is transparent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/isTransparent
func (b_ ButtonCell) SetTransparent(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTransparent:"), value)
}


// The button’s key-equivalent character.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/keyEquivalent
func (b_ ButtonCell) KeyEquivalent() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](b_.ID, objc.Sel("keyEquivalent"))
	return rv
}


// The button’s key-equivalent character.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/keyEquivalent
func (b_ ButtonCell) SetKeyEquivalent(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setKeyEquivalent:"), value)
}


// The font used to draw the button’s key equivalent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/keyEquivalentFont
func (b_ ButtonCell) KeyEquivalentFont() IFont {
	rv := objc.Send[Font](b_.ID, objc.Sel("keyEquivalentFont"))
	return rv
}


// The font used to draw the button’s key equivalent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/keyEquivalentFont
func (b_ ButtonCell) SetKeyEquivalentFont(value IFont) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setKeyEquivalentFont:"), value)
}


// The mask that identifies the modifier keys for the button’s key equivalent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/keyEquivalentModifierMask
func (b_ ButtonCell) KeyEquivalentModifierMask() EventModifierFlags {
	rv := objc.Send[EventModifierFlags](b_.ID, objc.Sel("keyEquivalentModifierMask"))
	return rv
}


// The mask that identifies the modifier keys for the button’s key equivalent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/keyEquivalentModifierMask
func (b_ ButtonCell) SetKeyEquivalentModifierMask(value EventModifierFlags) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setKeyEquivalentModifierMask:"), value)
}


// A Boolean value that indicates if the button displays its border only when the pointer is over it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/showsBorderOnlyWhileMouseInside
func (b_ ButtonCell) ShowsBorderOnlyWhileMouseInside() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("showsBorderOnlyWhileMouseInside"))
	return rv
}


// A Boolean value that indicates if the button displays its border only when the pointer is over it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/showsBorderOnlyWhileMouseInside
func (b_ ButtonCell) SetShowsBorderOnlyWhileMouseInside(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setShowsBorderOnlyWhileMouseInside:"), value)
}


// The flags that indicate how the button cell shows its alternate state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/showsStateBy
func (b_ ButtonCell) ShowsStateBy() CellStyleMask {
	rv := objc.Send[CellStyleMask](b_.ID, objc.Sel("showsStateBy"))
	return rv
}


// The flags that indicate how the button cell shows its alternate state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/showsStateBy
func (b_ ButtonCell) SetShowsStateBy(value CellStyleMask) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setShowsStateBy:"), value)
}


// The sound that’s played when the user presses the button (that is during a mouse-down event).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/sound
func (b_ ButtonCell) Sound() ISound {
	rv := objc.Send[Sound](b_.ID, objc.Sel("sound"))
	return rv
}


// The sound that’s played when the user presses the button (that is during a mouse-down event).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/sound
func (b_ ButtonCell) SetSound(value ISound) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSound:"), value)
}


// The title displayed on the button when it’s in its normal state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/title
func (b_ ButtonCell) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](b_.ID, objc.Sel("title"))
	return rv
}


// The title displayed on the button when it’s in its normal state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/title
func (b_ ButtonCell) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTitle:"), value)
}


// A Boolean value that indicates if the button is opaque.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/isopaque
func (b_ ButtonCell) IsOpaque() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isOpaque"))
	return rv
}


// A Boolean value that indicates if the button is opaque.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/isopaque
func (b_ ButtonCell) SetIsOpaque(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsOpaque:"), value)
}


// A Boolean value that indicates if the button is transparent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/istransparent
func (b_ ButtonCell) IsTransparent() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isTransparent"))
	return rv
}


// A Boolean value that indicates if the button is transparent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/istransparent
func (b_ ButtonCell) SetIsTransparent(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsTransparent:"), value)
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
func (b_ ButtonCell) DoubleValue() float64 {
	rv := objc.Send[float64](b_.ID, objc.Sel("doubleValue"))
	return rv
}


// The value of the receiver’s cell as a double-precision floating-point number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/doublevalue
func (b_ ButtonCell) SetDoubleValue(value float64) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setDoubleValue:"), value)
}


