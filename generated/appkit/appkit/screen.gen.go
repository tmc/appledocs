// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Screen] class.
var (
	screenClass     _ScreenClass
	screenClassOnce sync.Once
)

func getScreenClass() _ScreenClass {
	screenClassOnce.Do(func() {
		screenClass = _ScreenClass{objc.GetClass("NSScreen")}
	})
	return screenClass
}

type _ScreenClass struct {
	class objc.Class
}

// An interface definition for the [Screen] class.
type IScreen interface {
	objectivec.IObject
	CanRepresentDisplayGamut(displayGamut unsafe.Pointer) bool
	UserSpaceScaleFactor() float64
}

// An object that describes the attributes of a computer’s monitor or screen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen

type Screen struct {
	objectivec.Object
}

// ScreenFrom constructs a [Screen] from an unsafe.Pointer.
//
// An object that describes the attributes of a computer’s monitor or screen.
func ScreenFrom(ptr unsafe.Pointer) Screen {
	return Screen{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (sc _ScreenClass) Alloc() Screen {
	rv := objc.Send[Screen](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (sc _ScreenClass) New() Screen {
	rv := objc.Send[Screen](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ Screen) Init() Screen {
	rv := objc.Send[Screen](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ Screen) Autorelease() Screen {
	rv := objc.Send[Screen](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScreen creates a new Screen instance.
func NewScreen() Screen {
	return getScreenClass().New()
}


// A Boolean value indicating whether the color space of the screen is capable of representing the specified display gamut. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/canRepresent(_:)
func (s_ Screen) CanRepresentDisplayGamut(displayGamut unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("canRepresentDisplayGamut:"), displayGamut)
	return rv
}
// Returns the scaling factor from user space to device space on the screen. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/userSpaceScaleFactor
func (s_ Screen) UserSpaceScaleFactor() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("userSpaceScaleFactor"))
	return rv
}


