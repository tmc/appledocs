// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Button] class.
var ButtonClass objc.Class

func init() {
	ButtonClass = objc.GetClass("NSButton")
}

type Button struct {
	objc.ID
}

func ButtonFrom(ptr unsafe.Pointer) Button {
	return Button{
		ID: objc.ID(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (bc Button) Alloc() Button {
	ret := objc.ID(ButtonClass).Send(objc.RegisterName("alloc"))
	return Button{ret}
}

// New creates and returns a new initialized instance.
func (bc Button) New() Button {
	ret := objc.ID(ButtonClass).Send(objc.RegisterName("new"))
	return Button{ret}
}

// NewButton creates and returns a new initialized instance.
func NewButton() Button {
	ret := objc.ID(ButtonClass).Send(objc.RegisterName("new"))
	return Button{ret}
}

// Init initializes the instance.
func (b_ Button) Init() Button {
	ret := b_.ID.Send(objc.RegisterName("init"))
	return Button{ret}
}
// Creates a standard checkbox with the title you specify. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSButton/init(checkboxWithTitle:target:action:)
func (bc Button) CheckboxWithTitleTargetAction(title unsafe.Pointer, target objc.ID, action objc.SEL) unsafe.Pointer {
	sel := objc.RegisterName("checkboxWithTitle:target:action:")
	ret := objc.ID(ButtonClass).Send(sel, title, target, action)
	return unsafe.Pointer(ret)
}
// Creates a standard push button with a title and image. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSButton/init(title:image:target:action:)
func (bc Button) ButtonWithTitleImageTargetAction(title unsafe.Pointer, image unsafe.Pointer, target objc.ID, action objc.SEL) unsafe.Pointer {
	sel := objc.RegisterName("buttonWithTitle:image:target:action:")
	ret := objc.ID(ButtonClass).Send(sel, title, image, target, action)
	return unsafe.Pointer(ret)
}
// Creates a standard push button with the title you specify. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSButton/init(title:target:action:)
func (bc Button) ButtonWithTitleTargetAction(title unsafe.Pointer, target objc.ID, action objc.SEL) unsafe.Pointer {
	sel := objc.RegisterName("buttonWithTitle:target:action:")
	ret := objc.ID(ButtonClass).Send(sel, title, target, action)
	return unsafe.Pointer(ret)
}
// Sets the priority compression options for this button. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSButton/compress(withPrioritizedCompressionOptions:)
func (b_ Button) CompressWithPrioritizedCompressionOptions(prioritizedOptions unsafe.Pointer) {
	sel := objc.RegisterName("compressWithPrioritizedCompressionOptions:")
	b_.ID.Send(sel, prioritizedOptions)
}
// Sets the message delay and interval periods for a continuous button. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSButton/setPeriodicDelay(_:interval:)
func (b_ Button) SetPeriodicDelayInterval(delay float32, interval float32) {
	sel := objc.RegisterName("setPeriodicDelay:interval:")
	b_.ID.Send(sel, delay, interval)
}

