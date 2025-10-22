// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/gameplaykit"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Contour] class.
var (
	ContourClass     _ContourClass
	ContourClassOnce sync.Once
)

func getContourClass() _ContourClass {
	ContourClassOnce.Do(func() {
		ContourClass = _ContourClass{objc.GetClass("VNContour")}
	})
	return ContourClass
}

type _ContourClass struct {
	class objc.Class
}

// An interface definition for the [Contour] class.
type IContour interface {
	objectivec.IObject
	AspectRatio() float32
	SetAspectRatio(value float32)
	ChildContourCount() int
	SetChildContourCount(value int)
	ChildContours() VNContour
	SetChildContours(value IVNContour)
	IndexPath() foundation.IndexPath
	SetIndexPath(value foundation.IIndexPath)
	NormalizedPath() gameplaykit.Path
	SetNormalizedPath(value gameplaykit.IPath)
	NormalizedPoints() unsafe.Pointer
	SetNormalizedPoints(value unsafe.Pointer)
	PointCount() int
	SetPointCount(value int)
	ContourCount() int
	SetContourCount(value int)
	TopLevelContourCount() int
	SetTopLevelContourCount(value int)
	TopLevelContours() VNContour
	SetTopLevelContours(value IVNContour)
}

// A class that represents a detected contour in an image.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNContour
type Contour struct {
	objectivec.Object
}

// ContourFrom constructs a [Contour] from an unsafe.Pointer.
//
// A class that represents a detected contour in an image.
func ContourFrom(ptr unsafe.Pointer) Contour {
	return Contour{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _ContourClass) Alloc() Contour {
	rv := objc.Send[Contour](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ContourClass) New() Contour {
	rv := objc.Send[Contour](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ Contour) Init() Contour {
	rv := objc.Send[Contour](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ Contour) Autorelease() Contour {
	rv := objc.Send[Contour](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewContour creates a new Contour instance.
func NewContour() Contour {
	return getContourClass().New()
}


// The aspect ratio of the contour.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontour/aspectratio
func (c_ Contour) AspectRatio() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("aspectRatio"))
	return rv
}


// SetAspectRatio sets the value of the aspectRatio property.
// The aspect ratio of the contour.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontour/aspectratio
func (c_ Contour) SetAspectRatio(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAspectRatio:"), value)
}

// The total number of detected child contours.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontour/childcontourcount
func (c_ Contour) ChildContourCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("childContourCount"))
	return rv
}


// SetChildContourCount sets the value of the childContourCount property.
// The total number of detected child contours.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontour/childcontourcount
func (c_ Contour) SetChildContourCount(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setChildContourCount:"), value)
}

// An array of contours that this contour encloses.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontour/childcontours
func (c_ Contour) ChildContours() VNContour {
	rv := objc.Send[VNContour](c_.ID, objc.Sel("childContours"))
	return rv
}


// SetChildContours sets the value of the childContours property.
// An array of contours that this contour encloses.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontour/childcontours
func (c_ Contour) SetChildContours(value IVNContour) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setChildContours:"), value)
}

// The contour object’s index path.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontour/indexpath
func (c_ Contour) IndexPath() foundation.IndexPath {
	rv := objc.Send[foundation.IndexPath](c_.ID, objc.Sel("indexPath"))
	return rv
}


// SetIndexPath sets the value of the indexPath property.
// The contour object’s index path.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontour/indexpath
func (c_ Contour) SetIndexPath(value foundation.IIndexPath) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIndexPath:"), value)
}

// The contour object as a path in normalized coordinates.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontour/normalizedpath
func (c_ Contour) NormalizedPath() gameplaykit.Path {
	rv := objc.Send[gameplaykit.Path](c_.ID, objc.Sel("normalizedPath"))
	return rv
}


// SetNormalizedPath sets the value of the normalizedPath property.
// The contour object as a path in normalized coordinates.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontour/normalizedpath
func (c_ Contour) SetNormalizedPath(value gameplaykit.IPath) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNormalizedPath:"), value)
}

// The contour’s array of points in normalized coordinates.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontour/normalizedpoints-8n2s5
func (c_ Contour) NormalizedPoints() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("normalizedPoints"))
	return rv
}


// SetNormalizedPoints sets the value of the normalizedPoints property.
// The contour’s array of points in normalized coordinates.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontour/normalizedpoints-8n2s5
func (c_ Contour) SetNormalizedPoints(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNormalizedPoints:"), value)
}

// The contour’s number of points.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontour/pointcount
func (c_ Contour) PointCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("pointCount"))
	return rv
}


// SetPointCount sets the value of the pointCount property.
// The contour’s number of points.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontour/pointcount
func (c_ Contour) SetPointCount(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPointCount:"), value)
}

// The total number of detected contours.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontoursobservation/contourcount
func (c_ Contour) ContourCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("contourCount"))
	return rv
}


// SetContourCount sets the value of the contourCount property.
// The total number of detected contours.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontoursobservation/contourcount
func (c_ Contour) SetContourCount(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContourCount:"), value)
}

// The total number of detected top-level contours.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontoursobservation/toplevelcontourcount
func (c_ Contour) TopLevelContourCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("topLevelContourCount"))
	return rv
}


// SetTopLevelContourCount sets the value of the topLevelContourCount property.
// The total number of detected top-level contours.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontoursobservation/toplevelcontourcount
func (c_ Contour) SetTopLevelContourCount(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTopLevelContourCount:"), value)
}

// An array of contours that don’t have another contour enclosing them.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontoursobservation/toplevelcontours
func (c_ Contour) TopLevelContours() VNContour {
	rv := objc.Send[VNContour](c_.ID, objc.Sel("topLevelContours"))
	return rv
}


// SetTopLevelContours sets the value of the topLevelContours property.
// An array of contours that don’t have another contour enclosing them.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontoursobservation/toplevelcontours
func (c_ Contour) SetTopLevelContours(value IVNContour) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTopLevelContours:"), value)
}



