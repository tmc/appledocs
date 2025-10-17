// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ATSTypesetter] class.
var ATSTypesetterClass objc.Class

func init() {
	ATSTypesetterClass = objc.GetClass("NSATSTypesetter")
}

type ATSTypesetter struct {
	objc.ID
}

func ATSTypesetterFrom(ptr unsafe.Pointer) ATSTypesetter {
	return ATSTypesetter{
		ID: objc.ID(ptr),
	}
}



