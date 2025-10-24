// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [ContoursObservation] class.
type IContoursObservation interface {
	IObservation
	

	// properties:
	ContourCount() int
	NormalizedPath() PathRef /* not a class type */
	TopLevelContourCount() int
	TopLevelContours() []Contour
	Results() IVNContoursObservation
	SetResults(value IVNContoursObservation)


	

	// methods:
	ContourAtIndexPathError(indexPath foundation.IndexPath, error_ objectivec.IObject) IContour
	ContourAtIndexError(contourIndex int, error_ objectivec.IObject) IContour


}





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




















// Retrieves the contour object at the specified index path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNContoursObservation/contour(at:)-52odo
func (c_ ContoursObservation) ContourAtIndexPathError(indexPath foundation.IndexPath, error_ objectivec.IObject) IContour {
	rv := objc.Send[Contour](c_.ID, objc.Sel("contourAtIndexPath:error:"), indexPath, error_)
	return rv
}


// Retrieves the contour object at the specified index, irrespective of hierarchy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNContoursObservation/contour(at:)-9on0y
func (c_ ContoursObservation) ContourAtIndexError(contourIndex int, error_ objectivec.IObject) IContour {
	rv := objc.Send[Contour](c_.ID, objc.Sel("contourAtIndex:error:"), contourIndex, error_)
	return rv
}







// The total number of detected contours.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNContoursObservation/contourCount
func (c_ ContoursObservation) ContourCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("contourCount"))
	return rv
}


// The detected contours as a path object in normalized coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNContoursObservation/normalizedPath
func (c_ ContoursObservation) NormalizedPath() PathRef /* not a class type */ {
	rv := objc.Send[PathRef](c_.ID, objc.Sel("normalizedPath"))
	return rv
}


// The total number of detected top-level contours.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNContoursObservation/topLevelContourCount
func (c_ ContoursObservation) TopLevelContourCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("topLevelContourCount"))
	return rv
}


// An array of contours that don’t have another contour enclosing them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNContoursObservation/topLevelContours
func (c_ ContoursObservation) TopLevelContours() []Contour {
	rv := objc.Send[[]Contour](c_.ID, objc.Sel("topLevelContours"))
	return rv
}


// The results of the request to detect contours.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectcontoursrequest/results
func (c_ ContoursObservation) Results() IVNContoursObservation {
	rv := objc.Send[ContoursObservation](c_.ID, objc.Sel("results"))
	return rv
}


// The results of the request to detect contours.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectcontoursrequest/results
func (c_ ContoursObservation) SetResults(value IVNContoursObservation) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setResults:"), value)
}








