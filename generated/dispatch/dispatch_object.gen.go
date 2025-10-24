// Code generated from Apple documentation for Dispatch. DO NOT EDIT.

package dispatch

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DispatchObject */


/* debug [class_header]: Header for DispatchObject */
// The class instance for the [DispatchObject] class.
var (
	DispatchObjectClass     _DispatchObjectClass
	DispatchObjectClassOnce sync.Once
)

func getDispatchObjectClass() _DispatchObjectClass {
	DispatchObjectClassOnce.Do(func() {
		DispatchObjectClass = _DispatchObjectClass{objc.GetClass("DispatchObject")}
	})
	return DispatchObjectClass
}

type _DispatchObjectClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DispatchObject */
// An interface definition for the [DispatchObject] class.
type IDispatchObject interface {
	IOS_object
	
/* debug [class_interface_properties]: Properties for DispatchObject */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DispatchObject */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DispatchObject */
// Alloc allocates a new instance without initialization.
func (dc _DispatchObjectClass) Alloc() DispatchObject {
	rv := objc.Send[DispatchObject](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DispatchObjectClass) New() DispatchObject {
	rv := objc.Send[DispatchObject](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DispatchObject) Init() DispatchObject {
	rv := objc.Send[DispatchObject](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DispatchObject) Autorelease() DispatchObject {
	rv := objc.Send[DispatchObject](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDispatchObject creates a new DispatchObject instance.
func NewDispatchObject() DispatchObject {
	return getDispatchObjectClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DispatchObject */
// The base class for most dispatch types.
//
// There are many types of dispatch objects, including , , and . The base dispatch object interfaces allow you to manage memory, pause and resume execution, define object context, log task data, and more.


// The base class for most dispatch types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/DispatchObject
type DispatchObject struct {
	OS_object
}

// DispatchObjectFrom constructs a [DispatchObject] from an unsafe.Pointer.
//
// The base class for most dispatch types.
func DispatchObjectFrom(ptr unsafe.Pointer) DispatchObject {
	return DispatchObject{
		OS_object: OS_objectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DispatchObject *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DispatchObject */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DispatchObject */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DispatchObject */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DispatchObject */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DispatchObject */



