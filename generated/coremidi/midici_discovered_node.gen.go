// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MIDICIDiscoveredNode */


/* debug [class_header]: Header for MIDICIDiscoveredNode */
// The class instance for the [MIDICIDiscoveredNode] class.
var (
	MIDICIDiscoveredNodeClass     _MIDICIDiscoveredNodeClass
	MIDICIDiscoveredNodeClassOnce sync.Once
)

func getMIDICIDiscoveredNodeClass() _MIDICIDiscoveredNodeClass {
	MIDICIDiscoveredNodeClassOnce.Do(func() {
		MIDICIDiscoveredNodeClass = _MIDICIDiscoveredNodeClass{objc.GetClass("MIDICIDiscoveredNode")}
	})
	return MIDICIDiscoveredNodeClass
}

type _MIDICIDiscoveredNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MIDICIDiscoveredNode */
// An interface definition for the [MIDICIDiscoveredNode] class.
type IMIDICIDiscoveredNode interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MIDICIDiscoveredNode */
	// properties:
	Destination() MIDIEntityRef /* typedef */
	DeviceInfo() IMIDICIDeviceInfo
	MaximumSysExSize() objc.IObject /* cross-framework: NSNumber */
	SupportsProfiles() bool
	SupportsProperties() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MIDICIDiscoveredNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MIDICIDiscoveredNode */
// Alloc allocates a new instance without initialization.
func (mc _MIDICIDiscoveredNodeClass) Alloc() MIDICIDiscoveredNode {
	rv := objc.Send[MIDICIDiscoveredNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MIDICIDiscoveredNodeClass) New() MIDICIDiscoveredNode {
	rv := objc.Send[MIDICIDiscoveredNode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDICIDiscoveredNode) Init() MIDICIDiscoveredNode {
	rv := objc.Send[MIDICIDiscoveredNode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDICIDiscoveredNode) Autorelease() MIDICIDiscoveredNode {
	rv := objc.Send[MIDICIDiscoveredNode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDICIDiscoveredNode creates a new MIDICIDiscoveredNode instance.
func NewMIDICIDiscoveredNode() MIDICIDiscoveredNode {
	return getMIDICIDiscoveredNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MIDICIDiscoveredNode */
// A discovered MIDI-CI node that represents a MIDI source and destination that respond to capability inquiries.


// A discovered MIDI-CI node that represents a MIDI source and destination that respond to capability inquiries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDiscoveredNode
type MIDICIDiscoveredNode struct {
	objectivec.Object
}

// MIDICIDiscoveredNodeFrom constructs a [MIDICIDiscoveredNode] from an unsafe.Pointer.
//
// A discovered MIDI-CI node that represents a MIDI source and destination that respond to capability inquiries.
func MIDICIDiscoveredNodeFrom(ptr unsafe.Pointer) MIDICIDiscoveredNode {
	return MIDICIDiscoveredNode{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MIDICIDiscoveredNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MIDICIDiscoveredNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MIDICIDiscoveredNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MIDICIDiscoveredNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MIDICIDiscoveredNode */

// The node’s MIDI destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDiscoveredNode/destination
func (m_ MIDICIDiscoveredNode) Destination() MIDIEntityRef /* typedef */ {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("destination"))
	return rv
}/* debug [instance_properties/getter]: destination */


// The available MIDI-CI device information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDiscoveredNode/deviceInfo
func (m_ MIDICIDiscoveredNode) DeviceInfo() IMIDICIDeviceInfo {
	rv := objc.Send[MIDICIDeviceInfo](m_.ID, objc.Sel("deviceInfo"))
	return rv
}/* debug [instance_properties/getter]: deviceInfo */


// The maximum size of a System Exclusive (SysEx) message this node supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDiscoveredNode/maximumSysExSize
func (m_ MIDICIDiscoveredNode) MaximumSysExSize() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("maximumSysExSize"))
	return rv
}/* debug [instance_properties/getter]: maximumSysExSize */


// A Boolean value that indicates whether this node supports MIDI-CI profiles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDiscoveredNode/supportsProfiles
func (m_ MIDICIDiscoveredNode) SupportsProfiles() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("supportsProfiles"))
	return rv
}/* debug [instance_properties/getter]: supportsProfiles */


// A Boolean value that indicates whether this node supports MIDI-CI properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDiscoveredNode/supportsProperties
func (m_ MIDICIDiscoveredNode) SupportsProperties() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("supportsProperties"))
	return rv
}/* debug [instance_properties/getter]: supportsProperties */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MIDICIDiscoveredNode */



