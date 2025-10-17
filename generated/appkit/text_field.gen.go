// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TextField] class.
var textFieldClass = _TextFieldClass{objc.GetClass("NSTextField")}

type _TextFieldClass struct {
	class objc.Class
}

// Text the user can select or edit to send an action message to a target when the user presses the Return key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField

type TextField struct {
	Control
}

// TextFieldFrom constructs a [TextField] from an unsafe.Pointer.
//
// Text the user can select or edit to send an action message to a target when the user presses the Return key.
func TextFieldFrom(ptr unsafe.Pointer) TextField {
	return TextField{
		Control: ControlFrom(ptr),
	}
}



