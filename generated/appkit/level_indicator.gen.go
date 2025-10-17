// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [LevelIndicator] class.
var levelIndicatorClass = _LevelIndicatorClass{objc.GetClass("NSLevelIndicator")}

type _LevelIndicatorClass struct {
	class objc.Class
}

// A visual representation of a level or quantity, using discrete values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator

type LevelIndicator struct {
	Control
}

// LevelIndicatorFrom constructs a [LevelIndicator] from an unsafe.Pointer.
//
// A visual representation of a level or quantity, using discrete values.
func LevelIndicatorFrom(ptr unsafe.Pointer) LevelIndicator {
	return LevelIndicator{
		Control: ControlFrom(ptr),
	}
}



