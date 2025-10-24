// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLMultiArray */


/* debug [class_header]: Header for MLMultiArray */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MultiArray */
// An interface definition for the [MultiArray] class.
type IMultiArray interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MultiArray */
	// properties:
	Count() int
	DataPointer() objectivec.IObject
	DataType() MultiArrayDataType
	PixelBuffer() PixelBufferRef /* not a class type */
	Shape() []foundation.Number
	Strides() []foundation.Number
	MultiArrayConstraint() IMLMultiArrayConstraint
	SetMultiArrayConstraint(value IMLMultiArrayConstraint)
	ModelDescription() IMLModelDescription
	SetModelDescription(value IMLModelDescription)
	InputDescriptionsByName() IMLFeatureDescription
	SetInputDescriptionsByName(value IMLFeatureDescription)
	OutputDescriptionsByName() IMLFeatureDescription
	SetOutputDescriptionsByName(value IMLFeatureDescription)
	ShapeConstraint() IMLMultiArrayShapeConstraint
	SetShapeConstraint(value IMLMultiArrayShapeConstraint)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MultiArray */
	// methods:
	GetBytesWithHandler(handler int)
	GetMutableBytesWithHandler(handler unsafe.Pointer)
	SetObjectAtIndexedSubscript(obj objc.IObject /* cross-framework: NSNumber */, idx int)
	SetObjectForKeyedSubscript(obj objc.IObject /* cross-framework: NSNumber */, key []foundation.Number)
	ObjectAtIndexedSubscript(idx int) foundation.Number
	ObjectForKeyedSubscript(key []foundation.Number) foundation.Number
	TransferToMultiArray(destinationMultiArray IMLMultiArray)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MultiArray */
// Alloc allocates a new instance without initialization.
func (mc _MultiArrayClass) Alloc() MultiArray {
	rv := objc.Send[MultiArray](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MultiArray */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MultiArray */

// Merges an array of multiarrays into one multiarray along an axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/init(byConcatenatingMultiArrays:alongAxis:dataType:)
func NewMultiArrayByConcatenatingMultiArraysAlongAxisDataType(multiArrays []MultiArray, axis int, dataType MultiArrayDataType) MultiArray {
	rv := objc.Send[MultiArray](objc.ID(getMultiArrayClass().class), objc.Sel("multiArrayByConcatenatingMultiArrays:alongAxis:dataType:"), multiArrays, axis, dataType)
	return rv
}/* debug [class_init_methods/constructor]: NewMultiArrayByConcatenatingMultiArraysAlongAxisDataType */


// Creates a multiarray from a data pointer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/init(dataPointer:shape:dataType:strides:deallocator:)
func NewMultiArrayWithDataPointerShapeDataTypeStridesDeallocatorError(dataPointer objectivec.IObject, shape []foundation.Number, dataType MultiArrayDataType, strides []foundation.Number, deallocator unsafe.Pointer, error_ objectivec.IObject) MultiArray {
	instance := getMultiArrayClass().Alloc()
	rv := objc.Send[MultiArray](instance.ID, objc.Sel("initWithDataPointer:shape:dataType:strides:deallocator:error:"), dataPointer, shape, dataType, strides, deallocator, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMultiArrayWithDataPointerShapeDataTypeStridesDeallocatorError */


// Creates a multiarray sharing the surface of a pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/init(pixelBuffer:shape:)
func NewMultiArrayWithPixelBufferShape(pixelBuffer PixelBufferRef /* not a class type */, shape []foundation.Number) MultiArray {
	instance := getMultiArrayClass().Alloc()
	rv := objc.Send[MultiArray](instance.ID, objc.Sel("initWithPixelBuffer:shape:"), pixelBuffer, shape)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMultiArrayWithPixelBufferShape */


// Creates a multidimensional array with a shape and type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/init(shape:dataType:)
func NewMultiArrayWithShapeDataTypeError(shape []foundation.Number, dataType MultiArrayDataType, error_ objectivec.IObject) MultiArray {
	instance := getMultiArrayClass().Alloc()
	rv := objc.Send[MultiArray](instance.ID, objc.Sel("initWithShape:dataType:error:"), shape, dataType, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMultiArrayWithShapeDataTypeError */


// Creates the object with specified strides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/initWithShape:dataType:strides:
func NewMultiArrayWithShapeDataTypeStrides(shape []foundation.Number, dataType MultiArrayDataType, strides []foundation.Number) MultiArray {
	instance := getMultiArrayClass().Alloc()
	rv := objc.Send[MultiArray](instance.ID, objc.Sel("initWithShape:dataType:strides:"), shape, dataType, strides)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMultiArrayWithShapeDataTypeStrides */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MultiArray */

// Merges an array of multiarrays into one multiarray along an axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/init(byConcatenatingMultiArrays:alongAxis:dataType:)
func (mc _MultiArrayClass) MultiArrayByConcatenatingMultiArraysAlongAxisDataType(multiArrays []MultiArray, axis int, dataType MultiArrayDataType) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("multiArrayByConcatenatingMultiArrays:alongAxis:dataType:"), multiArrays, axis, dataType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MultiArrayByConcatenatingMultiArraysAlongAxisDataType) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MultiArray */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MultiArray */

// Get the underlying buffer pointer to read.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/getBytesWithHandler:
func (m_ MultiArray) GetBytesWithHandler(handler int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("getBytesWithHandler:"), handler)
}/* debug [instance_methods/method]: GetBytesWithHandler */


// Get the underlying buffer pointer to mutate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/getMutableBytesWithHandler:
func (m_ MultiArray) GetMutableBytesWithHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("getMutableBytesWithHandler:"), handler)
}/* debug [instance_methods/method]: GetMutableBytesWithHandler */


// Assigns a number to the multiarray’s element at the location that the linear offset defines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/setObject:atIndexedSubscript:
func (m_ MultiArray) SetObjectAtIndexedSubscript(obj objc.IObject /* cross-framework: NSNumber */, idx int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setObject:atIndexedSubscript:"), obj, idx)
}/* debug [instance_methods/method]: SetObjectAtIndexedSubscript */


// Assigns a number to the multiarray’s element at the location that the number array defines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/setObject:forKeyedSubscript:
func (m_ MultiArray) SetObjectForKeyedSubscript(obj objc.IObject /* cross-framework: NSNumber */, key []foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setObject:forKeyedSubscript:"), obj, key)
}/* debug [instance_methods/method]: SetObjectForKeyedSubscript */


// Accesses the multiarray by using a linear offset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/subscript(_:)-2hh91
func (m_ MultiArray) ObjectAtIndexedSubscript(idx int) foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("objectAtIndexedSubscript:"), idx)
	return rv
}/* debug [instance_methods/method]: ObjectAtIndexedSubscript */


// Accesses the multiarray by using a number array that has an element for each dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/subscript(_:)-3d9el
func (m_ MultiArray) ObjectForKeyedSubscript(key []foundation.Number) foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("objectForKeyedSubscript:"), key)
	return rv
}/* debug [instance_methods/method]: ObjectForKeyedSubscript */


// Transfer the contents to the destination multi-array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/transfer(to:)
func (m_ MultiArray) TransferToMultiArray(destinationMultiArray IMLMultiArray) {
	objc.Send[objc.ID](m_.ID, objc.Sel("transferToMultiArray:"), destinationMultiArray)
}/* debug [instance_methods/method]: TransferToMultiArray */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MultiArray */

// The total number of elements in the multiarray.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/count
func (m_ MultiArray) Count() int {
	rv := objc.Send[int](m_.ID, objc.Sel("count"))
	return rv
}/* debug [instance_properties/getter]: count */


// A pointer to the multiarray’s underlying memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/dataPointer
func (m_ MultiArray) DataPointer() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("dataPointer"))
	return rv
}/* debug [instance_properties/getter]: dataPointer */


// The underlying type of the multiarray.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/dataType
func (m_ MultiArray) DataType() MultiArrayDataType {
	rv := objc.Send[MultiArrayDataType](m_.ID, objc.Sel("dataType"))
	return rv
}/* debug [instance_properties/getter]: dataType */


