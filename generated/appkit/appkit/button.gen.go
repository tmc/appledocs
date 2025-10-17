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


// Creates a standard checkbox with the title you specify. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSButton/init(checkboxWithTitle:target:action:)
func (bc Button) CheckboxWithTitleTargetAction(title string, target objc.ID, action objc.SEL) unsafe.Pointer {
	sel := objc.RegisterName("checkboxWithTitle:target:action:")
	ret := objc.ID(ButtonClass).Send(sel, title, target, action)
	return unsafe.Pointer(ret)
}
// Creates a standard push button with a title and image. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSButton/init(title:image:target:action:)
func (bc Button) ButtonWithTitleImageTargetAction(title string, image unsafe.Pointer, target objc.ID, action objc.SEL) unsafe.Pointer {
	sel := objc.RegisterName("buttonWithTitle:image:target:action:")
	ret := objc.ID(ButtonClass).Send(sel, title, image, target, action)
	return unsafe.Pointer(ret)
}
// Creates a standard push button with the title you specify. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSButton/init(title:target:action:)
func (bc Button) ButtonWithTitleTargetAction(title string, target objc.ID, action objc.SEL) unsafe.Pointer {
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


