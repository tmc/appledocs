// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNDArrayLUTQuantizationDescriptor */


/* debug [class_header]: Header for MPSNDArrayLUTQuantizationDescriptor */
// The class instance for the [NDArrayLUTQuantizationDescriptor] class.
var (
	NDArrayLUTQuantizationDescriptorClass     _NDArrayLUTQuantizationDescriptorClass
	NDArrayLUTQuantizationDescriptorClassOnce sync.Once
)

func getNDArrayLUTQuantizationDescriptorClass() _NDArrayLUTQuantizationDescriptorClass {
	NDArrayLUTQuantizationDescriptorClassOnce.Do(func() {
		NDArrayLUTQuantizationDescriptorClass = _NDArrayLUTQuantizationDescriptorClass{objc.GetClass("MPSNDArrayLUTQuantizationDescriptor")}
	})
	return NDArrayLUTQuantizationDescriptorClass
}

type _NDArrayLUTQuantizationDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NDArrayLUTQuantizationDescriptor */
// An interface definition for the [NDArrayLUTQuantizationDescriptor] class.
type INDArrayLUTQuantizationDescriptor interface {
	INDArrayQuantizationDescriptor
	
/* debug [class_interface_properties]: Properties for NDArrayLUTQuantizationDescriptor */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NDArrayLUTQuantizationDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NDArrayLUTQuantizationDescriptor */
// Alloc allocates a new instance without initialization.
func (nc _NDArrayLUTQuantizationDescriptorClass) Alloc() NDArrayLUTQuantizationDescriptor {
	rv := objc.Send[NDArrayLUTQuantizationDescriptor](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NDArrayLUTQuantizationDescriptorClass) New() NDArrayLUTQuantizationDescriptor {
	rv := objc.Send[NDArrayLUTQuantizationDescriptor](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayLUTQuantizationDescriptor) Init() NDArrayLUTQuantizationDescriptor {
	rv := objc.Send[NDArrayLUTQuantizationDescriptor](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayLUTQuantizationDescriptor) Autorelease() NDArrayLUTQuantizationDescriptor {
	rv := objc.Send[NDArrayLUTQuantizationDescriptor](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayLUTQuantizationDescriptor creates a new NDArrayLUTQuantizationDescriptor instance.
func NewNDArrayLUTQuantizationDescriptor() NDArrayLUTQuantizationDescriptor {
	return getNDArrayLUTQuantizationDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NDArrayLUTQuantizationDescriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayLUTQuantizationDescriptor
type NDArrayLUTQuantizationDescriptor struct {
	NDArrayQuantizationDescriptor
}

// NDArrayLUTQuantizationDescriptorFrom constructs a [NDArrayLUTQuantizationDescriptor] from an unsafe.Pointer.
func NDArrayLUTQuantizationDescriptorFrom(ptr unsafe.Pointer) NDArrayLUTQuantizationDescriptor {
	return NDArrayLUTQuantizationDescriptor{
		NDArrayQuantizationDescriptor: NDArrayQuantizationDescriptorFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NDArrayLUTQuantizationDescriptor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraylutquantizationdescriptor/4446139-initwithdatatype
func NewNDArrayLUTQuantizationDescriptorWithDataType(quantizationDataType DataType) NDArrayLUTQuantizationDescriptor {
	instance := getNDArrayLUTQuantizationDescriptorClass().Alloc()
	rv := objc.Send[NDArrayLUTQuantizationDescriptor](instance.ID, objc.Sel("initWithDataType:"), quantizationDataType)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNDArrayLUTQuantizationDescriptorWithDataType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraylutquantizationdescriptor/4446140-initwithdatatype
func NewNDArrayLUTQuantizationDescriptorWithDataTypeVectorAxis(quantizationDataType DataType, vectorAxis uint) NDArrayLUTQuantizationDescriptor {
	instance := getNDArrayLUTQuantizationDescriptorClass().Alloc()
	rv := objc.Send[NDArrayLUTQuantizationDescriptor](instance.ID, objc.Sel("initWithDataType:vectorAxis:"), quantizationDataType, vectorAxis)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNDArrayLUTQuantizationDescriptorWithDataTypeVectorAxis */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NDArrayLUTQuantizationDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NDArrayLUTQuantizationDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NDArrayLUTQuantizationDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NDArrayLUTQuantizationDescriptor */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNDArrayLUTQuantizationDescriptor */


