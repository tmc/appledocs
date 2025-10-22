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

// The class instance for the [Button] class.
var (
	ButtonClass     _ButtonClass
	ButtonClassOnce sync.Once
)

func getButtonClass() _ButtonClass {
	ButtonClassOnce.Do(func() {
		ButtonClass = _ButtonClass{objc.GetClass("NSButton")}
	})
	return ButtonClass
}

type _ButtonClass struct {
	class objc.Class
}

// An interface definition for the [Button] class.
type IButton interface {
	IControl
	CompressWithPrioritizedCompressionOptions(prioritizedOptions []UserInterfaceCompressionOptions)
	GetPeriodicDelayInterval(delay unsafe.Pointer, interval unsafe.Pointer)
	Highlight(flag bool)
	MinimumSizeWithPrioritizedCompressionOptions(prioritizedOptions []UserInterfaceCompressionOptions) coregraphics.CGSize
	PerformKeyEquivalent(key IEvent) bool
	SetButtonType(type_ ButtonType)
	SetNextState()
	SetPeriodicDelayInterval(delay float32, interval float32)
	SetTitleWithMnemonic(stringWithAmpersand string)
	ActiveCompressionOptions() NSUserInterfaceCompressionOptions
	AllowsMixedState() bool
	SetAllowsMixedState(value bool)
	AlternateImage() Image
	SetAlternateImage(value IImage)
	AlternateTitle() string
	SetAlternateTitle(value string)
	AttributedAlternateTitle() foundation.AttributedString
	SetAttributedAlternateTitle(value foundation.IAttributedString)
	AttributedTitle() foundation.AttributedString
	SetAttributedTitle(value foundation.IAttributedString)
	BezelColor() NSColor
	SetBezelColor(value IColor)
	BezelStyle() BezelStyle
	SetBezelStyle(value BezelStyle)
	BorderShape() ControlBorderShape
	SetBorderShape(value IControlBorderShape)
	ContentTintColor() NSColor
	SetContentTintColor(value IColor)
	HasDestructiveAction() bool
	SetHasDestructiveAction(value bool)
	Image() Image
	SetImage(value IImage)
	ImageHugsTitle() bool
	SetImageHugsTitle(value bool)
	ImagePosition() CellImagePosition
	SetImagePosition(value CellImagePosition)
	ImageScaling() ImageScaling
	SetImageScaling(value ImageScaling)
	Bordered() bool
	SetBordered(value bool)
	SpringLoaded() bool
	SetSpringLoaded(value bool)
	Transparent() bool
	SetTransparent(value bool)
	KeyEquivalent() string
	SetKeyEquivalent(value string)
	KeyEquivalentModifierMask() EventModifierFlags
	SetKeyEquivalentModifierMask(value EventModifierFlags)
	MaxAcceleratorLevel() int
	SetMaxAcceleratorLevel(value int)
	ShowsBorderOnlyWhileMouseInside() bool
	SetShowsBorderOnlyWhileMouseInside(value bool)
	Sound() NSSound
	SetSound(value ISound)
	State() ControlStateValue
	SetState(value IControlStateValue)
	SymbolConfiguration() ImageSymbolConfiguration
	SetSymbolConfiguration(value IImageSymbolConfiguration)
	TintProminence() TintProminence
	SetTintProminence(value ITintProminence)
	Title() string
	SetTitle(value string)
	IsBordered() bool
	SetIsBordered(value bool)
	IsSpringLoaded() bool
	SetIsSpringLoaded(value bool)
	IsTransparent() bool
	SetIsTransparent(value bool)
}

// A control that defines an area on the screen that a user clicks to trigger an action.
//
// Buttons are a standard control for initiating actions within your app. You can configure buttons with many different visual styles, but the behavior is the same. When a user clicks it, a button calls the action method of its associated target object. (If you configure a button as continuous, it calls its action method at timed intervals until the user releases the mouse button or the cursor leaves the button boundaries). You use the action method to perform your app-specific tasks. There are multiple types of buttons, each with a different user interface and behavior. The class defines the button types, and calling the method configures them. If you configure it as an accelerator button (type or ), you can set a button to send action messages when changes in pressure occur when the user clicks the button. Buttons can either have two states (on and off) or three states (on, off, and mixed). You enable a three-state button by calling the method. On and off (also referred to as alternate and normal) states indicate that the user clicked or didn’t click the button. Mixed is typically used for checkboxes or radio buttons, which allow for an additional intermediate state. For example, suppose the state of a checkbox denotes whether a text field contains bold text. If all text in the text field is bold, then the checkbox is on. If none of the text is bold, then the checkbox is off. If some of the text is bold, then the checkbox is mixed. For most types of buttons, the value of the button matches its state—the value is for on, for off, or for mixed. For pressure-sensitive buttons, the value of the button indicates pressure level instead. and both provide a control view, which displays an object. However, while a matrix requires you to access the button cell objects directly, most button class methods act as “covers” for identically declared button cell methods. In other words, the implementation of the button method invokes the corresponding button cell method for you, allowing you to be unconcerned with the existence of the button cell. The only button cell methods that don’t have covers relate to the font you use to display the key equivalent and to specific methods for highlighting or showing the state of the button.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton
type Button struct {
	Control
}

// ButtonFrom constructs a [Button] from an unsafe.Pointer.
//
// A control that defines an area on the screen that a user clicks to trigger an action.
func ButtonFrom(ptr unsafe.Pointer) Button {
	return Button{
		Control: ControlFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (bc _ButtonClass) Alloc() Button {
	rv := objc.Send[Button](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _ButtonClass) New() Button {
	rv := objc.Send[Button](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ Button) Init() Button {
	rv := objc.Send[Button](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ Button) Autorelease() Button {
	rv := objc.Send[Button](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewButton creates a new Button instance.
func NewButton() Button {
	return getButtonClass().New()
}




// Creates a standard checkbox with the title you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/init(checkboxWithTitle:target:action:)
func NewButtonCheckboxWithTitleTargetAction(title string, target objectivec.IObject, action objc.SEL) Button {
	rv := objc.Send[Button](objc.ID(getButtonClass().class), objc.Sel("checkboxWithTitle:target:action:"), objc.String(title), target, action)
	return rv
}



// Creates a standard radio button with the title you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/init(radioButtonWithTitle:target:action:)
func NewButtonRadioButtonWithTitleTargetAction(title string, target objectivec.IObject, action objc.SEL) Button {
	rv := objc.Send[Button](objc.ID(getButtonClass().class), objc.Sel("radioButtonWithTitle:target:action:"), objc.String(title), target, action)
	return rv
}



// Creates a standard push button with the image you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/init(image:target:action:)
func NewButtonWithImageTargetAction(image IImage, target objectivec.IObject, action objc.SEL) Button {
	rv := objc.Send[Button](objc.ID(getButtonClass().class), objc.Sel("buttonWithImage:target:action:"), image, target, action)
	return rv
}



// Creates a standard push button with a title and image.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/init(title:image:target:action:)
func NewButtonWithTitleImageTargetAction(title string, image IImage, target objectivec.IObject, action objc.SEL) Button {
	rv := objc.Send[Button](objc.ID(getButtonClass().class), objc.Sel("buttonWithTitle:image:target:action:"), objc.String(title), image, target, action)
	return rv
}



// Creates a standard push button with the title you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/init(title:target:action:)
func NewButtonWithTitleTargetAction(title string, target objectivec.IObject, action objc.SEL) Button {
	rv := objc.Send[Button](objc.ID(getButtonClass().class), objc.Sel("buttonWithTitle:target:action:"), objc.String(title), target, action)
	return rv
}


// Creates a standard checkbox with the title you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/init(checkboxWithTitle:target:action:)
func (bc _ButtonClass) CheckboxWithTitleTargetAction(title string, target objectivec.IObject, action objc.SEL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("checkboxWithTitle:target:action:"), objc.String(title), target, action)
	return rv
}

// Creates a standard push button with the image you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/init(image:target:action:)
func (bc _ButtonClass) ButtonWithImageTargetAction(image IImage, target objectivec.IObject, action objc.SEL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("buttonWithImage:target:action:"), image, target, action)
	return rv
}

// Creates a standard radio button with the title you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/init(radioButtonWithTitle:target:action:)
func (bc _ButtonClass) RadioButtonWithTitleTargetAction(title string, target objectivec.IObject, action objc.SEL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("radioButtonWithTitle:target:action:"), objc.String(title), target, action)
	return rv
}

// Creates a standard push button with a title and image.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/init(title:image:target:action:)
func (bc _ButtonClass) ButtonWithTitleImageTargetAction(title string, image IImage, target objectivec.IObject, action objc.SEL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("buttonWithTitle:image:target:action:"), objc.String(title), image, target, action)
	return rv
}

// Creates a standard push button with the title you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/init(title:target:action:)
func (bc _ButtonClass) ButtonWithTitleTargetAction(title string, target objectivec.IObject, action objc.SEL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("buttonWithTitle:target:action:"), objc.String(title), target, action)
	return rv
}

// Sets the priority compression options for this button.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/compress(withPrioritizedCompressionOptions:)
func (b_ Button) CompressWithPrioritizedCompressionOptions(prioritizedOptions []UserInterfaceCompressionOptions) {
	objc.Send[objc.ID](b_.ID, objc.Sel("compressWithPrioritizedCompressionOptions:"), prioritizedOptions)
}

// Returns by reference the delay and interval periods for a continuous button.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/getPeriodicDelay(_:interval:)
func (b_ Button) GetPeriodicDelayInterval(delay unsafe.Pointer, interval unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("getPeriodicDelay:interval:"), delay, interval)
}

// Highlights (or unhighlights) the button.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/highlight(_:)
func (b_ Button) Highlight(flag bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("highlight:"), flag)
}

