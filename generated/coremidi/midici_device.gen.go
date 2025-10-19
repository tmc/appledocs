// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MIDICIDevice] class.
var (
	mIDICIDeviceClass     _MIDICIDeviceClass
	mIDICIDeviceClassOnce sync.Once
)

func getMIDICIDeviceClass() _MIDICIDeviceClass {
	mIDICIDeviceClassOnce.Do(func() {
		mIDICIDeviceClass = _MIDICIDeviceClass{objc.GetClass("MIDICIDevice")}
	})
	return mIDICIDeviceClass
}

type _MIDICIDeviceClass struct {
	class objc.Class
}

// An interface definition for the [MIDICIDevice] class.
type IMIDICIDevice interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDevice
type MIDICIDevice struct {
	objectivec.Object
}

// MIDICIDeviceFrom constructs a [MIDICIDevice] from an unsafe.Pointer.
func MIDICIDeviceFrom(ptr unsafe.Pointer) MIDICIDevice {
	return MIDICIDevice{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MIDICIDeviceClass) Alloc() MIDICIDevice {
	rv := objc.Send[MIDICIDevice](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MIDICIDeviceClass) New() MIDICIDevice {
	rv := objc.Send[MIDICIDevice](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDICIDevice) Init() MIDICIDevice {
	rv := objc.Send[MIDICIDevice](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDICIDevice) Autorelease() MIDICIDevice {
	rv := objc.Send[MIDICIDevice](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDICIDevice creates a new MIDICIDevice instance.
func NewMIDICIDevice() MIDICIDevice {
	return getMIDICIDeviceClass().New()
}




