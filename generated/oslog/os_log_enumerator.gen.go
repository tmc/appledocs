// Code generated from Apple documentation for OSLog. DO NOT EDIT.

package oslog

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class OSLogEnumerator */


/* debug [class_header]: Header for OSLogEnumerator */
// The class instance for the [OSLogEnumerator] class.
var (
	OSLogEnumeratorClass     _OSLogEnumeratorClass
	OSLogEnumeratorClassOnce sync.Once
)

func getOSLogEnumeratorClass() _OSLogEnumeratorClass {
	OSLogEnumeratorClassOnce.Do(func() {
		OSLogEnumeratorClass = _OSLogEnumeratorClass{objc.GetClass("OSLogEnumerator")}
	})
	return OSLogEnumeratorClass
}

type _OSLogEnumeratorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for OSLogEnumerator */
// An interface definition for the [OSLogEnumerator] class.
type IOSLogEnumerator interface {
	foundation.IEnumerator
	
/* debug [class_interface_properties]: Properties for OSLogEnumerator */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for OSLogEnumerator */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for OSLogEnumerator */
// Alloc allocates a new instance without initialization.
func (oc _OSLogEnumeratorClass) Alloc() OSLogEnumerator {
	rv := objc.Send[OSLogEnumerator](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _OSLogEnumeratorClass) New() OSLogEnumerator {
	rv := objc.Send[OSLogEnumerator](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OSLogEnumerator) Init() OSLogEnumerator {
	rv := objc.Send[OSLogEnumerator](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OSLogEnumerator) Autorelease() OSLogEnumerator {
	rv := objc.Send[OSLogEnumerator](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOSLogEnumerator creates a new OSLogEnumerator instance.
func NewOSLogEnumerator() OSLogEnumerator {
	return getOSLogEnumeratorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for OSLogEnumerator */
// An enumerator that can access and list log entries.


// An enumerator that can access and list log entries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogEnumerator
type OSLogEnumerator struct {
	foundation.Enumerator
}

// OSLogEnumeratorFrom constructs a [OSLogEnumerator] from an unsafe.Pointer.
//
// An enumerator that can access and list log entries.
func OSLogEnumeratorFrom(ptr unsafe.Pointer) OSLogEnumerator {
	return OSLogEnumerator{
		Enumerator: foundation.EnumeratorFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for OSLogEnumerator *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for OSLogEnumerator */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for OSLogEnumerator */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for OSLogEnumerator */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for OSLogEnumerator */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class OSLogEnumerator */



