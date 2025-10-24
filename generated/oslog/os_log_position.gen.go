// Code generated from Apple documentation for OSLog. DO NOT EDIT.

package oslog

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class OSLogPosition */


/* debug [class_header]: Header for OSLogPosition */
// The class instance for the [OSLogPosition] class.
var (
	OSLogPositionClass     _OSLogPositionClass
	OSLogPositionClassOnce sync.Once
)

func getOSLogPositionClass() _OSLogPositionClass {
	OSLogPositionClassOnce.Do(func() {
		OSLogPositionClass = _OSLogPositionClass{objc.GetClass("OSLogPosition")}
	})
	return OSLogPositionClass
}

type _OSLogPositionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for OSLogPosition */
// An interface definition for the [OSLogPosition] class.
type IOSLogPosition interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for OSLogPosition */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for OSLogPosition */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for OSLogPosition */
// Alloc allocates a new instance without initialization.
func (oc _OSLogPositionClass) Alloc() OSLogPosition {
	rv := objc.Send[OSLogPosition](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _OSLogPositionClass) New() OSLogPosition {
	rv := objc.Send[OSLogPosition](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OSLogPosition) Init() OSLogPosition {
	rv := objc.Send[OSLogPosition](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OSLogPosition) Autorelease() OSLogPosition {
	rv := objc.Send[OSLogPosition](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOSLogPosition creates a new OSLogPosition instance.
func NewOSLogPosition() OSLogPosition {
	return getOSLogPositionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for OSLogPosition */
// A representation of a point in a sequence of entries in the unified logging system.
//
// Generate positions with instance methods and use them to view entries from a particular starting point.


// A representation of a point in a sequence of entries in the unified logging system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogPosition
type OSLogPosition struct {
	objectivec.Object
}

// OSLogPositionFrom constructs a [OSLogPosition] from an unsafe.Pointer.
//
// A representation of a point in a sequence of entries in the unified logging system.
func OSLogPositionFrom(ptr unsafe.Pointer) OSLogPosition {
	return OSLogPosition{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for OSLogPosition *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for OSLogPosition */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for OSLogPosition */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for OSLogPosition */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for OSLogPosition */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class OSLogPosition */



