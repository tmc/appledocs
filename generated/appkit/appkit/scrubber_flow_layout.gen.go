// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ScrubberFlowLayout] class.
var ScrubberFlowLayoutClass objc.Class

func init() {
	ScrubberFlowLayoutClass = objc.GetClass("NSScrubberFlowLayout")
}

type ScrubberFlowLayout struct {
	objc.ID
}

func ScrubberFlowLayoutFrom(ptr unsafe.Pointer) ScrubberFlowLayout {
	return ScrubberFlowLayout{
		ID: objc.ID(ptr),
	}
}




