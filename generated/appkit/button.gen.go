// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Button] class.
var (
	buttonClass     _ButtonClass
	buttonClassOnce sync.Once
)

func getButtonClass() _ButtonClass {
	buttonClassOnce.Do(func() {
		buttonClass = _ButtonClass{objc.GetClass("NSButton")}
	})
	return buttonClass
}

type _ButtonClass struct {
	class objc.Class
}

// An interface definition for the [Button] class.
type IButton interface {
	IControl
	CompressWithPrioritizedCompressionOptions(prioritizedOptions unsafe.Pointer)
	SetPeriodicDelayInterval(delay float32, interval float32)
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
func NewButtonCheckboxWithTitleTargetAction(title string, target objc.ID, action objc.SEL) Button {
	rv := objc.Send[Button](objc.ID(getButtonClass().class), objc.Sel("checkboxWithTitle:target:action:"), objc.String(title), target, action)
	return rv
}

// Creates a standard push button with a title and image.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/init(title:image:target:action:)
func NewButtonWithTitleImageTargetAction(title string, image unsafe.Pointer, target objc.ID, action objc.SEL) Button {
	rv := objc.Send[Button](objc.ID(getButtonClass().class), objc.Sel("buttonWithTitle:image:target:action:"), objc.String(title), image, target, action)
	return rv
}

// Creates a standard push button with the title you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/init(title:target:action:)
func NewButtonWithTitleTargetAction(title string, target objc.ID, action objc.SEL) Button {
	rv := objc.Send[Button](objc.ID(getButtonClass().class), objc.Sel("buttonWithTitle:target:action:"), objc.String(title), target, action)
	return rv
}


// Creates a standard checkbox with the title you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/init(checkboxWithTitle:target:action:)
func (bc _ButtonClass) CheckboxWithTitleTargetAction(title string, target objc.ID, action objc.SEL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("checkboxWithTitle:target:action:"), objc.String(title), target, action)
	return rv
}

// Creates a standard push button with a title and image.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/init(title:image:target:action:)
func (bc _ButtonClass) ButtonWithTitleImageTargetAction(title string, image unsafe.Pointer, target objc.ID, action objc.SEL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("buttonWithTitle:image:target:action:"), objc.String(title), image, target, action)
	return rv
}

// Creates a standard push button with the title you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/init(title:target:action:)
func (bc _ButtonClass) ButtonWithTitleTargetAction(title string, target objc.ID, action objc.SEL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("buttonWithTitle:target:action:"), objc.String(title), target, action)
	return rv
}

// Sets the priority compression options for this button.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/compress(withPrioritizedCompressionOptions:)
func (b_ Button) CompressWithPrioritizedCompressionOptions(prioritizedOptions unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("compressWithPrioritizedCompressionOptions:"), prioritizedOptions)
}

// Sets the message delay and interval periods for a continuous button.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/setPeriodicDelay(_:interval:)
func (b_ Button) SetPeriodicDelayInterval(delay float32, interval float32) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPeriodicDelay:interval:"), delay, interval)
}

// The compression options active for this button.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/activeCompressionOptions
func (b_ Button) ActiveCompressionOptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("activeCompressionOptions"))
	return rv
}
// The title that the button displays in an off state, as an attributed string.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/attributedTitle
func (b_ Button) AttributedTitle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("attributedTitle"))
	return rv
}

// SetAttributedTitle sets the value of the attributedTitle property.
// The title that the button displays in an off state, as an attributed string.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/attributedTitle
func (b_ Button) SetAttributedTitle(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAttributedTitle:"), value)
}
// The color of the button’s bezel, in appearances that support it.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/bezelColor
func (b_ Button) BezelColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("bezelColor"))
	return rv
}

// SetBezelColor sets the value of the bezelColor property.
// The color of the button’s bezel, in appearances that support it.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/bezelColor
func (b_ Button) SetBezelColor(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBezelColor:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/borderShape
func (b_ Button) BorderShape() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("borderShape"))
	return rv
}

// SetBorderShape sets the value of the borderShape property.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/borderShape
func (b_ Button) SetBorderShape(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBorderShape:"), value)
}
// A tint color to use for the template image and text content.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/contentTintColor
func (b_ Button) ContentTintColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("contentTintColor"))
	return rv
}

// SetContentTintColor sets the value of the contentTintColor property.
// A tint color to use for the template image and text content.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/contentTintColor
func (b_ Button) SetContentTintColor(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setContentTintColor:"), value)
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
// The sound that plays when the user clicks the button.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/sound
func (b_ Button) Sound() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("sound"))
	return rv
}

// SetSound sets the value of the sound property.
// The sound that plays when the user clicks the button.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/sound
func (b_ Button) SetSound(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSound:"), value)
}

