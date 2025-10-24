// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNDArrayQuantizedMatrixMultiplication */


/* debug [class_header]: Header for MPSNDArrayQuantizedMatrixMultiplication */
// The class instance for the [NDArrayQuantizedMatrixMultiplication] class.
var (
	NDArrayQuantizedMatrixMultiplicationClass     _NDArrayQuantizedMatrixMultiplicationClass
	NDArrayQuantizedMatrixMultiplicationClassOnce sync.Once
)

func getNDArrayQuantizedMatrixMultiplicationClass() _NDArrayQuantizedMatrixMultiplicationClass {
	NDArrayQuantizedMatrixMultiplicationClassOnce.Do(func() {
		NDArrayQuantizedMatrixMultiplicationClass = _NDArrayQuantizedMatrixMultiplicationClass{objc.GetClass("MPSNDArrayQuantizedMatrixMultiplication")}
	})
	return NDArrayQuantizedMatrixMultiplicationClass
}

type _NDArrayQuantizedMatrixMultiplicationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NDArrayQuantizedMatrixMultiplication */
// An interface definition for the [NDArrayQuantizedMatrixMultiplication] class.
type INDArrayQuantizedMatrixMultiplication interface {
	INDArrayMatrixMultiplication
	
/* debug [class_interface_properties]: Properties for NDArrayQuantizedMatrixMultiplication */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NDArrayQuantizedMatrixMultiplication */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NDArrayQuantizedMatrixMultiplication */
// Alloc allocates a new instance without initialization.
func (nc _NDArrayQuantizedMatrixMultiplicationClass) Alloc() NDArrayQuantizedMatrixMultiplication {
	rv := objc.Send[NDArrayQuantizedMatrixMultiplication](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NDArrayQuantizedMatrixMultiplicationClass) New() NDArrayQuantizedMatrixMultiplication {
	rv := objc.Send[NDArrayQuantizedMatrixMultiplication](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayQuantizedMatrixMultiplication) Init() NDArrayQuantizedMatrixMultiplication {
	rv := objc.Send[NDArrayQuantizedMatrixMultiplication](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayQuantizedMatrixMultiplication) Autorelease() NDArrayQuantizedMatrixMultiplication {
	rv := objc.Send[NDArrayQuantizedMatrixMultiplication](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayQuantizedMatrixMultiplication creates a new NDArrayQuantizedMatrixMultiplication instance.
func NewNDArrayQuantizedMatrixMultiplication() NDArrayQuantizedMatrixMultiplication {
	return getNDArrayQuantizedMatrixMultiplicationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NDArrayQuantizedMatrixMultiplication */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayQuantizedMatrixMultiplication
type NDArrayQuantizedMatrixMultiplication struct {
	NDArrayMatrixMultiplication
}

// NDArrayQuantizedMatrixMultiplicationFrom constructs a [NDArrayQuantizedMatrixMultiplication] from an unsafe.Pointer.
func NDArrayQuantizedMatrixMultiplicationFrom(ptr unsafe.Pointer) NDArrayQuantizedMatrixMultiplication {
	return NDArrayQuantizedMatrixMultiplication{
		NDArrayMatrixMultiplication: NDArrayMatrixMultiplicationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NDArrayQuantizedMatrixMultiplication */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayquantizedmatrixmultiplication/4446153-initwithdevice
func NewNDArrayQuantizedMatrixMultiplicationWithDeviceLeftQuantizationDescriptorRightQuantizationDescriptor(device unsafe.Pointer, leftQuantizationDescriptor INDArrayQuantizationDescriptor, rightQuantizationDescriptor INDArrayQuantizationDescriptor) NDArrayQuantizedMatrixMultiplication {
	instance := getNDArrayQuantizedMatrixMultiplicationClass().Alloc()
	rv := objc.Send[NDArrayQuantizedMatrixMultiplication](instance.ID, objc.Sel("initWithDevice:leftQuantizationDescriptor:rightQuantizationDescriptor:"), device, leftQuantizationDescriptor, rightQuantizationDescriptor)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNDArrayQuantizedMatrixMultiplicationWithDeviceLeftQuantizationDescriptorRightQuantizationDescriptor */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NDArrayQuantizedMatrixMultiplication */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NDArrayQuantizedMatrixMultiplication */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NDArrayQuantizedMatrixMultiplication */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NDArrayQuantizedMatrixMultiplication */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNDArrayQuantizedMatrixMultiplication */


