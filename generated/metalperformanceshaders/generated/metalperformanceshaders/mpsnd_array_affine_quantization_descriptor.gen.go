// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNDArrayAffineQuantizationDescriptor */


/* debug [class_header]: Header for MPSNDArrayAffineQuantizationDescriptor */
// The class instance for the [NDArrayAffineQuantizationDescriptor] class.
var (
	NDArrayAffineQuantizationDescriptorClass     _NDArrayAffineQuantizationDescriptorClass
	NDArrayAffineQuantizationDescriptorClassOnce sync.Once
)

func getNDArrayAffineQuantizationDescriptorClass() _NDArrayAffineQuantizationDescriptorClass {
	NDArrayAffineQuantizationDescriptorClassOnce.Do(func() {
		NDArrayAffineQuantizationDescriptorClass = _NDArrayAffineQuantizationDescriptorClass{objc.GetClass("MPSNDArrayAffineQuantizationDescriptor")}
	})
	return NDArrayAffineQuantizationDescriptorClass
}

type _NDArrayAffineQuantizationDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NDArrayAffineQuantizationDescriptor */
// An interface definition for the [NDArrayAffineQuantizationDescriptor] class.
type INDArrayAffineQuantizationDescriptor interface {
	INDArrayQuantizationDescriptor
	
/* debug [class_interface_properties]: Properties for NDArrayAffineQuantizationDescriptor */
	// properties:
	HasMinValue() objectivec.IObject
	SetHasMinValue(value objectivec.IObject)
	HasZeroPoint() objectivec.IObject
	SetHasZeroPoint(value objectivec.IObject)
	ImplicitZeroPoint() objectivec.IObject
	SetImplicitZeroPoint(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NDArrayAffineQuantizationDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NDArrayAffineQuantizationDescriptor */
// Alloc allocates a new instance without initialization.
func (nc _NDArrayAffineQuantizationDescriptorClass) Alloc() NDArrayAffineQuantizationDescriptor {
	rv := objc.Send[NDArrayAffineQuantizationDescriptor](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NDArrayAffineQuantizationDescriptorClass) New() NDArrayAffineQuantizationDescriptor {
	rv := objc.Send[NDArrayAffineQuantizationDescriptor](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayAffineQuantizationDescriptor) Init() NDArrayAffineQuantizationDescriptor {
	rv := objc.Send[NDArrayAffineQuantizationDescriptor](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayAffineQuantizationDescriptor) Autorelease() NDArrayAffineQuantizationDescriptor {
	rv := objc.Send[NDArrayAffineQuantizationDescriptor](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayAffineQuantizationDescriptor creates a new NDArrayAffineQuantizationDescriptor instance.
func NewNDArrayAffineQuantizationDescriptor() NDArrayAffineQuantizationDescriptor {
	return getNDArrayAffineQuantizationDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NDArrayAffineQuantizationDescriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayAffineQuantizationDescriptor
type NDArrayAffineQuantizationDescriptor struct {
	NDArrayQuantizationDescriptor
}

// NDArrayAffineQuantizationDescriptorFrom constructs a [NDArrayAffineQuantizationDescriptor] from an unsafe.Pointer.
func NDArrayAffineQuantizationDescriptorFrom(ptr unsafe.Pointer) NDArrayAffineQuantizationDescriptor {
	return NDArrayAffineQuantizationDescriptor{
		NDArrayQuantizationDescriptor: NDArrayQuantizationDescriptorFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NDArrayAffineQuantizationDescriptor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayaffinequantizationdescriptor/4446137-initwithdatatype
func NewNDArrayAffineQuantizationDescriptorWithDataTypeHasZeroPointHasMinValue(quantizationDataType DataType, hasZeroPoint bool, hasMinValue bool) NDArrayAffineQuantizationDescriptor {
	instance := getNDArrayAffineQuantizationDescriptorClass().Alloc()
	rv := objc.Send[NDArrayAffineQuantizationDescriptor](instance.ID, objc.Sel("initWithDataType:hasZeroPoint:hasMinValue:"), quantizationDataType, hasZeroPoint, hasMinValue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNDArrayAffineQuantizationDescriptorWithDataTypeHasZeroPointHasMinValue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NDArrayAffineQuantizationDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NDArrayAffineQuantizationDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NDArrayAffineQuantizationDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NDArrayAffineQuantizationDescriptor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayaffinequantizationdescriptor/4446134-hasminvalue
func (n_ NDArrayAffineQuantizationDescriptor) HasMinValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("hasMinValue"))
	return rv
}/* debug [instance_properties/getter]: hasMinValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayaffinequantizationdescriptor/4446134-hasminvalue
func (n_ NDArrayAffineQuantizationDescriptor) SetHasMinValue(value objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHasMinValue:"), value)
}/* debug [instance_properties/setter]: hasMinValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayaffinequantizationdescriptor/4446135-haszeropoint
func (n_ NDArrayAffineQuantizationDescriptor) HasZeroPoint() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("hasZeroPoint"))
	return rv
}/* debug [instance_properties/getter]: hasZeroPoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayaffinequantizationdescriptor/4446135-haszeropoint
func (n_ NDArrayAffineQuantizationDescriptor) SetHasZeroPoint(value objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHasZeroPoint:"), value)
}/* debug [instance_properties/setter]: hasZeroPoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayaffinequantizationdescriptor/4462739-implicitzeropoint
func (n_ NDArrayAffineQuantizationDescriptor) ImplicitZeroPoint() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("implicitZeroPoint"))
	return rv
}/* debug [instance_properties/getter]: implicitZeroPoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayaffinequantizationdescriptor/4462739-implicitzeropoint
func (n_ NDArrayAffineQuantizationDescriptor) SetImplicitZeroPoint(value objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setImplicitZeroPoint:"), value)
}/* debug [instance_properties/setter]: implicitZeroPoint */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNDArrayAffineQuantizationDescriptor */


