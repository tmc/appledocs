// Code generated from Apple documentation for OSLog. DO NOT EDIT.

package oslog

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class OSLogEntryBoundary */


/* debug [class_header]: Header for OSLogEntryBoundary */
// The class instance for the [OSLogEntryBoundary] class.
var (
	OSLogEntryBoundaryClass     _OSLogEntryBoundaryClass
	OSLogEntryBoundaryClassOnce sync.Once
)

func getOSLogEntryBoundaryClass() _OSLogEntryBoundaryClass {
	OSLogEntryBoundaryClassOnce.Do(func() {
		OSLogEntryBoundaryClass = _OSLogEntryBoundaryClass{objc.GetClass("OSLogEntryBoundary")}
	})
	return OSLogEntryBoundaryClass
}

type _OSLogEntryBoundaryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for OSLogEntryBoundary */
// An interface definition for the [OSLogEntryBoundary] class.
type IOSLogEntryBoundary interface {
	IOSLogEntry
	
/* debug [class_interface_properties]: Properties for OSLogEntryBoundary */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for OSLogEntryBoundary */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for OSLogEntryBoundary */
// Alloc allocates a new instance without initialization.
func (oc _OSLogEntryBoundaryClass) Alloc() OSLogEntryBoundary {
	rv := objc.Send[OSLogEntryBoundary](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _OSLogEntryBoundaryClass) New() OSLogEntryBoundary {
	rv := objc.Send[OSLogEntryBoundary](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OSLogEntryBoundary) Init() OSLogEntryBoundary {
	rv := objc.Send[OSLogEntryBoundary](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OSLogEntryBoundary) Autorelease() OSLogEntryBoundary {
	rv := objc.Send[OSLogEntryBoundary](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOSLogEntryBoundary creates a new OSLogEntryBoundary instance.
func NewOSLogEntryBoundary() OSLogEntryBoundary {
	return getOSLogEntryBoundaryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for OSLogEntryBoundary */
// The metadata that partitions sequences of other entries.


// The metadata that partitions sequences of other entries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogEntryBoundary
type OSLogEntryBoundary struct {
	OSLogEntry
}

// OSLogEntryBoundaryFrom constructs a [OSLogEntryBoundary] from an unsafe.Pointer.
//
// The metadata that partitions sequences of other entries.
func OSLogEntryBoundaryFrom(ptr unsafe.Pointer) OSLogEntryBoundary {
	return OSLogEntryBoundary{
		OSLogEntry: OSLogEntryFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for OSLogEntryBoundary *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for OSLogEntryBoundary */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for OSLogEntryBoundary */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for OSLogEntryBoundary */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for OSLogEntryBoundary */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class OSLogEntryBoundary */



