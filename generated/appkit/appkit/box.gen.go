// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Box] class.
var BoxClass objc.Class

func init() {
	BoxClass = objc.GetClass("NSBox")
}

type Box struct {
	objc.ID
}

func BoxFrom(ptr unsafe.Pointer) Box {
	return Box{
		ID: objc.ID(ptr),
	}
}


// Places the receiver so its content view lies on the specified frame. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSBox/setFrameFromContentFrame(_:)
func (b_ Box) SetFrameFromContentFrame(contentFrame unsafe.Pointer) {
	sel := objc.RegisterName("setFrameFromContentFrame:")
	b_.ID.Send(sel, contentFrame)
}
// Sets the title of the receiver with a character denoted as an access key. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSBox/setTitleWithMnemonic:
func (b_ Box) SetTitleWithMnemonic(stringWithAmpersand string) {
	sel := objc.RegisterName("setTitleWithMnemonic:")
	b_.ID.Send(sel, stringWithAmpersand)
}
// Resizes and moves the receiver’s content view so it just encloses its subviews. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSBox/sizeToFit()
func (b_ Box) SizeToFit() {
	sel := objc.RegisterName("sizeToFit")
	b_.ID.Send(sel)
}


