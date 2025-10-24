// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MIDICISession */


/* debug [class_header]: Header for MIDICISession */
// The class instance for the [MIDICISession] class.
var (
	MIDICISessionClass     _MIDICISessionClass
	MIDICISessionClassOnce sync.Once
)

func getMIDICISessionClass() _MIDICISessionClass {
	MIDICISessionClassOnce.Do(func() {
		MIDICISessionClass = _MIDICISessionClass{objc.GetClass("MIDICISession")}
	})
	return MIDICISessionClass
}

type _MIDICISessionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MIDICISession */
// An interface definition for the [MIDICISession] class.
type IMIDICISession interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MIDICISession */
	// properties:
	DeviceInfo() IMIDICIDeviceInfo
	MaxPropertyRequests() objc.IObject /* cross-framework: NSNumber */
	MaxSysExSize() objc.IObject /* cross-framework: NSNumber */
	MidiDestination() MIDIEntityRef /* typedef */
	ProfileChangedCallback() objectivec.IObject
	SetProfileChangedCallback(value objectivec.IObject)
	ProfileSpecificDataHandler() objectivec.IObject
	SetProfileSpecificDataHandler(value objectivec.IObject)
	SupportsProfileCapability() bool
	SupportsPropertyCapability() bool
	MIDIChannelsWholePort() MIDIChannelNumber /* typedef */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MIDICISession */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MIDICISession */
// Alloc allocates a new instance without initialization.
func (mc _MIDICISessionClass) Alloc() MIDICISession {
	rv := objc.Send[MIDICISession](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MIDICISessionClass) New() MIDICISession {
	rv := objc.Send[MIDICISession](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDICISession) Init() MIDICISession {
	rv := objc.Send[MIDICISession](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDICISession) Autorelease() MIDICISession {
	rv := objc.Send[MIDICISession](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDICISession creates a new MIDICISession instance.
func NewMIDICISession() MIDICISession {
	return getMIDICISessionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MIDICISession */
// An object that represents a MIDI-CI session.
//
// A MIDI-CI session is a bidirectional communication path between a MIDI source and destination identified using MIDI-CI discovery. Use a session to manipulate MIDI-CI profiles and to discover device capabilities.


// An object that represents a MIDI-CI session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICISession
type MIDICISession struct {
	objectivec.Object
}

// MIDICISessionFrom constructs a [MIDICISession] from an unsafe.Pointer.
//
// An object that represents a MIDI-CI session.
func MIDICISessionFrom(ptr unsafe.Pointer) MIDICISession {
	return MIDICISession{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MIDICISession */

// Creates a MIDI-CI session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICISession/init(discoveredNode:dataReadyHandler:disconnectHandler:)
func NewMIDICISessionWithDiscoveredNodeDataReadyHandlerDisconnectHandler(discoveredNode IMIDICIDiscoveredNode, handler unsafe.Pointer, disconnectHandler objectivec.IObject) MIDICISession {
	instance := getMIDICISessionClass().Alloc()
	rv := objc.Send[MIDICISession](instance.ID, objc.Sel("initWithDiscoveredNode:dataReadyHandler:disconnectHandler:"), discoveredNode, handler, disconnectHandler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMIDICISessionWithDiscoveredNodeDataReadyHandlerDisconnectHandler */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MIDICISession */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MIDICISession */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MIDICISession */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MIDICISession */

// Information about a MIDI-CI device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICISession/deviceInfo
func (m_ MIDICISession) DeviceInfo() IMIDICIDeviceInfo {
	rv := objc.Send[MIDICIDeviceInfo](m_.ID, objc.Sel("deviceInfo"))
	return rv
}/* debug [instance_properties/getter]: deviceInfo */


// The maximum number of simultaneous property exchange requests, if supported.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICISession/maxPropertyRequests
func (m_ MIDICISession) MaxPropertyRequests() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("maxPropertyRequests"))
	return rv
}/* debug [instance_properties/getter]: maxPropertyRequests */


// The maximum size of System Exclusive (SysEx) messages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICISession/maxSysExSize
func (m_ MIDICISession) MaxSysExSize() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("maxSysExSize"))
	return rv
}/* debug [instance_properties/getter]: maxSysExSize */


// The MIDI destination with which the session is communicating.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICISession/midiDestination
func (m_ MIDICISession) MidiDestination() MIDIEntityRef /* typedef */ {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("midiDestination"))
	return rv
}/* debug [instance_properties/getter]: midiDestination */


// An optional block the system calls after it enables or disables a profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICISession/profileChangedCallback
func (m_ MIDICISession) ProfileChangedCallback() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("profileChangedCallback"))
	return rv
}/* debug [instance_properties/getter]: profileChangedCallback */


// An optional block the system calls after it enables or disables a profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICISession/profileChangedCallback
func (m_ MIDICISession) SetProfileChangedCallback(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProfileChangedCallback:"), value)
}/* debug [instance_properties/setter]: profileChangedCallback */


// An optional block the system calls when a device sends profile-specific data to the session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICISession/profileSpecificDataHandler
func (m_ MIDICISession) ProfileSpecificDataHandler() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("profileSpecificDataHandler"))
	return rv
}/* debug [instance_properties/getter]: profileSpecificDataHandler */


// An optional block the system calls when a device sends profile-specific data to the session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICISession/profileSpecificDataHandler
func (m_ MIDICISession) SetProfileSpecificDataHandler(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProfileSpecificDataHandler:"), value)
}/* debug [instance_properties/setter]: profileSpecificDataHandler */


// A Boolean value that indicates whether the entity supports the MIDI-CI profile’s capability.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICISession/supportsProfileCapability
func (m_ MIDICISession) SupportsProfileCapability() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("supportsProfileCapability"))
	return rv
}/* debug [instance_properties/getter]: supportsProfileCapability */


// A Boolean value that indicates whether the entity supports the MIDI-CI property exchange capability.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICISession/supportsPropertyCapability
func (m_ MIDICISession) SupportsPropertyCapability() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("supportsPropertyCapability"))
	return rv
}/* debug [instance_properties/getter]: supportsPropertyCapability */


// A constant value that indicates to use all channels of the port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midichannelswholeport
func (m_ MIDICISession) MIDIChannelsWholePort() MIDIChannelNumber /* typedef */ {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("MIDIChannelsWholePort"))
	return rv
}/* debug [instance_properties/getter]: MIDIChannelsWholePort */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MIDICISession */


