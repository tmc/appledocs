// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

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

// An interface definition for the [MKOverlayView] class.
type IMKOverlayView interface {
	appkit.IView
}

// Defines the basic behavior associated with all overlay views.
//
// An overlay view provides the visual representation of an overlay object—that is, an object that conforms to the protocol. This class defines the drawing infrastructure used by the map view but does not do any actual drawing. Subclasses are expected to override the method in order to draw the contents of the overlay view. The Map Kit framework provides several concrete instances of overlay views. Specifically, it provides overlay views for each of the concrete overlay objects. You can use one of these existing overlay views or define your own subclass if you want to draw the overlay contents differently. In iOS 7 and later, use the class to display overlays instead.


// Defines the basic behavior associated with all overlay views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayView
type MKOverlayView struct {
	appkit.View
}

// MKOverlayViewFrom constructs a [MKOverlayView] from an unsafe.Pointer.
//
// Defines the basic behavior associated with all overlay views.
func MKOverlayViewFrom(ptr unsafe.Pointer) MKOverlayView {
	return MKOverlayView{
		View: appkit.ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MKOverlayViewClass) Alloc() MKOverlayView {
	rv := objc.Send[MKOverlayView](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




