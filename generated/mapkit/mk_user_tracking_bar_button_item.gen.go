// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKUserTrackingBarButtonItem */


/* debug [class_header]: Header for MKUserTrackingBarButtonItem */
// The class instance for the [MKUserTrackingBarButtonItem] class.
var (
	MKUserTrackingBarButtonItemClass     _MKUserTrackingBarButtonItemClass
	MKUserTrackingBarButtonItemClassOnce sync.Once
)

func getMKUserTrackingBarButtonItemClass() _MKUserTrackingBarButtonItemClass {
	MKUserTrackingBarButtonItemClassOnce.Do(func() {
		MKUserTrackingBarButtonItemClass = _MKUserTrackingBarButtonItemClass{objc.GetClass("MKUserTrackingBarButtonItem")}
	})
	return MKUserTrackingBarButtonItemClass
}

type _MKUserTrackingBarButtonItemClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKUserTrackingBarButtonItem */
// An interface definition for the [MKUserTrackingBarButtonItem] class.
type IMKUserTrackingBarButtonItem interface {
	IBarButtonItem
	
/* debug [class_interface_properties]: Properties for MKUserTrackingBarButtonItem */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKUserTrackingBarButtonItem */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKUserTrackingBarButtonItem */
// Alloc allocates a new instance without initialization.
func (mc _MKUserTrackingBarButtonItemClass) Alloc() MKUserTrackingBarButtonItem {
	rv := objc.Send[MKUserTrackingBarButtonItem](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKUserTrackingBarButtonItemClass) New() MKUserTrackingBarButtonItem {
	rv := objc.Send[MKUserTrackingBarButtonItem](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKUserTrackingBarButtonItem) Init() MKUserTrackingBarButtonItem {
	rv := objc.Send[MKUserTrackingBarButtonItem](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKUserTrackingBarButtonItem) Autorelease() MKUserTrackingBarButtonItem {
	rv := objc.Send[MKUserTrackingBarButtonItem](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKUserTrackingBarButtonItem creates a new MKUserTrackingBarButtonItem instance.
func NewMKUserTrackingBarButtonItem() MKUserTrackingBarButtonItem {
	return getMKUserTrackingBarButtonItemClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKUserTrackingBarButtonItem */
// A specialized bar button item that allows the user to toggle whether the map tracks to the heading the user is facing.
//
// Tapping the button lets the user toggles between modes for displaying the map with and without the current heading applied. The button also reflects the current user tracking mode if set elsewhere. This bar button item is associated to a single map view.


// A specialized bar button item that allows the user to toggle whether the map tracks to the heading the user is facing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKUserTrackingBarButtonItem
type MKUserTrackingBarButtonItem struct {
	BarButtonItem
}

// MKUserTrackingBarButtonItemFrom constructs a [MKUserTrackingBarButtonItem] from an unsafe.Pointer.
//
// A specialized bar button item that allows the user to toggle whether the map tracks to the heading the user is facing.
func MKUserTrackingBarButtonItemFrom(ptr unsafe.Pointer) MKUserTrackingBarButtonItem {
	return MKUserTrackingBarButtonItem{
		BarButtonItem: BarButtonItemFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKUserTrackingBarButtonItem */

// Initializes a newly created bar button item with the specified map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKUserTrackingBarButtonItem/init(mapView:)
func NewMKUserTrackingBarButtonItemWithMapView(mapView IMKMapView) MKUserTrackingBarButtonItem {
	instance := getMKUserTrackingBarButtonItemClass().Alloc()
	rv := objc.Send[MKUserTrackingBarButtonItem](instance.ID, objc.Sel("initWithMapView:"), mapView)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKUserTrackingBarButtonItemWithMapView */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKUserTrackingBarButtonItem */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKUserTrackingBarButtonItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKUserTrackingBarButtonItem */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKUserTrackingBarButtonItem */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKUserTrackingBarButtonItem */


