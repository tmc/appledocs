// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Sound] class.
var soundClass = _SoundClass{objc.GetClass("NSSound")}

type _SoundClass struct {
	class objc.Class
}

// An interface definition for the [Sound] class.
type ISound interface {
	objectivec.IObject
}

// A simple interface for loading and playing audio files. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound

type Sound struct {
	objectivec.Object
}

// SoundFrom constructs a [Sound] from an unsafe.Pointer.
//
// A simple interface for loading and playing audio files.
func SoundFrom(ptr unsafe.Pointer) Sound {
	return Sound{objectivec.Object{objc.ID(ptr)}}
}



