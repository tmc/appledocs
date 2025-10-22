// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [BackgroundExtensionView] class.
var (
	BackgroundExtensionViewClass     _BackgroundExtensionViewClass
	BackgroundExtensionViewClassOnce sync.Once
)

func getBackgroundExtensionViewClass() _BackgroundExtensionViewClass {
	BackgroundExtensionViewClassOnce.Do(func() {
		BackgroundExtensionViewClass = _BackgroundExtensionViewClass{objc.GetClass("NSBackgroundExtensionView")}
	})
	return BackgroundExtensionViewClass
}

type _BackgroundExtensionViewClass struct {
	class objc.Class
}

// An interface definition for the [BackgroundExtensionView] class.
type IBackgroundExtensionView interface {
	IView
	AutomaticallyPlacesContentView() bool
	SetAutomaticallyPlacesContentView(value bool)
	ContentView() NSView
	SetContentView(value IView)
}

// A view that extends content to fill its own bounds.
//
// A background extension view can be laid out to extend outside the safe area, such as under the titlebar, sidebar, or inspector. By default it lays out its content to stay within the safe area, and uses modifications of the content along the edges to fill the container view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBackgroundExtensionView
type BackgroundExtensionView struct {
	View
}

// BackgroundExtensionViewFrom constructs a [BackgroundExtensionView] from an unsafe.Pointer.
//
// A view that extends content to fill its own bounds.
func BackgroundExtensionViewFrom(ptr unsafe.Pointer) BackgroundExtensionView {
	return BackgroundExtensionView{
		View: ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (bc _BackgroundExtensionViewClass) Alloc() BackgroundExtensionView {
	rv := objc.Send[BackgroundExtensionView](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BackgroundExtensionViewClass) New() BackgroundExtensionView {
	rv := objc.Send[BackgroundExtensionView](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BackgroundExtensionView) Init() BackgroundExtensionView {
	rv := objc.Send[BackgroundExtensionView](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BackgroundExtensionView) Autorelease() BackgroundExtensionView {
	rv := objc.Send[BackgroundExtensionView](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBackgroundExtensionView creates a new BackgroundExtensionView instance.
func NewBackgroundExtensionView() BackgroundExtensionView {
	return getBackgroundExtensionViewClass().New()
}


// Controls the automatic safe area placement of the within the container.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBackgroundExtensionView/automaticallyPlacesContentView
func (b_ BackgroundExtensionView) AutomaticallyPlacesContentView() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("automaticallyPlacesContentView"))
	return rv
}


// SetAutomaticallyPlacesContentView sets the value of the automaticallyPlacesContentView property.
// Controls the automatic safe area placement of the within the container.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBackgroundExtensionView/automaticallyPlacesContentView
func (b_ BackgroundExtensionView) SetAutomaticallyPlacesContentView(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAutomaticallyPlacesContentView:"), value)
}

// The content view to extend to fill the .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBackgroundExtensionView/contentView
func (b_ BackgroundExtensionView) ContentView() NSView {
	rv := objc.Send[NSView](b_.ID, objc.Sel("contentView"))
	return rv
}


// SetContentView sets the value of the contentView property.
// The content view to extend to fill the .

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBackgroundExtensionView/contentView
func (b_ BackgroundExtensionView) SetContentView(value IView) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setContentView:"), value)
}



