// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [AppleEventDescriptor] class.
var AppleEventDescriptorClass objc.Class

func init() {
	AppleEventDescriptorClass = objc.GetClass("NSAppleEventDescriptor")
}

type AppleEventDescriptor struct {
	objc.ID
}

func AppleEventDescriptorFrom(ptr unsafe.Pointer) AppleEventDescriptor {
	return AppleEventDescriptor{
		ID: objc.ID(ptr),
	}
}



