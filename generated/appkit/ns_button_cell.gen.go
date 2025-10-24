// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSButtonCell */


/* debug [class_header]: Header for NSButtonCell */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ButtonCell */
// An interface definition for the [ButtonCell] class.
type IButtonCell interface {
	IActionCell
	
/* debug [class_interface_properties]: Properties for ButtonCell */
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
	IntValue() objectivec.IObject
	SetIntValue(value objectivec.IObject)
	State() objectivec.IObject
	SetState(value objectivec.IObject)
	DoubleValue() float64
	SetDoubleValue(value float64)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ButtonCell */
	// methods:
	DrawBezelWithFrameInView(frame Rect /* not a class type */, controlView IView)
	DrawImageWithFrameInView(image IImage, frame Rect /* not a class type */, controlView IView)
	DrawTitleWithFrameInView(title foundation.AttributedString, frame Rect /* not a class type */, controlView IView) Rect /* not a class type */
	GetPeriodicDelayInterval(delay objectivec.IObject, interval objectivec.IObject)
	MouseEntered(event IEvent)
	MouseExited(event IEvent)
	PerformClick(sender objc.IObject)
	SetButtonType(type_ ButtonType)
	SetPeriodicDelayInterval(delay float32, interval float32)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ButtonCell */
// Alloc allocates a new instance without initialization.
func (bc _ButtonCellClass) Alloc() ButtonCell {
	rv := objc.Send[ButtonCell](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ButtonCell */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ButtonCell */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/init(imageCell:)
func NewButtonCellImageCell(image IImage) ButtonCell {
	instance := getButtonCellClass().Alloc()
	rv := objc.Send[ButtonCell](instance.ID, objc.Sel("initImageCell:"), image)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewButtonCellImageCell */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/init(textCell:)
func NewButtonCellTextCell(string_ objc.IObject /* cross-framework: NSString */) ButtonCell {
	instance := getButtonCellClass().Alloc()
	rv := objc.Send[ButtonCell](instance.ID, objc.Sel("initTextCell:"), string_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewButtonCellTextCell */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/init(coder:)
func NewButtonCellWithCoder(coder foundation.Coder) ButtonCell {
	instance := getButtonCellClass().Alloc()
	rv := objc.Send[ButtonCell](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewButtonCellWithCoder */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ButtonCell */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ButtonCell */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ButtonCell */

// Draws the border of the button using the current bezel style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/drawBezel(withFrame:in:)
func (b_ ButtonCell) DrawBezelWithFrameInView(frame Rect /* not a class type */, controlView IView) {
	objc.Send[objc.ID](b_.ID, objc.Sel("drawBezelWithFrame:inView:"), frame, controlView)
}/* debug [instance_methods/method]: DrawBezelWithFrameInView */


// Draws the image associated with the button’s current state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/drawImage(_:withFrame:in:)
func (b_ ButtonCell) DrawImageWithFrameInView(image IImage, frame Rect /* not a class type */, controlView IView) {
	objc.Send[objc.ID](b_.ID, objc.Sel("drawImage:withFrame:inView:"), image, frame, controlView)
}/* debug [instance_methods/method]: DrawImageWithFrameInView */


// Draws the button’s title centered vertically in a specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/drawTitle(_:withFrame:in:)
func (b_ ButtonCell) DrawTitleWithFrameInView(title foundation.AttributedString, frame Rect /* not a class type */, controlView IView) Rect /* not a class type */ {
	rv := objc.Send[Rect](b_.ID, objc.Sel("drawTitle:withFrame:inView:"), title, frame, controlView)
	return rv
}/* debug [instance_methods/method]: DrawTitleWithFrameInView */


// Returns by reference the delay and interval periods for a continuous button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/getPeriodicDelay(_:interval:)
func (b_ ButtonCell) GetPeriodicDelayInterval(delay objectivec.IObject, interval objectivec.IObject) {
	objc.Send[objc.ID](b_.ID, objc.Sel("getPeriodicDelay:interval:"), delay, interval)
}/* debug [instance_methods/method]: GetPeriodicDelayInterval */


// Draws the button’s border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/mouseEntered(with:)
func (b_ ButtonCell) MouseEntered(event IEvent) {
	objc.Send[objc.ID](b_.ID, objc.Sel("mouseEntered:"), event)
}/* debug [instance_methods/method]: MouseEntered */


// Erases the button’s border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/mouseExited(with:)
func (b_ ButtonCell) MouseExited(event IEvent) {
	objc.Send[objc.ID](b_.ID, objc.Sel("mouseExited:"), event)
}/* debug [instance_methods/method]: MouseExited */


// Simulates the user clicking the button with the pointer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/performClick(_:)
func (b_ ButtonCell) PerformClick(sender objc.IObject) {
	objc.Send[objc.ID](b_.ID, objc.Sel("performClick:"), sender)
}/* debug [instance_methods/method]: PerformClick */


// Sets how the button highlights while pressed and how it shows its state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/setButtonType(_:)
func (b_ ButtonCell) SetButtonType(type_ ButtonType) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setButtonType:"), type_)
}/* debug [instance_methods/method]: SetButtonType */


// Sets the message delay and interval for the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/setPeriodicDelay(_:interval:)
func (b_ ButtonCell) SetPeriodicDelayInterval(delay float32, interval float32) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPeriodicDelay:interval:"), delay, interval)
}/* debug [instance_methods/method]: SetPeriodicDelayInterval */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ButtonCell */

// The image the button displays in its alternate state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/alternateImage
func (b_ ButtonCell) AlternateImage() IImage {
	rv := objc.Send[Image](b_.ID, objc.Sel("alternateImage"))
	return rv
}/* debug [instance_properties/getter]: alternateImage */


// The image the button displays in its alternate state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/alternateImage
func (b_ ButtonCell) SetAlternateImage(value IImage) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAlternateImage:"), value)
}/* debug [instance_properties/setter]: alternateImage */


// The string displayed by the button when it’s in its alternate state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/alternateTitle
func (b_ ButtonCell) AlternateTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](b_.ID, objc.Sel("alternateTitle"))
	return rv
}/* debug [instance_properties/getter]: alternateTitle */


// The string displayed by the button when it’s in its alternate state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/alternateTitle
func (b_ ButtonCell) SetAlternateTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAlternateTitle:"), value)
}/* debug [instance_properties/setter]: alternateTitle */


// The title displayed by the button when it’s in its alternate state, as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/attributedAlternateTitle
func (b_ ButtonCell) AttributedAlternateTitle() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](b_.ID, objc.Sel("attributedAlternateTitle"))
	return rv
}/* debug [instance_properties/getter]: attributedAlternateTitle */


// The title displayed by the button when it’s in its alternate state, as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/attributedAlternateTitle
func (b_ ButtonCell) SetAttributedAlternateTitle(value foundation.AttributedString) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAttributedAlternateTitle:"), value)
}/* debug [instance_properties/setter]: attributedAlternateTitle */


// The title displayed by the button when it’s in its normal state as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/attributedTitle
func (b_ ButtonCell) AttributedTitle() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](b_.ID, objc.Sel("attributedTitle"))
	return rv
}/* debug [instance_properties/getter]: attributedTitle */


// The title displayed by the button when it’s in its normal state as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/attributedTitle
func (b_ ButtonCell) SetAttributedTitle(value foundation.AttributedString) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAttributedTitle:"), value)
}/* debug [instance_properties/setter]: attributedTitle */


// The background color of the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/backgroundColor
func (b_ ButtonCell) BackgroundColor() IColor {
	rv := objc.Send[Color](b_.ID, objc.Sel("backgroundColor"))
	return rv
}/* debug [instance_properties/getter]: backgroundColor */


// The background color of the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/backgroundColor
func (b_ ButtonCell) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBackgroundColor:"), value)
}/* debug [instance_properties/setter]: backgroundColor */


// The appearance of the button’s border, if it has one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/bezelStyle
func (b_ ButtonCell) BezelStyle() BezelStyle {
	rv := objc.Send[BezelStyle](b_.ID, objc.Sel("bezelStyle"))
	return rv
}/* debug [instance_properties/getter]: bezelStyle */


// The appearance of the button’s border, if it has one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/bezelStyle
func (b_ ButtonCell) SetBezelStyle(value BezelStyle) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBezelStyle:"), value)
}/* debug [instance_properties/setter]: bezelStyle */


// The gradient of the button’s border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/gradientType
func (b_ ButtonCell) GradientType() GradientType {
	rv := objc.Send[GradientType](b_.ID, objc.Sel("gradientType"))
	return rv
}/* debug [instance_properties/getter]: gradientType */


// The gradient of the button’s border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/gradientType
func (b_ ButtonCell) SetGradientType(value GradientType) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setGradientType:"), value)
}/* debug [instance_properties/setter]: gradientType */


// A set of flags that indicate how the button highlights when it receives a mouse-down event (that is, when the button is pressed).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/highlightsBy
func (b_ ButtonCell) HighlightsBy() CellStyleMask {
	rv := objc.Send[CellStyleMask](b_.ID, objc.Sel("highlightsBy"))
	return rv
}/* debug [instance_properties/getter]: highlightsBy */


// A set of flags that indicate how the button highlights when it receives a mouse-down event (that is, when the button is pressed).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/highlightsBy
func (b_ ButtonCell) SetHighlightsBy(value CellStyleMask) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setHighlightsBy:"), value)
}/* debug [instance_properties/setter]: highlightsBy */


// A Boolean value that indicates if the button’s image and text appear “dim” when the button is disabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/imageDimsWhenDisabled
func (b_ ButtonCell) ImageDimsWhenDisabled() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("imageDimsWhenDisabled"))
	return rv
}/* debug [instance_properties/getter]: imageDimsWhenDisabled */


// A Boolean value that indicates if the button’s image and text appear “dim” when the button is disabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/imageDimsWhenDisabled
func (b_ ButtonCell) SetImageDimsWhenDisabled(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setImageDimsWhenDisabled:"), value)
}/* debug [instance_properties/setter]: imageDimsWhenDisabled */


// The position of the button’s image relative to its title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/imagePosition
func (b_ ButtonCell) ImagePosition() CellImagePosition {
	rv := objc.Send[CellImagePosition](b_.ID, objc.Sel("imagePosition"))
	return rv
}/* debug [instance_properties/getter]: imagePosition */


// The position of the button’s image relative to its title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/imagePosition
func (b_ ButtonCell) SetImagePosition(value CellImagePosition) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setImagePosition:"), value)
}/* debug [instance_properties/setter]: imagePosition */


// The scale factor for the button’s image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/imageScaling
func (b_ ButtonCell) ImageScaling() ImageScaling {
	rv := objc.Send[ImageScaling](b_.ID, objc.Sel("imageScaling"))
	return rv
}/* debug [instance_properties/getter]: imageScaling */


// The scale factor for the button’s image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/imageScaling
func (b_ ButtonCell) SetImageScaling(value ImageScaling) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setImageScaling:"), value)
}/* debug [instance_properties/setter]: imageScaling */


// A Boolean value that indicates if the button is opaque.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/isOpaque
func (b_ ButtonCell) Opaque() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("opaque"))
	return rv
}/* debug [instance_properties/getter]: opaque */


// A Boolean value that indicates if the button is transparent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/isTransparent
func (b_ ButtonCell) Transparent() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("transparent"))
	return rv
}/* debug [instance_properties/getter]: transparent */


// A Boolean value that indicates if the button is transparent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/isTransparent
func (b_ ButtonCell) SetTransparent(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTransparent:"), value)
}/* debug [instance_properties/setter]: transparent */


// The button’s key-equivalent character.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/keyEquivalent
func (b_ ButtonCell) KeyEquivalent() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](b_.ID, objc.Sel("keyEquivalent"))
	return rv
}/* debug [instance_properties/getter]: keyEquivalent */


// The button’s key-equivalent character.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/keyEquivalent
func (b_ ButtonCell) SetKeyEquivalent(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setKeyEquivalent:"), value)
}/* debug [instance_properties/setter]: keyEquivalent */


// The font used to draw the button’s key equivalent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/keyEquivalentFont
func (b_ ButtonCell) KeyEquivalentFont() IFont {
	rv := objc.Send[Font](b_.ID, objc.Sel("keyEquivalentFont"))
	return rv
}/* debug [instance_properties/getter]: keyEquivalentFont */


// The font used to draw the button’s key equivalent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/keyEquivalentFont
func (b_ ButtonCell) SetKeyEquivalentFont(value IFont) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setKeyEquivalentFont:"), value)
}/* debug [instance_properties/setter]: keyEquivalentFont */


// The mask that identifies the modifier keys for the button’s key equivalent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/keyEquivalentModifierMask
func (b_ ButtonCell) KeyEquivalentModifierMask() EventModifierFlags {
	rv := objc.Send[EventModifierFlags](b_.ID, objc.Sel("keyEquivalentModifierMask"))
	return rv
}/* debug [instance_properties/getter]: keyEquivalentModifierMask */


// The mask that identifies the modifier keys for the button’s key equivalent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/keyEquivalentModifierMask
func (b_ ButtonCell) SetKeyEquivalentModifierMask(value EventModifierFlags) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setKeyEquivalentModifierMask:"), value)
}/* debug [instance_properties/setter]: keyEquivalentModifierMask */


