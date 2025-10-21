// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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


// The family to which the device belongs.
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicideviceinfo/family
func (m_ MIDICIDeviceInfo) Family() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("family"))
	return rv
}


// SetFamily sets the value of the family property.
// The family to which the device belongs.

//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicideviceinfo/family
func (m_ MIDICIDeviceInfo) SetFamily(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFamily:"), value)
}

// The MIDI System Exclusive (SysEx) ID of the device manufacturer.
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicideviceinfo/manufacturerid
func (m_ MIDICIDeviceInfo) ManufacturerID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("manufacturerID"))
	return rv
}


// SetManufacturerID sets the value of the manufacturerID property.
// The MIDI System Exclusive (SysEx) ID of the device manufacturer.

//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicideviceinfo/manufacturerid
func (m_ MIDICIDeviceInfo) SetManufacturerID(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setManufacturerID:"), value)
}

// The MIDI destination the device’s MIDI entity uses for capability inquiries.
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicideviceinfo/mididestination
func (m_ MIDICIDeviceInfo) MidiDestination() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("midiDestination"))
	return rv
}


// SetMidiDestination sets the value of the midiDestination property.
// The MIDI destination the device’s MIDI entity uses for capability inquiries.

//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicideviceinfo/mididestination
func (m_ MIDICIDeviceInfo) SetMidiDestination(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMidiDestination:"), value)
}

// The model number of the device.
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicideviceinfo/modelnumber
func (m_ MIDICIDeviceInfo) ModelNumber() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("modelNumber"))
	return rv
}


// SetModelNumber sets the value of the modelNumber property.
// The model number of the device.

//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicideviceinfo/modelnumber
func (m_ MIDICIDeviceInfo) SetModelNumber(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setModelNumber:"), value)
}

// The revision number of the device model number.
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicideviceinfo/revisionlevel
func (m_ MIDICIDeviceInfo) RevisionLevel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("revisionLevel"))
	return rv
}


// SetRevisionLevel sets the value of the revisionLevel property.
// The revision number of the device model number.

//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicideviceinfo/revisionlevel
func (m_ MIDICIDeviceInfo) SetRevisionLevel(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRevisionLevel:"), value)
}

// The MIDI-CI device’s information.
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciresponder/deviceinfo
func (m_ MIDICIDeviceInfo) DeviceInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("deviceInfo"))
	return rv
}


// SetDeviceInfo sets the value of the deviceInfo property.
// The MIDI-CI device’s information.

//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciresponder/deviceinfo
func (m_ MIDICIDeviceInfo) SetDeviceInfo(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDeviceInfo:"), value)
}

// An array of initiators.
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciresponder/initiators
func (m_ MIDICIDeviceInfo) Initiators() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("initiators"))
	return rv
}


// SetInitiators sets the value of the initiators property.
// An array of initiators.

//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciresponder/initiators
func (m_ MIDICIDeviceInfo) SetInitiators(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInitiators:"), value)
}



