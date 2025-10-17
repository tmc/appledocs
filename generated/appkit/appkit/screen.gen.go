// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Screen] class.
var ScreenClass objc.Class

func init() {
	ScreenClass = objc.GetClass("NSScreen")
}

type Screen struct {
	objc.ID
}

func ScreenFrom(ptr unsafe.Pointer) Screen {
	return Screen{
		ID: objc.ID(ptr),
	}
}


// A Boolean value indicating whether the color space of the screen is capable of representing the specified display gamut. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScreen/canRepresent(_:)
func (s_ Screen) CanRepresentDisplayGamut(displayGamut unsafe.Pointer) bool {
	sel := objc.RegisterName("canRepresentDisplayGamut:")
	ret := s_.ID.Send(sel, displayGamut)
	return ret != 0
}
// Returns the scaling factor from user space to device space on the screen. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScreen/userSpaceScaleFactor
func (s_ Screen) UserSpaceScaleFactor() float64 {
	sel := objc.RegisterName("userSpaceScaleFactor")
	ret := s_.ID.Send(sel)
	return float64(ret)
}

