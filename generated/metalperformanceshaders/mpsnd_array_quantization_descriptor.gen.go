// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [NDArrayQuantizationDescriptor] class.
type INDArrayQuantizationDescriptor interface {
	objectivec.IObject
	

	// properties:
	QuantizationDataType() DataType get /* not a class type */
	SetQuantizationDataType(value DataType get /* not a class type */)
	QuantizationScheme() NDArrayQuantizationScheme get /* not a class type */
	SetQuantizationScheme(value NDArrayQuantizationScheme get /* not a class type */)


	

	// methods:


}





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







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayQuantizationDescriptor
type NDArrayQuantizationDescriptor struct {
	objectivec.Object
}

// NDArrayQuantizationDescriptorFrom constructs a [NDArrayQuantizationDescriptor] from an unsafe.Pointer.
func NDArrayQuantizationDescriptorFrom(ptr unsafe.Pointer) NDArrayQuantizationDescriptor {
	return NDArrayQuantizationDescriptor{objectivec.Object{objc.ID(ptr)}}
}

























// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayquantizationdescriptor/4446142-quantizationdatatype
func (n_ NDArrayQuantizationDescriptor) QuantizationDataType() DataType get /* not a class type */ {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("quantizationDataType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayquantizationdescriptor/4446142-quantizationdatatype
func (n_ NDArrayQuantizationDescriptor) SetQuantizationDataType(value DataType get /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setQuantizationDataType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayquantizationdescriptor/4446143-quantizationscheme
func (n_ NDArrayQuantizationDescriptor) QuantizationScheme() NDArrayQuantizationScheme get /* not a class type */ {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("quantizationScheme"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayquantizationdescriptor/4446143-quantizationscheme
func (n_ NDArrayQuantizationDescriptor) SetQuantizationScheme(value NDArrayQuantizationScheme get /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setQuantizationScheme:"), value)
}








