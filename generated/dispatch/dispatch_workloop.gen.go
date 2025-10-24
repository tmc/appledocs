// Code generated from Apple documentation for Dispatch. DO NOT EDIT.

package dispatch

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class DispatchWorkloop */


/* debug [class_header]: Header for DispatchWorkloop */
// The class instance for the [DispatchWorkloop] class.
var (
	DispatchWorkloopClass     _DispatchWorkloopClass
	DispatchWorkloopClassOnce sync.Once
)

func getDispatchWorkloopClass() _DispatchWorkloopClass {
	DispatchWorkloopClassOnce.Do(func() {
		DispatchWorkloopClass = _DispatchWorkloopClass{objc.GetClass("DispatchWorkloop")}
	})
	return DispatchWorkloopClass
}

type _DispatchWorkloopClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DispatchWorkloop */
// An interface definition for the [DispatchWorkloop] class.
type IDispatchWorkloop interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for DispatchWorkloop */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DispatchWorkloop */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DispatchWorkloop */
// Alloc allocates a new instance without initialization.
func (dc _DispatchWorkloopClass) Alloc() DispatchWorkloop {
	rv := objc.Send[DispatchWorkloop](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DispatchWorkloopClass) New() DispatchWorkloop {
	rv := objc.Send[DispatchWorkloop](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DispatchWorkloop) Init() DispatchWorkloop {
	rv := objc.Send[DispatchWorkloop](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DispatchWorkloop) Autorelease() DispatchWorkloop {
	rv := objc.Send[DispatchWorkloop](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDispatchWorkloop creates a new DispatchWorkloop instance.
func NewDispatchWorkloop() DispatchWorkloop {
	return getDispatchWorkloopClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DispatchWorkloop */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/DispatchWorkloop
type DispatchWorkloop struct {
	objectivec.Object
}

// DispatchWorkloopFrom constructs a [DispatchWorkloop] from an unsafe.Pointer.
func DispatchWorkloopFrom(ptr unsafe.Pointer) DispatchWorkloop {
	return DispatchWorkloop{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DispatchWorkloop *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DispatchWorkloop */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DispatchWorkloop */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DispatchWorkloop */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DispatchWorkloop */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DispatchWorkloop */



