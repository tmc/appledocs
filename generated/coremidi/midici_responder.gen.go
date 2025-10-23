// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MIDICIResponder] class.
var (
	MIDICIResponderClass     _MIDICIResponderClass
	MIDICIResponderClassOnce sync.Once
)

func getMIDICIResponderClass() _MIDICIResponderClass {
	MIDICIResponderClassOnce.Do(func() {
		MIDICIResponderClass = _MIDICIResponderClass{objc.GetClass("MIDICIResponder")}
	})
	return MIDICIResponderClass
}

type _MIDICIResponderClass struct {
	class objc.Class
}

// An interface definition for the [MIDICIResponder] class.
type IMIDICIResponder interface {
	objectivec.IObject
	DeviceInfo() MIDICIDeviceInfo
	SetDeviceInfo(value MIDICIDeviceInfo)
	Initiators() unsafe.Pointer
	SetInitiators(value unsafe.Pointer)
	ProfileDelegate() unsafe.Pointer
	SetProfileDelegate(value unsafe.Pointer)
}

// An object that responds to MIDI-CI inquiries from an initiator on behalf of a MIDI client, and handles profile and property exchange operations.


// An object that responds to MIDI-CI inquiries from an initiator on behalf of a MIDI client, and handles profile and property exchange operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIResponder
type MIDICIResponder struct {
	objectivec.Object
}

// MIDICIResponderFrom constructs a [MIDICIResponder] from an unsafe.Pointer.
//
// An object that responds to MIDI-CI inquiries from an initiator on behalf of a MIDI client, and handles profile and property exchange operations.
func MIDICIResponderFrom(ptr unsafe.Pointer) MIDICIResponder {
	return MIDICIResponder{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MIDICIResponderClass) Alloc() MIDICIResponder {
	rv := objc.Send[MIDICIResponder](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MIDICIResponderClass) New() MIDICIResponder {
	rv := objc.Send[MIDICIResponder](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDICIResponder) Init() MIDICIResponder {
	rv := objc.Send[MIDICIResponder](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDICIResponder) Autorelease() MIDICIResponder {
	rv := objc.Send[MIDICIResponder](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDICIResponder creates a new MIDICIResponder instance.
func NewMIDICIResponder() MIDICIResponder {
	return getMIDICIResponderClass().New()
}



// The MIDI-CI device’s information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciresponder/deviceinfo
func (m_ MIDICIResponder) DeviceInfo() MIDICIDeviceInfo {
	rv := objc.Send[MIDICIDeviceInfo](m_.ID, objc.Sel("deviceInfo"))
	return rv
}


// The MIDI-CI device’s information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciresponder/deviceinfo
func (m_ MIDICIResponder) SetDeviceInfo(value MIDICIDeviceInfo) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDeviceInfo:"), value)
}


// An array of initiators.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciresponder/initiators
func (m_ MIDICIResponder) Initiators() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("initiators"))
	return rv
}


// An array of initiators.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciresponder/initiators
func (m_ MIDICIResponder) SetInitiators(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInitiators:"), value)
}


// The profile delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciresponder/profiledelegate
func (m_ MIDICIResponder) ProfileDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("profileDelegate"))
	return rv
}


// The profile delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciresponder/profiledelegate
func (m_ MIDICIResponder) SetProfileDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProfileDelegate:"), value)
}