// Returns the minimum size of the button by using the compression options.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/minimumSize(withPrioritizedCompressionOptions:)
func (b_ Button) MinimumSizeWithPrioritizedCompressionOptions(prioritizedOptions []UserInterfaceCompressionOptions) coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](b_.ID, objc.Sel("minimumSizeWithPrioritizedCompressionOptions:"), prioritizedOptions)
	return rv
}

// Checks the button’s key equivalent against the specified event and, if they match, simulates the button being clicked.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/performKeyEquivalent(with:)
func (b_ Button) PerformKeyEquivalent(key IEvent) bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("performKeyEquivalent:"), key)
	return rv
}

// Sets the button’s type, which affects its user interface and behavior when clicked.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/setButtonType(_:)
func (b_ Button) SetButtonType(type_ ButtonType) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setButtonType:"), type_)
}

// Sets the button to its next state.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/setNextState()
func (b_ Button) SetNextState() {
	objc.Send[objc.ID](b_.ID, objc.Sel("setNextState"))
}

// Sets the message delay and interval periods for a continuous button.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/setPeriodicDelay(_:interval:)
func (b_ Button) SetPeriodicDelayInterval(delay float32, interval float32) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPeriodicDelay:interval:"), delay, interval)
}

// Sets the title of a button with a character denoting an access key.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/setTitleWithMnemonic:
func (b_ Button) SetTitleWithMnemonic(stringWithAmpersand string) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTitleWithMnemonic:"), objc.String(stringWithAmpersand))
}

// The compression options active for this button.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/activeCompressionOptions
func (b_ Button) ActiveCompressionOptions() NSUserInterfaceCompressionOptions {
	rv := objc.Send[NSUserInterfaceCompressionOptions](b_.ID, objc.Sel("activeCompressionOptions"))
	return rv
}

// A Boolean value that indicates whether the button allows a mixed state.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/allowsMixedState
func (b_ Button) AllowsMixedState() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("allowsMixedState"))
	return rv
}


// SetAllowsMixedState sets the value of the allowsMixedState property.
// A Boolean value that indicates whether the button allows a mixed state.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/allowsMixedState
func (b_ Button) SetAllowsMixedState(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAllowsMixedState:"), value)
}

// An alternate image that appears on the button when the button is in an on state.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/alternateImage
func (b_ Button) AlternateImage() Image {
	rv := objc.Send[Image](b_.ID, objc.Sel("alternateImage"))
	return rv
}


// SetAlternateImage sets the value of the alternateImage property.
// An alternate image that appears on the button when the button is in an on state.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/alternateImage
func (b_ Button) SetAlternateImage(value IImage) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAlternateImage:"), value)
}

// The title that the button displays when the button is in an on state.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/alternateTitle
func (b_ Button) AlternateTitle() string {
	rv := objc.Send[string](b_.ID, objc.Sel("alternateTitle"))
	return rv
}


// SetAlternateTitle sets the value of the alternateTitle property.
// The title that the button displays when the button is in an on state.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/alternateTitle
func (b_ Button) SetAlternateTitle(value string) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAlternateTitle:"), objc.String(value))
}

// The title that the button displays as an attributed string when the button is in an on state.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/attributedAlternateTitle
func (b_ Button) AttributedAlternateTitle() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](b_.ID, objc.Sel("attributedAlternateTitle"))
	return rv
}


// SetAttributedAlternateTitle sets the value of the attributedAlternateTitle property.
// The title that the button displays as an attributed string when the button is in an on state.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/attributedAlternateTitle
func (b_ Button) SetAttributedAlternateTitle(value foundation.IAttributedString) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAttributedAlternateTitle:"), value)
}

// The title that the button displays in an off state, as an attributed string.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/attributedTitle
func (b_ Button) AttributedTitle() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](b_.ID, objc.Sel("attributedTitle"))
	return rv
}


// SetAttributedTitle sets the value of the attributedTitle property.
// The title that the button displays in an off state, as an attributed string.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/attributedTitle
func (b_ Button) SetAttributedTitle(value foundation.IAttributedString) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAttributedTitle:"), value)
}

// The color of the button’s bezel, in appearances that support it.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/bezelColor
func (b_ Button) BezelColor() NSColor {
	rv := objc.Send[NSColor](b_.ID, objc.Sel("bezelColor"))
	return rv
}


// SetBezelColor sets the value of the bezelColor property.
// The color of the button’s bezel, in appearances that support it.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/bezelColor
func (b_ Button) SetBezelColor(value IColor) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBezelColor:"), value)
}

// The appearance of the button’s border.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/bezelStyle-swift.property
func (b_ Button) BezelStyle() BezelStyle {
	rv := objc.Send[BezelStyle](b_.ID, objc.Sel("bezelStyle"))
	return rv
}


// SetBezelStyle sets the value of the bezelStyle property.
// The appearance of the button’s border.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/bezelStyle-swift.property
func (b_ Button) SetBezelStyle(value BezelStyle) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBezelStyle:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/borderShape
func (b_ Button) BorderShape() ControlBorderShape {
	rv := objc.Send[ControlBorderShape](b_.ID, objc.Sel("borderShape"))
	return rv
}


// SetBorderShape sets the value of the borderShape property.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/borderShape
func (b_ Button) SetBorderShape(value IControlBorderShape) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBorderShape:"), value)
}

// A tint color to use for the template image and text content.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/contentTintColor
func (b_ Button) ContentTintColor() NSColor {
	rv := objc.Send[NSColor](b_.ID, objc.Sel("contentTintColor"))
	return rv
}


// SetContentTintColor sets the value of the contentTintColor property.
// A tint color to use for the template image and text content.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/contentTintColor
func (b_ Button) SetContentTintColor(value IColor) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setContentTintColor:"), value)
}

// A Boolean value that defines whether a button’s action has a destructive effect.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/hasDestructiveAction
func (b_ Button) HasDestructiveAction() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("hasDestructiveAction"))
	return rv
}


// SetHasDestructiveAction sets the value of the hasDestructiveAction property.
// A Boolean value that defines whether a button’s action has a destructive effect.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/hasDestructiveAction
func (b_ Button) SetHasDestructiveAction(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setHasDestructiveAction:"), value)
}

// The image that appears on the button when it’s in an off state, or if there is no such image.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/image
func (b_ Button) Image() Image {
	rv := objc.Send[Image](b_.ID, objc.Sel("image"))
	return rv
}


// SetImage sets the value of the image property.
// The image that appears on the button when it’s in an off state, or if there is no such image.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/image
func (b_ Button) SetImage(value IImage) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setImage:"), value)
}

// A Boolean value that determines how the button’s image and title are positioned together within the button bezel.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/imageHugsTitle
func (b_ Button) ImageHugsTitle() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("imageHugsTitle"))
	return rv
}


// SetImageHugsTitle sets the value of the imageHugsTitle property.
// A Boolean value that determines how the button’s image and title are positioned together within the button bezel.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/imageHugsTitle
func (b_ Button) SetImageHugsTitle(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setImageHugsTitle:"), value)
}

// The position of the button’s image relative to its title.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/imagePosition
func (b_ Button) ImagePosition() CellImagePosition {
	rv := objc.Send[CellImagePosition](b_.ID, objc.Sel("imagePosition"))
	return rv
}


// SetImagePosition sets the value of the imagePosition property.
// The position of the button’s image relative to its title.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/imagePosition
func (b_ Button) SetImagePosition(value CellImagePosition) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setImagePosition:"), value)
}

// The scaling mode applied to make the cell’s image fit the frame of the image view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/imageScaling
func (b_ Button) ImageScaling() ImageScaling {
	rv := objc.Send[ImageScaling](b_.ID, objc.Sel("imageScaling"))
	return rv
}


// SetImageScaling sets the value of the imageScaling property.
// The scaling mode applied to make the cell’s image fit the frame of the image view.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/imageScaling
func (b_ Button) SetImageScaling(value ImageScaling) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setImageScaling:"), value)
}

// A Boolean value that determines whether the button has a border.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/isBordered
func (b_ Button) Bordered() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("bordered"))
	return rv
}


// SetBordered sets the value of the bordered property.
// A Boolean value that determines whether the button has a border.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/isBordered
func (b_ Button) SetBordered(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBordered:"), value)
}

// A Boolean value that indicates whether spring loading is enabled for the button.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/isSpringLoaded
func (b_ Button) SpringLoaded() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("springLoaded"))
	return rv
}


// SetSpringLoaded sets the value of the springLoaded property.
// A Boolean value that indicates whether spring loading is enabled for the button.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/isSpringLoaded
func (b_ Button) SetSpringLoaded(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSpringLoaded:"), value)
}

// A Boolean value that indicates whether the button is transparent.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/isTransparent
func (b_ Button) Transparent() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("transparent"))
	return rv
}


// SetTransparent sets the value of the transparent property.
// A Boolean value that indicates whether the button is transparent.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/isTransparent
func (b_ Button) SetTransparent(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTransparent:"), value)
}

// The key-equivalent character of the button.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/keyEquivalent
func (b_ Button) KeyEquivalent() string {
	rv := objc.Send[string](b_.ID, objc.Sel("keyEquivalent"))
	return rv
}


// SetKeyEquivalent sets the value of the keyEquivalent property.
// The key-equivalent character of the button.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/keyEquivalent
func (b_ Button) SetKeyEquivalent(value string) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setKeyEquivalent:"), objc.String(value))
}

// The mask specifying the modifier keys for the button’s key equivalent.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/keyEquivalentModifierMask
func (b_ Button) KeyEquivalentModifierMask() EventModifierFlags {
	rv := objc.Send[EventModifierFlags](b_.ID, objc.Sel("keyEquivalentModifierMask"))
	return rv
}


// SetKeyEquivalentModifierMask sets the value of the keyEquivalentModifierMask property.
// The mask specifying the modifier keys for the button’s key equivalent.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/keyEquivalentModifierMask
func (b_ Button) SetKeyEquivalentModifierMask(value EventModifierFlags) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setKeyEquivalentModifierMask:"), value)
}

// An integer value indicating the maximum pressure level for a button of type .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/maxAcceleratorLevel
func (b_ Button) MaxAcceleratorLevel() int {
	rv := objc.Send[int](b_.ID, objc.Sel("maxAcceleratorLevel"))
	return rv
}


// SetMaxAcceleratorLevel sets the value of the maxAcceleratorLevel property.
// An integer value indicating the maximum pressure level for a button of type .

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/maxAcceleratorLevel
func (b_ Button) SetMaxAcceleratorLevel(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setMaxAcceleratorLevel:"), value)
}

// A Boolean value that determines whether the button displays its border only when the pointer is over it.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/showsBorderOnlyWhileMouseInside
func (b_ Button) ShowsBorderOnlyWhileMouseInside() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("showsBorderOnlyWhileMouseInside"))
	return rv
}


// SetShowsBorderOnlyWhileMouseInside sets the value of the showsBorderOnlyWhileMouseInside property.
// A Boolean value that determines whether the button displays its border only when the pointer is over it.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/showsBorderOnlyWhileMouseInside
func (b_ Button) SetShowsBorderOnlyWhileMouseInside(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setShowsBorderOnlyWhileMouseInside:"), value)
}

// The sound that plays when the user clicks the button.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/sound
func (b_ Button) Sound() NSSound {
	rv := objc.Send[NSSound](b_.ID, objc.Sel("sound"))
	return rv
}


// SetSound sets the value of the sound property.
// The sound that plays when the user clicks the button.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/sound
func (b_ Button) SetSound(value ISound) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSound:"), value)
}

// The button’s state.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/state
func (b_ Button) State() ControlStateValue {
	rv := objc.Send[ControlStateValue](b_.ID, objc.Sel("state"))
	return rv
}


// SetState sets the value of the state property.
// The button’s state.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/state
func (b_ Button) SetState(value IControlStateValue) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setState:"), value)
}

// The combination of point size, weight, and scale to use when sizing and displaying symbol images.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/symbolConfiguration
func (b_ Button) SymbolConfiguration() ImageSymbolConfiguration {
	rv := objc.Send[ImageSymbolConfiguration](b_.ID, objc.Sel("symbolConfiguration"))
	return rv
}


// SetSymbolConfiguration sets the value of the symbolConfiguration property.
// The combination of point size, weight, and scale to use when sizing and displaying symbol images.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/symbolConfiguration
func (b_ Button) SetSymbolConfiguration(value IImageSymbolConfiguration) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSymbolConfiguration:"), value)
}

// The tint prominence of the button. Use tint prominence to gently suggest a hierarchy when multiple buttons perform similar actions. A button with primary tint prominence suggests the most preferred option, while secondary prominence indicates a reasonable alternative. See for a list of possible values.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/tintProminence
func (b_ Button) TintProminence() TintProminence {
	rv := objc.Send[TintProminence](b_.ID, objc.Sel("tintProminence"))
	return rv
}


// SetTintProminence sets the value of the tintProminence property.
// The tint prominence of the button. Use tint prominence to gently suggest a hierarchy when multiple buttons perform similar actions. A button with primary tint prominence suggests the most preferred option, while secondary prominence indicates a reasonable alternative. See for a list of possible values.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/tintProminence
func (b_ Button) SetTintProminence(value ITintProminence) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTintProminence:"), value)
}

// The title displayed on the button when it’s in an off state.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/title
func (b_ Button) Title() string {
	rv := objc.Send[string](b_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
// The title displayed on the button when it’s in an off state.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/title
func (b_ Button) SetTitle(value string) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTitle:"), objc.String(value))
}

// A Boolean value that determines whether the button has a border.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/isbordered
func (b_ Button) IsBordered() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isBordered"))
	return rv
}


// SetIsBordered sets the value of the isBordered property.
// A Boolean value that determines whether the button has a border.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/isbordered
func (b_ Button) SetIsBordered(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsBordered:"), value)
}

// A Boolean value that indicates whether spring loading is enabled for the button.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/isspringloaded
func (b_ Button) IsSpringLoaded() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isSpringLoaded"))
	return rv
}


// SetIsSpringLoaded sets the value of the isSpringLoaded property.
// A Boolean value that indicates whether spring loading is enabled for the button.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/isspringloaded
func (b_ Button) SetIsSpringLoaded(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsSpringLoaded:"), value)
}

// A Boolean value that indicates whether the button is transparent.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/istransparent
func (b_ Button) IsTransparent() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isTransparent"))
	return rv
}


// SetIsTransparent sets the value of the isTransparent property.
// A Boolean value that indicates whether the button is transparent.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/istransparent
func (b_ Button) SetIsTransparent(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsTransparent:"), value)
}


