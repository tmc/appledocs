// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coreml"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	ModelDescription() objc.IObject /* cross-framework: ModelDescription */
	SetModelDescription(value objc.IObject /* cross-framework: ModelDescription */)
	OutputDescriptionsByName() objc.IObject /* cross-framework: FeatureDescription */
	SetOutputDescriptionsByName(value objc.IObject /* cross-framework: FeatureDescription */)
	FeatureName() objc.IObject /* cross-framework: NSString */
	SetFeatureName(value objc.IObject /* cross-framework: NSString */)
	PixelBuffer() PixelBuffer /* not a class type */
	SetPixelBuffer(value PixelBuffer /* not a class type */)
	// methods:
}

// An object that represents an image that an image-analysis request produces.
//
// This type of observation results from performing a image analysis with a Core ML model that has an image-to-image processing role. For example, this observation might result from a model that analyzes the style of one image and then transfers that style to a different image. Vision infers that an object is an image-to-image model if that model includes an image. Its object includes an image-typed feature description in its dictionary.


// An object that represents an image that an image-analysis request produces.
//
// [Full Topic]
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



// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/modelDescription
func (p_ PixelBufferObservation) ModelDescription() objc.IObject /* cross-framework: ModelDescription */ {
	rv := objc.Send[coreml.ModelDescription](p_.ID, objc.Sel("modelDescription"))
	return rv
}


// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/modelDescription
func (p_ PixelBufferObservation) SetModelDescription(value objc.IObject /* cross-framework: ModelDescription */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setModelDescription:"), value)
}


// A dictionary of output feature descriptions, which the model keys by the output’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/outputDescriptionsByName
func (p_ PixelBufferObservation) OutputDescriptionsByName() objc.IObject /* cross-framework: FeatureDescription */ {
	rv := objc.Send[coreml.FeatureDescription](p_.ID, objc.Sel("outputDescriptionsByName"))
	return rv
}


// A dictionary of output feature descriptions, which the model keys by the output’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/outputDescriptionsByName
func (p_ PixelBufferObservation) SetOutputDescriptionsByName(value objc.IObject /* cross-framework: FeatureDescription */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOutputDescriptionsByName:"), value)
}


// A feature name that the CoreML model defines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnpixelbufferobservation/featurename
func (p_ PixelBufferObservation) FeatureName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("featureName"))
	return rv
}


// A feature name that the CoreML model defines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnpixelbufferobservation/featurename
func (p_ PixelBufferObservation) SetFeatureName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFeatureName:"), value)
}


// The image that results from a request with image output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnpixelbufferobservation/pixelbuffer
func (p_ PixelBufferObservation) PixelBuffer() PixelBuffer /* not a class type */ {
	rv := objc.Send[PixelBuffer](p_.ID, objc.Sel("pixelBuffer"))
	return rv
}


// The image that results from a request with image output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnpixelbufferobservation/pixelbuffer
func (p_ PixelBufferObservation) SetPixelBuffer(value PixelBuffer /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPixelBuffer:"), value)
}



