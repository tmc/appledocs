// Code generated from Apple documentation for Dispatch. DO NOT EDIT.

package dispatch

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class OS_object */


/* debug [class_header]: Header for OS_object */
// The class instance for the [OS_object] class.
var (
	OS_objectClass     _OS_objectClass
	OS_objectClassOnce sync.Once
)

func getOS_objectClass() _OS_objectClass {
	OS_objectClassOnce.Do(func() {
		OS_objectClass = _OS_objectClass{objc.GetClass("OS_object")}
	})
	return OS_objectClass
}

type _OS_objectClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for OS_object */
// An interface definition for the [OS_object] class.
type IOS_object interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for OS_object */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for OS_object */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for OS_object */
// Alloc allocates a new instance without initialization.
func (oc _OS_objectClass) Alloc() OS_object {
	rv := objc.Send[OS_object](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _OS_objectClass) New() OS_object {
	rv := objc.Send[OS_object](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OS_object) Init() OS_object {
	rv := objc.Send[OS_object](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OS_object) Autorelease() OS_object {
	rv := objc.Send[OS_object](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOS_object creates a new OS_object instance.
func NewOS_object() OS_object {
	return getOS_objectClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for OS_object */
// A parent class referenced by other Dispatch classes.


// A parent class referenced by other Dispatch classes. [Full Topic]
type OS_object struct {
	objectivec.Object
}

// OS_objectFrom constructs a [OS_object] from an unsafe.Pointer.
//
// A parent class referenced by other Dispatch classes.
func OS_objectFrom(ptr unsafe.Pointer) OS_object {
	return OS_object{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for OS_object *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for OS_object */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for OS_object */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for OS_object */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for OS_object */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class OS_object */



