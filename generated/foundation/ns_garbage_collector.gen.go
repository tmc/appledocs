// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSGarbageCollector */


/* debug [class_header]: Header for NSGarbageCollector */
// The class instance for the [GarbageCollector] class.
var (
	GarbageCollectorClass     _GarbageCollectorClass
	GarbageCollectorClassOnce sync.Once
)

func getGarbageCollectorClass() _GarbageCollectorClass {
	GarbageCollectorClassOnce.Do(func() {
		GarbageCollectorClass = _GarbageCollectorClass{objc.GetClass("NSGarbageCollector")}
	})
	return GarbageCollectorClass
}

type _GarbageCollectorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GarbageCollector */
// An interface definition for the [GarbageCollector] class.
type IGarbageCollector interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for GarbageCollector */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GarbageCollector */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GarbageCollector */
// Alloc allocates a new instance without initialization.
func (gc _GarbageCollectorClass) Alloc() GarbageCollector {
	rv := objc.Send[GarbageCollector](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GarbageCollectorClass) New() GarbageCollector {
	rv := objc.Send[GarbageCollector](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GarbageCollector) Init() GarbageCollector {
	rv := objc.Send[GarbageCollector](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GarbageCollector) Autorelease() GarbageCollector {
	rv := objc.Send[GarbageCollector](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGarbageCollector creates a new GarbageCollector instance.
func NewGarbageCollector() GarbageCollector {
	return getGarbageCollectorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GarbageCollector */
// A convenient interface to the garbage collection system.
//
// Cocoa’s garbage collector is a conservative generational garbage collector. It uses “write-barriers” to detect cross generational stores of pointers so that “young” objects can be collected quickly. You enable garbage collection (GC) by using the option. This switch causes the generation of the write-barrier assignment primitives. You must use this option on your main application file , including frameworks and bundles. Bundles are ignored if they are not GC-capable. The collector determines what is garbage by recursively examining all nodes starting with globals, possible nodes referenced from the thread stacks, and all nodes marked as having “external” references. Nodes not reached by this search are deemed garbage. Weak references to garbage nodes are then cleared. Garbage nodes that are objects are sent (in an arbitrary order) a message, and after all messages have been sent their memory is recovered. It is a runtime error (referred to as “resurrection”) to store a object being finalized into one that is not. For more details, see Implementing a finalize Method in Garbage Collection Programming Guide. You can request collection from any thread (see and ).


// A convenient interface to the garbage collection system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGarbageCollector
type GarbageCollector struct {
	objectivec.Object
}

// GarbageCollectorFrom constructs a [GarbageCollector] from an unsafe.Pointer.
//
// A convenient interface to the garbage collection system.
func GarbageCollectorFrom(ptr unsafe.Pointer) GarbageCollector {
	return GarbageCollector{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GarbageCollector *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GarbageCollector */

// Returns the default garbage collector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGarbageCollector/defaultCollector
func (gc _GarbageCollectorClass) DefaultCollector() objc.ID {
	rv := objc.Send[objc.ID](objc.ID(gc.class), objc.Sel("defaultCollector"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DefaultCollector) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GarbageCollector */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GarbageCollector */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GarbageCollector */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSGarbageCollector */



