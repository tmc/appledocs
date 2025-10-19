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

// A control that defines an area on the screen that a user clicks to trigger an action. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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


// Creates a standard checkbox with the title you specify. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/init(checkboxWithTitle:target:action:)
func NewButtonCheckboxWithTitleTargetAction(title string, target objc.ID, action objc.SEL) Button {
	rv := objc.Send[Button](objc.ID(getButtonClass().class), objc.Sel("checkboxWithTitle:target:action:"), objc.String(title), target, action)
	rv.Autorelease()
	return rv
}
// Creates a standard push button with a title and image. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/init(title:image:target:action:)
func NewButtonWithTitleImageTargetAction(title string, image unsafe.Pointer, target objc.ID, action objc.SEL) Button {
	rv := objc.Send[Button](objc.ID(getButtonClass().class), objc.Sel("buttonWithTitle:image:target:action:"), objc.String(title), image, target, action)
	rv.Autorelease()
	return rv
}
// Creates a standard push button with the title you specify. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/init(title:target:action:)
func NewButtonWithTitleTargetAction(title string, target objc.ID, action objc.SEL) Button {
	rv := objc.Send[Button](objc.ID(getButtonClass().class), objc.Sel("buttonWithTitle:target:action:"), objc.String(title), target, action)
	rv.Autorelease()
	return rv
}


// Creates a standard checkbox with the title you specify. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/init(checkboxWithTitle:target:action:)
func (bc _ButtonClass) CheckboxWithTitleTargetAction(title string, target objc.ID, action objc.SEL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("checkboxWithTitle:target:action:"), objc.String(title), target, action)
	return rv
}
// Creates a standard push button with a title and image. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/init(title:image:target:action:)
func (bc _ButtonClass) ButtonWithTitleImageTargetAction(title string, image unsafe.Pointer, target objc.ID, action objc.SEL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("buttonWithTitle:image:target:action:"), objc.String(title), image, target, action)
	return rv
}
// Creates a standard push button with the title you specify. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/init(title:target:action:)
func (bc _ButtonClass) ButtonWithTitleTargetAction(title string, target objc.ID, action objc.SEL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("buttonWithTitle:target:action:"), objc.String(title), target, action)
	return rv
}
// Sets the priority compression options for this button. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/compress(withPrioritizedCompressionOptions:)
func (b_ Button) CompressWithPrioritizedCompressionOptions(prioritizedOptions unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("compressWithPrioritizedCompressionOptions:"), prioritizedOptions)
}
// Sets the message delay and interval periods for a continuous button. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/setPeriodicDelay(_:interval:)
func (b_ Button) SetPeriodicDelayInterval(delay float32, interval float32) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPeriodicDelay:interval:"), delay, interval)
}

