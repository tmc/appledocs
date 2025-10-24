// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VNContoursObservation */


/* debug [class_header]: Header for VNContoursObservation */
// The class instance for the [ContoursObservation] class.
var (
	ContoursObservationClass     _ContoursObservationClass
	ContoursObservationClassOnce sync.Once
)

func getContoursObservationClass() _ContoursObservationClass {
	ContoursObservationClassOnce.Do(func() {
		ContoursObservationClass = _ContoursObservationClass{objc.GetClass("VNContoursObservation")}
	})
	return ContoursObservationClass
}

type _ContoursObservationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ContoursObservation */
// An interface definition for the [ContoursObservation] class.
type IContoursObservation interface {
	IObservation
	
/* debug [class_interface_properties]: Properties for ContoursObservation */
	// properties:
	ContourCount() int
	NormalizedPath() PathRef /* not a class type */
	TopLevelContourCount() int
	TopLevelContours() []Contour
	Results() IVNContoursObservation
	SetResults(value IVNContoursObservation)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ContoursObservation */
	// methods:
	ContourAtIndexPathError(indexPath foundation.IndexPath, error_ objectivec.IObject) IContour
	ContourAtIndexError(contourIndex int, error_ objectivec.IObject) IContour
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ContoursObservation */
// Alloc allocates a new instance without initialization.
func (cc _ContoursObservationClass) Alloc() ContoursObservation {
	rv := objc.Send[ContoursObservation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ContoursObservationClass) New() ContoursObservation {
	rv := objc.Send[ContoursObservation](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ContoursObservation) Init() ContoursObservation {
	rv := objc.Send[ContoursObservation](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ContoursObservation) Autorelease() ContoursObservation {
	rv := objc.Send[ContoursObservation](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewContoursObservation creates a new ContoursObservation instance.
func NewContoursObservation() ContoursObservation {
	return getContoursObservationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ContoursObservation */
// An object that represents the detected contours in an image.


// An object that represents the detected contours in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNContoursObservation
type ContoursObservation struct {
	Observation
}

// ContoursObservationFrom constructs a [ContoursObservation] from an unsafe.Pointer.
//
// An object that represents the detected contours in an image.
func ContoursObservationFrom(ptr unsafe.Pointer) ContoursObservation {
	return ContoursObservation{
		Observation: ObservationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ContoursObservation *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ContoursObservation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ContoursObservation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ContoursObservation */

// Retrieves the contour object at the specified index path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNContoursObservation/contour(at:)-52odo
func (c_ ContoursObservation) ContourAtIndexPathError(indexPath foundation.IndexPath, error_ objectivec.IObject) IContour {
	rv := objc.Send[Contour](c_.ID, objc.Sel("contourAtIndexPath:error:"), indexPath, error_)
	return rv
}/* debug [instance_methods/method]: ContourAtIndexPathError */


// Retrieves the contour object at the specified index, irrespective of hierarchy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNContoursObservation/contour(at:)-9on0y
func (c_ ContoursObservation) ContourAtIndexError(contourIndex int, error_ objectivec.IObject) IContour {
	rv := objc.Send[Contour](c_.ID, objc.Sel("contourAtIndex:error:"), contourIndex, error_)
	return rv
}/* debug [instance_methods/method]: ContourAtIndexError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ContoursObservation */

// The total number of detected contours.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNContoursObservation/contourCount
func (c_ ContoursObservation) ContourCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("contourCount"))
	return rv
}/* debug [instance_properties/getter]: contourCount */


// The detected contours as a path object in normalized coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNContoursObservation/normalizedPath
func (c_ ContoursObservation) NormalizedPath() PathRef /* not a class type */ {
	rv := objc.Send[PathRef](c_.ID, objc.Sel("normalizedPath"))
	return rv
}/* debug [instance_properties/getter]: normalizedPath */


// The total number of detected top-level contours.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNContoursObservation/topLevelContourCount
func (c_ ContoursObservation) TopLevelContourCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("topLevelContourCount"))
	return rv
}/* debug [instance_properties/getter]: topLevelContourCount */


// An array of contours that don’t have another contour enclosing them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNContoursObservation/topLevelContours
func (c_ ContoursObservation) TopLevelContours() []Contour {
	rv := objc.Send[[]Contour](c_.ID, objc.Sel("topLevelContours"))
	return rv
}/* debug [instance_properties/getter]: topLevelContours */


// The results of the request to detect contours.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectcontoursrequest/results
func (c_ ContoursObservation) Results() IVNContoursObservation {
	rv := objc.Send[ContoursObservation](c_.ID, objc.Sel("results"))
	return rv
}/* debug [instance_properties/getter]: results */


// The results of the request to detect contours.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectcontoursrequest/results
func (c_ ContoursObservation) SetResults(value IVNContoursObservation) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setResults:"), value)
}/* debug [instance_properties/setter]: results */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNContoursObservation */



