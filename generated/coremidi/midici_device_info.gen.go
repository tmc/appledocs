// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MIDICIDeviceInfo */


/* debug [class_header]: Header for MIDICIDeviceInfo */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MIDICIDeviceInfo */
// An interface definition for the [MIDICIDeviceInfo] class.
type IMIDICIDeviceInfo interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MIDICIDeviceInfo */
	// properties:
	Family() objc.IObject /* cross-framework: NSData */
	ManufacturerID() objc.IObject /* cross-framework: NSData */
	MidiDestination() MIDIEndpointRef /* typedef */
	ModelNumber() objc.IObject /* cross-framework: NSData */
	RevisionLevel() objc.IObject /* cross-framework: NSData */
	DeviceInfo() IMIDICIDeviceInfo
	SetDeviceInfo(value IMIDICIDeviceInfo)
	Initiators() MIDICIInitiatiorMUID /* typedef */
	SetInitiators(value MIDICIInitiatiorMUID /* typedef */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MIDICIDeviceInfo */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MIDICIDeviceInfo */
// Alloc allocates a new instance without initialization.
func (mc _MIDICIDeviceInfoClass) Alloc() MIDICIDeviceInfo {
	rv := objc.Send[MIDICIDeviceInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MIDICIDeviceInfo */
// An object that provides basic information about a MIDI-CI device.


// An object that provides basic information about a MIDI-CI device.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MIDICIDeviceInfo */

// Creates a new device information instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceInfo/init(destination:manufacturer:family:model:revision:)
func NewMIDICIDeviceInfoWithDestinationManufacturerFamilyModelRevision(midiDestination MIDIEntityRef /* typedef */, manufacturer objc.IObject /* cross-framework: NSData */, family objc.IObject /* cross-framework: NSData */, modelNumber objc.IObject /* cross-framework: NSData */, revisionLevel objc.IObject /* cross-framework: NSData */) MIDICIDeviceInfo {
	instance := getMIDICIDeviceInfoClass().Alloc()
	rv := objc.Send[MIDICIDeviceInfo](instance.ID, objc.Sel("initWithDestination:manufacturer:family:model:revision:"), midiDestination, manufacturer, family, modelNumber, revisionLevel)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMIDICIDeviceInfoWithDestinationManufacturerFamilyModelRevision */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MIDICIDeviceInfo */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MIDICIDeviceInfo */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MIDICIDeviceInfo */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MIDICIDeviceInfo */

// The family to which the device belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceInfo/family
func (m_ MIDICIDeviceInfo) Family() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("family"))
	return rv
}/* debug [instance_properties/getter]: family */


// The MIDI System Exclusive (SysEx) ID of the device manufacturer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceInfo/manufacturerID
func (m_ MIDICIDeviceInfo) ManufacturerID() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("manufacturerID"))
	return rv
}/* debug [instance_properties/getter]: manufacturerID */


// The MIDI destination the device’s MIDI entity uses for capability inquiries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceInfo/midiDestination
func (m_ MIDICIDeviceInfo) MidiDestination() MIDIEndpointRef /* typedef */ {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("midiDestination"))
	return rv
}/* debug [instance_properties/getter]: midiDestination */


// The model number of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceInfo/modelNumber
func (m_ MIDICIDeviceInfo) ModelNumber() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("modelNumber"))
	return rv
}/* debug [instance_properties/getter]: modelNumber */


// The revision number of the device model number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceInfo/revisionLevel
func (m_ MIDICIDeviceInfo) RevisionLevel() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("revisionLevel"))
	return rv
}/* debug [instance_properties/getter]: revisionLevel */


// The MIDI-CI device’s information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciresponder/deviceinfo
func (m_ MIDICIDeviceInfo) DeviceInfo() IMIDICIDeviceInfo {
	rv := objc.Send[MIDICIDeviceInfo](m_.ID, objc.Sel("deviceInfo"))
	return rv
}/* debug [instance_properties/getter]: deviceInfo */


// The MIDI-CI device’s information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciresponder/deviceinfo
func (m_ MIDICIDeviceInfo) SetDeviceInfo(value IMIDICIDeviceInfo) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDeviceInfo:"), value)
}/* debug [instance_properties/setter]: deviceInfo */


// An array of initiators.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciresponder/initiators
func (m_ MIDICIDeviceInfo) Initiators() MIDICIInitiatiorMUID /* typedef */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("initiators"))
	return rv
}/* debug [instance_properties/getter]: initiators */


// An array of initiators.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciresponder/initiators
func (m_ MIDICIDeviceInfo) SetInitiators(value MIDICIInitiatiorMUID /* typedef */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInitiators:"), value)
}/* debug [instance_properties/setter]: initiators */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MIDICIDeviceInfo */


