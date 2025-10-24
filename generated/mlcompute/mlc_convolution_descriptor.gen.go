// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLCConvolutionDescriptor */


/* debug [class_header]: Header for MLCConvolutionDescriptor */
// The class instance for the [CConvolutionDescriptor] class.
var (
	CConvolutionDescriptorClass     _CConvolutionDescriptorClass
	CConvolutionDescriptorClassOnce sync.Once
)

func getCConvolutionDescriptorClass() _CConvolutionDescriptorClass {
	CConvolutionDescriptorClassOnce.Do(func() {
		CConvolutionDescriptorClass = _CConvolutionDescriptorClass{objc.GetClass("MLCConvolutionDescriptor")}
	})
	return CConvolutionDescriptorClass
}

type _CConvolutionDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CConvolutionDescriptor */
// An interface definition for the [CConvolutionDescriptor] class.
type ICConvolutionDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CConvolutionDescriptor */
	// properties:
	ConvolutionType() CConvolutionType
	DilationRateInX() uint
	DilationRateInY() uint
	GroupCount() uint
	InputFeatureChannelCount() uint
	IsConvolutionTranspose() bool
	KernelHeight() uint
	KernelWidth() uint
	OutputFeatureChannelCount() uint
	PaddingPolicy() CPaddingPolicy
	PaddingSizeInX() uint
	PaddingSizeInY() uint
	StrideInX() uint
	StrideInY() uint
	UsesDepthwiseConvolution() bool
	DilationRates() int
	SetDilationRates(value int)
	KernelSizes() int
	SetKernelSizes(value int)
	Strides() int
	SetStrides(value int)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CConvolutionDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CConvolutionDescriptor */
// Alloc allocates a new instance without initialization.
func (cc _CConvolutionDescriptorClass) Alloc() CConvolutionDescriptor {
	rv := objc.Send[CConvolutionDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CConvolutionDescriptorClass) New() CConvolutionDescriptor {
	rv := objc.Send[CConvolutionDescriptor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CConvolutionDescriptor) Init() CConvolutionDescriptor {
	rv := objc.Send[CConvolutionDescriptor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CConvolutionDescriptor) Autorelease() CConvolutionDescriptor {
	rv := objc.Send[CConvolutionDescriptor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCConvolutionDescriptor creates a new CConvolutionDescriptor instance.
func NewCConvolutionDescriptor() CConvolutionDescriptor {
	return getCConvolutionDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CConvolutionDescriptor */
// A configuration object you use to create a convolution or fully connected layer.


// A configuration object you use to create a convolution or fully connected layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionDescriptor
type CConvolutionDescriptor struct {
	objectivec.Object
}

// CConvolutionDescriptorFrom constructs a [CConvolutionDescriptor] from an unsafe.Pointer.
//
// A configuration object you use to create a convolution or fully connected layer.
func CConvolutionDescriptorFrom(ptr unsafe.Pointer) CConvolutionDescriptor {
	return CConvolutionDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CConvolutionDescriptor */

// Creates a descriptor for convolution transpose with the kernel sizes and number of feature channels you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionDescriptor/init(transposeWithKernelWidth:kernelHeight:inputFeatureChannelCount:outputFeatureChannelCount:)
func NewCConvolutionDescriptorConvolutionTransposeDescriptorWithKernelWidthKernelHeightInputFeatureChannelCountOutputFeatureChannelCount(kernelWidth uint, kernelHeight uint, inputFeatureChannelCount uint, outputFeatureChannelCount uint) CConvolutionDescriptor {
	rv := objc.Send[CConvolutionDescriptor](objc.ID(getCConvolutionDescriptorClass().class), objc.Sel("convolutionTransposeDescriptorWithKernelWidth:kernelHeight:inputFeatureChannelCount:outputFeatureChannelCount:"), kernelWidth, kernelHeight, inputFeatureChannelCount, outputFeatureChannelCount)
	return rv
}/* debug [class_init_methods/constructor]: NewCConvolutionDescriptorConvolutionTransposeDescriptorWithKernelWidthKernelHeightInputFeatureChannelCountOutputFeatureChannelCount */


// Creates a descriptor for depthwise convolution with the kernel sizes, number of input feature channels, and channel multiplier you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionDescriptor/init(depthwiseWithKernelWidth:kernelHeight:inputFeatureChannelCount:channelMultiplier:)
func NewCConvolutionDescriptorDepthwiseConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelCountChannelMultiplier(kernelWidth uint, kernelHeight uint, inputFeatureChannelCount uint, channelMultiplier uint) CConvolutionDescriptor {
	rv := objc.Send[CConvolutionDescriptor](objc.ID(getCConvolutionDescriptorClass().class), objc.Sel("depthwiseConvolutionDescriptorWithKernelWidth:kernelHeight:inputFeatureChannelCount:channelMultiplier:"), kernelWidth, kernelHeight, inputFeatureChannelCount, channelMultiplier)
	return rv
}/* debug [class_init_methods/constructor]: NewCConvolutionDescriptorDepthwiseConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelCountChannelMultiplier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CConvolutionDescriptor */

// Creates a convolution transpose descriptor with the kernel and padding options, number of feature channels and groups, and dilation rates you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionDescriptor/convolutionTransposeDescriptorWithKernelSizes:inputFeatureChannelCount:outputFeatureChannelCount:groupCount:strides:dilationRates:paddingPolicy:paddingSizes:
func (cc _CConvolutionDescriptorClass) ConvolutionTransposeDescriptorWithKernelSizesInputFeatureChannelCountOutputFeatureChannelCountGroupCountStridesDilationRatesPaddingPolicyPaddingSizes(kernelSizes []foundation.Number, inputFeatureChannelCount uint, outputFeatureChannelCount uint, groupCount uint, strides []foundation.Number, dilationRates []foundation.Number, paddingPolicy CPaddingPolicy, paddingSizes []foundation.Number) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("convolutionTransposeDescriptorWithKernelSizes:inputFeatureChannelCount:outputFeatureChannelCount:groupCount:strides:dilationRates:paddingPolicy:paddingSizes:"), kernelSizes, inputFeatureChannelCount, outputFeatureChannelCount, groupCount, strides, dilationRates, paddingPolicy, paddingSizes)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ConvolutionTransposeDescriptorWithKernelSizesInputFeatureChannelCountOutputFeatureChannelCountGroupCountStridesDilationRatesPaddingPolicyPaddingSizes) */


// Creates a convolution transpose descriptor with the kernel sizes, number of feature channels, strides, and padding options you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionDescriptor/convolutionTransposeDescriptorWithKernelSizes:inputFeatureChannelCount:outputFeatureChannelCount:strides:paddingPolicy:paddingSizes:
func (cc _CConvolutionDescriptorClass) ConvolutionTransposeDescriptorWithKernelSizesInputFeatureChannelCountOutputFeatureChannelCountStridesPaddingPolicyPaddingSizes(kernelSizes []foundation.Number, inputFeatureChannelCount uint, outputFeatureChannelCount uint, strides []foundation.Number, paddingPolicy CPaddingPolicy, paddingSizes []foundation.Number) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("convolutionTransposeDescriptorWithKernelSizes:inputFeatureChannelCount:outputFeatureChannelCount:strides:paddingPolicy:paddingSizes:"), kernelSizes, inputFeatureChannelCount, outputFeatureChannelCount, strides, paddingPolicy, paddingSizes)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ConvolutionTransposeDescriptorWithKernelSizesInputFeatureChannelCountOutputFeatureChannelCountStridesPaddingPolicyPaddingSizes) */


// Creates a convolution descriptor with the kernel and padding options, number of input channels, channel multiplier, and dilation rates you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionDescriptor/depthwiseConvolutionDescriptorWithKernelSizes:inputFeatureChannelCount:channelMultiplier:strides:dilationRates:paddingPolicy:paddingSizes:
func (cc _CConvolutionDescriptorClass) DepthwiseConvolutionDescriptorWithKernelSizesInputFeatureChannelCountChannelMultiplierStridesDilationRatesPaddingPolicyPaddingSizes(kernelSizes []foundation.Number, inputFeatureChannelCount uint, channelMultiplier uint, strides []foundation.Number, dilationRates []foundation.Number, paddingPolicy CPaddingPolicy, paddingSizes []foundation.Number) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("depthwiseConvolutionDescriptorWithKernelSizes:inputFeatureChannelCount:channelMultiplier:strides:dilationRates:paddingPolicy:paddingSizes:"), kernelSizes, inputFeatureChannelCount, channelMultiplier, strides, dilationRates, paddingPolicy, paddingSizes)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DepthwiseConvolutionDescriptorWithKernelSizesInputFeatureChannelCountChannelMultiplierStridesDilationRatesPaddingPolicyPaddingSizes) */


// Creates a depthwise convolution descriptor with the kernel and padding options, number of input feature channels, and channel multiplier you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionDescriptor/depthwiseConvolutionDescriptorWithKernelSizes:inputFeatureChannelCount:channelMultiplier:strides:paddingPolicy:paddingSizes:
func (cc _CConvolutionDescriptorClass) DepthwiseConvolutionDescriptorWithKernelSizesInputFeatureChannelCountChannelMultiplierStridesPaddingPolicyPaddingSizes(kernelSizes []foundation.Number, inputFeatureChannelCount uint, channelMultiplier uint, strides []foundation.Number, paddingPolicy CPaddingPolicy, paddingSizes []foundation.Number) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("depthwiseConvolutionDescriptorWithKernelSizes:inputFeatureChannelCount:channelMultiplier:strides:paddingPolicy:paddingSizes:"), kernelSizes, inputFeatureChannelCount, channelMultiplier, strides, paddingPolicy, paddingSizes)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DepthwiseConvolutionDescriptorWithKernelSizesInputFeatureChannelCountChannelMultiplierStridesPaddingPolicyPaddingSizes) */


// Creates a convolution descriptor with the kernel and padding options, number of feature channels and groups, and dilation rates you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionDescriptor/descriptorWithKernelSizes:inputFeatureChannelCount:outputFeatureChannelCount:groupCount:strides:dilationRates:paddingPolicy:paddingSizes:
func (cc _CConvolutionDescriptorClass) DescriptorWithKernelSizesInputFeatureChannelCountOutputFeatureChannelCountGroupCountStridesDilationRatesPaddingPolicyPaddingSizes(kernelSizes []foundation.Number, inputFeatureChannelCount uint, outputFeatureChannelCount uint, groupCount uint, strides []foundation.Number, dilationRates []foundation.Number, paddingPolicy CPaddingPolicy, paddingSizes []foundation.Number) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("descriptorWithKernelSizes:inputFeatureChannelCount:outputFeatureChannelCount:groupCount:strides:dilationRates:paddingPolicy:paddingSizes:"), kernelSizes, inputFeatureChannelCount, outputFeatureChannelCount, groupCount, strides, dilationRates, paddingPolicy, paddingSizes)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithKernelSizesInputFeatureChannelCountOutputFeatureChannelCountGroupCountStridesDilationRatesPaddingPolicyPaddingSizes) */


// Creates a convolution descriptor with the kernel sizes, number of feature channels, strides, padding policy, and padding sizes you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionDescriptor/descriptorWithKernelSizes:inputFeatureChannelCount:outputFeatureChannelCount:strides:paddingPolicy:paddingSizes:
func (cc _CConvolutionDescriptorClass) DescriptorWithKernelSizesInputFeatureChannelCountOutputFeatureChannelCountStridesPaddingPolicyPaddingSizes(kernelSizes []foundation.Number, inputFeatureChannelCount uint, outputFeatureChannelCount uint, strides []foundation.Number, paddingPolicy CPaddingPolicy, paddingSizes []foundation.Number) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("descriptorWithKernelSizes:inputFeatureChannelCount:outputFeatureChannelCount:strides:paddingPolicy:paddingSizes:"), kernelSizes, inputFeatureChannelCount, outputFeatureChannelCount, strides, paddingPolicy, paddingSizes)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithKernelSizesInputFeatureChannelCountOutputFeatureChannelCountStridesPaddingPolicyPaddingSizes) */


// Creates a convolution descriptor with the kernel sizes and number of feature channels you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionDescriptor/descriptorWithKernelWidth:kernelHeight:inputFeatureChannelCount:outputFeatureChannelCount:
func (cc _CConvolutionDescriptorClass) DescriptorWithKernelWidthKernelHeightInputFeatureChannelCountOutputFeatureChannelCount(kernelWidth uint, kernelHeight uint, inputFeatureChannelCount uint, outputFeatureChannelCount uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("descriptorWithKernelWidth:kernelHeight:inputFeatureChannelCount:outputFeatureChannelCount:"), kernelWidth, kernelHeight, inputFeatureChannelCount, outputFeatureChannelCount)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithKernelWidthKernelHeightInputFeatureChannelCountOutputFeatureChannelCount) */


// Creates a descriptor with the type, kernel sizes, number of feature channels and groups, strides, dilation rates, and padding policy you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionDescriptor/descriptorWithType:kernelSizes:inputFeatureChannelCount:outputFeatureChannelCount:groupCount:strides:dilationRates:paddingPolicy:paddingSizes:
func (cc _CConvolutionDescriptorClass) DescriptorWithTypeKernelSizesInputFeatureChannelCountOutputFeatureChannelCountGroupCountStridesDilationRatesPaddingPolicyPaddingSizes(convolutionType CConvolutionType, kernelSizes []foundation.Number, inputFeatureChannelCount uint, outputFeatureChannelCount uint, groupCount uint, strides []foundation.Number, dilationRates []foundation.Number, paddingPolicy CPaddingPolicy, paddingSizes []foundation.Number) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("descriptorWithType:kernelSizes:inputFeatureChannelCount:outputFeatureChannelCount:groupCount:strides:dilationRates:paddingPolicy:paddingSizes:"), convolutionType, kernelSizes, inputFeatureChannelCount, outputFeatureChannelCount, groupCount, strides, dilationRates, paddingPolicy, paddingSizes)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithTypeKernelSizesInputFeatureChannelCountOutputFeatureChannelCountGroupCountStridesDilationRatesPaddingPolicyPaddingSizes) */


// Creates a descriptor for depthwise convolution with the kernel sizes, number of input feature channels, and channel multiplier you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionDescriptor/init(depthwiseWithKernelWidth:kernelHeight:inputFeatureChannelCount:channelMultiplier:)
func (cc _CConvolutionDescriptorClass) DepthwiseConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelCountChannelMultiplier(kernelWidth uint, kernelHeight uint, inputFeatureChannelCount uint, channelMultiplier uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("depthwiseConvolutionDescriptorWithKernelWidth:kernelHeight:inputFeatureChannelCount:channelMultiplier:"), kernelWidth, kernelHeight, inputFeatureChannelCount, channelMultiplier)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DepthwiseConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelCountChannelMultiplier) */


// Creates a descriptor for convolution transpose with the kernel sizes and number of feature channels you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionDescriptor/init(transposeWithKernelWidth:kernelHeight:inputFeatureChannelCount:outputFeatureChannelCount:)
func (cc _CConvolutionDescriptorClass) ConvolutionTransposeDescriptorWithKernelWidthKernelHeightInputFeatureChannelCountOutputFeatureChannelCount(kernelWidth uint, kernelHeight uint, inputFeatureChannelCount uint, outputFeatureChannelCount uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("convolutionTransposeDescriptorWithKernelWidth:kernelHeight:inputFeatureChannelCount:outputFeatureChannelCount:"), kernelWidth, kernelHeight, inputFeatureChannelCount, outputFeatureChannelCount)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ConvolutionTransposeDescriptorWithKernelWidthKernelHeightInputFeatureChannelCountOutputFeatureChannelCount) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CConvolutionDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CConvolutionDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CConvolutionDescriptor */

// The type of convolution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionDescriptor/convolutionType
func (c_ CConvolutionDescriptor) ConvolutionType() CConvolutionType {
	rv := objc.Send[CConvolutionType](c_.ID, objc.Sel("convolutionType"))
	return rv
}/* debug [instance_properties/getter]: convolutionType */


// The dilation rate, or stride of elements, in the kernel in x.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionDescriptor/dilationRateInX
func (c_ CConvolutionDescriptor) DilationRateInX() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("dilationRateInX"))
	return rv
}/* debug [instance_properties/getter]: dilationRateInX */


// The dilation rate, or stride of elements, in the kernel in y.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionDescriptor/dilationRateInY
func (c_ CConvolutionDescriptor) DilationRateInY() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("dilationRateInY"))
	return rv
}/* debug [instance_properties/getter]: dilationRateInY */


// The number of groups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionDescriptor/groupCount
func (c_ CConvolutionDescriptor) GroupCount() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("groupCount"))
	return rv
}/* debug [instance_properties/getter]: groupCount */


// The number of feature channels in the input tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionDescriptor/inputFeatureChannelCount
func (c_ CConvolutionDescriptor) InputFeatureChannelCount() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("inputFeatureChannelCount"))
	return rv
}/* debug [instance_properties/getter]: inputFeatureChannelCount */


// A Boolean that indicates whether this is a convolution transpose.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionDescriptor/isConvolutionTranspose
func (c_ CConvolutionDescriptor) IsConvolutionTranspose() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isConvolutionTranspose"))
	return rv
}/* debug [instance_properties/getter]: isConvolutionTranspose */


// The kernel size in y.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionDescriptor/kernelHeight
func (c_ CConvolutionDescriptor) KernelHeight() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("kernelHeight"))
	return rv
}/* debug [instance_properties/getter]: kernelHeight */


// The kernel size in x.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionDescriptor/kernelWidth
func (c_ CConvolutionDescriptor) KernelWidth() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("kernelWidth"))
	return rv
}/* debug [instance_properties/getter]: kernelWidth */


// The number of feature channels in the output tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionDescriptor/outputFeatureChannelCount
func (c_ CConvolutionDescriptor) OutputFeatureChannelCount() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("outputFeatureChannelCount"))
	return rv
}/* debug [instance_properties/getter]: outputFeatureChannelCount */


// The padding policy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionDescriptor/paddingPolicy-7drfq
func (c_ CConvolutionDescriptor) PaddingPolicy() CPaddingPolicy {
	rv := objc.Send[CPaddingPolicy](c_.ID, objc.Sel("paddingPolicy"))
	return rv
}/* debug [instance_properties/getter]: paddingPolicy */


// The pooling size in x, left and right, to use if the padding policy is to use padding size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionDescriptor/paddingSizeInX
func (c_ CConvolutionDescriptor) PaddingSizeInX() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("paddingSizeInX"))
	return rv
}/* debug [instance_properties/getter]: paddingSizeInX */


// The pooling size in y, top and bottom, to use if the padding policy is to use padding size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionDescriptor/paddingSizeInY
func (c_ CConvolutionDescriptor) PaddingSizeInY() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("paddingSizeInY"))
	return rv
}/* debug [instance_properties/getter]: paddingSizeInY */


// The kernel stride in .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionDescriptor/strideInX
func (c_ CConvolutionDescriptor) StrideInX() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("strideInX"))
	return rv
}/* debug [instance_properties/getter]: strideInX */


// The kernel stride in .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionDescriptor/strideInY
func (c_ CConvolutionDescriptor) StrideInY() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("strideInY"))
	return rv
}/* debug [instance_properties/getter]: strideInY */


// A Boolean that indicates whether you use depthwise convolution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionDescriptor/usesDepthwiseConvolution
func (c_ CConvolutionDescriptor) UsesDepthwiseConvolution() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("usesDepthwiseConvolution"))
	return rv
}/* debug [instance_properties/getter]: usesDepthwiseConvolution */


// A tuple that contains the dilation rates for y and x.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcconvolutiondescriptor/dilationrates
func (c_ CConvolutionDescriptor) DilationRates() int {
	rv := objc.Send[int](c_.ID, objc.Sel("dilationRates"))
	return rv
}/* debug [instance_properties/getter]: dilationRates */


// A tuple that contains the dilation rates for y and x.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcconvolutiondescriptor/dilationrates
func (c_ CConvolutionDescriptor) SetDilationRates(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDilationRates:"), value)
}/* debug [instance_properties/setter]: dilationRates */


// A tuple that contains the kernel sizes for height and width.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcconvolutiondescriptor/kernelsizes
func (c_ CConvolutionDescriptor) KernelSizes() int {
	rv := objc.Send[int](c_.ID, objc.Sel("kernelSizes"))
	return rv
}/* debug [instance_properties/getter]: kernelSizes */


// A tuple that contains the kernel sizes for height and width.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcconvolutiondescriptor/kernelsizes
func (c_ CConvolutionDescriptor) SetKernelSizes(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelSizes:"), value)
}/* debug [instance_properties/setter]: kernelSizes */


// A tuple that contains the kernel strides for y and x.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcconvolutiondescriptor/strides
func (c_ CConvolutionDescriptor) Strides() int {
	rv := objc.Send[int](c_.ID, objc.Sel("strides"))
	return rv
}/* debug [instance_properties/getter]: strides */


// A tuple that contains the kernel strides for y and x.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcconvolutiondescriptor/strides
func (c_ CConvolutionDescriptor) SetStrides(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStrides:"), value)
}/* debug [instance_properties/setter]: strides */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCConvolutionDescriptor */


