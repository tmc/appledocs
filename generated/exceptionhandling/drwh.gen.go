// Code generated from Apple documentation for ExceptionHandling. DO NOT EDIT.

package exceptionhandling

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class drwh */


/* debug [class_header]: Header for drwh */
// The class instance for the [drwh] class.
var (
	DrwhClass     _drwhClass
	DrwhClassOnce sync.Once
)

func getdrwhClass() _drwhClass {
	DrwhClassOnce.Do(func() {
		DrwhClass = _drwhClass{objc.GetClass("drwh")}
	})
	return DrwhClass
}

type _drwhClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for drwh */
// An interface definition for the [drwh] class.
type Idrwh interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for drwh */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for drwh */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for drwh */
// Alloc allocates a new instance without initialization.
func (dc _drwhClass) Alloc() drwh {
	rv := objc.Send[drwh](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _drwhClass) New() drwh {
	rv := objc.Send[drwh](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ drwh) Init() drwh {
	rv := objc.Send[drwh](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ drwh) Autorelease() drwh {
	rv := objc.Send[drwh](d_.ID, objc.Sel("autorelease"))
	return rv
}

// Newdrwh creates a new drwh instance.
func Newdrwh() drwh {
	return getdrwhClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for drwh */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExceptionHandling/NSExceptionHandler/struct_(unnamed)/drwh
type drwh struct {
	objectivec.Object
}

// drwhFrom constructs a [drwh] from an unsafe.Pointer.
func drwhFrom(ptr unsafe.Pointer) drwh {
	return drwh{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for drwh *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for drwh */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for drwh */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for drwh */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for drwh */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class drwh */



