// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ScrubberLayoutAttributes] class.
var ScrubberLayoutAttributesClass objc.Class

func init() {
	ScrubberLayoutAttributesClass = objc.GetClass("NSScrubberLayoutAttributes")
}

type ScrubberLayoutAttributes struct {
	objc.ID
}

func ScrubberLayoutAttributesFrom(ptr unsafe.Pointer) ScrubberLayoutAttributes {
	return ScrubberLayoutAttributes{
		ID: objc.ID(ptr),
	}
}



