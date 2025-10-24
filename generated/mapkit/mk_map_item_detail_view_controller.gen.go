// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKMapItemDetailViewController */


/* debug [class_header]: Header for MKMapItemDetailViewController */
// The class instance for the [MKMapItemDetailViewController] class.
var (
	MKMapItemDetailViewControllerClass     _MKMapItemDetailViewControllerClass
	MKMapItemDetailViewControllerClassOnce sync.Once
)

func getMKMapItemDetailViewControllerClass() _MKMapItemDetailViewControllerClass {
	MKMapItemDetailViewControllerClassOnce.Do(func() {
		MKMapItemDetailViewControllerClass = _MKMapItemDetailViewControllerClass{objc.GetClass("MKMapItemDetailViewController")}
	})
	return MKMapItemDetailViewControllerClass
}

type _MKMapItemDetailViewControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKMapItemDetailViewController */
// An interface definition for the [MKMapItemDetailViewController] class.
type IMKMapItemDetailViewController interface {
	appkit.IViewController
	
/* debug [class_interface_properties]: Properties for MKMapItemDetailViewController */
	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	MapItem() IMKMapItem
	SetMapItem(value IMKMapItem)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKMapItemDetailViewController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKMapItemDetailViewController */
// Alloc allocates a new instance without initialization.
func (mc _MKMapItemDetailViewControllerClass) Alloc() MKMapItemDetailViewController {
	rv := objc.Send[MKMapItemDetailViewController](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKMapItemDetailViewControllerClass) New() MKMapItemDetailViewController {
	rv := objc.Send[MKMapItemDetailViewController](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKMapItemDetailViewController) Init() MKMapItemDetailViewController {
	rv := objc.Send[MKMapItemDetailViewController](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKMapItemDetailViewController) Autorelease() MKMapItemDetailViewController {
	rv := objc.Send[MKMapItemDetailViewController](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKMapItemDetailViewController creates a new MKMapItemDetailViewController instance.
func NewMKMapItemDetailViewController() MKMapItemDetailViewController {
	return getMKMapItemDetailViewControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKMapItemDetailViewController */
// An object that displays detailed information about a map item.
//
// The view controller presents modally and displays place information such as addresses and phone numbers. This class doesn’t support subclassing. The view hierarchy for this class is private and must not be modified.


// An object that displays detailed information about a map item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItemDetailViewController
type MKMapItemDetailViewController struct {
	appkit.ViewController
}

// MKMapItemDetailViewControllerFrom constructs a [MKMapItemDetailViewController] from an unsafe.Pointer.
//
// An object that displays detailed information about a map item.
func MKMapItemDetailViewControllerFrom(ptr unsafe.Pointer) MKMapItemDetailViewController {
	return MKMapItemDetailViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKMapItemDetailViewController */

// Create a map item detail view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItemDetailViewController/init(mapItem:)
func NewMKMapItemDetailViewControllerWithMapItem(mapItem IMKMapItem) MKMapItemDetailViewController {
	instance := getMKMapItemDetailViewControllerClass().Alloc()
	rv := objc.Send[MKMapItemDetailViewController](instance.ID, objc.Sel("initWithMapItem:"), mapItem)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKMapItemDetailViewControllerWithMapItem */


// Create a map item detail view controller
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItemDetailViewController/init(mapItem:displaysMap:)
func NewMKMapItemDetailViewControllerWithMapItemDisplaysMap(mapItem IMKMapItem, displaysMap bool) MKMapItemDetailViewController {
	instance := getMKMapItemDetailViewControllerClass().Alloc()
	rv := objc.Send[MKMapItemDetailViewController](instance.ID, objc.Sel("initWithMapItem:displaysMap:"), mapItem, displaysMap)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKMapItemDetailViewControllerWithMapItemDisplaysMap */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKMapItemDetailViewController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKMapItemDetailViewController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKMapItemDetailViewController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKMapItemDetailViewController */

// The map item detail view controller’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItemDetailViewController/delegate
func (m_ MKMapItemDetailViewController) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The map item detail view controller’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItemDetailViewController/delegate
func (m_ MKMapItemDetailViewController) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The map item to display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItemDetailViewController/mapItem
func (m_ MKMapItemDetailViewController) MapItem() IMKMapItem {
	rv := objc.Send[MKMapItem](m_.ID, objc.Sel("mapItem"))
	return rv
}/* debug [instance_properties/getter]: mapItem */


// The map item to display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItemDetailViewController/mapItem
func (m_ MKMapItemDetailViewController) SetMapItem(value IMKMapItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMapItem:"), value)
}/* debug [instance_properties/setter]: mapItem */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKMapItemDetailViewController */


