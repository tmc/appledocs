// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VNRecognizeAnimalsRequest */


/* debug [class_header]: Header for VNRecognizeAnimalsRequest */
// The class instance for the [RecognizeAnimalsRequest] class.
var (
	RecognizeAnimalsRequestClass     _RecognizeAnimalsRequestClass
	RecognizeAnimalsRequestClassOnce sync.Once
)

func getRecognizeAnimalsRequestClass() _RecognizeAnimalsRequestClass {
	RecognizeAnimalsRequestClassOnce.Do(func() {
		RecognizeAnimalsRequestClass = _RecognizeAnimalsRequestClass{objc.GetClass("VNRecognizeAnimalsRequest")}
	})
	return RecognizeAnimalsRequestClass
}

type _RecognizeAnimalsRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RecognizeAnimalsRequest */
// An interface definition for the [RecognizeAnimalsRequest] class.
type IRecognizeAnimalsRequest interface {
	IImageBasedRequest
	
/* debug [class_interface_properties]: Properties for RecognizeAnimalsRequest */
	// properties:
	Results() []RecognizedObjectObservation
	VNRecognizeAnimalsRequestRevision1() int
	VNRecognizeAnimalsRequestRevision2() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RecognizeAnimalsRequest */
	// methods:
	SupportedIdentifiersAndReturnError(error_ objectivec.IObject) []string
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RecognizeAnimalsRequest */
// Alloc allocates a new instance without initialization.
func (rc _RecognizeAnimalsRequestClass) Alloc() RecognizeAnimalsRequest {
	rv := objc.Send[RecognizeAnimalsRequest](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RecognizeAnimalsRequestClass) New() RecognizeAnimalsRequest {
	rv := objc.Send[RecognizeAnimalsRequest](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RecognizeAnimalsRequest) Init() RecognizeAnimalsRequest {
	rv := objc.Send[RecognizeAnimalsRequest](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RecognizeAnimalsRequest) Autorelease() RecognizeAnimalsRequest {
	rv := objc.Send[RecognizeAnimalsRequest](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRecognizeAnimalsRequest creates a new RecognizeAnimalsRequest instance.
func NewRecognizeAnimalsRequest() RecognizeAnimalsRequest {
	return getRecognizeAnimalsRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RecognizeAnimalsRequest */
// A request that recognizes animals in an image.
//
// Use the method to determine which animals the request supports.


// A request that recognizes animals in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizeAnimalsRequest
type RecognizeAnimalsRequest struct {
	ImageBasedRequest
}

// RecognizeAnimalsRequestFrom constructs a [RecognizeAnimalsRequest] from an unsafe.Pointer.
//
// A request that recognizes animals in an image.
func RecognizeAnimalsRequestFrom(ptr unsafe.Pointer) RecognizeAnimalsRequest {
	return RecognizeAnimalsRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RecognizeAnimalsRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RecognizeAnimalsRequest */

// Returns a list of animal identifiers the recognition algorithm supports for the specified revision.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizeAnimalsRequest/knownAnimalIdentifiers(forRevision:)
func (rc _RecognizeAnimalsRequestClass) KnownAnimalIdentifiersForRevisionError(requestRevision uint, error_ objectivec.IObject) []string {
	rv := objc.Send[[]string](objc.ID(rc.class), objc.Sel("knownAnimalIdentifiersForRevision:error:"), requestRevision, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=KnownAnimalIdentifiersForRevisionError) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RecognizeAnimalsRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RecognizeAnimalsRequest */

// Returns the identifiers of the animals that the request detects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizeAnimalsRequest/supportedIdentifiers()
func (r_ RecognizeAnimalsRequest) SupportedIdentifiersAndReturnError(error_ objectivec.IObject) []string {
	rv := objc.Send[[]string](r_.ID, objc.Sel("supportedIdentifiersAndReturnError:"), error_)
	return rv
}/* debug [instance_methods/method]: SupportedIdentifiersAndReturnError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RecognizeAnimalsRequest */

// The results of the request to recognize animals.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizeAnimalsRequest/results
func (r_ RecognizeAnimalsRequest) Results() []RecognizedObjectObservation {
	rv := objc.Send[[]RecognizedObjectObservation](r_.ID, objc.Sel("results"))
	return rv
}/* debug [instance_properties/getter]: results */


// A constant for specifying revision 1 of the animal recognition request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizeanimalsrequestrevision1
func (r_ RecognizeAnimalsRequest) VNRecognizeAnimalsRequestRevision1() int {
	rv := objc.Send[int](r_.ID, objc.Sel("VNRecognizeAnimalsRequestRevision1"))
	return rv
}/* debug [instance_properties/getter]: VNRecognizeAnimalsRequestRevision1 */


// A constant for specifying revision 2 of the animal recognition request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizeanimalsrequestrevision2
func (r_ RecognizeAnimalsRequest) VNRecognizeAnimalsRequestRevision2() int {
	rv := objc.Send[int](r_.ID, objc.Sel("VNRecognizeAnimalsRequestRevision2"))
	return rv
}/* debug [instance_properties/getter]: VNRecognizeAnimalsRequestRevision2 */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNRecognizeAnimalsRequest */



