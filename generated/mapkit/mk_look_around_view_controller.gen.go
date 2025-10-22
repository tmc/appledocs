// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [MKLookAroundViewController] class.
var (
	MKLookAroundViewControllerClass     _MKLookAroundViewControllerClass
	MKLookAroundViewControllerClassOnce sync.Once
)

func getMKLookAroundViewControllerClass() _MKLookAroundViewControllerClass {
	MKLookAroundViewControllerClassOnce.Do(func() {
		MKLookAroundViewControllerClass = _MKLookAroundViewControllerClass{objc.GetClass("MKLookAroundViewController")}
	})
	return MKLookAroundViewControllerClass
}

type _MKLookAroundViewControllerClass struct {
	class objc.Class
}

// An interface definition for the [MKLookAroundViewController] class.
type IMKLookAroundViewController interface {
	appkit.IViewController
	BadgePosition() unsafe.Pointer
	SetBadgePosition(value unsafe.Pointer)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	IsNavigationEnabled() bool
	SetIsNavigationEnabled(value bool)
	PointOfInterestFilter() MKPointOfInterestFilter
	SetPointOfInterestFilter(value IMKPointOfInterestFilter)
	Scene() MKLookAroundScene
	SetScene(value IMKLookAroundScene)
	ShowsRoadLabels() bool
	SetShowsRoadLabels(value bool)
}

// A class that manages the presentation and display of a LookAround view.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundViewController
type MKLookAroundViewController struct {
	appkit.ViewController
}

// MKLookAroundViewControllerFrom constructs a [MKLookAroundViewController] from an unsafe.Pointer.
//
// A class that manages the presentation and display of a LookAround view.
func MKLookAroundViewControllerFrom(ptr unsafe.Pointer) MKLookAroundViewController {
	return MKLookAroundViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MKLookAroundViewControllerClass) Alloc() MKLookAroundViewController {
	rv := objc.Send[MKLookAroundViewController](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKLookAroundViewControllerClass) New() MKLookAroundViewController {
	rv := objc.Send[MKLookAroundViewController](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKLookAroundViewController) Init() MKLookAroundViewController {
	rv := objc.Send[MKLookAroundViewController](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKLookAroundViewController) Autorelease() MKLookAroundViewController {
	rv := objc.Send[MKLookAroundViewController](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKLookAroundViewController creates a new MKLookAroundViewController instance.
func NewMKLookAroundViewController() MKLookAroundViewController {
	return getMKLookAroundViewControllerClass().New()
}


// A value that indicates the badge’s position on the LookAround view.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklookaroundviewcontroller/badgeposition
func (m_ MKLookAroundViewController) BadgePosition() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("badgePosition"))
	return rv
}


// SetBadgePosition sets the value of the badgePosition property.
// A value that indicates the badge’s position on the LookAround view.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklookaroundviewcontroller/badgeposition
func (m_ MKLookAroundViewController) SetBadgePosition(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBadgePosition:"), value)
}

// An object you provide to receive events related to the user’s interaction with the LookAround view controller.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklookaroundviewcontroller/delegate
func (m_ MKLookAroundViewController) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// An object you provide to receive events related to the user’s interaction with the LookAround view controller.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklookaroundviewcontroller/delegate
func (m_ MKLookAroundViewController) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean value that indicates whether the map’s navigation controls are visible.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklookaroundviewcontroller/isnavigationenabled
func (m_ MKLookAroundViewController) IsNavigationEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isNavigationEnabled"))
	return rv
}


// SetIsNavigationEnabled sets the value of the isNavigationEnabled property.
// A Boolean value that indicates whether the map’s navigation controls are visible.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklookaroundviewcontroller/isnavigationenabled
func (m_ MKLookAroundViewController) SetIsNavigationEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsNavigationEnabled:"), value)
}

// The filter used to determine the points of interest shown on the map.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklookaroundviewcontroller/pointofinterestfilter
func (m_ MKLookAroundViewController) PointOfInterestFilter() MKPointOfInterestFilter {
	rv := objc.Send[MKPointOfInterestFilter](m_.ID, objc.Sel("pointOfInterestFilter"))
	return rv
}


// SetPointOfInterestFilter sets the value of the pointOfInterestFilter property.
// The filter used to determine the points of interest shown on the map.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklookaroundviewcontroller/pointofinterestfilter
func (m_ MKLookAroundViewController) SetPointOfInterestFilter(value IMKPointOfInterestFilter) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPointOfInterestFilter:"), value)
}

// The LookAround scene.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklookaroundviewcontroller/scene
func (m_ MKLookAroundViewController) Scene() MKLookAroundScene {
	rv := objc.Send[MKLookAroundScene](m_.ID, objc.Sel("scene"))
	return rv
}


// SetScene sets the value of the scene property.
// The LookAround scene.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklookaroundviewcontroller/scene
func (m_ MKLookAroundViewController) SetScene(value IMKLookAroundScene) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setScene:"), value)
}

// A Boolean value that indicates whether the map display road labels.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklookaroundviewcontroller/showsroadlabels
func (m_ MKLookAroundViewController) ShowsRoadLabels() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsRoadLabels"))
	return rv
}


// SetShowsRoadLabels sets the value of the showsRoadLabels property.
// A Boolean value that indicates whether the map display road labels.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklookaroundviewcontroller/showsroadlabels
func (m_ MKLookAroundViewController) SetShowsRoadLabels(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsRoadLabels:"), value)
}



