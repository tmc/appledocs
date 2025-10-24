// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKOverlayView */


/* debug [class_header]: Header for MKOverlayView */
// The class instance for the [MKOverlayView] class.
var (
	MKOverlayViewClass     _MKOverlayViewClass
	MKOverlayViewClassOnce sync.Once
)

func getMKOverlayViewClass() _MKOverlayViewClass {
	MKOverlayViewClassOnce.Do(func() {
		MKOverlayViewClass = _MKOverlayViewClass{objc.GetClass("MKOverlayView")}
	})
	return MKOverlayViewClass
}

type _MKOverlayViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKOverlayView */
// An interface definition for the [MKOverlayView] class.
type IMKOverlayView interface {
	IView
	
/* debug [class_interface_properties]: Properties for MKOverlayView */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKOverlayView */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKOverlayView */
// Alloc allocates a new instance without initialization.
func (mc _MKOverlayViewClass) Alloc() MKOverlayView {
	rv := objc.Send[MKOverlayView](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKOverlayViewClass) New() MKOverlayView {
	rv := objc.Send[MKOverlayView](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKOverlayView) Init() MKOverlayView {
	rv := objc.Send[MKOverlayView](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKOverlayView) Autorelease() MKOverlayView {
	rv := objc.Send[MKOverlayView](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKOverlayView creates a new MKOverlayView instance.
func NewMKOverlayView() MKOverlayView {
	return getMKOverlayViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKOverlayView */
// Defines the basic behavior associated with all overlay views.
//
// An overlay view provides the visual representation of an overlay object—that is, an object that conforms to the protocol. This class defines the drawing infrastructure used by the map view but does not do any actual drawing. Subclasses are expected to override the method in order to draw the contents of the overlay view. The Map Kit framework provides several concrete instances of overlay views. Specifically, it provides overlay views for each of the concrete overlay objects. You can use one of these existing overlay views or define your own subclass if you want to draw the overlay contents differently. In iOS 7 and later, use the class to display overlays instead.


// Defines the basic behavior associated with all overlay views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayView
type MKOverlayView struct {
	View
}

// MKOverlayViewFrom constructs a [MKOverlayView] from an unsafe.Pointer.
//
// Defines the basic behavior associated with all overlay views.
func MKOverlayViewFrom(ptr unsafe.Pointer) MKOverlayView {
	return MKOverlayView{
		View: ViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKOverlayView */

// Initializes and returns the overlay view and associates it with the specified overlay object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayView/initWithOverlay:
func NewMKOverlayViewWithOverlay(overlay unsafe.Pointer) MKOverlayView {
	instance := getMKOverlayViewClass().Alloc()
	rv := objc.Send[MKOverlayView](instance.ID, objc.Sel("initWithOverlay:"), overlay)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKOverlayViewWithOverlay */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKOverlayView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKOverlayView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKOverlayView */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKOverlayView */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKOverlayView */


