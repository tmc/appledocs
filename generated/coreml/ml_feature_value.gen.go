// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLFeatureValue */


/* debug [class_header]: Header for MLFeatureValue */
// The class instance for the [FeatureValue] class.
var (
	FeatureValueClass     _FeatureValueClass
	FeatureValueClassOnce sync.Once
)

func getFeatureValueClass() _FeatureValueClass {
	FeatureValueClassOnce.Do(func() {
		FeatureValueClass = _FeatureValueClass{objc.GetClass("MLFeatureValue")}
	})
	return FeatureValueClass
}

type _FeatureValueClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FeatureValue */
// An interface definition for the [FeatureValue] class.
type IFeatureValue interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FeatureValue */
	// properties:
	DictionaryValue() foundation.IDictionary
	DoubleValue() float64
	ImageBufferValue() PixelBufferRef /* not a class type */
	Int64Value() int64
	Undefined() bool
	MultiArrayValue() IMLMultiArray
	SequenceValue() IMLSequence
	StringValue() objc.IObject /* cross-framework: NSString */
	Type() FeatureType
	IsUndefined() bool
	SetIsUndefined(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FeatureValue */
	// methods:
	IsEqualToFeatureValue(value IMLFeatureValue) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FeatureValue */
// Alloc allocates a new instance without initialization.
func (fc _FeatureValueClass) Alloc() FeatureValue {
	rv := objc.Send[FeatureValue](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FeatureValueClass) New() FeatureValue {
	rv := objc.Send[FeatureValue](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FeatureValue) Init() FeatureValue {
	rv := objc.Send[FeatureValue](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FeatureValue) Autorelease() FeatureValue {
	rv := objc.Send[FeatureValue](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFeatureValue creates a new FeatureValue instance.
func NewFeatureValue() FeatureValue {
	return getFeatureValueClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FeatureValue */
// A generic wrapper around an underlying value and the value’s type.
//
// A Core ML wraps an underlying value and bundles it with that value’s type, which is one of the types that defines. Apps typically access feature values indirectly by using the methods in the wrapper class Xcode automatically generates for Core ML model files. If your app accesses an directly, it must create and consume instances. For each prediction, Core ML accepts a feature provider for its inputs, and generates a separate feature provider for its outputs. The input feature provider contains one instance per input, and the output feature provider contains one per output. See for more information about the model input and output features.


// A generic wrapper around an underlying value and the value’s type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue
type FeatureValue struct {
	objectivec.Object
}

// FeatureValueFrom constructs a [FeatureValue] from an unsafe.Pointer.
//
// A generic wrapper around an underlying value and the value’s type.
func FeatureValueFrom(ptr unsafe.Pointer) FeatureValue {
	return FeatureValue{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FeatureValue */

// Creates a feature value with a type that represents an undefined or missing value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(undefined:)
func NewFeatureValueUndefinedFeatureValueWithType(type_ FeatureType) FeatureValue {
	rv := objc.Send[FeatureValue](objc.ID(getFeatureValueClass().class), objc.Sel("undefinedFeatureValueWithType:"), type_)
	return rv
}/* debug [class_init_methods/constructor]: NewFeatureValueUndefinedFeatureValueWithType */


// Creates a feature value that contains an image defined by a core graphics image and a constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(CGImage:constraint:options:)
func NewFeatureValueWithCGImageConstraintOptionsError(cgImage ImageRef /* not a class type */, constraint IMLImageConstraint, options foundation.IDictionary, error_ objectivec.IObject) FeatureValue {
	rv := objc.Send[FeatureValue](objc.ID(getFeatureValueClass().class), objc.Sel("featureValueWithCGImage:constraint:options:error:"), cgImage, constraint, options, error_)
	return rv
}/* debug [class_init_methods/constructor]: NewFeatureValueWithCGImageConstraintOptionsError */


// Creates a feature value that contains an image defined by a core graphics image, an orientation, and a constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(CGImage:orientation:constraint:options:)
func NewFeatureValueWithCGImageOrientationConstraintOptionsError(cgImage ImageRef /* not a class type */, orientation ImagePropertyOrientation /* not a class type */, constraint IMLImageConstraint, options foundation.IDictionary, error_ objectivec.IObject) FeatureValue {
	rv := objc.Send[FeatureValue](objc.ID(getFeatureValueClass().class), objc.Sel("featureValueWithCGImage:orientation:constraint:options:error:"), cgImage, orientation, constraint, options, error_)
	return rv
}/* debug [class_init_methods/constructor]: NewFeatureValueWithCGImageOrientationConstraintOptionsError */


// Creates a feature value that contains an image defined by a core graphics image and its orientation, size, and pixel format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(CGImage:orientation:pixelsWide:pixelsHigh:pixelFormatType:options:)
func NewFeatureValueWithCGImageOrientationPixelsWidePixelsHighPixelFormatTypeOptionsError(cgImage ImageRef /* not a class type */, orientation ImagePropertyOrientation /* not a class type */, pixelsWide int, pixelsHigh int, pixelFormatType uint32 /* not a class type */, options foundation.IDictionary, error_ objectivec.IObject) FeatureValue {
	rv := objc.Send[FeatureValue](objc.ID(getFeatureValueClass().class), objc.Sel("featureValueWithCGImage:orientation:pixelsWide:pixelsHigh:pixelFormatType:options:error:"), cgImage, orientation, pixelsWide, pixelsHigh, pixelFormatType, options, error_)
	return rv
}/* debug [class_init_methods/constructor]: NewFeatureValueWithCGImageOrientationPixelsWidePixelsHighPixelFormatTypeOptionsError */


// Creates a feature value that contains an image defined by a core graphics image and its size and pixel format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(CGImage:pixelsWide:pixelsHigh:pixelFormatType:options:)
func NewFeatureValueWithCGImagePixelsWidePixelsHighPixelFormatTypeOptionsError(cgImage ImageRef /* not a class type */, pixelsWide int, pixelsHigh int, pixelFormatType uint32 /* not a class type */, options foundation.IDictionary, error_ objectivec.IObject) FeatureValue {
	rv := objc.Send[FeatureValue](objc.ID(getFeatureValueClass().class), objc.Sel("featureValueWithCGImage:pixelsWide:pixelsHigh:pixelFormatType:options:error:"), cgImage, pixelsWide, pixelsHigh, pixelFormatType, options, error_)
	return rv
}/* debug [class_init_methods/constructor]: NewFeatureValueWithCGImagePixelsWidePixelsHighPixelFormatTypeOptionsError */


// Creates a feature value that contains a dictionary of numbers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(dictionary:)
func NewFeatureValueWithDictionaryError(value foundation.IDictionary, error_ objectivec.IObject) FeatureValue {
	rv := objc.Send[FeatureValue](objc.ID(getFeatureValueClass().class), objc.Sel("featureValueWithDictionary:error:"), value, error_)
	return rv
}/* debug [class_init_methods/constructor]: NewFeatureValueWithDictionaryError */


// Creates a feature value that contains a double.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(double:)
func NewFeatureValueWithDouble(value float64) FeatureValue {
	rv := objc.Send[FeatureValue](objc.ID(getFeatureValueClass().class), objc.Sel("featureValueWithDouble:"), value)
	return rv
}/* debug [class_init_methods/constructor]: NewFeatureValueWithDouble */


// Creates a feature value that contains an image defined by an image URL and a constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(imageAtURL:constraint:options:)
func NewFeatureValueWithImageAtURLConstraintOptionsError(url objc.IObject /* cross-framework: NSURL */, constraint IMLImageConstraint, options foundation.IDictionary, error_ objectivec.IObject) FeatureValue {
	rv := objc.Send[FeatureValue](objc.ID(getFeatureValueClass().class), objc.Sel("featureValueWithImageAtURL:constraint:options:error:"), url, constraint, options, error_)
	return rv
}/* debug [class_init_methods/constructor]: NewFeatureValueWithImageAtURLConstraintOptionsError */


// Creates a feature value that contains an image defined by an image URL, an orientation, and a constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(imageAtURL:orientation:constraint:options:)
func NewFeatureValueWithImageAtURLOrientationConstraintOptionsError(url objc.IObject /* cross-framework: NSURL */, orientation ImagePropertyOrientation /* not a class type */, constraint IMLImageConstraint, options foundation.IDictionary, error_ objectivec.IObject) FeatureValue {
	rv := objc.Send[FeatureValue](objc.ID(getFeatureValueClass().class), objc.Sel("featureValueWithImageAtURL:orientation:constraint:options:error:"), url, orientation, constraint, options, error_)
	return rv
}/* debug [class_init_methods/constructor]: NewFeatureValueWithImageAtURLOrientationConstraintOptionsError */


// Creates a feature value that contains an image defined by an image URL and the image’s orientation, size, and pixel format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(imageAtURL:orientation:pixelsWide:pixelsHigh:pixelFormatType:options:)
func NewFeatureValueWithImageAtURLOrientationPixelsWidePixelsHighPixelFormatTypeOptionsError(url objc.IObject /* cross-framework: NSURL */, orientation ImagePropertyOrientation /* not a class type */, pixelsWide int, pixelsHigh int, pixelFormatType uint32 /* not a class type */, options foundation.IDictionary, error_ objectivec.IObject) FeatureValue {
	rv := objc.Send[FeatureValue](objc.ID(getFeatureValueClass().class), objc.Sel("featureValueWithImageAtURL:orientation:pixelsWide:pixelsHigh:pixelFormatType:options:error:"), url, orientation, pixelsWide, pixelsHigh, pixelFormatType, options, error_)
	return rv
}/* debug [class_init_methods/constructor]: NewFeatureValueWithImageAtURLOrientationPixelsWidePixelsHighPixelFormatTypeOptionsError */


// Creates a feature value that contains an image defined by an image URL and the image’s size and pixel format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(imageAtURL:pixelsWide:pixelsHigh:pixelFormatType:options:)
func NewFeatureValueWithImageAtURLPixelsWidePixelsHighPixelFormatTypeOptionsError(url objc.IObject /* cross-framework: NSURL */, pixelsWide int, pixelsHigh int, pixelFormatType uint32 /* not a class type */, options foundation.IDictionary, error_ objectivec.IObject) FeatureValue {
	rv := objc.Send[FeatureValue](objc.ID(getFeatureValueClass().class), objc.Sel("featureValueWithImageAtURL:pixelsWide:pixelsHigh:pixelFormatType:options:error:"), url, pixelsWide, pixelsHigh, pixelFormatType, options, error_)
	return rv
}/* debug [class_init_methods/constructor]: NewFeatureValueWithImageAtURLPixelsWidePixelsHighPixelFormatTypeOptionsError */


// Creates a feature value that contains an integer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(int64:)
func NewFeatureValueWithInt64(value int64) FeatureValue {
	rv := objc.Send[FeatureValue](objc.ID(getFeatureValueClass().class), objc.Sel("featureValueWithInt64:"), value)
	return rv
}/* debug [class_init_methods/constructor]: NewFeatureValueWithInt64 */


// Creates a feature value that contains a multidimensional array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(multiArray:)
func NewFeatureValueWithMultiArray(value IMLMultiArray) FeatureValue {
	rv := objc.Send[FeatureValue](objc.ID(getFeatureValueClass().class), objc.Sel("featureValueWithMultiArray:"), value)
	return rv
}/* debug [class_init_methods/constructor]: NewFeatureValueWithMultiArray */


// Creates a feature value that contains an image from a pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(pixelBuffer:)
func NewFeatureValueWithPixelBuffer(value PixelBufferRef /* not a class type */) FeatureValue {
	rv := objc.Send[FeatureValue](objc.ID(getFeatureValueClass().class), objc.Sel("featureValueWithPixelBuffer:"), value)
	return rv
}/* debug [class_init_methods/constructor]: NewFeatureValueWithPixelBuffer */


// Creates a feature value that contains a sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(sequence:)
func NewFeatureValueWithSequence(sequence IMLSequence) FeatureValue {
	rv := objc.Send[FeatureValue](objc.ID(getFeatureValueClass().class), objc.Sel("featureValueWithSequence:"), sequence)
	return rv
}/* debug [class_init_methods/constructor]: NewFeatureValueWithSequence */


// Creates a feature value that contains a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(string:)
func NewFeatureValueWithString(value objc.IObject /* cross-framework: NSString */) FeatureValue {
	rv := objc.Send[FeatureValue](objc.ID(getFeatureValueClass().class), objc.Sel("featureValueWithString:"), value)
	return rv
}/* debug [class_init_methods/constructor]: NewFeatureValueWithString */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FeatureValue */

// Creates a feature value that contains an image defined by a core graphics image and a constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(CGImage:constraint:options:)
func (fc _FeatureValueClass) FeatureValueWithCGImageConstraintOptionsError(cgImage ImageRef /* not a class type */, constraint IMLImageConstraint, options foundation.IDictionary, error_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(fc.class), objc.Sel("featureValueWithCGImage:constraint:options:error:"), cgImage, constraint, options, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FeatureValueWithCGImageConstraintOptionsError) */


// Creates a feature value that contains an image defined by a core graphics image, an orientation, and a constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(CGImage:orientation:constraint:options:)
func (fc _FeatureValueClass) FeatureValueWithCGImageOrientationConstraintOptionsError(cgImage ImageRef /* not a class type */, orientation ImagePropertyOrientation /* not a class type */, constraint IMLImageConstraint, options foundation.IDictionary, error_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(fc.class), objc.Sel("featureValueWithCGImage:orientation:constraint:options:error:"), cgImage, orientation, constraint, options, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FeatureValueWithCGImageOrientationConstraintOptionsError) */


// Creates a feature value that contains an image defined by a core graphics image and its orientation, size, and pixel format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(CGImage:orientation:pixelsWide:pixelsHigh:pixelFormatType:options:)
func (fc _FeatureValueClass) FeatureValueWithCGImageOrientationPixelsWidePixelsHighPixelFormatTypeOptionsError(cgImage ImageRef /* not a class type */, orientation ImagePropertyOrientation /* not a class type */, pixelsWide int, pixelsHigh int, pixelFormatType uint32 /* not a class type */, options foundation.IDictionary, error_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(fc.class), objc.Sel("featureValueWithCGImage:orientation:pixelsWide:pixelsHigh:pixelFormatType:options:error:"), cgImage, orientation, pixelsWide, pixelsHigh, pixelFormatType, options, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FeatureValueWithCGImageOrientationPixelsWidePixelsHighPixelFormatTypeOptionsError) */


// Creates a feature value that contains an image defined by a core graphics image and its size and pixel format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(CGImage:pixelsWide:pixelsHigh:pixelFormatType:options:)
func (fc _FeatureValueClass) FeatureValueWithCGImagePixelsWidePixelsHighPixelFormatTypeOptionsError(cgImage ImageRef /* not a class type */, pixelsWide int, pixelsHigh int, pixelFormatType uint32 /* not a class type */, options foundation.IDictionary, error_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(fc.class), objc.Sel("featureValueWithCGImage:pixelsWide:pixelsHigh:pixelFormatType:options:error:"), cgImage, pixelsWide, pixelsHigh, pixelFormatType, options, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FeatureValueWithCGImagePixelsWidePixelsHighPixelFormatTypeOptionsError) */


// Creates a feature value that contains a dictionary of numbers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(dictionary:)
func (fc _FeatureValueClass) FeatureValueWithDictionaryError(value foundation.IDictionary, error_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(fc.class), objc.Sel("featureValueWithDictionary:error:"), value, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FeatureValueWithDictionaryError) */


// Creates a feature value that contains a double.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(double:)
func (fc _FeatureValueClass) FeatureValueWithDouble(value float64) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(fc.class), objc.Sel("featureValueWithDouble:"), value)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FeatureValueWithDouble) */


// Creates a feature value that contains an image defined by an image URL and a constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(imageAtURL:constraint:options:)
func (fc _FeatureValueClass) FeatureValueWithImageAtURLConstraintOptionsError(url objc.IObject /* cross-framework: NSURL */, constraint IMLImageConstraint, options foundation.IDictionary, error_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(fc.class), objc.Sel("featureValueWithImageAtURL:constraint:options:error:"), url, constraint, options, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FeatureValueWithImageAtURLConstraintOptionsError) */


// Creates a feature value that contains an image defined by an image URL, an orientation, and a constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(imageAtURL:orientation:constraint:options:)
func (fc _FeatureValueClass) FeatureValueWithImageAtURLOrientationConstraintOptionsError(url objc.IObject /* cross-framework: NSURL */, orientation ImagePropertyOrientation /* not a class type */, constraint IMLImageConstraint, options foundation.IDictionary, error_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(fc.class), objc.Sel("featureValueWithImageAtURL:orientation:constraint:options:error:"), url, orientation, constraint, options, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FeatureValueWithImageAtURLOrientationConstraintOptionsError) */


// Creates a feature value that contains an image defined by an image URL and the image’s orientation, size, and pixel format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(imageAtURL:orientation:pixelsWide:pixelsHigh:pixelFormatType:options:)
func (fc _FeatureValueClass) FeatureValueWithImageAtURLOrientationPixelsWidePixelsHighPixelFormatTypeOptionsError(url objc.IObject /* cross-framework: NSURL */, orientation ImagePropertyOrientation /* not a class type */, pixelsWide int, pixelsHigh int, pixelFormatType uint32 /* not a class type */, options foundation.IDictionary, error_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(fc.class), objc.Sel("featureValueWithImageAtURL:orientation:pixelsWide:pixelsHigh:pixelFormatType:options:error:"), url, orientation, pixelsWide, pixelsHigh, pixelFormatType, options, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FeatureValueWithImageAtURLOrientationPixelsWidePixelsHighPixelFormatTypeOptionsError) */


// Creates a feature value that contains an image defined by an image URL and the image’s size and pixel format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(imageAtURL:pixelsWide:pixelsHigh:pixelFormatType:options:)
func (fc _FeatureValueClass) FeatureValueWithImageAtURLPixelsWidePixelsHighPixelFormatTypeOptionsError(url objc.IObject /* cross-framework: NSURL */, pixelsWide int, pixelsHigh int, pixelFormatType uint32 /* not a class type */, options foundation.IDictionary, error_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(fc.class), objc.Sel("featureValueWithImageAtURL:pixelsWide:pixelsHigh:pixelFormatType:options:error:"), url, pixelsWide, pixelsHigh, pixelFormatType, options, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FeatureValueWithImageAtURLPixelsWidePixelsHighPixelFormatTypeOptionsError) */


// Creates a feature value that contains an integer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(int64:)
func (fc _FeatureValueClass) FeatureValueWithInt64(value int64) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(fc.class), objc.Sel("featureValueWithInt64:"), value)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FeatureValueWithInt64) */


// Creates a feature value that contains a multidimensional array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(multiArray:)
func (fc _FeatureValueClass) FeatureValueWithMultiArray(value IMLMultiArray) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(fc.class), objc.Sel("featureValueWithMultiArray:"), value)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FeatureValueWithMultiArray) */


// Creates a feature value that contains an image from a pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(pixelBuffer:)
func (fc _FeatureValueClass) FeatureValueWithPixelBuffer(value PixelBufferRef /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(fc.class), objc.Sel("featureValueWithPixelBuffer:"), value)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FeatureValueWithPixelBuffer) */


// Creates a feature value that contains a sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(sequence:)
func (fc _FeatureValueClass) FeatureValueWithSequence(sequence IMLSequence) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(fc.class), objc.Sel("featureValueWithSequence:"), sequence)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FeatureValueWithSequence) */


// Creates a feature value that contains a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(string:)
func (fc _FeatureValueClass) FeatureValueWithString(value objc.IObject /* cross-framework: NSString */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(fc.class), objc.Sel("featureValueWithString:"), value)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FeatureValueWithString) */


// Creates a feature value with a type that represents an undefined or missing value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(undefined:)
func (fc _FeatureValueClass) UndefinedFeatureValueWithType(type_ FeatureType) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(fc.class), objc.Sel("undefinedFeatureValueWithType:"), type_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=UndefinedFeatureValueWithType) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FeatureValue */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FeatureValue */

// Returns a Boolean value that indicates whether a feature value is equal to another.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/isEqual(to:)
func (f_ FeatureValue) IsEqualToFeatureValue(value IMLFeatureValue) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isEqualToFeatureValue:"), value)
	return rv
}/* debug [instance_methods/method]: IsEqualToFeatureValue */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FeatureValue */

// The underlying dictionary of the feature value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/dictionaryValue
func (f_ FeatureValue) DictionaryValue() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](f_.ID, objc.Sel("dictionaryValue"))
	return rv
}/* debug [instance_properties/getter]: dictionaryValue */


// The underlying double of the feature value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/doubleValue
func (f_ FeatureValue) DoubleValue() float64 {
	rv := objc.Send[float64](f_.ID, objc.Sel("doubleValue"))
	return rv
}/* debug [instance_properties/getter]: doubleValue */