// A Boolean value that indicates if the button displays its border only when the pointer is over it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/showsBorderOnlyWhileMouseInside
func (b_ ButtonCell) ShowsBorderOnlyWhileMouseInside() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("showsBorderOnlyWhileMouseInside"))
	return rv
}/* debug [instance_properties/getter]: showsBorderOnlyWhileMouseInside */


// A Boolean value that indicates if the button displays its border only when the pointer is over it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/showsBorderOnlyWhileMouseInside
func (b_ ButtonCell) SetShowsBorderOnlyWhileMouseInside(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setShowsBorderOnlyWhileMouseInside:"), value)
}/* debug [instance_properties/setter]: showsBorderOnlyWhileMouseInside */


// The flags that indicate how the button cell shows its alternate state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/showsStateBy
func (b_ ButtonCell) ShowsStateBy() CellStyleMask {
	rv := objc.Send[CellStyleMask](b_.ID, objc.Sel("showsStateBy"))
	return rv
}/* debug [instance_properties/getter]: showsStateBy */


// The flags that indicate how the button cell shows its alternate state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/showsStateBy
func (b_ ButtonCell) SetShowsStateBy(value CellStyleMask) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setShowsStateBy:"), value)
}/* debug [instance_properties/setter]: showsStateBy */


// The sound that’s played when the user presses the button (that is during a mouse-down event).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/sound
func (b_ ButtonCell) Sound() ISound {
	rv := objc.Send[Sound](b_.ID, objc.Sel("sound"))
	return rv
}/* debug [instance_properties/getter]: sound */


// The sound that’s played when the user presses the button (that is during a mouse-down event).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/sound
func (b_ ButtonCell) SetSound(value ISound) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSound:"), value)
}/* debug [instance_properties/setter]: sound */


// The title displayed on the button when it’s in its normal state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/title
func (b_ ButtonCell) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](b_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// The title displayed on the button when it’s in its normal state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/title
func (b_ ButtonCell) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTitle:"), value)
}/* debug [instance_properties/setter]: title */


// A Boolean value that indicates if the button is opaque.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/isopaque
func (b_ ButtonCell) IsOpaque() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isOpaque"))
	return rv
}/* debug [instance_properties/getter]: isOpaque */


// A Boolean value that indicates if the button is opaque.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/isopaque
func (b_ ButtonCell) SetIsOpaque(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsOpaque:"), value)
}/* debug [instance_properties/setter]: isOpaque */


// A Boolean value that indicates if the button is transparent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/istransparent
func (b_ ButtonCell) IsTransparent() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isTransparent"))
	return rv
}/* debug [instance_properties/getter]: isTransparent */


// A Boolean value that indicates if the button is transparent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttoncell/istransparent
func (b_ ButtonCell) SetIsTransparent(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsTransparent:"), value)
}/* debug [instance_properties/setter]: isTransparent */


// The font that the cell uses to display text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/font
func (b_ ButtonCell) Font() IFont {
	rv := objc.Send[Font](b_.ID, objc.Sel("font"))
	return rv
}/* debug [instance_properties/getter]: font */


// The font that the cell uses to display text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/font
func (b_ ButtonCell) SetFont(value IFont) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setFont:"), value)
}/* debug [instance_properties/setter]: font */


// The cell’s value as an integer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/intvalue
func (b_ ButtonCell) IntValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](b_.ID, objc.Sel("intValue"))
	return rv
}/* debug [instance_properties/getter]: intValue */


// The cell’s value as an integer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/intvalue
func (b_ ButtonCell) SetIntValue(value objectivec.IObject) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIntValue:"), value)
}/* debug [instance_properties/setter]: intValue */


// The cell’s current state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/state
func (b_ ButtonCell) State() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](b_.ID, objc.Sel("state"))
	return rv
}/* debug [instance_properties/getter]: state */


// The cell’s current state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/state
func (b_ ButtonCell) SetState(value objectivec.IObject) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setState:"), value)
}/* debug [instance_properties/setter]: state */


// The value of the receiver’s cell as a double-precision floating-point number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/doublevalue
func (b_ ButtonCell) DoubleValue() float64 {
	rv := objc.Send[float64](b_.ID, objc.Sel("doubleValue"))
	return rv
}/* debug [instance_properties/getter]: doubleValue */


// The value of the receiver’s cell as a double-precision floating-point number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/doublevalue
func (b_ ButtonCell) SetDoubleValue(value float64) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setDoubleValue:"), value)
}/* debug [instance_properties/setter]: doubleValue */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSButtonCell */


