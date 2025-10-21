// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [FeatureValue] class.
type IFeatureValue interface {
	objectivec.IObject
	IsEqualToFeatureValue(value unsafe.Pointer) bool
}

// A generic wrapper around an underlying value and the value’s type.
//
// A Core ML wraps an underlying value and bundles it with that value’s type, which is one of the types that defines. Apps typically access feature values indirectly by using the methods in the wrapper class Xcode automatically generates for Core ML model files. If your app accesses an directly, it must create and consume instances. For each prediction, Core ML accepts a feature provider for its inputs, and generates a separate feature provider for its outputs. The input feature provider contains one instance per input, and the output feature provider contains one per output. See for more information about the model input and output features.
//
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

// Alloc allocates a new instance without initialization.
func (fc _FeatureValueClass) Alloc() FeatureValue {
	rv := objc.Send[FeatureValue](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates a feature value with a type that represents an undefined or missing value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(undefined:)
func NewFeatureValueUndefinedFeatureValueWithType(type_ unsafe.Pointer) FeatureValue {
	rv := objc.Send[FeatureValue](objc.ID(getFeatureValueClass().class), objc.Sel("undefinedFeatureValueWithType:"), type_)
	return rv
}



// Creates a feature value that contains an image defined by a core graphics image and a constraint.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(CGImage:constraint:options:)
func NewFeatureValueWithCGImageConstraintOptionsError(cgImage coregraphics.CGImageRef, constraint unsafe.Pointer, options unsafe.Pointer, error_ unsafe.Pointer) FeatureValue {
	rv := objc.Send[FeatureValue](objc.ID(getFeatureValueClass().class), objc.Sel("featureValueWithCGImage:constraint:options:error:"), cgImage, constraint, options, error_)
	return rv
}



// Creates a feature value that contains an image defined by a core graphics image, an orientation, and a constraint.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(CGImage:orientation:constraint:options:)
func NewFeatureValueWithCGImageOrientationConstraintOptionsError(cgImage coregraphics.CGImageRef, orientation unsafe.Pointer, constraint unsafe.Pointer, options unsafe.Pointer, error_ unsafe.Pointer) FeatureValue {
	rv := objc.Send[FeatureValue](objc.ID(getFeatureValueClass().class), objc.Sel("featureValueWithCGImage:orientation:constraint:options:error:"), cgImage, orientation, constraint, options, error_)
	return rv
}



// Creates a feature value that contains an image defined by a core graphics image and its orientation, size, and pixel format.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(CGImage:orientation:pixelsWide:pixelsHigh:pixelFormatType:options:)
func NewFeatureValueWithCGImageOrientationPixelsWidePixelsHighPixelFormatTypeOptionsError(cgImage coregraphics.CGImageRef, orientation unsafe.Pointer, pixelsWide int, pixelsHigh int, pixelFormatType unsafe.Pointer, options unsafe.Pointer, error_ unsafe.Pointer) FeatureValue {
	rv := objc.Send[FeatureValue](objc.ID(getFeatureValueClass().class), objc.Sel("featureValueWithCGImage:orientation:pixelsWide:pixelsHigh:pixelFormatType:options:error:"), cgImage, orientation, pixelsWide, pixelsHigh, pixelFormatType, options, error_)
	return rv
}



// Creates a feature value that contains an image defined by a core graphics image and its size and pixel format.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(CGImage:pixelsWide:pixelsHigh:pixelFormatType:options:)
func NewFeatureValueWithCGImagePixelsWidePixelsHighPixelFormatTypeOptionsError(cgImage coregraphics.CGImageRef, pixelsWide int, pixelsHigh int, pixelFormatType unsafe.Pointer, options unsafe.Pointer, error_ unsafe.Pointer) FeatureValue {
	rv := objc.Send[FeatureValue](objc.ID(getFeatureValueClass().class), objc.Sel("featureValueWithCGImage:pixelsWide:pixelsHigh:pixelFormatType:options:error:"), cgImage, pixelsWide, pixelsHigh, pixelFormatType, options, error_)
	return rv
}



// Creates a feature value that contains a dictionary of numbers.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(dictionary:)
func NewFeatureValueWithDictionaryError(value unsafe.Pointer, error_ unsafe.Pointer) FeatureValue {
	rv := objc.Send[FeatureValue](objc.ID(getFeatureValueClass().class), objc.Sel("featureValueWithDictionary:error:"), value, error_)
	return rv
}



// Creates a feature value that contains a double.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(double:)
func NewFeatureValueWithDouble(value unsafe.Pointer) FeatureValue {
	rv := objc.Send[FeatureValue](objc.ID(getFeatureValueClass().class), objc.Sel("featureValueWithDouble:"), value)
	return rv
}



// Creates a feature value that contains an image defined by an image URL and a constraint.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(imageAtURL:constraint:options:)
func NewFeatureValueWithImageAtURLConstraintOptionsError(url foundation.URL, constraint unsafe.Pointer, options unsafe.Pointer, error_ unsafe.Pointer) FeatureValue {
	rv := objc.Send[FeatureValue](objc.ID(getFeatureValueClass().class), objc.Sel("featureValueWithImageAtURL:constraint:options:error:"), url, constraint, options, error_)
	return rv
}



// Creates a feature value that contains an image defined by an image URL, an orientation, and a constraint.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(imageAtURL:orientation:constraint:options:)
func NewFeatureValueWithImageAtURLOrientationConstraintOptionsError(url foundation.URL, orientation unsafe.Pointer, constraint unsafe.Pointer, options unsafe.Pointer, error_ unsafe.Pointer) FeatureValue {
	rv := objc.Send[FeatureValue](objc.ID(getFeatureValueClass().class), objc.Sel("featureValueWithImageAtURL:orientation:constraint:options:error:"), url, orientation, constraint, options, error_)
	return rv
}



// Creates a feature value that contains an image defined by an image URL and the image’s orientation, size, and pixel format.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(imageAtURL:orientation:pixelsWide:pixelsHigh:pixelFormatType:options:)
func NewFeatureValueWithImageAtURLOrientationPixelsWidePixelsHighPixelFormatTypeOptionsError(url foundation.URL, orientation unsafe.Pointer, pixelsWide int, pixelsHigh int, pixelFormatType unsafe.Pointer, options unsafe.Pointer, error_ unsafe.Pointer) FeatureValue {
	rv := objc.Send[FeatureValue](objc.ID(getFeatureValueClass().class), objc.Sel("featureValueWithImageAtURL:orientation:pixelsWide:pixelsHigh:pixelFormatType:options:error:"), url, orientation, pixelsWide, pixelsHigh, pixelFormatType, options, error_)
	return rv
}



// Creates a feature value that contains an image defined by an image URL and the image’s size and pixel format.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(imageAtURL:pixelsWide:pixelsHigh:pixelFormatType:options:)
func NewFeatureValueWithImageAtURLPixelsWidePixelsHighPixelFormatTypeOptionsError(url foundation.URL, pixelsWide int, pixelsHigh int, pixelFormatType unsafe.Pointer, options unsafe.Pointer, error_ unsafe.Pointer) FeatureValue {
	rv := objc.Send[FeatureValue](objc.ID(getFeatureValueClass().class), objc.Sel("featureValueWithImageAtURL:pixelsWide:pixelsHigh:pixelFormatType:options:error:"), url, pixelsWide, pixelsHigh, pixelFormatType, options, error_)
	return rv
}



// Creates a feature value that contains an integer.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(int64:)
func NewFeatureValueWithInt64(value unsafe.Pointer) FeatureValue {
	rv := objc.Send[FeatureValue](objc.ID(getFeatureValueClass().class), objc.Sel("featureValueWithInt64:"), value)
	return rv
}



// Creates a feature value that contains a multidimensional array.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(multiArray:)
func NewFeatureValueWithMultiArray(value unsafe.Pointer) FeatureValue {
	rv := objc.Send[FeatureValue](objc.ID(getFeatureValueClass().class), objc.Sel("featureValueWithMultiArray:"), value)
	return rv
}



// Creates a feature value that contains an image from a pixel buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(pixelBuffer:)
func NewFeatureValueWithPixelBuffer(value unsafe.Pointer) FeatureValue {
	rv := objc.Send[FeatureValue](objc.ID(getFeatureValueClass().class), objc.Sel("featureValueWithPixelBuffer:"), value)
	return rv
}



// Creates a feature value that contains a sequence.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(sequence:)
func NewFeatureValueWithSequence(sequence unsafe.Pointer) FeatureValue {
	rv := objc.Send[FeatureValue](objc.ID(getFeatureValueClass().class), objc.Sel("featureValueWithSequence:"), sequence)
	return rv
}



// Creates a feature value that contains a string.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(string:)
func NewFeatureValueWithString(value string) FeatureValue {
	rv := objc.Send[FeatureValue](objc.ID(getFeatureValueClass().class), objc.Sel("featureValueWithString:"), objc.String(value))
	return rv
}


// Creates a feature value that contains an image defined by a core graphics image and a constraint.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(CGImage:constraint:options:)
func (fc _FeatureValueClass) FeatureValueWithCGImageConstraintOptionsError(cgImage coregraphics.CGImageRef, constraint unsafe.Pointer, options unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("featureValueWithCGImage:constraint:options:error:"), cgImage, constraint, options, error_)
	return rv
}

// Creates a feature value that contains an image defined by a core graphics image, an orientation, and a constraint.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(CGImage:orientation:constraint:options:)
func (fc _FeatureValueClass) FeatureValueWithCGImageOrientationConstraintOptionsError(cgImage coregraphics.CGImageRef, orientation unsafe.Pointer, constraint unsafe.Pointer, options unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("featureValueWithCGImage:orientation:constraint:options:error:"), cgImage, orientation, constraint, options, error_)
	return rv
}

// Creates a feature value that contains an image defined by a core graphics image and its orientation, size, and pixel format.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(CGImage:orientation:pixelsWide:pixelsHigh:pixelFormatType:options:)
func (fc _FeatureValueClass) FeatureValueWithCGImageOrientationPixelsWidePixelsHighPixelFormatTypeOptionsError(cgImage coregraphics.CGImageRef, orientation unsafe.Pointer, pixelsWide int, pixelsHigh int, pixelFormatType unsafe.Pointer, options unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("featureValueWithCGImage:orientation:pixelsWide:pixelsHigh:pixelFormatType:options:error:"), cgImage, orientation, pixelsWide, pixelsHigh, pixelFormatType, options, error_)
	return rv
}

// Creates a feature value that contains an image defined by a core graphics image and its size and pixel format.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(CGImage:pixelsWide:pixelsHigh:pixelFormatType:options:)
func (fc _FeatureValueClass) FeatureValueWithCGImagePixelsWidePixelsHighPixelFormatTypeOptionsError(cgImage coregraphics.CGImageRef, pixelsWide int, pixelsHigh int, pixelFormatType unsafe.Pointer, options unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("featureValueWithCGImage:pixelsWide:pixelsHigh:pixelFormatType:options:error:"), cgImage, pixelsWide, pixelsHigh, pixelFormatType, options, error_)
	return rv
}

// Creates a feature value that contains a dictionary of numbers.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(dictionary:)
func (fc _FeatureValueClass) FeatureValueWithDictionaryError(value unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("featureValueWithDictionary:error:"), value, error_)
	return rv
}

// Creates a feature value that contains a double.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(double:)
func (fc _FeatureValueClass) FeatureValueWithDouble(value unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("featureValueWithDouble:"), value)
	return rv
}

// Creates a feature value that contains an image defined by an image URL and a constraint.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(imageAtURL:constraint:options:)
func (fc _FeatureValueClass) FeatureValueWithImageAtURLConstraintOptionsError(url foundation.URL, constraint unsafe.Pointer, options unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("featureValueWithImageAtURL:constraint:options:error:"), url, constraint, options, error_)
	return rv
}

// Creates a feature value that contains an image defined by an image URL, an orientation, and a constraint.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(imageAtURL:orientation:constraint:options:)
func (fc _FeatureValueClass) FeatureValueWithImageAtURLOrientationConstraintOptionsError(url foundation.URL, orientation unsafe.Pointer, constraint unsafe.Pointer, options unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("featureValueWithImageAtURL:orientation:constraint:options:error:"), url, orientation, constraint, options, error_)
	return rv
}

// Creates a feature value that contains an image defined by an image URL and the image’s orientation, size, and pixel format.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(imageAtURL:orientation:pixelsWide:pixelsHigh:pixelFormatType:options:)
func (fc _FeatureValueClass) FeatureValueWithImageAtURLOrientationPixelsWidePixelsHighPixelFormatTypeOptionsError(url foundation.URL, orientation unsafe.Pointer, pixelsWide int, pixelsHigh int, pixelFormatType unsafe.Pointer, options unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("featureValueWithImageAtURL:orientation:pixelsWide:pixelsHigh:pixelFormatType:options:error:"), url, orientation, pixelsWide, pixelsHigh, pixelFormatType, options, error_)
	return rv
}

// Creates a feature value that contains an image defined by an image URL and the image’s size and pixel format.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(imageAtURL:pixelsWide:pixelsHigh:pixelFormatType:options:)
func (fc _FeatureValueClass) FeatureValueWithImageAtURLPixelsWidePixelsHighPixelFormatTypeOptionsError(url foundation.URL, pixelsWide int, pixelsHigh int, pixelFormatType unsafe.Pointer, options unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("featureValueWithImageAtURL:pixelsWide:pixelsHigh:pixelFormatType:options:error:"), url, pixelsWide, pixelsHigh, pixelFormatType, options, error_)
	return rv
}

// Creates a feature value that contains an integer.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(int64:)
func (fc _FeatureValueClass) FeatureValueWithInt64(value unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("featureValueWithInt64:"), value)
	return rv
}

// Creates a feature value that contains a multidimensional array.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(multiArray:)
func (fc _FeatureValueClass) FeatureValueWithMultiArray(value unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("featureValueWithMultiArray:"), value)
	return rv
}

// Creates a feature value that contains an image from a pixel buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(pixelBuffer:)
func (fc _FeatureValueClass) FeatureValueWithPixelBuffer(value unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("featureValueWithPixelBuffer:"), value)
	return rv
}

// Creates a feature value that contains a sequence.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(sequence:)
func (fc _FeatureValueClass) FeatureValueWithSequence(sequence unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("featureValueWithSequence:"), sequence)
	return rv
}

// Creates a feature value that contains a string.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(string:)
func (fc _FeatureValueClass) FeatureValueWithString(value string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("featureValueWithString:"), objc.String(value))
	return rv
}

// Creates a feature value with a type that represents an undefined or missing value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(undefined:)
func (fc _FeatureValueClass) UndefinedFeatureValueWithType(type_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("undefinedFeatureValueWithType:"), type_)
	return rv
}

// Returns a Boolean value that indicates whether a feature value is equal to another.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/isEqual(to:)
func (f_ FeatureValue) IsEqualToFeatureValue(value unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isEqualToFeatureValue:"), value)
	return rv
}

// A Boolean value that indicates whether the feature value is undefined or missing.
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturevalue/isundefined
func (f_ FeatureValue) IsUndefined() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isUndefined"))
	return rv
}


// SetIsUndefined sets the value of the isUndefined property.
// A Boolean value that indicates whether the feature value is undefined or missing.

//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturevalue/isundefined
func (f_ FeatureValue) SetIsUndefined(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsUndefined:"), value)
}

// The underlying dictionary of the feature value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/dictionaryValue
func (f_ FeatureValue) DictionaryValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("dictionaryValue"))
	return rv
}

// The underlying double of the feature value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/doubleValue
func (f_ FeatureValue) DoubleValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("doubleValue"))
	return rv
}

// The underlying image of the feature value as a pixel buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/imageBufferValue
func (f_ FeatureValue) ImageBufferValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("imageBufferValue"))
	return rv
}

// The underlying integer of the feature value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/int64Value
func (f_ FeatureValue) Int64Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("int64Value"))
	return rv
}

// A Boolean value that indicates whether the feature value is undefined or missing.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/isUndefined
func (f_ FeatureValue) Undefined() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("undefined"))
	return rv
}

// The underlying multiarray of the feature value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/multiArrayValue
func (f_ FeatureValue) MultiArrayValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("multiArrayValue"))
	return rv
}

// The underlying sequence of the feature value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/sequenceValue
func (f_ FeatureValue) SequenceValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("sequenceValue"))
	return rv
}

// The underlying string of the feature value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/stringValue
func (f_ FeatureValue) StringValue() string {
	rv := objc.Send[string](f_.ID, objc.Sel("stringValue"))
	return rv
}

// The type of the feature value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureValue/type
func (f_ FeatureValue) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("type"))
	return rv
}


