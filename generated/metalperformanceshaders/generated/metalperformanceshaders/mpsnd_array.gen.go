// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNDArray */


/* debug [class_header]: Header for MPSNDArray */
// The class instance for the [NDArray] class.
var (
	NDArrayClass     _NDArrayClass
	NDArrayClassOnce sync.Once
)

func getNDArrayClass() _NDArrayClass {
	NDArrayClassOnce.Do(func() {
		NDArrayClass = _NDArrayClass{objc.GetClass("MPSNDArray")}
	})
	return NDArrayClass
}

type _NDArrayClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NDArray */
// An interface definition for the [NDArray] class.
type INDArray interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NDArray */
	// properties:
	DataType() DataType get /* not a class type */
	SetDataType(value DataType get /* not a class type */)
	DataTypeSize() objectivec.IObject
	SetDataTypeSize(value objectivec.IObject)
	Device() Device get /* not a class type */
	SetDevice(value Device get /* not a class type */)
	Label() objectivec.IObject
	SetLabel(value objectivec.IObject)
	NumberOfDimensions() objectivec.IObject
	SetNumberOfDimensions(value objectivec.IObject)
	Parent() IMPSNDArray
	SetParent(value IMPSNDArray)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NDArray */
	// methods:
	ArrayView()
	ArrayViewWithCommandBufferDescriptorAliasing(cmdBuf unsafe.Pointer, descriptor INDArrayDescriptor, aliasing AliasingStrategy) INDArray
	Descriptor()
	Length()
	LengthOfDimension(dimensionIndex uint) uint
	ReadBytes()
	ResourceSize()
	Synchronize()
	SynchronizeOnCommandBuffer(commandBuffer unsafe.Pointer)
	WriteBytes()
	ExportData()
	ExportDataWithCommandBufferToBufferDestinationDataTypeOffsetRowStrides(cmdBuf unsafe.Pointer, buffer unsafe.Pointer, destinationDataType DataType, offset uint, rowStrides int)
	ImportData()
	ImportDataWithCommandBufferFromBufferSourceDataTypeOffsetRowStrides(cmdBuf unsafe.Pointer, buffer unsafe.Pointer, sourceDataType DataType, offset uint, rowStrides int)
	ExportDataWithCommandBufferToImagesOffset(cmdBuf unsafe.Pointer, images ImageBatch /* not a class type */, offset objc.IObject /* cross-framework: MPSImageCoordinate */)
	ImportDataWithCommandBufferFromImagesOffset(cmdBuf unsafe.Pointer, images ImageBatch /* not a class type */, offset objc.IObject /* cross-framework: MPSImageCoordinate */)
	ArrayViewWithDimensionCountDimensionSizesStrides(numberOfDimensions uint, dimensionSizes uint, dimStrides uint) INDArray
	ArrayViewWithShapeStrides(shape Shape /* not a class type */, strides Shape /* not a class type */) INDArray
	UserBuffer()
	ArrayViewWithDescriptor(descriptor INDArrayDescriptor) INDArray
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NDArray */
// Alloc allocates a new instance without initialization.
func (nc _NDArrayClass) Alloc() NDArray {
	rv := objc.Send[NDArray](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NDArrayClass) New() NDArray {
	rv := objc.Send[NDArray](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArray) Init() NDArray {
	rv := objc.Send[NDArray](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArray) Autorelease() NDArray {
	rv := objc.Send[NDArray](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArray creates a new NDArray instance.
func NewNDArray() NDArray {
	return getNDArrayClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NDArray */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArray
type NDArray struct {
	objectivec.Object
}

// NDArrayFrom constructs a [NDArray] from an unsafe.Pointer.
func NDArrayFrom(ptr unsafe.Pointer) NDArray {
	return NDArray{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NDArray */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/4391636-initwithbuffer
func NewNDArrayWithBufferOffsetDescriptor(buffer unsafe.Pointer, offset uint, descriptor INDArrayDescriptor) NDArray {
	instance := getNDArrayClass().Alloc()
	rv := objc.Send[NDArray](instance.ID, objc.Sel("initWithBuffer:offset:descriptor:"), buffer, offset, descriptor)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNDArrayWithBufferOffsetDescriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114049-initwithdevice
func NewNDArrayWithDeviceDescriptor(device unsafe.Pointer, descriptor INDArrayDescriptor) NDArray {
	instance := getNDArrayClass().Alloc()
	rv := objc.Send[NDArray](instance.ID, objc.Sel("initWithDevice:descriptor:"), device, descriptor)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNDArrayWithDeviceDescriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114051-initwithdevice
func NewNDArrayWithDeviceScalar(device unsafe.Pointer, value float64) NDArray {
	instance := getNDArrayClass().Alloc()
	rv := objc.Send[NDArray](instance.ID, objc.Sel("initWithDevice:scalar:"), device, value)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNDArrayWithDeviceScalar */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NDArray */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3131728-defaultallocator
func (nc _NDArrayClass) DefaultAllocator() {
	objc.Send[objc.ID](objc.ID(nc.class), objc.Sel("defaultAllocator"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DefaultAllocator) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NDArray */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NDArray */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114040-arrayview
func (n_ NDArray) ArrayView() {
	objc.Send[objc.ID](n_.ID, objc.Sel("arrayView"))
}/* debug [instance_methods/method]: ArrayView */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114040-arrayviewwithcommandbuffer
func (n_ NDArray) ArrayViewWithCommandBufferDescriptorAliasing(cmdBuf unsafe.Pointer, descriptor INDArrayDescriptor, aliasing AliasingStrategy) INDArray {
	rv := objc.Send[NDArray](n_.ID, objc.Sel("arrayViewWithCommandBuffer:descriptor:aliasing:"), cmdBuf, descriptor, aliasing)
	return rv
}/* debug [instance_methods/method]: ArrayViewWithCommandBufferDescriptorAliasing */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114044-descriptor
func (n_ NDArray) Descriptor() {
	objc.Send[objc.ID](n_.ID, objc.Sel("descriptor"))
}/* debug [instance_methods/method]: Descriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114053-length
func (n_ NDArray) Length() {
	objc.Send[objc.ID](n_.ID, objc.Sel("length"))
}/* debug [instance_methods/method]: Length */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114053-lengthofdimension
func (n_ NDArray) LengthOfDimension(dimensionIndex uint) uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("lengthOfDimension:"), dimensionIndex)
	return rv
}/* debug [instance_methods/method]: LengthOfDimension */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114057-readbytes
func (n_ NDArray) ReadBytes() {
	objc.Send[objc.ID](n_.ID, objc.Sel("readBytes"))
}/* debug [instance_methods/method]: ReadBytes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114058-resourcesize
func (n_ NDArray) ResourceSize() {
	objc.Send[objc.ID](n_.ID, objc.Sel("resourceSize"))
}/* debug [instance_methods/method]: ResourceSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114059-synchronize
func (n_ NDArray) Synchronize() {
	objc.Send[objc.ID](n_.ID, objc.Sel("synchronize"))
}/* debug [instance_methods/method]: Synchronize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114059-synchronizeoncommandbuffer
func (n_ NDArray) SynchronizeOnCommandBuffer(commandBuffer unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("synchronizeOnCommandBuffer:"), commandBuffer)
}/* debug [instance_methods/method]: SynchronizeOnCommandBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114060-writebytes
func (n_ NDArray) WriteBytes() {
	objc.Send[objc.ID](n_.ID, objc.Sel("writeBytes"))
}/* debug [instance_methods/method]: WriteBytes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3131729-exportdata
func (n_ NDArray) ExportData() {
	objc.Send[objc.ID](n_.ID, objc.Sel("exportData"))
}/* debug [instance_methods/method]: ExportData */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3131729-exportdatawithcommandbuffer
func (n_ NDArray) ExportDataWithCommandBufferToBufferDestinationDataTypeOffsetRowStrides(cmdBuf unsafe.Pointer, buffer unsafe.Pointer, destinationDataType DataType, offset uint, rowStrides int) {
	objc.Send[objc.ID](n_.ID, objc.Sel("exportDataWithCommandBuffer:toBuffer:destinationDataType:offset:rowStrides:"), cmdBuf, buffer, destinationDataType, offset, rowStrides)
}/* debug [instance_methods/method]: ExportDataWithCommandBufferToBufferDestinationDataTypeOffsetRowStrides */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3131730-importdata
func (n_ NDArray) ImportData() {
	objc.Send[objc.ID](n_.ID, objc.Sel("importData"))
}/* debug [instance_methods/method]: ImportData */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3131730-importdatawithcommandbuffer
func (n_ NDArray) ImportDataWithCommandBufferFromBufferSourceDataTypeOffsetRowStrides(cmdBuf unsafe.Pointer, buffer unsafe.Pointer, sourceDataType DataType, offset uint, rowStrides int) {
	objc.Send[objc.ID](n_.ID, objc.Sel("importDataWithCommandBuffer:fromBuffer:sourceDataType:offset:rowStrides:"), cmdBuf, buffer, sourceDataType, offset, rowStrides)
}/* debug [instance_methods/method]: ImportDataWithCommandBufferFromBufferSourceDataTypeOffsetRowStrides */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3152526-exportdatawithcommandbuffer
func (n_ NDArray) ExportDataWithCommandBufferToImagesOffset(cmdBuf unsafe.Pointer, images ImageBatch /* not a class type */, offset objc.IObject /* cross-framework: MPSImageCoordinate */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("exportDataWithCommandBuffer:toImages:offset:"), cmdBuf, images, offset)
}/* debug [instance_methods/method]: ExportDataWithCommandBufferToImagesOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3152527-importdatawithcommandbuffer
func (n_ NDArray) ImportDataWithCommandBufferFromImagesOffset(cmdBuf unsafe.Pointer, images ImageBatch /* not a class type */, offset objc.IObject /* cross-framework: MPSImageCoordinate */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("importDataWithCommandBuffer:fromImages:offset:"), cmdBuf, images, offset)
}/* debug [instance_methods/method]: ImportDataWithCommandBufferFromImagesOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/4408693-arrayviewwithdimensioncount
func (n_ NDArray) ArrayViewWithDimensionCountDimensionSizesStrides(numberOfDimensions uint, dimensionSizes uint, dimStrides uint) INDArray {
	rv := objc.Send[NDArray](n_.ID, objc.Sel("arrayViewWithDimensionCount:dimensionSizes:strides:"), numberOfDimensions, dimensionSizes, dimStrides)
	return rv
}/* debug [instance_methods/method]: ArrayViewWithDimensionCountDimensionSizesStrides */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/4408694-arrayviewwithshape
func (n_ NDArray) ArrayViewWithShapeStrides(shape Shape /* not a class type */, strides Shape /* not a class type */) INDArray {
	rv := objc.Send[NDArray](n_.ID, objc.Sel("arrayViewWithShape:strides:"), shape, strides)
	return rv
}/* debug [instance_methods/method]: ArrayViewWithShapeStrides */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/4408695-userbuffer
func (n_ NDArray) UserBuffer() {
	objc.Send[objc.ID](n_.ID, objc.Sel("userBuffer"))
}/* debug [instance_methods/method]: UserBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/4438553-arrayviewwithdescriptor
func (n_ NDArray) ArrayViewWithDescriptor(descriptor INDArrayDescriptor) INDArray {
	rv := objc.Send[NDArray](n_.ID, objc.Sel("arrayViewWithDescriptor:"), descriptor)
	return rv
}/* debug [instance_methods/method]: ArrayViewWithDescriptor */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NDArray */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114041-datatype
func (n_ NDArray) DataType() DataType get /* not a class type */ {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("dataType"))
	return rv
}/* debug [instance_properties/getter]: dataType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114041-datatype
func (n_ NDArray) SetDataType(value DataType get /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDataType:"), value)
}/* debug [instance_properties/setter]: dataType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114042-datatypesize
func (n_ NDArray) DataTypeSize() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("dataTypeSize"))
	return rv
}/* debug [instance_properties/getter]: dataTypeSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114042-datatypesize
func (n_ NDArray) SetDataTypeSize(value objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDataTypeSize:"), value)
}/* debug [instance_properties/setter]: dataTypeSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114045-device
func (n_ NDArray) Device() Device get /* not a class type */ {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("device"))
	return rv
}/* debug [instance_properties/getter]: device */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114045-device
func (n_ NDArray) SetDevice(value Device get /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDevice:"), value)
}/* debug [instance_properties/setter]: device */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114052-label
func (n_ NDArray) Label() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114052-label
func (n_ NDArray) SetLabel(value objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114055-numberofdimensions
func (n_ NDArray) NumberOfDimensions() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("numberOfDimensions"))
	return rv
}/* debug [instance_properties/getter]: numberOfDimensions */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114055-numberofdimensions
func (n_ NDArray) SetNumberOfDimensions(value objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNumberOfDimensions:"), value)
}/* debug [instance_properties/setter]: numberOfDimensions */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114056-parent
func (n_ NDArray) Parent() IMPSNDArray {
	rv := objc.Send[NDArray](n_.ID, objc.Sel("parent"))
	return rv
}/* debug [instance_properties/getter]: parent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114056-parent
func (n_ NDArray) SetParent(value IMPSNDArray) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setParent:"), value)
}/* debug [instance_properties/setter]: parent */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNDArray */


