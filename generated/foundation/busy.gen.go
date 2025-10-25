// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class busy */


/* debug [class_header]: Header for busy */
// The class instance for the [busy] class.
var (
	BusyClass     _busyClass
	BusyClassOnce sync.Once
)

func getbusyClass() _busyClass {
	BusyClassOnce.Do(func() {
		BusyClass = _busyClass{objc.GetClass("busy")}
	})
	return BusyClass
}

type _busyClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for busy */
// An interface definition for the [busy] class.
type Ibusy interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for busy */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for busy */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for busy */
// Alloc allocates a new instance without initialization.
func (bc _busyClass) Alloc() busy {
	rv := objc.Send[busy](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _busyClass) New() busy {
	rv := objc.Send[busy](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ busy) Init() busy {
	rv := objc.Send[busy](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ busy) Autorelease() busy {
	rv := objc.Send[busy](b_.ID, objc.Sel("autorelease"))
	return rv
}

// Newbusy creates a new busy instance.
func Newbusy() busy {
	return getbusyClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for busy */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/busy
type busy struct {
	objectivec.Object
}

// busyFrom constructs a [busy] from an unsafe.Pointer.
func busyFrom(ptr unsafe.Pointer) busy {
	return busy{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for busy *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for busy */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for busy */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for busy */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for busy */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class busy */



