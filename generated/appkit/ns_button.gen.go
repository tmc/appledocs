// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSButton */


/* debug [class_header]: Header for NSButton */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Button */
// An interface definition for the [Button] class.
type IButton interface {
	IControl
	
/* debug [class_interface_properties]: Properties for Button */
	// properties:
	ActiveCompressionOptions() IUserInterfaceCompressionOptions
	AllowsMixedState() bool
	SetAllowsMixedState(value bool)
	AlternateImage() IImage
	SetAlternateImage(value IImage)
	AlternateTitle() objc.IObject /* cross-framework: NSString */
	SetAlternateTitle(value objc.IObject /* cross-framework: NSString */)
	AttributedAlternateTitle() foundation.AttributedString
	SetAttributedAlternateTitle(value foundation.AttributedString)
	AttributedTitle() foundation.AttributedString
	SetAttributedTitle(value foundation.AttributedString)
	BezelColor() IColor
	SetBezelColor(value IColor)
	BezelStyle() BezelStyle
	SetBezelStyle(value BezelStyle)
	BorderShape() ControlBorderShape
	SetBorderShape(value ControlBorderShape)
	ContentTintColor() IColor
	SetContentTintColor(value IColor)
	HasDestructiveAction() bool
	SetHasDestructiveAction(value bool)
	Image() IImage
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
	KeyEquivalent() objc.IObject /* cross-framework: NSString */
	SetKeyEquivalent(value objc.IObject /* cross-framework: NSString */)
	KeyEquivalentModifierMask() EventModifierFlags
	SetKeyEquivalentModifierMask(value EventModifierFlags)
	MaxAcceleratorLevel() int
	SetMaxAcceleratorLevel(value int)
	ShowsBorderOnlyWhileMouseInside() bool
	SetShowsBorderOnlyWhileMouseInside(value bool)
	Sound() ISound
	SetSound(value ISound)
	State() ControlStateValue /* typedef */
	SetState(value ControlStateValue /* typedef */)
	SymbolConfiguration() IImageSymbolConfiguration
	SetSymbolConfiguration(value IImageSymbolConfiguration)
	TintProminence() TintProminence
	SetTintProminence(value TintProminence)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	IsBordered() bool
	SetIsBordered(value bool)
	IsSpringLoaded() bool
	SetIsSpringLoaded(value bool)
	IsTransparent() bool
	SetIsTransparent(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Button */
	// methods:
	CompressWithPrioritizedCompressionOptions(prioritizedOptions []UserInterfaceCompressionOptions)
	GetPeriodicDelayInterval(delay objectivec.IObject, interval objectivec.IObject)
	Highlight(flag bool)
	MinimumSizeWithPrioritizedCompressionOptions(prioritizedOptions []UserInterfaceCompressionOptions) Size /* not a class type */
	PerformKeyEquivalent(key IEvent) bool
	SetButtonType(type_ ButtonType)
	SetNextState()
	SetPeriodicDelayInterval(delay float32, interval float32)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Button */
// Alloc allocates a new instance without initialization.
func (bc _ButtonClass) Alloc() Button {
	rv := objc.Send[Button](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Button */
// A control that defines an area on the screen that a user clicks to trigger an action.
//
// Buttons are a standard control for initiating actions within your app. You can configure buttons with many different visual styles, but the behavior is the same. When a user clicks it, a button calls the action method of its associated target object. (If you configure a button as continuous, it calls its action method at timed intervals until the user releases the mouse button or the cursor leaves the button boundaries). You use the action method to perform your app-specific tasks. There are multiple types of buttons, each with a different user interface and behavior. The class defines the button types, and calling the method configures them. If you configure it as an accelerator button (type or ), you can set a button to send action messages when changes in pressure occur when the user clicks the button. Buttons can either have two states (on and off) or three states (on, off, and mixed). You enable a three-state button by calling the method. On and off (also referred to as alternate and normal) states indicate that the user clicked or didn’t click the button. Mixed is typically used for checkboxes or radio buttons, which allow for an additional intermediate state. For example, suppose the state of a checkbox denotes whether a text field contains bold text. If all text in the text field is bold, then the checkbox is on. If none of the text is bold, then the checkbox is off. If some of the text is bold, then the checkbox is mixed. For most types of buttons, the value of the button matches its state—the value is for on, for off, or for mixed. For pressure-sensitive buttons, the value of the button indicates pressure level instead. and both provide a control view, which displays an object. However, while a matrix requires you to access the button cell objects directly, most button class methods act as “covers” for identically declared button cell methods. In other words, the implementation of the button method invokes the corresponding button cell method for you, allowing you to be unconcerned with the existence of the button cell. The only button cell methods that don’t have covers relate to the font you use to display the key equivalent and to specific methods for highlighting or showing the state of the button.


// A control that defines an area on the screen that a user clicks to trigger an action.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Button */

// Creates a standard checkbox with the title you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/init(checkboxWithTitle:target:action:)
func NewButtonCheckboxWithTitleTargetAction(title objc.IObject /* cross-framework: NSString */, target objc.IObject, action objc.SEL) Button {
	rv := objc.Send[Button](objc.ID(getButtonClass().class), objc.Sel("checkboxWithTitle:target:action:"), title, target, action)
	return rv
}/* debug [class_init_methods/constructor]: NewButtonCheckboxWithTitleTargetAction */


// Creates a standard radio button with the title you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/init(radioButtonWithTitle:target:action:)
func NewButtonRadioButtonWithTitleTargetAction(title objc.IObject /* cross-framework: NSString */, target objc.IObject, action objc.SEL) Button {
	rv := objc.Send[Button](objc.ID(getButtonClass().class), objc.Sel("radioButtonWithTitle:target:action:"), title, target, action)
	return rv
}/* debug [class_init_methods/constructor]: NewButtonRadioButtonWithTitleTargetAction */


// Creates a standard push button with the image you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/init(image:target:action:)
func NewButtonWithImageTargetAction(image IImage, target objc.IObject, action objc.SEL) Button {
	rv := objc.Send[Button](objc.ID(getButtonClass().class), objc.Sel("buttonWithImage:target:action:"), image, target, action)
	return rv
}/* debug [class_init_methods/constructor]: NewButtonWithImageTargetAction */


// Creates a standard push button with a title and image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/init(title:image:target:action:)
func NewButtonWithTitleImageTargetAction(title objc.IObject /* cross-framework: NSString */, image IImage, target objc.IObject, action objc.SEL) Button {
	rv := objc.Send[Button](objc.ID(getButtonClass().class), objc.Sel("buttonWithTitle:image:target:action:"), title, image, target, action)
	return rv
}/* debug [class_init_methods/constructor]: NewButtonWithTitleImageTargetAction */


// Creates a standard push button with the title you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/init(title:target:action:)
func NewButtonWithTitleTargetAction(title objc.IObject /* cross-framework: NSString */, target objc.IObject, action objc.SEL) Button {
	rv := objc.Send[Button](objc.ID(getButtonClass().class), objc.Sel("buttonWithTitle:target:action:"), title, target, action)
	return rv
}/* debug [class_init_methods/constructor]: NewButtonWithTitleTargetAction */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Button */

// Creates a standard checkbox with the title you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/init(checkboxWithTitle:target:action:)
func (bc _ButtonClass) CheckboxWithTitleTargetAction(title objc.IObject /* cross-framework: NSString */, target objc.IObject, action objc.SEL) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(bc.class), objc.Sel("checkboxWithTitle:target:action:"), title, target, action)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CheckboxWithTitleTargetAction) */


// Creates a standard push button with the image you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/init(image:target:action:)
func (bc _ButtonClass) ButtonWithImageTargetAction(image IImage, target objc.IObject, action objc.SEL) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(bc.class), objc.Sel("buttonWithImage:target:action:"), image, target, action)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ButtonWithImageTargetAction) */


// Creates a standard radio button with the title you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/init(radioButtonWithTitle:target:action:)
func (bc _ButtonClass) RadioButtonWithTitleTargetAction(title objc.IObject /* cross-framework: NSString */, target objc.IObject, action objc.SEL) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(bc.class), objc.Sel("radioButtonWithTitle:target:action:"), title, target, action)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RadioButtonWithTitleTargetAction) */


// Creates a standard push button with a title and image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/init(title:image:target:action:)
func (bc _ButtonClass) ButtonWithTitleImageTargetAction(title objc.IObject /* cross-framework: NSString */, image IImage, target objc.IObject, action objc.SEL) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(bc.class), objc.Sel("buttonWithTitle:image:target:action:"), title, image, target, action)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ButtonWithTitleImageTargetAction) */


// Creates a standard push button with the title you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/init(title:target:action:)
func (bc _ButtonClass) ButtonWithTitleTargetAction(title objc.IObject /* cross-framework: NSString */, target objc.IObject, action objc.SEL) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(bc.class), objc.Sel("buttonWithTitle:target:action:"), title, target, action)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ButtonWithTitleTargetAction) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Button */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Button */

// Sets the priority compression options for this button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/compress(withPrioritizedCompressionOptions:)
func (b_ Button) CompressWithPrioritizedCompressionOptions(prioritizedOptions []UserInterfaceCompressionOptions) {
	objc.Send[objc.ID](b_.ID, objc.Sel("compressWithPrioritizedCompressionOptions:"), prioritizedOptions)
}/* debug [instance_methods/method]: CompressWithPrioritizedCompressionOptions */


