// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSMatrixCopyToImage */


/* debug [class_header]: Header for MPSMatrixCopyToImage */
// The class instance for the [MatrixCopyToImage] class.
var (
	MatrixCopyToImageClass     _MatrixCopyToImageClass
	MatrixCopyToImageClassOnce sync.Once
)

func getMatrixCopyToImageClass() _MatrixCopyToImageClass {
	MatrixCopyToImageClassOnce.Do(func() {
		MatrixCopyToImageClass = _MatrixCopyToImageClass{objc.GetClass("MPSMatrixCopyToImage")}
	})
	return MatrixCopyToImageClass
}

type _MatrixCopyToImageClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MatrixCopyToImage */
// An interface definition for the [MatrixCopyToImage] class.
type IMatrixCopyToImage interface {
	IKernel
	
/* debug [class_interface_properties]: Properties for MatrixCopyToImage */
	// properties:
	DataLayout() DataLayout get /* not a class type */
	SetDataLayout(value DataLayout get /* not a class type */)
	SourceMatrixBatchIndex() objectivec.IObject
	SetSourceMatrixBatchIndex(value objectivec.IObject)
	SourceMatrixOrigin() Origin get set /* not a class type */
	SetSourceMatrixOrigin(value Origin get set /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MatrixCopyToImage */
	// methods:
	Encode()
	EncodeToCommandBufferSourceMatrixDestinationImage(commandBuffer unsafe.Pointer, sourceMatrix IMatrix, destinationImage IImage)
	EncodeBatch()
	EncodeBatchToCommandBufferSourceMatrixDestinationImages(commandBuffer unsafe.Pointer, sourceMatrix IMatrix, destinationImages ImageBatch /* not a class type */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MatrixCopyToImage */
// Alloc allocates a new instance without initialization.
func (mc _MatrixCopyToImageClass) Alloc() MatrixCopyToImage {
	rv := objc.Send[MatrixCopyToImage](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MatrixCopyToImageClass) New() MatrixCopyToImage {
	rv := objc.Send[MatrixCopyToImage](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MatrixCopyToImage) Init() MatrixCopyToImage {
	rv := objc.Send[MatrixCopyToImage](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MatrixCopyToImage) Autorelease() MatrixCopyToImage {
	rv := objc.Send[MatrixCopyToImage](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatrixCopyToImage creates a new MatrixCopyToImage instance.
func NewMatrixCopyToImage() MatrixCopyToImage {
	return getMatrixCopyToImageClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MatrixCopyToImage */
// A kernel that copies matrix data to a Metal Performance Shaders image.


// A kernel that copies matrix data to a Metal Performance Shaders image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixCopyToImage
type MatrixCopyToImage struct {
	Kernel
}

// MatrixCopyToImageFrom constructs a [MatrixCopyToImage] from an unsafe.Pointer.
//
// A kernel that copies matrix data to a Metal Performance Shaders image.
func MatrixCopyToImageFrom(ptr unsafe.Pointer) MatrixCopyToImage {
	return MatrixCopyToImage{
		Kernel: KernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MatrixCopyToImage */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopytoimage/2976459-initwithcoder
func NewMatrixCopyToImageWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) MatrixCopyToImage {
	instance := getMatrixCopyToImageClass().Alloc()
	rv := objc.Send[MatrixCopyToImage](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMatrixCopyToImageWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopytoimage/2976460-initwithdevice
func NewMatrixCopyToImageWithDeviceDataLayout(device unsafe.Pointer, dataLayout DataLayout) MatrixCopyToImage {
	instance := getMatrixCopyToImageClass().Alloc()
	rv := objc.Send[MatrixCopyToImage](instance.ID, objc.Sel("initWithDevice:dataLayout:"), device, dataLayout)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMatrixCopyToImageWithDeviceDataLayout */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MatrixCopyToImage */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MatrixCopyToImage */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MatrixCopyToImage */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopytoimage/2976458-encode
func (m_ MatrixCopyToImage) Encode() {
	objc.Send[objc.ID](m_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopytoimage/2976458-encodetocommandbuffer
func (m_ MatrixCopyToImage) EncodeToCommandBufferSourceMatrixDestinationImage(commandBuffer unsafe.Pointer, sourceMatrix IMatrix, destinationImage IImage) {
	objc.Send[objc.ID](m_.ID, objc.Sel("encodeToCommandBuffer:sourceMatrix:destinationImage:"), commandBuffer, sourceMatrix, destinationImage)
}/* debug [instance_methods/method]: EncodeToCommandBufferSourceMatrixDestinationImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopytoimage/3013770-encodebatch
func (m_ MatrixCopyToImage) EncodeBatch() {
	objc.Send[objc.ID](m_.ID, objc.Sel("encodeBatch"))
}/* debug [instance_methods/method]: EncodeBatch */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopytoimage/3013770-encodebatchtocommandbuffer
func (m_ MatrixCopyToImage) EncodeBatchToCommandBufferSourceMatrixDestinationImages(commandBuffer unsafe.Pointer, sourceMatrix IMatrix, destinationImages ImageBatch /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceMatrix:destinationImages:"), commandBuffer, sourceMatrix, destinationImages)
}/* debug [instance_methods/method]: EncodeBatchToCommandBufferSourceMatrixDestinationImages */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MatrixCopyToImage */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopytoimage/2976457-datalayout
func (m_ MatrixCopyToImage) DataLayout() DataLayout get /* not a class type */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("dataLayout"))
	return rv
}/* debug [instance_properties/getter]: dataLayout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopytoimage/2976457-datalayout
func (m_ MatrixCopyToImage) SetDataLayout(value DataLayout get /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDataLayout:"), value)
}/* debug [instance_properties/setter]: dataLayout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopytoimage/2976461-sourcematrixbatchindex
func (m_ MatrixCopyToImage) SourceMatrixBatchIndex() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("sourceMatrixBatchIndex"))
	return rv
}/* debug [instance_properties/getter]: sourceMatrixBatchIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopytoimage/2976461-sourcematrixbatchindex
func (m_ MatrixCopyToImage) SetSourceMatrixBatchIndex(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceMatrixBatchIndex:"), value)
}/* debug [instance_properties/setter]: sourceMatrixBatchIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopytoimage/2976462-sourcematrixorigin
func (m_ MatrixCopyToImage) SourceMatrixOrigin() Origin get set /* not a class type */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("sourceMatrixOrigin"))
	return rv
}/* debug [instance_properties/getter]: sourceMatrixOrigin */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopytoimage/2976462-sourcematrixorigin
func (m_ MatrixCopyToImage) SetSourceMatrixOrigin(value Origin get set /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceMatrixOrigin:"), value)
}/* debug [instance_properties/setter]: sourceMatrixOrigin */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSMatrixCopyToImage */


