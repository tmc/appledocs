// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CMLogItem */


/* debug [class_header]: Header for CMLogItem */
// The class instance for the [LogItem] class.
var (
	LogItemClass     _LogItemClass
	LogItemClassOnce sync.Once
)

func getLogItemClass() _LogItemClass {
	LogItemClassOnce.Do(func() {
		LogItemClass = _LogItemClass{objc.GetClass("CMLogItem")}
	})
	return LogItemClass
}

type _LogItemClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for LogItem */
// An interface definition for the [LogItem] class.
type ILogItem interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for LogItem */
	// properties:
	Timestamp() float64
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for LogItem */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for LogItem */
// Alloc allocates a new instance without initialization.
func (lc _LogItemClass) Alloc() LogItem {
	rv := objc.Send[LogItem](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (lc _LogItemClass) New() LogItem {
	rv := objc.Send[LogItem](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LogItem) Init() LogItem {
	rv := objc.Send[LogItem](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LogItem) Autorelease() LogItem {
	rv := objc.Send[LogItem](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLogItem creates a new LogItem instance.
func NewLogItem() LogItem {
	return getLogItemClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for LogItem */
// The base class for all motion-related data objects.
//
// The class defines a read-only property that records the time a motion-event measurement was taken.


// The base class for all motion-related data objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMLogItem
type LogItem struct {
	objectivec.Object
}

// LogItemFrom constructs a [LogItem] from an unsafe.Pointer.
//
// The base class for all motion-related data objects.
func LogItemFrom(ptr unsafe.Pointer) LogItem {
	return LogItem{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for LogItem *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for LogItem */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for LogItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for LogItem */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for LogItem */

// The time when the logged item is valid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMLogItem/timestamp
func (l_ LogItem) Timestamp() float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("timestamp"))
	return rv
}/* debug [instance_properties/getter]: timestamp */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CMLogItem */



