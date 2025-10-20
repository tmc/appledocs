// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MultiArray] class.
var (
	MultiArrayClass     _MultiArrayClass
	MultiArrayClassOnce sync.Once
)

func getMultiArrayClass() _MultiArrayClass {
	MultiArrayClassOnce.Do(func() {
		MultiArrayClass = _MultiArrayClass{objc.GetClass("MLMultiArray")}
	})
	return MultiArrayClass
}

type _MultiArrayClass struct {
	class objc.Class
}

// An interface definition for the [MultiArray] class.
type IMultiArray interface {
	objectivec.IObject
	GetBytesWithHandler(handler unsafe.Pointer)
	GetMutableBytesWithHandler(handler unsafe.Pointer)
	SetObjectAtIndexedSubscript(obj unsafe.Pointer, idx int)
	SetObjectForKeyedSubscript(obj unsafe.Pointer, key unsafe.Pointer)
	ObjectAtIndexedSubscript(idx int) unsafe.Pointer
	ObjectForKeyedSubscript(key unsafe.Pointer) unsafe.Pointer
	TransferToMultiArray(destinationMultiArray unsafe.Pointer)
}

// A machine learning collection type that stores numeric values in an array with multiple dimensions.
//
// A multidimensional array, or , is one of the underlying types of an that stores numeric values in multiple dimensions. All elements in an instance are one of the same type, and one of the types that defines: Each dimension in a multiarray is typically significant or meaningful. For example, a model could have an input that accepts images as a multiarray of pixels with three dimensions, C x H x W. The first dimension, ,_ _represents the number of color channels, and the second and third dimensions, and , represent the image’s height and width, respectively. The number of dimensions and size of each dimension define the multiarray’s . The property is an integer array that has an element for each dimension in the multiarray. Each element in defines the size of the corresponding dimension. To inspect the shape and constraints of a model’s multiarray input or output feature: Access the model’s property. Find the multiarray input or output feature in the model description’s or property, respectively. Access the feature description’s property. Inspect the multiarray constraint’s and .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray
type MultiArray struct {
	objectivec.Object
}

// MultiArrayFrom constructs a [MultiArray] from an unsafe.Pointer.
//
// A machine learning collection type that stores numeric values in an array with multiple dimensions.
func MultiArrayFrom(ptr unsafe.Pointer) MultiArray {
	return MultiArray{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MultiArrayClass) Alloc() MultiArray {
	rv := objc.Send[MultiArray](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MultiArrayClass) New() MultiArray {
	rv := objc.Send[MultiArray](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MultiArray) Init() MultiArray {
	rv := objc.Send[MultiArray](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MultiArray) Autorelease() MultiArray {
	rv := objc.Send[MultiArray](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMultiArray creates a new MultiArray instance.
func NewMultiArray() MultiArray {
	return getMultiArrayClass().New()
}


// Creates a multiarray from a data pointer.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/init(dataPointer:shape:dataType:strides:deallocator:)
func NewMultiArrayWithDataPointerShapeDataTypeStridesDeallocatorError(dataPointer unsafe.Pointer, shape unsafe.Pointer, dataType unsafe.Pointer, strides unsafe.Pointer, deallocator unsafe.Pointer, error unsafe.Pointer) MultiArray {
	instance := getMultiArrayClass().Alloc()
	rv := objc.Send[MultiArray](instance.ID, objc.Sel("initWithDataPointer:shape:dataType:strides:deallocator:error:"), dataPointer, shape, dataType, strides, deallocator, error)
	rv.Autorelease()
	return rv
}

// Creates a multiarray sharing the surface of a pixel buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/init(pixelBuffer:shape:)
func NewMultiArrayWithPixelBufferShape(pixelBuffer unsafe.Pointer, shape unsafe.Pointer) MultiArray {
	instance := getMultiArrayClass().Alloc()
	rv := objc.Send[MultiArray](instance.ID, objc.Sel("initWithPixelBuffer:shape:"), pixelBuffer, shape)
	rv.Autorelease()
	return rv
}

// Creates a multidimensional array with a shape and type.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/init(shape:dataType:)
func NewMultiArrayWithShapeDataTypeError(shape unsafe.Pointer, dataType unsafe.Pointer, error unsafe.Pointer) MultiArray {
	instance := getMultiArrayClass().Alloc()
	rv := objc.Send[MultiArray](instance.ID, objc.Sel("initWithShape:dataType:error:"), shape, dataType, error)
	rv.Autorelease()
	return rv
}

// Creates the object with specified strides.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/initWithShape:dataType:strides:
func NewMultiArrayWithShapeDataTypeStrides(shape unsafe.Pointer, dataType unsafe.Pointer, strides unsafe.Pointer) MultiArray {
	instance := getMultiArrayClass().Alloc()
	rv := objc.Send[MultiArray](instance.ID, objc.Sel("initWithShape:dataType:strides:"), shape, dataType, strides)
	rv.Autorelease()
	return rv
}

// Merges an array of multiarrays into one multiarray along an axis.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/init(byConcatenatingMultiArrays:alongAxis:dataType:)
func NewMultiArrayByConcatenatingMultiArraysAlongAxisDataType(multiArrays unsafe.Pointer, axis int, dataType unsafe.Pointer) MultiArray {
	rv := objc.Send[MultiArray](objc.ID(getMultiArrayClass().class), objc.Sel("multiArrayByConcatenatingMultiArrays:alongAxis:dataType:"), multiArrays, axis, dataType)
	return rv
}


// Merges an array of multiarrays into one multiarray along an axis.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/init(byConcatenatingMultiArrays:alongAxis:dataType:)
func (mc _MultiArrayClass) MultiArrayByConcatenatingMultiArraysAlongAxisDataType(multiArrays unsafe.Pointer, axis int, dataType unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("multiArrayByConcatenatingMultiArrays:alongAxis:dataType:"), multiArrays, axis, dataType)
	return rv
}

// Get the underlying buffer pointer to read.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/getBytesWithHandler:
func (m_ MultiArray) GetBytesWithHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("getBytesWithHandler:"), handler)
}

// Get the underlying buffer pointer to mutate.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/getMutableBytesWithHandler:
func (m_ MultiArray) GetMutableBytesWithHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("getMutableBytesWithHandler:"), handler)
}

// Assigns a number to the multiarray’s element at the location that the linear offset defines.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/setObject:atIndexedSubscript:
func (m_ MultiArray) SetObjectAtIndexedSubscript(obj unsafe.Pointer, idx int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setObject:atIndexedSubscript:"), obj, idx)
}

// Assigns a number to the multiarray’s element at the location that the number array defines.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/setObject:forKeyedSubscript:
func (m_ MultiArray) SetObjectForKeyedSubscript(obj unsafe.Pointer, key unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setObject:forKeyedSubscript:"), obj, key)
}

// Accesses the multiarray by using a linear offset.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/subscript(_:)-2hh91
func (m_ MultiArray) ObjectAtIndexedSubscript(idx int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("objectAtIndexedSubscript:"), idx)
	return rv
}

// Accesses the multiarray by using a number array that has an element for each dimension.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/subscript(_:)-3d9el
func (m_ MultiArray) ObjectForKeyedSubscript(key unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("objectForKeyedSubscript:"), key)
	return rv
}

// Transfer the contents to the destination multi-array.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/transfer(to:)
func (m_ MultiArray) TransferToMultiArray(destinationMultiArray unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("transferToMultiArray:"), destinationMultiArray)
}

// The total number of elements in the multiarray.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/count
func (m_ MultiArray) Count() int {
	rv := objc.Send[int](m_.ID, objc.Sel("count"))
	return rv
}

// The underlying type of the multiarray.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/dataType
func (m_ MultiArray) DataType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("dataType"))
	return rv
}

// A reference to the multiarray’s underlying pixel buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/pixelBuffer
func (m_ MultiArray) PixelBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("pixelBuffer"))
	return rv
}

// The multiarray’s multidimensional shape as a number array in which each element’s value is the size of the corresponding dimension.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/shape
func (m_ MultiArray) Shape() []unsafe.Pointer {
	rv := objc.Send[[]unsafe.Pointer](m_.ID, objc.Sel("shape"))
	return rv
}

// A number array in which each element is the number of memory locations that span the length of the corresponding dimension.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/strides
func (m_ MultiArray) Strides() []unsafe.Pointer {
	rv := objc.Send[[]unsafe.Pointer](m_.ID, objc.Sel("strides"))
	return rv
}


