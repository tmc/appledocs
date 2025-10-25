// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class refCount */


/* debug [class_header]: Header for refCount */
// The class instance for the [refCount] class.
var (
	RefCountClass     _refCountClass
	RefCountClassOnce sync.Once
)

func getrefCountClass() _refCountClass {
	RefCountClassOnce.Do(func() {
		RefCountClass = _refCountClass{objc.GetClass("refCount")}
	})
	return RefCountClass
}

type _refCountClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for refCount */
// An interface definition for the [refCount] class.
type IrefCount interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for refCount */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for refCount */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for refCount */
// Alloc allocates a new instance without initialization.
func (rc _refCountClass) Alloc() refCount {
	rv := objc.Send[refCount](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _refCountClass) New() refCount {
	rv := objc.Send[refCount](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ refCount) Init() refCount {
	rv := objc.Send[refCount](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ refCount) Autorelease() refCount {
	rv := objc.Send[refCount](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewrefCount creates a new refCount instance.
func NewrefCount() refCount {
	return getrefCountClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for refCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendarDate/refCount
type refCount struct {
	objectivec.Object
}

// refCountFrom constructs a [refCount] from an unsafe.Pointer.
func refCountFrom(ptr unsafe.Pointer) refCount {
	return refCount{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for refCount *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for refCount */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for refCount */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for refCount */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for refCount */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class refCount */



