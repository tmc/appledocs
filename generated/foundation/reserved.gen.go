// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class reserved */


/* debug [class_header]: Header for reserved */
// The class instance for the [reserved] class.
var (
	ReservedClass     _reservedClass
	ReservedClassOnce sync.Once
)

func getreservedClass() _reservedClass {
	ReservedClassOnce.Do(func() {
		ReservedClass = _reservedClass{objc.GetClass("reserved")}
	})
	return ReservedClass
}

type _reservedClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for reserved */
// An interface definition for the [reserved] class.
type Ireserved interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for reserved */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for reserved */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for reserved */
// Alloc allocates a new instance without initialization.
func (rc _reservedClass) Alloc() reserved {
	rv := objc.Send[reserved](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _reservedClass) New() reserved {
	rv := objc.Send[reserved](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ reserved) Init() reserved {
	rv := objc.Send[reserved](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ reserved) Autorelease() reserved {
	rv := objc.Send[reserved](r_.ID, objc.Sel("autorelease"))
	return rv
}

// Newreserved creates a new reserved instance.
func Newreserved() reserved {
	return getreservedClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for reserved */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArchiver/reserved
type reserved struct {
	objectivec.Object
}

// reservedFrom constructs a [reserved] from an unsafe.Pointer.
func reservedFrom(ptr unsafe.Pointer) reserved {
	return reserved{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for reserved *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for reserved */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for reserved */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for reserved */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for reserved */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class reserved */



