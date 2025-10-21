// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [MKUserTrackingButton] class.
var (
	MKUserTrackingButtonClass     _MKUserTrackingButtonClass
	MKUserTrackingButtonClassOnce sync.Once
)

func getMKUserTrackingButtonClass() _MKUserTrackingButtonClass {
	MKUserTrackingButtonClassOnce.Do(func() {
		MKUserTrackingButtonClass = _MKUserTrackingButtonClass{objc.GetClass("MKUserTrackingButton")}
	})
	return MKUserTrackingButtonClass
}

type _MKUserTrackingButtonClass struct {
	class objc.Class
}

// An interface definition for the [MKUserTrackingButton] class.
type IMKUserTrackingButton interface {
	appkit.IView
}

// A specialized button that allows the user to toggle whether the map tracks to the heading the user is facing.
//
// Use this class when you need a standard button that you can incorporate into your view hierarchy. Tapping the button lets the user toggles between modes for displaying the map with and without the current heading applied. The button also reflects the current user tracking mode if set elsewhere.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKUserTrackingButton
type MKUserTrackingButton struct {
	appkit.View
}

// MKUserTrackingButtonFrom constructs a [MKUserTrackingButton] from an unsafe.Pointer.
//
// A specialized button that allows the user to toggle whether the map tracks to the heading the user is facing.
func MKUserTrackingButtonFrom(ptr unsafe.Pointer) MKUserTrackingButton {
	return MKUserTrackingButton{
		View: appkit.ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MKUserTrackingButtonClass) Alloc() MKUserTrackingButton {
	rv := objc.Send[MKUserTrackingButton](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKUserTrackingButtonClass) New() MKUserTrackingButton {
	rv := objc.Send[MKUserTrackingButton](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKUserTrackingButton) Init() MKUserTrackingButton {
	rv := objc.Send[MKUserTrackingButton](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKUserTrackingButton) Autorelease() MKUserTrackingButton {
	rv := objc.Send[MKUserTrackingButton](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKUserTrackingButton creates a new MKUserTrackingButton instance.
func NewMKUserTrackingButton() MKUserTrackingButton {
	return getMKUserTrackingButtonClass().New()
}




// Initializes the button with the map view that it should control.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKUserTrackingButton/init(mapView:)
func NewMKUserTrackingButtonWithMapView(mapView IMKMapView) MKUserTrackingButton {
	rv := objc.Send[MKUserTrackingButton](objc.ID(getMKUserTrackingButtonClass().class), objc.Sel("userTrackingButtonWithMapView:"), mapView)
	return rv
}


// Initializes the button with the map view that it should control.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKUserTrackingButton/init(mapView:)
func (mc _MKUserTrackingButtonClass) UserTrackingButtonWithMapView(mapView IMKMapView) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("userTrackingButtonWithMapView:"), mapView)
	return rv
}

// The map view associated with the button.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKUserTrackingButton/mapView
func (m_ MKUserTrackingButton) MapView() MKMapView {
	rv := objc.Send[MKMapView](m_.ID, objc.Sel("mapView"))
	return rv
}


// SetMapView sets the value of the mapView property.
// The map view associated with the button.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKUserTrackingButton/mapView
func (m_ MKUserTrackingButton) SetMapView(value IMKMapView) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMapView:"), value)
}


