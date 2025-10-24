// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VNRecognizedTextObservation */


/* debug [class_header]: Header for VNRecognizedTextObservation */
// The class instance for the [RecognizedTextObservation] class.
var (
	RecognizedTextObservationClass     _RecognizedTextObservationClass
	RecognizedTextObservationClassOnce sync.Once
)

func getRecognizedTextObservationClass() _RecognizedTextObservationClass {
	RecognizedTextObservationClassOnce.Do(func() {
		RecognizedTextObservationClass = _RecognizedTextObservationClass{objc.GetClass("VNRecognizedTextObservation")}
	})
	return RecognizedTextObservationClass
}

type _RecognizedTextObservationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RecognizedTextObservation */
// An interface definition for the [RecognizedTextObservation] class.
type IRecognizedTextObservation interface {
	IRectangleObservation
	
/* debug [class_interface_properties]: Properties for RecognizedTextObservation */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RecognizedTextObservation */
	// methods:
	TopCandidates(maxCandidateCount uint) []RecognizedText
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RecognizedTextObservation */
// Alloc allocates a new instance without initialization.
func (rc _RecognizedTextObservationClass) Alloc() RecognizedTextObservation {
	rv := objc.Send[RecognizedTextObservation](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RecognizedTextObservationClass) New() RecognizedTextObservation {
	rv := objc.Send[RecognizedTextObservation](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RecognizedTextObservation) Init() RecognizedTextObservation {
	rv := objc.Send[RecognizedTextObservation](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RecognizedTextObservation) Autorelease() RecognizedTextObservation {
	rv := objc.Send[RecognizedTextObservation](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRecognizedTextObservation creates a new RecognizedTextObservation instance.
func NewRecognizedTextObservation() RecognizedTextObservation {
	return getRecognizedTextObservationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RecognizedTextObservation */
// A request that detects and recognizes regions of text in an image.
//
// This type of observation results from a . It contains information about both the location and content of text and glyphs that Vision recognized in the input image.


// A request that detects and recognizes regions of text in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizedTextObservation
type RecognizedTextObservation struct {
	RectangleObservation
}

// RecognizedTextObservationFrom constructs a [RecognizedTextObservation] from an unsafe.Pointer.
//
// A request that detects and recognizes regions of text in an image.
func RecognizedTextObservationFrom(ptr unsafe.Pointer) RecognizedTextObservation {
	return RecognizedTextObservation{
		RectangleObservation: RectangleObservationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RecognizedTextObservation *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RecognizedTextObservation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RecognizedTextObservation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RecognizedTextObservation */

// Requests the top candidates for a recognized text string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizedTextObservation/topCandidates(_:)
func (r_ RecognizedTextObservation) TopCandidates(maxCandidateCount uint) []RecognizedText {
	rv := objc.Send[[]RecognizedText](r_.ID, objc.Sel("topCandidates:"), maxCandidateCount)
	return rv
}/* debug [instance_methods/method]: TopCandidates */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RecognizedTextObservation */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNRecognizedTextObservation */



