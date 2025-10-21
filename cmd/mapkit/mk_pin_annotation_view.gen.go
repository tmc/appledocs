// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MKPinAnnotationView] class.
var (
	MKPinAnnotationViewClass     _MKPinAnnotationViewClass
	MKPinAnnotationViewClassOnce sync.Once
)

func getMKPinAnnotationViewClass() _MKPinAnnotationViewClass {
	MKPinAnnotationViewClassOnce.Do(func() {
		MKPinAnnotationViewClass = _MKPinAnnotationViewClass{objc.GetClass("MKPinAnnotationView")}
	})
	return MKPinAnnotationViewClass
}

type _MKPinAnnotationViewClass struct {
	class objc.Class
}

// An interface definition for the [MKPinAnnotationView] class.
type IMKPinAnnotationView interface {
	IMKAnnotationView
}

// An annotation view that displays a pin image on the map.
//
// Return instances of this class from the method of your map view delegate when you want to display a pin for one of your annotations. The pins displayed by this view are the same ones found in the Maps application. You can specify the type of pin you want to display and whether you want the pin to be animated into place.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPinAnnotationView
type MKPinAnnotationView struct {
	MKAnnotationView
}

// MKPinAnnotationViewFrom constructs a [MKPinAnnotationView] from an unsafe.Pointer.
//
// An annotation view that displays a pin image on the map.
func MKPinAnnotationViewFrom(ptr unsafe.Pointer) MKPinAnnotationView {
	return MKPinAnnotationView{
		MKAnnotationView: MKAnnotationViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MKPinAnnotationViewClass) Alloc() MKPinAnnotationView {
	rv := objc.Send[MKPinAnnotationView](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKPinAnnotationViewClass) New() MKPinAnnotationView {
	rv := objc.Send[MKPinAnnotationView](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKPinAnnotationView) Init() MKPinAnnotationView {
	rv := objc.Send[MKPinAnnotationView](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKPinAnnotationView) Autorelease() MKPinAnnotationView {
	rv := objc.Send[MKPinAnnotationView](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKPinAnnotationView creates a new MKPinAnnotationView instance.
func NewMKPinAnnotationView() MKPinAnnotationView {
	return getMKPinAnnotationViewClass().New()
}




