// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [TrackingArea] class.
type ITrackingArea interface {
	objectivec.IObject
	Options() TrackingAreaOptions
	Owner() objc.ID
	Rect() coregraphics.CGRect
	UserInfo() unsafe.Pointer
	VisibleRect() coregraphics.CGRect
	SetVisibleRect(value coregraphics.CGRect)
}

// A region of a view that generates mouse-tracking and cursor-update events when the pointer is over that region.
//
// When creating a tracking-area object, you specify a rectangle (in the view’s coordinate system), an owning object, and one or more options, along with (optionally) a dictionary of data. After it’s created, you add the tracking-area object to a view using the method. Depending on the options specified, the owner of the tracking area receives , , , and messages when the mouse cursor enters, moves within, and leaves the tracking area. Currently the tracking area is restricted to rectangles. An object belongs to its view rather than to its window. Consequently, you can add and remove tracking rectangles without needing to worry if the view has been added to a window. In addition, this design makes it possible for the AppKit to compute the geometry of tracking areas automatically when a view moves and, in some cases, when a view changes size. Using , you can configure the scope of activity for mouse tracking. There are four options: The tracking area is active only when the view is first responder. The tracking area is active when the view is in the key window. The tracking area is active when the application is active. The tracking area is active always (even when the application is inactive). Other options for objects include specifying that the tracking area should be synchronized with the visible rectangle of the view ( ) and for generating and : events when the mouse is dragged. Other methods related to objects (in addition to ) include and . Views can override the latter method to recompute and replace their objects in certain situations, such as a change in the size of the .
//
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

// Alloc allocates a new instance without initialization.
func (tc _TrackingAreaClass) Alloc() TrackingArea {
	rv := objc.Send[TrackingArea](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Initializes and returns an object defining a region of a view to receive mouse-tracking events, mouse-moved events, cursor-update events, or possibly all these events.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTrackingArea/init(rect:options:owner:userInfo:)
func NewTrackingAreaWithRectOptionsOwnerUserInfo(rect coregraphics.CGRect, options TrackingAreaOptions, owner objectivec.IObject, userInfo unsafe.Pointer) TrackingArea {
	instance := getTrackingAreaClass().Alloc()
	rv := objc.Send[TrackingArea](instance.ID, objc.Sel("initWithRect:options:owner:userInfo:"), rect, options, owner, userInfo)
	rv.Autorelease()
	return rv
}


// The options specified for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTrackingArea/options-swift.property
func (t_ TrackingArea) Options() TrackingAreaOptions {
	rv := objc.Send[TrackingAreaOptions](t_.ID, objc.Sel("options"))
	return rv
}

// The object owning the receiver, which is the recipient of mouse-tracking, mouse-movement, and cursor-update messages.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTrackingArea/owner
func (t_ TrackingArea) Owner() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("owner"))
	return rv
}

// The rectangle defining the area encompassed by the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTrackingArea/rect
func (t_ TrackingArea) Rect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](t_.ID, objc.Sel("rect"))
	return rv
}

// The dictionary containing the data associated with the receiver when it was created.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTrackingArea/userInfo
func (t_ TrackingArea) UserInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("userInfo"))
	return rv
}

// The portion of the view that isn’t clipped by its superviews.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/visiblerect
func (t_ TrackingArea) VisibleRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](t_.ID, objc.Sel("visibleRect"))
	return rv
}


// SetVisibleRect sets the value of the visibleRect property.
// The portion of the view that isn’t clipped by its superviews.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/visiblerect
func (t_ TrackingArea) SetVisibleRect(value coregraphics.CGRect) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setVisibleRect:"), value)
}


