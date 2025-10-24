// Code generated from Apple documentation for PhotosUI. DO NOT EDIT.

package photosui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHProjectRegionOfInterest] class.
var (
	PHProjectRegionOfInterestClass     _PHProjectRegionOfInterestClass
	PHProjectRegionOfInterestClassOnce sync.Once
)

func getPHProjectRegionOfInterestClass() _PHProjectRegionOfInterestClass {
	PHProjectRegionOfInterestClassOnce.Do(func() {
		PHProjectRegionOfInterestClass = _PHProjectRegionOfInterestClass{objc.GetClass("PHProjectRegionOfInterest")}
	})
	return PHProjectRegionOfInterestClass
}

type _PHProjectRegionOfInterestClass struct {
	class objc.Class
}

// An interface definition for the [PHProjectRegionOfInterest] class.
type IPHProjectRegionOfInterest interface {
	objectivec.IObject
	// properties:
	Quality() float64
	Weight() float64
	Identifier() unsafe.Pointer
	SetIdentifier(value unsafe.Pointer)
	Rect() objc.IObject /* cross-framework: Rect */
	SetRect(value objc.IObject /* cross-framework: Rect */)
	// methods:
}

// A representation of a region of interest in a photo asset.
//
// A region of interest defines a rectangular portion of a photo corresponding to a face. Use a region of interest to determine where to focus, zoom, or crop your image in your project extension; for example, you can customize your slideshow’s transitions based on each photo’s highest-quality region of interest, as shown in .


// A representation of a region of interest in a photo asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectRegionOfInterest
type PHProjectRegionOfInterest struct {
	objectivec.Object
}

// PHProjectRegionOfInterestFrom constructs a [PHProjectRegionOfInterest] from an unsafe.Pointer.
//
// A representation of a region of interest in a photo asset.
func PHProjectRegionOfInterestFrom(ptr unsafe.Pointer) PHProjectRegionOfInterest {
	return PHProjectRegionOfInterest{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHProjectRegionOfInterestClass) Alloc() PHProjectRegionOfInterest {
	rv := objc.Send[PHProjectRegionOfInterest](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHProjectRegionOfInterestClass) New() PHProjectRegionOfInterest {
	rv := objc.Send[PHProjectRegionOfInterest](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHProjectRegionOfInterest) Init() PHProjectRegionOfInterest {
	rv := objc.Send[PHProjectRegionOfInterest](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHProjectRegionOfInterest) Autorelease() PHProjectRegionOfInterest {
	rv := objc.Send[PHProjectRegionOfInterest](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHProjectRegionOfInterest creates a new PHProjectRegionOfInterest instance.
func NewPHProjectRegionOfInterest() PHProjectRegionOfInterest {
	return getPHProjectRegionOfInterestClass().New()
}



// The region’s quality.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectRegionOfInterest/quality
func (p_ PHProjectRegionOfInterest) Quality() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("quality"))
	return rv
}


// The face region’s weight.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectRegionOfInterest/weight
func (p_ PHProjectRegionOfInterest) Weight() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("weight"))
	return rv
}


// The region’s unique identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photosui/phprojectregionofinterest/identifier-swift.property
func (p_ PHProjectRegionOfInterest) Identifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("identifier"))
	return rv
}


// The region’s unique identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photosui/phprojectregionofinterest/identifier-swift.property
func (p_ PHProjectRegionOfInterest) SetIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIdentifier:"), value)
}


// The rectangle representing the region’s location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photosui/phprojectregionofinterest/rect
func (p_ PHProjectRegionOfInterest) Rect() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](p_.ID, objc.Sel("rect"))
	return rv
}


// The rectangle representing the region’s location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photosui/phprojectregionofinterest/rect
func (p_ PHProjectRegionOfInterest) SetRect(value objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRect:"), value)
}



