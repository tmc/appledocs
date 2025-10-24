// Code generated from Apple documentation for NearbyInteraction. DO NOT EDIT.

package nearbyinteraction

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NIConfiguration */


/* debug [class_header]: Header for NIConfiguration */
// The class instance for the [NIConfiguration] class.
var (
	NIConfigurationClass     _NIConfigurationClass
	NIConfigurationClassOnce sync.Once
)

func getNIConfigurationClass() _NIConfigurationClass {
	NIConfigurationClassOnce.Do(func() {
		NIConfigurationClass = _NIConfigurationClass{objc.GetClass("NIConfiguration")}
	})
	return NIConfigurationClass
}

type _NIConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NIConfiguration */
// An interface definition for the [NIConfiguration] class.
type INIConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NIConfiguration */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NIConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NIConfiguration */
// Alloc allocates a new instance without initialization.
func (nc _NIConfigurationClass) Alloc() NIConfiguration {
	rv := objc.Send[NIConfiguration](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NIConfigurationClass) New() NIConfiguration {
	rv := objc.Send[NIConfiguration](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NIConfiguration) Init() NIConfiguration {
	rv := objc.Send[NIConfiguration](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NIConfiguration) Autorelease() NIConfiguration {
	rv := objc.Send[NIConfiguration](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNIConfiguration creates a new NIConfiguration instance.
func NewNIConfiguration() NIConfiguration {
	return getNIConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NIConfiguration */
// An abstract base class for interaction configurations.
//
// The class serves as the common identity for configuration objects. Don’t instantiate this class directly. Instead, instantiate one if its concrete subclasses: or . Use your configuration object to specify the features you want to enable in a Nearby Interaction session, and pass the object to the session’s   method.


// An abstract base class for interaction configurations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIConfiguration
type NIConfiguration struct {
	objectivec.Object
}

// NIConfigurationFrom constructs a [NIConfiguration] from an unsafe.Pointer.
//
// An abstract base class for interaction configurations.
func NIConfigurationFrom(ptr unsafe.Pointer) NIConfiguration {
	return NIConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NIConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NIConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NIConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NIConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NIConfiguration */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NIConfiguration */



