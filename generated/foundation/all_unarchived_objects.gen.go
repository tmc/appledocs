// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class allUnarchivedObjects */


/* debug [class_header]: Header for allUnarchivedObjects */
// The class instance for the [allUnarchivedObjects] class.
var (
	AllUnarchivedObjectsClass     _allUnarchivedObjectsClass
	AllUnarchivedObjectsClassOnce sync.Once
)

func getallUnarchivedObjectsClass() _allUnarchivedObjectsClass {
	AllUnarchivedObjectsClassOnce.Do(func() {
		AllUnarchivedObjectsClass = _allUnarchivedObjectsClass{objc.GetClass("allUnarchivedObjects")}
	})
	return AllUnarchivedObjectsClass
}

type _allUnarchivedObjectsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for allUnarchivedObjects */
// An interface definition for the [allUnarchivedObjects] class.
type IallUnarchivedObjects interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for allUnarchivedObjects */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for allUnarchivedObjects */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for allUnarchivedObjects */
// Alloc allocates a new instance without initialization.
func (ac _allUnarchivedObjectsClass) Alloc() allUnarchivedObjects {
	rv := objc.Send[allUnarchivedObjects](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _allUnarchivedObjectsClass) New() allUnarchivedObjects {
	rv := objc.Send[allUnarchivedObjects](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ allUnarchivedObjects) Init() allUnarchivedObjects {
	rv := objc.Send[allUnarchivedObjects](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ allUnarchivedObjects) Autorelease() allUnarchivedObjects {
	rv := objc.Send[allUnarchivedObjects](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewallUnarchivedObjects creates a new allUnarchivedObjects instance.
func NewallUnarchivedObjects() allUnarchivedObjects {
	return getallUnarchivedObjectsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for allUnarchivedObjects */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnarchiver/allUnarchivedObjects
type allUnarchivedObjects struct {
	objectivec.Object
}

// allUnarchivedObjectsFrom constructs a [allUnarchivedObjects] from an unsafe.Pointer.
func allUnarchivedObjectsFrom(ptr unsafe.Pointer) allUnarchivedObjects {
	return allUnarchivedObjects{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for allUnarchivedObjects *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for allUnarchivedObjects */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for allUnarchivedObjects */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for allUnarchivedObjects */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for allUnarchivedObjects */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class allUnarchivedObjects */



