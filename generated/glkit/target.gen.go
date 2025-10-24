// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class target */


/* debug [class_header]: Header for target */
// The class instance for the [target] class.
var (
	TargetClass     _targetClass
	TargetClassOnce sync.Once
)

func gettargetClass() _targetClass {
	TargetClassOnce.Do(func() {
		TargetClass = _targetClass{objc.GetClass("target")}
	})
	return TargetClass
}

type _targetClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for target */
// An interface definition for the [target] class.
type Itarget interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for target */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for target */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for target */
// Alloc allocates a new instance without initialization.
func (tc _targetClass) Alloc() target {
	rv := objc.Send[target](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _targetClass) New() target {
	rv := objc.Send[target](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ target) Init() target {
	rv := objc.Send[target](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ target) Autorelease() target {
	rv := objc.Send[target](t_.ID, objc.Sel("autorelease"))
	return rv
}

// Newtarget creates a new target instance.
func Newtarget() target {
	return gettargetClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for target */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfo/target-c.ivar
type target struct {
	objectivec.Object
}

// targetFrom constructs a [target] from an unsafe.Pointer.
func targetFrom(ptr unsafe.Pointer) target {
	return target{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for target *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for target */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for target */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for target */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for target */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class target */



