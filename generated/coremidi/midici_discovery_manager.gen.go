// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MIDICIDiscoveryManager */


/* debug [class_header]: Header for MIDICIDiscoveryManager */
// The class instance for the [MIDICIDiscoveryManager] class.
var (
	MIDICIDiscoveryManagerClass     _MIDICIDiscoveryManagerClass
	MIDICIDiscoveryManagerClassOnce sync.Once
)

func getMIDICIDiscoveryManagerClass() _MIDICIDiscoveryManagerClass {
	MIDICIDiscoveryManagerClassOnce.Do(func() {
		MIDICIDiscoveryManagerClass = _MIDICIDiscoveryManagerClass{objc.GetClass("MIDICIDiscoveryManager")}
	})
	return MIDICIDiscoveryManagerClass
}

type _MIDICIDiscoveryManagerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MIDICIDiscoveryManager */
// An interface definition for the [MIDICIDiscoveryManager] class.
type IMIDICIDiscoveryManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MIDICIDiscoveryManager */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MIDICIDiscoveryManager */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MIDICIDiscoveryManager */
// Alloc allocates a new instance without initialization.
func (mc _MIDICIDiscoveryManagerClass) Alloc() MIDICIDiscoveryManager {
	rv := objc.Send[MIDICIDiscoveryManager](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MIDICIDiscoveryManagerClass) New() MIDICIDiscoveryManager {
	rv := objc.Send[MIDICIDiscoveryManager](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDICIDiscoveryManager) Init() MIDICIDiscoveryManager {
	rv := objc.Send[MIDICIDiscoveryManager](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDICIDiscoveryManager) Autorelease() MIDICIDiscoveryManager {
	rv := objc.Send[MIDICIDiscoveryManager](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDICIDiscoveryManager creates a new MIDICIDiscoveryManager instance.
func NewMIDICIDiscoveryManager() MIDICIDiscoveryManager {
	return getMIDICIDiscoveryManagerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MIDICIDiscoveryManager */
// A singleton object that performs systemwide MIDI-CI discovery.
//
// Use this class to retrieve information about MIDI-CI–capable nodes in the MIDI subsystem. You can create objects only from the destinations discovered using this API.


// A singleton object that performs systemwide MIDI-CI discovery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDiscoveryManager
type MIDICIDiscoveryManager struct {
	objectivec.Object
}

// MIDICIDiscoveryManagerFrom constructs a [MIDICIDiscoveryManager] from an unsafe.Pointer.
//
// A singleton object that performs systemwide MIDI-CI discovery.
func MIDICIDiscoveryManagerFrom(ptr unsafe.Pointer) MIDICIDiscoveryManager {
	return MIDICIDiscoveryManager{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MIDICIDiscoveryManager *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MIDICIDiscoveryManager */

// Returns the singleton discovery manager instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDiscoveryManager/sharedInstance()
func (mc _MIDICIDiscoveryManagerClass) SharedInstance() MIDICIDiscoveryManager {
	rv := objc.Send[MIDICIDiscoveryManager](objc.ID(mc.class), objc.Sel("sharedInstance"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SharedInstance) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MIDICIDiscoveryManager */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MIDICIDiscoveryManager */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MIDICIDiscoveryManager */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MIDICIDiscoveryManager */



