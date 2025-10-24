// Code generated from Apple documentation for ExceptionHandling. DO NOT EDIT.

package exceptionhandling

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class reserved1 */


/* debug [class_header]: Header for reserved1 */
// The class instance for the [reserved1] class.
var (
	Reserved1Class     _reserved1Class
	Reserved1ClassOnce sync.Once
)

func getreserved1Class() _reserved1Class {
	Reserved1ClassOnce.Do(func() {
		Reserved1Class = _reserved1Class{objc.GetClass("reserved1")}
	})
	return Reserved1Class
}

type _reserved1Class struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for reserved1 */
// An interface definition for the [reserved1] class.
type Ireserved1 interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for reserved1 */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for reserved1 */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for reserved1 */
// Alloc allocates a new instance without initialization.
func (rc _reserved1Class) Alloc() reserved1 {
	rv := objc.Send[reserved1](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _reserved1Class) New() reserved1 {
	rv := objc.Send[reserved1](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ reserved1) Init() reserved1 {
	rv := objc.Send[reserved1](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ reserved1) Autorelease() reserved1 {
	rv := objc.Send[reserved1](r_.ID, objc.Sel("autorelease"))
	return rv
}

// Newreserved1 creates a new reserved1 instance.
func Newreserved1() reserved1 {
	return getreserved1Class().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for reserved1 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExceptionHandling/NSExceptionHandler/struct_(unnamed)/reserved1
type reserved1 struct {
	objectivec.Object
}

// reserved1From constructs a [reserved1] from an unsafe.Pointer.
func reserved1From(ptr unsafe.Pointer) reserved1 {
	return reserved1{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for reserved1 *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for reserved1 */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for reserved1 */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for reserved1 */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for reserved1 */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class reserved1 */






