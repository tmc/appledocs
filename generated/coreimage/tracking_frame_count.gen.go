// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class trackingFrameCount */


/* debug [class_header]: Header for trackingFrameCount */
// The class instance for the [trackingFrameCount] class.
var (
	TrackingFrameCountClass     _trackingFrameCountClass
	TrackingFrameCountClassOnce sync.Once
)

func gettrackingFrameCountClass() _trackingFrameCountClass {
	TrackingFrameCountClassOnce.Do(func() {
		TrackingFrameCountClass = _trackingFrameCountClass{objc.GetClass("trackingFrameCount")}
	})
	return TrackingFrameCountClass
}

type _trackingFrameCountClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for trackingFrameCount */
// An interface definition for the [trackingFrameCount] class.
type ItrackingFrameCount interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for trackingFrameCount */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for trackingFrameCount */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for trackingFrameCount */
// Alloc allocates a new instance without initialization.
func (tc _trackingFrameCountClass) Alloc() trackingFrameCount {
	rv := objc.Send[trackingFrameCount](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _trackingFrameCountClass) New() trackingFrameCount {
	rv := objc.Send[trackingFrameCount](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ trackingFrameCount) Init() trackingFrameCount {
	rv := objc.Send[trackingFrameCount](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ trackingFrameCount) Autorelease() trackingFrameCount {
	rv := objc.Send[trackingFrameCount](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewtrackingFrameCount creates a new trackingFrameCount instance.
func NewtrackingFrameCount() trackingFrameCount {
	return gettrackingFrameCountClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for trackingFrameCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/trackingFrameCount-c.ivar
type trackingFrameCount struct {
	objectivec.Object
}

// trackingFrameCountFrom constructs a [trackingFrameCount] from an unsafe.Pointer.
func trackingFrameCountFrom(ptr unsafe.Pointer) trackingFrameCount {
	return trackingFrameCount{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for trackingFrameCount *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for trackingFrameCount */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for trackingFrameCount */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for trackingFrameCount */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for trackingFrameCount */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class trackingFrameCount */



