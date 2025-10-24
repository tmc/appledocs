// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class runLoops */


/* debug [class_header]: Header for runLoops */
// The class instance for the [runLoops] class.
var (
	RunLoopsClass     _runLoopsClass
	RunLoopsClassOnce sync.Once
)

func getrunLoopsClass() _runLoopsClass {
	RunLoopsClassOnce.Do(func() {
		RunLoopsClass = _runLoopsClass{objc.GetClass("runLoops")}
	})
	return RunLoopsClass
}

type _runLoopsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for runLoops */
// An interface definition for the [runLoops] class.
type IrunLoops interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for runLoops */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for runLoops */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for runLoops */
// Alloc allocates a new instance without initialization.
func (rc _runLoopsClass) Alloc() runLoops {
	rv := objc.Send[runLoops](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _runLoopsClass) New() runLoops {
	rv := objc.Send[runLoops](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ runLoops) Init() runLoops {
	rv := objc.Send[runLoops](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ runLoops) Autorelease() runLoops {
	rv := objc.Send[runLoops](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewrunLoops creates a new runLoops instance.
func NewrunLoops() runLoops {
	return getrunLoopsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for runLoops */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/runLoops
type runLoops struct {
	objectivec.Object
}

// runLoopsFrom constructs a [runLoops] from an unsafe.Pointer.
func runLoopsFrom(ptr unsafe.Pointer) runLoops {
	return runLoops{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for runLoops *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for runLoops */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for runLoops */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for runLoops */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for runLoops */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class runLoops */



