// Code generated from Apple documentation for AVRouting. DO NOT EDIT.

package avrouting

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVRoutingPlaybackArbiter */


/* debug [class_header]: Header for AVRoutingPlaybackArbiter */
// The class instance for the [RoutingPlaybackArbiter] class.
var (
	RoutingPlaybackArbiterClass     _RoutingPlaybackArbiterClass
	RoutingPlaybackArbiterClassOnce sync.Once
)

func getRoutingPlaybackArbiterClass() _RoutingPlaybackArbiterClass {
	RoutingPlaybackArbiterClassOnce.Do(func() {
		RoutingPlaybackArbiterClass = _RoutingPlaybackArbiterClass{objc.GetClass("AVRoutingPlaybackArbiter")}
	})
	return RoutingPlaybackArbiterClass
}

type _RoutingPlaybackArbiterClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RoutingPlaybackArbiter */
// An interface definition for the [RoutingPlaybackArbiter] class.
type IRoutingPlaybackArbiter interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for RoutingPlaybackArbiter */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RoutingPlaybackArbiter */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RoutingPlaybackArbiter */
// Alloc allocates a new instance without initialization.
func (rc _RoutingPlaybackArbiterClass) Alloc() RoutingPlaybackArbiter {
	rv := objc.Send[RoutingPlaybackArbiter](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RoutingPlaybackArbiterClass) New() RoutingPlaybackArbiter {
	rv := objc.Send[RoutingPlaybackArbiter](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RoutingPlaybackArbiter) Init() RoutingPlaybackArbiter {
	rv := objc.Send[RoutingPlaybackArbiter](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RoutingPlaybackArbiter) Autorelease() RoutingPlaybackArbiter {
	rv := objc.Send[RoutingPlaybackArbiter](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRoutingPlaybackArbiter creates a new RoutingPlaybackArbiter instance.
func NewRoutingPlaybackArbiter() RoutingPlaybackArbiter {
	return getRoutingPlaybackArbiterClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RoutingPlaybackArbiter */
// An object that manages playback routing preferences.
//
// This object manages instances of for arbitration of media playback routing priorities and preferences on restricted playback interfaces. The playback routing arbiter is responsible for collecting and applying preferences, such as priorities in non-mixable audio routes and external playback states where the number of allowed players is limited.


// An object that manages playback routing preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVRoutingPlaybackArbiter
type RoutingPlaybackArbiter struct {
	objectivec.Object
}

// RoutingPlaybackArbiterFrom constructs a [RoutingPlaybackArbiter] from an unsafe.Pointer.
//
// An object that manages playback routing preferences.
func RoutingPlaybackArbiterFrom(ptr unsafe.Pointer) RoutingPlaybackArbiter {
	return RoutingPlaybackArbiter{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RoutingPlaybackArbiter *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RoutingPlaybackArbiter */

// Returns the singleton playback arbiter instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVRoutingPlaybackArbiter/shared()
func (rc _RoutingPlaybackArbiterClass) SharedRoutingPlaybackArbiter() IRoutingPlaybackArbiter {
	rv := objc.Send[RoutingPlaybackArbiter](objc.ID(rc.class), objc.Sel("sharedRoutingPlaybackArbiter"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SharedRoutingPlaybackArbiter) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RoutingPlaybackArbiter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RoutingPlaybackArbiter */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RoutingPlaybackArbiter */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVRoutingPlaybackArbiter */


