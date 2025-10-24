// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

/* debug [class.gen.go]: Generating class WebFrameView */


/* debug [class_header]: Header for WebFrameView */
// The class instance for the [WebFrameView] class.
var (
	WebFrameViewClass     _WebFrameViewClass
	WebFrameViewClassOnce sync.Once
)

func getWebFrameViewClass() _WebFrameViewClass {
	WebFrameViewClassOnce.Do(func() {
		WebFrameViewClass = _WebFrameViewClass{objc.GetClass("WebFrameView")}
	})
	return WebFrameViewClass
}

type _WebFrameViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for WebFrameView */
// An interface definition for the [WebFrameView] class.
type IWebFrameView interface {
	appkit.IView
	
/* debug [class_interface_properties]: Properties for WebFrameView */
	// properties:
	AllowsScrolling() bool
	SetAllowsScrolling(value bool)
	CanPrintHeadersAndFooters() bool
	DocumentView() unsafe.Pointer
	DocumentViewShouldHandlePrint() bool
	WebFrame() IWebFrame
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for WebFrameView */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for WebFrameView */
// Alloc allocates a new instance without initialization.
func (wc _WebFrameViewClass) Alloc() WebFrameView {
	rv := objc.Send[WebFrameView](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (wc _WebFrameViewClass) New() WebFrameView {
	rv := objc.Send[WebFrameView](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WebFrameView) Init() WebFrameView {
	rv := objc.Send[WebFrameView](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WebFrameView) Autorelease() WebFrameView {
	rv := objc.Send[WebFrameView](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWebFrameView creates a new WebFrameView instance.
func NewWebFrameView() WebFrameView {
	return getWebFrameViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for WebFrameView */
// objects and their subviews display the web content contained in a frame. You never create instances of directly— objects create and manage a hierarchy of objects, one for each frame. objects use a scroll view whose document view conforms to the protocol.


// objects and their subviews display the web content contained in a frame. You never create instances of directly— objects create and manage a hierarchy of objects, one for each frame. objects use a scroll view whose document view conforms to the protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebFrameView
type WebFrameView struct {
	appkit.View
}

// WebFrameViewFrom constructs a [WebFrameView] from an unsafe.Pointer.
//
// objects and their subviews display the web content contained in a frame. You never create instances of directly— objects create and manage a hierarchy of objects, one for each frame. objects use a scroll view whose document view conforms to the protocol.
func WebFrameViewFrom(ptr unsafe.Pointer) WebFrameView {
	return WebFrameView{
		View: appkit.ViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for WebFrameView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for WebFrameView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for WebFrameView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for WebFrameView */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for WebFrameView */

// A Boolean that indicates whether the frame view should allow users to scroll.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebFrameView/allowsScrolling
func (w_ WebFrameView) AllowsScrolling() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("allowsScrolling"))
	return rv
}/* debug [instance_properties/getter]: allowsScrolling */


// A Boolean that indicates whether the frame view should allow users to scroll.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebFrameView/allowsScrolling
func (w_ WebFrameView) SetAllowsScrolling(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAllowsScrolling:"), value)
}/* debug [instance_properties/setter]: allowsScrolling */


// A Boolean value indicating whether the receiver can print headers and footers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebFrameView/canPrintHeadersAndFooters
func (w_ WebFrameView) CanPrintHeadersAndFooters() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("canPrintHeadersAndFooters"))
	return rv
}/* debug [instance_properties/getter]: canPrintHeadersAndFooters */


// The subview that displays the web content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebFrameView/documentView
func (w_ WebFrameView) DocumentView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("documentView"))
	return rv
}/* debug [instance_properties/getter]: documentView */


// A Boolean value indicating whether the document view should handle a print operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebFrameView/documentViewShouldHandlePrint
func (w_ WebFrameView) DocumentViewShouldHandlePrint() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("documentViewShouldHandlePrint"))
	return rv
}/* debug [instance_properties/getter]: documentViewShouldHandlePrint */


// The web frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebFrameView/webFrame
func (w_ WebFrameView) WebFrame() IWebFrame {
	rv := objc.Send[WebFrame](w_.ID, objc.Sel("webFrame"))
	return rv
}/* debug [instance_properties/getter]: webFrame */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class WebFrameView */



