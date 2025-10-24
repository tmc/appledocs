// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MIDICIResponder */


/* debug [class_header]: Header for MIDICIResponder */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MIDICIResponder */
// An interface definition for the [MIDICIResponder] class.
type IMIDICIResponder interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MIDICIResponder */
	// properties:
	DeviceInfo() IMIDICIDeviceInfo
	Initiators() []foundation.Number
	ProfileDelegate() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MIDICIResponder */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MIDICIResponder */
// Alloc allocates a new instance without initialization.
func (mc _MIDICIResponderClass) Alloc() MIDICIResponder {
	rv := objc.Send[MIDICIResponder](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MIDICIResponder */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MIDICIResponder */

// Creates a new responder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIResponder/init(deviceInfo:profileDelegate:profileStates:supportProperties:)
func NewMIDICIResponderWithDeviceInfoProfileDelegateProfileStatesSupportProperties(deviceInfo IMIDICIDeviceInfo, delegate unsafe.Pointer, profileList objectivec.IObject, propertiesSupported bool) MIDICIResponder {
	instance := getMIDICIResponderClass().Alloc()
	rv := objc.Send[MIDICIResponder](instance.ID, objc.Sel("initWithDeviceInfo:profileDelegate:profileStates:supportProperties:"), deviceInfo, delegate, profileList, propertiesSupported)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMIDICIResponderWithDeviceInfoProfileDelegateProfileStatesSupportProperties */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MIDICIResponder */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MIDICIResponder */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MIDICIResponder */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MIDICIResponder */

// The MIDI-CI device’s information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIResponder/deviceInfo
func (m_ MIDICIResponder) DeviceInfo() IMIDICIDeviceInfo {
	rv := objc.Send[MIDICIDeviceInfo](m_.ID, objc.Sel("deviceInfo"))
	return rv
}/* debug [instance_properties/getter]: deviceInfo */


// An array of initiators.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIResponder/initiators
func (m_ MIDICIResponder) Initiators() []foundation.Number {
	rv := objc.Send[[]foundation.Number](m_.ID, objc.Sel("initiators"))
	return rv
}/* debug [instance_properties/getter]: initiators */


// The profile delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIResponder/profileDelegate
func (m_ MIDICIResponder) ProfileDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("profileDelegate"))
	return rv
}/* debug [instance_properties/getter]: profileDelegate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MIDICIResponder */


