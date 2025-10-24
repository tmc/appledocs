// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MIDICIDeviceManager */


/* debug [class_header]: Header for MIDICIDeviceManager */
// The class instance for the [MIDICIDeviceManager] class.
var (
	MIDICIDeviceManagerClass     _MIDICIDeviceManagerClass
	MIDICIDeviceManagerClassOnce sync.Once
)

func getMIDICIDeviceManagerClass() _MIDICIDeviceManagerClass {
	MIDICIDeviceManagerClassOnce.Do(func() {
		MIDICIDeviceManagerClass = _MIDICIDeviceManagerClass{objc.GetClass("MIDICIDeviceManager")}
	})
	return MIDICIDeviceManagerClass
}

type _MIDICIDeviceManagerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MIDICIDeviceManager */
// An interface definition for the [MIDICIDeviceManager] class.
type IMIDICIDeviceManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MIDICIDeviceManager */
	// properties:
	DiscoveredCIDevices() []MIDICIDevice
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MIDICIDeviceManager */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MIDICIDeviceManager */
// Alloc allocates a new instance without initialization.
func (mc _MIDICIDeviceManagerClass) Alloc() MIDICIDeviceManager {
	rv := objc.Send[MIDICIDeviceManager](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MIDICIDeviceManagerClass) New() MIDICIDeviceManager {
	rv := objc.Send[MIDICIDeviceManager](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDICIDeviceManager) Init() MIDICIDeviceManager {
	rv := objc.Send[MIDICIDeviceManager](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDICIDeviceManager) Autorelease() MIDICIDeviceManager {
	rv := objc.Send[MIDICIDeviceManager](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDICIDeviceManager creates a new MIDICIDeviceManager instance.
func NewMIDICIDeviceManager() MIDICIDeviceManager {
	return getMIDICIDeviceManagerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MIDICIDeviceManager */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceManager
type MIDICIDeviceManager struct {
	objectivec.Object
}

// MIDICIDeviceManagerFrom constructs a [MIDICIDeviceManager] from an unsafe.Pointer.
func MIDICIDeviceManagerFrom(ptr unsafe.Pointer) MIDICIDeviceManager {
	return MIDICIDeviceManager{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MIDICIDeviceManager *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MIDICIDeviceManager */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MIDICIDeviceManager */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceManager/shared
func (mc _MIDICIDeviceManagerClass) SharedInstance() MIDICIDeviceManager {
	rv := objc.Send[MIDICIDeviceManager](objc.ID(mc.class), objc.Sel("sharedInstance"))
	return rv
}/* debug [class_properties_class/property]: sharedInstance */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MIDICIDeviceManager */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MIDICIDeviceManager */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceManager/discoveredCIDevices
func (m_ MIDICIDeviceManager) DiscoveredCIDevices() []MIDICIDevice {
	rv := objc.Send[[]MIDICIDevice](m_.ID, objc.Sel("discoveredCIDevices"))
	return rv
}/* debug [instance_properties/getter]: discoveredCIDevices */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceManager/shared
func (m_ MIDICIDeviceManager) SharedInstance() IMIDICIDeviceManager {
	rv := objc.Send[MIDICIDeviceManager](m_.ID, objc.Sel("sharedInstance"))
	return rv
}/* debug [instance_properties/getter]: sharedInstance */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MIDICIDeviceManager */



