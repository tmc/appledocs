// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class trackingID */


/* debug [class_header]: Header for trackingID */
// The class instance for the [trackingID] class.
var (
	TrackingIDClass     _trackingIDClass
	TrackingIDClassOnce sync.Once
)

func gettrackingIDClass() _trackingIDClass {
	TrackingIDClassOnce.Do(func() {
		TrackingIDClass = _trackingIDClass{objc.GetClass("trackingID")}
	})
	return TrackingIDClass
}

type _trackingIDClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for trackingID */
// An interface definition for the [trackingID] class.
type ItrackingID interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for trackingID */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for trackingID */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for trackingID */
// Alloc allocates a new instance without initialization.
func (tc _trackingIDClass) Alloc() trackingID {
	rv := objc.Send[trackingID](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _trackingIDClass) New() trackingID {
	rv := objc.Send[trackingID](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ trackingID) Init() trackingID {
	rv := objc.Send[trackingID](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ trackingID) Autorelease() trackingID {
	rv := objc.Send[trackingID](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewtrackingID creates a new trackingID instance.
func NewtrackingID() trackingID {
	return gettrackingIDClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for trackingID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/trackingID-c.ivar
type trackingID struct {
	objectivec.Object
}

// trackingIDFrom constructs a [trackingID] from an unsafe.Pointer.
func trackingIDFrom(ptr unsafe.Pointer) trackingID {
	return trackingID{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for trackingID *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for trackingID */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for trackingID */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for trackingID */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for trackingID */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class trackingID */