// Returns by reference the delay and interval periods for a continuous button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/getPeriodicDelay(_:interval:)
func (b_ Button) GetPeriodicDelayInterval(delay objectivec.IObject, interval objectivec.IObject) {
	objc.Send[objc.ID](b_.ID, objc.Sel("getPeriodicDelay:interval:"), delay, interval)
}/* debug [instance_methods/method]: GetPeriodicDelayInterval */


// Highlights (or unhighlights) the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/highlight(_:)
func (b_ Button) Highlight(flag bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("highlight:"), flag)
}/* debug [instance_methods/method]: Highlight */


// Returns the minimum size of the button by using the compression options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/minimumSize(withPrioritizedCompressionOptions:)
func (b_ Button) MinimumSizeWithPrioritizedCompressionOptions(prioritizedOptions []UserInterfaceCompressionOptions) Size /* not a class type */ {
	rv := objc.Send[Size](b_.ID, objc.Sel("minimumSizeWithPrioritizedCompressionOptions:"), prioritizedOptions)
	return rv
}/* debug [instance_methods/method]: MinimumSizeWithPrioritizedCompressionOptions */


// Checks the button’s key equivalent against the specified event and, if they match, simulates the button being clicked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/performKeyEquivalent(with:)
func (b_ Button) PerformKeyEquivalent(key IEvent) bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("performKeyEquivalent:"), key)
	return rv
}/* debug [instance_methods/method]: PerformKeyEquivalent */


// Sets the button’s type, which affects its user interface and behavior when clicked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/setButtonType(_:)
func (b_ Button) SetButtonType(type_ ButtonType) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setButtonType:"), type_)
}/* debug [instance_methods/method]: SetButtonType */


// Sets the button to its next state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/setNextState()
func (b_ Button) SetNextState() {
	objc.Send[objc.ID](b_.ID, objc.Sel("setNextState"))
}/* debug [instance_methods/method]: SetNextState */


// Sets the message delay and interval periods for a continuous button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/setPeriodicDelay(_:interval:)
func (b_ Button) SetPeriodicDelayInterval(delay float32, interval float32) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPeriodicDelay:interval:"), delay, interval)
}/* debug [instance_methods/method]: SetPeriodicDelayInterval */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Button */

// The compression options active for this button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/activeCompressionOptions
func (b_ Button) ActiveCompressionOptions() IUserInterfaceCompressionOptions {
	rv := objc.Send[UserInterfaceCompressionOptions](b_.ID, objc.Sel("activeCompressionOptions"))
	return rv
}/* debug [instance_properties/getter]: activeCompressionOptions */


// A Boolean value that indicates whether the button allows a mixed state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/allowsMixedState
func (b_ Button) AllowsMixedState() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("allowsMixedState"))
	return rv
}/* debug [instance_properties/getter]: allowsMixedState */


// A Boolean value that indicates whether the button allows a mixed state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/allowsMixedState
func (b_ Button) SetAllowsMixedState(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAllowsMixedState:"), value)
}/* debug [instance_properties/setter]: allowsMixedState */


// An alternate image that appears on the button when the button is in an on state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/alternateImage
func (b_ Button) AlternateImage() IImage {
	rv := objc.Send[Image](b_.ID, objc.Sel("alternateImage"))
	return rv
}/* debug [instance_properties/getter]: alternateImage */


// An alternate image that appears on the button when the button is in an on state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/alternateImage
func (b_ Button) SetAlternateImage(value IImage) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAlternateImage:"), value)
}/* debug [instance_properties/setter]: alternateImage */


// The title that the button displays when the button is in an on state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/alternateTitle
func (b_ Button) AlternateTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](b_.ID, objc.Sel("alternateTitle"))
	return rv
}/* debug [instance_properties/getter]: alternateTitle */


// The title that the button displays when the button is in an on state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/alternateTitle
func (b_ Button) SetAlternateTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAlternateTitle:"), value)
}/* debug [instance_properties/setter]: alternateTitle */


// The title that the button displays as an attributed string when the button is in an on state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/attributedAlternateTitle
func (b_ Button) AttributedAlternateTitle() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](b_.ID, objc.Sel("attributedAlternateTitle"))
	return rv
}/* debug [instance_properties/getter]: attributedAlternateTitle */


// The title that the button displays as an attributed string when the button is in an on state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/attributedAlternateTitle
func (b_ Button) SetAttributedAlternateTitle(value foundation.AttributedString) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAttributedAlternateTitle:"), value)
}/* debug [instance_properties/setter]: attributedAlternateTitle */


// The title that the button displays in an off state, as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/attributedTitle
func (b_ Button) AttributedTitle() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](b_.ID, objc.Sel("attributedTitle"))
	return rv
}/* debug [instance_properties/getter]: attributedTitle */


// The title that the button displays in an off state, as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/attributedTitle
func (b_ Button) SetAttributedTitle(value foundation.AttributedString) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAttributedTitle:"), value)
}/* debug [instance_properties/setter]: attributedTitle */


// The color of the button’s bezel, in appearances that support it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/bezelColor
func (b_ Button) BezelColor() IColor {
	rv := objc.Send[Color](b_.ID, objc.Sel("bezelColor"))
	return rv
}/* debug [instance_properties/getter]: bezelColor */


// The color of the button’s bezel, in appearances that support it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/bezelColor
func (b_ Button) SetBezelColor(value IColor) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBezelColor:"), value)
}/* debug [instance_properties/setter]: bezelColor */


// The appearance of the button’s border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/bezelStyle-swift.property
func (b_ Button) BezelStyle() BezelStyle {
	rv := objc.Send[BezelStyle](b_.ID, objc.Sel("bezelStyle"))
	return rv
}/* debug [instance_properties/getter]: bezelStyle */


// The appearance of the button’s border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/bezelStyle-swift.property
func (b_ Button) SetBezelStyle(value BezelStyle) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBezelStyle:"), value)
}/* debug [instance_properties/setter]: bezelStyle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/borderShape
func (b_ Button) BorderShape() ControlBorderShape {
	rv := objc.Send[ControlBorderShape](b_.ID, objc.Sel("borderShape"))
	return rv
}/* debug [instance_properties/getter]: borderShape */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/borderShape
func (b_ Button) SetBorderShape(value ControlBorderShape) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBorderShape:"), value)
}/* debug [instance_properties/setter]: borderShape */


// A tint color to use for the template image and text content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/contentTintColor
func (b_ Button) ContentTintColor() IColor {
	rv := objc.Send[Color](b_.ID, objc.Sel("contentTintColor"))
	return rv
}/* debug [instance_properties/getter]: contentTintColor */


// A tint color to use for the template image and text content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/contentTintColor
func (b_ Button) SetContentTintColor(value IColor) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setContentTintColor:"), value)
}/* debug [instance_properties/setter]: contentTintColor */


// A Boolean value that defines whether a button’s action has a destructive effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/hasDestructiveAction
func (b_ Button) HasDestructiveAction() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("hasDestructiveAction"))
	return rv
}/* debug [instance_properties/getter]: hasDestructiveAction */


// A Boolean value that defines whether a button’s action has a destructive effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/hasDestructiveAction
func (b_ Button) SetHasDestructiveAction(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setHasDestructiveAction:"), value)
}/* debug [instance_properties/setter]: hasDestructiveAction */


// The image that appears on the button when it’s in an off state, or if there is no such image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/image
func (b_ Button) Image() IImage {
	rv := objc.Send[Image](b_.ID, objc.Sel("image"))
	return rv
}/* debug [instance_properties/getter]: image */


// The image that appears on the button when it’s in an off state, or if there is no such image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/image
func (b_ Button) SetImage(value IImage) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setImage:"), value)
}/* debug [instance_properties/setter]: image */


// A Boolean value that determines how the button’s image and title are positioned together within the button bezel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/imageHugsTitle
func (b_ Button) ImageHugsTitle() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("imageHugsTitle"))
	return rv
}/* debug [instance_properties/getter]: imageHugsTitle */


// A Boolean value that determines how the button’s image and title are positioned together within the button bezel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/imageHugsTitle
func (b_ Button) SetImageHugsTitle(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setImageHugsTitle:"), value)
}/* debug [instance_properties/setter]: imageHugsTitle */


// The position of the button’s image relative to its title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/imagePosition
func (b_ Button) ImagePosition() CellImagePosition {
	rv := objc.Send[CellImagePosition](b_.ID, objc.Sel("imagePosition"))
	return rv
}/* debug [instance_properties/getter]: imagePosition */


// The position of the button’s image relative to its title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/imagePosition
func (b_ Button) SetImagePosition(value CellImagePosition) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setImagePosition:"), value)
}/* debug [instance_properties/setter]: imagePosition */


// The scaling mode applied to make the cell’s image fit the frame of the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/imageScaling
func (b_ Button) ImageScaling() ImageScaling {
	rv := objc.Send[ImageScaling](b_.ID, objc.Sel("imageScaling"))
	return rv
}/* debug [instance_properties/getter]: imageScaling */


// The scaling mode applied to make the cell’s image fit the frame of the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/imageScaling
func (b_ Button) SetImageScaling(value ImageScaling) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setImageScaling:"), value)
}/* debug [instance_properties/setter]: imageScaling */


// A Boolean value that determines whether the button has a border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/isBordered
func (b_ Button) Bordered() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("bordered"))
	return rv
}/* debug [instance_properties/getter]: bordered */


// A Boolean value that determines whether the button has a border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/isBordered
func (b_ Button) SetBordered(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBordered:"), value)
}/* debug [instance_properties/setter]: bordered */


// A Boolean value that indicates whether spring loading is enabled for the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/isSpringLoaded
func (b_ Button) SpringLoaded() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("springLoaded"))
	return rv
}/* debug [instance_properties/getter]: springLoaded */


// A Boolean value that indicates whether spring loading is enabled for the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/isSpringLoaded
func (b_ Button) SetSpringLoaded(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSpringLoaded:"), value)
}/* debug [instance_properties/setter]: springLoaded */


// A Boolean value that indicates whether the button is transparent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/isTransparent
func (b_ Button) Transparent() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("transparent"))
	return rv
}/* debug [instance_properties/getter]: transparent */


// A Boolean value that indicates whether the button is transparent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/isTransparent
func (b_ Button) SetTransparent(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTransparent:"), value)
}/* debug [instance_properties/setter]: transparent */


// The key-equivalent character of the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/keyEquivalent
func (b_ Button) KeyEquivalent() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](b_.ID, objc.Sel("keyEquivalent"))
	return rv
}/* debug [instance_properties/getter]: keyEquivalent */


// The key-equivalent character of the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/keyEquivalent
func (b_ Button) SetKeyEquivalent(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setKeyEquivalent:"), value)
}/* debug [instance_properties/setter]: keyEquivalent */


// The mask specifying the modifier keys for the button’s key equivalent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/keyEquivalentModifierMask
func (b_ Button) KeyEquivalentModifierMask() EventModifierFlags {
	rv := objc.Send[EventModifierFlags](b_.ID, objc.Sel("keyEquivalentModifierMask"))
	return rv
}/* debug [instance_properties/getter]: keyEquivalentModifierMask */


// The mask specifying the modifier keys for the button’s key equivalent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/keyEquivalentModifierMask
func (b_ Button) SetKeyEquivalentModifierMask(value EventModifierFlags) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setKeyEquivalentModifierMask:"), value)
}/* debug [instance_properties/setter]: keyEquivalentModifierMask */


// An integer value indicating the maximum pressure level for a button of type .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/maxAcceleratorLevel
func (b_ Button) MaxAcceleratorLevel() int {
	rv := objc.Send[int](b_.ID, objc.Sel("maxAcceleratorLevel"))
	return rv
}/* debug [instance_properties/getter]: maxAcceleratorLevel */


// An integer value indicating the maximum pressure level for a button of type .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/maxAcceleratorLevel
func (b_ Button) SetMaxAcceleratorLevel(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setMaxAcceleratorLevel:"), value)
}/* debug [instance_properties/setter]: maxAcceleratorLevel */


// A Boolean value that determines whether the button displays its border only when the pointer is over it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/showsBorderOnlyWhileMouseInside
func (b_ Button) ShowsBorderOnlyWhileMouseInside() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("showsBorderOnlyWhileMouseInside"))
	return rv
}/* debug [instance_properties/getter]: showsBorderOnlyWhileMouseInside */


// A Boolean value that determines whether the button displays its border only when the pointer is over it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/showsBorderOnlyWhileMouseInside
func (b_ Button) SetShowsBorderOnlyWhileMouseInside(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setShowsBorderOnlyWhileMouseInside:"), value)
}/* debug [instance_properties/setter]: showsBorderOnlyWhileMouseInside */


// The sound that plays when the user clicks the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/sound
func (b_ Button) Sound() ISound {
	rv := objc.Send[Sound](b_.ID, objc.Sel("sound"))
	return rv
}/* debug [instance_properties/getter]: sound */


// The sound that plays when the user clicks the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/sound
func (b_ Button) SetSound(value ISound) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSound:"), value)
}/* debug [instance_properties/setter]: sound */


// The button’s state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/state
func (b_ Button) State() ControlStateValue /* typedef */ {
	rv := objc.Send[int](b_.ID, objc.Sel("state"))
	return rv
}/* debug [instance_properties/getter]: state */


// The button’s state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/state
func (b_ Button) SetState(value ControlStateValue /* typedef */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setState:"), value)
}/* debug [instance_properties/setter]: state */


// The combination of point size, weight, and scale to use when sizing and displaying symbol images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/symbolConfiguration
func (b_ Button) SymbolConfiguration() IImageSymbolConfiguration {
	rv := objc.Send[ImageSymbolConfiguration](b_.ID, objc.Sel("symbolConfiguration"))
	return rv
}/* debug [instance_properties/getter]: symbolConfiguration */


// The combination of point size, weight, and scale to use when sizing and displaying symbol images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/symbolConfiguration
func (b_ Button) SetSymbolConfiguration(value IImageSymbolConfiguration) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSymbolConfiguration:"), value)
}/* debug [instance_properties/setter]: symbolConfiguration */


// The tint prominence of the button. Use tint prominence to gently suggest a hierarchy when multiple buttons perform similar actions. A button with primary tint prominence suggests the most preferred option, while secondary prominence indicates a reasonable alternative. See for a list of possible values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/tintProminence
func (b_ Button) TintProminence() TintProminence {
	rv := objc.Send[TintProminence](b_.ID, objc.Sel("tintProminence"))
	return rv
}/* debug [instance_properties/getter]: tintProminence */


// The tint prominence of the button. Use tint prominence to gently suggest a hierarchy when multiple buttons perform similar actions. A button with primary tint prominence suggests the most preferred option, while secondary prominence indicates a reasonable alternative. See for a list of possible values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/tintProminence
func (b_ Button) SetTintProminence(value TintProminence) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTintProminence:"), value)
}/* debug [instance_properties/setter]: tintProminence */


// The title displayed on the button when it’s in an off state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/title
func (b_ Button) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](b_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// The title displayed on the button when it’s in an off state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/title
func (b_ Button) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTitle:"), value)
}/* debug [instance_properties/setter]: title */


// A Boolean value that determines whether the button has a border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/isbordered
func (b_ Button) IsBordered() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isBordered"))
	return rv
}/* debug [instance_properties/getter]: isBordered */


// A Boolean value that determines whether the button has a border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/isbordered
func (b_ Button) SetIsBordered(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsBordered:"), value)
}/* debug [instance_properties/setter]: isBordered */


// A Boolean value that indicates whether spring loading is enabled for the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/isspringloaded
func (b_ Button) IsSpringLoaded() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isSpringLoaded"))
	return rv
}/* debug [instance_properties/getter]: isSpringLoaded */


// A Boolean value that indicates whether spring loading is enabled for the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/isspringloaded
func (b_ Button) SetIsSpringLoaded(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsSpringLoaded:"), value)
}/* debug [instance_properties/setter]: isSpringLoaded */


// A Boolean value that indicates whether the button is transparent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/istransparent
func (b_ Button) IsTransparent() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isTransparent"))
	return rv
}/* debug [instance_properties/getter]: isTransparent */


// A Boolean value that indicates whether the button is transparent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/istransparent
func (b_ Button) SetIsTransparent(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsTransparent:"), value)
}/* debug [instance_properties/setter]: isTransparent */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSButton */


