// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSBackgroundExtensionView */


/* debug [class_header]: Header for NSBackgroundExtensionView */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BackgroundExtensionView */
// An interface definition for the [BackgroundExtensionView] class.
type IBackgroundExtensionView interface {
	IView
	
/* debug [class_interface_properties]: Properties for BackgroundExtensionView */
	// properties:
	AutomaticallyPlacesContentView() bool
	SetAutomaticallyPlacesContentView(value bool)
	ContentView() IView
	SetContentView(value IView)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BackgroundExtensionView */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BackgroundExtensionView */
// Alloc allocates a new instance without initialization.
func (bc _BackgroundExtensionViewClass) Alloc() BackgroundExtensionView {
	rv := objc.Send[BackgroundExtensionView](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BackgroundExtensionView */
// A view that extends content to fill its own bounds.
//
// A background extension view can be laid out to extend outside the safe area, such as under the titlebar, sidebar, or inspector. By default it lays out its content to stay within the safe area, and uses modifications of the content along the edges to fill the container view.


// A view that extends content to fill its own bounds.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BackgroundExtensionView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BackgroundExtensionView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BackgroundExtensionView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BackgroundExtensionView */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BackgroundExtensionView */

// Controls the automatic safe area placement of the within the container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBackgroundExtensionView/automaticallyPlacesContentView
func (b_ BackgroundExtensionView) AutomaticallyPlacesContentView() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("automaticallyPlacesContentView"))
	return rv
}/* debug [instance_properties/getter]: automaticallyPlacesContentView */


// Controls the automatic safe area placement of the within the container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBackgroundExtensionView/automaticallyPlacesContentView
func (b_ BackgroundExtensionView) SetAutomaticallyPlacesContentView(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAutomaticallyPlacesContentView:"), value)
}/* debug [instance_properties/setter]: automaticallyPlacesContentView */


// The content view to extend to fill the .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBackgroundExtensionView/contentView
func (b_ BackgroundExtensionView) ContentView() IView {
	rv := objc.Send[View](b_.ID, objc.Sel("contentView"))
	return rv
}/* debug [instance_properties/getter]: contentView */


// The content view to extend to fill the .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBackgroundExtensionView/contentView
func (b_ BackgroundExtensionView) SetContentView(value IView) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setContentView:"), value)
}/* debug [instance_properties/setter]: contentView */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSBackgroundExtensionView */



