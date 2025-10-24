// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLCPoolingDescriptor */


/* debug [class_header]: Header for MLCPoolingDescriptor */
// The class instance for the [CPoolingDescriptor] class.
var (
	CPoolingDescriptorClass     _CPoolingDescriptorClass
	CPoolingDescriptorClassOnce sync.Once
)

func getCPoolingDescriptorClass() _CPoolingDescriptorClass {
	CPoolingDescriptorClassOnce.Do(func() {
		CPoolingDescriptorClass = _CPoolingDescriptorClass{objc.GetClass("MLCPoolingDescriptor")}
	})
	return CPoolingDescriptorClass
}

type _CPoolingDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CPoolingDescriptor */
// An interface definition for the [CPoolingDescriptor] class.
type ICPoolingDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CPoolingDescriptor */
	// properties:
	CountIncludesPadding() bool
	DilationRateInX() uint
	DilationRateInY() uint
	KernelHeight() uint
	KernelWidth() uint
	PaddingPolicy() CPaddingPolicy
	PaddingSizeInX() uint
	PaddingSizeInY() uint
	PoolingType() CPoolingType
	StrideInX() uint
	StrideInY() uint
	DilationRates() int
	SetDilationRates(value int)
	KernelSizes() int
	SetKernelSizes(value int)
	Strides() int
	SetStrides(value int)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CPoolingDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CPoolingDescriptor */
// Alloc allocates a new instance without initialization.
func (cc _CPoolingDescriptorClass) Alloc() CPoolingDescriptor {
	rv := objc.Send[CPoolingDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CPoolingDescriptorClass) New() CPoolingDescriptor {
	rv := objc.Send[CPoolingDescriptor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CPoolingDescriptor) Init() CPoolingDescriptor {
	rv := objc.Send[CPoolingDescriptor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CPoolingDescriptor) Autorelease() CPoolingDescriptor {
	rv := objc.Send[CPoolingDescriptor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPoolingDescriptor creates a new CPoolingDescriptor instance.
func NewCPoolingDescriptor() CPoolingDescriptor {
	return getCPoolingDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CPoolingDescriptor */
// A configuration object you use to create a pooling layer.


// A configuration object you use to create a pooling layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPoolingDescriptor
type CPoolingDescriptor struct {
	objectivec.Object
}

// CPoolingDescriptorFrom constructs a [CPoolingDescriptor] from an unsafe.Pointer.
//
// A configuration object you use to create a pooling layer.
func CPoolingDescriptorFrom(ptr unsafe.Pointer) CPoolingDescriptor {
	return CPoolingDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CPoolingDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CPoolingDescriptor */

// Creates an average pooling descriptor with the kernel sizes, strides, dilution rates, padding policy and sizes, and zero padding option you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPoolingDescriptor/averagePoolingDescriptorWithKernelSizes:strides:dilationRates:paddingPolicy:paddingSizes:countIncludesPadding:
func (cc _CPoolingDescriptorClass) AveragePoolingDescriptorWithKernelSizesStridesDilationRatesPaddingPolicyPaddingSizesCountIncludesPadding(kernelSizes []foundation.Number, strides []foundation.Number, dilationRates []foundation.Number, paddingPolicy CPaddingPolicy, paddingSizes []foundation.Number, countIncludesPadding bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("averagePoolingDescriptorWithKernelSizes:strides:dilationRates:paddingPolicy:paddingSizes:countIncludesPadding:"), kernelSizes, strides, dilationRates, paddingPolicy, paddingSizes, countIncludesPadding)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AveragePoolingDescriptorWithKernelSizesStridesDilationRatesPaddingPolicyPaddingSizesCountIncludesPadding) */


// Creates an average pooling descriptor with the kernel sizes, strides, padding policy, padding sizes, and zero padding option that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPoolingDescriptor/averagePoolingDescriptorWithKernelSizes:strides:paddingPolicy:paddingSizes:countIncludesPadding:
func (cc _CPoolingDescriptorClass) AveragePoolingDescriptorWithKernelSizesStridesPaddingPolicyPaddingSizesCountIncludesPadding(kernelSizes []foundation.Number, strides []foundation.Number, paddingPolicy CPaddingPolicy, paddingSizes []foundation.Number, countIncludesPadding bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("averagePoolingDescriptorWithKernelSizes:strides:paddingPolicy:paddingSizes:countIncludesPadding:"), kernelSizes, strides, paddingPolicy, paddingSizes, countIncludesPadding)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AveragePoolingDescriptorWithKernelSizesStridesPaddingPolicyPaddingSizesCountIncludesPadding) */


// Creates a descriptor for an L2 norm pooling function with the kernel sizes, strides, dilation rates, padding policy, and padding sizes you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPoolingDescriptor/l2NormPoolingDescriptorWithKernelSizes:strides:dilationRates:paddingPolicy:paddingSizes:
func (cc _CPoolingDescriptorClass) L2NormPoolingDescriptorWithKernelSizesStridesDilationRatesPaddingPolicyPaddingSizes(kernelSizes []foundation.Number, strides []foundation.Number, dilationRates []foundation.Number, paddingPolicy CPaddingPolicy, paddingSizes []foundation.Number) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("l2NormPoolingDescriptorWithKernelSizes:strides:dilationRates:paddingPolicy:paddingSizes:"), kernelSizes, strides, dilationRates, paddingPolicy, paddingSizes)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=L2NormPoolingDescriptorWithKernelSizesStridesDilationRatesPaddingPolicyPaddingSizes) */


// Creates a descriptor for an L2 norm pooling function with the kernel sizes, strides, padding policy, and padding sizes that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPoolingDescriptor/l2NormPoolingDescriptorWithKernelSizes:strides:paddingPolicy:paddingSizes:
func (cc _CPoolingDescriptorClass) L2NormPoolingDescriptorWithKernelSizesStridesPaddingPolicyPaddingSizes(kernelSizes []foundation.Number, strides []foundation.Number, paddingPolicy CPaddingPolicy, paddingSizes []foundation.Number) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("l2NormPoolingDescriptorWithKernelSizes:strides:paddingPolicy:paddingSizes:"), kernelSizes, strides, paddingPolicy, paddingSizes)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=L2NormPoolingDescriptorWithKernelSizesStridesPaddingPolicyPaddingSizes) */


// Creates a descriptor for a max pooling function with the kernel sizes, strides, dilation rates, padding policy, and padding sizes that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPoolingDescriptor/maxPoolingDescriptorWithKernelSizes:strides:dilationRates:paddingPolicy:paddingSizes:
func (cc _CPoolingDescriptorClass) MaxPoolingDescriptorWithKernelSizesStridesDilationRatesPaddingPolicyPaddingSizes(kernelSizes []foundation.Number, strides []foundation.Number, dilationRates []foundation.Number, paddingPolicy CPaddingPolicy, paddingSizes []foundation.Number) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("maxPoolingDescriptorWithKernelSizes:strides:dilationRates:paddingPolicy:paddingSizes:"), kernelSizes, strides, dilationRates, paddingPolicy, paddingSizes)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MaxPoolingDescriptorWithKernelSizesStridesDilationRatesPaddingPolicyPaddingSizes) */


