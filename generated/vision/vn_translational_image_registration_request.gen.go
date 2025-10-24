// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VNTranslationalImageRegistrationRequest */


/* debug [class_header]: Header for VNTranslationalImageRegistrationRequest */
// The class instance for the [TranslationalImageRegistrationRequest] class.
var (
	TranslationalImageRegistrationRequestClass     _TranslationalImageRegistrationRequestClass
	TranslationalImageRegistrationRequestClassOnce sync.Once
)

func getTranslationalImageRegistrationRequestClass() _TranslationalImageRegistrationRequestClass {
	TranslationalImageRegistrationRequestClassOnce.Do(func() {
		TranslationalImageRegistrationRequestClass = _TranslationalImageRegistrationRequestClass{objc.GetClass("VNTranslationalImageRegistrationRequest")}
	})
	return TranslationalImageRegistrationRequestClass
}

type _TranslationalImageRegistrationRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TranslationalImageRegistrationRequest */
// An interface definition for the [TranslationalImageRegistrationRequest] class.
type ITranslationalImageRegistrationRequest interface {
	IImageRegistrationRequest
	
/* debug [class_interface_properties]: Properties for TranslationalImageRegistrationRequest */
	// properties:
	Results() []ImageTranslationAlignmentObservation
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TranslationalImageRegistrationRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TranslationalImageRegistrationRequest */
// Alloc allocates a new instance without initialization.
func (tc _TranslationalImageRegistrationRequestClass) Alloc() TranslationalImageRegistrationRequest {
	rv := objc.Send[TranslationalImageRegistrationRequest](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TranslationalImageRegistrationRequestClass) New() TranslationalImageRegistrationRequest {
	rv := objc.Send[TranslationalImageRegistrationRequest](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TranslationalImageRegistrationRequest) Init() TranslationalImageRegistrationRequest {
	rv := objc.Send[TranslationalImageRegistrationRequest](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TranslationalImageRegistrationRequest) Autorelease() TranslationalImageRegistrationRequest {
	rv := objc.Send[TranslationalImageRegistrationRequest](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTranslationalImageRegistrationRequest creates a new TranslationalImageRegistrationRequest instance.
func NewTranslationalImageRegistrationRequest() TranslationalImageRegistrationRequest {
	return getTranslationalImageRegistrationRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TranslationalImageRegistrationRequest */
// An image-analysis request that determines the affine transform necessary to align the content of two images.
//
// Create and perform a translational image registration request to align content in two images through translation.


// An image-analysis request that determines the affine transform necessary to align the content of two images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTranslationalImageRegistrationRequest
type TranslationalImageRegistrationRequest struct {
	ImageRegistrationRequest
}

// TranslationalImageRegistrationRequestFrom constructs a [TranslationalImageRegistrationRequest] from an unsafe.Pointer.
//
// An image-analysis request that determines the affine transform necessary to align the content of two images.
func TranslationalImageRegistrationRequestFrom(ptr unsafe.Pointer) TranslationalImageRegistrationRequest {
	return TranslationalImageRegistrationRequest{
		ImageRegistrationRequest: ImageRegistrationRequestFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TranslationalImageRegistrationRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TranslationalImageRegistrationRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TranslationalImageRegistrationRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TranslationalImageRegistrationRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TranslationalImageRegistrationRequest */

// The results of a translational image alignment request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTranslationalImageRegistrationRequest/results
func (t_ TranslationalImageRegistrationRequest) Results() []ImageTranslationAlignmentObservation {
	rv := objc.Send[[]ImageTranslationAlignmentObservation](t_.ID, objc.Sel("results"))
	return rv
}/* debug [instance_properties/getter]: results */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNTranslationalImageRegistrationRequest */



