// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Sound] class.
var SoundClass objc.Class

func init() {
	SoundClass = objc.GetClass("NSSound")
}

type Sound struct {
	objc.ID
}

func SoundFrom(ptr unsafe.Pointer) Sound {
	return Sound{
		ID: objc.ID(ptr),
	}
}