// A reference to the multiarray’s underlying pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/pixelBuffer
func (m_ MultiArray) PixelBuffer() PixelBufferRef /* not a class type */ {
	rv := objc.Send[PixelBufferRef](m_.ID, objc.Sel("pixelBuffer"))
	return rv
}/* debug [instance_properties/getter]: pixelBuffer */


// The multiarray’s multidimensional shape as a number array in which each element’s value is the size of the corresponding dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/shape
func (m_ MultiArray) Shape() []foundation.Number {
	rv := objc.Send[[]foundation.Number](m_.ID, objc.Sel("shape"))
	return rv
}/* debug [instance_properties/getter]: shape */


// A number array in which each element is the number of memory locations that span the length of the corresponding dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArray/strides
func (m_ MultiArray) Strides() []foundation.Number {
	rv := objc.Send[[]foundation.Number](m_.ID, objc.Sel("strides"))
	return rv
}/* debug [instance_properties/getter]: strides */


// The constraints on a multidimensional array feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/multiarrayconstraint
func (m_ MultiArray) MultiArrayConstraint() IMLMultiArrayConstraint {
	rv := objc.Send[MultiArrayConstraint](m_.ID, objc.Sel("multiArrayConstraint"))
	return rv
}/* debug [instance_properties/getter]: multiArrayConstraint */


// The constraints on a multidimensional array feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/multiarrayconstraint
func (m_ MultiArray) SetMultiArrayConstraint(value IMLMultiArrayConstraint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMultiArrayConstraint:"), value)
}/* debug [instance_properties/setter]: multiArrayConstraint */


// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodel/modeldescription
func (m_ MultiArray) ModelDescription() IMLModelDescription {
	rv := objc.Send[ModelDescription](m_.ID, objc.Sel("modelDescription"))
	return rv
}/* debug [instance_properties/getter]: modelDescription */


// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodel/modeldescription
func (m_ MultiArray) SetModelDescription(value IMLModelDescription) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setModelDescription:"), value)
}/* debug [instance_properties/setter]: modelDescription */


// A dictionary of input feature descriptions, which the model keys by the input’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/inputdescriptionsbyname
func (m_ MultiArray) InputDescriptionsByName() IMLFeatureDescription {
	rv := objc.Send[FeatureDescription](m_.ID, objc.Sel("inputDescriptionsByName"))
	return rv
}/* debug [instance_properties/getter]: inputDescriptionsByName */


// A dictionary of input feature descriptions, which the model keys by the input’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/inputdescriptionsbyname
func (m_ MultiArray) SetInputDescriptionsByName(value IMLFeatureDescription) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInputDescriptionsByName:"), value)
}/* debug [instance_properties/setter]: inputDescriptionsByName */


// A dictionary of output feature descriptions, which the model keys by the output’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/outputdescriptionsbyname
func (m_ MultiArray) OutputDescriptionsByName() IMLFeatureDescription {
	rv := objc.Send[FeatureDescription](m_.ID, objc.Sel("outputDescriptionsByName"))
	return rv
}/* debug [instance_properties/getter]: outputDescriptionsByName */


// A dictionary of output feature descriptions, which the model keys by the output’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/outputdescriptionsbyname
func (m_ MultiArray) SetOutputDescriptionsByName(value IMLFeatureDescription) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOutputDescriptionsByName:"), value)
}/* debug [instance_properties/setter]: outputDescriptionsByName */


// The constraint on the shape of the multiarray.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarrayconstraint/shapeconstraint
func (m_ MultiArray) ShapeConstraint() IMLMultiArrayShapeConstraint {
	rv := objc.Send[MultiArrayShapeConstraint](m_.ID, objc.Sel("shapeConstraint"))
	return rv
}/* debug [instance_properties/getter]: shapeConstraint */


// The constraint on the shape of the multiarray.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarrayconstraint/shapeconstraint
func (m_ MultiArray) SetShapeConstraint(value IMLMultiArrayShapeConstraint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShapeConstraint:"), value)
}/* debug [instance_properties/setter]: shapeConstraint */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLMultiArray */


