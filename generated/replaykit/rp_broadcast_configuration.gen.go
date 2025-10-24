// Code generated from Apple documentation for ReplayKit. DO NOT EDIT.

package replaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class RPBroadcastConfiguration */


/* debug [class_header]: Header for RPBroadcastConfiguration */
// The class instance for the [RPBroadcastConfiguration] class.
var (
	RPBroadcastConfigurationClass     _RPBroadcastConfigurationClass
	RPBroadcastConfigurationClassOnce sync.Once
)

func getRPBroadcastConfigurationClass() _RPBroadcastConfigurationClass {
	RPBroadcastConfigurationClassOnce.Do(func() {
		RPBroadcastConfigurationClass = _RPBroadcastConfigurationClass{objc.GetClass("RPBroadcastConfiguration")}
	})
	return RPBroadcastConfigurationClass
}

type _RPBroadcastConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RPBroadcastConfiguration */
// An interface definition for the [RPBroadcastConfiguration] class.
type IRPBroadcastConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for RPBroadcastConfiguration */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RPBroadcastConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RPBroadcastConfiguration */
// Alloc allocates a new instance without initialization.
func (rc _RPBroadcastConfigurationClass) Alloc() RPBroadcastConfiguration {
	rv := objc.Send[RPBroadcastConfiguration](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RPBroadcastConfigurationClass) New() RPBroadcastConfiguration {
	rv := objc.Send[RPBroadcastConfiguration](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RPBroadcastConfiguration) Init() RPBroadcastConfiguration {
	rv := objc.Send[RPBroadcastConfiguration](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RPBroadcastConfiguration) Autorelease() RPBroadcastConfiguration {
	rv := objc.Send[RPBroadcastConfiguration](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRPBroadcastConfiguration creates a new RPBroadcastConfiguration instance.
func NewRPBroadcastConfiguration() RPBroadcastConfiguration {
	return getRPBroadcastConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RPBroadcastConfiguration */
// An object used to configure the movie clips produced during a live broadcast.


// An object used to configure the movie clips produced during a live broadcast.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastConfiguration
type RPBroadcastConfiguration struct {
	objectivec.Object
}

// RPBroadcastConfigurationFrom constructs a [RPBroadcastConfiguration] from an unsafe.Pointer.
//
// An object used to configure the movie clips produced during a live broadcast.
func RPBroadcastConfigurationFrom(ptr unsafe.Pointer) RPBroadcastConfiguration {
	return RPBroadcastConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RPBroadcastConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RPBroadcastConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RPBroadcastConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RPBroadcastConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RPBroadcastConfiguration */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class RPBroadcastConfiguration */


