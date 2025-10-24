// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNAdd */


/* debug [class_header]: Header for MPSCNNAdd */
// The class instance for the [CNNAdd] class.
var (
	CNNAddClass     _CNNAddClass
	CNNAddClassOnce sync.Once
)

func getCNNAddClass() _CNNAddClass {
	CNNAddClassOnce.Do(func() {
		CNNAddClass = _CNNAddClass{objc.GetClass("MPSCNNAdd")}
	})
	return CNNAddClass
}

type _CNNAddClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNAdd */
// An interface definition for the [CNNAdd] class.
type ICNNAdd interface {
	ICNNArithmetic
	
/* debug [class_interface_properties]: Properties for CNNAdd */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNAdd */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNAdd */
// Alloc allocates a new instance without initialization.
func (cc _CNNAddClass) Alloc() CNNAdd {
	rv := objc.Send[CNNAdd](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNAddClass) New() CNNAdd {
	rv := objc.Send[CNNAdd](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNAdd) Init() CNNAdd {
	rv := objc.Send[CNNAdd](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNAdd) Autorelease() CNNAdd {
	rv := objc.Send[CNNAdd](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNAdd creates a new CNNAdd instance.
func NewCNNAdd() CNNAdd {
	return getCNNAddClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNAdd */
// An addition operator.


// An addition operator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNAdd
type CNNAdd struct {
	CNNArithmetic
}

// CNNAddFrom constructs a [CNNAdd] from an unsafe.Pointer.
//
// An addition operator.
func CNNAddFrom(ptr unsafe.Pointer) CNNAdd {
	return CNNAdd{
		CNNArithmetic: CNNArithmeticFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNAdd */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnadd/2942501-initwithdevice
func NewCNNAddWithDevice(device unsafe.Pointer) CNNAdd {
	instance := getCNNAddClass().Alloc()
	rv := objc.Send[CNNAdd](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNAddWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNAdd */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNAdd */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNAdd */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNAdd */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNAdd */


