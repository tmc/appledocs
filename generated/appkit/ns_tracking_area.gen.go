// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSTrackingArea */


/* debug [class_header]: Header for NSTrackingArea */
// The class instance for the [TrackingArea] class.
var (
	TrackingAreaClass     _TrackingAreaClass
	TrackingAreaClassOnce sync.Once
)

func getTrackingAreaClass() _TrackingAreaClass {
	TrackingAreaClassOnce.Do(func() {
		TrackingAreaClass = _TrackingAreaClass{objc.GetClass("NSTrackingArea")}
	})
	return TrackingAreaClass
}

type _TrackingAreaClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TrackingArea */
// An interface definition for the [TrackingArea] class.
type ITrackingArea interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TrackingArea */
	// properties:
	Options() TrackingAreaOptions /* not a class type */
	Rect() Rect /* not a class type */
	SetRect(value Rect /* not a class type */)
	UserInfo() objectivec.IObject
	SetUserInfo(value objectivec.IObject)
	VisibleRect() Rect /* not a class type */
	SetVisibleRect(value Rect /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TrackingArea */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TrackingArea */
// Alloc allocates a new instance without initialization.
func (tc _TrackingAreaClass) Alloc() TrackingArea {
	rv := objc.Send[TrackingArea](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TrackingAreaClass) New() TrackingArea {
	rv := objc.Send[TrackingArea](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TrackingArea) Init() TrackingArea {
	rv := objc.Send[TrackingArea](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TrackingArea) Autorelease() TrackingArea {
	rv := objc.Send[TrackingArea](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTrackingArea creates a new TrackingArea instance.
func NewTrackingArea() TrackingArea {
	return getTrackingAreaClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TrackingArea */
// A region of a view that generates mouse-tracking and cursor-update events when the pointer is over that region.
//
// When creating a tracking-area object, you specify a rectangle (in the view’s coordinate system), an owning object, and one or more options, along with (optionally) a dictionary of data. After it’s created, you add the tracking-area object to a view using the method. Depending on the options specified, the owner of the tracking area receives , , , and messages when the mouse cursor enters, moves within, and leaves the tracking area. Currently the tracking area is restricted to rectangles. An object belongs to its view rather than to its window. Consequently, you can add and remove tracking rectangles without needing to worry if the view has been added to a window. In addition, this design makes it possible for the AppKit to compute the geometry of tracking areas automatically when a view moves and, in some cases, when a view changes size. Using , you can configure the scope of activity for mouse tracking. There are four options: The tracking area is active only when the view is first responder. The tracking area is active when the view is in the key window. The tracking area is active when the application is active. The tracking area is active always (even when the application is inactive). Other options for objects include specifying that the tracking area should be synchronized with the visible rectangle of the view ( ) and for generating and : events when the mouse is dragged. Other methods related to objects (in addition to ) include and . Views can override the latter method to recompute and replace their objects in certain situations, such as a change in the size of the .


// A region of a view that generates mouse-tracking and cursor-update events when the pointer is over that region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTrackingArea
type TrackingArea struct {
	objectivec.Object
}

// TrackingAreaFrom constructs a [TrackingArea] from an unsafe.Pointer.
//
// A region of a view that generates mouse-tracking and cursor-update events when the pointer is over that region.
func TrackingAreaFrom(ptr unsafe.Pointer) TrackingArea {
	return TrackingArea{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TrackingArea *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TrackingArea */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TrackingArea */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TrackingArea */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TrackingArea */

// The options specified for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTrackingArea/options-swift.property
func (t_ TrackingArea) Options() TrackingAreaOptions /* not a class type */ {
	rv := objc.Send[TrackingAreaOptions](t_.ID, objc.Sel("options"))
	return rv
}/* debug [instance_properties/getter]: options */


// The rectangle defining the area encompassed by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstrackingarea/rect
func (t_ TrackingArea) Rect() Rect /* not a class type */ {
	rv := objc.Send[Rect](t_.ID, objc.Sel("rect"))
	return rv
}/* debug [instance_properties/getter]: rect */


// The rectangle defining the area encompassed by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstrackingarea/rect
func (t_ TrackingArea) SetRect(value Rect /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRect:"), value)
}/* debug [instance_properties/setter]: rect */


// The dictionary containing the data associated with the receiver when it was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstrackingarea/userinfo
func (t_ TrackingArea) UserInfo() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](t_.ID, objc.Sel("userInfo"))
	return rv
}/* debug [instance_properties/getter]: userInfo */


// The dictionary containing the data associated with the receiver when it was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstrackingarea/userinfo
func (t_ TrackingArea) SetUserInfo(value objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUserInfo:"), value)
}/* debug [instance_properties/setter]: userInfo */


// The portion of the view that isn’t clipped by its superviews.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/visiblerect
func (t_ TrackingArea) VisibleRect() Rect /* not a class type */ {
	rv := objc.Send[Rect](t_.ID, objc.Sel("visibleRect"))
	return rv
}/* debug [instance_properties/getter]: visibleRect */


// The portion of the view that isn’t clipped by its superviews.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/visiblerect
func (t_ TrackingArea) SetVisibleRect(value Rect /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setVisibleRect:"), value)
}/* debug [instance_properties/setter]: visibleRect */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTrackingArea */



