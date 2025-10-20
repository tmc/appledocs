// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MIDICIDeviceManager] class.
var (
	MIDICIDeviceManagerClass     _MIDICIDeviceManagerClass
	MIDICIDeviceManagerClassOnce sync.Once
)

func getMIDICIDeviceManagerClass() _MIDICIDeviceManagerClass {
	MIDICIDeviceManagerClassOnce.Do(func() {
		MIDICIDeviceManagerClass = _MIDICIDeviceManagerClass{objc.GetClass("MIDICIDeviceManager")}
	})
	return MIDICIDeviceManagerClass
}

type _MIDICIDeviceManagerClass struct {
	class objc.Class
}

// An interface definition for the [MIDICIDeviceManager] class.
type IMIDICIDeviceManager interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceManager
type MIDICIDeviceManager struct {
	objectivec.Object
}

// MIDICIDeviceManagerFrom constructs a [MIDICIDeviceManager] from an unsafe.Pointer.
func MIDICIDeviceManagerFrom(ptr unsafe.Pointer) MIDICIDeviceManager {
	return MIDICIDeviceManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MIDICIDeviceManagerClass) Alloc() MIDICIDeviceManager {
	rv := objc.Send[MIDICIDeviceManager](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MIDICIDeviceManagerClass) New() MIDICIDeviceManager {
	rv := objc.Send[MIDICIDeviceManager](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDICIDeviceManager) Init() MIDICIDeviceManager {
	rv := objc.Send[MIDICIDeviceManager](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDICIDeviceManager) Autorelease() MIDICIDeviceManager {
	rv := objc.Send[MIDICIDeviceManager](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDICIDeviceManager creates a new MIDICIDeviceManager instance.
func NewMIDICIDeviceManager() MIDICIDeviceManager {
	return getMIDICIDeviceManagerClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceManager/discoveredCIDevices
func (m_ MIDICIDeviceManager) DiscoveredCIDevices() []MIDICIDevice {
	rv := objc.Send[[]MIDICIDevice](m_.ID, objc.Sel("discoveredCIDevices"))
	return rv
}



