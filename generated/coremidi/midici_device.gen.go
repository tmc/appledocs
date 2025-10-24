// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MIDICIDevice */


/* debug [class_header]: Header for MIDICIDevice */
// The class instance for the [MIDICIDevice] class.
var (
	MIDICIDeviceClass     _MIDICIDeviceClass
	MIDICIDeviceClassOnce sync.Once
)

func getMIDICIDeviceClass() _MIDICIDeviceClass {
	MIDICIDeviceClassOnce.Do(func() {
		MIDICIDeviceClass = _MIDICIDeviceClass{objc.GetClass("MIDICIDevice")}
	})
	return MIDICIDeviceClass
}

type _MIDICIDeviceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MIDICIDevice */
// An interface definition for the [MIDICIDevice] class.
type IMIDICIDevice interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MIDICIDevice */
	// properties:
	DeviceInfo() IMIDI2DeviceInfo
	DeviceType() MIDICIDeviceType
	MaxPropertyExchangeRequests() uint
	MaxSysExSize() uint
	MUID() MIDICIMUID /* typedef */
	Profiles() []MIDIUMPCIProfile
	SupportsProcessInquiry() bool
	SupportsProfileConfiguration() bool
	SupportsPropertyExchange() bool
	SupportsProtocolNegotiation() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MIDICIDevice */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MIDICIDevice */
// Alloc allocates a new instance without initialization.
func (mc _MIDICIDeviceClass) Alloc() MIDICIDevice {
	rv := objc.Send[MIDICIDevice](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MIDICIDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDevice
type MIDICIDevice struct {
	objectivec.Object
}

// MIDICIDeviceFrom constructs a [MIDICIDevice] from an unsafe.Pointer.
func MIDICIDeviceFrom(ptr unsafe.Pointer) MIDICIDevice {
	return MIDICIDevice{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MIDICIDevice *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MIDICIDevice */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MIDICIDevice */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MIDICIDevice */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MIDICIDevice */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDevice/deviceInfo
func (m_ MIDICIDevice) DeviceInfo() IMIDI2DeviceInfo {
	rv := objc.Send[MIDI2DeviceInfo](m_.ID, objc.Sel("deviceInfo"))
	return rv
}/* debug [instance_properties/getter]: deviceInfo */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDevice/deviceType
func (m_ MIDICIDevice) DeviceType() MIDICIDeviceType {
	rv := objc.Send[MIDICIDeviceType](m_.ID, objc.Sel("deviceType"))
	return rv
}/* debug [instance_properties/getter]: deviceType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDevice/maxPropertyExchangeRequests
func (m_ MIDICIDevice) MaxPropertyExchangeRequests() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maxPropertyExchangeRequests"))
	return rv
}/* debug [instance_properties/getter]: maxPropertyExchangeRequests */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDevice/maxSysExSize
func (m_ MIDICIDevice) MaxSysExSize() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maxSysExSize"))
	return rv
}/* debug [instance_properties/getter]: maxSysExSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDevice/muid
func (m_ MIDICIDevice) MUID() MIDICIMUID /* typedef */ {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("MUID"))
	return rv
}/* debug [instance_properties/getter]: MUID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDevice/profiles
func (m_ MIDICIDevice) Profiles() []MIDIUMPCIProfile {
	rv := objc.Send[[]MIDIUMPCIProfile](m_.ID, objc.Sel("profiles"))
	return rv
}/* debug [instance_properties/getter]: profiles */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDevice/supportsProcessInquiry
func (m_ MIDICIDevice) SupportsProcessInquiry() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("supportsProcessInquiry"))
	return rv
}/* debug [instance_properties/getter]: supportsProcessInquiry */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDevice/supportsProfileConfiguration
func (m_ MIDICIDevice) SupportsProfileConfiguration() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("supportsProfileConfiguration"))
	return rv
}/* debug [instance_properties/getter]: supportsProfileConfiguration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDevice/supportsPropertyExchange
func (m_ MIDICIDevice) SupportsPropertyExchange() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("supportsPropertyExchange"))
	return rv
}/* debug [instance_properties/getter]: supportsPropertyExchange */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDevice/supportsProtocolNegotiation
func (m_ MIDICIDevice) SupportsProtocolNegotiation() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("supportsProtocolNegotiation"))
	return rv
}/* debug [instance_properties/getter]: supportsProtocolNegotiation */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MIDICIDevice */