// The underlying image of the feature value as a pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/imageBufferValue
func (f_ FeatureValue) ImageBufferValue() PixelBufferRef /* not a class type */ {
	rv := objc.Send[PixelBufferRef](f_.ID, objc.Sel("imageBufferValue"))
	return rv
}/* debug [instance_properties/getter]: imageBufferValue */


// The underlying integer of the feature value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/int64Value
func (f_ FeatureValue) Int64Value() int64 {
	rv := objc.Send[int64](f_.ID, objc.Sel("int64Value"))
	return rv
}/* debug [instance_properties/getter]: int64Value */


// A Boolean value that indicates whether the feature value is undefined or missing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/isUndefined
func (f_ FeatureValue) Undefined() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("undefined"))
	return rv
}/* debug [instance_properties/getter]: undefined */


// The underlying multiarray of the feature value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/multiArrayValue
func (f_ FeatureValue) MultiArrayValue() IMLMultiArray {
	rv := objc.Send[MultiArray](f_.ID, objc.Sel("multiArrayValue"))
	return rv
}/* debug [instance_properties/getter]: multiArrayValue */


// The underlying sequence of the feature value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/sequenceValue
func (f_ FeatureValue) SequenceValue() IMLSequence {
	rv := objc.Send[Sequence](f_.ID, objc.Sel("sequenceValue"))
	return rv
}/* debug [instance_properties/getter]: sequenceValue */


// The underlying string of the feature value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/stringValue
func (f_ FeatureValue) StringValue() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("stringValue"))
	return rv
}/* debug [instance_properties/getter]: stringValue */


// The type of the feature value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/type
func (f_ FeatureValue) Type() FeatureType {
	rv := objc.Send[FeatureType](f_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// A Boolean value that indicates whether the feature value is undefined or missing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturevalue/isundefined
func (f_ FeatureValue) IsUndefined() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isUndefined"))
	return rv
}/* debug [instance_properties/getter]: isUndefined */


// A Boolean value that indicates whether the feature value is undefined or missing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturevalue/isundefined
func (f_ FeatureValue) SetIsUndefined(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsUndefined:"), value)
}/* debug [instance_properties/setter]: isUndefined */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLFeatureValue */


