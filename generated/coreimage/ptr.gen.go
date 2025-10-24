// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ptr */


/* debug [class_header]: Header for ptr */
// The class instance for the [ptr] class.
var (
	PtrClass     _ptrClass
	PtrClassOnce sync.Once
)

func getptrClass() _ptrClass {
	PtrClassOnce.Do(func() {
		PtrClass = _ptrClass{objc.GetClass("ptr")}
	})
	return PtrClass
}

type _ptrClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ptr */
// An interface definition for the [ptr] class.
type Iptr interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ptr */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ptr */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ptr */
// Alloc allocates a new instance without initialization.
func (pc _ptrClass) Alloc() ptr {
	rv := objc.Send[ptr](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _ptrClass) New() ptr {
	rv := objc.Send[ptr](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ ptr) Init() ptr {
	rv := objc.Send[ptr](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ ptr) Autorelease() ptr {
	rv := objc.Send[ptr](p_.ID, objc.Sel("autorelease"))
	return rv
}

// Newptr creates a new ptr instance.
func Newptr() ptr {
	return getptrClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ptr */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/union_(unnamed)/ptr
type ptr struct {
	objectivec.Object
}

// ptrFrom constructs a [ptr] from an unsafe.Pointer.
func ptrFrom(ptr unsafe.Pointer) ptr {
	return ptr{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ptr *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ptr */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ptr */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ptr */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ptr */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ptr */



