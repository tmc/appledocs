// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ScrubberProportionalLayout] class.
var ScrubberProportionalLayoutClass objc.Class

func init() {
	ScrubberProportionalLayoutClass = objc.GetClass("NSScrubberProportionalLayout")
}

type ScrubberProportionalLayout struct {
	objc.ID
}

func ScrubberProportionalLayoutFrom(ptr unsafe.Pointer) ScrubberProportionalLayout {
	return ScrubberProportionalLayout{
		ID: objc.ID(ptr),
	}
}




