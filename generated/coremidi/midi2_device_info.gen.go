// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MIDI2DeviceInfo] class.
var (
	MIDI2DeviceInfoClass     _MIDI2DeviceInfoClass
	MIDI2DeviceInfoClassOnce sync.Once
)

func getMIDI2DeviceInfoClass() _MIDI2DeviceInfoClass {
	MIDI2DeviceInfoClassOnce.Do(func() {
		MIDI2DeviceInfoClass = _MIDI2DeviceInfoClass{objc.GetClass("MIDI2DeviceInfo")}
	})
	return MIDI2DeviceInfoClass
}

type _MIDI2DeviceInfoClass struct {
	class objc.Class
}

// An interface definition for the [MIDI2DeviceInfo] class.
type IMIDI2DeviceInfo interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDI2DeviceInfo
type MIDI2DeviceInfo struct {
	objectivec.Object
}

// MIDI2DeviceInfoFrom constructs a [MIDI2DeviceInfo] from an unsafe.Pointer.
func MIDI2DeviceInfoFrom(ptr unsafe.Pointer) MIDI2DeviceInfo {
	return MIDI2DeviceInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MIDI2DeviceInfoClass) Alloc() MIDI2DeviceInfo {
	rv := objc.Send[MIDI2DeviceInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MIDI2DeviceInfoClass) New() MIDI2DeviceInfo {
	rv := objc.Send[MIDI2DeviceInfo](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDI2DeviceInfo) Init() MIDI2DeviceInfo {
	rv := objc.Send[MIDI2DeviceInfo](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDI2DeviceInfo) Autorelease() MIDI2DeviceInfo {
	rv := objc.Send[MIDI2DeviceInfo](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDI2DeviceInfo creates a new MIDI2DeviceInfo instance.
func NewMIDI2DeviceInfo() MIDI2DeviceInfo {
	return getMIDI2DeviceInfoClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDI2DeviceInfo/init(manufacturerID:family:modelNumber:revisionLevel:)
func NewMIDI2DeviceInfoWithManufacturerIDFamilyModelNumberRevisionLevel(manufacturerID unsafe.Pointer, family IMIDIUInteger14, modelNumber IMIDIUInteger14, revisionLevel unsafe.Pointer) MIDI2DeviceInfo {
	instance := getMIDI2DeviceInfoClass().Alloc()
	rv := objc.Send[MIDI2DeviceInfo](instance.ID, objc.Sel("initWithManufacturerID:family:modelNumber:revisionLevel:"), manufacturerID, family, modelNumber, revisionLevel)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDI2DeviceInfo/family
func (m_ MIDI2DeviceInfo) Family() MIDIUInteger14 {
	rv := objc.Send[MIDIUInteger14](m_.ID, objc.Sel("family"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDI2DeviceInfo/manufacturerID
func (m_ MIDI2DeviceInfo) ManufacturerID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("manufacturerID"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDI2DeviceInfo/modelNumber
func (m_ MIDI2DeviceInfo) ModelNumber() MIDIUInteger14 {
	rv := objc.Send[MIDIUInteger14](m_.ID, objc.Sel("modelNumber"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDI2DeviceInfo/revisionLevel
func (m_ MIDI2DeviceInfo) RevisionLevel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("revisionLevel"))
	return rv
}


