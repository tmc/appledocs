// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNDArrayQuantizationDescriptor */


/* debug [class_header]: Header for MPSNDArrayQuantizationDescriptor */
// The class instance for the [NDArrayQuantizationDescriptor] class.
var (
	NDArrayQuantizationDescriptorClass     _NDArrayQuantizationDescriptorClass
	NDArrayQuantizationDescriptorClassOnce sync.Once
)

func getNDArrayQuantizationDescriptorClass() _NDArrayQuantizationDescriptorClass {
	NDArrayQuantizationDescriptorClassOnce.Do(func() {
		NDArrayQuantizationDescriptorClass = _NDArrayQuantizationDescriptorClass{objc.GetClass("MPSNDArrayQuantizationDescriptor")}
	})
	return NDArrayQuantizationDescriptorClass
}

type _NDArrayQuantizationDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NDArrayQuantizationDescriptor */
// An interface definition for the [NDArrayQuantizationDescriptor] class.
type INDArrayQuantizationDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NDArrayQuantizationDescriptor */
	// properties:
	QuantizationDataType() DataType get /* not a class type */
	SetQuantizationDataType(value DataType get /* not a class type */)
	QuantizationScheme() NDArrayQuantizationScheme get /* not a class type */
	SetQuantizationScheme(value NDArrayQuantizationScheme get /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NDArrayQuantizationDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NDArrayQuantizationDescriptor */
// Alloc allocates a new instance without initialization.
func (nc _NDArrayQuantizationDescriptorClass) Alloc() NDArrayQuantizationDescriptor {
	rv := objc.Send[NDArrayQuantizationDescriptor](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NDArrayQuantizationDescriptorClass) New() NDArrayQuantizationDescriptor {
	rv := objc.Send[NDArrayQuantizationDescriptor](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayQuantizationDescriptor) Init() NDArrayQuantizationDescriptor {
	rv := objc.Send[NDArrayQuantizationDescriptor](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayQuantizationDescriptor) Autorelease() NDArrayQuantizationDescriptor {
	rv := objc.Send[NDArrayQuantizationDescriptor](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayQuantizationDescriptor creates a new NDArrayQuantizationDescriptor instance.
func NewNDArrayQuantizationDescriptor() NDArrayQuantizationDescriptor {
	return getNDArrayQuantizationDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NDArrayQuantizationDescriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayQuantizationDescriptor
type NDArrayQuantizationDescriptor struct {
	objectivec.Object
}

// NDArrayQuantizationDescriptorFrom constructs a [NDArrayQuantizationDescriptor] from an unsafe.Pointer.
func NDArrayQuantizationDescriptorFrom(ptr unsafe.Pointer) NDArrayQuantizationDescriptor {
	return NDArrayQuantizationDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NDArrayQuantizationDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NDArrayQuantizationDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NDArrayQuantizationDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NDArrayQuantizationDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NDArrayQuantizationDescriptor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayquantizationdescriptor/4446142-quantizationdatatype
func (n_ NDArrayQuantizationDescriptor) QuantizationDataType() DataType get /* not a class type */ {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("quantizationDataType"))
	return rv
}/* debug [instance_properties/getter]: quantizationDataType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayquantizationdescriptor/4446142-quantizationdatatype
func (n_ NDArrayQuantizationDescriptor) SetQuantizationDataType(value DataType get /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setQuantizationDataType:"), value)
}/* debug [instance_properties/setter]: quantizationDataType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayquantizationdescriptor/4446143-quantizationscheme
func (n_ NDArrayQuantizationDescriptor) QuantizationScheme() NDArrayQuantizationScheme get /* not a class type */ {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("quantizationScheme"))
	return rv
}/* debug [instance_properties/getter]: quantizationScheme */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayquantizationdescriptor/4446143-quantizationscheme
func (n_ NDArrayQuantizationDescriptor) SetQuantizationScheme(value NDArrayQuantizationScheme get /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setQuantizationScheme:"), value)
}/* debug [instance_properties/setter]: quantizationScheme */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNDArrayQuantizationDescriptor */



