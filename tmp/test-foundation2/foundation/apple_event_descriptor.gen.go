// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var appleEventDescriptorClass _AppleEventDescriptorClass

func init() {
	appleEventDescriptorClass = _AppleEventDescriptorClass{objc.GetClass("NSAppleEventDescriptor")}
}

type _AppleEventDescriptorClass struct {
	class objc.Class
}

type AppleEventDescriptor struct {
	objc.ID
}

func AppleEventDescriptorFrom(ptr unsafe.Pointer) AppleEventDescriptor {
	return AppleEventDescriptor{
		ID: objc.ID(ptr),
	}
}




