// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MKUserLocationView] class.
var (
	MKUserLocationViewClass     _MKUserLocationViewClass
	MKUserLocationViewClassOnce sync.Once
)

func getMKUserLocationViewClass() _MKUserLocationViewClass {
	MKUserLocationViewClassOnce.Do(func() {
		MKUserLocationViewClass = _MKUserLocationViewClass{objc.GetClass("MKUserLocationView")}
	})
	return MKUserLocationViewClass
}

type _MKUserLocationViewClass struct {
	class objc.Class
}

// An interface definition for the [MKUserLocationView] class.
type IMKUserLocationView interface {
	IMKAnnotationView
}

// A configurable annotation that shows the user’s location using the default MapKit style.
//
// If you don’t need additional configuration, you can show an annotation with the user’s location by setting on the map to . If you want to specify additional configuration, such as , create this annotation view directly. To display the annotation view, return the instance from . The user location view provides the MapKit default style and behavior. The visual display varies with the level of authorization the user grants your app.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKUserLocationView
type MKUserLocationView struct {
	MKAnnotationView
}

// MKUserLocationViewFrom constructs a [MKUserLocationView] from an unsafe.Pointer.
//
// A configurable annotation that shows the user’s location using the default MapKit style.
func MKUserLocationViewFrom(ptr unsafe.Pointer) MKUserLocationView {
	return MKUserLocationView{
		MKAnnotationView: MKAnnotationViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MKUserLocationViewClass) Alloc() MKUserLocationView {
	rv := objc.Send[MKUserLocationView](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKUserLocationViewClass) New() MKUserLocationView {
	rv := objc.Send[MKUserLocationView](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKUserLocationView) Init() MKUserLocationView {
	rv := objc.Send[MKUserLocationView](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKUserLocationView) Autorelease() MKUserLocationView {
	rv := objc.Send[MKUserLocationView](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKUserLocationView creates a new MKUserLocationView instance.
func NewMKUserLocationView() MKUserLocationView {
	return getMKUserLocationViewClass().New()
}


// A Boolean value that indicates whether the map tries to display the user’s location.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/showsuserlocation
func (m_ MKUserLocationView) ShowsUserLocation() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsUserLocation"))
	return rv
}


// SetShowsUserLocation sets the value of the showsUserLocation property.
// A Boolean value that indicates whether the map tries to display the user’s location.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/showsuserlocation
func (m_ MKUserLocationView) SetShowsUserLocation(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsUserLocation:"), value)
}

// The relative importance of the annotation view when in an unselected state with respect to its ordering along the z-axis.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/zpriority
func (m_ MKUserLocationView) ZPriority() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("zPriority"))
	return rv
}


// SetZPriority sets the value of the zPriority property.
// The relative importance of the annotation view when in an unselected state with respect to its ordering along the z-axis.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/zpriority
func (m_ MKUserLocationView) SetZPriority(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setZPriority:"), value)
}



