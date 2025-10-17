// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Shadow] class.
var ShadowClass objc.Class

func init() {
	ShadowClass = objc.GetClass("NSShadow")
}

type Shadow struct {
	objc.ID
}

func ShadowFrom(ptr unsafe.Pointer) Shadow {
	return Shadow{
		ID: objc.ID(ptr),
	}
}