// Creates a descriptor for a max pooling function with the kernel sizes, strides, padding policy, and padding sizes that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPoolingDescriptor/maxPoolingDescriptorWithKernelSizes:strides:paddingPolicy:paddingSizes:
func (cc _CPoolingDescriptorClass) MaxPoolingDescriptorWithKernelSizesStridesPaddingPolicyPaddingSizes(kernelSizes []foundation.Number, strides []foundation.Number, paddingPolicy CPaddingPolicy, paddingSizes []foundation.Number) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("maxPoolingDescriptorWithKernelSizes:strides:paddingPolicy:paddingSizes:"), kernelSizes, strides, paddingPolicy, paddingSizes)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MaxPoolingDescriptorWithKernelSizesStridesPaddingPolicyPaddingSizes) */


// Creates a pooling descriptor with the pooling function, kernel size, and stride you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPoolingDescriptor/poolingDescriptorWithType:kernelSize:stride:
func (cc _CPoolingDescriptorClass) PoolingDescriptorWithTypeKernelSizeStride(poolingType CPoolingType, kernelSize uint, stride uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("poolingDescriptorWithType:kernelSize:stride:"), poolingType, kernelSize, stride)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PoolingDescriptorWithTypeKernelSizeStride) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CPoolingDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CPoolingDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CPoolingDescriptor */

// A Boolean that indicates whether you include zero padding in the averaging calculation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPoolingDescriptor/countIncludesPadding
func (c_ CPoolingDescriptor) CountIncludesPadding() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("countIncludesPadding"))
	return rv
}/* debug [instance_properties/getter]: countIncludesPadding */


// The kernel dilation rate, or stride of elements, in x.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPoolingDescriptor/dilationRateInX
func (c_ CPoolingDescriptor) DilationRateInX() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("dilationRateInX"))
	return rv
}/* debug [instance_properties/getter]: dilationRateInX */


// The kernel dilation rate, or stride of elements, in y.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPoolingDescriptor/dilationRateInY
func (c_ CPoolingDescriptor) DilationRateInY() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("dilationRateInY"))
	return rv
}/* debug [instance_properties/getter]: dilationRateInY */


// The pooling kernel size in y.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPoolingDescriptor/kernelHeight
func (c_ CPoolingDescriptor) KernelHeight() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("kernelHeight"))
	return rv
}/* debug [instance_properties/getter]: kernelHeight */


// The pooling kernel size in x.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPoolingDescriptor/kernelWidth
func (c_ CPoolingDescriptor) KernelWidth() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("kernelWidth"))
	return rv
}/* debug [instance_properties/getter]: kernelWidth */


// The padding policy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPoolingDescriptor/paddingPolicy-1e6rg
func (c_ CPoolingDescriptor) PaddingPolicy() CPaddingPolicy {
	rv := objc.Send[CPaddingPolicy](c_.ID, objc.Sel("paddingPolicy"))
	return rv
}/* debug [instance_properties/getter]: paddingPolicy */


// The padding size in x, left and right, to use if the padding policy is to use padding size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPoolingDescriptor/paddingSizeInX
func (c_ CPoolingDescriptor) PaddingSizeInX() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("paddingSizeInX"))
	return rv
}/* debug [instance_properties/getter]: paddingSizeInX */


// The padding size in y, top and bottom, to use if the padding policy is to use padding size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPoolingDescriptor/paddingSizeInY
func (c_ CPoolingDescriptor) PaddingSizeInY() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("paddingSizeInY"))
	return rv
}/* debug [instance_properties/getter]: paddingSizeInY */


// The pooling operation type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPoolingDescriptor/poolingType-9daku
func (c_ CPoolingDescriptor) PoolingType() CPoolingType {
	rv := objc.Send[CPoolingType](c_.ID, objc.Sel("poolingType"))
	return rv
}/* debug [instance_properties/getter]: poolingType */


// The stride of the kernel in x.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPoolingDescriptor/strideInX
func (c_ CPoolingDescriptor) StrideInX() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("strideInX"))
	return rv
}/* debug [instance_properties/getter]: strideInX */


// The stride of the kernel in y.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPoolingDescriptor/strideInY
func (c_ CPoolingDescriptor) StrideInY() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("strideInY"))
	return rv
}/* debug [instance_properties/getter]: strideInY */


// A tuple that contains the dilation rates for y and x.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcpoolingdescriptor/dilationrates
func (c_ CPoolingDescriptor) DilationRates() int {
	rv := objc.Send[int](c_.ID, objc.Sel("dilationRates"))
	return rv
}/* debug [instance_properties/getter]: dilationRates */


// A tuple that contains the dilation rates for y and x.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcpoolingdescriptor/dilationrates
func (c_ CPoolingDescriptor) SetDilationRates(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDilationRates:"), value)
}/* debug [instance_properties/setter]: dilationRates */


// A tuple that contains the kernel sizes for height and width.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcpoolingdescriptor/kernelsizes
func (c_ CPoolingDescriptor) KernelSizes() int {
	rv := objc.Send[int](c_.ID, objc.Sel("kernelSizes"))
	return rv
}/* debug [instance_properties/getter]: kernelSizes */


// A tuple that contains the kernel sizes for height and width.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcpoolingdescriptor/kernelsizes
func (c_ CPoolingDescriptor) SetKernelSizes(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelSizes:"), value)
}/* debug [instance_properties/setter]: kernelSizes */


// A tuple that contains the kernel strides for y and x.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcpoolingdescriptor/strides
func (c_ CPoolingDescriptor) Strides() int {
	rv := objc.Send[int](c_.ID, objc.Sel("strides"))
	return rv
}/* debug [instance_properties/getter]: strides */


// A tuple that contains the kernel strides for y and x.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcpoolingdescriptor/strides
func (c_ CPoolingDescriptor) SetStrides(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStrides:"), value)
}/* debug [instance_properties/setter]: strides */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCPoolingDescriptor */



