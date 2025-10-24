// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNDArrayMatrixMultiplication */


/* debug [class_header]: Header for MPSNDArrayMatrixMultiplication */
// The class instance for the [NDArrayMatrixMultiplication] class.
var (
	NDArrayMatrixMultiplicationClass     _NDArrayMatrixMultiplicationClass
	NDArrayMatrixMultiplicationClassOnce sync.Once
)

func getNDArrayMatrixMultiplicationClass() _NDArrayMatrixMultiplicationClass {
	NDArrayMatrixMultiplicationClassOnce.Do(func() {
		NDArrayMatrixMultiplicationClass = _NDArrayMatrixMultiplicationClass{objc.GetClass("MPSNDArrayMatrixMultiplication")}
	})
	return NDArrayMatrixMultiplicationClass
}

type _NDArrayMatrixMultiplicationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NDArrayMatrixMultiplication */
// An interface definition for the [NDArrayMatrixMultiplication] class.
type INDArrayMatrixMultiplication interface {
	INDArrayMultiaryKernel
	
/* debug [class_interface_properties]: Properties for NDArrayMatrixMultiplication */
	// properties:
	Alpha() objectivec.IObject
	SetAlpha(value objectivec.IObject)
	Beta() objectivec.IObject
	SetBeta(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NDArrayMatrixMultiplication */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NDArrayMatrixMultiplication */
// Alloc allocates a new instance without initialization.
func (nc _NDArrayMatrixMultiplicationClass) Alloc() NDArrayMatrixMultiplication {
	rv := objc.Send[NDArrayMatrixMultiplication](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NDArrayMatrixMultiplicationClass) New() NDArrayMatrixMultiplication {
	rv := objc.Send[NDArrayMatrixMultiplication](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayMatrixMultiplication) Init() NDArrayMatrixMultiplication {
	rv := objc.Send[NDArrayMatrixMultiplication](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayMatrixMultiplication) Autorelease() NDArrayMatrixMultiplication {
	rv := objc.Send[NDArrayMatrixMultiplication](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayMatrixMultiplication creates a new NDArrayMatrixMultiplication instance.
func NewNDArrayMatrixMultiplication() NDArrayMatrixMultiplication {
	return getNDArrayMatrixMultiplicationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NDArrayMatrixMultiplication */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMatrixMultiplication
type NDArrayMatrixMultiplication struct {
	NDArrayMultiaryKernel
}

// NDArrayMatrixMultiplicationFrom constructs a [NDArrayMatrixMultiplication] from an unsafe.Pointer.
func NDArrayMatrixMultiplicationFrom(ptr unsafe.Pointer) NDArrayMatrixMultiplication {
	return NDArrayMatrixMultiplication{
		NDArrayMultiaryKernel: NDArrayMultiaryKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NDArrayMatrixMultiplication *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NDArrayMatrixMultiplication */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NDArrayMatrixMultiplication */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NDArrayMatrixMultiplication */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NDArrayMatrixMultiplication */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymatrixmultiplication/3131760-alpha
func (n_ NDArrayMatrixMultiplication) Alpha() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("alpha"))
	return rv
}/* debug [instance_properties/getter]: alpha */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymatrixmultiplication/3131760-alpha
func (n_ NDArrayMatrixMultiplication) SetAlpha(value objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAlpha:"), value)
}/* debug [instance_properties/setter]: alpha */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymatrixmultiplication/3131761-beta
func (n_ NDArrayMatrixMultiplication) Beta() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("beta"))
	return rv
}/* debug [instance_properties/getter]: beta */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymatrixmultiplication/3131761-beta
func (n_ NDArrayMatrixMultiplication) SetBeta(value objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setBeta:"), value)
}/* debug [instance_properties/setter]: beta */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNDArrayMatrixMultiplication */



