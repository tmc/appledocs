// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MIDICIDeviceInfo] class.
var (
	MIDICIDeviceInfoClass     _MIDICIDeviceInfoClass
	MIDICIDeviceInfoClassOnce sync.Once
)

func getMIDICIDeviceInfoClass() _MIDICIDeviceInfoClass {
	MIDICIDeviceInfoClassOnce.Do(func() {
		MIDICIDeviceInfoClass = _MIDICIDeviceInfoClass{objc.GetClass("MIDICIDeviceInfo")}
	})
	return MIDICIDeviceInfoClass
}

type _MIDICIDeviceInfoClass struct {
	class objc.Class
}

// An interface definition for the [MIDICIDeviceInfo] class.
type IMIDICIDeviceInfo interface {
	objectivec.IObject
}

// An object that provides basic information about a MIDI-CI device.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceInfo
type MIDICIDeviceInfo struct {
	objectivec.Object
}

// MIDICIDeviceInfoFrom constructs a [MIDICIDeviceInfo] from an unsafe.Pointer.
//
// An object that provides basic information about a MIDI-CI device.
func MIDICIDeviceInfoFrom(ptr unsafe.Pointer) MIDICIDeviceInfo {
	return MIDICIDeviceInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MIDICIDeviceInfoClass) Alloc() MIDICIDeviceInfo {
	rv := objc.Send[MIDICIDeviceInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MIDICIDeviceInfoClass) New() MIDICIDeviceInfo {
	rv := objc.Send[MIDICIDeviceInfo](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDICIDeviceInfo) Init() MIDICIDeviceInfo {
	rv := objc.Send[MIDICIDeviceInfo](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDICIDeviceInfo) Autorelease() MIDICIDeviceInfo {
	rv := objc.Send[MIDICIDeviceInfo](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDICIDeviceInfo creates a new MIDICIDeviceInfo instance.
func NewMIDICIDeviceInfo() MIDICIDeviceInfo {
	return getMIDICIDeviceInfoClass().New()
}




