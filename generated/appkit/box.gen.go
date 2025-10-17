// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Box] class.
var boxClass = _BoxClass{objc.GetClass("NSBox")}

type _BoxClass struct {
	class objc.Class
}

// A stylized rectangular box with an optional title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox

type Box struct {
	View
}

// BoxFrom constructs a [Box] from an unsafe.Pointer.
//
// A stylized rectangular box with an optional title.
func BoxFrom(ptr unsafe.Pointer) Box {
	return Box{
		View: ViewFrom(ptr),
	}
}

// Places the receiver so its content view lies on the specified frame. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/setFrameFromContentFrame(_:)
func (b_ Box) SetFrameFromContentFrame(contentFrame unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setFrameFromContentFrame:"), contentFrame)
}
// Sets the title of the receiver with a character denoted as an access key. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/setTitleWithMnemonic:
func (b_ Box) SetTitleWithMnemonic(stringWithAmpersand string) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTitleWithMnemonic:"), stringWithAmpersand)
}
// Resizes and moves the receiver’s content view so it just encloses its subviews. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/sizeToFit()
func (b_ Box) SizeToFit() {
	objc.Send[objc.ID](b_.ID, objc.Sel("sizeToFit"))
}


