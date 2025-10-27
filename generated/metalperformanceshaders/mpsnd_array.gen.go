// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [NDArray] class.
type INDArray interface {
	objectivec.IObject
	

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
	ExportDataWithCommandBufferToImagesOffset(cmdBuf unsafe.Pointer, images ImageBatch /* not a class type */, offset ImageCoordinate)
	ImportDataWithCommandBufferFromImagesOffset(cmdBuf unsafe.Pointer, images ImageBatch /* not a class type */, offset ImageCoordinate)
	ArrayViewWithDimensionCountDimensionSizesStrides(numberOfDimensions uint, dimensionSizes uint, dimStrides uint) INDArray
	ArrayViewWithShapeStrides(shape Shape /* not a class type */, strides Shape /* not a class type */) INDArray
	UserBuffer()
	ArrayViewWithDescriptor(descriptor INDArrayDescriptor) INDArray


}





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







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArray
type NDArray struct {
	objectivec.Object
}

// NDArrayFrom constructs a [NDArray] from an unsafe.Pointer.
func NDArrayFrom(ptr unsafe.Pointer) NDArray {
	return NDArray{objectivec.Object{objc.ID(ptr)}}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/4391636-initwithbuffer
func NewNDArrayWithBufferOffsetDescriptor(buffer unsafe.Pointer, offset uint, descriptor INDArrayDescriptor) NDArray {
	instance := getNDArrayClass().Alloc()
	rv := objc.Send[NDArray](instance.ID, objc.Sel("initWithBuffer:offset:descriptor:"), buffer, offset, descriptor)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114049-initwithdevice
func NewNDArrayWithDeviceDescriptor(device unsafe.Pointer, descriptor INDArrayDescriptor) NDArray {
	instance := getNDArrayClass().Alloc()
	rv := objc.Send[NDArray](instance.ID, objc.Sel("initWithDevice:descriptor:"), device, descriptor)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114051-initwithdevice
func NewNDArrayWithDeviceScalar(device unsafe.Pointer, value float64) NDArray {
	instance := getNDArrayClass().Alloc()
	rv := objc.Send[NDArray](instance.ID, objc.Sel("initWithDevice:scalar:"), device, value)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3131728-defaultallocator
func (nc _NDArrayClass) DefaultAllocator() {
	objc.Send[objc.ID](objc.ID(nc.class), objc.Sel("defaultAllocator"))
}












// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114040-arrayview
func (n_ NDArray) ArrayView() {
	objc.Send[objc.ID](n_.ID, objc.Sel("arrayView"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114040-arrayviewwithcommandbuffer
func (n_ NDArray) ArrayViewWithCommandBufferDescriptorAliasing(cmdBuf unsafe.Pointer, descriptor INDArrayDescriptor, aliasing AliasingStrategy) INDArray {
	rv := objc.Send[NDArray](n_.ID, objc.Sel("arrayViewWithCommandBuffer:descriptor:aliasing:"), cmdBuf, descriptor, aliasing)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114044-descriptor
func (n_ NDArray) Descriptor() {
	objc.Send[objc.ID](n_.ID, objc.Sel("descriptor"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114053-length
func (n_ NDArray) Length() {
	objc.Send[objc.ID](n_.ID, objc.Sel("length"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114053-lengthofdimension
func (n_ NDArray) LengthOfDimension(dimensionIndex uint) uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("lengthOfDimension:"), dimensionIndex)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114057-readbytes
func (n_ NDArray) ReadBytes() {
	objc.Send[objc.ID](n_.ID, objc.Sel("readBytes"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114058-resourcesize
func (n_ NDArray) ResourceSize() {
	objc.Send[objc.ID](n_.ID, objc.Sel("resourceSize"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114059-synchronize
func (n_ NDArray) Synchronize() {
	objc.Send[objc.ID](n_.ID, objc.Sel("synchronize"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114059-synchronizeoncommandbuffer
func (n_ NDArray) SynchronizeOnCommandBuffer(commandBuffer unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("synchronizeOnCommandBuffer:"), commandBuffer)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114060-writebytes
func (n_ NDArray) WriteBytes() {
	objc.Send[objc.ID](n_.ID, objc.Sel("writeBytes"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3131729-exportdata
func (n_ NDArray) ExportData() {
	objc.Send[objc.ID](n_.ID, objc.Sel("exportData"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3131729-exportdatawithcommandbuffer
func (n_ NDArray) ExportDataWithCommandBufferToBufferDestinationDataTypeOffsetRowStrides(cmdBuf unsafe.Pointer, buffer unsafe.Pointer, destinationDataType DataType, offset uint, rowStrides int) {
	objc.Send[objc.ID](n_.ID, objc.Sel("exportDataWithCommandBuffer:toBuffer:destinationDataType:offset:rowStrides:"), cmdBuf, buffer, destinationDataType, offset, rowStrides)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3131730-importdata
func (n_ NDArray) ImportData() {
	objc.Send[objc.ID](n_.ID, objc.Sel("importData"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3131730-importdatawithcommandbuffer
func (n_ NDArray) ImportDataWithCommandBufferFromBufferSourceDataTypeOffsetRowStrides(cmdBuf unsafe.Pointer, buffer unsafe.Pointer, sourceDataType DataType, offset uint, rowStrides int) {
	objc.Send[objc.ID](n_.ID, objc.Sel("importDataWithCommandBuffer:fromBuffer:sourceDataType:offset:rowStrides:"), cmdBuf, buffer, sourceDataType, offset, rowStrides)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3152526-exportdatawithcommandbuffer
func (n_ NDArray) ExportDataWithCommandBufferToImagesOffset(cmdBuf unsafe.Pointer, images ImageBatch /* not a class type */, offset ImageCoordinate) {
	objc.Send[objc.ID](n_.ID, objc.Sel("exportDataWithCommandBuffer:toImages:offset:"), cmdBuf, images, offset)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3152527-importdatawithcommandbuffer
func (n_ NDArray) ImportDataWithCommandBufferFromImagesOffset(cmdBuf unsafe.Pointer, images ImageBatch /* not a class type */, offset ImageCoordinate) {
	objc.Send[objc.ID](n_.ID, objc.Sel("importDataWithCommandBuffer:fromImages:offset:"), cmdBuf, images, offset)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/4408693-arrayviewwithdimensioncount
func (n_ NDArray) ArrayViewWithDimensionCountDimensionSizesStrides(numberOfDimensions uint, dimensionSizes uint, dimStrides uint) INDArray {
	rv := objc.Send[NDArray](n_.ID, objc.Sel("arrayViewWithDimensionCount:dimensionSizes:strides:"), numberOfDimensions, dimensionSizes, dimStrides)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/4408694-arrayviewwithshape
func (n_ NDArray) ArrayViewWithShapeStrides(shape Shape /* not a class type */, strides Shape /* not a class type */) INDArray {
	rv := objc.Send[NDArray](n_.ID, objc.Sel("arrayViewWithShape:strides:"), shape, strides)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/4408695-userbuffer
func (n_ NDArray) UserBuffer() {
	objc.Send[objc.ID](n_.ID, objc.Sel("userBuffer"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/4438553-arrayviewwithdescriptor
func (n_ NDArray) ArrayViewWithDescriptor(descriptor INDArrayDescriptor) INDArray {
	rv := objc.Send[NDArray](n_.ID, objc.Sel("arrayViewWithDescriptor:"), descriptor)
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114041-datatype
func (n_ NDArray) DataType() DataType get /* not a class type */ {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("dataType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114041-datatype
func (n_ NDArray) SetDataType(value DataType get /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDataType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114042-datatypesize
func (n_ NDArray) DataTypeSize() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("dataTypeSize"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114042-datatypesize
func (n_ NDArray) SetDataTypeSize(value objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDataTypeSize:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114045-device
func (n_ NDArray) Device() Device get /* not a class type */ {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("device"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114045-device
func (n_ NDArray) SetDevice(value Device get /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDevice:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114052-label
func (n_ NDArray) Label() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("label"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114052-label
func (n_ NDArray) SetLabel(value objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setLabel:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114055-numberofdimensions
func (n_ NDArray) NumberOfDimensions() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("numberOfDimensions"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114055-numberofdimensions
func (n_ NDArray) SetNumberOfDimensions(value objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNumberOfDimensions:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114056-parent
func (n_ NDArray) Parent() IMPSNDArray {
	rv := objc.Send[NDArray](n_.ID, objc.Sel("parent"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114056-parent
func (n_ NDArray) SetParent(value IMPSNDArray) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setParent:"), value)
}







