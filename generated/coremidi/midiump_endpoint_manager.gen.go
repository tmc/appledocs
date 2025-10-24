// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MIDIUMPEndpointManager */


/* debug [class_header]: Header for MIDIUMPEndpointManager */
// The class instance for the [MIDIUMPEndpointManager] class.
var (
	MIDIUMPEndpointManagerClass     _MIDIUMPEndpointManagerClass
	MIDIUMPEndpointManagerClassOnce sync.Once
)

func getMIDIUMPEndpointManagerClass() _MIDIUMPEndpointManagerClass {
	MIDIUMPEndpointManagerClassOnce.Do(func() {
		MIDIUMPEndpointManagerClass = _MIDIUMPEndpointManagerClass{objc.GetClass("MIDIUMPEndpointManager")}
	})
	return MIDIUMPEndpointManagerClass
}

type _MIDIUMPEndpointManagerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MIDIUMPEndpointManager */
// An interface definition for the [MIDIUMPEndpointManager] class.
type IMIDIUMPEndpointManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MIDIUMPEndpointManager */
	// properties:
	UMPEndpoints() []MIDIUMPEndpoint
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MIDIUMPEndpointManager */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MIDIUMPEndpointManager */
// Alloc allocates a new instance without initialization.
func (mc _MIDIUMPEndpointManagerClass) Alloc() MIDIUMPEndpointManager {
	rv := objc.Send[MIDIUMPEndpointManager](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MIDIUMPEndpointManagerClass) New() MIDIUMPEndpointManager {
	rv := objc.Send[MIDIUMPEndpointManager](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDIUMPEndpointManager) Init() MIDIUMPEndpointManager {
	rv := objc.Send[MIDIUMPEndpointManager](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDIUMPEndpointManager) Autorelease() MIDIUMPEndpointManager {
	rv := objc.Send[MIDIUMPEndpointManager](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDIUMPEndpointManager creates a new MIDIUMPEndpointManager instance.
func NewMIDIUMPEndpointManager() MIDIUMPEndpointManager {
	return getMIDIUMPEndpointManagerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MIDIUMPEndpointManager */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpointManager
type MIDIUMPEndpointManager struct {
	objectivec.Object
}

// MIDIUMPEndpointManagerFrom constructs a [MIDIUMPEndpointManager] from an unsafe.Pointer.
func MIDIUMPEndpointManagerFrom(ptr unsafe.Pointer) MIDIUMPEndpointManager {
	return MIDIUMPEndpointManager{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MIDIUMPEndpointManager *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MIDIUMPEndpointManager */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MIDIUMPEndpointManager */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpointManager/shared
func (mc _MIDIUMPEndpointManagerClass) SharedInstance() MIDIUMPEndpointManager {
	rv := objc.Send[MIDIUMPEndpointManager](objc.ID(mc.class), objc.Sel("sharedInstance"))
	return rv
}/* debug [class_properties_class/property]: sharedInstance */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MIDIUMPEndpointManager */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MIDIUMPEndpointManager */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpointManager/shared
func (m_ MIDIUMPEndpointManager) SharedInstance() IMIDIUMPEndpointManager {
	rv := objc.Send[MIDIUMPEndpointManager](m_.ID, objc.Sel("sharedInstance"))
	return rv
}/* debug [instance_properties/getter]: sharedInstance */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpointManager/umpEndpoints
func (m_ MIDIUMPEndpointManager) UMPEndpoints() []MIDIUMPEndpoint {
	rv := objc.Send[[]MIDIUMPEndpoint](m_.ID, objc.Sel("UMPEndpoints"))
	return rv
}/* debug [instance_properties/getter]: UMPEndpoints */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MIDIUMPEndpointManager */



