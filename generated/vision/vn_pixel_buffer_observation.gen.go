// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PixelBufferObservation] class.
var (
	PixelBufferObservationClass     _PixelBufferObservationClass
	PixelBufferObservationClassOnce sync.Once
)

func getPixelBufferObservationClass() _PixelBufferObservationClass {
	PixelBufferObservationClassOnce.Do(func() {
		PixelBufferObservationClass = _PixelBufferObservationClass{objc.GetClass("VNPixelBufferObservation")}
	})
	return PixelBufferObservationClass
}

type _PixelBufferObservationClass struct {
	class objc.Class
}

// An interface definition for the [PixelBufferObservation] class.
type IPixelBufferObservation interface {
	IObservation
}

// An object that represents an image that an image-analysis request produces.
//
// This type of observation results from performing a image analysis with a Core ML model that has an image-to-image processing role. For example, this observation might result from a model that analyzes the style of one image and then transfers that style to a different image. Vision infers that an object is an image-to-image model if that model includes an image. Its object includes an image-typed feature description in its dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNPixelBufferObservation
type PixelBufferObservation struct {
	Observation
}

// PixelBufferObservationFrom constructs a [PixelBufferObservation] from an unsafe.Pointer.
//
// An object that represents an image that an image-analysis request produces.
func PixelBufferObservationFrom(ptr unsafe.Pointer) PixelBufferObservation {
	return PixelBufferObservation{
		Observation: ObservationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PixelBufferObservationClass) Alloc() PixelBufferObservation {
	rv := objc.Send[PixelBufferObservation](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PixelBufferObservationClass) New() PixelBufferObservation {
	rv := objc.Send[PixelBufferObservation](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PixelBufferObservation) Init() PixelBufferObservation {
	rv := objc.Send[PixelBufferObservation](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PixelBufferObservation) Autorelease() PixelBufferObservation {
	rv := objc.Send[PixelBufferObservation](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPixelBufferObservation creates a new PixelBufferObservation instance.
func NewPixelBufferObservation() PixelBufferObservation {
	return getPixelBufferObservationClass().New()
}


// The image that results from a request with image output.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnpixelbufferobservation/pixelbuffer
func (p_ PixelBufferObservation) PixelBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("pixelBuffer"))
	return rv
}


// SetPixelBuffer sets the value of the pixelBuffer property.
// The image that results from a request with image output.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnpixelbufferobservation/pixelbuffer
func (p_ PixelBufferObservation) SetPixelBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPixelBuffer:"), value)
}

// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/modelDescription
func (p_ PixelBufferObservation) ModelDescription() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("modelDescription"))
	return rv
}


// SetModelDescription sets the value of the modelDescription property.
// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/modelDescription
func (p_ PixelBufferObservation) SetModelDescription(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setModelDescription:"), value)
}

// A feature name that the CoreML model defines.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnpixelbufferobservation/featurename
func (p_ PixelBufferObservation) FeatureName() string {
	rv := objc.Send[string](p_.ID, objc.Sel("featureName"))
	return rv
}


// SetFeatureName sets the value of the featureName property.
// A feature name that the CoreML model defines.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnpixelbufferobservation/featurename
func (p_ PixelBufferObservation) SetFeatureName(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFeatureName:"), objc.String(value))
}

// A dictionary of output feature descriptions, which the model keys by the output’s name.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/outputDescriptionsByName
func (p_ PixelBufferObservation) OutputDescriptionsByName() string {
	rv := objc.Send[string](p_.ID, objc.Sel("outputDescriptionsByName"))
	return rv
}


// SetOutputDescriptionsByName sets the value of the outputDescriptionsByName property.
// A dictionary of output feature descriptions, which the model keys by the output’s name.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/outputDescriptionsByName
func (p_ PixelBufferObservation) SetOutputDescriptionsByName(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOutputDescriptionsByName:"), objc.String(value))
}



