// Code generated from Apple documentation for OSLog. DO NOT EDIT.

package oslog

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class OSLogEntrySignpost */


/* debug [class_header]: Header for OSLogEntrySignpost */
// The class instance for the [OSLogEntrySignpost] class.
var (
	OSLogEntrySignpostClass     _OSLogEntrySignpostClass
	OSLogEntrySignpostClassOnce sync.Once
)

func getOSLogEntrySignpostClass() _OSLogEntrySignpostClass {
	OSLogEntrySignpostClassOnce.Do(func() {
		OSLogEntrySignpostClass = _OSLogEntrySignpostClass{objc.GetClass("OSLogEntrySignpost")}
	})
	return OSLogEntrySignpostClass
}

type _OSLogEntrySignpostClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for OSLogEntrySignpost */
// An interface definition for the [OSLogEntrySignpost] class.
type IOSLogEntrySignpost interface {
	IOSLogEntry
	
/* debug [class_interface_properties]: Properties for OSLogEntrySignpost */
	// properties:
	SignpostIdentifier() unsafe.Pointer
	SignpostName() objc.IObject /* cross-framework: NSString */
	SignpostType() OSLogEntrySignpostType
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for OSLogEntrySignpost */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for OSLogEntrySignpost */
// Alloc allocates a new instance without initialization.
func (oc _OSLogEntrySignpostClass) Alloc() OSLogEntrySignpost {
	rv := objc.Send[OSLogEntrySignpost](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _OSLogEntrySignpostClass) New() OSLogEntrySignpost {
	rv := objc.Send[OSLogEntrySignpost](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OSLogEntrySignpost) Init() OSLogEntrySignpost {
	rv := objc.Send[OSLogEntrySignpost](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OSLogEntrySignpost) Autorelease() OSLogEntrySignpost {
	rv := objc.Send[OSLogEntrySignpost](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOSLogEntrySignpost creates a new OSLogEntrySignpost instance.
func NewOSLogEntrySignpost() OSLogEntrySignpost {
	return getOSLogEntrySignpostClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for OSLogEntrySignpost */
// An entry containing a signpost.
//
// These entries are created by the os_signpost API. To learn more about signposts and how to create a signpost entry, see and .


// An entry containing a signpost.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogEntrySignpost
type OSLogEntrySignpost struct {
	OSLogEntry
}

// OSLogEntrySignpostFrom constructs a [OSLogEntrySignpost] from an unsafe.Pointer.
//
// An entry containing a signpost.
func OSLogEntrySignpostFrom(ptr unsafe.Pointer) OSLogEntrySignpost {
	return OSLogEntrySignpost{
		OSLogEntry: OSLogEntryFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for OSLogEntrySignpost *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for OSLogEntrySignpost */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for OSLogEntrySignpost */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for OSLogEntrySignpost */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for OSLogEntrySignpost */

// The signpost’s identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogEntrySignpost/signpostIdentifier
func (o_ OSLogEntrySignpost) SignpostIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("signpostIdentifier"))
	return rv
}/* debug [instance_properties/getter]: signpostIdentifier */


// The signpost’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogEntrySignpost/signpostName
func (o_ OSLogEntrySignpost) SignpostName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("signpostName"))
	return rv
}/* debug [instance_properties/getter]: signpostName */


// The signpost’s type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogEntrySignpost/signpostType-swift.property
func (o_ OSLogEntrySignpost) SignpostType() OSLogEntrySignpostType {
	rv := objc.Send[OSLogEntrySignpostType](o_.ID, objc.Sel("signpostType"))
	return rv
}/* debug [instance_properties/getter]: signpostType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class OSLogEntrySignpost */



