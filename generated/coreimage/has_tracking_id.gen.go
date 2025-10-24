// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class hasTrackingID */


/* debug [class_header]: Header for hasTrackingID */
// The class instance for the [hasTrackingID] class.
var (
	HasTrackingIDClass     _hasTrackingIDClass
	HasTrackingIDClassOnce sync.Once
)

func gethasTrackingIDClass() _hasTrackingIDClass {
	HasTrackingIDClassOnce.Do(func() {
		HasTrackingIDClass = _hasTrackingIDClass{objc.GetClass("hasTrackingID")}
	})
	return HasTrackingIDClass
}

type _hasTrackingIDClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for hasTrackingID */
// An interface definition for the [hasTrackingID] class.
type IhasTrackingID interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for hasTrackingID */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for hasTrackingID */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for hasTrackingID */
// Alloc allocates a new instance without initialization.
func (hc _hasTrackingIDClass) Alloc() hasTrackingID {
	rv := objc.Send[hasTrackingID](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _hasTrackingIDClass) New() hasTrackingID {
	rv := objc.Send[hasTrackingID](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ hasTrackingID) Init() hasTrackingID {
	rv := objc.Send[hasTrackingID](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ hasTrackingID) Autorelease() hasTrackingID {
	rv := objc.Send[hasTrackingID](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewhasTrackingID creates a new hasTrackingID instance.
func NewhasTrackingID() hasTrackingID {
	return gethasTrackingIDClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for hasTrackingID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasTrackingID-c.ivar
type hasTrackingID struct {
	objectivec.Object
}

// hasTrackingIDFrom constructs a [hasTrackingID] from an unsafe.Pointer.
func hasTrackingIDFrom(ptr unsafe.Pointer) hasTrackingID {
	return hasTrackingID{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for hasTrackingID *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for hasTrackingID */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for hasTrackingID */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for hasTrackingID */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for hasTrackingID */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class hasTrackingID */



