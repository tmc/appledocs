// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	Count() int
	DataPointer() unsafe.Pointer
	DataType() MultiArrayDataType
	PixelBuffer() PixelBufferRef /* not a class type */
	Shape() []objc.IObject /* cross-framework: Number */
	Strides() []objc.IObject /* cross-framework: Number */
	MultiArrayConstraint() objc.IObject /* cross-framework: MultiArrayConstraint */
	SetMultiArrayConstraint(value objc.IObject /* cross-framework: MultiArrayConstraint */)
	ModelDescription() IMLModelDescription
	SetModelDescription(value IMLModelDescription)
	InputDescriptionsByName() IMLFeatureDescription
	SetInputDescriptionsByName(value IMLFeatureDescription)
	OutputDescriptionsByName() IMLFeatureDescription
	SetOutputDescriptionsByName(value IMLFeatureDescription)
	ShapeConstraint() objc.IObject /* cross-framework: MultiArrayShapeConstraint */
	SetShapeConstraint(value objc.IObject /* cross-framework: MultiArrayShapeConstraint */)
	// methods:
	GetBytesWithHandler(handler int)
	GetMutableBytesWithHandler(handler unsafe.Pointer)
	SetObjectAtIndexedSubscript(obj objc.IObject /* cross-framework: NSNumber */, idx int)
	SetObjectForKeyedSubscript(obj objc.IObject /* cross-framework: NSNumber */, key []objc.IObject /* cross-framework: Number */)
	TransferToMultiArray(destinationMultiArray IMLMultiArray)
}

// A machine learning collection type that stores numeric values in an array with multiple dimensions.
//
// A multidimensional array, or , is one of the underlying types of an that stores numeric values in multiple dimensions. All elements in an instance are one of the same type, and one of the types that defines: Each dimension in a multiarray is typically significant or meaningful. For example, a model could have an input that accepts images as a multiarray of pixels with three dimensions, C x H x W. The first dimension, ,_ _represents the number of color channels, and the second and third dimensions, and , represent the image’s height and width, respectively. The number of dimensions and size of each dimension define the multiarray’s . The property is an integer array that has an element for each dimension in the multiarray. Each element in defines the size of the corresponding dimension. To inspect the shape and constraints of a model’s multiarray input or output feature: Access the model’s property. Find the multiarray input or output feature in the model description’s or property, respectively. Access the feature description’s property. Inspect the multiarray constraint’s and .


// A machine learning collection type that stores numeric values in an array with multiple dimensions.
//
// [Full Topic]
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



// Merges an array of multiarrays into one multiarray along an axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/init(byConcatenatingMultiArrays:alongAxis:dataType:)
func NewMultiArrayByConcatenatingMultiArraysAlongAxisDataType(multiArrays []IMultiArray, axis int, dataType MultiArrayDataType) MultiArray {
	rv := objc.Send[MultiArray](objc.ID(getMultiArrayClass().class), objc.Sel("multiArrayByConcatenatingMultiArrays:alongAxis:dataType:"), multiArrays, axis, dataType)
	return rv
}


// Creates a multiarray from a data pointer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/init(dataPointer:shape:dataType:strides:deallocator:)
func NewMultiArrayWithDataPointerShapeDataTypeStridesDeallocatorError(dataPointer unsafe.Pointer, shape []objc.IObject /* cross-framework: Number */, dataType MultiArrayDataType, strides []objc.IObject /* cross-framework: Number */, deallocator unsafe.Pointer, error_ unsafe.Pointer) MultiArray {
	instance := getMultiArrayClass().Alloc()
	rv := objc.Send[MultiArray](instance.ID, objc.Sel("initWithDataPointer:shape:dataType:strides:deallocator:error:"), dataPointer, shape, dataType, strides, deallocator, error_)
	rv.Autorelease()
	return rv
}


// Creates a multiarray sharing the surface of a pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/init(pixelBuffer:shape:)
func NewMultiArrayWithPixelBufferShape(pixelBuffer PixelBufferRef /* not a class type */, shape []objc.IObject /* cross-framework: Number */) MultiArray {
	instance := getMultiArrayClass().Alloc()
	rv := objc.Send[MultiArray](instance.ID, objc.Sel("initWithPixelBuffer:shape:"), pixelBuffer, shape)
	rv.Autorelease()
	return rv
}


// Creates a multidimensional array with a shape and type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/init(shape:dataType:)
func NewMultiArrayWithShapeDataTypeError(shape []objc.IObject /* cross-framework: Number */, dataType MultiArrayDataType, error_ unsafe.Pointer) MultiArray {
	instance := getMultiArrayClass().Alloc()
	rv := objc.Send[MultiArray](instance.ID, objc.Sel("initWithShape:dataType:error:"), shape, dataType, error_)
	rv.Autorelease()
	return rv
}


// Creates the object with specified strides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/initWithShape:dataType:strides:
func NewMultiArrayWithShapeDataTypeStrides(shape []objc.IObject /* cross-framework: Number */, dataType MultiArrayDataType, strides []objc.IObject /* cross-framework: Number */) MultiArray {
	instance := getMultiArrayClass().Alloc()
	rv := objc.Send[MultiArray](instance.ID, objc.Sel("initWithShape:dataType:strides:"), shape, dataType, strides)
	rv.Autorelease()
	return rv
}



// Merges an array of multiarrays into one multiarray along an axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/init(byConcatenatingMultiArrays:alongAxis:dataType:)
func (mc _MultiArrayClass) MultiArrayByConcatenatingMultiArraysAlongAxisDataType(multiArrays []IMultiArray, axis int, dataType MultiArrayDataType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("multiArrayByConcatenatingMultiArrays:alongAxis:dataType:"), multiArrays, axis, dataType)
	return rv
}


// Get the underlying buffer pointer to read.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/getBytesWithHandler:
func (m_ MultiArray) GetBytesWithHandler(handler int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("getBytesWithHandler:"), handler)
}


// Get the underlying buffer pointer to mutate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/getMutableBytesWithHandler:
func (m_ MultiArray) GetMutableBytesWithHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("getMutableBytesWithHandler:"), handler)
}


// Assigns a number to the multiarray’s element at the location that the linear offset defines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/setObject:atIndexedSubscript:
func (m_ MultiArray) SetObjectAtIndexedSubscript(obj objc.IObject /* cross-framework: NSNumber */, idx int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setObject:atIndexedSubscript:"), obj, idx)
}


// Assigns a number to the multiarray’s element at the location that the number array defines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/setObject:forKeyedSubscript:
func (m_ MultiArray) SetObjectForKeyedSubscript(obj objc.IObject /* cross-framework: NSNumber */, key []objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setObject:forKeyedSubscript:"), obj, key)
}


// Transfer the contents to the destination multi-array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/transfer(to:)
func (m_ MultiArray) TransferToMultiArray(destinationMultiArray IMLMultiArray) {
	objc.Send[objc.ID](m_.ID, objc.Sel("transferToMultiArray:"), destinationMultiArray)
}


// The total number of elements in the multiarray.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/count
func (m_ MultiArray) Count() int {
	rv := objc.Send[int](m_.ID, objc.Sel("count"))
	return rv
}


// A pointer to the multiarray’s underlying memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/dataPointer
func (m_ MultiArray) DataPointer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("dataPointer"))
	return rv
}


// The underlying type of the multiarray.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/dataType
func (m_ MultiArray) DataType() MultiArrayDataType {
	rv := objc.Send[MultiArrayDataType](m_.ID, objc.Sel("dataType"))
	return rv
}


// A reference to the multiarray’s underlying pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/pixelBuffer
func (m_ MultiArray) PixelBuffer() PixelBufferRef /* not a class type */ {
	rv := objc.Send[PixelBufferRef](m_.ID, objc.Sel("pixelBuffer"))
	return rv
}


// The multiarray’s multidimensional shape as a number array in which each element’s value is the size of the corresponding dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/shape
func (m_ MultiArray) Shape() []objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[[]foundation.Number](m_.ID, objc.Sel("shape"))
	return rv
}


// A number array in which each element is the number of memory locations that span the length of the corresponding dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/strides
func (m_ MultiArray) Strides() []objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[[]foundation.Number](m_.ID, objc.Sel("strides"))
	return rv
}


// The constraints on a multidimensional array feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/multiarrayconstraint
func (m_ MultiArray) MultiArrayConstraint() objc.IObject /* cross-framework: MultiArrayConstraint */ {
	rv := objc.Send[MultiArrayConstraint](m_.ID, objc.Sel("multiArrayConstraint"))
	return rv
}


// The constraints on a multidimensional array feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/multiarrayconstraint
func (m_ MultiArray) SetMultiArrayConstraint(value objc.IObject /* cross-framework: MultiArrayConstraint */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMultiArrayConstraint:"), value)
}


// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodel/modeldescription
func (m_ MultiArray) ModelDescription() IMLModelDescription {
	rv := objc.Send[ModelDescription](m_.ID, objc.Sel("modelDescription"))
	return rv
}


// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodel/modeldescription
func (m_ MultiArray) SetModelDescription(value IMLModelDescription) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setModelDescription:"), value)
}


// A dictionary of input feature descriptions, which the model keys by the input’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/inputdescriptionsbyname
func (m_ MultiArray) InputDescriptionsByName() IMLFeatureDescription {
	rv := objc.Send[FeatureDescription](m_.ID, objc.Sel("inputDescriptionsByName"))
	return rv
}


// A dictionary of input feature descriptions, which the model keys by the input’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/inputdescriptionsbyname
func (m_ MultiArray) SetInputDescriptionsByName(value IMLFeatureDescription) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInputDescriptionsByName:"), value)
}


// A dictionary of output feature descriptions, which the model keys by the output’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/outputdescriptionsbyname
func (m_ MultiArray) OutputDescriptionsByName() IMLFeatureDescription {
	rv := objc.Send[FeatureDescription](m_.ID, objc.Sel("outputDescriptionsByName"))
	return rv
}


// A dictionary of output feature descriptions, which the model keys by the output’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/outputdescriptionsbyname
func (m_ MultiArray) SetOutputDescriptionsByName(value IMLFeatureDescription) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOutputDescriptionsByName:"), value)
}


// The constraint on the shape of the multiarray.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarrayconstraint/shapeconstraint
func (m_ MultiArray) ShapeConstraint() objc.IObject /* cross-framework: MultiArrayShapeConstraint */ {
	rv := objc.Send[MultiArrayShapeConstraint](m_.ID, objc.Sel("shapeConstraint"))
	return rv
}


// The constraint on the shape of the multiarray.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarrayconstraint/shapeconstraint
func (m_ MultiArray) SetShapeConstraint(value objc.IObject /* cross-framework: MultiArrayShapeConstraint */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShapeConstraint:"), value)
}


