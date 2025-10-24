// Code generated from Apple documentation for ExceptionHandling. DO NOT EDIT.

package exceptionhandling

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class drwl */


/* debug [class_header]: Header for drwl */
// The class instance for the [drwl] class.
var (
	DrwlClass     _drwlClass
	DrwlClassOnce sync.Once
)

func getdrwlClass() _drwlClass {
	DrwlClassOnce.Do(func() {
		DrwlClass = _drwlClass{objc.GetClass("drwl")}
	})
	return DrwlClass
}

type _drwlClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for drwl */
// An interface definition for the [drwl] class.
type Idrwl interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for drwl */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for drwl */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for drwl */
// Alloc allocates a new instance without initialization.
func (dc _drwlClass) Alloc() drwl {
	rv := objc.Send[drwl](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _drwlClass) New() drwl {
	rv := objc.Send[drwl](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ drwl) Init() drwl {
	rv := objc.Send[drwl](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ drwl) Autorelease() drwl {
	rv := objc.Send[drwl](d_.ID, objc.Sel("autorelease"))
	return rv
}

// Newdrwl creates a new drwl instance.
func Newdrwl() drwl {
	return getdrwlClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for drwl */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExceptionHandling/NSExceptionHandler/struct_(unnamed)/drwl
type drwl struct {
	objectivec.Object
}

// drwlFrom constructs a [drwl] from an unsafe.Pointer.
func drwlFrom(ptr unsafe.Pointer) drwl {
	return drwl{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for drwl *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for drwl */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for drwl */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for drwl */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for drwl */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class drwl */



