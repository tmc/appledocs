// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class replMode */


/* debug [class_header]: Header for replMode */
// The class instance for the [replMode] class.
var (
	ReplModeClass     _replModeClass
	ReplModeClassOnce sync.Once
)

func getreplModeClass() _replModeClass {
	ReplModeClassOnce.Do(func() {
		ReplModeClass = _replModeClass{objc.GetClass("replMode")}
	})
	return ReplModeClass
}

type _replModeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for replMode */
// An interface definition for the [replMode] class.
type IreplMode interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for replMode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for replMode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for replMode */
// Alloc allocates a new instance without initialization.
func (rc _replModeClass) Alloc() replMode {
	rv := objc.Send[replMode](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _replModeClass) New() replMode {
	rv := objc.Send[replMode](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ replMode) Init() replMode {
	rv := objc.Send[replMode](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ replMode) Autorelease() replMode {
	rv := objc.Send[replMode](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewreplMode creates a new replMode instance.
func NewreplMode() replMode {
	return getreplModeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for replMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/replMode
type replMode struct {
	objectivec.Object
}

// replModeFrom constructs a [replMode] from an unsafe.Pointer.
func replModeFrom(ptr unsafe.Pointer) replMode {
	return replMode{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for replMode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for replMode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for replMode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for replMode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for replMode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class replMode */



