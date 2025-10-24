// Code generated from Apple documentation for ScreenTime. DO NOT EDIT.

package screentime

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class STScreenTimeConfiguration */


/* debug [class_header]: Header for STScreenTimeConfiguration */
// The class instance for the [STScreenTimeConfiguration] class.
var (
	STScreenTimeConfigurationClass     _STScreenTimeConfigurationClass
	STScreenTimeConfigurationClassOnce sync.Once
)

func getSTScreenTimeConfigurationClass() _STScreenTimeConfigurationClass {
	STScreenTimeConfigurationClassOnce.Do(func() {
		STScreenTimeConfigurationClass = _STScreenTimeConfigurationClass{objc.GetClass("STScreenTimeConfiguration")}
	})
	return STScreenTimeConfigurationClass
}

type _STScreenTimeConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for STScreenTimeConfiguration */
// An interface definition for the [STScreenTimeConfiguration] class.
type ISTScreenTimeConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for STScreenTimeConfiguration */
	// properties:
	EnforcesChildRestrictions() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for STScreenTimeConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for STScreenTimeConfiguration */
// Alloc allocates a new instance without initialization.
func (sc _STScreenTimeConfigurationClass) Alloc() STScreenTimeConfiguration {
	rv := objc.Send[STScreenTimeConfiguration](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _STScreenTimeConfigurationClass) New() STScreenTimeConfiguration {
	rv := objc.Send[STScreenTimeConfiguration](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ STScreenTimeConfiguration) Init() STScreenTimeConfiguration {
	rv := objc.Send[STScreenTimeConfiguration](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ STScreenTimeConfiguration) Autorelease() STScreenTimeConfiguration {
	rv := objc.Send[STScreenTimeConfiguration](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSTScreenTimeConfiguration creates a new STScreenTimeConfiguration instance.
func NewSTScreenTimeConfiguration() STScreenTimeConfiguration {
	return getSTScreenTimeConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for STScreenTimeConfiguration */
// The configuration for this device.


// The configuration for this device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STScreenTimeConfiguration
type STScreenTimeConfiguration struct {
	objectivec.Object
}

// STScreenTimeConfigurationFrom constructs a [STScreenTimeConfiguration] from an unsafe.Pointer.
//
// The configuration for this device.
func STScreenTimeConfigurationFrom(ptr unsafe.Pointer) STScreenTimeConfiguration {
	return STScreenTimeConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for STScreenTimeConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for STScreenTimeConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for STScreenTimeConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for STScreenTimeConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for STScreenTimeConfiguration */

// A Boolean that indicates whether the device is currently enforcing child restrictions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STScreenTimeConfiguration/enforcesChildRestrictions
func (s_ STScreenTimeConfiguration) EnforcesChildRestrictions() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("enforcesChildRestrictions"))
	return rv
}/* debug [instance_properties/getter]: enforcesChildRestrictions */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class STScreenTimeConfiguration */



