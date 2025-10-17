
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Screen] class.
var ScreenClass _ScreenClass

func init() {
	ScreenClass = _ScreenClass{objc.GetClass("NSScreen")}
}

type _ScreenClass struct {
	objc.Class
}

// An interface definition for the [Screen] class.
type IScreen interface {
	ID() objc.ID
	CanRepresentDisplayGamut(displayGamut unsafe.Pointer) bool
	UserSpaceScaleFactor() float64
}

type Screen struct {
	id objc.ID
}

func ScreenFrom(ptr unsafe.Pointer) Screen {
	return Screen{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ Screen) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _ScreenClass) Alloc() Screen {
	rv := objc.Send[Screen](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _ScreenClass) New() Screen {
	rv := objc.Send[Screen](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewScreen creates and returns a new initialized instance.
func NewScreen() Screen {
	return ScreenClass.New()
}

// Init initializes the instance.
func (s_ Screen) Init() Screen {
	rv := objc.Send[Screen](s_.ID(), selInit)
	return rv
}
// A Boolean value indicating whether the color space of the screen is capable of representing the specified display gamut. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScreen/canRepresent(_:)
func (s_ Screen) CanRepresentDisplayGamut(displayGamut unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID(), objc.RegisterName("canRepresentDisplayGamut:"), displayGamut)
	return rv
}
// Returns the scaling factor from user space to device space on the screen. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScreen/userSpaceScaleFactor
func (s_ Screen) UserSpaceScaleFactor() float64 {
	rv := objc.Send[float64](s_.ID(), objc.RegisterName("userSpaceScaleFactor"))
	return rv
}
// The color space of the screen. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScreen/colorSpace
func (s_ Screen) ColorSpace() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("colorSpace"))
	return rv
}
// The current bit depth and colorspace information of the screen. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScreen/depth
func (s_ Screen) Depth() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("depth"))
	return rv
}
// The device dictionary for the screen. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScreen/deviceDescription
func (s_ Screen) DeviceDescription() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("deviceDescription"))
	return rv
}
// The dimensions and location of the screen. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScreen/frame
func (s_ Screen) Frame() foundation.Rect {
	rv := objc.Send[foundation.Rect](s_.ID(), objc.RegisterName("frame"))
	return rv
}
// The localized name of the display. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScreen/localizedName
func (s_ Screen) LocalizedName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("localizedName"))
	return rv
}
// The maximum possible color component value for the screen when it’s in extended dynamic range (EDR) mode. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScreen/maximumPotentialExtendedDynamicRangeColorComponentValue
func (s_ Screen) MaximumPotentialExtendedDynamicRangeColorComponentValue() float64 {
	rv := objc.Send[float64](s_.ID(), objc.RegisterName("maximumPotentialExtendedDynamicRangeColorComponentValue"))
	return rv
}
// A zero-terminated array of the window depths supported by the screen. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScreen/supportedWindowDepths
func (s_ Screen) SupportedWindowDepths() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("supportedWindowDepths"))
	return rv
}
