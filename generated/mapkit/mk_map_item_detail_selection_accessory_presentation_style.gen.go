// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKMapItemDetailSelectionAccessoryPresentationStyle */


/* debug [class_header]: Header for MKMapItemDetailSelectionAccessoryPresentationStyle */
// The class instance for the [MKMapItemDetailSelectionAccessoryPresentationStyle] class.
var (
	MKMapItemDetailSelectionAccessoryPresentationStyleClass     _MKMapItemDetailSelectionAccessoryPresentationStyleClass
	MKMapItemDetailSelectionAccessoryPresentationStyleClassOnce sync.Once
)

func getMKMapItemDetailSelectionAccessoryPresentationStyleClass() _MKMapItemDetailSelectionAccessoryPresentationStyleClass {
	MKMapItemDetailSelectionAccessoryPresentationStyleClassOnce.Do(func() {
		MKMapItemDetailSelectionAccessoryPresentationStyleClass = _MKMapItemDetailSelectionAccessoryPresentationStyleClass{objc.GetClass("MKMapItemDetailSelectionAccessoryPresentationStyle")}
	})
	return MKMapItemDetailSelectionAccessoryPresentationStyleClass
}

type _MKMapItemDetailSelectionAccessoryPresentationStyleClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKMapItemDetailSelectionAccessoryPresentationStyle */
// An interface definition for the [MKMapItemDetailSelectionAccessoryPresentationStyle] class.
type IMKMapItemDetailSelectionAccessoryPresentationStyle interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKMapItemDetailSelectionAccessoryPresentationStyle */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKMapItemDetailSelectionAccessoryPresentationStyle */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKMapItemDetailSelectionAccessoryPresentationStyle */
// Alloc allocates a new instance without initialization.
func (mc _MKMapItemDetailSelectionAccessoryPresentationStyleClass) Alloc() MKMapItemDetailSelectionAccessoryPresentationStyle {
	rv := objc.Send[MKMapItemDetailSelectionAccessoryPresentationStyle](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKMapItemDetailSelectionAccessoryPresentationStyleClass) New() MKMapItemDetailSelectionAccessoryPresentationStyle {
	rv := objc.Send[MKMapItemDetailSelectionAccessoryPresentationStyle](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKMapItemDetailSelectionAccessoryPresentationStyle) Init() MKMapItemDetailSelectionAccessoryPresentationStyle {
	rv := objc.Send[MKMapItemDetailSelectionAccessoryPresentationStyle](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKMapItemDetailSelectionAccessoryPresentationStyle) Autorelease() MKMapItemDetailSelectionAccessoryPresentationStyle {
	rv := objc.Send[MKMapItemDetailSelectionAccessoryPresentationStyle](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKMapItemDetailSelectionAccessoryPresentationStyle creates a new MKMapItemDetailSelectionAccessoryPresentationStyle instance.
func NewMKMapItemDetailSelectionAccessoryPresentationStyle() MKMapItemDetailSelectionAccessoryPresentationStyle {
	return getMKMapItemDetailSelectionAccessoryPresentationStyleClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKMapItemDetailSelectionAccessoryPresentationStyle */
// The type of map item detail accessory presentation to use.


// The type of map item detail accessory presentation to use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKSelectionAccessory/MapItemDetailPresentationStyle
type MKMapItemDetailSelectionAccessoryPresentationStyle struct {
	objectivec.Object
}

// MKMapItemDetailSelectionAccessoryPresentationStyleFrom constructs a [MKMapItemDetailSelectionAccessoryPresentationStyle] from an unsafe.Pointer.
//
// The type of map item detail accessory presentation to use.
func MKMapItemDetailSelectionAccessoryPresentationStyleFrom(ptr unsafe.Pointer) MKMapItemDetailSelectionAccessoryPresentationStyle {
	return MKMapItemDetailSelectionAccessoryPresentationStyle{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKMapItemDetailSelectionAccessoryPresentationStyle *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKMapItemDetailSelectionAccessoryPresentationStyle */

// An appropriate presentation style will be chosen automatically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItemDetailSelectionAccessoryPresentationStyle/automaticWithPresentationViewController:
func (mc _MKMapItemDetailSelectionAccessoryPresentationStyleClass) AutomaticWithPresentationViewController(presentationViewController appkit.ViewController) MKMapItemDetailSelectionAccessoryPresentationStyle {
	rv := objc.Send[MKMapItemDetailSelectionAccessoryPresentationStyle](objc.ID(mc.class), objc.Sel("automaticWithPresentationViewController:"), presentationViewController)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AutomaticWithPresentationViewController) */


// Show map item detail as an annotation callout on the map
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItemDetailSelectionAccessoryPresentationStyle/calloutWithCalloutStyle:
func (mc _MKMapItemDetailSelectionAccessoryPresentationStyleClass) CalloutWithCalloutStyle(style MKMapItemDetailSelectionAccessoryCalloutStyle) MKMapItemDetailSelectionAccessoryPresentationStyle {
	rv := objc.Send[MKMapItemDetailSelectionAccessoryPresentationStyle](objc.ID(mc.class), objc.Sel("calloutWithCalloutStyle:"), style)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CalloutWithCalloutStyle) */


// Show map item detail by presenting a sheet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKSelectionAccessory/MapItemDetailPresentationStyle/sheet(presentedFrom:)
func (mc _MKMapItemDetailSelectionAccessoryPresentationStyleClass) SheetPresentedFromViewController(viewController appkit.ViewController) MKMapItemDetailSelectionAccessoryPresentationStyle {
	rv := objc.Send[MKMapItemDetailSelectionAccessoryPresentationStyle](objc.ID(mc.class), objc.Sel("sheetPresentedFromViewController:"), viewController)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SheetPresentedFromViewController) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKMapItemDetailSelectionAccessoryPresentationStyle */

// Show map item detail as an annotation callout on the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKSelectionAccessory/MapItemDetailPresentationStyle/callout
func (mc _MKMapItemDetailSelectionAccessoryPresentationStyleClass) Callout() MKMapItemDetailSelectionAccessoryPresentationStyle {
	rv := objc.Send[MKMapItemDetailSelectionAccessoryPresentationStyle](objc.ID(mc.class), objc.Sel("callout"))
	return rv
}/* debug [class_properties_class/property]: callout */

// Display a small “Open in Apple Maps” link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKSelectionAccessory/MapItemDetailPresentationStyle/openInMaps
func (mc _MKMapItemDetailSelectionAccessoryPresentationStyleClass) OpenInMaps() MKMapItemDetailSelectionAccessoryPresentationStyle {
	rv := objc.Send[MKMapItemDetailSelectionAccessoryPresentationStyle](objc.ID(mc.class), objc.Sel("openInMaps"))
	return rv
}/* debug [class_properties_class/property]: openInMaps */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKMapItemDetailSelectionAccessoryPresentationStyle */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKMapItemDetailSelectionAccessoryPresentationStyle */

// Show map item detail as an annotation callout on the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKSelectionAccessory/MapItemDetailPresentationStyle/callout
func (m_ MKMapItemDetailSelectionAccessoryPresentationStyle) Callout() IMKMapItemDetailSelectionAccessoryPresentationStyle {
	rv := objc.Send[MKMapItemDetailSelectionAccessoryPresentationStyle](m_.ID, objc.Sel("callout"))
	return rv
}/* debug [instance_properties/getter]: callout */


// Display a small “Open in Apple Maps” link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKSelectionAccessory/MapItemDetailPresentationStyle/openInMaps
func (m_ MKMapItemDetailSelectionAccessoryPresentationStyle) OpenInMaps() IMKMapItemDetailSelectionAccessoryPresentationStyle {
	rv := objc.Send[MKMapItemDetailSelectionAccessoryPresentationStyle](m_.ID, objc.Sel("openInMaps"))
	return rv
}/* debug [instance_properties/getter]: openInMaps */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKMapItemDetailSelectionAccessoryPresentationStyle */



