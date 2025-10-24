// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKLookAroundViewController */


/* debug [class_header]: Header for MKLookAroundViewController */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKLookAroundViewController */
// An interface definition for the [MKLookAroundViewController] class.
type IMKLookAroundViewController interface {
	appkit.IViewController
	
/* debug [class_interface_properties]: Properties for MKLookAroundViewController */
	// properties:
	BadgePosition() MKLookAroundBadgePosition
	SetBadgePosition(value MKLookAroundBadgePosition)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	NavigationEnabled() bool
	SetNavigationEnabled(value bool)
	PointOfInterestFilter() IMKPointOfInterestFilter
	SetPointOfInterestFilter(value IMKPointOfInterestFilter)
	Scene() IMKLookAroundScene
	SetScene(value IMKLookAroundScene)
	ShowsRoadLabels() bool
	SetShowsRoadLabels(value bool)
	IsNavigationEnabled() bool
	SetIsNavigationEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKLookAroundViewController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKLookAroundViewController */
// Alloc allocates a new instance without initialization.
func (mc _MKLookAroundViewControllerClass) Alloc() MKLookAroundViewController {
	rv := objc.Send[MKLookAroundViewController](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKLookAroundViewController */
// A class that manages the presentation and display of a LookAround view.


// A class that manages the presentation and display of a LookAround view.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKLookAroundViewController */

// Creates a new LookAround view controller object from a coder object provided by a storyboard or nib file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundViewController/init(coder:)
func NewMKLookAroundViewControllerWithCoder(coder foundation.Coder) MKLookAroundViewController {
	instance := getMKLookAroundViewControllerClass().Alloc()
	rv := objc.Send[MKLookAroundViewController](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKLookAroundViewControllerWithCoder */


// Creates a new LookAround view controller from the specified nib and bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundViewController/init(nibName:bundle:)
func NewMKLookAroundViewControllerWithNibNameBundle(nibNameOrNil objc.IObject /* cross-framework: NSString */, nibBundleOrNil foundation.Bundle) MKLookAroundViewController {
	instance := getMKLookAroundViewControllerClass().Alloc()
	rv := objc.Send[MKLookAroundViewController](instance.ID, objc.Sel("initWithNibName:bundle:"), nibNameOrNil, nibBundleOrNil)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKLookAroundViewControllerWithNibNameBundle */


// Creates a new LookAround view controller with the specified scene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundViewController/init(scene:)
func NewMKLookAroundViewControllerWithScene(scene IMKLookAroundScene) MKLookAroundViewController {
	instance := getMKLookAroundViewControllerClass().Alloc()
	rv := objc.Send[MKLookAroundViewController](instance.ID, objc.Sel("initWithScene:"), scene)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKLookAroundViewControllerWithScene */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKLookAroundViewController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKLookAroundViewController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKLookAroundViewController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKLookAroundViewController */

// A value that indicates the badge’s position on the LookAround view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundViewController/badgePosition
func (m_ MKLookAroundViewController) BadgePosition() MKLookAroundBadgePosition {
	rv := objc.Send[MKLookAroundBadgePosition](m_.ID, objc.Sel("badgePosition"))
	return rv
}/* debug [instance_properties/getter]: badgePosition */


// A value that indicates the badge’s position on the LookAround view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundViewController/badgePosition
func (m_ MKLookAroundViewController) SetBadgePosition(value MKLookAroundBadgePosition) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBadgePosition:"), value)
}/* debug [instance_properties/setter]: badgePosition */


// An object you provide to receive events related to the user’s interaction with the LookAround view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundViewController/delegate
func (m_ MKLookAroundViewController) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// An object you provide to receive events related to the user’s interaction with the LookAround view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundViewController/delegate
func (m_ MKLookAroundViewController) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// A Boolean value that indicates whether the map’s navigation controls are visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundViewController/isNavigationEnabled
func (m_ MKLookAroundViewController) NavigationEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("navigationEnabled"))
	return rv
}/* debug [instance_properties/getter]: navigationEnabled */


// A Boolean value that indicates whether the map’s navigation controls are visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundViewController/isNavigationEnabled
func (m_ MKLookAroundViewController) SetNavigationEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNavigationEnabled:"), value)
}/* debug [instance_properties/setter]: navigationEnabled */


// The filter used to determine the points of interest shown on the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundViewController/pointOfInterestFilter
func (m_ MKLookAroundViewController) PointOfInterestFilter() IMKPointOfInterestFilter {
	rv := objc.Send[MKPointOfInterestFilter](m_.ID, objc.Sel("pointOfInterestFilter"))
	return rv
}/* debug [instance_properties/getter]: pointOfInterestFilter */


// The filter used to determine the points of interest shown on the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundViewController/pointOfInterestFilter
func (m_ MKLookAroundViewController) SetPointOfInterestFilter(value IMKPointOfInterestFilter) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPointOfInterestFilter:"), value)
}/* debug [instance_properties/setter]: pointOfInterestFilter */


// The LookAround scene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundViewController/scene
func (m_ MKLookAroundViewController) Scene() IMKLookAroundScene {
	rv := objc.Send[MKLookAroundScene](m_.ID, objc.Sel("scene"))
	return rv
}/* debug [instance_properties/getter]: scene */


// The LookAround scene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundViewController/scene
func (m_ MKLookAroundViewController) SetScene(value IMKLookAroundScene) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setScene:"), value)
}/* debug [instance_properties/setter]: scene */


// A Boolean value that indicates whether the map display road labels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundViewController/showsRoadLabels
func (m_ MKLookAroundViewController) ShowsRoadLabels() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsRoadLabels"))
	return rv
}/* debug [instance_properties/getter]: showsRoadLabels */


// A Boolean value that indicates whether the map display road labels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundViewController/showsRoadLabels
func (m_ MKLookAroundViewController) SetShowsRoadLabels(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsRoadLabels:"), value)
}/* debug [instance_properties/setter]: showsRoadLabels */


// A Boolean value that indicates whether the map’s navigation controls are visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklookaroundviewcontroller/isnavigationenabled
func (m_ MKLookAroundViewController) IsNavigationEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isNavigationEnabled"))
	return rv
}/* debug [instance_properties/getter]: isNavigationEnabled */


// A Boolean value that indicates whether the map’s navigation controls are visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklookaroundviewcontroller/isnavigationenabled
func (m_ MKLookAroundViewController) SetIsNavigationEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsNavigationEnabled:"), value)
}/* debug [instance_properties/setter]: isNavigationEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKLookAroundViewController */


