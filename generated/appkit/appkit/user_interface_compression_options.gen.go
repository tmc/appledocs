// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UserInterfaceCompressionOptions] class.
var UserInterfaceCompressionOptionsClass objc.Class

func init() {
	UserInterfaceCompressionOptionsClass = objc.GetClass("NSUserInterfaceCompressionOptions")
}

type UserInterfaceCompressionOptions struct {
	objc.ID
}

func UserInterfaceCompressionOptionsFrom(ptr unsafe.Pointer) UserInterfaceCompressionOptions {
	return UserInterfaceCompressionOptions{
		ID: objc.ID(ptr),
	}
}




