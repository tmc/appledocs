// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var presentationIntentClass _PresentationIntentClass

func init() {
	presentationIntentClass = _PresentationIntentClass{objc.GetClass("NSPresentationIntent")}
}

type _PresentationIntentClass struct {
	class objc.Class
}

type PresentationIntent struct {
	objc.ID
}

func PresentationIntentFrom(ptr unsafe.Pointer) PresentationIntent {
	return PresentationIntent{
		ID: objc.ID(ptr),
	}
}




