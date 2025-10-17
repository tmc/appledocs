// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [UserInterfaceCompressionOptions] class.
var userInterfaceCompressionOptionsClass = _UserInterfaceCompressionOptionsClass{objc.GetClass("NSUserInterfaceCompressionOptions")}

type _UserInterfaceCompressionOptionsClass struct {
	class objc.Class
}

// An object that specifies how user interface elements resize themselves when space is constrained. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions

type UserInterfaceCompressionOptions struct {
	objectivec.Object
}

// UserInterfaceCompressionOptionsFrom constructs a [UserInterfaceCompressionOptions] from an unsafe.Pointer.
//
// An object that specifies how user interface elements resize themselves when space is constrained.
func UserInterfaceCompressionOptionsFrom(ptr unsafe.Pointer) UserInterfaceCompressionOptions {
	return UserInterfaceCompressionOptions{objectivec.Object{objc.ID(ptr)}}
}



