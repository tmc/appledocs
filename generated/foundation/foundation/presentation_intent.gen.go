// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PresentationIntent] class.
var PresentationIntentClass objc.Class

func init() {
	PresentationIntentClass = objc.GetClass("NSPresentationIntent")
}

type PresentationIntent struct {
	objc.ID
}

func PresentationIntentFrom(ptr unsafe.Pointer) PresentationIntent {
	return PresentationIntent{
		ID: objc.ID(ptr),
	}
}




