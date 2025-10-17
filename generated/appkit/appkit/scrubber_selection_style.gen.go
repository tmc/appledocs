// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ScrubberSelectionStyle] class.
var ScrubberSelectionStyleClass objc.Class

func init() {
	ScrubberSelectionStyleClass = objc.GetClass("NSScrubberSelectionStyle")
}

type ScrubberSelectionStyle struct {
	objc.ID
}

func ScrubberSelectionStyleFrom(ptr unsafe.Pointer) ScrubberSelectionStyle {
	return ScrubberSelectionStyle{
		ID: objc.ID(ptr),
	}
}




