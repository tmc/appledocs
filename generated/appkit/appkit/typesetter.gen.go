// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Typesetter] class.
var TypesetterClass objc.Class

func init() {
	TypesetterClass = objc.GetClass("NSTypesetter")
}

type Typesetter struct {
	objc.ID
}

func TypesetterFrom(ptr unsafe.Pointer) Typesetter {
	return Typesetter{
		ID: objc.ID(ptr),
	}
}




