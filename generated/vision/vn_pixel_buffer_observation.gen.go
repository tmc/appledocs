// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coreml"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class VNPixelBufferObservation */


/* debug [class_header]: Header for VNPixelBufferObservation */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PixelBufferObservation */
// An interface definition for the [PixelBufferObservation] class.
type IPixelBufferObservation interface {
	IObservation
	
/* debug [class_interface_properties]: Properties for PixelBufferObservation */
	// properties:
	FeatureName() objc.IObject /* cross-framework: NSString */
	PixelBuffer() PixelBufferRef /* not a class type */
	ModelDescription() coreml.ModelDescription
	SetModelDescription(value coreml.ModelDescription)
	OutputDescriptionsByName() coreml.FeatureDescription
	SetOutputDescriptionsByName(value coreml.FeatureDescription)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PixelBufferObservation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PixelBufferObservation */
// Alloc allocates a new instance without initialization.
func (pc _PixelBufferObservationClass) Alloc() PixelBufferObservation {
	rv := objc.Send[PixelBufferObservation](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PixelBufferObservation */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PixelBufferObservation *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PixelBufferObservation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PixelBufferObservation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PixelBufferObservation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PixelBufferObservation */

// A feature name that the CoreML model defines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNPixelBufferObservation/featureName
func (p_ PixelBufferObservation) FeatureName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("featureName"))
	return rv
}/* debug [instance_properties/getter]: featureName */


// The image that results from a request with image output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNPixelBufferObservation/pixelBuffer
func (p_ PixelBufferObservation) PixelBuffer() PixelBufferRef /* not a class type */ {
	rv := objc.Send[PixelBufferRef](p_.ID, objc.Sel("pixelBuffer"))
	return rv
}/* debug [instance_properties/getter]: pixelBuffer */


// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/modelDescription
func (p_ PixelBufferObservation) ModelDescription() coreml.ModelDescription {
	rv := objc.Send[coreml.ModelDescription](p_.ID, objc.Sel("modelDescription"))
	return rv
}/* debug [instance_properties/getter]: modelDescription */


// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/modelDescription
func (p_ PixelBufferObservation) SetModelDescription(value coreml.ModelDescription) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setModelDescription:"), value)
}/* debug [instance_properties/setter]: modelDescription */


// A dictionary of output feature descriptions, which the model keys by the output’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/outputDescriptionsByName
func (p_ PixelBufferObservation) OutputDescriptionsByName() coreml.FeatureDescription {
	rv := objc.Send[coreml.FeatureDescription](p_.ID, objc.Sel("outputDescriptionsByName"))
	return rv
}/* debug [instance_properties/getter]: outputDescriptionsByName */


// A dictionary of output feature descriptions, which the model keys by the output’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/outputDescriptionsByName
func (p_ PixelBufferObservation) SetOutputDescriptionsByName(value coreml.FeatureDescription) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOutputDescriptionsByName:"), value)
}/* debug [instance_properties/setter]: outputDescriptionsByName */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNPixelBufferObservation */



